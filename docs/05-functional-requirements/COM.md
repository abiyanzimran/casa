# 05 — Kebutuhan fungsional: `COM` (chat rumah tangga dan kartu objek)

**Sesudah checkpoint S4 (12 November 2026; ditinjau `critic` A dan B, 12 November 2026; temuan diperbaiki; perbaikan diperiksa lead dengan `grep`, tanpa tinjauan ketiga):** S4-7 A diputuskan pemilik produk sendiri, jadi FR-COM-019 dan FR-COM-020 Tier 1 Must tanpa syarat (D-114); FR-COM-019/AC-7 gugur. Kehadiran dan "terakhir aktif" tetap seperti keputusan pemilik produk (S4-5 A, A): kriteria FR-COM-013 yang ditulis bersyarat pada 10-A dan UX-5 berlaku cabang A, menurut tabel indeks 05 §2 "Sesudah checkpoint S4".

Bagian dari dokumen 05. Indeks, aturan umum, kebutuhan lintas modul (`LM`), dan tabel keterlacakan ada di [../05-functional-requirements.md](../05-functional-requirements.md). Pemilik: `product-analyst`. Tahap S3, putaran penyelarasan sesudah fase B, 12 November 2026; putaran perbaikan fase D (12 November 2026) atas tinjauan `critic` A (CA-05, CA-06), pemeriksaan `security-privacy` (SP-45, SP-49, SP-55, SP-57, SP-60), dan 06 revisi kesepuluh. Perbaikan fase D belum ditinjau ulang. **Sesudah checkpoint S3 (12 November 2026):** jawaban "lanjut" dicatat; P-21, P-22, P-24, P-25, P-32 (S3-7, S3-8 (vii), D-71, D-72) dan SP-48, SP-49 (S3-4 (c), D-90) kini "diterima lewat "lanjut"", dengan penulis aturannya tetap disebut; cabang alternatif (b) layar kunci tidak berlaku; angka **[Asumsi]** tidak berubah. **S4 fase D (12 November 2026, belum ditinjau `critic`):** FR-COM-019 (sembunyikan untuk saya) dan FR-COM-020 (laporkan) baru di bawah F-COM-014, bersyarat pada keputusan 10-G; BR-COM-019, BR-COM-020; FR-COM-013/AC-9, AC-10 bersyarat pada 10-A dan UX-5; FR-COM-006/AC-6 kelompok kartu tagihan rutin (08 §10.4); FR-COM-001/AC-2 menunjuk FR-FAM-002/AC-6; isi chat dalam ekspor kini 10.3 SEC-102 (f).

| Hal | Isi |
|---|---|
| Lingkup | F-COM-001, F-COM-008 sampai F-COM-012, F-COM-014 (Tier 1, kebutuhan lengkap dengan kriteria penerimaan); sisi chat dari F-CARE-003 (Tier 2, Must bersyarat, FR-COM-016) |
| Rentang ID | `FR-COM-001` sampai `FR-COM-020`; `BR-COM-001` sampai `BR-COM-020` |
| Sumber aturan | Alur, invarian, dan pagar: [06](../06-activity-approval-architecture.md) §16 (chat), §13.5 dan §13.6 (kepedulian dan pendamping `AI`), §9.4 AT-4 dan AT-5, §10 PG-04, PG-07, INV-A33, §11 (N-23, N-28), §12 S-11; [ADR-011](../adr/ADR-011-chat-rumah-tangga-dan-kartu-objek.md). Uang: tidak ada aturan uang di berkas ini; setiap angka pada kartu adalah angka objeknya menurut [07](../07-financial-domain.md) |
| Bukan isi dokumen ini | Tata letak, bentuk kartu, kelompok kartu, kata layar, dan desain visual (dokumen 08, S4); transport waktu nyata, penyimpanan pesan dan foto, mekanisme pengakuan penerimaan (dokumen 09, S4); ekspor, retensi, hak hapus isi chat, dan penilaian kehadiran sebagai risiko pengawasan (dokumen 10, S4) |

## Cara membaca

- **Bentuk tiap FR** sama dengan [REQ.md](REQ.md): tabel kedudukan, uraian singkat, kriteria penerimaan *Given/When/Then* bernomor lokal (dirujuk sebagai `FR-COM-007/AC-2`), lalu baris **Dua arah**.
- **Pelaku.** "A" dan "B" adalah dua anggota dewasa aktif rumah tangga. Tidak ada AC yang membaca siapa suami atau istri, siapa berpenghasilan, atau siapa pemegang uang (06 PG-01, INV-A13; `LM-01`).
- **Angka contoh** mengikuti [REQ.md](REQ.md): ambang permintaan Rp500.000 **[Asumsi]**, permintaan makan Rp100.000 (brief §5D). Waktu adalah waktu server menurut zona waktu rumah tangga (06 S-6).
- **Chat adalah permukaan, bukan alur baru** (06 §16.2; ADR-011 butir 2). Setiap tindakan atas objek dari chat adalah perintah objek yang sudah ada dan diuji oleh FR modul objek itu ([REQ.md](REQ.md), [FIN-1.md](FIN-1.md), [FIN-2.md](FIN-2.md), [tier-2.md](tier-2.md)). Berkas ini hanya menambah kriteria bahwa jalannya identik (INV-A44).
- **Istilah model bukan kata layar.** "Kartu", "baris peristiwa", "tanda air", "kehadiran", "belum sepakat" adalah bahasa model atau bahasa pertanyaan lead; kata layarnya milik dokumen 08.

## Kedudukan keputusan chat

Dikutip dari [_progress.md](../_progress.md), alinea "Pernyataan pemilik produk tentang chat", "Penjelasan pemilik produk sesudahnya", dan tabel sesudah "Pernyataan pemilik produk tentang pendamping AI" (12 November 2026). Jawaban yang lebih akhir menggantikan yang lebih awal di tempat yang disebut di sana.

| Hal | Kedudukan |
|---|---|
| Chat ada; dari chat dapat dikirim tugas, pengeluaran, rencana; tanda dibaca dan terkirim wajib | **Pernyataan pemilik produk sendiri**: "chat itu ada/diskusi di chat bisa kirim task/pengeluaran/rencana, ada tanda di baca/delivery chat itu wajib loh" |
| Satu chat suami-istri dengan kartu objek asli; diskusi per objek tetap ada dan tampil juga di chat | Diterima atas saran lead ("Chat umum + kartu") |
| Chat Tier 1, versi ringan (teks, foto, kartu, terkirim dan dibaca, online dan mengetik) | Diterima atas saran lead ("Chat naik Tier 1"); **menggantikan** jawaban "Tier 2" |
| Balasan atau tautan ke kartu tercatat di riwayat objek; chat biasa tetap bebas | Diterima atas saran lead ("Reply nempel ke kartu") |
| Kartu bertanda terkirim dan dibaca seperti pesan lain; aplikasi tidak pernah mengirim "sudah dibaca tapi belum dijawab" dan tidak menghitung lama menjawab | Diterima atas saran lead ("Ikut tanda dibaca"); menggantikan "permintaan, persetujuan, dan tugas tanpa tanda dibaca" untuk kartu di chat. Hanya bagian **kartu** ini yang diterima atas saran lead; tanda pada **pesan chat** adalah pernyataan pemilik produk sendiri (baris pertama) |
| Terakhir aktif, sedang online, sedang mengetik tampil | **Diputuskan pemilik produk sendiri, berlawanan dengan saran lead** ("Semua seperti WhatsApp") |
| Pesan dari saran `AI` adalah pesan chat biasa yang membalas kartu status sakit atau suasana hati | Diterima atas saran lead ("Ke chat"); mengubah brief §5O ("not through a general chat") |
| Voice note, grup, reaksi, pencarian tidak termasuk | Disebut di jawaban "Tier 2" yang kini digantikan; bahwa pengecualian itu tetap berlaku sesudah naik ke Tier 1 adalah **[Usulan]** `architect` (06 §16.1) |
| Selebihnya (model kartu, tanda air, tanpa waktu baca, penyimpanan kehadiran, hapus pesan, notifikasi N-23, satu kartu hidup per status) | **[Usulan]** `architect` (06 §16; ADR-011), dikutip di sini tanpa diubah |
| Terkirim hanya oleh penarikan saat aplikasi dibuka; kehadiran hanya oleh pemakaian di depan dengan interaksi; kartu status hilang dari chat; anggota yang bergabung kemudian mendapat chat baru; semua metadata foto dibuang | Usulan `security-privacy` (SP-55, SP-60, SP-45, SP-49, SP-57), diterapkan `architect` di 06 revisi kesepuluh; **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026) sebagai D-93, D-95, S3-7 (ii), D-90, D-94. Opsi kehadiran P1/P2 tetap S4 (Q-174) |

## Daftar FR

| ID | Judul | Fitur | Tier / MoSCoW |
|---|---|---|---|
| FR-COM-001 | Chat rumah tangga dan pesertanya | F-COM-001 | 1 / Must |
| FR-COM-002 | Mengirim pesan teks | F-COM-001 | 1 / Must |
| FR-COM-003 | Mengirim foto | F-COM-001 | 1 / Must |
| FR-COM-004 | Kartu yang lahir otomatis | F-COM-008 | 1 / Must |
| FR-COM-005 | Membuat atau mengirim objek dari chat | F-COM-008 | 1 / Must |
| FR-COM-006 | Isi kartu dan baris peristiwa | F-COM-008 | 1 / Must |
| FR-COM-007 | Menanggapi dari kartu | F-COM-009 | 1 / Must |
| FR-COM-008 | Kartu yang basi dan tanggapan bersamaan | F-COM-009 | 1 / Must |
| FR-COM-009 | Balasan yang menempel pada kartu | F-COM-010 | 1 / Must |
| FR-COM-010 | Yang tidak pernah menjadi kartu atau balasan | F-COM-008, F-COM-010 | 1 / Must |
| FR-COM-011 | Tanda terkirim dan dibaca | F-COM-011 | 1 / Must |
| FR-COM-012 | Tanpa tekanan dari tanda dibaca | F-COM-011 | 1 / Must |
| FR-COM-013 | Sedang online, sedang mengetik, terakhir aktif | F-COM-012 | 1 / Must |
| FR-COM-014 | Kehadiran tanpa riwayat dan tidak dibaca apa pun | F-COM-012 | 1 / Must |
| FR-COM-015 | Menghapus pesan sendiri | F-COM-001 | 1 / Must |
| FR-COM-016 | Kartu status kepedulian dan catatan dukungan | F-CARE-003 | 2 / **Must bila F-CARE-001 atau F-CARE-002 dibangun** |
| FR-COM-017 | Notifikasi chat (N-23) | F-COM-001 | 1 / Must |
| FR-COM-018 | Offline dan urutan pesan | F-COM-001, F-COM-009 | 1 / Must |
| FR-COM-019 | Sembunyikan untuk saya | F-COM-014 | 1 / Must (S4-7 A) |
| FR-COM-020 | Laporkan | F-COM-014 | 1 / Must (S4-7 A) |

## Business rule

Setiap baris adalah penunjuk ke aturan kanonis (`LM` indeks bagian 4). Bila bunyinya tampak berbeda dari sumber, sumber yang berlaku.

| ID | Pernyataan | Sumber kanonis | Kedudukan |
|---|---|---|---|
| BR-COM-001 | Tepat satu chat terbuka per rumah tangga; pesertanya anggota dewasa aktif saat percakapan dibuat, dibekukan di `participants`; rumah tangga satu dewasa tidak punya chat yang tampil; chat mulai kosong saat anggota kedua aktif, tanpa kartu surut. Sejak F06 yang keluar tidak membaca chat dan tidak ada yang menulis di sana (F08 memulihkannya); pada F07 percakapan ditutup dan bekas peserta yang tersisa tetap membacanya tanpa menulis. Anggota yang bergabung sesudah ada yang keluar mendapat percakapan baru; percakapan lama hanya terbaca bekas pesertanya; layar undangan menyebutnya | 06 §16.2 "Peserta"; §3.5; §3.6 F04, F06 sampai F08; INV-A47 (f); ADR-011 butir 1, 11 | Chat ada: pernyataan pemilik produk sendiri. Satu chat: diterima atas saran lead. Aturan peserta **[Usulan]** `architect`; anggota yang bergabung kemudian: usulan `security-privacy` SP-49, diterima lewat "lanjut" sebagai D-90 |
| BR-COM-002 | Chat adalah permukaan atas primitif Diskusi (`message`) dan Peristiwa; kartu adalah tampilan objek asli yang isinya dibaca dari state objek saat ini; kartu tidak punya state | 06 §16.2 "Isi kartu"; ADR-011 butir 2 | **[Usulan]** `architect` |
| BR-COM-003 | Kartu lahir otomatis dalam transaksi yang sama hanya pada B01, K01, T01, C01, dan saat status kepedulian yang dibagikan disimpan. Entri biasa menjadi kartu hanya bila dicatat atau dikirim dari chat. Satu kartu per objek di server | 06 §16.2 "Kartu otomatis" | Kartu permintaan di chat: diterima atas saran lead. Daftarnya **[Usulan]** `architect` |
| BR-COM-004 | Tombol kartu adalah perintah objek yang sudah ada (B, K, C, T, D), dikirim dengan versi syarat dan state yang diharapkan; tidak ada transisi, guard, atau pelaku baru; pihak yang tidak berhak tidak melihat tombolnya dan perintahnya ditolak server | 06 §16.4; S-3, S-4; INV-A44; ADR-011 butir 3 | Menanggapi dari kartu: diterima atas saran lead. Pemetaan **[Usulan]** `architect` |
| BR-COM-005 | Tidak ada pesan, apa pun isinya ("setuju", "oke", foto struk), yang mengubah state objek atau melahirkan otorisasi | 06 §16.3 "Kekuatan"; §2.1; INV-A44 | **[Usulan]** `architect` |
| BR-COM-006 | Pesan yang membalas kartu, atau membalas pesan yang sudah menempel pada objek yang sama, mendapat `subject_ref` objek itu dan tampil di chat dan di diskusi objek dengan isi yang sama. Penulis dapat menautkan pesannya sendiri ke satu kartu sesudah dikirim (`link_message`); tautan tidak dapat dilepas kecuali dengan menghapus pesannya. Pesan bebas tidak masuk riwayat objek mana pun; server tidak menebak hubungannya | 06 §16.3 | Menempel: diterima atas saran lead. `link_message` **[Usulan]** `architect` |
| BR-COM-007 | Hanya objek yang terlihat oleh semua peserta chat yang dapat menjadi kartu atau ditempeli. Tidak pernah: entri buku pribadi, isi dompet campuran, status atau laporan sakit yang tidak dibagikan, pertanyaan tindak lanjut sistem (I02), setelan jam tenang, tunda, atau apa pun yang hanya terbaca pemiliknya | 06 §16.2 "Yang tidak pernah"; INV-A47 (c); V-9 | **[Usulan]** `architect` |
| BR-COM-008 | Transisi atas objek yang punya kartu tampil sebagai baris peristiwa yang ditulis server dari Peristiwanya dalam transaksi yang sama, berpelaku sama; baris peristiwa dan pembuatan kartu tidak memicu notifikasi sendiri | 06 §16.2 "Baris peristiwa"; INV-A46 | **[Usulan]** `architect` |
| BR-COM-009 | Tanda terkirim dan dibaca hanya di chat, sama bagi kedua orang, sebagai tanda air per peserta yang hanya naik. **Terkirim** naik hanya saat perangkat penerima menarik pesan itu ketika aplikasinya dibuka di depan, menyimpannya, dan mengakuinya; push, Notification Service Extension, penerima pesan latar, dan sinkron latar tidak pernah menggerakkannya; satu aturan di semua platform. **Dibaca** = pesan tampil di layar chat saat aplikasi di depan. Membuka notifikasi atau layar objek tidak menggerakkan keduanya. Tidak ada waktu baca per pesan, yang tersimpan maupun yang tampil | 06 §16.5; AT-5 (a), (b); INV-A33; ADR-011 butir 5 | Tanda pada pesan chat: pernyataan pemilik produk sendiri ("ada tanda di baca/delivery chat itu wajib loh"). Tanda pada kartu: diterima atas saran lead ("Ikut tanda dibaca"). Arti terkirim: usulan `security-privacy` SP-55, diterapkan 06 revisi kesepuluh, diterima lewat "lanjut" sebagai D-93; dapat-tidaknya dijaga di kedua platform **[Asumsi]** yang diverifikasi dokumen 09. Tanda air dan tanpa waktu baca **[Usulan]** `architect` |
| BR-COM-010 | Tidak ada pengingat, lencana, warna, urutan, atau penyaring yang berasal dari "dibaca tapi belum dijawab"; tidak ada hitungan lama menjawab, hitungan pesan per orang, hitungan pesan yang belum dibaca pasangan, atau statistik chat apa pun. Lencana "belum dibaca" hanya milik pembacanya sendiri | 06 §16.5 "Yang tidak ada"; AT-4; PG-07; INV-A45 | Tanpa "dibaca tapi belum dijawab" dan tanpa hitungan lama menjawab: diterima atas saran lead. Selebihnya **[Usulan]** `architect` |
| BR-COM-011 | Kehadiran digerakkan **hanya oleh pemakaian di depan dengan interaksi** (aplikasi di depan dan orangnya menyentuh, menggulir, atau mengetik); sinkron latar, push, ekstensi notifikasi, widget, dan aplikasi yang terbuka tanpa disentuh tidak pernah menggerakkannya. Sedang online (aplikasi di depan, tersambung, dan dipakai), sedang mengetik (sementara, tidak pernah diantrekan), terakhir aktif (waktu server saat sesi pemakaian di depan dengan interaksi terakhir berakhir, satu nilai per anggota per rumah tangga yang ditimpa); tampil hanya di layar chat; tanpa riwayat di basis data, log, analitik, cadangan yang dapat dikueri, atau ekspor; tidak dibaca aturan, pengingat, ringkasan, notifikasi, atau fitur `AI`; tidak pernah memicu notifikasi | 06 §16.6; INV-A45; S-11 (d) | Kehadiran tampil: **diputuskan pemilik produk sendiri, berlawanan dengan saran lead**. Hanya oleh pemakaian di depan dengan interaksi: usulan `security-privacy` SP-60, diterapkan 06 revisi kesepuluh, diterima lewat "lanjut" sebagai D-95. Cara menyimpan **[Usulan]** `architect`. Lama tanpa interaksi yang mengakhiri "online" **[Asumsi]** dokumen 09 |
| BR-COM-012 | Jam tenang dan tunda tidak pernah menggerakkan tanda terkirim, tanda dibaca, atau kehadiran. Jam tenang menahan notifikasi di server dan melepas semua yang tertahan sebagai **satu** notifikasi saat berakhir; karena terkirim hanya naik oleh penarikan saat aplikasi dibuka (BR-COM-009), apa yang dikirim atau ditahan server selama jam tenang tidak terbaca lewat tanda. Yang terbaca pihak yang menunggu identik bila tindakan penerima sendiri sama | 06 §16.5 "Jam tenang dan tunda"; §10 "Yang tertahan"; INV-A33 | **[Usulan]** `architect` dan `security-privacy` (SP-55; bagian "tanpa push senyap" SP-39 ditarik `security-privacy`). Rumusan revisi kesembilan ("menahan bunyi, tidak menahan pengiriman data") dicabut; mekanisme per platform **[Asumsi]** yang diverifikasi dokumen 09 |
| BR-COM-013 | Tidak ada sunting pesan. Penulis menghapus pesannya sendiri kapan saja, bagi kedua orang; isi dan foto dibuang; tersisa penanda "pesan dihapus" beserta waktunya, di chat dan di diskusi objek bila pesan itu menempel; tanda air tidak turun; tanpa notifikasi. Kartu dan baris peristiwa tidak dapat dihapus siapa pun | 06 §16.7; INV-A47 (a), (b); P-22 | Usulan `architect` (P-22); diterima lewat "lanjut" di checkpoint S3 sebagai D-71 |
| BR-COM-014 | Status kepedulian yang dibagikan punya tepat satu kartu hidup per status saat ini. Saat status diganti, berakhir, dicabut, atau dihapus, **kartunya hilang dari chat bagi kedua orang tanpa penanda** (tanpa kartu kosong, tanpa "status sudah berganti", tanpa baris peristiwa); status baru melahirkan kartu baru; balasan yang menempel dilepas dari status dan tetap sebagai pesan bebas milik penulisnya. Satu-satunya kartu yang dapat hilang | 06 §13.5; §16.7 "Kartu dan baris peristiwa"; INV-A47 (b), (d), (e); P-21 | Catatan dukungan sebagai balasan chat: diterima atas saran lead. Kartu hilang: usulan `security-privacy` SP-45, diterapkan 06 revisi kesepuluh (menggantikan "kartu lama tampil tanpa isi"). Satu kartu hidup dan pelepasan balasan: usulan `architect` (P-21). Keduanya diterima lewat "lanjut" di checkpoint S3, S3-7 (ii) |
| BR-COM-015 | Pesan chat memicu N-23 `kabar` langsung kepada peserta lain, dilipat per pengirim (tidak ada N-23 baru dari pengirim yang sama dalam 10 menit sesudah yang terakhir dilepas selama chat belum dibuka), satu yang tertahan per pengirim selama jam tenang; di luar jendela 6 jam, plafon harian, dan batas pengingat; layar kunci memakai teks umum kelas `kabar` ("Ada kabar baru"), tanpa nama pengirim dan tanpa isi, juga bila penerima memilih rincian. Balasan yang menempel memicu N-23, bukan N-19 | 06 §16.8; §11.2 N-23; §11.1; INV-A46; P-24, P-32 | Usulan `architect`, diterima lewat "lanjut" sebagai D-72; angka 10 menit **[Asumsi]**. Teks layar kunci per kelas prioritas sejak 06 revisi kesepuluh (SP-46; revisi kesembilan: "Ada pesan baru"); bentuk (a) tanpa nama diterima lewat "lanjut" (S3-8 (vii)); alternatif (b) "ada kabar dari {nama}" tidak dipilih |
| BR-COM-016 | Pesan adalah perintah `send_message` ber-`idempotency_key` dengan antrean kirim lokal; urutan tampil menurut `seq` server per percakapan yang naik tanpa celah, kecuali `seq` kartu status yang sudah hilang (BR-COM-014); tindakan dari kartu tunduk pada S-1 sampai S-4; tanda air dikirim saat tersambung tanpa waktu jam perangkat; kehadiran tidak pernah diantrekan; aturan harus benar juga bila transportnya hanya push disusul tarik | 06 §12 S-11; S-1, S-2, S-5, S-6; INV-A47 (d) | **[Usulan]** `architect`; transport milik dokumen 09 |
| BR-COM-017 | Kartu yang dibaca dan tidak ditanggapi tetap menunggu dengan batas objeknya sendiri (kedaluwarsa, pengingat §10). Membaca tidak pernah menjadi tanggapan; diam tidak pernah menjadi keputusan | 06 §16.5 "Diam"; PG-04; INV-A10 | Keputusan pemilik nomor 6 (diam bukan jawaban); rumusan untuk chat **[Usulan]** `architect` |
| BR-COM-018 | Teks pesan, nama berkas foto, dan kehadiran tidak pernah ditulis ke log | 06 §16.7 "Log"; `ENGINEERING-STANDARDS.md` §4 | Standar teknik |
| BR-COM-019 | Menyembunyikan pesan bebas, foto, indikator mengetik, dan N-23 satu anggota hanya mengubah tampilan pembacanya; kartu objek, baris peristiwa, balasan yang menempel, dan notifikasi §11.2 tidak pernah tersembunyi; tanda dibaca tetap mengikuti tanda air; tanpa Peristiwa yang terbaca orang lain, tanpa notifikasi; setelannya privat dan tidak masuk ekspor orang lain | 06 §16.12, INV-A53; 10.4 SEC-153 | **[Usulan]** `security-privacy` dan `architect` S4; bentuk 10-G A diputuskan pemilik produk sendiri (S4-7 A) |
| BR-COM-020 | Laporan dikirim hanya bila pengelola bukan anggota rumah tangga pelapor; bila pengelola anggota rumah tangga itu, server menolak menyerahkannya dan aplikasi membuka Bantuan; laporan tidak pernah dialirkan kepada orang yang dilaporkan. Laporan berisi hanya yang dipilih dan ditulis pelapor; disimpan untuk operator lewat perintah operasi; yang dilaporkan tidak diberi tahu; operator tidak mendapat akses baca chat karena laporan; tidak ada pemindaian isi | 06 §16.12; 10.4 SEC-135, SEC-154 | **[Usulan]** `security-privacy` dan `architect` S4; bentuk 10-G A diputuskan pemilik produk sendiri (S4-7 A) |

---

## F-COM-001 — Chat rumah tangga: pesan teks dan foto

### FR-COM-001 — Chat rumah tangga dan pesertanya

| | |
|---|---|
| Fitur | F-COM-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` ("the other checks the budget and approves, discusses, or counters": bicara dan memutus di satu tempat); `MSL-7`; `KS-4` |
| Aktor dan izin | Peserta = anggota dewasa aktif (06 §3.1). Orang yang undangannya belum dikukuhkan pengundang tidak membaca apa pun (INV-A48 c) |
| Aturan acuan | 06 §16.2 "Peserta"; §3.5; §3.6 F04, F06 sampai F08; INV-A47 (f); ADR-011 butir 1, 11 |
| Business rule | BR-COM-001 |
| Kedudukan | Lihat "Kedudukan keputusan chat" baris 1 sampai 3. Masa tenggang dan percakapan baru bagi anggota yang bergabung kemudian: usulan `security-privacy` SP-48, SP-49, diterapkan 06 revisi kesepuluh; diterima lewat "lanjut" di checkpoint S3 (S3-4 (c), D-90). Isi chat dalam ekspor orang yang keluar: pesan dan foto kedua peserta sampai F06, tanpa kartu status (10.3 SEC-102 (f); FR-FAM-013/AC-7), **[Usulan]** `security-privacy`, L-33 **[Perlu validasi]** |

Satu garis waktu per rumah tangga yang memuat pesan bebas, pesan yang menempel pada objek, kartu objek, dan baris peristiwa. Tidak ada percakapan kedua, grup, atau percakapan pribadi di Tier 1 (F-COM-013, Tier 3).

**Kriteria penerimaan**
- **AC-1** — *Given* rumah tangga berisi A dan B yang aktif, *When* masing-masing membuka chat, *Then* keduanya melihat satu chat rumah tangga yang sama, dengan pesan dan kartu yang sama dalam urutan yang sama.
- **AC-2** — *Given* rumah tangga satu anggota dewasa, *When* A membuka aplikasi, *Then* tidak ada chat yang tampil, dan mencatat, anggaran, serta laporan berjalan penuh (06 §3.5). Tab Chat dan catatan penggantinya di Beranda: FR-FAM-002/AC-6.
- **AC-3** — *Given* A sudah mencatat dan membuat objek sebelum B bergabung, *When* B menjadi aktif lewat F04, *Then* chat mulai kosong bagi keduanya, tanpa kartu surut atas objek lama; objek itu tetap di layarnya sendiri dan di daftar "yang menunggu" (FR-REQ-043).
- **AC-4** — *Given* B diundang dan sudah menerima undangan tetapi A belum mengukuhkannya (F02, belum F04), *When* B mencoba membaca chat, *Then* tidak ada isi chat yang terbaca (INV-A48 c).
- **AC-5** — *Given* B keluar (F06, dengan autentikasi ulang), *When* A membuka chat selama masa tenggang, *Then* A dapat membaca pesan lama tetapi tidak dapat mengirim apa pun; perintah `send_message` yang tetap terkirim ditolak; B tidak lagi membaca chat dan tidak terlihat di sana (kehadiran, kartu statusnya) sejak F06.
- **AC-6** — *Given* AC-5, *When* B membatalkan keluarnya (F08) di dalam masa tenggang, *Then* chat yang sama kembali terbaca dan dapat ditulisi oleh keduanya, tanpa kartu atau pesan yang hilang; *When* masa tenggang lewat tanpa F08 (F07), *Then* percakapan ditutup (`closed_at` terisi): A tetap dapat membacanya, tidak ada yang dapat menulis di sana (06 §3.6, §16.2).
- **AC-7** — *Given* AC-6 sesudah F07 dan A kemudian mengundang C yang dikukuhkan (F04), *When* C dan A membuka chat, *Then* keduanya melihat percakapan baru yang kosong tanpa kartu surut; C tidak dapat membaca percakapan lama dengan cara apa pun (layar, API, ekspor), sedangkan A tetap dapat membacanya sebagai bekas peserta; tidak ada pesan atau kartu yang dipindahkan (06 §16.2, INV-A47 (f); SP-49; FR-FAM-004/AC-10, AC-11).

**Dua arah.** Simetris. Tidak ada pemilik chat; menukar A dan B pada setiap AC menghasilkan keadaan yang tertukar.

### FR-COM-002 — Mengirim pesan teks

| | |
|---|---|
| Fitur | F-COM-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `MSL-7` |
| Aktor dan izin | Peserta mana pun |
| Aturan acuan | 06 §16.2 (`message`), §16.3 "Pesan bebas", "Kekuatan"; §2.5 (pesan adalah *rincian*); S-2; INV-A44 |
| Business rule | BR-COM-005, BR-COM-006, BR-COM-016, BR-COM-018 |
| Kedudukan | Chat dengan diskusi: pernyataan pemilik produk sendiri ("chat itu ada/diskusi di chat bisa kirim task/pengeluaran/rencana, ada tanda di baca/delivery chat itu wajib loh"). Pesan bebas di samping balasan yang menempel: diterima atas saran lead, lewat pilihan "Reply nempel ke kartu"; "chat biasa tetap bebas" adalah bunyi pilihan lead itu, bukan kata pemilik produk (`_progress.md`). Panjang maksimum teks **[Dok. 09]** |

Pesan teks, boleh membalas pesan lain. Server memberi `seq` dan `sent_at`. Pesan tidak punya akibat uang atau state apa pun.

**Kriteria penerimaan**
- **AC-1** — *Given* A dan B, *When* A mengirim "aku pulang jam 7", *Then* pesan tersimpan dengan `seq` dan waktu server, tampil bagi keduanya pada urutannya, dan B menerima N-23 (FR-COM-017).
- **AC-2** — *Given* permintaan B Rp100.000 menunggu tanggapan A, *When* A mengirim "setuju, beli aja" sebagai pesan bebas (tidak membalas kartu), *Then* state permintaan, versi syarat, Peristiwa, dan otorisasi tidak berubah, dan pesan itu tidak masuk diskusi permintaan; server tidak menebak hubungannya (06 §16.3).
- **AC-3** — *Given* A membalas pesan bebas B, *When* balasan terkirim, *Then* balasan itu juga pesan bebas, tanpa `subject_ref`.
- **AC-4** — *Given* teks kosong atau hanya spasi, *When* A mengirimnya, *Then* server menolak dan tidak ada yang tersimpan.
- **AC-5** — *Given* pesan apa pun, *When* log server diperiksa, *Then* teksnya tidak ada di log mana pun.

**Dua arah.** Simetris.

### FR-COM-003 — Mengirim foto

| | |
|---|---|
| Fitur | F-COM-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` (memperlihatkan barang atau harga saat berunding); `MSL-6` (masalah rumah difoto) |
| Aktor dan izin | Peserta mana pun; foto lewat primitif Lampiran (06 §2.1) |
| Aturan acuan | 06 §16.2 (`attachment_ids`); AT-6; §16.7 "Log"; §11.1 (layar kunci) |
| Business rule | BR-COM-005, BR-COM-006, BR-COM-018 |
| Kedudukan | Foto di versi ringan: diterima atas saran lead. Pembuangan semua metadata, dikode ulang di perangkat dan diperiksa lagi di server: usulan `security-privacy` SP-57, diterima lewat "lanjut" sebagai D-94; daftar akhir medan metadata **[Dok. 10]**. Jenis, ukuran, dan jumlah foto per pesan **[Dok. 09]**. Foto di chat bukan lampiran struk entri: **[Usulan]** `product-analyst` |

**Kriteria penerimaan**
- **AC-1** — *Given* A dan B, *When* A mengirim foto dengan keterangan opsional, *Then* foto tampil bagi keduanya, B menerima N-23 dengan teks layar kunci umum, dan tidak ada pratinjau gambar di notifikasi.
- **AC-2** — *Given* foto yang membawa metadata (EXIF termasuk lokasi, waktu, dan model perangkat; XMP; thumbnail tertanam yang dapat menyimpan gambar sebelum dipotong), *When* A mengirimnya, *Then* perangkat A mengode ulang gambar dan membuang semua metadata itu sebelum mengunggah, server memeriksanya lagi dan membuang yang masih tersisa sebelum menyimpan, dan berkas yang tersimpan, yang diterima B, dan yang diekspor tidak memuat metadata apa pun (SP-57; AT-6). Aturan yang sama berlaku atas foto struk (FR-REQ-042/AC-4).
- **AC-3** — *Given* entri belanja A sudah tercatat dan punya kartu, *When* A mengirim foto struk sebagai balasan pada kartu itu, *Then* foto menjadi bagian diskusi entri (FR-COM-009) tetapi bukan lampiran struk entri dan tidak mengubah jumlah, realisasi, tautan, atau tanda; untuk melampirkannya sebagai struk, A memakai lampiran pada entri (FR-REQ-042).
- **AC-4** — *Given* foto apa pun, *When* log diperiksa, *Then* nama berkasnya tidak ada.

**Dua arah.** Simetris.

### FR-COM-015 — Menghapus pesan sendiri

| | |
|---|---|
| Fitur | F-COM-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5` (salah kirim tidak menakutkan); `PAGAR` (percakapan tidak dapat diubah diam-diam, PG-08) |
| Aktor dan izin | Hanya penulis, atas pesannya sendiri |
| Aturan acuan | 06 §16.7; INV-A47 (a), (b); PG-08; ADR-011 butir 9 |
| Business rule | BR-COM-013 |
| Kedudukan | Usulan `architect` (06 P-22), diterima lewat "lanjut" di checkpoint S3 (12 November 2026) sebagai D-71; harganya W-23. Penghapusan dari cache perangkat pasangan **[Dok. 09]** |

**Kriteria penerimaan**
- **AC-1** — *Given* pesan teks atau foto A, *When* A menghapusnya kapan saja, *Then* isi dan foto hilang bagi keduanya; yang tersisa penanda "pesan dihapus" beserta waktunya, di chat dan, bila pesan itu menempel, di diskusi objeknya; tidak ada notifikasi; tanda air siapa pun tidak turun.
- **AC-2** — *Given* pesan A, *When* B mengirim perintah menghapusnya, *Then* server menolak dan pesan tetap.
- **AC-3** — *Given* pesan siapa pun, *When* ada yang mencoba menyuntingnya, *Then* tidak ada perintah sunting; perintah seperti itu ditolak.
- **AC-4** — *Given* kartu objek atau baris peristiwa, *When* siapa pun mencoba menghapusnya, *Then* server menolak; objek hanya berakhir lewat perintahnya sendiri (INV-A47 b).
- **AC-5** — *Given* pesan A yang menempel pada permintaan, *When* A menghapusnya, *Then* syarat dan state permintaan tidak berubah, dan diskusi permintaan menampilkan penanda yang sama.
- **AC-6** — *Given* B pernah menerima pesan A di perangkatnya, *When* perangkat B tersinkron sesudah A menghapus, *Then* isi pesan tidak lagi terbaca di perangkat B.

**Dua arah.** Simetris; hak menghapus mengikuti penulis.

### FR-COM-017 — Notifikasi chat (N-23)

| | |
|---|---|
| Fitur | F-COM-001 (dibaca bersama F-NTF-001) |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` (layar kunci tidak membocorkan isi; chat yang ramai tidak menenggelamkan yang meminta tanggapan) |
| Aktor dan izin | Penerima: peserta lain. Pengirim tidak diberi tahu atas pesannya sendiri |
| Aturan acuan | 06 §16.8; §11.2 N-23; §11.1 "Isi layar kunci", "Bentuk chat dan notifikasi"; INV-A46 |
| Business rule | BR-COM-008, BR-COM-015 |
| Kedudukan | Usulan `architect`, diterima lewat "lanjut" sebagai D-72; angka 10 menit **[Asumsi]**; teks layar kunci bentuk (a) diterima lewat "lanjut" (S3-8 (vii); 06 P-32); rumusan layar kunci akhir **[Dok. 10]** (06 P-24) |

**Kriteria penerimaan**
- **AC-1** — *Given* B belum membuka chat, *When* A mengirim tiga pesan dalam 5 menit, *Then* B menerima satu N-23; N-23 berikutnya dari A dilepas paling cepat 10 menit sesudah yang terakhir dilepas, berbunyi umum.
- **AC-2** — *Given* B sedang dalam jam tenang, *When* A mengirim lima pesan, *Then* yang tertahan untuk B dari A hanya satu, dilepas saat jam tenang B berakhir.
- **AC-3** — *Given* notifikasi N-23, *When* tampil di layar kunci B, *Then* bunyinya teks umum kelas `kabar` yang sama dengan `kabar` lain ("Ada kabar baru"), tanpa nama pengirim dan tanpa isi, juga bila B memilih rincian untuk notifikasi lain, dan kanal atau kategori notifikasinya sama dengan `kabar` lain; payload push tidak memuat teks, nama, atau jumlah (S-11 f; 06 §11.1, §16.8). Bentuk ini (a) diterima lewat "lanjut" di checkpoint S3 (S3-8 (vii)); cabang alternatif (b) tidak berlaku.
- **AC-4** — *Given* A mengajukan permintaan (B01), *When* kartunya lahir di chat, *Then* B menerima N-01 saja, bukan N-01 dan N-23; baris peristiwa atas transisi berikutnya tidak memicu N-23 (INV-A46).
- **AC-5** — *Given* B membalas kartu permintaan A, *When* balasan tersimpan, *Then* A menerima N-23, bukan N-19.
- **AC-6** — *Given* A mengirim banyak pesan, *When* plafon harian dan jendela 6 jam dihitung untuk A sebagai pengirim (FR-NTF-011, FR-NTF-013), *Then* N-23 tidak ikut terhitung dan tidak menunda pemberitahuan `perlu_tanggapan` mana pun.
- **AC-7** — *Given* B membuka notifikasi N-23, *When* B belum membuka layar chat, *Then* tanda dibaca tidak bergerak (FR-COM-011/AC-3).

**Dua arah.** Simetris.

### FR-COM-018 — Offline dan urutan pesan

| | |
|---|---|
| Fitur | F-COM-001, F-COM-009 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5`; brief §1 "reliable on poor networks" |
| Aktor dan izin | Peserta mana pun |
| Aturan acuan | 06 §12 S-11 (a) sampai (f); S-1, S-2, S-5, S-6; INV-A47 (d). Transport, antrean lokal, dan penyimpanan: dokumen 09 (S4) |
| Business rule | BR-COM-016 |
| Kedudukan | **[Usulan]** `architect` |

**Kriteria penerimaan**
- **AC-1** — *Given* A tanpa sambungan, *When* A mengirim pesan, *Then* pesan tampil di perangkat A sebagai "menunggu terkirim" dan terkirim saat tersambung; kiriman ulang yang sama tampil satu kali bagi keduanya (`idempotency_key`).
- **AC-2** — *Given* A dan B mengirim pesan pada detik yang sama dari perangkat dengan jam yang berbeda, *When* keduanya membuka chat, *Then* urutannya sama di kedua perangkat menurut `seq` server, dan `seq` naik tanpa celah.
- **AC-3** — *Given* A tanpa sambungan, *When* A mengetik, *Then* tanda "sedang mengetik" tidak pernah dikirim belakangan; "terakhir aktif" A diisi server dari akhir sesi pemakaian di depan dengan interaksi terakhir yang diketahui server, bukan dari sambungan latar mana pun (FR-COM-013/AC-7).
- **AC-4** — *Given* hanya push disusul tarik yang tersedia, *When* semua AC berkas ini dijalankan, *Then* semuanya tetap lulus; yang boleh berbeda hanya kecepatan tampil (S-11 e).

**Dua arah.** Simetris.

---

## F-COM-008 — Kartu objek di chat

### FR-COM-004 — Kartu yang lahir otomatis

| | |
|---|---|
| Fitur | F-COM-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `KS-4` (permintaan tiba di aplikasi, bukan sebagai kejutan). Contoh pemilik produk: "ada notif, saya mau naik jatah, yaudah bisa terima/diskusi" (`_progress.md`) |
| Aktor dan izin | Server, dalam transaksi perintah objeknya; pelaku perintah adalah pelaku objek |
| Aturan acuan | 06 §16.2 "Kartu otomatis", "Yang dapat menjadi kartu"; §13.5; INV-A46 |
| Business rule | BR-COM-002, BR-COM-003, BR-COM-008 |
| Kedudukan | Permintaan tampil sebagai kartu di chat: diterima atas saran lead. Daftar pemicu **[Usulan]** `architect` |

**Kriteria penerimaan**
- **AC-1** — *Given* A dan B, *When* A mengajukan permintaan anggaran Rp100.000 (B01), *Then* dalam transaksi yang sama sebuah kartu lahir di chat pada `seq` berikutnya, terlihat oleh keduanya, dan B menerima N-01 saja.
- **AC-2** — *Given* A mengusulkan kenaikan jatah pribadinya (K01), *When* usulan tersimpan, *Then* kartu lahir di chat dan B menerima N-11 saja.
- **AC-3** — *Given* A mencatat serah terima ke tempat uang yang dipegang B (C01), *When* catatan tersimpan dengan tanda belum dikonfirmasi, *Then* kartu lahir dan B menerima N-13 saja.
- **AC-4** — *Given* `REQ` Tier 2 dibangun, *When* A meminta B mengerjakan sesuatu (T01), *Then* kartu lahir dan notifikasinya mengikuti 06 §13.1.
- **AC-5** — *Given* A mencatat pengeluaran dari layar catat (bukan dari chat), *When* entri tersimpan, *Then* tidak ada kartu yang lahir; entri tampil seperti biasa di layarnya.
- **AC-6** — *Given* A membuat empat tagihan berulang dengan usulan pos rutinnya pada satu hari, *When* B membuka chat, *Then* di server ada satu kartu per usulan (tampilan kelompoknya milik dokumen 08), dan B menerima paling banyak satu N-11 langsung dari A hari itu (SP-36).
- **AC-7** — *Given* perintah objek ditolak (misalnya B01 dengan jumlah Rp0), *When* server menolak, *Then* tidak ada kartu dan tidak ada apa pun di chat.

**Dua arah.** Simetris; kartu lahir sama apa pun pengirimnya.

### FR-COM-005 — Membuat atau mengirim objek dari chat

| | |
|---|---|
| Fitur | F-COM-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `MSL-5` (mencatat di tempat percakapan terjadi) |
| Aktor dan izin | Peserta mana pun, dengan hak yang sama seperti dari layar objeknya |
| Aturan acuan | 06 §16.2 "Kartu otomatis" (kalimat entri), "Yang dapat menjadi kartu", "Kirim ke chat" (`share_to_chat`); §16.8 "Kartu yang dikirim ke chat"; N-23; INV-A44; INV-A46; INV-A47 (c); S-7; §13.5 |
| Business rule | BR-COM-003, BR-COM-004, BR-COM-007 |
| Kedudukan | Kata pemilik produk: "chat itu ada/diskusi di chat bisa kirim task/pengeluaran/rencana, ada tanda di baca/delivery chat itu wajib loh" (pernyataan pemilik produk sendiri, `_progress.md`). Entri menjadi kartu bila dicatat atau dikirim dari chat: **[Usulan]** `architect`. "Kirim ke chat" (`share_to_chat`) untuk setiap objek yang boleh menjadi kartu, dan notifikasinya: **[Usulan]** `architect`, 06 §16.2 "Kirim ke chat" dan §16.8 (revisi kesepuluh; menutup PA-COM-1 dan PA-COM-2, mengikuti bawaan berkas ini) |

Dari chat, peserta dapat memulai permintaan anggaran, permintaan jalan-jalan ("rencana"), pengeluaran atau entri buku rumah tangga lain, dan (Tier 2) permintaan tindakan ("tugas"). Formulir dan perintahnya **sama** dengan dari layar objek; chat tidak punya jalan pintas.

**Kirim ke chat** (06 §16.2, dalam rumusan ringkas `architect`; bunyi lengkapnya di 06): "Setiap objek yang boleh menjadi kartu dan belum punya kartu di percakapan yang terbuka dapat dikirim ke chat (`share_to_chat`) oleh peserta yang melihatnya, bila objek itu terlihat oleh semua peserta. Tetap satu kartu per objek per percakapan; bila kartunya sudah ada, kartu itu yang dibuka. Status kepedulian tidak dapat dikirim dengan cara ini." Termasuk keberatan, kesepakatan yang sudah berlaku, masalah rumah (Tier 2), dan objek yang dibuat sebelum percakapan itu ada. Pengirim boleh menyertakan teks, yang menjadi pesan yang menempel pada kartu itu.

**Notifikasinya** (06 §16.8, dalam rumusan ringkas `architect`): "Mengirim ke chat, lewat `share_to_chat` atau dengan mencatat entri dari chat, memicu N-23 yang ikut lipat per pengirim, kecuali perintah yang sama sudah melepas notifikasi lain kepada penerima itu (misalnya N-10, N-13, N-14). Bila notifikasi itu `ringkasan`, N-23 tetap tidak dikirim." Teks yang disertakan tidak menambah notifikasi kedua.

**Kriteria penerimaan**
- **AC-1** — *Given* A di chat, *When* A memulai permintaan jalan-jalan dari chat dan mengirimnya, *Then* yang dijalankan adalah B01 dengan guard yang sama seperti dari layar permintaan, kartunya lahir (FR-COM-004/AC-1), dan tidak ada medan atau akibat yang berbeda (INV-A44).
- **AC-2** — *Given* A di chat, *When* A mencatat pengeluaran Rp75.000 dari chat, *Then* entri tersimpan lewat perintah yang sama dengan F-FIN-005 (validasi sama, tanda dinilai server, FR-REQ-032) dan kartu entri lahir di chat.
- **AC-3** — *Given* entri buku rumah tangga B tanpa kartu, *When* A mengirimnya ke chat, *Then* satu kartu lahir; *Given* entri itu sudah punya kartu, *When* siapa pun mengirimnya lagi, *Then* tidak ada kartu kedua.
- **AC-4** — *Given* entri di buku pribadi A atau hitungan dompet campuran, *When* A mencoba mencatatnya dari chat atau mengirimnya ke chat, *Then* pilihan itu tidak tersedia dan perintahnya ditolak (INV-A47 c); pencatatan buku pribadi tetap tersedia di layarnya sendiri.
- **AC-5** — *Given* A tanpa sambungan, *When* A mencatat pengeluaran dari chat, *Then* kartu tampil "menunggu terkirim" di perangkat A, entri tersimpan menurut S-7 saat tersambung, dan kartunya mendapat `seq` server saat diterima.
- **AC-6** — *Given* keberatan B atas entri A, kesepakatan pos rutin "internet" yang sudah berlaku, dan entri yang dicatat sebelum percakapan dibuat, semuanya terlihat oleh A dan B dan belum punya kartu, *When* A mengirim masing-masing ke chat (`share_to_chat`) dengan teks "yang ini kita bahas ya", *Then* tepat satu kartu lahir per objek, teks A menjadi pesan yang menempel pada kartu itu, dan B menerima tepat satu N-23 per kiriman (dilipat per pengirim), bukan dua; *When* A atau B mengirim objek yang sama lagi, atau perintah yang sama diulang dengan `idempotency_key` yang sama, *Then* tidak ada kartu baru dan tidak ada notifikasi, dan aplikasi membuka kartu yang sudah ada (06 §16.2 "Kirim ke chat", §16.8; menutup PA-COM-1).
- **AC-7** — *Given* ambang Rp500.000, *When* A mencatat dari chat pengeluaran Rp75.000 dari tempat uang yang A pegang, *Then* B menerima N-23 saja. *When* A mencatat dari chat pengeluaran Rp600.000 tanpa permintaan, *Then* entri ditandai (600.000 > 500.000) dan B menerima N-10 saja, tanpa N-23. *When* A mencatat dari chat setoran ke tempat uang yang dipegang B, *Then* B menerima N-13 saja, tanpa N-23. *When* entri yang dicatat dari chat melepas notifikasi berprioritas `ringkasan` kepada B (N-10 dasar `personal_share` atau `routine_amount`, 06 §6.2), *Then* N-23 juga tidak dikirim, sehingga mengirim ke chat tidak menaikkan prioritas kabar tanda (06 §16.8, §6.2, INV-A46; D-35; menutup PA-COM-2).
- **AC-8** — *Given* status suasana hati A yang dibagikan, *When* siapa pun mengirim perintah `share_to_chat` atasnya, *Then* ditolak; kartu status hanya lahir dari berbagi (06 §13.5; FR-COM-016). *Given* objek yang tidak terlihat oleh semua peserta (misalnya entri buku pribadi A, V-9), *When* A mengirim `share_to_chat` atasnya, *Then* ditolak tanpa jejak yang terbaca B (V-7; INV-A47 (c); FR-COM-010/AC-1).

**Dua arah.** Simetris.

### FR-COM-006 — Isi kartu dan baris peristiwa

| | |
|---|---|
| Fitur | F-COM-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1`, `MSL-4` (keadaan terbaru terlihat tanpa bertanya) |
| Aktor dan izin | Isi kartu mengikuti visibilitas objek bagi tiap pembaca (06 §9; V-7) |
| Aturan acuan | 06 §16.2 "Isi kartu", "Baris peristiwa"; §16.7 "Kartu dan baris peristiwa"; INV-A46, INV-A47 (b) |
| Business rule | BR-COM-002, BR-COM-008 |
| Kedudukan | **[Usulan]** `architect` |

**Kriteria penerimaan**
- **AC-1** — *Given* kartu permintaan A di chat, *When* B mengajukan usul balik Rp80.000 dari layar permintaan (B03), *Then* kartu bagi keduanya seketika menampilkan syarat dan state terbaru, satu baris peristiwa "usul balik" berpelaku B tampil di chat, dan A menerima N-02 saja.
- **AC-2** — *Given* kartu permintaan yang disetujui, *When* permintaan diselesaikan atau ditutup otomatis, *Then* kartu menampilkan state terbaru dan tiap transisi punya baris peristiwanya.
- **AC-3** — *Given* kartu atas objek yang kemudian tidak lagi terlihat oleh B (V-7), *When* B membuka chat, *Then* kartu itu tampil "tidak lagi tersedia" tanpa isi bagi B.
- **AC-4** — *Given* kartu yang memuat jumlah, *When* A dan B membandingkannya dengan layar objek, *Then* angkanya sama persis dengan layar objek bagi masing-masing; kartu tidak punya angka sendiri.
- **AC-5** — *Given* kartu atau baris peristiwa, *When* siapa pun mengirim perintah menghapusnya, *Then* server menolak (FR-COM-015/AC-4).
- **AC-6** — *Given* A mengusulkan empat kesepakatan pos rutin pada hari yang sama, *When* B membuka chat, *Then* keempatnya tampil sebagai satu kelompok kartu ("4 usulan tagihan rutin dari A"); membuka kelompok menampilkan keempat kartu, masing-masing dengan tombolnya sendiri; tidak ada tombol "Setujui semua" di kelompok itu; di server tetap satu kartu per objek dan setiap tanggapan tetap perintah K02 atau K03 atas satu usulan (**[Usulan]** `ux-designer`, 08 §10.4; BR-COM-002, BR-COM-004). Persetujuan banyak sekaligus dengan syarat yang dibaca satu per satu tetap hanya di layar pos rutin (FR-FIN-102). Usaha: mengurangi gulir, tidak menambah ketukan.

**Dua arah.** Simetris.

### FR-COM-010 — Yang tidak pernah menjadi kartu atau balasan

| | |
|---|---|
| Fitur | F-COM-008, F-COM-010 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` (dana pribadi tertutup, keputusan pemilik nomor 3; berbagi kepedulian pilihan pemiliknya, brief §1.1 "Guardrail") |
| Aktor dan izin | Ditolak server, siapa pun pengirimnya |
| Aturan acuan | 06 §16.2 "Yang tidak pernah menjadi kartu atau balasan"; INV-A47 (c); V-9; §13.3 I02; INV-A33 |
| Business rule | BR-COM-007, BR-COM-012 |
| Kedudukan | **[Usulan]** `architect` |

**Kriteria penerimaan**
- **AC-1** — *Given* tiap jenis objek terlarang (entri buku pribadi, isi dompet campuran, status atau laporan sakit yang tidak dibagikan, pertanyaan tindak lanjut sistem kepada yang sakit, setelan jam tenang, tunda), *When* siapa pun mencoba membuat kartu atasnya atau pesan ber-`subject_ref`-nya, *Then* server menolak dan tidak ada jejak yang terbaca anggota lain.
- **AC-2** — *Given* A menunda permintaan B, *When* B membaca chat, *Then* tidak ada apa pun di chat yang menyatakan tunda itu atau lamanya; satu-satunya yang tampil adalah tanda dibaca bila A sendiri sudah membaca kartunya (AT-5 d; FR-COM-011/AC-7).
- **AC-3** — *Given* status kepedulian A yang dibagikan, *When* B membaca chat, *Then* yang terlihat hanya kartu hidup status saat ini (FR-COM-016), tidak pernah status sebelumnya.

**Dua arah.** Simetris.

---

## F-COM-009 — Menanggapi dari kartu

### FR-COM-007 — Menanggapi dari kartu

| | |
|---|---|
| Fitur | F-COM-009 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` (menyetujui, berunding, atau usul balik dalam satu atau dua ketukan) |
| Aktor dan izin | Hanya pihak yang berhak atas perintah itu menurut modul objeknya. Pihak yang tidak berhak tidak melihat tombolnya |
| Aturan acuan | 06 §16.4 (tabel pemetaan); S-3, S-4; INV-A44; PG-02; INV-A04, INV-A19; P-25 |
| Business rule | BR-COM-004, BR-COM-005, BR-COM-017 |
| Kedudukan | Menanggapi dari kartu: diterima atas saran lead. Pemetaan, termasuk "usul lain" yang K01-nya ditolak guard (revisi kesepuluh, CA-05): **[Usulan]** `architect`. Arti "belum sepakat" = perintah menolak yang sudah ada: usulan `architect` (06 P-25), diterima lewat "lanjut" di checkpoint S3, S3-7 (i) |

Pemetaan tombol mengikuti 06 §16.4 tanpa diubah; kriterianya diuji di FR objeknya.

| Kartu | Tanggapan dari kartu | FR objeknya |
|---|---|---|
| Permintaan menunggu pemberi persetujuan | Terima B06; usul lain B03; belum sepakat B08. Pemohon: revisi B02, batal B09, ingatkan lagi | FR-REQ-007 sampai FR-REQ-009, FR-REQ-021, FR-REQ-024 |
| Permintaan menunggu pemohon | Terima B07; usul lain B04; batal B09 | FR-REQ-009 |
| Perubahan syarat menunggu | Terima B12; belum sepakat B13 | FR-REQ-012 |
| Konfirmasi susulan | Terima B06; belum sepakat B08; tanpa usul lain | FR-REQ-038 |
| Usulan kesepakatan | Terima K02; usul lain = K03 lalu K01 baru oleh penanggap; belum sepakat K03. Pengusul: tarik K04 | FR-REQ-026, FR-REQ-027 |
| Serah terima menunggu konfirmasi | Terima C03; keberatan C04 (bukan penolakan) | FR-FIN di [FIN-1.md](FIN-1.md) untuk F-FIN-006 |
| Permintaan tindakan (Tier 2) | Terima T02; tolak T03. Penerima: T05, T06 | FR-REQ-044, FR-REQ-045 |
| Entri | Keberatan D01; koreksi §8.1, menurut haknya | FR-FIN di [FIN-2.md](FIN-2.md) untuk F-FIN-018 |

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A Rp100.000 versi 1 menunggu B, *When* B mengetuk "terima" pada kartunya, *Then* server menjalankan B06 atas versi 1 dengan state yang diharapkan `awaiting_approver`, hasilnya identik dengan dari layar permintaan, baris peristiwa tampil, dan A menerima N-04.
- **AC-2** — *Given* AC-1 sebelum B menanggapi, *When* B memilih "usul lain" dengan Rp80.000, *Then* B03 dijalankan, kartu menampilkan "menunggu A", dan A menerima N-02.
- **AC-3** — *Given* AC-1 sebelum B menanggapi, *When* B memilih "belum sepakat" tanpa alasan, *Then* B08 dijalankan, permintaan `rejected`, dan A dapat mengajukannya lagi dalam rantai yang sama (FR-REQ-023).
- **AC-4** — *Given* AC-1 sebelum B menanggapi, *When* B hanya membalas kartu "nanti malam kita bahas ya", *Then* permintaan tetap `awaiting_approver`, pesan menempel pada permintaan (FR-COM-009), dan A menerima N-23 saja.
- **AC-5** — *Given* kartu permintaan A sendiri, *When* A membukanya, *Then* A tidak melihat tombol terima, usul lain, atau belum sepakat; A melihat revisi, batal, dan ingatkan lagi (dalam batas PG-06); perintah `approve` dari perangkat A yang tetap terkirim ditolak (INV-A04).
- **AC-6** — *Given* kartu serah terima yang dicatat A menunggu konfirmasi B, *When* B memilih "keberatan", *Then* yang dijalankan C04, bukan penolakan, dan catatannya tetap ada dengan tanda dikeberatani.
- **AC-7** — *Given* usulan A menaikkan jatah pribadinya, *When* B memilih "usul lain" dengan jumlah yang tidak di atas ambang, *Then* sebelum mengirim layar B menampilkan batas yang dibaca dari server (ambang yang berlaku dan siapa yang boleh mengusulkan jatah di bawahnya) tanpa menghitung guard sendiri (S-8); server menjalankan K03 lalu K01; K01 baru dari B ditolak guard K01 (INV-A19); K03 tetap berlaku, usulan A berakhir `ditolak` dan tidak hidup kembali, dan tidak ada usulan dari B; A menerima N-12 atas K03 seperti biasa; B diberi tahu hasilnya di layar (atau di hasil sinkron bila offline) dan boleh mengirim K01 lain yang sah atau tidak sama sekali (06 §16.4, revisi kesepuluh; PA-COM-4 ditutup; angka di FR-REQ-027/AC-8).
- **AC-8** — *Given* rangkaian tanggapan yang sama dikirim sekali dari kartu dan sekali dari layar objek pada dua salinan keadaan yang sama, *When* hasil server dibandingkan, *Then* rangkaian perintah, Peristiwa, dan akibat uangnya identik (INV-A44).

**Dua arah.** Simetris: setiap AC berlaku dengan A dan B ditukar; tombol ditentukan peran dalam objek (pemohon, pemberi persetujuan, pengusul, penerima), bukan orangnya.

### FR-COM-008 — Kartu yang basi dan tanggapan bersamaan

| | |
|---|---|
| Fitur | F-COM-009 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` (persetujuan tidak pernah pindah ke syarat yang belum dibaca pemberinya) |
| Aktor dan izin | Seperti FR-COM-007 |
| Aturan acuan | 06 S-1 sampai S-4; S-11 (b); diagram urutan §16.4 |
| Business rule | BR-COM-004, BR-COM-016 |
| Kedudukan | **[Usulan]** `architect` |

**Kriteria penerimaan**
- **AC-1** — *Given* perangkat B menampilkan kartu permintaan A versi 1, *When* A merevisi ke versi 2 (B02) sebelum B mengetuk "terima", *Then* B06 atas versi 1 ditolak `terms_superseded`, kartu menampilkan versi 2, dan tidak ada yang disetujui.
- **AC-2** — *Given* A membatalkan permintaannya (B09), *When* B mengetuk "usul lain" pada kartu yang belum diperbarui, *Then* ditolak `state_changed` dan kartu menampilkan "dibatalkan".
- **AC-3** — *Given* B tanpa sambungan, *When* B mengetuk "terima", *Then* kartu di perangkat B tampil "menunggu terkirim", tidak pernah "disetujui"; saat tersambung, perintah diterima hanya bila versi dan state masih sama.
- **AC-4** — *Given* B mengetuk "terima" dua kali, *When* keduanya sampai di server, *Then* hanya satu tanggapan yang tercatat (`idempotency_key`).

**Dua arah.** Simetris.

---

## F-COM-010 — Balasan yang menempel pada kartu

### FR-COM-009 — Balasan yang menempel pada kartu

| | |
|---|---|
| Fitur | F-COM-010 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` (diskusi menempel pada permintaannya, brief §5B "contextual conversations"); arah pemilik produk "semua ada catatanya ada datanya" |
| Aktor dan izin | Peserta mana pun yang melihat objeknya. `link_message` hanya oleh penulis atas pesannya sendiri |
| Aturan acuan | 06 §16.3; §4.4; N-19, N-23; INV-A44 |
| Business rule | BR-COM-005, BR-COM-006, BR-COM-007 |
| Kedudukan | Menempel: diterima atas saran lead ("Reply nempel ke kartu"). `link_message` dan batas offline **[Usulan]** `architect`. Pesan di layar objek atas objek tanpa kartu: **[Usulan]** `architect`, 06 §16.3 "Pesan dari layar objek tanpa kartu" (revisi kesepuluh; menutup PA-COM-3 dan PA-REQ-6, mengikuti bawaan berkas ini) |

**Kriteria penerimaan**
- **AC-1** — *Given* kartu permintaan A, *When* B membalas kartu itu, *Then* pesan mendapat `subject_ref` permintaan, tampil di chat pada urutannya dan di diskusi permintaan di layarnya dengan isi yang sama, dan A menerima N-23, bukan N-19.
- **AC-2** — *Given* pesan B yang menempel pada permintaan, *When* A membalas pesan B itu, *Then* balasan A juga menempel pada permintaan yang sama.
- **AC-3** — *Given* pesan bebas A, *When* A menautkannya ke kartu permintaan (`link_message`), *Then* pesan itu menempel, Peristiwa tautan tercatat pada permintaan, dan tautan itu tidak dapat dilepas; untuk membatalkannya A menghapus pesannya (FR-COM-015). *When* B mencoba menautkan pesan A, *Then* ditolak.
- **AC-4** — *Given* pesan yang menempel berbunyi "oke setuju", *When* tersimpan, *Then* state, versi syarat, dan otorisasi permintaan tidak berubah (INV-A44).
- **AC-5** — *Given* B tanpa sambungan membalas kartu objek yang sementara itu tidak lagi terlihat oleh semua peserta (V-7), *When* tersinkron, *Then* balasan tersimpan sebagai pesan bebas dan B diberi tahu hasil sinkronnya.
- **AC-6** — *Given* permintaan A punya kartu, *When* B menulis pesan di diskusi permintaan dari layar permintaan, *Then* pesan yang sama tampil juga di chat sebagai pesan yang menempel pada kartu itu.
- **AC-7** — *Given* entri buku rumah tangga A yang dicatat dari layar catat, terlihat oleh A dan B, dan belum punya kartu, *When* B menulis pesan di diskusi entri itu dari layar entri, *Then* sesuai 06 §16.3 (rumusan ringkas `architect`): "Pesan yang ditulis di layar objek atas objek yang terlihat oleh semua peserta tetapi belum punya kartu tampil di chat sebagai pesan yang menempel, dengan rujukan ringkas yang dibaca dari state saat ini. Ia tidak melahirkan kartu, dan notifikasinya N-23 saja, bukan N-19." Rujukan ringkas itu menampilkan "tidak lagi tersedia" bila entri kemudian menyempit (V-7); A yang membalas pesan itu juga menempel pada entri yang sama; *Given* objek yang tidak terlihat oleh semua peserta, *Then* pesan tetap di diskusi objek dan notifikasinya N-19 (menutup PA-COM-3; FR-REQ-041/AC-8).

**Dua arah.** Simetris.

---

## F-COM-011 — Tanda terkirim dan dibaca

### FR-COM-011 — Tanda terkirim dan dibaca

| | |
|---|---|
| Fitur | F-COM-011 |
| Tier / MoSCoW | 1 / Must; **Pantau** |
| Masalah (brief §1.1) | `MSL-4` (pemohon tahu permintaannya sudah sampai tanpa bertanya). Tanda pada pesan chat wajib menurut pernyataan pemilik produk sendiri ("ada tanda di baca/delivery chat itu wajib loh") |
| Aktor dan izin | Setiap peserta melihat tanda atas pesan dan kartu di chat, sama bagi keduanya |
| Aturan acuan | 06 §16.5; AT-5 (a), (b), (d); INV-A33, INV-A45; S-11 (c); ADR-011 butir 5, 7 |
| Business rule | BR-COM-009, BR-COM-012 |
| Kedudukan | Tanda pada pesan chat: pernyataan pemilik produk sendiri; pada kartu: diterima atas saran lead ("Ikut tanda dibaca"). Terkirim hanya oleh penarikan saat aplikasi dibuka, satu aturan di semua platform: usulan `security-privacy` SP-55, diterapkan 06 revisi kesepuluh, diterima lewat "lanjut" sebagai D-93; dapat-tidaknya dibedakan dari ekstensi dan tugas latar di kedua platform **[Asumsi]** yang diverifikasi dokumen 09. Tanda air dan tanpa waktu baca **[Usulan]** `architect` |

**Akibat yang dinyatakan terus terang** (06 §16.5). Karena terkirim hanya naik saat aplikasi penerima dibuka, terkirim dan dibaca sering bergerak hampir bersamaan. Tanda terkirim berarti "ponselnya sudah membuka aplikasi sejak pesan ini dikirim", bukan "pesan sudah sampai di ponsel".

**Kriteria penerimaan**
- **AC-1** — *Given* A mengirim pesan ber-`seq` n dan aplikasi B tertutup, *When* B membuka aplikasinya di depan dan perangkat B menarik pesan n dari server, menyimpannya, dan mengakuinya, *Then* baru saat itu `delivered_upto_seq` B naik dan A melihat "terkirim" pada n dan pada semua pesan A sebelum n.
- **AC-2** — *Given* B membuka chat dengan aplikasi di depan dan pesan n tampil di layarnya, *When* `read_upto_seq` B naik ke n, *Then* A melihat "dibaca" pada n dan pada semua pesan dan kartu sebelum n.
- **AC-3** — *Given* B menerima N-01 atas permintaan A, *When* B membuka notifikasi itu atau membuka layar permintaan tanpa membuka layar chat, *Then* tanda pada kartu permintaan tidak bergerak.
- **AC-4** — *Given* kartu permintaan A ber-`seq` n dan pesan A ber-`seq` n+1, *When* B membaca n+1, *Then* kartu n juga tampil dibaca bagi A (tanda air).
- **AC-5** — *Given* tanda apa pun, *When* A atau B memeriksa pesan, info pesan, ekspor, atau skema, *Then* tidak ada waktu baca per pesan yang tampil atau tersimpan, dan waktu pembaruan tanda air tidak terbaca siapa pun.
- **AC-6** — *Given* aplikasi B tertutup, sekali di dalam jam tenang B dan sekali di luarnya, dan pada kedua keadaan B membuka aplikasinya satu jam sesudah pesan dikirim, *When* A mengirim pesan yang sama, *Then* tanda terkirim yang dilihat A identik pada kedua keadaan dan baru bergerak saat B membuka aplikasinya, karena push yang dilepas di luar jam tenang tidak menggerakkannya (INV-A33; 06 §16.5).
- **AC-7** — *Given* B menunda permintaan A, *When* B membaca kartunya di chat, *Then* A melihat kartu itu dibaca; A tidak melihat bahwa B menunda atau sampai kapan (AT-5 d, dinyatakan terus terang).
- **AC-8** — *Given* perangkat B tanpa sambungan membaca pesan, *When* tersambung, *Then* tanda air terakhir terkirim tanpa waktu dari jam perangkat, dan tanda air tidak pernah turun, juga sesudah pesan dihapus.
- **AC-9** — *Given* layar objek, notifikasi, ringkasan notifikasi, atau ekspor, *When* dibuka, *Then* tidak ada tanda terkirim atau dibaca di sana (AT-5 b).
- **AC-10** — *Given* aplikasi B tertutup dan perangkat B menerima N-23 atas pesan n, *When* push itu diproses, Notification Service Extension (iOS) atau penerima pesan latar (Android) mengambil datanya, atau sinkron latar berjalan, *Then* `delivered_upto_seq` B tidak bergerak pada salah satu pun dari keempatnya, di kedua platform, walau datanya sudah ada di perangkat B; tanda terkirim baru bergerak menurut AC-1 (06 §16.5; SP-55). Bagian iOS (termasuk iPadOS) lulus sebelum GR-2; bagian Android lulus sebelum klien Android memegang data nyata, paling lambat sebelum GR-4 (17 keputusan 42; 13 §3.4).

**Dua arah.** Simetris: B melihat tanda pesan B dengan cara yang sama persis.

### FR-COM-012 — Tanpa tekanan dari tanda dibaca

| | |
|---|---|
| Fitur | F-COM-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; arah pemilik produk kedua ("bukan dipakai untuk menyuruh nyuruh") |
| Aktor dan izin | Tidak ada pelaku: larangan atas sistem |
| Aturan acuan | 06 §16.5 "Yang tidak ada", "Diam"; §16.9; AT-4; PG-04, PG-07; INV-A45 |
| Business rule | BR-COM-010, BR-COM-017 |
| Kedudukan | Tanpa "dibaca tapi belum dijawab" dan tanpa hitungan lama menjawab: diterima atas saran lead. Selebihnya **[Usulan]** `architect` |

**Kriteria penerimaan**
- **AC-1** — *Given* B membaca kartu permintaan A tiga hari lalu dan belum menanggapi, *When* jam dimajukan, *Then* tidak ada pengingat, lencana, warna, urutan, atau notifikasi yang berasal dari "dibaca tapi belum dijawab"; yang berlaku hanya batas permintaan itu sendiri (kedaluwarsa B10, pengingat 06 §10 dalam batasnya).
- **AC-2** — *Given* layar, ringkasan, laporan, ekspor, atau keluaran `AI` mana pun, *When* diperiksa, *Then* tidak ada hitungan pesan per orang, hitungan pesan yang belum dibaca pasangan, rata-rata atau tren lama membalas, jarak antara dibaca dan dijawab, atau statistik chat apa pun.
- **AC-3** — *Given* A punya pesan yang belum ia baca, *When* lencana tampil, *Then* lencana itu hanya menghitung pesan yang belum dibaca A dan hanya terlihat oleh A.
- **AC-4** — *Given* B membaca kartu lalu diam, *When* waktu berjalan, *Then* permintaan tetap menunggu; membaca tidak pernah dicatat sebagai tanggapan (PG-04).
- **AC-5** — *Given* dua jalankan rangkaian perintah yang sama dengan `read_state` diacak di antaranya, *When* semua keluaran di luar layar chat dibandingkan (jadwal pengingat, tanda, ringkasan, laporan, ekspor, masukan ke penyedia `AI`), *Then* semuanya identik byte demi byte (INV-A45).

**Dua arah.** Simetris.

---

## F-COM-012 — Kehadiran

### FR-COM-013 — Sedang online, sedang mengetik, terakhir aktif

| | |
|---|---|
| Fitur | F-COM-012 |
| Tier / MoSCoW | 1 / Must; **Pantau** |
| Masalah (brief §1.1) | Tidak menunjuk `MSL` tersendiri. Ada atas **keputusan pemilik produk sendiri, berlawanan dengan saran lead**; dasar terdekat di brief adalah §1 "Experience bar: WhatsApp-class". Dicatat sebagai pengecualian atas aturan "fitur tanpa masalah §1.1 turun ke Tier 3" (laporan untuk dokumen 17) |
| Aktor dan izin | Setiap peserta melihat kehadiran peserta lain, hanya di layar chat |
| Aturan acuan | 06 §16.6; AT-5 (a), (d); S-11 (d); INV-A33 |
| Business rule | BR-COM-011, BR-COM-012 |
| Kedudukan | Tampil: keputusan pemilik produk sendiri. Hanya oleh pemakaian di depan dengan interaksi: usulan `security-privacy` SP-60, diterapkan 06 revisi kesepuluh, diterima lewat "lanjut" sebagai D-95. Batas tampil dan cara simpan **[Usulan]** `architect`. Lama tanpa interaksi yang mengakhiri "online", lama "sedang mengetik", dan ketelitian waktu "terakhir aktif" **[Dok. 09]**, **[Dok. 08]**. Opsi P1 (setelan timbal balik per orang) dan P2 (kehadiran hanya di layar chat yang sedang dibuka) hanya disajikan kepada pemilik produk di S4 (06 P-28) dan **tidak** diterapkan tanpa persetujuannya |

**Kriteria penerimaan**
- **AC-1** — *Given* aplikasi B di depan, tersambung, dan B sedang menyentuh, menggulir, atau mengetik, *When* A membuka layar chat, *Then* A melihat B sedang online; *When* B menutup aplikasi, *Then* tanda online hilang dan "terakhir aktif" B berisi waktu server saat sesi pemakaian di depan dengan interaksi itu berakhir.
- **AC-2** — *Given* B mengetik di chat, *When* A membuka layar chat, *Then* A melihat B sedang mengetik; tanda itu berakhir sendiri dalam beberapa detik sesudah B berhenti.
- **AC-3** — *Given* kehadiran B, *When* A membuka layar depan, layar objek, daftar "yang menunggu", atau notifikasi, *Then* kehadiran B tidak tampil di sana.
- **AC-4** — *Given* B menjadi online, *When* A tidak sedang membuka chat, *Then* A tidak menerima notifikasi apa pun karenanya.
- **AC-5** — *Given* aplikasi B tertutup, sekali di dalam jam tenang B dan sekali di luarnya, *When* A membuka chat, *Then* kehadiran B yang dilihat A identik pada kedua keadaan (INV-A33).
- **AC-6** — *Given* B aktif beberapa kali sehari, *When* A atau siapa pun meminta riwayatnya, *Then* yang ada hanya satu nilai "terakhir aktif" saat ini; nilai sebelumnya tidak dapat diambil.
- **AC-7** — *Given* aplikasi B tidak dipakai di depan, *When* sinkron latar berjalan, push atau ekstensi notifikasi diproses, widget diperbarui, atau aplikasi B terbuka di depan tanpa disentuh, *Then* "sedang online" dan "terakhir aktif" B tidak bergerak; yang dilihat A identik dengan bila tidak satu pun dari itu terjadi (06 §16.6; SP-60).
- **AC-8** — *Given* B anggota dua rumah tangga, *When* B memakai aplikasi di rumah tangga lain, *Then* kehadiran B di rumah tangga ini tidak bergerak (06 §16.2 `presence` per anggota per rumah tangga; `LM-09`).
- **AC-9** — *Given* B mencatat belanja di layar Uang dengan aplikasi di depan, *When* A membuka chat, *Then* A melihat B sedang online, seperti AC-1; kehadiran digerakkan pemakaian di bagian mana pun aplikasi, tidak hanya di layar chat, dan tidak ada setelan menyembunyikan kehadiran (10-A A, S4-5 (i) A, diputuskan pemilik produk sendiri: "A tetap seperti WhatsApp"). Tanda dibaca tidak berubah (10 §3 10-A). Cabang P1 dan P2 gugur; risiko sisanya W-18, diterima pemilik produk (10 §3).
- **AC-10** — *Given* sesi B berakhir pukul 21.07, *When* A membuka chat pukul 21.30, *Then* "terakhir aktif" B berketelitian menit (21.07); kemarin dan tanggal lama tampil dengan jam (UX-5 A, S4-5 (ii) A; 10.2 SEC-60). Pembulatan 15 menit gugur.

**Dua arah.** Simetris: keduanya melihat kehadiran yang lain dengan cara dan batas yang sama.

### FR-COM-014 — Kehadiran tanpa riwayat dan tidak dibaca apa pun

| | |
|---|---|
| Fitur | F-COM-012 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` (brief §8 "surveillance"; 06 PG-14 asas pengawasan) |
| Aktor dan izin | Tidak ada pelaku: larangan atas sistem |
| Aturan acuan | 06 §16.6 "Penyimpanan", "Notifikasi"; §16.7 "Log"; INV-A45; `ENGINEERING-STANDARDS.md` §4 |
| Business rule | BR-COM-011, BR-COM-018 |
| Kedudukan | **[Usulan]** `architect`; penilaian risiko pengawasan, termasuk terhadap jam tenang, status yang dibagikan, tunda, dan IDEA-19, **[Dok. 10]** |

**Kriteria penerimaan**
- **AC-1** — *Given* skema basis data, ekspor, dan peristiwa analitik, *When* diperiksa, *Then* tidak ada tabel, kolom, atau peristiwa yang menyimpan riwayat online, mengetik, atau terakhir aktif.
- **AC-2** — *Given* log server dan klien, *When* diperiksa sesudah rangkaian sesi, *Then* tidak ada catatan kehadiran.
- **AC-3** — *Given* dua jalankan rangkaian perintah yang sama dengan `presence` diacak di antaranya, *When* keluaran aturan, pengingat, ringkasan, notifikasi, laporan, ekspor, dan masukan ke penyedia `AI` dibandingkan, *Then* semuanya identik byte demi byte (INV-A45).
- **AC-4** — *Given* cadangan yang dapat dikueri, *When* dipulihkan, *Then* tidak ada riwayat kehadiran di dalamnya **[Dok. 09]**.

**Dua arah.** Simetris.

---

## F-CARE-003 — Sisi chat dari catatan dukungan (Tier 2)

### FR-COM-016 — Kartu status kepedulian dan catatan dukungan

| | |
|---|---|
| Fitur | F-CARE-003 (memakai F-COM-008 dan F-COM-010) |
| Tier / MoSCoW | 2 / **Must bila F-CARE-001 atau F-CARE-002 dibangun** (pagar isi chat) |
| Masalah (brief §1.1) | `MSL-6` (suasana hati pasangan tidak diketahui kecuali ditanya); `MSL-7`; `KS-5`; `PAGAR` |
| Aktor dan izin | Kartu lahir hanya bila pemilik status sudah memilih berbagi (FR-CARE-002). Membalas: pasangan yang boleh melihat status itu |
| Aturan acuan | 06 §13.5; §13.6 "Mengirim"; §16.2; INV-A47 (e); N-28; P-21; FR-CARE-001 sampai FR-CARE-006 di [tier-2.md](tier-2.md) |
| Business rule | BR-COM-003, BR-COM-014 |
| Kedudukan | Catatan dukungan sebagai balasan chat atas kartu status: diterima atas saran lead, 12 November 2026 (mengubah brief §5O). Kartu hilang dari chat saat status diganti, berakhir, dicabut, atau dihapus: usulan `security-privacy` SP-45, diterapkan 06 revisi kesepuluh (menggantikan "kartu lama tampil tanpa isi"). Satu kartu hidup dan pelepasan balasan: usulan `architect`; pelepasan balasan menyimpang dari "semuanya hilang" (06 P-21). Keduanya diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-7 (ii); harganya W-23 |

**Kriteria penerimaan**
- **AC-1** — *Given* A sudah memilih berbagi kepedulian, *When* A menyimpan status suasana hati, *Then* tepat satu kartu hidup lahir di chat dan B menerima N-28 saja, bukan N-28 dan N-23.
- **AC-2** — *Given* AC-1, *When* A mengganti statusnya, *Then* kartu lama hilang dari chat bagi A dan B tanpa penanda (tanpa kartu kosong, tanpa "status sudah berganti", tanpa baris peristiwa), kartu hidup baru lahir untuk status baru, dan chat tidak memuat isi status lama di mana pun (Q-136; INV-A47 (e)).
- **AC-3** — *Given* kartu status hidup, *When* status berakhir sendiri ("hari ini"), berbaginya dicabut, atau dihapus A, *Then* kartu hilang dari chat bagi B dan bagi A tanpa penanda dan tanpa kartu pengganti, notifikasi yang masih tertahan untuk status itu gugur (V-7), dan tidak ada nomor urut lain yang hilang di chat selain milik kartu itu (INV-A47 (d)).
- **AC-4** — *Given* kartu status A, *When* B membalasnya dengan kalimat yang ia tulis sendiri atau yang ia sunting dari saran `AI`, *Then* balasan menempel pada status itu, tidak ada tanda buatan `AI` di mana pun, dan A menerima N-23.
- **AC-5** — *Given* balasan B menempel pada status A, *When* status itu diganti, berakhir, dicabut, atau dihapus, *Then* tautannya hilang bersama kartunya, tidak ada pesan yang masih ber-`subject_ref` atau `reply_to_message_id` ke kartu itu, dan balasan B tetap di chat sebagai pesan bebas milik B, tanpa kutipan kartu, yang dapat B hapus sendiri. Batas yang dinyatakan: kalimat seperti "semoga cepat sembuh" tetap terbaca sesudah laporan sakitnya dihapus (06 P-21); A diberi tahu hal itu sebelum menghapus (FR-CARE-004/AC-4).
- **AC-6** — *Given* status A yang tidak dibagikan, *When* B atau siapa pun mencoba membalas atau membuat kartunya, *Then* ditolak (FR-COM-010/AC-1).
- **AC-7** — *Given* kartu status, *When* tampil, *Then* tidak memuat alasan bebas kecuali A membagikan alasan itu (FR-CARE-002).
- **AC-8** — *Given* catatan dukungan, *When* jam dimajukan berapa pun, *Then* tidak ada hitungan catatan, rentetan, atau ajakan "belum mengirim dukungan" kepada siapa pun (PRN-14).

**Dua arah.** Simetris.

---

## F-COM-014 — Sembunyikan untuk saya dan laporkan

Diminta dokumen 10 (SEC-153, SEC-154) sebagai padanan Apple Guideline 1.2 tanpa pemindaian dan tanpa mengeluarkan orang dewasa, dan dirancang di 06 §16.12. Pemilik produk memilih 10-G A di checkpoint S4 (S4-7 A, keputusan sendiri), dan F-COM-014 masuk Tier 1 lewat D-114. Tier 1 Must karena 10 §4 baris "Kehadiran, chat, foto" menandai SEC-153 (a) sampai (c), (e) dan SEC-154 (c) "wajib sebelum pilot memproses data nyata"; bagian lain (pengaburan foto SEC-153 (d), pengiriman laporan ke operator, kontak publik, tanggapan 3 × 24 jam, catatan App Review) harus siap sebelum build pertama ke TestFlight eksternal atau App Store, bukan sebelum pilot. Itu jadwal pembangunan, bukan syarat kirim: laporan dikirim hanya bila pengelola bukan anggota rumah tangga pelapor, apa pun jenis build-nya (FR-COM-020/AC-5).

**Pagar pemantauan dan kendali.** Kedua fitur ini tidak memberi siapa pun kendali atas orang lain: menyembunyikan hanya mengubah layar pembacanya, dan laporan tidak menghalangi atau menghukum yang dilaporkan. Pilihan yang memungkinkan satu orang dewasa mengeluarkan orang dewasa lain tidak dipilih (S4-7 A) dan tidak ditulis, sejalan dengan FR-FAM-016, 06 INV-A48, PG-13, dan brief §4.2 "member exit".

**Usaha berulang.** Nol; dipakai hanya bila dibutuhkan.

### FR-COM-019 — Sembunyikan untuk saya

| | |
|---|---|
| Fitur | F-COM-014 |
| Tier / MoSCoW | 1 / Must (S4-7 A, diputuskan pemilik produk sendiri; masuk Tier 1 lewat D-114) |
| Masalah (brief §1.1) | `PAGAR` (brief §8 abuse prevention: pesan yang menyakitkan tanpa jalan selain keluar); syarat toko (10.4 §2.1) |
| Aktor dan izin | Setiap peserta chat, atas tampilannya sendiri; perintah `mute_member(muted_member)` dan `unmute_member` |
| Aturan acuan | 06 §16.12, INV-A53, §16.5; 10.4 SEC-153 |
| Business rule | BR-COM-019 |
| Kedudukan | **[Usulan]** `security-privacy` S4 (SEC-153), dirancang `architect` (06 §16.12). Bentuknya, 10-G A, **diputuskan pemilik produk sendiri** di checkpoint S4 ("S4-7: A") |

**Kriteria penerimaan**
- **AC-1** — *Given* chat rumah tangga, *When* A memilih "Sembunyikan pesan dari B" dan mengukuhkannya, *Then* di layar chat A setiap pesan bebas dan foto dari B, yang lama maupun yang baru, diganti satu baris "Pesan disembunyikan" yang dapat A buka satu per satu; indikator mengetik B tidak tampil bagi A; N-23 dari B tidak dikirim kepada A; *When* A membatalkannya, *Then* semuanya tampil lagi seperti biasa.
- **AC-2** — *Given* AC-1, *When* B mengajukan permintaan, mengusulkan kesepakatan, mencatat belanja di atas ambang, atau membalas pada kartu, *Then* kartu objek, baris peristiwa, balasan yang menempel, dan notifikasi §11.2 (N-01 sampai N-22, N-24 sampai N-32) tetap tampil dan terkirim kepada A; A tetap dapat menanggapi semuanya (06 PG-10).
- **AC-3** — *Given* AC-1, *When* B memakai aplikasi dan membuka chat, *Then* tidak ada yang berubah bagi B: tidak ada Peristiwa, notifikasi, atau tanda yang memberi tahu B; kehadiran A tetap terkirim kepada B seperti biasa; uji 06 INV-A53 lulus (keluaran bagi B identik byte demi byte dengan dan tanpa `mute_member`).
- **AC-4** — *Given* AC-1 dan A menggulir chat melewati baris "Pesan disembunyikan", *When* B melihat pesannya, *Then* tanda dibacanya mengikuti tanda air A seperti biasa, jadi pesan tersembunyi di bawah tanda air tampil "dibaca" bagi B; sebelum mengukuhkan, layar menyembunyikan menyatakan hal ini kepada A ("dibaca" berarti sudah dilewati, bukan sudah dilihat) (06 §16.12 baris "Tanda dibaca"; SEC-153 (b)).
- **AC-5** — *Given* setelan sembunyikan milik A, *When* B membuat ekspor, membaca API, atau keluar dengan ekspornya, *Then* setelan itu tidak ada di mana pun; ia tidak dibaca aturan alur, tanda, ringkasan, laporan, atau `AI` (SEC-153 (e); INV-A53).
- **AC-6** — *Given* build untuk TestFlight eksternal atau App Store, *When* peringatan konten sensitif sistem di ponsel A menyala, *Then* foto dari B dikaburkan di perangkat A sampai A membukanya; hasil pemeriksaan tidak meninggalkan perangkat (SEC-153 (d); tidak wajib di pilot).
- **AC-7** — Gugur sesudah checkpoint S4 (cabang 10-G B atau D tidak dipilih). Nomor tidak dipakai ulang.

**Dua arah.** Simetris; siapa pun dapat menyembunyikan pesan pasangannya dari layarnya sendiri, dan tidak ada yang dapat menyembunyikan pesan seseorang dari orang lain.

### FR-COM-020 — Laporkan

| | |
|---|---|
| Fitur | F-COM-014 |
| Tier / MoSCoW | 1 / Must (S4-7 A, diputuskan pemilik produk sendiri; masuk Tier 1 lewat D-114) |
| Masalah (brief §1.1) | `PAGAR`; syarat toko (10.4 §2.1) |
| Aktor dan izin | Setiap peserta chat; perintah `report_content(subject, attached_excerpt?)` |
| Aturan acuan | 06 §16.12; 10.4 SEC-135, SEC-154; FR-FAM-024 |
| Business rule | BR-COM-020 |
| Kedudukan | **[Usulan]** `security-privacy` S4 (SEC-154), dirancang `architect` (06 §16.12). Penentu jalurnya adalah apakah pengelola anggota rumah tangga pelapor, **bukan** jenis build: selama pengelola anggota rumah tangga pelapor (seluruh pilot, keputusan 25), tombol membuka Bantuan (SEC-154 (c)); laporan dikirim hanya bila pengelola bukan anggota rumah tangga pelapor. Kesiapan jalur kirim itu sebelum build pertama ke TestFlight eksternal adalah jadwal (SEC-154 kolom tahap), bukan syarat kirim |

Di pilot pengelola server adalah anggota rumah tangga, sehingga "lapor ke pengelola" berarti lapor kepada pasangan (10 §3 10-G; 06 §16.12 baris "Pilot"). Karena itu di pilot tombol "Laporkan" membuka halaman Bantuan dan jalan keluar, bukan mengirim laporan.

**Kriteria penerimaan**
- **AC-1** — *Given* pengelola server adalah anggota rumah tangga A (seluruh pilot), *When* A menekan "Laporkan" pada pesan B, *Then* aplikasi membuka halaman Bantuan (FR-FAM-024) dengan keterangan bahwa laporan tidak dikirim karena pengelola juga anggota rumah tangganya (08 §10.7), beserta jalan keluar dari rumah tangga; tidak ada laporan yang terkirim, tidak ada perintah server, baris, Peristiwa, notifikasi, atau log berpengenal akun.
- **AC-2** — *Given* pengelola bukan anggota rumah tangga pelapor, *When* A memilih satu pesan dan menekan "Laporkan" lalu menulis keterangannya, *Then* yang tersimpan di `content_report` hanya pesan yang dipilih A dan teks A; server tidak menambah isi chat lain; A melihat tanda terima di layar; B tidak diberi tahu dan keluaran bagi B tidak berubah (SEC-154 (a), (b); 06 §16.12).
- **AC-3** — *Given* AC-2, *When* operator membaca laporan, *Then* ia membacanya lewat perintah operasi di server (`familyos ops report`), tanpa API admin, dan tidak memperoleh akses baca ke chat karena laporan itu; kontak operator dipublikasikan dan tanggapan mengikuti runbook paling lambat 3 × 24 jam **[Asumsi]** (SEC-154 (d), (e); SEC-33).
- **AC-4** — *Given* isi chat apa pun, *When* server dan perangkat diperiksa, *Then* tidak ada pemindaian isi otomatis (SEC-135); laporan hanya lahir dari tindakan pelapor.
- **AC-5** — **Laporan tidak pernah sampai kepada orang yang dilaporkan.** *Given* akun pengelola adalah anggota rumah tangga A, dengan build apa pun (internal, TestFlight eksternal, atau App Store) dan dengan klien yang dimodifikasi, *When* perintah `report_content` dari A tiba di server, *Then* server menolaknya tanpa menyimpan baris `content_report`, tanpa Peristiwa, notifikasi, atau email kepada siapa pun, dan aplikasi membuka Bantuan seperti AC-1; `familyos ops report` yang dijalankan pengelola itu tidak menampilkan laporan apa pun dari rumah tangganya; *When* keanggotaan pengelola dinilai, *Then* yang dibaca adalah keanggotaan aktif saat perintah tiba, bukan setelan build (06 §16.12 baris "Laporkan"; 10.4 SEC-154 (c); 08 §10.7).

**Dua arah.** Simetris.

---

## Kesenjangan: **[Perlu aturan]**

Pengenal mengikuti skema README (`PA-COM-n`; sebelumnya `K-n`, yang bertabrakan dengan `K-NN` pertentangan sumber). Sesudah 06 revisi kesepuluh (§15.12 "Tindak lanjut") **tidak ada butir yang masih terbuka**; semua jawaban di bawah **[Usulan]** `architect`, mengikuti bawaan berkas ini.

**Yang sudah dijawab.**

| ID | Kesenjangan | FR | Jawaban | Kedudukan |
|---|---|---|---|---|
| PA-COM-1 | Keberatan, kesepakatan yang sudah berlaku, masalah rumah, dan objek yang dibuat sebelum chat ada tidak punya jalan ke chat | FR-COM-005/AC-6, AC-8 | 06 §16.2 "Kirim ke chat", dalam rumusan ringkas `architect`: "Setiap objek yang boleh menjadi kartu dan belum punya kartu di percakapan yang terbuka dapat dikirim ke chat (`share_to_chat`) oleh peserta yang melihatnya, bila objek itu terlihat oleh semua peserta. Tetap satu kartu per objek per percakapan; bila kartunya sudah ada, kartu itu yang dibuka. Status kepedulian tidak dapat dikirim dengan cara ini." | **[Usulan]** `architect` |
| PA-COM-2 | Kartu yang lahir karena entri dicatat atau dikirim dari chat tidak berasal dari transisi yang punya notifikasi | FR-COM-005/AC-6, AC-7 | 06 §16.8, N-23, INV-A46, dalam rumusan ringkas `architect`: "Mengirim ke chat, lewat `share_to_chat` atau dengan mencatat entri dari chat, memicu N-23 yang ikut lipat per pengirim, kecuali perintah yang sama sudah melepas notifikasi lain kepada penerima itu (misalnya N-10, N-13, N-14). Bila notifikasi itu `ringkasan`, N-23 tetap tidak dikirim." | **[Usulan]** `architect` |
| PA-COM-3 | Pesan yang ditulis di layar objek atas objek yang tidak punya kartu | FR-COM-009/AC-7; FR-REQ-041/AC-8 (PA-REQ-6) | 06 §16.3 "Pesan dari layar objek tanpa kartu", dalam rumusan ringkas `architect`: "Pesan yang ditulis di layar objek atas objek yang terlihat oleh semua peserta tetapi belum punya kartu tampil di chat sebagai pesan yang menempel, dengan rujukan ringkas yang dibaca dari state saat ini. Ia tidak melahirkan kartu, dan notifikasinya N-23 saja, bukan N-19." | **[Usulan]** `architect` |
| PA-COM-4 | "Usul lain" atas usulan kesepakatan bila K01 penanggap ditolak guard sesudah K03 berlaku | FR-COM-007/AC-7; FR-REQ-027/AC-8 (PA-REQ-5) | 06 §16.4 (revisi kesepuluh, CA-05): K03 tetap; usulan semula berakhir `ditolak` dan tidak hidup kembali; tidak ada usulan pengganti; penanggap diberi tahu hasilnya dan boleh mengirim K01 lain yang sah; N-12 atas K03 seperti biasa | **[Usulan]** `architect` |

Bukan kesenjangan aturan, dicatat untuk pemiliknya: panjang teks, jenis dan ukuran foto, lama "sedang mengetik", lama tanpa interaksi yang mengakhiri "online", penghapusan dari cache perangkat, cara membedakan penarikan di depan dari ekstensi dan tugas latar (dokumen 09); ketelitian waktu "terakhir aktif" dan bentuk kartu (dokumen 08); daftar akhir medan metadata foto, ekspor dan hak hapus isi chat, isi chat bagi yang keluar dan bagi anggota yang bergabung kemudian, opsi kehadiran P1 dan P2 (dokumen 10).

## Pertanyaan untuk checkpoint S3 dari berkas ini

Nomor Q diberikan instans fase C di dokumen 17. **Hasil (12 November 2026):** kelima bawaan diterima lewat "lanjut": butir 1 sebagai S3-7 (i) (Q-152), butir 3 sebagai S3-7 (ii) (Q-153), butir 2 sebagai D-71 (Q-154), butir 4 sebagai D-72 (Q-155) dengan layar kunci S3-8 (vii) bentuk (a), butir 5 sebagai D-73 (Q-160). Tabel dipertahankan sebagai catatan.

| # | Pertanyaan | Bawaan yang diusulkan | Asal |
|---|---|---|---|
| 1 | "Belum sepakat" pada kartu sama dengan menolak, atau berarti "tetap terbuka, bicara dulu"? | Sama dengan menolak (B08, K03, T03), alasan opsional; yang ingin bicara dulu cukup membalas kartu | 06 P-25 |
| 2 | Hapus pesan: penulis boleh menghapus pesannya sendiri kapan saja, dengan penanda "pesan dihapus" beserta waktunya; tanpa sunting | Ya | 06 P-22 |
| 3 | Balasan pada kartu status sesudah status dihapus: tetap sebagai pesan bebas penulisnya | Ya, dengan batas yang dinyatakan di FR-COM-016/AC-5 | 06 P-21 |
| 4 | Lipat notifikasi chat 10 menit per pengirim dan layar kunci tanpa nama pengirim | Ya; angka diuji di pilot. Teks layar kunci bersama `kabar` lain masuk S3-8 | 06 P-24, P-32 |
| 5 | Foto yang dikirim di chat bukan struk entri; struk dilampirkan dari entri | Ya | FR-COM-003/AC-3 |
