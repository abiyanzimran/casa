<script lang="ts">
	// The setup link <origin>/signup#<token> (09.4 §2; plan decision 8). The init hook
	// (hooks.client.ts) took the token before anything ran; this page sends the person on: to
	// LY-02 Buat akun with the token held, or in a Safari tab on iPhone or iPad to LY-90 with
	// ly90.fromSetup, the token dropped unsent (08 §7.6). Replacing this entry keeps Back from
	// landing here again.
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { displayContext } from '#lib/device.js';

	$effect(() => {
		const tab = displayContext() === 'ios-safari-tab';
		goto(resolve(tab ? '/pasang' : '/masuk/daftar'), { replaceState: true }).catch(console.error);
	});
</script>
