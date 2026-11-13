import 'fake-indexeddb/auto';
import { afterEach, beforeEach, expect, test, vi } from 'vitest';
import { LocalDb, deleteLocalDb, localDb, type OutboxEntry } from './local-db';
import { createOutbox, type Send, type SendResult } from './outbox';

let t = 0;
const clock = () => t;
// 5 seconds: the ordinary undo window (08 §9.3; 09.3 §1), given by the caller as the
// command kind's value from the contract.
const ordinary = {
	kind: 'test_command',
	payload: new Uint8Array([1]),
	holdSeconds: 5,
	protective: false
};
const protective = { ...ordinary, holdSeconds: 0, protective: true };
const uuidV4 = /^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/;

const accepted: SendResult = { kind: 'accepted' };
let accepting = vi.fn<Send>();
const extraDbs: LocalDb[] = [];

beforeEach(() => {
	t = 0;
	accepting = vi.fn<Send>(() => Promise.resolve(accepted));
});

afterEach(async () => {
	for (const db of extraDbs.splice(0)) db.close();
	await deleteLocalDb();
});

type Outbox = ReturnType<typeof createOutbox>;

/**
 * Saves and waits for the send pass that saving starts, so it runs at the save time and not
 * at a clock value a test sets afterwards.
 */
async function save(outbox: Outbox, command: typeof ordinary): Promise<OutboxEntry> {
	const entry = await outbox.enqueue(command);
	await outbox.flush();
	return entry;
}

function fakeDocument(visibilityState: DocumentVisibilityState) {
	return Object.assign(new EventTarget(), { visibilityState });
}

function secondTab(): LocalDb {
	const db = new LocalDb();
	extraDbs.push(db);
	return db;
}

test('FR-FIN-040/AC-17: a saved command is stored at once and not sent before its 5-second undo window ends', async () => {
	const outbox = createOutbox(localDb(), accepting, clock);
	await save(outbox, ordinary);
	expect(await outbox.unsentCount()).toBe(1);
	t = 4999;
	await outbox.flush();
	expect(accepting).not.toHaveBeenCalled();
	t = 5000;
	await outbox.flush();
	expect(accepting).toHaveBeenCalledTimes(1);
	expect(await outbox.unsentCount()).toBe(0);
});

test('FR-FIN-040/AC-17: undo inside the undo window removes the command and nothing is sent', async () => {
	const outbox = createOutbox(localDb(), accepting, clock);
	const entry = await save(outbox, ordinary);
	t = 4999;
	expect(await outbox.undo(entry.idempotencyKey)).toBe(true);
	t = 5000;
	await outbox.flush();
	expect(accepting).not.toHaveBeenCalled();
	expect(await outbox.unsentCount()).toBe(0);
});

test('FR-FIN-040/AC-17: undo after the undo window keeps the command', async () => {
	const outbox = createOutbox(
		localDb(),
		vi.fn<Send>(() => Promise.resolve({ kind: 'retry' })),
		clock
	);
	const entry = await save(outbox, ordinary);
	t = 5000;
	expect(await outbox.undo(entry.idempotencyKey)).toBe(false);
	expect(await outbox.unsentCount()).toBe(1);
});

test('FR-FIN-040/AC-17: undo of a rejected command keeps it', async () => {
	const rejecting = vi.fn<Send>(() => Promise.resolve({ kind: 'rejected', code: 'state_changed' }));
	const outbox = createOutbox(localDb(), rejecting, clock);
	const entry = await save(outbox, protective);
	expect(rejecting).toHaveBeenCalledTimes(1);
	expect(await outbox.undo(entry.idempotencyKey)).toBe(false);
	expect(await outbox.unsentCount()).toBe(1);
});

test('FR-FIN-040/AC-17: an undo window that is not a whole number of seconds is refused and nothing is queued', async () => {
	const outbox = createOutbox(localDb(), accepting, clock);
	for (const holdSeconds of [-1, 1.5, Number.NaN, Number.POSITIVE_INFINITY]) {
		await expect(outbox.enqueue({ ...ordinary, holdSeconds })).rejects.toThrow(RangeError);
	}
	expect(await outbox.unsentCount()).toBe(0);
});

test('FR-FIN-040/AC-17: a command saved before its tab closed is sent at the next open', async () => {
	const first = secondTab();
	await save(createOutbox(first, accepting, clock), ordinary);
	first.close();
	t = 5000;
	await createOutbox(secondTab(), accepting, clock).flush();
	expect(accepting).toHaveBeenCalledTimes(1);
});

test('FR-FIN-040/AC-17: a device clock moved back after saving does not hold the command', async () => {
	const outbox = createOutbox(localDb(), accepting, clock);
	t = 10_000;
	const entry = await save(outbox, ordinary);
	t = 1000;
	// Elapsed time is unknown once the clock went back, so the undo chance is gone too.
	expect(await outbox.undo(entry.idempotencyKey)).toBe(false);
	await outbox.flush();
	expect(accepting).toHaveBeenCalledTimes(1);
});

test('SEC-151: a protective command is not held and is sent before older ordinary commands', async () => {
	const sent: string[] = [];
	const outbox = createOutbox(
		localDb(),
		(entry) => {
			sent.push(entry.idempotencyKey);
			return Promise.resolve(accepted);
		},
		clock
	);
	const first = await save(outbox, ordinary);
	t = 1;
	const second = await save(outbox, ordinary);
	t = 6000;
	const urgent = await save(outbox, protective);
	await outbox.flush();
	expect(sent).toEqual([urgent.idempotencyKey, first.idempotencyKey, second.idempotencyKey]);
});

test('SEC-173 (a): two tabs flushing at once send every command exactly once', async () => {
	const request = vi.spyOn(navigator.locks, 'request');
	const sentKeys: string[] = [];
	const slowSend: Send = async (entry) => {
		await new Promise((resolve) => setTimeout(resolve, 5));
		sentKeys.push(entry.idempotencyKey);
		return accepted;
	};
	const a = createOutbox(secondTab(), slowSend, clock);
	const b = createOutbox(secondTab(), slowSend, clock);
	const saved = await Promise.all([a.enqueue(ordinary), a.enqueue(ordinary), b.enqueue(ordinary)]);
	await Promise.all([a.flush(), b.flush()]);
	t = 5000;
	await Promise.all([a.flush(), b.flush()]);
	expect(sentKeys.sort()).toEqual(saved.map((entry) => entry.idempotencyKey).sort());
	expect(request).toHaveBeenCalledWith('outbox', expect.any(Function));
});

test('FR-FIN-040/AC-18: a retry answer stops the pass and keeps every command queued', async () => {
	const retrying = vi.fn<Send>(() => Promise.resolve({ kind: 'retry' }));
	const outbox = createOutbox(localDb(), retrying, clock);
	await Promise.all([outbox.enqueue(ordinary), outbox.enqueue(ordinary), outbox.enqueue(ordinary)]);
	await outbox.flush();
	t = 5000;
	await outbox.flush();
	expect(retrying).toHaveBeenCalledTimes(1);
	expect(await outbox.unsentCount()).toBe(3);
});

test('FR-FIN-040/AC-18: a rejected command stays queued with its code and is not sent again', async () => {
	const rejecting = vi.fn<Send>(() => Promise.resolve({ kind: 'rejected', code: 'state_changed' }));
	const outbox = createOutbox(localDb(), rejecting, clock);
	const entry = await save(outbox, ordinary);
	t = 5000;
	await outbox.flush();
	const stored: OutboxEntry | undefined = await localDb().outbox.get(entry.idempotencyKey);
	expect(stored?.rejectedCode).toBe('state_changed');
	expect(await outbox.unsentCount()).toBe(1);
	await outbox.flush();
	expect(rejecting).toHaveBeenCalledTimes(1);
});

test('FR-FIN-040/AC-18: a send that throws leaves the command queued', async () => {
	const outbox = createOutbox(localDb(), () => Promise.reject(new Error('network down')), clock);
	await save(outbox, ordinary);
	t = 5000;
	await expect(outbox.flush()).rejects.toThrow('network down');
	expect(await outbox.unsentCount()).toBe(1);
});

test('SEC-171 (f): a failed write rejects the save with the original error and queues nothing', async () => {
	const db = localDb();
	const outbox = createOutbox(db, accepting, clock);
	vi.spyOn(db.outbox, 'add').mockRejectedValue(new DOMException('full', 'QuotaExceededError'));
	await expect(outbox.enqueue(protective)).rejects.toMatchObject({ name: 'QuotaExceededError' });
	expect(await outbox.unsentCount()).toBe(0);
	await outbox.flush();
	expect(accepting).not.toHaveBeenCalled();
});

test('ADR-005: the idempotency key is a UUIDv4 made at save time and kept on every attempt', async () => {
	const send = vi
		.fn<Send>()
		.mockResolvedValueOnce({ kind: 'retry' })
		.mockResolvedValueOnce(accepted);
	const outbox = createOutbox(localDb(), send, clock);
	const entry = await save(outbox, ordinary);
	expect(entry.idempotencyKey).toMatch(uuidV4);
	t = 5000;
	await outbox.flush();
	await outbox.flush();
	expect(send.mock.calls.map(([sent]) => sent.idempotencyKey)).toEqual([
		entry.idempotencyKey,
		entry.idempotencyKey
	]);
});

test('SEC-171 (d): the outbox sends on open, on online, and when the page becomes visible', async () => {
	const win = new EventTarget();
	const doc = fakeDocument('hidden');
	const retrying = vi.fn<Send>(() => Promise.resolve({ kind: 'retry' }));
	const outbox = createOutbox(localDb(), retrying, clock);
	await localDb().outbox.add({
		idempotencyKey: crypto.randomUUID(),
		kind: 'test_command',
		payload: new Uint8Array([1]),
		protective: true,
		createdAt: 0,
		releaseAt: 0
	});
	const sends = (count: number) =>
		vi.waitFor(() => {
			expect(retrying).toHaveBeenCalledTimes(count);
		});

	const stop = outbox.start(win, doc);
	await sends(1);
	win.dispatchEvent(new Event('online'));
	await sends(2);
	doc.dispatchEvent(new Event('visibilitychange'));
	doc.visibilityState = 'visible';
	doc.dispatchEvent(new Event('visibilitychange'));
	await sends(3);
	stop();
	win.dispatchEvent(new Event('online'));
	doc.dispatchEvent(new Event('visibilitychange'));
	await new Promise((resolve) => setTimeout(resolve, 20));
	expect(retrying).toHaveBeenCalledTimes(3);
});

test('SEC-171 (d): a new command is sent when its undo window ends without another trigger', async () => {
	// fake-indexeddb needs the real setImmediate, so only timers and Date are faked.
	vi.useFakeTimers({ toFake: ['setTimeout', 'clearTimeout', 'Date'] });
	try {
		const win = new EventTarget();
		const doc = fakeDocument('visible');
		const outbox = createOutbox(localDb(), accepting, Date.now);
		const stop = outbox.start(win, doc);
		await outbox.enqueue(ordinary);
		await outbox.flush();
		await vi.advanceTimersByTimeAsync(4999);
		expect(accepting).not.toHaveBeenCalled();
		await vi.advanceTimersByTimeAsync(1);
		await vi.waitFor(() => {
			expect(accepting).toHaveBeenCalledTimes(1);
		});
		stop();
	} finally {
		vi.useRealTimers();
	}
});
