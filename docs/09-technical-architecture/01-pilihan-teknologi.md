# 09.1 — Pilihan teknologi: matriks perbandingan

Bagian dari [dokumen 09](../09-technical-architecture.md). Berkas ini membandingkan opsi untuk setiap keputusan teknologi dan merekomendasikan satu. Keputusannya dicatat sebagai ADR-012 sampai ADR-026. Fakta riset dikutip dari [dokumen 18](../18-market-research.md) §5 (diakses 12 November 2026). Fakta yang tidak ada di 18 diperiksa langsung pada 12 November 2026 dan dicantumkan di [indeks §9](../09-technical-architecture.md#9-fakta-yang-diperiksa-untuk-dokumen-ini).

---

## 1. Batasan dan kriteria

Batasannya berasal dari brief §1, dan setiap pilihan di berkas ini harus lolos semuanya. Pengenal `BT-N` hanya dipakai di dokumen 09.

| # | Batasan | Sumber | Akibat bagi pilihan |
|---|---|---|---|
| BT-1 | Tim satu sampai tiga pengembang. Pemimpinnya kuat di Go dan Rust dan berpengalaman luas di mobile | Brief §1 | Setiap komponen tambahan menambah beban pemeliharaan bagi orang yang sama |
| BT-2 | Pilot berupa satu rumah tangga dua orang di satu VPS sederhana | Brief §1 | Tidak dirancang untuk lalu lintas yang tidak ada; tidak ada klaster |
| BT-3 | Pengembangan di Linux, tanpa Mac | Brief §1 | Setiap jalur iOS harus menyebut cara build, uji di iPhone, dan distribusi, beserta biayanya |
| BT-4 | Pilot memerlukan iOS dan Android (iPhone 13, Galaxy A55, iPad mini 7) | Brief §1 tabel | Kedua platform siap sebelum pilot |
| BT-5 | "No usage-priced managed services unless clearly justified"; penyedia `AI` adalah pengecualian yang diterima, di bawah batas biaya keras | Brief §1 tabel | Tidak ada BaaS, antrean terkelola, atau object storage per GB bila ada alternatif berharga tetap; layanan per pemakaian hanya dengan alasan yang dinyatakan (email, ADR-026) |
| BT-6 | Server aplikasi di-host di Indonesia; `AI` diinginkan Claude atau ChatGPT | Keputusan 10 (dokumen 17 bagian keputusan S1), "Ia menginginkan AI-nya Claude atau ChatGPT, dan server aplikasi di-host di Indonesia" | Hosting dan cadangan di Indonesia; jalur `AI` mengikuti 10-B |
| BT-7 | Tampilan dan rasa native mengalahkan berbagi kode | Brief §9.1, §10.1 | Framework yang menggambar widget sendiri ditolak kecuali terbukti |
| BT-8 | Satu implementasi aturan yang berwenang, dan server sebagai otoritas akhir | `ENGINEERING-STANDARDS.md` §2 | Tidak ada salinan kedua aturan uang atau persetujuan yang ditulis tangan; setiap penyimpangan memerlukan ADR yang disetujui pemilik produk (ES baris 3) |
| BT-9 | Hanya versi stabil terbaru; alpha, beta, RC, atau API eksperimental memerlukan ADR | `ENGINEERING-STANDARDS.md` §8 | Pustaka 0.x dan API eksperimental dinilai menurut [ADR-023](../adr/ADR-023-kebijakan-versi-0x.md) |

Kriteria umum untuk backend mengikuti brief §9.2: kinerja API, waktu nyata, integritas transaksi, pekerjaan latar, keamanan, konkurensi, produktivitas, beban operasi, skalabilitas, kemudahan pemeliharaan, dan biaya. BT-7 dan BT-8 adalah **gerbang**. Opsi yang gagal di gerbang tidak dinilai lebih lanjut.

---

## 2. Stack mobile (ringkasan; keputusan pemilik produk)

Matriks lengkap ada di 18 §5.1. Tabel ini hanya merangkum baris yang menentukan.

| Opsi | Lolos gerbang native di iOS | Lolos gerbang native di Android | Catatan dari 18 §5.1 |
|---|---|---|---|
| Flutter | Tidak: menggambar widget sendiri | Tidak: menggambar widget sendiri | Gaya Liquid Glass harus ditiru setiap tahun |
| React Native / Expo | Ya (komponen UIKit) | Belum terbukti untuk Material 3 Expressive | Lapisan JS tambahan |
| Compose Multiplatform | Tidak: menggambar sendiri di iOS | Ya | Sama dengan Flutter di sisi iOS |
| **Native penuh (SwiftUI dan Jetpack Compose), logika di server** | **Ya** | **Ya** | Liquid Glass otomatis dengan SDK 26. Opt-out `UIDesignRequiresCompatibility` diabaikan mulai SDK 27. Unggahan ke App Store wajib memakai Xcode 26 / SDK iOS 26 sejak 28 April 2026 |
| Native penuh dengan modul logika KMP | Ya | Ya | Swift export masih alpha; contoh Room KMP memakai 3.1.0-alpha03 |
| Native penuh dengan inti Rust lewat UniFFI | Ya | Ya | UniFFI pra-1.0; Element X memakai pola ini, tetapi FFI Android-nya berlabel eksperimental |

**Keputusan pemilik produk (12 November 2026):** "Swift + Kotlin saya setuju untuk apple dan android catat dulu yaa, tinggal backend." ([_progress.md](../_progress.md); dokumen 17 Q-34). Artinya UI native penuh: SwiftUI di iPhone dan iPad, dan Jetpack Compose dengan Material 3 di Android. Keputusan ini dicatat di [ADR-012](../adr/ADR-012-ui-native-per-platform-dan-versi-os-minimum.md) dan tidak dibuka ulang. Yang masih terbuka adalah letak logika bersama (§3).

**Akibat bagi ES §8 (Android).** Compose Material 3 stabil adalah **1.4.0**; **1.5.0-rc01** (4 November 2026) masih release candidate, tanpa beta atau alpha yang lebih baru ([developer.android.com/jetpack/androidx/releases/compose-material3](https://developer.android.com/jetpack/androidx/releases/compose-material3), diakses 12 November 2026) **[Terverifikasi]**. RC adalah pra-rilis menurut ES §8, dan komponen Expressive tidak ada di 1.4.0. Di dalam artefak stabil pun ada API yang masih bertanda eksperimental: `ModalBottomSheet` (`@ExperimentalMaterial3Api`, Material 3 1.4.0) dan `ListDetailPaneScaffold` (`@ExperimentalMaterial3AdaptiveApi`, Adaptive stabil 1.3.0, 12 Agustus 2026) **[Terverifikasi]** di kode stabil (18 §5.1.1, §5.1.3). `SwipeToDismissBox` dan `NavigationSuiteScaffold` **tidak** eksperimental di sumber 1.4.0 **[Terverifikasi]**. `rememberModalBottomSheetState` sudah `@Deprecated` di kode pengembangan, jadi perlu diganti saat naik ke 1.5. Cara menanganinya: ADR-012 butir 5 dan ADR-023 butir 5, ditanyakan di 09-G.

---

## 3. Letak logika bersama

Pertanyaannya: aturan apa yang perlu berjalan di perangkat saat offline, dan di mana kodenya hidup, tanpa membuat salinan kedua yang ditulis tangan (BT-8)?

**Inventaris komputasi yang mungkin berjalan di perangkat.** Dokumen 06 §12 dan 07 sudah mengecilkannya; dokumen 05 menambah dua bantuan isian:

| Hal | Di mana | Sumber |
|---|---|---|
| Transisi state permintaan dan persetujuan | Server saja; perangkat menampilkan "menunggu terkirim" | 06 S-1 |
| Tanda ambang, penilaian jatah, tanda tautan, kecocokan kategori | Server saja, saat entri diterima | 06 S-8; INV-A37; 07 INV-37 |
| Selisih rekonsiliasi | Server saja | 06 S-10; 07 INV-16 |
| Saldo, sisa anggaran, kekayaan bersih | Server saja; perangkat menampilkan angka terakhir server beserta waktunya | 06 S-9; 07 INV-04 |
| Perkiraan "angka server − entri yang belum terkirim" | **Tidak dihitung di pilot** **[Usulan]**. FR-FIN-010/AC-6 dan FR-FIN-091/AC-2 hanya mengaturnya "bila ditampilkan"; yang wajib adalah angka server dengan waktunya dan entri yang belum terkirim tampil terpisah | 06 S-9 (revisi ketiga belas) |
| (a) Selisih baris sebab terhadap jumlah saat mengisi; simpan tidak tersedia bila tidak sama | **Perangkat**, bantuan isian; server tetap menolak (07 INV-02) | FR-FIN-042/AC-2 |
| (b) "Bagi rata": sisa pembagian pada baris pertama | **Perangkat**, bantuan isian; server tidak menghitungnya dan menerima pembagian apa pun yang berimbang | FR-FIN-042/AC-3 |
| (c) Bentuk masukan: wajib diisi, bilangan bulat positif, panjang teks | **Perangkat**, agar perintah offline tidak baru ditolak berjam-jam kemudian; server memeriksanya ulang | ES §2, §4 |

**Bentuk masukan tidak dapat di-generate dari kontrak untuk kedua klien.** protovalidate (aturan validasi sebagai anotasi di berkas `.proto`) punya pustaka resmi untuk Go, Java, Python, C++, dan TypeScript/JavaScript, tanpa Swift dan tanpa Kotlin ([protovalidate.com](https://protovalidate.com), diakses 12 November 2026) **[Terverifikasi]**. Usulannya **[Usulan]**: aturan bentuk ditulis **satu kali** sebagai anotasi protovalidate di kontrak, dan server menegakkannya dengan `buf.build/go/protovalidate` (v1.4.0 terlihat di pkg.go.dev, 31 Agustus 2026; versi terbaru dicek saat dipasang). Klien tidak dapat membaca anotasi itu lewat pustaka resmi, sehingga butir (a) sampai (c) di perangkat **ditulis tangan per platform**.

Jadi ada satu pertanyaan yang hanya dapat dijawab pemilik produk: tiga bantuan isian kecil itu ditulis tangan di Swift dan Kotlin. Menurut bunyi ES §2 ("never a hand-written copy") hal itu **pengecualian ES §2**, walaupun butir (b) tidak punya pasangan di server dan butir (a) serta (c) diperiksa ulang server. ES baris 3 menuntut ADR yang disetujui pemilik produk. Pengecualian itu diajukan di ADR-013 dan ditanyakan terang di 09-F.

| Kriteria | **A. Server saja + tiga bantuan isian berpagar (pengecualian ES §2)** | B. Server saja, tanpa bantuan isian di perangkat | C. Modul KMP bersama | D. Inti Rust lewat UniFFI |
|---|---|---|---|---|
| Satu implementasi aturan uang dan persetujuan (BT-8) | **Ya**: semua aturan di server | Ya | Bantuan isian dibagi Swift dan Kotlin, tetapi server Go tetap punya implementasi sendiri (tidak "same implementation", kecuali server Kotlin) | Satu implementasi hanya bila server juga Rust (09-A = B) |
| Memenuhi bunyi ES §2 tanpa pengecualian | Tidak: tiga bantuan isian ditulis dua kali, dijaga vektor uji emas | **Ya** | Tidak (lihat baris atas) | Ya bila server Rust |
| Akibat bagi pengguna | Isian tidak dapat disimpan bila barisnya tidak berimbang; "bagi rata" tersedia; kesalahan bentuk diketahui saat itu juga | Perintah offline yang salah bentuk baru ditolak saat tersambung dan dikembalikan ke pencatat; FR-FIN-042/AC-2 dan AC-3 harus diubah | Sama dengan A | Sama dengan A |
| Build iOS dari Linux (BT-3) | Tidak ada kode bersama yang dikompilasi untuk iOS | Sama | **Tidak bisa**: target Apple Kotlin/Native hanya didukung di host macOS dan memerlukan Xcode ([kotlinlang.org/docs/native-target-support.html](https://kotlinlang.org/docs/native-target-support.html), diakses 12 November 2026) **[Terverifikasi]** | Kompilasi silang ke `aarch64-apple-ios` memerlukan SDK Apple untuk linker, sehingga menghadapi soal lisensi yang sama dengan xtool (17 L-31) **[Perlu validasi]** |
| Kematangan (BT-9) | Tidak ada pustaka tambahan | Sama | Swift export alpha (18 §5.1) | UniFFI pra-1.0 (18 §5.1) |
| Kecocokan dengan tim (BT-1) | Tinggi | Tinggi | Sedang | Tinggi untuk Rust, tetapi menambah lapisan FFI |

**Rekomendasi: A** **[Usulan]**, bila pemilik produk menyetujui pengecualiannya; bila tidak, B. Alasannya:

1. Tidak ada aturan uang atau persetujuan di perangkat. Yang tersisa adalah tiga bantuan isian yang masing-masing beberapa baris, diperiksa ulang server, atau (butir b) tidak punya akibat di server.
2. C dan D menambah satu bahasa dan satu langkah yang hanya dapat dikompilasi di macOS demi kode yang hampir kosong, dan C tetap tidak memenuhi "same implementation" karena server Go.
3. Pagarnya dapat diperiksa: **vektor uji emas**. Uji server Go menghasilkan berkas data (masukan dan keluaran yang diharapkan) dari implementasi acuan server (INV-02 dan aturan protovalidate) untuk butir (a) dan (c), dan dari satu fungsi acuan Go untuk (b). Uji unit Swift dan Kotlin di CI menjalankan berkas yang sama; bila hasilnya berbeda, merge diblokir.

Keputusannya dicatat di [ADR-013](../adr/ADR-013-logika-bersama-hanya-di-server.md).

**Perkiraan offline kelak.** Bila pilot menunjukkan perkiraan diperlukan, ADR baru menambahkannya sebagai deskriptor yang dikirim server (perintah mana mengubah angka mana, dengan tanda) dan satu penafsir per platform. Itu menambah satu butir pada daftar pengecualian ES §2 dan karena itu juga memerlukan persetujuan pemilik produk.

---

## 4. Kontrak API

ES §2 mewajibkan kontrak API didefinisikan satu kali dan tipe klien di-generate darinya. Kriteria tambahan dari ADR-008: jumlah int64 harus sampai di klien tanpa kehilangan presisi.

| Kriteria | OpenAPI 3.1 + generator per bahasa | **Protobuf + protokol Connect (buf)** | GraphQL | JSON tulis tangan |
|---|---|---|---|---|
| Didefinisikan sekali, tipe di-generate (ES §2) | Ya | Ya | Ya | **Tidak** (gagal gerbang) |
| int64 tanpa kehilangan (ADR-008) | Angka JSON; aman bila decoder Swift dan Kotlin terbukti lossless **[Perlu validasi]** | Tipe `int64` native di Go, Swift, dan Kotlin; bentuk JSON-nya string, sehingga lossless | Skalar khusus | — |
| Kompatibilitas aplikasi lama di tangan pengguna | Bergantung disiplin; nilai enum baru dapat merusak decoder | **Nomor medan dan nilai enum tak dikenal dirancang untuk kompatibel**; `buf breaking` memeriksanya di CI | Skema tunggal; versi lewat deprecasi | — |
| Bingkai waktu nyata dalam skema yang sama | Perlu AsyncAPI atau skema terpisah | Pesan bingkai SSE didefinisikan di berkas `.proto` yang sama | Subscription (lapisan tambahan) | — |
| Aturan validasi dan kode galat sekali (ES §2) | JSON Schema; generator klien beragam | Anotasi protovalidate (penegakan server); kode galat sebagai `enum` di kontrak, ter-generate untuk ketiga bahasa | Bergantung | — |
| Server Go | Generator pihak ketiga | connect-go v2.0.2 (18 §5.3) **[Terverifikasi]** | Pustaka pihak ketiga | — |
| Klien Swift | swift-openapi-generator (Apple) **[Perlu validasi]** versi | connect-swift **1.2.3** (18 §5, fase D) **[Terverifikasi]** | Apollo iOS (pihak ketiga) | — |
| Klien Kotlin | openapi-generator (komunitas) | connect-kotlin **0.10.0**, 0.x (18 §5, fase D) **[Terverifikasi]** | Apollo Kotlin (pihak ketiga) | — |
| Dapat dicoba dengan `curl` | Ya | Ya: unary Connect adalah HTTP POST dengan JSON atau protobuf | Ya | Ya |

**Rekomendasi: Protobuf + Connect** **[Usulan]**, dikelola dengan `buf`: lint, `buf breaking` terhadap cabang utama, dan generate untuk Go, Swift, dan Kotlin. Kode galat domain (misalnya `state_changed`, `terms_superseded`, `currency_locked`, galat per medan) adalah satu `enum` di kontrak, sehingga ketiga bahasa memakai daftar yang sama (ES §2).

connect-kotlin masih 0.x, sehingga ADR-023 berlaku. Jalan mundurnya murah: unary Connect hanyalah HTTP POST ke `/<paket>.<Layanan>/<Metode>` dengan isi protobuf. Aplikasi Android dapat memakai pesan protobuf hasil generate resmi bersama klien HTTP biasa tanpa connect-kotlin, dan kontraknya tidak berubah. Rinciannya di ADR-013 butir 4.

---

## 5. Backend: bahasa, framework, atau BaaS

Fakta dari 18 §5.3 dan §5.4 **[Terverifikasi]**, kecuali yang diberi label lain.

| Kriteria (brief §9.2) | **Go** | Rust | Supabase self-hosted | Firebase | PocketBase |
|---|---|---|---|---|---|
| Versi stabil terbaru (ES §8) | Go 1.27.4 (10 November 2026); routing `net/http` stdlib sejak 1.22; pgx v5.13.0; sqlc v1.31.1; goose v3.30.0; connect-go v2.0.2 | Rust 1.99.1 (1.100 per releases.rs **[Perlu validasi]**); axum **0.8.9**; sqlx **0.9.0**; actix-web v4.19.0; SeaORM 2.0.4; Diesel 2.3.14 | Minimal 4 GB RAM dan 11 komponen; Edge Runtime JS/TS | Firestore tersedia di Jakarta | **Diarsipkan** 26 Oktober 2026 |
| Integritas transaksi (Serializable + coba ulang) | pgx: kendali transaksi penuh; SQL eksplisit lewat sqlc | sqlx: setara | PostgreSQL di baliknya, tetapi logika di Edge Functions JS/TS atau di SQL | **Gagal**: transaksi gagal saat offline; last-write-wins (18 §5.4) | — |
| Satu implementasi aturan (BT-8) | Aturan di Go | Aturan di Rust | Aturan tersebar di SQL, RLS, dan fungsi JS | Aturan di klien atau Cloud Functions | — |
| Pekerjaan latar | Tabel `job` sendiri (§9 berkas 03) atau River v0.51.2 (0.x) | apalis 1.1.1 (1.0.0 baru 12 Oktober 2026) | pg_cron dan Edge Functions | Cloud Functions per pemakaian | — |
| Waktu nyata (SSE) | `net/http` + `http.Flusher`, tanpa dependensi | axum SSE (0.x) | Realtime bawaan (komponen tambahan) | Bawaan | — |
| Keamanan dan permukaan serangan | Satu biner; dependensi sedikit | Satu biner; dependensi lebih banyak **[Asumsi]** | 11 komponen untuk dijaga | Data dan aturan di luar Indonesia sebagian | — |
| Memori di VPS 4 GB | `GOMEMLIMIT`; tidak ada sumber primer pembanding (18 §5.3) | Tidak ada sumber primer pembanding | ≥ 4 GB untuk dirinya sendiri | — | — |
| Produktivitas tim (BT-1) | Tinggi; kompilasi cepat **[Asumsi]** | Tinggi bagi pemimpin; kurva belajar lebih curam bagi anggota kedua dan ketiga **[Asumsi]**; kompilasi lebih lambat **[Asumsi]**, tanpa pengukuran untuk proyek ini | Sedang | Sedang | — |
| Biaya (BT-5) | Termasuk VPS | Termasuk VPS | Termasuk VPS, tetapi butuh VPS lebih besar | **Per operasi** (gagal BT-5) | — |
| Pustaka 0.x yang dibutuhkan | **`golang.org/x/crypto` v0.59.0** (Argon2id, ACME) dan **`golang.org/x/oauth2` v0.38.0** (token FCM), keduanya dari proyek Go sendiri, di tepi sistem (hash kata sandi, sertifikat, token push) **[Terverifikasi]** (18 §5, fase D; indeks §9) | **axum** dan **sqlx** di inti (setiap permintaan HTTP dan setiap kueri) | — | — | — |

**Rekomendasi: Go** **[Usulan]**. Pembandingnya adalah Rust, dan keduanya layak dipakai. Keduanya memerlukan pustaka 0.x menurut ADR-023; yang berbeda adalah letak dan jenisnya.

- **Kelebihan Go di sini:**
  - Pustaka di jalur setiap permintaan (`net/http` stdlib, connect-go, pgx, sqlc, goose) semuanya ≥ 1.0. Pustaka 0.x yang dibutuhkan (`golang.org/x/crypto`, `golang.org/x/oauth2`) dipelihara proyek Go sendiri dan hanya dipakai di tepi sistem; mengganti salah satunya menyentuh satu paket.
  - SSE dan TLS cukup dengan stdlib dan `golang.org/x/crypto/acme/autocert`.
  - Anggota tim kedua dan ketiga lebih cepat produktif **[Asumsi]**.
- **Kelebihan Rust di sini:** jaminan tipe yang lebih kuat untuk aturan uang, dan kemungkinan memakai inti yang sama di perangkat (§3 opsi D) bila kelak dibutuhkan. Harganya: framework HTTP (axum) dan akses basis data (sqlx) sama-sama 0.x di jalur setiap permintaan; apalis baru mencapai 1.0; kompilasi lebih lambat **[Asumsi]**.
- **Yang tidak membedakan:** kinerja. Beban pilot sangat kecil (dua orang, puluhan perintah per hari **[Asumsi]**), dan target p95 ES §3 tercapai dengan keduanya bila kueri berindeks.

**Catatan versi `golang.org/x`.** pkg.go.dev menampilkan v0.59.0 dan v0.38.0 dengan peringatan bahwa halaman paket itu "not in the latest version of its module", dan `autocert` disebut tidak ada di versi modul terbaru menurut halaman itu, yang bertentangan dengan indeks proxy (18 §5, fase D). Versinya dicek ulang saat dipasang (ES §8); bila `acme/autocert` benar-benar dipindah atau dihentikan, sertifikat diambil dengan klien ACME lain yang dicatat di ADR-023, atau dengan certbot di luar biner **[Perlu validasi]**.

**BaaS tidak disarankan.** Supabase hosted tidak punya wilayah Indonesia, dan versi self-hosted terlalu berat untuk VPS 4 GB. Firebase gagal integritas transaksi offline dan BT-5. PocketBase sudah diarsipkan (18 §5.4).

Keputusannya dicatat di [ADR-015](../adr/ADR-015-bahasa-dan-framework-backend.md), dan akan diputuskan pemilik produk di checkpoint S4 (pertanyaan 09-A).

---

## 6. Monolit modular atau microservices

| Kriteria | **Monolit modular, satu proses** | Microservices |
|---|---|---|
| Transaksi lintas modul (entri + tanda + Peristiwa + notifikasi dalam satu transaksi; 06 S-4) | **Satu transaksi PostgreSQL** | Saga atau outbox lintas layanan; kompensasi untuk aturan uang |
| Beban operasi di satu VPS (BT-2) | Satu biner, satu unit systemd | Banyak proses, jaringan internal, penemuan layanan |
| Satu jalur baca (SEC-06) | Satu fungsi otorisasi baca dipakai semua modul | Harus disalin atau dipanggil lewat jaringan |
| Skala tahap publik | Proses dapat digandakan di belakang penyeimbang beban; PostgreSQL dipisah ke server sendiri | Skala per layanan |
| Tim (BT-1) | Cocok | Tidak sebanding dengan bebannya |

**Rekomendasi: monolit modular** **[Usulan]**. Batas modul dijaga paket `internal/` Go dan satu uji arsitektur di CI yang menolak impor lintas modul di luar daftar yang diizinkan. Rinciannya di [ADR-014](../adr/ADR-014-monolit-modular.md) dan berkas 02 §1.

---

## 7. Basis data

| Kriteria | **PostgreSQL 18** | SQLite di server (+ replikasi berkas) | MySQL | Firestore |
|---|---|---|---|---|
| Versi | 18.6, didukung sampai 14 November 2030; 19 masih RC1 (18 §5.5) **[Terverifikasi]** | — | Tidak diteliti | — |
| Serializable sejati dengan coba ulang 40001 | **Ya** (SSI) | Satu penulis; serial secara alami | Serializable berbasis kunci | Tidak |
| RLS sebagai lapis kedua (SEC-08) | **Ya**, dengan `FORCE ROW LEVEL SECURITY` | Tidak ada | Tidak ada | Aturan keamanan sendiri |
| PITR (SEC-43) | Base backup + WAL (pgBackRest, WAL-G) | Lewat alat replikasi berkas | Binlog | Terkelola |
| Pencarian bahasa Indonesia | Konfigurasi `indonesian` (Snowball, sejak PG12; mutu terbatas) | FTS5 tanpa stemmer Indonesia **[Asumsi]** | Terbatas | Tidak |
| Jalur ke tahap publik | Server sendiri dan replika | Batas satu penulis | — | Per operasi (BT-5) |

**Rekomendasi: PostgreSQL 18** **[Usulan]**. Rinciannya di [ADR-016](../adr/ADR-016-postgresql-tenancy-dan-transaksi.md) dan berkas 02.

---

## 8. Transport waktu nyata

Arah dari 18 §5.6: kanal waktu nyata saat aplikasi di depan, ditambah tarik berkursor sebagai sumber kebenaran, ditambah push langsung. 06 S-11 (e) menuntut aturan tetap benar bila hanya push disusul tarik yang tersedia.

| Kriteria | WebSocket (RFC 6455) | **SSE + tarik berkursor** | Long polling | Hanya push disusul tarik |
|---|---|---|---|---|
| Arah data yang dibutuhkan | Dua arah | Server ke klien; klien ke server lewat RPC biasa | Server ke klien | Server ke klien lewat Apple dan Google |
| Kebutuhan dua arah sungguhan? | Tidak: semua tulisan adalah perintah idempoten (06 S-2) yang lewat RPC | — | — | — |
| Dependensi server Go | Pustaka WebSocket pihak ketiga | **Tidak ada** (`net/http` + `Flusher`) | Tidak ada | APNs dan FCM |
| Melanjutkan setelah putus | Dirancang sendiri | `Last-Event-ID` bawaan; komentar setiap 15 detik menjaga sambungan (18 §5.6) | Kursor | Kursor |
| iOS saat aplikasi ditangguhkan | Soket dapat ditutup sistem (TN2277) | Sama: hanya dipakai saat aplikasi di depan | Sama | Tetap jalan (bila tidak diblokir) |
| Klien | `URLSessionWebSocketTask`; OkHttp | Pembacaan aliran `URLSession`; OkHttp; pengurainya pendek | Biasa | Biasa |
| Latensi pesan chat | Satu perjalanan | Satu dorongan + satu tarik | Satu tarik | Detik, tidak dijamin |

**Rekomendasi: SSE + tarik berkursor saat aplikasi di depan; push langsung saat tidak** **[Usulan]**. Kanal SSE hanya membawa (a) tanda bahwa umpan pembaca bergerak, dan (b) kehadiran yang bersifat sesaat. Isi selalu ditarik lewat satu jalur baca (SEC-06), sehingga hanya ada satu jalan data. Rinciannya di berkas 03 dan [ADR-018](../adr/ADR-018-model-sinkron-dan-waktu-nyata.md).

---

## 9. Push notification

| Kriteria | **Langsung ke APNs dan FCM dari server** | FCM untuk kedua platform | Agregator pihak ketiga |
|---|---|---|---|
| SDK pihak ketiga di aplikasi (SEC-49) | iOS: tidak ada. Android: `firebase-messaging` saja | SDK Firebase juga di iOS | SDK agregator di kedua aplikasi (gagal SEC-49) |
| Isi lewat pihak tambahan | Apple dan Google saja (L-16) | Google juga untuk iOS | Satu pihak lagi |
| Kendali collapse id, kanal, lencana (SEC-52, SEC-53, SEC-57) | Penuh | Sebagian lewat pemetaan FCM | Bergantung penyedia |
| Dependensi server | HTTP/2 stdlib, JWT ES256 (APNs); OAuth2 akun layanan (`golang.org/x/oauth2`, 0.x, ADR-023) untuk FCM HTTP v1 | Satu API | Satu API |

**Rekomendasi: langsung ke APNs dan FCM** **[Usulan]**; dicatat di [ADR-019](../adr/ADR-019-push-langsung-apns-fcm.md).

---

## 10. Hosting dan tujuan cadangan

Syarat: server di Indonesia (BT-6); harga tetap per bulan (BT-5); ADR mencatat badan hukum pihak kontrak dan lokasi pusat data, baik untuk VPS maupun tujuan cadangan (dokumen 10 §5).

| Kriteria | **Biznet Gio NEO Lite MS 4.2** | Biznet Gio NEO Lite Pro MS.4.2 | Akamai (Linode) Jakarta `id-cgk` | AWS Lightsail Jakarta | IDCloudHost |
|---|---|---|---|---|---|
| Spesifikasi dan harga per bulan | 2 vCPU, 4 GB, 60 GB SSD, **Rp139.000** (18 §5.8; diperiksa ulang 12 November 2026) **[Terverifikasi]** | 2 vCPU, 4 GB, 60 GB NVMe, Rp559.000 (harga tahunan dengan kode diskon) | Linode 4GB: USD 28,80 | 4 GB: USD 24 (transfer separuh) | "Rocket VPS 1" 2 core, 4 GB, Rp105.000 **[Perlu validasi]** |
| Pihak kontrak | **PT Biznet Gio Nusantara**, tunduk hukum Indonesia (biznetgio.com/legal; 18 §5, fase D) **[Terverifikasi]** | Sama | Tidak diteliti | PT Amazon Web Services Indonesia **[Terverifikasi]** | **PT Cloud Hosting Indonesia** (hanya dari [idcloudhost.com/llms.txt](https://idcloudhost.com/llms.txt), berkas untuk mesin, bukan syarat layanan) **[Perlu validasi]** |
| Lokasi pusat data | NEO Cloud di Cimanggis dan MidPlaza menurut halaman Biznet Gio sendiri (2024; 18 §5, fase D) **[Terverifikasi]**; yang melayani NEO Lite secara khusus **[Perlu validasi]** | Sama | "Jakarta, ID" | `ap-southeast-3` Jakarta | Pusat data di Indonesia dan Singapura, termasuk Bogor dan Cyber 1 Jakarta (llms.txt) **[Perlu validasi]**; lokasi dipilih saat memesan |
| Data akun di luar negeri | Kebijakan privasi Biznet menyatakan data pribadi "dapat diproses atau disimpan di luar wilayah Republik Indonesia"; cakupan terbaca: data akun pelanggan (18 §5, fase D) **[Terverifikasi]** bunyinya; apakah berlaku atas isi VPS **[Perlu validasi]** advokat | Sama | — | — | Tidak diteliti |
| PPN | Tidak disebut di halaman | Tidak disebut | — | — | — |
| Harga tetap (BT-5) | Ya | Ya | Ya | Ya | Ya |

**Rekomendasi** **[Usulan]**:

- **VPS utama:** Biznet Gio NEO Lite MS 4.2, dengan meminta penyedia menyatakan secara tertulis kota pusat data VPS itu, dan bahwa kalimat "di luar wilayah Republik Indonesia" di kebijakan privasinya tidak mencakup isi VPS pelanggan, sebelum data nyata masuk.
- **Tujuan cadangan:** satu VPS kecil di **penyedia lain di Indonesia**, yaitu IDCloudHost di lokasi Indonesia. Dengan begitu kegagalan atau penangguhan akun di satu penyedia tidak menghapus server dan cadangannya sekaligus.
- **Alternatif yang lebih sederhana:** VPS NEO Lite XS 1.1 kedua di Biznet Gio, Rp59.000 per bulan. Pihak kontraknya sama, sehingga risiko akun tidak terpisah. Pilihan ini muncul di checkpoint (09-B).

Rinciannya di [ADR-020](../adr/ADR-020-hosting-dan-cadangan-di-indonesia.md).

---

## 11. Pendamping `AI`: hosted atau self-hosted (brief §9.5)

| Kriteria | **Hosted (Claude lewat API Anthropic)** | Hosted (ChatGPT lewat API OpenAI) | Self-hosted (SEA-LION v3 9B Q4_K_M di CPU) |
|---|---|---|---|
| Mutu bahasa Indonesia | Global-MMLU Indonesia 97,1% (Haiku 5.5) sampai 98,6% (Sonnet 5.5) dari skor Inggris, klaim vendor (18 §5.9) | Tolok ukur primer tidak ditemukan (18 §5.9) | Belum terbukti; SEA-HELM tidak terbaca (18 §5.9) |
| Lokasi pemrosesan | `global` atau `us` saja | Indonesia tidak tercantum | **Di Indonesia** (VPS sendiri) |
| Kesesuaian dengan R1 (kumpulan kalimat tanpa data rumah tangga; 10-B A) | Ya: tidak ada data rumah tangga yang dikirim | Ya | Ya |
| Harga | Haiku 5.5 USD 0,10/0,50, Sonnet 5.5 USD 2/10 per juta token | gpt-6.1-terra USD 1/5 | VPS 8 sampai 16 GB: Rp269.000 sampai Rp459.000 per bulan, sebaiknya VPS terpisah |
| Satu pembangunan kumpulan (0,06 juta token masuk, 0,12 juta keluar **[Asumsi]**; berkas 03 §11.1) | Sonnet 5.5: 0,06 × 2 + 0,12 × 10 = **USD 1,32** | gpt-6.1-terra: 0,06 × 1 + 0,12 × 5 = **USD 0,66** | Sewa VPS beberapa jam **[Perlu validasi]** tagihan per jam |
| Kecepatan | Detik | Detik | 3 sampai 10 token per detik (sumber sekunder) **[Perlu validasi]** |
| Beban operasi | Kunci API, batas belanja | Sama | Model, runtime, pembaruan, pemantauan mutu |
| Batas belanja di penyedia | Spend limit per workspace, tidak untuk Default Workspace | Tidak diteliti | Tidak relevan |

**Rekomendasi** **[Usulan]**: **hosted, Claude lewat API Anthropic**, di jalur **R1** (bawaan dan saran 10-B A).

- Kumpulan kalimat dibangun oleh pekerjaan terjadwal yang tidak membawa data rumah tangga. Server di Indonesia memilih kalimat untuk tiap orang.
- **Model:** Sonnet 5.5 untuk membangun kumpulan, karena dipanggil jarang dan mutu kalimatnya yang terlihat pengguna. ChatGPT separuh biayanya per pembangunan (USD 0,66 lawan 1,32), tetapi mutu bahasa Indonesianya tanpa tolok ukur primer; pada satu pembangunan per bulan selisihnya USD 0,66.
- **Batas biaya:** USD 10 per bulan, dipasang di dua tempat (SEC-127): spend limit di workspace khusus, dan penghitung di server per pekerjaan pembangunan kumpulan dan per bulan (berkas 03 §11.1).
- **Fallback:** kumpulan yang sudah tersimpan tetap dipakai, sehingga penyedia mati tidak menghentikan saran apa pun. Bila kumpulan kosong, tidak ada saran. Pencatatan, permintaan, dan persetujuan tidak pernah tersentuh (brief §5O).

ChatGPT tetap pilihan bagi pemilik produk (09-D). Akibat pilihan 10-B B, C, dan D ada di berkas 03 §11 dan [ADR-021](../adr/ADR-021-jalur-ai-kumpulan-kalimat.md).

---

## 12. Build, uji, dan distribusi iOS tanpa Mac

Fakta dari 18 §5.2 **[Terverifikasi]**, kecuali yang diberi label lain.

| Jalur | Build | Uji di iPhone | Distribusi | Biaya | Catatan |
|---|---|---|---|---|---|
| **Xcode Cloud + TestFlight, dengan sewa Mac sekali untuk penyiapan** | Di cloud Apple | Pasang dari TestFlight di iPhone 13 dan iPad mini 7 | TestFlight (internal), lalu App Store | Termasuk Apple Developer Program **USD 99/tahun**; 25 jam komputasi per bulan; tambahan 100 jam USD 49,99. Sewa Mac sekali: Scaleway Mac mini M4 sekitar EUR 0,22/jam, minimal 24 jam **[Perlu validasi]** | Alur kerja pertama harus disiapkan di Xcode; artefak disimpan 30 hari |
| GitHub Actions macOS | Runner macOS | TestFlight | Unggah lewat API App Store Connect `buildUploads` | USD 0,062/menit; jatah Free USD 12 (sekitar 193 menit), Pro USD 18 (sekitar 290 menit) | Penandatanganan dikelola sendiri |
| Codemagic | Runner macOS M2 | TestFlight | Sama | 500 menit macOS gratis per bulan; USD 0,095/menit sesudahnya | Penyedia tambahan |
| Mac sewaan bulanan | Mac jarak jauh | Xcode + kabel tidak ada; lewat TestFlight atau jaringan | Sama | Scaleway sekitar EUR 149/bulan **[Perlu validasi]**; AWS EC2 Mac minimal 24 jam | Mengembalikan pratinjau, simulator, dan debugger |
| Mac mini sendiri | Lokal | Kabel, debugger | Sama | Mac mini M4 di iBox Rp12.749.000 **[Perlu validasi]** | Menyalahi "tanpa Mac" di brief §1 |
| xtool di Linux | Linux, memakai isi Xcode `.xip` | Langsung | — | Gratis | **Tidak dijadikan jalur resmi**: lisensi EA2002 hanya mengizinkan perangkat lunak Apple di perangkat keras Apple yang menjalankan macOS (17 L-31) |

**Rekomendasi** **[Usulan]**:

1. **Jalur utama:** Xcode Cloud untuk build, uji unit di simulator, arsip, dan TestFlight. Mac disewa satu kali untuk menyiapkan alur kerja pertama.
2. **Uji di perangkat:** iPhone 13 dan iPad mini 7 lewat TestFlight internal.
3. **Kode non-UI yang diuji di Linux:** jaringan, sinkron, antrean, dan bantuan isian ditulis sebagai paket Swift tanpa SwiftUI, agar dapat dibangun dan diuji di Linux CI **[Perlu validasi]** per dependensi.
4. **Bila pekerjaan UI iOS menjadi harian:** pengembangan SwiftUI tanpa pratinjau, simulator, dan debugger itu lambat (18 §5.2). Pemicu untuk menyewa Mac bulanan atau membeli Mac mini adalah fase UI iOS. Pilihannya ada di checkpoint (09-C).

Rinciannya di [ADR-022](../adr/ADR-022-jalur-build-ios-tanpa-mac.md).

---

## 13. Versi yang disebut dan statusnya menurut ES §8

Semua versi diperiksa ulang secara langsung saat fase bangun dimulai (ES §8). Tabel ini hanya memuat versi yang disebut dokumen ini.

| Komponen | Versi | Sumber | Status ES §8 |
|---|---|---|---|
| Go | 1.27.4 | 18 §5.3 | Stabil |
| pgx | v5.13.0 | 18 §5.3 | Stabil |
| sqlc | v1.31.1 | 18 §5.3 | Stabil |
| goose | v3.30.0 | 18 §5.3 | Stabil |
| connect-go | v2.0.2 | 18 §5.3 | Stabil |
| `golang.org/x/crypto` | v0.59.0 (11 November 2026, menurut pkg.go.dev) | [pkg.go.dev/golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto), diakses 12 November 2026; 18 §5 (fase D) | **0.x**: ADR-023 |
| `golang.org/x/oauth2` | v0.38.0 (6 November 2026, menurut pkg.go.dev) | [pkg.go.dev/golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2), diakses 12 November 2026; 18 §5 (fase D) | **0.x**: ADR-023 |
| protovalidate-go (`buf.build/go/protovalidate`) | v1.4.0 terlihat (31 Agustus 2026); halaman menyatakan bukan versi terbaru modulnya | [pkg.go.dev/buf.build/go/protovalidate](https://pkg.go.dev/buf.build/go/protovalidate), diakses 12 November 2026 | Stabil; versi terbaru dicek saat dipasang |
| PostgreSQL | 18.6 | 18 §5.5 | Stabil; 19 RC1 tidak dipakai |
| pgBackRest | 2.60.0 (18 Oktober 2026, "Current Stable Release") | [pgbackrest.org/release.html](https://pgbackrest.org/release.html), diakses 12 November 2026 | Stabil. Dukungan PG18 diperkenalkan 2.55.0 sebagai eksperimental karena dirilis sebelum beta PG18; catatan rilisnya: "Nothing that affects pgBackRest changed before PostgreSQL 18.0" **[Terverifikasi]**. Status "penuh" tidak dinyatakan **[Perlu validasi]** |
| Compose Material 3 | 1.4.0 stabil; 1.5.0-rc01 (4 November 2026) | [developer.android.com](https://developer.android.com/jetpack/androidx/releases/compose-material3), diakses 12 November 2026 | 1.4.0 stabil; RC tidak dipakai; API `@ExperimentalMaterial3Api` di 1.4.0 hanya lewat ADR-023 butir 5 |
| Compose Material 3 Adaptive | 1.3.0 stabil (12 Agustus 2026); 1.4.0-alpha03 (7 Oktober 2026) | [developer.android.com](https://developer.android.com/jetpack/androidx/releases/compose-material3-adaptive), diakses 12 November 2026 | 1.3.0 stabil; `ListDetailPaneScaffold` eksperimental (ADR-023 butir 5) |
| Xcode / SDK iOS | Stabil terbaru saat bangun; Xcode 27.1 sudah dipakai citra build EAS (18 §5.2). Unggahan wajib minimal Xcode 26 / SDK iOS 26 sejak 28 April 2026 (18 §5.1) | 18 §5.1, §5.2 | Stabil; membangun dengan SDK 27 membuat `UIDesignRequiresCompatibility` diabaikan, yang tidak berpengaruh karena Liquid Glass memang dipakai |
| connect-swift | 1.2.3 | 18 §5 (fase D) | Stabil |
| connect-kotlin | 0.10.0 | 18 §5 (fase D) | **0.x**: ADR-023 |
| Rust (alternatif) | 1.99.1; axum 0.8.9; sqlx 0.9.0 | 18 §5.3 | axum dan sqlx **0.x**: ADR-023 bila 09-A memilih Rust |
| River (alternatif antrean) | v0.51.2 | 18 §5.7 | **0.x**: tidak dipakai di rekomendasi |
| Target API Android | 36 (wajib untuk aplikasi baru dan pembaruan sejak 31 Agustus 2026) | 18 §5 (fase D); [developer.android.com/google/play/requirements/target-sdk](https://developer.android.com/google/play/requirements/target-sdk) | Hanya berlaku bila Play dipakai; dicatat di ADR-012 |
| Alat kualitas per bahasa | Lihat ADR-025 | ADR-025 | Dicek saat dipasang |
