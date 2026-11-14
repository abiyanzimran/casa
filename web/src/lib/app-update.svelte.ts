// The "Perbarui" row's state and action (08 §25.7; SEC-169 (c), (d)): a version below the
// server's minimum shows the row, and a waiting service worker is applied only while no form is
// open, so an update never takes away what a person is typing.
import { errorCode } from './api';
import { ErrorCode } from './gen/casa/v1/error_pb';

export const upgrade: { required: boolean } = $state({ required: false });

/** Sets upgrade.required from a failed call; every RPC caller reaches this (W4). */
export function noteUpgradeRequired(error: unknown): void {
	if (errorCode(error) === ErrorCode.CLIENT_UPGRADE_REQUIRED) upgrade.required = true;
}

let openCount = 0;
let deferred = false;

export function hasOpenForms(): boolean {
	return openCount > 0;
}

/** A form opened; the returned function closes it. A deferred update runs at zero. */
export function openForm(): () => void {
	openCount++;
	let closed = false;
	return () => {
		if (closed) return;
		closed = true;
		openCount--;
		if (openCount === 0 && deferred) {
			deferred = false;
			updateNow().catch((error: unknown) => {
				console.error(error);
			});
		}
	};
}

/**
 * Looks for a newer version and, if one waits and no form is open, activates it and reloads
 * once it controls the page. With a form open it is deferred until the last form closes.
 */
export async function applyUpdate(
	openForms: () => boolean
): Promise<'reloaded' | 'deferred' | 'none'> {
	const registration = await navigator.serviceWorker.getRegistration();
	if (!registration) return 'none';
	await registration.update();
	const waiting = registration.waiting;
	// ponytail: a version still installing is found on the next tap, not awaited here.
	if (!waiting) return 'none';
	if (openForms()) {
		deferred = true;
		return 'deferred';
	}
	const controlled = new Promise<void>((resolve) => {
		navigator.serviceWorker.addEventListener(
			'controllerchange',
			() => {
				resolve();
			},
			{ once: true }
		);
	});
	waiting.postMessage({ type: 'SKIP_WAITING' });
	// A page no worker controls gets no controllerchange; it reloads at once.
	if (navigator.serviceWorker.controller) await controlled;
	location.reload();
	return 'reloaded';
}

let inFlight = false;

/**
 * What "Perbarui" does (08 §25.7): applies a waiting version, and with none (shared mode has no
 * service worker, or the new one is still installing) reloads, which loads from the network.
 * With a form open either waits until the last one closes. One run at a time.
 */
export async function updateNow(): Promise<void> {
	if (inFlight) return;
	inFlight = true;
	try {
		if ((await applyUpdate(hasOpenForms)) !== 'none') return;
		if (hasOpenForms()) deferred = true;
		else location.reload();
	} finally {
		inFlight = false;
	}
}
