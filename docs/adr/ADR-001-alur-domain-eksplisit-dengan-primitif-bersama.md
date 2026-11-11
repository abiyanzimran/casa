# ADR-001 — Alur domain eksplisit di atas primitif bersama

| | |
|---|---|
| Status | Diusulkan; menunggu persetujuan pemilik produk di checkpoint S2 |
| Tanggal | 11 November 2026 |
| Pengusul | `architect` |
| Rincian | [06-activity-approval-architecture.md](../06-activity-approval-architecture.md) §1, §2 |

## Konteks

Brief §6 meminta penilaian apakah modul-modul Family OS berbagi primitif agar alurnya tersambung (masalah rumah → perbaikan → persetujuan anggaran → pembayaran), dan perbandingan antara satu workflow engine generik dan alur domain eksplisit. Brief juga melarang memaksakan satu model generik bila merusak kejelasan atau kebenaran uang.

Kendala yang menentukan: tim 1–3 developer; satu rumah tangga pilot; Tier 1 hanya memuat empat alur (permintaan anggaran dan jalan-jalan, kesepakatan rumah tangga, konfirmasi custody, keberatan); akibat uang setiap alur sudah ditetapkan dokumen 07 sebagai sepuluh kait (07 §10.2 butir 2); `ENGINEERING-STANDARDS.md` §1 melarang abstraksi tanpa pemakai kedua dan §6 menuntut property test atas logika persetujuan; Tier 3 tidak boleh terhalang (brief §1.1).

## Opsi

| | A. Workflow engine generik | B. Alur domain eksplisit di atas primitif bersama |
|---|---|---|
| Bentuk | State, transisi, guard, dan aksi sebagai data konfigurasi | Satu tabel transisi tertulis per alur di kode server |
| Kebenaran uang | Kait uang menjadi aksi generik; baru terlihat saat dijalankan | Tiap transisi menyebut kait dokumen 07-nya |
| Pengujian | Mesin dan setiap konfigurasi harus diuji | Property test per tabel tertutup |
| Biaya bagi tim kecil | Mesin, bahasa konfigurasi, alat pengamatan | Beberapa ratus baris per alur |
| Alur baru (Tier 3) | Konfigurasi | Tabel baru; memakai primitif yang sama |

## Keputusan

**Opsi B.** Primitif bersama yang dipertahankan ada enam: Permintaan (memuat versi syarat dan tanggapan), Tugas, Diskusi, Lampiran, Peristiwa, Notifikasi. Objek-objek ditautkan dengan `subject_ref = (kind, id)`.

Dari kandidat brief §6: Activity bukan primitif (tidak ada tipe induk generik); Proposal dan Approval adalah bagian dari Permintaan; Assignment adalah atribut Tugas; Audit Record sama dengan Peristiwa; Timeline Entry adalah proyeksi Tier 3; Budget dan Financial Transaction adalah milik domain keuangan (dokumen 07).

Yang dapat diatur keluarga adalah **siapa yang menyetujui**, lewat satu fungsi penentu pemberi persetujuan (ADR-010), bukan bentuk alurnya.

## Konsekuensi

- Tabel transisi di dokumen 06 menjadi spesifikasi sekaligus sumber property test; transisi yang tidak ada di tabel ditolak.
- Keluarga tidak dapat merancang alur sendiri, dan alur baru memerlukan rilis. Diterima: tidak ada kebutuhan pilot untuk itu.
- Rantai lintas modul adalah tautan antar-objek yang masing-masing punya state machine sendiri, bukan satu alur panjang.
- Risiko pengulangan bentuk antar-alur dijaga dengan menaruh inti perundingan (ajukan, usul balik, setujui, tolak, batalkan) satu kali di primitif Permintaan.
- Glosarium: *aktivitas* tetap kata kumpulan, bukan tipe.
- Ditinjau ulang bila jumlah alur yang berbentuk sama melewati apa yang dapat dirawat sebagai tabel terpisah, atau bila tahap publik menuntut alur yang dapat diatur keluarga.
