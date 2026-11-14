import { expect, test } from '@playwright/test';
import { HouseholdService } from '../../src/lib/gen/casa/v1/household_pb';
import { failUnmocked, mockRpc } from './api-mock';

// SEC-175 (a), ADR-030 item 9: the visible name is "Casa".
test('SEC-175: app shell has the Casa title and one h1', async ({ page }) => {
	// Past the session gate (SEC-172 (b)): a signed-in session, so the h1 is the route's.
	await failUnmocked(page);
	await mockRpc(page, HouseholdService.method.listMemberships, () => ({ memberships: [] }));
	await page.goto('/');
	await expect(page.locator('[aria-busy="true"]')).toHaveCount(0);
	await expect(page).toHaveTitle('Casa');
	await expect(page.locator('h1')).toHaveCount(1);
});

// ADR-030 item 2: SPA mode, the static fallback page serves every path.
test('ADR-030 item 2: unknown path falls back to the SPA shell', async ({ page }) => {
	await page.goto('/does-not-exist');
	await expect(page).toHaveTitle('Casa');
});
