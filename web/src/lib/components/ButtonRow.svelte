<script lang="ts">
	// The row of a screen's or a dialog's buttons (08 §19.13.5 "Tombol", §25.4): in Ringkas an
	// upright stack, every button the full width; from Sedang side by side and equally wide.
	// Buttons 8 apart (§19.13.5 "Dialog").
	import type { Snippet } from 'svelte';

	let { children }: { children: Snippet } = $props();
</script>

<div>{@render children()}</div>

<style>
	div {
		display: flex;
		flex-direction: column;
		align-items: stretch;
		gap: var(--jarak-2);
		margin-top: var(--jarak-6);
	}
	@media (min-width: 600px) {
		div {
			flex-direction: row;
		}
		div > :global(button:not(:only-child)) {
			flex: 1 1 0;
		}
	}
</style>
