# 10 — Keamanan dan privasi

| | |
|---|---|
| Status | Draf kedua, tahap S4, putaran perbaikan fase D (12 November 2026), sesudah tinjauan `critic` B (CS-B-01 sampai CS-B-09, CS-B-13, CS-B-14, CS-B-18) dan `critic` A (CS-A-03, CS-A-04, CS-A-12), serta permintaan 08 §23.1 dan 09 indeks §8. Belum ada isi yang diterima pemilik produk kecuali yang disebut dengan kedudukannya |
| Pemilik | `security-privacy` |
| Tingkat kedalaman | Tier 1, penuh (brief §1.1); bagian Tier 2 (`CARE`, `AI`) ringkas tetapi lengkap untuk pengamannya; Tier 3 hanya syarat sebelum dibangun |
| Masukan | Brief §1.1, §4, §5K, §5O, §8; `ENGINEERING-STANDARDS.md` §4; dokumen 00, 05, 06 (revisi kedua belas), 07, 08, 09, 11, 17, 18 (termasuk §2.12); ADR-007, ADR-011, ADR-016, ADR-020, ADR-024; Permenkominfo 5/2020 naskah 2020 (dibaca 12 November 2026, 10.4 §3.4) |
| Dipakai oleh | Dokumen 09 (`architect`) dan 08 (`ux-designer`); 05 dan 17 (`product-analyst`) |
| Bukan nasihat hukum | Setiap butir hukum memakai temuan dokumen 18 atau naskah yang dibaca dengan tanggal aksesnya, dan berakhir pada pertanyaan advokat (10.4 §3) |

## Peta berkas

| Berkas | Isi | Persyaratan |
|---|---|---|
| Berkas ini | Cara membaca, asas, keputusan untuk checkpoint S4, pilot lawan publik, daftar persyaratan untuk 09 dan 08, permintaan kepada dokumen lain dan pemetaan nomor L, tinjauan atas permintaan 08, status butir terbuka, pertentangan | — |
| [10.1 Ancaman dan akses](10-security-privacy/01-ancaman-dan-akses.md) | Aset, pelaku, model ancaman; kontrol akses dan isolasi rumah tangga; satu orang di beberapa rumah tangga; jejak audit; `adults` (P-38); pengukuhan V dan AT; autentikasi, sesi, perangkat, kunci aplikasi; pemulihan akun; undangan; pengelola server dan teks pemberitahuannya | SEC-01 sampai SEC-39 |
| [10.2 Data, enkripsi, notifikasi](10-security-privacy/02-data-enkripsi-notifikasi.md) | Kelas data sensitif; log; enkripsi dan cadangan; basis data perangkat; E2EE (Q-57); privasi notifikasi dan pintasan; kehadiran dan tanda dibaca; chat dan foto | SEC-40 sampai SEC-69 |
| [10.3 Anak, penyalahgunaan, siklus data](10-security-privacy/03-anak-penyalahgunaan-siklus.md) | Anak, Kepmen 142/2026, rating usia toko; yang tidak pernah terlihat; wewenang dua arah; skenario paksaan; retensi, satu aturan jalan ekspor, pengecualian hapus seketika, penghapusan sampai cadangan, rumah tangga kosong, anggota yang wafat; serah terima kepemilikan | SEC-80 sampai SEC-110 |
| [10.4 `AI`, toko, regulasi, perilaku](10-security-privacy/04-ai-toko-regulasi-perilaku.md) | Aliran data `AI`, batas biaya, daftar kalimat terlarang, tanda tekanan berat, layanan bantuan; toko aplikasi dan Apple 1.2; ringkasan regulasi, Permenkominfo 5/2020, butir advokat; sembilan perilaku brief §8; bahan DPIA | SEC-120 sampai SEC-154 |

---

## 0. Cara membaca

**Pengenal.** Dokumen ini memakai empat awalan, semuanya hanya di dokumen 10: `AS-NN` aset, `PL-NN` pelaku, `TM-NN` ancaman, `SEC-NN` persyaratan (sudah ada di tabel pengenal `README.md`). Nomor `SEC` dikelompokkan per berkas; celah di antara kelompok disengaja. Keputusan untuk checkpoint diberi huruf `10-A` sampai `10-G`; nomor `S4-N` dan `Q-NN`-nya diberikan pemilik dokumen 17. **Butir advokat** memakai nomor dokumen 17 bila isinya sudah ada di sana (L-14, L-26 sampai L-30); butir baru memakai nomor final **L-32 sampai L-38**, yang sudah ditulis `product-analyst` di 17. Pemetaan dari usulan draf pertama: §7.1.

**Label.** Mengikuti brief §16: **[Terverifikasi]**, **[Asumsi]**, **[Usulan]**, **[Perlu validasi]**. Setiap persyaratan punya kolom label. "**[Usulan]**" tanpa nama berarti usulan `security-privacy`.

**Kutipan pemilik produk.** Hanya dari `_progress.md` atau dokumen 17, dengan kedudukannya: diputuskan sendiri, diterima atas saran lead, atau diterima lewat "lanjut". Dokumen ini tidak mengubah keputusan pemilik produk; bila riset menyarankan perubahan, ia ditulis sebagai keputusan untuk checkpoint S4 (§3), dengan keputusan yang berlaku sebagai bawaannya.

**Kolom "Tahap".** "Pilot" berarti wajib sebelum pilot memproses data nyata; "Publik" berarti wajib sebelum tahap publik; "dianjurkan" tidak menghalangi. Persyaratan Tier 2 berlaku bila fiturnya dibangun. Beberapa persyaratan toko berlaku "sebelum TestFlight eksternal atau App Store", yang dapat jatuh sebelum tahap publik.

**Kolom "Untuk".** Dokumen atau pemilik yang menerapkannya: 09 (`architect`), 08 (`ux-designer`), 05, 06, 07, atau Ops (operasi server dan runbook).

## 1. Asas

Semua persyaratan menurun dari tujuan di `CLAUDE.md`: pasangan mengetahui keadaan uang dan urusan rumah tangga tanpa saling bertanya, dengan usaha mencatat serendah mungkin, dan aplikasi membantu mereka saling memperhatikan. Aplikasi yang dapat dipakai satu orang untuk mengawasi atau mengendalikan yang lain merusak tujuan itu, jadi pencegahan penyalahgunaan diperlakukan setara dengan pertahanan terhadap orang luar.

1. **Dua lawan, setara.** Orang luar, dan anggota rumah tangga yang memakai aplikasi untuk mengendalikan (10.1 §2, PL-02).
2. **Yang ditutup benar-benar ditutup, dan yang tidak ditutup dikatakan.** Tidak ada janji ketertutupan yang tidak dapat ditegakkan; kebocoran yang tersisa dinyatakan di layar (07 §5.3; SEC-87).
3. **Pribadi sampai pemiliknya memilih penerima**, per jenis, per penerima, dapat dicabut seketika dan senyap (brief §5K; 06 V-6, V-7).
4. **Menambah pengawasan atas orang lain memerlukan semua; mengurangi pengawasan atas diri sendiri boleh sendirian** (06 PG-14). Dirumuskan ulang dengan batas periode dan aturan dasar di 10.3 §2.4.1 (K-34 titik (2), (3)).
5. **Tidak ada wewenang yang mengalir satu arah secara bawaan.** Yang satu arah ditandai (10.3 §2.3).
6. **Kredensial akun untuk tindakan yang tidak dapat ditarik**; biometrik hanya membuka aplikasi (10.1 §8).
7. **Minimum keluar dari server, dan pada pilot tidak ada data rumah tangga yang keluar ke penyedia `AI`** (10.4 §1).
8. **Perlindungan dari orang lain tidak boleh menjadi alat atas orang lain.** Padanan "blokir" adalah menyembunyikan untuk diri sendiri dan keluar, bukan mengeluarkan orang lain (10.4 §2.1).
9. **Dirancang untuk tahap publik, dibangun untuk pilot.** Kontrol yang hanya diperlukan di tahap publik tetap dirancang di sini dan diberi tanda tahapnya (§4).

## 2. Yang ditetapkan dokumen ini

Semua **[Usulan]** `security-privacy` kecuali disebut lain. Angka bertanda **[Asumsi]** diganti bila pengukuran atau riset pengguna memberi alasan.

| Hal | Ketetapan | Rujukan |
|---|---|---|
| Siapa `adults` (P-38, Q-192) | Dikukuhkan menurut peran, dengan lima syarat | 10.1 §6 |
| V-1 sampai V-9, AT-3 sampai AT-7 | Dikukuhkan; AT-7 cukup dengan dua tambahan | 10.1 §7 |
| Satu orang di beberapa rumah tangga | `household_id` hanya pemilih yang diperiksa terhadap keanggotaan; tidak ada data lintas rumah tangga; keanggotaan lain tidak terlihat | SEC-38 |
| Jejak audit | Peristiwa dan jejak keamanan per akun hanya-menambah; jejak keamanan hanya terbaca pemiliknya; akses aparat tercatat | SEC-39 |
| Masa berlaku undangan (A-33) | 7 hari, token 128 bit, sekali pakai; teks yang terbagi tanpa nama | SEC-28, SEC-29 |
| F04 (SP-65) | Pengundang melihat email terverifikasi penerima | SEC-31 |
| Autentikasi ulang | Kredensial akun untuk keluar, batal keluar, ekspor sesudah keluar, hapus akun, dan tindakan akun | SEC-13 |
| Hash kata sandi | Argon2id dengan parameter 09 (t = 3, p = 4, m = 64 MiB) diterima | SEC-11 |
| Kunci aplikasi | Mati secara bawaan, pribadi; kode aplikasi 6 digit **[Asumsi]**; biometrik gugur bila pendaftaran berubah; salah 5 kali **[Asumsi]** berarti kredensial akun, tanpa menghapus data | SEC-19 |
| Cadangan | Terenkripsi sebelum keluar VPS, di Indonesia; jendela pemulihan 30 hari; data yang dihapus bertahan paling lama **38 hari** di salinan mana pun | SEC-43, SEC-96 |
| Basis data perangkat | iOS Class C dan Android CE (bawaan, terverifikasi), menyimpang dari anjuran Apple dengan alasan di ADR-024; status `CARE` lokal Class A bila 10-D A; dikecualikan dari cadangan dan transfer; SQLCipher tidak di pilot | SEC-46 |
| Pintasan ikon aplikasi | Statis tanpa data; kunci aplikasi diminta lebih dulu | SEC-54 |
| Metadata foto | Encoder standar Go disetujui dengan empat syarat; server menerima dan mengeluarkan JPEG saja (keluaran tanpa APPn/COM, masukan PNG ditolak; CR-A-07), dengan uji keluaran tanpa metadata | SEC-66 |
| Log | Tanpa isi peka; 14 hari | SEC-44 |
| Masa tenggang keluar | 7 hari **[Asumsi]** (SP-48, diterima lewat "lanjut" di S3) dikukuhkan; hapus seketika atas pilihan orang itu sendiri adalah pengecualian yang sudah diterima (S3-4 (c), Q-179) | 10.3 §3.1 |
| Jalan ekspor sesudah keluar | **Satu aturan**: terbuka sejak F06, tertutup pada akhir hari ke-90 sesudah F07; di aplikasi dengan kredensial akun; tidak lewat email | SEC-101 |
| Isi ekspor orang yang keluar | Termasuk rincian dan chat kedua peserta sampai F06; tanpa `CARE` (10-D A) (L-33) | SEC-102 |
| Buku pribadi yang ditutup | Dihapus 90 hari sesudah F07 | SEC-94 |
| Rumah tangga kosong (PA-FAM-5) | Semua data dihapus 90 hari sesudah F07 anggota terakhir | SEC-100 |
| Anggota yang wafat (PA-FAM-6) | Hanya bukti kematian resmi; "tidak dapat dihubungi" bukan dasar | SEC-106 |
| Serah terima kepemilikan | Tidak ada pemilik tunggal yang diserahkan; tidak ada pengalihan data pribadi; pindah pengendali hanya dengan persetujuan setiap anggota | SEC-110 |
| N-26 (P-29) | `ringkasan` sesudah F06; tanpa notifikasi atas F07 dan F08 | SEC-56 |
| Lipat N-28 (P-31) | Paling banyak satu bunyi per 60 menit per pasangan pemilik-penerima **[Asumsi]** | SEC-57 |
| Jam di chat (P-30) | Waktu server menerima; bawaan 06 dikukuhkan | SEC-64 |
| Layar luar perangkat lipat (SP-69) | Teks layar kunci yang sama; tanpa widget di Tier 1 | SEC-54 |
| Tanda dibaca kartu status | Dipertahankan dan dinyatakan; pengecualian tidak mungkin dengan tanda air | 10.2 §6.4 |
| Ketelitian "terakhir aktif" | Mengikuti UX-5 (08), dijawab bersama 10-A; bawaan keputusan pemilik yang berlaku | SEC-60 |
| Balasan sesudah status dihapus (Q-153) | Dikukuhkan; L-37 | SEC-68 |
| Riwayat suasana hati (Q-136), pelacakan siklus (Q-139) | Dikukuhkan: tidak ada | 10.2 §1.1 |
| Judul tulisan orang lain ke `AI` (P-T2-4) | Judul tidak pernah dikirim | SEC-123 |
| Batas biaya `AI` | R1: batas per pembangunan kumpulan dan bulanan, tanpa batas harian; R2: harian dan bulanan | SEC-127 |
| Perbandingan antar-orang di `AI` (P-T2-5) | Dilarang di semua keluaran, dijaga struktural | SEC-129 |
| Arahan bantuan bagi status berat (P-T2-9) | Ya, hanya bagi pemilik status; halaman Bantuan statis; tanpa pemindaian teks | SEC-133 sampai SEC-135 |
| Layanan bantuan | Healing119 dan 110 terverifikasi dan dicantumkan; SAPA 129 menunggu verifikasi di situs penyelenggara; 112 bukan nomor nasional | 10.4 §1.7 |
| Apple 1.2 pada chat | Sembunyikan untuk saya, lapor, kontak, keluar; tanpa pemindaian dan tanpa mengeluarkan orang dewasa (keputusan 10-G) | SEC-153, SEC-154 |
| Rating usia toko | Apple 18+ lewat override; "Dating and Relationships" dijawab sebelum 12 Januari 2027; Google target audiens 18+ saja | SEC-83 |
| Hitungan oleh bukan pemegang (T-33) | Dikukuhkan | SEC-91 |
| Kecocokan kategori pos rutin (SP-44) | Wajib cocok | SEC-92 |
| Asas pengawasan lawan penegasan 1 dan S3-5 (i) (K-34 titik (2), (3)) | Keduanya bertentangan dengan bunyi PG-14 tetapi terbatas; keduanya dipertahankan; PG-14 dirumuskan ulang (batas periode, aturan dasar) lewat ADR; tanpa pertanyaan baru | 10.3 §2.4.1 |
| Jumlah per sisi keluarga (IDEA-07) | Tidak pernah dihitung | SEC-93 |
| Kategori bertema agama (IDEA-09) | Tidak di katalog bawaan; tidak ke `AI` | SEC-90 |
| Anak kerabat (IDEA-08) | Satu pihak luar kelompok, tanpa nama anak | SEC-85 |
| Anak dan Kepmen 142/2026 | Syarat melarang di bawah 18; pernyataan usia; desain sasaran dewasa dengan bukti | SEC-80 sampai SEC-84 |
| Permenkominfo 5/2020 | Pasal kewajiban dibaca; Narahubung dan jalur pemutusan akses masuk rencana PSE; penerapan pada chat tertutup: L-38 | 10.4 §3.4; SEC-147 |

## 3. Keputusan untuk checkpoint S4

Tujuh keputusan bagi pemilik produk, masing-masing dapat dijawab dengan satu huruf. "Bawaan" berlaku bila ia menjawab "lanjut" atau tidak menjawab; untuk perubahan atas keputusannya sendiri, bawaannya adalah keputusan yang berlaku.

### 10-A — Kehadiran di chat (Q-174; P-28; SP-61)

Keputusan pemilik produk sendiri, "Semua seperti WhatsApp" (`_progress.md`, **diputuskan pemilik produk sendiri**, berlawanan dengan saran lead). Risikonya: online pukul 23.00 memperlihatkan bahwa ia memegang aplikasi selama jam tenang; online sesudah membaca kartu yang ia tunda; online sesudah membagikan "butuh tenang dulu"; jam pencatatan entri yang AT-6 sembunyikan; pola tidur yang dapat dicatat pasangan sendiri walau server tidak menyimpan riwayat (10.2 §6).

| Huruf | Pilihan |
|---|---|
| **A (bawaan)** | Tetap seperti sekarang (P0): online, mengetik, dan terakhir aktif digerakkan pemakaian di bagian mana pun aplikasi, tampil di layar chat |
| **B (disarankan)** | P2: online dan terakhir aktif hanya digerakkan pemakaian **di layar chat**; mencatat atau membuka bagian lain tidak terlihat. Tanpa setelan baru |
| C | P1: setiap orang boleh menyembunyikan kehadirannya; yang menyembunyikan juga tidak melihat kehadiran pasangannya |
| D | P1 dan P2 sekaligus |

**Ketelitian "terakhir aktif" dijawab bersama, sebagai UX-5 di dokumen 08, dan dinyatakan terus terang sebagai perubahan atas keputusannya.** Bawaan UX-5 A: ketelitian menit, seperti keputusan "Semua seperti WhatsApp" yang berlaku. Pilihan B dan C UX-5 membulatkan ke bawah ke 15 menit dan menghapus jam untuk kemarin dan sebelumnya (08 §10.3); dokumen ini menyarankannya (SEC-60), tetapi itu hanya berlaku bila ia memilihnya. Tanda dibaca tidak berubah di pilihan mana pun.

### 10-B — Jalur `AI` di pilot (Q-10 bagian (1)) dan batas biaya

Claude dan ChatGPT tidak dapat diproses di Indonesia lewat jalur yang ditemukan; memakainya dengan data rumah tangga kemungkinan transfer lintas batas, dan sesudah PP 33/2026 berlaku (16 atau 17 Januari 2027) peneliti tidak menemukan jalur bagi pemanggilan harian (18 §2.8, §5.9; **[Perlu validasi]**). Ada jalan yang memenuhi kedua keinginan pemilik produk (Claude atau ChatGPT, dan server di Indonesia): `AI` menulis **kumpulan kalimat tanpa data rumah tangga**, dan server di Indonesia memilih kalimat bagi tiap orang (R1; 10.4 §1.2). Sudah dipakai untuk kata penyemangat dan saran apresiasi (D-92).

| Huruf | Pilihan |
|---|---|
| **A (bawaan, disarankan)** | R1 untuk semua fitur `AI` pilot: tidak ada data rumah tangga yang keluar ke penyedia; judul objek tidak pernah dikirim (pengetatan "never names" di brief §5O keputusan 3); "minta ubah" berarti saran lain dari kumpulan. Batas biaya USD 10 per bulan, dengan batas per pembangunan kumpulan (SEC-127). Panggilan langsung (R2) dinilai lagi sebelum tahap publik sesudah advokat menjawab |
| B | Q-10 opsi B: panggilan langsung berdata minimum ke penyedia di luar negeri dengan persetujuan, sampai PP berlaku, lalu turun ke A. Biaya USD 10 per bulan. Bergantung pada L-26 |
| C | Q-10 opsi A: hanya model self-hosted atau penyedia yang memproses di Indonesia; mutu bahasa Indonesia dan kecepatan belum terbukti (18 §5.9 butir 3) |
| D | Q-10 opsi C: tanpa `AI` sampai advokat menjawab |

Bila angka batas biaya lain yang diinginkan, sebutkan bersama hurufnya.

### 10-C — Kapan advokat ditanya (Q-10 bagian (2))

| Huruf | Pilihan |
|---|---|
| A | (ii): sebelum tahap publik (bawaan 17 yang belum dikukuhkan) |
| **B (bawaan baru, disarankan)** | (iii) ditambah tiga butir yang dijawab **sebelum pilot memproses data nyata**: L-01 (pengecualian rumah tangga), L-14 (apakah pilot wajib mendaftar PSE sebelum dipakai), L-24 (operator yang juga anggota). DPIA singkat dan perjanjian prosesor ditulis selama fase bangun (SEC-150). Sisanya sebelum tahap publik |
| C | (i): semua butir sekarang, sebelum fase bangun |

Alasan B: pendaftaran PSE harus "sebelum Sistem Elektronik mulai digunakan" (Permenkominfo 5/2020 Pasal 2 ayat (3); 18 §2.11), jadi bila pilot termasuk, menunggu sampai tahap publik sudah terlambat. Dengan 10-B pilihan A, butir transfer `AI` tidak menghalangi pilot.

### 10-D — Enkripsi end-to-end (Q-57; Q-99)

Di pilot operator adalah pasangan, dan pasangan adalah penerima sah chat dan status yang dibagikan, jadi E2EE hampir tidak menutup apa pun darinya; buku dan syarat permintaan harus terbaca server karena aturan uang dihitung di sana. Satu-satunya data yang benar-benar dapat ditutup dari operator-pasangan adalah status kepedulian yang **tidak** dibagikan (10.2 §4).

| Huruf | Pilihan |
|---|---|
| **A (bawaan, disarankan)** | Pilot tanpa E2EE; **status kepedulian yang tidak dibagikan tidak pernah meninggalkan perangkat pemiliknya** (dikirim ke server hanya saat ia menambah penerima); E2EE untuk isi chat, foto chat, dan alasan bebas `CARE` dinilai dalam ADR sebelum tahap publik, sesudah L-14 dan L-38 |
| B | Pilot tanpa E2EE; status kepedulian tetap disimpan di server seperti 06 §13.5 sekarang; E2EE dinilai sebelum tahap publik |
| C | E2EE untuk chat sejak pilot: lebih rumit (pemulihan, ekspor, perangkat baru) dan tidak menutup chat dari lawan bicaranya |

Teks pemberitahuan pengelola (10.1 §11.3) dan isi ekspor (SEC-102 (b)) mengikuti huruf yang dipilih.

### 10-E — Mengakhiri kesepakatan ambang sendirian (Q-93; P-15; SP-41)

Sekarang ambang hanya berubah atau berakhir bila semua pihak setuju; satu orang dapat mempertahankannya atas pasangannya tanpa batas (10.3 §2.4).

| Huruf | Pilihan |
|---|---|
| A | Tetap: hanya dengan persetujuan semua |
| **B (disarankan)** | Setiap anggota dewasa penuh dapat mengakhirinya sendirian, berlaku mulai periode berikutnya, diberitahukan sebagai `kabar`; ambang baru perlu kesepakatan biasa. Sejak pilot |
| **C (bawaan)** | Seperti B, tetapi baru sebelum tahap publik; pilot tetap A |

### 10-F — Bentuk pemberitahuan pengelola server (Q-122; SP-42; syarat keputusan 25)

Syarat keputusan pemilik produk sendiri: istrinya diberi tahu dengan terus terang sebelum mulai memakai. Draf teks: 10.1 §11.3.

| Huruf | Pilihan |
|---|---|
| **A (bawaan)** | Ia membaca sendiri teks itu di aplikasi sebelum mengisi data pribadi; pengakuannya tercatat bertanggal dan berversi; tanpa kalimat janji |
| **B (disarankan bila pemilik produk bersedia)** | A, ditambah kalimat janji operator dan catatan operasi yang terbaca keduanya setiap kali basis data dibuka (SEC-37) |
| C | Pemberitahuan lisan saja. Tidak disarankan: tidak memenuhi FR-FAM-019 dan tidak meninggalkan bukti |

### 10-G — Apple Guideline 1.2 pada chat rumah tangga (CS-B-02; 18 §2.12.4)

Apple 1.2 meminta aplikasi dengan konten buatan pengguna punya penyaringan, pelaporan dengan tanggapan, pemblokiran pengguna yang kasar, dan kontak yang dipublikasikan, tanpa pengecualian bagi grup tertutup (**[Terverifikasi]** bunyinya); apakah App Review menerapkannya pada chat rumah tangga belum diketahui (**[Perlu validasi]**). Rancangan sekarang tidak memindai isi (SEC-135) dan tidak mengizinkan orang dewasa mengeluarkan orang dewasa lain (brief §4.2 "member exit"; 06 INV-A48). Pada pilot, pengelola server adalah anggota rumah tangga, jadi "lapor ke pengelola" berarti lapor kepada pasangan. Rincian: 10.4 §2.1.

| Huruf | Pilihan | Risiko toko **[Asumsi]**, tidak diteliti selain bunyi pedoman |
|---|---|---|
| **A (bawaan, disarankan)** | Padanan tanpa pemindaian dan tanpa mengeluarkan orang: **sembunyikan untuk saya** (pesan bebas dan foto satu anggota, senyap, kartu objek tetap tampil; SEC-153); **keluar** dari rumah tangga sebagai pemutus terakhir; **laporkan** butir yang dipilih pelapor ke kontak pengelola, tanpa memberi tahu yang dilaporkan (SEC-154); kontak dipublikasikan; foto diburamkan di perangkat bila peringatan konten sensitif sistem menyala; catatan App Review. Di pilot: sembunyikan, keluar, dan tombol lapor yang membuka halaman Bantuan; sisanya sebelum TestFlight eksternal atau App Store | Rendah sampai sedang: semua unsur 1.2 punya padanan, tetapi "blokir dari layanan" dipenuhi dengan menyembunyikan dan keluar, bukan menghapus akun lain; bila peninjau menuntut blokir harfiah, jawabannya lewat catatan App Review atau ditanyakan ke Apple (18 §6.4 butir 10) |
| B | Seperti A tanpa "sembunyikan untuk saya": hanya lapor, kontak, keluar, catatan App Review | Sedang: tidak ada padanan blokir selain keluar dari rumah tangga |
| C | Seperti A, ditambah perintah agar anggota dewasa dapat mengeluarkan anggota dewasa lain, **atas pilihan pemilik produk sendiri**. Mengubah brief §4.2 ("member exit", diubah atas persetujuannya 12 November 2026), 06 INV-A48, dan PG-13. **Tidak disarankan**: tombol itu dapat dipakai menekan dan menyingkirkan pasangan dari catatan bersama | Paling rendah untuk 1.2, karena ada blokir harfiah |
| D | Tidak menambah apa pun selain kontak; berargumen chat tertutup bukan konten buatan pengguna | Tinggi: dapat ditolak di Beta App Review atau App Review; bila L-38 menyatakan Permenkominfo 5/2020 Pasal 10 berlaku, sarana pelaporan tetap wajib |

## 4. Pilot lawan tahap publik

| Kelompok | Wajib sebelum pilot memproses data nyata | Wajib sebelum tahap publik | Dianjurkan di pilot |
|---|---|---|---|
| Akses dan isolasi | SEC-01 sampai SEC-07, SEC-09, SEC-10 (fungsi), SEC-38 (a), (b), (e), (g) | SEC-08 (RLS), SEC-10 (pemberian peran), SEC-38 sisanya bila keanggotaan ganda dibangun | SEC-08 |
| Jejak audit | SEC-39 (a), (b), (c) pilot | SEC-39 (c) publik, (d) | — |
| Autentikasi dan sesi | SEC-11 sampai SEC-17, SEC-19 sampai SEC-24, SEC-26, SEC-27 | SEC-18 (faktor kedua), SEC-25 (kode pemulihan) | — |
| Undangan | SEC-28 sampai SEC-32 | — | — |
| Operator | SEC-33 sampai SEC-35 | SEC-36 | SEC-37 |
| Enkripsi, log, perangkat | SEC-40 sampai SEC-44, SEC-46 sampai SEC-49 | SEC-45; tinjauan ulang SQLCipher; ADR E2EE (10-D) | SEC-45 bila tersedia |
| Notifikasi dan pintasan | SEC-50 sampai SEC-56, SEC-58; SEC-57 bila Tier 2 | — | — |
| Kehadiran, chat, foto | SEC-60, SEC-61, SEC-63 sampai SEC-67, SEC-69; SEC-62 dan SEC-68 bila Tier 2; SEC-153 (a) sampai (c), (e); SEC-154 (c) | SEC-153 (d) dan SEC-154 sisanya: sebelum build pertama ke TestFlight eksternal atau App Store (pilot memakai TestFlight internal, ADR-022; 10.4 §2.1) | — |
| Anak | SEC-80 sampai SEC-82, SEC-85 | SEC-83 (sebelum kiriman toko pertama; deskriptor Apple paling lambat 12 Januari 2027), SEC-84 | — |
| Penyalahgunaan | SEC-86 sampai SEC-93 | — | — |
| Siklus data | SEC-94 sampai SEC-106, SEC-108 sampai SEC-110 (a) sampai (c); SEC-105 di aplikasi | SEC-105 (tautan web), SEC-107 (salinan swalayan); SEC-110 (d) sebelum migrasi | — |
| `AI` (bila Tier 2) | SEC-120, SEC-121, SEC-123 sampai SEC-136, SEC-138 | SEC-122 (R2), SEC-137 (Tier 3) | — |
| Toko | SEC-139 (versi pilot), SEC-140 (di aplikasi), SEC-143 bila Play dipakai untuk pengujian, SEC-144 sampai SEC-146; draf SEC-141, SEC-142 | SEC-139 (versi publik), SEC-140 (tautan web), SEC-141, SEC-142 | — |
| Regulasi dan operasi | SEC-148 sampai SEC-150; L-01, L-14, L-24 menurut 10-C; verifikasi SAPA 129 (SEC-138) | SEC-147; butir advokat lain | — |
| Perilaku | SEC-151, SEC-152 | — | — |

Ponytail berlaku atas ruang lingkup bangun, tidak atas kedalaman keamanan Tier 1 (`CLAUDE.md`): tidak ada persyaratan "Pilot" di atas yang dapat dipotong untuk mempercepat pilot tanpa keputusan baru yang dicatat.

## 5. Persyaratan untuk dokumen 09 (`architect`)

Wajib dipenuhi rancangan arsitektur. Nomor dengan judul ringkas; isi dan cara mengujinya di berkas masing-masing.

| Kelompok | Persyaratan |
|---|---|
| Otorisasi dan isolasi | SEC-01 otorisasi server per perintah dan per bacaan · SEC-02 `household_id` di setiap kunci unik dan foreign key · SEC-03 pengenal tanpa waktu · SEC-04 urutan sinkron tanpa celah yang membocorkan entri pribadi · SEC-05 "tidak ada" sama dengan "tidak boleh" · SEC-06 satu jalur baca untuk layar, laporan, ekspor, `AI` · SEC-07 tanpa hitungan per orang · SEC-08 RLS (publik) · SEC-09 peran basis data bukan superuser · SEC-10 `full_adult_members` satu fungsi · SEC-38 `household_id` sebagai pemilih yang diperiksa, tanpa data lintas rumah tangga · SEC-39 jejak hanya-menambah dan jejak keamanan per akun |
| Akun | SEC-11 kata sandi dan Argon2id (parameter 09 diterima) · SEC-12 batas laju · SEC-13 autentikasi ulang · SEC-14 biometrik lokal saja · SEC-15 sesi dan rotasi token di Keychain/Keystore · SEC-16 daftar sesi pribadi · SEC-17 pemberitahuan masuk kepada pemilik saja · SEC-19 kunci aplikasi: kode 6 digit di perangkat, biometrik gugur saat pendaftaran berubah, salah 5 kali ke kredensial akun tanpa hapus data · SEC-20 pengalih aplikasi tertutup · SEC-21 hapus jarak jauh saat sesi dicabut · SEC-22 tidak ada pencabutan sesi orang lain · SEC-23 F06 mencabut semua sesi · SEC-24 pemulihan lewat email · SEC-27 tanpa pemulihan darurat oleh anggota lain |
| Undangan | SEC-28 token · SEC-30 tanggapan seragam · SEC-31 F04 dengan email · SEC-32 syarat sebelum aktif |
| Operator | SEC-33 tanpa API admin · SEC-34 pemberitahuan berversi dan pengakuan · SEC-36 akses darurat (publik) |
| Data dan enkripsi | SEC-40 TLS · SEC-41 pengerasan VPS · SEC-42 rahasia · SEC-43 cadangan terenkripsi, jendela 30 hari · SEC-44 log · SEC-45 enkripsi disk · SEC-46 kelas proteksi berkas dan pengecualian cadangan · SEC-47 perangkat hanya menyimpan yang terbaca · SEC-48 pembuangan saat akses menyempit · SEC-49 tanpa SDK pihak ketiga |
| Notifikasi | SEC-50 payload · SEC-51 teks per kelas · SEC-52 kanal per kelas · SEC-53 lencana · SEC-54 layar luar, tanpa widget, pintasan statis · SEC-55 satu notifikasi sesudah jam tenang · SEC-56 N-26 `ringkasan` · SEC-57 lipat N-28 · SEC-58 token push |
| Chat dan kehadiran | SEC-60 kehadiran dikirim server dengan ketelitian yang ditampilkan (UX-5) · SEC-61 kehadiran hanya di chat · SEC-63 chat sebagai kelas peka · SEC-64 waktu server · SEC-65 penyimpanan foto · SEC-66 encoder Go dengan empat syarat · SEC-67 pemilih foto sistem · SEC-69 foto tidak ke `AI` · SEC-153 sembunyikan untuk saya · SEC-154 lapor; pilihan 10-A, 10-D, 10-G |
| Siklus data | SEC-85 skema tanggungan minimum · SEC-86 uji "tidak pernah terlihat" · SEC-89 uji simetri berpermutasi · SEC-92 kecocokan kategori · SEC-93 tanpa agregasi sisi keluarga · SEC-94 penghapusan 90 hari · SEC-96 data terhapus paling lama 38 hari di cadangan dan salinannya · SEC-97 pemulihan tanpa menghidupkan kembali · SEC-98 perangkat · SEC-100 rumah tangga kosong · SEC-101 sampai SEC-104 ekspor dengan satu jendela · SEC-105 hapus akun · SEC-106 perintah operasi anggota wafat · SEC-108 uji hash inti · SEC-109 catatan persetujuan · SEC-110 tanpa pengalihan kepemilikan |
| `AI` | SEC-120, SEC-121 R1 · SEC-122 kunci R2 · SEC-123 yang tidak pernah dikirim · SEC-124 keluaran tidak membuka data orang lain · SEC-125 persetujuan tetap berlaku · SEC-127 batas per pembangunan dan bulanan · SEC-128 saklar mati · SEC-129 tanpa perbandingan · SEC-130 daftar terlarang otomatis · SEC-133 sampai SEC-136 tanda tekanan berat |
| Perilaku | SEC-151 otorisasi ulang perintah luring; penyempitan tidak ditahan jeda urungkan · SEC-152 penolakan tabrakan tanpa kebocoran |

Tambahan untuk ADR hosting (dari 18 §2.8): catat badan hukum pihak kontrak dan lokasi pusat data penyedia VPS dan tujuan cadangan.

## 6. Persyaratan untuk dokumen 08 (`ux-designer`)

| Kelompok | Persyaratan |
|---|---|
| Desain sasaran dewasa | SEC-82: tanpa unsur indikator D Kepmen 142/2026; catatan keputusan desain sasaran dewasa sebagai bukti |
| Keterbukaan | SEC-87 isi "siapa melihat apa" · SEC-88 kata "disepakati sebagai" dan pernyataan sebelum menyimpan · SEC-34, SEC-35 pemberitahuan pengelola server (draf 10.1 §11.3, kalimat bersyarat 10-D) · SEC-62 layar berbagi `CARE` |
| Akun dan perangkat | SEC-13 layar autentikasi ulang · SEC-14 kata layar biometrik · SEC-16 daftar sesi · SEC-19 kode aplikasi, gugurnya biometrik, dan jalan kredensial akun sesudah salah 5 kali · SEC-21 peringatan catatan belum terkirim saat mencabut · SEC-80 pernyataan usia |
| Undangan | SEC-29 teks terbagi tanpa nama · SEC-31 layar F04 dengan email · SEC-32 urutan sebelum aktif |
| Notifikasi dan pintasan | SEC-51 teks per kelas · SEC-52 nama kanal umum · SEC-53 lencana · SEC-54 layar luar perangkat lipat dan pintasan statis · SEC-56 baris N-26 di ringkasan dan layar keanggotaan tanpa menyebut masa tenggang |
| Chat | SEC-60 ketelitian terakhir aktif menurut UX-5 · SEC-61 kehadiran hanya di chat · SEC-64 waktu di chat · SEC-68 layar sebelum menghapus status · SEC-153 layar sembunyikan untuk saya, termasuk pernyataan tentang tanda dibaca · SEC-154 tombol lapor (pilot: membuka Bantuan) |
| Keluar dan data | SEC-101 jalan ekspor sesudah keluar dengan satu jendela · pilihan hapus seketika bukan bawaan dan menyebut hilangnya masa batal (10.3 §3.1) · SEC-105 hapus akun · FR-FAM-014/AC-4 batas penghapusan |
| `AI` dan bantuan | SEC-125, SEC-126 layar ikut serta · SEC-133 baris bantuan bagi pemilik status berat · SEC-134 halaman Bantuan menurut 10.4 §1.7 |
| Lain | SEC-90 katalog kategori netral · SEC-93 tanpa tampilan per sisi keluarga |

## 7. Yang diminta dari pemilik dokumen lain

Dokumen ini tidak menyunting dokumen lain. Permintaan berikut dikirim lewat lead.

| Kepada | Permintaan |
|---|---|
| `architect` (06) | N-26 menjadi `ringkasan` tanpa notifikasi F07 dan F08 (SEC-56); lipat N-28 60 menit (SEC-57); bila 10-D pilihan A: status `CARE` yang tidak dibagikan di perangkat saja (06 §13.3, §13.5); bila 10-A pilihan B atau D: §16.6 baris "Yang menggerakkan kehadiran"; daftar uji PG-14 diperluas (SP-44); kecocokan kategori tautan pos rutin (SEC-92); bila 10-E pilihan B atau C: transisi mengakhiri kesepakatan ambang sendirian; bila 10-G A atau B: perintah "sembunyikan untuk saya" tanpa Peristiwa dan tanpa mengubah tanda air dibaca (SEC-153), dan laporan tanpa Peristiwa (SEC-154); perintah penyempitan tidak ditahan jeda urungkan (SEC-151); ADR untuk pertentangan §8; **ADR rumusan ulang PG-14** (10.3 §2.4.1 butir (a) sampai (e), atau dua pengecualian terbatas bila bunyi PG-14 dipertahankan) dan dua kasus uji baru di daftar uji PG-14 |
| `architect` (09, ADR) | **CS-A-03 (SEC-127):** di 09 berkas 03 §11.1 dan ADR-021, ganti penghitung harian USD 0,33 pada jalur R1 dengan batas **per pekerjaan `ai.pool_build`** (bawaan USD 3 **[Asumsi]**) dan batas bulanan USD 10; batas harian hanya untuk jalur `ai.suggest` (R2); pekerjaan yang terhenti tidak menerbitkan kumpulan setengah jadi. **CS-A-04 (SEC-96):** di 09 berkas 04 §6 dan ADR-020, batas yang dijamin adalah **data terhapus tidak bertahan lebih dari 38 hari di salinan mana pun**; salinan pelindung di VPS cadangan hanya menambah berkas baru (bukan salinan penuh harian) dan dibersihkan di VPS cadangan menurut umur set cadangan, bukan umur salinan; uji SEC-96 ditambahkan. **SEC-66 (ADR-016 butir 12):** encoder standar Go **disetujui dengan empat syarat** SEC-66 (a) sampai (d); 09 berkas 02 :400 tetap **[Perlu validasi]** sampai uji korpus lulus (CS-A-12). **SEC-46:** catat di ADR-024 penyimpangan dari anjuran Apple (`Complete`) dan alasannya; bila 10-D A, berkas status `CARE` lokal ber-Class A. **SEC-19:** kode aplikasi dan aturan salah 5 kali menurut SEC-19 (a) sampai (e); label biometrik di 09 berkas 03 §2 dapat naik ke **[Terverifikasi]** menurut 18 §2.12.2. **SEC-11:** parameter Argon2id diterima. **SEC-38:** `household_id` di `SyncRequest` hanya pemilih yang diperiksa terhadap keanggotaan sesi (sejalan dengan CS-A-06). **SEC-39:** tabel Peristiwa dan `security_event` tanpa `UPDATE`/`DELETE` bagi peran aplikasi; retensi `security_event` 90 hari. **SEC-60:** muatan kehadiran dari server dengan ketelitian yang ditampilkan. **10-G:** ADR-022 menyebut apakah pasangan pilot dipasang sebagai penguji TestFlight internal atau eksternal, karena itu menentukan kapan SEC-153 (d) dan SEC-154 diperlukan |
| `finance-domain` (07) | SEC-92 kecocokan kategori; SEC-93 larangan agregasi per sisi keluarga; 10-E sisi uangnya; jangka 90 hari buku pribadi dan satu jendela ekspor (07 §8.2 butir 1a dan 2a merujuk dokumen 10; SEC-101); SEC-110 (c): pastikan 07 punya aturan mencatat pemegang baru tempat uang bersama yang pemegangnya keluar, tanpa pengalihan otomatis |
| `product-analyst` (05, 17) | **17:** L-32 sampai L-38 **selesai**, tertulis di 17 dengan nomor §7.1 (final); keputusan 10-A sampai 10-G (dengan UX-5) sebagai pertanyaan checkpoint S4; Q-192 selesai (10.1 §6); catat pengecualian hapus seketika sebagai bagian Q-179 yang sudah diterima. **05:** FR-FAM-003/AC-6 dipersempit (SEC-29); FR-FAM-004 menyebut email di F04 (SEC-31) dan pernyataan usia (SEC-80); FR-FAM-012/AC-7 menurut SEC-56; FR-FAM-013 dan FR-FAM-014 menurut satu jendela ekspor SEC-101, SEC-102 (tanpa `CARE` bila 10-D A), SEC-96 (38 hari); FR-FAM-017/AC-4 menurut SEC-106; FR-AI-002, FR-AI-003, FR-AI-004 menurut R1 dan SEC-123; FR-AI-007 menurut SEC-127; FR-AI-009 menurut SEC-133 sampai SEC-136; FR-CARE-003 menurut SEC-57; kebutuhan baru: kunci aplikasi (SEC-19), hapus akun (SEC-105), daftar sesi (SEC-16), halaman Bantuan (SEC-134), sembunyikan untuk saya (SEC-153), lapor (SEC-154), jejak keamanan pribadi (SEC-39 (b)), isolasi keanggotaan ganda (SEC-38), tanpa pengalihan kepemilikan (SEC-110) |
| Lead | Glosarium: *sembunyikan untuk saya*, *jalan ekspor sesudah keluar* (definisi kerja **[Usulan]**: SEC-153, SEC-101); checkpoint S4 menyajikan 10-A bersama UX-5 dengan pernyataan bahwa pembulatan mengubah keputusan pemilik produk |
| `ux-designer` (08) | §6 dan §7.2. **Fase D, susulan:** 08 §9.1 (:509) "dicatat 3 hari kemudian" tampil bagi **setiap pembaca entri**, bukan hanya pencatat (§7.2, usulan tidak diterima); teks SEC-153 (d) tidak menambah setelan atau tanda di aplikasi |
| `architect` (06, 09), susulan fase D | (a) 06 AT-6 dan 09 berkas 03 §2: usulan "hanya pencatat" tidak diterima; `recorded_on` dikirim kepada semua pembaca entri (§7.2). (b) 09 berkas 03 §1: daftar "tidak pernah ditahan" sudah sejalan dengan SEC-151; mohon uji kontrak `hold_seconds = 0` per jenis (SEC-151) dicantumkan di 09 berkas 04. (c) ADR-024 dan SEC-46 kini saling merujuk; ADR-022 TestFlight internal sudah dipakai sebagai dasar tahap SEC-153 (d) dan SEC-154. (d) Runbook rilis (dokumen 14): peran App Store Connect tersempit bagi istri pemilik produk (10.4 §2.1 "Pilot") |
| `researcher` (18) | (1) dan (2) sebagian **selesai** fase D (18 §2.11, §2.12.5, §2.12.6). Tersisa: kewajiban pemberitahuan atau persetujuan ketika pengendali beralih (SEC-110 (d), empat pertanyaan riset di butir itu); hak akses peran Developer dan Marketing di App Store Connect; ketersediaan `ContentSafetyManager` di Galaxy A55; halaman 38–41 Permenkominfo 5/2020 dan naskah 10/2021. SAPA 129 tetap menunggu pemilik produk atau operator (18 §6.4 butir 9) |

### 7.1 Pemetaan nomor butir advokat (CS-B-01)

Untuk `product-analyst`. Kolom kiri adalah nomor di draf pertama dokumen ini (12 November 2026), yang tidak dipakai lagi.

| Draf pertama 10 | Isi | Nomor sekarang | Tindakan di 17 |
|---|---|---|---|
| L-26 | Pernyataan usia cukup untuk indikator B | **L-27** | Sudah ada |
| L-27 | Wajib lapor hasil penilaian indikator | **L-28** | Sudah ada |
| L-28 | Anak memakai akun orang tua | **L-29** | Sudah ada |
| L-29 | Data anak dicatat bukan orang tua atau wali | **L-30** | Sudah ada |
| L-30 | R1 bukan transfer lintas batas | **L-32** | Tambahkan |
| L-31 | Ekspor memuat tulisan anggota yang tersisa; chat sebagai bukti | **L-33** | Tambahkan |
| L-32 | Penghapusan sampai cadangan lewat umur set | **L-34** | Tambahkan (38 hari) |
| L-33 | Hak ahli waris; prosedur SEC-106 | **L-35** | Tambahkan |
| L-34 | Kewajiban Permenkominfo 5/2020 bagi "percakapan dalam jaringan", akses aparat, E2EE | **L-14** (sudah diperluas S4) dan **L-38** (bagian *user generated content*, pelaporan, pemutusan akses, Konten Komunikasi) | L-14 tetap; tambahkan L-38 |
| L-35 | Lama penyimpanan bukti persetujuan | **L-36** | Tambahkan |
| L-36 | Hak pemilik status atas keterangan di pesan pasangan | **L-37** | Tambahkan |
| — (baru) | Permenkominfo 5/2020 Pasal 1 angka 7, 10, 11, 13, 15, 36 pada chat tertutup | **L-38** | Tambahkan |
| — | Persetujuan sebelum PP berlaku untuk `AI` luar negeri | **L-26** (milik 17) | Dirujuk 10.4 §1.8, 10-B |
| — | Build iOS lewat xtool | L-31 (milik 17) | Tidak disentuh dokumen ini |

Isi lengkap L-32 sampai L-38: 10.4 §3.3.

### 7.2 Tinjauan atas permintaan dokumen 08 (08 §23.1 dan pesan `ux-designer` sesudah perbaikan)

| Hal di 08 | Hasil | Syarat yang dapat diuji |
|---|---|---|
| "Sembunyikan angka" per perangkat (08 §12.1) | **Diterima** sebagai perlindungan dari orang yang melihat layar, bukan kontrol keamanan; teksnya tidak boleh menyebutnya demikian | (a) Setelan hanya di perangkat, tanpa Peristiwa atau notifikasi, tidak terbaca anggota lain; (b) saat tersamar, pembaca layar mengucapkan "angka disembunyikan", bukan jumlahnya; (c) lebar samaran tetap (sudah di 08); (d) tangkapan layar menampilkan samaran |
| Pintasan ikon aplikasi di bawah kunci aplikasi (08 §9.1) | **Diterima** dengan SEC-54 | Pintasan statis dengan label tetap saja; kunci aplikasi diminta sebelum lembar apa pun, termasuk sisa anggaran |
| Jeda urungkan 5 detik terhadap SEC-151 (08 §9.3) | **Diterima** dengan pengecualian | Perintah dinilai saat tiba di server; perintah penyempitan, penghapusan data sendiri, sembunyikan untuk saya, dan perintah berautentikasi ulang tidak ditahan; perintah yang tertahan saat aplikasi dihentikan tetap terkirim kemudian, tidak hilang (`ENGINEERING-STANDARDS.md`) |
| "dicatat … kemudian" **hanya bagi pencatat sendiri** terhadap AT-6 (08 §9.1; 06 AT-6 revisi ketiga belas; 09 berkas 03 §2) | **Tidak diterima** (fase D, menggantikan "Diterima" di draf sebelumnya). Tanggal pencatatan yang tampil kepada **setiap pembaca entri** adalah satu-satunya penjaga atas jalan menghindar "tanggal mundur" yang diterima untuk jatah pribadi dan ambang (06 §6.3 "Jatah pribadi", contoh baris 7; 07 §3.4, §3.11, T-26). Bila hanya pencatat yang melihatnya, pasangan tidak dapat melihat penyisihan yang ditanggalkan mundur ke periode yang jatahnya masih kosong, sehingga pengaman dua arah antar-pasangan hilang tanpa diganti. Biaya pengawasannya kecil: hanya tanggal, tidak pernah jam, dan hanya bila berbeda dari tanggal kejadian. Aturan yang berlaku tetap 06 AT-6 baris asal (SP-28) | (a) Server mengirim `recorded_on` (tanggal saja, tanpa jam) lewat jalur baca kepada **setiap anggota yang dapat membaca entri itu**, hanya bila berbeda dari tanggal kejadian; entri buku pribadi tetap hanya terbaca pemiliknya (SEC-47); (b) tidak ada medan jam pencatatan, selisih jam, atau waktu server tepat di muatan klien mana pun (SEC-03, SEC-64); (c) entri yang dicatat pada hari kejadiannya tidak membawa keterangan apa pun; (d) aturannya simetris, sama bagi kedua pencatat. Permintaan ke `ux-designer` (08 §9.1 :509) dan `architect` (06 AT-6, 09 berkas 03 §2: `pembaca = semua pembaca entri`) di §7 |
| UX-5, ketelitian "terakhir aktif" | **Diselaraskan**: SEC-60 dan 10-A kini bersyarat pada UX-5; bawaan A ketelitian menit (keputusan yang berlaku); B atau C membulatkan 15 menit (disarankan) | SEC-60 |
| Kunci aplikasi dengan kode 6 digit, biometrik opsional yang gugur saat pendaftaran berubah, batas salah 5 kali (08 §12.3) | **Diterima**; persyaratannya dinyatakan di SEC-19 (a) sampai (e) | Kode ≥ 6 digit, hanya di perangkat, kode mudah ditolak; salah 5 kali berarti kredensial akun, **tanpa menghapus data**; lupa kode sama; mematikan kunci dengan kredensial akun. Angka 6 dan 5 tetap **[Asumsi]** |

## 8. Pertentangan dengan tujuan produk atau arsitektur

Penyelesaiannya dicatat `architect` sebagai ADR; rekomendasi dokumen ini di kolom kanan.

| # | Pertentangan | Rekomendasi |
|---|---|---|
| 1 | Pengenal bercap waktu (UUIDv7 dan sejenisnya) memudahkan indeks, tetapi membuka jam pencatatan yang AT-6 sembunyikan | Pengenal yang keluar ke klien acak (SEC-03); pengenal internal bebas |
| 2 | Kursor sinkron global yang padat itu sederhana, tetapi celahnya menghitung entri pribadi pasangan (SP-40 c) | Kursor per pembaca atau opak (SEC-04) |
| 3 | Status `CARE` yang tidak dibagikan di perangkat saja (10-D A) memindahkan alur I01 sampai I04 sebelum berbagi ke perangkat | Diterima demi brief §5K terhadap operator; Tier 2, jadi tidak menyentuh bangun Tier 1 |
| 4 | Hapus jarak jauh saat sesi dicabut (SEC-21) membuang catatan yang belum terkirim, sedangkan `ENGINEERING-STANDARDS.md` mencegah kehilangan data | Diterima dengan peringatan di layar cabut; perangkat yang hilang sudah kehilangan catatan itu. Salah kode aplikasi **tidak** menghapus data (SEC-19 (c)) |
| 5 | R1 (kumpulan kalimat) berbeda dari brief §5O "asks the AI to revise it" dan dari 05 yang mengirim judul | R1 di pilot (10-B); revisi langsung dan R2 dinilai sebelum publik; lead mencatat selisih |
| 6 | Lipat N-28 (SEC-57) dapat dibaca menyimpang dari brief §1.1 "notified on every change" | Yang dilipat bunyinya, bukan kabarnya; lead menilai |
| 7 | Ekspor orang yang keluar memuat tulisan pasangan (SEC-102) lawan batas data orang lain di PP 33/2026 Pasal 76 ayat (2) | Dimuat di pilot; L-33 sebelum publik |
| 8 | Kehadiran P2 (10-A B) dan pembulatan "terakhir aktif" (UX-5 B, C) mengubah keputusan pemilik produk | Hanya dengan jawabannya; bawaan keduanya keputusan yang berlaku |
| 9 | Basis data perangkat dikecualikan dari cadangan awan: pindah ponsel tidak membawa catatan yang belum terkirim | Diterima; data lain disinkron ulang dari server |
| 10 | Apple 1.2 meminta blokir dan penyaringan; brief §4.2 ("member exit") dan 06 INV-A48 melarang mengeluarkan orang dewasa, SEC-135 melarang pemindaian | Padanan tanpa mengeluarkan dan tanpa memindai (10-G A; SEC-153, SEC-154) |
| 11 | Kelas proteksi berkas yang memungkinkan sinkron latar (Class C) lawan anjuran Apple agar data pengguna ber-`Complete` | Class C untuk basis data utama dengan alasan di ADR-024; Class A untuk status `CARE` lokal (SEC-46) |
| 12 | Salinan pelindung cadangan (perlindungan dari VPS utama yang dibobol, TM-05) lawan batas umur data terhapus (SEC-96) | Salinan hanya menambah dan dibersihkan menurut umur set, sehingga keduanya terpenuhi (38 hari) |

## 9. Status butir yang diserahkan ke dokumen 10

| Butir | Jawaban |
|---|---|
| SP-61 (Q-174) kehadiran | 10.2 §6; keputusan 10-A dengan UX-5 |
| SP-62 chat sesudah keluar dan bukti | 10.3 §3.4; SEC-102 (f); L-22, L-33 (bagian advokat SP-62 (a), (c)) |
| SP-63 pemberitahuan pengelola mencakup chat | SEC-34; draf 10.1 §11.3 |
| SP-64 (Q-175) N-26 | SEC-56 |
| SP-65 F04 identitas | SEC-31 |
| SP-66 kata `AI` tanpa tafsir dan ajakan memantau | 10.4 §1.5; SEC-130 |
| SP-67 (Q-176) jam baris peristiwa | SEC-64 |
| SP-68 teks izin `AI` | SEC-125, SEC-126 (SP-68 (a) sudah di FR-AI-005/AC-6) |
| SP-69 N-28 di layar luar perangkat lipat | SEC-54 |
| SP-39 sisa | SEC-55 |
| SP-41 (Q-93, P-15) ambang | 10.3 §2.4; keputusan 10-E |
| K-34 titik (2) dan (3); 06 §15.5 R7-6; P-26 terhadap PG-14 | 10.3 §2.4.1; ditutup bila ADR-nya dicatat; titik (1) tetap S4-9 (i) |
| SP-42 (Q-122) pemberitahuan istri | 10.1 §11; keputusan 10-F |
| SP-44 kecocokan kategori | SEC-92 |
| SP-56 (d) bagian advokat | L-37 |
| P-T2-4, P-T2-5, P-T2-9 | SEC-123; SEC-129; SEC-133 |
| FR-AI-009 tanda tekanan berat | 10.4 §1.6 |
| FR-AI-006 daftar medan per fitur | 10.4 §1.2 tabel fitur; SEC-123 |
| Q-10 (1), (2) | Keputusan 10-B, 10-C |
| Q-57, Q-99 | 10.2 §4; keputusan 10-D |
| Q-136, Q-139 | 10.2 §1.1 |
| Q-153 (P-21) | SEC-68 |
| Q-186 data anak | 10.3 §1 |
| Q-192 (P-38) `adults` | 10.1 §6 |
| A-33 masa berlaku undangan | SEC-28 |
| T-33 | SEC-91 |
| V-1 sampai V-9, AT-3 sampai AT-7 | 10.1 §7 |
| Visibilitas Peristiwa bagi yang keluar | 10.1 §7 alinea akhir |
| Penerima berbagi selain anggota dewasa penuh; `account.time_zone` | 10.1 §6 syarat 4 dan 5 |
| PA-FAM-3, PA-FAM-5, PA-FAM-6 | SEC-101, SEC-102; SEC-100; SEC-106 |
| Retensi buku pribadi, chat, rumah tangga kosong, cadangan | 10.3 §3.2 |
| Tanda dibaca kartu status | 10.2 §6.4 |
| Jumlah per sisi keluarga, anak kerabat, kategori bertema agama, IDEA-17, IDEA-20 | SEC-93; SEC-85; SEC-90; 10.3 §2.6 |
| Celah `seq` kartu status (06 INV-A47 (d)) | SEC-04; penutupnya dokumen 09 |
| Brief §8: banyak rumah tangga, serah terima kepemilikan, jejak audit | SEC-38; SEC-110; SEC-39 |
| Sembilan perilaku brief §8 | 10.4 §4 |
| Apple 1.2 (18 §2.12.4) | 10.4 §2.1; keputusan 10-G |
| Permenkominfo 5/2020 pasal kewajiban (18 §6.4 butir 7) | 10.4 §3.4; L-14, L-38 |

## 10. Yang belum dapat ditetapkan

- **Hukum.** Semua butir L, terutama L-01, L-14, L-24 (pilot), L-04, L-20 (ADR-007 tetap sementara), dan L-38 (Permenkominfo 5/2020 pada chat tertutup). Teks PP 33/2026 dibaca dari salinan teks saja; Permenkominfo 5/2020 dari naskah 2020 tanpa perubahannya.
- **Perilaku platform**: biometrik, kelas proteksi berkas, dan Keychain kini **[Terverifikasi]** (18 §2.12.2). Peringatan konten sensitif iOS (SEC-153 (d)) dan batas TestFlight kini **[Terverifikasi]** (18 §2.12.5, §2.12.6). Masih **[Perlu validasi]**: apakah ekstensi notifikasi dapat membaca berkas sebelum buka kunci pertama (tidak menghalangi, SEC-50); `ContentSafetyManager` di Galaxy A55; bahwa build TestFlight internal bebas Beta App Review (tidak dinyatakan eksplisit Apple); hak akses peran App Store Connect; isi laporan crash TestFlight.
- **Layanan bantuan**: SAPA 129 belum terverifikasi di situs penyelenggara; jam layanannya tidak diketahui; cakupan 112 terkini (10.4 §1.7).
- **Mutu R1**: apakah kumpulan kalimat cukup bervariasi dan hangat belum diuji; usulan peneliti, uji buta kecil dengan pasangan pilot (18 §5.9 butir 4).
- **Toko**: kategori deklarasi kesehatan Google; penerapan Apple 1.2 pada chat tertutup (10-G); jawaban "Dating and Relationships" (paling lambat 12 Januari 2027).
- **Keputusan 25**: apakah istri pemilik produk sudah diberi tahu belum diketahui; FR-FAM-019 menjadikannya syarat teknis sebelum data apa pun.
- **Angka [Asumsi]**: 7 hari (tenggang, undangan), 90 hari (ekspor, buku pribadi, jejak keamanan), 30 hari (jendela pemulihan), 38 hari (umur data terhapus di cadangan), 14 hari, 15 menit, 60 menit, 1 menit kunci aplikasi, 5 menit autentikasi ulang, kode 6 digit, 5 kali salah kode, USD 3 per pembangunan kumpulan.

## Riwayat

| Tanggal | Perubahan |
|---|---|
| 12 November 2026 | Draf pertama, tahap S4 |
| 12 November 2026 | Draf kedua, putaran perbaikan fase D. **CS-B-01**: butir advokat dipetakan ke L-14, L-26 sampai L-30 di 17; sisanya L-32 sampai L-37, ditambah L-38 baru (§7.1); semua rujukan diselaraskan. **CS-B-02**: keputusan 10-G (Apple 1.2), SEC-153, SEC-154, 10.4 §2.1. **CS-B-03**: SEC-38 (banyak rumah tangga), SEC-39 (jejak audit), SEC-110 (serah terima kepemilikan), TM-31. **CS-B-04**: teks pemberitahuan pengelola bersyarat 10-D. **CS-B-05**: SEC-102 (b) tanpa `CARE` di 10-D A. **CS-B-06**: satu aturan jalan ekspor (sejak F06 sampai akhir hari ke-90 sesudah F07). **CS-B-07**: hapus seketika dicatat sebagai pengecualian yang sudah diterima (S3-4 (c), Q-179) dengan pagarnya. **CS-B-08**: layanan bantuan menurut 18 §2.12.1. **CS-B-09**: SEC-83 rating usia menurut 18 §2.12.3. **CS-B-13**: SEC-66 encoder Go dengan syarat. **CS-B-14**: pembulatan "terakhir aktif" kini bersyarat pada UX-5 dan dinyatakan sebagai perubahan atas keputusan pemilik. **CS-B-18**: Permenkominfo 5/2020 pasal kewajiban dibaca (10.4 §3.4), SEC-147 diperluas. **CS-A-03**: SEC-127 batas per pembangunan kumpulan. **CS-A-04**: SEC-96 batas 38 hari dan bentuk salinan pelindung. Tinjauan permintaan 08 (§7.2) dan 09 (SEC-11, SEC-19, SEC-46, SEC-66) |
| 12 November 2026 | Susulan fase D (dari 18 §2.11, §2.12.5, §2.12.6 dan 09 indeks §8). 10.4 §3.4: sumber kanonis kini 18 §2.11; Pasal 42 dipisah sebagai kewajiban penyedia komputasi awan (VPS), bukan aplikasi; Pasal 7 ayat (2) (pemutusan akses bila tidak mendaftar) memperberat L-14 (§3.2); Pasal 9 ayat (2) (petunjuk berbahasa Indonesia) dikunci di SEC-139; Pasal 16 dan 34 ditambahkan. SEC-46 merujuk penyimpangan Class C di ADR-024 dengan alasannya. SEC-151: daftar kanonis tujuh jenis perintah pelindung yang tidak pernah ditahan, sejalan dengan 09 berkas 03 §1; P-41 diterima. §7.2: usulan "dicatat … kemudian hanya bagi pencatat" **tidak diterima** (penjaga tanggal mundur). SEC-49 dan SEC-109 (b): laporan crash otomatis TestFlight. SEC-153 (d) menurut 18 §2.12.5. SEC-153 (d) dan SEC-154 wajib sebelum TestFlight eksternal atau App Store; istri pemilik produk sebagai penguji internal dengan peran App Store Connect tersempit (10.4 §2.1). SEC-110 (d): riset yang masih diperlukan |
| 12 November 2026 | Susulan fase D (temuan `product-analyst` atas 17 K-34): 10.3 §2.4.1 menilai titik (2) (penegasan 1, jatah di atas ambang) dan (3) (S3-5 (i), mengakhiri pos rutin sendirian) terhadap PG-14; keduanya dipertahankan; rumusan ulang PG-14 butir (a) sampai (e) dan permintaan ADR kepada `architect`; tanpa pertanyaan checkpoint baru |
