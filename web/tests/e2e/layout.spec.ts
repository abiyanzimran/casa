import { Code } from '@connectrpc/connect';
import { expect, test, type Page } from '@playwright/test';
import { HouseholdService } from '../../src/lib/gen/casa/v1/household_pb';
import { failUnmocked, mockRpc, rpcError } from './api-mock';

// page.route does not see requests a service worker handles (plan decision 19; W2).
test.use({ serviceWorkers: 'block' });

/**
 * Every step-1 screen; each later task adds its rows. The base mocks answer the session check
 * with no session; setup adds what the screen needs and arrives at path, or path is opened.
 */
const screens: { path: string; setup?: (page: Page) => Promise<void> }[] = [
	{
		path: '/pasang',
		setup: async (page) => {
			await page.addInitScript(() => {
				Object.defineProperty(Navigator.prototype, 'standalone', { get: () => false });
			});
			await page.goto('/pasang');
		}
	},
	{ path: '/masuk/mode' },
	{
		path: '/masuk',
		setup: async (page) => {
			// The device mode is memory only, so LY-01 is reached through LY-91.
			await page.goto('/masuk/mode');
			await page.getByRole('button', { name: 'Perangkat pribadiku', exact: true }).click();
		}
	}
];

// 08 §25.3 width classes: Ringkas (320, 390), Sedang (744), Lebar (1133, 1280); 320 also with
// 200% text (WCAG 1.4.4, 1.4.10).
const widths: { width: number; textScale?: string }[] = [
	{ width: 320 },
	{ width: 320, textScale: '200%' },
	{ width: 390 },
	{ width: 744 },
	{ width: 1133 },
	{ width: 1280 }
];

for (const { path, setup } of screens) {
	test.describe(path, () => {
		test('08 §25.5: every step-1 screen reflows at 320 px with 200% text and keeps one column of at most 640 px', async ({
			page
		}) => {
			await failUnmocked(page);
			await mockRpc(page, HouseholdService.method.listMemberships, () =>
				rpcError(Code.Unauthenticated)
			);
			await (setup ?? ((page: Page) => page.goto(path)))(page);
			await expect(page).toHaveURL(new RegExp(`${path}$`));
			await expect(page.locator('h1')).toBeFocused();
			for (const { width, textScale } of widths) {
				const at = `${String(width)} px${textScale ? ` with ${textScale} text` : ''}`;
				await page.setViewportSize({ width, height: 800 });
				// CSSOM, not a style attribute, which the production CSP blocks (SEC-166).
				await page.evaluate((scale) => {
					document.documentElement.style.fontSize = scale;
				}, textScale ?? '');
				const layout = await page.evaluate(() => ({
					scrollWidth: document.documentElement.scrollWidth,
					clientWidth: document.documentElement.clientWidth,
					column: document.querySelector('main')?.getBoundingClientRect().width ?? 0,
					h1: document.querySelectorAll('h1').length,
					// Every button and field shows its label as text (WCAG 2.5.3).
					unlabelled: [...document.querySelectorAll('main button, main input, main select')]
						.filter((element) => {
							if (element instanceof HTMLButtonElement) return !element.innerText.trim();
							const labels = (element as HTMLInputElement).labels;
							return !labels || ![...labels].some((label) => label.innerText.trim());
						})
						.map((element) => element.id || element.tagName)
				}));
				expect(layout.scrollWidth, `horizontal scroll at ${at}`).toBeLessThanOrEqual(
					layout.clientWidth
				);
				expect(layout.h1, `one h1 at ${at}`).toBe(1);
				expect(layout.column, `column at ${at}`).toBeLessThanOrEqual(640);
				expect(layout.unlabelled, `visible labels at ${at}`).toEqual([]);
				await expect(page.locator('h1')).toBeInViewport();
			}
		});
	});
}
