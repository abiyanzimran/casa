# ADR-025 — Alat kualitas kode per bahasa

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: alat per bahasa seperti butir 1 sampai 7 (D-103, diterima dengan "lanjut"; [_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk"). Gerbang kerentanan Swift dikukuhkan di minggu pertama `ios-dev` (butir 2) |
| Tanggal | 12 November 2026 (S4 fase D; status checkpoint S4; butir 8 TypeScript dan Svelte sesudah S6); 13 November 2026 (butir 9 rahasia di repositori, FS-1 tugas 1.1a) |
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

## Catatan keputusan web lebih dulu (12 November 2026, sesudah S6)

Mencatat [ADR-030](ADR-030-klien-web-pwa-lebih-dulu.md) butir 8 dan 18. Butir 1 sampai 7 tetap. **Butir 8 (baru), TypeScript dan Svelte (klien web):**

| Peran | Alat | Label |
|---|---|---|
| Pengelola paket | pnpm, `strictDepBuilds` bawaan `true`, `allowBuilds` kosong kecuali disetujui `code-reviewer`; `--frozen-lockfile` di CI | **[Terverifikasi]** perilaku bawaan ([pnpm.io/settings/build](https://pnpm.io/settings/build), diakses 12 November 2026); versi saat dipasang |
| Formatter | Prettier dengan `prettier-plugin-svelte` | Versi dan status ≥ 1.0 **[Perlu validasi]** saat dipasang |
| Linter | ESLint dengan `typescript-eslint` dan `eslint-plugin-svelte`; `svelte-check` untuk tipe; aturan yang melarang `innerHTML`, `{@html}`, `eval`, `new Function`, `setTimeout` berisi string (SEC-167) | **[Perlu validasi]** saat dipasang |
| Pemindai kerentanan | OSV-Scanner atas `pnpm-lock.yaml`, alat yang sama dengan Kotlin | **[Terverifikasi]** `pnpm-lock.yaml` tercantum ([daftar resmi](https://google.github.io/osv-scanner/supported-languages-and-lockfiles/), diakses 12 November 2026) |
| Duplikasi | PMD CPD | Dukungan TypeScript dan berkas `.svelte` **[Perlu validasi]**; bila `.svelte` tidak didukung, CPD dijalankan atas modul `.ts` saja dan `code-reviewer` memeriksa komponen |
| Uji | Vitest (unit, vektor emas); Playwright Chromium dan WebKit (alur, header, CSRF, XSS; 10.5 §9) | **[Perlu validasi]** saat dipasang |
| Bundel | Pemeriksaan CI bahwa paket di peta bundel hanya dari daftar ADR-030 butir 8, dan bundel produksi tanpa penanda pemeriksaan diri atau alamat staging (SEC-179) | Skrip CI |

Semua menjadi gerbang merge (butir 7). Semua berjalan di CI Linux.

## Catatan alat rahasia di repositori (13 November 2026, FS-1 langkah 1 tugas 1.1a)

[14 §2.2](../14-engineering-roadmap.md) baris "Rahasia di repositori" menyerahkan pilihan alat ke `backend-dev` di minggu pertama, dengan revisi ADR ini; [09.4 §5](../09-technical-architecture/04-operasi-uji-biaya.md) menyebut pemeriksaannya. `backend-dev` memilih di tugas 1.1a; catatan ini mencatatnya. Butir 1 sampai 8 tetap.

**Fakta yang diperiksa 13 November 2026:**

| Fakta | Label dan sumber |
|---|---|
| gitleaks, pilihan bawaan yang lazim: README menyatakan "Gitleaks is feature complete", "Future releases will be security patches only", dan pemeliharanya mengalihkan fokus ke Betterleaks. Rilis terakhir v8.30.1, 21 Maret 2026, tanpa rilis sesudahnya. Isu #2305 meminta perbaikan CVE | **[Terverifikasi]** README dan rilis: [github.com/gitleaks/gitleaks](https://github.com/gitleaks/gitleaks), [rilis](https://github.com/gitleaks/gitleaks/releases), diakses 13 November 2026. Isu #2305 dicek `backend-dev`, 13 November 2026 |
| Betterleaks v1.10.1 adalah rilis terbaru yang bukan pra-rilis ("Latest", 12 November); sebelumnya v1.10.0 (4 November) dan v1.9.0 (29 September). v2.0.0 masih rc (rc.7, 4 November), jadi tidak dipakai (ES §8; ADR-023 butir 4) | **[Terverifikasi]** [github.com/betterleaks/betterleaks/releases](https://github.com/betterleaks/betterleaks/releases), diakses 13 November 2026 |
| Modul `github.com/betterleaks/betterleaks@v1.10.1` ada di proxy.golang.org (cap waktu proxy 6 November 2026, lebih awal dari tanggal terbit di halaman rilis; tag dibuat lebih dulu) | **[Terverifikasi]** dicek lead di proxy.golang.org, 13 November 2026 |
| Lisensi MIT | **[Terverifikasi]** berkas `LICENSE` di repositori Betterleaks, diakses 13 November 2026 |
| Validasi live kredensial mati secara bawaan; dinyalakan oleh `-v`/`--validate`, `BETTERLEAKS_VALIDATE=true`, atau `-a`/`--analyze` | **[Terverifikasi]** README Betterleaks, diakses 13 November 2026 |
| `go run` atau `go install` dengan sufiks versi mengabaikan `go.mod` repositori kita dan menyusun daftar build secara deterministik (MVS) dari `go.mod` alat itu | **[Terverifikasi]** [go.dev/ref/mod](https://go.dev/ref/mod), bagian `go install` dan MVS, diakses 13 November 2026 |

**Butir 9 (baru), rahasia di repositori:**

1. **Alat:** Betterleaks v1.10.1, dipatok di `Makefile` (`BETTERLEAKS_VERSION`) dan dijalankan target `make secrets`: `go run github.com/betterleaks/betterleaks@v1.10.1 git --no-banner --redact --ignore-gitleaks-allow .`. Tanpa biner atau pengelola paket tambahan: toolchain Go sudah ada.
2. **Cakupan:** seluruh riwayat dari semua ref (CI memakai `fetch-depth: 0`), dengan aturan bawaan alat. Satu pemindaian mencakup seluruh repositori, jadi kolom klien web, Android, dan iOS di 14 §2.2 tidak memerlukan alat lain.
3. **Pengecualian:** komentar izin inline diabaikan (`--ignore-gitleaks-allow`). Setiap pengecualian hanya boleh ada di berkas ignore di repositori, dengan alasannya, dan ditinjau `code-reviewer`.
4. **Tidak ada yang keluar dari mesin:** temuan disamarkan di keluaran (`--redact`), dan `--validate`, `--analyze`, serta `BETTERLEAKS_VALIDATE` dilarang di `Makefile` dan CI. Dengan begitu kriteria ADR ini (tidak mengirim kode ke layanan luar) terpenuhi.
5. **Gerbang merge** seperti butir 7, berjalan di job CI dengan `permissions: contents: read`, `persist-credentials: false`, dan tanpa rahasia CI.
6. **Jalan mundur:** gitleaks v8.30.x selama masih menerima patch keamanan (bahwa berkas ignore dan konfigurasi kita terbaca tanpa diubah **[Asumsi]**, dicek bila jalan mundur ini dipakai); atau pemindai lain ≥ 1.0 yang memenuhi kriteria ADR ini. Penggantiannya hanya menyentuh target `secrets`.

**ADR-023 dan dependensi transitif alat CI (keputusan `architect`):** ADR-023 **tidak berlaku** untuk dependensi transitif 0.x dari alat yang hanya berjalan di CI atau mesin pengembang. Alasannya: syarat ADR-023 (lockfile kita, pembungkus milik kita, jalan mundur di kode kita) menjaga API yang dipanggil kode produk, sedangkan kode kita hanya memanggil baris perintah alat itu, bukan API dependensinya. Pembacaan ini berlaku untuk semua alat di butir 1 sampai 9, dengan tiga syarat; bila salah satunya gagal, ADR-023 berlaku penuh:

- (a) alat itu sendiri lolos ES §8 dan kriteria ADR ini: ≥ 1.0 atau bagian toolchain resmi, bukan pra-rilis, dipatok tepat;
- (b) tidak ada kode produk yang mengimpornya sebagai pustaka, dan tidak ada bagiannya yang masuk biner, bundel web, atau aplikasi;
- (c) alat itu berjalan di job CI tanpa rahasia CI dan dengan izin baca saja, terpisah dari job yang memegang kunci unggah atau penandatanganan (09.4 §5).

Risiko rantai pasok alat CI ditangani lewat versi yang dipatok, daftar build yang deterministik, syarat (c), dan pemutakhiran bulanan ES §8, bukan lewat ADR-023.

**Konsekuensi:** gerbang rahasia aktif sejak tugas 1.1a untuk seluruh repositori. Pemutakhiran bulanan memeriksa apakah Betterleaks v2 sudah stabil; naik ke v2 diperlakukan sebagai pembaruan yang dapat merusak (dibaca catatan rilisnya, lalu dibuktikan merah dengan kunci palsu di klon sementara seperti di 1.1a).
