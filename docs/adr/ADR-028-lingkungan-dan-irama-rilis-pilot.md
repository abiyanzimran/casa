# ADR-028 — Lingkungan, irama rilis, dan kompatibilitas klien

| | |
|---|---|
| Status | **Diterima seluruhnya.** **Butir 1, 2, 3, dan 5** diterima atas saran lead lewat **keputusan 56** (12 November 2026, selama S6; jawaban pemilik produk "Terima semua (Saran)", [_progress.md](../_progress.md), "Keputusan pemilik produk 56 sampai 58"). **Butir 4** diterima atas saran lead lewat keputusan 44 (12 November 2026; jawabannya "Ya, 60/80 hari (Saran)", [_progress.md](../_progress.md), "Penutupan pertanyaan terbuka sebelum S6"). Semula Diusulkan (S5). Butir 1 diselaraskan dengan keputusan 45 (server uji sementara sampai menjelang GR-2) |
| Tanggal | 12 November 2026 (S5; diperbarui S5 fase D; butir 4 diselaraskan dengan keputusan 42; keputusan 44 dan 45; S6: keputusan 56; catatan web lebih dulu sesudah S6); 13 November 2026 (FS-1 1.3: sumber versi minimum `web` diselaraskan dengan ADR-030 butir 9) |
| Pengusul | `architect` |
| Mengubah | ADR-022 butir 1 (**diterima**, keputusan 44): batas keras 80 hari sejak build iOS terakhir diunggah ke TestFlight, sasaran tetap 60. Semula ADR-022 butir 1, keputusan pemilik produk sendiri (17 keputusan 32), berbunyi paling lambat setiap 90 hari, sasaran 60 |
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
| Komponen tambahan | Tidak ada | Satu unit systemd, satu klaster PostgreSQL kedua (semula "satu basis data di klaster yang sama"; catatan 13 November 2026 di butir 1), satu subdomain, rahasia sendiri | Satu server penuh |
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
   - **Sebelum GR-2**, instans utama berisi data uji saja dan berfungsi sebagai staging. Menurut keputusan 45 (14 §7.8), FS-1 langkah 1 sampai 6 memakai **server uji sementara** berisi data uji saja, yang dimatikan dan dihapus sesudah instans di VPS pilihan S4-2 hidup; sejak langkah 7 instans utama berjalan di VPS pilihan S4-2, masih dengan data uji saja. Pada GR-2 basis datanya dikosongkan dan dibuat ulang lewat migrasi dari nol, lalu data nyata pertama dicatat.
   - **Sejak GR-2**, VPS yang sama menjalankan dua instans: `casa` (produksi) dan `casa-staging` (data uji saja). Masing-masing punya **klaster PostgreSQL sendiri** (catatan 13 November 2026 di bawah butir ini; semula "basis data sendiri di klaster PostgreSQL yang sama"), peran basis data sendiri, rahasia `LoadCredential` sendiri, dan subdomain sendiri. Staging tidak pernah menerima salinan data produksi, dalam bentuk apa pun.
   - Aplikasi staging punya identitas terpisah agar dapat terpasang berdampingan dengan aplikasi pilot: Android lewat `applicationIdSuffix`, iOS lewat bundle ID dan catatan aplikasi App Store Connect kedua yang dipasang dari TestFlight internal. Alamat server ditetapkan saat build per identitas, tidak dapat diganti di aplikasi.
   - Staging tidak masuk cadangan pgBackRest (tidak ada data yang perlu dipulihkan).
   - **Syarat pemisahan** (10 indeks §7.4; diuji sebelum GR-2 di 14 QG-30 butir 9, dan lagi oleh SEC-157 sebelum GR-5, bila jalur publik diambil): staging tidak pernah menerima data produksi dalam bentuk apa pun, termasuk ekspor dan cadangan yang dipulihkan (uji pulih berjalan di VPS cadangan, 09.4 §6); akun uji memakai alamat email domain uji, bukan alamat orang nyata; berjalan sebagai pengguna sistem sendiri; peran basis data staging tidak punya `CONNECT` ke basis data produksi, dan sebaliknya; rahasianya terpisah (SEC-42); konfigurasi keamanannya sama dengan produksi (TLS, batas laju, log dengan daftar izin) dan tanpa titik akhir debug; aplikasi staging hanya dapat menjangkau server staging.
   - **Catatan 13 November 2026 (FS-1 1.2, temuan `code-reviewer` I4): satu klaster per lingkungan.** Peran PostgreSQL berlaku untuk seluruh klaster, dan nama peran (`migrator`, `app_rw`, `job_runner`, `lifecycle_rw`) tetap (09.2 §3 lapis 5). Bila staging dan produksi berbagi klaster, `app_rw` staging adalah `app_rw` produksi, dan syarat "tanpa `CONNECT` ke basis data produksi" tidak dapat dipenuhi. Opsi yang ditimbang:

     | Opsi | Perubahan kode | Pemisahan peran | Cadangan |
     |---|---|---|---|
     | **Klaster kedua di VPS yang sama** (`pg_createcluster 18 staging`, port berikutnya) | Tidak ada | Penuh: peran staging tidak ada di klaster produksi | pgBackRest hanya untuk klaster produksi; staging benar-benar tidak masuk cadangan |
     | Awalan nama peran per lingkungan | `roles.sql` dan setiap `GRANT`/kebijakan RLS di migrasi diparameterkan | Penuh bila tidak ada yang lupa | Staging tetap ikut cadangan fisik produksi |
     | `CONNECT` per basis data plus set peran berbeda | Sama dengan awalan | Sama | Sama |

     **Diputuskan: klaster kedua.** Alasan: tanpa perubahan kode, dan pgBackRest selalu menyalin isi seluruh klaster ("pgBackRest copies the entire contents of the database cluster to the backup", [pgbackrest.org/user-guide.html](https://pgbackrest.org/user-guide.html), diakses 13 November 2026) **[Terverifikasi]**, sehingga di klaster bersama data uji staging ikut cadangan produksi. `pg_createcluster` dari postgresql-common membuat klaster baru per versi dan nama, dan memberinya port bebas berikutnya mulai 5432 ([manpages.debian.org pg_createcluster(1)](https://manpages.debian.org/trixie/postgresql-common/pg_createcluster.1.en.html), diakses 13 November 2026) **[Terverifikasi]**. `roles.sql` dijalankan superuser masing-masing klaster. Biaya: satu proses `postgres` lagi; `shared_buffers` staging disetel kecil **[Asumsi]** ukurannya diukur di langkah 7. Uji QG-30 butir 9 tetap: kredensial staging ditolak oleh klaster produksi, dan sebaliknya. Disiapkan di FS-1 langkah 7; kode tugas 1.2 tidak berubah.
2. **Tanpa sistem feature flag di pilot.** Fitur Tier 2 dibangun sebagai potongan vertikal: server lebih dulu (API hidup, teruji), lalu layar beserta titik masuk navigasinya dalam satu perubahan. Titik masuk sebuah modul Tier 2 digabung hanya sesudah semua kebutuhan Must bersyarat yang menjaganya lulus (05 §6.7; 13 §3.4, misalnya F-AI-005 sebelum `AI` mana pun). **Saklar yang sudah dirancang tetap ada** karena alasan keamanan, bukan pelepasan fitur: saklar mati `AI` (SEC-128) dan kunci konfigurasi R2 (SEC-122).
3. **Saklar operasi "tahan tulis".** Satu subperintah `casa ops write-freeze on|off` (lewat SSH, tercatat di `security_event`, SEC-33).
   - **Yang ditolak:** setiap perintah tulis yang **bukan** perintah pelindung, dengan kode galat sendiri `WRITE_FROZEN` yang dapat dicoba ulang dan dapat dibedakan aplikasi dari "tidak ada jaringan", agar 08 §13.2 dapat memilih teks status yang benar. Perintah itu tetap di antrean perangkat (LM-02) dan dinilai server saat tiba sesudah saklar dimatikan (06 S-7). **Tahan per akun `NOTICE_UNREAD`** (SEC-34; 09 berkas 03 §1) adalah mekanisme terpisah yang tidak mengubah butir ini. Server memeriksa `WRITE_FROZEN` lebih dulu, lalu `NOTICE_UNREAD`. Keduanya tidak pernah menahan jenis `protective`.
   - **Yang tidak pernah ditolak: perintah pelindung**, yaitu daftar kanonis 10.4 SEC-151 (1) sampai (7): menyempitkan visibilitas (pencabutan berbagi V-7, termasuk status `CARE`; hapus status `CARE` yang dibagikan), menghapus data sendiri (termasuk hapus akun, SEC-105), sembunyikan untuk saya (SEC-153), melaporkan (SEC-154), mencabut sesi (SEC-16, SEC-21), keluar rumah tangga (F06), dan setiap perintah yang memerlukan autentikasi ulang (SEC-13). Selama tahan tulis perintah itu diterima dan berlaku seperti biasa, pada insiden tingkat 1 maupun tingkat 2, dan tidak ada saklar operator untuk menahannya (SEC-151). Ikut tidak ditahan, karena tanpa itu (5) dan (7) tidak dapat dijalankan: masuk, autentikasi ulang, pembaruan token sesi, dan pemulihan lewat email (SEC-13, SEC-15, SEC-24), beserta baris jejak keamanan yang ditulisnya (SEC-39 (b)); keempatnya RPC `identity`, bukan perintah antrean, dan berada di luar pemeriksaan tahan tulis. Daftar perintah tidak disalin: tabel jenis perintah di kontrak memberi setiap jenis tanda `protective` (09 berkas 03 §1). Satu daftar itu melayani kedua aturan: setiap jenis `protective` bernilai `hold_seconds = 0` (jeda urungkan) dan dikecualikan dari tahan tulis. Pesan chat (06 P-41) dan `unmute_member` juga bernilai `hold_seconds = 0` karena rasa pakai, tetapi tidak bertanda `protective` dan mengikuti tahan tulis biasa, demikian pula perluasan berbagi dan urungkan atas penyempitan. Jenis perintah baru yang menulis inti moneter tidak dapat diberi tanda `protective` tanpa penilaian `security-privacy` (SEC-151).
   - **Mengapa aman saat insiden kebenaran data.** (a) Efek perlindungan berlaku seketika (akses menyempit, sesi gugur, keanggotaan berakhir, sembunyi, laporan tercatat), tetapi tidak satu pun perintah pelindung menulis inti moneter selama saklar menyala: F06 adalah Peristiwa tanpa entri (07 CTH-38; INV-41), sisanya menyangkut visibilitas, sesi, atau data sendiri; satu-satunya tulisan inti moneter yang mengikuti perintah pelindung, `exit_closure` (butir (c)), ditunda. Ini diuji: hash inti moneter (INV-09) tidak berubah selama saklar menyala (14 §4.7). (b) Perintah pelindung lewat jalur yang sama dengan biasa: satu transaksi `SERIALIZABLE`, otorisasi saat tiba (SEC-151), dan semua konstrain; tidak ada jalan pintas. (c) `exit_closure` F07 (pada akhir masa tenggang, atau seketika pada hapus akun dan hapus seketika; FR-FAM-012/AC-3) ditunda sampai saklar dimatikan, lalu ditulis **tepat satu kali** dengan tanggal kejadian menurut waktu server saat perintah diterima atau saat masa tenggang berakhir, sehingga angkanya sama dengan tanpa saklar (FR-FAM-012/AC-10; 07 §8.2 butir 2b: buku pribadi dibekukan, saldo klaim tidak berubah). Selama itu buku pribadi tidak terbaca siapa pun. Pembersihan fisik yang menyusul (SEC-94, SEC-96, SEC-100, SEC-105) boleh tertunda selama saklar menyala, tetapi **batas waktunya dihitung dari saat perintah diterima** dan tidak diperpanjang tahan tulis; bila saklar akan menyala melewati batas terdekat, pekerjaan penghapusan yang jatuh tempo dijalankan lebih dulu dan dicatat di catatan insiden (SEC-151). (d) Bila basis data sendiri tidak dapat dipercaya untuk ditulis sama sekali, itu insiden ketersediaan (14 §7.4 tingkat 3: layanan dihentikan dan dipulihkan, 09.4 §6), bukan tahan tulis. Dalam keadaan itu perintah pelindung menunggu di perangkat seperti saat server tidak terjangkau dan dikirim pertama begitu server hidup.
   - **Di perangkat:** penolakan `WRITE_FROZEN` atas perintah lain tidak pernah membuat perintah pelindung menunggu di belakang antrean atau di belakang jeda coba ulang; perintah pelindung dikirim lebih dulu dan seketika, dan aplikasi menampilkan hasil server atasnya, bukan "menunggu" (09 berkas 03 §1; SEC-151). Jumlah catatan yang belum terkirim di layar keluar dihitung dari antrean perangkat (08 §18.3; 07 §8.2 butir 2b).
   - **Dipakai hanya** untuk 13 §6.2 (VL-10 bukan nol) dan untuk insiden keamanan (14 §7.4). Alasan ada sekarang: tanpa saklar ini, menahan data baru bergantung pada pasangan mengingat untuk tidak memakai aplikasi, dan catatan yang dibuat selama itu hilang dari antrean bila mereka tidak mencatat sama sekali (diterima, keputusan 56).
   - **Hubungan dengan aturan tanpa-kunci** (S6; 06 INV-A11, PG-10; ADR-009): aturan, tanda, kesepakatan, dan state tidak pernah menahan pencatatan. Satu-satunya penahanan adalah penahanan operasional ini dan `NOTICE_UNREAD` (SEC-34); keduanya membiarkan perintah di antrean perangkat (LM-02) dan tidak pernah menahan perintah pelindung (SEC-151). Bila akhir masa tenggang keluar jatuh selagi tahan tulis menyala, `exit_closure` dan penutupan buku pribadi berjalan dalam satu transaksi begitu tahan tulis dimatikan, tepat satu kali, dengan tanggal kejadian saat masa tenggang berakhir (butir (c) di atas; 06 F07 dan 07 INV-28 memuat kalimat yang sama).
4. **Irama rilis pilot (diterima, keputusan 44; mengubah ADR-022 butir 1).** Setiap gelombang Tier 2 adalah satu rilis untuk setiap klien yang sudah memakai data nyata: iOS dan iPadOS sejak GR-2, Android sejak QG-30 (Android) (keputusan 42; 14 §3.5, §6.3). Bila tidak ada rilis dalam 60 hari, rilis pemeliharaan dibuat. **Batas keras 80 hari** sejak build iOS terakhir **diunggah** ke TestFlight (18 §2.12.6), menyisakan 10 hari sebelum kedaluwarsa 90 hari untuk build yang gagal atau tertahan.
5. **Kompatibilitas klien.**
   - Server melayani setiap build klien yang masih dapat terpasang: build iOS yang belum kedaluwarsa dan dua rilis APK terakhir. `buf breaking` terhadap cabang utama menegakkannya di CI.
   - **Rollback klien adalah rilis perbaikan maju**, bukan pemasangan build lama. Rollback server adalah memasang biner sebelumnya; skema tetap pada tahap *expand* (09.2 §10), dan tahap *contract* digabung paling cepat satu rilis sesudah tahap *expand* terpasang.
   - **Tahap publik** (dirancang sekarang, dibangun di FS-3 bila jalur publik opsional kelak diambil; keputusan 51): kontrak memuat satu kode galat `CLIENT_UPGRADE_REQUIRED` dan server memuat satu konstanta versi build minimum per platform. Dipakai hanya bila perubahan yang merusak tidak terhindarkan, lewat keputusan yang dicatat. Di pilot tidak dibangun: perangkat pilot diperbarui pengelola server sendiri.

## Konsekuensi

- Tidak ada biaya uang tambahan. Biaya waktu: satu unit systemd, satu klaster PostgreSQL kedua (catatan 13 November 2026 di butir 1), satu subdomain, dan satu catatan aplikasi iOS kedua disiapkan sekali di FS-1 langkah 7 (14 §3.3).
- Xcode Cloud membangun dua identitas. Jatah 25 jam per bulan dibagi di antara keduanya (14 §6.3); bila hampir habis, build staging iOS hanya untuk perubahan yang menyentuh `ios/` atau `proto/`.
- Staging dan produksi berbagi CPU dan RAM. Uji beban dan uji kinerja ES §3 dijalankan saat staging diam.
- Saklar "tahan tulis" menambah satu jalur kode yang harus diuji (14 §4.7): perintah yang ditolak saat tahan tulis tidak hilang, dan diterima sesudah saklar dimatikan dengan hasil yang sama dengan tanpa saklar; **setiap jenis perintah pelindung diterima selama saklar menyala** tanpa mengubah hash inti moneter.
- Satu tanda baru di tabel jenis perintah kontrak (`protective`), di samping `hold_seconds`. Uji kontrak: setiap jenis SEC-151 (1) sampai (7) bertanda `protective`, dan `protective` selalu disertai `hold_seconds = 0`.
- Fitur Tier 2 yang belum selesai tidak terlihat oleh pasangan pilot tanpa mekanisme flag, dengan harga: layar baru tidak dapat digabung sebelum servernya dan pagarnya selesai.

## Catatan keputusan web lebih dulu (12 November 2026, sesudah S6)

Mencatat [ADR-030](ADR-030-klien-web-pwa-lebih-dulu.md) butir 3, 9. Butir 2 dan 3 tidak berubah. Yang berlaku bagi klien web:

- **Butir 1, lingkungan:**
  - Staging web di subdomain staging yang sudah ada (origin lain dari produksi), dengan bundel staging tersendiri.
  - Pemisahan identitas aplikasi diganti pemisahan origin. Data browser per origin, jadi web app layar utama produksi dan staging dapat terpasang berdampingan.
  - Bundel staging hanya menjangkau server staging (CSP `connect-src 'self'`), dan halaman pemeriksaan diri hanya ada di bundel staging (SEC-179).
  - Server uji sementara FS-1 langkah 1 sampai 6 melayani bundel staging dengan header produksi.
- **Butir 4, irama:**
  - Setiap gelombang Tier 2 tetap satu rilis, dan bagi web rilis itu adalah rilis server, karena bundel disematkan di biner.
  - **Tanpa TestFlight, tanpa App Review, tanpa batas 90 hari:** aturan 60/80 hari dan QG-23 baru berlaku sejak build TestFlight pertama di jalur iOS native.
  - **Alur deploy web:**
    1. CI membangun bundel dari komit bertanda, mencatat hash setiap berkas, lalu membangun biner Go yang menyematkannya.
    2. Pengelola memasang biner ke staging, memeriksa header dan halaman pemeriksaan diri, lalu memasang biner yang sama ke produksi (`scp` lalu systemd restart; 09.4 §1).
    3. Runbook membandingkan hash biner di VPS dengan artefak CI.
    4. Klien mendapat versi baru saat aplikasi dibuka berikutnya (ADR-030 butir 9).
- **Butir 5, kompatibilitas:**
  - Bagi web, **versi klien minimum berlaku sejak pilot**, berbeda dari bunyi butir 5 "di pilot tidak dibangun". Alasannya: service worker dapat terus menyajikan bundel lama yang memuat celah (TM-45), dan pengelola tidak dapat memperbarui browser seperti memasang build.
  - Kode galat `CLIENT_UPGRADE_REQUIRED` dan versi minimum per platform dari butir 5 dibangun sekarang, dengan satu nilai untuk `web`. Nilai `ios` dan `android` tetap tidak dipakai di pilot. Catatan 13 November 2026 (FS-1 1.3, kontrak dasar; putusan lead Q2): versi minimum `web` adalah konfigurasi server (flag `casa serve`) dengan nilai bawaan konstanta di kode, sesuai ADR-030 butir 9, bukan konstanta saja.
  - Perintah yang ditolak tetap di antrean.
  - **Rollback web** adalah memasang biner sebelumnya. Klien yang sudah memuat bundel lebih baru tetap dilayani selama kontraknya kompatibel, dan `buf breaking` menjaganya.
- **Konsekuensi:** satu catatan aplikasi iOS kedua dan pembagian jatah Xcode Cloud baru diperlukan di jalur iOS native. FS-1 langkah 7 hanya menyiapkan instans staging dan subdomainnya.
