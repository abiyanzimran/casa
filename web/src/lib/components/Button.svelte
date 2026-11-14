<script lang="ts">
	// The three button styles of 08 §25.4 (component table, 13 November 2026): primary, filled
	// aksen, at most one per screen; tonal, aksen-muda with aksen-tinta; outlined, garis-tegas with
	// aksen-tinta. All the same height, a full pill (08 §19.13.5).
	import type { Snippet } from 'svelte';

	let {
		variant,
		onclick,
		children
	}: {
		variant: 'primary' | 'outlined' | 'tonal';
		onclick: (event: MouseEvent & { currentTarget: HTMLButtonElement }) => void;
		children: Snippet;
	} = $props();
</script>

<!-- A data attribute, not class={variant}: that would pull clsx into the bundle (ADR-030 item 8). -->
<button type="button" data-variant={variant} {onclick}>{@render children()}</button>

<style>
	button {
		box-sizing: border-box;
		min-height: var(--tinggi-tombol);
		padding: 0 var(--jarak-6);
		border: 1px solid transparent;
		border-radius: var(--penuh);
		font: inherit;
		font-weight: 600;
		cursor: pointer;
		transition: background-color var(--gerak-warna) var(--lengkung);
	}
	button:focus-visible {
		/* 08 §25.5: 2 px aksen-tinta ring, 2 px away */
		outline: 2px solid var(--aksen-tinta);
		outline-offset: 2px;
	}
	[data-variant='primary'] {
		background: var(--aksen);
		color: var(--di-atas-aksen);
	}
	[data-variant='tonal'] {
		background: var(--aksen-muda);
		color: var(--aksen-tinta);
	}
	[data-variant='outlined'] {
		background: transparent;
		border-color: var(--garis-tegas);
		color: var(--aksen-tinta);
	}
</style>
