// The one in-memory session state of the web client (plan decisions 4 to 6): the result of the
// session check, the device mode, the memberships, and a link token held for this tab. Nothing
// here is stored; a new document starts from 'checking'.
import { api, ctx, errorCode, failureOf } from './api';
import { ErrorCode } from './gen/casa/v1/error_pb';
import type { Membership } from './gen/casa/v1/household_pb';
import { localDbExists, wipePending } from './local-db';
import { broadcastWipe, wipeLocalData } from './wipe';

export type DeviceMode = 'personal' | 'shared';

/** Why a session ended, for LY-94 (08 §25.6). */
export type SessionEndReason = 'revoked' | 'expired' | 'shared-ended';

export type SessionState =
	| { kind: 'checking' }
	| { kind: 'signed-out' }
	| { kind: 'unreachable' }
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

/** The session check waits this long before the app opens without it (SEC-172 (b)). */
const sessionCheckLimitMs = 3000;

/**
 * The device mode on reopening, until GetSession says it (plan decision 5): a local database
 * exists only in personal mode (SEC-165 (b)).
 */
async function reopenedMode(): Promise<DeviceMode> {
	return (await localDbExists()) ? 'personal' : 'shared';
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

/**
 * Shows LY-94 with memory reset first (plan decision 6). A revoked ending, whose wipe has
 * happened, outranks a later expired one, and a second revoked answer never lowers the unsent
 * count the first wipe dropped.
 */
export function showEnded(reason: SessionEndReason, unsentDropped: number): void {
	const current = session.state;
	if (current.kind === 'ended' && current.reason === 'revoked') {
		if (reason !== 'revoked') return;
		unsentDropped = Math.max(unsentDropped, current.unsentDropped);
	}
	resetSessionState();
	session.state = { kind: 'ended', reason, unsentDropped };
}

/**
 * Shows the unreachable state with its re-check (08 §7.6 `home.unreachable`): where a revoked
 * wipe failed again after its one reload, the app neither says the data is deleted nor reloads
 * in a loop (SEC-172 (a)).
 */
export function showUnreachable(): void {
	settle({ kind: 'unreachable' });
	recheckWhenBack();
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
			// Unauthenticated without a detail: no session, nothing to wipe (SEC-161 (c) (5)).
			if (session.state.kind === 'ended') break;
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
		// SEC-172 (a): a wipe a failure left pending (marker in IndexedDB, 10.5 §5) finishes before
		// the session check, so before any screen or queue send. A failure counts as unreachable.
		if (await wipePending()) await wipeLocalData();
		const { memberships } = await api.household.listMemberships(
			{ requestContext: ctx(userInitiated) },
			{ timeoutMs }
		);
		session.mode ??= await reopenedMode();
		settle({ kind: 'signed-in', memberships });
	} catch (error) {
		try {
			if (await handleSessionError(error)) return;
			if (failureOf(error) !== 'offline') console.error(error);
		} catch (handlingError) {
			// Acting on the session error failed (a revoked wipe could not count the queue): stay
			// unreachable and check again, never on the check surface for good.
			console.error(handlingError);
		}
		showUnreachable();
	}
}
