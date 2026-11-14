<script lang="ts">
	// LY-94 Sesi berakhir (08 §25.6), shown in place without a reload (plan decision 6). The
	// memory of the session is already reset; "Masuk" opens a fresh document at / (SEC-164 (c)).
	import { copy } from '#lib/copy.js';
	import type { SessionEndReason } from '#lib/session.svelte.js';
	import { broadcastWipe, unsentCount } from '#lib/wipe.js';
	import Button from './Button.svelte';
	import ButtonRow from './ButtonRow.svelte';
	import ConfirmDialog from './ConfirmDialog.svelte';
	import Screen from './Screen.svelte';

	let { reason, unsentDropped }: { reason: SessionEndReason; unsentDropped: number } = $props();

	const id = $props.id();
	const sentence = `${id}-sentence`;
	const dropped = `${id}-dropped`;
	let opener = $state<HTMLElement>();
	let unsent = $state(0);

	function signIn(): void {
		location.replace('/');
	}

	async function askOtherAccount(button: HTMLElement): Promise<void> {
		unsent = await unsentCount();
		opener = button;
	}

	function deleteAndSignIn(): void {
		// SEC-170 (a): one account per browser profile. The sign_out signal wipes every tab and
		// opens a fresh document at / (SEC-164 (c)); the expired session needs no server call.
		broadcastWipe('sign_out').catch((error: unknown) => {
			console.error(error);
		});
	}
</script>

<Screen
	heading="Casa"
	describedby={reason === 'revoked' && unsentDropped > 0 ? `${sentence} ${dropped}` : sentence}
>
	{#if reason === 'revoked'}
		<p id={sentence}>{copy.ended.revoked}</p>
		{#if unsentDropped > 0}
			<p id={dropped}>{copy.ended.dropped(unsentDropped)}</p>
		{/if}
	{:else if reason === 'expired'}
		<p id={sentence}>{copy.ended.expired}</p>
	{:else}
		<p id={sentence}>{copy.ended.sharedEnded}</p>
	{/if}
	<ButtonRow>
		<Button variant="primary" onclick={signIn}>{copy.ended.signIn}</Button>
		{#if reason === 'expired'}
			<Button
				variant="outlined"
				onclick={(event: MouseEvent & { currentTarget: HTMLButtonElement }) => {
					askOtherAccount(event.currentTarget).catch((error: unknown) => {
						console.error(error);
					});
				}}>{copy.ended.otherAccount.button}</Button
			>
		{/if}
	</ButtonRow>
</Screen>

{#if reason === 'expired'}
	<ConfirmDialog
		bind:opener
		title={copy.ended.otherAccount.title}
		confirmLabel={copy.ended.otherAccount.confirm}
		cancelLabel={copy.common.back}
		onconfirm={deleteAndSignIn}
	>
		<p>{copy.ended.otherAccount.body(unsent)}</p>
	</ConfirmDialog>
{/if}
