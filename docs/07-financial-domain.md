# 07 — Domain Keuangan

| | |
|---|---|
| Pemilik dokumen | `finance-domain` |
| Tahap | S2 (Fondasi domain) |
| Status | Draf untuk tinjauan `critic`; direvisi menurut keputusan pemilik produk 11 November 2026 dan permintaan `architect` (dokumen 06 §15.4); direvisi lagi menurut temuan pemeriksaan awal `security-privacy` atas dokumen 06 (SP-01, SP-05, SP-12, SP-13, SP-18); diselaraskan terakhir dengan dokumen 06 §15.4 butir (e) sampai (m) dan dengan pernyataan pemilik produk 11 November 2026 bahwa kedua pasangan berpenghasilan; direvisi menurut keputusan pemilik produk 11 November 2026 tentang jatah pribadi bulanan dan hari mulai periode anggaran (§3.4, §6.1, CTH-39); direvisi menurut tinjauan `critic` atas S2 (label keputusan lawan usulan pada jatah pribadi dan hari mulai, rujukan S2-1 butir (i) sampai (iv), fungsi `periode()`, andaian ambang CTH-06, CTH-33, CTH-34, CTH-39) dan temuan `security-privacy` SP-24 dan SP-28, tanpa satu angka contoh pun yang berubah; **direvisi menurut jawaban pemilik produk di checkpoint S2** (revisi yang ia setujui; ditinjau `critic` dan angkanya dihitung ulang lead dengan skrip, 12 November 2026): jatah pribadi menjadi batas (S2-1 butir (iii) pilihan b), INV-27 dan hadiah pribadi diputuskan, dompet campuran diperlakukan sebagai kasus biasa (S2-6), §8 menjadi keputusan sementara (S2-8), contoh baru CTH-40 dan CTH-41, butir baru T-28 dan T-29; **direvisi kedua kalinya sesudah checkpoint S2** menurut tiga rumusan lead yang dikukuhkan pemilik produk pada 11 November 2026 dan permintaan dokumen 06 revisi ketujuh, §15.4 butir (t) dan (u) (ditinjau `critic` dan tandanya dihitung ulang lead dengan skrip, 12 November 2026): T-29 ditutup dengan pilihan 1, setiap tagihan berulang disepakati saat dibuat dan jumlah yang disepakati menjadi batas, tanda punya tiga dasar (§3.8), tanpa satu angka saldo pun yang berubah; **putaran perbaikan 12 November 2026** menurut tinjauan `critic` (B2, S3, S4, S7, M4, M5), temuan `security-privacy` SP-32 sampai SP-40, dan empat jawaban pemilik produk (ditinjau ulang `critic` tanpa temuan menghalangi, temuannya diperbaiki tanpa tinjauan ketiga; tanda dihitung ulang lead dengan skrip sesudah perbaikan terakhir, semuanya cocok, 12 November 2026): bagian entri tertaut di atas batas otorisasi dibandingkan dengan ambang, bawaan pembagian selisih dompet campuran dicabut, jatah yang diakhiri dinilai sampai akhir periode hanya bila di atas ambang, tanpa satu angka saldo pun yang berubah; **putaran penyebaran S3 (12 November 2026)** atas celah aturan uang yang ditemukan dokumen 05 (PA-FIN-2 sampai PA-FIN-8, PA-FIN-10, G-01, G-02, G-03, G-05, G-07, PA-FAM-7) dan dokumen 11 (angkanya dihitung ulang lead dengan skrip, 12 November 2026, semuanya cocok): aturan baru seluruhnya **[Usulan]**, contoh baru CTH-42 sampai CTH-49, INV-31 dan INV-32, T-30 sampai T-38, tanpa satu angka saldo contoh lama pun yang berubah; **putaran perbaikan S3 fase D (12 November 2026)** menurut tinjauan `critic` A (CA-11) dan pemeriksaan `security-privacy` (SP-48 sampai SP-51), PA-FIN-11, dan 06 §6.3 "Pos rutin" (g) (ditinjau ulang `critic`; angkanya dihitung ulang lead dengan skrip, 12 November 2026: 58 tabel, 316 baris, 0 kesalahan); **dicatat menurut jawaban pemilik produk di checkpoint S3 (12 November 2026)**: jawabannya "lanjut", sehingga S3-1 sampai S3-9 pilihan A dan D-56 sampai D-97 *diterima lewat "lanjut"*; IDEA-03 diadopsi (§3.15, CTH-50); T-28, T-30 sampai T-40 ditutup, lalu T-34 dibuka ulang; **sesudah checkpoint S3, menurut pernyataan pemilik produk 12 November 2026**: THR yang diberikan mendapat aturan **[Usulan]** di §3.2 (CTH-51, T-42), dan sumber uang masuk menjadi daftar bawaan yang dapat disunting dengan dua kelompok keterlihatan (**[Usulan]** lead untuk checkpoint S4; §3.1, INV-24, T-34, T-43); tidak ada angka contoh lama yang berubah; CTH-50 dan CTH-51 baru dihitung ulang lead dengan skrip (12 November 2026: 60 tabel, 336 baris, 0 kesalahan); temuan `critic` CP-03, CP-05, CP-06, CP-09, CP-13 diperbaiki tanpa angka yang berubah. Rincian perubahan: §11 |
| Tanggal | 12 November 2026 |
| Dibaca sebelum menulis | `CLAUDE.md`; `family-os-brief.md` §1.1, §5C, §5D, §7, §16; `ENGINEERING-STANDARDS.md` §5; [00-glossary.md](00-glossary.md); [_progress.md](_progress.md); [17-open-questions-assumptions.md](17-open-questions-assumptions.md); [18-market-research.md](18-market-research.md) §2.1; untuk revisi: [06-activity-approval-architecture.md](06-activity-approval-architecture.md) §2 sampai §8, §15; ADR-001 sampai ADR-010; untuk putaran S3: [04-feature-map.md](04-feature-map.md) F-FIN-004, F-FIN-012, F-FIN-015; [05-functional-requirements/](05-functional-requirements/) `FIN-1.md` §13, `FIN-2.md` §12, `FAM.md` PA-FAM-7, dan FR yang memunculkan tiap celah; [11-innovation.md](11-innovation.md) §10.1; [02-product-principles.md](02-product-principles.md) PRN-05; untuk putaran perbaikan S3: [_progress.md](_progress.md) "Tinjauan `critic` A", "Tinjauan `critic` B", "Pemeriksaan `security-privacy` atas 05 dan 06"; [README.md](README.md) tabel pengenal; 06 §6.3 "Pos rutin" (g); `FIN-1.md` FR-FIN-030/AC-11 dan PA-FIN-11; 17 S3-3, S3-4, S3-9; untuk checkpoint S3: [_progress.md](_progress.md) "Checkpoint S3 dijawab"; 17 S3-1 sampai S3-10, D-56 sampai D-97, Q-140 sampai Q-182 |

Aturan uang Family OS, termasuk jumlah pada *penyelesaian*, kanonis di dokumen ini. Dokumen lain mengutipnya. Istilah mengikuti [glosarium bagian 7](00-glossary.md#7-istilah-keuangan), yang difinalkan bersama dokumen ini.

**Label.** Keputusan pemilik produk di checkpoint S1 dikutip sebagai "keputusan pemilik nomor *n*", dengan nomor menurut dokumen 17 (bagian "Yang diputuskan pemilik produk sendiri"); default yang ia terima lewat "lanjut" dikutip dengan nomor `Q`-nya. Keputusannya sesudah checkpoint itu, dalam percakapan dengan lead, dikutip sebagai "keputusan pemilik produk, 11 November 2026". Semua rekomendasi lain di dokumen ini adalah **[Usulan]** `finance-domain` kecuali diberi label lain.

**Label sesudah checkpoint S2 (11 November 2026).** Enam kedudukan dibedakan dan tidak pernah dipertukarkan; sumbernya [_progress.md](_progress.md) bagian "Jawaban checkpoint S2":

| Label | Arti |
|---|---|
| "keputusan pemilik produk, checkpoint S2" | Dinyatakan pemilik produk sendiri: S2-1 butir (i), (ii), dan (iii) pilihan b; S2-3; bagian pembayaran rutin S2-4; S2-7; S2-8; S2-10; dan fakta rumah tangga yang ia sampaikan (S2-4, S2-6) |
| "diterima pemilik produk di checkpoint S2 atas saran lead" | Ia menyerahkannya kepada lead dan lead menerapkan bawaan tim: S2-1 butir (iv) dan (v), S2-2, S2-5, S2-9. Bukan pernyataannya sendiri |
| "tidak disebut pemilik produk; diterapkan lead" | Bagian ambang S2-4 (ambang hanya berubah bila berdua setuju; tidak ada yang keluar sendiri): ia tidak menyebutnya; lead menerapkan pilihan A dan menyampaikannya kepadanya. Bukan keputusannya dan bukan penerimaannya |
| "diterima pemilik produk atas saran lead, 12 November 2026" | Jawabannya "ya bener gas" atas tiga dari empat pertanyaan lead yang lahir dari temuan `security-privacy` SP-32 sampai SP-43: tanda dasar jatah dan dasar pos rutin tanpa pemberitahuan langsung; *konfirmasi susulan* tanpa pengingat otomatis, paling banyak satu pengingat ulang, dan tidak dapat diminta lagi sesudah "belum sepakat"; selisih dompet campuran dipilih pemegang setiap kali, tanpa bawaan. Penutupan otorisasi konfirmasi susulan di dalam transaksi persetujuan **tidak** termasuk: ia **[Usulan]** `security-privacy` (SP-38) dan dokumen ini. Pagar "jatah yang tidak di atas ambang hanya diusulkan pemiliknya" **tidak** termasuk: ia **[Usulan]** `security-privacy`, disampaikan kepada pemilik produk dan tidak ia bantah |
| "default yang diterima di checkpoint S2 lewat 'lanjut'" | D-01 sampai D-55 di dokumen 17. Tidak satu pun ia nyatakan sendiri. Aturan dokumen ini yang tercantum di daftar itu dan masih tertulis **[Usulan]** dibaca demikian: usulan yang diterima sebagai default, dapat ia ubah di checkpoint mana pun |
| "rumusan lead yang dikukuhkan pemilik produk, 11 November 2026" | Tiga penegasan sesudah revisi pertama dokumen ini ([_progress.md](_progress.md), "Tiga penegasan sesudah revisi dokumen 07"): T-29 pilihan 1; tambahan satu bulan di atas jatah lewat permintaan biasa; setiap tagihan berulang disepakati saat dibuat dan jumlahnya menjadi batas. Kalimatnya disusun lead dan ia kukuhkan dengan "bener"; **bukan kata-katanya sendiri** |

Usulan yang **baru** lahir di revisi ini, dan karena itu belum pernah disajikan kepadanya, ditandai "**[Usulan]**, revisi pasca-checkpoint S2". Usulan yang lahir di putaran penyebaran S3 (12 November 2026) ditandai "**[Usulan]**, putaran S3"; semuanya belum pernah disajikan kepadanya sampai checkpoint S3, dan kedudukannya sesudah checkpoint itu ada di tabel berikut. Angka contoh selain angka brief adalah **[Asumsi]**: pemilik produk tidak memberi data nyata.

**Label sesudah checkpoint S3 (12 November 2026).** Sumbernya [_progress.md](_progress.md) bagian "Checkpoint S3 dijawab": lead menyatakan bahwa "lanjut" berarti semua bawaan diterima, dan jawaban pemilik produk, kata demi kata, adalah "lanjut". Kedudukan ketujuh karenanya "**diterima lewat "lanjut" di checkpoint S3 (12 November 2026)**": usulan yang diterima sebagai bawaan, **bukan** pernyataan pemilik produk sendiri, dan dapat ia ubah di checkpoint mana pun. Pengusulnya tetap disebut. Angka yang **[Asumsi]** tetap **[Asumsi]**. Aturan di bawah yang di bagiannya masih tertulis **[Usulan]**, putaran S3 atau putaran perbaikan S3, dibaca menurut tabel ini.

| Kedudukan | Butir dokumen ini |
|---|---|
| Diterima lewat "lanjut" di checkpoint S3, pertanyaan S3-1 sampai S3-9 pilihan A | **S3-1:** (b) satu baris "disisihkan untuk tabungan" di halaman depan, di luar angka utama; tidak ada aturan uang baru: angka utama tetap *uang yang ada* (§5) dan alokasi tetap tanpa entri (§3.9). (c) tidak ada `GOL` lanjutan di pilot; tingkat tabungan tetap ditunda (§10.4). (d) pengingat tagihan N-32 Tier 1 per *kejadian tagihan* (usulan `architect`); definisi kejadian dan "belum dibayar" yang dipakainya (§3.15, PA-FIN-5) ikut diterima bersamanya dan bersama D-97. **S3-2:** (i) IDEA-01, bentuk uangnya §3.10 "Menghitung lalu merinci seluruh selisih sekaligus" dan CTH-48; (ii) IDEA-03, diadopsi di §3.15 "Tagihan berulang tanpa `due_rule`"; (iii) IDEA-02 adalah isian bawaan layar: entri yang tersimpan sama dengan yang diisi tangan, jadi tidak ada aturan uang baru. **S3-3:** THR dan bonus dari pemberi kerja bersumber `salary` (§3.1, T-34); **dibuka ulang sebagian sesudah checkpoint** oleh pertanyaan pemilik produk sendiri (baris terakhir tabel ini; T-43); pertanyaan sampingan "apakah kalian menerima THR?" tidak dijawab di checkpoint itu; jawabannya sesudahnya ada di baris terakhir tabel ini. **S3-4:** (c) masa tenggang keluar (usulan `security-privacy` SP-48; 7 hari **[Asumsi]**, A-39): §8.2 butir 1a, 2b, INV-28, T-39, CTH-38; (a) dan (b) tidak punya akibat uang. **S3-5:** (i) satu orang boleh mengakhiri kesepakatan pos rutin sendirian (dokumen 06 K06); (ii) §3.15 aturan 5; (iv) tautan susulan ke tagihan berulang hanya dari pencatat entri (§3.8, INV-26). **S3-6:** jalur *konfirmasi susulan* (§3.8 "Penautan" butir 1), termasuk (iv) hanya pencatat entri yang dapat memintanya. **S3-9:** (i) hitungan bukan pemegang dengan batas SP-50 (§3.10 butir 1, `others_may_count`, INV-33, T-33, CTH-47); (ii) penghapusan piutang dinilai kumulatif per piutang (`H`, §3.7, §3.8, INV-34, T-32, CTH-25 varian SP-51) |
| Diterima lewat "lanjut" di checkpoint S3, bawaan D-56 sampai D-97 dokumen 17 | D-75 (§2.5: ajakan menghitung hanya saat saldo negatif); D-76 (§10.3 butir 10); D-77 (§3.4, INV-30: jatah 0 berbeda dari tanpa jatah; jatah yang tidak di atas ambang hanya diusulkan pemiliknya dan berakhir seketika, SP-32 sampai SP-34); D-78 (§2.3); D-79 (§3.9, §3.15, T-28); D-80 (§3.15 aturan 7a); D-81 (T-29: batas masa penilaian menurut waktu simpan server); D-85 (§2.6, §3.2, §3.9, §3.14, §3.15: T-10, T-30, T-31, T-35, T-36, T-38); D-86 (§8.2 butir 2a, 2b, T-37, T-39); D-89 (§3.1, INV-24); D-90 (§8.2 butir 3a, usulan `security-privacy` SP-49); D-97 (§3.15 "Urutan bawaan kejadian", T-40) |
| Tetap **[Usulan]**, belum pernah disajikan kepada pemilik produk | Tidak ada di daftar D dan tidak ditanyakan di S3-1 sampai S3-9, jadi tidak tercakup "lanjut": PA-FIN-2 (saldo awal oleh bukan pemegang bertanda, §2.5 baris (a), §3.17); PA-FIN-3 (INV-31); PA-FIN-8 (§3.10 butir 8); PA-FIN-10 (INV-32); G-02 (selisih dihitung server, §3.10 "Hitungan yang tiba belakangan"); pengarsipan jadwal penyisihan oleh pemilik jatah (§2.6, mengikuti dokumen 06 §6.3 "Pos rutin" (g)); penutupan otorisasi konfirmasi susulan di dalam transaksi persetujuan dengan batas sebesar entri dan tepat satu tautan (usulan `security-privacy` SP-38 dan dokumen ini; S3-6 tidak menyebutnya); §3.15 aturan 6 dan 7 (usulan `architect`). Lahir sesudah checkpoint S3: §3.15 baris "Ada-tidaknya `due_rule`" (T-41), §3.2 "THR dan angpau yang diberikan" (T-42, CTH-51), dan §3.1 "Sumber uang masuk" (dua kelompok: bawaan lead untuk checkpoint S4; siapa boleh mengubah: dokumen ini; T-43) |
| Tidak dijawab di checkpoint S3; dijawab sesudahnya | Pertanyaan sampingan S3-3: pemilik produk menulis "thr bisa menerima bisa ngasih", dan atas pertanyaan lead kepada siapa THR diberikan, "bro pikirkan kalau semua orang pakai app ini" (pernyataannya sendiri; [_progress.md](_progress.md), 12 November 2026). Bacaan lead bahwa rumah tangga menerima dan memberi THR belum dikukuhkan. THR yang diberikan: §3.2 (**[Usulan]** `finance-domain`). Atas S3-3 sendiri ia bertanya "kenapa thr dicatat sebagai gaji? kan pemasukan kita yang atur darimana nya?" dan menulis "ada kategori default bisa edit tambah atau yang lainnya" (kata-katanya sendiri; bahwa "kategori" berarti sumber uang masuk adalah bacaan lead **[Usulan]**): sumber uang masuk adalah daftar bawaan yang dapat disunting dan ditambah, dengan "lainnya". S3-3 A dibuka ulang sebagian; THR yang diterima kini sumber bawaan "THR" di kelompok `work` (**[Usulan]** lead, untuk checkpoint S4; §3.1 "Sumber uang masuk", T-34, T-43). S3-10 dijawab sebagian sesudahnya dan tidak menyentuh aturan uang |

**Dua lapis.** Dokumen ini memakai bahasa model internal (entri, baris, sebab, kewajiban). Keluarga tidak pernah melihat kata-kata itu; padanan layarnya ada di §1.4 dan ditetapkan dokumen 08.

**Kesetaraan.** Semua aturan ditulis untuk "anggota dewasa". Kata "suami" dan "istri" hanya muncul di contoh, karena brief §1.1 menggambarkan rumah tangga pilot dengan kata itu.

## Daftar isi

1. Keputusan ledger
2. Model
3. Aturan per jenis peristiwa
4. Invarian
5. Tampilan custody
6. Anggaran dan laporan
7. Contoh terhitung
8. Riwayat yang tidak dapat diubah lawan hak penghapusan
9. Keputusan atas pertanyaan terbuka
10. Butir yang masih terbuka dan permintaan ke pihak lain
11. Riwayat perubahan

---

## 1. Keputusan ledger

### 1.1 Yang dituntut rumah tangga ini

| # | Kebutuhan | Asal |
|---|---|---|
| N1 | Enam dimensi custody terjawab terpisah untuk setiap jumlah uang | Brief §5C |
| N2 | Uang yang berpindah antar-orang dan antar-tempat tidak pernah terhitung dua kali | Brief §5C, §7 |
| N3 | Koreksi hanya dengan menambah catatan | Brief §7; `ENGINEERING-STANDARDS.md` §5 |
| N4 | Pencatat pelupa: entri terlambat, selisih kas, tanpa penutupan periode | Brief §1.1; keputusan pemilik nomor 7 |
| N5 | Invarian dapat dijalankan sebagai property test | `ENGINEERING-STANDARDS.md` §5, §6 |
| N6 | Tim 1–3 developer; satu VPS | Brief §1 |
| N7 | Utang, piutang, cicilan, aset, dan kekayaan bersih ada di Tier 1 | Brief §1.1 |
| N8 | Dana pribadi tertutup, tetapi kontribusinya ke rumah tangga terlihat, dan penghasilan yang diarahkan ke dana pribadi saat dicatat terlihat beserta jumlahnya | Keputusan pemilik nomor 3; keputusan pemilik produk, 11 November 2026 |

### 1.2 Perbandingan

| Kriteria | A. Double-entry penuh | B. Ledger sederhana: entri bertipe dengan baris seimbang | C. Daftar transaksi polos |
|---|---|---|---|
| Bentuk | Bagan akun bebas; setiap jurnal berisi debit dan kredit ke akun mana pun | Jenis entri tertutup (§2.5). Tiap entri punya baris *posisi* (tempat uang, piutang, kewajiban, aset) dan baris *sebab* (pemasukan, pengeluaran, selisih kas, dan lain-lain) yang jumlahnya harus sama | Satu baris per transaksi: jenis, jumlah, satu atau dua tempat uang |
| N1 custody | Bisa, tetapi pemegang dan pemilik harus dikodekan sebagai akun; bagan akun membengkak | Pemegang dan pemilik adalah atribut tempat uang; tujuan, penerima manfaat, pelaku adalah atribut baris | Bisa untuk pemegang; tujuan dan pemilik ditempel sebagai kolom tambahan |
| N2 tanpa hitung ganda | Terjamin oleh keseimbangan jurnal | Terjamin oleh keseimbangan baris (INV-02) | Tidak terjamin: pemindahan adalah dua baris atau satu baris berpengecualian; kesalahannya baru terlihat di laporan |
| N3 koreksi | Jurnal pembalik | Entri pembalik: baris yang sama dengan tanda terbalik | Baris negatif; arti "membalik" berbeda untuk tiap jenis |
| N4 pencatat pelupa | Selisih kas adalah jurnal ke akun selisih | Selisih kas adalah jenis entri sendiri; entri terlambat dapat menutupnya tanpa menyentuh saldo | Selisih kas bisa; "menutup selisih" memerlukan aturan khusus |
| N5 invarian | Satu invarian umum, tetapi aturan domain (misalnya "talangan tidak dihitung dua kali") tetap harus ditulis di atasnya | Satu invarian umum ditambah aturan per jenis entri; keduanya dapat diuji | Tidak ada invarian struktural; setiap laporan membawa aturannya sendiri |
| N7 kekayaan bersih | Neraca langsung | Jumlah baris posisi; sama dengan jumlah baris sebab (INV-07) | Dihitung dari beberapa tabel terpisah; mudah tidak cocok |
| N8 dana pribadi | Perlu buku per pemilik | Tiap baris membawa pemilik; keseimbangan diperiksa per pemilik | Kolom pemilik; tanpa pemeriksaan |
| N6 biaya | Tertinggi: bagan akun, pemetaan setiap layar ke jurnal, risiko kosakata akuntansi bocor ke layar | Sedang: satu tabel entri, satu tabel baris, daftar jenis yang tertutup | Terendah di awal, tertinggi kemudian: setiap fitur Tier 1 (cicilan, hak ganti, aset) menambah pengecualian |

### 1.3 Rekomendasi

**Opsi B [Usulan].** Secara matematis B adalah double-entry: setiap entri seimbang. Bedanya dari A, tidak ada bagan akun yang dapat diatur: "akun" hanya ada dalam dua bentuk yang sudah dimiliki keluarga, yaitu *posisi* (tempat uang, piutang, kewajiban, aset yang mereka daftarkan sendiri) dan *sebab* (sembilan nilai tetap di §2.4).

Alasan yang terikat pada rumah tangga ini:

1. Arus uang pilot didominasi perpindahan, bukan belanja: gaji masuk rekening, tarik tunai, serah terima ke pasangan, pinjam kembali (brief §1.1). Daftar transaksi polos paling lemah justru di sini (N2), dan kriteria sukses nomor 1 bergantung padanya.
2. Tier 1 memuat cicilan, pinjaman ke keluarga besar, hak ganti, aset, dan kekayaan bersih (N7). Semua itu adalah posisi yang bukan uang. Pada opsi C tiap posisi menjadi tabel dengan aturan sendiri; pada opsi B semuanya baris dari entri yang sama dan tunduk pada satu invarian.
3. Dana pribadi tertutup tetapi kontribusinya terlihat (N8). Itu menuntut satu peristiwa tercatat di dua buku pemilik tanpa terhitung dua kali; keseimbangan per pemilik (INV-02, INV-03) menjaminnya.
4. Tim kecil (N6): opsi A menambah bagan akun dan pemetaan jurnal yang tidak dipakai siapa pun di rumah tangga ini.

Yang tidak diberikan opsi B dan diterima: akun bebas buatan pengguna, laporan laba-rugi dan neraca berformat akuntansi, multi-mata-uang dalam satu entri. Tidak satu pun dibutuhkan pilot; yang terakhir dapat ditambahkan kemudian karena tiap baris membawa mata uangnya.

### 1.4 Yang dilihat keluarga lawan model internal

Tidak ada kata di kolom kiri yang tampil di layar. Kolom kanan adalah usulan untuk dokumen 08, yang menetapkan kata akhirnya.

| Model internal | Yang dilihat keluarga (usulan) |
|---|---|
| Entri, baris posisi, baris sebab | "Catatan" |
| Tempat uang | "Dompet", "Rekening", "E-wallet" beserta nama pemegangnya |
| Pemindahan, serah terima | "Pindah uang", "Kasih ke …", "Ambil dari …" |
| Pemasukan | "Uang masuk" |
| Pengeluaran | "Belanja", "Bayar" |
| Kewajiban | "Utang kita", "Sisa cicilan" |
| Piutang | "Dipinjam …" |
| Hak ganti | "Minta diganti" |
| Kontribusi dari dana pribadi | "Pakai uang pribadi …" |
| Selisih kas | "Selisih hitung" |
| Entri pembalik, entri penyesuaian | "Dibatalkan", "Diperbaiki" |
| Pelepasan klaim dana bersama atas seorang anggota | "Utang dihapuskan" |
| Penutupan klaim saat anggota keluar | "Ditutup karena anggota keluar" |
| Otorisasi, terikat | "Disetujui", "Sudah dijanjikan" |
| Reservasi | "Disisihkan untuk …" |
| Alokasi (label tujuan pada uang yang ada, misalnya untuk tagihan yang jatuh temponya menyusul; §3.9, §3.15) | "Disiapkan untuk …" (**[Usulan]**, revisi pasca-checkpoint S2). Dokumen 08 perlu membedakan tiga hal yang kata sehari-harinya berdekatan: alokasi, reservasi, dan "disisihkan ke uang pribadi" |
| Jatah pribadi (`set_aside_schedule`) | "Jatah pribadi" |
| Periode anggaran | "Bulan ini", beserta rentang tanggalnya |
| Kekayaan bersih | "Harta bersih" |
| Uang yang ada | "Uang kita" |

### 1.5 Yang perlu dicatat `architect` sebagai ADR

Ketujuh keputusan di bawah sudah ditulis `architect` sebagai ADR-002 sampai ADR-008, berurutan (ADR-A = ADR-002, …, ADR-G = ADR-008). Kalimat ADR yang perlu disesuaikan setelah revisi dokumen ini didaftar di §10.2.

| # | Keputusan untuk ADR | Bagian |
|---|---|---|
| ADR-A | Model ledger: opsi B, entri bertipe dengan baris seimbang; jenis entri dan sebab adalah daftar tertutup | §1.3, §2 |
| ADR-B | Dana pribadi berada di luar *garis kendali rumah tangga*; satu buku per *pemilik uang*; keseimbangan per pemilik | §2.2 |
| ADR-C | Saldo adalah turunan dari baris; bila disimpan sebagai cache, diperbarui dalam transaksi basis data yang sama dengan entrinya | §2.5, INV-04 |
| ADR-D | Idempotency key per perintah tulis, unik per rumah tangga, disimpan bersama entri | §2.5, INV-10 |
| ADR-E | Tidak ada penutupan periode; laporan periode lampau dihitung ulang dan menandai perubahan | §3.11 |
| ADR-F | Tafsir "append-only" terhadap hak penghapusan: inti moneter tetap dan dipseudonimkan, rincian pribadi terpisah dan dapat dihapus. Keputusan sementara sampai L-04 dijawab | §8 |
| ADR-G | Jumlah uang: bilangan bulat bertanda 64-bit dalam satuan terkecil; eksponen satuan per mata uang (IDR = 0) | §2.1 |

---

## 2. Model

Sketsa di bagian ini adalah sketsa skema, bukan kode. Nama bahasa Inggris adalah usulan nama pengenal.

### 2.1 Jumlah uang

- `amount`: bilangan bulat bertanda dalam satuan terkecil mata uangnya, selalu bersama `currency`.
- Satuan terkecil IDR adalah **rupiah utuh** (eksponen 0). Saldo rekening bank yang memuat sen dicatat dibulatkan ke bawah ke rupiah utuh; sisa sen tidak dicatat. Bila kelak mengganggu, ia tertangkap sebagai *selisih kas* **[Usulan]**.
- Tidak ada pembagian yang menghasilkan pecahan di dalam model. Pembagian satu jumlah ke beberapa baris (pengeluaran terbagi) diisi sebagai jumlah per baris oleh pencatat; bila aplikasi menawarkan "bagi rata", sisa pembagian ditaruh pada baris pertama, sehingga jumlah baris selalu sama dengan jumlah entri.
- MVP memakai satu mata uang per rumah tangga. Entri yang memuat dua mata uang ditolak.

### 2.2 Pemilik uang, buku, dan garis kendali rumah tangga

`money_owner` adalah salah satu dari:

| Nilai | Arti |
|---|---|
| `shared` | Rumah tangga secara bersama (*dana bersama*) |
| `member(M)` | Satu anggota dewasa (*dana pribadi* M) |

Setiap pemilik uang punya satu **buku**: kumpulan semua baris yang bertanda pemilik itu. Buku `shared` adalah **buku rumah tangga**.

**Buku pribadi per anggota, bukan per akun [Usulan]** (dikukuhkan atas permintaan dokumen 06 §15.4 butir (c); dokumen 06 §14; ADR-003). `member(M)` berarti *anggota*, yaitu keanggotaan seseorang dalam **satu** rumah tangga, bukan *akun* login-nya; setiap buku, entri, dan baris milik tepat satu rumah tangga (`entry.household_id`, §2.5), dan semua invarian §4 diperiksa di dalam satu rumah tangga. Akibatnya, orang yang ikut dua rumah tangga punya dua buku pribadi yang tidak saling tahu, tanpa entri lintas rumah tangga: uang pribadinya yang secara fisik satu dicatat sendiri-sendiri di tiap rumah tangga tempat ia memakainya, dan baris kedua "uang pribadi saya" (§5) hanya memuat buku pribadi rumah tangga yang sedang dibuka.

**Uang antara dua rumah tangga yang sama-sama memakai aplikasi [Usulan]** (Q-86; dokumen 03 bagian 4.3). Dua catatan yang berdiri sendiri, tanpa tautan. Rumah tangga pengirim mencatatnya menurut aturannya sendiri dengan penerima sebagai *pihak luar* (kiriman keluarga, atau `loan_given`); rumah tangga penerima mencatatnya menurut aturannya sendiri dengan pengirim sebagai *pihak luar* (`income` bersumber `gift` atau `other`, atau `loan_received`). Tidak ada entri lintas rumah tangga, tidak ada pengenal bersama, tidak ada pencocokan otomatis, dan tidak ada notifikasi dari satu rumah tangga ke rumah tangga lain; invarian §4 tiap rumah tangga diperiksa sendiri, dan bila kedua catatan berbeda jumlah atau salah satunya tidak ada, tidak ada yang mendamaikannya. Ini berlaku juga bila satu orang menjadi anggota keduanya. Alasannya isolasi rumah tangga: tautan apa pun memberi tahu satu rumah tangga bahwa, kapan, dan berapa rumah tangga lain mencatat. Bentuk yang lebih kaya (misalnya penerima mengonfirmasi kiriman) ditunda ke tahap publik (Tier 3) dan memerlukan tinjauan `security-privacy`.

**Keputusan: dana pribadi berada di luar garis kendali rumah tangga [Usulan].** Di dalam garis hanya dana bersama. Akibatnya:

| Hal | Akibat |
|---|---|
| *Uang yang ada* | Jumlah saldo tercatat semua tempat uang berpemilik `shared`. Inilah angka utama, dan sama bagi kedua pasangan. Tidak perlu nama kedua (menutup K-13) |
| *Kekayaan bersih* rumah tangga | Hanya dari buku rumah tangga |
| Dana bersama → dana pribadi | Uang melintas keluar garis: sebab `owner_transfer` (*penyisihan*), bukan pemindahan |
| Dana pribadi → dana bersama, atau membayar keperluan rumah tangga | Uang atau nilai melintas masuk: sebab `owner_transfer` (*kontribusi dari dana pribadi*) |
| Penghasilan (gaji, hasil usaha) yang diarahkan pencatatnya ke dana pribadinya saat dicatat | Pemasukannya tetap tercatat penuh di buku rumah tangga; bagian pribadinya keluar garis dalam entri yang sama sebagai `owner_transfer` (penyisihan), sehingga terlihat oleh keduanya (§3.1, INV-24). Hanya bila pencatatnya penerima penghasilan itu (INV-27; keputusan pemilik produk, checkpoint S2) |
| Klaim antara anggota dan dana bersama | Tampil di buku rumah tangga sebagai kewajiban atau piutang rumah tangga kepada anggota itu |
| Buku pribadi | Tetap berada di ledger yang sama dan tunduk pada invarian yang sama; hanya terlihat oleh pemiliknya (§5) |

Alasan:

1. Keputusan pemilik nomor 3: saldo dana pribadi tertutup bagi pasangan. Angka rumah tangga yang memuat dana pribadi tidak dapat ditampilkan sama kepada keduanya tanpa membocorkannya lewat pengurangan (Q-23). Dengan dana pribadi di luar garis, soal itu hilang secara struktur, bukan lewat penyaringan tampilan.
2. Keputusan pemilik nomor 3 juga menyebut pembayaran keperluan rumah tangga dengan uang pribadi sebagai **kontribusi**. Kontribusi hanya bermakna bila ada garis yang dilintasi.
3. Keputusan pemilik nomor 2: hadiah untuk satu orang menjadi dana pribadinya. Dengan dana pribadi di luar garis, hadiah itu tidak menaikkan "uang kita", sesuai maksudnya.

Biaya yang diterima: kekayaan bersih rumah tangga tidak memuat harta pribadi. "Berapa harta keluarga seluruhnya" tidak dijawab di layar mana pun, karena jawabannya hanya dapat ditampilkan dengan membuka dana pribadi. Tiap anggota melihat baris kedua, "uang pribadi saya" (§5).

**Aturan tulis buku pribadi [Usulan]** (menjawab dokumen 06 §7.3 dan §15.4 butir 4; INV-25). Perintah yang menulis baris berpemilik `member(M)` hanya diterima dari M. Pencatat lain tidak mengenal tempat uang pribadi M, dan tidak seorang pun dapat dicatat menerima, memakai, atau berutang atas dana pribadinya oleh orang lain. Akibatnya:

| Entri dua buku (buku rumah tangga dan buku pribadi M) | Dicatat oleh |
|---|---|
| Penyisihan kepada M; pemasukan dengan bagian pribadi M (§3.1) | M. Untuk penghasilan: hanya bila M juga `earner`-nya (INV-27; keputusan pemilik produk, checkpoint S2) |
| Kontribusi uang, kontribusi belanja, talangan oleh M | M |
| Penggantian dari dana bersama kepada M | M, yaitu penerimanya. Baris yang mengurangi tempat uang dana bersama yang dipegang anggota dewasa lain bertanda `unconfirmed` di sisi itu (§3.3; dokumen 06 §7.2 "Aturan umum") |
| Pinjaman antara dana bersama dan dana pribadi M, dan pelunasannya | M; tanda `unconfirmed` yang sama atas baris yang mengurangi tempat uang dana bersama yang dipegang anggota dewasa lain |
| Pelepasan klaim M atas dana bersama (hak ganti) | M, sebagai pemilik piutangnya (§3.6) |
| Pelepasan klaim dana bersama atas M | Anggota dewasa selain M. **Satu-satunya pengecualian bagi pencatat manusia**: entri ini hanya memuat baris kewajiban − dan `owner_transfer` + di buku M, tidak menyentuh tempat uang M, dan hanya menguntungkan M (§3.6) |
| Entri pembalik atas entri dua buku | M (§3.12) |
| Penutupan klaim saat M keluar (`exit_closure`) | Sistem, bukan anggota mana pun (§8.2, INV-28) **[Usulan]** |

Aturan ini sejalan dengan keputusan pemilik nomor 1 (yang menerima uang yang mencatat) dan dengan keputusan pemilik produk 11 November 2026 (yang mencatat uang masuk memilih arahnya): arah "ke dana pribadi" selalu berarti dana pribadi **pencatat sendiri**. Untuk penghasilan, keputusan pemilik produk di checkpoint S2 (S2-1 butir (i)): arah itu hanya tersedia bila pencatatnya juga penerima penghasilan itu (INV-27, §3.1). Persinggungannya dengan "kedua pasangan boleh mengoreksi entri dana bersama" diatur di §3.12.

Batas yang diterima: entri yang menyentuh buku pribadi **dua** anggota (talangan, pinjaman, atau penggantian antara dua dana pribadi; CTH-17b) tidak dapat dihasilkan satu perintah, sebab tidak ada satu pelaku yang berhak atas kedua buku. Modelnya tetap ada dan invariannya tetap diuji, tetapi alur pencatatannya **ditunda dari MVP**: arus uang rumah tangga pilot (brief §1.1) tidak memuatnya, dan angka rumah tangga tidak tersentuh olehnya. Jalan menambahkannya kelak: satu tempat uang pribadi bawaan per anggota sebagai tujuan baris masuk yang dicatat orang lain, bertanda belum dikonfirmasi (T-12).

### 2.3 Posisi

Posisi adalah sesuatu yang bersaldo. Setiap posisi punya tepat satu pemilik uang. Semuanya dibuat, diganti nama, dan diarsipkan oleh rumah tangga sendiri (keputusan pemilik nomor 9).

```text
holding                              -- tempat uang
  id, household_id
  name
  kind            cash | bank | ewallet | entrusted_external
  owner           shared | member(M)          -- tepat satu pemilik (Q-44)
  holder          member(M) | joint | external_party(P)
  visibility      private | spouses | all_members   -- ditetapkan dokumen 06 dan 10
  others_may_count  boolean, bawaan false       -- "pasanganku dapat melihat saldo rekening ini";
                  -- hanya bila owner = shared, holder = member(M), kind bank | ewallet;
                  -- diubah hanya oleh M, sebagai peristiwa tanpa entri (§3.10 butir 1)
  currency
  archived_at     null | timestamp

receivable                           -- piutang
  id, household_id, owner
  counterparty    external_party(P) | member(M) | shared
  origin          loan_given | advance | opening
  archived_at

liability                            -- kewajiban
  id, household_id, owner
  counterparty    external_party(P) | member(M) | shared
  kind            debt | installment | reimbursement_claim | inter_owner_loan | opening
  archived_at

asset                                -- aset
  id, household_id, owner
  name, kind      vehicle | gold | property | device | other
  archived_at
```

Aturan:

- **Pemegang adalah atribut tempat uang** (Q-20). Uang berganti pemegang dengan berpindah ke tempat uang lain, tidak dengan mengubah atribut `holder`. Atribut `holder` hanya dapat diubah bila saldonya nol.
- **Satu tempat uang, satu pemilik** (Q-44). Uang tunai pribadi dan uang tunai bersama yang dipegang orang yang sama adalah dua tempat uang, walau secara fisik satu dompet.
- **Tempat uang pribadi selalu dipegang pemiliknya** (**[Usulan]**, putaran S3; PA-FIN-10; INV-32). `owner = member(M)` berarti `holder = member(M)` dan `kind ≠ entrusted_external`; perintah yang membuat atau mengubah tempat uang pribadi dengan pemegang lain ditolak. Alasannya: hanya M yang menulis buku pribadinya (INV-25), jadi tidak ada orang lain yang dapat mencatat pemakaian uang itu; pemegang lain akan membuat saldo yang tidak dapat dijaga siapa pun, dan memindahkan uang pribadi ke tangan pasangan sudah punya bentuknya sendiri (penyisihan terbalik atau utang antar-pasangan, §3.4, §3.7). Uang pribadi M yang dititipkan kepada pihak luar dicatat M pada tempat uang pribadinya sendiri, dengan nama yang ia pilih; model tidak melacak pemegang luarnya. Tempat uang `entrusted_external` tetap hanya berpemilik `shared` (butir "Titipan pada pihak luar" di bawah).
- **Dompet campuran adalah kasus biasa, bukan kasus tepi** (fakta rumah tangga pilot, checkpoint S2, pertanyaan S2-6; kata-kata pemilik produk: "bisa terpisah bisa tercampur bisa di bank bisa tunai"). Andaian dokumen ini sebelumnya, bahwa uang pribadi tunai disimpan terpisah atau tidak ada, **tidak berlaku**: semua bentuk terjadi. Aturan "satu tempat uang, satu pemilik" dipertahankan, dengan perlakuan berikut. Bahwa aturan itu dipertahankan sudah disampaikan lead kepada pemilik produk; ia belum menyatakan menerima atau menolaknya.

  | Hal | Aturan |
  |---|---|
  | Pencatatan | Satu dompet fisik (atau satu rekening, atau satu dompet digital) yang memuat dana bersama dan dana pribadi pemegangnya dicatat sebagai **dua tempat uang** berpemegang sama: satu berpemilik `shared`, satu berpemilik `member(M)`. Tiap belanja dicatat pada tempat uang yang pemiliknya menanggungnya |
  | Menghitung | Pemegang menghitung **seluruh isi dompet satu kali** (`C`) dan membandingkannya dengan jumlah saldo tercatat kedua tempat uang: `selisih = C − (tercatat bersama + tercatat pribadi)` |
  | Membagi selisih | Pemegang menetapkan berapa dari selisih itu jatuh pada tiap tempat uang: `selisih = selisih bersama + selisih pribadi`. **Tidak ada bawaan.** Setiap kali selisihnya bukan nol, pemegang memilih satu dari tiga: seluruhnya pada bagian bersama, seluruhnya pada bagian pribadi, atau dibagi dengan jumlah yang ia isi; tidak ada pilihan yang terisi lebih dulu, dan hitungan tidak tersimpan sebelum ia memilih (diterima pemilik produk atas saran lead, 12 November 2026; temuan `security-privacy` SP-40 (b); T-03). Bawaan "seluruhnya pada bagian bersama" dari revisi sebelumnya **dicabut**: ia menampilkan kekurangan uang pribadi sebagai selisih uang bersama atas nama pemegang tanpa pemegang memilihnya |
  | Bentuk di model | Tidak ada jenis entri dan tidak ada entitas baru: hasilnya **dua rekonsiliasi biasa** (§3.10), satu per tempat uang, masing-masing dengan entri `cash_difference`-nya sendiri hanya bila selisihnya bukan nol, dan masing-masing tunduk pada INV-16. Saldo terhitung tiap tempat uang adalah turunan: `tercatat + selisih yang ditetapkan padanya`. Pilihan "bersama" memberi satu entri di buku rumah tangga, pilihan "pribadi" satu entri di buku pribadi, pilihan "dibagi" dua entri |
  | Yang diterima server | **Server tidak pernah menerima `C`.** Pembagian dilakukan di perangkat pemegang; yang dikirim adalah dua perintah rekonsiliasi biasa, masing-masing dengan saldo terhitung tempat uangnya sendiri dan idempotency key-nya sendiri. Di sisi yang terbaca pasangan keduanya **tidak terkait**: tidak ada pengenal bersama, medan, *peristiwa*, atau tautan di buku rumah tangga yang menyatakan bahwa rekonsiliasi bagian bersama itu bagian dari hitungan dompet campuran. Dikukuhkan atas permintaan dokumen 06 §15.4 butir (u) 9; sama dengan dokumen 06 V-9 dan INV-A41. **Ini aturan bentuk data, bukan perlindungan terhadap pengelola server** (temuan `security-privacy` SP-40 (d)): server menerima kedua saldo terhitung, dan siapa pun yang dapat membaca kedua buku di server dapat menjumlahkannya menjadi `C`. Yang dijaga aturan ini hanyalah apa yang terbaca pasangan; akses pengelola server dinilai di dokumen 10 (Q-57) |
  | Yang dilihat pasangan | Hanya rekonsiliasi tempat uang yang berpemilik `shared`, persis seperti rekonsiliasi tempat uang dana bersama lain: saldo terhitung bagian bersama dan selisihnya, juga bila selisihnya nol ("dihitung, cocok"). Tidak terbaca olehnya: `C`, jumlah kedua saldo, selisih sebelum dibagi, bagian selisih pada tempat uang pribadi, dan bahwa hitungan itu hitungan dompet campuran. Alasannya: bersama saldo terhitung bagian bersama, `C` membuka saldo dana pribadi lewat pengurangan |
  | Siapa yang mencatat | Pemegang dompet itu sendiri, karena ia pemilik bagian pribadinya (INV-25). Pasangannya dapat merekonsiliasi tempat uang yang berpemilik `shared` saja, dengan syarat yang sama dengan tempat uang dana bersama lain: hanya rekening bank atau dompet digital yang ditandai pemegangnya, dan paling banyak satu hitungan terbuka; selisihnya bertanda `unconfirmed` di sisi pemegang (§3.10 butir 1; §2.5 baris (c)). Dompet tunai campuran tidak pernah dihitung pasangannya |
  | Syarat layar (dokumen 05 dan 08) | Menghitung dompet campuran adalah **satu layar**: satu angka hitungan, dan bila selisihnya bukan nol satu pilihan pembagian selisih (bersama, pribadi, atau dibagi) **tanpa pilihan terisi**. Bahwa kedua saldo terhitung berjumlah `C` dijaga layar itu; server menerima dua rekonsiliasi biasa |

  **Biayanya, dinyatakan terus terang.** (1) Pemegang harus mengingat pembagiannya: model tidak dapat mengetahui lembar uang mana milik siapa, dan pembagian selisih adalah pernyataan pemegang yang tidak dapat diperiksa. (2) Bila pemegang memilih "bersama" padahal uangnya uang pribadi, kekurangan yang sebenarnya terjadi pada uang pribadi (belanja pribadi yang lupa dicatat) tampil sebagai "selisih hitung" dana bersama atas nama pemegangnya, terlihat pasangan, dan mengurangi *uang yang ada*; ia dapat memperbaikinya kemudian dengan entri pembalik dan catatan yang benar. (3) Bila pemegang menetapkan kekurangan pada bagian pribadi, pasangan tidak melihat apa pun dan saldo dana bersama tidak berubah: akibatnya sama dengan pemegang menanggung kekurangan itu dari uang pribadinya tanpa mencatatnya sebagai kontribusi. Sebaliknya, kelebihan hitungan yang ditetapkan pada bagian pribadi menambah dana pribadi tanpa terlihat pasangan. Model tidak menutup (3): pemegang yang sama memang selalu dapat merekonsiliasi kedua tempat uang itu sendiri-sendiri dengan angka pilihannya. (4) Tanpa bawaan, pemegang yang tidak ingat harus tetap memilih; pilihannya pernyataan, bukan fakta yang diperiksa. (5) Orang yang melihat dompet itu sendiri dapat menghitung bagian pribadinya (§5.3, B6). Contoh terhitung: CTH-40.
- **Tempat uang bersama** (`holder = joint`) tidak punya satu pemegang. Custody melaporkannya sebagai baris tersendiri, "dipegang bersama", tidak dibagi dua. `joint` menyangkut akses, bukan kepemilikan.
- **Titipan pada pihak luar** adalah tempat uang berjenis `entrusted_external` dengan `holder = external_party(P)` dan `owner = shared`. Ia terhitung dalam *uang yang ada*. Karena pihak luar tidak punya login, pemakaiannya dicatat oleh anggota; bila tidak dicatat, saldonya salah sampai direkonsiliasi. Itu sebabnya kiriman bulanan kepada ibu suami **tidak** dimodelkan sebagai titipan (default Q-07 yang diterima pemilik produk).
- **Klaim antar-pemilik** (*klaim internal*): satu klaim antara dua pemilik uang di rumah tangga yang sama tercatat sebagai sepasang posisi, `receivable` di buku satu pemilik dan `liability` di buku pemilik lain, yang lahir, berubah, dan berakhir dalam entri yang sama dan selalu sama besar (INV-08).
- Saldo piutang, kewajiban, dan aset disimpan sebagai bilangan tidak negatif. Dalam persamaan keseimbangan, kewajiban bertanda minus.

### 2.4 Sebab

Baris sebab menjawab "mengapa kekayaan bersih pemilik ini berubah". Nilainya adalah **akibat pada kekayaan bersih**: positif menaikkan, negatif menurunkan. Daftarnya tertutup.

| `cause` | Arti | Tanda lazim | Membawa |
|---|---|---|---|
| `opening` | Saldo awal | + atau − | — |
| `income` | Pemasukan | + | sumber (`source_id`, menunjuk satu *sumber uang masuk* rumah tangga, §3.1 "Sumber uang masuk"), atas nama siapa (`earner`), kategori. Pemasukan yang sumbernya berkelompok `work` ("dari pekerjaan atau usaha") disebut **penghasilan**; sumber berkelompok itu adalah **daftar sumber penghasilan** (INV-24). Bawaan kelompok `work`: gaji, THR, bonus, lembur, honor atau pekerjaan lepas, hasil usaha; kelompok `gift_other`: hadiah, lainnya (**[Usulan]** lead, untuk checkpoint S4; §3.1). Nilai lama `salary`, `business`, `gift`, `other` di contoh dokumen ini adalah sumber bawaan gaji, hasil usaha, hadiah, dan lainnya; penghasilan lepas yang dahulu `business` kini sumber bawaan honor atau pekerjaan lepas (`freelance`). Model tidak mengandaikan irama penghasilan apa pun: harian, bulanan, beberapa kali sebulan, atau satu periode tanpa penghasilan sama sekali sama-sama sah, dan tidak ada aturan atau angka yang bergantung pada adanya, kekerapan, atau tanggal uang masuk (catatan pemilik produk untuk tahap publik, checkpoint S2). Baris `income` bersumber penghasilan selalu berpemilik `shared` (§3.1, INV-24) dan wajib membawa `earner`: tepat satu anggota dewasa rumah tangga itu, yaitu isian "punya siapa" di layar, diisi pencatat dan terpisah dari `actor` (INV-27) |
| `expense` | Pengeluaran. Pengembalian dana adalah baris `expense` bertanda + yang tertaut ke aslinya | − | kategori, penerima manfaat, pembayar, pihak luar (opsional), `funded_by` |
| `credit_cost` | Biaya kredit: selisih antara jumlah seluruh cicilan dan nilai yang dibiayai | − | cicilan terkait |
| `cash_difference` | Selisih kas | + atau − | rekonsiliasi terkait |
| `write_off` | Penghapusan piutang | − | piutang terkait, alasan |
| `valuation` | Perubahan nilai aset, termasuk untung-rugi penjualan | + atau − | aset terkait |
| `owner_transfer` | Perpindahan nilai antar-pemilik: penyisihan (keluar dari buku rumah tangga), termasuk bagian pemasukan yang diarahkan langsung ke dana pribadi (§3.1); kontribusi (masuk ke buku rumah tangga); pelepasan klaim antar-pemilik (§3.6); dan penutupan klaim saat anggota keluar (§8.2) | + atau − | pemilik lawan |
| `debt_forgiven` | Utang rumah tangga yang dibebaskan pihak luar | + | kewajiban terkait |

Pemindahan, serah terima, pinjaman (diterima atau diberikan), pelunasan, penggantian, pembayaran cicilan, dan pembelian aset secara tunai **tidak punya baris sebab**: kekayaan bersih tidak berubah.

### 2.5 Entri dan barisnya

```text
entry
  id, household_id
  type              -- daftar tertutup, tabel di bawah
  event_date        -- tanggal kejadian, diisi pencatat; tidak boleh di masa depan
  recorded_at       -- tanggal pencatatan, diisi server; tidak dapat diatur pencatat
  actor             member(M) | system    -- siapa yang mencatat; selalu ada. system hanya pada exit_closure (§8.2)
  idempotency_key   -- dibuat klien satu kali per niat mencatat; unik per rumah tangga
  corrects          null | entry_id       -- entri yang dikoreksi
  correction_kind   null | reversal | adjustment
  links             authorization_id?, installment_id?, recurring_bill_id?,
                    bill_occurrence_due?, -- tanggal jatuh tempo kejadian tagihan yang dibayar (§3.15)
                    set_aside_schedule_id?, reconciliation_id?, refund_of_entry_id?,
                    subject_ref?          -- rujukan induk (kind, id) menurut dokumen 06 §2.2
  detail_ref        -> entry_detail       -- catatan, struk, nama pihak lawan; disimpan terpisah (§8)

entry_link                                -- tautan susulan (§3.8); hanya ditambah, tidak mengubah entri
  id, household_id, entry_id
  kind              authorization | installment | recurring_bill | set_aside_schedule | subject_ref
  target            -- id sesuai kind
  occurrence_due?   -- hanya untuk kind recurring_bill: kejadian yang dibayar (§3.15)
  actor, recorded_at, idempotency_key
  revokes           null | entry_link_id  -- mencabut tautan susulan yang keliru

line                                      -- baris; tidak pernah berdiri tanpa entri
  entry_id, owner                         -- pemilik uang yang bukunya disentuh baris ini
  kind              position | cause | allocation
  -- position:
  position_ref      holding | receivable | liability | asset
  delta             bilangan bulat bertanda (untuk kewajiban: + berarti kewajiban bertambah)
  -- cause:
  cause             nilai §2.4
  effect            bilangan bulat bertanda (akibat pada kekayaan bersih)
  category_id?, beneficiary?, payer?, funded_by?, earner?, source?
  -- allocation:
  holding_id, purpose_id, delta
```

**Keseimbangan.** Untuk setiap entri dan setiap pemilik uang yang disentuhnya:

```text
Σ delta(tempat uang) + Σ delta(piutang) + Σ delta(aset) − Σ delta(kewajiban) = Σ effect(sebab)
```

Baris `allocation` tidak ikut persamaan ini; ia hanya memberi label tujuan (§3.9).

**Rujukan induk.** `links.subject_ref` adalah rujukan induk dokumen 06 §2.2: pasangan `(kind, id)` yang menunjuk permintaan, tugas, masalah rumah, kesepakatan, atau objek lain tempat entri itu menempel. "Aktivitas" bukan primitif (ADR-001); model keuangan tidak mengenal objek induk generik dan tidak menurunkan akibat uang apa pun dari rujukan ini. Akibat uang hanya lahir dari tautan bertipe: `authorization_id` (§3.8), `installment_id`, `recurring_bill_id`, `set_aside_schedule_id`, `reconciliation_id`, `refund_of_entry_id`.

**Penanda `unconfirmed`.** Tiga macam baris tempat uang membawanya (§3.3):

| | Baris | Sisi |
|---|---|---|
| (a) | `delta > 0` pada tempat uang yang pemegangnya anggota dewasa selain `actor` entri, apa pun jenis entrinya, termasuk `opening` dan selisih hitung positif | Penerima |
| (b) | `delta < 0` pada tempat uang berpemilik `shared` yang pemegangnya anggota dewasa selain `actor`, **bila entrinya entri dua buku**, yaitu juga memuat baris berpemilik `member(M)`: penyisihan, penggantian, pinjaman dana bersama kepada M, pelunasan utang dana bersama kepada M | Pembayar |
| (c) | `delta < 0` pada entri `cash_difference` atas tempat uang yang pemegangnya anggota dewasa selain `actor`: selisih hitung negatif yang dicatat bukan pemegangnya (**[Usulan]**, putaran S3; G-05; §3.10 butir 1, yang hanya menerima hitungan itu atas rekening atau dompet digital yang ditandai pemegangnya dan bila tidak ada yang masih terbuka, INV-33). Selisih positif yang dicatat bukan pemegang sudah tercakup (a) | Pemegang |

Penanda ini fakta turunan dari `actor`, `holder`, dan buku yang disentuh entri saat entri disimpan, bukan kolom yang kelak diubah; status konfirmasinya sesudah itu dipegang dokumen 06 §7 sebagai peristiwa atas baris itu.

Enam dimensi custody dijawab dari:

| Dimensi (brief §5C) | Dijawab oleh |
|---|---|
| Siapa pemiliknya | `holding.owner` |
| Siapa pemegangnya | `holding.holder` |
| Di mana disimpan | `holding.kind`, `holding.name` |
| Untuk tujuan apa | baris `allocation` per tempat uang dan tujuan |
| Siapa yang berwenang membelanjakan | aturan rumah tangga (dokumen 06) dan *otorisasi* yang terbuka; bukan atribut uang |
| Siapa penerima manfaatnya | `beneficiary` pada baris sebab `expense` |

`actor` (yang mencatat), `payer` (yang membayar), dan `holder` (yang memegang) adalah tiga hal terpisah. Model tidak mengandaikan satu pasangan sebagai pencatat (keputusan pemilik nomor 1 dan 4).

**Jenis entri.** "Posisi" dan "sebab" di tabel ini adalah baris yang wajib ada; buku yang disentuh disebut di kolom terakhir.

| `type` | Baris posisi | Baris sebab | Buku |
|---|---|---|---|
| `opening` | posisi apa pun ± | `opening` | satu |
| `income` | tempat uang + | `income` + | satu |
| `income` dengan bagian pribadi (§3.1) | buku rumah tangga: tempat uang + sebesar jumlah dikurangi bagian pribadi (tidak ada bila seluruhnya ke dana pribadi); buku pribadi pencatat: tempat uang + sebesar bagian pribadi | buku rumah tangga: `income` + sebesar jumlah penuh dan `owner_transfer` − sebesar bagian pribadi; buku pribadi: `owner_transfer` +. Untuk penghasilan: hanya bila `actor = earner` (INV-27; keputusan pemilik produk, checkpoint S2) | dua |
| `expense` | tempat uang − | `expense` − (satu atau lebih; boleh disertai `expense` untuk biaya admin) | satu |
| `refund` | tempat uang + | `expense` +, tertaut ke entri asal | satu |
| `transfer` (termasuk serah terima) | tempat uang −, tempat uang + (pemilik sama); biaya admin opsional | tidak ada; `expense` − bila ada biaya admin | satu |
| `owner_transfer` | tempat uang − di buku asal, tempat uang + di buku tujuan | `owner_transfer` − dan + | dua |
| `contribution_expense` | tempat uang − di buku pribadi | buku pribadi: `owner_transfer` −; buku rumah tangga: `owner_transfer` + dan `expense` − | dua |
| `advance` (talangan dengan "minta diganti") | buku pembayar: tempat uang −, piutang +; buku yang ditalangi: kewajiban + | buku yang ditalangi: `expense` − | dua |
| `reimbursement` | buku yang ditalangi: tempat uang −, kewajiban −; buku pembayar: tempat uang +, piutang − | tidak ada | dua |
| `claim_waiver` | piutang − dan kewajiban − | `owner_transfer` − di buku yang melepas, + di buku yang dibebaskan | dua |
| `exit_closure` (hanya oleh sistem, saat anggota keluar; §8.2) **[Usulan]** | untuk setiap klaim antar-pemilik yang menyangkut buku pribadi anggota yang keluar: piutang − dan kewajiban − sebesar seluruh sisanya; tidak pernah baris tempat uang atau aset | per klaim: `owner_transfer` − di buku pemilik piutang, + di buku yang berkewajiban | dua atau lebih |
| `loan_given` | tempat uang −, piutang + | tidak ada | satu (dua bila antar-pemilik) |
| `loan_received` | tempat uang +, kewajiban + | tidak ada | satu (dua bila antar-pemilik) |
| `repayment_in` | tempat uang +, piutang − | tidak ada | satu (dua bila antar-pemilik) |
| `repayment_out` (termasuk pembayaran cicilan) | tempat uang −, kewajiban − | tidak ada; `expense` − untuk denda | satu (dua bila antar-pemilik) |
| `write_off` | piutang − | `write_off` − | satu |
| `debt_forgiven` | kewajiban − | `debt_forgiven` + | satu |
| `asset_purchase` | tempat uang − (uang muka atau lunas), aset +, kewajiban + bila dicicil | `credit_cost` − bila dicicil dan ada biaya kredit | satu |
| `asset_valuation` | aset ± | `valuation` ± | satu |
| `asset_disposal` | aset −, tempat uang + (bila dijual) | `valuation` ± sebesar selisih hasil dan nilai tercatat | satu |
| `cash_difference` | tempat uang ± | `cash_difference` ± | satu |
| `recategorization` | tidak ada | dua baris sebab yang saling meniadakan (§3.10) | satu |
| `allocation` | tidak ada | tidak ada; hanya baris `allocation` | satu |

Koreksi bukan jenis tersendiri: entri pembalik membawa `type` entri asal, `corrects`, dan `correction_kind = reversal`.

**Penulisan.** Satu perintah tulis menghasilkan tepat satu entri dalam satu transaksi basis data, membawa satu `idempotency_key` (`ENGINEERING-STANDARDS.md` §5). Perintah dengan kunci yang sudah pernah diterima mengembalikan entri yang sama dan tidak menulis apa pun (INV-10). Saldo posisi adalah jumlah `delta` barisnya; cache saldo, bila ada, diperbarui di transaksi yang sama (INV-04).

**Saldo negatif.** Pencatatan tidak pernah ditolak karena saldo tercatat tidak cukup: pencatat pelupa sering mencatat belanja sebelum mencatat uang masuknya. Saldo tercatat tempat uang yang negatif adalah tanda "ada yang belum dicatat" dan memicu ajakan rekonsiliasi kas; ia bukan pelanggaran invarian.

### 2.6 Entitas pendukung

Tidak satu pun entitas di bawah mengubah saldo. Semuanya dibuat, diubah, dan diarsipkan rumah tangga sendiri.

```text
category          id, name, parent_id?, archived_at
                  -- tidak ada kategori yang terkunci dan tidak ada kategori yang diwajibkan model
                  -- (keputusan pemilik nomor 9; T-10); daftar awal saat onboarding hanya contoh
                  -- yang semuanya dapat dihapus, misalnya belanja bulanan, cicilan, pulsa/wifi

purpose           id, name, kind (spending | savings), target_amount?, target_date?,
                  recurring_bill_id?,   -- tujuan "siapkan uang" milik satu tagihan berulang (§3.15)
                  archive_reason?,      -- achieved | cancelled; hanya keterangan (§3.9)
                  archived_at
                  -- tujuan; tabungan adalah tujuan berjenis savings

budget_period_setting  household_id, start_day (1..28), effective_from, actor, recorded_at
                  -- hari mulai periode anggaran rumah tangga; bawaan 1; hanya ditambah (§6.1)

budget            id, category_id, limit_amount, period_start, period_end
                  -- satu baris per kategori per periode yang sudah dimulai; periode adalah rentang tanggal
budget_rule       category_id, limit_amount
                  -- pola batas bulanan per kategori; periodenya mengikuti budget_period_setting

authorization     id, request_id (dokumen 06), requester, limit_amount,
                  funding_owner (shared), category_id?, purpose_id?, source_holding_id?,
                  planned_date?, valid_until, terms_version,
                  status (open | settled | cancelled | expired), opened_at, closed_at
                  -- batas, kategori, tujuan, sumber, tanggal rencana, dan masa berlaku menyalin
                  -- versi syarat yang disetujui (dokumen 06); berubah hanya lewat authorization.amended (§3.8)

reservation       id, authorization_id, holding_id, amount, status (active | released)

installment       id, liability_id, asset_id?, name, counterparty,
                  amount_per_payment, payments_total, payments_made, due_day,
                  category_id, responsible_member
                  -- tanpa routine_agreement_ref: pembayaran cicilan tidak ditandai (dokumen 06 §6.2)

recurring_bill    id, name, counterparty, estimated_amount, due_rule?,
                  category_id, responsible_member, end_date?, routine_agreement_ref?, archived_at
                  -- due_rule Tier 1: monthly(day), day 1..31; bulan yang tidak punya hari itu
                  --   memakai hari terakhirnya (§3.15 "Kejadian tagihan")
                  -- tanpa due_rule (IDEA-03, S3-2 (ii)): tanpa kejadian dan tanpa pengingat (§3.15)

set_aside_schedule  id, recipient (member(M)), monthly_amount, due_rule?,
                  routine_agreement_ref?, archived_at
                  -- jatah pribadi M: satu jumlah per periode anggaran (§3.4); paling banyak satu
                  -- yang tidak terarsip per recipient; due_rule hanya untuk pengingat

income_source     id, household_id, name, group (work | gift_other),
                  default_key?, archived_at
                  -- [Usulan] lead dan dokumen ini, untuk checkpoint S4 (§3.1 "Sumber uang masuk")
                  -- default_key: salary | thr | bonus | overtime | freelance | business  (work)
                  --              gift | other                                          (gift_other)
                  -- name unik di antara sumber yang tidak terarsip; gift dan other tidak pindah
                  --   kelompok dan tidak diarsipkan; perubahan = peristiwa tanpa entri (INV-12)
                  -- sumber gift_other tidak memakai nama yang pernah dipakai sumber work (CP-05);
                  --   riwayat nama disimpan sebagai peristiwa

reconciliation    id, holding_id, counted_amount, recorded_amount_at_count, counted_at, actor
                  -- recorded_amount_at_count dihitung server saat perintah diterima, bukan
                  --   dikirim perangkat (§3.10 "Hitungan yang tiba belakangan")

external_party    id, nickname, relation      -- nama panggilan dan hubungan saja (Q-28)

entry_detail      entry_id, note?, receipt_refs?, counterparty_text?   -- dapat dihapus terpisah (§8)
```

**Rujukan ke kesepakatan rutin.** `routine_agreement_ref` hanya ada pada tagihan berulang dan jadwal penyisihan, dan menunjuk *kesepakatan rutin* yang berlaku atas jadwal itu, bila ada (dokumen 06 §6.3). Setiap tagihan berulang dibuat bersama usulan kesepakatannya (§3.15); rujukan itu kosong selama usulannya belum disetujui dan sesudah kesepakatannya diakhiri, dan selama itu pembayarannya mengikuti aturan ambang biasa. **Cicilan tidak memilikinya**: pembayaran cicilan adalah pelunasan kewajiban yang sudah tercatat dan tidak pernah dibandingkan dengan ambang permintaan (§3.8, §3.14; dokumen 06 §6.2), sehingga rujukan itu tidak punya pemakai. Jumlah yang disepakati disimpan pada kesepakatan itu sendiri (dokumen 06 §6.3 "Kesepakatan rumah tangga"), bukan pada jadwal: per kali untuk tagihan berulang, per periode anggaran untuk *jatah pribadi*. `recurring_bill.estimated_amount` adalah perkiraan untuk pengingat, dan `set_aside_schedule.monthly_amount` adalah jumlah yang direncanakan atau diusulkan penerimanya; keduanya metadata yang dapat diubah tanpa persetujuan, dan mengubahnya tidak mengubah jumlah yang disepakati (menjawab dokumen 06 §15.4 butir (d)). Jumlah jatah yang berlaku selalu jumlah pada kesepakatan itu, bukan `monthly_amount` (§3.4). Rujukan itu diisi, diganti, dan dikosongkan hanya oleh alur kesepakatan di dokumen 06; ia metadata jadwal, bukan inti moneter, dan tidak mengubah angka apa pun (INV-12). Entri pembayaran tidak menunjuk kesepakatan secara langsung: ia ditautkan ke jadwalnya, dan dari jadwal itulah dokumen 06 membaca apakah pembayaran itu tercakup kesepakatan rutin. Kiriman keluarga bulanan memakai `recurring_bill` dengan pihak luar penerimanya sebagai `counterparty`.

**Pengarsipan.** Menghapus sesuatu yang punya riwayat berarti mengarsipkannya (keputusan pemilik nomor 9; Q-58): ia hilang dari pilihan saat mencatat, riwayat dan laporannya tetap. Syarat:

| Yang diarsipkan | Syarat |
|---|---|
| Tempat uang | Saldo tercatat nol, tidak ada reservasi aktif, dan jumlah alokasinya nol. Bila belum nol, pemegang memindahkan isinya atau merekonsiliasinya lebih dulu; alokasi yang tersisa ditawarkan untuk dilepas di langkah yang sama, tanpa pilihan terisi (**[Usulan]**, putaran S3) |
| Piutang, kewajiban | Saldo nol (lunas, dihapus, atau dibebaskan) |
| Aset | Nilai tercatat nol lewat `asset_disposal` |
| Kategori | Selalu boleh. Anggaran kategori itu pada periode yang sedang berjalan tetap berlaku sampai akhir periode itu; periode berikutnya tidak lagi membuat baris `budget` untuknya |
| Tujuan (termasuk tujuan tabungan) | Jumlah alokasinya di semua tempat uang nol (**[Usulan]**, putaran S3; G-01; §3.9 "Menyelesaikan atau membatalkan tujuan"). Alokasi tidak berperiode, jadi tidak ada yang "berakhir di akhir periodenya": rumusan revisi sebelumnya untuk tujuan **keliru** dan dicabut |
| Cicilan | Kewajibannya nol |
| Tagihan berulang | Selalu boleh, oleh tindakan anggota; tagihan **tidak pernah terarsip sendiri**, juga sesudah tanggal akhirnya (§3.15). Tujuan "siapkan uang" milik tagihan itu diarsipkan bersamanya bila alokasinya nol; bila belum, alokasinya ditawarkan untuk dilepas di langkah yang sama. Dari sisi uang, kesepakatan rutin yang merujuknya tidak perlu diakhiri: tagihan terarsip tidak menerima tautan baru, jadi kesepakatan itu tidak lagi menilai apa pun. Apakah dan oleh siapa kesepakatan itu diakhiri saat pengarsipan adalah aturan dokumen 06 §6.3 |
| Jadwal penyisihan | Selalu boleh dari sisi uang (tidak ada saldo atau alokasi yang harus nol). **Pelakunya** (**[Usulan]** `architect`, dokumen 06 §6.3 "Pos rutin" (g); diikuti dokumen ini, putaran perbaikan S3): bila jadwal itu dirujuk kesepakatan jatah yang **berlaku**, hanya pemilik jatah (penerima jadwal itu) yang dapat mengarsipkannya, dan pengarsipannya dijalankan dalam satu transaksi dengan pengakhiran jatahnya oleh dia sendiri (K06). Bila tidak dirujuk kesepakatan yang berlaku, anggota dewasa mana pun. Alasannya dari sisi uang: mengarsipkan jadwal yang dirujuk jatah sama dengan mengakhiri jatah, dan jatah hanya diakhiri pemiliknya (§3.4, D-52). **Akibatnya pada penilaian sama persis dengan mengakhiri jatah sendiri** (§3.4): bila jatah itu di atas ambang yang berlaku saat diarsipkan, ia tetap dasar penilaian sampai `assessment_until`, walau jadwalnya sudah terarsip; bila tidak di atas ambang, berakhir seketika. `ada_jatah` dibaca dari kesepakatan, bukan dari status arsip jadwal, jadi tidak ada rumus yang berubah. Mengaktifkan kembali jadwal itu tidak menghidupkan jatah lama; menyepakatinya lagi adalah usulan baru |

**Posisi terarsip tidak menerima baris posisi baru** (INV-21). Entri `recategorization` tidak punya baris posisi, jadi perincian selisih hitung atas tempat uang terarsip tetap boleh (§3.10 butir 6). Tagihan berulang dan tujuan yang terarsip tidak menerima tautan baru atau alokasi baru.

**Mengaktifkan kembali** (**[Usulan]**, putaran S3; PA-FIN-4; T-30). Tempat uang, tagihan berulang, kategori, atau tujuan yang terarsip dapat diaktifkan kembali oleh anggota yang boleh mengarsipkannya (siapa itu: dokumen 06, PA-FIN-1; tempat uang pribadi hanya oleh pemiliknya). Pengaktifan kembali adalah perubahan metadata: tercatat sebagai *peristiwa* beserta pelakunya, terlihat sama seperti pengarsipannya, tidak membuat entri, dan tidak mengubah angka apa pun (INV-12). Sesudahnya butir itu menerima entri, tautan, dan alokasi seperti biasa dan boleh diarsipkan lagi bila syaratnya terpenuhi. Inilah jalan entri terlambat atas tempat uang terarsip yang tidak dapat dicatat sebagai perincian selisih: PRN-05 ("terlambat itu sah") terpenuhi tanpa melonggarkan INV-21. Bila entri itu membuat saldo bukan nol, berarti ada catatan lain yang keliru (misalnya pemindahan yang tercatat terlalu besar); perbaikannya mengikuti §3.12, lalu tempat uang itu diarsipkan lagi. Contoh: CTH-43.

Yang belum pernah dipakai entri mana pun boleh dihapus sungguhan. Ganti nama adalah perubahan metadata: tercatat sebagai *peristiwa* di jejak audit beserta pelakunya, tidak membuat entri, dan riwayat menampilkan nama terbaru.

---

## 3. Aturan per jenis peristiwa

### 3.1 Uang masuk

Keputusan pemilik nomor 1 dan 2. Layar "uang masuk" mencatat sumber, atas nama siapa, dan siapa pemegangnya. Yang menerima uang otomatis menjadi pemegangnya dan mencatatnya; pemegang ditentukan oleh tempat uang yang dipilih.

**Keputusan pemilik produk, 11 November 2026.** Yang mencatat uang masuk memilih ke mana uang itu pergi: masuk kas bersama, langsung dipakai untuk sesuatu, atau masuk dana pribadi; ini berlaku juga untuk gaji. Bila gaji diarahkan ke dana pribadi, pasangan melihat catatannya, misalnya "gaji Rp10 juta masuk, Rp1 juta ke pribadi": jumlah gajinya dan berapa yang ke pribadi. Yang tetap tertutup hanya saldo dana pribadi itu dan pemakaiannya sesudahnya. Dengan keputusan ini, usulan dokumen ini sebelumnya ("gaji selalu masuk dana bersama") **tidak** diambil sebagai larangan; yang diambil adalah keterlihatannya.

| Sumber di layar | Entri | Pemilik bawaan | Boleh dicatat hanya di buku pribadi? |
|---|---|---|---|
| Sumber berkelompok `work`, "dari pekerjaan atau usaha" (*penghasilan*): bawaannya gaji, THR, bonus, lembur, honor atau pekerjaan lepas, hasil usaha, ditambah sumber yang disepakati rumah tangga | `income` | `shared`, walau tercatat "gaji suami" atau "THR istri" lewat `earner` | **Tidak.** Pemasukannya selalu tercatat penuh di buku rumah tangga; bagian yang ke dana pribadi tercatat di entri yang sama (INV-24) |
| Hadiah untuk satu orang (kelompok `gift_other`) | `income` di buku pribadi orang itu | `member(M)` | Ya, dan itu bawaannya |
| Hadiah untuk keluarga (kelompok `gift_other`) | `income` | `shared` | — |
| Pinjaman yang diterima | `loan_received`; **bukan pemasukan** | `shared` | — |
| Lainnya, dan sumber buatan sendiri berkelompok `gift_other` | `income` | `shared` | Ya, bila memang uang pribadi (misalnya hasil menjual barang pribadi) |

**Sumber uang masuk.** Kata-kata pemilik produk, 12 November 2026, kata demi kata: ia bertanya "kenapa thr dicatat sebagai gaji? kan pemasukan kita yang atur darimana nya?", lalu, atas pertanyaan lead tentang siapa yang mengatur sumber, menulis "ada kategori default bisa edit tambah atau yang lainnya" ([_progress.md](_progress.md), "Pertanyaan pemilik produk atas S3-3"). Ia menulis "kategori"; bahwa yang ia maksud adalah **sumber** uang masuk (asal pemasukan, bukan kategori keperluan) adalah **bacaan lead [Usulan]** dari konteks pertanyaannya ("darimana nya"). Menurut bacaan itu: sumber uang masuk adalah **daftar bawaan yang dapat disunting dan ditambah rumah tangga, dan punya "lainnya"**. Yang **tidak** ia jawab: apakah sumber buatan sendiri boleh dicatat tanpa terlihat pasangan. Bentuk di bawah adalah bawaan lead **[Usulan]** yang ditanyakan di checkpoint S4, dengan rincian aturan uang **[Usulan]** dokumen ini (T-34, T-43).

| Hal | Aturan |
|---|---|
| Sumber | Satu baris `income_source` per sumber rumah tangga (§2.6): nama yang tampil, satu **kelompok**, dan penanda bawaan. Baris sebab `income` menunjuknya lewat `source_id`. Kategori pemasukan tetap terpisah dari sumber |
| Dua kelompok | **`work` ("dari pekerjaan atau usaha")**: pemasukannya *penghasilan*; berlaku semua aturan yang kini berlaku atas gaji: selalu tercatat penuh di buku rumah tangga (INV-24), wajib `earner`, bagian pribadi hanya oleh penerimanya (INV-27) dan terlihat sebagai penyisihan (§3.4). **`gift_other` ("hadiah dan lainnya")**: boleh dicatat di buku pribadi penerimanya saja dan tidak terlihat pasangan (hadiah untuk satu orang: keputusan pemilik produk, checkpoint S2, S2-1 butir (ii); "lainnya": D-47) |
| Daftar bawaan | `work`: gaji, THR, bonus, lembur, honor atau pekerjaan lepas, hasil usaha. `gift_other`: hadiah, lainnya. Nama bawaan dapat diganti rumah tangga |
| Menambah sumber `gift_other` | Anggota dewasa mana pun, sendirian, seketika, tunduk pada baris "Nama yang dicadangkan". Nama dan kelompoknya tampil bagi kedua anggota dewasa sebagai *peristiwa* metadata, tanpa entri (INV-12); pemasukan yang dicatat dengannya di buku pribadi tetap tidak terlihat. Alasan cukup sendirian: "lainnya" sudah ada, jadi sumber baru di kelompok ini tidak membuka jalan tersembunyi yang baru; yang bertambah hanya nama, dan nama itu terlihat |
| Menambah sumber `work`, mengganti namanya, mengarsipkannya, dan memindahkan sumber antarkelompok | **Kesepakatan rumah tangga** kedua anggota dewasa (dokumen 06 §6.3 `rule`, K01 sampai K05), berlaku ke depan saja. Dasarnya D-89 yang diterima lewat "lanjut" di checkpoint S3: daftar sumber penghasilan termasuk isi kesepakatan aturan rumah tangga dan hanya diubah bila keduanya setuju. Memindahkan sumber dari `work` ke `gift_other` melonggarkan keterlihatan atas penghasilan pasangan; dari `gift_other` ke `work` mewajibkan keterlihatan atas pemasukan pasangan yang selama ini boleh tertutup; keduanya menyentuh orang lain, jadi tidak pernah tindakan satu orang. Mengarsipkan atau mengganti nama sumber `work` juga lewat kesepakatan, karena arsip atau ganti nama ("THR" menjadi "THR kantor"), lalu sumber `gift_other` bernama "THR", akan menjadi pemindahan diam-diam (temuan `critic` CP-05). Paling sedikit satu sumber `work` tetap tidak terarsip |
| Mengganti nama sumber `gift_other` | Anggota dewasa mana pun, sendirian, sebagai *peristiwa* yang terlihat; kelompok tidak ikut berubah; tunduk pada baris "Nama yang dicadangkan". Entri lama tetap menunjuk `source_id` yang sama dan tampil dengan nama yang berlaku, dengan riwayat nama di jejak audit |
| Nama yang dicadangkan | Sumber `gift_other` tidak dapat diberi nama yang **pernah** dipakai sumber `work` mana pun di rumah tangga itu (nama sekarang, nama lama, dan sumber yang terarsip), maupun nama bawaan kelompok `work`; dibandingkan tanpa membedakan huruf besar dan kecil dan tanpa spasi di awal dan akhir. Perintah yang melanggar ditolak |
| Kelompok selalu tampil | Setiap catatan uang masuk dan setiap sumber di daftar menampilkan kelompoknya ("dari pekerjaan atau usaha" atau "hadiah dan lainnya") di samping namanya, bagi kedua anggota dewasa; kata-katanya milik dokumen 08. **Batas yang dinyatakan:** nama yang mirip tetapi tidak sama ("Tunjangan hari raya" untuk sumber `gift_other`) tidak tertangkap aturan nama; yang menjaganya adalah peristiwa penambahan yang terlihat dan kelompok yang selalu tampil |
| Mengarsipkan sumber `gift_other` buatan sendiri | Anggota dewasa mana pun, sendirian, sebagai *peristiwa* yang terlihat. "Hadiah" dan "lainnya" tidak dapat diarsipkan dan tidak dapat dipindah kelompok |
| Nama unik | Nama sumber yang tidak terarsip unik dalam rumah tangga (tanpa membedakan huruf besar dan kecil), supaya tidak ada "THR" di dua kelompok sekaligus |
| Entri lama | Kelompok dinilai saat entri diterima (INV-24); memindahkan atau mengarsipkan sumber tidak mengubah entri yang sudah tersimpan (INV-09) |
| Ringkasan | Tabel A baris "uang masuk, per sumber dan atas nama siapa" memakai nama sumber (§6.3); hanya baris buku rumah tangga yang masuk |
| Satu anggota dewasa | Selama rumah tangga belum pernah punya dua anggota dewasa, perubahan yang memerlukan kesepakatan mengikuti aturan dokumen 06 untuk kesepakatan rumah tangga satu dewasa (seperti zona waktu, D-96 (a)); itu aturan alur, bukan aturan uang |

**Batas yang dinyatakan.** Kelompok hanya menjaga pemasukan yang **dicatat** dengan sumber yang benar: pencatat tetap dapat memilih "lainnya" atas THR dan mencatatnya di buku pribadi saja (celah label, §10.3 butir 1). Model tidak dapat menutupnya tanpa membuka seluruh pemasukan dana pribadi.


Pemilik ditentukan oleh tempat uang tujuan (satu tempat uang, satu pemilik). Mengarahkan uang masuk berarti memilih tempat uang tujuannya.

**Tiga arah, satu aturan.**

| Arah yang dipilih pencatat | Bentuk di model |
|---|---|
| Masuk dana bersama | Entri `income`: tempat uang dana bersama + jumlah; sebab `income` + jumlah |
| Ke dana pribadi pencatat, sebagian atau seluruhnya (X dari jumlah J, 0 < X ≤ J). Untuk penghasilan: hanya bila pencatat adalah `earner`-nya (INV-27; keputusan pemilik produk, checkpoint S2) | **Satu** entri `income` dengan baris di dua buku. Buku rumah tangga: sebab `income` +J, sebab `owner_transfer` −X, tempat uang +(J − X). Buku pribadi pencatat: tempat uang +X, sebab `owner_transfer` +X. Keseimbangan buku rumah tangga: J − X = J − X; buku pribadi: X = X; jumlah `owner_transfer` nol (INV-03) |
| Langsung dipakai untuk sesuatu | Entri `income` ke tempat uang penerimanya, lalu entri biasa dari tempat uang itu (pengeluaran, pelunasan, pinjaman, dan seterusnya). Layar boleh menyajikannya sebagai satu alur; modelnya tetap dua entri dengan dua idempotency key **[Usulan]**, karena "sesuatu" itu dapat berupa jenis entri apa pun dan jenis entri adalah daftar tertutup |

Mengapa bagian pribadi adalah baris di entri yang sama, bukan pemasukan ke dana bersama disusul penyisihan sebagai entri kedua **[Usulan]**:

1. Satu perintah tulis menghasilkan tepat satu entri (§2.5). Pemasukan dan bagian pribadinya tersimpan bersama atau tidak sama sekali; tidak ada saat ketika gaji sudah tercatat tetapi arahnya belum.
2. Catatan mengikuti kenyataan. Uang yang tidak pernah berada di tempat uang dana bersama tidak dicatat seolah pernah di sana; tidak ada serah terima semu dan tidak ada saldo yang sesaat salah.
3. Yang dilihat pasangan adalah satu catatan, persis seperti contoh pemilik produk: jumlah penuh dan bagian pribadinya.

Bagian pribadi itu adalah *penyisihan* (§3.4) dan tunduk pada aturannya: terlihat oleh keduanya beserta jumlah, tanggal, pelaku, dan penerimanya; tampil di ringkasan sebagai "disisihkan ke dana pribadi" (§6.3). Ia terhitung terhadap *jatah pribadi* bulanan pencatatnya (§3.4; default yang diterima di checkpoint S2 lewat "lanjut", dokumen 17 D-48).

**Beberapa uang masuk dalam satu periode; rekening bank atau dompet digital** (pernyataan pemilik produk, 11 November 2026: "gajian saya bisa 2x bahkan uang masuk dari freelancer juga bisa sampe 5x", masuk ke rekening bank atau e-wallet). "Sebulan" dan jumlah seluruhnya, sampai sekitar tujuh uang masuk, adalah **bacaan lead**, bukan kata-katanya; kalimatnya juga dapat dibaca "seluruhnya sampai lima kali" (dokumen 17, kelompok ketiga). Tidak ada aturan yang bergantung pada angkanya. Tidak ada aturan baru. Setiap uang masuk adalah satu entri `income` dengan idempotency key-nya sendiri, ke tempat uang yang menerimanya, apa pun jenisnya (`bank`, `ewallet`, `cash`); model tidak mengandaikan satu gajian per periode, dan tidak ada angka yang bergantung pada "tanggal gajian". Penghasilan lepas dicatat dengan sumber berkelompok `work`: dahulu `business`, kini sumber bawaan "honor atau pekerjaan lepas" (diterima pemilik produk di checkpoint S2 atas saran lead, S2-1 butir (iv); bukan pernyataannya sendiri; T-27 ditutup; nama sumber bawaannya **[Usulan]** lead untuk checkpoint S4), sehingga ia *penghasilan*, terlihat seperti gaji, dan tunduk pada INV-24 dan INV-27. Satu periode tanpa uang masuk sama sekali juga sah: laporan periode itu memuat pemasukan nol, dan anggaran, jatah, serta periode berjalan seperti biasa (§2.4). Tiap entri boleh memuat bagian pribadi; jumlahnya sepanjang periode dihitung terhadap satu jatah (§3.4). Meneruskan uang itu sesudahnya memakai aturan yang sudah ada: ke pasangan adalah serah terima, juga dari rekening ke rekening (§3.3); tarik tunai adalah pemindahan, dengan biaya admin di entri yang sama (§3.3); kepada ibu, adik, atau ayah adalah kiriman keluarga atau pinjaman, dipilih per entri (§3.2). Contoh terhitung satu periode dengan tiga uang masuk: CTH-39.

**Sumber mana yang wajib terlihat** (default yang diterima di checkpoint S2 lewat "lanjut", dokumen 17 D-47; D-89 di checkpoint S3). Daftarnya bukan bagian tetap dari model: ia isi rumah tangga, *daftar sumber penghasilan* (INV-24; permintaan `security-privacy`, SP-18), supaya model tidak mengunci satu bentuk keluarga. Sejak kata-kata pemilik produk 12 November 2026 ("ada kategori default bisa edit tambah atau yang lainnya", dibaca lead sebagai sumber uang masuk, **[Usulan]**) daftar itu adalah **sumber berkelompok `work`** di "Sumber uang masuk" di atas. Bawaannya mencakup gaji dan hasil usaha, yang sejak checkpoint S2 adalah preset pilot (keduanya nafkah rumah tangga, dan contoh pemilik produk menyebut gaji), ditambah THR, bonus, lembur, serta honor atau pekerjaan lepas (**[Usulan]** lead, untuk checkpoint S4). Hadiah untuk satu orang dan "lainnya" tidak termasuk, lihat di bawah. Bila pemilik produk menghendaki lain (misalnya "lainnya" ikut wajib), yang berubah hanya kelompoknya (T-14). Perubahan daftar berlaku atas perintah yang diterima sesudahnya; entri lama tidak berubah (INV-09). Siapa yang boleh mengubahnya: tabel "Sumber uang masuk"; alurnya dokumen 06 §6.3 "Kesepakatan rumah tangga" (§10.2).

**THR, bonus, dan lembur yang diterima** (T-34, **dibuka ulang** 12 November 2026). Riwayatnya: usulan putaran S3 menjawab dokumen 11 §10.1 butir 5 dengan mencatat THR dan bonus dari pemberi kerja bersumber `salary`, dan itu diterima lewat "lanjut" di checkpoint S3 (S3-3 pilihan A). Sesudahnya pemilik produk bertanya, kata demi kata, "kenapa thr dicatat sebagai gaji? kan pemasukan kita yang atur darimana nya?", lalu menulis "ada kategori default bisa edit tambah atau yang lainnya" ([_progress.md](_progress.md), "Pertanyaan pemilik produk atas S3-3"). S3-3 A karenanya **dibuka ulang sebagian**: THR tidak lagi dipaksa berlabel gaji. Yang berlaku sekarang (**[Usulan]** lead, untuk checkpoint S4): THR, bonus, dan lembur adalah **sumber bawaan tersendiri berkelompok `work`** ("Sumber uang masuk" di atas), sehingga label di layar adalah "THR", bukan "gaji", sedangkan akibat uangnya sama dengan yang diterima di S3-3 A: THR adalah *penghasilan*, selalu tercatat penuh di buku rumah tangga dan terlihat pasangan (INV-24), wajib `earner`, bagian pribadinya hanya dapat diarahkan penerimanya (INV-27) serta terhitung terhadap jatah pribadinya (§3.4). Bonus atau bagi hasil dari usaha sendiri memakai sumber hasil usaha. Rumah tangga boleh memindahkan THR ke `gift_other` hanya lewat kesepakatan berdua (D-89). Alasan THR di kelompok `work` secara bawaan: uang hasil bekerja adalah nafkah rumah tangga, dan membiarkan THR menjadi "lainnya" membuka celah label tepat pada uang masuk terbesar dalam setahun bagi banyak keluarga **[Asumsi]**. Celah label §10.3 butir 1 tetap: THR yang dicatat pencatatnya sebagai "lainnya" di buku pribadi tidak terlihat. Pertanyaan sampingan S3-3 "apakah kalian menerima THR?" tidak dijawab di checkpoint S3; sesudahnya pemilik produk menulis, kata demi kata, "thr bisa menerima bisa ngasih" ([_progress.md](_progress.md), "Jawaban pemilik produk atas S3-10 dan pertanyaan THR", 12 November 2026). Bacaan lead, belum dikukuhkan: rumah tangga dapat menerima THR dan juga memberi THR kepada orang lain. Paragraf ini mengatur THR yang **diterima**; THR yang **diberikan** diatur di §3.2 "THR dan angpau yang diberikan". Angka contoh tidak berubah: CTH yang memakai `salary` membaca sumber bawaan "gaji".

**Hadiah untuk satu orang masuk dana pribadi orang itu dan tidak terlihat oleh pasangan** (keputusan pemilik produk, checkpoint S2; dokumen 17 S2-1 butir (ii); Q-77). Ia `income` di buku pribadi penerimanya dan tidak menyentuh buku rumah tangga (CTH-36). Pilihan yang tidak diambil, dicatat satu kali: hadiah terlihat, yaitu `gift` masuk daftar sumber INV-24 dan dicatat dengan bentuk CTH-03b. Alasan: keputusan pemilik nomor 2 menjadikannya dana pribadi secara bawaan, keputusan nomor 3 menutup dana pribadi, dan pemilik produk tidak meminta hadiah terlihat. Menampakkannya berarti membuka pemasukan dana pribadi beserta asalnya (misalnya pemberian dari orang tua sendiri), yaitu hal yang justru dijaga keputusan nomor 3 bagi kedua pasangan secara sama. Biayanya dinyatakan di §10.3 butir 1: sumber diisi pencatat sendiri.

**Siapa yang mencatat.** Karena hanya M yang menulis ke buku pribadi M (§2.2), arah "ke dana pribadi" hanya tersedia bagi pencatat untuk dana pribadinya sendiri.

**Bagian pribadi dari penghasilan hanya bagi penerimanya** (keputusan pemilik produk, checkpoint S2; dokumen 17 S2-1 butir (i); INV-27; temuan `security-privacy` SP-01; dokumen 06 §6.4 "Uang masuk yang diarahkan ke dana pribadi"). Aturan di atas saja belum cukup: ia hanya memeriksa bahwa buku pribadi yang ditulis milik pencatat, sehingga A dapat mencatat "gaji B 10000000, 3000000 ke pribadi A", dan B tidak dapat membatalkannya, karena entri itu entri dua buku (§3.12). Aturannya: entri `income` bersumber penghasilan hanya boleh memuat bagian pribadi bila **pencatatnya adalah penerima penghasilan itu** (`actor = earner`). Perintah yang melanggar ditolak seluruhnya; pemasukan yang sama tetap dapat dicatat tanpa bagian pribadi, jadi tidak ada penghasilan yang terhalang dicatat. Bila pasangan yang mencatat gaji itu, ia mencatatnya ke dana bersama; siapa pun yang menghendaki bagian pribadi mengambilnya lewat penyisihan biasa (§3.4), yang terlihat dan mengikuti aturan jatah dan ambang.

- *Bagaimana penerima penghasilan dikenali.* Model sudah memisahkan `earner` dari `actor`: `earner` adalah isian "punya siapa" pada layar uang masuk (Q-55), tersimpan pada baris sebab `income`. Untuk sumber penghasilan ia wajib diisi dan bernilai tepat satu anggota dewasa rumah tangga itu. Hasil usaha yang dikerjakan berdua dicatat atas nama salah satu; bagian pasangannya diambil lewat penyisihan.
- *Batas yang dinyatakan.* `earner` diisi pencatat, sama seperti sumber. Pencatat dapat menulis dirinya sebagai `earner` atas gaji pasangannya. Bedanya dari celah label sumber (§10.3 butir 1): isian itu tercatat di buku rumah tangga dan terlihat oleh keduanya ("gaji, atas nama A"), sehingga pasangannya dapat mengajukan keberatan atas hal yang ia lihat. Pagar ini mencegah pengalihan yang jujur labelnya dan membuat yang tidak jujur terlihat; ia tidak membuktikan siapa penerima gaji sebenarnya.
- *Syarat keadilan.* Aturannya per penerima penghasilan dan simetris: kedua pasangan di rumah tangga pilot berpenghasilan (pernyataan pemilik produk, 11 November 2026), masing-masing mencatat penghasilannya sendiri dan hanya dapat mengarahkan bagian dari penghasilannya sendiri. Atas penghasilan pasangannya, seorang anggota memperoleh dana pribadi hanya lewat penyisihan. Pagar ini adil karena bagian pribadi penerima penghasilan diperlakukan dengan aturan yang sama dengan penyisihan lain: terhitung terhadap jatah pribadi yang sama dan ditandai dengan cara yang sama begitu jatah terlewati (T-13 dijawab "ya": default yang diterima di checkpoint S2 lewat "lanjut", dokumen 17 D-18 dan D-48; perlakuan yang melewati jatah: keputusan pemilik produk, checkpoint S2, S2-1 butir (iii) pilihan b; §3.4).
- *Pilihan yang tidak diambil, dicatat satu kali.* Tanpa INV-27: siapa pun yang mencatat penghasilan boleh mengarahkan sebagian atau seluruhnya ke dana pribadinya sendiri, juga atas gaji pasangannya; penerima gaji melihatnya dan dapat mengajukan keberatan, tetapi tidak dapat membatalkannya. Bentuk entri, saldo, dan semua angka contoh sama pada kedua pilihan: pada CTH-03b pencatat dan `earner` sama-sama suami.

### 3.2 Pengeluaran

- Satu entri `expense` boleh berisi beberapa baris sebab (kategori atau penerima manfaat berbeda) dan beberapa baris tempat uang (dibayar dari dua tempat). Jumlah keduanya sama.
- **Satu belanja dari tempat uang dua pemilik adalah dua entri** (**[Usulan]**, putaran S3; PA-FIN-6; T-31). Semua baris tempat uang satu entri `expense` berpemilik sama (`expense` adalah entri satu buku, §2.5). Belanja yang dibayar sebagian dari dana bersama dan sebagian dari dana pribadi pencatat dicatat sebagai `expense` atas bagian dana bersama dan `contribution_expense` atas bagian dana pribadi (atau `advance` bila bagian itu "minta diganti", §3.5), masing-masing dengan idempotency key-nya sendiri. Layar boleh menyajikannya sebagai satu alur, seperti T-15. Alasannya: jenis entri tertutup dan satu perintah satu entri (§2.5); bentuk satu entri akan menuntut jenis baru yang memuat kontribusi dan pengeluaran dana bersama sekaligus, tanpa manfaat bagi angka mana pun. Akibatnya dinyatakan: (1) yang dibandingkan dengan ambang hanya bagian dana bersama (§3.8), karena hanya itu pemakaian dana bersama; bagian dana pribadi tidak pernah dinilai; (2) anggaran kategori menghitung keduanya penuh (§6.1), jadi pemakaian kategori sama dengan bila dibayar satu sumber; (3) bila hanya satu dari dua perintah sampai di server, yang sampai tetap tersimpan dan tampil sendiri; tidak ada yang dibatalkan otomatis. Contoh: CTH-44.
- Belanja biasa dicatat setelah kejadian tanpa persetujuan (keputusan pemilik nomor 5). Status permintaan tidak pernah menghalangi pencatatan (A-20).
- **Kiriman keluarga** adalah `expense` dengan pihak luar sebagai penerima manfaat; tidak dilacak setelah diserahkan (default Q-07 yang diterima pemilik produk). Kiriman yang memang rutin dijadwalkan sebagai tagihan berulang (§3.15); di rumah tangga pilot kiriman untuk orang tua tidak rutin dan mengikuti aturan biasa (pernyataan pemilik produk, checkpoint S2).
- Uang untuk saudara atau orang tua: pencatat memilih per entri antara `loan_given` dan kiriman keluarga.

**THR dan angpau yang diberikan** (**[Usulan]** `finance-domain`, 12 November 2026; T-42). Dasarnya dua pernyataan pemilik produk sendiri ([_progress.md](_progress.md), 12 November 2026): "thr bisa menerima bisa ngasih", dan atas pertanyaan lead kepada siapa THR biasanya diberikan (keluarga besar, pekerja rumah, anak-anak saat Lebaran), "bro pikirkan kalau semua orang pakai app ini". Bacaan lead **[Usulan]**, yang diikuti di sini: aturannya harus berlaku bagi penerima siapa pun di tahap publik, bukan satu pola rumah tangga pilot. Hasilnya **tidak ada jenis entri, sebab, sumber, atau invarian baru**: THR yang diberikan adalah pengeluaran biasa, dan semua baris di bawah mengikuti aturan yang sudah ada. "THR" di sini mencakup THR, angpau, salam tempel, uang Lebaran, dan pemberian hari raya sejenis, apa pun agamanya.

| Hal | Aturan |
|---|---|
| Jenis entri | Uang yang diberikan tanpa harapan kembali dari dana bersama: `expense` (§3.2), satu buku, tanggal kejadian hari diberikan. Bila diharapkan kembali, itu pinjaman (`loan_given`, §3.7), dipilih pencatat per entri seperti kiriman keluarga. Pemberian berupa barang (parsel, hampers) adalah belanja biasa dengan kategori dan penerima manfaat yang sama |
| Penerima | `beneficiary` pada baris sebab `expense`. Pihak luar mana pun, tersimpan sebagai nama panggilan dan hubungan (Q-28; §8.2 butir 7): keluarga besar, pekerja rumah tangga (asisten rumah tangga, sopir, pengasuh), anak kerabat atau tetangga, dan pihak lain (petugas keamanan, petugas kebersihan, guru mengaji). Beberapa penerima kecil boleh dicatat sebagai satu pihak luar berkelompok ("anak-anak tetangga"); model tidak menuntut identitas tiap penerima. Satu entri boleh memuat beberapa baris sebab dengan penerima berbeda (§3.2 butir 1). Hubungan penerima tidak menimbulkan akibat uang apa pun (glosarium bagian 7.2: tidak satu pun dimensi custody diturunkan dari hubungan); daftar hubungannya milik glosarium bagian 3 dan dokumen 06 |
| Penerima anggota rumah tangga | Bukan uang yang keluar dari garis kendali. Bila uang itu menjadi dana pribadi anggota itu (misalnya anak yang menjadi anggota menerima uang Lebaran untuk dirinya), ia *penyisihan* (§3.4): terlihat, terhitung terhadap jatah pribadinya bila ada, selain itu dibandingkan dengan ambang. Bila uang itu dibelanjakan untuknya, ia `expense` dengan penerima manfaat anggota itu |
| Kiriman keluarga atau bukan | THR kepada keluarga besar adalah *kiriman keluarga* (glosarium bagian 7). THR kepada penerima lain bukan kiriman keluarga, tetapi aturan uangnya sama: `expense` berpenerima manfaat pihak luar. Model tidak membedakan keduanya selain lewat kategori dan hubungan penerima, dan tidak ada aturan ambang yang bergantung pada hubungan |
| Ambang permintaan | Bagian dana bersama dibandingkan dengan ambang **per entri**, dasar `threshold`, seperti kiriman keluarga (§3.8). **Batas yang dinyatakan:** banyak amplop yang masing-masing tidak di atas ambang tidak ditandai, berapa pun jumlah semusimnya. Penilaian kumulatif seperti `write_off` (INV-34) tidak dipakai: penghapusan menunjuk satu piutang yang jumlahnya diketahui, sedangkan THR menunjuk banyak penerima dan tidak ada jumlah asal yang dapat dijumlahkan. Jalan yang disarankan untuk musim pemberian adalah satu permintaan musiman (baris berikut) |
| Musim Lebaran dan hari raya lain | (a) **Satu permintaan musiman**, misalnya "THR dan Lebaran", dengan batas, disetujui berdua lewat alur permintaan Tier 1 (dokumen 06); templat "Mudik dan Lebaran" (dokumen 11 IDEA-08, Tier 2, isi saja) hanya mengisinya. Permintaan itu tidak perlu menyebut tempat uang sumber, sehingga tidak ada reservasi (§3.8). Setiap entri THR yang ditautkan menambah realisasinya; bagian dalam batas tidak ditandai; bagian yang membawa realisasi ke atas batas (`N`) dibandingkan dengan ambang (§3.8 "Entri yang tertaut ke otorisasi"); kelebihannya tampil sebagai "melebihi Rp…" (`authorization.overrun`). Entri yang sudah tersimpan tanpa tautan boleh ditautkan belakangan oleh pencatatnya (§3.8, INV-26). (b) **Menyiapkan uang** tanpa permintaan: alokasi bertujuan `spending` "THR Lebaran" pada tempat uangnya (§3.9), tanpa persetujuan dan tanpa entri; pembayaran yang membawa tujuan itu mengurangi alokasinya. (c) **Laporan**: tidak ada laporan musiman baru. Musim pemberian dapat melintasi dua periode anggaran; jumlah semusimnya terbaca dari realisasi permintaan musiman bila ada, selain itu dari kategorinya di ringkasan tiap periode (§6.3) |
| Uang pecahan baru | Menukar pecahan di dalam satu tempat uang tidak membuat entri. Tarik tunai untuk itu adalah pemindahan (§3.3). Biaya jasa penukaran adalah `expense` dari tempat uang yang membayarnya, berkategori pilihan rumah tangga, dan boleh ditautkan ke permintaan musiman |
| Dari dana pribadi | Pemberinya memilih, seperti pada belanja lain dari dana pribadi: (i) urusan pribadinya: `expense` di buku pribadinya, tidak terlihat pasangan (keputusan pemilik nomor 3); (ii) untuk keperluan rumah tangga tanpa minta diganti: `contribution_expense`, terlihat, dan tidak dibandingkan dengan ambang (§3.2 butir 2 akibat (1); §3.4); (iii) minta diganti: `advance`, melahirkan hak ganti (§3.5), dan boleh ditautkan ke permintaan musiman. Yang mana yang "keperluan rumah tangga" dinyatakan pemberinya sendiri; model tidak dapat memeriksanya **(batas yang dinyatakan)**: THR rumah tangga yang dicatat sebagai urusan pribadi tidak terlihat |
| Kategori | Model tidak mewajibkan kategori apa pun (T-10). Saran untuk dokumen 05 dan 08: satu saran nama kategori saat onboarding, misalnya "THR dan pemberian hari raya", yang dapat dihapus atau diganti, sama dengan saran nama kategori dokumen 11 IDEA-09 |
| Yang tidak dinilai | Apakah pemberian THR kepada pekerja rumah tangga merupakan kewajiban hukum tidak dinilai dokumen ini; modelnya mencatatnya sama dengan pemberian lain |

Contoh: CTH-51.

### 3.3 Pemindahan dan serah terima

- `transfer` hanya antara dua tempat uang berpemilik sama. Total uang pemilik itu tidak berubah (INV-06).
- **Serah terima** adalah `transfer` yang kedua tempat uangnya berbeda pemegang. Saldo kedua tempat uang berubah seketika saat dicatat.
- **Konfirmasi custody, sisi uang** (Q-19; keputusan pemilik produk, 11 November 2026: serah terima tercatat seketika, bertanda "belum dikonfirmasi" sampai penerima menerimanya; penerima dapat mengajukan *keberatan* kapan saja; sama ke dua arah). Penanda `unconfirmed` melekat pada **setiap baris yang menambah saldo tempat uang berpemegang anggota dewasa selain pencatat**, bukan hanya pada serah terima: pemasukan, pelunasan piutang, atau pengembalian dana yang dicatat ke tempat uang pasangan, entri pembalik yang mengembalikan uang ke tempat uang pasangan, **saldo awal** (`opening`) tempat uang yang dipegang pasangan, dan selisih hitung positif yang dicatat bukan pemegangnya (§2.5 baris (a)). Saldo awal termasuk di dalamnya karena kalimat umumnya memang "setiap baris yang menambah" (**[Usulan]**, putaran S3; PA-FIN-2): angka awal uang di tangan seseorang yang diisi orang lain adalah pernyataan atas uang yang ia pegang, sama seperti serah terima kepadanya. Saldo awal 0 tidak membawa penanda, karena tidak ada baris yang menambah. Contoh: CTH-42. Yang tidak membawa penanda: baris ke tempat uang yang dipegang pencatat sendiri, ke tempat uang bersama (`holder = joint`), dan ke titipan pada pihak luar. Tempat uang pribadi tidak pernah membawanya, karena hanya pemiliknya yang menulis ke sana (INV-25).
- **Sisi pembayar pada entri dua buku [Usulan]** (temuan `security-privacy` SP-05; dokumen 06 §7.2 "Aturan umum"). Penanda yang sama melekat pada **baris yang mengurangi tempat uang dana bersama berpemegang anggota dewasa selain pencatat, bila entrinya entri dua buku** (§2.5, baris (b)): M mencatat penyisihan, penggantian, pinjaman dari dana bersama, atau pelunasan utang dana bersama kepada dirinya, dari tempat uang dana bersama yang dipegang pasangannya. Alasannya: entri itu memindahkan uang dari tangan pemegang ke dana pribadi pencatat, dan pemegangnya tidak dapat membalikkannya (§3.12); tanpa penanda, uang yang ia pegang berkurang di catatan tanpa pernyataannya. Belanja biasa yang dicatat atas tempat uang pasangan (entri satu buku) **tidak** membawa penanda ini: pemegangnya dapat mengoreksinya sendiri dengan entri pembalik.
- **Selisih hitung negatif yang dicatat bukan pemegangnya** (**[Usulan]**, putaran S3; G-05; §2.5 baris (c)). Baris `cash_difference` negatif atas tempat uang yang dipegang anggota dewasa selain pencatat membawa penanda yang sama di sisi pemegang. Bedanya dari belanja biasa: belanja menyatakan pemakaian yang diketahui pencatatnya, sedangkan selisih negatif menyatakan "uang di tangan pemegang kurang" tanpa peristiwa apa pun, atas nama pemegang (§10.3 butir 3), dan membaliknya tidak memulihkan uang yang memang tidak ada. Penanda membuat jelas bahwa angka itu belum dinyatakan pemegangnya; pemegang mengonfirmasi, berkeberatan, atau menghitung sendiri. Tempat uang `joint` tidak membawanya, seperti aturan umum.
- Penanda itu tidak mengubah angka apa pun: saldo berubah seketika, dan custody menampilkan jumlahnya pada pemegang sebagai "belum dikonfirmasi" (§5), terpisah untuk baris masuk dan baris keluar. Tidak ada batas waktu dan tidak ada konfirmasi otomatis. Bila pemegang mengajukan *keberatan*, entri tetap ada dan tetap terhitung sampai dikoreksi dengan entri pembalik. Alur konfirmasi dan keberatan adalah milik dokumen 06 §7 dan §8; kaitnya `handover.recorded`, `handover.confirmed`, `handover.disputed`, yang berlaku atas semua baris bertanda itu. `handover.confirmed` hanya terpicu oleh pernyataan pemegangnya sendiri (dokumen 06 C02 dan C03). **Menutup keberatan tanpa entri pembalik (dokumen 06 C05) tidak memicu kait uang apa pun dan bukan konfirmasi**: baris itu kembali ke status sebelum keberatan, sehingga jumlahnya berpindah dari "dikeberatani" kembali ke "belum dikonfirmasi" atau ke terkonfirmasi di custody (§5.1), tanpa angka lain yang berubah.
- **Pinjam kembali** dana bersama adalah serah terima, bukan utang (default Q-01 yang diterima pemilik produk).
- **Biaya admin** (Q-48) dicatat **di entri yang sama**: tempat uang asal berkurang sebesar jumlah ditambah biaya, tempat uang tujuan bertambah sebesar jumlah, dan satu baris `expense` sebesar biaya. Pencatat mengisinya di layar yang sama. **Kategorinya bukan kategori yang diwajibkan model** (**[Usulan]**, putaran S3; PA-FIN-7; keputusan pemilik nomor 9; T-10): layar mengisinya lebih dulu dengan kategori baris biaya admin terakhir di rumah tangga itu, bila kategori itu tidak terarsip; selain itu kosong. Pencatat boleh mengganti atau mengosongkannya; baris tanpa kategori tersimpan, tidak terhitung terhadap anggaran mana pun, dan tampil di baris "tanpa kategori" laporan (§6.3). Rumusan revisi sebelumnya, "kategori bawaan 'biaya admin'", dicabut.

### 3.4 Penyisihan dan kontribusi (melintasi garis)

| Peristiwa | Entri | Buku rumah tangga | Buku pribadi | Terlihat pasangan |
|---|---|---|---|---|
| Penyisihan: dana bersama menjadi dana pribadi (misalnya uang saku) | `owner_transfer` | tempat uang −X; `owner_transfer` −X | tempat uang +X; `owner_transfer` +X | Ya: jumlah, tanggal, pelaku, penerima. Saldo pribadi sesudahnya tidak |
| Kontribusi uang: dana pribadi diserahkan menjadi dana bersama | `owner_transfer` | tempat uang +X; `owner_transfer` +X | tempat uang −X; `owner_transfer` −X | Ya |
| Kontribusi belanja: dana pribadi membayar keperluan rumah tangga, tidak minta diganti | `contribution_expense` | `owner_transfer` +X; `expense` −X dengan `funded_by = member(M)` | tempat uang −X; `owner_transfer` −X | Ya: pengeluarannya dan sumbernya ("beras, pakai uang pribadi suami"). Tempat uang pribadi yang dipakai dan saldonya tidak |

Kontribusi belanja tidak mengubah *uang yang ada* dan tidak mengubah kekayaan bersih rumah tangga: rumah tangga menerima nilai X dan langsung memakainya. Pengeluarannya terhitung **satu kali** di laporan kategori dan anggaran; kontribusinya tampil satu kali sebagai baris "kontribusi dari dana pribadi" (§6.3). Menjumlahkan keduanya sebagai dua pengeluaran adalah kesalahan yang dicegah INV-02.

Ketiga entri di atas dicatat oleh anggota yang dana pribadinya tersentuh (§2.2). Bagian pemasukan yang diarahkan langsung ke dana pribadi (§3.1) adalah penyisihan yang sama, hanya saja barisnya berada di dalam entri `income`.

Penyisihan ditandai menurut salah satu dari dua aturan, bergantung pada ada-tidaknya *jatah pribadi* penerimanya yang berlaku (di bawah). Yang dinilai adalah jumlah baris `owner_transfer` negatif di buku rumah tangga (§3.8). **Bagian pribadi di dalam entri `income` diperlakukan sama dengan penyisihan lain** (default yang diterima di checkpoint S2 lewat "lanjut", dokumen 17 D-18 dan D-48; T-13; dokumen 06 §6.4; bukan pernyataan pemilik produk sendiri), supaya mengarahkan uang saat dicatat tidak menjadi jalan memutar aturan itu. Tanda itu netral dan tidak menghalangi apa pun. Pilihan yang tidak diambil: baris `owner_transfer` di dalam entri `income` dikecualikan dari penilaian.

**Jatah pribadi bulanan.** Yang diputuskan pemilik produk (11 November 2026, keputusan 14: "Jatah pribadi dijumlah per bulan bebas bisa diatur"; menjawab Q-98) hanyalah bentuknya: tiap anggota dewasa punya paling banyak satu *jatah pribadi*, satu jumlah per *periode anggaran* (§6.1), dijumlahkan sepanjang periode itu, dan besarnya dapat diatur. Bentuk ini menggantikan rumusan dokumen ini sebelumnya (jumlah per kali, dibandingkan per entri).

**Jatah adalah batas** (keputusan pemilik produk, checkpoint S2; dokumen 17 S2-1 butir (iii) pilihan b; Q-106). Kata-katanya: "jatah itu ya memang batas, dan harus ajukan kenaikan bukan keringanan, kecuali dia split jatah + uang pribadi bisa itu". Aturannya:

| Keadaan | Aturan tanda | Kedudukan |
|---|---|---|
| M punya jatah yang berlaku; sesudah entri ini jumlah penyisihan kepadanya dalam periode itu **tidak melebihi** jatah (`terpakai_sebelum + bagian ≤ jatah`) | Tidak ditandai dan tidak ada pemberitahuan, berapa pun jumlah entrinya, juga bila di atas *ambang permintaan* | Default yang diterima di checkpoint S2 lewat "lanjut" (dokumen 17 D-48); dasar pertanyaan lead yang tidak ia bantah, bukan kata-katanya |
| M punya jatah yang berlaku; sesudah entri ini jumlahnya **melebihi** jatah (`terpakai_sebelum + bagian > jatah`) | **Ditandai, berapa pun kecilnya kelebihan itu.** Ambang tidak lagi dibandingkan bagi anggota yang punya jatah | Keputusan pemilik produk, checkpoint S2 |
| M mengakhiri jatahnya sendiri, jatah itu **di atas ambang yang berlaku saat diakhiri**, dan entri ini disimpan **sebelum lewat akhir periode anggaran yang memuat pengakhiran itu** | Jatah yang diakhiri itu **tetap menjadi dasar penilaian**, dengan jumlahnya saat diakhiri: dua baris di atas berlaku seolah jatah itu masih ada | Rumusan lead yang dikukuhkan pemilik produk, 11 November 2026 (T-29 pilihan 1). Bahwa batasnya waktu simpan di server dan bukan tanggal kejadian: **[Usulan]** `architect` (dokumen 06 §6.3 "Mengakhiri jatah sendiri"), diikuti dokumen ini |
| M mengakhiri jatahnya sendiri dan jatah itu **tidak di atas ambang yang berlaku saat diakhiri** (termasuk jatah 0, dan bila tidak ada ambang yang berlaku; jatah seperti itu hanya lahir dari usulan pemiliknya, guard K01 dokumen 06) | Pengakhiran berlaku **seketika** atas penilaian: entri yang disimpan sesudahnya dinilai tanpa jatah. Jatah seperti itu tidak pernah mengurangi tanda dibandingkan tanpa jatah, jadi melepasnya bukan jalan keluar dari batas yang lebih longgar | **[Usulan]** `security-privacy` (SP-33, SP-34), 12 November 2026; mempersempit rumusan T-29 pilihan 1 yang dikukuhkan pemilik produk, dan belum ia kukuhkan dalam bentuk ini. Bahwa pemilik jatah kecil "boleh melepasnya seketika" termasuk pertanyaan (1) lead kepadanya, yang jawabannya tidak menyebut hal ini |
| M tidak punya jatah (belum pernah disepakati; atau diakhiri pemiliknya dan periode anggaran yang memuat pengakhiran itu sudah lewat) | Tidak berubah: tiap penyisihan dibandingkan dengan ambang sendiri-sendiri, sebesar seluruh `bagian`-nya | Default yang diterima di checkpoint S2 lewat "lanjut" (D-48) |
| Penyisihan yang tertaut ke otorisasi (sudah disetujui lewat permintaan) | Tidak terhitung terhadap jatah; ia realisasi otorisasinya. Bagiannya yang berada dalam batas otorisasi tidak ditandai; bagian yang membawa realisasi ke atas batas dibandingkan dengan ambang, seperti kelebihan lain (§3.8 "Tanda atas entri") | Default yang diterima di checkpoint S2 lewat "lanjut" (D-48); perlakuan bagian di atas batas: aturan lead dari D-07, 12 November 2026 (temuan `critic` S7) |

Pada setiap baris **tidak ada yang dihalangi**: entri selalu tersimpan, dan setiap penyisihan tetap terlihat oleh keduanya, di dalam maupun di atas jatah. Jatah bukan uang dan bukan anggaran.

Yang mengikuti dari keputusan itu:

- **Yang perlu lebih mengajukan, bukan mengambil.** Tambahan untuk satu periode saja adalah *permintaan* biasa; penyisihan yang ditautkan ke otorisasinya tidak terhitung terhadap jatah dan, sampai batas otorisasi itu, tidak ditandai (default D-53 dan D-48). Kebutuhan yang menetap adalah usulan menaikkan jatah, yang berlaku setelah pasangan mengonfirmasi (keputusan 17). Tidak ada kelonggaran otomatis di atas jatah: inilah arti "bukan keringanan" (rumusan lead yang dikukuhkan pemilik produk, 11 November 2026). Tambahan untuk satu bulan tetap mungkin, lewat permintaan biasa yang disetujui.
- **Memakai uang pribadi yang sudah dimiliki tidak pernah menyentuh jatah.** Jatah hanya menghitung dana bersama yang menjadi dana pribadi. Belanja dari dana pribadi, kontribusi, dan talangan tidak terhitung dan tidak ditandai oleh aturan ini. Begitulah lead membaca "split jatah + uang pribadi" dalam kata-kata pemilik produk (**[Usulan]** lead; bacaan ini belum ia kukuhkan): kebutuhan yang lebih besar daripada sisa jatah boleh ditutup sebagian dari jatah dan sisanya dari uang pribadi yang sudah ada, tanpa tanda.
- **Jatah yang disepakati sebesar 0 [Usulan], revisi pasca-checkpoint S2.** Kesepakatan jatah yang berlaku dengan jumlah 0 berbeda dari tidak punya jatah: pada yang pertama setiap penyisihan ditandai, pada yang kedua hanya yang di atas ambang. Jatah 0 hanya dapat lahir dengan persetujuan pemilik jatah itu (di bawah), jadi tidak seorang pun dapat memasangnya atas pasangannya sendirian. **Dikukuhkan** atas permintaan dokumen 06 §15.4 butir (u) 7: `monthly_amount` pada `set_aside_schedule` dan jumlah pada kesepakatan jatah boleh 0 (bilangan bulat ≥ 0), dan jatah 0 tidak sama dengan tanpa jatah; dokumen 06 kini membedakannya dengan cara yang sama (K01, INV-A19). Kedudukannya tetap **[Usulan]**.
- **Jatah yang tidak di atas ambang yang berlaku (≤ ambang, termasuk 0) hanya dapat diusulkan pemiliknya sendiri** (**[Usulan]** `security-privacy`, SP-32 dan SP-34; disampaikan kepada pemilik produk dan tidak ia bantah, 12 November 2026). Jawabannya atas pertanyaan itu, "yaa kalau ngerasa kurang maka boleh ajukan misal suami ke istri atau sebalikanya", dibaca lead sebagai penegasan jalan kenaikan ke dua arah; ia **tidak** menyatakan pagar ini. Alasannya dari sisi hitungan: jatah yang tidak di atas ambang hanya pernah menambah tanda dibandingkan tanpa jatah, dan jatah 0 adalah ambang nol untuk satu orang. Pagarnya milik dokumen 06 (guard K01); dokumen ini hanya mengutipnya. Akibatnya atas kalimat di paragraf berikut: pasangan yang ingin jatah itu lebih kecil tetap boleh mengusulkannya, asalkan jumlahnya tetap di atas ambang yang berlaku. Batasnya sama dengan batas pengakhiran seketika di tabel atas: "tidak di atas ambang". Dari sisi uang tidak ada akibat (INV-12).
- **Pemberitahuan** atas entri yang bertanda adalah aturan dokumen 06 §6.2. Pemberitahuan langsung (N-10) hanya ada untuk tanda dasar `threshold`, paling banyak satu per penerima per hari (D-35); tanda dasar `personal_share` dan `routine_amount` tampil pada entri dan di ringkasan saja, **tidak pernah sebagai pemberitahuan langsung** (diterima pemilik produk atas saran lead, 12 November 2026; SP-37, SP-43).

**Pilihan yang tidak diambil, dicatat satu kali.** (a) *Hanya bagian di atas jatah yang dibandingkan dengan ambang* (rancangan dokumen ini dan dokumen 06 sebelum checkpoint; direkomendasikan `product-analyst` dan didukung `security-privacy`, SP-30, dengan alasan jatah lalu hanya pernah mengurangi tanda). Ditolak pemilik produk: ia memaksudkan jatah sebagai batas, dan pada pilihan itu pengambilan di atas jatah dalam potongan yang masing-masing tidak melebihi ambang tidak pernah ditandai. (b) *Bentuk tengah*: jumlah di atas jatah dibandingkan dengan ambang secara kumulatif dalam periode. Tidak diambil dengan alasan yang sama: ia tetap memberi kelonggaran sebesar ambang di atas jatah. Harga pilihan yang diambil dinyatakan di §10.3 butir 7.

**Jumlahnya disepakati berdua** (keputusan pemilik produk, 11 November 2026: "a di sepakati berdua, nanti saling konfirmasi aja"). Satu pihak mengusulkan jumlah jatah seorang anggota, pihak lain mengonfirmasi; perubahan berjalan dengan cara yang sama. Jumlah yang berlaku adalah jumlah pada *kesepakatan rutin* yang berlaku atas `set_aside_schedule` anggota itu, dibaca lewat `routine_agreement_ref` (§2.6); alur usul dan konfirmasinya milik dokumen 06 §6.3 "Kesepakatan rumah tangga". Sampai usulan baru dikonfirmasi, jumlah lama berlaku. Selama belum ada kesepakatan yang berlaku, M tidak punya jatah dan tiap penyisihan kepadanya dibandingkan dengan ambang sendiri-sendiri; sesudah kesepakatannya diakhiri hal yang sama berlaku baru sejak periode anggaran berikutnya (di bawah). **Hanya pemilik jatah yang dapat mengakhiri kesepakatan jatahnya sendiri** (default yang diterima di checkpoint S2 lewat "lanjut", dokumen 17 D-52; temuan `security-privacy` SP-24): pasangannya tidak dapat mengakhirinya sendirian; ia dapat mengusulkan jumlah yang lebih kecil, asalkan tetap di atas ambang yang berlaku (jatah yang tidak di atas ambang, termasuk 0, hanya diusulkan pemiliknya sendiri; guard K01 dokumen 06; butir di atas), dan jumlah lama tetap berlaku sampai usulan itu diterima pemilik jatah. **Dasar aturan ini berubah sesudah jatah menjadi batas.** Semula mengakhiri jatah sendiri hanya memperketat aturan atas diri pengakhirnya. Kini tidak lagi: tanpa jatah, penyisihan yang masing-masing tidak melebihi ambang tidak ditandai, sedangkan dengan jatah yang sudah habis setiap penyisihan ditandai. Mengakhiri jatah sendiri karenanya dapat mengurangi tanda atas diri sendiri. **Jalan itu ditutup** (rumusan lead yang dikukuhkan pemilik produk, 11 November 2026; T-29 pilihan 1): jatah yang diakhiri pemiliknya tetap menjadi dasar penilaian sampai akhir periode anggaran yang memuat pengakhirannya; baru sejak periode berikutnya anggota itu tidak punya jatah. Ini hanya berlaku bila jatah yang diakhiri di atas ambang yang berlaku saat diakhiri; jatah yang tidak di atas ambang berakhir seketika juga atas penilaian (**[Usulan]** `security-privacy`, SP-33; tabel di atas). Selebihnya D-52 tidak berubah. Rinciannya **[Usulan]** `architect` (dokumen 06 §6.3 "Mengakhiri jatah sendiri"), diikuti dokumen ini karena tidak ada aturan uang yang berlawanan dengannya: (a) pengakhirannya sendiri berlaku seketika sebagai *peristiwa*, terlihat dan diberitahukan; yang tertunda hanya akibatnya atas penilaian; (b) batas masa itu (`assessment_until`) adalah hari terakhir periode anggaran yang memuat waktu pengakhiran diterima server, dan yang dibandingkan dengannya adalah **waktu simpan** entri di server, bukan tanggal kejadiannya, sama dengan penentuan jatah mana yang berlaku; (c) jumlah periode (`terpakai_sebelum`) tetap dihitung menurut tanggal kejadian; (d) selama masa itu jumlahnya tetap jumlah saat diakhiri; bila jatah baru disepakati di dalam masa itu, jatah baru itulah yang dipakai sejak disetujui; (e) tidak ada entri yang dinilai ulang. Yang tidak ditutup dinyatakan di §10.3 butir 7. Dari sisi uang tidak ada yang berubah: tidak ada angka saldo yang bergantung pada jatah (INV-12). Jumlah kedua pasangan boleh berbeda; aturannya sama. Seorang anggota dapat mengajukan kenaikan jatahnya dengan menyertakan alasan, dan kenaikan itu berlaku setelah pasangannya mengonfirmasi (pernyataan pemilik produk, 11 November 2026: "bisa ajukan kenaikan sesuai alasan"); alurnya milik dokumen 06 §6.3, dan usul, alasan, konfirmasi, maupun penolakannya tidak punya akibat uang apa pun (INV-12). Jatah yang baru berlaku ke depan saja: entri yang sudah tersimpan tidak dinilai ulang. Pilihan yang tidak diambil, dicatat satu kali: tiap orang menetapkan jatahnya sendiri dan pasangannya hanya melihat.

Yang disimpan (§2.6): `set_aside_schedule` per penerima, berisi `monthly_amount` (jumlah yang direncanakan atau sedang diusulkan), pengingat opsional, dan `routine_agreement_ref`. Ia bukan entri dan tidak mengubah angka apa pun (INV-12). Uang saku anggota dewasa dan bagian pribadi dari penghasilan bukan dua hal: keduanya cara mengambil jatah yang sama, yang satu lewat entri `owner_transfer`, yang lain lewat baris di dalam entri `income` (§3.1). Jatah boleh diambil dalam berapa pun bagian sepanjang periode.

**Hitungannya.** Untuk entri `e` dan anggota M:

```text
bagian(e, M)    = − Σ effect baris owner_transfer negatif di buku rumah tangga pada e
                  yang pasangannya di buku M, bila e berjenis owner_transfer atau income;
                  0 untuk jenis lain (claim_waiver dan exit_closure bukan penyisihan)
terhitung(e)    = bagian(e, M) > 0, e bukan entri pembalik, dan e tidak tertaut ke otorisasi
                  -- penyisihan yang disetujui lewat permintaan menjadi realisasi otorisasinya (§3.8)
ada_jatah(e, M) = ada kesepakatan jatah M yang berlaku saat e disimpan, berapa pun jumlahnya,
                  termasuk 0;
                  -- jatah ≤ ambang yang berlaku, termasuk 0, hanya lahir dari usulan pemiliknya
                  --   sendiri (guard K01 dokumen 06; [Usulan] security-privacy)
                  ATAU kesepakatan jatah M yang terakhir diakhiri pemiliknya, belum ada
                  kesepakatan baru yang berlaku, kesepakatan itu punya assessment_until,
                  dan waktu simpan e (waktu server) belum melewatinya
                  -- assessment_until hanya ada bila, saat diakhiri, ada ambang yang berlaku
                  --   dan jumlah jatah itu > ambang itu (SP-33, [Usulan]); nilainya hari
                  --   terakhir periode anggaran yang memuat waktu pengakhiran diterima
                  --   server (zona waktu rumah tangga); T-29 pilihan 1
                  -- jatah yang diakhiri dengan jumlah ≤ ambang (termasuk 0): tanpa
                  --   assessment_until; pengakhirannya berlaku seketika
jatah(e, M)     = jumlah pada kesepakatan itu (yang berlaku, atau yang diakhiri selama
                  masa penilaiannya); 0 bila tidak ada
terpakai_sebelum(e, M)
                = Σ bagian(f, M) atas entri f yang terhitung, disimpan lebih dulu daripada e,
                  berada di periode anggaran yang sama dengan e, dan belum dibalik saat e disimpan
dalam_jatah(e)  = min(bagian(e, M), max(0, jatah(e, M) − terpakai_sebelum(e, M)))
luar_jatah(e)   = bagian(e, M) − dalam_jatah(e)

ditandai(e)     = bila ada_jatah(e, M):  luar_jatah(e) > 0
                      -- setara dengan terpakai_sebelum(e, M) + bagian(e, M) > jatah(e, M),
                      --   karena bagian > 0; ambang tidak dibandingkan
                      -- keputusan pemilik produk, checkpoint S2 (S2-1 butir (iii) pilihan b)
                  bila tidak:            bagian(e, M) > ambang permintaan yang berlaku
                      -- per entri, tidak dijumlahkan; tidak berubah
                  hanya dinilai bila terhitung(e), jadi hanya atas entri yang tidak tertaut
                  ke otorisasi. Penyisihan yang tertaut ke otorisasi tidak dinilai
                  terhadap jatah: bagiannya dalam batas otorisasi tidak ditandai, dan
                  bagian yang membawa realisasi ke atas batas dibandingkan dengan ambang
                  (§3.8 "Tanda atas entri")
```

Pemeriksaan kesetaraan: `luar_jatah > 0` berarti `bagian > max(0, jatah − terpakai_sebelum)`. Bila `terpakai_sebelum ≤ jatah`, itu sama dengan `terpakai_sebelum + bagian > jatah`. Bila `terpakai_sebelum > jatah`, kedua sisi benar untuk setiap `bagian > 0`.

Semua bilangan bulat; tidak ada pembagian. **Setiap penyisihan kepada M terhitung dengan sendirinya**, baik bagian pribadi di dalam entri `income` maupun entri `owner_transfer`: keanggotaannya diturunkan dari barisnya, bukan dari tautan. Pencatat tidak menambah langkah, dan tidak dapat memilih entri mana yang memakan jatah dengan menautkan atau tidak menautkannya. `links.set_aside_schedule_id` tetap sah untuk pengingat dan tidak mengubah hitungan. Yang tidak terhitung hanya dua: penyisihan yang tertaut ke otorisasi, dan `claim_waiver` (dicatat pasangannya, bukan pengambilan oleh M; §3.6). Kontribusi M dari dana pribadi tidak mengurangi hitungan: tidak ada angka bersih (§6.3). Aturan tanda `ditandai(e)` di atas kini **sama** dengan dokumen 06 revisi kedelapan (§6.2, dasar tanda `personal_share`: `lebih > 0`; `ada_jatah` termasuk jatah 0 dan jatah yang diakhiri sampai `assessment_until`). `ditandai(e)` di bagian ini hanya menilai penyisihan; tanda atas pemakaian lain, termasuk pembayaran tagihan berulang, ada di §3.8 "Tanda atas entri". Pembagian `dalam_jatah` dan `luar_jatah` sendiri tidak berubah dan sama dengan dokumen 06 §6.2 dan §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi": `terpakai_sebelum` adalah `jumlah_sebelum` di sana, dan `luar_jatah` adalah `lebih`, karena `bagian − min(bagian, max(0, jatah − sebelum)) = min(bagian, max(0, sebelum + bagian − jatah))`. Menjawab dokumen 06 §15.4 butir (o): *jumlah pemakaian* entri `owner_transfer` dan bagian pribadi entri `income` seluruhnya penyisihan kepada satu anggota; kedua jenis entri itu tidak memuat baris `expense` (tidak ada biaya admin di dalamnya, §2.5), sehingga `P − J = 0` bagi keduanya. Sebagai realisasi otorisasi, entri yang tertaut tetap terhitung penuh sebesar `bagian`-nya (§3.8). Contoh terhitung: CTH-39.

**"Periode yang sama" menurut tanggal apa.** Menurut **tanggal kejadian** entri. Ini ditetapkan dokumen 06 §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi", dan sama dengan dasar semua angka periode di dokumen ini. Tabel berikut mencatat alasannya dan dasar yang tidak dipilih.

| Dasar | Yang disediakan model | Akibat |
|---|---|---|
| Tanggal kejadian (`event_date`); **dipilih** | Periode `e` = periode anggaran yang memuat `event_date`-nya; urutan di dalam periode menurut waktu simpan | Sama dengan dasar anggaran dan laporan: jumlah `bagian` satu periode = baris "disisihkan ke dana pribadi M" pada ringkasan periode itu, dikurangi yang tertaut otorisasi. Entri terlambat memakai jatah periode kejadiannya, bukan periode berjalan. Celahnya: tanggal kejadian diisi pencatat, sehingga penyisihan dapat diberi tanggal periode lampau yang jatahnya belum terpakai; entri itu tampil sebagai entri terlambat dengan kedua tanggalnya |
| Waktu simpan (`recorded_at`); tidak dipilih | Periode `e` = periode anggaran yang memuat tanggal `recorded_at`-nya | Tidak dapat dimundurkan. Tetapi penghasilan akhir periode yang baru dicatat sesudah periode berganti memakan jatah periode baru, dan jumlah yang dipakai aturan jatah tidak lagi sama dengan baris ringkasan periode mana pun |

**Dinilai satu kali.** `dalam_jatah` dan `luar_jatah` dihitung saat `e` disimpan, dari entri yang sudah diterima server saat itu; hasilnya disimpan dokumen 06 sebagai catatan penilaian (dokumen 06 INV-A37). Entri yang datang kemudian, termasuk entri terlambat, entri pembalik, dan perubahan jatah, tidak mengubah pembagian entri yang sudah ada. **Entri pembalik melonggarkan jatah ke depan saja**: sejak pembalik tersimpan, entri asalnya tidak lagi masuk `terpakai_sebelum` entri berikutnya. Satu-satunya penilaian ulang adalah atas entri itu sendiri, bila tautan susulannya ke otorisasi ditambah atau dicabut (dokumen 06 §6.3).

**Jalan menghindar yang diterima: tanggal mundur.** Penyisihan hari ini dapat diberi tanggal kejadian periode lampau yang jatahnya belum terpakai. Model tidak dapat memeriksa tanggal kejadian, dan batas "paling jauh mundur" akan menghalangi entri terlambat yang jujur (§3.11). Yang melawannya hanya keterlihatan, dan untuk itu model menjamin dua hal (menjawab dokumen 06 §15.4 butir (s); disempurnakan menurut temuan `security-privacy` SP-28): setiap entri **menyimpan** `recorded_at` yang persis, diisi server dan tidak dapat diatur pencatat (§2.5, INV-22); dan setiap entri yang tanggal pencatatannya berbeda dari tanggal kejadiannya **menampilkan** tanggal pencatatan itu, hanya tanggalnya, sama bagi kedua pasangan (§3.11). Jam pencatatan tidak ditampilkan kepada siapa pun.

**Cocok dengan ringkasan** (menjawab dokumen 06 §15.4 butir (r)). Untuk satu periode dan satu anggota, baris "disisihkan ke dana pribadi" tabel A (§6.3) = jumlah `bagian` entri yang terhitung terhadap jatah dan tidak dibalik + jumlah penyisihan yang tertaut ke otorisasi. Dasarnya sama: tanggal kejadian; entri pembalik bertanggal kejadian sama dengan entri asalnya dan mengurangi baris yang sama (§6.3).

**Batas yang dinyatakan.** (1) Jatah tidak berpindah periode: yang tidak terpakai hangus, dan kelebihan satu periode tidak mengurangi periode berikutnya. (2) Tidak ada angka jatah atau pemakaiannya yang lintas periode, kumulatif, atau bersih per anggota (§6.3; SP-13); yang tampil hanya periode yang sedang dilihat, dan angkanya tidak membuka hal baru: ia jumlah penyisihan periode itu, yang memang sudah terlihat (§5.3, B1). (3) Jatah tidak punya akibat pada saldo, anggaran, *terikat*, reservasi, atau kekayaan bersih (INV-12, INV-30). (4) Periode peralihan saat hari mulai periode diganti (§6.1) membawa jatah penuh, tidak diprorata.

### 3.5 Talangan, hak ganti, dan penggantian

*Talangan* kini berarti: pembayaran dari dana pribadi untuk keperluan rumah tangga atau untuk anggota lain **yang ditandai "minta diganti"**. Tanpa tanda itu ia kontribusi (§3.4).

| Langkah | Buku pembayar | Buku yang ditalangi |
|---|---|---|
| Talangan X (`advance`) | tempat uang −X; piutang (hak ganti) +X | kewajiban (hak ganti) +X; `expense` −X |
| Penggantian Y ≤ sisa (`reimbursement`) | tempat uang +Y; piutang −Y | tempat uang −Y; kewajiban −Y |

Talangan dicatat oleh yang menalangi; penggantian dari dana bersama dicatat oleh **penerimanya**; bila tempat uang dana bersama yang berkurang dipegang anggota dewasa lain, baris itu bertanda `unconfirmed` sampai pemegangnya mengonfirmasi (§2.2, §3.3; dokumen 06 §7.2 "Aturan umum"). Talangan tidak menyentuh tempat uang dana bersama dan tidak membawa penanda; kewajiban yang dilahirkannya dapat dikeberatani dan tampil demikian (§3.12). Pengeluaran tercatat satu kali, saat talangan. Penggantian bukan pengeluaran dan bukan pemasukan. Hak ganti tercatat sejak talangan dicatat (Q-21), sebagai kewajiban rumah tangga kepada anggota itu, sehingga langsung mengurangi kekayaan bersih rumah tangga.

Bawaan tanda "minta diganti" **[Usulan]**: mati untuk belanja biasa (keputusan pemilik nomor 3); hidup bila entri ditautkan ke otorisasi atas dana bersama (§3.8, situasi C), karena persetujuannya menyatakan dana bersama yang menanggung. Pencatat dapat mengubahnya.

### 3.6 Pelepasan hak ganti dan klaim antar-pemilik lain

Melepas klaim berarti pemiliknya menghadiahkan jumlah itu kepada pihak yang berutang. Entri `claim_waiver`: piutang −X dan `owner_transfer` −X di buku yang melepas; kewajiban −X dan `owner_transfer` +X di buku yang dibebaskan.

Bila yang dibebaskan adalah rumah tangga, kekayaan bersih rumah tangga naik X dan laporan menampilkannya sebagai "kontribusi dari dana pribadi" (menjawab temuan N-14: kenaikan itu punya baris bernama, bukan muncul tanpa sebab). Hasil akhirnya sama persis dengan kontribusi belanja sejak awal (CTH-16).

Hanya pemilik piutang yang dapat melepasnya. Pihak yang berutang tidak dapat menghapus utangnya sendiri. Untuk klaim yang pemilik piutangnya dana bersama (dana bersama meminjamkan kepada M), yang melepas adalah anggota dewasa selain M; inilah satu-satunya entri yang dicatat anggota lain dan menulis ke buku pribadi M (§2.2; di luar itu hanya entri sistem `exit_closure`, §8.2), dan sebagai pemakaian dana bersama ia mengikuti ambang permintaan sebesar jumlah yang dilepas: *jumlah pemakaian*-nya adalah baris `owner_transfer` negatif di buku rumah tangga (§3.8). Di laporan ia tampil sebagai baris tersendiri, "utang dihapuskan", **bukan** sebagai "disisihkan ke dana pribadi M" (§6.3; temuan `security-privacy` SP-13): tidak ada uang yang berpindah ke M pada saat itu, dan menggabungkannya akan membesarkan angka penyisihan M. Arah sebaliknya, M melepas hak gantinya atas dana bersama, bukan pemakaian dana bersama dan tidak dibandingkan dengan ambang.

### 3.7 Utang, piutang, pelunasan, penghapusan

- `loan_given` kepada pihak luar: uang −X, piutang +X; kekayaan bersih tetap.
- `repayment_in`: uang +Y, piutang −Y, dengan Y ≤ sisa piutang. Kelebihan bayar dicatat sebagai `income` bersumber `other`.
- `write_off` (Q-47): piutang −Z, sebab `write_off` −Z. Dilaporkan sebagai baris tersendiri, "piutang diikhlaskan", pada bulan tanggal kejadian penghapusannya; tidak dihitung terhadap anggaran mana pun dan tidak muncul di arus kas, karena uangnya sudah keluar saat dipinjamkan. Tidak digabung dengan kiriman keluarga, supaya "berapa yang kita kirim" dan "berapa yang tidak kembali" tetap dua jawaban.
- **Penghapusan piutang dana bersama dibandingkan dengan ambang** (usulan putaran S3; G-03; T-32; mengadopsi bawaan dokumen 05; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-9 (ii)). *Jumlah pemakaian* `write_off` atas piutang berpemilik `shared` adalah jumlah yang dihapus, dan ia dinilai dengan dasar `threshold` seperti pemakaian dana bersama lain (§3.8), **tetapi secara kumulatif per piutang** (usulan `security-privacy` SP-51, diadopsi putaran perbaikan S3; diterima lewat "lanjut" di checkpoint S3, S3-9 (ii); INV-34; T-32). Yang dibandingkan dengan ambang adalah jumlah bagian yang belum disetujui (`N`, §3.8) atas **semua** penghapusan pada piutang yang sama yang sudah tersimpan dan belum dibalik, ditambah bagian entri ini; entri ditandai bila jumlah itu melebihi ambang dan entri ini sendiri punya bagian yang belum disetujui (`N > 0`). Akibatnya piutang 600000 yang diikhlaskan dalam dua kali 300000 dengan ambang 500000 menandai entri kedua (300000 + 300000 = 600000 > 500000) dan **setiap** penghapusan berikutnya atas piutang itu, walau masing-masing di bawah ambang. Alasan kumulatif: berbeda dari kiriman keluarga, penghapusan selalu menunjuk satu piutang yang jumlahnya diketahui, jadi memecah satu keputusan mengikhlaskan menjadi potongan di bawah ambang adalah jalan menghindar yang dapat dikenali; aturan per entri membiarkannya. Batasnya dinyatakan: (1) kumulatif seumur piutang, tidak per periode; (2) yang dijumlahkan hanya penghapusan atas posisi `receivable` yang sama; piutang tersendiri kepada orang yang sama mulai dari nol, dan pinjaman yang melahirkannya sudah dinilai sendiri saat dipinjamkan (§3.8 baris `loan_given`); (3) bagian yang sudah disetujui lewat permintaan tidak ikut dijumlahkan dan tidak pernah ditandai, karena persetujuan itu pernyataan pasangan atas jumlah itu (§3.8 "Entri yang tertaut ke otorisasi"); (4) penghapusan yang sudah dibalik sebelum entri ini disimpan tidak ikut dijumlahkan; (5) dinilai satu kali saat disimpan, seperti tanda lain. Contoh: CTH-25 varian SP-51. Alasan dasar `threshold` itu sendiri: pada saat dihapus, akibatnya bagi rumah tangga sama dengan memberi pihak luar itu uang sebesar itu (kekayaan bersih turun sebesar kiriman yang sama), dan kiriman keluarga sebesar itu dibandingkan dengan ambang; tanpa ini satu orang dapat menurunkan kekayaan bersih rumah tangga sebesar apa pun tanpa tanda, sedangkan pelepasan klaim atas seorang anggota sudah dibandingkan (§3.6). Pinjaman dan penghapusannya adalah dua keputusan, jadi keduanya dinilai sendiri-sendiri: meminjamkan menukar uang dengan piutang, menghapus melepaskan piutang itu. Penghapusan boleh diajukan lebih dulu sebagai permintaan dan ditautkan ke otorisasinya, kecuali ke otorisasi yang sudah memuat `loan_given` asal piutang itu: otorisasi itu menyetujui meminjamkan, bukan mengikhlaskan (INV-26). Penghapusan tetap tidak terhitung terhadap anggaran dan tidak muncul di arus kas. Siapa yang boleh mencatatnya adalah aturan dokumen 06; dari sisi uang anggota dewasa mana pun, karena piutang itu milik dana bersama. Piutang dana bersama kepada seorang anggota tidak dihapus lewat `write_off` melainkan dilepas lewat `claim_waiver` (§3.6).
- `loan_received` dan `repayment_out` adalah cerminnya di sisi kewajiban.
- **Utang antar-pasangan** hanya lahir bila uang berpindah antara dua pemilik berbeda dengan janji kembali: `loan_given` di buku pemberi dan `loan_received` di buku penerima, dalam satu entri. Kedua pihak selalu melihatnya (Q-24). Bila satu sisinya dana bersama, entrinya dicatat anggota yang dana pribadinya tersentuh; bila kedua sisinya dana pribadi, alur pencatatannya ditunda dari MVP (§2.2, T-12).

### 3.8 Otorisasi, reservasi, penyelesaian

Dokumen 06 memiliki state machine permintaan. Bagian ini hanya menetapkan akibat uang tiap peristiwanya.

| Peristiwa (kait untuk dokumen 06) | Akibat uang |
|---|---|
| `request.submitted`, `request.countered`, `request.rejected` | Tidak ada |
| `request.approved` | Lahir `authorization` berstatus `open`. Tidak ada entri, tidak ada saldo yang berubah (INV-12). Bila permintaan menyebut tempat uang sumber, lahir satu `reservation` aktif sebesar jumlah yang disetujui |
| `handover` untuk permintaan itu | Entri `transfer`. Reservasi dilepas sebesar yang diserahkan |
| Belanja ditautkan ke otorisasi | Entri biasa (`expense`, `loan_given`, `asset_purchase`, `owner_transfer`, `advance`, …). Reservasi yang masih aktif dilepas sebesar belanja itu, paling banyak sampai habis |
| `authorization.amended` (perubahan syarat diterima; dokumen 06 B12) | Tidak ada entri dan tidak ada saldo yang berubah (INV-12). Batas, kategori, tujuan, tempat uang sumber, tanggal rencana, dan masa berlaku otorisasi menjadi nilai versi syarat yang baru. *Terikat* dihitung ulang dari batas baru dan mengikuti periode tanggal rencana yang baru (§6.1). Reservasi: bila tempat uang sumbernya tetap, reservasi aktif berubah sebesar perubahan batas, paling rendah sampai nol; bila sumbernya berganti, reservasi lama dilepas dan lahir reservasi sebesar *terikat* yang baru pada sumber baru; bila sumbernya dihapus, reservasi dilepas. Bila batas baru lebih kecil daripada realisasi, selisihnya adalah kelebihan dan `authorization.overrun` terpicu; bila batas baru tidak lebih kecil daripada realisasi, kelebihan yang pernah tampil hilang. Entri yang sudah tertaut tidak berubah |
| `authorization.settled` | Otorisasi ditutup; *terikat* menjadi nol; reservasi yang tersisa dilepas. Penyelesaian sendiri tidak membuat entri; entri hanya ada bila sisa dikembalikan (situasi B) |
| `authorization.cancelled`, `authorization.expired` | Sama dengan `settled` tanpa belanja: terikat nol, reservasi dilepas. Belanja yang sudah tertaut tetap tercatat |
| `authorization.overrun` (realisasi melampaui batas) | Tidak ada entri tambahan. Lihat "kelebihan" di bawah |

**Reservasi** (Q-12). Persetujuan **tidak** membuat reservasi, kecuali permintaannya menyebut tempat uang sumber; itulah "aturan eksplisit" yang diminta brief §7. Reservasi mengurangi *saldo tersedia* tempat uang itu, tidak pernah saldo tercatat. Ia tidak menghalangi pencatatan apa pun (A-20): bila pemegang membelanjakan uang yang direservasi untuk hal lain, saldo tersedia menjadi negatif dan tampil sebagai "kurang Rp… untuk yang sudah dijanjikan".

**Realisasi** otorisasi = jumlah *jumlah pemakaian* entri yang ditautkan kepadanya, apa pun jenisnya (pengeluaran, pinjaman yang diberikan, pembelian aset, penyisihan, pembayaran cicilan), dikurangi pengembalian dana yang tertaut pada entri-entri itu. **Terikat** = max(0, batas − realisasi) selama otorisasi `open`; nol sesudahnya.

**Jumlah pemakaian** satu entri adalah satu angka yang dipakai untuk dua hal: sebagai realisasi bila entri itu tertaut ke otorisasi, dan sebagai jumlah yang dinilai untuk tanda bila tidak. Penilaian itu punya **tiga dasar**: *ambang permintaan*, *jatah pribadi* (§3.4), dan jumlah per kali *tagihan berulang* yang disepakati (§3.15); lihat "Tanda atas entri" di bawah tabel. Alur dan pemberitahuannya milik dokumen 06 §6.2.

| Jenis entri | Jumlah pemakaian | Dibandingkan dengan ambang? |
|---|---|---|
| Baris `expense` berpemilik `shared` yang ditanggung dana bersama, **di entri jenis apa pun**: entri `expense`; baris denda di dalam `repayment_out`; baris biaya admin di dalam `transfer`; baris `expense` di dalam `recategorization` yang merinci selisih kas (§3.10). Tidak termasuk baris `expense` pada `contribution_expense` dan `advance`, yang punya barisnya sendiri di bawah | Jumlah bersih baris itu (`−Σ effect`), bila positif. Bagian pelunasan dan bagian pemindahan dari entri yang sama tidak termasuk | Ya, **kecuali** entri itu tertaut ke tagihan berulang yang kesepakatan rutinnya berlaku: maka **tidak dibandingkan dengan ambang**; entri ditandai tepat bila jumlah pemakaiannya melebihi jumlah per kali yang disepakati, berapa pun kelebihannya dan berapa pun ambangnya (rumusan lead yang dikukuhkan pemilik produk, 11 November 2026; §3.15) |
| `loan_given` dari dana bersama | Jumlah pinjaman | Ya |
| `write_off` atas piutang dana bersama kepada pihak luar (§3.7) | Jumlah yang dihapus | Ya (usulan putaran S3; G-03), **kumulatif per piutang**: yang dibandingkan adalah `hapus_sebelum + N`, bukan `N` saja (usulan `security-privacy` SP-51; "Tanda atas entri" di bawah); keduanya diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-9 (ii) |
| `owner_transfer` keluar dari buku rumah tangga (penyisihan) | Jumlah baris `owner_transfer` negatif di buku rumah tangga | Bergantung pada jatah penerimanya (§3.4). **Dengan jatah yang berlaku: tidak dibandingkan dengan ambang**; entri ditandai bila `luar_jatah > 0`, yaitu bila jumlah penyisihan periode itu melebihi jatah, berapa pun kelebihannya (keputusan pemilik produk, checkpoint S2, S2-1 butir (iii) pilihan b); di dalam jatah tidak ditandai (default D-48). **Tanpa jatah yang berlaku: ya**, sebesar seluruh jumlahnya, per entri |
| `income` dengan bagian pribadi (§3.1) | Jumlah baris `owner_transfer` negatif di buku rumah tangga, yaitu bagian pribadinya | Sama persis dengan baris penyisihan di atas (default yang diterima di checkpoint S2 lewat "lanjut": T-13; dokumen 17 D-18, D-48; aturan di atas jatah: S2-1 butir (iii) pilihan b; §3.4) |
| `claim_waiver` atas klaim dana bersama kepada seorang anggota (§3.6) | Jumlah yang dilepas, yaitu baris `owner_transfer` negatif di buku rumah tangga | Ya |
| `claim_waiver` oleh anggota atas hak gantinya kepada dana bersama (§3.6) | Nol: baris `owner_transfer` di buku rumah tangga positif; nilai masuk, bukan keluar | Tidak: bukan pemakaian dana bersama |
| `advance` atas keperluan rumah tangga | Jumlah hak ganti yang lahir | Ya |
| `asset_purchase` tunai | Jumlah yang dibayar | Ya |
| `asset_purchase` secara cicilan | **Uang muka ditambah jumlah seluruh cicilan** (= harga ditambah biaya kredit), yaitu seluruh dana bersama yang dipakai dan diikat entri itu | Ya |
| `contribution_expense` | Jumlah baris `expense` | Tidak: bukan dana bersama |
| `repayment_out` (cicilan, utang): bagian yang mengurangi kewajiban | Jumlah pelunasan itu; menjadi realisasi bila ditautkan ke otorisasi | **Tidak.** Pelunasan kewajiban yang sudah tercatat tidak pernah dibandingkan dengan ambang dan tidak memerlukan kesepakatan rutin (dokumen 06 §6.2, §6.3). Denda di entri yang sama adalah baris `expense` dan mengikuti baris pertama tabel ini |
| `opening` atas kewajiban dana bersama kepada pihak luar | Jumlah saldo awal kewajiban itu, satu kali, pada entri `opening`-nya | Ya, terhadap ambang yang berlaku saat entri disimpan (dokumen 06 §6.2, putusan 2). Saldo awal yang diisi saat onboarding, sebelum ada ambang yang disepakati, tidak pernah ditandai, juga tidak surut sesudah ambang berlaku. Pembayaran kewajiban itu sesudahnya adalah pelunasan dan tidak dibandingkan |
| `opening` atas tempat uang, piutang, aset, dan kewajiban lain | Nol | Tidak |
| `exit_closure` (§8.2) | Nol: entri sistem, bukan pemakaian oleh siapa pun | Tidak; tidak ditandai dan tidak dapat ditautkan ke otorisasi |
| `debt_forgiven` | Nol: nilai masuk, bukan keluar | Tidak |
| Bagian pemindahan pada `transfer`; `reimbursement`, `allocation`, `refund`, dan lainnya | Nol; tidak ditautkan sebagai pemakaian | Tidak |

**Entri yang memuat lebih dari satu bagian.** Pelunasan dengan denda dan pemindahan dengan biaya admin memuat bagian dari dua baris tabel. Sebagai realisasi, *jumlah pemakaian* entri itu adalah jumlah semua bagiannya (pelunasan ditambah denda; biaya admin saja pada pemindahan, juga bila pemindahan itu serah terima untuk suatu permintaan). Yang dibandingkan dengan ambang hanyalah bagian yang kolom ketiganya "ya". Setiap baris entri terhitung paling banyak satu kali: pada `advance`, baris `expense` dan hak ganti yang lahir adalah jumlah yang sama dan dihitung sebagai satu angka. Pengembalian dana dan entri pembalik tidak punya jumlah pemakaian sendiri; keduanya mengurangi realisasi entri asalnya (§3.12, §3.13).

**Tanda atas entri.** Satu fungsi, tiga dasar; nama dasarnya sama dengan `flag.basis` di dokumen 06 §6.2. Untuk entri `e` dengan *jumlah pemakaian* `P`:

```text
tertaut      = e tertaut ke satu otorisasi A (tautan yang berlaku)
B            = batas A saat e dinilai
real_sebelum = realisasi A saat e dinilai, tanpa e
dalam_batas  = min(P, max(0, B − real_sebelum)) bila tertaut; selain itu 0
N            = P − dalam_batas
               -- bagian e yang dinilai. Tidak tertaut: N = P.
               -- Tertaut: N = bagian e yang membawa realisasi A ke atas batasnya

J      = bagian(e, M) bila terhitung(e) dan ada_jatah(e, M); selain itu 0          (§3.4)
         -- terhitung(e) menuntut e tidak tertaut, jadi J = 0 bila tertaut
lebih  = luar_jatah(e) bila J > 0; selain itu 0                                     (§3.4)
rutin  = e tidak tertaut ke otorisasi, e bukan contribution_expense, dan e tertaut ke
         satu tagihan berulang yang kesepakatan rutinnya berlaku saat e disimpan, atau
         saat tautan susulannya ditambahkan                                          (§3.15)
         -- contribution_expense boleh tertaut (pengingat), tetapi tidak dinilai: ia
         --   bukan pemakaian dana bersama (tabel jumlah pemakaian, kolom ketiga)
R      = jumlah per kali pada kesepakatan itu     -- kekerapan tidak diperiksa

H      = Σ N(f) atas setiap entri write_off f pada piutang yang sama dengan e, yang
         disimpan lebih dulu daripada e dan belum dibalik saat e disimpan, bila e
         berjenis write_off; selain itu 0                                           (§3.7)
         -- N(f) adalah N entri f saat f dinilai: bagian yang belum disetujui
         -- usulan security-privacy SP-51, putaran perbaikan S3;
         --   diterima lewat "lanjut" di checkpoint S3 (S3-9 (ii))

e ditandai bila rumah tangga punya ≥ dua anggota dewasa saat e disimpan,
dan salah satu berikut benar:
  dasar threshold       : bukan rutin; ada ambang yang berlaku; N − J > 0; H + N − J > ambang
                          -- untuk jenis selain write_off H = 0, jadi syarat ini sama dengan
                          --   N − J > ambang seperti sebelumnya
  dasar personal_share  : lebih > 0
  dasar routine_amount  : rutin; P > R
```

Pemeriksaan kesetaraan untuk jenis selain `write_off`: `H = 0`, dan `N − J > ambang` dengan ambang ≥ 0 sudah memuat `N − J > 0`, jadi tidak ada tanda lain yang berubah. Untuk `write_off`, `J = 0` (bukan penyisihan), sehingga syaratnya `N > 0` dan `H + N > ambang`.

**Entri yang tertaut ke otorisasi** (aturan lead dari dokumen 17 D-07, 12 November 2026; temuan `critic` S7). Bagian entri yang berada dalam batas otorisasinya **tidak pernah ditandai**, dengan dasar apa pun. Bagian yang membawa realisasi ke atas batas (`N`) dibandingkan dengan ambang, dengan dasar `threshold`; ia tidak dinilai terhadap jatah maupun terhadap jumlah per kali tagihan berulang. Rumusan revisi sebelumnya, "e tidak tertaut ke otorisasi" sebagai syarat setiap tanda, **keliru** dan dicabut: ia membuat kelebihan sebesar apa pun atas batas otorisasi tidak pernah bertanda. Yang dihitung untuk `N` hanya entri ini: kelebihan yang sudah ada sebelum `e` ditanggung entri yang melahirkannya. Bahwa entri yang tertaut ke otorisasi dan ke tagihan berulang sekaligus dinilai menurut otorisasinya saja adalah **[Usulan]** dokumen ini; dokumen 06 mengikutinya.

Aturan pelengkapnya:

- **Paling banyak satu tanda per entri**, dengan satu dasar. Ketiga dasar tidak dapat benar bersamaan: entri penyisihan tidak ditautkan ke tagihan berulang; bagi penyisihan yang tidak tertaut `N − J` adalah 0 bila ada jatah; dan bagi entri yang tertaut ke otorisasi hanya dasar `threshold` yang dinilai.
- **Kelebihan atas batas otorisasi dinilai per entri, tidak dijumlahkan.** Dua entri yang masing-masing membawa kelebihan tidak di atas ambang tidak ditandai, walau jumlah kelebihannya di atas ambang; kelebihan itu tetap tampil pada permintaan sebagai "melebihi Rp…" (`authorization.overrun`). Ini sama dengan sifat dasar `threshold` atas entri lain: per entri. **Satu-satunya pengecualian adalah `write_off`**, yang bagian belum disetujuinya dijumlahkan per piutang (`H`, §3.7; SP-51).
- **Pos rutin dinilai per pembayaran.** Tidak ada jumlah periode dan tidak ada catatan penilaian: tandanya fungsi dari entri itu dan kesepakatan yang berlaku saja. Dua pembayaran atas tagihan yang sama dalam satu periode masing-masing dibandingkan dengan `R`.
- **Tagihan berulang tanpa kesepakatan yang berlaku** (usulannya belum dijawab, ditolak, ditarik, atau kesepakatannya diakhiri): `rutin` salah, sehingga pembayarannya dinilai dengan dasar `threshold` persis seperti entri yang sama tanpa tautan itu (**[Usulan]** `architect`, dokumen 06 §6.3 "Pos rutin" baris a dan d, INV-A43; diikuti dokumen ini).
- **Pembayaran cicilan dan pelunasan utang** (`repayment_out`) tidak pernah ditandai (D-44); dendanya adalah baris `expense` dan mengikuti dasar `threshold`.
- **Satu anggota dewasa:** tidak ada tanda apa pun.
- **Dinilai satu kali**, saat entri disimpan (dokumen 06 INV-A37), dan sekali lagi hanya ketika tautan susulan ditambahkan atau dicabut. Mengganti ambang, jatah, jumlah per kali, atau batas otorisasi, dan pengembalian dana atau entri pembalik atas entri lain, tidak mengubah tanda entri yang sudah tersimpan.
- **Pemberitahuan** adalah aturan dokumen 06. Tanda dasar `personal_share` dan `routine_amount` tampil pada entri dan di ringkasan saja, tidak pernah sebagai pemberitahuan langsung (diterima pemilik produk atas saran lead, 12 November 2026).
- **Tautan susulan ke tagihan berulang** menghapus tanda dasar `threshold` hanya bila `P ≤ R`; bila `P > R` dasarnya berganti menjadi `routine_amount`. Entri di bawah ambang yang tidak bertanda dapat **menjadi bertanda** lewat tautan itu; karena itu tautan susulan ke tagihan berulang hanya diterima dari pencatat entrinya ("Penautan" butir 3).
- Tanda tidak menahan uang, tidak menolak entri, dan tidak mengubah saldo, anggaran, *terikat*, atau reservasi (INV-12).

Untuk pembelian secara cicilan, yang dibandingkan dengan ambang bukan harga barang (terlalu kecil sebesar biaya kredit) dan bukan uang muka saja (pembelian besar akan lolos dengan uang muka kecil). Pembayaran cicilan atas pembelian yang sudah tertaut penuh ke suatu otorisasi tidak ditautkan lagi ke otorisasi itu, supaya tidak terhitung dua kali (CTH-34). Hal yang sama berlaku atas kewajiban yang masuk sebagai saldo awal dan entri `opening`-nya tertaut ke suatu otorisasi: pembayarannya tidak ditautkan lagi ke otorisasi itu. Pada kedua kasus ini aturan server, bukan aturan layar saja: perintah tautannya ditolak dan entri pembayarannya tetap tersimpan (INV-26; dokumen 06 §4.5 "Masa berlaku dan kedaluwarsa").

**Dua batas baris saldo awal kewajiban, dinyatakan terbuka.** (1) Model tidak membedakan utang lama yang baru teringat dari utang baru; keduanya dibandingkan dengan cara yang sama, dan tandanya hanya hilang lewat tautan ke otorisasi atau entri pembalik (dokumen 06 §6.2). (2) Utang uang yang dicatat sebagai `loan_received` tidak punya jumlah pemakaian dan tidak dibandingkan dengan ambang, sedangkan utang yang sama yang dicatat sebagai saldo awal kewajiban dibandingkan. Ketimpangan ini tidak membuka jalan menghindar: jalan yang tidak ditandai adalah jalan yang mencatat uangnya masuk ke tempat uang dana bersama, terlihat oleh keduanya. Apakah `loan_received` perlu ikut dibandingkan adalah aturan dokumen 06 (§10.2).

**Penautan entri ke otorisasi.**

1. **Saat dicatat.** Entri boleh ditautkan ke otorisasi yang `open`. Entri terlambat boleh pula ditautkan ke otorisasi yang sudah `settled` atau `expired`, asalkan tanggal kejadiannya tidak lebih awal dari tanggal persetujuan dan tidak lebih baru dari tanggal penutupan otorisasi itu. Otorisasi yang `cancelled` tidak menerima tautan: pemohonnya sendiri yang menariknya; entrinya tetap tersimpan tanpa tautan (pencatatan tidak pernah ditolak). **Tautan ke otorisasi yang `open` tidak bersyarat tanggal**: entri yang tanggal kejadiannya lebih awal dari tanggal persetujuan boleh ditautkan kepadanya. Inilah yang dipakai *konfirmasi susulan* dokumen 06 (§6.2, INV-A40; usulan lead; jalurnya diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-6 (i) sampai (v)): belanja yang sudah terjadi diajukan sesudahnya, dan saat disetujui entrinya ditautkan ke otorisasi baru itu. **Syarat tautan, termasuk syarat tanggal, diperiksa hanya saat tautan dibuat.** Tautan yang sah saat dibuat tetap sah ketika otorisasinya kemudian `settled` atau `expired`, sehingga tanda entri itu tidak muncul lagi saat otorisasi ditutup. Tidak ada aturan uang yang berlawanan dengannya: realisasi tidak bergantung pada tanggal maupun status (INV-14), anggaran menghitung entri pada periode tanggal kejadiannya sendiri (§3.11), dan reservasi serta *terikat* dihitung dari realisasi yang sama. Tautan yang lahir di dalam transaksi persetujuan adalah `entry_link` biasa dengan kait yang sudah ada; dokumen ini tidak meminta kait tersendiri.

   **Konfirmasi susulan: disetujui dan ditutup dalam satu transaksi** (**[Usulan]** `security-privacy`, SP-38, dan dokumen ini; dokumen 06 §6.2 mengikutinya). Jawaban pemilik produk atas pertanyaan 3 lead (12 November 2026, diterima atas saran lead) hanya mencakup tanpa pengingat otomatis, satu pengingat ulang, dan tidak dapat diminta lagi sesudah "belum sepakat"; penutupan seketika, batas yang ditetapkan server, dan satu tautan tidak termasuk di dalamnya. Checkpoint S3 (S3-6, diterima lewat "lanjut") juga tidak menyebutnya: ketiganya tetap **[Usulan]**. Otorisasi yang lahir dari konfirmasi susulan ditutup di dalam transaksi persetujuannya, dengan batas tepat sebesar *jumlah pemakaian* entri yang dikonfirmasi. Dari sisi uang ini sah dengan aturan yang ada, dalam urutan berikut di dalam satu transaksi: `request.approved` (otorisasi `open`, batas `B = P` entri itu) → `entry_link` dari entri itu ke otorisasi (sah: otorisasi `open`, tanpa syarat tanggal) → `authorization.settled`. Akibatnya: realisasi = `B`; *terikat* = max(0, B − B) = 0 sebelum dan sesudah ditutup; sisa 0; kelebihan 0, jadi `authorization.overrun` tidak terpicu; tidak ada entri baru dan tidak ada saldo yang berubah (INV-12); tanda entri itu hilang karena `N = P − min(P, B) = 0`. Tiga syarat yang membuatnya tepat, **[Usulan]** dokumen ini untuk dokumen 06:
   - permintaannya **tidak menyebut tempat uang sumber**, sehingga tidak ada reservasi yang lahir (uangnya sudah keluar); permintaan konfirmasi susulan yang menyebut tempat uang sumber **ditolak saat diajukan** (dokumen 06 B01), tanpa akibat apa pun (**[Usulan]** `security-privacy`, SP-38, dan dokumen ini);
   - otorisasi itu **tidak menerima tautan lain**, juga bukan dari entri terlambat yang tanggal kejadiannya sama dengan tanggal persetujuannya; tanpa syarat ini aturan "tanggal kejadian di antara tanggal persetujuan dan tanggal penutupan" masih membuka satu hari, dan tautan kedua akan melahirkan kelebihan atas otorisasi yang dimaksudkan memuat tepat satu entri (INV-26);
   - entri yang sudah tertaut ke otorisasi lain tidak dapat dikonfirmasi susulan (satu entri, paling banyak satu otorisasi).

   Sesudah ditutup, aturan biasa tetap berlaku atas entri itu: pengembalian dana dan entri pembaliknya mengurangi realisasi (§3.12, §3.13), dan tautannya dapat dicabut oleh pembuat tautan atau pencatat entrinya (butir 3), yang mengembalikan penilaian tanda entri itu seperti tanpa tautan. Tidak ada yang menghalangi rancangan ini dari sisi uang. Pelaku `entry_link` itu adalah **pemohon**, dan pemohon harus pencatat entri yang dikonfirmasi (dokumen 06 §15.4 butir (v)); dengan itu tautannya memenuhi butir 3 (tautan susulan ke otorisasi hanya dari pencatat entri atau pemohon otorisasi) tanpa pengecualian.
2. **Akibat tautan ke otorisasi yang sudah tertutup.** Status otorisasi tidak berubah dan tidak ada entri tambahan. Realisasi dihitung ulang (INV-14 tidak bergantung pada status); *terikat* tetap nol; tidak ada reservasi yang hidup kembali. Sisa atau kelebihan pada permintaan dihitung ulang, dan `authorization.overrun` terpicu bila realisasi untuk pertama kali melampaui batas. Anggaran menghitung entri itu pada periode tanggal kejadiannya, seperti entri terlambat lain (§3.11).
3. **Tautan susulan.** Tautan boleh ditambahkan pada entri yang sudah tersimpan **tanpa mengubah entri itu**: ia catatan `entry_link` tersendiri yang hanya ditambah (§2.5), dengan pelaku, waktu, dan idempotency key-nya sendiri. Syarat tanggalnya sama dengan butir 1. Tautan susulan yang keliru dicabut dengan `entry_link` lain yang menunjuknya (`revokes`); tautan yang dibawa entri sejak dicatat hanya dapat dibatalkan dengan entri pembalik. Tautan yang berlaku atas suatu entri = tautan pada entri itu ditambah tautan susulan yang tidak dicabut; satu entri tertaut ke paling banyak satu otorisasi (INV-26). Menambah atau mencabut tautan susulan menghitung ulang realisasi, *terikat*, dan reservasi persis seperti bila entri itu dicatat dengan tautan tersebut, dan tidak mengubah saldo apa pun. **Siapa yang boleh** ditetapkan dokumen 06 §4.5 dan INV-A34, dan dikutip di sini tanpa diubah: tautan susulan **ke otorisasi** hanya diterima dari pencatat entri itu atau pemohon otorisasi itu; pencabutan tautan susulan hanya dari pembuat tautan itu atau pencatat entrinya; tautan susulan **ke tagihan berulang** hanya diterima dari pencatat entri itu (**[Usulan]** `architect`, revisi ketujuh dokumen 06; diikuti dokumen ini), karena tautan itu kini dapat memasang tanda dasar `routine_amount` pada entri yang semula tidak bertanda; tautan susulan ke jadwal penyisihan tetap boleh ditambahkan anggota dewasa mana pun, karena ia tidak mengubah hitungan jatah maupun tanda (§3.4). Perintah dari pelaku yang tidak berwenang ditolak tanpa akibat apa pun pada tautan, realisasi, *terikat*, atau reservasi. Pelakunya selalu terlihat. Dari sisi uang pembatasan ini tidak mengubah hitungan apa pun; ia hanya mempersempit siapa yang dapat menggerakkan realisasi suatu otorisasi lewat tautan.

**Sisa = batas − realisasi, bila positif.** Arti "diselesaikan dengan benar" (brief §5D langkah 7) per situasi custody:

| Situasi | Keadaan | Sisa | Entri saat penyelesaian | Contoh |
|---|---|---|---|---|
| A | Pemohon membelanjakan dana bersama yang sudah ia pegang | Tidak pernah berpindah; tetap di tempat uang pemohon sebagai dana bersama | Tidak ada | CTH-21 |
| B | Dana bersama diserahkan kepada pemohon untuk permintaan itu | Berada di tempat uang pemohon, tetap dana bersama, dan terlihat di custody. Pemohon memilih: (1) mengembalikannya, atau (2) tetap memegangnya | (1) satu `transfer` sebesar sisa; (2) tidak ada | CTH-22 |
| C | Pemohon membayar dengan dana pribadinya atas otorisasi dana bersama | Tidak ada uang sisa. Hak ganti lahir sebesar **realisasi**, bukan sebesar batas | Tidak ada; penggantian adalah entri tersendiri | CTH-23 |
| D | Pemegang lain membayar langsung dari tempat uang sumber yang direservasi | Tetap di tempat uang sumber; reservasinya dilepas | Tidak ada | §7.5, varian CTH-22 |

Pada semua situasi: sisa tidak pernah menjadi pemasukan, tidak pernah menjadi milik pribadi pemohon, dan tidak mengubah *uang yang ada*. Menjadikan sisa sebagai dana pribadi adalah penyisihan (§3.4) dan mengikuti aturannya.

**Kelebihan = realisasi − batas, bila positif** (Q-13). Pencatatan tidak dihalangi. Seluruh belanja, termasuk kelebihannya, tercatat sebagai entri biasa dan terhitung penuh terhadap anggaran kategorinya. Kelebihan tampil pada permintaan sebagai "melebihi Rp…".

| Kelebihan dibayar dari | Perlakuan |
|---|---|
| Dana bersama | Pengeluaran dana bersama biasa; tidak ada klaim. Dokumen 06 memutuskan apakah ia memerlukan tanggapan pasangan; uangnya sudah tercatat benar apa pun jawabannya |
| Dana pribadi pemohon | Kontribusi (§3.4) secara bawaan; talangan bila pemohon menandai "minta diganti" untuk bagian itu |

### 3.9 Alokasi dan tabungan

- Alokasi adalah baris `allocation` (tempat uang, tujuan, ±jumlah). Jumlah alokasi per (tempat uang, tujuan) tidak pernah negatif.
- **Batas** (Q-45): saat perintah alokasi dicatat, jumlah seluruh alokasi pada satu tempat uang tidak boleh melebihi saldo tercatatnya; perintah yang melampaui ditolak.
- Entri yang mengurangi tempat uang boleh membawa tujuan; alokasi tujuan itu ikut berkurang sebesar jumlahnya, paling banyak sampai nol. Entri tanpa tujuan mengurangi bagian *belum dialokasikan*.
- **Belum dialokasikan** = saldo tercatat − jumlah alokasi. Bila belanja membuatnya negatif, tidak ada yang ditolak dan tidak ada alokasi yang dikurangi otomatis: tempat uang itu tampil "alokasi melebihi uang Rp…" sampai rumah tangga mengurangi salah satu alokasinya. Aplikasi tidak menebak tujuan mana yang dikorbankan.
- Pemindahan boleh membawa tujuan: alokasi berkurang di asal dan bertambah di tujuan dalam entri yang sama.
- Anggaran tidak memerlukan alokasi (keputusan pemilik nomor 8: tanpa amplop). Alokasi bersifat pilihan.
- **Tabungan** adalah tujuan berjenis `savings`. Menabung adalah alokasi, atau `transfer` ke tempat uang khusus yang membawa tujuan itu; tidak pernah pengeluaran. Kemajuan = jumlah alokasi tujuan itu di semua tempat uang; *target* tidak masuk angka mana pun.

**Menyelesaikan atau membatalkan tujuan** (**[Usulan]**, putaran S3; G-01; T-38; mengadopsi bawaan dokumen 05 FR-FIN-128). Berlaku atas tujuan jenis apa pun, termasuk tujuan tabungan.

| Hal | Aturan |
|---|---|
| Langkahnya | Satu langkah menampilkan alokasi tujuan itu per tempat uang dan menawarkan melepas semuanya, **tanpa pilihan terisi**. Melepas adalah satu entri `allocation` dengan baris −jumlah per tempat uang; tidak ada uang yang berpindah dan tidak ada saldo yang berubah (INV-12, INV-17) |
| Arsip | Tujuan baru terarsip bila jumlah alokasinya nol (§2.6). Tanpa melepas, tujuan tetap aktif dan tidak ada yang berubah. Tidak ada alokasi yang dilepas otomatis |
| Selesai karena uangnya dipakai | Pengeluaran atau pembelian aset yang membawa tujuan itu mengurangi alokasinya (butir ketiga di atas); sisanya dilepas dengan langkah yang sama |
| Selesai lawan batal | `archive_reason` (`achieved` atau `cancelled`) hanya keterangan untuk daftar tujuan; tidak mengubah angka atau laporan apa pun |
| Siapa | Anggota dewasa mana pun atas tujuan dan tempat uang dana bersama; alokasi di tempat uang pribadi hanya oleh pemiliknya (INV-25). Aturan pelakunya milik dokumen 06 |
| Tabel D (§6.3) | Tujuan tabungan tampil di tabel D setiap periode yang alokasinya bukan nol pada awal atau akhir periode, atau berubah di dalamnya, **termasuk bila tujuan itu terarsip** sesudahnya; pada periode pelepasannya ia tampil dengan perubahan negatifnya dan keterangan "diarsipkan". Sesudah periode itu ia tidak tampil lagi |

Contoh: CTH-45.

### 3.10 Rekonsiliasi kas, selisih kas, dan entri terlambat

Keputusan pemilik nomor 7: catatan tidak pernah disunting; beda antara terhitung dan tercatat menjadi satu entri *selisih kas* yang terlihat; entri terlambat bertanggal kejadian asli kelak menguranginya; siapa yang lupa diselesaikan di luar aplikasi.

1. Pemegang mengisi saldo terhitung. `selisih = terhitung − tercatat`, dengan `tercatat` dihitung server (di bawah, "Hitungan yang tiba belakangan"). **Siapa yang boleh menghitung** (usulan putaran S3; G-05; T-33; menyatukan butir ini dengan §2.3; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-9 (i)): untuk tempat uang dana bersama, pemegangnya, atau anggota dewasa mana pun bila pemegangnya `joint` atau pihak luar (`entrusted_external`). **Hitungan oleh anggota dewasa selain pemegang** atas tempat uang dana bersama yang dipegang anggota dewasa lain diterima hanya bila ketiga syarat ini benar (usulan `security-privacy` SP-50 (a) dan (b), diadopsi putaran perbaikan S3; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-9 (i); INV-33; T-33):

   | # | Syarat | Bila tidak terpenuhi |
   |---|---|---|
   | (a) | Pemegangnya sudah menandai tempat uang itu "pasanganku dapat melihat saldo rekening ini" (`holding.others_may_count = true`, §2.3). **Bawaan mati.** Hanya pemegang yang menyalakan dan mematikannya, kapan saja; perubahannya *peristiwa* metadata yang terlihat, tanpa entri (INV-12). Tanda itu hanya tersedia pada `kind = bank` atau `ewallet`: uang tunai yang dipegang orang lain tidak pernah dihitung bukan pemegangnya (sesuai rumusan S3-9 (i): aplikasi tidak menawarkan menghitung uang tunai di dompet pasangan) | Perintah ditolak tanpa akibat apa pun |
   | (b) | Tidak ada **hitungan bukan pemegang yang masih terbuka** atas tempat uang itu. *Terbuka*: rekonsiliasi oleh bukan pemegang yang selisihnya bukan nol, dan baris selisihnya belum dijawab pemegang, yaitu belum dikonfirmasi, belum dikeberatani, dan pemegang belum mencatat rekonsiliasinya sendiri atas tempat uang itu sesudahnya. Hitungan "dihitung, cocok" tidak pernah terbuka. Akibatnya paling banyak satu terbuka per tempat uang, dari siapa pun | Perintah ditolak tanpa akibat apa pun, termasuk hitungan yang cocok, sampai pemegang menjawab |
   | (c) | Pemegangnya anggota aktif: tidak dalam masa tenggang keluar dan belum keluar (§8.2 butir 2b), karena ia tidak dapat lagi menjawab | Perintah ditolak; tempat uang itu diselesaikan anggota yang tersisa menurut §8.2 butir 5 |

   Mematikan tanda (a) berlaku atas perintah sesudahnya: rekonsiliasi lama tetap, dan hitungan yang masih terbuka tetap terbuka sampai dijawab. Selisih yang dicatat bukan pemegang membawa penanda `unconfirmed` di sisi pemegang, positif maupun negatif (§2.5 baris (a) dan (c), §3.3). Hitungan yang cocok (selisih nol) tidak membuat entri dan tidak membawa penanda; ia tampil "dihitung, cocok" beserta pelakunya. Tempat uang pribadi hanya dihitung pemiliknya (INV-25). Alasan kedua syarat itu: tanpa (a), pasangan dapat mencatat "uang di tanganmu kurang" atas uang yang tidak dapat ia lihat; tanpa (b), ia dapat mencatatnya berulang kali sebelum pemegang sempat menjawab, dan tiap catatan menjadi tanda baru atas nama pemegang (W-26). Rumusan lama butir ini ("pemegang mengisi") dan §2.3 ("pasangannya tetap dapat merekonsiliasi") kini satu aturan.
2. Bila selisih bukan nol, dicatat entri `cash_difference`: tempat uang ±selisih, sebab `cash_difference` ±selisih, tanggal kejadian = tanggal menghitung. Sesudahnya saldo tercatat = saldo terhitung.
3. **Selisih negatif** (uang keluar yang tidak tercatat) menurunkan kekayaan bersih dan terhitung dalam total uang keluar bulan itu sebagai "belum dirinci".
4. **Selisih positif** (uang masuk yang tidak tercatat, atau pengeluaran tercatat terlalu besar) menaikkan kekayaan bersih dan tampil sebagai baris tersendiri di sisi uang masuk; ia **bukan** pemasukan dan tidak mengurangi pengeluaran kategori mana pun.
5. **Merinci selisih** (`recategorization`): dua baris sebab tanpa baris posisi. Untuk selisih negatif: `cash_difference` +X (mengurangi selisih yang belum dirinci) dan `expense` −X dengan kategori dan tanggal kejadian aslinya. Saldo tidak berubah, karena uangnya sudah keluar dari catatan lewat selisih itu. Untuk selisih positif cerminnya: `cash_difference` −X dan `income` +X, atau `expense` +X bila ternyata suatu pengeluaran tercatat terlalu besar.
6. **Entri terlambat yang menutup selisih.** Saat pencatat memasukkan pengeluaran pada suatu tempat uang dengan tanggal kejadian tidak lebih baru daripada rekonsiliasi terakhirnya, dan rekonsiliasi itu masih punya selisih negatif yang belum dirinci, aplikasi menawarkan "ini bagian dari selisih hitung Rp…?" dengan bawaan ya. Bila ya, entrinya adalah `recategorization` (butir 5) sampai sebesar selisih yang belum dirinci; bagian yang melebihinya dicatat sebagai `expense` biasa yang mengurangi tempat uang. Bila tidak, seluruhnya `expense` biasa. Tawaran ini juga berlaku atas tempat uang yang sudah terarsip, karena `recategorization` tidak punya baris posisi (§2.6, INV-21); bagian yang melebihi selisih memerlukan pengaktifan kembali tempat uang itu (§2.6, CTH-43).
7. Selisih yang belum dirinci = |selisih| − jumlah yang sudah dirinci; tidak pernah negatif (INV-16).
8. Selisih kas dihitung terhadap anggaran hanya setelah dirinci ke kategori, pada tanggal kejadian yang diisi. Sebelum itu, **layar anggaran menampilkan satu keterangan** "ada selisih yang belum dirinci Rp…", sebesar jumlah selisih negatif yang belum dirinci atas tempat uang dana bersama yang tanggal kejadian selisihnya berada di periode yang dilihat, dan tempat uang yang bersangkutan menampilkan bagiannya sendiri. Keterangan itu tidak menempel pada kategori mana pun (**[Usulan]**, putaran S3; PA-FIN-8; mengadopsi bawaan dokumen 05). Rumusan lama, "anggaran kategori yang lazim memakai tempat uang itu", dicabut: model tidak punya hubungan antara tempat uang dan kategori, dan menebaknya akan menempelkan uang yang belum diketahui keperluannya pada satu kategori. Selisih positif tidak masuk keterangan ini, karena tidak pernah mengurangi pemakaian kategori mana pun (butir 4).

Rekonsiliasi berlaku untuk semua jenis tempat uang, bukan hanya uang tunai.

**Hitungan yang tiba belakangan** (**[Usulan]**, putaran S3; G-02; bagian bersama dengan `architect`). Perintah rekonsiliasi membawa **saldo terhitung**, bukan selisih. Server menghitung selisihnya saat perintah diterima:

```text
tercatat(h, t)  = Σ delta baris tempat uang h yang sudah diterima server saat perintah
                  rekonsiliasi diterima, dari entri bertanggal kejadian ≤ t
                  -- t = tanggal hitung (tanggal counted_at, zona waktu rumah tangga)
selisih         = terhitung − tercatat(h, t)
recorded_amount_at_count = tercatat(h, t)          -- disimpan pada reconciliation
```

| Keadaan | Aturan |
|---|---|
| Entri lain atas tempat uang itu sampai di server **sebelum** hitungan, walau dicatat sesudah perangkat penghitung terakhir sinkron (misalnya serah terima dari pasangan) | Masuk `tercatat`. Selisih dihitung terhadap saldo server, **bukan** saldo yang tampil di perangkat penghitung. Alasannya: uang yang dihitung adalah kenyataan fisik, dan entri yang sudah ada di server menyatakan bagian dari kenyataan itu; menghitung terhadap saldo perangkat akan mencatat uang yang sama dua kali (CTH-46) |
| Selisih yang tampil di perangkat saat menghitung | Perkiraan dari saldo yang dikenal perangkat. Sesudah sinkron yang berlaku adalah selisih server; bila keduanya berbeda, pencatat diberi tahu hasilnya (cara dan kata-katanya: dokumen 06, 08, 09) |
| Entri bertanggal kejadian ≤ `t` yang sampai di server **sesudah** hitungan | Entri terlambat biasa. Bila berupa pengeluaran dan rekonsiliasi itu masih punya selisih negatif yang belum dirinci, ia ditawarkan sebagai bagian dari selisih (butir 6). Bila pencatatnya baru mengetahui rekonsiliasi itu sesudah entrinya tersimpan, tawaran yang sama diberikan sesudah sinkron dan dijalankan dengan dua perintah yang sudah ada: entri pembalik atas pengeluaran itu, lalu `recategorization` sebesar jumlah yang sama dengan tanggal dan kategori yang sama. Entri jenis lain tersimpan biasa; selisih yang ditimbulkannya tertangkap hitungan berikutnya |
| Entri bertanggal sama dengan `t` | Dianggap terjadi sebelum hitungan bila sudah di server saat hitungan diterima, dan sesudahnya bila tiba belakangan. Model tidak mengenal jam kejadian; urutan yang keliru pada hari yang sama tertangkap hitungan berikutnya **(batas yang dinyatakan)** |
| Dompet campuran | Kedua rekonsiliasi dihitung sendiri-sendiri dengan aturan yang sama atas saldo terhitung masing-masing; pembagian `C` di perangkat tidak berubah (CTH-40) |

INV-16 berlaku atas `tercatat(h, t)`: sesudah entri `cash_difference`, saldo tercatat per tanggal hitung sama dengan saldo terhitung. Yang diandaikan dokumen ini atas sisi `architect` (dokumen 06 §12, dokumen 09 di S4): perintah dari satu perangkat tiba di server menurut urutan pencatatannya, sehingga entri penghitung sendiri yang dicatat sebelum hitungan selalu masuk `tercatat`; server satu-satunya yang menghitung selisih (06 S-1); pemberitahuan hasil dan kata-katanya milik dokumen 06 dan 08.

**Menghitung lalu merinci seluruh selisih sekaligus** (usulan putaran S3; menjawab dokumen 11 §10.1 butir 1, IDEA-01; IDEA-01 beserta bentuk uang ini diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-2 (i), Tier 1, hanya untuk uang tunai berpemilik tunggal dan hanya bila seluruh selisih dirinci). Sah, dan bentuknya **bukan** `cash_difference` disusul `recategorization`, melainkan:

| Hal | Aturan |
|---|---|
| Bentuk | Satu perintah: rekonsiliasi beserta **satu entri `expense`** atas tempat uang itu, sebesar selisih negatif yang seluruhnya dirinci pemegang ke satu kategori atau lebih, bertanggal kejadian = tanggal hitung, membawa `reconciliation_id`. Sesudah entri itu `tercatat` = terhitung, jadi tidak ada entri `cash_difference` dan tidak ada selisih yang belum dirinci (INV-16) |
| Alasan bentuk itu | Dengan `cash_difference` lalu `recategorization`, arus kas menampilkan uang itu sebagai "selisih hitung negatif" (§6.4), padahal keperluannya sudah dinyatakan; dengan `expense` ia tampil sebagai belanja di semua laporan, dan kata "selisih" tidak muncul di mana pun (PRN-18). Satu perintah tetap satu entri (§2.5) |
| Syarat | Selisih negatif dan jumlah baris kategori **sama persis** dengan selisih yang dihitung server. Bila server menghitung selisih lain (entri lain tiba lebih dulu, di atas), perintah itu disimpan sebagai rekonsiliasi biasa dengan `cash_difference` sebesar selisih server, dan perinciannya ditawarkan lagi sesudah sinkron sebagai `recategorization` (butir 5). Perincian sebagian tidak memakai bentuk ini: pemegang memakai jalur biasa (selisih penuh, lalu merinci bagian yang ia ketahui) |
| Ringkasan bulanan (§6.3) | Tabel A dan B: belanja per kategori; tidak ada baris "selisih hitung". Arus kas: "belanja". Anggaran: terhitung pada periode tanggal hitung. Ambang: baris `expense`-nya dinilai seperti belanja lain (§3.8), sama dengan baris `expense` dalam `recategorization` |
| Yang hilang | Tanggal dan rincian tiap pembelian; semuanya bertanggal hari menghitung **(batas yang dinyatakan)** |
| Dompet campuran | Dari sisi uang sah atas tiap tempat uang sesudah pembagian, karena tiap bagian adalah rekonsiliasi biasa. Apakah layar dompet campuran menawarkannya adalah keputusan dokumen 05 dan 08; dokumen 11 mengusulkan tidak. S3-2 (i) yang diterima membatasi IDEA-01 pada uang tunai berpemilik tunggal; dokumen ini membacanya sebagai tidak termasuk dompet campuran, sesuai usulan dokumen 11 (bacaan, bukan aturan uang; penegasannya milik dokumen 05) |

Contoh: CTH-48.

**Dompet campuran** (§2.3; S2-6). Satu hitungan atas dompet yang memuat dana bersama dan dana pribadi pemegangnya menghasilkan dua rekonsiliasi menurut butir 1 sampai 8, satu per tempat uang. Hitungan di bawah dilakukan di perangkat pemegang; server hanya menerima kedua saldo terhitung, tidak pernah `C` (§2.3; dokumen 06 V-9, INV-A41):

```text
C                = seluruh isi dompet yang dihitung pemegang
selisih          = C − (tercatat(bersama) + tercatat(pribadi))
selisih_pribadi  = ditetapkan pemegang setiap kali selisih ≠ 0; tidak ada bawaan   -- T-03
                   -- "bersama": 0;  "pribadi": selisih;  "dibagi": jumlah yang ia isi,
                   --   bertanda sama dengan selisih dan |selisih_pribadi| ≤ |selisih|
selisih_bersama  = selisih − selisih_pribadi
terhitung(bersama) = tercatat(bersama) + selisih_bersama
terhitung(pribadi) = tercatat(pribadi) + selisih_pribadi
                 -- jadi terhitung(bersama) + terhitung(pribadi) = C
```

Semua bilangan bulat. Butir 3 sampai 8 berlaku atas tiap bagian sendiri-sendiri: selisih bagian bersama terlihat keduanya, masuk ringkasan, dan dapat dirinci atau ditutup entri terlambat; selisih bagian pribadi hanya ada di buku pribadi. `C` tidak dikirim ke server dan tidak tersimpan di buku mana pun. Contoh: CTH-40.

### 3.11 Entri terlambat dan periode

- Entri terlambat boleh sejauh apa pun ke belakang, tetapi tidak lebih awal dari saldo awal tempat uang yang disentuhnya.
- Tidak ada penutupan periode. Laporan periode lampau dihitung dari tanggal kejadian, sehingga ikut berubah, dan menampilkan "berubah sejak terakhir dilihat" beserta daftar entri yang `recorded_at`-nya lebih baru daripada akhir periode itu.
- Tanggal kejadian di masa depan ditolak; rencana ke depan adalah tagihan berulang, cicilan, atau permintaan.
- **Tanggal pencatatan tampil bila berbeda dari tanggal kejadian** (temuan `security-privacy` SP-28). Setiap entri menyimpan `recorded_at` yang persis dari server (§2.5, INV-22). Yang **ditampilkan** hanyalah tanggal pencatatan (tanggal dari `recorded_at`, tanpa jam), dan hanya pada entri yang tanggal pencatatannya berbeda dari tanggal kejadiannya; entri itu tampil sebagai entri terlambat dengan kedua tanggalnya, sama bagi kedua pasangan. Entri yang dicatat pada hari kejadiannya hanya menampilkan tanggal kejadian. `recorded_at` yang persis tidak ditampilkan kepada siapa pun, pada entri maupun riwayatnya, karena jam pencatatan yang terlihat menjadi tanda "terakhir aktif" seseorang; server tetap memakainya untuk urutan, idempotency, dan audit. Ini syarat model bagi dokumen 08: tanggal kejadian diisi pencatat dan menentukan periode anggaran, laporan, dan jatah pribadi (§3.4), sehingga tanggal pencatatan adalah satu-satunya hal yang membuat tanggal mundur terlihat.

### 3.12 Koreksi dan keberatan

- **Entri pembalik**: salinan entri asal dengan semua `delta` dan `effect` berlawanan tanda, tanggal kejadian sama dengan entri asal, `corrects` menunjuk entri asal. Satu entri dibalik paling banyak satu kali; entri pembalik tidak dapat dibalik (pencatat mencatat ulang entri yang benar).
- **Entri penyesuaian**: entri sejenis yang hanya memuat selisihnya, dengan `corrects` menunjuk entri asal.
- **Saldo awal hanya disesuaikan, tidak dibalik** (**[Usulan]**, putaran S3; §3.17, INV-31): penyesuaiannya bertanggal kejadian sama dengan saldo awal asli.
- Yang berubah bukan jumlah (kategori, penerima manfaat) diperbaiki dengan `recategorization`.
- Catatan teks dan struk bukan inti moneter; penggantiannya dicatat di jejak audit tanpa entri (§8).
- Kedua pasangan boleh menambahkan koreksi atas entri dana bersama buatan siapa pun; entri asal tetap, pasangan diberi tahu, dan *keberatan* dapat diajukan (default Q-26 yang diterima pemilik produk). Keberatan tidak mengubah angka apa pun; alurnya milik dokumen 06 (kait `entry.disputed`). Entri buku pribadi hanya dapat dikoreksi pemiliknya.
- **Entri dua buku** (kontribusi, talangan, penggantian, penyisihan, pemasukan dengan bagian pribadi, pinjaman antar-pemilik) memuat baris buku pribadi M. Entri pembaliknya menyalin baris itu, jadi hanya M yang dapat mencatatnya (INV-25). Pasangannya tetap dapat (a) menambahkan entri penyesuaian yang seluruh barisnya di buku rumah tangga, misalnya `income` +500000 atas gaji yang tercatat terlalu kecil; dan (b) mengajukan keberatan. Penyesuaian sepihak atas bagian yang berpasangan dengan buku pribadi (baris `owner_transfer`, posisi klaim internal) tertolak dengan sendirinya oleh INV-03 dan INV-08. Batas ini dinyatakan terbuka: "keduanya boleh mengoreksi" berlaku penuh atas entri yang seluruhnya di buku rumah tangga; atas entri dua buku, pasangan dapat menyesuaikan sisi rumah tangganya dan berkeberatan, tetapi tidak dapat membatalkannya sendirian.
- Entri koreksi yang menambah saldo tempat uang berpemegang anggota dewasa selain pengoreksi membawa penanda `unconfirmed` seperti entri lain (§3.3).
- **Klaim internal yang dikeberatani [Usulan]** (temuan `security-privacy` SP-05; dokumen 06 §8.2 "Keberatan"). Selama ada keberatan terbuka atas entri yang memuat baris suatu klaim internal (talangan yang melahirkan hak ganti, pinjaman antar-pemilik, penggantian, pelunasan, pelepasan), klaim itu tampil di daftar utang dan piutang rumah tangga (§5.2; §6.3 tabel E) dengan keterangan "dikeberatani" dan jumlah baris klaim pada entri yang dikeberatani itu. **Tidak ada angka yang berubah**: saldo klaim, jumlah kewajiban dan piutang rumah tangga, dan kekayaan bersih tetap menghitungnya penuh sampai ada entri pembalik; tampilannya sama bagi keduanya. Keterangan itu hilang ketika keberatannya ditutup atau entri sasarannya dibalik.
- Reservasi, anggaran, dan alokasi mengikuti entri: entri pembalik atas belanja yang tertaut ke otorisasi mengurangi realisasinya.

### 3.13 Pengembalian dana

Q-15, K-02. `refund` menambah tempat uang dan membawa baris `expense` bertanda + dengan kategori entri asalnya; wajib tertaut ke entri asal. Jumlah seluruh pengembalian atas satu entri tidak melebihi jumlah entri itu (INV-13). Ia mengurangi pemakaian anggaran kategori itu pada periode **tanggal kejadian pengembalian**, bukan periode belanja aslinya, sehingga laporan lama hanya berubah karena entri terlambat, tidak karena peristiwa baru. Bila pengeluaran aslinya tidak ada di aplikasi, uang yang kembali dicatat sebagai `income` bersumber `other`. Pengembalian ke tempat uang lain daripada asalnya diperbolehkan.

### 3.14 Cicilan

Q-16, K-11.

- Kewajiban cicilan dicatat sebesar **jumlah seluruh pembayaran yang tersisa**, tanpa memisahkan pokok dan bunga. Keluarga tahu "1.500.000 kali 6 lagi"; mereka jarang tahu rincian bunganya **[Asumsi]**.
- Pembayaran cicilan adalah `repayment_out`: uang −X, kewajiban −X. Ia membawa kategori dan **terhitung terhadap anggaran kategori itu** serta tampil di belanja per kategori (§6), walaupun bukan baris sebab `expense`. Denda keterlambatan adalah baris `expense` di entri yang sama. **Kategorinya adalah `installment.category_id`** yang dipilih rumah tangga saat mendaftarkan cicilan, bukan kategori yang diwajibkan model (**[Usulan]**, putaran S3; PA-FIN-7; keputusan pemilik nomor 9; T-10; menggantikan "bawaan 'cicilan'"). Bila kategori itu kosong atau sudah terarsip, pembayaran tetap tersimpan tanpa kategori dan layar menawarkan memilih kategori; pembayaran tanpa kategori tidak terhitung terhadap anggaran mana pun tetapi tetap tampil di tabel B pada baris "tanpa kategori" (§6.3), sehingga setiap rupiah cicilan tetap terhitung satu kali di laporan pengeluaran.
- **Biaya kredit.** Saat barang dibeli secara cicilan di dalam aplikasi: aset bertambah sebesar harga barang, uang berkurang sebesar uang muka, kewajiban bertambah sebesar jumlah seluruh cicilan, dan selisihnya, `credit_cost = uang muka + jumlah seluruh cicilan − harga`, dicatat sebagai baris sebab di entri yang sama. Dengan itu turunnya kekayaan bersih sebesar bunga **punya entri yang menjelaskannya**, pada tanggal pembelian. Ia tampil di ringkasan bulanan sebagai baris "biaya kredit" dan tidak dihitung terhadap anggaran, karena uangnya keluar lewat cicilan bulanan yang sudah dihitung.
- **Biaya kredit tidak pernah negatif** (**[Usulan]**, putaran S3; G-07; T-35; mengadopsi bawaan dokumen 05 FR-FIN-111). Perintah `asset_purchase` secara cicilan dengan `uang muka + jumlah seluruh cicilan < harga` ditolak di batas seluruhnya, dengan keterangan bahwa yang dibayar lebih kecil dari harga. Bila `uang muka + jumlah seluruh cicilan = harga` (cicilan 0%), entri tidak memuat baris `credit_cost`. Alasannya: "harga" di model adalah nilai perolehan aset, yaitu yang benar-benar dibayar rumah tangga; selisih negatif berarti potongan harga atau cashback, dan mencatatnya sebagai biaya kredit negatif akan menaikkan kekayaan bersih pada tanggal beli tanpa peristiwa uang. Layar menawarkan mengisi harga sebesar uang muka ditambah jumlah seluruh cicilan; bila rumah tangga menilai barang itu lebih tinggi, mereka mencatat penilaian ulang (`asset_valuation`, §3.16) yang tampil sebagai "perubahan nilai aset", bukan sebagai biaya kredit. Contoh: CTH-49.
- Cicilan yang sudah berjalan sebelum rumah tangga mulai mencatat masuk sebagai saldo awal kewajiban; biaya kreditnya sudah berada di dalam saldo awal dan tidak dicatat lagi.
- Pelunasan dipercepat dengan potongan: `repayment_out` sebesar yang dibayar ditambah `debt_forgiven` sebesar potongannya.
- **Cicilan dan ambang permintaan** (dokumen 06 §6.2, §6.3). Pembelian secara cicilan yang dicatat di aplikasi dibandingkan dengan ambang permintaan **satu kali, saat dibeli**, sebesar uang muka ditambah jumlah seluruh cicilan (§3.8). Pembayaran bulanannya melunasi kewajiban yang sudah tercatat dan terlihat oleh keduanya: ia **tidak pernah** dibandingkan dengan ambang dan **tidak memerlukan kesepakatan rutin**, sehingga cicilan tidak membawa `routine_agreement_ref` (§2.6). Dari sisi uang ini benar: seluruh dana bersama yang diikat kredit itu sudah dihitung satu kali pada saat pembelian, dan menandai tiap pembayaran akan menghitung jumlah yang sama untuk kedua kalinya. Yang dikecualikan hanya bagian yang mengurangi kewajiban: denda keterlambatan di entri yang sama adalah baris `expense`, pengeluaran baru yang belum pernah tercatat sebagai kewajiban, dan dibandingkan dengan ambang seperti pengeluaran lain (§3.8).
- **Kewajiban yang masuk sebagai saldo awal** (dokumen 06 §6.2, putusan 2). Kewajiban dana bersama kepada pihak luar yang didaftarkan sebagai saldo awal saat ada ambang yang berlaku dibandingkan dengan ambang **satu kali, sebesar saldo awalnya** (§3.8); pembayarannya sesudah itu tidak. Dengan itu kredit baru yang dicatat sebagai saldo awal kewajiban, bukan sebagai pembelian secara cicilan, tidak lagi lolos saat dibuka dan saat dibayar. Saldo awal yang diisi saat onboarding, sebelum ambang disepakati, tidak pernah ditandai (**[Asumsi]** contoh: saldo awal hari pertama pada CTH-01, termasuk sisa cicilan motor). Dari sisi uang tidak ada yang terhitung dua kali: saldo awal bukan pengeluaran, tidak masuk anggaran, dan pelunasannya tidak dibandingkan lagi.

Setiap rupiah cicilan terhitung satu kali di tiap tampilan: satu kali di arus kas dan anggaran (saat dibayar), satu kali sebagai pengurang kewajiban, dan bunganya satu kali di perubahan kekayaan bersih (saat kredit dibuka).

### 3.15 Tagihan berulang

Jadwal dan pengingat, bukan entri. Tagihan yang jatuh tempo tetapi belum dibayar **bukan kewajiban** (Q-17): jumlahnya perkiraan, dan kewajiban semu akan mengubah kekayaan bersih tanpa peristiwa uang. Pembayarannya adalah `expense` yang ditautkan ke tagihan itu dengan jumlah sebenarnya. Tagihan boleh diberi **tanggal akhir** (Q-46; D-12); setelah lewat, tidak ada kejadian baru dan pengingat berhenti. **Tagihan tidak terarsip sendiri** (**[Usulan]**, putaran S3; T-36; mencabut "dan tagihan terarsip sendiri" dari revisi sebelumnya, yang bukan bagian D-12): pengarsipan selalu tindakan anggota (§2.6), sehingga tidak ada peristiwa tanpa pelaku atas kesepakatan rutinnya (dokumen 06 §6.3 "Pos rutin" baris (d) tetap: mengarsipkan jadwal adalah tindakan tersendiri). Sesudah tanggal akhir, pembayaran terlambat atas kejadian yang jatuh tempo sampai tanggal itu tetap dapat ditautkan dan dinilai seperti biasa. **Setiap** tagihan berulang, berapa pun jumlahnya, disepakati berdua saat dibuat, lewat satu *kesepakatan rutin* yang dirujuk `routine_agreement_ref` (§2.6; dokumen 06 §6.3 "Pos rutin"); sesudah itu pembayarannya berjalan tanpa diajukan, dan siapa pun dari keduanya boleh membayarnya. Kedudukannya dua, dan tidak dipertukarkan: bahwa tagihan rutin disepakati dulu dan siapa pun boleh membayarnya adalah keputusan pemilik produk, checkpoint S2 (S2-4 pilihan A; kata-katanya: "karena itu rutin untuk membuat tagihan rutin ya disepakati dulu"); bahwa itu berlaku atas **setiap** tagihan berulang berapa pun jumlahnya, dan bahwa jumlah yang disepakati adalah batasnya, adalah rumusan lead yang dikukuhkan pemilik produk, 11 November 2026, bukan keputusannya dengan kata-katanya sendiri (paragraf berikut).

**Jumlah yang disepakati adalah batas** (rumusan lead yang dikukuhkan pemilik produk, 11 November 2026; kata-kata pemilik produk yang mendasarinya: "kalau belanja rutin ya itu kan batas kesepakatan bersama"). Ini mengubah rancangan revisi pasca-checkpoint S2, yang hanya meminta kesepakatan bagi tagihan di atas ambang dan membandingkan kelebihannya dengan ambang.

| # | Aturan | Kedudukan |
|---|---|---|
| 1 | Setiap tagihan berulang disepakati saat dibuat, berapa pun jumlahnya, bukan hanya yang di atas ambang | Dikukuhkan pemilik produk, 11 November 2026 |
| 2 | Jumlah per kali yang disepakati (`R`) adalah batasnya. Untuk tagihan yang jumlahnya berubah-ubah, yang disepakati adalah jumlah tertinggi yang wajar | Dikukuhkan, sda. |
| 3 | Pembayaran yang tertaut padanya dengan *jumlah pemakaian* `P ≤ R` **tidak ditandai**, juga bila `P` di atas ambang, siapa pun yang membayar | Dikukuhkan, sda. |
| 4 | Pembayaran dengan `P > R` **ditandai** sebagai di atas kesepakatan (dasar `routine_amount`, §3.8), **tanpa** membandingkan kelebihannya dengan ambang. Pembayaran di bawah ambang karenanya dapat ditandai | Dikukuhkan, sda. |
| 5 | Selama kesepakatannya belum berlaku (usulan belum dijawab, ditolak, ditarik) atau sesudah diakhiri: aturan biasa, yaitu ambang per entri. Tidak ada yang dihalangi atau ditahan; jadwal dan pengingatnya tetap berjalan | Usulan `architect` (dokumen 06 §6.3 "Pos rutin: yang tidak disebut pemilik produk", INV-A43), diikuti dokumen ini; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-5 (ii) |
| 6 | Rumah tangga dengan satu anggota dewasa: tidak ada kesepakatan yang dapat berlaku dan tidak ada tanda | **[Usulan]** `architect`, sda.; sama dengan ambang dan jatah (D-33) |
| 7 | Penilaiannya per pembayaran, satu kali saat entri disimpan; kekerapan tidak diperiksa; mengganti atau mengakhiri kesepakatan tidak mengubah tanda entri yang sudah tersimpan | **[Usulan]** `architect`, sda. |
| 7a | **Jenis entri yang dapat ditautkan ke tagihan berulang**: yang memuat baris `expense` berpemilik `shared`, yaitu `expense`, `contribution_expense`, dan `advance`. Ketiganya mematikan pengingat periode itu. Yang dinilai terhadap `R`: `expense` dan `advance` (keduanya pemakaian dana bersama; pada `advance` `P` adalah hak ganti yang lahir). `contribution_expense` tertaut tetapi **tidak pernah ditandai**: tagihan itu dibayar dengan uang pribadi tanpa minta diganti, jadi tidak ada dana bersama yang dipakai di atas kesepakatan; ini sejalan dengan kolom ketiga tabel *jumlah pemakaian* (§3.8). Jenis lain, termasuk `repayment_out` dan `owner_transfer`, tidak dapat ditautkan; perintahnya ditolak dan entrinya tetap tersimpan | Jenis yang dapat ditautkan: dokumen 06 §15.4 butir (v) (**[Usulan]** `architect`). Bahwa `contribution_expense` tidak dinilai: **[Usulan]** dokumen ini |
| 8 | Cicilan bukan tagihan berulang: tidak memerlukan kesepakatan rutin dan pembayarannya tidak pernah ditandai (§3.14) | Diterima pemilik produk di checkpoint S2 lewat "lanjut" (D-44); tidak berubah |

Dari sisi uang kesepakatan dan tandanya tidak punya akibat apa pun: pembayaran selalu tersimpan dengan jumlah sebenarnya, dan tidak ada saldo, anggaran, alokasi, *terikat*, atau reservasi yang bergantung padanya (INV-12). `estimated_amount` tetap hanya perkiraan untuk pengingat dan untuk tawaran jumlah alokasi; ia tidak dipakai penilaian dan boleh berbeda dari `R`.

**Fakta rumah tangga pilot** (pernyataan pemilik produk, checkpoint S2). Begitu penghasilannya masuk, ia membayar tagihan lebih dulu, baru membagikan sisanya. Tanggal tagihan tersebar sesudah tanggal uang masuk (contohnya: uang masuk tanggal 11; tagihan tanggal 12, 13, 20, 29), dan jumlah sebagian tagihan dapat jauh berbeda antarbulan. **Kiriman untuk orang tua tidak rutin di rumah tangga pilot**: ia mengikuti aturan biasa (§3.2; di atas ambang diajukan dulu atau bertanda), bukan kesepakatan rutin. Kiriman bulanan yang muncul di contoh dokumen ini (CTH-06, CTH-39) adalah ilustrasi atas angka brief §1.1, bukan gambaran pembayaran rutin rumah tangga pilot; model tetap mendukungnya bagi rumah tangga yang memang mengirim rutin.

Tiga hal yang dituntut pola itu, dan aturan yang sudah menjawabnya. Tidak ada aturan uang baru:

| Kebutuhan | Dijawab oleh | Contoh |
|---|---|---|
| Membayar sebagian tagihan pada hari uang masuk, sebelum jatuh temponya | Pembayaran adalah `expense` bertanggal kejadian hari dibayar, ditautkan ke tagihannya. Tanggal jatuh tempo (`due_rule`) hanya menggerakkan pengingat; membayar lebih awal tidak memerlukan apa pun dan mematikan pengingat periode itu | CTH-05, CTH-41 |
| Menyiapkan uang untuk tagihan yang jatuh temponya menyusul | **Alokasi** (§3.9): label tujuan pada uang di tempat uangnya, tanpa entri uang dan tanpa persetujuan. Pembayaran yang membawa tujuan itu mengurangi alokasinya; *belum dialokasikan* menunjukkan berapa yang masih bebas dibagikan. **Bukan reservasi**: reservasi hanya lahir dari permintaan yang disetujui dan menyebut tempat uang sumber (§3.8), dan tagihan rutin justru tidak diajukan (**[Usulan]**, revisi pasca-checkpoint S2; T-28) | CTH-41; CTH-09 |
| Tagihan yang jumlahnya jauh berbeda antarbulan | `estimated_amount` hanya perkiraan untuk pengingat; yang dicatat selalu jumlah sebenarnya (CTH-05). Jumlah per kali pada kesepakatan rutinnya adalah **batas atas**: pembayaran sampai jumlah itu tidak ditandai, dan yang melampauinya ditandai berapa pun kecilnya kelebihan itu (aturan 3 dan 4 di atas). Rumah tangga karenanya menyepakati jumlah tertinggi yang wajar, bukan rata-ratanya dan bukan perkiraannya; tagihan yang disepakati sebesar perkiraannya akan bertanda setiap kali jatuh sedikit di atasnya (CTH-41, air) | CTH-05, CTH-41 |

Jumlah yang ditawarkan layar saat menyiapkan uang untuk satu tagihan **[Usulan]**, revisi pasca-checkpoint S2: jumlah per kali pada kesepakatan rutinnya bila kesepakatan itu berlaku, selain itu `estimated_amount`; pencatat boleh mengubahnya. "Siapkan uang untuk semua tagihan yang belum dibayar periode ini" sebagai satu langkah adalah syarat layar bagi dokumen 05, bukan aturan uang.

**Kejadian tagihan dan "belum dibayar periode ini"** (usulan putaran S3; PA-FIN-5; diterima lewat "lanjut" di checkpoint S3 (12 November 2026) bersama S3-1 (d), pengingat tagihan per kejadian, dan D-97). Bawaan dokumen 05 ("pembayaran tertaut bertanggal kejadian di periode itu") tidak diambil: tagihan yang dibayar lebih awal pada akhir periode sebelumnya, pola rumah tangga pilot bila uang masuk sebelum hari mulai periode (§3.15 "Fakta rumah tangga pilot"), akan tampil belum dibayar lagi. Penggantinya:

| Hal | Aturan |
|---|---|
| Kejadian | `due_rule` Tier 1 adalah satu hari dalam bulan, 1 sampai 31; pada bulan yang tidak punya hari itu, hari terakhir bulan itu. Tiap bulan sejak tagihan dibuat sampai tanggal akhirnya (bila ada) memberi tepat satu **kejadian**, dikenali dari tanggal jatuh temponya. Karena periode biasa mulai pada hari yang sama tiap bulan (§6.1), satu periode biasa memuat tepat satu kejadian tiap tagihan bulanan; periode peralihan memuat nol atau satu |
| Pembayaran menunjuk kejadiannya | Tautan pembayaran ke tagihan berulang membawa tanggal jatuh tempo kejadian yang dibayarnya (`links.bill_occurrence_due`, atau `entry_link.occurrence_due` bila susulan). Layar mengisinya sendiri menurut urutan bawaan di bawah; pencatat dapat menggantinya dengan kejadian mana pun yang ada. Tidak ada ketukan tambahan |
| Urutan bawaan kejadian (usulan putaran perbaikan S3; PA-FIN-11; T-40; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), D-97) | "Periode berjalan" = periode anggaran yang memuat tanggal kejadian pembayaran. (1) Kejadian paling awal yang belum dibayar dan jatuh tempo paling lambat pada akhir periode berjalan. (2) Bila tidak ada: kejadian paling awal yang belum dibayar sesudah akhir periode itu, paling lambat tanggal akhir tagihan (bila ada); inilah pembayaran lebih awal, pola rumah tangga pilot bila uang masuk sebelum hari mulai periode. (3) Bila tidak ada juga (tagihan sudah lewat tanggal akhirnya dan semua kejadiannya sudah dibayar): kejadian terakhir; pembayaran itu menjadi pembayaran kedua atas kejadian itu, yang sah (baris "Belum dibayar"). Kejadian sesudah periode berjalan sudah ada di model sejak tagihan dibuat (baris "Kejadian"), jadi (2) tidak memerlukan entitas baru. Alasan (2), bukan kosong: tautan ke tagihan berulang selalu membawa satu kejadian (INV-26), dan tanpa isian bawaan pencatat yang membayar lebih awal harus memilih sendiri setiap kali; kejadian berikutnya adalah yang memang ia bayar. Akibatnya: kejadian itu tidak lagi tampil "belum dibayar" di periode berikutnya, dan pengingatnya mati (FR-FIN-030/AC-11). Membayar lebih awal beberapa bulan berturut-turut mengisi kejadian berikutnya satu per satu |
| Belum dibayar | Kejadian yang belum punya pembayaran tertaut yang menunjuknya dan tidak dibalik. Satu pembayaran cukup; model tidak menilai lunas tidaknya, karena jumlah tagihan hanya perkiraan. Dua pembayaran atas kejadian yang sama sah, masing-masing dinilai terhadap `R` (§3.8) |
| "Belum dibayar periode ini" (F-FIN-004) | Semua kejadian yang belum dibayar dari tagihan yang tidak terarsip, yang tanggal jatuh temponya tidak lebih dari akhir periode anggaran berjalan. Kejadian periode lampau yang belum dibayar ikut tampil, sesuai Q-17 ("pengingat sampai dibayar"); kejadian sebelum tagihan itu dibuat tidak ada |
| Sudah disiapkan | Alokasi untuk satu tagihan memakai tujuan berjenis `spending` yang menunjuk tagihan itu (`purpose.recurring_bill_id`), dibuat sekali saat pertama disiapkan. Jumlah yang masih ditawarkan untuk disiapkan = `max(0, jumlah yang ditawarkan − jumlah alokasi tujuan itu di semua tempat uang dana bersama)`. Pembayaran tertaut ke tagihan itu membawa tujuannya dengan sendirinya, sehingga alokasinya berkurang (§3.9; CTH-41) |
| Pengingat | Pengingat tagihan di dokumen 06 (§13.4 di sana) memakai definisi "belum dibayar" ini; jam dan kekerapannya milik dokumen 06 |

Di CTH-41 setiap pembayaran menunjuk kejadian bulan itu (12, 13, 20, dan 29 November); tidak ada angka yang berubah.

**Tagihan berulang tanpa `due_rule`: pos rutin tanpa tanggal** (dokumen 11 §10.1 butir 2, IDEA-03; token listrik, gas, galon, pulsa, infak mingguan; **diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-2 (ii), Tier 1**, dan karena itu **diadopsi** di sini). Ia tetap `recurring_bill`, hanya tanpa `due_rule` (§2.6). Aturan 1 sampai 7a di atas tidak bergantung pada `due_rule`, jadi berlaku tanpa perubahan; yang berbeda hanya apa yang bergantung pada kejadian. Semua baris tabel di bawah mengikuti dari aturan yang sudah ada, kecuali baris "Ada-tidaknya `due_rule`", yang **[Usulan]** dokumen ini (T-41).

| Hal | Aturan |
|---|---|
| Bentuk | `recurring_bill` dengan `due_rule` kosong: nama, pihak lawan, kategori, `estimated_amount`, penanggung jawab, `end_date?`, `routine_agreement_ref`. Saat dibuat, satu angka "jumlah tertinggi yang wajar" sekaligus menjadi `R` dan `estimated_amount` (D-76); `estimated_amount` tidak dipakai penilaian (di atas) |
| Kesepakatan | Sama dengan tagihan bertanggal: disepakati berdua saat dibuat, berapa pun jumlahnya (aturan 1); jumlah per kali `R` adalah batasnya (aturan 2); siapa pun dari keduanya boleh membayar tanpa mengajukan |
| Penilaian per pembayaran | Satu kali, saat entri disimpan (aturan 7). *Jumlah pemakaian* `P` dibandingkan dengan `R`: `P ≤ R` tidak ditandai, juga bila `P` di atas ambang, siapa pun yang membayar (aturan 3); `P > R` ditandai dasar `routine_amount`, tanpa membandingkan kelebihannya dengan ambang (aturan 4). Selama kesepakatannya belum berlaku atau sesudah diakhiri: ambang per entri (aturan 5, S3-5 (ii)). Rumah tangga dengan satu anggota dewasa: tidak ada tanda (aturan 6). Jenis entri yang dapat ditautkan dan `contribution_expense` yang tidak pernah ditandai: aturan 7a. Entri yang juga tertaut ke otorisasi dinilai menurut otorisasinya saja (§3.8) |
| Kekerapan **(batas yang dinyatakan)** | Tidak diperiksa (aturan 7). Tiga pembelian masing-masing `≤ R` dalam satu periode tidak ditandai, walau jumlahnya di atas `R` atau di atas ambang. Jumlah sebulan hanya ditangkap anggaran kategori pos itu (§6.1). Layar kesepakatan wajib menyatakan bahwa batasnya per kali, bukan per bulan (syarat bagi dokumen 05 dan 08) |
| Tautan | Tautan ke pos tanpa `due_rule` **tidak membawa kejadian**: `links.bill_occurrence_due` dan `entry_link.occurrence_due` kosong; tautan yang membawanya ditolak, entrinya tetap tersimpan (INV-26). Tautan susulan hanya dari pencatat entri (S3-5 (iv)). Tautan ke pos terarsip ditolak |
| Kejadian dan pengingat | Tidak ada kejadian, jadi: tidak ada pengingat tagihan (N-32, dokumen 06 §13.4); tidak pernah tampil di "belum dibayar periode ini" (F-FIN-004); tidak termasuk langkah "siapkan uang untuk semua tagihan yang belum dibayar periode ini". Tabel "Kejadian tagihan" di atas, termasuk "Urutan bawaan kejadian" (T-40), hanya berlaku atas tagihan yang punya `due_rule` |
| Alur bagikan | Model tidak menawarkan jumlah apa pun untuk pos ini, karena tidak ada kejadian dan kekerapannya tidak diketahui. Rumah tangga tetap boleh menyiapkan uang untuknya sebagai alokasi biasa (§3.9) dengan tujuan berjenis `spending` yang menunjuk pos itu (`purpose.recurring_bill_id`); pembayaran yang tertaut padanya membawa tujuan itu dan mengurangi alokasinya, seperti tagihan bertanggal (baris "Sudah disiapkan"). "Bagikan seperti terakhir kali" (IDEA-02, S3-2 (iii)) boleh mengisi jumlah alokasi terakhir pencatat sendiri untuk tujuan itu; itu isian layar, bukan aturan uang |
| Ringkasan periode | Pembayarannya `expense` biasa, jadi tampil di tabel A dan B menurut kategorinya (atau di baris "tanpa kategori"), di tabel C menurut pembayar, dan terhitung terhadap anggaran kategorinya (§6.1, §6.3). Tidak ada baris ringkasan baru: §6.3 tidak punya baris per tagihan berulang, bertanggal maupun tidak. Siapa membeli dan kapan terbaca dari daftar entri yang tertaut ke pos itu, sama bagi keduanya |
| Ada-tidaknya `due_rule` (**[Usulan]** dokumen ini; T-41) | Ditetapkan saat pos dibuat dan tidak diganti sesudahnya. Rumah tangga yang ingin memberi tanggal pada pos tanpa tanggal, atau sebaliknya, mengarsipkan pos itu dan membuat pos baru dengan kesepakatan baru. Alasannya: kejadian lahir "sejak tagihan dibuat" (baris "Kejadian" di atas); mengganti ada-tidaknya tanggal di tengah jalan akan memerlukan aturan baru untuk kejadian sebelum dan sesudah penggantian, dan untuk pembayaran lama yang tidak menunjuk kejadian |

Tidak ada invarian, saldo, atau aturan tanda yang berubah; INV-26 hanya dipertegas (tepat satu kejadian bila ada `due_rule`, tidak ada bila tidak). Contoh: CTH-50.

### 3.16 Aset

Q-18.

- Pembelian tunai (`asset_purchase`): uang −X, aset +X. Bukan pengeluaran; kekayaan bersih tetap. Tampil di arus kas sebagai "beli aset". Tidak terhitung terhadap anggaran kecuali pencatat memberinya kategori.
- Nilai diperbarui manual dengan `asset_valuation` bertanggal: aset ±Δ, sebab `valuation` ±Δ. Tidak ada penyusutan otomatis.
- Penjualan atau pelepasan (`asset_disposal`): aset berkurang sebesar nilai tercatat, uang bertambah sebesar hasil, selisihnya `valuation`.
- Mana yang disebut aset dan mana yang belanja adalah pilihan pencatat. Bawaan **[Usulan]**: barang dicatat sebagai aset hanya bila keluarga ingin melihatnya di daftar "aset apa yang kita punya".
- Pembelian aset di atas ambang permintaan diajukan lebih dulu (keputusan pemilik nomor 5). Yang dibandingkan dengan ambang: jumlah yang dibayar bila tunai; uang muka ditambah jumlah seluruh cicilan bila dicicil (§3.8, CTH-34).

### 3.17 Saldo awal

Entri `opening` per posisi, dengan sebab `opening`. Bukan pemasukan, tidak terhitung terhadap anggaran, tidak muncul di arus kas periode mana pun. Tempat uang, piutang, kewajiban, atau aset yang didaftarkan belakangan mendapat `opening`-nya sendiri pada tanggal pendaftaran. Satu-satunya saldo awal yang punya *jumlah pemakaian* adalah kewajiban dana bersama kepada pihak luar (§3.8, §3.14).

**Satu saldo awal per posisi, tidak negatif** (**[Usulan]**, putaran S3; PA-FIN-3; INV-31; mengadopsi bawaan dokumen 05).

| Hal | Aturan |
|---|---|
| Berapa | Setiap tempat uang punya **tepat satu** entri `opening` asli (tanpa `corrects`), dicatat saat tempat uang itu didaftarkan; jumlahnya boleh 0 (tempat uang yang mulai kosong). Piutang, kewajiban, dan aset punya **paling banyak satu**, dan hanya bila posisi itu lahir sebagai saldo awal; yang lahir dari entri lain (`loan_given`, `asset_purchase` secara cicilan, `advance`, dan seterusnya) tidak punya `opening`. Perintah `opening` asli kedua atas posisi yang sama ditolak |
| Tanda | Jumlah saldo awal setiap posisi ≥ 0, karena saldo posisi memang disimpan tidak negatif (§2.3); hanya sebabnya yang bertanda − untuk kewajiban. Saldo awal tempat uang yang negatif ditolak: uang yang sudah berutang dicatat sebagai kewajiban, bukan saldo tempat uang negatif |
| Koreksi | Saldo awal yang salah diperbaiki dengan **entri penyesuaian** berjenis `opening` yang `corrects`-nya menunjuk saldo awal asli, bertanggal kejadian sama, sebesar selisihnya (§3.12); jumlah saldo awal bersih (asli ditambah penyesuaiannya) tetap ≥ 0. **Entri pembalik atas `opening` ditolak**: saldo awal adalah titik mulai posisi itu (INV-22), dan membaliknya meninggalkan posisi tanpa titik mulai; hasil yang sama dicapai dengan penyesuaian ke 0. Koreksi lewat penyesuaian, bukan lewat hitung uang, supaya angka awal yang salah tidak tampil sebagai "selisih hitung" atas nama pemegang |
| Penanda | Saldo awal (dan penyesuaian positifnya) atas tempat uang dana bersama yang dipegang anggota dewasa selain pencatat membawa penanda `unconfirmed` (§3.3; PA-FIN-2). Penyesuaian negatif atasnya adalah entri satu buku dan tidak membawanya, seperti belanja biasa |
| Siapa | Tempat uang pribadi: hanya pemiliknya (INV-25). Tempat uang dana bersama: aturannya milik dokumen 06 (PA-FIN-1) |
| Laporan | Saldo awal dan penyesuaiannya tidak masuk tabel A dan arus kas periode mana pun, juga bila bertanggal di dalam periode. Ringkasan periode yang memuatnya menyebutnya di bawah tabel A dan arus kas sebagai "saldo awal yang didaftarkan periode ini Rp…", sehingga `KB` awal + jumlah tabel A + saldo awal itu = `KB` akhir (INV-19) |

Contoh: CTH-42.

---

## 4. Invarian

Setiap invarian menjadi automated test pada fase pembangunan (`ENGINEERING-STANDARDS.md` §5, §6): property test atas rangkaian perintah acak, ditambah contoh §7 dengan angka yang sama. Semua invarian diperiksa atas **data penuh** (semua buku), bukan atas tampilan seorang anggota.

Notasi: `saldo(p)` = jumlah `delta` semua baris posisi `p`; `U(o)`, `P(o)`, `A(o)`, `K(o)` = jumlah saldo tempat uang, piutang, aset, kewajiban berpemilik `o`; `KB(o) = U(o) + P(o) + A(o) − K(o)`.

| ID | Invarian | Cara menguji |
|---|---|---|
| INV-01 | Setiap `amount`, `delta`, `effect` adalah bilangan bulat dalam satuan terkecil dan membawa mata uang; semua baris satu entri bermata uang sama | Pemeriksaan tipe dan skema; perintah berisi pecahan ditolak |
| INV-02 | **Keseimbangan per pemilik.** Untuk setiap entri `e` dan setiap pemilik `o`: `Σ delta(tempat uang) + Σ delta(piutang) + Σ delta(aset) − Σ delta(kewajiban) = Σ effect(sebab)`, dihitung atas baris `e` yang berpemilik `o` | Untuk setiap entri yang tersimpan |
| INV-03 | **Perpindahan antar-pemilik berjumlah nol.** Dalam satu entri, `Σ effect` semua baris `owner_transfer` di semua buku = 0 | Untuk setiap entri |
| INV-04 | `saldo(p)` yang ditampilkan atau disimpan sebagai cache = jumlah `delta` baris `p` | Hitung ulang dari baris setelah setiap perintah |
| INV-05 | *Uang yang ada* = `U(shared)` = jumlah saldo tercatat semua tempat uang berpemilik `shared`, termasuk yang terarsip dan titipan pada pihak luar | Hitung ulang |
| INV-06 | **Pemindahan tidak mengubah total.** Entri `transfer` tanpa biaya admin: `ΔU(o) = 0` dan `ΔKB(o) = 0`. Dengan biaya admin `f`: `ΔU(o) = ΔKB(o) = −f` | Untuk setiap entri `transfer` |
| INV-07 | **Identitas kekayaan bersih.** Untuk setiap pemilik `o` pada setiap saat: `KB(o) = Σ effect` semua baris sebab berpemilik `o` | Akibat INV-02; diuji langsung |
| INV-08 | **Klaim antar-pemilik dua sisi.** Untuk setiap klaim antara dua pemilik di satu rumah tangga, saldo piutang di buku yang satu = saldo kewajiban di buku yang lain. Akibatnya `Σ KB(o)` atas semua pemilik = uang semua tempat uang + piutang kepada pihak luar + aset − kewajiban kepada pihak luar | Setelah setiap perintah |
| INV-09 | **Hanya menambah.** Entri dan baris yang tersimpan tidak pernah berubah atau hilang: setelah perintah apa pun, himpunan (id, isi inti moneter) entri lama identik dengan sebelumnya. Pengecualian tunggal di §8 hanya menyentuh `entry_detail` dan pemetaan identitas | Bandingkan hash inti moneter sebelum dan sesudah |
| INV-10 | **Idempoten.** Perintah dengan (`household_id`, `idempotency_key`) yang sudah diterima tidak menambah entri, baris, tautan susulan, reservasi, atau otorisasi, dan mengembalikan hasil perintah pertama | Kirim ulang setiap perintah; jumlah baris dan semua saldo tetap |
| INV-11 | `saldo tersedia(h) = saldo tercatat(h) − Σ reservasi aktif(h)`; setiap reservasi ≥ 0; maka saldo tersedia ≤ saldo tercatat | Setelah setiap perintah |
| INV-12 | Persetujuan, penolakan, usul balik, perubahan syarat (`authorization.amended`), otorisasi, reservasi, anggaran, target, tagihan berulang, jadwal cicilan, jadwal penyisihan dan jumlah jatah pribadi, hari mulai periode anggaran, kesepakatan rutin dan rujukannya, konfirmasi custody, keberatan, dan tautan susulan tidak membuat baris posisi atau sebab: semua saldo dan semua `KB(o)` identik sebelum dan sesudahnya | Untuk setiap perintah jenis itu |
| INV-13 | **Batas tautan.** Saldo piutang dan kewajiban ≥ 0 (pelunasan, penggantian, penghapusan, pelepasan tidak melebihi sisa); nilai aset ≥ 0; jumlah pengembalian dana atas satu entri ≤ jumlah entri itu | Perintah yang melanggar ditolak |
| INV-14 | `realisasi(a)` = jumlah *jumlah pemakaian* (§3.8) entri yang tertaut ke otorisasi `a`, lewat tautan pada entri atau tautan susulan yang berlaku, dikurangi pengembalian dana tertaut; berlaku apa pun status `a`. `terikat(a) = max(0, batas(a) − realisasi(a))` bila `a` terbuka, 0 bila tidak; `batas(a)` adalah batas versi syarat yang berlaku | Hitung ulang, termasuk setelah `authorization.amended` dan setelah tautan ke otorisasi tertutup |
| INV-15 | **Anggaran.** Satu baris terhitung terhadap paling banyak satu anggaran (yang kategorinya sama dan periodenya memuat tanggal kejadiannya). `sisa = batas − terpakai − terikat` (§6.1) | Hitung ulang; jumlah `terpakai` semua kategori = jumlah baris berkategori pada periode itu |
| INV-16 | **Rekonsiliasi.** Sesaat setelah entri `cash_difference`, atau entri `expense` yang merinci seluruh selisih saat dihitung (§3.10), saldo tercatat tempat uang itu atas entri bertanggal kejadian ≤ tanggal hitung = saldo terhitung. Selisih dihitung server terhadap baris yang sudah ia terima saat perintah diterima, tidak pernah dikirim perangkat (§3.10 "Hitungan yang tiba belakangan"). Untuk tiap rekonsiliasi, `0 ≤ jumlah dirinci ≤ |selisih|` | Setelah setiap rekonsiliasi dan `recategorization`; kirim serah terima ke tempat uang itu dari perangkat lain lebih dulu, lalu hitungan bersaldo terhitung sama dengan saldo server: tidak ada entri `cash_difference` (CTH-46); perincian sekaligus yang jumlahnya tidak sama dengan selisih server tersimpan sebagai rekonsiliasi biasa |
| INV-17 | **Alokasi.** Alokasi per (tempat uang, tujuan) ≥ 0. `belum dialokasikan(h) = saldo tercatat(h) − Σ alokasi(h)`. Perintah `allocation` yang membuat `Σ alokasi(h) > saldo tercatat(h)` ditolak | Setelah setiap perintah |
| INV-18 | **Koreksi.** Entri pembalik memuat baris yang sama dengan entri asal dengan tanda berlawanan; jumlah akibat keduanya pada setiap saldo dan setiap `KB(o)` nol. Satu entri punya paling banyak satu pembalik | Untuk setiap pembalikan |
| INV-19 | **Jembatan laporan.** Untuk setiap pemilik dan setiap rentang tanggal kejadian: `ΔKB = Σ effect` baris sebab dalam rentang; `ΔU = Σ delta` baris tempat uang dalam rentang. Jumlah baris arus kas (§6.4) = `ΔU` dikurangi baris tempat uang entri `opening` dalam rentang; jumlah baris ringkasan (§6.3) = `ΔKB` dikurangi baris sebab `opening` dalam rentang; selisih keduanya adalah "saldo awal yang didaftarkan periode ini" (§3.17) | Hitung dua cara, juga atas rentang yang memuat saldo awal tempat uang yang didaftarkan belakangan |
| INV-20 | **Buku rumah tangga bebas dari dana pribadi.** Tidak ada angka rumah tangga (uang yang ada, kekayaan bersih, anggaran, laporan) yang nilainya berubah ketika baris berpemilik `member(M)` ditambah, kecuali lewat baris berpemilik `shared` dalam entri yang sama. Angka rumah tangga identik bagi kedua pelihat | Tambahkan entri buku pribadi murni; angka rumah tangga tetap. Hitung angka rumah tangga sebagai tiap pelihat; hasilnya sama |
| INV-21 | **Arsip.** Posisi terarsip bersaldo nol, tempat uang terarsip beralokasi nol, dan keduanya tidak menerima baris posisi atau baris alokasi baru; `recategorization` atas rekonsiliasi tempat uang terarsip diterima, karena tidak punya baris posisi. Tagihan berulang dan tujuan yang terarsip tidak menerima tautan atau alokasi baru. Pengaktifan kembali (§2.6) adalah peristiwa tanpa baris; sesudahnya butir itu menerima baris seperti biasa | Perintah yang melanggar ditolak; perincian selisih atas tempat uang terarsip diterima dan tidak mengubah saldo (CTH-43); belanja atas tempat uang yang sama ditolak sebelum diaktifkan kembali dan diterima sesudahnya |
| INV-22 | `event_date ≤` tanggal `recorded_at`; `event_date ≥` tanggal `opening` setiap posisi yang disentuh | Perintah yang melanggar ditolak |
| INV-23 | **Bersamaan.** Hasil dua perintah sah yang diterima bersamaan sama dengan hasil menjalankannya berurutan dalam salah satu urutan; tidak ada baris yang hilang | Uji konkurensi (CTH-31) |
| INV-24 | **Penghasilan yang dicatat sebagai penghasilan selalu tercatat penuh di buku rumah tangga.** *Daftar sumber penghasilan* adalah sumber rumah tangga yang berkelompok `work` (§3.1 "Sumber uang masuk"; sebelum pernyataan pemilik produk 12 November 2026, parameter preset berisi `salary`, `business`). Setiap baris sebab `income` yang sumbernya berkelompok `work` saat entri diterima berpemilik `shared`; perintah yang menulisnya di buku pribadi ditolak. Dalam entri `income` **yang memuat baris sebab `income` bersumber penghasilan**, tempat uang berpemilik `member(M)` hanya bertambah lewat sepasang baris `owner_transfer` di entri yang sama (−X di buku rumah tangga, +X di buku M), dengan 0 ≤ X ≤ jumlah pemasukan. Entri `income` yang sumbernya di luar daftar itu tidak dibatasi kalimat ini: hadiah untuk satu orang, atau pemasukan lain yang memang uang pribadi, dicatat sebagai `income` di buku pribadi penerimanya saja (§3.1, CTH-36). Akibatnya jumlah penuh setiap penghasilan dan bagian yang diarahkan ke dana pribadi terbaca dari buku rumah tangga saja, sama bagi kedua pelihat | Bangkitkan perintah uang masuk acak dari kedua anggota: tidak ada baris `income` bersumber penghasilan di buku pribadi; untuk setiap entri `income` bersumber penghasilan, kenaikan tempat uang pribadi = −Σ `owner_transfer` di buku rumah tangga; baca buku rumah tangga sebagai tiap pelihat, hasilnya sama (CTH-03b). Entri `income` bersumber di luar daftar yang seluruh barisnya di buku pribadi pencatat diterima (CTH-36). Ulangi dengan daftar sumber lain (misalnya THR dipindah ke `gift_other` lewat kesepakatan, atau sumber buatan sendiri ditambahkan ke tiap kelompok): aturan yang sama berlaku atas daftar itu, dan entri yang diterima sebelum perubahan tidak berubah. Perintah yang menambah sumber `work`, mengganti namanya, mengarsipkannya, atau memindahkan sumber antarkelompok tanpa kesepakatan yang berlaku ditolak; ganti nama "THR" menjadi "THR kantor" lewat kesepakatan, lalu sumber `gift_other` baru atau yang diganti namanya menjadi "thr ": ditolak; pemindahan atau pengarsipan "hadiah" dan "lainnya" ditolak; nama yang sama dengan sumber lain yang tidak terarsip ditolak; semua perubahan sumber tidak mengubah saldo atau `KB(o)` mana pun (§3.1 "Sumber uang masuk") |
| INV-25 | **Tulis buku pribadi.** Entri yang memuat baris berpemilik `member(M)` hanya diterima bila `actor = M`. Dua pengecualian, keduanya tanpa baris tempat uang atau aset milik M: (1) `claim_waiver` atas kewajiban M kepada dana bersama, dicatat anggota dewasa selain M, yang di buku M hanya memuat baris kewajiban − dan `owner_transfer` + (§2.2); (2) `exit_closure` oleh sistem saat M keluar (INV-28) | Perintah dari anggota lain yang menyentuh buku M ditolak tanpa akibat apa pun; tidak ada entri tersimpan yang memuat baris tempat uang atau aset milik M dengan `actor ≠ M` |
| INV-26 | **Tautan.** Menambah atau mencabut tautan susulan tidak mengubah entri, baris, atau saldo apa pun (INV-09). Satu entri tertaut ke paling banyak satu otorisasi yang berlaku. Tautan ke otorisasi hanya sah bila, **saat tautan itu dibuat**, otorisasi itu `open` (tanpa syarat tanggal: entri yang tanggal kejadiannya lebih awal dari tanggal persetujuan boleh ditautkan; *konfirmasi susulan*, dokumen 06 INV-A40), atau `settled`/`expired` dengan tanggal kejadian entri di antara tanggal persetujuan dan tanggal penutupannya; tidak pernah ke yang `cancelled`. Syarat itu tidak diperiksa ulang sesudahnya: tautan yang sah saat dibuat tetap sah ketika otorisasinya ditutup. Tautan susulan ke tagihan berulang hanya diterima dari pencatat entri itu. Otorisasi yang lahir dari *konfirmasi susulan* disetujui, ditautkan, dan ditutup dalam satu transaksi, dengan batas sebesar *jumlah pemakaian* entri yang dikonfirmasi, dan memuat tepat satu tautan, yang pelakunya pemohon dan pemohon itu pencatat entri yang dikonfirmasi: perintah tautan lain kepadanya, dan permintaan konfirmasi susulan dari yang bukan pencatat entri itu, ditolak (**[Usulan]** `security-privacy` dan dokumen ini; §3.8). Entri `repayment_out` atas suatu kewajiban tidak dapat ditautkan ke otorisasi yang sudah ditaut penuh oleh entri yang melahirkan kewajiban itu: `asset_purchase` secara cicilan, atau `opening` kewajiban dana bersama kepada pihak luar (§3.8; dokumen 06 §4.5 "Masa berlaku dan kedaluwarsa"). Entri `write_off` atas suatu piutang tidak dapat ditautkan ke otorisasi yang memuat `loan_given` asal piutang itu (§3.7; **[Usulan]**, putaran S3). Tautan ke tagihan berulang yang terarsip ditolak; tautan ke tagihan berulang membawa tepat satu kejadian bila tagihan itu punya `due_rule`, dan tidak membawa kejadian bila tidak (§3.15; IDEA-03) | Hash inti moneter sebelum dan sesudah; perintah tautan yang melanggar ditolak; entri itu sendiri tetap tersimpan; tautkan entri bertanggal sebelum persetujuan ke otorisasi `open`, lalu tutup otorisasi itu: tautan tetap berlaku dan realisasinya tidak berubah; perintah yang sama sesudah otorisasi ditutup ditolak; tautan susulan ke tagihan berulang dari bukan pencatat ditolak; konfirmasi susulan atas entri berjumlah pemakaian `P`: sesudah transaksi persetujuan, otorisasinya `settled`, batas = realisasi = `P`, *terikat* 0, tidak ada reservasi aktif, tidak ada `authorization.overrun`, semua saldo identik; tautan kedua ke otorisasi itu ditolak; tautan ke tagihan ber-`due_rule` tanpa kejadian, dan tautan ke pos tanpa `due_rule` yang membawa kejadian, ditolak sedangkan entrinya tetap tersimpan (CTH-50) |
| INV-27 | **Bagian pribadi dari penghasilan hanya bagi penerimanya.** Keputusan pemilik produk, checkpoint S2 (S2-1 butir (i); §3.1). Baris `income` bersumber penghasilan (INV-24) membawa `earner` yang bernilai tepat satu anggota dewasa rumah tangga itu. Entri `income` bersumber penghasilan hanya boleh memuat baris `owner_transfer` (bagian pribadi X > 0) bila `actor = earner`; perintah yang melanggar ditolak seluruhnya, dan perintah yang sama tanpa bagian pribadi diterima | Bangkitkan perintah uang masuk bersumber penghasilan dari kedua anggota, dengan dan tanpa bagian pribadi, dengan `earner` diri sendiri dan pasangan: yang berbagian pribadi dengan `actor ≠ earner` ditolak tanpa akibat apa pun; tiga kombinasi lain diterima. Tidak ada entri tersimpan bersumber penghasilan yang memuat `owner_transfer` dengan `actor ≠ earner` (CTH-03b: keduanya suami) |
| INV-28 | **Penutupan klaim saat anggota keluar.** Keputusan sementara pemilik produk, checkpoint S2 (S2-8 pilihan A, butir (d)), sampai advokat menjawab (§8.2). Entri `exit_closure` hanya ditulis sistem (`actor = system`), tepat satu kali per anggota yang keluar, dalam transaksi yang sama dengan penutupan buku pribadinya, **pada akhir masa tenggang keluar** (§8.2 butir 2b; usulan `security-privacy` SP-48, putaran perbaikan S3; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-4 (c); 7 hari **[Asumsi]**). Selama masa tenggang buku pribadi anggota itu tidak menerima baris baru dari siapa pun, termasuk `claim_waiver` oleh anggota lain, sehingga saldo setiap klaim pada akhir masa tenggang sama dengan saat akses berakhir; bila keluar dibatalkan di dalam masa itu, tidak ada `exit_closure` dan buku itu menerima baris lagi seperti biasa. *Penutupan* berarti buku itu tidak lagi menerima baris baru, dan terjadi pada setiap keluarnya anggota, dengan atau tanpa permintaan penghapusan; *penghapusan* baris buku itu terjadi atas permintaan atau pada akhir jangka retensi (§8.2 butir 2a). Ia hanya memuat baris piutang, kewajiban, dan `owner_transfer`; tidak pernah baris tempat uang atau aset. Sesudahnya saldo setiap klaim antar-pemilik yang menyangkut buku pribadi anggota itu nol di kedua sisinya. Ia memenuhi INV-02 dan INV-03, tidak mengubah `U(o)` pemilik mana pun, dan tidak dapat dibalik | Bangkitkan klaim acak antara buku pribadi M dan buku lain, lalu keluarkan M: semua klaim itu bersaldo nol di kedua sisi; `ΔU` semua pemilik nol; `ΔKB(shared)` = Σ kewajiban rumah tangga kepada M − Σ piutang rumah tangga kepada M yang ditutup; perintah `exit_closure` dari anggota, perintah kedua untuk anggota yang sama, dan perintah pembaliknya ditolak (CTH-38). Masa tenggang: sesudah akses berakhir dan sebelum masa itu lewat, setiap perintah yang menulis buku M ditolak tanpa akibat apa pun, dan belum ada `exit_closure`; pada akhir masa itu `ΔKB(shared)` sama dengan bila ditutup saat akses berakhir; keluar yang dibatalkan di dalam masa itu tidak menghasilkan `exit_closure` dan semua saldo identik dengan sebelum keluar diminta |
| INV-29 | **Periode anggaran membagi habis garis tanggal** (§6.1). Hari mulai adalah bilangan bulat 1 sampai 28. Untuk setiap rumah tangga, setiap tanggal berada di tepat satu periode; periode berurutan tanpa celah dan tanpa tumpang-tindih. Periode biasa panjangnya 28 sampai 31 hari; periode peralihan 1 sampai 30 hari. Mengganti hari mulai tidak mengubah batas periode mana pun yang sudah dimulai pada tanggal perubahan itu diterima, dan tidak mengubah saldo atau `KB(o)` mana pun (INV-12) | Rangkaian acak penggantian hari mulai pada tanggal acak: setiap tanggal dalam rentang uji jatuh di tepat satu periode; batas periode yang sudah dimulai identik sebelum dan sesudah tiap penggantian; jumlah `terpakai` semua periode per kategori = jumlah baris berkategori itu (INV-15); hari mulai di luar 1 sampai 28 ditolak |
| INV-30 | **Jatah pribadi hanya membagi, tidak mengubah** (§3.4). Untuk setiap entri `e` dan anggota M: `0 ≤ dalam_jatah(e) ≤ bagian(e, M)` dan `dalam_jatah(e) + luar_jatah(e) = bagian(e, M)`. Dalam satu periode, jumlah `dalam_jatah` atas entri yang tidak dibalik ≤ jumlah jatah terbesar yang berlaku dalam periode itu. Hitungan satu periode tidak bergantung pada entri periode lain. Jatah, perubahannya, dan hitungannya tidak membuat baris dan tidak mengubah saldo, anggaran, *terikat*, reservasi, atau `KB(o)`. **Tanda** (keputusan pemilik produk, checkpoint S2): bagi entri yang terhitung dan anggota yang punya jatah berlaku, `ditandai(e)` tepat bila `luar_jatah(e) > 0`, yaitu `terpakai_sebelum + bagian > jatah`, tanpa bergantung pada ambang; tanpa jatah berlaku, tepat bila `bagian > ambang`. **Jatah yang diakhiri pemiliknya** (rumusan lead yang dikukuhkan pemilik produk, 11 November 2026; T-29 pilihan 1) dihitung sebagai jatah berlaku, dengan jumlahnya saat diakhiri, bagi setiap entri yang waktu simpannya di server tidak melewati akhir periode anggaran yang memuat pengakhiran itu, kecuali jatah baru sudah disepakati; entri yang disimpan sesudahnya dinilai tanpa jatah. Itu hanya berlaku bila jatah yang diakhiri **di atas ambang yang berlaku saat diakhiri**; jatah yang tidak di atas ambang, termasuk 0, berakhir seketika juga atas penilaian (**[Usulan]** `security-privacy`, SP-33). Jatah berjumlah 0 adalah jatah berlaku. Jatah yang tidak di atas ambang yang berlaku (≤ ambang, termasuk 0) hanya dapat diusulkan pemiliknya sendiri; pagarnya guard K01 dokumen 06, dikutip di sini (**[Usulan]** `security-privacy`, SP-32, SP-34). Entri yang tertaut ke otorisasi **tidak dinilai terhadap jatah**: bagiannya yang berada dalam batas otorisasi tidak ditandai, dan bagian yang membawa realisasi ke atas batas dibandingkan dengan ambang dengan dasar `threshold` (§3.8 "Tanda atas entri"). Tidak ada entri yang ditolak karena jatah | Bangkitkan penyisihan dan pemasukan berbagian pribadi acak dalam beberapa periode, dengan jatah yang berubah: ketiga hubungan itu berlaku; untuk setiap entri, tanda yang dihasilkan sama dengan `terpakai_sebelum + bagian > jatah` bila ada jatah dan dengan `bagian > ambang` bila tidak, dan mengubah ambang tidak mengubah tanda entri anggota yang punya jatah; urutan 400000, 400000, 900000 atas jatah 1000000 memberi tepat satu tanda dan urutan 900000, 400000, 400000 tepat dua (§10.3 butir 7); dengan ambang 500000, akhiri jatah 1000000 sesudah 1000000 terpakai, lalu simpan penyisihan 100000 dalam periode yang sama: bertanda (1000000 + 100000 > 1000000); simpan penyisihan 100000 pada hari pertama periode berikutnya: tidak bertanda (100000 tidak di atas 500000); akhiri jatah 500000 (tidak di atas ambang 500000) sesudah 500000 terpakai, lalu simpan penyisihan 100000 dalam periode yang sama: tidak bertanda, karena pengakhirannya berlaku seketika; hal yang sama atas jatah 300000 dan jatah 0; lima penyisihan 400000 dalam satu periode: 0 tanda tanpa jatah, 3 tanda dengan jatah 1000000, 5 tanda dengan jatah 300000 (§10.3 butir 11); penyisihan 800000 yang ditautkan ke otorisasi berbatas 200000 tanpa realisasi lain: bertanda dengan dasar `threshold` (800000 − 200000 = 600000 > 500000) dan tidak mengubah `terpakai_sebelum` entri mana pun; penyisihan 230000 yang ditautkan ke otorisasi lain berbatas 200000 tanpa realisasi lain: tidak bertanda (230000 − 200000 = 30000, tidak di atas 500000); entri yang disimpan sesudah akhir periode itu tetapi bertanggal kejadian di dalamnya: dinilai tanpa jatah; jatah 0 yang berlaku: setiap penyisihan bertanda; semua perintah diterima; menghapus semua entri periode lain dari masukan tidak mengubah hasil suatu periode; semua saldo dan `KB(o)` identik dengan dan tanpa jatah (CTH-39) |
| INV-31 | **Saldo awal** (**[Usulan]**, putaran S3; §3.17). Setiap tempat uang punya tepat satu entri `opening` tanpa `corrects`; setiap piutang, kewajiban, dan aset paling banyak satu. Jumlah saldo awal bersih tiap posisi (entri asli ditambah penyesuaiannya) ≥ 0. Entri pembalik atas `opening` tidak ada. Baris `opening` yang menambah tempat uang berpemegang anggota dewasa selain `actor` membawa penanda `unconfirmed` (§2.5 baris (a)) | Perintah `opening` asli kedua, saldo awal negatif, penyesuaian yang membuat saldo awal bersih negatif, dan pembalik atas `opening` ditolak tanpa akibat apa pun; penyesuaian −20000 atas saldo awal 200000 diterima dan saldo tempat uang 180000 (CTH-42); saldo awal yang dicatat bukan pemegangnya bertanda, yang dicatat pemegangnya tidak |
| INV-32 | **Bentuk tempat uang** (**[Usulan]**, putaran S3; §2.3). `owner = member(M)` ⇒ `holder = member(M)` dan `kind ≠ entrusted_external`; `kind = entrusted_external` ⇔ `holder = external_party(P)`, dan keduanya hanya dengan `owner = shared`; tempat uang berpemilik `shared` tidak `private` (§5.2) | Perintah membuat atau mengubah tempat uang yang melanggar ditolak; tidak ada tempat uang tersimpan yang melanggar |
| INV-33 | **Hitungan bukan pemegang** (usulan `security-privacy` SP-50, putaran perbaikan S3; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-9 (i); §3.10 butir 1). Rekonsiliasi atas tempat uang yang pemegangnya anggota dewasa selain `actor` hanya diterima bila tempat uang itu berpemilik `shared`, berjenis `bank` atau `ewallet`, bertanda `others_may_count = true` saat perintah diterima, pemegangnya anggota aktif, dan tidak ada hitungan bukan pemegang yang masih terbuka atasnya. Akibatnya, pada setiap saat, setiap tempat uang punya paling banyak satu hitungan bukan pemegang yang terbuka. `others_may_count` hanya diubah pemegangnya dan tidak mengubah saldo apa pun (INV-12) | Hitungan bukan pemegang atas tempat uang tanpa tanda, atas uang tunai, dan hitungan kedua (juga yang cocok) selagi yang pertama terbuka: ditolak tanpa akibat apa pun. Sesudah pemegang mengonfirmasi, berkeberatan, atau mencatat hitungannya sendiri, hitungan bukan pemegang berikutnya diterima. Perintah mengubah tanda dari bukan pemegang ditolak. Mematikan tanda tidak mengubah rekonsiliasi yang sudah ada (CTH-47) |
| INV-34 | **Penghapusan piutang dinilai kumulatif** (usulan `security-privacy` SP-51, putaran perbaikan S3; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-9 (ii); §3.7, §3.8). Untuk entri `write_off` `e` atas piutang dana bersama `r`: `H(e)` = jumlah `N` entri `write_off` lain atas `r` yang tersimpan lebih dulu dan belum dibalik saat `e` disimpan; `e` bertanda dasar `threshold` tepat bila ada ambang yang berlaku, rumah tangga punya ≥ dua anggota dewasa, `N(e) > 0`, dan `H(e) + N(e) > ambang`. Tanda tidak menahan apa pun; semua saldo dan `KB(o)` identik dengan dan tanpa aturan ini | Rangkaian penghapusan acak atas beberapa piutang, sebagian tertaut ke otorisasi, sebagian dibalik: tanda yang dihasilkan sama dengan rumus di atas; dengan ambang 500000, 300000 lalu 300000 atas piutang yang sama memberi tanda pada yang kedua saja, dan 50000 sesudahnya bertanda; 300000 atas piutang lain tidak bertanda; bagian yang tertaut ke otorisasi dalam batasnya tidak bertanda dan tidak menambah `H` (CTH-25 varian SP-51) |

Yang sengaja **bukan** invarian: saldo tercatat tempat uang ≥ 0 (§2.5), *belum dialokasikan* ≥ 0 (§3.9), dan saldo tersedia ≥ 0 (§3.8). Ketiganya adalah sinyal bagi pengguna, bukan penolakan.

---

## 5. Tampilan custody

Kriteria sukses nomor 1: "sisa berapa dan siapa yang pegang", tanpa bertanya.

### 5.1 Perhitungan

```text
custody rumah tangga =
  untuk setiap tempat uang h berpemilik shared, tidak terarsip atau bersaldo bukan nol:
    (pemegang(h), jenis(h), nama(h), saldo tercatat(h), saldo tersedia(h),
     jumlah baris masuk yang belum dikonfirmasi,
     jumlah baris keluar yang belum dikonfirmasi,                            -- §2.5 baris (b) dan (c)
     jumlah yang dikeberatani,                                               -- §3.3; dokumen 06 §7
     alokasi per tujuan, belum dialokasikan)
  dikelompokkan per pemegang: anggota M | dipegang bersama | pihak luar P (titipan)

angka utama = uang yang ada = Σ saldo tercatat(h)        -- INV-05
```

Jumlah semua baris custody selalu sama dengan angka utama. Tidak ada baris "lainnya".

### 5.2 Yang dilihat tiap anggota dewasa

| Bagian | Isi | Pelihat A | Pelihat B |
|---|---|---|---|
| Angka utama | *Uang yang ada* | Sama | Sama |
| Rincian per pemegang | Tiap tempat uang dana bersama: pemegang, tempat, saldo, yang belum dikonfirmasi, yang "sudah dijanjikan" (reservasi) | Sama | Sama |
| Dipinjam orang, utang kita, sisa cicilan, aset, harta bersih | Dari buku rumah tangga, termasuk hak ganti yang belum dibayar kepada A atau B. Klaim internal yang sedang dikeberatani diberi keterangan "dikeberatani" beserta jumlah yang dikeberatani; angkanya tetap terhitung penuh (§3.12) | Sama | Sama |
| Baris kedua: "uang pribadi saya" | Tempat uang berpemilik pelihat itu sendiri, beserta klaimnya | Hanya milik A | Hanya milik B |

Buku rumah tangga identik bagi keduanya (INV-20). Baris kedua tidak dijumlahkan ke angka utama. Visibilitas per tempat uang (`private | spouses | all_members`) ditetapkan dokumen 06 dan 10; model hanya menuntut bahwa tempat uang berpemilik `shared` tidak dapat `private`, sebab angka utama harus sama bagi keduanya. Bila seorang anggota memilih memperlihatkan tempat uang pribadinya kepada pasangan, ia tampil di layar pasangan sebagai bagian terpisah dan tetap tidak masuk angka rumah tangga.

### 5.3 Kebocoran yang tersisa

Dinyatakan terbuka. Anggota diberi tahu sebelum mencatat hal yang akan terlihat pasangannya (Q-24).

| # | Yang diketahui pasangan | Sebab | Dapat dihindari? |
|---|---|---|---|
| B1 | Jumlah dan tanggal setiap penyisihan dana bersama menjadi dana pribadi. Termasuk bagian penghasilan yang diarahkan langsung ke dana pribadi saat dicatat: pasangan melihat jumlah penuh penghasilannya dan bagian pribadinya | Entri buku rumah tangga; keputusan pemilik produk, 11 November 2026; INV-24 | Tidak, dan memang tidak boleh: ini pagar terhadap pengalihan sepihak |
| B2 | Bahwa anggota punya dana pribadi sekurangnya sebesar kontribusi atau talangannya, pada tanggal itu | Keputusan pemilik nomor 3: kontribusi terlihat beserta sumbernya | Hanya dengan tidak mencatatnya sebagai keperluan rumah tangga |
| B3 | Hak ganti dan utang antar-pasangan: jumlah, arah, dan pelunasannya | Klaim punya dua sisi | Tidak |
| B4 | Batas bawah kasar dana pribadi: jumlah penyisihan yang diterima dikurangi kontribusi yang terlihat | Akibat B1 dan B2 | Tidak sepenuhnya; pemasukan dan belanja pribadi yang tidak terlihat membuatnya bukan saldo sebenarnya. Aplikasi **tidak menghitungkannya**: tidak ada angka kumulatif, lintas periode, atau bersih per anggota atas penyisihan dan kontribusi (§6.3). Pasangan yang menjumlahkan sendiri baris per periode tetap dapat memperolehnya |
| B5 | Hadiah untuk satu orang dan pemasukan lain yang dicatat langsung di buku pribadi **tidak** terlihat sama sekali: untuk hadiah, keputusan pemilik produk di checkpoint S2 (S2-1 butir (ii); §3.1); untuk pemasukan "lainnya", default yang diterima lewat "lanjut" (D-47). Penghasilan tidak dapat dicatat demikian (INV-24); tetapi bila gaji suatu bulan tidak dicatat, pasangan melihat gajinya tidak muncul | Keputusan pemilik nomor 2 dan 3 menjadikannya dana pribadi yang tertutup; bahwa catatan hadiah ikut tertutup kini keputusannya sendiri (checkpoint S2). Ketiadaan entri juga informasi | Tidak |
| B6 | Saldo bagian pribadi di dompet campuran, bagi siapa pun yang melihat isi dompet itu atau saldo rekening atau dompet digital campuran itu sendiri (di tangan, di layar ATM, di aplikasi bank): `bagian pribadi = isi yang terlihat − saldo bagian bersama`, dan saldo bagian bersama terlihat pasangan (temuan `security-privacy` SP-40 (a)) | Dana bersama terlihat keduanya (INV-20); benda atau rekeningnya ada di luar aplikasi. Aturan bahwa `C` tidak terbaca pasangan (§2.3) hanya menutup jalan lewat aplikasi | Hanya dengan menyimpan uang pribadi di tempat terpisah. Model tidak dapat menutupnya |

Yang tetap tertutup: saldo tempat uang pribadi, jumlah dan nama tempat uang pribadi, pemasukan pribadi selain penghasilan, belanja pribadi, dan aset pribadi. Dari penghasilan yang diarahkan ke dana pribadi, yang tertutup adalah tempat uang pribadi tujuannya, saldonya, dan pemakaiannya sesudah itu.

Risiko yang berlawanan arah, untuk `security-privacy`: jaminan INV-24 berlaku atas pemasukan yang **diberi sumber** berkelompok `work` (bawaannya gaji, THR, bonus, lembur, honor atau pekerjaan lepas, hasil usaha, ditambah sumber `work` yang disepakati rumah tangga; §3.1). Sumber diisi pencatat; penghasilan yang dicatat sebagai hadiah atau "lainnya" di buku pribadi, atau tidak dicatat sama sekali, tetap tidak terlihat. Lihat §10.3 butir 1.

---

## 6. Anggaran dan laporan

Semua angka di bagian ini dihitung dari buku rumah tangga, menurut **tanggal kejadian**, dan sama bagi kedua pasangan.

### 6.1 Anggaran

Keputusan pemilik nomor 8: anggaran adalah batas atas satu kategori buatan rumah tangga, jumlahnya dapat diubah, periodenya bulanan, tanpa amplop. Periode disimpan sebagai rentang tanggal [`period_start`, `period_end`], sehingga periode mingguan atau harian (Tier 3) tidak mengubah struktur.

**Periode anggaran adalah setelan rumah tangga** (keputusan pemilik produk, 11 November 2026: "bisa di tentukan pasti ada yang fixed tanggal ada juga yang gak tertentu"; menyempurnakan keputusan S1 "bulanan mulai tanggal gajian"). Menurut pemilik produk, gajinya "bisa 2x" dan uang masuk dari pekerjaan lepas "bisa sampe 5x"; bahwa itu berarti sampai sekitar tujuh uang masuk sebulan adalah **bacaan lead**, bukan kata-katanya (§3.1). Pada bacaan mana pun tidak ada satu "tanggal gajian" yang dapat dijadikan patokan; rumah tangga memilih sendiri hari mulainya. Asasnya keputusan pemilik produk; rentang hari, bentuk peralihan, dan batas anggaran pada periode peralihan adalah **[Usulan]** dokumen ini (T-23, T-24).

| Hal | Aturan |
|---|---|
| Setelan | Satu **hari mulai** per rumah tangga, bilangan bulat 1 sampai 28; bawaan 1. Berlaku atas semua kategori; tidak ada hari mulai per kategori |
| Mengapa 1 sampai 28 | Setiap bulan punya tanggal 28, jadi periode selalu mulai pada tanggal yang sama tanpa aturan tambahan untuk Februari dan bulan 30 hari. Rumah tangga yang penghasilannya tiba pada tanggal 29 sampai 31 memilih 28 atau 1. Biayanya: "akhir bulan" tidak dapat dipilih persis |
| Periode biasa | Dari hari mulai suatu bulan sampai sehari sebelum hari mulai bulan berikutnya; panjangnya 28 sampai 31 hari |
| Mengganti hari mulai | Berlaku **mulai periode berikutnya**. Periode yang sedang berjalan saat penggantian diterima, dan semua periode sebelumnya, tidak berubah batasnya. Penggantian tercatat sebagai baris baru `budget_period_setting` (§2.6) dan sebagai *peristiwa* beserta pelakunya; tidak membuat entri (INV-12) |
| Periode peralihan | **Satu periode pendek.** Ia mulai sehari sesudah periode berjalan berakhir dan berakhir sehari sebelum hari mulai baru yang pertama sesudahnya; panjangnya 1 sampai 30 hari. Sesudah itu periode biasa menurut hari mulai baru |
| Mengapa pendek, bukan satu periode panjang | Periode panjang dapat mencapai hampir dua bulan dan memuat dua kali cicilan, dua kali tagihan, dan dua kali kiriman bulanan terhadap batas satu bulan: anggaran pasti tampil terlampaui padahal tidak ada yang salah. Periode pendek paling banyak memuat satu kali tiap pembayaran bulanan; akibat terburuknya sisa anggaran yang besar, yang tidak berpindah (tanpa rollover) |
| Anggaran pada periode peralihan | Baris `budget` dibuat dari `budget_rule` dengan batas penuh, tidak diprorata: prorata menghasilkan pecahan dan menebak pola belanja. Rumah tangga boleh mengubah batas periode itu seperti batas periode lain |
| *Terikat* | Tidak disimpan per periode; dihitung dari tanggal rencana atau tanggal persetujuan otorisasi (di bawah). Otorisasi terbuka yang tanggalnya jatuh di periode peralihan terhitung di periode itu |
| Laporan | Ringkasan bulanan (§6.3) mengikuti periode anggaran; periode peralihan punya ringkasannya sendiri dan tampil dengan rentang tanggalnya. Laporan periode lampau tidak berubah karena penggantian. Rincian per anggota tetap hanya di dalam satu periode (§6.3) |
| Entri terlambat dan pengembalian dana | Mengikuti tanggal kejadiannya ke periode yang memuat tanggal itu, menurut batas yang berlaku bagi periode tersebut |
| *Jatah pribadi* | Per periode, termasuk periode peralihan, dengan jumlah penuh (§3.4) |
| Penggantian beruntun | Yang berlaku adalah penggantian terakhir yang diterima sebelum periode berjalan berakhir |
| Siapa yang boleh mengganti | **Kesepakatan rumah tangga: kedua anggota dewasa, seperti nilai ambang** (putusan lead, 11 November 2026; bukan keputusan pemilik produk; T-25). Satu mengusulkan, yang lain menyetujui; hari mulai lama berlaku sampai disetujui, dan tidak seorang pun menggantinya sendirian. Alurnya milik dokumen 06 §6.3 "Kesepakatan rumah tangga". Alasannya dari sisi uang: penggantian membuka satu periode pendek dengan batas anggaran dan jatah pribadi penuh |
| Terlihat | Usulan, persetujuan, dan penggantian itu sendiri tercatat sebagai *peristiwa* beserta pelakunya dan terlihat oleh kedua pasangan; periode peralihan tampil dengan rentang tanggalnya |

```text
urutan baris budget_period_setting: menurut (effective_from, recorded_at)
s0 = baris pertama menurut urutan itu   -- setelan awal rumah tangga; bawaan start_day = 1

periode(tanggal t) =
  s = baris terakhir menurut urutan itu dengan s.effective_from ≤ t
  bila tidak ada:  s = s0
      -- s0 berlaku pula atas semua tanggal sebelum effective_from-nya
      --   (entri terlambat dan saldo awal bertanggal sebelum rumah tangga mulai mencatat)
  p = tanggal pertama ≥ s.effective_from yang harinya s.start_day
  bila s ≠ s0 dan t < p:
      [s.effective_from, p − 1]                               -- periode peralihan
  selain itu:
      [hari s.start_day terakhir yang ≤ t, sehari sebelum hari s.start_day berikutnya]
      -- s0 tidak pernah membuka periode peralihan, walaupun effective_from-nya
      --   bukan hari start_day: tanggal-tanggal pertamanya masuk periode biasa
      --   yang mulai pada hari start_day sebelumnya

penggantian hari mulai menjadi d (d ≠ start_day yang berlaku), diterima pada tanggal c:
  E = akhir periode(c);  tambah baris (start_day = d, effective_from = E + 1)
  -- penggantian kedua dalam periode yang sama menghasilkan effective_from yang sama;
  --   yang berlaku adalah yang recorded_at-nya terakhir (baris "Penggantian beruntun")
```

Pemeriksaan kedua contoh di bawah dengan fungsi ini. Pada keduanya setelan lama adalah s0 (atau baris lain yang periodenya sudah biasa) dan c = 10 November 2026. **25 → 1:** periode(10 November) = [25 Oktober, 24 November], jadi E = 24 November dan baris baru (1, 25 November). Untuk t dari 25 November: s = baris baru ≠ s0; p = 1 Desember; t < p memberi [25 November, 30 November], 6 hari; sejak 1 Desember periode biasa [1 Desember, 31 Desember]. **1 → 25:** periode(10 November) = [1 November, 30 November], jadi E = 30 November dan baris baru (25, 1 Desember); p = 25 Desember; peralihan [1 Desember, 24 Desember], 24 hari; sesudahnya [25 Desember, 24 Januari]. Hasilnya sama dengan tabel. Fungsi lama memberi hasil yang sama untuk kedua contoh ini; yang dikoreksi hanya dua kasus tepi: tanggal sebelum baris pertama (dulu tanpa setelan) dan baris pertama yang tidak mulai pada hari `start_day` (dulu membuka periode peralihan yang tidak pernah disepakati).

| Hari mulai lama → baru, diterima 10 November 2026 | Periode berjalan (tetap) | Periode peralihan | Periode biasa pertama sesudahnya |
|---|---|---|---|
| 25 → 1 | 25 Oktober–24 November | 25–30 November (6 hari) | 1–31 Desember |
| 1 → 25 | 1–30 November | 1–24 Desember (24 hari) | 25 Desember–24 Januari |

Invarian yang berlaku: INV-29 (setiap tanggal di tepat satu periode; periode lampau tidak berubah), INV-15 (satu baris terhitung terhadap paling banyak satu anggaran), INV-19 (jembatan laporan berlaku atas rentang tanggal apa pun, jadi juga atas periode peralihan), INV-12 dan INV-09 (penggantian tidak membuat atau mengubah entri), dan INV-30.

```text
baris anggaran(k, periode) = semua baris b dengan b.category = k, b.owner = shared,
                             event_date entri b di dalam periode, yang berupa:
    sebab expense (termasuk pengembalian dana bertanda +, dan expense hasil recategorization)
    posisi kewajiban pada repayment_out yang berkategori (pembayaran cicilan)
    baris lain yang diberi kategori oleh pencatat: owner_transfer keluar, asset_purchase, loan_given

terpakai(k, periode)  = − Σ akibat baris anggaran itu pada uang atau kekayaan bersih
                        -- pengeluaran 150.000 menambah terpakai 150.000; pengembalian 20.000 menguranginya
terikat(k, periode)   = Σ terikat(a) untuk otorisasi a terbuka dengan a.category = k
                        dan tanggal rencana a (atau tanggal persetujuannya) di dalam periode
sisa(k, periode)      = batas − terpakai − terikat          -- boleh negatif
```

- Jumlah *terikat* mengurangi sisa anggaran dan tampil terpisah dari yang terpakai (Q-14). Jalan-jalan dan acara tidak punya anggaran periode; batasnya adalah otorisasinya sendiri. Bila permintaannya tidak berkategori, atau kategorinya tidak beranggaran, tidak ada sisa anggaran yang dikurangi.
- `credit_cost`, `write_off`, `valuation`, `cash_difference` yang belum dirinci, dan `opening` tidak pernah terhitung terhadap anggaran.
- Kontribusi belanja dan talangan terhitung penuh: yang diukur anggaran adalah pemakaian untuk kategori itu, bukan dari kantong mana uangnya.
- Mengubah batas anggaran tidak membuat entri; perubahan tercatat di jejak audit.
- Tidak ada rollover di Tier 1: sisa periode lalu tidak berpindah.

### 6.2 Realisasi dan terikat

Didefinisikan di §3.8 dan INV-14. *Realisasi* atas satu anggaran sama dengan `terpakai` di §6.1.

### 6.3 Ringkasan bulanan

Kriteria sukses nomor 3. Periode ringkasan adalah *periode anggaran* rumah tangga (§6.1). Setiap baris sebab buku rumah tangga dalam periode masuk tepat satu baris tabel A; jumlah tabel A = perubahan kekayaan bersih rumah tangga pada periode itu (INV-19).

**Tabel A — mengapa harta bersih berubah.**

| Baris | Rumus |
|---|---|
| Uang masuk (pemasukan), per sumber dan atas nama siapa | `Σ effect(income)` |
| Belanja per kategori | `− Σ effect(expense)` per kategori, sesudah pengembalian dana; baris `expense` tanpa kategori dijumlahkan pada baris "tanpa kategori" (§3.3, §3.14) |
| Biaya kredit | `− Σ effect(credit_cost)` |
| Selisih hitung yang belum dirinci | `Σ effect(cash_difference)`, negatif dan positif ditampilkan terpisah |
| Piutang diikhlaskan | `− Σ effect(write_off)` |
| Perubahan nilai aset | `Σ effect(valuation)` |
| Kontribusi dari dana pribadi, per anggota | `Σ effect(owner_transfer > 0)` pada entri `owner_transfer`, `contribution_expense`, dan `claim_waiver` (anggota melepas hak gantinya, §3.6) |
| Disisihkan ke dana pribadi, per anggota (termasuk bagian pemasukan yang langsung ke dana pribadi, §3.1; pemasukannya sendiri tampil penuh di baris pertama) | `− Σ effect(owner_transfer < 0)` pada entri `owner_transfer` dan `income` |
| Utang anggota kepada dana bersama dihapuskan ("utang dihapuskan") | `− Σ effect(owner_transfer < 0)` pada entri `claim_waiver` (§3.6); tidak pernah dijumlahkan ke baris "disisihkan" |
| Klaim ditutup karena anggota keluar | `Σ effect(owner_transfer)` pada entri `exit_closure`, per klaim (§8.2) |
| Utang yang dibebaskan pihak luar | `Σ effect(debt_forgiven)` |

Setiap baris sebab `owner_transfer` di buku rumah tangga masuk tepat satu dari empat baris di atas, menurut jenis entrinya dan tandanya; jenis entri lain tidak memuat sebab itu (§2.5). Baris entri pembalik masuk baris yang sama dengan baris entri asalnya dan menguranginya, bukan baris yang sesuai dengan tandanya sendiri: penyisihan yang dibalik mengurangi "disisihkan", tidak menambah "kontribusi".

**Angka per anggota atas penyisihan dan kontribusi hanya per periode [Usulan]** (temuan `security-privacy` SP-13; dokumen 06 §9.4 "Keterangan pelaku tanpa menjadi alat pantau", AT-4). Baris "kontribusi dari dana pribadi" dan "disisihkan ke dana pribadi" dirinci per anggota hanya di dalam **satu periode laporan**. Tidak ada tampilan, laporan, ekspor, atau narasi yang memuat, per anggota: (a) jumlah kumulatif sejak awal; (b) jumlah lintas periode, termasuk tren, rata-rata, dan perbandingan antar-periode; (c) angka bersih penyisihan dikurangi kontribusi. Laporan yang rentangnya lebih dari satu periode menampilkan kedua baris itu sebagai jumlah rumah tangga, tanpa rincian per anggota. Alasannya: penyisihan dikurangi kontribusi adalah batas atas kasar dana pribadi seorang anggota yang dana pribadinya hanya berasal dari dana bersama, yaitu dari penyisihan termasuk bagian pribadi penghasilannya sendiri (§5.3, B4); aplikasi tidak boleh menjadi alat yang menghitungnya. Batasnya dinyatakan di B4. Aturan yang sama berlaku atas kolom "dana pribadi, kontribusi" pada tabel C.

**Tabel B — uang keluar per kategori** (dasar yang sama dengan anggaran): `terpakai(k, periode)` untuk setiap kategori, termasuk cicilan yang dibayar, ditambah satu baris **"tanpa kategori"** bila ada: baris sebab `expense` tanpa kategori dan pembayaran cicilan (`repayment_out` yang tertaut ke cicilan) tanpa kategori (**[Usulan]**, putaran S3; PA-FIN-7). Baris itu tidak punya anggaran; ia ada supaya pengeluaran yang kategorinya kosong atau terarsip tetap terhitung satu kali. Pada CTH-37 baris itu tidak ada. Tabel B menjawab "uang kita dipakai untuk apa"; tabel A menjawab "mengapa harta kita berubah". Pembayaran cicilan ada di B dan tidak di A; biaya kredit ada di A dan tidak di B. Kedua tabel tidak pernah dijumlahkan.

**Tabel C — siapa membayar apa.** Baris `expense` dan pembayaran cicilan dikelompokkan menurut `payer`, dengan tiga kolom sumber: dana bersama; dana pribadi sebagai kontribusi; dana pribadi yang diganti (talangan). Jumlah tabel C per kategori = tabel B. Laporan per orang hanya memuat buku rumah tangga dan sama bagi keduanya (Q-27).

**Tabel D — tabungan.** Per tujuan berjenis `savings`: jumlah alokasi sekarang, perubahan dalam periode, target bila ada. Tujuan yang terarsip tampil pada periode pelepasan alokasinya dengan keterangan "diarsipkan", lalu tidak lagi (§3.9; G-01).

**Tabel E — aset, piutang, kewajiban.** Saldo akhir tiap posisi dan perubahannya dalam periode. Klaim internal yang sedang dikeberatani diberi keterangan "dikeberatani" tanpa perubahan angka (§3.12).

Laporan periode lampau yang berubah karena entri terlambat menandainya (§3.11).

### 6.4 Arus kas

`ΔU(shared)` dalam periode, dirinci dari baris tempat uang menurut jenis entrinya. Jumlah semua baris = perubahan *uang yang ada* (INV-19). Pemindahan tidak muncul.

| Baris arus kas | Dari |
|---|---|
| Uang masuk: pemasukan | `income`: yang benar-benar masuk tempat uang dana bersama, yaitu jumlah pemasukan dikurangi bagian yang langsung ke dana pribadi; bagian itu disebut di samping barisnya dan tidak menjadi baris "disisihkan" tersendiri di arus kas, karena tidak pernah berada di tempat uang dana bersama |
| Uang masuk: pinjaman diterima | `loan_received` |
| Uang masuk: piutang dibayar | `repayment_in` |
| Uang masuk: kontribusi uang dari dana pribadi | `owner_transfer` masuk |
| Uang masuk: hasil jual aset | `asset_disposal` |
| Uang masuk: selisih hitung positif | `cash_difference` > 0 |
| Uang keluar: belanja (sesudah pengembalian dana, termasuk biaya admin) | `expense`, `refund`, baris tempat uang saja |
| Uang keluar: cicilan dan utang dibayar | `repayment_out` |
| Uang keluar: penggantian dibayar | `reimbursement` |
| Uang keluar: pinjaman diberikan | `loan_given` |
| Uang keluar: beli aset (termasuk uang muka) | `asset_purchase` |
| Uang keluar: disisihkan ke dana pribadi | `owner_transfer` keluar |
| Uang keluar: selisih hitung negatif | `cash_difference` < 0, sebesar selisih penuh |

Pengeluaran yang tidak memakai dana bersama saat itu (kontribusi belanja, talangan, biaya kredit, selisih yang dirinci) tidak punya baris tempat uang di buku rumah tangga dan karena itu tidak muncul di arus kas; uangnya muncul lewat baris penggantian, cicilan, atau selisih.

### 6.5 Kekayaan bersih

```text
kekayaan bersih rumah tangga = U(shared) + P(shared) + A(shared) − K(shared)
```

`P(shared)` dan `K(shared)` memuat klaim kepada pihak luar **dan** kepada anggota (hak ganti, utang antar-pasangan dengan dana bersama), karena dana pribadi di luar garis. Titipan pada pihak luar ada di `U`, bukan `P`. Target, anggaran, otorisasi, dan reservasi tidak masuk.

---

## 7. Contoh terhitung

### 7.1 Cara membaca

Setiap contoh kelak menjadi executable test dengan angka yang sama (`ENGINEERING-STANDARDS.md` §5).

**Rumah tangga contoh [Asumsi].** Angka brief §1.1 (gaji 10.000.000; kiriman ibu 3.000.000) ditambah angka ilustrasi. Rangkaian contoh hanya memuat satu penghasilan dalam periodenya, gaji suami, yang ia terima dan ia catat sendiri. Itu penyederhanaan contoh, bukan gambaran rumah tangga pilot: menurut pemilik produk (11 November 2026) kedua pasangan berpenghasilan. Penghasilan istri dicatat dengan bentuk yang sama (CTH-03, CTH-03b) dengan `earner` istri; semua aturan berlaku per penerima penghasilan dan sama bagi anggota dewasa mana pun. Satu periode dengan beberapa penghasilan, ke rekening bank dan dompet digital, ada di CTH-39. Hari mulai periode anggaran rumah tangga contoh adalah 25 (pilihan rumah tangga itu; bawaannya 1, §6.1); periode contoh 25 Oktober–24 November 2026. Ambang permintaan 500.000.

**Kiriman kepada ibu di dalam contoh adalah ilustrasi.** Angka 3000000 berasal dari brief §1.1 dan tetap dipakai agar setiap langkah arus uang di brief punya contoh. Menurut pemilik produk (checkpoint S2), kiriman untuk orang tua **tidak rutin** di rumah tangga pilot; CTH-06 dan CTH-39 karenanya bukan gambaran pembayaran rutin rumah tangga itu. Urutan yang ia nyatakan (uang masuk, bayar tagihan, siapkan uang untuk tagihan yang menyusul, bagikan sisanya) ada di CTH-41.

**Kolom.** Semua angka adalah rupiah utuh, ditulis tanpa pemisah ribuan agar dapat dihitung ulang dengan skrip.

| Kolom | Arti | Pemilik | Pemegang |
|---|---|---|---|
| `BNK-S` | Rekening bank tempat gaji masuk | `shared` | suami |
| `TUN-S` | Uang tunai bersama di suami | `shared` | suami |
| `TUN-I` | Uang tunai bersama di istri | `shared` | istri |
| `EWL-I` | Dompet digital | `shared` | istri |
| `TAB-I` | Rekening tabungan dana darurat | `shared` | istri |
| `UANG` | *Uang yang ada* = `BNK-S + TUN-S + TUN-I + EWL-I + TAB-I` | | |
| `PIU` | Piutang rumah tangga (pinjaman kepada adik suami) | `shared` | |
| `AST` | Nilai aset rumah tangga | `shared` | |
| `KWJ-C` | Kewajiban cicilan: jumlah seluruh pembayaran tersisa | `shared` | |
| `KWJ-U` | Utang rumah tangga kepada pihak luar | `shared` | |
| `KWJ-G` | Hak ganti yang belum dibayar kepada anggota (kewajiban rumah tangga; sisi lainnya piutang anggota itu di buku pribadinya) | `shared` | |
| `KB` | Kekayaan bersih rumah tangga = `UANG + PIU + AST − KWJ-C − KWJ-U − KWJ-G` | | |
| `PRI-S` | Uang tunai pribadi suami (tidak terlihat istri) | `member(suami)` | suami |
| `PRI-I` | Uang tunai pribadi istri (tidak terlihat suami) | `member(istri)` | istri |

Baris `Δ` adalah entri; baris `Sesudah` = baris sebelumnya + `Δ`, kolom demi kolom. Contoh CTH-01 dan CTH-03 sampai CTH-36 **berantai**: `Sesudah` terakhir satu contoh adalah `Sebelum` contoh berikutnya. CTH-02 dan contoh bertanda "berdiri sendiri" tidak mengubah keadaan berantai; semuanya memakai kolomnya sendiri, kecuali CTH-03b, yang memakai kolom rantai dengan `Sebelum` yang sama dengan CTH-03.

**Peta kasus brief §7 ke contoh.**

| Kasus brief §7 | Contoh |
|---|---|
| Spending from personal cash | CTH-12 |
| Spending from shared household cash | CTH-11 |
| Husband transfers money to wife | CTH-08 |
| Allocating money to mother-in-law | CTH-06 (kiriman keluarga), CTH-02 (titipan) |
| Paying on behalf of another member | CTH-17, CTH-17b |
| Reimbursement | CTH-14, CTH-15, CTH-23 |
| Moving money between bank, e-wallet, and cash | CTH-07, CTH-18 |
| Full and partial refunds | CTH-19 |
| Recurring bills | CTH-05 |
| Split expenses | CTH-20 |
| Debt repayment | CTH-28 |
| Savings contributions | CTH-10 |
| Budget reservations | CTH-22 |
| Settling unused budget | CTH-21, CTH-22, CTH-23 |
| Cash reconciliation | CTH-29 |
| Incorrect transactions and corrections | CTH-30 |
| Concurrent updates | CTH-31 |
| Duplicate submissions | CTH-32 |
| Salary arrival and allocation as in §1.1 | CTH-03 sampai CTH-09 |
| An installment payment | CTH-04, CTH-34 |
| A loan to extended family and its repayment | CTH-25 |
| One spouse borrowing back | CTH-27 |
| An asset purchase and its effect on net worth | CTH-33, CTH-34 |
| Tambahan: custody §5C | CTH-02 |
| Tambahan: alur acuan §5D per situasi custody | CTH-21 (A), CTH-22 (B dan D), CTH-23 (C) |
| Tambahan: kelebihan | CTH-24 |
| Tambahan: kontribusi dana pribadi tanpa dan dengan "minta diganti"; pelepasan hak ganti | CTH-13, CTH-14, CTH-16 |
| Tambahan: entri terlambat yang mengurangi selisih kas | CTH-29 |
| Tambahan: bantuan tidak rutin ke keluarga besar; penghapusan piutang | CTH-26, CTH-25 |
| Tambahan: penyisihan ke dana pribadi; hadiah pribadi | CTH-35, CTH-36 |
| Tambahan: gaji dengan bagian langsung ke dana pribadi (keputusan pemilik produk, 11 November 2026) | CTH-03b |
| Tambahan: beberapa penghasilan dalam satu periode, jatah pribadi bulanan sebagai batas, tarik tunai berbiaya, serah terima dari rekening, dua kiriman keluarga (keputusan dan pernyataan pemilik produk, 11 November 2026 dan checkpoint S2) | CTH-39 |
| Tambahan: menghitung dompet campuran, dana bersama dan dana pribadi dalam satu dompet (fakta rumah tangga pilot, checkpoint S2, S2-6) | CTH-40 |
| Tambahan: uang masuk, bayar tagihan, siapkan uang untuk tagihan yang jatuh temponya menyusul, bagikan sisanya; tagihan yang jumlahnya berubah-ubah (fakta rumah tangga pilot, checkpoint S2, S2-4) | CTH-41 |
| Tambahan: perubahan syarat atas otorisasi terbuka | §7.6, varian CTH-22 |
| Tambahan: klaim antar-pemilik ditutup saat anggota keluar | CTH-38 |
| Tambahan: ringkasan bulanan, arus kas, anggaran | CTH-37 |
| Putaran S3: saldo awal oleh bukan pemegang dan koreksinya (PA-FIN-2, PA-FIN-3) | CTH-42 |
| Putaran S3: entri terlambat atas tempat uang terarsip (PA-FIN-4) | CTH-43 |
| Putaran S3: satu belanja dari dana bersama dan dana pribadi (PA-FIN-6) | CTH-44 |
| Putaran S3: membatalkan tujuan tabungan yang beralokasi (G-01) | CTH-45 |
| Putaran S3: hitungan offline yang tiba sesudah serah terima (G-02) | CTH-46 |
| Putaran S3: hitungan oleh bukan pemegang (G-05); tanda pemegang dan satu hitungan terbuka (SP-50) | CTH-47 |
| Putaran perbaikan S3: mengikhlaskan piutang bertahap, dinilai kumulatif (SP-51) | CTH-25 varian SP-51 |
| Putaran perbaikan S3: klaim ditutup pada akhir masa tenggang keluar (SP-48) | CTH-38 |
| Putaran S3: menghitung lalu merinci seluruh selisih sekaligus (dokumen 11, IDEA-01) | CTH-48 |
| Putaran S3: cicilan yang jumlah bayarnya di bawah harga; cicilan 0% (G-07) | CTH-49 |
| Sesudah checkpoint S3: pos rutin tanpa tanggal (IDEA-03) | CTH-50 |
| Sesudah checkpoint S3: THR dan angpau yang diberikan, satu permintaan musiman (pernyataan pemilik produk, 12 November 2026) | CTH-51 |

### 7.2 Saldo awal dan custody

#### CTH-01 — Saldo awal

Hari pertama mencatat (24 Oktober 2026). Motor bernilai 12000000 dengan sisa cicilan 6 × 1500000 = 9000000. Tiap kolom posisi adalah satu entri `opening`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Δ saldo awal | 500000 | 200000 | 300000 | 100000 | 5000000 | 6100000 | 0 | 12000000 | 9000000 | 0 | 0 | 9100000 | 400000 | 250000 |
| Sesudah | 500000 | 200000 | 300000 | 100000 | 5000000 | 6100000 | 0 | 12000000 | 9000000 | 0 | 0 | 9100000 | 400000 | 250000 |

Menguji: INV-02, INV-05, INV-07. Saldo awal bukan pemasukan dan tidak masuk laporan periode.

#### CTH-02 — Custody brief §5C (berdiri sendiri)

Kolom: `TUN-S` tunai suami, `TUN-I` tunai istri, `TTP-IBU` titipan pada ibu mertua untuk makan rumah tangga (jenis `entrusted_external`, pemilik `shared`), `BNK-B` rekening bersama (`holder = joint`), `UANG` = jumlah keempatnya.

| Baris | TUN-S | TUN-I | TTP-IBU | BNK-B | UANG |
|---|---|---|---|---|---|
| Sebelum | 5000000 | 3000000 | 0 | 2000000 | 10000000 |
| Δ suami menitipkan 1000000 kepada ibu mertua (`transfer`) | -1000000 | 0 | 1000000 | 0 | 0 |
| Sesudah (keadaan brief §5C) | 4000000 | 3000000 | 1000000 | 2000000 | 10000000 |
| Δ ibu mertua membeli lauk 150000, dicatat istri (`expense`) | 0 | 0 | -150000 | 0 | -150000 |
| Sesudah | 4000000 | 3000000 | 850000 | 2000000 | 9850000 |

Custody menampilkan empat baris: dipegang suami 4000000; dipegang istri 3000000; dititipkan pada ibu mertua 850000; dipegang bersama (rekening) 2000000. Titipan bukan pengeluaran sampai dibelanjakan. Menguji: INV-05, INV-06.

### 7.3 Gaji tiba dan dibagi (brief §1.1)

#### CTH-03 — Gaji tiba

Suami mencatat gaji 10000000 masuk `BNK-S`: `income`, sumber `salary`, `earner` suami, pemilik `shared` (bawaan).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 200000 | 300000 | 100000 | 5000000 | 6100000 | 0 | 12000000 | 9000000 | 0 | 0 | 9100000 | 400000 | 250000 |
| Δ gaji | 10000000 | 0 | 0 | 0 | 0 | 10000000 | 0 | 0 | 0 | 0 | 0 | 10000000 | 0 | 0 |
| Sesudah | 10500000 | 200000 | 300000 | 100000 | 5000000 | 16100000 | 0 | 12000000 | 9000000 | 0 | 0 | 19100000 | 400000 | 250000 |

Menguji: INV-02, INV-07.

#### CTH-03b — Gaji masuk, sebagian langsung ke dana pribadi (berdiri sendiri)

Varian CTH-03 menurut keputusan pemilik produk, 11 November 2026; `Sebelum`-nya sama dengan CTH-03 dan hasilnya **tidak** diteruskan ke CTH-04. Suami mencatat gaji 10000000 dan mengarahkan 1000000 ke dana pribadinya: 9000000 masuk `BNK-S`, 1000000 masuk `PRI-S`. Satu entri `income`, sumber `salary`, `earner` suami, dengan baris di dua buku. Bila seluruh gaji secara fisik tiba di satu rekening, bagian pribadinya dicatat pada tempat uang pribadi tersendiri (§2.3: satu tempat uang, satu pemilik); kolom `PRI-S` mewakilinya di sini.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 200000 | 300000 | 100000 | 5000000 | 6100000 | 0 | 12000000 | 9000000 | 0 | 0 | 9100000 | 400000 | 250000 |
| Δ gaji 10000000, 1000000 ke pribadi | 9000000 | 0 | 0 | 0 | 0 | 9000000 | 0 | 0 | 0 | 0 | 0 | 9000000 | 1000000 | 0 |
| Sesudah | 9500000 | 200000 | 300000 | 100000 | 5000000 | 15100000 | 0 | 12000000 | 9000000 | 0 | 0 | 18100000 | 1400000 | 250000 |

Baris entrinya:

| Buku | Baris | Nilai |
|---|---|---|
| Rumah tangga | Tempat uang `BNK-S` | 9000000 |
| Rumah tangga | Sebab `income` (gaji, atas nama suami) | 10000000 |
| Rumah tangga | Sebab `owner_transfer` (ke dana pribadi suami) | -1000000 |
| Pribadi suami | Tempat uang `PRI-S` | 1000000 |
| Pribadi suami | Sebab `owner_transfer` (dari dana bersama) | 1000000 |

Keseimbangan buku rumah tangga: 9000000 = 10000000 − 1000000. Buku pribadi suami: 1000000 = 1000000. Jumlah `owner_transfer`: −1000000 + 1000000 = 0. Istri melihat "gaji 10000000 masuk, 1000000 ke pribadi suami"; ia tidak melihat `PRI-S`, saldonya 1400000, atau pemakaiannya sesudah itu. Ringkasan bulanan: pemasukan 10000000, disisihkan ke dana pribadi suami 1000000, perubahan `KB` 9000000; arus kas: pemasukan 9000000. Dibanding CTH-03, `UANG` dan `KB` lebih kecil tepat 1000000 dan `PRI-S` lebih besar tepat 1000000: tidak ada rupiah yang hilang atau terhitung dua kali. Apakah bagian pribadi 1000000 ditandai bergantung pada jatah pribadi suami (§3.4); contoh ini tidak menetapkan jatahnya, dan angka di tabel sama pada ketiga keadaan. Dengan jatah berlaku 1000000 atau lebih dan belum terpakai: 0 + 1000000 tidak melebihi jatah, tidak ditandai. Dengan jatah berlaku lebih kecil, misalnya 600000: 0 + 1000000 > 600000, ditandai. Tanpa jatah: 1000000 di atas ambang contoh 500000, ditandai. Pencatat dan `earner` sama-sama suami, jadi entri ini memenuhi INV-27; perintah yang sama dari istri (bagian pribadi ke `PRI-I` atas gaji atas nama suami) ditolak (INV-27; keputusan pemilik produk, checkpoint S2). Aturannya simetris: penghasilan istri dicatat istri dengan bentuk yang sama, `earner` istri, dan bagian pribadinya ke `PRI-I`. Menguji: INV-02, INV-03, INV-20, INV-24, INV-25, INV-27.

#### CTH-04 — Pembayaran cicilan

Cicilan motor 1500000 dibayar dari `BNK-S`: `repayment_out`, kategori "cicilan". Sisa 5 kali.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 10500000 | 200000 | 300000 | 100000 | 5000000 | 16100000 | 0 | 12000000 | 9000000 | 0 | 0 | 19100000 | 400000 | 250000 |
| Δ cicilan | -1500000 | 0 | 0 | 0 | 0 | -1500000 | 0 | 0 | -1500000 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 9000000 | 200000 | 300000 | 100000 | 5000000 | 14600000 | 0 | 12000000 | 7500000 | 0 | 0 | 19100000 | 400000 | 250000 |

Anggaran "cicilan" (batas 1500000): terpakai 1500000, sisa 0. Kekayaan bersih tidak berubah; uang keluar tampil di arus kas dan di belanja per kategori. Menguji: INV-02, INV-13, INV-15.

#### CTH-05 — Tagihan berulang

Tiga tagihan dibayar suami dari `BNK-S`, masing-masing `expense` yang ditautkan ke tagihan berulangnya, kategori "tagihan rutin": listrik 450000 (perkiraan 400000), wifi 350000, langganan kerja 200000.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 9000000 | 200000 | 300000 | 100000 | 5000000 | 14600000 | 0 | 12000000 | 7500000 | 0 | 0 | 19100000 | 400000 | 250000 |
| Δ listrik | -450000 | 0 | 0 | 0 | 0 | -450000 | 0 | 0 | 0 | 0 | 0 | -450000 | 0 | 0 |
| Sesudah | 8550000 | 200000 | 300000 | 100000 | 5000000 | 14150000 | 0 | 12000000 | 7500000 | 0 | 0 | 18650000 | 400000 | 250000 |
| Δ wifi | -350000 | 0 | 0 | 0 | 0 | -350000 | 0 | 0 | 0 | 0 | 0 | -350000 | 0 | 0 |
| Sesudah | 8200000 | 200000 | 300000 | 100000 | 5000000 | 13800000 | 0 | 12000000 | 7500000 | 0 | 0 | 18300000 | 400000 | 250000 |
| Δ langganan kerja | -200000 | 0 | 0 | 0 | 0 | -200000 | 0 | 0 | 0 | 0 | 0 | -200000 | 0 | 0 |
| Sesudah | 8000000 | 200000 | 300000 | 100000 | 5000000 | 13600000 | 0 | 12000000 | 7500000 | 0 | 0 | 18100000 | 400000 | 250000 |

Sebelum dibayar, tagihan yang jatuh tempo tidak mengubah angka apa pun (INV-12). Yang tercatat adalah jumlah sebenarnya, bukan perkiraan. Anggaran "tagihan rutin" (batas 1100000): terpakai 1000000, sisa 100000.

**Penilaian tanda** (§3.8 "Tanda atas entri", §3.15; revisi kedua pasca-checkpoint S2). Ambang contoh 500000. **[Asumsi]** contoh: ketiga tagihan disepakati berdua saat dibuat, sebelum periode contoh, dengan jumlah per kali berikut. Kesepakatan tidak membuat entri dan tidak mengubah angka (INV-12), sehingga rantai contoh tidak memuatnya; **tidak ada saldo di tabel atas yang berubah**.

| Tagihan | Perkiraan | Disepakati per kali `R` [Asumsi] | Dibayar `P` | `P > R`? | Tanda |
|---|---|---|---|---|---|
| Listrik | 400000 | 500000 | 450000 | Tidak | Tidak |
| Wifi | 350000 | 350000 | 350000 | Tidak | Tidak |
| Langganan kerja | 200000 | 200000 | 200000 | Tidak | Tidak |

Listrik berubah-ubah, jadi yang disepakati adalah jumlah tertingginya yang wajar (500000), bukan perkiraannya. **Varian:** bila listrik disepakati sebesar perkiraannya, 400000, pembayaran 450000 **ditandai** dengan dasar `routine_amount` (450000 > 400000), walau 450000 tidak di atas ambang 500000; kelebihannya, 50000, tidak dibandingkan dengan apa pun. Bila usulan kesepakatan listrik belum dijawab pasangan, pembayaran 450000 dinilai dengan aturan biasa: 450000 tidak di atas 500000, tidak ditandai.

#### CTH-06 — Kiriman bulanan kepada ibu suami

**Ilustrasi atas angka brief §1.1, bukan pembayaran rutin rumah tangga pilot** (pernyataan pemilik produk, checkpoint S2: kiriman untuk orang tua tidak rutin di rumah tangganya; §7.1). Angkanya tidak berubah. Suami mentransfer 3000000 dari `BNK-S` kepada ibunya: `expense`, kategori "kiriman keluarga", penerima manfaat pihak luar "Ibu". Tidak dilacak sesudahnya.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 8000000 | 200000 | 300000 | 100000 | 5000000 | 13600000 | 0 | 12000000 | 7500000 | 0 | 0 | 18100000 | 400000 | 250000 |
| Δ kiriman ibu | -3000000 | 0 | 0 | 0 | 0 | -3000000 | 0 | 0 | 0 | 0 | 0 | -3000000 | 0 | 0 |
| Sesudah | 5000000 | 200000 | 300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |

**Andaian tentang ambang [Asumsi].** 3000000 di atas ambang contoh 500000. Contoh ini mengandaikan entri itu ditautkan ke tagihan berulang "kiriman Ibu" (§3.15) yang dirujuk satu *kesepakatan rutin* sebesar 3000000 per kali, disepakati berdua sebelum periode contoh (dokumen 06 §6.3, §7.4; dokumen 17 S2-4 pilihan A, bagian pembayaran rutin: keputusan pemilik produk, checkpoint S2), sehingga entri ini tidak ditandai: jumlah pemakaiannya, 3000000, tidak melebihi jumlah per kali yang disepakati, 3000000 (dasar `routine_amount`, §3.8; ambang tidak dibandingkan). Kiriman 3100000 atas kesepakatan yang sama akan ditandai, berapa pun ambangnya. Bagi rumah tangga yang kirimannya tidak rutin, seperti rumah tangga pilot, andaian ini tidak berlaku: kiriman 3000000 diajukan lebih dulu sebagai permintaan (seperti CTH-33), atau tersimpan dengan tanda netral; angkanya sama. Rantai contoh tidak memuat kesepakatan itu karena kesepakatan tidak membuat entri dan tidak mengubah angka (INV-12). Tanpa kesepakatan itu, entri yang sama tersimpan dengan angka yang sama dan membawa *tanda di atas ambang* yang netral (dokumen 06 §6.2); tanda tidak mengubah angka mana pun di tabel ini.

#### CTH-07 — Tarik tunai (bank ke tunai)

Suami menarik 4500000 di ATM bank sendiri, tanpa biaya: `transfer` `BNK-S` → `TUN-S`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 5000000 | 200000 | 300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |
| Δ tarik tunai | -4500000 | 4500000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 500000 | 4700000 | 300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |

Menguji: INV-06.

#### CTH-08 — Suami menyerahkan uang kepada istri

Suami menyerahkan 4000000: `transfer` (serah terima) `TUN-S` → `TUN-I`, dicatat suami, bertanda belum dikonfirmasi sampai istri mengonfirmasi (tanpa batas waktu; istri dapat pula mengajukan keberatan kapan saja). Bila istri yang mencatatnya, tanda itu tidak ada dan angkanya sama.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 4700000 | 300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |
| Δ serah terima | 0 | -4000000 | 4000000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 500000 | 700000 | 4300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |

Custody sesudahnya: suami memegang 1200000 (`BNK-S` 500000 + `TUN-S` 700000); istri memegang 9400000 (`TUN-I` 4300000, di antaranya 4000000 belum dikonfirmasi, + `EWL-I` 100000 + `TAB-I` 5000000); jumlah 10600000. Pemilik uang tidak berubah: semuanya dana bersama. Menguji: INV-05, INV-06.

Dari gaji 10000000, rangkaian CTH-03 sampai CTH-08 mencatat: 1500000 cicilan, 1000000 tagihan, 3000000 kiriman ibu, 4000000 diserahkan kepada istri, 500000 tetap di suami; 1500000 + 1000000 + 3000000 + 4000000 + 500000 = 10000000. Tidak satu rupiah pun terhitung dua kali: `UANG` naik 10000000 lalu turun 5500000 (1500000 + 1000000 + 3000000), dan dua pemindahan tidak mengubahnya.

#### CTH-09 — Alokasi

Istri mengalokasikan 2500000 dari `TUN-I` untuk tujuan "belanja bulanan", dan 5000000 di `TAB-I` untuk tujuan tabungan "dana darurat" (target 10000000). Entri `allocation`; tidak ada saldo yang berubah, jadi tabel utama tidak berubah dari `Sesudah` CTH-08.

| Tempat uang | Saldo tercatat | Alokasi sebelum | Δ alokasi | Alokasi sesudah | Belum dialokasikan sesudah |
|---|---|---|---|---|---|
| TUN-I | 4300000 | 0 | 2500000 | 2500000 | 1800000 |
| TAB-I | 5000000 | 0 | 5000000 | 5000000 | 0 |

Menguji: INV-12, INV-17.

#### CTH-10 — Setoran tabungan

Istri menyetor 500000 dari bagian yang belum dialokasikan di `TUN-I` ke `TAB-I`: `transfer` yang membawa tujuan "dana darurat".

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 4300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |
| Δ setor tabungan | 0 | 0 | -500000 | 0 | 500000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 500000 | 700000 | 3800000 | 100000 | 5500000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |

| Tempat uang | Saldo tercatat | Alokasi sebelum | Δ alokasi | Alokasi sesudah | Belum dialokasikan sesudah |
|---|---|---|---|---|---|
| TUN-I | 3800000 | 2500000 | 0 | 2500000 | 1300000 |
| TAB-I | 5500000 | 5000000 | 500000 | 5500000 | 0 |

Kemajuan "dana darurat": 5500000 dari target 10000000. Menabung bukan pengeluaran: `UANG` dan `KB` tetap. Menguji: INV-06, INV-17.

### 7.4 Belanja, dana pribadi, dan hak ganti

#### CTH-11 — Belanja dari kas bersama

Istri membeli beras dan lauk 150000 dari `TUN-I`: `expense`, kategori "belanja bulanan", tujuan "belanja bulanan", penerima manfaat rumah tangga.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3800000 | 100000 | 5500000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |
| Δ belanja | 0 | 0 | -150000 | 0 | 0 | -150000 | 0 | 0 | 0 | 0 | 0 | -150000 | 0 | 0 |
| Sesudah | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 400000 | 250000 |

Anggaran "belanja bulanan" (batas 2500000): terpakai 150000, sisa 2350000. Alokasi "belanja bulanan" di `TUN-I`: 2500000 − 150000 = 2350000; belum dialokasikan 3650000 − 2350000 = 1300000.

Sesudah contoh ini istri melepas alokasi "belanja bulanan" (entri `allocation` −2350000; alokasi 0, belum dialokasikan `TUN-I` 3650000), karena bagi rumah tangga contoh anggaran sebagai batas sudah cukup (keputusan pemilik nomor 8). Contoh berikutnya tidak membawa tujuan; alokasi "dana darurat" 5500000 di `TAB-I` tidak berubah sampai akhir.

#### CTH-12 — Belanja pribadi dari uang pribadi

Suami membeli kopi 30000 dari `PRI-S`: `expense` di buku pribadinya.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 400000 | 250000 |
| Δ kopi | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | -30000 | 0 |
| Sesudah | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 370000 | 250000 |

Tidak ada angka rumah tangga yang berubah dan istri tidak melihat entri ini. Menguji: INV-20.

#### CTH-13 — Kontribusi dari dana pribadi (tidak minta diganti)

Suami membeli beras 80000 untuk rumah dengan `PRI-S` dan tidak menandai "minta diganti": `contribution_expense`. Buku pribadi: tempat uang −80000, `owner_transfer` −80000. Buku rumah tangga: `owner_transfer` +80000, `expense` −80000 (kategori "belanja bulanan", `funded_by` suami).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 370000 | 250000 |
| Δ beras pakai uang pribadi suami | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | -80000 | 0 |
| Sesudah | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 290000 | 250000 |

Keseimbangan buku rumah tangga: sisi posisi 0; sisi sebab +80000 − 80000 = 0. Istri melihat "beras 80000, pakai uang pribadi suami"; ia tidak melihat `PRI-S`. Anggaran "belanja bulanan": terpakai 150000 + 80000 = 230000, sisa 2270000. Laporan: belanja 80000 terhitung satu kali; kontribusi dari dana pribadi suami 80000. Menguji: INV-02, INV-03, INV-15, INV-20.

#### CTH-14 — Talangan dengan "minta diganti"

Suami membayar gas dan galon 60000 dengan `PRI-S` dan menandai "minta diganti": `advance`. Buku pribadi: tempat uang −60000, piutang (hak ganti) +60000. Buku rumah tangga: kewajiban `KWJ-G` +60000, `expense` −60000 (kategori "belanja bulanan").

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 290000 | 250000 |
| Δ talangan | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 60000 | -60000 | -60000 | 0 |
| Sesudah | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 60000 | 14890000 | 230000 | 250000 |

Kekayaan bersih pribadi suami tidak berubah: `PRI-S` 230000 + hak ganti 60000 = 290000. Anggaran "belanja bulanan": terpakai 290000, sisa 2210000. Menguji: INV-02, INV-08.

#### CTH-15 — Penggantian

Istri mengganti 60000 dari `TUN-I` kepada suami: `reimbursement`, dicatat suami sebagai penerimanya (§2.2). Baris yang mengurangi `TUN-I`, yang dipegang istri, bertanda belum dikonfirmasi sampai istri mengonfirmasi atau mengajukan keberatan (§3.3, sisi pembayar); saldonya sudah berubah.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 60000 | 14890000 | 230000 | 250000 |
| Δ penggantian | 0 | 0 | -60000 | 0 | 0 | -60000 | 0 | 0 | 0 | 0 | -60000 | 0 | 60000 | 0 |
| Sesudah | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 0 | 14890000 | 290000 | 250000 |

Dari CTH-14 dan CTH-15 bersama: kekayaan bersih rumah tangga turun tepat 60000, pengeluaran tercatat satu kali, anggaran tidak berubah oleh penggantian. Menguji: INV-02, INV-08, INV-13.

#### CTH-16 — Hak ganti dilepas

Istri membayar sabun 40000 dengan `PRI-I` dan menandai "minta diganti", lalu melepasnya (`claim_waiver`).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 0 | 14890000 | 290000 | 250000 |
| Δ talangan | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 40000 | -40000 | 0 | -40000 |
| Sesudah | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 40000 | 14850000 | 290000 | 210000 |
| Δ hak ganti dilepas | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | -40000 | 40000 | 0 | 0 |
| Sesudah | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 0 | 14890000 | 290000 | 210000 |

Kenaikan 40000 pada pelepasan punya baris sebab `owner_transfer` +40000 di buku rumah tangga dan tampil sebagai "kontribusi dari dana pribadi istri". Keadaan akhirnya sama dengan bila sejak awal dicatat sebagai kontribusi (bandingkan CTH-13). Anggaran "belanja bulanan": terpakai 330000, sisa 2170000. Menguji: INV-02, INV-03, INV-08.

#### CTH-17 — Membayarkan keperluan anggota lain dari dana bersama

Istri membayar potong rambut suami 50000 dari `TUN-I`: `expense`, kategori "perawatan diri", penerima manfaat suami, pembayar istri. Tidak ada klaim: dana bersama memang untuk keperluan anggota.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 0 | 14890000 | 290000 | 210000 |
| Δ potong rambut | 0 | 0 | -50000 | 0 | 0 | -50000 | 0 | 0 | 0 | 0 | 0 | -50000 | 0 | 0 |
| Sesudah | 500000 | 700000 | 3540000 | 100000 | 5500000 | 10340000 | 0 | 12000000 | 7500000 | 0 | 0 | 14840000 | 290000 | 210000 |

#### CTH-17b — Membayarkan keperluan pribadi anggota lain dari dana pribadi (berdiri sendiri)

Varian: istri membayar 50000 keperluan pribadi suami dengan `PRI-I` dan minta diganti; suami kemudian mengganti dari `PRI-S`. Buku rumah tangga tidak tersentuh. Kolom: `KLAIM` = piutang istri = kewajiban suami; `KB-I = PRI-I + KLAIM`; `KB-S = PRI-S − KLAIM`.

| Baris | PRI-I | PRI-S | KLAIM | KB-I | KB-S |
|---|---|---|---|---|---|
| Sebelum | 210000 | 290000 | 0 | 210000 | 290000 |
| Δ istri menalangi (`advance`) | -50000 | 0 | 50000 | 0 | -50000 |
| Sesudah | 160000 | 290000 | 50000 | 210000 | 240000 |
| Δ suami mengganti (`reimbursement`) | 50000 | -50000 | -50000 | 0 | 0 |
| Sesudah | 210000 | 240000 | 0 | 210000 | 240000 |

Pengeluaran 50000 tercatat satu kali, di buku pribadi suami. Keduanya melihat klaim ini (kebocoran B3). Contoh ini contoh tingkat model: kedua entrinya menyentuh buku pribadi dua anggota, sehingga alur pencatatannya ditunda dari MVP (§2.2, T-12); angkanya tetap menjadi test INV-02 dan INV-08. Menguji: INV-02, INV-08.

### 7.5 Pemindahan, pengembalian dana, pengeluaran terbagi

#### CTH-18 — Bank ke dompet digital dengan biaya admin, lalu tunai ke bank

Suami mengisi `EWL-I` 300000 dari `BNK-S` dengan biaya admin 1000 (satu entri `transfer` dengan baris `expense` 1000, kategori "biaya admin"). Lalu ia menyetor tunai 200000 dari `TUN-S` ke `BNK-S`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3540000 | 100000 | 5500000 | 10340000 | 0 | 12000000 | 7500000 | 0 | 0 | 14840000 | 290000 | 210000 |
| Δ isi dompet digital + biaya admin | -301000 | 0 | 0 | 300000 | 0 | -1000 | 0 | 0 | 0 | 0 | 0 | -1000 | 0 | 0 |
| Sesudah | 199000 | 700000 | 3540000 | 400000 | 5500000 | 10339000 | 0 | 12000000 | 7500000 | 0 | 0 | 14839000 | 290000 | 210000 |
| Δ setor tunai | 200000 | -200000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 400000 | 5500000 | 10339000 | 0 | 12000000 | 7500000 | 0 | 0 | 14839000 | 290000 | 210000 |

Hanya biaya admin yang menjadi pengeluaran. Baris +300000 pada `EWL-I` dicatat suami ke tempat uang yang dipegang istri, jadi ia bertanda belum dikonfirmasi sampai istri mengonfirmasi (§3.3); saldonya sudah berubah. Menguji: INV-06.

#### CTH-19 — Pengembalian dana sebagian dan penuh

Istri membeli panci 120000 lewat `EWL-I` (kategori "perlengkapan rumah"); penjual mengembalikan 20000. Lalu ia membeli blender 100000, membatalkannya, dan menerima 100000 kembali.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 500000 | 3540000 | 400000 | 5500000 | 10339000 | 0 | 12000000 | 7500000 | 0 | 0 | 14839000 | 290000 | 210000 |
| Δ beli panci | 0 | 0 | 0 | -120000 | 0 | -120000 | 0 | 0 | 0 | 0 | 0 | -120000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 280000 | 5500000 | 10219000 | 0 | 12000000 | 7500000 | 0 | 0 | 14719000 | 290000 | 210000 |
| Δ pengembalian sebagian | 0 | 0 | 0 | 20000 | 0 | 20000 | 0 | 0 | 0 | 0 | 0 | 20000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 300000 | 5500000 | 10239000 | 0 | 12000000 | 7500000 | 0 | 0 | 14739000 | 290000 | 210000 |
| Δ beli blender | 0 | 0 | 0 | -100000 | 0 | -100000 | 0 | 0 | 0 | 0 | 0 | -100000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 200000 | 5500000 | 10139000 | 0 | 12000000 | 7500000 | 0 | 0 | 14639000 | 290000 | 210000 |
| Δ pengembalian penuh | 0 | 0 | 0 | 100000 | 0 | 100000 | 0 | 0 | 0 | 0 | 0 | 100000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 300000 | 5500000 | 10239000 | 0 | 12000000 | 7500000 | 0 | 0 | 14739000 | 290000 | 210000 |

"Perlengkapan rumah" terpakai: 120000 − 20000 + 100000 − 100000 = 100000. Tidak ada pemasukan yang tercatat. Pengembalian ketiga atas panci sebesar lebih dari 100000 akan ditolak (INV-13).

#### CTH-20 — Pengeluaran terbagi

Istri membayar 200000 di minimarket dari `TUN-I`: satu entri `expense` dengan dua baris sebab, 150000 "belanja bulanan" dan 50000 "perlengkapan rumah".

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 500000 | 3540000 | 300000 | 5500000 | 10239000 | 0 | 12000000 | 7500000 | 0 | 0 | 14739000 | 290000 | 210000 |
| Δ belanja terbagi | 0 | 0 | -200000 | 0 | 0 | -200000 | 0 | 0 | 0 | 0 | 0 | -200000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3340000 | 300000 | 5500000 | 10039000 | 0 | 12000000 | 7500000 | 0 | 0 | 14539000 | 290000 | 210000 |

150000 + 50000 = 200000. "Belanja bulanan": terpakai 480000, sisa 2020000. "Perlengkapan rumah": terpakai 150000. Pembagian menurut pembayar memakai bentuk yang sama: beberapa baris tempat uang dalam satu entri. Menguji: INV-02, INV-15.

### 7.6 Permintaan, otorisasi, penyelesaian

Kategori "jajan dan makan di luar" beranggaran 500000.

#### CTH-21 — Alur acuan §5D, situasi A: pemohon memegang uangnya

Suami meminta 100000 untuk membeli makanan; istri menyetujui; ia membeli 75000 dari `TUN-S`; penyelesaian.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 500000 | 3340000 | 300000 | 5500000 | 10039000 | 0 | 12000000 | 7500000 | 0 | 0 | 14539000 | 290000 | 210000 |
| Δ disetujui (otorisasi 100000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3340000 | 300000 | 5500000 | 10039000 | 0 | 12000000 | 7500000 | 0 | 0 | 14539000 | 290000 | 210000 |
| Δ beli makanan 75000 | 0 | -75000 | 0 | 0 | 0 | -75000 | 0 | 0 | 0 | 0 | 0 | -75000 | 0 | 0 |
| Sesudah | 399000 | 425000 | 3340000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |
| Δ penyelesaian (sisa 25000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 425000 | 3340000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |

| Saat | Batas anggaran | Terpakai | Terikat | Sisa |
|---|---|---|---|---|
| Sesudah disetujui | 500000 | 0 | 100000 | 400000 |
| Sesudah beli | 500000 | 75000 | 25000 | 400000 |
| Sesudah penyelesaian | 500000 | 75000 | 0 | 425000 |

Tanda pada CTH-21 sampai CTH-23 (§3.8 "Tanda atas entri"): tiap belanja 75000 berada seluruhnya dalam batas otorisasinya 100000 (`N` = 0), jadi tidak ada entri yang ditandai; serah terima dan penggantian tidak punya jumlah pemakaian.

"Benar" di situasi A: sisa 25000 tidak pernah berpindah, tetap dana bersama di `TUN-S`; penyelesaian hanya melepas terikat. Menguji: INV-12, INV-14, INV-15.

#### CTH-22 — Situasi B dengan reservasi: uang diserahkan untuk permintaan

Istri meminta 100000 untuk kue arisan dan menyebut `TUN-S` sebagai sumber; suami menyetujui (lahir reservasi 100000 di `TUN-S`); suami menyerahkan 100000; istri membeli 75000; istri mengembalikan sisa 25000.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 425000 | 3340000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |
| Δ disetujui (otorisasi dan reservasi 100000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 425000 | 3340000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |
| Δ serah terima 100000 | 0 | -100000 | 100000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 325000 | 3440000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |
| Δ beli kue 75000 | 0 | 0 | -75000 | 0 | 0 | -75000 | 0 | 0 | 0 | 0 | 0 | -75000 | 0 | 0 |
| Sesudah | 399000 | 325000 | 3365000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 0 | 14389000 | 290000 | 210000 |
| Δ penyelesaian: sisa 25000 dikembalikan | 0 | 25000 | -25000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 0 | 14389000 | 290000 | 210000 |

| Saat | TUN-S tercatat | Reservasi aktif | TUN-S tersedia |
|---|---|---|---|
| Sebelum | 425000 | 0 | 425000 |
| Sesudah disetujui | 425000 | 100000 | 325000 |
| Sesudah serah terima | 325000 | 0 | 325000 |
| Sesudah penyelesaian | 350000 | 0 | 350000 |

| Saat | Batas anggaran | Terpakai | Terikat | Sisa |
|---|---|---|---|---|
| Sesudah disetujui | 500000 | 75000 | 100000 | 325000 |
| Sesudah beli | 500000 | 150000 | 25000 | 325000 |
| Sesudah penyelesaian | 500000 | 150000 | 0 | 350000 |

"Benar" di situasi B: sisa 25000 berada di tangan pemohon sebagai dana bersama. Pilihan (1), dipakai di atas: dikembalikan lewat serah terima. Pilihan (2): tetap dipegang pemohon, tanpa entri; keadaan akhirnya `TUN-S` 325000, `TUN-I` 3365000, `UANG` 9889000. Pada kedua pilihan `UANG` dan `KB` sama. Menguji: INV-06, INV-11, INV-12, INV-14.

**Varian situasi D (berdiri sendiri).** Bila suami, pemegang `TUN-S`, membayar kue 75000 langsung dari `TUN-S` atas permintaan yang sama:

| Saat | TUN-S tercatat | Reservasi aktif | TUN-S tersedia |
|---|---|---|---|
| Sesudah disetujui | 425000 | 100000 | 325000 |
| Sesudah membayar 75000 | 350000 | 25000 | 325000 |
| Sesudah penyelesaian | 350000 | 0 | 350000 |

Sisa 25000 tidak pernah meninggalkan `TUN-S`; penyelesaian hanya melepas reservasinya.

**Varian perubahan syarat (berdiri sendiri).** Bila sesudah disetujui, dan sebelum ada serah terima atau belanja, istri mengusulkan batas 150000 dan suami menerimanya (`authorization.amended`): tidak ada entri dan tidak ada saldo tercatat yang berubah; reservasi dan *terikat* mengikuti batas baru.

| Saat | Batas otorisasi | TUN-S tercatat | Reservasi aktif | TUN-S tersedia | Terpakai anggaran | Terikat | Sisa anggaran |
|---|---|---|---|---|---|---|---|
| Sesudah disetujui | 100000 | 425000 | 100000 | 325000 | 75000 | 100000 | 325000 |
| Sesudah perubahan syarat diterima | 150000 | 425000 | 150000 | 275000 | 75000 | 150000 | 275000 |

Batas anggaran 500000: 500000 − 75000 − 150000 = 275000. `TUN-S` tersedia: 425000 − 150000 = 275000. Menguji: INV-11, INV-12, INV-14.

#### CTH-23 — Situasi C: dibayar dengan dana pribadi atas otorisasi dana bersama

Suami meminta 100000 untuk obat (kategori "kesehatan"); istri menyetujui; ia membayar 75000 dengan `PRI-S` (tanda "minta diganti" hidup secara bawaan karena tertaut ke otorisasi dana bersama); penyelesaian; istri mengganti dari `TUN-I`. Talangan dan penggantiannya dicatat suami, karena keduanya menyentuh buku pribadinya (§2.2); pada penggantian, baris yang mengurangi `TUN-I` bertanda belum dikonfirmasi sampai istri, pemegangnya, mengonfirmasi (§3.3, sisi pembayar).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 0 | 14389000 | 290000 | 210000 |
| Δ disetujui (otorisasi 100000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 0 | 14389000 | 290000 | 210000 |
| Δ talangan 75000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 75000 | -75000 | -75000 | 0 |
| Sesudah | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 75000 | 14314000 | 215000 | 210000 |
| Δ penyelesaian (sisa 25000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 75000 | 14314000 | 215000 | 210000 |
| Δ penggantian 75000 | 0 | 0 | -75000 | 0 | 0 | -75000 | 0 | 0 | 0 | 0 | -75000 | 0 | 75000 | 0 |
| Sesudah | 399000 | 350000 | 3265000 | 300000 | 5500000 | 9814000 | 0 | 12000000 | 7500000 | 0 | 0 | 14314000 | 290000 | 210000 |

"Benar" di situasi C: hak ganti sebesar realisasi 75000, bukan sebesar batas 100000; sisa 25000 tidak pernah ada sebagai uang dan hanya dilepas dari terikat. Menguji: INV-08, INV-12, INV-14.

#### CTH-24 — Kelebihan

Istri mendapat otorisasi 200000 untuk hadiah ulang tahun keponakan (kategori "hadiah") dan membayar 230000 dari `TUN-I`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 350000 | 3265000 | 300000 | 5500000 | 9814000 | 0 | 12000000 | 7500000 | 0 | 0 | 14314000 | 290000 | 210000 |
| Δ beli hadiah 230000 | 0 | 0 | -230000 | 0 | 0 | -230000 | 0 | 0 | 0 | 0 | 0 | -230000 | 0 | 0 |
| Sesudah | 399000 | 350000 | 3035000 | 300000 | 5500000 | 9584000 | 0 | 12000000 | 7500000 | 0 | 0 | 14084000 | 290000 | 210000 |

Realisasi 230000; batas 200000; kelebihan 30000; terikat max(0, 200000 − 230000) = 0. Pencatatan tidak ditolak dan tidak ada entri tambahan. Kelebihan 30000 dibayar dari dana bersama, jadi ia pengeluaran dana bersama biasa tanpa klaim; permintaan menampilkan "melebihi 30000" (kait `authorization.overrun`). **Tanda** (§3.8 "Tanda atas entri"; ambang contoh 500000): bagian dalam batas, min(230000, 200000 − 0) = 200000, tidak ditandai; bagian di atas batas, `N` = 230000 − 200000 = 30000, dibandingkan dengan ambang: 30000 tidak di atas 500000, jadi entri ini **tidak ditandai**. Varian: bila ia membayar 800000 atas otorisasi yang sama, `N` = 800000 − 200000 = 600000 > 500000, dan entrinya ditandai dengan dasar `threshold`; saldo dan realisasinya tetap dicatat penuh. Menguji: INV-14.

### 7.7 Keluarga besar, utang, dan pinjam kembali

#### CTH-25 — Pinjaman kepada adik suami, pelunasan sebagian, penghapusan sebagian

Istri menyerahkan 500000 dari `TUN-I` kepada adik suami sebagai pinjaman (`loan_given`; tidak di atas ambang 500000, jadi tanpa permintaan). Adik membayar 200000 kepada suami (`repayment_in` ke `TUN-S`). Kemudian suami mengikhlaskan 100000 (`write_off`, dengan alasan).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 350000 | 3035000 | 300000 | 5500000 | 9584000 | 0 | 12000000 | 7500000 | 0 | 0 | 14084000 | 290000 | 210000 |
| Δ pinjaman diberikan | 0 | 0 | -500000 | 0 | 0 | -500000 | 500000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 350000 | 2535000 | 300000 | 5500000 | 9084000 | 500000 | 12000000 | 7500000 | 0 | 0 | 14084000 | 290000 | 210000 |
| Δ adik membayar 200000 | 0 | 200000 | 0 | 0 | 0 | 200000 | -200000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 550000 | 2535000 | 300000 | 5500000 | 9284000 | 300000 | 12000000 | 7500000 | 0 | 0 | 14084000 | 290000 | 210000 |
| Δ 100000 diikhlaskan | 0 | 0 | 0 | 0 | 0 | 0 | -100000 | 0 | 0 | 0 | 0 | -100000 | 0 | 0 |
| Sesudah | 399000 | 550000 | 2535000 | 300000 | 5500000 | 9284000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13984000 | 290000 | 210000 |

Pinjaman dan pelunasannya bukan pengeluaran dan bukan pemasukan. Hanya penghapusan yang menurunkan kekayaan bersih, sebagai baris "piutang diikhlaskan". Sisa piutang 200000. Penilaian tanda penghapusan (putaran S3; §3.7, §3.8): jumlah pemakaiannya 100000, tidak di atas ambang 500000, jadi tidak ditandai; penghapusan seluruh sisa 200000 sesudahnya juga tidak, karena yang dibandingkan adalah jumlah kumulatif atas piutang ini, 100000 + 200000 = 300000, tidak di atas 500000 (putaran perbaikan S3; SP-51); penghapusan 600000 atas piutang sebesar itu akan ditandai dengan dasar `threshold`. Tidak ada angka yang berubah. Menguji: INV-02, INV-13, INV-34.

**Varian SP-51: mengikhlaskan bertahap** (berdiri sendiri; putaran perbaikan S3; semua angka **[Asumsi]**). Aturan: §3.7, §3.8 `H`, INV-34. Suami meminjamkan 700000 kepada adiknya dari `TUN-S` lewat permintaan yang disetujui (di atas ambang 500000), lalu mengikhlaskannya dalam tiga kali, tanpa permintaan. Kolom: `TUN-S` uang tunai bersama di suami; `UANG = TUN-S`; `PIU` piutang kepada adik; `KB = UANG + PIU`.

| Baris | TUN-S | UANG | PIU | KB |
|---|---|---|---|---|
| Sebelum | 1000000 | 1000000 | 0 | 1000000 |
| Δ pinjaman 700000, tertaut otorisasi berbatas 700000 (`loan_given`) | -700000 | -700000 | 700000 | 0 |
| Sesudah | 300000 | 300000 | 700000 | 1000000 |
| Δ diikhlaskan 300000 (`write_off`); `H` 0; 0 + 300000 = 300000, tidak di atas 500000: tidak ditandai | 0 | 0 | -300000 | -300000 |
| Sesudah | 300000 | 300000 | 400000 | 700000 |
| Δ diikhlaskan 300000 (`write_off`); `H` 300000; 300000 + 300000 = 600000 > 500000: ditandai | 0 | 0 | -300000 | -300000 |
| Sesudah | 300000 | 300000 | 100000 | 400000 |
| Δ diikhlaskan 50000 (`write_off`); `H` 600000; 600000 + 50000 = 650000 > 500000: ditandai | 0 | 0 | -50000 | -50000 |
| Sesudah | 300000 | 300000 | 50000 | 350000 |

Pemeriksaan: `KB` akhir 300000 + 50000 = 350000 = 1000000 − 300000 − 300000 − 50000. `UANG` tidak berubah oleh ketiga penghapusan; hanya `PIU` dan `KB`. Dengan aturan per entri yang lama, ketiganya tidak ditandai (300000, 300000, dan 50000 masing-masing tidak di atas 500000), padahal 650000 diikhlaskan tanpa persetujuan. **Bila penghapusan kedua diajukan dulu** dan ditautkan ke otorisasi berbatas 300000: `N` = 300000 − min(300000, max(0, 300000 − 0)) = 0, jadi tidak ditandai dan tidak menambah `H`; penghapusan ketiga lalu dinilai 300000 + 0 + 50000 = 350000, tidak di atas 500000, tidak ditandai. Angka saldo sama pada kedua jalan. Tiga baris "piutang diikhlaskan" tampil di ringkasan periode tanggal kejadian masing-masing. Menguji: INV-02, INV-13, INV-34.

#### CTH-26 — Bantuan tidak rutin kepada ayah suami

Suami memberi 200000 dari `TUN-S` kepada ayahnya dan memilih "bantuan", bukan pinjaman: `expense`, kategori "kiriman keluarga".

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 550000 | 2535000 | 300000 | 5500000 | 9284000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13984000 | 290000 | 210000 |
| Δ bantuan ayah | 0 | -200000 | 0 | 0 | 0 | -200000 | 0 | 0 | 0 | 0 | 0 | -200000 | 0 | 0 |
| Sesudah | 399000 | 350000 | 2535000 | 300000 | 5500000 | 9084000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13784000 | 290000 | 210000 |

#### CTH-27 — Pinjam kembali

Suami mengambil lagi 300000 dari uang bersama yang dipegang istri: `transfer` (serah terima) `TUN-I` → `TUN-S`. Bukan utang.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 350000 | 2535000 | 300000 | 5500000 | 9084000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13784000 | 290000 | 210000 |
| Δ pinjam kembali | 0 | 300000 | -300000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 650000 | 2235000 | 300000 | 5500000 | 9084000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13784000 | 290000 | 210000 |

Tidak ada kewajiban atau piutang yang lahir; yang berubah hanya siapa memegang. Menguji: INV-06.

#### CTH-28 — Utang rumah tangga dan pelunasannya

Rumah tangga meminjam 1000000 dari kakak istri, diterima istri ke `TUN-I` (`loan_received`); lalu istri membayar kembali 400000 (`repayment_out`).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 650000 | 2235000 | 300000 | 5500000 | 9084000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13784000 | 290000 | 210000 |
| Δ pinjaman diterima | 0 | 0 | 1000000 | 0 | 0 | 1000000 | 0 | 0 | 0 | 1000000 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 650000 | 3235000 | 300000 | 5500000 | 10084000 | 200000 | 12000000 | 7500000 | 1000000 | 0 | 13784000 | 290000 | 210000 |
| Δ pelunasan 400000 | 0 | 0 | -400000 | 0 | 0 | -400000 | 0 | 0 | 0 | -400000 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 650000 | 2835000 | 300000 | 5500000 | 9684000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13784000 | 290000 | 210000 |

Uang pinjaman menaikkan `UANG` tanpa menjadi pemasukan; pelunasan menurunkannya tanpa menjadi pengeluaran. Menguji: INV-02, INV-13.

### 7.8 Rekonsiliasi, koreksi, bersamaan, kiriman ganda

#### CTH-29 — Rekonsiliasi kas dan entri terlambat

Tanggal 10 November istri menghitung `TUN-I`: terhitung 2750000, tercatat 2835000, selisih −85000. Tanggal 12 November ia teringat membeli sayur 60000 pada 8 November dan mencatatnya; aplikasi menawarkannya sebagai bagian dari selisih itu (`recategorization`, kategori "belanja bulanan", tanggal kejadian 8 November). Tanggal 13 November suami menghitung `TUN-S`: terhitung 660000, tercatat 650000, selisih +10000.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 650000 | 2835000 | 300000 | 5500000 | 9684000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13784000 | 290000 | 210000 |
| Δ selisih kas TUN-I | 0 | 0 | -85000 | 0 | 0 | -85000 | 0 | 0 | 0 | 0 | 0 | -85000 | 0 | 0 |
| Sesudah | 399000 | 650000 | 2750000 | 300000 | 5500000 | 9599000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13699000 | 290000 | 210000 |
| Δ entri terlambat sayur 60000 (menutup selisih) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 650000 | 2750000 | 300000 | 5500000 | 9599000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13699000 | 290000 | 210000 |
| Δ selisih kas TUN-S | 0 | 10000 | 0 | 0 | 0 | 10000 | 0 | 0 | 0 | 0 | 0 | 10000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2750000 | 300000 | 5500000 | 9609000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13709000 | 290000 | 210000 |

| Rekonsiliasi | Terhitung | Tercatat | Selisih | Dirinci | Belum dirinci |
|---|---|---|---|---|---|
| TUN-I, sesudah dihitung | 2750000 | 2835000 | -85000 | 0 | 85000 |
| TUN-I, sesudah entri terlambat | 2750000 | 2835000 | -85000 | 60000 | 25000 |
| TUN-S, sesudah dihitung | 660000 | 650000 | 10000 | 0 | 10000 |

Entri terlambat berisi dua baris sebab: `cash_difference` +60000 dan `expense` −60000; jumlahnya 0, jadi tidak ada saldo yang berubah dan uangnya tidak terhitung dua kali. "Belanja bulanan": terpakai 480000 + 60000 = 540000, sisa 1960000. Selisih positif 10000 bukan pemasukan. Tidak ada entri lama yang disunting. Menguji: INV-02, INV-09, INV-16.

#### CTH-30 — Catatan salah dan koreksinya

Istri mencatat parkir 50000 dari `TUN-I`, padahal 5000. Ia memperbaikinya: entri pembalik, lalu entri yang benar (kategori "transportasi").

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 660000 | 2750000 | 300000 | 5500000 | 9609000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13709000 | 290000 | 210000 |
| Δ parkir, salah (50000) | 0 | 0 | -50000 | 0 | 0 | -50000 | 0 | 0 | 0 | 0 | 0 | -50000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2700000 | 300000 | 5500000 | 9559000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13659000 | 290000 | 210000 |
| Δ entri pembalik | 0 | 0 | 50000 | 0 | 0 | 50000 | 0 | 0 | 0 | 0 | 0 | 50000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2750000 | 300000 | 5500000 | 9609000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13709000 | 290000 | 210000 |
| Δ parkir, benar (5000) | 0 | 0 | -5000 | 0 | 0 | -5000 | 0 | 0 | 0 | 0 | 0 | -5000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2745000 | 300000 | 5500000 | 9604000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13704000 | 290000 | 210000 |

Ketiga entri tetap ada di jejak audit. "Transportasi" terpakai: 50000 − 50000 + 5000 = 5000. Hasil yang sama dapat dicapai dengan satu entri penyesuaian +45000. Menguji: INV-09, INV-18.

#### CTH-31 — Pembaruan bersamaan

Pada saat yang sama, istri mencatat ojek 20000 dari `TUN-I` (kategori "transportasi") dan suami mencatat serah terima 100000 dari `TUN-S` ke `TUN-I`. Kedua perangkat membaca `TUN-I` = 2745000.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 660000 | 2745000 | 300000 | 5500000 | 9604000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13704000 | 290000 | 210000 |
| Δ ojek (istri) | 0 | 0 | -20000 | 0 | 0 | -20000 | 0 | 0 | 0 | 0 | 0 | -20000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2725000 | 300000 | 5500000 | 9584000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13684000 | 290000 | 210000 |
| Δ serah terima (suami) | 0 | -100000 | 100000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 560000 | 2825000 | 300000 | 5500000 | 9584000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13684000 | 290000 | 210000 |

Hasil yang benar untuk `TUN-I` adalah 2745000 − 20000 + 100000 = 2825000, apa pun urutan tibanya. Hasil 2725000 atau 2845000 berarti satu entri hilang (lost update) dan merupakan cacat. Model menjaminnya karena entri hanya ditambahkan dan saldo adalah jumlah baris, bukan nilai yang ditimpa klien. Menguji: INV-04, INV-23.

#### CTH-32 — Kiriman ganda

Suami mencatat bensin 40000 dari `TUN-S` (kategori "transportasi"). Jaringan lambat; ia mengetuk "simpan" dua kali, dan aplikasi mengirim ulang dengan `idempotency_key` yang sama.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 560000 | 2825000 | 300000 | 5500000 | 9584000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13684000 | 290000 | 210000 |
| Δ bensin, kiriman pertama | 0 | -40000 | 0 | 0 | 0 | -40000 | 0 | 0 | 0 | 0 | 0 | -40000 | 0 | 0 |
| Sesudah | 399000 | 520000 | 2825000 | 300000 | 5500000 | 9544000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13644000 | 290000 | 210000 |
| Δ bensin, kiriman kedua (kunci sama) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 520000 | 2825000 | 300000 | 5500000 | 9544000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13644000 | 290000 | 210000 |

Jumlah entri bertambah satu, bukan dua; kiriman kedua mengembalikan entri pertama. Menguji: INV-10.

### 7.9 Aset, cicilan baru, penyisihan, hadiah

#### CTH-33 — Pembelian aset dan penilaian ulang

Istri membeli laptop bekas 2000000 dari `TUN-I` (`asset_purchase`; di atas ambang, jadi diajukan lebih dulu sebagai permintaan). Kemudian nilai motor diperbarui dari 12000000 menjadi 11000000 (`asset_valuation` −1000000).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 520000 | 2825000 | 300000 | 5500000 | 9544000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13644000 | 290000 | 210000 |
| Δ beli laptop | 0 | 0 | -2000000 | 0 | 0 | -2000000 | 0 | 2000000 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 520000 | 825000 | 300000 | 5500000 | 7544000 | 200000 | 14000000 | 7500000 | 600000 | 0 | 13644000 | 290000 | 210000 |
| Δ nilai motor turun | 0 | 0 | 0 | 0 | 0 | 0 | 0 | -1000000 | 0 | 0 | 0 | -1000000 | 0 | 0 |
| Sesudah | 399000 | 520000 | 825000 | 300000 | 5500000 | 7544000 | 200000 | 13000000 | 7500000 | 600000 | 0 | 12644000 | 290000 | 210000 |

Pembelian aset tidak mengubah kekayaan bersih: uang berubah bentuk menjadi barang. Ia tampil di arus kas sebagai "beli aset". Menguji: INV-02, INV-07.

**Yang tidak ditampilkan rantai [Asumsi].** "Diajukan lebih dulu" berarti: sebelum membeli, istri mengajukan permintaan; suami menyetujuinya; lahir otorisasi atas dana bersama dengan batas tidak kurang dari 2000000, tanpa tempat uang sumber, jadi tanpa reservasi; entri pembelian ditautkan kepadanya, sehingga realisasinya 2000000, seluruhnya dalam batas otorisasi (`N` = 0), dan entri itu tidak ditandai (§3.8 "Tanda atas entri"); lalu otorisasinya diselesaikan pada situasi A, tanpa entri (seperti CTH-21). Rantai ini tidak memuat baris persetujuan dan penyelesaian karena keduanya tidak membuat entri dan tidak mengubah kolom mana pun (INV-12); bila ditambahkan, barisnya nol seluruhnya. Sesudah penyelesaian *terikat*-nya nol, jadi CTH-37 tidak tersentuh. Bila istri membeli tanpa mengajukan, entri yang sama tersimpan dengan angka yang sama dan membawa tanda netral di atas ambang (dokumen 06 §6.2; dokumen 17 S2-2 pilihan A, diterima pemilik produk di checkpoint S2 atas saran lead).

#### CTH-34 — Pembelian aset secara cicilan dan biaya kredit

Kulkas seharga 4000000 dibeli dengan uang muka 400000 dari `TUN-I` dan 12 cicilan × 330000 = 3960000. Biaya kredit = 400000 + 3960000 − 4000000 = 360000. Satu entri `asset_purchase`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 520000 | 825000 | 300000 | 5500000 | 7544000 | 200000 | 13000000 | 7500000 | 600000 | 0 | 12644000 | 290000 | 210000 |
| Δ beli kulkas secara cicilan | 0 | 0 | -400000 | 0 | 0 | -400000 | 0 | 4000000 | 3960000 | 0 | 0 | -360000 | 0 | 0 |
| Sesudah | 399000 | 520000 | 425000 | 300000 | 5500000 | 7144000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12284000 | 290000 | 210000 |

Keseimbangan: −400000 + 4000000 − 3960000 = −360000 = `credit_cost`. `KWJ-C` 11460000 = motor 7500000 + kulkas 3960000. Kekayaan bersih turun sebesar biaya kredit pada tanggal pembelian, dengan baris sebab yang menjelaskannya (menutup K-11). Jumlah pemakaian entri ini, yang dibandingkan dengan ambang permintaan dan menjadi realisasi otorisasinya, adalah 400000 + 3960000 = 4360000 (§3.8): di atas ambang contoh 500000, jadi pembelian ini diajukan lebih dulu, dengan andaian yang sama dengan CTH-33 (otorisasi tanpa reservasi, batas tidak kurang dari jumlah pemakaian itu, diselesaikan tanpa entri; tidak ditampilkan rantai). Dengan batas itu seluruh 4360000 berada dalam batas otorisasi (`N` = 0), jadi entri ini tidak ditandai. Dua belas pembayaran 330000 berikutnya masing-masing seperti CTH-04: terhitung terhadap anggaran "cicilan", tanpa mengubah kekayaan bersih. Menguji: INV-02, INV-07.

#### CTH-35 — Penyisihan ke dana pribadi

Uang saku bulanan yang disepakati: 100000 untuk masing-masing. Suami memindahkan 100000 dari `TUN-S` ke `PRI-S`; istri 100000 dari `TUN-I` ke `PRI-I`. Dua entri `owner_transfer`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 520000 | 425000 | 300000 | 5500000 | 7144000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12284000 | 290000 | 210000 |
| Δ uang saku suami | 0 | -100000 | 0 | 0 | 0 | -100000 | 0 | 0 | 0 | 0 | 0 | -100000 | 100000 | 0 |
| Sesudah | 399000 | 420000 | 425000 | 300000 | 5500000 | 7044000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12184000 | 390000 | 210000 |
| Δ uang saku istri | 0 | 0 | -100000 | 0 | 0 | -100000 | 0 | 0 | 0 | 0 | 0 | -100000 | 0 | 100000 |
| Sesudah | 399000 | 420000 | 325000 | 300000 | 5500000 | 6944000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12084000 | 390000 | 310000 |

Keduanya melihat kedua entri ini (kebocoran B1); tidak seorang pun melihat saldo pribadi pasangannya sesudahnya. Tanda: tidak ada pada kedua entri, dan itu tidak berubah oleh keputusan checkpoint S2. Dengan jatah pribadi 100000 masing-masing (**[Asumsi]** contoh), tiap pengambilan adalah yang pertama di periodenya: 0 + 100000 tidak melebihi 100000. Tanpa jatah, 100000 tidak di atas ambang 500000. Pengambilan kedua dalam periode yang sama, berapa pun jumlahnya, akan ditandai bila jatah 100000 itu berlaku (§3.4). Menguji: INV-02, INV-03, INV-30.

#### CTH-36 — Hadiah untuk satu orang

Istri menerima hadiah 500000 dari ibunya dan mencatatnya ke `PRI-I`: `income` di buku pribadinya, sumber `gift` (bawaan: dana pribadi).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 420000 | 325000 | 300000 | 5500000 | 6944000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12084000 | 390000 | 310000 |
| Δ hadiah untuk istri | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 500000 |
| Sesudah | 399000 | 420000 | 325000 | 300000 | 5500000 | 6944000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12084000 | 390000 | 810000 |

Suami tidak melihat entri ini (keputusan pemilik produk, checkpoint S2, S2-1 butir (ii); §3.1) dan angka rumah tangga tidak berubah. Menguji: INV-20.

### 7.10 CTH-37 — Ringkasan bulanan rumah tangga contoh

Dihitung dari CTH-03 sampai CTH-36 (CTH-01 adalah saldo awal dan tidak masuk periode). Semua angka dari buku rumah tangga.

**Custody akhir** (kriteria sukses nomor 1).

| Pemegang | Tempat uang | Saldo |
|---|---|---|
| Suami | BNK-S | 399000 |
| Suami | TUN-S | 420000 |
| Istri | TUN-I | 325000 |
| Istri | EWL-I | 300000 |
| Istri | TAB-I | 5500000 |
| Jumlah = UANG | | 6944000 |

Suami memegang 819000; istri memegang 6125000; 819000 + 6125000 = 6944000.

**Tabel A — mengapa harta bersih berubah.**

| Baris | Jumlah |
|---|---|
| Pemasukan (gaji, atas nama suami) | 10000000 |
| Belanja semua kategori (lihat tabel B, tanpa cicilan) | -5461000 |
| Biaya kredit | -360000 |
| Selisih hitung negatif yang belum dirinci | -25000 |
| Selisih hitung positif | 10000 |
| Piutang diikhlaskan | -100000 |
| Perubahan nilai aset | -1000000 |
| Kontribusi dari dana pribadi (suami 80000, istri 40000) | 120000 |
| Disisihkan ke dana pribadi (suami 100000, istri 100000) | -200000 |
| Jumlah = perubahan KB | 2984000 |

`KB` awal 9100000 (CTH-01) + 2984000 = 12084000 = `KB` akhir (CTH-36).

**Tabel B — uang keluar per kategori.**

| Kategori | Rincian | Terpakai |
|---|---|---|
| Tagihan rutin | 450000 + 350000 + 200000 | 1000000 |
| Kiriman keluarga | 3000000 + 200000 | 3200000 |
| Belanja bulanan | 150000 + 80000 + 60000 + 40000 + 150000 + 60000 | 540000 |
| Perawatan diri | 50000 | 50000 |
| Biaya admin | 1000 | 1000 |
| Perlengkapan rumah | 120000 − 20000 + 100000 − 100000 + 50000 | 150000 |
| Jajan dan makan di luar | 75000 + 75000 | 150000 |
| Kesehatan | 75000 | 75000 |
| Hadiah | 230000 | 230000 |
| Transportasi | 5000 + 20000 + 40000 | 65000 |
| Jumlah belanja | | 5461000 |
| Cicilan | 1500000 | 1500000 |
| Jumlah tabel B | | 6961000 |

**Tabel C — siapa membayar apa.**

| Pembayar | Dari dana bersama | Dana pribadi, kontribusi | Dana pribadi, diganti | Jumlah |
|---|---|---|---|---|
| Suami | 5816000 | 80000 | 135000 | 6031000 |
| Istri | 890000 | 40000 | 0 | 930000 |
| Jumlah | 6706000 | 120000 | 135000 | 6961000 |

Suami, dana bersama: 1000000 (tagihan) + 3000000 (ibu) + 1000 (biaya admin) + 75000 (CTH-21) + 200000 (ayah) + 40000 (bensin) + 1500000 (cicilan) = 5816000. Istri, dana bersama: 150000 + 50000 + 100000 (CTH-19, sesudah pengembalian) + 200000 + 75000 + 230000 + 60000 (sayur) + 5000 + 20000 = 890000. Jumlah tabel C = jumlah tabel B. Angka ini mencerminkan siapa yang memegang uang untuk pos apa, bukan siapa yang lebih boros; lihat §10.3.

**Anggaran periode.**

| Kategori | Batas | Terpakai | Terikat | Sisa |
|---|---|---|---|---|
| Belanja bulanan | 2500000 | 540000 | 0 | 1960000 |
| Cicilan | 1500000 | 1500000 | 0 | 0 |
| Tagihan rutin | 1100000 | 1000000 | 0 | 100000 |
| Jajan dan makan di luar | 500000 | 150000 | 0 | 350000 |

**Arus kas.**

| Baris | Jumlah |
|---|---|
| Masuk: pemasukan | 10000000 |
| Masuk: pinjaman diterima | 1000000 |
| Masuk: piutang dibayar | 200000 |
| Masuk: selisih hitung positif | 10000 |
| Jumlah masuk | 11210000 |
| Keluar: belanja dari dana bersama | -5146000 |
| Keluar: cicilan dan utang dibayar (1500000 + 400000) | -1900000 |
| Keluar: penggantian dibayar (60000 + 75000) | -135000 |
| Keluar: pinjaman diberikan | -500000 |
| Keluar: beli aset (2000000 + 400000) | -2400000 |
| Keluar: disisihkan ke dana pribadi | -200000 |
| Keluar: selisih hitung negatif | -85000 |
| Jumlah keluar | -10366000 |
| Perubahan UANG | 844000 |

`UANG` awal 6100000 + 844000 = 6944000 = `UANG` akhir. Belanja dari dana bersama 5146000 = belanja tabel B 5461000 − kontribusi 120000 − talangan 135000 − yang dirinci dari selisih 60000.

**Tabungan, aset, piutang, kewajiban.**

| Posisi | Awal periode | Akhir periode |
|---|---|---|
| Tabungan "dana darurat" (target 10000000) | 5000000 | 5500000 |
| Aset: motor | 12000000 | 11000000 |
| Aset: laptop | 0 | 2000000 |
| Aset: kulkas | 0 | 4000000 |
| Piutang: adik suami | 0 | 200000 |
| Cicilan motor (sisa 5 kali) | 9000000 | 7500000 |
| Cicilan kulkas (sisa 12 kali) | 0 | 3960000 |
| Utang kepada kakak istri | 0 | 600000 |

Aset 11000000 + 2000000 + 4000000 = 17000000. `KB` akhir = 6944000 + 200000 + 17000000 − 11460000 − 600000 − 0 = 12084000.

**Pemeriksaan data penuh** (tidak ditampilkan kepada siapa pun). Kekayaan bersih per pemilik: `shared` 12084000; suami 390000; istri 810000; jumlah 13284000. Menurut INV-08: semua tempat uang (6944000 + 390000 + 810000 = 8144000) + piutang kepada pihak luar 200000 + aset 17000000 − kewajiban kepada pihak luar (11460000 + 600000 = 12060000) = 13284000.

Menguji: INV-05, INV-07, INV-08, INV-15, INV-19.

### 7.11 CTH-38 — Klaim antar-pemilik ditutup saat anggota keluar (berdiri sendiri)

Keputusan sementara pemilik produk, checkpoint S2 (S2-8; §8.2). Tidak mengubah keadaan berantai dan tidak masuk CTH-37; semua angkanya **[Asumsi]**. Anggota M punya dua klaim terbuka dengan dana bersama saat ia keluar: hak ganti 60000 (rumah tangga berutang kepadanya) dan pinjaman 200000 dari dana bersama (ia berutang kepada rumah tangga).

Kolom: `UANG` = uang yang ada; `PIU-A` = piutang rumah tangga kepada M; `KWJ-G` = hak ganti yang belum dibayar kepada M; `KB = UANG + PIU-A − KWJ-G`; `PRI-M` = uang pribadi M; `KLAIM-M` = hak ganti milik M (piutangnya); `UTG-M` = utang M kepada dana bersama; `KB-M = PRI-M + KLAIM-M − UTG-M`.

| Baris | UANG | PIU-A | KWJ-G | KB | PRI-M | KLAIM-M | UTG-M | KB-M |
|---|---|---|---|---|---|---|---|---|
| Sebelum | 5000000 | 0 | 0 | 5000000 | 300000 | 0 | 0 | 300000 |
| Δ M menalangi 60000, minta diganti (`advance`) | 0 | 0 | 60000 | -60000 | -60000 | 60000 | 0 | 0 |
| Sesudah | 5000000 | 0 | 60000 | 4940000 | 240000 | 60000 | 0 | 300000 |
| Δ dana bersama meminjamkan 200000 kepada M (`loan_given`, antar-pemilik) | -200000 | 200000 | 0 | 0 | 200000 | 0 | 200000 | 0 |
| Sesudah | 4800000 | 200000 | 60000 | 4940000 | 440000 | 60000 | 200000 | 300000 |
| Δ hari D: M keluar; aksesnya berakhir, masa tenggang 7 hari mulai (peristiwa, tanpa entri) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 4800000 | 200000 | 60000 | 4940000 | 440000 | 60000 | 200000 | 300000 |
| Δ hari D + 7: masa tenggang lewat tanpa dibatalkan; `exit_closure` oleh sistem | 0 | -200000 | -60000 | -140000 | 0 | -60000 | -200000 | 140000 |
| Sesudah | 4800000 | 0 | 0 | 4800000 | 440000 | 0 | 0 | 440000 |

Baris entri `exit_closure`:

| Buku | Baris | Nilai |
|---|---|---|
| Rumah tangga | Kewajiban hak ganti kepada M | -60000 |
| Rumah tangga | Sebab `owner_transfer` (hak ganti ditutup) | 60000 |
| Rumah tangga | Piutang kepada M | -200000 |
| Rumah tangga | Sebab `owner_transfer` (pinjaman ditutup) | -200000 |
| Pribadi M | Piutang hak ganti | -60000 |
| Pribadi M | Sebab `owner_transfer` (hak ganti ditutup) | -60000 |
| Pribadi M | Kewajiban kepada dana bersama | -200000 |
| Pribadi M | Sebab `owner_transfer` (pinjaman ditutup) | 200000 |

Keseimbangan buku rumah tangga (piutang − kewajiban = sebab): −200000 − (−60000) = −140000 = 60000 − 200000. Buku pribadi M: −60000 − (−200000) = 140000 = −60000 + 200000. Jumlah `owner_transfer` di semua buku: 60000 − 200000 − 60000 + 200000 = 0. `UANG` dan `PRI-M` tidak berubah: entri ini tidak menyatakan pembayaran apa pun. Kekayaan bersih rumah tangga turun 140000 dan tampil di ringkasan sebagai "klaim ditutup karena anggota keluar": hak ganti 60000, pinjaman −200000. Data penuh sebelum dan sesudah penutupan: 4940000 + 300000 = 5240000 = 4800000 + 440000. Sesudah entri ini buku pribadi M ditutup, lalu dihapus seketika bila M memintanya, atau pada akhir jangka retensi bila tidak (§8.2 butir 2a); angkanya sama pada keduanya. **Masa tenggang** (putaran perbaikan S3; SP-48; §8.2 butir 2b): dari hari D sampai D + 7 kedua klaim tetap terbuka dan tampil di buku rumah tangga (`PIU-A` 200000, `KWJ-G` 60000, `KB` 4940000), karena tidak seorang pun dapat menulis buku M selama itu; angka `exit_closure` karenanya sama dengan bila ditulis pada hari D: −200000 + 60000 = −140000. Yang berubah hanya tanggalnya: tanggal kejadian entri itu hari D + 7, jadi baris "klaim ditutup karena anggota keluar" tampil di ringkasan periode yang memuat D + 7. Bila M membatalkan keluarnya sebelum D + 7, baris terakhir tidak ada dan keadaan tetap seperti baris "Sesudah" hari D. Yang tersisa di buku rumah tangga adalah `UANG` 4800000 dan `KB` 4800000, tanpa klaim terbuka kepada orang yang sudah pergi. Menguji: INV-02, INV-03, INV-08, INV-25, INV-28.

### 7.12 CTH-39 — Satu periode dengan beberapa penghasilan dan jatah pribadi (berdiri sendiri)

Menurut keputusan dan pernyataan pemilik produk, 11 November 2026 (§3.1, §3.4). Tidak mengubah keadaan berantai dan tidak masuk CTH-37; semua angkanya **[Asumsi]** kecuali jumlah gaji 10000000 dan kiriman ibu 3000000 (brief §1.1). Kiriman kepada ibu di sini ilustrasi, bukan pembayaran rutin rumah tangga pilot (§7.1). **Revisi pasca-checkpoint S2: tidak ada angka saldo yang berubah; yang berubah hanya dasar tanda** (di bawah tabel jatah). Hari mulai periode 1; periode 1–31 Oktober 2026. Ambang permintaan 500000. Jatah pribadi suami yang disepakati berdua: 1000000 per periode. Semua entri dicatat suami; ia `earner` ketiga penghasilan.

Kolom: `BNK-S` rekening bank suami; `EWL-S` dompet digital suami; `TUN-S` uang tunai bersama di suami; `BNK-I` rekening bank istri; keempatnya berpemilik `shared`. `UANG = BNK-S + EWL-S + TUN-S + BNK-I`. `KB` = kekayaan bersih rumah tangga = `UANG` (tidak ada piutang, aset, atau kewajiban). `PRI-S` = tempat uang pribadi suami (tidak terlihat istri).

| Baris | BNK-S | EWL-S | TUN-S | BNK-I | UANG | KB | PRI-S |
|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 100000 | 200000 | 300000 | 1100000 | 1100000 | 400000 |
| Δ 1 Okt: gaji termin pertama 6000000, 600000 ke pribadi (`income`) | 5400000 | 0 | 0 | 0 | 5400000 | 5400000 | 600000 |
| Sesudah | 5900000 | 100000 | 200000 | 300000 | 6500000 | 6500000 | 1000000 |
| Δ 3 Okt: kiriman kepada ibu 3000000 (`expense`) | -3000000 | 0 | 0 | 0 | -3000000 | -3000000 | 0 |
| Sesudah | 2900000 | 100000 | 200000 | 300000 | 3500000 | 3500000 | 1000000 |
| Δ 10 Okt: penghasilan lepas 1500000 ke dompet digital (`income`, sumber bawaan honor atau pekerjaan lepas, `freelance`, kelompok `work`; dahulu `business`) | 0 | 1500000 | 0 | 0 | 1500000 | 1500000 | 0 |
| Sesudah | 2900000 | 1600000 | 200000 | 300000 | 5000000 | 5000000 | 1000000 |
| Δ 15 Okt: gaji termin kedua 4000000, 1000000 ke pribadi (`income`) | 3000000 | 0 | 0 | 0 | 3000000 | 3000000 | 1000000 |
| Sesudah | 5900000 | 1600000 | 200000 | 300000 | 8000000 | 8000000 | 2000000 |
| Δ 16 Okt: tarik tunai 2000000, biaya admin 6500 (`transfer`) | -2006500 | 0 | 2000000 | 0 | -6500 | -6500 | 0 |
| Sesudah | 3893500 | 1600000 | 2200000 | 300000 | 7993500 | 7993500 | 2000000 |
| Δ 16 Okt: serah terima 3500000 ke rekening istri (`transfer`) | -3500000 | 0 | 0 | 3500000 | 0 | 0 | 0 |
| Sesudah | 393500 | 1600000 | 2200000 | 3800000 | 7993500 | 7993500 | 2000000 |
| Δ 17 Okt: bantuan kepada adik 500000 dari tunai (`expense`) | 0 | 0 | -500000 | 0 | -500000 | -500000 | 0 |
| Sesudah | 393500 | 1600000 | 1700000 | 3800000 | 7493500 | 7493500 | 2000000 |

Jatah pribadi suami dalam periode ini (§3.4). `Dalam jatah = min(bagian, max(0, jatah − terpakai sebelum))`; `di luar jatah = bagian − dalam jatah`; `terpakai sesudah = terpakai sebelum + bagian`; `terpakai sebelum` baris kedua = `terpakai sesudah` baris pertama.

| Entri | Bagian pribadi | Jatah | Terpakai sebelum | Dalam jatah | Di luar jatah | Terpakai sesudah | Ditandai |
|---|---|---|---|---|---|---|---|
| Gaji termin pertama | 600000 | 1000000 | 0 | 600000 | 0 | 600000 | Tidak |
| Gaji termin kedua | 1000000 | 1000000 | 600000 | 400000 | 600000 | 1600000 | Ya |

Bagian pertama: 0 + 600000 = 600000, tidak melebihi jatah 1000000; tidak ditandai dan tidak ada pemberitahuan, walaupun 600000 di atas ambang 500000 (default yang diterima di checkpoint S2, dokumen 17 D-48). Bagian kedua: 600000 + 1000000 = 1600000 > 1000000; **ditandai karena jatah terlewati** (keputusan pemilik produk, checkpoint S2, S2-1 butir (iii) pilihan b). Ambang tidak dibandingkan: suami punya jatah yang berlaku. Entri itu tetap tersimpan; tidak ada yang dihalangi.

**Yang berubah dari versi sebelum checkpoint S2.** Angka kedua baris tabel jatah dan seluruh tabel saldo tidak berubah. Hasil tandanya juga sama (termin pertama tidak, termin kedua ya); yang berubah **dasarnya**: semula termin kedua ditandai karena bagian di luar jatah, 600000, di atas ambang 500000; kini karena jumlah periode itu melebihi jatah. Perbedaannya baru tampak bila kelebihannya kecil. Tiga varian termin kedua berikut berdiri sendiri: masing-masing mengganti bagian pribadi termin kedua saja, tidak diteruskan ke tabel saldo, dan `terpakai sebelum`-nya tetap 600000.

| Varian bagian pribadi termin kedua | Dalam jatah | Di luar jatah | Terpakai sesudah | Ditandai (aturan yang berlaku) | Ditandai menurut pilihan (a) yang ditolak |
|---|---|---|---|---|---|
| 400000 | 400000 | 0 | 1000000 | Tidak: 1000000 tidak melebihi 1000000 | Tidak |
| 500000 | 400000 | 100000 | 1100000 | **Ya**: 1100000 > 1000000 | Tidak: 100000 tidak di atas 500000 |
| 1000000 (contoh di atas) | 400000 | 600000 | 1600000 | Ya: 1600000 > 1000000 | Ya: 600000 di atas 500000 |

Bila suami memerlukan 500000 pada termin kedua, jalan tanpa tanda tersedia menurut bacaan lead atas "split jatah + uang pribadi" (**[Usulan]** lead; bacaan itu belum dikukuhkan pemilik produk): 400000 dari jatah, dan 100000 dari uang pribadinya yang sudah ada (`PRI-S`), yang tidak menyentuh jatah; atau ia mengajukan tambahan 100000 sebagai permintaan, atau mengusulkan kenaikan jatah (§3.4). Bagaimana tanda itu ditampilkan dan diberitahukan adalah aturan dokumen 06 §6.2. Jatah tidak mengubah satu angka pun di tabel pertama: tanpa jatah, ketujuh baris `Δ` sama persis.

Keseimbangan gaji termin kedua, buku rumah tangga: 3000000 = 4000000 − 1000000; buku pribadi suami: 1000000 = 1000000; jumlah `owner_transfer`: −1000000 + 1000000 = 0. Tarik tunai: `BNK-S` berkurang 2000000 + 6500 = 2006500, `TUN-S` bertambah 2000000, dan hanya biaya admin 6500 yang menjadi pengeluaran. Serah terima dicatat suami ke tempat uang yang dipegang istri, jadi baris +3500000 pada `BNK-I` bertanda belum dikonfirmasi sampai istri mengonfirmasi (§3.3); saldonya sudah berubah. Bantuan kepada adik 500000 tidak di atas ambang; kiriman kepada ibu 3000000 di atas ambang dan, seperti CTH-06, dijadwalkan sebagai tagihan berulang (§3.15); dengan andaian yang sama dengan CTH-06 (**[Asumsi]**: tercakup kesepakatan rutin 3000000 per kali), ia tidak ditandai, karena 3000000 tidak melebihi 3000000 (dasar `routine_amount`, §3.8). Tanpa kesepakatan yang berlaku ia dinilai dengan aturan biasa dan membawa tanda netral (3000000 > 500000); angkanya tetap.

Ringkasan periode: pemasukan 6000000 + 1500000 + 4000000 = 11500000 (gaji 10000000, hasil usaha 1500000); disisihkan ke dana pribadi suami 600000 + 1000000 = 1600000; belanja 3000000 + 6500 + 500000 = 3506500; perubahan `KB` = 11500000 − 1600000 − 3506500 = 6393500, dan 1100000 + 6393500 = 7493500. Arus kas: pemasukan yang masuk tempat uang dana bersama 5400000 + 1500000 + 3000000 = 9900000; belanja 3506500; perubahan `UANG` = 9900000 − 3506500 = 6393500. Custody akhir: suami memegang 393500 + 1600000 + 1700000 = 3693500; istri memegang 3800000, di antaranya 3500000 belum dikonfirmasi; 3693500 + 3800000 = 7493500. `PRI-S`: 400000 + 1600000 = 2000000. Aturannya simetris: penghasilan istri dicatat istri dengan bentuk yang sama dan dihitung terhadap jatahnya sendiri.

Menguji: INV-02, INV-03, INV-06, INV-19, INV-20, INV-24, INV-27, INV-30.

### 7.13 CTH-40 — Menghitung dompet campuran (berdiri sendiri)

Fakta rumah tangga pilot, checkpoint S2 (S2-6): uang pribadi dan uang bersama dapat bercampur di satu dompet. Aturannya di §2.3 dan §3.10. Tidak mengubah keadaan berantai dan tidak masuk CTH-37; semua angkanya **[Asumsi]**. Tidak ada bawaan: pemegang memilih satu dari tiga pilihan di bawah setiap kali selisihnya bukan nol (diterima pemilik produk atas saran lead, 12 November 2026; T-03). Ketiga tabel menunjukkan hasil tiap pilihan dari keadaan yang sama; angkanya tidak berubah dari revisi sebelumnya.

Satu dompet fisik istri dicatat sebagai dua tempat uang. Kolom: `DMP-B` bagian dana bersama di dompet istri (pemilik `shared`, pemegang istri); `DMP-P` bagian dana pribadi istri di dompet yang sama (pemilik `member(istri)`; tidak terlihat suami); `ISI = DMP-B + DMP-P`, isi dompet menurut catatan, hanya dapat dihitung istri; `BNK-S` rekening bank suami (pemilik `shared`); `UANG = DMP-B + BNK-S`; `KB` = kekayaan bersih rumah tangga = `UANG` (tidak ada piutang, aset, atau kewajiban).

| Baris | DMP-B | DMP-P | ISI | BNK-S | UANG | KB |
|---|---|---|---|---|---|---|
| Sebelum | 800000 | 300000 | 1100000 | 2000000 | 2800000 | 2800000 |
| Δ sayur 120000, keperluan rumah (`expense`, buku rumah tangga) | -120000 | 0 | -120000 | 0 | -120000 | -120000 |
| Sesudah | 680000 | 300000 | 980000 | 2000000 | 2680000 | 2680000 |
| Δ pulsa pribadi 50000 (`expense`, buku pribadi istri) | 0 | -50000 | -50000 | 0 | 0 | 0 |
| Sesudah | 680000 | 250000 | 930000 | 2000000 | 2680000 | 2680000 |

Istri menghitung seluruh isi dompet: `C` = 890000. Tercatat 680000 + 250000 = 930000. Selisih = 890000 − 930000 = −40000. Ia tidak ingat untuk apa.

**Pilihan 1: pemegang menetapkan seluruh selisih pada bagian bersama.** Satu entri `cash_difference` atas `DMP-B`.

| Baris | DMP-B | DMP-P | ISI | BNK-S | UANG | KB |
|---|---|---|---|---|---|---|
| Sebelum (sesudah pulsa) | 680000 | 250000 | 930000 | 2000000 | 2680000 | 2680000 |
| Δ selisih hitung pada bagian bersama | -40000 | 0 | -40000 | 0 | -40000 | -40000 |
| Sesudah | 640000 | 250000 | 890000 | 2000000 | 2640000 | 2640000 |

**Pilihan 2: pemegang menetapkan seluruh selisih pada bagian pribadinya** (dari `Sebelum` yang sama). Satu entri `cash_difference` atas `DMP-P`, di buku pribadi istri.

| Baris | DMP-B | DMP-P | ISI | BNK-S | UANG | KB |
|---|---|---|---|---|---|---|
| Sebelum (sesudah pulsa) | 680000 | 250000 | 930000 | 2000000 | 2680000 | 2680000 |
| Δ selisih hitung pada bagian pribadi | 0 | -40000 | -40000 | 0 | 0 | 0 |
| Sesudah | 680000 | 210000 | 890000 | 2000000 | 2680000 | 2680000 |

**Pilihan 3: dibagi, 25000 pada bagian bersama dan 15000 pada bagian pribadi** (dari `Sebelum` yang sama). Dua entri `cash_difference`, satu per tempat uang.

| Baris | DMP-B | DMP-P | ISI | BNK-S | UANG | KB |
|---|---|---|---|---|---|---|
| Sebelum (sesudah pulsa) | 680000 | 250000 | 930000 | 2000000 | 2680000 | 2680000 |
| Δ selisih hitung pada bagian bersama | -25000 | 0 | -25000 | 0 | -25000 | -25000 |
| Sesudah | 655000 | 250000 | 905000 | 2000000 | 2655000 | 2655000 |
| Δ selisih hitung pada bagian pribadi | 0 | -15000 | -15000 | 0 | 0 | 0 |
| Sesudah | 655000 | 235000 | 890000 | 2000000 | 2655000 | 2655000 |

Rekonsiliasi yang lahir, per tempat uang (`terhitung = tercatat + selisih yang ditetapkan`):

| Pilihan | Tempat uang | Tercatat | Selisih ditetapkan | Terhitung | Entri | Terlihat suami |
|---|---|---|---|---|---|---|
| 1 | DMP-B | 680000 | -40000 | 640000 | `cash_difference` | Ya |
| 1 | DMP-P | 250000 | 0 | 250000 | Tidak ada (selisih 0) | Tidak |
| 2 | DMP-B | 680000 | 0 | 680000 | Tidak ada (selisih 0) | Ya: "dihitung, cocok", terhitung 680000 |
| 2 | DMP-P | 250000 | -40000 | 210000 | `cash_difference` | Tidak |
| 3 | DMP-B | 680000 | -25000 | 655000 | `cash_difference` | Ya |
| 3 | DMP-P | 250000 | -15000 | 235000 | `cash_difference` | Tidak |

Pada ketiga pilihan `ISI` sesudahnya 890000 = `C`: 640000 + 250000; 680000 + 210000; 655000 + 235000. Jumlah kedua selisih selalu −40000: −40000 + 0; 0 − 40000; −25000 − 15000. `UANG` dan `KB` hanya berubah sebesar selisih bagian bersama: −40000, 0, −25000. Server menerima, pada tiap pilihan, dua perintah rekonsiliasi yang tidak terkait: saldo terhitung `DMP-B` (640000, 680000, atau 655000) dan saldo terhitung `DMP-P` (250000, 210000, atau 235000); angka 890000 dan 930000 tidak pernah dikirim. Suami tidak pernah melihat `C`, `DMP-P`, atau `ISI`, dan dari yang ia lihat ketiganya tidak dapat diturunkan (dokumen 06 V-9, INV-A41). Pada pilihan 1 ia melihat "selisih hitung 40000" pada uang bersama yang dipegang istri, terhitung 640000; pada pilihan 2 ia melihat uang bersama itu dihitung dan cocok, 680000; pada pilihan 3 selisih 25000, terhitung 655000. Yang ia lihat sama persis dengan rekonsiliasi `DMP-B` yang berdiri sendiri dengan selisih yang sama.

**Biaya memilih "bersama" ketika uangnya uang pribadi, pada angka ini.** Andaikan 40000 itu sebenarnya jajan pribadi istri yang lupa dicatat. Pada pilihan 1 kekurangan uang pribadinya tampil sebagai selisih hitung uang bersama atas namanya, dan `UANG` tercatat 40000 lebih kecil daripada yang sebenarnya dimiliki rumah tangga. Bila kemudian ia teringat, ia memperbaikinya dengan dua entri: entri pembalik atas selisih itu (`DMP-B` 640000 + 40000 = 680000; `UANG` 2640000 + 40000 = 2680000), lalu belanja pribadi 40000 (`DMP-P` 250000 − 40000 = 210000). Hasilnya sama dengan pilihan 2, dan `ISI` tetap 890000. Bila sebaliknya 40000 itu belanja rumah yang lupa dicatat, pada pilihan 1 ia mencatatnya sebagai entri terlambat yang menutup selisih tanpa mengubah saldo, seperti CTH-29.

Menguji: INV-02, INV-05, INV-16 (per tempat uang), INV-20 (pilihan 2 tidak mengubah angka rumah tangga), INV-25.

### 7.14 CTH-41 — Uang masuk, bayar tagihan, siapkan uang untuk tagihan yang menyusul, bagikan sisanya (berdiri sendiri)

Pola rumah tangga pilot menurut pemilik produk (checkpoint S2, S2-4): begitu penghasilan masuk ia membayar tagihan lebih dulu, lalu membagikan sisanya; tanggal tagihan tersebar sesudah tanggal uang masuk (contohnya 11, lalu 12, 13, 20, 29), dan sebagian tagihan jumlahnya berubah-ubah. Aturannya di §3.15 dan §3.9. Tidak mengubah keadaan berantai dan tidak masuk CTH-37. Tanggal mengikuti contohnya; **semua jumlah dan nama tagihan [Asumsi]**.

Hari mulai periode 11; periode 11 November–10 Desember 2026. Ambang permintaan 500000. Jatah pribadi suami yang disepakati: 500000 per periode. Empat tagihan berulang. Sejak revisi kedua pasca-checkpoint S2 **keempatnya** punya kesepakatan rutin (§3.15); **[Asumsi]** contoh: semuanya disepakati berdua saat dibuat, sebelum periode ini, dengan jumlah per kali `R` di bawah. Air sengaja disepakati sebesar perkiraannya, bukan jumlah tertingginya yang wajar, untuk menunjukkan akibatnya.

| Tagihan | Jatuh tempo | Perkiraan | Disepakati per kali `R` [Asumsi] | Jumlah sebenarnya bulan ini `P` | `P > R`? | Tanda |
|---|---|---|---|---|---|---|
| Internet | 12 | 350000 | 350000 | 350000 | Tidak | Tidak |
| Iuran kesehatan | 13 | 150000 | 150000 | 150000 | Tidak | Tidak |
| Listrik (berubah-ubah) | 20 | 600000 | 900000 | 780000 | Tidak | Tidak |
| Air (berubah-ubah) | 29 | 250000 | 250000 | 310000 | **Ya** | **Ya**, dasar `routine_amount` |

Kolom: `BNK-S` rekening bank suami; `BNK-I` rekening bank istri; keduanya berpemilik `shared`. `UANG = BNK-S + BNK-I`. `KB` = kekayaan bersih rumah tangga = `UANG`. `PRI-S` = tempat uang pribadi suami (tidak terlihat istri). Semua entri dicatat suami.

| Baris | BNK-S | BNK-I | UANG | KB | PRI-S |
|---|---|---|---|---|---|
| Sebelum | 400000 | 250000 | 650000 | 650000 | 150000 |
| Δ 11 Nov: penghasilan 8000000 (`income`, sumber `salary`) | 8000000 | 0 | 8000000 | 8000000 | 0 |
| Sesudah | 8400000 | 250000 | 8650000 | 8650000 | 150000 |
| Δ 11 Nov: bayar internet 350000, jatuh tempo 12 (`expense`) | -350000 | 0 | -350000 | -350000 | 0 |
| Sesudah | 8050000 | 250000 | 8300000 | 8300000 | 150000 |
| Δ 11 Nov: bayar iuran kesehatan 150000, jatuh tempo 13 (`expense`) | -150000 | 0 | -150000 | -150000 | 0 |
| Sesudah | 7900000 | 250000 | 8150000 | 8150000 | 150000 |
| Δ 11 Nov: siapkan uang untuk listrik 900000 dan air 250000 (`allocation`) | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 7900000 | 250000 | 8150000 | 8150000 | 150000 |
| Δ 11 Nov: serah terima 5500000 ke rekening istri (`transfer`) | -5500000 | 5500000 | 0 | 0 | 0 |
| Sesudah | 2400000 | 5750000 | 8150000 | 8150000 | 150000 |
| Δ 11 Nov: 500000 ke uang pribadi suami (`owner_transfer`) | -500000 | 0 | -500000 | -500000 | 500000 |
| Sesudah | 1900000 | 5750000 | 7650000 | 7650000 | 650000 |
| Δ 20 Nov: bayar listrik 780000, membawa tujuan "listrik" (`expense`) | -780000 | 0 | -780000 | -780000 | 0 |
| Sesudah | 1120000 | 5750000 | 6870000 | 6870000 | 650000 |
| Δ 20 Nov: lepas sisa label "listrik" 120000 (`allocation`) | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 1120000 | 5750000 | 6870000 | 6870000 | 650000 |
| Δ 29 Nov: bayar air 310000, membawa tujuan "air" (`expense`) | -310000 | 0 | -310000 | -310000 | 0 |
| Sesudah | 810000 | 5750000 | 6560000 | 6560000 | 650000 |

Uang yang disiapkan di `BNK-S` (`jumlah alokasi = listrik + air`; `belum dialokasikan = saldo tercatat − jumlah alokasi`):

| Saat | Saldo tercatat | Alokasi listrik | Alokasi air | Jumlah alokasi | Belum dialokasikan |
|---|---|---|---|---|---|
| Sesudah dua tagihan dibayar, sebelum alokasi | 7900000 | 0 | 0 | 0 | 7900000 |
| Sesudah alokasi | 7900000 | 900000 | 250000 | 1150000 | 6750000 |
| Sesudah serah terima | 2400000 | 900000 | 250000 | 1150000 | 1250000 |
| Sesudah ke uang pribadi | 1900000 | 900000 | 250000 | 1150000 | 750000 |
| Sesudah bayar listrik 780000 | 1120000 | 120000 | 250000 | 370000 | 750000 |
| Sesudah sisa label listrik dilepas | 1120000 | 0 | 250000 | 250000 | 870000 |
| Sesudah bayar air 310000 | 810000 | 0 | 0 | 0 | 810000 |

Yang ditunjukkan contoh ini:

- **Bayar lebih awal.** Internet dan iuran kesehatan dibayar tanggal 11, sebelum jatuh tempo 12 dan 13; tanggal kejadiannya 11. Keduanya sama dengan jumlah yang disepakati (350000 dan 150000): tidak ditandai.
- **Menyiapkan uang, lalu membagikan sisanya.** Sesudah alokasi, layar menunjukkan 6750000 yang masih bebas. Suami membagikan 5500000 + 500000 = 6000000 dan menyisakan 750000 bebas di rekeningnya; 6750000 − 6000000 = 750000. Alokasi tidak mengubah saldo dan tidak menghalangi apa pun: bila ia membagikan lebih dari 6750000, rekening itu tampil "alokasi melebihi uang" (§3.9).
- **Tagihan yang berubah-ubah, di atas ambang.** Listrik 780000 di atas ambang 500000 dan di atas perkiraannya 600000, tetapi tidak melebihi jumlah per kali yang disepakati, 900000: tidak ditandai; ambang tidak dibandingkan. Yang disiapkan adalah batas atasnya; sisanya 900000 − 780000 = 120000 dilepas dan kembali bebas: 750000 + 120000 = 870000. Bila tagihannya 950000, ia **ditandai** dengan dasar `routine_amount`, karena 950000 > 900000; kelebihannya, 950000 − 900000 = 50000, tidak dibandingkan dengan ambang (rumusan lead yang dikukuhkan pemilik produk, 11 November 2026; sebelum revisi ini kalimat ini berbunyi "tidak ditandai"). Bila kesepakatannya belum berlaku atau sudah diakhiri, 780000 dinilai dengan aturan biasa dan membawa tanda dasar `threshold` (780000 > 500000); angkanya sama.
- **Tagihan yang berubah-ubah, di bawah ambang.** Air 310000 melebihi yang disiapkan, 250000. Alokasi "air" berkurang sampai nol, dan 310000 − 250000 = 60000 diambil dari yang bebas: 870000 − 60000 = 810000. Tidak ada yang ditolak. Entri ini **ditandai** dengan dasar `routine_amount`: 310000 > 250000, walau 310000 tidak di atas ambang 500000 (sebelum revisi ini: tidak ditandai). Bila air disepakati sebesar jumlah tertingginya yang wajar, misalnya 350000, 310000 tidak ditandai; bila kesepakatannya belum berlaku, aturan biasa berlaku dan 310000 juga tidak ditandai. Inilah harga menyepakati perkiraan dan bukan batas atasnya (§3.15; §10.3 butir 9).
- **Jumlah tanda.** Dari tujuh entri uang contoh ini, **tepat satu** bertanda: pembayaran air. Sebelum revisi ini tidak ada.
- **Bagian untuk diri sendiri.** 0 + 500000 = 500000 tidak melebihi jatah 500000: tidak ditandai. Pengambilan berikutnya dalam periode ini, berapa pun, akan ditandai (§3.4).
- **Serah terima.** Baris +5500000 pada `BNK-I` bertanda belum dikonfirmasi sampai istri mengonfirmasi; saldonya sudah berubah (§3.3).

Ringkasan periode: pemasukan 8000000; belanja 350000 + 150000 + 780000 + 310000 = 1590000; disisihkan ke dana pribadi suami 500000; perubahan `KB` = 8000000 − 1590000 − 500000 = 5910000, dan 650000 + 5910000 = 6560000. Custody akhir: suami memegang 810000; istri memegang 5750000, di antaranya 5500000 belum dikonfirmasi; 810000 + 5750000 = 6560000. `PRI-S`: 150000 + 500000 = 650000. Di layar, enam langkah tanggal 11 dapat disajikan sebagai satu alur "uang masuk, lalu bagikan" (dokumen 17 D-55); di model tiap langkah tetap entri tersendiri.

Menguji: INV-02, INV-03, INV-06, INV-12, INV-17, INV-19, INV-30.

### 7.15 Contoh putaran S3 (semuanya berdiri sendiri)

Delapan contoh untuk aturan yang lahir di putaran penyebaran S3 (12 November 2026). Tidak satu pun mengubah keadaan berantai atau masuk CTH-37; semua angkanya **[Asumsi]**. Tiap contoh memakai kolomnya sendiri; aturan membaca baris `Δ` dan `Sesudah` sama dengan §7.1. Baris yang ditolak atau yang hanya peristiwa tetap ditulis dengan Δ nol supaya skrip dapat memeriksanya.

#### CTH-42 — Saldo awal yang dicatat bukan pemegangnya, lalu dikoreksi

Aturan: §3.3 (PA-FIN-2), §3.17 (PA-FIN-3), INV-31. Kolom: `TUN-S` uang tunai bersama yang dipegang suami, baru didaftarkan; `UANG = TUN-S`; `KB = UANG`.

| Baris | TUN-S | UANG | KB |
|---|---|---|---|
| Sebelum (baru didaftarkan, belum bersaldo awal) | 0 | 0 | 0 |
| Δ 24 Okt: istri mencatat saldo awal 200000 (`opening`) | 200000 | 200000 | 200000 |
| Sesudah | 200000 | 200000 | 200000 |
| Δ 25 Okt: suami mencatat penyesuaian saldo awal −20000 (`opening`, `corrects` saldo awal asli, bertanggal kejadian 24 Okt) | -20000 | -20000 | -20000 |
| Sesudah | 180000 | 180000 | 180000 |
| Δ perintah saldo awal asli kedua 180000: ditolak | 0 | 0 | 0 |
| Sesudah | 180000 | 180000 | 180000 |
| Δ perintah penyesuaian −200000: ditolak (180000 − 200000 = −20000 < 0) | 0 | 0 | 0 |
| Sesudah | 180000 | 180000 | 180000 |

Baris +200000 dicatat istri pada tempat uang yang dipegang suami, jadi bertanda belum dikonfirmasi di sisi suami sampai ia mengonfirmasi atau berkeberatan (§2.5 baris (a)); penyesuaian −20000 dicatat suami sendiri dan tidak bertanda. Penyesuaian bukan konfirmasi: keduanya tindakan terpisah (dokumen 06 C02, C03). Bila suami sendiri yang mencatat saldo awal, tidak ada penanda dan angkanya sama. Saldo awal bersih 200000 − 20000 = 180000. Ringkasan periode yang memuat 24 Oktober tidak memuat saldo awal itu di tabel A maupun arus kas; ia disebut di bawahnya sebagai "saldo awal yang didaftarkan periode ini 180000" (§3.17, INV-19). Bila angka awal yang salah dikoreksi dengan menghitung uang, hasilnya adalah selisih hitung −20000 atas nama suami; karena itu koreksinya penyesuaian. Menguji: INV-02, INV-19, INV-31.

#### CTH-43 — Entri terlambat atas tempat uang yang sudah terarsip

Aturan: §2.6 "Mengaktifkan kembali" (PA-FIN-4), §3.10 butir 6, INV-21. Kolom: `EWL-X` dompet digital dana bersama yang dipegang istri; `TUN-I` uang tunai bersama di istri; `UANG = EWL-X + TUN-I`; `KB = UANG`.

**Bagian A: belanja yang sudah tertangkap selisih hitung.** Tidak perlu diaktifkan kembali.

| Baris | EWL-X | TUN-I | UANG | KB |
|---|---|---|---|---|
| Sebelum | 30000 | 500000 | 530000 | 530000 |
| Δ 5 Nov: istri menghitung `EWL-X` 0; selisih 0 − 30000 = −30000 (`cash_difference`) | -30000 | 0 | -30000 | -30000 |
| Sesudah | 0 | 500000 | 500000 | 500000 |
| Δ 6 Nov: `EWL-X` diarsipkan (peristiwa, tanpa entri) | 0 | 0 | 0 | 0 |
| Sesudah | 0 | 500000 | 500000 | 500000 |
| Δ 10 Nov: parkir 30000 bertanggal 3 Nov dari `EWL-X`, merinci selisih (`recategorization`, kategori "transportasi") | 0 | 0 | 0 | 0 |
| Sesudah | 0 | 500000 | 500000 | 500000 |

Rekonsiliasi `EWL-X`: selisih −30000, dirinci 30000, belum dirinci 0. Perinciannya tidak punya baris posisi, jadi diterima walau `EWL-X` terarsip. "Transportasi" terpakai bertambah 30000 pada periode yang memuat 3 November.

**Bagian B: belanja yang tidak tertangkap selisih** (berdiri sendiri dari bagian A). Pada 5 November istri memindahkan seluruh isi `EWL-X` yang tercatat, 50000, ke `TUN-I`, lalu mengarsipkan `EWL-X`. Sebenarnya 15000 sudah terpakai dari `EWL-X` pada 4 November, dan yang benar-benar pindah hanya 35000.

| Baris | EWL-X | TUN-I | UANG | KB |
|---|---|---|---|---|
| Sebelum (sesudah pemindahan 50000 dan pengarsipan) | 0 | 550000 | 550000 | 550000 |
| Δ 12 Nov: perintah belanja 15000 atas `EWL-X` terarsip: ditolak (INV-21) | 0 | 0 | 0 | 0 |
| Sesudah | 0 | 550000 | 550000 | 550000 |
| Δ 12 Nov: `EWL-X` diaktifkan kembali (peristiwa, tanpa entri) | 0 | 0 | 0 | 0 |
| Sesudah | 0 | 550000 | 550000 | 550000 |
| Δ belanja 15000 bertanggal 4 Nov dari `EWL-X` (`expense`) | -15000 | 0 | -15000 | -15000 |
| Sesudah | -15000 | 550000 | 535000 | 535000 |
| Δ penyesuaian atas pemindahan 5 Nov sebesar −15000 (`transfer`, `corrects` pemindahan itu) | 15000 | -15000 | 0 | 0 |
| Sesudah | 0 | 535000 | 535000 | 535000 |
| Δ `EWL-X` diarsipkan lagi (peristiwa) | 0 | 0 | 0 | 0 |
| Sesudah | 0 | 535000 | 535000 | 535000 |

Saldo −15000 sesudah belanja terlambat menunjukkan bahwa catatan lain keliru: pemindahan tercatat 50000, padahal 35000. Penyesuaiannya mengembalikan `EWL-X` ke 0 dan menurunkan `TUN-I` ke 535000, yaitu 500000 + 35000. Penyesuaian atas pemindahan tidak mengubah `UANG` (INV-06). Tidak ada entri lama yang disunting. Menguji: INV-06, INV-09, INV-16, INV-21.

#### CTH-44 — Satu belanja dibayar dari dana bersama dan dana pribadi

Aturan: §3.2 (PA-FIN-6), §3.4. Istri membeli beras 300000: 200000 dari uang tunai bersama, 100000 dari uang pribadinya, tanpa minta diganti. Kolom: `TUN-I` uang tunai bersama di istri; `UANG = TUN-I`; `KB = UANG`; `PRI-I` uang pribadi istri (tidak terlihat suami).

| Baris | TUN-I | UANG | KB | PRI-I |
|---|---|---|---|---|
| Sebelum | 500000 | 500000 | 500000 | 250000 |
| Δ entri 1: bagian dana bersama 200000 (`expense`, "belanja bulanan") | -200000 | -200000 | -200000 | 0 |
| Sesudah | 300000 | 300000 | 300000 | 250000 |
| Δ entri 2: bagian dana pribadi 100000 (`contribution_expense`, "belanja bulanan") | 0 | 0 | 0 | -100000 |
| Sesudah | 300000 | 300000 | 300000 | 150000 |

Baris entri 2: buku pribadi istri, `PRI-I` −100000 dan `owner_transfer` −100000; buku rumah tangga, `owner_transfer` +100000 dan `expense` −100000 (`funded_by` istri). Keseimbangan buku rumah tangga: 0 = 100000 − 100000. Ringkasan: "belanja bulanan" terpakai bertambah 200000 + 100000 = 300000; tabel A: belanja −300000, kontribusi istri +100000, jumlah −200000 = perubahan `KB` (300000 − 500000); tabel C: istri, dari dana bersama 200000, kontribusi 100000, jumlah 300000 = tabel B; arus kas: belanja −200000 = perubahan `UANG`. Ambang 500000: entri 1 berjumlah pemakaian 200000, tidak ditandai; entri 2 tidak dinilai (§3.8). Bila berasnya 700000, dibayar 400000 + 300000, bagian bersama 400000 tetap tidak ditandai: yang dinilai hanya dana bersama. Menguji: INV-02, INV-03, INV-15, INV-19, INV-20.

#### CTH-45 — Membatalkan tujuan tabungan yang masih beralokasi

Aturan: §3.9 "Menyelesaikan atau membatalkan tujuan" (G-01), §2.6. Tujuan "dana darurat" beralokasi di dua tempat uang dana bersama yang dipegang istri. Istri membatalkannya dan memilih melepas semua alokasinya: satu entri `allocation`, dua baris. Tidak ada saldo yang berubah.

| Tempat uang | Saldo tercatat | Alokasi "belanja bulanan" | Alokasi "dana darurat" sebelum | Δ alokasi "dana darurat" | Alokasi "dana darurat" sesudah | Belum dialokasikan sebelum | Belum dialokasikan sesudah |
|---|---|---|---|---|---|---|---|
| TUN-I | 3800000 | 2500000 | 300000 | -300000 | 0 | 1000000 | 1300000 |
| TAB-I | 5500000 | 0 | 5500000 | -5500000 | 0 | 0 | 5500000 |

`Belum dialokasikan = saldo tercatat − alokasi "belanja bulanan" − alokasi "dana darurat"`: `TUN-I` 3800000 − 2500000 − 300000 = 1000000, sesudahnya 3800000 − 2500000 − 0 = 1300000; `TAB-I` 5500000 − 0 − 5500000 = 0, sesudahnya 5500000. Kemajuan "dana darurat": 300000 + 5500000 = 5800000 menjadi 0, lalu tujuan itu terarsip dengan keterangan `cancelled`. `UANG` dan `KB` tidak berubah. Tabel D periode itu (**[Asumsi]**: alokasi 5800000 sudah ada pada awal periode): awal 5800000, perubahan −5800000, akhir 0, keterangan "diarsipkan"; periode berikutnya tujuan itu tidak tampil. Bila istri tidak memilih melepas, tujuan tidak terarsip dan semua angka tetap. Menguji: INV-12, INV-17, INV-21.

#### CTH-46 — Hitungan offline yang tiba sesudah serah terima

Aturan: §3.10 "Hitungan yang tiba belakangan" (G-02), INV-16. Kolom: `TUN-S` uang tunai bersama di suami; `TUN-I` uang tunai bersama di istri; `UANG = TUN-S + TUN-I`.

| Baris | TUN-S | TUN-I | UANG |
|---|---|---|---|
| Sebelum (server dan ponsel istri sama) | 900000 | 500000 | 1400000 |
| Δ 10 Nov: suami menyerahkan 200000 tunai; dicatat di ponselnya dan sampai di server (`transfer`) | -200000 | 200000 | 0 |
| Sesudah (server) | 700000 | 700000 | 1400000 |
| Δ 10 Nov: hitungan istri, dibuat tanpa sinyal, terhitung 700000, sampai di server sesudahnya; selisih server 700000 − 700000 = 0, tanpa entri | 0 | 0 | 0 |
| Sesudah | 700000 | 700000 | 1400000 |

Ponsel istri, yang belum menerima serah terima itu, menampilkan tercatat 500000 dan perkiraan selisih 700000 − 500000 = +200000. Bila selisih diambil dari perangkat, entri `cash_difference` +200000 membuat `TUN-I` 700000 + 200000 = 900000 dan `UANG` 1600000, padahal uang di tangan istri 700000: 200000 yang sama tercatat dua kali. Dengan selisih server, rekonsiliasinya "dihitung, cocok", `recorded_amount_at_count` 700000. Serah terimanya tetap bertanda belum dikonfirmasi sampai istri mengonfirmasi; menghitung bukan konfirmasi. Menguji: INV-16, INV-23.

#### CTH-47 — Hitungan oleh anggota yang bukan pemegang

Aturan: §3.10 butir 1 (G-05; syarat (a) sampai (c) dari SP-50, putaran perbaikan S3), §2.5 baris (c), §3.3, INV-33. Kolom: `BNK-I` rekening dana bersama yang dipegang istri, saldonya dapat dilihat suami, dan istri sudah menandainya "pasanganku dapat melihat saldo rekening ini" (`others_may_count = true`); `UANG = BNK-I`; `KB = UANG`.

| Baris | BNK-I | UANG | KB |
|---|---|---|---|
| Sebelum | 3800000 | 3800000 | 3800000 |
| Δ 10 Nov: suami melihat saldo rekening 3750000 dan merekonsiliasinya: selisih 3750000 − 3800000 = −50000 (`cash_difference`) | -50000 | -50000 | -50000 |
| Sesudah | 3750000 | 3750000 | 3750000 |
| Δ 11 Nov: suami menghitung lagi, terhitung 3740000: ditolak, hitungan 10 Nov masih terbuka (INV-33) | 0 | 0 | 0 |
| Sesudah | 3750000 | 3750000 | 3750000 |
| Δ 12 Nov: istri mengonfirmasi baris −50000 (peristiwa, tanpa entri) | 0 | 0 | 0 |
| Sesudah | 3750000 | 3750000 | 3750000 |
| Δ 12 Nov: suami menghitung 3740000: selisih 3740000 − 3750000 = −10000 (`cash_difference`) | -10000 | -10000 | -10000 |
| Sesudah | 3740000 | 3740000 | 3740000 |

Baris −50000 dicatat suami pada tempat uang yang dipegang istri, jadi bertanda belum dikonfirmasi di sisi istri (§2.5 baris (c)). Custody sesudah baris 10 November: istri memegang `BNK-I` 3750000, dengan 50000 keluar belum dikonfirmasi; sesudah baris terakhir: 3740000, dengan 10000 keluar belum dikonfirmasi. Angkanya berubah seketika, seperti setiap baris bertanda (§3.3). Istri dapat mengonfirmasinya, merinci selisih itu (misalnya biaya administrasi bank 50000, sebagai `recategorization` tanpa mengubah saldo), atau berkeberatan; keberatan tidak mengubah angka sampai ada entri pembalik. Bila suami menghitung 3800000, hasilnya "dihitung, cocok", tanpa entri dan tanpa penanda. Bila istri sendiri yang mencatat selisih yang sama, angkanya sama dan tidak ada penanda. Hitungan 11 November ditolak walau angkanya mungkin benar: sampai istri menjawab, satu catatan "uangmu kurang" atas namanya sudah cukup; ia juga ditolak bila hasilnya cocok. Jawaban istri 12 November juga dapat berupa keberatan atau hitungannya sendiri atas `BNK-I`; ketiganya menutup hitungan 10 November. Baris −10000 bertanda belum dikonfirmasi seperti baris −50000, dan sejak itu hitungan 12 November yang terbuka. Jumlah selisih yang dicatat suami: −50000 + (−10000) = −60000 = 3740000 − 3800000. **Bila istri belum menandai `BNK-I`**, atau tempat uang itu uang tunai, hitungan 10 November pun ditolak dan semua angka tetap 3800000. Bila istri mematikan tandanya sesudah 10 November, hitungan itu tetap terbuka sampai ia jawab, dan hitungan suami berikutnya ditolak karena tandanya mati. Menguji: INV-02, INV-16, INV-33.

#### CTH-48 — Menghitung lalu merinci seluruh selisih sekaligus

Aturan: §3.10 "Menghitung lalu merinci seluruh selisih sekaligus" (dokumen 11, IDEA-01). Angka mengikuti skenario dokumen 11; nama kategori mengikuti contoh dokumen ini. Kolom: `TUN-I` uang tunai bersama di istri; `UANG = TUN-I`; `KB = UANG`.

| Baris | TUN-I | UANG | KB |
|---|---|---|---|
| Sebelum | 450000 | 450000 | 450000 |
| Δ 14 Nov: istri pulang dari pasar, menghitung 285000, dan merinci seluruh selisih 450000 − 285000 = 165000: "belanja bulanan" 145000, "transportasi" 20000 (`expense`, membawa `reconciliation_id`) | -165000 | -165000 | -165000 |
| Sesudah | 285000 | 285000 | 285000 |

145000 + 20000 = 165000, sama dengan selisih server, jadi bentuk ini dipakai. Rekonsiliasinya: terhitung 285000, tercatat saat hitung 450000, tanpa entri `cash_difference`, belum dirinci 0. Perbandingan dengan jalur biasa atas angka yang sama:

| Laporan | Merinci sekaligus (`expense`) | Jalur biasa (`cash_difference` −165000, lalu `recategorization`) |
|---|---|---|
| Tabel A: belanja bulanan | -145000 | -145000 |
| Tabel A: transportasi | -20000 | -20000 |
| Tabel A: selisih hitung yang belum dirinci | 0 | 0 |
| Arus kas: belanja | -165000 | 0 |
| Arus kas: selisih hitung negatif | 0 | -165000 |
| Anggaran "belanja bulanan": tambahan terpakai | 145000 | 145000 |
| `TUN-I` akhir | 285000 | 285000 |

Kedua jalur sama di tabel A, anggaran, saldo, dan perubahan `UANG` (−165000); bedanya hanya arus kas, yang pada jalur biasa menyebut uang yang sudah diketahui keperluannya sebagai "selisih hitung". Ambang 500000: baris belanja 165000 tidak ditandai. Menguji: INV-02, INV-15, INV-16, INV-19.

#### CTH-49 — Cicilan yang jumlah bayarnya di bawah harga; cicilan 0%

Aturan: §3.14 "Biaya kredit tidak pernah negatif" (G-07). Kolom: `TUN-I` uang tunai bersama di istri; `UANG = TUN-I`; `AST` aset; `KWJ-C` kewajiban cicilan; `KB = UANG + AST − KWJ-C`.

**Bagian A: dibayar lebih kecil dari harga.** Barang berlabel 4000000, uang muka 0, 12 × 300000 = 3600000.

| Baris | TUN-I | UANG | AST | KWJ-C | KB |
|---|---|---|---|---|---|
| Sebelum | 500000 | 500000 | 0 | 0 | 500000 |
| Δ perintah harga 4000000: ditolak (0 + 3600000 < 4000000) | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 500000 | 500000 | 0 | 0 | 500000 |
| Δ dicatat ulang dengan harga 3600000; biaya kredit 0 + 3600000 − 3600000 = 0, tanpa baris `credit_cost` (`asset_purchase`) | 0 | 0 | 3600000 | 3600000 | 0 |
| Sesudah | 500000 | 500000 | 3600000 | 3600000 | 500000 |

**Bagian B: cicilan 0% dengan uang muka** (dari `Sebelum` yang sama). Harga 4000000, uang muka 400000, 12 × 300000 = 3600000; 400000 + 3600000 = 4000000.

| Baris | TUN-I | UANG | AST | KWJ-C | KB |
|---|---|---|---|---|---|
| Sebelum | 500000 | 500000 | 0 | 0 | 500000 |
| Δ beli secara cicilan 0% (`asset_purchase`, tanpa baris `credit_cost`) | -400000 | -400000 | 4000000 | 3600000 | 0 |
| Sesudah | 100000 | 100000 | 4000000 | 3600000 | 500000 |

Keseimbangan bagian A: 0 + 3600000 − 3600000 = 0; bagian B: −400000 + 4000000 − 3600000 = 0; keduanya tanpa sebab. `KB` akhir bagian A: 500000 + 3600000 − 3600000 = 500000; bagian B: 100000 + 4000000 − 3600000 = 500000. Jumlah pemakaian yang dibandingkan dengan ambang (§3.8): bagian A 0 + 3600000 = 3600000, bagian B 400000 + 3600000 = 4000000; keduanya di atas ambang contoh 500000, jadi diajukan lebih dulu atau bertanda. Bila rumah tangga pada bagian A menilai barang itu 4000000, mereka mencatat `asset_valuation` +400000 yang tampil sebagai perubahan nilai aset, bukan biaya kredit. Menguji: INV-02, INV-07, INV-13.

### 7.16 Contoh sesudah checkpoint S3 (berdiri sendiri)

Dua contoh: CTH-50 untuk IDEA-03, yang diadopsi sesudah checkpoint S3, dan CTH-51 untuk THR yang diberikan (**[Usulan]** `finance-domain`). Keduanya tidak mengubah keadaan berantai atau CTH-37; semua angka dan tanggalnya **[Asumsi]**. Aturan membaca baris `Δ` dan `Sesudah` sama dengan §7.1.

#### CTH-50 — Pos rutin tanpa tanggal: token listrik

Aturan: §3.15 "Tagihan berulang tanpa `due_rule`: pos rutin tanpa tanggal", aturan 3, 4, 7, 7a; §3.8 dasar `routine_amount`; INV-26. Pos "Token listrik" tanpa `due_rule`, kategori "listrik", kesepakatan rutin berlaku dengan `R` = 300000 per kali (angka skenario dokumen 11 IDEA-03). Ambang 500000; anggaran "listrik" periode November 600000; rumah tangga dua anggota dewasa. Ketiga pembelian dicatat suami, ditautkan ke pos itu saat dicatat, tanpa kejadian. Kolom: `TUN-S` uang tunai bersama di suami; `UANG = TUN-S`; `KB = UANG`.

| Baris | TUN-S | UANG | KB |
|---|---|---|---|
| Sebelum | 1000000 | 1000000 | 1000000 |
| Δ 3 Nov: token 200000 (`expense` "listrik", tertaut ke pos; `P` 200000 ≤ `R` 300000: tidak ditandai) | -200000 | -200000 | -200000 |
| Sesudah | 800000 | 800000 | 800000 |
| Δ 15 Nov: token 300000 (`P` 300000 ≤ `R` 300000: tidak ditandai) | -300000 | -300000 | -300000 |
| Sesudah | 500000 | 500000 | 500000 |
| Δ 20 Nov: perintah tautan susulan yang membawa kejadian 20 November untuk entri 15 Nov: ditolak (pos tanpa `due_rule` tidak punya kejadian; INV-26) | 0 | 0 | 0 |
| Sesudah | 500000 | 500000 | 500000 |
| Δ 27 Nov: token 350000 (`P` 350000 > `R` 300000: ditandai dasar `routine_amount`; tidak dibandingkan dengan ambang 500000) | -350000 | -350000 | -350000 |
| Sesudah | 150000 | 150000 | 150000 |

Jumlah perubahan: −200000 − 300000 + 0 − 350000 = −850000 = 150000 − 1000000. Tanda: satu, pada entri 27 November, walau 350000 tidak di atas ambang; jumlah November 850000 di atas ambang dan di atas `R` tidak menandai apa pun, karena kekerapan tidak diperiksa (batas yang dinyatakan). Yang menangkap jumlah bulanannya adalah anggaran "listrik": terpakai 200000, 500000, 850000; sisa 600000 − 200000 = 400000, 400000 − 300000 = 100000, 100000 − 350000 = −250000. Ringkasan November: tabel A dan B baris "listrik" −850000 dan 850000; tabel C 850000 dibayar suami dari dana bersama; tidak ada baris per pos. Pos ini tidak pernah tampil di "belum dibayar periode ini" dan tidak memicu pengingat. Bila kesepakatannya belum berlaku, ketiga entri dibandingkan dengan ambang sendiri-sendiri (aturan 5): 200000, 300000, dan 350000 tidak di atas 500000, jadi tidak ada tanda. Menguji: INV-02, INV-12, INV-15, INV-26.

#### CTH-51 — THR dan angpau yang diberikan dalam satu musim Lebaran

Aturan: §3.2 "THR dan angpau yang diberikan"; §3.3 (tarik tunai); §3.8 "Entri yang tertaut ke otorisasi". Ambang 500000; rumah tangga dua anggota dewasa. Suami dan istri menyetujui satu permintaan musiman "THR dan Lebaran" berbatas 2000000, tanpa tempat uang sumber, jadi tanpa reservasi. Semua entri di tabel dicatat suami, dari dana bersama, dan ditautkan ke otorisasi itu saat dicatat; kategori "THR dan pemberian hari raya". Kolom: `BNK-S` rekening dana bersama di suami; `TUN-S` uang tunai bersama di suami; `UANG = BNK-S + TUN-S`; `KB = UANG`.

| Baris | BNK-S | TUN-S | UANG | KB |
|---|---|---|---|---|
| Sebelum | 5000000 | 0 | 5000000 | 5000000 |
| Δ tarik tunai 1500000 untuk amplop (`transfer`, tanpa biaya admin) | -1500000 | 1500000 | 0 | 0 |
| Sesudah | 3500000 | 1500000 | 5000000 | 5000000 |
| Δ biaya jasa tukar uang pecahan baru 50000 (`expense`; realisasi 50000) | 0 | -50000 | -50000 | -50000 |
| Sesudah | 3500000 | 1450000 | 4950000 | 4950000 |
| Δ THR asisten rumah tangga 600000 (`expense`, penerima manfaat pihak luar, hubungan pekerja rumah tangga; realisasi 650000 ≤ 2000000: tidak ditandai) | 0 | -600000 | -600000 | -600000 |
| Sesudah | 3500000 | 850000 | 4350000 | 4350000 |
| Δ angpau satu entri, dua baris sebab: keponakan 300000, "anak-anak tetangga" 200000 (`expense`; realisasi 1150000: tidak ditandai) | 0 | -500000 | -500000 | -500000 |
| Sesudah | 3500000 | 350000 | 3850000 | 3850000 |
| Δ THR untuk ibu suami 1000000 lewat transfer bank (`expense`, kiriman keluarga; realisasi 2150000 > 2000000; `N` = 1000000 − (2000000 − 1150000) = 150000, tidak di atas 500000: tidak ditandai; tampil "melebihi Rp150.000") | -1000000 | 0 | -1000000 | -1000000 |
| Sesudah | 2500000 | 350000 | 2850000 | 2850000 |

Jumlah perubahan: 0 − 50000 − 600000 − 500000 − 1000000 = −2150000 = 2850000 − 5000000; jumlah pengeluaran 50000 + 600000 + 300000 + 200000 + 1000000 = 2150000 = realisasi. Ringkasan periode: tabel A dan B baris "THR dan pemberian hari raya" 2150000 (bila seluruhnya jatuh dalam satu periode anggaran); tabel C 2150000 dibayar suami dari dana bersama. Tarik tunai tidak tampil di tabel A, B, maupun arus kas.

**Tanpa permintaan musiman** (angka saldo sama persis; hanya tanda yang berbeda): 50000 tidak di atas 500000; 600000 di atas 500000, ditandai; 500000 tidak di atas 500000 (sama dengan ambang), tidak ditandai; 1000000 ditandai. Dua tanda. Bila angpau 500000 dicatat sebagai dua entri 300000 dan 200000, keduanya juga tidak ditandai: penilaian per entri (batas yang dinyatakan di §3.2).

**Dari dana pribadi** (di luar tabel): istri memberi 100000 kepada anak sahabatnya dari dana pribadinya sebagai urusan pribadinya: `expense` di buku pribadinya, tidak terlihat suami, dan tidak mengubah satu angka pun di tabel di atas. Bila ia memberi THR asisten rumah tangga itu dari dana pribadinya untuk keperluan rumah tangga tanpa minta diganti, entrinya `contribution_expense` 600000: `UANG` tidak berubah, pengeluarannya terlihat keduanya, dan tidak dibandingkan dengan ambang. Menguji: INV-02, INV-06, INV-14, INV-26.

---

## 8. Riwayat yang tidak dapat diubah lawan hak penghapusan

**Status: keputusan sementara pemilik produk, checkpoint S2 (S2-8 pilihan A, dinyatakannya sendiri): bacaan append-only di bagian ini diterima sampai advokat menjawab (L-04, L-20, L-22, L-23).** Rinciannya (dokumen 17 D-41 dan D-42) adalah default yang diterima di checkpoint S2 lewat "lanjut", bukan pernyataannya sendiri. Sisi hukumnya tetap **[Perlu validasi]**; ADR-007 (`architect`) tetap berstatus sementara. Label **[Usulan]** yang masih tertulis di bagian ini dibaca menurut status ini. Bukan nasihat hukum.

### 8.1 Ketegangannya

- Brief §7 dan `ENGINEERING-STANDARDS.md` §5: riwayat keuangan hanya ditambah.
- UU PDP Pasal 43–44 mewajibkan penghapusan dan pemusnahan atas permintaan subjek data; pengecualiannya (Pasal 50) hanya untuk kepentingan negara, dan tidak ada pengecualian untuk keutuhan catatan yang dimiliki bersama orang lain (**[Terverifikasi]** primer, dokumen 18 §2.1, diakses 11 November 2026). Salinan teks PP 33/2026 mengukur pemusnahan dari hasilnya, "tidak lagi dapat digunakan untuk mengidentifikasi Subjek Data Pribadi" (Pasal 84 ayat (2)), menyebut salinan cadangan (Penjelasan Pasal 82), dan tidak menyebut pseudonimisasi sebagai cara memenuhi kewajiban hapus (dokumen 18 §2.1).
- Apakah pseudonimisasi identitas pada entri memenuhi kewajiban itu: **[Perlu validasi]** hukum (L-04).

### 8.2 Aturan yang diusulkan

Entri dipisah menjadi dua bagian sejak awal (§2.5, §2.6):

| Bagian | Isi | Sifat |
|---|---|---|
| **Inti moneter** | `entry` dan `line`: jenis, tanggal, jumlah, posisi, sebab, kategori, tautan; pelaku dan pihak sebagai **pengenal internal**, bukan nama | Hanya ditambah; tidak pernah dihapus |
| **Rincian** | `entry_detail`: catatan teks, struk, nama pihak lawan dalam teks bebas | Dapat dihapus tanpa menyentuh inti moneter |
| **Pemetaan identitas** | Tabel yang mengaitkan pengenal internal anggota dan pihak luar dengan orangnya (nama, akun) | Dapat diputus: setelah diputus, pengenal itu tidak lagi menunjuk siapa pun |

Atas permintaan penghapusan dari seorang anggota, atau ketika ia keluar dari rumah tangga dan memintanya (yang berlaku bila ia keluar **tanpa** memintanya ada di butir 2a):

1. **Ekspor ditawarkan lebih dulu** (data miliknya sendiri).
1a. **Ekspor buku rumah tangga sesudah keluar** (usulan `security-privacy` SP-48, putaran perbaikan S3; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-4 (c)). Sesudah aksesnya berakhir, yang keluar tetap dapat mengekspor buku rumah tangga **sampai saat aksesnya berakhir**: entri dan tautan yang sudah diterima server sebelum saat itu, dengan saldo per saat itu. Entri yang diterima sesudahnya tidak termasuk, juga bila tanggal kejadiannya lebih awal, karena itu catatan anggota yang tersisa sesudah ia pergi. Rincian (`entry_detail`) dalam ekspor itu mengikuti butir 4 pada saat ekspor dibuat. Buku pribadinya sendiri dapat ia ekspor sampai dihapus (butir 2a). Berapa lama dan lewat jalan apa ekspor itu tersedia adalah aturan dokumen 10 (jangka retensi); batasnya terhadap data anggota yang tersisa dinyatakan di §8.3.
2. **Klaim antar-pemilik yang masih terbuka ditutup oleh sistem, lalu buku pribadinya dihapus seluruhnya.** Buku pribadi bukan catatan bersama; tidak ada orang lain yang bergantung pada keutuhannya. Penghapusan **tidak bersyarat** pada penggantian, pelunasan, atau pelepasan oleh siapa pun (temuan `security-privacy` SP-12): keluar dengan aman tidak boleh menunggu tindakan pasangan, dan anggota yang tersisa tidak boleh tertinggal dengan klaim yang hanya dapat ditutup orang yang sudah pergi. Saat anggota keluar, sistem menulis satu entri `exit_closure` (di bawah) yang menolkan setiap klaim antara buku pribadinya dan buku lain, sehingga buku rumah tangga dan buku pasangan tetap seimbang (INV-08). Kedua pihak tetap boleh menyelesaikan klaim itu lebih dulu dengan entri biasa bila mau; itu pilihan, bukan syarat. Sisi klaim yang berada di buku rumah tangga atau buku pasangan tetap ada sebagai riwayat bersaldo nol, dengan pihak lawan dipseudonimkan. Penghapusan baris buku pribadi ini **satu-satunya** penghapusan baris yang diusulkan, dan ADR harus menyatakannya sebagai pengecualian eksplisit terhadap INV-09. Entri dua buku (kontribusi, talangan, penyisihan) kehilangan baris sisi pribadinya; baris sisi rumah tangganya tetap dan tetap seimbang per pemilik (INV-02), sedangkan INV-03 dan INV-08 sejak itu hanya diperiksa atas pemilik yang masih ada.
2a. **Keluar tanpa meminta penghapusan** (usulan putaran S3; PA-FAM-7; T-37; mengadopsi bawaan dokumen 05 `FAM.md`; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), D-86). Dua hal dipisah: *penutupan* buku pribadi dan *penghapusan* barisnya. **Penutupan terjadi pada setiap keluarnya anggota**, diminta atau tidak, pada akhir masa tenggang (butir 2b): sistem menulis `exit_closure` (INV-28) dan buku pribadi itu sejak saat itu tidak menerima baris baru dari siapa pun (selama masa tenggang ia sudah dibekukan). Alasannya: sesudah keluar ia bukan lagi `member(M)` rumah tangga itu, jadi tidak ada pelaku yang berhak menulis buku itu (INV-25), dan klaim antar-pemilik yang dibiarkan terbuka tidak akan pernah dapat ditutup. **Penghapusan** baris buku pribadi terjadi seketika bila diminta (butir 2), dan bila tidak diminta pada akhir jangka retensi yang ditetapkan dokumen 10; sampai itu isinya tidak terbaca siapa pun di rumah tangga itu, dan hanya dapat diekspor pemiliknya. Cara mantan anggota mengakses ekspor itu, dan jangka retensinya, adalah aturan dokumen 10. Butir 3 sampai 8 berlaku sama, kecuali butir 4 (rincian yang ia tulis): tanpa permintaan, rincian di buku pribadi ikut jangka retensi yang sama, dan rincian di buku rumah tangga tetap. Tidak ada angka rumah tangga yang berbeda antara keluar dengan atau tanpa permintaan; yang berbeda hanya kapan baris buku pribadi dihapus.
2b. **Masa tenggang keluar** (usulan `security-privacy` SP-48, putaran perbaikan S3; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-4 (c) dan D-86; panjang 7 hari **[Asumsi]**, A-39; alurnya, autentikasi ulang, dan siapa yang diberi tahu milik dokumen 06 dan 10). Saat anggota keluar, **aksesnya berakhir seketika**; `exit_closure`, penutupan buku pribadinya, dan penghapusan yang ia minta dijalankan **pada akhir masa tenggang**. Hanya yang keluar yang dapat membatalkannya, di dalam masa itu. Akibat uangnya:

    | Hal | Selama masa tenggang | Pada akhir masa tenggang | Bila dibatalkan |
    |---|---|---|---|
    | Keanggotaan dalam model uang | M masih `member(M)`: tetap terhitung anggota dewasa (termasuk syarat "≥ dua anggota dewasa" tanda, §3.8), tetap pemegang tempat uangnya. Yang berakhir hanya aksesnya | Berakhir bersama `exit_closure` | Tidak pernah berakhir |
    | Klaim antar-pemilik dengan buku M (hak ganti, pinjaman dana bersama kepada M) | **Tetap terbuka** dan tetap tampil di buku rumah tangga serta di `KB` | Ditutup `exit_closure` sebesar sisanya | Tetap terbuka |
    | Buku pribadi M | **Dibekukan**: tidak menerima baris dari siapa pun. M tidak punya akses; `claim_waiver` atas klaim dana bersama kepada M oleh anggota lain ditolak. Karena setiap klaim antar-pemilik hanya berubah lewat entri yang menyentuh buku M, saldo klaim pada akhir masa tenggang **sama** dengan saat akses berakhir. **Angka `exit_closure` tidak berubah, hanya waktunya** (CTH-38; skenario SCN-44: −200000 + 60000 = −140000) | Ditutup; lalu dihapus bila diminta (butir 2), atau pada akhir jangka retensi (butir 2a) | Menerima baris lagi seperti biasa |
    | Buku rumah tangga | Berjalan biasa; anggota yang tersisa mencatat apa saja. Entri atas tempat uang dana bersama yang dipegang M boleh (misalnya serah terima isinya menurut butir 5); baris yang menambah tempat uang M bertanda belum dikonfirmasi dan tidak dapat dijawab M selama masa itu. Hitungan bukan pemegang atas tempat uang M ditolak (§3.10 butir 1 syarat (c)) | Entri masa tenggang tetap, tidak dinilai ulang | Entri masa tenggang tetap, tidak dinilai ulang; penanda yang belum dijawab kini dapat dijawab M |
    | Perintah M yang belum terkirim saat akses berakhir | Ditolak, karena tidak ada akses. Syarat bagi dokumen 06 dan 09: keluar dijalankan sesudah perangkatnya mengirim perintah yang tertunda, atau layar keluar menyebut berapa catatan yang belum terkirim dan bahwa catatan itu tidak akan masuk | — | Dapat dikirim ulang dengan idempotency key yang sama (INV-10) |
    | Tanggal kejadian `exit_closure` | — | Hari akhir masa tenggang; baris "klaim ditutup karena anggota keluar" tampil di ringkasan periode yang memuat hari itu | — |
    | Ekspor buku rumah tangga | Batasnya saat akses berakhir (butir 1a), bukan akhir masa tenggang | Sama | — |

    Alasan bentuk ini: keluar karena terdesak (misalnya di bawah tekanan) dapat disesali, dan `exit_closure` serta penghapusan tidak dapat dibalik (tabel di bawah); masa tenggang memberi jalan kembali tanpa membuat angka rumah tangga bergantung pada pilihan itu. Membekukan buku M, bukan membiarkannya ditulis anggota lain, menjaga angka penutupan tetap, dan `claim_waiver` yang ditolak tidak merugikan siapa pun: `exit_closure` sebesar yang sama mengubah `KB(shared)` dengan jumlah yang sama, hanya dengan nama baris yang lain.
3. **Di buku rumah tangga, inti moneter tetap.** Pelaku, pembayar, penerima manfaat, `earner`, dan pemegang yang menunjuk anggota itu tetap berupa pengenal internal; pemetaan identitasnya diputus, sehingga layar menampilkan "mantan anggota".
3a. **Anggota yang bergabung sesudah ada yang keluar** (usulan `security-privacy` SP-49, putaran perbaikan S3; diterima lewat "lanjut" di checkpoint S3 (12 November 2026), D-90). Ia membaca inti moneter buku rumah tangga seperti anggota lain, termasuk entri sebelum ia bergabung, karena angka rumah tangga identik bagi semua pelihat (INV-20); pelaku dan pihak yang menunjuk orang yang sudah keluar tampil "mantan anggota". **Rincian (`entry_detail`: catatan teks, foto struk, nama pihak lawan dalam teks bebas) pada entri yang dicatat mantan anggota itu tidak terbaca olehnya**, juga bila rincian itu tetap ada bagi anggota yang sudah ada saat orang itu keluar (butir 4). Tidak ada angka, saldo, atau laporan yang berubah karenanya; yang dibatasi hanya rincian. Pembatasan ini cocok dengan butir 4 dan §8.3: rincian yang dipertahankan sebagai "data bersama" adalah data bersama orang-orang yang hidup bersama saat itu, bukan data yang dibuka bagi orang yang datang kemudian.
4. **Rincian yang ia tulis dihapus**: catatan teks dan struk pada entri yang ia catat, kecuali pada entri dana bersama yang rinciannya juga data anggota yang tersisa. Untuk kasus itu usulannya: struk dan catatan tetap, nama orang yang pergi di dalam teks bebas tidak dapat dijamin terhapus, dan hal ini dinyatakan terbuka sebagai batas (lihat 8.3).
5. **Tempat uang dana bersama yang ia pegang** tidak dihapus dan tidak dinolkan otomatis: uangnya nyata. Anggota yang tersisa mencatat apa yang sebenarnya terjadi (serah terima, atau pembagian sebagai `owner_transfer`), lalu mengarsipkannya. Pembagian harta ketika rumah tangga berpisah adalah keputusan keluarga atau hukum, bukan aturan aplikasi.
6. **Cadangan**: penghapusan rincian dan pemutusan pemetaan harus menjangkau salinan cadangan dalam jangka yang ditetapkan dokumen 10.
7. **Pihak luar** (ibu, adik, ayah) hanya tersimpan sebagai nama panggilan dan hubungan (Q-28); penghapusannya adalah pemutusan pemetaan yang sama.
8. **Entri dua buku lama.** Sesudah anggota itu keluar, tidak ada lagi yang dapat membalik entri dua buku yang pernah ia catat (§3.12). Anggota yang tersisa tetap dapat menyesuaikan sisi rumah tangganya. Bagaimana objek yang satu-satunya pelakunya sudah pergi ditampilkan dan dikeluarkan dari daftar terbuka adalah aturan dokumen 06 (§8.2 "Keberatan", §7.2 "Aturan umum").

**Entri `exit_closure`: keputusan sementara pemilik produk, checkpoint S2 (S2-8 butir (d)); sisi hukumnya [Perlu validasi] (L-04, L-22).**

| Hal | Aturan |
|---|---|
| Jenis dan pelaku | `type = exit_closure`, `actor = system`. Bukan tindakan anggota mana pun: tidak memerlukan persetujuan, konfirmasi, atau pernyataan dari yang keluar maupun yang tersisa, dan tidak dapat dicatat lewat layar |
| Kapan | Satu entri per anggota yang keluar, dengan atau tanpa permintaan penghapusan (butir 2a), ditulis dalam transaksi basis data yang sama dengan penutupan buku pribadinya, **pada akhir masa tenggang keluar** (butir 2b; tidak ditulis bila keluar dibatalkan di dalam masa itu); tanggal kejadian = hari akhir masa tenggang; idempotency key diturunkan dari peristiwa keluarnya, sehingga pengulangan tidak menulis apa pun (INV-10) |
| Baris | Untuk setiap klaim antar-pemilik yang menyangkut buku pribadi anggota itu dan bersaldo X > 0: piutang −X dan `owner_transfer` −X di buku pemilik piutang; kewajiban −X dan `owner_transfer` +X di buku yang berkewajiban. Tidak ada baris tempat uang, aset, atau alokasi |
| Yang dinyatakannya | Hanya bahwa aplikasi berhenti melacak klaim itu karena salah satu pihaknya keluar. Ia **tidak** menyatakan klaim itu dibayar, dan **tidak** menyatakan siapa pun mengikhlaskannya. Layar: "ditutup karena anggota keluar" |
| Akibat pada uang | Tidak ada: `U(o)` semua pemilik tetap; *uang yang ada* tetap |
| Akibat pada kekayaan bersih rumah tangga | Naik sebesar kewajiban rumah tangga kepada anggota itu yang ditutup (hak ganti, utang dana bersama kepadanya); turun sebesar piutang rumah tangga kepadanya yang ditutup. Klaim antara dua buku pribadi tidak menyentuh angka rumah tangga |
| Laporan | Baris tersendiri di tabel A, "klaim ditutup karena anggota keluar", per klaim (§6.3). Bukan kontribusi, bukan penyisihan, bukan "utang dihapuskan"; tidak terhitung terhadap anggaran; tidak muncul di arus kas; *jumlah pemakaian* nol, jadi tidak dibandingkan dengan ambang dan tidak ditandai (§3.8) |
| Invarian yang berlaku | INV-01, INV-02 (seimbang di tiap buku), INV-03 (jumlah `owner_transfer` nol), INV-05 dan INV-06 dalam arti `ΔU = 0`, INV-07, INV-08 (kedua sisi tiap klaim nol), INV-09 (entri hanya ditambah), INV-10, INV-13 (X tepat sebesar sisa), INV-19, INV-20, INV-28. INV-25 mendapat pengecualian kedua |
| Tidak dapat dibalik | Sesudah buku pribadinya dihapus tidak ada sisi lawan untuk dipulihkan. Bila orang itu bergabung lagi, ia anggota baru dengan buku baru; klaim lama, bila masih diakui keduanya, dicatat ulang sebagai saldo awal |
| Sisi hukum **[Perlu validasi]** | Apakah penutupan catatan di aplikasi dapat dibaca sebagai pelepasan hak atas utang yang sebenarnya, dan bagaimana kedudukannya terhadap harta bersama, belum ditinjau advokat (L-04). Maksud model: tidak; utang-piutang yang sebenarnya antara kedua orang itu tidak diputuskan aplikasi. Ekspor yang ditawarkan lebih dulu (butir 1) memuat daftar klaim terbuka beserta jumlahnya bagi yang keluar, dan sisi rumah tangganya tetap terbaca oleh yang tersisa |

Contoh terhitung: CTH-38.

### 8.3 Batas yang dinyatakan

| Batas | Akibat |
|---|---|
| Di rumah tangga dua orang, "mantan anggota" mudah dikenali oleh anggota yang tersisa | Pseudonimisasi tidak menyembunyikan apa pun dari pasangan; ia hanya memutus kaitan bagi operator dan pihak lain. Apakah itu memenuhi Pasal 84 ayat (2): L-04 |
| Jumlah dan tanggal belanja dana bersama tetap ada | Bila advokat menyatakan inti moneter pun harus dimusnahkan, satu-satunya jalan yang menjaga angka adalah **entri penutup**: buku rumah tangga dirangkum menjadi saldo awal baru per posisi, lalu entri lama dihapus. ADR harus menyebut jalan ini sebagai rencana cadangan dan biayanya: jejak audit sebelum tanggal rangkuman hilang bagi anggota yang tersisa |
| Teks bebas dan foto struk dapat memuat data orang lain | Tidak dapat dipilah otomatis. Usulan: teks bebas pada entri dana bersama diperlakukan sebagai data bersama |
| Hak anggota yang tersisa atas catatan bersama tidak diatur UU PDP sebagai pengecualian | Inilah inti L-04 |
| Ekspor buku rumah tangga oleh yang keluar (butir 1a; SP-48) | Yang keluar membawa salinan catatan bersama sampai saat aksesnya berakhir, termasuk catatan dan struk yang ditulis anggota yang tersisa. Ini sisi lain dari butir di atas: catatan bersama adalah data kedua pihak, dan masing-masing menyimpan haknya atas catatan sampai ia pergi. Jangka ketersediaannya dan apakah rincian tulisan anggota yang tersisa ikut diekspor: dokumen 10 dan L-04. Angkanya tidak dapat berbeda dari buku rumah tangga per saat itu; koreksi yang dicatat sesudahnya tidak ikut |
| Anggota yang bergabung kemudian (butir 3a; SP-49) | Rincian tulisan mantan anggota tidak terbaca baginya, sedangkan inti moneternya terbaca. Ini tidak bertentangan dengan baris "teks bebas dan foto struk" di atas: perlakuan "data bersama" berlaku bagi orang yang berbagi rumah tangga dengan penulisnya. Yang tidak dapat dijaga: angka dan tanggal entri lama tetap terbaca, dan nama mantan anggota dapat muncul di rincian yang ditulis anggota lain |

---

## 9. Keputusan atas pertanyaan terbuka

Keputusan dokumen ini atas butir dokumen 17 dan butir glosarium yang semula bertanda *(diputuskan di S2)*. Kedudukan tiap baris sesudah checkpoint S2 (11 November 2026): baris bertanda "keputusan pemilik" diputuskannya sendiri; baris bertanda "default yang diterima" ia terima lewat "lanjut" di checkpoint S1 (keduanya dibedakan sesuai K-24); **semua baris lain adalah default yang diterima di checkpoint S2 lewat "lanjut"** (dokumen 17 D-01 sampai D-55), bukan pernyataannya sendiri, kecuali yang disebut lain di barisnya. Jawaban checkpoint S2 yang mengubah atau mengukuhkan aturan dokumen ini ada di baris-baris terakhir tabel.

| ID | Keputusan | Bagian |
|---|---|---|
| Q-01 | Keputusan pemilik: gaji dana bersama secara bawaan, dapat diubah saat mencatat. *Pinjam kembali* adalah serah terima: default yang diterima pemilik produk lewat "lanjut" di checkpoint S1, bukan keputusannya sendiri | §3.1, §3.3, CTH-27 |
| Q-56, K-16 | Keputusan pemilik produk, 11 November 2026: pencatat uang masuk memilih arahnya, termasuk gaji ke dana pribadi; pasangan melihat jumlah gaji dan bagian pribadinya; saldo dan pemakaian dana pribadi tetap tertutup. Bentuknya: satu entri `income` dengan baris di dua buku; penghasilan (`salary`, `business`) tidak dapat dicatat hanya di buku pribadi (INV-24). Di luar keputusan itu: hadiah untuk satu orang tidak terlihat (Q-77) kini keputusan pemilik produk, checkpoint S2; bagian pribadi di dalam entri `income` diperlakukan sama dengan penyisihan lain (T-13, Q-78) adalah default yang diterima di checkpoint S2 lewat "lanjut" | §3.1, §3.4, §5.3, CTH-03b |
| Q-55 (sisi model) | "Punya siapa" pada layar adalah `earner`, terpisah dari pemilik uang; pinjaman yang diterima adalah `loan_received`, bukan pemasukan, dan bawaannya utang rumah tangga | §2.4, §3.1, CTH-28 |
| Q-02, Q-23, K-13 | Dana pribadi **di luar** garis kendali rumah tangga. *Uang yang ada* = dana bersama saja = angka utama, identik bagi keduanya; tidak ada nama kedua. Baris kedua "uang pribadi saya" hanya milik pelihat | §2.2, §5 |
| Q-05 | Keputusan pemilik untuk asasnya. Rinciannya: selisih negatif terhitung sebagai uang keluar yang belum dirinci; selisih positif baris tersendiri di sisi masuk, bukan pemasukan; terhitung terhadap anggaran hanya setelah dirinci; entri terlambat dapat menutup selisih tanpa menyentuh saldo; tanpa penutupan periode | §3.10, §3.11, CTH-29 |
| Q-06, Q-29 | Keputusan pemilik: anggaran adalah batas per kategori, bulanan, tanpa amplop. Hari mulai periode ditentukan rumah tangga (keputusan pemilik produk 15, 11 November 2026, menyempurnakan "mulai tanggal gajian" dari checkpoint S1); "bawaan 1, berlaku mulai periode berikutnya" adalah bacaan lead yang dikukuhkan kepadanya, bukan kata-katanya. **[Usulan]** dokumen ini: rentang 1 sampai 28 dan satu periode peralihan pendek dengan batas penuh. Periode disimpan sebagai rentang tanggal; tanpa rollover | §6.1 |
| Q-07 | Default yang diterima pemilik produk lewat "lanjut" di checkpoint S1: kiriman ibu adalah pengeluaran saat diserahkan; saudara dan ayah dipilih per entri | §3.2, CTH-06, CTH-25, CTH-26 |
| Q-11 | Opsi B: ledger sederhana, entri bertipe dengan baris seimbang | §1 |
| Q-12 | Persetujuan tidak membuat reservasi, kecuali permintaan menyebut tempat uang sumber | §3.8, CTH-22 |
| Q-13 | Sisa: tetap di tempatnya (A, D); dikembalikan atau tetap dipegang pemohon sebagai dana bersama (B); hak ganti sebesar realisasi (C). Kelebihan: tercatat penuh; dari dana bersama tanpa klaim; dari dana pribadi kontribusi kecuali "minta diganti" | §3.8, CTH-21 sampai CTH-24 |
| Q-14 | Ya: terikat mengurangi sisa anggaran kategori otorisasinya dan tampil terpisah | §6.1, CTH-21 |
| Q-15, K-02 | Pengembalian dana mengurangi pengeluaran kategori asal pada periode tanggal pengembalian; pemasukan hanya bila aslinya tidak tercatat | §3.13, CTH-19 |
| Q-16, K-11 | Kewajiban cicilan = jumlah pembayaran tersisa; pembayaran adalah pelunasan yang terhitung terhadap anggaran; bunga diakui sebagai `credit_cost` saat kredit dibuka. *Realisasi* didefinisikan ulang sebagai jumlah baris tertaut, apa pun jenisnya | §3.14, §3.8, CTH-04, CTH-34 |
| Q-17 | Bukan kewajiban; pengingat sampai dibayar | §3.15 |
| Q-18 | Pembelian aset bukan pengeluaran; tampil di arus kas; nilai diperbarui manual dengan entri bertanggal | §3.16, CTH-33 |
| Q-19, K-20 (sisi uang) | Keputusan pemilik produk, 11 November 2026: tercatat seketika, bertanda belum dikonfirmasi sampai penerima menerimanya, keberatan kapan saja, sama ke dua arah. Rinciannya: saldo berubah seketika; penanda berlaku atas setiap baris yang menambah tempat uang berpemegang anggota dewasa selain pencatat; keberatan tidak mengubah angka sampai ada entri pembalik | §3.3, §5 |
| Q-20 | Pemegang adalah atribut tempat uang; tempat uang bersama dilaporkan sebagai baris "dipegang bersama" | §2.3, CTH-02 |
| Q-21 | Hak ganti tercatat sejak talangan sebagai kewajiban rumah tangga kepada anggota, berpasangan dengan piutang di buku pribadinya | §3.5, CTH-14 |
| N-14 | Pelepasan hak ganti adalah `owner_transfer`: kenaikan kekayaan bersih rumah tangga tampil sebagai kontribusi dari dana pribadi | §3.6, CTH-16 |
| Q-22, K-10 | §8: inti moneter buku rumah tangga tetap dan dipseudonimkan; rincian dan buku pribadi dapat dihapus. Keputusan sementara pemilik produk, checkpoint S2 (S2-8 pilihan A), sampai advokat menjawab (L-04); ADR-007 berstatus sementara | §8 |
| Q-24 | Diterima sebagai sifat klaim dua sisi; daftar kebocoran B1–B5 | §5.3 |
| Q-26 (sisi uang) | Default yang diterima pemilik produk: keduanya boleh mengoreksi entri dana bersama. Bentuknya entri pembalik atau penyesuaian; entri buku pribadi hanya oleh pemiliknya; entri dua buku hanya dapat dibalik pemilik buku pribadinya, sedangkan pasangannya dapat menyesuaikan sisi rumah tangga dan berkeberatan | §3.12, §2.2, CTH-30 |
| Dokumen 06 §15.4 butir 1 | Kait `authorization.amended`: tanpa entri; batas, tanggal, masa berlaku, *terikat*, dan reservasi mengikuti versi syarat baru | §3.8, varian CTH-22 |
| Dokumen 06 §15.4 butir 2 | Entri terlambat boleh ditautkan ke otorisasi `settled` atau `expired` bila tanggal kejadiannya di dalam rentang hidup otorisasi itu; tidak ke `cancelled`. Tautan boleh ditambahkan sesudah entri tersimpan sebagai catatan `entry_link` yang hanya ditambah | §3.8, §2.5, INV-26 |
| Dokumen 06 §15.4 butir 3 | Penanda `unconfirmed` berlaku atas setiap baris yang menambah tempat uang berpemegang anggota dewasa selain pencatat | §3.3, §2.5 |
| Dokumen 06 §15.4 butir 4 | Dikukuhkan dan dipertegas: setiap baris buku pribadi M hanya ditulis M, dengan satu pengecualian bagi anggota lain (pelepasan klaim dana bersama atas M) dan satu bagi sistem (`exit_closure` saat M keluar, **[Usulan]**). Entri antara dua buku pribadi ditunda dari MVP | §2.2, INV-25 |
| Dokumen 06 §15.4 butir 5 | `set_aside_schedule`, dan `routine_agreement_ref` pada tagihan berulang dan jadwal penyisihan (tidak pada cicilan, lihat butir (a)) | §2.6, §3.4 |
| Dokumen 06 §15.4 butir 6 | `links.activity_id` diganti `links.subject_ref` (rujukan induk) | §2.5 |
| Dokumen 06 §15.4 butir 7 | Pembelian aset secara cicilan: uang muka ditambah jumlah seluruh cicilan | §3.8, §3.16, CTH-34 |
| Dokumen 06 §15.4 butir (a) | Diterapkan: `routine_agreement_ref` dihapus dari `installment`; pembayaran cicilan dan pelunasan utang yang sudah tercatat (`repayment_out`) tidak dibandingkan dengan ambang dan tidak memerlukan kesepakatan rutin. Tidak ada keberatan dari sisi uang; batas yang semula dinyatakan (kewajiban yang didaftarkan belakangan sebagai saldo awal) kini dijawab butir (m) | §2.6, §3.8, §3.14 |
| Dokumen 06 §15.4 butir (b) | Diterapkan: tabel *jumlah pemakaian* memuat `claim_waiver` untuk kedua arahnya | §3.8, §3.6 |
| Dokumen 06 §15.4 butir (c) | Dikukuhkan: `member(M)` adalah anggota dalam satu rumah tangga; buku pribadi per anggota, bukan per akun; tanpa entri lintas rumah tangga | §2.2 |
| Dokumen 06 §15.4 butir (d) | Tidak bertentangan: jumlah yang disepakati (per kali untuk tagihan berulang, per periode anggaran untuk jatah pribadi) dibaca dari kesepakatan; `estimated_amount` dan `monthly_amount` pada jadwal adalah metadata perkiraan dan rencana | §2.6 |
| Dokumen 06 §15.4 butir (e), (f), (i), (j), (k) | Sudah diterapkan pada revisi sebelumnya; lihat baris SP-01, SP-05, SP-12, SP-13 di bawah. Judul INV-24 sudah diganti (SP-22) | INV-27, §2.5, §3.3, §3.12, §8.2, §6.3, INV-24 |
| Dokumen 06 §15.4 butir (g) | Dikukuhkan: penutupan keberatan (C05) tidak memicu kait uang apa pun dan bukan konfirmasi; `handover.confirmed` hanya dari C02 dan C03 | §3.3, §10.2 |
| Dokumen 06 §15.4 butir (h) | Diterapkan: tautan susulan ke otorisasi hanya oleh pencatat entri atau pemohon otorisasi; pencabutan hanya oleh pembuat tautan atau pencatat entri; aturannya milik dokumen 06 (INV-A34) dan dikutip | §3.8 |
| Dokumen 06 §15.4 butir (l) | Diterapkan: kalimat kedua INV-24 dibatasi pada entri `income` yang memuat baris bersumber penghasilan; hadiah pribadi di buku pribadi saja tetap sah (CTH-36) | INV-24 |
| Dokumen 06 §15.4 butir (m) | Diterapkan dengan dua penegasan sisi uang: (1) baris `expense` berpemilik `shared` dihitung di entri jenis apa pun, **kecuali** baris `expense` pada `contribution_expense` (bukan dana bersama) dan pada `advance` (sudah dihitung sebagai hak ganti); (2) saldo awal kewajiban dana bersama kepada pihak luar punya jumlah pemakaian sebesar saldo awalnya, dibandingkan satu kali terhadap ambang yang berlaku saat disimpan; pembayarannya tidak ditautkan lagi bila `opening`-nya tertaut ke otorisasi | §3.8, §3.14, §3.17 |
| Pernyataan pemilik produk, 11 November 2026: kedua pasangan berpenghasilan | Tidak ada aturan baru. Aturan sudah per penerima penghasilan (`earner`) dan simetris; kalimat yang menyiratkan satu pencari nafkah diperbaiki. Bagian pribadi yang rutin dari penghasilan memakai `set_aside_schedule` per penerima dan dapat dicakup *kesepakatan rutin*; tidak ada kolom yang kurang. Bentuk jatah pribadi kini diputuskan pemilik produk (dua baris berikut) | §3.1, §3.4, §7.1, CTH-03b, §10.3 |
| Keputusan pemilik produk, 11 November 2026: jatah pribadi (Q-98) | Keputusan pemilik produk (14, 16, 17): jatah pribadi adalah satu jumlah per orang per periode anggaran, dijumlahkan sepanjang periode; besarnya dapat diatur dan **disepakati berdua**: satu mengusulkan, yang lain mengonfirmasi; kenaikan diajukan dengan alasan. Selebihnya rancangan dokumen ini dan dokumen 06. Bentuknya: `set_aside_schedule` per penerima yang dirujuk satu kesepakatan rutin. Keputusan pemilik produk, checkpoint S2 (S2-1 butir (iii) pilihan b): jatah adalah batas; begitu jumlah penyisihan periode itu melebihi jatah, entrinya ditandai berapa pun kelebihannya, dan ambang tidak lagi dibandingkan bagi anggota yang punya jatah. Default yang diterima di checkpoint S2 lewat "lanjut" (dokumen 17 D-48), bukan pernyataannya sendiri: penyisihan kepada M (entri `owner_transfer` dan bagian pribadi entri `income`) di dalam jatah tidak ditandai; tanpa jatah yang berlaku tiap penyisihan dibandingkan dengan ambang sendiri-sendiri; tanpa akibat pada saldo, anggaran, atau kekayaan bersih; tidak berpindah periode; tanpa angka lintas periode. Setiap penyisihan kepada M terhitung dengan sendirinya, tanpa tautan; periodenya menurut tanggal kejadian; entri pembalik melonggarkan jatah ke depan saja (sama dengan dokumen 06 §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi"). Kenaikan jatah diajukan dengan alasan dan dikonfirmasi pasangan (pernyataan pemilik produk); alurnya milik dokumen 06, tanpa akibat uang | §2.6, §3.4, §3.8, INV-30, CTH-39 |
| Keputusan pemilik produk, 11 November 2026: hari mulai periode anggaran | Keputusan pemilik produk (15): hari mulai ditentukan rumah tangga. Bacaan lead yang dikukuhkan kepadanya, bukan kata-katanya: satu hari mulai, bawaan 1, berlaku mulai periode berikutnya; periode lampau tidak berubah. **[Usulan]**: 1 sampai 28; satu periode peralihan pendek (1 sampai 30 hari) dengan batas anggaran dan jatah penuh. Putusan lead, diterima pemilik produk sebagai default di checkpoint S2 lewat "lanjut" (Q-102; dokumen 17 D-49; bukan pernyataannya sendiri): penggantian hari mulai adalah kesepakatan rumah tangga (kedua anggota dewasa, seperti ambang), dan terlihat oleh keduanya | §2.6, §6.1, INV-29 |
| Pernyataan pemilik produk, 11 November 2026: penghasilan datang beberapa kali ("sebulan" adalah bacaan lead), ke rekening bank atau dompet digital; diteruskan ke istri atau ditarik tunai lalu dibagikan | Tidak ada aturan baru: tiap uang masuk satu entri; serah terima dari rekening (§3.3; CTH-08, CTH-18), tarik tunai (CTH-07) dan biaya admin (CTH-18), kiriman keluarga (CTH-06, CTH-26) dan pinjaman (CTH-25) sudah tercakup; satu contoh menggabungkannya | §3.1, §3.2, §3.3, CTH-39 |
| Q-27 (masukan) | Laporan per orang hanya dari buku rumah tangga dan sama bagi keduanya | §6.3 |
| SP-01 (pemeriksaan awal `security-privacy`) | Keputusan pemilik produk, checkpoint S2 (S2-1 butir (i)): bagian pribadi di dalam entri `income` bersumber penghasilan hanya bila `actor = earner`; `earner` adalah isian "punya siapa", wajib untuk penghasilan. Pilihan yang tidak diambil: tanpa pagar, pencatat mana pun boleh mengarahkan ke dana pribadinya sendiri | §3.1, INV-27 |
| SP-18 | Daftar sumber penghasilan adalah parameter preset aturan; preset pilot `salary`, `business` | §2.4, §3.1, INV-24 |
| SP-05 (sisi uang) | **[Usulan]**: penanda `unconfirmed` juga atas baris entri dua buku yang mengurangi tempat uang dana bersama berpemegang anggota dewasa selain pencatat; klaim internal yang dikeberatani tampil "dikeberatani" di daftar utang dan piutang, tanpa perubahan angka | §2.5, §3.3, §3.12, §5 |
| SP-12 (sisi uang) | **[Usulan]**; sisi hukum **[Perlu validasi]**: penghapusan buku pribadi tidak bersyarat; klaim antar-pemilik yang terbuka ditutup entri sistem `exit_closure` yang tidak menyatakan pembayaran dan tidak memerlukan tindakan siapa pun | §8.2, INV-28, CTH-38 |
| SP-13 | **[Usulan]**: tidak ada angka kumulatif, lintas periode, atau bersih per anggota atas penyisihan dan kontribusi; hanya baris per periode. Pelepasan klaim dana bersama atas M tampil sebagai "utang dihapuskan", bukan penyisihan kepada M | §6.3, §3.6, §5.3 |
| Q-86 | **[Usulan]**: uang antara dua rumah tangga yang sama-sama memakai aplikasi adalah dua catatan yang berdiri sendiri, tanpa tautan; bentuk yang lebih kaya ditunda ke tahap publik | §2.2 |
| Q-43 | Rupiah utuh; sen bank dibulatkan ke bawah | §2.1 |
| Q-44 | Tidak: satu tempat uang, satu pemilik | §2.3 |
| Q-45 | Jumlah alokasi dibatasi saldo tercatat saat perintah alokasi; belanja tidak pernah ditolak dan tidak mengurangi alokasi lain otomatis; kekurangannya tampil sebagai peringatan | §3.9 |
| Q-46 | Ya, tanggal akhir opsional | §3.15 |
| Q-47 | Baris tersendiri "piutang diikhlaskan", pada bulan penghapusan, di luar anggaran dan arus kas | §3.7, CTH-25 |
| Q-48 | Di entri yang sama, sebagai baris pengeluaran terpisah | §3.3, CTH-18 |
| Q-49 (sisi uang) | Penyisihan adalah `owner_transfer`, bukan pemindahan dan bukan pengeluaran; selalu terlihat keduanya; yang rutin punya jadwal penyisihan | §3.4, CTH-35 |
| Glosarium: saldo awal, entri terlambat, arsip | Lihat bagiannya | §3.17, §3.11, §2.6 |
| Checkpoint S2, S2-1 butir (i), (ii), (iii) | Keputusan pemilik produk (dinyatakannya sendiri): (i) hanya penerima penghasilan yang mengarahkan sebagian penghasilannya ke dana pribadinya; (ii) hadiah untuk satu orang masuk dana pribadinya dan tidak terlihat pasangan; (iii) pilihan b: jatah pribadi adalah batas | §3.1, §3.4, INV-27, INV-30, CTH-36, CTH-39 |
| Checkpoint S2, S2-1 butir (iv), (v) | Diterima pemilik produk atas saran lead, bukan pernyataannya sendiri: (iv) penghasilan pekerjaan lepas adalah *penghasilan* (sumber `business`; T-27 ditutup); (v) alasan pada usulan kenaikan jatah sendiri tidak wajib, yang merupakan aturan dokumen 06 dan tidak punya akibat uang (INV-12). Kata-kata pemilik produk atas butir (iv) (`_progress.md`): "ini masuknya uang masuk aja ga sih? misal kalau public kan kita gatau ada yang gajinya harian, bulanan, bahkan sebulan gada pemasukan coba pikirkan". Dibaca lead: irama penghasilan tidak boleh diandaikan; model memang tidak mengandaikannya (§2.4) | §2.4, §3.1, §3.4 |
| Checkpoint S2, S2-4 | Pilihan A. Bagian pembayaran rutin dinyatakan pemilik produk sendiri (tagihan rutin disepakati dulu; siapa pun boleh membayar); bagian ambang **tidak ia sebut**: lead menerapkan pilihan A dan menyampaikannya kepadanya, jadi bukan keputusannya dan bukan "diterima atas saran lead". Fakta rumah tangga pilot: tagihan dibayar lebih dulu, lalu sisanya dibagikan; kiriman orang tua tidak rutin; tanggal tagihan tersebar; sebagian jumlahnya berubah-ubah. Tidak ada aturan uang baru: alokasi untuk tagihan yang menyusul; jumlah per kali kesepakatan rutin sebagai batas atas | §3.15, §3.9, CTH-41 |
| Tiga penegasan 11 November 2026, sesudah revisi dokumen ini (`_progress.md`) | Rumusan lead yang dikukuhkan pemilik produk dengan "bener", bukan kata-katanya sendiri: (1) T-29 pilihan 1: jatah yang diakhiri pemiliknya tetap dasar penilaian sampai akhir periode anggaran itu; (2) tambahan satu bulan di atas jatah tetap mungkin lewat permintaan biasa yang disetujui (D-53); "bukan keringanan" berarti tidak ada kelonggaran otomatis; (3) setiap tagihan berulang disepakati saat dibuat, jumlah yang disepakati adalah batasnya, dan pembayaran di atasnya ditandai tanpa membandingkan kelebihannya dengan ambang. Tidak ada angka saldo yang berubah; tanda pada CTH-41 berubah | §3.4, §3.8 "Tanda atas entri", §3.15, INV-30, CTH-05, CTH-06, CTH-39, CTH-41, T-29, §10.3 butir 7, 9, 10 |
| Putaran perbaikan 12 November 2026: tinjauan `critic` (B2, S3, S4, S7, M4, M5), temuan `security-privacy` SP-32 sampai SP-40, empat jawaban pemilik produk | Diterapkan. Diterima pemilik produk atas saran lead ("ya bener gas"): tanda dasar jatah dan dasar pos rutin tanpa pemberitahuan langsung; konfirmasi susulan tanpa pengingat otomatis, satu pengingat ulang, tidak dapat diminta lagi sesudah "belum sepakat" (aturan dokumen 06); selisih dompet campuran dipilih pemegang setiap kali, bawaan dicabut. **[Usulan]** `security-privacy` (SP-38) dan dokumen ini: otorisasi konfirmasi susulan ditutup di transaksi persetujuan dengan batas sebesar entri dan tepat satu tautan. **[Usulan]** `security-privacy`, disampaikan dan tidak ia bantah: jatah yang tidak di atas ambang hanya diusulkan pemiliknya. **[Usulan]** `security-privacy`: jatah yang diakhiri tetap dinilai sampai akhir periode hanya bila di atas ambang; B6. Aturan lead dari D-07 (S7): bagian entri tertaut yang di atas batas otorisasi dibandingkan dengan ambang. Tidak ada angka saldo yang berubah | §2.3, §3.4, §3.8, §3.10, §5.3, INV-26, INV-30, CTH-21 sampai CTH-24, CTH-33, CTH-34, CTH-40, T-03, T-29, §10.3 butir 5, 7, 10, 11 |
| Dokumen 06 §15.4 butir (t) | **Diterapkan.** Entri yang tanggal kejadiannya lebih awal dari tanggal persetujuan boleh ditautkan ke otorisasi yang `open`; syarat tautan diperiksa hanya saat tautan dibuat. Tidak ada aturan uang yang berlawanan (INV-14: realisasi tidak bergantung pada tanggal atau status). Tidak diminta kait tersendiri | §3.8 "Penautan" butir 1, INV-26 |
| Dokumen 06 §15.4 butir (u) 1 sampai 9 | **Diterapkan seluruhnya.** (1), (2) §3.15 dan §3.8: tiga dasar tanda, nama sama dengan `flag.basis`. (3) tautan susulan ke tagihan berulang hanya oleh pencatat entri. (4), (5) CTH-41, CTH-05, CTH-06, CTH-39 dinilai ulang dengan jumlah yang disepakati dinyatakan **[Asumsi]**. (6) `ada_jatah`, INV-30, T-29 ditutup; batasnya waktu simpan server. (7) **Dikukuhkan**: `set_aside_schedule` dan kesepakatan jatah boleh berjumlah 0; jatah 0 tidak sama dengan tanpa jatah (**[Usulan]**). (8) glosarium bagian 7. (9) **Dikukuhkan**: server tidak pernah menerima `C`; kedua rekonsiliasi dikirim sebagai dua perintah yang tidak terkait di sisi yang terbaca pasangan | §2.3, §3.4, §3.8, §3.10, §3.15, CTH-40, CTH-41 |
| Checkpoint S2, S2-6 (Q-61) | Fakta rumah tangga pilot: uang pribadi dan uang bersama dapat terpisah maupun bercampur, di bank maupun tunai. "Satu tempat uang, satu pemilik" dipertahankan; dompet campuran dicatat sebagai dua tempat uang dan dihitung satu kali. Pembagian selisih dipilih pemegang setiap kali, tanpa bawaan (diterima pemilik produk atas saran lead, 12 November 2026; SP-40 (b)); bawaan "seluruhnya ke bagian bersama" dicabut | §2.3, §3.10, §5.3 B6, CTH-40, T-03 |
| Checkpoint S2, S2-7 (Q-67) | Keputusan pemilik produk: otorisasi berlaku sampai 7 hari sesudah tanggal rencananya, atau 30 hari sesudah disetujui bila tanpa tanggal. Kedua angka tetap **[Asumsi]**, diuji di pilot. Angkanya milik dokumen 06 §4.5; dokumen ini hanya menetapkan akibat uang `authorization.expired` dan tautan entri terlambat, yang tidak bergantung pada angkanya | §3.8, INV-26 |
| Checkpoint S2, S2-8 (Q-81, Q-96) | Keputusan sementara pemilik produk sampai advokat menjawab: bacaan append-only §8 diterima, termasuk penghapusan buku pribadi dan `exit_closure`. Sisi hukum **[Perlu validasi]** | §8, INV-28, CTH-38 |
| Checkpoint S2, D-01 sampai D-55 | Diterima lewat "lanjut", tidak satu pun diubah. Semuanya default yang diterima, bukan pernyataan pemilik produk | Seluruh dokumen |
| Putaran penyebaran S3 (12 November 2026): celah dokumen 05 `FIN-1.md` §13 | Semuanya **[Usulan]**, putaran S3. PA-FIN-2: saldo awal atas tempat uang yang dipegang orang lain bertanda belum dikonfirmasi. PA-FIN-3: tepat satu saldo awal per tempat uang, paling banyak satu per posisi lain, jumlah ≥ 0, koreksi hanya lewat penyesuaian, pembalik ditolak. PA-FIN-4: perincian selisih atas tempat uang terarsip boleh; selebihnya lewat pengaktifan kembali. PA-FIN-5: kejadian tagihan per bulan dan pembayaran menunjuk kejadiannya; bawaan dokumen 05 diganti. PA-FIN-6: dua entri. PA-FIN-7: tidak ada kategori yang diwajibkan; baris tanpa kategori di tabel B. PA-FIN-8: satu keterangan di layar anggaran dan pada tempat uangnya. PA-FIN-10: tempat uang pribadi selalu dipegang pemiliknya | §2.3, §2.6, §3.2, §3.3, §3.10, §3.14, §3.15, §3.17, §6.3, INV-21, INV-31, INV-32, CTH-42 sampai CTH-44 |
| Putaran penyebaran S3: celah dokumen 05 `FIN-2.md` §12, `FAM.md` PA-FAM-7, dokumen 11 §10.1 | Semuanya **[Usulan]**, putaran S3. G-01: tujuan diarsipkan hanya bila alokasinya nol, pelepasan ditawarkan tanpa pilihan terisi, tabel D memuat periode pelepasannya. G-02 (bersama `architect`): selisih dihitung server terhadap saldo yang sudah ia terima. G-03: `write_off` dibandingkan dengan ambang. G-05 (bersama `architect`): bukan pemegang boleh menghitung tempat uang dana bersama, selisihnya bertanda di sisi pemegang. G-07: biaya kredit tidak pernah negatif. PA-FAM-7: buku pribadi ditutup pada setiap keluar, dihapus atas permintaan atau pada akhir retensi. THR dan bonus: sumber `salary`. Menghitung lalu merinci seluruhnya: satu entri `expense` bertaut rekonsiliasi. Tagihan berulang tanpa `due_rule`: tidak diizinkan saat ini; perubahan minimalnya dinyatakan. Tagihan tidak terarsip sendiri | §2.6, §3.1, §3.7, §3.8, §3.9, §3.10, §3.14, §3.15, §8.2, INV-16, INV-19, INV-26, INV-28, CTH-25, CTH-38, CTH-45 sampai CTH-49, §10.4 |
| Putaran perbaikan S3, fase D (12 November 2026): CA-11, SP-48 sampai SP-51, PA-FIN-11, 06 §6.3 "Pos rutin" (g) | Semuanya **[Usulan]**; yang perlu pemilik produk menunggu checkpoint S3-4 (c) dan S3-9. SP-51: penghapusan piutang dinilai kumulatif per piutang. SP-50: hitungan bukan pemegang hanya atas rekening atau dompet digital yang ditandai pemegangnya, paling banyak satu terbuka per tempat uang. SP-48: masa tenggang keluar 7 hari **[Asumsi]**; `exit_closure` pada akhirnya dengan angka yang sama; buku pribadi dibekukan; ekspor buku rumah tangga sampai saat akses berakhir. SP-49: anggota yang bergabung kemudian tidak membaca rincian tulisan mantan anggota. PA-FIN-11: kejadian bawaan berikutnya yang belum dibayar. Jadwal penyisihan yang dirujuk jatah yang berlaku hanya diarsipkan pemilik jatah (mengikuti 06). CA-11: tanda bilangan teks CTH-44 | §2.3, §2.5, §2.6, §3.7, §3.8, §3.10, §3.15, §8.2, §8.3, INV-28, INV-33, INV-34, CTH-25, CTH-38, CTH-44, CTH-47, T-32, T-33, T-37, T-39, T-40, §10.4 |
| Checkpoint S3 (12 November 2026; `_progress.md` "Checkpoint S3 dijawab"), S3-1 sampai S3-9, D-56 sampai D-97 | Jawaban pemilik produk: "lanjut", yang menurut pernyataan lead berarti semua bawaan diterima. Diterima lewat "lanjut", bukan pernyataannya sendiri: S3-3 A (THR dan bonus `salary`; dibuka ulang sesudahnya, dua baris berikut), S3-9 A (i) dan (ii), S3-4 A (c) masa tenggang 7 hari **[Asumsi]**, S3-2 A (IDEA-01 dan IDEA-03 Tier 1; IDEA-02 tanpa aturan uang), S3-1 A (baris tabungan di halaman depan tanpa aturan uang baru; tanpa `GOL` lanjutan; pengingat tagihan N-32 Tier 1 per kejadian), S3-5 A, S3-6 A, dan setiap D yang merujuk dokumen ini (tabel label di awal dokumen). **IDEA-03 diadopsi**: `due_rule?` opsional (CTH-50). Butir yang tidak tercakup "lanjut" tetap **[Usulan]** (tabel label) | Label di awal dokumen, §2.6, §3.1, §3.7, §3.8, §3.10, §3.15, §8.2, INV-28, INV-33, INV-34, §10.1 (T-28 sampai T-40), §10.3 butir 10, §10.4 |
| Sesudah checkpoint S3 (12 November 2026), pertanyaan sampingan S3-3 dan pertanyaan susulan lead (`_progress.md`) | Pernyataan pemilik produk sendiri: "thr bisa menerima bisa ngasih"; "bro pikirkan kalau semua orang pakai app ini". Bacaan lead **[Usulan]**: THR yang diberikan dirancang untuk semua penerima di tahap publik. Diterapkan sebagai **[Usulan]** `finance-domain`: THR yang diberikan adalah `expense` biasa, tanpa jenis entri baru, ambang per entri, musim lewat satu permintaan musiman atau alokasi | §3.2, CTH-51, T-42 |
| Sesudah checkpoint S3 (12 November 2026), "Pertanyaan pemilik produk atas S3-3" (`_progress.md`) | Kata-kata pemilik produk sendiri: "kenapa thr dicatat sebagai gaji? kan pemasukan kita yang atur darimana nya?"; "ada kategori default bisa edit tambah atau yang lainnya". Bahwa "kategori" berarti sumber uang masuk: bacaan lead **[Usulan]**. S3-3 A dibuka ulang sebagian. Bawaan lead **[Usulan]**, ditanyakan di checkpoint S4: dua kelompok sumber (`work` selalu terlihat; `gift_other` boleh di buku pribadi); THR, bonus, lembur sumber bawaan `work`. Rincian siapa boleh menambah, mengganti nama, mengarsipkan, dan memindah kelompok: **[Usulan]** dokumen ini, mengikuti D-89. Akibat uang THR yang diterima sama dengan S3-3 A; tidak ada angka yang berubah | §2.4, §2.6 `income_source`, §3.1, INV-24, T-14, T-18, T-34, T-43, §10.4 butir 1 |

---

## 10. Butir yang masih terbuka dan permintaan ke pihak lain

### 10.1 Butir terbuka untuk dokumen 17 (`product-analyst`)

Asumsi dan bawaan yang dipakai dokumen ini tanpa pernyataan keluarga. Nomor `Q` dan `A` diberikan pemilik dokumen 17.

| # | Butir | Bawaan yang diusulkan | Akibat bila lain |
|---|---|---|---|
| T-01 | Tanda "minta diganti" pada belanja dana pribadi yang tertaut ke otorisasi dana bersama | Hidup secara bawaan (§3.5) | Bila mati, situasi C menjadi kontribusi dan CTH-23 berubah |
| T-02 | Sisa pada situasi B: dikembalikan atau tetap dipegang pemohon | Pemohon memilih; tanpa pilihan, tetap dipegang (tanpa entri) | Bila wajib dikembalikan, penyelesaian menunggu serah terima |
| T-03 | Dompet fisik yang memuat dana bersama dan dana pribadi sekaligus | **Bukan lagi kasus tepi**: fakta rumah tangga pilot (checkpoint S2, S2-6: "bisa terpisah bisa tercampur bisa di bank bisa tunai"). Dua tempat uang; seluruh isi dompet dihitung satu kali; selisihnya dibagi pemegang **setiap kali, tanpa bawaan**: bersama, pribadi, atau dibagi (diterima pemilik produk atas saran lead, 12 November 2026; SP-40 (b)). Bawaan "seluruhnya ke yang `shared`" dari revisi sebelumnya **dicabut**. §2.3, §3.10, CTH-40 | Usahanya terasa setiap kali menghitung dompet campuran: pemegang harus mengingat pembagiannya dan harus memilih. Yang memilih "bersama" atas uang yang sebenarnya pribadi menampilkan kekurangannya sebagai selisih hitung uang bersama atas namanya. Yang selalu memilih "pribadi" membuat selisih dompet campuran tidak pernah terlihat pasangan, dan *uang yang ada* dapat tercatat terlalu besar; model tidak dapat memeriksa pilihan itu. Bila pemilik produk menghendaki satu dompet satu saldo, aturan "satu tempat uang, satu pemilik" dan INV-20 harus dibuka kembali |
| T-04 | Entri terlambat ditawarkan sebagai "bagian dari selisih hitung" | Bawaan ya | Bila tidak, saldo turun dua kali dan rekonsiliasi berikutnya menghasilkan selisih positif |
| T-05 | Barang mana yang dicatat sebagai aset | Pilihan pencatat; tanpa batas nilai | Batas nilai memerlukan satu pengaturan lagi |
| T-06 | Saldo tercatat negatif | Diperbolehkan, memicu ajakan rekonsiliasi | Bila ditolak, pencatat pelupa terhalang mencatat |
| T-07 | Sen pada saldo bank | Dibulatkan ke bawah | Bila dicatat persis, satuan IDR menjadi sen dan semua layar menampilkan dua desimal |
| T-08 | Arah uang masuk dipilih pencatat, termasuk gaji ke dana pribadi | Bukan lagi asumsi: keputusan pemilik produk, 11 November 2026. Modelnya di §3.1; risiko yang tersisa di §10.3 butir 1 | — |
| T-09 | Angka rumah tangga contoh | Semua **[Asumsi]** kecuali 10000000 dan 3000000 | Data pilot yang sebenarnya dimasukkan saat onboarding, bukan ke dokumen ini |
| T-10 | Kategori bawaan saat onboarding | Daftar contoh yang semuanya dapat dihapus. **Putaran S3:** model tidak mewajibkan kategori apa pun, juga untuk biaya admin dan cicilan; baris tanpa kategori tersimpan dan tampil di baris "tanpa kategori" tabel B (§3.3, §3.14, §6.3; PA-FIN-7); bagian ini diterima lewat "lanjut" di checkpoint S3 (D-85, Q-66) | Diputuskan dokumen 05 dan 08 |
| T-11 | Titipan pada pihak luar tanpa login | Tersedia di model; tidak dipakai pilot | Bila dipakai, seseorang harus mencatat pemakaiannya |
| T-12 | Entri antara dua dana pribadi (talangan, pinjaman, penggantian antara uang pribadi suami dan uang pribadi istri) | Alur pencatatannya ditunda dari MVP; model dan contohnya (CTH-17b) tetap | Bila pilot membutuhkannya: satu tempat uang pribadi bawaan per anggota sebagai tujuan baris masuk yang dicatat orang lain, bertanda belum dikonfirmasi; INV-25 mendapat pengecualian kedua |
| T-13 | Bagian pemasukan yang diarahkan langsung ke dana pribadi, terhadap ambang permintaan | **Ditutup di checkpoint S2.** Diperlakukan sama dengan penyisihan lain: default yang diterima lewat "lanjut" (dokumen 17 D-18 dan D-48; Q-78), bukan pernyataan pemilik produk sendiri. Aturan di atas jatah: keputusan pemilik produk (S2-1 butir (iii) pilihan b): ditandai begitu jatah terlewati (§3.4). Tidak pernah dihalangi | Bila dikecualikan, penyisihan besar dapat lolos dari tanda dengan dicatat pada hari gajian. Aturan tandanya milik dokumen 06; bentuk entri dan angka tidak berubah pada jawaban mana pun |
| T-14 | Sumber pemasukan yang wajib meninggalkan catatan di buku rumah tangga | `salary` dan `business`; `gift` untuk satu orang dan `other` boleh dicatat di buku pribadi saja. **Sejak 12 November 2026** (kata-kata pemilik produk "ada kategori default bisa edit tambah atau yang lainnya", dibaca lead sebagai sumber uang masuk, **[Usulan]**): sumber berkelompok `work`, bawaannya gaji, THR, bonus, lembur, honor atau pekerjaan lepas, hasil usaha; kelompok `gift_other`: hadiah, lainnya (**[Usulan]** lead, untuk checkpoint S4; T-43) | Bila "lainnya" ikut wajib, pemasukan dana pribadi sendiri (misalnya menjual barang pribadi) ikut terlihat, dan yang tertutup tinggal hadiah |
| T-15 | "Langsung dipakai untuk sesuatu" pada layar uang masuk | Dua entri (pemasukan, lalu pemakaiannya) dalam satu alur layar | Bila harus satu entri, jenis entri `income` perlu memuat baris jenis lain, dan daftar jenis entri tidak lagi tertutup |
| T-16 | Arah "ke dana pribadi" bila yang mencatat gaji bukan penerimanya | Hanya ke dana pribadi pencatat sendiri (INV-25); dan, menurut keputusan pemilik produk di checkpoint S2, tidak tersedia sama sekali bagi pencatat yang bukan `earner` (INV-27, T-17): ia mencatat ke dana bersama, dan bagian pribadi siapa pun diambil lewat penyisihan | Bila pasangan boleh mengarahkan ke dana pribadi orang lain, perlu pengecualian T-12 |
| T-17 | Pencatat yang bukan penerima penghasilan mengarahkan sebagian penghasilan itu ke dana pribadinya sendiri (SP-01) | **Ditutup: keputusan pemilik produk, checkpoint S2** (S2-1 butir (i)): tidak boleh; hanya `actor = earner` (INV-27). Adil karena T-13 berlaku "ya" | Bila boleh: A dapat mengarahkan gaji B ke dana pribadi A dalam entri yang tidak dapat dibatalkan B; B hanya dapat berkeberatan. Angka contoh tidak berubah pada jawaban mana pun |
| T-18 | Daftar sumber penghasilan sebagai parameter preset (SP-18) | Preset pilot: `salary`, `business`; perubahan berlaku ke depan. **Sejak 12 November 2026** daftar itu adalah sumber rumah tangga berkelompok `work`, yang dapat ditambah lewat kesepakatan (D-89; T-43) | Siapa yang boleh mengubahnya: §3.1 "Sumber uang masuk" (**[Usulan]**); alurnya: dokumen 06 |
| T-19 | Klaim antar-pemilik yang terbuka saat anggota keluar (SP-12) | **[Usulan]**: ditutup entri sistem `exit_closure`, tanpa syarat dan tanpa tindakan siapa pun; tidak menyatakan pembayaran. Sisi hukum **[Perlu validasi]** (L-04) | Bila advokat menilai penutupan itu dapat dibaca sebagai pelepasan hak: klaim dibiarkan terbuka sebagai klaim kepada "mantan anggota" (pihak luar), dengan biaya kekayaan bersih rumah tangga memuat klaim yang tidak dapat lagi ditutup dari dalam aplikasi kecuali dengan penghapusan piutang |
| T-20 | Rincian per anggota atas penyisihan dan kontribusi (SP-13) | Hanya di dalam satu periode laporan; tidak ada angka kumulatif, lintas periode, atau bersih | Bila dibuka, aplikasi menghitungkan batas atas dana pribadi anggota yang dana pribadinya hanya berasal dari dana bersama |
| T-21 | Uang antara dua rumah tangga yang sama-sama memakai aplikasi (Q-86) | Dua catatan terpisah tanpa tautan | Tautan apa pun adalah saluran antar-rumah tangga dan memerlukan tinjauan `security-privacy` |
| T-22 | Siapa yang menetapkan jumlah jatah pribadi | Bukan lagi asumsi: keputusan pemilik produk, 11 November 2026 ("a di sepakati berdua, nanti saling konfirmasi aja"). Modelnya di §3.4 | — |
| T-23 | Rentang hari mulai periode anggaran | 1 sampai 28; rumah tangga yang penghasilannya tiba tanggal 29 sampai 31 memilih 28 atau 1 | Bila 29 sampai 31 harus dapat dipilih, perlu aturan "hari terakhir bulan bila tanggal itu tidak ada", dan periode tidak lagi selalu mulai pada tanggal yang sama |
| T-24 | Bentuk peralihan saat hari mulai diganti | Satu periode pendek (1 sampai 30 hari) dengan batas anggaran dan jatah pribadi penuh, tidak diprorata | Bila satu periode panjang: sampai sekitar dua bulan, memuat dua kali pembayaran bulanan terhadap batas satu bulan. Bila diprorata: pembagian menghasilkan pecahan dan aturan pembulatan baru |
| T-25 | Siapa yang boleh mengganti hari mulai periode | Putusan lead, 11 November 2026: kesepakatan rumah tangga, yaitu kedua anggota dewasa, seperti nilai ambang; terlihat oleh keduanya. Belum dinyatakan pemilik produk dengan kata-katanya sendiri | Bila satu orang boleh sendiri: ia dapat membuka periode pendek baru, dengan batas anggaran dan jatah pribadi penuh, kapan saja; penggantiannya tetap terlihat |
| T-26 | Dasar periode hitungan jatah pribadi | Ditetapkan dokumen 06 §6.3: tanggal kejadian, dinilai satu kali saat entri disimpan. Jalan menghindar yang diterima: tanggal mundur, yang hanya dilawan keterlihatan tanggal pencatatan (§3.4, §3.11) | Bila waktu simpan: penghasilan yang dicatat terlambat memakan jatah periode berjalan, dan angka jatah tidak sama dengan baris "disisihkan" ringkasan mana pun |
| T-27 | Sumber untuk penghasilan lepas (freelance) | **Ditutup di checkpoint S2**: `business` (sejak 12 November 2026 sumber bawaan tersendiri honor atau pekerjaan lepas, `freelance`, di kelompok `work` yang sama, **[Usulan]** lead untuk checkpoint S4; akibatnya tidak berubah), sehingga ia *penghasilan*, terlihat seperti gaji, dan selalu tercatat di buku rumah tangga (INV-24). Diterima pemilik produk atas saran lead (S2-1 butir (iv)), bukan pernyataannya sendiri. Catatannya untuk tahap publik: irama penghasilan tidak boleh diandaikan (§2.4) | Pilihan yang tidak diambil: dicatat sebagai `other`, boleh masuk buku pribadi saja dan tidak terlihat. Celah label §10.3 butir 1 tetap ada bagi pencatat yang memilih sumber "lainnya" |
| T-28 | Uang untuk tagihan yang jatuh temponya menyusul sesudah uang masuk (fakta rumah tangga pilot, checkpoint S2) | **Ditutup di checkpoint S3**: diterima lewat "lanjut" (dokumen 17 D-79, Q-117), bukan pernyataan pemilik produk sendiri. Usulan revisi pasca-checkpoint S2: *alokasi* per tagihan pada tempat uangnya, tanpa persetujuan dan tanpa entri uang; bukan reservasi. Jumlah yang ditawarkan: jumlah per kali kesepakatan rutin bila kesepakatan itu berlaku, selain itu perkiraan tagihan (§3.15, CTH-41) | Bila uang itu harus mengurangi *saldo tersedia* seperti reservasi, reservasi harus dapat lahir tanpa permintaan yang disetujui: `reservation` kehilangan `authorization_id` wajibnya, dan §3.8, INV-11, serta dokumen 06 berubah |
| T-29 | Mengakhiri jatah sendiri sesudah jatah menjadi batas (akibat S2-1 butir (iii) pilihan b atas default D-52) | **Ditutup: pilihan (1)**, rumusan lead yang dikukuhkan pemilik produk, 11 November 2026. Jatah yang diakhiri pemiliknya tetap menjadi dasar penilaian sampai akhir periode anggaran yang memuat pengakhirannya; sejak periode berikutnya anggota itu tidak punya jatah. Selebihnya D-52 tidak berubah: hanya pemilik jatah yang mengakhirinya. Rincian usulan `architect`, diikuti dokumen ini: batasnya waktu simpan entri di server, bukan tanggal kejadian (diterima lewat "lanjut" di checkpoint S3, D-81); jatah baru yang disepakati dalam masa itu menggantikannya (§3.4, INV-30). **Dipersempit 12 November 2026** (usulan `security-privacy`, SP-33; diterima lewat "lanjut" di checkpoint S3 sebagai bagian D-77, bukan pernyataan pemilik produk sendiri): berlaku hanya bila jatah yang diakhiri di atas ambang yang berlaku saat diakhiri; selain itu pengakhiran berlaku seketika | Pilihan yang tidak diambil: (2) mengakhiri jatah memerlukan persetujuan pasangan, yang bertentangan dengan asas SP-24; (3) anggota tanpa jatah diperlakukan seperti berjatah 0, yang membalik W-14. Bila batasnya tanggal kejadian: entri yang disimpan sesudah masa itu dengan tanggal mundur masih dinilai terhadap jatah lama. Yang tidak ditutup: §10.3 butir 7 |

**Butir baru putaran S3 (12 November 2026).** Semuanya lahir sebagai usulan putaran S3; T-39 dan T-40 lahir di putaran perbaikan S3 (fase D); T-32, T-33, dan T-37 diperbarui di putaran itu. **Semuanya ditutup di checkpoint S3 (12 November 2026)**: bawaannya diterima lewat "lanjut", bukan pernyataan pemilik produk sendiri; nomor penutupnya di kolom kedua. T-34 dibuka ulang sesudahnya oleh pertanyaan pemilik produk (baris T-34; T-43). Angka **[Asumsi]** di dalamnya tetap **[Asumsi]**.

| # | Butir | Bawaan yang diusulkan | Akibat bila lain |
|---|---|---|---|
| T-30 | Entri terlambat atas tempat uang terarsip (PA-FIN-4) | **Ditutup di checkpoint S3** (D-85, Q-168). Perincian selisih diterima tanpa apa pun; selebihnya lewat pengaktifan kembali oleh yang boleh mengarsipkan, sebagai peristiwa (§2.6) | Bila entri terlambat diterima atas tempat uang terarsip tanpa pengaktifan: tempat uang terarsip dapat bersaldo bukan nol tanpa tampil di custody, dan INV-21 gugur |
| T-31 | Satu belanja dari tempat uang dua pemilik (PA-FIN-6) | **Ditutup di checkpoint S3** (D-85, Q-169). Dua entri, satu alur layar (§3.2) | Bila satu entri: jenis entri baru yang memuat kontribusi dan pengeluaran dana bersama sekaligus |
| T-32 | Penghapusan piutang terhadap ambang (G-03; SP-51) | Dibandingkan dengan dasar `threshold`, **kumulatif per piutang**: jumlah bagian yang belum disetujui atas semua penghapusan piutang itu yang belum dibalik, termasuk entri ini; entri ditandai bila jumlah itu di atas ambang dan entri ini punya bagian yang belum disetujui. 300000 + 300000 atas piutang 600000 menandai yang kedua dan setiap penghapusan sesudahnya. Boleh diajukan dulu (§3.7, §3.8 `H`, INV-34; CTH-25 varian SP-51). **Ditutup di checkpoint S3**: S3-9 (ii) pilihan A, diterima lewat "lanjut" (Q-149) | Bila tidak dibandingkan: satu orang dapat menurunkan kekayaan bersih rumah tangga sebesar apa pun tanpa tanda, sedangkan kiriman keluarga sebesar itu ditandai. Bila per entri (rumusan putaran S3 sebelumnya): mengikhlaskan dalam potongan di bawah ambang tidak pernah ditandai |
| T-33 | Hitungan tempat uang dana bersama oleh bukan pemegangnya (G-05; SP-50) | Boleh, dengan tiga syarat: hanya rekening bank atau dompet digital yang ditandai pemegangnya "pasanganku dapat melihat saldo rekening ini" (bawaan mati, dapat dicabut pemegang); paling banyak satu hitungan bukan pemegang yang terbuka per tempat uang, hitungan berikutnya ditolak sampai pemegang menjawab; pemegangnya bukan dalam masa tenggang keluar. Selisihnya bertanda belum dikonfirmasi di sisi pemegang, positif maupun negatif (§3.10 butir 1, INV-33; CTH-47). **Ditutup di checkpoint S3**: S3-9 (i) pilihan A, diterima lewat "lanjut" (Q-148) | Bila hanya pemegang (dan siapa pun atas `joint`): pasangan yang dapat melihat saldo rekening harus meminta pemegang menghitung. Bila boleh tanpa penanda: kekurangan uang di tangan seseorang tercatat atas namanya tanpa pernyataannya. Bila tanpa tanda pemegang: pasangan dapat mencatat "uangmu kurang" atas uang yang tidak dapat ia lihat. Bila tanpa batas satu terbuka: catatan itu dapat diulang sebelum pemegang sempat menjawab |
| T-34 | THR dan bonus dari pemberi kerja (dokumen 11 §10.1 butir 5) | **Dibuka ulang 12 November 2026.** Di checkpoint S3 ditutup dengan S3-3 pilihan A lewat "lanjut" (Q-145: sumber `salary`). Sesudahnya pemilik produk bertanya "kenapa thr dicatat sebagai gaji? kan pemasukan kita yang atur darimana nya?" dan menulis "ada kategori default bisa edit tambah atau yang lainnya" (kata-katanya sendiri; bahwa "kategori" berarti sumber uang masuk adalah bacaan lead **[Usulan]**). Bawaan baru (**[Usulan]** lead, ditanyakan di checkpoint S4): THR, bonus, dan lembur adalah sumber bawaan tersendiri di kelompok `work`: berlabel "THR", tetap *penghasilan*, selalu terlihat pasangan, `earner` wajib (§3.1 "Sumber uang masuk", INV-24, INV-27). Pemindahan ke `gift_other` hanya lewat kesepakatan berdua (D-89). Angka contoh tidak berubah. Pertanyaan sampingan "apakah kalian menerima THR?": "thr bisa menerima bisa ngasih"; THR yang diberikan: T-42 | Bila THR bawaannya di `gift_other`: boleh dicatat di buku pribadi saja dan tidak terlihat; celah label §10.3 butir 1 melebar tepat pada uang masuk tahunan terbesar |
| T-35 | Pembelian cicilan dengan uang muka + cicilan < harga (G-07) | **Ditutup di checkpoint S3** (D-85, Q-170). Ditolak; harga diisi sebesar yang dibayar; penilaian ulang bila perlu (§3.14) | Bila diterima: biaya kredit negatif menaikkan kekayaan bersih pada tanggal beli tanpa peristiwa uang |
| T-36 | Tagihan berulang sesudah tanggal akhirnya | **Ditutup di checkpoint S3** (D-85, Q-171). Tidak terarsip sendiri; tidak ada kejadian baru dan pengingat berhenti; diarsipkan anggota (§3.15) | Bila terarsip sendiri: ada peristiwa tanpa pelaku atas kesepakatan rutinnya (dokumen 06 §6.3 baris (d)), dan pembayaran terlambat atas kejadian terakhir memerlukan pengaktifan kembali |
| T-37 | Buku pribadi anggota yang keluar tanpa meminta penghapusan (PA-FAM-7) | **Ditutup di checkpoint S3** (D-86, Q-172; S3-4 (c)). Buku itu ditutup pada akhir masa tenggang keluar (`exit_closure`; T-39); dihapus pada akhir jangka retensi dokumen 10; sampai itu hanya dapat diekspor pemiliknya (§8.2 butir 2a) | Bila tetap terbuka: klaim antar-pemilik tidak pernah dapat ditutup, dan tidak ada pelaku yang berhak menulis buku itu |
| T-38 | Menyelesaikan atau membatalkan tujuan yang beralokasi (G-01) | **Ditutup di checkpoint S3** (D-85, Q-173). Arsip hanya bila alokasi nol; pelepasan ditawarkan di langkah yang sama tanpa pilihan terisi; tidak ada pelepasan otomatis (§3.9) | Bila pelepasan otomatis saat diarsipkan: uang yang disiapkan untuk tujuan itu menjadi bebas tanpa pernyataan siapa pun |
| T-39 | Masa tenggang keluar, sisi uang (SP-48; SP-49) | 7 hari **[Asumsi]**; akses berakhir seketika; `exit_closure`, penutupan, dan penghapusan pada akhir masa itu; hanya yang keluar dapat membatalkan; buku pribadinya dibekukan selama itu, klaim tetap terbuka dan tampil; angka `exit_closure` tidak berubah, hanya tanggalnya; ekspor buku rumah tangga sampai saat akses berakhir; anggota yang bergabung kemudian tidak membaca rincian tulisan mantan anggota (§8.2 butir 1a, 2b, 3a; INV-28; CTH-38). **Ditutup di checkpoint S3**: S3-4 (c) pilihan A, D-86, D-90, diterima lewat "lanjut" (Q-179); 7 hari tetap **[Asumsi]** (A-39) | Bila tanpa masa tenggang: keluar yang disesali tidak dapat dibatalkan karena `exit_closure` dan penghapusan tidak dapat dibalik. Bila buku pribadi tidak dibekukan: angka penutupan bergantung pada apa yang dicatat orang lain selama yang keluar tidak dapat melihatnya |
| T-40 | Kejadian bawaan bagi pembayaran tagihan yang tidak punya kejadian belum dibayar sampai akhir periodenya (PA-FIN-11) | **Ditutup di checkpoint S3** (D-97, Q-182). Kejadian paling awal yang belum dibayar sesudahnya; bila tidak ada, kejadian terakhir; pencatat dapat menggantinya (§3.15) | Bila kosong: pencatat yang membayar lebih awal harus memilih kejadian setiap kali, dan tautan tanpa kejadian tidak sah (INV-26) |

**Butir baru sesudah checkpoint S3 (12 November 2026).** Semuanya **[Usulan]**, belum disajikan kepada pemilik produk. T-41 dan T-42 (`finance-domain`) menurut dokumen ini tidak menyisakan pilihan nyata bagi pemilik produk (jawaban mana pun tidak mengubah angka contoh), jadi cukup dicatat sebagai bawaan di dokumen 17, bukan pertanyaan checkpoint; pemilik dokumen 17 yang menilainya. T-43 adalah bawaan lead yang memang ditanyakan di checkpoint S4, karena pemilik produk belum menjawab apakah sumber buatan sendiri boleh tidak terlihat.

| # | Butir | Bawaan yang diusulkan | Akibat bila lain |
|---|---|---|---|
| T-41 | Mengganti ada-tidaknya `due_rule` pada pos rutin yang sudah dibuat (IDEA-03) | Tidak dapat diganti; pos diarsipkan dan pos baru dibuat dengan kesepakatan baru (§3.15) | Bila dapat diganti: perlu aturan kejadian sebelum dan sesudah penggantian, dan aturan bagi pembayaran lama yang tidak menunjuk kejadian |
| T-42 | THR dan angpau yang diberikan (pernyataan pemilik produk, 12 November 2026: "thr bisa menerima bisa ngasih"; "bro pikirkan kalau semua orang pakai app ini") | `expense` biasa dengan penerima manfaat pihak luar mana pun; kiriman keluarga bila kepada keluarga besar; ambang per entri seperti kiriman keluarga; musim pemberian lewat satu permintaan musiman atau alokasi; dari dana pribadi menurut pilihan pemberi (§3.2, CTH-51) | Bila dinilai kumulatif per musim terhadap ambang: perlu batas "musim" yang tidak ada di model (tanggal hari raya berbeda per agama dan per tahun) dan tanda yang bergantung pada urutan tiba amplop; bila THR diberi jenis entri sendiri: daftar jenis entri tertutup bertambah tanpa angka yang berbeda |
| T-43 | Sumber uang masuk yang dapat disunting (kata-kata pemilik produk, 12 November 2026: "ada kategori default bisa edit tambah atau yang lainnya"; dibaca lead sebagai sumber uang masuk, **[Usulan]**); yang tidak ia jawab: apakah sumber buatan sendiri boleh dicatat tanpa terlihat pasangan | **[Usulan]** lead untuk checkpoint S4, dengan rincian dokumen ini (§3.1 "Sumber uang masuk"): tiap sumber termasuk kelompok `work` (selalu terlihat, INV-24, INV-27) atau `gift_other` (boleh di buku pribadi); bawaan `work`: gaji, THR, bonus, lembur, honor atau pekerjaan lepas, hasil usaha; `gift_other`: hadiah, lainnya. Menambah dan mengganti nama sumber `gift_other`: sendirian, sebagai peristiwa yang terlihat, tidak dengan nama yang pernah dipakai sumber `work`. Menambah, mengganti nama, dan mengarsipkan sumber `work`, dan memindahkan sumber antarkelompok: kesepakatan berdua, ke depan saja (D-89; CP-05). "Hadiah" dan "lainnya" tetap di `gift_other` dan tidak diarsipkan. Nama unik; kelompok selalu tampil di samping nama | Bila sumber buatan sendiri selalu terlihat: pemasukan pribadi yang bukan hadiah (menjual barang pribadi) harus memakai "lainnya". Bila memindahkan kelompok boleh sendirian: satu orang dapat membuat THR (miliknya maupun pasangannya) ke depan boleh dicatat tertutup, atau mewajibkan pemasukan pribadi pasangan terlihat, hanya dengan peristiwa. Bila menambah sumber `work` boleh sendirian (lebih ringan, karena tidak mengikat siapa pun): bertentangan dengan bunyi D-89 |

Butir dokumen 17 yang dapat ditutup atau diperbarui: K-02, K-11, K-13 (selesai lewat dokumen ini dan glosarium); Q-11 sampai Q-21, Q-43 sampai Q-48 (ditetapkan dokumen 07; diterima sebagai default di checkpoint S2 lewat "lanjut"); Q-22 (keputusan sementara pemilik produk di checkpoint S2, S2-8; menunggu L-04). Sesudah checkpoint S2: T-13, T-16, T-17, T-27 ditutup; T-03 diperbarui; T-28 dan T-29 baru. Sesudah tiga penegasan 11 November 2026: T-29 ditutup (pilihan 1). Sesudah checkpoint S3 (12 November 2026): T-28 dan T-30 sampai T-40 ditutup lewat "lanjut"; penyempitan T-29 (SP-33) dan batas waktu simpannya diterima lewat D-77 dan D-81; T-10 bagian putaran S3 diterima lewat D-85. Pertanyaan sampingan S3-3 dijawab sesudah checkpoint ("thr bisa menerima bisa ngasih"); akibatnya T-42. T-34 **dibuka ulang** oleh pertanyaan pemilik produk atas S3-3 ("kenapa thr dicatat sebagai gaji? ..."; "ada kategori default bisa edit tambah atau yang lainnya"); akibatnya T-43, ditanyakan di checkpoint S4. T-41, T-42, T-43 baru, **[Usulan]**.

### 10.2 Permintaan

**Kepada `architect`.**

1. ADR-A sampai ADR-G sudah ditulis sebagai ADR-002 sampai ADR-008. Kalimat yang kini tidak lagi sesuai dengan dokumen ini:
   - ADR-003, Konsekuensi: "Celah yang tidak ditutup model: uang masuk yang langsung dicatat sebagai dana pribadi tidak terlihat (07 §10.3 butir 1; K-16). Keputusannya milik pemilik produk di checkpoint S2 (dokumen 06 §6.4)." Sudah diputuskan; penghasilan tidak dapat lagi dicatat hanya di buku pribadi (INV-24). Yang tersisa adalah celah label sumber (§10.3 butir 1).
   - ADR-003, Keputusan: perlu memuat aturan tulis buku pribadi (INV-25), dan baris "Rincian" perlu menyebut INV-24 dan INV-25.
   - ADR-002, tabel Opsi: "Dua tabel dan daftar jenis tertutup". Kini ada tabel ketiga yang hanya ditambah, `entry_link`; dan jenis `income` boleh memuat baris di dua buku.
   - ADR-004, Keputusan: "dihitung dari baris dan dari otorisasi serta reservasi yang berlaku". Realisasi kini juga bergantung pada tautan susulan yang berlaku.
   - ADR-005, Konsekuensi: "jumlah entri, baris, otorisasi, reservasi, peristiwa, dan notifikasi tidak bertambah". Tambahkan tautan susulan.
   - ADR-006, ADR-007, ADR-008: tidak ada kalimat yang berubah.
2. Dokumen 06 memakai kait berikut dan tidak mendefinisikan ulang akibat uangnya: `request.approved`, `authorization.amended`, `authorization.settled`, `authorization.cancelled`, `authorization.expired`, `authorization.overrun`, `handover.recorded`, `handover.confirmed`, `handover.disputed`, `entry.disputed` (§3.3, §3.8, §3.12). Jumlahnya sepuluh. Transisi dokumen 06 yang **tidak** memicu kait uang: penutupan keberatan atas baris custody tanpa entri pembalik (C05), yang juga bukan konfirmasi (§3.3); `handover.confirmed` hanya dipicu C02 dan C03.
3. Dokumen 06 memutuskan: jenis permintaan untuk pemakaian dana bersama yang bukan pengeluaran (pinjaman, pembelian aset, penyisihan); apakah kelebihan yang dibayar dari dana bersama memerlukan tanggapan; penandaan pemakaian di atas ambang yang dicatat tanpa permintaan; tanggal yang menentukan periode *terikat* suatu otorisasi (§6.1 memakai tanggal rencana, atau tanggal persetujuan bila tidak ada).
4. Syarat dari model untuk dokumen 06 dan 10: tempat uang berpemilik `shared` tidak dapat bervisibilitas `private`; entri buku pribadi hanya dapat dikoreksi pemiliknya; hanya pemilik piutang yang dapat melepas klaim.
5. Jawaban atas kedelapan permintaan dokumen 06 §15.4 ada di tabel §9. Yang perlu disesuaikan di dokumen 06 sesudahnya: (a) §6.2 mengenali penyisihan dari jenis entri; bagian pribadi di dalam entri `income` adalah baris `owner_transfer`, jadi syaratnya perlu membaca baris, bukan jenis entri (T-13). (b) §6.4 sudah dijawab pemilik produk: hasilnya pilihan 3 dalam keterlihatannya, dengan bentuk satu entri. (c) selesai: dokumen 06 §6.2 dan §6.3 kini menetapkan bahwa pelunasan tidak ditandai dan cicilan tidak memerlukan kesepakatan rutin; dokumen ini diselaraskan (§2.6, §3.8, §3.14). (d) §8.1 "entri dua buku: hanya pemilik buku pribadinya" dipertegas di §3.12: itu berlaku atas entri pembalik; penyesuaian sisi rumah tangga tetap terbuka bagi keduanya. (e) §7.3: aturan dikukuhkan; tambahkan pengecualian pelepasan klaim dan penundaan entri antara dua buku pribadi.
6. Jawaban atas keempat permintaan baru dokumen 06 §15.4, butir (a) sampai (d), ada di tabel §9; keempatnya diterapkan atau dikukuhkan. Hal yang semula diserahkan kepada dokumen 06 (kewajiban yang didaftarkan sebagai saldo awal sesudah rumah tangga mulai mencatat) sudah diputuskan di dokumen 06 §6.2, putusan 2, dan diterapkan di §3.8 dan §3.14.
7. Jawaban atas permintaan dokumen 06 §15.4 butir (e) sampai (m) ada di tabel §9. **Selesai** (diperiksa pada dokumen 06 yang berlaku saat revisi ini): §6.2 kini memuat kedua pengecualian `contribution_expense` dan `advance`, mengukuhkan baris `expense` di dalam `recategorization`, dan menyatakan ketimpangan `loan_received` sebagai disengaja; §6.4 kini berbunyi menurut pernyataan bahwa kedua pasangan berpenghasilan; kalimat "satu-satunya pengecualian" atas INV-25 tidak ada lagi di dokumen 06; jatah pribadi kumulatif per periode ada di §6.3. Tidak ada yang tersisa dari butir ini.
8. Sisi uang temuan pemeriksaan awal `security-privacy` sudah diterapkan di dokumen ini (tabel §9). Di dokumen 06, SP-01, SP-05, SP-13, dan SP-22 **selesai**, begitu pula ADR-003 (INV-25 dengan dua pengecualian). Yang masih perlu diperiksa `architect`:
   - SP-12: ADR-007 keputusan 3 dan ADR-002 belum menyebut entri sistem `exit_closure` (§8.2, INV-28): penghapusan buku pribadi tidak bersyarat pada penyelesaian klaim oleh anggota; jenis entri bertambah satu; `actor` boleh `system` hanya pada jenis itu.
   - SP-18: siapa yang boleh mengubah *daftar sumber penghasilan* dan lewat alur apa (usulan: kesepakatan rumah tangga sebagai bagian preset aturan, dokumen 06 §6.3 "Kesepakatan rumah tangga"); dokumen 06 belum menyebut daftar itu.

9. Keputusan pemilik produk 11 November 2026 tentang jatah pribadi dan hari mulai periode anggaran. Yang disediakan dokumen ini, dan yang perlu ditetapkan dokumen 06:
   - Sudah sama dengan dokumen 06 §6.2 dan §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi", dan menjawab §15.4 butir (n) sampai (s): (n) setiap penyisihan kepada M terhitung dengan sendirinya, tanpa tautan; jumlahnya per periode anggaran (§3.4). (o) Kolom jumlah pemakaian tabel §3.8 tidak berubah; jumlah pemakaian entri `owner_transfer` dan bagian pribadi entri `income` seluruhnya penyisihan (§3.4). (p) Kedua sifat periode dijamin INV-29; penggantian hari mulai adalah kesepakatan rumah tangga dan terlihat (§6.1). (q) CTH-39. (r) Dasar yang sama dengan tabel A; entri pembalik mengurangi baris asalnya (§3.4, §6.3). (s) `recorded_at` disimpan server; tanggal pencatatan (tanpa jam) tampil hanya pada entri yang tanggal pencatatannya berbeda dari tanggal kejadiannya, dan jam pencatatan tidak ditampilkan kepada siapa pun (§3.11; SP-28).
   - Yang masih perlu di dokumen 06: §6.3 mendaftar isi kesepakatan `rule` sebagai nilai ambang dan preset aturan. Menurut putusan lead, **hari mulai periode anggaran** masuk daftar itu (K01 sampai K05; hari mulai lama berlaku sampai versi baru disetujui), dengan guard 1 sampai 28.
   - Nama: `luar_jatah` di sini adalah `lebih` di dokumen 06, dan `terpakai_sebelum` adalah `jumlah_sebelum`; rumusnya setara (§3.4).
   - §4.5 "Masa berlaku dan kedaluwarsa", baris "Pembayaran atas kewajiban yang sudah tertaut penuh": kini juga bagian dari INV-26.
   - ADR dan bagian dokumen 06 yang masih berbunyi "mulai tanggal gajian" atau "jumlah per kali" untuk penyisihan perlu mengikuti.
10. **Sesudah checkpoint S2** (revisi yang disetujui pemilik produk). Yang perlu berubah di dokumen 06 dan ADR supaya sama dengan dokumen ini:
   - **Jatah sebagai batas** (S2-1 butir (iii) pilihan b; §3.4). Dokumen 06 §6.2 dan §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi"; INV-A17, INV-A37, INV-A38: bagi anggota yang punya jatah berlaku, entri ditandai bila `jumlah_sebelum + bagian > jatah` (`lebih > 0`), tanpa membandingkan `lebih` dengan ambang; di dalam jatah tanpa tanda dan tanpa pemberitahuan; tanpa jatah tidak berubah. "Bentuk tengah" dan pilihan (a) dicatat sebagai pilihan yang ditolak. Kalimat §6.3 "Akibat lain yang diterima" tentang urutan tiba memakai contoh baru (§10.3 butir 7).
   - **Jatah berlaku sebesar 0 lawan tanpa jatah** (§3.4, **[Usulan]**): keduanya kini berakibat berbeda dan perlu dibedakan dalam penilaian.
   - **K06 atas jatah pribadi** (D-52): dasar "mengakhiri jatah sendiri hanya memperketat" tidak lagi benar (§3.4; T-29). Dokumen ini tidak mengubah aturannya; mohon dinilai ulang bersama `security-privacy`.
   - **INV-27 dan hadiah pribadi** kini keputusan pemilik produk: label **[Usulan]** di dokumen 06 §6.4 dan rujukan "menunggu checkpoint S2" dicabut. S2-1 butir (v): rujukan dokumen ini ke alasan kenaikan jatah tidak berubah (alasan opsional adalah aturan dokumen 06).
   - **Tagihan berulang** (§3.15): pernyataan pemilik produk "untuk membuat tagihan rutin ya disepakati dulu" dapat dibaca lebih luas daripada dokumen 06 §6.3, yang hanya mensyaratkan kesepakatan bagi tagihan yang jumlah per kalinya di atas ambang (§10.3 butir 9). Kiriman orang tua bukan pembayaran rutin di rumah tangga pilot; contoh di dokumen 06 §7.4 yang menyebutnya rutin perlu ditandai ilustrasi.
   - **Dompet campuran** (§2.3, §3.10): tidak ada kait atau alur baru. Yang perlu dijaga dokumen 06 dan 10: jumlah seluruh isi dompet tidak pernah menjadi data yang terlihat pasangan.
   - **ADR-003**: tambahkan bahwa satu dompet fisik boleh dicatat sebagai dua tempat uang berpemegang sama dan bahwa itu kasus biasa. **ADR-007**: status tetap sementara; dasarnya kini keputusan sementara pemilik produk (S2-8), bukan usulan.
   - **Alokasi untuk tagihan** (T-28): tidak mengubah dokumen 06; reservasi tetap hanya lahir dari permintaan yang disetujui.
11. **Sesudah revisi ketujuh dokumen 06** (tiga penegasan 11 November 2026). Butir 10 **selesai** kecuali ADR: jatah sebagai batas, jatah 0, K06 atas jatah pribadi, pos rutin, dan dompet campuran (V-9, INV-A41) kini sama di kedua dokumen. Permintaan dokumen 06 §15.4 butir (t) dan (u) diterapkan seluruhnya (tabel §9); tidak ada yang ditolak. Yang masih perlu dilihat `architect`:
   - **Pelaku tautan pada *konfirmasi susulan*.** §3.8 "Penautan" butir 3 mengutip INV-A34: tautan susulan ke otorisasi hanya dari pencatat entri atau pemohon otorisasi. Tautan konfirmasi susulan lahir di dalam transaksi persetujuan, yang pelakunya pemberi persetujuan. Dokumen ini membacanya sebagai tautan atas nama pemohon, yang menyebut entrinya saat mengajukan; mohon dinyatakan di dokumen 06 siapa `actor` pada `entry_link` itu, dan apakah pemohon harus pencatat entri yang dikonfirmasi. Tidak ada akibat uang pada jawaban mana pun. **Dijawab** (dokumen 06 §15.4 butir (v)): pelakunya pemohon, dan pemohon harus pencatat entri itu; diterapkan di §3.8 dan INV-26.
   - **Entri berjenis selain `expense` yang ditautkan ke tagihan berulang.** Dokumen ini mengandaikan pembayaran tagihan berulang selalu membawa baris `expense` (§3.15); penyisihan tidak ditautkan ke tagihan berulang, sehingga paling banyak satu dasar tanda yang benar per entri (§3.8). Bila dokumen 06 mengizinkan tautan jenis lain, mohon disebut. **Dijawab** (dokumen 06 §15.4 butir (v)): `expense`, `contribution_expense`, `advance`. Diterapkan di §3.15 aturan 7a; dokumen ini menambahkan bahwa `contribution_expense` tertaut tidak pernah ditandai (**[Usulan]**), mohon disamakan di dokumen 06.
   - **Putaran perbaikan 12 November 2026**, untuk diselaraskan di dokumen 06: (a) fungsi tanda §3.8 kini memakai `N = P − dalam_batas` bagi entri yang tertaut ke otorisasi (aturan lead dari D-07; `critic` S7); mohon nama dan rumusnya disamakan, juga dua hal yang **[Usulan]** dokumen ini: kelebihan dinilai per entri dan tidak dijumlahkan, dan entri yang tertaut ke otorisasi sekaligus ke tagihan berulang dinilai menurut otorisasinya saja. (b) Konfirmasi susulan: disetujui, ditautkan, dan ditutup dalam satu transaksi sah dari sisi uang; dokumen ini menambahkan bahwa otorisasinya tidak menerima tautan lain dan permintaannya tidak menyebut tempat uang sumber (§3.8, INV-26). (c) Guard K01 atas jatah yang tidak di atas ambang dan `assessment_until` yang hanya ada bila jatah di atas ambang: dokumen ini mengutipnya (§3.4, INV-30). (d) V-9: kalimat "server tidak menerima `C`" bukan perlindungan terhadap pengelola server (§2.3).
   - **ADR-009** (aturan permintaan menandai, bukan menghalangi) perlu memuat tiga dasar tanda dan bahwa pembayaran di bawah ambang dapat bertanda; **ADR-003**: butir 10 masih berlaku.

**Kepada lead (glosarium di luar bagian 7).**

1. Bagian 4, *permintaan penggantian*: sesuaikan dengan keputusan pemilik nomor 3. Usulan bunyi: "Meminta dana bersama mengganti uang pribadi yang dipakai untuk keperluan bersama dan ditandai *minta diganti*."
2. Bagian 4, *persetujuan* dan *kedaluwarsa*: rujuk §3.8 untuk akibat uangnya.
3. Bagian 3, *data pribadi anggota*: tambahkan bahwa buku pribadi (dana pribadi beserta entri, piutang, kewajiban, dan asetnya) termasuk di dalamnya, kecuali sisi yang menyentuh dana bersama (§5.3).
4. `_progress.md`: temuan N-14, K-11, K-13 ditutup oleh dokumen ini.
5. Dokumen 17: T-12 sampai T-16 sebagai butir baru; K-16 dan Q-56 dijawab (§9); T-08 bukan lagi asumsi.
6. Glosarium bagian 4, *rujukan induk* (subject reference) belum menjadi istilah; dokumen ini merujuknya ke dokumen 06 §2.2.
7. Glosarium bagian 4, *konfirmasi custody*: definisinya hanya menyebut penerima. Tambahkan sisi pembayar: catatan entri dua buku yang mengurangi tempat uang dana bersama yang dipegang anggota dewasa selain pencatatnya juga bertanda *belum dikonfirmasi* (§3.3).
8. Glosarium bagian 3, *keluar dengan aman*: rujuk §8.2 untuk klaim antar-pemilik yang ditutup sistem saat keluar.
9. Dokumen 17: T-17 sampai T-21 sebagai butir baru; INV-27 menjadi pertanyaan pemilik produk di checkpoint S2 bersama S2-1 (keduanya saling bergantung, §3.1); sisi hukum `exit_closure` masuk daftar tinjauan hukum bersama L-04; Q-86 dijawab sebagai **[Usulan]** (§2.2).
10. CTH-38 adalah contoh berdiri sendiri baru untuk skrip hitung ulang; tidak ada angka CTH-01 sampai CTH-37 yang berubah.
11. Penyelarasan terakhir (dokumen 06 §15.4 butir (e) sampai (m); kedua pasangan berpenghasilan) tidak mengubah satu angka pun di tabel mana pun. Rujukan lama ke nomor baris dokumen 17 dicabut: rumusan bersyarat yang tersisa di sana ("hanya satu orangnya menerima gaji", "tidak berpenghasilan") dinilai dokumen 17 bagian 5 tetap benar sebagai kalimat umum, dan kalimat satu pencari nafkah di dokumen 06 dan dokumen ini sudah diperbaiki (K-27 butir 1, selesai). Tidak ada yang tersisa dari butir ini.

12. Revisi jatah pribadi dan periode anggaran: CTH-39 adalah contoh berdiri sendiri baru untuk skrip hitung ulang (dua tabel, 17 baris); tabel peta kasus §7.1 bertambah satu baris; tidak ada angka CTH-01 sampai CTH-38 yang berubah. Glosarium bagian 4, *kesepakatan rutin*: untuk jatah pribadi jumlahnya per periode anggaran, bukan per kali. `_progress.md` dan dokumen 17 (Q-06, Q-98; T-22 sampai T-27) masih memuat "mulai tanggal gajian" dan jatah sebagai usul yang belum dijawab (keadaan saat butir ini ditulis, sebelum checkpoint S2; sejak itu bentuk jatah diputuskan pemilik produk dan butir ini selesai).

13. Sesudah checkpoint S2: CTH-40 (lima tabel) dan CTH-41 (tiga tabel) adalah contoh berdiri sendiri baru untuk skrip hitung ulang; CTH-39 mendapat kolom "Ditandai" dan satu tabel varian; contoh urutan tiba di §10.3 butir 7 berubah. Tidak ada angka saldo CTH-01 sampai CTH-39 yang berubah. Daftar lengkap angka dan tanda yang berubah ada di §11. Glosarium di luar bagian 7: bagian 4, *tanda di atas ambang* dan *kesepakatan rutin* perlu memuat tanda karena jatah terlewati (bukan hanya karena ambang); bagian 4, *ambang permintaan*: tidak dibandingkan atas penyisihan anggota yang punya jatah berlaku. Brief §1.1 masih menyebut kiriman Rp3.000.000 kepada ibu suami sebagai bagian arus uang; pemilik produk menyatakan kiriman orang tua tidak rutin (pertentangan sumber untuk dokumen 17 bagian 5).

14. Revisi kedua pasca-checkpoint S2 (tiga penegasan 11 November 2026; dokumen 06 revisi ketujuh): **tidak ada angka saldo yang berubah di contoh mana pun**. Yang berubah hanya tanda dan tabel penilaiannya: CTH-05 (tabel penilaian baru, tiga baris), CTH-41 (tabel tagihan mendapat kolom `R`, `P > R`, dan Tanda; satu entri kini bertanda), teks CTH-06 dan CTH-39. Daftar lengkap di §11. Glosarium di luar bagian 7: bagian 4, *kesepakatan rutin* (setiap tagihan berulang, bukan hanya yang di atas ambang; jumlah per kali adalah batas), *tanda di atas ambang* (kini tiga dasar; namanya tidak lagi mencakup semua tanda), dan *ambang permintaan* (tidak dibandingkan atas pembayaran tagihan berulang yang kesepakatannya berlaku).

**Kepada `product-analyst` (dokumen 17), sesudah checkpoint S2.** T-13, T-16, T-17, T-27 ditutup; T-03 diperbarui; T-28 dan T-29 butir baru, dan T-29 memerlukan pemilik produk. W-13 ditulis ulang menurut §10.3 butir 7 (cara (1) gugur; contoh urutan tiba baru; jalan keluar lewat mengakhiri jatah sendiri). W-14 perlu dibaca ulang: tanpa jatah, pengambilan yang masing-masing tidak melebihi ambang tidak ditandai sama sekali, jadi anggota tanpa jatah kini **lebih jarang** ditandai atas jumlah kecil daripada anggota berjatah. S2-6: bawaan "selisih ke bagian bersama" belum diterima pemilik produk (keadaan saat itu; **dicabut** 12 November 2026, T-03). Dua bacaan atas kata-katanya dicatat di §10.3 butir 9. **Sesudah tiga penegasan 11 November 2026:** T-29 ditutup (pilihan 1); kedua bacaan di §10.3 butir 9 sudah dijawab; W-13 perlu memuat bahwa jalan "mengakhiri jatah sendiri" tertutup sampai akhir periode; butir baru untuk dokumen 17: jumlah per kali yang disepakati untuk tiap tagihan di contoh adalah **[Asumsi]** (CTH-05, CTH-41), dan rincian pos rutin yang tidak disebut pemilik produk (§3.15 aturan 5 sampai 7) adalah **[Usulan]** `architect` yang belum disajikan kepadanya (dokumen 06 P-20). Syarat layar bagi dokumen 05: saat membuat tagihan berulang yang jumlahnya berubah-ubah, layar meminta jumlah tertinggi yang wajar, bukan perkiraannya (§10.3 butir 10).

**Putaran penyebaran S3 (12 November 2026).**

*Kepada `architect` (dokumen 06; bagian bersama).*

1. **G-02.** Sisi uang di §3.10 "Hitungan yang tiba belakangan": perintah rekonsiliasi membawa saldo terhitung, server menghitung selisih. Yang diandaikan dari dokumen 06 dan 09: perintah satu perangkat tiba menurut urutan pencatatannya; selisih yang tampil di perangkat adalah perkiraan; pencatat diberi tahu bila selisih server berbeda; tawaran "bagian dari selisih" sesudah sinkron dijalankan sebagai dua perintah yang sudah ada.
2. **G-05.** Sisi uang di §3.10 butir 1 dan §2.5 baris (c): bukan pemegang boleh merekonsiliasi tempat uang dana bersama, dan selisihnya bertanda di sisi pemegang, positif maupun negatif. Yang diandaikan: alur konfirmasi dan keberatan dokumen 06 §7 berlaku atas baris itu tanpa kait baru; pemberitahuan kepada pemegang (N-13 atau lainnya) dan apakah layar menawarkan menghitung **uang tunai** yang dipegang orang lain (PRN-10; dokumen 11 IDEA-01 membatasinya pada pemegang) diputuskan dokumen 06 dan 08.
3. **Pos rutin baris (d) dan pengarsipan otomatis.** Dokumen ini mencabut "tagihan terarsip sendiri" (§3.15, T-36); dengan itu tidak ada peristiwa tanpa pelaku atas kesepakatan rutin. Dari sisi uang, tagihan terarsip tidak menerima tautan baru, jadi kesepakatannya tidak perlu diakhiri saat pengarsipan; apakah 06 mengakhirinya (dan dengan pelaku siapa) bebas dipilih 06 (G-04).
4. **PA-FIN-1 dan pengaktifan kembali.** §2.6 merujuk "yang boleh mengarsipkan" ke dokumen 06; mohon pelaku membuat, mengganti pemegang, mengarsipkan, dan mengaktifkan kembali tempat uang dana bersama ditetapkan bersama, dan pengaktifan kembali tercatat sebagai peristiwa yang terlihat seperti pengarsipan.
5. **PA-FIN-5.** "Belum dibayar" didefinisikan per kejadian (§3.15); pengingat tagihan di dokumen 06 §13.4 sebaiknya memakainya. Tautan ke tagihan berulang kini membawa `bill_occurrence_due`; tautan ke tagihan terarsip ditolak (INV-26).
6. **G-03.** `write_off` kini punya jumlah pemakaian dan dasar `threshold` (§3.8); mohon dokumen 06 §6.2 memuatnya dan menetapkan pelakunya.
7. **PA-FAM-7.** `exit_closure` dan penutupan buku pribadi terjadi pada setiap keluarnya anggota (§8.2 butir 2a, INV-28); ADR-007 perlu membedakan penutupan dari penghapusan.

*Kepada `product-analyst` (dokumen 04, 05, 17).*

1. **04 F-FIN-012**: ganti "Mencatat cicilan dengan sisa pokok dan biaya kredit" menjadi "Mencatat cicilan dengan sisa yang harus dibayar (jumlah seluruh pembayaran yang tersisa, tanpa memisahkan pokok dan bunga) dan, untuk pembelian yang dicatat di aplikasi, biaya kreditnya" (07 §3.14).
2. **05 `FIN-1.md`**: FR-FIN-002/AC-4 dan BR terkait: PA-FIN-2 dan PA-FIN-3 kini aturan; tambahkan AC penyesuaian saldo awal dan penolakan pembalik atas saldo awal (CTH-42). FR-FIN-004/AC-6: ganti "ditolak" dengan dua jalur: perincian selisih diterima; belanja lain diterima sesudah pengaktifan kembali (CTH-43); tambahkan FR atau AC "mengaktifkan kembali". FR-FIN-030, FR-FIN-031: "belum dibayar periode ini" menurut kejadian tagihan (§3.15), bukan menurut tanggal pembayaran; jumlah yang masih ditawarkan dikurangi alokasi tujuan tagihan itu. FR-FIN-042: PA-FIN-6 kini aturan (CTH-44). FR-FIN-052/AC-6 dan FR-FIN-080/AC-7: kategori diisi dari baris biaya admin terakhir atau dari kategori cicilan; tanpa kategori tampil di baris "tanpa kategori" tabel B. FR-FIN-090/AC-5: keterangan satu kali di layar anggaran. FR-FIN-001: PA-FIN-10 kini aturan (INV-32). Hapus label **[Perlu aturan]** PA-FIN-2 sampai PA-FIN-8 dan PA-FIN-10; PA-FIN-1 dan PA-FIN-9 tetap.
3. **05 `FIN-2.md`**: FR-FIN-128 bawaannya diambil (§3.9; CTH-45); hapus **[Perlu aturan]** G-01. FR-FIN-135: tambahkan AC offline G-02 (CTH-46) dan AC hitungan oleh bukan pemegang dengan penanda (CTH-47). FR-FIN-117: tambahkan AC tanda dasar `threshold` atas penghapusan di atas ambang (CTH-25 teks). FR-FIN-111/AC-4: tetap, tambahkan AC cicilan 0% tanpa baris biaya kredit (CTH-49). FR-FIN-107: tagihan tidak terarsip sendiri sesudah tanggal akhir. Tabel D (FR yang memuatnya): tujuan terarsip tampil di periode pelepasannya.
4. **05 `FAM.md`**: PA-FAM-7 bawaannya diambil (§8.2 butir 2a); jangka retensi dan cara mengekspor tetap milik dokumen 10.
5. **Dokumen 17**: T-30 sampai T-38 sebagai butir baru; T-10 diperbarui; tiga pertanyaan checkpoint S3 di §10.4 di bawah.
6. **Dokumen 11** IDEA-01 baris (7) "Kerumitan": bentuknya satu entri `expense` bertaut rekonsiliasi, bukan `cash_difference` lalu `recategorization` (§3.10; CTH-48). IDEA-03: jawaban di §3.15 "Tagihan berulang tanpa `due_rule`". IDEA-08: THR dan bonus adalah `salary` (§3.1).

*Kepada lead (glosarium di luar bagian 7).* Bagian 4, *arsip*: tambahkan bahwa butir terarsip dapat diaktifkan kembali oleh yang boleh mengarsipkannya, sebagai peristiwa tanpa entri (§2.6). Tabel pengenal README: INV-31, INV-32, T-30 sampai T-38, CTH-42 sampai CTH-49.

**Putaran perbaikan S3, fase D (12 November 2026).**

*Kepada `architect` (dokumen 06 dan ADR).*

1. **SP-51, 06 §6.2 dasar `threshold`.** Mohon dimuat: untuk `write_off`, syaratnya `N > 0` dan `H + N > ambang`, dengan `H` = jumlah `N` penghapusan lain atas piutang yang sama yang tersimpan lebih dulu dan belum dibalik (07 §3.8 "Tanda atas entri", INV-34). Jenis lain tidak berubah.
2. **SP-50, 06 §7.2 dan §8.1.** Mohon dimuat: atribut tempat uang `others_may_count` (bawaan mati; hanya pada tempat uang dana bersama berjenis `bank` atau `ewallet` yang dipegang satu anggota; diubah hanya pemegangnya; *peristiwa* metadata yang terlihat, tanpa entri); hitungan bukan pemegang ditolak bila tanda mati, bila tempat uang itu tunai, bila pemegangnya dalam masa tenggang keluar, atau bila masih ada hitungan bukan pemegang yang terbuka atasnya; *terbuka* berakhir saat pemegang mengonfirmasi (C02, C03), berkeberatan, atau mencatat hitungannya sendiri atas tempat uang itu (07 §3.10 butir 1, INV-33). Batas N-14 adalah aturan 06.
3. **SP-48, 06 §3.6 (keluar).** Dari sisi uang dokumen ini mengandaikan: akses berakhir seketika; `exit_closure` (pelaku `system`) dan penutupan buku pribadi pada akhir masa tenggang 7 hari **[Asumsi]**; pembatalan hanya oleh yang keluar di dalam masa itu; selama masa itu yang keluar tetap `member(M)` dalam model uang (terhitung anggota dewasa untuk syarat tanda, tetap pemegang tempat uangnya), dan buku pribadinya tidak menerima perintah dari siapa pun, termasuk `claim_waiver` atas M. Yang perlu 06 tetapkan: autentikasi ulang, siapa diberi tahu dan kapan, pemberitahuan kepada yang keluar selama masa itu (dokumen ini menyarankan tidak ada), dan syarat perangkat mengirim perintah tertunda sebelum keluar dijalankan (07 §8.2 butir 2b). ADR-007 butir 10 perlu memuat bahwa penutupan dan penghapusan terjadi pada akhir masa tenggang.
4. **SP-49.** Rincian (`entry_detail`) entri yang dicatat mantan anggota tidak terbaca anggota yang bergabung sesudah orang itu keluar; inti moneter terbaca (07 §8.2 butir 3a). Aturan tampilannya milik 06 (larangan visibilitas `V-N`) dan 10.
5. **06 §6.3 "Pos rutin" (g)** diikuti 07 §2.6 tanpa perubahan. Satu tambahan dari sisi uang yang perlu dikukuhkan 06: pengarsipan jadwal oleh pemilik jatah **adalah** pengakhiran jatah oleh pemiliknya, jadi `assessment_until` berlaku bila jatahnya di atas ambang (07 §3.4). Satu celah untuk 06: usulan jatah yang masih **diusulkan** atas jadwal yang diarsipkan anggota lain (karena belum ada kesepakatan yang berlaku): ditarik atau gugur, seperti baris (f).
6. **PA-FIN-11** tidak mengubah 06 §13.4: pengingat membaca "belum dibayar" per kejadian; pembayaran lebih awal yang menunjuk kejadian berikutnya mematikan pengingat kejadian itu.

*Kepada `product-analyst` (dokumen 05, 12, 17).*

1. **FIN-1:** FR-FIN-042/AC-6 "(500000 − 300000)" menjadi "(300000 − 500000)" (CA-11 yang tersalin). FR-FIN-030/AC-11: ganti "[Perlu aturan PA-11]" dengan aturan 07 §3.15 "Urutan bawaan kejadian"; tutup PA-FIN-11 di tabel §13. Ganti nama pengenal celah menurut README.
2. **FIN-2:** FR-FIN-117: AC penghapusan kumulatif dengan angka CTH-25 varian SP-51 (700000; 300000 tidak bertanda, `PIU` 400000, `KB` 700000; 300000 bertanda, `PIU` 100000, `KB` 400000; 50000 bertanda, `PIU` 50000, `KB` 350000; varian disetujui: ketiga tidak bertanda, 300000 + 0 + 50000 = 350000). FR-FIN-135: AC tanda pemegang, hitungan kedua yang ditolak, dan hitungan sesudah jawaban (CTH-47: 3800000 → 3750000; ditolak; dikonfirmasi; 3750000 → 3740000); AC menyalakan dan mematikan tanda oleh pemegang (FR baru atau AC di FR tempat uang). AC pengarsipan jadwal penyisihan (07 §2.6).
3. **FAM:** masa tenggang keluar (07 §8.2 butir 2b), ekspor buku rumah tangga sesudah keluar (butir 1a), pembacaan rincian oleh anggota yang bergabung kemudian (butir 3a); angka AC dari CTH-38.
4. **12:** SCN-44: `exit_closure` pada hari akhir masa tenggang, angka −200000 + 60000 = −140000 tetap; skenario yang memuat hitungan bukan pemegang mengikuti syarat SP-50.
5. **17:** S3-9 (i) dan (ii) memuat batas SP-50 dan SP-51 (07 §10.4 butir 2 dan 3); S3-4 (c) masa tenggang (07 §10.4 butir 4); T-39, T-40 baru; T-32, T-33, T-37 berubah; W-26 menyebut kedua batas SP-50.

*Kepada lead.* Glosarium bagian 3: istilah *masa tenggang keluar* dan *mantan anggota* (sisi uangnya di bagian 7, *penutupan klaim saat keluar*). README tabel pengenal: INV-33, INV-34, T-39, T-40; tidak ada `CTH` baru (varian SP-51 berada di bawah CTH-25). Angka baru untuk skrip: CTH-25 varian SP-51, CTH-38 (dua baris baru), CTH-47 (enam baris baru).

**Kepada `security-privacy`.** Putaran perbaikan S3: SP-48 sampai SP-51 diadopsi seperti di atas. Dua hal untuk dinilai di dokumen 10: (a) selama masa tenggang yang keluar tetap terhitung anggota dewasa, sehingga tanda atas catatan pasangannya tetap lahir walau ia tidak dapat melihatnya; pilihan lain (tidak terhitung sejak akses berakhir) membuat tanda hilang tujuh hari dan tidak kembali bila keluar dibatalkan; (b) ekspor buku rumah tangga oleh yang keluar memuat rincian tulisan anggota yang tersisa sampai saat itu (§8.3). Putaran S3: selisih hitung oleh bukan pemegang (§3.10 butir 1, T-33) adalah jalan baru untuk mencatat "uang di tanganmu kurang" atas nama orang lain; penanda dan keterlihatan pelakunya adalah pagarnya, mohon dinilai di dokumen 10 bersama PRN-10. Buku pribadi yang ditutup tetapi belum dihapus (§8.2 butir 2a) memerlukan jangka retensi dan cara akses mantan anggota. Sesudah tiga penegasan 11 November 2026: pos rutin di bawah ambang kini dapat bertanda, dan mengakhiri kesepakatannya sendirian menghapus batas itu ke depan (§10.3 butir 10); jatah yang diakhiri tetap dinilai sampai akhir periode (T-29 ditutup). Sesudah checkpoint S2: T-29 (mengakhiri jatah sendiri kini juga mengurangi tanda atas diri sendiri; asas SP-24 perlu dinilai ulang terhadapnya); jatah berlaku sebesar 0 (§3.4); pembagian selisih dompet campuran dan ketertutupan jumlah isi dompet (§2.3). Jatah pribadi dan periode peralihan (§3.4, §6.1; §10.3 butir 7). Kebocoran B1–B5 (§5.3); risiko §10.3 butir 1 dan 2; aturan tulis buku pribadi dan batasnya atas koreksi (§2.2, §3.12); usulan §8 beserta pengecualiannya terhadap INV-09.

### 10.3 Yang menurut dokumen ini berisiko dalam keputusan pemilik produk

Dinyatakan terus terang, sesuai brief §17. Tidak satu pun dibatalkan dokumen ini; semuanya dijalankan sebagaimana diputuskan.

1. **Gaji ke dana pribadi kini terlihat, tetapi tiga hal tetap terbuka** (keputusan pemilik produk, 11 November 2026, bersama keputusan nomor 2 dan 3).
   - *Celah label.* INV-24 menjamin gaji yang **dicatat sebagai gaji** tidak dapat tersembunyi. Sumber diisi pencatat. Gaji yang dicatat sebagai "hadiah" atau "lainnya" langsung di buku pribadi, atau tidak dicatat, tetap tidak terlihat, dan model tidak dapat menutupnya tanpa membuka seluruh pemasukan dana pribadi. Yang tersisa bagi pasangan hanyalah bahwa gaji bulan itu tidak muncul (B5).
   - *Tidak ada batas atas bagian pribadi.* Penerima penghasilan boleh mengarahkan seluruh penghasilannya ke dana pribadinya (hanya ia sendiri: INV-27, keputusan pemilik produk, checkpoint S2). Pasangannya melihatnya dan dapat berkeberatan, tetapi tidak dapat membatalkannya: entri itu entri dua buku (§3.12). Jatah pribadi tidak mengubahnya: jatah menentukan kapan entri ditandai, bukan berapa yang dapat diarahkan (§3.4). Keterlihatan dan tanda adalah satu-satunya pagar.
   - *Simetris dalam aturan; setara dalam praktik hanya sejauh penghasilannya setara.* Kedua pasangan di rumah tangga pilot berpenghasilan (pernyataan pemilik produk, 11 November 2026). Aturannya per penerima penghasilan: masing-masing mencatat penghasilannya sendiri dan dapat mengarahkan bagian dari penghasilannya sendiri, tidak dari penghasilan pasangannya (INV-27; keputusan pemilik produk, checkpoint S2). Yang tersisa sebagai risiko: besarnya yang dapat diarahkan mengikuti besarnya penghasilan masing-masing. Pada rumah tangga yang hanya satu anggotanya berpenghasilan (mungkin ada di tahap publik), hak mengarahkan dalam kenyataan dimiliki satu orang dan anggota yang lain hanya memperoleh dana pribadi lewat penyisihan.
2. **Laporan "siapa membayar apa" mengukur siapa yang memegang uang, bukan siapa yang berbelanja.** Pada CTH-37 suami tercatat membayar 6031000 dan istri 930000 hanya karena gaji tiba di rekeningnya dan ia yang membayar tagihan dan kiriman. Kriteria sukses nomor 3 meminta laporan ini; tanpa keterangan ia mudah dibaca sebagai penilaian. Hal yang sama berlaku untuk "kontribusi dari dana pribadi per anggota", yang dapat menjadi papan skor.
3. **Selisih kas melekat pada pemegang** (keputusan nomor 7). "Siapa yang lupa diselesaikan di luar aplikasi", tetapi selisih tercatat per tempat uang dan tiap tempat uang punya pemegang, sehingga aplikasi tetap memperlihatkan di tangan siapa uang yang tidak tercatat itu. Model tidak dapat menghindarinya; dokumen 08 dapat memilih kata yang tidak menuduh.
4. **Ambang permintaan adalah kesepakatan, bukan kunci** (keputusan nomor 5 bersama A-20). Aplikasi tidak dapat mencegah pemakaian di atas ambang; ia hanya membuatnya terlihat. Kriteria sukses nomor 4 bergantung pada kebiasaan, bukan pada aturan sistem.
5. **Satu tempat uang, satu pemilik** adalah usulan dokumen ini, bukan keputusan pemilik, tetapi ia akibat langsung dari keputusan nomor 3. Menurut pemilik produk (checkpoint S2, S2-6) uang pribadi dan uang bersama di rumah tangga pilot dapat bercampur ("bisa terpisah bisa tercampur bisa di bank bisa tunai"), jadi ini kasus biasa, bukan kasus tepi. Aturannya dipertahankan dengan perlakuan §2.3: satu dompet, dua tempat uang, dihitung satu kali. Harganya: (a) pemegang harus mengingat berapa bagian pribadinya, dan bagian pribadi di dompet campuran tidak pernah dapat dicocokkan sendiri, hanya bersama bagian bersamanya; (b) pemegang memilih pembagian selisih setiap kali, tanpa bawaan (diterima pemilik produk atas saran lead, 12 November 2026; T-03): bila ia memilih "bersama" atas uang yang sebenarnya pribadi, kekurangannya tampil sebagai "selisih hitung" uang bersama atas namanya dan terlihat pasangan (butir 3; CTH-40); bila ia selalu memilih "pribadi", selisih dompet itu tidak pernah terlihat pasangan; (c) pembagian selisih adalah pernyataan pemegang yang tidak dapat diperiksa; (d) siapa pun yang melihat dompet atau saldo rekening campuran itu sendiri dapat menghitung bagian pribadinya dengan mengurangkan bagian bersama yang terlihat (§5.3, B6; SP-40 (a)); (e) bahwa server tidak menerima jumlah seluruh isi dompet hanya menjaga apa yang terbaca pasangan, bukan apa yang terbaca pengelola server (SP-40 (d)).
6. **Kontribusi yang terlihat membocorkan dana pribadi** (B2, B4). Ini akibat yang dipilih keputusan nomor 3 dan diterima; anggota perlu diberi tahu sebelum mencatat.
7. **Jatah pribadi adalah batas yang pelampauannya selalu bertanda; ia tetap bukan kunci.** Keputusan pemilik produk: bentuk jatah (11 November 2026: satu jumlah per orang per periode, dijumlahkan sepanjang periode, dapat diatur, disepakati berdua, kenaikan diajukan dengan alasan), dan di checkpoint S2 bahwa jatah itu batas (S2-1 butir (iii) pilihan b): begitu jumlah periode itu melebihi jatah, entrinya ditandai berapa pun kelebihannya. Bahwa di dalam jatah tidak ada tanda dan tidak ada pemberitahuan adalah default yang ia terima lewat "lanjut" (dokumen 17 D-48), bukan pernyataannya sendiri. Mengambil di atas jatah tidak dihalangi; yang ada hanya tanda dan keterlihatan. Yang tersisa, dinyatakan terus terang:
   - **Jalan yang gugur.** Mengambil di atas jatah dalam potongan yang masing-masing tidak melebihi ambang, yang pada rancangan sebelumnya tidak pernah ditandai (W-13 cara (1)), kini tertutup bagi anggota yang punya jatah: setiap potongan sesudah jatah terlewati ditandai.
   - **Jalan yang tetap terbuka dan hanya dijaga keterlihatan.** (2) Karena periode dihitung menurut tanggal kejadian, penyisihan dapat diberi tanggal periode lampau yang jatahnya belum terpakai; yang melawannya hanya tanggal pencatatan yang tampil pada entri terlambat (§3.11). (3) Mengganti hari mulai periode membuka satu periode pendek dengan jatah dan batas anggaran penuh; karena itu penggantiannya kesepakatan rumah tangga dan terlihat (§6.1, T-25), tetapi dua orang yang sepakat tetap dapat melakukannya berulang kali.
   - **Jalan baru yang lahir dari keputusan ini: tidak punya jatah.** Tanpa jatah yang berlaku, tiap penyisihan dibandingkan dengan ambang sendiri-sendiri, sehingga potongan yang masing-masing tidak melebihi ambang tidak pernah ditandai, tanpa batas jumlah. Anggota yang punya jatah karenanya lebih sering ditandai atas jumlah kecil daripada anggota yang tidak punya. Pemilik jatah dapat mengakhiri jatahnya sendiri kapan saja (D-52); pengakhiran itu terlihat dan diberitahukan. **Jalan keluar seketika ditutup** (T-29 pilihan 1; rumusan lead yang dikukuhkan pemilik produk, 11 November 2026): jatah yang diakhiri tetap menjadi dasar penilaian sampai akhir periode anggaran yang memuat pengakhirannya, sehingga mengakhiri jatah begitu jatah habis tidak menghapus tanda atas pengambilan berikutnya di periode itu. Itu hanya berlaku atas jatah yang di atas ambang saat diakhiri; jatah yang tidak di atas ambang, termasuk 0, berakhir seketika (**[Usulan]** `security-privacy`, SP-33), karena jatah seperti itu tidak pernah memberi kelonggaran yang dapat ditinggalkan. **Yang tidak ditutup, dinyatakan:** (a) sejak periode berikutnya anggota itu tanpa jatah, dan potongan yang masing-masing tidak melebihi ambang tidak ditandai lagi; anggota yang mengakhiri jatahnya pada hari terakhir periode bebas sejak esoknya. Itu sama dengan tidak pernah menyepakati jatah, yang memang boleh (W-14). (b) Batas masa itu adalah waktu simpan di server (**[Usulan]** `architect`): pengambilan yang dicatat sesudah masa itu dengan tanggal kejadian di dalamnya dinilai tanpa jatah; yang melawannya hanya tanggal pencatatan yang tampil pada entri terlambat (§3.11). Menunda pencatatan sampai periode berikutnya karenanya lolos dari jatah lama, tetapi uangnya tetap tercatat dan terlihat.
   - **Urutan tiba menentukan jumlah tanda, dan tanda tidak dihitung ulang.** Contoh (angka **[Asumsi]**): jatah 1000000, tiga penyisihan 400000, 400000, dan 900000 dalam periode yang sama. Tiba dalam urutan itu: 0 + 400000 = 400000 dan 400000 + 400000 = 800000 tidak melebihi jatah; 800000 + 900000 = 1700000 > 1000000; hanya yang ketiga ditandai, **satu tanda**. Bila 900000 tiba lebih dulu: 0 + 900000 = 900000 tidak melebihi jatah; 900000 + 400000 = 1300000 > 1000000; 1300000 + 400000 = 1700000 > 1000000; kedua penyisihan sesudahnya ditandai, **dua tanda**. Jumlah yang diambil sama, 1700000, dan kelebihannya sama, 700000. Jumlah tanda karenanya bukan ukuran besarnya kelebihan, dan dokumen 08 tidak boleh menyajikannya demikian. Pada rancangan sebelumnya contoh yang sama memberi satu tanda lawan tidak ada tanda.
   - **Tanda atas jumlah yang sangat kecil.** Kelebihan satu rupiah pun menandai entrinya. Tanda itu tampil pada entri dan di ringkasan saja, tidak pernah sebagai pemberitahuan langsung (diterima pemilik produk atas saran lead, 12 November 2026; aturannya milik dokumen 06). Ini akibat pilihan (iii) b yang dinyatakan pemilik produk; harga yang disebut `product-analyst` dan `security-privacy` (SP-30) tetap berlaku: lebih banyak tanda atas nama orang yang jatahnya lebih kecil atau yang lebih sering perlu mengambil.
   - Penyisihan dinilai bersama penyisihan lain di periode yang sama, tidak lagi sendiri-sendiri; itu perubahan atas sifat tanda di dokumen 06 (INV-A17).
8. **Makin sering penghasilan datang, makin besar celah label.** Dengan uang masuk pekerjaan lepas yang "bisa sampe 5x" (kata-kata pemilik produk; "sebulan" adalah bacaan lead, §3.1), setiap uang masuk adalah kesempatan memberinya sumber "lainnya" dan mencatatnya di buku pribadi saja (butir 1, T-27). Model tidak dapat menutupnya; pasangan hanya melihat bahwa penghasilan itu tidak muncul. Bahwa pekerjaan lepas dihitung sebagai penghasilan kini diterima pemilik produk atas saran lead (checkpoint S2, S2-1 butir (iv)); celah ini tetap ada bagi pencatat yang memilih sumber "lainnya".
9. **Dua kalimat pemilik produk di checkpoint S2 yang dapat dibaca lebih dari satu cara: keduanya kini dijawab.** Lead menanyakan keduanya; jawabannya adalah rumusan lead yang dikukuhkan pemilik produk dengan "bener" (11 November 2026), bukan kata-katanya sendiri.
   - *"harus ajukan kenaikan bukan keringanan"*. **Dikukuhkan: bacaan pertama.** "Bukan keringanan" berarti tidak ada kelonggaran otomatis di atas jatah (pilihan (a) dan bentuk tengah ditolak). Tambahan untuk satu bulan tetap mungkin lewat permintaan biasa yang disetujui (D-53); penyisihan yang tertaut ke otorisasinya tidak terhitung terhadap jatah dan, sampai batas otorisasi itu, tidak ditandai (D-48). Kebutuhan yang menetap adalah usulan kenaikan jatah. Bacaan kedua ("tambahan satu kali pun tidak dikehendaki") **tidak berlaku**; D-48 dan D-53 tidak berubah.
   - *"untuk membuat tagihan rutin ya disepakati dulu"*. **Dikukuhkan: bacaan kedua, dan lebih jauh daripada yang dicatat revisi sebelumnya.** Setiap tagihan berulang, berapa pun jumlahnya, disepakati saat dibuat; jumlah yang disepakati adalah batasnya; pembayaran di atasnya ditandai tanpa membandingkan kelebihannya dengan ambang (§3.15). Bacaan pertama, yang dipakai revisi sebelumnya (hanya tagihan di atas ambang), gugur. Pernyataan revisi sebelumnya bahwa pada bacaan ini "aturan uang dan angka contoh tidak berubah" **keliru separuh**: saldo memang tidak berubah, tetapi aturan tanda berubah dan satu entri CTH-41 kini bertanda (§11).
10. **Pos rutin sebagai batas: harganya.** Dijalankan sebagaimana dikukuhkan; yang berikut dinyatakan terus terang.
   - **Pembayaran di bawah ambang kini dapat bertanda.** Sebelum revisi ini tidak ada pembayaran di bawah ambang yang pernah ditandai. Kini tagihan air yang disepakati 250000 dan dibayar 310000 bertanda walau ambangnya 500000 (CTH-41), dan selisih seribu rupiah pun menandai. Tanda jatuh pada orang yang membayar tagihan, yang di rumah tangga pilot adalah penerima penghasilan yang membayar tagihan lebih dulu; jumlah tagihannya ditentukan pihak luar, bukan olehnya.
   - **Yang melindungi hanya jumlah yang disepakati.** Menyepakati perkiraan atau rata-rata membuat tagihan yang berubah-ubah bertanda hampir setiap kali naik; menyepakati jumlah tertinggi yang wajar menghindarinya. Itu syarat layar bagi dokumen 05 dan 08, bukan aturan uang.
   - **Pos yang disepakati lebih buruk kedudukannya daripada yang tidak, bagi tagihan kecil.** Tagihan di bawah ambang yang kesepakatannya belum berlaku atau sudah diakhiri tidak pernah bertanda; yang kesepakatannya berlaku bertanda begitu lewat `R`. Salah satu pihak dapat mengakhiri kesepakatan sendirian (dokumen 06 K06) dan dengan itu menghapus batasnya ke depan; pengakhirannya terlihat dan diberitahukan. Tidak menautkan pembayaran ke tagihannya berakibat sama. Dokumen ini tidak menutupnya. Keduanya usulan `architect` (dokumen 06 P-20, P-26) yang diterima lewat "lanjut" di checkpoint S3 (12 November 2026): mengakhiri sendirian S3-5 (i), aturan ambang biasa sesudahnya S3-5 (ii); risiko ini karenanya tetap.
   - **Jumlah tanda bukan ukuran.** Seperti pada jatah (butir 7), banyaknya tanda atas tagihan tidak mengukur besarnya kelebihan; dokumen 08 tidak boleh menyajikannya sebagai hitungan per orang atau per pos.
   - **Tanda tanpa pemberitahuan langsung.** Tanda dasar `personal_share` dan `routine_amount` tampil pada entri dan di ringkasan saja, tidak pernah sebagai pemberitahuan langsung (diterima pemilik produk atas saran lead, 12 November 2026; temuan `security-privacy` SP-37, SP-43; aturan pemberitahuannya milik dokumen 06). Ini meringankan dua butir pertama di atas tanpa mengubah aturan tanda: entri yang sama tetap bertanda. Pemberitahuan langsung (N-10) hanya ada untuk tanda dasar `threshold`, paling banyak satu per penerima per hari (D-35); kedua dasar lain hanya tampil pada entri dan di ringkasan.
   - Dari sisi uang tidak ada harga: tidak ada saldo, anggaran, atau alokasi yang bergantung pada kesepakatan (INV-12).
11. **Menyepakati jatah dapat merugikan pemiliknya: ketimpangan yang dinyatakan** (temuan `security-privacy` SP-32; dibacakan lead kepada pemilik produk, 12 November 2026). Anggota tanpa jatah dapat mengambil berkali-kali, tiap kali tidak di atas ambang, tanpa satu tanda pun dan tanpa batas jumlah; anggota berjatah ditandai begitu jatahnya terlewati. Contoh (angka **[Asumsi]**): ambang 500000, lima pengambilan 400000 dalam satu periode, jumlah 2000000.

    | Keadaan | Jumlah sesudah tiap pengambilan | Yang ditandai | Jumlah tanda |
    |---|---|---|---|
    | Tanpa jatah | — (tiap 400000 dibandingkan dengan ambang; 400000 tidak di atas 500000) | Tidak ada | 0 |
    | Jatah 1000000 | 400000, 800000, 1200000, 1600000, 2000000 | Ketiga, keempat, kelima (1200000, 1600000, 2000000 > 1000000) | 3 |
    | Jatah 300000 | 400000, 800000, 1200000, 1600000, 2000000 | Kelimanya (400000 > 300000 sejak yang pertama) | 5 |

    Uang yang berpindah sama pada ketiga baris. Akibatnya: (a) satu orang boleh berjatah sementara pasangannya tidak, dan yang berjatah lebih sering ditandai; (b) jatah yang tidak di atas ambang hanya pernah menambah tanda, karena itu jatah yang tidak di atas ambang, termasuk 0, hanya dapat diusulkan pemiliknya sendiri dan dapat ia lepas seketika (**[Usulan]** `security-privacy`, SP-32 sampai SP-34; disampaikan kepada pemilik produk dan tidak ia bantah; §3.4); (c) temuan SP-30, bahwa jatah hanya pernah mengurangi tanda, tidak lagi berlaku. Dokumen ini tidak menutup (a): menutupnya berarti memperlakukan anggota tanpa jatah seperti berjatah 0 (T-29 pilihan 3, tidak diambil) atau menjumlahkan pengambilan tanpa jatah terhadap ambang per periode, yang keduanya mengubah D-48. Tanda dasar jatah tidak pernah menjadi pemberitahuan langsung (butir 10), yang mengurangi beratnya tetapi tidak menghapus ketimpangannya. Syarat layar bagi dokumen 05 dan 08: sebelum jatah disetujui, layar menyebut apa yang ditandai dengan dan tanpa jatah itu.

### 10.4 Putaran S3: yang ditunda dan yang perlu pemilik produk

**Ditunda dengan sengaja** (permintaan dokumen 11 §10.1 dan dokumen 04). Tidak satu pun menjawab masalah brief §1.1 bagi pilot sekarang; aturan Tier 1 tidak bergantung padanya.

| Butir | Asal | Alasan ditunda |
|---|---|---|
| Rumus "perkiraan uang bebas" (tagihan mana dikurangkan, perkiraan atau jumlah sepakat, perlakuan *terikat*) | Dokumen 11 IDEA-04 tingkat 2, §10.1 butir 3 | Tingkat 1 IDEA-04 hanya menampilkan *belum dialokasikan* dan daftar kejadian yang belum dibayar (§3.15), tanpa rumus baru; tingkat 2 belum diminta |
| Laju anggaran ("habis sekitar tanggal *n*") | Dokumen 11 IDEA-05, §10.1 butir 4 | Konsepnya belum diterima; bila diterima, rumusnya turunan `terpakai` dan `terikat` §6.1 dan ditetapkan saat itu |
| Arisan | Dokumen 11 IDEA-10, §10.1 butir 6 | Tier 3; tidak diketahui apakah pilot ikut arisan; sampai itu dicatat dengan kategori buatan sendiri |
| Pola jatuh tempo tahunan | Dokumen 11 IDEA-12, §10.1 butir 7 | `due_rule` Tier 1 hanya bulanan (§3.15); tidak ada kebutuhan pembayaran tahunan pilot yang sudah diketahui, dan pembayaran tahunan apa pun (misalnya biaya sekolah tahunan, kini mungkin karena rumah tangga pilot punya anak berusia 12 tahun ke atas, jawaban pemilik produk sendiri "Ada anak umur 12+"; A-31 gugur) tetap dapat dicatat sebagai belanja biasa, disiapkan dengan alokasi, dan bila jumlahnya di atas ambang diajukan sebagai permintaan, tanpa pengingat jatuh tempo. Menambahkan `yearly(month, day)` kelak tidak mengubah kejadian bulanan |
| Tingkat tabungan | Permintaan dokumen 04 (_progress.md, fase A) | `GOL` lanjutan Tier 3; definisinya bergantung pada pembagi (pemasukan dengan atau tanpa bagian pribadi) yang belum perlu diputuskan |

**Pertanyaan untuk pemilik produk di checkpoint S3.** Ditanyakan lewat dokumen 17, bukan dari sini; nomor checkpoint di bawah adalah nomor di 17. Semua bawaan lahir sebagai usulan putaran S3 atau putaran perbaikan S3. **Keempatnya dijawab di checkpoint S3 (12 November 2026)**: jawaban pemilik produk "lanjut", yang menurut pernyataan lead berarti semua bawaan diterima ([_progress.md](_progress.md), "Checkpoint S3 dijawab"). Hasilnya di tiap butir; kedudukannya *diterima lewat "lanjut"*, bukan pernyataan pemilik produk sendiri.

1. **THR dan bonus** (checkpoint **S3-3**, Q-145) dicatat sebagai gaji, jadi selalu terlihat pasangan beserta bagian yang diarahkan ke uang pribadi? Bawaan: ya (T-34). **Dijawab: A, diterima lewat "lanjut"**; T-34 ditutup (§3.1). **Dibuka ulang sebagian sesudahnya**: pemilik produk bertanya "kenapa thr dicatat sebagai gaji? kan pemasukan kita yang atur darimana nya?" dan menulis "ada kategori default bisa edit tambah atau yang lainnya" (kata-katanya sendiri; [_progress.md](_progress.md), "Pertanyaan pemilik produk atas S3-3"; bahwa "kategori" berarti sumber uang masuk adalah bacaan lead **[Usulan]**). THR tidak lagi dipaksa berlabel gaji: ia sumber bawaan "THR" di kelompok `work`, dengan akibat uang yang sama (**[Usulan]** lead; §3.1 "Sumber uang masuk"; T-34 dibuka ulang, T-43). Yang belum ia jawab, dan ditanyakan di checkpoint S4: apakah sumber buatan sendiri boleh dicatat tanpa terlihat pasangan; bawaan lead: boleh, hanya bila sumber itu di kelompok `gift_other`, dan perubahan kelompok lewat kesepakatan berdua. Pertanyaan sampingan "apakah kalian menerima THR?" tidak dijawab di checkpoint; sesudahnya pemilik produk menulis "thr bisa menerima bisa ngasih" dan "bro pikirkan kalau semua orang pakai app ini" ([_progress.md](_progress.md), 12 November 2026). THR yang diberikan kini diatur di §3.2 (**[Usulan]** `finance-domain`; T-42; CTH-51).
2. **Menghitung uang yang dipegang pasangan** (checkpoint **S3-9 (i)**, Q-148). Bolehkah satu orang mencatat hasil hitung atas uang bersama yang dipegang pasangannya, dengan catatan "belum dikonfirmasi" sampai pemegangnya menjawab? Bawaan: boleh, dengan penanda, dan dengan dua batas tambahan dari `security-privacy` (SP-50): **hanya atas rekening bank atau dompet digital yang pemegangnya tandai "pasanganku dapat melihat saldo rekening ini"** (bawaan mati, dapat dicabut kapan saja; uang tunai tidak pernah), dan **paling banyak satu hitungan yang belum dijawab per rekening**: hitungan berikutnya ditolak sampai pemegang mengonfirmasi, berkeberatan, atau menghitung sendiri (T-33, §3.10 butir 1, CTH-47). **Dijawab: A, diterima lewat "lanjut"**, dengan kedua batas SP-50; T-33 ditutup (INV-33).
3. **Mengikhlaskan pinjaman keluarga** (checkpoint **S3-9 (ii)**, Q-149). Bila jumlahnya di atas ambang, apakah ia ditandai seperti kiriman sebesar itu? Bawaan: ya, **dihitung kumulatif per pinjaman** (SP-51): mengikhlaskan 600000 dalam dua kali 300000 menandai yang kedua dan setiap penghapusan sesudahnya atas pinjaman itu; bagian yang sudah disetujui lewat permintaan tidak dihitung (T-32, §3.7, CTH-25 varian SP-51). **Dijawab: A, diterima lewat "lanjut"**, kumulatif per piutang; T-32 ditutup (INV-34).
4. **Masa tenggang keluar** (checkpoint **S3-4 (c)**, usulan `security-privacy` SP-48; pertanyaan dan pilihannya disusun pemilik dokumen 17). Sisi uangnya bila diterima: aksesnya berakhir seketika; klaim antara uang pribadinya dan uang bersama tetap tampil sampai masa tenggang (7 hari **[Asumsi]**) lewat, lalu ditutup dengan angka yang sama; tidak ada yang dapat menulis buku pribadinya selama itu; ia tetap dapat mengekspor buku rumah tangga sampai saat ia keluar (T-39, §8.2 butir 1a, 2b, CTH-38). Bawaan: terima. **Dijawab: S3-4 pilihan A, termasuk (c), diterima lewat "lanjut"**; T-39 ditutup (INV-28). Panjang 7 hari tetap **[Asumsi]** (A-39), diuji di pilot.

**Dua konsep dokumen 11 yang sisi uangnya di sini, juga dijawab di checkpoint S3** (S3-2 pilihan A, diterima lewat "lanjut"): IDEA-01 Tier 1, bentuk uangnya §3.10 "Menghitung lalu merinci seluruh selisih sekaligus" (CTH-48); IDEA-03 Tier 1, **diadopsi** di §3.15 "Tagihan berulang tanpa `due_rule`" dengan perubahan minimal yang sudah dinyatakan, tanpa angka contoh yang berubah.

---

## 11. Riwayat perubahan

Hanya revisi sesudah checkpoint S2 yang didaftar di sini; revisi sebelumnya tercatat di baris Status.

### Revisi pasca-checkpoint S2 (11 November 2026)

Dasar: [_progress.md](_progress.md) bagian "Jawaban checkpoint S2". Ditinjau `critic` dan angkanya dihitung ulang lead dengan skrip (12 November 2026).

| Bagian | Perubahan |
|---|---|
| Kepala dokumen, §1.4 | Tabel label sesudah checkpoint S2; baris kata layar untuk alokasi |
| §2.2, §2.5, §3.1, INV-27, §9, T-16, T-17 | INV-27 dari **[Usulan]** menjadi keputusan pemilik produk (S2-1 butir (i)) |
| §2.3, §3.10, §10.3 butir 5, T-03 | Dompet campuran menjadi kasus biasa (S2-6); cara menghitung dan membagi selisih; biaya dinyatakan |
| §2.4, §3.1, T-27 | Pekerjaan lepas adalah penghasilan (diterima atas saran lead); model tidak mengandaikan irama penghasilan |
| §3.1, §5.3 B5, CTH-36 | Hadiah untuk satu orang tidak terlihat: keputusan pemilik produk (S2-1 butir (ii)) |
| §3.4, §3.8 (tabel *jumlah pemakaian*), INV-30, §9, T-13 | Jatah adalah batas (S2-1 butir (iii) pilihan b): rumus `ditandai(e)`; pilihan (a) dan bentuk tengah menjadi pilihan yang ditolak; jatah berlaku sebesar 0; dasar D-52 berubah |
| §3.2, §3.15, §7.1, CTH-06, CTH-39 | Fakta rumah tangga pilot (S2-4); kiriman orang tua ditandai sebagai ilustrasi; alokasi untuk tagihan yang menyusul; jumlah per kali sebagai batas atas |
| §8, INV-28, CTH-38 | Dari usulan menjadi keputusan sementara pemilik produk (S2-8) |
| §10.1 | T-28, T-29 baru |
| §10.2 | Butir 10 (`architect`), butir 13 (lead), permintaan kepada `product-analyst` dan `security-privacy` |
| §10.3 | Butir 1, 5, 7, 8 ditulis ulang atau dilengkapi; butir 9 baru |

**Angka dan tanda yang berubah atau baru** (untuk hitung ulang lead):

| Tempat | Sebelum | Sesudah |
|---|---|---|
| Semua tabel saldo CTH-01 sampai CTH-39; CTH-37 | — | **Tidak ada angka yang berubah** |
| CTH-03b, teks tanda | "Belum diputuskan" | Bergantung jatah: jatah ≥ 1000000 tidak ditandai (0 + 1000000 ≤ jatah); jatah 600000 ditandai (1000000 > 600000); tanpa jatah ditandai (1000000 > 500000) |
| CTH-06, CTH-33, CTH-34 | Tanda bergantung ambang dan kesepakatan rutin atau otorisasi | Tidak berubah: tidak satu pun penyisihan. CTH-06 ditandai ilustrasi; bila kiriman tidak rutin, 3000000 diajukan dulu atau bertanda |
| CTH-35, teks tanda | Tidak disebut | Tidak ditandai: 0 + 100000 tidak melebihi jatah 100000; tanpa jatah 100000 tidak di atas 500000 |
| CTH-39, tabel jatah | Tujuh kolom | Kolom "Ditandai" ditambah: termin pertama Tidak (600000 ≤ 1000000); termin kedua Ya (1600000 > 1000000). Hasil tanda sama dengan sebelumnya; dasarnya berubah |
| CTH-39, tabel varian (baru) | — | Bagian 400000: dalam 400000, luar 0, sesudah 1000000, tidak ditandai. Bagian 500000: dalam 400000, luar 100000, sesudah 1100000, ditandai (pada pilihan (a) yang ditolak: tidak). Bagian 1000000: dalam 400000, luar 600000, sesudah 1600000, ditandai |
| §10.3 butir 7, urutan tiba (semula hanya di dokumen 17 W-13) | Jatah 1000000, ambang 500000; 400000, 400000, 900000: satu tanda; 900000 lebih dulu: tidak ada tanda | Jatah 1000000; 400000, 400000, 900000: satu tanda (800000 + 900000 = 1700000); 900000 lebih dulu: dua tanda (1300000 dan 1700000) |
| CTH-40 (baru) | — | Lima tabel; `C` = 890000; selisih −40000; hasil pilihan 1: `DMP-B` 640000, `UANG` 2640000; pilihan 2: `DMP-P` 210000, `UANG` 2680000; pilihan 3: `DMP-B` 655000, `DMP-P` 235000, `UANG` 2655000 |
| CTH-41 (baru) | — | Tiga tabel; `UANG` 650000 menjadi 6560000; `BNK-S` akhir 810000; `BNK-I` 5750000; `PRI-S` 650000; belum dialokasikan 6750000, 1250000, 750000, 750000, 870000, 810000; tidak ada entri yang ditandai |

Pengenal baru: CTH-40, CTH-41, T-28, T-29. Tidak ada INV baru; INV-27, INV-28, dan INV-30 berubah isi atau labelnya, nomornya tetap.

### Revisi kedua pasca-checkpoint S2 (11 November 2026)

Dasar: [_progress.md](_progress.md) paragraf "Tiga penegasan sesudah revisi dokumen 07" (rumusan lead yang dikukuhkan pemilik produk, 11 November 2026) dan dokumen 06 revisi ketujuh, §15.4 butir (t) dan (u). Ditinjau `critic`; angka dan tanda dihitung ulang lead dengan skrip, semuanya cocok (12 November 2026).

| Bagian | Perubahan |
|---|---|
| Kepala dokumen | Label keempat: rumusan lead yang dikukuhkan pemilik produk |
| §2.3, §3.10, CTH-40 | Dikukuhkan: server tidak pernah menerima `C`; dua perintah rekonsiliasi yang tidak terkait; baris "Yang diterima server" baru. CTH-40 pilihan 2: rekonsiliasi `DMP-B` yang cocok tetap tampil bagi pasangan |
| §2.6 | Setiap tagihan berulang dibuat bersama usulan kesepakatannya |
| §3.4, INV-30, T-29 | T-29 pilihan 1: jatah yang diakhiri pemiliknya tetap dasar penilaian sampai akhir periode; `ada_jatah` diperluas; batas waktu simpan server; jatah 0 dikukuhkan; "bukan keringanan" tidak lagi bacaan |
| §3.8 | Tabel *jumlah pemakaian* baris pertama; blok baru "Tanda atas entri" dengan tiga dasar (`threshold`, `personal_share`, `routine_amount`); "Penautan" butir 1 (entri bertanggal sebelum persetujuan; syarat diperiksa saat tautan dibuat) dan butir 3 (tautan susulan ke tagihan berulang hanya oleh pencatat) |
| INV-26 | Syarat tautan berlaku saat tautan dibuat; otorisasi `open` tanpa syarat tanggal; tautan susulan ke tagihan berulang hanya oleh pencatat; uji ditambah |
| §3.15 | Setiap tagihan berulang disepakati saat dibuat; jumlah per kali adalah batas; delapan aturan dengan kedudukannya; baris ketiga tabel kebutuhan |
| CTH-05, CTH-06, CTH-39, CTH-41 | Penilaian tanda terhadap jumlah yang disepakati (**[Asumsi]**) |
| §9 | Tiga baris baru: tiga penegasan; dokumen 06 §15.4 butir (t); butir (u) |
| §10.2 | Butir 11 (`architect`), butir 14 (lead), tambahan bagi `product-analyst` dan `security-privacy` |
| §10.3 | Butir 7 ("tidak punya jatah") dan butir 9 ditulis ulang; butir 10 baru |

**Angka dan tanda yang berubah atau baru** (untuk hitung ulang lead):

| Tempat | Sebelum | Sesudah |
|---|---|---|
| Semua tabel saldo, alokasi, dan ringkasan CTH-01 sampai CTH-41; CTH-37 | — | **Tidak ada angka yang berubah** |
| CTH-05, tabel penilaian (baru) | Tanda tidak disebut | `R` **[Asumsi]**: listrik 500000, wifi 350000, langganan kerja 200000. `P`: 450000, 350000, 200000. Ketiganya `P ≤ R`: tidak ditandai. Varian: listrik disepakati 400000, 450000 > 400000: ditandai (dasar `routine_amount`), walau 450000 tidak di atas ambang 500000. Varian kesepakatan belum berlaku: 450000 tidak di atas 500000, tidak ditandai |
| CTH-06, teks | Tidak ditandai: tercakup kesepakatan rutin 3000000 | Tidak ditandai, dasar berubah: 3000000 tidak melebihi `R` 3000000. Varian baru: 3100000 ditandai |
| CTH-39, kiriman ibu 3000000 | Tidak ditandai (tercakup kesepakatan) | Tidak ditandai: 3000000 tidak melebihi `R` 3000000; tanpa kesepakatan berlaku: ditandai (3000000 > 500000). Hasil sama, dasar berubah |
| CTH-40, tabel pilihan 2, kolom "terlihat pasangan" untuk `DMP-B` | "Tidak ada yang tampil" | Ya: dihitung dan cocok, terhitung 680000. Angka tidak berubah |
| CTH-41, tabel tagihan | Kesepakatan: internet, iuran kesehatan, air "tidak perlu"; listrik paling banyak 900000 | `R` **[Asumsi]**: internet 350000, iuran kesehatan 150000, listrik 900000, air 250000. `P`: 350000, 150000, 780000, 310000 |
| CTH-41, internet 350000 | Tidak ditandai (di bawah ambang) | Tidak ditandai: 350000 tidak melebihi 350000 |
| CTH-41, iuran kesehatan 150000 | Tidak ditandai (di bawah ambang) | Tidak ditandai: 150000 tidak melebihi 150000 |
| CTH-41, listrik 780000 | Tidak ditandai (tercakup 900000) | Tidak ditandai: 780000 tidak melebihi 900000. Hasil sama |
| CTH-41, varian listrik 950000 | Tidak ditandai (kelebihan 50000 tidak di atas ambang 500000) | **Ditandai**, dasar `routine_amount`: 950000 > 900000 |
| CTH-41, varian listrik 780000 tanpa kesepakatan berlaku | Ditandai | Ditandai, dasar `threshold`: 780000 > 500000. Hasil sama |
| CTH-41, air 310000 | Tidak ditandai (310000 tidak di atas ambang 500000) | **Ditandai**, dasar `routine_amount`: 310000 > 250000. Varian: disepakati 350000, tidak ditandai; kesepakatan belum berlaku, tidak ditandai |
| CTH-41, penyisihan 500000 | Tidak ditandai (0 + 500000 tidak melebihi jatah 500000) | Tidak berubah |
| CTH-41, jumlah entri bertanda | 0 | **1** dari tujuh entri uang (pembayaran air) |
| INV-30, uji baru | — | Jatah 500000 habis terpakai lalu diakhiri: penyisihan 100000 yang disimpan dalam periode yang sama bertanda (500000 + 100000 > 500000); 100000 yang disimpan pada hari pertama periode berikutnya, ambang 500000, tidak bertanda |

Pengenal baru: tidak ada. T-29 ditutup. INV-26 dan INV-30 berubah isi, nomornya tetap. Butir §10.3 nomor 10 baru.

### Putaran perbaikan (12 November 2026)

Dasar: tinjauan `critic` atas revisi pasca-checkpoint S2 (belum siap; B2, S3, S4, S7, M4, M5 menyangkut dokumen ini), temuan `security-privacy` SP-32 sampai SP-40, dan empat jawaban pemilik produk ([_progress.md](_progress.md), "Empat pertanyaan dari SP-32 sampai SP-43"). Ditinjau ulang `critic` tanpa temuan menghalangi (temuannya diperbaiki, tanpa tinjauan ketiga); angka dan tanda dihitung ulang lead dengan skrip sesudah perbaikan terakhir, semuanya cocok (12 November 2026).

| Temuan | Bagian | Perubahan |
|---|---|---|
| B2 | Tabel label, §9 baris S2-4 | Bagian ambang S2-4: "tidak disebut pemilik produk; diterapkan lead", bukan "diterima atas saran lead" |
| S3 | §3.15 kalimat pembuka | Dipisah: "disepakati dulu, siapa pun boleh membayar" keputusannya sendiri; "setiap, berapa pun jumlahnya, jumlah = batas" rumusan lead yang ia kukuhkan |
| S4 | §3.4, CTH-39 | Bacaan "split jatah + uang pribadi" diberi label **[Usulan]** lead, belum dikukuhkan |
| S7 | §3.4 (tabel, `ditandai(e)`), §3.8 "Tanda atas entri", INV-30, CTH-21 sampai CTH-24, CTH-33, CTH-34 | Syarat "tidak tertaut ke otorisasi" dicabut; `N = P − dalam_batas`; bagian di atas batas otorisasi dibandingkan dengan ambang |
| M4 | §9 baris Q-22; §10.2 butir 12 | Label usang diperbarui |
| M5 | §9 baris S2-1 (iv), (v) | Kata-kata pemilik produk atas butir (iv) dikutip; "irama penghasilan tidak diandaikan" dinyatakan sebagai bacaan lead |
| Jawaban 4, SP-40 (b) | §2.3, §3.10, T-03, §10.3 butir 5, CTH-40, §9 | Bawaan "selisih ke bagian bersama" dicabut; pemegang memilih setiap kali |
| SP-40 (a), (d) | §5.3 B6 (baru), §2.3, §10.3 butir 5 | Kebocoran lewat melihat dompet; "server tidak menerima `C`" bukan perlindungan terhadap pengelola server |
| SP-32, SP-34 | §3.4, INV-30, §10.3 butir 11 (baru) | Jatah yang tidak di atas ambang hanya diusulkan pemiliknya (guard dokumen 06, dikutip); ketimpangan dinyatakan |
| SP-33 | §3.4 (tabel, `ada_jatah`), INV-30, T-29, §10.3 butir 7 | `assessment_until` hanya bila jatah yang diakhiri di atas ambang |
| Jawaban 2, SP-37 | §3.4, §3.8, §10.3 butir 7 dan 10 | Tanda dasar `personal_share` dan `routine_amount` tanpa pemberitahuan langsung (aturan dokumen 06, dikutip) |
| SP-38 (jawaban 3 hanya untuk pengingat; bagian ini **[Usulan]**) | §3.8 "Penautan" butir 1, INV-26 | Konfirmasi susulan disetujui dan ditutup dalam satu transaksi; otorisasinya memuat tepat satu tautan; pelaku tautannya pemohon, yang harus pencatat entri itu |

**Angka dan tanda yang berubah atau baru** (untuk hitung ulang lead):

| Tempat | Sebelum | Sesudah |
|---|---|---|
| Semua tabel saldo, alokasi, dan ringkasan CTH-01 sampai CTH-41; CTH-37 | — | **Tidak ada angka yang berubah** |
| CTH-21, CTH-22, CTH-23 | Tanda tidak disebut | Tidak ditandai: tiap 75000 dalam batas 100000, `N` = 0. Hasil sama dengan aturan lama |
| CTH-24, 230000 atas batas 200000 | Tanda tidak disebut (aturan lama: tidak ditandai karena tertaut) | Tidak ditandai: `N` = 30000, tidak di atas 500000. Hasil sama, dasar berubah |
| CTH-24, varian baru | — | 800000 atas batas 200000: `N` = 600000 > 500000, **ditandai**, dasar `threshold` (aturan lama: tidak ditandai) |
| CTH-33, 2000000; CTH-34, 4360000 | Tidak ditandai (tertaut) | Tidak ditandai: batas tidak kurang dari jumlah pemakaian, `N` = 0. Hasil sama |
| CTH-40 | Pilihan 1 berlabel "bawaan" | Tiga pilihan setara tanpa bawaan; angka tidak berubah |
| INV-30, uji jatah yang diakhiri (revisi kedua: jatah 500000, ambang 500000, 100000 dalam periode yang sama bertanda) | Bertanda | **Tidak bertanda**: 500000 tidak di atas ambang 500000, pengakhiran berlaku seketika. Uji diganti: jatah 1000000 habis lalu diakhiri, 100000 dalam periode yang sama bertanda (1100000 > 1000000); hari pertama periode berikutnya tidak bertanda |
| INV-30, uji baru | — | Penyisihan 800000 tertaut ke otorisasi berbatas 200000: bertanda (600000 > 500000); 230000 tertaut ke otorisasi **lain** berbatas 200000 tanpa realisasi lain: tidak bertanda (30000). Diperbaiki pada tinjauan ulang `critic`: rumusan semula "atas otorisasi yang sama" memberi `N` = 230000, bukan 30000 |
| §10.3 butir 11 (baru) | — | Ambang 500000, lima kali 400000: tanpa jatah 0 tanda; jatah 1000000 3 tanda (1200000, 1600000, 2000000); jatah 300000 5 tanda |

Pengenal baru: B6 (§5.3); butir §10.3 nomor 11. T-03 diperbarui; T-29 dipersempit. INV-26 dan INV-30 berubah isi, nomornya tetap.

**Tambahan atas laporan `architect` (dokumen 06 §15.4 butir (v)).** (1) Label: penutupan otorisasi konfirmasi susulan di transaksi persetujuan, batas yang ditetapkan server, dan tepat satu tautan adalah **[Usulan]** `security-privacy` (SP-38) dan dokumen ini, bukan "diterima pemilik produk atas saran lead"; jawaban 3 pemilik produk hanya mencakup pengingat (tabel label, §3.8, §9, tabel di atas). (2) §3.15 aturan 7a: jenis yang dapat ditautkan ke tagihan berulang adalah `expense`, `contribution_expense`, `advance`; `contribution_expense` tertaut tidak pernah ditandai (**[Usulan]**); fungsi tanda §3.8 diselaraskan. Pelaku tautan konfirmasi susulan adalah pemohon, yang harus pencatat entrinya (§3.8, INV-26). (3) §3.4 mengutip guard K01 dokumen 06 di setiap tempat yang menyebut jatah 0 atau jatah yang tidak di atas ambang; batas N-10 dirumuskan: pemberitahuan langsung hanya untuk dasar `threshold`, paling banyak satu per penerima per hari. **Tidak ada angka atau tanda yang berubah**: tidak ada contoh yang memuat `contribution_expense` tertaut ke tagihan berulang.

**Tinjauan ulang `critic` (tanpa temuan menghalangi).** (1) Batas pagar K01 disamakan dengan batas pengakhiran seketika: "tidak di atas ambang" (≤ ambang, termasuk 0), di §3.4, INV-30, §9, §10.2, §10.3 butir 11, tabel §11, dan glosarium *jatah pribadi*; sebelumnya kalimat pagar berbunyi "di bawah ambang" (keputusan lead, **[Usulan]** `security-privacy`). (2) Uji INV-30: penyisihan 230000 kini ditautkan ke otorisasi **lain** berbatas 200000, sehingga `N` = 30000 dan kesimpulan "tidak bertanda" tetap; rumusan semula ("otorisasi yang sama", sesudah 800000) memberi `N` = 230000. (3) §3.4: dokumen 06 revisi kedelapan. (4) §3.8: "untuk diselaraskan" dihapus; dokumen 06 mengikuti. (5) §3.4: kalimat pemberitahuan yang berulang dirapikan menjadi satu. (6) §3.8: permintaan konfirmasi susulan yang menyebut tempat uang sumber ditolak di B01 dokumen 06 (**[Usulan]** `security-privacy`, SP-38, dan dokumen ini), menggantikan "reservasi lahir dan dilepas". Angka lain tidak berubah.

### Putaran penyebaran S3 (12 November 2026)

Dasar: celah aturan uang yang ditemukan dokumen 05 (`FIN-1.md` §13, `FIN-2.md` §12, `FAM.md` PA-FAM-7), ketidakcocokan yang dilaporkan fase B ([_progress.md](_progress.md), "Fase B selesai"), dan permintaan dokumen 11 §10.1. Ditinjau `critic` A; angka dihitung ulang lead dengan skrip (12 November 2026), semuanya cocok. Pengenal celah 05 di tabel ini memakai skema README (`PA-FIN-n`, `PA-FAM-n`), diganti nama di putaran perbaikan berikut. Semua aturan baru **[Usulan]**, putaran S3. Pernyataan pemilik produk tentang chat (12 November 2026) tidak mengubah aturan uang mana pun dan tidak dipakai di sini.

| Celah | Bagian | Perubahan |
|---|---|---|
| PA-FIN-2 | §2.5 tabel penanda baris (a), §3.3, §3.17 | Saldo awal atas tempat uang yang dipegang orang lain bertanda belum dikonfirmasi |
| PA-FIN-3 | §3.17, §3.12 (dirujuk), INV-31, INV-19, §6.3 | Tepat satu saldo awal per tempat uang, paling banyak satu per posisi lain, jumlah ≥ 0; koreksi hanya penyesuaian; pembalik ditolak; saldo awal di luar tabel A dan arus kas, disebut di bawahnya; INV-19 dikoreksi untuk saldo awal yang bertanggal di dalam rentang |
| PA-FIN-4 | §2.6, §3.10 butir 6, INV-21 | Perincian selisih atas tempat uang terarsip boleh; pengaktifan kembali sebagai peristiwa; tempat uang terarsip beralokasi nol |
| PA-FIN-5 | §2.5 (`bill_occurrence_due`, `occurrence_due`), §2.6 (`due_rule`, `purpose.recurring_bill_id`), §3.15, INV-26 | Kejadian tagihan bulanan; pembayaran menunjuk kejadiannya; definisi "belum dibayar periode ini"; jumlah yang masih ditawarkan dikurangi alokasi tujuan tagihan. Bawaan dokumen 05 diganti, alasannya dinyatakan |
| PA-FIN-6 | §3.2 | Dua entri |
| PA-FIN-7 | §2.6 (`category`), §3.3, §3.14, §6.3 tabel A dan B, T-10 | "Kategori bawaan 'biaya admin'" dan "bawaan 'cicilan'" dicabut; baris "tanpa kategori" |
| PA-FIN-8 | §3.10 butir 8 | Satu keterangan di layar anggaran dan pada tempat uangnya |
| PA-FIN-10 | §2.3, INV-32 | Tempat uang pribadi selalu dipegang pemiliknya |
| G-01 | §2.6 tabel pengarsipan, §3.9, §6.3 tabel D | Rumusan "alokasi berakhir di akhir periodenya" dicabut; arsip tujuan bila alokasi nol; pelepasan tanpa pilihan terisi |
| G-02 | §2.6 (`reconciliation`), §3.10, INV-16 | Selisih dihitung server; perangkat hanya memperkirakan |
| G-03 | §3.7, §3.8 tabel *jumlah pemakaian*, INV-26, CTH-25 teks | `write_off` dibandingkan dengan ambang; tidak ditautkan ke otorisasi pinjaman asalnya; baris `debt_forgiven` (nol) ditambahkan ke tabel agar lengkap |
| G-05; 07 §3.10 butir 1 lawan §2.3 | §2.3, §2.5 baris (c), §3.3, §3.10 butir 1, §5.1 | Satu aturan siapa yang menghitung; selisih oleh bukan pemegang bertanda di sisi pemegang |
| G-07 | §3.14 | Biaya kredit tidak pernah negatif; cicilan 0% tanpa baris biaya kredit |
| 07 §2.6 lawan 06 §6.3 "Pos rutin" (d) | §2.6, §3.15 | Tagihan tidak terarsip sendiri |
| 04 F-FIN-012 "sisa pokok" | §10.2 | Permintaan rumusan kepada `product-analyst`; 07 §3.14 tidak berubah |
| PA-FAM-7 | §8.2 butir 2a, tabel `exit_closure`, INV-28, CTH-38 teks | Penutupan pada setiap keluar; penghapusan atas permintaan atau pada akhir retensi |
| Dokumen 11 §10.1 | §3.1 (THR), §3.10 (merinci sekaligus), §3.15 (tanpa `due_rule`), §10.4 | THR dan bonus `salary`; bentuk satu entri `expense`; perubahan minimal bila IDEA-03 diterima; lima butir ditunda |
| — | §9, §10.1, §10.2, §10.4, §7.1 | Baris §9; T-30 sampai T-38; permintaan; pertanyaan checkpoint S3; peta kasus |

**Angka dan tanda yang berubah atau baru** (untuk hitung ulang lead):

| Tempat | Sebelum | Sesudah |
|---|---|---|
| Semua tabel saldo, alokasi, dan ringkasan CTH-01 sampai CTH-41; CTH-37 | — | **Tidak ada angka yang berubah** |
| CTH-25, teks tanda | Tidak disebut | Penghapusan 100000 tidak ditandai (100000 tidak di atas 500000); varian 600000 ditandai |
| CTH-42 (baru) | — | Satu tabel: `TUN-S` 0 → 200000 → 180000; dua perintah ditolak (Δ 0) |
| CTH-43 (baru) | — | Bagian A: `EWL-X` 30000 → 0, `UANG` 530000 → 500000, perincian Δ 0. Bagian B: `UANG` 550000 → 535000; `EWL-X` 0 → −15000 → 0; `TUN-I` 550000 → 535000 |
| CTH-44 (baru) | — | `TUN-I` 500000 → 300000; `PRI-I` 250000 → 150000; tabel A −300000 + 100000 = −200000 |
| CTH-45 (baru) | — | Belum dialokasikan `TUN-I` 1000000 → 1300000, `TAB-I` 0 → 5500000; kemajuan 5800000 → 0 |
| CTH-46 (baru) | — | `TUN-S` 900000 → 700000; `TUN-I` 500000 → 700000; selisih server 0; varian perangkat yang ditolak: 900000 dan 1600000 |
| CTH-47 (baru) | — | `BNK-I` 3800000 → 3750000 |
| CTH-48 (baru) | — | `TUN-I` 450000 → 285000; tabel pembanding tujuh baris |
| CTH-49 (baru) | — | Bagian A: `AST` 3600000, `KWJ-C` 3600000, `KB` 500000. Bagian B: `TUN-I` 100000, `AST` 4000000, `KWJ-C` 3600000, `KB` 500000 |

Pengenal baru: INV-31, INV-32; T-30 sampai T-38; CTH-42 sampai CTH-49; §7.15; §8.2 butir 2a; §10.4; baris (c) tabel penanda §2.5. INV-16, INV-19, INV-21, INV-26, INV-28 berubah isi, nomornya tetap. T-10 diperbarui.

### Putaran perbaikan S3, fase D (12 November 2026)

Dasar: tinjauan `critic` A (CA-11), pemeriksaan `security-privacy` atas 05 dan 06 (SP-48 sampai SP-51), celah PA-FIN-11 dari `FIN-1.md`, 06 §6.3 "Pos rutin" (g), dan skema pengenal README. Sesudahnya ditinjau ulang `critic` dan angkanya dihitung ulang lead dengan skrip (12 November 2026; [_progress.md](_progress.md): 58 tabel, 316 baris, 0 kesalahan). Semua aturan baru lahir sebagai **[Usulan]**; yang perlu pemilik produk dijawab di checkpoint S3-3, S3-4 (c), dan S3-9 (bagian berikut).

| Butir | Bagian | Perubahan |
|---|---|---|
| CA-11 | CTH-44 teks | "(500000 − 300000)" menjadi "(300000 − 500000)"; angka tabel tidak berubah |
| SP-51 | §3.7, §3.8 tabel *jumlah pemakaian* dan "Tanda atas entri" (`H`), INV-34 (baru), T-32, CTH-25 teks dan varian SP-51 (baru), §10.4 butir 3 | Penghapusan piutang dinilai kumulatif per piutang atas bagian yang belum disetujui; entri ditandai bila `N > 0` dan `H + N > ambang`. Tanda jenis lain tidak berubah (pemeriksaan kesetaraan di §3.8) |
| SP-50 (a), (b) | §2.3 (`others_may_count`; dompet campuran), §2.5 baris (c), §3.10 butir 1 (syarat (a) sampai (c)), INV-33 (baru), T-33, CTH-47, §10.4 butir 2 | Hitungan bukan pemegang hanya atas rekening atau dompet digital yang ditandai pemegangnya; paling banyak satu terbuka per tempat uang; tidak atas pemegang dalam masa tenggang |
| SP-48 | §8.2 butir 1a dan 2b (baru), butir 2a, tabel `exit_closure` baris "Kapan", §8.3 (baris baru), INV-28, T-37, T-39 (baru), CTH-38, §10.4 butir 4 | Masa tenggang keluar 7 hari **[Asumsi]**: akses berakhir seketika, `exit_closure` dan penutupan pada akhirnya, buku pribadi dibekukan, angka penutupan tetap; ekspor buku rumah tangga sampai saat akses berakhir |
| SP-49 | §8.2 butir 3a (baru), §8.3 (baris baru) | Anggota yang bergabung kemudian membaca inti moneter, bukan rincian tulisan mantan anggota |
| PA-FIN-11 | §3.15 tabel kejadian (baris baru "Urutan bawaan kejadian"), T-40 (baru) | Kejadian bawaan: paling awal yang belum dibayar sampai akhir periode tanggal kejadian pembayaran; bila tidak ada, paling awal sesudahnya; bila tidak ada juga, kejadian terakhir. "Periode berjalan" kini dibaca sebagai periode tanggal kejadian pembayaran, sama untuk entri terlambat |
| 06 §6.3 "Pos rutin" (g) | §2.6 tabel pengarsipan, baris "Jadwal penyisihan" | Jadwal yang dirujuk jatah yang berlaku hanya diarsipkan pemilik jatah, bersama pengakhiran jatahnya; akibatnya sama dengan mengakhiri jatah sendiri |
| Skema pengenal README | Seluruh dokumen, termasuk tabel riwayat ini | PA-n (FIN-1) menjadi PA-FIN-n; P-FAM-n menjadi PA-FAM-n. Dokumen ini tidak memuat P-n FIN-2 atau PA-S3-n |
| — | Status, "Dibaca sebelum menulis", §7.1 peta kasus, §9, §10.1, §10.2, §10.4 | Baris §9; permintaan; pertanyaan checkpoint menunjuk S3-3, S3-4 (c), S3-9 |

**Angka yang berubah atau baru** (untuk hitung ulang lead):

| Tempat | Sebelum | Sesudah |
|---|---|---|
| Semua tabel saldo CTH-01 sampai CTH-49 yang sudah ada | — | **Tidak ada angka yang berubah**; CTH-38 dan CTH-47 hanya mendapat baris baru |
| CTH-44 teks | (500000 − 300000) | (300000 − 500000) = −200000 |
| CTH-25 teks | Penghapusan sisa 200000 tidak ditandai | Sama, kini dengan alasan kumulatif 100000 + 200000 = 300000 |
| CTH-25 varian SP-51 (baru) | — | `TUN-S` 1000000 → 300000; `PIU` 0 → 700000 → 400000 → 100000 → 50000; `KB` 1000000 → 1000000 → 700000 → 400000 → 350000; `H + N`: 300000, 600000, 650000; varian disetujui 300000 + 0 + 50000 = 350000 |
| CTH-38 | `exit_closure` langsung sesudah pinjaman | Baris peristiwa hari D (Δ nol; `KB` 4940000, `KB-M` 300000), lalu `exit_closure` hari D + 7 dengan angka yang sama (−140000, +140000) |
| CTH-47 | `BNK-I` 3800000 → 3750000 | Ditambah: hitungan ditolak (Δ 0), konfirmasi (Δ 0), hitungan −10000 → 3740000; jumlah −60000 = 3740000 − 3800000 |

Pengenal baru: INV-33, INV-34; T-39, T-40; §8.2 butir 1a, 2b, 3a; atribut `holding.others_may_count`; nama `H` di §3.8. INV-28, T-32, T-33, T-37 berubah isi, nomornya tetap.

### Pencatatan jawaban checkpoint S3 (12 November 2026)

Dasar: [_progress.md](_progress.md) bagian "Checkpoint S3 dijawab". Jawaban pemilik produk, kata demi kata: "lanjut", sesudah lead menyatakan bahwa "lanjut" berarti semua bawaan diterima. Kedudukannya *diterima lewat "lanjut"*, tidak pernah pernyataan pemilik produk sendiri. Tidak dijawab di checkpoint: S3-10 dan pertanyaan sampingan S3-3. Sesudahnya, pada hari yang sama, tiga masukan pemilik produk dicatat juga ([_progress.md](_progress.md), "Jawaban pemilik produk atas S3-10 dan pertanyaan THR", "Pertanyaan susulan lead dan jawabannya", "Pertanyaan pemilik produk atas S3-3"), dan atas permintaan lead aturan IDEA-03 ditulis lengkap.

| Bagian | Perubahan |
|---|---|
| Status, "Dibaca sebelum menulis" | Pencatatan checkpoint S3; status fase D (ditinjau ulang `critic`, angka dihitung ulang lead) |
| Label di awal dokumen | Kedudukan ketujuh "diterima lewat "lanjut" di checkpoint S3 (12 November 2026)"; tabel butir yang diterima (S3-1 sampai S3-9 pilihan A; D-75 sampai D-97 yang merujuk dokumen ini), yang tetap **[Usulan]** karena tidak disajikan (PA-FIN-2, PA-FIN-3/INV-31, PA-FIN-8, PA-FIN-10/INV-32, G-02, pengarsipan jadwal penyisihan, penutupan otorisasi konfirmasi susulan SP-38, §3.15 aturan 6 dan 7), dan yang tidak dijawab |
| §2.6 `recurring_bill` | `due_rule` menjadi `due_rule?` (IDEA-03) |
| §2.4, §2.6 (`income_source`), §3.1 "Sumber uang masuk", INV-24, T-14, T-18 | Sumber uang masuk menjadi daftar bawaan yang dapat disunting dan ditambah, dengan "lainnya" (kata-kata pemilik produk "ada kategori default bisa edit tambah atau yang lainnya"; bahwa itu tentang sumber: bacaan lead **[Usulan]**); dua kelompok `work` dan `gift_other` (**[Usulan]** lead, untuk checkpoint S4); siapa boleh menambah, mengganti nama, mengarsipkan, dan memindah kelompok (**[Usulan]** dokumen ini, mengikuti D-89). INV-24 membaca kelompok, bukan daftar preset |
| §3.1 THR, T-34, §10.4 butir 1 | Diterima (S3-3 A), lalu **dibuka ulang sebagian**: THR sumber bawaan tersendiri di kelompok `work`, akibat uang sama. Pertanyaan sampingan dijawab sesudahnya ("thr bisa menerima bisa ngasih") |
| §3.2 "THR dan angpau yang diberikan" (baru), CTH-51 (baru), T-42 (baru) | THR yang diberikan: `expense` biasa kepada pihak luar mana pun; penyisihan bila kepada anggota; ambang per entri; musim lewat satu permintaan musiman atau alokasi; tiga jalur dana pribadi (**[Usulan]** `finance-domain`) |
| §3.7, §3.8 tabel dan pseudocode `H`, INV-34 | Diterima (S3-9 (ii)) |
| §3.8 "Penautan" butir 1 | Jalur konfirmasi susulan diterima (S3-6); penutupan di dalam transaksi tetap **[Usulan]** |
| §3.10 butir 1, INV-33 | Diterima (S3-9 (i)) |
| §3.10 "Menghitung lalu merinci" | IDEA-01 diterima (S3-2 (i)); baris dompet campuran: bacaan "uang tunai berpemilik tunggal" |
| §3.15 | Aturan 5 diterima (S3-5 (ii)); kejadian tagihan dan urutan bawaan diterima (S3-1 (d), D-97); **IDEA-03 diadopsi** (S3-2 (ii)) dan ditulis lengkap: tabel "pos rutin tanpa tanggal" (bentuk, kesepakatan, penilaian per pembayaran, kekerapan, tautan, kejadian dan pengingat, alur bagikan, ringkasan, ada-tidaknya `due_rule` **[Usulan]** T-41); CTH-50 (baru); INV-26 dipertegas: tepat satu kejadian bila ada `due_rule`, tidak ada bila tidak |
| §8.2 butir 1a, 2a, 2b, 3a; INV-28 | Diterima (S3-4 (c), D-86, D-90); 7 hari tetap **[Asumsi]** |
| §9 | Baris checkpoint S3 |
| §10.1 | T-28 dan T-30 sampai T-40 ditutup; T-29 penyempitan dan batas waktu simpan diterima (D-77, D-81); T-10 bagian putaran S3 diterima (D-85) |
| §10.3 butir 10 | Mengakhiri sendirian dan aturan ambang biasa sesudahnya diterima (S3-5 (i), (ii)); risikonya tetap |
| §10.4 | Butir 1 sampai 4 dijawab; IDEA-01 dan IDEA-03 dijawab |

**Angka yang berubah atau baru** (untuk hitung ulang lead):

| Tempat | Sebelum | Sesudah |
|---|---|---|
| Semua tabel CTH-01 sampai CTH-49, termasuk CTH-25 varian SP-51, CTH-38, CTH-47 | — | **Tidak ada angka yang berubah**. Pemetaan label: `salary` → sumber bawaan gaji; `business` → hasil usaha, kecuali penghasilan lepas CTH-39 (10 Oktober) → honor atau pekerjaan lepas (`freelance`); `gift` → hadiah; `other` → lainnya. Kelompok tiap sumber sama dengan keterlihatan lamanya, jadi tidak ada entri yang berpindah buku |
| CTH-50 (baru) | — | `TUN-S` 1000000 → 800000 → 500000 → 500000 (tautan ditolak, Δ 0) → 150000; jumlah −850000; satu tanda (350000 > `R` 300000); anggaran "listrik" sisa 400000, 100000, −250000 |
| CTH-51 (baru) | — | `BNK-S` 5000000 → 3500000 → 2500000; `TUN-S` 0 → 1500000 → 1450000 → 850000 → 350000; `UANG` 5000000 → 5000000 → 4950000 → 4350000 → 3850000 → 2850000; realisasi 50000, 650000, 1150000, 2150000; `N` 150000; tanpa permintaan: dua tanda (600000, 1000000) |

Pengenal baru: §7.16, CTH-50, CTH-51, T-41, T-42, T-43, `income_source`, kelompok `work` dan `gift_other`. INV-24 dan INV-26 berubah isi, nomornya tetap; T-14, T-18, T-34 diperbarui. Angka dihitung ulang lead dengan skrip: 60 tabel, 336 baris, 0 kesalahan.

**Perbaikan atas temuan `critic` (CP, 12 November 2026); tidak ada angka yang berubah.** CP-05: mengganti nama sumber `work` lewat kesepakatan; sumber `gift_other` tidak memakai nama yang pernah dipakai sumber `work`; kelompok selalu tampil di samping nama (§3.1, §2.6, INV-24 uji, T-43). CP-03: kalimat risiko INV-24 (§5.3) menyebut seluruh kelompok `work`; CTH-39 penghasilan lepas berlabel `freelance` (dahulu `business`); pemetaan label di tabel angka di atas; T-27 dan glosarium *penghasilan* diselaraskan. CP-06: alasan penundaan pola jatuh tempo tahunan (§10.4) ditulis ulang tanpa A-31, yang gugur; tingkat tidak berubah. CP-09: baris Status. CP-13: "ada kategori default bisa edit tambah atau yang lainnya" dikutip sebagai kata-kata pemilik produk; bahwa "kategori" berarti sumber uang masuk dilabeli bacaan lead **[Usulan]** di setiap tempat.

