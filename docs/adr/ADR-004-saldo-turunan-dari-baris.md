# ADR-004 — Saldo diturunkan dari baris

| | |
|---|---|
| Status | **Diterima.** Diterima pemilik produk di checkpoint S2 lewat "lanjut" (dokumen 17 D-04, bersama ADR-001, 002, 005, dan 008; tidak dibacakan satu per satu). Bukan pernyataannya sendiri |
| Tanggal | 11 November 2026 |
| Pengusul | `finance-domain` (rekomendasi, 07 §1.5 ADR-C); ditulis `architect` |
| Rincian | [07-financial-domain.md](../07-financial-domain.md) §2.5, §3.8; INV-04, INV-05, INV-11, INV-14, INV-26 |

Aturan uang kanonis di dokumen 07. ADR ini mencatat keputusannya pada tingkat aturan; mekanisme basis datanya diputuskan di S4.

## Konteks

Kriteria sukses nomor 1 bergantung pada saldo tiap tempat uang. Bila saldo disimpan sebagai angka yang diubah langsung, ia dapat menyimpang dari riwayat entrinya tanpa terdeteksi, dan koreksi lewat entri pembalik (brief §7) kehilangan jaminannya.

## Opsi

| | A. Saldo sebagai angka tersimpan yang diubah tiap transaksi | B. Saldo adalah jumlah baris; cache boleh ada |
|---|---|---|
| Sumber kebenaran | Angka saldo | Baris entri |
| Penyimpangan | Mungkin, dan diam-diam | Terdeteksi: cache dapat dihitung ulang kapan saja |
| Entri terlambat dan koreksi | Perlu logika penyesuaian saldo tersendiri | Tidak perlu: jumlahnya memang berubah |
| Kecepatan baca | Tinggi | Sama, bila cache dipakai |

## Keputusan

**Opsi B.** Saldo tercatat suatu posisi adalah jumlah `delta` baris-barisnya. Semua angka turunan (saldo tersedia, *uang yang ada*, kekayaan bersih, realisasi, terikat, sisa anggaran) dihitung dari baris, dari otorisasi serta reservasi yang berlaku, dan, untuk realisasi dan terikat, dari tautan yang berlaku antara entri dan otorisasi: tautan pada entri ditambah tautan susulan yang tidak dicabut (07 §3.8, INV-14, INV-26). Tautan susulan mengubah realisasi tanpa mengubah baris atau saldo apa pun. Bila saldo disimpan sebagai cache, cache itu diperbarui dalam transaksi basis data yang sama dengan entrinya, dan tidak pernah menjadi sumber kebenaran.

## Konsekuensi

- INV-04 menjadi automated test: setelah setiap perintah, saldo yang ditampilkan sama dengan hasil hitung ulang dari baris.
- Tidak ada perintah "ubah saldo". Satu-satunya cara mengubah saldo adalah menambah entri.
- Realisasi dan terikat dihitung ulang setelah entri baru, setelah `authorization.amended`, dan setelah tautan susulan ditambah atau dicabut, termasuk atas otorisasi yang sudah tertutup (INV-14).
- Apakah cache diperlukan sama sekali untuk satu rumah tangga adalah pertanyaan pengukuran di fase pembangunan (`ENGINEERING-STANDARDS.md` §3: tanpa optimasi tanpa pengukuran).
- Tindak lanjut S4: bentuk cache bila ada, indeks, dan cara hitung ulang diverifikasi di produksi (dokumen 09).
