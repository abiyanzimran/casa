# 00 — Glosarium

Satu-satunya definisi setiap istilah bersama di Family OS. Dokumen lain menautkan ke sini dan tidak mendefinisikan ulang.

**Status:** versi pertama (S1). Definisi bertanda *(definisi kerja)* difinalkan pada tahap yang disebut di barisnya.

**Cara baca.** Kolom *Istilah* adalah kata yang dipakai di dokumen berbahasa Indonesia. Kolom *Term (EN)* adalah padanan bahasa Inggris yang kelak dipakai sebagai nama di kode, sesuai `ENGINEERING-STANDARDS.md` §1. Kata yang dilihat pengguna di layar bisa berbeda dari keduanya dan ditetapkan di dokumen 08; istilah akuntansi tidak pernah tampil ke pengguna.

**Kepemilikan.** Bagian 7 (istilah keuangan) ditulis oleh `finance-domain`. Bagian lain ditulis oleh lead. Permintaan istilah baru disampaikan lewat lead.

---

## 1. Produk dan proses perencanaan

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Family OS | Family OS | Nama kerja produk: aplikasi mobile tempat satu keluarga mengatur uang, permintaan, dan urusan rumah tangga, serta saling memperhatikan. |
| Rumah tangga pilot | pilot household | Rumah tangga pemilik produk: suami dan istri sebagai dua pengguna pertama. Sumber kebenaran utama untuk cakupan MVP (brief §1.1). |
| Tahap publik | public stage | Tahap kedua tujuan produk: produk untuk keluarga Indonesia secara umum, setelah pilot menunjukkan pemakaian harian yang bertahan. |
| MVP | MVP | Produk terkecil yang utuh dan dipakai rumah tangga pilot setiap hari. Cakupannya = Tier 1 + Tier 2. |
| Tier 1 — Inti | Tier 1 | Cakupan yang didokumentasikan penuh dengan kriteria penerimaan: `FAM` untuk satu rumah tangga dengan dua akun dewasa, `FIN`, `NTF` untuk permintaan dan persetujuan, serta `REQ` anggaran dan jalan-jalan. Rinciannya di brief §1.1. |
| Tier 2 — Versi ringan | Tier 2 | Cakupan yang didokumentasikan ringkas: versi ringan `SHP`, `MNT`, `CARE`, `CAL`/`NTF`, `PLN`, permintaan tindakan, dan pendamping AI versi ringan. Rinciannya di brief §1.1. |
| Tier 3 — Ditunda | Tier 3 | Segala sesuatu yang tidak disebut di Tier 1 atau Tier 2. Didokumentasikan sebagai garis besar satu halaman per modul. Menunda modul tidak menunda fondasinya (isolasi rumah tangga, satu orang di lebih dari satu rumah tangga). |
| Tahap (S1–S6) | stage | Enam langkah kerja perencanaan di brief §0.3. |
| Checkpoint | checkpoint | Titik **STOP** di akhir tahap: lead menyajikan ringkasan dan menunggu persetujuan eksplisit pemilik produk. |
| Pemilik produk | product owner | Orang yang memutuskan arah produk dan menyetujui setiap checkpoint. Berbeda dari *pemilik rumah tangga* (peran di dalam aplikasi, bagian 3). |
| Label bukti | evidence label | Salah satu dari **[Terverifikasi]**, **[Asumsi]**, **[Usulan]**, **[Perlu validasi]**; arti masing-masing ada di [README](README.md#label-bukti). |
| ADR | ADR (Architecture Decision Record) | Catatan satu keputusan arsitektur: konteks, opsi, keputusan, konsekuensi. Disimpan di `docs/adr/`. |

## 2. Kode modul

| Kode | Modul | Tingkatan di MVP |
|---|---|---|
| `FAM` | Pengelolaan keluarga dan rumah tangga | Tier 1 sebatas satu rumah tangga dengan dua akun dewasa; sisanya Tier 3 |
| `COM` | Komunikasi keluarga (chat umum) | Tier 3. Diskusi yang menempel pada permintaan termasuk `REQ`, bukan `COM` |
| `FIN` | Pengelolaan keuangan keluarga | Tier 1 untuk butir yang disebut brief §1.1 (custody, alokasi pemasukan, pengeluaran, anggaran per kategori dan periode, tagihan berulang, cicilan, utang dan pinjaman, tabungan sederhana, aset dan kekayaan bersih, laporan ringkas); selebihnya Tier 3 |
| `REQ` | Permintaan, izin, dan persetujuan | Tier 1 untuk permintaan anggaran dan jalan-jalan; Tier 2 untuk permintaan tindakan |
| `SHP` | Belanja | Tier 2 untuk daftar belanja dengan perkiraan dan biaya aktual; integrasi belanja online Tier 3 |
| `INV` | Persediaan rumah tangga | Tier 3 |
| `PLN` | Rencana dan kegiatan keluarga | Tier 2 sebatas yang dibutuhkan permintaan jalan-jalan |
| `GOL` | Tujuan dan tabungan | Pencatatan tabungan sederhana masuk `FIN` Tier 1; selebihnya Tier 3 |
| `MNT` | Masalah dan perawatan rumah | Tier 2 (masalah menjadi tugas) |
| `MEAL` | Masak dan rencana makan | Tier 3 |
| `CARE` | Kesejahteraan dan kepedulian | Tier 2 (status suasana hati sukarela, tindak lanjut sakit) |
| `CAL` | Kalender dan rutinitas | Tier 2 (pengingat rutinitas dan tagihan) |
| `TML` | Linimasa dan riwayat | Tier 3 |
| `NTF` | Notifikasi dan umpan aktivitas | Tier 1 untuk permintaan dan persetujuan; Tier 2 untuk pengingat dan jam tenang |
| `AI` | Pendamping keluarga | Tier 2 versi ringan; percakapan nasihat terbuka Tier 3 |

## 3. Keluarga, rumah tangga, dan identitas

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Rumah tangga | household | Satuan isolasi data di Family OS: sekumpulan anggota yang berbagi catatan uang dan urusan rumah. Setiap data bersama milik tepat satu rumah tangga, dan anggota rumah tangga lain tidak dapat membacanya. |
| Keluarga | family | Orang-orang yang berhubungan kerabat, tinggal serumah atau tidak. Satu keluarga dapat terdiri atas beberapa rumah tangga; keluarga bukan satuan isolasi data. |
| Keluarga besar | extended family | Kerabat di luar rumah tangga (misalnya ibu, ayah, atau saudara dari suami). Di MVP mereka muncul sebagai *pihak luar* dalam catatan uang, bukan sebagai anggota. |
| Akun | account (user account) | Identitas login milik satu orang. Satu orang punya satu akun, terlepas dari jumlah rumah tangga yang ia ikuti. Jangan disamakan dengan rekening bank, yang di glosarium ini adalah salah satu jenis *tempat uang* (bagian 7). |
| Anggota | member | Satu akun dalam kaitannya dengan satu rumah tangga. Orang yang ikut dua rumah tangga adalah dua anggota dengan satu akun. |
| Keanggotaan | membership | Ikatan antara satu akun dan satu rumah tangga; membawa peran, izin, tanggung jawab, dan wewenang persetujuan anggota itu di rumah tangga tersebut. |
| Undangan | invitation | Ajakan dari anggota yang berwenang kepada seseorang untuk bergabung ke rumah tangga; baru menjadi keanggotaan setelah diterima. |
| Pihak luar | external party | Orang atau lembaga yang tercatat dalam catatan rumah tangga tetapi bukan anggota rumah tangga itu (misalnya ibu suami yang menerima kiriman bulanan, saudara yang berutang, penjual). Pihak luar dapat pula memegang uang rumah tangga sebagai *titipan* (bagian 7); uang itu tetap terhitung milik rumah tangga. |
| Tanggungan | dependent | Orang yang berada di bawah pengawasan seorang wali, misalnya anak. Dapat tercatat sebagai anggota dengan akun sendiri atau sebagai *profil tanggungan tanpa login* (bagian 8); bentuk mana yang dipakai adalah pertanyaan terbuka di dokumen 17. Tier 3 di MVP. |
| Wali | guardian | Anggota dewasa yang bertanggung jawab atas seorang tanggungan dan memberi persetujuan atas namanya. Tier 3 di MVP. |
| Isolasi rumah tangga | household isolation | Jaminan bahwa anggota satu rumah tangga tidak dapat membaca atau mengubah data rumah tangga lain, termasuk ketika satu orang menjadi anggota keduanya. |
| Data pribadi anggota | personal data (member-private) | Data milik satu anggota yang tidak terlihat oleh anggota lain kecuali ia memilih membagikannya (misalnya dana pribadi, status suasana hati). Lawannya *data bersama*. Batas persisnya ditetapkan di dokumen 10 (S4). |
| Data bersama | shared data | Data rumah tangga yang terlihat oleh anggota sesuai izinnya. |
| Keluar dengan aman | safe exit | Kemampuan anggota dewasa meninggalkan rumah tangga tanpa memerlukan persetujuan anggota lain dan tanpa kehilangan data pribadinya. Rinciannya di dokumen 10 (S4). |

### Lima konsep yang berdiri sendiri

Kelima konsep ini dimodelkan terpisah dan tidak saling menurunkan secara otomatis (brief §4.3). Suami tidak otomatis punya izin lebih dari istri, dan sebaliknya.

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Hubungan | relationship | Bagaimana dua anggota berkerabat: suami, istri, anak, orang tua, mertua, saudara, wali. Hanya keterangan; tidak memberi hak apa pun. |
| Peran | role | Kedudukan anggota di tingkat sistem dalam satu rumah tangga: pemilik, admin, anggota, tanggungan, tamu. Peran tanggungan hanya berlaku bagi *tanggungan* yang punya akun; profil tanggungan tanpa login bukan anggota dan tidak berperan. |
| Izin | permission | Apa yang boleh dilihat atau diubah seorang anggota (misalnya melihat anggaran bersama, mengubah daftar belanja). |
| Tanggung jawab | responsibility | Apa yang diharapkan diurus seorang anggota (misalnya belanja dapur, tagihan listrik). Tidak memberi izin atau wewenang dengan sendirinya. |
| Wewenang persetujuan | approval authority | Permintaan mana yang boleh disetujui seorang anggota, dan sampai batas berapa. |

Nilai peran *(definisi kerja; difinalkan di dokumen 06 dan 10)*:

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Pemilik rumah tangga | household owner | Anggota yang memegang kendali administratif tertinggi atas rumah tangga (misalnya memindahkan kepemilikan). Satu rumah tangga boleh punya lebih dari satu pemilik; apakah kedua pasangan menjadi pemilik setara secara bawaan adalah pertanyaan terbuka di dokumen 17. Berbeda dari *pemilik uang* (bagian 7) dan dari *pemilik produk* (bagian 1). |
| Admin | admin | Anggota yang boleh mengelola keanggotaan dan aturan rumah tangga. Admin tidak pernah mendapat akses ke data pribadi anggota dewasa lain. |
| Anggota (peran) | member (role) | Peran bawaan anggota dewasa biasa. |
| Tamu | guest | Anggota dengan akses terbatas, misalnya asisten rumah tangga atau pengasuh. Tier 3 di MVP. |
| Preset aturan | rule preset | Kumpulan izin dan wewenang persetujuan siap pakai, agar keluarga tidak harus mengatur semuanya di hari pertama. |

## 4. Permintaan dan persetujuan

Definisi di bagian ini adalah *definisi kerja*; state machine dan aturan persetujuannya difinalkan di dokumen 06 (S2). Jumlah uang pada *penyelesaian* (settlement, bagian 7) kanonis di dokumen 07.

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Permintaan | request | Sesuatu yang diajukan satu anggota kepada anggota lain dan menunggu tanggapan. Permintaan berjalan dua arah: tiap pasangan dapat meminta kepada pasangannya. Ada lima jenis di bawah, yang tidak dimodelkan sebagai satu transaksi keuangan. |
| Permintaan izin | permission request | Meminta persetujuan untuk melakukan sesuatu, tanpa uang yang terikat (misalnya anak meminta izin keluar). |
| Permintaan anggaran | budget request | Meminta *otorisasi* (bagian 7) untuk memakai sejumlah uang bagi tujuan tertentu sebelum uang itu dibelanjakan. Kata "anggaran" di sini mengikuti bahasa sehari-hari; *anggaran* dalam arti batas belanja per kategori dan periode didefinisikan di bagian 7. |
| Permintaan jalan-jalan | outing request | Permintaan anggaran untuk satu kegiatan keluar bersama, yang juga memuat tanggal, tempat, rencana pembelian, dan anggaran (brief §1.1). |
| Permintaan penggantian | reimbursement request | Meminta uang pribadi yang sudah terlanjur dipakai untuk keperluan bersama diganti. Arti *penggantian* ada di bagian 7. |
| Permintaan tindakan | action request | Tugas yang diminta satu anggota kepada anggota lain. Yang diminta menerima atau menolak, yang meminta boleh membatalkan, dan tugas yang diterima tetap tertunda, dengan pengingat, sampai yang diminta menandainya selesai. |
| Usulan keputusan | decision proposal | Mengajukan suatu keputusan bersama untuk disepakati (misalnya pembelian besar). |
| Pemohon | requester | Anggota yang mengajukan permintaan. |
| Pemberi persetujuan | approver | Anggota yang berwenang menanggapi permintaan sesuai wewenang persetujuannya. |
| Penerima tugas | assignee | Anggota yang diminta mengerjakan suatu permintaan tindakan atau tugas. |
| Persetujuan | approval | Tanggapan pemberi persetujuan yang mengizinkan permintaan dijalankan. Persetujuan saja tidak mengurangi uang yang ada (brief §7); persetujuan atas permintaan anggaran melahirkan *otorisasi* (bagian 7), dan akibat keuangannya diatur di dokumen 07. |
| Penolakan | rejection | Tanggapan yang mengakhiri permintaan tanpa mengizinkannya, dapat disertai alasan. Apakah alasan wajib *(diputuskan di S2)*. |
| Usul balik | counter-proposal | Tanggapan yang mengajukan syarat berbeda (misalnya jumlah atau tanggal lain) pada permintaan yang sama, untuk ditanggapi pemohon. |
| Diskusi | discussion | Percakapan yang menempel pada satu permintaan, tugas, atau kegiatan tertentu. Bukan chat umum (`COM`). |
| Pembatalan | cancellation | Pemohon menarik permintaannya sebelum selesai. |
| Kedaluwarsa | expiry | Permintaan atau otorisasi berakhir dengan sendirinya karena waktunya lewat, tanpa disetujui dan tanpa ditolak. Aturannya *(diputuskan di S2)*. |
| Ambang permintaan | request threshold | Jumlah uang yang, bila dilampaui suatu pengeluaran tidak rutin, membuat pengeluaran itu diajukan sebagai permintaan lebih dulu. Berlaku sama bagi setiap anggota dewasa. Nilainya dan ada-tidaknya *(diputuskan di S2)*. |
| Keberatan | dispute | Pernyataan seorang anggota bahwa ia tidak sepakat dengan suatu catatan atau tindakan anggota lain; catatannya tetap ada dan keberatannya terlihat oleh keduanya. Penanganannya *(diputuskan di S2)*. |
| Tugas | task | Pekerjaan dengan penerima tugas dan status selesai/belum. Dapat lahir dari permintaan tindakan, masalah rumah, atau rutinitas. |
| Aktivitas | activity | *(definisi kerja; dokumen 06 memutuskan apakah ini menjadi primitif bersama)* Sebutan umum untuk hal yang punya siklus hidup dan dapat membawa diskusi, lampiran, dan catatan uang: permintaan, tugas, rencana, masalah rumah. |
| Jejak audit | audit trail | Riwayat yang tidak dapat diubah tentang siapa melakukan apa dan kapan atas suatu catatan. Berbeda dari *linimasa* (`TML`), album kenangan yang dikurasi dan boleh disunting. |

## 5. Urusan rumah, kepedulian, dan notifikasi

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Daftar belanja | shopping list | Daftar barang yang akan dibeli, dengan perkiraan biaya dan biaya aktual terhadap suatu *anggaran* (bagian 7). |
| Masalah rumah | household issue | Laporan kerusakan atau gangguan di rumah (misalnya atap bocor) yang menjadi tugas dan ditutup dengan foto dan catatan. |
| Rutinitas | routine | Pekerjaan rumah yang berulang dan diingatkan (misalnya menanak nasi). |
| Rencana | plan | Kegiatan keluarga yang dijadwalkan. Di MVP hanya sejauh yang dibutuhkan permintaan jalan-jalan. |
| Status suasana hati | mood status | Pernyataan sukarela seorang anggota tentang apa yang ia rasakan dan, bila mau, alasannya. Pribadi secara bawaan. Setelah anggota memilih membagikannya kepada pasangan, pasangan diberi tahu pada setiap perubahan. |
| Tindak lanjut sakit | illness follow-up | Alur setelah seorang anggota melaporkan sakit: keluarga menanyakan lagi setelah N hari, dan yang sakit menjawab sudah sembuh, masih sakit, atau akan ke klinik. Tanpa bahasa diagnosis. |
| Persetujuan berbagi | sharing consent | Pilihan eksplisit, terperinci, dan dapat dicabut dari pemilik data kepedulian tentang siapa yang boleh melihatnya. Jangan disamakan dengan *persetujuan* atas permintaan (bagian 4). |
| Notifikasi | notification | Pemberitahuan kepada anggota tentang hal yang perlu ia ketahui atau tanggapi. Isi di layar kunci tidak pernah memuat jumlah uang atau rincian kepedulian. |
| Jam tenang | quiet hours | Rentang waktu yang ditetapkan anggota ketika notifikasi tidak mengganggunya. |
| Ringkasan notifikasi | digest | Beberapa notifikasi berprioritas rendah yang digabung menjadi satu kiriman. |
| Pengingat | reminder | Notifikasi terjadwal untuk rutinitas, tagihan, atau tugas yang tertunda. |

## 6. Pendamping keluarga (AI)

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Pendamping keluarga | family companion | Fitur AI yang membantu anggota saling memperhatikan. Ia membantu; orang yang bertindak. Ia tidak pernah mengirim pesan, menyetujui permintaan, atau mengubah catatan atas nama siapa pun, dan hanya memakai data yang memang boleh dilihat anggota yang dibantunya. Setiap anggota ikut serta secara terpisah dan dapat mematikannya. |
| Sapaan pagi | morning greeting | Satu kalimat pendek penyemangat atau kasih sayang yang muncul saat aplikasi dibuka. |
| Saran kata dukungan | suggested words of support | Draf pesan penyemangat yang ditawarkan kepada pasangan ketika anggota membagikan suasana hati yang rendah atau sakit. Pasangan membaca lalu mengirim, menyunting, atau meminta revisi; pesan terkirim sebagai catatan pendek yang menempel pada status itu, bukan lewat chat umum. |
| Pengingat hangat | warm reminder | Pengingat tugas, permintaan, atau jalan-jalan yang disusun dengan kata-kata ramah, bukan peringatan polos. |
| Narasi bulanan | monthly narrative | Cerita pendek tentang uang bulan itu, apa yang dicapai pasangan bersama, dan kesan. Untuk bagian uangnya AI hanya menerima angka agregat. |
| Penyedia AI | AI provider | Pihak di luar server Family OS yang menjalankan model AI. Hanya menerima data minimum, tanpa nama, dan hanya atas persetujuan anggota yang datanya dikirim (brief §5O). |
| Batas biaya AI | AI cost cap | Batas keras pengeluaran untuk pemakaian AI; ketika tercapai, fitur AI berhenti dan aplikasi tetap berfungsi penuh. |

## 7. Istilah keuangan

*Bagian ini ditulis oleh `finance-domain`. Versi S1 berisi definisi kerja; difinalkan di S2 bersama keputusan ledger di dokumen 07.*

**Cara baca bagian ini.** Semua definisi di bawah adalah *definisi kerja*. Bagian yang bergantung pada keputusan S2 (model ledger, apakah dana pribadi berada di dalam garis kendali rumah tangga, apakah persetujuan mencadangkan uang, nasib sisa anggaran di tiap situasi custody) atau pada kesepakatan keluarga yang belum dinyatakan ditandai *(diputuskan di S2)*; definisinya ditulis agar tetap benar apa pun pilihannya. Bila hasilnya berbeda menurut letak dana pribadi, definisi menyebut kedua hasilnya dengan frasa "bila dana pribadi di dalam garis" dan "bila di luar garis". Istilah di sini adalah bahasa model internal dan boleh berbau akuntansi; kata yang dilihat keluarga ditetapkan di dokumen 08.

Family OS **mencatat** uang; ia tidak menyimpan, memindahkan, atau memproses uang (lihat *PJP*, bagian 8). Setiap "pemindahan" di bawah adalah catatan atas perpindahan yang terjadi di dunia nyata.

Tiga aturan dari brief §7 berlaku untuk semua definisi:

1. Pemindahan antar-tempat uang yang berada dalam kendali rumah tangga tidak pernah menjadi pemasukan atau pengeluaran.
2. Persetujuan dan otorisasi saja tidak mengurangi uang yang ada.
3. Catatan uang tidak pernah disunting atau dihapus; kesalahan diperbaiki dengan entri pembalik atau entri penyesuaian.

### 7.1 Dasar pencatatan

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Jumlah uang | amount | Bilangan bulat dalam satuan terkecil mata uangnya, selalu disertai mata uangnya. Tidak pernah pecahan atau floating point (`ENGINEERING-STANDARDS.md` §5). Satuan terkecil yang dipakai untuk IDR *(diputuskan di S2)*: usulan bawaannya rupiah utuh, karena uang tunai dan hampir semua transaksi rumah tangga tidak mengenal sen; pilihan lainnya sen (1/100 rupiah), agar saldo rekening bank yang memuat sen dapat dicatat persis. |
| Entri | entry | Satu catatan atas satu peristiwa uang (pemasukan, pengeluaran, pemindahan, dan seterusnya) yang, setelah tersimpan, tidak dapat diubah atau dihapus. Bentuk entri (berpasangan debit-kredit, satu baris bertipe, atau lainnya) bergantung pada model ledger *(diputuskan di S2)*. Cara menyelaraskan sifat ini dengan hak penghapusan data anggota diputuskan di dokumen 07 dan 10. |
| Garis kendali rumah tangga | household boundary | Garis yang memisahkan uang yang dihitung sebagai uang rumah tangga dari uang lainnya. Selalu di dalam garis: dana bersama, siapa pun pemegangnya dan di mana pun disimpan, termasuk titipan pada pihak luar. Selalu di luar garis: uang milik pihak luar. Dana pribadi anggota berada di dalam atau di luar garis *(diputuskan di S2)*. Garis ini menentukan jenis setiap pergerakan uang: yang melintas masuk adalah pemasukan (atau melahirkan kewajiban, atau melunasi piutang), yang melintas keluar adalah pengeluaran (atau melahirkan piutang, atau melunasi kewajiban), dan yang bergerak di dalamnya adalah pemindahan. |
| Dalam kendali rumah tangga | household-controlled | Sifat tempat uang yang isinya berada di dalam *garis kendali rumah tangga*. Apakah satu tempat uang boleh memuat uang lebih dari satu pemilik (misalnya dompet yang berisi dana bersama dan dana pribadi sekaligus) *(diputuskan di S2)*. |
| Tanggal kejadian | event date | Tanggal peristiwa uangnya terjadi di dunia nyata, diisi pencatat. Menentukan periode anggaran dan laporan tempat entri dihitung. |
| Tanggal pencatatan | recorded date | Waktu entri disimpan, diisi sistem dan tidak dapat diatur pencatat. Menentukan urutan di jejak audit. |
| Entri terlambat | late entry | Entri yang tanggal kejadiannya lebih awal daripada tanggal pencatatannya (belanja minggu lalu yang baru dicatat hari ini). Entri terlambat adalah entri baru yang sah, bukan koreksi dan bukan penyuntingan riwayat. Seberapa jauh ke belakang ia diperbolehkan, dan bagaimana laporan periode yang sudah lewat menandai perubahannya, *(diputuskan di S2)*. |
| Saldo awal | opening balance | Entri pertama pada satu tempat uang ketika rumah tangga mulai mencatat, atau ketika tempat uang itu baru didaftarkan: jumlah yang sudah ada di sana saat itu, biasanya hasil menghitung atau melihat saldo. Bukan pemasukan dan tidak dihitung terhadap anggaran mana pun. Kewajiban, piutang, dan aset yang sudah ada saat mulai mencatat dicatat dengan cara yang sama. |
| Kategori | category | Pengelompokan pemasukan atau pengeluaran menurut jenis keperluannya (misalnya dapur, listrik, transportasi). Dipakai oleh anggaran dan laporan. Kategori menjawab "untuk jenis apa"; *tujuan* (7.2) menjawab "uang ini disisihkan untuk apa". |

### 7.2 Custody dan enam dimensinya

Untuk setiap jumlah uang, model harus dapat menjawab enam pertanyaan berikut secara terpisah (brief §5C). Tidak satu pun diturunkan otomatis dari yang lain, dan tidak satu pun diturunkan dari *hubungan* atau *peran* (bagian 3).

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Custody | custody | Keadaan "siapa memegang uang yang mana, berapa, dan di mana" pada suatu saat. Custody adalah jawaban atas kriteria sukses pertama di brief §1.1. Custody tidak sama dengan kepemilikan: memegang uang tidak menjadikan seseorang pemiliknya. |
| Pemilik uang | money owner | Pihak yang berhak atas uang itu menurut hukum atau kesepakatan keluarga: rumah tangga secara bersama, atau satu anggota. Berbeda dari *pemilik rumah tangga* (peran, bagian 3) dan dari *pemegang*. |
| Pemegang | holder | Orang yang saat ini menguasai uang itu secara fisik atau lewat aksesnya: anggota, atau pihak luar yang dititipi. Pemegang berubah lewat *serah terima*; pemilik tidak ikut berubah karenanya. |
| Tempat uang | holding | Wadah bernama tempat uang disimpan: uang tunai di tangan seseorang, satu rekening bank, atau satu dompet digital. Setiap tempat uang punya saldo tercatat sendiri. Istilah ini sengaja bukan "akun": *akun* hanya berarti identitas login (bagian 3). Apakah pemegang merupakan atribut tempat uang atau dimensi tersendiri *(diputuskan di S2)*. |
| Uang tunai | cash holding | Jenis tempat uang: uang kertas dan logam yang dipegang satu orang. Satu-satunya jenis yang saldonya hanya bisa dipastikan dengan dihitung (lihat *rekonsiliasi kas*). |
| Rekening bank | bank holding | Jenis tempat uang: simpanan di bank atas nama satu orang atau lebih. |
| Dompet digital | e-wallet holding | Jenis tempat uang: saldo di layanan uang elektronik (misalnya GoPay, OVO, DANA). |
| Tempat uang bersama | joint holding | Tempat uang yang dapat diakses lebih dari satu anggota, misalnya rekening bank atas nama berdua, sehingga tidak punya satu pemegang (brief §5C: "shared bank account"). "Bersama" di sini menyangkut akses, bukan kepemilikan: jangan disamakan dengan *dana bersama*, yang tetap dana bersama walau disimpan di rekening atas nama satu orang. Cara custody melaporkan isinya *(diputuskan di S2, bersama keputusan tentang pemegang)*. |
| Tujuan | purpose | Untuk apa sejumlah uang disisihkan (misalnya belanja dapur bulan ini, dana darurat). Ditetapkan lewat *alokasi*; uang tanpa tujuan disebut *belum dialokasikan*. |
| Yang berwenang membelanjakan | authorized spender | Anggota yang boleh membelanjakan uang itu tanpa meminta lagi, karena aturan rumah tangga atau karena memegang *otorisasi*. Berbeda dari *wewenang persetujuan* (bagian 3), yaitu hak menyetujui permintaan orang lain. |
| Penerima manfaat | beneficiary | Untuk siapa suatu pengeluaran dilakukan: seluruh rumah tangga, satu anggota, atau pihak luar. Berbeda dari pihak yang menerima pembayaran (penjual) dan dari anggota yang membayar. |
| Dana bersama | shared funds | Uang yang pemiliknya rumah tangga secara bersama, siapa pun pemegangnya. Uang bersama yang dipegang istri tetap dana bersama. |
| Dana pribadi | personal funds | Uang yang pemiliknya satu anggota. Apakah ia dihitung sebagai uang rumah tangga bergantung pada letaknya terhadap *garis kendali rumah tangga* *(diputuskan di S2)*. Visibilitasnya bagi anggota lain mengikuti aturan *data pribadi anggota* (bagian 3, dokumen 10). |
| Titipan | entrusted money | Uang yang pemegangnya bukan pemiliknya: dana bersama yang dipegang satu anggota, atau uang rumah tangga yang dipegang pihak luar untuk keperluan rumah tangga (contoh brief §5C: ibu mertua memegang uang makan). Titipan masih dalam kendali rumah tangga dan masih dihitung sebagai uang yang ada; ia bukan pengeluaran dan bukan piutang. |

### 7.3 Pergerakan uang

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Pemasukan | income | Uang yang melintas masuk ke dalam garis kendali rumah tangga dan menambah kekayaan bersih tanpa kewajiban mengembalikan (gaji, hasil usaha, hadiah). Bukan pemasukan: pemindahan, uang pinjaman yang diterima (itu utang), pelunasan piutang, penggantian, dan saldo awal. Uang pribadi yang diserahkan seorang anggota menjadi dana bersama adalah pemindahan bila dana pribadi di dalam garis, dan pemasukan bila di luar garis. Pemilik pemasukan dicatat saat ia masuk; siapa pemilik gaji adalah kesepakatan keluarga *(diputuskan di S2)*. |
| Pengeluaran | expense | Pemakaian uang untuk barang, jasa, atau pemberian tanpa harapan kembali. Lazimnya berupa uang yang melintas keluar dari garis kendali rumah tangga. Satu-satunya pengecualian adalah *talangan* dari dana yang berada di luar garis: pengeluarannya diakui lewat lahirnya hak ganti, tanpa ada uang yang keluar dari tempat uang dalam kendali rumah tangga saat itu. Membawa kategori, penerima manfaat, dan anggota yang membayar. Bukan pengeluaran: pemindahan, titipan, pinjaman yang diberikan (itu piutang), pelunasan pokok utang, dan penggantian. Dana bersama yang diserahkan menjadi dana pribadi seorang anggota (misalnya uang saku) adalah pemindahan bila dana pribadi di dalam garis, dan pengeluaran bila di luar garis. Berbeda dari *kewajiban*: pengeluaran sudah terjadi, kewajiban adalah uang yang masih harus dikeluarkan. |
| Pemindahan | transfer | Perpindahan uang antara dua tempat uang yang sama-sama dalam kendali rumah tangga: tarik tunai ATM, isi saldo dompet digital, setor ke rekening, serah terima. Satu tempat uang berkurang dan yang lain bertambah dengan jumlah yang sama, sehingga total uang yang ada tidak berubah. Tidak pernah pemasukan atau pengeluaran. Berbeda dari *alokasi*, yang tidak memindahkan uang sama sekali. Biaya admin yang menyertainya adalah uang yang keluar dari garis kendali rumah tangga, jadi ia pengeluaran dan bukan bagian dari jumlah yang dipindahkan; apakah biaya itu dicatat dalam entri yang sama atau entri terpisah *(diputuskan di S2)*. |
| Serah terima | handover | Pemindahan yang mengganti *pemegang*: satu orang menyerahkan uang kepada orang lain (misalnya suami menyerahkan sisa gaji kepada istri). Pemilik uang tidak berubah karena serah terima saja. Apakah penerima harus mengonfirmasi *(diputuskan di S2, bersama dokumen 06)*. |
| Alokasi | allocation | Penetapan atau perubahan *tujuan* atas uang yang benar-benar ada, tanpa memindahkannya dan tanpa mengubah saldo tercatat tempat uang mana pun. Membagi gaji menjadi "tagihan", "dapur", dan "tabungan" adalah alokasi; menyerahkan bagian dapur kepada istri adalah serah terima yang terpisah. Berbeda dari *reservasi*: alokasi memberi label tujuan, reservasi mengunci jumlah untuk satu otorisasi. Apakah jumlah seluruh alokasi dibatasi oleh uang yang ada, dan apa yang terjadi bila pengeluaran membuat alokasi melampaui uang yang tersisa, *(diputuskan di S2)*. |
| Belum dialokasikan | unallocated | Bagian uang yang ada yang belum diberi tujuan. |
| Pengembalian dana | refund | Uang yang dikembalikan penjual, seluruhnya atau sebagian, atas suatu pengeluaran yang sudah tercatat, dan selalu ditautkan ke pengeluaran itu. Berbeda dari *koreksi*: pengembalian dana adalah peristiwa nyata, bukan perbaikan catatan yang salah. Apakah dilaporkan sebagai pengurang pengeluaran atau sebagai pemasukan *(diputuskan di S2)*. |
| Kiriman keluarga | family support | Uang yang diberikan kepada keluarga besar tanpa harapan kembali (misalnya kiriman bulanan Rp3.000.000 kepada ibu suami). Uang itu keluar dari kendali rumah tangga, jadi ia pengeluaran dengan pihak luar sebagai penerima manfaat. Berbeda dari *pinjaman yang diberikan* (diharapkan kembali, menjadi piutang) dan dari *titipan* (masih uang rumah tangga). Golongan mana yang berlaku untuk tiap pemberian adalah kesepakatan keluarga *(diputuskan di S2)*. |

### 7.4 Anggaran, otorisasi, dan penyelesaian

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Anggaran | budget | Batas belanja yang direncanakan untuk satu kategori dalam satu periode. Anggaran adalah rencana, bukan uang: membuat atau mengubahnya tidak mengubah saldo apa pun. Apakah anggaran harus didukung uang yang dialokasikan (cara amplop) atau cukup sebagai batas *(diputuskan di S2)*. Kata "anggaran" dalam *permintaan anggaran* (bagian 4) berarti jumlah yang dimintakan otorisasi, bukan anggaran dalam arti ini. |
| Periode anggaran | budget period | Rentang waktu berlakunya satu anggaran (harian, mingguan, bulanan, atau khusus). Tanggal mulai periode bulanan, misalnya tanggal gajian, ditetapkan rumah tangga. |
| Sisa anggaran | budget remaining | Anggaran dikurangi jumlah pengeluaran berkategori itu dalam periodenya; bisa negatif bila terlampaui. Ini angka rencana dan bukan saldo: sisa anggaran bisa positif ketika uangnya sudah tidak ada, dan sebaliknya. Apakah jumlah *terikat*, pembayaran cicilan, dan selisih kas ikut mengurangi *(diputuskan di S2)*; apa pun pilihannya, satu entri dihitung paling banyak satu kali terhadap satu anggaran. |
| Otorisasi | authorization | Catatan keuangan yang lahir ketika suatu permintaan anggaran disetujui: pemohon boleh membelanjakan sampai jumlah tertentu, dari dana tertentu, untuk tujuan tertentu. *Persetujuan* (bagian 4) adalah tindakan pemberi persetujuan; otorisasi adalah akibat keuangannya. Otorisasi tidak mengubah saldo tercatat tempat uang mana pun. Apakah ia membuat reservasi *(diputuskan di S2)*. Berakhir lewat penyelesaian, pembatalan, atau kedaluwarsa. |
| Reservasi | reservation | Sejumlah uang di satu tempat uang atau satu alokasi yang dikunci untuk satu otorisasi sehingga tidak dapat dipakai untuk hal lain. Mengurangi saldo tersedia, tidak mengurangi saldo tercatat. Dilepas saat uangnya dibelanjakan, otorisasinya diselesaikan, atau dibatalkan. Kapan reservasi dibuat, atau apakah dibuat sama sekali, mengikuti aturan eksplisit *(diputuskan di S2)*. |
| Terikat | committed | Bagian dari satu otorisasi yang masih terbuka dan belum dibelanjakan: jumlah yang diotorisasi dikurangi realisasinya, selama otorisasi itu belum berakhir. Terikat adalah angka rencana, bukan uang: ia tidak mengubah saldo tercatat maupun saldo tersedia. Berbeda dari *reservasi*, yang mengunci uang di satu tempat uang atau alokasi; suatu jumlah bisa terikat tanpa direservasi. Apakah jumlah terikat mengurangi sisa anggaran *(diputuskan di S2)*. |
| Realisasi | actual spend | Jumlah yang benar-benar dibelanjakan atas satu otorisasi atau satu anggaran, yaitu jumlah entri pengeluaran yang tertaut padanya. |
| Penyelesaian | settlement | Penutupan satu otorisasi setelah belanjanya selesai: realisasi dibandingkan dengan jumlah yang diotorisasi, lalu selisihnya (sisa atau kelebihan) diperlakukan menurut situasi custody-nya. Penyelesaian sendiri bukan pengeluaran; pengeluarannya sudah tercatat saat membeli. Aturan sisa dan kelebihan untuk tiap situasi custody *(diputuskan di S2)* dan kanonis di dokumen 07. Berbeda dari *pelunasan* (membayar utang) dan dari *penggantian*. Brief §5D menyebut langkah ini "reconciled"; di glosarium ini kata *rekonsiliasi* hanya dipakai untuk rekonsiliasi kas. |
| Talangan | advance payment | Pembayaran oleh seorang anggota dengan dana pribadinya untuk keperluan bersama atau untuk anggota lain. Dicatat satu kali, saat dibayar, sebagai pengeluaran pihak yang ditalangi (dana bersama, bila keperluannya bersama), dan melahirkan *hak ganti*. Akibat talangan keperluan bersama pada angka rumah tangga: bila dana pribadi di dalam garis, uang yang ada berkurang sebesar talangan dan hak gantinya adalah klaim internal; bila di luar garis, uang yang ada tidak berubah dan hak gantinya menjadi kewajiban rumah tangga kepada anggota itu. Pada kedua pilihan kekayaan bersih rumah tangga berkurang tepat sebesar talangan, satu kali. |
| Hak ganti | reimbursement claim | Hak anggota yang menalangi untuk menerima kembali jumlah talangannya dari pihak yang ditalangi. Satu klaim dengan dua sisi yang sama besar: piutang bagi yang menalangi, kewajiban bagi yang ditalangi. Lahir saat talangan dicatat dan berakhir saat penggantian dibayar, atau saat dilepaskan dengan entri yang terlihat. Bentuk pencatatannya *(diputuskan di S2)*. |
| Penggantian | reimbursement | Pembayaran dari pihak yang ditalangi (dana bersama atau anggota lain) kepada anggota yang menalangi, yang melunasi hak ganti. Bukan pemasukan bagi penerimanya dan bukan pengeluaran kedua; pengeluarannya sudah tercatat pada talangan. Bila dana pribadi di dalam garis, penggantian adalah pemindahan dan uang yang ada tidak berubah; bila di luar garis, uang yang ada dan kewajiban rumah tangga berkurang sama besar. Pada kedua pilihan penggantian tidak mengubah kekayaan bersih rumah tangga. Berbeda dari *penyelesaian*: penggantian melunasi hak ganti atas belanja yang dibayar dengan dana pribadi; penyelesaian menutup satu otorisasi dengan membandingkan realisasi dan jumlah yang diotorisasi, entah uangnya sudah dipegang pemohon atau diserahkan untuk permintaan itu. |

### 7.5 Kewajiban, utang, dan piutang

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Kewajiban | liability | Jumlah yang masih harus dibayar satu pemilik uang (rumah tangga atas dana bersama, atau satu anggota atas dana pribadinya) kepada pihak lain: utang, sisa cicilan, hak ganti yang belum dibayar. Setiap kewajiban punya cermin, yaitu *piutang* di sisi pihak yang berhak. Kewajiban mengurangi kekayaan bersih pemilik yang menanggungnya, tetapi tidak mengurangi uang yang ada sampai dibayar. *Kewajiban rumah tangga*, angka yang masuk kekayaan bersih rumah tangga, hanya memuat kewajiban kepada pihak di luar garis kendali rumah tangga; kewajiban antara dua pemilik yang sama-sama di dalam garis adalah *klaim internal*. Berbeda dari *pengeluaran*: timbulnya kewajiban belum mengeluarkan uang. Apakah tagihan yang sudah jatuh tempo tetapi belum dibayar dihitung sebagai kewajiban *(diputuskan di S2)*. |
| Piutang | receivable | Hak satu pemilik uang untuk menerima uang dari pihak lain: cermin kewajiban, di sisi pihak yang berhak. Lahir dari pinjaman yang diberikan atau dari talangan (hak ganti). Menambah kekayaan bersih pemiliknya tetapi bukan uang yang ada. *Piutang rumah tangga*, angka yang masuk kekayaan bersih rumah tangga, hanya memuat piutang kepada pihak di luar garis kendali rumah tangga. |
| Klaim internal | internal claim | Klaim antara dua pemilik uang yang sama-sama berada di dalam garis kendali rumah tangga: hak ganti dan utang antar-pasangan. Selalu tercatat dengan dua sisi yang sama besar (piutang satu pemilik, kewajiban pemilik lain), sehingga tidak mengubah kekayaan bersih rumah tangga dan tidak masuk angka kewajiban maupun piutang rumah tangga; ia hanya tampil pada rincian per pemilik. Bila dana pribadi di luar garis, tidak ada klaim internal: klaim antara dana bersama dan seorang anggota adalah kewajiban atau piutang rumah tangga kepada anggota itu. |
| Utang | debt (payable) | Kewajiban yang lahir karena menerima uang atau barang dengan janji mengembalikan. Uang pinjaman yang diterima dari luar garis kendali rumah tangga menambah uang yang ada dan menambah utang; ia bukan pemasukan. |
| Pinjaman yang diberikan | loan given | Uang dari dalam garis kendali rumah tangga yang diserahkan kepada pihak luar (misalnya saudara atau ayah suami) dengan harapan dikembalikan. Uang yang ada berkurang dan piutang bertambah dengan jumlah yang sama, sehingga kekayaan bersih tidak berubah; ia bukan pengeluaran. Pinjaman antara dua pemilik uang di dalam satu rumah tangga diatur di *utang antar-pasangan*, karena akibatnya berbeda. Bila kelak diikhlaskan, lihat *penghapusan piutang*. |
| Penghapusan piutang | receivable write-off | Entri yang menghapus sebagian atau seluruh piutang yang diikhlaskan atau tidak akan tertagih. Piutang dan kekayaan bersih berkurang sebesar yang dihapus; uang yang ada tidak berubah, karena uangnya sudah keluar saat dipinjamkan. Selalu berupa entri yang terlihat beserta pelaku dan alasannya; catatan pinjamannya tidak pernah dihapus. Melepaskan klaim internal tidak mengubah kekayaan bersih rumah tangga. Apakah jumlah yang dihapus dilaporkan sebagai kiriman keluarga atau sebagai kerugian tersendiri, dan terhadap anggaran periode mana, *(diputuskan di S2)*. |
| Pelunasan | repayment | Pembayaran yang mengurangi utang di sisi yang membayar dan piutang di sisi yang menerima, sebagian atau seluruhnya. Pelunasan pokok bukan pengeluaran dan bukan pemasukan. |
| Utang antar-pasangan | inter-spouse debt | Klaim antara dua pemilik uang di satu rumah tangga (dana pribadi satu anggota, dana pribadi anggota lain, dana bersama): utang bagi satu sisi, piutang bagi sisi lain, selalu sama besar. Hanya ada bila uang berpindah antara dana yang pemiliknya berbeda dengan janji dikembalikan. Akibatnya pada angka rumah tangga bergantung pada letak dana pribadi. Bila di dalam garis: perpindahan uangnya adalah pemindahan, uang yang ada tidak berubah, dan klaimnya klaim internal. Bila di luar garis dan salah satu sisinya dana bersama: uang yang ada berubah sebesar pinjaman, dan rumah tangga mencatat piutang (bila dana bersama yang meminjamkan) atau kewajiban (bila dana bersama yang meminjam) kepada anggota itu. Bila di luar garis dan kedua sisinya dana pribadi: angka rumah tangga tidak tersentuh. Pada semua pilihan kekayaan bersih rumah tangga tidak berubah oleh pinjaman maupun pelunasannya. |
| Pinjam kembali | borrow-back | Satu pasangan mengambil lagi uang dari yang dipegang pasangannya (brief §1.1). Bila uang itu dana bersama, ia serah terima saja, tanpa utang. Bila uang itu dana pribadi pemegangnya, ia melahirkan utang antar-pasangan, dan akibatnya pada angka rumah tangga mengikuti baris itu. Golongan mana yang berlaku bergantung pada siapa pemilik uang yang dipegang *(diputuskan di S2; kesepakatan keluarga)*. |
| Cicilan | installment | Pembayaran berkala berjadwal atas satu kewajiban dengan jumlah dan jangka yang diketahui (misalnya kredit kendaraan atau barang). Tiap pembayaran mengurangi uang yang ada. Apakah ia dicatat sebagai pengurang kewajiban, sebagai pengeluaran, atau dipisah menjadi pokok dan bunga *(diputuskan di S2)*. |
| Tagihan berulang | recurring bill | Pembayaran yang diperkirakan datang lagi tiap periode kepada pihak luar yang sama (listrik, wifi, langganan kerja), dengan perkiraan jumlah, tanggal jatuh tempo, dan anggota yang bertanggung jawab membayar. Tagihan berulang adalah jadwal dan pengingat, bukan entri; entri pengeluaran baru ada saat dibayar. Berbeda dari *cicilan*: tagihan berulang tidak melunasi suatu kewajiban yang jumlah seluruhnya diketahui, jadi tidak punya pokok yang berkurang. Apakah tagihan berulang boleh diberi tanggal akhir (misalnya langganan berjangka) *(diputuskan di S2)*. |

### 7.6 Saldo, tabungan, dan kekayaan

| Istilah | Term (EN) | Definisi |
|---|---|---|
| Saldo tercatat | recorded balance | Jumlah uang di satu tempat uang menurut catatan: hasil penjumlahan semua entri yang menyangkut tempat uang itu. Hanya berubah oleh entri, tidak oleh persetujuan, otorisasi, reservasi, alokasi, atau anggaran. |
| Saldo tersedia | available balance | Saldo tercatat dikurangi reservasi yang masih aktif: bagian yang masih bebas dipakai. Tidak pernah lebih besar dari saldo tercatat. Bila S2 memutuskan tidak ada reservasi, saldo tersedia sama dengan saldo tercatat. |
| Saldo terhitung | counted balance | Jumlah uang yang benar-benar ditemukan ketika seorang pemegang menghitung uang tunainya atau melihat saldo rekening atau dompet digitalnya. Angka dari dunia nyata, bukan dari catatan. |
| Uang yang ada | actual money | Jumlah saldo tercatat semua tempat uang dalam kendali rumah tangga. Inilah jawaban "berapa uang kita"; rinciannya per pemegang adalah custody. Rencana, target, anggaran, piutang, dan klaim internal tidak termasuk. Angka mana yang ditampilkan kepada tiap anggota bila ada dana pribadi yang tidak dibagikan *(diputuskan di S2 dan S4)*. |
| Rekonsiliasi kas | cash reconciliation | Membandingkan saldo terhitung dengan saldo tercatat satu tempat uang. Selisihnya tidak diperbaiki dengan menyunting entri lama, tetapi dicatat sebagai entri penyesuaian yang terlihat (*selisih kas*), sehingga saldo tercatat kembali sama dengan saldo terhitung. |
| Selisih kas | cash difference | Saldo terhitung dikurangi saldo tercatat pada satu rekonsiliasi kas, dan entri penyesuaian yang mencatatnya. Negatif berarti ada uang keluar yang tidak tercatat (paling sering belanja yang lupa dicatat); positif berarti ada uang masuk yang tidak tercatat atau pengeluaran yang tercatat terlalu besar. Selisih kas mengubah saldo tercatat dan uang yang ada, dan selalu tampil sebagai angka tersendiri. Apakah ia dihitung terhadap anggaran, bagaimana ia tampil di laporan, dan apakah ia boleh dirinci belakangan ke kategori lewat koreksi *(diputuskan di S2)*. |
| Koreksi | correction | Perbaikan atas entri yang salah tanpa mengubah entri itu: *entri pembalik* (reversing entry) meniadakan seluruh akibat entri asli, lalu entri baru yang benar dicatat; atau *entri penyesuaian* (adjusting entry) mencatat selisihnya saja. Entri asli, koreksinya, pelakunya, dan alasannya tetap terlihat di jejak audit. |
| Tabungan | savings | Uang yang benar-benar ada dan dialokasikan untuk disimpan, bukan dibelanjakan dalam periode berjalan (misalnya dana darurat). Menabung adalah alokasi, atau pemindahan bila uangnya dipindah ke tempat uang khusus; ia tidak pernah pengeluaran. |
| Target | target | Jumlah yang ingin dicapai suatu tabungan atau tujuan. Target adalah niat, bukan uang: tidak masuk saldo, uang yang ada, atau kekayaan bersih. Kemajuan dihitung dari uang yang benar-benar dialokasikan, bukan dari target. |
| Aset | asset | Barang bernilai milik rumah tangga atau anggota yang bukan uang di tempat uang (kendaraan, emas, rumah, perangkat), dicatat dengan perkiraan nilainya. Dalam akuntansi uang dan piutang juga aset; di sini kata ini sengaja dipersempit agar sama dengan pertanyaan keluarga "aset apa yang kita punya". Membeli aset mengurangi uang yang ada dan menambah aset; apakah pembeliannya juga dilaporkan sebagai pengeluaran, dan bagaimana nilainya diperbarui, *(diputuskan di S2)*. Aset milik pribadi seorang anggota mengikuti letak dana pribadi terhadap garis kendali rumah tangga. |
| Kekayaan bersih | net worth | Uang yang ada, ditambah piutang rumah tangga, ditambah nilai aset di dalam garis kendali rumah tangga, dikurangi kewajiban rumah tangga. Piutang dan kewajiban rumah tangga hanya memuat klaim terhadap pihak di luar garis; klaim internal tidak dihitung karena kedua sisinya sama besar (menghitung kedua sisinya memberi hasil yang sama; menghitung satu sisi saja adalah kesalahan). Bila dana pribadi di luar garis, seorang anggota terhitung pihak di luar garis untuk dana pribadinya, sehingga hak ganti dan utang antar-pasangan yang menyangkut dana bersama masuk sebagai kewajiban atau piutang rumah tangga. Kekayaan bersih juga dapat dihitung per pemilik uang dengan rumus yang sama, dengan klaim internal ikut dihitung; jumlah kekayaan bersih semua pemilik di dalam garis sama dengan kekayaan bersih rumah tangga. Pemindahan, alokasi, otorisasi, reservasi, penggantian, dan pelunasan pokok tidak mengubahnya. |
| Arus kas | cash flow | Perubahan uang yang ada dalam satu periode, dirinci menurut sebabnya: uang yang melintas masuk dan keluar garis kendali rumah tangga, serta selisih kas. Pemindahan tidak muncul karena tidak mengubah uang yang ada. Lebih luas daripada pemasukan dikurangi pengeluaran, karena memuat juga uang masuk dan keluar yang bukan keduanya, misalnya uang pinjaman dan pelunasan. |

## 8. Regulasi dan perlindungan data

Istilah hukum di bawah mengikuti pemakaian di [18-market-research.md](18-market-research.md), yang memuat sumber, tanggal akses, dan label buktinya. Glosarium ini hanya menyeragamkan kata; ia bukan nasihat hukum.

| Istilah | Term (EN) | Definisi |
|---|---|---|
| UU PDP | UU PDP (Personal Data Protection Law) | Undang-Undang No. 27 Tahun 2022 tentang Pelindungan Data Pribadi. |
| Data pribadi | personal data (legal) | Data tentang orang perseorangan yang teridentifikasi atau dapat diidentifikasi, dalam arti UU PDP. Lebih luas daripada *data pribadi anggota* (bagian 3), yang merupakan istilah produk tentang visibilitas. |
| Data pribadi spesifik | specific personal data | Kategori data pribadi yang menurut UU PDP berisiko lebih tinggi; yang relevan bagi Family OS: data kesehatan, data anak, dan data keuangan pribadi. |
| Subjek data | data subject | Orang perseorangan yang datanya diproses; di Family OS, setiap anggota dan setiap pihak luar yang tercatat. |
| Pengendali data | data controller | Pihak yang menentukan tujuan dan mengendalikan pemrosesan data pribadi; di Family OS, operator layanan. |
| Prosesor data | data processor | Pihak yang memproses data pribadi atas nama pengendali, misalnya penyedia AI atau penyedia hosting. |
| Sub-prosesor | sub-processor | Pihak lain yang dilibatkan prosesor data untuk mengerjakan sebagian pemrosesan itu. |
| Persetujuan pemrosesan | processing consent | Persetujuan subjek data, dalam arti UU PDP, atas pemrosesan datanya untuk tujuan tertentu. Berbeda dari *persetujuan* atas permintaan (bagian 4); *persetujuan berbagi* (bagian 5) adalah pilihan produk yang harus memenuhinya bila data itu keluar dari rumah tangga. |
| Transfer lintas batas | cross-border transfer | Pengiriman data pribadi ke luar wilayah Indonesia, misalnya ke penyedia AI di luar negeri. |
| DPIA | DPIA (data protection impact assessment) | Penilaian dampak pelindungan data pribadi yang dilakukan pengendali untuk pemrosesan berisiko tinggi. |
| Lembaga PDP | PDP agency | Lembaga pengawas pelindungan data pribadi yang diamanatkan UU PDP. |
| Anak | child | Orang berusia di bawah 18 tahun. |
| Profil risiko | risk profile | Tingkat risiko suatu produk atau fitur digital bagi anak menurut aturan pelindungan anak di sistem elektronik; menentukan apakah anak pada rentang usia tertentu boleh punya akun. Rincian dan sumbernya di dokumen 18 §2.2. |
| Penilaian mandiri | self-assessment | Penilaian yang dilakukan penyelenggara sendiri atas profil risiko produknya, sebagaimana diminta aturan yang sama. |
| Pejabat pelindungan data | data protection officer (PPDP/DPO) | Pejabat atau petugas yang menjalankan fungsi pelindungan data pribadi pada pengendali atau prosesor, dalam arti UU PDP. |
| Profil tanggungan tanpa login | dependent profile (no login) | *(usulan; lihat dokumen 17)* Tanggungan yang tercatat di rumah tangga tanpa akun sendiri; datanya dikelola walinya. |
| PSE | PSE (electronic system operator) | Penyelenggara Sistem Elektronik menurut aturan Indonesia. |
| PJP | PJP (payment service provider) | Penyedia Jasa Pembayaran yang berizin Bank Indonesia. Family OS bukan PJP: ia mencatat uang, tidak menyimpan, memindahkan, atau memprosesnya. |
| De-identifikasi | de-identification | Membuang nama dan penanda identitas lain dari data sebelum data itu dikirim keluar, misalnya ke penyedia AI, sehingga penerima tidak dapat mengenali orangnya. Berbeda dari *pseudonimisasi*: tidak ada penanda pengganti yang disimpan untuk mengaitkannya kembali. |
| Pseudonimisasi | pseudonymization | Mengganti identitas pada suatu catatan dengan penanda yang tidak dapat dikaitkan ke orangnya tanpa data tambahan yang disimpan terpisah. |
