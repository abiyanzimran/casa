# ADR-028 — Lingkungan, irama rilis, dan kompatibilitas klien

| | |
|---|---|
| Status | **Diusulkan** (S5, 12 November 2026). Disajikan kepada pemilik produk bersama ringkasan S5 |
| Tanggal | 12 November 2026 (S5; diperbarui S5 fase D; butir 4 diselaraskan dengan keputusan 42) |
| Pengusul | `architect` |
| Mengubah | ADR-022 butir 1 (**usulan; menunggu jawaban pemilik produk**). ADR-022 butir 1 diputuskan pemilik produk sendiri (17 keputusan 32): build baru paling lambat setiap 90 hari, sasaran 60. Butir 4 di bawah mengusulkan batas keras 80 hari (pertanyaan Q-219 di 17). Sampai pemilik produk menjawab, yang berlaku tetap ADR-022 butir 1 |
| Rincian | [14 §6](../14-engineering-roadmap.md#6-strategi-rilis); [09.4 §1, §5](../09-technical-architecture/04-operasi-uji-biaya.md) |

## Konteks

- Sejak GR-2 (13 §3.3) server pilot memegang data nyata dua orang, salah satunya bukan pengelola server. Sebelum GR-2 hanya ada data uji.
- `ENGINEERING-STANDARDS.md` §9 butir 5 menuntut setiap tugas dijalankan di perangkat pilot nyata atau terhadap API nyata. Sesudah GR-2, menguji build baru di ponsel yang sama dengan data rumah tangga nyata berarti menguji di atas data istri pemilik produk.
- Build TestFlight aktif 90 hari sejak **diunggah**, lalu kedaluwarsa (**[Terverifikasi]** 18 §2.12.6, "All builds uploaded to TestFlight are active for 90 days"). Apakah aplikasi yang sudah terpasang berhenti dapat dibuka sesudahnya **[Perlu validasi]** (18 §2.12.6). ADR-022 butir 1 (keputusan pemilik produk): paling lambat setiap 90 hari, sasaran 60.
- `buf breaking` (ADR-013 butir 4) sudah mencegah perubahan kontrak yang merusak aplikasi lama tanpa keputusan sadar. Di tahap publik, pengguna tidak selalu memperbarui aplikasinya.
- Ponytail (`CLAUDE.md`): tidak ada infrastruktur spekulatif untuk pilot. Satu VPS (brief §1).

## Opsi

### Lingkungan sesudah GR-2

| Kriteria | A. Satu instans; rumah tangga uji di basis data produksi | **B. Instans staging di VPS yang sama, identitas aplikasi terpisah** | C. VPS staging terpisah |
|---|---|---|---|
| Data nyata tersentuh uji | Ya: perintah uji, migrasi, dan pemeriksaan invarian berjalan di basis data yang sama | Tidak | Tidak |
| Biaya uang | 0 | 0 (RAM dan disk VPS yang sama) | Satu VPS lagi per bulan |
| Komponen tambahan | Tidak ada | Satu unit systemd, satu basis data di klaster yang sama, satu subdomain, rahasia sendiri | Satu server penuh |
| Uji migrasi sebelum produksi | Tidak ada tempatnya | Ya, di atas data uji | Ya |
| Gangguan ke produksi saat staging sibuk | — | Mungkin (CPU, RAM bersama) **[Asumsi]** kecil pada skala pilot | Tidak |

### Fitur yang dilepas bertahap (gelombang Tier 2)

| Kriteria | A. Sistem feature flag | **B. Tanpa feature flag; potongan vertikal; titik masuk digabung terakhir** |
|---|---|---|
| Kode tambahan | Penyimpanan flag, evaluasi, layar atau perintah ops | Tidak ada |
| Kesesuaian ES §1 (tanpa kode mati) | Flag mati menyimpan kode yang tidak terjangkau | Kode server terjangkau lewat API dan ujinya; layar baru masuk bersama titik masuknya |
| Pemakaian kedua yang nyata di pilot | Tidak ada (dua pengguna, semua fitur untuk keduanya) | — |

## Keputusan

1. **Lingkungan: opsi B.**
   - **Sebelum GR-2**, instans utama berisi data uji saja dan berfungsi sebagai staging. Pada GR-2 basis datanya dikosongkan dan dibuat ulang lewat migrasi dari nol, lalu data nyata pertama dicatat.
   - **Sejak GR-2**, VPS yang sama menjalankan dua instans: `familyos` (produksi) dan `familyos-staging` (data uji saja). Masing-masing punya basis data sendiri di klaster PostgreSQL yang sama, peran basis data sendiri, rahasia `LoadCredential` sendiri, dan subdomain sendiri. Staging tidak pernah menerima salinan data produksi, dalam bentuk apa pun.
   - Aplikasi staging punya identitas terpisah agar dapat terpasang berdampingan dengan aplikasi pilot: Android lewat `applicationIdSuffix`, iOS lewat bundle ID dan catatan aplikasi App Store Connect kedua yang dipasang dari TestFlight internal **[Usulan]**. Alamat server ditetapkan saat build per identitas, tidak dapat diganti di aplikasi.
   - Staging tidak masuk cadangan pgBackRest (tidak ada data yang perlu dipulihkan) **[Usulan]**.
   - **Syarat pemisahan** (10 indeks §7.4; diuji sebelum GR-2 di 14 QG-30 butir 9, dan lagi oleh SEC-157 sebelum GR-5): staging tidak pernah menerima data produksi dalam bentuk apa pun, termasuk ekspor dan cadangan yang dipulihkan (uji pulih berjalan di VPS cadangan, 09.4 §6); akun uji memakai alamat email domain uji, bukan alamat orang nyata; berjalan sebagai pengguna sistem sendiri; peran basis data staging tidak punya `CONNECT` ke basis data produksi, dan sebaliknya; rahasianya terpisah (SEC-42); konfigurasi keamanannya sama dengan produksi (TLS, batas laju, log dengan daftar izin) dan tanpa titik akhir debug; aplikasi staging hanya dapat menjangkau server staging.
2. **Tanpa sistem feature flag di pilot.** Fitur Tier 2 dibangun sebagai potongan vertikal: server lebih dulu (API hidup, teruji), lalu layar beserta titik masuk navigasinya dalam satu perubahan. Titik masuk sebuah modul Tier 2 digabung hanya sesudah semua kebutuhan Must bersyarat yang menjaganya lulus (05 §6.7; 13 §3.4, misalnya F-AI-005 sebelum `AI` mana pun). **Saklar yang sudah dirancang tetap ada** karena alasan keamanan, bukan pelepasan fitur: saklar mati `AI` (SEC-128) dan kunci konfigurasi R2 (SEC-122).
3. **Saklar operasi "tahan tulis".** Satu subperintah `familyos ops write-freeze on|off` (lewat SSH, tercatat di `security_event`, SEC-33).
   - **Yang ditolak:** setiap perintah tulis yang **bukan** perintah pelindung, dengan kode galat sendiri `WRITE_FROZEN` yang dapat dicoba ulang dan dapat dibedakan aplikasi dari "tidak ada jaringan", agar 08 §13.2 dapat memilih teks status yang benar. Perintah itu tetap di antrean perangkat (LM-02) dan dinilai server saat tiba sesudah saklar dimatikan (06 S-7).
   - **Yang tidak pernah ditolak: perintah pelindung**, yaitu daftar kanonis 10.4 SEC-151 (1) sampai (7): menyempitkan visibilitas (pencabutan berbagi V-7, termasuk status `CARE`; hapus status `CARE` yang dibagikan), menghapus data sendiri (termasuk hapus akun, SEC-105), sembunyikan untuk saya (SEC-153), melaporkan (SEC-154), mencabut sesi (SEC-16, SEC-21), keluar rumah tangga (F06), dan setiap perintah yang memerlukan autentikasi ulang (SEC-13). Selama tahan tulis perintah itu diterima dan berlaku seperti biasa, pada insiden tingkat 1 maupun tingkat 2, dan tidak ada saklar operator untuk menahannya (SEC-151). Ikut tidak ditahan, karena tanpa itu (5) dan (7) tidak dapat dijalankan: masuk, autentikasi ulang, pembaruan token sesi, dan pemulihan lewat email (SEC-13, SEC-15, SEC-24), beserta baris jejak keamanan yang ditulisnya (SEC-39 (b)); keempatnya RPC `identity`, bukan perintah antrean, dan berada di luar pemeriksaan tahan tulis. Daftar perintah tidak disalin: tabel jenis perintah di kontrak memberi setiap jenis tanda `protective` (09 berkas 03 §1). Satu daftar itu melayani kedua aturan: setiap jenis `protective` bernilai `hold_seconds = 0` (jeda urungkan) dan dikecualikan dari tahan tulis. Pesan chat (06 P-41) dan `unmute_member` juga bernilai `hold_seconds = 0` karena rasa pakai, tetapi tidak bertanda `protective` dan mengikuti tahan tulis biasa, demikian pula perluasan berbagi dan urungkan atas penyempitan. Jenis perintah baru yang menulis inti moneter tidak dapat diberi tanda `protective` tanpa penilaian `security-privacy` (SEC-151).
   - **Mengapa aman saat insiden kebenaran data.** (a) Efek perlindungan berlaku seketika (akses menyempit, sesi gugur, keanggotaan berakhir, sembunyi, laporan tercatat), tetapi tidak satu pun perintah pelindung menulis inti moneter selama saklar menyala: F06 adalah Peristiwa tanpa entri (07 CTH-38; INV-41), sisanya menyangkut visibilitas, sesi, atau data sendiri; satu-satunya tulisan inti moneter yang mengikuti perintah pelindung, `exit_closure` (butir (c)), ditunda. Ini diuji: hash inti moneter (INV-09) tidak berubah selama saklar menyala (14 §4.7). (b) Perintah pelindung lewat jalur yang sama dengan biasa: satu transaksi `SERIALIZABLE`, otorisasi saat tiba (SEC-151), dan semua konstrain; tidak ada jalan pintas. (c) `exit_closure` F07 (pada akhir masa tenggang, atau seketika pada hapus akun dan hapus seketika; FR-FAM-012/AC-3) ditunda sampai saklar dimatikan, lalu ditulis **tepat satu kali** dengan tanggal kejadian menurut waktu server saat perintah diterima atau saat masa tenggang berakhir, sehingga angkanya sama dengan tanpa saklar (FR-FAM-012/AC-10; 07 §8.2 butir 2b: buku pribadi dibekukan, saldo klaim tidak berubah). Selama itu buku pribadi tidak terbaca siapa pun. Pembersihan fisik yang menyusul (SEC-94, SEC-96, SEC-100, SEC-105) boleh tertunda selama saklar menyala, tetapi **batas waktunya dihitung dari saat perintah diterima** dan tidak diperpanjang tahan tulis; bila saklar akan menyala melewati batas terdekat, pekerjaan penghapusan yang jatuh tempo dijalankan lebih dulu dan dicatat di catatan insiden (SEC-151). (d) Bila basis data sendiri tidak dapat dipercaya untuk ditulis sama sekali, itu insiden ketersediaan (14 §7.4 tingkat 3: layanan dihentikan dan dipulihkan, 09.4 §6), bukan tahan tulis. Dalam keadaan itu perintah pelindung menunggu di perangkat seperti saat server tidak terjangkau dan dikirim pertama begitu server hidup.
   - **Di perangkat:** penolakan `WRITE_FROZEN` atas perintah lain tidak pernah membuat perintah pelindung menunggu di belakang antrean atau di belakang jeda coba ulang; perintah pelindung dikirim lebih dulu dan seketika, dan aplikasi menampilkan hasil server atasnya, bukan "menunggu" (09 berkas 03 §1; SEC-151). Jumlah catatan yang belum terkirim di layar keluar dihitung dari antrean perangkat (08 §18.3; 07 §8.2 butir 2b).
   - **Dipakai hanya** untuk 13 §6.2 (VL-10 bukan nol) dan untuk insiden keamanan (14 §7.4). Alasan ada sekarang: tanpa saklar ini, menahan data baru bergantung pada pasangan mengingat untuk tidak memakai aplikasi, dan catatan yang dibuat selama itu hilang dari antrean bila mereka tidak mencatat sama sekali **[Usulan]**.
4. **Irama rilis pilot (usulan; mengubah ADR-022 butir 1 bila pemilik produk menyetujuinya).** Setiap gelombang Tier 2 adalah satu rilis untuk setiap klien yang sudah memakai data nyata: iOS dan iPadOS sejak GR-2, Android sejak QG-30 (Android) (keputusan 42; 14 §3.5, §6.3). Bila tidak ada rilis dalam 60 hari, rilis pemeliharaan dibuat. **Usulan:** batas keras **80 hari** sejak build iOS terakhir **diunggah** ke TestFlight (18 §2.12.6), menyisakan 10 hari sebelum kedaluwarsa 90 hari untuk build yang gagal atau tertahan. **Sampai pemilik produk menjawab Q-219**, yang berlaku adalah ADR-022 butir 1: paling lambat 90 hari, sasaran 60, dihitung dari unggahan.
5. **Kompatibilitas klien.**
   - Server melayani setiap build klien yang masih dapat terpasang: build iOS yang belum kedaluwarsa dan dua rilis APK terakhir. `buf breaking` terhadap cabang utama menegakkannya di CI.
   - **Rollback klien adalah rilis perbaikan maju**, bukan pemasangan build lama. Rollback server adalah memasang biner sebelumnya; skema tetap pada tahap *expand* (09.2 §10), dan tahap *contract* digabung paling cepat satu rilis sesudah tahap *expand* terpasang.
   - **Tahap publik** (dirancang sekarang, dibangun di FS-3): kontrak memuat satu kode galat `CLIENT_UPGRADE_REQUIRED` dan server memuat satu konstanta versi build minimum per platform. Dipakai hanya bila perubahan yang merusak tidak terhindarkan, lewat keputusan yang dicatat. Di pilot tidak dibangun: perangkat pilot diperbarui pengelola server sendiri.

## Konsekuensi

- Tidak ada biaya uang tambahan. Biaya waktu: satu unit systemd, satu basis data, satu subdomain, dan satu catatan aplikasi iOS kedua disiapkan sekali di FS-1 langkah 7 (14 §3.3).
- Xcode Cloud membangun dua identitas. Jatah 25 jam per bulan dibagi di antara keduanya (14 §6.3); bila hampir habis, build staging iOS hanya untuk perubahan yang menyentuh `ios/` atau `proto/`.
- Staging dan produksi berbagi CPU dan RAM. Uji beban dan uji kinerja ES §3 dijalankan saat staging diam.
- Saklar "tahan tulis" menambah satu jalur kode yang harus diuji (14 §4.7): perintah yang ditolak saat tahan tulis tidak hilang, dan diterima sesudah saklar dimatikan dengan hasil yang sama dengan tanpa saklar; **setiap jenis perintah pelindung diterima selama saklar menyala** tanpa mengubah hash inti moneter.
- Satu tanda baru di tabel jenis perintah kontrak (`protective`), di samping `hold_seconds`. Uji kontrak: setiap jenis SEC-151 (1) sampai (7) bertanda `protective`, dan `protective` selalu disertai `hold_seconds = 0`.
- Fitur Tier 2 yang belum selesai tidak terlihat oleh pasangan pilot tanpa mekanisme flag, dengan harga: layar baru tidak dapat digabung sebelum servernya dan pagarnya selesai.
