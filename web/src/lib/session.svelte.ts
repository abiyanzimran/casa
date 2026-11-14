// The one in-memory session state of the web client (plan decisions 4 to 6): the result of the
// session check, the device mode, the memberships, and a link token held for this tab. Nothing
// here is stored; a new document starts from 'checking'.
import { Code, ConnectError } from '@connectrpc/connect';
import { noteUpgradeRequired } from './app-update.svelte';
import { api, ctx, errorCode, failureOf } from './api';
import { displayContext } from './device';
import { ErrorCode } from './gen/casa/v1/error_pb';
import type { Membership } from './gen/casa/v1/household_pb';
import {
	localDb,
	localDbExists,
	localDbWiped,
	requestPersistence,
	wipePending,
	type LocalDb
} from './local-db';
import { registerServiceWorker } from './service-worker-registration';
import { broadcastWipe, unsentCount, wipeLocalData } from './wipe';

export type DeviceMode = 'personal' | 'shared';

/** Why a session ended, for LY-94 (08 §25.6). */
export type SessionEndReason = 'revoked' | 'expired' | 'shared-ended';

/**
 * Why the app opened without the session (08 §7.6): no connection (home.unreachable), the
 * server erred or did not answer while online (home.unreachableServer), or a wipe failed again
 * (home.unreachableStorage).
 */
export type UnreachableCause = 'offline' | 'server' | 'storage';

export type SessionState =
	| { kind: 'checking' }
	| { kind: 'signed-out' }
	| { kind: 'unreachable'; cause: UnreachableCause }
	| { kind: 'signed-in'; memberships: Membership[] }
	| { kind: 'ended'; reason: SessionEndReason; unsentDropped: number };

interface Session {
	state: SessionState;
	mode?: DeviceMode;
	/** A link token taken from the address bar, held in this tab's memory only (SEC-175 (c)). */
	pendingToken?: { kind: 'setup' | 'invitation' | 'verification'; value: string };
	/** LY-04 was ticked in this session (plan decision 9). */
	ageDeclared: boolean;
}

export const session: Session = $state({ state: { kind: 'checking' }, ageDeclared: false });

/** The entry screens before a session (08 §25.6). */
export type EntryScreen = '/pasang' | '/masuk/mode' | '/masuk';

/** The screens after the entry screens, before a session: LY-02 Masuk (08 §7.6). */
export type SignInScreen = '/masuk/akun';

/**
 * The screen the app opens at (08 §25.6 flow): LY-90 in a Safari tab on iPhone or iPad until the
 * shared way is chosen (SEC-171 (b)); LY-91 until a device mode is chosen (SEC-165); then LY-01.
 * A member's screen is "/".
 */
export function nextScreen(): EntryScreen | '/' {
	if (session.state.kind === 'signed-in') return '/'; // ponytail: Task 8 adds LY-04, LY-05, LY-06
	if (displayContext() === 'ios-safari-tab' && session.mode !== 'shared') return '/pasang';
	return session.mode ? '/masuk' : '/masuk/mode';
}

/**
 * Where an entry screen sends the person instead of showing, so no step is skipped by its
 * address: LY-90 only in a Safari tab, LY-91 never there (SEC-171 (b) (1)), LY-01 and LY-02
 * only with a mode chosen (SEC-165). An earlier step stays reachable, so Back can change the
 * choice. Once signed in, every one of them sends the person on to "/".
 */
export function entryRedirect(path: EntryScreen | SignInScreen): EntryScreen | '/' | undefined {
	const tab = displayContext() === 'ios-safari-tab';
	const shows =
		session.state.kind === 'signed-out' &&
		(path === '/pasang' ? tab : path === '/masuk/mode' ? !tab : nextScreen() === '/masuk');
	return shows ? undefined : nextScreen();
}

/** The session check waits this long before the app opens without it (SEC-172 (b)). */
const sessionCheckLimitMs = 3000;

/**
 * The device mode on reopening, until GetSession says it (plan decision 5): a local database
 * exists only in personal mode (SEC-165 (b)).
 */
async function reopenedMode(): Promise<DeviceMode> {
	return (await localDbExists()) ? 'personal' : 'shared';
}

/**
 * What personal mode keeps on the device (SEC-171 (a); ADR-030 item 9): the service worker,
 * and the browser's answer to persistent storage in the local database. Shared mode never
 * calls this (SEC-165 (b)). Neither step holds up the screen; a failure is logged (ES §1).
 */
export function enablePersonalMode(db: LocalDb): void {
	registerServiceWorker().catch(logFailure);
	requestPersistence(db).catch(logFailure);
}

function logFailure(error: unknown): void {
	console.error(error);
}

/**
 * After SignIn opened a session in mode (SEC-165 (a)): personal mode keeps what it keeps on
 * the device, shared mode nothing (SEC-165 (b)); then the session is checked as the tap's
 * result, which brings the member's screen (SEC-165 (g)).
 */
export async function afterSignIn(mode: DeviceMode): Promise<void> {
	session.mode = mode;
	if (mode === 'personal') {
		// Created before the session check looks for a pending wipe: that look opens without a
		// version and deletes what it finds missing, which must not race the creation.
		const db = localDb();
		await db.open().catch(logFailure);
		enablePersonalMode(db);
	}
	await checkSession({ userInitiated: true });
}

/** Empties the memory of the session before LY-94 or the sign-in screen renders (decision 6). */
export function resetSessionState(): void {
	session.state = { kind: 'checking' };
	session.mode = undefined;
	session.pendingToken = undefined;
	session.ageDeclared = false;
}

/**
 * Sets a state other than LY-94. An ended session is final for this document: only a fresh
 * document leaves LY-94, so a late check, a restore, or a re-check never hides it or its unsent
 * line (08 §25.6; ADR-024 item 6).
 */
function settle(next: Exclude<SessionState, { kind: 'ended' }>): void {
	if (session.state.kind !== 'ended') session.state = next;
}

/** Unsent commands a wipe of this tab dropped before LY-94 could say so (ADR-024 item 6). */
let unsentDroppedCarried = 0;

/**
 * Keeps the number of unsent commands a wipe drops until LY-94 shows it: a wipe left pending,
 * or one that failed and reloads, would otherwise leave nothing to count (ADR-024 item 6;
 * SEC-171 (c)). Returns the largest count carried so far.
 */
export function carryUnsentDropped(count: number): number {
	unsentDroppedCarried = Math.max(unsentDroppedCarried, count);
	return unsentDroppedCarried;
}

/**
 * A revoked wipe this page carries: one left pending (its marker says session_revoked), or one
 * that failed and reloaded (page.state.wipeRetried). The server expires the cookie with its
 * session_revoked answer (1.7d), so the next answer is Unauthenticated without a detail; on
 * this page that still means revoked, so LY-94 keeps its unsent count (SEC-172 (a);
 * ADR-024 item 6).
 */
let revokedWipeCarried = false;

/** Called with page.state after the one reload of a failed revoked wipe (+layout.svelte). */
export function carryRevokedWipe(unsentDropped: number): void {
	revokedWipeCarried = true;
	carryUnsentDropped(unsentDropped);
}

/**
 * Shows LY-94 with memory reset first (plan decision 6). A revoked ending, whose wipe has
 * happened, outranks a later expired one, and its unsent count never drops below what an
 * earlier wipe of this tab dropped.
 */
export function showEnded(reason: SessionEndReason, unsentDropped: number): void {
	const current = session.state;
	if (current.kind === 'ended' && current.reason === 'revoked') {
		if (reason !== 'revoked') return;
		unsentDropped = Math.max(unsentDropped, current.unsentDropped);
	}
	if (reason === 'revoked') unsentDropped = carryUnsentDropped(unsentDropped);
	resetSessionState();
	session.state = { kind: 'ended', reason, unsentDropped };
}

/**
 * Shows the unreachable state with its re-check (08 §7.6): where a revoked wipe failed again
 * after its one reload ('storage'), the app neither says the data is deleted nor reloads in a
 * loop (SEC-172 (a)).
 */
export function showUnreachable(cause: UnreachableCause): void {
	settle({ kind: 'unreachable', cause });
	recheckWhenBack();
}

/**
 * A failed session check as the person sees it (08 §7.6): no connection, or the server's. A
 * timeout while the browser is online is the server's, though failureOf counts it offline.
 */
function unreachableCause(error: unknown): UnreachableCause {
	const timedOutOnline =
		navigator.onLine && ConnectError.from(error).code === Code.DeadlineExceeded;
	return failureOf(error) === 'offline' && !timedOutOnline ? 'offline' : 'server';
}

/**
 * Ends the session. A revoked one is wiped in every tab, this one included, and each tab shows
 * LY-94 with the unsent count once its wipe is done (SEC-172 (a)); an expired one keeps the
 * local data (SEC-161 (d), SEC-172 (d)).
 */
export async function endSession(reason: SessionEndReason): Promise<void> {
	if (reason === 'revoked') {
		await broadcastWipe('session_revoked');
		return;
	}
	showEnded(reason, 0);
}

/**
 * Acts on a failed call if it is a session error and says whether it was. Every RPC caller
 * calls this first (ADR-030 item 10).
 */
export async function handleSessionError(error: unknown): Promise<boolean> {
	noteUpgradeRequired(error); // the "Perbarui" row, from any answer (SEC-169 (d))
	if (failureOf(error) !== 'session') return false;
	switch (errorCode(error)) {
		case ErrorCode.SESSION_REVOKED:
			await endSession('revoked');
			break;
		case ErrorCode.SESSION_EXPIRED: {
			const mode = session.mode ?? (await reopenedMode());
			await endSession(mode === 'shared' ? 'shared-ended' : 'expired');
			break;
		}
		default:
			if (session.state.kind === 'ended') break;
			if (revokedWipeCarried) {
				await endSession('revoked');
				break;
			}
			// Unauthenticated without a detail: no session, nothing to wipe (SEC-161 (c) (5)).
			resetSessionState();
			settle({ kind: 'signed-out' });
	}
	return true;
}

let recheckArmed = false;

/** One listener checks again once the browser is back online or the page is shown again. */
function recheckWhenBack(): void {
	if (recheckArmed) return;
	recheckArmed = true;
	const listening = new AbortController();
	const recheck = () => {
		if (document.visibilityState !== 'visible') return;
		listening.abort();
		recheckArmed = false;
		// Only the unreachable state is waiting for this; LY-94 or a signed-in screen is not.
		if (session.state.kind !== 'unreachable') return;
		checkSession().catch((error: unknown) => {
			console.error(error);
		});
	};
	window.addEventListener('online', recheck, { signal: listening.signal });
	document.addEventListener('visibilitychange', recheck, { signal: listening.signal });
}

/**
 * Finishes a wipe a failure left pending (the marker in IndexedDB, 10.5 §5) before the session
 * check, so before any screen or queue send, counting its unsent commands first (SEC-172 (a);
 * ADR-024 item 6). A failure is thrown and shows the 'storage' unreachable state.
 */
async function finishPendingWipe(): Promise<void> {
	const kind = await wipePending();
	if (!kind) return;
	carryUnsentDropped(await unsentCount());
	if (kind === 'session_revoked') revokedWipeCarried = true;
	await wipeLocalData(kind);
}

/**
 * The session check on open (SEC-172 (b); plan decision 4): ListMemberships, waited for at
 * most 3 seconds. A session error ends the session; anything else opens the app as
 * unreachable and checks again once back. user_initiated is false unless a tap caused it
 * (SEC-165 (g)).
 */
export async function checkSession({
	userInitiated = false,
	timeoutMs = sessionCheckLimitMs
}: { userInitiated?: boolean; timeoutMs?: number } = {}): Promise<void> {
	try {
		await finishPendingWipe();
	} catch (error) {
		console.error(error);
		showUnreachable('storage');
		return;
	}
	try {
		const { memberships } = await api.household.listMemberships(
			{ requestContext: ctx(userInitiated) },
			{ timeoutMs }
		);
		if (localDbWiped()) {
			// A wipe on this page, on open or before an unreachable check, leaves its local
			// database closed for good (deleteLocalDb), so a valid session starts on a fresh page,
			// which has no marker any more.
			location.reload();
			return;
		}
		if (!session.mode) {
			// Reopened with a session: a personal-mode browser re-registers its service worker
			// (an update check) and asks for persistence again for LY-95 (SEC-171 (a)).
			session.mode = await reopenedMode();
			if (session.mode === 'personal') enablePersonalMode(localDb());
		}
		settle({ kind: 'signed-in', memberships });
	} catch (error) {
		try {
			if (await handleSessionError(error)) return;
			if (failureOf(error) !== 'offline') console.error(error);
		} catch (handlingError) {
			// Acting on the session error failed (local storage could not be read): stay
			// unreachable and check again, never on the check surface for good.
			console.error(handlingError);
		}
		showUnreachable(unreachableCause(error));
	}
}
