# 10 — Keamanan dan privasi

| | |
|---|---|
| Status | Draf keempat, putaran perbaikan fase D S5 (12 November 2026): perintah pelindung tidak ditahan tahan tulis (SEC-151, SEC-148), SEC-08 sejak pilot, butir sebelum GR-4 (§4.2), SEC-156 langganan Plus, SEC-157 uji penetrasi, tinjauan permintaan 13, 14, 15 (§7.4). Revisi keputusan 42 (platform pilot iOS dan iPadOS): §0, §4.1, §4.2, §4.3 baru, SEC-59 baru, §7.6. Penerapan keputusan 43, 45, 46, 47, 49 (penutupan sebelum S6): §0 uji iOS tanpa Mac, §2, §3.1, §4.1 butir 7 dan 11, pemicu Q-209, §4.2, §7.4 (termasuk B-04 dan VL-02), §7.5, §7.7 baru, §10, SEC-46, SEC-49, SEC-150, SEC-157; belum ditinjau `critic`. Keputusan pemilik produk atas 10-A sampai 10-G dicatat di §3 dengan kedudukannya. Isi lain tetap **[Usulan]** `security-privacy` kecuali disebut dengan kedudukannya. Draf ketiga sudah melewati tinjauan `critic` S4; draf keempat belum ditinjau `critic` |
| Pemilik | `security-privacy` |
| Tingkat kedalaman | Tier 1, penuh (brief §1.1); bagian Tier 2 (`CARE`, `AI`) ringkas tetapi lengkap untuk pengamannya; Tier 3 hanya syarat sebelum dibangun |
| Masukan | Brief §1.1, §4, §5K, §5O, §8; `ENGINEERING-STANDARDS.md` §4; dokumen 00, 05, 06 (revisi kedua belas), 07, 08, 09, 11, 17, 18 (termasuk §2.12); ADR-007, ADR-011, ADR-016, ADR-020, ADR-024, ADR-027; Permenkominfo 5/2020 naskah 2020 (dibaca 12 November 2026, 10.4 §3.4); `_progress.md` "Checkpoint S4: jawaban pemilik produk (12 November 2026)" dan "Penutupan pertanyaan terbuka sebelum S6 (12 November 2026)" (keputusan 43 sampai 50) |
| Dipakai oleh | Dokumen 09 (`architect`) dan 08 (`ux-designer`); 05 dan 17 (`product-analyst`) |
| Bukan nasihat hukum | Setiap butir hukum memakai temuan dokumen 18 atau naskah yang dibaca dengan tanggal aksesnya, dan berakhir pada pertanyaan advokat (10.4 §3) |

## Peta berkas

| Berkas | Isi | Persyaratan |
|---|---|---|
| Berkas ini | Cara membaca, asas, hasil checkpoint S4 (10-A sampai 10-G, DPIA, risiko sisa yang diterima), pilot lawan publik dan daftar sebelum data nyata, daftar persyaratan untuk 09 dan 08, permintaan kepada dokumen lain dan pemetaan nomor L, tinjauan atas permintaan 08, status butir terbuka, pertentangan | — |
| [10.1 Ancaman dan akses](10-security-privacy/01-ancaman-dan-akses.md) | Aset, pelaku, model ancaman; kontrol akses dan isolasi rumah tangga; satu orang di beberapa rumah tangga; jejak audit; `adults` (P-38); pengukuhan V dan AT; autentikasi, sesi, perangkat, kunci aplikasi; pemulihan akun; undangan; pengelola server dan teks pemberitahuannya | SEC-01 sampai SEC-39 |
| [10.2 Data, enkripsi, notifikasi](10-security-privacy/02-data-enkripsi-notifikasi.md) | Kelas data sensitif; log; enkripsi dan cadangan; basis data perangkat; E2EE (Q-57); privasi notifikasi dan pintasan; kehadiran dan tanda dibaca; chat dan foto | SEC-40 sampai SEC-69 |
| [10.3 Anak, penyalahgunaan, siklus data](10-security-privacy/03-anak-penyalahgunaan-siklus.md) | Anak, Kepmen 142/2026, rating usia toko; yang tidak pernah terlihat; wewenang dua arah; skenario paksaan; retensi, satu aturan jalan ekspor, pengecualian hapus seketika, penghapusan sampai cadangan, rumah tangga kosong, anggota yang wafat; serah terima kepemilikan | SEC-80 sampai SEC-110 |
| [10.4 `AI`, toko, regulasi, perilaku](10-security-privacy/04-ai-toko-regulasi-perilaku.md) | Aliran data `AI`, batas biaya, daftar kalimat terlarang, tanda tekanan berat, layanan bantuan; toko aplikasi dan Apple 1.2; ringkasan regulasi, Permenkominfo 5/2020, butir advokat; sembilan perilaku brief §8; bahan DPIA; sebutan anggota per pembaca (F-FAM-021); langganan Plus (SEC-156); uji penetrasi (SEC-157) | SEC-120 sampai SEC-157 |

---

## 0. Cara membaca

**Pengenal.** Dokumen ini memakai empat awalan, semuanya hanya di dokumen 10: `AS-NN` aset, `PL-NN` pelaku, `TM-NN` ancaman, `SEC-NN` persyaratan (sudah ada di tabel pengenal `README.md`). Nomor `SEC` dikelompokkan per berkas; celah di antara kelompok disengaja. Keputusan untuk checkpoint diberi huruf `10-A` sampai `10-G`; nomor `S4-N` dan `Q-NN`-nya diberikan pemilik dokumen 17, dan hasilnya di checkpoint S4 dicatat di §3. **Butir advokat** memakai nomor dokumen 17 bila isinya sudah ada di sana (L-14, L-26 sampai L-30); butir baru memakai nomor final **L-32 sampai L-38**, yang sudah ditulis `product-analyst` di 17. Pemetaan dari usulan draf pertama: §7.1.

**Label.** Mengikuti brief §16: **[Terverifikasi]**, **[Asumsi]**, **[Usulan]**, **[Perlu validasi]**. Setiap persyaratan punya kolom label. "**[Usulan]**" tanpa nama berarti usulan `security-privacy`.

**Kutipan pemilik produk.** Hanya dari `_progress.md` atau dokumen 17, dengan kedudukannya: diputuskan sendiri, diterima atas saran lead, atau diterima lewat "lanjut". Dokumen ini tidak mengubah keputusan pemilik produk; bila riset menyarankan perubahan, ia ditulis sebagai keputusan untuk checkpoint berikutnya, dengan keputusan yang berlaku sebagai bawaannya.

**Kolom "Tahap".** "Pilot" berarti wajib sebelum pilot memproses data nyata; "Sebelum GR-4" berarti wajib sebelum rumah tangga beta memakai server yang sama (13 §6.4; §4.2); "Publik" berarti wajib sebelum tahap publik (GR-5); "dianjurkan" tidak menghalangi. Persyaratan Tier 2 berlaku bila fiturnya dibangun. Beberapa persyaratan toko berlaku "sebelum TestFlight eksternal atau App Store", yang dapat jatuh sebelum tahap publik.

**"Pilot" per platform (keputusan 42).** Pilot dengan data nyata berjalan di iOS dan iPadOS: istri di iPhone 13, pemilik produk di iPad mini 7; GR-2 tidak menunggu Android (17, keputusan 42). Karena itu "Pilot" berarti: bagi server, iOS, dan iPadOS, wajib sebelum GR-2; bagi aplikasi Android, wajib **sebelum aplikasi Android dipasang di perangkat yang memuat data nyata** (bawaan lead **[Usulan]**: sesudah GR-2, paling lambat sebelum GR-4), dengan isi yang sama dan tanpa pelonggaran. Bagian persyaratan yang hanya ada karena Android (FCM, Keystore, APK, kanal notifikasi Android, aturan cadangan Google, pemilih foto Android) bukan syarat GR-2; daftarnya di §4.3. Tidak ada persyaratan sisi iOS yang berubah.

**Uji iOS tanpa Mac (keputusan 43).** Pemilik produk memilih "Tetap Xcode Cloud saja (Saran)" (`_progress.md`, "Penutupan pertanyaan terbuka sebelum S6 (12 November 2026)", keputusan 43; diterima atas saran lead): tanpa Mac di rumah, tanpa debugger di iPhone 13 dan iPad mini 7; build, uji unit, dan XCUITest di simulator Xcode Cloud, uji perangkat lewat TestFlight (ADR-022). Tidak ada persyaratan atau uji SEC yang dilonggarkan karenanya. Uji yang menyebut isi atau perilaku perangkat dijalankan begini **[Usulan]**:

| Jenis uji | Dijalankan di | SEC |
|---|---|---|
| Isi basis data lokal, berkas, foto, dan antrean sesudah sinkron, pencabutan, keluar, atau salah kode | XCUITest dan uji unit di simulator Xcode Cloud terhadap staging; di perangkat cukup pengamatan layar | SEC-19 (c), SEC-21, SEC-47, SEC-48, SEC-98 |
| Muatan yang dikirim server (push, sinkron, kehadiran) | Uji server di CI Linux; tidak memerlukan perangkat | SEC-50, SEC-60, SEC-155 |
| Isi permintaan yang dikirim aplikasi | Perekam permintaan di dalam uji simulator | SEC-19 (kode tidak terkirim) |
| Perilaku sistem yang hanya ada di perangkat | Pengamatan tangan di perangkat pilot dengan build TestFlight | SEC-19 (b), SEC-20, SEC-51, SEC-53, SEC-54, SEC-59 |
| Kelas proteksi berkas dan tanda pengecualian cadangan | Pemeriksaan diri yang hanya ada di build aplikasi staging (SEC-41 (d)), dijalankan di perangkat pilot: menampilkan nama kelas dan nilai tanda setiap berkas aplikasi, tanpa isi, tanpa mengirim apa pun ke mana pun. Build produksi tidak memuatnya; kode pembuat berkas sama di kedua build **[Asumsi]**. Bukti kelas antrean memakai layar yang sama dengan basis data uji terpisah berisi perintah sintetis (ADR-024, "Catatan keputusan 43") | SEC-46; ADR-024 |
| Domain yang dihubungi aplikasi | App Privacy Report iOS dan iPadOS, bagian App Network Activity: "domains that have been contacted either directly or from content within an app in the past 7 days", iOS dan iPadOS 15.2 ke atas (**[Terverifikasi]**, Apple Support, "About App Privacy Report", https://support.apple.com/en-us/102188, diakses 12 November 2026) | SEC-49 |

**Pemeriksaan arsip di CI** (14 §4.8; QG-30 butir 13; **[Usulan]** `architect`): **diterima** sebagai penjaga lapis kedua, dengan tiga syarat **[Usulan]**. (a) Penjaga pertamanya tetap pengecualian saat kompilasi (target atau kondisi kompilasi tersendiri); pemeriksaan arsip hanya menangkap kesalahan konfigurasi. (b) Yang dicari bukan hanya simbol, yang dapat dibuang atau tersamar pada build rilis Swift, tetapi juga satu string penanda unik yang hanya ada di kode pemeriksaan diri; arsip produksi juga ditolak bila memuat alamat server staging (SEC-41 (d)). (c) Pemeriksaan berjalan atas arsip yang sama yang diunggah, dan kegagalannya menghentikan workflow sebelum unggahan ke TestFlight; apakah skrip sesudah build di Xcode Cloud dapat menghentikan distribusi **[Perlu validasi]**, dan bila tidak, pemeriksaan dijalankan sebagai langkah tersendiri yang menghalangi rilis menurut runbook. Uji: satu build produksi sengaja diberi penanda, lalu ditolak.

Tidak ada sertifikat root, profil konfigurasi, atau proxy pihak ketiga yang dipasang di perangkat pilot untuk menguji, karena perangkat itu memegang data nyata (uji: Pengaturan, Umum, VPN & Device Management tidak memuat profil uji). Bila satu uji SEC ternyata tidak dapat dijalankan dengan cara di atas, uji itu tidak dihapus: gerbang yang memerlukannya menunggu, dan hal itu dibawa ke pemilik produk sebagai bahan meninjau keputusan 43.

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
| Basis data perangkat | iOS Class C dan Android CE (bawaan, terverifikasi), menyimpang dari anjuran Apple dengan alasan di ADR-024; status `CARE` lokal Class A (10-D A); dikecualikan dari cadangan dan transfer; SQLCipher tidak di pilot | SEC-46 |
| Pintasan ikon aplikasi | Statis tanpa data; kunci aplikasi diminta lebih dulu | SEC-54 |
| Metadata foto | Encoder standar Go disetujui dengan empat syarat; server menerima dan mengeluarkan JPEG saja (keluaran tanpa APPn/COM, masukan PNG ditolak; CR-A-07), dengan uji keluaran tanpa metadata | SEC-66 |
| Log | Tanpa isi peka; 14 hari | SEC-44 |
| Masa tenggang keluar | 7 hari **[Asumsi]** (SP-48, diterima lewat "lanjut" di S3) dikukuhkan; hapus seketika atas pilihan orang itu sendiri adalah pengecualian yang sudah diterima (S3-4 (c), Q-179) | 10.3 §3.1 |
| Jalan ekspor sesudah keluar | **Satu aturan**: terbuka sejak F06, tertutup pada akhir hari ke-90 sesudah F07; di aplikasi dengan kredensial akun; tidak lewat email | SEC-101 |
| Isi ekspor orang yang keluar | Termasuk rincian dan chat kedua peserta sampai F06; tanpa `CARE` (10-D A); tanpa sebutan (SEC-155 (f)); tulisan anggota yang tersisa: L-33 | SEC-102 |
| Buku pribadi yang ditutup | Dihapus 90 hari sesudah F07 | SEC-94 |
| Rumah tangga kosong (PA-FAM-5) | Semua data dihapus 90 hari sesudah F07 anggota terakhir | SEC-100 |
| Anggota yang wafat (PA-FAM-6) | Hanya bukti kematian resmi; "tidak dapat dihubungi" bukan dasar | SEC-106 |
| Serah terima kepemilikan | Tidak ada pemilik tunggal yang diserahkan; tidak ada pengalihan data pribadi; pindah pengendali hanya dengan persetujuan setiap anggota | SEC-110 |
| N-26 (P-29) | `ringkasan` sesudah F06; tanpa notifikasi atas F07 dan F08 | SEC-56 |
| Lipat N-28 (P-31) | Paling banyak satu bunyi per 60 menit per pasangan pemilik-penerima **[Asumsi]** | SEC-57 |
| Jam di chat (P-30) | Waktu server menerima; bawaan 06 dikukuhkan | SEC-64 |
| Layar luar perangkat lipat (SP-69) | Teks layar kunci yang sama; tanpa widget di Tier 1 | SEC-54 |
| Tanda dibaca kartu status | Dipertahankan dan dinyatakan; pengecualian tidak mungkin dengan tanda air | 10.2 §6.4 |
| Kehadiran dan ketelitian "terakhir aktif" | Seperti WhatsApp: digerakkan pemakaian di bagian mana pun aplikasi, terakhir aktif menit tepat (10-A (i) A, keputusan pemilik produk sendiri; (ii) A, diterima lewat "lanjut") | SEC-60, SEC-61 |
| Balasan sesudah status dihapus (Q-153) | Dikukuhkan; L-37 | SEC-68 |
| Riwayat suasana hati (Q-136), pelacakan siklus (Q-139) | Dikukuhkan: tidak ada | 10.2 §1.1 |
| Judul tulisan orang lain ke `AI` (P-T2-4) | Judul tidak pernah dikirim | SEC-123 |
| Jalur dan penyedia `AI` pilot | R1 untuk semua fitur `AI` pilot; penyedia Claude (API Anthropic) (10-B A, A) | 10.4 §1.2 |
| Batas biaya `AI` | R1: USD 3 per pembangunan kumpulan **[Asumsi]** dan USD 10 per bulan, tanpa batas harian (10-B A); R2: harian dan bulanan | SEC-127 |
| Kapan advokat ditanya | Sebelum tahap publik (10-C A). DPIA singkat, catatan ketentuan prosesor, dan kebijakan privasi versi pilot tetap dikerjakan selama fase bangun, sebelum data nyata (Q-208, **diterima atas saran lead**) | §3.1; SEC-139, SEC-150 |
| Kesepakatan ambang | Dibuat, diubah, dan diakhiri hanya dengan persetujuan semua pihak, di semua rumah tangga (10-E A); risiko sisanya diterima pemilik produk | §3.2; 10.3 §2.4 |
| Pemberitahuan pengelola server | Dibaca sendiri di aplikasi, pengakuan bertanggal dan berversi, tanpa kalimat janji (10-F A) | SEC-34, SEC-35; 10.1 §11.3 |
| Sebutan anggota per pembaca (F-FAM-021, Tier 2) | Hanya terlihat pembuatnya; tidak pernah di notifikasi, layar kunci, chat, kartu objek, `AI`, atau ekspor orang lain; dihapus pada F07 salah satu pihak (bukan pada F06; F08 memulihkannya) atau saat akun dihapus | SEC-155 |
| Perbandingan antar-orang di `AI` (P-T2-5) | Dilarang di semua keluaran, dijaga struktural | SEC-129 |
| Arahan bantuan bagi status berat (P-T2-9) | Ya, hanya bagi pemilik status; halaman Bantuan statis; tanpa pemindaian teks | SEC-133 sampai SEC-135 |
| Layanan bantuan | Healing119 dan 110 terverifikasi dan dicantumkan; SAPA 129 menunggu verifikasi di situs penyelenggara; 112 bukan nomor nasional | 10.4 §1.7 |
| Apple 1.2 pada chat | Sembunyikan untuk saya, lapor, kontak, keluar; tanpa pemindaian dan tanpa mengeluarkan orang dewasa (10-G A) | SEC-153, SEC-154 |
| Rating usia toko | Apple 18+ lewat override; "Dating and Relationships" dijawab sebelum 12 Januari 2027; Google target audiens 18+ saja | SEC-83 |
| Hitungan oleh bukan pemegang (T-33) | Dikukuhkan | SEC-91 |
| Kecocokan kategori pos rutin (SP-44) | Wajib cocok | SEC-92 |
| Asas pengawasan lawan penegasan 1 dan S3-5 (i) (K-34 titik (2), (3)) | Keduanya bertentangan dengan bunyi PG-14 tetapi terbatas; keduanya dipertahankan; PG-14 dirumuskan ulang (batas periode, aturan dasar) lewat ADR; tanpa pertanyaan baru | 10.3 §2.4.1 |
| Jumlah per sisi keluarga (IDEA-07) | Tidak pernah dihitung | SEC-93 |
| Kategori bertema agama (IDEA-09) | Tidak di katalog bawaan; tidak ke `AI` | SEC-90 |
| Anak kerabat (IDEA-08) | Satu pihak luar kelompok, tanpa nama anak | SEC-85 |
| Anak dan Kepmen 142/2026 | Syarat melarang di bawah 18; pernyataan usia; desain sasaran dewasa dengan bukti | SEC-80 sampai SEC-84 |
| Permenkominfo 5/2020 | Pasal kewajiban dibaca; Narahubung dan jalur pemutusan akses masuk rencana PSE; penerapan pada chat tertutup: L-38 | 10.4 §3.4; SEC-147 |
| Tahan tulis (ADR-028 butir 3) | Tidak pernah menahan perintah pelindung SEC-151 (1) sampai (7) beserta masuk dan pembaruan token; `exit_closure` yang mengikutinya ditunda dengan tanggal saat perintah diterima (hapus akun, hapus seketika) atau saat masa tenggang berakhir (F06 biasa); tindakan pertama insiden tingkat 1 dan 2 | SEC-151, SEC-148 |
| Butir sebelum GR-4 | SEC-25, SEC-34 teks umum, SEC-36, SEC-39 (c) publik, SEC-84, SEC-139 versi publik; SEC-147 menurut L-14 (SEC-107 swalayan sudah Pilot, sejak GR-2) | §4.2 |
| Langganan Plus (Q-217) | Pembayar tidak terbaca siapa pun, termasuk lewat toko asal; data keuangan pribadi pembayar | SEC-156 |
| Uji penetrasi pihak luar | Ya, sebelum GR-5, terhadap staging dengan data sintetis; tidak untuk pilot dan GR-4 | SEC-157 |
| Ukuran validasi (VL-00 sampai VL-10) | Hanya data yang terbaca kedua anggota, lewat jalur baca aplikasi; tanpa pengumpulan baru. Catatan bug VL-02 (b) di luar aplikasi dan server, tanpa isi nyata dan nama pelapor; bug "data bocor" juga insiden SEC-148 | §7.4 |
| Platform pilot (keputusan 42) | "Pilot" bagi Android berarti sebelum Android dipasang dengan data nyata (paling lambat GR-4); bagian khusus Android bukan syarat GR-2 | §0; §4.3 |
| Wawancara dan uji prototipe statis dengan keluarga lain (Q-215) | Bukan "pengguna di luar rumah tangga pilot" selama empat syarat §7.4 terpenuhi (keputusan 47, diterima atas saran lead) | §4.1 pemicu; §7.4 |
| Penyedia hosting dan instans sebelum GR-2 (Q-220) | Penyedia dipilih sebelum deploy pertama yang memegang data nyata; sampai itu instans sementara berisi data uji, dibuang sesudahnya (keputusan 45, diterima atas saran lead) | §4.1 butir 7 |
| Uji iOS tanpa Mac (keputusan 43) | Simulator Xcode Cloud, pengamatan di perangkat, pemeriksaan diri build staging; tanpa profil uji di perangkat pilot | §0 |
| Tinjauan sebelum gerbang (keputusan 49) | Tinjauan hanya-temuan `security-privacy` sebelum GR-2 dan sebelum GR-4 | §4.1 butir 11; §4.2 |
| Tablet yang dipakai bersama (iPad mini 7 pemilik produk) | Persyaratan iPhone berlaku sama di iPad; tawaran kunci aplikasi di iPad menyebut tablet yang dipinjamkan; tanpa indeks Spotlight dan saran Siri berisi data | SEC-59 |

## 3. Hasil checkpoint S4 atas 10-A sampai 10-G

Dikukuhkan pemilik produk pada 12 November 2026 (`_progress.md`, "Checkpoint S4: jawaban pemilik produk (12 November 2026)", paragraf "Keputusan checkpoint S4 (dikukuhkan …)"). Kutipan disalin utuh dari bagian itu. Pilihan yang tidak dipilih tercatat di 17 "Checkpoint S4" dan tidak diulang di sini.

| Huruf | Pertanyaan 17 | Hasil | Kedudukan dan kutipan | Akibat di dokumen ini |
|---|---|---|---|---|
| 10-A (i) | S4-5 (i) (Q-174) | **A.** Kehadiran tetap P0: online, mengetik, dan terakhir aktif digerakkan pemakaian di bagian mana pun aplikasi, tampil di layar chat | **Diputuskan pemilik produk sendiri**, mempertahankan keputusannya "Semua seperti WhatsApp" dan berlawanan dengan saran P2: "A tetap seperti WhatsApp" | SEC-61 berlaku; P1 dan P2 tidak diterapkan. Risiko sisa 10.2 §6.2 (W-18) diterima pemilik produk |
| 10-A (ii) | S4-5 (ii) (Q-200; UX-5) | **A.** "Terakhir aktif" menit tepat; kemarin dan tanggal lama dengan jam | **Diterima lewat "lanjut"**: bacaan lead atas kalimat yang sama, dikukuhkan dengan "lanjut" sesudah rekap (17 baris S4-5 (ii)); bukan keputusan yang ia nyatakan sendiri. Saran pembulatan 15 menit tidak diambil | SEC-60 menit tepat; pembulatan tidak diterapkan |
| 10-B | S4-4 (i), (ii) (Q-10 bagian (1)) | **A, A.** R1 untuk semua fitur `AI` pilot; penyedia Claude (API Anthropic); batas USD 10 per bulan dan USD 3 per pembangunan kumpulan **[Asumsi]** | Penyedia **diputuskan pemilik produk sendiri**: "Claude, nanti aja tapi integerasiin aja dulu". Bacaan lead atas "integerasiin aja dulu" (R1 dirancang dan dibangun) dan angka batas **diterima lewat "lanjut"** sesudah rekap | 10.4 §1 tidak lagi usulan; SEC-127 berlaku dengan angka itu; R2 tetap terkunci (SEC-122) |
| 10-C | S4-4 (iii) (Q-10 bagian (2)) | **A.** Advokat ditanya sebelum tahap publik | **Diputuskan pemilik produk sendiri**, berlawanan dengan saran B, sesudah W-31 dan W-32 dibacakan: "#S4-4 (iii): nanti aja kan butuh beberaapa bulan lagi inimah" | L-01, L-14, L-24 tidak lagi syarat sebelum pilot memproses data nyata; W-32 menjadi risiko yang diterima (10.4 §3.2); DPIA dan catatan prosesor tetap fase bangun (§3.1; Q-208, diterima atas saran lead); §4 dinyatakan ulang |
| 10-D | S4-6 (i) (Q-57, Q-99) | **A.** Pilot tanpa E2EE; status kepedulian yang tidak dibagikan tidak pernah meninggalkan perangkat pemiliknya; E2EE chat, foto chat, dan alasan bebas `CARE` dinilai lewat ADR sebelum tahap publik | **Diputuskan pemilik produk sendiri** (ia menyatakan isinya sendiri; sama dengan saran tim; 17 K-64): "oke gausa e2ee" | SEC-46 (b), SEC-102 (b), dan teks 10.1 §11.3 memakai bentuk ini tanpa syarat |
| 10-E | S4-9 (i) (Q-93) | **A.** Ambang dibuat, diubah, **dan diakhiri** hanya dengan persetujuan semua pihak, di semua rumah tangga, pilot maupun publik. S4-9 (ii) gugur | **Diputuskan pemilik produk sendiri**, berlawanan dengan saran B dan bawaan C: "Ambang itu batas yang di tentukan/disepakati 2 orang"; "ambang di akhiri berdua juga" | Risiko sisa yang diterima: §3.2. 10.3 §2.3, §2.4, §2.4.1 butir (c) disesuaikan |
| 10-F | S4-6 (ii) (Q-122) | **A.** Ia membaca sendiri teks di aplikasi sebelum mengisi data pribadi; pengakuan bertanggal dan berversi; tanpa kalimat janji | **Diterima lewat "lanjut"** (tidak disebut dengan huruf) | Kalimat janji dihapus dari 10.1 §11.3; SEC-37 tetap dianjurkan, bukan wajib |
| 10-G | S4-7 (Q-201) | **A.** Sembunyikan untuk saya, lapor (di pilot membuka Bantuan), kontak, keluar; tanpa pemindaian dan tanpa mengeluarkan orang dewasa | **Diputuskan pemilik produk sendiri** (ia menyebut hurufnya sendiri; sama dengan saran tim; 17 K-64): "S4-7: A" | SEC-153, SEC-154 berlaku; F-COM-014 masuk Tier 1 (FR-COM-019, FR-COM-020) |

**Keputusan S4 lain yang menyentuh dokumen ini.** S4-6 (iii) A: kunci aplikasi, sesi, jejak keamanan milik sendiri, halaman Bantuan, dan hapus akun masuk Tier 1 (F-FAM-019, F-FAM-020; FR-FAM-021 sampai FR-FAM-024), sejalan dengan tanda "Pilot" SEC-16, SEC-19, SEC-39 (b), SEC-105, SEC-134. S4-8 (ii) A: pilihan awal penerima berbagi adalah semua anggota dewasa penuh lain, dengan berbagi tetap mati secara bawaan (10.1 §6 syarat 4 tetap). S4-8 (v) A: SEC-92 dikukuhkan. S4-1 (ii) A: tiga bantuan isian di perangkat; validasi server tetap wajib ("hp boleh ikut, validasi di app dan server wajib jadi a"), jadi tidak ada persyaratan dokumen ini yang melonggar. Tambahan pemilik produk: sebutan anggota per pembaca (F-FAM-021), SEC-155.

**Pernyataan pemilik produk tentang keputusan 25.** Dalam jawabannya atas S4-6 ia menulis "masalah istri sudah tahu gdong". Itu dicatat sebagai pernyataannya, tidak diverifikasi dokumen ini; pembacaan teks pemberitahuan di aplikasi beserta pengakuannya tetap syarat teknis sebelum data apa pun (SEC-34; FR-FAM-019; 10-F A).

### 3.1 DPIA, catatan prosesor, dan kebijakan privasi sesudah 10-C A

**Ketetapan [Usulan] `security-privacy`: ketiganya tetap pekerjaan fase bangun dan selesai sebelum pilot memproses data nyata** (SEC-139 versi pilot, SEC-150). Yang ditunda ke tahap publik hanya tinjauan advokat atas teksnya.

| Alasan | Isi |
|---|---|
| 1. Tidak memerlukan advokat | Jawaban pemilik produk mengatur **kapan advokat ditanya**. DPIA, catatan ketentuan prosesor, dan kebijakan privasi adalah dokumen pengendali sendiri; bahannya sudah ada di dokumen ini (10.4 §5; 10.2 §1; 10.3 §3.2) |
| 2. Waktunya ditentukan teks, bukan advokat | Salinan teks PP 33/2026 Pasal 121 ayat (1): DPIA "dilaksanakan sebelum dilakukan pemrosesan Data Pribadi" (**[Perlu validasi] — salinan teks**; 17 catatan pertanyaan 10). Pilot diperkirakan berjalan sesudah PP berlaku **[Asumsi]**; DPIA yang ditulis menjelang tahap publik tidak lagi "sebelum pemrosesan". Dokumen ini dirancang seolah pengecualian rumah tangga tidak berlaku (18 §2.1), dan dengan 10-C A tidak ada jawaban L-01 yang dapat membebaskan pilot sebelum data nyata |
| 3. Biaya kecil | DPIA singkat disusun dari 10.4 §5. Catatan prosesor tidak berarti merundingkan perjanjian: ketentuan baku setiap prosesor diunduh, dicocokkan dengan sebelas butir PP 33/2026 Pasal 14 (salinan teks), dan disimpan bertanggal. Kebijakan privasi versi pilot disusun dari dokumen ini dalam bahasa Indonesia (Permenkominfo 5/2020 Pasal 9 ayat (2)) |
| 4. Sejalan dengan S4-2 yang ditunda | Penyedia hosting, cadangan, dan email baru dipilih menjelang deploy (S4-2, "nanti aja pas mau deploy"), yaitu sebelum deploy pertama yang memegang data nyata (keputusan 45, Q-220: "Sebelum data nyata (Saran)"; diterima atas saran lead). Catatan prosesor untuk ketiganya dibuat saat penyedianya dipilih, sebelum data nyata. Untuk Claude pada R1 tidak ada data rumah tangga yang dikirim (SEC-121), jadi ketentuan bakunya dicatat tanpa diperlakukan sebagai prosesor data pribadi **[Usulan]**; kedudukannya L-32 |
| 5. Kebijakan privasi juga memenuhi asas 2 | Istri pemilik produk dapat membaca secara tertulis apa yang diproses, oleh siapa, dan berapa lama, di samping pemberitahuan pengelola (SEC-34) |

**Kedudukan: diterima atas saran lead** (Q-208). Karena 17 catatan pertanyaan 10 pilihan (ii) berbunyi "advokat dan DPIA baru sebelum tahap publik" sedangkan S4-4 (iii) A hanya menyebut advokat, lead menyajikannya sebagai Q-208 di akhir checkpoint S4; pemilik produk menjawab "pasang semua" (`_progress.md`, "Hal terakhir checkpoint S4"). Pemicu bagi L-01, L-14, L-24 bila ada pengguna di luar rumah tangga pilot (§4.1) diterima dengan cara yang sama (Q-209).

### 3.2 Risiko sisa yang diterima pemilik produk: ambang yang diakhiri berdua (10-E A)

**Risiko.** Karena ambang hanya berubah atau berakhir dengan persetujuan semua pihak, satu pihak dapat mempertahankan ambang atas pihak lain tanpa batas waktu dengan tidak menyetujui perubahan atau pengakhirannya. Ini berlaku bagi setiap pihak dengan cara yang sama.

**Kedudukannya terhadap asas pengawasan.** Menurut rumusan asas pengawasan (06 PG-14; ADR-027 butir (b), (c); 10.3 §2.4.1), melepas kesepakatan yang juga membatasi diri sendiri boleh dilakukan sendirian. Butir (e) asas yang sama (ADR-027, diterima lewat D-113) menyatakan asas itu batas bawah, bukan izin: keputusan pemilik produk yang menuntut semua pihak tetap berlaku, dan menyebut ambang (S4-9 (i) A) secara langsung. Keputusan 10-E A termasuk butir (e), jadi tidak bertentangan dengan asas dan tidak memerlukan pengecualian baru. Risikonya dicatat sebagai **risiko sisa yang diterima pemilik produk**, sama seperti W-18 untuk kehadiran.

**Mitigasi yang berlaku**, semuanya sudah ada di rancangan dan dapat diuji:

| # | Mitigasi | Rujukan | Cara menguji |
|---|---|---|---|
| 1 | Tidak ada yang berada di bawah ambang tanpa persetujuannya sendiri: ambang lahir hanya dari kesepakatan semua pihak, dan menyepakatinya terpisah dari menerima undangan | 06 K02; FR-FAM-004/AC-6 | Terima undangan tanpa menyepakati ambang: keanggotaan aktif, tidak ada ambang yang berlaku atasnya |
| 2 | Sebelum menyetujui, layar kesepakatan ambang menyatakan bahwa ambang hanya dapat diubah atau diakhiri bila semua pihak setuju | SP-41 syarat kedua (10.3 §2.4); 08 | Tinjauan kata layar: kalimat itu tampil sebelum tombol setuju |
| 3 | Ambang menandai, tidak menghalangi: pencatatan dan pembayaran tidak pernah ditolak karena ambang | 06 INV-A11; ADR-009 | Uji INV-A11 dengan ambang aktif |
| 4 | Ambang mengikat setiap pihak sama, termasuk pihak yang ingin mempertahankannya | 06 PG-01; SEC-89 | Uji simetri berpermutasi |
| 5 | Tanda ambang jatuh pada pembayaran rumah tangga, bukan pada uang pribadi satu orang | 06 PG-14 (alasan yang dikutip 10.3 §2.4.1); 07 §3.8 | Uji milik 07: entri yang hanya menyentuh buku pribadi seseorang tidak dinilai terhadap ambang |
| 6 | Usulan mengubah atau mengakhiri selalu dapat diajukan; batas pengingat berlaku, dan tidak ada hitungan penolakan per orang, sehingga aplikasi tidak menekan pihak mana pun ke arah mana pun | 06 §10; AT-4; SEC-07 | Uji AT-4; pengingat melewati batas tidak terkirim |
| 7 | Keluar dari rumah tangga tetap pemutus terakhir bagi siapa pun atas setiap kesepakatan rumah tangga; halaman Bantuan terbuka tanpa jejak | 06 F06; SEC-134 | FR-FAM-012; SEC-134 |

Dokumen ini tidak membuka kembali keputusan itu.

## 4. Pilot lawan tahap publik

| Kelompok | Wajib sebelum pilot memproses data nyata | Wajib sebelum tahap publik | Dianjurkan di pilot |
|---|---|---|---|
| Akses dan isolasi | SEC-01 sampai SEC-09 (SEC-08 RLS sejak pilot: ADR-016 butir 3, D-99), SEC-10 (fungsi), SEC-38 (a), (b), (e), (g) | SEC-10 (pemberian peran), SEC-38 sisanya bila keanggotaan ganda dibangun | — |
| Jejak audit | SEC-39 (a), (b), (c) pilot | SEC-39 (c) publik: sebelum GR-4 (§4.2); (d) | — |
| Autentikasi dan sesi | SEC-11 sampai SEC-17, SEC-19 sampai SEC-24, SEC-26, SEC-27 | SEC-18 (faktor kedua); SEC-25 (kode pemulihan): sebelum GR-4 (§4.2) | — |
| Undangan | SEC-28 sampai SEC-32 | — | — |
| Operator | SEC-33 sampai SEC-35 | SEC-36 dan teks umum SEC-34: sebelum GR-4 (§4.2) | SEC-37 |
| Enkripsi, log, perangkat | SEC-40 sampai SEC-44, SEC-46 sampai SEC-49 | SEC-45; tinjauan ulang SQLCipher; ADR E2EE (10-D A), yang juga menilai sebutan anggota (SEC-155) | SEC-45 bila tersedia |
| Notifikasi dan pintasan | SEC-50 sampai SEC-56, SEC-58, SEC-59; SEC-57 bila Tier 2 | — | — |
| Kehadiran, chat, foto | SEC-60, SEC-61, SEC-63 sampai SEC-67, SEC-69; SEC-62 dan SEC-68 bila Tier 2; SEC-153 (a) sampai (c), (e); SEC-154 (c) | SEC-153 (d) dan SEC-154 sisanya: sebelum build pertama ke TestFlight eksternal atau App Store (pilot memakai TestFlight internal, ADR-022; 10.4 §2.1) | — |
| Anak | SEC-80 sampai SEC-82, SEC-85 | SEC-83 (sebelum kiriman toko pertama; deskriptor Apple paling lambat 12 Januari 2027), SEC-84 (sebelum GR-4, §4.2) | — |
| Penyalahgunaan | SEC-86 sampai SEC-93 | — | — |
| Siklus data | SEC-94 sampai SEC-110 (a) sampai (c); SEC-105 di aplikasi; SEC-107 swalayan sejak GR-2 | SEC-105 (tautan web); SEC-110 (d) sebelum migrasi | — |
| `AI` (bila Tier 2) | SEC-120, SEC-121, SEC-123 sampai SEC-136, SEC-138 | SEC-122 (R2), SEC-137 (Tier 3) | — |
| Toko | SEC-139 (versi pilot), SEC-140 (di aplikasi), SEC-143 bila Play dipakai untuk pengujian, SEC-144 sampai SEC-146; draf SEC-141, SEC-142 | SEC-139 (versi publik: sebelum GR-4, §4.2), SEC-140 (tautan web), SEC-141, SEC-142; SEC-156 sebelum Plus dibangun (FS-3) | — |
| Regulasi dan operasi | SEC-148, SEC-149; SEC-150 (a), (b) (Q-208, diterima atas saran lead); verifikasi SAPA 129 (SEC-138). **Tidak ada butir advokat** (10-C A) | SEC-147 (lebih awal bila jawaban L-14 sebelum GR-4 memintanya, §4.2); semua butir advokat, termasuk L-01, L-14, L-24 (sebelum GR-4 menurut Q-209); tinjauan advokat atas kebijakan privasi (SEC-139), DPIA (SEC-150), dan catatan prosesor; SEC-157 uji penetrasi | — |
| Perilaku | SEC-151, SEC-152 | — | — |
| Sebutan anggota (F-FAM-021, Tier 2) | SEC-155 (a) sampai (i), bila fiturnya dibangun | SEC-155 (j) | — |

Ponytail berlaku atas ruang lingkup bangun, tidak atas kedalaman keamanan Tier 1 (`CLAUDE.md`): tidak ada persyaratan "Pilot" di atas yang dapat dipotong untuk mempercepat pilot tanpa keputusan baru yang dicatat.

### 4.1 Daftar sebelum pilot memproses data nyata

Dinyatakan ulang sesudah checkpoint S4. Setiap butir dicentang dengan tanggal di runbook pilot sebelum data nyata pertama dicatat.

| # | Butir | Rujukan |
|---|---|---|
| 1 | Setiap persyaratan berkolom "Pilot" di tabel atas dibangun dan lulus ujinya di server, iPhone 13, dan iPad mini 7, termasuk perlindungan akun yang kini Tier 1 (S4-6 (iii) A) dan F-COM-014 (S4-7 A). Bagian khusus Android (§4.3) tidak termasuk (keputusan 42) | §4; §4.3; 05 |
| 2 | Istri pemilik produk membaca sendiri pemberitahuan pengelola versi pilot di aplikasi; pengakuannya tercatat bertanggal dan berversi | SEC-34; FR-FAM-019; 10-F A; keputusan 25 |
| 3 | Catatan persetujuan uji dibaca dan diakui kedua penguji sebelum build TestFlight pertama dipasang | SEC-109 (b) |
| 4 | DPIA singkat selesai, dan catatan ketentuan setiap prosesor tersimpan bertanggal (Q-208, diterima atas saran lead). Catatan FCM menyusul sebelum Android (§4.3) | SEC-150; §3.1 |
| 5 | Kebijakan privasi versi pilot dalam bahasa Indonesia ada di aplikasi (Q-208, diterima atas saran lead) | SEC-139 |
| 6 | Rencana respons insiden tertulis dan dilatih sekali; catatan kegiatan pemrosesan ada | SEC-148, SEC-149 |
| 7 | Penyedia hosting dan cadangan dipilih sebelum deploy pertama yang memegang data nyata (S4-2; keputusan 45, Q-220), dengan pihak kontrak dan lokasi pusat data tercatat dan pernyataan tertulis bahwa isi VPS dan cadangan tetap di Indonesia. Sampai itu hanya ada **instans sementara berisi data uji** (*server uji sementara* di glosarium; keputusan 45; 14 §7.8), dengan syarat **[Usulan]**: (a) SEC-40 sampai SEC-42 berlaku, tanpa cadangan; (b) tidak pernah menerima data nyata: setiap akun memakai alamat email domain uji, seperti SEC-41 (a), dan pasangan pilot tidak membuat akun dengan email pribadinya di sana; (c) kunci penandatangan token, kata sandi basis data, dan kunci cadangan server produksi dibuat baru, tidak disalin dari instans sementara; (d) sesudah server produksi berjalan, instans sementara beserta disk dan snapshot-nya dihapus di penyedianya, dan tanggalnya dicatat di runbook. Uji: daftar akun instans sementara hanya memuat domain uji; fingerprint setiap kunci produksi berbeda dari kunci instans sementara; konsol penyedia tidak lagi memuat instans itu | §5 alinea akhir; ADR-020; RISK-04; keputusan 10; keputusan 45 |
| 8 | Pemulihan cadangan diuji sekali; uji SEC-96 dan SEC-97 lulus | SEC-43, SEC-96, SEC-97 |
| 9 | SAPA 129 terverifikasi dan dicatat bertanggal, atau tidak ditampilkan di halaman Bantuan | SEC-138; W-33 |
| 10 | Pemilik produk ditanya apakah iPad mini 7 dipegang orang lain (anak, kerabat) dan apakah sidik jari orang lain terdaftar di Touch ID-nya (iPad mini 7 memakai Touch ID di tombol atas, **[Terverifikasi]** 18 §2.12.9); jawabannya dicatat di runbook pilot. Menyalakan kunci aplikasi tetap pilihannya (SEC-19) | SEC-59 (e) |
| 11 | Tinjauan hanya-temuan `security-privacy` sebelum GR-2 dijalankan dan temuannya tercatat (keputusan 49: "Setuju (Saran)", diterima atas saran lead; dicatat lead di `CLAUDE.md` "Build team"). Bahannya daftar ini, uji SEC berkolom "Pilot", dan §0 uji iOS tanpa Mac | Keputusan 49 |

**Bukan syarat sebelum data nyata** (10-C A): jawaban advokat atas L-01, L-14, L-24, dan pendaftaran PSE. Akibatnya bila pilot ternyata wajib mendaftar PSE, yaitu kemungkinan pemutusan akses, adalah risiko yang diterima pemilik produk (W-32; 10.4 §3.2). **Pemicu (Q-209, diterima atas saran lead):** risiko itu diterima untuk pilot satu rumah tangga milik pemilik produk. Bila sebelum tahap publik ada pengguna di luar rumah tangga itu (penguji lain, grup TestFlight eksternal, Google Play, App Store), L-01, L-14, dan L-24 dijawab advokat lebih dulu. Pemilik produk menerimanya bersama Q-208 dengan "pasang semua" (`_progress.md`, "Hal terakhir checkpoint S4"). **Wawancara dan uji prototipe statis dengan keluarga lain bukan pengguna di luar rumah tangga itu** selama empat syarat §7.4 terpenuhi (keputusan 47, Q-215: "Bukan, dengan 4 syarat (Saran)"; diterima atas saran lead; `_progress.md`, "Penutupan pertanyaan terbuka sebelum S6 (12 November 2026)"). Begitu satu syarat gugur, misalnya seorang peserta memasang build atau mengisi data, pemicu ini berlaku.

### 4.2 Butir "sebelum tahap publik" yang jatuh sebelum GR-4 (beta tertutup)

Menjawab 15 §11 (a). Sejak GR-4 rumah tangga beta (13 §6.4) memakai server yang sama, dan pengelola server bukan anggota mereka. Keputusan 25 dan pemberitahuan versi pilot hanya mengikat rumah tangga pilot; bagi rumah tangga beta tidak ada dasar yang setara. **[Usulan]** `security-privacy`.

| Butir | Mengapa sebelum GR-4 |
|---|---|
| SEC-36 dan SEC-39 (c) publik | Tanpa keduanya, pengelola server memegang akses basis data tetap atas keluarga yang tidak pernah menerima keputusan 25. Log akses darurat harus keluar dari VPS yang dikuasai operator yang sama (SEC-36) |
| SEC-34 teks umum, SEC-139 versi publik | Teks versi pilot menyebut pengelola sebagai anggota rumah tangga; bagi rumah tangga beta kalimat itu keliru (SEC-35). Versi dipilih per rumah tangga (10.1 §11.3 alinea akhir) |
| SEC-25 | Pemulihan dan ganti email tanpa operator yang mengenal penggunanya secara pribadi; TM-11 (pasangan yang menguasai email) tidak lagi dapat diselesaikan dengan berbicara langsung |
| SEC-84 | Pengguna di luar pilot; hanya dokumen, biayanya kecil |
| SEC-147 | Q-209 mewajibkan jawaban L-14 sebelum ada pengguna di luar rumah tangga pilot; bila jawabannya mewajibkan pendaftaran, pendaftaran juga sebelum GR-4 |
| Sudah jatuh di GR-4 menurut bunyinya sendiri | SEC-153 (d) dan SEC-154 sisanya (TestFlight eksternal), SEC-142 (jalur Play selain internal testing), SEC-143 |
| Tinjauan hanya-temuan `security-privacy` sebelum GR-4 | Keputusan 49 ("Setuju (Saran)", diterima atas saran lead; `CLAUDE.md` "Build team"); bahannya tabel ini dan §4.3 bila Android dipakai rumah tangga beta |
| Daftar §4.3, bila Android belum terpasang di dalam pilot | Bawaan keputusan 42 **[Usulan]** lead: Android paling lambat sebelum GR-4. Rumah tangga beta dapat memakai Android, jadi setiap persyaratan "Pilot" berlaku bagi aplikasi Android sebelum GR-4 |

**Tetap sebelum GR-5**, dengan alasannya: SEC-18 faktor kedua (pada skala beta SEC-12, SEC-16, SEC-17 menutup pengambilalihan akun oleh orang luar, yang tumbuh dengan jumlah pengguna; dianjurkan sebelum GR-4 bila biayanya kecil); SEC-45 (risikonya berasal dari penyedia, tidak berubah dengan rumah tangga beta); tinjauan ulang SQLCipher; ADR E2EE (10-D A menetapkan "sebelum tahap publik"; selama beta akses operator dibatasi SEC-36); SEC-10 pemberian peran dan SEC-38 sisanya (Tier 3); SEC-105 dan SEC-140 tautan web (Play produksi); SEC-157.

### 4.3 Sebelum aplikasi Android dipasang (keputusan 42)

Wajib sebelum APK pertama dipasang di perangkat yang memuat data nyata (Galaxy A55 pemilik produk; sesudah GR-2, paling lambat sebelum GR-4). Selain baris di bawah, **setiap** persyaratan berkolom "Pilot" berlaku bagi aplikasi Android dan diuji di Galaxy A55 sebelum pemasangan itu (§0). Bagian iOS dari baris yang sama sudah wajib sebelum GR-2. **[Usulan]** `security-privacy`.

| SEC | Bagian khusus Android | Uji di Galaxy A55 |
|---|---|---|
| SEC-15 | Token di Android Keystore, tidak ikut cadangan atau transfer | Pulihkan cadangan ke perangkat lain: tidak ada sesi |
| SEC-19 (a), (b) | Kode dibungkus kunci Keystore "hanya perangkat ini"; `setInvalidatedByBiometricEnrollment(true)` | Daftarkan sidik jari baru: biometrik gugur |
| SEC-20 | Pratinjau di recents tertutup tanpa memblokir tangkapan layar: `setRecentsScreenshotEnabled(false)`, API 33 ke atas (**[Terverifikasi]**, 18 §2.12.8). Di API 29 sampai 32 tidak ditemukan API resmi lain (**[Perlu validasi]**, 18 §2.12.8); `FLAG_SECURE` memblokir tangkapan layar (**[Terverifikasi]**) dan diduga juga menutup recents (**[Perlu validasi]**), bertentangan dengan SEC-20; tidak menyentuh Galaxy A55, diputuskan sebelum tahap publik (§7.6) | Buka recents: tidak ada isi; tangkapan layar tetap bisa |
| SEC-42 | Akun layanan FCM di penyimpanan rahasia; **kunci tanda tangan APK** dibuat dan disimpan sebelum APK pertama, karena tidak dapat diputar (14 §7.5). Pilihan Play App Signing tetap sebelum unggahan pertama ke Play | Pemindaian rahasia; lokasi kunci tercatat di runbook rilis |
| SEC-46 (a) sampai (c) | Penyimpanan CE; berkas `CARE` lokal di CE; basis data, foto, dan antrean dikecualikan dari cadangan Google dan `<device-transfer>` | Cadangan Google dan transfer: tanpa basis data aplikasi |
| SEC-49 | Lalu lintas hanya ke server dan FCM | Rekam lalu lintas |
| SEC-50, SEC-52 | Payload FCM hanya dua medan; kanal notifikasi Android per kelas dengan nama umum | Tangkap payload FCM setiap N-NN; setelan notifikasi sistem |
| SEC-54 | App shortcuts statis; layar luar perangkat lipat di emulator | Tekan lama ikon: dua label tetap |
| SEC-58 | Token FCM dan Firebase installation ID dihapus saat keluar akun | Uji kirim sesudah keluar akun ditolak |
| SEC-67, SEC-146 | Photo Picker; manifest tanpa izin galeri penuh, lokasi, kontak, mikrofon | Daftar izin di manifest |
| SEC-150 (b) | Catatan ketentuan FCM (Google) sebagai prosesor, bertanggal, sebelum token FCM pertama | Catatan ada |

## 5. Persyaratan untuk dokumen 09 (`architect`)

Wajib dipenuhi rancangan arsitektur. Nomor dengan judul ringkas; isi dan cara mengujinya di berkas masing-masing.

| Kelompok | Persyaratan |
|---|---|
| Otorisasi dan isolasi | SEC-01 otorisasi server per perintah dan per bacaan · SEC-02 `household_id` di setiap kunci unik dan foreign key · SEC-03 pengenal tanpa waktu · SEC-04 urutan sinkron tanpa celah yang membocorkan entri pribadi · SEC-05 "tidak ada" sama dengan "tidak boleh" · SEC-06 satu jalur baca untuk layar, laporan, ekspor, `AI` · SEC-07 tanpa hitungan per orang · SEC-08 RLS sejak pilot · SEC-09 peran basis data bukan superuser · SEC-10 `full_adult_members` satu fungsi · SEC-38 `household_id` sebagai pemilih yang diperiksa, tanpa data lintas rumah tangga · SEC-39 jejak hanya-menambah dan jejak keamanan per akun |
| Akun | SEC-11 kata sandi dan Argon2id (parameter 09 diterima) · SEC-12 batas laju · SEC-13 autentikasi ulang · SEC-14 biometrik lokal saja · SEC-15 sesi dan rotasi token di Keychain/Keystore · SEC-16 daftar sesi pribadi · SEC-17 pemberitahuan masuk kepada pemilik saja · SEC-19 kunci aplikasi: kode 6 digit di perangkat, biometrik gugur saat pendaftaran berubah, salah 5 kali ke kredensial akun tanpa hapus data · SEC-20 pengalih aplikasi tertutup · SEC-21 hapus jarak jauh saat sesi dicabut · SEC-22 tidak ada pencabutan sesi orang lain · SEC-23 F06 mencabut semua sesi · SEC-24 pemulihan lewat email · SEC-27 tanpa pemulihan darurat oleh anggota lain |
| Undangan | SEC-28 token · SEC-30 tanggapan seragam · SEC-31 F04 dengan email · SEC-32 syarat sebelum aktif |
| Operator | SEC-33 tanpa API admin · SEC-34 pemberitahuan berversi dan pengakuan · SEC-36 akses darurat dengan log hanya-menambah di luar VPS (sebelum GR-4) |
| Data dan enkripsi | SEC-40 TLS · SEC-41 pengerasan VPS dan pemisahan staging (sebelum GR-2) · SEC-42 rahasia · SEC-43 cadangan terenkripsi, jendela 30 hari · SEC-44 log · SEC-45 enkripsi disk · SEC-46 kelas proteksi berkas dan pengecualian cadangan · SEC-47 perangkat hanya menyimpan yang terbaca · SEC-48 pembuangan saat akses menyempit · SEC-49 tanpa SDK pihak ketiga |
| Notifikasi | SEC-50 payload · SEC-51 teks per kelas · SEC-52 kanal per kelas · SEC-53 lencana · SEC-54 layar luar, tanpa widget, pintasan statis · SEC-55 satu notifikasi sesudah jam tenang · SEC-56 N-26 `ringkasan` · SEC-57 lipat N-28 · SEC-58 token push · SEC-59 tanpa indeks Spotlight, saran Siri, atau judul jendela berisi data |
| Chat dan kehadiran | SEC-60 kehadiran dikirim server dengan ketelitian menit, tidak lebih halus daripada yang ditampilkan · SEC-61 kehadiran hanya di chat · SEC-63 chat sebagai kelas peka · SEC-64 waktu server · SEC-65 penyimpanan foto · SEC-66 encoder Go dengan empat syarat · SEC-67 pemilih foto sistem · SEC-69 foto tidak ke `AI` · SEC-153 sembunyikan untuk saya · SEC-154 lapor; hasil 10-A A, 10-D A, 10-G A |
| Sebutan anggota (Tier 2) | SEC-155 sebutan hanya dikirim ke sesi pembuatnya, divalidasi server, tidak pernah di payload push, muatan `AI`, log, atau ekspor orang lain, dihapus pada F07 salah satu pihak dan hapus akun |
| Siklus data | SEC-85 skema tanggungan minimum · SEC-86 uji "tidak pernah terlihat" · SEC-89 uji simetri berpermutasi · SEC-92 kecocokan kategori · SEC-93 tanpa agregasi sisi keluarga · SEC-94 penghapusan 90 hari · SEC-96 data terhapus paling lama 38 hari di cadangan dan salinannya · SEC-97 pemulihan tanpa menghidupkan kembali · SEC-98 perangkat · SEC-100 rumah tangga kosong · SEC-101 sampai SEC-104 ekspor dengan satu jendela · SEC-105 hapus akun · SEC-106 perintah operasi anggota wafat · SEC-108 uji hash inti · SEC-109 catatan persetujuan · SEC-110 tanpa pengalihan kepemilikan |
| `AI` | SEC-120, SEC-121 R1 · SEC-122 kunci R2 · SEC-123 yang tidak pernah dikirim · SEC-124 keluaran tidak membuka data orang lain · SEC-125 persetujuan tetap berlaku · SEC-127 batas per pembangunan dan bulanan · SEC-128 saklar mati · SEC-129 tanpa perbandingan · SEC-130 daftar terlarang otomatis · SEC-133 sampai SEC-136 tanda tekanan berat |
| Perilaku | SEC-151 otorisasi ulang perintah luring; perintah pelindung (1) sampai (7) tidak ditahan jeda urungkan maupun tahan tulis; `exit_closure` ditunda selama tahan tulis dengan tanggal saat perintah diterima atau saat masa tenggang berakhir · SEC-152 penolakan tabrakan tanpa kebocoran |
| Penagihan (FS-3) | SEC-156 status Plus sama bagi setiap anggota berakun, tanpa pembayar dan tanpa toko asal; `store_purchase` hanya modul penagihan; pengenal toko acak per rumah tangga |
| Uji penetrasi | SEC-157 pihak luar sebelum GR-5, terhadap staging |

Tambahan untuk ADR hosting (dari 18 §2.8): catat badan hukum pihak kontrak dan lokasi pusat data penyedia VPS dan tujuan cadangan.

## 6. Persyaratan untuk dokumen 08 (`ux-designer`)

| Kelompok | Persyaratan |
|---|---|
| Desain sasaran dewasa | SEC-82: tanpa unsur indikator D Kepmen 142/2026; catatan keputusan desain sasaran dewasa sebagai bukti |
| Keterbukaan | SEC-87 isi "siapa melihat apa" · SEC-88 kata "disepakati sebagai" dan pernyataan sebelum menyimpan · SEC-34, SEC-35 pemberitahuan pengelola server (draf 10.1 §11.3, tanpa kalimat janji, 10-F A) · SEC-62 layar berbagi `CARE` · indeks §3.2 butir 2: layar kesepakatan ambang menyatakan bahwa ambang hanya berubah atau berakhir bila semua setuju |
| Sebutan anggota (Tier 2) | SEC-155: diatur saat mengundang dan di pengaturan anggota; kalimat layar pengaturan (i); tampil hanya di layar pembuat; layar penerima data atau wewenang selalu menampilkan nama akun (d); tidak pernah di teks notifikasi, pintasan, atau lembar bagikan (c) |
| Akun dan perangkat | SEC-13 layar autentikasi ulang · SEC-14 kata layar biometrik · SEC-16 daftar sesi · SEC-19 kode aplikasi, gugurnya biometrik, dan jalan kredensial akun sesudah salah 5 kali · SEC-21 peringatan catatan belum terkirim saat mencabut · SEC-80 pernyataan usia |
| Undangan | SEC-29 teks terbagi tanpa nama · SEC-31 layar F04 dengan email · SEC-32 urutan sebelum aktif |
| Notifikasi dan pintasan | SEC-51 teks per kelas · SEC-52 nama kanal umum · SEC-53 lencana · SEC-54 layar luar perangkat lipat dan pintasan statis · SEC-56 baris N-26 di ringkasan dan layar keanggotaan tanpa menyebut masa tenggang · SEC-59 (b) kalimat tablet di tawaran kunci aplikasi, (d) kalimat perangkat penerima di layar berbagi `CARE` dan "siapa melihat apa" |
| Chat | SEC-60 terakhir aktif menit tepat (10-A A, UX-5 A) · SEC-61 kehadiran hanya di chat · SEC-64 waktu di chat · SEC-68 layar sebelum menghapus status · SEC-153 layar sembunyikan untuk saya, termasuk pernyataan tentang tanda dibaca · SEC-154 tombol lapor (pilot: membuka Bantuan) |
| Keluar dan data | SEC-101 jalan ekspor sesudah keluar dengan satu jendela · pilihan hapus seketika bukan bawaan dan menyebut hilangnya masa batal (10.3 §3.1) · SEC-105 hapus akun · FR-FAM-014/AC-4 batas penghapusan |
| `AI` dan bantuan | SEC-125, SEC-126 layar ikut serta · SEC-133 baris bantuan bagi pemilik status berat · SEC-134 halaman Bantuan menurut 10.4 §1.7 |
| Lain | SEC-90 katalog kategori netral · SEC-93 tanpa tampilan per sisi keluarga |
| S5 | SEC-151: selama tahan tulis, perintah pelindung menampilkan hasil server, bukan keadaan "menunggu" luring · SEC-156: layar status Plus dan tombol pengelolaan sama bagi setiap anggota, tanpa nama pembayar dan tanpa ikon toko asal |

## 7. Yang diminta dari pemilik dokumen lain

Dokumen ini tidak menyunting dokumen lain. Permintaan berikut dikirim lewat lead.

| Kepada | Permintaan |
|---|---|
| `architect` (06) | N-26 menjadi `ringkasan` tanpa notifikasi F07 dan F08 (SEC-56); lipat N-28 60 menit (SEC-57); status `CARE` yang tidak dibagikan di perangkat saja (10-D A; 06 §13.3, §13.5); daftar uji PG-14 diperluas (SP-44); kecocokan kategori tautan pos rutin (SEC-92); perintah "sembunyikan untuk saya" (10-G A) tanpa Peristiwa dan tanpa mengubah tanda air dibaca (SEC-153), dan laporan tanpa Peristiwa (SEC-154); perintah penyempitan tidak ditahan jeda urungkan (SEC-151); ADR untuk pertentangan §8; **ADR rumusan ulang PG-14** (10.3 §2.4.1 butir (a) sampai (e), atau dua pengecualian terbatas bila bunyi PG-14 dipertahankan) dan dua kasus uji baru di daftar uji PG-14 |
| `architect` (09, ADR) | **CS-A-03 (SEC-127):** di 09 berkas 03 §11.1 dan ADR-021, ganti penghitung harian USD 0,33 pada jalur R1 dengan batas **per pekerjaan `ai.pool_build`** (bawaan USD 3 **[Asumsi]**) dan batas bulanan USD 10; batas harian hanya untuk jalur `ai.suggest` (R2); pekerjaan yang terhenti tidak menerbitkan kumpulan setengah jadi. **CS-A-04 (SEC-96):** di 09 berkas 04 §6 dan ADR-020, batas yang dijamin adalah **data terhapus tidak bertahan lebih dari 38 hari di salinan mana pun**; salinan pelindung di VPS cadangan hanya menambah berkas baru (bukan salinan penuh harian) dan dibersihkan di VPS cadangan menurut umur set cadangan, bukan umur salinan; uji SEC-96 ditambahkan. **SEC-66 (ADR-016 butir 12):** encoder standar Go **disetujui dengan empat syarat** SEC-66 (a) sampai (d); 09 berkas 02 :400 tetap **[Perlu validasi]** sampai uji korpus lulus (CS-A-12). **SEC-46:** catat di ADR-024 penyimpangan dari anjuran Apple (`Complete`) dan alasannya; berkas status `CARE` lokal ber-Class A (10-D A). **SEC-19:** kode aplikasi dan aturan salah 5 kali menurut SEC-19 (a) sampai (e); label biometrik di 09 berkas 03 §2 dapat naik ke **[Terverifikasi]** menurut 18 §2.12.2. **SEC-11:** parameter Argon2id diterima. **SEC-38:** `household_id` di `SyncRequest` hanya pemilih yang diperiksa terhadap keanggotaan sesi (sejalan dengan CS-A-06). **SEC-39:** tabel Peristiwa dan `security_event` tanpa `UPDATE`/`DELETE` bagi peran aplikasi; retensi `security_event` 90 hari. **SEC-60:** muatan kehadiran dari server dengan ketelitian yang ditampilkan. **10-G:** ADR-022 menyebut apakah pasangan pilot dipasang sebagai penguji TestFlight internal atau eksternal, karena itu menentukan kapan SEC-153 (d) dan SEC-154 diperlukan |
| `finance-domain` (07) | SEC-92 kecocokan kategori; SEC-93 larangan agregasi per sisi keluarga; 10-E sisi uangnya (kini: ambang diakhiri hanya dengan persetujuan semua, §7.3); jangka 90 hari buku pribadi dan satu jendela ekspor (07 §8.2 butir 1a dan 2a merujuk dokumen 10; SEC-101); SEC-110 (c): pastikan 07 punya aturan mencatat pemegang baru tempat uang bersama yang pemegangnya keluar, tanpa pengalihan otomatis |
| `product-analyst` (05, 17) | **17:** L-32 sampai L-38 **selesai**, tertulis di 17 dengan nomor §7.1 (final); keputusan 10-A sampai 10-G (dengan UX-5) sebagai pertanyaan checkpoint S4, kini dijawab (§3, §7.3); Q-192 selesai (10.1 §6); catat pengecualian hapus seketika sebagai bagian Q-179 yang sudah diterima. **05:** FR-FAM-003/AC-6 dipersempit (SEC-29); FR-FAM-004 menyebut email di F04 (SEC-31) dan pernyataan usia (SEC-80); FR-FAM-012/AC-7 menurut SEC-56; FR-FAM-013 dan FR-FAM-014 menurut satu jendela ekspor SEC-101, SEC-102 (tanpa `CARE`, 10-D A), SEC-96 (38 hari); FR-FAM-017/AC-4 menurut SEC-106; FR-AI-002, FR-AI-003, FR-AI-004 menurut R1 dan SEC-123; FR-AI-007 menurut SEC-127; FR-AI-009 menurut SEC-133 sampai SEC-136; FR-CARE-003 menurut SEC-57; kebutuhan baru: kunci aplikasi (SEC-19), hapus akun (SEC-105), daftar sesi (SEC-16), halaman Bantuan (SEC-134), sembunyikan untuk saya (SEC-153), lapor (SEC-154), jejak keamanan pribadi (SEC-39 (b)), isolasi keanggotaan ganda (SEC-38), tanpa pengalihan kepemilikan (SEC-110) |
| Lead | Glosarium: *sembunyikan untuk saya*, *jalan ekspor sesudah keluar* (definisi kerja **[Usulan]**: SEC-153, SEC-101); checkpoint S4 menyajikan 10-A bersama UX-5 dengan pernyataan bahwa pembulatan mengubah keputusan pemilik produk (sudah dilakukan; hasilnya A, A) |
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
| UX-5, ketelitian "terakhir aktif" | **Selesai**: pemilik produk memilih UX-5 A dan 10-A A (ketelitian menit); pembulatan 15 menit tidak diterapkan | SEC-60 |
| Kunci aplikasi dengan kode 6 digit, biometrik opsional yang gugur saat pendaftaran berubah, batas salah 5 kali (08 §12.3) | **Diterima**; persyaratannya dinyatakan di SEC-19 (a) sampai (e) | Kode ≥ 6 digit, hanya di perangkat, kode mudah ditolak; salah 5 kali berarti kredensial akun, **tanpa menghapus data**; lupa kode sama; mematikan kunci dengan kredensial akun. Angka 6 dan 5 tetap **[Asumsi]** |

### 7.3 Permintaan sesudah checkpoint S4

Dikirim lewat lead. Nomor FR dan BR untuk F-FAM-021 diberikan `product-analyst` (sesudah FR-FAM-024 dan BR-FAM-028).

| Kepada | Permintaan |
|---|---|
| `product-analyst` (05) | (1) FR baru untuk F-FAM-021 dengan AC dari SEC-155 (a) sampai (i), satu AC per butir yang diuji. (2) FR-FAM-004 (menambah anggota): isian sebutan opsional yang menempel pada undangan, tidak masuk teks undangan (SEC-155 (a), (c); SEC-29). (3) FR-FAM-010 "siapa melihat apa": sebutan hanya terlihat pembuatnya, dan pemilih penerima selalu menampilkan nama akun (SEC-155 (d)). (4) FR-FAM-013: bagian "Sebutan yang kamu buat" di ekspor pembuat selama aktif, tidak di jalan sesudah keluar; FR-FAM-014 dan FR-FAM-023 (hapus akun): sebutan ke dua arah dihapus (SEC-155 (f), (g)). (5) FR-FAM-019: kalimat **[Tier 2]** tentang sebutan di pemberitahuan pengelola; tanpa kalimat janji (10-F A). (6) FR-AI-006: daftar medan yang tidak pernah dikirim memuat sebutan (SEC-123). (7) FR kesepakatan ambang: diakhiri hanya dengan persetujuan semua; layar menyatakannya sebelum tombol setuju (§3.2 butir 2); hapus setiap AC bersyarat 10-E B/C. (8) Hapus sisa AC bersyarat "bila 10-D A/B" dan "bila 10-A B" bila ada; FR CARE: status yang tidak dibagikan hanya di perangkat, tanpa syarat |
| `product-analyst` (17) | (1) Q-10 (1), (2), Q-57, Q-93, Q-99, Q-122, Q-174, Q-200, Q-201 ditutup dengan hasil §3. (2) W-32 dicatat sebagai **risiko yang diterima pemilik produk** (10-C A). (3) Risiko sisa baru yang diterima: satu pihak dapat mempertahankan ambang atas pihak lain tanpa batas (10-E A; §3.2), dengan rujukan ke mitigasinya. (4) Catatan pertanyaan 10: advokat sebelum tahap publik (10-C A); DPIA, catatan prosesor, dan kebijakan privasi versi pilot tetap fase bangun (Q-208, **diterima atas saran lead**; selesai). (5) Kolom "Menghalangi" L-01, L-14, L-24: "tahap publik". (6) Pemicu §4.1: pengguna di luar rumah tangga pilot sebelum tahap publik berarti L-01, L-14, L-24 dijawab lebih dulu (Q-209, **diterima atas saran lead**; selesai). (7) Butir advokat baru dari 10.4 §6.4 (usulan nomor L-39). (8) K-34 titik (1) ditutup oleh S4-9 (i) A. (9) W-31: butir "jawaban advokat … (bila S4-4 (iii) B)" tidak berlaku lagi |
| `architect` (06) | (1) Hapus transisi K07 bersyarat (mengakhiri ambang sendirian) dan tutup P-42: ambang berakhir hanya lewat kesepakatan semua pihak (10-E A). (2) §13.3, §13.5: status `CARE` yang tidak dibagikan di perangkat saja, tanpa syarat (10-D A). (3) §16.6: kehadiran tetap P0, menit tepat; tidak ada perubahan. (4) Sebutan anggota: bukan Peristiwa, tidak dipakai penyusun teks notifikasi (§11) dan tidak masuk badan pesan atau kartu objek (§16); F07 salah satu pihak menghapus sebutan ke dua arah; undangan yang dibatalkan atau lewat masa menghapus sebutannya (SEC-155 (b), (c), (g)) |
| `architect` (09, ADR) | (1) Model data dan API sebutan: per (`household_id`, akun pembuat, anggota yang disebut); jalur baca hanya ke sesi pembuat lewat fungsi otorisasi yang sama (SEC-01, SEC-06); validasi server SEC-155 (h); pekerjaan hapus untuk F07, hapus akun, dan undangan kedaluwarsa; tidak di payload push, muatan `AI`, log; uji SEC-155 di 09 berkas 04. (2) **ADR-027**: sudah diselaraskan `architect` (butir (c), (e); K-34 titik (1)); tidak ada permintaan tersisa. (3) **ADR-021**: R1 dan Claude diterima (10-B A, A); batas USD 3 per `ai.pool_build` dan USD 10 per bulan. (4) ADR E2EE sebelum tahap publik mencakup sebutan anggota (SEC-155 (j); §8 butir 13). (5) 09.4 §6 (daftar sebelum data nyata) diselaraskan dengan §4.1: tanpa butir advokat; catatan prosesor dan pernyataan lokasi data dibuat saat S4-2 diputuskan |
| `ux-designer` (08) | (1) Layar sebutan: diatur saat mengundang dan di pengaturan anggota; kalimat SEC-155 (i); tampilan di layar pembuat menggantikan atau mendampingi nama akun, tetapi layar penerima data atau wewenang (F04, pemilih penerima berbagi, pihak permintaan, siapa melihat apa, keluar, ekspor) selalu menampilkan nama akun (SEC-155 (d)). (2) Teks final pemberitahuan pengelola dari 10.1 §11.3 tanpa kalimat janji, dengan dua kalimat **[Tier 2]**. (3) Layar kesepakatan ambang: kalimat bahwa ambang hanya berubah atau berakhir bila semua setuju, tampil sebelum tombol setuju (§3.2 butir 2). (4) UX-5 A: "terakhir aktif" menit tepat; hapus varian pembulatan |
| `finance-domain` (07) | Untuk diketahui: usulan pengakhiran ambang yang disepakati berlaku seketika diterima dengan syarat waktu server menerima entri sebagai ukurannya (10.3 §2.4) |

### 7.4 Tinjauan atas permintaan S5 (13 §12.2, 14 §12, 15 §11) dan temuan `critic` S5

| Hal | Hasil | Syarat yang dapat diuji |
|---|---|---|
| CS5-A-01: tahan tulis lawan perintah pelindung (ADR-028 butir 3; 14 §4.7, §7.4) | **Diselesaikan di SEC-151 dan SEC-148** menurut ketetapan lead. Daftar yang tidak pernah ditahan tahan tulis: SEC-151 (1) sampai (7), ditambah masuk, autentikasi ulang, pembaruan token, dan pemulihan lewat email. Berlaku pada setiap tingkat insiden, tanpa saklar pengecualian | Uji tahan tulis di SEC-151; daftar jenis yang tidak ditahan = jenis bertanda `protective` (09 berkas 03 §1) ditambah kelompok masuk dan token; setiap `protective` bernilai `hold_seconds = 0`, tidak sebaliknya |
| CS5-A-07: SEC-08 | **Dipindah ke kolom Pilot** (ADR-016 butir 3 diterima lewat D-99; 09 berkas 02 §3 butir 4). Tidak ada alasan menundanya | Uji SEC-08 di 10.1 |
| B-04 dan 13 §12.2 butir 1: ukuran validasi | **Dikukuhkan dengan batas.** Ukuran VL-00 sampai VL-10 hanya membaca (a) jawaban pasangan; (b) yang terbaca **kedua** anggota dewasa, lewat layar aplikasi atau jalur baca yang sama (SEC-06), tidak pernah lewat kueri basis data langsung oleh pengelola server, yang di pilot adalah suami; (c) hasil lulus atau gagal pemeriksaan teknis tanpa isi (SEC-44). Buku pribadi, tempat uang pribadi, setelan `AI`, kehadiran, tanda dibaca, dan `CARE` yang tidak dibagikan (yang memang tidak ada di server, 10-D A) tidak pernah menjadi sumber. Akibatnya: **bahan frekuensi** VL-02 (bukan syarat sejak keputusan 46) hanya dari entri yang terbaca keduanya, tanpa pemisahan per orang; gerbang VL-02 (a) dan (b) diatur baris "VL-02 sesudah keputusan 46" di bawah; "fitur Tier 2 tidak dipakai enam minggu" untuk `CARE` dan `AI` hanya dari jawaban masing-masing, untuk modul lain dari objek yang terbaca keduanya. Validasi tidak menambah instrumen, medan, tabel, log, atau analitik; ukuran waktu VL-07 diambil dalam sesi uji yang disepakati, tidak dari build yang dipasang; jawaban VL-08 tidak disimpan di aplikasi atau server | Tinjauan 13 §6.1, §6.2 terhadap butir ini; daftar perubahan skema dan log di antara GR-2 dan GR-3 tidak memuat apa pun untuk validasi |
| VL-02 sesudah keputusan 46 (13 §3.5, §6.1; bacaan lead di `_progress.md`, "Penutupan pertanyaan terbuka sebelum S6") | **Sejalan dengan B-04, SEC-44, SEC-49; catatan bug bukan pengumpulan baru**, dengan syarat di kolom kanan **[Usulan]**. (a) Jawaban "nyaman dan aman" termasuk B-04 (a), jawaban pasangan, dengan saluran VL-08. (b) Catatan bug adalah catatan kerja pengembangan tentang **perangkat lunak**, bukan data tentang rumah tangga atau orang: tiga medan yang dibaca gerbang (tanggal ditemukan, golongan, status perbaikan) tidak memuat data pribadi. Sumbernya laporan pasangan, pemeriksaan teknis tanpa isi (B-04 (c), VL-10), dan laporan crash TestFlight yang sudah diberitahukan (SEC-49, SEC-109 (b)); tidak ada sumber baru. Satu hal yang wajib ditambahkan: **bug golongan "data bocor" juga insiden**, bukan hanya bug | (1) Disimpan di pelacak isu pengembangan atau runbook, di luar aplikasi dan server Family OS; tidak ada tabel, medan, log, Peristiwa, atau fitur lapor bug di aplikasi untuk validasi; tombol lapor SEC-154 bukan jalurnya. (2) Uraian perbaikan dan langkah reproduksi ditulis dengan data uji di staging (SEC-41 (a)); tanpa jumlah, nama, kategori, judul, teks pesan, atau `CARE` nyata, dan tanpa tangkapan layar data nyata; isi nyata yang terlanjur tertulis dihapus dari catatan saat ditemukan. (3) Tanpa nama pelapor dan tanpa hitungan laporan per orang (SEC-07). Medan platform boleh ada karena diperlukan untuk memperbaiki; di pilot ia dapat menunjuk pelapor (iPad mini 7 pemilik produk, iPhone 13 istrinya), jadi gerbang hanya membaca tiga medan di atas. (4) Bug yang membuat data pribadi terbaca orang lain, termasuk pasangan, dijalankan juga menurut rencana insiden SEC-148 (pemberitahuan kepada subjek data, tahan tulis bila perlu); catatannya menyebut "data bocor" tanpa menyebut data apa. (5) Tidak ada SDK, analitik, atau pelaporan crash tambahan untuk mendeteksi bug (SEC-49); log server tetap menurut SEC-44. (6) Hasil VL-02 (a) dicatat sebagai tanggal dan "keduanya sudah" atau "belum", tanpa siapa yang menjawab "belum", tidak disimpan di aplikasi atau server. **Ini bukan janji kerahasiaan** (asas 2): di rumah tangga dua penjawab, siapa pun yang menjawab "sudah" tahu bahwa "belum" berasal dari pasangannya. Di pilot pemilik produk adalah salah satu penjawab sekaligus pemutus GR-3, jadi ia selalu tahu sumbernya, dan istrinya tahu bila ia sendiri menjawab "sudah". Catatan tanpa nama hanya mencegah sumber itu tertulis dan terbaca pihak ketiga; ketersamaran sumber hanya ada bila penjawab lebih dari dua, dan tetap berkurang makin sedikit penjawabnya. Karena itu pertanyaan tetap tentang aplikasi, "belum" tidak perlu dijelaskan, dan satu "belum" menahan tanpa ditimpa (13 §6.2), sehingga mengetahui sumbernya tidak memberi alat menekan. Kalimat ini dinyatakan kepada kedua pasangan saat pertanyaan pertama diajukan. Uji: tinjauan catatan bug dan runbook sebelum GR-3 tanpa isi nyata atau nama; daftar perubahan skema dan log GR-2 sampai GR-3 tanpa apa pun untuk validasi; setiap bug "data bocor" punya catatan SEC-148 |
| 13 §12.2 butir 2: Q-215 | **Ketentuan: wawancara dan uji prototipe statis dengan keluarga lain bukan "pengguna di luar rumah tangga pilot"** (Q-209, §4.1) selama empat syarat berikut terpenuhi. **Keputusan 47**, Q-215: "Bukan, dengan 4 syarat (Saran)"; diterima atas saran lead (`_progress.md`, "Penutupan pertanyaan terbuka sebelum S6 (12 November 2026)"). Alasannya: tidak ada layanan yang memproses data peserta. (1) Tidak ada build terpasang, akun, atau data peserta yang masuk server. (2) Prototipe tanpa data nyata dan, bila dibuka daring, tanpa analitik atau rekaman sesi. (3) Catatan wawancara tanpa nama di samping jawaban dan tanpa angka keuangan rumah tangga peserta, disimpan di luar repositori dan server, dan dihapus sesudah dirangkum. (4) Peserta diberi tahu tujuan dan haknya untuk menolak di awal. Begitu satu syarat gugur, misalnya seorang peserta memasang build atau mengisi data, Q-209 berlaku dan L-01, L-14, L-24 dijawab advokat lebih dulu. Apakah catatan wawancara itu sendiri pemrosesan data pribadi oleh pemilik produk di luar pengecualian rumah tangga masuk lingkup L-01 **[Perlu validasi]**; keputusan 47 tidak menjawabnya, dan karena 10-C A advokat ditanya sebelum tahap publik | Daftar syarat (1) sampai (4) dicentang dengan tanggal sebelum wawancara pertama |
| 13 §12.2 butir 3; 14 §12 butir 2; 15 §11 (c): penagihan dan Q-217 | **SEC-156 baru**, wajib sebelum Plus dibangun (FS-3); kelas data di 10.2 §1 baris "Penagihan" | Uji SEC-156 |
| 13 §12.2 butir 4: SEC-107 dan F-FIN-023 | **Tidak memerlukan F-FIN-023.** SEC-107 adalah salinan data milik peminta, dipenuhi pembangkit ekspor FR-FAM-013 yang dibuka selama aktif; F-FIN-023 (ekspor data rumah tangga) tetap Could. Swalayan sejak GR-2 (FR-FAM-013/AC-12, FS-1; ketetapan lead atas CR5-B-09) | Uji SEC-107 |
| 13 §12.2 butir 5: browser (Q-210) | **Dicatat, tidak dirancang.** Bila Q-210 dibuka, dokumen ini menulis model keamanan browser lebih dulu: sesi dan cookie, CSRF, CSP dan XSS pada isi pengguna (chat, sebutan), penyimpanan di browser, komputer bersama di rumah sebagai jalan PL-02, kunci aplikasi, notifikasi browser | — |
| 14 §12 butir 1: uji penetrasi | **Ya, sebelum GR-5** (SEC-157); tidak untuk pilot dan GR-4 | SEC-157 |
| 14 §12 butir 3: staging tanpa data nyata dan tanpa cadangan (ADR-028 butir 1) | **Tidak bertentangan**, dengan syarat: (1) staging tidak pernah menerima data produksi dalam bentuk apa pun, termasuk ekspor dan pemulihan cadangan produksi (uji pulih tetap di VPS cadangan, 09 berkas 04 §6); akun uji memakai alamat email domain uji, bukan alamat orang nyata; (2) staging berjalan sebagai pengguna sistem tersendiri; peran basis datanya tidak punya `CONNECT` ke basis data produksi dan sebaliknya; rahasianya terpisah (SEC-42); (3) konfigurasi keamanannya sama dengan produksi (TLS, SEC-12, SEC-44), tanpa titik API atau saklar debug yang tidak ada di produksi, karena staging yang lemah di VPS yang sama adalah jalan masuk ke produksi; (4) aplikasi staging hanya dapat menghubungi staging (ADR-028). Tanpa data nyata, SEC-43 dan SEC-96 tidak berlaku bagi staging. Syarat ini kini persyaratan Pilot di **SEC-41**, diperiksa sebelum GR-2 (CR5-A-01) | Uji SEC-41 (staging), sebelum GR-2 (QG-30) |
| 14 §12 butir 4: tahan tulis di rencana insiden | **Dirujuk di SEC-148**, dengan pengecualian SEC-151 | Latihan SEC-148 |
| 15 §11 (a): butir sebelum GR-4 | **§4.2** (SEC-08 sudah Pilot) | — |
| 15 §11 (b): verifikasi ulang layanan bantuan | **SEC-138**: setiap gerbang dan sedikitnya setiap 6 bulan **[Asumsi]** | SEC-138 |

### 7.5 Permintaan sesudah putaran perbaikan S5

| Kepada | Permintaan |
|---|---|
| `architect` (ADR-028, 14, 09) | (1) ADR-028 butir 3 dan 14 §4.7, §7.4: tahan tulis menolak setiap perintah tulis **kecuali** SEC-151 (1) sampai (7) serta masuk, autentikasi ulang, pembaruan token, dan pemulihan lewat email; tingkat 1 dan 2 sama; `exit_closure` dan pembersihan fisik ditunda menurut SEC-151, dengan tanggal saat perintah diterima atau saat masa tenggang berakhir; baris uji "Tahan tulis" di 14 §4.7 memakai uji SEC-151. (2) 09 berkas 03 §1: daftar "tidak pernah ditahan" berlaku juga bagi tahan tulis lewat tanda `protective`; setiap `protective` bernilai `hold_seconds = 0`, tidak sebaliknya (**selesai** `architect`). (3) 15 :569, :573, :1298 (RISK-44): RLS sejak pilot (SEC-08). (4) 15 RISK-40, RISK-44 dan 13 GR-4: butir §4.2; SEC-36 memerlukan tujuan log hanya-menambah di luar VPS (09 berkas 04 atau ADR). (5) 14 §7.7: pengingat SEC-138 di setiap gerbang dan setiap 6 bulan. (6) ADR-028 butir 1 atau 14 §6: syarat staging SEC-41 (a) sampai (d); uji SEC-41 (staging) masuk QG-30 sebelum GR-2. (7) 14 §8 baris "Pembeli": kelas data menurut SEC-156 (c); `source_store` tidak keluar ke jalur baca anggota; 06 PG-15 (a) "siapa lagi": setiap anggota berakun (SEC-156 (b)). (8) 14 §4.8 dan QG-31 atau QG-32: SEC-157 sebelum GR-5 |
| `product-analyst` (13, 05, 17) | (1) 13 §6.1, §6.2: batas ukuran validasi di §7.4 (B-04), termasuk `CARE` dan `AI` hanya dari jawaban masing-masing. (2) 17 Q-215: hasil dan syarat §7.4; **dijawab** keputusan 47 (§7.7). (3) 17 Q-217: dijawab sebagai persyaratan SEC-156 dengan 06 PG-15; isi dan harga Plus tetap Q-36, Q-216. (4) 05: kriteria penerimaan salinan data milik sendiri selama aktif (SEC-107): **selesai** sebagai FR-FAM-013/AC-12, fase FS-1 (ketetapan lead atas CR5-B-09). (5) 13 §3.6 GR-4 dan §6.4: butir §4.2. (6) 17 L-40: tambahkan lama simpan bukti pembelian dan apakah identitas pembayar data keuangan spesifik (SEC-156 (c), (f)); L-01: catatan wawancara Phase 0 (§7.4 Q-215) |
| `ux-designer` (08) | §6 baris "S5": hasil perintah pelindung selama tahan tulis; layar Plus yang sama bagi setiap anggota |
| `researcher` (18) | **Selesai** (18 §2.12.7, §2.13.2, §5.2.8): dipakai di SEC-157, SEC-156 (d), (e), dan SEC-42 |
| Lead | Glosarium *Tahan tulis*: tambahkan "kecuali perintah pelindung SEC-151"; 16 dan 01 bila menyebut tahan tulis atau butir sebelum tahap publik |

### 7.6 Permintaan sesudah keputusan 42 (platform pilot)

| Kepada | Permintaan |
|---|---|
| `architect` (14, 09, ADR-022) | (1) Gerbang atau daftar rilis "sebelum APK pertama dengan data nyata" memuat §4.3; QG-30 (sebelum GR-2) tidak memuat bagian Android. (2) SEC-59 (c): tanpa Core Spotlight, `NSUserActivity` yang dapat dicari, donasi App Intents berisi data, dan judul scene berisi data; uji di 09 berkas 04. (3) SEC-20 di Android API 29 sampai 32 (§4.3): pilih antara `FLAG_SECURE` hanya di API itu (tangkapan layar terblokir) atau pratinjau recents terbuka di API itu, sebelum tahap publik; usulan dokumen ini **[Usulan]**: `FLAG_SECURE` hanya di API 29 sampai 32, karena pratinjau yang terbuka membuka isi kepada siapa pun yang memegang ponsel. Dasar (18 §2.12.8): `setRecentsScreenshotEnabled` tersedia sejak API 33 dan hanya mengosongkan pratinjau Recents (**[Terverifikasi]**); `FLAG_SECURE` memblokir tangkapan layar (**[Terverifikasi]**); bahwa `FLAG_SECURE` juga mengosongkan pratinjau Recents hanya tersirat (**[Perlu validasi]**); API resmi lain untuk API 29 sampai 32 tidak ditemukan, halaman dibaca sebagian (**[Perlu validasi]**) |
| `ux-designer` (08) | SEC-59 (b), (d): kalimat tablet di tawaran kunci aplikasi bila berjalan di iPad; kalimat perangkat penerima di layar berbagi `CARE` (SEC-62) dan "siapa melihat apa" (SEC-87) |
| `product-analyst` (05) | AC untuk SEC-59 (b), (c) di FR kunci aplikasi dan FR notifikasi atau pintasan yang bersangkutan |
| `researcher` (18) | **Selesai** (18 §2.12.8): dipakai di §4.3 baris SEC-20 dan butir (3) di atas |

### 7.7 Permintaan sesudah penutupan pertanyaan terbuka sebelum S6 (keputusan 43 sampai 50)

Keputusan 44 (irama rilis 60/80 hari, Q-219) tidak menyentuh dokumen ini: tidak ada SEC yang bergantung pada irama rilis atau masa berlaku build TestFlight. Keputusan 48 dan 50 juga tidak. Keputusan 46 (pemakaian harian yang bertahan, VL-02) dinilai di §7.4 baris "VL-02 sesudah keputusan 46". Q-222 tetap terbuka dan tidak diubah di sini.

| Kepada | Permintaan |
|---|---|
| `architect` (14, ADR-022, ADR-024, 09 berkas 04) | (1) 14 §7.8: Q-220 dijawab (keputusan 45); syarat instans sementara (a) sampai (d) di §4.1 butir 7, termasuk kunci produksi dibuat baru dan penghapusan instans dicatat. (2) Cara uji iOS tanpa Mac (§0): pemeriksaan diri kelas proteksi dan tanda pengecualian cadangan hanya di build aplikasi staging, tidak di build produksi; uji yang menyebut isi basis data lokal di simulator Xcode Cloud; tanpa profil atau sertifikat uji di perangkat pilot. (3) ADR-024: pembuktian **[Asumsi]** bahwa Class A merusak antrean saat perangkat dikunci kini tanpa debugger; mohon dinyatakan cara membuktikannya lewat build staging di perangkat. (4) QG-30: butir tinjauan hanya-temuan `security-privacy` sebelum GR-2, dan daftar GR-4 sebelum GR-4 (keputusan 49) |
| `product-analyst` (17, 13) | (1) Q-215 ditutup dengan keputusan 47 dan empat syarat §7.4 bernomor (1) sampai (4); catatan L-01 tetap. (2) Q-220 ditutup dengan keputusan 45; syarat instans sementara merujuk §4.1 butir 7. (3) 13 §3.2 atau §6 (wawancara Phase 0): daftar syarat dicentang dengan tanggal sebelum wawancara pertama. (4) 13 §6.1 VL-02 (b): catatan bug disimpan di luar aplikasi dan server; reproduksi dengan data uji di staging; tanpa hitungan per orang; medan platform tidak dibaca gerbang (§7.4 syarat (1) sampai (3), (5)). (5) 13 §6.1 VL-02 (b) dan §6.2: bug golongan "data bocor" juga dijalankan menurut SEC-148 (§7.4 syarat (4)). (6) 13 §6.1 VL-02 (a): hasilnya dicatat sebagai "keduanya sudah" atau "belum" bertanggal, tanpa siapa yang menjawab "belum", tidak di aplikasi atau server; kebocoran sumber "belum" di rumah tangga dua penjawab dinyatakan terbuka kepada kedua pasangan, tanpa janji kerahasiaan (§7.4 syarat (6)) |
| Lead | Glosarium bila memuat *instans sementara* atau *tinjauan hanya-temuan*: rujuk §4.1 butir 7 dan butir 11 |

## 8. Pertentangan dengan tujuan produk atau arsitektur

Penyelesaiannya dicatat `architect` sebagai ADR; rekomendasi dokumen ini di kolom kanan.

| # | Pertentangan | Rekomendasi |
|---|---|---|
| 1 | Pengenal bercap waktu (UUIDv7 dan sejenisnya) memudahkan indeks, tetapi membuka jam pencatatan yang AT-6 sembunyikan | Pengenal yang keluar ke klien acak (SEC-03); pengenal internal bebas |
| 2 | Kursor sinkron global yang padat itu sederhana, tetapi celahnya menghitung entri pribadi pasangan (SP-40 c) | Kursor per pembaca atau opak (SEC-04) |
| 3 | Status `CARE` yang tidak dibagikan di perangkat saja (10-D A) memindahkan alur I01 sampai I04 sebelum berbagi ke perangkat | Diterima demi brief §5K terhadap operator; Tier 2, jadi tidak menyentuh bangun Tier 1 |
| 4 | Hapus jarak jauh saat sesi dicabut (SEC-21) membuang catatan yang belum terkirim, sedangkan `ENGINEERING-STANDARDS.md` mencegah kehilangan data | Diterima dengan peringatan di layar cabut; perangkat yang hilang sudah kehilangan catatan itu. Salah kode aplikasi **tidak** menghapus data (SEC-19 (c)) |
| 5 | R1 (kumpulan kalimat) berbeda dari brief §5O "asks the AI to revise it" dan dari 05 yang mengirim judul | **Selesai**: R1 di pilot (10-B A); brief §5O kini memuat kalimat R1, K-56 ditutup; revisi langsung dan R2 dinilai sebelum publik |
| 6 | Lipat N-28 (SEC-57) dapat dibaca menyimpang dari brief §1.1 "notified on every change" | Yang dilipat bunyinya, bukan kabarnya; lead menilai |
| 7 | Ekspor orang yang keluar memuat tulisan pasangan (SEC-102) lawan batas data orang lain di PP 33/2026 Pasal 76 ayat (2) | Dimuat di pilot; L-33 sebelum publik |
| 8 | Kehadiran P2 dan pembulatan "terakhir aktif" mengubah keputusan pemilik produk | **Selesai**: ia mempertahankan keputusannya atas kehadiran (10-A (i) A, keputusannya sendiri); ketelitian menit diterima lewat "lanjut" (10-A (ii) A, UX-5 A); risikonya diterima (W-18) |
| 9 | Basis data perangkat dikecualikan dari cadangan awan: pindah ponsel tidak membawa catatan yang belum terkirim | Diterima; data lain disinkron ulang dari server |
| 10 | Apple 1.2 meminta blokir dan penyaringan; brief §4.2 ("member exit") dan 06 INV-A48 melarang mengeluarkan orang dewasa, SEC-135 melarang pemindaian | Padanan tanpa mengeluarkan dan tanpa memindai (10-G A; SEC-153, SEC-154) |
| 11 | Kelas proteksi berkas yang memungkinkan sinkron latar (Class C) lawan anjuran Apple agar data pengguna ber-`Complete` | Class C untuk basis data utama dengan alasan di ADR-024; Class A untuk status `CARE` lokal (SEC-46) |
| 12 | Salinan pelindung cadangan (perlindungan dari VPS utama yang dibobol, TM-05) lawan batas umur data terhapus (SEC-96) | Salinan hanya menambah dan dibersihkan menurut umur set, sehingga keduanya terpenuhi (38 hari) |
| 13 | Sebutan anggota disimpan di server agar ikut ke perangkat lain dan ke ekspor pembuatnya, tetapi di pilot operator, yang juga anggota, dapat membaca sebutan yang diberikan pasangannya untuknya (TM-34) | Simpan di server dan nyatakan di pemberitahuan pengelola serta layar pengaturan sebutan (SEC-155 (i)). Menyimpannya di perangkat saja, seperti status `CARE` yang tidak dibagikan, menutup operator tetapi kehilangan sinkron antarperangkat dan ekspor; untuk data Tier 2 yang biasanya tidak peka, harga itu tidak sebanding. Dinilai ulang dalam ADR E2EE sebelum tahap publik (SEC-155 (j)) |
| 14 | Asas pengawasan (ADR-027 butir (b), (c)) membolehkan melepas kesepakatan yang juga membatasi diri sendiri sendirian; pemilik produk memutuskan ambang diakhiri hanya berdua (10-E A) | Tidak ada pertentangan yang perlu ADR baru: butir (e) ADR-027 menempatkan keputusan pemilik produk di atas batas bawah asas. Risiko sisanya diterima dan dicatat dengan mitigasinya (§3.2); ADR-027 diselaraskan bunyinya (§7.3) |

## 9. Status butir yang diserahkan ke dokumen 10

| Butir | Jawaban |
|---|---|
| SP-61 (Q-174) kehadiran | 10.2 §6; 10-A A dengan UX-5 A (ditutup) |
| SP-62 chat sesudah keluar dan bukti | 10.3 §3.4; SEC-102 (f); L-22, L-33 (bagian advokat SP-62 (a), (c)) |
| SP-63 pemberitahuan pengelola mencakup chat | SEC-34; draf 10.1 §11.3 |
| SP-64 (Q-175) N-26 | SEC-56 |
| SP-65 F04 identitas | SEC-31 |
| SP-66 kata `AI` tanpa tafsir dan ajakan memantau | 10.4 §1.5; SEC-130 |
| SP-67 (Q-176) jam baris peristiwa | SEC-64 |
| SP-68 teks izin `AI` | SEC-125, SEC-126 (SP-68 (a) sudah di FR-AI-005/AC-6) |
| SP-69 N-28 di layar luar perangkat lipat | SEC-54 |
| SP-39 sisa | SEC-55 |
| SP-41 (Q-93, P-15) ambang | 10.3 §2.4; 10-E A (ditutup; risiko sisa yang diterima, §3.2) |
| K-34 titik (2) dan (3); 06 §15.5 R7-6; P-26 terhadap PG-14 | 10.3 §2.4.1; titik (2), (3) ditutup lewat ADR-027 (D-113); titik (1) ditutup S4-9 (i) A |
| SP-42 (Q-122) pemberitahuan istri | 10.1 §11; 10-F A (ditutup) |
| SP-44 kecocokan kategori | SEC-92 |
| SP-56 (d) bagian advokat | L-37 |
| P-T2-4, P-T2-5, P-T2-9 | SEC-123; SEC-129; SEC-133 |
| FR-AI-009 tanda tekanan berat | 10.4 §1.6 |
| FR-AI-006 daftar medan per fitur | 10.4 §1.2 tabel fitur; SEC-123 |
| Q-10 (1), (2) | 10-B A, A; 10-C A (ditutup); DPIA dan catatan prosesor §3.1 |
| Q-57, Q-99 | 10.2 §4; 10-D A (ditutup) |
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
| Apple 1.2 (18 §2.12.4) | 10.4 §2.1; 10-G A (ditutup) |
| F-FAM-021 sebutan anggota per pembaca (checkpoint S4) | 10.4 §6; SEC-155 |
| Permenkominfo 5/2020 pasal kewajiban (18 §6.4 butir 7) | 10.4 §3.4; L-14, L-38 |

## 10. Yang belum dapat ditetapkan

- **Hukum.** Semua butir L, dijawab sebelum tahap publik (10-C A), terutama L-01, L-14, L-24 (pilot berjalan tanpa jawabannya; W-32 diterima), L-04, L-20 (ADR-007 tetap sementara), L-38 (Permenkominfo 5/2020 pada chat tertutup), dan butir baru tentang sebutan anggota (10.4 §6.4). Teks PP 33/2026 dibaca dari salinan teks saja; Permenkominfo 5/2020 dari naskah 2020 tanpa perubahannya.
- **Perilaku platform**: biometrik, kelas proteksi berkas, dan Keychain kini **[Terverifikasi]** (18 §2.12.2). Peringatan konten sensitif iOS (SEC-153 (d)) dan batas TestFlight kini **[Terverifikasi]** (18 §2.12.5, §2.12.6). Masih **[Perlu validasi]**: apakah ekstensi notifikasi dapat membaca berkas sebelum buka kunci pertama (tidak menghalangi, SEC-50); `ContentSafetyManager` di Galaxy A55; bahwa build TestFlight internal bebas Beta App Review (tidak dinyatakan eksplisit Apple); hak akses peran App Store Connect; isi laporan crash TestFlight; apakah `FLAG_SECURE` juga mengosongkan pratinjau Recents dan apakah ada cara resmi lain di API 29 sampai 32 (18 §2.12.8; §4.3).
- **Layanan bantuan**: SAPA 129 belum terverifikasi di situs penyelenggara; jam layanannya tidak diketahui; cakupan 112 terkini (10.4 §1.7).
- **Mutu R1**: apakah kumpulan kalimat cukup bervariasi dan hangat belum diuji; usulan peneliti, uji buta kecil dengan pasangan pilot (18 §5.9 butir 4).
- **Toko**: kategori deklarasi kesehatan Google; penerapan Apple 1.2 pada chat tertutup (padanan 10-G A; W-30); jawaban "Dating and Relationships" (paling lambat 12 Januari 2027).
- **Keputusan 25**: pemilik produk menyatakan "masalah istri sudah tahu gdong" (checkpoint S4; tidak diverifikasi dokumen ini). FR-FAM-019 dan SEC-34 tetap menjadikan pembacaan di aplikasi syarat teknis sebelum data apa pun (§4.1 butir 2).
- **Hosting**: penyedia, pihak kontrak, dan lokasi pusat data dipilih sebelum deploy pertama yang memegang data nyata (S4-2; keputusan 45); sampai itu instans sementara berisi data uji. Keputusan 10 (server di Indonesia) tetap mengikat (§4.1 butir 7).
- **Operator pengganti (Q-222)**: tetap terbuka; akan diputuskan pemilik produk bersama istrinya sebelum GR-2 (`_progress.md`, "Penutupan pertanyaan terbuka sebelum S6").
- **Angka [Asumsi]**: 7 hari (tenggang, undangan), 90 hari (ekspor, buku pribadi, jejak keamanan), 30 hari (jendela pemulihan), 38 hari (umur data terhapus di cadangan), 14 hari, 15 menit, 60 menit, 1 menit kunci aplikasi, 5 menit autentikasi ulang, kode 6 digit, 5 kali salah kode, USD 3 per pembangunan kumpulan, 30 karakter sebutan anggota.

## Riwayat

| Tanggal | Perubahan |
|---|---|
| 12 November 2026 | Draf pertama, tahap S4 |
| 12 November 2026 | Draf kedua, putaran perbaikan fase D. **CS-B-01**: butir advokat dipetakan ke L-14, L-26 sampai L-30 di 17; sisanya L-32 sampai L-37, ditambah L-38 baru (§7.1); semua rujukan diselaraskan. **CS-B-02**: keputusan 10-G (Apple 1.2), SEC-153, SEC-154, 10.4 §2.1. **CS-B-03**: SEC-38 (banyak rumah tangga), SEC-39 (jejak audit), SEC-110 (serah terima kepemilikan), TM-31. **CS-B-04**: teks pemberitahuan pengelola bersyarat 10-D. **CS-B-05**: SEC-102 (b) tanpa `CARE` di 10-D A. **CS-B-06**: satu aturan jalan ekspor (sejak F06 sampai akhir hari ke-90 sesudah F07). **CS-B-07**: hapus seketika dicatat sebagai pengecualian yang sudah diterima (S3-4 (c), Q-179) dengan pagarnya. **CS-B-08**: layanan bantuan menurut 18 §2.12.1. **CS-B-09**: SEC-83 rating usia menurut 18 §2.12.3. **CS-B-13**: SEC-66 encoder Go dengan syarat. **CS-B-14**: pembulatan "terakhir aktif" kini bersyarat pada UX-5 dan dinyatakan sebagai perubahan atas keputusan pemilik. **CS-B-18**: Permenkominfo 5/2020 pasal kewajiban dibaca (10.4 §3.4), SEC-147 diperluas. **CS-A-03**: SEC-127 batas per pembangunan kumpulan. **CS-A-04**: SEC-96 batas 38 hari dan bentuk salinan pelindung. Tinjauan permintaan 08 (§7.2) dan 09 (SEC-11, SEC-19, SEC-46, SEC-66) |
| 12 November 2026 | Susulan fase D (dari 18 §2.11, §2.12.5, §2.12.6 dan 09 indeks §8). 10.4 §3.4: sumber kanonis kini 18 §2.11; Pasal 42 dipisah sebagai kewajiban penyedia komputasi awan (VPS), bukan aplikasi; Pasal 7 ayat (2) (pemutusan akses bila tidak mendaftar) memperberat L-14 (§3.2); Pasal 9 ayat (2) (petunjuk berbahasa Indonesia) dikunci di SEC-139; Pasal 16 dan 34 ditambahkan. SEC-46 merujuk penyimpangan Class C di ADR-024 dengan alasannya. SEC-151: daftar kanonis tujuh jenis perintah pelindung yang tidak pernah ditahan, sejalan dengan 09 berkas 03 §1; P-41 diterima. §7.2: usulan "dicatat … kemudian hanya bagi pencatat" **tidak diterima** (penjaga tanggal mundur). SEC-49 dan SEC-109 (b): laporan crash otomatis TestFlight. SEC-153 (d) menurut 18 §2.12.5. SEC-153 (d) dan SEC-154 wajib sebelum TestFlight eksternal atau App Store; istri pemilik produk sebagai penguji internal dengan peran App Store Connect tersempit (10.4 §2.1). SEC-110 (d): riset yang masih diperlukan |
| 12 November 2026 | Susulan fase D (temuan `product-analyst` atas 17 K-34): 10.3 §2.4.1 menilai titik (2) (penegasan 1, jatah di atas ambang) dan (3) (S3-5 (i), mengakhiri pos rutin sendirian) terhadap PG-14; keduanya dipertahankan; rumusan ulang PG-14 butir (a) sampai (e) dan permintaan ADR kepada `architect`; tanpa pertanyaan checkpoint baru |
| 12 November 2026 | Draf ketiga, sesudah checkpoint S4. §3 diganti hasil 10-A sampai 10-G dengan kedudukan dan kutipan pemilik produk (10-A A, A; 10-B A, A; 10-C A; 10-D A; 10-E A untuk semua tahap; 10-F A; 10-G A); semua kalimat "bila 10-x A/B/C", pilihan, dan bawaan yang belum dijawab dihapus dari indeks dan 10.1 sampai 10.4. §3.1 baru: DPIA, catatan prosesor, dan kebijakan privasi versi pilot tetap fase bangun. §3.2 baru: ambang yang diakhiri berdua sebagai risiko sisa yang diterima pemilik produk, dengan tujuh mitigasi. §4 dan §4.1 baru: daftar sebelum data nyata tanpa butir advokat; W-32 diterima; pemicu L-01, L-14, L-24 bila ada pengguna di luar rumah tangga pilot. SEC-60 menit tepat; SEC-35 tanpa kalimat janji; 10.1 §11.3 tanpa varian; SEC-46 (b), SEC-102 (b) tanpa syarat; SEC-147, SEC-150 (a), (b), SEC-139 disesuaikan; 10.4 §3.2 dinyatakan ulang. **SEC-155 baru** (F-FAM-021, sebutan anggota per pembaca), TM-32 sampai TM-36, 10.4 §6, butir advokat baru (usulan L-39), §8 butir 13 dan 14. Atas permintaan `finance-domain` lewat lead: 10.3 §2.4 mencatat saran B dan bawaan C tidak diambil; cabang S4-9 (ii) B dan 10-E B/C dihapus dari 10.3 §2.4.1; usulan 07 (pengakhiran yang disepakati berlaku seketika bagi entri yang diterima server sesudahnya) dinilai tanpa keberatan, dengan satu syarat. §7.3 permintaan sesudah checkpoint S4. Penyelarasan akhir dari lead: 10-D dan 10-G berkedudukan diputuskan pemilik produk sendiri (17 K-64); SEC-155 (h) tanpa sebutan ganda per pembuat per rumah tangga (sesudah NFC dan case folding; 08 §7.4); baris SEC-60 di §5 dan §6 menyebut menit tepat. Perbaikan `critic`: CRB-2 SEC-155 (h) pemeriksaan kelas karakter, NFC, dan sebutan ganda di server saja (aplikasi hanya wajib isi dan panjang, ES §2 (c)); CRB-3 10-A dipisah, (ii) diterima lewat "lanjut"; CRB-4 DPIA, catatan prosesor, dan kebijakan privasi versi pilot ditandai **[Usulan], menunggu Q-208**, pemicu §4.1 menunggu Q-209; CRB-8 10.3 §2.4.1 diselaraskan dengan ADR-027 yang diterima, kalimat ADR basi dihapus; CRB-14 SEC-102 (g) tanpa sebutan, §2 merujuk SEC-155 (f); CRB-15 sebutan dihapus pada F07. Sesudah "pasang semua" (`_progress.md`, "Hal terakhir checkpoint S4"): Q-208 (DPIA, catatan ketentuan prosesor, kebijakan privasi versi pilot di fase bangun) dan Q-209 (pemicu L-01, L-14, L-24) berkedudukan **diterima atas saran lead** di §2, §3, §3.1, §4, §4.1, 10.4 §3.2, SEC-139, SEC-150; brief §5O memuat kalimat R1, K-56 ditutup (§8 butir 5; 10.4 §1.2) |
| 12 November 2026 | Draf keempat, putaran perbaikan fase D S5. **CS5-A-01**: SEC-151 menyatakan tahan tulis tidak pernah menahan perintah pelindung (1) sampai (7) beserta masuk dan pembaruan token, pada setiap tingkat insiden; `exit_closure` dan pembersihan fisik ditunda, dengan tanggal saat perintah diterima atau saat masa tenggang berakhir; SEC-148 merujuk tahan tulis. **CS5-A-07**: SEC-08 sejak pilot. **B-04**: batas ukuran validasi (§7.4). Permintaan 15 §11: §4.2 baru (SEC-25, SEC-34 teks umum, SEC-36 dengan log di luar VPS, SEC-39 (c) publik, SEC-84, SEC-107, SEC-139 versi publik sebelum GR-4); SEC-138 verifikasi ulang. **SEC-156** baru (Plus, Q-217, PG-15) dan baris kelas data "Penagihan" di 10.2 §1. **SEC-157** baru (uji penetrasi sebelum GR-5). §7.4 tinjauan permintaan 13, 14, 15 (Q-215, SEC-107 tanpa F-FIN-023, browser, staging); §7.5 permintaan. Susulan dari 18 §2.12.7, §2.13.2, §5.2.8: SEC-156 (d) medan `appAccountToken` dan `obfuscatedAccountId`, (e) baru pemberitahuan server ke server (RTDN lewat Pub/Sub; retensi menjadi (f)); SEC-157 kisaran biaya, penyedia CREST, isi kontrak; SEC-42 kunci APK di Play App Signing. Tinjauan ulang `critic` A2: CR5-A-01 syarat staging menjadi persyaratan Pilot di SEC-41, diperiksa sebelum GR-2; CR5-A-06 kisaran biaya SEC-157 per sumber 18 §2.12.7; CR5-A-09 tanggal `exit_closure` di indeks menurut SEC-151 (saat diterima atau akhir masa tenggang); uji SEC-151 memakai tanda `protective`. CR5-B-09 (ketetapan lead): SEC-107 swalayan sejak GR-2 (FR-FAM-013/AC-12, FS-1), keluar dari §4.2 |
| 12 November 2026 | Revisi keputusan 42 (pilot dengan data nyata di iOS dan iPadOS; Android sesudah GR-2, paling lambat GR-4, bawaan lead [Usulan]). §0: "Pilot" per platform. §4.3 baru: bagian khusus Android (SEC-15, SEC-19 (a), (b), SEC-20, SEC-42, SEC-46, SEC-49, SEC-50, SEC-52, SEC-54, SEC-58, SEC-67, SEC-146, SEC-150 (b)) wajib sebelum APK pertama dengan data nyata, bukan sebelum GR-2; SEC-42 dan SEC-58 kolom Tahap. §4.1 butir 1 dan 4 menurut §4.3, butir 10 baru (pemakaian iPad oleh orang lain). §4.2 baris Android sebelum GR-4. **SEC-59 baru** (10.2 §5): tablet yang dipakai bersama; Shared iPad memerlukan MDM (Apple Support, diakses 12 November 2026); kalimat tablet di tawaran kunci aplikasi; tanpa Spotlight, saran Siri, atau judul jendela berisi data. SEC-20 di Android API 29 sampai 32 dicatat sebagai pertentangan untuk tahap publik, dengan dasar 18 §2.12.8. §7.6 permintaan; §10 butir platform. Tidak ada persyaratan sisi iOS yang berubah |
| 12 November 2026 | Penerapan penutupan pertanyaan terbuka sebelum S6 (`_progress.md`, keputusan 43 sampai 50). **Keputusan 47** (Q-215): penilaian **[Usulan]** §7.4 menjadi ketentuan, empat syarat dinomori (1) sampai (4), dirujuk dari pemicu Q-209 di §4.1 dan dari tabel waktu advokat 10.4; catatan L-01 dan daftar centang sebelum wawancara pertama tetap. **Keputusan 45** (Q-220): §3.1 butir 4, §4.1 butir 7 (penyedia sebelum data nyata; syarat instans sementara (a) sampai (d)), §10, SEC-150 (b). **Keputusan 49**: §4.1 butir 11 dan baris §4.2 (tinjauan hanya-temuan sebelum GR-2 dan GR-4); SEC-157 tidak lagi menyebutnya usulan proses. **Keputusan 43**: §0 cara menguji iOS tanpa Mac dan debugger, tanpa melonggarkan SEC; uji SEC-46 lewat pemeriksaan diri build staging, SEC-49 lewat App Privacy Report (Apple Support 102188, diakses 12 November 2026). Keputusan 44 tidak menyentuh SEC mana pun. §2 empat baris; §7.5 butir Q-215; §7.7 baru. **Keputusan 46** (atas permintaan `product-analyst` lewat lead): B-04 di §7.4 kini membatasi hanya bahan frekuensi VL-02; baris baru "VL-02 sesudah keputusan 46" menilai VL-02 (a) dan (b) terhadap B-04, SEC-44, SEC-49: catatan bug bukan pengumpulan baru dengan enam syarat, dan bug "data bocor" juga insiden SEC-148. Atas susulan `architect` lewat lead: §0 merujuk ADR-024 "Catatan keputusan 43" untuk bukti kelas antrean; pemeriksaan arsip CI (14 §4.8; QG-30 butir 13) diterima dengan tiga syarat (penjaga kompilasi tetap utama, string penanda dan alamat staging, menghalangi sebelum unggahan); istilah glosarium *server uji sementara* sama dengan instans sementara §4.1 butir 7. `critic` P-3: §7.4 syarat (6) VL-02 tidak lagi menjanjikan kerahasiaan sumber "belum"; kebocoran sisa di rumah tangga dua penjawab dinyatakan (asas 2). Label keputusan 43, 45, 47, 49 diperiksa: semuanya "diterima atas saran lead". Q-222 tetap terbuka |
