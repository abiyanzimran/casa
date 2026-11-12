# 05 — Kebutuhan fungsional: `REQ` (permintaan, persetujuan, kesepakatan)

Bagian dari dokumen 05. Indeks, tabel keterlacakan, dan aturan umum ada di [../05-functional-requirements.md](../05-functional-requirements.md). Pemilik: `product-analyst`. Tahap S3 fase B, 12 November 2026; diselaraskan dengan revisi kesembilan dokumen 06 (chat rumah tangga dan kartu objek, §16) pada putaran penyelarasan S3, 12 November 2026. **Sesudah checkpoint S3 (12 November 2026):** jawaban "lanjut" dicatat; butir yang dahulu "pertanyaan checkpoint S3" (S3-4 (c), S3-5 (i), S3-6, S3-7 (i), D-71) kini "diterima lewat "lanjut"", dengan penulis aturannya tetap disebut; angka **[Asumsi]** tidak berubah. **Penyisiran andaian khusus pilot S4 (12 November 2026, M1; belum ditinjau `critic`):** BR-REQ-019 dan FR-REQ-027 menyatakan kesepakatan oleh semua pihak yang disyaratkan (06 §6.3, bawaan `all`, alternatif P-37). Perilaku pilot sama. **S4 fase D (12 November 2026, belum ditinjau `critic`):** FR-REQ-029/AC-5 diperjelas dan AC-6 baru (mengakhiri ambang sendirian, bersyarat pada 10-E, bawaan C lebih dulu; angka dari uji 07 INV-40); BR-REQ-038 baru; FR-REQ-001/AC-6 draf permintaan (08 §9.3). **Sesudah checkpoint S4 (12 November 2026; ditinjau `critic` A dan B, 12 November 2026; temuan diperbaiki; perbaikan diperiksa lead dengan `grep`, tanpa tinjauan ketiga):** S4-9 (i) A, diputuskan pemilik produk sendiri: ambang dibuat, diubah, dan diakhiri hanya dengan persetujuan semua pihak yang disyaratkan, di rumah tangga mana pun. BR-REQ-020 dirumuskan ulang; BR-REQ-038 hanya cabang A (nomornya tidak dipakai ulang); FR-REQ-029/AC-5, AC-6 dan FR-REQ-030/AC-3 tidak lagi bersyarat; PA-REQ-8 dibuka dan ditutup (06 §6.3 K01 revisi keempat belas menerima "tanpa ambang"; sisi uangnya 07 §3.8 "Ambang hanya lewat kesepakatan"); AC-6 memakai angka 07 INV-40 yang baru (permintaan 07 §10.2). Hitungan FR dan Must tidak berubah.

| Hal | Isi |
|---|---|
| Lingkup | F-REQ-001 sampai F-REQ-012 (Tier 1, kebutuhan lengkap dengan kriteria penerimaan); F-REQ-013 (Tier 2, ringkas) |
| Rentang ID | `FR-REQ-001` sampai `FR-REQ-045`; `BR-REQ-001` sampai `BR-REQ-038` |
| Sumber aturan | Alur dan pagar: [06](../06-activity-approval-architecture.md). Chat, kartu, tanda terkirim dan dibaca: 06 §16. Uang: [07](../07-financial-domain.md) §3.8, §6.1. Dokumen ini tidak mendefinisikan ulang aturan; ia menyatakan apa yang harus dapat dilakukan dan dilihat pengguna, lalu menunjuk sumbernya |
| Chat | Kebutuhan chat itu sendiri (pesan, foto, tanda, kehadiran, hapus) ada di [COM.md](COM.md). Dokumen ini hanya memuat akibat chat atas permintaan dan kesepakatan: kartu, tanggapan dari kartu, dan diskusi yang tampil di chat |
| Bukan isi dokumen ini | Tata letak, kata-kata layar, dan desain visual (dokumen 08, S4); mekanisme teknis (dokumen 09, 10) |

## Cara membaca

- **Bentuk tiap FR.** Tabel kedudukan, uraian singkat (masukan, validasi, transisi yang dirujuk ke 06, apa yang dilihat tiap pihak, perilaku offline bila 06 menetapkannya), lalu kriteria penerimaan *Given/When/Then*. AC bernomor lokal dan dirujuk sebagai `FR-REQ-017/AC-2`.
- **Pelaku.** "A" dan "B" adalah dua anggota dewasa rumah tangga. Tidak ada AC yang membaca siapa suami atau istri, siapa yang bergaji, atau siapa yang memegang uang. Tiap FR ditutup baris **Dua arah** yang menyatakan apakah AC-nya simetris dan mengapa (06 PG-01, INV-A13: menukar kedua orang dewasa menghasilkan keadaan yang tertukar).
- **Angka contoh.** Ambang permintaan Rp500.000 (contoh 07 §7.1, **[Asumsi]**). Anggaran kategori "jajan dan makan di luar" Rp500.000. Angka tempat uang dikutip dari CTH-21 sampai CTH-24 dan kasus uji 06 §12 tanpa diubah. Hari mulai periode anggaran 1 (bawaan, 07 §6.1). Tanggal dibandingkan pada akhir hari menurut zona waktu rumah tangga, dengan waktu server (06 §4.5, S-6).
- **Istilah model bukan kata layar.** "Pemberi persetujuan", "ditolak", "tanda", "keberatan", "otorisasi" adalah bahasa model; kata-kata layar milik dokumen 08 (Q-119).
- **Kedudukan.** Tingkat kepastian dikutip dari dokumen 17 dan `_progress.md`: *diputuskan pemilik produk sendiri*; *diterima atas saran lead*; *diterima lewat "lanjut"* (default D-nn); **[Usulan]** dengan nama pengusulnya; **[Asumsi]** untuk angka yang diuji di pilot. **[Perlu aturan]** menandai kebutuhan yang aturannya belum ada di 06 maupun 07; permintaannya tercantum di bagian akhir, bersama permintaan yang sudah dijawab revisi kesembilan dokumen 06.
- **Kartu.** "Kartu" adalah tampilan objek di chat rumah tangga (06 §16.2). Setiap AC yang menyebut "dari kartu" berlaku dengan hasil yang sama dari layar objek (06 INV-A44); AC kartu hanya ditulis bila jalur kartu menambah sesuatu yang dapat salah.
- **Keluhan.** AC bertanda "(Keluhan n)" menguji baris keluhan dokumen 17 terhadap aturan yang sudah ada. AC itu tidak menambah aturan.

## Daftar FR

| ID | Judul | Fitur |
|---|---|---|
| FR-REQ-001 | Mengajukan permintaan anggaran | F-REQ-001 |
| FR-REQ-002 | Pemberi persetujuan ditentukan aturan dan dibekukan | F-REQ-001 |
| FR-REQ-003 | Masa berlaku dan jumlah terikat | F-REQ-001 |
| FR-REQ-004 | Draf, pengajuan ganda, dan pengajuan offline | F-REQ-001 |
| FR-REQ-005 | Tempat uang sumber dan reservasi | F-REQ-001 |
| FR-REQ-006 | Syarat permintaan jalan-jalan | F-REQ-002 |
| FR-REQ-007 | Menyetujui | F-REQ-003 |
| FR-REQ-008 | Menolak | F-REQ-003 |
| FR-REQ-009 | Usul balik dan revisi syarat | F-REQ-003 |
| FR-REQ-010 | Tanggapan selalu atas versi yang dibaca | F-REQ-003 |
| FR-REQ-011 | Tunda senyap dan kabar "nanti aku jawab" | F-REQ-003 |
| FR-REQ-012 | Mengusulkan dan menanggapi perubahan syarat | F-REQ-004 |
| FR-REQ-013 | Persetujuan tidak dapat dicabut sendirian | F-REQ-004 |
| FR-REQ-014 | Menautkan belanja saat mencatat | F-REQ-005 |
| FR-REQ-015 | Tautan susulan dan pencabutannya | F-REQ-005 |
| FR-REQ-016 | Kelebihan atas batas | F-REQ-005 |
| FR-REQ-017 | Menyelesaikan dan sisa menurut situasi custody | F-REQ-005 |
| FR-REQ-018 | Penutupan otomatis dan kedaluwarsa otorisasi | F-REQ-005 |
| FR-REQ-019 | Tampilan yang sama bagi kedua pihak | F-REQ-005 |
| FR-REQ-020 | Mencatat offline dengan tautan | F-REQ-005 |
| FR-REQ-021 | Membatalkan | F-REQ-006 |
| FR-REQ-022 | Kedaluwarsa permintaan yang menunggu | F-REQ-006 |
| FR-REQ-023 | Mengajukan lagi dan rantai permintaan | F-REQ-006 |
| FR-REQ-024 | Pengingat ulang oleh pihak yang menunggu | F-REQ-007 |
| FR-REQ-025 | Pengingat otomatis | F-REQ-007 |
| FR-REQ-026 | Mengusulkan kesepakatan | F-REQ-008 |
| FR-REQ-027 | Menyetujui, menolak, menarik, dan menyetujui beberapa pos sekaligus | F-REQ-008 |
| FR-REQ-028 | Penggantian kesepakatan dan kapan berlaku | F-REQ-008 |
| FR-REQ-029 | Mengakhiri kesepakatan rutin | F-REQ-008 |
| FR-REQ-030 | Ambang saat mulai memakai dan rumah tangga satu dewasa | F-REQ-008 |
| FR-REQ-031 | Jalan yang terlihat untuk mengubah kesepakatan | F-REQ-008 |
| FR-REQ-032 | Penilaian tanda oleh server | F-REQ-009 |
| FR-REQ-033 | Yang tampil pada tanda dan tanggapan pasangan | F-REQ-009 |
| FR-REQ-034 | Pemberitahuan atas tanda | F-REQ-009 |
| FR-REQ-035 | Tanda hilang dan muncul kembali | F-REQ-009 |
| FR-REQ-036 | Tanpa hitungan dan papan skor | F-REQ-009 |
| FR-REQ-037 | Mengajukan konfirmasi susulan | F-REQ-010 |
| FR-REQ-038 | Menanggapi konfirmasi susulan | F-REQ-010 |
| FR-REQ-039 | Pengingat dan pengajuan ulang konfirmasi susulan | F-REQ-010 |
| FR-REQ-040 | Tampilan konfirmasi susulan | F-REQ-010 |
| FR-REQ-041 | Diskusi pada objek | F-REQ-011 |
| FR-REQ-042 | Lampiran pada objek | F-REQ-011 |
| FR-REQ-043 | Daftar yang menunggu saya dan yang menunggu pasangan | F-REQ-012 |
| FR-REQ-044 | Mengajukan dan menanggapi permintaan tindakan (Tier 2) | F-REQ-013 |
| FR-REQ-045 | Tugas yang diterima dan pengingatnya (Tier 2) | F-REQ-013 |

## Business rule

Setiap baris adalah pernyataan tingkat produk yang menunjuk sumber kanonisnya. Bila bunyinya tampak berbeda dari sumber, sumber yang berlaku.

| ID | Pernyataan | Sumber kanonis | Kedudukan |
|---|---|---|---|
| BR-REQ-001 | Permintaan anggaran hanya atas dana bersama; jumlah yang diminta bilangan rupiah utuh > 0 dan sama dengan *jumlah pemakaian* yang kelak menjadi realisasinya. Pemakaian dana pribadi tidak pernah memerlukan permintaan | 06 §4.1, B01; 07 §3.8, §2.1 | D-09, D-15 (diterima lewat "lanjut") |
| BR-REQ-002 | Pemberi persetujuan ditentukan `resolve_approvers` saat pengajuan dan dibekukan; pemohon selalu dikecualikan; di pilot itu anggota dewasa yang lain. Dengan satu anggota dewasa pengajuan ditolak `no_approver` | 06 §3.2, §3.5; ADR-010 | D-19, D-33 (lewat "lanjut"); Q-30 masih terbuka |
| BR-REQ-003 | "Disetujui" berarti pemohon dan pemberi persetujuan sama-sama menyatakan setuju atas versi syarat yang sama. `system` tidak pernah membawa permintaan ke disetujui atau ditolak | 06 §3.3, INV-A04, INV-A10 | Keputusan pemilik nomor 6 (diam bukan jawaban) |
| BR-REQ-004 | Versi syarat tidak dapat diubah; setiap tanggapan menyebut versi yang ditanggapinya; tanggapan atas versi yang digantikan ditolak `terms_superseded` | 06 INV-A06, INV-A07, S-3 | Dirancang 06 |
| BR-REQ-005 | Diam tidak pernah berarti setuju atau menolak. Permintaan tanpa tanggal tidak kedaluwarsa selama menunggu; yang bertanggal kedaluwarsa, tidak terselesaikan sendiri | 06 PG-04, PG-05, B10, §4.5 | Keputusan pemilik nomor 6; D-20 |
| BR-REQ-006 | Otorisasi selalu punya masa berlaku; bawaan tanggal rencana + 7 hari, atau tanggal persetujuan + 30 hari bila tanpa tanggal rencana; masa berlaku dapat dirundingkan | 06 §4.5 | Diputuskan pemilik produk sendiri (keputusan 23, S2-7); angka **[Asumsi]** |
| BR-REQ-007 | Realisasi, *terikat*, sisa, dan kelebihan dihitung menurut 07; *terikat* dihitung pada periode anggaran tanggal rencana pada versi yang disetujui, atau tanggal persetujuan | 07 §3.8, §6.1, INV-14; 06 §5.5 | D-08 (lewat "lanjut") |
| BR-REQ-008 | Reservasi hanya lahir bila permintaan menyebut tempat uang sumber; ia mengurangi saldo tersedia, bukan saldo tercatat, dan tidak menghalangi pencatatan apa pun | 07 §3.8 "Reservasi" | D-06 (lewat "lanjut") |
| BR-REQ-009 | Sisa tetap dana bersama dan tidak pernah menjadi pemasukan atau dana pribadi. Situasi custody A sampai D dibaca dari entri yang tertaut, tidak dipilih; pada situasi B pemohon ditanya satu kali (bawaan: tetap dipegang, T-02); pada situasi C hak ganti sebesar realisasi | 07 §3.8 tabel sisa; 06 §5.3 | D-07, D-14 (lewat "lanjut") |
| BR-REQ-010 | Kelebihan atas batas tidak memerlukan tanggapan pasangan; pasangan diberi tahu satu kali (N-08); bagian yang melebihi batas dibandingkan dengan ambang, per entri | 06 §5.4; 07 §3.8 "Tanda atas entri" | D-07 (lewat "lanjut") |
| BR-REQ-011 | Pencatatan tidak pernah dihalangi atau ditahan oleh state permintaan atau kesepakatan mana pun | 06 INV-A11, S-7; 07 A-20 | Dirancang 06; ADR-009 |
| BR-REQ-012 | Penyelesaian hanya oleh pemohon (B15) atau `system` (B16); pemberi persetujuan tidak dapat menyelesaikan atau menutup permintaan orang lain | 06 §5.1 | Dirancang 06 |
| BR-REQ-013 | Persetujuan tidak dapat dicabut sepihak. Otorisasi terbuka hanya berubah lewat perubahan syarat (B11 sampai B13); paling banyak satu perubahan yang menunggu | 06 §4.4, PG-11, INV-A14 | Diputuskan pemilik produk sendiri (keputusan 21, S2-3) |
| BR-REQ-014 | Tautan susulan ke otorisasi hanya oleh pencatat entri atau pemohon otorisasi; pencabutan hanya oleh pembuat tautan atau pencatat entri; tautan ke tagihan berulang hanya oleh pencatat entri. Entri terlambat dapat ditautkan ke permintaan `settled` atau `expired` bila tanggalnya di antara tanggal persetujuan dan tanggal penutupan; permintaan `cancelled` tidak menerima tautan | 06 §4.5, INV-A34; 07 §3.8 "Penautan", INV-26 | D-39 (lewat "lanjut"); tautan ke tagihan berulang **[Usulan]** `architect` |
| BR-REQ-015 | Permintaan yang ditolak, dibatalkan, atau kedaluwarsa dapat diajukan lagi tanpa batas jumlah, menunjuk asalnya; semua yang terhubung lewat `resubmitted_from` satu rantai permintaan | 06 §4.4 | D-24 (lewat "lanjut") |
| BR-REQ-016 | Pengingat ulang: paling banyak satu per 6 jam per pengirim per rantai; paling banyak tiga per pengirim per objek seumur objek; satu atas konfirmasi susulan. Isinya kalimat baku tanpa jumlah | 06 §10 tabel batas, INV-A32 | 6 jam: keputusan pemilik nomor 13 (sendiri). Tiga: diterima atas saran lead (S2-5), **[Asumsi]**. Isi: D-29 |
| BR-REQ-017 | Pengingat otomatis paling banyak dua per rantai permintaan; tidak ada atas konfirmasi susulan | 06 §10 tabel batas | D-20 (lewat "lanjut"); konfirmasi susulan diterima atas saran lead, 12 November 2026 |
| BR-REQ-018 | Tunda senyap 24 jam tidak terbaca pihak yang menunggu; kabar "nanti aku jawab" adalah perintah terpisah yang bawaannya tidak dikirim. Batas yang dinyatakan terus terang: di chat, kartu yang sudah dibaca memperlihatkan bahwa penunda sudah melihatnya; yang tetap tertutup adalah apakah ia menunda dan sampai kapan | 06 §4.4, INV-A33, AT-5 (c)(d), §16.9 | Diterima atas saran lead (S2-5); 24 jam **[Asumsi]**. Batas di chat: akibat "Ikut tanda dibaca" (diterima atas saran lead, 12 November 2026), dirumuskan 06 AT-5 (d) |
| BR-REQ-019 | Kesepakatan hanya berlaku lewat pernyataan setuju semua pihak yang disyaratkan (bawaan semua anggota dewasa penuh selain pengusul, dibekukan saat K01; satu "belum sepakat" cukup untuk menolak; pemilik jatah selalu wajib; alternatif mayoritas P-37; pada dua anggota dewasa: pasangan pengusul; penyisiran S4, M1); usulan tidak kedaluwarsa; yang lama berlaku sampai penggantinya disetujui | 06 §6.3 K01 sampai K06, INV-A18, INV-A19, PG-04 | Dirancang 06 |
| BR-REQ-020 | Ambang permintaan, selama ada, selalu > 0; ia dibuat, diubah, dan diakhiri hanya bila semua pihak yang disyaratkan setuju (pada dua anggota dewasa: keduanya). Tidak ada perintah sepihak untuk mengubah atau mengakhirinya; mengakhiri adalah usulan `rule` "tanpa ambang" yang disetujui semua pihak | 07 §3.8 "Ambang hanya lewat kesepakatan", INV-40; 06 §6.3 K01 (usulan `rule` "tanpa ambang", revisi keempat belas), K02, K05; tanpa perintah sepihak | Keputusan pemilik nomor 5 (satu ambang bersama). Dibuat, diubah, dan diakhiri hanya bersama: **diputuskan pemilik produk sendiri** di checkpoint S4, S4-9 (i) A ("Ambang itu batas yang di tentukan/disepakati 2 orang"; "ambang di akhiri berdua juga"), berlaku bagi semua rumah tangga. Sebelumnya bagian "hanya berdua" berasal dari S2-4 yang tidak ia sebut dan diterapkan lead |
| BR-REQ-021 | Hari mulai periode anggaran 1 sampai 28; penggantiannya kesepakatan berdua dan berlaku mulai periode berikutnya lewat satu periode peralihan | 07 §6.1; 06 §6.3 | D-49, D-50 (lewat "lanjut"); bahwa penggantian berdua adalah putusan lead |
| BR-REQ-022 | Jatah pribadi yang tidak di atas ambang, termasuk 0, hanya dapat diusulkan pemiliknya; jatah hanya dapat diakhiri pemiliknya; jatah di atas ambang yang diakhiri tetap menjadi dasar penilaian sampai akhir periode itu | 06 §6.3 K01, K06 | D-52 (lewat "lanjut"); masa penilaian: penegasan 1; batas "tidak di atas ambang" **[Usulan]** `security-privacy` SP-32 sampai SP-34, disampaikan dan tidak dibantah; batas "sama dengan" pengetatan lead **[Usulan]** |
| BR-REQ-023 | Setiap tagihan berulang disepakati saat dibuat, berapa pun jumlahnya; jumlah per kali yang disepakati adalah batasnya; salah satu anggota dewasa boleh mengakhiri kesepakatannya | 06 §6.3 "Pos rutin", INV-A43; 07 §3.15 | Keputusan 22 dan penegasan 3 (sendiri); mengakhiri sendirian: bawaan lama dengan akibat yang belum disajikan (P-20) |
| BR-REQ-024 | Tanda dibuat server, satu kali saat entri diterima (dan saat tautan susulan ditambahkan atau dicabut), dengan salah satu dari tiga dasar menurut rumus 07; tidak dihitung ulang karena entri lain atau perubahan aturan | 06 §6.2, INV-A17, INV-A37, S-8; 07 §3.8 "Tanda atas entri" | Dirancang 06 dan 07; ADR-009 |
| BR-REQ-025 | Yang tidak pernah ditandai: serah terima, pemindahan, `contribution_expense`, penggantian, pelunasan kewajiban yang sudah tercatat (termasuk cicilan), saldo awal sebelum ada ambang, pembayaran pos rutin sampai jumlah per kalinya, penyisihan di dalam jatah | 06 §6.2 daftar "tidak ditandai" | D-44, D-46, D-48 (lewat "lanjut") |
| BR-REQ-026 | Tanda hanya ada bila rumah tangga punya sekurangnya dua anggota dewasa saat entri disimpan; tanda dasar ambang hanya bila ada ambang yang berlaku | 06 §3.5; 07 §3.8 | D-33, D-42 (lewat "lanjut") |
| BR-REQ-027 | Pemberitahuan langsung atas tanda hanya atas dasar ambang, paling banyak satu per penerima per hari; dasar jatah dan dasar pos rutin hanya ke ringkasan; di ringkasan semua butir tanda dilebur menjadi satu baris tanpa hitungan | 06 §6.2 "Notifikasi", N-10 | Satu per hari: D-35 (lewat "lanjut"). Hanya dasar ambang: diterima atas saran lead, 12 November 2026. Satu baris: **[Usulan]** `security-privacy` SP-43 butir b |
| BR-REQ-028 | Tanda hilang hanya lewat tautan susulan yang sah atau pembalikan entri; muncul kembali hanya bila pembuat tautan atau pencatat entri mencabut tautan itu | 06 §6.2 "Tanda hilang", "Tanda muncul kembali" | D-39 (lewat "lanjut") |
| BR-REQ-029 | Konfirmasi susulan adalah permintaan anggaran dengan `confirms_entry_id`: hanya pencatat yang memulai, syaratnya ditetapkan server, tanpa usul balik dan tanpa medan alasan penolakan, disetujui dalam satu transaksi yang menutupnya, tanpa pengingat otomatis, satu pengingat ulang, tidak dapat diajukan lagi sesudah ditolak; N-01 dan pengingat ulangnya selalu `ringkasan` | 06 §6.2 "Konfirmasi susulan", INV-A40, N-01, N-20; 07 §3.8 | Jalur: usulan lead (Q-112). Tiga batas: diterima atas saran lead, 12 November 2026. Selalu `ringkasan`: usulan SP-52. Rincian lain: usulan `security-privacy` SP-38 dan `architect` (Q-124, Q-125, Q-130). Jalur, SP-52, dan rinciannya diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-6 (i) sampai (v) |
| BR-REQ-030 | Diskusi menempel pada objek induknya, tersedia di setiap state, tidak pernah mengubah state, tidak dapat disunting; penulis dapat menghapus pesannya sendiri, dengan penanda "pesan dihapus" yang tetap. Atas objek yang punya kartu di chat, diskusinya adalah balasan yang menempel pada kartu itu, tampil di chat dan di layar objek dengan isi yang sama, dan notifikasinya N-23 `kabar`. N-19 `ringkasan` tinggal untuk diskusi objek yang tidak tampil di chat | 06 §2.1, §4.4, §16.3, §16.7, §16.8, N-19, N-23, INV-A44, INV-A46; ADR-001, ADR-011 | Dirancang 06. Diskusi tampil di chat: diterima atas saran lead ("Chat umum + kartu", "Reply nempel ke kartu"), 12 November 2026 |
| BR-REQ-031 | Syarat, tanggapan, diskusi, dan lampiran permintaan hanya terbaca para pihaknya; kesepakatan terbaca semua anggota dewasa; permintaan tidak dapat disembunyikan dari pemberi persetujuannya; visibilitas tidak menyembunyikan akibat uangnya | 06 §9, V-2, V-3, INV-A16 | D-25 (lewat "lanjut") |
| BR-REQ-032 | Tidak ada umpan kegiatan per orang, hitungan atau peringkat per orang (tanda, keberatan, pengingat, lama menanggapi), lokasi, atau penyaring dan kolom ekspor menurut tanda. Tanda "sudah dibaca", "terakhir aktif", dan "sedang online" tidak ada di luar layar chat; di dalam chat, tanda terkirim dan dibaca pada pesan dan kartu serta kehadiran ada, sama bagi keduanya, tanpa waktu baca, dan tidak dibaca aturan, pengingat, ringkasan, atau ekspor mana pun. Tidak ada pengingat "dibaca tapi belum dijawab" | 06 §9.4 AT-1 sampai AT-7 (AT-5 dirumuskan ulang), PG-07, PG-09, §16.5, §16.6, INV-A45 | D-23, D-43 (lewat "lanjut"). Tanda di chat: pernyataan pemilik produk sendiri ("ada tanda di baca/delivery chat itu wajib loh") dan "Ikut tanda dibaca" (diterima atas saran lead). Kehadiran: **diputuskan pemilik produk sendiri**, berlawanan dengan saran lead, 12 November 2026. Tanpa waktu baca: **[Usulan]** `architect` |
| BR-REQ-033 | Tidak seorang pun menyetujui, menolak, mengonfirmasi, atau menandai selesai atas nama orang dewasa lain | 06 PG-02 | Dirancang 06 |
| BR-REQ-034 | Server satu-satunya yang memutuskan transisi; setiap perintah membawa `idempotency_key` dan state yang diharapkannya; waktu server yang menentukan | 06 S-1, S-2, S-4, S-6; `ENGINEERING-STANDARDS.md` §2 | Dirancang 06 |
| BR-REQ-035 | Keluar dari rumah tangga tidak menunggu objek alur apa pun; objek yang satu-satunya pelaku berwenangnya sudah keluar keluar dari daftar terbuka | 06 §3.5, INV-A35 | D-42 (lewat "lanjut") |
| BR-REQ-036 | Permintaan tindakan mengikuti T01 sampai T06: hanya penerima tugas yang menerima, menolak, menandai selesai, atau melepas; tanpa kedaluwarsa dan tanpa transisi `system` | 06 §13.1 | D-27 (lewat "lanjut"); T06 **[Usulan]** `architect` |
| BR-REQ-038 | **Hanya cabang A (sesudah checkpoint S4).** Tidak ada pengakhiran ambang sendirian di rumah tangga mana pun, pilot maupun tahap publik; perintah sepihak atas ambang ditolak tanpa akibat; ambang diakhiri hanya lewat kesepakatan (BR-REQ-020). Nomor tidak dipakai ulang | 06 §6.3 K01, K02, K05 (revisi keempat belas; K07 tidak ada); 07 §3.8 "Ambang hanya lewat kesepakatan", INV-40; 10 §3 10-E | S4-9 (i) A, **diputuskan pemilik produk sendiri**; saran `security-privacy` (B) dan bawaan C tidak diambil |
| BR-REQ-037 | Permintaan, kesepakatan, konfirmasi susulan, serah terima yang menunggu, dan permintaan tindakan mendapat kartu di chat saat dibuat. Isi kartu dibaca dari state objek saat ini. Tombol kartu adalah perintah objek yang sama dengan layar objek, dengan versi syarat dan state yang diharapkan; ketidakcocokan ditolak `terms_superseded` atau `state_changed` dan kartu menampilkan state terbaru. Pesan apa pun, termasuk "setuju", tidak mengubah state. Kartu tidak memicu notifikasi sendiri dan tidak dapat dihapus | 06 §16.2, §16.4, S-3, S-4, INV-A44, INV-A46, INV-A47 (b) | Dari chat dapat dikirim tugas, pengeluaran, dan rencana: pernyataan pemilik produk sendiri ("chat itu ada/diskusi di chat bisa kirim task/pengeluaran/rencana"). Permintaan muncul sebagai kartu yang ditanggapi dari chat: diterima atas saran lead ("Chat umum + kartu", "Chat naik Tier 1"). Daftar kartu otomatis dan rinciannya **[Usulan]** `architect` |

---

## F-REQ-001 — Permintaan anggaran

### FR-REQ-001 — Mengajukan permintaan anggaran

| | |
|---|---|
| Fitur | F-REQ-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `KS-4` |
| Aktor dan izin | Anggota dewasa mana pun menjadi pemohon. Syarat hanya terbaca pemohon dan pemberi persetujuan (06 §9, INV-A16); tidak dapat disembunyikan dari pemberi persetujuan (V-2) |
| Aturan acuan | 06 §4.1, §4.3 B01, §3.5; 07 §3.8 (tabel *jumlah pemakaian*), §2.6 |
| Business rule | BR-REQ-001, BR-REQ-002, BR-REQ-011, BR-REQ-034 |
| Kedudukan | Kapan diajukan: keputusan pemilik nomor 5. Jenis pemakaian bukan pengeluaran lewat formulir yang sama: D-09 (lewat "lanjut"). Tambahan dana pribadi untuk satu periode lewat permintaan biasa: penegasan 2 (D-53) |

**Masukan.** Judul; jumlah; jenis pemakaian (`use_kind`: pengeluaran, pinjaman yang diberikan, pembelian aset, penyisihan ke dana pribadi); kategori, tujuan, pihak luar (untuk pinjaman), tanggal rencana, masa berlaku (FR-REQ-003), tempat uang sumber (FR-REQ-005), dan catatan, semuanya opsional.

**Validasi server (guard B01).** Jumlah bilangan rupiah utuh > 0; sumbernya dana bersama; `resolve_approvers` tidak kosong. Untuk pembelian aset secara cicilan, jumlah yang diminta adalah uang muka ditambah seluruh cicilan (06 §4.1; CTH-34: 400.000 + 3.960.000 = 4.360.000). Formulir menjelaskan hal ini; aplikasi tidak menebak.

**Hasil.** State `awaiting_approver`, syarat versi 1, Peristiwa `request.submitted` tanpa akibat uang (INV-A09), N-01 kepada pemberi persetujuan. Sebelum disetujui tidak ada *terikat*, reservasi, atau entri.

**Kapan dipakai.** Atas pilihan pemohon untuk rencana berapa pun jumlahnya; menurut kesepakatan untuk pemakaian tidak rutin di atas ambang (06 §4.1). Bila tidak diajukan, pencatatan tetap berjalan dan entrinya mungkin bertanda (FR-REQ-032).

**Kriteria penerimaan**
- **AC-1** — *Given* A dan B anggota dewasa, anggaran "jajan dan makan di luar" Rp500.000 dengan terpakai Rp0, *When* A mengajukan Rp100.000 berjenis pengeluaran pada kategori itu, *Then* permintaan berada di `awaiting_approver` versi 1, B menerima N-01 berprioritas `perlu_tanggapan`, dan anggaran kategori itu tetap menampilkan terikat Rp0 dan sisa Rp500.000 kepada keduanya.
- **AC-2** — *Given* formulir permintaan, *When* A mengirim jumlah Rp0, jumlah negatif, atau jumlah bukan rupiah utuh, *Then* server menolak, tidak ada permintaan yang tersimpan, dan B tidak menerima apa pun.
- **AC-3** — *Given* rumah tangga baru berisi satu anggota dewasa, *When* A membuka pengajuan permintaan, *Then* formulir pengajuan tidak ditawarkan dan alasannya disebut; perintah `submit` yang tetap dikirim ditolak `no_approver`; mencatat, anggaran, dan laporan tetap berjalan penuh (06 §3.5).
- **AC-4** — *Given* A memilih tempat uang pribadinya sebagai sumber, *When* A mengirim permintaan, *Then* server menolak karena permintaan hanya atas dana bersama (BR-REQ-001).
- **AC-5** — *Given* A ingin tambahan Rp300.000 ke dana pribadinya untuk satu periode saja, *When* A mengajukannya dengan jenis penyisihan ke dana pribadi dan B menyetujuinya, lalu penyisihan itu dicatat dan ditautkan, *Then* penyisihan itu tidak terhitung terhadap jatah pribadi A (D-48: yang sudah disetujui lewat permintaan tidak terhitung) dan tidak bertanda dasar jatah.
- **AC-6** — *Given* A mengisi permintaan Rp100.000 lalu menutup lembar tanpa mengajukan, *When* A membukanya lagi, *Then* aplikasi menawarkan "Lanjutkan yang tadi"; draf hanya di perangkat A, tidak terbaca B, tidak menghasilkan Peristiwa, N-01, atau *terikat*, dan dibuang saat diajukan atau saat A memilih mulai baru (**[Usulan]** `ux-designer`, 08 §9.3; sama dengan FR-FIN-040/AC-16). Pengajuan tetap tunduk pada jeda urungkan 5 detik (LM-12).

**Dua arah.** Simetris. Guard B01 tidak membaca peran, penghasilan, atau pemegang uang; menukar A dan B pada setiap AC menghasilkan keadaan yang tertukar (INV-A13).

### FR-REQ-002 — Pemberi persetujuan ditentukan aturan dan dibekukan

| | |
|---|---|
| Fitur | F-REQ-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` (setara sejak awal) |
| Aktor dan izin | Ditentukan server, bukan dipilih pemohon. Pemohon tidak pernah menyetujui permintaannya sendiri (INV-A04); hanya anggota di `required_approvals` yang bertindak sebagai pemberi persetujuan (INV-A05) |
| Aturan acuan | 06 §3.2, §3.3, §3.5 (tabel "Dari dua menjadi satu", "Selama masa tenggang"), §3.6 F06, F08, INV-A48 (e); 07 §8.2 butir 2b; ADR-010 |
| Business rule | BR-REQ-002, BR-REQ-003, BR-REQ-035 |
| Kedudukan | D-19, D-33 (lewat "lanjut"). Tanpa layar pengaturan izin di pilot: Q-30, diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-1 (a). Pemberi persetujuan selama masa tenggang keluar (AC-5): usulan `security-privacy` SP-48 dan `architect` (06 §3.5 "Selama masa tenggang", menutup PA-FAM-9), diterima lewat "lanjut", S3-4 (c); harganya W-24; 7 hari **[Asumsi]** |

Daftar pemberi persetujuan dihitung satu kali saat B01 dan dibekukan; perubahan keanggotaan sesudahnya tidak menghitungnya ulang. Tidak ada layar pengaturan izin di pilot.

**Kriteria penerimaan**
- **AC-1** — *Given* rumah tangga berisi A dan B, *When* A mengajukan permintaan, *Then* `required_approvals` berisi B saja; *When* B mengajukan permintaan, *Then* berisi A saja.
- **AC-2** — *Given* permintaan A menunggu di `awaiting_approver`, *When* A mengirim perintah `approve` atasnya, *Then* server menolak, state tidak berubah, dan tidak ada Peristiwa tanggapan.
- **AC-3** — *Given* permintaan A tanpa tanggal yang menunggu 90 hari, *When* waktu berjalan, *Then* permintaan tetap `awaiting_approver`; tidak ada Peristiwa ber-`actor` `system` yang menyetujui atau menolaknya (INV-A10).
- **AC-4** — *Given* permintaan A menunggu tanggapan B, *When* B keluar dari rumah tangga, *Then* daftar pemberi persetujuan tidak dihitung ulang, permintaan tampil "pemberi persetujuan sudah keluar" dan keluar dari daftar menunggu A, dan A tetap dapat membatalkannya (B09); tidak ada jalan menuju `approved`. Tampilan ini berlaku sejak F06, saat akses B berakhir, bukan baru pada akhir masa tenggang (06 §3.5 "Kapan tabel di bawah berlaku").
- **AC-5** — *Given* B menjalankan F06 dan berada dalam masa tenggang keluar (`keluar_tenggang`), *When* A mengajukan permintaan baru, *Then* formulir pengajuan tidak ditawarkan dan alasannya disebut, dan perintah `submit` yang tetap dikirim ditolak `no_approver`, karena B bukan anggota dewasa aktif dan bukan pemberi persetujuan sejak F06 (06 §3.2, §3.5 "Selama masa tenggang", INV-A48 (e)); *Given* usulan kesepakatan A yang menunggu B, *When* waktu berjalan atau siapa pun mengirim K02 atasnya, *Then* tidak ada yang dapat menyetujuinya: K02 ditolak, A tetap dapat menariknya (K04), dan usulan B yang menunggu A dapat A tolak (K03), tetapi tidak dapat A setujui; *When* B kembali lewat F08 sebelum F07, *Then* permintaan dan usulan yang masih menunggu kembali ke daftar menunggu B dan dapat B tanggapi, dan tidak ada yang berlaku surut (06 F08). Mencatat tidak terhalang selama masa tenggang (LM-02), dan entri A di atas ambang tetap ditandai karena B masih terhitung anggota dewasa dalam aturan uang sampai F07 (07 §8.2 butir 2b; FR-FAM-012/AC-4).

**Dua arah.** Simetris. Satu-satunya sifat yang dibaca adalah "bukan pemohon"; tidak ada urutan atau peran yang tetap.

### FR-REQ-003 — Masa berlaku dan jumlah terikat

| | |
|---|---|
| Fitur | F-REQ-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` (sisa anggaran memperhitungkan yang sudah dijanjikan); `KS-1` |
| Aktor dan izin | Masa berlaku diisi pemohon atau diisi bawaan saat disetujui; dirundingkan seperti syarat lain. Akibatnya pada anggaran terlihat oleh siapa pun yang melihat anggaran itu (V-3) |
| Aturan acuan | 06 §4.5, §5.5; 07 §3.8, §6.1 |
| Business rule | BR-REQ-006, BR-REQ-007 |
| Kedudukan | Bawaan +7 dan +30 hari: diputuskan pemilik produk sendiri (keputusan 23, S2-7); angka **[Asumsi]**. Periode *terikat*: D-08 (lewat "lanjut") |

Bila pemohon tidak mengisi masa berlaku, server mengisinya pada versi yang disetujui: tanggal rencana + 7 hari, atau tanggal persetujuan + 30 hari. Selama menunggu, permintaan hanya kedaluwarsa bila versi berjalan memuat tanggal rencana atau masa berlaku yang diisi pemohon (FR-REQ-022). *Terikat* = max(0, batas − realisasi) dan mengurangi sisa anggaran kategori pada periode tanggal rencana, atau periode tanggal persetujuan.

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A Rp100.000 tanpa tanggal rencana dan tanpa masa berlaku, *When* B menyetujuinya pada 3 Desember 2026, *Then* masa berlaku otorisasi 2 Januari 2027, terlihat sama oleh keduanya.
- **AC-2** — *Given* permintaan jalan-jalan dengan tanggal rencana 20 Desember 2026, *When* disetujui pada 5 Desember 2026, *Then* masa berlakunya 27 Desember 2026.
- **AC-3** — *Given* anggaran "jajan dan makan di luar" Rp500.000 dengan terpakai Rp0 (CTH-21), *When* permintaan A Rp100.000 berkategori itu disetujui, *Then* anggaran menampilkan terpakai Rp0, terikat Rp100.000, sisa Rp400.000; *When* belanja Rp75.000 ditautkan, *Then* Rp75.000, Rp25.000, Rp400.000; *When* A menyelesaikannya, *Then* Rp75.000, Rp0, Rp425.000.
- **AC-4** — *Given* periode anggaran Desember (1 sampai 31 Desember 2026), *When* permintaan bertanggal rencana 5 Januari 2027 disetujui pada 20 Desember 2026, *Then* *terikat*-nya mengurangi sisa anggaran periode Januari, bukan Desember.
- **AC-5** — *Given* B mengajukan usul balik yang mengubah masa berlaku, *When* A menerimanya, *Then* masa berlaku pada versi yang disetujui itulah yang dipakai otorisasi.
- **AC-6** — *Given* permintaan tanpa kategori, *When* disetujui, *Then* tidak ada sisa anggaran kategori mana pun yang berkurang.

**Dua arah.** Simetris. Perhitungan tanggal dan jumlah tidak bergantung pada siapa pemohonnya.

### FR-REQ-004 — Draf, pengajuan ganda, dan pengajuan offline

| | |
|---|---|
| Fitur | F-REQ-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `MSL-5` (sinyal buruk tidak boleh menggandakan atau menghilangkan pengajuan) |
| Aktor dan izin | Draf hanya ada di perangkat pemohon dan tidak terbaca siapa pun |
| Aturan acuan | 06 §4.2 (draf bukan state server), S-1, S-2, S-6, INV-A12; §10 "Yang tidak ditutup" (permintaan yang diketik ulang) |
| Business rule | BR-REQ-034 |
| Kedudukan | Dirancang 06. Draf disimpan kembali bila pengajuan offline ditolak server: **[Usulan]** `product-analyst` |

**Kriteria penerimaan**
- **AC-1** — *Given* A mengetuk "ajukan" dua kali, atau perangkat mengirim ulang karena jaringan buruk, dengan `idempotency_key` yang sama, *When* server menerima kedua kiriman, *Then* hanya ada satu permintaan dan satu N-01; kiriman kedua mengembalikan hasil yang pertama.
- **AC-2** — *Given* A offline, *When* A mengajukan permintaan, *Then* perangkat A menampilkannya "menunggu terkirim", tidak pernah "diajukan", dan B tidak melihat apa pun sampai server menerimanya; saat tersambung server menjalankan B01 dengan waktu server.
- **AC-3** — *Given* pengajuan jalan-jalan dibuat offline dengan tanggal rencana 14 Desember 2026, *When* perangkat baru tersambung pada 15 Desember 2026, *Then* server menolaknya karena tanggal sudah lampau, A diberi tahu alasannya, dan isian A tetap tersedia sebagai draf di perangkatnya.
- **AC-4** — *Given* A menyimpan draf, *When* B membuka aplikasi, *Then* B tidak melihat draf itu, tidak ada Peristiwa, dan tidak ada notifikasi.
- **AC-5** — *Given* A mengetik dua permintaan dengan isi sama tanpa memakai "ajukan lagi", *When* keduanya diajukan, *Then* keduanya permintaan terpisah dengan rantai masing-masing; aplikasi tidak menggabungkan atau menebak kesamaannya.

**Dua arah.** Simetris; perilaku jaringan tidak bergantung pada pelaku.

### FR-REQ-005 — Tempat uang sumber dan reservasi

| | |
|---|---|
| Fitur | F-REQ-001 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`, `MSL-2` (uang yang sudah dijanjikan terlihat di tempat uangnya) |
| Aktor dan izin | Pemohon boleh menyebut tempat uang dana bersama mana pun, juga yang dipegang pasangannya; reservasi terlihat oleh siapa pun yang melihat tempat uang itu |
| Aturan acuan | 07 §3.8 "Reservasi", tabel sisa situasi B dan D; CTH-22 |
| Business rule | BR-REQ-008, BR-REQ-011 |
| Kedudukan | D-06 (lewat "lanjut") |

Reservasi lahir saat disetujui, sebesar batas, pada tempat uang sumber. Ia berkurang oleh serah terima atau pembayaran yang tertaut dari tempat uang itu dan dilepas saat penyelesaian, pembatalan, atau kedaluwarsa. Ia tidak pernah menolak pencatatan: bila saldo tersedia menjadi negatif, tempat uang menampilkan "kurang Rp… untuk yang sudah dijanjikan".

**Kriteria penerimaan**
- **AC-1** — *Given* `TUN-S`, tempat uang dana bersama yang dipegang B, tercatat Rp425.000 (CTH-22), *When* A mengajukan Rp100.000 dengan sumber `TUN-S` dan B menyetujuinya, *Then* `TUN-S` tercatat Rp425.000, reservasi Rp100.000, tersedia Rp325.000, dan tidak ada entri yang lahir.
- **AC-2** — *Given* AC-1, *When* B mencatat serah terima Rp100.000 dari `TUN-S` ke tempat uang A dan menautkannya ke permintaan itu, *Then* `TUN-S` tercatat Rp325.000, reservasi Rp0, tersedia Rp325.000.
- **AC-3** — *Given* AC-1, *When* B membayar langsung Rp75.000 dari `TUN-S` dan menautkannya (situasi D), *Then* `TUN-S` tercatat Rp350.000, reservasi Rp25.000, tersedia Rp325.000; *When* A menyelesaikan permintaan, *Then* reservasi Rp0 dan tersedia Rp350.000.
- **AC-4** — *Given* AC-1, *When* B mencatat belanja lain Rp400.000 dari `TUN-S` yang tidak ditautkan, *Then* belanja itu tersimpan, `TUN-S` tercatat Rp25.000 dan tersedia −Rp75.000, dan tempat uang itu menampilkan "kurang Rp75.000 untuk yang sudah dijanjikan" kepada keduanya; tidak ada yang ditolak.
- **AC-5** — *Given* permintaan yang disetujui tanpa tempat uang sumber, *When* disetujui, *Then* tidak ada reservasi di tempat uang mana pun.

**Dua arah.** Simetris. Pemohon boleh menyebut tempat uang yang dipegang siapa pun; reservasi hanya keterangan dan tidak memberi kendali atas pemegangnya.

---

## F-REQ-002 — Permintaan jalan-jalan

### FR-REQ-006 — Syarat permintaan jalan-jalan

| | |
|---|---|
| Fitur | F-REQ-002 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `KS-4` |
| Aktor dan izin | Sama dengan FR-REQ-001. Tempat, peserta, dan nama butir rencana termasuk *rincian* yang dapat dihapus (06 §2.5) |
| Aturan acuan | 06 §2.4 (`request_terms` jalan-jalan), §4 pembuka, B01, B10 |
| Business rule | BR-REQ-001, BR-REQ-005 |
| Kedudukan | Dirancang 06. Perkiraan butir tidak membatasi apa pun: **[Usulan]** `product-analyst`. Tanggal lampau pada versi sesudah v1: 06 §4.3 "Guard tanggal dan batas bawah pada versi baru" baris B01 sampai B05 (**[Usulan]** `architect`, revisi kesembilan) |

Permintaan jalan-jalan adalah permintaan anggaran yang syaratnya juga memuat tanggal rencana (wajib), tempat, rentang waktu, peserta, dan butir rencana pembelian (nama, perkiraan). Satu state machine; tidak ada objek rencana tersendiri. Batas otorisasi adalah jumlah yang diminta, bukan jumlah perkiraan butir; jumlah perkiraan ditampilkan di sampingnya sebagai keterangan.

**Kriteria penerimaan**
- **AC-1** — *Given* hari ini 10 Desember 2026, *When* A mengajukan jalan-jalan bertanggal rencana 9 Desember 2026, *Then* server menolak; *When* bertanggal 14 Desember 2026, *Then* permintaan berada di `awaiting_approver`.
- **AC-2** — *Given* formulir jalan-jalan, *When* A mengirimnya tanpa tanggal rencana, *Then* server menolak.
- **AC-3** — *Given* A mengajukan jalan-jalan Rp500.000 dengan tiga butir berperkiraan Rp150.000, Rp200.000, dan Rp100.000, *When* B membukanya, *Then* keduanya melihat batas Rp500.000 dan jumlah perkiraan Rp450.000; bila disetujui, otorisasinya Rp500.000; perkiraan butir tidak pernah membatasi, menandai, atau mengubah realisasi.
- **AC-4** — *Given* jalan-jalan bertanggal rencana 14 Desember 2026 yang belum ditanggapi, *When* hari 14 Desember berakhir menurut zona waktu rumah tangga, *Then* permintaan `expired` (B10), keduanya menerima N-07, dan tidak ada akibat uang.
- **AC-5** — *Given* hari ini 10 Desember 2026 dan jalan-jalan bertanggal 14 Desember menunggu tanggapan, *When* B mengajukan usul balik dengan tanggal rencana 8 Desember 2026, *Then* server menolaknya karena tanggal itu sebelum hari ini, permintaan tetap `awaiting_approver` atas versi 1, dan tidak ada versi baru (06 §4.3, guard B01 sampai B05); *When* usul baliknya bertanggal 12 Desember 2026, *Then* versi 2 tersimpan dan state `awaiting_requester`.
- **AC-6** — *Given* hari ini 10 Desember 2026 dan B mengajukan usul balik bertanggal 12 Desember dengan `valid_until` 9 Desember 2026, *When* server menerimanya, *Then* usul balik ditolak karena `valid_until` sebelum hari ini (06 §4.3, guard B01 sampai B05).

**Dua arah.** Simetris; sama dengan FR-REQ-001.

---

## F-REQ-003 — Menanggapi dan merundingkan

### FR-REQ-007 — Menyetujui

| | |
|---|---|
| Fitur | F-REQ-003 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `KS-4` |
| Aktor dan izin | B06: pemberi persetujuan, saat `awaiting_approver`. B07: pemohon, saat `awaiting_requester`, atas usul balik pemberi persetujuan. Tidak ada yang menyetujui atas nama orang lain (PG-02) |
| Aturan acuan | 06 §3.3, §4.3 B06, B07, §12 S-1 sampai S-4, kasus uji §12; §16.4 (tombol kartu), INV-A44, INV-A46; 07 §3.8 (kait `request.approved`), INV-10, CTH-32 |
| Business rule | BR-REQ-003, BR-REQ-004, BR-REQ-034, BR-REQ-037 |
| Kedudukan | Dirancang 06. Menyetujui dari kartu chat: 06 §16.4 (**[Usulan]** `architect`, atas keputusan chat 12 November 2026) |

Menyetujui permintaan sederhana adalah satu ketukan, dari layar permintaan atau dari kartunya di chat; keduanya perintah B06 atau B07 yang sama. Hasilnya: state `approved`, `approved_version` terisi, otorisasi terbuka lahir (beserta reservasi bila ada tempat uang sumber), N-04 kepada pihak lainnya, dan satu baris peristiwa di chat. Pelaku tidak diberi tahu atas tindakannya sendiri.

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A Rp100.000 versi 1 di `awaiting_approver`, *When* B menyetujui dengan menyebut versi 1, *Then* state `approved`, otorisasi berbatas Rp100.000 terbuka, A menerima N-04 berprioritas `kabar`, dan B tidak menerima notifikasi atas tindakannya.
- **AC-2** — *Given* B mengajukan usul balik Rp80.000 (versi 2, `awaiting_requester`), *When* A menerimanya (B07), *Then* state `approved` dengan batas Rp80.000 dan B menerima N-04.
- **AC-3** — *Given* B offline, *When* B menyetujui, *Then* perangkat B menampilkan "menunggu terkirim", tidak pernah "disetujui", dan A tetap melihat `awaiting_approver` sampai server menerimanya (S-1).
- **AC-4** — *Given* persetujuan B terkirim dua kali karena jaringan buruk, *When* server menerima keduanya, *Then* hanya ada satu otorisasi (CTH-32).
- **AC-5** — *Given* B menyetujui pada saat yang sama A membatalkan, *When* server menerima keduanya, *Then* yang lebih dulu diterima menang dan tidak pernah ada otorisasi terbuka pada permintaan `cancelled` (INV-A08).
- **AC-6** — *Given* permintaan A Rp100.000 versi 1 tampil sebagai kartu di chat, *When* B mengetuk "setuju" pada kartu itu, *Then* server menerima B06 yang membawa versi 1 dan state `awaiting_approver` yang diharapkan; hasilnya, Peristiwa, dan notifikasinya identik dengan AC-1; chat mendapat satu baris peristiwa yang tidak memicu notifikasi tersendiri (06 §16.4, INV-A44, INV-A46).
- **AC-7** — *Given* kartu permintaan A, *When* B membalas kartu itu dengan pesan "setuju" atau dengan foto struk, *Then* permintaan tetap `awaiting_approver`, tidak ada otorisasi, dan pesan itu menjadi bagian diskusi permintaan (06 §16.3, INV-A44).
- **AC-8** — *Given* kartu permintaan A di chat, *When* A membuka kartunya sendiri, *Then* A tidak melihat tombol setuju; *When* perintah setuju dari A tetap terkirim, *Then* server menolaknya seperti dari layar permintaan (06 §16.4).

**Dua arah.** Simetris per peran. Pemberi persetujuan adalah "yang bukan pemohon"; ketika B pemohon, A yang menyetujui dengan aturan yang sama, dari layar objek maupun dari kartu.

### FR-REQ-008 — Menolak

| | |
|---|---|
| Fitur | F-REQ-003 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; arah 2 pemilik produk ("setuju atau tidak") |
| Aktor dan izin | B08 hanya oleh pemberi persetujuan, dari `awaiting_approver` atau `awaiting_requester`. Alasan, bila diisi, terbaca para pihak saja dan tidak pernah tampil di layar kunci |
| Aturan acuan | 06 §4.3 B08, §4.4; N-05; §16.4 (tombol "belum sepakat" pada kartu) |
| Business rule | BR-REQ-003, BR-REQ-015, BR-REQ-037 |
| Kedudukan | Alasan tidak wajib: Q-50, D-24 (lewat "lanjut"). "Belum sepakat" pada kartu = B08: usulan `architect` (06 P-25), diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-7 (i) |

Penolak ditawari menulis alasan singkat atau beralih ke usul balik. Tidak ada state "usul balik ditolak": pemohon yang tidak setuju dengan usul balik mengajukan usul balik lagi (B04) atau membatalkan (B09). Penolakan tidak menyentuh uang.

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A menunggu B, *When* B menolak tanpa alasan, *Then* state `rejected`, A menerima N-05 berprioritas `kabar`, dan tidak ada medan alasan yang wajib.
- **AC-2** — *Given* B menolak dengan alasan "minggu depan saja", *When* A membuka permintaannya, *Then* alasan itu terbaca oleh A dan B saja; notifikasi di layar kunci tetap teks umum.
- **AC-3** — *Given* B mengajukan usul balik (state `awaiting_requester`), *When* A mengirim perintah `reject`, *Then* server menolak perintah itu; pilihan A adalah menerima, mengajukan usul balik, atau membatalkan.
- **AC-4** — *Given* penolakan, *When* keduanya membuka anggaran kategori permintaan itu, *Then* tidak ada *terikat*, reservasi, atau entri yang lahir.
- **AC-5** — *Given* kartu permintaan A di chat, *When* B mengetuk "belum sepakat" dengan atau tanpa alasan, *Then* server menjalankan B08 seperti AC-1 dan AC-2, permintaan `rejected`, dan A dapat mengajukan lagi dalam rantai yang sama; *When* B hanya membalas kartu itu dengan pesan, *Then* permintaan tetap menunggu (06 §16.4; P-25).

**Dua arah.** Simetris per peran (AC-3 berlaku sama ketika B pemohon). Bahwa hanya pemberi persetujuan yang menolak adalah sifat peran, bukan orang.

### FR-REQ-009 — Usul balik dan revisi syarat

| | |
|---|---|
| Fitur | F-REQ-003 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` (anggaran, waktu, dan tambahan dapat dirundingkan) |
| Aktor dan izin | B02: pemohon merevisi saat `awaiting_approver`. B03: pemberi persetujuan mengusulkan balik. B04: pemohon mengusulkan balik. B05: pengusul versi berjalan merevisi saat `awaiting_requester`. Versi lama terbaca para pihak sebagai "digantikan" |
| Aturan acuan | 06 §4.3 B02 sampai B05, §4.4, "Guard tanggal dan batas bawah pada versi baru"; N-02, N-03; §10 jendela 6 jam; §16.4 (tombol "usul lain" dan revisi dari kartu) |
| Business rule | BR-REQ-004, BR-REQ-037 |
| Kedudukan | Keputusan pemilik nomor 5 (dapat dirundingkan). Guard tanggal pada versi baru: **[Usulan]** `architect` (06 §4.3, revisi kesembilan) |

Usul balik boleh mengubah syarat mana pun: jumlah, tanggal, tempat, butir rencana, tempat uang sumber, masa berlaku. Ia menggantikan versi sebelumnya seluruhnya. State berubah seketika; pemberitahuannya tunduk pada pembatas (FR-NTF-011).

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A Rp100.000 versi 1, *When* B mengajukan usul balik Rp80.000, *Then* versi 2 diusulkan B, state `awaiting_requester`, A menerima N-02, dan versi 1 tetap terbaca keduanya sebagai digantikan.
- **AC-2** — *Given* AC-1, *When* A mengajukan usul balik Rp90.000, *Then* versi 3 diusulkan A, state `awaiting_approver`, dan B dapat membaca versi 3 seketika di aplikasi walau pemberitahuannya masih ditahan jendela 6 jam.
- **AC-3** — *Given* state `awaiting_requester` dengan versi berjalan diusulkan B, *When* B merevisi jumlahnya dari Rp80.000 ke Rp85.000 (B05), *Then* versi baru diusulkan B dan N-03 menggantikan pemberitahuan versi lama.
- **AC-4** — *Given* AC-3, *When* A mengirim `revise_terms`, *Then* server menolak karena A bukan pengusul versi berjalan.
- **AC-5** — *Given* permintaan konfirmasi susulan, *When* siapa pun mengirim revisi atau usul balik, *Then* server menolak (FR-REQ-037).
- **AC-6** — *Given* kartu permintaan A Rp100.000 versi 1 di chat, *When* B memilih "usul lain" dari kartu dan mengirim Rp80.000, *Then* hasilnya identik dengan AC-1 (B03), kartu yang sama kini menampilkan versi 2 dan state `awaiting_requester` bagi keduanya, dan tidak ada kartu kedua (06 §16.2 "Isi kartu", §16.4).
- **AC-7** — *Given* hari ini 10 Desember 2026, *When* siapa pun mengirim revisi atau usul balik yang `planned_date` atau `valid_until`-nya sebelum hari ini, atau yang jumlahnya ≤ 0, *Then* server menolak dan versi berjalan tidak berubah (06 §4.3 guard B01 sampai B05; FR-REQ-006/AC-5).

**Dua arah.** Simetris per peran: setiap pihak punya satu langkah merevisi versinya sendiri dan satu langkah mengusulkan balik versi pihak lain.

### FR-REQ-010 — Tanggapan selalu atas versi yang dibaca

| | |
|---|---|
| Fitur | F-REQ-003 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` (persetujuan tidak pindah ke syarat yang belum dibaca) |
| Aktor dan izin | Semua penanggap |
| Aturan acuan | 06 S-3, S-4, S-11 (b), INV-A06, INV-A07, INV-A44; kasus uji §12 baris 1 dan baris kartu (revisi kesembilan); §16.2 "Isi kartu", §16.4 |
| Business rule | BR-REQ-004, BR-REQ-034, BR-REQ-037 |
| Kedudukan | Dirancang 06. Kartu basi: S-11 (b), **[Usulan]** `architect` |

Kartu di chat tidak punya state sendiri. Ia selalu menampilkan state dan versi syarat objek saat ini, dan setiap tombolnya membawa versi dan state yang dibaca pengetuknya. Kartu yang dibaca sebelum objeknya berubah adalah kartu basi; tanggapan dari kartu basi diperlakukan persis seperti tanggapan dari layar objek yang basi.

**Kriteria penerimaan**
- **AC-1** — *Given* B membaca versi 1 dan A merevisi ke versi 2, *When* revisi A diterima server lebih dulu dan persetujuan B atas versi 1 tiba sesudahnya, *Then* persetujuan ditolak `terms_superseded`, tidak ada otorisasi, dan B ditunjukkan versi 2.
- **AC-2** — *Given* situasi AC-1, *When* persetujuan B tiba lebih dulu, *Then* revisi A ditolak `state_changed`, permintaan `approved` atas versi 1, dan A dapat mengusulkan perubahan syarat (FR-REQ-012).
- **AC-3** — *Given* B menyetujui versi 1 saat offline dan A merevisi sebelum B tersambung, *When* perangkat B mengirim persetujuannya, *Then* ditolak `terms_superseded` dan perangkat B menyatakan bahwa persetujuan itu tidak terjadi serta menampilkan versi terbaru.
- **AC-4** — *Given* B offline mengetuk "setuju" pada kartu permintaan A versi 1, dan sementara itu A merevisi ke versi 2, *When* perangkat B tersambung dan perintahnya tiba, *Then* persetujuan ditolak `terms_superseded`, tidak ada otorisasi, kartu di perangkat B berganti dari "menunggu terkirim" ke versi 2, dan B diberi tahu bahwa persetujuannya tidak terjadi (06 S-11 (b); kasus uji §12 baris kartu).
- **AC-5** — *Given* kartu permintaan A yang sudah dibatalkan A (B09), *When* B mengetuk tombol apa pun pada kartu yang belum diperbarui di perangkatnya, *Then* server menolak `state_changed` dan kartu menampilkan state `cancelled` tanpa tombol tanggapan.

**Dua arah.** Simetris; aturan versi berlaku atas setiap penanggap, dari layar objek maupun dari kartu.

### FR-REQ-011 — Tunda senyap dan kabar "nanti aku jawab"

| | |
|---|---|
| Fitur | F-REQ-003 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` (tidak ditagih); `KS-5` |
| Aktor dan izin | Pihak yang ditunggu tanggapannya menunda satu permintaan. `reminders_paused_until` hanya terbaca penekannya (INV-A33). Kabar N-21 hanya dikirim atas perintah tersendiri pihak yang ditunggu |
| Aturan acuan | 06 §4.4 "Tunda senyap", §10 tabel batas baris "Tunda senyap", INV-A33 (dipersempit), AT-5 (c)(d), §16.5 "Jam tenang dan tunda", §16.9; N-21 |
| Business rule | BR-REQ-005, BR-REQ-018 |
| Kedudukan | Diterima atas saran lead (S2-5 pilihan A butir c); 24 jam **[Asumsi]**. Batas di chat (kartu yang dibaca terlihat): 06 AT-5 (d), akibat keputusan chat 12 November 2026 |

Tunda bukan jawaban. Selama tunda, pemberitahuan `perlu_tanggapan` atas rantai itu tidak dilepas kepada penekannya; "ingatkan lagi" dari pihak yang menunggu diterima, dihitung, dan ditahan. Tunda tidak berbatas jumlah dan tidak menghentikan kedaluwarsa. Kabar "nanti aku jawab" adalah kalimat baku berprioritas `ringkasan`, tidak menyebut tunda atau lamanya, dan bawaannya tidak dikirim.

**Batas yang dinyatakan terus terang (06 AT-5 d).** Tunda tetap senyap, tetapi chat punya tanda dibaca. Bila B membaca kartu permintaan di chat lalu menunda, A melihat kartu itu sudah dibaca. A tidak dapat mengetahui apakah B menunda, sampai kapan, atau apa setelan jam tenangnya. Tunda dan jam tenang tidak menggerakkan tanda apa pun. Penjelasan "siapa melihat apa" (FR-FAM-010) menyebut batas ini kepada keduanya sebelum chat dipakai.

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A menunggu B dan B menunda pada 08.00, *When* A menekan "ingatkan lagi" pada 10.00, *Then* perintah A diterima, hitungan A naik satu, A melihat waktu tekan 10.00, dan B tidak menerima pemberitahuan apa pun atas rantai itu sampai tunda berakhir pada 08.00 esoknya; saat itu B menerima satu pemberitahuan (Keluhan 5).
- **AC-2** — *Given* rangkaian perintah dan pembacaan A yang sama, dan tindakan B selain tunda (termasuk apakah dan kapan B membaca chat) juga sama, *When* dijalankan sekali dengan B menunda dan sekali tanpa tunda, *Then* semua tanggapan server dan semua yang terbaca A, termasuk tanda terkirim dan dibaca di chat, identik; tidak ada tanda, Peristiwa, atau notifikasi tentang tunda (INV-A33 sesudah revisi kesembilan).
- **AC-3** — *Given* jalan-jalan bertanggal 14 Desember 2026 yang ditunda B pada 14 Desember pukul 20.00, *When* hari itu berakhir, *Then* permintaan tetap `expired` dan keduanya menerima N-07; tunda tidak memperpanjangnya.
- **AC-4** — *Given* permintaan A menunggu B, *When* B memilih "nanti aku jawab", *Then* A menerima N-21 berprioritas `ringkasan` dengan kalimat baku yang tidak menyebut tunda atau lamanya; *When* B hanya menunda, *Then* N-21 tidak dikirim.
- **AC-5** — *Given* B menunda, *When* keduanya membuka permintaan, *Then* state tidak berubah dan tidak ada baris tanggapan yang tercatat.
- **AC-6** — *Given* B menunda permintaan A, *When* B kemudian membuka chat dan kartu permintaan itu tampil di layarnya, *Then* kartu itu bertanda dibaca bagi A seperti pesan lain, tanpa waktu baca; tidak ada pengingat, lencana, atau warna "dibaca tapi belum dijawab", dan tunda tetap berlaku (06 AT-5 (d), §16.5).

**Dua arah.** Simetris per peran. Pihak yang ditunggu dapat siapa pun: pemberi persetujuan saat `awaiting_approver`, pemohon saat `awaiting_requester`; keduanya punya tunda yang sama.

---

## F-REQ-004 — Perubahan syarat setelah disetujui

### FR-REQ-012 — Mengusulkan dan menanggapi perubahan syarat

| | |
|---|---|
| Fitur | F-REQ-004 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` |
| Aktor dan izin | B11: pemohon atau pemberi persetujuan mengusulkan. B12: pihak lainnya menerima. B13: pihak lainnya menolak, atau pengusulnya menarik |
| Aturan acuan | 06 §4.3 B11 sampai B13 dan "Guard tanggal dan batas bawah pada versi baru", §4.4, INV-A14, kasus uji §12 baris 3, §16.4 (kartu perubahan syarat); 07 §3.8 (`authorization.amended`), varian perubahan syarat CTH-22 |
| Business rule | BR-REQ-013, BR-REQ-007, BR-REQ-008, BR-REQ-010 |
| Kedudukan | Keputusan 21 (diputuskan pemilik produk sendiri, S2-3) |

Sampai pihak lain menerimanya, syarat yang disetujui tetap berlaku. Batas baru menggerakkan *terikat* dan reservasi. Bila batas baru lebih kecil daripada realisasi, kelebihan lahir dan N-08 dikirim. Perubahan yang masih menunggu saat permintaan ditutup gugur dan tercatat demikian.

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A Rp100.000 bersumber `TUN-S` disetujui B (CTH-22: `TUN-S` tercatat Rp425.000, reservasi Rp100.000, tersedia Rp325.000; anggaran terpakai Rp75.000, terikat Rp100.000, sisa Rp325.000), *When* A mengusulkan batas Rp150.000 dan B menerimanya, *Then* batas Rp150.000, reservasi Rp150.000, tersedia Rp275.000, terikat Rp150.000, sisa anggaran Rp275.000, dan tidak ada entri yang lahir.
- **AC-2** — *Given* A mengusulkan batas Rp150.000 atas otorisasi Rp100.000 dan B belum menjawab, *When* A menautkan belanja Rp120.000, *Then* batas yang berlaku tetap Rp100.000, permintaan menampilkan "melebihi Rp20.000", dan B menerima N-08 satu kali; *When* B kemudian menerima Rp150.000, *Then* tampilan kelebihan hilang dan terikat Rp30.000.
- **AC-3** — *Given* otorisasi terbuka, *When* A dan B mengusulkan perubahan pada saat bersamaan, *Then* yang pertama tersimpan dan yang kedua ditolak karena sudah ada yang menunggu.
- **AC-4** — *Given* otorisasi Rp100.000 dengan realisasi Rp75.000, *When* B mengusulkan batas Rp50.000 dan A menerimanya, *Then* permintaan menampilkan "melebihi Rp25.000" kepada keduanya dan N-08 dikirim kepada pemberi persetujuan (B).
- **AC-5** — *Given* perubahan syarat menunggu, *When* A menyelesaikan permintaan (B15), *Then* perubahan itu gugur dan tercatat gugur pada permintaan.
- **AC-6** — *Given* B mengusulkan perubahan, *When* B menariknya sebelum dijawab, *Then* syarat yang disetujui tetap dan A menerima N-05.
- **AC-7** — *Given* hari ini 20 Desember 2026 dan jalan-jalan A bertanggal rencana 14 Desember 2026 disetujui dengan batas Rp500.000, *When* A mengusulkan batas Rp600.000 tanpa menggeser tanggal, *Then* usulan diterima server; *When* A mengusulkan batas Rp600.000 dengan tanggal rencana digeser ke 15 Desember 2026, *Then* server menolak karena tanggal yang digeser sebelum hari ini; *When* tanggalnya digeser ke 22 Desember 2026, *Then* usulan diterima (06 §4.3 baris B11).
- **AC-8** — *Given* perubahan syarat yang diusulkan A tampil di kartu permintaan di chat, *When* B mengetuk "terima" atau "belum sepakat" pada kartu, *Then* server menjalankan B12 atau B13 yang sama dengan dari layar permintaan, dengan versi perubahan yang dibaca B; bila A sudah menarik perubahan itu, perintah B ditolak `state_changed` dan kartu menampilkan syarat yang berlaku (06 §16.4).

**Dua arah.** Simetris: kedua pihak dapat mengusulkan, dan selalu pihak lain yang menerima.

### FR-REQ-013 — Persetujuan tidak dapat dicabut sendirian

| | |
|---|---|
| Fitur | F-REQ-004 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` |
| Aktor dan izin | Tidak ada perintah pencabutan bagi siapa pun |
| Aturan acuan | 06 §4.4, PG-11, §5.1; §4.3 "Guard tanggal dan batas bawah pada versi baru" baris B11 |
| Business rule | BR-REQ-012, BR-REQ-013 |
| Kedudukan | Keputusan 21 (diputuskan pemilik produk sendiri). Batas bawah jumlah pada perubahan syarat: 06 §4.3 baris B11 (**[Usulan]** `architect`, revisi kesembilan) |

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A disetujui B, *When* B mencoba membawa permintaan itu ke `rejected` atau `cancelled`, *Then* server menolak; tidak ada perintah pencabutan.
- **AC-2** — *Given* B berubah pikiran, *When* B membuka permintaan itu, *Then* yang tersedia baginya adalah diskusi dan usulan perubahan syarat; persetujuan tetap berlaku sampai A menerima perubahan itu.
- **AC-3** — *Given* A sudah menautkan belanja Rp75.000, *When* B menulis di diskusi bahwa ia keberatan, *Then* otorisasi, realisasi, dan entrinya tidak berubah.
- **AC-4** — *Given* B ingin menghentikan otorisasi sepenuhnya, *When* B mengusulkan batas Rp0, *Then* server menolak usulan itu karena batas baru harus > 0 (06 §4.3 baris B11); menghentikan otorisasi sepenuhnya hanya lewat pembatalan oleh pemohon (B17) bila realisasi 0, atau penyelesaian oleh pemohon (B15) (06 §5.1, PG-11).

**Dua arah.** Simetris; tidak ada pihak yang punya pencabutan.

---

## F-REQ-005 — Belanja tertaut dan penyelesaian sisa

### FR-REQ-014 — Menautkan belanja saat mencatat

| | |
|---|---|
| Fitur | F-REQ-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`, `MSL-2`, `MSL-5` (satu pilihan saat mencatat) |
| Aktor dan izin | Anggota dewasa yang mencatat entri, termasuk pemberi persetujuan yang membayar dari tempat uang sumber (situasi D) |
| Aturan acuan | 06 §4.3 B14, §5.2; 07 §3.8 "Penautan entri ke otorisasi", tabel *jumlah pemakaian*, INV-23, INV-26, CTH-21, CTH-31 |
| Business rule | BR-REQ-007, BR-REQ-011, BR-REQ-024 |
| Kedudukan | Dirancang 06 dan 07 |

Saat mencatat, pencatat dapat memilih satu otorisasi terbuka, atau permintaan yang sudah ditutup bila tanggal entrinya memenuhi rentang FR-REQ-015. Satu otorisasi boleh menerima banyak entri pada hari berbeda; state tetap `approved` ("sebagian terpakai" bukan state). Realisasi adalah jumlah *jumlah pemakaian* entri tertaut menurut 07; serah terima itu sendiri bukan realisasi.

**Kriteria penerimaan**
- **AC-1** — *Given* otorisasi A Rp100.000 berkategori "jajan dan makan di luar", *When* A mencatat belanja Rp75.000 dari `TUN-S` yang dipegang A dan memilih permintaan itu, *Then* realisasi Rp75.000, terikat Rp25.000, sisa Rp25.000 tampil kepada keduanya, dan entri itu tidak bertanda (N = 0).
- **AC-2** — *Given* otorisasi yang sama, *When* A mencatat Rp40.000 lalu Rp35.000 pada hari berbeda, *Then* state tetap `approved` dan realisasi Rp75.000.
- **AC-3** — *Given* otorisasi A, *When* B mencatat pembayaran Rp75.000 dari tempat uang dana bersama yang dipegang B dan menautkannya, *Then* tautan diterima dan realisasi Rp75.000.
- **AC-4** — *Given* dua entri Rp40.000 dan Rp35.000 ditautkan bersamaan dari dua perangkat, *When* server menerima keduanya, *Then* keduanya tersimpan dan realisasi Rp75.000 (CTH-31).
- **AC-5** — *Given* serah terima Rp100.000 yang ditautkan ke permintaan, *When* belum ada belanja, *Then* realisasi Rp0.
- **AC-6** — *Given* entri yang sudah tertaut ke satu otorisasi, *When* pencatat mencoba menautkannya ke otorisasi lain, *Then* ditolak; satu entri tertaut ke paling banyak satu otorisasi.

**Dua arah.** Simetris. Siapa pun dari keduanya boleh mencatat dan menautkan saat mencatat; tidak ada yang membaca siapa pemegang tempat uang.

### FR-REQ-015 — Tautan susulan dan pencabutannya

| | |
|---|---|
| Fitur | F-REQ-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5` (pencatat pelupa tidak dihukum oleh tenggat); `PAGAR` (pasangan tidak dapat memasang tanda pada entri orang lain) |
| Aktor dan izin | Menautkan ke otorisasi: pencatat entri atau pemohon otorisasi. Mencabut: pembuat tautan atau pencatat entri. Menautkan ke tagihan berulang: hanya pencatat entri. Pelaku tampil pada permintaan dan entri |
| Aturan acuan | 06 §4.5 baris "Entri terlambat setelah ditutup", "Tautan yang tidak sah", "Tautan susulan", "Pembayaran atas kewajiban yang sudah tertaut penuh"; INV-A34; 07 §3.8 butir 3, INV-26; CTH-34 |
| Business rule | BR-REQ-014, BR-REQ-011, BR-REQ-028 |
| Kedudukan | D-39 (lewat "lanjut"). Tautan ke tagihan berulang hanya oleh pencatat: **[Usulan]** `architect`, belum disajikan |

**Kriteria penerimaan**
- **AC-1** — *Given* entri Rp75.000 dicatat A tanpa tautan dan otorisasi terbuka milik A sebagai pemohon, *When* B (bukan pencatat dan bukan pemohon) mencoba menautkannya, *Then* ditolak; *When* A menautkannya, *Then* diterima dan pelakunya tampil.
- **AC-2** — *Given* entri dicatat B dan otorisasi milik A sebagai pemohon, *When* A menautkan entri B, *Then* diterima; *When* kemudian B mencabutnya, *Then* diterima (B pencatat); pencabutan oleh anggota yang bukan pembuat tautan dan bukan pencatat ditolak.
- **AC-3** — *Given* permintaan disetujui 1 Desember 2026 dan diselesaikan 10 Desember 2026, *When* pada 15 Desember A mencatat belanja Rp30.000 bertanggal 8 Desember dan menautkannya, *Then* tautan diterima, realisasi dihitung ulang, *terikat* tetap Rp0, state tidak berubah, dan entri bertanda "dicatat setelah ditutup"; *When* entri bertanggal 12 Desember, *Then* tautan ditolak dan entrinya tersimpan tanpa tautan.
- **AC-4** — *Given* permintaan `cancelled`, *When* siapa pun menautkan entri kepadanya, *Then* tautan ditolak dan entrinya tetap tersimpan tanpa tautan; tanda dihitung seperti atas entri tanpa permintaan.
- **AC-5** — *Given* pembelian aset secara cicilan Rp4.360.000 yang sudah tertaut penuh (CTH-34), *When* pembayaran cicilan dicatat dan pencatat mencoba menautkannya ke permintaan itu, *Then* tautan ditolak dan pembayarannya tersimpan.
- **AC-6** — *Given* entri dicatat A, *When* B mencoba menautkannya ke tagihan berulang, *Then* ditolak; hanya A yang dapat.

**Dua arah.** Simetris. Wewenang mengikuti hubungan dengan objek (pencatat, pemohon, pembuat tautan), bukan orangnya; menukar A dan B memberi hasil yang tertukar.

### FR-REQ-016 — Kelebihan atas batas

| | |
|---|---|
| Fitur | F-REQ-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `MSL-5` (belanja nyata tetap tercatat penuh) |
| Aktor dan izin | Tidak ada tanggapan yang diperlukan dari siapa pun. Pemohon boleh mengusulkan kenaikan batas (B11) |
| Aturan acuan | 06 §5.4, B14 (`authorization.overrun`), N-08, kasus uji §12 rumus `N`; 07 §3.8 "Tanda atas entri", CTH-24 |
| Business rule | BR-REQ-010, BR-REQ-011, BR-REQ-024 |
| Kedudukan | D-07 (lewat "lanjut"; melunakkan default S1) |

**Kriteria penerimaan**
- **AC-1** — *Given* otorisasi B Rp200.000 berkategori "hadiah" (CTH-24), *When* B mencatat dan menautkan pembayaran Rp230.000, *Then* entri tersimpan penuh, realisasi Rp230.000, terikat Rp0, permintaan menampilkan "melebihi Rp30.000" kepada keduanya, A menerima N-08 satu kali, tidak ada tanggapan yang diminta, dan entri itu tidak bertanda karena N = Rp30.000 tidak di atas Rp500.000.
- **AC-2** — *Given* otorisasi yang sama, *When* yang dibayar Rp800.000, *Then* N = Rp600.000, entri bertanda dasar ambang, dan N-08 dikirim satu kali.
- **AC-3** — *Given* ambang Rp500.000 dan otorisasi Rp1.000.000, *When* entri tiba berurutan Rp700.000, Rp400.000, Rp600.000, *Then* yang pertama tidak bertanda (N = 0); yang kedua tidak bertanda (N = Rp100.000) tetapi memicu kelebihan dan N-08; yang ketiga bertanda dasar ambang (N = Rp600.000); N-08 tidak dikirim lagi.
- **AC-4** — *Given* kelebihan sudah terjadi, *When* pemohon tidak mengusulkan apa pun, *Then* tidak ada akibat lain dan penyelesaian tidak menunggu apa pun.

**Dua arah.** Simetris. Penerima N-08 adalah pemberi persetujuan, siapa pun dia.

### FR-REQ-017 — Menyelesaikan dan sisa menurut situasi custody

| | |
|---|---|
| Fitur | F-REQ-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`, `MSL-2`, `MSL-3`; alur acuan brief §5D (100.000 / 75.000 / 25.000) |
| Aktor dan izin | B15 hanya oleh pemohon, bila realisasi > 0. Pemberi persetujuan tidak dapat menyelesaikan (06 §5.1) |
| Aturan acuan | 06 §5.1 sampai §5.3; 07 §3.8 tabel sisa A sampai D, T-02; CTH-21, CTH-22, CTH-23 |
| Business rule | BR-REQ-009, BR-REQ-012 |
| Kedudukan | D-07, D-14 (lewat "lanjut"); bawaan T-02 |

Menyelesaikan adalah satu ketukan. Pemohon tidak ditanya "situasi yang mana"; server membacanya dari entri tertaut. Hanya situasi B dengan sisa > 0 yang memunculkan satu pertanyaan: kembalikan sisa, atau tetap dipegang. Pengembalian adalah serah terima biasa yang tertaut dan mengikuti konfirmasi custody. Sisa tidak pernah menjadi pemasukan atau dana pribadi. N-09 (`ringkasan`) kepada pihak lainnya.

**Kriteria penerimaan**
- **AC-1** — *Given* situasi A (CTH-21): otorisasi Rp100.000, belanja Rp75.000 dari `TUN-S` yang dipegang pemohon, *When* pemohon menyelesaikan, *Then* tidak ada pertanyaan, sisa Rp25.000 tetap di `TUN-S` sebagai dana bersama, terikat Rp0, sisa anggaran Rp425.000, dan pihak lain menerima N-09 di ringkasan.
- **AC-2** — *Given* situasi B (CTH-22): `TUN-S` dipegang B, A menerima serah terima Rp100.000 ke `TUN-I` yang dipegang A (Rp3.340.000 menjadi Rp3.440.000) dan membelanjakan Rp75.000 (Rp3.365.000), *When* A menyelesaikan dan memilih mengembalikan, *Then* lahir serah terima Rp25.000 dari `TUN-I` ke `TUN-S` yang tertaut dan menunggu konfirmasi B; `TUN-S` Rp350.000, `TUN-I` Rp3.340.000.
- **AC-3** — *Given* AC-2, *When* A memilih tetap memegang atau tidak menjawab pertanyaan itu, *Then* tidak ada entri; `TUN-S` Rp325.000, `TUN-I` Rp3.365.000; sisa Rp25.000 tampil sebagai dana bersama yang dipegang A pada tampilan custody kedua pihak.
- **AC-4** — *Given* situasi C (CTH-23): pemohon membayar Rp75.000 dari dana pribadinya atas otorisasi Rp100.000, *When* ia menyelesaikan, *Then* tidak ada pertanyaan; hak ganti Rp75.000 sudah lahir saat talangan dicatat, bukan Rp100.000; penggantiannya entri tersendiri.
- **AC-5** — *Given* permintaan A disetujui, *When* B mengirim perintah `settle`, *Then* server menolak; B tetap dapat mencatat belanja tertaut dan mengusulkan perubahan syarat.
- **AC-6** — *Given* realisasi Rp0, *When* pemohon membuka permintaan, *Then* menyelesaikan tidak tersedia dan yang tersedia adalah membatalkan (FR-REQ-021).
- **AC-7** — *Given* pemohon menyelesaikan pada saat `system` menutup otomatis, *When* keduanya tiba, *Then* tepat satu `authorization.settled` (INV-A15).

**Dua arah.** Simetris per peran. CTH-21 dan CTH-23 memakai pemohon yang satu, CTH-22 pemohon yang lain; hasilnya ditentukan entri, bukan orangnya.

### FR-REQ-018 — Penutupan otomatis dan kedaluwarsa otorisasi

| | |
|---|---|
| Fitur | F-REQ-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5` (nol usaha bila lupa menyelesaikan); `MSL-4` |
| Aktor dan izin | `system`; ini bukan keputusan atas nama siapa pun |
| Aturan acuan | 06 §4.3 B16, B18, §4.5 "Saat masa berlaku habis", S-6; 07 §3.8, T-02 |
| Business rule | BR-REQ-006, BR-REQ-009 |
| Kedudukan | Keputusan 23 (sendiri); angka **[Asumsi]** |

**Kriteria penerimaan**
- **AC-1** — *Given* otorisasi tanpa tanggal rencana disetujui 3 Desember 2026 (berlaku sampai 2 Januari 2027) dengan realisasi Rp75.000, *When* hari 2 Januari 2027 berakhir, *Then* permintaan `settled` bertanda "ditutup otomatis", terikat Rp0, sisa tetap di tempatnya (T-02), dan keduanya menerima N-09 berprioritas `ringkasan`.
- **AC-2** — *Given* otorisasi yang sama dengan realisasi Rp0, *When* masa berlakunya habis, *Then* permintaan `expired`, terikat Rp0, reservasi dilepas, dan keduanya menerima N-07.
- **AC-3** — *Given* situasi B dengan sisa Rp25.000 di tangan pemohon, *When* ditutup otomatis, *Then* tidak ada entri; sisa tetap dipegang pemohon sebagai dana bersama.
- **AC-4** — *Given* jam perangkat salah, *When* perangkat menampilkan permintaan, *Then* penutupan mengikuti waktu server saja.

**Dua arah.** Simetris; `system` tidak membaca siapa pemohonnya.

### FR-REQ-019 — Tampilan yang sama bagi kedua pihak

| | |
|---|---|
| Fitur | F-REQ-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1`, `MSL-2`; `KS-1` (tahu tanpa bertanya) |
| Aktor dan izin | Para pihak permintaan (pilot: kedua anggota dewasa). Akibat uangnya terbaca semua yang melihat anggaran atau tempat uang itu (V-3) |
| Aturan acuan | 06 §5.6, §9, INV-A16, V-2, V-3 |
| Business rule | BR-REQ-031, BR-REQ-032 |
| Kedudukan | Dirancang 06; D-25 |

Yang terlihat identik bagi pemohon dan pemberi persetujuan: syarat yang disetujui; batas, realisasi, sisa atau kelebihan; entri tertaut beserta pencatat dan pembayarnya; di mana sisa berada; perubahan syarat yang menunggu; siapa yang menutup dan kapan, termasuk "ditutup otomatis". Yang berbeda hanya tindakan yang tersedia menurut peran.

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan disetujui dengan dua entri tertaut, *When* A dan B membukanya, *Then* angka, entri, pencatat, pembayar, dan riwayatnya sama persis.
- **AC-2** — *Given* anggota rumah tangga yang bukan pihak permintaan (kasus uji; tidak ada di pilot), *When* ia mencari permintaan itu, *Then* tidak ditemukan, tetapi *terikat* pada anggaran dan reservasi pada tempat uang tetap terlihat olehnya bila ia melihat anggaran atau tempat uang itu.
- **AC-3** — *Given* permintaan apa pun, *When* para pihak membukanya di layar permintaan, *Then* tidak ada tanda "sudah dibaca", waktu dibuka, atau lama menanggapi per orang; membuka layar permintaan atau notifikasinya tidak menggerakkan tanda dibaca kartunya di chat. Tanda terkirim dan dibaca hanya ada pada kartu di chat, tanpa waktu baca (06 AT-5 (a)(b), §16.5).
- **AC-4** — *Given* permintaan yang punya kartu di chat, *When* A dan B membuka kartu dan layar permintaannya, *Then* syarat, state, dan angka pada kartu sama dengan pada layar permintaan, karena isi kartu dibaca dari state objek saat ini (06 §16.2).

**Dua arah.** Simetris; satu tampilan bagi kedua peran.

### FR-REQ-020 — Mencatat offline dengan tautan

| | |
|---|---|
| Fitur | F-REQ-005 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5` (belanja di pasar tanpa sinyal selalu tercatat) |
| Aktor dan izin | Pencatat |
| Aturan acuan | 06 S-7, S-8, kasus uji §12 baris "Belanja dicatat offline"; §4.5; 07 INV-26 |
| Business rule | BR-REQ-011, BR-REQ-014, BR-REQ-024 |
| Kedudukan | Dirancang 06 |

**Kriteria penerimaan**
- **AC-1** — *Given* A offline mencatat Rp75.000 tertaut ke otorisasi terbuka, *When* tersambung dan otorisasi masih terbuka, *Then* entri dan tautannya diterima; sebelum itu perangkat menampilkan entri "menunggu terkirim" tanpa tanda.
- **AC-2** — *Given* B offline mencatat Rp600.000 tertaut ke permintaan A, dan sementara itu A membatalkan permintaan itu (realisasi masih 0), *When* perangkat B tersambung, *Then* entri tersimpan tanpa tautan, tanda dihitung server atasnya (bertanda dasar ambang karena Rp600.000 > Rp500.000), dan hasil sinkronisasi memberi tahu B bahwa tautannya tidak berlaku.
- **AC-3** — *Given* entri offline bertanggal di dalam rentang permintaan yang sementara itu ditutup otomatis, *When* tersambung, *Then* tautan diterima dan entri bertanda "dicatat setelah ditutup".
- **AC-4** — *Given* entri offline, *When* belum tersambung, *Then* tidak ada perhitungan tanda di perangkat.

**Dua arah.** Simetris; aturan sinkronisasi tidak membaca pelaku.

---

## F-REQ-006 — Batal, kedaluwarsa, ajukan lagi

### FR-REQ-021 — Membatalkan

| | |
|---|---|
| Fitur | F-REQ-006 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` |
| Aktor dan izin | Hanya pemohon. B09 dari state menunggu; B17 dari `approved` bila realisasi 0 |
| Aturan acuan | 06 §4.3 B09, B17; N-06; 07 §3.8 (`authorization.cancelled`) |
| Business rule | BR-REQ-012, BR-REQ-033 |
| Kedudukan | Dirancang 06 |

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A menunggu, *When* A membatalkan, *Then* state `cancelled`, B menerima N-06 berprioritas `kabar`, dan tidak ada akibat uang.
- **AC-2** — *Given* permintaan A Rp100.000 disetujui tanpa entri tertaut, *When* A membatalkan, *Then* otorisasi `cancelled`, terikat pada anggaran turun dari Rp100.000 ke Rp0, dan reservasi dilepas.
- **AC-3** — *Given* realisasi Rp75.000, *When* A mencoba membatalkan, *Then* ditolak dan yang ditawarkan adalah menyelesaikan.
- **AC-4** — *Given* permintaan A, *When* B mengirim `cancel`, *Then* server menolak.

**Dua arah.** Simetris per peran.

### FR-REQ-022 — Kedaluwarsa permintaan yang menunggu

| | |
|---|---|
| Fitur | F-REQ-006 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` (diam bukan alat) |
| Aktor dan izin | `system`, hanya bila versi berjalan bertanggal |
| Aturan acuan | 06 §4.3 B10, §4.5, PG-04, PG-05; N-07 |
| Business rule | BR-REQ-005 |
| Kedudukan | Keputusan pemilik nomor 6; D-20 |

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A tanpa tanggal yang tidak ditanggapi 60 hari, *When* waktu berjalan, *Then* ia tetap menunggu dan tampil di daftar menunggu dengan sejak kapan (Keluhan 6).
- **AC-2** — *Given* permintaan A dengan masa berlaku 20 Desember 2026 yang diisi A, *When* hari 20 Desember berakhir tanpa tanggapan, *Then* `expired` dan keduanya menerima N-07.
- **AC-3** — *Given* permintaan kedaluwarsa, *When* keduanya membukanya, *Then* ia tampil kedaluwarsa, bukan ditolak, tanpa alasan, dan dapat diajukan lagi.
- **AC-4** — *Given* permintaan apa pun yang menunggu, *When* tidak ada tanggapan, *Then* tidak pernah ada transisi ke `approved` atau `rejected` (Keluhan 6).

**Dua arah.** Simetris; kedaluwarsa hanya membaca tanggal.

### FR-REQ-023 — Mengajukan lagi dan rantai permintaan

| | |
|---|---|
| Fitur | F-REQ-006 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4` |
| Aktor dan izin | Pemohon asal, atas permintaan `rejected`, `cancelled`, atau `expired`. Kaitan terbaca kedua pihak |
| Aturan acuan | 06 §4.4 (`resubmitted_from`, rantai permintaan), §10 (hitungan per rantai), §10 "Risiko yang tersisa" |
| Business rule | BR-REQ-015, BR-REQ-016, BR-REQ-017 |
| Kedudukan | D-24 (lewat "lanjut"); jumlah pengajuan ulang tidak dibatasi di pilot (P-03) |

**Kriteria penerimaan**
- **AC-1** — *Given* B menolak permintaan A Rp1.000.000, *When* A memakai "ajukan lagi" dengan Rp700.000, *Then* lahir permintaan baru versi 1 yang menunjuk asalnya, kaitannya terlihat keduanya, dan N-01 kepada B tunduk pada jendela 6 jam rantai itu.
- **AC-2** — *Given* A sudah memakai tiga pengingat ulang pada rantai itu, *When* A mengajukan lagi lalu menekan "ingatkan lagi", *Then* ditolak; hitungan diwarisi.
- **AC-3** — *Given* pengingat otomatis rantai itu sudah terpakai, *When* A mengajukan lagi, *Then* tidak ada pengingat otomatis baru.
- **AC-4** — *Given* konfirmasi susulan yang ditolak, *When* pencatat mengajukannya lagi, *Then* ditolak (FR-REQ-039).

**Dua arah.** Simetris; rantai dan hitungannya per pengirim, siapa pun pengirimnya.

---

## F-REQ-007 — Pengingat atas permintaan

### FR-REQ-024 — Pengingat ulang oleh pihak yang menunggu

| | |
|---|---|
| Fitur | F-REQ-007 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `MSL-4`; arah 2 pemilik produk ("mengingat dan mencatat") |
| Aktor dan izin | Pihak yang menunggu tanggapan atas objek: pemohon saat `awaiting_approver`, pemberi persetujuan saat `awaiting_requester`, pengusul perubahan syarat atau kesepakatan, pencatat serah terima yang belum dikonfirmasi |
| Aturan acuan | 06 §10 tabel batas, PG-06, INV-A30 sampai INV-A33; N-20 |
| Business rule | BR-REQ-016 |
| Kedudukan | 6 jam: keputusan pemilik nomor 13 (sendiri). Tiga dan hitungan seumur rantai: diterima atas saran lead (S2-5), **[Asumsi]**. Isi dan keterbukaan: D-29 (lewat "lanjut") |

Perintah "ingatkan lagi" diterima atau ditolak hanya menurut riwayat pengirimnya sendiri: waktu ia terakhir mengirim pengingat ulang atas rantai itu, dan hitungannya. Pelepasan kepada penerima diatur pembatas `NTF` (jendela bersama, jam tenang, tunda, plafon). Isinya kalimat baku tanpa jumlah dan tanpa teks bebas.

**Kriteria penerimaan**
- **AC-1** — *Given* A mengajukan permintaan pukul 08.00 dan belum pernah mengingatkan, *When* A menekan "ingatkan lagi" pukul 09.00, *Then* perintah diterima, hitungan A menjadi 1, dan pemberitahuannya baru dilepas kepada B pukul 14.00 (6 jam sesudah N-01 dilepas).
- **AC-2** — *Given* AC-1, *When* A menekan lagi pukul 12.00, *Then* ditolak karena belum 6 jam sejak tekanan terakhirnya; A dapat melihat bahwa tekan berikutnya tersedia pukul 15.00.
- **AC-3** — *Given* hitungan A sudah 3 atas rantai itu, *When* A menekan lagi, *Then* ditolak; permintaan tetap menunggu dan tetap terlihat (Keluhan 5).
- **AC-4** — *Given* A sudah mengingatkan dua kali, *When* A atau B membuka permintaan, *Then* keduanya melihat jumlah 2 dan waktu A menekan; tidak seorang pun melihat kapan pengingat dilepas, sampai, ditahan, atau dibuka.
- **AC-5** — *Given* pengingat ulang dilepas, *When* tampil di perangkat B, *Then* isinya kalimat baku aplikasi tanpa jumlah uang dan tanpa teks A; layar kunci menampilkan teks umum (Keluhan 5).
- **AC-6** — *Given* B mengajukan usul balik dan menunggu jawaban A, *When* B menekan "ingatkan lagi", *Then* batas yang sama berlaku bagi B.

**Dua arah.** Ditulis untuk kedua arah: AC-1 sampai AC-5 dengan A menunggu, AC-6 dengan B menunggu; batasnya per pengirim, sama bagi keduanya.

### FR-REQ-025 — Pengingat otomatis

| | |
|---|---|
| Fitur | F-REQ-007 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `MSL-4`; arah 2 pemilik produk (aplikasi yang mengingatkan, bukan pasangan yang menagih) |
| Aktor dan izin | `system`, kepada pihak yang saat itu ditunggu; dihitung sebagai berasal dari pihak yang menunggu untuk jendela dan plafon |
| Aturan acuan | 06 §10 baris "Pengingat otomatis", "Keterbukaan"; §11.1 baris "Waktu pengingat yang berpatokan tanggal"; N-20 |
| Business rule | BR-REQ-017 |
| Kedudukan | D-20 (lewat "lanjut"); angka **[Asumsi]**. Saat pemicu bagi tanggal tanpa jam: 06 §11.1 (**[Usulan]** `architect`, revisi kesembilan) |

Paling banyak dua per rantai: 24 jam sesudah permintaan pertama rantai itu diajukan, dan sehari sebelum tanggalnya bila versi berjalan bertanggal. Karena tanggal rencana tidak berjam, pengingat "sehari sebelum" dijadwalkan pukul 00.00 hari sebelum tanggal itu menurut zona waktu rumah tangga, lalu tunduk pada jam tenang penerima (06 §11.1). Usul balik, revisi, perubahan syarat, dan pengajuan ulang tidak memicunya lagi. Ia gugur, tidak ditunda, bila jendela pengirim belum lewat atau ada yang tertahan atas objek itu. Ia tidak mengurangi batas tiga.

**Kriteria penerimaan**
- **AC-1** — *Given* A mengajukan permintaan tanpa tanggal pada 1 Desember pukul 08.00 dan tidak ada pemberitahuan `perlu_tanggapan` lain dari A atas rantai itu sesudahnya, *When* 2 Desember pukul 08.00 tiba, *Then* B menerima satu pengingat otomatis dan hitungan pengingat ulang A tetap.
- **AC-2** — *Given* A menekan "ingatkan lagi" pada 2 Desember pukul 05.00, ditahan jam tenang B dan dilepas pukul 06.00, *When* pengingat otomatis jatuh pukul 08.00, *Then* ia gugur karena jendela 6 jam sejak 06.00 belum lewat, dan tidak dijadwalkan ulang.
- **AC-3** — *Given* pengingat otomatis pertama sudah terpakai, *When* B mengajukan usul balik dan A mengajukan usul balik lagi, *Then* tidak ada pengingat otomatis baru selain yang bertanggal.
- **AC-4** — *Given* konfirmasi susulan, *When* waktu berjalan, *Then* tidak pernah ada pengingat otomatis.
- **AC-5** — *Given* pengingat otomatis dilepas atau gugur, *When* A membuka permintaan, *Then* A tidak melihat per kejadian apakah ia dilepas atau gugur; yang terbaca hanya aturannya.
- **AC-6** — *Given* jalan-jalan A bertanggal rencana 14 Desember 2026 menunggu B, pengingat otomatis pertamanya sudah terpakai, B memakai jam tenang bawaan 22.00 sampai 06.00, dan tidak ada pemberitahuan `perlu_tanggapan` lain atas rantai itu sesudah 12 Desember pukul 17.00, *When* 13 Desember pukul 00.00 tiba menurut zona waktu rumah tangga, *Then* pengingat otomatis kedua dijadwalkan, ditahan jam tenang, dan dilepas kepada B pukul 06.00 tanggal 13 Desember (06 §11.1).

**Dua arah.** Simetris; pengingat ditujukan kepada siapa pun yang ditunggu, dan dihitung atas nama siapa pun yang menunggu.

---

## F-REQ-008 — Kesepakatan rumah tangga

### FR-REQ-026 — Mengusulkan kesepakatan

| | |
|---|---|
| Fitur | F-REQ-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`, `MSL-3`; `PAGAR` |
| Aktor dan izin | Anggota dewasa mana pun (K01), kecuali jatah yang tidak di atas ambang atau 0: hanya pemiliknya. Kesepakatan terbaca semua anggota dewasa (06 §9.2) |
| Aturan acuan | 06 §6.3 K01, "Pos rutin", "Jatah pribadi", INV-A19, kasus uji §12 "Guard K01"; 07 §6.1, §3.15; N-11 |
| Business rule | BR-REQ-019 sampai BR-REQ-023 |
| Kedudukan | Ambang: keputusan 5. Hari mulai: D-49. Pos rutin: keputusan 22. Jatah tidak di atas ambang: **[Usulan]** SP-32/SP-34 (tidak dibantah); "sama dengan": pengetatan lead **[Usulan]**. Alasan tidak wajib: diterima atas saran lead (S2-1 v). N-11 pos rutin satu per hari: **[Usulan]** SP-36. Daftar sumber penghasilan lewat `rule`: **[Usulan]** 06 (Q-95) |

Jenis `rule`: ambang > 0, hari mulai 1 sampai 28, zona waktu dari daftar zona waktu yang dikenal server (revisi kesembilan; FR-FAM-011), daftar sumber penghasilan tidak kosong; setiap versi memuat seluruh isi. Jenis `routine`: pos rutin (jumlah per kali > 0) atau jatah pribadi (per periode ≥ 0). Bila sudah ada yang berlaku atas sasaran yang sama, usulan menunjuknya sebagai yang diganti. Satu usulan menunggu per sasaran. Alasan tidak pernah wajib; pada jatah, bila diisi, dari daftar tertutup; penjelasan bebas di diskusi. Membuat tagihan berulang dan mengusulkan kesepakatannya satu perintah; tagihannya dapat dipakai seketika.

**Kriteria penerimaan**
- **AC-1** — *Given* A dan B, *When* A mengusulkan ambang Rp0, *Then* ditolak; *When* Rp500.000, *Then* usulan `diusulkan` dan B menerima N-11.
- **AC-2** — *Given* usulan hari mulai, *When* nilainya 29, *Then* ditolak; *When* 25, *Then* diterima sebagai usulan.
- **AC-3** — *Given* ambang Rp500.000 berlaku, *When* B mengusulkan jatah A sebesar Rp300.000, Rp0, atau Rp500.000, *Then* ditolak; *When* Rp500.001, *Then* diterima sebagai usulan; *When* A sendiri mengusulkan Rp300.000, Rp500.000, atau Rp0, *Then* diterima sebagai usulan yang berlaku bila B menyetujuinya.
- **AC-4** — *Given* usulan ambang menunggu, *When* B mengusulkan ambang lain, *Then* ditolak karena sudah ada usulan menunggu atas sasaran itu.
- **AC-5** — *Given* A membuat tagihan berulang "listrik" Rp400.000 per kali, *When* sebelum B menyetujui kesepakatannya A mencatat dan menautkan pembayaran Rp450.000, *Then* pembayaran dinilai dengan aturan ambang per entri dan tidak bertanda (Rp450.000 tidak di atas Rp500.000).
- **AC-6** — *Given* A mengusulkan tiga pos rutin pada satu hari, *When* B menerima pemberitahuannya, *Then* hanya satu N-11 langsung; dua lainnya masuk ringkasan B.
- **AC-7** — *Given* A mengusulkan kenaikan jatahnya sendiri, *When* A mengisi formulir, *Then* alasan tidak wajib, bila diisi dipilih dari daftar tertutup, dan tidak ada medan teks bebas pada usulan.

**Dua arah.** Simetris, kecuali guard jatah yang sengaja membedakan "pemilik jatah" dari "pasangannya"; itu peran terhadap objek, bukan orangnya, dan berlaku sama atas jatah A maupun jatah B (06 PG-14).

### FR-REQ-027 — Menyetujui, menolak, menarik, dan menyetujui beberapa pos sekaligus

| | |
|---|---|
| Fitur | F-REQ-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` |
| Aktor dan izin | K02 dan K03: setiap pihak yang disyaratkan, yaitu anggota dewasa penuh yang bukan pengusul menurut `required_parties` (pada dua anggota dewasa: pasangan pengusul). K04: pengusul |
| Aturan acuan | 06 §6.3 K02 sampai K04, "Pihak yang disyaratkan" (revisi kedua belas; P-37), "Pos rutin" (menyetujui sekaligus), INV-A18; N-12; §16.4 baris "Usulan kesepakatan" |
| Business rule | BR-REQ-019, BR-REQ-037 |
| Kedudukan | Dirancang 06. Menyetujui sekaligus: **[Usulan]** SP-36. Tanggapan dari kartu, termasuk "usul lain" yang K01-nya ditolak guard: 06 §16.4 (**[Usulan]** `architect`, revisi kesepuluh, CA-05) |

**Kriteria penerimaan**
- **AC-1** — *Given* usulan ambang Rp500.000 dari A, *When* B menyetujuinya, *Then* ambang berlaku atas entri yang disimpan sesudahnya dan A menerima N-12.
- **AC-2** — *Given* usulan A, *When* A sendiri mengirim `agree`, *Then* ditolak.
- **AC-3** — *Given* usulan A, *When* B menolak tanpa alasan, *Then* usulan `ditolak`, yang lama tetap berlaku, dan A menerima N-12.
- **AC-4** — *Given* usulan A yang belum dijawab, *When* A menariknya, *Then* `ditarik`.
- **AC-5** — *Given* usulan yang tidak dijawab berapa lama pun, *When* waktu berjalan, *Then* tidak ada yang menjadi berlaku dan usulan tidak kedaluwarsa.
- **AC-6** — *Given* ambang Rp500.000 dan A mengusulkan empat pos rutin Rp150.000, Rp300.000, Rp200.000, dan Rp750.000, *When* B memakai "setujui sekaligus", *Then* yang ditawarkan hanya tiga pos pertama, masing-masing dengan jumlahnya, tanpa pilihan yang sudah terisi; pos Rp750.000 disetujui sendiri; di server tiap pos yang disetujui satu K02.
- **AC-7** — *Given* ambang Rp500.000 dan usulan A menaikkan jatahnya sendiri dari Rp1.000.000 ke Rp1.500.000 tampil sebagai kartu di chat, *When* B memilih "usul lain" pada kartu dan mengisi Rp1.200.000, *Then* server menjalankan K03 atas usulan A lalu K01 baru oleh B dengan `replaces` menunjuk jatah Rp1.000.000 yang berlaku, sebagai satu langkah layar; A menerima N-12 dan N-11; jatah Rp1.000.000 tetap berlaku sampai A menyetujui usulan B (06 §16.4).
- **AC-8** — *Given* AC-7, *When* B mengisi usul lain Rp400.000 (tidak di atas ambang), *Then* sebelum mengirim layar menampilkan batas yang dibaca dari server (ambang Rp500.000 dan bahwa jatah di bawahnya hanya diusulkan pemiliknya) tanpa menghitung guard sendiri (06 S-8); server menjalankan K03 lalu K01; K01 B ditolak guard jatah (INV-A19); K03 tetap berlaku, usulan A berakhir `ditolak` dan tidak hidup kembali, dan tidak ada usulan dari B; jatah Rp1.000.000 tetap berlaku; A menerima N-12 atas K03 seperti biasa dan tidak menerima N-11; B diberi tahu hasilnya di layar (di hasil sinkron bila offline) dan boleh mengirim K01 lain yang sah, misalnya Rp1.200.000, atau tidak sama sekali (06 §16.4 "Usulan kesepakatan", revisi kesepuluh; PA-REQ-5 ditutup).

**Dua arah.** Simetris; "bukan pengusul" adalah satu-satunya syarat.

### FR-REQ-028 — Penggantian kesepakatan dan kapan berlaku

| | |
|---|---|
| Fitur | F-REQ-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `MSL-1` (periode anggaran yang dapat dipercaya) |
| Aktor dan izin | K05 berjalan dalam transaksi K02 atas usulan yang menunjuk kesepakatan berlaku |
| Aturan acuan | 06 §6.3 K05, butir "Hari mulai periode anggaran"; 07 §6.1 tabel penggantian; INV-A37 |
| Business rule | BR-REQ-019, BR-REQ-020, BR-REQ-021 |
| Kedudukan | D-49, D-50 (lewat "lanjut"). Usulan ambang yang naik sampai atau melewati jatah yang berlaku: syarat layar usulannya dan akibatnya atas jatah di FR-FIN-077 ([FIN-1.md](FIN-1.md); 06 §6.3 "Jatah yang tidak di atas ambang", butir "Bila ambang kemudian naik"), tidak diulang di sini |

**Kriteria penerimaan**
- **AC-1** — *Given* ambang Rp500.000 berlaku dan entri Rp600.000 tersimpan bertanda, *When* usulan Rp1.000.000 disetujui pukul 10.00, *Then* entri yang disimpan sesudah 10.00 dinilai terhadap Rp1.000.000 dan entri Rp600.000 tetap bertanda.
- **AC-2** — *Given* hari mulai 1, *When* penggantian ke 25 disetujui pada 10 November 2026, *Then* periode 1 sampai 30 November tidak berubah, periode peralihan 1 sampai 24 Desember (24 hari), lalu 25 Desember sampai 24 Januari.
- **AC-3** — *Given* hari mulai 25, *When* penggantian ke 1 disetujui pada 10 November 2026, *Then* periode 25 Oktober sampai 24 November tidak berubah, peralihan 25 sampai 30 November (6 hari), lalu 1 sampai 31 Desember.
- **AC-4** — *Given* usulan pengganti menunggu, *When* belum disetujui, *Then* nilai lama berlaku; aturan rumah tangga tidak pernah kosong.

**Dua arah.** Simetris; penggantian selalu memerlukan anggota dewasa yang bukan pengusulnya.

### FR-REQ-029 — Mengakhiri kesepakatan rutin

| | |
|---|---|
| Fitur | F-REQ-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` (asas pengawasan, 06 PG-14) |
| Aktor dan izin | Pos rutin: salah satu anggota dewasa. Jatah: hanya pemiliknya. `rule` tidak dapat diakhiri |
| Aturan acuan | 06 §6.3 K06, INV-A38, kasus uji §12 baris K06 K01 (usulan `rule` "tanpa ambang", revisi keempat belas), K02, K05; 07 §3.8 "Ambang hanya lewat kesepakatan", INV-40; ADR-027 (asas pengawasan, diterima lewat D-113) |
| Business rule | BR-REQ-020, BR-REQ-022, BR-REQ-023 |
| Kedudukan | D-52 (lewat "lanjut"); masa penilaian: penegasan 1; jatah tidak di atas ambang dilepas seketika: **[Usulan]** SP-33. Mengakhiri pos rutin sendirian: usulan `architect` (06 P-26, menggantikan P-20; dokumen 03 §7 butir 9), diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-5 (i); harganya W-25. Asas pengawasan atas jatah dan pos rutin: ADR-027 opsi A, diterima lewat "lanjut" di checkpoint S4 (D-113; K-34 ditutup). Ambang tidak diakhiri sendirian: S4-9 (i) A, **diputuskan pemilik produk sendiri** |

**Kriteria penerimaan**
- **AC-1** — *Given* pos rutin "listrik" berkesepakatan Rp400.000, *When* B mengakhirinya, *Then* berlaku seketika dan ke depan, pembayaran berikutnya dinilai dengan ambang per entri, dan A menerima N-12.
- **AC-2** — *Given* jatah A Rp1.000.000 berlaku, *When* B mengirim `end` atasnya, *Then* ditolak; B dapat mengusulkan jumlah lain yang hanya berlaku bila A menyetujuinya.
- **AC-3** — *Given* ambang Rp500.000 dan jatah A Rp1.000.000 dalam periode 1 sampai 30 November, *When* A mengakhirinya pada 15 November, *Then* penyisihan yang disimpan sampai 30 November masih dinilai terhadap Rp1.000.000; sejak 1 Desember dinilai terhadap ambang per entri.
- **AC-4** — *Given* jatah A Rp300.000 (tidak di atas ambang), *When* A mengakhirinya, *Then* berlaku seketika juga atas penilaian, tanpa masa penilaian.
- **AC-5** — *Given* kesepakatan `rule`, *When* siapa pun mengirim `end` (K06) atau perintah sepihak lain atas ambang, *Then* ditolak; `rule` tidak pernah diakhiri lewat K06. Hari mulai periode dan zona waktu tidak pernah dapat diakhiri; ambang diakhiri hanya lewat usulan `rule` "tanpa ambang" yang disetujui semua pihak yang disyaratkan (07 §3.8 "Ambang hanya lewat kesepakatan"; 06 §6.3 K01, K02, K05).
- **AC-6** — **Ambang dibuat, diubah, dan diakhiri hanya lewat kesepakatan** (BR-REQ-020; S4-9 (i) A, diputuskan pemilik produk sendiri; berlaku bagi setiap rumah tangga, pilot dan tahap publik; angka dari 07 INV-40, periode 1 sampai 30 November). *Given* ambang Rp500.000 berlaku, *When* A mengirim perintah mengakhiri atau mengubah ambang sendirian, *Then* ditolak tanpa akibat: tidak ada Peristiwa yang terbaca B, tidak ada notifikasi, dan tidak ada entri, penilaian, atau saldo yang berubah; layar ambang menyatakan bahwa ambang tidak dapat diubah atau diakhiri sendirian (FR-FAM-004/AC-6). *Given* A mengusulkan "tanpa ambang" pada 10 November dan B menyetujuinya pada 12 November, *When* belanja Rp600.000 disimpan 11 November, *Then* bertanda dan tetap bertanda sesudah persetujuan; *When* belanja Rp600.000 disimpan 12 November sesudah persetujuan, *Then* tidak bertanda. *Given* ambang Rp300.000 disepakati 15 November, *When* belanja Rp400.000 disimpan 16 November, *Then* bertanda. *Given* rumah tangga dengan tiga anggota dewasa penuh (tahap publik), *When* dua di antaranya menyetujui usulan yang mengubah atau mengakhiri ambang, *Then* usulan tetap `diusulkan` sampai anggota ketiga juga menyetujuinya (06 §6.3 "Pihak yang disyaratkan", bawaan `all`, S4-8 (i) A). Cabang B dan C yang semula ditulis di sini gugur.

**Dua arah.** Simetris per peran. Pos rutin: siapa pun dari keduanya. Jatah: pemiliknya, siapa pun dia. Ambang: tidak seorang pun mengubah atau mengakhirinya sendirian; karena itu satu orang dapat menahan ambang atas yang lain tanpa batas waktu (keputusan S4-9 (i) A; **Pagar**, risiko yang diterima pemilik produk, dicatat sebagai peringatan bertanggal W-34 di dokumen 17). Ambang mengikat keduanya sama, jadi tidak memberi salah satu kendali yang tidak dimiliki yang lain.

### FR-REQ-030 — Ambang saat mulai memakai dan rumah tangga satu dewasa

| | |
|---|---|
| Fitur | F-REQ-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` (setara sejak awal; tidak ada yang berlaku surut) |
| Aktor dan izin | Anggota pertama boleh mengusulkan; tidak ada yang berlaku sampai anggota dewasa kedua menyetujui |
| Aturan acuan | 06 §3.5 (kedua tabel), §6.3; SP-41 (06 §6.3 "Alternatif yang tidak diambil") |
| Business rule | BR-REQ-019, BR-REQ-026, BR-REQ-035 |
| Kedudukan | D-33, D-42 (lewat "lanjut"). Langkah terpisah dari menerima undangan dan pemberitahuan "tidak dapat diubah sendirian": **[Usulan]** `security-privacy` SP-41 |

**Kriteria penerimaan**
- **AC-1** — *Given* A sendirian mengusulkan ambang Rp500.000, *When* A mencatat belanja Rp2.000.000, *Then* usulan tetap `diusulkan` dan belanja tidak bertanda.
- **AC-2** — *Given* AC-1, *When* B bergabung dan menyetujui ambang pada 12 November pukul 10.00, *Then* hanya entri yang disimpan sesudahnya yang dapat bertanda; belanja Rp2.000.000 tetap tidak bertanda dan B dapat mengajukan keberatan atas entri mana pun sebelumnya.
- **AC-3** — *Given* B menerima undangan, *When* B bergabung, *Then* menerima undangan tidak sama dengan menyetujui ambang: ambang disetujui lewat langkah tersendiri yang menyatakan apa yang tidak dapat dilakukan sendirian **[Usulan]**; kalimatnya ditetapkan di FR-FAM-004/AC-6: "ambang tidak dapat diubah atau diakhiri sendirian", di pilot dan tahap publik, tanpa syarat (S4-9 (i) A).
- **AC-4** — *Given* B keluar dari rumah tangga, *When* A mencatat belanja Rp2.000.000, *Then* tidak bertanda selama hanya ada satu anggota dewasa.

**Dua arah.** Simetris; siapa pun anggota pertamanya.

### FR-REQ-031 — Jalan yang terlihat untuk mengubah kesepakatan

| | |
|---|---|
| Fitur | F-REQ-008 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` |
| Aktor dan izin | Pencatat pembayaran yang melebihi jumlah per kali; anggota dewasa mana pun dari tanda apa pun (dengan guard K01 jatah) |
| Aturan acuan | 06 §6.3 "Pos rutin" baris "Sesudah pembayaran di atas jumlah yang disepakati", SP-35; dokumen 17 Keluhan 1 |
| Business rule | BR-REQ-019, BR-REQ-023 |
| Kedudukan | Usulan satu ketukan: **[Usulan]** SP-37, syarat dokumen 05. Tempat status: **[Usulan]** SP-35 |

**Kriteria penerimaan**
- **AC-1** — *Given* pos rutin "listrik" berkesepakatan Rp400.000, *When* A mencatat pembayaran Rp430.000 yang tertaut, *Then* A ditawari satu ketukan untuk mengusulkan Rp430.000 sebagai jumlah per kali baru; bila diabaikan, tidak ada akibat.
- **AC-2** — *Given* entri bertanda dasar apa pun, *When* pencatatnya atau pasangannya membuka tanda itu, *Then* tanda menyebut dasarnya dan menyediakan jalan ke kesepakatan yang menjadi dasarnya (ambang, jatah, atau jumlah per kali) untuk mengusulkan perubahan (Keluhan 1).
- **AC-3** — *Given* pos rutin yang kesepakatannya belum dijawab, *When* keduanya melihat pembayarannya, layar utama, atau daftar apa pun, *Then* status kesepakatan hanya tampil di layar pos itu sendiri, tidak sebagai hitungan.

**Dua arah.** Simetris; tawaran mengikuti pencatat pembayaran, siapa pun dia.

---

## F-REQ-009 — Tanda atas catatan

### FR-REQ-032 — Penilaian tanda oleh server

| | |
|---|---|
| Fitur | F-REQ-009 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `MSL-4` (belanja tak terencana terlihat tanpa harus bertanya) |
| Aktor dan izin | Server saja (S-8). Tanda sama bagi kedua anggota dewasa |
| Aturan acuan | 06 §6.2, INV-A17, INV-A37, S-8; 07 §3.8 "Tanda atas entri"; ADR-009 |
| Business rule | BR-REQ-024, BR-REQ-025, BR-REQ-026, BR-REQ-011 |
| Kedudukan | Dirancang 06 dan 07. Jatah sebagai batas: keputusan 20 (sendiri). Pos rutin: keputusan 22 dan penegasan 3 |

Rumus dan ketiga dasar (`threshold`, `personal_share`, `routine_amount`) kanonis di 07 §3.8. Dokumen ini hanya menuntut perilakunya: paling banyak satu tanda per entri, dinilai satu kali saat entri diterima server, tidak pernah menghalangi pencatatan.

**Kriteria penerimaan**
- **AC-1** — *Given* ambang Rp500.000, *When* A mencatat belanja dana bersama Rp600.000 tanpa permintaan, *Then* entri tersimpan dan bertanda dasar ambang dengan nilai ambang saat itu.
- **AC-2** — *Given* ambang Rp500.000, *When* belanja tanpa permintaan sebesar Rp500.000, *Then* tidak bertanda (perbandingannya "di atas").
- **AC-3** — *Given* jatah A Rp1.000.000 per periode, *When* tiga penyisihan Rp400.000, Rp400.000, dan Rp900.000 tiba dalam urutan itu, *Then* hanya yang ketiga bertanda dasar jatah; *When* yang Rp900.000 tiba lebih dulu, *Then* dua tanda; jumlah tanda tidak disajikan sebagai ukuran kelebihan (W-13).
- **AC-4** — *Given* ambang Rp500.000 dan jatah A Rp1.000.000, *When* lima penyisihan Rp400.000 dalam satu periode, *Then* yang ketiga, keempat, dan kelima bertanda (jumlah periode Rp1.200.000, Rp1.600.000, Rp2.000.000); tanpa jatah, tidak satu pun bertanda (W-17).
- **AC-5** — *Given* pos rutin "listrik" berkesepakatan Rp400.000, *When* siapa pun membayar Rp400.000, *Then* tidak bertanda; *When* Rp401.000, *Then* bertanda dasar pos rutin.
- **AC-6** — *Given* pembayaran cicilan Rp1.200.000 dengan denda Rp50.000 di entri yang sama, *When* dicatat, *Then* pelunasannya tidak pernah bertanda dan dendanya dibandingkan dengan ambang (tidak bertanda).
- **AC-7** — *Given* entri bertanda Rp600.000, *When* ambang kemudian dinaikkan ke Rp1.000.000, *Then* tanda entri itu tidak berubah.
- **AC-8** — *Given* rumah tangga dengan satu anggota dewasa, *When* entri berapa pun dicatat, *Then* tidak ada tanda.

**Dua arah.** Simetris; rumusnya tidak membaca siapa pencatatnya. Akibat yang tidak sama yang dinyatakan terbuka (W-14, W-17: orang berjatah lebih sering ditandai atas jumlah kecil) adalah hasil aturan yang sama, bukan aturan yang berbeda.

### FR-REQ-033 — Yang tampil pada tanda dan tanggapan pasangan

| | |
|---|---|
| Fitur | F-REQ-009 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `MSL-4`; arah 1 pemilik produk ("bukan menunjukan siapa yang salah") |
| Aktor dan izin | Keterangan "mendesak" dan "sudah ada sebelum kami mulai mencatat" hanya oleh pencatat. Tanggapan pasangan tidak wajib |
| Aturan acuan | 06 §6.2 tabel "Yang tampil", "Tanggapan pasangan", "Keperluan mendesak", putusan 2; §8 |
| Business rule | BR-REQ-024, BR-REQ-032 |
| Kedudukan | Tanpa persetujuan susulan wajib: diterima atas saran lead (S2-2 A) |

**Kriteria penerimaan**
- **AC-1** — *Given* entri bertanda, *When* A dan B membukanya, *Then* keduanya melihat tanda yang sama dengan dasarnya (nilai ambang, catatan penilaian jatah, atau jumlah per kali).
- **AC-2** — *Given* entri A bertanda, *When* A menambahkan keterangan "mendesak", *Then* keterangan tampil bersama tanda, tanda tetap ada, dan tidak ada yang diminta dari B.
- **AC-3** — *Given* entri A bertanda, *When* B tidak melakukan apa pun, *Then* tidak ada yang berubah dan tidak ada pengingat kepada siapa pun.
- **AC-4** — *Given* entri A bertanda, *When* B memilih menanggapi, *Then* yang tersedia membuka diskusi pada entri atau mengajukan keberatan dengan alasan dari daftar tertutup; keduanya tidak mengubah angka.

**Dua arah.** Simetris; pencatat dan pasangan adalah peran terhadap entri.

### FR-REQ-034 — Pemberitahuan atas tanda

| | |
|---|---|
| Fitur | F-REQ-009 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `MSL-4`; `PAGAR` (tidak menjadi pantauan per belanja) |
| Aktor dan izin | Penerima: anggota dewasa selain pencatat |
| Aturan acuan | 06 §6.2 "Notifikasi", N-10, kasus uji §12 "Prioritas N-10 menurut dasar" |
| Business rule | BR-REQ-027 |
| Kedudukan | Satu per hari: D-35. Hanya dasar ambang: diterima atas saran lead, 12 November 2026. Satu baris tanpa hitungan: **[Usulan]** SP-43 butir b |

**Kriteria penerimaan**
- **AC-1** — *Given* satu hari dan satu penerima B, *When* entri A bertanda berturut-turut dasar jatah, dasar pos rutin, lalu dua kali dasar ambang, *Then* dua yang pertama hanya masuk ringkasan; tanda ambang pertama menjadi N-10 `kabar` (tunduk pada jam tenang); yang kedua masuk ringkasan; ringkasan hari itu memuat satu baris tanpa angka dan tanpa nama untuk semua butir tanda. Urutan dibalik memberi hasil per dasar yang sama.
- **AC-2** — *Given* jatah A sudah terlewati, *When* A mengambil lima kali dalam sehari, *Then* lima entri bertanda dan B tidak menerima satu pun pemberitahuan langsung (Keluhan 1).
- **AC-3** — *Given* tanda pada satu entri hilang lalu muncul kembali karena tautan dicabut, *When* tanda muncul kembali, *Then* N-10 tidak dikirim lagi untuk entri itu.

**Dua arah.** Ditulis dengan A sebagai pencatat; menukar A dan B memberi hasil yang tertukar karena aturan dihitung per penerima, bukan per orang tertentu.

### FR-REQ-035 — Tanda hilang dan muncul kembali

| | |
|---|---|
| Fitur | F-REQ-009 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `MSL-4`, `MSL-5` |
| Aktor dan izin | Menurut FR-REQ-015 (siapa boleh menautkan dan mencabut) |
| Aturan acuan | 06 §6.2 "Tanda hilang", "Tanda muncul kembali", INV-A34; 07 §3.8 aturan pelengkap |
| Business rule | BR-REQ-028, BR-REQ-014 |
| Kedudukan | D-39 (lewat "lanjut") |

**Kriteria penerimaan**
- **AC-1** — *Given* entri A Rp600.000 bertanda dasar ambang dan otorisasi terbuka milik A Rp1.000.000 dengan realisasi Rp0, *When* A menautkannya, *Then* tanda hilang, pelaku tautan tampil, dan riwayat tandanya tetap di Peristiwa.
- **AC-2** — *Given* entri yang sama dan otorisasi milik A Rp50.000 dengan realisasi Rp0, *When* A menautkannya, *Then* N = Rp550.000, tanda tetap (dasar ambang), dan permintaan menampilkan "melebihi Rp550.000" dengan N-08.
- **AC-3** — *Given* entri A Rp450.000 tanpa tanda dan pos rutin berkesepakatan Rp400.000, *When* A menautkannya ke pos itu, *Then* entri menjadi bertanda dasar pos rutin; *When* jumlah per kalinya Rp500.000, *Then* tetap tanpa tanda.
- **AC-4** — *Given* entri bertanda, *When* pencatatnya membaliknya, *Then* tanda hilang.
- **AC-5** — *Given* tanda hilang karena A menautkan entri A, *When* B mencoba mencabut tautan itu, *Then* ditolak; tanda tidak dapat dipasang kembali oleh pasangan.
- **AC-6** — *Given* entri bertanda karena kelebihan (CTH-24 varian Rp800.000), *When* perubahan syarat menaikkan batas ke Rp800.000 diterima, *Then* tampilan "melebihi Rp…" pada permintaan hilang, tetapi tanda pada entri yang sudah tersimpan tidak berubah (07 §3.8 "Dinilai satu kali"). Yang hilang adalah tampilan kelebihan pada permintaan, bukan tanda entri (dikukuhkan 06 §5.4, revisi kesembilan).

**Dua arah.** Simetris; ditulis dengan A sebagai pencatat, dan wewenangnya peran terhadap entri.

### FR-REQ-036 — Tanpa hitungan dan papan skor

| | |
|---|---|
| Fitur | F-REQ-009 |
| Tier / MoSCoW | 1 / Must; Pantau |
| Masalah (brief §1.1) | `PAGAR`; arah 1 pemilik produk |
| Aktor dan izin | Berlaku atas setiap layar, laporan, ringkasan, dan ekspor |
| Aturan acuan | 06 §6.2 "Hitungan", PG-07, PG-09, AT-1 sampai AT-7 (AT-5 dirumuskan ulang), §16.5, §16.6, INV-A45; dokumen 17 Keluhan 1 dan 7 |
| Business rule | BR-REQ-032 |
| Kedudukan | D-23, D-43 (lewat "lanjut"). Tanda dan kehadiran di chat: lihat BR-REQ-032 |

**Kriteria penerimaan**
- **AC-1** — *Given* berapa pun tanda, keberatan, pengingat, permintaan, dan penolakan yang ada, *When* siapa pun membuka layar, laporan, ringkasan notifikasi, atau ekspor mana pun, *Then* tidak ada jumlah, peringkat, rekap, atau tren per orang atas semua itu, dan tidak ada lama menanggapi per orang (Keluhan 1, Keluhan 7).
- **AC-2** — *Given* daftar entri, *When* siapa pun mencari cara menyaring, mengurutkan, menyimpan tampilan, atau mengekspor kolom menurut tanda, keberatan, koreksi, atau catatan terlambat, *Then* tidak ada.
- **AC-3** — *Given* aplikasi apa pun, *When* B membuka apa pun milik A di luar layar chat, *Then* tidak ada umpan kegiatan A, tanda "sudah dibaca", "terakhir aktif", "sedang online", atau lokasi (Keluhan 7; 06 AT-5 (b)). Di layar chat, tanda dan kehadiran mengikuti 06 §16.5 dan §16.6 (lihat COM.md).
- **AC-4** — *Given* tanda dibaca dan kehadiran di chat diubah acak di antara dua jalankan rangkaian perintah yang sama, *When* layar permintaan, daftar menunggu, ringkasan, notifikasi, laporan, dan ekspor dibandingkan, *Then* semuanya identik; tidak ada pengingat, lencana, warna, atau urutan "dibaca tapi belum dijawab", dan tidak ada hitungan lama menjawab (06 INV-A45, PG-07).

**Dua arah.** Simetris; larangannya berlaku atas data siapa pun.

---

## F-REQ-010 — Konfirmasi susulan

Seluruh fitur ini berkedudukan: jalurnya usulan lead (Q-112); tiga batasnya (tanpa pengingat otomatis, satu pengingat ulang, tidak diajukan lagi sesudah ditolak) diterima atas saran lead, 12 November 2026; rincian lain usulan `security-privacy` SP-38 dan `architect` (Q-124, Q-125, Q-130). N-01 dan satu pengingat ulangnya selalu `ringkasan`, tidak pernah pemberitahuan langsung: usulan `security-privacy` SP-52, diterapkan 06 revisi kesepuluh (N-01, N-20, INV-A40 (f)). Jalur, rincian, dan SP-52 **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026), S3-6 (i) sampai (v).

### FR-REQ-037 — Mengajukan konfirmasi susulan

| | |
|---|---|
| Fitur | F-REQ-010 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`, `MSL-5` (kesepakatan lisan tidak perlu diketik ulang sebagai permintaan) |
| Aktor dan izin | Hanya pencatat entri, atas kehendaknya sendiri. Pasangan tidak punya perintah "minta konfirmasi"; aplikasi tidak pernah mengajak pencatat memintanya |
| Aturan acuan | 06 §6.2 "Konfirmasi susulan" (Bentuk, Siapa yang memulai, Guard B01 tambahan), §4.3 paragraf konfirmasi susulan, INV-A40 butir a |
| Business rule | BR-REQ-029 |
| Kedudukan | Lihat pembuka fitur |

Syaratnya ditetapkan server dari entri: jenis pemakaian, kategori, dan batas = *jumlah pemakaian* entri. Tanpa tempat uang sumber. Guard tambahan: pelaku pencatat; entri saat ini bertanda (dasar apa pun), belum dibalik, belum tertaut ke otorisasi; tidak ada permintaan lain yang belum terminal atas entri itu; tidak pernah ada yang berakhir `rejected` atas entri itu.

**Kriteria penerimaan**
- **AC-1** — *Given* entri A Rp600.000 bertanda dasar ambang, *When* A mengajukan konfirmasi susulan, *Then* lahir permintaan berbatas Rp600.000 yang ditetapkan server, tanpa tempat uang sumber, dan N-01 kepada B berprioritas `ringkasan`: ia masuk ringkasan B berikutnya dan tidak pernah dilepas sebagai pemberitahuan langsung, juga bila plafon harian B masih longgar (06 N-01, SP-52).
- **AC-2** — *Given* entri A bertanda, *When* B mengajukan konfirmasi susulan atasnya, *Then* ditolak.
- **AC-3** — *Given* entri tanpa tanda, entri yang sudah tertaut, entri yang sudah dibalik, atau entri yang sudah punya konfirmasi susulan yang belum terminal, *When* pencatatnya mengajukan, *Then* ditolak.
- **AC-4** — *Given* pengajuan konfirmasi susulan, *When* pencatat mengirim jumlah lain atau tempat uang sumber, *Then* ditolak.
- **AC-5** — *Given* penyisihan A bertanda dasar jatah, atau pembayaran bertanda dasar pos rutin, *When* A mengajukan konfirmasi susulan, *Then* diterima.
- **AC-6** — *Given* entri bertanda yang tidak pernah dimintakan konfirmasi susulan, *When* siapa pun melihat entri itu, *Then* ia tampil persis seperti entri bertanda lain dan tidak ada ajakan kepada pencatat.
- **AC-7** — *Given* N-01 konfirmasi susulan A masuk antrean ringkasan B, *When* B sudah menyetujui atau menolaknya, atau A membatalkannya, sebelum ringkasan dilepas, *Then* butir itu tidak ada di ringkasan B, karena ia meminta tindakan atas objek yang tidak lagi menunggu B (06 §11.1 "Jadwal ringkasan", butir yang dibuang (2)).

**Dua arah.** Simetris; pencatat adalah peran terhadap entri.

### FR-REQ-038 — Menanggapi konfirmasi susulan

| | |
|---|---|
| Fitur | F-REQ-010 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`, `MSL-5` |
| Aktor dan izin | Pemberi persetujuan menyetujui (satu ketukan), menolak, atau tidak melakukan apa pun. Yang disetujui adalah pemakaian dana bersama sebesar itu, bukan "entri itu" |
| Aturan acuan | 06 §6.2 "Konfirmasi susulan" (Saat disetujui, Pelaku pada `entry_link`, Bila entrinya sudah dibalik, Bila pasangan diam, Bila ditolak), INV-A40; kasus uji §12; 07 §3.8, INV-26 |
| Business rule | BR-REQ-029, BR-REQ-005 |
| Kedudukan | Lihat pembuka fitur |

**Kriteria penerimaan**
- **AC-1** — *Given* konfirmasi susulan A Rp600.000, *When* B menyetujuinya, *Then* dalam satu transaksi lahir otorisasi Rp600.000, tautan dari entri itu ber-`actor` A, dan otorisasi `settled` dengan realisasi Rp600.000, terikat Rp0, tanpa reservasi; tanda entri hilang (N = 0); A menerima N-04 dan tidak ada N-09 tersendiri.
- **AC-2** — *Given* konfirmasi susulan A, *When* B menolaknya, *Then* tidak ada medan alasan, entri tidak berubah dan tidak mendapat tanda tambahan, tanda semula tetap, dan diskusi tetap tersedia.
- **AC-3** — *Given* konfirmasi susulan A, *When* B tidak melakukan apa pun, *Then* tidak ada akibat; permintaan menunggu tanpa batas dan setiap pembacaan entri identik dengan keadaan tanpa permintaan itu.
- **AC-4** — *Given* A membalik entri itu sebelum persetujuan B tiba di server, *When* persetujuan tiba, *Then* ditolak `state_changed`, permintaan `cancelled`, dan tidak ada otorisasi.
- **AC-5** — *Given* A menautkan entri itu ke otorisasi lain lebih dulu, *When* persetujuan B tiba, *Then* sama dengan AC-4.
- **AC-6** — *Given* otorisasi konfirmasi susulan sudah `settled`, *When* entri terlambat bertanggal hari persetujuan ditautkan kepadanya, *Then* ditolak dan entrinya tersimpan tanpa tautan.
- **AC-7** — *Given* tautan yang lahir pada AC-1, *When* B mencoba mencabutnya, *Then* ditolak; hanya A yang dapat.

**Dua arah.** Simetris per peran (pencatat dan pasangannya).

### FR-REQ-039 — Pengingat dan pengajuan ulang konfirmasi susulan

| | |
|---|---|
| Fitur | F-REQ-010 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR` (tidak bergeser menjadi "persetujuan susulan yang diharapkan") |
| Aktor dan izin | Pencatat sebagai pihak yang menunggu |
| Aturan acuan | 06 §6.2 "Notifikasi dan pengingat", Guard B01 butir 4, INV-A32, INV-A40 butir f |
| Business rule | BR-REQ-016, BR-REQ-017, BR-REQ-029 |
| Kedudukan | Ketiga batas: diterima atas saran lead, 12 November 2026. Pengingat ulang `ringkasan`: **[Usulan]** SP-52 (lihat pembuka fitur) |

**Kriteria penerimaan**
- **AC-1** — *Given* konfirmasi susulan A menunggu, *When* waktu berjalan berapa pun, *Then* tidak pernah ada pengingat otomatis.
- **AC-2** — *Given* A sudah menekan "ingatkan lagi" satu kali pada rantai itu, *When* A menekan lagi, *Then* ditolak.
- **AC-3** — *Given* konfirmasi susulan A ditolak B, *When* A mengajukan konfirmasi susulan lagi atas entri yang sama, *Then* ditolak.
- **AC-4** — *Given* konfirmasi susulan A dibatalkan A, *When* A mengajukannya lagi lewat "ajukan lagi", *Then* diterima dalam rantai yang sama dan hitungan pengingat ulangnya diwarisi.
- **AC-5** — *Given* konfirmasi susulan A menunggu dan A belum memakai pengingat ulangnya, *When* A menekan "ingatkan lagi", *Then* perintah diterima dan N-20 kepada B berprioritas `ringkasan`: ia masuk ringkasan B berikutnya, tidak pernah dilepas langsung, dan dibuang dari ringkasan bila konfirmasi itu tidak lagi menunggu B saat ringkasan dilepas (06 N-20, §11.1; SP-52).

**Dua arah.** Simetris per peran.

### FR-REQ-040 — Tampilan konfirmasi susulan

| | |
|---|---|
| Fitur | F-REQ-010 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `PAGAR`; `MSL-1` |
| Aktor dan izin | Para pihak |
| Aturan acuan | 06 §6.2 "Yang ditampilkan", lima hal penjaga, INV-A40 butir b dan g |
| Business rule | BR-REQ-029, BR-REQ-032 |
| Kedudukan | **[Usulan]** SP-38 |

**Kriteria penerimaan**
- **AC-1** — *Given* entri bertanggal 8 Desember 2026 dan konfirmasi susulan diajukan 10 Desember 2026, *When* permintaan tampil di layar atau ekspor mana pun, *Then* ia tampil diajukan sesudah belanjanya terjadi, dengan kedua tanggal, dan tidak pernah sebagai permintaan yang diajukan lebih dulu.
- **AC-2** — *Given* entri bertanda sebelum disetujui, sesudah ditolak, atau tanpa permintaan, *When* siapa pun membukanya, *Then* tandanya tidak pernah berbunyi "belum disetujui" atau sejenisnya, dan entri tidak punya state "menunggu", "ditolak", atau "dikonfirmasi".
- **AC-3** — *Given* daftar entri, *When* siapa pun mencari penyaring, urutan, atau kolom ekspor menurut ada-tidaknya atau hasil konfirmasi susulan, *Then* tidak ada; tidak ada hitungan per orang atas konfirmasi susulan.
- **AC-4** — *Given* saldo, anggaran, dan laporan, *When* dibandingkan sebelum dan sesudah konfirmasi susulan disetujui, *Then* sama; yang berubah hanya realisasi otorisasi baru itu dan hilangnya tanda.

**Dua arah.** Simetris.

---

## F-REQ-011 — Diskusi dan lampiran pada objek

### FR-REQ-041 — Diskusi pada objek

| | |
|---|---|
| Fitur | F-REQ-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-4`; brief §5B "contextual conversations" |
| Aktor dan izin | Siapa pun yang boleh melihat objek induknya: permintaan (para pihak), kesepakatan (semua anggota dewasa), entri (menurut visibilitas entri), keberatan (yang melihat sasarannya). Penulis menghapus pesannya sendiri; tidak ada yang menyunting |
| Aturan acuan | 06 §2.1, §4.4, §2.5 (rincian), §8.2, §16.2, §16.3, §16.7, §16.8, N-19, N-23, INV-A44, INV-A46, INV-A47; ADR-001, ADR-011; D-38 |
| Business rule | BR-REQ-030, BR-REQ-031, BR-REQ-037 |
| Kedudukan | Dirancang 06; D-38 (lewat "lanjut"). Diskusi tampil di chat sebagai balasan yang menempel pada kartu: diterima atas saran lead ("Chat umum + kartu", "Reply nempel ke kartu"), 12 November 2026; rinciannya **[Usulan]** `architect` (06 §16.3). Pesan dari layar objek tanpa kartu: **[Usulan]** `architect` (06 §16.3, revisi kesepuluh; PA-REQ-6 ditutup). Hapus dengan penanda: usulan `architect` (06 P-22), diterima lewat "lanjut" di checkpoint S3 sebagai D-71 |

**Diskusi kini tampil di chat.** Atas objek yang punya kartu di chat rumah tangga, diskusi objek itu adalah pesan yang membalas kartunya (atau membalas pesan lain yang sudah menempel padanya), ditambah pesan yang ditautkan penulisnya ke kartu sesudah dikirim (`link_message`). Pesan yang sama tampil di chat pada urutannya dan di layar objek sebagai diskusinya. Pesan bebas di chat tanpa balasan atau tautan tidak masuk riwayat objek mana pun; server tidak menebak hubungannya. Diskusi yang tidak tampil di chat hanya ada atas objek yang tidak terlihat oleh semua peserta chat; di pilot dua orang dewasa praktis tidak ada (06 N-19). Fitur chat itu sendiri: lihat COM.md.

**Kriteria penerimaan**
- **AC-1** — *Given* permintaan A menunggu dan kartunya ada di chat, *When* B membalas kartu itu dengan pesan, *Then* state tidak berubah; pesan tampil di chat dan di diskusi permintaan dengan isi yang sama; A menerima N-23 `kabar` (dilipat per pengirim), bukan N-19; pesan itu tidak terhitung dalam pembatas `perlu_tanggapan` (06 §16.3, §16.8, INV-A46).
- **AC-2** — *Given* permintaan `settled`, `rejected`, atau `cancelled`, *When* A membukanya, *Then* diskusi tetap dapat dibaca dan ditulisi, dari layar permintaan maupun dengan membalas kartunya di chat.
- **AC-3** — *Given* pesan A, *When* A menghapusnya, *Then* keduanya melihat penanda "pesan dihapus" beserta waktunya, di chat dan di diskusi objek; *When* B mencoba menghapus atau siapa pun mencoba menyunting pesan A, *Then* ditolak (06 §16.7, INV-A47 (a)).
- **AC-4** — *Given* B mengajukan keberatan atas entri A, *When* B ingin menjelaskan, *Then* keberatan hanya memuat alasan dari daftar tertutup dan penjelasan bebas ditulis di diskusi yang menempel padanya, yang dapat dihapus B sendiri (Keluhan 10).
- **AC-5** — *Given* entri berpemilik pribadi A, *When* B mencari diskusinya, *Then* tidak ditemukan; dan tidak ada kartu, balasan, atau pesan ber-`subject_ref` entri itu di chat (06 INV-A47 (c)).
- **AC-6** — *Given* A mengirim pesan bebas "jangan lupa beli gas" tanpa membalas kartu mana pun, *When* B membuka layar permintaan apa pun, *Then* pesan itu tidak ada di diskusi objek mana pun; *When* A kemudian menautkannya ke kartu permintaan belanja dapur (`link_message`), *Then* pesan itu masuk diskusi permintaan itu, tautannya tercatat sebagai Peristiwa, dan tautan itu tidak dapat dilepas selain dengan menghapus pesannya (06 §16.3).
- **AC-7** — *Given* A membalas kartu permintaan saat offline dan sementara itu objeknya menyempit sehingga tidak lagi terlihat oleh B, *When* perangkat A tersinkron, *Then* pesan tersimpan sebagai pesan bebas tanpa `subject_ref` dan A diberi tahu hasil sinkronnya (06 §16.3 "Batas").
- **AC-8** — *Given* entri buku rumah tangga A yang tidak punya kartu, *When* B menulis pesan di diskusi entri itu dari layar entri, *Then* pesan tampil di chat sebagai pesan yang menempel pada entri itu, dengan `subject_ref` entri dan rujukan ringkas yang dibaca dari state saat ini (bukan salinan; "tidak lagi tersedia" bila entri menyempit); tidak ada kartu yang lahir; membalas pesan itu juga menempel pada entri yang sama; A menerima N-23 saja, bukan N-19 (06 §16.3 "Pesan dari layar objek tanpa kartu", revisi kesepuluh, **[Usulan]** `architect`; menutup PA-REQ-6 dan PA-COM-3). *Given* objek yang tidak terlihat oleh semua peserta, *Then* pesan tetap di diskusi objek dan notifikasinya N-19.

**Dua arah.** Simetris; hak menulis dan menghapus mengikuti penulis dan visibilitas objek.

### FR-REQ-042 — Lampiran pada objek

| | |
|---|---|
| Fitur | F-REQ-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-5` (foto struk cukup); `MSL-1` |
| Aktor dan izin | Sama dengan diskusi objeknya. Lampiran termasuk *rincian* yang dapat dihapus (06 §2.5) |
| Aturan acuan | 06 §2.1 (primitif lampiran), §2.5; §11.1 isi layar kunci |
| Business rule | BR-REQ-030, BR-REQ-031 |
| Kedudukan | Dirancang 06. Siapa yang boleh menghapus lampiran: **[Usulan]** `product-analyst`, sama dengan pesan (pengunggahnya). Pembuangan semua metadata foto: **[Usulan]** `security-privacy` SP-57, sama dengan FR-COM-003/AC-2. Jenis dan ukuran berkas: dokumen 09 dan 10 (S4) |

**Kriteria penerimaan**
- **AC-1** — *Given* entri belanja A, *When* A melampirkan foto struk, *Then* lampiran terlihat oleh siapa pun yang melihat entri itu dan tidak mengubah angka entri.
- **AC-2** — *Given* lampiran A, *When* A menghapusnya, *Then* lampiran hilang dan penanda penghapusan tampil; inti entri tidak berubah.
- **AC-3** — *Given* lampiran pada objek apa pun, *When* notifikasi tentang objek itu tampil, *Then* tidak memuat gambar atau pratinjaunya.
- **AC-4** — *Given* foto struk yang membawa metadata (EXIF termasuk lokasi, waktu, dan perangkat; XMP; thumbnail tertanam), *When* A melampirkannya, *Then* perangkat A mengode ulang gambar dan membuang semua metadata itu sebelum mengunggah, server memeriksanya lagi dan membuang yang masih tersisa sebelum menyimpan, dan berkas yang tersimpan maupun yang diterima B tidak memuat metadata apa pun (SP-57; 06 AT-6).

**Dua arah.** Simetris.

---

## F-REQ-012 — Yang menunggu saya dan yang menunggu pasangan

### FR-REQ-043 — Daftar yang menunggu saya dan yang menunggu pasangan

| | |
|---|---|
| Fitur | F-REQ-012 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | `MSL-1`, `MSL-4`; `KS-1`; PRN-01 |
| Aktor dan izin | Setiap anggota dewasa melihat dua bagian: objek yang menunggu tanggapannya, dan objek miliknya yang menunggu tanggapan pasangan. Isinya hanya objek yang memang boleh ia lihat |
| Aturan acuan | 06 AT-4, AT-5, INV-A33, INV-A35, INV-A45; §4.2 (yang ditunggu per state), §6.3, §7.2, §16.2 "Menunggu sejak"; dokumen 17 Keluhan 6 dan 7 |
| Business rule | BR-REQ-032, BR-REQ-035, BR-REQ-011 |
| Kedudukan | Fitur dari dokumen 04. "Sejak kapan" sebagai umur objek: dikukuhkan 06 §16.2 baris "Menunggu sejak" (revisi kesembilan). Tanpa angka pada bagian "menunggu pasangan": **[Usulan]** `product-analyst` |

Yang masuk: permintaan dan konfirmasi susulan yang menunggu tanggapan, usul balik, perubahan syarat yang menunggu, usulan kesepakatan, serah terima atau catatan yang belum dikonfirmasi pemegangnya, dan (Tier 2) permintaan tindakan yang belum diterima. "Sejak kapan" adalah waktu server saat objek mulai menunggu pihak itu (B01, B03, B04, B11, K01, C01), yaitu umur objek, bukan lama tanggapan seseorang, bukan waktu kartunya dibaca, dan bukan waktu notifikasinya dilepas; nilai yang sama tampil di daftar dan di kartu chat (06 §16.2). Daftar ini tidak membaca tanda dibaca chat: kartu yang sudah dibaca tetap tampil sama dengan yang belum. Keberatan tidak masuk: ia tidak menunggu tanggapan pasangan dan tampil di daftar "belum sepakat" (F-FIN-018).

**Kriteria penerimaan**
- **AC-1** — *Given* A mengajukan permintaan 1 Desember pukul 08.00 dan B belum menanggapi, *When* A membuka daftar, *Then* permintaan itu ada di "menunggu pasangan" dengan "sejak 1 Desember 08.00"; *When* B membuka daftar, *Then* permintaan yang sama ada di "menunggu saya" dengan waktu yang sama (Keluhan 6).
- **AC-2** — *Given* AC-1, *When* B mengajukan usul balik pada 2 Desember pukul 09.00, *Then* permintaan pindah ke "menunggu saya" A dan "menunggu pasangan" B, sejak 2 Desember 09.00.
- **AC-3** — *Given* A mencatat serah terima ke tempat uang yang dipegang B, *When* B belum mengonfirmasi, *Then* catatan itu ada di "menunggu saya" B dan "menunggu pasangan" A, tanpa batas waktu dan tanpa konfirmasi otomatis.
- **AC-4** — *Given* B sedang menunda, dalam jam tenang, atau tidak keduanya, *When* A membuka daftar, *Then* isinya identik pada ketiga keadaan (INV-A33).
- **AC-5** — *Given* daftar apa pun, *When* dibuka, *Then* tidak ada angka "berapa yang menunggu pasangan", rata-rata lama menanggapi, urutan menurut orang, tanda "dibaca", "terakhir aktif", atau "sedang online"; *When* B sudah membaca kartu permintaan A di chat tetapi belum menanggapi, *Then* daftar A tetap identik dengan sebelum B membacanya (Keluhan 7; 06 AT-5 (b), INV-A45).
- **AC-6** — *Given* permintaan A menunggu, *When* A ingin bertindak, *Then* A tetap dapat mencatat apa pun, membatalkan, atau mengajukan lagi; tidak ada yang menjadi disetujui atau ditolak karena diam (Keluhan 6).
- **AC-7** — *Given* B keluar dari rumah tangga, *When* A membuka daftar, *Then* objek yang hanya dapat ditutup B tidak ada di daftar dan tampil "sudah keluar" pada objeknya.

**Dua arah.** Simetris: "menunggu saya" milik yang satu adalah "menunggu pasangan" milik yang lain, isinya sama.

---

## F-REQ-013 — Permintaan tindakan (Tier 2, ringkas)

### FR-REQ-044 — Mengajukan dan menanggapi permintaan tindakan

| | |
|---|---|
| Fitur | F-REQ-013 |
| Tier / MoSCoW | 2 / Should |
| Masalah (brief §1.1) | `MSL-6`, `MSL-7`; `KB-13` |
| Aktor dan izin | T01 pemohon; T02, T03 hanya penerima tugas; T04 pemohon. Pemohon ≠ penerima tugas |
| Aturan acuan | 06 §13.1 T01 sampai T04; §16.2 (kartu otomatis T01), §16.4 baris "Permintaan tindakan"; PRN-10 |
| Business rule | BR-REQ-036, BR-REQ-033, BR-REQ-037 |
| Kedudukan | D-27 (lewat "lanjut"). Kartu dan tombolnya: **[Usulan]** `architect` (06 §16.4) |

Judul, catatan, tenggat opsional. Tidak ada tugas sebelum penerima menerima; menolak tidak wajib beralasan; tidak ada usul balik (penjelasan di diskusi, yaitu balasan pada kartunya di chat). Permintaan tindakan dapat dikirim dari chat; ia lahir sebagai kartu di chat dalam transaksi yang sama dengan T01.

**Kriteria penerimaan**
- **AC-1** — *Given* A meminta B "ganti lampu dapur" dengan tenggat 15 Desember, *When* diajukan, *Then* B menerima pemberitahuan `perlu_tanggapan`, kartunya tampil di chat, dan belum ada tugas yang lahir.
- **AC-2** — *Given* AC-1, *When* B menolak tanpa alasan, *Then* `ditolak`, A menerima `kabar`, dan tidak ada pengingat sesudahnya.
- **AC-3** — *Given* AC-1, *When* A mencoba menerima atas nama B, *Then* ditolak.
- **AC-4** — *Given* kartu permintaan tindakan A di chat, *When* B mengetuk "terima" atau "belum sepakat", *Then* server menjalankan T02 atau T03 yang sama dengan dari layar objek; A tidak melihat tombol itu pada kartunya; bila A sudah membatalkan (T04), perintah B ditolak `state_changed` dan kartu menampilkan state terbaru (06 §16.4, INV-A44).

### FR-REQ-045 — Tugas yang diterima dan pengingatnya

| | |
|---|---|
| Fitur | F-REQ-013 |
| Tier / MoSCoW | 2 / Should |
| Masalah (brief §1.1) | `MSL-6` (permintaan tetap ada sampai selesai) |
| Aktor dan izin | T05 dan T06 hanya penerima tugas; T04 pemohon |
| Aturan acuan | 06 §13.1 T04 sampai T06, tabel "Rincian pengingat"; §11.1 "Waktu pengingat yang berpatokan tanggal"; N-30; PG-06; §16.4 baris "Permintaan tindakan" |
| Business rule | BR-REQ-036, BR-REQ-016, BR-REQ-037 |
| Kedudukan | T06 **[Usulan]** `architect`. Rincian pengingat: 06 §13.1 tabel "Rincian pengingat" dan N-30 (**[Usulan]** `architect`, revisi kesembilan) |

Empat jenis pengingat, menurut 06 §13.1: (1) T01 dan pengingat ulang pemohon selama `diajukan`, `perlu_tanggapan` dengan semua batas biasa; (2) tenggat tiba (N-30), satu kali, `kabar`, tunduk jam tenang, tidak dihitung dalam plafon harian; (3) pengingat mingguan sesudah tenggat, atau sejak T02 bila tanpa tenggat, hanya sebagai baris ringkasan, paling sering sekali seminggu, tanpa batas akhir selain penutupan tugas; (4) pengingat ulang pemohon sesudah `diterima`, `perlu_tanggapan`, dengan batas tiga dihitung seumur objek.

**Kriteria penerimaan**
- **AC-1** — *Given* B menerima tugas bertenggat 15 Desember dan memakai jam tenang bawaan 22.00 sampai 06.00, *When* 15 Desember pukul 00.00 tiba menurut zona waktu rumah tangga, *Then* N-30 `kabar` dijadwalkan, ditahan jam tenang, dan dilepas kepada B pukul 06.00, satu kali, tanpa dihitung dalam plafon harian; sesudahnya B mendapat baris pengingat di ringkasan paling sering sekali seminggu sampai B menandai selesai, melepas, atau A membatalkan (06 §13.1, §11.1, N-30).
- **AC-2** — *Given* tugas yang diterima B, *When* A mencoba menandai selesai, *Then* ditolak; *When* B melepasnya dengan atau tanpa alasan, *Then* `dilepas` dan A menerima `kabar`.
- **AC-3** — *Given* tugas tanpa tenggat, *When* waktu berjalan, *Then* tidak ada kedaluwarsa dan tidak ada transisi `system`; baris pengingat mingguan di ringkasan B berjalan sejak T02.
- **AC-4** — *Given* B menandai selesai (T05) pada 14 Desember, *When* 15 Desember tiba, *Then* N-30 tidak dikirim dan baris mingguan berhenti.
- **AC-5** — *Given* A sudah mengirim dua pengingat ulang selama `diajukan`, *When* B menerima tugas lalu A mengirim pengingat ulang sesudah `diterima`, *Then* yang ketiga diterima dan yang keempat ditolak, karena batas tiga dihitung seumur objek (06 §13.1 baris 4).
- **AC-6** — *Given* tugas yang diterima B tampil sebagai kartu di chat, *When* B mengetuk "selesai" atau "lepas" pada kartu, *Then* server menjalankan T05 atau T06 yang sama dengan dari layar tugas; pesan "sudah beres" yang membalas kartu tidak mengubah state (06 §16.4, INV-A44).

**Dua arah (F-REQ-013).** Simetris per peran; pemohon dan penerima tugas dapat siapa pun dari keduanya.

---

## Kesenjangan: [Perlu aturan]

Kebutuhan yang aturannya belum ada di 06 maupun 07. Default yang diusulkan dipakai dalam AC dengan tanda **[Perlu aturan]** sampai pemilik aturannya menjawab.

Pengenal butir mengikuti skema README `PA-REQ-n` (fase D; sebelumnya "butir n" atau "permintaan n").

### Yang sudah dijawab (06 revisi kesembilan dan kesepuluh)

| ID | Kebutuhan | Di mana | Jawaban | Kedudukan |
|---|---|---|---|---|
| PA-REQ-1 | Guard tanggal "tidak lampau" pada versi baru (B02 sampai B05, B11) | FR-REQ-006/AC-5, AC-6; FR-REQ-009/AC-7; FR-REQ-012/AC-7 | 06 §4.3 "Guard tanggal dan batas bawah pada versi baru" | **[Usulan]** `architect`, mengikuti default di sini |
| PA-REQ-2 | Batas bawah batas otorisasi pada B11 | FR-REQ-013/AC-4 | 06 §4.3 baris B11: > 0, boleh di bawah realisasi | **[Usulan]** `architect` |
| PA-REQ-3 | Saat pemicu pengingat "sehari sebelum" bagi tanggal tanpa jam | FR-REQ-025/AC-6 | 06 §11.1 "Waktu pengingat yang berpatokan tanggal": 00.00 hari sebelumnya, lalu jam tenang | **[Usulan]** `architect` |
| PA-REQ-4 | Pengingat permintaan tindakan | FR-REQ-045/AC-1, AC-3 sampai AC-5 | 06 §13.1 tabel "Rincian pengingat"; N-30 | **[Usulan]** `architect` |
| PA-REQ-5 | "Usul lain" pada kartu kesepakatan bila K01 penanggap ditolak guard sesudah K03 berlaku | FR-REQ-027/AC-8; sama dengan PA-COM-4 | 06 §16.4 baris "Usulan kesepakatan" (revisi kesepuluh, CA-05): K03 tetap; usulan semula berakhir `ditolak` dan tidak hidup kembali; tidak ada usulan pengganti; penanggap diberi tahu hasilnya dan boleh mengirim K01 lain yang sah; N-12 atas K03 seperti biasa | **[Usulan]** `architect` |
| PA-REQ-6 | Pesan yang ditulis di layar diskusi objek tanpa kartu (misalnya entri buku rumah tangga yang tidak dicatat dari chat) | FR-REQ-041/AC-8 | 06 §16.3 "Pesan dari layar objek tanpa kartu" (revisi kesepuluh): "Pesan yang ditulis di layar objek atas objek yang terlihat oleh semua peserta tetapi belum punya kartu (misalnya entri yang dicatat dari layar catat) **tampil di chat sebagai pesan yang menempel**, dengan `subject_ref` objek itu dan **rujukan ringkas** ke objeknya yang dibaca dari state saat ini"; tanpa kartu; N-23 saja, bukan N-19. Sama dengan PA-COM-3 | **[Usulan]** `architect`, mengikuti default di sini |
| PA-REQ-7 | Rumusan pemicu pengingat otomatis kedua: 06 §10 dulu menyebut "24 jam sebelum tanggalnya", §11.1 "sehari sebelum tanggal" | FR-REQ-025 | 06 §10 (revisi kesepuluh): "sehari sebelum", yaitu 00.00 hari sebelumnya menurut zona waktu rumah tangga; isi FR-REQ-025 tidak berubah | **[Usulan]** `architect` |
| — | Bacaan "tanda kelebihan hilang" sebagai tampilan pada permintaan | FR-REQ-035/AC-6 | 06 §5.4 | Dikukuhkan `architect` |
| — | "Sejak kapan" sebagai umur objek | FR-REQ-043 | 06 §16.2 "Menunggu sejak" | Dikukuhkan `architect` |

### Yang masih terbuka

Tidak ada. **PA-REQ-8** (sisi alur mengakhiri ambang bersama, dibuka sesudah checkpoint S4) **ditutup**: 06 §6.3 K01 revisi keempat belas menerima usulan `rule` "tanpa ambang" (mengakhiri ambang lewat kesepakatan semua pihak, K02, K05) dan nilai > 0 sesudahnya (membuat ambang baru); `rule` tetap tidak pernah kosong. Sisi uangnya 07 §3.8 "Ambang hanya lewat kesepakatan", INV-40. Dokumen 17 Q-207 ditutup.
