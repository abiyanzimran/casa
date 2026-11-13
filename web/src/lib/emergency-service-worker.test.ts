import 'fake-indexeddb/auto';
import { expect, test, vi } from 'vitest';
import { LocalDb } from './local-db';

type Listener = (event: unknown) => void;

/** A service worker global scope that records listeners and calls (SEC-169 (e)). */
function fakeScope() {
	const listeners = new Map<string, Listener>();
	const cacheNames = ['casa-shell-v1', 'other'];
	const scope = {
		addEventListener: (type: string, listener: Listener) => listeners.set(type, listener),
		skipWaiting: vi.fn(() => Promise.resolve()),
		caches: {
			keys: vi.fn(() => Promise.resolve([...cacheNames])),
			delete: vi.fn((name: string) => Promise.resolve(cacheNames.includes(name)))
		},
		clients: { matchAll: vi.fn(), claim: vi.fn() },
		registration: {
			unregister: vi.fn(),
			showNotification: vi.fn(() => Promise.resolve())
		}
	};
	return { scope, listeners };
}

/** Runs one event through its listener and waits for what it passed to waitUntil(). */
async function dispatch(listener: Listener | undefined, extra: object = {}): Promise<void> {
	const pending: Promise<unknown>[] = [];
	listener?.({ waitUntil: (promise: Promise<unknown>) => pending.push(promise), ...extra });
	await Promise.all(pending);
}

async function loadEmergencyWorker() {
	const { scope, listeners } = fakeScope();
	vi.stubGlobal('self', scope);
	vi.resetModules();
	// Through a variable, so the root TypeScript program (DOM types) does not pull in the
	// worker, which emergency/tsconfig.json checks with WebWorker types.
	const emergencyWorker = '../../emergency/service-worker.js';
	await import(/* @vite-ignore */ emergencyWorker);
	return { scope, listeners };
}

test('SEC-169 (e): the emergency service worker takes over at once and empties Cache Storage without touching IndexedDB', async () => {
	const db = new LocalDb();
	await db.outbox.add({
		idempotencyKey: crypto.randomUUID(),
		kind: 'test_command',
		payload: new Uint8Array([1]),
		protective: false,
		createdAt: 0,
		releaseAt: 0
	});
	try {
		const { scope, listeners } = await loadEmergencyWorker();
		await dispatch(listeners.get('install'));
		expect(scope.skipWaiting).toHaveBeenCalledTimes(1);
		await dispatch(listeners.get('activate'));
		expect(scope.caches.delete.mock.calls.map(([name]) => name).sort()).toEqual([
			'casa-shell-v1',
			'other'
		]);
		expect(await db.outbox.count()).toBe(1);
	} finally {
		await db.delete();
	}
});

test('SEC-169 (e): the emergency service worker has no fetch handler, never navigates pages, and never unregisters', async () => {
	const { scope, listeners } = await loadEmergencyWorker();
	await dispatch(listeners.get('install'));
	await dispatch(listeners.get('activate'));
	expect(listeners.has('fetch')).toBe(false);
	expect(scope.clients.matchAll).not.toHaveBeenCalled();
	expect(scope.registration.unregister).not.toHaveBeenCalled();
});

test('SEC-176 (b): the emergency service worker always shows a notification for a push', async () => {
	const { scope, listeners } = await loadEmergencyWorker();
	await dispatch(listeners.get('push'), { data: { json: () => ({ text: 'Ada kabar baru' }) } });
	await dispatch(listeners.get('push'), { data: null });
	await dispatch(listeners.get('push'), {
		data: {
			json: () => {
				throw new SyntaxError('not JSON');
			}
		}
	});
	expect(scope.registration.showNotification.mock.calls).toEqual([
		['Casa', { body: 'Ada kabar baru' }],
		['Casa', {}],
		['Casa', {}]
	]);
});
