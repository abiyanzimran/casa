# 05 — Kebutuhan Fungsional: Tier 3 (garis besar)

| | |
|---|---|
| Pemilik | `product-analyst` |
| Tahap | S3, fase B (12 November 2026) |
| Status | Draf pertama; ditinjau `critic` A (fase D); garis besar `COM` diperbarui putaran perbaikan fase D (CA-04) |
| Bagian dari | [05-functional-requirements.md](../05-functional-requirements.md) |
| Kedalaman | Satu halaman per modul, tanpa nomor FR (brief §1.1: "A scenario that falls in Tier 3 scope is written at outline level and needs no requirement IDs") |

**Bentuk tiap garis besar.** Tujuan; apa yang ditunda dan mengapa; apa yang tidak boleh dihalangi fondasi Tier 1; syarat naik tingkat. Fitur Tier 3 yang dirujuk memakai nomor dokumen [04](../04-feature-map.md). Ditunda membatasi dokumentasi dan pembangunan, **bukan** fondasi (brief §1.1): isolasi rumah tangga dan satu orang di beberapa rumah tangga diuji sejak Tier 1 (FR-FAM-020 di [FAM.md](FAM.md)).

**Syarat naik tingkat yang berlaku untuk semua modul.** Bagian Tier 3 baru naik bila (1) menyebut masalah brief §1.1 yang dipecahkannya, (2) menyebut usaha berulang yang ditambahkannya dan mengapa sepadan (PRN-02), (3) tidak membutuhkan perubahan pagar 06 PG-01 sampai PG-14 kecuali lewat ADR, dan (4) disetujui pemilik produk di checkpoint. Syarat khusus per modul ditulis di bawah.

**`COM`.** Chat rumah tangga suami-istri adalah **Tier 1** (pernyataan pemilik produk sendiri tentang chat; naik ke Tier 1 diterima pemilik produk atas saran lead, 12 November 2026; `_progress.md`, "Pernyataan pemilik produk tentang chat"); kebutuhannya di [COM.md](COM.md). Yang tetap Tier 3 hanya garis besar ini (CA-04, putaran perbaikan fase D):

| Fitur 04 | Tingkat | Mengapa ditunda | Yang tidak boleh dihalangi fondasi Tier 1 | Syarat naik |
|---|---|---|---|---|
| F-COM-003 Pesan suara | Could | Tidak ada masalah §1.1 yang kuat (`MSL-5` lemah); isi suara data pribadi baru **[Dok. 10]** | Model pesan yang dapat membawa lampiran berjenis lain | Hasil pilot menunjukkan teks tidak cukup; penilaian dokumen 10 |
| F-COM-004 Reaksi | Could | `MSL-7` lemah; risiko menjadi hitungan per orang (06 AT-4) | Tidak ada | Rancangan tanpa hitungan, rekap, atau tren per orang |
| F-COM-007 Pencarian pesan | Could | Tidak ada masalah di pilot; objek sudah dapat dicari di layarnya | Isi chat tidak dikunci ke format yang mencegah indeks kelak | Penilaian dokumen 09 dan 10 atas indeks isi chat |
| F-COM-013 Grup dan percakapan di antara lebih dari dua anggota | Future | Tahap publik | Percakapan menyimpan `participants`, bukan "suami dan istri" (06 §16.2); objek yang tidak terlihat semua peserta tidak menjadi kartu | F-FAM-010 (rumah tangga lebih besar) dibangun |
| F-COM-002 utas bebas, F-COM-005 pin, F-COM-006 polling | Will not build | Alasannya di dokumen 04 | — | — |

---

## Daftar isi

1. [`INV` Persediaan rumah tangga](#1-inv-persediaan-rumah-tangga)
2. [`GOL` di luar tabungan sederhana](#2-gol-di-luar-tabungan-sederhana)
3. [`MEAL` Masak dan rencana menu](#3-meal-masak-dan-rencana-menu)
4. [`TML` Garis waktu dan album kenangan](#4-tml-garis-waktu-dan-album-kenangan)
5. [Lebih dari satu rumah tangga](#5-lebih-dari-satu-rumah-tangga)
6. [Rancangan layar khusus tablet dan perangkat lipat](#6-rancangan-layar-khusus-tablet-dan-perangkat-lipat)
7. [Integrasi belanja online](#7-integrasi-belanja-online)
8. [`AI` percakapan saran terbuka](#8-ai-percakapan-saran-terbuka)
9. [Sisa Tier 3 per modul](#9-sisa-tier-3-per-modul)

---

## 1. `INV` Persediaan rumah tangga

| | |
|---|---|
| Tujuan | Tahu barang rumah tangga yang hampir habis tanpa memeriksa lemari |
| Fitur 04 | F-INV-001 (Could), F-INV-002 bentuk manual (Will not build) |
| Masalah | `MSL-6` lemah. Tidak ada masalah brief §1.1 yang menuntut stok |

**Ditunda, dan mengapa.** Stok yang diperbarui manual menuntut isian setiap kali barang dipakai, padahal pencatat utama pelupa (`MSL-5`); itu contoh pelanggaran PRN-02 sendiri (dokumen 04 §4.2). Bentuk termurah, "hampir habis" satu ketukan, sudah tercakup dengan menambah barang ke daftar belanja (FR-SHP-001 di [tier-2.md](tier-2.md)).

**Yang tidak boleh dihalangi Tier 1.** Barang daftar belanja dan entri pengeluaran dapat kelak dikaitkan ke satu "barang" tanpa mengubah inti entri (07; PRN-16). Stok, bila ada, adalah proyeksi dari pembelian, bukan sumber angka uang (PRN-17).

**Syarat naik tingkat.** Pilot menunjukkan barang yang sama berulang kali habis tanpa disadari **dan** stok dapat lahir sebagai akibat samping pembelian dan interval beli, tanpa isian pemakaian.

---

## 2. `GOL` di luar tabungan sederhana

| | |
|---|---|
| Tujuan | Merencanakan impian dan dana darurat: kapan tercapai, mana yang didahulukan |
| Fitur 04 | F-GOL-001 tanggal target dan proyeksi (Could); F-GOL-002 tonggak dan prioritas (Future); F-GOL-003 kontribusi per orang lintas periode (Will not build) |
| Masalah | `MSL-1` lemah. Tabungan sederhana dan dana darurat sebagai satu tujuan tabungan adalah Tier 1 (07 §3.9, F-FIN-015) |

**Ditunda, dan mengapa.** Pertanyaan pemilik produk "dana darurat, impian/goals. udah?" (kata-katanya, `_progress.md`, 12 November 2026; Q-123) dijawab Tier 1: dana darurat adalah satu tujuan tabungan dengan target dan kemajuan. Yang di luar itu menambah isian dan urutan yang harus dipelihara, dan proyeksi memerlukan aturan uang yang belum ada di 07 (dokumen 04 §10.4: definisi laju tabungan). Bawaan Q-123 butir (b): tidak ada bagian `GOL` yang naik sebelum pilot menunjukkan perlunya **[Usulan]**.

**Yang tidak boleh dihalangi Tier 1.** Tujuan tabungan Tier 1 dapat kelak membawa tanggal target dan prioritas sebagai kolom tambahan. Mengakhiri tujuan yang masih beralokasi belum punya aturan di 07 (permintaan dokumen 04 kepada `finance-domain`); Tier 1 tidak boleh menghapus alokasi diam-diam. Kontribusi per orang lintas periode tidak dapat naik dalam bentuk apa pun (06 AT-4; 07 §6.3; D-43).

**Syarat naik tingkat.** Calon termurah F-GOL-001 (satu isian tanggal target), dan hanya setelah 07 menetapkan rumus proyeksinya. Pemilik produk dapat menaikkannya di checkpoint S3 (Q-123 butir (b)).

---

## 3. `MEAL` Masak dan rencana menu

| | |
|---|---|
| Tujuan | Menjawab "masak apa hari ini?" tanpa berdebat |
| Fitur 04 | F-MEAL-001 daftar ide masakan (Could); F-MEAL-002 rencana menu, resep, giliran masak (Future); F-MEAL-003 menu menjadi daftar belanja (Future) |
| Masalah | `MSL-6` ("what should we cook?"); `KB-20` |

**Ditunda, dan mengapa.** Masalahnya nyata, tetapi setiap bentuk yang menjawabnya menuntut pemeliharaan (resep, menu, stok) yang lebih besar dari masalahnya, dan menu yang menghasilkan daftar belanja bergantung pada `INV` yang juga ditunda (dokumen 04 §4.3). Giliran masak mendekati penugasan; bila ada, hanya lewat permintaan tindakan yang diterima (06 §13.1, §13.4).

**Yang tidak boleh dihalangi Tier 1.** Daftar belanja (Tier 2) dapat kelak menerima barang dari satu ide masakan tanpa mengubah aturannya. Tidak ada data preferensi makan atau kesehatan yang dikumpulkan sebelum modul ini ada.

**Syarat naik tingkat.** Pilot menunjukkan "masak apa" sering menjadi gesekan; bentuk pertama F-MEAL-001 saja (daftar ide bersama, tanpa jadwal, tanpa resep).

---

## 4. `TML` Garis waktu dan album kenangan

| | |
|---|---|
| Tujuan | Mengenang yang dicapai bersama: tabungan tercapai, perbaikan selesai, perjalanan, pembelian besar |
| Fitur 04 | F-TML-001 album kenangan (Future); F-TML-002 garis waktu tonggak sebagai proyeksi hanya-baca (Could) |
| Masalah | `MSL-7` lemah |

**Ditunda, dan mengapa.** Tidak memecahkan masalah inti pilot, dan album menuntut kurasi berulang. Narasi bulanan AI (FR-AI-004) sudah memberi "apa yang dicapai bersama" sekali sebulan tanpa usaha.

**Yang tidak boleh dihalangi Tier 1.** Batasnya ditetapkan dokumen 04 §4.5: riwayat audit hanya bertambah dan tidak pernah membaca dari album; album dikurasi, dapat disunting dan dihapus, dan paling banyak mengambil salinan dari Peristiwa. Menghapus butir album tidak pernah menyentuh catatan asal. Album tidak boleh menyimpan salinan catatan `CARE` atau buku pribadi yang lepas dari hak hapus pemiliknya (ADR-007 butir 7).

**Syarat naik tingkat.** F-TML-002 lebih dulu, karena tanpa isian baru. Album baru setelah dokumen 10 menetapkan penyimpanan dan penghapusan foto.

---

## 5. Lebih dari satu rumah tangga

| | |
|---|---|
| Tujuan | Tahap publik: satu orang di beberapa rumah tangga (orang tua dan rumah tangga sendiri, keluarga campuran, pengasuhan bersama), dengan pengalih rumah tangga |
| Fitur 04 | F-FAM-009 (Future); terkait F-FAM-010 wali dan tanggungan, F-FAM-011 preset lain, F-FAM-012 delegasi |
| Masalah | Tahap publik (brief §1.1 "Goal"; §4.1) |

**Ditunda, dan mengapa.** Pilot satu rumah tangga; ibu suami tidak punya login (Q-25). Membangun pengalih dan preset lain sekarang menambah layar tanpa pengguna.

**Yang tidak boleh dihalangi Tier 1** (diuji di FR-FAM-020):

| Fondasi | Sumber |
|---|---|
| `household_id` pada setiap objek; akses di luar keanggotaan ditolak tanpa membedakan "tidak ada" dari "tidak boleh" | 06 §14; `ENGINEERING-STANDARDS.md` §4 |
| Akun dan anggota dua entitas; buku pribadi melekat pada anggota, bukan akun | ADR-003; D-45 |
| Pemberi persetujuan adalah daftar; jenis permintaan adalah diskriminator | 06 §14 |
| Lima konsep terpisah; aturan tidak membaca hubungan | 06 §3.1, PG-01 |

**Jalur bocor yang harus ditutup sebelum naik** (dokumen 03 §4.3): notifikasi dua rumah tangga di satu perangkat (teks layar kunci tidak menyebut rumah tangga, D-34); buku pribadi yang terpisah per rumah tangga; konteks AI per rumah tangga (FR-AI-006); pencarian, ekspor, dan penghapusan akun per keanggotaan; pihak luar bernama sama tidak dicocokkan lintas rumah tangga (FR-FAM-018 AC-4).

**Syarat naik tingkat.** Tahap publik dimulai, dan L-04 serta L-20 dijawab advokat (07 §8.2). Preset baru masing-masing memerlukan ADR dan uji PG-01 sampai PG-14.

---

## 6. Rancangan layar khusus tablet dan perangkat lipat

| | |
|---|---|
| Tujuan | Memanfaatkan layar lebar: dua panel, layar terbelah pada perangkat lipat dan layar ganda |
| Fitur 04 | Tidak ada nomor fitur; kebutuhan platform (brief §1.1 Tier 3, §10.3) |
| Masalah | Tidak ada masalah pilot; persyaratan platform |

**Ditunda, dan mengapa.** Yang ditunda hanya **rancangan khusus** per bentuk layar. Brief §1.1 Tier 3 menyatakan tata letak tetap harus menyesuaikan diri dengan lebar berapa pun sejak hari pertama. Syarat pemilik produk yang tercatat di `_progress.md`: responsif di ponsel Android, iPhone, iPad, tablet Android, dan ponsel lipat termasuk layar ganda seperti Duo, dengan mode gelap dan terang. Cakupan uji layar ganda terbuka di Q-120.

**Yang tidak boleh dihalangi Tier 1.** Tidak ada kebutuhan Tier 1 atau Tier 2 di dokumen 05 yang mengandaikan satu lebar layar, satu orientasi, atau satu panel. Tata letak dan uji per bentuk layar milik dokumen 08 (S4).

**Syarat naik tingkat.** Dokumen 08 menetapkan pola tata letak dasar yang menyesuaikan, dan pilot menunjukkan perangkat yang dipakai keluarga.

---

## 7. Integrasi belanja online

| | |
|---|---|
| Tujuan | Belanja daring (Shopee, Alfagift, KlikIndomaret, dan lainnya) tercatat tanpa mengetik ulang |
| Fitur 04 | F-SHP-005 (Future, **[Perlu validasi]**); terkait F-FIN-025 impor mutasi, F-CAL-005 sinkron kalender |
| Masalah | `MSL-5` (usaha mencatat) |

**Ditunda, dan mengapa.** Riset kemampuan integrasi **belum dilakukan**; dijadwalkan sebagai riset teknologi S4 (dokumen 04 §4.6; dokumen 18). Tidak ada klaim di sini bahwa API resmi, deep link, impor struk, atau pelacakan pesanan tersedia.

**Yang tidak boleh dihalangi Tier 1.** Tidak ada fitur Tier 1 atau Tier 2 yang bergantung pada integrasi apa pun; pilot berjalan dengan pencatatan manual. Entri dari impor kelak tetap entri biasa menurut 07, dengan pencatat orang (06 PG-12: tidak ada pelaku otomatis) dan idempoten (07 INV-10).

**Syarat naik tingkat.** Dokumen 18 mengklasifikasikan tiap layanan menurut brief §5E (API mitra resmi, deep link, tautan produk, impor struk, pelacakan manual), dan dokumen 10 menilai data yang ditarik.

---

## 8. `AI` percakapan saran terbuka

| | |
|---|---|
| Tujuan | Saran dan semangat atas permintaan anggota (brief §5O "Advice and encouragement on request") |
| Fitur 04 | F-AI-006 (Future); terkait F-AI-007 entri berbantuan AI (Future; Q-31) |
| Masalah | `MSL-7` |

**Ditunda, dan mengapa.** Keputusan pemilik produk nomor 1 (brief §5O): percakapan saran terbuka adalah Tier 3. Teks bebas menaikkan risiko yang belum ditangani: tanda tekanan berat (dokumen 10), kebocoran data pasangan lewat parafrase, dan biaya.

**Yang tidak boleh dihalangi Tier 1.** Semua pagar F-AI-005 (FR-AI-005 sampai FR-AI-009 di [tier-2.md](tier-2.md)) berlaku tanpa pengecualian: AI bukan pelaku, hanya data yang boleh dilihat anggota yang dibantu, tanpa nama, batas biaya. Konteks AI per rumah tangga (bagian 5).

**Syarat naik tingkat.** Dokumen 10 menetapkan penanganan tanda tekanan berat dan penilaian UU PDP atas keputusan 2 sampai 4; pilot versi ringan berjalan tanpa insiden privasi; pemilik produk menyetujuinya.

---

## 9. Sisa Tier 3 per modul

Bagian modul Tier 1 dan Tier 2 yang tidak disebut di tingkat itu adalah Tier 3 (brief §1.1). Rinciannya cukup di dokumen 04; di sini hanya penunjuk.

| Modul | Fitur Tier 3 di dokumen 04 | Fondasi yang sudah dijaga |
|---|---|---|
| `FAM` | F-FAM-010 sampai F-FAM-013; F-FAM-014 sampai 017 Will not build | Lima konsep terpisah (FR-FAM-001 AC-2); tanpa riwayat kegiatan per anggota (06 AT-3) |
| `FIN` | F-FIN-021 sampai F-FIN-026 | Isinya milik [FIN-1.md](FIN-1.md) dan [FIN-2.md](FIN-2.md) |
| `REQ` | F-REQ-014 sampai F-REQ-017 | Jenis permintaan sebagai diskriminator (06 §14) |
| `SHP` | F-SHP-003, F-SHP-004 | Penugasan pembelanja hanya lewat permintaan tindakan |
| `PLN` | F-PLN-002 sampai F-PLN-004 | Rencana umum kelak memakai primitif permintaan yang sama |
| `MNT` | F-MNT-004, F-MNT-005 | Perawatan berkala di pilot dapat dibuat sebagai rutinitas (FR-CAL-001) |
| `CARE` | F-CARE-004, F-CARE-005; F-CARE-006 pelacakan siklus Will not build (dokumen 04 §4.4) | Catatan `CARE` dapat dihapus seluruhnya (FR-CARE-004) |
| `CAL` | F-CAL-003 sampai F-CAL-006 | — |
| `NTF` | F-NTF-006 sampai F-NTF-008 | Isinya milik [NTF.md](NTF.md) |
