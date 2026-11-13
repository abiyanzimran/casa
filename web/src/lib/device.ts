// Browser feature and display-mode detection (SEC-160 (e), SEC-171 (b); ADR-030 items 4, 11).

/** The browser properties these checks read; globalThis in the app, a plain object in tests. */
export interface BrowserScope {
	indexedDB?: unknown;
	BroadcastChannel?: unknown;
	crypto?: { subtle?: unknown; randomUUID?: unknown };
	navigator?: { serviceWorker?: unknown; storage?: unknown; locks?: unknown; standalone?: boolean };
	matchMedia?: (query: string) => { matches: boolean };
}

export type Feature =
	| 'indexedDB'
	| 'serviceWorker'
	| 'storage'
	| 'crypto.subtle'
	| 'crypto.randomUUID'
	| 'locks'
	| 'BroadcastChannel';

/**
 * The features the web client's controls depend on that this browser lacks (SEC-160 (e);
 * ADR-030 item 4, with BroadcastChannel and crypto.randomUUID added by the lead's ruling of
 * 13 November 2026). Checked by value, not with `in`: outside a secure context
 * navigator.serviceWorker and crypto.subtle exist as undefined.
 */
export function missingFeatures(scope: BrowserScope = globalThis): Feature[] {
	const present: Record<Feature, unknown> = {
		indexedDB: scope.indexedDB,
		serviceWorker: scope.navigator?.serviceWorker,
		storage: scope.navigator?.storage,
		'crypto.subtle': scope.crypto?.subtle,
		'crypto.randomUUID': scope.crypto?.randomUUID,
		locks: scope.navigator?.locks,
		BroadcastChannel: scope.BroadcastChannel
	};
	return (Object.keys(present) as Feature[]).filter((feature) => !present[feature]);
}

export type DisplayContext = 'standalone' | 'ios-safari-tab' | 'browser';

/**
 * Where the page runs (SEC-171 (b); ADR-030 item 11). navigator.standalone exists only in
 * Safari on iPhone and iPad, so false there means a Safari tab. Anything uncertain is an
 * ordinary browser, so no account owner is locked out (08 §25.6).
 */
export function displayContext(scope: BrowserScope = globalThis): DisplayContext {
	const standalone = scope.navigator?.standalone;
	if (scope.matchMedia?.('(display-mode: standalone)').matches || standalone === true) {
		return 'standalone';
	}
	return standalone === false ? 'ios-safari-tab' : 'browser';
}
