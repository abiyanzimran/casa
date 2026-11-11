# 05 — Kebutuhan Fungsional: `FAM` (Tier 1)

| | |
|---|---|
| Pemilik | `product-analyst` |
| Tahap | S3, fase B (12 November 2026); putaran penyelarasan dengan revisi kesembilan dokumen 06 (§3.1 baris hubungan, §3.5, §3.6 F01 sampai F06, PG-13, INV-A48, N-24 sampai N-26, AT-5, §16) dan 07 §8.2 butir 2a, 12 November 2026; putaran perbaikan fase D (12 November 2026) atas 06 revisi kesepuluh (§3.5, §3.6 F06 sampai F08, masa tenggang keluar, INV-A48 (d)) dan 07 §8.2 butir 1a, 2b, 3a, temuan `security-privacy` SP-48 dan SP-49, dan `critic` A CA-06, CA-07, CA-09 |
| Status | Ditinjau `critic` A dan diperiksa `security-privacy` (fase D); perbaikan fase D belum ditinjau ulang |
| Bagian dari | [05-functional-requirements.md](../05-functional-requirements.md) (indeks, format blok FR, kebutuhan lintas modul `LM-nn`) |
| Fitur yang dicakup | F-FAM-001 sampai F-FAM-008 ([04-feature-map.md](../04-feature-map.md) §3.1). F-FAM-009 sampai 017 Tier 3: [tier-3.md](tier-3.md) |
| Dibaca sebelum menulis | `CLAUDE.md`; brief §1.1, §4, §5A, §5O, §8, §12, §16; [00-glossary.md](../00-glossary.md) bagian 3 sampai 6 dan 8; [_progress.md](../_progress.md); dokumen [04](../04-feature-map.md) seluruhnya; [06](../06-activity-approval-architecture.md) §2.5, §3, §9, §10, §11, §13, §14; [07](../07-financial-domain.md) §2.2, §3.2, §3.3, §6.1 sampai §6.3, §8; [ADR-007](../adr/ADR-007-append-only-dan-hak-penghapusan.md); [02](../02-product-principles.md); [03](../03-personas-household-models.md); [17](../17-open-questions-assumptions.md) keputusan 1 sampai 25, kelompok (a) sampai (i), Q-04, Q-25, Q-28, Q-30, Q-76, Q-99, Q-120 sampai Q-123, D-23, D-28, D-31 sampai D-45 |

**Batas berkas ini.** Berkas ini menyebut apa yang harus dapat dilakukan dan dilihat anggota. Ia tidak menetapkan aturan uang (kanonis di 07), aturan alur dan pagar (kanonis di 06), keamanan undangan, autentikasi, pemulihan akun, atau isi teks hukum (dokumen 10, S4), dan tidak menetapkan kata-kata layar atau tata letak (dokumen 08, S4). Kata di dalam tanda kutip pada kriteria penerimaan adalah isi yang harus tersampaikan, bukan kata layar.

**Arah.** Semua kebutuhan di berkas ini simetris: ditulis untuk "anggota dewasa", tidak membaca hubungan, jenis kelamin, penghasilan, pemegang uang, atau pembuat rumah tangga (06 PG-01; PRN-09; `LM-01`). Kriteria yang tidak ditulis dua kali menyatakan kesimetrisannya dan diuji dengan menukar kedua anggota dewasa (06 INV-A13).

---

## Daftar isi

1. [Ringkasan kebutuhan](#1-ringkasan-kebutuhan)
2. [Business rule (penunjuk)](#2-business-rule-penunjuk)
3. [Rumah tangga dan keanggotaan](#3-rumah-tangga-dan-keanggotaan)
4. [Wewenang, peran, dan visibilitas](#4-wewenang-peran-dan-visibilitas)
5. [Keluar, ekspor, dan penghapusan](#5-keluar-ekspor-dan-penghapusan)
6. [Pihak luar, pemberitahuan operator, dan fondasi](#6-pihak-luar-pemberitahuan-operator-dan-fondasi)
7. [Diagram keanggotaan](#7-diagram-keanggotaan)
8. [Yang tidak diatur di sini: [Perlu aturan], pertanyaan, dan permintaan](#8-yang-tidak-diatur-di-sini)

---

## 1. Ringkasan kebutuhan

Semua Tier 1 / Must.

| FR | Judul | Fitur | Masalah |
|---|---|---|---|
| FR-FAM-001 | Membuat rumah tangga | F-FAM-001 | `MSL-1`, `MSL-2` (fondasi); `KS-1` |
| FR-FAM-002 | Rumah tangga dengan satu anggota dewasa | F-FAM-001 | `MSL-1`; `PAGAR` |
| FR-FAM-003 | Membuat, membagikan, dan membatalkan undangan | F-FAM-002 | `MSL-1`, `MSL-2` |
| FR-FAM-004 | Menerima undangan dan mengaktifkan keanggotaan | F-FAM-002 | `MSL-1`, `MSL-2`; `PAGAR` |
| FR-FAM-005 | Yang berlaku saat anggota dewasa kedua bergabung | F-FAM-002 | `MSL-1`; `KS-1` |
| FR-FAM-006 | Profil anggota minimal | F-FAM-003 | Fondasi (`KS-3`: nama pada keterangan pelaku) |
| FR-FAM-007 | Wewenang setara tanpa layar pengaturan izin | F-FAM-004 | `PAGAR`; `MSL-4` |
| FR-FAM-008 | Yang tidak pernah dapat dilihat atau dilakukan peran mana pun | F-FAM-004, F-FAM-005 | `PAGAR` |
| FR-FAM-009 | Visibilitas data milik sendiri | F-FAM-005 | `PAGAR`; `MSL-1` |
| FR-FAM-010 | Keterangan "siapa melihat apa" | F-FAM-005 | `PAGAR` |
| FR-FAM-011 | Pengaturan rumah tangga | F-FAM-001 | Fondasi |
| FR-FAM-012 | Keluar dengan aman | F-FAM-006 | `PAGAR` |
| FR-FAM-013 | Ekspor data saat keluar | F-FAM-006 | `PAGAR` |
| FR-FAM-014 | Menghapus data pribadi atas permintaan | F-FAM-006 | `PAGAR` |
| FR-FAM-015 | Objek terbuka sesudah seorang anggota keluar | F-FAM-006 | `PAGAR`; `MSL-1` |
| FR-FAM-016 | Tidak ada jalan mengeluarkan orang dewasa lain | F-FAM-004, F-FAM-006 | `PAGAR` |
| FR-FAM-017 | Kehilangan akses ke akun atau perangkat | F-FAM-004 | `PAGAR`; `MSL-1` |
| FR-FAM-018 | Pihak luar tanpa login | F-FAM-007 | `MSL-3`; `MSL-2` |
| FR-FAM-019 | Pemberitahuan akses pengelola server | F-FAM-008 | `PAGAR` |
| FR-FAM-020 | Fondasi: isolasi rumah tangga, akun terpisah dari anggota | F-FAM-001 | Fondasi tahap publik (PRN-03) |

**Usaha berulang seluruh `FAM`.** Sekali per orang (membuat atau bergabung, membaca pemberitahuan operator), sekali per pihak luar, dan nol sesudahnya. Tidak ada kebutuhan `FAM` yang meminta tindakan harian (PRN-02).

---

## 2. Business rule (penunjuk)

BR di bawah adalah **penunjuk** ke aturan kanonis. Isinya tidak didefinisikan ulang di sini; bila bunyi ringkas di kolom kedua berbeda dari sumbernya, sumbernya yang berlaku.

| BR | Pokok | Sumber kanonis |
|---|---|---|
| BR-FAM-001 | Pemberi persetujuan ditentukan satu fungsi aturan; pemohon selalu dikecualikan | 06 §3.2 (`resolve_approvers`), INV-A04, PG-02 |
| BR-FAM-002 | Lima konsep ditulis terpisah saat rumah tangga dibuat; aturan tidak membaca hubungan | 06 §3.1, PG-01, INV-A13; PRN-03, PRN-09 |
| BR-FAM-003 | Perilaku rumah tangga satu anggota dewasa, dan saat anggota dewasa kedua bergabung | 06 §3.5 (D-33) |
| BR-FAM-004 | Dari dua menjadi satu: keluar tidak menunggu objek apa pun; objek yang hanya dapat ditutup orang yang keluar keluar dari daftar terbuka | 06 §3.5, INV-A35 (D-42) |
| BR-FAM-005 | Tindakan yang mengenai orang dewasa lain hanya berlaku dengan persetujuannya; keluar tanpa persetujuan siapa pun. Di Tier 1 tidak ada perintah yang mengeluarkan, menonaktifkan, membekukan, atau menurunkan peran orang dewasa lain, juga tidak dengan persetujuannya; keanggotaan hanya aktif lewat konfirmasi pengundang | 06 PG-13 (dipersempit, revisi kesembilan), §3.6, INV-A48; default Q-04 |
| BR-FAM-006 | Tiga tingkat visibilitas; bawaan dan rentang per objek | 06 §9.1, §9.2 |
| BR-FAM-007 | Kombinasi visibilitas yang tidak diizinkan; perluasan tercatat dan diberitahukan, penyempitan senyap | 06 §9.3 V-1 sampai V-9, INV-A36, INV-A41, N-22 |
| BR-FAM-008 | Keterangan pelaku tanpa menjadi alat pantau | 06 §9.4 AT-1 sampai AT-7, PG-07, PG-09 |
| BR-FAM-009 | Larangan di dalam aplikasi tidak mengikat pengelola server | 06 §9.4 "Batas semua larangan"; keputusan 25 |
| BR-FAM-010 | Klaim antar-pemilik ditutup sistem saat anggota keluar (`exit_closure`), pada akhir masa tenggang (F07), dengan angka yang sama seperti bila ditulis saat akses berakhir | 07 §8.2 butir 2b, INV-28, CTH-38; 06 §3.6 F07; ADR-007 butir 8 |
| BR-FAM-011 | Inti, rincian, dan pemetaan identitas; penghapusan buku pribadi; catatan `CARE` di luar hanya-menambah | 07 §8.2, §8.3; 06 §2.5; ADR-007 butir 1 sampai 7 |
| BR-FAM-012 | Data pihak luar: nama panggilan dan hubungan saja | 07 §2.6; Q-28 |
| BR-FAM-013 | Menghapus butir beriwayat berarti mengarsipkannya; butir bersaldo tidak dapat diarsipkan | Keputusan 9; Q-58; D-26; 07 INV-21 |
| BR-FAM-014 | Setiap objek milik tepat satu rumah tangga; buku pribadi melekat pada anggota, bukan akun | 06 §14; ADR-003; D-45 |
| BR-FAM-015 | Jam tenang, tunda, dan waktu pelepasan milik orangnya; tanpa orakel | 06 INV-A33, INV-A42 |
| BR-FAM-016 | Alur keanggotaan F01 sampai F08 dan notifikasinya (N-24 kepada pengundang, N-25 kepada anggota baru, N-26 kepada yang tersisa). F06 keluar dengan autentikasi ulang: akses berakhir seketika; F07 langkah yang merusak pada akhir masa tenggang 7 hari, atau seketika bila yang keluar memilih hapus seketika; F08 pembatalan hanya oleh yang keluar, dengan autentikasi ulang, tanpa F04; F01 ditolak selama ada anggota dalam masa tenggang | 06 §3.6, INV-A48, §11.2; 07 §8.2 butir 2b |
| BR-FAM-017 | Zona waktu rumah tangga diisi saat dibuat. Selama rumah tangga belum pernah punya dua anggota dewasa aktif, satu-satunya anggota dewasa dapat menggantinya sendiri (`set_household_timezone`), berlaku mulai periode berikutnya; sesudah itu penggantiannya kesepakatan `rule` yang berlaku mulai periode berikutnya, juga bila kemudian tinggal satu | 06 §3.5 tabel baris "Zona waktu rumah tangga", §6.3 `rule`, K01, K05; 07 §6.1, INV-29 |
| BR-FAM-018 | Buku pribadi dibekukan sejak F06 dan ditutup pada F07 dalam transaksi `exit_closure`; dihapus seketika bila diminta atau pada akhir jangka retensi; sampai itu tidak terbaca siapa pun di rumah tangga dan hanya pemiliknya yang dapat mengekspornya | 07 §8.2 butir 2a, 2b; ADR-007 butir 10; 06 F06, F07 |
| BR-FAM-019 | Sesudah aksesnya berakhir, yang keluar tetap dapat mengekspor buku rumah tangga sampai saat aksesnya berakhir (entri dan tautan yang diterima server sebelum saat itu, saldo per saat itu); jangka dan jalannya dokumen 10 | 07 §8.2 butir 1a, §8.3 |
| BR-FAM-020 | Anggota yang bergabung sesudah ada yang keluar membaca inti uang buku rumah tangga seperti anggota lain; pelaku dan pihak yang menunjuk orang yang sudah keluar tampil "mantan anggota"; rincian tulisan mantan anggota (catatan, foto struk, lampiran, pesan diskusi) dan chat lama tidak terbaca olehnya; layar undangan menyebutnya kepada pengundang dan penerima sebelum F02 | 06 §3.5 "Saat anggota dewasa kedua bergabung", V-1, §16.2, INV-A47 (f); 07 §8.2 butir 3a, §8.3 |

---

## 3. Rumah tangga dan keanggotaan

### FR-FAM-001 — Membuat rumah tangga

| | |
|---|---|
| Fitur | F-FAM-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1`, `MSL-2` (fondasi: tanpa rumah tangga tidak ada buku bersama); `KS-1` |
| Aktor dan izin | Siapa pun yang punya akun dan belum menjadi anggota rumah tangga mana pun. Pembuat menjadi anggota dewasa berperan *pemilik rumah tangga*; membuat tidak memberi hak lebih (06 PG-01) |
| Aturan acuan | 06 §3.1, §3.5; 07 §6.1 (setelan awal hari mulai, baris `s0`); 06 P-11 |
| Business rule | BR-FAM-002, BR-FAM-003, BR-FAM-014 |
| Kedudukan | Kesetaraan penuh: keputusan pemilik nomor 4 (diputuskan sendiri, checkpoint S1). Pembuat tidak mendapat hak lebih: bagian default Q-04 yang diterima lewat "lanjut" (checkpoint S1). Satu zona waktu per rumah tangga diisi saat onboarding: Q-76, D-28 (diterima lewat "lanjut", checkpoint S2). Hari mulai bawaan tanggal 1: keputusan 15 (bacaan lead yang dikukuhkan kepada pemilik produk) dan D-49 (lewat "lanjut", Pagar). Isian minimal: **[Usulan]** `product-analyst` |

Membuat rumah tangga meminta satu isian wajib: nama rumah tangga. Zona waktu terisi dari perangkat dan dapat diganti saat itu juga. Hari mulai periode anggaran tampil dengan bawaan tanggal 1 dan dapat dipilih 1 sampai 28 (isinya milik F-FIN-009). Tidak ada pertanyaan tentang hubungan, struktur keluarga, jumlah anggota, penghasilan, atau data lain yang tidak dipakai aturan mana pun (PRN-02, PRN-03). Ambang permintaan **tidak** diisi saat membuat: ia kesepakatan rumah tangga yang baru dapat berlaku setelah anggota dewasa kedua menyetujuinya (F-REQ-008; 06 §3.5).

**Kriteria penerimaan**

- **AC-1** — *Given* seseorang punya akun dan bukan anggota rumah tangga mana pun, *When* ia membuat rumah tangga dengan nama dan zona waktu, *Then* rumah tangga tersimpan dengan pengenal baru, ia menjadi anggota dewasa berperan pemilik rumah tangga, dan setelan hari mulai periode bernilai 1 kecuali ia memilih nilai lain dari 1 sampai 28.
- **AC-2** — *Given* rumah tangga baru tersimpan, *When* model diperiksa, *Then* kelima konsep ditulis sebagai baris terpisah (peran, kumpulan izin preset, wewenang persetujuan untuk semua jenis permintaan tanpa batas jumlah, aturan rumah tangga dengan preset "dua dewasa setara" tanpa nilai ambang), tidak ada yang dihitung dari yang lain; konsep kelima, hubungan, punya tabelnya sendiri di skema tetapi tidak ada baris yang ditulis di pilot (06 §3.1 baris "Hubungan": "Tidak diisi"; FR-FAM-006).
- **AC-3** — *Given* layar membuat rumah tangga, *When* pembuat hanya mengisi nama, *Then* rumah tangga dapat dibuat; tidak ada isian wajib lain.
- **AC-4** — *Given* rumah tangga yang kelak punya dua anggota dewasa, *When* rangkaian perintah acak dijalankan dua kali dengan pembuat dan anggota kedua ditukar, *Then* hasil setiap perintah dan setiap tampilan identik dengan anggota tertukar; tidak ada perintah, izin, atau layar yang hanya tersedia bagi pembuat (06 INV-A13, PG-01).
- **AC-5** — *Given* seseorang sudah menjadi anggota satu rumah tangga, *When* ia membuka aplikasi, *Then* aplikasi pilot tidak menawarkan pembuatan rumah tangga kedua (Tier 3, F-FAM-009); larangan itu aturan produk pilot, bukan batas skema (FR-FAM-020).
- **AC-6** — *Given* perangkat tanpa sambungan, *When* seseorang mencoba membuat rumah tangga, *Then* aplikasi menyatakan bahwa pembuatan memerlukan sambungan dan tidak menampilkan rumah tangga sebagai sudah ada sebelum server menerimanya (06 S-1).

### FR-FAM-002 — Rumah tangga dengan satu anggota dewasa

| | |
|---|---|
| Fitur | F-FAM-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1` (mencatat dapat dimulai sebelum pasangan bergabung); `PAGAR` (tidak ada jalan menyetujui sendiri) |
| Aktor dan izin | Satu-satunya anggota dewasa |
| Aturan acuan | 06 §3.5 tabel "Selama hanya ada satu anggota dewasa", §3.2, INV-A04, INV-A10, INV-A11 |
| Business rule | BR-FAM-001, BR-FAM-003 |
| Kedudukan | D-33 (diterima lewat "lanjut", checkpoint S2; dibacakan sebagai baris Pagar). Rincian yang tidak disebut D-33 tetap **[Usulan]** `architect` sebagaimana ditandai di 06 §3.5 |

Selama rumah tangga baru berisi satu anggota dewasa, semua pencatatan, anggaran, laporan, dan koreksi berjalan penuh; yang tidak ada adalah segala hal yang memerlukan orang kedua. Aplikasi menyebut alasannya, bukan menyembunyikan fiturnya tanpa penjelasan.

**Kriteria penerimaan**

- **AC-1** — *Given* rumah tangga dengan satu anggota dewasa, *When* ia mencatat uang masuk, pengeluaran, pemindahan, anggaran, atau koreksi, *Then* semuanya tersimpan dan terhitung seperti di rumah tangga dua anggota (06 INV-A11).
- **AC-2** — *Given* rumah tangga dengan satu anggota dewasa, *When* ia membuka pengajuan permintaan anggaran atau jalan-jalan, *Then* aplikasi tidak menawarkan pengajuan dan menyatakan bahwa permintaan memerlukan anggota dewasa kedua; rencana dapat disimpan sebagai draf di perangkat yang tidak berakibat apa pun; perintah pengajuan yang tetap dikirim ditolak server dengan `no_approver` (06 §3.2).
- **AC-3** — *Given* rumah tangga dengan satu anggota dewasa, *When* ia mengusulkan ambang, hari mulai baru, atau kesepakatan pos rutin, *Then* usulan itu tersimpan sebagai menunggu dan tidak berlaku; tidak ada entri yang ditandai dengan dasar apa pun (06 §3.5).
- **AC-4** — *Given* rumah tangga dengan satu anggota dewasa, *When* ia mencatat apa pun, *Then* tidak ada catatan yang bertanda "belum dikonfirmasi", karena tidak ada tempat uang yang dipegang anggota dewasa lain.
- **AC-5** — *Given* rumah tangga dengan satu anggota dewasa, *When* jam dimajukan berapa pun, *Then* tidak ada usulan, draf, atau permintaan yang menjadi berlaku atau disetujui (06 INV-A10, PG-04).

### FR-FAM-003 — Membuat, membagikan, dan membatalkan undangan

| | |
|---|---|
| Fitur | F-FAM-002 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1`, `MSL-2` (kedua pasangan harus melihat buku yang sama) |
| Aktor dan izin | Anggota dewasa mana pun di rumah tangga itu. Pada preset "dua dewasa setara" undangan hanya tersedia selama rumah tangga berisi kurang dari dua anggota dewasa |
| Aturan acuan | 06 §3.5 (anggota kedua melihat seluruh buku, V-1, kecuali rincian mantan anggota), §3.6 F01 (guard: kurang dari dua anggota dewasa aktif, tidak ada undangan lain yang belum gugur, tidak ada anggota dalam `keluar_tenggang`; tanpa notifikasi), F03 (batal atau lewat masa berlaku; tanpa notifikasi); 07 §8.2 butir 3a; keamanan undangan: dokumen 10 **[Dok. 10]** |
| Business rule | BR-FAM-003, BR-FAM-005, BR-FAM-016, BR-FAM-020 |
| Kedudukan | **[Usulan]** `product-analyst`, dikukuhkan `architect` sebagai alur 06 §3.6, kecuali bagian yang dikutip. Tanpa undangan selama masa tenggang dan keterangan "mantan anggota" di layar undangan: **[Usulan]** `security-privacy` SP-48, SP-49, diterapkan 06 revisi kesepuluh. Masa berlaku undangan **[Asumsi]**; bentuk tautan atau kode dan entropinya ditetapkan dokumen 10 (S4) |

Undangan dibuat dari dalam aplikasi dan dibagikan lewat saluran apa pun yang dipilih pengundang (misalnya lembar bagikan sistem ke WhatsApp). Undangan sekali pakai, berlaku selama jangka terbatas (usulan: 7 hari **[Asumsi]**, angkanya ditetapkan dokumen 10), dan dapat dibatalkan pengundangnya selama belum diterima. Sebelum membuat undangan, pengundang diberi tahu satu hal yang tidak dapat ditarik kembali: anggota yang bergabung akan melihat seluruh catatan uang bersama sejak awal, termasuk yang dicatat sebelum ia bergabung (06 §3.5, V-1). Dana pribadi pengundang tetap tertutup.

**Kriteria penerimaan**

- **AC-1** — *Given* rumah tangga berisi satu anggota dewasa dan tidak ada undangan yang menunggu, *When* anggota itu membuat undangan, *Then* sebelum undangan dibuat ia diberi tahu bahwa yang bergabung akan melihat seluruh catatan uang bersama sejak awal dan tidak melihat dana pribadinya; sesudah ia melanjutkan, satu undangan sekali pakai dibuat dan dapat dibagikan lewat saluran pilihannya.
- **AC-2** — *Given* sudah ada satu undangan yang menunggu, *When* anggota yang sama atau anggota lain mencoba membuat undangan kedua, *Then* aplikasi menampilkan undangan yang sudah ada dan tidak membuat yang baru; pada preset pilot paling banyak satu undangan menunggu.
- **AC-3** — *Given* undangan yang belum diterima, *When* pengundang membatalkannya (06 F03), *Then* undangan itu tidak dapat lagi dipakai, upaya memakainya menghasilkan pesan "undangan tidak berlaku" tanpa menyebut nama rumah tangga, pembatalan tercatat sebagai Peristiwa, dan tidak ada notifikasi kepada siapa pun.
- **AC-4** — *Given* undangan yang masa berlakunya lewat, *When* seseorang mencoba memakainya, *Then* hasilnya sama dengan AC-3, dan pengundang dapat membuat undangan baru.
- **AC-5** — *Given* rumah tangga sudah berisi dua anggota dewasa, *When* salah satu membuka menu keanggotaan, *Then* tidak ada pilihan mengundang anggota dewasa ketiga (Tier 3, F-FAM-011); yang tersedia hanya informasi keanggotaan dan keluar (FR-FAM-012).
- **AC-6** — *Given* undangan dibuat, *When* undangan itu dibagikan, *Then* isi yang ikut terbagi tidak memuat jumlah uang, isi catatan, atau data pribadi siapa pun; paling banyak nama tampilan pengundang dan nama rumah tangga, sesuai rumusan dokumen 10.
- **AC-7** — *Given* B baru saja keluar (F06) dan masih dalam masa tenggang 7 hari **[Asumsi]**, *When* A mencoba membuat undangan, *Then* server menolak (guard F01) dan layar menyatakan bahwa undangan baru tersedia sesudah masa tenggang selesai, tanpa menyebut apakah B akan kembali; *When* masa tenggang lewat tanpa F08 (F07 berjalan), *Then* A dapat membuat undangan seperti AC-1. Dengan begitu kepulangan B lewat F08 tidak pernah menghasilkan tiga anggota dewasa (06 §3.6).
- **AC-8** — *Given* rumah tangga yang pernah punya anggota dewasa yang sudah keluar, *When* A akan membuat undangan, *Then* sebelum undangan dibuat A diberi tahu bahwa yang bergabung akan melihat inti uang seluruh buku sejak awal, tetapi tidak membaca chat lama dan tidak membaca nama, catatan, foto, atau lampiran tulisan anggota yang sudah keluar, yang tampil sebagai "mantan anggota" (06 §3.5; 07 §8.2 butir 3a; SP-49).

### FR-FAM-004 — Menerima undangan dan mengaktifkan keanggotaan

| | |
|---|---|
| Fitur | F-FAM-002 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1`, `MSL-2`; `PAGAR` |
| Aktor dan izin | Penerima undangan (orang yang punya akun atau membuatnya saat itu; autentikasi milik dokumen 10); pengundang untuk konfirmasi orangnya |
| Aturan acuan | 06 §3.1, §3.5, §3.6 F02, F04, F05, INV-A48 (b)(c); PG-13; N-24, N-25; syarat SP-41 (dokumen 17 kelompok (i)); pemberitahuan operator FR-FAM-019 |
| Business rule | BR-FAM-002, BR-FAM-003, BR-FAM-005, BR-FAM-009, BR-FAM-016 |
| Kedudukan | Setara sejak bergabung: keputusan pemilik nomor 4 (diputuskan sendiri). Konfirmasi orangnya oleh pengundang: usulan `product-analyst`, dikukuhkan `architect` sebagai F04 (06 §3.6; 06 P-23), diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-4 (b); dinilai dokumen 10. N-24 dan N-25: usulan `architect` (menjawab PA-FAM-4), diterima lewat "lanjut" sebagai bagian D-84. Anggota yang bergabung sesudah ada yang keluar (AC-10, AC-11): usulan `security-privacy` SP-49, diterapkan 06 revisi kesepuluh dan 07 §8.2 butir 3a, diterima lewat "lanjut" sebagai D-90. Menyepakati ambang terpisah dari menerima undangan: syarat SP-41, **[Usulan]** `security-privacy`, ditanyakan di checkpoint S4 (Q-93) |

Alurnya kanonis di 06 §3.6: F02 (penerima menerima; N-24 `perlu_tanggapan` kepada pengundang), F04 (hanya pengundang mengukuhkan; N-25 `kabar` kepada anggota baru), F05 (pengundang menyatakan bukan orangnya, penerima menarik diri, atau masa berlaku lewat; tanpa notifikasi). Menerima undangan terdiri atas tiga langkah, semuanya sebelum penerima melihat data rumah tangga: (1) penerima membaca apa yang ia masuki: nama rumah tangga, siapa yang mengundang, bahwa kedua anggota dewasa setara, apa yang akan dilihat pasangannya tentang dirinya (FR-FAM-010), dan pemberitahuan pengelola server (FR-FAM-019), lalu menerima; (2) pengundang mengonfirmasi bahwa nama yang menerima memang orang yang ia undang; (3) keanggotaan aktif. Langkah (2) ada karena tidak ada jalan mengeluarkan orang dewasa lain sesudah keanggotaan aktif (FR-FAM-016): undangan yang diteruskan ke orang yang salah harus berhenti sebelum orang itu menjadi anggota. Sebelum keanggotaan aktif, penerima tidak membaca data apa pun dan pengundang tidak memperoleh data apa pun tentang penerima selain nama tampilannya.

**Kriteria penerimaan**

- **AC-1** — *Given* undangan yang berlaku, *When* penerima membukanya, *Then* sebelum ada tombol terima ia membaca nama rumah tangga, nama tampilan pengundang, bahwa kedua anggota dewasa punya wewenang yang sama, apa yang akan terlihat oleh pasangannya tentang dirinya, dan pemberitahuan pengelola server; tidak ada data uang rumah tangga yang tampil.
- **AC-2** — *Given* penerima menekan terima, *When* pengundang belum mengonfirmasi, *Then* keanggotaan berstatus menunggu konfirmasi, penerima tidak dapat membaca atau menulis data rumah tangga, pengundang menerima N-24 berprioritas `perlu_tanggapan` dengan teks umum di layar kunci, dan pengundang melihat nama tampilan penerima beserta pilihan "ya, ini orangnya" dan "bukan" (06 F02).
- **AC-3** — *Given* keanggotaan menunggu konfirmasi, *When* pengundang memilih "bukan", *Then* keanggotaan tidak pernah aktif, undangan gugur, tidak ada data rumah tangga yang pernah terbaca penerima, tidak ada notifikasi kepada siapa pun, dan pengundang dapat membuat undangan baru (06 F05).
- **AC-4** — *Given* keanggotaan menunggu konfirmasi, *When* pengundang memilih "ya, ini orangnya", *Then* keanggotaan aktif dengan peran pemilik rumah tangga, kumpulan izin, dan wewenang persetujuan yang identik dengan pengundang (06 §3.1), peristiwa bergabung tercatat dengan kedua pelakunya, anggota baru menerima N-25 berprioritas `kabar`, dan chat rumah tangga mulai tampil bagi keduanya dalam keadaan kosong, tanpa kartu surut atas objek lama (06 F04, §16.2 "Peserta").
- **AC-5** — *Given* keanggotaan menunggu konfirmasi, *When* penerima menarik penerimaannya, *Then* ia keluar dari proses tanpa jejak data di rumah tangga itu selain Peristiwa undangan.
- **AC-6** — *Given* keanggotaan baru aktif, *When* ada usulan ambang atau kesepakatan lain yang menunggu, *Then* usulan itu disampaikan kepadanya sebagai hal yang menunggu tanggapan, terpisah dari langkah menerima undangan; menerima undangan tidak menyetujui ambang, hari mulai, atau kesepakatan apa pun, dan layar usulan ambang menyatakan bahwa ambang tidak dapat diubah sendirian (SP-41).
- **AC-7** — *Given* keanggotaan menunggu konfirmasi lebih lama dari masa berlaku undangan, *When* masa itu lewat, *Then* keanggotaan gugur seperti AC-3; tidak ada konfirmasi otomatis (06 PG-04).
- **AC-8** — *Given* rumah tangga sudah berisi dua anggota dewasa aktif, *When* undangan lama yang entah bagaimana masih ada dipakai, *Then* penerimaan ditolak dan tidak ada keanggotaan ketiga.
- **AC-9** — *Given* keanggotaan menunggu konfirmasi, *When* penerima sendiri, `system`, atau siapa pun selain pengundang mencoba membuatnya aktif, *Then* ditolak; keanggotaan hanya menjadi aktif lewat F04 oleh pengundang (06 INV-A48 (b)).
- **AC-10** — *Given* rumah tangga yang pernah punya anggota dewasa M yang sudah keluar (F07 sudah berjalan), *When* C membuka undangan dari A, *Then* sebelum tombol terima C membaca, selain isi AC-1, bahwa ia akan melihat inti uang seluruh buku rumah tangga sejak awal, bahwa catatan yang dibuat anggota yang sudah keluar tampil sebagai "mantan anggota" tanpa nama, catatan, foto, atau lampirannya, dan bahwa chat lama tidak terbaca olehnya (06 §3.5; SP-49).
- **AC-11** — *Given* AC-10 dan A mengukuhkan C (F04), *When* C membuka buku rumah tangga, entri yang dicatat M, dan chat, *Then* jumlah, tanggal, kategori, tempat uang, dan saldo setiap entri identik dengan yang dilihat A (07 INV-20); pelaku yang menunjuk M tampil "mantan anggota"; catatan teks, foto struk, lampiran, dan pesan diskusi yang ditulis M tidak terbaca C lewat layar, API, maupun ekspor, walau tetap terbaca A; C hanya melihat percakapan chat baru (FR-COM-001/AC-7; 07 §8.2 butir 3a).

### FR-FAM-005 — Yang berlaku saat anggota dewasa kedua bergabung

| | |
|---|---|
| Fitur | F-FAM-002 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1` (angka yang sama bagi keduanya); `KS-1` |
| Aktor dan izin | Kedua anggota dewasa |
| Aturan acuan | 06 §3.5 "Saat anggota dewasa kedua bergabung", V-1, §8.2 (keberatan) |
| Business rule | BR-FAM-003, BR-FAM-007 |
| Kedudukan | D-33 (diterima lewat "lanjut", checkpoint S2) |

**Kriteria penerimaan**

- **AC-1** — *Given* anggota pertama mencatat entri dana bersama sebelum anggota kedua bergabung, *When* anggota kedua aktif, *Then* ia melihat seluruh buku rumah tangga sejak awal, angka utama dan rincian custody yang identik dengan yang dilihat anggota pertama (06 §3.5, V-1). Satu pengecualian atas rincian, bukan atas angka: rincian tulisan anggota dewasa yang sudah keluar sebelum ia bergabung tidak terbaca olehnya (FR-FAM-004/AC-11).
- **AC-2** — *Given* entri yang dicatat sebelum anggota kedua bergabung, *When* anggota kedua aktif, *Then* tidak ada entri lama yang menjadi bertanda di atas ambang atau berstatus "belum dikonfirmasi"; tidak ada yang berlaku surut (06 §3.5; D-33).
- **AC-3** — *Given* anggota kedua aktif, *When* ia tidak sepakat dengan entri yang dicatat sebelum ia bergabung, *Then* ia dapat mengajukan keberatan atas entri itu tanpa batas waktu (06 §8.2).
- **AC-4** — *Given* anggota pertama punya tempat uang pribadi dan buku pribadi, *When* anggota kedua aktif, *Then* anggota kedua tidak melihat saldo, isi, atau keberadaan tempat uang pribadi yang tidak dibagikan (06 V-4; 17 D-31).
- **AC-5** — *Given* anggota pertama menyimpan draf permintaan di perangkatnya, *When* anggota kedua aktif, *Then* draf itu tetap di perangkat, tidak dikirim otomatis, dan sejak saat itu dapat diajukan dengan anggota kedua sebagai pemberi persetujuannya (06 §3.2).

### FR-FAM-006 — Profil anggota minimal

| | |
|---|---|
| Fitur | F-FAM-003 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | Fondasi: nama dipakai pada keterangan pelaku (`KS-3` "siapa membayar apa"; keputusan 4) |
| Aktor dan izin | Hanya pemilik profil yang mengubah profilnya sendiri |
| Aturan acuan | 06 §3.1 (baris hubungan), PG-01; PRN-03; 07 §8.2 butir 3 ("mantan anggota") |
| Business rule | BR-FAM-002, BR-FAM-011 |
| Kedudukan | **[Usulan]** `product-analyst`. Baris hubungan tidak diisi: 06 §3.1 (**[Usulan]** `architect`, revisi kesembilan, menjawab PA-FAM-1). Layar kunci tanpa nama: S3-8 (vii) bentuk (a), diterima lewat "lanjut" di checkpoint S3 (12 November 2026); alternatif (b), nama di layar kunci, tidak dipilih (AC-5) |

Profil memuat nama tampilan (wajib) dan foto (opsional). Tidak ada tanggal lahir, nomor telepon yang terlihat anggota lain, alamat, hubungan, atau data lain yang tidak dipakai aturan mana pun. Konsep *hubungan* tetap terpisah di model: tabelnya ada di skema karena kelima konsep harus ada sejak Tier 1 (PRN-03), tetapi pilot tidak menanyakannya dan **tidak menulis baris apa pun**; Tier 3 cukup mengisinya (06 §3.1 baris "Hubungan", PG-01).

**Nama tampilan dan layar kunci** (06 §11.1 "Isi layar kunci", P-32; checkpoint S3-8). Yang diterima pemilik produk atas saran lead adalah pilihan "Umum di layar kunci", yang teksnya berbunyi "ada kabar dari istrimu". Teks pastinya bentuk (a), usulan `security-privacy` SP-46, **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026), S3-8 (vii): layar kunci tidak memuat nama siapa pun, jadi berkas ini tidak menambah apa pun. **Historis:** alternatif (b) (setiap `kabar` dan N-23 dari pasangan berbunyi "ada kabar dari {nama}", dengan `{nama}` nama tampilan yang dipilih pengirim sendiri; CA-07) tidak dipilih. Tetap berlaku: tidak ada isian nama panggilan untuk orang lain.

**Kriteria penerimaan**

- **AC-1** — *Given* anggota dewasa, *When* ia mengubah nama tampilan atau fotonya, *Then* perubahan berlaku pada semua keterangan pelaku yang sudah ada dan yang akan datang, karena nama dibaca dari pemetaan identitas, bukan disalin ke catatan (07 §8.2).
- **AC-2** — *Given* anggota dewasa, *When* ia mencoba mengubah profil anggota lain, *Then* perintah itu tidak tersedia dan ditolak server.
- **AC-3** — *Given* layar profil dan layar bergabung, *When* diperiksa, *Then* tidak ada isian hubungan, tanggal lahir, jenis kelamin, atau penghasilan; *When* basis data rumah tangga pilot diperiksa, *Then* tabel hubungan ada dan tidak berisi baris (06 §3.1).
- **AC-4** — *Given* anggota yang sudah keluar dan pemetaan identitasnya sudah diputus, *When* anggota yang tersisa membuka catatan lama, *Then* pelakunya tampil sebagai "mantan anggota", bukan nama lamanya (07 §8.2 butir 3). *Given* anggota yang bergabung sesudah orang itu keluar, *When* ia membuka catatan yang sama, *Then* pelakunya tampil "mantan anggota" apa pun keadaan pemetaannya (07 §8.2 butir 3a; FR-FAM-004/AC-11).
- **AC-5** — *Given* A mengganti nama tampilannya, *When* sesudahnya B menerima `kabar` atau N-23 dari A, *Then* layar kunci B tidak memuat nama A sama sekali (bentuk (a), S3-8 (vii), diterima lewat "lanjut"; FR-NTF-005/AC-5), dan B tidak punya isian untuk memberi A nama lain. (Cabang semula untuk alternatif (b) dihapus karena (b) tidak dipilih.)

---

## 4. Wewenang, peran, dan visibilitas

### FR-FAM-007 — Wewenang setara tanpa layar pengaturan izin

| | |
|---|---|
| Fitur | F-FAM-004 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` (permintaan dua arah); `MSL-4` (permintaan dan persetujuan memerlukan pemberi persetujuan yang jelas) |
| Aktor dan izin | Kedua anggota dewasa; wewenang persetujuan identik (06 §3.1) |
| Aturan acuan | 06 §3.1, §3.2, §3.3, PG-01, PG-02; ADR-010 |
| Business rule | BR-FAM-001, BR-FAM-002 |
| Kedudukan | Kesetaraan: keputusan pemilik nomor 4 (diputuskan sendiri). Tanpa layar pengaturan izin: Q-30, usulan `product-analyst`, **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026), S3-1 (a) |

Tidak ada layar untuk mengatur peran, izin, tanggung jawab, atau wewenang persetujuan di pilot. Pemberi persetujuan setiap permintaan ditentukan aturan saat permintaan diajukan, tidak dipilih pemohon, dan selalu anggota dewasa yang bukan pemohon. Kebutuhan ini juga menjadi dasar `REQ` dan `FIN` untuk menanyakan "siapa yang berwenang" (keduanya memanggil fungsi yang sama, 06 §3.2).

**Kriteria penerimaan**

- **AC-1** — *Given* rumah tangga dua anggota dewasa, *When* anggota dewasa A mengajukan permintaan, *Then* pemberi persetujuannya B; *When* B mengajukan permintaan, *Then* pemberi persetujuannya A. Pemohon tidak pernah dapat memilih atau mengubah pemberi persetujuan (simetris; 06 §3.2).
- **AC-2** — *Given* permintaan yang diajukan A, *When* A mencoba menyetujui permintaannya sendiri lewat jalur apa pun, *Then* perintah ditolak (06 INV-A04).
- **AC-3** — *Given* aplikasi pilot, *When* menu mana pun dibuka, *Then* tidak ada layar pengaturan izin atau wewenang; satu-satunya pengaturan rumah tangga yang menyangkut aturan adalah kesepakatan rumah tangga yang memerlukan kedua anggota (F-REQ-008).
- **AC-4** — *Given* basis kode alur, *When* pemeriksaan statis dijalankan, *Then* tidak ada aturan, guard, bawaan, atau batas yang membaca hubungan, jenis kelamin, penghasilan, pemegang uang, atau pembuat rumah tangga (06 PG-01).
- **AC-5** — *Given* aturan rumah tangga diubah sesudah sebuah permintaan diajukan, *When* permintaan itu ditanggapi, *Then* pemberi persetujuannya tetap yang dibekukan saat diajukan (06 §3.2).

### FR-FAM-008 — Yang tidak pernah dapat dilihat atau dilakukan peran mana pun

| | |
|---|---|
| Fitur | F-FAM-004, F-FAM-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; brief §8 "what an admin can never see" |
| Aktor dan izin | Setiap peran, termasuk pemilik rumah tangga dan admin (Tier 3) |
| Aturan acuan | 06 V-4, V-5, V-9, PG-02, PG-03, PG-07, PG-13, AT-3 sampai AT-6, INV-A33, INV-A42, INV-A45, INV-A47, INV-A48; 07 INV-25 |
| Business rule | BR-FAM-005, BR-FAM-007, BR-FAM-008, BR-FAM-009, BR-FAM-015 |
| Kedudukan | Larangan dikutip dari 06 (D-23, D-31, D-34, D-40 diterima lewat "lanjut", checkpoint S2; V-9 **[Usulan]**). Batasnya terhadap pengelola server: keputusan 25 (diputuskan sendiri, dengan syarat) |

Daftar tertutup. Peran apa pun, sekarang atau kelak, **tidak pernah** dapat:

| # | Yang tidak dapat | Sumber |
|---|---|---|
| 1 | Membaca saldo, isi, atau keberadaan tempat uang dan buku pribadi orang dewasa lain yang tidak dibagikan pemiliknya | V-4, V-5; 07 INV-25 |
| 2 | Membaca status suasana hati, laporan sakit, atau alasannya yang tidak dibagikan pemiliknya | V-4, V-7; brief §5K |
| 3 | Membaca jumlah seluruh isi dompet campuran orang lain dan pembagian selisihnya | V-9, INV-A41 |
| 4 | Membaca jam tenang, tunda, dan waktu notifikasi dilepas atau dibuka pada orang lain; membaca tanda dibaca atau kehadiran di luar layar chat, atau waktu baca per pesan di mana pun | INV-A33, INV-A42, INV-A45; AT-5 (b)(c) |
| 5 | Membaca umpan kegiatan per orang, hitungan, peringkat, atau tren per orang | AT-3, AT-4, PG-09 |
| 6 | Menyetujui, menolak, mengonfirmasi, menerima, atau menandai selesai atas nama orang dewasa lain | PG-02 |
| 7 | Menyembunyikan apa pun yang menyangkut dana bersama dari orang dewasa yang berhak melihatnya | V-1, V-5, PG-03 |
| 8 | Mengubah izin, wewenang, profil, jam tenang, setelan berbagi, atau setelan AI orang dewasa lain; mengeluarkan, menonaktifkan, membekukan, atau menurunkan perannya, juga dengan persetujuannya | PG-13 (dipersempit), INV-A48 (a); FR-FAM-016 |
| 9 | Menulis ke buku pribadi orang lain | 07 INV-25; D-31 |
| 10 | Menghapus pesan chat orang lain, atau menghapus kartu dan baris peristiwa di chat | 06 §16.7, INV-A47 (a)(b) |

**Chat dan daftar ini.** Keputusan pemilik produk 12 November 2026 tentang chat (`_progress.md`, "Pernyataan pemilik produk tentang chat") menaikkan chat rumah tangga ke Tier 1 ("Chat naik Tier 1", diterima atas saran lead), dengan tanda terkirim dan dibaca pada pesan dan kartu di chat ("Ikut tanda dibaca", diterima atas saran lead) serta terakhir aktif, sedang online, dan sedang mengetik (diputuskan pemilik produk sendiri, berlawanan dengan saran lead). Tanda dan kehadiran itu bukan pelanggaran daftar ini: 06 AT-5 dan PG-07 sudah dirumuskan ulang sehingga keduanya ada **hanya di layar chat**, sama bagi kedua orang, tanpa waktu baca, tanpa riwayat kehadiran, dan tidak dibaca aturan, pengingat, ringkasan, laporan, ekspor, atau `AI` mana pun (06 AT-5, PG-07, §16.5, §16.6, INV-A45). Butir 4 tetap berlaku utuh di luar chat. Kebutuhan chat: [COM.md](COM.md).

**Kriteria penerimaan**

- **AC-1** — *Given* uji akses per peran atas setiap butir 1 sampai 5, *When* dijalankan sebagai setiap peran yang ada di model, *Then* hasilnya identik dan tidak bergantung pada peran: data itu tidak ditemukan, tidak dapat diturunkan, dan tidak muncul di layar, Peristiwa, notifikasi, laporan, ekspor, atau tanggapan API mana pun (06 PG-03).
- **AC-2** — *Given* setiap perintah di butir 6, 8, 9, dan 10, *When* dikirim oleh anggota yang bukan orang yang pernyataannya diminta atau bukan penulis pesannya, *Then* ditolak server, apa pun perannya (06 PG-02, PG-13, INV-A47, INV-A48).
- **AC-3** — *Given* butir 7, *When* seorang anggota mencoba membuat objek dana bersama berstatus `private` atau mempersempitnya di bawah `spouses`, *Then* perintah ditolak (06 V-1).
- **AC-4** — *Given* layar keterangan "siapa melihat apa" (FR-FAM-010), *When* dibaca, *Then* daftar di atas tersampaikan dengan bahasa sehari-hari, bersama batasnya: larangan ini mengikat peran di aplikasi, tidak akses pengelola server (FR-FAM-019; 06 §9.4).
- **AC-5** — *Given* peran Tier 3 kelak ditambahkan (admin, tamu, wali), *When* uji AC-1 dan AC-2 dijalankan, *Then* uji yang sama lulus tanpa perubahan; peran baru tidak boleh menambah pengecualian atas daftar ini kecuali wali atas tanggungan, lewat ADR (06 §3.4, PG-02).

### FR-FAM-009 — Visibilitas data milik sendiri

| | |
|---|---|
| Fitur | F-FAM-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; `MSL-1` (dana bersama tidak dapat disembunyikan) |
| Aktor dan izin | Pemilik data. Objek dana bersama: tidak dapat dipersempit siapa pun (V-1) |
| Aturan acuan | 06 §9.1, §9.2, V-1 sampai V-9, INV-A36, N-22; PRN-12, PRN-13 |
| Business rule | BR-FAM-006, BR-FAM-007 |
| Kedudukan | Tingkat visibilitas per hal: keputusan pemilik nomor 4 (diputuskan sendiri). Bawaan dan rentang per objek: Q-51, D-25 (lewat "lanjut"). Penyempitan senyap: D-40 (lewat "lanjut", dibacakan sebagai Pagar). Rumusan akhir: dokumen 10 **[Dok. 10]** |

Tingkat visibilitas tidak ditanyakan pada tiap pencatatan; setiap jenis objek punya bawaan (06 §9.2). Yang dapat diubah pemilik data di Tier 1 hanyalah membagikan tempat uang pribadinya kepada pasangan (`private` → `spouses`) dan mencabutnya lagi. Berbagi status `CARE` diatur di [tier-2.md](tier-2.md) dengan aturan yang sama.

**Kriteria penerimaan**

- **AC-1** — *Given* anggota dewasa mencatat apa pun, *When* layar entri dibuka, *Then* tidak ada pertanyaan visibilitas; catatan mengikuti bawaan jenis objeknya (06 §9.2).
- **AC-2** — *Given* tempat uang pribadi M, *When* M membagikannya kepada pasangan, *Then* pasangan dapat membaca tempat uang dan saldonya sejak saat itu, perubahan tercatat sebagai Peristiwa, dan N-22 dikirim kepada setiap anggota selain M yang datanya menjadi lebih terlihat (06 V-6, INV-A36).
- **AC-3** — *Given* tempat uang pribadi M yang sedang dibagikan, *When* M mencabutnya, *Then* berlaku seketika tanpa persetujuan siapa pun, pasangan tidak lagi menemukan tempat uang itu maupun Peristiwanya, dan tidak ada notifikasi apa pun kepada siapa pun (06 V-6, V-7, INV-A36).
- **AC-4** — *Given* pasangan M, *When* ia mencoba membagikan atau mencabut berbagi data milik M, *Then* perintah ditolak (06 V-4).
- **AC-5** — *Given* objek apa pun yang menyangkut dana bersama, *When* pemilik, pencatat, atau peran mana pun mencoba mempersempit visibilitasnya di bawah `spouses`, *Then* ditolak (06 V-1, V-3).
- **AC-6** — *Given* kesimetrisan, *When* AC-2 sampai AC-5 dijalankan dengan M dan pasangannya ditukar, *Then* hasilnya identik.

### FR-FAM-010 — Keterangan "siapa melihat apa"

| | |
|---|---|
| Fitur | F-FAM-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; brief §8 "transparency about who can see what" |
| Aktor dan izin | Setiap anggota, atas dirinya sendiri |
| Aturan acuan | 07 §5.3 "Kebocoran yang tersisa" (kalimat pembuka: layar mengatakannya sebelum menyimpan); PRN-12 cara diuji; 06 §9.4, AT-5 (a)(d), §16.5, §16.6, §13.5 "Tanda dibaca" |
| Business rule | BR-FAM-007, BR-FAM-009 |
| Kedudukan | **[Usulan]** `product-analyst`. Isi teks: dokumen 08 dan 10 (S4). Versi dinamis ("lihat sebagai pasangan") Tier 3 (dokumen 11 IDEA-20) |

Dua bentuk, keduanya statis di pilot. (a) **Sebelum menyimpan**: setiap layar yang akan membuat sesuatu dari sisi pribadi terlihat pasangan (penyisihan, kontribusi dari dana pribadi, talangan "minta diganti", bagian pribadi penghasilan) menyatakannya sebelum disimpan. (b) **Kapan saja**: satu penjelasan "yang dilihat pasanganmu tentang kamu, dan yang tidak", yang juga tampil saat bergabung (FR-FAM-004 AC-1).

**Kriteria penerimaan**

- **AC-1** — *Given* layar yang akan mencatat entri dua buku yang menyentuh dana pribadi pencatat, *When* pencatat akan menyimpan, *Then* layar menyatakan apa yang akan terlihat oleh pasangan (jumlah dan pelakunya di sisi rumah tangga) dan apa yang tetap tertutup (saldo dan pemakaian dana pribadi), sebelum tombol simpan (07 §5.3).
- **AC-2** — *Given* anggota mana pun, *When* ia membuka penjelasan "siapa melihat apa", *Then* isinya mencakup daftar persentuhan dana pribadi yang terlihat menurut 07 §5.3, daftar FR-FAM-008, dan batas terhadap pengelola server (FR-FAM-019), dan isinya sama bagi kedua anggota dewasa.
- **AC-3** — *Given* tinjauan dokumen 08, *When* daftar di AC-2 dibandingkan dengan 07 §5.3 dan 06 §9.2, *Then* tidak ada persentuhan yang terlihat pasangan tetapi tidak disebut (PRN-12 cara diuji).
- **AC-4** — *Given* chat rumah tangga Tier 1 (keputusan 12 November 2026), *When* penjelasan "siapa melihat apa" dibaca, *Then* ia juga menyebut, sebelum chat pertama kali dipakai: (a) pasangan melihat tanda terkirim dan dibaca pada pesan dan kartu di chat, tanpa waktu baca; (b) pasangan melihat terakhir aktif, sedang online, dan sedang mengetik, hanya di layar chat; (c) kartu permintaan yang sudah dibaca memperlihatkan bahwa kamu sudah melihatnya walau kamu menunda, tetapi tidak memperlihatkan apakah kamu menunda atau setelan jam tenangmu; (d) bila `CARE` (Tier 2) dibangun, pemilik status kepedulian dapat melihat bahwa pasangannya sudah membaca kartu statusnya (06 AT-5 (a)(d), §16.5, §16.6, §13.5 "Tanda dibaca"). Rumusan akhirnya dokumen 08 dan 10.

### FR-FAM-011 — Pengaturan rumah tangga

| | |
|---|---|
| Fitur | F-FAM-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | Fondasi |
| Aktor dan izin | Anggota dewasa mana pun untuk nama. Zona waktu: diisi pembuat saat rumah tangga dibuat; selama rumah tangga belum pernah punya dua anggota dewasa aktif, satu-satunya anggota dewasa dapat menggantinya sendiri; sesudahnya diganti lewat usulan kesepakatan `rule` oleh anggota dewasa mana pun, berlaku bila anggota dewasa lain menyetujuinya |
| Aturan acuan | 06 PG-08 (tidak ada perubahan diam-diam); P-11 (Q-76; revisi kesembilan dan kesepuluh); §3.5 tabel baris "Zona waktu rumah tangga" (`set_household_timezone`); §6.3 `rule`, K01 (guard zona waktu), K05; 07 §6.1, INV-29; hari mulai periode: 06 §6.3 dan 07 §6.1 (kesepakatan `rule`, F-REQ-008) |
| Business rule | BR-FAM-017 |
| Kedudukan | **[Usulan]** `product-analyst` untuk nama. Zona waktu lewat `rule`, berlaku mulai periode berikutnya: **[Usulan]** `architect` (revisi kesembilan, menjawab PA-FAM-2). Penggantian sendiri sebelum pernah ada dua anggota dewasa: **[Usulan]** `architect` (revisi kesepuluh, menjawab PA-FAM-8, mengikuti bawaan di sini) |

Zona waktu rumah tangga menggeser batas periode anggaran, kedaluwarsa, jam tenang, jadwal ringkasan, dan pengingat bertanggal kedua orang, jadi ia diperlakukan seperti hari mulai periode: penggantiannya kesepakatan `rule` (satu versi memuat seluruh isi aturan rumah tangga) dan berlaku mulai periode berikutnya tanpa menggeser periode yang sudah dimulai (06 §6.3 K05; 07 INV-29). Pengecualiannya rumah tangga yang belum pernah punya dua anggota dewasa aktif: di sana belum ada jatah atau kesepakatan yang dapat berlaku, jadi satu-satunya anggota dewasa dapat membetulkan zona waktunya sendiri, juga berlaku mulai periode berikutnya (06 §3.5). Hari mulai periode tidak ikut dilonggarkan. Usaha berulang: nol; zona waktu hampir tidak pernah berganti.

**Kriteria penerimaan**

- **AC-1** — *Given* anggota dewasa mana pun, *When* ia mengganti nama rumah tangga, *Then* perubahan tercatat sebagai Peristiwa beserta pelakunya dan terlihat oleh kedua anggota (06 PG-08); tidak memerlukan persetujuan dan tidak membuat notifikasi.
- **AC-2** — *Given* rumah tangga yang pernah punya dua anggota dewasa aktif, *When* anggota membuka pengaturan, *Then* zona waktu tampil dan hanya dapat diubah lewat usulan kesepakatan `rule`, sama dengan hari mulai periode (F-FIN-009, F-REQ-008); tidak ada pengubahan langsung dari layar pengaturan, dan perintah `set_household_timezone` ditolak, juga bila kini tinggal satu anggota dewasa.
- **AC-3** — *Given* rumah tangga dua anggota dewasa dengan zona waktu Asia/Jakarta dan hari mulai 1, *When* A mengusulkan Asia/Makassar pada 10 November 2026 dan B menyetujuinya hari itu, *Then* periode 1 sampai 30 November tetap menurut Asia/Jakarta dan zona waktu baru berlaku sejak periode yang dimulai 1 Desember 2026; sampai itu, kedaluwarsa, jam tenang, dan ringkasan tetap memakai Asia/Jakarta (06 §6.3 K05).
- **AC-4** — *Given* usulan zona waktu yang tidak ada di daftar zona waktu yang dikenal server, *When* dikirim, *Then* ditolak (06 K01).
- **AC-5** — *Given* rumah tangga yang belum pernah punya dua anggota dewasa aktif, dibuat dengan Asia/Jakarta dan hari mulai 1, *When* satu-satunya anggota dewasa mengganti zona waktu ke Asia/Makassar pada 10 November 2026, *Then* perubahan tersimpan langsung sebagai Peristiwa berpelaku, tanpa usulan dan tanpa persetujuan; periode 1 sampai 30 November tetap menurut Asia/Jakarta, dan Asia/Makassar berlaku sejak periode yang dimulai 1 Desember 2026 (06 §3.5 baris "Zona waktu rumah tangga"; 07 INV-29). Hari mulai periode tetap hanya lewat `rule`, yang menunggu anggota dewasa kedua.

---

## 5. Keluar, ekspor, dan penghapusan

### FR-FAM-012 — Keluar dengan aman

| | |
|---|---|
| Fitur | F-FAM-006 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; brief §8 "safe exit from a household", "a couple separates", "a member leaves the family" |
| Aktor dan izin | Setiap anggota dewasa, atas dirinya sendiri; tidak memerlukan persetujuan siapa pun |
| Aturan acuan | 06 PG-13, §3.2, §3.5 "Dari dua menjadi satu", "Kapan tabel di bawah berlaku", dan "Selama masa tenggang", §3.6 F06 sampai F08 dan "Masa tenggang keluar", INV-A35, INV-A48 (a)(b)(d)(e), N-26, §16.2 "Peserta"; 07 §8.2 butir 1, 1a, 2, 2a, 2b, 5, 8 dan tabel `exit_closure`; CTH-38; ADR-007 butir 8 dan 10 |
| Business rule | BR-FAM-004, BR-FAM-005, BR-FAM-010, BR-FAM-011, BR-FAM-016, BR-FAM-018, BR-FAM-019 |
| Kedudukan | Keluar tanpa persetujuan: default Q-04 (lewat "lanjut", checkpoint S1). Keluar tidak menunggu apa pun dan penutupan klaim oleh sistem: D-42 (lewat "lanjut", Pagar) di atas keputusan 24 (S2-8 "a", sementara sampai advokat menjawab). **Autentikasi ulang, masa tenggang 7 hari [Asumsi], pembatalan hanya oleh yang keluar, dan hapus seketika sebagai pilihan:** usulan `security-privacy` SP-48, diterapkan 06 revisi kesepuluh (F06 sampai F08) dan 07 §8.2 butir 2b; **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026), S3-4 (c); harganya W-24. Akses dan keanggotaan uang dipisah selama masa tenggang (AC-4, AC-14): usulan `architect`, diterima lewat "lanjut" bersama S3-4 (c), 06 §3.5 "Selama masa tenggang", mengikuti 07 §8.2 butir 2b (menjawab PA-FAM-9). Perlakuan catatan `CARE` saat keluar: **[Usulan]** `product-analyst`. Buku pribadi ditutup lalu dihapus: 07 §8.2 butir 2a dan ADR-007 butir 10 (menjawab PA-FAM-7). N-26: **[Usulan]** `architect`, apakah dan kapan dikirim dinilai dokumen 10 (06 P-29) |

Keluar adalah satu perintah orang itu sendiri (F06), dengan **autentikasi ulang memakai kredensial akun**, bukan sesi yang sedang terbuka, supaya orang yang hanya memegang ponsel yang terbuka tidak dapat menjalankannya. Sebelum mengukuhkan, ia membaca akibatnya, ditawari ekspor (FR-FAM-013), dan memilih apakah ia juga meminta penghapusan data pribadinya (FR-FAM-014) dan, bila ya, apakah seketika. **Aksesnya berakhir seketika** pada F06. Langkah yang merusak (`exit_closure`, penutupan buku pribadi, penghapusan catatan `CARE`, penutupan chat baginya) berjalan pada F07, **7 hari [Asumsi]** kemudian, atau dalam transaksi F06 bila ia memilih hapus seketika. Selama masa tenggang hanya ia sendiri yang dapat membatalkan (F08, dengan autentikasi ulang, tanpa F04), dan tidak ada undangan baru (FR-FAM-003/AC-7). Sesudah F07 keluar tidak dapat dibatalkan; bergabung lagi berarti menjadi anggota baru (07 §8.2 tabel `exit_closure`, baris "Tidak dapat dibalik"). Pembagian harta ketika pasangan berpisah adalah keputusan keluarga atau hukum, bukan aturan aplikasi (07 §8.2 butir 5).

**Kriteria penerimaan**

- **AC-1** — *Given* anggota dewasa A (simetris: berlaku sama bagi B), *When* A memilih keluar, *Then* sebelum mengukuhkan ia membaca: (a) aksesnya berakhir seketika; (b) klaim antar-pemilik yang masih terbuka beserta jumlahnya akan ditutup sebagai "ditutup karena anggota keluar", tanpa menyatakan lunas, pada akhir masa tenggang 7 hari; (c) tempat uang dana bersama yang ia pegang tetap tercatat dan diurus anggota yang tersisa; (d) catatan dana bersama tetap ada; (e) selama 7 hari hanya ia sendiri yang dapat membatalkan keluarnya, dan sesudahnya tidak dapat dibatalkan; (f) bila ada catatan di perangkatnya yang belum terkirim, berapa jumlahnya dan bahwa catatan itu tidak akan masuk (07 §8.2 butir 2b); lalu ditawari ekspor.
- **AC-2** — *Given* A punya permintaan, usulan, konfirmasi custody, keberatan, atau klaim yang masih terbuka dalam state apa pun, *When* A mengukuhkan keluar, *Then* perintah keluar diterima; tidak ada state objek yang menolak atau menahannya (06 INV-A35).
- **AC-3** — *Given* A mengukuhkan keluar tanpa memilih hapus seketika, *When* server menerima F06, *Then* belum ada langkah yang merusak: tidak ada `exit_closure`, buku pribadi A dibekukan (tidak menerima baris dari siapa pun, termasuk `claim_waiver` oleh anggota lain atas klaim dana bersama kepada A), catatan `CARE` A masih ada tetapi tidak terbaca siapa pun, dan klaim terbuka tetap tampil di buku rumah tangga. *When* masa tenggang 7 hari lewat menurut waktu server tanpa F08, *Then* F07 berjalan tepat satu kali, dalam satu transaksi: buku pribadi A ditutup, tepat satu entri `exit_closure` ditulis bila ada klaim terbuka dengan tanggal kejadian hari akhir masa tenggang, catatan `CARE` A dihapus, chat ditutup baginya; *uang yang ada* rumah tangga tidak berubah; pengulangan tidak menulis apa pun lagi (06 F07, INV-A48 (d); 07 §8.2 butir 2a, 2b; INV-10, INV-28). Penutupan terjadi pada setiap keluar, juga bila A tidak meminta penghapusan. Angkanya: FR-FAM-012/AC-10.
- **AC-4** — *Given* A sudah keluar (F06), *When* anggota yang tersisa membuka aplikasi, selama maupun sesudah masa tenggang, *Then* ia tetap dapat mencatat, membaca seluruh buku rumah tangga, dan mengurus tempat uang yang dulu dipegang A; tidak ada pencatatan yang terhalang (06 INV-A11). Selama masa tenggang A tetap terhitung anggota dewasa untuk syarat tanda dan tetap pemegang tempat uangnya; baris yang menambah tempat uang A bertanda belum dikonfirmasi dan tidak dapat dijawab A selama itu; hitungan bukan pemegang atas tempat uang A ditolak (07 §8.2 butir 2b).
- **AC-5** — *Given* A sudah keluar (F06), *When* A membuka aplikasi dengan akunnya, *Then* A tidak lagi membaca atau menulis data rumah tangga itu dan tidak menerima notifikasinya; yang masih dapat ia lakukan hanyalah membatalkan selama masa tenggang (F08, AC-11) dan satu jalan ekspor berautentikasi ulang: buku rumah tangga per saat aksesnya berakhir dan buku pribadinya sendiri (FR-FAM-013/AC-6, AC-8; 06 §3.5, INV-A48 (c)); jangka dan jalannya **[Dok. 10]**. Perintah A yang tertunda di perangkatnya dan tiba sesudah F06 ditolak.
- **AC-6** — *Given* status suasana hati dan laporan sakit milik A, *When* A keluar, *Then* sejak F06 status dan kartunya tidak lagi terlihat oleh siapa pun, dan pada F07 catatan `CARE` milik A dihapus seluruhnya, kecuali A memasukkannya ke ekspornya (06 §2.5, F06, F07; ADR-007 butir 7; **[Usulan]**).
- **AC-7** — *Given* A keluar (F06), *When* anggota yang tersisa membuka keanggotaan, *Then* ia melihat bahwa A sudah keluar dan tanggalnya; Peristiwa keluar tercatat (06 PG-08); anggota yang tersisa menerima N-26 berprioritas `kabar` dengan teks umum di layar kunci, kecuali dokumen 10 menetapkan lain demi keluar dengan aman (06 F06, N-26, P-29) **[Dok. 10]**. Apakah layar yang tersisa menyebut masa tenggang dan kemungkinan A kembali ditetapkan dokumen 08 dan 10.
- **AC-8** — *Given* kedua anggota dewasa keluar satu per satu, *When* anggota terakhir keluar, *Then* ia ditawari ekspor seperti AC-1, dan nasib rumah tangga tanpa anggota mengikuti aturan retensi dokumen 10 (PA-FAM-5).
- **AC-9** — *Given* A keluar (F06), *When* anggota yang tersisa membuka chat rumah tangga, *Then* ia dapat membaca chat lama tetapi tidak ada yang dapat menulis di sana; sesudah F07 percakapan itu ditutup dan tetap terbaca olehnya; isi chat dalam ekspor A dan hak hapusnya **[Dok. 10]** (06 §16.2 "Peserta"; FR-COM-001/AC-5, AC-6).
- **AC-10** — *Given* keadaan CTH-38 (07 §7.11): A berutang pinjaman dana bersama 200000 dan punya hak ganti 60000, `KB` rumah tangga 4940000, *When* A keluar pada hari D dan masa tenggang lewat tanpa F08, *Then* sampai hari D + 7 kedua klaim tetap terbuka dan `KB` tetap 4940000; pada hari D + 7 `exit_closure` ditulis dengan angka yang sama seperti bila ditulis hari D (−200000 + 60000 = −140000 bagi rumah tangga; `KB` 4940000 − 140000 = 4800000; `UANG` tetap 4800000), dan baris "klaim ditutup karena anggota keluar" tampil di ringkasan periode yang memuat hari D + 7 (07 §8.2 butir 2b).
- **AC-11** — *Given* A dalam masa tenggang, *When* A menjalankan F08 dengan autentikasi ulang, *Then* keanggotaan, peran, izin, dan wewenangnya kembali seperti sebelum F06 tanpa F04; tidak ada yang berlaku surut (entri masa tenggang tetap dan tidak dinilai ulang; untuk syarat tanda A terhitung anggota dewasa selama masa tenggang, 07 §8.2 butir 2b; 06 §3.5 "Selama masa tenggang", PA-FAM-9 dijawab); penanda yang belum dijawab kini dapat dijawab A; tidak ada `exit_closure` (07 §8.2 butir 2b, kolom "Bila dibatalkan"). *When* F08 dikirim tanpa autentikasi ulang, oleh anggota lain, atau sesudah F07, *Then* ditolak (06 INV-A48 (b)).
- **AC-12** — *Given* A memilih hapus seketika pada layar keluar, *When* server menerima F06, *Then* F07 berjalan dalam transaksi yang sama (06 F06): `exit_closure`, penutupan dan penghapusan buku pribadi, penghapusan catatan `CARE`, dan penutupan chat baginya terjadi seketika, dan tidak ada masa tenggang yang dapat dibatalkan.
- **AC-13** — *Given* A mengirim F06 tanpa autentikasi ulang, misalnya dari sesi yang terbuka tanpa memasukkan kredensial, *When* server menerimanya, *Then* ditolak tanpa akibat (06 INV-A48 (a)).
- **AC-14** — *Given* A dalam masa tenggang (sesudah F06, sebelum F07 atau F08), *When* B mengirim `submit` permintaan baru, *Then* ditolak `no_approver` karena A bukan anggota dewasa aktif; *When* siapa pun mengirim K02 atas usulan kesepakatan yang menunggu, *Then* ditolak, karena tidak ada pemberi persetujuan; B tetap dapat menarik usulannya sendiri (K04) dan menolak usulan A (K03) (06 §3.2, §3.5 "Selama masa tenggang" dan tabel "Dari dua menjadi satu", INV-A48 (e); FR-REQ-002/AC-5). *When* B mencatat pemakaian 600000 di atas ambang 500000 selama itu, *Then* entri ditandai, karena rumah tangga tidak dinilai satu dewasa sebelum F07 (07 §8.2 butir 2b). *When* A kemudian menjalankan F08, *Then* tanda entri itu tidak berubah, dan A kini dapat menjawab penanda dan permintaan yang menunggunya.

### FR-FAM-013 — Ekspor data saat keluar

| | |
|---|---|
| Fitur | F-FAM-006 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` (keluar tanpa kehilangan data pribadinya; glosarium *keluar dengan aman*) |
| Aktor dan izin | Anggota yang akan keluar, atas datanya sendiri |
| Aturan acuan | 07 §8.2 butir 1, 1a, 2a, 2b (baris "Ekspor buku rumah tangga"), §8.3, dan tabel `exit_closure` baris "Sisi hukum" (ekspor memuat daftar klaim terbuka); ADR-007 butir 10; 06 AT-4 (kolom ekspor), V-9, LM-11 (indeks dokumen 05); format, cakupan rincian, dan jangka retensi: dokumen 10 **[Dok. 10]** |
| Business rule | BR-FAM-008, BR-FAM-011, BR-FAM-018, BR-FAM-019 |
| Kedudukan | Ekspor ditawarkan lebih dulu: 07 §8.2 (D-41, D-42 lewat "lanjut"). Ekspor buku pribadi yang ditutup sesudah keluar, hanya oleh pemiliknya, sampai dihapus: 07 §8.2 butir 2a (menjawab PA-FAM-7). Ekspor buku rumah tangga sesudah keluar, sampai saat akses berakhir: usulan `security-privacy` SP-48, 07 §8.2 butir 1a, diterima lewat "lanjut" di checkpoint S3, S3-4 (c) (menjawab PA-FAM-3). Jangka ketersediaan, jalannya, dan apakah rincian tulisan anggota yang tersisa ikut **[Dok. 10]**. Ekspor di luar keluar: Tier 3 (F-FIN-023) |

Ekspor ditawarkan, tidak diwajibkan, dan tidak menahan keluar. Isinya paling sedikit: seluruh buku pribadi orang itu, catatan `CARE` miliknya, rincian yang ia tulis, dan daftar klaim antar-pemilik yang masih terbuka beserta jumlahnya. Bila ia tidak mengekspor saat keluar dan tidak meminta penghapusan, buku pribadinya yang sudah ditutup tetap dapat ia ekspor sendiri sampai dihapus pada akhir jangka retensi dokumen 10; sampai itu tidak seorang pun di rumah tangga dapat membacanya (07 §8.2 butir 2a). **Buku rumah tangga** tetap dapat ia ekspor sesudah aksesnya berakhir, **sampai saat aksesnya berakhir** (F06): entri dan tautan yang diterima server sebelum saat itu, dengan saldo per saat itu. Entri yang diterima sesudahnya tidak termasuk, juga bila tanggal kejadiannya lebih awal (07 §8.2 butir 1a). Ekspor itu tidak pernah memuat data pribadi pasangannya, `read_state`, atau `presence` (LM-11).

**Kriteria penerimaan**

- **AC-1** — *Given* anggota akan keluar, *When* layar keluar terbuka, *Then* ekspor ditawarkan sebelum tombol keluar; menolak ekspor tidak menghalangi keluar.
- **AC-2** — *Given* ekspor dibuat, *When* isinya diperiksa, *Then* memuat buku pribadi orang itu, catatan `CARE` miliknya, dan daftar klaim terbuka beserta jumlahnya, dan tidak memuat apa pun milik anggota lain yang tidak terbaca olehnya di aplikasi (FR-FAM-008).
- **AC-3** — *Given* ekspor dibuat, *When* kolomnya diperiksa, *Then* tidak ada kolom hitungan atau tren per orang atas tanda, keberatan, koreksi, selisih kas, entri terlambat, pengingat, atau lama menanggapi (06 AT-4).
- **AC-4** — *Given* gagal membuat ekspor (misalnya jaringan putus), *When* anggota tetap ingin keluar, *Then* keluar tetap dapat dikukuhkan setelah aplikasi menyatakan bahwa ekspor belum berhasil dan bahwa buku pribadinya masih dapat ia ekspor sendiri sampai jangka retensi berakhir.
- **AC-5** — *Given* A sudah keluar tanpa meminta penghapusan, *When* anggota yang tersisa atau peran mana pun mencari buku pribadi A, *Then* tidak ditemukan; *When* A sendiri meminta ekspor lewat jalur yang ditetapkan dokumen 10, *Then* ia memperolehnya; *When* jangka retensi lewat, *Then* buku itu dihapus seperti FR-FAM-014/AC-1 (07 §8.2 butir 2a; ADR-007 butir 10) **[Dok. 10]**.
- **AC-6** — *Given* A keluar (F06) pada 12 November pukul 10.00 dengan `KB` rumah tangga 4940000, dan sesudahnya B mencatat belanja dana bersama 100000 bertanggal kejadian 11 November yang diterima server pukul 13.00, *When* A meminta ekspor buku rumah tangga lewat jalur dokumen 10, selama masa tenggang maupun sesudah F07, *Then* ekspor memuat entri yang diterima server sebelum 12 November pukul 10.00 dengan saldo per saat itu (`KB` 4940000), tidak memuat entri B bertanggal 11 November itu, dan tidak memuat `exit_closure` yang ditulis kemudian; angkanya tidak dapat berbeda dari buku rumah tangga per saat itu (07 §8.2 butir 1a, §8.3).
- **AC-7** — *Given* ekspor AC-6, *When* kolomnya diperiksa, *Then* tidak ada data pribadi B (buku pribadi, catatan `CARE`, termasuk status kepedulian B), `read_state`, `presence` B, atau kolom per orang yang dilarang AT-4; rincian tulisan B ikut atau tidak menurut dokumen 10, dan isi chat dalam ekspor menunggu dokumen 10 **[Dok. 10]**.
- **AC-8** — *Given* A sudah keluar (F06), *When* A meminta ekspor, *Then* sesuai 06 §3.5 (rumusan ringkas `architect`): "Sesudah F06, yang keluar tidak punya akses kecuali satu jalan ekspor berautentikasi ulang: buku rumah tangga per saat aksesnya berakhir dan buku pribadinya sendiri (07 §8.2 butir 1a, 2a)." Jalan itu meminta kredensial akun, bukan sesi yang terbuka, dan tanpa itu ditolak; ia tidak membuka layar rumah tangga mana pun; hasilnya tidak memuat apa pun yang diterima server sesudah F06, juga bila tanggal kejadiannya lebih awal (AC-6), dan tidak memuat kehadiran atau status kepedulian pasangan (AC-7). Berapa lama dan lewat jalan apa ia tersedia **[Dok. 10]** (06 INV-A48 (c), **[Usulan]** `architect`).

### FR-FAM-014 — Menghapus data pribadi atas permintaan

| | |
|---|---|
| Fitur | F-FAM-006 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; brief §7 "legitimate deletion rights" |
| Aktor dan izin | Anggota itu sendiri; bersama keluar, atau sesudahnya selama akunnya ada (jalurnya sesudah keluar: dokumen 10) |
| Aturan acuan | 07 §8.2 butir 2, 2a sampai 7, §8.3; 06 §2.5; ADR-007 butir 1 sampai 8 dan 10 |
| Business rule | BR-FAM-010, BR-FAM-011, BR-FAM-018 |
| Kedudukan | Keputusan 24 (S2-8 "a", diputuskan sendiri, sementara sampai advokat menjawab L-04, L-20, L-22, L-23); rincian D-41, D-42 (lewat "lanjut"). Penghapusan seketika bila diminta, selain itu pada akhir jangka retensi: 07 §8.2 butir 2a, ADR-007 butir 10 (menjawab PA-FAM-7); jangkanya **[Dok. 10]**. "Seketika" berarti pada F07 yang dijalankan dalam transaksi F06 bila yang keluar memilih hapus seketika (06 F06; 07 §8.2 butir 2b; **[Usulan]** SP-48); tanpa pilihan itu permintaan yang diajukan saat keluar dijalankan pada F07 akhir masa tenggang. Sisi hukum **[Perlu validasi]** |

**Kriteria penerimaan**

- **AC-1** — *Given* anggota yang keluar dan meminta penghapusan, *When* permintaan diproses pada F07 (seketika bila ia memilih hapus seketika, selain itu pada akhir masa tenggang; FR-FAM-012/AC-3, AC-12), *Then* buku pribadinya dihapus seluruhnya, tanpa menunggu jangka retensi (07 §8.2 butir 2a, 2b), rincian yang ia tulis pada entri dihapus kecuali rincian entri dana bersama yang juga data anggota tersisa (07 §8.2 butir 4), dan pemetaan identitasnya diputus.
- **AC-2** — *Given* penghapusan selesai, *When* anggota yang tersisa membaca buku rumah tangga, *Then* inti moneter setiap entri identik dengan sebelumnya (hash inti sama), angka rumah tangga tidak berubah, dan pelaku yang menunjuk orang itu tampil "mantan anggota" (ADR-007 bagian Konsekuensi).
- **AC-3** — *Given* catatan `CARE` milik orang itu, *When* penghapusan selesai, *Then* tidak ada baris atau Peristiwa `CARE` miliknya yang tersisa (ADR-007 butir 7).
- **AC-4** — *Given* layar permintaan penghapusan, *When* dibaca, *Then* ia menyatakan terus terang dua batas: di rumah tangga dua orang, "mantan anggota" tetap dikenali pasangannya; dan jumlah serta tanggal belanja dana bersama tetap ada (07 §8.3).
- **AC-5** — *Given* permintaan penghapusan, *When* diproses, *Then* tidak ada syarat berupa tindakan anggota lain (pelunasan, konfirmasi, persetujuan) (ADR-007 butir 8).
- **AC-6** — *Given* salinan cadangan, *When* jangka yang ditetapkan dokumen 10 lewat, *Then* penghapusan rincian dan pemutusan pemetaan juga berlaku di cadangan (07 §8.2 butir 6) **[Dok. 10]**.
- **AC-7** — *Given* A keluar dengan meminta penghapusan tanpa memilih hapus seketika, *When* A membatalkan lewat F08 di dalam masa tenggang, *Then* tidak ada yang sudah terhapus, permintaan penghapusan itu gugur bersama keluarnya, dan buku pribadinya menerima baris lagi seperti biasa (07 §8.2 butir 2b kolom "Bila dibatalkan").

### FR-FAM-015 — Objek terbuka sesudah seorang anggota keluar

| | |
|---|---|
| Fitur | F-FAM-006 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` (yang tersisa tidak tertinggal dengan hal yang tidak dapat ditutup); `MSL-1` |
| Aktor dan izin | Anggota yang tersisa |
| Aturan acuan | 06 §3.5 tabel "Objek saat seorang anggota dewasa keluar", INV-A35; 07 §8.2 butir 5 dan 8 |
| Business rule | BR-FAM-004 |
| Kedudukan | D-42 (lewat "lanjut", Pagar) |

Kebutuhan ini hanya menyatakan apa yang harus terlihat bagi anggota yang tersisa; aturannya kanonis di 06 §3.5.

**Kriteria penerimaan**

- **AC-1** — *Given* objek pada setiap state (permintaan, usulan, baris custody belum dikonfirmasi, keberatan, permintaan tindakan) dan satu anggota keluar, *When* daftar "yang menunggu" anggota yang tersisa dibuka, *Then* daftar itu hanya memuat objek yang dapat ia tindak sendiri; objek yang hanya dapat ditutup orang yang keluar tampil dengan keterangan "sudah keluar" pada objeknya (06 INV-A35).
- **AC-2** — *Given* permintaan yang menunggu dengan anggota yang tersisa sebagai pemohon, *When* pemberi persetujuannya keluar, *Then* anggota yang tersisa tetap dapat membatalkannya (06 §3.5).
- **AC-3** — *Given* tempat uang dana bersama yang dipegang anggota yang keluar, *When* anggota yang tersisa membukanya, *Then* saldo tetap terhitung di angka utama dengan keterangan pemegang sudah keluar, dan ia dapat mencatat apa yang sebenarnya terjadi lalu mengarsipkannya setelah bersaldo nol (07 §8.2 butir 5; INV-21).
- **AC-4** — *Given* ambang yang berlaku dan kini tinggal satu anggota dewasa sesudah F07, *When* ia mencatat pemakaian di atas ambang, *Then* tidak ada tanda (06 §3.5). Selama masa tenggang yang keluar masih terhitung anggota dewasa untuk syarat tanda, jadi pemakaian di atas ambang tetap ditandai (07 §8.2 butir 2b; 06 §3.5 "Selama masa tenggang" dan baris ambang "Dari dua menjadi satu"; PA-FAM-9 dijawab).

### FR-FAM-016 — Tidak ada jalan mengeluarkan orang dewasa lain

| | |
|---|---|
| Fitur | F-FAM-004, F-FAM-006 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; brief §8 abuse prevention |
| Aktor dan izin | Tidak ada aktor: perintahnya tidak ada |
| Aturan acuan | 06 PG-13 (dipersempit pada revisi kesembilan), §3.6 (F06 sampai F08), INV-A48 (a)(b); PRN-10 contoh pelanggaran ("pemilik rumah tangga mengeluarkan pasangannya") |
| Business rule | BR-FAM-005, BR-FAM-016 |
| Kedudukan | **[Usulan]** `product-analyst`, dikukuhkan `architect` sebagai PG-13 yang dipersempit dan INV-A48 (06 §3.6). Tantangan atas brief §4.2 "member removal" untuk Tier 1 (06 P-23), **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026), S3-4 (a); brief §4.2 diubah menjadi "member exit …" 12 November 2026, diterima atas saran lead (K-41) |

06 PG-13 semula membolehkan mengeluarkan orang dewasa lain dengan persetujuan orang itu. Persetujuan untuk dikeluarkan sama artinya dengan keluar sendiri, tetapi menambah satu tombol yang dapat dipakai menekan ("aku sudah minta kamu keluar"). Revisi kesembilan 06 mempersempitnya: Tier 1 tidak punya perintah mengeluarkan, menonaktifkan, membekukan, atau menurunkan peran orang dewasa lain, juga tidak dengan persetujuannya; satu-satunya jalan adalah orang itu keluar sendiri (F06; FR-FAM-012). Undangan yang jatuh ke orang yang salah dihentikan sebelum keanggotaan aktif (FR-FAM-004 AC-3). Biaya yang diterima: bila pasangan tidak mau keluar, yang ingin berpisah keluar sendiri dengan ekspornya dan dapat membuat rumah tangga baru; aplikasi tidak memutus siapa yang berhak tinggal.

**Kriteria penerimaan**

- **AC-1** — *Given* rumah tangga dua anggota dewasa, *When* menu mana pun dibuka oleh peran mana pun, *Then* tidak ada pilihan mengeluarkan, menonaktifkan, membekukan, atau menurunkan peran anggota dewasa lain.
- **AC-2** — *Given* perintah keanggotaan apa pun atas anggota lain (mengeluarkan, menonaktifkan, membekukan, mengubah peran, izin, atau wewenang) dikirim langsung ke server oleh setiap pelaku, *When* diterima, *Then* semuanya ditolak tanpa akibat, apa pun perannya dan juga bila disertai pernyataan setuju dari orang itu (06 INV-A48 (a)). Hal yang sama berlaku bagi F06 dan F08: hanya orang itu sendiri, dengan autentikasi ulang; F06 atau F08 atas nama orang lain ditolak (06 INV-A48 (a)(b); FR-FAM-012/AC-11, AC-13).
- **AC-3** — *Given* anggota yang tidak aktif berbulan-bulan, *When* jam dimajukan berapa pun, *Then* keanggotaannya tidak berubah dan tidak ada wewenang yang berpindah kepada anggota lain (06 PG-04).

### FR-FAM-017 — Kehilangan akses ke akun atau perangkat

| | |
|---|---|
| Fitur | F-FAM-004 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; `MSL-1` (yang lain tetap dapat bekerja); brief §8 "device lost or stolen", "admin loses access"; brief §12 nomor 20 |
| Aktor dan izin | Anggota yang kehilangan akses; anggota lainnya |
| Aturan acuan | 06 PG-02, PG-03, PG-13; autentikasi, sesi, perangkat, pemulihan akun: dokumen 10 **[Dok. 10]** |
| Business rule | BR-FAM-005 |
| Kedudukan | Akibat langsung keputusan 4 dan default Q-04 (Q-04 "Jika jawabannya lain": admin kehilangan akses bukan risiko Tier 1 bila keduanya pemilik). Mekanisme: dokumen 10 |

Karena kedua anggota dewasa adalah pemilik dengan wewenang identik, kehilangan akses satu akun tidak pernah menghentikan rumah tangga. Yang dijamin kebutuhan ini adalah apa yang **tidak** terjadi; caranya memulihkan akses milik dokumen 10.

**Kriteria penerimaan**

- **AC-1** — *Given* A kehilangan akses ke akunnya (simetris), *When* B memakai aplikasi, *Then* B tetap dapat mencatat, menanggapi, dan membaca seperti biasa, dan objek yang menunggu tanggapan A tetap menunggu tanpa berubah (06 PG-04).
- **AC-2** — *Given* A kehilangan akses, *When* B mencari jalan apa pun, *Then* tidak ada jalan bagi B untuk membaca data pribadi A, menanggapi atas nama A, atau mengambil alih akun A (06 PG-02, PG-03).
- **AC-3** — *Given* perangkat A hilang, *When* A memulihkan akses lewat jalur dokumen 10, *Then* data dana bersama dan data pribadi A kembali seperti sebelumnya; entri yang belum terkirim dari perangkat yang hilang ditangani menurut dokumen 09 dan 10 **[Dok. 10]**.
- **AC-4** — *Given* anggota yang tidak akan pernah kembali (misalnya wafat), *When* anggota yang tersisa memerlukan penyelesaian, *Then* penanganannya mengikuti dokumen 10 (PA-FAM-6); sampai ada, berlaku AC-1 dan AC-2.

---

## 6. Pihak luar, pemberitahuan operator, dan fondasi

### FR-FAM-018 — Pihak luar tanpa login

| | |
|---|---|
| Fitur | F-FAM-007 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-3` (kiriman ibu, bantuan dan pinjaman kepada saudara dan ayah); `MSL-2` (titipan) |
| Aktor dan izin | Anggota dewasa mana pun membuat dan mengarsipkan; terlihat oleh keduanya (`spouses`) |
| Aturan acuan | 07 §2.6 "Entitas pendukung", §3.2 "Kiriman keluarga", §3.7, §8.2 butir 7; 06 §9.2 |
| Business rule | BR-FAM-012, BR-FAM-013 |
| Kedudukan | Tanpa login di pilot: Q-25 (lewat "lanjut", checkpoint S1). Minimisasi data: Q-28, **terbuka** (S4), bawaan dipakai. Arsip bukan hapus: keputusan 9 (diputuskan sendiri) dan Q-58 (disampaikan lead, tidak dibantah) |

Pihak luar dicatat dengan nama panggilan dan hubungan saja (misalnya "Ibu", "Adik"). Tidak ada nomor telepon, alamat, foto, atau tanggal lahir. Pihak luar dibuat sekali, dipilih kembali sesudahnya, dan tidak pernah dicocokkan dengan pihak luar di rumah tangga lain (dokumen 03 §4.3, **[Usulan]**).

**Kriteria penerimaan**

- **AC-1** — *Given* anggota dewasa mencatat kiriman, pinjaman, atau titipan, *When* penerimanya belum ada, *Then* ia dapat membuat pihak luar dari layar yang sama dengan isian nama panggilan (wajib) dan hubungan (opsional), tanpa isian lain.
- **AC-2** — *Given* pihak luar yang sudah punya catatan, *When* anggota "menghapusnya", *Then* pihak luar diarsipkan: hilang dari daftar pilihan, catatannya tetap; pihak luar yang masih punya piutang atau titipan bersaldo tidak dapat diarsipkan sebelum dinolkan (keputusan 9; D-26; 07 INV-21).
- **AC-3** — *Given* pihak luar, *When* aplikasi diperiksa, *Then* tidak ada jalan pihak luar itu masuk, menerima notifikasi, atau melihat apa pun.
- **AC-4** — *Given* pihak luar bernama sama di dua rumah tangga (Tier 3 kelak), *When* data diperiksa, *Then* tidak ada tautan, pencocokan, atau saran yang menghubungkan keduanya.

### FR-FAM-019 — Pemberitahuan akses pengelola server

| | |
|---|---|
| Fitur | F-FAM-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; brief §8 "transparency about who can see what" |
| Aktor dan izin | Setiap anggota rumah tangga, dengan syarat yang sama (termasuk anggota yang juga pengelola server) |
| Aturan acuan | 06 §9.4 "Batas semua larangan"; ADR-007 bagian Konsekuensi; dokumen 10 (bentuk dan bukti) **[Dok. 10]** |
| Business rule | BR-FAM-009 |
| Kedudukan | Pengelola server sekaligus anggota, dengan syarat istrinya diberi tahu dengan terus terang sebelum mulai memakai: **keputusan 25** (diputuskan pemilik produk sendiri, S2-10 "a"). Bentuknya (membaca sendiri di aplikasi sebelum mengisi data pribadi; pengakuan bertanggal): SP-42, **[Usulan]** `security-privacy`, Q-122, disajikan di checkpoint S4. Ditampilkan kepada setiap anggota, bukan hanya kepada yang bukan pengelola: **[Usulan]** `product-analyst` (PRN-09) |

Selama pilot, pengelola server adalah salah satu anggota rumah tangga, dan data pribadi pasangannya secara teknis dapat ia buka dari server. Larangan di FR-FAM-008 tidak melindungi dari itu. Kebutuhan ini memastikan pemberitahuan terjadi di dalam aplikasi, terbaca sebelum data pribadi diisi, dan tercatat. Teksnya ditulis dokumen 10. Kebutuhan ini sengaja tidak membaca siapa pengelolanya: setiap anggota membaca teks yang sama, sehingga aturan tetap simetris.

**Kriteria penerimaan**

- **AC-1** — *Given* anggota yang baru membuat rumah tangga atau menerima undangan, *When* ia sampai pada langkah sebelum keanggotaannya aktif, *Then* ia membaca pemberitahuan dokumen 10 secara utuh di aplikasi dan menekan tanda sudah membaca; tanpa itu keanggotaannya tidak aktif dan ia tidak dapat mengisi data apa pun.
- **AC-2** — *Given* tanda sudah membaca ditekan, *When* tersimpan, *Then* tercatat tanggal dan versi teks yang dibaca, sebagai catatan milik anggota itu yang terbaca olehnya (SP-42).
- **AC-3** — *Given* teks pemberitahuan berubah versi (misalnya sesudah keputusan enkripsi di S4, Q-57), *When* anggota berikutnya membuka aplikasi, *Then* ia membaca versi baru dan menandainya sebelum melanjutkan; versi lama dan tanggalnya tetap tercatat.
- **AC-4** — *Given* anggota mana pun, *When* ia membuka penjelasan "siapa melihat apa" (FR-FAM-010), *Then* pemberitahuan yang sama dapat dibaca ulang kapan saja.
- **AC-5** — *Given* teks di layar mana pun yang menjelaskan privasi, *When* ditinjau, *Then* tidak ada kalimat yang menyiratkan bahwa larangan di aplikasi, V-9, atau pseudonimisasi melindungi dari pengelola server (06 §9.4; 06 V-9).
- **AC-6** — *Given* tahap publik, saat pengelola bukan anggota, *When* kebutuhan ini dijalankan, *Then* teksnya menjadi pemberitahuan operator umum menurut dokumen 10; mekanisme membaca dan mencatat tetap sama.

### FR-FAM-020 — Fondasi: isolasi rumah tangga, akun terpisah dari anggota

| | |
|---|---|
| Fitur | F-FAM-001 (fondasi untuk F-FAM-009, Tier 3) |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | Tidak ada masalah pilot; diwajibkan brief §1.1 "the Tier 1 data model and access model must not block Tier 3" dan `CLAUDE.md` (fondasi tahap kedua tidak dipotong) |
| Aktor dan izin | Server; diuji otomatis |
| Aturan acuan | 06 §14, ADR-003; PRN-03; `ENGINEERING-STANDARDS.md` §4; D-45 |
| Business rule | BR-FAM-002, BR-FAM-014 |
| Kedudukan | Buku pribadi per anggota: D-45 (lewat "lanjut"). Selebihnya brief §1.1 dan PRN-03 |

Kebutuhan model, bukan layar. Pilot tidak menawarkan rumah tangga kedua (FR-FAM-001 AC-5), tetapi ujinya ada sejak Tier 1.

**Kriteria penerimaan**

- **AC-1** — *Given* dua rumah tangga di satu server, *When* anggota rumah tangga 1 mengirim baca atau tulis apa pun dengan pengenal objek rumah tangga 2, *Then* ditolak, dan tidak ada tanggapan yang membedakan "tidak ada" dari "tidak boleh" (`ENGINEERING-STANDARDS.md` §4).
- **AC-2** — *Given* satu akun yang (lewat data uji) menjadi anggota dua rumah tangga, *When* uji AC-1 dijalankan atas tiap keanggotaannya, *Then* hasilnya sama; tidak ada data, angka, buku pribadi, notifikasi, atau keluaran AI satu rumah tangga yang terbaca lewat keanggotaan di rumah tangga lain.
- **AC-3** — *Given* skema, *When* ditinjau, *Then* setiap objek membawa `household_id`; akun dan anggota adalah dua entitas; buku pribadi melekat pada anggota; tidak ada kolom atau tabel bermakna "suami" atau "istri" (PRN-03; 06 §14).

---

## 7. Diagram keanggotaan

State machine kanonisnya 06 §3.6 (F01 sampai F08). Diagram di bawah sudut pandang satu anggota, dengan nomor transisi 06 dan FR yang menguji layarnya; satu-satunya tambahan atas 06 adalah jalur membuat rumah tangga (FR-FAM-001), yang tidak melalui undangan.

```mermaid
stateDiagram-v2
    [*] --> undangan_terbuka: F01 undangan dibuat (FR-FAM-003)
    undangan_terbuka --> gugur: F03 dibatalkan pengundang, atau masa berlaku lewat
    undangan_terbuka --> menunggu_konfirmasi_pengundang: F02 penerima membaca dan menerima; N-24 (FR-FAM-004)
    menunggu_konfirmasi_pengundang --> gugur: F05 pengundang memilih "bukan", penerima menarik, atau masa berlaku lewat
    menunggu_konfirmasi_pengundang --> aktif: F04 hanya pengundang, "ya, ini orangnya"; N-25
    [*] --> aktif: membuat rumah tangga (FR-FAM-001)
    aktif --> keluar_tenggang: F06 orang itu sendiri, autentikasi ulang; akses berakhir; N-26 (FR-FAM-012)
    keluar_tenggang --> aktif: F08 hanya orang itu sendiri, autentikasi ulang, tanpa F04 (FR-FAM-012/AC-11)
    keluar_tenggang --> keluar: F07 system, 7 hari sesudah F06, atau seketika bila memilih hapus seketika; exit_closure (FR-FAM-012/AC-3, AC-12)
    gugur --> [*]
    keluar --> [*]
```

Tidak ada transisi `aktif → keluar_tenggang` yang berpelaku orang lain, waktu, atau `system`; F07 adalah satu-satunya transisi `system`, hanya berangkat dari F06 orang itu sendiri, dan tidak ada transisi `system` menuju `aktif` (FR-FAM-016; 06 PG-04, PG-13, INV-A48). Selama ada anggota dalam `keluar_tenggang`, F01 ditolak (FR-FAM-003/AC-7). Pemberitahuan operator (FR-FAM-019) adalah syarat masuk ke `aktif` dari kedua jalur.

---

## 8. Yang tidak diatur di sini

### 8.1 [Perlu aturan]

Butir yang aturannya belum ada di 06, 07, atau 10. Pengenal `PA-FAM-n` (skema `PA-<MOD>-N`, README dokumen; dulu `P-FAM-n`, diganti pada putaran perbaikan fase D, CA-06) hanya untuk berkas ini; nomor `Q` diberikan pemilik dokumen 17.

**Sudah dijawab** (tidak dihapus agar rujukan tetap dapat dilacak):

| # | Jawaban | Kedudukan | Dipakai di |
|---|---|---|---|
| PA-FAM-1 | 06 §3.1 baris "Hubungan": tidak diisi; tidak ada baris yang ditulis, tabelnya tetap ada | **[Usulan]** `architect` | FR-FAM-001/AC-2, FR-FAM-006/AC-3 |
| PA-FAM-2 | 06 §3.5 tabel, §6.3 `rule`, K01, K05: zona waktu diisi saat dibuat; penggantiannya kesepakatan `rule`, berlaku mulai periode berikutnya | **[Usulan]** `architect` | FR-FAM-011/AC-2 sampai AC-4 |
| PA-FAM-3 | 07 §8.2 butir 1a, §8.3: yang keluar dapat mengekspor buku rumah tangga sesudah aksesnya berakhir, sampai saat aksesnya berakhir; jangka, jalan, dan cakupan rincian tulisan anggota yang tersisa dokumen 10 | Usulan `security-privacy` SP-48, diterima lewat "lanjut" di checkpoint S3, S3-4 (c); cakupan rincian tulisan tetap dokumen 10 | FR-FAM-013/AC-6, AC-7 |
| PA-FAM-4 | 06 §3.6 dan §11.2: N-24 `perlu_tanggapan` kepada pengundang; N-25 `kabar` kepada anggota baru; N-26 `kabar` kepada yang tersisa, apakah dan kapan ditinjau dokumen 10 | **[Usulan]** `architect` | FR-FAM-004/AC-2, AC-4; FR-FAM-012/AC-7 |
| PA-FAM-7 | 07 §8.2 butir 2a, T-37; ADR-007 butir 10: ditutup pada setiap keluar; dihapus seketika bila diminta, selain itu pada akhir jangka retensi; sampai itu hanya pemiliknya yang dapat mengekspor. Waktunya kini F07 (07 §8.2 butir 2b) | Bawaan `product-analyst` diambil `finance-domain` (07 §10.2 butir 4); jangka retensi dan cara ekspor **[Dok. 10]** | FR-FAM-012/AC-3, FR-FAM-013/AC-4, AC-5, FR-FAM-014/AC-1 |
| PA-FAM-8 | 06 §3.5 tabel baris "Zona waktu rumah tangga" (revisi kesepuluh): selama rumah tangga belum pernah punya dua anggota dewasa aktif, satu-satunya anggota dewasa dapat mengganti zona waktu sendiri, berlaku mulai periode berikutnya; sesudah itu hanya lewat `rule` | **[Usulan]** `architect`, mengikuti bawaan berkas ini | FR-FAM-011/AC-2, AC-5 |
| PA-FAM-9 | Dulu **pertentangan** 06 F06 dan F08 lawan 07 §8.2 butir 2b. Kini 06 §3.5 "Selama masa tenggang" dan INV-A48 (e) (revisi kesepuluh, §15.12 "Tindak lanjut"): "Selama masa tenggang, akses berakhir pada F06: yang keluar tidak membaca, tidak menulis, tidak terlihat, tidak menerima notifikasi, dan bukan pemberi persetujuan. Dalam aturan uang ia tetap `member(M)` sampai F07: tetap terhitung anggota dewasa untuk tanda, tetap pemegang tempat uangnya, buku pribadinya dibekukan (07 §8.2 butir 2b). Rumah tangga tidak dinilai satu dewasa sebelum F07. Bila F08: entri masa tenggang tetap dan tidak dinilai ulang; penanda yang belum dijawab kini dapat ia jawab." Rumusan "dinilai sebagai rumah tangga satu dewasa" pada F08 dicabut | **[Usulan]** `architect`, mengikuti 07 dan bawaan berkas ini | FR-FAM-012/AC-4, AC-11, AC-14; FR-FAM-015/AC-4; FR-REQ-002/AC-5 |

**Masih terbuka:**

| # | Yang belum diatur | Bawaan yang diusulkan **[Usulan]** | Pemutus |
|---|---|---|---|
| PA-FAM-5 | Nasib rumah tangga dan datanya ketika anggota terakhir keluar | Ekspor ditawarkan, lalu seluruh data rumah tangga dihapus setelah jangka retensi dokumen 10 | `security-privacy` (dokumen 10) |
| PA-FAM-6 | Anggota yang tidak akan kembali (wafat, tidak dapat dihubungi selamanya): apa yang dapat dilakukan anggota yang tersisa atas objek yang menunggu orang itu dan atas datanya | Tidak ada tindakan atas nama orang itu; objeknya tampil seperti anggota yang keluar (06 §3.5) setelah prosedur dokumen 10 yang melibatkan bukti di luar aplikasi | `security-privacy` (dokumen 10) |

### 8.2 Tantangan terhadap brief

| Brief | Usulan | Alasan |
|---|---|---|
| §4.2 "member removal" | Tier 1 tanpa perintah mengeluarkan orang dewasa lain (FR-FAM-016); dikukuhkan `architect` sebagai PG-13 yang dipersempit dan INV-A48 (06 P-23); diterima lewat "lanjut" di checkpoint S3, S3-4 (a); brief §4.2 diubah 12 November 2026, diterima atas saran lead (K-41) | Dengan PG-13, mengeluarkan sama dengan keluar sendiri ditambah satu tombol penekan. Kasus undangan yang salah ditutup sebelum keanggotaan aktif (F04) |
| §4.3 "each family configures its own household rules" | Tanpa layar pengaturan izin di pilot (FR-FAM-007; Q-30) | Dua orang dewasa setara tidak punya apa pun untuk diatur; model tetap menyimpan kelima konsep |

### 8.3 Untuk dokumen 17 (diajukan ke lead)

| Usulan | Pertanyaan atau asumsi | Bawaan | Kapan |
|---|---|---|---|
| Pertanyaan (dijawab) | Q-30: pilot tanpa layar pengaturan izin | Ya | Diterima lewat "lanjut" di checkpoint S3 (S3-1 (a)) |
| Pertanyaan (dijawab) | Tier 1 tanpa perintah mengeluarkan orang dewasa lain, juga dengan persetujuannya; undangan dikukuhkan pengundang sebelum aktif (06 P-23) | Ya | Diterima lewat "lanjut" di checkpoint S3 (S3-4 (a), (b)); dinilai `security-privacy` di S4 |
| Pertanyaan | Catatan `CARE` milik orang yang keluar dihapus saat keluar kecuali ia mengekspornya | Ya | `security-privacy` (S4) |
| Pertanyaan | N-26: anggota yang tersisa diberi tahu bahwa pasangannya keluar | Ya, `kabar` dengan teks umum; dokumen 10 dapat menunda atau meniadakannya demi keluar dengan aman | `security-privacy` (S4) |
| Pertanyaan (dijawab) | PA-FAM-8: zona waktu dapat dibetulkan sendiri selama rumah tangga belum pernah punya dua anggota dewasa | Diterapkan 06 §3.5 revisi kesepuluh **[Usulan]** | Lead dapat menutup barisnya di dokumen 17 |
| Pertanyaan | Masa tenggang keluar 7 hari, autentikasi ulang, pembatalan hanya oleh yang keluar, ekspor buku rumah tangga sesudah keluar, anggota yang bergabung kemudian tidak membaca rincian mantan anggota dan chat lama (SP-48, SP-49) | Ya | Diterima lewat "lanjut" di checkpoint S3 (S3-4 (c), D-90); 7 hari **[Asumsi]** |
| Pertentangan (diselesaikan) | PA-FAM-9: 06 F06 dan F08 lawan 07 §8.2 butir 2b tentang hitungan anggota dewasa selama masa tenggang | Diselesaikan 06 §3.5 "Selama masa tenggang", INV-A48 (e), mengikuti 07 | Lead dapat mencatat penyelesaiannya di dokumen 17 |
| Asumsi | Masa berlaku undangan 7 hari | Diuji di pilot; angka akhir dokumen 10 | S4 |
| Pertentangan (selesai) | 06 §3.1 (nilai hubungan "suami, istri" pada preset pilot) lawan dokumen 04 F-FAM-003 (tanpa isian hubungan) | Diselesaikan 06 §3.1 revisi kesembilan: tidak diisi (PA-FAM-1). Lead dapat menutup barisnya di dokumen 17 | — |
