# Log Kemajuan

Dicatat lead setelah setiap tahap (brief §0.2). Bila konteks terhapus atau sesi dimulai ulang: baca berkas ini dan [00-glossary.md](00-glossary.md), lalu lanjutkan dari **Tahap berikutnya**.

## Posisi saat ini

| | |
|---|---|
| Tahap terakhir yang selesai | S2 — Fondasi domain (11 November 2026): dokumen 02, 03, 06, 07, dan ADR-001 sampai ADR-010 ditulis, diperiksa, dan diperbaiki |
| Status | **Checkpoint S2 dijawab (11 dan 12 November 2026).** Jawabannya disebarkan ke 06, 07, 17, 02, 03, 18, glosarium, dan ADR; putaran perbaikan atas temuan `critic` dan SP-32 sampai SP-44 selesai, ditinjau ulang `critic`, dan angkanya dihitung ulang lead (12 November 2026); satu commit revisi |
| Tahap berikutnya | S3 — Definisi produk (04, 05, 11, 12), **hanya setelah pemilik produk berkata lanjut** |

**Cara melanjutkan di sesi baru:** baca berkas ini dan glosarium. Bila `git status` masih menunjukkan perubahan yang belum di-commit, putaran perbaikan belum selesai: periksa temuan B1, B2, dan S1 sampai S15 di bagian "Tinjauan `critic`" di bawah terhadap berkasnya, lalu commit. Bila sudah di-commit, tunggu pemilik produk berkata lanjut, lalu mulai S3 dengan daftar "Dibawa ke S3" di dokumen 17.

**Catatan proses dari pemilik produk (11 November 2026):** ia menyerahkan cara kerja tim kepada lead asal hasilnya profesional dan valid; semua hal dijelaskan dan dikonfirmasi kepadanya; keputusan ditanyakan langsung. Demi kuota, pekerjaan S3 dijalankan dengan model yang lebih hemat dan gambaran UI/UX dikerjakan di S4; setelah tiap tugas selesai dan aman, lead memberi tahu bahwa percakapan boleh dibersihkan.

## Jawaban checkpoint S2

Dicatat lead per pertanyaan, dengan kata-kata pemilik produk. Bagian ini adalah catatan asal; dokumen lain yang mengutip pemilik produk harus dapat dilacak ke sini.

| Pertanyaan | Jawaban | Catatan |
|---|---|---|
| S2-1 | **B** | Butir (i) ya, (ii) tidak terlihat, (iii) **b**: dinyatakan pemilik produk sendiri. Kata-katanya atas (i), sesudah lead menjelaskan dua jalan mendapat uang pribadi dan menyarankan "hanya penerimanya": "butir 1 sudah bener". Atas (ii) ia mengulang rumusan pertanyaannya sebagai jawabannya: "Hadiah untuk satu orang: Masuk uang pribadinya dan tidak terlihat pasangan". Kata-katanya atas (iii): "jatah itu ya memang batas, dan harus ajukan kenaikan bukan keringanan, kecuali dia split jatah + uang pribadi bisa itu". Butir (iv) ya dan (v) **b**: ia menyerahkan kepada lead ("coba pikirkan bagusnya gimana ... wajib mudah dipahami dan profesional"); lead menerapkan bawaan tim dan menyampaikannya kepadanya; belum ia nyatakan sendiri. Atas (v) ia berkata alasan ditulis di catatan ("nanti ada notenya dan di sana lah alasannya"); dibaca lead sebagai catatan bebas di diskusi yang menempel pada usulan |
| S2-2 | **A** | Diterima atas saran lead ("saya ikut saran kamu aja"). Ia menegaskan barang mahal tetap diajukan dulu, dan bertanya tentang kasus sepakat secara langsung lalu baru dicatat sepulangnya. Syarat layar baru untuk dokumen 05 di S3 **[Usulan]** lead: dari catatan bertanda di atas ambang, pencatat dapat meminta konfirmasi susulan; pasangan menyetujui dengan satu ketukan; catatan tertaut ke persetujuan itu dan tandanya hilang (06 §6.2 "Tanda hilang"). Konfirmasi itu sukarela; diam tidak berakibat apa pun. Apakah permintaan boleh dibuat sesudah belanjanya terjadi perlu dipastikan `architect` di dokumen 06 |
| S2-3 | **A** | Dinyatakan pemilik produk sendiri ("a") |
| S2-4 | **A** | Bagian pembayaran rutin dinyatakan pemilik produk sendiri: "karena itu rutin untuk membuat tagihan rutin ya disepakati dulu"; siapa pun dari keduanya boleh membayar. Bagian ambang (hanya berubah bila berdua setuju; tidak ada yang keluar sendiri) **tidak ia sebut**; lead menerapkan A dan menyampaikannya; pendapat `security-privacy` (B) dibawa ke dokumen 10. Fakta rumah tangga pilot yang ia nyatakan: begitu penghasilannya masuk, ia membayar tagihan lebih dulu, baru membagikan sisanya; kiriman untuk orang tua **tidak rutin** di rumah tangganya (mengikuti aturan biasa, bukan kesepakatan rutin); tanggal tagihan tersebar sesudah tanggal uang masuk (contohnya: uang masuk tanggal 11, tagihan tanggal 12, 13, 20, 29) dan jumlah sebagian tagihan dapat jauh berbeda antarbulan. Untuk S3 (dokumen 05): alur "uang masuk, lalu bayar atau sisihkan untuk tagihan yang jatuh temponya menyusul" perlu skenario sendiri |
| S2-5 | **A** | Diterima atas saran lead. Lead menutup pesannya dengan dua butir bernomor: (1) "Setuju jam tenang bawaan 22.00–06.00, atau jam lain?" dan (2) "Pertanyaan 5 tetap perlu satu huruf: A …, B, C, atau D. Saran saya A". Jawabannya: "gas 2 2 nya aja", dibaca lead sebagai "dua-duanya": bawaan jam tenang 22.00–06.00 **dan** pilihan A. Jadi keduanya berkedudukan "diterima atas saran lead"; angka jamnya tetap **[Asumsi]** lead. Tambahan yang lahir dari pertanyaannya ("kalau tidak set ... bisa langsung diingatkan?"): **jam tenang punya bawaan 22.00 sampai 06.00 per orang**, dapat digeser atau dimatikan pemiliknya sendiri; angkanya **[Asumsi]** lead, diuji di pilot. Dokumen 06 §10 dan §11.1 belum menetapkan bawaan jam tenang; diteruskan ke `architect` saat penyebaran |
| S2-6 | **Semua bentuk terjadi** (bukan satu huruf) | Fakta rumah tangga, kata-katanya: "bisa terpisah bisa tercampur bisa di bank bisa tunai". Andaian bawaan "A atau C" **tidak berlaku sepenuhnya**: dompet campuran memang ada (pilihan B). Aturan "satu tempat uang, satu pemilik" dipertahankan dengan perlakuan 07 §2.3 dan T-03 yang sudah ada: satu dompet fisik dicatat sebagai dua tempat uang; menghitung seluruh isi dompet lalu membandingkannya dengan jumlah kedua saldo; selisih dibebankan ke yang bersama kecuali pemegang memilih lain. Lead menyampaikan perlakuan itu dan akibatnya kepadanya; ia belum menyatakan menerima atau menolak. Untuk `finance-domain` saat penyebaran: T-03 bukan lagi kasus tepi, perlu contoh terhitung dompet campuran bila belum ada; untuk dokumen 05 dan 08 (S3, S4): menghitung dompet campuran harus satu layar |
| S2-7 | **A** | Dinyatakan pemilik produk sendiri ("a"); angka 7 dan 30 hari tetap **[Asumsi]** |
| S2-8 | **A** | Dinyatakan pemilik produk sendiri ("a"): bacaan append-only diterima sebagai keputusan sementara sampai advokat menjawab (L-04, L-20, L-22, L-23) |
| S2-9 | **A** | Diterima atas saran lead ("ikut saran kamu asal profesional bener dan valid"). Kalimat persis untuk brief dan `CLAUDE.md` ditunjukkan kepadanya lebih dulu; ia menjawab "oke"; lead lalu memasangnya di brief (§0.3, baris Tier 1, §5C) dan di `CLAUDE.md` ("Sources of truth") |
| S2-10 | **A** | Dinyatakan pemilik produk sendiri ("a"). Syarat yang melekat: istrinya diberi tahu dengan terus terang sebelum mulai memakai; lead menyarankan istrinya ikut mendengar pertanyaan ini, belum diketahui apakah itu terjadi. Enkripsi terhadap pengelola server dinilai di S4 (Q-57) |
| W-1 sampai W-14 | dibacakan | W-15 dan W-16 tidak dibacakan (pilihan pemicunya tidak diambil). W-13 cara (1) gugur pada (iii) b; contoh urutan tiba pada (iii) b dihitung lead (urutan 400, 400, 900 ribu atas jatah 1 juta: satu tanda; 900 ribu lebih dulu: dua tanda) dan rumusannya ditulis ulang pemilik dokumen saat penyebaran. Ia tidak mengubah satu jawaban pun dan menyerahkan penanganan batas-batas itu kepada lead, dengan pernyataan arah di bawah tabel ini |
| D bertanda Pagar | dibacakan satu per satu, per tema | Ke-24 baris dibacakan dengan pilihan lainnya (D-01, D-02, D-49, D-52, D-54). D-48 dibacakan dengan penyesuaian S2-1 (iii) b |
| D-01 sampai D-55 | **diterima lewat "gas lanjut"** | Tidak satu pun diubah. Ke-31 baris yang tidak bertanda Pagar tidak dibacakan satu per satu; lead menawarkan membacakannya dan ia memilih lanjut. Semuanya default yang diterima, bukan pernyataannya sendiri |

**Status penyebaran:** dimulai 11 November 2026. Urutan: 07 dan glosarium bagian 7 (`finance-domain`), lalu 06 dan ADR (`architect`), lalu 17 dan pencocokan 02 (`product-analyst`), lalu `critic` dan hitung ulang angka oleh lead, lalu glosarium selebihnya, kalimat brief dan `CLAUDE.md` ditunjukkan kepada pemilik produk, satu commit revisi.

**Pernyataan arah pemilik produk sesudah peringatan dibacakan (11 November 2026), kata-katanya:** "yang penting perselisihan dan tidak bisa di otak atik, aplikasi ini bukan menunjukan siapa yang salah, tapi mencatat apa yang perlu di catat, liat perkembangan, status, bukan dibuat untuk perselisihan tapi untuk menyimpan data ('semua ada catatanya ada datanya')". Ia meminta lead memikirkan keluhan yang akan muncul dan jalan keluarnya. Bacaan lead **[Usulan]**, belum ia kukuhkan kata per kata: (1) catatan tidak dapat diubah diam-diam; (2) aplikasi mencatat dan menunjukkan keadaan, tidak menilai siapa yang salah; (3) tanda, keberatan, dan selisih memakai bahasa keterangan, bukan bahasa pelanggaran. Diteruskan saat penyebaran: ke `product-analyst` untuk dicocokkan dengan prinsip di dokumen 02 dan untuk daftar keluhan di dokumen 05; ke `ux-designer` sebagai syarat kata-kata layar di dokumen 08 (S4).

**Tiga penegasan sesudah revisi dokumen 07 (11 November 2026).** Lead menanyakan tiga hal yang ditemukan `finance-domain` (T-29; arti "bukan keringanan"; tagihan rutin mana yang disepakati). Kata-kata pemilik produk: "jatah pribadi itu ya harusnya jelas gitu, jatah pribadi bisa masuk uang pribadi, kalau belanja rutin ya itu kan batas kesepakatan bersama". Lead merumuskan ulang modelnya dan tiga akibatnya, dan ia menjawab "bener". Jadi ketiganya **rumusan lead yang ia kukuhkan**, bukan kata-katanya sendiri:

1. **T-29, pilihan (1):** jatah yang diakhiri pemiliknya tetap menjadi dasar penilaian sampai akhir periode anggaran itu; periode berikutnya barulah tanpa jatah. D-52 selebihnya tetap (hanya pemilik yang mengakhiri).
2. **Tambahan untuk satu bulan** boleh, lewat permintaan biasa yang disetujui pasangan (D-53 tetap). "Bukan keringanan" berarti tidak ada kelonggaran otomatis di atas jatah.
3. **Setiap pos rutin disepakati saat dibuat**, berapa pun jumlahnya, bukan hanya yang di atas ambang; jumlah yang disepakati adalah batasnya; pembayaran sampai jumlah itu tidak ditandai, siapa pun yang membayar; yang melewatinya tampil sebagai di atas kesepakatan, tanpa membandingkan kelebihannya dengan ambang. **Ini mengubah rancangan** 06 §6.2 dan §6.3 dan bagian S2-4 (sebelumnya: hanya tagihan di atas ambang, dan kelebihannya dibandingkan dengan ambang). Cicilan tetap tidak ditandai (D-44). Akibatnya atas pagar (pasangan yang diam atas usulan pos rutin; beban satu ketukan per pos) dinilai `security-privacy` sebelum checkpoint S3.

**Arahan pemilik produk tentang tampilan (11 November 2026).** Atas permintaannya, lead membuat satu sketsa kasar halaman depan (HTML, di luar repositori, bukan keluaran tahap mana pun) sebagai referensi isi. Tanggapannya atas tampilannya: "keliatan banget AI SLOP nya kaya gak ada seni tersendiri, standar app atau layout profesional, pallet warna font ... harusnya kita membangun kesan untuk aplikasi kita sendiri". Untuk `ux-designer` di dokumen 08 (S4), sebagai syarat: aplikasi punya identitas visual sendiri (warna, tipografi, ikon, ilustrasi, gerak, nada bahasa) dan tidak boleh tampak seperti templat umum; identitas itu dibangun di dalam komponen native tiap platform (brief §10.1), bukan dengan menggantinya. Usulan lead, belum ia jawab: di S4 disajikan dua atau tiga arah tampilan yang berbeda untuk ia pilih, dan ia boleh mengirim contoh aplikasi yang ia sukai. Sketsa kasar itu **bukan** acuan tampilan. Dua arah tampilan yang kemudian dibuat lead tanpa referensi ("Nila dan Kunyit", "Struk") juga ia tolak ("ai slop semua") dan **tidak dipakai**. Ia meminta referensi nyata dicari lebih dulu (ia menunjuk galeri Dribbble bertanda finance-app sebagai contoh arah) dan menetapkan syarat: responsif di ponsel Android, iPhone, iPad, tablet Android, dan ponsel lipat termasuk layar ganda seperti Duo; mode gelap dan terang; UI percakapan (diskusi pada permintaan dan pendamping AI) dikerjakan dengan sungguh-sungguh. Riset referensi desain dipesan lebih awal kepada `researcher` (dokumen 18, bagian desain platform), sebagai bahan `ux-designer` di S4. Sesudah riset itu lead membuat satu contoh halaman depan Android selebar 384 dp untuk Galaxy A55 (lebar itu hitungan lead, **belum terverifikasi** di perangkat), bergaya Material 3, hitam dengan satu aksen hijau limau, huruf Manrope, mode terang dan gelap. Tanggapannya: "oke sudah bagus untuk gambaran mah, tapi nanti layout tata letak tetep di s4 yaa supaya enak nyaman di semua menu dan fiturnya". Jadi contoh itu **gambaran yang ia terima sebagai arah kasar, bukan keputusan desain**; tata letak, navigasi, dan sistem visual diputuskan di S4 oleh `ux-designer` untuk semua menu dan fitur. Kanvasnya di luar repositori: https://claude.ai/artifact/PbC4BvowPG1GDzwqJZQ99N (artboard "Beranda Android, Galaxy A55"; hanya pemiliknya yang dapat membuka).

**Pemeriksaan `security-privacy` atas aturan yang berubah sesudah checkpoint S2 (hanya temuan; SP-32 sampai SP-44).** Tidak ada berkas temuan `security-privacy` di `docs/`; ringkasan ini satu-satunya catatannya sampai dokumen 10 ditulis. Tidak ada temuan yang menghalangi. Semua usulan perubahan **[Usulan]**.

| ID | Pokok | Kapan | Usulan perubahan |
|---|---|---|---|
| SP-32 | Jatah sebagai batas membuat menyepakati jatah merugikan dalam dua hal: jatah di bawah atau sama dengan ambang hanya menambah tanda; yang tanpa jatah dapat mengambil berkali-kali di bawah ambang tanpa tanda (ambang 500 ribu, lima kali 400 ribu: 0 tanda tanpa jatah, 3 tanda dengan jatah 1 juta, 5 tanda dengan jatah 300 ribu). SP-30 berbalik | Sebelum S3 | Jatah yang tidak di atas ambang (di bawah atau sama dengan ambang), termasuk 0, hanya dapat diusulkan pemiliknya sendiri (K01). Batas "sama dengan" adalah pengetatan lead sesudah tinjauan ulang `critic` (12 November 2026), supaya pagar dan jalan keluar seketika memakai batas yang sama; pertanyaan kepada pemilik produk berbunyi "nol atau di bawah ambang"; layar menyebut apa yang ditandai dengan dan tanpa jatah sebelum disetujui; peringatan baru di dokumen 17 |
| SP-33 | T-29 pilihan 1 berlawanan dengan PG-14 tetapi terbatas (paling lama satu periode) | Sebelum S3 | `assessment_until` hanya bila jatah yang diakhiri di atas ambang; selain itu berlaku seketika; layar menyebut tanggalnya; PG-14 dirumuskan ulang di dokumen 10 dengan ADR |
| SP-34 | Jatah 0 adalah ambang nol untuk satu orang | Sebelum S3; keputusan pemilik produk | Pagar SP-32 dan jalan keluar seketika SP-33 |
| SP-35 | Diam atau menolak atas usulan pos rutin: tidak ada daya halang baru | Dokumen 10; satu perbaikan sebelum S3 | Status kesepakatan pos hanya tampil di layar pos itu; kata layar "belum sepakat" |
| SP-36 | Beban satu ketukan jatuh menurut peran (pembayar tagihan mengusulkan hampir semuanya) | Sebelum S3 | N-11 pos rutin paling banyak satu pemberitahuan langsung per pengirim per hari; setuju sekaligus menampilkan tiap jumlah, tanpa pilihan terisi, dan tidak mencakup pos di atas ambang |
| SP-37 | Tagihan kecil berkesepakatan lebih rugi daripada yang tanpa kesepakatan; insentif menghindari kesepakatan | Sebelum S3; P-20 disampaikan ke pemilik produk | Tanda dasar `routine_amount` hanya masuk ringkasan, tidak pernah pemberitahuan langsung; layar menawarkan satu ketukan untuk mengusulkan jumlah baru |
| SP-38 | Konfirmasi susulan dapat bergeser menjadi "persetujuan susulan diharapkan" | Sebelum S3 | Tanpa pengingat otomatis; paling banyak satu pengingat ulang; tidak dapat diajukan lagi sesudah "belum sepakat"; tampil sebagai diajukan sesudah belanja dengan kedua tanggal; tanda yang tersisa tidak pernah berbunyi "belum disetujui"; otorisasi ditutup di transaksi persetujuan dengan batas sebesar entri |
| SP-39 | Jam tenang: tidak ditemukan orakel di aplikasi | Dokumen 10 dan 09 | Yang tertahan dilepas sebagai satu notifikasi; tidak ada push apa pun selama jam tenang |
| SP-40 | Dompet campuran: empat jalan tersisa | (b) sebelum S3 | (b) T-03 tanpa bawaan terisi: pemegang memilih bersama, pribadi, atau dibagi; (a) orang yang melihat dompet dapat menghitung bagian pribadi, dinyatakan terus terang; (c) urutan sinkron tidak boleh membocorkan adanya entri pribadi (09); (d) "server tidak menerima C" tidak melindungi dari operator |
| SP-41 | Ambang, S2-4 pilihan A | Dokumen 10 | Dapat diterima untuk pilot dengan dua syarat (menyepakati ambang terpisah dari menerima undangan; keduanya diberi tahu ambang tidak dapat diubah sendirian); sebelum tahap publik disarankan boleh menarik diri sendirian, berlaku periode berikutnya; ditanyakan di S4 |
| SP-42 | Pemberitahuan kepada istri tentang akses operator | Sebelum pilot dimulai | Ia membaca sendiri teksnya di aplikasi sebelum mengisi data pribadi; pengakuannya dicatat bertanggal; draf teks ada di laporan agen dan dibawa ke dokumen 10; kalimat janji operator hanya bila pemilik produk bersedia |
| SP-43 | Mekanisme yang masih melawan arah "mencatat, bukan menilai" | (a), (b) sebelum S3 | (a) pemberitahuan atas dasar jatah dan dasar pos rutin hanya ringkasan; (b) butir bertanda di ringkasan dilebur menjadi satu baris tanpa hitungan; sisanya syarat kata-kata untuk dokumen 08 |
| SP-44 | Daftar uji PG-14 belum memuat jatah pertama, kesepakatan pos rutin, tautan ke pos rutin | Sebelum S3 (kecil) | Perluas daftar uji; pemeriksaan kecocokan kategori diusulkan di dokumen 10 |

**Tinjauan `critic` atas revisi pasca-checkpoint S2 (12 November 2026): belum siap.** 2 menghalangi (B1: dua kutipan di 06 §10 yang bukan kata-kata pemilik produk; B2: bagian ambang S2-4 dinaikkan menjadi "diterima atas saran lead"), 15 perlu diperbaiki, 7 minor. Aritmetika dan tanda yang ia hitung ulang dengan penalaran benar. Dikerjakan dalam satu putaran perbaikan bersama SP-32 sampai SP-44 dan empat jawaban di bawah.

**Dua kata-kata pemilik produk yang belum tercatat di atas, ditambahkan agar dapat dilacak.** (1) Atas S2-1 butir (iv), kata-katanya: "ini masuknya uang masuk aja ga sih? misal kalau public kan kita gatau ada yang gajinya harian, bulanan, bahkan sebulan gada pemasukan coba pikirkan". Inilah dasar kalimat "irama penghasilan tidak diandaikan" di 07. (2) Kalimat "tenang, bukan aplikasi trading atau kripto" di pesan tugas riset desain adalah **kata-kata lead**, bukan pemilik produk; ia tidak pernah menyebut trading atau kripto.

**Empat pertanyaan dari SP-32 sampai SP-43 diajukan lead kepada pemilik produk (12 November 2026):** (1) jatah nol atau di bawah ambang hanya boleh diusulkan pemiliknya sendiri dan boleh ia lepas seketika; (2) tanda karena lewat jatah atau lewat jumlah pos rutin hanya tampil di catatan dan ringkasan, tanpa pemberitahuan langsung; (3) konfirmasi susulan tanpa pengingat otomatis, satu pengingat ulang, tidak dapat diminta lagi sesudah "belum sepakat"; (4) selisih dompet campuran dipilih pemegang setiap kali, tanpa bawaan. Peringatan SP-32 (yang tanpa jatah lebih jarang ditandai; satu orang boleh berjatah sementara pasangannya tidak) dibacakan. Jawabannya sejauh ini, atas (1): "yaa kalau ngerasa kurang maka boleh ajukan misal suami ke istri atau sebalikanya". Bacaan lead: ia menegaskan jalan kenaikan (keputusan 17), sama ke dua arah; ia **tidak** menyatakan pagar "jatah di bawah ambang hanya diusulkan pemiliknya". Pagar itu, beserta jalan keluar seketika bagi jatah di bawah ambang (SP-33), diterapkan sebagai **[Usulan]** `security-privacy` yang disampaikan kepadanya dalam pertanyaan (1) dan tidak ia bantah. Butir (2), (3), (4): **diterima atas saran lead** ("ya bener gas"). Akibatnya atas yang dibacakan di checkpoint: pemberitahuan langsung "paling banyak satu per hari" (D-35) kini hanya berlaku atas dasar ambang; tanda karena lewat jatah dan lewat jumlah pos rutin tidak pernah menjadi pemberitahuan langsung. T-03: bawaan "selisih ke bagian bersama" **dicabut**; pemegang memilih setiap kali. Diterapkan dalam satu putaran perbaikan bersama temuan `critic`.

**Pertanyaan pemilik produk tentang dana darurat dan impian (12 November 2026):** "dana darurat, impian/goals. udah?". Jawaban lead: sudah ada dalam bentuk sederhana di Tier 1 (07 §3.9: tujuan berjenis `savings` dengan nama dan target; menabung adalah alokasi atau pemindahan, bukan pengeluaran; contoh "dana darurat" target 10 juta); tanggal target, tonggak, prioritas, proyeksi, dan kontribusi per orang adalah `GOL` Tier 3 (brief §1.1, §5H). Untuk S3 (dokumen 04, 05): tempat tujuan tabungan di halaman depan dan apakah ada bagian `GOL` yang ia ingin naikkan ke pilot ditanyakan di checkpoint S3.

**Pernyataan arah kedua (11 November 2026), kata-katanya:** "karena semua butuh kesepakatan, setuju atau tidak, tujuan bukan dipakai untuk menyuruh nyuruh, bukan untuk membebankan, tapi mengingat dan mencatat". Bacaan lead **[Usulan]**: permintaan adalah ajakan bersepakat antara dua orang setara, bukan perintah dan bukan izin; aplikasi yang mengingatkan, bukan pasangan yang menagih; tugas antar-anggota adalah permintaan tolong yang boleh diterima atau tidak. Usulan kata-kata layar yang lahir darinya (untuk dokumen 08, S4; belum ia kukuhkan): tanpa kata "melanggar", "salah", "curang", "perintah", "izin"; "ditolak" dan "keberatan" ditinjau (calon: "belum sepakat", "perlu dibahas"). Alasan yang diberikan lead kepadanya adalah penalaran lead, **bukan hasil riset**; bila dipakai di dokumen, `researcher` mencari sumbernya atau labelnya **[Asumsi]**.

## Keputusan pemilik produk di checkpoint S1

Catatan lengkap, termasuk mana yang diputuskan sendiri, mana yang diterima lewat "lanjut", dan mana yang belum dikukuhkan, ada di [17-open-questions-assumptions.md](17-open-questions-assumptions.md) bagian "Keputusan pemilik produk di checkpoint S1". Ringkasnya:

- Tiap uang masuk dicatat dengan sumber, punya siapa, dan pemegangnya; penerima otomatis menjadi pemegang dan pencatat.
- Gaji berbawaan dana bersama (dapat diubah saat entri); pemberian kepada satu orang berbawaan dana pribadi orang itu.
- Dana pribadi ada dan saldonya tertutup; bila dipakai untuk keperluan rumah tangga, pengeluarannya terlihat oleh keduanya dan tidak diganti kecuali ditandai "minta diganti".
- Kedua pasangan setara penuh; pelaku tiap tindakan terlihat; tiap hal punya tingkat visibilitas.
- Belanja biasa dicatat setelah terjadi tanpa persetujuan; rencana diajukan dulu dan dapat dirundingkan; pemakaian besar tak terencana di atas satu ambang permintaan yang sama bagi keduanya diajukan dulu.
- Permintaan tanpa tanggapan tetap menunggu; pemohon dapat mengingatkan lagi.
- Catatan lama tidak pernah disunting; selisih dicatat sebagai selisih kas; entri terlambat memakai tanggal aslinya.
- Anggaran per kategori buatan rumah tangga, bulanan (perubahan cakupan atas brief §5C disetujui). Hari mulai periode ditentukan rumah tangga sendiri (keputusan 11 November 2026, di bawah), menggantikan rumusan awal "mulai tanggal gajian".
- Semua daftar buatan pengguna; menghapus butir beriwayat berarti mengarsipkannya.
- AI: preferensi Claude atau ChatGPT, server aplikasi di Indonesia; pilihan penyedia **ditunda ke checkpoint S4**. Waktu melibatkan advokat dan catatan tentang apa yang diterima pemilik sebagai operator selama pilot **belum dikukuhkan** dan disajikan lagi di S4.
- Keputusan pemilik di checkpoint berkedudukan setara brief untuk hal yang diputuskannya; rumusan brief dan `CLAUDE.md` yang sepadan diusulkan lead di checkpoint S2.

## Keputusan dan pernyataan pemilik produk selama S2 (11 November 2026)

Catatan lengkap dengan kata-katanya ada di dokumen 17 (keputusan 11 sampai 17). Ringkasnya:

- Uang masuk, termasuk gaji, boleh diarahkan ke dana pribadi saat dicatat; catatannya (jumlah penuh dan bagian pribadi) terlihat pasangan; saldo dan pemakaian dana pribadi tetap tertutup.
- Serah terima tercatat seketika dan bertanda "belum dikonfirmasi" sampai penerima menekan terima; penerima boleh berkeberatan kapan saja; sama ke dua arah.
- Pengingat ulang paling sering satu kali per 6 jam per permintaan dan menghormati jam tenang penerima.
- Kedua pasangan berpenghasilan. Uang masuk pemilik produk bisa beberapa kali sebulan (gaji sampai dua kali, pekerjaan lepas sampai lima kali), ke rekening bank atau dompet digital; lalu ditransfer ke istri atau ditarik tunai dan dibagikan.
- *Jatah pribadi*: per orang, dijumlahkan per bulan, angkanya disepakati berdua dan saling dikonfirmasi; kenaikan diajukan dengan alasan.
- Hari mulai periode anggaran ditentukan rumah tangga sendiri.

**Bukan keputusan pemilik produk saat S2 ditulis** (semuanya kemudian dijawab di checkpoint S2; lihat "Jawaban checkpoint S2" di atas): di dalam jatah tidak ada tanda (D-48); perlakuan atas bagian di luar jatah (S2-1 iii); hanya penerima penghasilan yang boleh mengarahkannya (S2-1 i); hadiah untuk satu orang tidak terlihat (S2-1 ii); alasan kenaikan wajib atau tidak (S2-1 v); hari mulai 1 sampai 28 dan penggantiannya disepakati berdua (putusan lead, D-49); seluruh daftar D.

## S2 — Fondasi domain

### Dokumen

| Dokumen | Pemilik | Status | Catatan |
|---|---|---|---|
| [07-financial-domain.md](07-financial-domain.md) | `finance-domain` | ditinjau, direvisi sesudah checkpoint S2 | Ledger entri bertipe; satu buku per pemilik uang; dana pribadi di luar garis kendali; INV-01 sampai INV-30; CTH-01 sampai CTH-41; butir terbuka T-01 sampai T-29 |
| [06-activity-approval-architecture.md](06-activity-approval-architecture.md) | `architect` | ditinjau | Alur domain eksplisit di atas enam primitif; INV-A, PG, N, V, AT; butir terbuka P-01 sampai P-17 |
| [adr/](adr/) ADR-001 sampai ADR-010 | `architect` | diterima di checkpoint S2 | Kedudukan tiap ADR di 06 §15.7 dan dokumen 17; ADR-007 sementara sampai advokat menjawab (S2-8; L-04, L-20) |
| [02-product-principles.md](02-product-principles.md) | `product-analyst` | ditinjau | PRN-01 sampai PRN-18, tiap prinsip dengan cara ujinya; ketegangan KT |
| [03-personas-household-models.md](03-personas-household-models.md) | `product-analyst` | ditinjau | PSN-01 sampai PSN-07; sepuluh struktur rumah tangga; KB-01 sampai KB-25 |
| [00-glossary.md](00-glossary.md) | lead; bagian 7 `finance-domain` | ditinjau (versi S2) | Istilah keuangan dan persetujuan difinalkan; usulan yang menunggu checkpoint berlabel **[Usulan]** |
| [17-open-questions-assumptions.md](17-open-questions-assumptions.md) | `product-analyst` | ditinjau, direvisi sesudah checkpoint S2 | Keputusan 11 sampai 25, penegasan 1 sampai 3; hasil checkpoint S2; SP-32 sampai SP-44; Q sampai Q-132; daftar "Dibawa ke S3" |

### Tinjauan

| Tinjauan | Oleh | Hasil |
|---|---|---|
| Pemeriksaan awal dokumen 06 (hanya temuan) | `security-privacy` | SP-01 sampai SP-22: 2 menghalangi, 12 sebelum S3, 8 catatan untuk dokumen 10. Semua yang menghalangi dan sebelum-S3 dikerjakan pemilik dokumennya |
| Tinjauan S2 | `critic` | Belum siap: 1 menghalangi (usulan tentang jatah pribadi ditulis sebagai keputusan pemilik), 7 perlu diperbaiki, 9 minor. Semuanya dikerjakan |
| Pemeriksaan aturan jatah pribadi (hanya temuan) | `security-privacy` | SP-24 sampai SP-31: 2 menghalangi (mengakhiri jatah pasangan sendirian; bentuk alasan kenaikan), 3 sebelum S3, 3 catatan. Yang menghalangi dan sebelum-S3 dikerjakan |

**Yang tidak ditinjau ulang:** perbaikan atas temuan `critic` dan atas SP-24 sampai SP-31 diperiksa lead terhadap laporan pemilik dokumen dan lewat hitung ulang angka, **tidak** ditinjau `critic` untuk kedua kalinya. `architect` melaporkan membaca ulang dokumen 06 secara utuh hanya sebagian pada tiap revisi; `critic` membacanya utuh satu kali, sebelum putaran perbaikan terakhir. Tinjauan konsistensi berikutnya ada di S3 dan S6.

**Angka yang dihitung ulang lead dengan skrip (di luar repositori), setelah revisi terakhir dokumen 07:** 42 tabel saldo, 213 baris: jumlah berjalan tiap kolom, rumus `UANG` dan `KB`, kesinambungan rantai CTH-01 sampai CTH-36; tabel samping (alokasi, anggaran, reservasi, rekonsiliasi, CTH-02, CTH-17b, CTH-37 A/B/C, arus kas); CTH-03b, varian perubahan syarat CTH-22, CTH-34, CTH-38, CTH-39 termasuk tabel jatah; contoh urutan tiba pada W-13. Semuanya cocok. `critic` memeriksa dengan penalaran bahwa angka per kategori CTH-37 berasal dari entri satuannya dan bahwa baris entri seimbang per buku; itu tidak diperiksa skrip.

### Revisi sesudah checkpoint S2 (11 dan 12 November 2026)

| Tinjauan | Oleh | Hasil |
|---|---|---|
| Pemeriksaan aturan yang berubah (hanya temuan) | `security-privacy` | SP-32 sampai SP-44: tidak ada yang menghalangi; delapan sebelum S3, sisanya untuk dokumen 10 dan 09. Yang sebelum S3 dikerjakan pemilik dokumennya |
| Tinjauan revisi | `critic` | Belum siap: 2 menghalangi (dua kutipan karangan di 06 §10; bagian ambang S2-4 dinaikkan kedudukannya), 15 perlu diperbaiki, 7 minor. Semuanya dikerjakan |
| Tinjauan ulang putaran perbaikan | `critic` | Tidak ada yang menghalangi; 5 perlu diperbaiki (dua atribusi terlalu luas atas konfirmasi susulan, sisa label Q-53, batas pagar K01 tidak sama dengan jalan keluar seketika, catatan hitung ulang), 7 minor. Semuanya dikerjakan; tidak ditinjau ketiga kalinya |

**Asal dua kutipan karangan (B1).** `architect` menyatakan ia menyusunnya dari ringkasan percakapan sesudah konteks dipadatkan, tanpa membaca ulang berkas ini. Keduanya dan semua yang diturunkan darinya dicabut. Pelajaran proses: kutipan pemilik produk hanya diambil dari berkas ini atau dokumen 17, tidak pernah dari ringkasan.

**Angka yang dihitung ulang lead dengan skrip (di luar repositori), sesudah putaran perbaikan terakhir dokumen 07 (12 November 2026):** jumlah berjalan tiap kolom pada 44 tabel saldo (90 baris "Sesudah"); tidak ada baris saldo lama yang berubah atau terhapus menurut `git diff`. CTH-40: ketiga pilihan selisih Rp40.000, isi dompet Rp890.000. CTH-41: saldo, alokasi, dan tepat satu tanda (air 310000 atas kesepakatan 250000). Jatah sebagai batas: contoh urutan tiba (satu lawan dua tanda) dan varian CTH-39. Rumus `N` untuk entri tertaut: CTH-21 sampai CTH-24 termasuk varian 800000 atas batas 200000, CTH-33, CTH-34, dan kasus 06 §12 (700000, 400000, 600000 atas batas 1000000: hanya yang ketiga bertanda). 07 §10.3 butir 11 (lima kali 400000: 0, 3, 5 tanda). SP-33 (jatah 500000 diakhiri: seketika; jatah 1000000 diakhiri: tetap dasar sampai akhir periode). Semuanya cocok. **Tidak** diperiksa skrip kali ini: rumus kolom `UANG` dan `KB` dan kesinambungan rantai CTH-01 sampai CTH-36 (dihitung sebelum checkpoint; barisnya tidak tersentuh revisi), dan kasus K01 serta kenaikan ambang di 06 §12 (bukan hitungan uang; diperiksa `critic` dengan penalaran).

### Keputusan

ADR-001 sampai ADR-010 diusulkan di S2 dan diterima di checkpoint S2. Keputusan proses lead saat S2 ditulis: pembayaran cicilan tidak memerlukan kesepakatan rutin (pembelian secara cicilan dibandingkan dengan ambang satu kali); penggantian hari mulai periode disepakati berdua (menunggu pemilik produk); pertanyaan checkpoint dibaca dari dokumen 17 saja.

### Dibawa ke S3 dan sesudahnya

- Syarat layar untuk dokumen 05 dan 08: "bagikan ke beberapa penerima" dan "langsung dipakai" sebagai satu alur layar; tanggal pencatatan tampil hanya bila berbeda dari tanggal kejadian; sisa jatah pasangan hanya bila dibuka; teks layar kunci umum; catatan penilaian tanpa bahasa pelanggaran.
- Untuk dokumen 10 (S4): catatan SP-15 sampai SP-22, SP-29 sampai SP-31; asas yang diusulkan `security-privacy` (seorang dewasa boleh sendirian mengurangi pengawasan aplikasi atau menambah pengawasan atas dirinya; menambah pengawasan atas dewasa lain butuh persetujuan orang itu); mengakhiri kesepakatan rutin atas pembayaran yang dilakukan pasangan (03 §7).
- Untuk dokumen 09 (S4): tanda jatah bergantung pada urutan penerimaan server (INV-A37); perlu urutan tersimpan dan penilaian berurutan di dalam transaksi tulis.
- Untuk advokat: L-22 sampai L-25 ditambah butir lama.
- Untuk `product-analyst` saat menyebarkan jawaban checkpoint: di dokumen 17, tandai K-27 butir (3), (4), (9), (12) selesai (dokumen 06 sudah diperbaiki) dan catat pada Q-95 bahwa daftar sumber penghasilan diubah lewat kesepakatan rumah tangga oleh kedua anggota dewasa **[Usulan]**.
- Pemilik produk: uji langsung dua atau tiga aplikasi pesaing sebelum S3 selesai; keputusan S4 (penyedia AI, waktu advokat, catatan operator, build iOS tanpa Mac).


Bukti baru setelah checkpoint: pemilik memberi salinan teks PP 33/2026; dokumen 18 ditulis ulang darinya (label **[Perlu validasi] — salinan teks**; pindaian resmi belum terbuka). Menurut bacaan peneliti atas Pasal 160 ayat (3), ukuran transfer lintas batas adalah kedudukan penerima, sehingga "diproses di Indonesia" oleh penyedia asing belum tentu cukup; riset S4 (R-05) harus memeriksa kedudukan hukum penyedia AI dan hosting.

## S1 — Framing

### Dokumen

| Dokumen | Pemilik | Status | Catatan |
|---|---|---|---|
| [README.md](README.md) | lead | ditinjau | Kerangka, tautan, konvensi; skema pengenal ditambah `Q`, `A`, `R`, `K`, `L` |
| [00-glossary.md](00-glossary.md) | lead; bagian 7 `finance-domain` | ditinjau (versi pertama) | Istilah keuangan adalah definisi kerja; titik bertanda *(diputuskan di S2)* difinalkan bersama dokumen 07 |
| [18-market-research.md](18-market-research.md) | `researcher` | ditinjau | Pasar dan regulasi terisi; integrasi, desain platform, dan teknologi masih placeholder (S3/S4) |
| [17-open-questions-assumptions.md](17-open-questions-assumptions.md) | `product-analyst` | ditinjau; direvisi setelah checkpoint (revisi belum ditinjau `critic`) | 58 pertanyaan, asumsi, riset yang perlu divalidasi, pertentangan sumber, butir tinjauan hukum |

### Tinjauan

`critic` meninjau dua kali.

| Tinjauan | Menghalangi | Perlu diperbaiki | Minor | Putusan |
|---|---|---|---|---|
| Pertama | 1 | 16 | 19 | Belum siap |
| Kedua | 1 | 4 | 16 | Belum siap |

Semua temuan menghalangi dan perlu-diperbaiki dari kedua tinjauan sudah dikerjakan pemilik dokumennya. Perbaikan atas tinjauan kedua diperiksa lead terhadap teks temuan, **tidak** ditinjau `critic` untuk ketiga kalinya. Minor yang sengaja dibiarkan ada di "Dibawa ke S2" di bawah.

Angka yang dihitung ulang lead dengan skrip (di luar repositori): contoh talangan Rp100.000 lalu penggantian, untuk dua letak dana pribadi terhadap garis kendali rumah tangga. Hasil: kekayaan bersih turun tepat Rp100.000, pengeluaran tercatat satu kali, pada kedua pilihan.

### Keputusan

Belum ada ADR. S1 tidak memutuskan apa pun atas nama pemilik produk; semua default berstatus **[Usulan]** sampai checkpoint.

Keputusan proses yang diambil lead:

- Riset S1 dibatasi pada pasar dan regulasi, sesuai brief §0.3.
- Label **[Terverifikasi]** hanya untuk sumber primer yang benar-benar dibuka; sumber sekunder saja atau cuplikan pencarian berlabel **[Perlu validasi]** (dokumen 18 §0.2).
- Di glosarium, *tempat uang* (`holding`) dipakai agar *akun* hanya berarti identitas login; peran pemilik bernama `household owner` agar tidak tertukar dengan *pemilik uang*.

### Asumsi yang dipakai

Daftar lengkap di dokumen 17 bagian 3. Yang paling berpengaruh:

- Semua `[DEFAULT]` brief §1 berlaku.
- Pencatatan pengeluaran tidak pernah dihalangi status permintaan (A-20).
- Desain dan rencana pilot diperlakukan seolah pengecualian "kegiatan pribadi atau rumah tangga" UU PDP **tidak** berlaku, sampai advokat menjawab (A-13, dokumen 18 §2.1).
- Anak dimodelkan sebagai profil tanggungan tanpa login; login anak adalah keputusan tahap publik.

### Menunggu pemilik produk

Sepuluh pertanyaan checkpoint S1 (dokumen 17 §1.1), urut prioritas: Q-09, Q-01, Q-02, Q-04, Q-03, Q-08, Q-05, Q-06, Q-29, Q-10. Empat keputusan tambahan yang dibutuhkan sebelum S2 (§1.2): Q-07, Q-49, Q-25, Q-26. "Lanjut pakai default" mencakup keempat belas butir dan catatan pertanyaan 10.

Tindakan yang hanya bisa dilakukan pemilik (§1.3):

- Uji langsung dua atau tiga aplikasi pesaing, paling lambat S3.
- Mengunduh teks resmi PP 33/2026 bila `researcher` tetap gagal membukanya, paling lambat S4.

### Dibawa ke S2

Untuk `finance-domain` (dokumen 07 dan glosarium bagian 7):

- Letak dana pribadi terhadap garis kendali rumah tangga (Q-02), dan apakah angka utama "hanya dana bersama" diberi nama sendiri atau definisi *uang yang ada* disesuaikan (K-13).
- Definisi *realisasi* masih "jumlah entri pengeluaran", bertentangan dengan cicilan sebagai pelunasan yang dihitung terhadap anggaran; bunga cicilan menurunkan kekayaan bersih tanpa entri pengeluaran (K-11, Q-16).
- Pelepasan *hak ganti* bila dana pribadi di luar garis menaikkan kekayaan bersih tanpa pemasukan (temuan minor N-14 tinjauan kedua).
- Perlakuan kelebihan yang dibayar dari dana bersama (Q-13) dan selisih kas positif (Q-05).
- Q-43 sampai Q-48 belum punya default.

Untuk `architect` (dokumen 06 dan ADR):

- ADR atas tafsir "append-only" untuk penghapusan data anggota: inti moneter entri tetap dan dipseudonimkan, rincian pribadi dihapus (Q-22, K-10). UU PDP Pasal 43–44 tidak memberi pengecualian bagi keutuhan catatan bersama.
- Konfirmasi serah terima (Q-19), jenis permintaan untuk pemakaian dana bersama yang bukan pengeluaran (Q-03), alasan penolakan (Q-50).

Untuk lead:

- Rumusan *ambang permintaan* di glosarium masih "pengeluaran tidak rutin"; diperluas menjadi "pemakaian dana bersama" bila pemilik menerima default Q-03 (K-12).
- "Angka utama" dan "baris kedua" (Q-02) belum menjadi istilah glosarium; ditetapkan setelah dokumen 07.

Untuk `security-privacy` (pemeriksaan awal dokumen 06 sebelum checkpoint S2, lalu dokumen 10 di S4):

- Utang antar-pasangan yang menyentuh dana pribadi pasti terlihat oleh keduanya; dua kebocoran dana pribadi yang dinyatakan di Q-02.
- Keberadaan tempat uang pribadi yang tidak dibagikan: usulan tanpa penanda bagi pasangan.
- Pilihan (iii) pada Q-10 (DPIA ringkas dan perjanjian prosesor selama fase pembangunan) menyiratkan pekerjaan yang belum dijadwalkan.

Untuk `researcher` (diperluas di S3/S4):

- L-16: apakah pemrosesan di wilayah Indonesia oleh penyedia asing, atau hosting di luar negeri, terhitung transfer lintas batas.
- R-05: kelayakan penyedia AI yang memproses di Indonesia dan model yang di-host sendiri.
- Teks resmi PP 33/2026; pendaftaran PSE lingkup privat; label privasi Apple; angka Epstein dkk. dicocokkan dengan PDF.
- Kategori komunikasi keluarga hanya dua halaman vendor (celah terhadap brief §15).
