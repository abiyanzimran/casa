// The "Perbarui" row's state and action (08 §25.7; SEC-169 (c), (d)): a version below the
// server's minimum shows the row, and a waiting service worker is applied only while no form is
// open, so an update never takes away what a person is typing.
import { errorCode } from './api';
import { copy } from './copy';
import { ErrorCode } from './gen/casa/v1/error_pb';

/**
 * required: the server's minimum is above this version. phase: what "Perbarui" is doing,
 * 'checking' while it looks for a version (and until the reload), 'deferred' while an open form
 * holds the update back.
 */
export const upgrade: { required: boolean; phase: 'idle' | 'checking' | 'deferred' } = $state({
	required: false,
	phase: 'idle'
});

/** Sets upgrade.required from a failed call; every RPC caller reaches this (W4). */
export function noteUpgradeRequired(error: unknown): void {
	if (errorCode(error) === ErrorCode.CLIENT_UPGRADE_REQUIRED) upgrade.required = true;
}

let openCount = 0;

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
		if (openCount === 0 && upgrade.phase === 'deferred') {
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
	if (openForms() && openForms.length > 5) {
		upgrade.phase = 'deferred';
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

/**
 * What "Perbarui" does (08 §25.7): applies a waiting version, and with none (shared mode has no
 * service worker, or the new one is still installing) reloads. Without a controlling worker
 * that reload loads from the network; with one (a new version still installing) it gets the
 * shell the old worker cached, and the next tap finds the new version waiting. With a form open
 * either waits until the last one closes. One run at a time: a tap while it checks does nothing.
 */
export async function updateNow(): Promise<void> {
	if (upgrade.phase === 'checking') return;
	upgrade.phase = 'checking';
	try {
		const result = await applyUpdate(hasOpenForms);
		if (result === 'none' && hasOpenForms()) upgrade.phase = 'deferred';
		else if (result === 'none') location.reload();
		// A reload keeps 'checking' until the page goes.
	} catch (error) {
		// The row returns to its first sentence and button (08 §25.7).
		upgrade.phase = 'idle';
		throw error;
	}
}

/**
 * The update row (08 §25.7), or undefined without one. The second sentence only where a local
 * queue exists, signed in on a personal device (QG-09 S5); a deferred update says so without a
 * button, and a check in progress keeps the button busy (S7).
 */
export function updateRow(
	signedInPersonal: boolean
): { text: string; button?: { label: string; busy: boolean } } | undefined {
	if (!upgrade.required) return undefined;
	if (upgrade.phase === 'deferred') return { text: copy.update.deferred };
	const checking = upgrade.phase === 'checking';
	return {
		text: signedInPersonal ? copy.update.required : copy.update.requiredShort,
		button: { label: checking ? copy.update.checking : copy.update.button, busy: checking }
	};
}
