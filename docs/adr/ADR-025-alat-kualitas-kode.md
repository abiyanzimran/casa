# ADR-025 — Alat kualitas kode per bahasa

| | |
|---|---|
| Status | **Diusulkan** (`architect`, 12 November 2026); akan diputuskan pemilik produk di checkpoint S4 |
| Tanggal | 12 November 2026 (S4 fase D) |
| Pengusul | `architect`, atas temuan `critic` CS-A-07 |
| Mengubah | — |
| Rincian | [09.4 §4, §5](../09-technical-architecture/04-operasi-uji-biaya.md) |

## Konteks

`ENGINEERING-STANDARDS.md` menuntut formatter, linter, dan pemindai kerentanan dependensi untuk setiap bahasa, dipilih lewat ADR di S4, dan menjadikannya gerbang merge (ES baris 3, §4, §7). ES §2 juga menuntut pemeriksaan duplikasi. Ada tiga bahasa: Go (server), Swift (iOS dan iPadOS), dan Kotlin (Android), ditambah kontrak Protobuf.

Fakta yang diperiksa 12 November 2026 **[Terverifikasi]** (sumber di [indeks 09 §9](../09-technical-architecture.md#9-fakta-yang-diperiksa-untuk-dokumen-ini)):

- golangci-lint v2.15.0 (16 Oktober 2026); govulncheck (`golang.org/x/vuln`) v1.9.0 (11 Oktober 2026).
- swift-format ikut toolchain Swift 6 ke atas sebagai `swift format`, termasuk mode `lint`. SwiftLint 0.65.1 masih 0.x.
- ktlint 1.8.0 (≥ 1.0), bertanggal "14 Nov" tanpa tahun, jadi November 2025 atau lebih awal; ktlint 2.0 alpha sudah ada dan tidak dipakai (ES §8). ktfmt 0.65 masih 0.x; detekt stabil terakhir 1.23.8 (2025).
- OSV-Scanner v2.7.1 (9 November 2026) membaca `go.mod` dan `gradle.lockfile`; daftar resminya **tidak** memuat lockfile Swift.
- Dependabot mendukung Swift Package Manager (`Package.resolved`, GitHub Changelog 31 Maret 2026; 18 §5.3), Gradle, dan Go modules.
- PMD CPD mendukung Go, Kotlin, dan Swift.

## Opsi

| Peran | Go | Swift | Kotlin |
|---|---|---|---|
| Formatter | **gofmt** (stdlib toolchain) / gofumpt | **swift format** (toolchain) / SwiftFormat (Lockwood) | **ktlint 1.x** / ktfmt 0.x |
| Linter | **go vet + golangci-lint v2** / staticcheck saja | **swift format lint** / SwiftLint 0.x | **ktlint + Android Lint** / detekt 1.23 |
| Pemindai kerentanan | **govulncheck** (analisis jangkauan panggilan) / OSV-Scanner | **Dependabot alerts** / OSV-Scanner (tanpa dukungan Swift tercantum) | **OSV-Scanner** atas `gradle.lockfile` / Dependabot saja |
| Duplikasi | **PMD CPD** | **PMD CPD** | **PMD CPD** |

Kriteria: ≥ 1.0 atau bagian toolchain resmi (ES §8, ADR-023), berjalan di CI Linux (BT-3), satu alat per peran bila mungkin, dan tidak mengirim kode ke layanan luar.

## Keputusan

1. **Go:** `gofmt -l` harus kosong; `go vet` dan golangci-lint v2 dengan konfigurasi di repositori (termasuk larangan `float32`/`float64` di paket uang lewat aturan `forbidigo` atau uji arsitektur, ADR-015 butir 6); `govulncheck ./...`.
2. **Swift:** `swift format` sebagai formatter dan `swift format lint --strict` sebagai linter, keduanya dari toolchain sehingga tidak ada dependensi 0.x. Kerentanan: peringatan Dependabot atas `Package.resolved`. **Gerbang CI** untuk Swift **[Perlu validasi]**: bila Dependabot tidak dapat dijadikan pemeriksaan yang memblokir merge, `ios-dev` menetapkan di minggu pertama salah satu dari (a) OSV-Scanner bila dukungan Swift sudah tercantum, atau (b) skrip CI yang memeriksa daftar peringatan Dependabot terbuka lewat API GitHub, dan butir ini diperbarui.
3. **Kotlin:** ktlint 1.x (format dan gaya) dan Android Lint (dengan peringatan dijadikan galat untuk kategori keamanan dan kinerja); OSV-Scanner atas `gradle.lockfile`, yang karena itu diaktifkan (Gradle dependency locking).
4. **Duplikasi:** PMD CPD per bahasa, dengan ambang token yang ditetapkan `code-reviewer` di minggu pertama.
5. **Kontrak:** `buf lint` dan `buf breaking` (ADR-013).
6. **Versi:** semua alat dipatok di repositori dan dicek live saat dipasang (ES §8). SwiftLint, ktfmt, dan detekt 2.x pra-rilis tidak dipakai.
7. **Semua alat menjadi gerbang merge** di CI (09.4 §5). Alat yang berjalan di Xcode Cloud (uji UI iOS) tidak menggantikan gerbang Linux.

## Konsekuensi

- Gerbang ES §7 lengkap untuk Go dan Kotlin sejak hari pertama. Swift punya formatter dan linter sebagai gerbang; pemindai kerentanannya bergantung pada butir 2 (RISK diterima sampai minggu pertama `ios-dev`).
- `swift format` lebih sedikit aturannya daripada SwiftLint. Harga itu diterima demi tanpa dependensi 0.x; bila kurang, SwiftLint dapat ditambah lewat butir ADR-023.
- Tidak ada alat yang mengunggah kode ke layanan luar selain GitHub, tempat repositori memang disimpan.
