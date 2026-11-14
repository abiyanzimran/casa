import { expect, type Page } from '@playwright/test';

/** Resolves once the page's service worker is active (SEC-169). */
export async function serviceWorkerReady(page: Page): Promise<void> {
	await page.evaluate(async () => {
		await navigator.serviceWorker.ready;
	});
}

/**
 * Records the effective directive of every CSP violation in each document the page opens,
 * from before any page script runs. Read the list with violations().
 */
export async function collectViolations(page: Page): Promise<void> {
	await page.addInitScript(() => {
		const seen: string[] = [];
		Object.assign(window, { violations: seen });
		document.addEventListener('securitypolicyviolation', (event) => {
			seen.push(event.effectiveDirective);
		});
	});
}

/** The violations collectViolations() recorded in the current document. */
export function violations(page: Page): Promise<string[]> {
	return page.evaluate(() => (window as unknown as { violations: string[] }).violations);
}

/** Every leaf of a parsed JSON value that is not a number. */
function nonNumbers(value: unknown): unknown[] {
	if (typeof value === 'number') return [];
	if (typeof value === 'object' && value !== null) return Object.values(value).flatMap(nonNumbers);
	return [value];
}

/**
 * The SEC-179 storage check with the allowlist of 10.5 §5 (note of 14 November 2026, N3):
 * localStorage is empty, and sessionStorage holds only SvelteKit's navigation keys, whose
 * values carry numbers only, never a string from data. Returns the sessionStorage keys.
 */
export async function expectNavigationStateOnly(page: Page): Promise<string[]> {
	const { local, session } = await page.evaluate(() => ({
		local: Object.keys(localStorage),
		session: Object.entries(sessionStorage) as [string, string][]
	}));
	expect(local).toEqual([]);
	const allowed = ['sveltekit:history-info', 'sveltekit:snapshot', 'sveltekit:navigation-snapshot'];
	for (const [key, value] of session) {
		expect(allowed, `sessionStorage key ${key}`).toContain(key);
		expect(nonNumbers(JSON.parse(value)), `sessionStorage ${key}`).toEqual([]);
	}
	return session.map(([key]) => key);
}
