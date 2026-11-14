<script lang="ts">
	import { tick, type Snippet } from 'svelte';
	import { on } from 'svelte/events';
	import '#lib/design/tokens.css';
	import Button from '#lib/components/Button.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import SessionEnded from '#lib/components/SessionEnded.svelte';
	import { copy } from '#lib/copy.js';
	import { missingFeatures } from '#lib/device.js';
	import { registerServiceWorker } from '#lib/service-worker-registration.js';
	import { checkSession, session, showEnded } from '#lib/session.svelte.js';
	import { onTabSignal, type WipeKind } from '#lib/tab-signal.js';
	import { wipeLocalData } from '#lib/wipe.js';

	let { children }: { children: Snippet } = $props();

	// SEC-160 (e): checked once at start; without a required feature nothing else loads (LY-97).
	const unsupported = missingFeatures().length > 0;

	// SEC-172 (a): set first, so the screen's data leaves the DOM before the wipe starts.
	let wiping = $state(false);

	function logFailure(error: unknown): void {
		console.error(error);
	}

	function check(userInitiated = false): void {
		checkSession({ userInitiated }).catch(logFailure);
	}

	/**
	 * Every tab's answer to a wipe signal, the sending tab's included (SEC-173 (b)). A revoked
	 * session shows LY-94 in place with the unsent count once the wipe is done (plan decision 6);
	 * a sign-out the person asked for, or leaving the household, opens a fresh document at /,
	 * which empties memory and replaces the history entry (SEC-164 (c)). Sign-out (Task 7)
	 * arrives here through the sign_out signal.
	 */
	async function wipe(signal: { kind: WipeKind; unsentDropped: number }): Promise<void> {
		wiping = true;
		await tick();
		let wiped = true;
		try {
			await wipeLocalData();
		} catch (error) {
			wiped = false;
			logFailure(error);
		}
		// LY-94 says the data is deleted, so a failed wipe never shows it: the fresh document's
		// session check runs the wipe again (SEC-172 (a)).
		if (signal.kind !== 'session_revoked' || !wiped) {
			location.replace('/');
			return;
		}
		showEnded('revoked', signal.unsentDropped);
		wiping = false;
	}

	$effect(() => {
		if (unsupported) return;
		registerServiceWorker().catch(logFailure);
		// SEC-172 (b): the session is checked before anything else shows.
		check();
		// SEC-164 (d): a page restored from the back-forward cache shows nothing until the
		// session is checked again.
		const stopRestore = on(window, 'pageshow', (event) => {
			// Only a screen that depends on the session; LY-94 stays as it is.
			const kind = session.state.kind;
			if (!event.persisted || (kind !== 'signed-in' && kind !== 'unreachable')) return;
			session.state = { kind: 'checking' };
			check();
		});
		// ponytail: app_lock is handled with the app lock in step 6 (SEC-174).
		const stopSignals = onTabSignal((signal) => {
			if (signal.kind !== 'app_lock') wipe(signal).catch(logFailure);
		});
		return () => {
			stopRestore();
			stopSignals();
		};
	});
</script>

{#if unsupported}
	<Screen heading="Casa">
		<!-- 08 §25.6 LY-97 -->
		<p>
			Browser ini belum bisa menjalankan Casa dengan aman. Perbarui browser ini, atau buka Casa di
			browser lain yang lebih baru.
		</p>
		<p>Bila Casa terbuka di dalam aplikasi lain atau di jendela pribadi, buka di browser biasa.</p>
	</Screen>
{:else if wiping || session.state.kind === 'checking'}
	<!-- 08 §25.7: the session check shows the paper surface with the knot, no sentence, no spinner -->
	<div class="pemeriksaan" aria-busy="true"><span class="simpul"></span></div>
{:else if session.state.kind === 'unreachable'}
	<Screen heading="Casa">
		<p>{copy.home.unreachable}</p>
		<div class="tombol">
			<Button
				variant="primary"
				onclick={() => {
					check(true);
				}}>{copy.common.retry}</Button
			>
		</div>
	</Screen>
{:else if session.state.kind === 'ended'}
	<SessionEnded reason={session.state.reason} unsentDropped={session.state.unsentDropped} />
{:else}
	{@render children()}
{/if}

<style>
	.pemeriksaan {
		display: grid;
		place-items: center;
		min-height: 100dvh;
		background: var(--kertas);
	}
	.simpul {
		/* 08 §19.5: the knot in tinta-2, never aksen */
		width: var(--jarak-16);
		height: var(--jarak-16);
		background: var(--tinta-2);
		mask: url('/glyphs/simpul.svg') center / contain no-repeat;
	}
	@media (forced-colors: active) {
		.simpul {
			forced-color-adjust: none;
			background: CanvasText;
		}
	}
	.tombol {
		margin-top: var(--jarak-6);
	}
</style>
