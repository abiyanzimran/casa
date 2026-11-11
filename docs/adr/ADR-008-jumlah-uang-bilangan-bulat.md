# ADR-008 — Jumlah uang: bilangan bulat bertanda 64-bit dengan eksponen per mata uang

| | |
|---|---|
| Status | Diusulkan; menunggu persetujuan pemilik produk di checkpoint S2 |
| Tanggal | 11 November 2026 |
| Pengusul | `finance-domain` (rekomendasi, 07 §1.5 ADR-G); ditulis `architect` |
| Rincian | [07-financial-domain.md](../07-financial-domain.md) §2.1, INV-01; Q-43 |

Aturan uang kanonis di dokumen 07. ADR ini mencatat keputusannya dan tidak mengubah isinya.

## Konteks

`ENGINEERING-STANDARDS.md` §5: uang disimpan dan dihitung sebagai bilangan bulat dalam satuan terkecil mata uangnya; floating point tidak pernah dipakai. Pasar pertama adalah Indonesia dengan IDR, dirancang untuk internasionalisasi kemudian (brief §1). Uang tunai dan hampir semua transaksi rumah tangga tidak mengenal sen; saldo bank kadang memuatnya.

## Opsi

| | A. Bilangan bulat 64-bit, satuan terkecil, eksponen per mata uang | B. Desimal presisi tetap | C. Floating point |
|---|---|---|---|
| Sesuai `ENGINEERING-STANDARDS.md` §5 | Ya | Tidak: standar meminta bilangan bulat | Dilarang |
| Pembulatan diam-diam | Tidak mungkin | Mungkin pada pembagian | Pasti |
| Sama di server dan klien | Ya, di bahasa apa pun | Bergantung pustaka | Tidak |

## Keputusan

**Opsi A.** `amount` adalah bilangan bulat bertanda 64-bit dalam satuan terkecil mata uangnya, selalu disertai `currency`. Tiap mata uang punya eksponen satuan dalam **tabel milik produk**; untuk IDR eksponennya 0, yaitu rupiah utuh. Sen pada saldo bank dibulatkan ke bawah (07 §2.1). MVP memakai satu mata uang per rumah tangga; entri yang memuat dua mata uang ditolak. Tidak ada pembagian yang menghasilkan pecahan di dalam model: pada "bagi rata", sisa pembagian ditaruh pada baris pertama.

## Konsekuensi

- Rentangnya ±9.223.372.036.854.775.807 satuan terkecil (2⁶³ − 1), jauh di atas angka rumah tangga mana pun dalam rupiah utuh.
- **Eksponen tidak diambil mentah dari ISO 4217.** Daftar ISO 4217 mencantumkan dua angka desimal untuk IDR (**[Perlu validasi]**: hanya dari cuplikan hasil pencarian atas sumber sekunder, [localizely.com/currency-code/IDR](https://localizely.com/currency-code/IDR/) dan [handwiki.org/wiki/ISO_4217](https://handwiki.org/wiki/ISO_4217), 11 November 2026; halamannya sendiri dan daftar resmi ISO tidak dibuka, dan harus dibuka di S4 sebelum pustaka mata uang dipilih). Pustaka mata uang yang mengikuti ISO akan menganggap satuan terkecil IDR adalah sen. Tabel eksponen produk yang menentukan, dan setiap pustaka pemformat yang dipilih di S4 diuji terhadapnya.
- Mengubah eksponen suatu mata uang setelah ada data berarti migrasi seluruh jumlah bermata uang itu; eksponen ditetapkan sekali per mata uang.
- Tipe jumlah uang didefinisikan satu kali dan dibagi ke klien lewat kontrak API (`ENGINEERING-STANDARDS.md` §2). Klien yang bahasanya tidak punya bilangan bulat 64-bit asli harus memakai representasi yang tidak kehilangan presisi; itu menjadi kriteria pemilihan stack di S4.
- Multi-mata-uang dalam satu entri tidak tersedia; dapat ditambahkan kemudian karena tiap baris membawa mata uangnya.
- 07 INV-01 menjadi pemeriksaan tipe dan skema: perintah berisi pecahan ditolak di batas server.
