<script lang="ts">
	// The status row under the top bar (08 §25.7): one sentence and at most one button, never a
	// modal. Used by the update row and, later, the unsent-count and notice rows. The role="status"
	// container is always on the page and filled when needed, so a change of text is announced
	// (WCAG 4.1.3; QG-09 S6). The button is outlined: the row is already aksen-muda (QG-09 S1).
	import Button from './Button.svelte';

	let {
		text,
		action
	}: {
		text?: string;
		action?: { label: string; onclick: () => void; busy?: boolean };
	} = $props();
</script>

<div role="status">
	{#if text}
		<div class="baris">
			<div class="isi">
				<p>{text}</p>
				{#if action}
					<Button variant="outlined" busy={action.busy} onclick={action.onclick}>
						{action.label}
					</Button>
				{/if}
			</div>
		</div>
	{/if}
</div>

<style>
	.baris {
		padding: var(--jarak-3) 0;
		background: var(--aksen-muda);
		color: var(--tinta);
	}
	.isi {
		/* the content lines up with the screen's one column (Screen.svelte) */
		box-sizing: border-box;
		display: flex;
		flex-wrap: wrap;
		align-items: center;
		justify-content: space-between;
		gap: var(--jarak-3);
		max-width: var(--lebar-satu-kolom);
		margin: 0 auto;
		padding: 0 var(--margin);
	}
	p {
		margin: 0;
		flex: 1 1 16rem;
	}
</style>
