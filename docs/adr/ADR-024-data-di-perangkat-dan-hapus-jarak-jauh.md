# ADR-024 — Data di perangkat: proteksi, cadangan, hapus jarak jauh, dan `CARE` lokal

| | |
|---|---|
| Status | **Diusulkan** (`architect` bersama `security-privacy`, 12 November 2026); akan diputuskan pemilik produk di checkpoint S4. Menyelesaikan pertentangan dokumen 10 §8 butir 3, 4, dan 9 menurut rekomendasi `security-privacy` di sana. Butir 5 berlaku hanya bila pemilik produk memilih 10-D A |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D) |
| Pengusul | `architect`, atas SEC-21, SEC-46 sampai SEC-48 (`security-privacy`) |
| Mengubah | 06 §13.3, §13.5 bila 10-D A (06 revisi ketiga belas) |
| Rincian | [09.3 §2, §3](../09-technical-architecture/03-sinkron-waktu-nyata-notifikasi.md) |

## Konteks

Dokumen 10 menuntut:

- **SEC-46:** basis data perangkat dilindungi proteksi berkas OS pada kelas yang masih memungkinkan sinkron latar, dan dikecualikan dari cadangan perangkat dan awan;
- **SEC-47:** perangkat hanya menyimpan yang terbaca;
- **SEC-48:** salinan dibuang saat akses menyempit;
- **SEC-21:** perangkat yang sesinya dicabut menghapus data dan antreannya sebelum menampilkan apa pun.

Dokumen 10 §8 mencatat tiga pertentangan:

- **butir 3:** status `CARE` yang tidak dibagikan disimpan di perangkat saja (10-D A), sehingga alur sebelum berbagi pindah ke perangkat;
- **butir 4:** hapus jarak jauh membuang catatan yang belum terkirim, sedangkan ES mencegah kehilangan data;
- **butir 9:** pengecualian dari cadangan berarti pindah ponsel tidak membawa catatan yang belum terkirim.

Perilaku platform yang diperiksa 12 November 2026:

- **Apple:** Class C (`NSFileProtectionCompleteUntilFirstUserAuthentication`) menjaga kunci kelas di memori sesudah perangkat dikunci, dan merupakan kelas bawaan data aplikasi pihak ketiga ([Apple Platform Security](https://support.apple.com/guide/security/data-protection-classes-secb010e978a/web)) **[Terverifikasi]**.
- **Apple** menyarankan kelas proteksi yang lebih ketat daripada Class C bila aplikasi tidak memerlukan akses saat terkunci (dokumen 10, SEC-46). Pilot tidak punya jalur kode latar (09.3 §8), sehingga memakai Class C adalah **penyimpangan dari saran itu** dan harus beralasan.
- **Android:** `allowBackup="false"` pada sebagian produsen tidak menghentikan transfer antarperangkat. Data tetap di perangkat hanya bila dikecualikan di `<cloud-backup>`, `<device-transfer>`, **dan `<cross-platform-transfer>`** dalam `dataExtractionRules` ([developer.android.com](https://developer.android.com/identity/data/autobackup); 18 §5 fase D) **[Terverifikasi]**.

## Opsi

| Butir | Opsi A | **Opsi yang dipilih** |
|---|---|---|
| Kelas proteksi iOS (basis data, antrean, foto) | Class A (Complete): data tidak terbaca saat terkunci; tulisan yang masih berjalan saat perangkat dikunci gagal | **Class C**: kelas bawaan; tulisan SQLite dan WAL yang selesai sesudah perangkat dikunci tidak gagal |
| Kelas proteksi iOS (berkas `CARE` lokal, 10-D A) | Class C | **Class A**: tidak pernah disinkron, sehingga tidak ada tulisan latar |
| Cadangan Android | `allowBackup="false"` saja | **Aturan eksklusi eksplisit untuk cadangan awan, transfer antarperangkat, dan transfer lintas platform**, ditambah `fullBackupContent` untuk Android 11 ke bawah |
| Hapus jarak jauh | Kirim antrean dulu, baru hapus | **Hapus sebelum layar pertama; antrean yang belum terkirim ikut terhapus** |
| `CARE` tak dibagikan | Di server (06 §13.5 sekarang) | **Di perangkat saja bila 10-D A** |

## Keputusan

1. **iOS dan iPadOS:**
   - berkas basis data, WAL, antrean, dan cache foto memakai **Class C** yang ditetapkan eksplisit. **Penyimpangan dari saran Apple, beserta alasannya:** antrean perintah dan jawaban `Sync` ditulis ke SQLite; tulisan yang masih berjalan saat pengguna mengunci perangkat (misalnya menghapus perintah dari antrean sesudah server menerimanya) gagal di Class A, dan antrean dapat tertinggal setengah tertulis, sehingga perintah terkirim dua kali dari sisi pengguna atau hilang dari tampilan **[Asumsi]**, dibuktikan `ios-dev` di perangkat. Idempotensi server mencegah efek ganda, tetapi tidak mencegah antrean rusak. Bila uji di perangkat menunjukkan Class B (`CompleteUnlessOpen`) atau Class A aman untuk pola tulis ini, kelas yang lebih ketat dipakai;
   - berkas `CARE` lokal (butir 5) memakai **Class A, `NSFileProtectionComplete`**, dalam basis data terpisah;
   - direktori itu diberi `isExcludedFromBackup = true` **[Perlu validasi]** perilakunya di perangkat;
   - token disimpan di Keychain dengan kelas "sesudah pertama dibuka, perangkat ini saja".
2. **Android:**
   - basis data di penyimpanan internal (*credential encrypted*);
   - `dataExtractionRules` meng-`exclude` domain `database`, `file`, dan `sharedpref` pada `<cloud-backup>`, `<device-transfer>`, dan `<cross-platform-transfer>`, ditambah `fullBackupContent` setara; uji di Galaxy A55 membuktikan bahwa pindah perangkat tidak membawa basis data;
   - token dibungkus kunci Android Keystore yang tidak dapat diekspor.
3. **Tanpa SQLCipher di pilot.** Dinilai ulang sebelum tahap publik (SEC-46).
4. **Hapus jarak jauh.**
   - Setiap jawaban server atas sesi yang dicabut adalah galat `session_revoked`.
   - Aplikasi menghapus basis data, berkas, antrean, dan notifikasi yang tampil **sebelum** menampilkan layar apa pun.
   - Perintah lama dari perangkat itu ditolak server.
   - Layar cabut di perangkat lain menyatakan bahwa catatan yang belum terkirim di perangkat itu akan hilang (SEC-21). Alasannya: perangkat yang hilang sudah kehilangan catatan itu bagi pemiliknya, dan mengirimnya tanpa pemegang yang sah lebih berbahaya.
5. **`CARE` tak dibagikan, bila 10-D A:**
   - status yang tidak dibagikan disimpan hanya di basis data perangkat pemiliknya;
   - perintah berbagi mengunggah status saat itu kepada penerima yang dipilih;
   - mencabut berbagi menghapusnya dari server dan menghasilkan `purge` bagi penerima.
   - Status tak dibagikan hilang bila perangkat hilang, dan itu dinyatakan di layar `CARE`. Ia Tier 2, sehingga tidak menyentuh bangun Tier 1.
6. **Pindah ponsel** (dokumen 10 §8 butir 9): catatan yang sudah terkirim disinkron ulang dari server. Yang belum terkirim tidak ikut. Sebelum keluar akun, aplikasi menampilkan jumlah catatan yang belum terkirim (SEC-47).

## Konsekuensi

- Penemu perangkat yang terkunci sesudah dinyalakan ulang tidak dapat membaca data. Sesudah pembukaan pertama, perlindungannya bergantung pada kunci perangkat dan kunci aplikasi (SEC-19). Status `CARE` lokal tidak terbaca setiap kali perangkat terkunci.
- Tidak ada salinan data rumah tangga di iCloud atau Google Drive (SEC-145).
- Dua jalan yang sengaja menghilangkan data yang belum terkirim (hapus jarak jauh dan pindah ponsel) dinyatakan di layar. Keduanya diterima `security-privacy` di dokumen 10 §8.
- Bila 10-D A, sebagian alur `CARE` (I01 sampai I04 sebelum berbagi) berjalan di perangkat. Itu bukan aturan uang atau persetujuan, sehingga tidak bertentangan dengan ES §2. Yang dibagikan tetap diputus server.
