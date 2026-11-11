# Log Kemajuan

Dicatat lead setelah setiap tahap (brief §0.2). Bila konteks terhapus atau sesi dimulai ulang: baca berkas ini dan [00-glossary.md](00-glossary.md), lalu lanjutkan dari **Tahap berikutnya**.

## Posisi saat ini

| | |
|---|---|
| Tahap terakhir yang selesai | S1 — Framing (11 November 2026) |
| Status | **Menunggu pemilik produk di checkpoint S1.** S2 tidak dimulai sebelum pemilik menjawab sepuluh pertanyaan atau menyatakan "lanjut pakai default" |
| Tahap berikutnya | S2 — Fondasi domain: 07, lalu 06, lalu 02 dan 03 |

## S1 — Framing

### Dokumen

| Dokumen | Pemilik | Status | Catatan |
|---|---|---|---|
| [README.md](README.md) | lead | ditinjau | Kerangka, tautan, konvensi; skema pengenal ditambah `Q`, `A`, `R`, `K`, `L` |
| [00-glossary.md](00-glossary.md) | lead; bagian 7 `finance-domain` | ditinjau (versi pertama) | Istilah keuangan adalah definisi kerja; titik bertanda *(diputuskan di S2)* difinalkan bersama dokumen 07 |
| [18-market-research.md](18-market-research.md) | `researcher` | ditinjau | Pasar dan regulasi terisi; integrasi, desain platform, dan teknologi masih placeholder (S3/S4) |
| [17-open-questions-assumptions.md](17-open-questions-assumptions.md) | `product-analyst` | ditinjau | 50 pertanyaan, asumsi, riset yang perlu divalidasi, pertentangan sumber, butir tinjauan hukum |

### Tinjauan

`critic` meninjau dua kali.

| Tinjauan | Menghalangi | Perlu diperbaiki | Minor | Putusan |
|---|---|---|---|---|
| Pertama | 1 | 16 | 19 | Belum siap |
| Kedua | 1 | 4 | 16 | Belum siap |

Semua temuan menghalangi dan perlu-diperbaiki dari kedua tinjauan sudah dikerjakan pemilik dokumennya. Perbaikan atas tinjauan kedua diperiksa lead terhadap teks temuan, **tidak** ditinjau `critic` untuk ketiga kalinya. Minor yang sengaja dibiarkan ada di "Dibawa ke S2" di bawah.

Angka yang dihitung ulang lead dengan skrip (di luar repositori): contoh talangan Rp100.000 lalu penggantian, untuk dua letak dana pribadi terhadap garis kendali rumah tangga. Hasil: kekayaan bersih turun tepat Rp100.000, pengeluaran tercatat satu kali, pada kedua pilihan.

### Keputusan

Belum ada ADR. S1 tidak memutuskan apa pun atas nama pemilik produk; semua default berstatus **[Usulan]** sampai checkpoint.

Keputusan proses yang diambil lead:

- Riset S1 dibatasi pada pasar dan regulasi, sesuai brief §0.3.
- Label **[Terverifikasi]** hanya untuk sumber primer yang benar-benar dibuka; sumber sekunder saja atau cuplikan pencarian berlabel **[Perlu validasi]** (dokumen 18 §0.2).
- Di glosarium, *tempat uang* (`holding`) dipakai agar *akun* hanya berarti identitas login; peran pemilik bernama `household owner` agar tidak tertukar dengan *pemilik uang*.

### Asumsi yang dipakai

Daftar lengkap di dokumen 17 bagian 3. Yang paling berpengaruh:

- Semua `[DEFAULT]` brief §1 berlaku.
- Pencatatan pengeluaran tidak pernah dihalangi status permintaan (A-20).
- Desain dan rencana pilot diperlakukan seolah pengecualian "kegiatan pribadi atau rumah tangga" UU PDP **tidak** berlaku, sampai advokat menjawab (A-13, dokumen 18 §2.1).
- Anak dimodelkan sebagai profil tanggungan tanpa login; login anak adalah keputusan tahap publik.

### Menunggu pemilik produk

Sepuluh pertanyaan checkpoint S1 (dokumen 17 §1.1), urut prioritas: Q-09, Q-01, Q-02, Q-04, Q-03, Q-08, Q-05, Q-06, Q-29, Q-10. Empat keputusan tambahan yang dibutuhkan sebelum S2 (§1.2): Q-07, Q-49, Q-25, Q-26. "Lanjut pakai default" mencakup keempat belas butir dan catatan pertanyaan 10.

Tindakan yang hanya bisa dilakukan pemilik (§1.3):

- Uji langsung dua atau tiga aplikasi pesaing, paling lambat S3.
- Mengunduh teks resmi PP 33/2026 bila `researcher` tetap gagal membukanya, paling lambat S4.

### Dibawa ke S2

Untuk `finance-domain` (dokumen 07 dan glosarium bagian 7):

- Letak dana pribadi terhadap garis kendali rumah tangga (Q-02), dan apakah angka utama "hanya dana bersama" diberi nama sendiri atau definisi *uang yang ada* disesuaikan (K-13).
- Definisi *realisasi* masih "jumlah entri pengeluaran", bertentangan dengan cicilan sebagai pelunasan yang dihitung terhadap anggaran; bunga cicilan menurunkan kekayaan bersih tanpa entri pengeluaran (K-11, Q-16).
- Pelepasan *hak ganti* bila dana pribadi di luar garis menaikkan kekayaan bersih tanpa pemasukan (temuan minor N-14 tinjauan kedua).
- Perlakuan kelebihan yang dibayar dari dana bersama (Q-13) dan selisih kas positif (Q-05).
- Q-43 sampai Q-48 belum punya default.

Untuk `architect` (dokumen 06 dan ADR):

- ADR atas tafsir "append-only" untuk penghapusan data anggota: inti moneter entri tetap dan dipseudonimkan, rincian pribadi dihapus (Q-22, K-10). UU PDP Pasal 43–44 tidak memberi pengecualian bagi keutuhan catatan bersama.
- Konfirmasi serah terima (Q-19), jenis permintaan untuk pemakaian dana bersama yang bukan pengeluaran (Q-03), alasan penolakan (Q-50).

Untuk lead:

- Rumusan *ambang permintaan* di glosarium masih "pengeluaran tidak rutin"; diperluas menjadi "pemakaian dana bersama" bila pemilik menerima default Q-03 (K-12).
- "Angka utama" dan "baris kedua" (Q-02) belum menjadi istilah glosarium; ditetapkan setelah dokumen 07.

Untuk `security-privacy` (pemeriksaan awal dokumen 06 sebelum checkpoint S2, lalu dokumen 10 di S4):

- Utang antar-pasangan yang menyentuh dana pribadi pasti terlihat oleh keduanya; dua kebocoran dana pribadi yang dinyatakan di Q-02.
- Keberadaan tempat uang pribadi yang tidak dibagikan: usulan tanpa penanda bagi pasangan.
- Pilihan (iii) pada Q-10 (DPIA ringkas dan perjanjian prosesor selama fase pembangunan) menyiratkan pekerjaan yang belum dijadwalkan.

Untuk `researcher` (diperluas di S3/S4):

- L-16: apakah pemrosesan di wilayah Indonesia oleh penyedia asing, atau hosting di luar negeri, terhitung transfer lintas batas.
- R-05: kelayakan penyedia AI yang memproses di Indonesia dan model yang di-host sendiri.
- Teks resmi PP 33/2026; pendaftaran PSE lingkup privat; label privasi Apple; angka Epstein dkk. dicocokkan dengan PDF.
- Kategori komunikasi keluarga hanya dua halaman vendor (celah terhadap brief §15).
