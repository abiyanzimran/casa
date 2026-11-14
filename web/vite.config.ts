import adapter from '@sveltejs/adapter-static';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import pkg from './package.json' with { type: 'json' };

// SvelteKit 3 takes its options here; svelte.config.js is no longer read.
export default defineConfig({
	// The client build version every request carries (ADR-030 item 9; SEC-169 (d)).
	define: { __CASA_VERSION__: JSON.stringify(pkg.version) },
	plugins: [
		sveltekit({
			compilerOptions: {
				// Runes mode for project code; libraries keep their own mode.
				runes: ({ filename }) =>
					filename.split(/[/\\]/).includes('node_modules') ? undefined : true
			},
			// SEC-167 (c): registered only by registerServiceWorker() in
			// src/lib/service-worker-registration.ts, behind the casa-sw policy.
			serviceWorker: { register: false },
			adapter: adapter({
				// SPA mode (ADR-030 item 2): every unknown path gets the app shell.
				fallback: 'index.html',
				// scripts/check-bundle.mjs builds into a temporary directory, so a
				// build with source maps never lands in build/. Unset: 'build'.
				pages: process.env.BUNDLE_CHECK_OUT || undefined
			})
		})
	],
	build: {
		// Source maps never ship in the production bundle. scripts/check-bundle.mjs
		// sets BUNDLE_CHECK=1 to read package names from them in a temp build.
		sourcemap: process.env.BUNDLE_CHECK === '1' ? 'hidden' : false
	},
	// Unit tests and golden vectors (ADR-030 item 4; 14 §2.2). Playwright owns tests/e2e.
	test: {
		include: ['src/**/*.test.ts'],
		environment: 'node',
		restoreMocks: true,
		unstubGlobals: true
	}
});
