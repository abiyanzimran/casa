import adapter from '@sveltejs/adapter-static';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

// SvelteKit 3 takes its options here; svelte.config.js is no longer read.
export default defineConfig({
	plugins: [
		sveltekit({
			compilerOptions: {
				// Runes mode for project code; libraries keep their own mode.
				runes: ({ filename }) =>
					filename.split(/[/\\]/).includes('node_modules') ? undefined : true
			},
			adapter: adapter({
				// SPA mode (ADR-030 item 2): every unknown path gets the app shell.
				fallback: 'index.html',
				// scripts/check-bundle.mjs builds into a temporary directory, so a
				// build with source maps never lands in build/. Unset: 'build'.
				pages: process.env.BUNDLE_CHECK_OUT
			})
		})
	],
	build: {
		// Source maps never ship in the production bundle. scripts/check-bundle.mjs
		// sets BUNDLE_CHECK=1 to read package names from them in a temp build.
		sourcemap: process.env.BUNDLE_CHECK === '1' ? 'hidden' : false
	}
});
