# ADR-018 — Model sinkron dan waktu nyata

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: SSE + tarik berkursor, offline per modul, jeda urungkan 5 detik (D-100, diterima dengan "lanjut"); kehadiran menurut S4-5 (i) A dan (ii) A, "A tetap seperti WhatsApp" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban pertama) |
| Tanggal | 12 November 2026 (S4 fase C) |
| Pengusul | `architect` |
| Mengubah | 06 §16.5 dan §16.6: **[Asumsi]** platform diganti rancangan; batas diam kehadiran ditetapkan (permintaan revisi 06, indeks 09 §8) |
| Rincian | [09.1 §8](../09-technical-architecture/01-pilihan-teknologi.md); [09.3 §1 sampai §7, §10](../09-technical-architecture/03-sinkron-waktu-nyata-notifikasi.md) |

## Konteks

Brief §9.4 meminta keputusan offline-first atau online-first per modul, beserta sinkron, penyelesaian tabrakan, dan transport waktu nyata.

Dokumen 06 §12 menetapkan aturan domainnya:

- server memutus transisi;
- perintah idempoten dan compare-and-set;
- waktu server;
- tanda dihitung server;
- pencatatan offline tidak pernah ditolak karena state;
- aturan tetap benar bila hanya push disusul tarik yang tersedia.

Chat memerlukan tanda terkirim yang naik hanya saat aplikasi penerima di depan menarik pesan (D-93, W-27), dan kehadiran yang hanya digerakkan pemakaian di depan dengan interaksi (D-95).

Riset 18 §5.6 menemukan:

- iOS dapat menutup soket aplikasi yang ditangguhkan;
- push senyap iOS tidak dijamin;
- FCM berprioritas tinggi diturunkan bila tidak menampilkan notifikasi.

## Opsi

| Kriteria | WebSocket | **SSE + tarik berkursor** | Hanya push + tarik |
|---|---|---|---|
| Kebutuhan dua arah | Tidak ada (tulisan adalah RPC idempoten) | — | — |
| Dependensi server Go | Pustaka pihak ketiga | Tidak ada | APNs, FCM |
| Satu jalan data (SEC-06) | Isi dapat lewat soket | **Tanda saja; isi lewat `Sync`** | Isi lewat `Sync` |
| Latensi chat di depan | Rendah | Rendah (dorongan + tarik) | Detik, tidak dijamin |

## Keputusan

1. **Pola per modul.**
   - **Offline-first:** pencatatan uang, tanggapan permintaan, chat, dan daftar belanja.
   - **Online-first:** pengaturan, kesepakatan, keanggotaan, sesi, `AI`, laporan, dan ekspor.
   - Tabel lengkapnya di 09.3 §1, dengan penyelesaian tabrakan per baris: entri tidak bertabrakan; objek berstate memakai compare-and-set dengan `state_changed` atau `terms_superseded`; tanda air hanya naik; daftar belanja per butir menurut urutan server menerima.
2. **Sinkron.** `Sync(cursor, purpose)` atas umpan per pembaca (ADR-017), yang mengembalikan `UPSERT` dan `PURGE`. Isi dibangun lewat jalur baca. Perintah dikirim dari antrean lokal FIFO per rumah tangga dengan `idempotency_key`, dan diotorisasi ulang saat tiba (SEC-151). Kursor kosong dimulai dengan potret awal (ADR-017 butir 7).
2a. **Jeda urungkan** (dokumen 08 §9.3). Perintah ditulis ke antrean seketika dengan `release_at` = ketukan + `hold_seconds` jenis perintah itu (5 detik bawaan, dari kontrak). "Urungkan" sebelum `release_at` menghapus baris antrean tanpa jejak di server. **Tidak pernah ditahan** (`hold_seconds = 0`): tujuh jenis perintah pelindung dokumen 10 SEC-151, yaitu (1) menyempitkan visibilitas, (2) menghapus data sendiri (pesan, rincian, akun), (3) "Sembunyikan untuk saya", (4) melapor (di pilot tanpa perintah server, SEC-154 (c)), (5) mencabut sesi, (6) keluar rumah tangga, (7) perintah dengan autentikasi ulang; uji kontrak menggagalkan jenis baru dari golongan itu yang nilainya bukan 0. Pesan chat juga 0 (06 P-41, diterima 08 §9.3). Perintah yang ditahan tidak menahan perintah lain yang tidak bergantung padanya. Pencatatan tetap tidak pernah ditolak (06 S-7); kartu dan notifikasi bagi pasangan paling cepat muncul sesudah jeda (09.3 §1).
3. **Waktu nyata.** SSE (`GET /v1/stream`) hanya saat aplikasi di depan, dengan peristiwa `feed` tanpa isi dan `presence` sesaat, serta komentar setiap 15 detik. Klien selalu menarik `Sync` sesudah tersambung. Saat tidak di depan, push umum (ADR-019).
4. **Tanda terkirim.**
   - Hanya `Sync` dengan `purpose = FOREGROUND` yang menaikkan `delivered_upto_seq`.
   - Pilot tidak punya jalur kode latar (tanpa Notification Service Extension, `content-available`, `BGAppRefreshTask`, atau sinkron `WorkManager`). Push Android adalah *notification message*, yang di latar tidak menjalankan kode aplikasi ([dokumentasi FCM](https://firebase.google.com/docs/cloud-messaging/android/receive-messages), diakses 12 November 2026) **[Terverifikasi]**.
   - `FOREGROUND` hanya dikirim dari jalur siklus hidup layar aktif. Sinkron latar kelak memakai klien terpisah yang hanya dapat mengirim `BACKGROUND`.
5. **Kehadiran.**
   - Denyut paling sering setiap 15 detik selama ada interaksi di layar yang memenuhi syarat.
   - Online berakhir 30 detik sesudah denyut terakhir atau seketika saat aplikasi ke latar. Mengetik tampil sampai 6 detik.
   - Disimpan di memori proses saja. Terakhir aktif satu baris yang ditimpa, dipotong ke **menit** sebelum disimpan (SEC-60; S4-5 (ii) A).
   - Denyut dikirim dari **setiap layar aplikasi** saat di depan dengan interaksi (S4-5 (i) A); tidak ada setelan `presence_hidden`. Disebarkan hanya ke layar chat, tidak lebih rinci daripada yang tampil di layar (UX-5).
6. **Tanpa message broker.** Satu transaksi menulis domain, `object_event`, `reader_feed`, `notification`, dan `job`. `NOTIFY` sesudah commit membangunkan hub dan pekerja.

## Konsekuensi

- Satu jalan data: semua isi lewat `Sync` dan jalur baca. SSE dan push hanya mempercepat. Kehilangan satu peristiwa SSE atau push tidak menghilangkan data.
- Anggapan lama 06 §16.5 bahwa platform membedakan tarik di depan dari tugas latar tidak lagi diperlukan, karena pembedaannya dibuat oleh struktur kode kita (06 revisi ketiga belas sudah mengganti **[Asumsi]** itu). Harganya: pilot tidak punya sinkron latar, sehingga isi baru selalu ditarik saat aplikasi dibuka. Untuk dua pengguna dengan volume kecil, tarikan itu cepat.
- Server tidak dapat membuktikan klien sungguh di depan. Klien yang dimodifikasi hanya dapat memalsukan tanda miliknya sendiri.
- Hub SSE dan kehadiran di memori proses hanya berlaku untuk satu proses. Tahap publik dengan beberapa proses memakai `LISTEN/NOTIFY` (ADR-014 butir 5).
