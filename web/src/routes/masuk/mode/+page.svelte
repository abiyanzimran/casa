<script lang="ts">
	// 08 §25.6 LY-91 "Mode perangkat" (SEC-165): two equal choices, neither preselected, and the
	// limit the app cannot prevent. The choice lives in memory for this sign-in (session.mode).
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import ChoiceButton from '#lib/components/ChoiceButton.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import { copy } from '#lib/copy.js';
	import { entryRedirect, session, type DeviceMode } from '#lib/session.svelte.js';

	const redirect = $derived(entryRedirect('/masuk/mode'));
	const id = $props.id();

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});

	// A verification link opened before a mode was chosen goes back to its page (08 §7.6
	// "Halaman verifikasi"); otherwise LY-01.
	function choose(mode: DeviceMode): void {
		session.mode = mode;
		const back = session.pendingToken?.kind === 'verification';
		goto(resolve(back ? '/signup/verify' : '/masuk')).catch(console.error);
	}
</script>

{#if !redirect}
	<Screen heading={copy.ly91.question} destination="Mode perangkat" describedby="{id}-limit">
		<div class="pilihan">
			<ChoiceButton
				title={copy.ly91.personal.title}
				note={copy.ly91.personal.note}
				onclick={() => {
					choose('personal');
				}}
			/>
			<ChoiceButton
				title={copy.ly91.shared.title}
				note={copy.ly91.shared.note}
				onclick={() => {
					choose('shared');
				}}
			/>
		</div>
		<p id="{id}-limit">{copy.ly91.limit}</p>
	</Screen>
{/if}

<style>
	.pilihan {
		display: flex;
		flex-direction: column;
		gap: var(--jarak-3);
		margin-bottom: var(--jarak-6);
	}
</style>
