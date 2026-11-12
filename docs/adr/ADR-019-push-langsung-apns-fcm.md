# ADR-019 — Push langsung ke APNs dan FCM dengan payload umum

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **langsung ke APNs dan FCM** dengan payload umum (D-101, diterima dengan "lanjut"; [_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk") |
| Tanggal | 12 November 2026 (S4 fase C; status checkpoint S4; catatan web lebih dulu sesudah S6) |
| Pengusul | `architect`, atas SEC-49 sampai SEC-58 (`security-privacy`) |
| Mengubah | — |
| Rincian | [09.1 §9](../09-technical-architecture/01-pilihan-teknologi.md); [09.3 §8](../09-technical-architecture/03-sinkron-waktu-nyata-notifikasi.md) |

## Konteks

Aturan notifikasi berasal dari dua sumber:

- **06 §11:** tiga kelas prioritas, jam tenang per orang, ringkasan pukul 19.00, dan katalog N-01 sampai N-32.
- **Dokumen 10:** payload hanya kelas dan satu pengenal opak (SEC-50); teks layar kunci per kelas (SEC-51); kanal per kelas (SEC-52); lencana hanya `perlu_tanggapan` (SEC-53); satu notifikasi sesudah jam tenang (SEC-55); lipatan N-28 (SEC-57); token dihapus (SEC-58); tanpa SDK pihak ketiga selain yang perlu, dan lalu lintas hanya ke server, APNs, dan FCM (SEC-49).

Tanda terkirim tidak boleh digerakkan push (ADR-018).

## Opsi

| Kriteria | **Langsung APNs + FCM** | FCM untuk kedua platform | Agregator |
|---|---|---|---|
| SDK di iOS | Tidak ada | Firebase | SDK agregator |
| Pihak yang melihat metadata | Apple, Google | Google juga untuk iOS | Satu pihak lagi |
| Kendali collapse, kanal, lencana | Penuh | Lewat pemetaan | Bergantung penyedia |

## Keputusan

1. **Server mengirim langsung.**
   - **APNs:** HTTP/2 dengan token JWT ES256 dari kunci `.p8`, diperbarui tiap 20 sampai 60 menit.
   - **FCM HTTP v1:** OAuth2 akun layanan.
   - Pengiriman berjalan sebagai pekerjaan `push.send`.
2. **Payload:**
   - teks kelas dalam bahasa penerima, misalnya "Ada yang menunggu jawabanmu" (SEC-51);
   - satu medan `r` = `push_ref`, UUIDv4 per notifikasi yang tidak bermakna tanpa sesi.
   - **Tidak ada** nama, jumlah, judul, jenis peristiwa, status, atau pengenal objek.
3. **Tanpa kode latar.**
   - iOS tanpa `mutable-content` dan tanpa `content-available`, sehingga tidak ada Notification Service Extension.
   - Android memakai *notification message* dengan `channel_id` per kelas.
4. **Kanal dan kategori per kelas:** "Perlu tanggapan", "Kabar", "Ringkasan". Thread per kelas.
5. **Lencana:** jumlah `perlu_tanggapan` penerima, dihitung server (`aps.badge`; `notification_count`).
6. **Lipatan N-28:** collapse id (APNs) atau tag (FCM) per (pemilik status, penerima), dengan bunyi paling banyak sekali per 60 menit.
7. **Jam tenang dan ringkasan** dijalankan server: `hold_until` menurut zona penerima, satu pelepasan, dan ringkasan pukul 19.00. Butir yang gugur hilang tanpa jejak.
8. **SDK:** iOS memakai UserNotifications sistem. Android memakai `firebase-messaging` saja, dengan pengumpulan analitik Firebase dimatikan.
9. **Token** disimpan per sesi dan dihapus saat keluar akun, sesi dicabut, atau F06. Firebase installation ID dihapus lewat API-nya saat keluar akun.
10. **Penarikan kembali:** `PURGE` membawa `retract_push_refs`, dan aplikasi menghapus notifikasi yang masih tampil saat dibuka (SEC-48).

## Konsekuensi

- Apple dan Google tetap melihat token perangkat dan waktu kirim (TM-18; L-16). Isinya hanya teks kelas.
- Tanpa NSE, notifikasi iOS tidak dapat menampilkan rincian apa pun di layar kunci. Itu memang dikehendaki SEC-51.
- Pemeliharaan dua klien push di server (sekitar dua ratus baris **[Asumsi]**) lebih murah daripada satu dependensi agregator.
- Bila kelak sinkron latar diperlukan (tahap publik), ADR baru menilai `content-available` bersama aturan tanda terkirim.

## Catatan keputusan web lebih dulu (12 November 2026, sesudah S6)

Mencatat [ADR-030](ADR-030-klien-web-pwa-lebih-dulu.md) butir 14. **Web Push menjadi kanal ketiga**, dan di jalur GR-2 satu-satunya kanal yang dipakai; APNs dan FCM baru dipakai saat aplikasi native tersedia.

- **Butir 1:** server mengirim langsung ke endpoint layanan push browser dengan VAPID (RFC 8292) dan payload terenkripsi (RFC 8291), sebagai pekerjaan `push.send` yang sama. Pustaka `github.com/SherClockHolmes/webpush-go` v1.4.0 (perawatan **[Perlu validasi]**); jalan mundur di atas stdlib. JWT VAPID ES256, `exp` paling lama 24 jam, diperbarui paling sering sekali per jam, `sub` alamat operasi umum (18 §5.13.3; SEC-177).
- **Butir 2:** payload sama dengan SEC-50 ditambah angka lencana, di bawah 4 KB, **selalu ditampilkan** (Safari mencabut izin bila notifikasi tidak tampil). `tag` (lipatan, butir 6) dan header `Topic` opak, paling banyak 32 karakter.
- **Butir 3:** handler `push` di service worker hanya `showNotification` dan lencana, tanpa panggilan jaringan (ADR-018 catatan web).
- **Butir 5:** lencana lewat Badging API di web app layar utama iOS dan iPadOS dan desktop; tidak didukung di Android (18 §5.13.3), jadi di Galaxy A55 hanya notifikasi.
- **Butir 9:** langganan disimpan per sesi di `push_token` (`platform = 'webpush'`, endpoint, `p256dh`, `auth`) dan dihapus pada jalur yang sama dengan token native (SEC-58). Endpoint diperiksa terhadap daftar izin host tanpa redirect saat berlangganan dan saat mengirim (SEC-176 (g); TM-43; daftarnya di ADR-030 butir 14).
- **Butir 10:** penarikan kembali lewat `getNotifications({tag})` dan `close()` saat aplikasi dibuka.
- **Konsekuensi:** penyedia layanan push browser (Apple, Google, Microsoft, Mozilla) melihat endpoint dan waktu kirim, sama seperti APNs dan FCM (TM-18). Tidak ada biaya dan tidak memerlukan keanggotaan Apple Developer untuk web (18 §5.13.3).
