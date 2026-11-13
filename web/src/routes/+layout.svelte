<script lang="ts">
	import { tick, type Snippet } from 'svelte';
	import '#lib/design/tokens.css';
	import { missingFeatures } from '#lib/device.js';
	import { registerServiceWorker } from '#lib/service-worker-registration.js';
	import { onTabSignal } from '#lib/tab-signal.js';
	import { wipeLocalData } from '#lib/wipe.js';

	let { children }: { children: Snippet } = $props();

	// SEC-160 (e): checked once at start; without a required feature nothing else loads (LY-97).
	const unsupported = missingFeatures().length > 0;

	// SEC-172 (a): set first, so the screen's data leaves the DOM before the wipe starts.
	let wiping = $state(false);

	async function wipeAndRestart(): Promise<void> {
		wiping = true;
		await tick();
		try {
			await wipeLocalData();
		} catch (error) {
			// The session check on the next open (1.9b) runs the wipe again.
			console.error(error);
		}
		// A fresh document empties memory and replaces the history entry (SEC-164 (c)).
		location.replace('/');
	}

	$effect(() => {
		if (unsupported) return;
		registerServiceWorker().catch((error: unknown) => {
			console.error(error);
		});
		// ponytail: app_lock is handled with the app lock in step 6 (SEC-174).
		return onTabSignal((signal) => {
			if (signal.kind !== 'app_lock') void wipeAndRestart();
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
{:else if !wiping}
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
