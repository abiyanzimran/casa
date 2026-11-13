import { expect, test } from '@playwright/test';

// 08 §25.6 LY-97, word for word: two paragraphs.
const ly97 = [
	'Browser ini belum bisa menjalankan Casa dengan aman. Perbarui browser ini, atau buka Casa di browser lain yang lebih baru.',
	'Bila Casa terbuka di dalam aplikasi lain atau di jendela pribadi, buka di browser biasa.'
];

// One init script per required feature (SEC-160 (e); ADR-030 item 4 and the lead's ruling 3
// of 13 November 2026, plus CSS light-dark()), each making that feature read as missing before the app starts.
const removals = [
	() => Object.defineProperty(window, 'indexedDB', { value: undefined }),
	() => Object.defineProperty(Navigator.prototype, 'serviceWorker', { get: () => undefined }),
	() => Object.defineProperty(Navigator.prototype, 'storage', { get: () => undefined }),
	() => Object.defineProperty(Crypto.prototype, 'subtle', { get: () => undefined }),
	() => Object.defineProperty(Crypto.prototype, 'randomUUID', { value: undefined }),
	() => Object.defineProperty(Navigator.prototype, 'locks', { get: () => undefined }),
	() => Object.defineProperty(window, 'BroadcastChannel', { value: undefined }),
	() => Object.defineProperty(CSS, 'supports', { value: () => false })
];

test('SEC-160 (e): without any one required feature only the unsupported-browser sentence shows', async ({
	browser
}) => {
	for (const removal of removals) {
		const page = await browser.newPage();
		await page.addInitScript(removal);
		await page.goto('/');
		await expect(page.locator('main p')).toHaveText(ly97);
		await expect(page.locator('input, form, button')).toHaveCount(0); // "Tanpa isian masuk"
		await expect(page).toHaveTitle('Casa');
		await expect(page.locator('h1')).toHaveClass(/\bjudul-layar\b/);
		await page.close();
	}
});

test('SEC-160 (e): a supported browser does not show the unsupported-browser sentence', async ({
	page
}) => {
	await page.goto('/');
	await expect(page.locator('h1')).toHaveCount(1);
	for (const paragraph of ly97) await expect(page.getByText(paragraph)).toHaveCount(0);
});

test('LY-97: the page follows the system text size and has no browser body margin (08 §25.4, §19.13.6)', async ({
	browser
}) => {
	const page = await browser.newPage();
	await page.addInitScript(removals[0]);
	await page.goto('/');
	await expect(page.locator('meta[name="text-scale"]')).toHaveAttribute('content', 'scale');
	expect(await page.evaluate(() => getComputedStyle(document.body).margin)).toBe('0px');
	await page.close();
});
