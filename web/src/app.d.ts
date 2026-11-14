// Build-time constants set by `define` in vite.config.ts, and SvelteKit's app types.

declare global {
	/** `version` in web/package.json: the client build version (ADR-030 item 9). */
	const __CASA_VERSION__: string;

	namespace App {
		interface PageState {
			/** A revoked wipe failed and the page reloaded once (SEC-172 (a)); see +layout.svelte. */
			wipeRetried?: true;
			/** The unsent count that wipe dropped, for LY-94 after the reload (ADR-024 item 6). */
			unsentDropped?: number;
		}
	}
}

export {};
