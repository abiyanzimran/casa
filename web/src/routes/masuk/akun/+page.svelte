<script lang="ts">
	// 08 §5.1 LY-02 "Masuk" with the words of 08 §7.6 signIn.*: email and password, then SignIn
	// with the device mode chosen on LY-91 or LY-90 (SEC-165 (a)). One sentence for a wrong email
	// or password (SEC-12); a failure that is no one field's sits above the button as an alert,
	// and only CREDENTIALS_REJECTED empties the password (08 §7.6). Shared mode saves no password
	// and says so (SEC-165 (e), SEC-175 (d)).
	import { tick } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import Button from '#lib/components/Button.svelte';
	import ButtonRow from '#lib/components/ButtonRow.svelte';
	import InlineMessage from '#lib/components/InlineMessage.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import TextField from '#lib/components/TextField.svelte';
	import TextLink from '#lib/components/TextLink.svelte';
	import { api, ctx, errorCode } from '#lib/api.js';
	import { copy } from '#lib/copy.js';
	import { actionFailure, fieldErrors, focusFirstInvalid } from '#lib/form.js';
	import { ErrorCode } from '#lib/gen/casa/v1/error_pb.js';
	import { DeviceMode } from '#lib/gen/casa/v1/identity_pb.js';
	import { afterSignIn, entryRedirect, handleSessionError, session } from '#lib/session.svelte.js';

	const redirect = $derived(entryRedirect('/masuk/akun'));

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});

	// LY-02 shows only with a mode chosen (entryRedirect), so personal unless shared was chosen.
	const personal = $derived(session.mode !== 'shared');

	let email = $state('');
	let password = $state('');
	let failure: string | undefined = $state();
	let emailError: string | undefined = $state();
	let passwordError: string | undefined = $state();
	let busy = $state(false);
	let form: HTMLFormElement | undefined = $state();
	let passwordBox: HTMLDivElement | undefined = $state();

	// buf.validate on SignInRequest names only lengths (SEC-12: no email rule; 08 §7.6).
	const fieldSentences = {
		email: { 'string.max_len': copy.signIn.emailInvalid },
		password: { 'string.max_len': copy.signIn.passwordTooLong }
	};

	async function signIn(): Promise<void> {
		const mode = personal ? 'personal' : 'shared';
		failure = undefined; // so the same sentence again is announced again
		emailError = undefined;
		passwordError = undefined;
		try {
			await api.identity.signIn({
				requestContext: ctx(true),
				email,
				password,
				deviceMode: mode === 'personal' ? DeviceMode.PERSONAL : DeviceMode.SHARED
			});
		} catch (error) {
			if (await handleSessionError(error)) return;
			// The sentence for a failed SignIn (08 §7.6, §13.2); only the unexpected is logged.
			if (errorCode(error) === ErrorCode.CREDENTIALS_REJECTED) {
				failure = copy.signIn.rejected;
				password = '';
				// The emptied password is what to type again (QG-09 PR B minor 9).
				await tick();
				passwordBox?.querySelector('input')?.focus();
				return;
			}
			const fields = fieldErrors(error, fieldSentences);
			emailError = fields?.email;
			passwordError = fields?.password;
			if (fields) await focusFirstInvalid(form);
			else failure = actionFailure(error);
			return;
		}
		await afterSignIn(mode);
	}
</script>

{#if !redirect}
	<Screen heading={copy.signIn.title} destination="Masuk" back={resolve('/masuk')}>
		<form
			bind:this={form}
			novalidate
			onsubmit={(event) => {
				event.preventDefault();
				if (busy) return;
				busy = true;
				signIn()
					.catch(console.error)
					.finally(() => {
						busy = false;
					});
			}}
		>
			<TextField
				label={copy.signIn.email}
				type="email"
				autocomplete={personal ? 'username' : 'off'}
				error={emailError}
				bind:value={email}
			/>
			<div bind:this={passwordBox}>
				<TextField
					label={copy.signIn.password}
					type="password"
					autocomplete={personal ? 'current-password' : 'off'}
					hint={personal ? undefined : copy.signIn.doNotSave}
					error={passwordError}
					bind:value={password}
				/>
			</div>
			<InlineMessage role="alert" text={failure} />
			<ButtonRow>
				<Button variant="primary" type="submit" {busy}>{copy.signIn.title}</Button>
			</ButtonRow>
		</form>
		<!-- 08 §7.6 signIn.links; "Saya diundang" comes with /masuk/undangan (Task 11) -->
		<TextLink href={resolve('/masuk/daftar')}>{copy.signIn.signUp}</TextLink>
	</Screen>
{/if}
