<script lang="ts">
	import { tick, type Snippet } from 'svelte';
	import { on } from 'svelte/events';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import '#lib/design/tokens.css';
	import Button from '#lib/components/Button.svelte';
	import StatusRow from '#lib/components/StatusRow.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import SessionEnded from '#lib/components/SessionEnded.svelte';
	import { updateNow, upgrade } from '#lib/app-update.svelte.js';
	import { copy } from '#lib/copy.js';
	import { missingFeatures } from '#lib/device.js';
	import { registerServiceWorker } from '#lib/service-worker-registration.js';
	import {
		carryUnsentDropped,
		checkSession,
		session,
		showEnded,
		showUnreachable
	} from '#lib/session.svelte.js';
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
		if (signal.kind !== 'session_revoked') {
			location.replace('/');
			return;
		}
		const unsentDropped = carryUnsentDropped(signal.unsentDropped);
		const current = session.state;
		if (wiped) showEnded('revoked', unsentDropped);
		else if (current.kind === 'ended' && current.reason === 'revoked') {
			// An earlier wipe of this tab finished and LY-94 says so truthfully; a later one that
			// fails has nothing left to say, so LY-94 stays (the failure is logged above).
		} else if (!(await retryWipeOnce(unsentDropped))) return; // hidden until the reload
		wiping = false;
	}

	/**
	 * LY-94 says the data is deleted, so a failed revoked wipe never shows it (SEC-172 (a)).
	 * The first failure reloads at /, where the session check runs the wipe again (the
	 * wipe_pending marker, or the server's next session_revoked). The flag is navigation state
	 * in this tab's history entry, not storage (10.5 §5), and survives that one reload; a second
	 * failure in this tab shows the unreachable state with its retry instead of reloading in a
	 * loop. Returns whether it showed that state rather than reloading.
	 */
	async function retryWipeOnce(unsentDropped: number): Promise<boolean> {
		if (page.state.wipeRetried) {
			showUnreachable();
			return true;
		}
		// The unsent count goes along, so the reloaded page's LY-94 still shows it (ADR-024 item 6).
		await goto(resolve('/'), {
			replace: true,
			shallow: true,
			state: { wipeRetried: true, unsentDropped },
			persistState: true
		});
		location.reload();
		return false;
	}

	$effect(() => {
		if (unsupported) return;
		registerServiceWorker().catch(logFailure);
		carryUnsentDropped(page.state.unsentDropped ?? 0); // from a wipe that failed and reloaded
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
	{#if upgrade.required}
		<!-- 08 §25.7: applied only without an open form (SEC-169 (c)) -->
		<StatusRow
			text={copy.update.required}
			action={{
				label: copy.update.button,
				onclick: () => {
					updateNow().catch(logFailure);
				}
			}}
		/>
	{/if}
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
