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
	await wipeLocalData();
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
	await wipeLocalData();
	answer({ kind: 'accepted' });
	await Promise.allSettled([flushing]);
	expect(await Dexie.exists('casa')).toBe(false);
});

test('SEC-172 (a): every wipe step runs when one fails, and the failure is reported', async () => {
	await queueOne();
	cacheStorage.keys.mockRejectedValue(new Error('cache storage unavailable'));
	await expect(wipeLocalData()).rejects.toBeInstanceOf(AggregateError);
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
	const wiping = wipeLocalData();
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
	await wipeLocalData();
	expect(getSubscription).not.toHaveBeenCalled();
	vi.stubGlobal('Notification', undefined);
	await wipeLocalData();
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
