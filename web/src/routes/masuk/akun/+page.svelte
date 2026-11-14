<script lang="ts">
	// 08 §5.1 LY-02 "Masuk" with the words of 08 §7.6 signIn.*: email and password, then SignIn
	// with the device mode chosen on LY-91 or LY-90 (SEC-165 (a)). One sentence for a wrong email
	// or password (SEC-12); a failure that is no one field's sits above the button as an alert,
	// and only CREDENTIALS_REJECTED empties the password (08 §7.6). Shared mode saves no password
	// and says so (SEC-165 (e), SEC-175 (d)).
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import Button from '#lib/components/Button.svelte';
	import ButtonRow from '#lib/components/ButtonRow.svelte';
	import Screen from '#lib/components/Screen.svelte';
	import TextField from '#lib/components/TextField.svelte';
	import { api, ctx, errorCode, failureOf } from '#lib/api.js';
	import { copy } from '#lib/copy.js';
	import { ErrorCode } from '#lib/gen/casa/v1/error_pb.js';
	import { DeviceMode } from '#lib/gen/casa/v1/identity_pb.js';
	import { afterSignIn, entryRedirect, handleSessionError, session } from '#lib/session.svelte.js';

	const redirect = $derived(entryRedirect('/masuk/akun'));

	$effect(() => {
		if (redirect) goto(resolve(redirect), { replaceState: true }).catch(console.error);
	});

	// LY-02 shows only with a mode chosen (entryRedirect), so personal unless shared was chosen.
	const personal = $derived(session.mode !== 'shared');

	let email = $state('');
	let password = $state('');
	let failure: string | undefined = $state();
	let busy = $state(false);

	async function signIn(): Promise<void> {
		const mode = personal ? 'personal' : 'shared';
		failure = undefined; // so the same sentence again is announced again
		try {
			await api.identity.signIn({
				requestContext: ctx(true),
				email,
				password,
				deviceMode: mode === 'personal' ? DeviceMode.PERSONAL : DeviceMode.SHARED
			});
		} catch (error) {
			if (await handleSessionError(error)) return;
			// The sentence for a failed SignIn (08 §7.6, §13.2); only the unexpected is logged.
			const code = errorCode(error);
			if (code === ErrorCode.CREDENTIALS_REJECTED) {
				failure = copy.signIn.rejected;
				password = '';
			} else if (code === ErrorCode.RATE_LIMITED) failure = copy.common.rateLimited;
			else if (failureOf(error) === 'offline') failure = copy.serverAction.offline;
			else {
				failure = copy.serverAction.server;
				console.error(error);
			}
			return;
		}
		await afterSignIn(mode);
	}
</script>

{#if !redirect}
	<Screen heading={copy.signIn.title} destination="Masuk">
		<form
			novalidate
			onsubmit={(event) => {
				event.preventDefault();
				if (busy) return;
				busy = true;
				signIn()
					.catch(console.error)
					.finally(() => {
						busy = false;
					});
			}}
		>
			<TextField
				label={copy.signIn.email}
				type="email"
				autocomplete={personal ? 'username' : 'off'}
				bind:value={email}
			/>
			<TextField
				label={copy.signIn.password}
				type="password"
				autocomplete={personal ? 'current-password' : 'off'}
				hint={personal ? undefined : copy.signIn.doNotSave}
				bind:value={password}
			/>
			<div role="alert">
				{#if failure}<p class="galat">{failure}</p>{/if}
			</div>
			<ButtonRow>
				<Button variant="primary" type="submit" {busy}>{copy.signIn.title}</Button>
			</ButtonRow>
		</form>
	</Screen>
{/if}

<style>
	.galat {
		margin: var(--jarak-4) 0 0;
		color: var(--galat);
	}
</style>
