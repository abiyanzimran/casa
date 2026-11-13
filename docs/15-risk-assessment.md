# 15 — Penilaian Risiko

| | |
|---|---|
| Pemilik | `architect` |
| Tahap | S5 — Rencana dan sintesis, fase B (12 November 2026; izin pemilik produk "gas", `_progress.md` "S5 dimulai"); **S6 (12 November 2026): pass konsistensi**: §0 tinjauan tiga bulan dan GR-3; RISK-25, RISK-27, RISK-83 (jalur rumah tangga lewat GR-4 (a)); §8 baris FS-2, GR-4 (a), FS-3/GR-5, FS-4–FS-8; §10.2 Q-222, Q-224, Q-227, Q-228; ADR-028 diterima seluruhnya (keputusan 56). Tidak ada K, D, prioritas, status jalur, atau hitungan yang berubah; **sesudah S6 (12 November 2026): web lebih dulu** (keputusan 69; [ADR-030](adr/ADR-030-klien-web-pwa-lebih-dulu.md)) |
| Status | **Revisi web lebih dulu (12 November 2026):** RISK-01, RISK-18, RISK-91 dinilai ulang (RISK-91 D Tinggi → Sedang, P2 → P3); **RISK-94 sampai RISK-97 baru** (TM-44, TM-42, TM-38 dan TM-37, TM-45, web tidak terasa native); RISK-88 ditutup; §8 GR-1 diperbaiki (temuan audit 5) dan baris "Jalur iOS native" baru; §10.1 TM-37 sampai TM-45; hitungan di §12 butir 15. Sebelumnya: revisi S5 fase D sesudah tinjauan `critic` A dan B, lalu revisi keputusan 42 (platform pilot, 12 November 2026), lalu revisi keputusan 43 sampai 50 (12 November 2026), lalu revisi keputusan 51 dan 52 (tujuan tanpa rencana go public; jalur publik opsional; pengguna luar pertama; RISK-92 dinilai dengan 18 §5.2.9; RISK-93 baru, 12 November 2026), lalu revisi keputusan 53 sampai 55 (RISK-92 Diterima, RISK-40 faktor kedua opsional di FS-2). Hitungan §2 dan §12 dihitung ulang lead dengan skrip sebelum revisi; yang berubah di revisi ini ditandai di §12 |
| Sumber yang mengikat | Brief §1, §1.1, §4.3, §5D, §5O, §7, §8, §14, §16; `CLAUDE.md`; `ENGINEERING-STANDARDS.md` (ES); dokumen 02, 03, 06, 07, 08, 09 (indeks dan berkas 04 §9), 10 (indeks dan lima berkas, termasuk 10.5 klien web), 12, 13, 17, 18; ADR-002 sampai ADR-030 |
| Dipakai oleh | 14 (`architect`, quality gate per gerbang), 16 dan 01 (lead), 17 (`product-analyst`) |

---

## 0. Cara membaca

**Yang ditetapkan dokumen ini:** satu register risiko produk, UX, keamanan dan privasi, keuangan, teknis, bisnis, dan hukum; penilaian kemungkinan dan dampaknya; mitigasi yang **sudah ada** di dokumen lain beserta di fase atau gerbang mana ia harus berlaku; tanda awal; dan risiko sisa. **Yang tidak:** persyaratan baru. Bila sebuah mitigasi memerlukan persyaratan yang belum ada, ia ditulis sebagai permintaan kepada pemilik dokumennya (§11).

**Satu register.** Dokumen ini kanonis untuk pengenal `RISK-NN`.

| Sumber | Kedudukannya terhadap dokumen ini |
|---|---|
| 09.4 §9 (`RISK-01` sampai `RISK-17`) | Diambil alih dengan pengenal yang sama. Uraiannya dilengkapi di sini; 09.4 §9 kini hanya penunjuk ke dokumen ini (S5 fase D) |
| Dokumen 10 (`TM-NN`, `SEC-NN`, `PL-NN`, `AS-NN`) | Tetap kanonis untuk ancaman dan persyaratan keamanan. Dokumen ini mengelompokkan ancaman menjadi risiko, menunjuk persyaratannya, dan menyatakan risiko sisanya. Pemetaan lengkap `TM` ke `RISK` di §10.1 |
| Dokumen 17 (`L`, `R`, `A`, `W`, `K`, `D`, `Q`) | Tetap kanonis untuk butir hukum, riset, asumsi, peringatan, pertentangan, dan default. Dokumen ini menunjuknya dan tidak menyalin bunyinya |
| 06 (`PG`, `INV-A`), 07 (`INV`, §10.3), 08, 13 (`FS`, `GR`, `VL`) | Sumber mitigasi dan gerbang; dirujuk, tidak disalin |

**Catatan pengenal.** Di dokumen 10, `AS-NN` adalah **aset** (10.1 §1), bukan skenario penyalahgunaan. Skenario paksaan ada di 10.3 §2.7 tanpa pengenal; dokumen ini merujuknya dengan nomor bagiannya. Ancaman dokumen 10 sekarang `TM-01` sampai `TM-45` (TM-32 sampai TM-36 di 10.4 §6.2; TM-37 sampai TM-45 klien web di 10.5 §1).

**Skala.** Kemungkinan (K) dan dampak (D) dinilai **sebelum mitigasi**; risiko sisa dinilai sesudahnya. Semua nilai K, D, prioritas, dan risiko sisa adalah **[Usulan]** `architect`, kecuali yang disebut diterima pemilik produk.

| K | Arti, dalam fase yang disebut di kolom "Harus ada pada" |
|---|---|
| Rendah | Tidak diharapkan terjadi, tetapi mungkin |
| Sedang | Dapat terjadi |
| Tinggi | Diperkirakan terjadi bila tidak dimitigasi |
| Pasti | Keadaannya sudah berlaku menurut rancangan atau keputusan |

| D | Arti, diukur terhadap tujuan di `CLAUDE.md` dan brief §1.1 |
|---|---|
| Rendah | Gangguan kecil; tidak menyentuh KS-1 sampai KS-5 atau data |
| Sedang | Menunda satu fase atau melemahkan satu kriteria sukses |
| Tinggi | Menggagalkan salah satu KS-1 sampai KS-4, merusak kepercayaan pada angka, membuat aplikasi terasa sebagai alat tekanan, atau menahan satu gerbang |
| Kritis | Merugikan seorang anggota (dikendalikan, diawasi, data pribadi atau kesehatannya terbuka, akses ke pertolongan salah), angka uang salah tanpa terdeteksi, atau layanan diputus |

| D \ K | Rendah | Sedang | Tinggi | Pasti |
|---|---|---|---|---|
| **Kritis** | P2 | P1 | P1 | P1 |
| **Tinggi** | P3 | P2 | P1 | P1 |
| **Sedang** | P3 | P3 | P2 | P2 |
| **Rendah** | P3 | P3 | P3 | P3 |

P1 dibahas di ringkasan (§1) dan harus punya mitigasi yang teruji sebelum gerbangnya. P2 punya mitigasi tertulis dan pemilik. P3 dipantau.

**Kategori.** Tujuh: Produk, UX, Keamanan dan privasi, Keuangan, Teknis, Bisnis, Hukum dan regulasi. **Hukum dipisah** dari teknis dan bisnis karena tiga hal: (1) pemutusnya advokat dan pemilik produk, bukan tim; (2) mitigasinya kebanyakan pertanyaan advokat dan gerbang GR-4 atau GR-5, bukan rancangan; (3) sejak S4-4 (iii) A beberapa risiko hukum **sengaja diterima** untuk pilot (W-32), dan menggabungkannya ke kategori lain akan menyamarkan penerimaan itu.

**Penanda Pagar.** Risiko yang menyangkut penyalahgunaan antaranggota atau keselamatan seorang anggota (brief §4.3, §5D kalimat terakhir, §8; pelaku PL-02, PL-03, PL-05) diberi tanda **Pagar** di samping kategorinya. Brief §8 menjadikan pencegahan penyalahgunaan persyaratan utama; tanda ini memastikan risiko itu tidak tenggelam di kategori teknis.

**Kedalaman** (brief §1.1; `CLAUDE.md`):

| Kedalaman | Berlaku untuk | Bentuk |
|---|---|---|
| Penuh | Risiko yang mengenai Tier 1, data nyata pilot, atau gerbang GR-1 sampai GR-5 (GR-4 (a) sebelum pengguna luar pertama; GR-3, GR-4 (b), dan GR-5 di jalur publik opsional) | Kartu dengan semua kolom (§3) |
| Ringkas | Risiko fitur Tier 2 | Satu baris tabel dengan semua kolom (§4) |
| Satu baris | Risiko Tier 3 | Satu kalimat (§5) |

**Pemilik mitigasi** adalah peran, bukan orang: pemilik produk; pengelola server (glosarium; di pilot pemilik produk sendiri); peran tim bangun di `CLAUDE.md` (`backend-dev`, `web-dev`, `ios-dev`, `android-dev`, `ux-designer`, `code-reviewer`, `qa`; `web-dev` diminta ke lead, 14 §3.1); pemilik dokumen perencanaan (`architect`, `finance-domain`, `security-privacy`, `product-analyst`, `researcher`, lead); advokat.

**Fase dan gerbang** memakai 13 §1: GR-1 mulai bangun, FS-1 bangun Tier 1, GR-2 data nyata, FS-2 pilot dengan *tinjauan tiga bulan* di dalamnya (bukan gerbang; glosarium), GR-3 evaluasi go public (jalur publik opsional), GR-4 gerbang pengguna luar pertama dengan dua pintu, (a) di jalur keluarga dan (b) di jalur publik (glosarium; 13 §3.6), FS-3 persiapan publik dan beta, GR-5 rilis publik, FS-4 tahap publik, dan FS-5 sampai FS-8 Tier 3, yang dapat dipicu dari rumah tangga sendiri maupun dari jalur publik (13 §3.8).

**Jalur publik opsional (keputusan 51).** Diputuskan pemilik produk dengan mengukuhkan bacaan lead: aplikasi ini untuk keluarganya sendiri dengan standar kelas publik, dan go public tidak direncanakan tetapi tidak ditutup. Kata-kata jawabannya: "yap benar seperti itu, untuk sekarang gitu aja, tapi tidak menutup kemungkinan untuk go public" (`_progress.md`, "Keputusan pemilik produk 51"). Akibatnya bagi register ini:

- GR-3, FS-3, GR-4 pintu (b), GR-5, dan FS-4 adalah **jalur publik opsional** yang tidak dijadwalkan. GR-4 pintu (a) ada di jalur keluarga (butir pengguna luar pertama di bawah). Fase Tier 3 FS-5 sampai FS-8 tidak termasuk jalur publik (13 §3.8). Jalur rumah tangga berakhir pada rumah tangga yang memakai aplikasi sehari-hari: FS-1, GR-2, FS-2 yang berlanjut, gelombang Tier 2, dan jalur Android untuk Galaxy A55.
- Risiko yang **seluruhnya** timbul dari beta, toko, penjualan, keluarga lain, atau keputusan go public diberi status "**jalur publik**" di register (§2) dan di kartunya. Risiko itu tidak dihapus dan tidak dinilai ulang: K dan D tetap dinilai seolah jalur itu diambil (sesuai arti K di bawah, "dalam fase yang disebut di kolom Harus ada pada"), dan mitigasinya wajib berlaku bila pemilik produk kelak mengambil jalur itu.
- **Pengguna luar pertama (keputusan 52)** adalah jalur opsional tersendiri, terpisah dari tahap publik: orang tua yang diberi login dihitung pengguna di luar rumah tangga pilot (Q-209). Diputuskan pemilik produk sendiri, berlawanan dengan usulan `security-privacy`; kata-katanya: "orang tua yaa hitung pengguna luar, yang harus di tanyakan ke advokat apa?" (`_progress.md`, "Keputusan pemilik produk 52"). Jadi syarat sebelum pengguna luar pertama (10 §4.2; 14 QG-31 (a)) dapat diperlukan tanpa beta, tanpa rilis toko, dan tanpa go public. Risiko yang timbul sejak pengguna luar pertama, bukan hanya di tahap publik, diberi status "**pengguna luar**".
- Risiko yang sebagian berlaku bagi rumah tangga dan sebagian hanya di jalur opsional tetap berstatus biasa; bagiannya ditulis "sebelum pengguna luar pertama" atau "bila jalur publik diambil" di kolom "Harus ada pada".
- Tidak ada rancangan yang dipotong karena keputusan ini (brief §1.1; keputusan 51). Risiko Tier 3 (§5) mengikuti fasenya (FS-5 sampai FS-8) dan penempatannya di 13; karena belum dinilai, kecuali RISK-83, mereka tidak masuk hitungan pembagian jalur di §1.1.

**Web lebih dulu (keputusan 69; ADR-030).** Pemilik produk memutuskan: "web dibangun duluan aja yang penting web jalan dulu aja. ios dan android nanti. web akan jadi yang pertama untuk user experience", dan menerima bahwa web tidak terasa native (`_progress.md`, "Keputusan pemilik produk: web lebih dulu" dan "Jawaban pemilik produk"). FS-1 dan GR-2 kini memakai klien web di iPhone 13, iPad mini 7, Galaxy A55, dan desktop; iOS native (14 §3.9) lalu Android (14 §3.8) menyusul. **Aturan baca** untuk kartu yang belum ditulis ulang: pemilik mitigasi `ios-dev` atau `android-dev` untuk bagian klien yang jatuh di FS-1 atau GR-2 dibaca `web-dev`; bagian khusus iOS native (Keychain, kelas proteksi berkas, TestFlight, Xcode Cloud, XCUITest) jatuh di jalur iOS native (QG-30 (iOS)), bukan GR-2. K dan D risiko lama tidak berubah karena aturan ini; yang dinilai ulang disebut di kartunya dan di §12 butir 15.

**Risiko sisa** dipakai dengan arti: risiko yang tetap ada sesudah semua mitigasi yang disebut berlaku, beserta siapa yang menerimanya. Istilah ini belum ada di glosarium; diminta ke lead (§11).

**Kata pemilik produk** hanya dikutip dari `_progress.md` atau dokumen 17, dengan kedudukannya. Dokumen ini tidak membuka kembali keputusannya; risiko dari keputusannya dinyatakan sebagai risiko sisa yang diterimanya.

---

## 1. Ringkasan: risiko yang dapat menggagalkan tujuan

Tujuan (`CLAUDE.md`): suami dan istri mengetahui keadaan uang dan urusan rumah tangga tanpa saling bertanya, dengan usaha mencatat serendah mungkin, dan aplikasi membantu mereka saling memperhatikan dan bersikap hangat. Empat kelompok risiko dapat menggagalkannya; kelompok kelima hanya berlaku bila kelak ada pengguna luar pertama (keputusan 52) atau pemilik produk mengambil jalur publik opsional (keputusan 51; §0). Semua P1 ada di sini.

| # | Kelompok | Risiko | Mengapa dapat menggagalkan tujuan | Pengaman terkuat yang sudah dirancang |
|---|---|---|---|---|
| 1 | **Masalah inti tidak terjawab** | RISK-18, RISK-19, RISK-20 | Bila hanya satu orang mencatat, atau mencatat terlalu lama, atau catatan terlewat, angka utama tidak menjawab "sisa berapa dan siapa memegang" (KS-1) dan orang kembali bertanya. Tidak ada fitur lain yang menggantikan ini | Data nyata dimulai di klien web yang sama bagi kedua pasangan, dengan pencatat utama di iPhone 13 dan pemilik produk di iPad mini 7 dan Galaxy A55 yang dibawanya ke luar rumah (ADR-030; semula aplikasi iOS, keputusan 42; 13 §3.2); ES §3 tiga ketukan dan simpan lokal seketika; entri terlambat diterima (07 §3.11); rekonsiliasi kas yang menampilkan selisih (07 §3.10) |
| 2 | **Angka uang salah** | RISK-47, RISK-48 | Satu angka yang terbukti salah meruntuhkan kepercayaan pada semua angka (brief §7: "non-negotiable") | Satu implementasi aturan di server (ADR-013); setiap invarian 06 dan 07 dan setiap contoh terhitung sebagai uji yang memblokir merge (ES §5); kunci idempotensi dan satu transaksi `SERIALIZABLE` per perintah; VL-10 nol; bila tidak nol, server masuk *tahan tulis* sementara pencatatan berlanjut di antrean perangkat (13 §6.2; ADR-028 butir 3) |
| 3 | **Aplikasi menjadi alat kendali antaranggota** | RISK-11, RISK-35, RISK-38, RISK-40 (P1); RISK-21, RISK-31, RISK-34, RISK-37, RISK-39 (P2) | Aplikasi yang dipakai satu orang untuk mengawasi atau menekan yang lain merusak "attentive and kind" dan melanggar brief §8. Untuk tahap publik, satu kasus seperti itu cukup untuk merusak kepercayaan pada produk | Pagar 06 PG-01 sampai PG-14 sebagai uji; "tidak pernah terlihat" per peran (SEC-86); uji simetri berpermutasi (SEC-89); keluar dengan aman tanpa persetujuan siapa pun; tidak ada perintah atas orang dewasa lain; halaman Bantuan tanpa jejak |
| 4 | **Operasi bergantung pada satu orang** | RISK-10 (P1) | Pemilik produk adalah pemimpin teknis, pengelola server, dan pemegang kunci cadangan; bila ia berhalangan lama, rumah tangga kehilangan aplikasi yang dipakainya sehari-hari | Runbook di fase bangun, sebelum GR-2; komponen sedikit; salinan kunci cadangan offline; Q-222 diputuskan sebelum GR-2 |
| 5 | **Jalur opsional, bila kelak diambil: pengguna luar pertama atau tahap publik** (keputusan 51, 52) | RISK-93 (pengguna luar pertama, termasuk orang tua dan kerabat dewasa lain yang diberi login); RISK-25, RISK-58 (tahap publik) (P1); RISK-26; RISK-65 kini P2 (S6, jawaban advokat, keputusan 59) | Menurut jawaban advokat (keputusan 59), aplikasi tertutup untuk keluarga, termasuk orang tua yang diberi login, tidak wajib PSE dan tetap dalam pengecualian rumah tangga; kewajiban PSE baru timbul bila dibuka untuk umum. Pengelola server yang anak mereka sendiri dapat membuka data orang tua, sementara orang tua sungkan menolak (RISK-93, Pagar); pilot satu rumah tangga dapat menghasilkan "lanjut" di GR-3 yang tidak berlaku bagi keluarga lain; model bisnis belum punya bukti. Tidak satu pun mengancam tujuan rumah tangga selama tidak ada pengguna luar dan jalur publik tidak diambil | GR-4 mensyaratkan jawaban advokat atas L-01, L-14, L-24 (dijawab lewat konsultasi advokat, keputusan 59; kutipan pasal adalah pernyataan advokat) (Q-209; 14 QG-31 (a)), satu konsultasi advokat ulang sebelum go public, yang menggantikan pemicu usulan 10 §3.4 (pengguna non-keluarga, penagihan, App Store tersembunyi, Play publik) (keputusan 67, diputuskan pemilik produk sendiri), dan pendaftaran PSE (SEC-147) baru sebelum dibuka untuk umum; SEC-159 pengelola tidak membuka data akun orang lain tanpa izin; SEC-36 dengan baris akses darurat di jejak keamanan setiap akun dan SEC-34 versi keluarga pengelola sebelum login pertama orang tua; pertanyaan tentang aplikasi, bukan tentang pasangan; beta tertutup dengan pengelola yang bukan anggota sebagai uji pagar sebenarnya (13 §6.4, §6.5); wawancara dan beta sebelum GR-5 |

**Satu P1 yang diterima pemilik produk dengan syarat**, bukan ditutup: RISK-11 (pengelola server pilot adalah pasangan; keputusan 25, 10-F A; sejak S6 ditambah aturan pengelola server, keputusan 60). RISK-65 (risiko PSE, W-32, S4-4 (iii) A) semula P1 yang diterima untuk pilot; sejak jawaban advokat (keputusan 59) dinilai ulang menjadi P2 (S6). Keduanya dinyatakan di kartunya beserta syaratnya.

**Penyalahgunaan antaranggota dan kebenaran uang** mendapat perhatian khusus karena keduanya fondasi yang brief §1.1 larang dipotong untuk pilot. Dari 94 risiko terbuka, **23 bertanda Pagar** dan **6 berkategori Keuangan**; keduanya punya mitigasi yang berupa uji otomatis, bukan niat (§3.3, §3.4, §6). Hanya satu risiko Pagar (RISK-63, langganan Plus) yang berstatus jalur publik, dan satu (RISK-93, pengelola server yang anak orang tua yang diberi login) berstatus pengguna luar; semua risiko Keuangan berlaku bagi rumah tangga.

### 1.1 Peta prioritas risiko berkedalaman penuh

| D \ K | Rendah | Sedang | Tinggi | Pasti |
|---|---|---|---|---|
| **Kritis** | 05, 42, 44, 46, 65, 95, 96 | 35, 38, 40, 47 | 48 | 11, 93 |
| **Tinggi** | 43, 59, 66, 71 | 04, 12, 14, 16, 21, 24, 31, 34, 37, 39, 41, 45, 49, 60, 63, 67, 68, 92, 94 | 10, 18, 19, 20, 25, 58 | — |
| **Sedang** | 06, 09, 56, 57 | 02, 03, 22, 23, 26, 27, 30, 33, 36, 50, 51, 54, 55, 61, 62, 64, 69, 70, 72, 91 | 01, 52, 53 | 97 |
| **Rendah** | 07 | 08, 13, 17, 29, 32 | — | — |

Angka adalah nomor `RISK-NN`. P1: 13 risiko; P2: 30; P3: 34 (berkedalaman penuh, 77 risiko; sesudah S6, web lebih dulu: RISK-94 sampai RISK-97 baru, semuanya P2, dan RISK-91 turun dari P2 ke P3, §12 butir 15; di S6 RISK-65 turun dari P1 ke P2 dan RISK-66 dari P2 ke P3 atas jawaban advokat, keputusan 59; RISK-59 turun dari P2 ke P3 di S5 fase D; RISK-28 ditutup dan RISK-91 baru oleh keputusan 42; RISK-92 baru oleh keputusan 51, dinilai dengan 18 §5.2.9; RISK-93 baru oleh keputusan 52 dan 10.1 §11.3.1). Risiko Tier 2 dinilai di §4; Tier 3 tidak dinilai sampai diangkat, kecuali RISK-83 (P3, §5, karena brief §14).

**Pembagian menurut jalur (keputusan 51, 52; §0).** Dari 77 risiko berkedalaman penuh yang terbuka, **12 berstatus jalur publik**: P1 RISK-25, RISK-58; P2 RISK-16, RISK-60, RISK-63, RISK-67; P3 RISK-26, RISK-59, RISK-61, RISK-62, RISK-69, RISK-72. **2 berstatus pengguna luar**: P1 RISK-93; P3 RISK-70. **Jalur rumah tangga: 63 risiko, P1 10, P2 26, P3 27** (web lebih dulu: RISK-94 sampai RISK-97 masuk, RISK-91 dari P2 ke P3; S6: RISK-69 pindah dari pengguna luar ke jalur publik, RISK-65 dan RISK-66 turun satu prioritas; keputusan 59) (termasuk RISK-92 baru, yang timbul justru karena jalur rumah tangga tidak berbatas waktu). Beberapa risiko rumah tangga punya bagian yang jatuh tempo sebelum pengguna luar pertama (RISK-06, RISK-11, RISK-12, RISK-40, RISK-44, RISK-65, RISK-68). Peta di atas tetap memuat keduanya, karena K dan D tidak berubah.

---

## 2. Register ringkas

Satu baris per risiko. Kolom "Harus ada pada" adalah fase atau gerbang paling lambat bagi mitigasinya. Status "Diterima" berarti risiko sisanya diterima pemilik produk dengan kedudukan yang disebut di kartunya. Status "jalur publik" berarti risiko itu hanya berlaku bila pemilik produk kelak mengambil jalur publik opsional (keputusan 51; §0). Status "pengguna luar" berarti risiko itu berlaku sejak pengguna luar pertama, termasuk orang tua yang diberi login, walaupun tahap publik tidak diambil (keputusan 52; §0).

| ID | Kategori | Pagar | Judul | K | D | P | Pemilik mitigasi utama | Harus ada pada | Status |
|---|---|---|---|---|---|---|---|---|---|
| RISK-01 | Teknis | | Siklus UI iOS tanpa Mac terlalu lambat | Tinggi | Sedang | P2 | Pemilik produk; `ios-dev` | Jalur iOS native, langkah I1 (14 §3.9); semula GR-1 | Diterima |
| RISK-02 | Teknis | | Pustaka 0.x berubah tidak kompatibel | Sedang | Sedang | P3 | `backend-dev`; `android-dev` | FS-1 | Terbuka |
| RISK-03 | Teknis | | Satu VPS titik gagal tunggal | Sedang | Sedang | P3 | Pengelola server | GR-2 | Terbuka |
| RISK-04 | Hukum | | Letak pusat data dan kebijakan privasi penyedia hosting | Sedang | Tinggi | P2 | Pemilik produk; pengelola server | GR-2 | Terbuka |
| RISK-05 | Teknis | | pgBackRest dan PG18: cadangan tidak andal | Rendah | Kritis | P2 | Pengelola server | GR-2 | Terbuka |
| RISK-06 | Teknis | | Email pemulihan tidak sampai | Rendah | Sedang | P3 | `backend-dev`; pengelola server | GR-2 | Terbuka |
| RISK-07 | Teknis | | Foto di PostgreSQL membesarkan cadangan | Rendah | Rendah | P3 | `backend-dev` | FS-2 | Terbuka |
| RISK-08 | Teknis | | SSE terputus di jaringan seluler | Sedang | Rendah | P3 | `backend-dev` | FS-1 | Terbuka |
| RISK-09 | Keuangan | | Eksponen IDR produk lawan ISO 4217 | Rendah | Sedang | P3 | `backend-dev` | FS-1 | Terbuka |
| RISK-10 | Teknis | | Satu pemegang pengetahuan (bus factor) | Tinggi | Tinggi | P1 | Pemilik produk; `architect` (14) | GR-2 | Terbuka |
| RISK-11 | Keamanan dan privasi | Pagar | Pengelola server pilot adalah pasangan | Pasti | Kritis | P1 | Pengelola server; `security-privacy` | GR-2 | Diterima dengan syarat |
| RISK-12 | Hukum | | Jalur yang bergantung pada tafsir hukum: xtool, transfer push dan `AI` | Sedang | Tinggi | P2 | Advokat; `security-privacy` | GR-4, sebelum pengguna luar pertama | Terbuka |
| RISK-13 | Teknis | | Pembungkus SQLite iOS tidak dapat dibangun di Linux | Sedang | Rendah | P3 | `ios-dev` | Jalur iOS native, langkah I1 | Terbuka |
| RISK-14 | Keamanan dan privasi | Pagar | Biometrik membuka kunci aplikasi bagi pasangan | Sedang | Tinggi | P2 | `ios-dev`; `android-dev` | GR-2 | Terbuka |
| RISK-15 | — | | Pengecualian ES §2 ditolak | — | — | — | — | — | **Ditutup** (§9) |
| RISK-16 | Bisnis | | App Review menolak karena Guideline 1.2 | Sedang | Tinggi | P2 | `ios-dev`; `security-privacy` | GR-4 | Terbuka; jalur publik |
| RISK-17 | Teknis | | API eksperimental Compose berubah | Sedang | Rendah | P3 | `android-dev` | Jalur Android | Terbuka |
| RISK-18 | Produk | | Hanya satu pasangan yang aktif mencatat | Tinggi | Tinggi | P1 | Pemilik produk; `ux-designer`; `web-dev` | GR-2 (klien web di ketiga perangkat) | Terbuka |
| RISK-19 | Produk | | Usaha mencatat terlalu tinggi bagi pencatat yang pelupa | Tinggi | Tinggi | P1 | `ux-designer`; `ios-dev`; `android-dev` | GR-2 | Terbuka |
| RISK-20 | Produk | | Catatan terlewat membuat angka utama tidak dipercaya | Tinggi | Tinggi | P1 | `finance-domain`; `ux-designer` | FS-1 | Terbuka |
| RISK-21 | Produk | Pagar | Persetujuan terasa mengekang atau tidak dipakai | Sedang | Tinggi | P2 | `architect`; `ux-designer` | FS-1 | Terbuka |
| RISK-22 | Produk | | Notifikasi terlalu banyak | Sedang | Sedang | P3 | `architect`; `ux-designer` | FS-1 | Terbuka |
| RISK-23 | Produk | | Cakupan dan navigasi tumbuh | Sedang | Sedang | P3 | Pemilik produk; `product-analyst` | FS-2; GR-3 (pemangkasan cakupan publik), bila jalur publik diambil | Terbuka |
| RISK-24 | Produk | | Penyiapan awal terlalu berat | Sedang | Tinggi | P2 | `ux-designer`; `product-analyst` | GR-2 | Terbuka |
| RISK-25 | Produk | | Bias validasi pilot | Tinggi | Tinggi | P1 | Pemilik produk; `product-analyst` | GR-3 | Terbuka; jalur publik |
| RISK-26 | Produk | | "Pemakaian harian yang bertahan" dinilai dari pernyataan sendiri | Sedang | Sedang | P3 | Pemilik produk | GR-3 | Terbuka; jalur publik |
| RISK-27 | Produk | | Ramadan dan Idulfitri di jendela pilot | Sedang | Sedang | P3 | `product-analyst` | FS-2 | Terbuka |
| RISK-28 | — | | Urutan platform belum diputuskan | — | — | — | — | — | **Ditutup** (§9) |
| RISK-29 | Produk | | Chat rumah tangga tidak dipakai | Sedang | Rendah | P3 | Pemilik produk | FS-2 | Terbuka |
| RISK-30 | UX | | Tampilan terasa templat atau kurang native | Sedang | Sedang | P3 | `ux-designer` | FS-1 | Terbuka |
| RISK-31 | UX | Pagar | Tanda dan laporan dibaca sebagai tuduhan atau papan skor | Sedang | Tinggi | P2 | `ux-designer`; `finance-domain` | FS-1 | Terbuka |
| RISK-32 | UX | | Arti "terkirim" berbeda dari WhatsApp | Sedang | Rendah | P3 | `ux-designer` | FS-1 | Terbuka |
| RISK-33 | UX | | Perubahan desain OS tahunan dan komponen belum stabil | Sedang | Sedang | P3 | `ios-dev`; `android-dev` | FS-1 | Terbuka |
| RISK-34 | Keamanan dan privasi | Pagar | Pengawasan lewat kehadiran dan tanda dibaca | Sedang | Tinggi | P2 | Pemilik produk; `security-privacy` | FS-1 | Diterima |
| RISK-35 | Keamanan dan privasi | Pagar | Tekanan di luar aplikasi | Sedang | Kritis | P1 | `security-privacy`; `architect` | GR-2 | Terbuka |
| RISK-36 | Keamanan dan privasi | Pagar | Satu pihak menahan ambang | Sedang | Sedang | P3 | Pemilik produk | FS-1 | Diterima |
| RISK-37 | Keamanan dan privasi | Pagar | Pelecehan lewat permintaan, pengingat, keberatan, chat | Sedang | Tinggi | P2 | `architect`; `ux-designer` | FS-1 | Terbuka |
| RISK-38 | Keamanan dan privasi | Pagar | Membaca atau menyimpulkan data pribadi pasangan | Sedang | Kritis | P1 | `backend-dev`; `qa` | GR-2 | Terbuka |
| RISK-39 | Keamanan dan privasi | Pagar | Perpisahan: keluar sendiri satu-satunya jalan | Sedang | Tinggi | P2 | `architect`; `finance-domain` | GR-2 | Terbuka |
| RISK-40 | Keamanan dan privasi | Pagar | Akses tidak sah ke akun atau rumah tangga | Sedang | Kritis | P1 | `backend-dev`; `security-privacy` | GR-2 | Terbuka |
| RISK-41 | Keamanan dan privasi | Pagar | Perangkat hilang atau dipegang orang lain | Sedang | Tinggi | P2 | `ios-dev`; `android-dev` | GR-2 | Terbuka |
| RISK-42 | Keamanan dan privasi | | Server atau cadangan dibobol | Rendah | Kritis | P2 | Pengelola server | GR-2 | Terbuka |
| RISK-43 | Keamanan dan privasi | | Data peka keluar lewat log, push, metadata foto | Rendah | Tinggi | P3 | `backend-dev`; `qa` | FS-1 | Terbuka |
| RISK-44 | Keamanan dan privasi | | Kebocoran lintas rumah tangga | Rendah | Kritis | P2 | `backend-dev`; `qa` | GR-2; GR-4, bila pengguna luar membentuk rumah tangga sendiri atau beta diadakan | Terbuka |
| RISK-45 | Keamanan dan privasi | | Penghapusan tidak tuntas | Sedang | Tinggi | P2 | `backend-dev`; pengelola server | GR-2 | Terbuka |
| RISK-46 | Keamanan dan privasi | Pagar | Halaman Bantuan memuat nomor yang salah atau basi | Rendah | Kritis | P2 | `security-privacy`; pengelola server | GR-2 | Terbuka |
| RISK-47 | Keuangan | | Aturan uang salah diimplementasikan | Sedang | Kritis | P1 | `backend-dev`; `qa`; `finance-domain` | FS-1 | Terbuka |
| RISK-48 | Keuangan | | Kiriman ganda, sinkron luring, perintah bersamaan | Tinggi | Kritis | P1 | `backend-dev`; `ios-dev`; `android-dev` | FS-1 | Terbuka |
| RISK-49 | Keuangan | | Pengguna memilih jenis catatan yang salah | Sedang | Tinggi | P2 | `ux-designer`; `finance-domain` | FS-1 | Terbuka |
| RISK-50 | Keuangan | | Celah label penghasilan | Sedang | Sedang | P3 | `finance-domain` | FS-1 | Diterima |
| RISK-51 | Keuangan | Pagar | Tanda dapat dihindari atau timpang | Sedang | Sedang | P3 | `finance-domain`; `architect` | FS-1 | Diterima |
| RISK-52 | Teknis | | Build TestFlight kedaluwarsa; jalur rilis iOS bergantung pada Apple | Tinggi | Sedang | P2 | `ios-dev`; pemilik produk | Jalur iOS native (QG-23 sejak build TestFlight pertama) | Terbuka |
| RISK-53 | Teknis | | Kapasitas: klien web dan dua UI native dengan cakupan Tier 1 untuk tim 1–3 orang | Tinggi | Sedang | P2 | Pemilik produk; `architect` (14) | GR-1 | Terbuka |
| RISK-54 | Teknis | | Sasaran kinerja ES §3 tidak tercapai di perangkat pilot | Sedang | Sedang | P3 | `web-dev`; `ios-dev`; `android-dev`; `backend-dev` | GR-2 (klien web); QG-30 (iOS); QG-30 (Android) | Terbuka |
| RISK-55 | Teknis | | Rekayasa berlebihan | Sedang | Sedang | P3 | `architect`; `code-reviewer` | FS-1 | Terbuka |
| RISK-56 | Teknis | | Kunci tanda tangan APK dan distribusi APK | Rendah | Sedang | P3 | `android-dev`; pemilik produk | Jalur Android A1 | Terbuka |
| RISK-57 | Teknis | | Ketergantungan pada layanan pihak ketiga | Rendah | Sedang | P3 | `backend-dev` | FS-1 | Terbuka |
| RISK-58 | Bisnis | | Kesediaan membayar tidak diketahui | Tinggi | Tinggi | P1 | Pemilik produk; `product-analyst` | GR-5 | Terbuka; jalur publik |
| RISK-59 | Bisnis | | Komisi toko, pembelian dalam aplikasi, hak lintas platform | Rendah | Tinggi | P3 | `researcher`; `architect` | GR-5 | Terbuka; jalur publik |
| RISK-60 | Bisnis | | Keluarga lain berbeda dari rumah tangga pilot | Sedang | Tinggi | P2 | `product-analyst` | GR-5 | Terbuka; jalur publik |
| RISK-61 | Bisnis | | Pembeda belum teruji | Sedang | Sedang | P3 | `product-analyst`; pemilik produk | GR-5 | Terbuka; jalur publik |
| RISK-62 | Bisnis | | Biaya operasi dan waktu pengembang di tahap publik | Sedang | Sedang | P3 | Pemilik produk | GR-5 | Terbuka; jalur publik |
| RISK-63 | Bisnis | Pagar | Langganan Plus menjadi alat kendali | Sedang | Tinggi | P2 | `architect`; `security-privacy` | FS-3 | Terbuka; jalur publik |
| RISK-64 | Bisnis | | Deskriptor rating usia Apple 12 Januari 2027 | Sedang | Sedang | P3 | Pemilik produk | Jalur iOS native, sebelum unggahan pertama | Terbuka |
| RISK-65 | Hukum | | Status regulasi layanan: PSE dan lingkup jasa keuangan | Rendah | Kritis | P2 | Pemilik produk; advokat | GR-4, sebelum pengguna luar pertama; FS-3 sebelum GR-5 bila jalur publik diambil | Diterima untuk pilot (W-32); dinilai ulang S6 atas jawaban advokat (keputusan 59), semula P1 |
| RISK-66 | Hukum | | Kewajiban pengendali selama pilot | Rendah | Tinggi | P3 | Pemilik produk; `security-privacy` | GR-2; dokumen 10 §4.1 butir 4 dan 5 sebelum go public (keputusan 68) | Terbuka |
| RISK-67 | Hukum | | Catatan hanya-menambah lawan hak hapus | Sedang | Tinggi | P2 | Advokat; `finance-domain` | GR-5 | Terbuka; jalur publik |
| RISK-68 | Hukum | Pagar | Anak: indikator Kepmen 142/2026 dan data anak | Sedang | Tinggi | P2 | `security-privacy`; `ux-designer` | GR-2; GR-4, sebelum pengguna luar pertama | Terbuka |
| RISK-69 | Hukum | | Konten pengguna dan permintaan aparat (Permenkominfo 5/2020) | Sedang | Sedang | P3 | Advokat; `security-privacy` | FS-3, sebelum GR-5 bila jalur publik diambil | Terbuka; jalur publik (S6, keputusan 59) |
| RISK-70 | Hukum | | Regulasi dibaca dari salinan teks; Lembaga PDP belum ada | Sedang | Sedang | P3 | `researcher`; pemilik produk | GR-4, sebelum pengguna luar pertama | Terbuka; pengguna luar |
| RISK-71 | Hukum | | Kedudukan catatan dalam perkawinan dan perpisahan | Rendah | Tinggi | P3 | Advokat; `ux-designer` | FS-1; GR-5, bila jalur publik diambil | Terbuka |
| RISK-72 | Hukum | | Kewajiban menjual langganan | Sedang | Sedang | P3 | Advokat | GR-5 | Terbuka; jalur publik |
| RISK-73 sampai RISK-80 | Tier 2 | sebagian | Lihat §4 | | | | | FS-2 | Terbuka |
| RISK-81 sampai RISK-90 | Tier 3 | — | Lihat §5 (RISK-88 **ditutup**, §9) | | | | | FS-4 sampai FS-8 | Terbuka |
| RISK-91 | Teknis | | Android native menyusul: aplikasi Android teruji belakangan | Sedang | Sedang | P3 | Pemilik produk; `android-dev`; `qa` | QG-30 (Android); sebelum jalur pengujian Google Play dibuka (GR-4 butir 8; QG-31 (b)) | Terbuka; dinilai ulang (web lebih dulu) |
| RISK-92 | Teknis | | TestFlight internal sebagai jalur iOS keluarga tanpa batas waktu tidak sesuai aturan Apple | Sedang | Tinggi | P2 | `architect` (ADR-022); `researcher`; pemilik produk | Jalur iOS native: jalur dipilih (keputusan 55, selesai); GR-2 (SEC-158, tetap) | Diterima (keputusan 55) |
| RISK-93 | Keamanan dan privasi | Pagar | Pengelola server adalah anak atau kerabat pengguna luar dari keluarganya sendiri | Pasti | Kritis | P1 | Pengelola server; `security-privacy`; `backend-dev` | GR-4, sebelum login pertama orang tua atau kerabat | Terbuka; pengguna luar |
| RISK-94 | Teknis | | Browser menghapus antrean catatan yang belum terkirim (TM-44) | Sedang | Tinggi | P2 | `web-dev`; `ux-designer`; `qa` | GR-2 (QG-30 butir 16) | Terbuka |
| RISK-95 | Keamanan dan privasi | Pagar | Pengelola menyajikan kode web yang diubah bagi satu akun (TM-42) | Rendah | Kritis | P2 | Pengelola server; `security-privacy`; `web-dev` | GR-2 | Diterima (bawaan 69.7, diterima atas saran lead) |
| RISK-96 | Keamanan dan privasi | | Kode yang berjalan di klien web bukan kode yang dimaksud: dependensi npm disusupi, XSS, service worker versi lama (TM-38, TM-37, TM-45) | Rendah | Kritis | P2 | `web-dev`; `code-reviewer`; pengelola server | FS-1 langkah 1; GR-2 (QG-30 butir 17) | Terbuka |
| RISK-97 | UX | | Klien web tidak terasa native di ponsel; brief §10.1 dipenuhi belakangan oleh aplikasi native | Pasti | Sedang | P2 | Pemilik produk; `ux-designer`; `web-dev` | GR-2 (gerbang contoh layar web 08); jalur iOS native dan Android | Diterima (pemilik produk, web lebih dulu) |

---

## 3. Rincian risiko berkedalaman penuh

### 3.1 Produk

#### RISK-18 — Hanya satu pasangan yang aktif mencatat

| | |
|---|---|
| Penilaian | Produk · K **Tinggi** · D **Tinggi** · **P1** |
| Uraian dan sebab | Nilai produk lahir ketika kedua orang melihat angka yang sama tanpa bertanya (PRN-01; KS-1). Bila hanya satu yang mencatat, custody di tangan pasangannya kosong atau basi, dan KS-2 gagal. Brief §14 menyebut "only one family member actively using the app". Sebabnya: pemrakarsa aplikasi akan memakainya apa pun bentuknya, sedangkan pasangannya belum tentu (A-25; 03 §2.1; dipertanyakan Q-218); pencatatan terasa sebagai tugas satu orang. **Sejak keputusan 42:** sampai klien Android tersedia, pemilik produk mencatat di iPad mini 7, yang bukan perangkat yang dibawa saat berbelanja, sehingga pengeluarannya di luar rumah baru tercatat belakangan dari iPad **[Asumsi]** (13 §3.2 "Risiko sisa"; 17 keputusan 42). Catatan yang tertunda membuat angka utama tertinggal sampai catatan itu masuk, dan pasangannya kembali bertanya. **Dinilai ulang sesudah ADR-030 (web lebih dulu):** sebab khusus keputusan 42 itu hilang, karena pemilik produk mencatat dari Galaxy A55 lewat Chrome sejak GR-2. Sebab utamanya tidak berubah (pemrakarsa lawan pasangan, pencatatan terasa tugas satu orang), jadi K Tinggi dan D Tinggi tetap. Yang baru: di iPhone 13 pencatatan memakai web app layar utama, yang lebih lambat dibuka dan kurang terasa native daripada aplikasi iOS (RISK-97), dan catatan luring dapat hilang bila browser menghapus data situs (RISK-94) |
| Mitigasi | **Sejak ADR-030:** ketiga perangkat pilot memakai klien web yang sama sejak GR-2; uji kecepatan alur inti di ketiga perangkat (14 §3.3 langkah 2); ES §3 diukur di klien web (ADR-029 catatan web); aplikasi iOS native menyusul untuk pencatat utama (14 §3.9). Mitigasi sebelumnya tetap: istri, yang menurut pemilik produk paling sering mencatat ("istri saya yang sering catat", `_progress.md` 12 November 2026), memakai platform pertama, iPhone 13, yang mendapat kehalusan paling banyak (keputusan 42; Q-211 terjawab). Siapa pemrakarsanya tidak ditetapkan di sini (A-25 **[Asumsi]**; Q-218 terbuka); kedua pasangan memakai aplikasi iOS yang sama sejak data nyata (13 §3.2); entri terlambat diterima dengan tanggal kejadiannya (PRN-06; 07 §3.11); pemilik produk "jarang kemana-mana" (17 keputusan 42); bila VL-07 butir (2) menunjukkan bahwa momen yang terlewat terutama pengeluaran pemilik produk di luar rumah, Android dinaikkan urutannya di dalam FS-2 (13 §3.2 **[Usulan]**; 14 §3.8); usaha mencatat rendah (RISK-19); serah terima dikonfirmasi pemegangnya, jadi pencatatan dua arah sejak hari pertama (07 §2.5); tidak ada hitungan per orang yang membuat satu orang tampak malas (06 PG-09, AT-4); kartu objek di chat sehingga tanggapan terjadi di tempat percakapan (ADR-011) |
| Pemilik mitigasi | Pemilik produk (urutan jalur native); `ux-designer`; `web-dev`; `ios-dev`; `android-dev` |
| Harus ada pada | GR-2: klien web terpasang di iPhone 13 dan iPad mini 7 sebagai web app layar utama dan dipakai di Galaxy A55 (14 QG-30 butir 3 dan 16; ADR-030). Aplikasi native di jalurnya (QG-30 (iOS), QG-30 (Android)); sebelum jalur pengujian Google Play dibuka (GR-4 butir 8; QG-31 (b)) bila jalur publik diambil |
| Tanda awal | VL-03 bulan pertama "tidak" dari salah satu; bahan frekuensi VL-02 turun; baris "belum dikonfirmasi" menumpuk di tempat uang satu pemegang (dibaca di tingkat rumah tangga); VL-01 tetap memperlihatkan pertanyaan "sisa berapa" seperti di VL-00 |
| Risiko sisa | K Sedang, D Tinggi. Aplikasi tidak dapat memaksa orang mencatat dan memang tidak boleh menekan (PG-06). Bila VL-03 gagal, 13 §6.2: FS-2 diperpanjang dan model "dua pencatat" ditinjau ulang |
| Sumber | Brief §1.1 KS-1, KS-2; brief §14; A-19, A-25; Q-211, Q-218; keputusan 42; ADR-030; 03 §2.1; 13 §3.2, §6.2, VL-07; SCN-27; RISK-94, RISK-97 |

#### RISK-19 — Usaha mencatat terlalu tinggi bagi pencatat yang pelupa

| | |
|---|---|
| Penilaian | Produk · K **Tinggi** · D **Tinggi** · **P1** |
| Uraian dan sebab | Brief §1.1: "The main recorder is forgetful, so entry must take seconds and tolerate late entry". Usaha berulang Tier 1 diperkirakan 7 sampai 29 tindakan singkat per orang per minggu (A-29 **[Asumsi]**). Bila mencatat lebih lama daripada membayar, catatan terlewat (RISK-20). Sebabnya: banyak jenis peristiwa (belanja, tarik tunai, serah terima, dompet campuran), banyak tempat uang dan kategori, layar isian yang berat, cold start lambat, jaringan buruk |
| Mitigasi | ES §3: paling banyak 3 ketukan sesudah layar isian terbuka, tersimpan lokal seketika, cold start ≤ 2 detik (`[DEFAULT]`, dikukuhkan ADR sesudah diukur); satu tombol Catat (08); isian bawaan dari entri terakhir (IDEA-02, FS-1); pengingat tagihan dan pos rutin (IDEA-03); entri terlambat diterima dengan tanggal kejadiannya (07 §3.11); pencatatan offline-first (ADR-018); uji kecepatan tiga alur inti di perangkat pilot sebelum GR-2 (13 §3.1) |
| Pemilik mitigasi | `ux-designer`; `ios-dev`; `android-dev`; `qa` |
| Harus ada pada | FS-1 (uji kecepatan); GR-2 butir 5 (ES §3 diukur di iPhone 13 dan iPad mini 7); QG-30 (Android) untuk Galaxy A55 (14 §5.3) |
| Tanda awal | VL-07: ketukan dan waktu di atas ES §3, atau momen yang dilewati karena "terlalu lama"; keluhan seperti SCN-27 ("Capek mencatat") |
| Risiko sisa | K Sedang, D Tinggi. Sasaran ES §3 tetap `[DEFAULT]` sampai diukur; pembanding pesaing (R-02) belum diuji |
| Sumber | Brief §1.1, §14; ES §3; A-19, A-29; 04 §6; 13 §3.1, VL-07; SCN-27 |

#### RISK-20 — Catatan terlewat membuat angka utama tidak dipercaya

| | |
|---|---|
| Penilaian | Produk · K **Tinggi** · D **Tinggi** · **P1** |
| Uraian dan sebab | *Uang yang ada* dan custody hanya benar bila catatannya lengkap. Tarik tunai, serah terima, dan kiriman untuk keluarga besar adalah yang paling sering tidak tercatat di rumah tangga pilot (brief §1.1). Sekali angka terbukti meleset dari dompet, orang kembali bertanya, dan KS-1 gagal walaupun semua invarian benar |
| Mitigasi | Rekonsiliasi kas dengan selisih yang tampil, tidak disembunyikan (07 §3.10, INV-16); penanda "belum dikonfirmasi" pada catatan atas uang orang lain (07 §2.5); hitungan bukan pemegang hanya atas tempat uang yang pemegangnya tandai (INV-33); saldo awal tempat uang boleh didaftarkan belakangan dan laporan menyebutnya (07 §3.17, INV-19); VL-01 bulanan memuat menghitung uang bersama dan mencocokkannya (13 §6.1) |
| Pemilik mitigasi | `finance-domain` (aturan); `ux-designer`; pasangan pilot (rekonsiliasi) |
| Harus ada pada | FS-1 (rekonsiliasi kas Tier 1); FS-2 (VL-01) |
| Tanda awal | Selisih kas besar atau berulang di tempat uang yang sama; jawaban VL-01 berbeda antara kedua pasangan; keluhan seperti SCN-28 |
| Risiko sisa | K Sedang, D Sedang. Selisih tetap ada tetapi terlihat; siapa yang lupa diselesaikan di luar aplikasi (keputusan 7; 07 §10.3 butir 3; lihat RISK-31) |
| Sumber | Brief §1.1, KS-1; 07 §3.10, §5; 13 VL-01; SCN-28 |

#### RISK-21 — Persetujuan terasa mengekang atau tidak dipakai

| | |
|---|---|
| Penilaian | Produk · **Pagar** · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Brief §5D: alur persetujuan "must support cooperation, not coercive control or surveillance". Bila permintaan terasa minta izin satu arah, pasangan berhenti mengajukannya (KS-4 gagal) atau merasa dikendalikan. Brief §14 menyebut "intrusive approvals". Sebabnya: ambang dipahami sebagai kunci, pengingat yang berulang, diam dibaca sebagai penolakan, tanda di atas ambang tampil seperti pelanggaran |
| Mitigasi | Dua arah dan simetris (06 PG-01, PG-02); ambang menandai, tidak menghalangi (ADR-009; PG-10; INV-A11); diam bukan keputusan (PG-04); pengingat berbatas (PG-06); diskusi dan usul balik pada permintaan itu sendiri (brief §1.1 Tier 1); persetujuan satu ketukan dari kartu chat (08); belanja biasa dicatat sesudah terjadi tanpa persetujuan (A-20); tanpa "Setujui semua" pada kelompok usulan (D-112) |
| Pemilik mitigasi | `architect` (06); `ux-designer` |
| Harus ada pada | FS-1; dinilai di FS-2 (VL-05, VL-08) |
| Tanda awal | Jumlah permintaan per bulan turun ke nol sementara VL-05 melaporkan kejutan; VL-08 "ditagih" atau "didesak"; keluhan SCN-30, SCN-31 |
| Risiko sisa | K Rendah, D Tinggi. Ambang tetap kesepakatan, bukan kunci (07 §10.3 butir 4), jadi KS-4 bergantung pada kebiasaan. Untuk tahap publik, 13 §6.3 sudah menyebut akibatnya bila keluarga lain menganggap persetujuan mengekang |
| Sumber | Brief §5D, §14; 06 §10; ADR-009; A-20; 13 VL-05, VL-08, §6.3 |

#### RISK-22 — Notifikasi terlalu banyak

| | |
|---|---|
| Penilaian | Produk · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Brief §14 "notification overload"; brief §5N "minimize notification fatigue". Permintaan, pengingat, chat, tagihan, dan (Tier 2) kabar kepedulian dan `AI` dapat menumpuk |
| Mitigasi | Jam tenang per orang (A-28) dan ringkasan (A-32), keduanya Tier 1; plafon enam pemberitahuan "menunggu jawabanmu" per hari dan tiga pengingat ulang per objek (PG-06; W-11); lipat notifikasi chat per pengirim (A-34); tanda dasar jatah dan pos rutin tidak pernah menjadi notifikasi langsung (07 §10.3 butir 10); kabar kepedulian paling banyak berbunyi sekali per 60 menit (D-108) |
| Pemilik mitigasi | `architect` (06 §11); `ux-designer` |
| Harus ada pada | FS-1 |
| Tanda awal | VL-09 "terlalu banyak" |
| Risiko sisa | K Rendah, D Sedang. Dengan banyak permintaan terbuka, pengingat datang lebih sering daripada tiga kali sehari (W-11); angkanya **[Asumsi]** dan disetel di pilot |
| Sumber | Brief §5N, §14; 06 §10, §11; A-28, A-32, A-34; 13 VL-09 |

#### RISK-23 — Cakupan dan navigasi tumbuh

| | |
|---|---|
| Penilaian | Produk · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Brief §14 "too many features", "complex navigation". Tier 1 kini 55 fitur dengan 198 FR Must, dan 18 fitur Tier 2 masuk selama pilot (`_progress.md`, "Putaran pencatatan sesudah lanjut"; 13 §1). Setiap gelombang menambah layar dan kata |
| Mitigasi | Empat tujuan navigasi di Tier 1, ditambah satu di Tier 2 (08); fitur muncul kontekstual lewat kartu objek, bukan menu; Tier 2 dilepas bergelombang (13 §3.4); fitur Tier 2 yang tidak dipakai enam minggu sesudah dilepas menjadi kandidat dikeluarkan (13 §6.2; PRN-02) |
| Pemilik mitigasi | Pemilik produk; `product-analyst`; `ux-designer` |
| Harus ada pada | FS-2 (bergelombang); GR-3 (pemangkasan cakupan publik), bila jalur publik diambil |
| Tanda awal | VL-08 "apa yang ingin kamu ubah" menyebut menu atau mencari; gelombang Tier 2 yang tidak meninggalkan objek apa pun |
| Risiko sisa | K Rendah, D Sedang |
| Sumber | Brief §14; 04; 08; 13 §3.4, §6.2 |

#### RISK-24 — Penyiapan awal terlalu berat

| | |
|---|---|
| Penilaian | Produk · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Brief §14 "difficult onboarding". Data nyata dimulai dengan mendaftar tempat uang, saldo awal, kategori, pos rutin, dan bila disepakati ambang dan jatah. Sesi pertama yang panjang membuat data nyata tertunda, atau dimulai setengah sehingga angka utama salah sejak awal (RISK-20) |
| Mitigasi | Preset aturan sehingga tidak semua diatur di hari pertama (brief §4.3); ambang dan jatah boleh tidak ada (07 INV-40; W-14); saldo awal per tempat uang boleh menyusul (07 §3.17); menyepakati ambang terpisah dari menerima undangan (FR-FAM-004/AC-6); alur pembuatan rumah tangga dan undangan diuji lewat SCN-42 |
| Pemilik mitigasi | `ux-designer`; `product-analyst` |
| Harus ada pada | GR-2 |
| Tanda awal | Tempat uang yang dipakai sehari-hari belum terdaftar sesudah minggu pertama data nyata; VL-07 minggu kedua menyebut penyiapan |
| Risiko sisa | K Rendah, D Sedang |
| Sumber | Brief §4.3, §14; 07 §3.17; FR-FAM-004; SCN-42 |

#### RISK-25 — Bias validasi pilot

| | |
|---|---|
| Penilaian | Produk · K **Tinggi** · D **Tinggi** · **P1** · **Jalur publik** (keputusan 51: berlaku hanya bila pemilik produk kelak mempertimbangkan go public; §0) |
| Uraian dan sebab | 13 §6.5: pilot dijalankan salah satu pasangan, yang juga pengelola server dan pewawancara; tidak ada pihak ketiga. Jawaban jujur pasangan yang lain tentang pagar (VL-08) tidak dapat dijamin. Satu rumah tangga tidak dapat digeneralisasi (PRN-04). Pemrakarsa memakai aplikasi apa pun bentuknya (A-25). Akibatnya tinjauan tiga bulan (bukan gerbang; tempat VL-08 dibacakan bagi rumah tangga) dapat menyatakan semuanya baik, dan GR-3, bila jalur publik diambil, dapat memutuskan "lanjut", atas bukti yang lemah, sehingga masalah pagar baru tampak di tahap publik |
| Mitigasi | Pertanyaan VL-08 tentang aplikasi, bukan tentang pasangan; boleh dijawab tertulis kapan saja (13 §6.1); VL-10 dan VL-02 (b) dibaca dari data, bukan pendapat, sedangkan VL-02 (a) adalah pernyataan sendiri kedua pasangan (keputusan 46) dan memikul bias yang sama (RISK-26), sehingga satu "belum" menahan GR-3 dan jawabannya lewat saluran VL-08; wawancara keluarga lain dengan pasangan diwawancarai terpisah (13 §6.3); beta tertutup sekurangnya 5 rumah tangga tanpa batas atas (sekurangnya 14 penguji Android bila akun Play pribadi; 13 §6.4, 14 §6.7), yang pengelola servernya bukan anggota mereka, sebagai uji pagar yang sebenarnya, dengan VL-08 tanpa temuan menghalangi sebagai syarat GR-5 (13 §6.4). Permintaan tambahan: saluran jawaban VL-08 yang tidak melewati pemilik produk (§11) |
| Pemilik mitigasi | Pemilik produk; `product-analyst` |
| Harus ada pada | GR-2 (VL-00 dicatat); GR-3 dan GR-5 (beta), bila jalur publik diambil. Bagi rumah tangga, VL-08 dibacakan di tinjauan tiga bulan (bukan gerbang; 13 §6.1) |
| Tanda awal | Semua VL terpenuhi sejak bulan pertama tanpa satu keluhan pun; VL-08 tidak pernah dijawab tertulis |
| Risiko sisa | K Sedang, D Tinggi. Tidak dapat dihilangkan dalam pilot satu rumah tangga; karena itu GR-5 bergantung pada beta, bukan pada pilot |
| Sumber | 13 §6.1, §6.3 sampai §6.5; A-17, A-25; PRN-04; Q-218 |

#### RISK-26 — "Pemakaian harian yang bertahan" dinilai dari pernyataan sendiri

| | |
|---|---|
| Penilaian | Produk · K **Sedang** · D **Sedang** · **P3** · **Jalur publik** (keputusan 51; §0) |
| Uraian dan sebab | Brief §1 "First users": "Going public is not planned (decision 51); if the owner later decides on it, only after the pilot shows sustained daily use". Jadi GR-3 sendiri kini opsional: ia dinilai hanya bila pemilik produk kelak mempertimbangkan go public. Risiko semula, yaitu ambang ditetapkan sesudah hasil pilot terlihat, **sudah tertangani**: pemilik produk menjawab Q-213 sebelum GR-2 (keputusan 46, 12 November 2026). Menurut bacaan lead atas keputusan itu (`_progress.md`, "Penutupan pertanyaan terbuka sebelum S6"; 13 §3.5, VL-02), GR-3 paling cepat tiga bulan sesudah GR-2, bila (a) kedua pasangan menyatakan sendiri merasa nyaman dan aman memakai aplikasi, dan (b) dalam 14 hari terakhir tidak ditemukan bug baru yang berat (angka uang salah, data hilang atau bocor, aplikasi tidak dapat dipakai). Ambang 5 dari 7 hari dicabut; frekuensi catatan hanya bahan. Risiko yang tersisa: (a) adalah pernyataan sendiri, dan pemilik produk sekaligus pengelola server, pengembang, dan penanya, sehingga "nyaman" dapat dinyatakan tanpa kebiasaan mencatat yang bertahan, dan tekanan untuk menyatakannya tidak terlihat (RISK-25). Tanpa ambang frekuensi, "lanjut" di GR-3 tidak membuktikan pemakaian harian bagi keluarga lain |
| Mitigasi | Satu "belum" dari pasangan mana pun menahan GR-3, dan tidak ada yang menimpanya (13 §6.2); jawaban (a) per orang, terpisah, tentang aplikasi, lewat saluran VL-08 yang tidak melewati pemilik produk (13 VL-02, §6.5); (b) dibaca dari catatan bug yang hanya memuat tanggal, golongan, dan status perbaikan, dengan golongan yang diperdebatkan dibaca berat (13 VL-02; 14 §5.3 masukan teknis GR-3), bersama VL-10 dari data; bug berat yang ditemukan lebih dari 14 hari sebelum GR-3 tetapi belum diperbaiki juga menahan GR-3 (bacaan lead **[Usulan]** atas `critic` P-4, `_progress.md`; 14 §5.3), sehingga (b) tidak dapat dipenuhi dengan menunggu 14 hari tanpa memperbaiki; frekuensi catatan tetap dibaca sebagai bahan tanpa pemisahan per orang (PRN-11); GR-5 bergantung pada beta dengan pengelola yang bukan anggota, bukan pada pilot (RISK-25) |
| Pemilik mitigasi | Pemilik produk; `product-analyst`; `qa` (catatan bug) |
| Harus ada pada | GR-3, bila jalur publik diambil |
| Tanda awal | Pernyataan "nyaman" dari kedua pasangan sementara bahan frekuensi menurun atau catatan terlambat menumpuk; catatan bug tanpa satu pun golongan berat selama berbulan-bulan sementara VL-04 "angkanya dipercaya?" dijawab ragu |
| Risiko sisa | K Sedang, D Sedang. Tidak dapat dihilangkan dalam pilot satu rumah tangga. Ini akibat definisi keputusan 46, tetapi risiko sisa ini tidak disajikan kepada pemilik produk saat ia menjawab, jadi statusnya tetap Terbuka, bukan Diterima. Bukti pemakaian bagi keluarga lain datang dari beta (13 §6.4) |
| Sumber | Brief §1; keputusan 46, 51; glosarium *Pemakaian harian yang bertahan*; 13 §3.5, §6.2, VL-02; Q-213 (selesai); RISK-25 |

#### RISK-27 — Ramadan dan Idulfitri di jendela pilot

| | |
|---|---|
| Penilaian | Produk · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Bulan mulai pilot belum diketahui (A-37 **[Asumsi]**, bawaannya sesudah pertengahan Januari 2027). Bila jendela tiga bulan FS-2 mencakup Ramadan 2027 (awalnya sekitar sebulan sebelum Idulfitri **[Asumsi]**) dan Idulfitri 1448 H (libur nasional 10 sampai 11 Maret 2027, cuti bersama 9, 12, dan 15 Maret 2027 **[Terverifikasi]**, SKB tiga menteri, 18 §2.14), maka di jendela itu: pengeluaran melonjak, mudik menjauhkan pasangan dari rutinitas dan jaringan, kebiasaan mencatat terganggu. Bila akhir jendela bertepatan dengan mudik, bahan frekuensi VL-02 turun karena kalender, bukan karena produk; sejak keputusan 46 frekuensi bukan syarat GR-3, tetapi dapat disalahbaca |
| Mitigasi | Tidak ada ukuran VL yang membandingkan besar pengeluaran antarbulan (13 §6.1); THR dan angpau sudah Tier 1 (FR-FIN-040/AC-14); IDEA-08 di FS-2 bila jendela pilot mencakup Ramadan 2027; bila tidak, sebelum Ramadan berikutnya di FS-2 yang berlanjut (13 §8, §9); pencatatan luring selama perjalanan (ADR-018); FS-2 berlanjut tanpa batas akhir sesudah tinjauan tiga bulan, jadi kalender tidak memotong pemakaian (13 §3.4) |
| Pemilik mitigasi | `product-analyst`; pemilik produk |
| Harus ada pada | FS-2 (sebelum dijadwalkan; tanggal resmi R-18) |
| Tanda awal | Tanggal GR-2 ditetapkan dan jendela dihitung |
| Risiko sisa | K Rendah, D Sedang. Tinjauan tiga bulan membaca bahan frekuensi catatan (tingkat rumah tangga, PRN-11) dengan konteks kalender; VL-02 hanya di jalur publik opsional |
| Sumber | A-37, A-40; R-18; 13 §6.1, §8; 11 IDEA-08 |

#### RISK-28 — Urutan platform belum diputuskan

| | |
|---|---|
| Penilaian | **Ditutup** 12 November 2026 (§9). Penilaian sebelumnya: Produk · K Sedang · D Tinggi · P2. Pengenal tidak dipakai ulang |
| Dasar penutupan | Keputusan 42 (17): iOS dan iPadOS dibangun lebih dulu (menjawab Q-37), pilot dengan data nyata di iPhone 13 istri dan iPad mini 7 pemilik produk, Android sesudah GR-2 paling lambat sebelum GR-4 (semula bawaan lead; diterima atas saran lead lewat keputusan 50; sejak keputusan 51, tenggatnya menjadi "sebelum jalur pengujian Google Play dibuka (GR-4 butir 8)"). Siapa memakai perangkat mana kini diketahui (Q-211 terjawab). Urutan kerjanya di 14 §3.4 dan §3.8 |
| Harga pilihan itu, dinilai di risiko lain | RISK-01 (siklus UI iOS lambat sejak awal FS-1, diterima lewat keputusan 43; 14 §9); RISK-18 (belanja pemilik produk di luar rumah tercatat belakangan dari iPad sampai Android tersedia); RISK-91 (pengalaman Android teruji belakangan, dan dekat beta bila beta kelak diadakan) |
| Sumber | Keputusan 42; 17 K-66; Q-37, Q-211; 13 §3.2, §3.4; 14 §3.4, §3.8 |

#### RISK-29 — Chat rumah tangga tidak dipakai

| | |
|---|---|
| Penilaian | Produk · K **Sedang** · D **Rendah** · **P3** |
| Uraian dan sebab | Chat adalah bagian termahal untuk dibangun (W-22). Bila pasangan tetap bercakap di WhatsApp, chat menjadi kotak kartu yang sepi |
| Mitigasi | Beranda yang menjawab "tanpa bertanya" tidak bergantung pada chat (W-22); kartu objek tetap dapat ditanggapi dari notifikasi; tidak ada yang rusak bila chat sepi |
| Pemilik mitigasi | Pemilik produk |
| Harus ada pada | FS-2 (dinilai pemilik produk selama rumah tangga memakai aplikasi; juga di GR-3 bila jalur publik diambil) |
| Tanda awal | Chat hanya berisi kartu sesudah bulan pertama |
| Risiko sisa | K Sedang, D Rendah: biaya bangun sudah dikeluarkan |
| Sumber | W-22; ADR-011; 04 §4.1 |

### 3.2 UX

#### RISK-30 — Tampilan terasa templat atau kurang native

| | |
|---|---|
| Penilaian | UX · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Brief §10.1 menjadikan tampilan native syarat keras, dan pemilik produk pernah menanggapi contoh yang terasa standar tanpa identitas (dicatat lead di `_progress.md`, "Riset tampilan"). Aplikasi yang terasa templat menurunkan kesediaan pasangan memakainya setiap hari |
| Mitigasi | Komponen sistem stabil per platform (ADR-012); gerbang contoh layar arah C kepada pemilik produk dan istrinya sebelum token dibekukan (08 §19.8; "pasang semua", diterima atas saran lead); pemeriksaan anti-templat 08 §19.8 |
| Pemilik mitigasi | `ux-designer`; `ios-dev`; `android-dev` |
| Harus ada pada | FS-1, sebelum token dibekukan |
| Tanda awal | Gerbang 08 §19.8 tidak lulus; VL-08 "ingin diubah" menyebut tampilan |
| Risiko sisa | K Rendah, D Sedang. Selera dua orang bukan riset pasar (08 §24). **Sejak ADR-030** risiko ini berlaku untuk aplikasi native di jalurnya dan untuk sistem desain web (tampilan templat); bahwa web memang tidak native dinilai terpisah di RISK-97, karena keadaannya pasti dan diterima pemilik produk |
| Sumber | Brief §10.1; 08 §19.8; ADR-012; ADR-030; RISK-97 |

#### RISK-97 — Klien web tidak terasa native di ponsel

| | |
|---|---|
| Penilaian | UX · K **Pasti** · D **Sedang** · **P2** · **Diterima pemilik produk** (baru, web lebih dulu) |
| Uraian dan sebab | Brief §10.1 menjadikan tampilan native syarat keras. Klien web memakai komponen yang digambar sendiri; ia dapat menyamai IA, token arah C, dan kata-kata 08, tetapi tidak komponen sistem iOS atau Android, gestur sistem, dan perubahan desain OS tahunan (18 §5.13; 08 bagian web). Di iPhone 13 dan iPad mini 7 aplikasi berjalan sebagai web app layar utama; di Galaxy A55 lewat Chrome. Keadaan ini pasti selama web menjadi klien utama di ponsel. Akibatnya kesediaan memakai setiap hari dapat turun (RISK-18, RISK-19), dan VL-08 dapat menyebut tampilan |
| Mitigasi | Pemilik produk menerimanya: "ya saya terima kan ini versi webnya, selebihnya saya ikut saran kamu" (`_progress.md`, "Jawaban pemilik produk", 12 November 2026), dengan bacaan lead yang dikukuhkannya: syarat brief §10.1 dipenuhi aplikasi native kemudian, bukan oleh web. Yang meringankan: sistem desain web 08 mengikuti konvensi tiap platform sejauh web dapat (navigasi, ukuran sentuh, tipografi sistem, tema gelap dan terang) dan diuji di gerbang contoh layar web sebelum token dibekukan; aksesibilitas penuh (14 §4.9); jalur iOS native (14 §3.9) dan Android (14 §3.8) direncanakan dengan stack native ADR-012 yang tidak dicabut; setiap aplikasi native menggantikan web di perangkatnya sesudah lulus gerbangnya |
| Pemilik mitigasi | Pemilik produk (kapan jalur native dimulai); `ux-designer`; `web-dev` |
| Harus ada pada | GR-2: gerbang contoh layar web 08 lulus; jalur iOS native dan Android sebagai penutup |
| Tanda awal | VL-08 "ingin diubah" menyebut rasa aplikasi atau kecepatan buka; salah satu pasangan kembali ke catatan di luar aplikasi |
| Risiko sisa | K Pasti, D Sedang sampai aplikasi native menggantikan web di perangkat itu; sesudahnya hilang bagi perangkat itu. **Diterima pemilik produk.** Kapan jalur iOS native dimulai diputuskan pemilik produk di GR-2 (14 §3.9) |
| Sumber | Brief §10.1; ADR-012, ADR-030; 08 (bagian web); 18 §5.13; `_progress.md` "Jawaban pemilik produk"; RISK-18, RISK-30 |

#### RISK-31 — Tanda dan laporan dibaca sebagai tuduhan atau papan skor

| | |
|---|---|
| Penilaian | UX · **Pagar** · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Tanda "belum dikonfirmasi", selisih kas, dan tanda di atas ambang hampir selalu jatuh pada pemegang uang belanja (W-12; 07 §10.3 butir 3). Laporan "siapa membayar apa" mengukur lewat tangan siapa uang dibayarkan, bukan siapa yang menanggung (W-6; 07 §10.3 butir 2). Jumlah tanda bukan ukuran besar kelebihan (07 §10.3 butir 7, 10), dan pemilik jatah lebih sering ditandai daripada yang tidak berjatah (butir 11; W-17). Hitungan oleh bukan pemegang mencatat "uang di tanganmu kurang" atas nama orang lain (W-26; TM-29). Dibaca sebagai tuduhan, semua ini merusak "attentive and kind" |
| Mitigasi | Tidak ada papan skor dan tidak ada hitungan per orang (06 PG-09, AT-4; SEC-07); tanda dasar jatah dan pos rutin tidak pernah notifikasi langsung (07 §10.3 butir 10); kata layar di samping istilah domain yang tidak menuduh (07 §1.4; 08); syarat layar 07 §10.3 butir 2, 7, 11 bagi 05 dan 08; hitungan bukan pemegang hanya atas tempat uang yang pemegangnya tandai dan bertanda "belum dikonfirmasi" (INV-33; SEC-91); keberatan selalu tersedia (06 §7) |
| Pemilik mitigasi | `ux-designer`; `finance-domain`; `architect` |
| Harus ada pada | FS-1; dinilai di FS-2 (VL-08) |
| Tanda awal | Keluhan SCN-26 ("Aku terlalu sering ditandai") atau SCN-33 ("Semua aku yang bayar"); VL-08 |
| Risiko sisa | K Sedang, D Sedang. Model tidak dapat menghindari bahwa selisih melekat pada pemegang (07 §10.3 butir 3) |
| Sumber | W-3, W-6, W-12, W-17, W-26; 07 §10.3 butir 2, 3, 7, 10, 11; TM-29; SCN-26, SCN-33 |

#### RISK-32 — Arti "terkirim" berbeda dari WhatsApp

| | |
|---|---|
| Penilaian | UX · K **Sedang** · D **Rendah** · **P3** |
| Uraian dan sebab | "Terkirim" naik hanya saat aplikasi penerima dibuka, bukan saat pesan sampai di ponsel (W-27; ADR-011 butir 5). Pengguna yang terbiasa WhatsApp dapat mengira pesan bermasalah, atau membacanya sebagai tanda pasangan membuka aplikasi |
| Mitigasi | Batas ini dinyatakan kepada pemilik produk sebagai W-27; penjelasan "siapa melihat apa" memuat tanda dibaca dan kehadiran (SEC-87 (c)) |
| Pemilik mitigasi | `ux-designer` |
| Harus ada pada | FS-1 |
| Tanda awal | Pertanyaan pasangan pilot di minggu pertama |
| Risiko sisa | K Rendah, D Rendah. Imbalannya: tanda itu tidak membocorkan apa pun saat ponsel tidak dipakai |
| Sumber | W-27; ADR-011; 06 §16.5 |

#### RISK-33 — Perubahan desain OS tahunan dan komponen yang belum stabil

| | |
|---|---|
| Penilaian | UX · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Brief §9.1 menuntut tampilan yang mengikuti perubahan desain OS setiap tahun, di dua aplikasi native. Komponen Material 3 Expressive hanya ada di jalur 1.5.0 yang belum stabil (08 §17.1; 09 §9: 1.5.0-rc01, 4 November 2026 **[Terverifikasi]**); iOS minimum 26 (ADR-012) |
| Mitigasi | Hanya komponen sistem dan API stabil (ADR-012; ES §8; ADR-023); pemeriksaan pembaruan bulanan (ES §8); padanan stabil per komponen dan penggantian sesudah 1.5.0 stabil (08 §17.4; 09 §6) |
| Pemilik mitigasi | `ios-dev`; `android-dev`; `ux-designer` |
| Harus ada pada | FS-1, lalu setiap rilis OS besar |
| Tanda awal | Rilis beta OS tahunan; catatan rilis Compose Material 3 |
| Risiko sisa | K Sedang, D Rendah. Dua perubahan desain per tahun adalah harga keputusan UI native per platform |
| Sumber | Brief §9.1, §10.1; ADR-012, ADR-023; 08 §17; RISK-17 |

### 3.3 Keamanan dan privasi

#### 3.3.1 Penyalahgunaan antaranggota

##### RISK-11 — Pengelola server pilot adalah pasangan

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Pasti** · D **Kritis** · **P1** · **Diterima pemilik produk dengan syarat** |
| Uraian dan sebab | Selama pilot pemilik produk mengoperasikan server dan basis data sekaligus anggota rumah tangga (keputusan 25; PL-05). Secara teknis ia dapat membaca buku pribadi dan chat istrinya (TM-12), mengubah basis data langsung (TM-31), dan membaca sebutan yang diberikan istrinya untuknya (TM-34). Wewenang ini satu arah (10.3 §2.3). Larangan V-4, V-5, PG-03 mengikat peran di dalam aplikasi, bukan akses server (06 §9) |
| Mitigasi | Tanpa API admin; operasi lewat subperintah di SSH yang tercatat di `security_event` (SEC-33); pemberitahuan pengelola server yang dibaca istri sendiri di aplikasi, dengan pengakuan bertanggal dan berversi, tanpa kalimat janji (SEC-34; 10-F A). Setiap versi baru dan setiap pergantian versi menahan tulisan biasa istri sampai ia mengakui versi terakhir di LY-05, dan menulis satu baris di jejak keamanannya; tindakan pelindung tidak pernah ditahan (10.1 SEC-34 aturan (5); 09.3 §1); status kepedulian yang tidak dibagikan tidak pernah meninggalkan perangkat (10-D A); di pilot "Laporkan" tidak mengirim apa pun dan kontak pengelola tidak tampil (SEC-154; 06 §16.12); beta tertutup, bila kelak diadakan, memakai pengelola yang bukan anggota (13 §6.4); ADR E2EE sebelum tahap publik, bila diambil (10-D A). **Sejak S6:** aturan pengelola server SEC-159, "admin tidak membuka data akun orang lain tanpa izin", diterima atas saran lead (keputusan 60; 10.1 §11.2) atas jawaban advokat untuk L-24 (keputusan 59: pengecualian rumah tangga tidak gugur selama istri tahu dan setuju cara kerjanya; yang dilarang membuka data akun istri diam-diam, menurut advokat UU ITE Pasal 30). Aturan itu janji, bukan pagar teknis; K dan D tidak berubah |
| Pemilik mitigasi | Pengelola server; `security-privacy` |
| Harus ada pada | GR-2 (10 §4.1 butir 2; SEC-107 salinan swalayan, dibangun di FS-1 lewat FR-FAM-013/AC-12, sehingga operator tidak menyusun salinan dari basis data); GR-4, sebelum pengguna luar pertama termasuk orang tua yang diberi login (keputusan 52; pengelola bukan anggota rumah tangga beta bila beta diadakan; SEC-36 akses darurat dengan log hanya-menambah di luar VPS dan satu baris tanpa isi di jejak keamanan setiap akun, SEC-39 (c) publik, SEC-34 tiga versi per akun, menurut 10 §4.2; risiko bagi orang tua sendiri di RISK-93) |
| Tanda awal | Sulit diamati: pengelola yang sama mengendalikan log dan basis data (TM-31). VL-08 adalah satu-satunya tanda dari sisi istri |
| Risiko sisa | Bentuknya tetap Pasti dan Kritis selama pilot, **diterima pemilik produk** (keputusan 25; TM-12 "Diterima untuk pilot"). Arsitektur tidak dapat menutupnya tanpa E2EE, yang tidak dipakai di pilot (S4-6 (i) A). Penerimaan ini berlaku untuk pemakaian rumah tangga tanpa batas waktu (10 §3.3): menurut penilaian `security-privacy`, keputusan 25 tetap mencakupnya sesudah keputusan 51, dan teks pemberitahuan pengelola tidak berbatas waktu (ketetapan lead atas §11 (h)). Bila orang tua diberi login (keputusan 52), pengelola yang sama dapat membaca data mereka; syarat sebelum pengguna luar pertama di atas berlaku, dan risikonya bagi mereka dinilai tersendiri di RISK-93. **Tahan berulang lewat versi pemberitahuan** (SEC-34 aturan (5)): pengelola dapat menerbitkan versi berulang kali untuk menahan pencatatan istrinya. Batasnya:<br>- tidak ada data yang hilang, karena catatan tetap di antrean;<br>- satu kali baca dan satu ketukan mengakhirinya, kapan saja istri mau;<br>- tindakan pelindung tidak pernah ditahan;<br>- setiap versi terlihat di Riwayat keamanan.<br><br>Pengelola juga dapat menulis langsung ke basis data. Selama server hanya dipakai pasangan, keduanya termasuk penerimaan keputusan 25 di atas. Sejak pengguna luar pertama, akses basis data hanya lewat akses darurat SEC-36 yang tercatat di jejak keamanan setiap akun |
| Sumber | 09.4 §9; 10.1 §11; TM-12, TM-31, TM-34; L-24 (dijawab advokat, keputusan 59); keputusan 25, 60; 10 §3 10-D, 10-F, 10-G |

##### RISK-93 — Pengelola server adalah anak atau kerabat pengguna luar dari keluarganya sendiri

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Pasti** · D **Kritis** · **P1** · **Pengguna luar** (keputusan 52; cermin RISK-11) |
| Uraian dan sebab | Bila orang tua atau anggota keluarga lain di luar pasangan diberi login (keputusan 52; kerabat dewasa lain sama dengan orang tua, keputusan 63), pengelola server, yaitu pemilik produk, adalah anak atau kerabat mereka. Secara teknis ia dapat membaca data yang mereka simpan di server, termasuk dana pribadi, kabar sakit yang dibagikan, serta chat beserta foto (TM-12), dan dapat mengubah basis data langsung (TM-31). Ada tiga perbedaan dari RISK-11 (10.1 §11.3.1):<br>- keputusan 25 tidak mengikat mereka, dan versi pilot pemberitahuan tidak berlaku bagi mereka;<br>- anak dewasa yang dapat membuka data keuangan dan kesehatan orang tuanya adalah hubungan kuasa tersendiri;<br>- orang tua dapat sungkan menolak anaknya, sehingga "jangan mengisi data pribadi" lebih sulit dipilih.<br><br>Wewenangnya satu arah. Bila mereka membuat rumah tangga sendiri, pengelola bahkan bukan anggota rumah tangga itu |
| Mitigasi | Sebelum login pertama mereka (14 QG-31 (a); 10 §4.2):<br>- **SEC-36**: tidak ada akses basis data tetap bagi seluruh basis data produksi; akses darurat per kejadian dengan alasan tertulis; log hanya-menambah dikirim keluar dari VPS. Di server keluarga pemilik produk, setiap akses darurat juga menulis **satu baris tanpa isi data** (tanggal, lama, alasan) di jejak keamanan setiap akun, dan baris itu hanya terbaca pemilik akunnya;<br>- **SEC-34 versi keluarga pengelola**, dipilih per akun sebagai setelan server: menyebut pengelola sebagai anggota keluarga, cara akses darurat, dan bahwa pencatatan tidak mencegah semuanya. Pengakuannya bertanggal dan berversi oleh orang itu sendiri; versi yang salah terlihat oleh pembacanya di "Siapa melihat apa" (FR-FAM-019/AC-9), dan **setiap pergantian versi sebuah akun menulis satu baris di jejak keamanan akun itu** (SEC-34; 10 §3.3), sehingga pengelola tidak dapat diam-diam memindahkan orang tua ke versi lain. Setiap versi baru dan setiap pergantian versi menahan tulisan biasa orang tua sampai mereka mengakui versi terakhir di LY-05; tindakan pelindung tidak pernah ditahan (10.1 SEC-34 aturan (5); 09.3 §1);<br>- status kepedulian yang tidak dibagikan tidak pernah meninggalkan perangkat (10-D A);<br>- SEC-109 (b) catatan persetujuan uji per orang, dan jawaban advokat L-01, L-14, L-24 (Q-209). SEC-139 versi pengguna luar dan SEC-150 (a) tinjauan ulang DPIA tidak lagi di sini (keputusan 68; ditunda ke go public); SEC-34 versi keluarga pengelola tetap memberi tahu mereka.<br><br>Bila mereka membuat rumah tangga sendiri: QG-31 (a+) (SEC-154 (a), (b), (d) sampai (f); SEC-134) |
| Pemilik mitigasi | Pengelola server; `security-privacy`; `backend-dev` (baris SEC-36 di jejak keamanan; setelan versi per akun); pemilik produk (menetapkan versi saat memberi akses) |
| Harus ada pada | GR-4 (a), sebelum login pertama orang tua atau anggota keluarga lain di luar pasangan (keputusan 52; 14 QG-31 (a)) |
| Tanda awal | Baris akses darurat atau baris pergantian versi pemberitahuan di jejak keamanan akun orang tua; orang tua tidak mengisi data pribadi atau berhenti memakai sesudah membaca pemberitahuan; versi pemberitahuan yang salah terlihat oleh pembacanya |
| Risiko sisa | Pengelola tetap menguasai server dan kunci cadangan, jadi ia tetap dapat membuka data tanpa tercatat (10.1 §11.3.1). Ia juga dapat menerbitkan versi pemberitahuan berulang kali untuk menahan pencatatan orang tua (SEC-34 aturan (5)). Batasnya: tidak ada data yang hilang; satu kali baca dan satu ketukan mengakhirinya kapan saja; tindakan pelindung tidak pernah ditahan; setiap versi terlihat di Riwayat keamanan mereka. Akses basis data sejak login pertama orang tua hanya lewat akses darurat SEC-36 yang tercatat di jejak keamanan akun mereka. Arsitektur tidak dapat menutupnya tanpa E2EE, yang tidak dipakai (S4-6 (i) A). Akses lewat jalur resmi selalu terlihat oleh pemilik akun; yang tersisa adalah akses di luar jalur itu: K Sedang · D Kritis **[Usulan]**. **Pemilik produk tidak dapat menerimanya atas nama orang tua**, karena keputusan 25 tidak mengikat mereka. Dasarnya hanya pembacaan dan pengakuan mereka sendiri atas versi keluarga pengelola (10.1 §11.3.1). Status tetap Terbuka |
| Sumber | Keputusan 25, 52; 10 §3.3 (RISK-93 dikukuhkan `security-privacy`), §4.2; 10.1 SEC-34, SEC-36, §11.3.1; TM-12, TM-31; Q-209; RISK-11; 14 QG-31 (a) |

##### RISK-95 — Pengelola menyajikan kode web yang diubah bagi satu akun

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Rendah** · D **Kritis** · **P2** (baru, web lebih dulu; cermin RISK-11 dan RISK-93) |
| Uraian dan sebab | Di klien web, kode yang memegang data perangkat dimuat dari server setiap pembaruan. Pengelola server dapat menyajikan bundel yang diubah, juga hanya bagi satu akun, tanpa jejak pemasangan seperti build TestFlight atau APK (TM-42; 10.5 §0). Bundel itu dapat mengirim apa yang sengaja tidak pernah sampai ke server: status `CARE` yang tidak dibagikan (10-D A), kode kunci aplikasi web, atau isi buku pribadi yang sudah di perangkat. Ini menambah satu jalan pada wewenang RISK-11: di aplikasi native, data yang hanya di perangkat di luar jangkauan pengelola; di web tidak (10 §8 butir 17). K Rendah karena memerlukan tindakan aktif dan disengaja, bukan keadaan; D Kritis karena yang terbuka justru data yang dijanjikan tidak meninggalkan perangkat |
| Mitigasi | Bundel dibangun di CI dari repositori, hash setiap berkas dicatat sebagai artefak, dan hash berkas yang disajikan VPS dibandingkan dengan artefak sesudah setiap deploy (SEC-169 (a); 14 §7.1); kalimat [Klien web] di pemberitahuan pengelola yang dibaca istri sendiri menyatakan keadaan ini (SEC-178 (b); SEC-34); aturan pengelola SEC-159 (keputusan 60); versi dan hash bundel yang berjalan tampil di LY-95 "Perangkat ini" (08 §25.6; SEC-169 (f), dianjurkan), sebagai alat periksa, bukan bukti, karena bundel yang diubah dapat menampilkan apa saja. ADR E2EE sebelum tahap publik harus menilai bahwa klien web tidak dapat memberi perlindungan berarti terhadap pengelola, dan perlindungan itu hanya mungkin di aplikasi native (10 §8 butir 17). Penutup sebenarnya bagi perangkat itu: aplikasi native di jalurnya (14 §3.8, §3.9) |
| Pemilik mitigasi | Pengelola server; `security-privacy`; `web-dev`; `backend-dev` |
| Harus ada pada | GR-2 (SEC-169, SEC-178 (b) di teks pemberitahuan; 14 QG-30 butir 17); GR-4 (a) bila pengguna luar memakai klien web, dengan teks versi keluarga pengelola (RISK-93) |
| Tanda awal | Sulit diamati, sama dengan RISK-11: pengelola yang sama mengendalikan CI, VPS, dan pemeriksaan hash. VL-08 adalah satu-satunya tanda dari sisi istri |
| Risiko sisa | K Rendah, D Kritis. Pemeriksaan hash dijalankan pengelola sendiri, jadi ia mencegah salah deploy, bukan pengelola yang berniat; hash di LY-95 "Perangkat ini" dapat dipalsukan oleh bundel yang diubah (SEC-169 (f)). Di server pasangan saja, keadaan ini termasuk lingkup keputusan 25 menurut bacaan `security-privacy` (10 §8 butir 17), tetapi keputusan 25 diambil sebelum web ada, jadi penerimaannya dicatat tersendiri: **Diterima** di bawah keputusan 25 dengan kalimat [Klien web] di pemberitahuan pengelola (bawaan 69.7, diterima atas saran lead; 17). **Pagar penerimaan:** batasnya dinyatakan kepada istri sebelum data nyata (10 §4.1 butir 2), dan SEC-159 tetap. K, D, dan P2 tidak berubah karena penerimaan. Bagi orang tua yang memakai web berlaku RISK-93: pemilik produk tidak dapat menerimanya atas nama mereka |
| Sumber | TM-42; 10.5 §0, SEC-169, SEC-178; 10 §8 butir 17; 10-D A; keputusan 25, 60; ADR-030; RISK-11, RISK-93 |

##### RISK-14 — Biometrik membuka kunci aplikasi bagi pasangan

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Biometrik perangkat menerima siapa pun yang terdaftar dan tidak memberi tahu aplikasi siapa yang cocok (PL-03, **[Terverifikasi]** 18 §2.12.2). Bila sidik jari atau wajah pasangan terdaftar di ponsel, kunci aplikasi tidak menahannya |
| Mitigasi | Kode aplikasi 6 digit milik aplikasi; biometrik gugur bila daftar biometrik perangkat berubah; salah 5 kali berarti kredensial akun tanpa menghapus data (SEC-19; D-107); kredensial akun untuk tindakan yang tidak dapat ditarik (SEC-13) |
| Pemilik mitigasi | `ios-dev`; `android-dev`; `security-privacy` (kata layar SEC-14) |
| Harus ada pada | GR-2 |
| Tanda awal | — (tidak teramati aplikasi) |
| Risiko sisa | K Rendah, D Tinggi. Biometrik pasangan yang sudah terdaftar sebelum kunci dinyalakan tetap diterima; jalan keluarnya memakai kode saja. **Klien web (ADR-030):** kunci aplikasi web di pilot memakai kode saja, diturunkan dengan PBKDF2, tanpa biometrik (SEC-174; 10 §8 butir 20), jadi sebab risiko ini tidak ada di web; ia kembali di aplikasi native (QG-30 (iOS) I-5; QG-30 (Android) A-6) |
| Sumber | 09.4 §9; 10.1 §8; SEC-14, SEC-19, SEC-174; TM-10 |

##### RISK-34 — Pengawasan lewat kehadiran dan tanda dibaca

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Sedang** · D **Tinggi** · **P2** · **Diterima pemilik produk** |
| Uraian dan sebab | Online, mengetik, dan terakhir aktif menit tepat, digerakkan pemakaian di bagian mana pun aplikasi dan tampil di layar chat (10-A A, diputuskan pemilik produk sendiri, berlawanan dengan saran lead dan `security-privacy`; UX-5 A). Itu memperlihatkan kapan pasangan memegang aplikasi, juga saat jam tenang, saat menunda menjawab, atau saat sakit (W-18; TM-08). Kartu permintaan dan kartu status juga bertanda dibaca (W-19, W-20) |
| Mitigasi | Hanya di layar chat, tanpa riwayat, tidak dibaca aturan, pengingat, ringkasan, `AI`, atau ekspor (06 INV-A45, PG-07; SEC-61); tidak ada "dibaca tapi belum dijawab" (PG-07); dinyatakan di "siapa melihat apa" (SEC-87 (c)) |
| Pemilik mitigasi | Pemilik produk (keputusan); `security-privacy` (batas); `backend-dev` (INV-A45) |
| Harus ada pada | FS-1; dinilai di FS-2 (VL-08) |
| Tanda awal | VL-08 "merasa diawasi" (SCN-32), yang menurut 13 §6.1 adalah temuan menghalangi untuk GR-3 |
| Risiko sisa | K Sedang, D Tinggi, **diterima pemilik produk** (10 §3 10-A; W-18). Pilihan P1 dan P2 tidak diterapkan; keputusan ini dapat dibuka pemilik produk di checkpoint mana pun, dan untuk tahap publik ia berlaku bagi setiap rumah tangga |
| Sumber | 10.2 §6; 10 §3 10-A; TM-08; W-18, W-19, W-20; 06 §16.6; SCN-32 |

##### RISK-35 — Tekanan di luar aplikasi

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Sedang** · D **Kritis** · **P1** |
| Uraian dan sebab | Seorang anggota dewasa memaksa pasangannya menyetujui, membagikan status kepedulian, membuka ponsel, tidak keluar, atau memakai ponselnya yang terbuka untuk bertindak atas namanya (TM-09, TM-10; 10.3 §2.7). Kemungkinan di rumah tangga pilot dinilai rendah **[Asumsi]**; di tahap publik, dengan banyak rumah tangga, sedang. Rancangan dibuat untuk tahap publik (brief §1.1) |
| Mitigasi | Pagar 06 PG-01 sampai PG-14 sebagai uji; tidak ada persetujuan atas nama orang lain (PG-02) dan tidak ada perintah atas orang dewasa lain (PG-13; INV-A48); diam bukan persetujuan (PG-04); berbagi dapat dicabut seketika dan senyap (V-7); status yang tidak diisi tidak dapat dibedakan dari yang tidak dibagikan (FR-CARE-001/AC-3); keluar dan batal keluar memerlukan kredensial akun (SEC-13); anggota yang "tidak dapat dihubungi" tidak dapat diambil alih, hanya bukti kematian resmi (SEC-106); halaman Bantuan tanpa jejak (SEC-134); sembunyikan untuk saya (SEC-153); pemberitahuan keluar hanya lewat ringkasan (SEC-56) |
| Pemilik mitigasi | `security-privacy`; `architect` (06); `ux-designer` |
| Harus ada pada | GR-2 |
| Tanda awal | Tidak ada, menurut rancangan: aplikasi tidak memantau perilaku. Di pilot VL-08; di tahap publik laporan ke pengelola yang bukan anggota (SEC-154) |
| Risiko sisa | K Sedang, D Kritis untuk paksaan di luar aplikasi, yang tidak dapat dicegah aplikasi (10.3 §2.7 kolom kanan). Yang dapat dijamin: aplikasi tidak menambah alat tekanan |
| Sumber | Brief §5D, §8; 10.3 §2, §2.7; TM-09, TM-10; 06 §10 |

##### RISK-36 — Satu pihak menahan ambang

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Sedang** · D **Sedang** · **P3** · **Diterima pemilik produk** |
| Uraian dan sebab | Ambang dibuat, diubah, dan diakhiri hanya dengan persetujuan semua pihak (10-E A, S4-9 (i) A, diputuskan pemilik produk sendiri). Satu pihak dapat mempertahankan ambang atas pihak lain tanpa batas waktu (W-34; 07 §10.3 butir 12) |
| Mitigasi | Tujuh mitigasi di 10 §3.2: ambang lahir hanya dari kesepakatan; layar menyatakannya sebelum tombol setuju; ambang menandai, tidak menghalangi (INV-A11); mengikat setiap pihak sama (PG-01); hanya atas pembayaran rumah tangga; usulan selalu dapat diajukan tanpa hitungan penolakan; keluar tetap pemutus terakhir |
| Pemilik mitigasi | Pemilik produk (keputusan); `architect`; `finance-domain` |
| Harus ada pada | FS-1 |
| Tanda awal | Usulan "tanpa ambang" berulang yang tidak disetujui (dibaca di tingkat objek, bukan dihitung per orang) |
| Risiko sisa | K Sedang, D Sedang, **diterima pemilik produk** (W-34) |
| Sumber | 10 §3.2; 10.3 §2.4; W-34; 07 §10.3 butir 12; ADR-027 |

##### RISK-37 — Pelecehan lewat permintaan, pengingat, keberatan, dan chat

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Permintaan, pengingat, keberatan, atau pesan chat dipakai untuk menagih atau melecehkan (TM-28). Permintaan yang ditolak dapat diajukan lagi tanpa batas (W-8) |
| Mitigasi | Pengingat ulang paling banyak sekali per 6 jam, tiga per objek, plafon enam per hari, menghormati jam tenang (PG-06); alasan tertutup (06 §8.2); keberatan terlihat tetapi tidak menghalangi; sembunyikan untuk saya yang senyap (SEC-153); "Laporkan" (SEC-154); halaman Bantuan; keluar |
| Pemilik mitigasi | `architect` (06); `ux-designer`; `security-privacy` |
| Harus ada pada | FS-1 |
| Tanda awal | Keluhan SCN-30 ("Aku ditagih terus"), SCN-35 ("Keberatan dipakai untuk menyerang"), SCN-49; VL-08 |
| Risiko sisa | K Sedang, D Sedang. Pengajuan ulang tanpa batas (W-8); pesan yang sudah terbaca |
| Sumber | TM-28; W-8, W-11; 06 §10; SEC-153, SEC-154; SCN-30, SCN-35, SCN-49 |

##### RISK-38 — Membaca atau menyimpulkan data pribadi pasangan

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Sedang** · D **Kritis** · **P1** |
| Uraian dan sebab | Anggota dewasa memakai API langsung atau aplikasi yang dimodifikasi (TM-06), saluran samping seperti celah kursor sinkron, waktu di pengenal, beda pesan galat, atau hitungan (TM-07), fitur "admin" (TM-13), atau keberadaan sebutan (TM-36) untuk membaca atau menyimpulkan data pribadi pasangan. Brief §8: "Family membership never grants unlimited access to another adult's private data" |
| Mitigasi | Otorisasi server per perintah dan per bacaan (SEC-01; ES §4); satu jalur baca (SEC-06); pengenal acak dan kursor per pembaca (ADR-017; SEC-03, SEC-04); "tidak ada" sama dengan "tidak boleh" (SEC-05); tanpa hitungan per orang (SEC-07); uji "tidak pernah terlihat" sebagai setiap peran (SEC-86); uji simetri berpermutasi (SEC-89); tanpa API admin (SEC-33); pernyataan sebelum menyimpan hal yang akan terlihat pasangan (SEC-88; FR-FAM-010) |
| Pemilik mitigasi | `backend-dev`; `qa`; `code-reviewer`; `security-privacy` |
| Harus ada pada | FS-1 (uji di CI); GR-2 (10 §4.1 butir 1) |
| Tanda awal | Kegagalan uji privasi di CI, termasuk urutan kursor yang tidak identik byte demi byte (09.4 §4) |
| Risiko sisa | K Rendah, D Tinggi. Kebocoran B1 sampai B6 (07 §5.3; W-7) dan keterlihatan catatan sebelum bergabung (W-9) dinyatakan dan diterima lewat keputusan yang disebut di 07; pengelola server di luar cakupan ini (RISK-11) |
| Sumber | Brief §8; TM-06, TM-07, TM-13, TM-36; 10.3 §2.1; 07 §5.3; W-7, W-9; D-98, D-99 |

##### RISK-39 — Perpisahan: keluar sendiri satu-satunya jalan

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Brief §8 meminta perilaku saat pasangan berpisah dan anggota keluar. Di Tier 1 satu-satunya jalan adalah keluar sendiri (PG-13). Selama 7 hari masa tenggang pasangan yang tinggal tidak dapat memperoleh persetujuan apa pun dan tidak ada yang dapat menulis di chat (W-24). Ekspor orang yang keluar memuat tulisan pasangannya sampai F06 (L-33). Penutupan klaim saat keluar dapat dibaca sebagai pelepasan hak (L-22). Anggota yang tidak dapat dihubungi tidak dapat dikeluarkan (PA-FAM-6) |
| Mitigasi | Keluar dengan aman tanpa persetujuan siapa pun (06 F06); satu jendela ekspor dari F06 sampai akhir hari ke-90 sesudah F07, di aplikasi dengan kredensial akun, tidak lewat email (SEC-101; D-105); daftar klaim terbuka ditawarkan untuk diekspor lebih dulu (07 §8.2); pemberitahuan keluar hanya lewat ringkasan (SEC-56); keluar tidak memindahkan pemegang atau pemilik uang (07 INV-41); tidak ada pemilik tunggal yang diserahkan (SEC-110); diuji lewat SCN-44 |
| Pemilik mitigasi | `architect` (06); `finance-domain` (07); `security-privacy`; `ux-designer` |
| Harus ada pada | GR-2 |
| Tanda awal | — (peristiwa jarang; diuji lewat skenario) |
| Risiko sisa | K Sedang, D Tinggi. Chat tidak dapat dipakai pada pekan ketika berbicara paling diperlukan (W-24); kedudukan hukum catatan dalam perpisahan belum diketahui (RISK-71) |
| Sumber | Brief §8; W-24; 06 §3.5, §3.6; 07 §8.2, INV-28, INV-41; SEC-101 sampai SEC-110; L-22, L-33; SCN-44 |

#### 3.3.2 Pihak luar, perangkat, dan siklus data

##### RISK-40 — Akses tidak sah ke akun atau rumah tangga

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** (TM-11) · K **Sedang** · D **Kritis** · **P1** |
| Uraian dan sebab | Kredensial bocor dari layanan lain dipakai menebak kata sandi (TM-01); pasangan yang menguasai email mengatur ulang kata sandi (TM-11); tautan undangan diteruskan ke orang yang salah (TM-16) |
| Mitigasi | Argon2id, minimal 12 karakter, daftar kata sandi bocor di server, batas laju (SEC-11, SEC-12; 09.4 §2); token akses 15 menit, token penyegar berputar, pencabutan seketika (SEC-15); pemberitahuan masuk hanya kepada pemilik akun (SEC-17); pemulihan lewat tautan email sekali pakai 30 menit (SEC-24); undangan 7 hari sekali pakai, pengundang mengukuhkan email terverifikasi penerima (SEC-28 sampai SEC-31; 06 F04); pendaftaran tertutup di server, hanya dengan token undangan atau token pemasangan sekali pakai dari pengelola untuk satu orang tertentu, dengan tanggapan penolakan seragam (SEC-158, Pilot); **faktor kedua (SEC-18) dan kode pemulihan (SEC-25) sebagai satu pasangan fitur opsional per akun**, dibangun di FS-2: bila pemilik akun menyalakannya, pemulihan lewat email tidak melewati faktor kedua, dan kode pemulihan adalah jalan kembali bila perangkat faktor kedua hilang (keputusan 54; 10.1 SEC-18, SEC-25). **Klien web (TM-39):** cookie sesi `__Host-` `HttpOnly` `SameSite=Strict` dengan rotasi dan deteksi pemakaian ulang; RPC hanya POST dengan pemeriksaan `Origin` dan `Sec-Fetch-Site`; `frame-ancestors 'none'` (SEC-161 sampai SEC-163, SEC-166; ADR-030 butir 5 sampai 7) |
| Pemilik mitigasi | `backend-dev`; `web-dev`; `ios-dev`; `security-privacy` |
| Harus ada pada | GR-2 (SEC-158 dan mitigasi lainnya); FS-2 (pasangan SEC-18 dan SEC-25, tidak menahan GR-2; keputusan 54); sudah ada sebelum pengguna luar pertama (QG-31 (a)) |
| Tanda awal | Pemberitahuan masuk dari perangkat yang tidak dikenal (SEC-17); lonjakan penolakan batas laju |
| Risiko sisa | **Sejak GR-2 sampai pasangan SEC-18 dan SEC-25 terbangun di FS-2:** K Sedang, D Tinggi. Penguasa email dapat mengatur ulang kata sandi, dan pemiliknya mengetahuinya lewat SEC-17 (TM-11). Ini diterima pemilik produk sebagai harga "Opsional, setelah dipakai harian (Saran)" (keputusan 54, diterima atas saran lead; menjawab §11 (g) dan Q-227). **Sesudah FS-2:** K Rendah, D Tinggi bagi akun yang menyalakan pasangan itu. Bagi akun yang tidak menyalakannya, risiko sisa tetap K Sedang, D Tinggi, karena fiturnya opsional per akun. Selama tahan tulis, masuk, pemulihan, dan pencabutan sesi tetap berjalan (SEC-151; ADR-028 butir 3) |
| Sumber | TM-01, TM-11, TM-16; 10.1 §8 sampai §10; D-107 |

##### RISK-41 — Perangkat hilang, dicuri, atau dipegang orang lain

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Perangkat hilang (TM-14); ponsel yang terbuka dipakai bertindak atas nama pemiliknya (TM-10); anak 12+ di rumah tangga pilot membaca chat, uang, atau status pasangan di ponsel orang tua (TM-15; Q-186); layar kunci atau layar luar perangkat lipat membuka isi (TM-26) |
| Mitigasi | Kunci aplikasi Tier 1 (F-FAM-019; SEC-19); pengalih aplikasi tertutup (SEC-20); hapus jarak jauh saat sesi dicabut (SEC-21; ADR-024); proteksi berkas Class C dan CE, status kepedulian lokal Class A, dikecualikan dari cadangan awan dan transfer perangkat (SEC-46; D-102); teks layar kunci per kelas tanpa jumlah atau keadaan kesehatan (SEC-50 sampai SEC-54; ES §4). **Klien web (TM-40, TM-41; ADR-030):** IndexedDB tanpa proteksi berkas OS, jadi pengamannya kunci aplikasi web (SEC-174), penutup netral saat aplikasi tersembunyi, mode perangkat bersama untuk PC atau tablet yang dipakai bersama (sesi 30 menit diam dan 12 jam mutlak, tanpa salinan luring; SEC-165), keluar akun dan hapus jarak jauh di semua tab (SEC-164, SEC-172), satu akun per profil browser (SEC-170), judul, riwayat, dan isian otomatis tanpa data (SEC-175) |
| Pemilik mitigasi | `web-dev` (GR-2); `ios-dev`; `android-dev` (jalurnya) |
| Harus ada pada | GR-2 (klien web); QG-30 (iOS) dan QG-30 (Android) untuk bagian native |
| Tanda awal | — |
| Risiko sisa | K Sedang, D Sedang. Kunci aplikasi mati secara bawaan (D-107); ponsel yang terbuka di tangan orang lain memperlihatkan semuanya; catatan yang belum terkirim hilang saat sesi dicabut (D-102). Di web: browser yang sesinya dicabut tetapi tidak pernah dibuka lagi menyimpan salinan tanpa batas waktu, dinyatakan di layar cabut sesi (TM-41; SEC-172); pengelola kata sandi dan perekam di PC bersama di luar kendali aplikasi (TM-40) |
| Sumber | TM-10, TM-14, TM-15, TM-26, TM-40, TM-41; PL-03, PL-04; 10.1 §8; 10.2 §3; 10.5 |

##### RISK-42 — Server atau cadangan dibobol

| | |
|---|---|
| Penilaian | Keamanan dan privasi · K **Rendah** · D **Kritis** · **P2** |
| Uraian dan sebab | Lalu lintas disadap (TM-03), VPS dikuasai (TM-04), salinan cadangan bocor (TM-05) |
| Mitigasi | TLS ≥ 1.2, hanya porta 443 dan SSH berkunci (SEC-40, SEC-41); rahasia lewat `LoadCredential` (SEC-42); cadangan dienkripsi di VPS utama sebelum dikirim, kata sandinya tidak pernah di VPS cadangan (SEC-43); salinan pelindung yang tidak dapat ditulis VPS utama (09.4 §6); runbook insiden dan satu latihan (SEC-148 (a), Pilot); rencana pemberitahuan tertulis SEC-148 (b) ditunda ke go public (keputusan 68) |
| Pemilik mitigasi | Pengelola server; `backend-dev` |
| Harus ada pada | GR-2 (10 §4.1 butir 6, 8) |
| Tanda awal | Pemantau uptime; `ops.backup_verify` gagal (09.4 §3); pemberitahuan masuk SSH yang tidak dikenal |
| Risiko sisa | K Rendah, D Kritis. Tanpa E2EE seluruh data terbaca bila VPS dikuasai (TM-04); enkripsi disk hanya bila penyedia menawarkannya (SEC-45). **Sejak klien web:** VPS yang dikuasai juga menyajikan kode web kepada setiap perangkat (RISK-96 bagian bundel); endpoint langganan push palsu untuk SSRF ditolak oleh daftar izin host (TM-43; SEC-176 (g)) |
| Sumber | TM-03 sampai TM-05, TM-43; 10.2 §3; 09.4 §1, §6 |

##### RISK-96 — Kode yang berjalan di klien web bukan kode yang dimaksud

| | |
|---|---|
| Penilaian | Keamanan dan privasi · K **Rendah** · D **Kritis** · **P2** (baru, web lebih dulu) |
| Uraian dan sebab | Tiga jalan bagi kode asing atau basi di origin aplikasi: (1) dependensi npm atau alat build yang disusupi masuk ke bundel (TM-38); (2) XSS, yaitu skrip asing yang membaca IndexedDB dan memanggil API dengan cookie sesi selama halaman terbuka; `HttpOnly` mencegah pencurian token, bukan pemakaiannya (TM-37); (3) service worker terus menyajikan versi lama yang memuat celah yang sudah diperbaiki (TM-45). Akibatnya semua yang ada di perangkat dan semua yang terbaca anggota itu terbuka. K Rendah, sejajar RISK-42, karena daftar dependensi bundel tertutup dan kecil |
| Mitigasi | **Rantai pasok (TM-38):** daftar dependensi bundel tertutup (svelte, @sveltejs/kit, dexie, @connectrpc/connect, @connectrpc/connect-web, @bufbuild/protobuf; `workbox-*` dihapus 13 November 2026, ADR-030 butir 8), tambahan hanya lewat ADR; `pnpm install --frozen-lockfile`; skrip dependensi diblokir (`strictDepBuilds`, `allowBuilds` kosong); OSV-Scanner atas `pnpm-lock.yaml` di setiap perubahan; masa tunggu 7 hari bagi versi baru bila usulan ES §8 disetujui (ADR-030 butir 8, 17; SEC-168; 14 QG-03). **XSS (TM-37):** CSP ketat tanpa `unsafe-inline` untuk skrip, Trusted Types, larangan sink HTML di lint, tanpa skrip pihak ketiga (SEC-166, SEC-167; ADR-030 butir 7). **Versi lama (TM-45):** tanpa `skipWaiting` otomatis tetapi dengan tombol muat ulang, `client_version` minimum dengan `CLIENT_UPGRADE_REQUIRED` sejak pilot, `no-cache` untuk `index.html` dan service worker, service worker pengganti yang dilatih di staging (SEC-169 (b) sampai (e); 14 §6.5, §7.7). Hash artefak CI dibandingkan dengan berkas yang disajikan (SEC-169 (a)) |
| Pemilik mitigasi | `web-dev`; `code-reviewer` (QG-03, QG-11); pengelola server (deploy, latihan) |
| Harus ada pada | FS-1 langkah 1 (CSP, header, lockfile, pemindai); GR-2 (14 QG-30 butir 17; 10 §4.1 butir 13) |
| Tanda awal | Temuan OSV-Scanner; penolakan CSP di uji Playwright; perangkat yang `client_version`-nya tertinggal di daftar sesi |
| Risiko sisa | K Rendah, D Kritis. Penyusupan yang belum diketahui publik saat versi dipasang (TM-38 kolom sisa); masa tunggu 7 hari mengurangi, tidak menghapus. Skrip yang sudah berjalan memakai sesi yang sedang terbuka |
| Sumber | TM-37, TM-38, TM-45; 10.5 SEC-166 sampai SEC-169; ADR-025 butir 8; ADR-030; ES §4, §8 |

##### RISK-43 — Data peka keluar lewat log, push, atau metadata foto

| | |
|---|---|
| Penilaian | Keamanan dan privasi · K **Rendah** · D **Tinggi** · **P3** |
| Uraian dan sebab | Log memuat jumlah, kesehatan, atau teks pesan (TM-27); isi notifikasi terbaca di infrastruktur push (TM-18); metadata foto membuka lokasi atau perangkat (TM-25) |
| Mitigasi | Log dengan daftar izin yang dijaga tipe dan pemindaian log di CI (SEC-44; 09.4 §3, §4); payload push umum dengan teks kelas, diuji per N-NN (SEC-50, SEC-51); foto di-encode ulang menjadi JPEG tanpa penanda APPn atau COM, masukan PNG ditolak (SEC-66) |
| Pemilik mitigasi | `backend-dev`; `qa` |
| Harus ada pada | FS-1 |
| Tanda awal | Kegagalan uji log atau payload di CI |
| Risiko sisa | K Rendah, D Sedang. Token perangkat dan waktu kirim tetap lewat Apple dan Google (TM-18; L-16) |
| Sumber | TM-18, TM-25, TM-27; SEC-44, SEC-50, SEC-51, SEC-66 |

##### RISK-44 — Kebocoran lintas rumah tangga

| | |
|---|---|
| Penilaian | Keamanan dan privasi · K **Rendah** · D **Kritis** · **P2** |
| Uraian dan sebab | Objek rumah tangga lain terbaca dengan mengganti pengenal (TM-02) atau lewat constraint, pesan galat, cache, atau orang yang menjadi anggota dua rumah tangga (TM-22). Di pilot hanya ada satu rumah tangga; kemungkinannya naik saat rumah tangga beta memakai server yang sama (GR-4) |
| Mitigasi | `household_id` di setiap kunci (SEC-02); uji lintas rumah tangga yang dibangkitkan untuk setiap kueri (09.4 §4; ES §4); `household_id` sebagai pemilih yang diperiksa terhadap keanggotaan (SEC-38); **RLS (SEC-08) dipaksakan sejak pilot** (`FORCE ROW LEVEL SECURITY`; ADR-016 butir 3, diterima; D-99; 09.2 §3; 10 §4); uji isolasi dengan banyak rumah tangga sebelum GR-4 (14 QG-31) |
| Pemilik mitigasi | `backend-dev`; `qa` |
| Harus ada pada | GR-2 (uji isolasi dan RLS); GR-4 (uji isolasi dengan banyak rumah tangga), bila pengguna luar membentuk rumah tangga sendiri di server yang sama (keputusan 52) atau beta diadakan |
| Tanda awal | Kegagalan uji isolasi di CI |
| Risiko sisa | K Rendah, D Kritis (akibat kebocoran tidak berkurang; kemungkinannya yang ditekan). RLS dan uji isolasi menjadi dua lapis; yang tersisa adalah kesalahan kebijakan RLS dan kode yang sama-sama keliru, dan orang di beberapa rumah tangga (RISK-81, Tier 3) |
| Sumber | TM-02, TM-22; SEC-02, SEC-08, SEC-38; ES §4 |

##### RISK-45 — Penghapusan tidak tuntas

| | |
|---|---|
| Penilaian | Keamanan dan privasi · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Data yang dihapus tetap ada di cadangan, cache, ekspor, atau perangkat (TM-23); mantan anggota tetap membaca lewat sesi lama atau cache (TM-17); pemulihan cadangan menghidupkan kembali data yang sudah dihapus |
| Mitigasi | F06 mencabut semua sesi (SEC-23); pembuangan di perangkat saat akses menyempit (SEC-48); data yang dihapus bertahan paling lama 38 hari di salinan mana pun (SEC-96; 09.4 §6); jurnal penghapusan dijalankan ulang sebelum menerima lalu lintas sesudah pemulihan (SEC-97); uji pemulihan sebelum data nyata dan bulanan |
| Pemilik mitigasi | `backend-dev`; pengelola server |
| Harus ada pada | GR-2 (10 §4.1 butir 8) |
| Tanda awal | `ops.backup_verify` menemukan set cadangan lebih tua dari jangka simpan |
| Risiko sisa | K Rendah, D Sedang. Ekspor yang sudah diunduh pemiliknya (TM-23); cukup-tidaknya 38 hari menurut hukum: L-34 (RISK-67) |
| Sumber | TM-17, TM-23; SEC-94 sampai SEC-99; D-104 |

##### RISK-46 — Halaman Bantuan memuat nomor yang salah atau basi

| | |
|---|---|
| Penilaian | Keamanan dan privasi · **Pagar** · K **Rendah** · D **Kritis** · **P2** |
| Uraian dan sebab | Halaman Bantuan adalah Tier 1 (FR-FAM-024) dan jalan keluar bagi orang yang ditekan atau dalam bahaya. Healing119 dan 110 terverifikasi; SAPA 129 belum (W-33; R-12); 112 bukan nomor nasional (10.4 §1.7). Nomor yang salah atau sudah berubah merugikan orang pada saat paling rentan |
| Mitigasi | Hanya nomor terverifikasi yang tampil; SAPA 129 disembunyikan sampai diverifikasi bertanggal (SEC-138; 10 §4.1 butir 9); halaman statis yang dibuka tanpa jejak (SEC-134); verifikasi ulang di setiap gerbang dan sekurangnya setiap 6 bulan (SEC-138; kalender operasi 14 §7.7) |
| Pemilik mitigasi | `security-privacy`; pengelola server (verifikasi R-12) |
| Harus ada pada | GR-2, lalu setiap gerbang |
| Tanda awal | Tanggal verifikasi terakhir lebih tua dari 6 bulan |
| Risiko sisa | K Rendah, D Tinggi: nomor dapat berubah sesudah diverifikasi |
| Sumber | W-33; R-12; SEC-134, SEC-138; 10.4 §1.7; FR-FAM-024; SCN-49 |

### 3.4 Keuangan

#### RISK-47 — Aturan uang salah diimplementasikan

| | |
|---|---|
| Penilaian | Keuangan · K **Sedang** · D **Kritis** · **P1** |
| Uraian dan sebab | Brief §7: "Financial accuracy is non-negotiable". Satu aturan yang salah (pemindahan terhitung pengeluaran, persetujuan mengurangi uang, koreksi dengan menyunting) membuat angka yang dipercaya salah tanpa tanda. Sebabnya kerumitan aturan: 41 invarian keuangan (INV-01 sampai INV-41), 53 invarian alur (INV-A01 sampai INV-A53), dan aturan tanda (jatah, ambang, pos rutin) |
| Mitigasi | Satu implementasi aturan di server (ES §2; ADR-013); bilangan bulat (ADR-008); entri bertipe (ADR-002); saldo diturunkan dari baris (ADR-004); hanya-menambah (ADR-007); setiap invarian sebagai property test berstate di PostgreSQL nyata dan setiap contoh terhitung sebagai uji, memblokir merge (ES §5; 09.4 §4); vektor uji emas untuk tiga bantuan isian di klien (ES §2); `ops.invariant_check` dengan peringatan ke pengelola (09.4 §3); VL-10 nol; bila tidak nol, server masuk *tahan tulis* sampai penyebabnya diperbaiki, pencatatan berlanjut di antrean perangkat, dan perintah pelindung tetap berlaku (13 §6.2; ADR-028 butir 3); migrasi expand–contract yang dapat dibalik (ES §7) |
| Pemilik mitigasi | `backend-dev`; `qa`; `code-reviewer`; `finance-domain` (spesifikasi) |
| Harus ada pada | FS-1; GR-2 butir 3; FS-2 (VL-10) |
| Tanda awal | Kegagalan uji invarian di CI; peringatan `ops.invariant_check`; keluhan seperti SCN-28 |
| Risiko sisa | K Rendah, D Kritis. Uji hanya membuktikan aturan yang ditulis; aturan yang salah di 07 lulus semua uji. Pengamannya tinjauan `critic` dan hitung ulang contoh oleh lead (`CLAUDE.md` "Computed, not read") |
| Sumber | Brief §7; ES §2, §5, §6; 07 §4, §7; 06 invarian; ADR-002, ADR-004, ADR-007, ADR-008, ADR-013; 13 VL-10 |

#### RISK-48 — Kiriman ganda, sinkron luring, dan perintah bersamaan

| | |
|---|---|
| Penilaian | Keuangan · K **Tinggi** · D **Kritis** · **P1** |
| Uraian dan sebab | Brief §14 "unreliable offline sync"; brief §7 "concurrent updates; duplicate submissions". Jaringan buruk (brief §1 "reliable on poor networks") berarti kiriman ulang, ketukan ganda, dan dua perangkat yang luring bersamaan |
| Mitigasi | Kunci idempotensi per perintah tulis (ADR-005; INV-10); satu transaksi `SERIALIZABLE` per perintah (ADR-016; INV-23); server memeriksa ulang setiap perintah luring (SEC-151); antrean perangkat dengan jeda urungkan 5 detik, perintah pelindung tidak ditahan (06 S-7); tanpa perkiraan luring di pilot (D-100); penolakan tabrakan tanpa kebocoran (SEC-152); uji: setiap perintah dikirim ulang, perintah bersamaan dibandingkan dengan hasil berurutan (09.4 §4) |
| Pemilik mitigasi | `backend-dev`; `ios-dev`; `android-dev`; `qa` |
| Harus ada pada | FS-1; GR-2 butir 3 |
| Tanda awal | Metrik coba ulang `40001` (09.4 §3); perintah luring yang ditolak saat sinkron |
| Risiko sisa | K Rendah, D Sedang. Catatan yang belum terkirim hilang bila sesi perangkat dicabut, dinyatakan di layar cabut (D-102); pengguna luring melihat angka server tanpa catatannya yang tertunda |
| Sumber | Brief §1, §7, §14; ADR-005, ADR-016, ADR-018; 07 INV-10, INV-23, CTH-31; 06 §12 |

#### RISK-49 — Pengguna memilih jenis catatan yang salah

| | |
|---|---|
| Penilaian | Keuangan · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Model membedakan pemindahan, pengeluaran, alokasi, reservasi, dan klaim; kosakata keluarga tidak. Serah terima yang dicatat sebagai belanja membuat ringkasan menampilkan pengeluaran yang sebenarnya perpindahan custody. **Termasuk "unclear money ownership" (brief §14):** memegang uang dikira memiliki, atau uang bersama dicatat di tempat uang pribadi (dan sebaliknya), sehingga *uang yang ada* tidak memuat uang yang semestinya bersama, atau klaim antar-pemilik tidak tercatat. Semua invarian tetap benar, tetapi angka KS-1 dan KS-3 salah |
| Mitigasi | Kata layar di samping istilah domain (07 §1.4; 08); alur tersendiri untuk serah terima dan tarik tunai; koreksi lewat pembalik dan `recategorization`, tidak pernah menyunting (07 §3.12); jembatan laporan yang membuat ketidakcocokan terlihat (INV-19). **Pencegahan atas kepemilikan:** pemilik uang dan pemegang adalah dua hal terpisah di model (07 §2.2; glosarium *Pemilik uang*, *Pemegang*, *Custody*); serah terima memindahkan pemegang, tidak pernah pemilik (07 §2.2; INV-41 saat keluar); setiap tempat uang punya satu pemilik, `shared` atau satu anggota, dan dana pribadi berada di luar garis rumah tangga sehingga angka rumah tangga sama bagi keduanya (INV-05, INV-20); uang yang melintas garis selalu `owner_transfer` yang terlihat keduanya (07 §2.2); hanya M yang menulis buku pribadi M (INV-25). **Mitigasi:** catatan dengan pemilik yang keliru dibetulkan lewat jalur koreksi yang sama (pembalik, 07 §3.12), tanpa menyunting riwayat |
| Pemilik mitigasi | `ux-designer`; `finance-domain` |
| Harus ada pada | FS-1 |
| Tanda awal | VL-04 "angkanya dipercaya?" dijawab tidak; koreksi yang sering |
| Risiko sisa | K Sedang, D Sedang. Aplikasi tidak dapat mengetahui maksud pencatat |
| Sumber | Brief §5C ("no accounting vocabulary"), §14 ("unclear money ownership"); 07 §1.4, §2.2, §3.3, §3.12; INV-05, INV-20, INV-25, INV-41; 13 VL-04; SCN-29 |

#### RISK-50 — Celah label penghasilan

| | |
|---|---|
| Penilaian | Keuangan · K **Sedang** · D **Sedang** · **P3** · **Diterima** lewat keputusan yang disebut 07 §10.3 |
| Uraian dan sebab | Gaji yang dicatat sebagai hadiah atau "lainnya" di buku pribadi, atau tidak dicatat, tidak terlihat; seluruh penghasilan dapat diarahkan ke dana pribadi; makin sering uang masuk, makin besar celahnya (W-1, W-2, W-4; 07 §10.3 butir 1, 8) |
| Mitigasi | Penghasilan bersumber kelompok `work` selalu tercatat penuh di buku rumah tangga (INV-24); bagian pribadi hanya oleh penerimanya (INV-27); keberatan terlihat |
| Pemilik mitigasi | `finance-domain` |
| Harus ada pada | FS-1 |
| Tanda awal | Keluhan SCN-34 |
| Risiko sisa | K Sedang, D Sedang: pasangan hanya melihat bahwa penghasilan itu tidak muncul |
| Sumber | 07 §5.3 B5, §10.3 butir 1, 8; W-1, W-2, W-4; SCN-34 |

#### RISK-51 — Tanda dapat dihindari atau timpang

| | |
|---|---|
| Penilaian | Keuangan · **Pagar** · K **Sedang** · D **Sedang** · **P3** · **Diterima** lewat keputusan yang disebut 07 §10.3 |
| Uraian dan sebab | Penyisihan bertanggal periode lampau, penggantian hari mulai periode, tidak menautkan pembayaran ke pos rutin, dan mengakhiri pos rutin sendirian dapat menghindari tanda; anggota tanpa jatah lebih jarang ditandai daripada yang berjatah (07 §10.3 butir 7, 10, 11; W-5, W-13, W-14, W-17, W-25). Pos rutin berkategori lain pernah menjadi jalan lolos (TM-30) |
| Mitigasi | Tanggal pencatatan tampil bagi setiap pembaca entri (06 AT-6); penggantian hari mulai adalah kesepakatan (07 §6.1); kategori entri dan pos rutin wajib cocok (SEC-92; INV-37); asas pengawasan (ADR-027); tanda tidak pernah menghalangi (PG-10) |
| Pemilik mitigasi | `finance-domain`; `architect` |
| Harus ada pada | FS-1 |
| Tanda awal | — |
| Risiko sisa | K Sedang, D Rendah: semua jalan yang tersisa hanya dijaga keterlihatan, sebagaimana diterima (D-48; S3-5) |
| Sumber | 07 §10.3 butir 7, 10, 11; W-5, W-13, W-14, W-17, W-25; TM-30; ADR-027 |

#### RISK-09 — Eksponen IDR produk berbeda dari ISO 4217

| | |
|---|---|
| Penilaian | Keuangan · K **Rendah** · D **Sedang** · **P3** |
| Uraian dan sebab | Eksponen produk untuk IDR adalah 0; ISO 4217 menyebut 2 (09 §9 **[Terverifikasi]**). Alat lain atau pustaka platform yang membaca 2 salah membaca ekspor |
| Mitigasi | Eksponen hanya dari tabel produk; setiap ekspor membawa `minor_unit_exponent` (07 INV-39) |
| Pemilik mitigasi | `backend-dev` |
| Harus ada pada | FS-1 |
| Tanda awal | Uji impor ulang ekspor (SEC-103) |
| Risiko sisa | K Rendah, D Rendah |
| Sumber | 09.4 §9; 07 INV-39 |

### 3.5 Teknis

#### RISK-01 — Siklus UI iOS tanpa Mac terlalu lambat

| | |
|---|---|
| Penilaian | Teknis · K **Tinggi** · D **Sedang** · **P2** (dinilai ulang sesudah ADR-030: K dan D tetap, karena keadaannya sama begitu UI iOS dikerjakan; yang berubah fasenya) |
| Uraian dan sebab | Pengembangan di Linux; pratinjau SwiftUI, simulator, dan debugger hanya di macOS. Dengan ADR-022 opsi A setiap perubahan UI iOS diuji lewat build Xcode Cloud lalu TestFlight (W-28). **Sejak ADR-030 risiko ini keluar dari jalur kritis GR-2:** FS-1 dan GR-2 memakai klien web, yang dibangun dan diuji di Linux; harga ini baru dibayar di jalur iOS native (14 §3.9), ketika keluarga sudah memakai data nyata lewat web, sehingga keterlambatannya tidak menunda pemakaian harian. Yang tersisa di jalur web: Web Inspector untuk iPhone dan iPad juga memerlukan Mac (18 §5.13.8), dan ditanggung dengan cara 14 §4.8 "Uji klien web di iPhone dan iPad tanpa Mac" |
| Mitigasi | **Keputusan 43** (pemilik produk memilih "Tetap Xcode Cloud saja (Saran)"; `_progress.md`, "Penutupan pertanyaan terbuka sebelum S6"): tanpa Mac sewaan bulanan atau Mac sendiri sampai dana ada, jadi mitigasi lewat Mac tidak dipakai, dan siklus yang lebih lambat tanpa debugger di perangkat diterima (ADR-022 butir 4 dan catatan keputusan 43). *Sejak ADR-030 kalimat berikut digantikan:* "Sejak keputusan 42 iOS dan iPadOS satu-satunya platform UI di FS-1, jadi harga ini berlaku sejak langkah 1"; kini harga ini berlaku sejak jalur iOS langkah I1, dan sumber Mac sekali adalah AWS EC2 Mac atau Mac pinjaman (bawaan 69.6, diterima atas saran lead; pilihan di antara keduanya saat itu; ADR-022 catatan web). Yang tetap meringankan, tanpa biaya: kode non-UI sebagai paket Swift yang diuji di Linux (toolchain Swift di Linux **[Terverifikasi]**, 18 §5.2.7.4; perilaku API Foundation per dependensi **[Perlu validasi]**); XCUITest di simulator Xcode Cloud; perubahan layar dikumpulkan per build; uji perangkat menurut 14 §4.8 "Uji iOS tanpa Mac" (10 §0); perkiraan waktu FS-1 memperhitungkan siklus ini (14 §9). Bila dana ada, Mac sewaan bulanan (USD 149–164 **[Terverifikasi]**, Scaleway EUR 149 **[Perlu validasi]**; pratinjau dan simulator, tanpa debugger di iPhone fisik) atau Mac mini M6 sendiri (16 GB/512 GB Rp21.499.000 **[Perlu validasi]**; satu-satunya yang memberi debugger di perangkat pilot), dengan titik impas sekitar 7,1 sampai 7,9 bulan sebagai acuan (14 §9). xtool tidak dipakai sampai L-31 dijawab |
| Pemilik mitigasi | Pemilik produk (peninjauan Mac bila dana ada); `ios-dev` |
| Harus ada pada | Jalur iOS native, langkah I1 (14 §3.9); semula GR-1 dan awal FS-1 |
| Tanda awal | Layar iOS di jalur native tertinggal dari layar web yang sudah ada lebih dari satu langkah I2; jatah Xcode Cloud melewati 80% (14 §6.3); satu uji perangkat ternyata memerlukan Mac (14 §4.8) |
| Risiko sisa | K Tinggi, D Sedang. **Diterima** atas saran lead (keputusan 43); ditinjau bila dana ada. Bila satu uji keamanan perangkat memerlukan Mac, gerbangnya menunggu dan hal itu dibawa ke pemilik produk (10 §0) |
| Sumber | 09.4 §5, §9; ADR-022; ADR-030; W-28; L-31; keputusan 42, 43; 14 §3.9, §4.8, §9 |

#### RISK-02 — Pustaka 0.x berubah tidak kompatibel

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | connect-kotlin, `golang.org/x/crypto`, dan `golang.org/x/oauth2` masih 0.x (09 §9 **[Terverifikasi]**) |
| Mitigasi | Dipatok dengan lockfile (ES §8); jalan mundur per pustaka di ADR-023; pemutakhiran bulanan dengan membaca catatan rilis |
| Pemilik mitigasi | `backend-dev`; `android-dev` |
| Harus ada pada | FS-1 |
| Tanda awal | Catatan rilis yang menyebut perubahan tidak kompatibel |
| Risiko sisa | K Sedang, D Rendah: pekerjaan ulang di satu paket |
| Sumber | 09.4 §9; ADR-023 |

#### RISK-03 — Satu VPS titik gagal tunggal

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Satu VPS menjalankan API, SSE, pekerja, dan PostgreSQL (ADR-014). Bila mati, chat, persetujuan, dan sinkron berhenti |
| Mitigasi | RPO ≤ 1 menit, RTO ≤ 4 jam **[Usulan]**; pencatatan luring tetap jalan (LM-02); cadangan di penyedia lain (09.4 §6) |
| Pemilik mitigasi | Pengelola server |
| Harus ada pada | GR-2 |
| Tanda awal | Pemantau uptime `/healthz` |
| Risiko sisa | K Sedang, D Rendah: beberapa jam tanpa layanan |
| Sumber | 09.4 §6, §9 |

#### RISK-05 — Cadangan tidak andal (pgBackRest dan PG18)

| | |
|---|---|
| Penilaian | Teknis · K **Rendah** · D **Kritis** · **P2** |
| Uraian dan sebab | Dukungan PG18 di pgBackRest diperkenalkan sebagai eksperimental (2.55.0) dan status "penuh" belum dikukuhkan; `repo1-type=sftp` **[Perlu validasi]** (09 §9; ADR-020). Cadangan yang tidak dapat dipulihkan berarti data rumah tangga hilang |
| Mitigasi | Uji pemulihan sebelum data nyata dan bulanan, dengan `ops.invariant_check` (09.4 §6); WAL-G sebagai alternatif |
| Pemilik mitigasi | Pengelola server |
| Harus ada pada | GR-2 (10 §4.1 butir 8) |
| Tanda awal | Uji pemulihan gagal; `ops.backup_verify` gagal |
| Risiko sisa | K Rendah, D Sedang |
| Sumber | 09.4 §6, §9; ADR-020 |

#### RISK-06 — Email pemulihan tidak sampai

| | |
|---|---|
| Penilaian | Teknis · K **Rendah** · D **Sedang** · **P3** |
| Uraian dan sebab | Email masuk spam atau melampaui kuota sandbox; penyedia email menunggu S4-2 (ADR-026; K-62) |
| Mitigasi | SPF, DKIM, DMARC; uji kirim sebelum pilot; pengiriman sebagai job yang dicoba ulang (ADR-026); keluar dari sandbox sebelum pengguna luar (09 §6) |
| Pemilik mitigasi | `backend-dev`; pengelola server |
| Harus ada pada | GR-2; GR-4 (sandbox), sebelum pengguna luar pertama |
| Tanda awal | Job email yang gagal berulang |
| Risiko sisa | K Rendah, D Rendah |
| Sumber | 09.4 §2, §9; ADR-026 |

#### RISK-07 — Foto di PostgreSQL membesarkan cadangan

| | |
|---|---|
| Penilaian | Teknis · K **Rendah** · D **Rendah** · **P3** |
| Uraian dan sebab | Foto disimpan di basis data (09.2 §11), jadi ikut setiap cadangan |
| Mitigasi | Batas pindah ke penyimpanan objek: 5 GB foto atau 20 GB basis data (09 §6) |
| Pemilik mitigasi | `backend-dev` |
| Harus ada pada | FS-2 (pemantauan) |
| Tanda awal | Ukuran mendekati batas |
| Risiko sisa | K Rendah, D Rendah |
| Sumber | 09.4 §9; 09 §6 |

#### RISK-08 — SSE terputus di jaringan seluler

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Rendah** · **P3** |
| Uraian dan sebab | Proxy operator seluler dapat memutus koneksi panjang |
| Mitigasi | Tarik `Sync` saat tersambung ulang; push saat aplikasi tidak di depan (ADR-018) |
| Pemilik mitigasi | `backend-dev` |
| Harus ada pada | FS-1 |
| Tanda awal | Pengujian `qa` di jaringan seluler pilot |
| Risiko sisa | K Sedang, D Rendah: chat terasa lebih lambat daripada WhatsApp |
| Sumber | 09.4 §9; ADR-018 |

#### RISK-10 — Satu pemegang pengetahuan (bus factor)

| | |
|---|---|
| Penilaian | Teknis · K **Tinggi** · D **Tinggi** · **P1** |
| Uraian dan sebab | Pemilik produk adalah pemimpin teknis, pengelola server, dan pemegang kunci cadangan. Bila ia berhalangan lama, layanan tidak dioperasikan siapa pun |
| Mitigasi | Runbook pemulihan, rilis, dan rotasi rahasia ditulis di fase bangun (09.4 §9; 14); komponen sedikit (ADR-014); salinan kunci cadangan offline (09.4 §1) |
| Pemilik mitigasi | Pemilik produk; `architect` (14) |
| Harus ada pada | GR-2 (runbook ada) |
| Tanda awal | Runbook belum ada saat GR-2 |
| Risiko sisa | K Sedang, D Tinggi. Memberi istri akses pengelola bila pemilik produk berhalangan juga memberinya akses teknis ke data pribadi suaminya, cermin RISK-11. Itu keputusan pemilik produk, tidak diusulkan di sini (pertanyaan diminta di §11) |
| Sumber | 09.4 §9; RISK-11 |

#### RISK-13 — Pembungkus SQLite iOS tidak dapat dibangun di Linux

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Rendah** · **P3** |
| Uraian dan sebab | Bila pembungkus yang dipilih tidak dapat dibangun di Linux, uji non-UI iOS hanya berjalan di Xcode Cloud |
| Mitigasi | Dinilai `ios-dev` di minggu pertama jalur iOS native (09 §6; 09.3 §2; 14 §3.9 I1) |
| Pemilik mitigasi | `ios-dev` |
| Harus ada pada | Jalur iOS native, langkah I1 (semula FS-1 minggu pertama; ADR-030) |
| Tanda awal | Build paket Swift gagal di CI Linux |
| Risiko sisa | K Sedang, D Rendah |
| Sumber | 09.4 §9; 09 §6 |

#### RISK-17 — API eksperimental Compose berubah

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Rendah** · **P3** |
| Uraian dan sebab | `ModalBottomSheet` bertanda eksperimental di Material 3 1.4.0, dan `rememberModalBottomSheetState` sudah `@Deprecated` di kode pengembangan (09 §9 **[Terverifikasi]**) |
| Mitigasi | Dipatok ke 1.4.0; jalan mundur dialog layar penuh (ADR-012 butir 5; ADR-023 butir 5); uji UI lembar Catat dijalankan ulang setiap naik versi |
| Pemilik mitigasi | `android-dev` |
| Harus ada pada | Jalur Android (14 §3.8) |
| Tanda awal | Catatan rilis 1.5.0 |
| Risiko sisa | K Sedang, D Rendah |
| Sumber | 09.4 §9; ADR-012, ADR-023; 08 §17 |

#### RISK-52 — Build TestFlight kedaluwarsa; jalur rilis iOS bergantung pada Apple

| | |
|---|---|
| Penilaian | Teknis · K **Tinggi** · D **Sedang** · **P2** |
| Uraian dan sebab | Build TestFlight aktif 90 hari sejak diunggah, lalu tidak tersedia bagi penguji (**[Terverifikasi]** 18 §2.12.6), hampir sama dengan jendela tiga bulan. Apakah aplikasi yang sudah terpasang berhenti dapat dibuka sesudahnya **[Perlu validasi]** (18 §2.12.6); bila ya, aplikasi tidak dapat dibuka di iPhone 13 dan iPad mini 7 sampai build baru terpasang, sehingga KS-1 terganggu dan keadaan itu terhitung bug berat "aplikasi tidak dapat dipakai" yang menunda GR-3 (VL-02 (b), keputusan 46). Jalur iOS juga bergantung pada Xcode Cloud dan akun Apple; artefak hanya disimpan 30 hari (ADR-022) |
| Mitigasi | Pekerjaan kalender di runbook rilis untuk tanggal 60, 80, dan 90 hari sejak unggahan (ADR-022 butir 1 sebagaimana diubah ADR-028 butir 4; 14 §6.3, §7.7); rilis pemeliharaan bila tidak ada gelombang Tier 2 dalam 60 hari, dan batas keras 80 hari sejak unggahan (ADR-028 butir 4, diterima atas saran lead, keputusan 44; 14 QG-23). Jalan mundur GitHub Actions macOS dengan API `buildUploads` sebagai pemeriksaan wajib sebelum merge (ADR-022 catatan S5; 14 §5.1); versi rilis dapat dibangun ulang dari tag |
| Pemilik mitigasi | `ios-dev`; pemilik produk (pemegang akun) |
| Harus ada pada | Jalur iOS native, sejak build TestFlight pertama (pekerjaan kalender; 14 QG-23); semula FS-1. Klien web tidak punya batas umur build (ADR-028 catatan web), jadi risiko ini tidak menyentuh GR-2 |
| Tanda awal | Lebih dari 60 hari sejak build TestFlight terakhir diunggah |
| Risiko sisa | K Rendah, D Sedang. Lihat juga RISK-64 bila deskriptor rating usia ternyata menahan unggahan |
| Sumber | ADR-022; 13 §3.4, §8; 18 §2.12.6 |

#### RISK-53 — Kapasitas: dua UI native dan cakupan Tier 1 untuk tim 1–3 orang

| | |
|---|---|
| Penilaian | Teknis · K **Tinggi** · D **Sedang** · **P2** |
| Uraian dan sebab | Pemilik produk memutuskan UI native penuh per platform (Swift dan Kotlin; Q-34), jadi setiap layar dibangun dua kali; **sejak ADR-030 tiga kali** (web, lalu iOS, lalu Android), dengan TypeScript sebagai bahasa keempat. Tier 1 kini 55 fitur dan 198 FR Must (`_progress.md`). Brief §1 tidak memberi tenggat dan lebih memilih cakupan terkecil yang dipakai harian. Risikonya GR-2 menjauh sampai motivasi turun, atau tekanan memotong uji dan keamanan Tier 1, yang dilarang (`CLAUDE.md` "Instruction priority" butir 3) |
| Mitigasi | Semua aturan sekali di server, sehingga klien tipis (ADR-013); tipe klien dibangkitkan dari kontrak (ES §2); layar dari komponen design system (ES §2); fitur Tier 2 tidak menahan GR-2 (13 §3.2); satu klien UI sampai GR-2, kini klien web yang dibangun dan diuji seluruhnya di Linux, lalu iOS native dan Android sebagai jalur sendiri sesudahnya (ADR-030; 14 §3.4, §3.8, §3.9); jalur native tidak menahan pemakaian harian karena web sudah memegang data nyata; urutan pekerjaan di dalam Tier 1 di 14 |
| Pemilik mitigasi | Pemilik produk; `architect` (14); lead |
| Harus ada pada | GR-1 |
| Tanda awal | Alur ujung ke ujung pertama belum terpasang di perangkat pilot jauh sesudah FS-1 dimulai (14 menetapkan titik ukurnya) |
| Risiko sisa | K Sedang, D Sedang. Harga keputusan native diterima (brief §9.1 menaruh tampilan native di atas berbagi kode). Klien web mempercepat GR-2 tetapi menambah satu UI yang terus dirawat; bila beban tiga UI terlalu berat sesudah aplikasi native ada, pensiun klien web di perangkat yang sudah pindah adalah keputusan pemilik produk, bukan dokumen ini. Mengubah tier untuk mempercepat pilot hanya lewat keputusan pemilik produk di checkpoint (brief §1.1); dokumen ini tidak mengusulkannya |
| Sumber | Brief §1, §9.1; Q-34; ADR-012, ADR-013; 13 §3.2 |

#### RISK-54 — Sasaran kinerja ES §3 tidak tercapai di perangkat pilot

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Brief §1 menuntut cold start seketika dan aplikasi ringan; ES §3 menetapkan sasaran `[DEFAULT]`. Galaxy A55 adalah perangkat pilot dengan kinerja paling rendah **[Asumsi]** |
| Mitigasi | **Sejak ADR-030:** pengukuran klien web di iPhone 13, iPad mini 7, dan Galaxy A55 sebelum GR-2, termasuk ukuran app shell terkompresi ≤ 500 KB **[Usulan]** dan cold start dari ikon web app (ADR-029 catatan web; ADR-030 butir 16; 14 §4.10), dikukuhkan atau diubah lewat ADR (ES §3; 13 §3.3 butir 5); aplikasi native diukur sebelum QG-30 (iOS) dan QG-30 (Android), dengan ADR yang sama direvisi bila perlu; tanpa kerja jaringan atau disk di thread UI, daftar terpaginasi, indeks untuk setiap kueri layar (ES §3) |
| Pemilik mitigasi | `web-dev`; `ios-dev`; `android-dev`; `backend-dev` |
| Harus ada pada | GR-2 (klien web di ketiga perangkat); QG-30 (iOS); QG-30 (Android) |
| Tanda awal | Pengukuran pertama di perangkat |
| Risiko sisa | K Rendah, D Sedang |
| Sumber | Brief §1; ES §3; 09.4 §4 |

#### RISK-55 — Rekayasa berlebihan

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Brief §2 dan §14: "do not over-engineer v1 for traffic it will not have". Godaan menambah broker, layanan terkelola, atau abstraksi sebelum ada kebutuhan |
| Mitigasi | Monolit modular satu proses di satu VPS (ADR-014); tanpa broker; pemicu tertulis untuk setiap penskalaan (09 §6); tanpa abstraksi tanpa pemakaian kedua (ES §1); tinjauan dua lapis (ES §7) |
| Pemilik mitigasi | `architect`; `code-reviewer` |
| Harus ada pada | FS-1 |
| Tanda awal | Komponen baru tanpa pemicu 09 §6 |
| Risiko sisa | K Rendah, D Sedang. Model yang dirancang untuk semua rumah tangga tahap publik (aturan N anggota dewasa, isolasi rumah tangga) adalah fondasi yang disengaja, bukan rekayasa berlebihan (brief §1.1) |
| Sumber | Brief §2, §14; ADR-014; ES §1; 09 §6 |

#### RISK-56 — Kunci tanda tangan APK dan distribusi APK

| | |
|---|---|
| Penilaian | Teknis · K **Rendah** · D **Sedang** · **P3** |
| Uraian dan sebab | Pilot Android memasang APK langsung tanpa Play (09.4 §5). Bila kunci tanda tangan hilang, pembaruan tidak dapat dipasang di atas aplikasi lama; memasang ulang menghapus antrean catatan yang belum terkirim |
| Mitigasi | Kunci disimpan offline oleh pemilik produk dan sebagai rahasia CI (09.4 §5), dibuat sebelum APK pertama (10 §4.3 SEC-42; 14 §3.8 A1) |
| Pemilik mitigasi | `android-dev`; pemilik produk |
| Harus ada pada | Jalur Android langkah A1 (14 §3.8) |
| Tanda awal | — |
| Risiko sisa | K Rendah, D Rendah |
| Sumber | 09.4 §5 |

#### RISK-57 — Ketergantungan pada layanan pihak ketiga

| | |
|---|---|
| Penilaian | Teknis · K **Rendah** · D **Sedang** · **P3** |
| Uraian dan sebab | Brief §14 "integration dependencies". Tier 1 bergantung pada layanan push browser (Apple Web Push, layanan push Chrome, dan lainnya; satu-satunya kanal push sampai jalur native) dan penyedia email; sejak jalur native juga APNs, FCM, dan Xcode Cloud; Tier 2 pada penyedia `AI`. Tidak ada integrasi bank atau marketplace di MVP (A-05) |
| Mitigasi | Push gagal: SSE dan tarik saat aplikasi dibuka (ADR-018, ADR-019); host layanan push browser dibatasi daftar izin, dengan host yang belum terverifikasi ditandai **[Perlu validasi]** di ADR-030 butir 14 (SEC-176 (g)); email gagal: job dicoba ulang (ADR-026); `AI` gagal: kalimat baku, dan kegagalan `AI` tidak pernah menghalangi pencatatan atau persetujuan (brief §5O; PG-12); Xcode Cloud gagal: GitHub Actions macOS (ADR-022) |
| Pemilik mitigasi | `backend-dev` |
| Harus ada pada | FS-1 |
| Tanda awal | Kegagalan job berulang |
| Risiko sisa | K Rendah, D Rendah |
| Sumber | Brief §14, §5O; ADR-018, ADR-019, ADR-022, ADR-026, ADR-030; A-05 |

#### RISK-91 — Android native menyusul: aplikasi Android teruji belakangan

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Sedang** · **P3** (baru, keputusan 42; **dinilai ulang sesudah ADR-030: D Tinggi → Sedang, P2 → P3**) |
| Penilaian ulang (web lebih dulu) | Galaxy A55 memegang data nyata lewat klien web di Chrome sejak GR-2, jadi akibat terberat versi sebelumnya hilang: pengeluaran pemilik produk di luar rumah tidak lagi tercatat belakangan dari iPad, dan kinerja, sinkron, tata letak, serta Web Push di Galaxy A55 diuji sejak GR-2 (14 QG-30 butir 1, 3, 4). Yang tersisa: komponen Compose, kunci aplikasi Android, SEC-46, dan FCM pertama kali dipakai lebih belakang lagi, sesudah jalur iOS native (14 §3.8); aplikasi Android bergabung ke rumah tangga yang sudah berisi data berbulan-bulan (butir (4) di bawah); dan bila jalur publik diambil, beta tetap menjadi uji Android native yang sebenarnya. Itu melemahkan satu gerbang opsional, bukan KS rumah tangga, jadi D Sedang. K tetap Sedang. Teks di bawah adalah penilaian sebelumnya; kalimat tentang iPad dan RISK-18 digantikan paragraf ini |
| Uraian dan sebab | Keputusan 42: pilot dengan data nyata di iOS dan iPadOS; Android dibangun sesudah GR-2 untuk Galaxy A55 pemilik produk dan bergabung sebelum beta, bila beta kelak diadakan (semula bawaan lead; diterima atas saran lead lewat keputusan 50; bunyi brief §1 "Pilot devices" sesudah keputusan 51). Jalur Android tetap dibangun untuk keluarga, walaupun jalur publik tidak diambil. Akibatnya: (1) layar Compose, kunci aplikasi Android, pengecualian cadangan (SEC-46), FCM, dan kinerja di Galaxy A55, perangkat pilot dengan kinerja paling rendah **[Asumsi]** (RISK-54), pertama kali dipakai sungguh-sungguh jauh sesudah iOS; (2) bila jalur publik diambil dan jalur Android selesai dekat GR-4, pemakaian Android dengan data nyata sebelum penguji beta (sekurangnya 14 penguji Android bila akun Play pribadi, 14 §6.7) singkat atau tidak ada, jadi beta menjadi uji Android yang sebenarnya; (3) jalur Android bersaing dengan gelombang Tier 2 untuk waktu tim yang sama (RISK-53); (4) aplikasi Android bergabung ke rumah tangga yang sudah berisi data nyata berbulan-bulan, jadi potret awal dan sinkron pertamanya menghadapi volume yang tidak dialami iOS saat GR-2. Bila jalur itu tergelincir, pengeluaran pemilik produk di luar rumah tetap tercatat belakangan dari iPad (RISK-18), dan, bila jalur publik diambil, GR-4 menunggu (14 §3.8) |
| Mitigasi | Satu server, satu kontrak, semua aturan di server (ADR-013), jadi Android tidak membawa aturan baru dan `buf breaking` menjaga kontrak sejak FS-1; vektor emas Kotlin lulus atas seluruh berkas yang ada saat modul pertama ditambahkan (14 §4.5, QG-06); **QG-30 (Android)** sebelum APK memakai data nyata: setiap AC Tier 1 yang perlu perangkat, 38 skenario, daftar 10 §4.3, ES §3 di Galaxy A55, tata letak 08 §16.4, potret awal atas rumah tangga sintetis seukuran data pilot **[Usulan]** (14 §5.3); bila jalur publik diambil, QG-31 mensyaratkan QG-30 (Android) dan APK terpasang sebelum jalur Play dibuka (13 GR-4 butir 8); rilis Play publik hanya sesudah konsultasi advokat sebelum go public, satu-satunya pemicu konsultasi ulang (keputusan 67, diputuskan pemilik produk sendiri); waktu mulai jalur Android diputuskan pemilik produk di GR-2 dan dapat dimajukan bila RISK-18 menuntutnya (13 §3.2) |
| Pemilik mitigasi | Pemilik produk (jadwal); `android-dev`; `qa` |
| Harus ada pada | QG-30 (Android), sesudah GR-2; sebelum jalur pengujian Google Play dibuka (GR-4 butir 8; QG-31 (b)) |
| Tanda awal | Jalur Android belum dimulai sesudah gelombang Tier 2 terakhir dilepas; temuan QG-30 (Android) di sinkron, potret awal, atau kinerja |
| Risiko sisa | K Rendah, D Sedang. Lamanya pemakaian Android oleh keluarga bergantung jadwal; bila jalur publik diambil dan pemakaian itu singkat, beta adalah uji Android sesungguhnya, dengan pagar QG-31 yang sama |
| Sumber | Keputusan 42, 51; ADR-030; 17 K-66; 13 §3.2, §3.4, GR-4 butir 8; 14 §3.8, §3.9, §5.3, §6.7; 10 §4.3; RISK-53, RISK-54 |

#### RISK-92 — TestFlight internal sebagai jalur iOS keluarga tanpa batas waktu tidak sesuai aturan Apple

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Tinggi** · **P2** · **Diterima atas saran lead** (keputusan 55) (baru, keputusan 51; dinilai dengan 18 §5.2.9; K dan D tidak berubah dari penilaian sementara) |
| Uraian dan sebab | Jalur iOS satu-satunya adalah Xcode Cloud ke TestFlight internal (ADR-022 butir 1). Sejak keputusan 51, jalur itu menjadi saluran distribusi keluarga tanpa batas waktu, bukan uji menuju rilis. Temuan `researcher` (18 §5.2.9):<br>- Guideline 2.2 kemungkinan tidak jelas menjangkau build internal **[Usulan]**.<br>- Kontraknya, DPLA 7.4.1 (29 Oktober 2026), membatasi TestFlight internal pada "pre-release versions", "solely for their internal use in testing, evaluating and/or developing", untuk "members of Your company or organization", dan Apple dapat menghentikannya "at any time in its sole discretion" **[Terverifikasi]**.<br>- Catatan penegakan tidak ditemukan.<br><br>K **Sedang** karena pemakaian harian keluarga berada di tepi bunyi kontrak dan penghentian sepenuhnya kewenangan Apple, tetapi tanpa riwayat penegakan. **Akibatnya kehilangan akses, bukan kebocoran data:** build baru tidak terpasang, dan build terpasang kedaluwarsa dalam 10 sampai 90 hari (RISK-52; irama 80 hari). Akibatnya iPhone 13 (pencatat utama) dan iPad mini 7 kehilangan aplikasi. Data tetap di server; yang dapat hilang hanya antrean perangkat yang belum terkirim.<br><br>Tidak ada saluran Apple yang sekaligus resmi untuk keluarga tanpa badan hukum, tanpa Mac, dan tanpa batas 90 hari (18 §5.2.9.5). Ad Hoc memerlukan Mac (keputusan 43) dan kemungkinan Developer Mode; App Store unlisted memerlukan App Review dan prasyarat toko |
| Mitigasi | **Keputusan 55** (Q-228; diterima atas saran lead): tetap TestFlight internal, pindah ke App Store unlisted (C) bila Apple menghentikannya. ADR-022 (catatan keputusan 51) memuat dua pemicu C, syarat C-1 sampai C-7, dan jalan akun demo App Review (demo mode di perangkat dengan izin Apple). Ad Hoc ditolak selama keputusan 43 berlaku dan selama memerlukan Developer Mode (10 §0). Pendukung lainnya:<br>- server sebagai sumber kebenaran (ADR-013), sehingga pindah saluran tidak memindahkan data;<br>- batas keras 80 hari (ADR-028 butir 4), sehingga build terpasang berlaku sekurangnya 10 hari lagi saat penghentian;<br>- saat pemicu datang, antrean perangkat dikosongkan lebih dulu;<br>- saluran baru ditinjau `security-privacy` atas 10 §0 dan SEC-41 (d) sebelum dipakai (10 §7.8);<br>- pertanyaan advokat DPLA 7.4.1, 7.3, dan pengecualian rumah tangga (18 §5.2.9), tidak dikejar sampai go public (keputusan 67) |
| Pemilik mitigasi | Pemilik produk (pilihan jalur; pemegang akun); `architect` (ADR-022); `ios-dev`; `security-privacy` (10 §0); advokat |
| Harus ada pada | Jalur dipilih (keputusan 55, selesai); GR-2: SEC-158 (syarat C-3, Pilot; berlaku bagi semua klien). **Sejak ADR-030** jalur TestFlight baru dipakai di jalur iOS native (14 §3.9), jadi risiko ini mulai berlaku saat build TestFlight pertama; selama web menjadi klien iPhone dan iPad, penghentian TestFlight tidak menghilangkan aplikasi keluarga. Syarat C lainnya hanya saat pemicu datang |
| Tanda awal | Galat kontrak beta atau build TestFlight yang tiba-tiba tidak tersedia di App Store Connect (cuplikan forum `BETA_CONTRACT_MISSING`, 18 §5.2.9.1 **[Perlu validasi]**); permintaan Apple untuk menghentikan distribusi |
| Risiko sisa | **K Sedang · D Sedang, diterima atas saran lead (keputusan 55, pilihan 1).** Wujudnya jeda tanpa aplikasi iOS selama syarat C dan App Review berjalan, dikurangi sisa masa build (10 sampai 90 hari); lama App Review **[Asumsi]** belum diketahui. Data tidak hilang, kecuali antrean perangkat yang belum terkirim (syarat C-7). Pilihan yang tidak diambil: pilihan 2 (K Sedang, D Rendah) dan pilihan 3 (risiko ini hampir hilang, diganti beban App Review per rilis, RISK-16), dicatat di ADR-022 sebagai bahan bila keputusan 55 dibuka kembali |
| Sumber | Keputusan 43, 51, 52; ADR-022 (catatan keputusan 51, catatan web); ADR-028 butir 4; ADR-030; 10 §3.3, §7.8, §8; 18 §2.12.6, §5.2.9; RISK-16, RISK-52 |

#### RISK-94 — Browser menghapus antrean catatan yang belum terkirim

| | |
|---|---|
| Penilaian | Teknis · K **Sedang** · D **Tinggi** · **P2** (baru, web lebih dulu) |
| Uraian dan sebab | Klien web menyimpan antrean perintah luring di IndexedDB. Browser dapat menghapus data situs: aturan 7 hari Safari untuk data skrip di tab tanpa interaksi (web app layar utama berpenghitung sendiri), tekanan ruang, atau pengguna menghapus data situs (TM-44; 18 §5.13.2, §5.13.9 risiko 1). Safari tidak punya Background Sync, jadi antrean hanya terkirim saat aplikasi terbuka (ADR-018 catatan web). Catatan yang hilang sebelum terkirim membuat angka utama kurang tanpa diketahui, persis kegagalan yang diwaspadai RISK-20 dan brief §7, dan bertabrakan dengan "pencatatan tanpa beban" (brief §1.1; 10 §8 butir 19) |
| Mitigasi | Di iOS dan iPadOS hanya web app layar utama yang menawarkan masuk; tab Safari hanya menampilkan petunjuk pasang **[Usulan]** (ADR-030 butir 11; SEC-171), lalu `navigator.storage.persist()` diminta dan hasilnya dicatat per perangkat (14 QG-30 butir 16); antrean dikirim seketika saat daring dan dikunci satu pengirim lewat Web Locks, sehingga umurnya pendek (SEC-173); jumlah catatan belum terkirim tampil, dan peringatan hanya muncul bila ada yang belum terkirim (SEC-171); IndexedDB yang terhapus sementara cookie sesi masih ada terdeteksi dan pesan kehilangan tampil, bukan diam (SEC-171 (e)); uji Playwright menghapus IndexedDB dan menutup konteks dengan antrean berisi (14 §4.7) |
| Pemilik mitigasi | `web-dev`; `ux-designer` (kata peringatan); `qa` (pemeriksaan di perangkat) |
| Harus ada pada | GR-2 (14 QG-30 butir 16; 10 §4.1 butir 12) |
| Tanda awal | `persist()` ditolak di salah satu perangkat; pesan kehilangan antrean pernah tampil; jumlah catatan belum terkirim yang sering lebih dari nol saat aplikasi ditutup |
| Risiko sisa | K Rendah, D Tinggi. Berkurang, tidak hilang (TM-44 kolom sisa): pengguna yang menghapus data situs sendiri atau browser yang menolak `persist()` tetap dapat kehilangan catatan yang belum terkirim, tetapi kini tahu. Penutup sebenarnya adalah aplikasi native (10 §8 butir 19) |
| Sumber | TM-44; 10.5 SEC-171, SEC-173; 10 §8 butir 19; 18 §5.13.2, §5.13.4, §5.13.9; ADR-018, ADR-030; brief §1.1, §7; RISK-20 |

### 3.6 Bisnis

#### RISK-16 — App Review menolak karena Guideline 1.2

| | |
|---|---|
| Penilaian | Bisnis · K **Sedang** · D **Tinggi** · **P2** · **Jalur publik** (§0; TestFlight internal rumah tangga tidak melewati App Review) |
| Uraian dan sebab | Guideline 1.2 menuntut penyaringan, laporan, dan blokir atas konten buatan pengguna dan tidak mengecualikan chat tertutup (18 §2.12). Padanan 10-G A memenuhi setiap unsur tanpa blokir harfiah; peninjau dapat menuntut blokir harfiah (W-30) |
| Mitigasi | SEC-153 (d) dan sisa SEC-154 sebelum build pertama ke TestFlight eksternal atau App Store (ADR-022; 06 §16.12); penjelasan di catatan App Review; R-13 |
| Pemilik mitigasi | `ios-dev`; `security-privacy` |
| Harus ada pada | GR-4 |
| Tanda awal | Hasil Beta App Review pertama |
| Risiko sisa | K Sedang, D Sedang: rilis tertunda sampai dijelaskan (W-30). Blokir harfiah yang mengeluarkan orang dewasa lain tidak akan dibangun (brief §4.2) |
| Sumber | 09.4 §9; 10.4 §2.1; W-30; R-13; K-58 |

#### RISK-58 — Kesediaan membayar tidak diketahui

| | |
|---|---|
| Penilaian | Bisnis · K **Tinggi** · D **Tinggi** · **P1** · **Jalur publik** (§0) |
| Uraian dan sebab | Tidak ada bukti primer kesediaan membayar aplikasi keuangan rumah tangga di Indonesia (R-16; 18 §0.3). Rekomendasi freemium per rumah tangga (13 §7.5) bersandar pada pola harga pesaing |
| Mitigasi | Pilot tidak dimonetisasi; inti gratis sehingga tahap publik dapat berjalan bila biayanya ditanggung (13 §6.3, §7.5); harga uji di wawancara (preferensi yang dinyatakan, bukti lemah); beta; harga dan isi Plus diputuskan pemilik produk di GR-5 (Q-216) |
| Pemilik mitigasi | Pemilik produk; `product-analyst`; `researcher` |
| Harus ada pada | GR-5 |
| Tanda awal | Jawaban pertanyaan 10 wawancara 13 §6.3 |
| Risiko sisa | K Sedang, D Tinggi sampai ada pembelian nyata sesudah GR-5 |
| Sumber | R-16; 13 §7; Q-36, Q-216 |

#### RISK-59 — Komisi toko, pembelian dalam aplikasi, dan hak lintas platform

| | |
|---|---|
| Penilaian | Bisnis · K **Rendah** · D **Tinggi** · **P3** (dinilai ulang di S5 fase D sesudah R-15; semula K Sedang, P2) · **Jalur publik** (§0) |
| Uraian dan sebab | Rumah tangga pilot sendiri memakai iPhone dan Android, jadi langganan per rumah tangga lintas platform adalah kasus biasa. Hasil R-15 (18 §2.13) **[Terverifikasi]**: komisi langganan sekitar 15% di kedua toko bagi pengembang kecil yang mendaftar (Apple Small Business Program; tingkat 15% Google); langganan wajib lewat pembelian dalam aplikasi di kedua toko; PPN pelanggan Indonesia dipungut dan disetor toko; Apple 3.1.3(b) mengizinkan hak yang dibeli di platform lain asalkan Plus juga dapat dibeli di aplikasi iOS. **Yang tersisa:** aturan Google tentang hak lintas platform tidak ditemukan **[Perlu validasi]**; tingkat harga rupiah belum dibaca **[Perlu validasi]** |
| Mitigasi | Hak akses disimpan server per rumah tangga, bukan dibaca dari toko perangkat (13 §7.5; 14 §8); Plus dapat dibeli di dalam aplikasi di kedua platform; Family Sharing Apple tidak diaktifkan untuk Plus **[Usulan]** `researcher` (18 §2.13.3 butir 3); rekomendasi 13 menyebut syarat pembatalannya bila aturan Google tidak mengizinkan |
| Pemilik mitigasi | `researcher` (sisa R-15); `architect` (14) |
| Harus ada pada | FS-3, sebelum GR-5; tarif dan aturan diperiksa ulang sebelum Plus dijual (18 §2.13) |
| Tanda awal | Kebijakan Play yang melarang hak dari pembelian di luar Play |
| Risiko sisa | K Rendah, D Sedang |
| Sumber | R-15; 18 §2.13; 13 §7.5; 14 §8 |

#### RISK-60 — Keluarga lain berbeda dari rumah tangga pilot

| | |
|---|---|
| Penilaian | Bisnis · K **Sedang** · D **Tinggi** · **P2** · **Jalur publik** (§0) |
| Uraian dan sebab | Custody tunai, kiriman ke orang tua, dan persetujuan antarpasangan adalah pola rumah tangga pilot; keluarga lain mungkin hampir tanpa tunai, satu rekening bersama, atau menganggap persetujuan mengekang (R-08) |
| Mitigasi | Wawancara 8 sampai 12 rumah tangga dengan tabel "hasil yang mengubah rencana" (13 §6.3; A-41); beta; segmen rilis pertama yang sempit (Q-214) |
| Pemilik mitigasi | `product-analyst` |
| Harus ada pada | GR-5 |
| Tanda awal | Tiga wawancara pertama |
| Risiko sisa | K Sedang, D Sedang |
| Sumber | R-08; A-41; 13 §6.3, §3.7 |

#### RISK-61 — Pembeda belum teruji

| | |
|---|---|
| Penilaian | Bisnis · K **Sedang** · D **Sedang** · **P3** · **Jalur publik** (§0) |
| Uraian dan sebab | Pembeda yang diusulkan (custody, permintaan dua arah sampai uangnya, arus uang keluarga besar) belum diuji terhadap pesaing (A-15; R-02); produk bank seperti Bank Jago Kantong Bersama belum diteliti (R-09) |
| Mitigasi | Uji langsung dua atau tiga aplikasi pesaing sebelum materi publik (13 §3.1; R-02, memerlukan pemilik produk); materi publik tanpa positioning yang belum diuji (13 §3.6 GR-5) |
| Pemilik mitigasi | `product-analyst`; pemilik produk |
| Harus ada pada | GR-5 |
| Tanda awal | — |
| Risiko sisa | K Sedang, D Rendah |
| Sumber | A-15; R-02, R-09; 18 §1.4 |

#### RISK-62 — Biaya operasi dan waktu pengembang di tahap publik

| | |
|---|---|
| Penilaian | Bisnis · K **Sedang** · D **Sedang** · **P3** · **Jalur publik** (§0; waktu pengembang untuk keluarga sendiri dinilai di RISK-53) |
| Uraian dan sebab | Brief §14 "operating cost". Biaya infrastruktur pilot kecil (09.4 §7.1): Rp260.667 + USD 1,32 per bulan selama jalur web (batas atas USD 10), dan Rp260.667 + USD 9,57 (batas atas USD 18,25) sejak jalur iOS native dimulai, karena Apple Developer Program USD 99 per tahun baru diperlukan saat itu dan per rumah tangga di tahap publik diperkirakan kecil (A-42 **[Asumsi]**). Yang tidak diperkirakan dokumen mana pun: waktu pengembang, pembaruan OS tahunan untuk dua aplikasi native, dukungan pengguna, tinjauan hukum (13 §7.3) |
| Mitigasi | Satu VPS dengan pemicu penskalaan (09 §6); `AI` R1 tidak tumbuh dengan jumlah rumah tangga (09.4 §7.2); model bisnis di GR-5 |
| Pemilik mitigasi | Pemilik produk |
| Harus ada pada | GR-5 |
| Tanda awal | Waktu pemeliharaan per bulan selama pilot |
| Risiko sisa | K Sedang, D Sedang |
| Sumber | Brief §1, §14; 09.4 §7; 13 §7.3; A-42 |

#### RISK-63 — Langganan Plus menjadi alat kendali

| | |
|---|---|
| Penilaian | Bisnis · **Pagar** · K **Sedang** · D **Tinggi** · **P2** · **Jalur publik** (§0; Plus tidak dibangun untuk rumah tangga, 14 §8) |
| Uraian dan sebab | Pembayar dapat memakai status "aku yang bayar" sebagai wewenang, atau berhenti membayar untuk mencabut fitur yang diandalkan pasangannya. Brief §8 dan PRN-09, PRN-10 melarang wewenang satu arah |
| Mitigasi | Aturan langganan 13 §7.5 (a) sampai (d) (Q-217): Plus melekat pada rumah tangga, tanpa nama pembayar, pembayar tidak mendapat peran, izin, wewenang, atau pantauan, berakhir di akhir masa bayar tanpa catatan hilang, tidak lintas rumah tangga; daftar "tidak pernah berbayar" (13 §7.5). Kini persyaratan yang dapat diuji: 06 PG-15 dan SEC-156 (setiap anggota berakun membaca baris Plus yang sama; pembayar dan toko asal tidak terbaca; 14 §8) |
| Pemilik mitigasi | `architect` (06, 09); `security-privacy` (10); `product-analyst` |
| Harus ada pada | FS-3 |
| Tanda awal | — |
| Risiko sisa | K Rendah, D Sedang: pengaruh sosial di luar aplikasi |
| Sumber | 13 §7.1, §7.5; Q-217; PRN-09, PRN-10, PRN-13; 06 PG-15; SEC-156 |

#### RISK-64 — Deskriptor rating usia Apple 12 Januari 2027

| | |
|---|---|
| Penilaian | Bisnis · K **Sedang** · D **Sedang** · **P3** |
| Uraian dan sebab | Semua aplikasi wajib menjawab deskriptor baru sebelum kiriman pertama sesudah 16 November 2026, paling lambat 12 Januari 2027 (**[Terverifikasi]** 18 §2.12.3). Apakah itu berlaku bagi aplikasi yang hanya di TestFlight internal belum diketahui (R-18). Bila berlaku dan belum dijawab, unggahan build iOS dapat tertahan |
| Mitigasi | Kuesioner dijawab saat catatan aplikasi di App Store Connect dibuat (13 §8 **[Usulan]**); jawaban "Dating and Relationships" dan override 18+ (SEC-83) |
| Pemilik mitigasi | Pemilik produk (pemegang akun) |
| Harus ada pada | Jalur iOS native, sebelum unggahan pertama (14 §3.9 I1; semula FS-1). Klien web tidak melewati App Store Connect, jadi risiko ini tidak menyentuh GR-2 |
| Tanda awal | Peringatan di App Store Connect |
| Risiko sisa | K Rendah, D Rendah |
| Sumber | SEC-83; R-18; 13 §8 |

### 3.7 Hukum dan regulasi

Semua butir bukan nasihat hukum. Tafsirnya adalah pertanyaan advokat di 17 bagian 6.

#### RISK-04 — Letak pusat data dan kebijakan privasi penyedia hosting

| | |
|---|---|
| Penilaian | Hukum · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Kota pusat data penyedia usulan tidak dinyatakan, dan kebijakan privasinya menyebut data pribadi dapat diproses di luar Indonesia (09 §9 **[Terverifikasi]**; R-14). Syarat "server di Indonesia" (keputusan 10) tidak dapat dibuktikan tertulis |
| Mitigasi | Penyedia dipilih di S4-2 menjelang deploy; pernyataan tertulis penyedia yang dipilih, sebelum data nyata (10 §4.1 butir 7; ADR-020); catatan prosesor ditunda ke go public (keputusan 68; semula Q-208) |
| Pemilik mitigasi | Pemilik produk (S4-2); pengelola server |
| Harus ada pada | GR-2 |
| Tanda awal | Penyedia menolak memberi pernyataan tertulis |
| Risiko sisa | K Rendah, D Sedang; tafsir L-12, L-16 sebelum pengguna luar pertama (Q-209; keputusan 52) atau tahap publik, mana yang lebih dulu |
| Sumber | 09.4 §9; ADR-020; R-14; L-12, L-16 |

#### RISK-12 — Jalur yang bergantung pada tafsir hukum: xtool, transfer push dan `AI`

| | |
|---|---|
| Penilaian | Hukum · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Lisensi alat build iOS di Linux (L-31); apakah token push dan pemanggilan penyedia `AI` di luar negeri adalah transfer lintas batas (L-16, L-05, L-18, L-19, L-26, L-32). Menurut bacaan `researcher` **[Usulan]**, sesudah PP 33/2026 berlaku tidak ditemukan jalan bagi pengiriman harian ke penyedia `AI` luar negeri sampai Lembaga PDP menerbitkan daftar atau klausul (17 Q-10; R-11) |
| Mitigasi | Xcode Cloud sebagai jalur resmi, xtool tidak dipakai (ADR-022); jalur R1: tidak ada data rumah tangga yang dikirim ke penyedia `AI` (SEC-121; ADR-021); R2 terkunci sampai ADR baru sesudah advokat (SEC-122); payload push umum (SEC-50) |
| Pemilik mitigasi | Advokat; `security-privacy`; pemilik produk |
| Harus ada pada | GR-4 (L-16), sebelum pengguna luar pertama; FS-7 (R2) |
| Tanda awal | Terbitnya Peraturan Lembaga atau daftar negara setara |
| Risiko sisa | K Rendah, D Sedang untuk pilot. Bahwa R1 bukan transfer adalah bacaan `security-privacy` (L-32; W-29); token perangkat tetap lewat Apple dan Google (TM-18) |
| Sumber | 09.4 §9; L-05, L-16, L-18, L-19, L-26, L-31, L-32; R-05, R-11; W-29; TM-19 |

#### RISK-65 — Status regulasi layanan: PSE dan lingkup jasa keuangan

| | |
|---|---|
| Penilaian | Hukum · K **Rendah** · D **Kritis** · **P2** · **Diterima pemilik produk untuk pilot** (W-32). **Dinilai ulang di S6** atas jawaban advokat (keputusan 59; sumbernya dinyatakan pemilik produk sendiri); sebelumnya K Sedang · P1 |
| Uraian dan sebab | Permenkominfo 5/2020 mewajibkan PSE lingkup privat mendaftar sebelum dipakai; yang wajib tetapi tidak mendaftar dapat diputus aksesnya (18 §2.11 **[Terverifikasi]** naskah 2020; W-32; L-14). **Jawaban advokat atas L-14 dan L-38** (keputusan 59; pernyataan advokat, tidak diverifikasi lead): aplikasi tertutup untuk keluarga, termasuk dengan fitur chat, tidak wajib daftar PSE; baru wajib bila dibuka untuk umum. Atas L-01: orang tua yang diberi login tetap keluarga, dan pengecualian rumah tangga tetap berlaku selama aplikasi pribadi dan nonkomersial. Sisa: bacaan itu belum diuji di luar konsultasi, dan bila jalur publik diambil kewajiban PSE pasti timbul. Terpisah: apakah pengingat tagihan dengan tautan ke aplikasi pembayaran tetap di luar "payment initiation" (L-13, belum dijawab) |
| Mitigasi | Pilot dan pengguna luar di jalur keluarga: risiko diterima pemilik produk (S4-4 (iii) A; W-32), kini didukung jawaban advokat atas L-01, L-14, L-24 (dijawab lewat konsultasi advokat, keputusan 59), yang memenuhi syarat advokat GR-4 butir 1 (13 §3.6; Q-209). Bila jalur publik diambil: pendaftaran PSE (SEC-147, sejak S6 keluar dari QG-31 (a)) serta badan hukum dan konsultasi advokat ulang sebelum dibuka untuk umum (FS-3, sebelum GR-5). Syarat "nonkomersial dan tertutup untuk keluarga" dijaga dengan satu konsultasi advokat ulang sebelum go public, yang menggantikan empat pemicu usulan 10 §3.4 (pengguna non-keluarga, penagihan atau Plus, App Store tersembunyi, Play publik); sebelum itu tidak ada konsultasi lanjutan, dan L-04, L-20 tidak dikejar (keputusan 67, diputuskan pemilik produk sendiri); pendaftaran tertutup di server (SEC-158) menjaga "tertutup" secara teknis. Menurut 10 §3.4, W-32 ditutup; statusnya di 17 milik `product-analyst`. Lingkup uang: aplikasi tidak memegang atau memindahkan dana (brief §1) |
| Pemilik mitigasi | Pemilik produk; advokat |
| Harus ada pada | GR-4, sebelum pengguna luar pertama (jawaban advokat sudah ada); FS-3 sebelum GR-5, bila jalur publik diambil (pendaftaran PSE) |
| Tanda awal | Rencana membuka aplikasi bagi orang di luar keluarga; perubahan aturan PSE |
| Risiko sisa | Jalur keluarga, termasuk orang tua dan kerabat dewasa lain yang diberi login (keputusan 52, 63): K Rendah, D Kritis, diterima pemilik produk (W-32). Jalur publik sesudah pendaftaran PSE: K Rendah, D Sedang. P-51-1 (kapan advokat menjawab L-14) selesai bersama Q-224 (keputusan 59) |
| Sumber | W-32; L-01, L-13, L-14, L-24, L-38; R-07; Q-209; Q-224 (selesai); keputusan 59; 10.4 §3.2; 13 §3.6 |

#### RISK-66 — Kewajiban pengendali selama pilot

| | |
|---|---|
| Penilaian | Hukum · K **Rendah** · D **Tinggi** · **P3**. **Dinilai ulang di S6** atas jawaban advokat (keputusan 59); sebelumnya K Sedang · P2 |
| Uraian dan sebab | **Jawaban advokat** (keputusan 59; pernyataan advokat, tidak diverifikasi lead): pilot masuk pengecualian kegiatan pribadi atau rumah tangga (L-01; UU PDP Pasal 2 ayat (2) dan Putusan MK 108/PUU-XX/2022 menurut advokat), juga dengan server sewaan dan notifikasi Apple dan Google, selama aplikasi nonkomersial dan tertutup untuk keluarga; pengecualian tidak gugur karena pengelola servernya anggota, selama istri tahu dan setuju cara kerjanya dan admin tidak membuka data akun orang lain tanpa izin (L-24; aturan itu diterima atas saran lead, keputusan 60); data pihak luar yang dicatat (ibu, ayah, saudara) tidak wajib diberi tahu, selama tidak disebarkan atau disalahgunakan (L-03, L-17). Sisa: pengecualian gugur bila syaratnya tidak dijaga (komersial, terbuka untuk umum, admin membuka akun orang lain), dan L-11, L-21 belum dijawab. PP 33/2026 diperkirakan berlaku pertengahan Januari 2027 **[Perlu validasi]** |
| Mitigasi | Pilot berpegang pada pengecualian rumah tangga dan menjaga syaratnya: nonkomersial, tertutup untuk keluarga (SEC-158 pendaftaran tertutup), dan pengelola tidak membuka data akun orang lain tanpa izin (SEC-159; keputusan 60). Kontrol keamanan teknis kolom "Pilot" tetap. DPIA, catatan prosesor, kebijakan privasi, dan rencana pemberitahuan insiden (10 §4.1 butir 4 dan 5; SEC-139, SEC-148 (b), SEC-149, SEC-150) **tidak disiapkan untuk pilot** dan dikerjakan sebelum go public (keputusan 68, diputuskan pemilik produk sendiri; semula Q-208) |
| Pemilik mitigasi | Pemilik produk (sebagai pengendali); `security-privacy` |
| Harus ada pada | GR-2 (syarat pengecualian dijaga); sebelum go public (dokumen 10 §4.1 butir 4 dan 5) |
| Tanda awal | Rencana penjualan, pembukaan bagi orang di luar keluarga, atau pengelola membuka data akun orang lain; SEC-158 atau SEC-159 gagal uji |
| Risiko sisa | K Rendah, D Tinggi, diterima pemilik produk (keputusan 68): bila pengecualian ternyata tidak berlaku, dokumen kepatuhan belum ada dan disusun sesudahnya; prioritas tetap P3 karena K dan D penilaian tidak berubah |
| Sumber | L-01, L-03, L-11, L-17, L-21, L-24; keputusan 59, 60, 68; A-16; Q-208; 10 §3.1 |

#### RISK-67 — Catatan hanya-menambah lawan hak hapus

| | |
|---|---|
| Penilaian | Hukum · K **Sedang** · D **Tinggi** · **P2** · **Jalur publik** (§0; bagi rumah tangga, pilot berpegang pada pengecualian rumah tangga dengan syaratnya dijaga, RISK-66 dan keputusan 68, dan pemisahan inti dan rincian sudah dibangun di FS-1) |
| Uraian dan sebab | Hak hapus menurut UU PDP berhadapan dengan riwayat uang yang tidak dapat diubah (brief §7). Pseudonimisasi inti moneter dan penghapusan rincian adalah keputusan sementara sampai advokat menjawab (keputusan 24; ADR-007; L-04, L-20, L-23). Di rumah tangga dua orang, "mantan anggota" tetap jelas siapa orangnya (W-10). Cukup-tidaknya 38 hari di cadangan (L-34), hak ahli waris (L-35), dan lama menyimpan bukti persetujuan (L-36) juga belum dijawab |
| Mitigasi | Pemisahan inti, rincian, dan pemetaan identitas (06 §2.5; 07 §8.2) sehingga perubahan, bila diperlukan, mengenai bagian penghapusan dan bukan inti ledger; ADR-007 menyebut apa yang berubah bila jawabannya tidak |
| Pemilik mitigasi | Advokat; `finance-domain`; `architect` |
| Harus ada pada | FS-3, sebelum GR-5 |
| Tanda awal | Jawaban advokat atas L-04 (tidak dikejar sampai go public diputuskan; keputusan 67) |
| Risiko sisa | K Sedang, D Sedang sampai dijawab |
| Sumber | Brief §7; ADR-007; L-04, L-20, L-23, L-34, L-35, L-36; W-10 |

#### RISK-68 — Anak: indikator Kepmen 142/2026 dan data anak

| | |
|---|---|
| Penilaian | Hukum · **Pagar** · K **Sedang** · D **Tinggi** · **P2** |
| Uraian dan sebab | Satu indikator Kepmen Komdigi 142/2026 cukup untuk menggolongkan aplikasi "mungkin digunakan atau diakses oleh anak" (18 §2.9 **[Terverifikasi]**). Rumah tangga pilot punya anak 12+ (Q-186), yang dapat memakai ponsel orang tua (L-29); data anak dicatat orang dewasa, kadang bukan orang tua atau walinya (L-30) |
| Mitigasi | Syarat layanan melarang di bawah 18 dan pernyataan usia tanpa tanggal lahir (SEC-80); tanpa akun anak (SEC-81); desain sasaran dewasa dengan catatan bukti (SEC-82); rating 18+ (SEC-83); penilaian mandiri indikator (SEC-84); minimisasi data anak (SEC-85); kunci aplikasi (RISK-41) |
| Pemilik mitigasi | `security-privacy`; `ux-designer`; pemilik produk |
| Harus ada pada | GR-2 (SEC-80 sampai SEC-82, SEC-85); GR-4: SEC-84 sebelum pengguna luar pertama, SEC-83 (rating toko) bila jalur publik diambil |
| Tanda awal | — |
| Risiko sisa | K Sedang, D Sedang sampai L-27 sampai L-30 dijawab |
| Sumber | 10.3 §1; L-27 sampai L-30; Q-186; D-109 |

#### RISK-69 — Konten pengguna dan permintaan aparat

| | |
|---|---|
| Penilaian | Hukum · K **Sedang** · D **Sedang** · **P3** · **Jalur publik** (S6; semula "pengguna luar" menurut keputusan 52). Menurut jawaban advokat atas L-14 dan L-38 (keputusan 59), aplikasi tertutup untuk keluarga, termasuk chat dan orang tua yang diberi login, bukan layanan publik dan tidak wajib PSE; kewajiban PSE dan UGC baru timbul bila dibuka untuk umum. K dan D tetap dinilai seolah jalur publik diambil (§0) |
| Uraian dan sebab | Bila chat tertutup menjadikan layanan "PSE Lingkup Privat User Generated Content", ada kewajiban sarana pelaporan publik dan pemutusan akses konten dalam jam atas chat yang tidak dipindai server (L-38). Permintaan data oleh aparat (TM-24; L-14). Apakah aplikasi keluarga tertutup termasuk "layanan jejaring dan media sosial" (L-07) |
| Mitigasi | "Laporkan" (SEC-154); rekam jejak akses aparat (SEC-39 (d)); minimisasi data; L-38 dijawab advokat untuk jalur keluarga (keputusan 59); L-07 dan kewajiban UGC di jalur publik belum dijawab |
| Pemilik mitigasi | Advokat; `security-privacy` |
| Harus ada pada | FS-3, sebelum GR-5, bila jalur publik diambil |
| Tanda awal | — |
| Risiko sisa | K Sedang, D Sedang sampai dijawab; berkaitan dengan ADR E2EE |
| Sumber | TM-24; L-07, L-14, L-38; 10.4 §3.4 |

#### RISK-70 — Regulasi dibaca dari salinan teks; Lembaga PDP belum ada

| | |
|---|---|
| Penilaian | Hukum · K **Sedang** · D **Sedang** · **P3** · **Pengguna luar** (§0; keputusan 52; bagi rumah tangga, rancangan sudah memilih bacaan yang lebih ketat) |
| Uraian dan sebab | Banyak ketentuan PP 33/2026 dibaca dari salinan teks yang belum dicocokkan dengan pindaian resmi, dan memuat salah baca serta ayat terpotong (R-01). Lembaga PDP belum terbentuk menurut berita (R-11) |
| Mitigasi | Label **[Perlu validasi] — salinan teks** di setiap kutipan; pencocokan oleh pemilik produk atau advokat; rancangan memilih bacaan yang lebih ketat bila ragu (10 §3.1) |
| Pemilik mitigasi | `researcher`; pemilik produk |
| Harus ada pada | GR-4 |
| Tanda awal | Pindaian resmi tersedia |
| Risiko sisa | K Sedang, D Rendah |
| Sumber | R-01, R-11; 18 §2.6, §2.7 |

#### RISK-71 — Kedudukan catatan dalam perkawinan dan perpisahan

| | |
|---|---|
| Penilaian | Hukum · K **Rendah** · D **Tinggi** · **P3** |
| Uraian dan sebab | Menyebut sebagian penghasilan "dana pribadi" atau menyepakati jatah dapat dibaca sebagai perjanjian tentang harta (L-25). **Jawaban advokat atas L-25** (keputusan 59; pernyataan advokat): tidak; itu hanya pengaturan anggaran, penghasilan selama perkawinan tetap harta bersama, dan pemisahan harta hanya lewat perjanjian perkawinan di notaris. Penilaian K Rendah tetap, kini dengan dasar itu. Yang belum dijawab: catatan aplikasi dan ekspor chat dapat dipakai sebagai bukti dalam perpisahan (L-22, L-33) |
| Mitigasi | Kata antarmuka "disepakati sebagai", bukan "milik" (SEC-88); ekspor daftar klaim sebelum penutupan |
| Pemilik mitigasi | Advokat; `ux-designer` |
| Harus ada pada | FS-1 (kata layar); GR-5 (jawaban advokat), bila jalur publik diambil |
| Tanda awal | — |
| Risiko sisa | K Rendah, D Sedang |
| Sumber | L-22, L-25, L-33; SEC-88 |

#### RISK-72 — Kewajiban menjual langganan

| | |
|---|---|
| Penilaian | Hukum · K **Sedang** · D **Sedang** · **P3** · **Jalur publik** (§0) |
| Uraian dan sebab | Perlindungan konsumen (perpanjangan otomatis, pembatalan, pengembalian dana), pajak pengembang, dan apakah penagihan lewat toko mengubah kedudukan PSE belum diteliti (L-40). Sejak jawaban advokat (keputusan 59): pengecualian rumah tangga berlaku "selama aplikasi nonkomersial", jadi penjualan apa pun, termasuk Plus, dapat menggugurkannya |
| Mitigasi | L-40 sebelum Plus dijual (13 §7.5); Plus dan penagihan apa pun hanya sesudah konsultasi advokat sebelum go public, satu-satunya pemicu konsultasi ulang (keputusan 67, diputuskan pemilik produk sendiri; 14 §8) |
| Pemilik mitigasi | Advokat |
| Harus ada pada | GR-5 |
| Tanda awal | — |
| Risiko sisa | K Sedang, D Sedang sampai dijawab |
| Sumber | L-40; 13 §7.5 |

---

## 4. Risiko Tier 2 (ringkas)

Berlaku bila fiturnya dibangun di FS-2 (13 §3.4). Mitigasi yang disebut adalah persyaratan "Must bersyarat" yang ikut gelombangnya.

| ID | Kategori | Risiko dan sebab | K · D · P | Mitigasi | Pemilik · harus ada | Tanda awal · risiko sisa | Sumber |
|---|---|---|---|---|---|---|---|
| RISK-73 | Keamanan dan privasi · **Pagar** | Tekanan untuk membagikan status kepedulian; status yang dibagikan bertanda dibaca; balasan dukungan memuat keadaan kesehatan dan tetap ada sesudah status dihapus | Sedang · Tinggi · P2 | Pribadi sampai pemilik memilih penerima; cabut seketika dan senyap (V-7); "belum diisi" tidak dapat dibedakan dari "tidak dibagikan"; tanpa riwayat suasana hati; layar sebelum menghapus status (SEC-62, SEC-68) | `security-privacy`; `ux-designer` · T2-1 | VL-08 · K Rendah, D Sedang; balasan pasangan tetap ada (W-23; L-37) | TM-20; W-20, W-23; L-02, L-37; 10.3 §2.7; D-108 |
| RISK-74 | Produk | Saran `AI` terasa palsu, berulang, atau saling menggemakan bila keduanya memakainya; mutu kumpulan kalimat R1 belum diuji | Tinggi · Sedang · P2 | Masuk kolom tulis dulu, bervariasi, tanpa hitungan, dapat dimatikan tanpa diketahui pasangan; saran yang sama tidak diulang 30 hari (A-38) | `product-analyst`; `security-privacy` · T2-2 | VL-06 · K Sedang, D Rendah; aplikasi tetap utuh tanpa `AI` (PRN-14) | W-21, W-29; SCN-40 |
| RISK-75 | Teknis | Penyedia `AI` tidak tersedia atau batas biaya tercapai | Sedang · Rendah · P3 | Batas USD 3 per pembangunan kumpulan **[Asumsi]** dan USD 10 per bulan (SEC-127); saklar mati (SEC-128); kalimat baku; `AI` gagal tidak pernah menghalangi apa pun (PG-12) | `backend-dev` · T2-2 | Penghitung `ai_usage` · K Rendah, D Rendah | SEC-127, SEC-128; SCN-41; 09.3 §11 |
| RISK-76 | Keamanan dan privasi · **Pagar** | Keluaran `AI` membuka kepada satu orang apa yang ditutup orang lain, mendiagnosis, menyuruh mengawasi, atau membandingkan pasangan; penyedia menyimpan muatan | Rendah · Kritis · P2 | R1 tanpa data rumah tangga (SEC-121); judul tidak pernah dikirim (SEC-123); pemilih di server lewat jalur baca (SEC-124); tanpa perbandingan, dijaga struktural (SEC-129); daftar kalimat terlarang diuji otomatis (SEC-130) | `security-privacy`; `backend-dev` · T2-2 | Uji muatan dan keluaran · K Rendah, D Sedang | TM-19, TM-20, TM-21; D-110 |
| RISK-77 | Keamanan dan privasi · **Pagar** | Permintaan tindakan dipakai menagih pasangan; tugas yang menunggu menjadi alat tekanan | Sedang · Sedang · P3 | Penerima menerima atau menolak; pemohon membatalkan; pengingat berbatas PG-06; tanpa hitungan per orang | `architect`; `ux-designer` · T2-1 | SCN-30, VL-08 · K Rendah, D Sedang | Brief §5D; TM-28; F-REQ-013 |
| RISK-78 | Keamanan dan privasi · **Pagar** | Sebutan anggota terbaca dari layar pembuat, ditukar untuk menyesatkan pilihan penerima, dibaca pengelola server, dipakai sebagai catatan pengawasan, atau keberadaannya tersimpulkan | Rendah · Sedang · P3 | SEC-155 (a) sampai (i): hanya pembuat; layar penerima data atau wewenang selalu menampilkan nama akun; 30 karakter tanpa riwayat; dihapus pada F07 | `security-privacy`; `backend-dev` · T2-1 | Uji SEC-155 · K Rendah, D Rendah; sebagai data pribadi orang yang disebut: L-39 | TM-32 sampai TM-36; L-39 |
| RISK-79 | Produk | Fitur Tier 2 lain (`SHP`, `MNT`, `CAL`, `PLN`) tidak dipakai dan menambah beban navigasi | Sedang · Rendah · P3 | Dilepas bergelombang; dikeluarkan dari cakupan publik bila enam minggu tidak dipakai (13 §6.2; PRN-02) | `product-analyst` · T2-3 | Jawaban masing-masing pasangan, ditambah ada tidaknya objek yang terlihat keduanya, hanya untuk `SHP`, `MNT`, `PLN`, `CAL`, dan F-REQ-013; status `CARE` yang tidak dibagikan, sebutan anggota (F-FAM-021), dan pemakaian `AI` hanya dari jawaban penggunanya sendiri, tidak pernah dari data server (13 §6.2) · K Sedang, D Rendah | 13 §3.4, §6.2 |
| RISK-80 | Keamanan dan privasi · **Pagar** | Status berat, tanda menyakiti diri, atau tindak lanjut sakit ditangani salah: bahasa diagnosis, arahan yang terlambat, atau pemindaian yang mengawasi | Rendah · Kritis · P2 | Arahan bantuan hanya kepada pemilik status, lewat halaman Bantuan statis, tanpa pemindaian teks (SEC-133 sampai SEC-136); tanpa bahasa diagnosis (brief §5K) | `security-privacy`; `ux-designer` · T2-1 | Tinjauan kata layar · K Rendah, D Tinggi; aplikasi bukan layanan darurat | Brief §5K, §5O; 10.4 §1.5 sampai §1.7; RISK-46 |

---

## 5. Risiko Tier 3 (satu baris)

Dinilai saat fiturnya diangkat; tidak diberi K, D, atau P sekarang, kecuali RISK-83 (brief §14).

- **RISK-81** — Keamanan dan privasi · satu orang di beberapa rumah tangga membocorkan data antar-rumah tangganya (SEC-38; TM-22) · FS-5.
- **RISK-82** — Hukum · akun anak dan profil tanggungan memicu PP 17/2025 dan Permen 9/2026 (L-08 sampai L-10; R-10) · FS-5.
- **RISK-83** — UX · anggota lanjut usia tidak nyaman memakai aplikasi (brief §14 "low adoption by older members") · FS-5. **Dinilai sekarang, satu-satunya di bagian ini**, karena brief §14 menamainya dan meminta kemungkinan, dampak, pencegahan, dan mitigasinya: K **Sedang** · D **Sedang** · **P3** **[Usulan]**. *Mengapa penting:* di rumah tangga yang menopang orang tua (13 §6.4 sebaran beta; brief §4.1), anggota lanjut usia sering memegang uang belanja atau menerima kiriman; bila ia tidak mengonfirmasi, custody tetap "belum dikonfirmasi" dan KS-1 melemah bagi rumah tangga itu. *Pencegahan (sudah Tier 1):* teks sampai 200%, sasaran sentuh 48 pt/dp, pembaca layar, jalan non-gestur, kurangi gerak (08 §15); kata tanpa istilah akuntansi (brief §5C; 07 §1.4); alur PSN-04 mengonfirmasi uang yang dipegang dengan dua tombol besar, urungkan 5 detik, tanpa batas waktu menjawab (08 §8.4). *Mitigasi:* **di jalur rumah tangga** (S6; keputusan 52): orang tua dapat diberi login tanpa tahap publik, lewat GR-4 pintu (a) dan FS-5 (13 §3.6, §3.8), jadi risiko ini dapat datang lebih dulu di rumah tangga sendiri; sebelum login pertama orang tua, pencegahan Tier 1 di atas dicoba bersama orang tua itu di perangkatnya, berdampingan dengan QG-31 (a) **[Usulan]** `architect`; **bila jalur publik diambil:** beta mencari rumah tangga yang menopang orang tua (13 §6.4), dan "orang tua serumah" yang muncul berulang menaikkan urutan FS-5 (13 §6.4). *Pemilik:* pemilik produk (rumah tangga sendiri); `ux-designer`; `product-analyst` · GR-4 pintu (a) bila orang tua diberi login; GR-5 bila jalur publik diambil. *Tanda awal:* baris "belum dikonfirmasi" yang menumpuk di tempat uang yang dipegang anggota lanjut usia, di rumah tangga sendiri atau di rumah tangga beta (dibaca dari jawaban, bukan analitik). *Risiko sisa:* K Rendah, D Sedang.
- **RISK-84** — Produk · rumah tangga lebih dari dua dewasa macet karena kesepakatan `all` (S4-8 (i) A; 07 §10.3 butir 12 (a)) · FS-5.
- **RISK-85** — Teknis · impor mutasi bank tanpa jalur resmi bagi non-PJP (18 §3.1; L-15; R-04) · FS-8.
- **RISK-86** — Teknis · integrasi belanja daring tanpa API pembeli dan sinkron kalender yang memerlukan verifikasi scope (18 §3.2, §3.3) · FS-8.
- **RISK-87** — Hukum · `AI` R2, percakapan saran, dan entri berbantuan `AI` memerlukan dasar transfer lintas batas (L-05, L-16, L-18, L-19, L-32) · FS-7.
- **RISK-88** — **Ditutup** (web lebih dulu, §9). Semula: Teknis · klien web atau PC tanpa model keamanan browser di dokumen 10 (Q-210; R-17) · kandidat sesudah FS-4. Model keamanan browser kini ada di 10.5 (SEC-160 sampai SEC-179) dan klien web dirancang di ADR-030; risikonya dinilai ulang sebagai RISK-94 sampai RISK-97 berkedalaman penuh. Pengenal tidak dipakai ulang.
- **RISK-89** — Produk · `INV` dan `MEAL` menuntut pembaruan harian yang tidak akan dilakukan keluarga (brief §5F, §5J) · FS-6.
- **RISK-90** — UX · rancangan khusus tablet dan perangkat lipat tertunda (Q-120); tata letak Tier 1 tetap menyesuaikan lebar apa pun (brief §1.1) · FS-6.

---

## 6. Penyalahgunaan antaranggota dan kebenaran uang: pemeriksaan silang

Tujuan bagian ini: memastikan setiap ancaman dari pelaku di dalam rumah tangga (PL-02 anggota dewasa sebagai lawan, PL-03 orang yang memegang ponsel, PL-05 pengelola server) dan setiap invarian uang punya risiko, mitigasi yang dapat diuji, dan risiko sisa yang dinyatakan.

### 6.1 Pelaku di dalam rumah tangga

| Pelaku | Ancaman 10 | Risiko | Mitigasi yang berupa uji | Risiko sisa yang dinyatakan |
|---|---|---|---|---|
| PL-02 | TM-06, TM-07, TM-13, TM-36 | RISK-38 | SEC-86, SEC-89, urutan kursor identik (SEC-04) | B1 sampai B6 (07 §5.3) |
| PL-02 | TM-08 | RISK-34 | INV-A45 | Kehadiran seperti WhatsApp, diterima pemilik produk (W-18) |
| PL-02 | TM-09 | RISK-35, RISK-36 | PG-01 sampai PG-14; INV-A48 | Paksaan di luar aplikasi; ambang yang ditahan (W-34) |
| PL-02 | TM-11 | RISK-40 | SEC-17 | Pengaturan ulang lewat email yang dikuasai (sampai SEC-25) |
| PL-02 | TM-28 | RISK-37, RISK-77 | PG-06 | Pengajuan ulang tanpa batas (W-8) |
| PL-02 | TM-29, TM-30 | RISK-31, RISK-51 | INV-33, INV-37 | Selisih melekat pada pemegang |
| PL-02, PL-03 | TM-10 | RISK-35, RISK-41, RISK-14 | SEC-13, SEC-19 | Ponsel yang terbuka |
| PL-03 | TM-14, TM-15, TM-26 | RISK-41, RISK-68 | SEC-19 sampai SEC-21, SEC-46, SEC-50 sampai SEC-54 | Kunci aplikasi mati secara bawaan |
| PL-05 | TM-12, TM-31, TM-34 | RISK-11 | SEC-33, SEC-34 (pengakuan berversi) | Seluruhnya, diterima untuk pilot (keputusan 25) |
| PL-05 lewat klien web | TM-42 | RISK-95 | SEC-169 (a) hash artefak lawan berkas yang disajikan; kalimat [Klien web] SEC-178 (b) | Pengelola yang berniat tidak tertahan; diterima di server pasangan (bawaan 69.7, di bawah keputusan 25) |
| PL-02 lewat `AI` | TM-20, TM-21 | RISK-76 | SEC-124, SEC-129, SEC-130 | — |
| PL-02, PL-03, PL-05 | TM-32 sampai TM-36 | RISK-78 | SEC-155 | Layar pembuat yang terbuka; pengelola pilot |
| Pembayar Plus | SEC-156 | RISK-63 | 06 PG-15; SEC-156 (permutasi pembayar; tanggapan baca identik byte demi byte bagi setiap anggota) | Pengaruh sosial |

Skenario paksaan 10.3 §2.7 (sepuluh baris) seluruhnya tercakup RISK-34, RISK-35, RISK-36, RISK-37, RISK-39, RISK-40, RISK-41, dan RISK-73.

### 6.2 Kebenaran uang

| Ancaman terhadap kebenaran | Risiko | Uji yang memblokir merge (ES §5) |
|---|---|---|
| Aturan salah diimplementasikan | RISK-47 | Semua INV-01 sampai INV-41 dan INV-A01 sampai INV-A53 sebagai property test; semua CTH sebagai uji tabel |
| Kiriman ganda, bersamaan, luring | RISK-48 | INV-10, INV-23; SEC-151 |
| Riwayat diubah | RISK-47; RISK-11 (pengelola) | INV-09 hash inti moneter; uji arsitektur tanpa `UPDATE`/`DELETE` atas tabel inti uang (09.4 §4) |
| Pemindahan terhitung pemasukan atau pengeluaran | RISK-47 | INV-06 |
| Persetujuan mengurangi uang | RISK-47 | INV-12 |
| Catatan lengkap tetapi jenisnya salah | RISK-49 | Tidak dapat diuji mesin; VL-04 |
| Catatan tidak lengkap | RISK-20 | Tidak dapat diuji mesin; VL-01, rekonsiliasi kas |
| Ekspor dibaca salah | RISK-09 | INV-39 |

Dua baris terakhir sebelum baris ekspor adalah batas yang jujur: invarian menjamin aplikasi menghitung benar atas apa yang dicatat, bukan bahwa yang dicatat benar dan lengkap. Itulah sebab RISK-20 dan RISK-49 P1 dan P2 walaupun semua uji lulus.

---

## 7. Mengurangi beban pikir dan usaha berulang

Brief §14 meminta cara konkret agar aplikasi menghemat waktu, bukan menjadi tugas baru. Semua di bawah sudah ada di dokumen yang dirujuk; dokumen ini hanya menautkannya ke risiko.

| Beban | Mekanisme yang sudah dirancang | Rujukan | Risiko |
|---|---|---|---|
| Mencatat setiap pengeluaran | Satu tombol Catat; ≤ 3 ketukan; isian bawaan dari entri terakhir; simpan lokal seketika | ES §3; 08; IDEA-02 | RISK-19 |
| Lupa mencatat | Entri terlambat diterima dengan tanggal kejadian; pengingat tagihan dan pos rutin | 07 §3.11; IDEA-03 | RISK-19, RISK-20 |
| Mencocokkan uang | Rekonsiliasi kas sekali hitung; selisih tampil, tidak harus dijelaskan | 07 §3.10 | RISK-20 |
| Menyetujui | Satu ketukan dari kartu chat; dua dari notifikasi; belanja biasa tanpa persetujuan | 08; A-20 | RISK-21 |
| Mengatur aturan | Preset; ambang dan jatah boleh tidak ada; diatur kapan saja lewat kesepakatan | Brief §4.3; 07 INV-40 | RISK-24 |
| Membaca keadaan | Angka utama dan custody di Beranda; ringkasan bulanan otomatis | 07 §5, §6.3; 08 | RISK-18, RISK-20 |
| Notifikasi | Jam tenang, ringkasan, plafon harian, lipat | 06 §10, §11 | RISK-22 |
| Memikirkan siapa yang salah | Tanpa papan skor; tanda tanpa notifikasi langsung | PG-09; 07 §10.3 butir 10 | RISK-31 |
| Fitur yang tidak dipakai | Tier 2 bergelombang; dipangkas dari cakupan publik di GR-3, bila jalur publik diambil | 13 §3.4, §6.2 | RISK-23, RISK-79 |

---

## 8. Yang harus ada di setiap gerbang

Masukan bagi quality gate dokumen 14. Kolom kanan menyebut risiko yang mitigasinya harus sudah berlaku saat gerbang dilewati. Baris GR-3, GR-4 (b), dan FS-3 sampai FS-8 adalah **jalur publik opsional** (keputusan 51; §0): berlaku hanya bila pemilik produk kelak mengambilnya. Baris GR-4 (a) berlaku sebelum pengguna luar pertama, yang dapat terjadi tanpa tahap publik (keputusan 52). Tidak ada yang dihapus.

| Gerbang | Yang harus ada | Risiko |
|---|---|---|
| **GR-1** Mulai bangun | Urutan platform tercatat: **klien web lebih dulu** diputuskan pemilik produk sendiri (keputusan 69; ADR-030; menggantikan urutan keputusan 42 untuk FS-1 dan GR-2); sesudah GR-2 iOS native lalu Android (bawaan 69.5, diterima atas saran lead); stack web ADR-030 ditetapkan, dan bunyi ES §2, §3, §8 (ADR-030 butir 15 sampai 17) disetujui pemilik produk dan dipasang (13 November 2026; terpenuhi); **Mac tidak diperlukan sampai jalur iOS native** (ADR-022 catatan web; keputusan 43 tetap berlaku untuk jalur itu); urutan pekerjaan Tier 1 di 14. *Diperbaiki (temuan audit 5): "server uji sementara memenuhi 10 §4.1 butir 7" dan "versi dipatok (ES §8)" adalah syarat FS-1 langkah 1, bukan GR-1, dan kini di baris FS-1; jalur distribusi iOS keluarga (keputusan 55) pindah ke baris "Jalur iOS native".* | RISK-53, RISK-97 |
| **FS-1** Bangun Tier 1 | **Langkah 1:** server uji sementara memenuhi syarat 10 §4.1 butir 7 (keputusan 45; 14 §7.8); versi dipatok saat dipasang (ES §8); CSP dan header, sesi cookie, CSRF, lockfile beku dan pemindai `pnpm-lock.yaml`, daftar dependensi bundel tertutup (ADR-030; 14 §3.3). **Sepanjang FS-1:** semua uji invarian dan contoh memblokir merge; uji isolasi, privasi, log, payload; uji header, CSRF, XSS, dan antrean web dengan Playwright; uji kecepatan alur inti di klien web pada ketiga perangkat; gerbang contoh layar 08 §19.8 versi web. Mitigasi yang dikerjakan `ios-dev` atau `android-dev` jatuh di jalurnya, bukan di FS-1 (14 §3.8, §3.9) | RISK-02, RISK-09, RISK-19, RISK-30, RISK-33, RISK-38, RISK-43, RISK-47, RISK-48, RISK-49, RISK-55, RISK-57, RISK-94, RISK-96 |
| **GR-2** Data nyata (klien web) | Butir 10 §4.1 kecuali butir 3 (SEC-109 (b), bukan butir klien web, 10.5 SEC-178 (c); di jalur iOS native) dan butir 4 dan 5, yang tidak disiapkan untuk pilot (keputusan 68) (W-31: tanpa butir advokat); butir 7 termasuk server uji sementara dihapus (14 §7.8, QG-30 butir 8); butir 10 (iPad mini 7 yang mungkin dipegang orang lain) tercatat; butir 11 tinjauan hanya-temuan `security-privacy` (keputusan 49; 14 QG-30 butir 12); **butir 12** perangkat pilot di klien web (web app layar utama, `persist()`, izin notifikasi, browser desktop dan modenya; 14 QG-30 butir 16) dan **butir 13** origin produksi lulus pemeriksaan (header, CSRF, SSE, XSS, pemindai npm, hash bundel, bundel tanpa pemeriksaan diri; 14 QG-30 butir 17); uji klien web tanpa Mac lengkap (14 QG-30 butir 13); semua AC Tier 1 lulus, termasuk FR-FAM-013/AC-12 salinan swalayan (SEC-107); klien web di iPhone 13, iPad mini 7, dan Galaxy A55 lulus daftar uji tata letak web 08; ES §3 diukur di klien web pada ketiga perangkat (ADR-029 catatan web); runbook ada, termasuk deploy web, service worker pengganti, versi minimum, dan kunci VAPID; VL-00 dicatat | RISK-03 sampai RISK-06, RISK-10, RISK-11, RISK-18, RISK-24, RISK-35, RISK-38 sampai RISK-42, RISK-44 sampai RISK-46, RISK-54, RISK-66, RISK-68, RISK-92 (SEC-158), RISK-94 sampai RISK-97 |
| **Jalur iOS native** (sesudah GR-2; untuk keluarga; paling lambat sebelum TestFlight eksternal atau App Store, GR-4 pintu (b); ADR-030) | QG-30 (iOS) sebelum build iOS memakai data nyata (14 §3.9, §5.3): Apple Developer Program; sumber Mac sekali: AWS EC2 Mac atau Mac pinjaman (bawaan 69.6, diterima atas saran lead), pilihan di antara keduanya oleh pemilik produk saat jalur dimulai; Xcode Cloud sebagai pemeriksaan wajib pull request; pembungkus SQLite dinilai di minggu pertama; deskriptor rating usia sebelum unggahan pertama; SEC-109 (b) sebelum build TestFlight pertama (QG-25); pekerjaan kalender 60, 80, 90 hari (QG-23); jalur distribusi iOS keluarga: tetap TestFlight internal, App Store unlisted bila dihentikan (keputusan 55; ADR-022); daftar 10 §4.4; uji iOS tanpa Mac; ES §3 di aplikasi native; tinjauan hanya-temuan `security-privacy` sebelum iOS native memegang data nyata (bawaan 69.4, diterima atas saran lead; QG-30 (iOS) I-8). Bagian iOS dari mitigasi RISK-14, RISK-41, RISK-48 berlaku di sini | RISK-01, RISK-13, RISK-14, RISK-41, RISK-48, RISK-52, RISK-54, RISK-64, RISK-92, RISK-97 |
| **Jalur Android** (sesudah jalur iOS native; untuk keluarga; sebelum jalur pengujian Google Play dibuka, GR-4 butir 8, QG-31 (b); keputusan 42, 51; ADR-030) | QG-30 (Android) sebelum APK memakai data nyata: tinjauan hanya-temuan `security-privacy` (keputusan 53); daftar 10 §4.3, AC dan skenario di Galaxy A55, ES §3, vektor emas Kotlin, kunci APK (14 §3.8, §5.3). Bagian Android dari mitigasi RISK-14, RISK-41, RISK-48 berlaku di sini | RISK-14, RISK-17, RISK-41, RISK-48, RISK-54, RISK-56, RISK-91, RISK-97 |
| **FS-2** Pilot berjalan | VL-01, VL-03 sampai VL-10 (VL-02 hanya jalur publik opsional; 13 §6.1; PRN-11); sejak jalur iOS native, rilis pemeliharaan pada 60 hari tanpa rilis dan build iOS berikutnya paling lambat 80 hari sesudah unggahan sebelumnya (ADR-028 butir 4, keputusan 44); risiko Tier 2 per gelombang; pasangan opsional SEC-18 faktor kedua dan SEC-25 kode pemulihan per akun (keputusan 54); tinjauan hanya-temuan `security-privacy` tahunan (keputusan 53) | RISK-07, RISK-20, RISK-21, RISK-22, RISK-23, RISK-27, RISK-31, RISK-34, RISK-37, RISK-40, RISK-52, RISK-73 sampai RISK-80 |
| **GR-3** Evaluasi (jalur publik opsional; hanya bila pemilik produk kelak mempertimbangkan go public, keputusan 51) | Paling cepat tiga bulan sesudah GR-2; tabel keputusan 13 §6.2; VL-02 (a) kedua pasangan menyatakan nyaman dan aman, satu "belum" menahan; VL-02 (b) tanpa bug berat baru dalam 14 hari dari catatan bug (14 §5.3 masukan teknis GR-3); bug berat yang ditemukan lebih dari 14 hari sebelumnya tetapi belum diperbaiki juga menahan GR-3 (bacaan lead **[Usulan]** atas `critic` P-4, `_progress.md`); VL-08 tanpa temuan menghalangi; VL-10 nol (keputusan 46) | RISK-18, RISK-19, RISK-20, RISK-25, RISK-26, RISK-29, RISK-79 |
| **GR-4 (a)** Pengguna luar pertama, termasuk orang tua dan kerabat dewasa lain yang diberi login, tanpa rilis toko dan tanpa tahap publik (opsional; keputusan 52, 63) | 14 QG-31 (a): tinjauan hanya-temuan `security-privacy` sebelum pengguna luar pertama (keputusan 49); L-01, L-14, L-24 dijawab (sudah, lewat konsultasi advokat, keputusan 59, bagi pengguna luar dari keluarga; kutipan pasal adalah pernyataan advokat; Q-209); tanpa konsultasi advokat lanjutan, karena pemicu konsultasi ulang diringkas menjadi sebelum go public (keputusan 67); SEC-159 aturan pengelola server (keputusan 60); SEC-147 tidak di sini sejak S6 (wajib saat dibuka untuk umum, FS-3 dan GR-5); tabel 10 §4.2 kelompok (a): SEC-36 dengan log hanya-menambah di luar VPS dan satu baris tanpa isi di jejak keamanan setiap akun pada setiap akses darurat, SEC-39 (c), SEC-34 tiga versi dipilih per akun, SEC-18 dan SEC-25 (pasangan opsional yang dibangun di FS-2, dipastikan sudah lulus; keputusan 54; 13 §3.6 butir 7), SEC-84, SEC-109 (b) dan peran App Store Connect tersempit per orang baru bila ia dipasang lewat TestFlight (bagi klien web tidak berlaku, 10.5 SEC-178 (c)); QG-30 (Android) bila pengguna luar memakai aplikasi Android; email ke pengguna luar terbukti sampai. Kelompok (a+) bila pengguna luar membentuk rumah tangga sendiri: uji isolasi dengan banyak rumah tangga, SEC-154 (a), (b), (d) sampai (f), kontak pengelola (SEC-134) | RISK-06, RISK-11, RISK-12, RISK-40, RISK-44, RISK-65, RISK-68, RISK-70, RISK-93 |
| **GR-4 (b)** Beta (jalur publik opsional) | 14 QG-31 (b): QG-30 (Android) lulus dan APK terpasang sebelum jalur Play dibuka (13 GR-4 butir 8); SEC-153 (d) (atau di (a) bila pengguna luar dipasang lewat TestFlight eksternal), kelompok (a+), SEC-142, SEC-143; rating usia (SEC-83); SEC-139 versi publik; SEC-148 (b), SEC-149, SEC-150 dan 10 §4.1 butir 4 dan 5, dari GR-2 dan GR-4 (a) (keputusan 68); pengelola bukan anggota rumah tangga beta; uji isolasi dengan banyak rumah tangga; prasyarat beta 14 §6.7 | RISK-16, RISK-44, RISK-66, RISK-68, RISK-91 |
| **FS-3 dan GR-5** Rilis publik (jalur publik opsional) | Butir L "tahap publik"; badan hukum dan konsultasi advokat ulang sebelum dibuka untuk umum (jawaban advokat butir 7); SEC-147 pendaftaran PSE (keluar dari GR-4 (a) sejak S6); aturan langganan (SEC-156), hanya sesudah konsultasi advokat itu (keputusan 67; 14 §8); 10 §4.1 butir 4 dan 5 dan SEC-148 (b) bila belum lewat GR-4 (b) (keputusan 68); R-15, R-16; wawancara dan beta; tinjauan toko; SEC-18 (rumusan tahap publik: passkey menurut SEC-11; TOTP sudah di FS-2, keputusan 54), SEC-45; uji penetrasi pihak luar SEC-157 dengan temuan tinggi dan kritis diperbaiki (14 QG-32) | RISK-16, RISK-58 sampai RISK-63, RISK-65 (pendaftaran PSE), RISK-67, RISK-69, RISK-71, RISK-72 |
| **FS-4 sampai FS-8** | Dinilai saat diangkat, kecuali RISK-83 yang sudah dinilai (§5); pencegahannya Tier 1, mitigasinya di rumah tangga sendiri sebelum login pertama orang tua (GR-4 (a), FS-5) dan, bila jalur publik diambil, beta (GR-5) | RISK-81 sampai RISK-90 |

---

## 9. Risiko yang ditutup atau diterima

| ID | Keadaan | Dasar |
|---|---|---|
| RISK-15 | **Ditutup.** Pemilik produk menyetujui tegas pengecualian ES §2 untuk tiga bantuan isian (09-F A, checkpoint S4). Pengenal tidak dipakai ulang | 09.4 §9; 09 §7; ADR-013 |
| RISK-28 | **Ditutup.** Pemilik produk memutuskan urutan dan platform pilot (keputusan 42; urutannya kini digantikan web lebih dulu, keputusan 69); harga pilihannya dinilai di RISK-01, RISK-18, RISK-91, RISK-97. Pengenal tidak dipakai ulang | Keputusan 42; ADR-030; 17 K-66; Q-37, Q-211 |
| RISK-88 | **Ditutup** (web lebih dulu). Klien web kini Tier 1 dengan model keamanan 10.5 dan ADR-030; risikonya dinilai di RISK-94 sampai RISK-97. Pengenal tidak dipakai ulang | ADR-030; 10.5; Q-210 |
| RISK-01 | **Diterima** (risiko sisa: siklus UI iOS lambat, tanpa debugger di perangkat); ditinjau pemilik produk bila dana ada. Sejak ADR-030 berlaku mulai jalur iOS native, di luar jalur kritis GR-2 | Keputusan 43 (diterima atas saran lead); ADR-022 butir 4; W-28; ADR-030 |
| RISK-11 | **Diterima dengan syarat** untuk pilot | Keputusan 25; 10-F A; TM-12 |
| RISK-34 | **Diterima** (risiko sisa) | 10-A A, diputuskan pemilik produk sendiri; W-18 |
| RISK-36 | **Diterima** (risiko sisa) | 10-E A, S4-9 (i) A, diputuskan pemilik produk sendiri; W-34 |
| RISK-50, RISK-51 | **Diterima** lewat keputusan yang disebut 07 §10.3 | 07 §10.3 butir 1, 7, 8, 10, 11 |
| RISK-65 (pilot) | **Diterima** untuk pilot; sejak S6 berprioritas P2 atas jawaban advokat (keputusan 59) | S4-4 (iii) A; W-32 |
| RISK-92 | **Diterima** (risiko sisa: jeda tanpa aplikasi iOS bila Apple menghentikan TestFlight sebelum App Store unlisted siap; sejak ADR-030 klien web tetap tersedia selama jeda itu) | Keputusan 55 (Q-228; diterima atas saran lead); ADR-022 |
| RISK-95 | **Diterima** di server pasangan, di bawah keputusan 25 (risiko sisa: pengelola yang berniat dapat menyajikan bundel yang diubah bagi satu akun; pemeriksaan hash mencegah salah deploy, bukan niat; batasnya dinyatakan kepada istri sebelum data nyata, 10 §4.1 butir 2; SEC-159 tetap). Tidak berlaku bagi pengguna luar (RISK-93) | 17 bawaan 69.7 (diterima atas saran lead); 10 §8 butir 17; 10.5 SEC-178 (b) |
| RISK-97 | **Diterima** (risiko sisa: klien web tidak terasa native di ponsel sampai aplikasi native menggantikannya di perangkat itu) | `_progress.md` "Jawaban pemilik produk" (12 November 2026): "ya saya terima kan ini versi webnya, selebihnya saya ikut saran kamu"; ADR-030 |

Penerimaan ini tidak menutup kemungkinan pemilik produk membukanya kembali di checkpoint mana pun; dokumen ini tidak mengajukannya kembali.

---

## 10. Pemetaan sumber

### 10.1 Ancaman dokumen 10 ke risiko

| TM | RISK | TM | RISK | TM | RISK |
|---|---|---|---|---|---|
| TM-01 | 40 | TM-13 | 38 | TM-25 | 43 |
| TM-02 | 44 | TM-14 | 41 | TM-26 | 41 |
| TM-03 | 42 | TM-15 | 41 | TM-27 | 43 |
| TM-04 | 42 | TM-16 | 40 | TM-28 | 37, 77 |
| TM-05 | 42 | TM-17 | 45 | TM-29 | 31 |
| TM-06 | 38 | TM-18 | 43 | TM-30 | 51 |
| TM-07 | 38 | TM-19 | 12, 76 | TM-31 | 11, 93 |
| TM-08 | 34 | TM-20 | 76 | TM-32 | 78 |
| TM-09 | 35 | TM-21 | 76 | TM-33 | 78 |
| TM-10 | 14, 35, 41 | TM-22 | 44, 81 | TM-34 | 11, 78 |
| TM-11 | 40 | TM-23 | 45 | TM-35 | 78 |
| TM-12 | 11, 93 | TM-24 | 69 | TM-36 | 38, 78 |
| TM-37 | 96 | TM-40 | 41 | TM-43 | 42 |
| TM-38 | 96 | TM-41 | 41, 45 | TM-44 | 94 |
| TM-39 | 40 | TM-42 | 95, 11 | TM-45 | 96 |

### 10.2 Butir dokumen 17 ke risiko

| Jenis | Butir → RISK |
|---|---|
| W | W-1, W-2, W-4 → 50 · W-3, W-6, W-12, W-26 → 31 · W-5, W-13, W-14, W-17, W-25 → 51 · W-7, W-9 → 38 · W-8, W-11 → 37 (W-11 juga 22) · W-10 → 67 · W-18, W-19 → 34 · W-20, W-23 → 73 · W-21 → 74 · W-22 → 29 · W-24 → 39 · W-27 → 32 · W-28 → 01 · W-29 → 12, 74 · W-30 → 16 · W-31 → gerbang GR-2 (§8, butir 10 §4.1 tanpa butir advokat) dan 04 (pernyataan lokasi data) · W-32 → 65 · W-33 → 46 · W-34 → 36. W-15 dan W-16 bersyarat dan tidak berlaku |
| L | L-01, L-03, L-11, L-17, L-21, L-24 → 66 (L-24 juga 11) · L-02, L-37 → 73 · L-04, L-20, L-23, L-34, L-35, L-36 → 67 · L-05, L-06, L-16, L-18, L-19, L-26, L-31, L-32 → 12 (L-05, L-06, L-16, L-18, L-19, L-32 juga 87; L-06 menyangkut teks yang dide-identifikasi ke penyedia `AI`, yang di pilot tidak dikirim karena R1 tanpa data rumah tangga) · L-07, L-38 → 69 · L-08, L-09, L-10 → 82 · L-12 → 04, 66 · L-13, L-14 → 65 · L-15 → 85 · L-22, L-25, L-33 → 71 (L-22, L-33 juga 39) · L-27 sampai L-30 → 68 · L-39 → 78 · L-40 → 72 |
| R | R-01, R-11 → 70 · R-02, R-09 → 61 · R-03: angka Epstein dkk. tidak dikutip di sini karena masih diblokir (17 bagian 4) · R-04 → 85 · R-05 → 12 · R-06 → 16 · R-07 → 65 · R-08 → 60 · R-10 → 82 · R-12 → 46 · R-13 → 16 · R-14 → 04 · R-15 → 59 · R-16 → 58 · R-17 → 88 (ditutup; kini 94 sampai 97) · R-18 → 27, 64 |
| A | A-05 → 57 · A-16 → 66 · A-17, A-25 → 18, 25 · A-19, A-29 → 19 · A-20 → 21 · A-28, A-32, A-34 → 22 · A-37, A-40 → 27, 26 · A-38 → 74 · A-41 → 60 · A-42 → 62 |
| D (Pagar) | D-98, D-99 → 38 · D-102 → 41, 48 · D-104 → 45 · D-105, D-106 → 39 · D-107 → 14, 40, 41 · D-108 → 22, 73 · D-110 → 76 · D-112 → 21 · D-113 → 51 · D-114 → 41, 46 |
| K | K-58 → 16 · K-61 → 34 · K-62 → 06 |
| Q | Q-37, Q-211 → 28 (ditutup; keduanya terjawab keputusan 42), 18, 91 · Q-213 → 26 (dijawab keputusan 46) · Q-216 → 58 · Q-217 → 63 · Q-218 → 18, 25 · Q-210 → 88 (ditutup), 94, 95, 96 · Q-222 → 10, 11 (terbuka; sebelum GR-2) · Q-224 → 65 (selesai, keputusan 59: konsultasi advokat sudah terjadi) · Q-227 → 40 (dijawab keputusan 54) · Q-228 → 92 (dijawab keputusan 55) |

---

## 11. Permintaan kepada pemilik dokumen lain (lewat lead)

Dokumen ini tidak menyunting dokumen lain dan tidak membuat persyaratan baru.

| Kepada | Permintaan | Risiko |
|---|---|---|
| `architect` (09, 06, ADR; instans yang menulis 14) | (a) 09.4 §9: ganti tabel dengan penunjuk ke dokumen ini (pengenal dan judul satu baris, atau satu kalimat), agar tidak ada dua salinan; rujukan `RISK-NN` di ADR-008, ADR-012, ADR-013, ADR-020, ADR-022, ADR-023, ADR-026 tetap berlaku. (b) RISK-12 di 09.4 §9 berjudul sempit; dokumen ini memperluas uraiannya ke L-05, L-18, L-19, L-26, L-32 tanpa mengganti pengenal; selaraskan judulnya atau biarkan penunjuk. (c) RISK-04 berkategori Hukum di sini. (d) Q-217 (13 §12.2): baris 06 yang menyatakan Plus tidak memberi peran, izin, wewenang, atau pantauan dan tidak lintas rumah tangga (RISK-63). (e) 14: pakai §8 dokumen ini sebagai masukan quality gate per gerbang | 04, 11, 12, 63 |
| `security-privacy` (10) | (a) **Waktu butir "sebelum tahap publik" terhadap GR-4.** Rumah tangga beta (FS-3) memakai server yang sama dan pengelolanya bukan anggota mereka. Mohon dinyatakan butir mana yang wajib sebelum GR-4, bukan GR-5: setidaknya SEC-18 faktor kedua dan SEC-25 kode pemulihan (RISK-40), SEC-36 akses darurat. (b) Jangka verifikasi ulang nomor layanan di halaman Bantuan (Healing119, 110, SAPA 129 bila tampil), misalnya di setiap gerbang dan tahunan, sebagai bagian SEC-138 (RISK-46). (c) Kelas data penagihan dan aturan Q-217 sebagai persyaratan, sebelum FS-3 (sudah diminta 13; RISK-63) | 40, 44, 46, 63 |
| `product-analyst` (13, 17) | (a) 13 §6.5 dan VL-08: pertimbangkan saluran jawaban yang tidak melewati pemilik produk, misalnya jawaban tertulis yang disimpan istrinya sendiri dan dibacakan di GR-3 atas pilihannya (RISK-25). (b) Q-213: usulan `architect` agar dijawab sebelum GR-2, bukan sebelum GR-3 (RISK-26); **selesai**, dijawab pemilik produk lewat keputusan 46 sebelum GR-2. (c) Selaraskan 13 §3.4 "80 hari" dengan ADR-022 butir 1 ("paling lambat setiap 90 hari, sasaran setiap 60 hari"); keduanya tidak bertentangan, tetapi pembaca dapat mengira sasarannya 80 (RISK-52). (d) 17 R-03 menyebut dokumen 15 memerlukan angka Epstein dkk.; dokumen ini tidak mengutipnya, jadi kolom "Diperlukan oleh" dapat diperbarui. (e) 17: catat bahwa `RISK-NN` kanonis di dokumen 15. (f) Pertanyaan baru untuk pemilik produk bila lead menyetujuinya: bila pemilik produk berhalangan lama, siapa mengoperasikan server, mengingat memberi istri akses pengelola juga memberinya akses teknis ke data pribadi suaminya (RISK-10, cermin RISK-11) | 10, 25, 26, 52 |
| `researcher` (18) | Tidak ada permintaan baru. Angka RISK-01 sudah diperbarui dari 18 §5.2.7 (fase D) | 01 |
| Lead (status fase D) | **`architect`:** (a) 09.4 §9 kini penunjuk; (b) judul RISK-12 di penunjuk itu sama dengan di sini; (c) RISK-04 tetap Hukum di sini, 09.4 tidak lagi menilai; (d) Q-217: 06 PG-15 sudah menyatakannya (revisi kelima belas), kini juga SEC-156; (e) 14 §5.3 memakai §8 dokumen ini sebagai masukan QG-30 sampai QG-32. **`security-privacy`:** (a) dijawab 10 §4.2 (diterapkan di §8 GR-4, RISK-11, RISK-40, RISK-44); (b) SEC-138 setiap gerbang dan sekurangnya 6 bulan (RISK-46); (c) SEC-156. **`product-analyst`:** (b) Q-213 sebelum GR-2 **[Usulan]** di 13 dan 17; (c) menjadi Q-219 (batas 80 hari usulan; ADR-022 butir 1 berlaku); (f) menjadi Q-222. **Sesudah keputusan 43 sampai 50:** Q-213 dijawab (keputusan 46), Q-219 dijawab (keputusan 44, 80 hari berlaku), Q-222 tetap terbuka sampai sebelum GR-2 | — |
| `security-privacy` (10), sesudah keputusan 51 | (g) Sejak GR-4 dan GR-5 opsional, SEC-25 kode pemulihan dan SEC-18 faktor kedua tidak punya tenggat di jalur rumah tangga, jadi risiko sisa RISK-40 (penguasa email dapat mengatur ulang kata sandi pasangannya; TM-11, Pagar) berlaku tanpa batas waktu. Mohon dinyatakan apakah keduanya tetap terikat GR-4 dan GR-5, atau diberi tenggat di jalur rumah tangga (misalnya sebelum gelombang Tier 2 tertentu). Dokumen ini tidak mengusulkan jawabannya. **Diteruskan lead ke `security-privacy`.** Sejak keputusan 52, SEC-25 jatuh tempo sebelum pengguna luar pertama (QG-31 (a)). **Dijawab pemilik produk lewat keputusan 54** (Q-227): SEC-18 dan SEC-25 sepasang fitur opsional per akun, dibangun di FS-2, tidak menahan GR-2; diterapkan di RISK-40 | 40 |
| Lead, sesudah keputusan 51 | (h) RISK-11 (keputusan 25, "Diterima untuk pilot") dan RISK-65 (W-32, S4-4 (iii) A, untuk pilot) diterima pemilik produk dengan kata "pilot", yang saat itu berbatas waktu menuju GR-4. Sejak keputusan 51, pemakaian rumah tangga ini tidak berbatas waktu (glosarium *Rumah tangga pilot*). Mohon dinyatakan apakah kedua penerimaan dibaca berlaku untuk pemakaian itu, atau dibawa ke pemilik produk. **Dijawab lead:** RISK-11 tetap Diterima dan berlaku untuk pemakaian rumah tangga tanpa batas waktu (penilaian `security-privacy`, 10 §3.3); bagian PSE RISK-65 dibawa ke pertanyaan pemilik produk P-51-1, dan status tetap "Diterima untuk pilot" sampai ia menjawab. Diterapkan di kartu keduanya. **S6:** P-51-1 selesai bersama Q-224; jawaban advokat atas L-14 sudah ada (keputusan 59) | 11, 65 |
| `product-analyst` (17), web lebih dulu | (i) Pertanyaan untuk pemilik produk: apakah ia menerima RISK-95 (pengelola dapat menyajikan kode web yang diubah bagi istrinya, termasuk mengirim data yang hanya di perangkat) sebagai bagian penerimaan keputusan 25 di server pasangan, sebagaimana dibaca `security-privacy` di 10 §8 butir 17; bawaan **[Usulan]**: ya, dengan kalimat [Klien web] di pemberitahuan yang dibaca istri sendiri, dan status RISK-95 menjadi "Diterima dengan syarat". **Selesai:** diterima sebagai bawaan 69.7 atas saran lead (17); diterapkan di register, kartu, §6.1, §9. (ii) Keputusan 69 dan K atas brief §1 dan `CLAUDE.md` "Platform order" (kutipan di §0) | 95, 97 |
| Lead, web lebih dulu | (a) README: RISK-94 sampai RISK-97 baru, RISK-88 ditutup. (b) 16 dan 01: P2 kini 30, RISK-97 diterima, RISK-01 keluar dari jalur kritis. (c) Hitung ulang §12 butir 15 | 94 sampai 97 |
| Lead | (a) README baris `RISK-NN`: sebut "kanonis di dokumen 15". (b) Glosarium: istilah *risiko sisa* (residual risk), definisi kerja di §0. (c) Brief tugas S5 menyebut `AS-01` sampai `AS-14` sebagai skenario penyalahgunaan; di dokumen 10 itu aset, dan ancaman sudah sampai TM-36 (§0). (d) Hitung ulang §12 | — |

---

## 12. Angka yang harus dihitung ulang lead

| # | Angka | Letak |
|---|---|---|
| 1 | **Diubah (keputusan 42).** 91 pengenal (RISK-01 sampai RISK-91): 2 ditutup (RISK-15, RISK-28), 89 terbuka; 71 berkedalaman penuh (RISK-28 keluar, RISK-91 masuk), 8 Tier 2, 10 Tier 3 | §2, §4, §5 |
| 2 | **Diubah (keputusan 42).** Per kategori (terbuka): Produk 15, UX 6, Keamanan dan privasi 21, Keuangan 6, Teknis 21, Bisnis 8, Hukum 12; jumlah 89 | §2, §4, §5 |
| 3 | **Diubah.** Prioritas berkedalaman penuh: P1 13, P2 26, P3 32 (jumlah 71; RISK-59 dari P2 ke P3, K Rendah · D Tinggi); Tier 2: P2 4, P3 4 | §1.1, §2, §3.6 RISK-59, §4 |
| 4 | **Diubah (keputusan 42).** Setiap sel peta §1.1 cocok dengan K dan D di §2, dan prioritasnya cocok dengan matriks §0 (RISK-59 di sel Tinggi × Rendah; sel Tinggi × Sedang tetap 19 risiko: RISK-28 keluar, RISK-91 masuk). Prioritas tidak berubah: P1 13, P2 26, P3 32 | §0, §1.1, §2 |
| 5 | Tanda Pagar: 21 risiko (RISK-11, 14, 21, 31, 34 sampai 41, 46, 51, 63, 68, 73, 76, 77, 78, 80) | §1, §2, §4 |
| 6 | **Diubah (keputusan 44).** TestFlight kedaluwarsa 90 hari sejak unggahan; rilis pemeliharaan pada 60 hari; batas keras 80 < 90, sisa 10 (ADR-028 butir 4, diterima) | §3.5 RISK-52; §8 |
| 7 | **Diubah (ADR-030).** Biaya pilot dikutip dari 09.4 §7.1: Rp260.667 + USD 1,32 per bulan selama jalur web (9,57 − 99 ÷ 12 = 1,32; batas atas USD 10), dan Rp260.667 + USD 9,57 sejak jalur iOS native (batas atas USD 18,25) | RISK-62 |
| 8 | **Baru.** RISK-83 diberi K Sedang · D Sedang · P3 di §5 (brief §14); butir 3 tetap atas risiko berkedalaman penuh dan Tier 2, sehingga P3 bertambah satu di luar keduanya (jumlah semua P3: 32 + 4 + 1 = 37) | §5 |
| 9 | **Diubah.** Titik impas Mac 16 GB/512 GB 7,1 sampai 7,9 bulan (dikutip dari 14 §11 butir 6) | §3.5 RISK-01 |
| 10 | **Baru (keputusan 51).** **RISK-92 baru** (Teknis, K Sedang · D Tinggi · P2; dinilai dengan 18 §5.2.9, K dan D tetap). Akibatnya butir 1: 92 pengenal, 2 ditutup, **90 terbuka**; **72 berkedalaman penuh**, 8 Tier 2, 10 Tier 3. Butir 2: **Teknis 22**, jumlah 90; kategori lain tetap. Butir 3: **P1 13, P2 27, P3 32** (jumlah 72); Tier 2 tetap. Butir 4: sel Tinggi × Sedang kini **20** risiko (RISK-92 masuk). Butir 5 tetap 21 Pagar. Jalur publik di antara 72: 11 = P1 2 (RISK-25, RISK-58) + P2 4 (RISK-16, RISK-60, RISK-63, RISK-67) + P3 5 (RISK-26, RISK-59, RISK-61, RISK-62, RISK-72). Pengguna luar (keputusan 52): 2 = P3 2 (RISK-69, RISK-70). Jalur rumah tangga: 72 − 11 − 2 = 59 = P1 11 + P2 23 + P3 25. Per kategori, jalur publik: Produk 2, Bisnis 7, Hukum 2; pengguna luar: Hukum 2. Pagar di jalur publik: 1 (RISK-63), jadi Pagar di jalur rumah tangga 20 dari 21. Semua P3: 32 + 4 + 1 = 37 (butir 8 tetap). K dan D risiko lama tidak berubah | §1, §1.1, §2, §3.5 |
| 11 | **Baru (keputusan 52; 10.1 §11.3.1).** **RISK-93 baru** (Keamanan dan privasi, Pagar, K Pasti · D Kritis · P1, pengguna luar). Akibatnya butir 1: 93 pengenal, 2 ditutup, **91 terbuka**; **73 berkedalaman penuh**, 8 Tier 2, 10 Tier 3. Butir 2: **Keamanan dan privasi 22**, Teknis 22, jumlah 91; kategori lain tetap. Butir 3: **P1 14, P2 27, P3 32** (jumlah 73); Tier 2 tetap. Butir 4: sel Kritis × Pasti kini **2** risiko (RISK-11, RISK-93). Butir 5: **22 Pagar** (RISK-93 masuk). Butir 10: jalur publik tetap 11 (P1 2, P2 4, P3 5); pengguna luar **3** = P1 1 (RISK-93) + P3 2 (RISK-69, RISK-70); jalur rumah tangga tetap 73 − 11 − 3 = 59 = P1 11 + P2 23 + P3 25. Per kategori, pengguna luar: Hukum 2, Keamanan dan privasi 1. Pagar: jalur publik 1 (RISK-63), pengguna luar 1 (RISK-93), jalur rumah tangga 20. Semua P3 tetap 37. K dan D risiko lain tidak berubah | §1, §1.1, §2, §3.3, §8, §10.1 |
| 12 | **Tetap (keputusan 53 sampai 55).** RISK-92 kini Diterima (keputusan 55) dan RISK-40 berubah mitigasi dan risiko sisanya (keputusan 54); K, D, dan prioritas keduanya tetap. Risiko yang Diterima tetap dihitung terbuka, seperti RISK-01 dan RISK-11, jadi butir 1 sampai 5 dan 11 tidak berubah. Daftar §9 bertambah satu baris (RISK-92) | §2, §3.3 RISK-40, §3.5 RISK-92, §9 |
| 13 | **Diubah (S6; keputusan 59, jawaban advokat).** RISK-65 K Sedang → **Rendah**, D Kritis tetap, P1 → **P2** (matriks §0: Rendah × Kritis = P2); RISK-66 K Sedang → **Rendah**, D Tinggi tetap, P2 → **P3** (Rendah × Tinggi = P3); RISK-69 status pengguna luar → **jalur publik**, K, D, P tetap. Akibatnya butir 1, 2, 5 tetap (93 pengenal, 91 terbuka, 73 berkedalaman penuh; per kategori tetap; 22 Pagar). Butir 3: **P1 13, P2 27, P3 33** (jumlah 73); Tier 2 tetap. Butir 4: sel Kritis × Sedang **4** (RISK-65 keluar), Kritis × Rendah **5** (RISK-65 masuk), Tinggi × Sedang **19** (RISK-66 keluar), Tinggi × Rendah **4** (RISK-66 masuk). Butir 10 dan 11: jalur publik **12** = P1 2 + P2 4 + P3 6 (RISK-69 masuk); pengguna luar **2** = P1 1 (RISK-93) + P3 1 (RISK-70); jalur rumah tangga 73 − 12 − 2 = 59 = **P1 10 + P2 23 + P3 26**. Per kategori, jalur publik: Produk 2, Bisnis 7, Hukum 3; pengguna luar: Hukum 1, Keamanan dan privasi 1. Semua P3: 33 + 4 + 1 = **38**. P1 yang diterima dengan syarat: satu (RISK-11) | §1, §1.1, §2, §3 RISK-65, RISK-66, RISK-69, §9, §10.2 |
| 14 | **Tetap (keputusan 67, 68).** Mitigasi dan tenggat RISK-04, RISK-42, RISK-65, RISK-66, RISK-67, RISK-72, RISK-91, RISK-93 berubah; K, D, prioritas, dan status jalur tidak berubah. Butir 1 sampai 5, 10, 11, 13 tetap: P1 13, P2 27, P3 33 (73 berkedalaman penuh) | §1, §2, §3, §8 |
| 15 | **Diubah (web lebih dulu; ADR-030).** **RISK-94** Teknis, K Sedang · D Tinggi · P2; **RISK-95** Keamanan dan privasi, Pagar, K Rendah · D Kritis · P2; **RISK-96** Keamanan dan privasi, K Rendah · D Kritis · P2; **RISK-97** UX, K Pasti · D Sedang · P2, Diterima; semuanya jalur rumah tangga, berkedalaman penuh. **RISK-91** D Tinggi → Sedang, P2 → P3 (Sedang × Sedang = P3). **RISK-88** (Tier 3, Teknis) ditutup. RISK-01 dan RISK-18 dinilai ulang dengan K dan D tetap. Akibatnya butir 1: 97 pengenal, 3 ditutup (RISK-15, RISK-28, RISK-88), **94 terbuka**; **77 berkedalaman penuh**, 8 Tier 2, **9 Tier 3**; 77 + 8 + 9 = 94. Butir 2: Produk 15, **UX 7**, **Keamanan dan privasi 24**, Keuangan 6, Teknis 22 (−1 RISK-88, +1 RISK-94), Bisnis 8, Hukum 12; jumlah 94. Butir 3: **P1 13, P2 30, P3 34** (27 + 4 − 1 = 30; 33 + 1 = 34; jumlah 77); Tier 2 tetap. Butir 4: sel Kritis × Rendah **7** (RISK-95, RISK-96 masuk), Tinggi × Sedang tetap **19** (RISK-91 keluar, RISK-94 masuk), Sedang × Sedang **20** (RISK-91 masuk), Sedang × Pasti **1** (RISK-97). Periksa: P1 = 4 + 1 + 2 + 6 = 13; P2 = 7 + 19 + 3 + 1 = 30; P3 = 4 + 4 + 20 + 6 = 34. Butir 5: **23 Pagar** (RISK-95 masuk); Pagar jalur rumah tangga 21. Butir 10 dan 11: jalur publik tetap 12, pengguna luar tetap 2, jalur rumah tangga 77 − 12 − 2 = **63** = **P1 10 + P2 26 + P3 27** (23 + 4 − 1 = 26; 26 + 1 = 27). Semua P3: 34 + 4 + 1 = **39**. Butir 7 kini memuat kedua angka biaya (USD 1,32 selama jalur web; 9,57 sejak iOS native; 09.4 §7.1; 14 §11 butir 8) | §0, §1, §1.1, §2, §3, §5, §9 |
| 16 | **Tetap (bawaan 69.7).** RISK-95 kini Diterima; K, D, prioritas, dan tanda Pagar tetap. Risiko yang Diterima tetap dihitung terbuka (butir 12), jadi butir 15 tidak berubah: 94 terbuka, 77 berkedalaman penuh, P1 13, P2 30, P3 34, 23 Pagar, jalur rumah tangga 63. Daftar §9 bertambah satu baris (RISK-95) | §2, RISK-95, §6.1, §9 |

---

## Riwayat

| Tanggal | Tahap | Perubahan |
|---|---|---|
| 12 November 2026 | S5 fase B | Versi pertama. RISK-01 sampai RISK-17 diambil alih dari 09.4 §9 dengan pengenal tetap; RISK-18 sampai RISK-90 baru. Belum ditinjau `critic` |
| 12 November 2026 | S5 fase D | CS5-A-03 (irama: ADR-022 butir 1 berlaku, 80 hari usulan Q-219), A-07 (RLS sejak pilot di RISK-44), A-08 (RISK-01 harga Mac dari 18 §5.2.7), A-11 (RISK-83 dinilai; "unclear money ownership" dipetakan eksplisit ke RISK-49), A-14 (RISK-52 label perilaku aplikasi terpasang), A-18 (peta §10.2: K-61, L-06, W-31), A-21 (RISK-27 kalimat Ramadan), CS5-B-03 (Q-213 sebelum GR-2 [Usulan]); tahan tulis di §1 dan RISK-47; 10 §4.2 di §8 GR-4, RISK-11, RISK-40; SEC-138 di RISK-46; SEC-156 di RISK-63 dan §6.1; SEC-157 di §8; RISK-79 kolom tanda awal dibatasi seperti 13 §6.2. Tinjauan ulang `critic` A2: RISK-27 tanggal Idulfitri 1448 H terverifikasi; RISK-59 dinilai ulang dengan hasil R-15 (P2 → P3); §1.1, §8, §12 menyebut RISK-83 yang sudah dinilai |
| 12 November 2026 | Keputusan 42 | RISK-28 ditutup (§9); **RISK-91 baru** (Android menyusul, P2); RISK-18 memuat catatan terlambat dari iPad sampai Android tersedia (13 §3.2; 17 keputusan 42); §1 baris 1 tanpa "kedua platform"; RISK-01, RISK-17, RISK-19, RISK-53, RISK-54, RISK-56 menurut jalur Android 14 §3.8 dan QG-30 (Android); §8 GR-2 sepuluh butir 10 §4.1, baris "Jalur Android" baru, GR-4 dengan 13 GR-4 butir 8; §10.2 Q-37, Q-211; §12 butir 1, 2, 4 |
| 12 November 2026 | Keputusan 43 sampai 50 | RISK-01 **Diterima** (keputusan 43; register, kartu, §9; tanda awal dan risiko sisa baru: K Tinggi, D Sedang; penilaian sebelum mitigasi tetap); RISK-26 ditulis ulang untuk keputusan 46 (judul, uraian, mitigasi, "Harus ada pada" GR-2 → GR-3; K, D, P tetap); RISK-25, RISK-18, RISK-27, RISK-52 tanpa ambang 5 dari 7 dan dengan batas 80 hari diterima (keputusan 44); §8 GR-1 (Mac, server uji sementara), GR-2 (sebelas butir 10 §4.1, tinjauan `security-privacy`, uji iOS tanpa Mac, ADR-029; RISK-01 masuk, RISK-26 pindah ke GR-3), FS-2, GR-3, GR-4 (tinjauan `security-privacy`, keputusan 49); §10.2, §11, §12 butir 6. Tidak ada K, D, atau prioritas yang berubah, jadi §1.1 dan §12 butir 1 sampai 5 tetap |
| 12 November 2026 | Keputusan 51 | §0 aturan **jalur publik opsional** (status "jalur publik"; K dan D tidak dinilai ulang; tidak ada yang dihapus); §1 kelompok 4 kini RISK-10 (operasi bergantung pada satu orang), kelompok 5 jalur publik (RISK-25, RISK-58, RISK-65 bagian publik, RISK-26); §1.1 pembagian per jalur; §2 status 13 risiko "jalur publik" dan "bila jalur publik diambil" di kolom "Harus ada pada" RISK-12, 23, 44, 65, 68, 71, 91; RISK-29 "Harus ada pada" GR-3 → FS-2; kartu RISK-11, 16, 18, 23, 25, 26 (kutipan brief §1 yang baru; aturan bug berat lama yang belum diperbaiki, bacaan lead [Usulan]), 28, 29, 40, 44, 06, 04, 12, 58 sampai 63, 65 sampai 72, 91; §7; §8 GR-3 (aturan bug lama), GR-4, FS-3 dan GR-5 bertanda jalur publik opsional, baris Jalur Android; §11 (g), (h); §12 butir 10. Atas permintaan `security-privacy` (10 §7.8): **RISK-92 baru** (TestFlight internal sebagai jalur iOS keluarga tanpa batas waktu; kehilangan akses, bukan kebocoran; menunggu riset `researcher`), di §1.1, §2, §3.5, §8 GR-1 dan GR-2 |
| 12 November 2026 | Keputusan 52 | Orang tua yang diberi login dihitung pengguna luar (Q-209): §0 jalur **pengguna luar pertama** dipisah dari tahap publik, status baru "pengguna luar"; RISK-69 dan RISK-70 dari "jalur publik" ke "pengguna luar"; "Harus ada pada" RISK-06, 11, 12, 40, 44, 65, 68 kini "sebelum pengguna luar pertama" untuk bagian keamanan dan hukumnya; RISK-04 risiko sisa; §1 kelompok 5; §1.1 pembagian jalur; §8 GR-4 dibagi (a) pengguna luar pertama dan (b) beta, mengikuti 14 QG-31 (a) dan (b). Ketetapan lead atas §11 (g), (h): RISK-11 berlaku untuk pemakaian rumah tangga tanpa batas waktu (10 §3.3); RISK-65 bagian PSE menunggu P-51-1; §12 butir 10. K, D, dan prioritas tidak berubah |
| 12 November 2026 | Temuan 18 §5.2.9 | RISK-92 dinilai dengan temuan `researcher` (DPLA 7.4.1 **[Terverifikasi]**; penegakan tidak ditemukan): uraian, mitigasi (rekomendasi **[Usulan]** ADR-022: tetap A, App Store unlisted sebagai jalan keluar, Ad Hoc ditolak selama keputusan 43), pemilik mitigasi, "Harus ada pada", tanda awal, dan risiko sisa per pilihan pemilik produk; status register "menunggu pilihan pemilik produk"; §8 GR-1; §12 butir 10. K Sedang dan D Tinggi tetap, jadi §1.1 dan hitungan §12 tidak berubah |
| 12 November 2026 | Pembaruan `security-privacy` (10 §4.2 tabel baru; 10.1 §11.3.1) | **RISK-93 baru** (pengelola server adalah anak atau kerabat pengguna luar dari keluarganya sendiri; cermin RISK-11; Pagar, K Pasti · D Kritis · P1, pengguna luar; tidak dapat diterima pemilik produk atas nama orang tua) di §1 kelompok 5, §1.1, §2, §3.3, §8 GR-4 (a), §10.1 (TM-12, TM-31), §12 butir 11. RISK-11 "Harus ada pada" dan risiko sisa menunjuk RISK-93. §8 GR-4 (a) dan (b) mengikuti tiga kelompok QG-31: (a), (a+), (b); SEC-153 (d) di (b) kecuali lewat TestFlight eksternal |
| 12 November 2026 | Keputusan 53 sampai 55; susulan `security-privacy` (10 §3.3, §7.8) | Keputusan 55: RISK-92 **Diterima** (register, kartu, §8 GR-1, §9), risiko sisa menurut pilihan 1, syarat C dan akun demo di ADR-022. Keputusan 54: RISK-40 mitigasi pasangan opsional SEC-18 dan SEC-25 per akun di FS-2, "Harus ada pada", risiko sisa sebelum dan sesudah FS-2; §11 (g) dijawab; §8 FS-2. Keputusan 53: §8 jalur Android dan FS-2 (tinjauan `security-privacy` sebelum Android memegang data nyata dan tahunan). SEC-158 (pendaftaran tertutup, Pilot) di mitigasi RISK-40. RISK-93: pergantian versi pemberitahuan menulis baris di jejak keamanan; dikukuhkan `security-privacy`. §12 butir 12: tidak ada hitungan yang berubah |
| 12 November 2026 | Temuan `critic` | §0 fase dan gerbang: GR-4 dua pintu, FS-5 sampai FS-8 bukan jalur publik (13 §3.8); tenggat Android menjadi "sebelum jalur pengujian Google Play dibuka (GR-4 butir 8; QG-31 (b))" di RISK-91, RISK-18, dan §8; RISK-11 dan RISK-93 ditambah batas pola tahan berulang oleh pengelola (satu tahan per 30 hari, usulan lead); kedudukan keputusan 55 menjadi "diterima atas saran lead". Tidak ada hitungan yang berubah |
| 12 November 2026 | Pengukuhan `security-privacy` | Batas satu tahan per 30 hari menjadi 10.1 SEC-34 aturan (5), dihitung sejak awal tahan sebelumnya, tanpa pengecualian; dikutip di RISK-11 dan RISK-93. Risiko sisa keduanya ditambah: pengelola dengan akses basis data langsung dapat melewati batas itu, yang tercakup keputusan 25 selama server hanya dipakai pasangan, dan sejak pengguna luar pertama hanya lewat akses darurat SEC-36 yang tercatat. Tidak ada hitungan yang berubah |
| 12 November 2026 | SEC-34 aturan (5) diganti `security-privacy` (temuan `critic` R-2) | Mitigasi RISK-11 dan RISK-93: paling banyak satu versi sampai ke akun per 30 hari; versi di dalam jendela menunggu di slot tertunda (yang terbaru menimpa) dan tidak tampil; akun mencatat di bawah versi yang terakhir diakui. Kalimat "versi berikutnya tetap dibaca tanpa menahan" dihapus. Risiko sisa akses basis data langsung tetap. Tidak ada hitungan yang berubah |
| 12 November 2026 | SEC-34 aturan (5) opsi A (`security-privacy`; temuan `critic` R-3) | Batas 30 hari dan slot versi tertunda dicabut: setiap versi dan setiap pergantian versi menahan sampai versi terakhir diakui. Mitigasi RISK-11 dan RISK-93 diganti. Risiko sisa keduanya: tahan berulang lewat versi, dengan batasnya (tanpa data hilang, satu baca dan satu ketukan, tindakan pelindung tidak ditahan, terlihat di Riwayat keamanan), keputusan 25 bagi server pasangan, SEC-36 sejak pengguna luar pertama. Tidak ada hitungan yang berubah |
| 12 November 2026 | S6: pass konsistensi | Temuan `critic` S6: E-02 (RISK-27 tanpa GR-3 memperpanjang FS-2), E-03 (§0 fase: tinjauan tiga bulan bukan gerbang), E-04 (§8 FS-2: VL-02 hanya jalur publik), E-08 (RISK-83 mitigasi jalur keluarga **[Usulan]** `architect`), E-14 (RISK-25 VL-08 di tinjauan tiga bulan), E-15, E-16 (SEC-18 dan SEC-25 di §8), E-17 (§10.2 Q-222, Q-224, Q-227, Q-228). Keputusan 56: ADR-028 butir 3 diterima. **Jawaban advokat (keputusan 59, 60):** RISK-65 dinilai ulang (K Rendah, P2; tetap Diterima untuk pilot), RISK-66 (K Rendah, P3), RISK-69 ke jalur publik, RISK-71 dan RISK-11 menyebut jawabannya; §1, §1.1, §2, §8 GR-4 (a) dan FS-3, §9, §10.2, §11 (h), §12 butir 13. Keputusan 63: kerabat dewasa lain sama dengan orang tua (§8 GR-4 (a), RISK-93). Atas 10 §3.4 (`security-privacy`): SEC-147 keluar dari GR-4 (a) ke FS-3 dan GR-5; pemicu konsultasi advokat ulang **[Usulan]** di §1, §8, RISK-65, RISK-72 (Plus), RISK-91 (Play); keputusan 60 kini SEC-159 (RISK-11). Tidak ada hitungan yang berubah karena butir ini |
| 12 November 2026 | S6: keputusan 67 dan 68 | Pemicu konsultasi advokat ulang diringkas menjadi satu, sebelum go public, dan L-04, L-20 tidak dikejar (keputusan 67): §1, §8 GR-4 (a) dan FS-3, RISK-65, RISK-67, RISK-72, RISK-91. DPIA, catatan prosesor, kebijakan privasi, dan rencana pemberitahuan insiden (10 §4.1 butir 4 dan 5; SEC-139, SEC-148 (b), SEC-149, SEC-150; SEC-148 (a) tetap Pilot) keluar dari GR-2 dan GR-4 (a) ke GR-4 (b) dan go public (keputusan 68): §2 dan kartu RISK-66 (mitigasi diganti dengan syarat pengecualian; risiko sisa K Rendah, D Tinggi), RISK-04, RISK-42, RISK-93, §8 GR-2, GR-4 (a), GR-4 (b), FS-3; §12 butir 14. Tidak ada hitungan yang berubah |
| 12 November 2026 | Sesudah S6: web lebih dulu (keputusan 69; ADR-030; 10.5) | §0 aturan baca web lebih dulu, TM sampai TM-45, peran `web-dev`; §1 kelompok 1 dan jumlah Pagar; §1.1 peta dan pembagian jalur; §2 register (RISK-01, RISK-13, RISK-18, RISK-52, RISK-53, RISK-54, RISK-64, RISK-91, RISK-92 "Harus ada pada" atau penilaian; RISK-94 sampai RISK-97 baru); kartu RISK-18 (sebab iPad hilang; K dan D tetap), RISK-30, RISK-01 (keluar dari jalur kritis GR-2; K dan D tetap), RISK-13, RISK-14, RISK-40, RISK-41, RISK-42, RISK-52, RISK-53, RISK-54, RISK-57, RISK-64, RISK-91 (**D Tinggi → Sedang, P2 → P3**), RISK-92; kartu baru **RISK-94** (TM-44), **RISK-95** (TM-42, Pagar), **RISK-96** (TM-38, TM-37, TM-45), **RISK-97** (web tidak terasa native, Diterima); §5 dan §9 RISK-88 ditutup; §6.1 PL-05 lewat klien web; **§8 GR-1 diperbaiki (temuan audit 5)**: server uji sementara dan versi dipatok pindah ke FS-1 langkah 1; GR-1 kini urutan web lebih dulu, stack ADR-030, tanpa Mac; FS-1, GR-2 (10 §4.1 butir 12 dan 13), baris baru **Jalur iOS native**, Jalur Android, FS-2, GR-4 (a); §9 RISK-97; §10.1 TM-37 sampai TM-45; §10.2 R-17, Q-210; §11 (i), (ii) dan lead; §12 butir 15 |
| 13 November 2026 | Sesudah tinjauan `security-privacy` atas ADR-030 | RISK-95 merujuk LY-95 "Perangkat ini" (08 §25.6). Tidak ada hitungan yang berubah |
| 13 November 2026 | Temuan `critic` atas putaran web | "Calon keputusan 69" menjadi keputusan 69; §8 GR-1 memisahkan keputusan 69 (web lebih dulu) dari bawaan 69.5 (iOS lalu Android); RISK-95 **Diterima** (bawaan 69.7; register, kartu, §6.1, §9, §11 (i) selesai); baris Jalur iOS native: tinjauan `security-privacy` diterima (bawaan 69.4) dan sumber Mac AWS EC2 Mac atau Mac pinjaman (bawaan 69.6), juga di RISK-01; RISK-62 dan §12 butir 7 memakai biaya jalur web dan jalur iOS native dari 09.4 §7.1; §12 butir 16. §8 GR-1: bunyi ES §2, §3, §8 disetujui pemilik produk dan dipasang (13 November 2026). Tidak ada hitungan yang berubah |
