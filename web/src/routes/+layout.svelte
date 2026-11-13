<script lang="ts">
	import type { Snippet } from 'svelte';
	import '#lib/design/tokens.css';
	import { missingFeatures } from '#lib/device.js';
	import { registerServiceWorker } from '#lib/service-worker-registration.js';

	let { children }: { children: Snippet } = $props();

	// SEC-160 (e): checked once at start; without a required feature nothing else loads (LY-97).
	const unsupported = missingFeatures().length > 0;

	$effect(() => {
		if (unsupported) return;
		registerServiceWorker().catch((error: unknown) => {
			console.error(error);
		});
	});
</script>

{#if unsupported}
	<main class="unsupported">
		<h1>Casa</h1>
		<!-- 08 §25.6 LY-97 -->
		<p>
			Browser ini belum bisa menjalankan Casa dengan aman. Perbarui browser ini, atau buka Casa di
			browser lain yang lebih baru.
		</p>
	</main>
{:else}
	{@render children()}
{/if}

<style>
	.unsupported {
		box-sizing: border-box;
		max-width: var(--lebar-satu-kolom);
		margin: 0 auto;
		padding: var(--jarak-12) var(--margin);
		font-family: var(--huruf);
		font-size: var(--ukuran-isi);
		line-height: 1.5;
		color: var(--tinta);
	}
	h1 {
		margin: 0 0 var(--jarak-4);
		font-size: var(--ukuran-judul-layar);
		font-weight: 700;
	}
	p {
		margin: 0;
	}
</style>
