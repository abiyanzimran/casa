# ADR-022 — Jalur build iOS tanpa Mac

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **opsi A**, Xcode Cloud + TestFlight internal + sewa Mac sekali (S4-3 (i) A). Kata-katanya: "#S4-3: A Xcode Cloud + sewa Mac sekali - A iOS 26, Android 10" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban pertama; dikukuhkan dengan "lanjut"). Pemicu naik ke B atau C (butir 4) tetap berlaku |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4) |
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
- **biaya:** Apple Developer Program USD 99 per tahun; Xcode Cloud termasuk 25 jam per bulan; GitHub Actions macOS USD 0,062 per menit; Codemagic 500 menit gratis lalu USD 0,095 per menit; sewa Mac Scaleway sekitar EUR 0,22 per jam minimal 24 jam atau sekitar EUR 149 per bulan **[Perlu validasi]**; Mac mini M4 Rp12.749.000 **[Perlu validasi]**.

## Opsi

| Kriteria | **A. Xcode Cloud + TestFlight + sewa Mac sekali** | B. A + Mac sewaan bulanan selama fase UI | C. Mac mini sendiri | D. GitHub Actions macOS | E. xtool |
|---|---|---|---|---|---|
| Sesuai lisensi | Ya | Ya | Ya | Ya | **Belum** (L-31) |
| Biaya tetap | USD 99 per tahun | + sekitar EUR 149 per bulan | + Rp12.749.000 sekali | USD 99 + per menit | USD 99 |
| Pratinjau, simulator, debugger | Tidak | Ya | Ya | Tidak | Tidak |
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
3. **Kode non-UI** (klien API, antrean, sinkron, bantuan isian) ditulis sebagai paket Swift tanpa SwiftUI, agar dapat dibangun dan diuji di Linux **[Perlu validasi]**. Dengan begitu sebagian besar logika iOS teruji tanpa build cloud.
4. **Pemicu naik ke B atau C:** saat pekerjaan UI iOS menjadi harian, atau saat waktu tunggu build cloud melebihi waktu kerja yang wajar. Platform pertama adalah keputusan pemilik produk saat fase bangun dimulai. Bila iOS dibangun lebih dulu, B atau C hampir pasti diperlukan sejak awal.
5. **xtool tidak dijadikan jalur resmi** sampai 17 L-31 dijawab advokat.

## Konsekuensi

- Biaya iOS pilot hampir hanya keanggotaan Apple (USD 8,25 per bulan) selama jatah 25 jam cukup. Dengan build sekitar 10 menit **[Asumsi]**, jatah itu sekitar 150 build per bulan.
- Siklus UI iOS lambat tanpa Mac (RISK-01). Android sebagai platform pertama meringankan beban ini. Itu tetap keputusan pemilik produk, dan tradeoff-nya dinyatakan di sini.
- Ketergantungan pada Xcode Cloud dan akun Apple: bila Xcode Cloud bermasalah, jalur D (GitHub Actions macOS, dengan API `buildUploads`) dapat diaktifkan dengan biaya per menit.
- Artefak hanya disimpan 30 hari di Xcode Cloud. Versi yang dirilis dapat dibangun ulang dari tag.
