# ADR-017 — Pengenal acak dan kursor sinkron per pembaca

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: UUIDv4 untuk semua pengenal dan umpan sinkron per pembaca (D-99, diterima dengan "lanjut"; [_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk"). Menyelesaikan pertentangan dokumen 10 §8 butir 1 dan 2 menurut rekomendasi `security-privacy` di sana |
| Tanggal | 12 November 2026 (S4 fase C) |
| Pengusul | `architect`, atas persyaratan SEC-03 dan SEC-04 (`security-privacy`) |
| Mengubah | 06 S-11 (a) dan INV-A47 (d): `seq` pesan menjadi internal (permintaan revisi 06, indeks 09 §8) |
| Rincian | [09.2 §4, §5.1](../09-technical-architecture/02-data-dan-skema.md); [09.3 §3, §7](../09-technical-architecture/03-sinkron-waktu-nyata-notifikasi.md) |

## Konteks

06 AT-6 menyimpan jam pencatatan entri tetapi tidak menampilkannya. Pengenal bercap waktu (UUIDv7, ULID, snowflake) membuka jam itu kepada siapa pun yang membaca pengenal (SEC-03).

Kursor sinkron global yang padat itu sederhana. Tetapi bila anggota B menerima posisi 10 lalu 12, ia tahu ada satu perubahan yang tidak boleh ia baca, misalnya entri pribadi pasangannya (SP-40 c; SEC-04). Uji SEC-04 menuntut urutan kursor yang diterima anggota lain **identik byte demi byte**, apakah pasangannya mencatat 0, 1, atau 5 entri pribadi di antaranya.

Chat punya `seq` per percakapan. Satu-satunya celah yang mungkin terbaca adalah kartu status `CARE` yang hilang (INV-A47 (d)), dan dokumen 09 diminta menutupnya.

## Opsi

| Kriteria | **UUIDv4 acak untuk semua** | UUIDv7 internal + UUIDv4 eksternal | Bigint berurutan internal + UUIDv4 eksternal |
|---|---|---|---|
| Waktu tidak terbaca (SEC-03) | Ya | Ya, bila tidak pernah bocor | Ya, bila tidak pernah bocor |
| Jumlah kolom pengenal | Satu | Dua | Dua |
| Lokalitas indeks B-tree | Acak; cukup pada skala yang direncanakan **[Asumsi]** | Baik | Baik |
| Risiko kebocoran tidak sengaja | Tidak ada | Pengenal internal masuk log atau ekspor | Sama |

| Kriteria | Kursor global padat | Kursor global terenkripsi | **Umpan per pembaca** |
|---|---|---|---|
| Celah membocorkan entri pribadi | Ya | Tidak terbaca, tetapi jumlah perubahan per tarik dapat berbeda | **Tidak**: umpan hanya bergerak untuk perubahan yang terbaca |
| Uji byte demi byte SEC-04 | Gagal | Gagal: nonce acak dan isi berbeda | **Lulus**: posisi dan isi bergantung hanya pada yang terbaca |
| Biaya tulis | Satu baris | Satu baris | Satu baris per anggota yang dapat membaca (2 sampai 5) |
| Pembuangan saat akses menyempit (SEC-48) | Perlu mekanisme lain | Sama | Baris `purge` alami |

## Keputusan

1. **Pengenal.**
   - Setiap pengenal baris adalah **UUIDv4** (`gen_random_uuid()`, 122 bit acak), dan satu pengenal dipakai per baris.
   - Tidak ada UUIDv7, ULID, snowflake, atau bigint berurutan sebagai pengenal yang dapat keluar ke klien, log, atau ekspor.
   - Urutan selalu dari kolom eksplisit (`seq`, `event_date`, `server_received_at`), tidak pernah dari pengenal.
2. **Pengenal ditetapkan server.** Objek yang dibuat offline dirujuk lewat `pending_ref` = `idempotency_key` perintah pembuatnya. Klien tidak dapat memilih pengenal.
3. **Kursor sinkron per pembaca.**
   - Tabel `reader_feed(household_id, member_id, pos, object_kind, object_id, change)`; `pos` naik tanpa celah per (rumah tangga, anggota) lewat baris penghitung yang dikunci di transaksi tulis.
   - Setiap perintah menulis satu baris per anggota yang `can_read` objek sesudah perubahan, dan baris `purge` bagi yang kehilangan akses.
   - Kursor dikirim ke klien sebagai string opak (posisi + anggota + MAC server). Kursor anggota lain atau kursor yang diubah ditolak sama dengan "tidak ada".
4. **Isi diambil lewat jalur baca saat ditarik** (SEC-06). Umpan tidak menyimpan isi.
5. **`seq` pesan chat internal.**
   - Klien menerima `order_key` opak yang dapat dibandingkan, dan tanda air dikirim sebagai pengenal pesan.
   - Kartu status yang hilang dikirim sebagai `purge` hanya kepada yang pernah menerimanya.
   - Tidak ada celah apa pun yang terbaca di chat.
6. **Pemangkasan umpan** sesudah 90 hari. Kursor yang lebih tua mendapat `resync_required`.
7. **Potret awal untuk kursor kosong** (ponsel baru, anggota baru, cache dibuang, atau `resync_required`). Server tidak membaca umpan dari awal, karena umpan sudah dipangkas. Ia membangun potret berhalaman dari tabel domain lewat jalur baca, dalam satu transaksi `REPEATABLE READ READ ONLY`: posisi umpan pembaca dicatat di awal, lalu setiap objek hidup yang `can_read` dikembalikan, beserta entri periode anggaran berjalan dan sebelumnya. Kursor sesudah halaman terakhir adalah posisi yang dicatat di awal. Riwayat yang lebih tua dibaca lewat RPC baca biasa saat layar memintanya (09.3 §3). Potret memakai proyeksi yang sama dengan layar, sehingga tidak membuka jalan data kedua (SEC-06), dan urutan halamannya menurut pengenal objek, bukan waktu, sehingga tidak membocorkan jumlah objek tak terbaca (SEC-04).
8. **`household_id` di `SyncRequest` hanya pemilih**, dicocokkan dengan keanggotaan aktif akun sesi; kursor memuat anggota dan rumah tangga di dalam MAC-nya.

## Konsekuensi

- SEC-03 dan SEC-04 dipenuhi oleh struktur, bukan oleh disiplin. Uji SEC-04 menjadi uji otomatis yang memblokir merge (09.4 §4).
- Fan-out saat menulis memerlukan evaluasi `can_read` untuk setiap anggota di setiap perintah. Untuk rumah tangga 2 sampai 5 orang biayanya kecil. Rumah tangga yang jauh lebih besar di tahap publik perlu diukur.
- Baris penghitung per pembaca membuat dua perintah yang mengenai pembaca yang sama saling mengunci sebentar. Pada beban rumah tangga hal itu tidak terasa.
- Indeks B-tree atas UUID acak kurang rapat dibandingkan pengenal berurutan. Pada skala yang direncanakan hal itu diterima. Bila kelak terukur menjadi masalah, pengenal internal berurutan ditambahkan dengan ADR baru dan larangan keluar dijaga uji.
