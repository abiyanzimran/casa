// StartSignUp from LY-02 Buat akun and LY-03 "Kirim ulang" (08 §7.6; plan decision 7): the
// email and the one link token held in memory (SEC-158).
import { api, ctx, errorCode } from './api';
import { copy } from './copy';
import { actionFailure, fieldErrors } from './form';
import { ErrorCode } from './gen/casa/v1/error_pb';
import { handleSessionError, session } from './session.svelte';

/** A link token sign-up can start with. */
export interface SignUpToken {
	kind: 'setup' | 'invitation';
	value: string;
}

/** The setup or invitation token this tab holds, if any (SEC-175 (c): memory only). */
export function heldSignUpToken(): SignUpToken | undefined {
	const token = session.pendingToken;
	if (!token || token.kind === 'verification') return undefined;
	return { kind: token.kind, value: token.value };
}

/**
 * What a failed start shows: a sentence on the email field, or one above the button.
 * tokenRejected: the server refused the held link (TOKEN_REJECTED).
 */
export interface SignUpFailure {
	email?: string;
	alert?: string;
	tokenRejected?: true;
}

const emailSentences = {
	email: { 'string.email': copy.signUp.emailInvalid, 'string.max_len': copy.signUp.emailInvalid }
};

/**
 * Starts sign-up with token in its own field, never both (SEC-158). Returns undefined once the
 * server took it, otherwise what to show. A rejected token gets one sentence for every reason,
 * by the kind of link held (SEC-30; 08 §7.6, plan Q1).
 */
export async function startSignUp(
	email: string,
	token: SignUpToken
): Promise<SignUpFailure | undefined> {
	try {
		await api.identity.startSignUp({
			requestContext: ctx(true),
			email,
			setupToken: token.kind === 'setup' ? token.value : '',
			invitationToken: token.kind === 'invitation' ? token.value : ''
		});
		return undefined;
	} catch (error) {
		if (await handleSessionError(error)) return {};
		if (errorCode(error) === ErrorCode.TOKEN_REJECTED) {
			return {
				alert: token.kind === 'setup' ? copy.signUp.setupRejected : copy.invitation.rejected,
				tokenRejected: true
			};
		}
		return fieldErrors(error, emailSentences) ?? { alert: actionFailure(error) };
	}
}
