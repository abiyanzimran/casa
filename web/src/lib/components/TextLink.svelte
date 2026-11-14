<script lang="ts">
	// A link to another screen, in a column under a screen's buttons (08 §7.6 signIn.links,
	// signUp.closed, checkEmail.change): tinta with an underline (08 §19.12), a 48 px target
	// (08 §25.5), and the 2 px aksen-tinta focus ring. href is a resolve()d route.
	import type { Snippet } from 'svelte';

	let { href, children }: { href: string; children: Snippet } = $props();
</script>

<!-- eslint-disable-next-line svelte/no-navigation-without-resolve -- callers pass resolve() -->
<a {href}>{@render children()}</a>

<style>
	a {
		display: flex;
		align-items: center;
		min-height: var(--sasaran);
		margin-top: var(--jarak-2);
		color: var(--tinta);
		font-weight: 600;
		text-decoration: underline;
		text-underline-offset: 0.2em;
	}
	a:focus-visible {
		/* 08 §25.5: 2 px aksen-tinta ring, 2 px away */
		outline: 2px solid var(--aksen-tinta);
		outline-offset: 2px;
	}
</style>
