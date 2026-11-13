import { expect, test, type Page } from '@playwright/test';
import { readdirSync } from 'node:fs';
import { join, relative, sep } from 'node:path';

/** Every pathname in every Cache Storage cache of the page's origin. */
function cachedUrls(page: Page): Promise<string[]> {
	return page.evaluate(async () => {
		const urls: string[] = [];
		for (const name of await caches.keys()) {
			const cache = await caches.open(name);
			for (const request of await cache.keys()) urls.push(new URL(request.url).pathname);
		}
		return urls;
	});
}

/**
 * The app shell as built (SEC-169 (b)): every file in build/ except the service worker
 * itself, the version poll file, source maps, and the licence and readme texts. The SPA
 * fallback page index.html is cached under /.
 */
function expectedShell(): string[] {
	const root = 'build';
	return readdirSync(root, { recursive: true, withFileTypes: true })
		.filter((entry) => entry.isFile())
		.map((entry) => '/' + relative(root, join(entry.parentPath, entry.name)).split(sep).join('/'))
		.filter(
			(path) =>
				path !== '/service-worker.js' &&
				path !== '/_app/version.json' &&
				!path.endsWith('.map') &&
				!path.endsWith('.txt') &&
				!/\/(README\.md|LICENSE)$/.test(path)
		)
		.map((path) => (path === '/index.html' ? '/' : path));
}

async function serviceWorkerReady(page: Page): Promise<void> {
	await page.evaluate(async () => {
		await navigator.serviceWorker.ready;
	});
}

test('SEC-169 (b): the service worker caches exactly the app shell, in one cache named with the version', async ({
	page
}) => {
	await page.goto('/');
	await serviceWorkerReady(page);
	expect((await cachedUrls(page)).sort()).toEqual(expectedShell().sort());
	const names = await page.evaluate(() => caches.keys());
	expect(names).toHaveLength(1);
	expect(names[0]).toMatch(/^casa-shell-.+/);
});

test('SEC-169 (b): API, stream, and other requests are never cached', async ({ page }) => {
	await page.goto('/');
	await serviceWorkerReady(page);
	await page.reload(); // now controlled by the service worker
	await page.evaluate(() =>
		Promise.all([fetch('/v1/stream'), fetch(`/probe-${crypto.randomUUID()}`)])
	);
	expect(
		(await cachedUrls(page)).filter((url) => url.startsWith('/v1/') || url.startsWith('/probe-'))
	).toEqual([]);
});

test('SEC-169 (b): the service worker never answers a navigation to the API', async ({ page }) => {
	await page.goto('/');
	await serviceWorkerReady(page);
	await page.reload(); // now controlled by the service worker
	const response = await page.goto('/v1/probe');
	expect(response?.fromServiceWorker()).toBe(false);
});

test('SEC-169 (b): after the first visit the app opens without a network', async ({
	page,
	context
}) => {
	await page.goto('/');
	await serviceWorkerReady(page);
	await context.setOffline(true);
	await page.reload();
	await expect(page).toHaveTitle('Casa');
	await expect(page.locator('h1')).toHaveCount(1);
	await context.setOffline(false);
});

/**
 * Installs a second copy of the worker at another URL while a page is controlled by the
 * first. The controlled page is the manifest file, which the worker serves as itself, so no
 * app code runs and the app's own registration on load cannot race the second one. Same
 * options as registerServiceWorker(). Returns whether the copy was still waiting 500 ms
 * after it installed.
 */
async function installSecondCopy(page: Page): Promise<boolean> {
	await page.goto('/');
	await serviceWorkerReady(page);
	await page.goto('/manifest.webmanifest');
	return page.evaluate(async () => {
		const original = navigator.serviceWorker.controller;
		if (!original) throw new Error('page not controlled');
		// eslint-disable-next-line no-restricted-syntax -- the test installs a second copy of the worker (SEC-167 (c) allows this only in tests)
		const probe = await navigator.serviceWorker.register('/service-worker.js?probe=1', {
			type: 'module',
			scope: '/'
		});
		const worker = probe.installing ?? probe.waiting;
		if (!worker) throw new Error('no new worker');
		await new Promise<void>((resolve) => {
			if (worker.state === 'installed') resolve();
			worker.addEventListener('statechange', () => {
				if (worker.state === 'installed') resolve();
			});
		});
		// Give an automatic skipWaiting() the chance to happen; it must not.
		await new Promise((resolve) => setTimeout(resolve, 500));
		return probe.waiting === worker && navigator.serviceWorker.controller === original;
	});
}

/** Tells the waiting copy to take over; returns whether it then controls the page. */
async function takeOver(page: Page): Promise<boolean> {
	return page.evaluate(async () => {
		const worker = (await navigator.serviceWorker.ready).waiting;
		if (!worker) throw new Error('nothing waiting');
		const changed = new Promise<void>((resolve) => {
			navigator.serviceWorker.addEventListener('controllerchange', () => {
				resolve();
			});
		});
		// controllerchange comes at "activating"; the activate handler is done at "activated".
		const activated = new Promise<void>((resolve) => {
			worker.addEventListener('statechange', () => {
				if (worker.state === 'activated') resolve();
			});
		});
		worker.postMessage({ type: 'SKIP_WAITING' });
		await Promise.all([changed, activated]);
		return navigator.serviceWorker.controller === worker;
	});
}

test('SEC-169 (b): activating a new version deletes every other cache', async ({ page }) => {
	await page.goto('/');
	await serviceWorkerReady(page);
	await page.evaluate(async () => {
		await (await caches.open('casa-shell-old')).put('/old', new Response('x'));
	});
	expect(await installSecondCopy(page)).toBe(true);
	expect(await takeOver(page)).toBe(true);
	const names = await page.evaluate(() => caches.keys());
	expect(names).toHaveLength(1);
	expect(names).not.toContain('casa-shell-old');
});

test('SEC-169 (c): a new service worker waits until it is told to take over', async ({ page }) => {
	expect(await installSecondCopy(page)).toBe(true);
	expect(await takeOver(page)).toBe(true);
});

test('SEC-167 (c): with Trusted Types enforced the service worker registers without a violation', async ({
	page
}) => {
	// The production directive (10.5 SEC-167 (c)); the server sets it from 1.5.
	const directive = "require-trusted-types-for 'script'; trusted-types svelte-trusted-html casa-sw";
	await page.route('**/*', async (route) => {
		if (route.request().resourceType() !== 'document') {
			await route.continue();
			return;
		}
		const response = await route.fetch();
		await route.fulfill({
			response,
			headers: { ...response.headers(), 'content-security-policy': directive }
		});
	});
	await page.addInitScript(() => {
		const seen: string[] = [];
		Object.assign(window, { violations: seen });
		document.addEventListener('securitypolicyviolation', (event) => {
			seen.push(event.violatedDirective);
		});
	});
	await page.goto('/');
	await serviceWorkerReady(page);
	expect(
		await page.evaluate(() => (window as unknown as { violations: string[] }).violations)
	).toEqual([]);
	expect(await page.evaluate(() => 'trustedTypes' in window)).toBe(true);
});

test('SEC-175 (g): the web app manifest names the app Casa', async ({ request }) => {
	const manifest: unknown = await (await request.get('/manifest.webmanifest')).json();
	expect(manifest).toMatchObject({
		name: 'Casa',
		short_name: 'Casa',
		display: 'standalone',
		start_url: '/',
		scope: '/'
	});
});
