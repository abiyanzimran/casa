// Screen text, word for word from 08, named by the keys of 08 §7.6 (plan decision 17). Text
// without a §7.6 key keeps the name of its place in 08. Tests write the sentences as literals
// and never import this file.

export const copy = {
	common: {
		retry: 'Coba lagi',
		// While a tapped re-check runs (QG-09 S3).
		retrying: 'Mencoba…',
		back: 'Kembali'
	},
	home: {
		// 08 §7.6 `home.unreachable`: opened without a connection; step 1 stores no data.
		unreachable: 'Belum tersambung. Rumah tanggamu tampil begitu tersambung lagi.',
		// Online, but the server errs or does not answer in time (§13.2 sentence).
		unreachableServer: 'Ada gangguan di server. Coba lagi beberapa saat lagi.',
		// A wipe failed a second time in this tab (SEC-172 (a)); never says the data is deleted.
		unreachableStorage:
			'Data Casa di perangkat ini belum bisa dibersihkan. Coba lagi. Kalau tetap gagal, tutup tab ini lalu buka Casa lagi.'
	},
	// 08 §25.7 error table row "Versi aplikasi di bawah minimum" (SEC-169 (d)).
	update: {
		// The second sentence only in personal mode once signed in, where a local queue exists.
		required:
			'Aplikasi perlu diperbarui untuk mengirim catatan. Catatanmu tersimpan di perangkat ini.',
		requiredShort: 'Aplikasi perlu diperbarui untuk mengirim catatan.',
		button: 'Perbarui',
		checking: 'Memeriksa…',
		deferred: 'Diperbarui setelah isianmu ditutup.'
	},
	// 08 §25.6 LY-94 "Sesi berakhir"; never says who or what ended the session.
	ended: {
		signIn: 'Masuk',
		revoked: 'Kamu sudah keluar di perangkat ini, dan data Casa di browser ini sudah dihapus.',
		dropped: (n: number) =>
			`${String(n)} catatan yang belum terkirim dari perangkat ini ikut terhapus.`,
		expired:
			'Waktu masukmu di perangkat ini sudah habis. Masuk lagi dengan akun yang sama; catatan yang belum terkirim tetap tersimpan dan terkirim sesudah kamu masuk.',
		sharedEnded: 'Kamu sudah keluar otomatis. Catatanmu tidak tersimpan di perangkat ini.',
		otherAccount: {
			button: 'Masuk dengan akun lain',
			title: 'Masuk dengan akun lain?',
			body: (n: number) =>
				n > 0
					? `Data Casa di perangkat ini akan dihapus, termasuk ${String(n)} catatan yang belum terkirim.`
					: 'Data Casa di perangkat ini akan dihapus.',
			confirm: 'Hapus dan masuk'
		}
	}
} as const;
