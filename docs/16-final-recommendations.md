# 16 — Rekomendasi Akhir

| | |
|---|---|
| Pemilik | Lead |
| Tahap | S5 — Rencana dan sintesis, fase C (12 November 2026); S6 (12 November 2026): pass konsistensi |
| Status | Ditinjau `critic` (S5 B, lalu tinjauan ulang B2, 12 November 2026); direvisi untuk keputusan 42 dan ditinjau `critic` lagi (12 November 2026); temuan diperbaiki; direvisi untuk keputusan 43–50 dan 51–55 (12 November 2026); S6: keputusan 56–68 dimasukkan ke §1, §3, §4, §5, §10; sesudah S6 (13 November 2026): audit kesiapan bangun dan keputusan 69 (klien web lebih dulu, bawaan 69.0–69.10) dan 70 (nama Casa) dimasukkan ke §1 sampai §5, §8 sampai §10 |
| Sumber yang mengikat | Brief §0.3, §1, §1.1, §16 (baris 16: "MVP, architecture, critical decisions, deferred decisions, features to reconsider, top opportunities"); `CLAUDE.md`; keputusan pemilik produk di [17](17-open-questions-assumptions.md) dan [_progress.md](_progress.md) |
| Merangkum | 02 sampai 15 dan 18. Dokumen ini tidak menetapkan aturan, persyaratan, atau keputusan baru; setiap butir menunjuk sumbernya |
| Dipakai oleh | [01](01-executive-summary.md) |

---

## 0. Cara membaca

- **Kedudukan setiap butir** ditulis terang (definisi tunggal: *Tingkat kepastian* di [glosarium](00-glossary.md)): *diputuskan pemilik produk* (keputusannya sendiri), *diterima* (atas saran lead atau lewat "lanjut" di checkpoint), atau **[Usulan]** (rekomendasi tim yang belum diputuskan). Kata pemilik produk hanya dikutip dari [17](17-open-questions-assumptions.md) atau [_progress.md](_progress.md).
- **Tier mengatur urutan membangun, bukan batas rancangan** (brief §1.1). Semua model, aturan, alur, istilah, dan layar dirancang untuk setiap rumah tangga tahap publik; rumah tangga pilot adalah pengguna pertamanya.
- Label bukti mengikuti [README](README.md#label-bukti). Istilah mengikuti [00-glossary.md](00-glossary.md).

---

## 1. Rekomendasi dalam satu halaman

1. **Bangun MVP yang menjawab satu pertanyaan lebih dulu:** "sisa uang berapa, dan siapa yang memegang" bisa dijawab kedua pasangan dari aplikasi tanpa saling bertanya (brief §1.1 kriteria sukses 1; PRN-01). Seluruh Tier 1 melayani pertanyaan itu, permintaan dan persetujuan yang mendahului pengeluaran, dan chat rumah tangga tempat keduanya dibahas.
2. **Pertahankan arsitektur yang sudah diputuskan** (kedudukan per lapisan di §3): **klien web (PWA) lebih dulu**, responsif dan dapat dipakai luring, di semua perangkat pilot (keputusan 69; ADR-030); aplikasi native per platform (SwiftUI; Jetpack Compose Material 3) menyusul sesudah GR-2, stack diputuskan pemilik produk; satu server Go (keputusan 29) dengan PostgreSQL di server di Indonesia (keputusan 10; bentuk basis data diterima lewat "lanjut", D-99); semua aturan uang dan persetujuan hanya di server (keputusan 30; ADR-013); dan pencatatan yang tidak pernah dihalangi jaringan (PRN-06).
3. **Masuk ke data nyata hanya lewat gerbang GR-2** (13 §3.3; 14 QG-30): semua kriteria penerimaan Tier 1 lulus sebagai uji bernama, tiga belas butir 10 §4.1 selesai (butir 4 dan 5 bukan butir pilot, keputusan 68; butir 3, catatan persetujuan uji TestFlight, hanya untuk aplikasi iOS native), dan klien web terpasang sebagai web app layar utama di iPhone 13 istri dan iPad mini 7 pemilik produk, dan di Chrome pada Galaxy A55 pemilik produk; aplikasi native tidak ditunggu (keputusan 69, bawaan 69.0).
4. **Ukur pemakaian dengan ukuran yang tidak mengawasi siapa pun** (13 §6). Roadmap berakhir saat rumah tangga memakai aplikasi sehari-hari; tinjauan tiga bulan sesudah GR-2 membaca ukurannya dan bukan gerbang (13 §3.4; keputusan 51). Go public tidak direncanakan. Bila kelak dipertimbangkan, GR-3 mensyaratkan kedua pasangan merasa nyaman dan aman, dan 14 hari terakhir tanpa bug baru yang berat (keputusan 46).
5. **Sebelum orang di luar suami dan istri memakai aplikasi, termasuk orang tua** (keputusan 52): gerbang pengguna luar pertama (13 GR-4; 14 QG-31 (a)), dengan pemberitahuan pengelola server versi keluarga dan butir keamanannya. Menurut jawaban advokat, pengecualian rumah tangga UU PDP tetap berlaku untuk keluarga yang diberi login, termasuk orang tua, dan pendaftaran PSE tidak perlu selama aplikasi tertutup untuk keluarga dan nonkomersial (keputusan 59); kerabat dewasa lain dihitung sama dengan orang tua (keputusan 63). **Bila kelak go public** (jalur publik opsional, keputusan 51): freemium per rumah tangga dengan satu langganan rumah tangga (Plus) **[Usulan]** (13 §7.5) dan beta tertutup lebih dulu (GR-5).
6. **Keputusan berikutnya ada pada pemilik produk** (§5). Yang terdekat: siapa mengoperasikan server bila pemilik produk berhalangan lama (Q-222, bersama istrinya, sebelum GR-2), dan hosting dan email: menurut keputusan 45 sebelum deploy pertama yang memegang data nyata, dalam praktik sebelum FS-1 langkah 7, saat instans utama dipasang di VPS pilihan (14 §3.3).

---

## 2. MVP yang direkomendasikan

MVP = Tier 1 + Tier 2 (glosarium *MVP*). Rinciannya di [04 §9](04-feature-map.md#9-potongan-mvp-untuk-checkpoint-s3); persyaratannya di [05](05-functional-requirements.md).

| | Isi | Hitungan | Kedudukan |
|---|---|---|---|
| **Tier 1 — Inti** (FS-1, dibangun sebelum data nyata) | `FAM` satu rumah tangga dua dewasa setara, termasuk kunci aplikasi dan halaman Bantuan; `FIN` custody, uang masuk dan pembagiannya, pengeluaran cepat, serah terima, jatah pribadi, anggaran bulanan per kategori, tagihan berulang dan pos rutin, cicilan, utang dan piutang termasuk keluarga besar, tabungan sederhana, aset dan kekayaan bersih, ringkasan; `REQ` permintaan dua arah sampai penyelesaian sisa; `NTF` dengan jam tenang dan ringkasan; `COM` chat rumah tangga ringan dengan kartu objek, tanda terkirim dan dibaca, kehadiran, sembunyikan untuk saya dan laporkan; `CAL` pengingat tagihan | 55 fitur | Potongan S3 diterima lewat "lanjut". Tiga fitur Tier 1 tambahan di checkpoint S4: F-COM-014 **diputuskan pemilik produk** (keputusan 37, S4-7); F-FAM-019, F-FAM-020 rumusan lead yang dikukuhkan lewat "lanjut" (S4-6 (iii); D-114) |
| **Tier 2 — Versi ringan** (FS-2, bergelombang di dalam pilot) | Sebutan anggota; permintaan tindakan; daftar belanja; masalah rumah menjadi tugas; status suasana hati dan tindak lanjut sakit; rutinitas; rencana mendatang; pendamping `AI` versi ringan (kumpulan kalimat R1); faktor kedua dan kode pemulihan opsional per akun (F-FAM-022, keputusan 54) | 19 fitur | Diterima (checkpoint S3 dan S4; sebutan anggota permintaan pemilik produk sendiri, keputusan 40) |
| **Tier 3 — Ditunda** | Selebihnya, termasuk struktur keluarga lain, `INV`, `MEAL`, `TML`, integrasi | 58 fitur | Garis besar satu halaman per modul. Klien web bukan fitur 04: ia klien pertama yang membawa Tier 1 dan Tier 2 (keputusan 69; ADR-030), bukan fitur tersendiri |

Seluruhnya: 132 fitur; 219 FR (Must 198, Should 21), 207 BR (05); 49 skenario (12). Setiap Must terpetakan ke fitur, skenario, dan fase roadmap: 190 di FS-1 dan 8 Must bersyarat di FS-2 (05 §6; 13 §4).

**Tiga konsep dokumen 11 di Tier 1 sebagai sub-fitur** (diterima lewat "lanjut", S3-2): hitung sisa uang tunai sepulang belanja (IDEA-01), pos rutin tanpa tanggal jatuh tempo seperti token listrik dan galon (IDEA-03), dan "bagikan seperti terakhir kali" untuk uang masuk (IDEA-02). Ketiganya memangkas usaha mencatat yang berulang setiap bulan, sesuai PRN-02 dan PRN-05.

**Yang sengaja tidak dibangun di pilot** (04 §8; 17): pelacakan siklus (D-63); kalkulator zakat (11 IDEA-09); voice note, reaksi, grup, dan pencarian chat (D-70); sinkron bank (brief §1; 18 §3.1); perintah mengeluarkan orang dewasa lain (S3-4 A). Aplikasi mencatat dan menata uang; ia tidak memegang, memindahkan, atau memproses uang sungguhan (brief §1).

---

## 3. Arsitektur

Rinciannya di [09](09-technical-architecture.md) dan ADR-012 sampai ADR-030; roadmap teknisnya di [14](14-engineering-roadmap.md).

| Lapisan | Rekomendasi | Kedudukan |
|---|---|---|
| Klien pertama | Klien web (PWA): SvelteKit mode SPA dengan Svelte 5, service worker, Dexie di IndexedDB, kontrak Connect yang sama; satu origin dengan server; responsif untuk ponsel, tablet, dan browser desktop; di iPhone dan iPad dipasang ke layar utama sebelum masuk. Tata letak, arah visual, dan kata-katanya sama dengan aplikasi native (08 §25), tetapi tidak terasa sepenuhnya iPhone atau Android | Web lebih dulu: **diputuskan pemilik produk** (keputusan 69). Penerimaan bahwa web belum terasa native: **diputuskan pemilik produk** (keputusan 69 (b)). Stack dan bawaan 69.0–69.10: diterima atas saran lead (ADR-030) |
| Aplikasi native | UI native penuh: Swift (SwiftUI) untuk iPhone dan iPad, Kotlin (Jetpack Compose Material 3) untuk Android; iOS dan iPadOS 26, Android 10 (API 29); dibangun sesudah GR-2, iOS dan iPadOS lebih dulu, lalu Android (bawaan 69.5) | **Diputuskan pemilik produk**: stack (`_progress.md` "Keputusan pemilik produk: stack mobile"); OS minimum (17 keputusan 32, S4-3). Urutan sesudah GR-2: diterima atas saran lead (bawaan 69.5) |
| Letak aturan | Hanya di server; perangkat tidak menghitung aturan uang atau persetujuan; tiga bantuan isian dengan vektor emas sebagai pengecualian ES §2 | **Diputuskan pemilik produk** (keputusan 30, S4-1 (ii); ADR-013) |
| Server | Satu monolit modular Go, satu PostgreSQL 18 dengan RLS dan transaksi `SERIALIZABLE` per perintah tulis, kontrak Protobuf + Connect | Go: **diputuskan pemilik produk** (keputusan 29, S4-1 (i); ADR-015). Bentuk server dan basis data: diterima sebagai default lewat "lanjut" (D-99, D-100; ADR-014, ADR-016, ADR-017) |
| Sinkron | Antrean di perangkat, SSE + tarik berkursor saat aplikasi terbuka, push umum saat tertutup: Web Push untuk klien web, APNs dan FCM untuk aplikasi native | Diterima sebagai default lewat "lanjut" (D-100, D-101; ADR-018, ADR-019); Web Push: diterima atas saran lead (keputusan 69; ADR-030, catatan web ADR-019) |
| Hosting, cadangan, email | Satu VPS di Indonesia, cadangan terenkripsi ke penyedia Indonesia lain, email transaksional minimal | **Ditunda sampai sebelum deploy pertama yang memegang data nyata**, keputusan pemilik produk (keputusan 31, S4-2; arti "deploy" keputusan 45); dalam praktik sebelum FS-1 langkah 7 (14 §3.3); sampai itu server uji sementara berisi data uji. Syarat data di Indonesia (keputusan 10) tetap mengikat |
| `AI` | Jalur R1: kumpulan kalimat dibangun tanpa data rumah tangga; Claude; batas USD 10 per bulan dan USD 3 per pembangunan | Claude: **diputuskan pemilik produk** (keputusan 33, S4-4 (ii)). R1 dan batas biaya: rumusan lead yang dikukuhkan lewat "lanjut" (S4-4 (i)); angka batas **[Asumsi]** (ADR-021) |
| Build iOS (jalur iOS native, sesudah GR-2) | Tidak diperlukan untuk klien web: web dibangun dan dipasang tanpa Mac dan tanpa akun Apple Developer (ADR-030 butir 19). Saat jalur iOS native dimulai: Xcode Cloud + TestFlight internal, dengan Mac sekali dari AWS EC2 Mac atau Mac pinjaman (bawaan 69.6); tanpa sewa bulanan dan tanpa membeli Mac sampai ada dana. Pemasangan keluarga tetap lewat TestFlight internal; pindah ke App Store tersembunyi (unlisted) bila Apple menghentikannya (perjanjian Apple membatasi TestFlight untuk pengujian; RISK-92) | Jalur A: **diputuskan pemilik produk** (keputusan 32, S4-3 (i); ADR-022). Tanpa sewa bulanan dan tanpa membeli: diterima atas saran lead (keputusan 43). Jalur pasang: diterima atas saran lead (keputusan 55) |
| Lingkungan dan rilis pilot | Staging di VPS yang sama sejak GR-2, tanpa sistem feature flag, saklar *tahan tulis* yang tidak pernah menahan perintah pelindung; batas rilis iOS 80 hari sejak unggah dengan rilis pemeliharaan di 60 hari (berlaku sejak aplikasi iOS native dipasang lewat TestFlight); klien web dirilis lewat deploy server dengan service worker pengganti untuk rollback (ADR-030; 14 §6) | Batas rilis 60/80: diterima atas saran lead (keputusan 44, menjawab Q-219; ADR-028 butir 4). Butir ADR-028 lainnya: diterima atas saran lead (keputusan 56); ADR-028 seluruhnya Diterima |

**Biaya pilot per bulan** menurut 09.4 §7.1: **Rp260.667** + PPN bila berlaku, ditambah **USD 1,32** selama jalur web (paling tinggi USD 10 bila batas `AI` tercapai), belum termasuk email; Rp244.000 dari angka rupiah itu bergantung pada S4-2. Selisihnya dari angka sebelumnya adalah keanggotaan Apple Developer Program, USD 99 ÷ 12 = USD 8,25 per bulan, yang baru dibayar saat jalur iOS native dimulai; sesudah itu USD 9,57 (paling tinggi USD 18,25). Satu VPS sederhana, sesuai brief §1.

**Tiga tradeoff yang paling terasa saat membangun** (09 §1):

1. **Server saja** menghemat satu bahasa dan satu langkah build yang hanya bisa di macOS, tetapi perangkat tidak menampilkan sisa atau tanda saat offline sampai server menjawab.
2. **Web lebih dulu** membuat aplikasi dipakai di semua perangkat pilot tanpa Mac dan dengan satu UI, tetapi tampilannya belum terasa sepenuhnya iPhone atau Android (15 RISK-97, Diterima), catatan luring bergantung pada penyimpanan browser (RISK-94), dan permukaan keamanan browser bertambah (RISK-96; 10.5). Pemilik produk menerima harga ini (keputusan 69). Siklus UI iOS tanpa Mac (keputusan 43; RISK-01) baru berlaku saat jalur iOS native dimulai.
3. **R1** menjaga data rumah tangga tetap di Indonesia dan biaya `AI` tetap, dengan harga bahwa saran tidak menyebut hal yang spesifik.

---

## 4. Keputusan kritis yang sudah diambil

Dicatat agar tidak ditanyakan ulang. Rincian dan kata-kata pemilik produk di 17.

| # | Keputusan | Kedudukan | Sumber |
|---|---|---|---|
| 1 | Dua dewasa setara, dengan pelaku terlihat pada setiap tindakan; tidak ada perintah mengeluarkan orang dewasa lain | Kesetaraan: **diputuskan pemilik produk** (keputusan 4, S1). Tanpa perintah mengeluarkan: diterima sebagai default lewat "lanjut" (S3-4 A) | PRN-09, PRN-10; 06 PG-01 sampai PG-15 |
| 2 | Setiap hal punya tingkat visibilitas; berbagi adalah pilihan pemiliknya; dana bersama tidak dapat disembunyikan | Tingkat visibilitas: **diputuskan pemilik produk** (keputusan 4). Rumusan prinsipnya diterima di checkpoint S2 | PRN-12, PRN-13 |
| 3 | Catatan lama tidak pernah disunting; selisih dan entri terlambat dicatat sebagai entri baru; jumlah uang bilangan bulat | Tidak menyunting catatan lama: **diputuskan pemilik produk** (keputusan 7, S1). Bacaan append-only: **diputuskan pemilik produk** sebagai keputusan sementara sampai advokat menjawab (keputusan 24, S2-8); advokat sudah dikonsultasi (keputusan 59) dan konsultasi lanjutan tidak dilakukan sampai go public (keputusan 67), jadi bacaan ini berlaku untuk pemakaian keluarga | ADR-002, ADR-004, ADR-007, ADR-008 |
| 4 | Pemakaian besar yang tidak diajukan lebih dulu tetap tersimpan dan ditandai netral, tidak dihalangi | Diterima atas saran lead (S2-2) | ADR-009; PRN-06 |
| 5 | Chat rumah tangga; dari chat dapat dikirim tugas, pengeluaran, dan rencana; tanda terkirim dan dibaca wajib; kehadiran seperti WhatsApp | Chat dan tanda: **diputuskan pemilik produk** (keputusan 26). Bentuk dan cakupannya: diterima atas saran lead. Kehadiran: **diputuskan pemilik produk**, berlawanan dengan saran lead (keputusan 27, dipertahankan di keputusan 35) | ADR-011 |
| 6 | Stack mobile Swift + Kotlin, native penuh | **Diputuskan pemilik produk** | ADR-012 |
| 7 | Backend Go; aturan hanya di server | **Diputuskan pemilik produk** (keputusan 29, 30) | ADR-013, ADR-015 |
| 8 | Tanpa E2EE di pilot; status kepedulian yang tidak dibagikan tidak meninggalkan perangkat; E2EE dinilai lewat ADR sebelum tahap publik, bila diambil | **Diputuskan pemilik produk** (keputusan 36, S4-6 (i)) | 10-D; ADR-024 |
| 9 | Pengelola server pilot adalah pasangan, dengan syarat istrinya diberi tahu terus terang sebelum mulai memakai | **Diputuskan pemilik produk** (keputusan 25, S2-10). Bentuk pemberitahuannya diterima sebagai default lewat "lanjut" (S4-6 (ii) A) | FR-FAM-019/AC-7; 15 RISK-11 |
| 10 | Advokat ditanya sebelum tahap publik; risiko PSE selama pilot diterima | **Diputuskan pemilik produk**, berlawanan dengan saran B tim (keputusan 34, S4-4 (iii); W-32). Konsultasi sudah dilakukan pemilik produk (keputusan 59; Q-224 selesai): tanpa pendaftaran PSE sampai dibuka untuk publik. Konsultasi lanjutan tidak dilakukan; urusan hukum dibuka lagi hanya bila go public diputuskan (keputusan 67) | 15 RISK-65; 10 §3.4 |
| 11 | Sembunyikan untuk saya, laporkan, kontak, keluar; tanpa pemindaian dan tanpa mengeluarkan orang dewasa | **Diputuskan pemilik produk** (keputusan 37, S4-7) | F-COM-014; 10.4 SEC-153, SEC-154 |
| 12 | Arah visual C, terang dan tenang | Diterima atas saran lead (S4-10 (i)) | 08 §19 |
| 13 | Aplikasi menyapa dengan "kamu" | Rumusan lead yang dikukuhkan lewat "lanjut" (S4-10 (iv)) | 08 |
| 14 | ~~Web atau PC tidak dibangun di pilot~~ **Digantikan keputusan 69**: klien web dibangun lebih dulu | Semula diterima atas saran lead (keputusan 41) | 17 §2.20; Q-210 |
| 15 | Ambang dibuat, diubah, dan diakhiri hanya dengan persetujuan semua pihak, di setiap rumah tangga, pilot dan publik | **Diputuskan pemilik produk**, berlawanan dengan saran B dan bawaan C tim (keputusan 39, S4-9 (i)). Risiko sisa yang diterimanya: satu pihak dapat mempertahankan ambang atas pihak lain tanpa batas waktu dengan tidak menyetujui perubahan atau pengakhirannya, berlaku sama bagi setiap pihak (W-34, Pagar) | 17 keputusan 39, W-34 |
| 16 | Pilot di iOS dan iPadOS: istri di iPhone 13, pemilik produk di iPad mini 7; iOS dan iPadOS dibangun lebih dulu dan data nyata tidak menunggu Android. Android (Kotlin) untuk Galaxy A55 pemilik produk menyusul sesudah GR-2, selesai sebelum jalur pengujian Google Play dibuka bila jalur publik kelak diambil (GR-4 butir 8; keputusan 51; bawaan lead di keputusan 42; bunyinya diterima atas saran lead lewat keputusan 50, brief §1 "Pilot devices") | **Diputuskan pemilik produk** di luar checkpoint, sesudah S5 (keputusan 42; menjawab Q-37). Bunyi brief §1 "Pilot devices" dan "Platforms" serta ES §3 disesuaikan (keputusan 50; K-66, K-67). Risiko sisa: pengeluaran pemilik produk di luar rumah dicatat belakangan dari iPad sampai Android terpasang. **Urutan platformnya digantikan keputusan 69**: klien web lebih dulu di ketiga perangkat, termasuk Galaxy A55 lewat Chrome, sehingga risiko sisa itu tidak berlaku lagi; fakta perangkat dan dasar "iOS sebelum Android" untuk aplikasi native tetap (bawaan 69.5) | 17 keputusan 42, 50, 69; 13 §3.2; 15 RISK-18 |
| 17 | Tetap Xcode Cloud + TestFlight; tanpa sewa Mac bulanan dan tanpa membeli Mac sampai ada dana | Diterima atas saran lead (keputusan 43; ADR-022 butir 4). Sejak keputusan 69 berlaku saat jalur iOS native dimulai, bukan di jalur menuju GR-2. Risiko sisa: siklus UI iOS lebih lambat, tanpa debugger di perangkat (15 RISK-01, kini di luar jalur kritis) | 17 keputusan 43, 69 |
| 18 | Batas rilis iOS: rilis pemeliharaan bila 60 hari tanpa rilis, batas keras 80 hari sejak unggah; berlaku bagi aplikasi iOS native | Diterima atas saran lead (keputusan 44, Q-219) | ADR-028 butir 4 |
| 19 | Pemakaian harian yang bertahan: paling cepat tiga bulan sesudah GR-2; kedua pasangan merasa nyaman dan aman; 14 hari terakhir tanpa bug baru yang berat | Batas tiga bulan dan syarat nyaman, aman, tanpa bug baru: **diputuskan pemilik produk** (keputusan 46, Q-213). Rumusan gabungan adalah bacaan lead; arti "bug berat" (14 hari) diterima atas saran lead; bug berat lama yang belum diperbaiki juga menahan GR-3 adalah bacaan lead **[Usulan]**. Selisih dengan brief §1 "sustained daily use": 17 K-68 | 13 §3.5, §6; glosarium |
| 20 | Wawancara dan prototipe statis dengan keluarga lain bukan "pengguna luar", dengan empat syarat | Diterima atas saran lead (keputusan 47, Q-215, Pagar) | 10 §7.4; 13 §6.3 |
| 21 | Bila jalur publik diambil (keputusan 51): akun Google Play pribadi; sebelum dibuat, dukungan Play Console ditanya apakah aplikasi wajib akun organisasi; bila wajib, dibawa kembali ke pemilik produk | Diterima atas saran lead (keputusan 48, Q-221) | 13 §6.4; 14 §6.7 |
| 22 | Tinjauan `security-privacy` hanya-temuan sebelum GR-2, sebelum aplikasi iOS native dan aplikasi Android memegang data nyata, sebelum pengguna luar pertama, dan setahun sekali | Diterima atas saran lead (keputusan 49, 53; sebelum iOS native: bawaan 69.4); bunyi baris `CLAUDE.md` rumusan lead (pembaruannya menunggu pratinjau, K-77) | `CLAUDE.md` "Build team" |
| 23 | Aplikasi untuk keluarga pemilik produk sendiri, berstandar produk publik; go public tidak direncanakan tetapi tetap terbuka tanpa pembangunan ulang. Roadmap berakhir di pemakaian sehari-hari; GR-3, GR-5, FS-3, FS-4 dan isinya menjadi jalur publik opsional (13 §1) | **Diputuskan pemilik produk** (keputusan 51; bacaan lead dikukuhkannya). Bunyi brief §1.1 dan `CLAUDE.md` rumusan lead | Brief §1.1; 13 §1, §3.4 |
| 24 | Orang tua yang diberi login dihitung pengguna luar: gerbang pengguna luar pertama berlaku walaupun tidak go public | **Diputuskan pemilik produk**, berlawanan dengan usulan `security-privacy` (keputusan 52) | 13 GR-4; 14 QG-31 (a); 10 §4.2; 15 RISK-93 |
| 25 | Faktor kedua (TOTP; passkey tetap tahap publik) dan kode pemulihan sebagai satu pasang fitur opsional per akun, dibangun di FS-2, tidak menahan GR-2 | Diterima atas saran lead (keputusan 54, Q-227; Pagar) | 10 SEC-18, SEC-25; 15 RISK-40 |
| 26 | Pemasangan iOS keluarga tetap lewat TestFlight internal; pindah ke App Store tersembunyi hanya bila Apple menghentikannya; berlaku bagi aplikasi iOS native | Diterima atas saran lead (keputusan 55, Q-228) | ADR-022; 15 RISK-92; 18 §5.2.9 |
| 27 | ADR-028 seluruhnya: staging di VPS yang sama dengan basis data terpisah, tanpa feature flag, saklar *tahan tulis*, server melayani build lama yang masih terpasang | Diterima atas saran lead (keputusan 56) | ADR-028; 14 §6.1 |
| 28 | Bunyi brief §1.1 Tier 2 `FAM` (faktor kedua dan kode pemulihan) dan dua kalimat `CLAUDE.md` (urutan platform; satu pengembang per platform native) | Diterima atas saran lead (keputusan 57, 58); bunyinya rumusan lead | Brief §1.1; `CLAUDE.md` |
| 29 | Jawaban hukum dari konsultasi advokat: pengecualian rumah tangga UU PDP berlaku selama nonkomersial dan tertutup untuk keluarga, termasuk orang tua; tanpa pendaftaran PSE sampai dibuka untuk publik; uang saku pribadi hanya penganggaran, harta bersama tetap menurut UU Perkawinan | Sumbernya **dinyatakan pemilik produk sendiri** (keputusan 59); kutipan pasal adalah pernyataan advokat, dicek sebagian oleh `researcher` | 17 §6; 18 §2.15; 10 §3.4 |
| 30 | Pengelola server tidak membuka data akun orang lain tanpa izin pemiliknya | Diterima atas saran lead (keputusan 60); janji, bukan kontrol teknis | 10 SEC-159; 15 RISK-11 |
| 31 | Bagian lebih dari permintaan yang disetujui, dibayar dengan uang pribadi, bawaannya sumbangan kecuali dicentang | Diterima atas saran lead (keputusan 61) | 07 §3.5 |
| 32 | Nuansa religi pendamping AI: satu pilihan per rumah tangga, Netral bawaan lalu enam agama yang diakui, tanpa Kepercayaan; diubah hanya lewat kesepakatan semua pihak yang wajib | Diterima atas saran lead (keputusan 62, 64, 66) | 05 tier-2 FR-AI-001; 06 §6.3; 08 §5.2 |
| 33 | Kerabat dewasa selain orang tua yang diberi login dihitung pengguna luar | Diterima atas saran lead (keputusan 63) | 13 GR-4; 15 RISK-93 |
| 34 | Teks notifikasi dan kalimat AI dalam bahasa penerima | Diterima atas saran lead (keputusan 65) | 06 §11.1 |
| 35 | Tidak ada konsultasi advokat lanjutan; L-04, L-20 dan klarifikasi kutipan dibuka lagi hanya bila go public diputuskan | **Diputuskan pemilik produk sendiri**, menolak saran lead (keputusan 67); rinciannya bacaan lead | 17 §6; 10 §3.4 |
| 36 | DPIA, catatan pemroses, kebijakan privasi, dan rencana pemberitahuan insiden tertulis tidak disiapkan untuk pilot, ditunda ke go public; kontrol keamanan teknis, termasuk runbook *tahan tulis* dan latihannya (SEC-148 (a)), tetap | **Diputuskan pemilik produk sendiri**, menolak saran lead (keputusan 68); rinciannya bacaan lead | 10 §4.1; 13 §3.3; 14 QG-30 |
| 37 | Klien web (PWA) dibangun lebih dulu dan menjadi klien pertama; aplikasi iOS dan Android menyusul; web belum terasa sepenuhnya native, dan syarat platform-native brief §10.1 dipenuhi aplikasi native kemudian | **Diputuskan pemilik produk sendiri** (keputusan 69), atas usulannya sendiri. Bunyi brief, `CLAUDE.md`, dan ES yang menyesuaikannya menunggu pratinjau (K-71 sampai K-78) | 17 keputusan 69; ADR-030; 08 §25; 10.5 |
| 38 | Bawaan klien web: GR-2 di web saja; bantuan isian juga di TypeScript; ukur kinerja di web; masa tunggu 7 hari dependensi npm; tinjauan sebelum iOS native; iOS lalu Android sesudah GR-2; Mac dari AWS atau pinjaman; RISK-95 diterima; pasang ke layar utama sebelum masuk di iPhone dan iPad; tema mengikuti sistem; Galaxy A55 lewat Chrome | Diterima atas saran lead (bawaan 69.0 sampai 69.10; "selebihnya saya ikut saran kamu") | 17 keputusan 69 |
| 39 | Nama aplikasi sementara **Casa**, selama pemakaian keluarga dan belum go public | **Diputuskan pemilik produk sendiri** (keputusan 70); merek dan ketersediaan dicek sebelum go public (Q-229) | 17 keputusan 70; glosarium *Casa* |

---

## 5. Keputusan yang ditunda, dan kapan diambil

| Keputusan | Kapan | Bahan yang sudah ada | Siapa |
|---|---|---|---|
| **Mac: sumber Mac sekali (AWS EC2 Mac atau pinjaman), lalu sewa bulanan atau beli Mac mini** | Saat jalur iOS native dimulai sesudah GR-2 (bawaan 69.6); sewa bulanan atau beli bila dana ada, sampai itu tetap Xcode Cloud (keputusan 43). Klien web tidak memerlukan Mac | Mac mini M6 16/512 GB, konfigurasi minimum yang disarankan `researcher`, Rp21.499.000 **[Perlu validasi]**; sewa sekitar USD 149 atau EUR 149 per bulan; titik impas untuk 16/512 GB sekitar 7,1 sampai 7,9 bulan, untuk 16/256 GB sekitar 5,8 sampai 6,4 bulan (8,8 bulan terhadap sewa berkomitmen 12 bulan) **[Usulan]** (18 §5.2.7.6); hanya Mac sendiri memberi debugger di iPhone 13 dan iPad mini 7 (18 §5.2.7; 14 §9) | Pemilik produk |
| **Hosting, cadangan, email (S4-2)** | Sebelum deploy pertama yang memegang data nyata (keputusan 31 dan 45), dalam praktik sebelum FS-1 langkah 7 (14 §3.3); sampai itu server uji sementara berisi data uji, dibuang sesudahnya | ADR-020, ADR-026 (Diusulkan); harga dan pihak kontrak dicek ulang saat itu | Pemilik produk |
| **Jenis akun Google Play: konfirmasi** | Hanya bila jalur publik diambil (keputusan 51), sebelum akun dibuat | Akun pribadi, diterima atas saran lead (keputusan 48). Sebelum dibuat, dukungan Play Console ditanya apakah aplikasi keuangan rumah tangga wajib akun organisasi ("Financial products and services"; 18 §2.4). Akun pribadi baru wajib closed testing dengan sedikitnya 12 penguji selama 14 hari sebelum produksi (14 §6.7 **[Terverifikasi]**); ukuran beta di 13 §6.4 | Pemilik produk |
| **Isi dan harga Plus** | Hanya bila jalur publik diambil (GR-5) | Calon isi dan daftar "tidak pernah berbayar" (13 §7.5; Q-36, Q-216); aturan PG-15 | Pemilik produk, sesudah wawancara dan beta |
| **Go public atau tidak** | Tidak dijadwalkan (keputusan 51). Bila pemilik produk kelak mempertimbangkannya: GR-3, paling cepat tiga bulan sesudah GR-2; mundur bila syarat keputusan 46 belum terpenuhi | Syarat keputusan 46 (nyaman dan aman menurut kedua pasangan; 14 hari tanpa bug baru yang berat); VL-01 sampai VL-10 dan tabel keputusan 13 §6.2 | Pemilik produk |
| **Urusan hukum untuk go public** (L-04, L-20, klarifikasi kutipan; badan hukum, kebijakan privasi, DPIA, pendaftaran PSE) | Hanya bila go public diputuskan (keputusan 67 dan 68) | Jawaban advokat dan pemeriksaan kutipannya (17 §6; 18 §2.15); 10 §3.4 | Pemilik produk |
| **Siapa mengoperasikan server bila pemilik produk berhalangan lama** | Sebelum GR-2 | Memberi istri akses pengelola juga memberinya akses teknis ke data pribadi suaminya, cermin RISK-11 (Q-222, Pagar) | Pemilik produk dan istrinya |
| **Bunyi baru brief, `CLAUDE.md`, dan ES untuk klien web** (K-71 sampai K-78) | Sebelum fase bangun dimulai (ADR-030: sebelum FS-1 langkah 1) | Pratinjau kata demi kata dari lead | Pemilik produk |
| **Nama tetap bila go public** | Hanya bila go public diputuskan | Casa sementara (keputusan 70); merek dan ketersediaan belum dicek (Q-229) | Pemilik produk |
| Butir keamanan sebelum pengguna luar pertama (akses darurat SEC-36 dengan log di luar VPS, pemberitahuan pengelola server versi keluarga, dan lainnya) | Sebelum pengguna luar pertama, termasuk orang tua (keputusan 52; 14 QG-31 (a)) | 10 §4.2 | Tim |
| E2EE chat dan foto, SQLCipher, passkey, jalur `AI` R2, uji penetrasi pihak luar (SEC-157) | Sebelum tahap publik, bila diambil (GR-5) | 09 §6; 10 §4 | Tim, satu ADR per butir bila mengubah arsitektur |

---

## 6. Yang perlu ditimbang ulang

Bukan usulan perubahan sekarang. Ini hal yang tim sarankan diperiksa dengan data pilot, karena buktinya belum ada atau risikonya tinggi.

| # | Hal | Mengapa ditimbang ulang | Kapan dan dengan apa |
|---|---|---|---|
| 1 | **Kehadiran "Semua seperti WhatsApp"** (online, terakhir aktif menit tepat, sedang mengetik) | Keputusan pemilik produk sendiri; `security-privacy` dan `ux-designer` mencatat bahwa kehadiran memperlihatkan kapan pasangan memegang aplikasi, juga di jam tenangnya (17 S4-5; W-18). Tidak diubah tanpa pemilik produk | Tinjauan tiga bulan: pertanyaan VL-08 apakah ada yang merasa diawasi (13 §3.4, §6) |
| 2 | **Modul Tier 2 yang mungkin tidak dipakai** (`SHP`, `MNT`, `CAL`, `PLN`) | Setiap fitur harus membayar usaha berulangnya (PRN-02); risiko menambah beban navigasi (15 RISK-79) | Gelombang T2-3 (Q-212); dicabut bila tidak dipakai |
| 3 | **Mutu kumpulan kalimat `AI` R1** | Belum diuji; saran bisa terasa palsu atau berulang bila keduanya memakainya (15 RISK-74) | Gelombang T2-2; F-AI-005 lebih dulu |
| 4 | **Alur persetujuan yang terasa mengekang** | Permintaan dua arah adalah pembeda (18 §1.4), tetapi dapat dibaca sebagai kontrol (15 RISK-21, RISK-31) | VL-08 dan percakapan akhir bulan pilot |
| 5 | **"Sembunyikan angka" (F-FIN-027)** | Tetap Tier 3 Could (S4-10 (iii) A); angka tampil saat aplikasi dibuka, tombol mata menyembunyikan | Bila pilot atau beta menunjukkan perlunya |
| 6 | **Tiga konsep Tier 2 yang belum punya FR** (IDEA-15 tinjauan bulanan, IDEA-19 "butuh tenang hari ini", IDEA-08 templat Mudik dan Lebaran) | Diterima di S3-2, tetapi persyaratannya belum ditulis (04 §9.4) | FR ditulis `product-analyst` sebelum gelombang T2-3 (13 §12); IDEA-08 sebelum Ramadan pertama di dalam pilot (15 RISK-27) |
| 7 | **Struktur Phase 0 dan Phase 2 brief §13** | 13 mengubahnya: validasi pilot tidak menahan pembangunan, dan Phase 2 berarti membawa MVP ke publik, bukan menambah modul (13 §10 TB-2, TB-3, TB-6 sesudah keputusan 51) | Diterima bila 13 diterima; brief §13 sendiri mengizinkan struktur lain |

---

## 7. Peluang teratas

Dasar pembedaannya [18 §1.4](18-market-research.md#14-peluang-pembedaan-positioning-dan-value-proposition), seluruhnya **[Usulan]**: halaman pesaing yang diperiksa tidak menyebut hal-hal ini, tetapi itu belum membuktikan fiturnya tidak ada di produk mereka. Dua atau tiga aplikasi pesaing perlu diuji langsung sebelum diklaim di materi publik (18 §6.4; 13 GR-5). Bagian ini hanya relevan bila jalur publik kelak diambil (keputusan 51).

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

Register lengkapnya di [15](15-risk-assessment.md) (RISK-01 sampai RISK-97). Empat kelompok yang dapat menggagalkan tujuan, dan satu yang hanya berlaku di jalur opsional (15 §1):

1. **Hanya satu pasangan mencatat, mencatat terlalu lama, atau catatan terlewat** (RISK-18 sampai RISK-20). Penawarnya: data nyata dimulai lewat klien web di perangkat kedua pasangan (iPhone istri, pencatat utama; iPad dan Galaxy A55 pemilik produk; keputusan 69), tiga ketukan, entri terlambat diterima, rekonsiliasi kas. Khusus klien web: catatan luring yang belum terkirim dapat hilang bila browser menghapus penyimpanannya; penawarnya web app layar utama, antrean dikirim begitu daring, dan kehilangan dinyatakan ke pengguna (RISK-94; SEC-170 sampai SEC-173).
2. **Angka uang salah** (RISK-47, RISK-48). Penawarnya: setiap invarian 06 dan 07 sebagai uji yang memblokir merge, kunci idempotensi, satu transaksi per perintah; bila pemeriksaan invarian harian tidak nol, server berhenti menerima tulisan baru lewat *tahan tulis* sampai penyebabnya diperbaiki, sementara perangkat tetap menyimpan catatan di antrean dan perintah pelindung tetap berlaku seketika (ADR-028 butir 3, diterima atas saran lead, keputusan 56; SEC-151; PRN-06).
3. **Aplikasi menjadi alat kendali antaranggota** (RISK-11, RISK-35, RISK-38, RISK-40). Penawarnya: pagar 06 sebagai uji, simetri per peran, keluar dengan aman tanpa persetujuan siapa pun, halaman Bantuan tanpa jejak.
4. **Operasi bergantung pada satu orang** (RISK-10): pemilik produk adalah pemimpin teknis, pengelola server, dan pemegang kunci cadangan. Penawarnya: runbook sebelum GR-2, komponen sedikit, salinan kunci cadangan offline, Q-222.
5. **Jalur opsional, bila kelak diambil** (keputusan 51, 52): pengguna luar pertama, termasuk orang tua (PSE bila wajib, RISK-65; pengelola server yang anak mereka sendiri, RISK-93, Pagar), atau tahap publik (keputusan GR-3 atas bukti satu rumah tangga, RISK-25; kesediaan membayar, RISK-58).

---

## 9. Usulan proses untuk fase bangun

| # | Usulan | Asal | Kedudukan |
|---|---|---|---|
| 1 | Tinjauan `security-privacy` hanya-temuan sebelum GR-2 (data nyata pertama) dan sebelum GR-4 (pengguna luar pertama), sebagai pengganti peran keamanan khusus di tim bangun; ditambah sebelum aplikasi iOS native dan aplikasi Android memegang data nyata, dan setahun sekali | 14 §12 (`architect`); 10 §3.3, §4.4 (`security-privacy`) | Diterima atas saran lead (keputusan 49, 53; bawaan 69.4); dipasang di `CLAUDE.md` "Build team" (pembaruan untuk iOS native menunggu pratinjau, K-77) |
| 2 | Peran bangun baru `web-dev` untuk klien web; `ios-dev` dan `android-dev` mulai sesudah GR-2 | 14 WS-8 (`architect`) | Diterima atas saran lead (keputusan 69); bunyi `CLAUDE.md` menunggu pratinjau (K-77) |

---

## 10. Yang tersisa sesudah S6

- **S6 — pass konsistensi** (brief §0.3) selesai 12 November 2026: kontradiksi, tautan, istilah, dan keterlacakan diperbaiki; butir W1–W7 diterapkan; laporan perubahan di [_progress.md](_progress.md), bagian "S6 — Pass konsistensi".
- **Sesudah S6** (12–13 November 2026): audit kesiapan bangun oleh `critic` (siap dengan perbaikan; temuannya diperbaiki di putaran ini), lalu putaran web lebih dulu (keputusan 69) dan nama Casa (keputusan 70): 08 §25, 10.5, ADR-030, 09, 13, 14, 15, 17, 05, 18 §5.13 diperbarui; rinciannya di [_progress.md](_progress.md) bagian "Sesudah S6".
- **Fase perencanaan berakhir hanya bila pemilik produk menyatakannya** (`CLAUDE.md`). Sebelumnya: pratinjau bunyi brief, `CLAUDE.md`, dan ES (K-71 sampai K-78). Sesudahnya: agen tim bangun ditulis lead menurut stack yang dipilih (`web-dev` lebih dulu), lalu GR-1.

---

## Riwayat

| Tanggal | Perubahan |
|---|---|
| 12 November 2026 | Draf pertama (lead, S5 fase C) |
| 12 November 2026 | Perbaikan atas `critic` S5 B (CS5-B-01 kedudukan keputusan) dan tinjauan ulang B2 (CR5-B-01, CR5-B-05, CR5-B-06, CR5-B-07, CR5-B-13): §1 butir 2 per keputusan bernomor; §4 baris 15 keputusan 39 dan W-34; §5 Q-215, syarat akun Play berlabel, kutipan keputusan 31 persis; salinan data sendiri (FR-FAM-013/AC-12) dibangun di FS-1 menurut ketetapan lead, jadi keluar dari daftar sebelum GR-4; §8 status ADR-028 |
| 12 November 2026 | Keputusan 42 (pilot di iOS dan iPadOS; Android menyusul): §1 butir 3 dan 6, §4 baris 16, §5 baris platform dihapus dan waktu keputusan Mac, §8 butir 1 |
| 12 November 2026 | Keputusan 43–50 (penutupan pertanyaan terbuka sebelum S6): §1 butir 4 dan 6; §3 baris build iOS, lingkungan dan rilis, hosting, tradeoff 2; §4 baris 16 sampai 22; §5 baris Mac, hosting, akun Play, pemakaian harian (dihapus), wawancara (dihapus), GR-3; §9 diterima |
| 12 November 2026 | Keputusan 51–55 (aplikasi untuk keluarga sendiri berstandar publik; orang tua pengguna luar; tinjauan keamanan tambahan; faktor kedua di FS-2; jalur pasang iOS): §1 butir 4 sampai 6; §3 baris build iOS; §4 baris 8, 10, 16, 21, 22 dan baris baru 23 sampai 26; §5 baris akun Play, Plus, go public, advokat (baru), klien web, keamanan sebelum pengguna luar, tahap publik; §6 baris 1, 7; §7; §8 kelompok 4 dan 5, RISK-93; §9; §2 hitungan F-FAM-022 |
| 12 November 2026 | S6 (pass konsistensi): keputusan 56–68; §0 tautan *Tingkat kepastian*; §1 butir 5 dan 6; §3 baris lingkungan dan rilis; §4 baris 10 dan baris baru 27 sampai 36; §5 baris advokat diganti urusan hukum untuk go public; §10 |
| 13 November 2026 | Sesudah S6: audit kesiapan bangun (temuan 4, 7, 8 dokumen ini: waktu S4-2, "tiga belas butir", baris advokat keputusan 24) dan keputusan 69, 70: §1 butir 2, 3, 6; §2 baris Tier 3 dan daftar yang tidak dibangun; §3 baris klien pertama (baru), aplikasi native, sinkron, hosting, build iOS, lingkungan dan rilis, biaya, tradeoff 2; §4 baris 3, 14, 16, 17, 18, 22, 26 dan baris baru 37 sampai 39; §5 baris Mac, hosting, bunyi dokumen aturan (baru), nama (baru), klien web tahap publik (dihapus); §8; §9 baris 2 (baru); §10 |
