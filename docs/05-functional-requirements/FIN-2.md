# 05 — Kebutuhan Fungsional: FIN bagian 2 (F-FIN-011 sampai F-FIN-020)

Bagian dari dokumen 05. Indeks, tabel keterlacakan, dan pemetaan ke skenario (dokumen 12) serta fase roadmap (dokumen 13) ada di [../05-functional-requirements.md](../05-functional-requirements.md). Penulis: `product-analyst`, S3 fase B, 12 November 2026.

| Hal | Isi |
|---|---|
| Status | **Sesudah checkpoint S4, 12 November 2026 (ditinjau `critic` A dan B, 12 November 2026; temuan diperbaiki; perbaikan diperiksa lead dengan `grep`, tanpa tinjauan ketiga):** acuan 07 T-46 (`joint(S)`, S4-8 (iv) A) dan T-47 (kecocokan kategori pos rutin, S4-8 (v) A) kini keputusan pemilik produk di checkpoint S4, bukan usulan yang menunggu; acuan `others_may_count` mengikuti S4-8 (ii) A. Tidak ada kriteria yang berubah isinya; hitungan tidak berubah. **S4 fase D, 12 November 2026 (belum ditinjau `critic`):** kecocokan kategori pos rutin (07 INV-37, §3.15 aturan 7b, T-47; 06 INV-A52; 10.3 SEC-92): FR-FIN-100/AC-9, FR-FIN-103/AC-14, FR-FIN-141/AC-4 baru. Tidak ada angka uang baru; contoh kategori disalin dari uji INV-37. **Penyisiran andaian khusus pilot S4, 12 November 2026 (M2, M4, M6, M7; belum ditinjau `critic`):** tingkat `spouses` dibaca `adults`; BR-FIN-080 dan FR-FIN-135 membaca `others_may_count` sebagai himpunan ("penghitung termasuk `others_may_count`") dan menambah `joint(S)`; FR rujukan 07 §3.7 kini "Utang antar-pemilik". Tidak ada angka yang berubah. Sebelumnya: ditinjau `critic` A dan diperiksa `security-privacy` (fase D) |
| Putaran perbaikan fase D (12 November 2026) | Diselaraskan dengan 07 putaran perbaikan S3 (§3.7, §3.8 `H`, §3.10 butir 1 (a) sampai (c), §2.6, §8.2 butir 2b, INV-33, INV-34, CTH-25 varian SP-51, CTH-47) dan 06 revisi kesepuluh (N-14, N-32, §8.1 "Mengarsipkan tujuan tabungan"): FR-FIN-100/AC-1 (N-32; CA-08), FR-FIN-117 AC-4, AC-7 sampai AC-9 (SP-51), FR-FIN-135 AC-8, AC-9, AC-11 sampai AC-13 (SP-50), FR-FIN-128 (G-08 tertutup), pengenal pertanyaan `P-n` menjadi `PA-FIN-12` sampai `PA-FIN-18` (CA-06). Aturan baru itu usulan `finance-domain`, `architect`, dan `security-privacy` (SP-50, SP-51); **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026), S3-9 (i), (ii), dan N-32 lewat S3-1 (d). Angka **[Asumsi]** tetap; angka baru belum dihitung ulang lead |
| Sesudah checkpoint S3 (12 November 2026) | Jawaban "lanjut" dicatat: penanda "checkpoint S3-N" dan cabang "bila ditolak" atas butir yang kini diterima diganti, dengan penulis aturannya tetap disebut. Tidak ada angka atau AC lama yang berubah selain hilangnya cabang "bila ditolak" pada FR-FIN-100/AC-1. **Ditambahkan:** FR-FIN-100/AC-8 pos rutin tanpa tanggal (IDEA-03; 07 §3.15, CTH-50, T-41), angkanya dari CTH-50 saja; sisi alur 06 §6.3 baris (h), N-32; skenario SCN-11. FR-FIN-135/AC-14 (IDEA-01, hanya uang tunai berpemilik tunggal) baru, angkanya dari CTH-47 saja; skenario SCN-28 |
| Putaran penyelarasan (12 November 2026) | Diselaraskan dengan 07 putaran penyebaran S3 (§10.2 "Kepada `product-analyst`" butir 2; §11 "Putaran penyebaran S3"; T-30 sampai T-38; CTH-45 sampai CTH-49) dan 06 revisi kesembilan (§6.2, §6.3 "Pos rutin" (f) dan (g), §7.2, §8.1, §12 S-9 sampai S-11, §15.11, §16). G-01 sampai G-07 tertutup; daftarnya di bagian 12. Angka baru atau berubah belum dihitung ulang lead |
| Chat rumah tangga | Persyaratan chat milik `COM.md`. Yang menyentuh `FIN` di berkas ini merujuk aturan 06 §16 dan FIN-1.md KU-8, bukan pengenal `COM` |

## Pengantar

**Lingkup.** Sepuluh fitur Tier 1 modul `FIN` dari dokumen 04 §3.3, semuanya Must: tagihan berulang dan pos rutin, cicilan, utang dan piutang, alokasi, tabungan, aset dan kekayaan bersih, hitung uang dan dompet campuran, koreksi dan keberatan, riwayat satu catatan, serta ringkasan bulanan dan arus kas. Fitur F-FIN-001 sampai F-FIN-010 ada di [FIN-1.md](FIN-1.md).

**Rentang ID.** Kebutuhan fungsional FR-FIN-100 sampai FR-FIN-199; business rule BR-FIN-050 sampai BR-FIN-099. Kriteria penerimaan bernomor lokal per FR dan dirujuk sebagai FR-FIN-112/AC-3.

**Cara membaca.**

| Hal | Arti di berkas ini |
|---|---|
| Aturan uang dan alur | Tidak didefinisikan di sini. Aturan uang kanonis di dokumen 07, alur, visibilitas, dan pagar di dokumen 06. BR di bagian 1 hanya pernyataan tingkat produk yang menunjuk sumbernya |
| **[Perlu aturan]** | Hal yang tidak diatur 07 maupun 06. Bawaan yang ditulis di situ usulan, bukan aturan; permintaannya ada di bagian 12 |
| Kedudukan | Keputusan pemilik produk hanya dikutip dari dokumen 17 atau `_progress.md`, dengan tingkat kepastiannya. Bawaan yang diterima lewat "lanjut" bukan pernyataannya |
| Angka | Rupiah bulat tanpa pemisah ribuan, sama dengan contoh terhitung 07 (CTH). Angka yang bukan dari CTH ditandai **[Asumsi]** ilustrasi |
| Kata dalam tanda kutip | Isi yang harus tampil, bukan kata-kata akhir; kata-kata layar milik dokumen 08. Istilah model ("tanda", "ditolak", "keberatan", `routine_amount`) bukan kata layar (Q-119) |
| Dua arah | Kecuali disebut lain, contoh memakai suami dan istri sebagai ilustrasi; menukar perannya menghasilkan hal yang sama, karena setiap izin di 06 diberikan menurut peran dalam alur, bukan menurut orang (06 PG-01, INV-A13). FR yang perannya tidak dapat dipertukarkan (pemegang dompet campuran, pencatat entri, pengaju keberatan) menyebutnya di baris "Aktor dan izin" |
| Ketentuan umum offline | Berlaku untuk setiap perintah tulis di berkas ini dan tidak diulang per FR: setiap perintah membawa `idempotency_key` sekali per niat (07 INV-10; 06 S-2); pencatatan offline tidak pernah ditolak (06 S-7); state dan tanda hanya diputuskan server, dan perangkat menampilkan "menunggu terkirim" sampai diterima (06 S-1, S-8). FR hanya menulis perilaku offline yang khusus |
| Tata letak | Tidak ada. FR menyebut apa yang harus dapat dilakukan dan dilihat; susunan layar milik dokumen 08 |

**Daftar FR.**

| ID | Judul | Fitur |
|---|---|---|
| FR-FIN-100 | Membuat tagihan berulang beserta usulan kesepakatannya | F-FIN-011 |
| FR-FIN-101 | Menanggapi usulan pos rutin | F-FIN-011 |
| FR-FIN-102 | Menyetujui beberapa pos rutin sekaligus | F-FIN-011 |
| FR-FIN-103 | Membayar tagihan dengan entri yang tertaut | F-FIN-011 |
| FR-FIN-104 | Status kesepakatan hanya di layar pos itu | F-FIN-011 |
| FR-FIN-105 | Mengubah jumlah yang disepakati dan tawaran satu ketukan | F-FIN-011 |
| FR-FIN-106 | Mengakhiri kesepakatan pos rutin | F-FIN-011 |
| FR-FIN-107 | Tanggal akhir, pengarsipan, pengaktifan kembali, dan penghapusan tagihan | F-FIN-011 |
| FR-FIN-108 | Menyiapkan uang untuk tagihan periode ini | F-FIN-011 |
| FR-FIN-110 | Mendaftarkan cicilan yang sudah berjalan | F-FIN-012 |
| FR-FIN-111 | Membeli barang secara cicilan | F-FIN-012 |
| FR-FIN-112 | Membayar cicilan | F-FIN-012 |
| FR-FIN-113 | Melunasi cicilan lebih awal dengan potongan | F-FIN-012 |
| FR-FIN-115 | Memberi pinjaman kepada pihak luar dan menerima pelunasannya | F-FIN-013 |
| FR-FIN-116 | Menerima pinjaman dari pihak luar dan membayarnya kembali | F-FIN-013 |
| FR-FIN-117 | Mengikhlaskan piutang | F-FIN-013 |
| FR-FIN-118 | Kiriman keluarga atau pinjaman | F-FIN-013 |
| FR-FIN-119 | Utang antara dana bersama dan dana pribadi, dan pembebasannya | F-FIN-013 |
| FR-FIN-120 | Membuat alokasi | F-FIN-014 |
| FR-FIN-121 | Memakai, mengurangi, dan memindahkan alokasi | F-FIN-014 |
| FR-FIN-122 | Alokasi yang melebihi uang | F-FIN-014 |
| FR-FIN-125 | Membuat tujuan tabungan | F-FIN-015 |
| FR-FIN-126 | Menabung | F-FIN-015 |
| FR-FIN-127 | Kemajuan tabungan dan baris di halaman depan | F-FIN-015 |
| FR-FIN-128 | Menyelesaikan atau membatalkan tujuan tabungan | F-FIN-015 |
| FR-FIN-130 | Membeli aset | F-FIN-016 |
| FR-FIN-131 | Memperbarui nilai aset | F-FIN-016 |
| FR-FIN-132 | Menjual atau melepas aset | F-FIN-016 |
| FR-FIN-133 | Menampilkan kekayaan bersih | F-FIN-016 |
| FR-FIN-135 | Menghitung isi satu tempat uang | F-FIN-017 |
| FR-FIN-136 | Merinci selisih dan entri terlambat yang menutupnya | F-FIN-017 |
| FR-FIN-137 | Menghitung dompet campuran | F-FIN-017 |
| FR-FIN-138 | Kapan menghitung ditawarkan | F-FIN-017 |
| FR-FIN-140 | Membalik entri | F-FIN-018 |
| FR-FIN-141 | Menyesuaikan entri dan memindah kategorinya | F-FIN-018 |
| FR-FIN-142 | Batas koreksi atas buku pribadi dan entri dua buku | F-FIN-018 |
| FR-FIN-143 | Mengajukan keberatan | F-FIN-018 |
| FR-FIN-144 | Menutup keberatan dan daftar "belum sepakat" | F-FIN-018 |
| FR-FIN-145 | Riwayat satu catatan | F-FIN-019 |
| FR-FIN-146 | Tidak ada riwayat lintas catatan menurut orang | F-FIN-019 |
| FR-FIN-150 | Periode ringkasan dan penjelasan perubahan kekayaan bersih | F-FIN-020 |
| FR-FIN-151 | Pengeluaran per kategori | F-FIN-020 |
| FR-FIN-152 | Siapa membayar apa | F-FIN-020 |
| FR-FIN-153 | Tabungan, aset, piutang, dan kewajiban di ringkasan | F-FIN-020 |
| FR-FIN-154 | Arus kas dana bersama | F-FIN-020 |
| FR-FIN-155 | Angka per anggota hanya per periode, tanpa hitungan | F-FIN-020 |
| FR-FIN-156 | Ringkasan periode lampau yang berubah | F-FIN-020 |

## 1. Business rule

Setiap BR menunjuk sumber kanonisnya dan tidak mendefinisikan aturan untuk kedua kalinya. Bila bunyinya terasa berbeda dari sumbernya, sumbernya yang berlaku.

| ID | Pernyataan tingkat produk | Sumber kanonis |
|---|---|---|
| BR-FIN-050 | Tagihan berulang adalah jadwal dan pengingat, bukan entri. Tagihan yang jatuh tempo tetapi belum dibayar bukan kewajiban dan tidak mengubah angka apa pun. Kejadian per bulan, pembayaran yang menunjuk satu kejadian, dan arti "belum dibayar periode ini": BR-FIN-047 (`FIN-1.md`). Tagihan tidak pernah terarsip sendiri; sesudah tanggal akhir tidak ada kejadian baru dan pengingat berhenti | 07 §3.15, INV-12, INV-26, T-36; D-12 |
| BR-FIN-051 | Setiap pos rutin disepakati saat dibuat, berapa pun jumlahnya; jumlah per kali yang disepakati adalah batasnya; untuk tagihan berubah-ubah yang disepakati jumlah tertinggi yang wajar. Kategori pos wajib dan tidak pernah berganti; entri hanya dapat ditautkan bila setiap baris `expense`-nya berkategori pos itu, dan entri tertaut tidak dapat dipindah ke kategori lain (S4 fase D) | 07 §3.15 butir 1–2, aturan 7b, INV-37, T-47; 06 §6.3 "Pos rutin" butir 1–2, INV-A52 |
| BR-FIN-052 | Pembayaran tertaut dengan jumlah pemakaian sampai jumlah yang disepakati tidak ditandai, siapa pun pembayarnya, juga bila di atas ambang atau dibayar lebih awal; yang melebihinya ditandai sebagai di atas kesepakatan tanpa dibandingkan dengan ambang. Satu ID untuk aturan ini; BR-FIN-047 di `FIN-1.md` merujuknya | 07 §3.15 butir 3–4, §3.8 "Tanda atas entri"; 06 §6.3 "Pos rutin" butir 3–4, INV-A43 |
| BR-FIN-053 | Selama kesepakatan suatu pos belum berlaku atau sesudah diakhiri, pembayarannya mengikuti aturan ambang biasa; tidak ada yang dihalangi | 07 §3.15 butir 5; 06 §6.3 "Pos rutin" baris (a), (d) |
| BR-FIN-054 | Di rumah tangga satu anggota dewasa tidak ada kesepakatan yang berlaku dan tidak ada tanda; usulan menunggu anggota dewasa kedua | 07 §3.15 butir 6; 06 §6.3 "Pos rutin" baris (b) |
| BR-FIN-055 | Penilaian pembayaran pos rutin terjadi sekali per pembayaran; mengganti atau mengakhiri kesepakatan tidak menilai ulang entri yang sudah tersimpan | 07 §3.15 butir 7; 06 §6.3 "Pos rutin" baris (c) |
| BR-FIN-056 | Yang dapat ditautkan ke tagihan berulang hanya `expense`, `contribution_expense`, dan `advance`; `contribution_expense` tertaut tidak pernah ditandai; jenis lain ditolak tautannya dan entrinya tetap tersimpan | 07 §3.15 butir 7a; 06 §6.3 "Pos rutin" baris "Entri apa yang dapat ditautkan" |
| BR-FIN-057 | Tautan susulan ke pos rutin hanya dapat dibuat pencatat entri itu | 06 §4.5, §6.3 "Pos rutin"; 07 §3.8 "Penautan" butir 3; Q-116 |
| BR-FIN-058 | Tanda tidak pernah menghalangi, menahan, atau mengubah angka apa pun | 06 §6.2; 07 §3.8; PRN-06 |
| BR-FIN-059 | Pemberitahuan atas tanda dasar kesepakatan pos rutin hanya masuk ringkasan, tidak pernah pemberitahuan langsung; butir bertanda di ringkasan dilebur menjadi satu baris tanpa hitungan | 06 §6.3 "Pos rutin" baris (e), §6.2, N-10, AT-4; SP-37, SP-43 |
| BR-FIN-060 | Mengubah jumlah yang disepakati adalah usulan pengganti; jumlah lama berlaku sampai usulan baru disetujui | 06 §6.3 K01, K05, "Pos rutin" baris (c) |
| BR-FIN-061 | Kesepakatan pos rutin dapat diakhiri salah satu anggota dewasa sendirian, berlaku seketika dan ke depan saja | 06 §6.3 K06, "Pos rutin" baris (d) (**[Usulan]** `architect`) |
| BR-FIN-062 | Cicilan bukan pos rutin: pembayarannya tidak memerlukan kesepakatan dan tidak pernah ditandai; denda di entri yang sama dinilai seperti pengeluaran lain | 07 §3.14, §3.15 butir 8, §3.8; D-44 |
| BR-FIN-063 | Kewajiban cicilan adalah **sisa yang harus dibayar**: jumlah seluruh pembayaran yang tersisa, tanpa memisahkan pokok dan bunga | 07 §3.14 |
| BR-FIN-064 | Biaya kredit = uang muka + jumlah seluruh cicilan − harga, dicatat saat beli sebagai baris ringkasan, tidak dihitung terhadap anggaran. Uang muka + jumlah seluruh cicilan yang lebih kecil daripada harga ditolak; pada cicilan 0% tidak ada baris biaya kredit | 07 §3.14, §6.1, T-35 (G-07) |
| BR-FIN-065 | Pembelian secara cicilan dibandingkan dengan ambang satu kali sebesar uang muka ditambah seluruh cicilan; kewajiban yang didaftarkan sebagai saldo awal sesudah ada ambang dibandingkan satu kali sebesar saldo awalnya | 07 §3.14, §3.8 |
| BR-FIN-066 | Pelunasan tidak melebihi sisa piutang atau kewajiban; kelebihan bayar atas piutang dicatat sebagai pemasukan "lainnya" | 07 §3.7, INV-13 |
| BR-FIN-067 | Piutang yang diikhlaskan tampil sebagai baris tersendiri "piutang diikhlaskan" pada bulan kejadiannya, tidak masuk anggaran atau arus kas, dan tidak digabung dengan kiriman keluarga. Mengikhlaskan piutang dana bersama dinilai terhadap ambang secara kumulatif per piutang (ditandai bila `N` > 0 dan `H + N` > ambang), dengan alasan wajib, oleh anggota dewasa mana pun | 07 §3.7, §3.8, INV-34, T-32 (G-03; SP-51); 06 §6.2; D-17 |
| BR-FIN-068 | Kiriman keluarga adalah pengeluaran yang tidak dilacak sesudahnya; untuk saudara atau orang tua pencatat memilih per entri antara pinjaman dan kiriman | 07 §3.2 "Kiriman keluarga" |
| BR-FIN-069 | Utang antar-pemilik hanya lahir bila uang berpindah dengan janji kembali, dan selalu terlihat oleh kedua pihaknya; antara dua dana pribadi ditunda dari MVP | 07 §3.7, INV-08, T-12; Q-24 |
| BR-FIN-070 | Hanya pemilik piutang yang dapat membebaskannya; utang anggota kepada dana bersama dibebaskan oleh anggota dewasa lainnya, mengikuti ambang, dan tampil sebagai "utang dihapuskan" | 07 §3.6, §3.8; D-43 |
| BR-FIN-071 | Saat dibuat, jumlah alokasi pada satu tempat uang tidak boleh melebihi saldo tercatatnya | 07 §3.9, INV-17; Q-45; D-17 |
| BR-FIN-072 | Belanja yang membuat alokasi melebihi uang tidak ditolak dan tidak mengurangi alokasi secara otomatis; tempat uang menampilkan "alokasi melebihi uang Rp…" | 07 §3.9 |
| BR-FIN-073 | Anggaran tidak memerlukan alokasi | 07 §3.9; keputusan 8 |
| BR-FIN-074 | Menabung adalah alokasi atau pemindahan yang membawa tujuan, tidak pernah pengeluaran; kemajuan = jumlah alokasi tujuan itu di semua tempat uang; target tidak masuk angka mana pun | 07 §3.9, §6.5; PRN-17 |
| BR-FIN-075 | Menyelesaikan atau membatalkan tujuan tabungan adalah satu langkah yang menawarkan melepas alokasinya tanpa pilihan terisi (satu entri `allocation` berbaris negatif); tujuan terarsip hanya bila alokasinya 0; alasan arsip (tercapai atau batal) hanya keterangan | 07 §3.9 "Menyelesaikan atau membatalkan tujuan", §2.6, T-38 (G-01) |
| BR-FIN-076 | Membeli aset bukan pengeluaran dan tidak mengubah kekayaan bersih; barang mana yang disebut aset adalah pilihan pencatat | 07 §3.16, T-05; D-13 |
| BR-FIN-077 | Nilai aset diperbarui manual dan bertanggal; tidak ada penyusutan otomatis | 07 §3.16; 04 §8.1 (F-FIN-026) |
| BR-FIN-078 | Menjual aset: aset berkurang sebesar nilai tercatat, uang bertambah sebesar hasil, selisihnya perubahan nilai | 07 §3.16 |
| BR-FIN-079 | Kekayaan bersih rumah tangga = uang + piutang + aset − kewajiban dana bersama; target, anggaran, otorisasi, dan reservasi tidak masuk | 07 §6.5 |
| BR-FIN-080 | Beda antara terhitung dan tercatat menjadi satu entri selisih kas yang terlihat; catatan tidak pernah disunting; siapa yang lupa diselesaikan di luar aplikasi. Selisihnya dihitung server; perangkat hanya menampilkan perkiraan berlabel. Bukan pemegang boleh menghitung tempat uang dana bersama hanya bila tempat uang itu `bank` atau `ewallet` dan penghitung termasuk `others_may_count` tempat uang itu, himpunan anggota pilihan pemegang (a; penyisiran S4, M4), tidak ada hitungan bukan pemegang lain yang terbuka atasnya (b), dan pemegangnya aktif, tidak dalam masa tenggang keluar (c); selisihnya bertanda di sisi pemegang, dan N-14-nya `ringkasan`. Atas tempat uang `joint(S)` hanya anggota `S` yang menghitung (07 INV-33, M6). Uang tunai tidak pernah dihitung anggota lain. Hitungan yang cocok tidak membuat entri | 07 §3.10 butir 1 (a) sampai (c), "Hitungan yang tiba belakangan", INV-33, T-33 (G-02, G-05; SP-50); 06 §7.2 C01, §12 S-10, N-14; keputusan 7 |
| BR-FIN-081 | Pengeluaran terlambat bertanggal tidak lebih baru daripada hitungan terakhir ditawari menjadi bagian selisih, dengan bawaan ya, sampai sebesar yang belum dirinci | 07 §3.10 butir 6, T-04; D-05 |
| BR-FIN-082 | Selisih yang belum dirinci tidak pernah negatif dan baru dihitung terhadap anggaran sesudah dirinci | 07 §3.10 butir 7–8, INV-16 |
| BR-FIN-083 | Dompet campuran dihitung sekali; pemegangnya memilih pembagian selisih setiap kali tanpa pilihan terisi; isi seluruhnya, pilihan pembagiannya, dan bagian pribadinya tidak terbaca pasangan | 07 §2.3, §3.10, T-03; 06 V-9, INV-A41 |
| BR-FIN-084 | Saldo tercatat negatif tidak ditolak; ia memicu ajakan menghitung | 07 §2.5 "Saldo negatif", T-06; D-05 |
| BR-FIN-085 | Koreksi adalah pembalikan atau penyesuaian dengan alasan wajib; entri asal tetap terlihat; satu entri hanya dapat dibalik sekali dan pembalikan tidak dapat dibalik | 07 §3.12, INV-09, INV-18; 06 §8.1 |
| BR-FIN-086 | Entri yang hanya di buku rumah tangga dapat dikoreksi anggota dewasa mana pun; entri buku pribadi hanya oleh pemiliknya; entri dua buku dibalik hanya oleh pemilik buku pribadinya | 07 §3.12; 06 §8.1; Q-26 |
| BR-FIN-087 | Keberatan tidak mengubah angka; alasannya dari daftar tertutup yang netral; hanya pengajunya yang dapat menutupnya; tidak ada penutupan otomatis | 06 §8.2, INV-A25 sampai INV-A28; D-22, D-38 |
| BR-FIN-088 | Daftar "belum sepakat" satu dan sama bagi keduanya, tanpa hitungan per orang | 06 §8.2, AT-4 |
| BR-FIN-089 | Riwayat selalu per catatan; tidak ada umpan, penyaring, atau riwayat lintas catatan menurut orang; tidak ada tanda dibaca dan lokasi | 06 §9.4 AT-1 sampai AT-6, PG-08 |
| BR-FIN-090 | Yang tampil tanggal kejadian, dan tanggal pencatatan hanya bila berbeda; jam pencatatan tidak pernah tampil | 07 §3.11; 06 AT-6 |
| BR-FIN-091 | Periode ringkasan sama dengan periode anggaran; jumlah tabel A sama dengan perubahan kekayaan bersih; jumlah arus kas sama dengan perubahan uang yang ada | 07 §6.3, §6.4, INV-19 |
| BR-FIN-092 | "Siapa membayar apa" dikelompokkan menurut pembayar dengan tiga kolom sumber, per periode, tanpa persentase dan tanpa urutan menurut besarnya, sama bagi keduanya, dan menerangkan bahwa ia menunjukkan lewat tangan siapa uang dibayarkan | 07 §6.3 tabel C, §10.3 butir 2; 06 AT-7; D-01 |
| BR-FIN-093 | Penyisihan dan kontribusi per anggota hanya tampil per periode: tidak ada angka kumulatif, lintas periode, rata-rata, tren, atau bersih per anggota | 07 §6.3; 06 AT-4; D-43 |
| BR-FIN-094 | Tidak ada penutupan bulan; periode lampau yang berubah karena entri terlambat ditandai berubah | 07 §3.11, §6.3; D-05 |
| BR-FIN-095 | Angka rumah tangga identik bagi kedua anggota dewasa | 07 INV-20; 06 AT-2 |
| BR-FIN-096 | Tidak ada layar yang menyatakan atau menyiratkan bahwa semua penghasilan sudah tercatat | 07 INV-24, §10.3 butir 1 dan 8; dokumen 17 "Keluhan" 9; W-1 |
| BR-FIN-097 | Tidak ada hitungan, peringkat, skor, atau tren per orang atas tanda, keberatan, koreksi, selisih kas, catatan terlambat, atau status pos rutin; tidak ada hitungan tanda tingkat rumah tangga | 06 AT-4, PG-09; SP-35, SP-43 |
| BR-FIN-098 | Setiap perintah tulis idempoten: kiriman ulang dengan kunci yang sama menghasilkan satu akibat | 07 INV-10; 06 S-2 |
| BR-FIN-099 | Tanggal kejadian tidak boleh di masa depan dan tidak boleh sebelum saldo awal tempat uangnya; entri terlambat sah sejauh apa pun ke belakang | 07 §3.11, INV-22; PRN-05 |

## 2. F-FIN-011 — Tagihan berulang dan pos rutin

Masalah: `MSL-6` (siapa membayar tagihan apa), `MSL-3`. Beban berulang: sekali per tagihan saat mulai (4 sampai 8 tagihan **[Asumsi]**, dokumen 04), satu persetujuan pasangan per pos, dan satu lagi setiap kali suatu jumlah berubah (06 §6.3 "Batas dan beban"). Pembayaran bulanan memakai usaha F-FIN-005; pengingatnya N-32 Tier 1 (`ringkasan`, tiga hari sebelum dan pada hari jatuh tempo, per kejadian; 06 §11.2, §13.4; usulan `architect`, diterima lewat "lanjut" di checkpoint S3, S3-1 (d)) dan berjalan apa pun state kesepakatannya (04 KM-04). FR-CAL-002 di `tier-2.md` kini hanya penunjuk ke aturan ini.

### FR-FIN-100 — Membuat tagihan berulang beserta usulan kesepakatannya
| | |
|---|---|
| Fitur | F-FIN-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6, MSL-3 |
| Aktor dan izin | Anggota dewasa mana pun membuat; tagihan dan usulannya `adults`, terlihat keduanya (06 §9.2, V-1). Simetris (06 PG-01) |
| Aturan acuan | 07 §3.15 butir 1–2, §2.6 `recurring_bill`; 06 §6.3 K01, "Pos rutin" baris "Membuat pos rutin" dan (b), INV-A19, INV-A11, §11.2 N-32 (pengingat tagihan Tier 1, `ringkasan`, usulan `architect`, diterima lewat "lanjut" di checkpoint S3, S3-1 (d)), §13.4; 07 INV-12 |
| Business rule | BR-FIN-050, BR-FIN-051, BR-FIN-054, BR-FIN-098 |
| Kedudukan | Setiap tagihan disepakati saat dibuat dan jumlahnya batas: rumusan lead yang dikukuhkan pemilik produk ("bener", 11 November 2026; 06 §6.3). Satu perintah untuk jadwal dan usulan: **[Usulan]** `architect`. Satu medan jumlah yang juga menjadi perkiraan: **[Usulan]** dokumen ini. Kategori bawaan: **[Asumsi]** (07 T-10) |

Masukan: nama (wajib); pihak penerima (opsional, `counterparty`); aturan jatuh tempo (`due_rule`, **opsional**, bentuknya mengikuti 07 §2.6; dikosongkan untuk pos rutin tanpa tanggal seperti token listrik, gas, atau galon: 07 §3.15 "Tagihan berulang tanpa `due_rule`", 06 §6.3 "Pos rutin" baris (h), IDEA-03, AC-8); kategori (bawaan "tagihan rutin"); penanggung jawab (opsional; tidak ada pilihan terisi); tanggal akhir (opsional, tidak sebelum hari ini); dan **jumlah per kali yang diusulkan** (wajib, bilangan bulat rupiah > 0). Layar meminta jumlah itu sebagai **jumlah tertinggi yang wajar**, bukan rata-rata, dan menerangkan dalam satu kalimat bahwa pembayaran di atasnya akan tampil sebagai di atas kesepakatan (dokumen 17 "Dibawa ke S3"; 06 §6.3 "Batas dan beban", tanda atas selisih kecil). Perkiraan (`estimated_amount`) bawaannya sama dengan jumlah itu dan boleh diisi lain; ia hanya dipakai pengingat dan tawaran alokasi (07 §3.15). Alasan satu medan: pencatat tidak perlu memikirkan dua angka untuk satu tagihan.

Yang dicatat: satu `recurring_bill` dan satu usulan kesepakatan rutin (K01) atasnya, dalam satu transaksi dengan satu kunci. Bila `due_rule` diisi, jadwalnya langsung dipakai pengingat dan pembayaran; bila kosong, pos itu tidak punya kejadian, jadi tanpa pengingat N-32, tanpa "belum dibayar periode ini", dan tanpa "siapkan uang untuk semua", dan pembayarannya dinilai per pembayaran terhadap `R` (AC-8). Tidak ada tagihan berulang tanpa usulan. Pasangan menerima N-11, dengan batas satu pemberitahuan langsung per pengirim per hari untuk usulan pos rutin (06 N-11; SP-36 **[Usulan]**). Offline: perintah menunggu terkirim; jadwal berlaku sejak diterima server.

**Kriteria penerimaan**
- **AC-1** — *Given* rumah tangga CTH-41 (dua anggota dewasa, ambang 500000, hari mulai periode 11) *When* suami membuat tagihan "listrik", jatuh tempo tanggal 20, jumlah per kali 900000 *Then* tercipta satu tagihan berulang dan satu usulan kesepakatan 900000 per kali; istri menerima N-11; tidak ada saldo, anggaran, alokasi, atau kekayaan bersih yang berubah; jadwalnya langsung dipakai pengingat N-32: karena penanggung jawab tidak diisi, kedua anggota dewasa mendapat satu baris `ringkasan` tanpa jumlah pada 17 November (tiga hari sebelum jatuh tempo **[Asumsi]**) dan satu pada 20 November untuk kejadian 20 November, dan baris itu dibuang dari ringkasan bila kejadian itu sudah dibayar saat ringkasan dilepas (06 N-32, §13.4; FR-NTF-009/AC-5). Bagian N-32 ini usulan `architect`, diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-1 (d); cabang lama "bila ditolak, AC ini dipersempit" tidak berlaku.
- **AC-2** — *Given* rumah tangga yang sama *When* istri membuat tagihan "air", jatuh tempo tanggal 29, jumlah per kali 250000 *Then* hasilnya sama dengan AC-1 dengan peran tertukar: suami menerima N-11.
- **AC-3** — *Given* formulir tagihan baru *When* jumlah per kali kosong, 0, negatif, atau berpecahan *Then* perintah ditolak di batas dengan keterangan pada medan itu, dan tidak ada tagihan maupun usulan yang tercipta.
- **AC-4** — *Given* suami mengirim AC-1 dan ketukannya terkirim dua kali dengan kunci yang sama *When* server menerima keduanya *Then* hanya ada satu tagihan "listrik", satu usulan, dan satu N-11.
- **AC-5** — *Given* suami hari ini sudah memicu satu pemberitahuan langsung N-11 atas usulan pos rutin *When* ia membuat "iuran kesehatan" 150000 pada hari yang sama *Then* usulannya tercatat dan jadwalnya berlaku seketika, tetapi N-11-nya masuk ringkasan istri, bukan pemberitahuan langsung.
- **AC-6** — *Given* perkiraan dibiarkan kosong *When* tagihan "internet" 350000 disimpan *Then* `estimated_amount` = 350000. *Given* pencatat mengisi perkiraan 600000 untuk "listrik" dengan jumlah per kali 900000 *Then* pengingat memakai 600000 dan penilaian pembayaran memakai 900000 bila kesepakatannya berlaku.
- **AC-7** — *Given* rumah tangga dengan satu anggota dewasa *When* ia membuat tagihan "internet" 350000 *Then* jadwalnya berlaku untuk pengingat dan pembayaran tanpa tanda apa pun; usulannya menunggu dan disampaikan kepada anggota dewasa kedua saat ia bergabung; tidak ada pembayaran lama yang ditandai surut.
- **AC-8** — **Pos rutin tanpa tanggal** (IDEA-03, diterima lewat "lanjut" di checkpoint S3, S3-2 (ii); aturan 07 §3.15 "Tagihan berulang tanpa `due_rule`: pos rutin tanpa tanggal", CTH-50, T-41, diadopsi `finance-domain` sesudah checkpoint). Pos tanpa `due_rule` tidak punya kejadian dan tidak memicu N-32; tiap pembayaran dinilai terhadap `R`; ada-tidaknya `due_rule` ditetapkan saat pos dibuat dan tidak dapat diganti (T-41 **[Usulan]**). *Given* keadaan CTH-50 (pos "Token listrik" tanpa `due_rule`, kesepakatan berlaku `R` 300000, ambang 500000, `TUN-S` 1000000, anggaran "listrik" 600000) *When* suami mencatat token 200000 lalu 300000, tertaut ke pos itu *Then* keduanya tidak ditandai dan `TUN-S` 500000; *When* tautan susulan yang membawa kejadian 20 November dikirim untuk entri 300000 *Then* ditolak dan tidak ada angka yang berubah; *When* token 350000 dicatat *Then* ditandai dasar `routine_amount` (350000 > 300000, tidak dibandingkan dengan ambang), `TUN-S` 150000, sisa anggaran "listrik" −250000; pos itu tidak pernah tampil di "belum dibayar periode ini" dan tidak memicu pengingat. Sisi alur: 06 §6.3 "Pos rutin" baris (h) (revisi kesebelas) dan pengecualian N-32 di 06 §11.2 N-32 ("tidak pernah lahir atas pos rutin tanpa `due_rule`") dan §13.4. Skenario: dokumen 12 SCN-11.
- **AC-9** — **Kategori pos wajib dan tetap** (07 §2.6 `recurring_bill`, §3.15 aturan 7b, INV-37, T-47; 10.3 SEC-92; usulan `finance-domain` S4, **diputuskan pemilik produk** di checkpoint S4, S4-8 (v) A, Q-202 ditutup). *Given* formulir pos rutin baru, *When* kategori dikosongkan, *Then* perintah ditolak di batas dengan keterangan pada medan kategori (bawaannya tetap "tagihan rutin", jadi pencatat yang tidak menyentuhnya tidak terhalang). *Given* pos "Listrik" berkategori "listrik", *When* siapa pun mengirim perintah mengganti kategori pos itu, *Then* ditolak; mengganti kategori berarti membuat pos baru. Alasannya: pembayaran tertaut dinilai terhadap `R` dan tidak terhadap ambang, jadi kategori pos yang dapat diganti akan menjadi jalan meloloskan belanja lain dari tanda (`PAGAR`).

### FR-FIN-101 — Menanggapi usulan pos rutin
| | |
|---|---|
| Fitur | F-FIN-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6, MSL-3 |
| Aktor dan izin | Menyetujui (K02) atau tidak menyetujui (K03, alasan opsional): anggota dewasa selain pengusul. Menarik (K04): pengusul. Keduanya melihat usulan dan jawabannya (06 §9.2). Simetris (PG-01): siapa pun dapat mengusulkan, dan pihak lainnya yang menjawab |
| Aturan acuan | 06 §6.3 K02, K03, K04, "Pos rutin" baris (a); 06 §10 (jendela 6 jam, tiga pengingat ulang per pos, SP-27); 06 S-1, S-3, S-4; N-12 |
| Business rule | BR-FIN-051, BR-FIN-053, BR-FIN-058, BR-FIN-098 |
| Kedudukan | Bahwa tagihan rutin disepakati dulu dan siapa pun boleh membayarnya: keputusan pemilik produk, checkpoint S2 (S2-4 pilihan A; keputusan 22). Akibat usulan yang belum dijawab atau tidak disetujui: **[Usulan]** `architect`, belum disajikan (Q-113) |

Penjawab melihat nama pos, jatuh tempo, dan jumlah per kali yang diusulkan. Jawaban "tidak setuju" tampil dengan kata layar "belum sepakat" (SP-35; dokumen 08). Diam tidak pernah berarti setuju atau menolak, dan usulan tidak kedaluwarsa (06 PG-04). Pengingat atas usulan mengikuti batas 06 §10. Persetujuan offline tampil "menunggu terkirim", tidak pernah "disepakati", di perangkat mana pun.

**Kriteria penerimaan**
- **AC-1** — *Given* usulan suami "listrik" 900000 belum dijawab *When* istri menyetujuinya *Then* kesepakatan berlaku sejak diterima server; pembayaran yang disimpan sesudahnya dinilai terhadap 900000; suami menerima N-12.
- **AC-2** — *Given* usulan istri "air" 250000 *When* suami menyetujuinya *Then* hasilnya sama dengan AC-1 dengan peran tertukar.
- **AC-3** — *Given* usulan suami "listrik" *When* suami sendiri mencoba menyetujuinya *Then* perintah ditolak; usulan tetap belum dijawab.
- **AC-4** — *Given* istri tidak menyetujui "listrik" 900000 tanpa alasan, ambang 500000 *When* suami membayar listrik 780000 tertaut *Then* tagihan tetap dapat dibayar dan pengingatnya tetap berjalan; pembayarannya mengikuti aturan ambang dan ditandai dasar ambang (780000 > 500000; CTH-41 varian); tidak ada yang dihalangi.
- **AC-5** — *Given* usulan "air" 250000 tidak dijawab selama 30 hari *When* suami membayar air 310000 *Then* usulan masih belum dijawab (tidak berlaku, tidak dianggap tidak disetujui); pembayaran mengikuti aturan ambang dan tidak ditandai (310000 ≤ 500000).
- **AC-6** — *Given* istri menyetujui "listrik" saat offline dan suami sementara itu menarik usulannya *When* persetujuan istri tiba *Then* server menolaknya dengan state terbaru, dan perangkat istri menampilkan bahwa usulan sudah ditarik.
- **AC-7** — *Given* persetujuan istri terkirim dua kali dengan kunci yang sama *Then* hanya ada satu K02 dan satu N-12.
- **AC-8** — *Given* usulan suami "listrik" 900000 tampil sebagai kartu di chat (FIN-1.md KU-8) *When* istri mengetuk setuju pada kartu *Then* hasilnya sama dengan AC-1 (K02 yang sama, guard yang sama); *When* ia hanya membalas kartu dengan "oke" *Then* usulan tetap belum dijawab; suami tidak melihat tombol setuju pada kartunya sendiri; kartu yang diketuk offline tampil "menunggu terkirim" dan, bila usulan sudah ditarik, ditolak dengan state terbaru seperti AC-6 (06 §16.4, INV-A44, S-11 b).

### FR-FIN-102 — Menyetujui beberapa pos rutin sekaligus
| | |
|---|---|
| Fitur | F-FIN-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6; MSL-5 (beban satu ketukan per pos saat mulai) |
| Aktor dan izin | Anggota dewasa yang menjadi penjawab usulan-usulan itu. Simetris (PG-01) |
| Aturan acuan | 06 §6.3 "Pos rutin" baris "Menyetujui beberapa pos sekaligus"; 06 S-3 |
| Business rule | BR-FIN-051, BR-FIN-098 |
| Kedudukan | **[Usulan]** `security-privacy` (SP-36), diterapkan 06; belum disajikan kepada pemilik produk |

Satu langkah layar atas beberapa K02; di server tiap pos tetap satu K02 dengan kuncinya sendiri dan berhasil atau gagal sendiri. Tiga syarat: **setiap jumlah per kali tampil**; **tidak ada pilihan yang sudah terisi**; **pos yang jumlah per kalinya di atas ambang yang berlaku tidak dapat dipilih** dan disetujui satu per satu lewat FR-FIN-101.

**Kriteria penerimaan**
- **AC-1** — *Given* ambang 500000 dan empat usulan suami yang belum dijawab: internet 350000, iuran kesehatan 150000, listrik 900000, air 250000 (CTH-41) *When* istri membuka langkah setujui sekaligus *Then* ia melihat internet, iuran kesehatan, dan air masing-masing dengan jumlahnya, tidak satu pun terpilih; listrik tampil tidak dapat dipilih dengan keterangan bahwa ia disetujui sendiri.
- **AC-2** — *Given* AC-1 *When* istri memilih internet dan air lalu mengukuhkan *Then* server mencatat dua K02; iuran kesehatan dan listrik tetap belum dijawab.
- **AC-3** — *Given* AC-1 *When* tidak ada yang dipilih *Then* tidak ada yang dapat dikirim.
- **AC-4** — *Given* istri memilih internet dan air, dan sementara itu suami menarik usulan air *When* istri mengukuhkan *Then* K02 internet berhasil; K02 air ditolak dengan state terbarunya, dan istri melihat bahwa usulan air sudah ditarik.
- **AC-5** — *Given* usulan-usulan itu dibuat istri *When* suami menyetujuinya sekaligus *Then* berlaku sama dengan AC-1 sampai AC-4.

### FR-FIN-103 — Membayar tagihan dengan entri yang tertaut
| | |
|---|---|
| Fitur | F-FIN-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6, MSL-3 |
| Aktor dan izin | Anggota dewasa mana pun membayar, siapa pun penanggung jawabnya. Tautan dibuat pencatat saat mencatat; tautan susulan hanya oleh pencatat entri itu (peran tidak dapat dipertukarkan). Entri dan tandanya terlihat sama oleh keduanya (06 §9.2, AT-2) |
| Aturan acuan | 07 §3.15 butir 3–7a dan "Kejadian tagihan dan 'belum dibayar periode ini'", §2.5 `bill_occurrence_due`, §3.8 "Tanda atas entri" dan "Penautan", INV-26; 06 §6.2, §6.3 "Pos rutin", §4.5, §12 (uji K06 bersamaan dengan pembayaran), §13.4, N-10 |
| Business rule | BR-FIN-047 (`FIN-1.md`), BR-FIN-051, BR-FIN-052, BR-FIN-053, BR-FIN-055, BR-FIN-056, BR-FIN-057, BR-FIN-058, BR-FIN-059, BR-FIN-098, BR-FIN-099 |
| Kedudukan | Butir 3–4: rumusan lead yang dikukuhkan pemilik produk (11 November 2026). Jenis yang dapat ditautkan dan `contribution_expense` tanpa tanda: **[Usulan]** `architect` dan `finance-domain` (Q-128, Q-129). Pemberitahuan hanya ringkasan: diterima pemilik produk atas saran lead (12 November 2026). Satu baris tanpa hitungan di ringkasan: **[Usulan]** `security-privacy` (SP-43 b). Jumlah terisi dari pembayaran terakhir: **[Usulan]** dokumen ini. Kejadian tagihan dan tautan ke tagihan terarsip ditolak: **[Usulan]** `finance-domain`, putaran S3 (07 §3.15, INV-26; PA-FIN-5) |

Pembayaran dimulai dari pengingat atau dari layar pos. Medan jumlah terisi jumlah pembayaran terakhir tagihan itu dan dapat diubah; kategori dari tagihannya; tempat uang dipilih; tanggal kejadian bawaan hari ini. Pembayaran menunjuk satu **kejadian** tagihan; bawaannya kejadian paling awal yang belum dibayar dan jatuh tempo paling lambat akhir periode berjalan, dan pencatat dapat menggantinya; dari pengingat, kejadian pengingat itu (07 §3.15). Pembayaran yang tersimpan mematikan pengingat kejadian itu saja (06 §13.4). Tautan ke tagihan terarsip ditolak dan entrinya tetap tersimpan tanpa tautan; untuk menautkan, tagihan diaktifkan kembali lebih dulu (FR-FIN-107; 07 INV-26). Server menilai tanda saat entri diterima: tertaut ke pos dengan kesepakatan berlaku → dibandingkan dengan jumlah per kali `R`; tanpa kesepakatan berlaku → ambang. Tanda menampilkan dasarnya dan `R` saat dinilai (06 §6.3 "Yang ditampilkan pada tanda"). Tanda dasar kesepakatan hanya masuk ringkasan pasangan; di ringkasan, butir bertanda dilebur menjadi satu baris tanpa hitungan (rinciannya milik NTF.md).

**Kriteria penerimaan**
- **AC-1** — *Given* "listrik" berkesepakatan 900000, ambang 500000, `BNK-S` 1900000 (CTH-41, 20 November) *When* suami mencatat pembayaran listrik 780000 dari `BNK-S` tertaut *Then* `BNK-S` menjadi 1120000 dan entrinya tidak ditandai, walaupun 780000 di atas ambang.
- **AC-2** — *Given* istri yang membayar listrik 780000 dari `BNK-I` tertaut *Then* tidak ditandai juga; siapa pun yang membayar, hasilnya sama.
- **AC-3** — *Given* "air" berkesepakatan 250000 *When* pembayar mencatat air 310000 tertaut (29 November) *Then* entri tersimpan dan ditandai sebagai di atas kesepakatan dengan keterangan 250000; kelebihannya tidak dibandingkan dengan ambang; tidak ada pemberitahuan langsung; pasangan melihatnya di ringkasan sebagai bagian dari satu baris tanpa hitungan; dari tujuh entri CTH-41 hanya ini yang bertanda.
- **AC-4** — *Given* "listrik" berkesepakatan 400000 (CTH-05 varian) *When* listrik 450000 dibayar tertaut *Then* entri ditandai di atas kesepakatan, walaupun 450000 di bawah ambang 500000.
- **AC-5** — *Given* "listrik" yang usulannya belum dijawab (CTH-05 varian) *When* listrik 450000 dibayar tertaut *Then* aturan ambang berlaku dan entri tidak ditandai (450000 ≤ 500000).
- **AC-6** — *Given* "internet" berkesepakatan 350000 *When* suami membayarnya 400000 dari dana pribadinya sebagai kontribusi tertaut *Then* entri tidak pernah ditandai dan pengingat periode itu mati.
- **AC-7** — *Given* pembayaran cicilan atau penyisihan *When* pencatat mencoba menautkannya ke "listrik" *Then* tautannya ditolak dan entrinya tetap tersimpan tanpa tautan.
- **AC-8** — *Given* istri mencatat pengeluaran 600000 tanpa tautan dan entrinya ditandai dasar ambang, sementara "listrik" berkesepakatan 900000 *When* istri menautkannya ke "listrik" *Then* tanda ambangnya hilang (600000 ≤ 900000). *Given* jumlah per kali 500000 *Then* tandanya menjadi di atas kesepakatan dengan keterangan 500000. *When* suami, bukan pencatatnya, mencoba menautkan entri itu *Then* perintahnya ditolak.
- **AC-9** — *Given* suami mencatat pembayaran listrik 780000 saat offline, dan istri mengakhiri kesepakatan "listrik" yang tiba di server lebih dulu *When* pembayaran tiba *Then* entri tersimpan dan dinilai dengan aturan ambang (ditandai, 780000 > 500000); sebelum tiba, perangkat suami menampilkan "menunggu terkirim" tanpa status tanda. Bila pembayaran tiba lebih dulu, ia tidak ditandai, dan pengakhiran sesudahnya tidak mengubah itu.
- **AC-10** — *Given* pembayaran AC-1 terkirim dua kali dengan kunci yang sama *Then* hanya ada satu entri dan `BNK-S` 1120000.
- **AC-11** — *Given* tanggal kejadian di masa depan, atau sebelum saldo awal `BNK-S` *Then* entri ditolak di batas dengan keterangan pada medan tanggal.
- **AC-12** — *Given* AC-1 dibuka dari pengingat listrik jatuh tempo 20 November *Then* pembayaran menunjuk kejadian 20 November dan hanya pengingat kejadian itu yang mati; kejadian 20 Desember tetap diingatkan. *Given* kejadian 20 Oktober belum dibayar dan pembayaran dibuka dari layar pos *Then* bawaannya kejadian 20 Oktober (yang paling awal belum dibayar) dan pencatat dapat memilih 20 November.
- **AC-13** — *Given* tagihan "internet" terarsip *When* pengeluaran 350000 dicatat dengan tautan ke "internet" *Then* tautannya ditolak, entrinya tetap tersimpan tanpa tautan dan dinilai dengan aturan ambang (tidak ditandai, 350000 ≤ 500000), dan layar menawarkan mengaktifkan kembali tagihan itu lalu menautkannya lewat tautan susulan (07 INV-26; 06 §6.3 "Pos rutin" baris (f)).
- **AC-14** — **Tautan hanya bila kategori sama** (07 §3.15 aturan 7b, INV-37; 06 INV-A52; 10.3 SEC-92). *Given* pos "Listrik" berkategori "listrik", *When* pencatat menautkan entri berkategori "belanja bulanan" ke pos itu, saat mencatat maupun lewat tautan susulan, *Then* tautannya ditolak dengan `category_mismatch`, entrinya tetap tersimpan tanpa tautan dan dinilai terhadap ambang seperti entri biasa; *When* entri dua baris ("listrik" dan "belanja bulanan") ditautkan, *Then* ditolak juga, karena setiap baris `expense` entri tertaut harus berkategori pos itu. Pemeriksaannya satu fungsi server yang sama untuk kedua jalan (06 INV-A52). CTH-50 tidak berubah: entri tokennya berkategori "listrik".

### FR-FIN-104 — Status kesepakatan hanya di layar pos itu
| | |
|---|---|
| Fitur | F-FIN-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6; MSL-7 (tidak menjadi alat tekan) |
| Aktor dan izin | Kedua anggota dewasa melihat hal yang sama (AT-2) |
| Aturan acuan | 06 §6.3 "Pos rutin" baris (a), "Yang ditampilkan pada tanda"; INV-A43 butir f; AT-4 |
| Business rule | BR-FIN-097 |
| Kedudukan | **[Usulan]** `security-privacy` (SP-35), diterapkan 06; belum disajikan kepada pemilik produk |

Status kesepakatan suatu pos (diusulkan, berlaku, belum sepakat, ditarik, diakhiri) dan riwayatnya hanya tampil di layar pos itu. Usulan yang menunggu jawaban seseorang tampil kepadanya di tempat tanggapan yang diatur REQ.md, sebagai butir satu per satu, tanpa hitungan status.

**Kriteria penerimaan**
- **AC-1** — *Given* "listrik" berstatus belum sepakat *When* salah satu membuka layar pembayaran listrik, halaman depan, daftar tagihan, atau ringkasan *Then* status itu tidak tampil di sana; ia hanya tampil di layar pos "listrik".
- **AC-2** — *Given* beberapa pos dengan status berbeda *Then* tidak ada penyaring, urutan, hitungan, daftar per orang, atau kolom ekspor menurut status kesepakatan.
- **AC-3** — *Given* entri air 310000 bertanda di atas kesepakatan *Then* tandanya hanya menyebut 250000 sebagai jumlah yang disepakati saat dinilai; tidak ada keterangan berapa kali pos itu pernah di atas kesepakatan.
- **AC-4** — *Given* suami dan istri membuka layar pos "air" *Then* keduanya melihat isi yang sama.

### FR-FIN-105 — Mengubah jumlah yang disepakati dan tawaran satu ketukan
| | |
|---|---|
| Fitur | F-FIN-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6; MSL-5 |
| Aktor dan izin | Usulan pengganti: anggota dewasa mana pun; disetujui pihak lainnya. Tawaran satu ketukan: hanya pencatat pembayaran yang melebihi kesepakatan (tidak dapat dipertukarkan). Simetris untuk usulan dan persetujuan (PG-01) |
| Aturan acuan | 06 §6.3 K01 `replaces`, K05, "Pos rutin" baris (c) dan "Sesudah pembayaran di atas jumlah yang disepakati"; INV-A19; 07 §3.15 butir 7 |
| Business rule | BR-FIN-055, BR-FIN-060 |
| Kedudukan | Usulan pengganti: **[Usulan]** `architect`. Tawaran satu ketukan: **[Usulan]** `security-privacy` (SP-37). Tawaran disembunyikan bila sudah ada usulan menunggu: **[Usulan]** dokumen ini |

Sesudah pembayaran tertaut dengan jumlah pemakaian di atas `R`, pencatatnya ditawari satu ketukan untuk mengusulkan jumlah yang dibayar sebagai `R` baru. Sukarela; diabaikan tanpa akibat. Bila pos itu sudah punya usulan yang menunggu (paling banyak satu per sasaran, INV-A19), tawaran tidak tampil dan layar menunjuk usulan yang ada.

**Kriteria penerimaan**
- **AC-1** — *Given* air berkesepakatan 250000 dan pembayaran 310000 bertanda (FR-FIN-103/AC-3) *When* pencatatnya memakai tawaran *Then* tercatat usulan pengganti 310000; pasangannya menerima N-11 dengan batas SP-36; kesepakatan 250000 tetap berlaku sampai disetujui.
- **AC-2** — *Given* AC-1 *Then* pasangan yang bukan pencatat pembayaran itu tidak melihat tawaran tersebut.
- **AC-3** — *Given* usulan 310000 belum disetujui *When* air bulan berikutnya dibayar 320000 *Then* ia dinilai terhadap 250000 dan ditandai. *When* usulan 310000 kemudian disetujui *Then* pembayaran 310000 dan 320000 yang sudah tersimpan tetap bertanda; pembayaran sesudahnya dinilai terhadap 310000.
- **AC-4** — *Given* listrik berkesepakatan 900000 *When* istri mengusulkan 950000 dan suami menyetujuinya *Then* kesepakatan 950000 menggantikan 900000 sejak diterima server.
- **AC-5** — *Given* pembayaran ≤ `R`, atau pos tanpa kesepakatan berlaku, atau pos dengan usulan yang menunggu *Then* tawaran tidak tampil.

### FR-FIN-106 — Mengakhiri kesepakatan pos rutin
| | |
|---|---|
| Fitur | F-FIN-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6; MSL-7 (persetujuan dapat ditarik) |
| Aktor dan izin | Salah satu anggota dewasa, sendirian (K06). Simetris (PG-01). **Ditandai menurut tugas agen ini:** pengakhiran sepihak mengubah perlakuan atas pembayaran yang dilakukan pasangan, sehingga dapat menambah pengawasan atas orang dewasa lain (06 PG-14; 03 §7 butir 9) |
| Aturan acuan | 06 §6.3 K06, "Pos rutin" baris (d), "Batas dan beban" (mengakhiri sebagai jalan keluar), §12 uji bersamaan; N-12; 07 §3.15 butir 5 dan 7 |
| Business rule | BR-FIN-053, BR-FIN-055, BR-FIN-061 |
| Kedudukan | Usulan `architect` (06 P-26; 03 §7 butir 9; Q-113), **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026), S3-5 (i); harganya W-25; tetap dinilai terhadap PG-14 di dokumen 10 (K-34, S4) |

Pengakhiran berlaku seketika saat diterima server dan ke depan saja; jadwal dan pengingat tetap; mengarsipkan adalah tindakan tersendiri (FR-FIN-107); menyepakatinya lagi adalah usulan baru. Sebelum mengukuhkan, layar menyebut akibatnya: pembayaran pos ini sesudahnya mengikuti aturan ambang, yang di atas ambang akan ditandai, dan pembayaran yang sudah tersimpan tidak berubah. Pasangan menerima N-12.

**Kriteria penerimaan**
- **AC-1** — *Given* listrik berkesepakatan 900000, ambang 500000, dan pembayaran 20 November 780000 tanpa tanda *When* istri mengakhiri kesepakatan sendirian *Then* state menjadi diakhiri seketika; suami menerima N-12; pembayaran listrik berikutnya 780000 oleh suami ditandai dasar ambang; pembayaran 20 November tetap tanpa tanda.
- **AC-2** — *Given* air berkesepakatan 250000 *When* suami mengakhirinya sendirian *Then* pembayaran air berikutnya 310000 tidak ditandai (≤ ambang 500000); pengakhirannya terlihat di layar pos dan diberitahukan kepada istri.
- **AC-3** — *Given* pengakhiran dan pembayaran tiba bersamaan *Then* yang lebih dulu diterima server yang berlaku; tanda entri yang sudah tersimpan tidak berubah.
- **AC-4** — *Given* layar konfirmasi pengakhiran *Then* akibat di atas tampil sebelum pengakhiran dapat dikukuhkan.
- **AC-5** — *Given* kesepakatan diakhiri *Then* tagihannya tetap ada, pengingatnya tetap berjalan, dan tagihan itu tidak terarsip.

### FR-FIN-107 — Tanggal akhir, pengarsipan, dan penghapusan tagihan
| | |
|---|---|
| Fitur | F-FIN-011 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6 |
| Aktor dan izin | Anggota dewasa mana pun mengarsipkan, mengaktifkan kembali, mengganti nama, atau menghapus tagihan yang belum pernah dipakai (07 §2.6 "Selalu boleh", "Mengaktifkan kembali"; 06 §6.3 "Pos rutin" baris (f)). Simetris (PG-01). **Ditandai menurut tugas agen ini:** mengarsipkan sendirian juga mengakhiri kesepakatan rutin yang disetujui berdua, sama dengan FR-FIN-106 (06 P-26) |
| Aturan acuan | 07 §3.15 (tanggal akhir, Q-46), §2.6 tabel pengarsipan dan "Mengaktifkan kembali", INV-26, T-36; 06 §6.3 "Pos rutin" baris (d) dan (f), K04, K06, N-12, PG-04 |
| Business rule | BR-FIN-050, BR-FIN-061 |
| Kedudukan | Tanggal akhir: D-12 (diterima di checkpoint S1). Pengarsipan selalu boleh: 07 §2.6; keputusan 9 dan D-26. Tagihan tidak terarsip sendiri: **[Usulan]** `finance-domain`, putaran S3 (07 §3.15, T-36). Pengarsipan beserta K06 dalam satu transaksi dan pengaktifan kembali yang tidak menghidupkan kesepakatan: **[Usulan]** `architect`, revisi kesembilan (06 §6.3 "Pos rutin" baris (f); menutup G-04) |

Sesudah tanggal akhir lewat, tidak ada kejadian baru dan pengingat berhenti; **tagihan tidak terarsip sendiri** dan tetap tampil sampai seorang anggota mengarsipkannya. Pembayaran terlambat atas kejadian yang jatuh tempo sampai tanggal akhir tetap dapat ditautkan dan dinilai seperti biasa (07 §3.15). Tagihan yang belum pernah dipakai benar-benar dapat dihapus; yang sudah dipakai hanya diarsipkan. Ganti nama adalah Peristiwa metadata yang tampil di riwayat pos. Pembayaran lama dan tautannya tetap terlihat di riwayat dan ringkasan.

**Mengarsipkan** adalah satu perintah oleh pengarsip. Bila ada kesepakatan rutin yang berlaku, perintah itu juga mengakhirinya (K06) dalam transaksi yang sama: dua Peristiwa berpelaku pengarsip (tagihan diarsipkan, kesepakatan diakhiri) dan satu N-12 (`kabar`) kepada anggota dewasa lainnya yang menyebut keduanya. Kesepakatan yang masih diusulkan atas tagihan itu ditarik (K04) bila pengarsip pengusulnya, atau gugur dengan Peristiwa berpelaku pengarsip bila bukan; tanpa notifikasi kedua. Sebelum mengukuhkan, layar menyebut bahwa kesepakatannya ikut berakhir. Tujuan "siapkan uang" milik tagihan itu ikut diarsipkan bila alokasinya nol; bila belum, alokasinya ditawarkan untuk dilepas di langkah yang sama, tanpa pilihan terisi (07 §2.6). Tagihan terarsip tidak menerima tautan atau alokasi baru (07 INV-26).

**Mengaktifkan kembali** adalah Peristiwa berpelaku tanpa entri, terlihat seperti pengarsipannya, tanpa notifikasi (07 §2.6). Ia **tidak** menghidupkan kesepakatan lama: pembayaran sesudahnya mengikuti aturan ambang sampai usulan baru (K01) disetujui (06 §6.3 "Pos rutin" baris (f), PG-04).

**Kriteria penerimaan**
- **AC-1** — *Given* "langganan kerja" berkesepakatan 200000 **[Asumsi]**, jatuh tempo tanggal 25 **[Asumsi]**, bertanggal akhir 31 Desember 2026 *When* tanggal itu lewat menurut waktu server *Then* tidak ada kejadian sesudah 25 Desember dan tidak ada pengingat lagi; tagihan **tidak** terarsip dan kesepakatannya tetap berlaku; pembayaran November dan Desember beserta tautannya tetap tampil di riwayat dan ringkasan periodenya. *When* pembayaran 200000 untuk kejadian 25 Desember yang belum dibayar dicatat terlambat pada 5 Januari 2027 *Then* tautannya diterima dan dinilai terhadap 200000 (tidak ditandai).
- **AC-2** — *Given* tagihan "air" yang dibuat keliru dan belum pernah dibayar *When* pembuatnya menghapusnya *Then* tagihan dan usulannya hilang. *Given* tagihan yang sudah pernah dibayar *Then* pilihan yang tersedia hanya arsip.
- **AC-3** — *Given* "wifi" diganti nama menjadi "internet rumah" *Then* riwayat pos menampilkan Peristiwa ganti nama; jumlah yang disepakati tidak berubah; pembayaran lama tampil dengan nama barunya.
- **AC-4** — *Given* tanggal akhir sebelum hari ini *When* disimpan *Then* ditolak di batas.
- **AC-5** — *Given* "listrik" berkesepakatan 900000 *When* istri mengarsipkannya *Then* layar menyebut lebih dulu bahwa kesepakatannya ikut berakhir; sesudah dikukuhkan tercatat dua Peristiwa berpelaku istri (tagihan diarsipkan, kesepakatan diakhiri) dalam satu transaksi, suami menerima tepat satu N-12, dan tidak ada Peristiwa berpelaku `system`. *Given* suami yang mengarsipkan *Then* hal yang sama dengan peran tertukar.
- **AC-6** — *Given* "air" yang usulan 250000-nya dari suami belum dijawab *When* suami mengarsipkannya *Then* usulannya ditarik (K04) dalam transaksi yang sama; *When* istri yang mengarsipkannya *Then* usulan suami gugur dengan Peristiwa berpelaku istri; dalam kedua hal hanya ada satu N-12.
- **AC-7** — *Given* "listrik" terarsip dengan kesepakatan 900000 yang diakhiri saat pengarsipan *When* suami mengaktifkannya kembali dan membayar listrik 780000 tertaut *Then* pengaktifan kembali tercatat sebagai Peristiwa berpelaku suami tanpa entri dan tanpa notifikasi; pembayaran itu dinilai dengan aturan ambang dan ditandai (780000 > 500000); kesepakatan 900000 tidak hidup lagi, dan layar pos menawarkan mengusulkannya kembali (K01).
- **AC-8** — *Given* "listrik" belum terarsip dan alokasi listrik 120000 **[Asumsi]** masih ada *When* pengarsipan dibuka *Then* layar menawarkan melepas 120000 di langkah yang sama tanpa pilihan terisi; tagihan tetap dapat diarsipkan, tetapi tujuan "listrik" ikut diarsipkan hanya bila alokasinya 0 (07 §2.6).

### FR-FIN-108 — Menyiapkan uang untuk tagihan periode ini
| | |
|---|---|
| Fitur | F-FIN-011 (memakai F-FIN-014) |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-6, MSL-1 (uang yang sudah ditujukan terlihat terpisah dari uang bebas) |
| Aktor dan izin | Anggota dewasa mana pun atas tempat uang dana bersama yang terlihat olehnya; alokasi `adults` (06 §9.2). Simetris (PG-01) |
| Aturan acuan | 07 §3.15 "Tiga hal yang dituntut pola itu" (alokasi, bukan reservasi; jumlah tawaran) dan "Kejadian tagihan dan 'belum dibayar periode ini'" baris "Sudah disiapkan", §2.6 `purpose.recurring_bill_id`, §3.9, INV-17; T-28 |
| Business rule | BR-FIN-047 (`FIN-1.md`), BR-FIN-071, BR-FIN-072, BR-FIN-073 |
| Kedudukan | Alokasi untuk tagihan dan jumlah tawarannya: **[Usulan]** `finance-domain` (T-28, Q-117). Jumlah yang masih ditawarkan sesudah dikurangi alokasi: **[Usulan]** `finance-domain`, putaran S3 (07 §3.15; PA-FIN-5). Semua tagihan termasuk sebagai bawaan dan tawaran melepas sisa: **[Usulan]** dokumen ini. Fitur ini pilihan; tagihan tetap dapat dibayar tanpanya |

Satu langkah "siapkan uang untuk semua tagihan yang belum dibayar periode ini" (syarat layar 07 §3.15): menampilkan tiap tagihan yang punya kejadian belum dibayar menurut BR-FIN-047, dengan **jumlah yang masih ditawarkan** = max(0, jumlah tawaran − alokasi tujuan tagihan itu di semua tempat uang dana bersama); jumlah tawaran = `R` bila kesepakatannya berlaku, selain itu `estimated_amount`. Tagihan yang jumlah masih ditawarkannya 0 tidak ikut. Tiap baris dapat diubah atau dikeluarkan; satu tempat uang dipilih. Ini label pada uang, bukan pemakaian: tidak ada tanda, persetujuan, atau perubahan uang yang ada. Pembayaran yang membawa tujuan tagihan itu mengurangi alokasinya. Bila sisa alokasi tertinggal sesudah tagihan dibayar, pencatatnya ditawari melepasnya, tanpa pilihan terisi; tidak ada yang dilepas otomatis.

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-41 sesudah internet dan iuran dibayar: `BNK-S` 7900000 tanpa alokasi; listrik (`R` 900000) dan air (`R` 250000) belum dibayar *When* suami menyiapkan uang dari `BNK-S` *Then* tercatat alokasi listrik 900000 dan air 250000; belum dialokasikan 6750000; uang yang ada dan kekayaan bersih tidak berubah.
- **AC-2** — *Given* `BNK-S` 1900000 dengan alokasi listrik 900000 dan air 250000 *When* listrik 780000 dibayar membawa tujuan "listrik" *Then* `BNK-S` 1120000, alokasi listrik 120000, dan pencatat ditawari melepas 120000. *When* ia menerimanya *Then* belum dialokasikan 870000.
- **AC-3** — *Given* `BNK-S` 1120000 dan alokasi air 250000 sesudah sisa listrik dilepas *When* air 310000 dibayar membawa tujuan "air" *Then* `BNK-S` 810000; alokasi air turun ke 0, bukan negatif; kelebihan 60000 mengurangi bagian belum dialokasikan menjadi 810000; tidak ada yang ditolak.
- **AC-4** — *Given* tempat uang dengan saldo tercatat 1000000 tanpa alokasi **[Asumsi]** *When* menyiapkan 1150000 untuk dua tagihan *Then* langkah itu ditolak seluruhnya, layar menyebut bahwa yang dapat dialokasikan 1000000, dan pencatat dapat mengeluarkan atau mengubah satu baris.
- **AC-5** — *Given* "listrik" yang usulan 900000-nya belum dijawab dan perkiraannya 600000 *When* uang disiapkan *Then* jumlah tawaran listrik 600000; sesudah kesepakatan 900000 berlaku, tawarannya 900000.
- **AC-6** — *Given* AC-1 sudah dijalankan (alokasi listrik 900000 dan air 250000 atas `BNK-S`) *When* suami membuka langkah itu lagi sebelum membayar *Then* jumlah yang masih ditawarkan listrik dan air 0, keduanya tidak ikut, dan belum dialokasikan tetap 6750000. *Given* AC-5 dengan alokasi listrik 600000 yang disiapkan sebelum kesepakatan berlaku *When* kesepakatan 900000 berlaku dan langkah dibuka lagi *Then* yang masih ditawarkan untuk listrik 300000 (900000 − 600000) (07 §3.15).

## 3. F-FIN-012 — Cicilan

Masalah: `MSL-3`, `MSL-1`; `KB-06`. Beban berulang: sekali per cicilan; per pembayaran sama dengan mencatat satu pengeluaran (F-FIN-005). Yang dicatat fitur ini: cicilan dengan **sisa yang harus dibayar** (jumlah seluruh pembayaran yang tersisa, tanpa memisahkan pokok dan bunga) dan, untuk pembelian yang dicatat di aplikasi, biaya kreditnya (07 §3.14; rumusan yang diminta 07 §11 untuk 04 F-FIN-012).

### FR-FIN-110 — Mendaftarkan cicilan yang sudah berjalan
| | |
|---|---|
| Fitur | F-FIN-012 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, MSL-3 |
| Aktor dan izin | Anggota dewasa mana pun untuk cicilan dana bersama; cicilan dan kewajibannya `adults` (06 §9.2), sama bagi keduanya (07 §5.2). Simetris (PG-01) |
| Aturan acuan | 07 §3.14, §3.17, §2.6 `installment`, §3.8 baris `opening` kewajiban; CTH-01 |
| Business rule | BR-FIN-063, BR-FIN-065, BR-FIN-098 |
| Kedudukan | 07 §3.14 (Q-16, K-11). Penilaian satu kali atas saldo awal sesudah ada ambang: 06 §6.2 putusan 2 |

Masukan: nama, pihak pemberi kredit, jumlah per pembayaran, jumlah pembayaran seluruhnya, jumlah yang sudah dibayar, tanggal jatuh tempo bulanan, kategori (`installment.category_id`, dipilih rumah tangga; kategori contoh "cicilan" terisi bila ada dan tidak terarsip **[Usulan]** dokumen ini, tetapi bukan kategori yang diwajibkan model; boleh kosong; 07 §3.14, T-10), penanggung jawab (opsional), aset terkait (opsional). Validasi di batas: semua bilangan bulat; jumlah per pembayaran > 0; yang sudah dibayar < seluruhnya. Yang dicatat: `installment` dan saldo awal kewajiban sebesar jumlah per pembayaran × sisa pembayaran. Saldo awal bukan pemasukan dan tidak masuk anggaran atau arus kas.

**Kriteria penerimaan**
- **AC-1** — *Given* onboarding rumah tangga CTH-01 sebelum ada ambang *When* cicilan motor dicatat 1500000 per bulan dengan 6 pembayaran tersisa *Then* kewajiban `KWJ-C` 9000000 dan kekayaan bersih 9100000; entri itu tidak ditandai.
- **AC-2** — *Given* ambang 500000 sudah berlaku *When* cicilan baru didaftarkan sebagai saldo awal 300000 × 10 = 3000000 **[Asumsi]** *Then* saldo awalnya dibandingkan dengan ambang satu kali dan ditandai; pembayaran sesudahnya tidak dibandingkan.
- **AC-3** — *Given* jumlah yang sudah dibayar sama dengan atau lebih dari jumlah seluruhnya, atau jumlah per pembayaran 0 *Then* ditolak di batas.
- **AC-4** — *Given* pendaftaran terkirim dua kali dengan kunci yang sama *Then* hanya ada satu cicilan dan satu kewajiban.

### FR-FIN-111 — Membeli barang secara cicilan
| | |
|---|---|
| Fitur | F-FIN-012 (dengan F-FIN-016) |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, MSL-3 |
| Aktor dan izin | Anggota dewasa mana pun; di atas ambang mengikuti alur permintaan (REQ.md) atau ditandai. Simetris (PG-01) |
| Aturan acuan | 07 §3.14 "Biaya kredit" dan "Biaya kredit tidak pernah negatif", §3.16, §3.8 baris `asset_purchase` secara cicilan, T-35; CTH-34, CTH-49; 06 §6.2 |
| Business rule | BR-FIN-063, BR-FIN-064, BR-FIN-065, BR-FIN-076 |
| Kedudukan | 07 §3.14. Penolakan bila uang muka + cicilan < harga, dan cicilan 0% tanpa baris biaya kredit: **[Usulan]** `finance-domain`, putaran S3 (07 §3.14, T-35; menutup G-07), mengadopsi bawaan berkas ini |

Masukan: barang, harga, uang muka, tempat uang uang muka, jumlah per cicilan, jumlah cicilan, pemberi kredit. Biaya kredit dihitung, tidak diisi: uang muka + jumlah seluruh cicilan − harga. Satu entri: aset + harga, uang − uang muka, kewajiban + seluruh cicilan, dan baris biaya kredit bila biaya kreditnya lebih dari 0; pada cicilan 0% tidak ada baris biaya kredit. Bila uang muka + seluruh cicilan lebih kecil daripada harga, perintah ditolak dan layar menawarkan harga = uang muka + seluruh cicilan; bila rumah tangga menilai barang itu lebih tinggi, selisihnya dicatat sebagai penilaian ulang aset (FR-FIN-131), bukan biaya kredit (07 §3.14).

**Kriteria penerimaan**
- **AC-1** — *Given* `TUN-I` 825000, uang yang ada 7544000, aset 13000000, `KWJ-C` 7500000, kekayaan bersih 12644000 (CTH-33 akhir) *When* istri membeli kulkas seharga 4000000 dengan uang muka 400000 dari `TUN-I` dan 12 × 330000 *Then* `TUN-I` 425000, uang yang ada 7144000, aset 17000000, `KWJ-C` 11460000, biaya kredit 360000, kekayaan bersih 12284000.
- **AC-2** — *Given* AC-1 *Then* jumlah pemakaiannya 4360000 dan dibandingkan dengan ambang satu kali saat dibeli; tidak satu pun dari 12 pembayarannya dibandingkan.
- **AC-3** — *Given* AC-1 *Then* ringkasan periode itu menampilkan baris "biaya kredit" −360000 di tabel A dan tidak menghitungnya terhadap anggaran mana pun.
- **AC-4** — *Given* CTH-49 bagian A: `TUN-I` 500000, uang yang ada 500000, aset 0, `KWJ-C` 0, kekayaan bersih 500000; barang berlabel 4000000, uang muka 0, 12 × 300000 = 3600000 *When* disimpan dengan harga 4000000 *Then* ditolak tanpa akibat apa pun (0 + 3600000 < 4000000) dan layar menawarkan harga 3600000. *When* dicatat ulang dengan harga 3600000 *Then* aset 3600000, `KWJ-C` 3600000, `TUN-I` dan uang yang ada tetap 500000, kekayaan bersih 500000, tidak ada baris biaya kredit, dan jumlah pemakaian 3600000 dibandingkan dengan ambang satu kali (di atas ambang 500000: diajukan lebih dulu atau bertanda). *When* rumah tangga menilai barang itu 4000000 *Then* mereka mencatat penilaian ulang +400000 yang tampil sebagai perubahan nilai aset, bukan biaya kredit.
- **AC-5** — *Given* `Sebelum` yang sama dengan AC-4 (CTH-49 bagian B) *When* istri membeli barang seharga 4000000 secara cicilan 0% dengan uang muka 400000 dari `TUN-I` dan 12 × 300000 *Then* `TUN-I` 100000, uang yang ada 100000, aset 4000000, `KWJ-C` 3600000, kekayaan bersih 500000; tidak ada baris biaya kredit di tabel A; jumlah pemakaian 4000000 dibandingkan dengan ambang satu kali.

### FR-FIN-112 — Membayar cicilan
| | |
|---|---|
| Fitur | F-FIN-012 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-3, MSL-6 |
| Aktor dan izin | Anggota dewasa mana pun, siapa pun penanggung jawabnya. Simetris (PG-01) |
| Aturan acuan | 07 §3.14, §3.8 baris `repayment_out`, INV-13; CTH-04; 06 §6.2 daftar yang tidak ditandai |
| Business rule | BR-FIN-062, BR-FIN-066, BR-FIN-098 |
| Kedudukan | Pembayaran cicilan tidak ditandai dan tanpa kesepakatan: D-44, diterima pemilik produk lewat "lanjut" di checkpoint S2; kalimat keduanya bagian penegasan 3 |

Pembayaran adalah `repayment_out` berkategori `installment.category_id`, terhitung terhadap anggaran kategori itu dan tampil di pengeluaran per kategori. Bila kategori itu kosong atau terarsip, pembayaran tetap tersimpan tanpa kategori, layar menawarkan memilih kategori, dan pembayaran itu tidak terhitung terhadap anggaran mana pun tetapi tampil di baris "tanpa kategori" tabel B (07 §3.14, §6.3; FIN-1.md FR-FIN-080/AC-7). Denda keterlambatan adalah baris pengeluaran di entri yang sama. Pengingat jatuh tempo mengikuti F-CAL-002 **[Asumsi]**.

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-04: `BNK-S` 10500000, uang yang ada 16100000, `KWJ-C` 9000000, kekayaan bersih 19100000, ambang 500000 *When* suami membayar cicilan motor 1500000 dari `BNK-S` *Then* `BNK-S` 9000000, uang yang ada 14600000, `KWJ-C` 7500000, kekayaan bersih tetap 19100000; anggaran "cicilan" terpakai 1500000, sisa 0; tersisa 5 pembayaran; entri tidak ditandai walaupun di atas ambang.
- **AC-2** — *Given* istri yang membayar cicilan itu dari `BNK-I` *Then* hasilnya sama dan tidak ditandai.
- **AC-3** — *Given* pembayaran 1500000 dengan denda 50000 **[Asumsi]** *Then* bagian 1500000 tidak dinilai; baris denda 50000 dinilai terhadap ambang (tidak ditandai, 50000 ≤ 500000), masuk pengeluaran kategori, dan menurunkan kekayaan bersih 50000.
- **AC-4** — *Given* sisa kewajiban 1500000 *When* dicatat pelunasan 3000000 *Then* ditolak di batas karena melebihi sisa.
- **AC-5** — *Given* pembayaran terakhir tercatat *Then* kewajiban 0 dan cicilan dapat diarsipkan (07 §2.6).
- **AC-6** — *Given* cicilan motor yang kategorinya "cicilan" sudah diarsipkan *When* suami membayar 1500000 *Then* entri tersimpan tanpa kategori, layar menawarkan memilih kategori, tidak ada anggaran yang berubah, dan 1500000 tampil satu kali di baris "tanpa kategori" tabel B; `BNK-S`, `KWJ-C`, dan kekayaan bersih berubah sama dengan AC-1.

### FR-FIN-113 — Melunasi cicilan lebih awal dengan potongan
| | |
|---|---|
| Fitur | F-FIN-012 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Anggota dewasa mana pun. Simetris (PG-01) |
| Aturan acuan | 07 §3.14 "Pelunasan dipercepat", §6.3 tabel A baris "utang dibebaskan pihak luar", INV-13 |
| Business rule | BR-FIN-062, BR-FIN-066 |
| Kedudukan | 07 §3.14 |

Satu langkah: jumlah yang dibayar dan potongan. Jumlah keduanya tidak boleh melebihi sisa kewajiban.

**Kriteria penerimaan**
- **AC-1** — *Given* `KWJ-C` 7500000 sesudah CTH-04 *When* suami melunasi dengan membayar 7000000 dan potongan 500000 **[Asumsi]** *Then* uang yang ada −7000000, `KWJ-C` 0, kekayaan bersih +500000 di tabel A sebagai "utang dibebaskan pihak luar"; pembayarannya tidak ditandai.
- **AC-2** — *Given* sisa 7500000 *When* dibayar 7000000 dengan potongan 600000 *Then* ditolak di batas karena 7600000 > 7500000.

## 4. F-FIN-013 — Utang, piutang, dan bantuan keluarga besar

Masalah: `MSL-3`; `KB-07`, `KB-21`. Beban berulang: per kejadian, sesekali **[Asumsi]**. Pihak luar dicatat dengan nama panggilan dan hubungan (07 §2.6 `external_party`; F-FAM-007).

### FR-FIN-115 — Memberi pinjaman kepada pihak luar dan menerima pelunasannya
| | |
|---|---|
| Fitur | F-FIN-013 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-3, MSL-1 |
| Aktor dan izin | Anggota dewasa mana pun dari dana bersama; piutang `adults` (06 §9.2). Pelunasan dicatat ke tempat uang mana pun; bila masuk ke tempat uang yang dipegang anggota lain, berlaku konfirmasi custody (FIN-1.md). Simetris (PG-01) |
| Aturan acuan | 07 §3.7, §3.8 baris `loan_given`, INV-13; CTH-25 |
| Business rule | BR-FIN-066, BR-FIN-068, BR-FIN-098 |
| Kedudukan | 07 §3.7 (Q-47); D-17 |

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-25: uang yang ada 9584000, piutang 0, aset 12000000, `KWJ-C` 7500000, kekayaan bersih 14084000, ambang 500000 *When* istri meminjamkan 500000 dari dana bersama kepada adik suami *Then* uang yang ada 9084000, piutang 500000, kekayaan bersih tetap 14084000; pinjaman dinilai terhadap ambang dan tidak ditandai (500000 tidak melebihi ambang).
- **AC-2** — *Given* AC-1 *When* adik membayar 200000 ke `TUN-S` *Then* uang yang ada 9284000, piutang 300000; kekayaan bersih tetap.
- **AC-3** — *Given* sisa piutang 300000 *When* adik membayar 350000 *Then* piutang 0 dan kelebihan 50000 tercatat sebagai pemasukan "lainnya" **[Asumsi]** angka.
- **AC-4** — *Given* piutang kepada adik *Then* suami dan istri melihat sisa yang sama di daftar piutang dan di kekayaan bersih.

### FR-FIN-116 — Menerima pinjaman dari pihak luar dan membayarnya kembali
| | |
|---|---|
| Fitur | F-FIN-013 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-3, MSL-1 |
| Aktor dan izin | Anggota dewasa mana pun; kewajiban `adults`. Simetris (PG-01) |
| Aturan acuan | 07 §3.7 (cermin), §3.8 baris `repayment_out`, INV-13; CTH-28; D-26 |
| Business rule | BR-FIN-062, BR-FIN-066 |
| Kedudukan | Pinjaman yang diterima adalah utang, bukan pemasukan: D-26 |

**Kriteria penerimaan**
- **AC-1** — *Given* uang yang ada 9084000 dan kekayaan bersih 13784000 (CTH-28) *When* istri menerima pinjaman 1000000 dari kakaknya ke `TUN-I` *Then* uang yang ada 10084000, kewajiban `KWJ-U` 1000000, kekayaan bersih tetap 13784000; pinjaman tidak tampil sebagai pemasukan.
- **AC-2** — *Given* AC-1 *When* dibayar kembali 400000 *Then* uang yang ada 9684000, `KWJ-U` 600000; pembayaran kembali tidak dibandingkan dengan ambang (07 §3.8 baris `repayment_out`).
- **AC-3** — *Given* `KWJ-U` 600000 *When* dicatat pembayaran kembali 700000 *Then* ditolak di batas.

### FR-FIN-117 — Mengikhlaskan piutang
| | |
|---|---|
| Fitur | F-FIN-013 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, MSL-3 |
| Aktor dan izin | Anggota dewasa mana pun atas piutang dana bersama kepada pihak luar, juga bila pemberi pinjamannya pasangannya (CTH-25: istri meminjamkan, suami mengikhlaskan) (06 §6.2 baris `write_off`). Simetris (PG-01). Satu orang dapat menurunkan kekayaan bersih rumah tangga; pagarnya: penghapusan dibandingkan dengan ambang, ditandai bila di atasnya, dan pasangannya menerima N-10 seperti pemakaian lain (07 §3.7, T-32). Piutang dana bersama atas seorang anggota tidak dihapus di sini, melainkan dilepas lewat pelepasan klaim (FR-FIN-119; 06 §7.3) |
| Aturan acuan | 07 §3.7 `write_off`, §3.8 tabel *jumlah pemakaian* dan "Tanda atas entri" (`H`), §6.1, §6.3 tabel A, INV-26, INV-34, T-32; CTH-25 dan varian SP-51; 06 §6.2 baris `write_off`, N-10 |
| Business rule | BR-FIN-067, BR-FIN-070 |
| Kedudukan | 07 §3.7; D-17. Dibandingkan dengan ambang: **[Usulan]** `finance-domain`, putaran S3 (07 §3.7, T-32); **kumulatif per piutang**: **[Usulan]** `security-privacy` SP-51, diterapkan 07 putaran perbaikan S3 (§3.8, INV-34). Pelaku anggota dewasa mana pun dan larangan menautkan ke otorisasi pinjaman asal: **[Usulan]** `architect`, revisi kesembilan (06 §6.2); menutup G-03. Ditandai kumulatif: usulan `finance-domain` dan `security-privacy` SP-51 (07 §10.4 butir 3), diterima lewat "lanjut" di checkpoint S3, S3-9 (ii) |

Masukan: piutang, jumlah (≤ sisa), tanggal kejadian, **alasan (wajib)** (07 §3.7; 06 §6.2). Server menilai tandanya saat diterima, dengan dasar ambang: entri ditandai bila bagian yang belum disetujuinya `N` > 0 dan `H + N` > ambang, dengan `H` jumlah `N` penghapusan lain atas piutang yang sama yang tersimpan lebih dulu dan belum dibalik (07 §3.8 "Tanda atas entri", INV-34). Tanpa itu, mengikhlaskan dalam beberapa potongan di bawah ambang akan menghapus jumlah berapa pun tanpa tanda. Penghapusan di atas ambang boleh diajukan lebih dulu sebagai permintaan anggaran dan ditautkan ke otorisasinya, tetapi tidak ke otorisasi yang memuat pinjaman asal piutang itu: menyetujui meminjamkan bukan menyetujui mengikhlaskan (06 §6.2; 07 INV-26).

**Kriteria penerimaan**
- **AC-1** — *Given* piutang adik suami 300000, kekayaan bersih 14084000 (CTH-25) *When* suami mengikhlaskan 100000 dengan alasan *Then* piutang 200000, kekayaan bersih 13984000; ringkasan bulan tanggal kejadiannya menampilkan "piutang diikhlaskan" −100000 sebagai baris tersendiri; anggaran dan arus kas tidak berubah; entri tidak ditandai (100000 ≤ 500000).
- **AC-2** — *Given* sisa piutang 200000 *When* diikhlaskan 250000 *Then* ditolak di batas.
- **AC-3** — *Given* AC-1 dan kiriman keluarga 200000 pada periode yang sama *Then* "piutang diikhlaskan" dan "kiriman keluarga" tampil sebagai dua baris terpisah; istri melihat entri itu dengan pelakunya di riwayat piutang (FR-FIN-145).
- **AC-4** — *Given* AC-1 (sisa 200000, ambang 500000) *When* suami mengikhlaskan seluruh sisa 200000 *Then* tidak ditandai: `H` = 100000 dari AC-1, dan 100000 + 200000 = 300000 ≤ 500000 (CTH-25 catatan penilaian). *Given* piutang dana bersama 600000 **[Asumsi]** kepada pihak luar *When* seluruhnya diikhlaskan tanpa permintaan *Then* entri tersimpan, angka berubah seperti biasa, entri ditandai dengan dasar ambang (600000 > 500000), dan pasangan menerima N-10; tidak ada yang dihalangi (CTH-25 catatan penilaian).
- **AC-5** — *Given* formulir mengikhlaskan *When* alasan kosong *Then* ditolak di batas tanpa akibat apa pun.
- **AC-6** — *Given* permintaan meminjamkan 600000 **[Asumsi]** yang disetujui dan dipakai pinjaman asal *When* penghapusan 600000 yang sudah tersimpan dan bertanda ditautkan susulan ke otorisasi tersebut *Then* tautannya ditolak, entri dan tandanya tetap; *When* layar mengikhlaskan dibuka *Then* otorisasi itu tidak ditawarkan sebagai tautan, dan jalannya adalah permintaan baru.
- **AC-7** — *Given* CTH-25 varian SP-51: suami meminjamkan 700000 kepada adiknya dari `TUN-S` 1000000 lewat permintaan yang disetujui; sesudahnya `TUN-S` 300000, `PIU` 700000, `KB` 1000000, ambang 500000 *When* ia mengikhlaskan 300000 tanpa permintaan *Then* `H` 0, 0 + 300000 = 300000 tidak di atas 500000, tidak ditandai; `PIU` 400000, `KB` 700000. *When* ia mengikhlaskan 300000 lagi *Then* `H` 300000, 300000 + 300000 = 600000 > 500000, **ditandai** dan istri menerima N-10; `PIU` 100000, `KB` 400000. *When* ia mengikhlaskan 50000 *Then* `H` 600000, 600000 + 50000 = 650000 > 500000, ditandai; `PIU` 50000, `KB` 350000. `TUN-S` tetap 300000 sepanjang ketiganya, dan tiga baris "piutang diikhlaskan" tampil di ringkasan periode tanggal kejadian masing-masing (07 CTH-25 varian SP-51; INV-34).
- **AC-8** — *Given* AC-7 sampai penghapusan pertama, tetapi penghapusan kedua 300000 diajukan dulu lewat permintaan yang disetujui dan ditautkan ke otorisasi berbatas 300000 (bukan otorisasi pinjaman asal, AC-6) *When* tersimpan *Then* `N` = 300000 − min(300000, max(0, 300000 − 0)) = 0, tidak ditandai dan tidak menambah `H`; *When* penghapusan ketiga 50000 tanpa permintaan *Then* dinilai 300000 + 0 + 50000 = 350000, tidak di atas 500000, tidak ditandai. Angka saldo sama dengan AC-7 (`PIU` 50000, `KB` 350000).
- **AC-9** — *Given* AC-7 sampai penghapusan kedua (`PIU` 100000, `KB` 400000), lalu penghapusan kedua dibalik (`PIU` 400000, `KB` 700000) **[Asumsi]** *When* suami mengikhlaskan 50000 *Then* `H` = 300000 (hanya penghapusan pertama; yang dibalik tidak termasuk, 07 §3.8 "belum dibalik saat e disimpan"), 300000 + 50000 = 350000 tidak di atas 500000, tidak ditandai; `PIU` 350000, `KB` 650000. Tanda yang sudah lahir pada entri lain tidak dinilai ulang.

### FR-FIN-118 — Kiriman keluarga atau pinjaman
| | |
|---|---|
| Fitur | F-FIN-013 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-3; KS-3 (pengeluaran per kategori) |
| Aktor dan izin | Pencatat memilih per entri; kiriman mengikuti aturan pengeluaran biasa, termasuk ambang. Simetris (PG-01) |
| Aturan acuan | 07 §3.2 "Kiriman keluarga"; 06 §6.3 "Pos rutin" (kiriman orang tua bukan pos rutin di pilot); CTH-26 |
| Business rule | BR-FIN-068 |
| Kedudukan | Kiriman orang tua tidak rutin di pilot: pernyataan pemilik produk, checkpoint S2 (dokumen 17 kelompok (e)). Q-07 diterima |

Saat penerimanya saudara atau orang tua, layar menanyakan sekali apakah ini kiriman atau pinjaman, tanpa pilihan terisi **[Usulan]** dokumen ini, karena salah pilih mengubah kekayaan bersih. Kiriman tidak dilacak sesudahnya.

**Kriteria penerimaan**
- **AC-1** — *Given* uang yang ada 9284000 dan kekayaan bersih 13984000 (CTH-26) *When* suami mencatat bantuan untuk ayah 200000 dari `TUN-S` sebagai kiriman *Then* pengeluaran kategori "kiriman keluarga" 200000; uang yang ada 9084000; kekayaan bersih 13784000; tidak ada piutang.
- **AC-2** — *Given* penerima "adik suami" *When* pencatat belum memilih kiriman atau pinjaman *Then* entri belum dapat disimpan dan tidak ada pilihan yang terisi.
- **AC-3** — *Given* kiriman tercatat keliru sebagai pinjaman *Then* perbaikannya lewat koreksi FR-FIN-140, bukan penyuntingan.

### FR-FIN-119 — Utang antara dana bersama dan dana pribadi, dan pembebasannya
| | |
|---|---|
| Fitur | F-FIN-013 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, MSL-3 |
| Aktor dan izin | Dicatat anggota yang dana pribadinya tersentuh (peran tidak dapat dipertukarkan); kedua pihak selalu melihatnya (Q-24). Pembebasan utang anggota M kepada dana bersama: anggota dewasa selain M; pembebasan hak ganti M atas dana bersama: M sendiri. Antara dua dana pribadi: tidak dibangun di MVP (T-12) |
| Aturan acuan | 07 §3.7 "Utang antar-pemilik" (`inter_owner_loan`; dahulu "utang antar-pasangan", penyisiran S4, M7), §3.6, §3.8 baris `claim_waiver`, §6.3 tabel A ("utang dihapuskan", "kontribusi dari dana pribadi"), INV-08, INV-25 |
| Business rule | BR-FIN-069, BR-FIN-070, BR-FIN-093 |
| Kedudukan | 07 §3.6, §3.7; D-43 |

**Kriteria penerimaan**
- **AC-1** — *Given* istri meminjam 300000 **[Asumsi]** dari dana bersama ke dana pribadinya dengan janji kembali *When* ia mencatatnya *Then* satu entri menulis kedua sisi; suami melihat utang istri kepada dana bersama sebesar 300000; saldo dana pribadi istri tetap tidak terlihat suami.
- **AC-2** — *Given* AC-1 *When* suami mencoba mencatat pinjaman itu atas nama istri *Then* ditolak: sisi pribadi hanya ditulis pemiliknya.
- **AC-3** — *Given* utang istri 300000 dan ambang 500000 *When* suami membebaskannya *Then* ringkasan menampilkan "utang dihapuskan" 300000, bukan "disisihkan ke dana pribadi istri"; entri dinilai terhadap ambang sebesar 300000 dan tidak ditandai. *When* istri mencoba membebaskan utangnya sendiri *Then* ditolak.
- **AC-4** — *Given* suami memegang hak ganti 50000 **[Asumsi]** atas dana bersama *When* suami melepasnya *Then* kekayaan bersih rumah tangga naik 50000 sebagai "kontribusi dari dana pribadi" suami pada periode itu; entri tidak dinilai terhadap ambang. Berlaku sama bila istri yang melepas hak gantinya.

## 5. F-FIN-014 — Alokasi uang

Masalah: `MSL-1` (uang yang disisihkan lawan uang yang ada). Beban berulang: pilihan; per uang masuk bila dipakai. Anggaran tidak memerlukannya.

### FR-FIN-120 — Membuat alokasi
| | |
|---|---|
| Fitur | F-FIN-014 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Anggota dewasa mana pun atas tempat uang dana bersama; alokasi `adults`. Alokasi di tempat uang pribadi hanya oleh pemiliknya dan tertutup (V-1, 07 INV-25). Simetris (PG-01) |
| Aturan acuan | 07 §3.9, INV-17; CTH-09 |
| Business rule | BR-FIN-071, BR-FIN-073, BR-FIN-079 |
| Kedudukan | Batas alokasi: D-17 (Q-45, diterima di checkpoint S1) |

Masukan: tempat uang, tujuan, jumlah (bilangan bulat > 0). Layar menampilkan bagian belum dialokasikan tempat uang itu selama mengisi.

**Kriteria penerimaan**
- **AC-1** — *Given* `TUN-I` saldo 4300000 tanpa alokasi (CTH-09) *When* istri mengalokasikan 2500000 untuk "belanja bulanan" *Then* belum dialokasikan 1800000; uang yang ada dan kekayaan bersih tidak berubah.
- **AC-2** — *Given* AC-1 *When* suami mencoba menambah alokasi 2000000 untuk tujuan lain di `TUN-I` *Then* ditolak karena 4500000 > 4300000, dan layar menyebut 1800000 sebagai yang dapat dialokasikan.
- **AC-3** — *Given* `TAB-I` 5000000 *When* seluruhnya dialokasikan untuk "dana darurat" *Then* belum dialokasikan 0.

### FR-FIN-121 — Memakai, mengurangi, dan memindahkan alokasi
| | |
|---|---|
| Fitur | F-FIN-014 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Sama dengan FR-FIN-120 |
| Aturan acuan | 07 §3.9; CTH-10, CTH-41 |
| Business rule | BR-FIN-071, BR-FIN-072 |
| Kedudukan | 07 §3.9 |

Entri yang mengurangi tempat uang boleh membawa tujuan: alokasinya berkurang paling jauh sampai 0, sisanya mengurangi bagian belum dialokasikan. Entri tanpa tujuan mengurangi bagian belum dialokasikan. Pemindahan boleh membawa tujuan, sehingga alokasinya ikut pindah. Mengurangi alokasi tidak memindahkan uang.

**Kriteria penerimaan**
- **AC-1** — *Given* `TUN-I` 4300000 (alokasi 2500000, belum dialokasikan 1800000) dan `TAB-I` 5000000 (alokasi "dana darurat" 5000000) *When* 500000 dipindah dari `TUN-I` ke `TAB-I` membawa tujuan "dana darurat" *Then* `TUN-I` 3800000 (alokasi 2500000, belum dialokasikan 1300000); `TAB-I` 5500000 (alokasi 5500000); uang yang ada dan kekayaan bersih tidak berubah (CTH-10).
- **AC-2** — *Given* FR-FIN-108/AC-2 dan AC-3 *Then* alokasi listrik dan air berkurang sesuai pembayaran yang membawa tujuannya, tidak pernah di bawah 0.
- **AC-3** — *Given* alokasi "belanja bulanan" 2500000 *When* istri menguranginya menjadi 2000000 *Then* belum dialokasikan bertambah 500000 dan saldo tidak berubah.

### FR-FIN-122 — Alokasi yang melebihi uang
| | |
|---|---|
| Fitur | F-FIN-014 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 (yang disisihkan tidak boleh tampil sebagai uang yang ada) |
| Aktor dan izin | Keduanya melihat keterangan yang sama atas tempat uang dana bersama |
| Aturan acuan | 07 §3.9 ("alokasi melebihi uang Rp…"; bukan invarian bahwa belum dialokasikan ≥ 0) |
| Business rule | BR-FIN-058, BR-FIN-072 |
| Kedudukan | 07 §3.9 |

**Kriteria penerimaan**
- **AC-1** — *Given* `TUN-I` 3800000 dengan alokasi 2500000 (CTH-10) *When* belanja 1500000 tanpa tujuan dicatat **[Asumsi]** *Then* entri tersimpan; `TUN-I` 2300000; alokasi tetap 2500000; tempat uang menampilkan "alokasi melebihi uang Rp200.000" sampai salah satu mengurangi alokasinya.
- **AC-2** — *Given* AC-1 *Then* aplikasi tidak memilih alokasi mana yang dikurangi dan tidak menguranginya sendiri.
- **AC-3** — *Given* AC-1 *When* alokasi dikurangi menjadi 2300000 *Then* keterangan hilang.

## 6. F-FIN-015 — Tabungan sederhana

Masalah: `MSL-1`; `KS-3` (tabungan di ringkasan). Beban berulang: sekali per tujuan; per kali menabung (bulanan **[Asumsi]**). Dana darurat adalah satu tujuan tabungan, bukan fitur tersendiri.

### FR-FIN-125 — Membuat tujuan tabungan
| | |
|---|---|
| Fitur | F-FIN-015 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, KS-3 |
| Aktor dan izin | Anggota dewasa mana pun; tujuan `adults` (06 §9.2). Simetris (PG-01) |
| Aturan acuan | 07 §2.6 `purpose` (jenis `savings`, target dan tanggal target opsional), §3.9, §6.5 |
| Business rule | BR-FIN-074, BR-FIN-079 |
| Kedudukan | 07 §3.9. Q-123 diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-1 (b), (c) |

Masukan: nama (wajib), target (opsional, bilangan bulat > 0), tanggal target (opsional, tidak sebelum hari ini). Membuat tujuan tidak mengubah angka apa pun.

**Kriteria penerimaan**
- **AC-1** — *Given* rumah tangga mana pun *When* suami membuat "dana darurat" bertarget 10000000 *Then* tujuan tampil di daftar tabungan dengan kemajuan 0; uang yang ada dan kekayaan bersih tidak berubah; target tidak masuk angka mana pun.
- **AC-2** — *Given* target 0, negatif, atau berpecahan *Then* ditolak di batas. *Given* target kosong *Then* tujuan tersimpan tanpa target.

### FR-FIN-126 — Menabung
| | |
|---|---|
| Fitur | F-FIN-015 (memakai F-FIN-014) |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Sama dengan FR-FIN-120 |
| Aturan acuan | 07 §3.9 (menabung = alokasi atau pemindahan yang membawa tujuan; tidak pernah pengeluaran); CTH-09, CTH-10 |
| Business rule | BR-FIN-071, BR-FIN-074 |
| Kedudukan | 07 §3.9; brief §5H |

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-10 *When* 500000 dipindah ke `TAB-I` membawa "dana darurat" *Then* kemajuan "dana darurat" 5500000 dari 10000000; pengeluaran periode itu tidak bertambah; anggaran tidak berubah.
- **AC-2** — *Given* `TUN-I` dengan belum dialokasikan 1300000 *When* istri mengalokasikan 300000 di `TUN-I` untuk "dana darurat" tanpa memindah uang **[Asumsi]** *Then* kemajuan menjadi 5800000, dihitung dari alokasi di kedua tempat uang.

### FR-FIN-127 — Kemajuan tabungan dan baris di halaman depan
| | |
|---|---|
| Fitur | F-FIN-015 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, KS-3 |
| Aktor dan izin | Keduanya melihat angka yang sama (07 INV-20) |
| Aturan acuan | 07 §3.9 (kemajuan dari alokasi), §5.1 (angka utama = uang yang ada); 04 §4.7 |
| Business rule | BR-FIN-074, BR-FIN-095 |
| Kedudukan | Tabungan di luar angka utama dan satu baris di halaman depan: usulan dokumen 04 §4.7 atas Q-123 butir (a), diterima lewat "lanjut" di checkpoint S3 (12 November 2026), S3-1 (b). Pertanyaan pemilik produk yang memicunya: "dana darurat, impian/goals. udah?" (dokumen 17 Q-123) |

Angka utama tetap uang yang ada. Halaman depan memuat paling banyak satu baris "disisihkan untuk tabungan Rp…", yaitu jumlah alokasi semua tujuan tabungan; satu ketukan membuka daftar tujuan dengan target dan kemajuannya. Target tidak dijumlahkan ke angka mana pun, dan tidak ada angka "kurang Rp…" yang terasa seperti utang **[Usulan]** dokumen ini.

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-10 *Then* angka utama tetap uang yang ada (tidak berubah oleh pemindahan) dan halaman depan menampilkan "disisihkan untuk tabungan Rp5.500.000"; daftar tujuan menampilkan "dana darurat" 5500000 dari target 10000000.
- **AC-2** — *Given* tidak ada tujuan tabungan beralokasi *Then* baris itu tidak tampil.
- **AC-3** — *Given* suami dan istri membuka daftar tujuan *Then* angka keduanya identik.

### FR-FIN-128 — Menyelesaikan atau membatalkan tujuan tabungan
| | |
|---|---|
| Fitur | F-FIN-015 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Anggota dewasa mana pun atas tujuan tabungan dana bersama, karena alokasi tidak mengubah saldo dan tidak memerlukan persetujuan (07 §3.9); mengarsipkan dan mengaktifkan kembali tujuan dana bersama: anggota dewasa mana pun, dengan syarat alokasi nol, dan pengarsipannya masuk ringkasan anggota dewasa lainnya (N-14 `ringkasan`) (06 §8.1 baris "Mengarsipkan tujuan tabungan", revisi kesepuluh, **[Usulan]** `architect`; menutup G-08). Terlihat keduanya. Simetris (PG-01) |
| Aturan acuan | 07 §3.9 "Menyelesaikan atau membatalkan tujuan", §2.6 tabel pengarsipan baris "Tujuan", §6.3 tabel D, INV-12, INV-17, INV-21, T-38; CTH-45 |
| Business rule | BR-FIN-074, BR-FIN-075 |
| Kedudukan | **[Usulan]** `finance-domain`, putaran S3 (07 §3.9, T-38; menutup G-01), mengadopsi bawaan berkas ini dan menambah perlakuan tabel D. Dokumen 04 F-FIN-015 menyebut "melepas alokasinya"; kini punya aturannya |

Menyelesaikan atau membatalkan adalah **satu langkah** yang menampilkan alokasi tujuan itu per tempat uang dan menawarkan melepasnya di langkah yang sama, **tanpa pilihan terisi**. Melepas adalah satu entri alokasi berbaris negatif; tidak ada saldo yang berubah. Tujuan terarsip hanya bila alokasinya di semua tempat uang 0; tidak ada yang dilepas otomatis, dan tanpa memilih melepas, tujuan tidak terarsip. Alasan arsip (tercapai atau batal) hanya keterangan dan tidak mengubah angka apa pun. Tabel D periode pelepasan menampilkan tujuan itu dengan keterangan "diarsipkan"; periode sesudahnya tidak lagi (07 §3.9, §6.3).

**Kriteria penerimaan** (CTH-45)
- **AC-1** — *Given* "dana darurat" beralokasi 300000 di `TUN-I` (saldo 3800000, alokasi "belanja bulanan" 2500000, belum dialokasikan 1000000) dan 5500000 di `TAB-I` (saldo 5500000, belum dialokasikan 0); kemajuan 5800000 *When* istri membatalkannya dan memilih melepas semua *Then* tercatat satu entri alokasi dengan dua baris (−300000, −5500000); alokasi "dana darurat" 0 di kedua tempat uang; belum dialokasikan `TUN-I` 1300000 dan `TAB-I` 5500000; kemajuan 0; tujuan terarsip dengan alasan batal; uang yang ada dan kekayaan bersih tidak berubah.
- **AC-2** — *Given* AC-1 tanpa memilih melepas *Then* tujuan tidak terarsip dan semua angka tetap. *When* perintah arsip dikirim atas tujuan yang alokasinya belum 0 *Then* ditolak (07 §2.6).
- **AC-3** — *Given* AC-1 dan alokasi 5800000 sudah ada pada awal periode **[Asumsi]** CTH-45 *Then* tabel D periode itu menampilkan "dana darurat": awal 5800000, perubahan −5800000, akhir 0, keterangan "diarsipkan"; tabel D periode berikutnya tidak memuatnya.
- **AC-4** — *Given* tujuan yang sama diselesaikan dengan alasan tercapai *Then* semua angka sama dengan AC-1 sampai AC-3; hanya keterangan alasannya yang berbeda.

## 7. F-FIN-016 — Aset dan kekayaan bersih

Masalah: `MSL-1` ("aset apa yang kita punya"); `KB-09`. Beban berulang: sekali per aset; penilaian ulang bila mau, tanpa pengingat **[Usulan]** dokumen 04.

### FR-FIN-130 — Membeli aset
| | |
|---|---|
| Fitur | F-FIN-016 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Anggota dewasa mana pun; aset `adults`; di atas ambang mengikuti alur permintaan atau ditandai. Simetris (PG-01) |
| Aturan acuan | 07 §3.16, §3.8 baris `asset_purchase` tunai, §6.4; CTH-33 |
| Business rule | BR-FIN-076, BR-FIN-079 |
| Kedudukan | D-13 (diterima di checkpoint S1); 07 T-05 |

Masukan: nama aset, harga, tempat uang, tanggal kejadian, kategori opsional. Barang mana yang disebut aset adalah pilihan pencatat, tanpa batas nilai.

**Kriteria penerimaan**
- **AC-1** — *Given* `TUN-I` 2825000, uang yang ada 9544000, aset 12000000, kekayaan bersih 13644000 (CTH-33), dan permintaan laptop sudah disetujui *When* istri membeli laptop 2000000 dari `TUN-I` *Then* `TUN-I` 825000, uang yang ada 7544000, aset 14000000, kekayaan bersih tetap 13644000; arus kas menampilkan "beli aset"; anggaran tidak berubah kecuali entri diberi kategori.
- **AC-2** — *Given* suami yang membeli aset yang sama tanpa permintaan *Then* entri tersimpan dan ditandai dasar ambang (2000000 > 500000), tidak dihalangi.

### FR-FIN-131 — Memperbarui nilai aset
| | |
|---|---|
| Fitur | F-FIN-016 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Anggota dewasa mana pun atas aset dana bersama; anggota dewasa lainnya menerima N-14 sebagai `ringkasan` (06 §8.1, N-14; menutup G-06). Simetris (PG-01) |
| Aturan acuan | 07 §3.16 `asset_valuation`, §6.1, §6.3 tabel A; CTH-33; 06 §8.1, N-14 |
| Business rule | BR-FIN-077, BR-FIN-079 |
| Kedudukan | 07 §3.16; D-13. Pelaku dan N-14: **[Usulan]** `architect`, revisi kesembilan (06 §8.1). Tanpa pengingat: **[Usulan]** dokumen 04 |

Masukan: aset, nilai baru (bilangan bulat ≥ 0), tanggal penilaian. Tidak ada pengingat, saran nilai, atau penyusutan otomatis.

**Kriteria penerimaan**
- **AC-1** — *Given* motor bernilai 12000000 dan kekayaan bersih 13644000 *When* suami menilainya 11000000 *Then* aset 13000000, kekayaan bersih 12644000; ringkasan menampilkan perubahan nilai aset −1000000; anggaran dan arus kas tidak berubah; istri menerima N-14 sebagai `ringkasan`, bukan pemberitahuan langsung.
- **AC-2** — *Given* nilai negatif atau berpecahan *Then* ditolak di batas.

### FR-FIN-132 — Menjual atau melepas aset
| | |
|---|---|
| Fitur | F-FIN-016 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Anggota dewasa mana pun. Simetris (PG-01) |
| Aturan acuan | 07 §3.16 `asset_disposal`, §2.6 pengarsipan aset |
| Business rule | BR-FIN-078 |
| Kedudukan | 07 §3.16 |

**Kriteria penerimaan**
- **AC-1** — *Given* motor bernilai tercatat 11000000 **[Asumsi]** *When* dijual 10500000 ke `BNK-S` *Then* aset −11000000, `BNK-S` +10500000, perubahan nilai −500000 di tabel A; aset bernilai 0 dan dapat diarsipkan.
- **AC-2** — *Given* aset rusak dan dibuang tanpa hasil *Then* aset −nilai tercatat dan perubahan nilai sebesar itu; tidak ada uang masuk.

### FR-FIN-133 — Menampilkan kekayaan bersih
| | |
|---|---|
| Fitur | F-FIN-016 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Keduanya melihat angka yang identik; dana pribadi tidak masuk (07 §5.2) |
| Aturan acuan | 07 §6.5, §5.2, INV-20 |
| Business rule | BR-FIN-079, BR-FIN-095 |
| Kedudukan | 07 §6.5 |

Kekayaan bersih tampil bersama empat komponennya (uang yang ada, piutang, aset, kewajiban) dan satu ketukan ke penjelasan perubahannya per periode (FR-FIN-150).

**Kriteria penerimaan**
- **AC-1** — *Given* akhir periode CTH-37 *Then* kekayaan bersih 12084000 tampil sama di kedua ponsel, bersama komponennya.
- **AC-2** — *Given* tujuan "dana darurat" bertarget 10000000, anggaran, dan permintaan yang disetujui *Then* tidak satu pun dari ketiganya mengubah kekayaan bersih.

## 8. F-FIN-017 — Hitung uang dan dompet campuran

Masalah: `MSL-2`, `MSL-5` (lupa mencatat tertangkap di sini). Beban berulang: tidak dijadwalkan; perkiraan sekali seminggu sampai sebulan **[Asumsi]**. Satu hitungan menggantikan banyak catatan yang terlupa, jadi beban rendah ini sejalan dengan pencatat yang pelupa.

### FR-FIN-135 — Menghitung isi satu tempat uang
| | |
|---|---|
| Fitur | F-FIN-017 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-2, MSL-5 |
| Aktor dan izin | Pemegang tempat uang. Anggota dewasa lain hanya atas tempat uang dana bersama yang memenuhi **ketiga syarat** 07 §3.10 butir 1 (SP-50): (a) berjenis `bank` atau `ewallet` dan **penghitung termasuk `others_may_count`** tempat uang itu (himpunan anggota pilihan pemegang, 07 §2.3; penyisiran S4, M4; bawaan kosong; saat dinyalakan pilihan awalnya semua anggota dewasa lain; hanya pemegang yang dapat mengubahnya, termasuk mencabut per anggota, sebagai *peristiwa* tanpa entri yang terlihat keduanya; di rumah tangga pilot terbaca sebagai tanda "pasanganku dapat melihat saldo rekening ini"; FR-FIN-001 untuk pembuatan tempat uang); (b) tidak ada hitungan bukan pemegang lain yang masih terbuka atas tempat uang itu (INV-33); (c) pemegangnya anggota aktif: tidak dalam masa tenggang keluar dan belum keluar (§8.2 butir 2b). Atas tempat uang `joint(S)`, setiap anggota `S` adalah pemegang dan menghitung tanpa penanda; anggota di luar `S` tidak menghitungnya (07 INV-33, INV-36, T-46; M6). Uang tunai, termasuk dompet campuran, tidak pernah dihitung anggota lain. Hitungan *terbuka* berakhir saat pemegang mengonfirmasi, berkeberatan, atau mencatat hitungannya sendiri atas tempat uang itu. Bila penghitung bukan pemegang dan selisihnya bukan nol, baris selisih bertanda "belum dikonfirmasi" di sisi pemegang, dan pemegang menerima N-13 (`perlu_tanggapan`) bila positif atau N-14 **`ringkasan`** bila negatif (06 N-14 revisi kesepuluh); ia mengonfirmasi, berkeberatan, atau merinci seperti biasa (06 §7.2 baris G-05). Layar **tidak menawarkan** menghitung tempat uang yang tidak memenuhi syarat itu (06 §7.2). Rekonsiliasi tempat uang dana bersama terlihat keduanya, termasuk "dihitung, cocok" beserta pelakunya bila selisih 0. Simetris (PG-01). **Ditandai menurut tugas agen ini:** bukan pemegang dapat mengubah saldo tempat uang yang dipegang pasangannya; pagarnya penanda, notifikasi, dan tidak adanya tombol hitung atas uang tunai pasangan |
| Aturan acuan | 07 §3.10 butir 1 (syarat (a) sampai (c)), 2–4, 7–8, "Hitungan yang tiba belakangan", "Menghitung lalu merinci seluruh selisih sekaligus", §2.5 baris (a) dan (c), §2.6 `reconciliation` dan atribut `others_may_count`, §8.2 butir 2b, §10.3 butir 3, INV-16, INV-33, T-33; CTH-29, CTH-46, CTH-47, CTH-48; 06 §7.2, §8.1, §12 S-10, N-13, N-14 |
| Business rule | BR-FIN-080, BR-FIN-082, BR-FIN-084 |
| Kedudukan | Keputusan pemilik produk nomor 7 (checkpoint S1, diputuskan sendiri). Selisih dihitung server untuk hitungan yang tiba belakangan: **[Usulan]** `finance-domain` (07 §3.10, putaran S3) dan `architect` (06 S-10); menutup G-02. Hitungan oleh bukan pemegang: **[Usulan]** `finance-domain` (07 §3.10 butir 1, T-33) dan `architect` (06 §7.2); menutup G-05. Batasnya (tanda pemegang, satu hitungan terbuka, pemegang aktif, N-14 `ringkasan`): usulan `security-privacy` SP-50, diterapkan 07 dan 06 putaran perbaikan S3 (07 §10.4 butir 2), diterima lewat "lanjut" di checkpoint S3, S3-9 (i); harganya W-26. Merinci sekaligus: **[Usulan]** `finance-domain` (07 §3.10; dokumen 11 IDEA-01) |

Masukan: tempat uang, jumlah terhitung (bilangan bulat ≥ 0), tanggal menghitung (bawaan hari ini). Layar menampilkan saldo tercatat dan perkiraan selisihnya sebelum disimpan. **Selisih yang tercatat dihitung server** terhadap saldo tercatat saat perintah diterima, dan server mengisi saldo tercatat saat hitung; perkiraan di perangkat hanya tampilan berlabel (07 §3.10 "Hitungan yang tiba belakangan"). Selisih bukan nol menjadi satu entri selisih kas bertanggal hari menghitung; selisih 0 tercatat sebagai rekonsiliasi tanpa entri. Bila selisih server berbeda dari perkiraan yang tampil, hasil sinkron di perangkat penghitung menampilkan selisih server di tempat perkiraan itu dengan satu kalimat bahwa catatan lain sudah lebih dulu sampai; itu bukan notifikasi (06 S-10). Tawaran "ini bagian dari selisih hitung?" yang baru muncul sesudah sinkron dijalankan sebagai dua perintah yang sudah ada, masing-masing dengan kuncinya (06 S-10). Pemegang yang menghitung boleh merinci seluruh selisih negatif di langkah yang sama; bila jumlah rinciannya sama dengan selisih server, hasilnya satu entri pengeluaran yang membawa rekonsiliasinya, tanpa entri selisih kas (07 §3.10 "Menghitung lalu merinci seluruh selisih sekaligus"). Tidak ada kalimat yang menyebut siapa yang lupa.

**Kriteria penerimaan**
- **AC-1** — *Given* `TUN-I` tercatat 2835000, uang yang ada 9684000, kekayaan bersih 13784000 (CTH-29) *When* istri menghitung 2750000 pada 10 November *Then* tercatat selisih −85000 bertanggal 10 November; `TUN-I` 2750000; uang yang ada 9599000; kekayaan bersih 13699000; ringkasan menghitungnya sebagai uang keluar "belum dirinci".
- **AC-2** — *Given* `TUN-S` tercatat 650000 *When* suami menghitung 660000 pada 13 November *Then* selisih +10000 tampil sebagai baris tersendiri di sisi uang masuk, bukan pemasukan; uang yang ada 9609000; kekayaan bersih 13709000.
- **AC-3** — *Given* selisih −85000 belum dirinci *Then* anggaran "belanja bulanan" tidak berkurang karenanya, dan layar anggaran menampilkan satu keterangan "ada selisih yang belum dirinci Rp85000" yang tidak menempel pada kategori mana pun (07 §3.10 butir 8; FIN-1.md FR-FIN-090/AC-5).
- **AC-4** — *Given* hitungan sama dengan saldo tercatat *Then* tidak ada entri; pasangan melihat "dihitung, cocok" dengan tanggal dan pelakunya.
- **AC-5** — *Given* jumlah terhitung negatif atau berpecahan *Then* ditolak di batas.
- **AC-6** — *Given* hitungan terkirim dua kali dengan kunci yang sama *Then* hanya ada satu rekonsiliasi dan satu entri selisih.
- **AC-7** — *Given* CTH-46: `TUN-S` 900000, `TUN-I` 500000, uang yang ada 1400000 di server dan ponsel istri; suami menyerahkan 200000 tunai pada 10 November dan serah terimanya sampai di server (`TUN-S` 700000, `TUN-I` 700000) *When* hitungan istri 700000, dibuat tanpa sinyal di ponsel yang masih menampilkan tercatat 500000 dan perkiraan selisih +200000, sampai di server sesudahnya *Then* selisih server 0 (700000 − 700000), tidak ada entri, saldo tercatat saat hitung 700000, rekonsiliasinya "dihitung, cocok", `TUN-I` 700000 dan uang yang ada 1400000 (bukan 900000 dan 1600000); ponsel istri mengganti perkiraan +200000 dengan selisih server 0 dan satu kalimat, tanpa notifikasi; serah terima tetap belum dikonfirmasi sampai istri mengonfirmasi, karena menghitung bukan konfirmasi.
- **AC-8** — *Given* CTH-47: `BNK-I` dana bersama dipegang istri, tercatat 3800000, dan istri sudah menyalakan `others_may_count` sehingga himpunannya memuat suami *When* pada 10 November suami merekonsiliasinya dengan 3750000 *Then* tercatat selisih 3750000 − 3800000 = −50000 yang bertanda "belum dikonfirmasi" di sisi istri, `BNK-I`, uang yang ada, dan kekayaan bersih 3750000 seketika, dan istri menerima N-14 sebagai `ringkasan`, tidak pernah pemberitahuan langsung; istri dapat mengonfirmasi, merinci (misalnya biaya administrasi bank 50000, sebagai `recategorization` tanpa mengubah saldo), atau berkeberatan, dan keberatan tidak mengubah angka sampai ada entri pembalik. *When* suami menghitung 3800000 *Then* "dihitung, cocok" beserta pelakunya, tanpa entri, penanda, atau notifikasi. *When* istri sendiri mencatat selisih −50000 *Then* angkanya sama, tanpa penanda. *Given* selisih positif yang dicatat bukan pemegang *Then* pemegang menerima N-13 (`perlu_tanggapan`).
- **AC-9** — *Given* `TUN-S` dipegang suami, atau dompet campuran suami *When* istri membukanya *Then* tidak ada tombol menghitung, dan perintah hitung istri yang dikirim langsung ke server ditolak. *Given* `BNK-I` dipegang istri dengan `others_may_count` memuat suami *When* suami membukanya *Then* tombol menghitung ada; *Given* himpunannya kosong (bawaan) *Then* tidak ada tombol, dan hitungan suami atas `BNK-I` 3800000 ditolak, semua angka tetap 3800000 (07 §3.10 butir 1 (a); CTH-47 catatan) (06 §7.2).
- **AC-10** — *Given* CTH-48: `TUN-I` tercatat 450000, ambang 500000 *When* pada 14 November istri menghitung 285000 dan merinci seluruh selisih di langkah yang sama, "belanja bulanan" 145000 dan "transportasi" 20000 *Then* tercatat satu entri pengeluaran 165000 yang membawa rekonsiliasinya, tanpa entri selisih kas; `TUN-I` dan uang yang ada 285000; belum dirinci 0; tabel A belanja bulanan −145000 dan transportasi −20000; arus kas belanja −165000 dan selisih hitung 0; anggaran belanja bulanan terpakai +145000; entri tidak ditandai (165000 ≤ 500000). *Given* istri hanya mengetahui sebagian (jumlah rincian lebih kecil dari 165000) *Then* bentuk ini tidak ditawarkan; ia memakai jalur biasa (selisih kas penuh, lalu merinci yang ia ketahui). *Given* entri lain atas `TUN-I` sampai di server lebih dulu sehingga selisih server bukan 165000 *Then* perintah disimpan sebagai rekonsiliasi biasa dengan selisih kas sebesar selisih server, dan perinciannya ditawarkan lagi sesudah sinkron (07 §3.10). Bentuk ini hanya ditawarkan kepada pemegang, dan tidak pada layar dompet campuran **[Usulan]** dokumen ini, mengikuti dokumen 11.
- **AC-11** — *Given* AC-8 (hitungan suami 10 November masih terbuka, `BNK-I` 3750000) *When* pada 11 November suami menghitung lagi, terhitung 3740000 *Then* ditolak (INV-33), juga bila hasilnya cocok dengan saldo tercatat; `BNK-I` tetap 3750000, dan layar menyatakan bahwa hitungan sebelumnya masih menunggu jawaban istri. *When* pada 12 November istri mengonfirmasi baris −50000 (peristiwa, tanpa entri) dan sesudahnya suami menghitung 3740000 *Then* tercatat selisih 3740000 − 3750000 = −10000 yang bertanda belum dikonfirmasi seperti baris −50000, `BNK-I` 3740000, dan sejak itu hitungan 12 November yang terbuka. Jumlah selisih yang dicatat suami −50000 + (−10000) = −60000 = 3740000 − 3800000 (CTH-47). Keberatan istri atau hitungannya sendiri atas `BNK-I` juga menutup hitungan 10 November.
- **AC-12** — *Given* `BNK-I` dengan hitungan suami 10 November yang masih terbuka *When* istri mematikan `others_may_count` atau mencabut suami darinya *Then* hitungan 10 November tetap terbuka sampai istri menjawab, dan hitungan suami berikutnya ditolak karena suami tidak lagi termasuk himpunannya (CTH-47 catatan). *When* suami mencoba mengubah `others_may_count` atas `BNK-I` *Then* ditolak; hanya pemegang yang dapat mengubahnya, dan setiap perubahan tercatat sebagai peristiwa yang terlihat keduanya, tanpa entri (07 §2.6; 06 §8.1).
- **AC-13** — *Given* istri, pemegang `BNK-I` dengan tanda menyala, keluar (F06) dan berada dalam masa tenggang *When* suami menghitung `BNK-I` *Then* ditolak (07 §3.10 butir 1 (c), §8.2 butir 2b); entri lain atas `BNK-I` tetap boleh (FAM.md FR-FAM-012/AC-4).
- **AC-14** — **IDEA-01 hanya atas uang tunai berpemilik tunggal** (batas itu bagian dari S3-2 (i), diterima lewat "lanjut" di checkpoint S3, Tier 1; 07 §3.10 "Menghitung lalu merinci seluruh selisih sekaligus", baris "Syarat" dan "Dompet campuran"). Bahwa "berpemilik tunggal" berarti tidak termasuk dompet campuran adalah bacaan 07 yang penegasannya diserahkan kepada dokumen 05; dokumen ini menegaskannya. AC-10 sudah menguji bentuknya, syarat "seluruh selisih dirinci", dan penolakan perincian sebagian. *Given* CTH-47: `BNK-I` dana bersama dipegang istri, tercatat 3800000 *When* istri sendiri menghitung 3750000 *Then* bentuk "hitung lalu rinci seluruh selisih" **tidak ditawarkan** karena `BNK-I` bukan uang tunai; tercatat selisih −50000 lewat jalur biasa, tanpa penanda, dan istri dapat merincinya sesudahnya (`recategorization`). *Given* dompet campuran *Then* bentuk ini juga tidak ditawarkan (AC-10; 07 §3.10 baris "Dompet campuran"). Skenario: dokumen 12 SCN-28.
### FR-FIN-136 — Merinci selisih dan entri terlambat yang menutupnya
| | |
|---|---|
| Fitur | F-FIN-017 (dengan F-FIN-005) |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-5, MSL-1 |
| Aktor dan izin | Pencatat entri terlambat atau anggota dewasa mana pun atas selisih tempat uang dana bersama. Simetris (PG-01) |
| Aturan acuan | 07 §3.10 butir 5–8, §3.11; 06 §6.2 (baris `expense` dalam `recategorization` dinilai terhadap ambang); CTH-29; 04 KM-11 |
| Business rule | BR-FIN-081, BR-FIN-082, BR-FIN-099 |
| Kedudukan | Bawaan ya: D-05 (07 T-04) |

Saat pengeluaran dicatat pada tempat uang dengan tanggal kejadian tidak lebih baru daripada rekonsiliasi terakhirnya yang masih punya selisih negatif belum dirinci, layar bertanya "ini bagian dari selisih hitung Rp…?" dengan bawaan ya. Merinci juga dapat dilakukan langsung dari entri selisih.

**Kriteria penerimaan**
- **AC-1** — *Given* selisih `TUN-I` −85000 belum dirinci (CTH-29) *When* pada 12 November istri mencatat sayur 60000 bertanggal 8 November dan menjawab ya *Then* saldo `TUN-I` tidak berubah; yang belum dirinci 25000; "belanja bulanan" terpakai 540000, sisa 1960000, pada periode tanggal 8 November.
- **AC-2** — *Given* sisa belum dirinci 25000 *When* dicatat pengeluaran terlambat 40000 **[Asumsi]** bertanggal 9 November dengan jawaban ya *Then* 25000 menjadi perincian dan 15000 menjadi pengeluaran biasa yang mengurangi `TUN-I`; yang belum dirinci 0, tidak negatif.
- **AC-3** — *Given* pengeluaran terlambat bertanggal sesudah hitungan terakhir *Then* pertanyaan itu tidak muncul.
- **AC-4** — *Given* ambang 500000 *When* perincian 600000 **[Asumsi]** disimpan *Then* baris pengeluarannya dinilai terhadap ambang seperti pengeluaran yang dicatat tepat waktu.
- **AC-5** — *Given* suami yang mencatat perincian atas selisih `TUN-I` *Then* hasilnya sama; tidak ada keterangan yang menyebut siapa yang lupa.
- **AC-6** — *Given* tempat uang yang diarsipkan sesudah hitungannya dan masih punya selisih negatif yang belum dirinci (CTH-43 bagian A) *When* perincian dicatat *Then* diterima tanpa mengaktifkan kembali tempat uang itu, karena perincian tidak punya baris posisi (07 §2.6, §3.10 butir 6; FIN-1.md FR-FIN-004/AC-6).
- **AC-7** — *Given* pengeluaran terlambat yang tersimpan di server sebelum pencatatnya tahu tentang rekonsiliasi yang masih punya selisih negatif *When* sesudah sinkron ia menerima tawaran "ini bagian dari selisih hitung?" *Then* dijalankan sebagai dua perintah dengan kunci sendiri: entri pembalik atas pengeluaran itu, lalu perincian sebesar jumlah yang sama dengan tanggal dan kategori yang sama; tidak ada perintah gabungan (07 §3.10 "Hitungan yang tiba belakangan"; 06 S-10).

### FR-FIN-137 — Menghitung dompet campuran
| | |
|---|---|
| Fitur | F-FIN-017 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-2, MSL-5 |
| Aktor dan izin | Hanya pemegang dompet itu, karena ia pemilik bagian pribadinya (07 INV-25; peran tidak dapat dipertukarkan). Pasangannya hanya melihat rekonsiliasi tempat uang dana bersama, persis seperti rekonsiliasi lain, dan tetap dapat merekonsiliasi tempat uang dana bersama itu sendirian. Isi seluruhnya, jumlah kedua saldo, bagian pribadi selisihnya, pilihan pembagiannya, dan **bahwa** hitungan itu bagian dompet campuran tidak terbaca pasangan (06 V-9, INV-A41) |
| Aturan acuan | 07 §2.3, §3.10 "Dompet campuran", T-03, §10.3 butir 5; 06 V-9, INV-A41, §9.4 "Batas semua larangan"; CTH-40 |
| Business rule | BR-FIN-080, BR-FIN-083 |
| Kedudukan | Pembagian tanpa bawaan, dipilih setiap kali: diterima pemilik produk atas saran lead, 12 November 2026. V-9: **[Usulan]** (Q-115). Dompet campuran adalah kasus biasa: pernyataan pemilik produk, checkpoint S2 (S2-6: "bisa terpisah bisa tercampur bisa di bank bisa tunai") |

Syarat layar (07 §2.3; dokumen 17 "Dibawa ke S3"): **satu layar, satu angka hitungan**, dan pilihan pembagian "bersama", "pribadi", atau "dibagi" dengan **tidak ada pilihan terisi**. Hitungan tidak dapat disimpan sebelum pilihan diambil bila selisih bukan nol. Untuk "dibagi", bagian pribadi bertanda sama dengan selisih dan nilai mutlaknya tidak melebihi selisih. Validasi ini terjadi di perangkat, karena server tidak menerima isi seluruhnya; server menerima dua rekonsiliasi biasa dengan kuncinya masing-masing. Layar menyatakan terus terang dua batas: orang yang melihat dompet fisiknya dapat menghitung bagian pribadi, dan larangan ini tidak melindungi dari pengelola server (07 §2.3; 06 V-9). Kata-katanya milik dokumen 08.

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-40: `DMP-B` 680000, `DMP-P` 250000, `BNK-S` 2000000 *When* pemegang menghitung isi dompet 890000 *Then* layar menampilkan selisih −40000 dan tiga pilihan, tidak satu pun terpilih; hitungan belum dapat disimpan sampai satu dipilih.
- **AC-2** — *Given* AC-1 *When* dipilih "bersama" *Then* `DMP-B` 640000, uang yang ada 2640000; pasangan melihat selisih −40000 pada `DMP-B`.
- **AC-3** — *Given* AC-1 *When* dipilih "pribadi" *Then* `DMP-P` 210000, uang yang ada 2680000; pasangan melihat `DMP-B` "dihitung, cocok" di 680000.
- **AC-4** — *Given* AC-1 *When* dipilih "dibagi" dengan bagian pribadi −15000 *Then* `DMP-B` 655000, `DMP-P` 235000, uang yang ada 2655000. *When* bagian pribadi diisi +10000 atau −50000 *Then* tidak dapat disimpan.
- **AC-5** — *Given* AC-2, AC-3, dan AC-4 *Then* dari seluruh yang terbaca pasangan (layar, riwayat, notifikasi, ringkasan, ekspor), 890000, 930000, saldo `DMP-P`, dan pilihan pembagian tidak dapat ditemukan atau diturunkan; yang terbaca identik dengan rekonsiliasi `DMP-B` saja dengan selisih bersama yang sama.
- **AC-6** — *Given* hitungan dengan selisih bukan nol dikirim tanpa pilihan pembagian (misalnya dari klien yang rusak) *Then* tidak ada entri selisih dan tidak ada saldo yang berubah.
- **AC-7** — *Given* AC-2 dan kemudian pemegang sadar 40000 itu pengeluaran pribadinya *Then* perbaikannya dua catatan: pembalikan selisih bersama (+40000, `DMP-B` 680000) dan pengeluaran pribadi 40000 (`DMP-P` 210000); entri asal tetap terlihat.
- **AC-8** — *Given* pasangan (bukan pemegang) *When* ia merekonsiliasi `DMP-B` sendirian dengan hitungan 680000 *Then* tercatat "dihitung, cocok" tanpa menyentuh `DMP-P`.

### FR-FIN-138 — Kapan menghitung ditawarkan
| | |
|---|---|
| Fitur | F-FIN-017 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-5, MSL-2 |
| Aktor dan izin | Ajakan hanya kepada pemegang tempat uang itu; tidak pernah kepada pasangannya tentang tempat uang yang dipegang orang lain |
| Aturan acuan | 07 §2.5 "Saldo negatif", T-06; 03 §7 butir 8; 04 §3.3 (tanpa pengingat "kamu belum mencatat"), PRN-08 |
| Business rule | BR-FIN-084, BR-FIN-097 |
| Kedudukan | Ajakan saat saldo negatif: D-05. Tanpa jadwal dan tanpa pengingat "belum mencatat": usulan dokumen 04 dan 03 §7 butir 8 (Q-90, Q-91), diterima lewat "lanjut" di checkpoint S3 sebagai D-75 |

Menghitung selalu tersedia dari tempat uang. Aplikasi menawarkannya hanya saat saldo tercatat suatu tempat uang negatif, sebagai keterangan pada tempat uang itu di layar pemegangnya, bukan pemberitahuan **[Usulan]** dokumen ini (menutup 03 §7 butir 8). Tidak ada jadwal, hitungan "terakhir dihitung … hari lalu" yang dibandingkan antarorang, atau pengingat mencatat.

**Kriteria penerimaan**
- **AC-1** — *Given* `TUN-I` tercatat −20000 **[Asumsi]** *Then* istri, pemegangnya, melihat ajakan menghitung pada `TUN-I`; entri yang membuatnya negatif tetap tersimpan.
- **AC-2** — *Given* AC-1 *Then* suami tidak menerima ajakan atau pemberitahuan tentang `TUN-I`; ia melihat saldo negatifnya seperti saldo lain.
- **AC-3** — *Given* tidak ada saldo negatif *Then* tidak ada ajakan, jadwal, atau pengingat menghitung.

## 9. F-FIN-018 — Koreksi dan keberatan

Masalah: `MSL-1` (angka kembali benar), `MSL-5` (salah catat tidak menakutkan). Beban berulang: jarang **[Asumsi]**. Menjawab dokumen 17 "Keluhan" 4 dan 10.

### FR-FIN-140 — Membalik entri
| | |
|---|---|
| Fitur | F-FIN-018 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, MSL-5 |
| Aktor dan izin | Menurut FR-FIN-142. Asal, koreksi, pelaku, tanggal, dan alasan terlihat sama oleh keduanya (06 §8.1). Simetris (PG-01) untuk entri buku rumah tangga |
| Aturan acuan | 07 §3.12, INV-09, INV-18; 06 §8.1, N-18; CTH-30 |
| Business rule | BR-FIN-085, BR-FIN-086, BR-FIN-090 |
| Kedudukan | Tidak ada penyuntingan: keputusan pemilik produk nomor 7 (checkpoint S1). Kedua pasangan boleh mengoreksi entri bersama: Q-26 (bawaan diterima lewat "lanjut") |

Alasan wajib: pilihan singkat ("salah jumlah", "salah tempat uang", "tercatat dua kali", "tidak terjadi") atau teks bebas; tanpa alasan ditolak. Pembalikan bertanggal sama dengan entri asal. Tidak ada tombol hapus; layar menerangkan dalam satu kalimat mengapa, dan bahwa angkanya kembali benar (dokumen 17 "Keluhan" 4). Koreksi tidak memerlukan persetujuan. Bila pengoreksi bukan pencatat asal, pencatat asal menerima N-18. Pembalikan yang menambah tempat uang yang dipegang anggota dewasa lain membawa penanda belum dikonfirmasi (07 §3.12).

**Kriteria penerimaan**
- **AC-1** — *Given* parkir tercatat 50000 dari `TUN-I`, sehingga `TUN-I` 2700000 dan uang yang ada 9559000 (CTH-30) *When* istri membaliknya dengan alasan "salah jumlah" lalu mencatat parkir 5000 *Then* `TUN-I` 2745000, uang yang ada 9604000, kekayaan bersih 13704000; "transportasi" terpakai 5000; entri 50000 tetap terlihat bersama pembalikannya.
- **AC-2** — *Given* AC-1 dengan suami sebagai pengoreksi atas entri yang dicatat istri *Then* angkanya sama dan istri menerima N-18. Sebaliknya juga berlaku.
- **AC-3** — *Given* pembalikan tanpa alasan *Then* ditolak di batas.
- **AC-4** — *Given* entri yang sudah dibalik *When* ia dibalik lagi, atau pembalikannya dibalik *Then* ditolak.
- **AC-5** — *Given* bensin 40000 terkirim dua kali dengan kunci yang sama (CTH-32) *Then* hanya ada satu entri, sehingga tidak perlu koreksi; bila dua ketukan menghasilkan dua kunci, salah satunya dibalik dengan alasan "tercatat dua kali".

### FR-FIN-141 — Menyesuaikan entri dan memindah kategorinya
| | |
|---|---|
| Fitur | F-FIN-018 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, MSL-5 |
| Aktor dan izin | Sama dengan FR-FIN-140 |
| Aturan acuan | 07 §3.12 (adjustment, recategorization untuk perubahan bukan jumlah), §3.15 aturan 7b, INV-37; 06 §8.1 (perubahan bukan uang adalah Peristiwa), INV-A52; CTH-30 |
| Business rule | BR-FIN-051, BR-FIN-085, BR-FIN-086 |
| Kedudukan | 07 §3.12 |

Penyesuaian adalah satu entri sebesar bedanya; pindah kategori tidak mengubah saldo. Layar menawarkan jalur yang paling sedikit langkahnya dan tidak pernah mengubah entri asal.

**Kriteria penerimaan**
- **AC-1** — *Given* parkir 50000 dari `TUN-I` (`TUN-I` 2700000) *When* disesuaikan +45000 dengan alasan "salah jumlah" *Then* `TUN-I` 2745000, hasil akhirnya sama dengan FR-FIN-140/AC-1.
- **AC-2** — *Given* belanja 150000 berkategori "jajan" yang seharusnya "perlengkapan rumah" **[Asumsi]** *When* kategorinya dipindah *Then* saldo tidak berubah; anggaran kedua kategori berubah pada periode tanggal kejadian aslinya; periode lampau yang terdampak ditandai berubah.
- **AC-3** — *Given* catatan teks entri diubah *Then* tercatat sebagai Peristiwa, tanpa entri uang.
- **AC-4** — *Given* pembayaran yang tertaut ke pos "Listrik" berkategori "listrik" *When* siapa pun memindah kategorinya ke "belanja bulanan" *Then* `recategorization` ditolak; layar menyebut bahwa tautannya harus dilepas lebih dulu, dan sesudah dilepas entri dapat dipindah kategori dan dinilai terhadap ambang seperti entri tanpa tautan (07 §3.15 aturan 7b, INV-37; 06 INV-A52; FR-FIN-103/AC-14).

### FR-FIN-142 — Batas koreksi atas buku pribadi dan entri dua buku
| | |
|---|---|
| Fitur | F-FIN-018 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1; pagar privasi |
| Aktor dan izin | Entri buku pribadi M: hanya M. Entri dua buku (menyentuh buku rumah tangga dan buku pribadi M): dibalik hanya oleh M; pasangannya dapat menyesuaikan sisi rumah tangga atau mengajukan keberatan. Peran tidak dapat dipertukarkan: hak mengikuti pemilik buku |
| Aturan acuan | 07 §3.12, INV-25; 06 §8.1, V-1 |
| Business rule | BR-FIN-086 |
| Kedudukan | 07 §3.12; 06 §8.1 |

**Kriteria penerimaan**
- **AC-1** — *Given* suami memiliki entri di buku pribadinya *Then* istri tidak melihatnya dan tidak dapat mengoreksinya.
- **AC-2** — *Given* kontribusi istri 40000 dari dana pribadinya untuk belanja rumah tangga *When* suami mencoba membaliknya *Then* ditolak; ia dapat menyesuaikan sisi rumah tangga atau mengajukan keberatan. *When* istri membaliknya *Then* kedua sisi terbalik.
- **AC-3** — *Given* peran tertukar pada AC-2 *Then* hasilnya sama.

### FR-FIN-143 — Mengajukan keberatan
| | |
|---|---|
| Fitur | F-FIN-018 (dengan F-REQ-011 untuk diskusi) |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1; MSL-7 (tidak menjadi alat serang) |
| Aktor dan izin | Anggota dewasa yang dapat melihat sasarannya dan bukan pencatatnya (peran tidak dapat dipertukarkan dalam satu keberatan; setiap orang dapat menjadi pengaju). Keberatan terlihat keduanya |
| Aturan acuan | 06 §8.2 D01, INV-A25 sampai INV-A28, N-16; 07 §3.12 (klaim internal yang disengketakan tetap angkanya) |
| Business rule | BR-FIN-087, BR-FIN-088, BR-FIN-097 |
| Kedudukan | D-22, D-38. Daftar alasan tertutup: **[Usulan]** `architect`, 06 §8.2 |

Alasan dipilih dari daftar tertutup yang netral (06 §8.2): "jumlahnya tidak cocok", "uang ini tidak saya terima atau tidak saya serahkan", "tercatat dua kali", "tempat uang atau pemegangnya keliru", "perlu kita bicarakan"; tanpa teks bebas di alasan. Penjelasan ditulis di diskusi catatan itu. Tidak ada batas waktu. Angka tidak berubah. Pencatat sasaran menerima N-16. Kata layar untuk "keberatan" milik dokumen 08 (Q-119).

**Kriteria penerimaan**
- **AC-1** — *Given* suami mencatat pengeluaran "hadiah" 230000 dari dana bersama *When* istri mengajukan keberatan "jumlahnya tidak cocok" dan menulis penjelasan di diskusi *Then* keberatan tercatat; tidak ada saldo atau angka yang berubah; suami menerima N-16; keduanya melihat butir itu di daftar "belum sepakat".
- **AC-2** — *Given* peran tertukar *Then* hasilnya sama.
- **AC-3** — *Given* suami mencoba mengajukan keberatan atas entrinya sendiri *Then* ditolak.
- **AC-4** — *Given* istri sudah punya keberatan terbuka atas sasaran itu *When* ia mengajukan lagi *Then* ditolak dan ia diarahkan ke keberatan yang ada.
- **AC-5** — *Given* alasan kosong atau di luar daftar *Then* ditolak di batas.

### FR-FIN-144 — Menutup keberatan dan daftar "belum sepakat"
| | |
|---|---|
| Fitur | F-FIN-018 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, MSL-7 |
| Aktor dan izin | Menutup: hanya pengaju. Daftar: satu, sama bagi keduanya |
| Aturan acuan | 06 §8.2 D02, N-17, daftar "belum sepakat", "sasaran sudah dikoreksi", AT-4 |
| Business rule | BR-FIN-087, BR-FIN-088, BR-FIN-097 |
| Kedudukan | D-22, D-38; 06 §8.2 |

Daftar memuat keberatan terbuka yang sasarannya belum dibalik dan pengajunya masih anggota. Tidak ada hitungan per orang, urutan menurut orang, atau penyaring menurut pengaju. Aplikasi tidak memutuskan siapa yang benar dan tidak menutup apa pun sendiri.

**Kriteria penerimaan**
- **AC-1** — *Given* keberatan istri atas pengeluaran "hadiah" *When* suami mencoba menutupnya *Then* ditolak. *When* istri menutupnya dan menunjuk koreksi yang menyelesaikannya *Then* keberatan tertutup dan suami menerima N-17.
- **AC-2** — *Given* sasaran keberatan dibalik *Then* keberatan tampil "sasaran sudah dikoreksi" dan keluar dari daftar, walaupun state tersimpannya tetap terbuka.
- **AC-3** — *Given* keberatan dari keduanya *Then* daftar tidak menampilkan berapa yang diajukan masing-masing.
- **AC-4** — *Given* keberatan terbuka 60 hari *Then* tidak ada penutupan otomatis.

## 10. F-FIN-019 — Riwayat satu catatan

Masalah: `MSL-1`; arah 1 pemilik produk ("semua ada catatanya ada datanya", dokumen 04). Beban berulang: nol (dibaca). Fitur berstatus **Pantau** (04 §7).

### FR-FIN-145 — Riwayat satu catatan
| | |
|---|---|
| Fitur | F-FIN-019 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1 |
| Aktor dan izin | Siapa pun yang dapat melihat catatan itu melihat riwayat yang sama (AT-2) |
| Aturan acuan | 06 §9.4 AT-1, AT-2, AT-5, AT-6, PG-08, INV-A01, S-5; 07 §3.11, INV-09 |
| Business rule | BR-FIN-089, BR-FIN-090 |
| Kedudukan | Pelaku terlihat "bukan audit": keputusan pemilik produk nomor 4 (checkpoint S1, diputuskan sendiri) |

Pada tiap catatan (entri, tagihan, cicilan, piutang, kewajiban, aset, tujuan, alokasi, kesepakatan): siapa mencatat, siapa membayar, siapa menyetujui, siapa mengoreksi; tanggal kejadian, dan tanggal pencatatan hanya bila berbeda; koreksi dan keberatannya beserta alasannya. Urutan sama di semua perangkat.

**Kriteria penerimaan**
- **AC-1** — *Given* parkir CTH-30 dengan pembalikan dan koreksinya *Then* riwayat entri parkir menampilkan pencatat, pengoreksi, alasan "salah jumlah", dan tanggal; jam tidak tampil.
- **AC-2** — *Given* sayur 60000 bertanggal kejadian 8 November dan dicatat 12 November *Then* tampil kedua tanggal. *Given* entri yang dicatat pada hari kejadiannya *Then* hanya tanggal kejadian.
- **AC-3** — *Given* catatan apa pun *Then* tidak ada lokasi, tanda dibaca, atau "terakhir aktif".

### FR-FIN-146 — Tidak ada riwayat lintas catatan menurut orang
| | |
|---|---|
| Fitur | F-FIN-019 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-7 (dokumen 17 "Keluhan" 7) |
| Aktor dan izin | Berlaku sama bagi keduanya |
| Aturan acuan | 06 AT-3, AT-4, PG-08, PG-09 |
| Business rule | BR-FIN-089, BR-FIN-097 |
| Kedudukan | 06 §9.4; D-43 |

**Kriteria penerimaan**
- **AC-1** — *Given* layar mana pun *Then* tidak ada umpan kegiatan, penyaring "dicatat oleh", atau daftar "semua koreksi/keberatan/catatan terlambat seseorang".
- **AC-2** — *Given* ekspor *Then* tidak ada kolom atau tampilan tersimpan menurut tanda, keberatan, koreksi, atau catatan terlambat (D-43).

## 11. F-FIN-020 — Ringkasan bulanan dan arus kas

Masalah: `MSL-1`; `KS-3`; `KB-08`. Beban berulang: nol (dibaca), sekali sebulan **[Asumsi]**. Fitur berstatus **Pantau** (04 §7). Menjawab kriteria sukses brief §1.1 nomor 3 dan dokumen 17 "Keluhan" 7 sampai 9. Kata-kata dan keterangannya milik dokumen 08.

### FR-FIN-150 — Periode ringkasan dan penjelasan perubahan kekayaan bersih
| | |
|---|---|
| Fitur | F-FIN-020 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, KS-3 |
| Aktor dan izin | Keduanya melihat angka identik; hanya buku rumah tangga (07 INV-20; 06 AT-2) |
| Aturan acuan | 07 §6.3 tabel A, INV-19; CTH-37 |
| Business rule | BR-FIN-091, BR-FIN-093, BR-FIN-095, BR-FIN-096 |
| Kedudukan | 07 §6.3 |

Periode = periode anggaran rumah tangga. Tabel A menjelaskan mengapa kekayaan bersih berubah; setiap baris sebab jatuh tepat di satu baris; jumlahnya sama dengan perubahan kekayaan bersih. Pemasukan tampil per sumber dan penerima seperti yang tercatat; tidak ada kalimat bahwa itu seluruh penghasilan.

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-37 *Then* tabel A menampilkan: pemasukan 10000000; pengeluaran −5461000; biaya kredit −360000; selisih kas belum dirinci −25000 dan +10000 terpisah; piutang diikhlaskan −100000; perubahan nilai aset −1000000; kontribusi dari dana pribadi 120000 (suami 80000, istri 40000); disisihkan ke dana pribadi −200000 (suami 100000, istri 100000); jumlah 2984000; kekayaan bersih 9100000 menjadi 12084000.
- **AC-2** — *Given* AC-1 dibuka suami dan istri *Then* semua angka identik.
- **AC-3** — *Given* ringkasan mana pun *Then* tidak ada teks "semua penghasilan", "total penghasilan kalian", atau yang menyiratkan kelengkapan; yang tampil "pemasukan yang tercatat".

### FR-FIN-151 — Pengeluaran per kategori
| | |
|---|---|
| Fitur | F-FIN-020 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | KS-3, MSL-1 |
| Aktor dan izin | Sama dengan FR-FIN-150 |
| Aturan acuan | 07 §6.3 tabel B, §3.13 (pengembalian dana); CTH-37 |
| Business rule | BR-FIN-091 |
| Kedudukan | 07 §6.3 |

Tabel B memuat pemakaian per kategori sesudah pengembalian dana, termasuk cicilan yang dibayar; biaya kredit ada di tabel A dan tidak pernah dijumlahkan dengan cicilan.

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-37 *Then* tabel B: tagihan rutin 1000000; kiriman keluarga 3200000; belanja bulanan 540000; perawatan diri 50000; biaya admin 1000; perlengkapan rumah 150000; jajan 150000; kesehatan 75000; hadiah 230000; transportasi 65000; jumlah pengeluaran 5461000; cicilan 1500000; jumlah tabel B 6961000.
- **AC-2** — *Given* AC-1 *Then* biaya kredit 360000 tidak tampil di tabel B.

### FR-FIN-152 — Siapa membayar apa
| | |
|---|---|
| Fitur | F-FIN-020 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | KS-3 (kriteria sukses nomor 3, "who paid what"); MSL-7 (dokumen 17 "Keluhan" 8) |
| Aktor dan izin | Sama bagi keduanya; hanya buku rumah tangga. **Ditandai menurut tugas agen ini:** tabel per orang dapat dibaca sebagai papan perbandingan; pagarnya di bawah |
| Aturan acuan | 07 §6.3 tabel C, §10.3 butir 2; 06 AT-7; CTH-37 |
| Business rule | BR-FIN-092, BR-FIN-093, BR-FIN-097 |
| Kedudukan | D-01 (pagar pemantauan, checkpoint S1): tiga kolom, sama bagi keduanya, tanpa persentase dan peringkat, dengan keterangan artinya; lingkupnya: ditampilkan selama pilot dalam bentuk ini, diterima lewat "lanjut" di checkpoint S3 (12 November 2026) sebagai D-74 (Q-27). Contoh layarnya dokumen 08 |

Tabel C: pengeluaran dan pembayaran cicilan dikelompokkan menurut pembayar, dengan tiga kolom sumber: dana bersama; dana pribadi sebagai kontribusi; dana pribadi yang diganti (talangan). Per kategori jumlahnya sama dengan tabel B. Urutan pembayar tetap (misalnya urutan anggota di rumah tangga), tidak menurut besarnya. Tidak ada persentase, bagian, grafik perbandingan, atau angka lintas periode. Keterangan yang selalu tampil bersama tabel, bukan tersembunyi: tabel ini menunjukkan **lewat tangan siapa uang dibayarkan**, bukan siapa yang lebih banyak menanggung atau membelanjakan; orang yang menerima gaji dan membayar tagihan dengan sendirinya tercatat membayar lebih banyak.

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-37 *Then* tabel C: suami 5816000 / 80000 / 135000 = 6031000; istri 890000 / 40000 / 0 = 930000; jumlah 6706000 / 120000 / 135000 = 6961000, sama dengan tabel B.
- **AC-2** — *Given* AC-1 *Then* tidak ada persentase, peringkat, urutan menurut besar, atau angka bulan lain; keterangan "lewat tangan siapa uang dibayarkan" tampil bersama tabel.
- **AC-3** — *Given* AC-1 dibuka suami dan istri *Then* tabel dan keterangannya identik.
- **AC-4** — *Given* ringkasan beberapa periode *Then* tabel C tidak dijumlahkan lintas periode per orang; yang ada hanya total rumah tangga.

### FR-FIN-153 — Tabungan, aset, piutang, dan kewajiban di ringkasan
| | |
|---|---|
| Fitur | F-FIN-020 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | KS-3, MSL-1 |
| Aktor dan izin | Sama dengan FR-FIN-150 |
| Aturan acuan | 07 §6.3 tabel D dan E, §3.9 "Menyelesaikan atau membatalkan tujuan", §3.12 (klaim yang disengketakan); CTH-37, CTH-45 |
| Business rule | BR-FIN-074, BR-FIN-075, BR-FIN-091 |
| Kedudukan | 07 §6.3. Tujuan terarsip di tabel D: **[Usulan]** `finance-domain`, putaran S3 (07 §3.9, §6.3; menutup G-01) |

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-37 *Then* tabel D: "dana darurat" 5000000 menjadi 5500000 (perubahan 500000), target 10000000.
- **AC-2** — *Given* CTH-37 *Then* tabel E: motor 12000000 → 11000000; laptop 0 → 2000000; kulkas 0 → 4000000; piutang adik 0 → 200000; cicilan motor 9000000 → 7500000; cicilan kulkas 0 → 3960000; utang kakak istri 0 → 600000; kekayaan bersih akhir 12084000.
- **AC-3** — *Given* klaim internal yang punya keberatan terbuka *Then* tabel E menampilkannya dengan keterangan sengketa (kata layar dokumen 08) dan jumlahnya tidak berubah.
- **AC-4** — *Given* CTH-45 (FR-FIN-128/AC-3) *Then* tabel D periode pelepasan menampilkan "dana darurat" awal 5800000, perubahan −5800000, akhir 0, keterangan "diarsipkan"; tabel D periode berikutnya tidak memuatnya; tabel A dan kekayaan bersih tidak berubah karenanya.

### FR-FIN-154 — Arus kas dana bersama
| | |
|---|---|
| Fitur | F-FIN-020 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-1, MSL-3 |
| Aktor dan izin | Sama dengan FR-FIN-150 |
| Aturan acuan | 07 §6.4, INV-19; CTH-37 |
| Business rule | BR-FIN-091 |
| Kedudukan | 07 §6.4 |

Pemindahan, kontribusi, talangan, biaya kredit, dan selisih yang sudah dirinci tidak tampil di arus kas.

**Kriteria penerimaan**
- **AC-1** — *Given* CTH-37 *Then* uang masuk 11210000 (pemasukan 10000000, pinjaman 1000000, pelunasan piutang 200000, selisih positif 10000); uang keluar −10366000 (pengeluaran −5146000, cicilan dan utang −1900000, penggantian −135000, pinjaman diberikan −500000, aset −2400000, disisihkan −200000, selisih negatif −85000); perubahan 844000; uang yang ada 6100000 menjadi 6944000.
- **AC-2** — *Given* pemindahan 500000 `TUN-I` → `TAB-I` *Then* tidak muncul di arus kas.

### FR-FIN-155 — Angka per anggota hanya per periode, tanpa hitungan
| | |
|---|---|
| Fitur | F-FIN-020 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-7 (dokumen 17 "Keluhan" 7) |
| Aktor dan izin | Berlaku sama bagi keduanya |
| Aturan acuan | 07 §6.3 (per anggota hanya dalam satu periode); 06 AT-4, AT-7; SP-43 butir (b) |
| Business rule | BR-FIN-093, BR-FIN-097 |
| Kedudukan | D-43 (pagar, checkpoint S1). Tanpa hitungan tanda rumah tangga: **[Usulan]** `security-privacy` (SP-43 b) |

**Kriteria penerimaan**
- **AC-1** — *Given* ringkasan beberapa periode atau tampilan tahunan *Then* tidak ada angka kumulatif, rata-rata, tren, atau bersih per anggota atas penyisihan dan kontribusi; yang ada total rumah tangga.
- **AC-2** — *Given* periode dengan entri bertanda, keberatan, koreksi, dan catatan terlambat *Then* ringkasan bulanan tidak memuat hitungan atasnya, per orang maupun per rumah tangga.
- **AC-3** — *Given* sisa jatah pasangan *Then* tidak tampil di ringkasan bulanan (D-54: hanya bila dibuka di layar kesepakatan jatah).

### FR-FIN-156 — Ringkasan periode lampau yang berubah
| | |
|---|---|
| Fitur | F-FIN-020 |
| Tier / MoSCoW | 1 / Must |
| Masalah (brief §1.1) | MSL-5 (catatan terlambat sah), MSL-1 |
| Aktor dan izin | Keduanya melihat keterangan yang sama |
| Aturan acuan | 07 §3.11, §6.3 |
| Business rule | BR-FIN-094, BR-FIN-099 |
| Kedudukan | D-05 |

**Kriteria penerimaan**
- **AC-1** — *Given* ringkasan periode Oktober sudah dibuka *When* entri bertanggal Oktober dicatat pada November *Then* ringkasan Oktober berubah dan menampilkan "berubah sejak terakhir dilihat" kepada setiap orang yang sudah membukanya; tidak menyebut siapa yang mencatatnya terlambat.
- **AC-2** — *Given* periode mana pun *Then* tidak ada perintah "tutup bulan" yang mencegah entri terlambat.

## 12. Butir terbuka dari berkas ini

**Celah aturan.** Putaran penyelarasan 12 November 2026: G-01 sampai G-07 tertutup. Pengenalnya dipertahankan sebagai jejak dan tidak dipakai ulang.

| # | Hal | Status dan tempat aturannya | Yang berubah di berkas ini |
|---|---|---|---|
| G-01 | Menyelesaikan atau membatalkan tujuan tabungan yang masih beralokasi; tujuan terarsip di tabel D | **Tertutup.** 07 §3.9 "Menyelesaikan atau membatalkan tujuan", §2.6 baris "Tujuan", §6.3, T-38, CTH-45; bawaan berkas ini diambil, ditambah perlakuan tabel D | BR-FIN-075; FR-FIN-128 (ditulis ulang, AC-1 sampai AC-4); FR-FIN-153/AC-4 |
| G-02 | Hitungan offline yang tiba sesudah entri lain | **Tertutup.** 07 §3.10 "Hitungan yang tiba belakangan", CTH-46; 06 §12 S-10. Selisih dihitung server | BR-FIN-080; FR-FIN-135 (spesifikasi, AC-7); FR-FIN-136/AC-7 |
| G-03 | `write_off` terhadap ambang | **Tertutup.** 07 §3.7, §3.8, T-32, CTH-25; 06 §6.2 baris `write_off` (pelaku, alasan wajib, tidak ditautkan ke otorisasi pinjaman asal) | BR-FIN-067; FR-FIN-117 (aktor, masukan: alasan kini **wajib**, bukan opsional; AC-1, AC-4 sampai AC-6; kumulatif SP-51 fase D: AC-4, AC-7 sampai AC-9) |
| G-04 | Kesepakatan rutin saat tagihan diarsipkan; tagihan yang terarsip sendiri | **Tertutup.** 07 §3.15, §2.6, T-36 (tagihan tidak terarsip sendiri); 06 §6.3 "Pos rutin" baris (f) (arsip + K06 satu transaksi). Bawaan berkas ini untuk arsip otomatis **dicabut** | BR-FIN-050; FR-FIN-107 (ditulis ulang, AC-1, AC-5 sampai AC-8); FR-FIN-103/AC-13 |
| G-05 | Penghitung tempat uang dana bersama yang dipegang orang lain | **Tertutup.** 07 §3.10 butir 1, §2.5 baris (a) dan (c), T-33, CTH-47; 06 §7.2 baris G-05 | BR-FIN-080; FR-FIN-135 (aktor, AC-8, AC-9; batas SP-50 fase D: AC-11 sampai AC-13) |
| G-06 | Pelaku penilaian aset dan pemberitahuannya | **Tertutup.** 06 §8.1, N-14 (`ringkasan`) | FR-FIN-131 (aktor, AC-1) |
| G-07 | Uang muka + cicilan < harga | **Tertutup.** 07 §3.14 "Biaya kredit tidak pernah negatif", T-35, CTH-49 | BR-FIN-064; FR-FIN-111 (spesifikasi, AC-4, AC-5) |
| G-08 | Siapa yang boleh mengarsipkan (dan mengaktifkan kembali) tujuan tabungan dana bersama | **Tertutup.** 06 §8.1 baris "Mengarsipkan tujuan tabungan" (revisi kesepuluh, **[Usulan]** `architect`): anggota dewasa mana pun, dengan syarat alokasi nol (07 §2.6); N-14 `ringkasan` kepada anggota dewasa lainnya. Bawaan berkas ini diambil, ditambah N-14 | FR-FIN-128 (aktor) |

Pengenal `G-n` dipertahankan karena dokumen 07 dan 06 merujuknya; pertanyaan checkpoint di bawah diganti dari `P-n` menjadi `PA-FIN-12` sampai `PA-FIN-18` (README dokumen, skema `PA-<MOD>-N`; putaran perbaikan fase D, CA-06), dengan `PA-FIN-(11 + n)` untuk `P-n` lama.

**Pertanyaan untuk checkpoint S3** (diusulkan masuk dokumen 17; nomornya diberikan lead). **Hasil (12 November 2026):** semua bawaan di tabel ini diterima lewat "lanjut" (PA-FIN-12 sebagai S3-5 (i), PA-FIN-17 sebagai S3-9 (ii), PA-FIN-18 sebagai S3-9 (i), PA-FIN-15 sebagai S3-1 (b), PA-FIN-13 sebagai D-74, PA-FIN-14 sebagai D-75, PA-FIN-16 sebagai D-76); tabel dipertahankan sebagai catatan:

| # | Pertanyaan dalam bahasa sehari-hari | Bawaan yang diusulkan |
|---|---|---|
| PA-FIN-12 (dulu P-1) | Kini 06 P-26; tidak ditanyakan dua kali. Bila salah satu dari kalian mengakhiri sendirian kesepakatan tagihan rutin, atau mengarsipkan tagihannya (yang ikut mengakhiri kesepakatannya, FR-FIN-107), pembayaran berikutnya kembali ke aturan ambang. Boleh sendirian, atau harus berdua? (03 §7 butir 9; Q-113) | Boleh sendirian, berlaku seketika, diberitahukan (06 P-26, §6.3 baris (d) dan (f), **[Usulan]** `architect`); dinilai terhadap PG-14 di dokumen 10 |
| PA-FIN-17 (dulu P-6) | Kini 07 §10.4 butir 3, checkpoint S3-9 (ii). Mengikhlaskan pinjaman keluarga di atas ambang: ditandai seperti kiriman sebesar itu, dan alasannya wajib diisi? Dihitung kumulatif per pinjaman: mengikhlaskan 600000 dalam dua kali 300000 menandai yang kedua (FR-FIN-117/AC-7, penghapusan pertama dan kedua: 0 + 300000 = 300000 tidak ditandai, 300000 + 300000 = 600000 > 500000 ditandai); bagian yang disetujui lewat permintaan tidak dihitung (AC-8); satu potong 600000 sekaligus ditandai (AC-4) | Ya, ditandai, kumulatif per pinjaman; alasan wajib (07 T-32, INV-34; 06 §6.2; SP-51) |
| PA-FIN-18 (dulu P-7) | Kini 07 §10.4 butir 2, checkpoint S3-9 (i). Bolehkah satu orang mencatat hasil hitung atas rekening bersama yang dipegang pasangannya, dengan catatan "belum dikonfirmasi" sampai pemegangnya menjawab? Hanya bila pemegang menandai rekening itu "pasanganku dapat melihat saldo rekening ini", hanya satu hitungan yang menunggu jawaban sekaligus, dan tidak pernah atas uang tunai (FR-FIN-135) | Boleh, dengan batas itu (07 T-33, INV-33; 06 §7.2; SP-50) |
| PA-FIN-13 (dulu P-2) | Ringkasan "siapa membayar apa" ditampilkan dalam bentuk FR-FIN-152 (tiga kolom, keterangan "lewat tangan siapa"), atau tidak ditampilkan selama pilot? (Q-27, D-01) | Ditampilkan dalam bentuk FR-FIN-152 |
| PA-FIN-14 (dulu P-3) | Ajakan menghitung uang hanya muncul saat saldo suatu tempat uang negatif, kepada pemegangnya; tidak ada jadwal atau pengingat mencatat (03 §7 butir 8; Q-90) | Ya |
| PA-FIN-15 (dulu P-4) | Tabungan tidak masuk angka utama; satu baris "disisihkan untuk tabungan" di halaman depan (Q-123 butir a) | Ya |
| PA-FIN-16 (dulu P-5) | Saat membuat tagihan rutin, satu angka "jumlah tertinggi yang wajar" sekaligus menjadi perkiraan pengingat | Ya, perkiraan dapat diisi lain bila mau |
