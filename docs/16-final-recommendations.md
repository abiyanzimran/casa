# 16 — Rekomendasi Akhir

| | |
|---|---|
| Pemilik | Lead |
| Tahap | S5 — Rencana dan sintesis, fase C (12 November 2026) |
| Status | Ditinjau `critic` (S5 B, lalu tinjauan ulang B2, 12 November 2026); temuan diperbaiki |
| Sumber yang mengikat | Brief §0.3, §1, §1.1, §16 (baris 16: "MVP, architecture, critical decisions, deferred decisions, features to reconsider, top opportunities"); `CLAUDE.md`; keputusan pemilik produk di [17](17-open-questions-assumptions.md) dan [_progress.md](_progress.md) |
| Merangkum | 02 sampai 15 dan 18. Dokumen ini tidak menetapkan aturan, persyaratan, atau keputusan baru; setiap butir menunjuk sumbernya |
| Dipakai oleh | [01](01-executive-summary.md) |

---

## 0. Cara membaca

- **Kedudukan setiap butir** ditulis terang: *diputuskan pemilik produk* (keputusannya sendiri), *diterima* (atas saran lead atau lewat "lanjut" di checkpoint), atau **[Usulan]** (rekomendasi tim yang belum diputuskan). Kata pemilik produk hanya dikutip dari [17](17-open-questions-assumptions.md) atau [_progress.md](_progress.md).
- **Tier mengatur urutan membangun, bukan batas rancangan** (brief §1.1). Semua model, aturan, alur, istilah, dan layar dirancang untuk setiap rumah tangga tahap publik; rumah tangga pilot adalah pengguna pertamanya.
- Label bukti mengikuti [README](README.md#label-bukti). Istilah mengikuti [00-glossary.md](00-glossary.md).

---

## 1. Rekomendasi dalam satu halaman

1. **Bangun MVP yang menjawab satu pertanyaan lebih dulu:** "sisa uang berapa, dan siapa yang memegang" bisa dijawab kedua pasangan dari aplikasi tanpa saling bertanya (brief §1.1 kriteria sukses 1; PRN-01). Seluruh Tier 1 melayani pertanyaan itu, permintaan dan persetujuan yang mendahului pengeluaran, dan chat rumah tangga tempat keduanya dibahas.
2. **Pertahankan arsitektur yang sudah diputuskan** (kedudukan per lapisan di §3): aplikasi native per platform (SwiftUI; Jetpack Compose Material 3), stack diputuskan pemilik produk; satu server Go (keputusan 29) dengan PostgreSQL di server di Indonesia (keputusan 10; bentuk basis data diterima lewat "lanjut", D-99); semua aturan uang dan persetujuan hanya di server (keputusan 30; ADR-013); dan pencatatan yang tidak pernah dihalangi jaringan (PRN-06).
3. **Masuk ke data nyata hanya lewat gerbang GR-2** (13 §3.3; 14 QG-30): semua kriteria penerimaan Tier 1 lulus sebagai uji bernama, sembilan butir 10 §4.1 selesai, dan aplikasi terpasang di ponsel kedua pasangan.
4. **Ukur pilot dengan ukuran yang tidak mengawasi siapa pun** (13 §6; VL-00 sampai VL-10), lalu putuskan di GR-3 sesudah tiga bulan apakah lanjut ke tahap publik.
5. **Untuk tahap publik:** freemium per rumah tangga dengan satu langganan rumah tangga (Plus) **[Usulan]** (13 §7.5); advokat, pendaftaran PSE, dan beta tertutup lebih dulu (GR-4, GR-5).
6. **Keputusan berikutnya ada pada pemilik produk** (§5, §9). Yang terdekat: platform yang dibangun lebih dulu, Mac sewa atau beli, hosting dan email (S4-2), jenis akun Google Play, dan usulan tinjauan keamanan sebelum data nyata dan sebelum pengguna luar pertama.

---

## 2. MVP yang direkomendasikan

MVP = Tier 1 + Tier 2 (glosarium *MVP*). Rinciannya di [04 §9](04-feature-map.md#9-potongan-mvp-untuk-checkpoint-s3); persyaratannya di [05](05-functional-requirements.md).

| | Isi | Hitungan | Kedudukan |
|---|---|---|---|
| **Tier 1 — Inti** (FS-1, dibangun sebelum data nyata) | `FAM` satu rumah tangga dua dewasa setara, termasuk kunci aplikasi dan halaman Bantuan; `FIN` custody, uang masuk dan pembagiannya, pengeluaran cepat, serah terima, jatah pribadi, anggaran bulanan per kategori, tagihan berulang dan pos rutin, cicilan, utang dan piutang termasuk keluarga besar, tabungan sederhana, aset dan kekayaan bersih, ringkasan; `REQ` permintaan dua arah sampai penyelesaian sisa; `NTF` dengan jam tenang dan ringkasan; `COM` chat rumah tangga ringan dengan kartu objek, tanda terkirim dan dibaca, kehadiran, sembunyikan untuk saya dan laporkan; `CAL` pengingat tagihan | 55 fitur | Potongan S3 diterima lewat "lanjut". Tiga fitur Tier 1 tambahan di checkpoint S4: F-COM-014 **diputuskan pemilik produk** (keputusan 37, S4-7); F-FAM-019, F-FAM-020 rumusan lead yang dikukuhkan lewat "lanjut" (S4-6 (iii); D-114) |
| **Tier 2 — Versi ringan** (FS-2, bergelombang di dalam pilot) | Sebutan anggota; permintaan tindakan; daftar belanja; masalah rumah menjadi tugas; status suasana hati dan tindak lanjut sakit; rutinitas; rencana mendatang; pendamping `AI` versi ringan (kumpulan kalimat R1) | 18 fitur | Diterima (checkpoint S3 dan S4; sebutan anggota permintaan pemilik produk sendiri, keputusan 40) |
| **Tier 3 — Ditunda** | Selebihnya, termasuk struktur keluarga lain, `INV`, `MEAL`, `TML`, integrasi | 58 fitur | Garis besar satu halaman per modul. Klien web bukan fitur 04; ia calon di FS-8 (13 §3.8) |

Seluruhnya: 131 fitur; 218 FR (Must 198, Should 20), 206 BR (05); 49 skenario (12). Setiap Must terpetakan ke fitur, skenario, dan fase roadmap: 190 di FS-1 dan 8 Must bersyarat di FS-2 (05 §6; 13 §4).

**Tiga konsep dokumen 11 di Tier 1 sebagai sub-fitur** (diterima lewat "lanjut", S3-2): hitung sisa uang tunai sepulang belanja (IDEA-01), pos rutin tanpa tanggal jatuh tempo seperti token listrik dan galon (IDEA-03), dan "bagikan seperti terakhir kali" untuk uang masuk (IDEA-02). Ketiganya memangkas usaha mencatat yang berulang setiap bulan, sesuai PRN-02 dan PRN-05.

**Yang sengaja tidak dibangun di pilot** (04 §8; 17): web atau PC (keputusan 41); pelacakan siklus (D-63); kalkulator zakat (11 IDEA-09); voice note, reaksi, grup, dan pencarian chat (D-70); sinkron bank (brief §1; 18 §3.1); perintah mengeluarkan orang dewasa lain (S3-4 A). Aplikasi mencatat dan menata uang; ia tidak memegang, memindahkan, atau memproses uang sungguhan (brief §1).

---

## 3. Arsitektur

Rinciannya di [09](09-technical-architecture.md) dan ADR-012 sampai ADR-028; roadmap teknisnya di [14](14-engineering-roadmap.md).

| Lapisan | Rekomendasi | Kedudukan |
|---|---|---|
| Aplikasi | UI native penuh: Swift (SwiftUI) untuk iPhone dan iPad, Kotlin (Jetpack Compose Material 3) untuk Android; iOS dan iPadOS 26, Android 10 (API 29) | **Diputuskan pemilik produk**: stack (`_progress.md` "Keputusan pemilik produk: stack mobile"); OS minimum (17 keputusan 32, S4-3) |
| Letak aturan | Hanya di server; perangkat tidak menghitung aturan uang atau persetujuan; tiga bantuan isian dengan vektor emas sebagai pengecualian ES §2 | **Diputuskan pemilik produk** (keputusan 30, S4-1 (ii); ADR-013) |
| Server | Satu monolit modular Go, satu PostgreSQL 18 dengan RLS dan transaksi `SERIALIZABLE` per perintah tulis, kontrak Protobuf + Connect | Go: **diputuskan pemilik produk** (keputusan 29, S4-1 (i); ADR-015). Bentuk server dan basis data: diterima sebagai default lewat "lanjut" (D-99, D-100; ADR-014, ADR-016, ADR-017) |
| Sinkron | Antrean di perangkat, SSE + tarik berkursor saat aplikasi terbuka, push umum ke APNs dan FCM saat tertutup | Diterima sebagai default lewat "lanjut" (D-100, D-101; ADR-018, ADR-019) |
| Hosting, cadangan, email | Satu VPS di Indonesia, cadangan terenkripsi ke penyedia Indonesia lain, email transaksional minimal | **Ditunda sampai menjelang deploy**, keputusan pemilik produk (keputusan 31, S4-2); syarat data di Indonesia (keputusan 10) tetap mengikat |
| `AI` | Jalur R1: kumpulan kalimat dibangun tanpa data rumah tangga; Claude; batas USD 10 per bulan dan USD 3 per pembangunan | Claude: **diputuskan pemilik produk** (keputusan 33, S4-4 (ii)). R1 dan batas biaya: rumusan lead yang dikukuhkan lewat "lanjut" (S4-4 (i)); angka batas **[Asumsi]** (ADR-021) |
| Build iOS | Xcode Cloud + TestFlight internal, dengan sewa Mac sekali | **Diputuskan pemilik produk** (keputusan 32, S4-3 (i); ADR-022) |
| Lingkungan dan rilis pilot | Staging di VPS yang sama sejak GR-2, tanpa sistem feature flag, saklar *tahan tulis* yang tidak pernah menahan perintah pelindung; batas rilis iOS 80 hari sejak unggah dengan rilis pemeliharaan di 60 hari | **[Usulan]** `architect` (ADR-028, Diusulkan). Batas rilis yang mengikat sampai pemilik produk menjawab Q-219 tetap ADR-022 butir 1: paling lambat 90 hari, sasaran 60 |

**Biaya pilot per bulan** menurut 09.4 §7.1: **Rp260.667** + PPN bila berlaku, ditambah **USD 9,57** (paling tinggi USD 18,25 bila batas `AI` tercapai), belum termasuk email; Rp244.000 dari angka rupiah itu bergantung pada S4-2. Satu VPS sederhana, sesuai brief §1.

**Tiga tradeoff yang paling terasa saat membangun** (09 §1):

1. **Server saja** menghemat satu bahasa dan satu langkah build yang hanya bisa di macOS, tetapi perangkat tidak menampilkan sisa atau tanda saat offline sampai server menjawab.
2. **Tanpa Mac**, preview SwiftUI, simulator, dan debugger tidak ada di Linux (18 §5.2.7.4); setiap perubahan UI iOS dilihat lewat build Xcode Cloud dan TestFlight. Pemicu untuk punya Mac ada di ADR-022 butir 4 (§5).
3. **R1** menjaga data rumah tangga tetap di Indonesia dan biaya `AI` tetap, dengan harga bahwa saran tidak menyebut hal yang spesifik.

---

## 4. Keputusan kritis yang sudah diambil

Dicatat agar tidak ditanyakan ulang. Rincian dan kata-kata pemilik produk di 17.

| # | Keputusan | Kedudukan | Sumber |
|---|---|---|---|
| 1 | Dua dewasa setara, dengan pelaku terlihat pada setiap tindakan; tidak ada perintah mengeluarkan orang dewasa lain | Kesetaraan: **diputuskan pemilik produk** (keputusan 4, S1). Tanpa perintah mengeluarkan: diterima sebagai default lewat "lanjut" (S3-4 A) | PRN-09, PRN-10; 06 PG-01 sampai PG-15 |
| 2 | Setiap hal punya tingkat visibilitas; berbagi adalah pilihan pemiliknya; dana bersama tidak dapat disembunyikan | Tingkat visibilitas: **diputuskan pemilik produk** (keputusan 4). Rumusan prinsipnya diterima di checkpoint S2 | PRN-12, PRN-13 |
| 3 | Catatan lama tidak pernah disunting; selisih dan entri terlambat dicatat sebagai entri baru; jumlah uang bilangan bulat | Tidak menyunting catatan lama: **diputuskan pemilik produk** (keputusan 7, S1). Bacaan append-only: **diputuskan pemilik produk sebagai keputusan sementara sampai advokat menjawab** (keputusan 24, S2-8) | ADR-002, ADR-004, ADR-007, ADR-008 |
| 4 | Pemakaian besar yang tidak diajukan lebih dulu tetap tersimpan dan ditandai netral, tidak dihalangi | Diterima atas saran lead (S2-2) | ADR-009; PRN-06 |
| 5 | Chat rumah tangga; dari chat dapat dikirim tugas, pengeluaran, dan rencana; tanda terkirim dan dibaca wajib; kehadiran seperti WhatsApp | Chat dan tanda: **diputuskan pemilik produk** (keputusan 26). Bentuk dan cakupannya: diterima atas saran lead. Kehadiran: **diputuskan pemilik produk**, berlawanan dengan saran lead (keputusan 27, dipertahankan di keputusan 35) | ADR-011 |
| 6 | Stack mobile Swift + Kotlin, native penuh | **Diputuskan pemilik produk** | ADR-012 |
| 7 | Backend Go; aturan hanya di server | **Diputuskan pemilik produk** (keputusan 29, 30) | ADR-013, ADR-015 |
| 8 | Tanpa E2EE di pilot; status kepedulian yang tidak dibagikan tidak meninggalkan perangkat; E2EE dinilai lewat ADR sebelum tahap publik | **Diputuskan pemilik produk** (keputusan 36, S4-6 (i)) | 10-D; ADR-024 |
| 9 | Pengelola server pilot adalah pasangan, dengan syarat istrinya diberi tahu terus terang sebelum mulai memakai | **Diputuskan pemilik produk** (keputusan 25, S2-10). Bentuk pemberitahuannya diterima sebagai default lewat "lanjut" (S4-6 (ii) A) | FR-FAM-019/AC-7; 15 RISK-11 |
| 10 | Advokat ditanya sebelum tahap publik; risiko PSE selama pilot diterima | **Diputuskan pemilik produk**, berlawanan dengan saran B tim (keputusan 34, S4-4 (iii); W-32) | 15 RISK-65 |
| 11 | Sembunyikan untuk saya, laporkan, kontak, keluar; tanpa pemindaian dan tanpa mengeluarkan orang dewasa | **Diputuskan pemilik produk** (keputusan 37, S4-7) | F-COM-014; 10.4 SEC-153, SEC-154 |
| 12 | Arah visual C, terang dan tenang | Diterima atas saran lead (S4-10 (i)) | 08 §19 |
| 13 | Aplikasi menyapa dengan "kamu" | Rumusan lead yang dikukuhkan lewat "lanjut" (S4-10 (iv)) | 08 |
| 14 | Web atau PC tidak dibangun di pilot | Diterima atas saran lead (keputusan 41) | 13 §3.8; Q-210 |
| 15 | Ambang dibuat, diubah, dan diakhiri hanya dengan persetujuan semua pihak, di setiap rumah tangga, pilot dan publik | **Diputuskan pemilik produk**, berlawanan dengan saran B dan bawaan C tim (keputusan 39, S4-9 (i)). Risiko sisa yang diterimanya: satu pihak dapat mempertahankan ambang atas pihak lain tanpa batas waktu dengan tidak menyetujui perubahan atau pengakhirannya, berlaku sama bagi setiap pihak (W-34, Pagar) | 17 keputusan 39, W-34 |

---

## 5. Keputusan yang ditunda, dan kapan diambil

| Keputusan | Kapan | Bahan yang sudah ada | Siapa |
|---|---|---|---|
| **Platform yang dibangun lebih dulu** | Awal fase bangun (`CLAUDE.md`) | Tiga pilihan beserta urutan kerja masing-masing. Saran tim **[Usulan]**: Android lebih dulu, karena siklus kerjanya di Linux paling cepat dan sewa Mac tertunda sampai layar SwiftUI dimulai; GR-2 tetap menunggu kedua platform (13 §3.2, Q-37; 14 §3.4). Kedua pasangan harus dapat memakai aplikasi sejak data nyata (brief §1) | Pemilik produk |
| **Mac: sewa bulanan atau beli Mac mini** | Saat kerja UI iOS menjadi harian (ADR-022 butir 4) | Mac mini M6 16/512 GB, konfigurasi minimum yang disarankan `researcher`, Rp21.499.000 **[Perlu validasi]**; sewa sekitar USD 149 atau EUR 149 per bulan; titik impas untuk 16/512 GB sekitar 7,1 sampai 7,9 bulan, untuk 16/256 GB sekitar 5,8 sampai 6,4 bulan (8,8 bulan terhadap sewa berkomitmen 12 bulan) **[Usulan]** (18 §5.2.7.6); hanya Mac sendiri memberi debugger di iPhone 13 dan iPad mini 7 (18 §5.2.7; 14 §9) | Pemilik produk |
| **Hosting, cadangan, email (S4-2)** | "nanti aja pas mau deploy" (keputusan 31). Apakah itu deploy server uji pertama (14 FS-1 langkah 1) atau deploy pertama yang memuat data nyata (13) ditanyakan di Q-220, sebelum FS-1 langkah 1 selesai | ADR-020, ADR-026 (Diusulkan); harga dan pihak kontrak dicek ulang saat itu | Pemilik produk |
| **Batas rilis iOS** | Sebelum build iOS pertama (FS-1) | ADR-022 butir 1 mengikat: paling lambat 90 hari sejak unggah ke TestFlight, sasaran 60 (keputusan 32). ADR-028 butir 4 mengusulkan batas keras 80 hari dengan rilis pemeliharaan di 60 hari **[Usulan]** (Q-219) | Pemilik produk |
| **Jenis akun Google Play** (pribadi atau organisasi) | Sebelum GR-4 | Akun pribadi baru wajib closed testing dengan sedikitnya 12 penguji selama 14 hari sebelum produksi (14 §6.7; 18 §2.4 **[Terverifikasi]**). Bahwa hanya pemakai Android yang dihitung adalah bacaan `researcher` **[Usulan]**, karena Google tidak menyatakannya terang; ukuran beta di 13 §6.4 bergantung pada jawabannya. Akun organisasi wajib bagi kategori "Financial products and services" dan "Health apps" dan memerlukan nomor D-U-N-S; apakah Family OS termasuk, dan apakah akun organisasi bebas dari syarat 12 penguji, **[Perlu validasi]** | Pemilik produk (Q-221) |
| **Definisi pemakaian harian yang bertahan** | Sebelum GR-2 **[Usulan]** (garis dasar VL-00 diambil sebelum GR-2) | VL-02 **[Usulan]**: sedikitnya 5 dari 7 hari pada setiap dari empat minggu terakhir (13 §6.1; Q-213) | Pemilik produk |
| **Isi dan harga Plus** | GR-5 | Calon isi dan daftar "tidak pernah berbayar" (13 §7.5; Q-36, Q-216); aturan PG-15 | Pemilik produk, sesudah wawancara dan beta |
| **Lanjut atau tidak ke tahap publik** | GR-3, sesudah tiga bulan pilot | VL-01 sampai VL-10 dan tabel keputusan 13 §6.2 | Pemilik produk |
| **Siapa mengoperasikan server bila pemilik produk berhalangan lama** | Sebelum GR-2 | Memberi istri akses pengelola juga memberinya akses teknis ke data pribadi suaminya, cermin RISK-11 (Q-222, Pagar) | Pemilik produk dan istrinya |
| **Wawancara dan uji prototipe dengan keluarga lain sebelum GR-3**: termasuk "pengguna di luar rumah tangga pilot" atau tidak | Sebelum wawancara pertama | Penilaian `security-privacy` **[Usulan]**: tidak termasuk, dengan empat syarat (tanpa build terpasang dan tanpa data peserta di server; prototipe tanpa data nyata dan tanpa analitik; catatan tanpa nama dan tanpa angka keuangan peserta, di luar repositori dan server; peserta diberi tahu haknya menolak) (17 Q-215; 10 §7.4). Ketetapan lead atas L-01: jadwal advokat tetap menurut keputusan 34 | Pemilik produk (Q-215, Pagar) |
| **Klien web untuk tahap publik** | Sesudah pilot | Q-210; prasyarat model keamanan browser di 10 | Pemilik produk |
| Butir keamanan sebelum pengguna luar pertama (akses darurat SEC-36 dengan log di luar VPS, kode pemulihan SEC-25, dan lainnya) | Sebelum GR-4 | 10 §4.2 | Tim |
| E2EE chat dan foto, SQLCipher, passkey, jalur `AI` R2, uji penetrasi pihak luar (SEC-157) | Sebelum tahap publik (GR-5) | 09 §6; 10 §4 | Tim, satu ADR per butir bila mengubah arsitektur |

---

## 6. Yang perlu ditimbang ulang

Bukan usulan perubahan sekarang. Ini hal yang tim sarankan diperiksa dengan data pilot, karena buktinya belum ada atau risikonya tinggi.

| # | Hal | Mengapa ditimbang ulang | Kapan dan dengan apa |
|---|---|---|---|
| 1 | **Kehadiran "Semua seperti WhatsApp"** (online, terakhir aktif menit tepat, sedang mengetik) | Keputusan pemilik produk sendiri; `security-privacy` dan `ux-designer` mencatat bahwa kehadiran memperlihatkan kapan pasangan memegang aplikasi, juga di jam tenangnya (17 S4-5; W-18). Tidak diubah tanpa pemilik produk | GR-3: pertanyaan VL-08 apakah ada yang merasa diawasi (13 §6) |
| 2 | **Modul Tier 2 yang mungkin tidak dipakai** (`SHP`, `MNT`, `CAL`, `PLN`) | Setiap fitur harus membayar usaha berulangnya (PRN-02); risiko menambah beban navigasi (15 RISK-79) | Gelombang T2-3 (Q-212); dicabut bila tidak dipakai |
| 3 | **Mutu kumpulan kalimat `AI` R1** | Belum diuji; saran bisa terasa palsu atau berulang bila keduanya memakainya (15 RISK-74) | Gelombang T2-2; F-AI-005 lebih dulu |
| 4 | **Alur persetujuan yang terasa mengekang** | Permintaan dua arah adalah pembeda (18 §1.4), tetapi dapat dibaca sebagai kontrol (15 RISK-21, RISK-31) | VL-08 dan percakapan akhir bulan pilot |
| 5 | **"Sembunyikan angka" (F-FIN-027)** | Tetap Tier 3 Could (S4-10 (iii) A); angka tampil saat aplikasi dibuka, tombol mata menyembunyikan | Bila pilot atau beta menunjukkan perlunya |
| 6 | **Tiga konsep Tier 2 yang belum punya FR** (IDEA-15 tinjauan bulanan, IDEA-19 "butuh tenang hari ini", IDEA-08 templat Mudik dan Lebaran) | Diterima di S3-2, tetapi persyaratannya belum ditulis (04 §9.4) | FR ditulis `product-analyst` sebelum gelombang T2-3 (13 §12); IDEA-08 sebelum Ramadan pertama di dalam pilot (15 RISK-27) |
| 7 | **Struktur Phase 0 dan Phase 2 brief §13** | 13 mengubahnya: validasi pilot tidak menahan pembangunan, dan Phase 2 berarti membawa MVP ke publik, bukan menambah modul (13 §10 TB-2, TB-3) | Diterima bila 13 diterima; brief §13 sendiri mengizinkan struktur lain |

---

## 7. Peluang teratas

Dasar pembedaannya [18 §1.4](18-market-research.md#14-peluang-pembedaan-positioning-dan-value-proposition), seluruhnya **[Usulan]**: halaman pesaing yang diperiksa tidak menyebut hal-hal ini, tetapi itu belum membuktikan fiturnya tidak ada di produk mereka. Dua atau tiga aplikasi pesaing perlu diuji langsung sebelum diklaim di materi publik (18 §6.4; 13 GR-5).

| # | Peluang | Mengapa ini peluang | Di MVP sebagai |
|---|---|---|---|
| 1 | **Custody sebagai konsep utama**: "sisa berapa dan siapa memegang" di layar pertama | Kriteria sukses 1; tidak disebut di halaman pesaing yang diperiksa khusus | F-FIN-001, F-FIN-002, F-FIN-006 (Tier 1) |
| 2 | **Permintaan dua arah yang selesai sampai uangnya**: usul, diskusi, usul balik, persetujuan, belanja, penyelesaian sisa | Kriteria sukses 4; menjawab "jalan-jalan diumumkan mendadak" (brief §1.1) | `REQ` Tier 1 |
| 3 | **Arus uang keluarga Indonesia apa adanya**: kiriman ke orang tua, bantuan dan pinjaman ke saudara, uang tunai, tarik ATM, serah terima, THR | Pesaing memodelkan utang secara umum; polanya belum diuji di luar pilot (18 §1.4 **[Perlu validasi]**) | F-FIN-013, IDEA-01, IDEA-03; IDEA-08 Tier 2 |
| 4 | **Uang dan perhatian dalam satu tempat**: status suasana hati sukarela, tindak lanjut sakit, kata dukungan | Bagian dari tujuan produk ("attentive and kind"); nilainya bagi keluarga lain belum diuji | `CARE` dan `AI` Tier 2 |
| 5 | **Privasi yang dapat diperiksa**: tanpa iklan, tanpa penjualan data, data di Indonesia, tidak ada yang mengawasi siapa | Syarat masuk, bukan pembeda besar (pesaing lokal sudah mengklaim tanpa iklan); menjadi pembeda bila pagar antaranggota terbukti di beta | 06 pagar; 10; model bisnis tanpa iklan (13 §7.1) |

**Yang bukan pembeda:** pencatatan cepat dengan `AI` sudah ditawarkan pesaing lokal. Family OS harus setara cepatnya (tiga ketukan, simpan lokal seketika; ES §3), tidak bertumpu padanya.

---

## 8. Risiko yang paling menentukan

Register lengkapnya di [15](15-risk-assessment.md) (RISK-01 sampai RISK-90). Lima kelompok yang dapat menggagalkan tujuan (15 §1):

1. **Hanya satu pasangan mencatat, mencatat terlalu lama, atau catatan terlewat** (RISK-18 sampai RISK-20). Penawarnya: data nyata dimulai di kedua platform, tiga ketukan, entri terlambat diterima, rekonsiliasi kas.
2. **Angka uang salah** (RISK-47, RISK-48). Penawarnya: setiap invarian 06 dan 07 sebagai uji yang memblokir merge, kunci idempotensi, satu transaksi per perintah; bila pemeriksaan invarian harian tidak nol, server berhenti menerima tulisan baru lewat *tahan tulis* sampai penyebabnya diperbaiki, sementara ponsel tetap menyimpan catatan di antrean dan perintah pelindung tetap berlaku seketika (ADR-028 butir 3, masih Diusulkan; SEC-151; PRN-06).
3. **Aplikasi menjadi alat kendali antaranggota** (RISK-11, RISK-35, RISK-38, RISK-40). Penawarnya: pagar 06 sebagai uji, simetri per peran, keluar dengan aman tanpa persetujuan siapa pun, halaman Bantuan tanpa jejak.
4. **Keputusan GR-3 atas bukti lemah**, karena pilot satu rumah tangga dijalankan salah satu pasangan yang juga pengelola server (RISK-25). Penawarnya: pertanyaan tentang aplikasi, bukan tentang pasangan; beta dengan pengelola yang bukan anggota.
5. **Tahap publik terhalang**: PSE, kesediaan membayar, satu orang memegang seluruh pengetahuan operasi (RISK-65, RISK-58, RISK-10).

---

## 9. Usulan proses untuk fase bangun

Butir ini menyangkut `CLAUDE.md`, jadi **tidak diterapkan tanpa persetujuan pemilik produk**.

| # | Usulan | Asal | Akibat bila diterima |
|---|---|---|---|
| 1 | Tinjauan `security-privacy` hanya-temuan sebelum GR-2 (data nyata pertama) dan sebelum GR-4 (pengguna luar pertama), sebagai pengganti peran keamanan khusus di tim bangun | 14 §12 (`architect`) **[Usulan]** | Satu baris di `CLAUDE.md` bagian "Build team" |

---

## 10. Yang tersisa sesudah S5

- **S6 — pass konsistensi** (brief §0.3): memperbaiki kontradiksi, tautan rusak, istilah tanpa definisi, dan persyaratan yang tidak terlacak, termasuk butir W yang dibawa ke S6 di 17; laporan perubahan untuk pemilik produk.
- **Fase perencanaan berakhir hanya bila pemilik produk menyatakannya** (`CLAUDE.md`). Sesudah itu: agen tim bangun ditulis lead menurut stack yang dipilih, lalu GR-1.

---

## Riwayat

| Tanggal | Perubahan |
|---|---|
| 12 November 2026 | Draf pertama (lead, S5 fase C) |
| 12 November 2026 | Perbaikan atas `critic` S5 B (CS5-B-01 kedudukan keputusan) dan tinjauan ulang B2 (CR5-B-01, CR5-B-05, CR5-B-06, CR5-B-07, CR5-B-13): §1 butir 2 per keputusan bernomor; §4 baris 15 keputusan 39 dan W-34; §5 Q-215, syarat akun Play berlabel, kutipan keputusan 31 persis; salinan data sendiri (FR-FAM-013/AC-12) dibangun di FS-1 menurut ketetapan lead, jadi keluar dari daftar sebelum GR-4; §8 status ADR-028 |
