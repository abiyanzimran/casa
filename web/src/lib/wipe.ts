// The local wipe on sign-out, session_revoked, and leaving the household (SEC-164 (c),
// SEC-172 (a); ADR-024 item 7 (b), (d)). Revoking the push subscription on the server is the
// server's part (SEC-164 (b), SEC-176 (f)).
import { APP_SHELL_CACHE_PREFIX } from './app-shell';
import { deleteLocalDb, localDb, localDbExists } from './local-db';
import { postTabSignal, type WipeKind } from './tab-signal';

async function deleteNonShellCaches(): Promise<void> {
	const names = await caches.keys();
	await Promise.all(
		names
			.filter((name) => !name.startsWith(APP_SHELL_CACHE_PREFIX))
			.map((name) => caches.delete(name))
	);
}

async function unsubscribePush(): Promise<void> {
	// A subscription needs notification permission, so without it there is none to remove;
	// iPhone and iPad Safari tabs have no Notification at all. This also keeps
	// getSubscription() out of Playwright's WebKit, where it never settles.
	if (typeof Notification === 'undefined' || Notification.permission !== 'granted') return;
	const registration = await navigator.serviceWorker.getRegistration();
	const subscription = await registration?.pushManager.getSubscription();
	await subscription?.unsubscribe();
}

/**
 * Deletes the local database, every cache but the app shell, and this browser's push
 * subscription. Every step runs even when another fails; failures are then thrown together
 * (ES §1: nothing fails silently).
 */
export async function wipeLocalData(): Promise<void> {
	const results = await Promise.allSettled([
		deleteLocalDb(),
		deleteNonShellCaches(),
		unsubscribePush()
	]);
	const failures: unknown[] = results.flatMap((result) =>
		result.status === 'rejected' ? [result.reason as unknown] : []
	);
	if (failures.length > 0) throw new AggregateError(failures, 'local wipe incomplete');
}

/**
 * Tells every tab, this one included, to wipe (SEC-173 (b)), with the number of unsent
 * commands that go with it for LY-94. Without a local database it counts 0 and creates none.
 */
export async function broadcastWipe(kind: WipeKind): Promise<number> {
	const unsentDropped = (await localDbExists()) ? await localDb().outbox.count() : 0;
	postTabSignal({ kind, unsentDropped });
	return unsentDropped;
}
