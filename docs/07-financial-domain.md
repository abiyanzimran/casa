# 07 — Domain Keuangan

| | |
|---|---|
| Pemilik dokumen | `finance-domain` |
| Tahap | S2 (Fondasi domain) |
| Status | Draf untuk tinjauan `critic`; direvisi menurut keputusan pemilik produk 11 November 2026 dan permintaan `architect` (dokumen 06 §15.4); direvisi lagi menurut temuan pemeriksaan awal `security-privacy` atas dokumen 06 (SP-01, SP-05, SP-12, SP-13, SP-18); diselaraskan terakhir dengan dokumen 06 §15.4 butir (e) sampai (m) dan dengan pernyataan pemilik produk 11 November 2026 bahwa kedua pasangan berpenghasilan; direvisi menurut keputusan pemilik produk 11 November 2026 tentang jatah pribadi bulanan dan hari mulai periode anggaran (§3.4, §6.1, CTH-39); direvisi menurut tinjauan `critic` atas S2 (label keputusan lawan usulan pada jatah pribadi dan hari mulai, rujukan S2-1 butir (i) sampai (iv), fungsi `periode()`, andaian ambang CTH-06, CTH-33, CTH-34, CTH-39) dan temuan `security-privacy` SP-24 dan SP-28, tanpa satu angka contoh pun yang berubah |
| Tanggal | 11 November 2026 |
| Dibaca sebelum menulis | `CLAUDE.md`; `family-os-brief.md` §1.1, §5C, §5D, §7, §16; `ENGINEERING-STANDARDS.md` §5; [00-glossary.md](00-glossary.md); [_progress.md](_progress.md); [17-open-questions-assumptions.md](17-open-questions-assumptions.md); [18-market-research.md](18-market-research.md) §2.1; untuk revisi: [06-activity-approval-architecture.md](06-activity-approval-architecture.md) §2 sampai §8, §15; ADR-001 sampai ADR-010 |

Aturan uang Family OS, termasuk jumlah pada *penyelesaian*, kanonis di dokumen ini. Dokumen lain mengutipnya. Istilah mengikuti [glosarium bagian 7](00-glossary.md#7-istilah-keuangan), yang difinalkan bersama dokumen ini.

**Label.** Keputusan pemilik produk di checkpoint S1 dikutip sebagai "keputusan pemilik nomor *n*", dengan nomor menurut dokumen 17 (bagian "Yang diputuskan pemilik produk sendiri"); default yang ia terima lewat "lanjut" dikutip dengan nomor `Q`-nya. Keputusannya sesudah checkpoint itu, dalam percakapan dengan lead, dikutip sebagai "keputusan pemilik produk, 11 November 2026". Semua rekomendasi lain di dokumen ini adalah **[Usulan]** `finance-domain` kecuali diberi label lain. Angka contoh selain angka brief adalah **[Asumsi]**: pemilik produk tidak memberi data nyata.

**Dua lapis.** Dokumen ini memakai bahasa model internal (entri, baris, sebab, kewajiban). Keluarga tidak pernah melihat kata-kata itu; padanan layarnya ada di §1.4 dan ditetapkan dokumen 08.

**Kesetaraan.** Semua aturan ditulis untuk "anggota dewasa". Kata "suami" dan "istri" hanya muncul di contoh, karena brief §1.1 menggambarkan rumah tangga pilot dengan kata itu.

## Daftar isi

1. Keputusan ledger
2. Model
3. Aturan per jenis peristiwa
4. Invarian
5. Tampilan custody
6. Anggaran dan laporan
7. Contoh terhitung
8. Riwayat yang tidak dapat diubah lawan hak penghapusan
9. Keputusan atas pertanyaan terbuka
10. Butir yang masih terbuka dan permintaan ke pihak lain

---

## 1. Keputusan ledger

### 1.1 Yang dituntut rumah tangga ini

| # | Kebutuhan | Asal |
|---|---|---|
| N1 | Enam dimensi custody terjawab terpisah untuk setiap jumlah uang | Brief §5C |
| N2 | Uang yang berpindah antar-orang dan antar-tempat tidak pernah terhitung dua kali | Brief §5C, §7 |
| N3 | Koreksi hanya dengan menambah catatan | Brief §7; `ENGINEERING-STANDARDS.md` §5 |
| N4 | Pencatat pelupa: entri terlambat, selisih kas, tanpa penutupan periode | Brief §1.1; keputusan pemilik nomor 7 |
| N5 | Invarian dapat dijalankan sebagai property test | `ENGINEERING-STANDARDS.md` §5, §6 |
| N6 | Tim 1–3 developer; satu VPS | Brief §1 |
| N7 | Utang, piutang, cicilan, aset, dan kekayaan bersih ada di Tier 1 | Brief §1.1 |
| N8 | Dana pribadi tertutup, tetapi kontribusinya ke rumah tangga terlihat, dan penghasilan yang diarahkan ke dana pribadi saat dicatat terlihat beserta jumlahnya | Keputusan pemilik nomor 3; keputusan pemilik produk, 11 November 2026 |

### 1.2 Perbandingan

| Kriteria | A. Double-entry penuh | B. Ledger sederhana: entri bertipe dengan baris seimbang | C. Daftar transaksi polos |
|---|---|---|---|
| Bentuk | Bagan akun bebas; setiap jurnal berisi debit dan kredit ke akun mana pun | Jenis entri tertutup (§2.5). Tiap entri punya baris *posisi* (tempat uang, piutang, kewajiban, aset) dan baris *sebab* (pemasukan, pengeluaran, selisih kas, dan lain-lain) yang jumlahnya harus sama | Satu baris per transaksi: jenis, jumlah, satu atau dua tempat uang |
| N1 custody | Bisa, tetapi pemegang dan pemilik harus dikodekan sebagai akun; bagan akun membengkak | Pemegang dan pemilik adalah atribut tempat uang; tujuan, penerima manfaat, pelaku adalah atribut baris | Bisa untuk pemegang; tujuan dan pemilik ditempel sebagai kolom tambahan |
| N2 tanpa hitung ganda | Terjamin oleh keseimbangan jurnal | Terjamin oleh keseimbangan baris (INV-02) | Tidak terjamin: pemindahan adalah dua baris atau satu baris berpengecualian; kesalahannya baru terlihat di laporan |
| N3 koreksi | Jurnal pembalik | Entri pembalik: baris yang sama dengan tanda terbalik | Baris negatif; arti "membalik" berbeda untuk tiap jenis |
| N4 pencatat pelupa | Selisih kas adalah jurnal ke akun selisih | Selisih kas adalah jenis entri sendiri; entri terlambat dapat menutupnya tanpa menyentuh saldo | Selisih kas bisa; "menutup selisih" memerlukan aturan khusus |
| N5 invarian | Satu invarian umum, tetapi aturan domain (misalnya "talangan tidak dihitung dua kali") tetap harus ditulis di atasnya | Satu invarian umum ditambah aturan per jenis entri; keduanya dapat diuji | Tidak ada invarian struktural; setiap laporan membawa aturannya sendiri |
| N7 kekayaan bersih | Neraca langsung | Jumlah baris posisi; sama dengan jumlah baris sebab (INV-07) | Dihitung dari beberapa tabel terpisah; mudah tidak cocok |
| N8 dana pribadi | Perlu buku per pemilik | Tiap baris membawa pemilik; keseimbangan diperiksa per pemilik | Kolom pemilik; tanpa pemeriksaan |
| N6 biaya | Tertinggi: bagan akun, pemetaan setiap layar ke jurnal, risiko kosakata akuntansi bocor ke layar | Sedang: satu tabel entri, satu tabel baris, daftar jenis yang tertutup | Terendah di awal, tertinggi kemudian: setiap fitur Tier 1 (cicilan, hak ganti, aset) menambah pengecualian |

### 1.3 Rekomendasi

**Opsi B [Usulan].** Secara matematis B adalah double-entry: setiap entri seimbang. Bedanya dari A, tidak ada bagan akun yang dapat diatur: "akun" hanya ada dalam dua bentuk yang sudah dimiliki keluarga, yaitu *posisi* (tempat uang, piutang, kewajiban, aset yang mereka daftarkan sendiri) dan *sebab* (sembilan nilai tetap di §2.4).

Alasan yang terikat pada rumah tangga ini:

1. Arus uang pilot didominasi perpindahan, bukan belanja: gaji masuk rekening, tarik tunai, serah terima ke pasangan, pinjam kembali (brief §1.1). Daftar transaksi polos paling lemah justru di sini (N2), dan kriteria sukses nomor 1 bergantung padanya.
2. Tier 1 memuat cicilan, pinjaman ke keluarga besar, hak ganti, aset, dan kekayaan bersih (N7). Semua itu adalah posisi yang bukan uang. Pada opsi C tiap posisi menjadi tabel dengan aturan sendiri; pada opsi B semuanya baris dari entri yang sama dan tunduk pada satu invarian.
3. Dana pribadi tertutup tetapi kontribusinya terlihat (N8). Itu menuntut satu peristiwa tercatat di dua buku pemilik tanpa terhitung dua kali; keseimbangan per pemilik (INV-02, INV-03) menjaminnya.
4. Tim kecil (N6): opsi A menambah bagan akun dan pemetaan jurnal yang tidak dipakai siapa pun di rumah tangga ini.

Yang tidak diberikan opsi B dan diterima: akun bebas buatan pengguna, laporan laba-rugi dan neraca berformat akuntansi, multi-mata-uang dalam satu entri. Tidak satu pun dibutuhkan pilot; yang terakhir dapat ditambahkan kemudian karena tiap baris membawa mata uangnya.

### 1.4 Yang dilihat keluarga lawan model internal

Tidak ada kata di kolom kiri yang tampil di layar. Kolom kanan adalah usulan untuk dokumen 08, yang menetapkan kata akhirnya.

| Model internal | Yang dilihat keluarga (usulan) |
|---|---|
| Entri, baris posisi, baris sebab | "Catatan" |
| Tempat uang | "Dompet", "Rekening", "E-wallet" beserta nama pemegangnya |
| Pemindahan, serah terima | "Pindah uang", "Kasih ke …", "Ambil dari …" |
| Pemasukan | "Uang masuk" |
| Pengeluaran | "Belanja", "Bayar" |
| Kewajiban | "Utang kita", "Sisa cicilan" |
| Piutang | "Dipinjam …" |
| Hak ganti | "Minta diganti" |
| Kontribusi dari dana pribadi | "Pakai uang pribadi …" |
| Selisih kas | "Selisih hitung" |
| Entri pembalik, entri penyesuaian | "Dibatalkan", "Diperbaiki" |
| Pelepasan klaim dana bersama atas seorang anggota | "Utang dihapuskan" |
| Penutupan klaim saat anggota keluar | "Ditutup karena anggota keluar" |
| Otorisasi, terikat | "Disetujui", "Sudah dijanjikan" |
| Reservasi | "Disisihkan untuk …" |
| Jatah pribadi (`set_aside_schedule`) | "Jatah pribadi" |
| Periode anggaran | "Bulan ini", beserta rentang tanggalnya |
| Kekayaan bersih | "Harta bersih" |
| Uang yang ada | "Uang kita" |

### 1.5 Yang perlu dicatat `architect` sebagai ADR

Ketujuh keputusan di bawah sudah ditulis `architect` sebagai ADR-002 sampai ADR-008, berurutan (ADR-A = ADR-002, …, ADR-G = ADR-008). Kalimat ADR yang perlu disesuaikan setelah revisi dokumen ini didaftar di §10.2.

| # | Keputusan untuk ADR | Bagian |
|---|---|---|
| ADR-A | Model ledger: opsi B, entri bertipe dengan baris seimbang; jenis entri dan sebab adalah daftar tertutup | §1.3, §2 |
| ADR-B | Dana pribadi berada di luar *garis kendali rumah tangga*; satu buku per *pemilik uang*; keseimbangan per pemilik | §2.2 |
| ADR-C | Saldo adalah turunan dari baris; bila disimpan sebagai cache, diperbarui dalam transaksi basis data yang sama dengan entrinya | §2.5, INV-04 |
| ADR-D | Idempotency key per perintah tulis, unik per rumah tangga, disimpan bersama entri | §2.5, INV-10 |
| ADR-E | Tidak ada penutupan periode; laporan periode lampau dihitung ulang dan menandai perubahan | §3.11 |
| ADR-F | Tafsir "append-only" terhadap hak penghapusan: inti moneter tetap dan dipseudonimkan, rincian pribadi terpisah dan dapat dihapus. Keputusan sementara sampai L-04 dijawab | §8 |
| ADR-G | Jumlah uang: bilangan bulat bertanda 64-bit dalam satuan terkecil; eksponen satuan per mata uang (IDR = 0) | §2.1 |

---

## 2. Model

Sketsa di bagian ini adalah sketsa skema, bukan kode. Nama bahasa Inggris adalah usulan nama pengenal.

### 2.1 Jumlah uang

- `amount`: bilangan bulat bertanda dalam satuan terkecil mata uangnya, selalu bersama `currency`.
- Satuan terkecil IDR adalah **rupiah utuh** (eksponen 0). Saldo rekening bank yang memuat sen dicatat dibulatkan ke bawah ke rupiah utuh; sisa sen tidak dicatat. Bila kelak mengganggu, ia tertangkap sebagai *selisih kas* **[Usulan]**.
- Tidak ada pembagian yang menghasilkan pecahan di dalam model. Pembagian satu jumlah ke beberapa baris (pengeluaran terbagi) diisi sebagai jumlah per baris oleh pencatat; bila aplikasi menawarkan "bagi rata", sisa pembagian ditaruh pada baris pertama, sehingga jumlah baris selalu sama dengan jumlah entri.
- MVP memakai satu mata uang per rumah tangga. Entri yang memuat dua mata uang ditolak.

### 2.2 Pemilik uang, buku, dan garis kendali rumah tangga

`money_owner` adalah salah satu dari:

| Nilai | Arti |
|---|---|
| `shared` | Rumah tangga secara bersama (*dana bersama*) |
| `member(M)` | Satu anggota dewasa (*dana pribadi* M) |

Setiap pemilik uang punya satu **buku**: kumpulan semua baris yang bertanda pemilik itu. Buku `shared` adalah **buku rumah tangga**.

**Buku pribadi per anggota, bukan per akun [Usulan]** (dikukuhkan atas permintaan dokumen 06 §15.4 butir (c); dokumen 06 §14; ADR-003). `member(M)` berarti *anggota*, yaitu keanggotaan seseorang dalam **satu** rumah tangga, bukan *akun* login-nya; setiap buku, entri, dan baris milik tepat satu rumah tangga (`entry.household_id`, §2.5), dan semua invarian §4 diperiksa di dalam satu rumah tangga. Akibatnya, orang yang ikut dua rumah tangga punya dua buku pribadi yang tidak saling tahu, tanpa entri lintas rumah tangga: uang pribadinya yang secara fisik satu dicatat sendiri-sendiri di tiap rumah tangga tempat ia memakainya, dan baris kedua "uang pribadi saya" (§5) hanya memuat buku pribadi rumah tangga yang sedang dibuka.

**Uang antara dua rumah tangga yang sama-sama memakai aplikasi [Usulan]** (Q-86; dokumen 03 bagian 4.3). Dua catatan yang berdiri sendiri, tanpa tautan. Rumah tangga pengirim mencatatnya menurut aturannya sendiri dengan penerima sebagai *pihak luar* (kiriman keluarga, atau `loan_given`); rumah tangga penerima mencatatnya menurut aturannya sendiri dengan pengirim sebagai *pihak luar* (`income` bersumber `gift` atau `other`, atau `loan_received`). Tidak ada entri lintas rumah tangga, tidak ada pengenal bersama, tidak ada pencocokan otomatis, dan tidak ada notifikasi dari satu rumah tangga ke rumah tangga lain; invarian §4 tiap rumah tangga diperiksa sendiri, dan bila kedua catatan berbeda jumlah atau salah satunya tidak ada, tidak ada yang mendamaikannya. Ini berlaku juga bila satu orang menjadi anggota keduanya. Alasannya isolasi rumah tangga: tautan apa pun memberi tahu satu rumah tangga bahwa, kapan, dan berapa rumah tangga lain mencatat. Bentuk yang lebih kaya (misalnya penerima mengonfirmasi kiriman) ditunda ke tahap publik (Tier 3) dan memerlukan tinjauan `security-privacy`.

**Keputusan: dana pribadi berada di luar garis kendali rumah tangga [Usulan].** Di dalam garis hanya dana bersama. Akibatnya:

| Hal | Akibat |
|---|---|
| *Uang yang ada* | Jumlah saldo tercatat semua tempat uang berpemilik `shared`. Inilah angka utama, dan sama bagi kedua pasangan. Tidak perlu nama kedua (menutup K-13) |
| *Kekayaan bersih* rumah tangga | Hanya dari buku rumah tangga |
| Dana bersama → dana pribadi | Uang melintas keluar garis: sebab `owner_transfer` (*penyisihan*), bukan pemindahan |
| Dana pribadi → dana bersama, atau membayar keperluan rumah tangga | Uang atau nilai melintas masuk: sebab `owner_transfer` (*kontribusi dari dana pribadi*) |
| Penghasilan (gaji, hasil usaha) yang diarahkan pencatatnya ke dana pribadinya saat dicatat | Pemasukannya tetap tercatat penuh di buku rumah tangga; bagian pribadinya keluar garis dalam entri yang sama sebagai `owner_transfer` (penyisihan), sehingga terlihat oleh keduanya (§3.1, INV-24). **[Usulan]**: hanya bila pencatatnya penerima penghasilan itu (INV-27) |
| Klaim antara anggota dan dana bersama | Tampil di buku rumah tangga sebagai kewajiban atau piutang rumah tangga kepada anggota itu |
| Buku pribadi | Tetap berada di ledger yang sama dan tunduk pada invarian yang sama; hanya terlihat oleh pemiliknya (§5) |

Alasan:

1. Keputusan pemilik nomor 3: saldo dana pribadi tertutup bagi pasangan. Angka rumah tangga yang memuat dana pribadi tidak dapat ditampilkan sama kepada keduanya tanpa membocorkannya lewat pengurangan (Q-23). Dengan dana pribadi di luar garis, soal itu hilang secara struktur, bukan lewat penyaringan tampilan.
2. Keputusan pemilik nomor 3 juga menyebut pembayaran keperluan rumah tangga dengan uang pribadi sebagai **kontribusi**. Kontribusi hanya bermakna bila ada garis yang dilintasi.
3. Keputusan pemilik nomor 2: hadiah untuk satu orang menjadi dana pribadinya. Dengan dana pribadi di luar garis, hadiah itu tidak menaikkan "uang kita", sesuai maksudnya.

Biaya yang diterima: kekayaan bersih rumah tangga tidak memuat harta pribadi. "Berapa harta keluarga seluruhnya" tidak dijawab di layar mana pun, karena jawabannya hanya dapat ditampilkan dengan membuka dana pribadi. Tiap anggota melihat baris kedua, "uang pribadi saya" (§5).

**Aturan tulis buku pribadi [Usulan]** (menjawab dokumen 06 §7.3 dan §15.4 butir 4; INV-25). Perintah yang menulis baris berpemilik `member(M)` hanya diterima dari M. Pencatat lain tidak mengenal tempat uang pribadi M, dan tidak seorang pun dapat dicatat menerima, memakai, atau berutang atas dana pribadinya oleh orang lain. Akibatnya:

| Entri dua buku (buku rumah tangga dan buku pribadi M) | Dicatat oleh |
|---|---|
| Penyisihan kepada M; pemasukan dengan bagian pribadi M (§3.1) | M. Untuk penghasilan, **[Usulan]**: hanya bila M juga `earner`-nya (INV-27) |
| Kontribusi uang, kontribusi belanja, talangan oleh M | M |
| Penggantian dari dana bersama kepada M | M, yaitu penerimanya. Baris yang mengurangi tempat uang dana bersama yang dipegang anggota dewasa lain bertanda `unconfirmed` di sisi itu (§3.3; dokumen 06 §7.2 "Aturan umum") |
| Pinjaman antara dana bersama dan dana pribadi M, dan pelunasannya | M; tanda `unconfirmed` yang sama atas baris yang mengurangi tempat uang dana bersama yang dipegang anggota dewasa lain |
| Pelepasan klaim M atas dana bersama (hak ganti) | M, sebagai pemilik piutangnya (§3.6) |
| Pelepasan klaim dana bersama atas M | Anggota dewasa selain M. **Satu-satunya pengecualian bagi pencatat manusia**: entri ini hanya memuat baris kewajiban − dan `owner_transfer` + di buku M, tidak menyentuh tempat uang M, dan hanya menguntungkan M (§3.6) |
| Entri pembalik atas entri dua buku | M (§3.12) |
| Penutupan klaim saat M keluar (`exit_closure`) | Sistem, bukan anggota mana pun (§8.2, INV-28) **[Usulan]** |

Aturan ini sejalan dengan keputusan pemilik nomor 1 (yang menerima uang yang mencatat) dan dengan keputusan pemilik produk 11 November 2026 (yang mencatat uang masuk memilih arahnya): arah "ke dana pribadi" selalu berarti dana pribadi **pencatat sendiri**. Untuk penghasilan, **[Usulan]** yang menunggu pemilik produk di checkpoint S2: arah itu hanya tersedia bila pencatatnya juga penerima penghasilan itu (INV-27, §3.1). Persinggungannya dengan "kedua pasangan boleh mengoreksi entri dana bersama" diatur di §3.12.

Batas yang diterima: entri yang menyentuh buku pribadi **dua** anggota (talangan, pinjaman, atau penggantian antara dua dana pribadi; CTH-17b) tidak dapat dihasilkan satu perintah, sebab tidak ada satu pelaku yang berhak atas kedua buku. Modelnya tetap ada dan invariannya tetap diuji, tetapi alur pencatatannya **ditunda dari MVP**: arus uang rumah tangga pilot (brief §1.1) tidak memuatnya, dan angka rumah tangga tidak tersentuh olehnya. Jalan menambahkannya kelak: satu tempat uang pribadi bawaan per anggota sebagai tujuan baris masuk yang dicatat orang lain, bertanda belum dikonfirmasi (T-12).

### 2.3 Posisi

Posisi adalah sesuatu yang bersaldo. Setiap posisi punya tepat satu pemilik uang. Semuanya dibuat, diganti nama, dan diarsipkan oleh rumah tangga sendiri (keputusan pemilik nomor 9).

```text
holding                              -- tempat uang
  id, household_id
  name
  kind            cash | bank | ewallet | entrusted_external
  owner           shared | member(M)          -- tepat satu pemilik (Q-44)
  holder          member(M) | joint | external_party(P)
  visibility      private | spouses | all_members   -- ditetapkan dokumen 06 dan 10
  currency
  archived_at     null | timestamp

receivable                           -- piutang
  id, household_id, owner
  counterparty    external_party(P) | member(M) | shared
  origin          loan_given | advance | opening
  archived_at

liability                            -- kewajiban
  id, household_id, owner
  counterparty    external_party(P) | member(M) | shared
  kind            debt | installment | reimbursement_claim | inter_owner_loan | opening
  archived_at

asset                                -- aset
  id, household_id, owner
  name, kind      vehicle | gold | property | device | other
  archived_at
```

Aturan:

- **Pemegang adalah atribut tempat uang** (Q-20). Uang berganti pemegang dengan berpindah ke tempat uang lain, tidak dengan mengubah atribut `holder`. Atribut `holder` hanya dapat diubah bila saldonya nol.
- **Satu tempat uang, satu pemilik** (Q-44). Uang tunai pribadi dan uang tunai bersama yang dipegang orang yang sama adalah dua tempat uang, walau secara fisik satu dompet. Biaya usahanya nyata: pemegang harus memisahkannya di kepala. Rekonsiliasi kas pada dompet campuran menghitung seluruh isi dompet lalu membandingkannya dengan jumlah saldo kedua tempat uang; selisihnya dibebankan pemegang ke salah satu (bawaan: yang berpemilik `shared`, karena itulah yang dilihat pasangan).
- **Tempat uang bersama** (`holder = joint`) tidak punya satu pemegang. Custody melaporkannya sebagai baris tersendiri, "dipegang bersama", tidak dibagi dua. `joint` menyangkut akses, bukan kepemilikan.
- **Titipan pada pihak luar** adalah tempat uang berjenis `entrusted_external` dengan `holder = external_party(P)` dan `owner = shared`. Ia terhitung dalam *uang yang ada*. Karena pihak luar tidak punya login, pemakaiannya dicatat oleh anggota; bila tidak dicatat, saldonya salah sampai direkonsiliasi. Itu sebabnya kiriman bulanan kepada ibu suami **tidak** dimodelkan sebagai titipan (default Q-07 yang diterima pemilik produk).
- **Klaim antar-pemilik** (*klaim internal*): satu klaim antara dua pemilik uang di rumah tangga yang sama tercatat sebagai sepasang posisi, `receivable` di buku satu pemilik dan `liability` di buku pemilik lain, yang lahir, berubah, dan berakhir dalam entri yang sama dan selalu sama besar (INV-08).
- Saldo piutang, kewajiban, dan aset disimpan sebagai bilangan tidak negatif. Dalam persamaan keseimbangan, kewajiban bertanda minus.

### 2.4 Sebab

Baris sebab menjawab "mengapa kekayaan bersih pemilik ini berubah". Nilainya adalah **akibat pada kekayaan bersih**: positif menaikkan, negatif menurunkan. Daftarnya tertutup.

| `cause` | Arti | Tanda lazim | Membawa |
|---|---|---|---|
| `opening` | Saldo awal | + atau − | — |
| `income` | Pemasukan | + | sumber (`salary`, `gift`, `business`, `other`), atas nama siapa (`earner`), kategori. Sumber yang termasuk **daftar sumber penghasilan** rumah tangga disebut **penghasilan**; daftar itu parameter preset aturan, dan pada preset pilot berisi `salary` dan `business` (INV-24). Baris `income` bersumber penghasilan selalu berpemilik `shared` (§3.1, INV-24) dan wajib membawa `earner`: tepat satu anggota dewasa rumah tangga itu, yaitu isian "punya siapa" di layar, diisi pencatat dan terpisah dari `actor` (INV-27) |
| `expense` | Pengeluaran. Pengembalian dana adalah baris `expense` bertanda + yang tertaut ke aslinya | − | kategori, penerima manfaat, pembayar, pihak luar (opsional), `funded_by` |
| `credit_cost` | Biaya kredit: selisih antara jumlah seluruh cicilan dan nilai yang dibiayai | − | cicilan terkait |
| `cash_difference` | Selisih kas | + atau − | rekonsiliasi terkait |
| `write_off` | Penghapusan piutang | − | piutang terkait, alasan |
| `valuation` | Perubahan nilai aset, termasuk untung-rugi penjualan | + atau − | aset terkait |
| `owner_transfer` | Perpindahan nilai antar-pemilik: penyisihan (keluar dari buku rumah tangga), termasuk bagian pemasukan yang diarahkan langsung ke dana pribadi (§3.1); kontribusi (masuk ke buku rumah tangga); pelepasan klaim antar-pemilik (§3.6); dan penutupan klaim saat anggota keluar (§8.2) | + atau − | pemilik lawan |
| `debt_forgiven` | Utang rumah tangga yang dibebaskan pihak luar | + | kewajiban terkait |

Pemindahan, serah terima, pinjaman (diterima atau diberikan), pelunasan, penggantian, pembayaran cicilan, dan pembelian aset secara tunai **tidak punya baris sebab**: kekayaan bersih tidak berubah.

### 2.5 Entri dan barisnya

```text
entry
  id, household_id
  type              -- daftar tertutup, tabel di bawah
  event_date        -- tanggal kejadian, diisi pencatat; tidak boleh di masa depan
  recorded_at       -- tanggal pencatatan, diisi server; tidak dapat diatur pencatat
  actor             member(M) | system    -- siapa yang mencatat; selalu ada. system hanya pada exit_closure (§8.2)
  idempotency_key   -- dibuat klien satu kali per niat mencatat; unik per rumah tangga
  corrects          null | entry_id       -- entri yang dikoreksi
  correction_kind   null | reversal | adjustment
  links             authorization_id?, installment_id?, recurring_bill_id?,
                    set_aside_schedule_id?, reconciliation_id?, refund_of_entry_id?,
                    subject_ref?          -- rujukan induk (kind, id) menurut dokumen 06 §2.2
  detail_ref        -> entry_detail       -- catatan, struk, nama pihak lawan; disimpan terpisah (§8)

entry_link                                -- tautan susulan (§3.8); hanya ditambah, tidak mengubah entri
  id, household_id, entry_id
  kind              authorization | installment | recurring_bill | set_aside_schedule | subject_ref
  target            -- id sesuai kind
  actor, recorded_at, idempotency_key
  revokes           null | entry_link_id  -- mencabut tautan susulan yang keliru

line                                      -- baris; tidak pernah berdiri tanpa entri
  entry_id, owner                         -- pemilik uang yang bukunya disentuh baris ini
  kind              position | cause | allocation
  -- position:
  position_ref      holding | receivable | liability | asset
  delta             bilangan bulat bertanda (untuk kewajiban: + berarti kewajiban bertambah)
  -- cause:
  cause             nilai §2.4
  effect            bilangan bulat bertanda (akibat pada kekayaan bersih)
  category_id?, beneficiary?, payer?, funded_by?, earner?, source?
  -- allocation:
  holding_id, purpose_id, delta
```

**Keseimbangan.** Untuk setiap entri dan setiap pemilik uang yang disentuhnya:

```text
Σ delta(tempat uang) + Σ delta(piutang) + Σ delta(aset) − Σ delta(kewajiban) = Σ effect(sebab)
```

Baris `allocation` tidak ikut persamaan ini; ia hanya memberi label tujuan (§3.9).

**Rujukan induk.** `links.subject_ref` adalah rujukan induk dokumen 06 §2.2: pasangan `(kind, id)` yang menunjuk permintaan, tugas, masalah rumah, kesepakatan, atau objek lain tempat entri itu menempel. "Aktivitas" bukan primitif (ADR-001); model keuangan tidak mengenal objek induk generik dan tidak menurunkan akibat uang apa pun dari rujukan ini. Akibat uang hanya lahir dari tautan bertipe: `authorization_id` (§3.8), `installment_id`, `recurring_bill_id`, `set_aside_schedule_id`, `reconciliation_id`, `refund_of_entry_id`.

**Penanda `unconfirmed`.** Dua macam baris tempat uang membawanya (§3.3):

| | Baris | Sisi |
|---|---|---|
| (a) | `delta > 0` pada tempat uang yang pemegangnya anggota dewasa selain `actor` entri, apa pun jenis entrinya | Penerima |
| (b) | `delta < 0` pada tempat uang berpemilik `shared` yang pemegangnya anggota dewasa selain `actor`, **bila entrinya entri dua buku**, yaitu juga memuat baris berpemilik `member(M)`: penyisihan, penggantian, pinjaman dana bersama kepada M, pelunasan utang dana bersama kepada M | Pembayar |

Penanda ini fakta turunan dari `actor`, `holder`, dan buku yang disentuh entri saat entri disimpan, bukan kolom yang kelak diubah; status konfirmasinya sesudah itu dipegang dokumen 06 §7 sebagai peristiwa atas baris itu.

Enam dimensi custody dijawab dari:

| Dimensi (brief §5C) | Dijawab oleh |
|---|---|
| Siapa pemiliknya | `holding.owner` |
| Siapa pemegangnya | `holding.holder` |
| Di mana disimpan | `holding.kind`, `holding.name` |
| Untuk tujuan apa | baris `allocation` per tempat uang dan tujuan |
| Siapa yang berwenang membelanjakan | aturan rumah tangga (dokumen 06) dan *otorisasi* yang terbuka; bukan atribut uang |
| Siapa penerima manfaatnya | `beneficiary` pada baris sebab `expense` |

`actor` (yang mencatat), `payer` (yang membayar), dan `holder` (yang memegang) adalah tiga hal terpisah. Model tidak mengandaikan satu pasangan sebagai pencatat (keputusan pemilik nomor 1 dan 4).

**Jenis entri.** "Posisi" dan "sebab" di tabel ini adalah baris yang wajib ada; buku yang disentuh disebut di kolom terakhir.

| `type` | Baris posisi | Baris sebab | Buku |
|---|---|---|---|
| `opening` | posisi apa pun ± | `opening` | satu |
| `income` | tempat uang + | `income` + | satu |
| `income` dengan bagian pribadi (§3.1) | buku rumah tangga: tempat uang + sebesar jumlah dikurangi bagian pribadi (tidak ada bila seluruhnya ke dana pribadi); buku pribadi pencatat: tempat uang + sebesar bagian pribadi | buku rumah tangga: `income` + sebesar jumlah penuh dan `owner_transfer` − sebesar bagian pribadi; buku pribadi: `owner_transfer` +. Untuk penghasilan, **[Usulan]**: hanya bila `actor = earner` (INV-27) | dua |
| `expense` | tempat uang − | `expense` − (satu atau lebih; boleh disertai `expense` untuk biaya admin) | satu |
| `refund` | tempat uang + | `expense` +, tertaut ke entri asal | satu |
| `transfer` (termasuk serah terima) | tempat uang −, tempat uang + (pemilik sama); biaya admin opsional | tidak ada; `expense` − bila ada biaya admin | satu |
| `owner_transfer` | tempat uang − di buku asal, tempat uang + di buku tujuan | `owner_transfer` − dan + | dua |
| `contribution_expense` | tempat uang − di buku pribadi | buku pribadi: `owner_transfer` −; buku rumah tangga: `owner_transfer` + dan `expense` − | dua |
| `advance` (talangan dengan "minta diganti") | buku pembayar: tempat uang −, piutang +; buku yang ditalangi: kewajiban + | buku yang ditalangi: `expense` − | dua |
| `reimbursement` | buku yang ditalangi: tempat uang −, kewajiban −; buku pembayar: tempat uang +, piutang − | tidak ada | dua |
| `claim_waiver` | piutang − dan kewajiban − | `owner_transfer` − di buku yang melepas, + di buku yang dibebaskan | dua |
| `exit_closure` (hanya oleh sistem, saat anggota keluar; §8.2) **[Usulan]** | untuk setiap klaim antar-pemilik yang menyangkut buku pribadi anggota yang keluar: piutang − dan kewajiban − sebesar seluruh sisanya; tidak pernah baris tempat uang atau aset | per klaim: `owner_transfer` − di buku pemilik piutang, + di buku yang berkewajiban | dua atau lebih |
| `loan_given` | tempat uang −, piutang + | tidak ada | satu (dua bila antar-pemilik) |
| `loan_received` | tempat uang +, kewajiban + | tidak ada | satu (dua bila antar-pemilik) |
| `repayment_in` | tempat uang +, piutang − | tidak ada | satu (dua bila antar-pemilik) |
| `repayment_out` (termasuk pembayaran cicilan) | tempat uang −, kewajiban − | tidak ada; `expense` − untuk denda | satu (dua bila antar-pemilik) |
| `write_off` | piutang − | `write_off` − | satu |
| `debt_forgiven` | kewajiban − | `debt_forgiven` + | satu |
| `asset_purchase` | tempat uang − (uang muka atau lunas), aset +, kewajiban + bila dicicil | `credit_cost` − bila dicicil dan ada biaya kredit | satu |
| `asset_valuation` | aset ± | `valuation` ± | satu |
| `asset_disposal` | aset −, tempat uang + (bila dijual) | `valuation` ± sebesar selisih hasil dan nilai tercatat | satu |
| `cash_difference` | tempat uang ± | `cash_difference` ± | satu |
| `recategorization` | tidak ada | dua baris sebab yang saling meniadakan (§3.10) | satu |
| `allocation` | tidak ada | tidak ada; hanya baris `allocation` | satu |

Koreksi bukan jenis tersendiri: entri pembalik membawa `type` entri asal, `corrects`, dan `correction_kind = reversal`.

**Penulisan.** Satu perintah tulis menghasilkan tepat satu entri dalam satu transaksi basis data, membawa satu `idempotency_key` (`ENGINEERING-STANDARDS.md` §5). Perintah dengan kunci yang sudah pernah diterima mengembalikan entri yang sama dan tidak menulis apa pun (INV-10). Saldo posisi adalah jumlah `delta` barisnya; cache saldo, bila ada, diperbarui di transaksi yang sama (INV-04).

**Saldo negatif.** Pencatatan tidak pernah ditolak karena saldo tercatat tidak cukup: pencatat pelupa sering mencatat belanja sebelum mencatat uang masuknya. Saldo tercatat tempat uang yang negatif adalah tanda "ada yang belum dicatat" dan memicu ajakan rekonsiliasi kas; ia bukan pelanggaran invarian.

### 2.6 Entitas pendukung

Tidak satu pun entitas di bawah mengubah saldo. Semuanya dibuat, diubah, dan diarsipkan rumah tangga sendiri.

```text
category          id, name, parent_id?, archived_at
                  -- tidak ada kategori bawaan yang terkunci; contoh awal: belanja bulanan, cicilan, pulsa/wifi

purpose           id, name, kind (spending | savings), target_amount?, target_date?, archived_at
                  -- tujuan; tabungan adalah tujuan berjenis savings

budget_period_setting  household_id, start_day (1..28), effective_from, actor, recorded_at
                  -- hari mulai periode anggaran rumah tangga; bawaan 1; hanya ditambah (§6.1)

budget            id, category_id, limit_amount, period_start, period_end
                  -- satu baris per kategori per periode yang sudah dimulai; periode adalah rentang tanggal
budget_rule       category_id, limit_amount
                  -- pola batas bulanan per kategori; periodenya mengikuti budget_period_setting

authorization     id, request_id (dokumen 06), requester, limit_amount,
                  funding_owner (shared), category_id?, purpose_id?, source_holding_id?,
                  planned_date?, valid_until, terms_version,
                  status (open | settled | cancelled | expired), opened_at, closed_at
                  -- batas, kategori, tujuan, sumber, tanggal rencana, dan masa berlaku menyalin
                  -- versi syarat yang disetujui (dokumen 06); berubah hanya lewat authorization.amended (§3.8)

reservation       id, authorization_id, holding_id, amount, status (active | released)

installment       id, liability_id, asset_id?, name, counterparty,
                  amount_per_payment, payments_total, payments_made, due_day,
                  category_id, responsible_member
                  -- tanpa routine_agreement_ref: pembayaran cicilan tidak ditandai (dokumen 06 §6.2)

recurring_bill    id, name, counterparty, estimated_amount, due_rule,
                  category_id, responsible_member, end_date?, routine_agreement_ref?, archived_at

set_aside_schedule  id, recipient (member(M)), monthly_amount, due_rule?,
                  routine_agreement_ref?, archived_at
                  -- jatah pribadi M: satu jumlah per periode anggaran (§3.4); paling banyak satu
                  -- yang tidak terarsip per recipient; due_rule hanya untuk pengingat

reconciliation    id, holding_id, counted_amount, recorded_amount_at_count, counted_at, actor

external_party    id, nickname, relation      -- nama panggilan dan hubungan saja (Q-28)

entry_detail      entry_id, note?, receipt_refs?, counterparty_text?   -- dapat dihapus terpisah (§8)
```

**Rujukan ke kesepakatan rutin.** `routine_agreement_ref` hanya ada pada tagihan berulang dan jadwal penyisihan, dan menunjuk *kesepakatan rutin* yang berlaku atas jadwal itu, bila ada (dokumen 06 §6.3). **Cicilan tidak memilikinya**: pembayaran cicilan adalah pelunasan kewajiban yang sudah tercatat dan tidak pernah dibandingkan dengan ambang permintaan (§3.8, §3.14; dokumen 06 §6.2), sehingga rujukan itu tidak punya pemakai. Jumlah yang disepakati disimpan pada kesepakatan itu sendiri (dokumen 06 §6.3 "Kesepakatan rumah tangga"), bukan pada jadwal: per kali untuk tagihan berulang, per periode anggaran untuk *jatah pribadi*. `recurring_bill.estimated_amount` adalah perkiraan untuk pengingat, dan `set_aside_schedule.monthly_amount` adalah jumlah yang direncanakan atau diusulkan penerimanya; keduanya metadata yang dapat diubah tanpa persetujuan, dan mengubahnya tidak mengubah jumlah yang disepakati (menjawab dokumen 06 §15.4 butir (d)). Jumlah jatah yang berlaku selalu jumlah pada kesepakatan itu, bukan `monthly_amount` (§3.4). Rujukan itu diisi, diganti, dan dikosongkan hanya oleh alur kesepakatan di dokumen 06; ia metadata jadwal, bukan inti moneter, dan tidak mengubah angka apa pun (INV-12). Entri pembayaran tidak menunjuk kesepakatan secara langsung: ia ditautkan ke jadwalnya, dan dari jadwal itulah dokumen 06 membaca apakah pembayaran itu tercakup kesepakatan rutin. Kiriman keluarga bulanan memakai `recurring_bill` dengan pihak luar penerimanya sebagai `counterparty`.

**Pengarsipan.** Menghapus sesuatu yang punya riwayat berarti mengarsipkannya (keputusan pemilik nomor 9; Q-58): ia hilang dari pilihan saat mencatat, riwayat dan laporannya tetap. Syarat:

| Yang diarsipkan | Syarat |
|---|---|
| Tempat uang | Saldo tercatat nol dan tidak ada reservasi aktif. Bila belum nol, pemegang memindahkan isinya atau merekonsiliasinya lebih dulu |
| Piutang, kewajiban | Saldo nol (lunas, dihapus, atau dibebaskan) |
| Aset | Nilai tercatat nol lewat `asset_disposal` |
| Kategori, tujuan | Selalu boleh; anggaran dan alokasi yang masih berjalan atasnya berakhir di akhir periodenya |
| Cicilan | Kewajibannya nol |
| Tagihan berulang, jadwal penyisihan | Selalu boleh; kesepakatan rutin yang merujuknya diakhiri menurut dokumen 06 §6.3 |

Yang belum pernah dipakai entri mana pun boleh dihapus sungguhan. Ganti nama adalah perubahan metadata: tercatat sebagai *peristiwa* di jejak audit beserta pelakunya, tidak membuat entri, dan riwayat menampilkan nama terbaru.

---

## 3. Aturan per jenis peristiwa

### 3.1 Uang masuk

Keputusan pemilik nomor 1 dan 2. Layar "uang masuk" mencatat sumber, atas nama siapa, dan siapa pemegangnya. Yang menerima uang otomatis menjadi pemegangnya dan mencatatnya; pemegang ditentukan oleh tempat uang yang dipilih.

**Keputusan pemilik produk, 11 November 2026.** Yang mencatat uang masuk memilih ke mana uang itu pergi: masuk kas bersama, langsung dipakai untuk sesuatu, atau masuk dana pribadi; ini berlaku juga untuk gaji. Bila gaji diarahkan ke dana pribadi, pasangan melihat catatannya, misalnya "gaji Rp10 juta masuk, Rp1 juta ke pribadi": jumlah gajinya dan berapa yang ke pribadi. Yang tetap tertutup hanya saldo dana pribadi itu dan pemakaiannya sesudahnya. Dengan keputusan ini, usulan dokumen ini sebelumnya ("gaji selalu masuk dana bersama") **tidak** diambil sebagai larangan; yang diambil adalah keterlihatannya.

| Sumber di layar | Entri | Pemilik bawaan | Boleh dicatat hanya di buku pribadi? |
|---|---|---|---|
| Gaji, hasil usaha (*penghasilan*) | `income` | `shared`, walau tercatat "gaji suami" atau "gaji istri" lewat `earner` | **Tidak.** Pemasukannya selalu tercatat penuh di buku rumah tangga; bagian yang ke dana pribadi tercatat di entri yang sama (INV-24) |
| Hadiah untuk satu orang | `income` di buku pribadi orang itu | `member(M)` | Ya, dan itu bawaannya |
| Hadiah untuk keluarga | `income` | `shared` | — |
| Pinjaman yang diterima | `loan_received`; **bukan pemasukan** | `shared` | — |
| Lainnya | `income` | `shared` | Ya, bila memang uang pribadi (misalnya hasil menjual barang pribadi) |

Pemilik ditentukan oleh tempat uang tujuan (satu tempat uang, satu pemilik). Mengarahkan uang masuk berarti memilih tempat uang tujuannya.

**Tiga arah, satu aturan.**

| Arah yang dipilih pencatat | Bentuk di model |
|---|---|
| Masuk dana bersama | Entri `income`: tempat uang dana bersama + jumlah; sebab `income` + jumlah |
| Ke dana pribadi pencatat, sebagian atau seluruhnya (X dari jumlah J, 0 < X ≤ J). Untuk penghasilan, **[Usulan]**: hanya bila pencatat adalah `earner`-nya (INV-27) | **Satu** entri `income` dengan baris di dua buku. Buku rumah tangga: sebab `income` +J, sebab `owner_transfer` −X, tempat uang +(J − X). Buku pribadi pencatat: tempat uang +X, sebab `owner_transfer` +X. Keseimbangan buku rumah tangga: J − X = J − X; buku pribadi: X = X; jumlah `owner_transfer` nol (INV-03) |
| Langsung dipakai untuk sesuatu | Entri `income` ke tempat uang penerimanya, lalu entri biasa dari tempat uang itu (pengeluaran, pelunasan, pinjaman, dan seterusnya). Layar boleh menyajikannya sebagai satu alur; modelnya tetap dua entri dengan dua idempotency key **[Usulan]**, karena "sesuatu" itu dapat berupa jenis entri apa pun dan jenis entri adalah daftar tertutup |

Mengapa bagian pribadi adalah baris di entri yang sama, bukan pemasukan ke dana bersama disusul penyisihan sebagai entri kedua **[Usulan]**:

1. Satu perintah tulis menghasilkan tepat satu entri (§2.5). Pemasukan dan bagian pribadinya tersimpan bersama atau tidak sama sekali; tidak ada saat ketika gaji sudah tercatat tetapi arahnya belum.
2. Catatan mengikuti kenyataan. Uang yang tidak pernah berada di tempat uang dana bersama tidak dicatat seolah pernah di sana; tidak ada serah terima semu dan tidak ada saldo yang sesaat salah.
3. Yang dilihat pasangan adalah satu catatan, persis seperti contoh pemilik produk: jumlah penuh dan bagian pribadinya.

Bagian pribadi itu adalah *penyisihan* (§3.4) dan tunduk pada aturannya: terlihat oleh keduanya beserta jumlah, tanggal, pelaku, dan penerimanya; tampil di ringkasan sebagai "disisihkan ke dana pribadi" (§6.3). Ia terhitung terhadap *jatah pribadi* bulanan pencatatnya (§3.4; **[Usulan]**, dokumen 17 D-48).

**Beberapa uang masuk dalam satu periode; rekening bank atau dompet digital** (pernyataan pemilik produk, 11 November 2026: "gajian saya bisa 2x bahkan uang masuk dari freelancer juga bisa sampe 5x", masuk ke rekening bank atau e-wallet). "Sebulan" dan jumlah seluruhnya, sampai sekitar tujuh uang masuk, adalah **bacaan lead**, bukan kata-katanya; kalimatnya juga dapat dibaca "seluruhnya sampai lima kali" (dokumen 17, kelompok ketiga). Tidak ada aturan yang bergantung pada angkanya. Tidak ada aturan baru. Setiap uang masuk adalah satu entri `income` dengan idempotency key-nya sendiri, ke tempat uang yang menerimanya, apa pun jenisnya (`bank`, `ewallet`, `cash`); model tidak mengandaikan satu gajian per periode, dan tidak ada angka yang bergantung pada "tanggal gajian". Penghasilan lepas dicatat dengan sumber `business` **[Usulan]** (T-27; dokumen 17 S2-1 butir (iv)), sehingga ia *penghasilan* dan tunduk pada INV-24 dan INV-27. Tiap entri boleh memuat bagian pribadi; jumlahnya sepanjang periode dihitung terhadap satu jatah (§3.4). Meneruskan uang itu sesudahnya memakai aturan yang sudah ada: ke pasangan adalah serah terima, juga dari rekening ke rekening (§3.3); tarik tunai adalah pemindahan, dengan biaya admin di entri yang sama (§3.3); kepada ibu, adik, atau ayah adalah kiriman keluarga atau pinjaman, dipilih per entri (§3.2). Contoh terhitung satu periode dengan tiga uang masuk: CTH-39.

**Sumber mana yang wajib terlihat [Usulan].** Daftarnya bukan bagian tetap dari model: ia **parameter preset aturan** rumah tangga, *daftar sumber penghasilan* (INV-24; permintaan `security-privacy`, SP-18), supaya model tidak mengunci satu bentuk keluarga. Preset pilot ("dua dewasa setara"): `salary` dan `business`: keduanya nafkah rumah tangga, dan contoh pemilik produk menyebut gaji. `gift` untuk satu orang dan `other` tidak termasuk, lihat di bawah. Bila pemilik produk menghendaki lain (misalnya `other` ikut wajib), yang berubah hanya isi daftar itu (T-14). Perubahan daftar berlaku atas perintah yang diterima sesudahnya; entri lama tidak berubah (INV-09). Siapa yang boleh mengubahnya dan lewat alur apa adalah aturan dokumen 06 §6.3 "Kesepakatan rumah tangga" (§10.2).

**Hadiah untuk satu orang tidak terlihat sama sekali oleh pasangan [Usulan]; belum ditanyakan kepada pemilik produk dan menunggu jawabannya di checkpoint S2** (dokumen 17, bagian "Checkpoint S2: pertanyaan untuk pemilik produk", pertanyaan S2-1 butir (ii); Q-77). Yang sudah diputuskan pemilik produk hanyalah bahwa hadiah itu menjadi dana pribadi penerimanya; bahwa catatannya tidak terlihat adalah usulan dokumen ini. Bila ia memilih sebaliknya, `gift` masuk daftar sumber INV-24 dan dicatat dengan bentuk CTH-03b; tidak ada angka contoh yang berubah selain CTH-36. Menurut usulan ini, ia `income` di buku pribadi penerimanya dan tidak menyentuh buku rumah tangga. Alasan: keputusan pemilik nomor 2 menjadikannya dana pribadi secara bawaan, keputusan nomor 3 menutup dana pribadi, dan pemilik produk tidak meminta hadiah terlihat. Menampakkannya berarti membuka pemasukan dana pribadi beserta asalnya (misalnya pemberian dari orang tua sendiri), yaitu hal yang justru dijaga keputusan nomor 3 bagi kedua pasangan secara sama. Biayanya dinyatakan di §10.3 butir 1: sumber diisi pencatat sendiri.

**Siapa yang mencatat.** Karena hanya M yang menulis ke buku pribadi M (§2.2), arah "ke dana pribadi" hanya tersedia bagi pencatat untuk dana pribadinya sendiri.

**Bagian pribadi dari penghasilan hanya bagi penerimanya [Usulan]; menunggu pemilik produk di checkpoint S2** (dokumen 17 S2-1 butir (i); INV-27; temuan `security-privacy` SP-01; dokumen 06 §6.4 "Uang masuk yang diarahkan ke dana pribadi"). Aturan di atas saja belum cukup: ia hanya memeriksa bahwa buku pribadi yang ditulis milik pencatat, sehingga A dapat mencatat "gaji B 10000000, 3000000 ke pribadi A", dan B tidak dapat membatalkannya, karena entri itu entri dua buku (§3.12). Usulannya: entri `income` bersumber penghasilan hanya boleh memuat bagian pribadi bila **pencatatnya adalah penerima penghasilan itu** (`actor = earner`). Perintah yang melanggar ditolak seluruhnya; pemasukan yang sama tetap dapat dicatat tanpa bagian pribadi, jadi tidak ada penghasilan yang terhalang dicatat. Bila pasangan yang mencatat gaji itu, ia mencatatnya ke dana bersama; siapa pun yang menghendaki bagian pribadi mengambilnya lewat penyisihan biasa (§3.4), yang terlihat dan mengikuti ambang.

- *Bagaimana penerima penghasilan dikenali.* Model sudah memisahkan `earner` dari `actor`: `earner` adalah isian "punya siapa" pada layar uang masuk (Q-55), tersimpan pada baris sebab `income`. Untuk sumber penghasilan ia wajib diisi dan bernilai tepat satu anggota dewasa rumah tangga itu. Hasil usaha yang dikerjakan berdua dicatat atas nama salah satu; bagian pasangannya diambil lewat penyisihan.
- *Batas yang dinyatakan.* `earner` diisi pencatat, sama seperti sumber. Pencatat dapat menulis dirinya sebagai `earner` atas gaji pasangannya. Bedanya dari celah label sumber (§10.3 butir 1): isian itu tercatat di buku rumah tangga dan terlihat oleh keduanya ("gaji, atas nama A"), sehingga pasangannya dapat mengajukan keberatan atas hal yang ia lihat. Pagar ini mencegah pengalihan yang jujur labelnya dan membuat yang tidak jujur terlihat; ia tidak membuktikan siapa penerima gaji sebenarnya.
- *Syarat keadilan.* Aturannya per penerima penghasilan dan simetris: kedua pasangan di rumah tangga pilot berpenghasilan (pernyataan pemilik produk, 11 November 2026), masing-masing mencatat penghasilannya sendiri dan hanya dapat mengarahkan bagian dari penghasilannya sendiri. Atas penghasilan pasangannya, seorang anggota memperoleh dana pribadi hanya lewat penyisihan, yang ditandai bila di atas ambang. Pagar ini adil hanya bila bagian pribadi penerima penghasilan diperlakukan dengan aturan yang sama dengan penyisihan lain: terhitung terhadap jatah pribadi yang sama, dan bagian di luar jatah dibandingkan dengan ambang dengan cara yang sama (T-13 dijawab "ya"; dokumen 17 D-18 dan D-48, keduanya **[Usulan]**; perlakuan bagian di luar jatah: S2-1 butir (iii)). Pagar itu sendiri ditanyakan sebagai S2-1 butir (i). Bila pemilik produk memilih bagian itu tidak ditandai, aturan ambang atas dana pribadi makin tidak mengikat makin besar penghasilan seseorang, dan pada rumah tangga yang hanya satu anggotanya berpenghasilan ia hanya mengikat anggota yang lain (§10.3 butir 1).
- *Alternatifnya.* Tanpa INV-27: siapa pun yang mencatat penghasilan boleh mengarahkan sebagian atau seluruhnya ke dana pribadinya sendiri, juga atas gaji pasangannya; penerima gaji melihatnya dan dapat mengajukan keberatan, tetapi tidak dapat membatalkannya. Kata-kata pemilik produk ("yang mencatat memilih arahnya") dapat dibaca demikian; karena itu pagar ini usulan, bukan tafsir atas keputusannya. Pada kedua jawaban bentuk entri, saldo, dan semua angka contoh tidak berubah: pada CTH-03b pencatat dan `earner` sama-sama suami.

### 3.2 Pengeluaran

- Satu entri `expense` boleh berisi beberapa baris sebab (kategori atau penerima manfaat berbeda) dan beberapa baris tempat uang (dibayar dari dua tempat). Jumlah keduanya sama.
- Belanja biasa dicatat setelah kejadian tanpa persetujuan (keputusan pemilik nomor 5). Status permintaan tidak pernah menghalangi pencatatan (A-20).
- **Kiriman keluarga** adalah `expense` dengan pihak luar sebagai penerima manfaat; tidak dilacak setelah diserahkan (default Q-07 yang diterima pemilik produk). Kiriman bulanan dijadwalkan sebagai tagihan berulang (§3.15).
- Uang untuk saudara atau orang tua: pencatat memilih per entri antara `loan_given` dan kiriman keluarga.

### 3.3 Pemindahan dan serah terima

- `transfer` hanya antara dua tempat uang berpemilik sama. Total uang pemilik itu tidak berubah (INV-06).
- **Serah terima** adalah `transfer` yang kedua tempat uangnya berbeda pemegang. Saldo kedua tempat uang berubah seketika saat dicatat.
- **Konfirmasi custody, sisi uang** (Q-19; keputusan pemilik produk, 11 November 2026: serah terima tercatat seketika, bertanda "belum dikonfirmasi" sampai penerima menerimanya; penerima dapat mengajukan *keberatan* kapan saja; sama ke dua arah). Penanda `unconfirmed` melekat pada **setiap baris yang menambah saldo tempat uang berpemegang anggota dewasa selain pencatat**, bukan hanya pada serah terima: pemasukan, pelunasan piutang, atau pengembalian dana yang dicatat ke tempat uang pasangan, dan entri pembalik yang mengembalikan uang ke tempat uang pasangan (§2.5). Yang tidak membawa penanda: baris ke tempat uang yang dipegang pencatat sendiri, ke tempat uang bersama (`holder = joint`), dan ke titipan pada pihak luar. Tempat uang pribadi tidak pernah membawanya, karena hanya pemiliknya yang menulis ke sana (INV-25).
- **Sisi pembayar pada entri dua buku [Usulan]** (temuan `security-privacy` SP-05; dokumen 06 §7.2 "Aturan umum"). Penanda yang sama melekat pada **baris yang mengurangi tempat uang dana bersama berpemegang anggota dewasa selain pencatat, bila entrinya entri dua buku** (§2.5, baris (b)): M mencatat penyisihan, penggantian, pinjaman dari dana bersama, atau pelunasan utang dana bersama kepada dirinya, dari tempat uang dana bersama yang dipegang pasangannya. Alasannya: entri itu memindahkan uang dari tangan pemegang ke dana pribadi pencatat, dan pemegangnya tidak dapat membalikkannya (§3.12); tanpa penanda, uang yang ia pegang berkurang di catatan tanpa pernyataannya. Belanja biasa yang dicatat atas tempat uang pasangan (entri satu buku) **tidak** membawa penanda ini: pemegangnya dapat mengoreksinya sendiri dengan entri pembalik.
- Penanda itu tidak mengubah angka apa pun: saldo berubah seketika, dan custody menampilkan jumlahnya pada pemegang sebagai "belum dikonfirmasi" (§5), terpisah untuk baris masuk dan baris keluar. Tidak ada batas waktu dan tidak ada konfirmasi otomatis. Bila pemegang mengajukan *keberatan*, entri tetap ada dan tetap terhitung sampai dikoreksi dengan entri pembalik. Alur konfirmasi dan keberatan adalah milik dokumen 06 §7 dan §8; kaitnya `handover.recorded`, `handover.confirmed`, `handover.disputed`, yang berlaku atas semua baris bertanda itu. `handover.confirmed` hanya terpicu oleh pernyataan pemegangnya sendiri (dokumen 06 C02 dan C03). **Menutup keberatan tanpa entri pembalik (dokumen 06 C05) tidak memicu kait uang apa pun dan bukan konfirmasi**: baris itu kembali ke status sebelum keberatan, sehingga jumlahnya berpindah dari "dikeberatani" kembali ke "belum dikonfirmasi" atau ke terkonfirmasi di custody (§5.1), tanpa angka lain yang berubah.
- **Pinjam kembali** dana bersama adalah serah terima, bukan utang (default Q-01 yang diterima pemilik produk).
- **Biaya admin** (Q-48) dicatat **di entri yang sama**: tempat uang asal berkurang sebesar jumlah ditambah biaya, tempat uang tujuan bertambah sebesar jumlah, dan satu baris `expense` sebesar biaya (kategori bawaan "biaya admin"). Pencatat mengisinya di layar yang sama.

### 3.4 Penyisihan dan kontribusi (melintasi garis)

| Peristiwa | Entri | Buku rumah tangga | Buku pribadi | Terlihat pasangan |
|---|---|---|---|---|
| Penyisihan: dana bersama menjadi dana pribadi (misalnya uang saku) | `owner_transfer` | tempat uang −X; `owner_transfer` −X | tempat uang +X; `owner_transfer` +X | Ya: jumlah, tanggal, pelaku, penerima. Saldo pribadi sesudahnya tidak |
| Kontribusi uang: dana pribadi diserahkan menjadi dana bersama | `owner_transfer` | tempat uang +X; `owner_transfer` +X | tempat uang −X; `owner_transfer` −X | Ya |
| Kontribusi belanja: dana pribadi membayar keperluan rumah tangga, tidak minta diganti | `contribution_expense` | `owner_transfer` +X; `expense` −X dengan `funded_by = member(M)` | tempat uang −X; `owner_transfer` −X | Ya: pengeluarannya dan sumbernya ("beras, pakai uang pribadi suami"). Tempat uang pribadi yang dipakai dan saldonya tidak |

Kontribusi belanja tidak mengubah *uang yang ada* dan tidak mengubah kekayaan bersih rumah tangga: rumah tangga menerima nilai X dan langsung memakainya. Pengeluarannya terhitung **satu kali** di laporan kategori dan anggaran; kontribusinya tampil satu kali sebagai baris "kontribusi dari dana pribadi" (§6.3). Menjumlahkan keduanya sebagai dua pengeluaran adalah kesalahan yang dicegah INV-02.

Ketiga entri di atas dicatat oleh anggota yang dana pribadinya tersentuh (§2.2). Bagian pemasukan yang diarahkan langsung ke dana pribadi (§3.1) adalah penyisihan yang sama, hanya saja barisnya berada di dalam entri `income`.

Penyisihan mengikuti ambang permintaan (default Q-49 yang diterima pemilik produk). Jumlah yang dibandingkan dengan ambang adalah jumlah baris `owner_transfer` negatif di buku rumah tangga (§3.8), dikurangi bagian yang berada di dalam *jatah pribadi* penerimanya (di bawah; **[Usulan]**, D-48). **Bukan keputusan pemilik produk: apakah bagian pribadi di dalam entri `income` ikut dibandingkan dengan ambang.** Pemilik produk hanya memutuskan bahwa catatannya terlihat (§3.1); ia tidak menyebut ambang. **[Usulan]** dokumen ini dan dokumen 06 (T-13; dokumen 06 §6.4), berlaku sebagai default yang dibacakan di checkpoint S2 (dokumen 17 D-18 dan D-48; Q-78), dengan perlakuan bagian di luar jatah ditanyakan sebagai S2-1 butir (iii): bagian itu diperlakukan sama dengan penyisihan lain, supaya mengarahkan uang saat dicatat tidak menjadi jalan memutar ambang; tanda itu netral dan tidak menghalangi apa pun. Bila pemilik produk memilih sebaliknya, baris `owner_transfer` di dalam entri `income` dikecualikan dari perbandingan dengan ambang; bentuk entri, saldo, dan semua angka contoh tidak berubah.

**Jatah pribadi bulanan.** Yang diputuskan pemilik produk (11 November 2026, keputusan 14: "Jatah pribadi dijumlah per bulan bebas bisa diatur"; menjawab Q-98) hanyalah bentuknya: tiap anggota dewasa punya paling banyak satu *jatah pribadi*, satu jumlah per *periode anggaran* (§6.1), dijumlahkan sepanjang periode itu, dan besarnya dapat diatur. Bentuk ini menggantikan rumusan dokumen ini sebelumnya (jumlah per kali, dibandingkan per entri). Hubungannya dengan ambang **bukan** keputusannya:

- *Di dalam jatah tidak dibandingkan dengan ambang* **[Usulan]**. Penyisihan kepada M sampai sejumlah jatahnya dalam satu periode tidak dibandingkan dengan *ambang permintaan*, jadi tidak ditandai. Ini dasar pertanyaan lead (butir 5 usul lead), disampaikan kepada pemilik produk dan tidak ia bantah; ia bukan kata-katanya, dan berlaku sebagai default D-48 (dokumen 17). Menurut usulan ini, pengecualian ambang itulah satu-satunya akibat jatah.
- *Di luar jatah dibandingkan seperti penyisihan lain* **[Usulan]; terbuka di checkpoint S2** (dokumen 17 S2-1 butir (iii); Q-106). Pilihan (a), rancangan sekarang: bagian yang melampaui jatah dibandingkan dengan ambang, sehingga catatannya ditandai hanya bila bagian di luar jatah itu sendiri di atas ambang. Pilihan (b): selama ada jatah yang berlaku, catatan yang melampaui jatah ditandai berapa pun bagian di luarnya. Hitungan di bawah menyediakan keduanya: (a) membandingkan `luar_jatah(e)` dengan ambang; (b) menandai bila `luar_jatah(e) > 0`. Tanpa jatah yang berlaku, kedua pilihan membandingkan seluruh penyisihan dengan ambang (di bawah).

Pada setiap pilihan tidak ada yang dihalangi, dan setiap penyisihan tetap terlihat oleh keduanya, di dalam maupun di luar jatah. Jatah bukan uang dan bukan anggaran. Bila bagian pribadi di dalam entri `income` diputuskan tidak dibandingkan dengan ambang (T-13), jatah hanya bekerja atas entri `owner_transfer`.

**Jumlahnya disepakati berdua** (keputusan pemilik produk, 11 November 2026: "a di sepakati berdua, nanti saling konfirmasi aja"). Satu pihak mengusulkan jumlah jatah seorang anggota, pihak lain mengonfirmasi; perubahan berjalan dengan cara yang sama. Jumlah yang berlaku adalah jumlah pada *kesepakatan rutin* yang berlaku atas `set_aside_schedule` anggota itu, dibaca lewat `routine_agreement_ref` (§2.6); alur usul dan konfirmasinya milik dokumen 06 §6.3 "Kesepakatan rumah tangga". Sampai usulan baru dikonfirmasi, jumlah lama berlaku. Selama belum ada kesepakatan yang berlaku, atau sesudah kesepakatannya diakhiri, jatahnya 0 dan seluruh penyisihan dibandingkan dengan ambang. **Hanya pemilik jatah yang dapat mengakhiri kesepakatan jatahnya sendiri [Usulan]** (temuan `security-privacy` SP-24): pasangannya tidak dapat mengakhirinya sendirian; ia dapat mengusulkan jumlah yang lebih kecil, termasuk 0, dan jumlah lama tetap berlaku sampai usulan itu diterima pemilik jatah. Mengakhiri jatah sendiri hanya memperketat aturan atas diri pengakhirnya, jadi tidak memerlukan persetujuan. Dari sisi uang kedua jalan sama: jatah baru berlaku ke depan saja (di bawah), dan tidak ada angka saldo yang berubah (INV-12). Jumlah kedua pasangan boleh berbeda; aturannya sama. Seorang anggota dapat mengajukan kenaikan jatahnya dengan menyertakan alasan, dan kenaikan itu berlaku setelah pasangannya mengonfirmasi (pernyataan pemilik produk, 11 November 2026: "bisa ajukan kenaikan sesuai alasan"); alurnya milik dokumen 06 §6.3, dan usul, alasan, konfirmasi, maupun penolakannya tidak punya akibat uang apa pun (INV-12). Jatah yang baru berlaku ke depan saja: entri yang sudah tersimpan tidak dinilai ulang. Pilihan yang tidak diambil, dicatat satu kali: tiap orang menetapkan jatahnya sendiri dan pasangannya hanya melihat.

Yang disimpan (§2.6): `set_aside_schedule` per penerima, berisi `monthly_amount` (jumlah yang direncanakan atau sedang diusulkan), pengingat opsional, dan `routine_agreement_ref`. Ia bukan entri dan tidak mengubah angka apa pun (INV-12). Uang saku anggota dewasa dan bagian pribadi dari penghasilan bukan dua hal: keduanya cara mengambil jatah yang sama, yang satu lewat entri `owner_transfer`, yang lain lewat baris di dalam entri `income` (§3.1). Jatah boleh diambil dalam berapa pun bagian sepanjang periode.

**Hitungannya.** Untuk entri `e` dan anggota M:

```text
bagian(e, M)    = − Σ effect baris owner_transfer negatif di buku rumah tangga pada e
                  yang pasangannya di buku M, bila e berjenis owner_transfer atau income;
                  0 untuk jenis lain (claim_waiver dan exit_closure bukan penyisihan)
terhitung(e)    = bagian(e, M) > 0, e bukan entri pembalik, dan e tidak tertaut ke otorisasi
                  -- penyisihan yang disetujui lewat permintaan menjadi realisasi otorisasinya (§3.8)
jatah(e, M)     = jumlah jatah M yang berlaku saat e disimpan; 0 bila tidak ada
terpakai_sebelum(e, M)
                = Σ bagian(f, M) atas entri f yang terhitung, disimpan lebih dulu daripada e,
                  berada di periode anggaran yang sama dengan e, dan belum dibalik saat e disimpan
dalam_jatah(e)  = min(bagian(e, M), max(0, jatah(e, M) − terpakai_sebelum(e, M)))
luar_jatah(e)   = bagian(e, M) − dalam_jatah(e)       -- (a): inilah yang dibandingkan dengan ambang
                                                      -- (b): tanda bila > 0 dan ada jatah berlaku
                                                      -- [Usulan]; dokumen 17 S2-1 butir (iii)
```

Semua bilangan bulat; tidak ada pembagian. **Setiap penyisihan kepada M terhitung dengan sendirinya**, baik bagian pribadi di dalam entri `income` maupun entri `owner_transfer`: keanggotaannya diturunkan dari barisnya, bukan dari tautan. Pencatat tidak menambah langkah, dan tidak dapat memilih entri mana yang memakan jatah dengan menautkan atau tidak menautkannya. `links.set_aside_schedule_id` tetap sah untuk pengingat dan tidak mengubah hitungan. Yang tidak terhitung hanya dua: penyisihan yang tertaut ke otorisasi, dan `claim_waiver` (dicatat pasangannya, bukan pengambilan oleh M; §3.6). Kontribusi M dari dana pribadi tidak mengurangi hitungan: tidak ada angka bersih (§6.3). Rumusan ini sama dengan dokumen 06 §6.2 dan §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi": `terpakai_sebelum` adalah `jumlah_sebelum` di sana, dan `luar_jatah` adalah `lebih`, karena `bagian − min(bagian, max(0, jatah − sebelum)) = min(bagian, max(0, sebelum + bagian − jatah))`. Menjawab dokumen 06 §15.4 butir (o): *jumlah pemakaian* entri `owner_transfer` dan bagian pribadi entri `income` seluruhnya penyisihan kepada satu anggota; kedua jenis entri itu tidak memuat baris `expense` (tidak ada biaya admin di dalamnya, §2.5), sehingga `P − J = 0` bagi keduanya. Sebagai realisasi otorisasi, entri yang tertaut tetap terhitung penuh sebesar `bagian`-nya (§3.8). Contoh terhitung: CTH-39.

**"Periode yang sama" menurut tanggal apa.** Menurut **tanggal kejadian** entri. Ini ditetapkan dokumen 06 §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi", dan sama dengan dasar semua angka periode di dokumen ini. Tabel berikut mencatat alasannya dan dasar yang tidak dipilih.

| Dasar | Yang disediakan model | Akibat |
|---|---|---|
| Tanggal kejadian (`event_date`); **dipilih** | Periode `e` = periode anggaran yang memuat `event_date`-nya; urutan di dalam periode menurut waktu simpan | Sama dengan dasar anggaran dan laporan: jumlah `bagian` satu periode = baris "disisihkan ke dana pribadi M" pada ringkasan periode itu, dikurangi yang tertaut otorisasi. Entri terlambat memakai jatah periode kejadiannya, bukan periode berjalan. Celahnya: tanggal kejadian diisi pencatat, sehingga penyisihan dapat diberi tanggal periode lampau yang jatahnya belum terpakai; entri itu tampil sebagai entri terlambat dengan kedua tanggalnya |
| Waktu simpan (`recorded_at`); tidak dipilih | Periode `e` = periode anggaran yang memuat tanggal `recorded_at`-nya | Tidak dapat dimundurkan. Tetapi penghasilan akhir periode yang baru dicatat sesudah periode berganti memakan jatah periode baru, dan jumlah yang dipakai aturan jatah tidak lagi sama dengan baris ringkasan periode mana pun |

**Dinilai satu kali.** `dalam_jatah` dan `luar_jatah` dihitung saat `e` disimpan, dari entri yang sudah diterima server saat itu; hasilnya disimpan dokumen 06 sebagai catatan penilaian (dokumen 06 INV-A37). Entri yang datang kemudian, termasuk entri terlambat, entri pembalik, dan perubahan jatah, tidak mengubah pembagian entri yang sudah ada. **Entri pembalik melonggarkan jatah ke depan saja**: sejak pembalik tersimpan, entri asalnya tidak lagi masuk `terpakai_sebelum` entri berikutnya. Satu-satunya penilaian ulang adalah atas entri itu sendiri, bila tautan susulannya ke otorisasi ditambah atau dicabut (dokumen 06 §6.3).

**Jalan menghindar yang diterima: tanggal mundur.** Penyisihan hari ini dapat diberi tanggal kejadian periode lampau yang jatahnya belum terpakai. Model tidak dapat memeriksa tanggal kejadian, dan batas "paling jauh mundur" akan menghalangi entri terlambat yang jujur (§3.11). Yang melawannya hanya keterlihatan, dan untuk itu model menjamin dua hal (menjawab dokumen 06 §15.4 butir (s); disempurnakan menurut temuan `security-privacy` SP-28): setiap entri **menyimpan** `recorded_at` yang persis, diisi server dan tidak dapat diatur pencatat (§2.5, INV-22); dan setiap entri yang tanggal pencatatannya berbeda dari tanggal kejadiannya **menampilkan** tanggal pencatatan itu, hanya tanggalnya, sama bagi kedua pasangan (§3.11). Jam pencatatan tidak ditampilkan kepada siapa pun.

**Cocok dengan ringkasan** (menjawab dokumen 06 §15.4 butir (r)). Untuk satu periode dan satu anggota, baris "disisihkan ke dana pribadi" tabel A (§6.3) = jumlah `bagian` entri yang terhitung terhadap jatah dan tidak dibalik + jumlah penyisihan yang tertaut ke otorisasi. Dasarnya sama: tanggal kejadian; entri pembalik bertanggal kejadian sama dengan entri asalnya dan mengurangi baris yang sama (§6.3).

**Batas yang dinyatakan.** (1) Jatah tidak berpindah periode: yang tidak terpakai hangus, dan kelebihan satu periode tidak mengurangi periode berikutnya. (2) Tidak ada angka jatah atau pemakaiannya yang lintas periode, kumulatif, atau bersih per anggota (§6.3; SP-13); yang tampil hanya periode yang sedang dilihat, dan angkanya tidak membuka hal baru: ia jumlah penyisihan periode itu, yang memang sudah terlihat (§5.3, B1). (3) Jatah tidak punya akibat pada saldo, anggaran, *terikat*, reservasi, atau kekayaan bersih (INV-12, INV-30). (4) Periode peralihan saat hari mulai periode diganti (§6.1) membawa jatah penuh, tidak diprorata.

### 3.5 Talangan, hak ganti, dan penggantian

*Talangan* kini berarti: pembayaran dari dana pribadi untuk keperluan rumah tangga atau untuk anggota lain **yang ditandai "minta diganti"**. Tanpa tanda itu ia kontribusi (§3.4).

| Langkah | Buku pembayar | Buku yang ditalangi |
|---|---|---|
| Talangan X (`advance`) | tempat uang −X; piutang (hak ganti) +X | kewajiban (hak ganti) +X; `expense` −X |
| Penggantian Y ≤ sisa (`reimbursement`) | tempat uang +Y; piutang −Y | tempat uang −Y; kewajiban −Y |

Talangan dicatat oleh yang menalangi; penggantian dari dana bersama dicatat oleh **penerimanya**; bila tempat uang dana bersama yang berkurang dipegang anggota dewasa lain, baris itu bertanda `unconfirmed` sampai pemegangnya mengonfirmasi (§2.2, §3.3; dokumen 06 §7.2 "Aturan umum"). Talangan tidak menyentuh tempat uang dana bersama dan tidak membawa penanda; kewajiban yang dilahirkannya dapat dikeberatani dan tampil demikian (§3.12). Pengeluaran tercatat satu kali, saat talangan. Penggantian bukan pengeluaran dan bukan pemasukan. Hak ganti tercatat sejak talangan dicatat (Q-21), sebagai kewajiban rumah tangga kepada anggota itu, sehingga langsung mengurangi kekayaan bersih rumah tangga.

Bawaan tanda "minta diganti" **[Usulan]**: mati untuk belanja biasa (keputusan pemilik nomor 3); hidup bila entri ditautkan ke otorisasi atas dana bersama (§3.8, situasi C), karena persetujuannya menyatakan dana bersama yang menanggung. Pencatat dapat mengubahnya.

### 3.6 Pelepasan hak ganti dan klaim antar-pemilik lain

Melepas klaim berarti pemiliknya menghadiahkan jumlah itu kepada pihak yang berutang. Entri `claim_waiver`: piutang −X dan `owner_transfer` −X di buku yang melepas; kewajiban −X dan `owner_transfer` +X di buku yang dibebaskan.

Bila yang dibebaskan adalah rumah tangga, kekayaan bersih rumah tangga naik X dan laporan menampilkannya sebagai "kontribusi dari dana pribadi" (menjawab temuan N-14: kenaikan itu punya baris bernama, bukan muncul tanpa sebab). Hasil akhirnya sama persis dengan kontribusi belanja sejak awal (CTH-16).

Hanya pemilik piutang yang dapat melepasnya. Pihak yang berutang tidak dapat menghapus utangnya sendiri. Untuk klaim yang pemilik piutangnya dana bersama (dana bersama meminjamkan kepada M), yang melepas adalah anggota dewasa selain M; inilah satu-satunya entri yang dicatat anggota lain dan menulis ke buku pribadi M (§2.2; di luar itu hanya entri sistem `exit_closure`, §8.2), dan sebagai pemakaian dana bersama ia mengikuti ambang permintaan sebesar jumlah yang dilepas: *jumlah pemakaian*-nya adalah baris `owner_transfer` negatif di buku rumah tangga (§3.8). Di laporan ia tampil sebagai baris tersendiri, "utang dihapuskan", **bukan** sebagai "disisihkan ke dana pribadi M" (§6.3; temuan `security-privacy` SP-13): tidak ada uang yang berpindah ke M pada saat itu, dan menggabungkannya akan membesarkan angka penyisihan M. Arah sebaliknya, M melepas hak gantinya atas dana bersama, bukan pemakaian dana bersama dan tidak dibandingkan dengan ambang.

### 3.7 Utang, piutang, pelunasan, penghapusan

- `loan_given` kepada pihak luar: uang −X, piutang +X; kekayaan bersih tetap.
- `repayment_in`: uang +Y, piutang −Y, dengan Y ≤ sisa piutang. Kelebihan bayar dicatat sebagai `income` bersumber `other`.
- `write_off` (Q-47): piutang −Z, sebab `write_off` −Z. Dilaporkan sebagai baris tersendiri, "piutang diikhlaskan", pada bulan tanggal kejadian penghapusannya; tidak dihitung terhadap anggaran mana pun dan tidak muncul di arus kas, karena uangnya sudah keluar saat dipinjamkan. Tidak digabung dengan kiriman keluarga, supaya "berapa yang kita kirim" dan "berapa yang tidak kembali" tetap dua jawaban.
- `loan_received` dan `repayment_out` adalah cerminnya di sisi kewajiban.
- **Utang antar-pasangan** hanya lahir bila uang berpindah antara dua pemilik berbeda dengan janji kembali: `loan_given` di buku pemberi dan `loan_received` di buku penerima, dalam satu entri. Kedua pihak selalu melihatnya (Q-24). Bila satu sisinya dana bersama, entrinya dicatat anggota yang dana pribadinya tersentuh; bila kedua sisinya dana pribadi, alur pencatatannya ditunda dari MVP (§2.2, T-12).

### 3.8 Otorisasi, reservasi, penyelesaian

Dokumen 06 memiliki state machine permintaan. Bagian ini hanya menetapkan akibat uang tiap peristiwanya.

| Peristiwa (kait untuk dokumen 06) | Akibat uang |
|---|---|
| `request.submitted`, `request.countered`, `request.rejected` | Tidak ada |
| `request.approved` | Lahir `authorization` berstatus `open`. Tidak ada entri, tidak ada saldo yang berubah (INV-12). Bila permintaan menyebut tempat uang sumber, lahir satu `reservation` aktif sebesar jumlah yang disetujui |
| `handover` untuk permintaan itu | Entri `transfer`. Reservasi dilepas sebesar yang diserahkan |
| Belanja ditautkan ke otorisasi | Entri biasa (`expense`, `loan_given`, `asset_purchase`, `owner_transfer`, `advance`, …). Reservasi yang masih aktif dilepas sebesar belanja itu, paling banyak sampai habis |
| `authorization.amended` (perubahan syarat diterima; dokumen 06 B12) | Tidak ada entri dan tidak ada saldo yang berubah (INV-12). Batas, kategori, tujuan, tempat uang sumber, tanggal rencana, dan masa berlaku otorisasi menjadi nilai versi syarat yang baru. *Terikat* dihitung ulang dari batas baru dan mengikuti periode tanggal rencana yang baru (§6.1). Reservasi: bila tempat uang sumbernya tetap, reservasi aktif berubah sebesar perubahan batas, paling rendah sampai nol; bila sumbernya berganti, reservasi lama dilepas dan lahir reservasi sebesar *terikat* yang baru pada sumber baru; bila sumbernya dihapus, reservasi dilepas. Bila batas baru lebih kecil daripada realisasi, selisihnya adalah kelebihan dan `authorization.overrun` terpicu; bila batas baru tidak lebih kecil daripada realisasi, kelebihan yang pernah tampil hilang. Entri yang sudah tertaut tidak berubah |
| `authorization.settled` | Otorisasi ditutup; *terikat* menjadi nol; reservasi yang tersisa dilepas. Penyelesaian sendiri tidak membuat entri; entri hanya ada bila sisa dikembalikan (situasi B) |
| `authorization.cancelled`, `authorization.expired` | Sama dengan `settled` tanpa belanja: terikat nol, reservasi dilepas. Belanja yang sudah tertaut tetap tercatat |
| `authorization.overrun` (realisasi melampaui batas) | Tidak ada entri tambahan. Lihat "kelebihan" di bawah |

**Reservasi** (Q-12). Persetujuan **tidak** membuat reservasi, kecuali permintaannya menyebut tempat uang sumber; itulah "aturan eksplisit" yang diminta brief §7. Reservasi mengurangi *saldo tersedia* tempat uang itu, tidak pernah saldo tercatat. Ia tidak menghalangi pencatatan apa pun (A-20): bila pemegang membelanjakan uang yang direservasi untuk hal lain, saldo tersedia menjadi negatif dan tampil sebagai "kurang Rp… untuk yang sudah dijanjikan".

**Realisasi** otorisasi = jumlah *jumlah pemakaian* entri yang ditautkan kepadanya, apa pun jenisnya (pengeluaran, pinjaman yang diberikan, pembelian aset, penyisihan, pembayaran cicilan), dikurangi pengembalian dana yang tertaut pada entri-entri itu. **Terikat** = max(0, batas − realisasi) selama otorisasi `open`; nol sesudahnya.

**Jumlah pemakaian** satu entri adalah satu angka yang dipakai untuk dua hal: sebagai realisasi bila entri itu tertaut ke otorisasi, dan sebagai jumlah yang dibandingkan dengan *ambang permintaan* bila tidak (aturan tandanya milik dokumen 06 §6.2).

| Jenis entri | Jumlah pemakaian | Dibandingkan dengan ambang? |
|---|---|---|
| Baris `expense` berpemilik `shared` yang ditanggung dana bersama, **di entri jenis apa pun**: entri `expense`; baris denda di dalam `repayment_out`; baris biaya admin di dalam `transfer`; baris `expense` di dalam `recategorization` yang merinci selisih kas (§3.10). Tidak termasuk baris `expense` pada `contribution_expense` dan `advance`, yang punya barisnya sendiri di bawah | Jumlah bersih baris itu (`−Σ effect`), bila positif. Bagian pelunasan dan bagian pemindahan dari entri yang sama tidak termasuk | Ya |
| `loan_given` dari dana bersama | Jumlah pinjaman | Ya |
| `owner_transfer` keluar dari buku rumah tangga (penyisihan) | Jumlah baris `owner_transfer` negatif di buku rumah tangga | Ya, sebesar bagian yang di luar *jatah pribadi* penerimanya (`luar_jatah`, §3.4); bagian di dalam jatah tidak dibandingkan. Keduanya **[Usulan]** (dokumen 17 D-48; bagian di luar jatah: S2-1 butir (iii)) |
| `income` dengan bagian pribadi (§3.1) | Jumlah baris `owner_transfer` negatif di buku rumah tangga, yaitu bagian pribadinya | **[Usulan]** ya, seperti penyisihan lain, yaitu sebesar bagian yang di luar jatah pribadi; default yang dibacakan di checkpoint S2 (T-13; dokumen 17 D-18, D-48, S2-1 butir (iii); §3.4) |
| `claim_waiver` atas klaim dana bersama kepada seorang anggota (§3.6) | Jumlah yang dilepas, yaitu baris `owner_transfer` negatif di buku rumah tangga | Ya |
| `claim_waiver` oleh anggota atas hak gantinya kepada dana bersama (§3.6) | Nol: baris `owner_transfer` di buku rumah tangga positif; nilai masuk, bukan keluar | Tidak: bukan pemakaian dana bersama |
| `advance` atas keperluan rumah tangga | Jumlah hak ganti yang lahir | Ya |
| `asset_purchase` tunai | Jumlah yang dibayar | Ya |
| `asset_purchase` secara cicilan | **Uang muka ditambah jumlah seluruh cicilan** (= harga ditambah biaya kredit), yaitu seluruh dana bersama yang dipakai dan diikat entri itu | Ya |
| `contribution_expense` | Jumlah baris `expense` | Tidak: bukan dana bersama |
| `repayment_out` (cicilan, utang): bagian yang mengurangi kewajiban | Jumlah pelunasan itu; menjadi realisasi bila ditautkan ke otorisasi | **Tidak.** Pelunasan kewajiban yang sudah tercatat tidak pernah dibandingkan dengan ambang dan tidak memerlukan kesepakatan rutin (dokumen 06 §6.2, §6.3). Denda di entri yang sama adalah baris `expense` dan mengikuti baris pertama tabel ini |
| `opening` atas kewajiban dana bersama kepada pihak luar | Jumlah saldo awal kewajiban itu, satu kali, pada entri `opening`-nya | Ya, terhadap ambang yang berlaku saat entri disimpan (dokumen 06 §6.2, putusan 2). Saldo awal yang diisi saat onboarding, sebelum ada ambang yang disepakati, tidak pernah ditandai, juga tidak surut sesudah ambang berlaku. Pembayaran kewajiban itu sesudahnya adalah pelunasan dan tidak dibandingkan |
| `opening` atas tempat uang, piutang, aset, dan kewajiban lain | Nol | Tidak |
| `exit_closure` (§8.2) | Nol: entri sistem, bukan pemakaian oleh siapa pun | Tidak; tidak ditandai dan tidak dapat ditautkan ke otorisasi |
| Bagian pemindahan pada `transfer`; `reimbursement`, `allocation`, `refund`, dan lainnya | Nol; tidak ditautkan sebagai pemakaian | Tidak |

**Entri yang memuat lebih dari satu bagian.** Pelunasan dengan denda dan pemindahan dengan biaya admin memuat bagian dari dua baris tabel. Sebagai realisasi, *jumlah pemakaian* entri itu adalah jumlah semua bagiannya (pelunasan ditambah denda; biaya admin saja pada pemindahan, juga bila pemindahan itu serah terima untuk suatu permintaan). Yang dibandingkan dengan ambang hanyalah bagian yang kolom ketiganya "ya". Setiap baris entri terhitung paling banyak satu kali: pada `advance`, baris `expense` dan hak ganti yang lahir adalah jumlah yang sama dan dihitung sebagai satu angka. Pengembalian dana dan entri pembalik tidak punya jumlah pemakaian sendiri; keduanya mengurangi realisasi entri asalnya (§3.12, §3.13).

Untuk pembelian secara cicilan, yang dibandingkan dengan ambang bukan harga barang (terlalu kecil sebesar biaya kredit) dan bukan uang muka saja (pembelian besar akan lolos dengan uang muka kecil). Pembayaran cicilan atas pembelian yang sudah tertaut penuh ke suatu otorisasi tidak ditautkan lagi ke otorisasi itu, supaya tidak terhitung dua kali (CTH-34). Hal yang sama berlaku atas kewajiban yang masuk sebagai saldo awal dan entri `opening`-nya tertaut ke suatu otorisasi: pembayarannya tidak ditautkan lagi ke otorisasi itu. Pada kedua kasus ini aturan server, bukan aturan layar saja: perintah tautannya ditolak dan entri pembayarannya tetap tersimpan (INV-26; dokumen 06 §4.5 "Masa berlaku dan kedaluwarsa").

**Dua batas baris saldo awal kewajiban, dinyatakan terbuka.** (1) Model tidak membedakan utang lama yang baru teringat dari utang baru; keduanya dibandingkan dengan cara yang sama, dan tandanya hanya hilang lewat tautan ke otorisasi atau entri pembalik (dokumen 06 §6.2). (2) Utang uang yang dicatat sebagai `loan_received` tidak punya jumlah pemakaian dan tidak dibandingkan dengan ambang, sedangkan utang yang sama yang dicatat sebagai saldo awal kewajiban dibandingkan. Ketimpangan ini tidak membuka jalan menghindar: jalan yang tidak ditandai adalah jalan yang mencatat uangnya masuk ke tempat uang dana bersama, terlihat oleh keduanya. Apakah `loan_received` perlu ikut dibandingkan adalah aturan dokumen 06 (§10.2).

**Penautan entri ke otorisasi.**

1. **Saat dicatat.** Entri boleh ditautkan ke otorisasi yang `open`. Entri terlambat boleh pula ditautkan ke otorisasi yang sudah `settled` atau `expired`, asalkan tanggal kejadiannya tidak lebih awal dari tanggal persetujuan dan tidak lebih baru dari tanggal penutupan otorisasi itu. Otorisasi yang `cancelled` tidak menerima tautan: pemohonnya sendiri yang menariknya; entrinya tetap tersimpan tanpa tautan (pencatatan tidak pernah ditolak).
2. **Akibat tautan ke otorisasi yang sudah tertutup.** Status otorisasi tidak berubah dan tidak ada entri tambahan. Realisasi dihitung ulang (INV-14 tidak bergantung pada status); *terikat* tetap nol; tidak ada reservasi yang hidup kembali. Sisa atau kelebihan pada permintaan dihitung ulang, dan `authorization.overrun` terpicu bila realisasi untuk pertama kali melampaui batas. Anggaran menghitung entri itu pada periode tanggal kejadiannya, seperti entri terlambat lain (§3.11).
3. **Tautan susulan.** Tautan boleh ditambahkan pada entri yang sudah tersimpan **tanpa mengubah entri itu**: ia catatan `entry_link` tersendiri yang hanya ditambah (§2.5), dengan pelaku, waktu, dan idempotency key-nya sendiri. Syarat tanggalnya sama dengan butir 1. Tautan susulan yang keliru dicabut dengan `entry_link` lain yang menunjuknya (`revokes`); tautan yang dibawa entri sejak dicatat hanya dapat dibatalkan dengan entri pembalik. Tautan yang berlaku atas suatu entri = tautan pada entri itu ditambah tautan susulan yang tidak dicabut; satu entri tertaut ke paling banyak satu otorisasi (INV-26). Menambah atau mencabut tautan susulan menghitung ulang realisasi, *terikat*, dan reservasi persis seperti bila entri itu dicatat dengan tautan tersebut, dan tidak mengubah saldo apa pun. **Siapa yang boleh** ditetapkan dokumen 06 §4.5 dan INV-A34, dan dikutip di sini tanpa diubah: tautan susulan **ke otorisasi** hanya diterima dari pencatat entri itu atau pemohon otorisasi itu; pencabutan tautan susulan hanya dari pembuat tautan itu atau pencatat entrinya; tautan susulan ke jadwal (tagihan berulang, jadwal penyisihan) boleh ditambahkan anggota dewasa mana pun. Perintah dari pelaku yang tidak berwenang ditolak tanpa akibat apa pun pada tautan, realisasi, *terikat*, atau reservasi. Pelakunya selalu terlihat. Dari sisi uang pembatasan ini tidak mengubah hitungan apa pun; ia hanya mempersempit siapa yang dapat menggerakkan realisasi suatu otorisasi lewat tautan.

**Sisa = batas − realisasi, bila positif.** Arti "diselesaikan dengan benar" (brief §5D langkah 7) per situasi custody:

| Situasi | Keadaan | Sisa | Entri saat penyelesaian | Contoh |
|---|---|---|---|---|
| A | Pemohon membelanjakan dana bersama yang sudah ia pegang | Tidak pernah berpindah; tetap di tempat uang pemohon sebagai dana bersama | Tidak ada | CTH-21 |
| B | Dana bersama diserahkan kepada pemohon untuk permintaan itu | Berada di tempat uang pemohon, tetap dana bersama, dan terlihat di custody. Pemohon memilih: (1) mengembalikannya, atau (2) tetap memegangnya | (1) satu `transfer` sebesar sisa; (2) tidak ada | CTH-22 |
| C | Pemohon membayar dengan dana pribadinya atas otorisasi dana bersama | Tidak ada uang sisa. Hak ganti lahir sebesar **realisasi**, bukan sebesar batas | Tidak ada; penggantian adalah entri tersendiri | CTH-23 |
| D | Pemegang lain membayar langsung dari tempat uang sumber yang direservasi | Tetap di tempat uang sumber; reservasinya dilepas | Tidak ada | §7.5, varian CTH-22 |

Pada semua situasi: sisa tidak pernah menjadi pemasukan, tidak pernah menjadi milik pribadi pemohon, dan tidak mengubah *uang yang ada*. Menjadikan sisa sebagai dana pribadi adalah penyisihan (§3.4) dan mengikuti aturannya.

**Kelebihan = realisasi − batas, bila positif** (Q-13). Pencatatan tidak dihalangi. Seluruh belanja, termasuk kelebihannya, tercatat sebagai entri biasa dan terhitung penuh terhadap anggaran kategorinya. Kelebihan tampil pada permintaan sebagai "melebihi Rp…".

| Kelebihan dibayar dari | Perlakuan |
|---|---|
| Dana bersama | Pengeluaran dana bersama biasa; tidak ada klaim. Dokumen 06 memutuskan apakah ia memerlukan tanggapan pasangan; uangnya sudah tercatat benar apa pun jawabannya |
| Dana pribadi pemohon | Kontribusi (§3.4) secara bawaan; talangan bila pemohon menandai "minta diganti" untuk bagian itu |

### 3.9 Alokasi dan tabungan

- Alokasi adalah baris `allocation` (tempat uang, tujuan, ±jumlah). Jumlah alokasi per (tempat uang, tujuan) tidak pernah negatif.
- **Batas** (Q-45): saat perintah alokasi dicatat, jumlah seluruh alokasi pada satu tempat uang tidak boleh melebihi saldo tercatatnya; perintah yang melampaui ditolak.
- Entri yang mengurangi tempat uang boleh membawa tujuan; alokasi tujuan itu ikut berkurang sebesar jumlahnya, paling banyak sampai nol. Entri tanpa tujuan mengurangi bagian *belum dialokasikan*.
- **Belum dialokasikan** = saldo tercatat − jumlah alokasi. Bila belanja membuatnya negatif, tidak ada yang ditolak dan tidak ada alokasi yang dikurangi otomatis: tempat uang itu tampil "alokasi melebihi uang Rp…" sampai rumah tangga mengurangi salah satu alokasinya. Aplikasi tidak menebak tujuan mana yang dikorbankan.
- Pemindahan boleh membawa tujuan: alokasi berkurang di asal dan bertambah di tujuan dalam entri yang sama.
- Anggaran tidak memerlukan alokasi (keputusan pemilik nomor 8: tanpa amplop). Alokasi bersifat pilihan.
- **Tabungan** adalah tujuan berjenis `savings`. Menabung adalah alokasi, atau `transfer` ke tempat uang khusus yang membawa tujuan itu; tidak pernah pengeluaran. Kemajuan = jumlah alokasi tujuan itu di semua tempat uang; *target* tidak masuk angka mana pun.

### 3.10 Rekonsiliasi kas, selisih kas, dan entri terlambat

Keputusan pemilik nomor 7: catatan tidak pernah disunting; beda antara terhitung dan tercatat menjadi satu entri *selisih kas* yang terlihat; entri terlambat bertanggal kejadian asli kelak menguranginya; siapa yang lupa diselesaikan di luar aplikasi.

1. Pemegang mengisi saldo terhitung. `selisih = terhitung − tercatat`.
2. Bila selisih bukan nol, dicatat entri `cash_difference`: tempat uang ±selisih, sebab `cash_difference` ±selisih, tanggal kejadian = tanggal menghitung. Sesudahnya saldo tercatat = saldo terhitung.
3. **Selisih negatif** (uang keluar yang tidak tercatat) menurunkan kekayaan bersih dan terhitung dalam total uang keluar bulan itu sebagai "belum dirinci".
4. **Selisih positif** (uang masuk yang tidak tercatat, atau pengeluaran tercatat terlalu besar) menaikkan kekayaan bersih dan tampil sebagai baris tersendiri di sisi uang masuk; ia **bukan** pemasukan dan tidak mengurangi pengeluaran kategori mana pun.
5. **Merinci selisih** (`recategorization`): dua baris sebab tanpa baris posisi. Untuk selisih negatif: `cash_difference` +X (mengurangi selisih yang belum dirinci) dan `expense` −X dengan kategori dan tanggal kejadian aslinya. Saldo tidak berubah, karena uangnya sudah keluar dari catatan lewat selisih itu. Untuk selisih positif cerminnya: `cash_difference` −X dan `income` +X, atau `expense` +X bila ternyata suatu pengeluaran tercatat terlalu besar.
6. **Entri terlambat yang menutup selisih.** Saat pencatat memasukkan pengeluaran pada suatu tempat uang dengan tanggal kejadian tidak lebih baru daripada rekonsiliasi terakhirnya, dan rekonsiliasi itu masih punya selisih negatif yang belum dirinci, aplikasi menawarkan "ini bagian dari selisih hitung Rp…?" dengan bawaan ya. Bila ya, entrinya adalah `recategorization` (butir 5) sampai sebesar selisih yang belum dirinci; bagian yang melebihinya dicatat sebagai `expense` biasa yang mengurangi tempat uang. Bila tidak, seluruhnya `expense` biasa.
7. Selisih yang belum dirinci = |selisih| − jumlah yang sudah dirinci; tidak pernah negatif (INV-16).
8. Selisih kas dihitung terhadap anggaran hanya setelah dirinci ke kategori, pada tanggal kejadian yang diisi. Sebelum itu, anggaran kategori yang lazim memakai tempat uang itu menampilkan pemberitahuan "ada selisih yang belum dirinci Rp…".

Rekonsiliasi berlaku untuk semua jenis tempat uang, bukan hanya uang tunai.

### 3.11 Entri terlambat dan periode

- Entri terlambat boleh sejauh apa pun ke belakang, tetapi tidak lebih awal dari saldo awal tempat uang yang disentuhnya.
- Tidak ada penutupan periode. Laporan periode lampau dihitung dari tanggal kejadian, sehingga ikut berubah, dan menampilkan "berubah sejak terakhir dilihat" beserta daftar entri yang `recorded_at`-nya lebih baru daripada akhir periode itu.
- Tanggal kejadian di masa depan ditolak; rencana ke depan adalah tagihan berulang, cicilan, atau permintaan.
- **Tanggal pencatatan tampil bila berbeda dari tanggal kejadian** (temuan `security-privacy` SP-28). Setiap entri menyimpan `recorded_at` yang persis dari server (§2.5, INV-22). Yang **ditampilkan** hanyalah tanggal pencatatan (tanggal dari `recorded_at`, tanpa jam), dan hanya pada entri yang tanggal pencatatannya berbeda dari tanggal kejadiannya; entri itu tampil sebagai entri terlambat dengan kedua tanggalnya, sama bagi kedua pasangan. Entri yang dicatat pada hari kejadiannya hanya menampilkan tanggal kejadian. `recorded_at` yang persis tidak ditampilkan kepada siapa pun, pada entri maupun riwayatnya, karena jam pencatatan yang terlihat menjadi tanda "terakhir aktif" seseorang; server tetap memakainya untuk urutan, idempotency, dan audit. Ini syarat model bagi dokumen 08: tanggal kejadian diisi pencatat dan menentukan periode anggaran, laporan, dan jatah pribadi (§3.4), sehingga tanggal pencatatan adalah satu-satunya hal yang membuat tanggal mundur terlihat.

### 3.12 Koreksi dan keberatan

- **Entri pembalik**: salinan entri asal dengan semua `delta` dan `effect` berlawanan tanda, tanggal kejadian sama dengan entri asal, `corrects` menunjuk entri asal. Satu entri dibalik paling banyak satu kali; entri pembalik tidak dapat dibalik (pencatat mencatat ulang entri yang benar).
- **Entri penyesuaian**: entri sejenis yang hanya memuat selisihnya, dengan `corrects` menunjuk entri asal.
- Yang berubah bukan jumlah (kategori, penerima manfaat) diperbaiki dengan `recategorization`.
- Catatan teks dan struk bukan inti moneter; penggantiannya dicatat di jejak audit tanpa entri (§8).
- Kedua pasangan boleh menambahkan koreksi atas entri dana bersama buatan siapa pun; entri asal tetap, pasangan diberi tahu, dan *keberatan* dapat diajukan (default Q-26 yang diterima pemilik produk). Keberatan tidak mengubah angka apa pun; alurnya milik dokumen 06 (kait `entry.disputed`). Entri buku pribadi hanya dapat dikoreksi pemiliknya.
- **Entri dua buku** (kontribusi, talangan, penggantian, penyisihan, pemasukan dengan bagian pribadi, pinjaman antar-pemilik) memuat baris buku pribadi M. Entri pembaliknya menyalin baris itu, jadi hanya M yang dapat mencatatnya (INV-25). Pasangannya tetap dapat (a) menambahkan entri penyesuaian yang seluruh barisnya di buku rumah tangga, misalnya `income` +500000 atas gaji yang tercatat terlalu kecil; dan (b) mengajukan keberatan. Penyesuaian sepihak atas bagian yang berpasangan dengan buku pribadi (baris `owner_transfer`, posisi klaim internal) tertolak dengan sendirinya oleh INV-03 dan INV-08. Batas ini dinyatakan terbuka: "keduanya boleh mengoreksi" berlaku penuh atas entri yang seluruhnya di buku rumah tangga; atas entri dua buku, pasangan dapat menyesuaikan sisi rumah tangganya dan berkeberatan, tetapi tidak dapat membatalkannya sendirian.
- Entri koreksi yang menambah saldo tempat uang berpemegang anggota dewasa selain pengoreksi membawa penanda `unconfirmed` seperti entri lain (§3.3).
- **Klaim internal yang dikeberatani [Usulan]** (temuan `security-privacy` SP-05; dokumen 06 §8.2 "Keberatan"). Selama ada keberatan terbuka atas entri yang memuat baris suatu klaim internal (talangan yang melahirkan hak ganti, pinjaman antar-pemilik, penggantian, pelunasan, pelepasan), klaim itu tampil di daftar utang dan piutang rumah tangga (§5.2; §6.3 tabel E) dengan keterangan "dikeberatani" dan jumlah baris klaim pada entri yang dikeberatani itu. **Tidak ada angka yang berubah**: saldo klaim, jumlah kewajiban dan piutang rumah tangga, dan kekayaan bersih tetap menghitungnya penuh sampai ada entri pembalik; tampilannya sama bagi keduanya. Keterangan itu hilang ketika keberatannya ditutup atau entri sasarannya dibalik.
- Reservasi, anggaran, dan alokasi mengikuti entri: entri pembalik atas belanja yang tertaut ke otorisasi mengurangi realisasinya.

### 3.13 Pengembalian dana

Q-15, K-02. `refund` menambah tempat uang dan membawa baris `expense` bertanda + dengan kategori entri asalnya; wajib tertaut ke entri asal. Jumlah seluruh pengembalian atas satu entri tidak melebihi jumlah entri itu (INV-13). Ia mengurangi pemakaian anggaran kategori itu pada periode **tanggal kejadian pengembalian**, bukan periode belanja aslinya, sehingga laporan lama hanya berubah karena entri terlambat, tidak karena peristiwa baru. Bila pengeluaran aslinya tidak ada di aplikasi, uang yang kembali dicatat sebagai `income` bersumber `other`. Pengembalian ke tempat uang lain daripada asalnya diperbolehkan.

### 3.14 Cicilan

Q-16, K-11.

- Kewajiban cicilan dicatat sebesar **jumlah seluruh pembayaran yang tersisa**, tanpa memisahkan pokok dan bunga. Keluarga tahu "1.500.000 kali 6 lagi"; mereka jarang tahu rincian bunganya **[Asumsi]**.
- Pembayaran cicilan adalah `repayment_out`: uang −X, kewajiban −X. Ia membawa kategori (bawaan "cicilan") dan **terhitung terhadap anggaran kategori itu** serta tampil di belanja per kategori (§6), walaupun bukan baris sebab `expense`. Denda keterlambatan adalah baris `expense` di entri yang sama.
- **Biaya kredit.** Saat barang dibeli secara cicilan di dalam aplikasi: aset bertambah sebesar harga barang, uang berkurang sebesar uang muka, kewajiban bertambah sebesar jumlah seluruh cicilan, dan selisihnya, `credit_cost = uang muka + jumlah seluruh cicilan − harga`, dicatat sebagai baris sebab di entri yang sama. Dengan itu turunnya kekayaan bersih sebesar bunga **punya entri yang menjelaskannya**, pada tanggal pembelian. Ia tampil di ringkasan bulanan sebagai baris "biaya kredit" dan tidak dihitung terhadap anggaran, karena uangnya keluar lewat cicilan bulanan yang sudah dihitung.
- Cicilan yang sudah berjalan sebelum rumah tangga mulai mencatat masuk sebagai saldo awal kewajiban; biaya kreditnya sudah berada di dalam saldo awal dan tidak dicatat lagi.
- Pelunasan dipercepat dengan potongan: `repayment_out` sebesar yang dibayar ditambah `debt_forgiven` sebesar potongannya.
- **Cicilan dan ambang permintaan** (dokumen 06 §6.2, §6.3). Pembelian secara cicilan yang dicatat di aplikasi dibandingkan dengan ambang permintaan **satu kali, saat dibeli**, sebesar uang muka ditambah jumlah seluruh cicilan (§3.8). Pembayaran bulanannya melunasi kewajiban yang sudah tercatat dan terlihat oleh keduanya: ia **tidak pernah** dibandingkan dengan ambang dan **tidak memerlukan kesepakatan rutin**, sehingga cicilan tidak membawa `routine_agreement_ref` (§2.6). Dari sisi uang ini benar: seluruh dana bersama yang diikat kredit itu sudah dihitung satu kali pada saat pembelian, dan menandai tiap pembayaran akan menghitung jumlah yang sama untuk kedua kalinya. Yang dikecualikan hanya bagian yang mengurangi kewajiban: denda keterlambatan di entri yang sama adalah baris `expense`, pengeluaran baru yang belum pernah tercatat sebagai kewajiban, dan dibandingkan dengan ambang seperti pengeluaran lain (§3.8).
- **Kewajiban yang masuk sebagai saldo awal** (dokumen 06 §6.2, putusan 2). Kewajiban dana bersama kepada pihak luar yang didaftarkan sebagai saldo awal saat ada ambang yang berlaku dibandingkan dengan ambang **satu kali, sebesar saldo awalnya** (§3.8); pembayarannya sesudah itu tidak. Dengan itu kredit baru yang dicatat sebagai saldo awal kewajiban, bukan sebagai pembelian secara cicilan, tidak lagi lolos saat dibuka dan saat dibayar. Saldo awal yang diisi saat onboarding, sebelum ambang disepakati, tidak pernah ditandai (**[Asumsi]** contoh: saldo awal hari pertama pada CTH-01, termasuk sisa cicilan motor). Dari sisi uang tidak ada yang terhitung dua kali: saldo awal bukan pengeluaran, tidak masuk anggaran, dan pelunasannya tidak dibandingkan lagi.

Setiap rupiah cicilan terhitung satu kali di tiap tampilan: satu kali di arus kas dan anggaran (saat dibayar), satu kali sebagai pengurang kewajiban, dan bunganya satu kali di perubahan kekayaan bersih (saat kredit dibuka).

### 3.15 Tagihan berulang

Jadwal dan pengingat, bukan entri. Tagihan yang jatuh tempo tetapi belum dibayar **bukan kewajiban** (Q-17): jumlahnya perkiraan, dan kewajiban semu akan mengubah kekayaan bersih tanpa peristiwa uang. Pembayarannya adalah `expense` yang ditautkan ke tagihan itu dengan jumlah sebenarnya. Tagihan boleh diberi **tanggal akhir** (Q-46); setelah lewat, pengingat berhenti dan tagihan terarsip sendiri. Tagihan berulang yang jumlah per kalinya di atas ambang permintaan (misalnya kiriman keluarga bulanan) dirujuk satu *kesepakatan rutin* lewat `routine_agreement_ref` (§2.6; dokumen 06 §6.3).

### 3.16 Aset

Q-18.

- Pembelian tunai (`asset_purchase`): uang −X, aset +X. Bukan pengeluaran; kekayaan bersih tetap. Tampil di arus kas sebagai "beli aset". Tidak terhitung terhadap anggaran kecuali pencatat memberinya kategori.
- Nilai diperbarui manual dengan `asset_valuation` bertanggal: aset ±Δ, sebab `valuation` ±Δ. Tidak ada penyusutan otomatis.
- Penjualan atau pelepasan (`asset_disposal`): aset berkurang sebesar nilai tercatat, uang bertambah sebesar hasil, selisihnya `valuation`.
- Mana yang disebut aset dan mana yang belanja adalah pilihan pencatat. Bawaan **[Usulan]**: barang dicatat sebagai aset hanya bila keluarga ingin melihatnya di daftar "aset apa yang kita punya".
- Pembelian aset di atas ambang permintaan diajukan lebih dulu (keputusan pemilik nomor 5). Yang dibandingkan dengan ambang: jumlah yang dibayar bila tunai; uang muka ditambah jumlah seluruh cicilan bila dicicil (§3.8, CTH-34).

### 3.17 Saldo awal

Entri `opening` per posisi, dengan sebab `opening`. Bukan pemasukan, tidak terhitung terhadap anggaran, tidak muncul di arus kas periode mana pun. Tempat uang, piutang, kewajiban, atau aset yang didaftarkan belakangan mendapat `opening`-nya sendiri pada tanggal pendaftaran. Satu-satunya saldo awal yang punya *jumlah pemakaian* adalah kewajiban dana bersama kepada pihak luar (§3.8, §3.14).

---

## 4. Invarian

Setiap invarian menjadi automated test pada fase pembangunan (`ENGINEERING-STANDARDS.md` §5, §6): property test atas rangkaian perintah acak, ditambah contoh §7 dengan angka yang sama. Semua invarian diperiksa atas **data penuh** (semua buku), bukan atas tampilan seorang anggota.

Notasi: `saldo(p)` = jumlah `delta` semua baris posisi `p`; `U(o)`, `P(o)`, `A(o)`, `K(o)` = jumlah saldo tempat uang, piutang, aset, kewajiban berpemilik `o`; `KB(o) = U(o) + P(o) + A(o) − K(o)`.

| ID | Invarian | Cara menguji |
|---|---|---|
| INV-01 | Setiap `amount`, `delta`, `effect` adalah bilangan bulat dalam satuan terkecil dan membawa mata uang; semua baris satu entri bermata uang sama | Pemeriksaan tipe dan skema; perintah berisi pecahan ditolak |
| INV-02 | **Keseimbangan per pemilik.** Untuk setiap entri `e` dan setiap pemilik `o`: `Σ delta(tempat uang) + Σ delta(piutang) + Σ delta(aset) − Σ delta(kewajiban) = Σ effect(sebab)`, dihitung atas baris `e` yang berpemilik `o` | Untuk setiap entri yang tersimpan |
| INV-03 | **Perpindahan antar-pemilik berjumlah nol.** Dalam satu entri, `Σ effect` semua baris `owner_transfer` di semua buku = 0 | Untuk setiap entri |
| INV-04 | `saldo(p)` yang ditampilkan atau disimpan sebagai cache = jumlah `delta` baris `p` | Hitung ulang dari baris setelah setiap perintah |
| INV-05 | *Uang yang ada* = `U(shared)` = jumlah saldo tercatat semua tempat uang berpemilik `shared`, termasuk yang terarsip dan titipan pada pihak luar | Hitung ulang |
| INV-06 | **Pemindahan tidak mengubah total.** Entri `transfer` tanpa biaya admin: `ΔU(o) = 0` dan `ΔKB(o) = 0`. Dengan biaya admin `f`: `ΔU(o) = ΔKB(o) = −f` | Untuk setiap entri `transfer` |
| INV-07 | **Identitas kekayaan bersih.** Untuk setiap pemilik `o` pada setiap saat: `KB(o) = Σ effect` semua baris sebab berpemilik `o` | Akibat INV-02; diuji langsung |
| INV-08 | **Klaim antar-pemilik dua sisi.** Untuk setiap klaim antara dua pemilik di satu rumah tangga, saldo piutang di buku yang satu = saldo kewajiban di buku yang lain. Akibatnya `Σ KB(o)` atas semua pemilik = uang semua tempat uang + piutang kepada pihak luar + aset − kewajiban kepada pihak luar | Setelah setiap perintah |
| INV-09 | **Hanya menambah.** Entri dan baris yang tersimpan tidak pernah berubah atau hilang: setelah perintah apa pun, himpunan (id, isi inti moneter) entri lama identik dengan sebelumnya. Pengecualian tunggal di §8 hanya menyentuh `entry_detail` dan pemetaan identitas | Bandingkan hash inti moneter sebelum dan sesudah |
| INV-10 | **Idempoten.** Perintah dengan (`household_id`, `idempotency_key`) yang sudah diterima tidak menambah entri, baris, tautan susulan, reservasi, atau otorisasi, dan mengembalikan hasil perintah pertama | Kirim ulang setiap perintah; jumlah baris dan semua saldo tetap |
| INV-11 | `saldo tersedia(h) = saldo tercatat(h) − Σ reservasi aktif(h)`; setiap reservasi ≥ 0; maka saldo tersedia ≤ saldo tercatat | Setelah setiap perintah |
| INV-12 | Persetujuan, penolakan, usul balik, perubahan syarat (`authorization.amended`), otorisasi, reservasi, anggaran, target, tagihan berulang, jadwal cicilan, jadwal penyisihan dan jumlah jatah pribadi, hari mulai periode anggaran, kesepakatan rutin dan rujukannya, konfirmasi custody, keberatan, dan tautan susulan tidak membuat baris posisi atau sebab: semua saldo dan semua `KB(o)` identik sebelum dan sesudahnya | Untuk setiap perintah jenis itu |
| INV-13 | **Batas tautan.** Saldo piutang dan kewajiban ≥ 0 (pelunasan, penggantian, penghapusan, pelepasan tidak melebihi sisa); nilai aset ≥ 0; jumlah pengembalian dana atas satu entri ≤ jumlah entri itu | Perintah yang melanggar ditolak |
| INV-14 | `realisasi(a)` = jumlah *jumlah pemakaian* (§3.8) entri yang tertaut ke otorisasi `a`, lewat tautan pada entri atau tautan susulan yang berlaku, dikurangi pengembalian dana tertaut; berlaku apa pun status `a`. `terikat(a) = max(0, batas(a) − realisasi(a))` bila `a` terbuka, 0 bila tidak; `batas(a)` adalah batas versi syarat yang berlaku | Hitung ulang, termasuk setelah `authorization.amended` dan setelah tautan ke otorisasi tertutup |
| INV-15 | **Anggaran.** Satu baris terhitung terhadap paling banyak satu anggaran (yang kategorinya sama dan periodenya memuat tanggal kejadiannya). `sisa = batas − terpakai − terikat` (§6.1) | Hitung ulang; jumlah `terpakai` semua kategori = jumlah baris berkategori pada periode itu |
| INV-16 | **Rekonsiliasi.** Sesaat setelah entri `cash_difference`, saldo tercatat = saldo terhitung. Untuk tiap rekonsiliasi, `0 ≤ jumlah dirinci ≤ |selisih|` | Setelah setiap rekonsiliasi dan `recategorization` |
| INV-17 | **Alokasi.** Alokasi per (tempat uang, tujuan) ≥ 0. `belum dialokasikan(h) = saldo tercatat(h) − Σ alokasi(h)`. Perintah `allocation` yang membuat `Σ alokasi(h) > saldo tercatat(h)` ditolak | Setelah setiap perintah |
| INV-18 | **Koreksi.** Entri pembalik memuat baris yang sama dengan entri asal dengan tanda berlawanan; jumlah akibat keduanya pada setiap saldo dan setiap `KB(o)` nol. Satu entri punya paling banyak satu pembalik | Untuk setiap pembalikan |
| INV-19 | **Jembatan laporan.** Untuk setiap pemilik dan setiap rentang tanggal kejadian: `ΔKB = Σ effect` baris sebab dalam rentang; `ΔU = Σ delta` baris tempat uang dalam rentang. Jumlah baris arus kas (§6.4) = `ΔU`; jumlah baris ringkasan (§6.3) = `ΔKB` | Hitung dua cara |
| INV-20 | **Buku rumah tangga bebas dari dana pribadi.** Tidak ada angka rumah tangga (uang yang ada, kekayaan bersih, anggaran, laporan) yang nilainya berubah ketika baris berpemilik `member(M)` ditambah, kecuali lewat baris berpemilik `shared` dalam entri yang sama. Angka rumah tangga identik bagi kedua pelihat | Tambahkan entri buku pribadi murni; angka rumah tangga tetap. Hitung angka rumah tangga sebagai tiap pelihat; hasilnya sama |
| INV-21 | **Arsip.** Posisi terarsip bersaldo nol dan tidak menerima baris baru | Perintah yang melanggar ditolak |
| INV-22 | `event_date ≤` tanggal `recorded_at`; `event_date ≥` tanggal `opening` setiap posisi yang disentuh | Perintah yang melanggar ditolak |
| INV-23 | **Bersamaan.** Hasil dua perintah sah yang diterima bersamaan sama dengan hasil menjalankannya berurutan dalam salah satu urutan; tidak ada baris yang hilang | Uji konkurensi (CTH-31) |
| INV-24 | **Penghasilan yang dicatat sebagai penghasilan selalu tercatat penuh di buku rumah tangga.** *Daftar sumber penghasilan* adalah parameter preset aturan rumah tangga; preset pilot: `salary`, `business` (§3.1). Setiap baris sebab `income` yang sumbernya termasuk daftar yang berlaku saat entri diterima berpemilik `shared`; perintah yang menulisnya di buku pribadi ditolak. Dalam entri `income` **yang memuat baris sebab `income` bersumber penghasilan**, tempat uang berpemilik `member(M)` hanya bertambah lewat sepasang baris `owner_transfer` di entri yang sama (−X di buku rumah tangga, +X di buku M), dengan 0 ≤ X ≤ jumlah pemasukan. Entri `income` yang sumbernya di luar daftar itu tidak dibatasi kalimat ini: hadiah untuk satu orang, atau pemasukan lain yang memang uang pribadi, dicatat sebagai `income` di buku pribadi penerimanya saja (§3.1, CTH-36). Akibatnya jumlah penuh setiap penghasilan dan bagian yang diarahkan ke dana pribadi terbaca dari buku rumah tangga saja, sama bagi kedua pelihat | Bangkitkan perintah uang masuk acak dari kedua anggota: tidak ada baris `income` bersumber penghasilan di buku pribadi; untuk setiap entri `income` bersumber penghasilan, kenaikan tempat uang pribadi = −Σ `owner_transfer` di buku rumah tangga; baca buku rumah tangga sebagai tiap pelihat, hasilnya sama (CTH-03b). Entri `income` bersumber di luar daftar yang seluruh barisnya di buku pribadi pencatat diterima (CTH-36). Ulangi dengan daftar sumber lain (misalnya ditambah `other`): aturan yang sama berlaku atas daftar itu, dan entri yang diterima sebelum perubahan tidak berubah |
| INV-25 | **Tulis buku pribadi.** Entri yang memuat baris berpemilik `member(M)` hanya diterima bila `actor = M`. Dua pengecualian, keduanya tanpa baris tempat uang atau aset milik M: (1) `claim_waiver` atas kewajiban M kepada dana bersama, dicatat anggota dewasa selain M, yang di buku M hanya memuat baris kewajiban − dan `owner_transfer` + (§2.2); (2) `exit_closure` oleh sistem saat M keluar (INV-28) | Perintah dari anggota lain yang menyentuh buku M ditolak tanpa akibat apa pun; tidak ada entri tersimpan yang memuat baris tempat uang atau aset milik M dengan `actor ≠ M` |
| INV-26 | **Tautan.** Menambah atau mencabut tautan susulan tidak mengubah entri, baris, atau saldo apa pun (INV-09). Satu entri tertaut ke paling banyak satu otorisasi yang berlaku. Tautan ke otorisasi hanya sah bila otorisasi itu `open`, atau `settled`/`expired` dengan tanggal kejadian entri di antara tanggal persetujuan dan tanggal penutupannya; tidak pernah ke yang `cancelled`. Entri `repayment_out` atas suatu kewajiban tidak dapat ditautkan ke otorisasi yang sudah ditaut penuh oleh entri yang melahirkan kewajiban itu: `asset_purchase` secara cicilan, atau `opening` kewajiban dana bersama kepada pihak luar (§3.8; dokumen 06 §4.5 "Masa berlaku dan kedaluwarsa") | Hash inti moneter sebelum dan sesudah; perintah tautan yang melanggar ditolak; entri itu sendiri tetap tersimpan |
| INV-27 | **Bagian pribadi dari penghasilan hanya bagi penerimanya. [Usulan]; menunggu pemilik produk di checkpoint S2** (§3.1). Baris `income` bersumber penghasilan (INV-24) membawa `earner` yang bernilai tepat satu anggota dewasa rumah tangga itu. Entri `income` bersumber penghasilan hanya boleh memuat baris `owner_transfer` (bagian pribadi X > 0) bila `actor = earner`; perintah yang melanggar ditolak seluruhnya, dan perintah yang sama tanpa bagian pribadi diterima. Alternatif bila pemilik produk menolak usulan ini: invarian ini dihapus dan yang berlaku hanya INV-24 dan INV-25 | Bangkitkan perintah uang masuk bersumber penghasilan dari kedua anggota, dengan dan tanpa bagian pribadi, dengan `earner` diri sendiri dan pasangan: yang berbagian pribadi dengan `actor ≠ earner` ditolak tanpa akibat apa pun; tiga kombinasi lain diterima. Tidak ada entri tersimpan bersumber penghasilan yang memuat `owner_transfer` dengan `actor ≠ earner` (CTH-03b: keduanya suami) |
| INV-28 | **Penutupan klaim saat anggota keluar. [Usulan]** (§8.2). Entri `exit_closure` hanya ditulis sistem (`actor = system`), tepat satu kali per anggota yang keluar, dalam transaksi yang sama dengan penutupan buku pribadinya. Ia hanya memuat baris piutang, kewajiban, dan `owner_transfer`; tidak pernah baris tempat uang atau aset. Sesudahnya saldo setiap klaim antar-pemilik yang menyangkut buku pribadi anggota itu nol di kedua sisinya. Ia memenuhi INV-02 dan INV-03, tidak mengubah `U(o)` pemilik mana pun, dan tidak dapat dibalik | Bangkitkan klaim acak antara buku pribadi M dan buku lain, lalu keluarkan M: semua klaim itu bersaldo nol di kedua sisi; `ΔU` semua pemilik nol; `ΔKB(shared)` = Σ kewajiban rumah tangga kepada M − Σ piutang rumah tangga kepada M yang ditutup; perintah `exit_closure` dari anggota, perintah kedua untuk anggota yang sama, dan perintah pembaliknya ditolak (CTH-38) |
| INV-29 | **Periode anggaran membagi habis garis tanggal** (§6.1). Hari mulai adalah bilangan bulat 1 sampai 28. Untuk setiap rumah tangga, setiap tanggal berada di tepat satu periode; periode berurutan tanpa celah dan tanpa tumpang-tindih. Periode biasa panjangnya 28 sampai 31 hari; periode peralihan 1 sampai 30 hari. Mengganti hari mulai tidak mengubah batas periode mana pun yang sudah dimulai pada tanggal perubahan itu diterima, dan tidak mengubah saldo atau `KB(o)` mana pun (INV-12) | Rangkaian acak penggantian hari mulai pada tanggal acak: setiap tanggal dalam rentang uji jatuh di tepat satu periode; batas periode yang sudah dimulai identik sebelum dan sesudah tiap penggantian; jumlah `terpakai` semua periode per kategori = jumlah baris berkategori itu (INV-15); hari mulai di luar 1 sampai 28 ditolak |
| INV-30 | **Jatah pribadi hanya membagi, tidak mengubah** (§3.4). Untuk setiap entri `e` dan anggota M: `0 ≤ dalam_jatah(e) ≤ bagian(e, M)` dan `dalam_jatah(e) + luar_jatah(e) = bagian(e, M)`. Dalam satu periode, jumlah `dalam_jatah` atas entri yang tidak dibalik ≤ jumlah jatah terbesar yang berlaku dalam periode itu. Hitungan satu periode tidak bergantung pada entri periode lain. Jatah, perubahannya, dan hitungannya tidak membuat baris dan tidak mengubah saldo, anggaran, *terikat*, reservasi, atau `KB(o)` | Bangkitkan penyisihan dan pemasukan berbagian pribadi acak dalam beberapa periode, dengan jatah yang berubah: ketiga hubungan itu berlaku; menghapus semua entri periode lain dari masukan tidak mengubah hasil suatu periode; semua saldo dan `KB(o)` identik dengan dan tanpa jatah (CTH-39) |

Yang sengaja **bukan** invarian: saldo tercatat tempat uang ≥ 0 (§2.5), *belum dialokasikan* ≥ 0 (§3.9), dan saldo tersedia ≥ 0 (§3.8). Ketiganya adalah sinyal bagi pengguna, bukan penolakan.

---

## 5. Tampilan custody

Kriteria sukses nomor 1: "sisa berapa dan siapa yang pegang", tanpa bertanya.

### 5.1 Perhitungan

```text
custody rumah tangga =
  untuk setiap tempat uang h berpemilik shared, tidak terarsip atau bersaldo bukan nol:
    (pemegang(h), jenis(h), nama(h), saldo tercatat(h), saldo tersedia(h),
     jumlah baris masuk yang belum dikonfirmasi,
     jumlah baris keluar entri dua buku yang belum dikonfirmasi,
     jumlah yang dikeberatani,                                               -- §3.3; dokumen 06 §7
     alokasi per tujuan, belum dialokasikan)
  dikelompokkan per pemegang: anggota M | dipegang bersama | pihak luar P (titipan)

angka utama = uang yang ada = Σ saldo tercatat(h)        -- INV-05
```

Jumlah semua baris custody selalu sama dengan angka utama. Tidak ada baris "lainnya".

### 5.2 Yang dilihat tiap anggota dewasa

| Bagian | Isi | Pelihat A | Pelihat B |
|---|---|---|---|
| Angka utama | *Uang yang ada* | Sama | Sama |
| Rincian per pemegang | Tiap tempat uang dana bersama: pemegang, tempat, saldo, yang belum dikonfirmasi, yang "sudah dijanjikan" (reservasi) | Sama | Sama |
| Dipinjam orang, utang kita, sisa cicilan, aset, harta bersih | Dari buku rumah tangga, termasuk hak ganti yang belum dibayar kepada A atau B. Klaim internal yang sedang dikeberatani diberi keterangan "dikeberatani" beserta jumlah yang dikeberatani; angkanya tetap terhitung penuh (§3.12) | Sama | Sama |
| Baris kedua: "uang pribadi saya" | Tempat uang berpemilik pelihat itu sendiri, beserta klaimnya | Hanya milik A | Hanya milik B |

Buku rumah tangga identik bagi keduanya (INV-20). Baris kedua tidak dijumlahkan ke angka utama. Visibilitas per tempat uang (`private | spouses | all_members`) ditetapkan dokumen 06 dan 10; model hanya menuntut bahwa tempat uang berpemilik `shared` tidak dapat `private`, sebab angka utama harus sama bagi keduanya. Bila seorang anggota memilih memperlihatkan tempat uang pribadinya kepada pasangan, ia tampil di layar pasangan sebagai bagian terpisah dan tetap tidak masuk angka rumah tangga.

### 5.3 Kebocoran yang tersisa

Dinyatakan terbuka. Anggota diberi tahu sebelum mencatat hal yang akan terlihat pasangannya (Q-24).

| # | Yang diketahui pasangan | Sebab | Dapat dihindari? |
|---|---|---|---|
| B1 | Jumlah dan tanggal setiap penyisihan dana bersama menjadi dana pribadi. Termasuk bagian penghasilan yang diarahkan langsung ke dana pribadi saat dicatat: pasangan melihat jumlah penuh penghasilannya dan bagian pribadinya | Entri buku rumah tangga; keputusan pemilik produk, 11 November 2026; INV-24 | Tidak, dan memang tidak boleh: ini pagar terhadap pengalihan sepihak |
| B2 | Bahwa anggota punya dana pribadi sekurangnya sebesar kontribusi atau talangannya, pada tanggal itu | Keputusan pemilik nomor 3: kontribusi terlihat beserta sumbernya | Hanya dengan tidak mencatatnya sebagai keperluan rumah tangga |
| B3 | Hak ganti dan utang antar-pasangan: jumlah, arah, dan pelunasannya | Klaim punya dua sisi | Tidak |
| B4 | Batas bawah kasar dana pribadi: jumlah penyisihan yang diterima dikurangi kontribusi yang terlihat | Akibat B1 dan B2 | Tidak sepenuhnya; pemasukan dan belanja pribadi yang tidak terlihat membuatnya bukan saldo sebenarnya. Aplikasi **tidak menghitungkannya**: tidak ada angka kumulatif, lintas periode, atau bersih per anggota atas penyisihan dan kontribusi (§6.3). Pasangan yang menjumlahkan sendiri baris per periode tetap dapat memperolehnya |
| B5 | Hadiah untuk satu orang dan pemasukan lain yang dicatat langsung di buku pribadi **tidak** terlihat sama sekali **[Usulan]**; untuk hadiah, menunggu pemilik produk di checkpoint S2 (§3.1). Penghasilan tidak dapat dicatat demikian (INV-24); tetapi bila gaji suatu bulan tidak dicatat, pasangan melihat gajinya tidak muncul | Keputusan pemilik nomor 2 dan 3 menjadikannya dana pribadi yang tertutup; bahwa catatannya ikut tertutup adalah usulan dokumen ini. Ketiadaan entri juga informasi | Tidak |

Yang tetap tertutup: saldo tempat uang pribadi, jumlah dan nama tempat uang pribadi, pemasukan pribadi selain penghasilan, belanja pribadi, dan aset pribadi. Dari penghasilan yang diarahkan ke dana pribadi, yang tertutup adalah tempat uang pribadi tujuannya, saldonya, dan pemakaiannya sesudah itu.

Risiko yang berlawanan arah, untuk `security-privacy`: jaminan INV-24 berlaku atas pemasukan yang **diberi sumber** gaji atau hasil usaha. Sumber diisi pencatat; penghasilan yang dicatat sebagai hadiah atau "lainnya" di buku pribadi, atau tidak dicatat sama sekali, tetap tidak terlihat. Lihat §10.3 butir 1.

---

## 6. Anggaran dan laporan

Semua angka di bagian ini dihitung dari buku rumah tangga, menurut **tanggal kejadian**, dan sama bagi kedua pasangan.

### 6.1 Anggaran

Keputusan pemilik nomor 8: anggaran adalah batas atas satu kategori buatan rumah tangga, jumlahnya dapat diubah, periodenya bulanan, tanpa amplop. Periode disimpan sebagai rentang tanggal [`period_start`, `period_end`], sehingga periode mingguan atau harian (Tier 3) tidak mengubah struktur.

**Periode anggaran adalah setelan rumah tangga** (keputusan pemilik produk, 11 November 2026: "bisa di tentukan pasti ada yang fixed tanggal ada juga yang gak tertentu"; menyempurnakan keputusan S1 "bulanan mulai tanggal gajian"). Menurut pemilik produk, gajinya "bisa 2x" dan uang masuk dari pekerjaan lepas "bisa sampe 5x"; bahwa itu berarti sampai sekitar tujuh uang masuk sebulan adalah **bacaan lead**, bukan kata-katanya (§3.1). Pada bacaan mana pun tidak ada satu "tanggal gajian" yang dapat dijadikan patokan; rumah tangga memilih sendiri hari mulainya. Asasnya keputusan pemilik produk; rentang hari, bentuk peralihan, dan batas anggaran pada periode peralihan adalah **[Usulan]** dokumen ini (T-23, T-24).

| Hal | Aturan |
|---|---|
| Setelan | Satu **hari mulai** per rumah tangga, bilangan bulat 1 sampai 28; bawaan 1. Berlaku atas semua kategori; tidak ada hari mulai per kategori |
| Mengapa 1 sampai 28 | Setiap bulan punya tanggal 28, jadi periode selalu mulai pada tanggal yang sama tanpa aturan tambahan untuk Februari dan bulan 30 hari. Rumah tangga yang penghasilannya tiba pada tanggal 29 sampai 31 memilih 28 atau 1. Biayanya: "akhir bulan" tidak dapat dipilih persis |
| Periode biasa | Dari hari mulai suatu bulan sampai sehari sebelum hari mulai bulan berikutnya; panjangnya 28 sampai 31 hari |
| Mengganti hari mulai | Berlaku **mulai periode berikutnya**. Periode yang sedang berjalan saat penggantian diterima, dan semua periode sebelumnya, tidak berubah batasnya. Penggantian tercatat sebagai baris baru `budget_period_setting` (§2.6) dan sebagai *peristiwa* beserta pelakunya; tidak membuat entri (INV-12) |
| Periode peralihan | **Satu periode pendek.** Ia mulai sehari sesudah periode berjalan berakhir dan berakhir sehari sebelum hari mulai baru yang pertama sesudahnya; panjangnya 1 sampai 30 hari. Sesudah itu periode biasa menurut hari mulai baru |
| Mengapa pendek, bukan satu periode panjang | Periode panjang dapat mencapai hampir dua bulan dan memuat dua kali cicilan, dua kali tagihan, dan dua kali kiriman bulanan terhadap batas satu bulan: anggaran pasti tampil terlampaui padahal tidak ada yang salah. Periode pendek paling banyak memuat satu kali tiap pembayaran bulanan; akibat terburuknya sisa anggaran yang besar, yang tidak berpindah (tanpa rollover) |
| Anggaran pada periode peralihan | Baris `budget` dibuat dari `budget_rule` dengan batas penuh, tidak diprorata: prorata menghasilkan pecahan dan menebak pola belanja. Rumah tangga boleh mengubah batas periode itu seperti batas periode lain |
| *Terikat* | Tidak disimpan per periode; dihitung dari tanggal rencana atau tanggal persetujuan otorisasi (di bawah). Otorisasi terbuka yang tanggalnya jatuh di periode peralihan terhitung di periode itu |
| Laporan | Ringkasan bulanan (§6.3) mengikuti periode anggaran; periode peralihan punya ringkasannya sendiri dan tampil dengan rentang tanggalnya. Laporan periode lampau tidak berubah karena penggantian. Rincian per anggota tetap hanya di dalam satu periode (§6.3) |
| Entri terlambat dan pengembalian dana | Mengikuti tanggal kejadiannya ke periode yang memuat tanggal itu, menurut batas yang berlaku bagi periode tersebut |
| *Jatah pribadi* | Per periode, termasuk periode peralihan, dengan jumlah penuh (§3.4) |
| Penggantian beruntun | Yang berlaku adalah penggantian terakhir yang diterima sebelum periode berjalan berakhir |
| Siapa yang boleh mengganti | **Kesepakatan rumah tangga: kedua anggota dewasa, seperti nilai ambang** (putusan lead, 11 November 2026; bukan keputusan pemilik produk; T-25). Satu mengusulkan, yang lain menyetujui; hari mulai lama berlaku sampai disetujui, dan tidak seorang pun menggantinya sendirian. Alurnya milik dokumen 06 §6.3 "Kesepakatan rumah tangga". Alasannya dari sisi uang: penggantian membuka satu periode pendek dengan batas anggaran dan jatah pribadi penuh |
| Terlihat | Usulan, persetujuan, dan penggantian itu sendiri tercatat sebagai *peristiwa* beserta pelakunya dan terlihat oleh kedua pasangan; periode peralihan tampil dengan rentang tanggalnya |

```text
urutan baris budget_period_setting: menurut (effective_from, recorded_at)
s0 = baris pertama menurut urutan itu   -- setelan awal rumah tangga; bawaan start_day = 1

periode(tanggal t) =
  s = baris terakhir menurut urutan itu dengan s.effective_from ≤ t
  bila tidak ada:  s = s0
      -- s0 berlaku pula atas semua tanggal sebelum effective_from-nya
      --   (entri terlambat dan saldo awal bertanggal sebelum rumah tangga mulai mencatat)
  p = tanggal pertama ≥ s.effective_from yang harinya s.start_day
  bila s ≠ s0 dan t < p:
      [s.effective_from, p − 1]                               -- periode peralihan
  selain itu:
      [hari s.start_day terakhir yang ≤ t, sehari sebelum hari s.start_day berikutnya]
      -- s0 tidak pernah membuka periode peralihan, walaupun effective_from-nya
      --   bukan hari start_day: tanggal-tanggal pertamanya masuk periode biasa
      --   yang mulai pada hari start_day sebelumnya

penggantian hari mulai menjadi d (d ≠ start_day yang berlaku), diterima pada tanggal c:
  E = akhir periode(c);  tambah baris (start_day = d, effective_from = E + 1)
  -- penggantian kedua dalam periode yang sama menghasilkan effective_from yang sama;
  --   yang berlaku adalah yang recorded_at-nya terakhir (baris "Penggantian beruntun")
```

Pemeriksaan kedua contoh di bawah dengan fungsi ini. Pada keduanya setelan lama adalah s0 (atau baris lain yang periodenya sudah biasa) dan c = 10 November 2026. **25 → 1:** periode(10 November) = [25 Oktober, 24 November], jadi E = 24 November dan baris baru (1, 25 November). Untuk t dari 25 November: s = baris baru ≠ s0; p = 1 Desember; t < p memberi [25 November, 30 November], 6 hari; sejak 1 Desember periode biasa [1 Desember, 31 Desember]. **1 → 25:** periode(10 November) = [1 November, 30 November], jadi E = 30 November dan baris baru (25, 1 Desember); p = 25 Desember; peralihan [1 Desember, 24 Desember], 24 hari; sesudahnya [25 Desember, 24 Januari]. Hasilnya sama dengan tabel. Fungsi lama memberi hasil yang sama untuk kedua contoh ini; yang dikoreksi hanya dua kasus tepi: tanggal sebelum baris pertama (dulu tanpa setelan) dan baris pertama yang tidak mulai pada hari `start_day` (dulu membuka periode peralihan yang tidak pernah disepakati).

| Hari mulai lama → baru, diterima 10 November 2026 | Periode berjalan (tetap) | Periode peralihan | Periode biasa pertama sesudahnya |
|---|---|---|---|
| 25 → 1 | 25 Oktober–24 November | 25–30 November (6 hari) | 1–31 Desember |
| 1 → 25 | 1–30 November | 1–24 Desember (24 hari) | 25 Desember–24 Januari |

Invarian yang berlaku: INV-29 (setiap tanggal di tepat satu periode; periode lampau tidak berubah), INV-15 (satu baris terhitung terhadap paling banyak satu anggaran), INV-19 (jembatan laporan berlaku atas rentang tanggal apa pun, jadi juga atas periode peralihan), INV-12 dan INV-09 (penggantian tidak membuat atau mengubah entri), dan INV-30.

```text
baris anggaran(k, periode) = semua baris b dengan b.category = k, b.owner = shared,
                             event_date entri b di dalam periode, yang berupa:
    sebab expense (termasuk pengembalian dana bertanda +, dan expense hasil recategorization)
    posisi kewajiban pada repayment_out yang berkategori (pembayaran cicilan)
    baris lain yang diberi kategori oleh pencatat: owner_transfer keluar, asset_purchase, loan_given

terpakai(k, periode)  = − Σ akibat baris anggaran itu pada uang atau kekayaan bersih
                        -- pengeluaran 150.000 menambah terpakai 150.000; pengembalian 20.000 menguranginya
terikat(k, periode)   = Σ terikat(a) untuk otorisasi a terbuka dengan a.category = k
                        dan tanggal rencana a (atau tanggal persetujuannya) di dalam periode
sisa(k, periode)      = batas − terpakai − terikat          -- boleh negatif
```

- Jumlah *terikat* mengurangi sisa anggaran dan tampil terpisah dari yang terpakai (Q-14). Jalan-jalan dan acara tidak punya anggaran periode; batasnya adalah otorisasinya sendiri. Bila permintaannya tidak berkategori, atau kategorinya tidak beranggaran, tidak ada sisa anggaran yang dikurangi.
- `credit_cost`, `write_off`, `valuation`, `cash_difference` yang belum dirinci, dan `opening` tidak pernah terhitung terhadap anggaran.
- Kontribusi belanja dan talangan terhitung penuh: yang diukur anggaran adalah pemakaian untuk kategori itu, bukan dari kantong mana uangnya.
- Mengubah batas anggaran tidak membuat entri; perubahan tercatat di jejak audit.
- Tidak ada rollover di Tier 1: sisa periode lalu tidak berpindah.

### 6.2 Realisasi dan terikat

Didefinisikan di §3.8 dan INV-14. *Realisasi* atas satu anggaran sama dengan `terpakai` di §6.1.

### 6.3 Ringkasan bulanan

Kriteria sukses nomor 3. Periode ringkasan adalah *periode anggaran* rumah tangga (§6.1). Setiap baris sebab buku rumah tangga dalam periode masuk tepat satu baris tabel A; jumlah tabel A = perubahan kekayaan bersih rumah tangga pada periode itu (INV-19).

**Tabel A — mengapa harta bersih berubah.**

| Baris | Rumus |
|---|---|
| Uang masuk (pemasukan), per sumber dan atas nama siapa | `Σ effect(income)` |
| Belanja per kategori | `− Σ effect(expense)` per kategori, sesudah pengembalian dana |
| Biaya kredit | `− Σ effect(credit_cost)` |
| Selisih hitung yang belum dirinci | `Σ effect(cash_difference)`, negatif dan positif ditampilkan terpisah |
| Piutang diikhlaskan | `− Σ effect(write_off)` |
| Perubahan nilai aset | `Σ effect(valuation)` |
| Kontribusi dari dana pribadi, per anggota | `Σ effect(owner_transfer > 0)` pada entri `owner_transfer`, `contribution_expense`, dan `claim_waiver` (anggota melepas hak gantinya, §3.6) |
| Disisihkan ke dana pribadi, per anggota (termasuk bagian pemasukan yang langsung ke dana pribadi, §3.1; pemasukannya sendiri tampil penuh di baris pertama) | `− Σ effect(owner_transfer < 0)` pada entri `owner_transfer` dan `income` |
| Utang anggota kepada dana bersama dihapuskan ("utang dihapuskan") | `− Σ effect(owner_transfer < 0)` pada entri `claim_waiver` (§3.6); tidak pernah dijumlahkan ke baris "disisihkan" |
| Klaim ditutup karena anggota keluar | `Σ effect(owner_transfer)` pada entri `exit_closure`, per klaim (§8.2) |
| Utang yang dibebaskan pihak luar | `Σ effect(debt_forgiven)` |

Setiap baris sebab `owner_transfer` di buku rumah tangga masuk tepat satu dari empat baris di atas, menurut jenis entrinya dan tandanya; jenis entri lain tidak memuat sebab itu (§2.5). Baris entri pembalik masuk baris yang sama dengan baris entri asalnya dan menguranginya, bukan baris yang sesuai dengan tandanya sendiri: penyisihan yang dibalik mengurangi "disisihkan", tidak menambah "kontribusi".

**Angka per anggota atas penyisihan dan kontribusi hanya per periode [Usulan]** (temuan `security-privacy` SP-13; dokumen 06 §9.4 "Keterangan pelaku tanpa menjadi alat pantau", AT-4). Baris "kontribusi dari dana pribadi" dan "disisihkan ke dana pribadi" dirinci per anggota hanya di dalam **satu periode laporan**. Tidak ada tampilan, laporan, ekspor, atau narasi yang memuat, per anggota: (a) jumlah kumulatif sejak awal; (b) jumlah lintas periode, termasuk tren, rata-rata, dan perbandingan antar-periode; (c) angka bersih penyisihan dikurangi kontribusi. Laporan yang rentangnya lebih dari satu periode menampilkan kedua baris itu sebagai jumlah rumah tangga, tanpa rincian per anggota. Alasannya: penyisihan dikurangi kontribusi adalah batas atas kasar dana pribadi seorang anggota yang dana pribadinya hanya berasal dari dana bersama, yaitu dari penyisihan termasuk bagian pribadi penghasilannya sendiri (§5.3, B4); aplikasi tidak boleh menjadi alat yang menghitungnya. Batasnya dinyatakan di B4. Aturan yang sama berlaku atas kolom "dana pribadi, kontribusi" pada tabel C.

**Tabel B — uang keluar per kategori** (dasar yang sama dengan anggaran): `terpakai(k, periode)` untuk setiap kategori, termasuk cicilan yang dibayar. Tabel B menjawab "uang kita dipakai untuk apa"; tabel A menjawab "mengapa harta kita berubah". Pembayaran cicilan ada di B dan tidak di A; biaya kredit ada di A dan tidak di B. Kedua tabel tidak pernah dijumlahkan.

**Tabel C — siapa membayar apa.** Baris `expense` dan pembayaran cicilan dikelompokkan menurut `payer`, dengan tiga kolom sumber: dana bersama; dana pribadi sebagai kontribusi; dana pribadi yang diganti (talangan). Jumlah tabel C per kategori = tabel B. Laporan per orang hanya memuat buku rumah tangga dan sama bagi keduanya (Q-27).

**Tabel D — tabungan.** Per tujuan berjenis `savings`: jumlah alokasi sekarang, perubahan dalam periode, target bila ada.

**Tabel E — aset, piutang, kewajiban.** Saldo akhir tiap posisi dan perubahannya dalam periode. Klaim internal yang sedang dikeberatani diberi keterangan "dikeberatani" tanpa perubahan angka (§3.12).

Laporan periode lampau yang berubah karena entri terlambat menandainya (§3.11).

### 6.4 Arus kas

`ΔU(shared)` dalam periode, dirinci dari baris tempat uang menurut jenis entrinya. Jumlah semua baris = perubahan *uang yang ada* (INV-19). Pemindahan tidak muncul.

| Baris arus kas | Dari |
|---|---|
| Uang masuk: pemasukan | `income`: yang benar-benar masuk tempat uang dana bersama, yaitu jumlah pemasukan dikurangi bagian yang langsung ke dana pribadi; bagian itu disebut di samping barisnya dan tidak menjadi baris "disisihkan" tersendiri di arus kas, karena tidak pernah berada di tempat uang dana bersama |
| Uang masuk: pinjaman diterima | `loan_received` |
| Uang masuk: piutang dibayar | `repayment_in` |
| Uang masuk: kontribusi uang dari dana pribadi | `owner_transfer` masuk |
| Uang masuk: hasil jual aset | `asset_disposal` |
| Uang masuk: selisih hitung positif | `cash_difference` > 0 |
| Uang keluar: belanja (sesudah pengembalian dana, termasuk biaya admin) | `expense`, `refund`, baris tempat uang saja |
| Uang keluar: cicilan dan utang dibayar | `repayment_out` |
| Uang keluar: penggantian dibayar | `reimbursement` |
| Uang keluar: pinjaman diberikan | `loan_given` |
| Uang keluar: beli aset (termasuk uang muka) | `asset_purchase` |
| Uang keluar: disisihkan ke dana pribadi | `owner_transfer` keluar |
| Uang keluar: selisih hitung negatif | `cash_difference` < 0, sebesar selisih penuh |

Pengeluaran yang tidak memakai dana bersama saat itu (kontribusi belanja, talangan, biaya kredit, selisih yang dirinci) tidak punya baris tempat uang di buku rumah tangga dan karena itu tidak muncul di arus kas; uangnya muncul lewat baris penggantian, cicilan, atau selisih.

### 6.5 Kekayaan bersih

```text
kekayaan bersih rumah tangga = U(shared) + P(shared) + A(shared) − K(shared)
```

`P(shared)` dan `K(shared)` memuat klaim kepada pihak luar **dan** kepada anggota (hak ganti, utang antar-pasangan dengan dana bersama), karena dana pribadi di luar garis. Titipan pada pihak luar ada di `U`, bukan `P`. Target, anggaran, otorisasi, dan reservasi tidak masuk.

---

## 7. Contoh terhitung

### 7.1 Cara membaca

Setiap contoh kelak menjadi executable test dengan angka yang sama (`ENGINEERING-STANDARDS.md` §5).

**Rumah tangga contoh [Asumsi].** Angka brief §1.1 (gaji 10.000.000; kiriman ibu 3.000.000) ditambah angka ilustrasi. Rangkaian contoh hanya memuat satu penghasilan dalam periodenya, gaji suami, yang ia terima dan ia catat sendiri. Itu penyederhanaan contoh, bukan gambaran rumah tangga pilot: menurut pemilik produk (11 November 2026) kedua pasangan berpenghasilan. Penghasilan istri dicatat dengan bentuk yang sama (CTH-03, CTH-03b) dengan `earner` istri; semua aturan berlaku per penerima penghasilan dan sama bagi anggota dewasa mana pun. Satu periode dengan beberapa penghasilan, ke rekening bank dan dompet digital, ada di CTH-39. Hari mulai periode anggaran rumah tangga contoh adalah 25 (pilihan rumah tangga itu; bawaannya 1, §6.1); periode contoh 25 Oktober–24 November 2026. Ambang permintaan 500.000.

**Kolom.** Semua angka adalah rupiah utuh, ditulis tanpa pemisah ribuan agar dapat dihitung ulang dengan skrip.

| Kolom | Arti | Pemilik | Pemegang |
|---|---|---|---|
| `BNK-S` | Rekening bank tempat gaji masuk | `shared` | suami |
| `TUN-S` | Uang tunai bersama di suami | `shared` | suami |
| `TUN-I` | Uang tunai bersama di istri | `shared` | istri |
| `EWL-I` | Dompet digital | `shared` | istri |
| `TAB-I` | Rekening tabungan dana darurat | `shared` | istri |
| `UANG` | *Uang yang ada* = `BNK-S + TUN-S + TUN-I + EWL-I + TAB-I` | | |
| `PIU` | Piutang rumah tangga (pinjaman kepada adik suami) | `shared` | |
| `AST` | Nilai aset rumah tangga | `shared` | |
| `KWJ-C` | Kewajiban cicilan: jumlah seluruh pembayaran tersisa | `shared` | |
| `KWJ-U` | Utang rumah tangga kepada pihak luar | `shared` | |
| `KWJ-G` | Hak ganti yang belum dibayar kepada anggota (kewajiban rumah tangga; sisi lainnya piutang anggota itu di buku pribadinya) | `shared` | |
| `KB` | Kekayaan bersih rumah tangga = `UANG + PIU + AST − KWJ-C − KWJ-U − KWJ-G` | | |
| `PRI-S` | Uang tunai pribadi suami (tidak terlihat istri) | `member(suami)` | suami |
| `PRI-I` | Uang tunai pribadi istri (tidak terlihat suami) | `member(istri)` | istri |

Baris `Δ` adalah entri; baris `Sesudah` = baris sebelumnya + `Δ`, kolom demi kolom. Contoh CTH-01 dan CTH-03 sampai CTH-36 **berantai**: `Sesudah` terakhir satu contoh adalah `Sebelum` contoh berikutnya. CTH-02 dan contoh bertanda "berdiri sendiri" tidak mengubah keadaan berantai; semuanya memakai kolomnya sendiri, kecuali CTH-03b, yang memakai kolom rantai dengan `Sebelum` yang sama dengan CTH-03.

**Peta kasus brief §7 ke contoh.**

| Kasus brief §7 | Contoh |
|---|---|
| Spending from personal cash | CTH-12 |
| Spending from shared household cash | CTH-11 |
| Husband transfers money to wife | CTH-08 |
| Allocating money to mother-in-law | CTH-06 (kiriman keluarga), CTH-02 (titipan) |
| Paying on behalf of another member | CTH-17, CTH-17b |
| Reimbursement | CTH-14, CTH-15, CTH-23 |
| Moving money between bank, e-wallet, and cash | CTH-07, CTH-18 |
| Full and partial refunds | CTH-19 |
| Recurring bills | CTH-05 |
| Split expenses | CTH-20 |
| Debt repayment | CTH-28 |
| Savings contributions | CTH-10 |
| Budget reservations | CTH-22 |
| Settling unused budget | CTH-21, CTH-22, CTH-23 |
| Cash reconciliation | CTH-29 |
| Incorrect transactions and corrections | CTH-30 |
| Concurrent updates | CTH-31 |
| Duplicate submissions | CTH-32 |
| Salary arrival and allocation as in §1.1 | CTH-03 sampai CTH-09 |
| An installment payment | CTH-04, CTH-34 |
| A loan to extended family and its repayment | CTH-25 |
| One spouse borrowing back | CTH-27 |
| An asset purchase and its effect on net worth | CTH-33, CTH-34 |
| Tambahan: custody §5C | CTH-02 |
| Tambahan: alur acuan §5D per situasi custody | CTH-21 (A), CTH-22 (B dan D), CTH-23 (C) |
| Tambahan: kelebihan | CTH-24 |
| Tambahan: kontribusi dana pribadi tanpa dan dengan "minta diganti"; pelepasan hak ganti | CTH-13, CTH-14, CTH-16 |
| Tambahan: entri terlambat yang mengurangi selisih kas | CTH-29 |
| Tambahan: bantuan tidak rutin ke keluarga besar; penghapusan piutang | CTH-26, CTH-25 |
| Tambahan: penyisihan ke dana pribadi; hadiah pribadi | CTH-35, CTH-36 |
| Tambahan: gaji dengan bagian langsung ke dana pribadi (keputusan pemilik produk, 11 November 2026) | CTH-03b |
| Tambahan: beberapa penghasilan dalam satu periode, jatah pribadi bulanan, tarik tunai berbiaya, serah terima dari rekening, dua kiriman keluarga (keputusan dan pernyataan pemilik produk, 11 November 2026; akibat jatah pada ambang **[Usulan]**) | CTH-39 |
| Tambahan: perubahan syarat atas otorisasi terbuka | §7.6, varian CTH-22 |
| Tambahan: klaim antar-pemilik ditutup saat anggota keluar | CTH-38 |
| Tambahan: ringkasan bulanan, arus kas, anggaran | CTH-37 |

### 7.2 Saldo awal dan custody

#### CTH-01 — Saldo awal

Hari pertama mencatat (24 Oktober 2026). Motor bernilai 12000000 dengan sisa cicilan 6 × 1500000 = 9000000. Tiap kolom posisi adalah satu entri `opening`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Δ saldo awal | 500000 | 200000 | 300000 | 100000 | 5000000 | 6100000 | 0 | 12000000 | 9000000 | 0 | 0 | 9100000 | 400000 | 250000 |
| Sesudah | 500000 | 200000 | 300000 | 100000 | 5000000 | 6100000 | 0 | 12000000 | 9000000 | 0 | 0 | 9100000 | 400000 | 250000 |

Menguji: INV-02, INV-05, INV-07. Saldo awal bukan pemasukan dan tidak masuk laporan periode.

#### CTH-02 — Custody brief §5C (berdiri sendiri)

Kolom: `TUN-S` tunai suami, `TUN-I` tunai istri, `TTP-IBU` titipan pada ibu mertua untuk makan rumah tangga (jenis `entrusted_external`, pemilik `shared`), `BNK-B` rekening bersama (`holder = joint`), `UANG` = jumlah keempatnya.

| Baris | TUN-S | TUN-I | TTP-IBU | BNK-B | UANG |
|---|---|---|---|---|---|
| Sebelum | 5000000 | 3000000 | 0 | 2000000 | 10000000 |
| Δ suami menitipkan 1000000 kepada ibu mertua (`transfer`) | -1000000 | 0 | 1000000 | 0 | 0 |
| Sesudah (keadaan brief §5C) | 4000000 | 3000000 | 1000000 | 2000000 | 10000000 |
| Δ ibu mertua membeli lauk 150000, dicatat istri (`expense`) | 0 | 0 | -150000 | 0 | -150000 |
| Sesudah | 4000000 | 3000000 | 850000 | 2000000 | 9850000 |

Custody menampilkan empat baris: dipegang suami 4000000; dipegang istri 3000000; dititipkan pada ibu mertua 850000; dipegang bersama (rekening) 2000000. Titipan bukan pengeluaran sampai dibelanjakan. Menguji: INV-05, INV-06.

### 7.3 Gaji tiba dan dibagi (brief §1.1)

#### CTH-03 — Gaji tiba

Suami mencatat gaji 10000000 masuk `BNK-S`: `income`, sumber `salary`, `earner` suami, pemilik `shared` (bawaan).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 200000 | 300000 | 100000 | 5000000 | 6100000 | 0 | 12000000 | 9000000 | 0 | 0 | 9100000 | 400000 | 250000 |
| Δ gaji | 10000000 | 0 | 0 | 0 | 0 | 10000000 | 0 | 0 | 0 | 0 | 0 | 10000000 | 0 | 0 |
| Sesudah | 10500000 | 200000 | 300000 | 100000 | 5000000 | 16100000 | 0 | 12000000 | 9000000 | 0 | 0 | 19100000 | 400000 | 250000 |

Menguji: INV-02, INV-07.

#### CTH-03b — Gaji masuk, sebagian langsung ke dana pribadi (berdiri sendiri)

Varian CTH-03 menurut keputusan pemilik produk, 11 November 2026; `Sebelum`-nya sama dengan CTH-03 dan hasilnya **tidak** diteruskan ke CTH-04. Suami mencatat gaji 10000000 dan mengarahkan 1000000 ke dana pribadinya: 9000000 masuk `BNK-S`, 1000000 masuk `PRI-S`. Satu entri `income`, sumber `salary`, `earner` suami, dengan baris di dua buku. Bila seluruh gaji secara fisik tiba di satu rekening, bagian pribadinya dicatat pada tempat uang pribadi tersendiri (§2.3: satu tempat uang, satu pemilik); kolom `PRI-S` mewakilinya di sini.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 200000 | 300000 | 100000 | 5000000 | 6100000 | 0 | 12000000 | 9000000 | 0 | 0 | 9100000 | 400000 | 250000 |
| Δ gaji 10000000, 1000000 ke pribadi | 9000000 | 0 | 0 | 0 | 0 | 9000000 | 0 | 0 | 0 | 0 | 0 | 9000000 | 1000000 | 0 |
| Sesudah | 9500000 | 200000 | 300000 | 100000 | 5000000 | 15100000 | 0 | 12000000 | 9000000 | 0 | 0 | 18100000 | 1400000 | 250000 |

Baris entrinya:

| Buku | Baris | Nilai |
|---|---|---|
| Rumah tangga | Tempat uang `BNK-S` | 9000000 |
| Rumah tangga | Sebab `income` (gaji, atas nama suami) | 10000000 |
| Rumah tangga | Sebab `owner_transfer` (ke dana pribadi suami) | -1000000 |
| Pribadi suami | Tempat uang `PRI-S` | 1000000 |
| Pribadi suami | Sebab `owner_transfer` (dari dana bersama) | 1000000 |

Keseimbangan buku rumah tangga: 9000000 = 10000000 − 1000000. Buku pribadi suami: 1000000 = 1000000. Jumlah `owner_transfer`: −1000000 + 1000000 = 0. Istri melihat "gaji 10000000 masuk, 1000000 ke pribadi suami"; ia tidak melihat `PRI-S`, saldonya 1400000, atau pemakaiannya sesudah itu. Ringkasan bulanan: pemasukan 10000000, disisihkan ke dana pribadi suami 1000000, perubahan `KB` 9000000; arus kas: pemasukan 9000000. Dibanding CTH-03, `UANG` dan `KB` lebih kecil tepat 1000000 dan `PRI-S` lebih besar tepat 1000000: tidak ada rupiah yang hilang atau terhitung dua kali. Bagian pribadi 1000000 berada di atas ambang contoh 500000; apakah ia ditandai belum diputuskan dan menunggu pemilik produk di checkpoint S2 (**[Usulan]** dokumen ini dan dokumen 06 §6.2: ya, sebesar bagian yang di luar jatah pribadi suami, §3.4; T-13). Pencatat dan `earner` sama-sama suami, jadi entri ini sah dengan maupun tanpa INV-27; perintah yang sama dari istri (bagian pribadi ke `PRI-I` atas gaji atas nama suami) ditolak bila INV-27 diterima. Aturannya simetris: penghasilan istri dicatat istri dengan bentuk yang sama, `earner` istri, dan bagian pribadinya ke `PRI-I`. Menguji: INV-02, INV-03, INV-20, INV-24, INV-25, INV-27.

#### CTH-04 — Pembayaran cicilan

Cicilan motor 1500000 dibayar dari `BNK-S`: `repayment_out`, kategori "cicilan". Sisa 5 kali.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 10500000 | 200000 | 300000 | 100000 | 5000000 | 16100000 | 0 | 12000000 | 9000000 | 0 | 0 | 19100000 | 400000 | 250000 |
| Δ cicilan | -1500000 | 0 | 0 | 0 | 0 | -1500000 | 0 | 0 | -1500000 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 9000000 | 200000 | 300000 | 100000 | 5000000 | 14600000 | 0 | 12000000 | 7500000 | 0 | 0 | 19100000 | 400000 | 250000 |

Anggaran "cicilan" (batas 1500000): terpakai 1500000, sisa 0. Kekayaan bersih tidak berubah; uang keluar tampil di arus kas dan di belanja per kategori. Menguji: INV-02, INV-13, INV-15.

#### CTH-05 — Tagihan berulang

Tiga tagihan dibayar suami dari `BNK-S`, masing-masing `expense` yang ditautkan ke tagihan berulangnya, kategori "tagihan rutin": listrik 450000 (perkiraan 400000), wifi 350000, langganan kerja 200000.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 9000000 | 200000 | 300000 | 100000 | 5000000 | 14600000 | 0 | 12000000 | 7500000 | 0 | 0 | 19100000 | 400000 | 250000 |
| Δ listrik | -450000 | 0 | 0 | 0 | 0 | -450000 | 0 | 0 | 0 | 0 | 0 | -450000 | 0 | 0 |
| Sesudah | 8550000 | 200000 | 300000 | 100000 | 5000000 | 14150000 | 0 | 12000000 | 7500000 | 0 | 0 | 18650000 | 400000 | 250000 |
| Δ wifi | -350000 | 0 | 0 | 0 | 0 | -350000 | 0 | 0 | 0 | 0 | 0 | -350000 | 0 | 0 |
| Sesudah | 8200000 | 200000 | 300000 | 100000 | 5000000 | 13800000 | 0 | 12000000 | 7500000 | 0 | 0 | 18300000 | 400000 | 250000 |
| Δ langganan kerja | -200000 | 0 | 0 | 0 | 0 | -200000 | 0 | 0 | 0 | 0 | 0 | -200000 | 0 | 0 |
| Sesudah | 8000000 | 200000 | 300000 | 100000 | 5000000 | 13600000 | 0 | 12000000 | 7500000 | 0 | 0 | 18100000 | 400000 | 250000 |

Sebelum dibayar, tagihan yang jatuh tempo tidak mengubah angka apa pun (INV-12). Yang tercatat adalah jumlah sebenarnya, bukan perkiraan. Anggaran "tagihan rutin" (batas 1100000): terpakai 1000000, sisa 100000.

#### CTH-06 — Kiriman bulanan kepada ibu suami

Suami mentransfer 3000000 dari `BNK-S` kepada ibunya: `expense`, kategori "kiriman keluarga", penerima manfaat pihak luar "Ibu". Tidak dilacak sesudahnya.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 8000000 | 200000 | 300000 | 100000 | 5000000 | 13600000 | 0 | 12000000 | 7500000 | 0 | 0 | 18100000 | 400000 | 250000 |
| Δ kiriman ibu | -3000000 | 0 | 0 | 0 | 0 | -3000000 | 0 | 0 | 0 | 0 | 0 | -3000000 | 0 | 0 |
| Sesudah | 5000000 | 200000 | 300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |

**Andaian tentang ambang [Asumsi].** 3000000 di atas ambang contoh 500000. Contoh ini mengandaikan entri itu ditautkan ke tagihan berulang "kiriman Ibu" (§3.15) yang dirujuk satu *kesepakatan rutin* sebesar 3000000 per kali, disepakati berdua sebelum periode contoh (dokumen 06 §6.3, §7.4; dokumen 17 S2-4, default A), sehingga entri ini tidak ditandai. Rantai contoh tidak memuat kesepakatan itu karena kesepakatan tidak membuat entri dan tidak mengubah angka (INV-12). Tanpa kesepakatan itu, entri yang sama tersimpan dengan angka yang sama dan membawa *tanda di atas ambang* yang netral (dokumen 06 §6.2); tanda tidak mengubah angka mana pun di tabel ini.

#### CTH-07 — Tarik tunai (bank ke tunai)

Suami menarik 4500000 di ATM bank sendiri, tanpa biaya: `transfer` `BNK-S` → `TUN-S`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 5000000 | 200000 | 300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |
| Δ tarik tunai | -4500000 | 4500000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 500000 | 4700000 | 300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |

Menguji: INV-06.

#### CTH-08 — Suami menyerahkan uang kepada istri

Suami menyerahkan 4000000: `transfer` (serah terima) `TUN-S` → `TUN-I`, dicatat suami, bertanda belum dikonfirmasi sampai istri mengonfirmasi (tanpa batas waktu; istri dapat pula mengajukan keberatan kapan saja). Bila istri yang mencatatnya, tanda itu tidak ada dan angkanya sama.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 4700000 | 300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |
| Δ serah terima | 0 | -4000000 | 4000000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 500000 | 700000 | 4300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |

Custody sesudahnya: suami memegang 1200000 (`BNK-S` 500000 + `TUN-S` 700000); istri memegang 9400000 (`TUN-I` 4300000, di antaranya 4000000 belum dikonfirmasi, + `EWL-I` 100000 + `TAB-I` 5000000); jumlah 10600000. Pemilik uang tidak berubah: semuanya dana bersama. Menguji: INV-05, INV-06.

Dari gaji 10000000, rangkaian CTH-03 sampai CTH-08 mencatat: 1500000 cicilan, 1000000 tagihan, 3000000 kiriman ibu, 4000000 diserahkan kepada istri, 500000 tetap di suami; 1500000 + 1000000 + 3000000 + 4000000 + 500000 = 10000000. Tidak satu rupiah pun terhitung dua kali: `UANG` naik 10000000 lalu turun 5500000 (1500000 + 1000000 + 3000000), dan dua pemindahan tidak mengubahnya.

#### CTH-09 — Alokasi

Istri mengalokasikan 2500000 dari `TUN-I` untuk tujuan "belanja bulanan", dan 5000000 di `TAB-I` untuk tujuan tabungan "dana darurat" (target 10000000). Entri `allocation`; tidak ada saldo yang berubah, jadi tabel utama tidak berubah dari `Sesudah` CTH-08.

| Tempat uang | Saldo tercatat | Alokasi sebelum | Δ alokasi | Alokasi sesudah | Belum dialokasikan sesudah |
|---|---|---|---|---|---|
| TUN-I | 4300000 | 0 | 2500000 | 2500000 | 1800000 |
| TAB-I | 5000000 | 0 | 5000000 | 5000000 | 0 |

Menguji: INV-12, INV-17.

#### CTH-10 — Setoran tabungan

Istri menyetor 500000 dari bagian yang belum dialokasikan di `TUN-I` ke `TAB-I`: `transfer` yang membawa tujuan "dana darurat".

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 4300000 | 100000 | 5000000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |
| Δ setor tabungan | 0 | 0 | -500000 | 0 | 500000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 500000 | 700000 | 3800000 | 100000 | 5500000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |

| Tempat uang | Saldo tercatat | Alokasi sebelum | Δ alokasi | Alokasi sesudah | Belum dialokasikan sesudah |
|---|---|---|---|---|---|
| TUN-I | 3800000 | 2500000 | 0 | 2500000 | 1300000 |
| TAB-I | 5500000 | 5000000 | 500000 | 5500000 | 0 |

Kemajuan "dana darurat": 5500000 dari target 10000000. Menabung bukan pengeluaran: `UANG` dan `KB` tetap. Menguji: INV-06, INV-17.

### 7.4 Belanja, dana pribadi, dan hak ganti

#### CTH-11 — Belanja dari kas bersama

Istri membeli beras dan lauk 150000 dari `TUN-I`: `expense`, kategori "belanja bulanan", tujuan "belanja bulanan", penerima manfaat rumah tangga.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3800000 | 100000 | 5500000 | 10600000 | 0 | 12000000 | 7500000 | 0 | 0 | 15100000 | 400000 | 250000 |
| Δ belanja | 0 | 0 | -150000 | 0 | 0 | -150000 | 0 | 0 | 0 | 0 | 0 | -150000 | 0 | 0 |
| Sesudah | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 400000 | 250000 |

Anggaran "belanja bulanan" (batas 2500000): terpakai 150000, sisa 2350000. Alokasi "belanja bulanan" di `TUN-I`: 2500000 − 150000 = 2350000; belum dialokasikan 3650000 − 2350000 = 1300000.

Sesudah contoh ini istri melepas alokasi "belanja bulanan" (entri `allocation` −2350000; alokasi 0, belum dialokasikan `TUN-I` 3650000), karena bagi rumah tangga contoh anggaran sebagai batas sudah cukup (keputusan pemilik nomor 8). Contoh berikutnya tidak membawa tujuan; alokasi "dana darurat" 5500000 di `TAB-I` tidak berubah sampai akhir.

#### CTH-12 — Belanja pribadi dari uang pribadi

Suami membeli kopi 30000 dari `PRI-S`: `expense` di buku pribadinya.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 400000 | 250000 |
| Δ kopi | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | -30000 | 0 |
| Sesudah | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 370000 | 250000 |

Tidak ada angka rumah tangga yang berubah dan istri tidak melihat entri ini. Menguji: INV-20.

#### CTH-13 — Kontribusi dari dana pribadi (tidak minta diganti)

Suami membeli beras 80000 untuk rumah dengan `PRI-S` dan tidak menandai "minta diganti": `contribution_expense`. Buku pribadi: tempat uang −80000, `owner_transfer` −80000. Buku rumah tangga: `owner_transfer` +80000, `expense` −80000 (kategori "belanja bulanan", `funded_by` suami).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 370000 | 250000 |
| Δ beras pakai uang pribadi suami | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | -80000 | 0 |
| Sesudah | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 290000 | 250000 |

Keseimbangan buku rumah tangga: sisi posisi 0; sisi sebab +80000 − 80000 = 0. Istri melihat "beras 80000, pakai uang pribadi suami"; ia tidak melihat `PRI-S`. Anggaran "belanja bulanan": terpakai 150000 + 80000 = 230000, sisa 2270000. Laporan: belanja 80000 terhitung satu kali; kontribusi dari dana pribadi suami 80000. Menguji: INV-02, INV-03, INV-15, INV-20.

#### CTH-14 — Talangan dengan "minta diganti"

Suami membayar gas dan galon 60000 dengan `PRI-S` dan menandai "minta diganti": `advance`. Buku pribadi: tempat uang −60000, piutang (hak ganti) +60000. Buku rumah tangga: kewajiban `KWJ-G` +60000, `expense` −60000 (kategori "belanja bulanan").

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 0 | 14950000 | 290000 | 250000 |
| Δ talangan | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 60000 | -60000 | -60000 | 0 |
| Sesudah | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 60000 | 14890000 | 230000 | 250000 |

Kekayaan bersih pribadi suami tidak berubah: `PRI-S` 230000 + hak ganti 60000 = 290000. Anggaran "belanja bulanan": terpakai 290000, sisa 2210000. Menguji: INV-02, INV-08.

#### CTH-15 — Penggantian

Istri mengganti 60000 dari `TUN-I` kepada suami: `reimbursement`, dicatat suami sebagai penerimanya (§2.2). Baris yang mengurangi `TUN-I`, yang dipegang istri, bertanda belum dikonfirmasi sampai istri mengonfirmasi atau mengajukan keberatan (§3.3, sisi pembayar); saldonya sudah berubah.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3650000 | 100000 | 5500000 | 10450000 | 0 | 12000000 | 7500000 | 0 | 60000 | 14890000 | 230000 | 250000 |
| Δ penggantian | 0 | 0 | -60000 | 0 | 0 | -60000 | 0 | 0 | 0 | 0 | -60000 | 0 | 60000 | 0 |
| Sesudah | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 0 | 14890000 | 290000 | 250000 |

Dari CTH-14 dan CTH-15 bersama: kekayaan bersih rumah tangga turun tepat 60000, pengeluaran tercatat satu kali, anggaran tidak berubah oleh penggantian. Menguji: INV-02, INV-08, INV-13.

#### CTH-16 — Hak ganti dilepas

Istri membayar sabun 40000 dengan `PRI-I` dan menandai "minta diganti", lalu melepasnya (`claim_waiver`).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 0 | 14890000 | 290000 | 250000 |
| Δ talangan | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 40000 | -40000 | 0 | -40000 |
| Sesudah | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 40000 | 14850000 | 290000 | 210000 |
| Δ hak ganti dilepas | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | -40000 | 40000 | 0 | 0 |
| Sesudah | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 0 | 14890000 | 290000 | 210000 |

Kenaikan 40000 pada pelepasan punya baris sebab `owner_transfer` +40000 di buku rumah tangga dan tampil sebagai "kontribusi dari dana pribadi istri". Keadaan akhirnya sama dengan bila sejak awal dicatat sebagai kontribusi (bandingkan CTH-13). Anggaran "belanja bulanan": terpakai 330000, sisa 2170000. Menguji: INV-02, INV-03, INV-08.

#### CTH-17 — Membayarkan keperluan anggota lain dari dana bersama

Istri membayar potong rambut suami 50000 dari `TUN-I`: `expense`, kategori "perawatan diri", penerima manfaat suami, pembayar istri. Tidak ada klaim: dana bersama memang untuk keperluan anggota.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3590000 | 100000 | 5500000 | 10390000 | 0 | 12000000 | 7500000 | 0 | 0 | 14890000 | 290000 | 210000 |
| Δ potong rambut | 0 | 0 | -50000 | 0 | 0 | -50000 | 0 | 0 | 0 | 0 | 0 | -50000 | 0 | 0 |
| Sesudah | 500000 | 700000 | 3540000 | 100000 | 5500000 | 10340000 | 0 | 12000000 | 7500000 | 0 | 0 | 14840000 | 290000 | 210000 |

#### CTH-17b — Membayarkan keperluan pribadi anggota lain dari dana pribadi (berdiri sendiri)

Varian: istri membayar 50000 keperluan pribadi suami dengan `PRI-I` dan minta diganti; suami kemudian mengganti dari `PRI-S`. Buku rumah tangga tidak tersentuh. Kolom: `KLAIM` = piutang istri = kewajiban suami; `KB-I = PRI-I + KLAIM`; `KB-S = PRI-S − KLAIM`.

| Baris | PRI-I | PRI-S | KLAIM | KB-I | KB-S |
|---|---|---|---|---|---|
| Sebelum | 210000 | 290000 | 0 | 210000 | 290000 |
| Δ istri menalangi (`advance`) | -50000 | 0 | 50000 | 0 | -50000 |
| Sesudah | 160000 | 290000 | 50000 | 210000 | 240000 |
| Δ suami mengganti (`reimbursement`) | 50000 | -50000 | -50000 | 0 | 0 |
| Sesudah | 210000 | 240000 | 0 | 210000 | 240000 |

Pengeluaran 50000 tercatat satu kali, di buku pribadi suami. Keduanya melihat klaim ini (kebocoran B3). Contoh ini contoh tingkat model: kedua entrinya menyentuh buku pribadi dua anggota, sehingga alur pencatatannya ditunda dari MVP (§2.2, T-12); angkanya tetap menjadi test INV-02 dan INV-08. Menguji: INV-02, INV-08.

### 7.5 Pemindahan, pengembalian dana, pengeluaran terbagi

#### CTH-18 — Bank ke dompet digital dengan biaya admin, lalu tunai ke bank

Suami mengisi `EWL-I` 300000 dari `BNK-S` dengan biaya admin 1000 (satu entri `transfer` dengan baris `expense` 1000, kategori "biaya admin"). Lalu ia menyetor tunai 200000 dari `TUN-S` ke `BNK-S`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 700000 | 3540000 | 100000 | 5500000 | 10340000 | 0 | 12000000 | 7500000 | 0 | 0 | 14840000 | 290000 | 210000 |
| Δ isi dompet digital + biaya admin | -301000 | 0 | 0 | 300000 | 0 | -1000 | 0 | 0 | 0 | 0 | 0 | -1000 | 0 | 0 |
| Sesudah | 199000 | 700000 | 3540000 | 400000 | 5500000 | 10339000 | 0 | 12000000 | 7500000 | 0 | 0 | 14839000 | 290000 | 210000 |
| Δ setor tunai | 200000 | -200000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 400000 | 5500000 | 10339000 | 0 | 12000000 | 7500000 | 0 | 0 | 14839000 | 290000 | 210000 |

Hanya biaya admin yang menjadi pengeluaran. Baris +300000 pada `EWL-I` dicatat suami ke tempat uang yang dipegang istri, jadi ia bertanda belum dikonfirmasi sampai istri mengonfirmasi (§3.3); saldonya sudah berubah. Menguji: INV-06.

#### CTH-19 — Pengembalian dana sebagian dan penuh

Istri membeli panci 120000 lewat `EWL-I` (kategori "perlengkapan rumah"); penjual mengembalikan 20000. Lalu ia membeli blender 100000, membatalkannya, dan menerima 100000 kembali.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 500000 | 3540000 | 400000 | 5500000 | 10339000 | 0 | 12000000 | 7500000 | 0 | 0 | 14839000 | 290000 | 210000 |
| Δ beli panci | 0 | 0 | 0 | -120000 | 0 | -120000 | 0 | 0 | 0 | 0 | 0 | -120000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 280000 | 5500000 | 10219000 | 0 | 12000000 | 7500000 | 0 | 0 | 14719000 | 290000 | 210000 |
| Δ pengembalian sebagian | 0 | 0 | 0 | 20000 | 0 | 20000 | 0 | 0 | 0 | 0 | 0 | 20000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 300000 | 5500000 | 10239000 | 0 | 12000000 | 7500000 | 0 | 0 | 14739000 | 290000 | 210000 |
| Δ beli blender | 0 | 0 | 0 | -100000 | 0 | -100000 | 0 | 0 | 0 | 0 | 0 | -100000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 200000 | 5500000 | 10139000 | 0 | 12000000 | 7500000 | 0 | 0 | 14639000 | 290000 | 210000 |
| Δ pengembalian penuh | 0 | 0 | 0 | 100000 | 0 | 100000 | 0 | 0 | 0 | 0 | 0 | 100000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3540000 | 300000 | 5500000 | 10239000 | 0 | 12000000 | 7500000 | 0 | 0 | 14739000 | 290000 | 210000 |

"Perlengkapan rumah" terpakai: 120000 − 20000 + 100000 − 100000 = 100000. Tidak ada pemasukan yang tercatat. Pengembalian ketiga atas panci sebesar lebih dari 100000 akan ditolak (INV-13).

#### CTH-20 — Pengeluaran terbagi

Istri membayar 200000 di minimarket dari `TUN-I`: satu entri `expense` dengan dua baris sebab, 150000 "belanja bulanan" dan 50000 "perlengkapan rumah".

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 500000 | 3540000 | 300000 | 5500000 | 10239000 | 0 | 12000000 | 7500000 | 0 | 0 | 14739000 | 290000 | 210000 |
| Δ belanja terbagi | 0 | 0 | -200000 | 0 | 0 | -200000 | 0 | 0 | 0 | 0 | 0 | -200000 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3340000 | 300000 | 5500000 | 10039000 | 0 | 12000000 | 7500000 | 0 | 0 | 14539000 | 290000 | 210000 |

150000 + 50000 = 200000. "Belanja bulanan": terpakai 480000, sisa 2020000. "Perlengkapan rumah": terpakai 150000. Pembagian menurut pembayar memakai bentuk yang sama: beberapa baris tempat uang dalam satu entri. Menguji: INV-02, INV-15.

### 7.6 Permintaan, otorisasi, penyelesaian

Kategori "jajan dan makan di luar" beranggaran 500000.

#### CTH-21 — Alur acuan §5D, situasi A: pemohon memegang uangnya

Suami meminta 100000 untuk membeli makanan; istri menyetujui; ia membeli 75000 dari `TUN-S`; penyelesaian.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 500000 | 3340000 | 300000 | 5500000 | 10039000 | 0 | 12000000 | 7500000 | 0 | 0 | 14539000 | 290000 | 210000 |
| Δ disetujui (otorisasi 100000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 500000 | 3340000 | 300000 | 5500000 | 10039000 | 0 | 12000000 | 7500000 | 0 | 0 | 14539000 | 290000 | 210000 |
| Δ beli makanan 75000 | 0 | -75000 | 0 | 0 | 0 | -75000 | 0 | 0 | 0 | 0 | 0 | -75000 | 0 | 0 |
| Sesudah | 399000 | 425000 | 3340000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |
| Δ penyelesaian (sisa 25000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 425000 | 3340000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |

| Saat | Batas anggaran | Terpakai | Terikat | Sisa |
|---|---|---|---|---|
| Sesudah disetujui | 500000 | 0 | 100000 | 400000 |
| Sesudah beli | 500000 | 75000 | 25000 | 400000 |
| Sesudah penyelesaian | 500000 | 75000 | 0 | 425000 |

"Benar" di situasi A: sisa 25000 tidak pernah berpindah, tetap dana bersama di `TUN-S`; penyelesaian hanya melepas terikat. Menguji: INV-12, INV-14, INV-15.

#### CTH-22 — Situasi B dengan reservasi: uang diserahkan untuk permintaan

Istri meminta 100000 untuk kue arisan dan menyebut `TUN-S` sebagai sumber; suami menyetujui (lahir reservasi 100000 di `TUN-S`); suami menyerahkan 100000; istri membeli 75000; istri mengembalikan sisa 25000.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 425000 | 3340000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |
| Δ disetujui (otorisasi dan reservasi 100000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 425000 | 3340000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |
| Δ serah terima 100000 | 0 | -100000 | 100000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 325000 | 3440000 | 300000 | 5500000 | 9964000 | 0 | 12000000 | 7500000 | 0 | 0 | 14464000 | 290000 | 210000 |
| Δ beli kue 75000 | 0 | 0 | -75000 | 0 | 0 | -75000 | 0 | 0 | 0 | 0 | 0 | -75000 | 0 | 0 |
| Sesudah | 399000 | 325000 | 3365000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 0 | 14389000 | 290000 | 210000 |
| Δ penyelesaian: sisa 25000 dikembalikan | 0 | 25000 | -25000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 0 | 14389000 | 290000 | 210000 |

| Saat | TUN-S tercatat | Reservasi aktif | TUN-S tersedia |
|---|---|---|---|
| Sebelum | 425000 | 0 | 425000 |
| Sesudah disetujui | 425000 | 100000 | 325000 |
| Sesudah serah terima | 325000 | 0 | 325000 |
| Sesudah penyelesaian | 350000 | 0 | 350000 |

| Saat | Batas anggaran | Terpakai | Terikat | Sisa |
|---|---|---|---|---|
| Sesudah disetujui | 500000 | 75000 | 100000 | 325000 |
| Sesudah beli | 500000 | 150000 | 25000 | 325000 |
| Sesudah penyelesaian | 500000 | 150000 | 0 | 350000 |

"Benar" di situasi B: sisa 25000 berada di tangan pemohon sebagai dana bersama. Pilihan (1), dipakai di atas: dikembalikan lewat serah terima. Pilihan (2): tetap dipegang pemohon, tanpa entri; keadaan akhirnya `TUN-S` 325000, `TUN-I` 3365000, `UANG` 9889000. Pada kedua pilihan `UANG` dan `KB` sama. Menguji: INV-06, INV-11, INV-12, INV-14.

**Varian situasi D (berdiri sendiri).** Bila suami, pemegang `TUN-S`, membayar kue 75000 langsung dari `TUN-S` atas permintaan yang sama:

| Saat | TUN-S tercatat | Reservasi aktif | TUN-S tersedia |
|---|---|---|---|
| Sesudah disetujui | 425000 | 100000 | 325000 |
| Sesudah membayar 75000 | 350000 | 25000 | 325000 |
| Sesudah penyelesaian | 350000 | 0 | 350000 |

Sisa 25000 tidak pernah meninggalkan `TUN-S`; penyelesaian hanya melepas reservasinya.

**Varian perubahan syarat (berdiri sendiri).** Bila sesudah disetujui, dan sebelum ada serah terima atau belanja, istri mengusulkan batas 150000 dan suami menerimanya (`authorization.amended`): tidak ada entri dan tidak ada saldo tercatat yang berubah; reservasi dan *terikat* mengikuti batas baru.

| Saat | Batas otorisasi | TUN-S tercatat | Reservasi aktif | TUN-S tersedia | Terpakai anggaran | Terikat | Sisa anggaran |
|---|---|---|---|---|---|---|---|
| Sesudah disetujui | 100000 | 425000 | 100000 | 325000 | 75000 | 100000 | 325000 |
| Sesudah perubahan syarat diterima | 150000 | 425000 | 150000 | 275000 | 75000 | 150000 | 275000 |

Batas anggaran 500000: 500000 − 75000 − 150000 = 275000. `TUN-S` tersedia: 425000 − 150000 = 275000. Menguji: INV-11, INV-12, INV-14.

#### CTH-23 — Situasi C: dibayar dengan dana pribadi atas otorisasi dana bersama

Suami meminta 100000 untuk obat (kategori "kesehatan"); istri menyetujui; ia membayar 75000 dengan `PRI-S` (tanda "minta diganti" hidup secara bawaan karena tertaut ke otorisasi dana bersama); penyelesaian; istri mengganti dari `TUN-I`. Talangan dan penggantiannya dicatat suami, karena keduanya menyentuh buku pribadinya (§2.2); pada penggantian, baris yang mengurangi `TUN-I` bertanda belum dikonfirmasi sampai istri, pemegangnya, mengonfirmasi (§3.3, sisi pembayar).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 0 | 14389000 | 290000 | 210000 |
| Δ disetujui (otorisasi 100000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 0 | 14389000 | 290000 | 210000 |
| Δ talangan 75000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 75000 | -75000 | -75000 | 0 |
| Sesudah | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 75000 | 14314000 | 215000 | 210000 |
| Δ penyelesaian (sisa 25000) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 350000 | 3340000 | 300000 | 5500000 | 9889000 | 0 | 12000000 | 7500000 | 0 | 75000 | 14314000 | 215000 | 210000 |
| Δ penggantian 75000 | 0 | 0 | -75000 | 0 | 0 | -75000 | 0 | 0 | 0 | 0 | -75000 | 0 | 75000 | 0 |
| Sesudah | 399000 | 350000 | 3265000 | 300000 | 5500000 | 9814000 | 0 | 12000000 | 7500000 | 0 | 0 | 14314000 | 290000 | 210000 |

"Benar" di situasi C: hak ganti sebesar realisasi 75000, bukan sebesar batas 100000; sisa 25000 tidak pernah ada sebagai uang dan hanya dilepas dari terikat. Menguji: INV-08, INV-12, INV-14.

#### CTH-24 — Kelebihan

Istri mendapat otorisasi 200000 untuk hadiah ulang tahun keponakan (kategori "hadiah") dan membayar 230000 dari `TUN-I`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 350000 | 3265000 | 300000 | 5500000 | 9814000 | 0 | 12000000 | 7500000 | 0 | 0 | 14314000 | 290000 | 210000 |
| Δ beli hadiah 230000 | 0 | 0 | -230000 | 0 | 0 | -230000 | 0 | 0 | 0 | 0 | 0 | -230000 | 0 | 0 |
| Sesudah | 399000 | 350000 | 3035000 | 300000 | 5500000 | 9584000 | 0 | 12000000 | 7500000 | 0 | 0 | 14084000 | 290000 | 210000 |

Realisasi 230000; batas 200000; kelebihan 30000; terikat max(0, 200000 − 230000) = 0. Pencatatan tidak ditolak dan tidak ada entri tambahan. Kelebihan 30000 dibayar dari dana bersama, jadi ia pengeluaran dana bersama biasa tanpa klaim; permintaan menampilkan "melebihi 30000" (kait `authorization.overrun`). Menguji: INV-14.

### 7.7 Keluarga besar, utang, dan pinjam kembali

#### CTH-25 — Pinjaman kepada adik suami, pelunasan sebagian, penghapusan sebagian

Istri menyerahkan 500000 dari `TUN-I` kepada adik suami sebagai pinjaman (`loan_given`; tidak di atas ambang 500000, jadi tanpa permintaan). Adik membayar 200000 kepada suami (`repayment_in` ke `TUN-S`). Kemudian suami mengikhlaskan 100000 (`write_off`, dengan alasan).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 350000 | 3035000 | 300000 | 5500000 | 9584000 | 0 | 12000000 | 7500000 | 0 | 0 | 14084000 | 290000 | 210000 |
| Δ pinjaman diberikan | 0 | 0 | -500000 | 0 | 0 | -500000 | 500000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 350000 | 2535000 | 300000 | 5500000 | 9084000 | 500000 | 12000000 | 7500000 | 0 | 0 | 14084000 | 290000 | 210000 |
| Δ adik membayar 200000 | 0 | 200000 | 0 | 0 | 0 | 200000 | -200000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 550000 | 2535000 | 300000 | 5500000 | 9284000 | 300000 | 12000000 | 7500000 | 0 | 0 | 14084000 | 290000 | 210000 |
| Δ 100000 diikhlaskan | 0 | 0 | 0 | 0 | 0 | 0 | -100000 | 0 | 0 | 0 | 0 | -100000 | 0 | 0 |
| Sesudah | 399000 | 550000 | 2535000 | 300000 | 5500000 | 9284000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13984000 | 290000 | 210000 |

Pinjaman dan pelunasannya bukan pengeluaran dan bukan pemasukan. Hanya penghapusan yang menurunkan kekayaan bersih, sebagai baris "piutang diikhlaskan". Sisa piutang 200000. Menguji: INV-02, INV-13.

#### CTH-26 — Bantuan tidak rutin kepada ayah suami

Suami memberi 200000 dari `TUN-S` kepada ayahnya dan memilih "bantuan", bukan pinjaman: `expense`, kategori "kiriman keluarga".

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 550000 | 2535000 | 300000 | 5500000 | 9284000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13984000 | 290000 | 210000 |
| Δ bantuan ayah | 0 | -200000 | 0 | 0 | 0 | -200000 | 0 | 0 | 0 | 0 | 0 | -200000 | 0 | 0 |
| Sesudah | 399000 | 350000 | 2535000 | 300000 | 5500000 | 9084000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13784000 | 290000 | 210000 |

#### CTH-27 — Pinjam kembali

Suami mengambil lagi 300000 dari uang bersama yang dipegang istri: `transfer` (serah terima) `TUN-I` → `TUN-S`. Bukan utang.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 350000 | 2535000 | 300000 | 5500000 | 9084000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13784000 | 290000 | 210000 |
| Δ pinjam kembali | 0 | 300000 | -300000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 650000 | 2235000 | 300000 | 5500000 | 9084000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13784000 | 290000 | 210000 |

Tidak ada kewajiban atau piutang yang lahir; yang berubah hanya siapa memegang. Menguji: INV-06.

#### CTH-28 — Utang rumah tangga dan pelunasannya

Rumah tangga meminjam 1000000 dari kakak istri, diterima istri ke `TUN-I` (`loan_received`); lalu istri membayar kembali 400000 (`repayment_out`).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 650000 | 2235000 | 300000 | 5500000 | 9084000 | 200000 | 12000000 | 7500000 | 0 | 0 | 13784000 | 290000 | 210000 |
| Δ pinjaman diterima | 0 | 0 | 1000000 | 0 | 0 | 1000000 | 0 | 0 | 0 | 1000000 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 650000 | 3235000 | 300000 | 5500000 | 10084000 | 200000 | 12000000 | 7500000 | 1000000 | 0 | 13784000 | 290000 | 210000 |
| Δ pelunasan 400000 | 0 | 0 | -400000 | 0 | 0 | -400000 | 0 | 0 | 0 | -400000 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 650000 | 2835000 | 300000 | 5500000 | 9684000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13784000 | 290000 | 210000 |

Uang pinjaman menaikkan `UANG` tanpa menjadi pemasukan; pelunasan menurunkannya tanpa menjadi pengeluaran. Menguji: INV-02, INV-13.

### 7.8 Rekonsiliasi, koreksi, bersamaan, kiriman ganda

#### CTH-29 — Rekonsiliasi kas dan entri terlambat

Tanggal 10 November istri menghitung `TUN-I`: terhitung 2750000, tercatat 2835000, selisih −85000. Tanggal 12 November ia teringat membeli sayur 60000 pada 8 November dan mencatatnya; aplikasi menawarkannya sebagai bagian dari selisih itu (`recategorization`, kategori "belanja bulanan", tanggal kejadian 8 November). Tanggal 13 November suami menghitung `TUN-S`: terhitung 660000, tercatat 650000, selisih +10000.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 650000 | 2835000 | 300000 | 5500000 | 9684000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13784000 | 290000 | 210000 |
| Δ selisih kas TUN-I | 0 | 0 | -85000 | 0 | 0 | -85000 | 0 | 0 | 0 | 0 | 0 | -85000 | 0 | 0 |
| Sesudah | 399000 | 650000 | 2750000 | 300000 | 5500000 | 9599000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13699000 | 290000 | 210000 |
| Δ entri terlambat sayur 60000 (menutup selisih) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 650000 | 2750000 | 300000 | 5500000 | 9599000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13699000 | 290000 | 210000 |
| Δ selisih kas TUN-S | 0 | 10000 | 0 | 0 | 0 | 10000 | 0 | 0 | 0 | 0 | 0 | 10000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2750000 | 300000 | 5500000 | 9609000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13709000 | 290000 | 210000 |

| Rekonsiliasi | Terhitung | Tercatat | Selisih | Dirinci | Belum dirinci |
|---|---|---|---|---|---|
| TUN-I, sesudah dihitung | 2750000 | 2835000 | -85000 | 0 | 85000 |
| TUN-I, sesudah entri terlambat | 2750000 | 2835000 | -85000 | 60000 | 25000 |
| TUN-S, sesudah dihitung | 660000 | 650000 | 10000 | 0 | 10000 |

Entri terlambat berisi dua baris sebab: `cash_difference` +60000 dan `expense` −60000; jumlahnya 0, jadi tidak ada saldo yang berubah dan uangnya tidak terhitung dua kali. "Belanja bulanan": terpakai 480000 + 60000 = 540000, sisa 1960000. Selisih positif 10000 bukan pemasukan. Tidak ada entri lama yang disunting. Menguji: INV-02, INV-09, INV-16.

#### CTH-30 — Catatan salah dan koreksinya

Istri mencatat parkir 50000 dari `TUN-I`, padahal 5000. Ia memperbaikinya: entri pembalik, lalu entri yang benar (kategori "transportasi").

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 660000 | 2750000 | 300000 | 5500000 | 9609000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13709000 | 290000 | 210000 |
| Δ parkir, salah (50000) | 0 | 0 | -50000 | 0 | 0 | -50000 | 0 | 0 | 0 | 0 | 0 | -50000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2700000 | 300000 | 5500000 | 9559000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13659000 | 290000 | 210000 |
| Δ entri pembalik | 0 | 0 | 50000 | 0 | 0 | 50000 | 0 | 0 | 0 | 0 | 0 | 50000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2750000 | 300000 | 5500000 | 9609000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13709000 | 290000 | 210000 |
| Δ parkir, benar (5000) | 0 | 0 | -5000 | 0 | 0 | -5000 | 0 | 0 | 0 | 0 | 0 | -5000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2745000 | 300000 | 5500000 | 9604000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13704000 | 290000 | 210000 |

Ketiga entri tetap ada di jejak audit. "Transportasi" terpakai: 50000 − 50000 + 5000 = 5000. Hasil yang sama dapat dicapai dengan satu entri penyesuaian +45000. Menguji: INV-09, INV-18.

#### CTH-31 — Pembaruan bersamaan

Pada saat yang sama, istri mencatat ojek 20000 dari `TUN-I` (kategori "transportasi") dan suami mencatat serah terima 100000 dari `TUN-S` ke `TUN-I`. Kedua perangkat membaca `TUN-I` = 2745000.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 660000 | 2745000 | 300000 | 5500000 | 9604000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13704000 | 290000 | 210000 |
| Δ ojek (istri) | 0 | 0 | -20000 | 0 | 0 | -20000 | 0 | 0 | 0 | 0 | 0 | -20000 | 0 | 0 |
| Sesudah | 399000 | 660000 | 2725000 | 300000 | 5500000 | 9584000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13684000 | 290000 | 210000 |
| Δ serah terima (suami) | 0 | -100000 | 100000 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 560000 | 2825000 | 300000 | 5500000 | 9584000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13684000 | 290000 | 210000 |

Hasil yang benar untuk `TUN-I` adalah 2745000 − 20000 + 100000 = 2825000, apa pun urutan tibanya. Hasil 2725000 atau 2845000 berarti satu entri hilang (lost update) dan merupakan cacat. Model menjaminnya karena entri hanya ditambahkan dan saldo adalah jumlah baris, bukan nilai yang ditimpa klien. Menguji: INV-04, INV-23.

#### CTH-32 — Kiriman ganda

Suami mencatat bensin 40000 dari `TUN-S` (kategori "transportasi"). Jaringan lambat; ia mengetuk "simpan" dua kali, dan aplikasi mengirim ulang dengan `idempotency_key` yang sama.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 560000 | 2825000 | 300000 | 5500000 | 9584000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13684000 | 290000 | 210000 |
| Δ bensin, kiriman pertama | 0 | -40000 | 0 | 0 | 0 | -40000 | 0 | 0 | 0 | 0 | 0 | -40000 | 0 | 0 |
| Sesudah | 399000 | 520000 | 2825000 | 300000 | 5500000 | 9544000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13644000 | 290000 | 210000 |
| Δ bensin, kiriman kedua (kunci sama) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 520000 | 2825000 | 300000 | 5500000 | 9544000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13644000 | 290000 | 210000 |

Jumlah entri bertambah satu, bukan dua; kiriman kedua mengembalikan entri pertama. Menguji: INV-10.

### 7.9 Aset, cicilan baru, penyisihan, hadiah

#### CTH-33 — Pembelian aset dan penilaian ulang

Istri membeli laptop bekas 2000000 dari `TUN-I` (`asset_purchase`; di atas ambang, jadi diajukan lebih dulu sebagai permintaan). Kemudian nilai motor diperbarui dari 12000000 menjadi 11000000 (`asset_valuation` −1000000).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 520000 | 2825000 | 300000 | 5500000 | 9544000 | 200000 | 12000000 | 7500000 | 600000 | 0 | 13644000 | 290000 | 210000 |
| Δ beli laptop | 0 | 0 | -2000000 | 0 | 0 | -2000000 | 0 | 2000000 | 0 | 0 | 0 | 0 | 0 | 0 |
| Sesudah | 399000 | 520000 | 825000 | 300000 | 5500000 | 7544000 | 200000 | 14000000 | 7500000 | 600000 | 0 | 13644000 | 290000 | 210000 |
| Δ nilai motor turun | 0 | 0 | 0 | 0 | 0 | 0 | 0 | -1000000 | 0 | 0 | 0 | -1000000 | 0 | 0 |
| Sesudah | 399000 | 520000 | 825000 | 300000 | 5500000 | 7544000 | 200000 | 13000000 | 7500000 | 600000 | 0 | 12644000 | 290000 | 210000 |

Pembelian aset tidak mengubah kekayaan bersih: uang berubah bentuk menjadi barang. Ia tampil di arus kas sebagai "beli aset". Menguji: INV-02, INV-07.

**Yang tidak ditampilkan rantai [Asumsi].** "Diajukan lebih dulu" berarti: sebelum membeli, istri mengajukan permintaan; suami menyetujuinya; lahir otorisasi atas dana bersama dengan batas tidak kurang dari 2000000, tanpa tempat uang sumber, jadi tanpa reservasi; entri pembelian ditautkan kepadanya, sehingga realisasinya 2000000 dan entri itu tidak dibandingkan dengan ambang (§3.8); lalu otorisasinya diselesaikan pada situasi A, tanpa entri (seperti CTH-21). Rantai ini tidak memuat baris persetujuan dan penyelesaian karena keduanya tidak membuat entri dan tidak mengubah kolom mana pun (INV-12); bila ditambahkan, barisnya nol seluruhnya. Sesudah penyelesaian *terikat*-nya nol, jadi CTH-37 tidak tersentuh. Bila istri membeli tanpa mengajukan, entri yang sama tersimpan dengan angka yang sama dan membawa tanda netral di atas ambang (dokumen 06 §6.2; dokumen 17 S2-2).

#### CTH-34 — Pembelian aset secara cicilan dan biaya kredit

Kulkas seharga 4000000 dibeli dengan uang muka 400000 dari `TUN-I` dan 12 cicilan × 330000 = 3960000. Biaya kredit = 400000 + 3960000 − 4000000 = 360000. Satu entri `asset_purchase`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 520000 | 825000 | 300000 | 5500000 | 7544000 | 200000 | 13000000 | 7500000 | 600000 | 0 | 12644000 | 290000 | 210000 |
| Δ beli kulkas secara cicilan | 0 | 0 | -400000 | 0 | 0 | -400000 | 0 | 4000000 | 3960000 | 0 | 0 | -360000 | 0 | 0 |
| Sesudah | 399000 | 520000 | 425000 | 300000 | 5500000 | 7144000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12284000 | 290000 | 210000 |

Keseimbangan: −400000 + 4000000 − 3960000 = −360000 = `credit_cost`. `KWJ-C` 11460000 = motor 7500000 + kulkas 3960000. Kekayaan bersih turun sebesar biaya kredit pada tanggal pembelian, dengan baris sebab yang menjelaskannya (menutup K-11). Jumlah pemakaian entri ini, yang dibandingkan dengan ambang permintaan dan menjadi realisasi otorisasinya, adalah 400000 + 3960000 = 4360000 (§3.8): di atas ambang contoh 500000, jadi pembelian ini diajukan lebih dulu, dengan andaian yang sama dengan CTH-33 (otorisasi tanpa reservasi, batas tidak kurang dari jumlah pemakaian itu, diselesaikan tanpa entri; tidak ditampilkan rantai). Dua belas pembayaran 330000 berikutnya masing-masing seperti CTH-04: terhitung terhadap anggaran "cicilan", tanpa mengubah kekayaan bersih. Menguji: INV-02, INV-07.

#### CTH-35 — Penyisihan ke dana pribadi

Uang saku bulanan yang disepakati: 100000 untuk masing-masing. Suami memindahkan 100000 dari `TUN-S` ke `PRI-S`; istri 100000 dari `TUN-I` ke `PRI-I`. Dua entri `owner_transfer`.

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 520000 | 425000 | 300000 | 5500000 | 7144000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12284000 | 290000 | 210000 |
| Δ uang saku suami | 0 | -100000 | 0 | 0 | 0 | -100000 | 0 | 0 | 0 | 0 | 0 | -100000 | 100000 | 0 |
| Sesudah | 399000 | 420000 | 425000 | 300000 | 5500000 | 7044000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12184000 | 390000 | 210000 |
| Δ uang saku istri | 0 | 0 | -100000 | 0 | 0 | -100000 | 0 | 0 | 0 | 0 | 0 | -100000 | 0 | 100000 |
| Sesudah | 399000 | 420000 | 325000 | 300000 | 5500000 | 6944000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12084000 | 390000 | 310000 |

Keduanya melihat kedua entri ini (kebocoran B1); tidak seorang pun melihat saldo pribadi pasangannya sesudahnya. Menguji: INV-02, INV-03.

#### CTH-36 — Hadiah untuk satu orang

Istri menerima hadiah 500000 dari ibunya dan mencatatnya ke `PRI-I`: `income` di buku pribadinya, sumber `gift` (bawaan: dana pribadi).

| Baris | BNK-S | TUN-S | TUN-I | EWL-I | TAB-I | UANG | PIU | AST | KWJ-C | KWJ-U | KWJ-G | KB | PRI-S | PRI-I |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Sebelum | 399000 | 420000 | 325000 | 300000 | 5500000 | 6944000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12084000 | 390000 | 310000 |
| Δ hadiah untuk istri | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 500000 |
| Sesudah | 399000 | 420000 | 325000 | 300000 | 5500000 | 6944000 | 200000 | 17000000 | 11460000 | 600000 | 0 | 12084000 | 390000 | 810000 |

Suami tidak melihat entri ini (**[Usulan]**, menunggu pemilik produk di checkpoint S2; §3.1) dan angka rumah tangga tidak berubah. Menguji: INV-20.

### 7.10 CTH-37 — Ringkasan bulanan rumah tangga contoh

Dihitung dari CTH-03 sampai CTH-36 (CTH-01 adalah saldo awal dan tidak masuk periode). Semua angka dari buku rumah tangga.

**Custody akhir** (kriteria sukses nomor 1).

| Pemegang | Tempat uang | Saldo |
|---|---|---|
| Suami | BNK-S | 399000 |
| Suami | TUN-S | 420000 |
| Istri | TUN-I | 325000 |
| Istri | EWL-I | 300000 |
| Istri | TAB-I | 5500000 |
| Jumlah = UANG | | 6944000 |

Suami memegang 819000; istri memegang 6125000; 819000 + 6125000 = 6944000.

**Tabel A — mengapa harta bersih berubah.**

| Baris | Jumlah |
|---|---|
| Pemasukan (gaji, atas nama suami) | 10000000 |
| Belanja semua kategori (lihat tabel B, tanpa cicilan) | -5461000 |
| Biaya kredit | -360000 |
| Selisih hitung negatif yang belum dirinci | -25000 |
| Selisih hitung positif | 10000 |
| Piutang diikhlaskan | -100000 |
| Perubahan nilai aset | -1000000 |
| Kontribusi dari dana pribadi (suami 80000, istri 40000) | 120000 |
| Disisihkan ke dana pribadi (suami 100000, istri 100000) | -200000 |
| Jumlah = perubahan KB | 2984000 |

`KB` awal 9100000 (CTH-01) + 2984000 = 12084000 = `KB` akhir (CTH-36).

**Tabel B — uang keluar per kategori.**

| Kategori | Rincian | Terpakai |
|---|---|---|
| Tagihan rutin | 450000 + 350000 + 200000 | 1000000 |
| Kiriman keluarga | 3000000 + 200000 | 3200000 |
| Belanja bulanan | 150000 + 80000 + 60000 + 40000 + 150000 + 60000 | 540000 |
| Perawatan diri | 50000 | 50000 |
| Biaya admin | 1000 | 1000 |
| Perlengkapan rumah | 120000 − 20000 + 100000 − 100000 + 50000 | 150000 |
| Jajan dan makan di luar | 75000 + 75000 | 150000 |
| Kesehatan | 75000 | 75000 |
| Hadiah | 230000 | 230000 |
| Transportasi | 5000 + 20000 + 40000 | 65000 |
| Jumlah belanja | | 5461000 |
| Cicilan | 1500000 | 1500000 |
| Jumlah tabel B | | 6961000 |

**Tabel C — siapa membayar apa.**

| Pembayar | Dari dana bersama | Dana pribadi, kontribusi | Dana pribadi, diganti | Jumlah |
|---|---|---|---|---|
| Suami | 5816000 | 80000 | 135000 | 6031000 |
| Istri | 890000 | 40000 | 0 | 930000 |
| Jumlah | 6706000 | 120000 | 135000 | 6961000 |

Suami, dana bersama: 1000000 (tagihan) + 3000000 (ibu) + 1000 (biaya admin) + 75000 (CTH-21) + 200000 (ayah) + 40000 (bensin) + 1500000 (cicilan) = 5816000. Istri, dana bersama: 150000 + 50000 + 100000 (CTH-19, sesudah pengembalian) + 200000 + 75000 + 230000 + 60000 (sayur) + 5000 + 20000 = 890000. Jumlah tabel C = jumlah tabel B. Angka ini mencerminkan siapa yang memegang uang untuk pos apa, bukan siapa yang lebih boros; lihat §10.3.

**Anggaran periode.**

| Kategori | Batas | Terpakai | Terikat | Sisa |
|---|---|---|---|---|
| Belanja bulanan | 2500000 | 540000 | 0 | 1960000 |
| Cicilan | 1500000 | 1500000 | 0 | 0 |
| Tagihan rutin | 1100000 | 1000000 | 0 | 100000 |
| Jajan dan makan di luar | 500000 | 150000 | 0 | 350000 |

**Arus kas.**

| Baris | Jumlah |
|---|---|
| Masuk: pemasukan | 10000000 |
| Masuk: pinjaman diterima | 1000000 |
| Masuk: piutang dibayar | 200000 |
| Masuk: selisih hitung positif | 10000 |
| Jumlah masuk | 11210000 |
| Keluar: belanja dari dana bersama | -5146000 |
| Keluar: cicilan dan utang dibayar (1500000 + 400000) | -1900000 |
| Keluar: penggantian dibayar (60000 + 75000) | -135000 |
| Keluar: pinjaman diberikan | -500000 |
| Keluar: beli aset (2000000 + 400000) | -2400000 |
| Keluar: disisihkan ke dana pribadi | -200000 |
| Keluar: selisih hitung negatif | -85000 |
| Jumlah keluar | -10366000 |
| Perubahan UANG | 844000 |

`UANG` awal 6100000 + 844000 = 6944000 = `UANG` akhir. Belanja dari dana bersama 5146000 = belanja tabel B 5461000 − kontribusi 120000 − talangan 135000 − yang dirinci dari selisih 60000.

**Tabungan, aset, piutang, kewajiban.**

| Posisi | Awal periode | Akhir periode |
|---|---|---|
| Tabungan "dana darurat" (target 10000000) | 5000000 | 5500000 |
| Aset: motor | 12000000 | 11000000 |
| Aset: laptop | 0 | 2000000 |
| Aset: kulkas | 0 | 4000000 |
| Piutang: adik suami | 0 | 200000 |
| Cicilan motor (sisa 5 kali) | 9000000 | 7500000 |
| Cicilan kulkas (sisa 12 kali) | 0 | 3960000 |
| Utang kepada kakak istri | 0 | 600000 |

Aset 11000000 + 2000000 + 4000000 = 17000000. `KB` akhir = 6944000 + 200000 + 17000000 − 11460000 − 600000 − 0 = 12084000.

**Pemeriksaan data penuh** (tidak ditampilkan kepada siapa pun). Kekayaan bersih per pemilik: `shared` 12084000; suami 390000; istri 810000; jumlah 13284000. Menurut INV-08: semua tempat uang (6944000 + 390000 + 810000 = 8144000) + piutang kepada pihak luar 200000 + aset 17000000 − kewajiban kepada pihak luar (11460000 + 600000 = 12060000) = 13284000.

Menguji: INV-05, INV-07, INV-08, INV-15, INV-19.

### 7.11 CTH-38 — Klaim antar-pemilik ditutup saat anggota keluar (berdiri sendiri)

**[Usulan]** (§8.2). Tidak mengubah keadaan berantai dan tidak masuk CTH-37; semua angkanya **[Asumsi]**. Anggota M punya dua klaim terbuka dengan dana bersama saat ia keluar: hak ganti 60000 (rumah tangga berutang kepadanya) dan pinjaman 200000 dari dana bersama (ia berutang kepada rumah tangga).

Kolom: `UANG` = uang yang ada; `PIU-A` = piutang rumah tangga kepada M; `KWJ-G` = hak ganti yang belum dibayar kepada M; `KB = UANG + PIU-A − KWJ-G`; `PRI-M` = uang pribadi M; `KLAIM-M` = hak ganti milik M (piutangnya); `UTG-M` = utang M kepada dana bersama; `KB-M = PRI-M + KLAIM-M − UTG-M`.

| Baris | UANG | PIU-A | KWJ-G | KB | PRI-M | KLAIM-M | UTG-M | KB-M |
|---|---|---|---|---|---|---|---|---|
| Sebelum | 5000000 | 0 | 0 | 5000000 | 300000 | 0 | 0 | 300000 |
| Δ M menalangi 60000, minta diganti (`advance`) | 0 | 0 | 60000 | -60000 | -60000 | 60000 | 0 | 0 |
| Sesudah | 5000000 | 0 | 60000 | 4940000 | 240000 | 60000 | 0 | 300000 |
| Δ dana bersama meminjamkan 200000 kepada M (`loan_given`, antar-pemilik) | -200000 | 200000 | 0 | 0 | 200000 | 0 | 200000 | 0 |
| Sesudah | 4800000 | 200000 | 60000 | 4940000 | 440000 | 60000 | 200000 | 300000 |
| Δ M keluar: `exit_closure` oleh sistem | 0 | -200000 | -60000 | -140000 | 0 | -60000 | -200000 | 140000 |
| Sesudah | 4800000 | 0 | 0 | 4800000 | 440000 | 0 | 0 | 440000 |

Baris entri `exit_closure`:

| Buku | Baris | Nilai |
|---|---|---|
| Rumah tangga | Kewajiban hak ganti kepada M | -60000 |
| Rumah tangga | Sebab `owner_transfer` (hak ganti ditutup) | 60000 |
| Rumah tangga | Piutang kepada M | -200000 |
| Rumah tangga | Sebab `owner_transfer` (pinjaman ditutup) | -200000 |
| Pribadi M | Piutang hak ganti | -60000 |
| Pribadi M | Sebab `owner_transfer` (hak ganti ditutup) | -60000 |
| Pribadi M | Kewajiban kepada dana bersama | -200000 |
| Pribadi M | Sebab `owner_transfer` (pinjaman ditutup) | 200000 |

Keseimbangan buku rumah tangga (piutang − kewajiban = sebab): −200000 − (−60000) = −140000 = 60000 − 200000. Buku pribadi M: −60000 − (−200000) = 140000 = −60000 + 200000. Jumlah `owner_transfer` di semua buku: 60000 − 200000 − 60000 + 200000 = 0. `UANG` dan `PRI-M` tidak berubah: entri ini tidak menyatakan pembayaran apa pun. Kekayaan bersih rumah tangga turun 140000 dan tampil di ringkasan sebagai "klaim ditutup karena anggota keluar": hak ganti 60000, pinjaman −200000. Data penuh sebelum dan sesudah penutupan: 4940000 + 300000 = 5240000 = 4800000 + 440000. Sesudah entri ini buku pribadi M dihapus; yang tersisa di buku rumah tangga adalah `UANG` 4800000 dan `KB` 4800000, tanpa klaim terbuka kepada orang yang sudah pergi. Menguji: INV-02, INV-03, INV-08, INV-25, INV-28.

### 7.12 CTH-39 — Satu periode dengan beberapa penghasilan dan jatah pribadi (berdiri sendiri)

Menurut keputusan dan pernyataan pemilik produk, 11 November 2026 (§3.1, §3.4). Tidak mengubah keadaan berantai dan tidak masuk CTH-37; semua angkanya **[Asumsi]** kecuali jumlah gaji 10000000 dan kiriman ibu 3000000 (brief §1.1). Hari mulai periode 1; periode 1–31 Oktober 2026. Ambang permintaan 500000. Jatah pribadi suami yang disepakati berdua: 1000000 per periode. Semua entri dicatat suami; ia `earner` ketiga penghasilan.

Kolom: `BNK-S` rekening bank suami; `EWL-S` dompet digital suami; `TUN-S` uang tunai bersama di suami; `BNK-I` rekening bank istri; keempatnya berpemilik `shared`. `UANG = BNK-S + EWL-S + TUN-S + BNK-I`. `KB` = kekayaan bersih rumah tangga = `UANG` (tidak ada piutang, aset, atau kewajiban). `PRI-S` = tempat uang pribadi suami (tidak terlihat istri).

| Baris | BNK-S | EWL-S | TUN-S | BNK-I | UANG | KB | PRI-S |
|---|---|---|---|---|---|---|---|
| Sebelum | 500000 | 100000 | 200000 | 300000 | 1100000 | 1100000 | 400000 |
| Δ 1 Okt: gaji termin pertama 6000000, 600000 ke pribadi (`income`) | 5400000 | 0 | 0 | 0 | 5400000 | 5400000 | 600000 |
| Sesudah | 5900000 | 100000 | 200000 | 300000 | 6500000 | 6500000 | 1000000 |
| Δ 3 Okt: kiriman kepada ibu 3000000 (`expense`) | -3000000 | 0 | 0 | 0 | -3000000 | -3000000 | 0 |
| Sesudah | 2900000 | 100000 | 200000 | 300000 | 3500000 | 3500000 | 1000000 |
| Δ 10 Okt: penghasilan lepas 1500000 ke dompet digital (`income`, sumber `business`) | 0 | 1500000 | 0 | 0 | 1500000 | 1500000 | 0 |
| Sesudah | 2900000 | 1600000 | 200000 | 300000 | 5000000 | 5000000 | 1000000 |
| Δ 15 Okt: gaji termin kedua 4000000, 1000000 ke pribadi (`income`) | 3000000 | 0 | 0 | 0 | 3000000 | 3000000 | 1000000 |
| Sesudah | 5900000 | 1600000 | 200000 | 300000 | 8000000 | 8000000 | 2000000 |
| Δ 16 Okt: tarik tunai 2000000, biaya admin 6500 (`transfer`) | -2006500 | 0 | 2000000 | 0 | -6500 | -6500 | 0 |
| Sesudah | 3893500 | 1600000 | 2200000 | 300000 | 7993500 | 7993500 | 2000000 |
| Δ 16 Okt: serah terima 3500000 ke rekening istri (`transfer`) | -3500000 | 0 | 0 | 3500000 | 0 | 0 | 0 |
| Sesudah | 393500 | 1600000 | 2200000 | 3800000 | 7993500 | 7993500 | 2000000 |
| Δ 17 Okt: bantuan kepada adik 500000 dari tunai (`expense`) | 0 | 0 | -500000 | 0 | -500000 | -500000 | 0 |
| Sesudah | 393500 | 1600000 | 1700000 | 3800000 | 7493500 | 7493500 | 2000000 |

Jatah pribadi suami dalam periode ini (§3.4). `Dalam jatah = min(bagian, max(0, jatah − terpakai sebelum))`; `di luar jatah = bagian − dalam jatah`; `terpakai sesudah = terpakai sebelum + bagian`; `terpakai sebelum` baris kedua = `terpakai sesudah` baris pertama.

| Entri | Bagian pribadi | Jatah | Terpakai sebelum | Dalam jatah | Di luar jatah | Terpakai sesudah |
|---|---|---|---|---|---|---|
| Gaji termin pertama | 600000 | 1000000 | 0 | 600000 | 0 | 600000 |
| Gaji termin kedua | 1000000 | 1000000 | 600000 | 400000 | 600000 | 1600000 |

Bagian pertama seluruhnya di dalam jatah, jadi tidak dibandingkan dengan ambang (**[Usulan]**, dokumen 17 D-48). Dari bagian kedua, 400000 menghabiskan jatah dan 600000 berada di luarnya. Menurut pilihan (a) dokumen 17 S2-1 butir (iii) (**[Usulan]**, rancangan sekarang), 600000 itulah yang dibandingkan dengan ambang 500000; menurut pilihan (b), entri itu ditandai karena jatah terlewati, berapa pun bagian di luarnya. Karena 600000 di atas 500000, kedua pilihan memberi hasil yang sama pada contoh ini (§3.4). Apakah dan bagaimana entri itu ditandai adalah aturan dokumen 06 §6.2 "Pemakaian di atas ambang yang dicatat tanpa permintaan". Jatah tidak mengubah satu angka pun di tabel pertama: tanpa jatah, ketujuh baris `Δ` sama persis.

Keseimbangan gaji termin kedua, buku rumah tangga: 3000000 = 4000000 − 1000000; buku pribadi suami: 1000000 = 1000000; jumlah `owner_transfer`: −1000000 + 1000000 = 0. Tarik tunai: `BNK-S` berkurang 2000000 + 6500 = 2006500, `TUN-S` bertambah 2000000, dan hanya biaya admin 6500 yang menjadi pengeluaran. Serah terima dicatat suami ke tempat uang yang dipegang istri, jadi baris +3500000 pada `BNK-I` bertanda belum dikonfirmasi sampai istri mengonfirmasi (§3.3); saldonya sudah berubah. Bantuan kepada adik 500000 tidak di atas ambang; kiriman kepada ibu 3000000 di atas ambang dan, seperti CTH-06, dijadwalkan sebagai tagihan berulang (§3.15); dengan andaian yang sama dengan CTH-06 (**[Asumsi]**: tercakup kesepakatan rutin 3000000 per kali), ia tidak ditandai. Tanpa kesepakatan itu ia membawa tanda netral; angkanya tetap.

Ringkasan periode: pemasukan 6000000 + 1500000 + 4000000 = 11500000 (gaji 10000000, hasil usaha 1500000); disisihkan ke dana pribadi suami 600000 + 1000000 = 1600000; belanja 3000000 + 6500 + 500000 = 3506500; perubahan `KB` = 11500000 − 1600000 − 3506500 = 6393500, dan 1100000 + 6393500 = 7493500. Arus kas: pemasukan yang masuk tempat uang dana bersama 5400000 + 1500000 + 3000000 = 9900000; belanja 3506500; perubahan `UANG` = 9900000 − 3506500 = 6393500. Custody akhir: suami memegang 393500 + 1600000 + 1700000 = 3693500; istri memegang 3800000, di antaranya 3500000 belum dikonfirmasi; 3693500 + 3800000 = 7493500. `PRI-S`: 400000 + 1600000 = 2000000. Aturannya simetris: penghasilan istri dicatat istri dengan bentuk yang sama dan dihitung terhadap jatahnya sendiri.

Menguji: INV-02, INV-03, INV-06, INV-19, INV-20, INV-24, INV-27, INV-30.

---

## 8. Riwayat yang tidak dapat diubah lawan hak penghapusan

**Status: [Usulan] yang memerlukan ADR (`architect`, ADR-F) dan tinjauan hukum (L-04). Bukan nasihat hukum.**

### 8.1 Ketegangannya

- Brief §7 dan `ENGINEERING-STANDARDS.md` §5: riwayat keuangan hanya ditambah.
- UU PDP Pasal 43–44 mewajibkan penghapusan dan pemusnahan atas permintaan subjek data; pengecualiannya (Pasal 50) hanya untuk kepentingan negara, dan tidak ada pengecualian untuk keutuhan catatan yang dimiliki bersama orang lain (**[Terverifikasi]** primer, dokumen 18 §2.1, diakses 11 November 2026). Salinan teks PP 33/2026 mengukur pemusnahan dari hasilnya, "tidak lagi dapat digunakan untuk mengidentifikasi Subjek Data Pribadi" (Pasal 84 ayat (2)), menyebut salinan cadangan (Penjelasan Pasal 82), dan tidak menyebut pseudonimisasi sebagai cara memenuhi kewajiban hapus (dokumen 18 §2.1).
- Apakah pseudonimisasi identitas pada entri memenuhi kewajiban itu: **[Perlu validasi]** hukum (L-04).

### 8.2 Aturan yang diusulkan

Entri dipisah menjadi dua bagian sejak awal (§2.5, §2.6):

| Bagian | Isi | Sifat |
|---|---|---|
| **Inti moneter** | `entry` dan `line`: jenis, tanggal, jumlah, posisi, sebab, kategori, tautan; pelaku dan pihak sebagai **pengenal internal**, bukan nama | Hanya ditambah; tidak pernah dihapus |
| **Rincian** | `entry_detail`: catatan teks, struk, nama pihak lawan dalam teks bebas | Dapat dihapus tanpa menyentuh inti moneter |
| **Pemetaan identitas** | Tabel yang mengaitkan pengenal internal anggota dan pihak luar dengan orangnya (nama, akun) | Dapat diputus: setelah diputus, pengenal itu tidak lagi menunjuk siapa pun |

Atas permintaan penghapusan dari seorang anggota, atau ketika ia keluar dari rumah tangga dan memintanya:

1. **Ekspor ditawarkan lebih dulu** (data miliknya sendiri).
2. **Klaim antar-pemilik yang masih terbuka ditutup oleh sistem, lalu buku pribadinya dihapus seluruhnya.** Buku pribadi bukan catatan bersama; tidak ada orang lain yang bergantung pada keutuhannya. Penghapusan **tidak bersyarat** pada penggantian, pelunasan, atau pelepasan oleh siapa pun (temuan `security-privacy` SP-12): keluar dengan aman tidak boleh menunggu tindakan pasangan, dan anggota yang tersisa tidak boleh tertinggal dengan klaim yang hanya dapat ditutup orang yang sudah pergi. Saat anggota keluar, sistem menulis satu entri `exit_closure` (di bawah) yang menolkan setiap klaim antara buku pribadinya dan buku lain, sehingga buku rumah tangga dan buku pasangan tetap seimbang (INV-08). Kedua pihak tetap boleh menyelesaikan klaim itu lebih dulu dengan entri biasa bila mau; itu pilihan, bukan syarat. Sisi klaim yang berada di buku rumah tangga atau buku pasangan tetap ada sebagai riwayat bersaldo nol, dengan pihak lawan dipseudonimkan. Penghapusan baris buku pribadi ini **satu-satunya** penghapusan baris yang diusulkan, dan ADR harus menyatakannya sebagai pengecualian eksplisit terhadap INV-09. Entri dua buku (kontribusi, talangan, penyisihan) kehilangan baris sisi pribadinya; baris sisi rumah tangganya tetap dan tetap seimbang per pemilik (INV-02), sedangkan INV-03 dan INV-08 sejak itu hanya diperiksa atas pemilik yang masih ada.
3. **Di buku rumah tangga, inti moneter tetap.** Pelaku, pembayar, penerima manfaat, `earner`, dan pemegang yang menunjuk anggota itu tetap berupa pengenal internal; pemetaan identitasnya diputus, sehingga layar menampilkan "mantan anggota".
4. **Rincian yang ia tulis dihapus**: catatan teks dan struk pada entri yang ia catat, kecuali pada entri dana bersama yang rinciannya juga data anggota yang tersisa. Untuk kasus itu usulannya: struk dan catatan tetap, nama orang yang pergi di dalam teks bebas tidak dapat dijamin terhapus, dan hal ini dinyatakan terbuka sebagai batas (lihat 8.3).
5. **Tempat uang dana bersama yang ia pegang** tidak dihapus dan tidak dinolkan otomatis: uangnya nyata. Anggota yang tersisa mencatat apa yang sebenarnya terjadi (serah terima, atau pembagian sebagai `owner_transfer`), lalu mengarsipkannya. Pembagian harta ketika rumah tangga berpisah adalah keputusan keluarga atau hukum, bukan aturan aplikasi.
6. **Cadangan**: penghapusan rincian dan pemutusan pemetaan harus menjangkau salinan cadangan dalam jangka yang ditetapkan dokumen 10.
7. **Pihak luar** (ibu, adik, ayah) hanya tersimpan sebagai nama panggilan dan hubungan (Q-28); penghapusannya adalah pemutusan pemetaan yang sama.
8. **Entri dua buku lama.** Sesudah anggota itu keluar, tidak ada lagi yang dapat membalik entri dua buku yang pernah ia catat (§3.12). Anggota yang tersisa tetap dapat menyesuaikan sisi rumah tangganya. Bagaimana objek yang satu-satunya pelakunya sudah pergi ditampilkan dan dikeluarkan dari daftar terbuka adalah aturan dokumen 06 (§8.2 "Keberatan", §7.2 "Aturan umum").

**Entri `exit_closure` [Usulan]; sisi hukumnya [Perlu validasi] (L-04).**

| Hal | Aturan |
|---|---|
| Jenis dan pelaku | `type = exit_closure`, `actor = system`. Bukan tindakan anggota mana pun: tidak memerlukan persetujuan, konfirmasi, atau pernyataan dari yang keluar maupun yang tersisa, dan tidak dapat dicatat lewat layar |
| Kapan | Satu entri per anggota yang keluar, ditulis dalam transaksi basis data yang sama dengan penutupan buku pribadinya; tanggal kejadian = tanggal keluar; idempotency key diturunkan dari peristiwa keluarnya, sehingga pengulangan tidak menulis apa pun (INV-10) |
| Baris | Untuk setiap klaim antar-pemilik yang menyangkut buku pribadi anggota itu dan bersaldo X > 0: piutang −X dan `owner_transfer` −X di buku pemilik piutang; kewajiban −X dan `owner_transfer` +X di buku yang berkewajiban. Tidak ada baris tempat uang, aset, atau alokasi |
| Yang dinyatakannya | Hanya bahwa aplikasi berhenti melacak klaim itu karena salah satu pihaknya keluar. Ia **tidak** menyatakan klaim itu dibayar, dan **tidak** menyatakan siapa pun mengikhlaskannya. Layar: "ditutup karena anggota keluar" |
| Akibat pada uang | Tidak ada: `U(o)` semua pemilik tetap; *uang yang ada* tetap |
| Akibat pada kekayaan bersih rumah tangga | Naik sebesar kewajiban rumah tangga kepada anggota itu yang ditutup (hak ganti, utang dana bersama kepadanya); turun sebesar piutang rumah tangga kepadanya yang ditutup. Klaim antara dua buku pribadi tidak menyentuh angka rumah tangga |
| Laporan | Baris tersendiri di tabel A, "klaim ditutup karena anggota keluar", per klaim (§6.3). Bukan kontribusi, bukan penyisihan, bukan "utang dihapuskan"; tidak terhitung terhadap anggaran; tidak muncul di arus kas; *jumlah pemakaian* nol, jadi tidak dibandingkan dengan ambang dan tidak ditandai (§3.8) |
| Invarian yang berlaku | INV-01, INV-02 (seimbang di tiap buku), INV-03 (jumlah `owner_transfer` nol), INV-05 dan INV-06 dalam arti `ΔU = 0`, INV-07, INV-08 (kedua sisi tiap klaim nol), INV-09 (entri hanya ditambah), INV-10, INV-13 (X tepat sebesar sisa), INV-19, INV-20, INV-28. INV-25 mendapat pengecualian kedua |
| Tidak dapat dibalik | Sesudah buku pribadinya dihapus tidak ada sisi lawan untuk dipulihkan. Bila orang itu bergabung lagi, ia anggota baru dengan buku baru; klaim lama, bila masih diakui keduanya, dicatat ulang sebagai saldo awal |
| Sisi hukum **[Perlu validasi]** | Apakah penutupan catatan di aplikasi dapat dibaca sebagai pelepasan hak atas utang yang sebenarnya, dan bagaimana kedudukannya terhadap harta bersama, belum ditinjau advokat (L-04). Maksud model: tidak; utang-piutang yang sebenarnya antara kedua orang itu tidak diputuskan aplikasi. Ekspor yang ditawarkan lebih dulu (butir 1) memuat daftar klaim terbuka beserta jumlahnya bagi yang keluar, dan sisi rumah tangganya tetap terbaca oleh yang tersisa |

Contoh terhitung: CTH-38.

### 8.3 Batas yang dinyatakan

| Batas | Akibat |
|---|---|
| Di rumah tangga dua orang, "mantan anggota" mudah dikenali oleh anggota yang tersisa | Pseudonimisasi tidak menyembunyikan apa pun dari pasangan; ia hanya memutus kaitan bagi operator dan pihak lain. Apakah itu memenuhi Pasal 84 ayat (2): L-04 |
| Jumlah dan tanggal belanja dana bersama tetap ada | Bila advokat menyatakan inti moneter pun harus dimusnahkan, satu-satunya jalan yang menjaga angka adalah **entri penutup**: buku rumah tangga dirangkum menjadi saldo awal baru per posisi, lalu entri lama dihapus. ADR harus menyebut jalan ini sebagai rencana cadangan dan biayanya: jejak audit sebelum tanggal rangkuman hilang bagi anggota yang tersisa |
| Teks bebas dan foto struk dapat memuat data orang lain | Tidak dapat dipilah otomatis. Usulan: teks bebas pada entri dana bersama diperlakukan sebagai data bersama |
| Hak anggota yang tersisa atas catatan bersama tidak diatur UU PDP sebagai pengecualian | Inilah inti L-04 |

---

## 9. Keputusan atas pertanyaan terbuka

Keputusan dokumen ini atas butir dokumen 17 dan butir glosarium yang semula bertanda *(diputuskan di S2)*. Semua **[Usulan]** sampai ditinjau pemilik produk di checkpoint S2, kecuali baris bertanda "keputusan pemilik" (diputuskannya sendiri) dan baris bertanda "default yang diterima" (default yang ia terima lewat "lanjut" di checkpoint S1; keduanya dibedakan sesuai K-24).

| ID | Keputusan | Bagian |
|---|---|---|
| Q-01 | Keputusan pemilik: gaji dana bersama secara bawaan, dapat diubah saat mencatat. *Pinjam kembali* adalah serah terima: default yang diterima pemilik produk lewat "lanjut" di checkpoint S1, bukan keputusannya sendiri | §3.1, §3.3, CTH-27 |
| Q-56, K-16 | Keputusan pemilik produk, 11 November 2026: pencatat uang masuk memilih arahnya, termasuk gaji ke dana pribadi; pasangan melihat jumlah gaji dan bagian pribadinya; saldo dan pemakaian dana pribadi tetap tertutup. Bentuknya: satu entri `income` dengan baris di dua buku; penghasilan (`salary`, `business`) tidak dapat dicatat hanya di buku pribadi (INV-24). Di luar keputusan itu dan masih **[Usulan]**, menunggu pemilik produk di checkpoint S2: hadiah untuk satu orang tidak terlihat (Q-77), dan bagian pribadi di dalam entri `income` ikut dibandingkan dengan ambang (T-13, Q-78) | §3.1, §3.4, §5.3, CTH-03b |
| Q-55 (sisi model) | "Punya siapa" pada layar adalah `earner`, terpisah dari pemilik uang; pinjaman yang diterima adalah `loan_received`, bukan pemasukan, dan bawaannya utang rumah tangga | §2.4, §3.1, CTH-28 |
| Q-02, Q-23, K-13 | Dana pribadi **di luar** garis kendali rumah tangga. *Uang yang ada* = dana bersama saja = angka utama, identik bagi keduanya; tidak ada nama kedua. Baris kedua "uang pribadi saya" hanya milik pelihat | §2.2, §5 |
| Q-05 | Keputusan pemilik untuk asasnya. Rinciannya: selisih negatif terhitung sebagai uang keluar yang belum dirinci; selisih positif baris tersendiri di sisi masuk, bukan pemasukan; terhitung terhadap anggaran hanya setelah dirinci; entri terlambat dapat menutup selisih tanpa menyentuh saldo; tanpa penutupan periode | §3.10, §3.11, CTH-29 |
| Q-06, Q-29 | Keputusan pemilik: anggaran adalah batas per kategori, bulanan, tanpa amplop. Hari mulai periode ditentukan rumah tangga (keputusan pemilik produk 15, 11 November 2026, menyempurnakan "mulai tanggal gajian" dari checkpoint S1); "bawaan 1, berlaku mulai periode berikutnya" adalah bacaan lead yang dikukuhkan kepadanya, bukan kata-katanya. **[Usulan]** dokumen ini: rentang 1 sampai 28 dan satu periode peralihan pendek dengan batas penuh. Periode disimpan sebagai rentang tanggal; tanpa rollover | §6.1 |
| Q-07 | Default yang diterima pemilik produk lewat "lanjut" di checkpoint S1: kiriman ibu adalah pengeluaran saat diserahkan; saudara dan ayah dipilih per entri | §3.2, CTH-06, CTH-25, CTH-26 |
| Q-11 | Opsi B: ledger sederhana, entri bertipe dengan baris seimbang | §1 |
| Q-12 | Persetujuan tidak membuat reservasi, kecuali permintaan menyebut tempat uang sumber | §3.8, CTH-22 |
| Q-13 | Sisa: tetap di tempatnya (A, D); dikembalikan atau tetap dipegang pemohon sebagai dana bersama (B); hak ganti sebesar realisasi (C). Kelebihan: tercatat penuh; dari dana bersama tanpa klaim; dari dana pribadi kontribusi kecuali "minta diganti" | §3.8, CTH-21 sampai CTH-24 |
| Q-14 | Ya: terikat mengurangi sisa anggaran kategori otorisasinya dan tampil terpisah | §6.1, CTH-21 |
| Q-15, K-02 | Pengembalian dana mengurangi pengeluaran kategori asal pada periode tanggal pengembalian; pemasukan hanya bila aslinya tidak tercatat | §3.13, CTH-19 |
| Q-16, K-11 | Kewajiban cicilan = jumlah pembayaran tersisa; pembayaran adalah pelunasan yang terhitung terhadap anggaran; bunga diakui sebagai `credit_cost` saat kredit dibuka. *Realisasi* didefinisikan ulang sebagai jumlah baris tertaut, apa pun jenisnya | §3.14, §3.8, CTH-04, CTH-34 |
| Q-17 | Bukan kewajiban; pengingat sampai dibayar | §3.15 |
| Q-18 | Pembelian aset bukan pengeluaran; tampil di arus kas; nilai diperbarui manual dengan entri bertanggal | §3.16, CTH-33 |
| Q-19, K-20 (sisi uang) | Keputusan pemilik produk, 11 November 2026: tercatat seketika, bertanda belum dikonfirmasi sampai penerima menerimanya, keberatan kapan saja, sama ke dua arah. Rinciannya: saldo berubah seketika; penanda berlaku atas setiap baris yang menambah tempat uang berpemegang anggota dewasa selain pencatat; keberatan tidak mengubah angka sampai ada entri pembalik | §3.3, §5 |
| Q-20 | Pemegang adalah atribut tempat uang; tempat uang bersama dilaporkan sebagai baris "dipegang bersama" | §2.3, CTH-02 |
| Q-21 | Hak ganti tercatat sejak talangan sebagai kewajiban rumah tangga kepada anggota, berpasangan dengan piutang di buku pribadinya | §3.5, CTH-14 |
| N-14 | Pelepasan hak ganti adalah `owner_transfer`: kenaikan kekayaan bersih rumah tangga tampil sebagai kontribusi dari dana pribadi | §3.6, CTH-16 |
| Q-22, K-10 | Usulan §8: inti moneter buku rumah tangga tetap dan dipseudonimkan; rincian dan buku pribadi dapat dihapus; menunggu ADR-F dan L-04 | §8 |
| Q-24 | Diterima sebagai sifat klaim dua sisi; daftar kebocoran B1–B5 | §5.3 |
| Q-26 (sisi uang) | Default yang diterima pemilik produk: keduanya boleh mengoreksi entri dana bersama. Bentuknya entri pembalik atau penyesuaian; entri buku pribadi hanya oleh pemiliknya; entri dua buku hanya dapat dibalik pemilik buku pribadinya, sedangkan pasangannya dapat menyesuaikan sisi rumah tangga dan berkeberatan | §3.12, §2.2, CTH-30 |
| Dokumen 06 §15.4 butir 1 | Kait `authorization.amended`: tanpa entri; batas, tanggal, masa berlaku, *terikat*, dan reservasi mengikuti versi syarat baru | §3.8, varian CTH-22 |
| Dokumen 06 §15.4 butir 2 | Entri terlambat boleh ditautkan ke otorisasi `settled` atau `expired` bila tanggal kejadiannya di dalam rentang hidup otorisasi itu; tidak ke `cancelled`. Tautan boleh ditambahkan sesudah entri tersimpan sebagai catatan `entry_link` yang hanya ditambah | §3.8, §2.5, INV-26 |
| Dokumen 06 §15.4 butir 3 | Penanda `unconfirmed` berlaku atas setiap baris yang menambah tempat uang berpemegang anggota dewasa selain pencatat | §3.3, §2.5 |
| Dokumen 06 §15.4 butir 4 | Dikukuhkan dan dipertegas: setiap baris buku pribadi M hanya ditulis M, dengan satu pengecualian bagi anggota lain (pelepasan klaim dana bersama atas M) dan satu bagi sistem (`exit_closure` saat M keluar, **[Usulan]**). Entri antara dua buku pribadi ditunda dari MVP | §2.2, INV-25 |
| Dokumen 06 §15.4 butir 5 | `set_aside_schedule`, dan `routine_agreement_ref` pada tagihan berulang dan jadwal penyisihan (tidak pada cicilan, lihat butir (a)) | §2.6, §3.4 |
| Dokumen 06 §15.4 butir 6 | `links.activity_id` diganti `links.subject_ref` (rujukan induk) | §2.5 |
| Dokumen 06 §15.4 butir 7 | Pembelian aset secara cicilan: uang muka ditambah jumlah seluruh cicilan | §3.8, §3.16, CTH-34 |
| Dokumen 06 §15.4 butir (a) | Diterapkan: `routine_agreement_ref` dihapus dari `installment`; pembayaran cicilan dan pelunasan utang yang sudah tercatat (`repayment_out`) tidak dibandingkan dengan ambang dan tidak memerlukan kesepakatan rutin. Tidak ada keberatan dari sisi uang; batas yang semula dinyatakan (kewajiban yang didaftarkan belakangan sebagai saldo awal) kini dijawab butir (m) | §2.6, §3.8, §3.14 |
| Dokumen 06 §15.4 butir (b) | Diterapkan: tabel *jumlah pemakaian* memuat `claim_waiver` untuk kedua arahnya | §3.8, §3.6 |
| Dokumen 06 §15.4 butir (c) | Dikukuhkan: `member(M)` adalah anggota dalam satu rumah tangga; buku pribadi per anggota, bukan per akun; tanpa entri lintas rumah tangga | §2.2 |
| Dokumen 06 §15.4 butir (d) | Tidak bertentangan: jumlah yang disepakati (per kali untuk tagihan berulang, per periode anggaran untuk jatah pribadi) dibaca dari kesepakatan; `estimated_amount` dan `monthly_amount` pada jadwal adalah metadata perkiraan dan rencana | §2.6 |
| Dokumen 06 §15.4 butir (e), (f), (i), (j), (k) | Sudah diterapkan pada revisi sebelumnya; lihat baris SP-01, SP-05, SP-12, SP-13 di bawah. Judul INV-24 sudah diganti (SP-22) | INV-27, §2.5, §3.3, §3.12, §8.2, §6.3, INV-24 |
| Dokumen 06 §15.4 butir (g) | Dikukuhkan: penutupan keberatan (C05) tidak memicu kait uang apa pun dan bukan konfirmasi; `handover.confirmed` hanya dari C02 dan C03 | §3.3, §10.2 |
| Dokumen 06 §15.4 butir (h) | Diterapkan: tautan susulan ke otorisasi hanya oleh pencatat entri atau pemohon otorisasi; pencabutan hanya oleh pembuat tautan atau pencatat entri; aturannya milik dokumen 06 (INV-A34) dan dikutip | §3.8 |
| Dokumen 06 §15.4 butir (l) | Diterapkan: kalimat kedua INV-24 dibatasi pada entri `income` yang memuat baris bersumber penghasilan; hadiah pribadi di buku pribadi saja tetap sah (CTH-36) | INV-24 |
| Dokumen 06 §15.4 butir (m) | Diterapkan dengan dua penegasan sisi uang: (1) baris `expense` berpemilik `shared` dihitung di entri jenis apa pun, **kecuali** baris `expense` pada `contribution_expense` (bukan dana bersama) dan pada `advance` (sudah dihitung sebagai hak ganti); (2) saldo awal kewajiban dana bersama kepada pihak luar punya jumlah pemakaian sebesar saldo awalnya, dibandingkan satu kali terhadap ambang yang berlaku saat disimpan; pembayarannya tidak ditautkan lagi bila `opening`-nya tertaut ke otorisasi | §3.8, §3.14, §3.17 |
| Pernyataan pemilik produk, 11 November 2026: kedua pasangan berpenghasilan | Tidak ada aturan baru. Aturan sudah per penerima penghasilan (`earner`) dan simetris; kalimat yang menyiratkan satu pencari nafkah diperbaiki. Bagian pribadi yang rutin dari penghasilan memakai `set_aside_schedule` per penerima dan dapat dicakup *kesepakatan rutin*; tidak ada kolom yang kurang. Bentuk jatah pribadi kini diputuskan pemilik produk (dua baris berikut) | §3.1, §3.4, §7.1, CTH-03b, §10.3 |
| Keputusan pemilik produk, 11 November 2026: jatah pribadi (Q-98) | Keputusan pemilik produk (14, 16, 17): jatah pribadi adalah satu jumlah per orang per periode anggaran, dijumlahkan sepanjang periode; besarnya dapat diatur dan **disepakati berdua**: satu mengusulkan, yang lain mengonfirmasi; kenaikan diajukan dengan alasan. Selebihnya rancangan dokumen ini dan dokumen 06. Bentuknya: `set_aside_schedule` per penerima yang dirujuk satu kesepakatan rutin. **[Usulan]**, bukan keputusan pemilik produk: penyisihan kepada M (entri `owner_transfer` dan bagian pribadi entri `income`) di dalam jatah tidak dibandingkan dengan ambang (dokumen 17 D-48), dan bagian di luarnya dibandingkan seperti penyisihan lain (terbuka: S2-1 butir (iii); alternatifnya, yang melampaui jatah ditandai berapa pun jumlahnya); tanpa akibat pada saldo, anggaran, atau kekayaan bersih; tidak berpindah periode; tanpa angka lintas periode. Setiap penyisihan kepada M terhitung dengan sendirinya, tanpa tautan; periodenya menurut tanggal kejadian; entri pembalik melonggarkan jatah ke depan saja (sama dengan dokumen 06 §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi"). Kenaikan jatah diajukan dengan alasan dan dikonfirmasi pasangan (pernyataan pemilik produk); alurnya milik dokumen 06, tanpa akibat uang | §2.6, §3.4, §3.8, INV-30, CTH-39 |
| Keputusan pemilik produk, 11 November 2026: hari mulai periode anggaran | Keputusan pemilik produk (15): hari mulai ditentukan rumah tangga. Bacaan lead yang dikukuhkan kepadanya, bukan kata-katanya: satu hari mulai, bawaan 1, berlaku mulai periode berikutnya; periode lampau tidak berubah. **[Usulan]**: 1 sampai 28; satu periode peralihan pendek (1 sampai 30 hari) dengan batas anggaran dan jatah penuh. Putusan lead, **[Usulan]** menunggu pemilik produk (Q-102; dokumen 17 D-49): penggantian hari mulai adalah kesepakatan rumah tangga (kedua anggota dewasa, seperti ambang), dan terlihat oleh keduanya | §2.6, §6.1, INV-29 |
| Pernyataan pemilik produk, 11 November 2026: penghasilan datang beberapa kali ("sebulan" adalah bacaan lead), ke rekening bank atau dompet digital; diteruskan ke istri atau ditarik tunai lalu dibagikan | Tidak ada aturan baru: tiap uang masuk satu entri; serah terima dari rekening (§3.3; CTH-08, CTH-18), tarik tunai (CTH-07) dan biaya admin (CTH-18), kiriman keluarga (CTH-06, CTH-26) dan pinjaman (CTH-25) sudah tercakup; satu contoh menggabungkannya | §3.1, §3.2, §3.3, CTH-39 |
| Q-27 (masukan) | Laporan per orang hanya dari buku rumah tangga dan sama bagi keduanya | §6.3 |
| SP-01 (pemeriksaan awal `security-privacy`) | **[Usulan], menunggu pemilik produk di checkpoint S2**: bagian pribadi di dalam entri `income` bersumber penghasilan hanya bila `actor = earner`; `earner` adalah isian "punya siapa", wajib untuk penghasilan. Alternatif: tanpa pagar, pencatat mana pun boleh mengarahkan ke dana pribadinya sendiri | §3.1, INV-27 |
| SP-18 | Daftar sumber penghasilan adalah parameter preset aturan; preset pilot `salary`, `business` | §2.4, §3.1, INV-24 |
| SP-05 (sisi uang) | **[Usulan]**: penanda `unconfirmed` juga atas baris entri dua buku yang mengurangi tempat uang dana bersama berpemegang anggota dewasa selain pencatat; klaim internal yang dikeberatani tampil "dikeberatani" di daftar utang dan piutang, tanpa perubahan angka | §2.5, §3.3, §3.12, §5 |
| SP-12 (sisi uang) | **[Usulan]**; sisi hukum **[Perlu validasi]**: penghapusan buku pribadi tidak bersyarat; klaim antar-pemilik yang terbuka ditutup entri sistem `exit_closure` yang tidak menyatakan pembayaran dan tidak memerlukan tindakan siapa pun | §8.2, INV-28, CTH-38 |
| SP-13 | **[Usulan]**: tidak ada angka kumulatif, lintas periode, atau bersih per anggota atas penyisihan dan kontribusi; hanya baris per periode. Pelepasan klaim dana bersama atas M tampil sebagai "utang dihapuskan", bukan penyisihan kepada M | §6.3, §3.6, §5.3 |
| Q-86 | **[Usulan]**: uang antara dua rumah tangga yang sama-sama memakai aplikasi adalah dua catatan yang berdiri sendiri, tanpa tautan; bentuk yang lebih kaya ditunda ke tahap publik | §2.2 |
| Q-43 | Rupiah utuh; sen bank dibulatkan ke bawah | §2.1 |
| Q-44 | Tidak: satu tempat uang, satu pemilik | §2.3 |
| Q-45 | Jumlah alokasi dibatasi saldo tercatat saat perintah alokasi; belanja tidak pernah ditolak dan tidak mengurangi alokasi lain otomatis; kekurangannya tampil sebagai peringatan | §3.9 |
| Q-46 | Ya, tanggal akhir opsional | §3.15 |
| Q-47 | Baris tersendiri "piutang diikhlaskan", pada bulan penghapusan, di luar anggaran dan arus kas | §3.7, CTH-25 |
| Q-48 | Di entri yang sama, sebagai baris pengeluaran terpisah | §3.3, CTH-18 |
| Q-49 (sisi uang) | Penyisihan adalah `owner_transfer`, bukan pemindahan dan bukan pengeluaran; selalu terlihat keduanya; yang rutin punya jadwal penyisihan | §3.4, CTH-35 |
| Glosarium: saldo awal, entri terlambat, arsip | Lihat bagiannya | §3.17, §3.11, §2.6 |

---

## 10. Butir yang masih terbuka dan permintaan ke pihak lain

### 10.1 Butir terbuka untuk dokumen 17 (`product-analyst`)

Asumsi dan bawaan yang dipakai dokumen ini tanpa pernyataan keluarga. Nomor `Q` dan `A` diberikan pemilik dokumen 17.

| # | Butir | Bawaan yang diusulkan | Akibat bila lain |
|---|---|---|---|
| T-01 | Tanda "minta diganti" pada belanja dana pribadi yang tertaut ke otorisasi dana bersama | Hidup secara bawaan (§3.5) | Bila mati, situasi C menjadi kontribusi dan CTH-23 berubah |
| T-02 | Sisa pada situasi B: dikembalikan atau tetap dipegang pemohon | Pemohon memilih; tanpa pilihan, tetap dipegang (tanpa entri) | Bila wajib dikembalikan, penyelesaian menunggu serah terima |
| T-03 | Dompet fisik yang memuat dana bersama dan dana pribadi sekaligus | Dua tempat uang; selisih rekonsiliasi dibebankan ke yang `shared` kecuali pemegang memilih lain | Bila rumah tangga pilot tidak memisahkan uang pribadi secara fisik, usaha ini terasa setiap kali menghitung |
| T-04 | Entri terlambat ditawarkan sebagai "bagian dari selisih hitung" | Bawaan ya | Bila tidak, saldo turun dua kali dan rekonsiliasi berikutnya menghasilkan selisih positif |
| T-05 | Barang mana yang dicatat sebagai aset | Pilihan pencatat; tanpa batas nilai | Batas nilai memerlukan satu pengaturan lagi |
| T-06 | Saldo tercatat negatif | Diperbolehkan, memicu ajakan rekonsiliasi | Bila ditolak, pencatat pelupa terhalang mencatat |
| T-07 | Sen pada saldo bank | Dibulatkan ke bawah | Bila dicatat persis, satuan IDR menjadi sen dan semua layar menampilkan dua desimal |
| T-08 | Arah uang masuk dipilih pencatat, termasuk gaji ke dana pribadi | Bukan lagi asumsi: keputusan pemilik produk, 11 November 2026. Modelnya di §3.1; risiko yang tersisa di §10.3 butir 1 | — |
| T-09 | Angka rumah tangga contoh | Semua **[Asumsi]** kecuali 10000000 dan 3000000 | Data pilot yang sebenarnya dimasukkan saat onboarding, bukan ke dokumen ini |
| T-10 | Kategori bawaan saat onboarding | Daftar contoh yang semuanya dapat dihapus | Diputuskan dokumen 05 dan 08 |
| T-11 | Titipan pada pihak luar tanpa login | Tersedia di model; tidak dipakai pilot | Bila dipakai, seseorang harus mencatat pemakaiannya |
| T-12 | Entri antara dua dana pribadi (talangan, pinjaman, penggantian antara uang pribadi suami dan uang pribadi istri) | Alur pencatatannya ditunda dari MVP; model dan contohnya (CTH-17b) tetap | Bila pilot membutuhkannya: satu tempat uang pribadi bawaan per anggota sebagai tujuan baris masuk yang dicatat orang lain, bertanda belum dikonfirmasi; INV-25 mendapat pengecualian kedua |
| T-13 | Bagian pemasukan yang diarahkan langsung ke dana pribadi, terhadap ambang permintaan | **[Usulan], belum diputuskan; menunggu pemilik produk di checkpoint S2** (dokumen 17 D-18 dan D-48 untuk perlakuan yang sama dan bagian di dalam jatah; S2-1 butir (iii) untuk bagian di luar jatah; Q-78). Usulannya: diperlakukan sama dengan penyisihan lain, yaitu terhitung terhadap jatah pribadi penerimanya dan, untuk bagian di luar jatah, ditandai menurut aturan penyisihan (§3.4); tidak pernah dihalangi | Bila dikecualikan, penyisihan besar dapat lolos dari tanda dengan dicatat pada hari gajian. Aturan tandanya milik dokumen 06; bentuk entri dan angka tidak berubah pada jawaban mana pun |
| T-14 | Sumber pemasukan yang wajib meninggalkan catatan di buku rumah tangga | `salary` dan `business`; `gift` untuk satu orang dan `other` boleh dicatat di buku pribadi saja | Bila `other` ikut wajib, pemasukan dana pribadi sendiri (misalnya menjual barang pribadi) ikut terlihat, dan yang tertutup tinggal hadiah |
| T-15 | "Langsung dipakai untuk sesuatu" pada layar uang masuk | Dua entri (pemasukan, lalu pemakaiannya) dalam satu alur layar | Bila harus satu entri, jenis entri `income` perlu memuat baris jenis lain, dan daftar jenis entri tidak lagi tertutup |
| T-16 | Arah "ke dana pribadi" bila yang mencatat gaji bukan penerimanya | Hanya ke dana pribadi pencatat sendiri (INV-25); dan, **[Usulan] menunggu pemilik produk di checkpoint S2**, tidak tersedia sama sekali bagi pencatat yang bukan `earner` (INV-27, T-17): ia mencatat ke dana bersama, dan bagian pribadi siapa pun diambil lewat penyisihan | Bila pasangan boleh mengarahkan ke dana pribadi orang lain, perlu pengecualian T-12 |
| T-17 | Pencatat yang bukan penerima penghasilan mengarahkan sebagian penghasilan itu ke dana pribadinya sendiri (SP-01) | **[Usulan]**: tidak boleh; hanya `actor = earner` (INV-27; dokumen 17 S2-1 butir (i)). Adil hanya bersama jawaban "ya" atas T-13 | Bila boleh: A dapat mengarahkan gaji B ke dana pribadi A dalam entri yang tidak dapat dibatalkan B; B hanya dapat berkeberatan. Angka contoh tidak berubah pada jawaban mana pun |
| T-18 | Daftar sumber penghasilan sebagai parameter preset (SP-18) | Preset pilot: `salary`, `business`; perubahan berlaku ke depan | Siapa yang boleh mengubahnya dan lewat alur apa: dokumen 06 |
| T-19 | Klaim antar-pemilik yang terbuka saat anggota keluar (SP-12) | **[Usulan]**: ditutup entri sistem `exit_closure`, tanpa syarat dan tanpa tindakan siapa pun; tidak menyatakan pembayaran. Sisi hukum **[Perlu validasi]** (L-04) | Bila advokat menilai penutupan itu dapat dibaca sebagai pelepasan hak: klaim dibiarkan terbuka sebagai klaim kepada "mantan anggota" (pihak luar), dengan biaya kekayaan bersih rumah tangga memuat klaim yang tidak dapat lagi ditutup dari dalam aplikasi kecuali dengan penghapusan piutang |
| T-20 | Rincian per anggota atas penyisihan dan kontribusi (SP-13) | Hanya di dalam satu periode laporan; tidak ada angka kumulatif, lintas periode, atau bersih | Bila dibuka, aplikasi menghitungkan batas atas dana pribadi anggota yang dana pribadinya hanya berasal dari dana bersama |
| T-21 | Uang antara dua rumah tangga yang sama-sama memakai aplikasi (Q-86) | Dua catatan terpisah tanpa tautan | Tautan apa pun adalah saluran antar-rumah tangga dan memerlukan tinjauan `security-privacy` |
| T-22 | Siapa yang menetapkan jumlah jatah pribadi | Bukan lagi asumsi: keputusan pemilik produk, 11 November 2026 ("a di sepakati berdua, nanti saling konfirmasi aja"). Modelnya di §3.4 | — |
| T-23 | Rentang hari mulai periode anggaran | 1 sampai 28; rumah tangga yang penghasilannya tiba tanggal 29 sampai 31 memilih 28 atau 1 | Bila 29 sampai 31 harus dapat dipilih, perlu aturan "hari terakhir bulan bila tanggal itu tidak ada", dan periode tidak lagi selalu mulai pada tanggal yang sama |
| T-24 | Bentuk peralihan saat hari mulai diganti | Satu periode pendek (1 sampai 30 hari) dengan batas anggaran dan jatah pribadi penuh, tidak diprorata | Bila satu periode panjang: sampai sekitar dua bulan, memuat dua kali pembayaran bulanan terhadap batas satu bulan. Bila diprorata: pembagian menghasilkan pecahan dan aturan pembulatan baru |
| T-25 | Siapa yang boleh mengganti hari mulai periode | Putusan lead, 11 November 2026: kesepakatan rumah tangga, yaitu kedua anggota dewasa, seperti nilai ambang; terlihat oleh keduanya. Belum dinyatakan pemilik produk dengan kata-katanya sendiri | Bila satu orang boleh sendiri: ia dapat membuka periode pendek baru, dengan batas anggaran dan jatah pribadi penuh, kapan saja; penggantiannya tetap terlihat |
| T-26 | Dasar periode hitungan jatah pribadi | Ditetapkan dokumen 06 §6.3: tanggal kejadian, dinilai satu kali saat entri disimpan. Jalan menghindar yang diterima: tanggal mundur, yang hanya dilawan keterlihatan tanggal pencatatan (§3.4, §3.11) | Bila waktu simpan: penghasilan yang dicatat terlambat memakan jatah periode berjalan, dan angka jatah tidak sama dengan baris "disisihkan" ringkasan mana pun |
| T-27 | Sumber untuk penghasilan lepas (freelance) | `business`, sehingga ia *penghasilan* dan selalu tercatat di buku rumah tangga (INV-24); ditanyakan sebagai dokumen 17 S2-1 butir (iv) | Bila dicatat sebagai `other`, ia boleh masuk buku pribadi saja dan tidak terlihat; celah label §10.3 butir 1 berlaku pada setiap uang masuk pekerjaan lepas, bukan satu kali |

Butir dokumen 17 yang dapat ditutup atau diperbarui: K-02, K-11, K-13 (selesai lewat dokumen ini dan glosarium); Q-11 sampai Q-21, Q-43 sampai Q-48 (ditetapkan dokumen 07, menunggu tinjauan pemilik di checkpoint S2); Q-22 (usulan, menunggu ADR-F dan L-04).

### 10.2 Permintaan

**Kepada `architect`.**

1. ADR-A sampai ADR-G sudah ditulis sebagai ADR-002 sampai ADR-008. Kalimat yang kini tidak lagi sesuai dengan dokumen ini:
   - ADR-003, Konsekuensi: "Celah yang tidak ditutup model: uang masuk yang langsung dicatat sebagai dana pribadi tidak terlihat (07 §10.3 butir 1; K-16). Keputusannya milik pemilik produk di checkpoint S2 (dokumen 06 §6.4)." Sudah diputuskan; penghasilan tidak dapat lagi dicatat hanya di buku pribadi (INV-24). Yang tersisa adalah celah label sumber (§10.3 butir 1).
   - ADR-003, Keputusan: perlu memuat aturan tulis buku pribadi (INV-25), dan baris "Rincian" perlu menyebut INV-24 dan INV-25.
   - ADR-002, tabel Opsi: "Dua tabel dan daftar jenis tertutup". Kini ada tabel ketiga yang hanya ditambah, `entry_link`; dan jenis `income` boleh memuat baris di dua buku.
   - ADR-004, Keputusan: "dihitung dari baris dan dari otorisasi serta reservasi yang berlaku". Realisasi kini juga bergantung pada tautan susulan yang berlaku.
   - ADR-005, Konsekuensi: "jumlah entri, baris, otorisasi, reservasi, peristiwa, dan notifikasi tidak bertambah". Tambahkan tautan susulan.
   - ADR-006, ADR-007, ADR-008: tidak ada kalimat yang berubah.
2. Dokumen 06 memakai kait berikut dan tidak mendefinisikan ulang akibat uangnya: `request.approved`, `authorization.amended`, `authorization.settled`, `authorization.cancelled`, `authorization.expired`, `authorization.overrun`, `handover.recorded`, `handover.confirmed`, `handover.disputed`, `entry.disputed` (§3.3, §3.8, §3.12). Jumlahnya sepuluh. Transisi dokumen 06 yang **tidak** memicu kait uang: penutupan keberatan atas baris custody tanpa entri pembalik (C05), yang juga bukan konfirmasi (§3.3); `handover.confirmed` hanya dipicu C02 dan C03.
3. Dokumen 06 memutuskan: jenis permintaan untuk pemakaian dana bersama yang bukan pengeluaran (pinjaman, pembelian aset, penyisihan); apakah kelebihan yang dibayar dari dana bersama memerlukan tanggapan; penandaan pemakaian di atas ambang yang dicatat tanpa permintaan; tanggal yang menentukan periode *terikat* suatu otorisasi (§6.1 memakai tanggal rencana, atau tanggal persetujuan bila tidak ada).
4. Syarat dari model untuk dokumen 06 dan 10: tempat uang berpemilik `shared` tidak dapat bervisibilitas `private`; entri buku pribadi hanya dapat dikoreksi pemiliknya; hanya pemilik piutang yang dapat melepas klaim.
5. Jawaban atas kedelapan permintaan dokumen 06 §15.4 ada di tabel §9. Yang perlu disesuaikan di dokumen 06 sesudahnya: (a) §6.2 mengenali penyisihan dari jenis entri; bagian pribadi di dalam entri `income` adalah baris `owner_transfer`, jadi syaratnya perlu membaca baris, bukan jenis entri (T-13). (b) §6.4 sudah dijawab pemilik produk: hasilnya pilihan 3 dalam keterlihatannya, dengan bentuk satu entri. (c) selesai: dokumen 06 §6.2 dan §6.3 kini menetapkan bahwa pelunasan tidak ditandai dan cicilan tidak memerlukan kesepakatan rutin; dokumen ini diselaraskan (§2.6, §3.8, §3.14). (d) §8.1 "entri dua buku: hanya pemilik buku pribadinya" dipertegas di §3.12: itu berlaku atas entri pembalik; penyesuaian sisi rumah tangga tetap terbuka bagi keduanya. (e) §7.3: aturan dikukuhkan; tambahkan pengecualian pelepasan klaim dan penundaan entri antara dua buku pribadi.
6. Jawaban atas keempat permintaan baru dokumen 06 §15.4, butir (a) sampai (d), ada di tabel §9; keempatnya diterapkan atau dikukuhkan. Hal yang semula diserahkan kepada dokumen 06 (kewajiban yang didaftarkan sebagai saldo awal sesudah rumah tangga mulai mencatat) sudah diputuskan di dokumen 06 §6.2, putusan 2, dan diterapkan di §3.8 dan §3.14.
7. Jawaban atas permintaan dokumen 06 §15.4 butir (e) sampai (m) ada di tabel §9. **Selesai** (diperiksa pada dokumen 06 yang berlaku saat revisi ini): §6.2 kini memuat kedua pengecualian `contribution_expense` dan `advance`, mengukuhkan baris `expense` di dalam `recategorization`, dan menyatakan ketimpangan `loan_received` sebagai disengaja; §6.4 kini berbunyi menurut pernyataan bahwa kedua pasangan berpenghasilan; kalimat "satu-satunya pengecualian" atas INV-25 tidak ada lagi di dokumen 06; jatah pribadi kumulatif per periode ada di §6.3. Tidak ada yang tersisa dari butir ini.
8. Sisi uang temuan pemeriksaan awal `security-privacy` sudah diterapkan di dokumen ini (tabel §9). Di dokumen 06, SP-01, SP-05, SP-13, dan SP-22 **selesai**, begitu pula ADR-003 (INV-25 dengan dua pengecualian). Yang masih perlu diperiksa `architect`:
   - SP-12: ADR-007 keputusan 3 dan ADR-002 belum menyebut entri sistem `exit_closure` (§8.2, INV-28): penghapusan buku pribadi tidak bersyarat pada penyelesaian klaim oleh anggota; jenis entri bertambah satu; `actor` boleh `system` hanya pada jenis itu.
   - SP-18: siapa yang boleh mengubah *daftar sumber penghasilan* dan lewat alur apa (usulan: kesepakatan rumah tangga sebagai bagian preset aturan, dokumen 06 §6.3 "Kesepakatan rumah tangga"); dokumen 06 belum menyebut daftar itu.

9. Keputusan pemilik produk 11 November 2026 tentang jatah pribadi dan hari mulai periode anggaran. Yang disediakan dokumen ini, dan yang perlu ditetapkan dokumen 06:
   - Sudah sama dengan dokumen 06 §6.2 dan §6.3 "Kesepakatan rumah tangga", bagian "Jatah pribadi", dan menjawab §15.4 butir (n) sampai (s): (n) setiap penyisihan kepada M terhitung dengan sendirinya, tanpa tautan; jumlahnya per periode anggaran (§3.4). (o) Kolom jumlah pemakaian tabel §3.8 tidak berubah; jumlah pemakaian entri `owner_transfer` dan bagian pribadi entri `income` seluruhnya penyisihan (§3.4). (p) Kedua sifat periode dijamin INV-29; penggantian hari mulai adalah kesepakatan rumah tangga dan terlihat (§6.1). (q) CTH-39. (r) Dasar yang sama dengan tabel A; entri pembalik mengurangi baris asalnya (§3.4, §6.3). (s) `recorded_at` disimpan server; tanggal pencatatan (tanpa jam) tampil hanya pada entri yang tanggal pencatatannya berbeda dari tanggal kejadiannya, dan jam pencatatan tidak ditampilkan kepada siapa pun (§3.11; SP-28).
   - Yang masih perlu di dokumen 06: §6.3 mendaftar isi kesepakatan `rule` sebagai nilai ambang dan preset aturan. Menurut putusan lead, **hari mulai periode anggaran** masuk daftar itu (K01 sampai K05; hari mulai lama berlaku sampai versi baru disetujui), dengan guard 1 sampai 28.
   - Nama: `luar_jatah` di sini adalah `lebih` di dokumen 06, dan `terpakai_sebelum` adalah `jumlah_sebelum`; rumusnya setara (§3.4).
   - §4.5 "Masa berlaku dan kedaluwarsa", baris "Pembayaran atas kewajiban yang sudah tertaut penuh": kini juga bagian dari INV-26.
   - ADR dan bagian dokumen 06 yang masih berbunyi "mulai tanggal gajian" atau "jumlah per kali" untuk penyisihan perlu mengikuti.

**Kepada lead (glosarium di luar bagian 7).**

1. Bagian 4, *permintaan penggantian*: sesuaikan dengan keputusan pemilik nomor 3. Usulan bunyi: "Meminta dana bersama mengganti uang pribadi yang dipakai untuk keperluan bersama dan ditandai *minta diganti*."
2. Bagian 4, *persetujuan* dan *kedaluwarsa*: rujuk §3.8 untuk akibat uangnya.
3. Bagian 3, *data pribadi anggota*: tambahkan bahwa buku pribadi (dana pribadi beserta entri, piutang, kewajiban, dan asetnya) termasuk di dalamnya, kecuali sisi yang menyentuh dana bersama (§5.3).
4. `_progress.md`: temuan N-14, K-11, K-13 ditutup oleh dokumen ini.
5. Dokumen 17: T-12 sampai T-16 sebagai butir baru; K-16 dan Q-56 dijawab (§9); T-08 bukan lagi asumsi.
6. Glosarium bagian 4, *rujukan induk* (subject reference) belum menjadi istilah; dokumen ini merujuknya ke dokumen 06 §2.2.
7. Glosarium bagian 4, *konfirmasi custody*: definisinya hanya menyebut penerima. Tambahkan sisi pembayar: catatan entri dua buku yang mengurangi tempat uang dana bersama yang dipegang anggota dewasa selain pencatatnya juga bertanda *belum dikonfirmasi* (§3.3).
8. Glosarium bagian 3, *keluar dengan aman*: rujuk §8.2 untuk klaim antar-pemilik yang ditutup sistem saat keluar.
9. Dokumen 17: T-17 sampai T-21 sebagai butir baru; INV-27 menjadi pertanyaan pemilik produk di checkpoint S2 bersama S2-1 (keduanya saling bergantung, §3.1); sisi hukum `exit_closure` masuk daftar tinjauan hukum bersama L-04; Q-86 dijawab sebagai **[Usulan]** (§2.2).
10. CTH-38 adalah contoh berdiri sendiri baru untuk skrip hitung ulang; tidak ada angka CTH-01 sampai CTH-37 yang berubah.
11. Penyelarasan terakhir (dokumen 06 §15.4 butir (e) sampai (m); kedua pasangan berpenghasilan) tidak mengubah satu angka pun di tabel mana pun. Rujukan lama ke nomor baris dokumen 17 dicabut: rumusan bersyarat yang tersisa di sana ("hanya satu orangnya menerima gaji", "tidak berpenghasilan") dinilai dokumen 17 bagian 5 tetap benar sebagai kalimat umum, dan kalimat satu pencari nafkah di dokumen 06 dan dokumen ini sudah diperbaiki (K-27 butir 1, selesai). Tidak ada yang tersisa dari butir ini.

12. Revisi jatah pribadi dan periode anggaran: CTH-39 adalah contoh berdiri sendiri baru untuk skrip hitung ulang (dua tabel, 17 baris); tabel peta kasus §7.1 bertambah satu baris; tidak ada angka CTH-01 sampai CTH-38 yang berubah. Glosarium bagian 4, *kesepakatan rutin*: untuk jatah pribadi jumlahnya per periode anggaran, bukan per kali. `_progress.md` dan dokumen 17 (Q-06, Q-98; T-22 sampai T-27) masih memuat "mulai tanggal gajian" dan jatah sebagai usul yang belum dijawab.

**Kepada `security-privacy`.** Jatah pribadi dan periode peralihan (§3.4, §6.1; §10.3 butir 7). Kebocoran B1–B5 (§5.3); risiko §10.3 butir 1 dan 2; aturan tulis buku pribadi dan batasnya atas koreksi (§2.2, §3.12); usulan §8 beserta pengecualiannya terhadap INV-09.

### 10.3 Yang menurut dokumen ini berisiko dalam keputusan pemilik produk

Dinyatakan terus terang, sesuai brief §17. Tidak satu pun dibatalkan dokumen ini; semuanya dijalankan sebagaimana diputuskan.

1. **Gaji ke dana pribadi kini terlihat, tetapi tiga hal tetap terbuka** (keputusan pemilik produk, 11 November 2026, bersama keputusan nomor 2 dan 3).
   - *Celah label.* INV-24 menjamin gaji yang **dicatat sebagai gaji** tidak dapat tersembunyi. Sumber diisi pencatat. Gaji yang dicatat sebagai "hadiah" atau "lainnya" langsung di buku pribadi, atau tidak dicatat, tetap tidak terlihat, dan model tidak dapat menutupnya tanpa membuka seluruh pemasukan dana pribadi. Yang tersisa bagi pasangan hanyalah bahwa gaji bulan itu tidak muncul (B5).
   - *Tidak ada batas atas bagian pribadi.* Pencatat boleh mengarahkan seluruh gaji ke dana pribadinya; bila INV-27 diterima, hanya penerima gaji itu sendiri. Pasangannya melihatnya dan dapat berkeberatan, tetapi tidak dapat membatalkannya: entri itu entri dua buku (§3.12). Keterlihatan adalah satu-satunya pagar; kesepakatan tentang berapa yang wajar terjadi di luar aplikasi, atau lewat kesepakatan rutin bila usulan T-13 diterima.
   - *Simetris dalam aturan; setara dalam praktik hanya sejauh penghasilannya setara.* Kedua pasangan di rumah tangga pilot berpenghasilan (pernyataan pemilik produk, 11 November 2026). Aturannya per penerima penghasilan: masing-masing mencatat penghasilannya sendiri dan dapat mengarahkan bagian dari penghasilannya sendiri, tidak dari penghasilan pasangannya (INV-27, **[Usulan]**). Yang tersisa sebagai risiko: besarnya yang dapat diarahkan mengikuti besarnya penghasilan masing-masing; dan bila bagian pribadi dari penghasilan **tidak** ikut ditandai (lawan dari T-13), aturan ambang atas dana pribadi makin longgar bagi yang berpenghasilan lebih besar. Pada rumah tangga yang hanya satu anggotanya berpenghasilan (mungkin ada di tahap publik), hak mengarahkan dalam kenyataan dimiliki satu orang dan anggota yang lain hanya memperoleh dana pribadi lewat penyisihan.
2. **Laporan "siapa membayar apa" mengukur siapa yang memegang uang, bukan siapa yang berbelanja.** Pada CTH-37 suami tercatat membayar 6031000 dan istri 930000 hanya karena gaji tiba di rekeningnya dan ia yang membayar tagihan dan kiriman. Kriteria sukses nomor 3 meminta laporan ini; tanpa keterangan ia mudah dibaca sebagai penilaian. Hal yang sama berlaku untuk "kontribusi dari dana pribadi per anggota", yang dapat menjadi papan skor.
3. **Selisih kas melekat pada pemegang** (keputusan nomor 7). "Siapa yang lupa diselesaikan di luar aplikasi", tetapi selisih tercatat per tempat uang dan tiap tempat uang punya pemegang, sehingga aplikasi tetap memperlihatkan di tangan siapa uang yang tidak tercatat itu. Model tidak dapat menghindarinya; dokumen 08 dapat memilih kata yang tidak menuduh.
4. **Ambang permintaan adalah kesepakatan, bukan kunci** (keputusan nomor 5 bersama A-20). Aplikasi tidak dapat mencegah pemakaian di atas ambang; ia hanya membuatnya terlihat. Kriteria sukses nomor 4 bergantung pada kebiasaan, bukan pada aturan sistem.
5. **Satu tempat uang, satu pemilik** adalah usulan dokumen ini, bukan keputusan pemilik, tetapi ia akibat langsung dari keputusan nomor 3. Bila pasangan tidak memisahkan uang pribadi secara fisik, dana pribadi menjadi angka yang tidak pernah dapat direkonsiliasi (T-03).
6. **Kontribusi yang terlihat membocorkan dana pribadi** (B2, B4). Ini akibat yang dipilih keputusan nomor 3 dan diterima; anggota perlu diberi tahu sebelum mencatat.
7. **Jatah pribadi adalah pengecualian tanda, bukan batas.** Yang diputuskan pemilik produk (11 November 2026) hanyalah bentuk jatah: satu jumlah per orang per periode, dijumlahkan sepanjang periode, dapat diatur, disepakati berdua, kenaikan diajukan dengan alasan. Bahwa di dalam jatah tidak ada tanda adalah **[Usulan]** lead yang disampaikan kepadanya dan tidak ia bantah (dokumen 17 D-48); bahwa hanya bagian di luar jatah yang dibandingkan dengan ambang adalah **[Usulan]** rancangan dokumen ini dan dokumen 06, terbuka di checkpoint S2 (dokumen 17 S2-1 butir (iii), pilihan (a); pilihan (b): yang melampaui jatah ditandai berapa pun jumlahnya). Mengambil di atas jatah tidak dihalangi pada pilihan mana pun. Risiko di bawah berlaku atas pilihan (a); pada pilihan (b), jalan (a) di bawah tertutup bagi bagian di luar jatah. Tiga jalan yang tetap terbuka dan hanya dijaga keterlihatan: (a) bagian di luar jatah yang dipecah menjadi beberapa entri, masing-masing tidak di atas ambang, tidak pernah ditandai; (b) karena periode dihitung menurut tanggal kejadian, penyisihan dapat diberi tanggal periode lampau yang jatahnya belum terpakai; yang melawannya hanya tanggal pencatatan yang tampil pada entri terlambat (§3.11); (c) mengganti hari mulai periode membuka satu periode pendek dengan jatah dan batas anggaran penuh; karena itu penggantiannya kesepakatan rumah tangga dan terlihat (§6.1, T-25), tetapi dua orang yang sepakat tetap dapat melakukannya berulang kali. Selain itu, penyisihan kini dinilai bersama penyisihan lain di periode yang sama, tidak lagi sendiri-sendiri; itu perubahan atas sifat tanda di dokumen 06 (INV-A17).
8. **Makin sering penghasilan datang, makin besar celah label.** Dengan uang masuk pekerjaan lepas yang "bisa sampe 5x" (kata-kata pemilik produk; "sebulan" adalah bacaan lead, §3.1), setiap uang masuk adalah kesempatan memberinya sumber "lainnya" dan mencatatnya di buku pribadi saja (butir 1, T-27). Model tidak dapat menutupnya; pasangan hanya melihat bahwa penghasilan itu tidak muncul.

