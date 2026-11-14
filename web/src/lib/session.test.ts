import 'fake-indexeddb/auto';
import { Code, ConnectError } from '@connectrpc/connect';
import { afterEach, beforeEach, expect, test, vi } from 'vitest';
import { ErrorCode, ErrorDetailSchema } from './gen/casa/v1/error_pb';
import type { TabSignal } from './tab-signal';

// Fresh modules per test: the session state is module state, and after a wipe localDb()
// refuses until the page reloads, which new modules stand for here.
let sessionModule: typeof import('./session.svelte');
let localDbModule: typeof import('./local-db');

beforeEach(async () => {
	vi.resetModules();
	sessionModule = await import('./session.svelte');
	localDbModule = await import('./local-db');
});

afterEach(async () => {
	await localDbModule.deleteLocalDb();
});

function failed(code: Code, detail?: ErrorCode): ConnectError {
	const details =
		detail === undefined ? [] : [{ desc: ErrorDetailSchema, value: { code: detail } }];
	return new ConnectError('failed', code, undefined, details);
}

/** A personal-mode browser: the local database exists, with one unsent command. */
async function personalBrowser(): Promise<void> {
	await localDbModule.localDb().outbox.add({
		idempotencyKey: crypto.randomUUID(),
		kind: 'test_command',
		payload: new Uint8Array([1]),
		protective: false,
		createdAt: 0,
		releaseAt: 0
	});
}

/** The next tab signal any tab of this profile sends. */
function nextSignal(): Promise<TabSignal> {
	const channel = new BroadcastChannel('casa');
	return new Promise((resolve) => {
		channel.onmessage = (event: MessageEvent<TabSignal>) => {
			channel.close();
			resolve(event.data);
		};
	});
}

test('SEC-172 (a): a session_revoked answer tells every tab to wipe, with the unsent count', async () => {
	await personalBrowser();
	const signal = nextSignal();
	const { handleSessionError } = sessionModule;
	expect(await handleSessionError(failed(Code.Unauthenticated, ErrorCode.SESSION_REVOKED))).toBe(
		true
	);
	expect(await signal).toEqual({ kind: 'session_revoked', unsentDropped: 1 });
});

test('SEC-161 (d): a session_expired answer in personal mode ends the session and keeps the local data', async () => {
	await personalBrowser();
	const { handleSessionError, session } = sessionModule;
	session.state = { kind: 'signed-in', memberships: [] };
	session.mode = 'personal';
	session.ageDeclared = true;
	session.pendingToken = { kind: 'invitation', value: 'x' };

	expect(await handleSessionError(failed(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED))).toBe(
		true
	);
	expect(session.state).toEqual({ kind: 'ended', reason: 'expired', unsentDropped: 0 });
	// Memory is reset before LY-94 renders (plan decision 6).
	expect([session.mode, session.pendingToken, session.ageDeclared]).toEqual([
		undefined,
		undefined,
		false
	]);
	expect(await localDbModule.localDb().outbox.count()).toBe(1);
});

test('SEC-175 (c): no session on open keeps the link token this tab was opened with', async () => {
	const { handleSessionError, session } = sessionModule;
	session.pendingToken = { kind: 'setup', value: 'x' };
	session.setupLinkDropped = true;
	session.mode = 'personal';

	expect(await handleSessionError(failed(Code.Unauthenticated))).toBe(true);
	expect(session.state).toEqual({ kind: 'signed-out' });
	// The rest of memory is emptied as before; the link is what the signed-out screens are for.
	expect(session.mode).toBeUndefined();
	expect([session.pendingToken, session.setupLinkDropped]).toEqual([
		{ kind: 'setup', value: 'x' },
		true
	]);
});

test('SEC-161 (d): on open, a session_expired answer with a local database counts as personal mode', async () => {
	await personalBrowser();
	const { handleSessionError, session } = sessionModule;
	await handleSessionError(failed(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED));
	expect(session.state).toEqual({ kind: 'ended', reason: 'expired', unsentDropped: 0 });
});

test('SEC-165: a session_expired answer in shared mode is the shared ending', async () => {
	const { handleSessionError, session } = sessionModule;
	session.mode = 'shared';
	await handleSessionError(failed(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED));
	expect(session.state).toEqual({ kind: 'ended', reason: 'shared-ended', unsentDropped: 0 });
});

test('SEC-165: on open, a session_expired answer without a local database is the shared ending', async () => {
	const { handleSessionError, session } = sessionModule;
	await handleSessionError(failed(Code.Unauthenticated, ErrorCode.SESSION_EXPIRED));
	expect(session.state).toEqual({ kind: 'ended', reason: 'shared-ended', unsentDropped: 0 });
	expect(await localDbModule.localDbExists()).toBe(false); // SEC-165 (b): nothing created
});

test('SEC-161 (c) (5): Unauthenticated without a detail means signed out, without a wipe', async () => {
	await personalBrowser();
	const { handleSessionError, session } = sessionModule;
	session.state = { kind: 'signed-in', memberships: [] };
	expect(await handleSessionError(failed(Code.Unauthenticated))).toBe(true);
	expect(session.state).toEqual({ kind: 'signed-out' });
	expect(await localDbModule.localDb().outbox.count()).toBe(1);
});

test('ADR-030 item 10: any other failure is not a session error and leaves the state alone', async () => {
	const { handleSessionError, session } = sessionModule;
	session.state = { kind: 'signed-in', memberships: [] };
	for (const error of [
		failed(Code.Unauthenticated, ErrorCode.CREDENTIALS_REJECTED),
		failed(Code.Unavailable),
		failed(Code.DeadlineExceeded),
		failed(Code.FailedPrecondition, ErrorCode.NOTICE_UNREAD),
		new TypeError('Failed to fetch')
	]) {
		expect(await handleSessionError(error)).toBe(false);
	}
	expect(session.state).toEqual({ kind: 'signed-in', memberships: [] });
});

/** The session check's answer, and what an unreachable state needs from the page. */
async function answerSessionCheck(
	...answers: (Error | 'signed-in')[]
): Promise<{ reload: () => void }> {
	const { api } = await import('./api');
	const spy = vi.spyOn(api.household, 'listMemberships');
	for (const answer of answers) {
		if (answer === 'signed-in') spy.mockResolvedValueOnce({ memberships: [] } as never);
		else spy.mockRejectedValueOnce(answer);
	}
	vi.stubGlobal('window', { addEventListener: vi.fn() });
	vi.stubGlobal('document', { addEventListener: vi.fn() });
	const location = { reload: vi.fn() };
	vi.stubGlobal('location', location);
	vi.stubGlobal('caches', { keys: () => Promise.resolve([]) });
	return location;
}

/** A wipe started for kind that did not finish: its marker in the local database. */
async function wipeLeftPending(kind: 'session_revoked' | 'sign_out'): Promise<void> {
	await personalBrowser();
	await localDbModule.localDb().meta.put({ key: 'wipe_pending', value: kind });
}

test('SEC-172 (a): a revoked wipe left pending, answered without a detail (after 1.7d), still ends as revoked with its unsent count', async () => {
	await wipeLeftPending('session_revoked');
	await answerSessionCheck(failed(Code.Unauthenticated));
	const signal = nextSignal();
	const { checkSession, carryUnsentDropped, session } = sessionModule;
	await checkSession();
	expect(await signal).toEqual({ kind: 'session_revoked', unsentDropped: 0 });
	expect(carryUnsentDropped(0)).toBe(1); // counted before the wipe, for LY-94
	expect(session.state.kind).not.toBe('signed-out');
	expect(await localDbModule.localDbExists()).toBe(false);
});

test('SEC-161 (c) (5): a sign-out wipe left pending, answered without a detail, is signed out', async () => {
	await wipeLeftPending('sign_out');
	await answerSessionCheck(failed(Code.Unauthenticated));
	const { checkSession, session } = sessionModule;
	await checkSession();
	expect(session.state).toEqual({ kind: 'signed-out' });
	expect(await localDbModule.localDbExists()).toBe(false);
});

test('SEC-172 (a): a page reloaded after a failed revoked wipe treats an answer without a detail as revoked', async () => {
	await answerSessionCheck(failed(Code.Unauthenticated));
	const signal = nextSignal();
	const { checkSession, carryRevokedWipe } = sessionModule;
	carryRevokedWipe(2); // from page.state after the one reload
	await checkSession();
	expect(await signal).toEqual({ kind: 'session_revoked', unsentDropped: 0 });
	expect(sessionModule.carryUnsentDropped(0)).toBe(2);
});

test('SEC-172 (a): a pending wipe still runs when its queue cannot be counted', async () => {
	await wipeLeftPending('session_revoked');
	await answerSessionCheck(failed(Code.Unauthenticated, ErrorCode.SESSION_REVOKED));
	// Dexie's shared Table prototype, so every connection's count fails (iOS connection loss).
	const table = Object.getPrototypeOf(
		Object.getPrototypeOf(new localDbModule.LocalDb().outbox)
	) as { count: () => Promise<number> };
	vi.spyOn(table, 'count').mockRejectedValue(new Error('connection lost'));
	vi.spyOn(console, 'error').mockImplementation(() => {});
	const signal = nextSignal();
	await sessionModule.checkSession();
	expect(await signal).toEqual({ kind: 'session_revoked', unsentDropped: 0 });
	expect(await localDbModule.localDbExists()).toBe(false);
});

test('SEC-172 (a): after a pending wipe and an unreachable check, a valid session starts on a fresh page', async () => {
	await wipeLeftPending('session_revoked');
	const { reload } = await answerSessionCheck(failed(Code.Unavailable), 'signed-in');
	vi.spyOn(console, 'error').mockImplementation(() => {});
	const { checkSession, session } = sessionModule;
	await checkSession();
	expect(session.state.kind).toBe('unreachable');
	await checkSession();
	expect(reload).toHaveBeenCalledTimes(1);
});
