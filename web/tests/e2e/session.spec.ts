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
