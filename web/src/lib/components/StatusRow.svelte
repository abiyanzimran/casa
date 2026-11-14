<script lang="ts">
	// The status row under the top bar (08 §25.7): one sentence and at most one button, never a
	// modal. Used by the update row and, later, the unsent-count and notice rows.
	import Button from './Button.svelte';

	let { text, action }: { text: string; action?: { label: string; onclick: () => void } } =
		$props();
</script>

<div class="baris" role="status">
	<p>{text}</p>
	{#if action}
		<Button variant="tonal" onclick={action.onclick}>{action.label}</Button>
	{/if}
</div>

<style>
	.baris {
		display: flex;
		flex-wrap: wrap;
		align-items: center;
		justify-content: space-between;
		gap: var(--jarak-3);
		padding: var(--jarak-3) var(--margin);
		background: var(--aksen-muda);
		color: var(--tinta);
	}
	p {
		margin: 0;
		flex: 1 1 16rem;
	}
</style>
