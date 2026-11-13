import { expect, test, type Page } from '@playwright/test';

/** Waits for the app to render, so its signal listener is installed, and for the worker. */
async function openApp(page: Page): Promise<void> {
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

test('SEC-172 (a): a session_revoked signal empties IndexedDB and non-shell caches in every open tab', async ({
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

	// The old documents are marked; the wipe ends with location.replace('/'), so a page without
	// the mark is a new document with nothing of the old one in memory (SEC-164 (c)).
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
		// Polled, because the page navigates in between.
		await expect
			.poll(() => afterWipe(page), { timeout: 10_000 })
			.toEqual({ newDocument: true, casaDatabase: false, probeCache: false, shellCaches: 1 });
	}
});
