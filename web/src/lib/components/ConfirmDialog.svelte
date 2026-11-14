<script lang="ts">
	// Confirmation of a heavy action (08 §9.3, §25.4): a modal dialog whose buttons name what they
	// do, never confirm(). Esc and the cancel button close it; focus then returns to the button
	// that opened it (08 §25.5). The caller opens it by setting opener to that button, because
	// Safari does not focus a clicked button. The cancel button is optional, for a one-button
	// notice (W15).
	import type { Snippet } from 'svelte';
	import Button from './Button.svelte';
	import ButtonRow from './ButtonRow.svelte';

	let {
		opener = $bindable(),
		title,
		confirmLabel,
		cancelLabel,
		onconfirm,
		children
	}: {
		/** The button that opened the dialog; set it to open, unset when the dialog closes. */
		opener?: HTMLElement;
		title: string;
		confirmLabel: string;
		cancelLabel?: string;
		onconfirm: () => void;
		children: Snippet;
	} = $props();

	const id = $props.id();
	let dialog: HTMLDialogElement;

	$effect(() => {
		if (opener && !dialog.open) dialog.showModal();
	});

	function returnFocus(): void {
		opener?.focus();
		opener = undefined;
	}

	function close(): void {
		dialog.close();
	}

	function confirm(): void {
		close();
		onconfirm();
	}
</script>

<dialog
	bind:this={dialog}
	aria-labelledby="{id}-title"
	aria-describedby="{id}-body"
	onclose={returnFocus}
>
	<h2 id="{id}-title" class="judul-bagian">{title}</h2>
	<div id="{id}-body">{@render children()}</div>
	<ButtonRow>
		{#if cancelLabel}
			<Button variant="outlined" onclick={close}>{cancelLabel}</Button>
		{/if}
		<Button variant="primary" onclick={confirm}>{confirmLabel}</Button>
	</ButtonRow>
</dialog>

<style>
	dialog {
		/* 08 §19.13.5 "Dialog": padding 24, width 280-560, corner 28, buttons 8 apart */
		box-sizing: border-box;
		width: min(560px, calc(100% - 2 * var(--margin)));
		min-width: min(280px, 100%);
		padding: var(--jarak-6);
		border: none;
		border-radius: var(--sudut-lembar);
		background: var(--permukaan-mengambang);
		color: var(--tinta);
		font-size: var(--ukuran-isi);
		line-height: 1.5;
	}
	dialog::backdrop {
		background: var(--selubung);
	}
	h2 {
		margin: 0 0 var(--jarak-4);
	}
	dialog :global(p) {
		margin: 0;
	}
	@media (forced-colors: active) {
		dialog {
			border: 1px solid CanvasText;
		}
	}
</style>
