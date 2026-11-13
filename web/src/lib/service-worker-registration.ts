// The one place that registers the service worker and the one Trusted Types policy for it,
// casa-sw (SEC-167 (c)). Lint rejects serviceWorker.register and trustedTypes.createPolicy
// anywhere else, so the policy can wrap nothing but this call.

/** Same-origin path, no query or fragment; the emergency worker is served here too (SEC-169 (e)). */
const serviceWorkerUrl = '/service-worker.js';

// TypeScript's DOM library has no Trusted Types yet; only what this module uses.
interface ScriptUrlPolicy {
	createScriptURL(input: string): object;
}
interface TrustedTypes {
	createPolicy(name: 'casa-sw', rules: { createScriptURL(input: string): string }): ScriptUrlPolicy;
}

/** Made once when the module loads, never exported. Absent in Safari 17 to 25. */
const policy = (globalThis as { trustedTypes?: TrustedTypes }).trustedTypes?.createPolicy(
	'casa-sw',
	{
		createScriptURL(input) {
			if (input !== serviceWorkerUrl) {
				throw new TypeError('casa-sw: only the service worker URL is allowed');
			}
			return input;
		}
	}
);

/** Registers the app shell service worker with fixed options (SEC-167 (c); ADR-030 item 9). */
export async function registerServiceWorker(): Promise<void> {
	// A TrustedScriptURL is what register() requires under Trusted Types; the DOM typings
	// only know strings, and the object stringifies to the same URL.
	const scriptUrl = policy
		? (policy.createScriptURL(serviceWorkerUrl) as unknown as string)
		: serviceWorkerUrl;
	await navigator.serviceWorker.register(scriptUrl, { type: 'module', scope: '/' });
}
