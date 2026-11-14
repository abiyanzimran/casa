<script lang="ts">
	// 08 §25.6 LY-90, in a Safari tab on iPhone and iPad (SEC-171 (b)): install to the home
	// screen, or the shared way in this tab, which opens LY-01 in shared mode without LY-91
	// (SEC-171 (b) (5)). Personal mode is never offered here; nothing is sent and nothing stored.
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import ChoiceButton from '#lib/components/ChoiceButton.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import { copy } from '#lib/copy.js';
	import { entryRedirect, session } from '#lib/session.svelte.js';

	const redirect = $derived(entryRedirect('/pasang'));
	const id = $props.id();
	// Opened from a setup link (/signup), whose token was dropped unsent (08 §7.6 ly90.fromSetup).
	const fromSetup = session.setupLinkDropped === true;

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});

	function useShared(): void {
		session.mode = 'shared';
		goto(resolve('/masuk')).catch(console.error);
	}
</script>

{#if !redirect}
	<Screen
		heading={copy.ly90.title}
		destination="Pasang"
		describedby="{id}-first {id}-reason{fromSetup ? ` ${id}-setup` : ''}"
	>
		{#snippet lead()}
			<p id="{id}-first">{copy.ly90.firstLine}</p>
		{/snippet}
		<p id="{id}-reason">{copy.ly90.reason}</p>
		{#if fromSetup}<p id="{id}-setup">{copy.ly90.fromSetup}</p>{/if}
		<ol>
			{#each copy.ly90.steps as step (step)}
				<li>{step}</li>
			{/each}
		</ol>
		<ChoiceButton title={copy.ly90.shared.title} note={copy.ly90.shared.note} onclick={useShared} />
	</Screen>
{/if}

<style>
	ol {
		/* the numbers are part of the text (08 §25.6 "Langkah") */
		list-style: none;
		margin: var(--jarak-4) 0 var(--jarak-6);
		padding: 0;
	}
	li + li {
		margin-top: var(--jarak-2);
	}
</style>
