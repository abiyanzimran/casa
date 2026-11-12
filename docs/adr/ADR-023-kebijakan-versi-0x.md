# ADR-023 — Kebijakan pustaka versi 0.x dan API eksperimental

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **opsi C** dengan kalimat ES §8 di butir 3 (pertanyaan S4-1 (iii), huruf A). Kata-katanya: "yang penting kan versi stable, versi stable ga musti dari 1 kan? apalagi yang baru? cek goolge aja" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban keempat). Bacaan lead: rilis 0.x final, bukan pra-rilis, boleh dipakai, sama dengan bunyi butir 3; dikukuhkan dengan "lanjut". Kalimat itu sudah dipasang lead di `ENGINEERING-STANDARDS.md` §8 |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4) |
| Pengusul | `architect` |
| Mengubah | `ENGINEERING-STANDARDS.md` §8 (satu kalimat, butir 3; disetujui pemilik produk) |
| Rincian | [09.1 §2, §5, §13](../09-technical-architecture/01-pilihan-teknologi.md); [ADR-012](ADR-012-ui-native-per-platform-dan-versi-os-minimum.md) butir 5 |

## Konteks

ES §8 menuntut hanya versi stabil terbaru: tidak ada alpha, beta, RC, atau API eksperimental tanpa ADR. Ia tidak menyebut versi 0.x.

Banyak pustaka yang matang dan banyak dipakai masih bernomor 0.x menurut semver, artinya API publiknya boleh berubah tanpa naik versi mayor. Yang diperlukan atau disebut dokumen 09 (diperiksa 12 November 2026; 18 §5 fase D) **[Terverifikasi]**:

- `golang.org/x/crypto` v0.59.0 dan `golang.org/x/oauth2` v0.38.0, dari proyek Go sendiri, diperlukan di server Go;
- connect-kotlin **0.10.0**, diperlukan klien Android (dikukuhkan sebagai rilis terbaru yang bukan pra-rilis di [github.com/connectrpc/connect-kotlin/releases](https://github.com/connectrpc/connect-kotlin/releases), diakses 12 November 2026);
- axum 0.8.9 dan sqlx 0.9.0, hanya bila Rust dipilih (tidak dipilih; ADR-015);
- River v0.51.2 (alternatif antrean, tidak dipakai);
- SwiftLint 0.65.1 dan ktfmt 0.65 (dihindari ADR-025).

Selain itu, artefak stabil dapat memuat **API bertanda eksperimental**. Compose Material 3 1.4.0 memuat `ModalBottomSheet` dengan `@ExperimentalMaterial3Api`, dan Material 3 Adaptive 1.3.0 (12 Agustus 2026) memuat `ListDetailPaneScaffold` dengan `@ExperimentalMaterial3AdaptiveApi`. `SwipeToDismissBox` dan `NavigationSuiteScaffold` tidak eksperimental di 1.4.0, sehingga tidak perlu dicatat di sini (18 §5.1.1, §5.1.3; indeks 09 §9). ES §8 menyebut API eksperimental, sehingga pemakaiannya memerlukan ADR.

Tanpa kebijakan, setiap butir ini menjadi perdebatan baru, atau ES §8 dibaca berbeda oleh agen build yang berbeda.

## Opsi

| Opsi | Akibat |
|---|---|
| A. 0.x dan API eksperimental selalu dianggap tidak stabil | Menolak `golang.org/x/crypto` (Argon2id tidak ada di stdlib), connect-kotlin, dan `ModalBottomSheet`; Rust praktis tertutup |
| B. 0.x selalu dianggap stabil bila bukan pra-rilis | Tidak ada pagar atas pustaka muda yang berubah tiap bulan |
| **C. Boleh dengan syarat dan tercatat dengan jalan mundurnya** | Pagar yang dapat diperiksa, tanpa menutup pustaka matang |

## Keputusan

1. **Rilis 0.x terbaru yang bukan pra-rilis** (bukan alpha, beta, RC, atau nightly) **boleh dipakai bila kelima syarat ini terpenuhi**:
   1. tidak ada alternatif ≥ 1.0 yang setara untuk kebutuhan yang sama **tanpa menambah kode atau komponen operasi milik kita** (misalnya implementasi protokol buatan sendiri, atau proses dan timer tambahan di VPS). Alternatif ≥ 1.0 yang menuntut tambahan seperti itu tidak menggugurkan pustaka 0.x; ia menjadi jalan mundurnya (syarat 5). Rumusan ini disempitkan pada revisi fase D (CR-A-05), karena untuk `golang.org/x/oauth2` (stdlib `crypto/rsa` + `net/http`) dan `acme/autocert` (certbot) alternatif ≥ 1.0 memang ada, tetapi masing-masing menambah kode buatan sendiri atau komponen operasi;
   2. proyeknya aktif: rilis dalam 6 bulan terakhir, dan isu keamanan ditanggapi;
   3. versinya dipatok tepat di lockfile, dan pembaruan minor diperlakukan sebagai pembaruan yang dapat merusak (dibaca catatan rilisnya, lalu seluruh uji dijalankan);
   4. pemakaiannya dibungkus satu modul atau lapisan tipis milik kita, sehingga penggantiannya terbatas;
   5. ada jalan mundur yang dinyatakan.
2. **Daftar pustaka 0.x yang disetujui ADR ini** (versi persis dicek saat dipasang):

   | Pustaka | Versi terlihat | Dipakai untuk | Jalan mundur |
   |---|---|---|---|
   | `golang.org/x/crypto/argon2` | v0.59.x | Hash kata sandi Argon2id (SEC-11) | Tidak ada padanan di stdlib. Implementasi Argon2id lain yang ≥ 1.0 bila ada saat itu, atau salinan paket `argon2` yang dipatok di repositori (lisensi BSD) **[Perlu validasi]** |
   | `golang.org/x/crypto/acme/autocert` | v0.59.x | Sertifikat TLS | certbot di VPS dan biner membaca berkas sertifikat (09.4 §1). pkg.go.dev menyebut `autocert` mungkin tidak ada di versi modul terbaru; dicek saat dipasang |
   | `golang.org/x/oauth2` (`google`) | v0.38.x | Token akses akun layanan FCM HTTP v1 | JWT akun layanan ditandatangani dengan `crypto/rsa` stdlib dan ditukar ke titik token Google lewat `net/http` |
   | connect-kotlin | 0.10.x | Klien RPC Android | Pesan protobuf hasil generate resmi + klien HTTP biasa (unary Connect = HTTP POST) |
   | axum, sqlx | 0.8.x, 0.9.x | **Tidak dipakai**: Rust tidak dipilih (ADR-015) | actix-web v4 (stabil); Diesel 2.x (stabil) |
   | River | 0.51.x | **Tidak dipakai**; alternatif antrean | Tabel `job` sendiri (rekomendasi) |

   Pustaka 0.x lain memerlukan baris baru di tabel ini lewat revisi ADR ini.
3. **Kalimat ES §8** (disetujui pemilik produk, S4-1 (iii) A; dipasang lead): *"A 0.x release that is not a pre-release, or an API marked experimental inside a stable release, counts as stable only if it is listed in ADR-023 with its fallback; anything else needs its own ADR."*
4. **Pra-rilis tetap dilarang** tanpa ADR pengecualian yang menyebut versi, alasan, dan jalan keluarnya, misalnya Compose Material 3 1.5.0-rc01 dengan komponen Expressive (ADR-012 butir 5, yang memilih tidak memakainya).
5. **API eksperimental di dalam artefak stabil** diperlakukan sama dengan pustaka 0.x: kelima syarat butir 1 berlaku (syarat 4 berarti satu composable pembungkus milik kita), dan setiap API tercatat di sini.

   | API | Artefak stabil | Dipakai untuk | Jalan mundur |
   |---|---|---|---|
   | `ModalBottomSheet`, `rememberModalBottomSheetState` (`@ExperimentalMaterial3Api`) | Compose Material 3 1.4.0 | Lembar Catat di Android pilot (dokumen 08; ADR-012 butir 5 (b)) | Dialog layar penuh dengan `Dialog` dan `Scaffold` stabil, hanya bila API itu rusak atau berubah saat naik versi. Catatan naik ke 1.5: `rememberModalBottomSheetState` sudah `@Deprecated` di kode pengembangan (18 §5.1.1) |
   | `ListDetailPaneScaffold` (`@ExperimentalMaterial3AdaptiveApi`) | Material 3 Adaptive 1.3.0 | **Tidak dipakai di pilot** (tidak ada tablet Android) | Tata letak dua panel dengan kelas ukuran jendela stabil |

## Konsekuensi

- Agen build punya satu daftar yang dapat diperiksa, bukan tafsiran.
- Pemutakhiran bulanan (ES §8) atas pustaka 0.x dan API eksperimental memerlukan pembacaan catatan rilis. Biayanya lebih mahal daripada pustaka 1.x, dan biaya itu diterima untuk daftar pendek ini (RISK-02, RISK-17).
- Bila sebuah pustaka di daftar mencapai 1.0, atau sebuah API kehilangan tanda eksperimentalnya, barisnya dihapus pada pemutakhiran berikutnya.
