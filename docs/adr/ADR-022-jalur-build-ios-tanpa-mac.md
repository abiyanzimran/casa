# ADR-022 — Jalur build iOS tanpa Mac

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **opsi A**, Xcode Cloud + TestFlight internal + sewa Mac sekali (S4-3 (i) A). Kata-katanya: "#S4-3: A Xcode Cloud + sewa Mac sekali - A iOS 26, Android 10" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban pertama; dikukuhkan dengan "lanjut"). Pemicu naik ke B atau C (butir 4) tetap berlaku |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4; fakta dan catatan S5) |
| Pengusul | `architect` |
| Mengubah | — |
| Rincian | [09.1 §12](../09-technical-architecture/01-pilihan-teknologi.md); [09.4 §5](../09-technical-architecture/04-operasi-uji-biaya.md) |

## Konteks

Pengembangan dilakukan di Linux tanpa Mac, dan pilot memerlukan iPhone 13 dan iPad mini 7 (brief §1). UI iOS native SwiftUI (ADR-012).

Riset 18 §5.2 menemukan:

- lisensi Xcode dan SDK Apple (EA2002) hanya mengizinkan pemakaian di perangkat keras Apple yang menjalankan macOS;
- simulator hanya ada di macOS;
- alur kerja Xcode Cloud pertama harus disiapkan di Xcode, dan artefaknya disimpan 30 hari;
- API App Store Connect `buildUploads` tersedia;
- xtool dapat membangun dari Linux tetapi memerlukan isi Xcode `.xip`, dan soal lisensinya dicatat sebagai 17 L-31;
- **TestFlight** (18 §2.12.6) **[Terverifikasi]**: penguji internal paling banyak 100, harus pengguna App Store Connect dengan peran Account Holder, Admin, App Manager, Developer, atau Marketing, masing-masing sampai 30 perangkat; penguji eksternal paling banyak 10.000, dan build eksternal pertama memerlukan Beta App Review; build kedaluwarsa sesudah 90 hari. Bahwa build internal tidak ditinjau Apple **[Perlu validasi]** (sangat tersirat dokumentasi Apple). TestFlight mengirim laporan crash otomatis, terlepas dari setelan perangkat (18 §5 fase D);
- **biaya:** Apple Developer Program USD 99 per tahun; Xcode Cloud termasuk 25 jam per bulan; GitHub Actions macOS USD 0,062 per menit; Codemagic 500 menit gratis lalu USD 0,095 per menit; sewa Mac Scaleway sekitar EUR 0,22 per jam atau sekitar EUR 149 per bulan **[Perlu validasi]** (harga dari cuplikan), dengan minimum 24 jam **[Terverifikasi]** (18 §5.2.7.2); penawaran bulanan M4 16 GB lain USD 149–164 **[Terverifikasi]**; AWS EC2 Mac On-Demand per jam, minimum 24 jam: Sydney `mac-m4` USD 1,538, Singapura `mac2` (M1) USD 0,779 **[Terverifikasi]**; **tidak ada Mac sewaan di Jakarta**, terdekat Malaysia dan Singapura (18 §5.2.7.2); Mac mini M6 16 GB/256 GB "mulai dari" Rp17.499.000 **[Perlu validasi]** (dua media, 18 §5.2.7.1). *Revisi S5:* angka sebelumnya, Mac mini M4 Rp12.749.000, sudah tidak berlaku karena M4 digantikan M6.
- **Mac jarak jauh dalam praktik** (18 §5.2.7.3): Mac sewaan memulihkan pratinjau SwiftUI dan simulator **[Usulan]** `researcher`, tetapi tidak dapat menjalankan atau men-debug aplikasi di iPhone 13 atau iPad mini 7 lewat jalur Apple, karena pemasangan menuntut kabel ke Mac atau Wi-Fi yang sama **[Terverifikasi]**; dengan Mac sewaan, uji di perangkat tetap lewat TestFlight.

## Opsi

| Kriteria | **A. Xcode Cloud + TestFlight + sewa Mac sekali** | B. A + Mac sewaan bulanan selama fase UI | C. Mac mini sendiri | D. GitHub Actions macOS | E. xtool |
|---|---|---|---|---|---|
| Sesuai lisensi | Ya | Ya | Ya | Ya | **Belum** (L-31) |
| Biaya tetap | USD 99 per tahun | + sekitar EUR 149 atau USD 149–164 per bulan | + mulai dari Rp17.499.000 sekali (Mac mini M6) | USD 99 + per menit | USD 99 |
| Pratinjau, simulator, debugger | Tidak | Pratinjau dan simulator ya; debugger di iPhone fisik **tidak** (uji perangkat tetap lewat TestFlight) | Ya, termasuk debugger di iPhone 13 dan iPad mini 7 | Tidak | Tidak |
| Kecepatan siklus UI | Lambat (build cloud + TestFlight) | Cepat | Cepat | Lambat | Sedang |
| Penandatanganan | Dikelola cloud Apple | Sama | Lokal | Dikelola sendiri | Sendiri |

## Keputusan

1. **Jalur resmi A.**
   - Xcode Cloud membangun, menjalankan uji unit dan XCUITest di simulator, mengarsipkan, dan mengunggah ke TestFlight internal.
   - iPhone 13 dan iPad mini 7 memasang dari TestFlight.
   - Penandatanganan dikelola Xcode Cloud.
   - **Pasangan pilot adalah penguji internal.** Pemilik produk memegang akun (Account Holder). Pasangannya ditambahkan sebagai pengguna App Store Connect dengan peran paling sempit di daftar itu, **Marketing** **[Usulan]**; peran itu dikukuhkan saat penyiapan bila ternyata tidak memberi akses TestFlight internal **[Perlu validasi]**. Akibatnya build pilot tidak melewati Beta App Review, sehingga bagian Guideline 1.2 yang hanya dituntut App Review, yaitu SEC-153 (d) dan SEC-154 selain (c) (06 §16.12), wajib selesai sebelum build pertama ke **TestFlight eksternal atau App Store**. SEC-153 (a) sampai (c), (e) dan SEC-154 (c) tetap wajib di pilot (dokumen 10), karena melindungi anggota, bukan karena App Review.
   - **Build baru paling lambat setiap 90 hari** (sasaran: setiap 60 hari), agar build di perangkat pilot tidak kedaluwarsa. Pengingatnya pekerjaan kalender di runbook rilis.
2. **Penyiapan.** Mac jarak jauh disewa satu kali, minimal 24 jam (sekitar EUR 5,28 **[Perlu validasi]**), untuk membuat proyek, menghubungkan repositori, dan menyiapkan alur kerja pertama. Bila disewa dari penyedia selain AWS, kesesuaian sewanya dengan lisensi macOS diperiksa dulu (17 L-31 butir 3).
3. **Kode non-UI** (klien API, antrean, sinkron, bantuan isian) ditulis sebagai paket Swift tanpa SwiftUI, agar dapat dibangun dan diuji di Linux. Toolchain Swift resmi, Swift Package Manager, dan pustaka uji (XCTest, Swift Testing) tersedia di Linux **[Terverifikasi]** (18 §5.2.7.4); perilaku setiap API Foundation yang dipakai di Linux tetap **[Perlu validasi]** per dependensi. Dengan begitu sebagian besar logika iOS teruji tanpa build cloud.
   - *Tambahan S5 (`architect`, 12 November 2026; bukan bagian jawaban checkpoint S4, dan tidak mengubah jalur yang dipilih):* karena Security (Keychain), LocalAuthentication, dan UserNotifications juga tidak tersedia di Linux **[Terverifikasi]** (18 §5.2.7.4), paket non-UI **tidak mengimpor ketiganya, maupun SwiftUI**; pembungkusnya tinggal di lapisan aplikasi iOS di belakang antarmuka yang dapat dipalsukan dalam uji Linux **[Usulan]**.
4. **Pemicu naik ke B atau C:** saat pekerjaan UI iOS menjadi harian, atau saat waktu tunggu build cloud melebihi waktu kerja yang wajar. Platform pertama adalah keputusan pemilik produk saat fase bangun dimulai. Bila iOS dibangun lebih dulu, B atau C hampir pasti diperlukan sejak awal.
5. **xtool tidak dijadikan jalur resmi** sampai 17 L-31 dijawab advokat.

## Konsekuensi

- Biaya iOS pilot hampir hanya keanggotaan Apple (USD 8,25 per bulan) selama jatah 25 jam cukup. Dengan build sekitar 10 menit **[Asumsi]**, jatah itu sekitar 150 build per bulan.
- Siklus UI iOS lambat tanpa Mac (RISK-01). Android sebagai platform pertama meringankan beban ini. Itu tetap keputusan pemilik produk, dan tradeoff-nya dinyatakan di sini.
- Ketergantungan pada Xcode Cloud dan akun Apple: bila Xcode Cloud bermasalah, jalur D (GitHub Actions macOS, dengan API `buildUploads`) dapat diaktifkan dengan biaya per menit.
- Artefak hanya disimpan 30 hari di Xcode Cloud. Versi yang dirilis dapat dibangun ulang dari tag.

## Catatan S5 (12 November 2026)

Keputusan ADR ini tidak berubah. Yang ditambahkan:

- **Butir 1, irama build:** tetap berlaku apa adanya: paling lambat setiap 90 hari, sasaran 60 hari, dihitung dari **unggahan** build ke TestFlight (18 §2.12.6: "All builds uploaded to TestFlight are active for 90 days"). [ADR-028](ADR-028-lingkungan-dan-irama-rilis-pilot.md) butir 4 **mengusulkan** batas keras 80 hari agar tersisa 10 hari untuk build yang gagal atau tertahan; usulan itu baru berlaku bila pemilik produk menyetujuinya (Q-219), karena butir 1 keputusannya sendiri (17 keputusan 32).
- **Butir 4, membeli atau menyewa:** menurut hitungan `researcher` (18 §5.2.7.6, **[Usulan]**, dihitung ulang lead), membeli Mac mini impas terhadap sewa bulanan sesudah sekitar **7,1 sampai 7,9 bulan** untuk 16 GB/512 GB (konfigurasi minimum usulan `researcher`); 5,8 sampai 6,4 bulan untuk 16 GB/256 GB (8,8 bulan terhadap sewa berkomitmen 12 bulan); 5,5 sampai 6,2 bulan untuk 24 GB/512 GB. Nilai jual kembali, pajak, dan biaya kartu tidak dihitung. Hanya Mac sendiri di jaringan rumah yang memberi debugger di perangkat pilot. Keputusan B atau C tetap pada pemilik produk **saat pemicu butir 4 terjadi**; pada urutan platform A dan C itu sejak FS-1 langkah 1, sehingga praktis bersamaan dengan pilihan platform pertama (14 §3.4, §9).
- **Konsekuensi, jalur D:** bila Xcode Cloud bermasalah, jalur D tetap jalan mundur seperti tertulis di Konsekuensi, dan dijalankan sebagai **pemeriksaan wajib sebelum merge**, tidak pernah sesudah merge (ES §7; 14 §5.1). Biayanya USD 0,062 per menit **[Terverifikasi]** (18 §5.2), dipotong dari jatah GitHub yang sama dengan CI Linux. Jalan mundur yang menggabungkan perubahan iOS sebelum pemeriksaannya lulus tidak dipakai; bila suatu saat diperlukan, ia memerlukan ADR baru yang disetujui pemilik produk.
- Fakta Konteks diperbarui dari 18 §5.2.7 (harga Mac mini M6, Mac sewaan, AWS, Swift di Linux).
