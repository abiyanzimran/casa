# 05 — Kebutuhan Fungsional (indeks)

| | |
|---|---|
| Pemilik | `product-analyst` |
| Tahap | S3, fase B (12 November 2026); diselaraskan sesudah fase B (chat Tier 1, 06 revisi kesembilan); putaran perbaikan fase D (12 November 2026; 06 revisi kesepuluh, 07 putaran perbaikan S3) |
| Status | Ditinjau `critic` dan diperiksa `security-privacy` (findings-only) di fase D; perbaikan fase D belum ditinjau ulang. Pengenal butir terbuka diseragamkan ke `PA-<MOD>-N` (CA-06; padanan lama disebut di daftar tiap berkas) |
| Membaca lebih dulu | [04-feature-map.md](04-feature-map.md) (fitur), [06-activity-approval-architecture.md](06-activity-approval-architecture.md) (alur, pagar), [07-financial-domain.md](07-financial-domain.md) (aturan uang), [02-product-principles.md](02-product-principles.md) (prinsip) |

## 1. Tujuan

Dokumen 05 menyatakan apa yang harus dapat dilakukan dan dilihat setiap anggota, per fitur dokumen 04, dalam bentuk yang dapat diuji. Ia tidak menetapkan aturan uang (kanonis di 07), aturan alur dan pagar (kanonis di 06), keamanan, autentikasi, dan teks hukum (dokumen 10, S4), atau tata letak dan kata layar (dokumen 08, S4). Bila sebuah kebutuhan memerlukan aturan yang belum ada, ia ditandai **[Perlu aturan]** dan diajukan kepada pemilik dokumen yang bersangkutan, bukan diisi sendiri.

## 2. Cara dokumen ini dipecah, dan mengapa

Kebutuhan Tier 1 ditulis lengkap dengan Given/When/Then (brief §1.1), sehingga satu berkas akan terlalu panjang untuk ditinjau dan terlalu besar untuk ditulis bersamaan. Dokumen 05 karena itu terdiri atas indeks ini dan satu berkas per kelompok modul di [`05-functional-requirements/`](05-functional-requirements/). Pembagiannya mengikuti kedalaman tier: Tier 1 per modul, Tier 2 satu berkas ringkas, Tier 3 satu berkas garis besar. Indeks ini memuat aturan bersama, kebutuhan lintas modul, dan tabel keterlacakan; berkas modul tidak mengulanginya.

| Berkas | Isi | Tier | Rentang pengenal | Baris | FR | BR |
|---|---|---|---|---|---|---|
| [FAM.md](05-functional-requirements/FAM.md) | Rumah tangga, keanggotaan, wewenang, visibilitas, keluar, pihak luar, pemberitahuan operator | 1 | FR-FAM-001 dst.; BR-FAM-001 dst. | 645 | 20 | 20 |
| [FIN-1.md](05-functional-requirements/FIN-1.md) | F-FIN-001 sampai F-FIN-010 | 1 | FR-FIN-001 sampai 099; BR-FIN-001 sampai 049 | 1.279 | 42 | 49 |
| [FIN-2.md](05-functional-requirements/FIN-2.md) | F-FIN-011 sampai F-FIN-020 | 1 | FR-FIN-100 sampai 199; BR-FIN-050 sampai 099 | 1.100 | 47 | 50 |
| [REQ.md](05-functional-requirements/REQ.md) | `REQ` Tier 1, dan permintaan tindakan F-REQ-013 (Tier 2) | 1, 2 | FR-REQ-*; BR-REQ-* | 1.209 | 45 | 37 |
| [NTF.md](05-functional-requirements/NTF.md) | `NTF` Tier 1 | 1 | FR-NTF-*; BR-NTF-* | 446 | 15 | 20 |
| [COM.md](05-functional-requirements/COM.md) | Chat rumah tangga: pesan teks dan foto, kartu objek, menanggapi dari kartu, balasan yang menempel, tanda terkirim dan dibaca, kehadiran, notifikasi N-23; sisi chat F-CARE-003 (Tier 2, Must bersyarat) | 1 (dan satu FR Tier 2) | FR-COM-001 sampai 018; BR-COM-001 sampai 018 | 563 | 18 | 18 |
| [tier-2.md](05-functional-requirements/tier-2.md) | `SHP`, `PLN`, `MNT`, `CARE`, `CAL`, `AI` versi ringan | 2 | FR-SHP-*, FR-PLN-*, FR-MNT-*, FR-CARE-*, FR-CAL-*, FR-AI-* | 673 | 24 | — |
| [tier-3.md](05-functional-requirements/tier-3.md) | Garis besar per modul Tier 3, termasuk garis besar `COM` Tier 3 | 3 | Tanpa nomor FR | 192 | — | — |
| **Jumlah** | | | | **6.107** | **211** | **194** |

Angka dihitung ulang dengan `grep` pada 12 November 2026 di akhir putaran perbaikan fase D, sesudah tindak lanjut atas 06 §15.12: baris = semua baris berkas; FR = judul `#… FR-`; BR = baris tabel yang diawali `| BR-`. Fase D tidak menambah atau menghapus FR (Must 192 termasuk FR-CAL-002 bersyarat S3-1 (d), Should 19; rinciannya §6); BR bertambah tiga (BR-FAM-019, BR-FAM-020, BR-NTF-020). FR-CAL-002 tetap berjudul FR di `tier-2.md` tetapi kini hanya penunjuk ke N-32 Tier 1 (FR-FIN-100/AC-1, FR-NTF-009/AC-5), berlabel Must bersyarat S3-1 (d), Tier 1 **[Usulan]**; hitungan Must dan Should di §6. Lead menghitung ulang sebelum checkpoint S3.

Chat rumah tangga (`COM`) kini **Tier 1 versi ringan** atas keputusan pemilik produk 12 November 2026 (`_progress.md`, "Pernyataan pemilik produk tentang chat" dan jawaban sesudahnya; "Chat naik Tier 1" menggantikan jawaban "Tier 2"). Alur dan pagarnya kanonis di 06 §16 dan ADR-011.

## 3. Format blok FR

Setiap kebutuhan Tier 1 memakai blok ini. Tier 2 memakai bentuk ringkasnya (dijelaskan di kepala [tier-2.md](05-functional-requirements/tier-2.md)).

```
### FR-FAM-0NN — <judul>
| | |
|---|---|
| Fitur | F-FAM-00X |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-… / KS-… |
| Aktor dan izin | … (cite 06 V/AT/PG) |
| Aturan acuan | 06 …; 07 … |
| Business rule | BR-FAM-0NN, … |
| Kedudukan | owner decision with certainty level / [Usulan] … / [Asumsi] … |

<short spec>

**Kriteria penerimaan**
- **AC-1** — *Given* … *When* … *Then* …
```

## 4. Konvensi

| Hal | Aturan |
|---|---|
| Nomor kriteria penerimaan | Lokal per FR, mulai dari AC-1. Dikutip sebagai `FR-XXX-NNN/AC-n`, misalnya `FR-FAM-012/AC-2` |
| Business rule | Setiap `BR-<MOD>-NNN` adalah **penunjuk** ke aturan kanonis di 07 atau 06 (atau keputusan di 17), bukan definisi kedua. Bila bunyi ringkas BR berbeda dari sumbernya, sumbernya yang berlaku dan selisihnya adalah temuan |
| Label bukti | Mengikuti [README](README.md#label-bukti): **[Terverifikasi]**, **[Asumsi]**, **[Usulan]**, **[Perlu validasi]** |
| Kedudukan keputusan | Kata atau keputusan pemilik produk hanya dikutip dari `_progress.md` atau dokumen 17, dengan tingkat kepastiannya (diputuskan pemilik produk sendiri; diterima atas saran lead; diterima lewat "lanjut"; dan seterusnya) |
| **[Perlu aturan]** | Kebutuhan memerlukan aturan yang belum ada di 06, 07, atau 10. Setiap berkas punya daftar butirnya beserta bawaan yang diusulkan dan pemutusnya; bawaan itu bukan aturan sampai pemutusnya menetapkannya |
| **[Dok. 10]** | Rincian keamanan atau privasi yang ditetapkan dokumen 10 di S4; kebutuhan menyatakan perilakunya, dokumen 10 parameternya |
| Kata di dalam tanda kutip | Isi yang harus tersampaikan, bukan kata layar; kata-katanya milik dokumen 08 |
| MoSCoW | Tier 1 semuanya Must. Tier 2 Should kecuali yang ditandai **Must bersyarat**: wajib bila fitur yang dijaganya dibangun |
| Simetri | Kebutuhan ditulis untuk "anggota dewasa"; kriteria yang tidak ditulis dua kali berlaku sama dengan kedua anggota ditukar (`LM-01`) |

## 5. Kebutuhan lintas modul

Berlaku untuk setiap FR di semua berkas tanpa perlu disebut ulang. Isinya kanonis di sumber yang ditunjuk; di sini hanya ringkasan dan cara diuji.

| ID | Kebutuhan | Sumber | Cara diuji |
|---|---|---|---|
| LM-01 | **Dua arah.** Tidak ada aturan, layar, bawaan, atau batas yang membaca hubungan, jenis kelamin, penghasilan, pemegang uang, atau pembuat rumah tangga | 06 PG-01, INV-A13; PRN-09 | Uji simetri: rangkaian perintah dengan kedua anggota dewasa ditukar |
| LM-02 | **Pencatatan tidak pernah terhalang** oleh state permintaan, sambungan, AI, atau anggota lain | 06 INV-A11, S-7; PRN-06; FR-AI-008 | Entri sah diterima pada setiap state, tanpa sambungan, dan dengan AI mati |
| LM-03 | **Notifikasi aman privasi dan berbatas.** Teks layar kunci umum; rincian atas pilihan penerima dan tanpa jumlah, judul, alasan, diskusi, atau hal kepedulian; jam tenang dan batas pengingat; jenis notifikasi baru hanya bila ada barisnya di 06 §11.2 | 06 §11.1, §11.2, INV-A30 sampai INV-A33, INV-A42; D-34; PRN-08 | Tinjauan: setiap notifikasi di dokumen 05 punya baris N-nn atau ditandai [Perlu aturan] |
| LM-04 | **Offline** menurut 06 §12: server memutus transisi, perintah idempoten, compare-and-set, waktu server, tanda dihitung server | 06 S-1 sampai S-8; 07 INV-10 | Uji putus-sambung dan pengulangan perintah |
| LM-05 | **Tanpa metrik pengawasan.** Tidak ada umpan kegiatan per orang, hitungan, peringkat, tren per orang, atau lokasi di mana pun. Tanda terkirim dan dibaca, sedang online, sedang mengetik, dan terakhir aktif **hanya** ada di layar chat (FR-COM-011, FR-COM-013), sama bagi keduanya, tanpa waktu baca dan tanpa riwayat; tidak pernah pada layar objek, notifikasi, ringkasan, laporan, atau ekspor; dan tidak ada pengingat "dibaca tapi belum dijawab" atau hitungan lama menjawab | 06 AT-3 sampai AT-7, PG-07, PG-09, §16.5, §16.6, INV-A45; PRN-11 (rumusan lama perlu diselaraskan, lihat permintaan di bawah) | Tinjauan layar, laporan, ekspor, dan keluaran AI; uji INV-A45 di FR-COM-012/AC-5 dan FR-COM-014/AC-3 |
| LM-10 | **Pesan tidak pernah mengubah state.** Pesan chat atau diskusi apa pun ("setuju", "oke", foto struk) tidak mengubah state objek, tidak melahirkan otorisasi, dan tidak dicatat sebagai tanggapan; tindakan dari kartu chat adalah perintah objek yang sama dengan dari layar objeknya | 06 §2.1, §16.3, §16.4, INV-A44 | FR-COM-002/AC-2, FR-COM-007/AC-8, FR-COM-009/AC-4 |
| LM-11 | **Tanda, kehadiran, dan isi chat tidak dibaca apa pun.** Aturan, pengingat, tanda, ringkasan, laporan, dan `AI` tidak membaca `read_state`, `presence`, atau isi chat; ekspor tidak memuat `read_state` atau `presence`; isi chat dalam ekspor menurut dokumen 10 | 06 INV-A45, §13.6 "Yang tidak dibaca"; SP-58 (rumusan) | Uji pengacakan `read_state` dan `presence`: keluaran di luar layar chat identik byte demi byte; skema ekspor tanpa kolom `read_state` atau `presence` |
| LM-06 | **Tidak ada tindakan atas nama orang dewasa lain, dan diam tidak memutuskan apa pun** | 06 PG-02, PG-04, INV-A10; PRN-10 | Majukan jam; kirim perintah atas nama orang lain |
| LM-07 | **Visibilitas dari bawaan per objek**, tidak ditanyakan per entri; dana bersama tidak dapat disembunyikan; penyempitan oleh pemilik senyap | 06 §9.2, V-1 sampai V-9, INV-A36; PRN-12, PRN-13 | Uji akses per objek dan per peran |
| LM-08 | **Pelaku terlihat pada objeknya** (siapa mencatat, membayar, menyetujui, mengoreksi) | 06 AT-1, AT-2; keputusan 4 | Setiap catatan menampilkan pelakunya |
| LM-09 | **Isolasi rumah tangga** dan satu orang di beberapa rumah tangga tidak terhalang | 06 §14; ADR-003; `ENGINEERING-STANDARDS.md` §4; FR-FAM-020 | Uji lintas rumah tangga |

**LM-05 sesudah keputusan chat.** Tanda terkirim dan dibaca pada pesan chat wajib menurut pernyataan pemilik produk sendiri ("ada tanda di baca/delivery chat itu wajib loh"); tanda yang sama pada kartu di chat diterima atas saran lead ("Ikut tanda dibaca"); kehadiran diputuskan pemilik produk sendiri, berlawanan dengan saran lead (`_progress.md`). 06 AT-5 dan PG-07 sudah dipersempit `architect` ke chat (revisi kesembilan), dan PRN-11 di dokumen 02 sudah diselaraskan (fase C). Arti "terkirim" (hanya oleh penarikan saat aplikasi dibuka) dan penggerak kehadiran (hanya pemakaian di depan dengan interaksi) adalah **[Usulan]** `security-privacy` SP-55 dan SP-60, diterapkan 06 revisi kesepuluh.

## 6. Tabel keterlacakan

Setiap kebutuhan Must (semua Tier 1, Must bersyarat Tier 2, dan Must bersyarat keputusan checkpoint) dipetakan ke fiturnya (dokumen 04), sekurangnya satu skenario ([dokumen 12](12-scenarios.md)), dan fase roadmap dokumen 13 (diisi di S5). Daftar FR disusun dari `grep` judul `#… FR-` dan baris `Tier / MoSCoW` di semua berkas modul, 12 November 2026: 211 FR, di antaranya 183 Tier 1 Must, 8 Tier 2 Must bersyarat (FR-COM-016, FR-CARE-002, FR-CARE-004, FR-AI-005 sampai FR-AI-009), dan 1 Must bersyarat S3-1 (d), FR-CAL-002 (Tier 1 **[Usulan]**, penunjuk ke N-32, mengikuti 04 F-CAL-002), jadi **191 FR Must ditambah 1 bersyarat S3-1 (d) = 192**. Yang tidak masuk tabel: **19** FR Tier 2 Should (termasuk FR-REQ-044, FR-REQ-045, FR-NTF-004). Bila S3-1 (d) ditolak, FR-CAL-002 kembali Should: 191 Must dan 20 Should. **Tidak ada FR Must tanpa skenario.** Fase D tidak menambah FR; kriteria penerimaan baru dipetakan ke skenario yang sudah ada dan kolom FR skenario itu diperluas (SCN-02, SCN-11, SCN-22, SCN-25, SCN-28, SCN-32, SCN-36, SCN-39, SCN-43, SCN-44, SCN-46).

### 6.1 `FAM`

| FR | Fitur | Skenario | Fase |
|---|---|---|---|
| FR-FAM-001 | F-FAM-001 | SCN-42 | S5 |
| FR-FAM-002 | F-FAM-001 | SCN-43 | S5 |
| FR-FAM-003 | F-FAM-002 | SCN-42, SCN-44 | S5 |
| FR-FAM-004 | F-FAM-002 | SCN-42, SCN-44 | S5 |
| FR-FAM-005 | F-FAM-002 | SCN-42 | S5 |
| FR-FAM-006 | F-FAM-003 | SCN-42 | S5 |
| FR-FAM-007 | F-FAM-004 | SCN-20 | S5 |
| FR-FAM-008 | F-FAM-004, F-FAM-005 | SCN-23, SCN-32 | S5 |
| FR-FAM-009 | F-FAM-005 | SCN-10, SCN-32 | S5 |
| FR-FAM-010 | F-FAM-005 | SCN-10, SCN-32 | S5 |
| FR-FAM-011 | F-FAM-001 | SCN-42, SCN-43 | S5 |
| FR-FAM-012 | F-FAM-006 | SCN-44 | S5 |
| FR-FAM-013 | F-FAM-006 | SCN-44 | S5 |
| FR-FAM-014 | F-FAM-006 | SCN-44 | S5 |
| FR-FAM-015 | F-FAM-006 | SCN-44 | S5 |
| FR-FAM-016 | F-FAM-004, F-FAM-006 | SCN-20 | S5 |
| FR-FAM-017 | F-FAM-004 | SCN-20 | S5 |
| FR-FAM-018 | F-FAM-007 | SCN-04, SCN-46 | S5 |
| FR-FAM-019 | F-FAM-008 | SCN-42 | S5 |
| FR-FAM-020 | F-FAM-001 | SCN-42 | S5 |

### 6.2 `COM`

| FR | Fitur | Skenario | Fase |
|---|---|---|---|
| FR-COM-001 | F-COM-001 | SCN-36, SCN-42, SCN-44 | S5 |
| FR-COM-002 | F-COM-001 | SCN-37 | S5 |
| FR-COM-003 | F-COM-001 | SCN-12, SCN-37 | S5 |
| FR-COM-004 | F-COM-008 | SCN-01, SCN-24, SCN-36 | S5 |
| FR-COM-005 | F-COM-008 | SCN-02, SCN-37 | S5 |
| FR-COM-006 | F-COM-008 | SCN-36 | S5 |
| FR-COM-007 | F-COM-009 | SCN-36, SCN-37 | S5 |
| FR-COM-008 | F-COM-009 | SCN-18, SCN-19, SCN-38 | S5 |
| FR-COM-009 | F-COM-010 | SCN-36, SCN-37 | S5 |
| FR-COM-010 | F-COM-008, F-COM-010 | SCN-38 | S5 |
| FR-COM-011 | F-COM-011 | SCN-32, SCN-36 | S5 |
| FR-COM-012 | F-COM-011 | SCN-32, SCN-36 | S5 |
| FR-COM-013 | F-COM-012 | SCN-32, SCN-36 | S5 |
| FR-COM-014 | F-COM-012 | SCN-32 | S5 |
| FR-COM-015 | F-COM-001 | SCN-37 | S5 |
| FR-COM-016 (Tier 2, Must bersyarat) | F-CARE-003 | SCN-09, SCN-10, SCN-39 | S5 |
| FR-COM-017 | F-COM-001 | SCN-36 | S5 |
| FR-COM-018 | F-COM-001, F-COM-009 | SCN-18, SCN-19, SCN-38 | S5 |

### 6.3 `FIN` (FIN-1.md)

| FR | Fitur | Skenario | Fase |
|---|---|---|---|
| FR-FIN-001 | F-FIN-001 | SCN-04, SCN-42 | S5 |
| FR-FIN-002 | F-FIN-001 | SCN-42 | S5 |
| FR-FIN-003 | F-FIN-001 | SCN-23 | S5 |
| FR-FIN-004 | F-FIN-001 | SCN-42 | S5 |
| FR-FIN-010 | F-FIN-002 | SCN-21 | S5 |
| FR-FIN-011 | F-FIN-002 | SCN-04, SCN-21 | S5 |
| FR-FIN-012 | F-FIN-002 | SCN-12, SCN-23 | S5 |
| FR-FIN-020 | F-FIN-003 | SCN-21, SCN-34 | S5 |
| FR-FIN-021 | F-FIN-003 | SCN-26 | S5 |
| FR-FIN-022 | F-FIN-003 | SCN-34 | S5 |
| FR-FIN-023 | F-FIN-003 | SCN-27 | S5 |
| FR-FIN-030 | F-FIN-004 | SCN-21, SCN-27, SCN-43 | S5 |
| FR-FIN-031 | F-FIN-004 (memakai F-FIN-014, F-FIN-011) | SCN-02, SCN-11, SCN-21 | S5 |
| FR-FIN-032 | F-FIN-004 | SCN-21 | S5 |
| FR-FIN-040 | F-FIN-005 | SCN-01, SCN-02, SCN-04, SCN-07, SCN-08, SCN-18, SCN-19, SCN-22, SCN-27 | S5 |
| FR-FIN-041 | F-FIN-005 | SCN-27 | S5 |
| FR-FIN-042 | F-FIN-005 | SCN-02 | S5 |
| FR-FIN-043 | F-FIN-005 | SCN-27, SCN-28 | S5 |
| FR-FIN-044 | F-FIN-005 | SCN-15 | S5 |
| FR-FIN-045 | F-FIN-005 | SCN-01, SCN-22 | S5 |
| FR-FIN-050 | F-FIN-006 | SCN-02, SCN-13, SCN-21 | S5 |
| FR-FIN-051 | F-FIN-006 | SCN-02, SCN-13, SCN-21, SCN-35 | S5 |
| FR-FIN-052 | F-FIN-006 | SCN-04 | S5 |
| FR-FIN-060 | F-FIN-007 | SCN-21, SCN-25 | S5 |
| FR-FIN-061 | F-FIN-007 | SCN-12 | S5 |
| FR-FIN-062 | F-FIN-007 | SCN-12, SCN-35 | S5 |
| FR-FIN-063 | F-FIN-007 | SCN-12 | S5 |
| FR-FIN-064 | F-FIN-007 | SCN-46 | S5 |
| FR-FIN-070 | F-FIN-008 | SCN-36, SCN-43 | S5 |
| FR-FIN-071 | F-FIN-008 | SCN-36 | S5 |
| FR-FIN-072 | F-FIN-008 | SCN-25, SCN-26, SCN-36 | S5 |
| FR-FIN-073 | F-FIN-008 | SCN-21, SCN-25, SCN-26, SCN-43 | S5 |
| FR-FIN-074 | F-FIN-008 | SCN-25, SCN-32 | S5 |
| FR-FIN-075 | F-FIN-008 | SCN-25 | S5 |
| FR-FIN-076 | F-FIN-008 | SCN-25 | S5 |
| FR-FIN-077 | F-FIN-008 | SCN-48 | S5 |
| FR-FIN-080 | F-FIN-009 | SCN-42 | S5 |
| FR-FIN-081 | F-FIN-009 | SCN-42 | S5 |
| FR-FIN-082 | F-FIN-009 | SCN-42 | S5 |
| FR-FIN-083 | F-FIN-009 | SCN-48 | S5 |
| FR-FIN-090 | F-FIN-010 | SCN-01, SCN-02 | S5 |
| FR-FIN-091 | F-FIN-010 | SCN-01 | S5 |

### 6.4 `FIN` (FIN-2.md)

| FR | Fitur | Skenario | Fase |
|---|---|---|---|
| FR-FIN-100 | F-FIN-011 | SCN-11, SCN-24, SCN-43 | S5 |
| FR-FIN-101 | F-FIN-011 | SCN-24 | S5 |
| FR-FIN-102 | F-FIN-011 | SCN-24 | S5 |
| FR-FIN-103 | F-FIN-011 | SCN-11, SCN-21, SCN-24 | S5 |
| FR-FIN-104 | F-FIN-011 | SCN-24 | S5 |
| FR-FIN-105 | F-FIN-011 | SCN-24 | S5 |
| FR-FIN-106 | F-FIN-011 | SCN-24 | S5 |
| FR-FIN-107 | F-FIN-011 | SCN-24 | S5 |
| FR-FIN-108 | F-FIN-011 (memakai F-FIN-014) | SCN-11 | S5 |
| FR-FIN-110 | F-FIN-012 | SCN-45 | S5 |
| FR-FIN-111 | F-FIN-012, F-FIN-016 | SCN-45 | S5 |
| FR-FIN-112 | F-FIN-012 | SCN-45 | S5 |
| FR-FIN-113 | F-FIN-012 | SCN-45 | S5 |
| FR-FIN-115 | F-FIN-013 | SCN-46 | S5 |
| FR-FIN-116 | F-FIN-013 | SCN-46 | S5 |
| FR-FIN-117 | F-FIN-013 | SCN-46 | S5 |
| FR-FIN-118 | F-FIN-013 | SCN-46 | S5 |
| FR-FIN-119 | F-FIN-013 | SCN-46 | S5 |
| FR-FIN-120 | F-FIN-014 | SCN-02, SCN-06 | S5 |
| FR-FIN-121 | F-FIN-014 | SCN-02 | S5 |
| FR-FIN-122 | F-FIN-014 | SCN-06, SCN-21 | S5 |
| FR-FIN-125 | F-FIN-015 | SCN-05, SCN-06 | S5 |
| FR-FIN-126 | F-FIN-015 | SCN-06 | S5 |
| FR-FIN-127 | F-FIN-015 | SCN-06 | S5 |
| FR-FIN-128 | F-FIN-015 | SCN-06 | S5 |
| FR-FIN-130 | F-FIN-016 | SCN-45 | S5 |
| FR-FIN-131 | F-FIN-016 | SCN-45 | S5 |
| FR-FIN-132 | F-FIN-016 | SCN-45 | S5 |
| FR-FIN-133 | F-FIN-016 | SCN-45 | S5 |
| FR-FIN-135 | F-FIN-017 | SCN-28, SCN-44 | S5 |
| FR-FIN-136 | F-FIN-017, F-FIN-005 | SCN-28 | S5 |
| FR-FIN-137 | F-FIN-017 | SCN-23 | S5 |
| FR-FIN-138 | F-FIN-017 | SCN-28 | S5 |
| FR-FIN-140 | F-FIN-018 | SCN-29 | S5 |
| FR-FIN-141 | F-FIN-018 | SCN-29 | S5 |
| FR-FIN-142 | F-FIN-018 | SCN-29 | S5 |
| FR-FIN-143 | F-FIN-018, F-REQ-011 | SCN-35 | S5 |
| FR-FIN-144 | F-FIN-018 | SCN-35 | S5 |
| FR-FIN-145 | F-FIN-019 | SCN-29 | S5 |
| FR-FIN-146 | F-FIN-019 | SCN-32 | S5 |
| FR-FIN-150 | F-FIN-020 | SCN-47 | S5 |
| FR-FIN-151 | F-FIN-020 | SCN-47 | S5 |
| FR-FIN-152 | F-FIN-020 | SCN-33 | S5 |
| FR-FIN-153 | F-FIN-020 | SCN-47 | S5 |
| FR-FIN-154 | F-FIN-020 | SCN-47 | S5 |
| FR-FIN-155 | F-FIN-020 | SCN-32, SCN-33 | S5 |
| FR-FIN-156 | F-FIN-020 | SCN-47 | S5 |

### 6.5 `REQ`

| FR | Fitur | Skenario | Fase |
|---|---|---|---|
| FR-REQ-001 | F-REQ-001 | SCN-01, SCN-12, SCN-37, SCN-43, SCN-45 | S5 |
| FR-REQ-002 | F-REQ-001 | SCN-01, SCN-31, SCN-44 | S5 |
| FR-REQ-003 | F-REQ-001 | SCN-01 | S5 |
| FR-REQ-004 | F-REQ-001 | SCN-19 | S5 |
| FR-REQ-005 | F-REQ-001 | SCN-13 | S5 |
| FR-REQ-006 | F-REQ-002 | SCN-05 | S5 |
| FR-REQ-007 | F-REQ-003 | SCN-01, SCN-12, SCN-13 | S5 |
| FR-REQ-008 | F-REQ-003 | SCN-37 | S5 |
| FR-REQ-009 | F-REQ-003 | SCN-05, SCN-37 | S5 |
| FR-REQ-010 | F-REQ-003 | SCN-18, SCN-38 | S5 |
| FR-REQ-011 | F-REQ-003 | SCN-30 | S5 |
| FR-REQ-012 | F-REQ-004 | SCN-14 | S5 |
| FR-REQ-013 | F-REQ-004 | SCN-14 | S5 |
| FR-REQ-014 | F-REQ-005 | SCN-01, SCN-12, SCN-13, SCN-19 | S5 |
| FR-REQ-015 | F-REQ-005 | SCN-19 | S5 |
| FR-REQ-016 | F-REQ-005 | SCN-14 | S5 |
| FR-REQ-017 | F-REQ-005 | SCN-01, SCN-12, SCN-13 | S5 |
| FR-REQ-018 | F-REQ-005 | SCN-01 | S5 |
| FR-REQ-019 | F-REQ-005 | SCN-13 | S5 |
| FR-REQ-020 | F-REQ-005 | SCN-19 | S5 |
| FR-REQ-021 | F-REQ-006 | SCN-31 | S5 |
| FR-REQ-022 | F-REQ-006 | SCN-31 | S5 |
| FR-REQ-023 | F-REQ-006 | SCN-31 | S5 |
| FR-REQ-024 | F-REQ-007 | SCN-30 | S5 |
| FR-REQ-025 | F-REQ-007 | SCN-30 | S5 |
| FR-REQ-026 | F-REQ-008 | SCN-24, SCN-36 | S5 |
| FR-REQ-027 | F-REQ-008 | SCN-24, SCN-36 | S5 |
| FR-REQ-028 | F-REQ-008 | SCN-48 | S5 |
| FR-REQ-029 | F-REQ-008 | SCN-24, SCN-25 | S5 |
| FR-REQ-030 | F-REQ-008 | SCN-42, SCN-43 | S5 |
| FR-REQ-031 | F-REQ-008 | SCN-26, SCN-48 | S5 |
| FR-REQ-032 | F-REQ-009 | SCN-08, SCN-14, SCN-22, SCN-26 | S5 |
| FR-REQ-033 | F-REQ-009 | SCN-08, SCN-26 | S5 |
| FR-REQ-034 | F-REQ-009 | SCN-08, SCN-14, SCN-22, SCN-26 | S5 |
| FR-REQ-035 | F-REQ-009 | SCN-26 | S5 |
| FR-REQ-036 | F-REQ-009 | SCN-26 | S5 |
| FR-REQ-037 | F-REQ-010 | SCN-22 | S5 |
| FR-REQ-038 | F-REQ-010 | SCN-22 | S5 |
| FR-REQ-039 | F-REQ-010 | SCN-22 | S5 |
| FR-REQ-040 | F-REQ-010 | SCN-22 | S5 |
| FR-REQ-041 | F-REQ-011 | SCN-35, SCN-37 | S5 |
| FR-REQ-042 | F-REQ-011 | SCN-12 | S5 |
| FR-REQ-043 | F-REQ-012 | SCN-31 | S5 |

### 6.6 `NTF`

| FR | Fitur | Skenario | Fase |
|---|---|---|---|
| FR-NTF-001 | F-NTF-001 | SCN-21, SCN-30 | S5 |
| FR-NTF-002 | F-NTF-001 | SCN-30 | S5 |
| FR-NTF-003 | F-NTF-001 | SCN-27, SCN-30 | S5 |
| FR-NTF-005 | F-NTF-002 | SCN-30, SCN-39 | S5 |
| FR-NTF-006 | F-NTF-002 | SCN-30 | S5 |
| FR-NTF-007 | F-NTF-003 | SCN-30 | S5 |
| FR-NTF-008 | F-NTF-003 | SCN-30 | S5 |
| FR-NTF-009 | F-NTF-004 | SCN-11, SCN-22, SCN-30 | S5 |
| FR-NTF-010 | F-NTF-004 | SCN-30 | S5 |
| FR-NTF-011 | F-NTF-005 | SCN-30 | S5 |
| FR-NTF-012 | F-NTF-005 | SCN-30 | S5 |
| FR-NTF-013 | F-NTF-005 | SCN-30 | S5 |
| FR-NTF-014 | F-NTF-005 | SCN-30 | S5 |
| FR-NTF-015 | F-NTF-005 | SCN-30 | S5 |

### 6.7 Must bersyarat (`tier-2.md`)

| FR | Fitur | Syarat | Skenario | Fase |
|---|---|---|---|---|
| FR-CAL-002 (penunjuk; diuji lewat FR-FIN-100/AC-1 dan FR-NTF-009/AC-5) | F-CAL-002 (Tier 1 **[Usulan]** di 04) | Bila S3-1 (d) diterima (N-32 Tier 1, 06 P-36); bila ditolak, Tier 2 Should | SCN-11 | S5 |
| FR-CARE-002 | F-CARE-001, F-CARE-002 | Bila F-CARE-001 atau F-CARE-002 dibangun | SCN-09, SCN-10, SCN-39 | S5 |
| FR-CARE-004 | F-CARE-001, F-CARE-002 | Sama | SCN-10 | S5 |
| FR-AI-005 | F-AI-005 | Bila F-AI-001 sampai F-AI-004 atau F-AI-008 dibangun | SCN-39, SCN-40 | S5 |
| FR-AI-006 | F-AI-005 | Sama | SCN-39, SCN-40 | S5 |
| FR-AI-007 | F-AI-005 | Sama | SCN-41 | S5 |
| FR-AI-008 | F-AI-005 | Sama | SCN-41 | S5 |
| FR-AI-009 | F-AI-005 | Sama | SCN-39 | S5 |

FR-COM-016 (Must bersyarat, `COM.md`) ada di 6.2.

**Pemeriksaan ulang sebelum checkpoint.** Bila FR Must ditambah atau dihapus di berkas modul, `grep` judul `#… FR-` dan baris `Tier / MoSCoW` dijalankan lagi; setiap FR yang tidak muncul di tabel ini, atau yang tidak disebut di kolom FR skenario yang dirujuknya, adalah temuan.
