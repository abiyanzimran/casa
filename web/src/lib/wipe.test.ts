import 'fake-indexeddb/auto';
import Dexie from 'dexie';
import { afterEach, beforeEach, expect, test, vi } from 'vitest';
import type { SendResult } from './outbox';
import type { TabSignal } from './tab-signal';

// Fresh modules per test: after a wipe, localDb() refuses until the page reloads, which new
// modules stand for here.
let localDb: typeof import('./local-db').localDb;
let createOutbox: typeof import('./outbox').createOutbox;
let onTabSignal: typeof import('./tab-signal').onTabSignal;
let broadcastWipe: typeof import('./wipe').broadcastWipe;
let wipeLocalData: typeof import('./wipe').wipeLocalData;

let cacheNames: string[] = [];
const cacheStorage = {
	keys: vi.fn(() => Promise.resolve([...cacheNames])),
	delete: vi.fn((name: string) => Promise.resolve(cacheNames.includes(name)))
};
const unsubscribe = vi.fn(() => Promise.resolve(true));
const getSubscription = vi.fn(() => Promise.resolve({ unsubscribe }));
const stops: (() => void)[] = [];

beforeEach(async () => {
	vi.resetModules();
	({ localDb } = await import('./local-db'));
	({ createOutbox } = await import('./outbox'));
	({ onTabSignal } = await import('./tab-signal'));
	({ broadcastWipe, wipeLocalData } = await import('./wipe'));
	cacheNames = ['casa-shell-v1', 'other'];
	cacheStorage.keys.mockImplementation(() => Promise.resolve([...cacheNames]));
	vi.stubGlobal('caches', cacheStorage);
	vi.stubGlobal('Notification', { permission: 'granted' });
	vi.stubGlobal('navigator', {
		locks: navigator.locks,
		serviceWorker: {
			getRegistration: () =>
				Promise.resolve({
					pushManager: { getSubscription }
				})
		}
	});
});

afterEach(async () => {
	for (const stop of stops.splice(0)) stop();
	await (await import('./local-db')).deleteLocalDb();
});

async function queueOne(): Promise<void> {
	await localDb().outbox.add({
		idempotencyKey: crypto.randomUUID(),
		kind: 'test_command',
		payload: new Uint8Array([1]),
		protective: true,
		createdAt: 0,
		releaseAt: 0
	});
}

test('SEC-172 (a): the wipe deletes the local database, every cache but the app shell, and the push subscription', async () => {
	await queueOne();
	await wipeLocalData('sign_out');
	expect(await Dexie.exists('casa')).toBe(false);
	expect(cacheStorage.delete.mock.calls).toEqual([['other']]);
	expect(unsubscribe).toHaveBeenCalledTimes(1);
});

test('SEC-172 (a): a send that finishes after the wipe does not bring the database back', async () => {
	await queueOne();
	let answer: (result: SendResult) => void = () => undefined;
	const sending = new Promise<SendResult>((resolve) => {
		answer = resolve;
	});
	const send = vi.fn(() => sending);
	const flushing = createOutbox(localDb(), send, () => 0).flush();
	await vi.waitFor(() => {
		expect(send).toHaveBeenCalled();
	});
	await wipeLocalData('sign_out');
	answer({ kind: 'accepted' });
	await Promise.allSettled([flushing]);
	expect(await Dexie.exists('casa')).toBe(false);
});

test('SEC-172 (a): every wipe step runs when one fails, and the failure is reported', async () => {
	await queueOne();
	cacheStorage.keys.mockRejectedValue(new Error('cache storage unavailable'));
	await expect(wipeLocalData('sign_out')).rejects.toBeInstanceOf(AggregateError);
	expect(await Dexie.exists('casa')).toBe(false);
	expect(unsubscribe).toHaveBeenCalledTimes(1);
});

test('SEC-172 (a): the push subscription step runs last, after the database and caches are gone', async () => {
	await queueOne();
	let releaseCaches: () => void = () => undefined;
	cacheStorage.delete.mockImplementationOnce(
		() =>
			new Promise<boolean>((resolve) => {
				releaseCaches = () => {
					resolve(true);
				};
			})
	);
	let databaseAtPush: boolean | undefined;
	getSubscription.mockImplementationOnce(async () => {
		databaseAtPush = await Dexie.exists('casa');
		return { unsubscribe };
	});
	const wiping = wipeLocalData('sign_out');
	await vi.waitFor(() => {
		expect(cacheStorage.delete).toHaveBeenCalled();
	});
	await new Promise((resolve) => setTimeout(resolve, 20));
	expect(getSubscription).not.toHaveBeenCalled();
	releaseCaches();
	await wiping;
	expect(getSubscription).toHaveBeenCalledTimes(1);
	expect(databaseAtPush).toBe(false);
});

test('SEC-172 (a): without notification permission there is no push subscription to look up', async () => {
	vi.stubGlobal('Notification', { permission: 'default' });
	await wipeLocalData('sign_out');
	expect(getSubscription).not.toHaveBeenCalled();
	vi.stubGlobal('Notification', undefined);
	await wipeLocalData('sign_out');
	expect(getSubscription).not.toHaveBeenCalled();
});

test('SEC-164 (c): broadcastWipe counts unsent commands and tells every tab', async () => {
	await queueOne();
	await queueOne();
	const received: TabSignal[] = [];
	stops.push(onTabSignal((signal) => received.push(signal)));
	expect(await broadcastWipe('session_revoked')).toBe(2);
	await vi.waitFor(() => {
		expect(received).toEqual([{ kind: 'session_revoked', unsentDropped: 2 }]);
	});
});

test('SEC-172 (a): broadcastWipe without a local database creates none and reports 0', async () => {
	expect(await broadcastWipe('sign_out')).toBe(0);
	expect(await indexedDB.databases()).toEqual([]);
});

test('SEC-172 (a): a database deletion that fails leaves the wipe_pending marker for the next open', async () => {
	await queueOne();
	const { wipePending } = await import('./local-db');
	expect(await wipePending()).toBeUndefined();
	vi.spyOn(Dexie, 'delete').mockRejectedValueOnce(new Error('deletion failed'));
	await expect(wipeLocalData('session_revoked')).rejects.toThrow('local wipe incomplete');
	// The marker keeps the kind, so the next open knows a revoked wipe is unfinished (I2).
	expect(await wipePending()).toBe('session_revoked');

	// The next wipe (fresh modules: a new page) removes the database, marker included.
	vi.resetModules();
	await (await import('./wipe')).wipeLocalData('sign_out');
	expect(await Dexie.exists('casa')).toBe(false);
	expect(await (await import('./local-db')).wipePending()).toBeUndefined();
});

test('SEC-165 (b): the wipe marks nothing without a local database', async () => {
	await wipeLocalData('sign_out');
	expect(await Dexie.exists('casa')).toBe(false);
});

test('SEC-172 (a): broadcastWipe still tells every tab when the queue cannot be counted, as unknown (0)', async () => {
	await queueOne();
	const { LocalDb } = await import('./local-db');
	// Dexie's shared Table prototype, so every connection's count fails (iOS connection loss).
	const table = Object.getPrototypeOf(Object.getPrototypeOf(new LocalDb().outbox)) as {
		count: () => Promise<number>;
	};
	vi.spyOn(table, 'count').mockRejectedValue(new Error('connection lost'));
	const logged = vi.spyOn(console, 'error').mockImplementation(() => {});
	const received: TabSignal[] = [];
	stops.push(onTabSignal((signal) => received.push(signal)));
	expect(await broadcastWipe('session_revoked')).toBe(0);
	await vi.waitFor(() => {
		expect(received).toEqual([{ kind: 'session_revoked', unsentDropped: 0 }]);
	});
	expect(logged).toHaveBeenCalled();
});

test('SEC-172 (a): the queue is counted after this page wiped once and another tab brought the database back', async () => {
	await wipeLocalData('sign_out'); // localDb() now refuses in this page
	const { LocalDb } = await import('./local-db');
	const otherTab = new LocalDb();
	await otherTab.outbox.add({
		idempotencyKey: crypto.randomUUID(),
		kind: 'test_command',
		payload: new Uint8Array([1]),
		protective: true,
		createdAt: 0,
		releaseAt: 0
	});
	otherTab.close();
	expect(await broadcastWipe('session_revoked')).toBe(1);
});
