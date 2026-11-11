# Log Kemajuan

Dicatat lead setelah setiap tahap (brief §0.2). Bila konteks terhapus atau sesi dimulai ulang: baca berkas ini dan [00-glossary.md](00-glossary.md), lalu lanjutkan dari **Tahap berikutnya**.

## Posisi saat ini

| | |
|---|---|
| Tahap terakhir yang selesai | S2 — Fondasi domain (11 November 2026): dokumen 02, 03, 06, 07, dan ADR-001 sampai ADR-010 ditulis, diperiksa, dan diperbaiki |
| Status | **Berhenti di checkpoint S2.** Menunggu pemilik produk menjawab sepuluh pertanyaan di [17-open-questions-assumptions.md](17-open-questions-assumptions.md) bagian "Checkpoint S2: pertanyaan untuk pemilik produk" (S2-1 sampai S2-10), mendengar peringatan W-1 sampai W-16, dan menerima atau mengubah default D-01 sampai D-55 |
| Tahap berikutnya | Setelah jawaban checkpoint S2: sebarkan jawabannya ke 06, 07, 17, glosarium, dan ADR (satu commit revisi), usulkan kalimat brief dan `CLAUDE.md` (S2-9), lalu S3 — Definisi produk (04, 05, 11, 12) setelah pemilik produk berkata lanjut |

**Cara melanjutkan di sesi baru:** baca berkas ini dan glosarium, lalu sajikan checkpoint S2 dari dokumen 17 kepada pemilik produk dalam bahasa sehari-hari: ringkasan model keuangan dan model persetujuan paling banyak satu halaman, sepuluh pertanyaan satu per satu, peringatan dibacakan apa adanya, dan default bertanda **Pagar** dibacakan satu per satu. Tidak ada tahap baru sebelum ia menjawab.

**Catatan proses dari pemilik produk (11 November 2026):** ia menyerahkan cara kerja tim kepada lead asal hasilnya profesional dan valid; semua hal dijelaskan dan dikonfirmasi kepadanya; keputusan ditanyakan langsung. Demi kuota, pekerjaan S3 dijalankan dengan model yang lebih hemat dan gambaran UI/UX dikerjakan di S4; setelah tiap tugas selesai dan aman, lead memberi tahu bahwa percakapan boleh dibersihkan.

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

**Bukan keputusan pemilik produk** (usulan yang menunggu jawabannya di checkpoint S2): di dalam jatah tidak ada tanda (D-48); perlakuan atas bagian di luar jatah (S2-1 iii); hanya penerima penghasilan yang boleh mengarahkannya (S2-1 i); hadiah untuk satu orang tidak terlihat (S2-1 ii); alasan kenaikan wajib atau tidak (S2-1 v); hari mulai 1 sampai 28 dan penggantiannya disepakati berdua (putusan lead, D-49); seluruh daftar D.

## S2 — Fondasi domain

### Dokumen

| Dokumen | Pemilik | Status | Catatan |
|---|---|---|---|
| [07-financial-domain.md](07-financial-domain.md) | `finance-domain` | ditinjau | Ledger entri bertipe; satu buku per pemilik uang; dana pribadi di luar garis kendali; INV-01 sampai INV-30; CTH-01 sampai CTH-39; butir terbuka T-01 sampai T-27 |
| [06-activity-approval-architecture.md](06-activity-approval-architecture.md) | `architect` | ditinjau | Alur domain eksplisit di atas enam primitif; INV-A, PG, N, V, AT; butir terbuka P-01 sampai P-17 |
| [adr/](adr/) ADR-001 sampai ADR-010 | `architect` | diusulkan | Belum disetujui pemilik produk; ADR-007 sementara (L-04, L-20) |
| [02-product-principles.md](02-product-principles.md) | `product-analyst` | ditinjau | PRN-01 sampai PRN-18, tiap prinsip dengan cara ujinya; ketegangan KT |
| [03-personas-household-models.md](03-personas-household-models.md) | `product-analyst` | ditinjau | PSN-01 sampai PSN-07; sepuluh struktur rumah tangga; KB-01 sampai KB-25 |
| [00-glossary.md](00-glossary.md) | lead; bagian 7 `finance-domain` | ditinjau (versi S2) | Istilah keuangan dan persetujuan difinalkan; usulan yang menunggu checkpoint berlabel **[Usulan]** |
| [17-open-questions-assumptions.md](17-open-questions-assumptions.md) | `product-analyst` | ditinjau | Keputusan 11 sampai 17; bahan checkpoint S2 |

### Tinjauan

| Tinjauan | Oleh | Hasil |
|---|---|---|
| Pemeriksaan awal dokumen 06 (hanya temuan) | `security-privacy` | SP-01 sampai SP-22: 2 menghalangi, 12 sebelum S3, 8 catatan untuk dokumen 10. Semua yang menghalangi dan sebelum-S3 dikerjakan pemilik dokumennya |
| Tinjauan S2 | `critic` | Belum siap: 1 menghalangi (usulan tentang jatah pribadi ditulis sebagai keputusan pemilik), 7 perlu diperbaiki, 9 minor. Semuanya dikerjakan |
| Pemeriksaan aturan jatah pribadi (hanya temuan) | `security-privacy` | SP-24 sampai SP-31: 2 menghalangi (mengakhiri jatah pasangan sendirian; bentuk alasan kenaikan), 3 sebelum S3, 3 catatan. Yang menghalangi dan sebelum-S3 dikerjakan |

**Yang tidak ditinjau ulang:** perbaikan atas temuan `critic` dan atas SP-24 sampai SP-31 diperiksa lead terhadap laporan pemilik dokumen dan lewat hitung ulang angka, **tidak** ditinjau `critic` untuk kedua kalinya. `architect` melaporkan membaca ulang dokumen 06 secara utuh hanya sebagian pada tiap revisi; `critic` membacanya utuh satu kali, sebelum putaran perbaikan terakhir. Tinjauan konsistensi berikutnya ada di S3 dan S6.

**Angka yang dihitung ulang lead dengan skrip (di luar repositori), setelah revisi terakhir dokumen 07:** 42 tabel saldo, 213 baris: jumlah berjalan tiap kolom, rumus `UANG` dan `KB`, kesinambungan rantai CTH-01 sampai CTH-36; tabel samping (alokasi, anggaran, reservasi, rekonsiliasi, CTH-02, CTH-17b, CTH-37 A/B/C, arus kas); CTH-03b, varian perubahan syarat CTH-22, CTH-34, CTH-38, CTH-39 termasuk tabel jatah; contoh urutan tiba pada W-13. Semuanya cocok. `critic` memeriksa dengan penalaran bahwa angka per kategori CTH-37 berasal dari entri satuannya dan bahwa baris entri seimbang per buku; itu tidak diperiksa skrip.

### Keputusan

ADR-001 sampai ADR-010, semuanya berstatus diusulkan. Keputusan proses lead: pembayaran cicilan tidak memerlukan kesepakatan rutin (pembelian secara cicilan dibandingkan dengan ambang satu kali); penggantian hari mulai periode disepakati berdua (menunggu pemilik produk); pertanyaan checkpoint dibaca dari dokumen 17 saja.

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
