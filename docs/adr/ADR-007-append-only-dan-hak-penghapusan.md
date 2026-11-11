# ADR-007 — Tafsir append-only terhadap hak penghapusan (sementara)

| | |
|---|---|
| Status | **Diterima sebagai keputusan sementara.** Keputusan pemilik produk, checkpoint S2 (dokumen 17 S2-8 pilihan A): bacaan ini dipakai sampai advokat menjawab. Tetap sementara sampai butir tinjauan hukum L-04 dan L-20 terjawab (pertanyaan tertulisnya: L-22, L-23). Butir tambahan 6 sampai 9 tercakup penerimaan yang sama; yang bertanda **[Usulan]** atau **[Perlu validasi]** tetap demikian. Butir 10 (12 November 2026, tahap S3) **[Usulan]**; ditinjau `critic` dan `security-privacy` dan direvisi pada putaran perbaikan fase D yang sama (dokumen 06 revisi kesepuluh): waktu penutupan pindah ke akhir masa tenggang keluar (SP-48), butir 7 menyebut kartu status di chat (SP-45), dan konsekuensi tentang pesan chat diperbaiki karena bertentangan dengan dokumen 06 §16.2 dan F-FAM-006 (CB-12) |
| Tanggal | 11 November 2026 |
| Pengusul | `finance-domain` (rekomendasi, 07 §1.5 ADR-F); ditulis `architect`; memerlukan masukan `security-privacy` di S4 |
| Rincian | [07-financial-domain.md](../07-financial-domain.md) §8, INV-09; [17-open-questions-assumptions.md](../17-open-questions-assumptions.md) Q-22, K-10, L-04, L-20; [18-market-research.md](../18-market-research.md) §2.1 |

ADR ini bukan nasihat hukum. Bunyi peraturan dikutip dari dokumen 18 dengan label buktinya.

## Konteks

Dua tuntutan bertemu:

- Brief §7 dan `ENGINEERING-STANDARDS.md` §5: riwayat keuangan hanya ditambah; koreksi adalah entri pembalik atau penyesuaian.
- UU PDP Pasal 43–44 mewajibkan penghapusan dan pemusnahan atas permintaan subjek data; pengecualiannya (Pasal 50) tidak mencakup keutuhan catatan yang dimiliki bersama orang lain (**[Terverifikasi]** primer, dokumen 18 §2.1, diakses 11 November 2026). Salinan teks PP 33/2026 mengukur pemusnahan dari hasilnya, "tidak lagi dapat digunakan untuk mengidentifikasi Subjek Data Pribadi" (Pasal 84 ayat (2)), menjangkau salinan cadangan, dan tidak menyebut pseudonimisasi sebagai cara memenuhi kewajiban hapus (**[Perlu validasi] — salinan teks**).

Catatan dana bersama adalah catatan dua orang. Menghapus entri seorang anggota mengubah saldo, laporan, dan jejak audit anggota yang tersisa.

## Opsi

| | A. Tidak ada yang pernah dihapus | B. Entri anggota dihapus atas permintaan | C. Inti moneter tetap dan dipseudonimkan; rincian dan buku pribadi dapat dihapus | D. Entri penutup: buku dirangkum menjadi saldo awal baru, entri lama dihapus |
|---|---|---|---|---|
| Hak penghapusan | Tidak dipenuhi | Dipenuhi | Dipenuhi sebagian; cukup-tidaknya L-04, L-20 | Dipenuhi |
| Angka anggota yang tersisa | Utuh | Rusak | Utuh | Utuh |
| Jejak audit anggota yang tersisa | Utuh | Rusak | Utuh, pelaku menjadi "mantan anggota" | Hilang sebelum tanggal rangkuman |

## Keputusan

**Opsi C sebagai keputusan sementara, dengan opsi D sebagai jalan cadangan yang sudah dinyatakan.**

1. Entri dipisah sejak awal menjadi tiga bagian (07 §8.2): **inti moneter** (hanya ditambah; pelaku dan pihak sebagai pengenal internal), **rincian** (catatan teks, struk; dapat dihapus tanpa menyentuh inti moneter), dan **pemetaan identitas** (dapat diputus).
2. "Append-only" di `ENGINEERING-STANDARDS.md` §5 dibaca sebagai berlaku atas **inti moneter**. ADR ini mengusulkan bacaan itu sebagai penjelasan resmi atas standar tersebut; standarnya sendiri tidak dilonggarkan untuk koreksi atau penyuntingan.
3. **Satu-satunya pengecualian** terhadap 07 INV-09: atas permintaan penghapusan yang sah dari seorang anggota, **buku pribadinya** ditutup dan dihapus seluruhnya, setelah ekspor ditawarkan dan setelah setiap klaim antar-pemilik yang masih terbuka diakhiri lewat entri yang terlihat. Buku pribadi bukan catatan bersama. Sisi rumah tangga dari entri dua buku tetap dan tetap seimbang per pemilik.
4. Di buku rumah tangga tidak ada baris yang dihapus: pemetaan identitas anggota itu diputus dan rincian yang ia tulis dihapus, dengan batas yang dinyatakan di 07 §8.3.
5. Penghapusan rincian dan pemutusan pemetaan menjangkau salinan cadangan dalam jangka yang ditetapkan dokumen 10.

**Tambahan 11 November 2026 (temuan `security-privacy` SP-11 dan SP-12).** Statusnya sama dengan butir 1 sampai 5: diterima sebagai keputusan sementara (checkpoint S2); sisi hukumnya **[Perlu validasi]** (L-04, L-20).

6. **Catatan alur mengikuti pemisahan tiga bagian yang sama** (dokumen 06 §2.5). *Inti alur* hanya ditambah: permintaan, versi syarat sejauh menyangkut jumlah dan tanggal, tanggapan, Peristiwa, kesepakatan, status konfirmasi custody, keberatan beserta alasan dari daftar tertutup; pelaku dan pihak sebagai pengenal internal. *Rincian* dapat dihapus tanpa menyentuh inti: judul, catatan, tempat, peserta, alasan teks bebas, pesan diskusi, lampiran. *Pemetaan identitas* adalah tabel yang sama dengan ledger dan diputus bersama. "Peristiwa tidak pernah dihapus" (06 INV-A01) dibaca atas inti alur.
7. **Catatan `CARE` berada di luar aturan hanya-menambah.** Tindak lanjut sakit, status suasana hati, dan Peristiwa keduanya dapat dihapus pemiliknya kapan saja, seluruhnya, tanpa sisa dan tanpa pemberitahuan kepada siapa pun. Catatan itu bukan catatan bersama; tidak ada angka atau hak orang lain yang bergantung padanya. Karena itu Peristiwanya disimpan terpisah dari Peristiwa alur lain. Pengecualian ini mencakup kartu status di chat rumah tangga: kartu itu hilang dari chat bagi kedua orang saat statusnya diganti, berakhir, dicabut, atau dihapus, tanpa penanda (06 §13.5, §16.7, INV-A47; ADR-011 butir 10; revisi 12 November 2026, SP-45).
8. **Keluar dan penghapusan tidak bersyarat pada tindakan orang lain.** Butir 3 semula mensyaratkan setiap klaim antar-pemilik yang terbuka "diakhiri lewat entri yang terlihat"; bila pengakhiran itu memerlukan tindakan pasangan, pasangan dapat menahan kepergian. Kini: klaim antar-pemilik yang masih terbuka ditutup saat keluar (sejak 12 November 2026: pada akhir masa tenggang, butir 10) oleh entri berjenis `exit_closure` ("ditutup karena anggota keluar"; ADR-002), yang hanya ditulis sistem, tepat satu kali per anggota yang keluar, tidak dapat dibalik, yang tidak menyatakan ada pembayaran dan tidak memerlukan tindakan siapa pun (aturan uangnya di 07 §8.2, yang direvisi `finance-domain`). Objek alur yang hanya dapat ditutup anggota yang keluar ditampilkan demikian dan tidak lagi masuk daftar terbuka (06 §3.5, INV-A35); tidak ada yang ditutup atas namanya.
9. **Jatah pribadi** (tambahan atas temuan `security-privacy` SP-25 dan SP-31; **[Usulan]**). Alasan pada usulan kenaikan jatah, walau dipilih dari daftar tertutup, termasuk *rincian* dan dihapus bersama rincian lain saat anggota keluar atau meminta penghapusan; jumlah jatah dan kesepakatannya termasuk inti. *Catatan penilaian* (06 §6.3) termasuk inti alur, hanya ditambah, dan pelakunya disimpan sebagai pengenal internal yang mengikuti pemetaan identitas saat anggota keluar; ia hanya dibaca per entri.

**Tambahan 12 November 2026 (tahap S3; 07 §8.2 butir 2a, T-37, INV-28; menjawab 07 §10.2 putaran S3 butir 7, PA-FAM-7).** **[Usulan]** `finance-domain` dan `architect`; belum disajikan kepada pemilik produk.

10. **Menutup dan menghapus adalah dua langkah.** *Menutup* buku pribadi terjadi pada **setiap** keluarnya anggota, dengan atau tanpa permintaan penghapusan: dalam satu transaksi dengan `exit_closure`, buku itu berhenti menerima entri (tidak ada pelaku yang berhak menulisnya lagi) dan tetap dapat diekspor pemiliknya. **Waktunya** (revisi 12 November 2026, SP-48; **[Usulan]**; pertanyaan checkpoint S3-4 butir (c)): pada akhir masa tenggang keluar (06 §3.6 F07, 7 hari **[Asumsi]** sesudah perintah keluar), bukan saat perintah keluar; seketika bila yang keluar sendiri memilih hapus seketika. Selama masa tenggang aksesnya sudah berakhir tetapi belum ada langkah yang merusak, dan hanya ia yang dapat membatalkan keluarnya (F08). Penghapusan catatan `CARE`-nya (butir 7) mengikuti waktu yang sama. *Menghapus* buku itu (butir 3) terjadi atas permintaan penghapusan yang sah, atau pada akhir jangka retensi yang ditetapkan dokumen 10 bila ia tidak memintanya. Butir 3 dibaca dengan urutan itu: penutupan selalu lebih dulu, penghapusan menyusul. Buku rumah tangga tidak tersentuh oleh keduanya (butir 4). Jangka retensi, cara mantan anggota mengekspor, dan cadangannya **[Dok. 10]**.

**Jalan cadangan.** Bila tinjauan hukum menyatakan inti moneter buku rumah tangga pun harus dimusnahkan, dipakai opsi D: buku rumah tangga dirangkum menjadi saldo awal baru per posisi, lalu entri lama dihapus. Biayanya: jejak audit dan laporan sebelum tanggal rangkuman hilang bagi anggota yang tersisa.

## Konsekuensi

- Pemisahan inti moneter, rincian, dan pemetaan identitas harus ada sejak skema pertama; menambahkannya kemudian berarti migrasi atas seluruh riwayat. Ini satu-satunya bagian ADR ini yang tidak dapat ditunda.
- Model harus mampu menjalankan opsi D tanpa rancang ulang: saldo awal per posisi sudah ada sebagai jenis entri (07 §3.17), dan invarian sudah mengizinkan entri pertama berupa saldo awal.
- Di rumah tangga dua orang, "mantan anggota" tetap dikenali pasangannya; pseudonimisasi hanya memutus kaitan bagi operator dan pihak lain (07 §8.3). Apakah itu cukup adalah isi L-04 dan L-20. **Pada pilot, pemilik produk adalah operator server sekaligus anggota** (catatan SP-17): "memutus kaitan bagi operator" tidak berarti apa pun di sana, dan larangan peran (06 V-5) mengikat peran aplikasi, bukan akses basis data. Batas ini dinyatakan dan diserahkan ke dokumen 10. Pemilik produk menerima keadaan itu untuk pilot dengan syarat istrinya diberi tahu dengan terus terang sebelum mulai memakai (keputusan pemilik produk, checkpoint S2; dokumen 17 S2-10 pilihan A); apakah pemberitahuan itu sudah terjadi belum diketahui.
- Test tambahan untuk butir 6 sampai 8: hash inti alur identik sebelum dan sesudah penghapusan rincian; setelah pemilik menghapus catatan `CARE`-nya tidak ada baris atau Peristiwa yang tersisa; perintah keluar tidak pernah ditolak karena klaim atau objek alur yang terbuka (06 INV-A35).
- Teks bebas dan foto struk pada entri dana bersama dapat memuat data orang yang pergi dan tidak dapat dipilah otomatis; diperlakukan sebagai data bersama dan dinyatakan sebagai batas.
- Test: hash inti moneter entri lama identik sebelum dan sesudah perintah apa pun, kecuali perintah penghapusan buku pribadi, yang punya test sendiri (angka rumah tangga tidak berubah; INV-02 tetap berlaku di buku rumah tangga).
- **Yang berubah bila jawabannya tidak:** opsi D menggantikan butir 4; ADR-006 mendapat satu bentuk penutupan; dokumen 10 dan 12 (skenario "anggota keluar", "pasangan berpisah") ditulis ulang untuk itu.
- Tahap publik tidak dimulai sebelum L-04 dan L-20 terjawab. Pilot dua orang berjalan dengan opsi C.
- Test untuk butir 10: selama masa tenggang tidak ada `exit_closure` dan buku pribadi belum ditutup; sesudah F07 tanpa permintaan penghapusan, setiap perintah tulis atas buku pribadi anggota itu ditolak dan ekspornya tetap tersedia baginya; sesudah jangka retensi berakhir, buku itu tidak lagi ada; angka buku rumah tangga identik di keempat titik.
- **Pesan chat** (dokumen 06 §16; ADR-011) adalah rincian: penulisnya dapat menghapusnya kapan saja. **Saat ia keluar tanpa meminta penghapusan, pesannya tidak dihapus**: percakapan ditutup dan tetap terbaca oleh bekas pesertanya yang tersisa (06 §16.2; 04 F-FAM-006), dan tidak pernah terbaca oleh anggota yang bergabung sesudahnya (06 §16.2, INV-A47 (f); SP-49). **Atas permintaan penghapusan yang sah**, isi dan fotonya dibuang seperti rincian lain menurut butir 4, dengan penanda yang sama seperti penghapusan oleh penulisnya; jangka dan cadangannya dokumen 10. Rumusan sebelumnya ("saat ia keluar atau meminta penghapusan pesannya diperlakukan seperti rincian lain") dapat dibaca bahwa keluar saja menghapus pesannya, yang bertentangan dengan 06 §16.2; ia diganti (CB-12).
- Tindak lanjut S4: mekanisme penghapusan pada penyimpanan dan cadangan (dokumen 09), jangka retensi dan alur permintaan penghapusan (dokumen 10).
