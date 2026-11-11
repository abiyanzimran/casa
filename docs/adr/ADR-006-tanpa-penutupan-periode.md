# ADR-006 — Tidak ada penutupan periode

| | |
|---|---|
| Status | **Diterima.** Diterima pemilik produk di checkpoint S2 lewat "lanjut" (dokumen 17 D-05; tidak dibacakan satu per satu). Bukan pernyataannya sendiri |
| Tanggal | 11 November 2026 |
| Pengusul | `finance-domain` (rekomendasi, 07 §1.5 ADR-E); ditulis `architect` |
| Rincian | [07-financial-domain.md](../07-financial-domain.md) §3.10, §3.11; INV-19, INV-22 |

Aturan uang kanonis di dokumen 07. ADR ini mencatat keputusannya dan tidak mengubah isinya.

## Konteks

Pencatat utama rumah tangga pilot pelupa; pencatatan harus menoleransi entri terlambat (brief §1.1). Keputusan pemilik produk nomor 7: catatan lama tidak pernah disunting; belanja yang terlupa ditambahkan sebagai entri terlambat dengan tanggal aslinya, dan selisih kas mengecil karenanya. Pembukuan lazim menutup periode agar laporan lama tidak berubah.

## Opsi

| | A. Periode ditutup; entri terlambat masuk periode berjalan | B. Tidak ada penutupan; laporan lama dihitung ulang |
|---|---|---|
| Laporan bulan lalu | Tetap | Dapat berubah |
| Kebenaran per kategori dan periode | Belanja Oktober tampil di November | Belanja tampil pada bulan terjadinya |
| Selaras dengan keputusan pemilik nomor 7 | Tidak: tanggal asli tidak dipakai | Ya |
| Usaha pengguna | Ada langkah "tutup bulan" | Tidak ada |

## Keputusan

**Opsi B.** Tidak ada penutupan periode. Entri terlambat boleh sejauh apa pun ke belakang, tetapi tidak lebih awal dari saldo awal posisi yang disentuhnya. Laporan dan anggaran dihitung dari *tanggal kejadian*; laporan periode lampau yang berubah menampilkan tanda "berubah sejak terakhir dilihat" beserta entri yang tanggal pencatatannya lebih baru daripada akhir periode itu.

## Konsekuensi

- Angka bulan lalu bukan angka tetap. Narasi bulanan dan ringkasan yang pernah dibaca dapat berbeda dari hitungan sekarang; tandanya wajib ada.
- Tidak ada operasi "tutup bulan" untuk dibangun, diuji, atau dilupakan pengguna.
- Urutan audit tetap terjaga oleh *tanggal pencatatan*, yang diisi server.
- Angka turunan yang disimpan per periode, bila ada, harus dapat dihitung ulang (ADR-004).
- Bila ADR-007 jatuh ke jalan cadangannya (entri penutup), rangkuman itu menjadi satu-satunya bentuk penutupan, dan hanya terjadi atas permintaan penghapusan.
- Ditinjau ulang bila tahap publik menuntut laporan yang dibekukan (misalnya ekspor untuk keperluan pajak).
