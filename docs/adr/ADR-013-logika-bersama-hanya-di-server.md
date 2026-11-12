# ADR-013 — Logika bersama hanya di server; kontrak API satu sumber; pengecualian ES §2 untuk bantuan isian

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **opsi A**, termasuk pengecualian ES §2 di butir 3, disetujui tegas sesuai ES baris 3 (S4-1 (ii) A). Kata-katanya: "hp boleh ikut, validasi di app dan server wajib jadi a" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban keempat; dikukuhkan dengan "lanjut"). Kontrak (butir 4 sampai 6) diterima lewat D-100. Pengecualian itu sudah dipasang lead di `ENGINEERING-STANDARDS.md` §2 |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4; catatan web lebih dulu sesudah S6) |
| Pengusul | `architect` |
| Mengubah | Menuntaskan 06 S-9 (tanpa perkiraan offline di pilot) |
| Rincian | [09.1 §3, §4](../09-technical-architecture/01-pilihan-teknologi.md); [09.2 §10](../09-technical-architecture/02-data-dan-skema.md); [09.4 §4](../09-technical-architecture/04-operasi-uji-biaya.md) |

## Konteks

ES §2 menuntut tiga hal:

- satu implementasi yang berwenang untuk setiap aturan uang dan persetujuan, dengan server sebagai otoritas akhir;
- logika yang dijalankan klien offline berasal dari kode bersama atau hasil generate, "never a hand-written copy", lalu diperiksa ulang server;
- kontrak API didefinisikan satu kali, dan tipe klien di-generate darinya.

UI sudah diputuskan native per platform (ADR-012 butir 1), sehingga ada tiga bahasa: server, Swift, dan Kotlin.

Dokumen 06 §12 sudah memindahkan semua penilaian ke server: transisi (S-1), tanda dan jatah (S-8), dan selisih rekonsiliasi (S-10). Yang tersisa di klien:

- **Perkiraan offline** (06 S-9). FR-FIN-010/AC-6 dan FR-FIN-091/AC-2 hanya mengaturnya "bila ditampilkan"; yang wajib adalah angka server dengan waktunya dan entri yang belum terkirim tampil terpisah.
- **Tiga bantuan isian** dari dokumen 05: (a) selisih baris sebab terhadap jumlah, dengan simpan tidak tersedia bila tidak sama (FR-FIN-042/AC-2); (b) "bagi rata" dengan sisa pada baris pertama (FR-FIN-042/AC-3); (c) bentuk masukan: wajib diisi, bilangan bulat positif, panjang teks.

**Fakta** (diakses 12 November 2026) **[Terverifikasi]**:

- protovalidate (aturan validasi sebagai anotasi `.proto`) punya pustaka resmi untuk Go, Java, Python, C++, dan TypeScript/JavaScript, **tanpa Swift dan Kotlin** ([protovalidate.com](https://protovalidate.com); 18 §5 fase D). Aturan bentuk dapat ditulis satu kali di kontrak, tetapi hanya server yang dapat menegakkannya dari sana.
- Target Apple Kotlin/Native hanya didukung di host macOS dan memerlukan Xcode ([kotlinlang.org](https://kotlinlang.org/docs/native-target-support.html)). Kompilasi silang Rust ke iOS memerlukan SDK Apple **[Perlu validasi]**. Pengembangan dilakukan di Linux (brief §1).

## Opsi

| Kriteria | **A. Server saja + tiga bantuan isian berpagar** | B. Server saja tanpa bantuan isian | C. Modul KMP | D. Inti Rust (UniFFI) |
|---|---|---|---|---|
| Satu implementasi aturan uang dan persetujuan | Ya | Ya | Bantuan isian dibagi Swift dan Kotlin, tetapi server Go tetap terpisah | Hanya bila server Rust (09-A = B) |
| Memenuhi bunyi ES §2 tanpa pengecualian | **Tidak** | **Ya** | Tidak | Ya bila server Rust |
| Pengalaman isian | Selisih dan bentuk diketahui saat mengisi | Kesalahan baru diketahui saat perintah tiba di server, bisa berjam-jam kemudian bila offline | Sama dengan A | Sama dengan A |
| Build iOS dari Linux | Tidak ada kode bersama untuk iOS | Sama | Perlu macOS untuk setiap perubahan modul | Perlu SDK Apple |
| Kematangan | — | — | Swift export alpha | UniFFI pra-1.0 |

| Kriteria kontrak | OpenAPI 3.1 | **Protobuf + Connect** | GraphQL |
|---|---|---|---|
| int64 lossless (ADR-008) | Bergantung decoder | Native; JSON-nya string | Skalar khusus |
| Kompatibilitas dengan aplikasi lama | Disiplin manual | Aturan nomor medan + `buf breaking` | Deprecasi |
| Skema bingkai SSE | Terpisah | Satu berkas `.proto` | Subscription |
| Aturan validasi sekali | JSON Schema | Anotasi protovalidate (server) | Bergantung |

## Keputusan

1. **Aturan hanya di server** (Go; ADR-015). Aplikasi tidak memegang salinan aturan uang, persetujuan, tanda, jatah, periode, kecocokan kategori, atau visibilitas.
2. **Tanpa perkiraan offline di pilot.** Perangkat menampilkan angka server terakhir dengan waktunya dan entri yang belum terkirim secara terpisah (06 S-9). Bila kelak perkiraan diperlukan, ADR baru menambahkannya sebagai deskriptor dari server dengan satu penafsir per platform, dan itu menjadi butir tambahan pengecualian di bawah.
3. **Pengecualian ES §2 untuk bantuan isian** (opsi A; disetujui pemilik produk, S4-1 (ii) A):
   - **Isi pengecualian, tertutup:** hanya butir (a), (b), dan (c) di Konteks. Butir lain memerlukan ADR baru dengan persetujuan yang sama.
   - **Server tetap otoritas:** (a) diperiksa ulang oleh INV-02 di modul `ledger`; (c) diperiksa ulang oleh protovalidate di server; (b) tidak punya pasangan di server, karena server menerima pembagian apa pun yang berimbang.
   - **Pagar:** vektor uji emas. Uji server Go membangkitkan berkas data (masukan dan hasil yang diharapkan) dari implementasi acuan: fungsi INV-02 untuk (a), aturan protovalidate untuk (c), dan satu fungsi acuan Go untuk (b). Uji unit Swift dan Kotlin menjalankan berkas yang sama di CI; beda hasil memblokir merge (09.4 §4).
   - **Opsi B tidak dipilih** (perangkat tanpa bantuan isian, server menolak dengan kode galat per medan). FR-FIN-042/AC-2 dan AC-3 tetap berlaku seperti bunyinya. Kode galat per medan tetap ada di kontrak, karena server selalu memeriksa ulang (a) dan (c).
4. **Kontrak API: Protobuf dengan protokol Connect**, dikelola `buf` (lint, `buf breaking` terhadap cabang utama, generate untuk Go, Swift, dan Kotlin). Jumlah uang bertipe `int64` beserta `currency`. Pesan bingkai SSE didefinisikan di kontrak yang sama. **Aturan bentuk** ditulis satu kali sebagai anotasi protovalidate dan ditegakkan server dengan `buf.build/go/protovalidate` (v1.x). **Kode galat** domain dan per medan adalah satu `enum` di kontrak, ter-generate untuk ketiga bahasa. Nilai per jenis perintah yang dibaca klien (misalnya `hold_seconds` jeda urungkan, 09.3 §1) juga didefinisikan di kontrak, bukan ditulis di aplikasi.
5. **Pustaka klien:** connect-go (v2.0.2), connect-swift (1.2.3), dan connect-kotlin (0.10.x; terlihat v0.10.0, rilis terbaru yang bukan pra-rilis di [github.com/connectrpc/connect-kotlin/releases](https://github.com/connectrpc/connect-kotlin/releases), diakses 12 November 2026 **[Terverifikasi]**; 0.x menurut ADR-023). Bila connect-kotlin bermasalah, aplikasi Android memakai pesan protobuf hasil generate resmi bersama klien HTTP biasa. Unary Connect adalah HTTP POST, sehingga kontraknya tidak berubah.
6. **Kode non-UI klien** (klien API, antrean, sinkron, bantuan isian) dipisah dari kode layar di kedua aplikasi. Di iOS ia ditulis sebagai paket Swift yang diusahakan dapat dibangun dan diuji di Linux **[Perlu validasi]** per dependensi.

## Konsekuensi

- Tidak ada bahasa keempat dan tidak ada langkah build khusus macOS di luar aplikasi iOS itu sendiri.
- Perangkat tidak menampilkan tanda, sisa jatah, sisa anggaran, atau perkiraan saat offline. Yang tampil adalah angka server terakhir dengan waktunya dan daftar entri yang belum terkirim, sesuai 06 S-8 dan S-9.
- Tiga potong kode kecil ada dua kali (Swift dan Kotlin) di samping acuan Go, dan pagarnya adalah uji, bukan pembangkitan kode. Ini dicatat terang sebagai pengecualian ES §2, bukan disebut "terpenuhi". RISK-15 (pengecualian ditolak) gugur.
- `buf breaking` membuat perubahan kontrak yang merusak aplikasi lama tidak dapat masuk tanpa keputusan sadar.

## Catatan keputusan web lebih dulu (12 November 2026, sesudah S6)

Mencatat [ADR-030](ADR-030-klien-web-pwa-lebih-dulu.md). Butir 1 dan 2 tidak berubah: klien web, seperti aplikasi native, tidak memegang aturan.

- **Butir 3, bantuan isian:** klien web memerlukan ketiga bantuan isian (a), (b), (c) yang sama, ditulis sekali lagi di **TypeScript**. Vektor uji emas yang sama dari acuan Go dijalankan di CI web; beda hasil memblokir merge (ADR-030 butir 15). Isi pengecualian tetap tertutup pada tiga butir itu; yang bertambah hanya bahasanya. Karena bunyi ES §2 menyebut "once in Swift and once in Kotlin", tambahan TypeScript **menunggu persetujuan pemilik produk** atas bunyi baru ES §2 (ADR-030 butir 15); sampai disetujui, layar web memakai opsi B butir 3 (tanpa bantuan isian, server menolak dengan kode galat per medan).
- **Butir 4, kontrak:** `buf generate` juga untuk TypeScript (`@bufbuild/protobuf` 2.x, protoc-gen-es). Kontrak memuat `client_kind` (`ios`, `android`, `web`), `client_version`, dan kode galat `CLIENT_UPGRADE_REQUIRED` (ADR-028 butir 5); server menolak perintah tulis dari klien web di bawah versi minimum `web` sejak pilot (SEC-169 (d)).
- **Butir 5, pustaka klien:** `@connectrpc/connect-web` 2.x (18 §5.13.1). Dari browser hanya unary dan server streaming; kontrak tidak memakai client streaming atau bidi.
- **Butir 6:** kode non-UI web (klien API, antrean, sinkron, bantuan isian) dipisah dari komponen layar, seperti di aplikasi native.
- **Konsekuensi pertama di atas berubah:** TypeScript menjadi bahasa keempat di repositori, untuk klien web saja. Tidak ada langkah build khusus macOS sampai jalur iOS native dimulai.
