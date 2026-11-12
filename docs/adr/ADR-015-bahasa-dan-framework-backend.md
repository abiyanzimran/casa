# ADR-015 — Bahasa dan framework backend: Go

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **opsi A, Go** (S4-1 (i) A). Kata-katanya: "#S4-1: A Go" ([_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban pertama; dikukuhkan dengan "lanjut"). Sebelumnya: "Tunggu dokumen 09 (Saran)", diterima atas saran lead (dokumen 17 Q-34) |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4) |
| Pengusul | `architect` |
| Mengubah | — |
| Rincian | [09.1 §5, §13](../09-technical-architecture/01-pilihan-teknologi.md) |

## Konteks

Pemimpin tim kuat di Go dan Rust (brief §1). Brief §9.2 meminta penilaian atas kinerja API, waktu nyata, integritas transaksi, pekerjaan latar, keamanan, konkurensi, produktivitas, beban operasi, skalabilitas, kemudahan pemeliharaan, dan biaya, termasuk BaaS bila mengurangi usaha.

Riset 18 §5.3, §5.4, dan §5 fase D (diakses 12 November 2026) menemukan:

- **Go:** Go 1.27.4, dengan routing `net/http` stdlib.
- **Pustaka Go di jalur setiap permintaan:** pgx v5.13.0, sqlc v1.31.1, goose v3.30.0, dan connect-go v2.0.2, semuanya ≥ 1.0.
- **Pustaka Go di tepi sistem:** `golang.org/x/crypto` v0.59.0 (Argon2id, ACME) dan `golang.org/x/oauth2` v0.38.0 (token FCM), **keduanya 0.x**, dipelihara proyek Go sendiri **[Terverifikasi]**. pkg.go.dev menampilkan peringatan versi modul terbaru yang bertentangan dengan indeks proxy, sehingga versi dicek saat dipasang.
- **Rust:** Rust 1.99.1 dengan axum 0.8.9 dan sqlx 0.9.0 (keduanya 0.x, di jalur setiap permintaan), dan apalis 1.1.1 (1.0 baru 12 Oktober 2026).
- **Perbandingan memori dan waktu kompilasi:** tidak ada sumber primer.
- **BaaS:** Supabase self-hosted perlu 4 GB dan 11 komponen. Firebase gagal transaksi offline dan berharga per operasi. PocketBase diarsipkan.

## Opsi

| Kriteria | **A. Go** | B. Rust | C. BaaS |
|---|---|---|---|
| Pustaka 0.x yang diperlukan (ES §8; ADR-023) | `golang.org/x/crypto`, `golang.org/x/oauth2`: dari proyek Go, di tepi sistem; mengganti satu menyentuh satu paket | axum, sqlx: di jalur setiap permintaan HTTP dan setiap kueri | — |
| Transaksi Serializable dengan coba ulang | pgx + SQL eksplisit (sqlc) | sqlx | Bergantung; Firebase gagal |
| SSE dan TLS | Stdlib + `golang.org/x/crypto/acme/autocert` | axum + crate | Bawaan |
| Produktivitas tim kecil | Tinggi; kompilasi cepat **[Asumsi]** | Tinggi bagi pemimpin; lebih curam bagi anggota baru **[Asumsi]**; kompilasi lebih lambat **[Asumsi]** | — |
| Jaminan tipe untuk aturan uang | Cukup (tipe uang khusus, tanpa `float`) | **Lebih kuat** | Lemah |
| Inti bersama di perangkat kelak | Tidak | Mungkin (UniFFI) | Tidak |
| Biaya | VPS | VPS | Per pemakaian atau VPS besar |

## Keputusan

1. **Go**, versi stabil terbaru saat bangun (sekarang 1.27.4), dicek live saat dipasang (ES §8).
2. **HTTP:** stdlib `net/http` (routing bawaan), dengan connect-go untuk RPC dan handler SSE biasa. Tanpa framework web tambahan. Validasi kontrak dengan `buf.build/go/protovalidate` v1.x (ADR-013).
3. **Basis data:** pgx v5, dengan kueri SQL eksplisit yang di-generate sqlc. Tanpa ORM. Migrasi dengan goose.
4. **Pekerjaan latar:** tabel `job` sendiri dengan `FOR UPDATE SKIP LOCKED` dan `LISTEN/NOTIFY` (09.3 §9), diambil peran `job_runner` (09.2 §3). River (v0.51.2, 0.x) adalah alternatifnya bila kebutuhan antrean tumbuh, dan memerlukan butir ADR-023.
5. **Kripto dan protokol:** `golang.org/x/crypto/argon2`, `golang.org/x/crypto/acme/autocert`, dan `golang.org/x/oauth2` (FCM), ketiganya 0.x dan tercatat di ADR-023 dengan jalan mundurnya. JWT APNs ES256 dibentuk dengan `crypto/ecdsa` dari stdlib. Pustaka email mengikuti ADR-026, yang masih diusulkan sampai S4-2 dijawab menjelang deploy (AWS SDK for Go v2 `sesv2` bila SES dipilih).
6. **Uang:** tipe `Money{Amount int64; Currency string}` di satu paket. Tidak ada `float32` atau `float64` di jalur uang; linter atau uji arsitektur menolaknya (ADR-025).
7. **Rust (opsi B) tidak dipilih.** axum dan sqlx karena itu tidak masuk daftar ADR-023.

## Konsekuensi

- Jalur server memerlukan dua modul 0.x (`golang.org/x/crypto`, `golang.org/x/oauth2`), tercatat di ADR-023. Tidak ada pustaka 0.x di jalur setiap permintaan.
- Sebagian invarian yang di Rust dapat dijaga tipe dijaga di Go oleh uji property-based (ES §6). Rangkaian uji invarian karena itu wajib sebelum modul `ledger` dianggap selesai.
- Kode SQL eksplisit (sqlc) membuat setiap kueri layar terlihat dan dapat diperiksa indeksnya (ES §3: tanpa N+1, setiap kueri berindeks).
- Tim build: `backend-dev` bekerja dalam Go. Agen build ditulis lead sesudah stack disetujui.
