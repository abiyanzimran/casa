<script lang="ts">
	// "/" chooses the screen (plan decision 13): a person without a session goes to the entry
	// screen nextScreen() names, replacing "/" in the history so Back does not bounce here.
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { nextScreen } from '#lib/session.svelte.js';

	const target = $derived(nextScreen());

	$effect(() => {
		if (target !== '/') goto(resolve(target), { replaceState: true }).catch(console.error);
	});
</script>

{#if target === '/'}
	<main><h1>Casa</h1></main>
{/if}
