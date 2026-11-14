<script lang="ts">
	// The verification link <origin>/signup/verify#<token> (identity/signup.go), with the words
	// of 08 §7.6 verify.*. The init hook (hooks.client.ts) took the token. Outside a Safari tab:
	// LY-91 first if no mode is chosen, then the password and CompleteSignUp with the mode
	// (SEC-165 (a)), then the next screen after sign-in. In a Safari tab on iPhone or iPad: LY-90,
	// the token dropped unsent, until the contract can complete it without a session (10.5
	// SEC-171 (b) (2), Q2-6). The client checks no length or list itself; the server's
	// violations are worded on the field (SEC-11; plan decision 7).
	import { tick } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import Button from '#lib/components/Button.svelte';
	import ButtonRow from '#lib/components/ButtonRow.svelte';
	import InlineMessage from '#lib/components/InlineMessage.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import TextField from '#lib/components/TextField.svelte';
	import { api, ctx, errorCode } from '#lib/api.js';
	import { copy } from '#lib/copy.js';
	import { displayContext } from '#lib/device.js';
	import { actionFailure, fieldErrors } from '#lib/form.js';
	import { ErrorCode } from '#lib/gen/casa/v1/error_pb.js';
	import { DeviceMode } from '#lib/gen/casa/v1/identity_pb.js';
	import { afterSignIn, handleSessionError, nextScreen, session } from '#lib/session.svelte.js';

	const passwordSentences = {
		password: {
			'string.min_len': copy.verify.tooShort,
			'string.max_len': copy.verify.tooLong,
			'password.common': copy.verify.common
		}
	};

	const token = $derived(
		session.pendingToken?.kind === 'verification' ? session.pendingToken.value : undefined
	);
	let rejected = $state(false);
	// Without a token (reloaded, or used by a sign-up that went through) the page has nothing to
	// complete, so the screen nextScreen() names follows: the member's screens once signed in.
	const redirect = $derived(
		displayContext() === 'ios-safari-tab'
			? '/pasang'
			: !session.mode
				? '/masuk/mode'
				: !token && !rejected
					? nextScreen()
					: undefined
	);

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});

	const personal = $derived(session.mode !== 'shared');

	let password = $state('');
	let passwordError: string | undefined = $state();
	let failure: string | undefined = $state();
	let busy = $state(false);
	let rejectedBox: HTMLDivElement | undefined = $state();

	async function complete(verificationToken: string): Promise<void> {
		const mode = personal ? 'personal' : 'shared';
		failure = undefined; // so the same sentence again is announced again
		passwordError = undefined;
		try {
			await api.identity.completeSignUp({
				requestContext: ctx(true),
				verificationToken,
				password,
				deviceMode: mode === 'personal' ? DeviceMode.PERSONAL : DeviceMode.SHARED
			});
		} catch (error) {
			if (await handleSessionError(error)) return;
			if (errorCode(error) === ErrorCode.TOKEN_REJECTED) {
				// One sentence for every reason (SEC-30); the token is of no further use.
				rejected = true;
				session.pendingToken = undefined;
				// The form is gone; the sentence's "Masuk" button takes focus (08 §25.5).
				await tick();
				rejectedBox?.querySelector('button')?.focus();
				return;
			}
			const fields = fieldErrors(error, passwordSentences);
			if (fields) passwordError = fields.password;
			else failure = actionFailure(error);
			return;
		}
		await afterSignIn(mode);
		session.pendingToken = undefined; // the redirect above then opens nextScreen()
	}
</script>

{#if !redirect}
	<Screen heading={copy.verify.title} destination="Buat kata sandi">
		{#if rejected}
			<div bind:this={rejectedBox}>
				<InlineMessage role="alert" text={copy.verify.rejected} />
				<ButtonRow>
					<Button
						variant="primary"
						onclick={() => {
							goto(resolve('/masuk/akun')).catch(console.error);
						}}>{copy.verify.signIn}</Button
					>
				</ButtonRow>
			</div>
		{:else if token}
			<form
				novalidate
				onsubmit={(event) => {
					event.preventDefault();
					if (busy) return;
					busy = true;
					complete(token)
						.catch(console.error)
						.finally(() => {
							busy = false;
						});
				}}
			>
				<TextField
					label={copy.verify.password}
					type="password"
					autocomplete={personal ? 'new-password' : 'off'}
					hint={personal ? copy.verify.hint : `${copy.verify.hint} ${copy.signIn.doNotSave}`}
					error={passwordError}
					bind:value={password}
				/>
				<InlineMessage role="alert" text={failure} />
				<ButtonRow>
					<Button variant="primary" type="submit" {busy}>{copy.verify.submit}</Button>
				</ButtonRow>
			</form>
		{/if}
	</Screen>
{/if}
