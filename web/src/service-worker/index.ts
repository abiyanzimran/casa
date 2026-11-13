// App shell service worker (ADR-030 items 2, 9; SEC-169 (b), (c)). SvelteKit 3's own
// service worker build, no plugin (lead's ruling 1 of 13 November 2026).
//
// - The file list is fixed at build: $app/manifest plus the SPA fallback page (index.html),
//   cached under / because servers answer /index.html with a redirect or not at all.
//   Nothing is added to a cache while handling fetch, so API answers, the stream, photos,
//   and data pages are never cached.
// - One cache per version, named with the build version and filled at install with requests
//   that bypass the HTTP cache, so fixed-name files (index.html, static/) never come from an
//   older version. Activate deletes every other cache (TM-45).
// - No automatic skipWaiting() and no clients.claim(): a new version waits until the app is
//   next opened, or until the page sends SKIP_WAITING (SEC-169 (c)).
import { version } from '$app/env';
import { assets, immutable } from '$app/manifest';
import { self } from '$app/service-worker';
import { APP_SHELL_CACHE_PREFIX } from '../lib/app-shell';
// @ts-expect-error -- red proof only: ms has no types; a package off the bundle allowlist
import ms from 'ms';

// eslint-disable-next-line @typescript-eslint/no-unsafe-call -- red proof only: untyped package
const cacheName = `${APP_SHELL_CACHE_PREFIX}-${version}-${String(ms(0))}`;
const indexPage = '/';

// Licence and readme texts in static/ are not part of the shell.
const isShellAsset = (path: string) => !/(^|\/)(README\.md|LICENSE)$|\.txt$/.test(path);

const shell = new Set([
	indexPage,
	...immutable.map((file) => `/${file.path}`),
	...assets.map((file) => `/${file.path}`).filter(isShellAsset)
]);

self.addEventListener('install', (event) => {
	event.waitUntil(
		caches
			.open(cacheName)
			.then((cache) => cache.addAll([...shell].map((url) => new Request(url, { cache: 'reload' }))))
	);
});

self.addEventListener('activate', (event) => {
	event.waitUntil(
		caches
			.keys()
			.then((names) =>
				Promise.all(names.filter((name) => name !== cacheName).map((name) => caches.delete(name)))
			)
	);
});

self.addEventListener('fetch', (event) => {
	const { request } = event;
	const url = new URL(request.url);
	if (request.method !== 'GET' || url.origin !== self.location.origin) return;
	// A shell file is served as itself; any other app path is the SPA page (ADR-030 item 2).
	// /v1/ is the API and stream, never answered here.
	let path = url.pathname;
	if (!shell.has(path)) {
		if (request.mode !== 'navigate' || path.startsWith('/v1/')) return;
		path = indexPage;
	}
	event.respondWith(caches.match(path, { cacheName }).then((cached) => cached ?? fetch(request)));
});

self.addEventListener('message', (event) => {
	const data: unknown = event.data;
	if (typeof data === 'object' && data !== null && 'type' in data && data.type === 'SKIP_WAITING') {
		event.waitUntil(self.skipWaiting());
	}
});
