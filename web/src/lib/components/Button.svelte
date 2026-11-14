<script lang="ts">
	// The three button styles of 08 §25.4 (component table, 13 November 2026): primary, filled
	// aksen, at most one per screen; tonal, aksen-muda with aksen-tinta; outlined, garis-tegas with
	// aksen-tinta. All the same height, a full pill (08 §19.13.5). A busy button is not active and
	// says so (aria-busy, aria-disabled), but keeps focus, which a disabled attribute would drop.
	// A submit button sends its form; the form's submit handler checks busy itself, because
	// Enter in a field submits without a click.
	import type { Snippet } from 'svelte';

	type Click = (event: MouseEvent & { currentTarget: HTMLButtonElement }) => void;

	let {
		variant,
		busy = false,
		children,
		...action
	}: {
		variant: 'primary' | 'outlined' | 'tonal';
		busy?: boolean;
		children: Snippet;
	} & ({ type?: 'button'; onclick: Click } | { type: 'submit'; onclick?: never }) = $props();
</script>

<!-- A data attribute, not class={variant}: that would pull clsx into the bundle (ADR-030 item 8). -->
<button
	type={action.type ?? 'button'}
	data-variant={variant}
	aria-busy={busy ? 'true' : undefined}
	aria-disabled={busy ? 'true' : undefined}
	onclick={(event) => {
		if (busy) event.preventDefault();
		else action.onclick?.(event);
	}}>{@render children()}</button
>

<style>
	button {
		box-sizing: border-box;
		min-height: var(--tinggi-tombol);
		/* Block padding, so a label that wraps to two lines keeps room (08 §25.4) */
		padding: var(--jarak-2) var(--jarak-6);
		border: 1px solid transparent;
		border-radius: var(--penuh);
		font: inherit;
		font-weight: 600;
		cursor: pointer;
		/* 08 §19.13.7 "Tekan kendali": back in 200 ms on release */
		transition:
			background-color var(--gerak-warna) var(--lengkung),
			transform var(--gerak) var(--lengkung);
	}
	button:active:not([aria-disabled='true']) {
		/* pressed: scale(0.97) in 100 ms */
		transform: scale(0.97);
		transition-duration: var(--gerak-warna), var(--gerak-tekan);
	}
	button[aria-disabled='true'] {
		cursor: default;
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
	@media (prefers-reduced-motion: reduce) {
		button:active:not([aria-disabled='true']) {
			/* 08 §19.13.7: no scale; opacity 0.8 at once while pressed */
			transform: none;
			opacity: 0.8;
		}
	}
</style>
