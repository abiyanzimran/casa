<script lang="ts">
	// LY-94 Sesi berakhir (08 §25.6), shown in place without a reload (plan decision 6). The
	// memory of the session is already reset; "Masuk" opens a fresh document at / (SEC-164 (c)).
	import { copy } from '#lib/copy.js';
	import type { SessionEndReason } from '#lib/session.svelte.js';
	import { broadcastWipe, unsentCount } from '#lib/wipe.js';
	import Button from './Button.svelte';
	import ConfirmDialog from './ConfirmDialog.svelte';
	import Screen from './Screen.svelte';

	let { reason, unsentDropped }: { reason: SessionEndReason; unsentDropped: number } = $props();

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

<Screen heading="Casa">
	{#if reason === 'revoked'}
		<p>{copy.ended.revoked}</p>
		{#if unsentDropped > 0}
			<p>{copy.ended.dropped(unsentDropped)}</p>
		{/if}
	{:else if reason === 'expired'}
		<p>{copy.ended.expired}</p>
	{:else}
		<p>{copy.ended.sharedEnded}</p>
	{/if}
	<div class="tombol">
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
	</div>
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

<style>
	.tombol {
		display: flex;
		flex-wrap: wrap;
		gap: var(--jarak-2);
		margin-top: var(--jarak-6);
	}
</style>
