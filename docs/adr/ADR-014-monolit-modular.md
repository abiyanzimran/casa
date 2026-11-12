# ADR-014 — Monolit modular satu proses

| | |
|---|---|
| Status | **Diusulkan** (`architect`, 12 November 2026); akan diputuskan pemilik produk di checkpoint S4 |
| Tanggal | 12 November 2026 (S4 fase C) |
| Pengusul | `architect` |
| Mengubah | — |
| Rincian | [09.1 §6](../09-technical-architecture/01-pilihan-teknologi.md); [09.2 §1](../09-technical-architecture/02-data-dan-skema.md) |

## Konteks

Brief §9.4 meminta pilihan antara monolit modular dan microservices, tanpa menganggap microservices lebih baik.

Dokumen 06 S-4 menuntut agar perubahan state, Peristiwa, dan kait uang berlaku dalam **satu transaksi**. Satu perintah persetujuan dapat menyentuh `request`, `authorization`, `reservation`, `entry_flag`, `object_event`, `reader_feed`, `notification`, dan `job` sekaligus. SEC-06 menuntut satu jalur baca untuk layar, laporan, ekspor, dan `AI`. Tim terdiri dari satu sampai tiga orang, dan semuanya berjalan di satu VPS (brief §1).

## Opsi

| Kriteria | **A. Monolit modular, satu proses** | B. Monolit tanpa batas modul | C. Microservices |
|---|---|---|---|
| Satu transaksi lintas modul | **Ya** | Ya | Tidak (saga, kompensasi) |
| Satu jalur baca | **Satu paket** | Mudah dilanggar | Disalin atau lewat jaringan |
| Beban operasi | Satu biner | Satu biner | Banyak proses dan jaringan internal |
| Batas kepemilikan data | **Dijaga uji** | Tidak ada | Dijaga jaringan |
| Skala tahap publik | Gandakan proses di belakang penyeimbang beban | Sama | Per layanan |

## Keputusan

1. **Satu biner Go, satu proses** berisi API, hub SSE, dan pekerja job. Ia ditemani satu PostgreSQL.
2. **Modul adalah paket `internal/`** dengan tabel milik sendiri (daftar di 09.2 §1). Modul lain hanya memanggil fungsinya, di dalam transaksi pemanggil.
3. **Aturan dependensi:**
   - `ledger` hanya ditulis `ledger`.
   - Semua bacaan data orang lain lewat `read`.
   - `chat` tidak membaca isi objek lain selain lewat `read`.
   - Hanya `chat` yang membaca `read_state` dan kehadiran.
4. **Uji arsitektur di CI:** daftar impor yang diizinkan antarpaket diperiksa dari `go list`. Impor di luar daftar memblokir merge.
5. **Jalan ke tahap publik:** proses digandakan tanpa mengubah modul. Yang berubah hanya hal yang saat ini berada di memori proses: hub SSE dan batas laju pindah ke `LISTEN/NOTIFY` dan tabel PostgreSQL, sedangkan kehadiran memakai `LISTEN/NOTIFY` tanpa penyimpanan.

## Konsekuensi

- Integritas transaksi lintas modul gratis. Tidak ada outbox antar-layanan dan tidak ada konsistensi eventual di aturan uang.
- Satu kesalahan memori atau panik dapat menjatuhkan seluruh layanan. Penanganannya: `recover` per permintaan dan per pekerjaan, systemd memulai ulang, dan klien menyambung ulang lalu menarik `Sync`.
- Batas modul hanya sekuat uji arsitekturnya. Uji itu wajib ada sebelum modul kedua ditulis.
- Memecah satu modul menjadi layanan sendiri kelak tetap mungkin, karena tabel dan fungsinya sudah terpisah. Hal itu tidak direncanakan.
