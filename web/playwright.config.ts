import { defineConfig, devices } from '@playwright/test';

const port = 4173;

// Both engines always run (ADR-030 item 4): Chromium for Chrome on the
// Galaxy A55, WebKit for Safari on the iPhone and iPad. No retries: a flaky
// test is fixed, not retried (ES §6).
export default defineConfig({
	testDir: 'tests/e2e',
	// The ID checker (QG-06) reads Playwright titles from *.spec.ts only.
	testMatch: '**/*.spec.ts',
	forbidOnly: true,
	retries: 0,
	reporter: [['list'], ['html', { open: 'never' }]],
	use: { baseURL: `http://localhost:${String(port)}` },
	projects: [
		{ name: 'chromium', use: { ...devices['Desktop Chrome'] } },
		{ name: 'webkit', use: { ...devices['Desktop Safari'] } }
	],
	webServer: {
		command: `pnpm build && pnpm preview --port ${String(port)} --strictPort`,
		port,
		reuseExistingServer: false
	}
});
