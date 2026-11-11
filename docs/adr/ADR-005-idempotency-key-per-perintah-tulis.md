# ADR-005 — Idempotency key per perintah tulis

| | |
|---|---|
| Status | **Diterima.** Diterima pemilik produk di checkpoint S2 lewat "lanjut" (dokumen 17 D-04, bersama ADR-001, 002, 004, dan 008; tidak dibacakan satu per satu). Bukan pernyataannya sendiri |
| Tanggal | 11 November 2026 |
| Pengusul | `finance-domain` (rekomendasi, 07 §1.5 ADR-D); diperluas ke perintah alur dan ditulis `architect` |
| Rincian | [07-financial-domain.md](../07-financial-domain.md) §2.5, INV-10, CTH-32; [06-activity-approval-architecture.md](../06-activity-approval-architecture.md) §12, INV-A12 |

## Konteks

`ENGINEERING-STANDARDS.md` §5 mewajibkan setiap penulisan uang membawa idempotency key. Aplikasi dipakai di jaringan buruk (brief §1: "reliable on poor networks"), sehingga kirim ulang dan ketukan ganda adalah keadaan biasa. Persetujuan, usul balik, dan pengingat yang tercatat dua kali sama merusaknya dengan belanja yang tercatat dua kali: persetujuan ganda dapat melahirkan dua otorisasi.

## Opsi

| | A. Deteksi duplikat dari isi (jumlah, waktu, pelaku) | B. Kunci per perintah, dibuat klien | C. Kunci hanya untuk perintah uang |
|---|---|---|---|
| Dua belanja yang memang sama | Salah ditolak | Diterima: kuncinya berbeda | Diterima |
| Kirim ulang | Tertangkap hanya bila isinya persis sama | Tertangkap | Tertangkap untuk uang saja |
| Perintah alur | Tidak tercakup | Tercakup | Tidak tercakup |

## Keputusan

**Opsi B, untuk setiap perintah tulis, uang maupun alur.** Klien membuat satu kunci per niat pengguna. Kunci unik per rumah tangga dan disimpan bersama hasil perintahnya. Perintah dengan (`household_id`, `idempotency_key`) yang sudah diterima tidak menulis apa pun dan mengembalikan hasil perintah pertama.

## Konsekuensi

- 07 INV-10 dan 06 INV-A12 menjadi automated test: setiap perintah dikirim ulang; jumlah entri, baris, tautan susulan, otorisasi, reservasi, peristiwa, dan notifikasi tidak bertambah. Menambah atau mencabut tautan susulan adalah perintah tulis dengan kuncinya sendiri (07 §2.5 `entry_link`).
- "Langsung dipakai" pada layar uang masuk adalah dua perintah dengan dua kunci (pemasukan, lalu pemakaiannya; 07 §3.1): masing-masing dikirim ulang sendiri-sendiri, dan gagalnya yang kedua tidak membatalkan yang pertama.
- Kunci dibuat saat niat lahir, bukan saat dikirim, agar perintah yang ditulis offline tetap membawa kunci yang sama pada setiap percobaan.
- Kunci yang sama dengan isi berbeda diperlakukan sebagai kesalahan klien dan ditolak; tidak pernah diam-diam menimpa.
- Idempotency key tidak menggantikan versi syarat dan pemeriksaan state (dokumen 06 §12, S-3 dan S-4): yang satu mencegah pengulangan, yang lain mencegah tanggapan atas keadaan yang sudah berubah.
- Tindak lanjut S4: berapa lama kunci disimpan, bentuk kuncinya, dan cara keunikan dipaksakan (dokumen 09).
