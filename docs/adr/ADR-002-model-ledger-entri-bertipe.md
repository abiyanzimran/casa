# ADR-002 — Model ledger: entri bertipe dengan baris seimbang

| | |
|---|---|
| Status | Diusulkan; menunggu persetujuan pemilik produk di checkpoint S2 |
| Tanggal | 11 November 2026 |
| Pengusul | `finance-domain` (rekomendasi, 07 §1.5 ADR-A); ditulis `architect` |
| Rincian | [07-financial-domain.md](../07-financial-domain.md) §1, §2.3 sampai §2.5, §8.2; INV-02, INV-24, INV-26, INV-28 |

Aturan uang kanonis di dokumen 07. ADR ini mencatat keputusannya dan tidak mengubah isinya.

## Konteks

Brief §5C meminta perbandingan double-entry penuh, ledger yang disederhanakan, dan daftar transaksi polos, dengan syarat akurat, dapat diaudit, dan tanpa kosakata akuntansi di layar. Kebutuhan rumah tangga pilot (07 §1.1): enam dimensi custody terjawab terpisah; perpindahan uang tidak terhitung dua kali; koreksi hanya dengan menambah; pencatat pelupa; invarian dapat diuji; tim kecil; utang, cicilan, aset, dan kekayaan bersih ada di Tier 1; dana pribadi tertutup tetapi kontribusinya terlihat.

## Opsi

| | A. Double-entry penuh | B. Entri bertipe dengan baris seimbang | C. Daftar transaksi polos |
|---|---|---|---|
| Tanpa hitung ganda | Terjamin | Terjamin oleh keseimbangan per pemilik (07 INV-02) | Tidak terjamin |
| Posisi bukan uang (piutang, kewajiban, aset) | Akun | Baris posisi pada entri yang sama | Tabel dan aturan terpisah per jenis |
| Biaya | Bagan akun dan pemetaan jurnal | Tabel entri, tabel baris, satu tabel tautan susulan yang hanya ditambah (`entry_link`), dan daftar jenis tertutup | Rendah di awal, tinggi kemudian |

Perbandingan lengkapnya di 07 §1.2.

## Keputusan

**Opsi B.** Setiap entri punya satu jenis dari daftar tertutup (07 §2.5) dan sejumlah baris: baris *posisi* (tempat uang, piutang, kewajiban, aset) dan baris *sebab* dari daftar tertutup (07 §2.4), yang seimbang per pemilik uang. Tidak ada bagan akun yang dapat diatur.

Tiga hal yang dipertegas revisi dokumen 07 (11 November 2026) dan tidak mengubah keputusan ini:

- Satu entri boleh memuat baris di dua buku. Selain jenis yang memang dua buku (penyisihan, kontribusi, talangan, penggantian), jenis `income` memuat baris di buku rumah tangga dan buku pribadi pencatat bila sebagian pemasukan diarahkan ke dana pribadi (07 §3.1, INV-24, CTH-03b). Keseimbangan tetap diperiksa per buku.
- Tautan dari entri ke otorisasi, jadwal, atau objek induknya dapat ditambahkan sesudah entri tersimpan sebagai catatan `entry_link` tersendiri yang hanya ditambah; entri dan barisnya tidak berubah (07 §2.5, §3.8, INV-26).
- Daftar jenis entri memuat satu jenis yang **hanya ditulis sistem**: `exit_closure`, penutupan klaim antar-pemilik yang menyangkut buku pribadi anggota yang keluar, sebesar seluruh sisanya, tanpa baris tempat uang atau aset dan tanpa menyatakan ada pembayaran (07 §2.5, §8.2, INV-28; **[Usulan]**). Pelakunya `system` (`actor = system` hanya pada jenis ini); tidak ada perintah anggota yang dapat membuatnya, dan ia tidak dapat dibalik (dokumen 06 §3.5, §7.3; ADR-007 butir 8). Jenis ini mengikuti invarian yang sama dengan jenis lain, termasuk keseimbangan per buku.

## Konsekuensi

- Menambah jenis entri atau sebab adalah perubahan model yang memerlukan pembaruan dokumen 07, invariannya, dan test-nya; bukan konfigurasi.
- Satu invarian struktural (INV-02) ditambah aturan per jenis entri menjadi automated test (`ENGINEERING-STANDARDS.md` §5).
- Tidak tersedia dan diterima: akun bebas buatan pengguna, laporan berformat akuntansi, multi-mata-uang dalam satu entri.
- Kata "entri", "baris", dan "sebab" tidak pernah tampil di layar; padanannya diusulkan di 07 §1.4 dan ditetapkan dokumen 08.
- Bentuk penyimpanannya (tabel, indeks, tingkat isolasi) diputuskan di S4.
