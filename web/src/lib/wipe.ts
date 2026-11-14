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
	// step 4: 3 s limit, SEC-164 (c) (on getSubscription and unsubscribe, when Web Push is built).
	// A subscription needs notification permission, so without it there is none to remove;
	// iPhone and iPad Safari tabs have no Notification at all. This also keeps
	// getSubscription() out of Playwright's WebKit, where it never settles.
	if (typeof Notification === 'undefined' || Notification.permission !== 'granted') return;
	const registration = await navigator.serviceWorker.getRegistration();
	const subscription = await registration?.pushManager.getSubscription();
	await subscription?.unsubscribe();
}

/**
 * Deletes the local database and every cache but the app shell, then this browser's push
 * subscription. The push step runs last, once the others have settled, so a push step that
 * hangs never leaves data behind (SEC-164 (c), SEC-172 (a)). Every step runs even when another
 * fails; failures are then thrown together (ES §1: nothing fails silently).
 */
export async function wipeLocalData(): Promise<void> {
	const results = await Promise.allSettled([deleteLocalDb(), deleteNonShellCaches()]);
	results.push(...(await Promise.allSettled([unsubscribePush()])));
	const failures: unknown[] = results.flatMap((result) =>
		result.status === 'rejected' ? [result.reason as unknown] : []
	);
	if (failures.length > 0) throw new AggregateError(failures, 'local wipe incomplete');
}

/** The number of unsent commands; without a local database 0, and none is created. */
export async function unsentCount(): Promise<number> {
	return (await localDbExists()) ? localDb().outbox.count() : 0;
}

/**
 * Tells every tab, this one included, to wipe (SEC-173 (b)), with the number of unsent
 * commands that go with it for LY-94.
 */
export async function broadcastWipe(kind: WipeKind): Promise<number> {
	const unsentDropped = await unsentCount();
	postTabSignal({ kind, unsentDropped });
	return unsentDropped;
}
