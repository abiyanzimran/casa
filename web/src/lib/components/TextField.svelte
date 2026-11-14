<script lang="ts">
	// A one-line text field (08 §19.13.5 "Isian teks", §25.4, §25.5): the label visible above,
	// 48 high, text at 1 rem so Safari on iOS does not zoom in on focus. The hint and the error
	// sit under the field and describe it (aria-describedby); an error also sets aria-invalid
	// (08 §7.6). A password field has a show button (08 §7.6 field.showPassword). Paste is
	// never blocked (WCAG 3.3.8). autocomplete is the caller's, per SEC-175 (d).
	import type { FullAutoFill } from 'svelte/elements';
	import { copy } from '../copy';

	let {
		label,
		type,
		autocomplete,
		value = $bindable(''),
		hint,
		error
	}: {
		label: string;
		type: 'email' | 'password' | 'text' | 'url';
		autocomplete: FullAutoFill;
		value?: string;
		hint?: string;
		error?: string;
	} = $props();

	const id = $props.id();
	let shown = $state(false);

	const describedby = $derived(
		[hint && `${id}-hint`, error && `${id}-error`].filter(Boolean).join(' ') || undefined
	);
	const show = copy.field.showPassword;
</script>

<div class="isian">
	<label for="{id}-input">{label}</label>
	<div class="baris">
		<input
			id="{id}-input"
			type={type === 'password' && shown ? 'text' : type}
			{autocomplete}
			autocapitalize={type === 'text' ? undefined : 'off'}
			spellcheck={type === 'text' ? undefined : false}
			aria-describedby={describedby}
			aria-invalid={error ? 'true' : undefined}
			bind:value
		/>
		{#if type === 'password'}
			<button
				type="button"
				aria-label={shown ? show.hideName : show.showName}
				onclick={() => {
					shown = !shown;
				}}>{shown ? show.hide : show.show}</button
			>
		{/if}
	</div>
	{#if hint}<p class="keterangan" id="{id}-hint">{hint}</p>{/if}
	{#if error}<p class="galat" id="{id}-error">{error}</p>{/if}
</div>

<style>
	.isian {
		display: flex;
		flex-direction: column;
		gap: var(--jarak-1);
		margin-top: var(--jarak-4);
	}
	label {
		font-weight: 600;
	}
	.baris {
		display: flex;
		gap: var(--jarak-2);
	}
	input {
		box-sizing: border-box;
		flex: 1 1 auto;
		min-width: 0;
		min-height: var(--tinggi-isian);
		padding: 0 var(--jarak-4);
		border: 1px solid var(--garis-tegas);
		border-radius: var(--sudut-sedang);
		background: var(--kartu);
		color: var(--tinta);
		font: inherit;
		font-size: max(1rem, var(--ukuran-isi));
	}
	input[aria-invalid='true'] {
		border-color: var(--galat);
	}
	input:focus-visible,
	button:focus-visible {
		/* 08 §25.5: 2 px aksen-tinta ring, 2 px away */
		outline: 2px solid var(--aksen-tinta);
		outline-offset: 2px;
	}
	button {
		flex: none; /* the field gives way, never the button's label (320 px, 200% text) */
		min-width: var(--sasaran);
		min-height: var(--sasaran);
		padding: 0 var(--jarak-3);
		border: 0;
		border-radius: var(--penuh);
		background: transparent;
		color: var(--aksen-tinta);
		font: inherit;
		font-weight: 600;
		cursor: pointer;
	}
	.keterangan,
	.galat {
		margin: 0;
	}
	.keterangan {
		color: var(--tinta-2);
	}
	.galat {
		color: var(--galat);
	}
</style>
