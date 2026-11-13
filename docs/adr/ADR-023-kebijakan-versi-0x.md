# ADR-023 — Kebijakan pustaka versi 0.x dan API eksperimental

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **opsi C** dengan kalimat ES §8 di butir 3 (pertanyaan S4-1 (iii), huruf A). Kata-katanya: "yang penting kan versi stable, versi stable ga musti dari 1 kan? apalagi yang baru? cek goolge aja" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban keempat). Bacaan lead: rilis 0.x final, bukan pra-rilis, boleh dipakai, sama dengan bunyi butir 3; dikukuhkan dengan "lanjut". Kalimat itu sudah dipasang lead di `ENGINEERING-STANDARDS.md` §8 |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4); 13 November 2026 (FS-1 1.2: butir 6, modul tidak langsung); 13 November 2026 (FS-1 1.3: butir 7, SDK Go hasil generate BSR); 14 November 2026 (FS-1 1.6: butir 8, `golang.org/x/text`) |
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
   | `golang.org/x/text/unicode/norm` (catatan 14 November 2026, FS-1 tugas 1.6) | v0.44.x | NFC atas nama rumah tangga dan nama tampilan sebelum panjangnya dihitung (putusan `product-analyst` atas rencana 1.6 Q5 (a)) | Tanpa NFC: panjang dihitung dalam rune sesudah dipangkas, dan klien menormalkan dengan fungsi platformnya sebelum mengirim; batas panjang itu **[Usulan]**, jadi selisih karakter terurai tidak melanggar aturan apa pun |
   | connect-kotlin | 0.10.x | Klien RPC Android | Pesan protobuf hasil generate resmi + klien HTTP biasa (unary Connect = HTTP POST) |
   | axum, sqlx | 0.8.x, 0.9.x | **Tidak dipakai**: Rust tidak dipilih (ADR-015) | actix-web v4 (stabil); Diesel 2.x (stabil) |
   | River | 0.51.x | **Tidak dipakai**; alternatif antrean | Tabel `job` sendiri (rekomendasi) |

   Pustaka 0.x lain memerlukan baris baru di tabel ini lewat revisi ADR ini, kecuali modul tidak langsung menurut butir 6.
3. **Kalimat ES §8** (disetujui pemilik produk, S4-1 (iii) A; dipasang lead): *"A 0.x release that is not a pre-release, or an API marked experimental inside a stable release, counts as stable only if it is listed in ADR-023 with its fallback; anything else needs its own ADR."*
4. **Pra-rilis tetap dilarang** tanpa ADR pengecualian yang menyebut versi, alasan, dan jalan keluarnya, misalnya Compose Material 3 1.5.0-rc01 dengan komponen Expressive (ADR-012 butir 5, yang memilih tidak memakainya).
5. **API eksperimental di dalam artefak stabil** diperlakukan sama dengan pustaka 0.x: kelima syarat butir 1 berlaku (syarat 4 berarti satu composable pembungkus milik kita), dan setiap API tercatat di sini.

   | API | Artefak stabil | Dipakai untuk | Jalan mundur |
   |---|---|---|---|
   | `ModalBottomSheet`, `rememberModalBottomSheetState` (`@ExperimentalMaterial3Api`) | Compose Material 3 1.4.0 | Lembar Catat di Android pilot (dokumen 08; ADR-012 butir 5 (b)) | Dialog layar penuh dengan `Dialog` dan `Scaffold` stabil, hanya bila API itu rusak atau berubah saat naik versi. Catatan naik ke 1.5: `rememberModalBottomSheetState` sudah `@Deprecated` di kode pengembangan (18 §5.1.1) |
   | `ListDetailPaneScaffold` (`@ExperimentalMaterial3AdaptiveApi`) | Material 3 Adaptive 1.3.0 | **Tidak dipakai di pilot** (tidak ada tablet Android) | Tata letak dua panel dengan kelas ukuran jendela stabil |
6. **Modul tidak langsung** (catatan 13 November 2026, FS-1 1.2, keputusan lead atas temuan `code-reviewer` I3). Butir 1 dan 2 berlaku bagi dependensi yang diimpor kode kita. Modul yang hanya ditarik oleh dependensi langsung yang sudah disetujui **diterima tanpa baris sendiri**, dengan ketentuan:
   1. versinya adalah versi yang dipilih resolusi modul dependensi langsung itu (versi terbaru yang diizinkan), dipatok di `go.mod` dan `go.sum` (ES §8 "Pinned");
   2. ia tercakup pemindaian kerentanan CI yang sama dengan dependensi langsung: govulncheck untuk modul Go; OSV-Scanner hanya memindai `pnpm-lock.yaml` (dependensi npm);
   3. PR yang pertama kali membawanya mencantumkan nama, versi, tanggal dicek, tautan sumber, dan dependensi langsung yang menariknya (ES §8 "Recorded");
   4. modul tidak langsung baru yang sendiri bernomor 0.x atau berupa pseudo-version **disebut namanya di PR** untuk diperiksa `code-reviewer`. Syarat "aktif" butir 1.2 dinilai pada dependensi langsungnya, bukan pada modul tidak langsung, karena modul kecil yang sudah selesai dapat lama tanpa rilis. Modul yang bertanda deprecated di `go.mod`-nya atau repositorinya diarsipkan tidak diterima lewat butir ini (ES §8 "Maintained");
   5. jalan mundurnya ikut jalan mundur dependensi langsung yang menariknya;
   6. bila kode kita mengimpor modul itu secara langsung, ia menjadi dependensi langsung dan butir 1 dan 2 berlaku penuh (baris sendiri bila 0.x).

   Yang pertama diterima lewat butir ini (PR #13, ditarik pgx v5 dan goose v3): `github.com/jackc/pgservicefile` (pseudo-version `v0.0.0-20240606120523-5a60cdf6a761`, satu-satunya versi di [pkg.go.dev/github.com/jackc/pgservicefile?tab=versions](https://pkg.go.dev/github.com/jackc/pgservicefile?tab=versions), tanpa tanda deprecated, diakses 13 November 2026) **[Terverifikasi]**, `github.com/mfridman/interpolate` v0.2.0, `github.com/sethvargo/go-retry` v0.5.0, `golang.org/x/sync` v0.25.0, dan `golang.org/x/text` v0.44.0 (versi menurut PR #13).
7. **SDK Go hasil generate BSR** (catatan 13 November 2026, FS-1 1.3, kontrak dasar; putusan lead Q8). Modul Go yang dibangkitkan Buf Schema Registry, misalnya `buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go` yang diimpor kode hasil generate kita, bernomor `v<versi plugin>-<waktu komit modul>-<12 karakter ID komit>.<revisi plugin>`. Dokumentasi Buf menyebut bagian waktu dan komit sebagai "pre-release" dalam sintaks semver, dan revisi plugin sebagai build metadata ([buf.build/docs/bsr/generated-sdks/go](https://buf.build/docs/bsr/generated-sdks/go/), diakses 13 November 2026) **[Terverifikasi]**; versi terbaru modul itu saat dicek `v1.36.12-20260825204119-511051f7f437.2` ([proxy.golang.org, `@latest`](https://proxy.golang.org/buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/@latest), diakses 13 November 2026) **[Terverifikasi]**. Bentuk ini adalah **build beralamat komit** dari modul sumber yang sudah dirilis, bukan pra-rilis (alpha, beta, RC, nightly) dalam arti ES §8 dan butir 4, dan diterima tanpa ADR pengecualian bila:
   1. versi plugin pembangkitnya dan versi modul sumbernya (komit di `buf.lock`) dipatok, dan komit di versi Go sama dengan komit di `buf.lock`;
   2. keduanya dicatat di PR yang memasang atau menaikkannya (ES §8 "Recorded"), bersama label rilis modul sumber yang menunjuk komit itu bila ada;
   3. versi plugin pembangkitnya sendiri bukan pra-rilis dan, bila 0.x, memenuhi butir 1.

8. **`golang.org/x/text` menjadi dependensi langsung** (catatan 14 November 2026, FS-1 tugas 1.6; menurut butir 6.6). Modul ini sudah masuk sebagai modul tidak langsung lewat butir 6 (PR #13); sejak 1.6 kode kita mengimpor `unicode/norm`, jadi butir 1 berlaku penuh dan barisnya ada di tabel butir 2. Kelima syarat: (1) stdlib tidak punya normalisasi Unicode, dan menulisnya sendiri adalah kode tambahan milik kita; (2) aktif: v0.44.0 terbit 9 November 2026 ([proxy.golang.org, `@latest`](https://proxy.golang.org/golang.org/x/text/@latest), diakses 14 November 2026) **[Terverifikasi]**, dari proyek Go sendiri seperti `x/crypto`; (3) dipatok di `go.mod` dan `go.sum`; (4) dipakai hanya lewat satu fungsi penormal isian teks di modul `household`, yang dipakai ulang modul lain bila kelak perlu; (5) jalan mundur di baris tabelnya. Versinya tetap versi yang dipilih resolusi modul bersama dependensi lain yang menariknya.

## Konsekuensi

- Agen build punya satu daftar yang dapat diperiksa, bukan tafsiran.
- Pemutakhiran bulanan (ES §8) atas pustaka 0.x dan API eksperimental memerlukan pembacaan catatan rilis. Biayanya lebih mahal daripada pustaka 1.x, dan biaya itu diterima untuk daftar pendek ini (RISK-02, RISK-17).
- Bila sebuah pustaka di daftar mencapai 1.0, atau sebuah API kehilangan tanda eksperimentalnya, barisnya dihapus pada pemutakhiran berikutnya.
