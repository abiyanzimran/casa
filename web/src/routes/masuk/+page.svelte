<script lang="ts">
	// 08 §5.1 LY-01 "Sambutan" with the words of 08 §7.6 ly01.*: one line on what Casa is, then
	// "Buat akun" and "Sudah punya akun". Shown once a device mode is chosen (SEC-165).
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import Button from '#lib/components/Button.svelte';
	import ButtonRow from '#lib/components/ButtonRow.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import { copy } from '#lib/copy.js';
	import { entryRedirect } from '#lib/session.svelte.js';

	const redirect = $derived(entryRedirect('/masuk'));
	const id = $props.id();

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});

	// ponytail: Task 5 (/masuk/akun) and Task 6 (/masuk/daftar) add the screens these buttons
	// open; resolve() accepts only routes that exist, so the taps are wired there.
	const notYetBuilt = () => undefined;
</script>

{#if !redirect}
	<Screen heading={copy.ly01.title} destination="Selamat datang" describedby="{id}-line">
		<p id="{id}-line">{copy.ly01.line}</p>
		<ButtonRow>
			<Button variant="primary" onclick={notYetBuilt}>{copy.ly01.signUp}</Button>
			<Button variant="outlined" onclick={notYetBuilt}>{copy.ly01.signIn}</Button>
		</ButtonRow>
	</Screen>
{/if}
