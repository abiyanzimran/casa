# 02 — Prinsip Produk

| | |
|---|---|
| Pemilik dokumen | `product-analyst` |
| Tahap | S2 (Fondasi domain) |
| Status | **Putaran perbaikan 12 November 2026** (temuan `critic` S12): bagian 7 butir 2 tidak lagi menyebut batas pengingat "belum dijawab"; bagian 0, PRN-05, dan bagian 7 butir 6 diselaraskan dengan S2-5 dan empat jawaban 12 November 2026. Tidak ada kalimat prinsip yang berubah. **Dicocokkan dengan jawaban checkpoint S2 dan dua pernyataan arah pemilik produk (11 November 2026):** bagian 0 memuat hasil checkpoint S2 dan kedua pernyataan arah; PRN-10 dan PRN-18 dipertajam (masing-masing satu pemeriksaan baru); PRN-08 mendapat satu pemeriksaan dan ketegangan baru KT-12; PRN-11, PRN-13, PRN-16 hanya mengutip; label yang usang di PRN-09, PRN-12, PRN-13, PRN-16, KT-1, KT-7, dan bagian 7 disesuaikan. Tidak ada prinsip baru. **Catatan sebelumnya:** Draf untuk tinjauan `critic`. Diperbarui dengan pernyataan pemilik produk 11 November 2026 bahwa kedua pasangan berpenghasilan (PRN-09 "Biaya yang diterima"; KT-9; bagian 7 butir 1 dan 2). **Putaran perbaikan tinjauan `critic` S2** (M-6, M-7, M-8): bagian 0 memuat keputusan 11 sampai 17 dan dua pernyataan pemilik produk; PRN-05 tidak lagi mengandaikan tombol tunda; PRN-09 diselaraskan dengan ADR-010 butir 8 dan 07 INV-27; PRN-13 dan KT-1 menyebut jatah pribadi; PRN-16 dan KT-7 diselaraskan dengan ADR-007 |
| Tanggal | 11 November 2026 |
| Dibaca sebelum menulis | `CLAUDE.md`; `family-os-brief.md` §1, §1.1, §2–§5, §8, §10.1, §11–§13, §16; `ENGINEERING-STANDARDS.md`; [00-glossary.md](00-glossary.md); [_progress.md](_progress.md); [07-financial-domain.md](07-financial-domain.md) §1–§6, §9–§10; [06-activity-approval-architecture.md](06-activity-approval-architecture.md); ADR-001 sampai ADR-010; [17-open-questions-assumptions.md](17-open-questions-assumptions.md) (keputusan pemilik produk 1 sampai 17, D-48 sampai D-54, dan daftar); ADR-007 butir 1 sampai 8 dan ADR-010 butir 8 untuk putaran perbaikan; [18-market-research.md](18-market-research.md) §1, §2.2 |

Dokumen ini memuat delapan belas prinsip Family OS. Tiap prinsip ditulis agar **dapat gagal dalam tinjauan**: ada pemeriksaan konkret yang lulus atau tidak lulus, ada contoh pelanggaran, dan ada biaya yang diterima. Prinsip yang tidak dapat gagal bukan prinsip; daftar kalimat yang ditolak karena itu ada di bagian 6.

## 0. Cara membaca

**Kedudukan.** Dokumen ini tidak membuat aturan uang atau aturan persetujuan. Aturan uang kanonis di dokumen 07 dan aturan alur di dokumen 06; prinsip di sini menunjuk invarian (`INV-nn`, `INV-Ann`), pagar (`PG-nn`), dan larangan visibilitas (`V-n`, `AT-n`) yang sudah ada di sana, dan tidak membuat padanannya. Bila sebuah prinsip dan dokumen 06 atau 07 tampak berbeda, yang berlaku untuk uang dan alur adalah dokumen 06 dan 07; perbedaannya dilaporkan ke lead.

**Keputusan pemilik produk** dikutip sebagai "keputusan pemilik nomor *n*", dengan nomor dokumen 17. Nomor 1 sampai 10 diambil di checkpoint S1; nomor 11 sampai 17 diambil sesudahnya, pada 11 November 2026 (dokumen 17, kelompok kedua dan ketiga). Rujukan "keputusan pemilik, 11 November 2026" yang lebih lama di dokumen ini berarti keputusan 11 sampai 13.

| # | Keputusan |
|---|---|
| 11 | Uang masuk boleh diarahkan ke dana pribadi saat dicatat, termasuk gaji, tetapi catatannya terlihat pasangan (jumlah gajinya dan berapa yang masuk dana pribadi). Yang tertutup hanya saldo dana pribadi dan pemakaiannya sesudah itu |
| 12 | Serah terima dicatat seketika dan bertanda "belum dikonfirmasi" sampai penerima menerimanya; penerima dapat mengajukan keberatan kapan saja; berlaku sama ke dua arah |
| 13 | Pengingat ulang oleh pemohon paling banyak satu kali per 6 jam per permintaan, dan menghormati jam tenang penerima |
| 14 | *Jatah pribadi* adalah satu jumlah per orang per bulan, dijumlahkan sepanjang bulan; jumlahnya dapat diatur |
| 15 | Hari mulai *periode anggaran* ditentukan rumah tangga (menyempurnakan keputusan 8, "mulai pada tanggal gajian"). "Bawaan tanggal 1, berlaku mulai periode berikutnya" adalah bacaan lead yang dikukuhkan kepadanya |
| 16 | Jumlah jatah disepakati berdua: satu mengusulkan, yang lain mengonfirmasi |
| 17 | Kenaikan jatah dapat diajukan dengan alasan dan dikonfirmasi pasangan |

Pada hari yang sama pemilik produk juga menyatakan dua fakta tentang rumah tangganya, yang bukan keputusan rancangan: kedua pasangan berpenghasilan; dan penghasilannya sendiri datang beberapa kali (gaji dapat dua kali, pekerjaan lepas sampai lima kali) ke rekening bank atau e-wallet.

**Yang bukan keputusan pemilik produk** dan tidak dikutip sebagai keputusan: di dalam jatah tanpa tanda dan tanpa pemberitahuan, dan setiap pemindahan uang bersama ke uang pribadi terhitung terhadap jatah (bawaan dokumen 17 D-48, **[Usulan]**); siapa yang boleh mengganti hari mulai periode (putusan lead, D-49); perlakuan atas pengambilan yang melewati jatah (dokumen 17, S2-1 butir iii; terbuka saat kalimat ini ditulis, kini keputusan 20, di bawah).

**Sesudah checkpoint S2** (dokumen 17, bagian "Keputusan pemilik produk di checkpoint S2"). Yang ia putuskan sendiri dan dipakai dokumen ini: hanya penerima penghasilan yang mengarahkan penghasilannya (keputusan 18); hadiah untuk satu orang tidak terlihat pasangan (19); jatah pribadi adalah batas, dan melewatinya ditandai berapa pun kelebihannya (20); persetujuan tidak dapat ditarik sendirian (21); tagihan rutin disepakati dulu (22); bacaan append-only diterima sebagai keputusan sementara (24). D-01 sampai D-55 ia terima lewat "lanjut"; bawaan itu berlaku tetapi **bukan** pernyataannya dan tidak dikutip di sini sebagai keputusan. Batas pengingat (tiga pengingat, plafon harian, tunda senyap) dan bawaan jam tenang 22.00 sampai 06.00 ia terima atas saran lead; angkanya **[Asumsi]**.

**Dua pernyataan arah pemilik produk** (11 November 2026; kata-katanya, dari `_progress.md`). Dikutip di dokumen ini sebagai "arah 1" dan "arah 2".

| # | Kata-kata pemilik produk | Bacaan lead **[Usulan]**, belum ia kukuhkan kata per kata | Prinsip yang sudah memuatnya | Yang dipertajam |
|---|---|---|---|---|
| Arah 1 | "yang penting perselisihan dan tidak bisa di otak atik, aplikasi ini bukan menunjukan siapa yang salah, tapi mencatat apa yang perlu di catat, liat perkembangan, status, bukan dibuat untuk perselisihan tapi untuk menyimpan data ('semua ada catatanya ada datanya')" | (1) Catatan tidak dapat diubah diam-diam. (2) Aplikasi mencatat dan menunjukkan keadaan, tidak menilai siapa yang salah. (3) Tanda, keberatan, dan selisih memakai bahasa keterangan, bukan bahasa pelanggaran | (1) PRN-16, PRN-13. (2) PRN-18 ("tidak memutus siapa yang benar"), PRN-11 (tanpa hitungan atau peringkat per orang), PRN-08 (tanpa notifikasi yang menilai orang) | (3): sebelumnya hanya ada sebagai kalimat di "Biaya yang diterima" PRN-18, yang tidak dapat gagal dalam tinjauan. Kini satu pemeriksaan di PRN-18 |
| Arah 2 | "karena semua butuh kesepakatan, setuju atau tidak, tujuan bukan dipakai untuk menyuruh nyuruh, bukan untuk membebankan, tapi mengingat dan mencatat" | Permintaan adalah ajakan bersepakat antara dua orang setara, bukan perintah dan bukan izin; aplikasi yang mengingatkan, bukan pasangan yang menagih; tugas antar-anggota adalah permintaan tolong yang boleh diterima atau tidak | PRN-09 (dua arah, simetris), PRN-10 (hanya orangnya sendiri yang menyetujui, menerima, menandai selesai; diam tidak memutuskan), PRN-08 (pengingat berbatas) | "Tidak membebankan": PRN-10 belum menyatakan bahwa tidak ada objek yang membebani orang dewasa lain sebelum ia sendiri menerimanya. Kini satu kalimat dan satu pemeriksaan di PRN-10. "Aplikasi yang mengingatkan" bergesekan dengan keputusan pemilik nomor 6: KT-12 |

Prinsip tidak menetapkan kata-kata layar; itu milik dokumen 08. Daftar kata yang diusulkan lead (dokumen 17, Q-119) belum dikukuhkan pemilik produk.

Dokumen 06 dan 07 versi revisi sudah memuat ketujuh keputusan itu. Karena nomor bagian dokumen 06 masih dapat bergeser, rujukan ke bagian yang terdampak menyebut **nama bagian** selain nomornya: 06 §6.3 "Kesepakatan rumah tangga" (bagian "Jatah pribadi"), 06 §6.4 "Uang masuk yang diarahkan ke dana pribadi", 06 §7 "Konfirmasi custody", 06 §10 "Pagar terhadap kendali koersif" (PG-06 "Pengingat berbatas"), 07 §3.1 "Uang masuk", 07 §3.4 "Penyisihan dan kontribusi", 07 §5.3 "Kebocoran yang tersisa", 07 §6.1 "Anggaran".

**Label bukti** mengikuti [README](README.md#label-bukti). Pernyataan prinsip adalah **[Usulan]** `product-analyst` kecuali bagian yang dikutip dari brief, keputusan pemilik produk, atau `ENGINEERING-STANDARDS.md`. Angka sasaran yang bukan dari `ENGINEERING-STANDARDS.md` berlabel **[Usulan]** dan dikukuhkan di S3 (dokumen 05) atau S4 (dokumen 08).

**Pengenal.** Brief §16 tidak punya pengenal untuk prinsip. Dokumen ini memakai `PRN-NN`, dan tiga singkatan untuk hal yang dirujuk berulang oleh dokumen 03, 04, dan 05. Lead diminta menambahkannya ke README.

| Kode | Arti | Sumber |
|---|---|---|
| `PRN-NN` | Prinsip produk | Dokumen ini |
| `MSL-1` | Uang tidak terlihat | Brief §1.1 butir 1 |
| `MSL-2` | Custody tidak jelas | Brief §1.1 butir 2 |
| `MSL-3` | Arus uang yang sebenarnya (gaji, cicilan, tagihan, kiriman ibu, sisa dipegang istri, pinjam kembali, bantuan keluarga besar) tidak tercatat | Brief §1.1 butir 3 |
| `MSL-4` | Belanja dan jalan-jalan tidak direncanakan | Brief §1.1 butir 4 |
| `MSL-5` | Mencatat itu beban; pencatat utama pelupa | Brief §1.1 butir 5 |
| `MSL-6` | Masalah sekunder: suasana hati pasangan, masalah rumah, rutinitas, "masak apa", tindak lanjut sakit, siapa membayar tagihan apa | Brief §1.1 butir 6 |
| `MSL-7` | Kehangatan butuh usaha | Brief §1.1 butir 7 |
| `KS-1` … `KS-5` | Lima kriteria sukses setelah tiga bulan pilot, urut seperti di brief | Brief §1.1 |
| `PAGAR` | Permintaan berjalan dua arah, dan berbagi suasana hati selalu pilihan yang membagikan | Brief §1.1 "Guardrail" |
| `TUJUAN` | Satu tujuan di `CLAUDE.md` | `CLAUDE.md` |

**Bentuk tiap prinsip.** Pernyataan; yang dilayani; cara diuji; contoh pelanggaran; biaya yang diterima. "Cara diuji" membedakan pemeriksaan otomatis pada fase pembangunan dari pemeriksaan dalam tinjauan dokumen atau tinjauan layar.

## Ringkasan

| ID | Prinsip | Kelompok |
|---|---|---|
| PRN-01 | Tahu tanpa bertanya | Produk dan desain |
| PRN-02 | Setiap fitur membayar usaha berulangnya | Produk dan desain |
| PRN-03 | Fondasi tahap publik tidak dipotong | Produk dan desain |
| PRN-04 | Jujur tentang bukti | Produk dan desain |
| PRN-05 | Mencatat dalam hitungan detik, dan terlambat itu sah | UX |
| PRN-06 | Pencatatan tidak pernah dihalangi | UX |
| PRN-07 | Asli di tiap platform, dan tenang | UX |
| PRN-08 | Notifikasi berbatas dan aman di layar kunci | UX |
| PRN-09 | Simetris karena bentuknya | Privasi dan keamanan |
| PRN-10 | Tidak ada yang bertindak atas nama orang dewasa lain, dan diam tidak memutuskan apa pun | Privasi dan keamanan |
| PRN-11 | Pelaku terlihat pada catatannya, tidak pernah menjadi alat pantau | Privasi dan keamanan |
| PRN-12 | Pribadi secara bawaan; berbagi adalah pilihan pemiliknya | Privasi dan keamanan |
| PRN-13 | Dana bersama tidak dapat disembunyikan | Privasi dan keamanan |
| PRN-14 | AI membantu menyusun kata; orang yang membaca, memutuskan, dan mengirim | Privasi dan keamanan |
| PRN-15 | Setiap rupiah terhitung satu kali | Integritas keuangan |
| PRN-16 | Catatan hanya bertambah | Integritas keuangan |
| PRN-17 | Rencana bukan uang | Integritas keuangan |
| PRN-18 | Yang belum pasti ditandai, tidak disembunyikan dan tidak ditebak | Integritas keuangan |

---

## 1. Produk dan desain

### PRN-01 — Tahu tanpa bertanya

**Pernyataan.** Tiap pasangan dapat menjawab "uang kita sisa berapa dan siapa yang memegangnya", serta "apa yang sedang menunggu saya dan apa yang sedang menunggu pasangan saya", paling jauh satu ketukan dari layar pertama **[Usulan]**, dan jawabannya sama bagi keduanya.

**Melayani.** `MSL-1`, `MSL-2`; `KS-1`; `TUJUAN`.

**Cara diuji.**

- Otomatis: *uang yang ada* dihitung dari baris (07 INV-05); angka rumah tangga identik bagi kedua pelihat (07 INV-20); jumlah baris custody sama dengan angka utama, termasuk bagian "belum dikonfirmasi" dan "dikeberatani" (06 INV-A24).
- Tinjauan layar (S4, lalu tiap rilis): dari aplikasi tertutup, penguji menemukan angka utama dan rincian per pemegang tanpa membuka menu; cold start sesuai `ENGINEERING-STANDARDS.md` §3.
- Pilot: pada akhir bulan pertama, kedua, dan ketiga, tiap pasangan diminta menjawab kedua pertanyaan itu dari aplikasi tanpa bertanya, lalu dicocokkan dengan menghitung uang (rekonsiliasi kas, 07 §3.10). Prinsip gagal bila salah satu tidak dapat menjawab, atau jawabannya berbeda.

**Contoh pelanggaran.** Angka utama yang memuat anggaran, target, atau dana pribadi; angka yang berbeda di ponsel suami dan ponsel istri; custody yang baru terlihat setelah tiga layar; daftar permintaan yang hanya terlihat oleh pemohonnya.

**Biaya yang diterima.** "Berapa harta keluarga seluruhnya" tidak dijawab di layar mana pun, karena dana pribadi berada di luar angka rumah tangga (07 §2.2; ADR-003). Angka utama hanya sebenar catatannya: di antara dua kali menghitung uang ia bisa meleset, dan itu ditampilkan, tidak ditutupi (PRN-18). Prinsip ini tidak menjawab "bagaimana keadaan pasangan saya" kecuali pasangan memilih membagikannya (PRN-12).

### PRN-02 — Setiap fitur membayar usaha berulangnya

**Pernyataan.** Setiap fitur menyebut masalah brief §1.1 yang diselesaikannya dan usaha berulang yang dimintanya: tindakan apa, oleh siapa, seberapa sering. Fitur yang tidak menyebut masalah dipindahkan ke Tier 3 atau ke "tidak dibangun" beserta alasannya. Fitur yang nilainya bergantung pada pemeliharaan manual harian, selain pencatatan uang itu sendiri, harus berubah bentuk (datanya menjadi akibat samping tindakan lain) atau ditunda.

**Melayani.** `MSL-5`; `KS-2`; brief §3 "Every feature must justify the recurring effort it asks of users".

**Cara diuji.**

- Tinjauan dokumen (S3): tiap baris fitur di dokumen 04 punya kolom "masalah (`MSL-n`)" dan "usaha berulang" yang terisi; baris kosong adalah temuan menghalangi. Tiap kebutuhan Tier 1 di dokumen 05 menunjuk fiturnya.
- Tinjauan dokumen: jumlah tindakan berulang per orang per minggu dijumlahkan untuk seluruh MVP di dokumen 04. Angka itu dibandingkan dengan keadaan sekarang, yaitu hanya belanja dapur yang dicatat (brief §1.1). Penambahan fitur yang menaikkan jumlah itu tanpa menyebut pertanyaan apa yang tidak perlu lagi ditanyakan adalah temuan.
- Pilot: `KS-2` (keduanya mencatat tanpa diingatkan) adalah ujian akhirnya.

**Contoh pelanggaran.** Persediaan rumah tangga yang stoknya diperbarui manual tiap hari (brief §5F sendiri meragukannya); isian suasana hati wajib tiap pagi; kategori, catatan, dan struk sebagai isian wajib pada setiap belanja; layar pengaturan izin yang harus diisi sebelum aplikasi dapat dipakai.

**Biaya yang diterima.** Beberapa hal yang diminta brief ditunda atau dibuat lebih miskin data: `INV` dan `MEAL` tetap Tier 3 walaupun "masak apa" adalah keluhan nyata (`MSL-6`); laporan lebih kasar karena isian opsional sering kosong. Konfirmasi custody menambah satu ketukan bagi penerima pada tiap serah terima yang dicatat pemberi (06 §7 "Konfirmasi custody"); usaha itu dibenarkan di ketegangan KT-8.

### PRN-03 — Fondasi tahap publik tidak dipotong

**Pernyataan.** Pemotongan cakupan berlaku atas fitur dan layar, tidak atas empat fondasi model data dan model akses: isolasi rumah tangga; akun terpisah dari anggota; lima konsep brief §4.3 disimpan terpisah; dan tidak ada aturan yang membaca hubungan suami-istri. Keempatnya ada sejak Tier 1 walaupun pilot hanya punya satu rumah tangga dengan dua anggota dewasa.

**Melayani.** `TUJUAN` tahap kedua; brief §1.1 "the Tier 1 data model and access model must not block Tier 3"; `CLAUDE.md` "The one goal", paragraf ketiga.

**Cara diuji.**

- Otomatis: test bahwa anggota satu rumah tangga tidak dapat membaca atau mengubah data rumah tangga lain (`ENGINEERING-STANDARDS.md` §4); test yang sama dengan satu akun yang menjadi anggota dua rumah tangga.
- Pemeriksaan statis: kode alur tidak membaca `relationship` (06 PG-01); tidak ada kolom, tabel, atau aturan yang bernama atau bermakna "suami" atau "istri"; setiap objek membawa `household_id` (06 §14).
- Tinjauan skema (S4): kelima baris 06 §3.1 ditulis saat rumah tangga dibuat, sehingga preset lain kelak hanya mengisi nilai berbeda.

**Contoh pelanggaran.** Tabel rumah tangga dengan kolom `husband_id` dan `wife_id`; saldo yang melekat pada akun, bukan pada tempat uang milik satu rumah tangga; izin yang diturunkan dari hubungan ("orang tua boleh melihat semua").

**Biaya yang diterima.** Sebagian pekerjaan model tidak punya pemakai di pilot (tabel keanggotaan, baris wewenang persetujuan, peran). Yang **tidak** dibangun di pilot: layar pengaturan izin, preset selain "dua dewasa setara", pengalih rumah tangga (Q-30; dokumen 03 bagian 4). Fondasi berarti model, bukan layar.

### PRN-04 — Jujur tentang bukti

**Pernyataan.** Setiap klaim yang tidak sepele di dokumen membawa satu dari empat label bukti, dan label tidak pernah dinaikkan saat dikutip ulang. Angka tanpa sumber primer tidak dikutip. Hasil pilot dilaporkan sebagaimana adanya: `KS-5` adalah pernyataan pasangan, bukan metrik.

**Melayani.** Brief §0.1 "Honesty about evidence", §16; `CLAUDE.md` "Evidence".

**Cara diuji.**

- Tinjauan `critic` tiap tahap: setiap angka tentang pasar, pengguna, atau regulasi punya label dan sumber; label **[Terverifikasi]** hanya untuk sumber primer yang dibuka (dokumen 18 §0.2); angka blog yang dilarang dokumen 18 §1.3.5 tidak muncul di dokumen mana pun.
- Pemeriksaan khusus: angka Epstein dkk. tidak dikutip di luar dokumen 18 sebelum dicocokkan dengan PDF naskahnya (dokumen 18 §1.3.5, §6.4), dan tidak pernah disajikan sebagai pembagian berjumlah 100%.
- Persona di dokumen 03: setiap atribut rumah tangga pilot yang tidak dinyatakan brief atau pemilik produk berlabel **[Asumsi]**.

**Contoh pelanggaran.** "82% pengguna berhenti karena input manual"; "tidak ada pesaing yang punya fitur custody" (yang diperiksa hanya halaman pemasaran, dokumen 18 §1.2); persona dengan usia, pekerjaan, dan kebiasaan yang dikarang.

**Biaya yang diterima.** Dokumen memuat lebih banyak kalimat bersyarat. Positioning "custody dan permintaan dua arah belum ada di pasar" belum boleh dipakai di materi publik sampai dua atau tiga aplikasi pesaing diuji langsung (dokumen 18 §1.4). Persona pilot tampak tipis, karena memang hanya itu yang diketahui.

---

## 2. UX

### PRN-05 — Mencatat dalam hitungan detik, dan terlambat itu sah

**Pernyataan.** Mencatat satu belanja memerlukan paling banyak tiga ketukan setelah layar entri terbuka dan langsung tersimpan di perangkat (`ENGINEERING-STANDARDS.md` §3). Menanggapi sesuatu yang ditunggu orang lain (menyetujui, menolak, mengonfirmasi serah terima) memerlukan paling banyak dua ketukan dari notifikasinya **[Usulan]**; tunda senyap dipertahankan (S2-5 pilihan A, diterima atas saran lead), jadi sasaran yang sama berlaku atasnya. Lupa adalah keadaan normal: entri terlambat dengan tanggal aslinya adalah jalur biasa, *periode anggaran* (hari mulainya ditentukan rumah tangga, keputusan pemilik nomor 15) tidak pernah ditutup, dan yang tetap terlupa tertangkap sebagai selisih kas.

**Melayani.** `MSL-5`; `KS-2`; brief §1 "Recording something must take seconds".

**Cara diuji.**

- Otomatis dan terukur di perangkat pilot: sasaran `ENGINEERING-STANDARDS.md` §3 untuk pencatatan belanja; UI test menghitung ketukan untuk menyetujui, menolak, dan mengonfirmasi dari notifikasi, dan untuk menunda (S2-5 pilihan A mempertahankannya) (06 §7.2 sudah menyebut konfirmasi satu ketukan).
- Otomatis: entri bertanggal kejadian 60 hari lalu diterima selama tidak lebih awal dari saldo awal tempat uangnya (07 INV-22), dan laporan periode itu menandai perubahannya (07 §3.11; ADR-006); sesudah rekonsiliasi, saldo tercatat sama dengan saldo terhitung (07 INV-16).
- Tinjauan kebutuhan (S3): isian wajib pada entri belanja didaftar satu per satu di dokumen 05; setiap isian wajib selain jumlah dan tempat uang memerlukan alasan tertulis.

**Contoh pelanggaran.** Kategori, catatan, dan foto struk wajib; bulan lalu dikunci sehingga belanja yang baru teringat tidak dapat dicatat; persetujuan yang baru dapat diberikan setelah membuka tiga layar.

**Biaya yang diterima.** Laporan bulan lampau dapat berubah. Sebagian uang keluar tampil sebagai "selisih hitung yang belum dirinci" dan tidak terhitung terhadap anggaran sampai dirinci (07 §3.10). Persetujuan dua ketukan memberi sedikit ruang membaca; yang menjaganya adalah persetujuan atas versi syarat tertentu (06 INV-A06). Kecepatan dicapai tanpa AI (PRN-14), walaupun dua pesaing lokal menjual entri berbantuan AI (dokumen 18 §1.2, isi halaman vendor **[Terverifikasi]**; Q-31).

### PRN-06 — Pencatatan tidak pernah dihalangi

**Pernyataan.** Entri yang sah menurut dokumen 07 selalu tersimpan. Status permintaan, ambang permintaan, kesepakatan, keberatan, saldo yang tidak cukup, AI yang gagal, dan jaringan yang putus tidak pernah menolaknya. Aturan rumah tangga menandai; ia tidak mengunci.

**Melayani.** `MSL-1`, `MSL-5`; `KS-1`; keputusan pemilik nomor 5 "recording is never blocked"; A-20.

**Cara diuji.** Otomatis: 06 INV-A11, INV-A20, PG-10, aturan S-7 (pencatatan offline); saldo tercatat negatif, *belum dialokasikan* negatif, dan saldo tersedia negatif sengaja bukan invarian (07 §4, paragraf penutup); semua kriteria penerimaan Tier 1 lulus dengan AI dimatikan (06 PG-12).

**Contoh pelanggaran.** "Belanja ditolak: belum disetujui"; "saldo tidak cukup"; entri di atas ambang ditahan sampai pasangan menjawab (pilihan yang ditolak di Q-54); layar entri yang menunggu saran kategori dari AI.

**Biaya yang diterima.** Ambang permintaan adalah kesepakatan yang pelanggarannya terlihat, bukan kunci (ADR-009; 07 §10.3 butir 4). `KS-4` karena itu bergantung pada kebiasaan pasangan, bukan pada aturan sistem; lihat ketegangan KT-5. Saldo negatif dan tanda "di atas ambang" akan muncul di layar.

### PRN-07 — Asli di tiap platform, dan tenang

**Pernyataan.** Di iPhone aplikasi memakai komponen, navigasi, dan gerak sistem iOS; di Android, milik Android. Yang dibagi kedua platform adalah warna merek, ikon, kata-kata, dan arsitektur informasi, bukan layar yang identik (brief §10.1). Layar pertama menjawab PRN-01, bukan menampilkan dashboard; dan tidak ada istilah akuntansi di layar mana pun. Rinciannya milik dokumen 08 (S4); di sini hanya batas produknya.

**Melayani.** `TUJUAN` (usaha serendah mungkin: pola yang sudah dikenal tidak perlu dipelajari); brief §1 "WhatsApp-class", §3 "must not feel like an enterprise dashboard", §10.1.

**Cara diuji.**

- Keputusan stack (S4): kerangka yang menggambar widget sendiri harus membuktikan kesetaraan dengan komponen sistem terkini di kedua platform, atau ditolak (brief §9.1).
- Tinjauan layar: `ux-designer` memeriksa setiap layar yang dibangun terhadap dokumen 08 dan pedoman platformnya (`CLAUDE.md`, tabel tim pembangunan).
- Otomatis: cold start, frame, dan ukuran unduhan sesuai `ENGINEERING-STANDARDS.md` §3; pemeriksaan teks antarmuka terhadap daftar kata terlarang yang diturunkan dari kolom kiri 07 §1.4 (entri, baris, kewajiban, piutang, otorisasi, reservasi, dan seterusnya).

**Contoh pelanggaran.** Bottom sheet buatan sendiri yang sama persis di kedua platform; layar pertama berisi enam grafik; kata "kewajiban" atau "rekonsiliasi" pada tombol.

**Biaya yang diterima.** Dua implementasi antarmuka, atau satu kerangka dengan batasan ketat; keputusannya di S4. Identitas visual lebih tidak mencolok daripada aplikasi yang menggambar semuanya sendiri. Sembilan kata sifat di brief §10.1 (hangat, premium, tenang, dan seterusnya) tidak dijadikan prinsip karena tidak dapat gagal; dokumen 08 menerjemahkannya menjadi pilihan yang dapat diperiksa.

### PRN-08 — Notifikasi berbatas dan aman di layar kunci

**Pernyataan.** Notifikasi hanya untuk hal yang perlu ditanggapi atau perlu diketahui penerimanya, dan tiap jenisnya terdaftar dengan prioritas. Pengingat berbatas jumlah dan waktunya, dan tidak menembus jam tenang. Layar kunci tidak pernah memuat jumlah uang atau rincian kepedulian. Aplikasi tidak mengirim pengingat "kamu belum mencatat hari ini" **[Usulan]**: `KS-2` meminta pencatatan tanpa diingatkan.

**Melayani.** `MSL-4`, `MSL-6`; `KS-4`; `PAGAR`; brief §5N "Minimize notification fatigue".

**Cara diuji.**

- Otomatis: batas pengingat menurut 06 PG-06 "Pengingat berbatas" dan INV-A30, dengan angka keputusan pemilik, 11 November 2026 (pengingat ulang kedua dalam 6 jam atas permintaan yang sama ditolak; pengingat pada jam tenang ditahan); isi layar kunci sesuai `ENGINEERING-STANDARDS.md` §4 dan 06 §11.1.
- Tinjauan: jenis notifikasi baru yang tidak ada di tabel peristiwa 06 §11.2 (atau padanan Tier 2-nya di dokumen 05) beserta prioritasnya adalah temuan.
- Tinjauan: tidak ada notifikasi yang isinya menilai seseorang ("sudah tiga hari belum mencatat", "pasanganmu belum menjawab").
- Otomatis (arah 2): isi pengingat ulang adalah kalimat baku aplikasi; pemohon tidak dapat menulis teks bebas atau jumlah uang di dalamnya (06 §10 PG-06; dokumen 17 D-29). Pengingat dengan isi dari pemohon adalah temuan.

**Contoh pelanggaran.** "Istri belanja Rp250.000 di minimarket" di layar kunci; pengingat harian untuk mencatat; pemohon dapat mengirim pengingat ulang sepuluh kali dalam satu jam.

**Biaya yang diterima.** Hal penting dapat tertahan sampai jam tenang berakhir; apakah masalah rumah yang "mendesak" boleh menembusnya belum diputuskan (06 §13.2). Tanpa pengingat mencatat, lupa lebih sering terjadi; yang menangkapnya adalah rekonsiliasi kas, bukan omelan. Brief §5N meminta preferensi notifikasi yang terperinci; untuk pilot diusulkan pengaturan per prioritas saja, bukan per peristiwa, karena layar pengaturan yang panjang adalah usaha tanpa pemakai (PRN-02) **[Usulan]**, diputuskan di S3.

---

## 3. Privasi dan keamanan

### PRN-09 — Simetris karena bentuknya

**Pernyataan.** Tidak ada aturan, bawaan, batas, atau tampilan yang bergantung pada hubungan atau label peran keluarga (suami, istri, kepala keluarga), jenis kelamin, siapa yang berpenghasilan atau berpenghasilan lebih besar, siapa pemegang uang, atau siapa pembuat rumah tangga. Sebuah aturan boleh bergantung pada **penghasilan siapa** sebuah catatan (penerima penghasilannya, `earner`), asalkan berlaku sama bagi anggota dewasa mana pun (ADR-010 butir 8); contohnya 07 INV-27 (keputusan pemilik nomor 18, checkpoint S2): bagian pribadi sebuah penghasilan hanya dapat diarahkan oleh penerima penghasilan itu sendiri. Permintaan berjalan dua arah; ambang permintaan adalah satu angka yang sama bagi semua anggota dewasa; dan *jatah pribadi* disepakati dengan cara yang sama bagi keduanya, walaupun jumlahnya per orang (keputusan pemilik nomor 14 dan 16).

**Melayani.** `PAGAR`; `KS-4`; brief §4.3 kalimat terakhir, §5D kalimat terakhir; keputusan pemilik nomor 4.

**Cara diuji.** Otomatis: menukar kedua anggota dewasa pada rangkaian perintah acak, **termasuk** `earner` pada tiap catatan penghasilan dan penerima tiap jatah, menghasilkan state yang sama dengan anggota tertukar (06 INV-A13); pemohon tidak pernah menyetujui permintaannya sendiri (06 INV-A04, INV-A05); pemeriksaan statis 06 PG-01. Tinjauan dokumen: aturan di dokumen 05 ditulis untuk "anggota dewasa", "pemohon", "pemberi persetujuan", "pemegang", "penerima penghasilan"; kata "suami" atau "istri" di dalam aturan adalah temuan; aturan yang menyebut `earner` hanya boleh membandingkannya dengan pelaku catatan yang sama (`actor = earner`), tidak dengan besarnya penghasilan dan tidak dengan anggota tertentu.

**Contoh pelanggaran.** Preset "suami menyetujui belanja istri di atas Rp500.000"; ambang per orang; ambang lebih tinggi bagi yang berpenghasilan lebih besar; layar yang hanya ada bagi "pemegang uang belanja".

**Biaya yang diterima.** Rumah tangga yang menginginkan aturan tidak setara antara dua orang dewasa tidak dilayani preset pilot. Jumlah jatah pribadi yang berbeda antara dua orang adalah angka yang mereka sepakati berdua, bukan aturan yang membaca siapa orangnya. Aturan yang simetris tetap dapat berakibat tidak simetris. Kedua pasangan pilot berpenghasilan (pernyataan pemilik produk, 11 November 2026), sehingga penghasilan keduanya terlihat pasangannya menurut keputusan pemilik, 11 November 2026; tetapi jumlah yang terbaca lebih besar pada yang berpenghasilan lebih besar, dan tanda serta selisih lebih sering jatuh pada yang lebih banyak mencatat, membayar, dan memegang uang. Akibat semacam ini dinyatakan terbuka; lihat ketegangan KT-9.

### PRN-10 — Tidak ada yang bertindak atas nama orang dewasa lain, dan diam tidak memutuskan apa pun

**Pernyataan.** Setiap persetujuan, penerimaan, konfirmasi, dan penandaan selesai dilakukan oleh orang yang pernyataannya diminta. Tidak ada peran, tenggat, atau AI yang menggantikannya. Waktu hanya dapat mengakhiri permintaan sebagai kedaluwarsa; ia tidak pernah menjadikan sesuatu disetujui, ditolak, berlaku, atau dikonfirmasi. **Tidak ada objek yang membebani orang dewasa lain sebelum ia sendiri menerimanya:** permintaan, usulan kesepakatan, dan permintaan tindakan adalah ajakan; sampai dijawab, tidak ada tugas, batas, atau kewajiban yang berlaku atas orang yang diajak, dan menolaknya tidak wajib beralasan.

**Melayani.** `PAGAR`; brief §5D "cooperation, not coercive control", §8; keputusan pemilik nomor 6, 16, 21, 22; keputusan pemilik, 11 November 2026 (serah terima); arah 2 ("bukan dipakai untuk menyuruh nyuruh, bukan untuk membebankan").

**Cara diuji.** Otomatis: 06 PG-02, PG-04, PG-05, PG-11, PG-13; INV-A10, INV-A18, INV-A22, INV-A23, INV-A28 (memajukan jam sejauh apa pun tidak menghasilkan persetujuan, penolakan, atau konfirmasi); hanya penerima tugas yang dapat menandai tugasnya selesai (06 §13.1). Otomatis (arah 2): permintaan tindakan baru melahirkan tugas setelah penerimanya menerima (06 §13.1, T02), dan hanya penerimanya yang dapat menerima, menolak, atau melepasnya; jatah, ambang, dan jumlah pos rutin baru berlaku setelah disetujui orang kedua (06 INV-A18); penolakan diterima tanpa alasan (06 §4.4; §13.1 T03). Tinjauan dokumen 05 dan 08: objek yang tampil di daftar "tugasku" atau berlaku sebagai batas atas seseorang sebelum ia menerimanya adalah temuan.

**Contoh pelanggaran.** Disetujui otomatis setelah 48 jam tanpa jawaban; serah terima dianggap diterima setelah tiga hari; pemilik rumah tangga mengeluarkan pasangannya; pemohon menandai selesai tugas yang ia minta dari pasangannya; tugas yang langsung masuk daftar pasangan dengan tenggat sebelum ia menerimanya; jatah yang berlaku atas pasangan begitu diusulkan.

**Biaya yang diterima.** Permintaan dapat menunggu tanpa batas, dan serah terima dapat selamanya "belum dikonfirmasi". Aplikasi tidak memecah kebuntuan; pasangan yang menyelesaikannya. Pemohon tidak pernah tahu apakah permintaannya sudah dibaca (PRN-11).

### PRN-11 — Pelaku terlihat pada catatannya, tidak pernah menjadi alat pantau

**Pernyataan.** Siapa yang mencatat, membayar, menyetujui, dan mengoreksi tampil pada objeknya, sama bagi kedua pasangan. Tidak ada umpan kegiatan per orang, hitungan atau peringkat per orang, tanda "sudah dibaca", "terakhir aktif", atau lokasi perangkat.

**Melayani.** `KS-3` (siapa membayar apa); `PAGAR`; brief §8 "Abuse prevention is a first-class requirement"; keputusan pemilik nomor 4 ("bukan audit"); arah 1 ("bukan menunjukan siapa yang salah").

**Cara diuji.** Otomatis dan tinjauan skema: 06 §9.4 "Keterangan pelaku tanpa menjadi alat pantau", AT-1 sampai AT-7; PG-07 (tidak ada medan "dibaca pada"); PG-09 (tidak ada kueri atau layar yang mengelompokkan tanda, keberatan, koreksi, selisih kas, entri terlambat, pengingat, atau lama menanggapi per orang). Laporan "siapa membayar apa" lulus hanya bila memenuhi AT-7.

**Contoh pelanggaran.** "Aktivitas istri hari ini"; "terakhir dilihat 10.42"; papan peringkat tugas rumah (dijual sebagai fitur oleh satu aplikasi di dokumen 18 §1.1); rekap "siapa paling sering lupa mencatat".

**Biaya yang diterima.** Tidak ada gamifikasi, padahal sebagian keluarga mungkin menyukainya **[Asumsi]**. Laporan per orang lebih miskin: tanpa persentase perbandingan dan tanpa urutan menurut besarnya. Dua hal tidak dapat dihindari model dan dinyatakan terbuka (07 §10.3 butir 2 dan 3): tabel "siapa membayar apa" mengukur siapa yang memegang uang, bukan siapa yang berbelanja; dan selisih kas melekat pada pemegang tempat uangnya.

### PRN-12 — Pribadi secara bawaan; berbagi adalah pilihan pemiliknya

**Pernyataan.** Data kepedulian (suasana hati, sakit) dan dana pribadi tertutup secara bawaan. Yang membukanya hanya pemiliknya, secara eksplisit, terperinci, dan dapat dicabut seketika. Tidak ada peran yang dapat membuka data pribadi orang dewasa lain. Dari dana pribadi, yang terlihat pasangan hanyalah persentuhannya dengan dana bersama, yang daftarnya tertutup dan diberitahukan sebelum dicatat.

**Melayani.** `PAGAR` (berbagi suasana hati selalu pilihan yang membagikan); `KS-5`; brief §5K, §8 "Family membership never grants unlimited access to another adult's private data"; keputusan pemilik nomor 3.

**Cara diuji.**

- Otomatis: data kepedulian hanya terlihat oleh audiens pilihan pemiliknya dan pilihan itu dapat dicabut (`ENGINEERING-STANDARDS.md` §4); 06 V-4, V-5, V-7, PG-03; uji akses per peran yang hasilnya tidak bergantung pada peran; buku pribadi tidak mengubah angka rumah tangga (07 INV-20); hanya pemilik buku pribadi yang dapat menulis ke bukunya (07 INV-25).
- Tinjauan: daftar persentuhan dana pribadi yang terlihat pasangan sama dengan daftar di 07 §5.3 "Kebocoran yang tersisa" sesudah diselaraskan dengan keputusan pemilik, 11 November 2026; tidak ada jalan lain. Layar yang akan memperlihatkan sesuatu kepada pasangan mengatakannya sebelum disimpan (07 §5.3, kalimat pembuka).
- Tinjauan AI: tidak ada keluaran bagi satu pasangan yang memuat, mengisyaratkan, atau menyadur data yang ditutup pasangannya (brief §5O).

**Contoh pelanggaran.** Pemilik rumah tangga melihat riwayat suasana hati pasangannya; angka "total harta keluarga" yang memuat dana pribadi; sapaan pagi "sepertinya pasanganmu sedang sedih" ketika status itu tidak dibagikan; status suasana hati yang terbagi otomatis ke semua anggota.

**Biaya yang diterima.** Pasangan tidak dapat diberi tahu tentang kesusahan yang tidak dibagikan; `KS-5` bergantung pada kesediaan berbagi. Kebocoran yang tersisa nyata dan diterima: penyisihan, kontribusi, klaim, dan, menurut keputusan pemilik 11 November 2026, uang masuk yang diarahkan ke dana pribadi saat dicatat. Dari situ pasangan dapat memperkirakan batas kasar dana pribadi. Hadiah untuk satu orang, yang bawaannya memang dana pribadi, tidak terlihat (07 §3.1 "Uang masuk", §5.3 butir B5, INV-24): keputusan pemilik nomor 19, checkpoint S2.

### PRN-13 — Dana bersama tidak dapat disembunyikan

**Pernyataan.** Segala sesuatu yang menyangkut dana bersama (tempat uang, entri, otorisasi, anggaran, piutang, kewajiban, aset rumah tangga) terlihat oleh kedua pasangan dan sama persis bagi keduanya. Tidak ada tingkat visibilitas, peran, atau jenis permintaan yang dapat menutupnya, dan tidak ada perubahan yang terjadi diam-diam.

**Melayani.** `MSL-1`, `MSL-2`; `KS-1`, `KS-3`.

**Cara diuji.** Otomatis: 06 V-1 (dana bersama tidak dapat `private`), V-3 (visibilitas permintaan tidak menutup akibat uangnya), V-6 dan PG-08 (setiap perubahan menulis peristiwa dan diberitahukan), INV-A17 (tanda di atas ambang tidak dapat dipasang atau dilepas pencatat); V-8 dan INV-A29 bersama 07 INV-24 (penghasilan yang diarahkan ke dana pribadi tetap meninggalkan catatan yang terlihat kedua anggota dewasa; keputusan pemilik, 11 November 2026); 07 INV-20.

**Contoh pelanggaran.** Belanja dari dana bersama yang ditandai "pribadi"; permintaan tertutup yang otorisasinya ikut tertutup; tempat uang bersama yang diarsipkan saat masih bersaldo sehingga uangnya hilang dari angka utama (dicegah 07 INV-21).

**Biaya yang diterima.** Hadiah kejutan yang dibeli dengan dana bersama akan terlihat. Dua orang dewasa yang setara tidak punya ruang tertutup di dalam dana bersama; ruang tertutup itu adalah dana pribadi (PRN-12), dan batas antara keduanya adalah ketegangan KT-1. *Jatah pribadi* (keputusan pemilik nomor 14) tidak membuka ruang tertutup di dalam dana bersama: setiap pengambilan ke uang pribadi tetap terlihat, di dalam maupun di luar jatah; yang diubah jatah hanyalah apakah catatannya ditandai: di dalam jatah tidak (bawaan dokumen 17 D-48, diterima lewat "lanjut"), melewati jatah ya, berapa pun kelebihannya (keputusan pemilik nomor 20; 07 §3.4). Tanda itu sendiri tunduk pada PRN-18: keterangan atas kesepakatan, bukan penilaian atas orangnya.

### PRN-14 — AI membantu menyusun kata; orang yang membaca, memutuskan, dan mengirim

**Pernyataan.** Pendamping keluarga hanya menyusun kata. Ia tidak pernah mengirim pesan, menyetujui, mengonfirmasi, mengingatkan atas kehendaknya sendiri, atau mengubah catatan. Setiap pesan kepada pasangan dibaca dan dikirim pengirimnya sendiri. Tiap anggota ikut serta secara terpisah dan dapat mematikannya; data yang keluar ke penyedia AI adalah yang minimum, tanpa nama; biayanya berbatas keras. Seluruh aplikasi berfungsi tanpa AI. Kehangatan tidak pernah diukur, dihitung, atau dijadikan tekanan.

**Melayani.** `MSL-7`; `KS-5`; `TUJUAN` ("membantu mereka saling memperhatikan"); brief §5O.

**Cara diuji.**

- Otomatis: tipe pelaku peristiwa tidak memuat AI (06 PG-12); tidak ada jalur yang membuat catatan untuk pasangan tanpa perintah kirim dari anggota; jumlah pengingat hangat sama dengan jumlah pengingat yang memang dikirim aturan 06 §11 (AI hanya mengganti kalimatnya).
- Otomatis: semua kriteria penerimaan Tier 1 dan Tier 2 non-AI lulus dengan AI dimatikan, dan dengan batas biaya AI tercapai.
- Contract test atas muatan ke penyedia AI: tanpa nama, tanpa jumlah uang per transaksi, tanpa alasan teks bebas kecuali pemiliknya mengizinkan (brief §5O, keputusan 3 dan 4; A-08, A-09).
- Tinjauan: tidak ada hitungan, rentetan, atau teguran tentang perhatian ("sudah lima hari kamu tidak menyapa pasanganmu").

**Contoh pelanggaran.** Pesan "semangat ya" terkirim otomatis pukul tujuh pagi atas nama suami; AI memilih kategori lalu menyimpan entri; sapaan pagi dari aplikasi yang tampil seolah-olah ditulis pasangan; skor kehangatan mingguan.

**Biaya yang diterima.** Tidak ada entri berbantuan AI di MVP (Q-31), sehingga kecepatan harus datang dari desain. Kehangatan tetap bergantung pada orang yang menekan kirim; aplikasi hanya menurunkan usahanya. Draf tidak membawa penanda AI (keputusan pemilik di brief §5O nomor 2), jadi risiko terasa palsu nyata; rekomendasinya ditulis di dokumen 05 (Q-32) dan diputuskan pemilik produk di checkpoint S3. Pilihan penyedia dan akibat hukumnya ditunda ke S4 (Q-10).

---

## 4. Integritas keuangan

Keempat prinsip di bagian ini adalah bentuk produk dari aturan yang kanonis di dokumen 07 dan `ENGINEERING-STANDARDS.md` §5. Dokumen ini tidak menambah aturan; ia menyatakan apa yang tidak boleh dikorbankan demi kemudahan.

### PRN-15 — Setiap rupiah terhitung satu kali

**Pernyataan.** Uang yang berpindah antar-orang, antar-tempat, atau antar-pemilik tidak pernah terhitung dua kali di saldo, anggaran, atau laporan mana pun. Pemindahan antar-tempat uang dalam kendali rumah tangga tidak pernah menjadi pemasukan atau pengeluaran.

**Melayani.** `MSL-1`, `MSL-2`, `MSL-3`; `KS-1`, `KS-3`; brief §5C, §7.

**Cara diuji.** Otomatis: 07 INV-02 (keseimbangan per pemilik), INV-03, INV-06 (pemindahan tidak mengubah total), INV-08 (klaim dua sisi), INV-15 (satu baris terhadap paling banyak satu anggaran), INV-19 (jembatan laporan); setiap contoh terhitung dokumen 07 menjadi test dengan angka yang sama (`ENGINEERING-STANDARDS.md` §5). Sebelum tiap checkpoint, lead menghitung ulang contoh dengan skrip (`CLAUDE.md` "Computed, not read").

**Contoh pelanggaran.** Tarik tunai ATM tampil sebagai pengeluaran lalu pemasukan; penggantian talangan terhitung sebagai pengeluaran kedua; belanja dari dana pribadi untuk rumah tangga terhitung di kategori dan sekali lagi sebagai kontribusi.

**Biaya yang diterima.** Model internal lebih rumit daripada daftar transaksi polos (ADR-002). Pencatat harus memilih antara "belanja" dan "pindah uang"; pilihan yang salah diperbaiki dengan koreksi (PRN-16), bukan dicegah.

### PRN-16 — Catatan hanya bertambah

**Pernyataan.** Inti catatan uang (jumlah, tanggal, jenis, tempat uang, pelaku dan pihak sebagai pengenal internal) dan inti catatan alur (siapa mengajukan, menanggapi, menyetujui, mengonfirmasi, atau berkeberatan, atas versi apa, kapan) hanya bertambah: tidak pernah disunting atau dihapus. Kesalahan diperbaiki dengan entri pembalik atau penyesuaian yang membawa pelaku dan alasannya, dan catatan asalnya tetap terlihat. Yang boleh dihapus hanyalah daftar tertutup di ADR-007, yang diterima pemilik produk sebagai keputusan sementara sampai advokat menjawab (keputusan pemilik nomor 24; dokumen 17 S2-8): *rincian* (catatan teks, struk, isi diskusi, alasan teks bebas), buku pribadi seorang anggota atas permintaan penghapusannya, dan catatan kepedulian (`CARE`) oleh pemiliknya. Tidak satu pun penghapusan itu mengubah angka rumah tangga; penutupan klaim saat anggota keluar adalah entri baru yang terlihat, bukan penghapusan (07 §8.2).

**Melayani.** `MSL-2`; `KS-1`; brief §7; keputusan pemilik nomor 7; arah 1 ("tidak bisa di otak atik"; "semua ada catatanya ada datanya").

**Cara diuji.**

- Otomatis: 07 INV-09 (hash inti moneter entri lama identik sesudah perintah apa pun, kecuali penghapusan buku pribadi menurut ADR-007 butir 3), INV-18; 06 INV-A01 (dibaca atas inti alur, ADR-007 butir 6), INV-A27, PG-08.
- Otomatis: perintah hapus atas rincian, atas buku pribadi seorang anggota (sesudah klaimnya ditutup lewat `exit_closure`), atau atas catatan `CARE` tidak mengubah satu pun angka buku rumah tangga maupun hash inti alur.
- Tinjauan skema (S4): setiap jalur hapus menunjuk satu butir ADR-007; jalur hapus atau sunting yang tidak ada di daftar itu adalah temuan.

**Contoh pelanggaran.** Mengganti jumlah belanja minggu lalu di tempat; menghapus entri yang salah; menghapus catatan uang bersama atas permintaan penghapusan data; mengubah ambang atau visibilitas tanpa jejak.

**Biaya yang diterima.** Kesalahan ketik pada jumlah tetap terlihat di riwayat selamanya. Ketegangan dengan hak penghapusan data belum selesai: ADR-007 adalah tafsir sementara, diterima pemilik produk di checkpoint S2 (keputusan nomor 24), yang menunggu tinjauan hukum (L-04, L-20, L-22, L-23); lihat KT-7.

### PRN-17 — Rencana bukan uang

**Pernyataan.** Persetujuan, otorisasi, reservasi, anggaran, alokasi, target, tagihan berulang, dan jadwal cicilan tidak pernah mengubah saldo tercatat, *uang yang ada*, atau kekayaan bersih. Hanya entri atas peristiwa uang yang benar-benar terjadi yang mengubahnya.

**Melayani.** `MSL-1`, `MSL-4`; `KS-1`, `KS-4`; brief §5H ("money earmarked … from money that actually exists"), §7.

**Cara diuji.** Otomatis: 07 INV-12 (semua saldo identik sebelum dan sesudah perintah jenis itu), INV-11 (saldo tersedia), INV-14; 06 INV-A09.

**Contoh pelanggaran.** Menyetujui permintaan Rp100.000 menurunkan "uang kita" Rp100.000; target tabungan terhitung dalam harta bersih; tagihan yang jatuh tempo otomatis tercatat sebagai utang.

**Biaya yang diterima.** Layar harus menjelaskan dua pasang angka tanpa istilah akuntansi: tercatat lawan tersedia, terpakai lawan terikat. Sisa anggaran dapat positif ketika uangnya sudah tidak ada, dan sebaliknya (glosarium, *sisa anggaran*).

### PRN-18 — Yang belum pasti ditandai, tidak disembunyikan dan tidak ditebak

**Pernyataan.** Ketika catatan dan kenyataan belum cocok, atau dua orang belum sepakat, aplikasi memperlihatkannya sebagai tanda yang sama bagi keduanya: "belum dikonfirmasi", "dikeberatani", "selisih hitung yang belum dirinci", "berubah sejak terakhir dilihat", saldo negatif. Aplikasi tidak memperbaikinya sendiri, tidak menebak niat, dan tidak memutus siapa yang benar. **Setiap tanda adalah keterangan atas catatan dan atas kesepakatan yang menjadi dasarnya, bukan penilaian atas orangnya:** ia menyebut apa yang belum cocok atau kesepakatan mana yang terlewati, dan tidak pernah berakibat apa pun selain terlihat.

**Melayani.** `MSL-2`; `KS-1`; keputusan pemilik nomor 7; keputusan pemilik, 11 November 2026 (serah terima); arah 1 ("bukan menunjukan siapa yang salah, tapi mencatat ... liat perkembangan, status").

**Cara diuji.** Otomatis: 06 INV-A21, INV-A24 (konfirmasi custody), INV-A25, INV-A26 (keberatan tidak mengubah angka dan hanya ditutup pengajunya); 07 INV-16 (selisih yang belum dirinci), §3.11 (laporan lampau menandai perubahan). Tinjauan: tidak ada perintah berpelaku `system` yang mengubah atau merapikan angka. Otomatis (arah 1): tidak ada aturan yang menahan, mengurangi, atau membatasi sesuatu karena sebuah catatan bertanda (06 INV-A11, INV-A20; 07 INV-12). Tinjauan kata-kata (dokumen 05, 08, dan setiap teks layar atau notifikasi pada fase pembangunan): teks yang menyertai tanda, selisih, keberatan, atau penolakan (1) menyebut catatannya atau kesepakatannya, bukan orangnya, sebagai pokok kalimat; (2) tidak memakai kata yang menyatakan kesalahan orang, sekurangnya "melanggar", "pelanggaran", "salah", "curang" (**[Usulan]**; daftar akhirnya dokumen 08, dokumen 17 Q-119); (3) tidak menyajikan jumlah tanda sebagai ukuran, per orang maupun per pos (07 §10.3 butir 7 dan 10; 06 PG-09). Teks yang gagal salah satunya adalah temuan.

**Contoh pelanggaran.** Serah terima dikonfirmasi otomatis; selisih kas disebar otomatis ke kategori yang "paling mungkin"; alokasi dikurangi otomatis ketika uangnya kurang (07 §3.9 menolaknya); pemakaian yang dipecah di bawah ambang "dideteksi" sebagai pelanggaran (06 P-09); tanda berbunyi "Kamu melanggar jatah"; rekap "tiga kali di atas kesepakatan bulan ini".

**Biaya yang diterima.** Layar memperlihatkan ketidaksempurnaan, dan sebagian tanda melekat pada orang: selisih pada pemegang, "belum dikonfirmasi" pada penerima. Dokumen 08 memilih kata yang tidak menuduh (07 §10.3 butir 3), dan kini diperiksa dengan pemeriksaan kata-kata di atas. Kata yang netral tidak menghapus kenyataan bahwa tanda tetap melekat pada catatan seseorang; sejak jatah menjadi batas dan pos rutin punya batas, tanda juga lebih sering muncul (07 §10.3 butir 7 dan 10).

---

## 5. Pemetaan dan ketegangan

### 5.1 Prinsip terhadap kriteria sukses dan pagar

Tiap kriteria dan pagar punya sekurangnya dua prinsip; tiap prinsip produk muncul sekurangnya satu kali. PRN-03 dan PRN-04 melayani tahap publik dan mutu perencanaan, bukan satu kriteria pilot, dan dicatat di baris terakhir.

| Kriteria (brief §1.1) | Prinsip yang menopang | Prinsip yang membatasinya |
|---|---|---|
| `KS-1` Menjawab "sisa berapa dan siapa yang pegang" tanpa bertanya | PRN-01, PRN-06, PRN-13, PRN-15, PRN-16, PRN-17, PRN-18 | PRN-12 (dana pribadi di luar angka) |
| `KS-2` Keduanya mencatat tanpa diingatkan | PRN-02, PRN-05, PRN-06, PRN-07 | PRN-08 (tanpa pengingat mencatat) |
| `KS-3` Ringkasan bulanan: per kategori, per orang, siapa membayar apa, tabungan, aset | PRN-13, PRN-15, PRN-16, PRN-17 | PRN-11 (laporan per orang tanpa perbandingan) |
| `KS-4` Jalan-jalan dan pembelian tidak rutin datang sebagai permintaan | PRN-05, PRN-08, PRN-09, PRN-10, PRN-17 | PRN-06 (menandai, tidak mengunci) |
| `KS-5` Tiap pasangan merasa terbantu memperhatikan pasangannya | PRN-14, PRN-08 | PRN-12 (hanya yang dibagikan), PRN-14 (tidak diukur) |
| `PAGAR` Dua arah; berbagi suasana hati pilihan yang membagikan | PRN-09, PRN-10, PRN-11, PRN-12 | — |
| Tahap publik dan mutu perencanaan | PRN-03, PRN-04 | — |

### 5.2 Ketegangan antarprinsip dan aturan penyelesaiannya

| # | Ketegangan | Aturan penyelesaian |
|---|---|---|
| KT-1 | **PRN-13 lawan PRN-12.** Keterbukaan dana bersama lawan privasi dana pribadi | Batasnya adalah *garis kendali rumah tangga* (07 §2.2; ADR-003). Yang melintasi garis terlihat di sisi rumah tangga beserta jumlah dan pelakunya; saldo dan pemakaian di sisi pribadi tertutup. Uang masuk yang diarahkan ke dana pribadi mengikuti keputusan pemilik nomor 11. Jatah pribadi tidak menggeser garis ini; ia hanya menentukan tanda atas pengambilan yang tetap terlihat (D-48, diterima lewat "lanjut"; melewati jatah: keputusan pemilik nomor 20). Daftar yang terlihat bersifat tertutup dan diberitahukan sebelum mencatat. Bila ragu apakah sesuatu menyangkut dana bersama, ia terlihat |
| KT-2 | **PRN-05 lawan PRN-15 dan PRN-18.** Usaha rendah lawan ketelitian | Kebenaran saldo dan custody didahulukan atas ketelitian kategori. Entri diterima dengan isian paling sedikit; ketelitian dipulihkan kemudian lewat rekonsiliasi kas dan perincian selisih (07 §3.10). Aplikasi tidak pernah menolak entri karena kurang rinci, dan tidak pernah mengisi kekurangannya dengan tebakan |
| KT-3 | **PRN-14 terhadap dirinya sendiri.** Kehangatan lawan kasih sayang yang pura-pura | Orang membaca dan mengirim sendiri; aplikasi tidak pernah berbicara sebagai pasangan; tidak ada hitungan. Cara menjaga saran tidak terasa palsu direkomendasikan di dokumen 05 dan diputuskan pemilik produk di checkpoint S3 (Q-32) |
| KT-4 | **PRN-11 lawan `KS-3`.** Laporan per orang lawan larangan papan skor | Laporan "siapa membayar apa" ada, dengan syarat 06 AT-7. Bila `security-privacy` menilai AT-7 tidak cukup (06 §15.5), laporannya yang berubah, bukan larangannya |
| KT-5 | **PRN-06 lawan `KS-4`.** Tidak pernah menghalangi lawan "bukan kejutan" | Menandai, tidak mengunci (ADR-009). `KS-4` diukur di pilot dari berapa pemakaian tidak rutin yang datang sebagai permintaan. Bila tanda di atas ambang sering muncul, yang dibahas pasangan adalah kebiasaan dan nilai ambangnya; kunci tidak ditambahkan |
| KT-6 | **PRN-10 lawan PRN-01.** Diam tidak memutuskan lawan "tahu tanpa bertanya" | Permintaan boleh menggantung. Yang dijamin PRN-01 adalah bahwa keduanya **melihat** ia menggantung dan sejak kapan, bukan bahwa ia terjawab. Pemohon punya pengingat ulang berbatas, pembatalan, dan pengajuan ulang; aplikasi tidak menjawab untuk siapa pun |
| KT-7 | **PRN-16 lawan hak penghapusan data** (brief §7, §8) | Sementara (ADR-007; 07 §8; 06 §2.5): inti moneter dan inti alur tetap, dan kaitan identitasnya dapat diputus; rincian, buku pribadi anggota yang meminta penghapusan, dan catatan `CARE` dapat dihapus, persis daftar yang dinyatakan PRN-16. Diterima pemilik produk sebagai keputusan sementara (keputusan nomor 24; dokumen 17 S2-8); menunggu tinjauan hukum (L-04, L-20, L-22, L-23). Bila hukum menuntut lain, daftar hapus PRN-16 yang disesuaikan lewat ADR baru; jalan cadangannya opsi D di ADR-007 |
| KT-8 | **PRN-02 lawan PRN-18.** Konfirmasi serah terima menambah satu ketukan | Dibenarkan: tanpa itu seseorang dapat tercatat memegang uang yang tidak ia terima (06 §7.1). Usahanya dibatasi satu ketukan dari notifikasi, tanpa tenggat, dan nol bila penerima sendiri yang mencatat |
| KT-9 | **PRN-09 lawan kenyataan yang tidak simetris.** Keduanya berpenghasilan (pernyataan pemilik produk, 11 November 2026), tetapi satu orang membayar cicilan dan tagihan lalu menyerahkan sisanya, satu orang memegang uang belanja, dan penghasilan keduanya belum tentu sama besar | Aturan tetap tidak membaca siapa orangnya. Tempat aturan simetris berakibat tidak simetris didaftar terbuka (07 §10.3; dokumen 03 bagian 6) dan disampaikan kepada pemilik produk di checkpoint, tidak diperbaiki dengan aturan khusus bagi satu pihak |
| KT-10 | **PRN-03 lawan PRN-02.** Fondasi tahap publik lawan memotong yang tidak dibutuhkan pilot | Fondasi terbatas pada empat hal di PRN-03 dan hanya di model. Apa pun di luar itu yang "untuk nanti" tidak dibangun dan tidak didokumentasikan lebih dari garis besar Tier 3 |
| KT-11 | **PRN-07 lawan PRN-01.** Layar pertama yang tenang lawan angka uang yang langsung terlihat | Angka utama dapat dicapai tanpa navigasi; menyamarkannya di layar adalah pilihan pelihatnya sendiri (brief §10.4). Bentuknya diputuskan dokumen 08 |
| KT-12 | **Arah 2 ("mengingat dan mencatat", bukan menagih) lawan keputusan pemilik nomor 6 dan 13** (pemohon dapat mengingatkan lagi). Pengingat ulang ditekan pasangan yang menunggu, lewat aplikasi; itu pasangan yang mengingatkan, bukan aplikasi (06 §10, ketegangan 1) | Keputusan 6 dan 13 tetap. Yang meredamnya adalah PRN-08: isi kalimat baku aplikasi tanpa teks bebas, jarak 6 jam, batas tiga pengingat, plafon harian, jam tenang, tunda senyap. Pengingat otomatis tetap berpelaku aplikasi. Cara menyajikan "sudah diingatkan berapa kali" pada permintaan diputuskan dokumen 08; bila pemilik produk menghendaki hanya aplikasi yang mengingatkan, keputusan 6 yang ia ubah, bukan prinsip ini |

---

## 6. Yang ditolak sebagai prinsip

Kalimat berikut ditimbang dan ditolak karena tidak dapat gagal dalam tinjauan, atau karena bertentangan dengan prinsip di atas.

| Kalimat | Mengapa ditolak | Gantinya |
|---|---|---|
| "Hangat, tenang, premium, modern" (brief §10.1) | Tidak ada layar yang dapat dinyatakan melanggarnya | PRN-07, dan pilihan terperiksa di dokumen 08 |
| "Sistem canggih, pengalaman tanpa usaha" (brief §3) | Slogan | PRN-02 dan PRN-05, yang punya angka |
| "Satu lingkungan terpadu untuk seluruh urusan keluarga" (brief §3) | Mendorong penambahan modul; "too many features" adalah risiko brief §14 sendiri, dan dokumen 18 §1.3.4 mencatatnya pada organizer serba ada **[Usulan]** peneliti | PRN-02: modul masuk karena menyelesaikan `MSL-n`, bukan karena melengkapi |
| "AI-first" atau "cerdas" | Brief §11 meminta fitur AI tanpa masalah nyata ditolak | PRN-14 |
| "Aman dan privat" | Tidak dapat gagal | PRN-09 sampai PRN-13, masing-masing dengan test |

## 7. Yang dibawa ke lead

1. **Hadiah untuk satu orang dan keputusan pemilik 11 November 2026.** Keputusan itu menyebut uang masuk yang "diarahkan" ke dana pribadi saat dicatat. Hadiah untuk satu orang tidak diarahkan; bawaannya memang dana pribadi (keputusan pemilik nomor 2). Bila catatan hadiah juga terlihat, setiap uang masuk pribadi terlihat pasangan dan "saldo tertutup" tinggal nama. Dokumen 07 kini mengusulkan: yang wajib terlihat hanya penghasilan (gaji, hasil usaha); hadiah untuk satu orang tidak terlihat (07 §3.1, INV-24, T-14). Dokumen 06 kini berbunyi sama (06 §6.4 "Uang masuk yang diarahkan ke dana pribadi", bagian "Cakupan"; V-8; P-13); perbedaan yang semula dicatat di dokumen 17 (K-25) selesai. **Selesai di checkpoint S2:** pemilik produk memutuskan hadiah untuk satu orang tidak terlihat (keputusan nomor 19).
2. **Batas pengingat.** 06 PG-06 kini berbunyi satu per 6 jam sesuai keputusan pemilik (INV-A30). **Selesai di checkpoint S2:** batas lainnya (paling banyak tiga pengingat ulang per orang selama satu rantai permintaan, plafon enam pemberitahuan yang meminta tanggapan per hari, tunda senyap 24 jam) dan bawaan jam tenang 22.00 sampai 06.00 per orang **diterima pemilik produk atas saran lead** (S2-5 pilihan A; dokumen 17, kelompok (b)); bukan pernyataannya sendiri, dan semua angkanya **[Asumsi]** (dokumen 17 A-28). Yang ia putuskan sendiri tetap hanya batas 6 jam dan jam tenang penerima (keputusan 13). Sebelum checkpoint S2 butir ini berbunyi: batas lainnya usulan `architect` yang belum dijawab (dokumen 17, Q-53 dan Q-69).
6. **Sesudah checkpoint S2.** (a) Kode `KT-12` baru. (b) Daftar kata pada pemeriksaan kata-kata PRN-18 adalah **[Usulan]** yang menunggu dokumen 08 dan pemilik produk (dokumen 17 Q-119). (c) Bacaan lead atas kedua pernyataan arah belum dikukuhkan pemilik produk kata per kata; bila ia mengubahnya, yang berubah adalah kalimat baru di PRN-10 dan PRN-18, bukan aturan di dokumen 06 dan 07. (d) Dokumen 06 §10 mengutip dua kalimat lain sebagai pernyataan arah pemilik produk yang tidak ada di `_progress.md` (dokumen 17 K-35); dokumen ini tidak memakainya. Lead memastikan pada 12 November 2026 bahwa kedua kalimat itu bukan kata-kata pemilik produk; dokumen 06 revisi kedelapan sudah mencabutnya (K-35 ditutup). (e) **Sesudah empat jawaban 12 November 2026** (dokumen 17, kelompok (i)): tanda dasar jatah dan dasar pos rutin tidak pernah menjadi pemberitahuan langsung, hanya tampil di catatan dan ringkasan (diterima atas saran lead); ini searah dengan PRN-08 dan PRN-18 dan tidak mengubah kalimat prinsip mana pun. Usulan `security-privacy` yang menyentuh prinsip dan belum disajikan (SP-43 (b): butir bertanda di ringkasan sebagai satu baris tanpa hitungan) ditimbang di dokumen 05 terhadap PRN-11.
3. **Sasaran dua ketukan** (PRN-05) dan **satu ketukan dari layar pertama** (PRN-01) adalah **[Usulan]** yang belum ada di `ENGINEERING-STANDARDS.md` §3; diusulkan untuk dikukuhkan di dokumen 08 dan, bila diterima, ditambahkan lewat ADR.
4. **Tidak ada pengingat mencatat** (PRN-08) adalah usulan yang membatasi `CAL`/`NTF` Tier 2; diputuskan di S3.
5. Kode `PRN-NN`, `MSL-n`, `KS-n`, `PAGAR`, `TUJUAN`, dan `KT-n` (ketegangan di 5.2) untuk README. `KT` dipilih agar tidak bertabrakan dengan `T-NN` milik dokumen 07 dan `K-NN` milik dokumen 17.
