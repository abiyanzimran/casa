<script lang="ts">
	// The one-column frame of a screen (08 §17.8, §25.5): at most --lebar-satu-kolom wide with the
	// screen margin, one h1 that takes focus when the screen opens, and the page title
	// "Casa · {destination}" or "Casa" alone (SEC-175 (a); 08 §7.6 title table). describedby names
	// the sentence paragraphs, so focusing the h1 announces them too (QG-09 S2).
	import type { Snippet } from 'svelte';

	let {
		heading,
		destination,
		describedby,
		children
	}: { heading: string; destination?: string; describedby?: string; children: Snippet } = $props();

	let title: HTMLHeadingElement;

	$effect(() => {
		title.focus();
	});
</script>

<svelte:head>
	<title>{destination ? `Casa · ${destination}` : 'Casa'}</title>
</svelte:head>

<main>
	<h1 class="judul-layar" tabindex="-1" aria-describedby={describedby} bind:this={title}>
		{heading}
	</h1>
	{@render children()}
</main>

<style>
	main {
		box-sizing: border-box;
		max-width: var(--lebar-satu-kolom);
		margin: 0 auto;
		padding: var(--jarak-12) var(--margin);
		font-size: var(--ukuran-isi);
		line-height: 1.5;
	}
	h1 {
		margin: 0 0 var(--jarak-4);
	}
	h1:focus {
		outline: none; /* focused by the app to announce the screen, never by Tab */
	}
	main > :global(p) {
		margin: 0;
	}
	main > :global(p + p) {
		margin-top: var(--jarak-3);
	}
</style>
