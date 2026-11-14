// SvelteKit's init hook runs before the router reads the address and before any app code, so
// a link token leaves the address bar before the first request and before the router writes
// the address back into history (SEC-175 (c); plan decision 8).
import { displayContext } from '#lib/device.js';
import { takeFragmentToken } from '#lib/link-token.js';
import { session } from '#lib/session.svelte.js';

/** The link paths of the server that carry a token, and its kind (plan decisions 8, 13). */
const linkKinds: Record<string, 'setup' | 'verification' | undefined> = {
	'/signup': 'setup',
	'/signup/verify': 'verification'
};

export function init(): void {
	const kind = linkKinds[location.pathname];
	if (!kind) return;
	const value = takeFragmentToken();
	// In a Safari tab on iPhone or iPad these tokens are dropped unsent: a setup link shows LY-90
	// with ly90.fromSetup (08 §7.6), and a verification link LY-90 until the contract can complete
	// it without a session (SEC-171 (b) (2), Q2-6).
	if (displayContext() === 'ios-safari-tab') session.setupLinkDropped = kind === 'setup';
	else if (value) session.pendingToken = { kind, value };
}
