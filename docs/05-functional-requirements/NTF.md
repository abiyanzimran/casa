# 05 — Kebutuhan fungsional: `NTF` (notifikasi)

Bagian dari dokumen 05. Indeks, tabel keterlacakan, dan aturan umum ada di [../05-functional-requirements.md](../05-functional-requirements.md). Pemilik: `product-analyst`. Tahap S3 fase B, 12 November 2026; diselaraskan dengan revisi kesembilan dokumen 06 (§11.1 jadwal ringkasan, penggabungan `kabar`, waktu pengingat bertanggal; §11.2 N-23 sampai N-31; §16 chat) pada putaran penyelarasan S3, 12 November 2026; putaran perbaikan fase D (12 November 2026) atas tinjauan `critic` A (CA-01, CA-07, CA-08), `security-privacy` (SP-46, SP-50, SP-52, SP-55), dan 06 revisi kesepuluh (N-32, butir ringkasan yang dibuang, teks layar kunci per kelas). Perbaikan fase D belum ditinjau ulang.

| Hal | Isi |
|---|---|
| Lingkup | F-NTF-001 sampai F-NTF-005 (Tier 1, kebutuhan lengkap). F-NTF-003 dan F-NTF-004 Tier 1 sebagai akibat keputusan yang sudah ada, berbeda dari teks brief §1.1 (04 §4.8; Q-135). Satu FR Tier 2 untuk peristiwa kepedulian, pendamping `AI`, tugas, dan rutinitas (N-28 sampai N-31). Pengingat tagihan N-32 Tier 1 (06 revisi kesepuluh); kebutuhannya di [FIN-2.md](FIN-2.md) FR-FIN-100/AC-1 |
| Rentang ID | `FR-NTF-001` sampai `FR-NTF-015`; `BR-NTF-001` sampai `BR-NTF-020` |
| Sumber aturan | [06](../06-activity-approval-architecture.md) §10 (pembatas pengingat), §11 (aturan dan daftar peristiwa N-01 sampai N-32), §16.8 (notifikasi chat). Rumusan akhir isi layar kunci dan mekanisme push: dokumen 10 dan 09 (S4) |
| Bukan isi berkas ini | Kebutuhan notifikasi chat N-23 ada di [COM.md](COM.md); berkas ini hanya memuat hubungannya dengan pembatas, ringkasan, dan layar kunci. Notifikasi keanggotaan N-24 sampai N-26: [FAM.md](FAM.md). Pemegang tempat uang N-27: [FIN-2.md](FIN-2.md) |
| Bukan isi dokumen ini | Kata-kata notifikasi, tampilan, dan layar pengaturan (dokumen 08) |

## Cara membaca

Sama dengan [REQ.md](REQ.md) bagian "Cara membaca": bentuk FR, AC lokal (`FR-NTF-011/AC-2`), pelaku "A" dan "B", label kedudukan, dan baris **Dua arah**. Tambahan untuk berkas ini:

- **Dilepas** berarti server menyerahkan pemberitahuan untuk dikirim ke perangkat penerima (06 §10). Model tidak menyimpan apakah perangkat menerimanya.
- **Jam tenang** dalam AC adalah bawaan 22.00 sampai 06.00 menurut zona waktu rumah tangga, kecuali disebut lain.
- **Objek** untuk pembatas: rantai permintaan, sasaran kesepakatan, baris custody, atau permintaan tindakan (06 §11.1).

## Daftar FR

| ID | Judul | Fitur |
|---|---|---|
| FR-NTF-001 | Peristiwa, penerima, dan prioritas | F-NTF-001 |
| FR-NTF-002 | Kegagalan kirim, tanpa tanda terima, keadaan terbaru | F-NTF-001 |
| FR-NTF-003 | Notifikasi yang tidak ada | F-NTF-001 |
| FR-NTF-004 | Prioritas peristiwa kepedulian (Tier 2) | F-NTF-001 |
| FR-NTF-005 | Teks umum di layar kunci | F-NTF-002 |
| FR-NTF-006 | Rincian atas pilihan penerima dan isi push | F-NTF-002 |
| FR-NTF-007 | Jam tenang bawaan dan pemiliknya | F-NTF-003 |
| FR-NTF-008 | Penahanan dan pelepasan sesudah jam tenang | F-NTF-003 |
| FR-NTF-009 | Isi ringkasan notifikasi | F-NTF-004 |
| FR-NTF-010 | Jadwal ringkasan notifikasi | F-NTF-004 |
| FR-NTF-011 | Jendela 6 jam bersama per objek per pengirim | F-NTF-005 |
| FR-NTF-012 | Satu yang tertahan per objek | F-NTF-005 |
| FR-NTF-013 | Plafon harian per pengirim kepada penerima | F-NTF-005 |
| FR-NTF-014 | Tanpa orakel | F-NTF-005 |
| FR-NTF-015 | Cakupan pembatas | F-NTF-005 |

## Business rule

| ID | Pernyataan | Sumber kanonis | Kedudukan |
|---|---|---|---|
| BR-NTF-001 | Tiga prioritas: `perlu_tanggapan` (penerima diminta bertindak), `kabar` (hasil atau perubahan yang perlu diketahui), `ringkasan` (boleh digabung ke ringkasan notifikasi) | 06 §11.1 | Dirancang 06 |
| BR-NTF-002 | Daftar peristiwa N-01 sampai N-32 beserta penerima dan prioritasnya kanonis di 06 §11.2 | 06 §11.2 | N-01 sampai N-22: dirancang 06; N-10 lihat BR-REQ-027; N-01 dan N-20 atas konfirmasi susulan `ringkasan` dan N-14 atas hitungan bukan pemegang `ringkasan`: **[Usulan]** SP-52, SP-50 (revisi kesepuluh). N-23 sampai N-31: **[Usulan]** `architect` (revisi kesembilan), kecuali yang disebut lain di BR-NTF-015 sampai BR-NTF-020. N-32: **[Usulan]** `architect` (revisi kesepuluh, CA-08), memindahkan pengingat tagihan ke Tier 1; bagian pertanyaan checkpoint S3-1 (06 P-36) |
| BR-NTF-003 | Pelaku tidak diberi tahu atas tindakannya sendiri; penerima hanya anggota yang boleh melihat objeknya | 06 §11.1 | Dirancang 06 |
| BR-NTF-004 | Notifikasi yang gagal tidak membatalkan transisi; state selalu terbaca di aplikasi; model tidak menyimpan atau menampilkan tanda terima notifikasi. Tanda terkirim dan dibaca di chat adalah keadaan pesan di chat, bukan tanda terima notifikasi; membuka notifikasi atau layar objek tidak menggerakkannya | 06 §11.1 "Tanpa tanda terima", AT-5, INV-A33, §16.5 | D-29 (lewat "lanjut"); batas dengan chat: revisi kesembilan, **[Usulan]** `architect` |
| BR-NTF-005 | Layar kunci bawaan hanya teks umum per kelas prioritas, tanpa jenis peristiwa, nama pelaku, atau nama rumah tangga; `ringkasan` tidak tampil sendiri-sendiri. Rincian hanya atas pilihan penerima per perangkat, dan tetap tanpa jumlah, judul, alasan, isi diskusi, atau rincian kepedulian. Push hanya memuat kelas prioritas dan pengenal buram. Kanal dan kategori notifikasi platform dibuat per kelas prioritas, bukan per jenis peristiwa. N-23 dan N-28 memakai teks umum kelas `kabar` yang sama persis dengan `kabar` lain ("Ada kabar baru"), tanpa nama dan tanpa isi; N-23 tetap demikian juga bila penerima memilih rincian. Bawaan itu menunggu checkpoint S3-8; alternatif (b): setiap `kabar` dan N-23 dari pasangan berbunyi "ada kabar dari {nama}", dengan `{nama}` nama tampilan pilihan pengirimnya sendiri (FR-FAM-006) | 06 §11.1 "Isi layar kunci", §16.8, N-28, P-32; `ENGINEERING-STANDARDS.md` §4 | D-34 (Pagar, lewat "lanjut"); rumusan akhir **[Dok. 10]**. Kanal per kelas: **[Usulan]** `security-privacy` SP-46. N-23: **[Usulan]** `architect`, pertanyaan checkpoint S3 (06 P-24). N-28: yang diterima pemilik produk atas saran lead (12 November 2026) adalah pilihan "Umum di layar kunci", yang teksnya berbunyi "ada kabar dari istrimu"; teks layar kunci yang pasti adalah pertanyaan checkpoint S3-8, dengan bawaan **[Usulan]** (a) "Ada kabar baru" (06 §11.1) |
| BR-NTF-006 | Setiap orang punya jam tenang sejak awal, 22.00 sampai 06.00 zona waktu rumah tangga; hanya ia yang menggeser atau mematikannya; setelannya tidak terbaca orang lain. Semua notifikasi Tier 1 ditahan selama jam tenang penerima | 06 §10 tabel batas, §11.1, INV-A42 | Untuk pengingat: keputusan pemilik nomor 13 (sendiri). Bawaan per orang: diterima atas saran lead (S2-5); angka **[Asumsi]** (A-28). Untuk notifikasi selain pengingat: **[Usulan]** (Q-132) |
| BR-NTF-007 | Paling banyak satu pemberitahuan tertahan per objek per penerima; yang datang kemudian melebur ke dalamnya dan menampilkan keadaan terbaru | 06 §10 baris "Yang tertahan", INV-A30 | D-29 (lewat "lanjut") |
| BR-NTF-008 | Semua yang tertahan selama jam tenang dilepas sebagai satu notifikasi saat jam tenang berakhir. Tanda terkirim di chat tidak bergantung pada push: ia naik hanya saat perangkat penerima menarik pesan ketika aplikasinya dibuka; push, Notification Service Extension, dan sinkron latar tidak pernah menggerakkannya; satu aturan di semua platform (BR-COM-009) | 06 §10 baris "Yang tertahan"; §16.5; INV-A33 | **[Usulan]** `security-privacy` SP-39 (satu notifikasi) dan SP-55 (arti terkirim), diterapkan 06 revisi kesepuluh. Bagian SP-39 "tanpa push apa pun, termasuk push senyap" **ditarik** `security-privacy`; rumusan lama BR ini dicabut. Mekanisme per platform **[Asumsi]** yang diverifikasi dokumen 09 |
| BR-NTF-009 | Antara dua pemberitahuan `perlu_tanggapan` atas satu objek dari satu pengirim kepada satu penerima yang dilepas, sekurangnya 6 jam; pengingat otomatis dihitung sebagai dari pihak yang menunggu | 06 §10 baris "Jendela 6 jam bersama", INV-A30 | 6 jam: keputusan pemilik nomor 13 (sendiri). Perluasan ke semua `perlu_tanggapan`: D-29 (lewat "lanjut") |
| BR-NTF-010 | Dalam 24 jam berjalan paling banyak enam `perlu_tanggapan` dari satu anggota kepada satu anggota lain dilepas langsung, lintas semua objek; kelebihannya ke ringkasan; tidak ada yang hilang dan tidak ada transisi yang ditolak | 06 §10 baris "Plafon harian", INV-A31 | Diterima atas saran lead (S2-5); enam **[Asumsi]** |
| BR-NTF-011 | Tunda senyap menahan pelepasan seperti jam tenang; penerimaan perintah pengirim hanya bergantung pada riwayat pengirim sendiri; hasil setiap perintah dan pembacaan pengirim identik apa pun keadaan penerima | 06 §10, INV-A33 | Diterima atas saran lead (S2-5); 24 jam **[Asumsi]** |
| BR-NTF-012 | N-11 atas usulan pos rutin dilepas langsung paling banyak satu per pengirim per hari, di dalam plafon harian | 06 §10, N-11 | **[Usulan]** `security-privacy` SP-36 |
| BR-NTF-013 | Ringkasan memuat butir `ringkasan` dan kelebihan plafon harian; semua butir N-10 dilebur menjadi satu baris tanpa hitungan dan tanpa nama | 06 §6.2 "Notifikasi", N-10, INV-A31 | Kelebihan plafon: S2-5. Satu baris: **[Usulan]** SP-43 butir b |
| BR-NTF-014 | Isi pengingat adalah kalimat baku aplikasi, tanpa teks bebas dan tanpa jumlah | 06 §10 baris "Isi" | D-29 (lewat "lanjut") |
| BR-NTF-015 | Chat tidak menggantikan notifikasi. Satu kejadian, satu notifikasi per penerima: kartu dan baris peristiwa di chat tidak memicu notifikasi sendiri; balasan yang menempel pada kartu memicu N-23, bukan N-19. N-19 tinggal untuk diskusi objek yang tidak tampil di chat. Tanda dibaca dan kehadiran tidak dibaca aturan notifikasi mana pun selain lipat N-23; tidak ada notifikasi "pasanganmu online" atau "dibaca tapi belum dijawab" | 06 §11.1 "Bentuk chat dan notifikasi", §16.5, §16.6, §16.8, N-19, INV-A45, INV-A46 | Tanpa "dibaca tapi belum dijawab": keputusan pemilik produk atas saran lead ("Ikut tanda dibaca"), 12 November 2026. Selebihnya **[Usulan]** `architect` |
| BR-NTF-016 | N-23 berprioritas `kabar`, dilipat per pengirim (paling banyak satu per 10 menit selama penerima belum membuka chat; satu yang tertahan per pengirim selama jam tenang), dan tidak masuk jendela 6 jam, plafon harian, atau batas pengingat ulang. Kebutuhannya: FR-COM-017 di COM.md | 06 §16.8, N-23 | **[Usulan]** `architect`; 10 menit **[Asumsi]**; pertanyaan checkpoint S3 (06 P-24) |
| BR-NTF-017 | `kabar` yang tertahan jam tenang atas objek yang sama dari pengirim yang sama dilebur menjadi satu, dengan isi menurut state objek saat dilepas; bila objeknya sudah tidak terlihat penerima saat dilepas, kabar itu gugur tanpa jejak | 06 §11.1 "Penggabungan `kabar`", V-7 | **[Usulan]** `architect` (revisi kesembilan) |
| BR-NTF-018 | Pengingat yang berpatokan tanggal tanpa jam: "sehari sebelum" dijadwalkan 00.00 hari sebelum tanggal itu, "pada tanggal" 00.00 tanggal itu, menurut zona waktu rumah tangga, lalu tunduk pada jam tenang penerima; dengan jam tenang bawaan keduanya dilepas 06.00. Server memegang jadwalnya | 06 §11.1 "Waktu pengingat yang berpatokan tanggal", S-6 | **[Usulan]** `architect` (revisi kesembilan) |
| BR-NTF-019 | Ringkasan notifikasi sekali sehari pukul 19.00 zona waktu rumah tangga, per penerima; hanya penerimanya yang menggesernya dan setelannya tidak terbaca orang lain; di dalam jam tenang dilepas saat jam tenang berakhir; yang kosong tidak dikirim; isinya state saat dilepas; butir yang dibuang menurut BR-NTF-020 | 06 §11.1 "Jadwal ringkasan", INV-A42 | **[Usulan]** `architect`, menjawab Q-133; 19.00 **[Asumsi]** |
| BR-NTF-020 | Saat ringkasan dilepas, yang dibuang hanya: (1) butir yang objeknya tidak lagi terlihat penerima, selalu; (2) butir yang **meminta tindakan atau mengingatkan** (pemberitahuan `perlu_tanggapan` yang masuk ringkasan karena plafon, N-01 dan N-20 konfirmasi susulan, pengingat mingguan permintaan tindakan, N-32) bila objeknya tidak lagi menunggu penerima itu. Butir yang memberitahukan hasil atau perubahan (N-05, N-09, N-14, N-15, N-17, dan sejenisnya) tidak dibuang karena objeknya sudah selesai; isinya memang hasil itu | 06 §11.1 "Jadwal ringkasan" (revisi kesepuluh), V-7 | **[Usulan]** `architect`, menjawab PA-NTF-3 |

---

## F-NTF-001 — Notifikasi alur Tier 1

### FR-NTF-001 — Peristiwa, penerima, dan prioritas

| | |
|---|---|
| Fitur | F-NTF-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`, `MSL-2` (yang perlu ditanggapi sampai tanpa harus bertanya); `KS-1` |
| Aktor dan izin | Server membuat notifikasi sebagai akibat samping transisi; penerima hanya anggota yang boleh melihat objeknya (06 §9) |
| Aturan acuan | 06 §11.1, §11.2 (N-01 sampai N-32), §16.8, INV-A46 |
| Business rule | BR-NTF-001, BR-NTF-002, BR-NTF-003, BR-NTF-015 |
| Kedudukan | N-01 sampai N-22: dirancang 06. N-23 sampai N-27: **[Usulan]** `architect` (revisi kesembilan). N-32 dan prioritas `ringkasan` atas konfirmasi susulan dan hitungan bukan pemegang: **[Usulan]** (revisi kesepuluh; lihat BR-NTF-002) |

Setiap transisi yang disebut 06 §11.2 menghasilkan peristiwanya, kepada penerima dan dengan prioritas di tabel itu. Dokumen ini tidak menambah peristiwa. Ringkasnya, untuk Tier 1:

| Prioritas | Peristiwa |
|---|---|
| `perlu_tanggapan` | N-01, N-02, N-03, N-11, N-13, N-20, N-24 (undangan diterima, menunggu pengundang mengukuhkan); Tier 2: permintaan tindakan baru. Pengecualian: N-01 dan N-20 atas konfirmasi susulan selalu `ringkasan` (FR-REQ-037/AC-1, FR-REQ-039/AC-5) |
| `kabar` | N-04 sampai N-08, N-12, N-16 sampai N-18, N-22, N-23 (pesan chat, lihat COM.md), N-25 (kamu sudah menjadi anggota), N-26 (anggota dewasa keluar; ditinjau dokumen 10) |
| `ringkasan` | N-09, N-15, N-19 (hanya diskusi yang tidak tampil di chat), N-21, N-27 (pemegang tempat uang berubah), N-32 (pengingat tagihan, tiga hari sebelum dan pada hari jatuh tempo; FR-FIN-100/AC-1) |
| Menurut dasar atau jenis | N-10 menurut dasarnya; N-14 `kabar` atau `ringkasan` menurut jenis entrinya, termasuk penilaian ulang aset dana bersama oleh pasangan, selisih negatif dari hitungan bukan pemegang (FR-FIN-135/AC-8), dan pengarsipan tujuan dana bersama (ketiganya `ringkasan`) |

Peristiwa Tier 2 (N-28 sampai N-31) ada di FR-NTF-004.

**Kriteria penerimaan**
- **AC-1** — *Given* A mengajukan permintaan, *When* server menerimanya, *Then* B menerima N-01 berprioritas `perlu_tanggapan` dan A tidak menerima notifikasi atas tindakannya sendiri.
- **AC-2** — *Given* B menyetujui permintaan A, *When* server menerimanya, *Then* A menerima N-04 berprioritas `kabar`; B tidak menerima apa pun.
- **AC-3** — *Given* permintaan bertanggal kedaluwarsa (B10), *When* `system` menjalankannya, *Then* A dan B sama-sama menerima N-07.
- **AC-4** — *Given* otorisasi ditutup otomatis (B16), *When* `system` menjalankannya, *Then* keduanya menerima N-09 berprioritas `ringkasan`; *Given* pemohon menyelesaikan sendiri (B15), *Then* hanya pihak lainnya yang menerima N-09.
- **AC-5** — *Given* B mempersempit siapa yang boleh melihat datanya, *When* perubahan tersimpan, *Then* tidak ada notifikasi kepada A; *Given* B memperluasnya, *Then* A menerima N-22 (D-40).
- **AC-6** — *Given* A mengajukan permintaan dan kartunya lahir di chat, *When* server menerimanya dan B kemudian menyetujuinya dari kartu, *Then* B menerima tepat satu notifikasi atas pengajuan (N-01) dan A tepat satu atas persetujuan (N-04); kartu dan baris peristiwa tidak menambah notifikasi apa pun (06 INV-A46).
- **AC-7** — *Given* aset dana bersama, *When* B menilai ulang nilainya, *Then* A menerima N-14 berprioritas `ringkasan` (06 §11.2 N-14, revisi kesembilan).

**Dua arah.** Simetris; penerima ditentukan peran terhadap objek, bukan orangnya.

### FR-NTF-002 — Kegagalan kirim, tanpa tanda terima, keadaan terbaru

| | |
|---|---|
| Fitur | F-NTF-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1`; `PAGAR` (notifikasi tidak menjadi petunjuk "terakhir aktif"; kehadiran hanya di layar chat, 06 §16.6) |
| Aktor dan izin | Tidak seorang pun melihat waktu dilepas, ditahan, diterima, atau dibuka notifikasi orang lain |
| Aturan acuan | 06 §11.1 "Kegagalan", "Tanpa tanda terima", §10 "Keterbukaan", AT-5, INV-A33, §16.5 |
| Business rule | BR-NTF-004 |
| Kedudukan | D-29 (lewat "lanjut"); batas dengan tanda chat: **[Usulan]** `architect` (revisi kesembilan) |

**Kriteria penerimaan**
- **AC-1** — *Given* layanan push gagal mengirim N-01 kepada B, *When* B membuka aplikasi, *Then* permintaan A terbaca menunggu tanggapan B dan state permintaan tidak terpengaruh kegagalan itu.
- **AC-2** — *Given* B membuka notifikasi dari A, *When* A membuka objek itu, *Then* A tidak melihat apakah atau kapan notifikasi itu sampai, ditampilkan, atau dibuka (Keluhan 7).
- **AC-3** — *Given* notifikasi tentang permintaan yang sudah berubah state sejak dibuat, *When* B membukanya, *Then* yang ditampilkan adalah keadaan terbaru objek itu, bukan keadaan saat notifikasi dibuat.
- **AC-4** — *Given* N-01 atas permintaan A yang kartunya ada di chat, *When* B membuka notifikasi itu dan membaca permintaannya di layar permintaan tanpa membuka chat, *Then* tanda dibaca kartu itu bagi A tidak bergerak (06 §16.5 "Dibaca", §11.1 "Tanpa tanda terima").

**Dua arah.** Simetris.

### FR-NTF-003 — Notifikasi yang tidak ada

| | |
|---|---|
| Fitur | F-NTF-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5`; `KS-2`; `PAGAR` (tidak ada notifikasi yang menilai orang) |
| Aktor dan izin | Berlaku atas semua orang |
| Aturan acuan | 02 PRN-08; 06 §11.2 (daftar tertutup), PG-07, PG-09, §16.5 "Yang tidak ada", §16.6 "Notifikasi"; 04 F-NTF-006, F-NTF-007 |
| Business rule | BR-NTF-002, BR-NTF-014, BR-NTF-015 |
| Kedudukan | Tanpa "kamu belum mencatat": **[Usulan]** PRN-08, Q-90 terbuka sampai checkpoint S3. Tanpa peringatan anggaran: Q-134 diusulkan dokumen 04. Preferensi hanya per prioritas, tidak per peristiwa: **[Usulan]** PRN-08 |

**Kriteria penerimaan**
- **AC-1** — *Given* A tidak mencatat apa pun selama tujuh hari, *When* waktu berjalan, *Then* tidak ada notifikasi kepada A atau B tentang itu.
- **AC-2** — *Given* sisa anggaran kategori tinggal sedikit atau negatif, *When* entri disimpan, *Then* tidak ada push; sisa anggaran terbaca di aplikasi saat mencatat.
- **AC-3** — *Given* peristiwa apa pun, *When* notifikasinya disusun, *Then* tidak ada kalimat yang menilai seseorang dan tidak ada hitungan per orang (misalnya "ini tanda ketiga minggu ini").
- **AC-4** — *Given* pilot, *When* seseorang membuka pengaturan notifikasi, *Then* tidak ada saklar per jenis peristiwa; yang dapat diatur hanya jam tenang (FR-NTF-007), rincian layar kunci (FR-NTF-006), dan jam ringkasan (FR-NTF-010). Setelan pendamping `AI` (Tier 2) ada di setelan `AI`, bukan di pengaturan notifikasi (FR-AI-005).
- **AC-5** — *Given* B sudah membaca pesan atau kartu dari A di chat dan tidak membalas berapa lama pun, atau B sedang online, *When* waktu berjalan, *Then* tidak ada notifikasi kepada siapa pun tentang itu: tidak ada "dibaca tapi belum dijawab", "pasanganmu online", atau "pasanganmu terakhir aktif …" (06 §16.5, §16.6, INV-A45).

**Dua arah.** Simetris.

### FR-NTF-004 — Peristiwa Tier 2: kepedulian, saran `AI`, tugas, dan rutinitas

| | |
|---|---|
| Fitur | F-NTF-001 (peristiwa milik F-CARE-001, F-CARE-002, F-AI-002, F-REQ-013, F-CAL-001) |
| Tier / MoSCoW | 2 / Should |
| Masalah (brief §1.1) | `MSL-6` (rutinitas terlupa; suasana hati pasangan tidak diketahui kecuali ditanya), `MSL-7`; `KS-5` |
| Aktor dan izin | Penerima N-28 hanya pasangan yang diizinkan melihat oleh pemilik datanya; N-29 hanya anggota yang menyalakannya untuk dirinya; N-30 penerima tugas; N-31 penanggung jawab rutinitas, atau keduanya bila "siapa saja" |
| Aturan acuan | 06 §11.1 "Isi layar kunci" (teks umum per kelas prioritas, berlaku juga bagi N-28 di bawah bawaan (a) checkpoint S3-8; nama pengirim hanya di bawah alternatif (b)), "Penggabungan `kabar`", "Jadwal ringkasan"; §11.2 N-28 sampai N-31; §13.1, §13.3, §13.4, §13.5, §13.6 |
| Business rule | BR-NTF-005, BR-NTF-006, BR-NTF-017, BR-NTF-019 |
| Kedudukan | Q-138 dijawab 06 §11.2 N-28 dan §13.5 (**[Usulan]** `architect`, mengikuti bawaan Q-138). Layar kunci N-28: yang diterima pemilik produk atas saran lead (12 November 2026) adalah pilihan "Umum di layar kunci", teksnya "ada kabar dari istrimu", dengan kata-kata dan saran yang tampil sesudah ponsel dibuka; teks layar kunci yang pasti adalah pertanyaan checkpoint S3-8, bawaan **[Usulan]** (a) "Ada kabar baru" (06 §11.1, P-32). Saran harian: kata-kata pemilik produk "minimal 1 atau 2 appresiasi atau dukungan perhari" adalah pernyataannya sendiri; bahwa `AI` yang menyarankan, rentang satu sampai dua, batas dua, dan waktunya adalah bacaan lead **[Usulan]** (Q-157; S3-8), angkanya **[Asumsi]**. Batas dua kabar bersaran per hari **[Asumsi]**. Rinciannya pertanyaan checkpoint S3 (06 P-27) |

Prioritas dan bentuk, menurut 06 §11.2:

| Peristiwa | Prioritas | Yang khas |
|---|---|---|
| N-28 status kepedulian yang dibagikan berubah (status suasana hati baru, laporan sakit, jawaban tindak lanjut) | `kabar`, tidak pernah `perlu_tanggapan` | Dilebur per objek selama tertahan; layar kunci bawaan "Ada kabar baru", sama dengan `kabar` lain, menunggu S3-8; kanal platform sama dengan `kabar` lain; berakhir sendiri, pencabutan, dan penghapusan tidak diberitahukan dan menggugurkan yang tertahan; bila syarat 06 §13.6 terpenuhi, membawa kata-kata menenangkan dan saran pesan, paling banyak dua kabar bersaran per penerima per hari |
| N-29 saran apresiasi atau dukungan harian (bacaan lead **[Usulan]**, Q-157) | Bukan pemberitahuan untuk saran pagi; saran kedua dilepas bersama ringkasan | Saran pagi tampil di sapaan pagi saat aplikasi pertama dibuka hari itu; saran kedua bersama ringkasan 19.00, atau sendiri pada jam itu bila ringkasan kosong; tidak pernah di jam tenang; tidak menyusul bila terlewat |
| N-30 tenggat permintaan tindakan tiba | `kabar`, satu kali | Di luar plafon harian; tidak dikirim bila tugas sudah ditutup (FR-REQ-045) |
| N-31 rutinitas berjam tiba | `kabar` pada jam itu | Satu per kemunculan; dalam jam tenang dilepas saat jam tenang berakhir dan gugur bila kemunculannya sudah `selesai`; rutinitas tanpa jam tetap `ringkasan` (FR-CAL-001); pengingat tagihan kini N-32 Tier 1, `ringkasan` (FR-FIN-100/AC-1) |

Tak satu pun dari keempatnya `perlu_tanggapan`, jadi tidak terhitung jendela 6 jam dan plafon harian. Pertanyaan tindak lanjut sakit (I02) dan pengingatnya hanya kepada yang sakit, `kabar`, tunduk pada jam tenang (06 §13.3).

**Kriteria penerimaan**
- **AC-1** — *Given* A membagikan suasana hatinya kepada B dan B memakai jam tenang bawaan, *When* A mengubahnya pukul 23.00 lalu sekali lagi pukul 23.30, *Then* B menerima tepat satu N-28 pukul 06.00, dengan teks layar kunci yang sama dengan `kabar` lain (FR-NTF-005/AC-5), yang saat dibuka menampilkan status A pukul 06.00 (06 §11.1 "Penggabungan `kabar`").
- **AC-2** — *Given* notifikasi kepedulian, *When* tampil di layar kunci atau di daftar kanal notifikasi pengaturan ponsel, *Then* tidak pernah memuat suasana hati, kata "sakit", rinciannya, atau nama kanal yang membedakannya dari `kabar` lain, juga bila B memilih rincian; dengan bawaan (a) tidak ada nama orang sama sekali (06 §11.1; SP-46; S3-8).
- **AC-3** — *Given* A memilih status "hari ini" pukul 21.00 yang N-28-nya tertahan karena B memasang jam tenang mulai 20.00, *When* status itu berakhir sendiri pada akhir hari sebelum jam tenang B selesai, *Then* N-28 itu gugur tanpa jejak dan B tidak menerima apa pun tentang berakhirnya (06 §13.5, V-7).
- **AC-4** — *Given* A melaporkan sakit tiga kali pada satu hari dengan saran `AI` aktif bagi keduanya, *When* N-28 ketiga dilepas kepada B, *Then* ia tetap terkirim tetapi tanpa saran, karena dua kabar bersaran sudah terpakai hari itu (06 §13.6 "Batas"; **[Asumsi]**).
- **AC-5** — *Given* B menyalakan saran harian dan ringkasan B hari itu kosong, *When* pukul 19.00 tiba di luar jam tenang B, *Then* saran kedua dilepas sendiri sebagai satu pemberitahuan berteks umum; *When* B tidak membuka aplikasi pagi itu, *Then* saran pagi tidak menyusul (06 §13.6).
- **AC-6** — *Given* rutinitas "menanak nasi" pukul 17.00 dengan penanggung jawab "siapa saja", *When* pukul 17.00 tiba, *Then* A dan B masing-masing menerima satu N-31 `kabar`, tanpa pengulangan. *Given* rutinitas pukul 05.30 dengan penanggung jawab B dan jam tenang bawaan B, *When* A menandai kemunculan itu selesai pukul 05.45, *Then* N-31 yang tertahan untuk B gugur dan B tidak menerima apa pun pukul 06.00 (06 §13.4).

**Dua arah.** Simetris; setiap peristiwa ditujukan menurut peran terhadap objek, dan saran `AI` hanya bagi yang menyalakannya untuk dirinya.

---

## F-NTF-002 — Aman di layar kunci

### FR-NTF-005 — Teks umum di layar kunci

| | |
|---|---|
| Fitur | F-NTF-002 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` (privasi dari siapa pun yang melihat ponsel, termasuk pasangan); PRN-08 |
| Aktor dan izin | Pemilik perangkat |
| Aturan acuan | 06 §11.1 "Isi layar kunci" (SP-07, SP-46, kabar N-28), §16.8 (N-23), P-32 |
| Business rule | BR-NTF-005 |
| Kedudukan | D-34 (Pagar, lewat "lanjut"); rumusan akhir **[Dok. 10]** (S4). N-23 dan N-28, serta alternatif (b) di S3-8: lihat BR-NTF-005 |

**Kriteria penerimaan**
- **AC-1** — *Given* bawaan, *When* B menerima N-01 dari A, *Then* layar kunci B hanya menampilkan teks umum kelas `perlu_tanggapan`, misalnya "Ada yang menunggu jawabanmu", tanpa jenis peristiwa, tanpa nama A, tanpa nama rumah tangga, tanpa jumlah.
- **AC-2** — *Given* bawaan, *When* A menerima N-16 (keberatan diajukan), *Then* layar kunci hanya menampilkan teks umum kelas `kabar`, misalnya "Ada kabar baru"; kata "keberatan" tidak tampil (Keluhan 10).
- **AC-3** — *Given* butir berprioritas `ringkasan`, *When* dibuat, *Then* tidak tampil sendiri-sendiri di layar kunci.
- **AC-4** — *Given* B menyalakan rincian di perangkatnya, *When* B menerima N-23 dari A, *Then* layar kunci tetap teks umum kelas `kabar` ("Ada kabar baru"), tanpa nama A dan tanpa isi pesan (06 §16.8; P-24). Di bawah alternatif (b) checkpoint S3-8 satu-satunya bedanya nama tampilan A sendiri; isi pesan tetap tidak pernah tampil.
- **AC-5** — *Given* bawaan **[Usulan]** (a), menunggu S3-8, *When* B menerima N-28 atas status A dan, pada waktu lain, N-04 atas permintaan B, *Then* teks layar kunci, kanal atau kategori notifikasi, dan bentuk payload keduanya identik ("Ada kabar baru"), tanpa nama, jenis peristiwa, atau isi; kata-kata menenangkan dan saran tampil hanya sesudah ponsel dibuka (06 §11.1 "Isi layar kunci", §13.6 "Layar kunci"; SP-46). Bila pemilik produk memilih (b) di S3-8, setiap `kabar` dan N-23 dari A berbunyi "ada kabar dari {nama}" dengan `{nama}` nama tampilan pilihan A sendiri (FR-FAM-006), tetap sama untuk N-28 dan `kabar` lain; AC ini lalu diperbarui.

**Dua arah.** Simetris.

### FR-NTF-006 — Rincian atas pilihan penerima dan isi push

| | |
|---|---|
| Fitur | F-NTF-002 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` |
| Aktor dan izin | Hanya pemilik perangkat yang menyalakan atau mematikan rincian, per perangkat |
| Aturan acuan | 06 §11.1 "Isi layar kunci"; `ENGINEERING-STANDARDS.md` §4 |
| Business rule | BR-NTF-005 |
| Kedudukan | D-34. Mekanisme pengambilan rincian sesudah autentikasi: dokumen 09 dan 10 (S4) |

**Kriteria penerimaan**
- **AC-1** — *Given* B menyalakan rincian di perangkatnya, *When* B menerima N-01, *Then* rincian boleh menyebut jenis peristiwa, tetapi tidak pernah jumlah uang, judul, alasan, isi diskusi, atau rincian kepedulian.
- **AC-2** — *Given* A tidak menyalakan rincian, *When* B menyalakannya di perangkat B, *Then* layar kunci A tidak berubah; *When* B mencoba mengatur rincian di perangkat A, *Then* tidak ada jalan untuk itu.
- **AC-3** — *Given* push apa pun, *When* isinya diperiksa di jalur kirim, *Then* hanya memuat kelas prioritas dan pengenal buram; rincian, bila dipilih, diambil perangkat dari server sesudah terautentikasi.

**Dua arah.** Simetris.

---

## F-NTF-003 — Jam tenang per orang

### FR-NTF-007 — Jam tenang bawaan dan pemiliknya

| | |
|---|---|
| Fitur | F-NTF-003 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; PRN-08 |
| Aktor dan izin | Hanya pemilik jam tenang. Setelan berlevel `private` |
| Aturan acuan | 06 §10 baris "Bawaan jam tenang", INV-A42, INV-A33 |
| Business rule | BR-NTF-006, BR-NTF-011 |
| Kedudukan | Lihat BR-NTF-006. Tier 1 berbeda dari teks brief: Q-135 |

**Kriteria penerimaan**
- **AC-1** — *Given* B anggota baru yang tidak pernah membuka pengaturan, *When* notifikasi untuk B dibuat pukul 23.00, *Then* ia dilepas pukul 06.00 atau sesudahnya; *When* dibuat pukul 21.59, *Then* dilepas seketika.
- **AC-2** — *Given* B mematikan jam tenangnya, *When* notifikasi dibuat pukul 23.00, *Then* dilepas seketika; tunda senyap dan plafon tetap berlaku.
- **AC-3** — *Given* B menggeser jam tenangnya ke 23.00 sampai 07.00, *When* notifikasi dibuat pukul 22.30, *Then* dilepas seketika; *When* dibuat pukul 06.30, *Then* dilepas pukul 07.00; yang sedang tertahan dilepas menurut setelan baru.
- **AC-4** — *Given* jam tenang B, *When* A mencoba melihat atau mengubahnya, *Then* tidak ada jalan untuk itu; perubahan oleh B tidak menghasilkan Peristiwa yang terbaca A atau notifikasi.

**Dua arah.** Simetris; tiap orang punya jam tenangnya sendiri.

### FR-NTF-008 — Penahanan dan pelepasan sesudah jam tenang

| | |
|---|---|
| Fitur | F-NTF-003 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; `MSL-5` (tidak dibanjiri) |
| Aktor dan izin | Server |
| Aturan acuan | 06 §10 baris "Jam tenang", "Yang tertahan", §11.1 "Jam tenang", "Penggabungan `kabar`", "Waktu pengingat yang berpatokan tanggal"; §16.5 "Jam tenang dan tunda" |
| Business rule | BR-NTF-006, BR-NTF-007, BR-NTF-008, BR-NTF-017, BR-NTF-018 |
| Kedudukan | Pengingat: keputusan 13. Notifikasi lain: **[Usulan]** Q-132. Satu notifikasi gabungan: **[Usulan]** SP-39 (bagian "tanpa push senyap" ditarik `security-privacy`). Terkirim di chat tidak bergantung pada push: **[Usulan]** SP-55 (06 revisi kesepuluh). Penggabungan `kabar` dan waktu pengingat bertanggal: **[Usulan]** `architect` (revisi kesembilan) |

**Kriteria penerimaan**
- **AC-1** — *Given* jalan-jalan kedaluwarsa pada pergantian hari 00.00, *When* N-07 dibuat, *Then* ia ditahan sampai jam tenang masing-masing penerima berakhir; tidak ada peristiwa Tier 1 yang menembus jam tenang.
- **AC-2** — *Given* selama jam tenang B tertahan N-02 atas satu permintaan, N-13 atas satu serah terima, dan N-10 `kabar`, *When* jam tenang berakhir pukul 06.00, *Then* **[Usulan]** B menerima satu notifikasi, yang saat dibuka menampilkan keadaan terbaru tiap objek.
- **AC-3** — *Given* jam tenang B berlaku dan selama itu tertahan beberapa notifikasi untuk B, *When* jam tenang B berakhir, *Then* **[Usulan]** semua yang tertahan dilepas sebagai satu notifikasi; tidak ada pemberitahuan yang tampil atau berbunyi di perangkat B selama jam tenang (SP-39; rumusan lama "tanpa push apa pun, termasuk push senyap" dicabut karena ditarik `security-privacy`).
- **AC-4** — *Given* penahanan, *When* state objek berubah selama jam tenang, *Then* perubahan state tidak tertunda; yang ditahan hanya pemberitahuannya.
- **AC-5** — *Given* selama jam tenang B, A menolak permintaan P milik B (N-05) dan menyetujui permintaan Q milik B (N-04), *When* jam tenang berakhir, *Then* kedua kabar itu tidak dilebur satu sama lain karena objeknya berbeda, dan keduanya terkandung dalam satu notifikasi yang dilepas (AC-2); *Given* dua `kabar` atas objek yang sama dari A tertahan, *When* dilepas, *Then* keduanya menjadi satu dengan isi menurut state objek saat dilepas; *When* objek itu sudah tidak terlihat B saat dilepas, *Then* kabar itu gugur tanpa jejak (06 §11.1 "Penggabungan `kabar`").
- **AC-6** — *Given* jalan-jalan bertanggal 14 Desember dan pengingat "sehari sebelum" atasnya, *When* 13 Desember pukul 00.00 tiba menurut zona waktu rumah tangga, *Then* pengingat dijadwalkan dan, dengan jam tenang bawaan penerima, dilepas pukul 06.00 tanggal 13 Desember (06 §11.1; FR-REQ-025/AC-6).
- **AC-7** — *Given* A mengirim pesan chat sekali saat B dalam jam tenang dan sekali saat B di luar jam tenang, dengan aplikasi B tertutup pada keduanya dan B membuka aplikasinya pada jarak waktu yang sama sesudah pesan dikirim, *When* tanda terkirim dibandingkan, *Then* tanda itu bergerak pada saat yang sama relatif terhadap pembukaan aplikasi B dan tidak lebih awal pada keadaan di luar jam tenang, karena push yang dilepas tidak menggerakkannya; A tidak dapat menyimpulkan jam tenang B dari tanda terkirim (06 §16.5 "Terkirim", "Jam tenang dan tunda"; INV-A33; FR-COM-011/AC-6, AC-10; SP-55).

**Dua arah.** Simetris.

---

## F-NTF-004 — Ringkasan notifikasi

### FR-NTF-009 — Isi ringkasan notifikasi

| | |
|---|---|
| Fitur | F-NTF-004 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5` (tidak dibanjiri); `PAGAR` |
| Aktor dan izin | Ringkasan milik penerimanya; isinya hanya objek yang boleh ia lihat |
| Aturan acuan | 06 §11.1 "Prioritas", "Jadwal ringkasan", §6.2 "Notifikasi", N-01 dan N-20 (konfirmasi susulan), N-09, N-10, N-14, N-15, N-19, N-21, N-27, N-32, INV-A31; §13.1 (baris tugas mingguan), §13.4 (rutinitas tanpa jam dan tagihan) |
| Business rule | BR-NTF-013, BR-NTF-010, BR-NTF-019, BR-NTF-020 |
| Kedudukan | Isi: dirancang 06. Satu baris tanpa hitungan untuk N-10: **[Usulan]** SP-43 butir b. Isi menurut state saat dilepas dan butir yang dibuang: **[Usulan]** `architect` (revisi kesembilan dan kesepuluh) |

Ringkasan memuat butir berprioritas `ringkasan` sejak ringkasan sebelumnya, ditambah pemberitahuan `perlu_tanggapan` yang melewati plafon harian. Isinya state saat ringkasan dilepas, bukan saat butirnya lahir. Yang dibuang hanya butir yang objeknya tidak lagi terlihat penerima, dan butir yang meminta tindakan atau mengingatkan bila objeknya tidak lagi menunggu penerima itu; butir yang memberitahukan hasil tetap ada (BR-NTF-020). Pesan chat tidak masuk ringkasan; ia diberitahukan lewat N-23 (COM.md).

**Kriteria penerimaan**
- **AC-1** — *Given* sejak ringkasan sebelumnya ada tiga entri A yang bertanda (satu dasar jatah, dua dasar ambang yang kedua dan ketiga hari itu), *When* ringkasan B disusun, *Then* ketiganya menjadi satu baris yang menyatakan ada catatan bertanda, tanpa angka dan tanpa nama (Keluhan 1).
- **AC-2** — *Given* pemberitahuan `perlu_tanggapan` ketujuh dari A kepada B dalam 24 jam, *When* ia tidak dilepas langsung, *Then* ia ada di ringkasan B dan objeknya ada di daftar "menunggu saya" B.
- **AC-3** — *Given* serah terima terkonfirmasi (N-15), kabar "nanti aku jawab" (N-21), dan perubahan pemegang tempat uang (N-27), *When* ringkasan disusun, *Then* ketiganya ada di ringkasan dan tidak satu pun dilepas sebagai pemberitahuan langsung; *Given* B membalas kartu permintaan A di chat, *When* ringkasan A disusun, *Then* balasan itu tidak ada di ringkasan, karena ia sudah diberitahukan lewat N-23 (06 N-19, INV-A46).
- **AC-4** — *Given* N-01 ketujuh dari A kepada B pada hari itu masuk ringkasan B (AC-2), *When* B sudah menyetujui permintaan itu sebelum pukul 19.00, *Then* butir itu tidak ada di ringkasan B yang dilepas pukul 19.00 (06 §11.1 "Jadwal ringkasan"); *Given* butir N-09 atas permintaan yang sudah diselesaikan, butir N-15 atas serah terima yang sudah dikonfirmasi, dan butir N-17 atas keberatan yang sudah ditutup, *When* ringkasan dilepas, *Then* ketiganya tetap ada dan menampilkan keadaan terbarunya, karena yang diberitahukannya adalah hasil itu sendiri (BR-NTF-020; PA-NTF-3 ditutup).
- **AC-5** — *Given* butir N-32 atas kejadian tagihan "listrik" jatuh tempo 12 Desember masuk antrean ringkasan B pada 9 Desember, *When* pembayaran yang ditautkan ke kejadian itu tersimpan pukul 15.00 dan ringkasan dilepas pukul 19.00, *Then* butir itu tidak ada di ringkasan; *When* yang dibayar hanya kejadian Januari, *Then* butir kejadian Desember tetap ada (BR-NTF-020; 06 §13.4; 07 §3.15).

**Dua arah.** Simetris; ringkasan dihitung per penerima.

### FR-NTF-010 — Jadwal ringkasan notifikasi

| | |
|---|---|
| Fitur | F-NTF-004 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5`; `PAGAR` |
| Aktor dan izin | Hanya pemiliknya yang menggeser jam ringkasannya; setelan tidak terbaca orang lain |
| Aturan acuan | 06 §11.1 "Jadwal ringkasan", INV-A42; 04 §10.1 |
| Business rule | BR-NTF-013, BR-NTF-019 |
| Kedudukan | Q-133 dijawab 06 §11.1 (**[Usulan]** `architect`, mengikuti bawaan Q-133); 19.00 **[Asumsi]** |

Ringkasan dikirim sekali sehari pukul 19.00 zona waktu rumah tangga, per penerima; hanya pemiliknya yang dapat menggesernya; bila jatuh di jam tenangnya, dilepas saat jam tenang berakhir; tidak dikirim bila kosong. Saran apresiasi harian kedua (Tier 2, N-29) menumpang pada ringkasan ini (FR-NTF-004).

**Kriteria penerimaan**
- **AC-1** — *Given* B tidak mengubah apa pun dan ada butir ringkasan, *When* pukul 19.00 tiba, *Then* B menerima satu notifikasi ringkasan dengan teks umum di layar kunci.
- **AC-2** — *Given* tidak ada butir sejak ringkasan sebelumnya, atau semua butirnya dibuang karena objeknya sudah selesai atau tidak lagi terlihat, *When* pukul 19.00 tiba, *Then* tidak ada yang dikirim.
- **AC-3** — *Given* B menggeser jam tenangnya ke 18.00 sampai 06.00, *When* pukul 19.00 tiba, *Then* ringkasan dilepas pukul 06.00 dengan isi menurut state pukul 06.00.
- **AC-4** — *Given* B menggeser jam ringkasannya, *When* A membuka aplikasi, *Then* A tidak dapat mengetahui jam ringkasan B.

**Dua arah.** Simetris.

---

## F-NTF-005 — Pembatas pemberitahuan yang meminta tanggapan

Seluruh fitur ini berkedudukan: jarak 6 jam dan jam tenang keputusan pemilik nomor 13 (sendiri); batas tiga, plafon enam, dan tunda 24 jam diterima atas saran lead (S2-5), angkanya **[Asumsi]**; cara menghitung dan perluasan cakupan D-29 (lewat "lanjut"). Penahanan dan plafon tidak pernah menunda perubahan state.

### FR-NTF-011 — Jendela 6 jam bersama per objek per pengirim

| | |
|---|---|
| Fitur | F-NTF-005 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `PAGAR` (tidak ditagih); arah 2 pemilik produk |
| Aktor dan izin | Server |
| Aturan acuan | 06 §10 baris "Jendela 6 jam bersama", INV-A30; §11.1 "Pemberitahuan yang meminta tanggapan" |
| Business rule | BR-NTF-009 |
| Kedudukan | Lihat pembuka fitur |

Jendela mencakup semua `perlu_tanggapan` dari satu pengirim atas satu objek: pengajuan, revisi, usul balik, usulan dan usulan ulang perubahan syarat, pengajuan ulang dalam rantai yang sama, pengingat ulang, dan pengingat otomatis atas nama pihak yang menunggu.

**Kriteria penerimaan**
- **AC-1** — *Given* N-01 dari A atas permintaan P dilepas kepada B pukul 08.00, *When* A merevisi syaratnya pukul 09.00, *Then* state berubah seketika dan B dapat membaca versi baru di aplikasi, tetapi N-03 baru dilepas pukul 14.00.
- **AC-2** — *Given* AC-1, *When* B mengajukan usul balik atas P pukul 10.00, *Then* N-02 kepada A dilepas seketika: jendela dihitung per pengirim.
- **AC-3** — *Given* AC-1, *When* A mengajukan permintaan lain Q pukul 09.30, *Then* N-01 atas Q dilepas seketika: jendela dihitung per objek (tunduk pada plafon harian).
- **AC-4** — *Given* A membatalkan P pukul 10.00 dan mengajukannya lagi pukul 10.05, *When* N-01 rantai itu dibuat, *Then* ia dilepas paling cepat pukul 14.00, karena pengajuan ulang berada di rantai yang sama.

**Dua arah.** Ditulis untuk kedua arah (AC-1 dari A, AC-2 dari B); aturan per pengirim.

### FR-NTF-012 — Satu yang tertahan per objek

| | |
|---|---|
| Fitur | F-NTF-005 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `PAGAR`; `MSL-5` |
| Aktor dan izin | Server |
| Aturan acuan | 06 §10 baris "Yang tertahan", "Mengapa 'ingatkan lagi' tidak ditolak", INV-A30 |
| Business rule | BR-NTF-007 |
| Kedudukan | Lihat pembuka fitur |

**Kriteria penerimaan**
- **AC-1** — *Given* N-03 dari A atas P tertahan sampai 14.00, *When* A menekan "ingatkan lagi" pukul 11.00 dan merevisi lagi pukul 12.00, *Then* pukul 14.00 B menerima tepat satu pemberitahuan atas P yang saat dibuka menampilkan versi terbaru.
- **AC-2** — *Given* AC-1, *When* A melihat permintaannya, *Then* pengingat ulang pukul 11.00 tetap terhitung terhadap batas tiga A walau melebur.

**Dua arah.** Simetris.

### FR-NTF-013 — Plafon harian per pengirim kepada penerima

| | |
|---|---|
| Fitur | F-NTF-005 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `PAGAR` (sepuluh objek tidak berarti sepuluh kali batas) |
| Aktor dan izin | Server |
| Aturan acuan | 06 §10 baris "Plafon harian", INV-A31 |
| Business rule | BR-NTF-010, BR-NTF-012 |
| Kedudukan | Lihat pembuka fitur; N-11 pos rutin: **[Usulan]** SP-36 |

**Kriteria penerimaan**
- **AC-1** — *Given* A membuat tujuh objek yang menunggu tanggapan B pada 08.00, 09.00, sampai 14.00 di hari yang sama, di luar jam tenang B, *When* pemberitahuannya dibuat, *Then* enam pertama dilepas langsung, yang ketujuh masuk ringkasan B, dan ketujuh objek berubah state seketika.
- **AC-2** — *Given* AC-1, *When* pukul 08.00 hari berikutnya tiba, *Then* jendela 24 jam berjalan memberi ruang lagi untuk satu pemberitahuan langsung dari A kepada B.
- **AC-3** — *Given* pengingat otomatis atas objek A kepada B, *When* plafon dihitung, *Then* pengingat itu terhitung sebagai dari A.
- **AC-4** — *Given* A mengusulkan tiga pos rutin pada satu hari, *When* N-11-nya dibuat, *Then* hanya satu dilepas langsung dan dua masuk ringkasan, dan yang satu itu terhitung di dalam plafon enam.

**Dua arah.** Simetris; plafon per pasangan pengirim dan penerima, ke kedua arah.

### FR-NTF-014 — Tanpa orakel

| | |
|---|---|
| Fitur | F-NTF-005 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `PAGAR` (pembatas tidak boleh menjadi petunjuk "pasangan sedang tidak aktif") |
| Aktor dan izin | Pengirim hanya melihat riwayatnya sendiri |
| Aturan acuan | 06 INV-A33, §10 baris "Tunda senyap", "Keterbukaan" |
| Business rule | BR-NTF-011 |
| Kedudukan | Lihat pembuka fitur |

**Kriteria penerimaan**
- **AC-1** — *Given* rangkaian perintah A yang sama (ajukan, ingatkan lagi, revisi), *When* dijalankan tiga kali dengan B menunda, B dalam jam tenang, dan B tanpa keduanya, *Then* semua tanggapan server dan semua yang terbaca A identik.
- **AC-2** — *Given* A menekan "ingatkan lagi" saat B menunda, *When* perintah diproses, *Then* diterima, dihitung, dan ditahan; tidak ditolak dan tidak menampilkan "sedang ditahan" (Keluhan 5).
- **AC-3** — *Given* pemberitahuan dari A tertahan di sisi B, *When* A membuka permintaannya, *Then* A melihat jumlah pengingat dan waktu ia menekan, tidak waktu dilepas.

**Dua arah.** Simetris.

### FR-NTF-015 — Cakupan pembatas

| | |
|---|---|
| Fitur | F-NTF-005 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `PAGAR` |
| Aktor dan izin | Server |
| Aturan acuan | 06 §10 baris "Berlaku atas", "Batas jumlah pengingat ulang" (SP-27); §11.1 |
| Business rule | BR-NTF-009, BR-NTF-010, BR-NTF-011 |
| Kedudukan | Lihat pembuka fitur |

**Kriteria penerimaan**
- **AC-1** — *Given* A mencatat serah terima ke tempat uang yang dipegang B, *When* A menekan "ingatkan lagi" atas catatan yang belum dikonfirmasi itu, *Then* berlaku jendela 6 jam, batas tiga, jam tenang, dan plafon yang sama dengan permintaan.
- **AC-2** — *Given* A mengusulkan jatah A Rp1.000.000, B menolak, lalu A mengusulkan Rp900.000 atas jadwal yang sama, *When* A menekan "ingatkan lagi", *Then* hitungan tiga dilanjutkan per sasaran kesepakatan, tidak dimulai lagi.
- **AC-3** — *Given* perubahan syarat dari B menunggu jawaban A, atau (Tier 2) permintaan tindakan dari B menunggu A, *When* B menekan "ingatkan lagi", *Then* batas yang sama berlaku.
- **AC-4** — *Given* pesan diskusi atau pesan chat, termasuk balasan pada kartu, *When* plafon dan jendela dihitung, *Then* pesan itu tidak terhitung: diskusi di luar chat berprioritas `ringkasan` (N-19), pesan chat `kabar` dengan lipatnya sendiri (N-23; 06 §16.8).
- **AC-5** — *Given* N-28, N-29, N-30, atau N-31 kepada B (Tier 2), *When* plafon harian dan jendela 6 jam untuk pengirim mana pun dihitung, *Then* peristiwa itu tidak terhitung, karena tak satu pun `perlu_tanggapan` (06 §11.2, §13.1).

**Dua arah.** Ditulis untuk kedua arah (AC-1, AC-2 dari A; AC-3 dari B).

---

## Usulan yang menunggu

Pengenal butir mengikuti skema README `PA-NTF-n` (fase D; sebelumnya "#n").

### Yang sudah dijawab (06 revisi kesembilan dan kesepuluh)

| ID | Pertanyaan | Jawaban | Kedudukan |
|---|---|---|---|
| Q-133 | Kapan ringkasan notifikasi dikirim? | 06 §11.1 "Jadwal ringkasan", sama dengan bawaan yang diusulkan, ditambah: isi menurut state saat dilepas | **[Usulan]** `architect`; 19.00 **[Asumsi]** |
| Q-138 | Prioritas notifikasi atas suasana hati yang dibagikan dan status sakit | 06 §11.2 N-28 `kabar`, sama dengan bawaan. Teks layar kuncinya pertanyaan checkpoint S3-8 (bawaan **[Usulan]** (a) "Ada kabar baru"); yang diterima pemilik produk atas saran lead adalah pilihan "Umum di layar kunci", teksnya "ada kabar dari istrimu" | **[Usulan]** `architect` |
| — | Penggabungan `kabar` yang tertahan (tier-2.md FR-CARE-003/AC-2) | 06 §11.1 "Penggabungan `kabar`" | **[Usulan]** `architect` |
| — | Waktu pengingat yang berpatokan tanggal | 06 §11.1 | **[Usulan]** `architect` |
| PA-NTF-3 | Butir ringkasan yang "objeknya sudah selesai" lawan butir yang memberitahukan hasil (N-09, N-15, N-17) | 06 §11.1 revisi kesepuluh: yang dibuang hanya butir yang meminta tindakan atau mengingatkan bila objeknya tidak lagi menunggu penerima, dan butir yang tidak lagi terlihat; butir hasil tetap (BR-NTF-020) | **[Usulan]** `architect` |
| PA-NTF-4 | Tanda terkirim yang sama di dalam dan di luar jam tenang lawan larangan push senyap (SP-39) | 06 §16.5 revisi kesepuluh: terkirim naik hanya oleh penarikan saat aplikasi dibuka, satu aturan di semua platform; larangan push senyap ditarik `security-privacy` (SP-55). Cabang cadangan sudah menjadi aturannya (BR-NTF-008, FR-NTF-008/AC-7) | **[Usulan]** `security-privacy`, diterapkan `architect`; mekanisme platform **[Asumsi]** dokumen 09 |

### Yang masih terbuka

| ID | Pertanyaan | Default yang diusulkan | Dipakai di | Pemilik aturan |
|---|---|---|---|---|
| PA-NTF-1 | Notifikasi chat N-23: lipat 10 menit per pengirim, layar kunci tanpa nama pengirim | Ya; angka diuji di pilot; teks layar kunci bersama `kabar` lain masuk S3-8 | BR-NTF-005, BR-NTF-016 | Checkpoint S3 (06 P-24, P-32) |
| PA-NTF-2 | Beban pemberitahuan pendamping `AI`: dua kabar bersaran per hari, saran harian satu pagi tanpa pemberitahuan dan satu bersama ringkasan (bacaan lead **[Usulan]**, Q-157) | Ya; angka dan jam diuji di pilot | FR-NTF-004 | Checkpoint S3-8 (06 P-27) |
