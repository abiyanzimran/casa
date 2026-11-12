# ADR-012 — UI native per platform dan versi OS minimum

| | |
|---|---|
| Status | **Butir 1: Diputuskan pemilik produk** (12 November 2026): "Swift + Kotlin saya setuju untuk apple dan android catat dulu yaa, tinggal backend." ([_progress.md](../_progress.md); dokumen 17 Q-34). **Butir 2 sampai 6: Diterima** pemilik produk di checkpoint S4, 12 November 2026: butir 2 dan 3 lewat S4-3 (ii) A (iOS dan iPadOS 26; Android 10); butir 5 lewat S4-1 (iii) A (kalimat ADR-023 masuk ES §8); butir 4 dan 6 lewat "lanjut" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk") |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4; catatan urutan platform sesudah S5) |
| Pengusul | Butir 1: matriks `researcher` (18 §5.1), diputuskan pemilik produk. Butir 2 sampai 6: `architect` |
| Mengubah | — |
| Rincian | [09.1 §2, §13](../09-technical-architecture/01-pilihan-teknologi.md); [ADR-013](ADR-013-logika-bersama-hanya-di-server.md); [ADR-023](ADR-023-kebijakan-versi-0x.md) butir 5 |

## Konteks

Brief §9.1 dan §10.1 menjadikan tampilan dan rasa native syarat keras. Framework yang menggambar widget sendiri harus membuktikan bahwa ia setara dengan komponen sistem terkini di kedua platform dan mengikuti perubahan desain tahunan, atau ditolak.

Matriks 18 §5.1 menemukan bahwa hanya UI native penuh (dengan logika di server, di KMP, atau di inti Rust) yang lolos di kedua platform:

- Flutter dan Compose Multiplatform menggambar sendiri di iOS.
- React Native belum terbukti untuk Material 3 Expressive.
- Liquid Glass datang otomatis dengan SDK 26, dan opt-out-nya diabaikan mulai SDK 27.
- Unggahan ke App Store wajib memakai Xcode 26 sejak 28 April 2026.

ES §8 meminta versi OS minimum ditetapkan lewat ADR di S4, dan melarang alpha, beta, RC, atau API eksperimental tanpa ADR.

**Keadaan Compose Material 3 (diperiksa 12 November 2026)** **[Terverifikasi]**:

- Stabil **1.4.0**; pra-rilis terbaru **1.5.0-rc01** (4 November 2026). API Expressive tidak ada di 1.4.0 (dikeluarkan sejak 1.4.0-beta01) dan hanya ada di jalur 1.5.0.
- Di dalam 1.4.0 pun ada API yang masih bertanda eksperimental. `ModalBottomSheet` dan `rememberModalBottomSheetState` bertanda `@ExperimentalMaterial3Api` di cabang utama androidx, sehingga juga di 1.4.0. Material 3 Adaptive stabil 1.3.0, dan `ListDetailPaneScaffold` di dalamnya bertanda `@ExperimentalMaterial3AdaptiveApi`.
- `SwipeToDismissBox` dan `NavigationSuiteScaffold` **tidak** bertanda eksperimental di sumber 1.4.0 (18 §5.1.1, §5.1.3) **[Terverifikasi]**.

Sumbernya di [indeks 09 §9](../09-technical-architecture.md#9-fakta-yang-diperiksa-untuk-dokumen-ini).

## Opsi

| Kriteria | Flutter | React Native / Expo | Compose Multiplatform | **UI native penuh** |
|---|---|---|---|---|
| Native di iOS | Tidak | Ya | Tidak | **Ya** |
| Native di Android | Tidak | Belum terbukti (M3 Expressive) | Ya | **Ya** |
| Mengikuti desain tahunan | Ditiru setiap tahun | Mengikuti UIKit | Ditiru di iOS | **Otomatis lewat SDK** |
| Berbagi kode UI | Ya | Ya | Ya | Tidak (berbagi merek, kata, arsitektur informasi; brief §10.1) |

Untuk versi minimum:

| Kriteria | **iOS/iPadOS 26 + Android 10 (API 29)** | iOS 18 + Android 9 (API 28) | iOS 27 + Android 12 |
|---|---|---|---|
| Perangkat pilot | iPhone 13, iPad mini 7, Galaxy A55: semua didukung | Sama | iPhone 13 dan iPad mini 7 didukung iOS/iPadOS 27 (18 §5.1, [apple.com/ios](https://www.apple.com/ios/), [apple.com/ipados](https://www.apple.com/ipados/)) **[Terverifikasi]** |
| Bahasa desain iOS yang diuji | Satu (Liquid Glass) | Dua (sebelum dan sesudah Liquid Glass) | Satu |
| Jangkauan tahap publik | Sedang | Paling luas | Paling sempit |
| Android | Perangkat yang diluncurkan dengan Android 10 ke atas memakai enkripsi berbasis berkas **[Perlu validasi]** | Lebih banyak perangkat lama | — |

Untuk API eksperimental di dalam artefak stabil (lembar Catat di Android memakai lembar bawah):

| Kriteria | **A. `ModalBottomSheet` 1.4.0 lewat ADR-023 butir 5** | B. Dialog layar penuh dengan API stabil saja | C. Lembar bawah buatan sendiri |
|---|---|---|---|
| Rasa native (BT-7) | **Komponen Material resmi** | Native, tetapi pola berbeda dari 08 | Meniru; gagal gerbang native |
| ES §8 | Pengecualian tercatat dengan jalan mundur | Tanpa pengecualian | Tanpa pengecualian, tetapi kode lebih banyak |
| Risiko perubahan API | Ada; dipatok dengan versi | Tidak ada | Tidak ada |

## Keputusan

1. **UI native penuh** (diputuskan pemilik produk): **SwiftUI** di iPhone dan iPad, dan **Jetpack Compose dengan Material 3** di Android. Dua aplikasi berbagi warna merek, ikon, kata, dan arsitektur informasi, tetapi tidak berbagi kode UI. Tim build terdiri dari `ios-dev` dan `android-dev` yang terpisah (`CLAUDE.md` "Build team").
2. **Versi minimum iOS dan iPadOS: 26** (diterima, S4-3 (ii) A). Satu bahasa desain diuji, dan perangkat pilot didukung.
3. **Versi minimum Android: 10 (API 29)** (diterima, S4-3 (ii) A), dengan target API sesuai syarat Google Play yang berlaku (36 sejak 31 Agustus 2026) bila Play dipakai.
4. **SDK build:** Xcode dan SDK iOS stabil terbaru saat bangun (Xcode 27.1 sudah terlihat di citra build, 18 §5.2; minimal 26 untuk unggahan); Android SDK dan Compose BOM stabil terbaru. Keduanya diperiksa live saat bangun (ES §8).
5. **Compose Material 3**:
   - (a) **Versi:** 1.4.0, atau 1.5.0 bila sudah stabil saat fase bangun dimulai. 1.5.0-rc01 tidak dipakai. Selama 1.4.0, komponen Expressive tidak tersedia; `ux-designer` menyebut padanan 1.4.0 untuk setiap komponen Expressive di dokumen 08.
   - (b) **Lembar Catat di Android pilot memakai `ModalBottomSheet`**, tercatat di ADR-023 butir 5 (S4-1 (iii) A). Dialog layar penuh (`Dialog` dan `Scaffold` stabil) hanya jalan mundur bila `ModalBottomSheet` rusak atau berubah saat naik versi.
   - (c) **`ListDetailPaneScaffold`** tidak dipakai di pilot, karena perangkat Android pilot adalah ponsel (Galaxy A55). Navigasi memakai `NavigationSuiteScaffold`, yang stabil di 1.4.0.
   - (d) **`SwipeToDismissBox`** stabil di 1.4.0 dan boleh dipakai tanpa butir ADR-023.
   - (e) Setiap API eksperimental lain memerlukan butir ADR-023 baru sebelum dipakai.
6. **Yang tetap terbuka** (tidak diputuskan ADR ini): platform mana yang dibangun lebih dulu, yang menjadi keputusan pemilik produk saat fase bangun dimulai (`CLAUDE.md`). Letak logika bersama sudah diputuskan di ADR-013 (diterima di checkpoint S4, S4-1 (ii) A). *Sudah dijawab pemilik produk: lihat catatan di bawah.*

## Catatan urutan platform (12 November 2026, sesudah S5)

Bukan perubahan keputusan ADR ini; mencatat jawaban atas butir 6.

- **Keputusan 42** (17; `_progress.md`): "ios ipad yang bakalan dipakai pilot ya". iOS dan iPadOS dibangun lebih dulu (menjawab Q-37), dan pilot dengan data nyata berjalan di iOS dan iPadOS: istri di iPhone 13 (pencatat utama), pemilik produk di iPad mini 7. GR-2 tidak menunggu Android. Android dibangun sesudah GR-2 untuk Galaxy A55 milik pemilik produk; sebelum jalur pengujian Google Play dibuka (GR-4 butir 8; QG-31 (b)): semula bawaan lead, kini diterima atas saran lead lewat keputusan 50, dengan syarat beta dari keputusan 51 (bunyi brief §1: "Android is built after real data starts and joins the pilot before the beta, if a beta is ever held"). Brief §1 "Pilot devices" ("the pilot needs iOS and Android") digantikan (17 K-66); bunyi brief §1 "Pilot devices" dan "Platforms" yang baru diterima atas saran lead (keputusan 50) dan dipasang lead, dan ES §3 diubah lewat [ADR-029](ADR-029-perangkat-pengukuran-kinerja-es-3.md).
- **Butir 1 tetap:** Android tetap Kotlin dan Compose Material 3, dan tetap dirancang penuh; yang berubah hanya waktu bangunnya (14 §3.8). Tim bangun tetap `ios-dev` dan `android-dev` terpisah; `android-dev` mulai bekerja saat jalur Android dimulai.
- **Butir 2 dan 3 tetap.** Butir 5 berlaku sejak jalur Android dimulai; versi Compose Material 3 diperiksa ulang saat itu (ES §8), dan bila 1.5.0 sudah stabil, butir 5 (a) memilihnya. Butir 5 (c) tetap: perangkat Android pilot adalah ponsel.
- **Akibat:** pekerjaan UI hanya SwiftUI di FS-1, dan dikerjakan tanpa Mac lewat Xcode Cloud dan TestFlight, karena pemilik produk memilih tetap di opsi A tanpa sewa bulanan atau membeli Mac sampai dana ada (keputusan 43; ADR-022 butir 4; 14 §9). Risiko yang dibawa urutan ini: RISK-01 (diterima), RISK-18, RISK-91 (15).

## Konsekuensi

- Dua basis kode UI. Kualitas native di kedua platform diperoleh dengan harga dua kali pekerjaan layar. Tim satu sampai tiga orang menanggungnya karena logika tidak diduplikasi (ADR-013).
- Pembaruan desain tahunan Apple dan Google datang lewat SDK tanpa menulis ulang.
- iOS dibangun tanpa Mac lewat Xcode Cloud (ADR-022). Siklus UI iOS lebih lambat daripada Android.
- Pengguna iOS di bawah 26 dan Android di bawah 10 tidak dapat memasang aplikasi pada tahap publik. Bila data pilot atau riset pasar menunjukkan porsi yang berarti, butir 2 dan 3 ditinjau ulang sebelum tahap publik, bila diambil (keputusan 51).
- Android pilot tampil Material 3 tanpa Expressive bila 1.5.0 belum stabil. Peralihan ke Expressive adalah pembaruan versi, bukan penulisan ulang.
- `ModalBottomSheet` dipatok ke versi Compose yang tercatat; pembaruan versi menjalankan ulang uji UI lembar Catat (RISK-17).
