import { expect, test, type Page } from '@playwright/test';
import { createServer } from 'node:http';
import type { AddressInfo } from 'node:net';
import { collectViolations, personalBrowser, serviceWorkerReady, violations } from './support';

// Every test here runs on the embedded bundle served by cmd/webpreview with the production
// headers (plan 1.5 decision 6), so it sees the CSP the casa binary sends (SEC-166).

/** The origin the app is served from. */
function appOrigin(baseURL: string | undefined): string {
	if (!baseURL) throw new Error('baseURL is not set');
	return new URL(baseURL).origin;
}

/** The CSP directives of an answer, by name. */
function directives(csp: string): Map<string, string> {
	return new Map(
		csp
			.split(';')
			.map((directive) => directive.trim().split(' '))
			.map(([name, ...sources]) => [name, sources.join(' ')])
	);
}

// The only style attribute the CSP allows: the SvelteKit route announcer's, by its hash
// (SEC-166 note of FS-1 task 1.5).
const announcerStyle = "'unsafe-hashes' 'sha256-S8qMpvofolR8Mpjy4kQvEm7m1q8clzU4dfDH0AmvZjo='";

/** Serves the real index page with markup inserted before marker, under the real headers. */
async function serveIndexWith(page: Page, marker: string, markup: string): Promise<void> {
	await page.route('/', async (route) => {
		const response = await route.fetch();
		const html = await response.text();
		expect(html).toContain(marker);
		const headers = { ...response.headers() };
		delete headers['content-length']; // the body grows
		await route.fulfill({ response, headers, body: html.replace(marker, markup + marker) });
	});
}

test('SEC-166: the app loads under the production CSP without a violation', async ({ page }) => {
	await collectViolations(page);
	await personalBrowser(page); // the service worker registers in personal mode only
	const response = await page.goto('/');
	const csp = response?.headers()['content-security-policy'] ?? '';
	const policy = directives(csp);
	expect(policy.get('require-trusted-types-for')).toBe("'script'");
	expect(policy.get('script-src')).toMatch(/^'self' 'sha256-[A-Za-z0-9+/]+=*'$/);
	expect(policy.get('style-src')).toBe("'self'");
	expect(policy.get('style-src-attr')).toBe(announcerStyle);
	expect(csp).not.toContain("'unsafe-inline'");
	await serviceWorkerReady(page);
	await expect(page.locator('h1')).toHaveCount(1);
	// The announcer stays visually hidden. Both engines apply its cloned style even when the
	// CSP blocks it, so the empty violation list is what proves the hash allows it.
	const announcer = page.locator('#svelte-announcer');
	await expect(announcer).toHaveCSS('position', 'absolute');
	await expect(announcer).toHaveCSS('width', '1px');
	expect(await violations(page)).toEqual([]);

	await page.goto('/does-not-exist');
	await expect(page.locator('h1')).toHaveCount(1);
	await expect(announcer).toHaveCSS('width', '1px');
	expect(await violations(page)).toEqual([]);
});

test('SEC-166: an inline script injected into the page does not run', async ({ page }) => {
	await collectViolations(page);
	await serveIndexWith(page, '</head>', '<script>window.injected = 1</script>');
	await page.goto('/');
	await expect(page.locator('h1')).toHaveCount(1); // the app's own hashed script still runs
	await expect.poll(() => violations(page)).toEqual(['script-src-elem']);
	expect(await page.evaluate(() => 'injected' in window)).toBe(false);
});

test('SEC-166: an injected style attribute does not apply', async ({ page }) => {
	await collectViolations(page);
	await serveIndexWith(page, '</body>', '<div id="injected" style="width: 7px"></div>');
	await page.goto('/');
	await expect(page.locator('h1')).toHaveCount(1);
	await expect.poll(() => violations(page)).toEqual(['style-src-attr']);
	expect(
		await page.locator('#injected').evaluate((element) => getComputedStyle(element).width)
	).not.toBe('7px');
});

test('SEC-162 (e): the app cannot be framed by another origin', async ({ page, baseURL }) => {
	const app = appOrigin(baseURL);
	// A real loopback server, not page.route: Chromium treats a fulfilled page as a public
	// address and would block the frame by Local Network Access before frame-ancestors is
	// ever read, so the test would pass without the header.
	const framer = createServer((_, response) => {
		response.setHeader('content-type', 'text/html');
		response.end(`<!doctype html><title>Other</title><iframe src="${app}/"></iframe>`);
	});
	await new Promise<void>((resolve) => framer.listen(0, '127.0.0.1', resolve));
	const { port } = framer.address() as AddressInfo;
	const appRequests: string[] = [];
	page.on('request', (request) => {
		const url = new URL(request.url());
		if (url.origin === app) appRequests.push(url.pathname);
	});
	try {
		// localhost, while the app is on 127.0.0.1: another site, not only another port.
		await page.goto(`http://localhost:${String(port)}/`); // waits for the frame's load event too
	} finally {
		framer.close();
	}
	// The frame asked for the page and the browser refused the answer: the document was never
	// parsed, so not even its preloads and stylesheet were requested, and no script ran.
	expect(appRequests).toEqual(['/']);
	const frame = page.frames().find((candidate) => candidate !== page.mainFrame());
	expect(frame).toBeDefined();
	expect(await frame?.title()).not.toBe('Casa');
});

test('SEC-160 (d): the app requests nothing from another host', async ({
	page,
	context,
	baseURL
}) => {
	const app = appOrigin(baseURL);
	const urls: string[] = [];
	// The context also sees the service worker's own requests (Chromium).
	context.on('request', (request) => urls.push(request.url()));
	await personalBrowser(page); // the service worker registers in personal mode only
	await page.goto('/');
	await serviceWorkerReady(page);
	expect(urls).toContain(`${app}/service-worker.js`);
	expect(urls.filter((url) => new URL(url).origin !== app)).toEqual([]);
});
