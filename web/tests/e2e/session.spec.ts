import { Code } from '@connectrpc/connect';
import { expect, test, type Page } from '@playwright/test';
import { ErrorCode } from '../../src/lib/gen/casa/v1/error_pb';
import { HouseholdService } from '../../src/lib/gen/casa/v1/household_pb';
import { failUnmocked, mockRpc, rpcError } from './api-mock';

// page.route does not see requests a service worker handles (plan decision 19; W2).
test.use({ serviceWorkers: 'block' });

// 08 §25.6 LY-94 and §7.6, word for word.
const revoked = 'Kamu sudah keluar di perangkat ini, dan data Casa di browser ini sudah dihapus.';
const dropped1 = '1 catatan yang belum terkirim dari perangkat ini ikut terhapus.';
const expired =
	'Waktu masukmu di perangkat ini sudah habis. Masuk lagi dengan akun yang sama; catatan yang belum terkirim tetap tersimpan dan terkirim sesudah kamu masuk.';
const sharedEnded = 'Kamu sudah keluar otomatis. Catatanmu tidak tersimpan di perangkat ini.';
const unreachable = 'Belum tersambung. Rumah tanggamu tampil begitu tersambung lagi.';

const listMemberships = HouseholdService.method.listMemberships;
const signedIn = { memberships: [] };

/** The paper surface of the session check (08 §25.7): no sentence, no spinner. */
const checking = (page: Page) => page.locator('[aria-busy="true"]');

/**
 * A personal-mode browser with count unsent commands: the local database as the app's Dexie
 * schema version 1 makes it (IndexedDB version 10), seeded from a page where no app code runs.
 */
async function seedUnsent(page: Page, count: number): Promise<void> {
	await page.goto('/manifest.webmanifest');
	await page.evaluate(
		(count) =>
			new Promise<void>((resolve, reject) => {
				const open = indexedDB.open('casa', 10);
				open.onupgradeneeded = () => {
					open.result.createObjectStore('outbox', { keyPath: 'idempotencyKey' });
					open.result.createObjectStore('meta', { keyPath: 'key' });
				};
				open.onerror = () => {
					reject(open.error ?? new Error('open failed'));
				};
				open.onsuccess = () => {
					const db = open.result;
					const transaction = db.transaction('outbox', 'readwrite');
					for (let i = 0; i < count; i++) {
						transaction.objectStore('outbox').add({
							idempotencyKey: crypto.randomUUID(),
							kind: 'test_command',
							payload: new Uint8Array([1]),
							protective: false,
							createdAt: 0,
							releaseAt: 0
						});
					}
					transaction.oncomplete = () => {
						db.close();
						resolve();
					};
					transaction.onerror = () => {
						reject(transaction.error ?? new Error('seed failed'));
					};
				};
			}),
		count
	);
}

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

/** The wipe_pending row a wipe writes before it deletes anything (SEC-172 (a)). */
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
					transaction.objectStore('meta').put({ key: 'wipe_pending', value: true });
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

/** Counts the wipes that reached the cache step (each calls caches.keys() once). */
async function countCacheWipes(page: Page): Promise<() => Promise<number>> {
	await page.addInitScript(() => {
		const keys = caches.keys.bind(caches);
		Object.assign(window, { cacheWipes: 0 });
		caches.keys = () => {
			(window as unknown as { cacheWipes: number }).cacheWipes++;
			return keys();
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
	const calls = await mockRpc(page, listMemberships, async () => {
		if (++answered === 1) await new Promise((resolve) => setTimeout(resolve, 10_000));
		return signedIn;
	});
	await page.goto('/');
	await expect(checking(page)).toBeVisible();
	await expect(checking(page)).toHaveText('');
	await expect(page.locator('progress, [role="progressbar"]')).toHaveCount(0);

	await expect(checking(page)).toHaveCount(0, { timeout: 3500 });
	await expect(page.getByText(unreachable)).toBeVisible();
	await expect(page).toHaveTitle('Casa');
	await expect(page.locator('h1')).toHaveText('Casa');

	await context.setOffline(true);
	await context.setOffline(false);
	await expect.poll(() => calls.length).toBe(2);
	expect(calls[1]?.requestContext?.userInitiated).toBe(false);
	await expect(page.getByText(unreachable)).toHaveCount(0);
});

test('SEC-165 (g): "Coba lagi" checks the session again as a person\'s action', async ({
	page
}) => {
	await failUnmocked(page);
	let answered = 0;
	const calls = await mockRpc(page, listMemberships, () =>
		++answered === 1 ? rpcError(Code.Unavailable) : signedIn
	);
	await page.goto('/');
	await expect(page.getByText(unreachable)).toBeVisible();
	await page.getByRole('button', { name: 'Coba lagi' }).click();
	await expect(page.getByText(unreachable)).toHaveCount(0);
	expect(calls.map((call) => call.requestContext?.userInitiated)).toEqual([false, true]);
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
		await expect(page).toHaveTitle('Casa');
		expect(await hasCasaDatabase(page)).toBe(false);
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
	expect(new URL(page.url()).pathname).toBe('/');
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
	await expect(page.getByText(unreachable)).toBeVisible(); // the re-check listener is armed

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

test('SEC-172 (a): a wipe that keeps failing reloads once, then shows the unreachable state, never LY-94', async ({
	page
}) => {
	await failUnmocked(page);
	// Every cache deletion fails, so the wipe can never finish.
	await page.addInitScript(() => {
		CacheStorage.prototype.delete = () => Promise.reject(new Error('cache deletion failed'));
	});
	let answered = 0;
	await mockRpc(page, listMemberships, () =>
		++answered === 1 ? signedIn : rpcError(Code.Unauthenticated, ErrorCode.SESSION_REVOKED)
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
	await expect(page.getByText(unreachable)).toBeVisible();
	await expect(page.getByText(revoked)).toHaveCount(0);
	expect(documentLoads).toBe(1);
	expect(new URL(page.url()).pathname).toBe('/');
});

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

test('ADR-024 item 6: a wipe left pending counts its unsent records for LY-94', async ({
	page
}) => {
	await failUnmocked(page);
	await seedUnsent(page, 1);
	await seedWipePending(page);
	await mockRpc(page, listMemberships, () =>
		rpcError(Code.Unauthenticated, ErrorCode.SESSION_REVOKED)
	);
	await page.goto('/');
	await expect(page.getByText(revoked)).toBeVisible();
	await expect(page.getByText(dropped1)).toBeVisible();
	expect(await hasCasaDatabase(page)).toBe(false);
});

test('ADR-024 item 6: after a failed cache deletion and its reload, LY-94 keeps the unsent count', async ({
	page
}) => {
	await failUnmocked(page);
	// Cache deletion fails in the first document only; the reloaded one wipes normally.
	await page.addInitScript(() => {
		const navigation = performance.getEntriesByType('navigation')[0] as
			PerformanceNavigationTiming | undefined;
		if (navigation?.type === 'reload') return;
		CacheStorage.prototype.delete = () => Promise.reject(new Error('cache deletion failed'));
	});
	await seedUnsent(page, 1);
	await seedProbeCache(page);
	await mockRpc(page, listMemberships, () =>
		rpcError(Code.Unauthenticated, ErrorCode.SESSION_REVOKED)
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
