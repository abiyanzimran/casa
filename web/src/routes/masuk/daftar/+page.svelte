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
	import { focusFirstInvalid } from '#lib/form.js';
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
	// A setup link the server refused: the email form stays, and another link can be pasted.
	let anotherLink = $state(false);
	let emailBox: HTMLDivElement | undefined = $state();
	let pasteBox: HTMLDivElement | undefined = $state();
	let form: HTMLFormElement | undefined = $state();

	/**
	 * A pasted link becomes the held token, or the not-Casa sentence under the field, with no
	 * request either way (SEC-175 (c) (2)); the field is emptied once the link is taken (c) (1).
	 * A sentence about the link held before goes with it.
	 */
	async function takePasted(): Promise<void> {
		failure = undefined;
		const link = parsePastedLink(pasted, location.origin);
		if (link === 'not-casa') {
			pasteError = copy.signUp.notCasa;
			// Its field takes focus, so the sentence is read (08 §25.5 "Formulir").
			await tick();
			pasteBox?.querySelector('input')?.focus();
			return;
		}
		session.pendingToken = { kind: link.kind, value: link.token };
		pasted = '';
		pasteError = undefined;
		anotherLink = false;
		// The paste field is gone; the email field takes focus (08 §25.5).
		await tick();
		emailBox?.querySelector('input')?.focus();
	}

	/**
	 * A refused link (SEC-30, one sentence per kind of link): an invitation is of no further
	 * use, so it is dropped and sign-up is closed again with its paste field; a setup link is
	 * kept, since the email may be what was wrong, and a field for another link is added. The
	 * sentence stays above the button either way (QG-09 PR B S5).
	 */
	async function refused(token: SignUpToken): Promise<void> {
		if (token.kind === 'setup') {
			anotherLink = true;
			return;
		}
		session.pendingToken = undefined;
		await tick();
		pasteBox?.querySelector('input')?.focus();
	}

	async function send(token: SignUpToken): Promise<void> {
		failure = undefined; // so the same sentence again is announced again
		emailError = undefined;
		const failed = await startSignUp(email, token);
		if (failed) {
			emailError = failed.email;
			failure = failed.alert;
			if (failed.tokenRejected) await refused(token);
			else if (emailError) await focusFirstInvalid(form);
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

{#snippet pasteField()}
	<div bind:this={pasteBox}>
		<TextField
			label={copy.signUp.paste.label}
			type="url"
			autocomplete="off"
			hint={copy.signUp.paste.hint}
			error={pasteError}
			bind:value={pasted}
		/>
	</div>
{/snippet}

{#if !redirect}
	<Screen
		heading={copy.signUp.title}
		destination="Buat akun"
		describedby="{id}-lead"
		back={resolve('/masuk')}
	>
		<!-- A status region from the start, so the sentence that confirms a taken link is
		     announced (QG-09 PR B minor 8). -->
		<p id="{id}-lead" role="status">
			{held
				? held.kind === 'setup'
					? copy.signUp.withSetup
					: copy.signUp.withInvitation
				: copy.signUp.closed}
		</p>
		<!-- One form whose fields follow the held link, so its alert stays mounted and is
		     announced when a refused invitation closes sign-up again (InlineMessage; 08 §25.5). -->
		<form bind:this={form} novalidate onsubmit={submit}>
			{#if held}
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
			{:else}
				{@render pasteField()}
			{/if}
			<InlineMessage role="alert" text={failure} />
			<ButtonRow>
				<Button variant="primary" type="submit" {busy}
					>{held ? copy.signUp.submit : copy.signUp.paste.submit}</Button
				>
			</ButtonRow>
		</form>
		{#if held && anotherLink}
			<form
				novalidate
				onsubmit={(event) => {
					event.preventDefault();
					takePasted().catch(console.error);
				}}
			>
				{@render pasteField()}
				<ButtonRow>
					<Button variant="outlined" type="submit">{copy.signUp.paste.submit}</Button>
				</ButtonRow>
			</form>
		{/if}
		<!-- With a link held too: an existing account signs in (QG-09 PR B minor 7). -->
		<TextLink href={resolve('/masuk/akun')}>{copy.signUp.signIn}</TextLink>
	</Screen>
{/if}

<style>
	.sebelum-tombol {
		margin: var(--jarak-4) 0 0;
	}
</style>
