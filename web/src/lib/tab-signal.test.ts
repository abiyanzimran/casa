import { afterEach, expect, test, vi } from 'vitest';
import { onTabSignal, postTabSignal, type TabSignal } from './tab-signal';

const stops: (() => void)[] = [];
afterEach(() => {
	for (const stop of stops.splice(0)) stop();
});

function listen(): TabSignal[] {
	const received: TabSignal[] = [];
	stops.push(onTabSignal((signal) => received.push(signal)));
	return received;
}

test('SEC-173 (b): every signal kind reaches another channel in the same browser', async () => {
	const received = listen();
	const signals: TabSignal[] = [
		{ kind: 'sign_out', unsentDropped: 2 },
		{ kind: 'session_revoked', unsentDropped: 2 },
		{ kind: 'household_exit', unsentDropped: 2 },
		{ kind: 'app_lock' }
	];
	for (const signal of signals) postTabSignal(signal);
	await vi.waitFor(() => {
		expect(received).toEqual(signals);
	});
});

test('SEC-173 (b): a malformed or unknown message is ignored', async () => {
	const received = listen();
	const raw = new BroadcastChannel('casa');
	for (const message of [
		'x',
		null,
		{ kind: 'nope' },
		{ kind: 'session_revoked' },
		{ kind: 'sign_out', unsentDropped: -1 },
		{ kind: 'sign_out', unsentDropped: 1.5 },
		{ kind: 'sign_out', unsentDropped: '2' }
	]) {
		raw.postMessage(message);
	}
	const sentinel: TabSignal = { kind: 'sign_out', unsentDropped: 0 };
	raw.postMessage(sentinel);
	raw.close();
	await vi.waitFor(() => {
		expect(received).toEqual([sentinel]);
	});
});
