<script lang="ts">
	let { s }: { s: string } = $props();
	let h = $state('');
</script>

{@html s}
<div contenteditable="true" bind:innerHTML={h}></div>
<iframe srcdoc={s} title="x"></iframe>
