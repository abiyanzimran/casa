import { expect, test } from 'vitest';
import { displayContext, missingFeatures, type BrowserScope } from './device';

const full = {
	indexedDB: {},
	BroadcastChannel: () => undefined,
	crypto: { subtle: {}, randomUUID: () => '' },
	navigator: { serviceWorker: {}, storage: {}, locks: {} },
	CSS: { supports: (property: string, value: string) => `${property}: ${value}` === lightDark }
};
const lightDark = 'color: light-dark(#000, #fff)';

/** A scope whose display-mode query and navigator.standalone are set only when given. */
function scope(options: {
	standaloneMedia?: boolean;
	navigatorStandalone?: boolean;
}): BrowserScope {
	return {
		matchMedia: (query) => ({
			matches: query === '(display-mode: standalone)' && options.standaloneMedia === true
		}),
		navigator:
			options.navigatorStandalone === undefined ? {} : { standalone: options.navigatorStandalone }
	};
}

test('SEC-160 (e): a browser with all eight required features misses none', () => {
	expect(missingFeatures(full)).toEqual([]);
});

test('SEC-160 (e): each required feature is reported when it is absent', () => {
	expect(missingFeatures({ ...full, indexedDB: undefined })).toEqual(['indexedDB']);
	expect(missingFeatures({ ...full, BroadcastChannel: undefined })).toEqual(['BroadcastChannel']);
	expect(missingFeatures({ ...full, crypto: { randomUUID: full.crypto.randomUUID } })).toEqual([
		'crypto.subtle'
	]);
	expect(missingFeatures({ ...full, crypto: { subtle: {} } })).toEqual(['crypto.randomUUID']);
	expect(missingFeatures({ ...full, CSS: { supports: () => false } })).toEqual(['css-light-dark']);
	expect(missingFeatures({ ...full, CSS: undefined })).toEqual(['css-light-dark']);
	for (const feature of ['serviceWorker', 'storage', 'locks'] as const) {
		expect(
			missingFeatures({ ...full, navigator: { ...full.navigator, [feature]: undefined } })
		).toEqual([feature]);
	}
});

test('SEC-171 (b): standalone display mode or navigator.standalone true is the home-screen web app', () => {
	expect(displayContext(scope({ standaloneMedia: true }))).toBe('standalone');
	expect(displayContext(scope({ navigatorStandalone: true }))).toBe('standalone');
});

test('SEC-171 (b): navigator.standalone false is a Safari tab on iPhone or iPad', () => {
	expect(displayContext(scope({ navigatorStandalone: false }))).toBe('ios-safari-tab');
});

test('SEC-171 (b): without navigator.standalone the page is an ordinary browser', () => {
	// 08 §25.6: when detection is uncertain (iPad reporting as a computer), LY-91, not LY-90.
	expect(displayContext(scope({}))).toBe('browser');
});
