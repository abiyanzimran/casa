# 13 — Roadmap Produk

| | |
|---|---|
| Pemilik | `product-analyst` |
| Tahap | S5 — Rencana dan sintesis, fase A (12 November 2026; izin pemilik produk "gas", `_progress.md` "S5 dimulai") |
| Status | **Sesudah keputusan 53 sampai 55 dan SEC-158** (12 November 2026; belum ditinjau `critic`): tinjauan `security-privacy` tambahan (53); F-FAM-022 faktor kedua dan kode pemulihan di FS-2, gelombang T2-1 (54); saluran iOS TestFlight internal (55); pendaftaran tertutup sebagai ketergantungan FS-1. Diterapkan di §3.2, §3.4, §3.6, §4.1, §12. **Sesudah keputusan 51 dan 52** (tujuan tanpa rencana go public; orang tua dihitung pengguna luar; 12 November 2026; 17 bagian "Keputusan pemilik produk 51" dan "52"): roadmap yang dijadwalkan berakhir saat rumah tangga pemilik produk memakai aplikasi sehari-hari (FS-2 berlanjut); GR-3, FS-3, GR-5, FS-4 menjadi **jalur publik opsional** yang tidak dijadwalkan dan tidak dihapus; **GR-4 menjadi gerbang bersyarat** yang juga dilewati sebelum login pertama orang tua, tanpa tahap publik (keputusan 52); VL-02 pindah ke jalur itu; FS-5 sampai FS-8 dapat dipicu kebutuhan rumah tangga sendiri. Diterapkan di §0, §1, §2.2, §3.1 sampai §3.8, §4.1, §5, §6, §7, §8, §9, §10 (TB-6 baru), §11, §12; belum ditinjau `critic`. Sebelumnya: **sebelum S6: keputusan 43 sampai 50** (penutupan pertanyaan terbuka, 12 November 2026; 17 bagian "Keputusan pemilik produk sebelum S6") diterapkan di §1, §3.2 sampai §3.6, §5, §6.1 sampai §6.4, §8, §10, §11, §12; belum ditinjau `critic`. Sebelumnya: **sesudah S5: keputusan 42** (iOS dan iPadOS lebih dulu; pilot dengan data nyata di iPhone 13 dan iPad mini 7; 17 keputusan 42) diterapkan di §1, §3.2 sampai §3.4, §6.4, §7.4, §12; belum ditinjau `critic`. Sebelumnya: direvisi di S5 fase D atas tinjauan `critic` S5 (A, B) dan permintaan 14 §12.2, 15 §11; ditinjau ulang `critic` (B2), temuannya diperbaiki. Semua angka §11, termasuk butir 8 dan 9, dihitung ulang lead dan cocok (`_progress.md`) |
| Sumber yang mengikat | 17 keputusan 51; brief §0.3, §1 (Pilot devices, First users, Target time to MVP, Monthly infrastructure budget, Business model, Money movement scope), §1.1 (tujuan, masalah inti, tier, kriteria sukses, pagar), §13, §16; `CLAUDE.md`; `ENGINEERING-STANDARDS.md` §3, §5, §6, §9; dokumen 02, 03, 04, 05, 08 §19.8, 09 (indeks §6, berkas 04 §5–§7), 10 (indeks §3, §4, §4.1), 11, 12, 17, 18 |
| Dipakai oleh | 14 dan 15 (`architect`), 16 dan 01 (lead), tabel keterlacakan [05 §6](05-functional-requirements.md#6-tabel-keterlacakan) |

---

## 0. Cara membaca

- **Yang ditetapkan dokumen ini:** urutan fase, gerbang di antara fase, apa yang diukur untuk memutuskan, penempatan setiap fitur dokumen 04 pada fase, dan rekomendasi model bisnis. **Yang tidak:** aturan uang (07), aturan alur dan pagar (06), persyaratan keamanan (10), workstream teknis, strategi uji, dan rilis (14), daftar risiko (15). Rujukan "lihat 14" atau "lihat 15" menunjuk dokumen yang ditulis sesudah ini.
- **Tidak ada tanggal kalender** kecuali yang ditetapkan pihak luar (§8). Brief §1: "No fixed deadline. Prefer the smallest scope the pilot household will use daily over a date". Setiap fase berakhir pada **kriteria keluar**, bukan tanggal.
- **Tier mengatur urutan membangun, bukan batas rancangan** (brief §1.1; PRN-03). Roadmap ini mengurutkan pembangunan; rancangan untuk semua rumah tangga tahap publik sudah ada di 02 sampai 10.
- **Jalur rumah tangga dan jalur publik opsional** (17 keputusan 51, diputuskan pemilik produk; bacaan lead dikukuhkan). Kata-kata pemilik produk: "masalah go public gini aja, pikirkan aplikasi ini untuk keluarga pribadi aja, tapi standarnya untuk public sampe sini kamu paham gaa kira kira", lalu, atas bacaan lead, "yap benar seperti itu, untuk sekarang gitu aja, tapi tidak menutup kemungkinan untuk go public". Akibatnya di dokumen ini: roadmap yang dijadwalkan **berakhir saat rumah tangga pemilik produk memakai aplikasi sehari-hari** (GR-1, FS-1, GR-2, lalu FS-2 yang berlanjut dengan gelombang Tier 2 dan klien Android untuk Galaxy A55). GR-3, FS-3, GR-5, dan FS-4, beserta seluruh isinya (beta tertutup, harga Plus, wawancara keluarga lain, akun Play; dan advokat serta PSE bila belum dijawab di GR-4), adalah **jalur publik opsional**: tidak dijadwalkan, tidak dihapus, dan baru dimasuki bila pemilik produk kelak memutuskan go public. Rancangan untuk tahap publik tidak berubah (glosarium *Tahap publik*). **Keputusan 52** (diputuskan pemilik produk sendiri): orang tua yang diberi login dihitung pengguna luar, kata-katanya: "orang tua yaa hitung pengguna luar, yang harus di tanyakan ke advokat apa?". Karena itu **GR-4 dipisahkan dari jalur publik**: ia gerbang bersyarat yang dilewati sebelum login pertama orang tua, walaupun tahap publik tidak diambil (§3.6).
- **Label bukti** mengikuti [README](README.md#label-bukti). Seluruh urutan fase, gerbang, ambang ukur, dan rekomendasi di sini **[Usulan]** `product-analyst` kecuali disebut lain. Ambang angka di rencana validasi tidak pernah fakta.
- **Kata pemilik produk** hanya dikutip dari `_progress.md` atau dokumen 17, dengan kedudukannya.
- **Pengenal dokumen ini**, semuanya terdaftar di [README](README.md):

| Kode | Arti |
|---|---|
| `FS-N` | Fase roadmap produk |
| `GR-N` | Gerbang (milestone) di antara fase: kriteria keluar satu fase dan syarat masuk fase berikutnya |
| `VL-NN` | Ukuran dalam rencana validasi |
| `T2-N` | Gelombang pelepasan fitur Tier 2 di dalam pilot (§3.4; glosarium *Gelombang Tier 2*). Bukan `T-NN` dokumen 07 dan bukan `P-T2-N` di 05 `tier-2.md` |
| `TB-N` | Tantangan dokumen ini terhadap brief (§10) |

---

## 1. Ringkasan

| Fase | Nama | Fase brief §13 | Isi | Kriteria keluar (gerbang) |
|---|---|---|---|---|
| **FS-0** | Validasi sebelum dan selama membangun | Phase 0 | Gerbang contoh layar (08 §19.8); uji alur inti di perangkat; garis dasar pilot (VL-00). Wawancara keluarga lain (R-08) dan uji pesaing (R-02) hanya di jalur publik opsional (§6.3) | Contoh layar disetujui kedua pasangan sebelum token dikunci |
| **FS-1** | Bangun Tier 1 | Phase 1 (MVP) | 55 fitur Tier 1, 190 FR Must, di iOS dan iPadOS (satu aplikasi SwiftUI), dengan data uji | **GR-2 Data nyata**: sebelas butir 10 §4.1, semua AC Tier 1 lulus, aplikasi terpasang di iPhone 13 istri dan iPad mini 7 pemilik produk (keputusan 42) |
| **FS-2** | Pilot berjalan, lalu pemakaian sehari-hari | Phase 1 (MVP) | Tier 1 dengan data nyata oleh kedua pasangan; 19 fitur Tier 2 menyusul bertahap; klien Android untuk Galaxy A55 sesudah GR-2 (§3.4); ukuran VL-00, VL-01, VL-03 sampai VL-10; tinjauan tiga bulan (§3.4) | **Tidak ada gerbang keluar.** Titik akhir roadmap: rumah tangga memakai aplikasi sehari-hari (§3.4). Sesudahnya FS-2 berlanjut sebagai pemakaian dan pemeliharaan |
| | ***Gerbang bersyarat*** | | *Tidak dijadwalkan; dilewati sebelum pengguna luar pertama, juga tanpa tahap publik (keputusan 52)* | |
| **GR-4** (bersyarat) | Pengguna luar pertama | — | Sebelum orang tua mendapat login (keputusan 52; Q-209), atau sebelum penguji beta di jalur publik opsional: advokat L-01, L-14, L-24; PSE bila wajib; butir 10 §4.2; tinjauan `security-privacy` (§3.6) | Login orang tua (FS-5), atau beta (FS-3) |
| | ***Jalur publik opsional*** | | *Tidak dijadwalkan; dimasuki hanya bila pemilik produk kelak memutuskan go public (keputusan 51)* | |
| **GR-3** (opsional) | Evaluasi go public | — | Paling cepat tiga bulan sesudah GR-2: KS-1 sampai KS-4 terpenuhi, pemakaian harian yang bertahan menurut keputusan 46 (VL-02), tanpa kegagalan invarian (§3.5) | Masuk FS-3 |
| **FS-3** (opsional) | Persiapan tahap publik dan beta tertutup | Phase 2 | Butir "Wajib sebelum tahap publik" 10 §4; advokat; PSE; toko; langganan rumah tangga; beta tertutup sekurangnya 5 rumah tangga, ditambah sampai sekurangnya 14 penguji Android karena akun Play pribadi (keputusan 48) memerlukan 12 penguji selama 14 hari (§6.4) | **GR-4 Pengguna luar pertama** (di awal fase, bila belum dilewati untuk orang tua) dan **GR-5 Rilis publik** (di akhir) |
| **FS-4** (opsional) | Tahap publik untuk pasangan | Phase 2 | MVP yang sama untuk rumah tangga dua dewasa (dan satu dewasa) di Indonesia | Dipakai rumah tangga di luar pilot dengan pola yang sama dengan GR-3; salah satu pemicu FS-5 sampai FS-8 |
| | ***Tier 3*** | | *Tidak dijadwalkan; dimulai bila pemicunya terpenuhi, dari kebutuhan rumah tangga sendiri (FS-2) atau dari jalur publik opsional (FS-4)* | |
| **FS-5** | Struktur keluarga lain | Phase 2–3 | Banyak rumah tangga per orang, tanggungan dan wali, preset lain, lebih dari dua dewasa | Dijadwalkan bila pemicunya terpenuhi (§3.8) |
| **FS-6** | Urusan rumah lanjutan | Phase 3 | Anggaran periode lain, rencana umum, `MEAL`, kalender bersama, album | Sama |
| **FS-7** | Kecerdasan dan otomasi | Phase 4 | `AI` di luar versi ringan: percakapan saran, entri berbantuan | Sama |
| **FS-8** | Ekosistem dan skala | Phase 5 | Impor mutasi bank, integrasi belanja, sinkron kalender; kandidat klien web | Sama |

```mermaid
flowchart LR
    S6["S6 selesai;<br/>pemilik produk<br/>mengakhiri perencanaan"] --> GR1{{"GR-1<br/>Mulai bangun"}}
    GR1 --> FS1["FS-1<br/>Bangun Tier 1<br/>(data uji)"]
    FS0["FS-0<br/>Validasi"] -.-> FS1
    FS1 --> GR2{{"GR-2<br/>Data nyata"}}
    GR2 --> FS2["FS-2<br/>Pilot berjalan, lalu<br/>pemakaian sehari-hari<br/>Tier 1 + Tier 2 bertahap"]
    GR2 -. "sesudah GR-2" .-> AND["Klien Android<br/>(Galaxy A55)"]
    AND -.-> FS2
    FS2 --> ENDPT(["Titik akhir roadmap:<br/>rumah tangga memakai<br/>aplikasi sehari-hari"])
    FS2 -. "kebutuhan rumah<br/>tangga sendiri" .-> T3["FS-5 sampai FS-8<br/>Tier 3, menurut pemicunya"]
    FS2 -. "orang tua akan<br/>mendapat login<br/>(keputusan 52)" .-> GR4{{"GR-4 (bersyarat)<br/>Pengguna luar pertama"}}
    GR4 -. "login orang tua" .-> T3
    subgraph OPS["Jalur publik opsional: tidak dijadwalkan (keputusan 51)"]
        GR3{{"GR-3<br/>Evaluasi go public"}}
        FS3["FS-3<br/>Persiapan publik<br/>+ beta tertutup"]
        FS3 --> GR5{{"GR-5<br/>Rilis publik"}}
        GR5 --> FS4["FS-4<br/>Tahap publik<br/>untuk pasangan"]
    end
    GR3 -- "bila GR-4<br/>belum dilewati" --> GR4
    GR3 -- "bila GR-4<br/>sudah dilewati" --> FS3
    GR4 -. "jalur publik" .-> FS3
    ENDPT -. "hanya bila pemilik produk<br/>memutuskan go public" .-> GR3
    FS4 -. "pemicu" .-> T3
```

**Mengapa struktur brief §13 diubah sedikit** (brief §13 mengizinkan "or justify a better one"):

1. **Phase 0 tidak mendahului seluruh pembangunan.** Rumah tangga pilot adalah sumber kebenaran (brief §1.1), jadi pilot itu sendiri adalah validasinya. Yang perlu divalidasi sebelum kode hanyalah tampilan (08 §19.8) dan kecepatan alur inti. Validasi untuk tahap publik (keluarga lain, §6.3) tetap tertulis tetapi berada di jalur publik opsional (keputusan 51).
2. **Phase 1 dipecah menjadi FS-1 dan FS-2** karena di antaranya ada gerbang yang tidak boleh dilewati, yaitu data nyata (10 §4.1), dan karena Tier 2 dibangun di dalam pilot yang sudah berjalan, bukan sebelum pilot dimulai.
3. **Phase 2 "Core Family Operations" berarti membawa MVP ke rumah tangga lain**, bukan menambah modul. Versi ringan `SHP`, `MNT`, `PLN`, `CAL`, dan `CARE` sudah masuk MVP lewat Tier 2 (brief §1.1), sehingga yang tersisa untuk "operasi inti" adalah gerbang hukum, toko, dan model bisnis. Sejak keputusan 51 seluruhnya jalur publik opsional, kecuali GR-4, yang sejak keputusan 52 juga dilewati sebelum login orang tua.
4. **Phase 3 dipecah menjadi FS-5 dan FS-6**, karena struktur keluarga lain bergantung pada tinjauan hukum anak (L-08 sampai L-10, L-27 sampai L-30), sedangkan urusan rumah lanjutan hanya bergantung pada bukti pemakaian. Keduanya dapat berjalan dalam urutan apa pun.
5. **GR-4 dipisahkan dari tahap publik** (keputusan 52). Pemilik produk memutuskan sendiri bahwa orang tua yang diberi login dihitung pengguna luar, jadi gerbang "pengguna luar pertama" dapat dilewati tanpa go public: sebelum login pertama orang tua, bukan hanya sebelum beta.
6. **Phase 2 sampai Phase 5 tidak dijadwalkan** (keputusan 51; TB-6). Roadmap yang dijadwalkan berakhir di FS-2. Phase 3 sampai 5 (FS-5 sampai FS-8) dimulai menurut pemicu, yang kini dapat datang dari rumah tangga sendiri, misalnya brief §1 "Parents and children get logins later" (§3.8).

---

## 2. Kerangka prioritas

### 2.1 Mengapa bukan RICE

Brief §13 memberi contoh "RICE scores plus MoSCoW labels". RICE ditolak **[Usulan]**:

- *Reach* di pilot adalah satu rumah tangga untuk setiap fitur, jadi tidak membedakan apa pun.
- *Impact* dan *Confidence* harus diisi angka yang tidak punya sumber: dokumen 18 §0.3 mencatat bahwa tidak ada bukti dari pengguna di luar rumah tangga pilot. Mengisi angka itu melanggar PRN-04 ("Angka tanpa sumber primer tidak dikutip").
- Satu-satunya angka yang berarti, *Effort*, adalah usaha **pengguna** (04 §6), bukan usaha tim; RICE mengukur yang kedua.

### 2.2 Kerangka yang dipakai: MoSCoW dengan lima syarat yang dapat diperiksa

Satu kerangka, dipakai untuk semua 132 fitur. Label **final** sejak dokumen ini (04 "Cara membaca" menyebutnya sementara sampai S5).

| Label | Syarat masuk (semuanya harus terpenuhi) | Fase |
|---|---|---|
| **Must** | (1) Tier 1 menurut brief §1.1 sebagaimana diubah keputusan checkpoint; (2) menyebut `MSL-1` sampai `MSL-5` atau persyaratan "Pilot" dokumen 10; (3) tanpa fitur ini salah satu dari KS-1 sampai KS-4 tidak dapat terpenuhi, atau pilot tidak boleh memproses data nyata | FS-1 |
| **Must bersyarat** | Pagar yang wajib ada bila fitur Tier 2 yang dijaganya dibangun (05 §4) | FS-2, bersama fiturnya |
| **Should** | Tier 2; menyebut `MSL-n`; usaha berulang nol atau sukarela (04 §6) | FS-2 |
| **Could** | Tier 3; menyebut `MSL-n`, tetapi pilot belum menunjukkan perlunya | "Menunggu bukti" (§4), ditimbang di tinjauan tiga bulan dan sesudahnya di FS-2; di jalur publik opsional juga di GR-3 dan sesudah GR-5 |
| **Future** | Tier 3; memerlukan fondasi, tinjauan hukum, atau rumah tangga yang belum ada di pilot | FS-5 sampai FS-8, menurut pemicunya |
| **Will not build** | Tidak menyebut masalah, atau melanggar pagar (06 AT-3, AT-4, PG-09; PRN-02, PRN-11) | Tidak pernah; alasannya di §4.2 |

Urutan **di dalam** satu fase memakai empat pertimbangan, dalam urutan ini: ketergantungan (04 §5.1); masalah yang lebih dekat ke masalah inti (`MSL-1` sampai `MSL-5` sebelum `MSL-6`, `MSL-7`); kriteria sukses yang butuh waktu lebih lama untuk dinilai (KS-5 butuh pemakaian berminggu-minggu); usaha berulang yang lebih kecil.

---

## 3. Fase dan gerbang

### 3.1 FS-0 — Validasi sebelum dan selama membangun

**Untuk siapa.** Bagian pilot: pemilik produk dan istrinya. Bagian publik: keluarga lain (R-08), hanya di jalur publik opsional (keputusan 51). Rinciannya di §6.

| Kegiatan | Kapan | Yang dihasilkan | Rujukan |
|---|---|---|---|
| Gerbang contoh layar arah C: Beranda, Catat, Chat di iOS dan Android, terang dan gelap, beserta ikon | Di FS-1, sebelum token dibekukan | Persetujuan kedua pasangan, atau perubahan pada lima unsur anti-templat | 08 §19.8; diterima atas saran lead ("pasang semua", 17 keputusan S4 (b)) |
| Uji kecepatan alur inti di perangkat pilot: catat belanja, catat dan konfirmasi serah terima, ajukan dan setujui permintaan | Sesudah alur pertama dapat dipasang (FS-1) | Jumlah ketukan dan waktu dibandingkan dengan ES §3 dan PRN-05 | ES §3 `[DEFAULT]` dikukuhkan ADR sesudah pengukuran pertama |
| Garis dasar pilot (VL-00) | Sebelum GR-2 | Satu percakapan 10 menit, §6.1 | §6.1 |
| Wawancara dan uji prototipe dengan 8–12 rumah tangga lain (**jalur publik opsional**) | Tidak dijadwalkan; hanya bila pemilik produk mempertimbangkan go public, selesai sebelum GR-5 | Pola kebutuhan, kesediaan membayar, keberatan terhadap persetujuan | R-08; §6.3; keputusan 51 |
| Uji langsung dua atau tiga aplikasi pesaing (**jalur publik opsional**) | Sebelum materi publik apa pun (GR-5) | Jumlah langkah untuk satu belanja, satu tarik tunai, satu serah terima | R-02 (memerlukan pemilik produk); 18 §1.4 |

### 3.2 GR-1 — Mulai bangun, dan FS-1 — Bangun Tier 1

**GR-1 (syarat masuk FS-1).**

| # | Syarat | Pemutus |
|---|---|---|
| 1 | Pemilik produk menyatakan fase perencanaan selesai (`CLAUDE.md` "Current phase") | Pemilik produk |
| 2 | Agen tim bangun dibuat lead sesudah stack disetujui di S4 (`CLAUDE.md` "Build team") | Lead |
| 3 | Platform yang dibangun lebih dulu: **diputuskan, iOS dan iPadOS** (Q-37 selesai; 17 keputusan 42) | Pemilik produk; sudah diputuskan |

**Platform: diputuskan pemilik produk sendiri** (17 keputusan 42, 12 November 2026), kata-katanya: "a saja karena saya ada ipad dan iphone jadi saya bisa pake ipad bener gaa?" dan "ios ipad yang bakalan dipakai pilot ya".

| Hal | Yang berlaku |
|---|---|
| Dibangun lebih dulu | iOS dan iPadOS, satu aplikasi SwiftUI, lewat Xcode Cloud dan TestFlight internal (keputusan 32); Mac disewa sekali untuk penyiapan. **Tanpa sewa Mac bulanan dan tanpa membeli Mac** selama dana belum ada (keputusan 43, diterima atas saran lead: "Tetap Xcode Cloud saja (Saran)"; menjawab ADR-022 butir 4): setiap perubahan tampilan iOS diuji lewat build di awan dan TestFlight, lebih lambat dan tanpa debug di perangkat (17 W-28; 15 RISK-01); ditinjau lagi bila dana ada. Server dan kontrak lebih dulu atau bersamaan, karena semua aturan ada di server (ADR-013); urutan teknisnya di 14 §3.4 |
| Pilot dengan data nyata | iPhone 13 istri (pencatat utama) dan iPad mini 7 pemilik produk. GR-2 tidak menunggu Android |
| Android (Kotlin; keputusan stack tetap) | Dibangun sesudah GR-2 untuk Galaxy A55 milik pemilik produk, di dalam FS-2 (§3.4). Semula bawaan lead **[Usulan]** (keputusan 42); kini **diterima atas saran lead lewat keputusan 50**, karena bunyi brief yang ia terima ("Setuju bunyi ini (Saran)") memuat "Android is built after real data starts and joins the pilot before the beta (decision 42)" (17 keputusan 50). Sejak keputusan 51 brief berbunyi "...before the beta, if a beta is ever held (decisions 42, 51)": batas "sebelum GR-4" hanya berlaku bila jalur publik opsional dibuka; di jalur rumah tangga tidak ada batas waktu. "Di dalam pilot bila waktu memungkinkan" tetap rincian lead |
| Brief §1 "Pilot devices" | "the pilot needs iOS and Android" tidak lagi menjadi syarat data nyata. Bunyi baru brief §1 baris "Pilot devices" dan "Platforms" diterima atas saran lead lewat keputusan 50 ("Setuju bunyi ini (Saran)"; 17 K-66 selesai) dan dipasang lead. Maksudnya, kedua pasangan dapat memakai aplikasi, tetap terpenuhi |

**Riwayat pilihan, singkat.** Draf S5 menyajikan tiga pilihan untuk Q-37: keduanya bersamaan, Android dulu, iOS dulu. Saran tim fase D adalah B (Android dulu: siklus kerja di Linux, Mac tertunda), disertai usulan bahwa data nyata tidak dimulai dengan satu platform. Saran B **digantikan** keputusan 42 setelah dua fakta yang belum diketahui saat itu (Q-211): istri tidak memakai Android, dan dialah yang sering mencatat. Usulan "tidak dengan satu platform" juga digantikan; alasannya (PRN-01, KS-1: kedua orang mencatat dan melihat angka yang sama sejak data nyata) tetap terpenuhi, karena kedua pasangan memakai aplikasi iOS yang sama.

**Risiko sisa.** Sampai Android tersedia, pemilik produk mencatat pengeluarannya di luar rumah belakangan dari iPad, karena iPad bukan perangkat yang dibawa saat berbelanja **[Asumsi]**. Entri terlambat diterima (PRN-06; 07 §3.11), tetapi catatan yang tertunda memperbesar 15 RISK-18 dan membuat angka utama tertinggal sampai catatan itu masuk. Yang meredamnya: pemilik produk menyatakan "saya jarang kemana-mana", dan VL-07 butir (2) menanyakan momen yang terlewat di minggu kedua; bila momen yang terlewat terutama pengeluaran pemilik produk di luar rumah, Android dinaikkan urutannya di dalam FS-2 **[Usulan]**.

**FS-1 berisi** semua fitur Tier 1 (§4.1) dan konsep dokumen 11 yang menjadi kriteria penerimaan FR induk Tier 1 (IDEA-01, IDEA-02, IDEA-03; THR yang diberikan; 05 §6). Tidak ada fitur Tier 2 yang menahan GR-2.

**Ketergantungan yang jatuh di dalam FS-1:**

| Butir | Kapan paling lambat | Rujukan |
|---|---|---|
| **S4-2: penyedia hosting, cadangan, dan email** | **Sebelum deploy pertama yang memegang data nyata** (keputusan 45, diterima atas saran lead: "Sebelum data nyata (Saran)"; Q-220 selesai). Sampai saat itu hanya ada instans sementara berisi data uji (glosarium *Server uji sementara*). Syaratnya **[Usulan]** `security-privacy` (10 §4.1 butir 7 (a) sampai (d)), bukan isi keputusan 45: SEC-40 sampai SEC-42 berlaku, tanpa cadangan; tidak pernah menerima data nyata (akun hanya dengan email domain uji), kunci produksi dibuat baru, dan instans beserta disk dan snapshot-nya dihapus sesudah server produksi berjalan, tanggalnya dicatat di runbook (14 §7.8). Pernyataan tertulis letak data tetap wajib sebelum data nyata (10 §4.1 butir 7, tidak dapat ditunda), dan server tetap di Indonesia (keputusan 10) | Keputusan 31, diputuskan pemilik produk sendiri: "#S4-2: nanti aja pas mau deploy", diartikan keputusan 45; ADR-020, ADR-026 tetap Diusulkan; R-14; K-62 |
| Catatan persetujuan uji (SEC-109 (b)) | Dibaca dan diakui kedua penguji sebelum build TestFlight pertama dipasang, jadi di FS-1, bukan baru di GR-2 | 10 §4.1 butir 3 |
| Email untuk verifikasi dan pemulihan (SEC-24) | Sebelum akun pertama dibuat di server yang di-deploy | 09 §5 |
| Pendaftaran tertutup: akun hanya lewat token undangan atau token pemasangan sekali pakai dari pengelola, terikat pada email satu orang yang memverifikasinya dan memasang kata sandinya sendiri; penolakan seragam. Akun pertama pilot (pengelola) dibuat lewat token pemasangan, pasangannya lewat undangan | Sebelum akun pertama dibuat di server yang di-deploy, karena API terbuka ke internet apa pun jalur pemasangan aplikasinya; diuji sebagai persyaratan "Pilot" dokumen 10 di GR-2 butir 1 | 10.4 SEC-158 **[Usulan]** `security-privacy`; 05 FR-FAM-001/AC-9, AC-10 |
| Akun Apple Developer Program dan sewa Mac sekali untuk alur kerja pertama Xcode Cloud | Sebelum build iOS pertama | Keputusan 32; 09.4 §5 |
| Kunci API Claude dan pembangunan kumpulan kalimat pertama (R1) | Tidak menahan GR-2; diperlukan sebelum fitur `AI` Tier 2 di FS-2 | Keputusan 33; 10 §3 10-B |
| Deskriptor rating usia Apple | Lihat §8 (12 Januari 2027) | 18 §2.12.3; SEC-83 |

### 3.3 GR-2 — Data nyata

Satu-satunya gerbang yang menyangkut keselamatan data orang yang bukan pemilik produk (istrinya), jadi daftarnya tidak dipersingkat (`CLAUDE.md` "Instruction priority" butir 3).

| # | Syarat | Rujukan |
|---|---|---|
| 1 | Sebelas butir 10 §4.1 (butir 10 baru sejak keputusan 42; butir 11, tinjauan hanya-temuan `security-privacy` sebelum GR-2, baru sejak keputusan 49), dicentang bertanggal di runbook pilot: persyaratan "Pilot" dokumen 10 lulus uji; istri membaca sendiri pemberitahuan pengelola dan pengakuannya tercatat (keputusan 25; 10-F A); catatan persetujuan uji (SEC-109 (b); sudah wajib sebelum build TestFlight pertama dipasang, §3.2); DPIA dan catatan prosesor (Q-208, diterima atas saran lead); kebijakan privasi versi pilot; rencana insiden; penyedia hosting dipilih dengan pernyataan letak data; pemulihan cadangan diuji; SAPA 129 terverifikasi atau tidak ditampilkan; pemilik produk ditanya apakah iPad mini 7 dipegang orang lain dan apakah sidik jari orang lain terdaftar di Touch ID-nya (iPad mini 7 memakai Touch ID di tombol atas, **[Terverifikasi]** 18 §2.12.9), jawabannya dicatat (SEC-59 (e) **[Usulan]** `security-privacy`; kunci aplikasi tetap pilihannya, SEC-19). Bagian khusus Android (10 §4.3) bukan syarat GR-2 | 10 §4.1 |
| 2 | Setiap kriteria penerimaan Tier 1 punya uji otomatis yang menyebut ID-nya dan lulus | ES §6; 05 |
| 3 | Uji invarian 07 dan 06, dan setiap contoh terhitung 07 sebagai uji yang dapat dijalankan, lulus | ES §5 |
| 4 | Aplikasi terpasang lewat TestFlight internal di iPhone 13 istri dan iPad mini 7 pemilik produk. Setiap alur Tier 1 yang dipakai pemilik produk (08 §16.2) lulus uji tata letak di iPad mini 7 menurut 08 §16.4 baris "Sebelum GR-2" (potret, lanskap, jendela berdampingan *compact* dan *regular*, jendela yang diubah ukurannya di tengah alur). Android tidak disyaratkan | Keputusan 42; bunyi brief §1 menurut keputusan 50 (17 K-66 selesai); 08 §16.2, §16.4; 09.4 §5 |
| 5 | Sasaran ES §3 diukur di iPhone 13 dan iPad mini 7 (perangkat harian pemilik produk), dan angkanya dikukuhkan atau diubah lewat ADR. Galaxy A55 diukur sebelum klien Android memegang data nyata (§3.4). Bunyi ES §3 yang sejalan diterima atas saran lead lewat keputusan 50 (17 K-67 selesai) dan dipasang lewat ADR-029 | ES §3; 14 QG-30 butir 4; keputusan 42, 50 |
| 6 | Gerbang contoh layar 08 §19.8 sudah dilewati | §3.1 |
| 7 | Garis dasar VL-00 sudah dicatat | §6.1 |
| 8 | Q-222 dijawab pemilik produk bersama istrinya: siapa mengoperasikan server bila pemilik produk berhalangan lama. Jawaban "tidak ada pengganti" juga jawaban. Yang diketahui keduanya sebelum menjawab: akses pengelola bagi istri juga memberinya akses teknis ke data pribadi suaminya, dan pemberitahuan pengelola server berlaku juga ke arah itu **[Usulan]** | 17 Q-222; 15 RISK-10, RISK-11; **Pagar** |

**Bukan syarat GR-2** (diputuskan pemilik produk sendiri, keputusan 34): jawaban advokat atas L-01, L-14, L-24 dan pendaftaran PSE; keduanya kini syarat GR-4 (butir 1 dan 2), jadi berlaku di kedua pintunya: sebelum login pertama orang tua (keputusan 52) atau di jalur publik opsional (keputusan 51); waktu advokat di luar GR-4 menunggu 17 Q-224. Risiko bila pilot ternyata wajib mendaftar PSE (W-32) diterima pemilik produk. Karena pemakaian rumah tangga kini berlanjut tanpa batas waktu, penerimaan itu juga berlaku tanpa batas waktu; lead diminta menyampaikannya kepada pemilik produk (17 W-32, 2.18) **[Usulan]**.

### 3.4 FS-2 — Pilot berjalan, lalu pemakaian sehari-hari

**Jendela tiga bulan.** Brief §1.1 menilai sukses "after three months of pilot use". **[Usulan]** (A-40): jendela dimulai di GR-2, saat kedua pasangan memakai Tier 1 dengan data nyata, bukan saat Tier 2 lengkap. Alasannya KS-1 sampai KS-4 hanya bergantung pada Tier 1. KS-5 bergantung pada `CARE` dan `AI` (Tier 2), jadi penilaiannya disyaratkan sekurangnya enam minggu pemakaian fitur itu **[Usulan]**; bila Tier 2 datang terlambat, tinjauan tiga bulan menilai KS-1 sampai KS-4 tepat waktu dan KS-5 menyusul. Batas bawah tiga bulan dalam keputusan 46 ("Tetap minimal 3 bulan") adalah syarat GR-3 di jalur publik opsional (§3.5), bukan syarat jalur rumah tangga.

**Tinjauan tiga bulan dan titik akhir roadmap** (keputusan 51) **[Usulan]**. Di akhir bulan ketiga sesudah GR-2, pasangan membaca ukuran §6.1 (VL-01, VL-03 sampai VL-10) terhadap KS-1 sampai KS-5. Tinjauan ini **bukan gerbang**: tidak ada fase yang menunggunya, dan hasilnya perbaikan di FS-2 (§6.2). **Titik akhir roadmap tercapai** bila KS-1 sampai KS-4 terpenuhi (VL-01, VL-03, VL-04, VL-05 memenuhi ambangnya), VL-10 nol, VL-08 tanpa temuan menghalangi, dan gelombang Tier 2 yang dipilih pasangan sudah dilepas. KS-5 dinilai pasangan sendiri (VL-06) dan tidak menahan. Titik akhir tidak menunggu klien Android. Sesudahnya FS-2 berlanjut sebagai pemakaian dan pemeliharaan: irama rilis di bawah tetap berlaku selama aplikasi dipasang lewat TestFlight, tinjauan hanya-temuan `security-privacy` berjalan setahun sekali (keputusan 53), Tier 3 dibuka menurut pemicunya (§3.8), dan jalur publik opsional (§3.5) hanya dimasuki atas keputusan pemilik produk.

**Tier 2 di dalam pilot.** 19 fitur, dilepas bertahap. Urutan bawaan di bawah diturunkan dari §2.2; **pasangan pilot boleh mengubahnya di GR-2** (Q-212), karena merekalah sumber kebenaran MVP.

| Gelombang | Fitur | Must bersyarat yang ikut | Mengapa di urutan ini |
|---|---|---|---|
| T2-1 | F-CARE-001, F-CARE-002, F-CARE-003; F-REQ-013; F-FAM-021; F-FAM-022 | FR-COM-016, FR-CARE-002, FR-CARE-004 | KS-5 perlu waktu paling lama untuk dinilai; permintaan tindakan menjawab `MSL-6`/`MSL-7` dan brief §1.1 ("things one asks of the other kept as tasks until done"); sebutan anggota hampir tanpa biaya. F-FAM-022 (faktor kedua dan kode pemulihan, opsional per akun; FR-FAM-026) dibangun di FS-2 dan tidak menahan GR-2: keputusan 54, diterima atas saran lead ("Opsional, setelah dipakai harian (Saran)"); di gelombang pertama **[Usulan]** karena ia kecil dan menutup RISK-40 lebih awal |
| T2-2 | F-AI-005 lebih dulu, lalu F-AI-001, F-AI-002, F-AI-003, F-AI-004, F-AI-008 | FR-AI-005 sampai FR-AI-009 | F-AI-002 membutuhkan `CARE`; F-AI-005 (persetujuan, batas biaya, gagal aman) harus ada sebelum `AI` mana pun |
| T2-3 | F-SHP-001, F-SHP-002; F-MNT-001 sampai F-MNT-003; F-CAL-001; F-PLN-001; konsep 11 IDEA-15 (tinjauan bulanan), IDEA-19 (butuh tenang) | — | `MSL-4` (anggaran dapur) sudah sebagian terjawab Tier 1; sisanya `MSL-6`. IDEA-15 dan IDEA-19 belum punya FR (17 Q-187); ditulis sebelum dibangun |
| Di luar gelombang | Konsep 11 IDEA-08 (musim Lebaran), menurut §8 | — | Waktunya ditentukan Ramadan, bukan urutan gelombang; belum punya FR (Q-187) |

**Rilis dalam pilot.** Build TestFlight berlaku 90 hari sejak diunggah ke TestFlight, bukan sejak dipasang (**[Terverifikasi]** 18 §2.12.6, "All builds uploaded to TestFlight are active for 90 days"), hampir sama dengan jendela tiga bulan. **Yang mengikat:** ADR-028 butir 4, diterima pemilik produk (keputusan 44, diterima atas saran lead: "Ya, 60/80 hari (Saran)"; Q-219 selesai): rilis pemeliharaan bila 60 hari tanpa rilis, dan batas keras 80 hari, keduanya dihitung dari unggahan ke TestFlight, agar tersisa 10 hari untuk build yang gagal atau tertahan di Xcode Cloud. Aturan ini memperketat ADR-022 butir 1 (paling lambat 90 hari, sasaran 60; keputusan 32). Setiap gelombang Tier 2 sekaligus memperbarui build. Jadwal rilisnya di 14 §6.3. Sejak keputusan 51 TestFlight internal menjadi jalur iOS tetap keluarga, bukan jalur beta sementara; apakah Apple Guideline 2.2 dan DPLA 7.4.1 mengizinkannya belum diketahui (17 R-19, L-41; risikonya ketersediaan, bukan kerahasiaan). Saluran pemasangan iOS jangka panjang: **keputusan 55, diterima atas saran lead** ("TestFlight, pindah bila dihentikan (Saran)"; 17 Q-228 selesai): tetap TestFlight internal, pindah ke App Store tersembunyi (unlisted) bila Apple menghentikannya. Keputusan 43 dan 44 tetap berlaku. Pindah ke App Store unlisted membuka tautan pasang kepada siapa pun yang memegangnya, jadi pendaftaran tertutup (10.4 SEC-158; FR-FAM-001/AC-9) sudah harus berlaku, dan memang berlaku sejak FS-1.

**Anak 12+ di rumah tangga pilot** (Q-186; 03 §4.1). **Tidak ada yang dijadwalkan untuk anak selama pilot** **[Usulan]**: tanpa akun anak dan tanpa profil tanggungan (bawaan dikukuhkan dokumen 10, D-109). Risiko anak memegang ponsel orang tua (PL-03; L-29) sudah dijawab kunci aplikasi Tier 1 (F-FAM-019). Akun dan profil anak menunggu FS-5 dan tinjauan advokat L-08 sampai L-10, L-27 sampai L-30. Tinjauan tiga bulan menanyakan sekali apakah pasangan ingin anak tercatat sebagai profil; jawabannya dapat menjadi pemicu FS-5 dari rumah tangga sendiri (brief §1: "Parents and children get logins later"; §3.8), bukan perubahan pilot.

**Klien Android di dalam pilot** (waktunya diterima atas saran lead lewat keputusan 50; semula bawaan lead di keputusan 42). Dibangun sesudah GR-2 untuk Galaxy A55 milik pemilik produk, di dalam FS-2 bila waktu memungkinkan, dan dipasang sebagai APK tanpa akun Google Play (§8). Di jalur rumah tangga tidak ada batas waktu; bila jalur publik opsional kelak dibuka, klien Android harus siap sebelum jalur pengujian Google Play dibuka (GR-4 butir 8; QG-31 (b)), bukan sebelum GR-4 itu sendiri, karena beta memakai jalur pengujian Google Play (§6.4) dan rumah tangga beta dapat memakai Android (brief §1: "if a beta is ever held"; keputusan 51). Ia tidak menahan GR-2 dan tidak menggeser jendela tiga bulan. Server, kontrak, dan aturan sama (ADR-013), jadi tidak ada aturan yang berubah. Sebelum APK memegang data nyata di Galaxy A55, syarat GR-2 yang bergantung pada klien diulang untuk klien itu **[Usulan]**: butir 2 (uji AC Tier 1 yang berjalan di klien, termasuk bagian Android FR-COM-011/AC-10, FR-NTF-005/AC-6 dan AC-7), butir 3 (uji invarian lewat klien), butir 5 (ES §3 di Galaxy A55), dan persyaratan "Pilot" dokumen 10 yang menyangkut klien Android; daftar teknisnya di 14. Ditambah satu tinjauan hanya-temuan `security-privacy` sebelum APK memegang data nyata, dengan temuan menghalangi diperbaiki lebih dulu (keputusan 53, diterima atas saran lead: "Ya, tambah keduanya (Saran)"). Selama itu klien Android memakai data uji di staging (ADR-028 butir 1). Sesudah klien Android memegang data nyata, VL-07 diulang di Galaxy A55.

**Web dan PC tidak dibangun di pilot.** Keputusan pemilik produk selama S5 (`_progress.md`, "Pertanyaan pemilik produk tentang web", 12 November 2026), kata-katanya: "A aja, catat dulu"; **diterima atas saran lead** (17 keputusan 41). Kandidat klien web ada di §3.8; pertanyaannya Q-210.

### 3.5 GR-3 — Evaluasi go public, paling cepat tiga bulan (jalur publik opsional)

**Jalur publik opsional (keputusan 51).** §3.5, FS-3 di §3.6, §3.7, §6.3, §6.4, dan §7 tidak dijadwalkan dan tidak dikerjakan kecuali pemilik produk kelak memutuskan mempertimbangkan go public. Isinya dipertahankan, tidak dihapus, agar jalan itu tidak perlu dirancang ulang. GR-4 (§3.6) bukan bagian jalur ini saja: ia gerbang bersyarat yang juga dilewati tanpa go public, sebelum orang tua mendapat login (keputusan 52). GR-3 dinilai hanya saat itu; keputusan 46 adalah syaratnya (brief §1: "if the owner later decides on it, only after the pilot shows sustained daily use"). Keputusan 47 dan 48 juga bagian jalur ini. Keputusan "lanjut ke tahap publik" adalah keputusan pemilik produk; GR-3 menyiapkan bahannya. Ukurannya di §6.1, tabel keputusannya di §6.2.

**Pemakaian harian yang bertahan** (brief §1: "only after the pilot shows sustained daily use"). **Diputuskan pemilik produk sendiri** (17 keputusan 46; Q-213 selesai), kata-katanya: "1 - 2 bulan dan bisa lebih, ketika kita sudah merasa nyaman dan aman serta gada bug baru publish". Lead lalu bertanya lanjut, karena brief §1.1 menilai "after three months of pilot use" (alasan itu milik pertanyaan lead, bukan pernyataan pemilik produk), dan pemilik produk memilih "Tetap minimal 3 bulan" serta, untuk arti bug baru, "2 minggu tanpa bug berat (Saran)". Kedudukannya: tiga bulan dan kedua syarat diputuskan pemilik produk sendiri; arti bug diterima atas saran lead; rumusan gabungan di bawah adalah bacaan lead. Menurut bacaan lead atas keputusan itu (`_progress.md`; glosarium *Pemakaian harian yang bertahan*; dapat dikoreksi pemilik produk), GR-3 dinilai **paling cepat tiga bulan sesudah GR-2**, bila:

| Syarat | Isi | Ukuran |
|---|---|---|
| (a) | Kedua pasangan, masing-masing, menyatakan sendiri merasa nyaman dan aman **memakai aplikasi**. Pertanyaannya tentang aplikasi, tidak pernah tentang pasangan (15 RISK-25) | VL-02 bagian (a) |
| (b) | Dalam 14 hari terakhir tidak ditemukan bug baru yang berat: angka uang salah, data hilang atau bocor, atau aplikasi tidak dapat dipakai. Bug tampilan kecil tidak menahan. **Bacaan lead [Usulan]** (`_progress.md`, atas `critic` P-4): bug berat yang ditemukan lebih dari 14 hari sebelumnya tetapi belum diperbaiki juga menahan | VL-02 bagian (b) |

**"Bisa lebih"** berarti GR-3 mundur selama (a) atau (b) belum terpenuhi; tidak ada batas atas. Ambang lama "sedikitnya 5 dari 7 hari" (usulan VL-02 sebelum keputusan 46) **tidak lagi menjadi syarat**; frekuensi catatan tetap boleh dibaca sebagai bahan, di tingkat rumah tangga saja (PRN-11). Karena ditetapkan sebelum GR-2, ambang ini tidak diatur sesudah hasil terlihat (15 RISK-26). Gerbang ini tidak lagi menguji kata "daily" brief §1; perbedaannya dicatat di 17 K-68 dan keputusan menang (`CLAUDE.md` "Sources of truth").

Ringkas:

| Hasil | Keputusan |
|---|---|
| KS-1 sampai KS-4 terpenuhi, pemakaian harian yang bertahan menurut keputusan 46 (VL-02 (a) dan (b)), tanpa kegagalan invarian (VL-10), tanpa tanda bahaya pagar (VL-08) | Lanjut ke GR-4 dan FS-3 |
| VL-02 (a) atau (b) belum terpenuhi | GR-3 mundur (keputusan 46, "bisa lebih"); §6.2 |
| Sebagian KS terpenuhi, atau Tier 2 datang terlambat | GR-3 dinilai ulang satu sampai dua bulan kemudian **[Usulan]**, sesudah perbaikan yang disebut di FS-2 |
| KS-1 tidak terpenuhi | Masalah inti belum terjawab: tahap publik tidak dimulai; penyebabnya dicari dan dipecahkan di FS-2 |
| Fitur Tier 2 tidak dipakai | Kandidat dikeluarkan dari cakupan publik (PRN-02); bagi rumah tangga sendiri, lihat §6.2 |

### 3.6 GR-4 bersyarat, dan FS-3 — Persiapan tahap publik dan beta tertutup (jalur publik opsional)

FS-3 tidak dijadwalkan (keputusan 51; §3.5). **GR-4 — Pengguna luar pertama: gerbang bersyarat, terpisah dari tahap publik** (keputusan 52). Dilewati sebelum siapa pun di luar kedua pasangan memakai layanan, termasuk penguji, grup TestFlight eksternal, atau jalur pengujian Google Play, dan **termasuk orang tua yang diberi login di rumah tangga pemilik produk sendiri**. Keputusan 52, diputuskan pemilik produk sendiri, berlawanan dengan usulan `security-privacy` (17 Q-226), kata-katanya: "orang tua yaa hitung pengguna luar, yang harus di tanyakan ke advokat apa?". Jadi GR-4 punya dua pintu masuk: **(a) jalur keluarga**, sebelum login pertama orang tua (FS-5; §3.8), tanpa tahap publik; **(b) jalur publik opsional**, sesudah GR-3 dan sebelum beta. Yang dilewati lebih dulu berlaku; GR-4 tidak diulang di pintu kedua kecuali ada butir yang bunyinya baru terpicu di sana. Semua butir di bawah berlaku di kedua pintu; butir yang bunyinya sendiri bersyarat (butir 3 TestFlight eksternal atau App Store; butir 4 bagian Google Play dan butir 5 bila Play dipakai; butir 6 bila SES; butir 8 sebelum jalur pengujian Play dibuka) berlaku bila syaratnya terjadi **[Usulan]**. Waktu advokat (butir 1) di luar GR-4: 17 Q-224, masih terbuka (pemilik produk: "saya tanya besok"; keputusan yang berlaku tetap sampai ia menjawab). Tinjauan `security-privacy` tambahan sebelum klien Android memegang data nyata dan setahun sekali: keputusan 53, diterima atas saran lead (17 Q-225 selesai; 14 dan `CLAUDE.md` milik pemiliknya masing-masing):

| # | Syarat | Rujukan |
|---|---|---|
| 1 | Advokat menjawab **L-01, L-14, L-24** | Q-209, diterima atas saran lead ("pasang semua"); 10 §4.1 |
| 2 | Pendaftaran PSE lingkup privat lewat OSS sebelum pengguna memakai layanan, bila advokat menyatakan wajib | L-14; R-07; 18 §2.11 (Permenkominfo 5/2020, **[Terverifikasi]** naskah 2020) |
| 3 | SEC-153 (d) dan sisa SEC-154: sebelum build pertama ke TestFlight eksternal atau App Store | 10 §4; R-13 |
| 4 | Rating usia: Apple 18+ dengan override, deskriptor dijawab; Google Play target audiens 18 ke atas | SEC-83; 18 §2.12.3 |
| 5 | Google Play bila dipakai: akun pengembang **pribadi** (keputusan 48, diterima atas saran lead: "Pribadi (Saran)"; Q-221 selesai; bagian jalur publik opsional, keputusan 51). **Langkah sebelum membuat akun:** menanyakan ke dukungan Play Console apakah aplikasi keuangan rumah tangga wajib memakai akun organisasi (18 §2.4; 18 §6.4 butir 12); bila jawabannya wajib, hal itu dibawa ke pemilik produk sebelum akun dibuat. Akun pengembang (USD 25 sekali, **[Terverifikasi]** 18 §5.1.1); karena akun pribadi itu dibuat sesudah 13 November 2023, rencana *closed testing* dengan sekurangnya 12 penguji yang ikut tanpa putus 14 hari sebelum produksi (**[Terverifikasi]** `architect`, 14 §6.7; syarat itu jatuh di GR-5, ukurannya di §6.4); target API 36 (**[Terverifikasi]** 18 §5.1.1), deklarasi aplikasi kesehatan (18 §2.4), Data safety di luar jalur internal testing (18 §2.4) | SEC-143 |
| 6 | Email keluar dari sandbox SES, bila S4-2 memilih SES | 09 §6; ADR-026 |
| 7 | Butir "sebelum tahap publik" yang jatuh sebelum GR-4, karena pengguna luar memakai server yang sama dan tidak terikat keputusan 25 (penerimaan pasangan pilot atas pengelola-anggota): di pintu (a) orang tua, yang pengelolanya anggota keluarga tetapi bukan pasangan mereka; di pintu (b) rumah tangga beta, yang pengelolanya bukan anggota mereka: SEC-36 dan SEC-39 (c) publik (akses darurat, log keluar dari VPS operator); SEC-34 teks umum dan SEC-139 versi publik; SEC-84; SEC-147 bila L-14 mewajibkan pendaftaran. Menurut bunyinya sendiri juga di GR-4: SEC-153 (d), SEC-154 (butir 3), SEC-142, SEC-143 (butir 5). SEC-25 dan SEC-18 dibangun di FS-2 sebagai pasangan opsional per akun (keputusan 54; FR-FAM-026), jadi di GR-4 cukup dipastikan sudah lulus; bila FS-2 belum melepasnya, ia menahan GR-4. SEC-107 swalayan tidak di sini: dibangun di FS-1 (FR-FAM-013/AC-12), jadi sudah ada sejak GR-2 | 10 §4.2 **[Usulan]** `security-privacy` |
| 8 | Klien Android memenuhi syarat ulang GR-2 di §3.4 dan terpasang di Galaxy A55, sebelum jalur pengujian Google Play dibuka | 17 keputusan 42, 50 (waktu Android diterima atas saran lead lewat keputusan 50); syarat ulang §3.4 **[Usulan]** |
| 9 | Tinjauan hanya-temuan `security-privacy` sebelum GR-4 selesai, dan temuan yang menghalangi sudah diperbaiki | Keputusan 49, diterima atas saran lead; `CLAUDE.md` "Build team" |

**Isi FS-3 sesudah GR-4:**

| Kelompok | Butir | Rujukan |
|---|---|---|
| Hukum | Semua butir L berstatus "Tahap publik", terutama L-04, L-20, L-22, L-23 (keputusan sementara append-only, keputusan 24); tinjauan advokat atas kebijakan privasi, DPIA, dan catatan prosesor | 17 bagian 6; 10 §4 baris "Regulasi dan operasi" |
| Keamanan | Sisa kolom "Wajib sebelum tahap publik" 10 §4, sebelum GR-5: SEC-45, SEC-105 tautan web, SEC-140 tautan web, SEC-141, SEC-157 (uji penetrasi pihak luar); ADR E2EE (10-D A); tinjauan ulang SQLCipher. Yang jatuh sebelum GR-4 ada di GR-4 butir 7. RLS tidak di daftar ini: dipaksakan sejak pilot (ADR-016; D-99; 14 §3.3 langkah 1) | 10 §4, §4.2; 09 §6 |
| `AI` | Jalur R2 hanya lewat ADR baru sesudah advokat; tanpa itu R1 tetap | 09 §6; SEC-122 |
| Model bisnis | Langganan rumah tangga lewat penagihan toko, hak akses per rumah tangga di server (§7.5) | §7 |
| Ekspor | SEC-107 swalayan dipenuhi FR-FAM-013 yang dibuka selama aktif (FR-FAM-013/AC-12, dibangun di FS-1, ada sejak GR-2); **tidak** memerlukan F-FIN-023, yang karena itu pindah ke "Menunggu bukti" (§4.1) | 10 §7.4 |
| Validasi | Beta tertutup sekurangnya 5 rumah tangga, ditambah bila syarat penguji Play menuntut (§6.4), sekurangnya 8 minggu; hasil wawancara R-08 (§6.3) | §6 |
| Cakupan | Segmen rilis publik pertama (Q-214): rumah tangga dua dewasa setara dan satu dewasa, tanpa akun anak | §3.7 |

**GR-5 — Rilis publik:** semua butir FS-3 selesai; beta memenuhi §6.4; lolos tinjauan App Store dan Google Play; harga dan isi Plus diputuskan pemilik produk sesudah hasil wawancara dan beta (§7.5); materi publik tidak memakai positioning yang belum diuji (PRN-04; 18 §1.4).

### 3.7 FS-4 — Tahap publik untuk pasangan (jalur publik opsional)

Tidak dijadwalkan (keputusan 51; §3.5). Yang dirilis adalah MVP yang sama dengan pilot (Tier 1 + Tier 2 yang lolos GR-3), untuk rumah tangga yang modelnya sudah dibangun: **dua dewasa setara** (struktur 1, dan struktur 2 tanpa akun anak) dan **satu dewasa** (SCN-43 sudah Tier 1). **[Usulan]** dan Q-214. Alasannya: itulah satu-satunya bentuk yang sudah divalidasi pilot; struktur lain dirancang (03 §4) tetapi belum dipakai siapa pun.

Infrastruktur tumbuh menurut 09.4 §7.2 dan pemicu 09 §6 (lihat 14). Usaha mencatat diukur ulang pada beta dan rilis publik dengan cara yang sama dengan pilot, tanpa analitik perilaku (SEC-49).

### 3.8 FS-5 sampai FS-8 — Tier 3 (garis besar)

Kedalaman satu halaman (`CLAUDE.md`, depth rule). Tidak ada FR; garis besarnya di [05 tier-3.md](05-functional-requirements/tier-3.md). Urutan di antara FS-5 sampai FS-8 tidak tetap: fase dimulai bila pemicunya terpenuhi. **Sejak keputusan 51** pemicunya dapat datang dari rumah tangga pemilik produk sendiri di FS-2, misalnya brief §1 "Parents and children get logins later", atau dari jalur publik opsional (FS-4). Fase Tier 3 karena itu tidak ikut menjadi jalur publik opsional: menjadikannya hanya-publik akan menutup login orang tua dan anak yang dijanjikan brief §1 bagi rumah tangga ini.

| Fase | Fitur | Pemicu **[Usulan]** | Prasyarat |
|---|---|---|---|
| FS-5 Struktur keluarga lain (12) | F-FAM-009, 010, 011, 012; F-COM-013; F-FIN-024; F-REQ-014, 015, 016, 017; F-CARE-005; F-CAL-006 | Kebutuhan nyata rumah tangga pemilik produk (FS-2; misalnya profil atau akun anak, login orang tua); di jalur publik opsional, permintaan rumah tangga tahap publik (FS-4) atau wawancara R-08; untuk anak, juga jawaban advokat. **Login orang tua mensyaratkan GR-4 lebih dulu** (keputusan 52; §3.6) | L-08 sampai L-10, L-27 sampai L-30, R-10 (anak); SEC-38 (satu orang di beberapa rumah tangga); Q-88 (orang dewasa sebagai tanggungan) |
| FS-6 Urusan rumah lanjutan (10) | F-FIN-021; F-PLN-003, 004; F-GOL-002; F-MNT-005; F-MEAL-002, 003; F-CARE-004; F-CAL-003; F-TML-001 | Fitur Tier 2 induknya dipakai rutin di FS-2 (atau FS-4), dan pemakainya meminta kedalaman | Pembukaan anggaran periode lain adalah keputusan pemilik produk (brief §5C, Tier 3) |
| FS-7 Kecerdasan dan otomasi (2) | F-AI-006 (percakapan saran), F-AI-007 (entri berbantuan `AI`) | Bukti bahwa usaha mencatat tetap penghalang terbesar di FS-2 (atau sesudah FS-4); jalur `AI` yang sah | L-05, L-16, L-18, L-19, L-32; ADR R2; kebijakan AI-Generated Content Google Play dinilai ulang (18 §2.4 butir 3) |
| FS-8 Ekosistem dan skala (3) | F-FIN-025 (impor mutasi), F-SHP-005 (belanja online), F-CAL-005 (sinkron kalender) | Permintaan nyata; jalur yang tersedia | 18 §3 (impor manual sebagai jalur realistis; tanpa API pembeli marketplace; kalender perangkat tanpa izin baca); L-13, L-15 |

**Kandidat, bukan komitmen: klien web atau PC.** Tidak dibangun di pilot (§3.4); hanya itu yang diputuskan (keputusan 41). Apakah dan kapan dibuka sesudah pilot adalah Q-210. Yang membuatnya mungkin tanpa membongkar backend: semua aturan ada di server (ADR-013). Prasyaratnya: model keamanan browser di dokumen 10 (belum dirancang), sesi di PC yang dipakai bersama, dan aturan data di perangkat (ADR-024) yang harus diterjemahkan ke penyimpanan browser. Klaim tentang ketersediaan klien browser untuk kontrak Connect dan status Windows Subsystem for Android **tidak dipakai** di sini karena belum diverifikasi di 18 (permintaan riset R-17).

---

## 4. Penempatan setiap fitur

### 4.1 Ringkasan per fase

| Penempatan | Fitur | Jumlah |
|---|---|---|
| **FS-1** (Must) | `FAM` 001–008, 019, 020; `COM` 001, 008–012, 014; `FIN` 001–020; `REQ` 001–012; `CAL` 002; `NTF` 001–005 | **55** |
| **FS-2** (Should; F-AI-005 Must bersyarat) | `FAM` 021, 022; `REQ` 013; `SHP` 001, 002; `PLN` 001; `MNT` 001–003; `CARE` 001–003; `CAL` 001; `AI` 001–005, 008 | **19** |
| **FS-3** (Could, bersyarat; jalur publik opsional) | Tidak ada. F-FIN-023 semula di sini bila SEC-107 memerlukannya; 10 §7.4 menyatakan tidak, jadi pindah ke "Menunggu bukti" | **0** |
| **FS-5** (Future) | §3.8 | **12** |
| **FS-6** (Future) | §3.8 | **10** |
| **FS-7** (Future) | §3.8 | **2** |
| **FS-8** (Future) | §3.8 | **3** |
| **Menunggu bukti** (Could; ditimbang di tinjauan tiga bulan dan sesudahnya; di jalur publik opsional juga di GR-3 dan sesudah GR-5) | `FAM` 013; `COM` 003, 004, 007; `FIN` 022, 023, 027; `SHP` 003, 004; `INV` 001; `PLN` 002; `GOL` 001; `MNT` 004; `MEAL` 001; `CAL` 004; `TML` 002; `NTF` 006, 007 | **18** |
| **Will not build** | §4.2 | **13** |
| **Jumlah** | | **132** (Tier 1 55, Tier 2 19, Tier 3 58; sama dengan 04 §1; F-FAM-022 baru sesudah keputusan 54) |

Label MoSCoW dokumen 04 tidak diubah: setiap fitur "Future" di 04 berada di FS-5 sampai FS-8, dan setiap "Could" di "Menunggu bukti" atau FS-3. Pengecualian yang disengaja tidak ada.

**"Menunggu bukti" bukan fase.** Fitur Could tidak dijadwalkan sampai pilot atau tahap publik menunjukkan masalah yang dijawabnya. Contoh pemicu: F-COM-003, 004, 007 (pesan suara, reaksi, pencarian) dinilai sesudah pilot (D-70); F-NTF-006 (peringatan anggaran) bila anggaran dapur tetap terlampaui sesudah tiga bulan (Q-134); F-FIN-027 (Sembunyikan angka) bila tombol mata (S4-10 (iii) A) tidak cukup.

### 4.2 Will not build

Alasannya kanonis di 04 §8.1; diulang satu baris agar brief §13 ("Will not build (with reasons)") terpenuhi di satu tempat.

| Fitur | Alasan singkat |
|---|---|
| F-COM-002, 005, 006 utas bebas, pin, polling | Balasan pada kartu dan F-REQ-012 sudah memberi tiap objek tempatnya; dua orang dewasa tidak memungut suara |
| F-FAM-014 visualisasi hubungan | Tidak ada masalah §1.1; aturan tidak membaca hubungan (06 PG-01) |
| F-FAM-015 direktori kontak | Tidak ada masalah §1.1; menambah data pihak luar (Q-28) |
| F-FAM-016 pengumuman | Tidak ada masalah §1.1 bagi dua orang dewasa |
| F-FAM-017, F-NTF-008 riwayat kegiatan per anggota, umpan kegiatan | Dilarang 06 AT-3 |
| F-FAM-018 mengeluarkan orang dewasa lain | Keluar sendiri mencapai keadaan yang sama tanpa tombol tekanan (06 INV-A48) |
| F-FIN-026 penyusutan aset otomatis | Tidak ada masalah §1.1 |
| F-INV-002 stok manual | Contoh pelanggaran PRN-02 |
| F-GOL-003 kontribusi per orang lintas periode | Dilarang 06 AT-4 |
| F-CARE-006 pelacakan siklus | 04 §4.4 (D-63) |

---

## 5. Ketergantungan dan urutan

```mermaid
flowchart TB
    subgraph Keputusan_pemilik["Keputusan pemilik produk"]
        Q37["Q-37 platform lebih dulu:<br/>iOS dan iPadOS (keputusan 42)"]
        S42["S4-2 hosting, cadangan, email:<br/>sebelum data nyata (keputusan 45)"]
        Q222["Q-222 pengganti pengelola server<br/>(pemilik produk dan istrinya)"]
    end
    DPIA["DPIA, catatan prosesor,<br/>kebijakan privasi pilot (Q-208)"]
    ADVANAK["L-08–L-10, L-27–L-30 (anak)"]
    Q37 --> FS1["FS-1"]
    S42 --> GR2{{"GR-2"}}
    DPIA --> GR2
    Q222 --> GR2
    FS1 --> GR2 --> FS2["FS-2 berlanjut:<br/>pemakaian sehari-hari"]
    FS2 -. "pemicu rumah tangga" .-> FS5["FS-5"]
    FS2 -. "pemicu rumah tangga" .-> FS678["FS-6, FS-7, FS-8"]
    ADVANAK --> FS5
    subgraph Bersyarat["Gerbang bersyarat: pengguna luar pertama (keputusan 52)"]
        ADV1["Advokat L-01, L-14, L-24 (Q-209)"]
        PSE["Pendaftaran PSE lewat OSS,<br/>bila advokat menyatakan wajib"]
        SEC153["SEC-153 (d), SEC-154<br/>(bila TestFlight eksternal)"]
        RATING["Rating usia, deskriptor Apple<br/>(12 Januari 2027)"]
        GR4{{"GR-4"}}
        ADV1 --> GR4
        PSE --> GR4
        SEC153 --> GR4
        RATING --> GR4
    end
    FS2 -. "orang tua akan<br/>mendapat login" .-> GR4
    GR4 -. "login orang tua" .-> FS5
    subgraph Opsional["Jalur publik opsional: tidak dijadwalkan (keputusan 51)"]
        Q213["Syarat GR-3: keputusan 46<br/>(Q-213 selesai)"]
        BM["Model bisnis (Q-36): harga dan isi Plus"]
        ADVALL["Semua butir L 'Tahap publik'"]
        BILL["Penagihan toko (R-15)"]
        R08["Wawancara R-08 (§6.3)"]
        GR3{{"GR-3"}}
        FS3["FS-3"] --> GR5{{"GR-5"}} --> FS4["FS-4"]
        Q213 --> GR3
        ADVALL --> GR5
        BM --> GR5
        BILL --> GR5
        R08 --> GR5
    end
    GR3 -- "bila GR-4 belum dilewati" --> GR4
    GR3 -- "bila GR-4 sudah dilewati" --> FS3
    GR4 -. "jalur publik" .-> FS3
    FS2 -. "hanya bila pemilik produk<br/>memutuskan go public" .-> GR3
    FS4 -. "pemicu publik" .-> FS5
    FS4 -. "pemicu publik" .-> FS678
```

Ketergantungan antarfitur di dalam FS-1 dan FS-2 kanonis di 04 §5.1 dan tidak diulang. Yang menyangkut urutan Tier 2: `CARE` sebelum F-AI-002; F-AI-005 sebelum semua `AI`; F-SHP-002 bertumpu pada F-FIN-005 dan F-REQ-005 (Tier 1; 04 §5.2 KM-17).

---

## 6. Rencana validasi

### 6.1 FS-2: apa yang diukur

**Sejak keputusan 51** hanya ukuran yang melayani rumah tangga sendiri yang dijalankan di jalur rumah tangga: apakah aplikasi benar-benar menghapus kebutuhan bertanya (VL-00, VL-01), kriteria sukses brief §1.1 lainnya (VL-03 sampai VL-06), usaha mencatat (VL-07), pagar (VL-08), beban notifikasi (VL-09), dan kebenaran data (VL-10). **VL-02 hanya melayani keputusan go public** (keputusan 46), jadi pindah ke jalur publik opsional dan tidak diukur kecuali pemilik produk mempertimbangkan go public. Tidak ada ukuran baru. Sesudah tinjauan tiga bulan tidak ada percakapan berkala lagi kecuali untuk ukuran yang belum memenuhi ambangnya, atau bila pasangan memintanya; VL-10 tetap berjalan otomatis **[Usulan]**.

**Batas yang mengikat semua ukuran.** Tidak ada SDK analitik perilaku (SEC-49); tidak ada hitungan, peringkat, atau tren per orang (PRN-11; 06 AT-4, PG-09); `read_state` dan kehadiran tidak dibaca apa pun (INV-A45); KS-5 adalah pernyataan pasangan, bukan metrik (brief §1.1; PRN-04). Karena itu setiap ukuran di bawah berasal dari salah satu dari tiga sumber: **(a)** jawaban pasangan sendiri; **(b)** data rumah tangga yang memang terbaca oleh **kedua** anggota dewasa, dibaca pada tingkat rumah tangga, tidak per orang, lewat layar aplikasi atau jalur baca yang sama (SEC-06), **tidak pernah** lewat kueri basis data langsung oleh pengelola server, yang di pilot adalah suami; **(c)** hasil lulus atau gagal pemeriksaan teknis yang sudah ada, tanpa isi (uji invarian, ES §5; SEC-44). Buku pribadi, tempat uang pribadi, setelan `AI`, kehadiran, tanda dibaca, dan `CARE` yang tidak dibagikan tidak pernah menjadi sumber. Tidak ada pengumpulan data baru: validasi tidak menambah instrumen, medan, tabel, log, atau analitik; ukuran waktu VL-07 diambil dalam sesi uji yang disepakati, bukan dari build yang dipasang; jawaban VL-08 tidak disimpan di aplikasi atau server. Batas ini dikukuhkan `security-privacy` (10 §7.4, baris B-04) **[Usulan]**.

**Usaha bagi pasangan:** satu percakapan akhir bulan sekitar 20 sampai 30 menit **[Asumsi]**, ditambah satu percakapan singkat di akhir minggu kedua. Tidak ada buku harian.

| ID | Mengukur | Cara | Sumber | Kapan | Ambang **[Usulan]** |
|---|---|---|---|---|---|
| VL-00 | Garis dasar | Tiga pertanyaan sebelum GR-2: berapa kali minggu lalu bertanya soal uang; jalan-jalan atau pembelian tidak rutin bulan lalu yang datang mendadak; apa yang sekarang dicatat | (a) | Sekali | — (pembanding) |
| VL-01 | **KS-1** tahu sisa dan pemegang | Masing-masing, terpisah, menjawab "uang kita sisa berapa dan siapa memegangnya" dari aplikasi tanpa bertanya, lalu keduanya menghitung uang (rekonsiliasi kas, 07 §3.10) dan mencocokkan. Sama dengan uji pilot PRN-01 | (a), (b) | Akhir bulan 1, 2, 3 | Bulan 2 dan 3: keduanya dapat menjawab dan jawabannya sama; selisih yang tersisa sudah tampil sebagai "belum dikonfirmasi" atau selisih kas, bukan tersembunyi |
| VL-02 | **Jalur publik opsional** (keputusan 51; hanya bila pemilik produk mempertimbangkan go public). **Pemakaian harian yang bertahan** (brief §1; definisinya keputusan 46, §3.5) | **(a)** Masing-masing, terpisah, menjawab tentang aplikasi, bukan tentang pasangan: "Apakah kamu merasa nyaman dan aman memakai aplikasi ini sehari-hari?" Saluran jawaban sama dengan VL-08: jawaban istri tidak melewati pemilik produk, dan ia memilih sendiri apa yang dibacakan (15 RISK-25). Jawaban "belum" tidak perlu dijelaskan; yang menjawabnya boleh menyatakan "sudah" kapan saja sesudahnya, dan tidak ditanya ulang lebih sering dari percakapan akhir bulan **[Usulan]**. Hasilnya dicatat bertanggal hanya sebagai "keduanya sudah" atau "belum", tanpa siapa yang menjawab "belum", tidak di aplikasi dan tidak di server (10 §7.4 baris "VL-02 sesudah keputusan 46"). **Terus terang:** di pilot dua orang, "belum" tetap dapat ditelusuri ke sumbernya, karena pemilik produk adalah salah satu dari dua penjawab sekaligus pemutus GR-3. Catatan tanpa nama hanya menjaga agar sumbernya tidak tertulis; ia bukan janji kerahasiaan (10 asas 2: tidak ada janji kerahasiaan yang tidak dapat ditegakkan). Karena itu yang melindungi penjawab adalah aturan di atas: "belum" tidak perlu dijelaskan dan tidak didesak. **(b)** Bug baru yang berat (angka uang salah, data hilang atau bocor, termasuk data pribadi yang terlihat pasangan; aplikasi tidak dapat dipakai) yang ditemukan dalam 14 hari terakhir, dari laporan pasangan dan dari pemeriksaan teknis (VL-10). Catatan bug disimpan di luar aplikasi dan server Family OS (issue tracker tim bangun atau runbook), tanpa isi rumah tangga, tangkapan layar data nyata, nama pelapor, atau hitungan per orang; reproduksinya memakai data uji di staging. Gerbang hanya membaca tanggal ditemukan, golongan (berat atau tampilan), dan status perbaikan, tidak medan perangkat (10 §7.4). Golongan yang diperdebatkan dibaca berat **[Usulan]**. **Bacaan lead [Usulan]** (atas `critic` P-4): bug berat yang ditemukan lebih dari 14 hari sebelum GR-3 tetapi belum diperbaiki juga menahan GR-3. Bug "data bocor" juga dijalankan lewat rencana insiden SEC-148. Bentuk catatan bug: 14 §5.3 "GR-3 — masukan teknis". **Bahan, bukan syarat:** jumlah hari per minggu dengan sedikitnya satu catatan oleh siapa pun, dari tanggal pencatatan entri yang terbaca keduanya (buku pribadi tidak dihitung), tanpa pemisahan per orang | (a); (a) dan (c); bahan: (b) | Di GR-3, paling cepat akhir bulan 3; (b) dihitung mundur 14 hari dari tanggal GR-3 | (a) "keduanya sudah"; (b) nol bug berat baru dalam 14 hari dan tidak ada bug berat yang belum diperbaiki (bacaan lead **[Usulan]**). Bahan frekuensi tanpa ambang. Ambang 5 dari 7 hari dicabut (keputusan 46) |
| VL-03 | **KS-2** mencatat tanpa diingatkan | Masing-masing menjawab sendiri: "Bulan ini kamu mencatat tanpa diingatkan pasanganmu?" dan "momen apa yang paling sering terlewat?" | (a) | Akhir bulan 1, 2, 3 | Bulan 3: keduanya "ya" |
| VL-04 | **KS-3** isi ringkasan bulanan | Bersama membuka ringkasan bulanan (F-FIN-020) dan mencentang: per kategori, per orang (siapa membayar apa, 06 AT-7), tabungan, aset; lalu "apakah angkanya dipercaya?" | (a), (b) | Akhir bulan 1, 2, 3 | Bulan 3: keempat butir ada dan dipercaya keduanya |
| VL-05 | **KS-4** rencana, bukan kejutan | "Bulan ini ada jalan-jalan atau pembelian tidak rutin yang datang mendadak?" dibandingkan dengan jumlah permintaan bulan itu (daftar objek, tingkat rumah tangga) | (a), (b) | Akhir bulan 1, 2, 3 | Bulan 3: tidak ada kejutan, atau setiap kejutan sudah dibahas dan pasangan tidak menganggapnya masalah |
| VL-06 | **KS-5** saling memperhatikan | Masing-masing menyatakan apakah aplikasi membantunya memperhatikan dan menanggapi keadaan pasangan, dengan satu kalimat bebas. Tidak dinilai angka, tidak dirata-rata | (a) | Akhir bulan 3, sesudah sekurangnya enam minggu `CARE` dan `AI` (A-40) | Dinilai pasangan sendiri (brief §1.1) |
| VL-07 | Usaha mencatat | (1) Ketukan dan waktu tiga alur inti di perangkat pilot (ES §3); (2) di akhir minggu kedua, momen mencatat yang dilewati dan alasannya, dicocokkan dengan tabel 04 §6 dan 03 §2.4 | (a), (c) | Minggu 2; ulang di bulan 3 | ES §3 terpenuhi; tidak ada momen yang dilewati karena "terlalu lama" |
| VL-08 | **Pagar** | Masing-masing menjawab tentang aplikasi, bukan tentang pasangan: "Pernah merasa diawasi, ditagih, atau didesak lewat aplikasi?" dan "Apa yang ingin kamu ubah?" Boleh dijawab tertulis kapan saja. **Saluran jawaban istri tidak melewati pemilik produk** (yang juga pengelola server) **[Usulan]**: ia menulisnya di tempat yang ia simpan sendiri di luar aplikasi dan di luar server pilot, dan memilih sendiri apa yang dibacakan di GR-3 (15 RISK-25) | (a) | Akhir bulan 1 dan 3 | Tidak ada "ya" yang tidak terjawab perbaikannya; satu "ya" pada SCN-30 atau SCN-32 adalah temuan menghalangi untuk GR-3 |
| VL-09 | Beban notifikasi | "Notifikasinya terasa terlalu banyak, pas, atau terlalu sedikit?" | (a) | Akhir bulan 1 dan 3 | Bukan "terlalu banyak" |
| VL-10 | Kebenaran data | Pemeriksaan invarian di server (`ops.invariant_check`, 09.4 §6) dan catatan yang hilang | (c) | Terus-menerus | **Nol** kegagalan invarian; nol catatan hilang (ES §5). Bukan usulan: syarat |

**Lebaran 2027 di dalam jendela.** Bila FS-2 berjalan melewati Ramadan dan Idulfitri 2027 (libur Idulfitri 10–11 Maret 2027 dan cuti bersama 9, 12, 15 Maret, **[Terverifikasi]** 18 §2.14; awal Ramadan 8 atau 9 Februari 2027 **[Asumsi]**; A-37), tidak satu pun ukuran di atas membandingkan besar pengeluaran antarbulan, jadi lonjakannya tidak merusak penilaian. Bulan itu justru ujian terbaik bagi KS-4 (THR, mudik, dan pembelian besar datang sebagai permintaan atau tidak). Rinciannya di §8.

### 6.2 Keputusan di tinjauan tiga bulan, dan di GR-3 (jalur publik opsional)

Baris bertanda *jalur publik opsional* hanya berlaku bila pemilik produk mempertimbangkan go public (keputusan 51); baris lain berlaku di tinjauan tiga bulan (§3.4) dan, bila jalur itu dibuka, juga di GR-3.

| Bila | Maka | Pemutus |
|---|---|---|
| Sudah sekurangnya tiga bulan sesudah GR-2; VL-01, VL-03, VL-04, VL-05 memenuhi ambang; VL-10 nol; VL-08 tanpa temuan menghalangi | **Titik akhir roadmap tercapai** (§3.4); FS-2 berlanjut sebagai pemakaian dan pemeliharaan | Pasangan; pemilik produk |
| *Jalur publik opsional:* syarat baris di atas, ditambah VL-02 (a) dan (b) terpenuhi | **Lanjut** ke GR-4 dan FS-3 | Pemilik produk |
| *Jalur publik opsional:* VL-02 (a): salah satu pasangan belum merasa nyaman dan aman memakai aplikasi | **GR-3 mundur** (keputusan 46, "bisa lebih"). Yang disebutnya tentang aplikasi, bila ia mau menyebutnya, menjadi perbaikan di FS-2. Satu "belum" cukup untuk menahan, dari pasangan mana pun; tidak ada yang menimpanya | Pasangan yang menjawab; pemilik produk menjadwalkan ulang |
| *Jalur publik opsional:* VL-02 (b): bug baru yang berat ditemukan dalam 14 hari terakhir, atau bug berat yang lebih tua tetapi belum diperbaiki (bacaan lead **[Usulan]**, atas `critic` P-4) | **GR-3 mundur** sampai perbaikannya terpasang dan 14 hari berlalu tanpa bug berat baru. Bug "data bocor" juga dijalankan lewat rencana insiden SEC-148. Catatan bugnya: 14 §5.3 "GR-3 — masukan teknis" | Tim bangun; pemilik produk |
| VL-06 "tidak" dari salah satu | Tidak menahan titik akhir, dan di jalur publik opsional tidak menahan tahap publik sendirian **[Usulan]**; `CARE` dan `AI` ditinjau ulang bentuknya bersama pasangan (di jalur publik opsional juga sebelum masuk cakupan publik) | Pemilik produk |
| VL-01 gagal di bulan 3 | **Titik akhir belum tercapai** (dan di jalur publik opsional: tidak lanjut). Masalah inti belum terjawab; penyebab dicari (catatan terlewat, angka tidak dipercaya, layar), diperbaiki di FS-2, dan VL-01 diulang di akhir bulan berikutnya | Pemilik produk |
| VL-03 gagal bagi salah satu | Perbaikan di FS-2 pada alur yang paling sering terlewat (VL-07); bila pencatat utama tetap tidak mencatat, model "dua pencatat" ditinjau ulang | Pemilik produk |
| *Jalur publik opsional:* bahan frekuensi VL-02 menurun tajam, tetapi VL-02 (a) dan (b) terpenuhi | Tidak menahan GR-3 (keputusan 46). Ditanyakan apakah aplikasi dibuka tanpa mencatat (tidak terukur tanpa analitik) dan apakah ada momen yang terlewat (VL-07); jawabannya bahan perbaikan | Pemilik produk |
| VL-08 dengan temuan menghalangi | Diperbaiki dan diuji ulang di FS-2 begitu ditemukan; tidak ada pengecualian. Juga syarat sebelum GR-4, di pintu mana pun (§3.6) | Pemilik produk; `security-privacy` |
| Satu fitur Tier 2 tidak pernah dipakai enam minggu sesudah dilepas. Dibaca dari jawaban masing-masing pasangan tentang pemakaiannya sendiri (a); untuk fitur yang objeknya memang terlihat keduanya (`SHP`, `MNT`, `PLN`, `CAL`, F-REQ-013), juga dari ada tidaknya objek yang terbaca keduanya, lewat layar aplikasi, bukan kueri basis data (b; §6.1). Status `CARE` yang tidak dibagikan tidak pernah meninggalkan perangkat (17 keputusan 36), dan sebutan anggota (F-FAM-021) hanya terlihat pembuatnya; keduanya, beserta pemakaian `AI`, **hanya** dinilai dari jawaban pemakainya sendiri, tidak pernah dari data di server, karena pengelola server adalah salah satu pasangan | Fitur diubah bentuknya atau dikeluarkan bersama pasangan (PRN-02); di jalur publik opsional juga dikeluarkan dari cakupan publik | Pemilik produk |
| VL-10 bukan nol | Server *tahan tulis* (glosarium; ADR-028 butir 3, Diusulkan): server menolak penulisan baru, ponsel tetap mencatat ke antrean dan mengirimnya sesudah tahan tulis dicabut, dan perintah pelindung (SEC-151) tetap diterima. Dicabut sesudah penyebabnya diperbaiki dan diuji (ES §5; PRN-06) | Tim bangun; pemilik produk |

### 6.3 Phase 0 untuk tahap publik: keluarga lain (jalur publik opsional)

Menjawab brief §13 Phase 0 dan R-08. Tidak memerlukan aplikasi terpasang. **Tidak dijadwalkan (keputusan 51).** Rencana ini dipertahankan agar brief §13 Phase 0 ("who to interview, how many families, which questions, which prototype tests, and what result would change the plan") tetap terjawab, tetapi wawancara baru dilakukan bila pemilik produk mempertimbangkan go public. Aplikasi untuk rumah tangga sendiri tidak memerlukannya: bagi rumah tangga itu, pilot adalah validasinya (TB-2).

| Hal | Usulan |
|---|---|
| Siapa | Pasangan menikah yang tinggal serumah di Indonesia, dari jaringan pemilik produk. Sebaran yang dicari: keduanya berpenghasilan dan hanya satu; banyak memakai uang tunai dan hampir tanpa tunai; menopang orang tua atau mertua; dengan anak; yang sudah pernah memakai aplikasi pencatat dan yang belum |
| Berapa | 8 sampai 12 rumah tangga **[Asumsi]** (A-41); berhenti lebih awal bila tiga wawancara berturut-turut tidak menambah pola baru. Hasilnya pola kualitatif, tidak pernah dilaporkan sebagai persentase pasar |
| Cara | Kedua pasangan diwawancarai **terpisah** bila mereka bersedia, karena pertanyaan tentang kendali dan privasi tidak dapat dijawab jujur di depan pasangan **[Asumsi]**; 30 sampai 45 menit; tanpa merekam isi keuangan mereka |
| Pewawancara | Pemilik produk; istrinya bila ia bersedia. Bias jaringan pertemanan diakui di laporan |
| Pertanyaan | (1) Ceritakan uang masuk bulan lalu sampai habis: siapa menerima, siapa memegang, ke mana. (2) Bagaimana kalian tahu sisa uang hari ini? (3) Pertanyaan soal uang apa yang paling sering kalian ajukan satu sama lain? (4) Bagaimana jalan-jalan atau pembelian besar diputuskan? Pernah ada kejutan? (5) Uang untuk orang tua, saudara, atau pinjaman: dicatat? (6) Apakah kalian punya uang pribadi yang tidak dilihat pasangan, dan apakah itu perlu tetap begitu? (7) Bila pasangan harus menyetujui pengeluaran di atas batas yang kalian sepakati, apakah itu terasa membantu atau mengekang? (8) Aplikasi apa yang pernah dipakai, dan mengapa berhenti? (9) Siapa yang akan mencatat, dan seberapa sering kalian sanggup? (10) Apakah kalian mau membayar, dengan cara apa (per bulan, per tahun, sekali), dan berapa yang terasa wajar? |
| Uji prototipe | Contoh layar statis atau dapat diklik dari arah C (yang dibuat untuk 08 §19.8), tanpa data nyata dan tanpa memasang aplikasi: (a) uji lima detik pada Beranda, "uang kalian sisa berapa dan siapa yang pegang?"; (b) tugas: catat belanja tunai; catat serah terima lalu konfirmasi; ajukan jalan-jalan dan tanggapi dengan usul balik; (c) tiga pilihan model bisnis dengan harga uji (§7.5), "mana yang kalian pilih dan mengapa". Jawaban harga adalah preferensi yang dinyatakan, bukti lemah |
| Kedudukan terhadap Q-209 | **Diputuskan** (keputusan 47, bagian jalur publik opsional menurut keputusan 51; diterima atas saran lead: "Bukan, dengan 4 syarat (Saran)"; Q-215 selesai): wawancara dan uji prototipe tanpa layanan dan tanpa data nyata **bukan** "pengguna di luar rumah tangga pilot", **selama empat syarat 10 §7.4 terpenuhi**; keempatnya kini aturan dan dicentang dengan tanggal sebelum wawancara pertama (daftar periksa di bawah). Begitu satu syarat gugur, misalnya seorang peserta memasang build atau mengisi data, Q-209 berlaku (advokat menjawab L-01, L-14, L-24 lebih dulu). Apakah catatan wawancara itu sendiri pemrosesan di luar pengecualian rumah tangga tetap L-01 **[Perlu validasi]**. **Pagar** (privasi peserta) |

**Daftar periksa sebelum wawancara pertama** (10 §7.4, keputusan 47). Dicentang dengan tanggal oleh pewawancara sebelum wawancara pertama, dan diperiksa ulang bila cara wawancara berubah (misalnya prototipe mulai dibuka daring). Disimpan bersama rencana wawancara, di luar repositori dan server, tanpa nama peserta.

| # | Syarat (10 §7.4) | Diperiksa dengan | Tanggal dicentang |
|---|---|---|---|
| (1) | Tidak ada build terpasang, akun, atau data peserta yang masuk server | Tidak ada undangan TestFlight, APK, atau akun untuk peserta; prototipe tidak terhubung ke server pilot | … |
| (2) | Prototipe tanpa data nyata dan, bila dibuka daring, tanpa analitik atau rekaman sesi | Isi prototipe hanya angka contoh; pengaturan alat prototipe tanpa analitik dan rekaman sesi | … |
| (3) | Catatan wawancara tanpa nama di samping jawaban dan tanpa angka keuangan rumah tangga peserta, disimpan di luar repositori dan server, dihapus sesudah dirangkum | Templat catatan memakai kode peserta, tanpa kolom nama atau angka; tempat simpan di luar repositori dan server; tanggal hapus dicatat sesudah rangkuman | … |
| (4) | Peserta diberi tahu tujuan dan haknya menolak di awal | Kalimat pembuka tertulis dibacakan sebelum pertanyaan pertama | … |

**Hasil yang mengubah rencana** (ambang **[Usulan]**; yang memutuskan perubahan tetap pemilik produk):

| Hasil wawancara | Perubahan |
|---|---|
| Sebagian besar rumah tangga hampir tidak memakai uang tunai dan memakai satu rekening bersama | Custody tetap, tetapi tidak dijadikan pembeda utama di materi publik; urutan halaman depan ditinjau 08 |
| Sebagian besar menganggap persetujuan antarpasangan mengekang | Permintaan tetap ada tetapi ambang tidak ditawarkan dalam penyiapan awal; positioning bergeser ke "rencanakan bersama" |
| Pola menopang orang tua jarang | Kiriman keluarga tetap (Tier 1), tidak dijadikan pesan pemasaran |
| Hampir tidak ada yang mau membayar dalam bentuk apa pun | Model bisnis ditinjau ulang sebelum GR-5; tahap publik tetap dapat berjalan dengan inti gratis bila biayanya ditanggung (§7.3) |
| Kebutuhan struktur lain (anak dengan akun, orang tua serumah) muncul berulang | FS-5 dinaikkan urutannya sesudah FS-4 |
| Tugas prototipe gagal di lebih dari sepertiga peserta | Layar yang gagal diubah sebelum beta |

### 6.4 Beta tertutup (FS-3; jalur publik opsional)

Tidak dijadwalkan (keputusan 51); keputusan 48 bagian jalur ini. Rumah tangga di luar pilot, sesudah GR-4, sekurangnya 8 minggu **[Usulan]**, dengan ukuran yang sama dengan VL-01 sampai VL-09 dan percakapan akhir bulan yang sama. Pengujian pagar yang sebenarnya terjadi di sini, karena pengelola server bukan anggota rumah tangga mereka (glosarium *Pengelola server*). Distribusi lewat TestFlight eksternal dan jalur *closed testing* Google Play, bukan APK langsung **[Usulan]**.

**Ukuran beta** **[Usulan]**. Untuk validasi, sekurangnya 5 rumah tangga. Syarat Play mungkin lebih besar: akun pengembang **pribadi** yang dibuat sesudah 13 November 2023 memerlukan sekurangnya 12 penguji yang ikut tanpa putus 14 hari di *closed testing* sebelum rilis produksi (**[Terverifikasi]** `architect`, 14 §6.7; apakah akun organisasi dikecualikan **[Perlu validasi]**, permintaan ke `researcher` di 14 §12.2). Yang dihitung adalah orang dewasa yang memakai perangkat ber-Google Play, bukan rumah tangga (18 §2.4 butir 6 **[Usulan]** `researcher`; Google tidak menyatakan eksplisit bahwa penguji harus memakai Android): rumah tangga dua dewasa yang keduanya Android menyumbang dua penguji, rumah tangga campuran (satu iPhone, satu Android) satu, rumah tangga yang keduanya iPhone nol. Karena itu, bila akunnya pribadi, jumlah rumah tangga ditetapkan dari jumlah pemakai Android, bukan sebaliknya: rekrut sampai **sekurangnya 14 penguji Android** (12 ditambah 2 cadangan bila ada yang keluar sebelum 14 hari habis) **[Usulan]**, tanpa batas atas 10 rumah tangga. 14 hari itu jatuh di dalam 8 minggu beta. **Akunnya pribadi** (keputusan 48; Q-221 selesai), jadi aturan 14 penguji Android ini berlaku; sebelum akun dibuat, dukungan Play Console ditanya apakah aplikasi keuangan rumah tangga wajib memakai akun organisasi (GR-4 butir 5). Keputusan 42 tidak mengubahnya, karena klien Android hadir paling lambat sebelum jalur pengujian Google Play dibuka (GR-4 butir 8; QG-31 (b)). Rumah tangga pilot sendiri, bila tetap di beta, menyumbang paling banyak satu penguji Android (pemilik produk di Galaxy A55).

GR-5 mensyaratkan: tidak ada temuan menghalangi VL-08; VL-10 nol; sebagian besar rumah tangga beta memenuhi VL-01 di bulan kedua; syarat 12 penguji × 14 hari akun Play pribadi sudah dipenuhi. Rumah tangga beta pertama baru masuk sesudah butir 10 §4.2 terpenuhi (GR-4 butir 7).

### 6.5 Batas rencana validasi

| Batas | Akibat | Mitigasi |
|---|---|---|
| Pilot dijalankan oleh salah satu pasangan, yang juga pengelola server; tidak ada pihak ketiga | Jawaban jujur pasangan yang lain tentang pagar tidak dapat dijamin | Pertanyaan tentang aplikasi, bukan tentang pasangan; boleh tertulis, di tempat yang ia simpan sendiri dan dibacakan atas pilihannya (VL-08). Beta (§6.4), uji pagar yang sebenarnya, hanya ada di jalur publik opsional; tanpanya batas ini diterima di jalur rumah tangga, dan yang menjaga adalah saluran VL-08 |
| Satu rumah tangga | Tidak ada hasil pilot yang boleh disebut berlaku umum (PRN-04) | Aplikasi untuk rumah tangga ini sendiri (keputusan 51) tidak memerlukan hasil yang berlaku umum. Klaim umum hanya sesudah wawancara §6.3 dan beta §6.4, di jalur publik opsional |
| Tanpa analitik | "Membuka tanpa mencatat" tidak terukur | Ditanyakan langsung (§6.2) |
| Pernyataan pemilik produk di checkpoint S4, kata-katanya: "istri saya juga programer dia yang nyruh ini kocak" (17 keputusan 36; tidak diverifikasi dokumen mana pun) | Bila istrinya ikut memprakarsai aplikasi, A-25 gugur dan KS-2 kehilangan "ujian adopsi" yang dimaksud 03 §2.1 | Ditanyakan (Q-218); pencatat utama yang pelupa (brief §1.1) tetap ujiannya |

---

## 7. Model bisnis

Rumah tangga pemilik produk tidak dimonetisasi. Rekomendasi ini hanya untuk jalur publik opsional (FS-3, FS-4; keputusan 51): ditulis karena brief §1 "Business model" memintanya, tidak dijadwalkan. Brief §1: "Evaluate freemium per household, family subscription, and one-time purchase; recommend one with reasoning" (A-03; Q-36).

### 7.1 Batasan dari dokumen lain

| Batasan | Sumber | Akibat bagi model bisnis |
|---|---|---|
| Tidak ada SDK iklan, analitik perilaku, atau perekam sesi | SEC-49 | Model iklan tertutup |
| Data kesehatan tidak dipakai untuk iklan, pemasaran, atau penambangan data | SEC-145; Apple 5.1.3 (**[Terverifikasi]** 18 §2.4) | Sama |
| Data pribadi dan sensitif di Google Play dilarang dijual | **[Terverifikasi]** 18 §2.4 | Menjual data tertutup |
| Aplikasi mencatat uang, tidak memegang atau memindahkan dana | Brief §1 "Money movement scope" | Pembayaran langganan tidak boleh membawa produk ke gateway pembayaran; penagihan lewat toko **[Usulan]** |
| Dana bersama terlihat setiap anggota dewasa penuh; tidak ada peran yang menutupnya | PRN-13 | Tidak boleh ada tingkat berbayar yang membatasi apa yang dilihat pasangan |
| Simetris; tidak ada yang bertindak atas nama orang dewasa lain | PRN-09, PRN-10 | Membayar tidak boleh memberi hak, peran, atau pantauan apa pun |
| Satuan isolasi adalah rumah tangga; satu orang dapat di beberapa rumah tangga | Glosarium; 03 §4.3; PRN-03 | Hak akses berbayar melekat pada rumah tangga, bukan pada akun; tidak bocor ke rumah tangga lain |
| Nilai produk perlu kedua pasangan aktif | PRN-01; KS-1, KS-2; 03 §2.5 | Pasangan kedua tidak boleh menemui dinding bayar untuk bergabung |
| Catatan hanya bertambah; hak hapus dan ekspor | PRN-16; 10 siklus data; UU 27/2022 (18 §2.1) | Langganan yang berhenti tidak pernah mengunci catatan |

### 7.2 Bukti pasar

Semua harga **[Terverifikasi]** primer di 18 §1.1–§1.3.7, tanggal akses 11 November 2026, kecuali disebut lain.

| Aplikasi | Model | Harga | Catatan |
|---|---|---|---|
| SakuSaku (Indonesia) | Freemium; dasar gratis tanpa iklan | Pro Rp85.000/tahun; uji coba 30 hari | Mode Pasangan |
| MoneyKita (Indonesia) | Langganan bulanan | Solo Rp49.000/bulan; Pro Rp79.000/bulan, satu langganan Pro memberi pasangan akses Pro | 18 mencatat halaman tidak konsisten soal pasangan perlu "upgrade" |
| CoupleWallet (dilokalkan) | Gratis dengan iklan | — | Label privasi: pelacakan dan iklan pihak ketiga |
| Cozi | Freemium, langganan keluarga tahunan | Gold USD 39/tahun; Max USD 79/tahun | Tanpa fitur keuangan |
| TimeTree | Freemium | Premium USD 4,49/bulan atau USD 44,99/tahun (bervariasi per wilayah) | Tanpa fitur keuangan |
| Monarch | Satu langganan untuk beberapa anggota rumah tangga | Harga tidak tampil | Menambah anggota tanpa biaya |

**Toko (18 §2.13.1, R-15, fase D).** Semua **[Terverifikasi]** di 18 kecuali disebut lain:

| Hal | Apple | Google Play |
|---|---|---|
| Komisi langganan | Tahun pertama tiap pelanggan pengembang menerima 70%, sesudah satu tahun 85%; peserta Small Business Program (hasil ≤ USD 1 juta per tahun, harus mendaftar) 85% sejak awal, yaitu komisi 15% | 15% untuk langganan berulang di Indonesia bagi pengembang yang terdaftar di tingkat 15% |
| Wajib pembelian dalam aplikasi | Ya (Guideline 3.1.1); hak lintas platform boleh asal Plus juga dapat dibeli di aplikasi iOS (3.1.3(b)) | Ya (Payments policy); aturan lintas platform tidak ditemukan **[Perlu validasi]** |
| PPN pelanggan | Dipungut dan disetor Apple; harga sudah termasuk pajak | Dipungut dan disetor Google, 11% |
| Berbagi | Family Sharing: sampai lima anggota keluarga Apple; **setelah diaktifkan tidak dapat dibatalkan** | Family Library tidak berlaku untuk pembelian dalam aplikasi; langganan tidak disebut **[Perlu validasi]** |

Bacaan `researcher` **[Usulan]**: komisi efektif sekitar 15% di kedua toko bagi pengembang kecil yang mendaftar (Apple tanpa pendaftaran 30% di tahun pertama tiap pelanggan). Tingkat harga rupiah kedua toko belum terbaca **[Perlu validasi]**. **Bukti kesediaan membayar keluarga Indonesia tetap tidak ada** (R-16 terbuka; 18 §1.5: survei untuk kategori ini tidak ditemukan; menurut media, hanya sebagian kecil pengguna internet berlangganan aplikasi digital, dan kebanyakan hiburan **[Perlu validasi]**).

### 7.3 Biaya

| Butir | Angka | Label |
|---|---|---|
| Pilot per bulan | Rp260.667 + PPN bila berlaku, ditambah USD 9,57 + email (paling tinggi USD 18,25 + email bila batas `AI` tercapai); Rp244.000 di antaranya bergantung pada S4-2 | 09.4 §7.1 baris "Jumlah" (dihitung ulang lead S4); harga sebagian **[Terverifikasi]**, sebagian **[Perlu validasi]** |
| Tahap publik, sekitar 500 rumah tangga | Rp400.000 sampai Rp500.000 per bulan + email, yaitu **Rp800 sampai Rp1.000 per rumah tangga per bulan** (Rp9.600 sampai Rp12.000 per tahun) | Angka 09.4 §7.2 **[Asumsi]**; pembagiannya hitungan dokumen ini |
| Tahap publik, sekitar 5.000 rumah tangga | Rp1,5 juta sampai Rp2 juta per bulan + email, yaitu **Rp300 sampai Rp400 per rumah tangga per bulan** (Rp3.600 sampai Rp4.800 per tahun) | Sama |
| Tetap, tidak tumbuh dengan pengguna | Apple USD 99/tahun; Google Play USD 25 sekali; `AI` R1 "tidak tumbuh dengan jumlah rumah tangga: satu kumpulan dipakai semua" | **[Terverifikasi]** 18 §5.1.1, §5.2.3; 09.4 §7.2 |
| Tumbuh dengan pengguna | Email per pesan; `AI` R2 bila kelak dipilih (per panggilan); dukungan pengguna | 09.4 §7.2 |
| Waktu pengembang | Biaya terbesar; tidak diperkirakan dokumen mana pun | — |

**Bacaan [Usulan]** (A-42). Biaya infrastruktur per rumah tangga kecil sekali dibanding harga pesaing; biaya yang sesungguhnya adalah pemeliharaan berkelanjutan (pembaruan OS tahunan, server yang harus terus hidup untuk sinkron dan chat, tinjauan hukum). Ilustrasi saja: pada 5.000 rumah tangga, Rp18 juta sampai Rp24 juta per tahun ditutup oleh sekitar 121 sampai 162 rumah tangga yang membayar Rp149.000 per tahun, **sebelum** PPN yang dipungut toko, komisi toko sekitar 15% (§7.2), dan pajak penghasilan pengembang (L-40), dan tanpa waktu pengembang.

### 7.4 Penilaian tiga pilihan

| Kriteria | Freemium per rumah tangga | Langganan keluarga (tanpa tingkat gratis; hanya uji coba) | Beli sekali |
|---|---|---|---|
| Menutup biaya berulang (server, sinkron, chat, pembaruan OS) | Ya, dari rumah tangga yang membayar | Ya | **Tidak**: pendapatan sekali untuk biaya yang tidak pernah berhenti **[Usulan]** |
| Pasangan kedua dapat bergabung tanpa dinding bayar | Ya, bila inti gratis | Hanya selama uji coba | Bergantung apakah pembelian di satu toko dikenali di toko lain; rumah tangga pilot sendiri memiliki iPhone, iPad, dan Android (brief §1; keputusan 42). Cara tokonya **[Perlu validasi]** (R-15) |
| Simetri (PRN-09) | Ya, bila hak melekat pada rumah tangga dan pembayar tidak mendapat apa pun selain status langganan | Sama | Sama |
| Tidak mengunci catatan | Ya: inti tetap gratis, jadi berhentinya langganan hanya menghapus tambahan | Berisiko: sesudah langganan berhenti, apa yang terjadi pada catatan harus dirancang agar tidak menjadi sandera | Ya |
| Kecocokan dengan pasar Indonesia yang terlihat | Pola SakuSaku (dasar gratis, Pro tahunan murah) | Pola MoneyKita (bulanan, pasangan ikut) | Tidak ada pembanding di 18 |
| Hambatan masuk di pasar yang sensitif harga **[Asumsi]** | Rendah | Tinggi | Sedang |
| Kerumitan bangun | Sedang: penagihan toko dua platform, hak akses per rumah tangga di server, garis "gratis lawan Plus" | Sedang: sama, tanpa garis gratis | Rendah di toko, tetapi tetap perlu pengenal hak lintas platform |
| Risiko prinsip | Godaan menaruh hal penting di balik bayar; dijaga daftar "tidak pernah berbayar" di §7.5 | Rumah tangga yang tidak membayar tidak mendapat apa pun | Dorongan menjual ulang "versi 2" |

### 7.5 Rekomendasi

**Freemium per rumah tangga** **[Usulan]**: inti produk gratis selamanya untuk satu rumah tangga, ditambah satu langganan **Plus per rumah tangga** (tahunan, dengan pilihan bulanan), dijual lewat penagihan App Store dan Google Play, dengan hak aksesnya disimpan server per rumah tangga. Ini memadukan pilihan pertama dan kedua brief: bentuk tagihannya langganan keluarga, tetapi ada tingkat gratis yang memuat masalah inti.

**Alasannya, singkat.** (1) Nilai produk baru muncul bila kedua pasangan aktif (PRN-01, KS-2); dinding bayar di pintu masuk menghalangi orang kedua, yang justru ujian adopsi (03 §2.1). (2) Biaya infrastruktur per rumah tangga kecil (§7.3), jadi tingkat gratis terjangkau. (3) Beli sekali tidak menutup biaya layanan yang terus hidup. (4) Pola harga Indonesia yang terlihat untuk kategori ini adalah dasar gratis dengan tingkat berbayar tahunan murah (SakuSaku). (5) Iklan dan penjualan data tertutup oleh dokumen 10 dan kebijakan toko (§7.1).

**Tidak pernah berbayar** (pagar model bisnis; diperiksa di setiap usulan isi Plus):

| Tidak pernah di balik bayar | Dasar |
|---|---|
| Semua fitur Tier 1: angka utama dan custody, mencatat, serah terima, permintaan dan persetujuan, chat rumah tangga, ringkasan bulanan, pengingat tagihan | Masalah inti; KS-1 sampai KS-4 |
| Bergabung, keluar dengan aman, hapus akun, halaman Bantuan, sembunyikan untuk saya, laporkan, kunci aplikasi | PRN-10; 10 §4; brief §8 |
| Apa yang dilihat pasangan atas dana bersama | PRN-13 |
| Ekspor data milik sendiri | 10 siklus data; UU 27/2022 |
| Jumlah anggota rumah tangga | Biaya per anggota menghukum keluarga yang menopang orang tua (brief §4.1 struktur 8) **[Usulan]** |

**Calon isi Plus** **[Usulan]**, diuji di wawancara dan beta, diputuskan pemilik produk di GR-5: laporan lintas periode dan tren rumah tangga (F-FIN-022), ruang foto yang lebih besar, pendamping `AI` versi ringan, templat musiman (IDEA-08). Apakah `AI` dan `CARE` pantas berbayar adalah pertanyaan nilai, bukan biaya (R1 tidak tumbuh dengan jumlah rumah tangga), dan dibuka di Q-36.

**Aturan langganan yang diusulkan** (bukan aturan uang; diminta ke `architect` dan `security-privacy`, §12.2): (a) Plus melekat pada rumah tangga; setiap anggota dewasa melihat "Plus aktif sampai tanggal …", tanpa nama pembayar **[Usulan]**; (b) pembayar tidak mendapat peran, izin, wewenang, atau pantauan; (c) bila pembayar keluar atau berhenti membayar, Plus berakhir di akhir masa bayar dan tidak ada catatan yang hilang; (d) Plus di satu rumah tangga tidak memberi apa pun di rumah tangga lain dan tidak terlihat di sana; (e) **Family Sharing Apple tidak diaktifkan untuk Plus** **[Usulan]**, sejalan dengan 18 §2.13.3 butir 3: keluarga Apple pembeli belum tentu anggota rumah tangga Family OS (hak bisa bocor ke orang di luar rumah tangga, berlawanan dengan (d)), pengaktifannya tidak dapat dibatalkan, dan hak per rumah tangga di server sudah menjangkau setiap anggota.

**Harga uji** **[Usulan]**, hanya untuk wawancara (§6.3): Rp99.000, Rp149.000, dan Rp249.000 per tahun per rumah tangga. Jangkarnya SakuSaku Pro Rp85.000/tahun dan MoneyKita Pro Rp79.000/bulan (Rp948.000 setahun). Bukan harga yang direkomendasikan.

**Yang dapat membatalkan rekomendasi ini:** aturan toko yang membuat langganan per rumah tangga lintas platform tidak mungkin (Apple mengizinkannya, 3.1.3(b); aturan Google belum ditemukan, R-15 sisa); wawancara dan beta yang menunjukkan hampir tidak ada yang mau membayar (§6.3); kewajiban hukum penjualan langganan yang belum diteliti (L-40).

---

## 8. Tanggal yang ditetapkan pihak luar

| Tanggal | Hal | Akibat pada roadmap | Label |
|---|---|---|---|
| 12 Januari 2027 | Semua aplikasi wajib menjawab deskriptor rating usia Apple baru ("Shared Location", "Dating and Relationships"); jawaban wajib sebelum kiriman pertama sesudah 16 November 2026 | Bila catatan aplikasi di App Store Connect dibuat sebelum tanggal itu, kuesioner dijawab saat itu juga, walaupun pilot hanya memakai TestFlight internal **[Usulan]**; apakah kewajibannya berlaku bagi aplikasi yang hanya di TestFlight internal belum diketahui (R-18) | Tanggal **[Terverifikasi]** 18 §2.12.3 |
| 16 atau 17 Januari 2027 (hitungan peneliti atas salinan teks); 15 Januari 2027 menurut metadata Hukumonline | PP 33/2026 mulai berlaku | Tidak mengubah GR-2: DPIA sudah wajib sebelum data nyata (Q-208) | **[Perlu validasi]** (17 R-01; salinan teks belum dicocokkan dengan pindaian resmi) |
| 10–11 Maret 2027 (libur nasional); cuti bersama 9, 12, 15 Maret | Idulfitri 1448 H (SKB tiga menteri, Setneg 15 September 2026); awal Ramadan belum ditetapkan Menteri Agama, perkiraan 8 atau 9 Februari 2027 | (1) Bila FS-2 mencakupnya: templat IDEA-08 dilepas sebelum Ramadan; THR dan angpau sudah Tier 1 (FR-FIN-040/AC-14); bulan itu ujian KS-4 (§6.1). (2) Bila FS-2 dimulai sesudahnya: IDEA-08 dijadwalkan sebelum Ramadan berikutnya di FS-2 yang berlanjut (keputusan 51) | Idulfitri dan cuti bersama **[Terverifikasi]** 18 §2.14 (halaman Setneg; naskah SKB tidak dibuka); awal Ramadan **[Asumsi]** (11 §11 butir 2; A-37; R-18) |
| 90 hari sesudah setiap unggahan | Build TestFlight kedaluwarsa | Rilis pemeliharaan bila 60 hari tanpa rilis, batas keras 80 hari sejak unggahan sebelumnya (ADR-028 butir 4, keputusan 44); di bawah batas 90 hari ADR-022 butir 1 (keputusan 32) (§3.4) | **[Terverifikasi]** 18 §2.12.6 |
| 31 Agustus 2026 (sudah lewat) | Google Play: aplikasi baru wajib menargetkan API 36 | Berlaku di GR-4 bila Play dipakai; di pilot, klien Android (sesudah GR-2) dipasang sebagai APK | **[Terverifikasi]** 18 §5.1.1 |

---

## 9. Konsep dokumen 11 dan butir 17 yang diserahkan ke dokumen ini

| Butir | Penempatan |
|---|---|
| IDEA-01, IDEA-02, IDEA-03; THR yang diberikan (Q-185) | **FS-1**, sebagai kriteria penerimaan FR induk Tier 1 (05 §6). Q-187 bagian fase ditutup |
| IDEA-15 tinjauan bulanan, IDEA-19 butuh tenang | **FS-2** gelombang T2-3 (§3.4); FR-nya ditulis sebelum dibangun (bukan Must) |
| IDEA-08 musim Lebaran | **FS-2** bila jendela pilot mencakup Ramadan 2027; bila tidak, sebelum Ramadan berikutnya di FS-2 yang berlanjut (§8) |
| IDEA-14 tujuh hari ke depan (D-69) | **Tercakup tanpa fitur baru**: Beranda bagian 7 "Minggu ini" (08) menampilkan tagihan tujuh hari ke depan, cicilan, dan persetujuan terbuka. Yang belum tercakup (tugas yang diterima, tindak lanjut sakit yang dibagikan) ditimbang di tinjauan tiga bulan |
| IDEA-13 peta urusan rumah (D-69) | **Menunggu bukti** (Could). *Tanggung jawab* sudah ada di model (06 §3.1) dan 08 menampilkan penanggung jawab yang disepakati. Ditimbang di tinjauan tiga bulan bila VL-00 dan VL-05 menunjukkan "siapa bayar apa" masih ditanyakan |
| Konsep Could lain dan Tier 3 dokumen 11 | Menunggu hasil pilot (D-69); yang ditolak dokumen 11 tetap tidak dibangun |
| Q-186 sisi S5 (anak 12+) | §3.4: tidak ada yang dijadwalkan untuk anak selama pilot; akun dan profil anak di FS-5 |
| A-37 (bulan mulai pilot) | Tetap **[Asumsi]**; roadmap ini berlaku untuk kedua kemungkinan (§8) |

---

## 10. Tantangan terhadap brief

| # | Brief | Tantangan | Usulan |
|---|---|---|---|
| TB-1 | §13 "RICE scores plus MoSCoW labels" (sebagai contoh) | RICE memaksa angka *Reach*, *Impact*, *Confidence* yang tidak punya sumber | MoSCoW dengan syarat yang dapat diperiksa (§2) |
| TB-2 | §13 Phase 0 sebelum Phase 1 | Bagi tahap pertama, pilot adalah validasinya; menahan pembangunan untuk wawancara keluarga lain tidak menambah apa pun bagi rumah tangga pilot | Phase 0 pilot sempit (§3.1); Phase 0 publik di jalur publik opsional (§6.3; keputusan 51) |
| TB-3 | §13 Phase 2 "Core Family Operations" | Versi ringan modul operasi rumah sudah masuk MVP lewat Tier 2; menambah modul sebelum rumah tangga lain memakai MVP adalah cakupan tanpa bukti | Phase 2 = membawa MVP ke publik (FS-3, FS-4), kini jalur publik opsional (keputusan 51) |
| TB-4 | §1 "only after the pilot shows sustained daily use" (semula "Public launch only after ..."; sejak keputusan 51 didahului "if the owner later decides on it") | "Sustained daily use" tanpa definisi tidak dapat menjadi gerbang | **Dijawab lewat keputusan 46** (Q-213 selesai): batas bawah tiga bulan dan kedua syarat (nyaman dan aman; tanpa bug baru) diputuskan pemilik produk sendiri; arti bug (14 hari tanpa bug berat) diterima atas saran lead; rumusan gabungan (paling cepat tiga bulan sesudah GR-2, (a) dan (b)) adalah bacaan lead (§3.5; VL-02). Usulan lama (5 dari 7 hari) digantikan. Gerbang tidak lagi menguji "daily": 17 K-68 |
| TB-5 | §1 Business model: tiga pilihan terpisah | Freemium dan langganan keluarga bukan pilihan yang saling meniadakan | Freemium per rumah tangga dengan Plus sebagai langganan rumah tangga (§7.5), kini jalur publik opsional |
| TB-6 | §13 Phase 0 sampai Phase 5 sebagai urutan fase yang dijalankan | Keputusan 51: aplikasi untuk rumah tangga sendiri, go public tidak direncanakan. Menjadwalkan fase publik berarti merencanakan pekerjaan yang tidak diminta (advokat, PSE, toko, beta, wawancara keluarga lain); menghapusnya berarti merancang ulang bila kelak go public | Roadmap yang dijadwalkan berakhir di FS-2 (pemakaian sehari-hari). Phase 2 (GR-3, FS-3, GR-5, FS-4) dan Phase 0 publik (§6.3) menjadi jalur publik opsional, tertulis tetapi tidak dijadwalkan; GR-4 gerbang bersyarat yang juga dilewati sebelum login orang tua (keputusan 52). Phase 3 sampai 5 (FS-5 sampai FS-8) mengikuti pemicu dari rumah tangga sendiri atau dari jalur opsional (§3.8) |

TB-1 sampai TB-3, TB-5, dan TB-6 tidak bertentangan dengan bunyi brief: brief §13 sendiri mengizinkan struktur lain ("or justify a better one"), rencana validasi Phase 0 tetap tertulis lengkap di §6.3 (hanya tidak dijadwalkan), dan brief §1 serta §1.1 sudah memuat keputusan 51. Karena itu keputusan 51 tidak memerlukan baris K baru di 17. Jawaban TB-4 lewat keputusan 46 berbeda dengan bunyi brief §1: tiga bulan brief §1.1 dipertahankan sebagai batas bawah, tetapi gerbang tidak lagi menguji "daily". Perbedaan itu dicatat di 17 K-68, dan keputusan menang.

---

## 11. Angka di dokumen ini yang harus dihitung ulang lead

| # | Angka | Masukan | Letak |
|---|---|---|---|
| 1 | 132 = 55 + 19 + 58 (F-FAM-022 baru, keputusan 54; semula 131 = 55 + 18 + 58); Tier 3 58 = 0 + 12 + 10 + 2 + 3 + 18 + 13 (fase D: F-FIN-023 dari FS-3 ke "Menunggu bukti"; semula 1 + … + 17); daftar ID per penempatan cocok dengan 04 §3 | 04 §1, §3; §3.8, §4.1 | §4.1 |
| 2 | 198 Must = 190 `FS-1` + 8 `FS-2` | 05 §6; `grep` baris `Tier / MoSCoW` | 05 §6 |
| 3 | Rp400.000 ÷ 500 = Rp800; Rp500.000 ÷ 500 = Rp1.000; ×12 = Rp9.600; Rp12.000 | 09.4 §7.2 | §7.3 |
| 4 | Rp1.500.000 ÷ 5.000 = Rp300; Rp2.000.000 ÷ 5.000 = Rp400; ×12 = Rp3.600; Rp4.800 | 09.4 §7.2 | §7.3 |
| 5 | Rp18 juta = 1,5 juta × 12; Rp24 juta = 2 juta × 12; 18.000.000 ÷ 149.000 ≈ 120,8 → 121; 24.000.000 ÷ 149.000 ≈ 161,1 → 162 | §7.3 baris 3; harga uji | §7.3 |
| 6 | MoneyKita Pro Rp79.000 × 12 = Rp948.000 | 18 §1.2 | §7.5 |
| 7 | Rp260.667 + USD 9,57; maks USD 18,25; Rp244.000 bergantung S4-2 (dikutip, sudah dihitung ulang S4) | 09.4 §7.1 | §7.3 |
| 8 | Yang mengikat (ADR-028 butir 4, keputusan 44): 60 < 80 < 90; 90 − 80 = 10 hari sisa. ADR-022 butir 1: sasaran 60 < batas 90 hari masa build | 18 §2.12.6; ADR-022 butir 1; ADR-028 butir 4 | §3.4, §8 |
| 9 | Beta (jalur publik opsional), akun Play pribadi: 12 penguji + 2 cadangan = 14 penguji Android; rumah tangga dua dewasa Android 2, campuran 1, dua iPhone 0; 5 rumah tangga dua dewasa Android = 10 < 12 | 14 §6.7 | §6.4 |

---

## 12. Pertanyaan, asumsi, dan permintaan

### 12.1 Butir baru di dokumen 17

Q-210 (klien web untuk tahap publik) beserta catatan keputusan web; Q-211 (siapa memakai perangkat mana); Q-212 (urutan Tier 2 di pilot); Q-213 (definisi pemakaian harian yang bertahan); Q-214 (segmen rilis publik pertama); Q-215 (uji prototipe tanpa layanan terhadap Q-209); Q-216 (isi Plus dan apakah `AI`/`CARE` boleh berbayar); Q-217 (aturan langganan per rumah tangga); Q-218 (peran istri dalam memprakarsai aplikasi, A-25); A-40, A-41, A-42; R-15 sampai R-18; L-40. Status Q-36, Q-37, Q-186, Q-187, R-08, D-69, A-37 diperbarui di 17 bagian 2.16.

**Fase D:** Q-219 (memperketat irama rilis ke 80 hari, ADR-028 butir 4); Q-220 (arti "pas mau deploy" untuk S4-2); Q-221 (jenis akun Google Play); Q-222 (siapa mengoperasikan server bila pemilik produk berhalangan lama).

**Sesudah S5:** keputusan 42; Q-37 selesai; Q-211 terjawab; A-19 diperbarui; K-66 (brief §1 "Pilot devices"); K-67 (ES §3); Q-223 (kalimat tablet di tablet Android).

**Sebelum S6:** keputusan 43 sampai 50. Selesai: Q-213 (46), Q-215 (47), Q-219 (44), Q-220 (45), Q-221 (48), K-66 dan K-67 (50). Tetap terbuka dengan sengaja: Q-222 (pemilik produk dan istrinya, sebelum GR-2; kini GR-2 butir 8).

**Keputusan 51:** keputusan 51 dicatat di 17; status butir yang mengandaikan peluncuran publik terencana diperbarui di 17 bagian 2.18 menjadi "jalur publik opsional (keputusan 51)", tanpa menutupnya; L-01 diperluas (18 §2.1 butir tinjauan 1; 10 §3.3 butir 3). Baru: Q-224 dan Q-225 (P-51-1, P-51-2 `security-privacy`; terbuka, keputusan yang berlaku tetap sampai pemilik produk menjawab), Q-226 (P-51-3; **selesai lewat keputusan 52**: orang tua dihitung pengguna luar, jadi GR-4 dipisahkan dari jalur publik, §3.6), Q-227 (P-51-4: SEC-25 dan SEC-18 di jalur keluarga; terbuka; bila B atau C dipilih, isi FS-2 atau FS-1 dokumen ini berubah) dan R-19 (TestFlight internal sebagai jalur tetap lawan Apple Guideline 2.2). Tidak ada baris K baru (§10, TB-6). Q-222 tetap terbuka.

**Keputusan 53 sampai 55** (12 November 2026, ketiganya diterima atas saran lead): Q-225 selesai (53; tinjauan `security-privacy` sebelum klien Android memegang data nyata dan setahun sekali; §3.4, §3.6); Q-227 selesai (54; SEC-18 dan SEC-25 sebagai pasangan opsional per akun di FS-2, F-FAM-022 baru di gelombang T2-1, GR-4 butir 7 dan daftar keamanan FS-3 disesuaikan; §3.4, §4.1); Q-228 selesai (55; TestFlight internal, pindah ke App Store unlisted bila Apple menghentikannya; §3.4). Q-224 tetap terbuka (pemilik produk: "saya tanya besok"). Susulan `security-privacy`: pendaftaran tertutup SEC-158 sebagai ketergantungan FS-1 (§3.2).

### 12.2 Permintaan kepada pemilik dokumen lain (lewat lead)

| Kepada | Permintaan |
|---|---|
| `architect` (14) | Urutan workstream yang memungkinkan GR-2 dengan iOS dan iPadOS saja (keputusan 42; semula menurut ketiga pilihan Q-37), lalu klien Android sesudah GR-2 dengan daftar teknis syarat ulang §3.4 sebelum APK memegang data nyata; daftar teknis GR-2 (§3.3 butir 2–5) sebagai quality gate; irama rilis pilot menurut ADR-028 butir 4 (60/80 hari; diterima, keputusan 44); hak akses Plus per rumah tangga di server dan penagihan toko dua platform di FS-3; jalur beta (TestFlight eksternal, jalur pengujian Play) |
| `architect` (fase D) | (1) ADR-028 butir 4 dan 14 §6.3, QG-23: hitungan 80 hari dari unggahan ke TestFlight, bukan dari pemasangan (18 §2.12.6). (2) 15 RISK-79 kolom pemantauan "ada tidaknya objeknya": dibatasi ke objek yang terlihat keduanya dan jawaban pemakai sendiri (§6.2) |
| `architect` (15) | Risiko: bias validasi pilot yang dijalankan satu pasangan (§6.5); build TestFlight kedaluwarsa; platform lebih dulu; kesediaan membayar dan komisi toko belum diketahui; Lebaran di jendela pilot; definisi "pemakaian bertahan" |
| `architect` (06/09) | Aturan langganan §7.5 (a)–(d): apakah memerlukan baris di 06 (Plus tidak memberi peran atau wewenang; tidak lintas rumah tangga) |
| `security-privacy` (10) | (1) Periksa bahwa VL-00 sampai VL-10 hanya memakai data yang sudah terlihat kedua anggota dan tidak menambah pengumpulan (SEC-44, SEC-49, PRN-11); (2) Q-215; (3) kelas data penagihan (bukti pembelian toko, identitas pembayar) sebelum FS-3; (4) apakah SEC-107 memerlukan F-FIN-023; (5) model keamanan browser bila Q-210 dibuka. **Dijawab di 10 §7.4 (fase D)** dan diterapkan di §3.6, §6.1, §6.2, 05 FR-FAM-013/AC-12, 17 Q-215, Q-217. **Sebelum S6 (keputusan 46):** periksa VL-02 bagian (a) (pertanyaan tentang aplikasi, saluran jawaban istri, larangan desakan) dan bagian (b) (catatan bug hanya tanggal, golongan, status; tanpa isi rumah tangga, tangkapan layar data nyata, atau nama pelapor) terhadap batas §6.1 dan SEC-44, SEC-49. **Dijawab di 10 §7.4 baris "VL-02 sesudah keputusan 46"** dan diterapkan di §6.1, §6.2 |
| `ux-designer` (08) | Contoh layar statis arah C untuk uji prototipe §6.3 (yang dibuat untuk 08 §19.8 dipakai ulang, tanpa data nyata); kelak kata layar status Plus tanpa nama pembayar |
| `finance-domain` (07) | Tidak ada; dokumen ini tidak membuat aturan uang |
| `researcher` (18) | R-15: komisi dan aturan langganan Apple dan Google, kewajiban pembelian dalam aplikasi, tingkat harga Indonesia, pajak langganan aplikasi di Indonesia, pengakuan hak lintas platform. R-16: bukti primer kesediaan membayar aplikasi keuangan rumah tangga di Indonesia. R-17: ketersediaan klien browser untuk Connect; status Windows Subsystem for Android. R-18: apakah deskriptor rating usia 12 Januari 2027 berlaku bagi aplikasi yang hanya di TestFlight internal; tanggal resmi Idulfitri dan awal Ramadan 2027. R-09 (Bank Jago Kantong Bersama) tetap relevan bagi positioning |
| Semua (keputusan 51, 52) | `architect`: 14 dan 15 mengikuti jalur rumah tangga lawan jalur publik opsional (workstream dan risiko FS-3, GR-5, beta, akun Play, penagihan toko tidak lagi terjadwal; GR-4 dan QG-31 bagian keamanannya menjadi gerbang bersyarat dengan dua pintu masuk, login orang tua atau jalur publik, keputusan 52; batas Android "sebelum GR-4" hanya bila jalur itu dibuka; irama rilis TestFlight berlanjut tanpa batas waktu). `security-privacy`: kolom "Wajib sebelum tahap publik" dan butir sebelum GR-4 di 10 dibaca sebagai syarat jalur opsional; W-32 tanpa batas waktu. Lead: 16, 01, glosarium (*Tinjauan tiga bulan*, *Titik akhir roadmap*, *Jalur publik opsional* bila disetujui) |
| Lead | README: awalan `FS-N`, `GR-N`, `VL-NN` (sudah); fase D: `T2-N` (gelombang Tier 2) dan `TB-N` (tantangan 13 terhadap brief). Glosarium, istilah baru bila disetujui: *langganan rumah tangga* (Plus), *beta tertutup*, *pemakaian harian yang bertahan* (kini menurut keputusan 46; sudah dipasang lead di glosarium). Menjadwalkan FR Tier 2 untuk IDEA-15, IDEA-19, IDEA-08 sebelum gelombang T2-3. Bahan 16 dan 01: §1, §7.5 |

### 12.3 Yang tidak diputuskan dokumen ini

Penyedia hosting dan email (S4-2; waktunya sudah diputuskan, sebelum data nyata, keputusan 45), kapan tepatnya Android dibangun di antara GR-2 dan GR-4 (rentangnya diterima atas saran lead lewat keputusan 50; Q-37 sendiri sudah diputuskan, keputusan 42), harga dan isi Plus (Q-36, Q-216), pengganti pengelola server (Q-222, pemilik produk bersama istrinya), sewa atau beli Mac bila dana kelak ada (keputusan 43 menundanya), apakah jalur publik opsional dimasuki sama sekali (keputusan 51), dan keputusan lanjut di GR-3 dan GR-5 adalah keputusan pemilik produk. Yang sudah diputuskan sebelum S6 dan hanya diterapkan di sini: Q-213 (46), Q-215 (47), Q-219 (44), Q-220 (45), Q-221 (48).

### 12.4 Riwayat

| Tanggal | Perubahan |
|---|---|
| 12 November 2026 | Draf pertama (S5 fase A) |
| 12 November 2026 | S5 fase D atas `critic` S5 A (A-03, A-05, A-06, A-07) dan B (B-02 sampai B-06, B-08, B-09, B-12 sampai B-14), 14 §12.2, 15 §11: irama rilis (ADR-022 butir 1 mengikat, ADR-028 usulan, Q-219); S4-2 **[Usulan]** dan Q-220; SEC-109 (b) sebelum TestFlight pertama; RLS keluar dari daftar FS-3; syarat penguji Play dan ukuran beta, Q-221; Q-213 sebelum GR-2; ukuran "fitur Tier 2 tidak dipakai" tanpa data yang tidak terlihat keduanya; VL-10 memicu tahan tulis; saluran jawaban VL-08; klien web tanpa batas waktu; `T-N` menjadi `TB-N`; IDEA-15, IDEA-19 masuk T2-3; biaya pilot dengan PPN dan email. Susulan atas 10 §4.2, §7.4: batas sumber ukuran validasi (§6.1; VL-02 hanya entri yang terbaca keduanya); GR-4 butir 7 (butir 10 §4.2); daftar keamanan FS-3 menjadi sisa sebelum GR-5 (dengan SEC-157); F-FIN-023 dari FS-3 ke "Menunggu bukti" (SEC-107 dipenuhi FR-FAM-013/AC-12); Q-215 dengan syarat `security-privacy`. Susulan atas 18 (§1.5, §2.4, §2.13, §2.14): penguji Play dihitung per orang dewasa; tabel toko §7.2 dan aturan (e) tanpa Family Sharing; tanggal Idulfitri 1448 H. Atas `critic` B2: FR-FAM-013/AC-12 di FS-1, jadi SEC-107 keluar dari GR-4 butir 7; Q-215 di §12.3; saran tim untuk Q-37 (§3.2, B Android dulu) |
| 12 November 2026 | Sesudah S5, keputusan 42 (17): §3.2 pilihan platform diganti keputusan (iOS dan iPadOS lebih dulu; saran B dan usulan "tidak dengan satu platform" digantikan, riwayat singkat; risiko sisa pencatatan dari iPad, 15 RISK-18); §1 baris FS-1, FS-2 dan diagram (klien Android sesudah GR-2); GR-1 butir 3; GR-2 butir 1 (sepuluh butir 10 §4.1, termasuk SEC-59 (e)), butir 4 (tata letak iPad mini 7 menurut 08 §16.4) dan 5, tanpa Android; §3.4 klien Android di dalam pilot dengan syarat ulang GR-2 sebelum data nyata di Galaxy A55; GR-4 butir 8; §5 diagram; §6.4 dan Q-221 tetap sebelum GR-4; §7.4, §8; §12. Atas `critic`: CR42-08 Touch ID iPad di GR-2 butir 1, kini **[Terverifikasi]** 18 §2.12.9; CR42-09 GR-2 butir 5 tanpa label usulan, mengikuti 14 QG-30 butir 4, bunyi ES §3 di 17 K-67 |
| 12 November 2026 | Sebelum S6, keputusan 43 sampai 50 (17): §3.5 dan VL-02 menurut keputusan 46 dan bacaan lead (GR-3 paling cepat tiga bulan sesudah GR-2; (a) kedua pasangan nyaman dan aman memakai aplikasi, ditanya tentang aplikasi, bukan pasangan, saluran istri seperti VL-08; (b) 14 hari tanpa bug baru yang berat; ambang 5 dari 7 hari dicabut, frekuensi hanya bahan); §6.2 baris GR-3 mundur; §1, diagram §1 dan §5; TB-4; §3.2 Mac (43) dan S4-2 sebelum data nyata (45); §3.4, §8, §11 butir 8 irama 60/80 hari mengikat (44); §6.3 empat syarat menjadi aturan (47); GR-4 butir 5 dan §6.4 akun Play pribadi dengan langkah tanya dukungan Play Console (48); GR-2 butir 1 (sebelas butir 10 §4.1, butir 11 tinjauan `security-privacy`) dan GR-4 butir 9 (49); GR-2 butir 4, 5 dan §3.2 bunyi brief dan ES (50); GR-2 butir 8 Q-222; S4-2 dengan syarat instans sementara 10 §4.1 butir 7; §6.3 daftar periksa bertanggal sebelum wawancara pertama (10 §7.4, §7.7); VL-02 (a) hasil "keduanya sudah"/"belum" tanpa nama, (b) catatan bug di luar aplikasi dan server, reproduksi di staging, bocor lewat SEC-148 (10 §7.4 baris "VL-02 sesudah keputusan 46"); §12. Atas `critic` (revisi keputusan 43 sampai 50): P-1 waktu Android diterima atas saran lead lewat keputusan 50 (§3.2, §3.4, GR-4 butir 8, §12.3); P-2 dan P-10 TB-4 dan §10 dengan kedudukan per bagian dan 17 K-68; P-3 VL-02 (a) terus terang tentang ketertelusuran "belum"; P-4 bug berat lama yang belum diperbaiki menahan GR-3 (bacaan lead **[Usulan]**); P-5 dan M-6 syarat instans sementara (a) sampai (d) berlabel **[Usulan]** `security-privacy`; M-7 "sudah"/"belum"; M-12 alasan tiga bulan milik pertanyaan lead |
| 12 November 2026 | Keputusan 51 (17): §0 jalur rumah tangga dan jalur publik opsional, dengan kata-kata pemilik produk; §1 tabel dan diagram (FS-2 tanpa gerbang keluar, titik akhir roadmap; GR-3 sampai FS-4 opsional; Tier 3 dari pemicu rumah tangga atau jalur opsional), butir 1, 3, 5; §2.2 baris Could; §3.1 wawancara dan uji pesaing opsional; §3.2 dan §3.4 Android tanpa batas waktu di jalur rumah tangga (brief §1 "if a beta is ever held"), APK tanpa Play; §3.3 W-32 tanpa batas waktu (diajukan ke lead); §3.4 tinjauan tiga bulan dan titik akhir roadmap, anak 12+; §3.5 sampai §3.7 bertanda opsional, keputusan 46 syarat **bila** go public, 47 dan 48 bagian jalur opsional; §3.8 pemicu dari rumah tangga sendiri; §4.1; §5 diagram; §6.1 VL-02 pindah ke jalur opsional, percakapan berkala berhenti sesudah tinjauan tiga bulan; §6.2 baris titik akhir; §6.3, §6.4, §6.5; §7; §8; §9; §10 TB-2 sampai TB-5, TB-6 baru, tanpa baris K baru; §11 butir 9; §12. Susulan `security-privacy` (10 §7.8): GR-4 bersyarat, Q-224 sampai Q-226, R-19 di "Rilis dalam pilot". **Keputusan 52**: §0, §1 tabel (pita "Gerbang bersyarat") dan butir 5, diagram §1 dan §5, §3.5, §3.6 (GR-4 dua pintu masuk: login orang tua atau jalur publik), §3.8 FS-5, §6.2 VL-08, §8. Atas sisa `critic`: tanda baca kutipan keputusan 46 di §3.5. Hitungan tidak berubah: 131 fitur (55, 18, 58); 198 Must = 190 FS-1 + 8 FS-2 |
| 12 November 2026 | Keputusan 53 sampai 55 (17) dan SEC-158: §3.2 baris ketergantungan pendaftaran tertutup (SEC-158; FR-FAM-001/AC-9, AC-10); §3.4 tinjauan `security-privacy` sebelum APK memegang data nyata dan setahun sekali (53), F-FAM-022 di T2-1 (54), saluran iOS (55) dengan catatan pendaftaran tertutup; §3.6 kalimat Q-224 dan Q-225, GR-4 butir 7 (SEC-25 dan SEC-18 dari FS-2), daftar keamanan FS-3 tanpa SEC-18; §4.1; §12. Hitungan: 132 fitur (55, 19, 58); 198 Must = 190 FS-1 + 8 FS-2 (tidak berubah) |
