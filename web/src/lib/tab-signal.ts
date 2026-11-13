// Signals every tab of this browser profile must act on: sign-out, session_revoked, leaving
// the household (F06), and the app lock (ADR-030 item 10; SEC-173 (b)). A BroadcastChannel
// also delivers to other channel objects in the sending tab, so the sender acts through the
// same listener as every other tab.

const channelName = 'casa';

export type WipeKind = 'sign_out' | 'session_revoked' | 'household_exit';

/** unsentDropped is the count for the LY-94 sentence (08 §25.6). */
export type TabSignal = { kind: WipeKind; unsentDropped: number } | { kind: 'app_lock' };

const wipeKinds: readonly unknown[] = ['sign_out', 'session_revoked', 'household_exit'];

// Messages from another app version (an old tab after an update) or anything malformed are
// not signals.
function isTabSignal(data: unknown): data is TabSignal {
	if (typeof data !== 'object' || data === null || !('kind' in data)) return false;
	if (data.kind === 'app_lock') return true;
	return (
		wipeKinds.includes(data.kind) &&
		'unsentDropped' in data &&
		Number.isSafeInteger(data.unsentDropped) &&
		(data.unsentDropped as number) >= 0
	);
}

export function postTabSignal(signal: TabSignal): void {
	const channel = new BroadcastChannel(channelName);
	channel.postMessage(signal);
	channel.close();
}

/** Calls the handler for every valid signal; returns a function that stops listening. */
export function onTabSignal(handler: (signal: TabSignal) => void): () => void {
	const channel = new BroadcastChannel(channelName);
	channel.onmessage = (event: MessageEvent<unknown>) => {
		if (isTabSignal(event.data)) handler(event.data);
	};
	return () => {
		channel.close();
	};
}
