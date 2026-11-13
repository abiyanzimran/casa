import type { Page } from '@playwright/test';

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
