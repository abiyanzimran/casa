// @ts-check
// Emergency service worker (SEC-169 (e); ADR-030 item 9). Served by the server at
// /service-worker.js, the normal worker's URL, when a broken worker is installed (runbook,
// FS-1 step 7); the browser picks it up as an ordinary update. Plain JavaScript with no
// imports, so it works even when the build does not.
//
// - install: skipWaiting(). This is outside SEC-169 (c): it carries no app code, so it cannot
//   swap code under an open form; it only stops the broken worker serving.
// - activate: deletes every cache (Cache Storage holds only the app shell, SEC-169 (b)).
// - No fetch handler: every request goes to the network.
// - Never navigates or reloads pages, never touches IndexedDB (forms and the outbox stay),
//   and never unregisters, because the push subscription belongs to the registration
//   (SEC-176 (f)). The normal worker returns through an ordinary update at the same URL.
// - push: always shows a notification (SEC-176 (b)); Safari revokes permission otherwise.

const scope = /** @type {ServiceWorkerGlobalScope} */ (/** @type {unknown} */ (self));

scope.addEventListener('install', (event) => {
	event.waitUntil(scope.skipWaiting());
});

scope.addEventListener('activate', (event) => {
	event.waitUntil(
		scope.caches.keys().then((names) => Promise.all(names.map((name) => scope.caches.delete(name))))
	);
});

/**
 * The class text from the push payload, if it has one (SEC-50, SEC-176 (a)).
 * @param {PushMessageData | null} data
 * @returns {NotificationOptions}
 */
function notificationOptions(data) {
	try {
		/** @type {unknown} */
		const payload = data?.json();
		if (typeof payload === 'object' && payload !== null && 'text' in payload) {
			const text = payload.text;
			if (typeof text === 'string') return { body: text };
		}
	} catch {
		// An unreadable payload still gets a notification, without text.
	}
	return {};
}

scope.addEventListener('push', (event) => {
	event.waitUntil(scope.registration.showNotification('Casa', notificationOptions(event.data)));
});
