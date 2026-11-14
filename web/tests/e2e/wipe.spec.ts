import { expect, test, type Page } from '@playwright/test';
import { personalBrowser } from './support';

/**
 * Waits for the app to render past the session check (a signed-in session in personal mode),
 * so its signal listener is installed, and for the worker.
 */
async function openApp(page: Page): Promise<void> {
	await personalBrowser(page);
	await page.goto('/');
	await expect(page.locator('h1')).toBeVisible();
	await page.evaluate(async () => {
		await navigator.serviceWorker.ready;
	});
}

/** What the wipe must leave behind: no local database, only the app shell cache. */
function storage(page: Page): Promise<{ databases: string[]; caches: string[] }> {
	return page.evaluate(async () => ({
		databases: (await indexedDB.databases()).map((db) => db.name ?? ''),
		caches: await caches.keys()
	}));
}

/** null while the page is between documents. */
async function afterWipe(page: Page) {
	try {
		const left = await storage(page);
		return {
			newDocument: await page.evaluate(() => !('oldDocument' in window)),
			casaDatabase: left.databases.includes('casa'),
			probeCache: left.caches.includes('casa-probe'),
			shellCaches: left.caches.filter((name) => name.startsWith('casa-shell')).length
		};
	} catch {
		return null;
	}
}

test('SEC-172 (a): a session_revoked signal empties IndexedDB and non-shell caches in every open tab and shows LY-94 there', async ({
	context
}) => {
	const a = await context.newPage();
	const b = await context.newPage();
	await openApp(a);
	await openApp(b);
	await a.evaluate(async () => {
		// The connection is closed again: an open raw connection would block the deletion.
		await new Promise<void>((resolve, reject) => {
			const request = indexedDB.open('casa');
			request.onsuccess = () => {
				request.result.close();
				resolve();
			};
			request.onerror = () => {
				reject(request.error ?? new Error('open failed'));
			};
		});
		await (await caches.open('casa-probe')).put('/probe', new Response('x'));
	});
	expect((await storage(a)).databases).toContain('casa');

	// The old documents are marked: a revoked session shows LY-94 in place, with memory reset and
	// no reload (plan decision 6), so the mark stays. A sign-out reloads (the test below).
	for (const page of [a, b]) {
		await page.evaluate(() => {
			Object.assign(window, { oldDocument: true });
		});
	}
	await a.evaluate(() => {
		const channel = new BroadcastChannel('casa');
		channel.postMessage({ kind: 'session_revoked', unsentDropped: 0 });
		channel.close();
	});

	for (const page of [a, b]) {
		await expect(
			page.getByText(
				'Kamu sudah keluar di perangkat ini, dan data Casa di browser ini sudah dihapus.'
			)
		).toBeVisible({ timeout: 10_000 });
		expect(await afterWipe(page)).toEqual({
			newDocument: false,
			casaDatabase: false,
			probeCache: false,
			shellCaches: 1
		});
	}
});

/**
 * Opens a raw connection that ignores versionchange: it holds the database deletion, so the
 * wipe stops in the middle and what is on screen then can be seen.
 */
function holdDeletion(page: Page): Promise<void> {
	return page.evaluate(
		() =>
			new Promise<void>((resolve, reject) => {
				const request = indexedDB.open('casa');
				request.onsuccess = () => {
					Object.assign(window, { heldConnection: request.result });
					resolve();
				};
				request.onerror = () => {
					reject(request.error ?? new Error('open failed'));
				};
			})
	);
}

function releaseDeletion(page: Page): Promise<void> {
	return page.evaluate(() => {
		(window as unknown as { heldConnection: IDBDatabase }).heldConnection.close();
	});
}

test('SEC-172 (a): a revoked session shows LY-94 only once the wipe has finished', async ({
	context
}) => {
	const a = await context.newPage();
	const b = await context.newPage();
	await openApp(a);
	await openApp(b);
	await holdDeletion(a);
	await a.evaluate(() => {
		const channel = new BroadcastChannel('casa');
		channel.postMessage({ kind: 'session_revoked', unsentDropped: 0 });
		channel.close();
	});
	const revokedSentence = b.getByText(
		'Kamu sudah keluar di perangkat ini, dan data Casa di browser ini sudah dihapus.'
	);

	await expect(b.locator('h1')).toHaveCount(0);
	await expect(revokedSentence).toHaveCount(0);
	expect((await storage(b)).databases).toContain('casa'); // the deletion is held

	await releaseDeletion(a);
	await expect(revokedSentence).toBeVisible({ timeout: 10_000 });
	expect((await storage(b)).databases).not.toContain('casa');
});

test('SEC-172 (a): every tab drops its screen before the wipe starts', async ({ context }) => {
	const a = await context.newPage();
	const b = await context.newPage();
	await openApp(a);
	await openApp(b);
	await holdDeletion(a);
	await b.evaluate(() => {
		Object.assign(window, { oldDocument: true });
	});
	await a.evaluate(() => {
		const channel = new BroadcastChannel('casa');
		channel.postMessage({ kind: 'sign_out', unsentDropped: 0 });
		channel.close();
	});

	await expect(b.locator('h1')).toHaveCount(0);
	expect(await b.evaluate(() => 'oldDocument' in window)).toBe(true); // the wipe has not ended

	await releaseDeletion(a);
	await expect
		.poll(() => afterWipe(b), { timeout: 10_000 })
		.toMatchObject({ newDocument: true, casaDatabase: false });
});
