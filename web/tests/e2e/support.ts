import { expect, type Page } from '@playwright/test';
import { HouseholdService } from '../../src/lib/gen/casa/v1/household_pb';
import { mockRpc } from './api-mock';

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

/**
 * A personal-mode browser with count unsent commands: the local database as the app's Dexie
 * schema version 1 makes it (IndexedDB version 10), seeded from a page where no app code runs.
 */
export async function seedUnsent(page: Page, count: number): Promise<void> {
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

/**
 * A personal-mode browser reopening Casa (plan decision 5): the local database exists and the
 * session check finds a session, so the app registers the service worker, which it does in
 * personal mode only (SEC-165 (b)). Call before the page opens the app.
 */
export async function personalBrowser(page: Page): Promise<void> {
	await seedUnsent(page, 0);
	await mockRpc(page, HouseholdService.method.listMemberships, () => ({ memberships: [] }));
}

/** Safari on iPhone or iPad in a tab: navigator.standalone exists and is false (08 §25.6). */
export async function safariTab(page: Page): Promise<void> {
	await page.addInitScript(() => {
		Object.defineProperty(Navigator.prototype, 'standalone', { get: () => false });
	});
}

export type Choice = 'personal' | 'shared';

/** 08 §25.6 LY-91: the button for each device mode, word for word. */
export const choiceButton: Record<Choice, string> = {
	personal: 'Perangkat pribadiku',
	shared: 'Perangkat bersama atau pinjaman'
};

/** Every path the page shows from now on, client-side navigations included. */
export function visitedPaths(page: Page): string[] {
	const paths: string[] = [];
	page.on('framenavigated', (frame) => {
		if (frame === page.mainFrame()) paths.push(new URL(frame.url()).pathname);
	});
	return paths;
}
