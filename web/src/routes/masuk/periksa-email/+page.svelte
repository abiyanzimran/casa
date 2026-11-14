<script lang="ts">
	// 08 §5.1 LY-03 "Periksa email" with the words of 08 §7.6 checkEmail.*: a link was sent, never
	// that an account exists or the email is known (SEC-12). "Kirim ulang" sends the same
	// StartSignUp; "Ganti email" goes back to LY-02 Buat akun with the email kept.
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import Button from '#lib/components/Button.svelte';
	import ButtonRow from '#lib/components/ButtonRow.svelte';
	import InlineMessage from '#lib/components/InlineMessage.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import TextLink from '#lib/components/TextLink.svelte';
	import { copy } from '#lib/copy.js';
	import { entryRedirect, session } from '#lib/session.svelte.js';
	import { heldSignUpToken, startSignUp, type SignUpToken } from '#lib/sign-up.js';

	const held = $derived(heldSignUpToken());
	const email = $derived(session.signUpEmail);
	// Memory only: after a reload nothing is held, so LY-02 Buat akun again.
	const redirect = $derived(
		entryRedirect('/masuk/periksa-email') ?? (held && email ? undefined : '/masuk/daftar')
	);
	const id = $props.id();

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});

	let failure: string | undefined = $state();
	let resent = $state(false);
	let busy = $state(false);

	async function resend(token: SignUpToken, to: string): Promise<void> {
		failure = undefined;
		resent = false;
		const failed = await startSignUp(to, token);
		if (failed) failure = failed.alert ?? failed.email;
		else resent = true;
	}
</script>

{#if !redirect && held && email}
	<Screen
		heading={copy.checkEmail.title}
		destination="Periksa email"
		describedby="{id}-sent {id}-help"
		back={resolve('/masuk/daftar')}
	>
		<p id="{id}-sent" class="kalimat-email">{copy.checkEmail.sent(email)}</p>
		<p id="{id}-help">{copy.checkEmail.help}</p>
		<InlineMessage role="alert" text={failure} />
		<InlineMessage role="status" text={resent ? copy.checkEmail.resent : undefined} />
		<ButtonRow>
			<Button
				variant="tonal"
				{busy}
				onclick={() => {
					busy = true;
					resend(held, email)
						.catch(console.error)
						.finally(() => {
							busy = false;
						});
				}}>{copy.checkEmail.resend}</Button
			>
		</ButtonRow>
		<TextLink href={resolve('/masuk/daftar')}>{copy.checkEmail.change}</TextLink>
	</Screen>
{/if}

<style>
	.kalimat-email {
		/* a long email wraps, never cut (08 §7.6) */
		overflow-wrap: anywhere;
	}
</style>
