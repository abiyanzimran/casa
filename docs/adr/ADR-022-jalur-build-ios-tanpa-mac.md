# ADR-022 — Jalur build iOS tanpa Mac

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **opsi A**, Xcode Cloud + TestFlight internal + sewa Mac sekali (S4-3 (i) A). Kata-katanya: "#S4-3: A Xcode Cloud + sewa Mac sekali - A iOS 26, Android 10" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban pertama; dikukuhkan dengan "lanjut"). **Butir 4 diubah keputusan 43** (12 November 2026; diterima atas saran lead): tetap opsi A tanpa sewa bulanan dan tanpa membeli Mac; B atau C ditinjau lagi hanya bila dana ada (catatan keputusan 43). **Butir 1 diubah ADR-028 butir 4** (keputusan 44): batas keras 80 hari sejak unggahan. **Catatan keputusan 51:** jalur distribusi keluarga tanpa batas waktu dinilai ulang dengan temuan 18 §5.2.9; **Keputusan 55** (Q-228, 12 November 2026; diterima atas saran lead): tetap TestFlight internal, pindah ke App Store unlisted (C) bila Apple menghentikannya. Kata-kata pilihannya: "TestFlight, pindah bila dihentikan (Saran)" ([_progress.md](../_progress.md), "Jawaban pemilik produk atas pertanyaan sesudah keputusan 51 dan 52") |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4; fakta dan catatan S5; catatan keputusan 42; keputusan 43 dan 44; catatan keputusan 51, dinilai dengan 18 §5.2.9; keputusan 55; catatan web lebih dulu sesudah S6: ADR ini berlaku sejak jalur iOS native, bukan FS-1); 13 November 2026 sumber Mac butir 2 ditetapkan menurut bawaan 69.6 (catatan web) |
| Pengusul | `architect` |
| Mengubah | — (diubah oleh [ADR-028](ADR-028-lingkungan-dan-irama-rilis-pilot.md) butir 4 pada butir 1) |
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
   - **Pasangan pilot adalah penguji internal.** Pemilik produk memegang akun (Account Holder). Pasangannya ditambahkan sebagai pengguna App Store Connect dengan peran paling sempit di daftar itu, **Marketing** **[Usulan]**; peran itu dikukuhkan saat penyiapan bila ternyata tidak memberi akses TestFlight internal **[Perlu validasi]**. Akibatnya build pilot tidak melewati Beta App Review, sehingga bagian Guideline 1.2 yang hanya dituntut App Review, yaitu SEC-153 (d) dan SEC-154 selain (c) (06 §16.12), wajib selesai sebelum build pertama ke **TestFlight eksternal atau App Store**; SEC-154 selain (c) juga sebelum pengguna luar pertama yang membuat rumah tangga tanpa pengelola, bila itu lebih dulu (10.4 SEC-154; keputusan 52; pemicu kedua ditambahkan di S6). SEC-153 (a) sampai (c), (e) dan SEC-154 (c) tetap wajib di pilot (dokumen 10), karena melindungi anggota, bukan karena App Review.
   - **Build baru paling lambat setiap 90 hari** (sasaran: setiap 60 hari), agar build di perangkat pilot tidak kedaluwarsa. Pengingatnya pekerjaan kalender di runbook rilis. *Diubah ADR-028 butir 4 (keputusan 44, 12 November 2026): batas keras **80 hari** sejak build terakhir diunggah ke TestFlight, sasaran tetap 60; 90 hari tinggal sebagai batas kedaluwarsa Apple.*
2. **Penyiapan.** Mac jarak jauh disewa satu kali, minimal 24 jam (sekitar EUR 5,28 **[Perlu validasi]**), untuk membuat proyek, menghubungkan repositori, dan menyiapkan alur kerja pertama. Bila disewa dari penyedia selain AWS, kesesuaian sewanya dengan lisensi macOS diperiksa dulu (17 L-31 butir 3).
3. **Kode non-UI** (klien API, antrean, sinkron, bantuan isian) ditulis sebagai paket Swift tanpa SwiftUI, agar dapat dibangun dan diuji di Linux. Toolchain Swift resmi, Swift Package Manager, dan pustaka uji (XCTest, Swift Testing) tersedia di Linux **[Terverifikasi]** (18 §5.2.7.4); perilaku setiap API Foundation yang dipakai di Linux tetap **[Perlu validasi]** per dependensi. Dengan begitu sebagian besar logika iOS teruji tanpa build cloud.
   - *Tambahan S5 (`architect`, 12 November 2026; bukan bagian jawaban checkpoint S4, dan tidak mengubah jalur yang dipilih):* karena Security (Keychain), LocalAuthentication, dan UserNotifications juga tidak tersedia di Linux **[Terverifikasi]** (18 §5.2.7.4), paket non-UI **tidak mengimpor ketiganya, maupun SwiftUI**; pembungkusnya tinggal di lapisan aplikasi iOS di belakang antarmuka yang dapat dipalsukan dalam uji Linux **[Usulan]**.
4. **Tetap opsi A untuk seluruh pekerjaan UI iOS** (diubah keputusan 43, 12 November 2026). Mac sewaan bulanan (B) dan Mac mini sendiri (C) tidak dipakai; pekerjaan UI iOS harian, termasuk seluruh FS-1, dikerjakan lewat Xcode Cloud dan TestFlight. **B atau C ditinjau lagi hanya bila dana tersedia**, atas prakarsa pemilik produk; pekerjaan UI yang menjadi harian atau build cloud yang lambat tidak lagi memaksa B atau C. Titik impas di catatan S5 tetap sebagai acuan untuk saat itu.
   - *Bunyi lama (S4), tidak berlaku:* "Pemicu naik ke B atau C: saat pekerjaan UI iOS menjadi harian, atau saat waktu tunggu build cloud melebihi waktu kerja yang wajar. Platform pertama adalah keputusan pemilik produk saat fase bangun dimulai. Bila iOS dibangun lebih dulu, B atau C hampir pasti diperlukan sejak awal."
5. **xtool tidak dijadikan jalur resmi** sampai 17 L-31 dijawab advokat.

## Konsekuensi

- Biaya iOS pilot hampir hanya keanggotaan Apple (USD 8,25 per bulan) selama jatah 25 jam cukup. Dengan build sekitar 10 menit **[Asumsi]**, jatah itu sekitar 150 build per bulan.
- Siklus UI iOS lambat tanpa Mac (RISK-01). Android sebagai platform pertama meringankan beban ini. Itu tetap keputusan pemilik produk, dan tradeoff-nya dinyatakan di sini. *Sejak keputusan 42 dan 43:* iOS dibangun lebih dulu dan tanpa Mac, jadi siklus lambat ini berlaku sepanjang FS-1 dan FS-2 dan diterima atas saran lead (catatan keputusan 43).
- Ketergantungan pada Xcode Cloud dan akun Apple: bila Xcode Cloud bermasalah, jalur D (GitHub Actions macOS, dengan API `buildUploads`) dapat diaktifkan dengan biaya per menit.
- Artefak hanya disimpan 30 hari di Xcode Cloud. Versi yang dirilis dapat dibangun ulang dari tag.

## Catatan S5 (12 November 2026)

Keputusan ADR ini tidak berubah. Yang ditambahkan:

- **Butir 1, irama build:** tetap berlaku apa adanya: paling lambat setiap 90 hari, sasaran 60 hari, dihitung dari **unggahan** build ke TestFlight (18 §2.12.6: "All builds uploaded to TestFlight are active for 90 days"). [ADR-028](ADR-028-lingkungan-dan-irama-rilis-pilot.md) butir 4 **mengusulkan** batas keras 80 hari agar tersisa 10 hari untuk build yang gagal atau tertahan; usulan itu baru berlaku bila pemilik produk menyetujuinya (Q-219), karena butir 1 keputusannya sendiri (17 keputusan 32). *Kemudian:* diterima atas saran lead (keputusan 44, jawaban "Ya, 60/80 hari (Saran)"), jadi yang berlaku sejak 12 November 2026 adalah batas keras 80 hari sejak unggahan, sasaran 60.
- **Butir 4, membeli atau menyewa:** menurut hitungan `researcher` (18 §5.2.7.6, **[Usulan]**, dihitung ulang lead), membeli Mac mini impas terhadap sewa bulanan sesudah sekitar **7,1 sampai 7,9 bulan** untuk 16 GB/512 GB (konfigurasi minimum usulan `researcher`); 5,8 sampai 6,4 bulan untuk 16 GB/256 GB (8,8 bulan terhadap sewa berkomitmen 12 bulan); 5,5 sampai 6,2 bulan untuk 24 GB/512 GB. Nilai jual kembali, pajak, dan biaya kartu tidak dihitung. Hanya Mac sendiri di jaringan rumah yang memberi debugger di perangkat pilot. *Sejak keputusan 43:* angka ini hanya acuan bila pemilik produk meninjau B atau C lagi saat dana tersedia (14 §9).
- **Konsekuensi, jalur D:** bila Xcode Cloud bermasalah, jalur D tetap jalan mundur seperti tertulis di Konsekuensi, dan dijalankan sebagai **pemeriksaan wajib sebelum merge**, tidak pernah sesudah merge (ES §7; 14 §5.1). Biayanya USD 0,062 per menit **[Terverifikasi]** (18 §5.2), dipotong dari jatah GitHub yang sama dengan CI Linux. Jalan mundur yang menggabungkan perubahan iOS sebelum pemeriksaannya lulus tidak dipakai; bila suatu saat diperlukan, ia memerlukan ADR baru yang disetujui pemilik produk.
- Fakta Konteks diperbarui dari 18 §5.2.7 (harga Mac mini M6, Mac sewaan, AWS, Swift di Linux).

## Catatan keputusan 42 (12 November 2026)

Keputusan ADR ini tidak berubah. Yang dicatat:

- **Platform pertama sudah diputuskan** (butir 4 kalimat kedua): iOS dan iPadOS lebih dulu, dan pilot dengan data nyata berjalan di iPhone 13 dan iPad mini 7; Android sesudah GR-2 (keputusan 42; ADR-012 catatan urutan platform). Kalimat Konsekuensi "Android sebagai platform pertama meringankan beban ini" kini hanya riwayat tradeoff.
- **Pekerjaan UI iOS harian sejak awal FS-1.** iOS dan iPadOS satu-satunya platform UI sampai GR-2, dan layar SwiftUI dikerjakan sejak FS-1 langkah 1 (14 §3.3, §3.4), jadi pekerjaan UI iOS menjadi harian sejak awal **[Asumsi]**. Catatan ini semula menyimpulkan bahwa pemicu butir 4 lama akan terjadi di GR-1; keputusan 43 menggantinya: pekerjaan harian itu dikerjakan dengan opsi A.
- **Sewa sekali untuk penyiapan (butir 2)** adalah sewa bertarif per jam dengan minimum 24 jam berturut-turut, bukan sewa sebulan (18 §5.2.7.2: minimum 24 jam **[Terverifikasi]**; 24 × EUR 0,22 = EUR 5,28 **[Perlu validasi]** harga). Sewa sekali ini tetap diperlukan, karena B dan C tidak dipakai (keputusan 43).
- Android tidak memerlukan Mac; datangnya Android sesudah GR-2 tidak mengubah butir mana pun di ADR ini.

## Catatan keputusan 43 dan 44 (12 November 2026)

**Keputusan 43 mengubah butir 4.** Pemilik produk bertanya, "mac bingung gada dana kan keputusannya pake xcode cloud?", lalu memilih "Tetap Xcode Cloud saja (Saran)" dari tiga pilihan (sewa Mac bulanan, beli Mac mini). Keterangan pilihan yang ia pilih: tanpa sewa bulanan dan tanpa beli; Mac disewa sekali untuk penyiapan; uji lewat TestFlight; lebih lambat, tanpa debug di perangkat; ditinjau lagi kalau dana ada ([_progress.md](../_progress.md), "Penutupan pertanyaan terbuka sebelum S6"). Butir 1 sampai 3 dan 5 tidak berubah.

Kedudukan: diterima atas saran lead (pemilik produk memilih pilihan bertanda "(Saran)"). Konsekuensi yang ikut diterima dengan pilihan itu, karena tercantum di keterangannya:

- **Siklus UI iOS lambat sepanjang FS-1 dan FS-2.** Setiap perubahan layar SwiftUI dilihat lewat build Xcode Cloud lalu TestFlight, tanpa pratinjau SwiftUI, simulator, atau debugger di iPhone 13 dan iPad mini 7. Bug UI dicari dari log, laporan crash TestFlight, dan XCUITest di simulator Xcode Cloud, bukan dengan menghentikan aplikasi di perangkat (RISK-01, diterima).
- **Yang meringankan, semuanya tanpa biaya tambahan:** logika non-UI di paket Swift yang diuji di Linux (butir 3), sehingga yang tersisa untuk siklus lambat hanya tampilan; perubahan layar dikumpulkan per build; jatah 25 jam Xcode Cloud cukup untuk sekitar 150 build per bulan **[Asumsi]** build 10 menit; bila jatah hampir habis, aturan 80% di 14 §6.3.
- **Jadwal FS-1 memuat harga ini.** Perkiraan waktu layar iOS di 14 tidak boleh dihitung seolah-olah ada pratinjau dan simulator **[Usulan]**.
- **Peninjauan ulang:** hanya bila dana tersedia, atas prakarsa pemilik produk. Titik impas membeli terhadap menyewa (catatan S5; 14 §9) dipakai sebagai acuan saat itu, dengan harga yang dicek ulang (ES §8 berlaku juga atas harga).

**Keputusan 44 mengubah butir 1** lewat ADR-028 butir 4: pemilik produk menjawab "Ya, 60/80 hari (Saran)" atas pertanyaan rilis pemeliharaan bila 60 hari tanpa rilis, batas keras 80 hari. Yang berlaku: rilis pemeliharaan bila 60 hari tanpa rilis; build iOS berikutnya diunggah paling lambat 80 hari sesudah unggahan sebelumnya; 90 hari adalah kedaluwarsa TestFlight (18 §2.12.6), bukan lagi batas kerja.

## Catatan keputusan 51: TestFlight internal sebagai jalur keluarga tanpa batas waktu

**Status catatan: diputuskan (keputusan 55).** Pemilik produk memilih pilihan 1 dari rekomendasi di bawah: butir 1 (jalur A) tetap berlaku, dan C menjadi jalan keluar bila Apple menghentikan TestFlight (bagian "Keputusan 55" di akhir catatan). Penilaian di bawah memakai temuan `researcher` di 18 §5.2.9 (diakses 12 November 2026). Bagian "Riwayat penilaian" di akhir mencatat versi yang menunggu riset.

### Konteks baru dari 18 §5.2.9

- **Tidak ada saluran Apple untuk keluarga tanpa badan hukum yang sekaligus resmi, tanpa Mac, dan bebas batas 90 hari** (18 §5.2.9.5, **[Usulan]** `researcher`).
- **TestFlight internal (A).** Guideline 2.2 kemungkinan tidak jelas menjangkau build internal yang tidak diajukan ke TestFlight App Review (**[Usulan]** `researcher`). Batas yang mengikat ada di kontrak, DPLA 7.4.1 (versi 29 Oktober 2026), **[Terverifikasi]**:
  - yang didistribusikan "pre-release versions";
  - pemakaiannya "solely for their internal use in testing, evaluating and/or developing";
  - penerimanya "members of Your company or organization", padahal pengguna tambahan pada akun perorangan "not part of your team";
  - Apple dapat meminta distribusi dihentikan "at any time in its sole discretion".
  
  Catatan penegakan atas pemakaian seperti ini tidak ditemukan **[Perlu validasi]**.
- **Ad Hoc (B).** Bunyi DPLA 7.3 ("otherwise affiliated with You", tanpa kata "testing") paling dekat dengan pemakaian keluarga, dan Introduction App Review Guidelines menyebutnya untuk "family and friends" **[Terverifikasi]**. Kendalanya:
  - 100 perangkat per keluarga produk per tahun;
  - pemeriksaan daring PPQ saat aplikasi pertama dibuka;
  - pemasangan `.ipa` lewat Device Hub atau Apple Configurator memerlukan **Developer Mode**, yang menurut Apple "reduces the security of your device" **[Terverifikasi]**; apakah pemasangan OTA juga memerlukannya **[Perlu validasi]**;
  - masa berlaku profil belum terverifikasi;
  - **Xcode Cloud tidak menawarkan ekspor Ad Hoc** menurut dokumentasinya, sehingga B memerlukan jalur build di Mac.
- **App Store unlisted (C).** Hanya lewat tautan, dengan App Review penuh untuk setiap versi; permintaan ditolak bila aplikasi berstatus beta atau prerelease; siapa pun yang memegang tautan dapat memasang. Kelayakan akun perorangan **[Perlu validasi]**. Dampak terhadap pengecualian kegiatan rumah tangga (PP 33/2026 Pasal 3 ayat (2) huruf c, "tidak diperuntukkan bagi publik") perlu advokat. Dapat dibangun dari Xcode Cloud tanpa Mac, dan tanpa batas 90 hari.
- **Custom Apps (D)** memerlukan organisasi dengan D-U-N-S, jadi tidak cocok. **App Store publik (E)** bertentangan dengan keputusan 51.

### Kriteria dan matriks

Kriteria: (1) kesesuaian dengan kontrak dan pedoman Apple; (2) tanpa Mac, karena keputusan 43 (diterima atas saran lead) menetapkan tanpa sewa bulanan dan tanpa membeli Mac sampai dana ada; (3) keamanan perangkat yang memegang data nyata; (4) batas waktu build terpasang; (5) beban rilis dan jarak dari jalur publik (keputusan 51); (6) biaya; (7) pertanyaan hukum yang dibuka. D dan E gugur karena syarat (badan hukum; keputusan 51) dan tidak dinilai lebih jauh.

| Kriteria | **A. TestFlight internal (tetap)** | B. Ad Hoc | C. App Store unlisted |
|---|---|---|---|
| Kontrak dan pedoman Apple | Ketegangan dengan DPLA 7.4.1 (pre-release, internal testing, anggota organisasi); Apple dapat menghentikan kapan saja; penegakan tidak ditemukan | Paling dekat (DPLA 7.3; Introduction pedoman) | Saluran rilis resmi; kelayakan akun perorangan dan keluarga sebagai *limited audience* belum terverifikasi |
| Tanpa Mac (keputusan 43) | Ya | **Tidak**: Xcode Cloud tanpa ekspor Ad Hoc; runner macOS CI lain **[Perlu validasi]**. Bertentangan dengan keputusan 43 kecuali pemilik produk membukanya kembali | Ya ("TestFlight and App Store" di Xcode Cloud) |
| Keamanan perangkat | Tidak berubah | **Developer Mode** di iPhone 13 dan iPad mini 7 yang memegang data nyata, kemungkinan; PPQ memerlukan internet saat pertama dibuka | Tidak berubah. Siapa pun yang memegang tautan dapat memasang, tetapi tanpa undangan yang dikukuhkan pengundang tidak masuk ke rumah tangga (F-FAM-002), dan pendaftaran tertutup di server (SEC-158, Pilot bagi semua opsi; ditambahkan sesudah penilaian ini, 10 §3.3) |
| Batas waktu build | 90 hari per build; irama 60/80 hari (ADR-028 butir 4) | Masa profil (belum terverifikasi) | Tidak ada |
| Beban rilis; jarak dari jalur publik | Tetap seperti sekarang | Pendaftaran UDID, profil, dan ekspor per rilis di Mac | App Review penuh setiap rilis, termasuk perbaikan mendesak; akun demo untuk peninjau; prasyarat toko dari jalur publik maju ke sebelum pemakaian: SEC-153 (d), sisa SEC-154, SEC-83, label privasi (14 QG-31 (b)) |
| Biaya di luar USD 99 per tahun | Tidak ada | Mac atau runner macOS (14 §9; USD 0,062 per menit untuk runner GitHub, 18 §5.2.3) | Tidak ada; waktu pengembang untuk prasyarat toko |
| Pertanyaan hukum | DPLA 7.4.1 (advokat) | DPLA 7.3 (advokat) | DPLA dan pengecualian rumah tangga PP 33/2026 Pasal 3 (advokat) |

### Rekomendasi [Usulan] `architect`: tetap A, dengan C sebagai jalan keluar yang sudah direncanakan

1. **Tetap TestFlight internal untuk aplikasi harian keluarga.** Hanya A yang memenuhi keputusan 43 tanpa menambah beban rilis atau melemahkan keamanan perangkat. Risikonya kontraktual dan berakibat kehilangan akses, bukan kebocoran data: data tetap di server (ADR-013). Irama 80 hari (ADR-028 butir 4) menjamin build terpasang masih berlaku sekurangnya 10 hari bila Apple menghentikan distribusi.
2. **Jalan keluar yang direncanakan adalah C**, karena C tetap tanpa Mac. Pemicunya:
   - Apple meminta distribusi TestFlight dihentikan, atau ada tanda dari App Store Connect (galat kontrak beta atau build yang tiba-tiba tidak tersedia);
   - advokat menyatakan pemakaian A melanggar DPLA 7.4.1.
   
   Syarat sebelum C dipakai ada di tabel "Syarat opsi C" di bawah.
3. **B tidak direkomendasikan.** B memerlukan Mac, yang ditolak keputusan 43 sampai dana ada, dan kemungkinan menyalakan Developer Mode di perangkat yang memegang data nyata. Sejak 10 §0, **Developer Mode tidak pernah menyala di perangkat berdata nyata**, sebagai aturan dengan uji; `security-privacy` menolak B dari sisi keamanan selama B memerlukannya (10 §3.3). B baru dinilai lagi bila pemilik produk membuka keputusan 43 **dan** pemasangan OTA tanpa Developer Mode terbukti.
4. **Bila A dihentikan sebelum C siap, ada jeda tanpa aplikasi iOS.** Lamanya jeda = waktu menyelesaikan prasyarat C + App Review, dikurangi sisa masa build terpasang (10 sampai 90 hari). Lama App Review untuk aplikasi ini **[Asumsi]** belum diketahui. Selama jeda, pencatatan berhenti di iOS, tetapi catatan yang terlambat diterima dengan tanggal kejadiannya (07 §3.11). Catatan di antrean perangkat yang belum terkirim saat build kedaluwarsa dapat hilang, jadi saat pemicu datang, aplikasi diminta mengosongkan antrean lebih dulu.
5. **Pertanyaan advokat baru** (diminta ke `product-analyst` untuk 17, sesuai 18 §5.2.9 butir tinjauan 1 sampai 3): DPLA 7.4.1 untuk A, DPLA 7.3 untuk B, dan pengaruh App Store unlisted terhadap pengecualian rumah tangga untuk C. Keputusan 67 (12 November 2026, diputuskan pemilik produk sendiri): pertanyaan ini tidak dikejar dan dibuka lagi hanya bila go public diputuskan.

### Syarat opsi C (dikukuhkan `security-privacy`, 10 §3.3)

Semua butir wajib lulus sebelum build App Store pertama dipasang di perangkat keluarga. Selama A berjalan, tidak ada yang dijadwalkan; daftarnya disiapkan supaya pindah ke C tidak dirancang ulang saat pemicu datang.

| # | Syarat | Pemilik | Sumber |
|---|---|---|---|
| C-1 | Akun Apple Developer Program perorangan terbukti layak untuk distribusi unlisted (permintaan ke Apple dijawab) | Pemilik produk | 18 §5.2.9.3 |
| C-2 | Advokat menjawab apakah distribusi lewat App Store, walau hanya lewat tautan, masih "tidak diperuntukkan bagi publik" (PP 33/2026 Pasal 3 ayat (2) huruf c; L-01). S6: jawaban advokat atas L-01 (keputusan 59) menyatakan pengecualian berlaku selama aplikasi nonkomersial dan tertutup untuk keluarga dan baru gugur bila benar-benar dibuka untuk umum, tetapi tidak menyebut distribusi App Store unlisted. Keputusan 67 (diputuskan pemilik produk sendiri): tidak ada konsultasi advokat lanjutan, dan pemicu konsultasi ulang usulan `security-privacy` (termasuk App Store tersembunyi) diringkas menjadi satu, yaitu sebelum go public. Jadi syarat advokat L-01 atas C dihapus (keputusan 67; dikukuhkan dokumen 10 sesudah keputusan 68): C-2 tidak lagi menahan C, dan pertanyaan App Store unlisted dibuka lagi hanya bila go public diputuskan | — | 10 §3.3; 18 §5.2.9 butir tinjauan 3 |
| C-3 | Pendaftaran tertutup di server (SEC-158). Ini sudah syarat Pilot sejak GR-2 bagi semua opsi, jadi di C hanya diperiksa ulang | `backend-dev` | 10.4 SEC-158 |
| C-4 | Jalan akun demo App Review dipilih dan siap (di bawah) | `ios-dev`; pemilik produk | Guideline 2.1(a) |
| C-5 | Prasyarat toko dari 14 QG-31 (b) yang berlaku bagi kiriman App Store: SEC-153 (d), SEC-83 rating usia, label privasi, SEC-139 versi publik di URL publik (keputusan 68 menunda SEC-139 semua versi ke go public; bila C dipakai, kiriman App Store menariknya maju, karena toko memintanya) | `ios-dev`; `security-privacy` | 14 §5.3 |
| C-6 | Tinjauan penuh `security-privacy` atas 10 §0 dan SEC-41 (d) sebelum C dipakai. Bahan tinjauan awal: build App Store adalah build produksi dengan pemeriksaan arsip CI yang sama; satu perangkat tidak dapat memegang build TestFlight staging dan build App Store dengan bundle ID yang sama, jadi uji perangkat 10 §0 dijalankan di build staging sebelum pindah, atau dipakai bundle ID staging tersendiri **[Perlu validasi]**; domain yang dihubungi (SEC-49) diperiksa ulang lewat App Privacy Report | `security-privacy`; `architect` | 10 §3.3, §7.8 |
| C-7 | Antrean perangkat dikosongkan sebelum build TestFlight terakhir kedaluwarsa | `ios-dev`; pemakai | Butir 4 di atas |

**Jalan akun demo App Review [Usulan] `architect`, mengikuti anjuran `security-privacy`.** Guideline 2.1(a) meminta info akun demo bila aplikasi memiliki login. Demo mode bawaan hanya boleh dipakai "with prior approval by Apple" **[Terverifikasi]** (10 §3.3). Build App Store hanya menghubungi server produksi (SEC-41 (d)), jadi akun demo di produksi adalah akun orang asing di server keluarga, yaitu pengguna luar pertama (Q-209).
- **Dipilih: (i) demo mode di perangkat.** Data sintetis dibundel di build, tanpa koneksi ke server mana pun, dan dimintakan izin Apple lebih dulu. Demo mode ini memakai alur dan layar yang sama. Aturan uang tetap di server (ADR-013): demo mode menampilkan hasil yang sudah dihitung dari vektor emas, bukan salinan aturan kedua di klien. Uji: dalam demo mode tidak ada permintaan jaringan ke domain produksi (App Privacy Report), dan tidak ada data demo yang tersimpan di luar wadah demo.
- **Cadangan, bila Apple menolak (i): (ii) rumah tangga demo di produksi**, hanya sesudah QG-31 (a) dan (a+) lulus. Isinya data sintetis saja, dibuat lewat token pemasangan SEC-158 (b), dan dikunci atau dihapus di antara peninjauan. Harganya: QG-31 (a) dan (a+) maju ke sebelum C.

### Pilihan yang tidak diambil (dicatat sebagai bahan bila keputusan 55 dibuka kembali)

- **Pilihan 2 (A sekarang, C disiapkan sebelum GR-2):** prasyarat toko dari 14 QG-31 (b) dikerjakan sebelum GR-2: SEC-153 (d), sisa SEC-154, SEC-83, label privasi, dan rencana akun demo yang hanya berisi data uji. Jeda pada butir 4 menyusut menjadi lama App Review saja. Harganya waktu pengembang di FS-1 dan pertanyaan advokat C yang dimajukan.
- **Pilihan 3 (C untuk aplikasi harian sejak GR-2; TestFlight internal hanya untuk aplikasi staging):** pembagian ini paling cocok dengan bunyi DPLA, karena TestFlight dipakai untuk menguji dan App Store untuk versi final. Harganya: App Review untuk setiap rilis produksi, termasuk perbaikan mendesak (14 §6.3); prasyarat toko sebelum GR-2; akun demo peninjau di instans produksi berisi data uji; keputusan atas pendaftaran terbuka; dan jawaban advokat soal pengecualian rumah tangga **sebelum** data nyata. 14 §6, QG-23, QG-25, dan §4.8 direvisi; QG-23 (80 hari) tidak lagi berlaku untuk aplikasi harian.

### Keputusan 52

Orang tua yang kelak diberi login juga memerlukan jalur pemasangan. Di A, setiap anggota keluarga yang memakai iPhone menjadi pengguna App Store Connect di akun pengembang pemilik produk (butir 1), dan ini menambah ketegangan dengan DPLA 7.4.1 ("members of Your company or organization"). Di C, mereka memasang lewat tautan. Ini dicatat sebagai bahan bagi pemilik produk saat orang tua diberi login, bukan pemicu C; pemicunya hanya yang ditetapkan keputusan 55. Bila mereka dipasang lewat TestFlight internal, setiap orang mendapat peran App Store Connect tersempit (14 QG-31 (a); 10.4 §2.1).

### Keputusan 55

Pertanyaan lead (Q-228) menawarkan tiga pilihan di atas; Ad Hoc tidak ditawarkan karena memerlukan Mac (keputusan 43). Pemilik produk memilih "TestFlight, pindah bila dihentikan (Saran)". **Kedudukan: diterima atas saran lead** (`_progress.md`, "Jawaban pemilik produk atas pertanyaan sesudah keputusan 51 dan 52", baris 55).

Akibatnya:
- butir 1 tetap: Xcode Cloud ke TestFlight internal, dengan irama 60/80 hari (ADR-028 butir 4);
- C adalah jalan keluar, dipicu hanya oleh dua pemicu di butir 2 rekomendasi: Apple menghentikan distribusi, atau advokat menyatakan A melanggar DPLA 7.4.1;
- syarat C-1 sampai C-7 tidak dijadwalkan sebelum pemicu, kecuali C-3 (SEC-158), yang memang syarat Pilot;
- risiko sisanya, yaitu jeda tanpa aplikasi iOS selama syarat C dan App Review berjalan, diterima atas saran lead (keputusan 55; 15 RISK-92).

### Riwayat penilaian: versi yang menunggu riset

- **Mengapa dibuka.** Saat ADR ini diterima, TestFlight internal adalah jalur uji untuk pilot yang berbatas waktu menuju beta dan rilis publik. Sejak keputusan 51, go public tidak direncanakan, jadi TestFlight internal dapat menjadi **satu-satunya jalur distribusi iOS keluarga tanpa batas waktu**.
- **Yang diangkat `security-privacy`** (10 §3.3 butir 7, §8 butir 15, §7.8 permintaan (1)): Apple Guideline 2.2 menyatakan aplikasi di TestFlight "should be intended for public distribution", dan Apple mengarahkan aplikasi yang hanya untuk keluarga sendiri ke pemasangan lewat Xcode atau distribusi Ad Hoc. Kutipan dan bacaan ini **[Perlu validasi]** di dokumen ini sampai `researcher` memeriksanya di halaman resmi Apple dan mencatatnya di 18.
- **Akibat yang dinilai: kehilangan akses, bukan kebocoran data.** Bila Apple menolak atau menghentikan pemakaian itu, build baru tidak dapat dipasang dan build terpasang kedaluwarsa dalam 90 hari, sehingga istri (iPhone 13) dan pemilik produk (iPad mini 7) kehilangan aplikasi. Data rumah tangga tetap utuh di server (sumber kebenaran; ADR-013); yang dapat hilang hanya catatan di antrean perangkat yang belum terkirim. Dinilai di 15 RISK-92.
- **Yang sedang diperiksa `researcher`:** (1) apakah Guideline 2.2 berlaku bagi build TestFlight internal yang tidak melewati Beta App Review; (2) distribusi Ad Hoc tanpa Mac: batas perangkat, masa berlaku profil, dan pendaftaran perangkat; (3) apakah Xcode Cloud dapat mengekspor build Ad Hoc tanpa Mac; (4) distribusi App Store *unlisted*.
- **Sesudah temuan tiba:** ADR ini dinilai ulang dengan matriks opsi (A tetap, Ad Hoc lewat Xcode Cloud, App Store *unlisted*, dan opsi lain yang ditemukan) terhadap kriteria ADR ini ditambah kesesuaian dengan aturan Apple untuk pemakaian tanpa batas waktu. Bila jalurnya diganti, `security-privacy` meninjau ulang 10 §0 (uji iOS tanpa Mac) dan SEC-41 (d) sebelum jalur baru dipakai (10 §7.8), dan pemilik produk yang memutuskan, karena butir 1 adalah keputusannya (S4-3 (i) A).
- **Keputusan 52 menambah satu kriteria.** Orang tua yang kelak diberi login juga memerlukan jalur pemasangan. Di jalur A, setiap penguji internal harus menjadi pengguna App Store Connect dengan satu peran (butir 1), sehingga setiap anggota keluarga yang memakai iPhone mendapat akun di akun pengembang pemilik produk. Matriks opsi menilai juga hal itu.
- **Kapan harus selesai** **[Usulan]**: sebelum GR-1, karena jalur build menentukan penyiapan FS-1 langkah 1 (Apple Developer Program, sewa Mac sekali, alur kerja Xcode Cloud pertama); paling lambat sebelum GR-2, karena sesudah itu kehilangan jalur berarti kehilangan aplikasi yang memegang data nyata. *Sejak catatan web lebih dulu, waktunya mengikuti jalur iOS native, bukan GR-1.*

## Catatan keputusan web lebih dulu (12 November 2026, sesudah S6)

Mencatat [ADR-030](ADR-030-klien-web-pwa-lebih-dulu.md) butir 19. Butir 1 sampai 5 dan keputusan 55 **tidak berubah isinya**; yang berubah adalah **kapan** ADR ini berlaku.

- **Tidak lagi di jalur FS-1 dan GR-2.** Klien web dibangun, diuji, dan dideploy di Linux tanpa Mac (18 §5.13.8). Apple Developer Program, Xcode Cloud, TestFlight, dan sewa Mac sekali baru diperlukan **saat jalur iOS native dimulai** (14 §3.9), sesudah GR-2.
- **Butir 2, sumber Mac (temuan audit 2):** pilihan sumber Mac sekali (pinjam, AWS, Scaleway) **diputuskan saat jalur iOS native dimulai**, sesuai pernyataan pemilik produk "itu nanti aja" (`_progress.md`, "Pernyataan pemilik produk tentang sewa Mac"). Kalimat butir 2 tentang pemeriksaan lisensi untuk penyedia selain AWS (17 L-31 butir 3) bertabrakan dengan keputusan 67, yang menghentikan konsultasi advokat sampai go public. **Ditetapkan (bawaan 69.6, diterima atas saran lead, 13 November 2026):** sumbernya **AWS EC2 Mac** (minimum 24 jam, Singapura `mac2` USD 18,696 untuk 24 jam, belum termasuk EBS; 09.4 §7.1) **atau Mac pinjaman** milik pemiliknya sendiri; pilihan di antara keduanya diambil pemilik produk saat jalur iOS native dimulai. Keduanya tidak memerlukan tafsir lisensi, jadi L-31 butir 3 tidak perlu dijawab dan kalimat lisensi butir 2 tidak berlaku. *Catatan riwayat:* Scaleway (EUR 5,28 **[Perlu validasi]**) tidak dipakai lagi.
- **Butir 4 dan keputusan 43 tetap**, tetapi beban yang diterima mengecil: siklus UI iOS lambat kini hanya berlaku di jalur iOS native, bukan di jalur kritis data nyata (15 RISK-01).
- **Butir 1 dan keputusan 55:** batas 80 hari (ADR-028 butir 4), QG-23, QG-25, syarat C, dan RISK-52 serta RISK-92 baru berlaku sejak build TestFlight pertama di jalur iOS native.
- **Biaya:** USD 99 per tahun keluar dari biaya pilot sampai jalur iOS native dimulai (09.4 §7.1).
- **Kalimat konsekuensi "Android sebagai platform pertama meringankan beban ini"** tidak lagi relevan: web adalah platform pertama, dan urutan native sesudah GR-2 dibahas di ADR-012 catatan web.
