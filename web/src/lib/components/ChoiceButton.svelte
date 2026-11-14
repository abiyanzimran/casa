<script lang="ts">
	// A large outlined button the column's width, with a title and a note under it, for a choice
	// between equal ways (08 §25.6 LY-90 second way, LY-91): never a selected state, so nothing is
	// chosen for the person (SEC-165). The title is the accessible name, the note its description.
	let { title, note, onclick }: { title: string; note: string; onclick: () => void } = $props();

	const id = $props.id();
</script>

<button type="button" aria-labelledby="{id}-title" aria-describedby="{id}-note" {onclick}>
	<span class="judul" id="{id}-title">{title}</span>
	<span class="keterangan" id="{id}-note">{note}</span>
</button>

<style>
	button {
		box-sizing: border-box;
		display: flex;
		flex-direction: column;
		gap: var(--jarak-1);
		width: 100%;
		min-height: var(--sasaran);
		padding: var(--jarak-4);
		border: 1px solid var(--garis-tegas);
		border-radius: var(--sudut-besar);
		background: var(--kartu);
		color: var(--tinta);
		font: inherit;
		text-align: start;
		cursor: pointer;
		/* 08 §19.13.7 "Tekan kendali", as Button */
		transition: transform var(--gerak) var(--lengkung);
	}
	button:active {
		transform: scale(0.97);
		transition-duration: var(--gerak-tekan);
	}
	button:focus-visible {
		/* 08 §25.5: 2 px aksen-tinta ring, 2 px away */
		outline: 2px solid var(--aksen-tinta);
		outline-offset: 2px;
	}
	.judul {
		font-weight: 600;
		color: var(--aksen-tinta);
	}
	.keterangan {
		color: var(--tinta-2);
	}
	@media (prefers-reduced-motion: reduce) {
		button:active {
			transform: none;
			opacity: 0.8;
		}
	}
</style>
