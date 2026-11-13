import { defineConfig, devices } from '@playwright/test';

// 127.0.0.1, not localhost: cmd/webpreview listens on the IPv4 loopback only. Both
// browsers treat it as a secure context, like localhost.
const baseURL = 'http://127.0.0.1:4173';

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
	use: { baseURL },
	projects: [
		{ name: 'chromium', use: { ...devices['Desktop Chrome'] } },
		{ name: 'webkit', use: { ...devices['Desktop Safari'] } }
	],
	// The embedded bundle under the production headers, as the casa binary serves it
	// (SEC-166; plan 1.5 decision 6), so every test runs under the production CSP.
	webServer: {
		command: 'make -C .. web-preview',
		url: baseURL,
		reuseExistingServer: false,
		timeout: 120_000 // vite build, then go build
	}
});
