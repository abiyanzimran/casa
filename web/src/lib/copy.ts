// Screen text, word for word from 08, named by the keys of 08 §7.6 (plan decision 17). Text
// without a §7.6 key keeps the name of its place in 08. Tests write the sentences as literals
// and never import this file.

export const copy = {
	common: {
		retry: 'Coba lagi',
		// While a tapped re-check runs (QG-09 S3).
		retrying: 'Mencoba…',
		back: 'Kembali',
		// A submit button while its request runs (QG-09 PR B S2).
		working: 'Memproses…',
		rateLimited: 'Terlalu banyak percobaan. Tunggu beberapa menit, lalu coba lagi.'
	},
	// 08 §7.6 field.showPassword: the visible text, and the accessible name with the field's.
	field: {
		showPassword: {
			show: 'Tampilkan',
			hide: 'Sembunyikan',
			showName: 'Tampilkan kata sandi',
			hideName: 'Sembunyikan kata sandi'
		}
	},
	// 08 §13.2 row "Tindakan yang memerlukan server": without a connection, and with the server's
	// fault (also RETRY_EXHAUSTED and errors without a code, 08 §7.6).
	serverAction: {
		offline: 'Ini perlu sambungan internet. Coba lagi saat tersambung.',
		server: 'Ada gangguan di server. Coba lagi beberapa saat lagi.'
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
	// 08 §25.6 LY-90 "Pasang ke layar utama", in a Safari tab on iPhone and iPad (SEC-171 (b)).
	ly90: {
		// Above both ways: covers a browser inside another app without detecting it.
		firstLine:
			'Kalau kamu membuka halaman ini dari aplikasi lain, buka halaman ini di Safari dulu.',
		title: 'Pasang Casa di layar utama dulu',
		reason:
			'Di iPhone dan iPad, Casa dibuka dari ikonnya di layar utama. Dengan begitu catatanmu tidak terhapus oleh aturan pembersihan data Safari, dan kamu bisa menerima pemberitahuan.',
		// Numbered text with the iOS button names in Indonesian, no screenshots.
		steps: [
			'1. Ketuk ⋯ di samping alamat situs, lalu ketuk Bagikan.',
			'2. Pilih Tambah ke Layar Utama.',
			'3. Biarkan Buka sebagai Web App menyala, lalu ketuk Tambah.',
			'4. Tutup Safari, lalu buka Casa dari layar utama.'
		],
		shared: {
			title: 'Pakai sebagai perangkat bersama atau pinjaman',
			note: 'Untuk iPhone atau iPad milik orang lain. Catatanmu tidak disimpan di perangkat ini, dan kamu keluar otomatis sesudah 30 menit tidak dipakai.'
		},
		// 08 §7.6 ly90.fromSetup: opened from a setup link, whose token is dropped unsent.
		fromSetup:
			'Kamu membuka tautan pemasangan. Sesudah Casa terpasang, bukalah, pilih Buat akun, lalu tempel tautan pemasangan dari pesan yang kamu terima.'
	},
	// 08 §25.6 LY-91 "Mode perangkat" (SEC-165): no preselected choice.
	ly91: {
		question: 'Ini perangkat pribadimu, atau perangkat bersama atau pinjaman?',
		personal: {
			title: 'Perangkat pribadiku',
			note: 'Salinan catatan tersimpan di browser perangkat ini, jadi kamu bisa mencatat tanpa internet. Siapa pun yang bisa membuka browser ini dapat membaca salinan itu. Bisa memakai pemberitahuan dan kunci aplikasi.'
		},
		shared: {
			title: 'Perangkat bersama atau pinjaman',
			note: 'Catatanmu tidak disimpan di perangkat ini. Kamu keluar otomatis sesudah 30 menit tidak dipakai, dan paling lama 12 jam. Mencatat perlu internet.'
		},
		limit:
			'Casa tidak bisa mencegah perekam layar atau ketikan, atau pengelola kata sandi, yang terpasang di perangkat ini.'
	},
	// 08 §7.6 ly01.*; the two buttons from 08 §5.1 LY-01.
	ly01: {
		title: 'Casa',
		line: 'Uang dan urusan rumah tangga, diketahui bersama tanpa perlu saling bertanya.',
		signUp: 'Buat akun',
		signIn: 'Sudah punya akun'
	},
	// 08 §7.6 signIn.* (LY-02 Masuk); doNotSave is the shared-mode sentence of 08 §25.6 LY-02 (2).
	signIn: {
		title: 'Masuk',
		email: 'Email',
		password: 'Kata sandi',
		rejected: 'Email atau kata sandi belum cocok. Periksa keduanya, lalu coba lagi.',
		doNotSave: 'Jangan simpan kata sandi di browser perangkat ini.',
		// string.max_len on email or password (SEC-12: no email rule); the sentences of
		// signUp.emailInvalid and verify.tooLong, no new ones.
		emailInvalid: 'Tulis email lengkap, misalnya nama@contoh.id.',
		passwordTooLong: 'Kata sandi ini terlalu panjang. Pakai paling banyak 1.024 karakter.',
		// 08 §7.6 signIn.links: the second link; "Saya diundang" comes with its screen (Task 11).
		signUp: 'Belum punya akun? Buat akun'
	},
	// 08 §7.6 signUp.* (LY-02 Buat akun).
	signUp: {
		title: 'Buat akun',
		// No link held: sign-up is closed (SEC-158), the same for anyone, with no request.
		closed:
			'Akun Casa dibuat lewat tautan: tautan undangan dari anggota keluarga, atau tautan pemasangan dari pengelola server.',
		paste: {
			label: 'Tempel tautan',
			hint: 'Salin tautan dari pesan yang kamu terima, lalu tempel di sini.',
			submit: 'Lanjut'
		},
		signIn: 'Sudah punya akun? Masuk',
		notCasa: 'Tautan ini bukan tautan Casa. Salin lagi tautan dari pesan yang kamu terima.',
		withInvitation: 'Kamu membuka undangan Casa. Buat akun dulu, lalu undangannya tampil.',
		withSetup: 'Tautan pemasangan siap dipakai. Pakai email yang kamu berikan ke pengelola server.',
		email: {
			label: 'Email',
			hint: 'Pakai email yang hanya kamu yang membukanya. Kami kirim tautan untuk melanjutkan ke email ini.'
		},
		passwordLater: 'Kata sandi dibuat sesudah kamu membuka tautan itu.',
		submit: 'Kirim tautan',
		// string.email or string.max_len on email.
		emailInvalid: 'Tulis email lengkap, misalnya nama@contoh.id.',
		// TOKEN_REJECTED on a setup link, one sentence for every reason (SEC-30, SEC-158 (b)).
		setupRejected:
			'Tautan pemasangan ini tidak dapat dipakai. Pastikan emailnya sama dengan yang kamu berikan ke pengelola server, atau minta tautan baru kepadanya.'
	},
	// 08 §7.2: an invitation that cannot be used, one sentence for every reason, without the
	// household's name (SEC-30).
	invitation: {
		rejected: 'Undangan ini tidak dapat dipakai. Minta undangan baru kepada yang mengundangmu.'
	},
	// 08 §7.6 checkEmail.* (LY-03); never says an account was made or an email is known (SEC-12).
	checkEmail: {
		title: 'Periksa emailmu',
		sent: (email: string) =>
			`Kami mengirim tautan ke ${email}. Buka tautan itu untuk membuat kata sandi dan melanjutkan.`,
		help: 'Belum masuk dalam beberapa menit? Periksa folder spam, atau kirim ulang.',
		resend: 'Kirim ulang',
		resent: 'Tautan baru sudah dikirim. Pakai tautan di email yang paling baru.',
		change: 'Ganti email'
	},
	// 08 §7.6 verify.* (/signup/verify); no composition rules (SEC-11).
	verify: {
		title: 'Buat kata sandi',
		password: 'Kata sandi',
		hint: 'Paling sedikit 12 karakter. Tidak perlu angka atau simbol; kalimat pendek yang mudah kamu ingat juga boleh.',
		submit: 'Lanjut',
		tooShort: 'Kata sandi paling sedikit 12 karakter.',
		tooLong: 'Kata sandi ini terlalu panjang. Pakai paling banyak 1.024 karakter.',
		common: 'Kata sandi ini terlalu umum dan mudah ditebak. Pilih yang lain.',
		// TOKEN_REJECTED, one sentence for every reason (SEC-30), with the button "Masuk".
		rejected:
			'Tautan ini sudah tidak dapat dipakai. Kalau akunmu sudah jadi, masuk saja. Kalau belum, buka lagi tautan undangan atau tautan pemasangan, lalu daftar lagi.',
		signIn: 'Masuk'
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
