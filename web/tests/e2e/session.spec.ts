import { Code } from '@connectrpc/connect';
import { expect, test, type Page } from '@playwright/test';
import { ErrorCode } from '../../src/lib/gen/casa/v1/error_pb';
import { HouseholdService } from '../../src/lib/gen/casa/v1/household_pb';
import { failUnmocked, mockRpc, rpcError } from './api-mock';
import { expectNavigationStateOnly, seedUnsent } from './support';

// page.route does not see requests a service worker handles (plan decision 19; W2).
test.use({ serviceWorkers: 'block' });

// 08 §25.6 LY-94 and §7.6, word for word.
const revoked = 'Kamu sudah keluar di perangkat ini, dan data Casa di browser ini sudah dihapus.';
const dropped1 = '1 catatan yang belum terkirim dari perangkat ini ikut terhapus.';
const expired =
	'Waktu masukmu di perangkat ini sudah habis. Masuk lagi dengan akun yang sama; catatan yang belum terkirim tetap tersimpan dan terkirim sesudah kamu masuk.';
const sharedEnded = 'Kamu sudah keluar otomatis. Catatanmu tidak tersimpan di perangkat ini.';
const unreachable = 'Belum tersambung. Rumah tanggamu tampil begitu tersambung lagi.';
const unreachableServer = 'Ada gangguan di server. Coba lagi beberapa saat lagi.';
const unreachableStorage =
	'Data Casa di perangkat ini belum bisa dibersihkan. Coba lagi. Kalau tetap gagal, tutup tab ini lalu buka Casa lagi.';

const listMemberships = HouseholdService.method.listMemberships;
const signedIn = { memberships: [] };

/** The paper surface of the session check (08 §25.7): no sentence, no spinner. */
const checking = (page: Page) => page.locator('[data-session-check]');

function hasCasaDatabase(page: Page): Promise<boolean> {
	return page.evaluate(async () =>
		(await indexedDB.databases()).some((database) => database.name === 'casa')
	);
}

/** Marks the current document, so a page without the mark is a new document. */
function markDocument(page: Page): Promise<void> {
	return page.evaluate(() => {
		Object.assign(window, { oldDocument: true });
	});
}

/** null while the page is between documents. */
async function isNewDocument(page: Page): Promise<boolean | null> {
	try {
		return await page.evaluate(() => !('oldDocument' in window));
	} catch {
		return null;
	}
}

/** The wipe_pending row a revoked wipe writes before it deletes anything (SEC-172 (a)). */
function seedWipePending(page: Page): Promise<void> {
	return page.evaluate(
		() =>
			new Promise<void>((resolve, reject) => {
				const open = indexedDB.open('casa', 10);
				open.onerror = () => {
					reject(open.error ?? new Error('open failed'));
				};
				open.onsuccess = () => {
					const transaction = open.result.transaction('meta', 'readwrite');
					transaction.objectStore('meta').put({ key: 'wipe_pending', value: 'session_revoked' });
					transaction.oncomplete = () => {
						open.result.close();
						resolve();
					};
				};
			})
	);
}

/** A cache the wipe must delete, so a failing cache deletion makes the wipe fail. */
function seedProbeCache(page: Page): Promise<void> {
	return page.evaluate(async () => {
		await (await caches.open('casa-probe')).put('/probe', new Response('x'));
	});
}

/**
 * Counts the wipes that reached the cache step (each calls caches.keys() once). Patched on the
 * prototype: WebKit may collect the `caches` wrapper and make a new one, which drops a
 * property set on the instance (seen under load, losing the count).
 */
async function countCacheWipes(page: Page): Promise<() => Promise<number>> {
	await page.addInitScript(() => {
		// eslint-disable-next-line @typescript-eslint/unbound-method -- called below with the instance as this
		const keys = CacheStorage.prototype.keys;
		Object.assign(window, { cacheWipes: 0 });
		CacheStorage.prototype.keys = function (this: CacheStorage) {
			(window as unknown as { cacheWipes: number }).cacheWipes++;
			return keys.call(this);
		};
	});
	return () => page.evaluate(() => (window as unknown as { cacheWipes: number }).cacheWipes);
}

test('SEC-172 (b): the app waits at most 3 seconds for the session check, then checks again when back online', async ({
	page,
	context
}) => {
	await failUnmocked(page);
	let answered = 0;
	let firstCheckAt = 0;
	const calls = await mockRpc(page, listMemberships, async () => {
		if (++answered === 1) {
			firstCheckAt = Date.now();
			await new Promise((resolve) => setTimeout(resolve, 10_000));
		}
		return signedIn;
	});
	await page.goto('/');
	await expect(checking(page)).toBeVisible();
	await expect(checking(page)).toHaveText('');
	await expect(page.locator('progress, [role="progressbar"]')).toHaveCount(0);
	// 08 §25.7: the status row's live region is on the page from the start, empty (QG-09 S6).
	await expect(page.locator('[role="status"]')).toHaveCount(1);
	await expect(page.locator('[role="status"]')).toHaveText('');

	// The 3 s cap is measured from the request to the surface leaving, read every frame, so a
	// slow runner gets time to look without the cap growing (fix wave A M5).
	await page.waitForFunction(() => !document.querySelector('[data-session-check]'), null, {
		polling: 'raf',
		timeout: 10_000
	});
	expect(Date.now() - firstCheckAt).toBeLessThan(3500);
	// A timeout while online is the server's (08 §7.6 home.unreachableServer; QG-09 S4).
	await expect(page.getByText(unreachableServer)).toBeVisible();
	await expect(page).toHaveTitle('Casa');
	await expect(page.locator('h1')).toHaveText('Casa');
	// The focused h1 announces the sentence too (QG-09 S2).
	await expect(page.locator('h1')).toHaveAccessibleDescription(unreachableServer);
	await expect(page.locator('[role="status"]')).toHaveCount(1);

	await context.setOffline(true);
	await context.setOffline(false);
	await expect.poll(() => calls.length).toBe(2);
	expect(calls[1]?.requestContext?.userInitiated).toBe(false);
	await expect(page.getByText(unreachableServer)).toHaveCount(0);
});

test('SEC-165 (g): "Coba lagi" checks the session again as a person\'s action, busy while it checks', async ({
	page
}) => {
	await failUnmocked(page);
	const secondCheck = Promise.withResolvers<undefined>();
	let answered = 0;
	const calls = await mockRpc(page, listMemberships, async () => {
		answered++;
		if (answered === 2) await secondCheck.promise;
		return answered === 3 ? signedIn : rpcError(Code.Unavailable);
	});
	await page.goto('/');
	await expect(page.getByText(unreachableServer)).toBeVisible();
	await page.getByRole('button', { name: 'Coba lagi' }).click();
	// 08 §7.6 common.retrying (QG-09 S3): not active and busy until the answer.
	const busy = page.getByRole('button', { name: 'Mencoba…' });
	await expect(busy).toBeDisabled();
	await expect(busy).toHaveAttribute('aria-busy', 'true');
	await expect(busy).toBeFocused();
	await busy.click({ force: true });
	secondCheck.resolve(undefined);
	// The same answer: back to "Coba lagi" with focus kept.
	const retry = page.getByRole('button', { name: 'Coba lagi' });
	await expect(retry).toBeEnabled();
	await expect(retry).toBeFocused();
	await retry.click();
	await expect(page.getByText(unreachableServer)).toHaveCount(0);
	expect(calls.map((call) => call.requestContext?.userInitiated)).toEqual([false, true, true]);
});

test('SEC-172 (b): opened without a connection, the app says it is not connected', async ({
	page
}) => {
	await failUnmocked(page);
	// A fetch that never reaches the server (08 §7.6 home.unreachable; QG-09 S4).
	await page.route(`**/${listMemberships.parent.typeName}/${listMemberships.name}`, (route) =>
		route.abort('internetdisconnected')
	);
	await page.goto('/');
	await expect(page.getByText(unreachable)).toBeVisible();
	await expect(page.locator('h1')).toHaveAccessibleDescription(unreachable);
});

test('SEC-172 (a): a session_revoked answer empties every tab and shows LY-94', async ({
	context
}) => {
	const b = await context.newPage();
	await failUnmocked(b);
	await mockRpc(b, listMemberships, () => signedIn);
	await seedUnsent(b, 1);
	await b.goto('/');
	await expect(b.locator('h1')).toBeVisible();
	await markDocument(b);

	const a = await context.newPage();
	await failUnmocked(a);
	// After the wipe the server no longer knows the cookie (see the report's server note).
	let answered = 0;
	await mockRpc(a, listMemberships, () =>
		++answered === 1
			? rpcError(Code.Unauthenticated, ErrorCode.SESSION_REVOKED)
			: rpcError(Code.Unauthenticated)
	);
	await a.goto('/');

	for (const page of [a, b]) {
		await expect(page.getByText(revoked)).toBeVisible();
		await expect(page.getByText(dropped1)).toBeVisible();
		await expect(page.locator('h1')).toHaveAccessibleDescription(`${revoked} ${dropped1}`);
		await expect(page).toHaveTitle('Casa');
		await expect.poll(() => hasCasaDatabase(page)).toBe(false);
	}
	// LY-94 in place, without a reload (plan decision 6).
	expect(await isNewDocument(b)).toBe(false);

	await markDocument(a);
	await a.getByRole('button', { name: 'Masuk' }).click();
	await expect.poll(() => isNewDocument(a)).toBe(true);
	expect(new URL(a.url()).pathname).toBe('/');
});

test('SEC-161 (d): a session_expired answer shows LY-94 and keeps the local data', async ({
	page
}) => {
	await failUnmocked(page);
	await seedUnsent(page, 1);
	let answered = 0;
	await mockRpc(page, listMemberships, () =>
		++answered === 1
			? rpcError(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED)
			: rpcError(Code.Unauthenticated)
	);
	await page.goto('/');
	await expect(page.getByText(expired)).toBeVisible();
	await expect(page.getByRole('button', { name: 'Masuk', exact: true })).toBeVisible();
	const otherAccount = page.getByRole('button', { name: 'Masuk dengan akun lain' });
	await expect(otherAccount).toBeVisible();
	expect(await hasCasaDatabase(page)).toBe(true);

	// SEC-170 (a): the second button asks first, with the unsent count; "Kembali" keeps all.
	await otherAccount.click();
	const dialog = page.getByRole('dialog', { name: 'Masuk dengan akun lain?' });
	await expect(dialog).toBeVisible();
	await expect(
		dialog.getByText(
			'Data Casa di perangkat ini akan dihapus, termasuk 1 catatan yang belum terkirim.'
		)
	).toBeVisible();
	await dialog.getByRole('button', { name: 'Kembali' }).click();
	await expect(dialog).toBeHidden();
	await expect(otherAccount).toBeFocused();
	await expect(page.getByText(expired)).toBeVisible();
	expect(await hasCasaDatabase(page)).toBe(true);
	// Esc closes it the same way (08 §25.5).
	await otherAccount.click();
	await expect(dialog).toBeVisible();
	await expect(dialog).toHaveAccessibleDescription(
		'Data Casa di perangkat ini akan dihapus, termasuk 1 catatan yang belum terkirim.'
	);
	await page.keyboard.press('Escape');
	await expect(dialog).toBeHidden();
	await expect(otherAccount).toBeFocused();
	expect(await hasCasaDatabase(page)).toBe(true);

	await markDocument(page);
	await otherAccount.click();
	await dialog.getByRole('button', { name: 'Hapus dan masuk' }).click();
	await expect.poll(() => isNewDocument(page)).toBe(true);
	// A fresh document at "/", which sends a person without a session on to sign-in (LY-91).
	await expect(page).toHaveURL(/\/masuk\/mode$/);
	expect(await hasCasaDatabase(page)).toBe(false);
});

test('SEC-170 (a): with nothing unsent the other-account question only says the data is deleted', async ({
	page
}) => {
	await failUnmocked(page);
	await seedUnsent(page, 0);
	await mockRpc(page, listMemberships, () =>
		rpcError(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED)
	);
	await page.goto('/');
	await page.getByRole('button', { name: 'Masuk dengan akun lain' }).click();
	const dialog = page.getByRole('dialog', { name: 'Masuk dengan akun lain?' });
	await expect(
		dialog.getByText('Data Casa di perangkat ini akan dihapus.', { exact: true })
	).toBeVisible();
	await expect(dialog.getByText('catatan yang belum terkirim')).toHaveCount(0);
});

test('SEC-165: an expired shared session shows the shared ending', async ({ page }) => {
	await failUnmocked(page);
	await mockRpc(page, listMemberships, () =>
		rpcError(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED)
	);
	await page.goto('/');
	await expect(page.getByText(sharedEnded)).toBeVisible();
	await expect(page.getByRole('button')).toHaveText(['Masuk']);
	expect(await hasCasaDatabase(page)).toBe(false);
});

/** The boxes of the buttons in a container, in document order. */
async function buttonBoxes(container: ReturnType<Page['locator']>) {
	const boxes = [];
	for (const button of await container.getByRole('button').all())
		boxes.push(await button.boundingBox());
	return boxes.map((box) => box ?? { x: 0, y: 0, width: 0, height: 0 });
}

test('08 §25.4: LY-94 and its dialog stack their buttons upright and equally wide in Ringkas, side by side in Sedang', async ({
	page
}) => {
	await failUnmocked(page);
	await seedUnsent(page, 0);
	await mockRpc(page, listMemberships, () =>
		rpcError(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED)
	);
	await page.setViewportSize({ width: 390, height: 844 });
	await page.goto('/');
	const screen = page.locator('main');
	await expect(page.getByText(expired)).toBeVisible();
	await expect(page.locator('h1')).toHaveAccessibleDescription(expired);
	let [first, second] = await buttonBoxes(screen);
	expect(second.width).toBe(first.width);
	expect(second.x).toBe(first.x);
	expect(second.y).toBeGreaterThan(first.y + first.height);

	await page.getByRole('button', { name: 'Masuk dengan akun lain' }).click();
	const dialog = page.getByRole('dialog');
	await expect(dialog).toBeVisible();
	[first, second] = await buttonBoxes(dialog);
	expect(second.width).toBe(first.width);
	expect(second.y).toBeGreaterThan(first.y + first.height);
	await page.keyboard.press('Escape');
	await expect(dialog).toBeHidden();

	await page.setViewportSize({ width: 744, height: 1133 });
	[first, second] = await buttonBoxes(screen);
	expect(second.width).toBe(first.width);
	expect(second.y).toBe(first.y);
	await page.getByRole('button', { name: 'Masuk dengan akun lain' }).click();
	await expect(dialog).toBeVisible();
	[first, second] = await buttonBoxes(dialog);
	expect(second.width).toBe(first.width);
	expect(second.y).toBe(first.y);
});

test('08 §25.4: a pressed button shrinks a little while held', async ({ page }) => {
	await failUnmocked(page);
	await mockRpc(page, listMemberships, () =>
		rpcError(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED)
	);
	await page.goto('/');
	const button = page.getByRole('button', { name: 'Masuk' });
	await button.hover();
	await page.mouse.down();
	// 08 §19.13.7 "Tekan kendali": scale(0.97) while :active.
	await expect
		.poll(() => button.evaluate((element) => getComputedStyle(element).transform))
		.toBe('matrix(0.97, 0, 0, 0.97, 0, 0)');
	await page.mouse.move(0, 0);
	await page.mouse.up();
});

for (const mode of ['shared', 'personal'] as const) {
	test.describe(`in ${mode} mode`, () => {
		test("SEC-179: sessionStorage holds only SvelteKit's navigation numbers and localStorage nothing", async ({
			page
		}) => {
			await failUnmocked(page);
			if (mode === 'personal') await seedUnsent(page, 1);
			await mockRpc(page, listMemberships, () =>
				rpcError(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED)
			);
			await page.goto('/');
			await markDocument(page);
			// "Masuk" opens a new document; SvelteKit writes its navigation state as the old one unloads.
			await page.getByRole('button', { name: 'Masuk', exact: true }).click();
			await expect.poll(() => isNewDocument(page)).toBe(true);
			await expect(page.getByRole('button', { name: 'Masuk', exact: true })).toBeVisible();
			expect(await expectNavigationStateOnly(page)).toContain('sveltekit:history-info');
		});
	});
}

test('SEC-164 (d): a page restored from the back-forward cache checks the session first', async ({
	page
}) => {
	await failUnmocked(page);
	const restoreCheck = Promise.withResolvers<undefined>();
	let answered = 0;
	const calls = await mockRpc(page, listMemberships, async () => {
		if (++answered === 2) await restoreCheck.promise;
		return signedIn;
	});
	await page.goto('/');
	await expect(page.locator('h1')).toBeVisible();
	expect(calls).toHaveLength(1);

	await page.evaluate(() => {
		window.dispatchEvent(new PageTransitionEvent('pageshow', { persisted: true }));
	});
	await expect.poll(() => calls.length).toBe(2);
	expect(calls[1]?.requestContext?.userInitiated).toBe(false);
	// Nothing of the restored screen shows until the session is known.
	await expect(checking(page)).toBeVisible();
	await expect(page.locator('h1')).toHaveCount(0);

	restoreCheck.resolve(undefined);
	await expect(page.locator('h1')).toBeVisible();
});

/** What every tab of this profile receives when one tab gets a session_revoked answer. */
function revokedSignal(page: Page, unsentDropped: number): Promise<void> {
	return page.evaluate((unsentDropped) => {
		const channel = new BroadcastChannel('casa');
		channel.postMessage({ kind: 'session_revoked', unsentDropped });
		channel.close();
	}, unsentDropped);
}

test('ADR-024 item 6: LY-94 after a revoked session stays, with its unsent count, whatever comes later', async ({
	page
}) => {
	await failUnmocked(page);
	const cacheWipes = await countCacheWipes(page);
	await seedUnsent(page, 1);
	let answered = 0;
	const calls = await mockRpc(page, listMemberships, () =>
		++answered === 1 ? rpcError(Code.Unavailable) : signedIn
	);
	await page.goto('/');
	await expect(page.getByText(unreachableServer)).toBeVisible(); // the re-check listener is armed

	await revokedSignal(page, 1);
	await expect(page.getByText(revoked)).toBeVisible();
	await expect(page.getByText(dropped1)).toBeVisible();

	await page.evaluate(() => {
		document.dispatchEvent(new Event('visibilitychange'));
		window.dispatchEvent(new Event('online'));
		window.dispatchEvent(new PageTransitionEvent('pageshow', { persisted: true }));
	});
	await revokedSignal(page, 0); // a second answer, counted after the first wipe
	// The second wipe hides LY-94 from its start to its end, so LY-94 seen again after the
	// second wipe's cache step is the state that wipe left.
	await expect.poll(cacheWipes).toBe(2);
	await expect(page.getByText(revoked)).toBeVisible();
	await expect(page.getByText(dropped1)).toBeVisible();
	expect(calls).toHaveLength(1);
});

/**
 * The two server behaviours after a session_revoked answer (fix wave I2): today every later call
 * answers session_revoked; after 1.7d that answer expires the cookie, so every later call
 * answers Unauthenticated without a detail.
 */
const afterRevoke = [
	{
		server: 'every call session_revoked',
		later: () => rpcError(Code.Unauthenticated, ErrorCode.SESSION_REVOKED)
	},
	{ server: 'then no session (1.7d)', later: () => rpcError(Code.Unauthenticated) }
];

for (const { server, later } of afterRevoke) {
	test.describe(`server: ${server}`, () => {
		test('SEC-172 (a): a wipe that keeps failing reloads once, then shows the unreachable state, never LY-94', async ({
			page
		}) => {
			await failUnmocked(page);
			// Every cache deletion fails, so the wipe can never finish.
			await page.addInitScript(() => {
				CacheStorage.prototype.delete = () => Promise.reject(new Error('cache deletion failed'));
			});
			let answered = 0;
			const calls = await mockRpc(page, listMemberships, () =>
				++answered === 1 ? signedIn : later()
			);
			await page.goto('/');
			await expect(page.locator('h1')).toBeVisible();
			await seedProbeCache(page);
			await markDocument(page);
			let documentLoads = 0;
			page.on('load', () => {
				documentLoads++;
			});

			await revokedSignal(page, 0);
			// The first failure reloads; the reloaded page's check is revoked again, its wipe fails again.
			await expect.poll(() => isNewDocument(page)).toBe(true);
			// 08 §7.6 home.unreachableStorage: never says the data is deleted (QG-09 S4).
			await expect(page.getByText(unreachableStorage)).toBeVisible();
			await expect(page.getByText(revoked)).toHaveCount(0);
			expect(documentLoads).toBe(1);
			expect(new URL(page.url()).pathname).toBe('/');
			// One check per document: the goto with state before the reload starts no check (I3).
			expect(calls).toHaveLength(2);
		});
	});
}

test('SEC-172 (a): a wipe left pending runs on open, before the session check and any screen; a valid session then starts on a fresh page', async ({
	page
}) => {
	await failUnmocked(page);
	await seedUnsent(page, 1);
	await seedWipePending(page);
	const databaseAtCheck: boolean[] = [];
	await mockRpc(page, listMemberships, async () => {
		databaseAtCheck.push(await hasCasaDatabase(page));
		return signedIn;
	});
	let documentLoads = 0;
	page.on('load', () => {
		documentLoads++;
	});
	await page.goto('/');
	// The page that wiped reloads once, so local data starts clean (no marker the second time).
	await expect.poll(() => documentLoads).toBe(2);
	await expect(page.locator('h1')).toBeVisible();
	expect(databaseAtCheck).toEqual([false, false]);
	expect(await hasCasaDatabase(page)).toBe(false);
});

for (const { server, later } of afterRevoke) {
	test(`ADR-024 item 6: a wipe left pending counts its unsent records for LY-94 (server: ${server})`, async ({
		page
	}) => {
		await failUnmocked(page);
		await seedUnsent(page, 1);
		await seedWipePending(page);
		// The session_revoked answer that left the wipe pending came to an earlier page.
		await mockRpc(page, listMemberships, later);
		await page.goto('/');
		await expect(page.getByText(revoked)).toBeVisible();
		await expect(page.getByText(dropped1)).toBeVisible();
		expect(await hasCasaDatabase(page)).toBe(false);
	});

	test(`ADR-024 item 6: after a failed cache deletion and its reload, LY-94 keeps the unsent count (server: ${server})`, async ({
		page
	}) => {
		await failUnmocked(page);
		await seedUnsent(page, 1);
		await seedProbeCache(page);
		// Cache deletion fails in the first app document only; the reloaded one wipes normally.
		// window.name survives the reload in this tab, so it counts the documents in both engines.
		await page.addInitScript(() => {
			const documents = Number(window.name || '0') + 1;
			window.name = String(documents);
			if (documents > 1) return;
			CacheStorage.prototype.delete = () => Promise.reject(new Error('cache deletion failed'));
		});
		let answered = 0;
		await mockRpc(page, listMemberships, () =>
			++answered === 1 ? rpcError(Code.Unauthenticated, ErrorCode.SESSION_REVOKED) : later()
		);
		let documentLoads = 0;
		page.on('load', () => {
			documentLoads++;
		});
		await page.goto('/');
		await expect(page.getByText(revoked)).toBeVisible();
		await expect(page.getByText(dropped1)).toBeVisible();
		expect(documentLoads).toBe(2); // the first document's wipe failed and reloaded once
	});
}

test('SEC-172 (a): a later revoked wipe that fails keeps the truthful LY-94 already shown', async ({
	page
}) => {
	await failUnmocked(page);
	await seedUnsent(page, 1);
	await mockRpc(page, listMemberships, () =>
		rpcError(Code.Unauthenticated, ErrorCode.SESSION_REVOKED)
	);
	await page.goto('/');
	await expect(page.getByText(dropped1)).toBeVisible();
	await markDocument(page);
	await seedProbeCache(page);
	await page.evaluate(() => {
		CacheStorage.prototype.delete = () => Promise.reject(new Error('cache deletion failed'));
	});
	const failedWipe = page.waitForEvent('console', (message) => message.type() === 'error');

	await revokedSignal(page, 0);
	await failedWipe;
	await expect(page.getByText(revoked)).toBeVisible();
	await expect(page.getByText(dropped1)).toBeVisible();
	expect(await isNewDocument(page)).toBe(false);
});
