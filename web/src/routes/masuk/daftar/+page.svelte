<script lang="ts">
	// 08 §5.1 LY-02 "Buat akun" with the words of 08 §7.6 signUp.*. With a setup or invitation
	// link held: the email, then StartSignUp with that one token (SEC-158); the password is made
	// on the verification page (plan decision 7). Without a link, sign-up is closed: the same
	// sentence for anyone and a field for a pasted link, checked here without a request
	// (SEC-158; SEC-175 (c)).
	import { tick } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import Button from '#lib/components/Button.svelte';
	import ButtonRow from '#lib/components/ButtonRow.svelte';
	import InlineMessage from '#lib/components/InlineMessage.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import TextField from '#lib/components/TextField.svelte';
	import TextLink from '#lib/components/TextLink.svelte';
	import { copy } from '#lib/copy.js';
	import { parsePastedLink } from '#lib/link-token.js';
	import { entryRedirect, session } from '#lib/session.svelte.js';
	import { heldSignUpToken, startSignUp, type SignUpToken } from '#lib/sign-up.js';

	const redirect = $derived(entryRedirect('/masuk/daftar'));
	const id = $props.id();

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});

	const held = $derived(heldSignUpToken());
	// LY-02 shows only with a mode chosen (entryRedirect), so personal unless shared was chosen.
	const personal = $derived(session.mode !== 'shared');

	// "Ganti email" on LY-03 comes back here with the email kept (08 §7.6 checkEmail.change).
	let email = $state(session.signUpEmail ?? '');
	let pasted = $state('');
	let pasteError: string | undefined = $state();
	let emailError: string | undefined = $state();
	let failure: string | undefined = $state();
	let busy = $state(false);
	let emailBox: HTMLDivElement | undefined = $state();

	/**
	 * A pasted link becomes the held token, or the not-Casa sentence under the field, with no
	 * request either way (SEC-175 (c) (2)); the field is emptied once the link is taken (c) (1).
	 */
	async function takePasted(): Promise<void> {
		const link = parsePastedLink(pasted, location.origin);
		if (link === 'not-casa') {
			pasteError = copy.signUp.notCasa;
			return;
		}
		session.pendingToken = { kind: link.kind, value: link.token };
		pasted = '';
		pasteError = undefined;
		// The paste field is gone; the email field it made way for takes focus (08 §25.5).
		await tick();
		emailBox?.querySelector('input')?.focus();
	}

	async function send(token: SignUpToken): Promise<void> {
		failure = undefined; // so the same sentence again is announced again
		emailError = undefined;
		const failed = await startSignUp(email, token);
		if (failed) {
			emailError = failed.email;
			failure = failed.alert;
			return;
		}
		session.signUpEmail = email;
		await goto(resolve('/masuk/periksa-email'));
	}

	function submit(event: SubmitEvent): void {
		event.preventDefault();
		if (!held) {
			takePasted().catch(console.error);
			return;
		}
		if (busy) return;
		busy = true;
		send(held)
			.catch(console.error)
			.finally(() => {
				busy = false;
			});
	}
</script>

{#if !redirect}
	<Screen heading={copy.signUp.title} destination="Buat akun" describedby="{id}-lead">
		{#if held}
			<p id="{id}-lead">
				{held.kind === 'setup' ? copy.signUp.withSetup : copy.signUp.withInvitation}
			</p>
			<form novalidate onsubmit={submit}>
				<div bind:this={emailBox}>
					<TextField
						label={copy.signUp.email.label}
						type="email"
						autocomplete={personal ? 'username' : 'off'}
						hint={copy.signUp.email.hint}
						error={emailError}
						bind:value={email}
					/>
				</div>
				<p class="sebelum-tombol">{copy.signUp.passwordLater}</p>
				<InlineMessage role="alert" text={failure} />
				<ButtonRow>
					<Button variant="primary" type="submit" {busy}>{copy.signUp.submit}</Button>
				</ButtonRow>
			</form>
		{:else}
			<p id="{id}-lead">{copy.signUp.closed}</p>
			<form novalidate onsubmit={submit}>
				<TextField
					label={copy.signUp.paste.label}
					type="url"
					autocomplete="off"
					hint={copy.signUp.paste.hint}
					error={pasteError}
					bind:value={pasted}
				/>
				<ButtonRow>
					<Button variant="primary" type="submit">{copy.signUp.paste.submit}</Button>
				</ButtonRow>
			</form>
			<TextLink href={resolve('/masuk/akun')}>{copy.signUp.signIn}</TextLink>
		{/if}
	</Screen>
{/if}

<style>
	.sebelum-tombol {
		margin: var(--jarak-4) 0 0;
	}
</style>
