<script lang="ts">
	// 08 §5.1 LY-01 "Sambutan" with the words of 08 §7.6 ly01.*: one line on what Casa is, then
	// "Buat akun" and "Sudah punya akun". Shown once a device mode is chosen (SEC-165).
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import Button from '#lib/components/Button.svelte';
	import ButtonRow from '#lib/components/ButtonRow.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import { copy } from '#lib/copy.js';
	import { displayContext } from '#lib/device.js';
	import { entryRedirect } from '#lib/session.svelte.js';

	const redirect = $derived(entryRedirect('/masuk'));
	const id = $props.id();
	// The step before (08 §25.4): LY-90 in a Safari tab, LY-91 elsewhere.
	const back = resolve(displayContext() === 'ios-safari-tab' ? '/pasang' : '/masuk/mode');

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});
</script>

{#if !redirect}
	<Screen heading={copy.ly01.title} destination="Selamat datang" describedby="{id}-line" {back}>
		<p id="{id}-line">{copy.ly01.line}</p>
		<ButtonRow>
			<Button
				variant="primary"
				onclick={() => {
					goto(resolve('/masuk/daftar')).catch(console.error);
				}}>{copy.ly01.signUp}</Button
			>
			<Button
				variant="outlined"
				onclick={() => {
					goto(resolve('/masuk/akun')).catch(console.error);
				}}>{copy.ly01.signIn}</Button
			>
		</ButtonRow>
	</Screen>
{/if}
