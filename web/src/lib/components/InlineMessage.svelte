<script lang="ts">
	// A sentence a screen announces in place (08 §7.6, §25.5 "Pengumuman"): role="alert" for a
	// failure that stops the step and is no one field's, above the main button; role="status" for
	// a result. The container is always on the page and filled when needed, so the same sentence
	// set again after being cleared is announced again (WCAG 4.1.3).
	let { role, text }: { role: 'alert' | 'status'; text?: string } = $props();
</script>

<div {role}>
	{#if text}<p data-role={role}>{text}</p>{/if}
</div>

<style>
	p {
		margin: var(--jarak-4) 0 0;
	}
	[data-role='alert'] {
		color: var(--galat);
	}
</style>
