# ADR-010 — Model persetujuan simetris

| | |
|---|---|
| Status | Diusulkan; menunggu persetujuan pemilik produk di checkpoint S2. Dua butir sudah diputuskan pemilik produk pada 11 November 2026 (jarak pengingat ulang dan jam tenang; konfirmasi serah terima) dan ditandai di tempatnya |
| Tanggal | 11 November 2026 |
| Pengusul | `architect`; diperiksa `security-privacy` sebelum checkpoint S2 |
| Rincian | [06-activity-approval-architecture.md](../06-activity-approval-architecture.md) §3, §4, §7, §10, §12; INV-A04 sampai INV-A10, INV-A13; PG-01 sampai PG-14 |

## Konteks

Brief §4.3 meminta lima konsep (hubungan, peran, izin, tanggung jawab, wewenang persetujuan) dimodelkan terpisah. Brief §5D dan §8 menjadikan pencegahan kendali koersif syarat utama. Keputusan pemilik produk di checkpoint S1: kedua pasangan setara penuh (nomor 4); permintaan tanpa tanggapan tetap menunggu dan pemohon dapat mengingatkan lagi (nomor 6); rencana dapat dirundingkan pada permintaannya sendiri (nomor 5). Pilot memakai satu preset tanpa layar pengaturan izin (Q-30). Model yang sama harus kelak menampung beberapa pemberi persetujuan, urutan, batas jumlah, delegasi, dan wali tanpa rancang ulang.

## Opsi

**Siapa yang menyetujui.**

| | A. Pemohon memilih pemberi persetujuan | B. Ditentukan peran (misalnya pemilik rumah tangga) | C. Ditentukan satu fungsi dari wewenang persetujuan, pemohon selalu dikecualikan |
|---|---|---|---|
| Simetris | Ya | Tidak bila peran berbeda | Ya |
| Dapat dihindari pemohon | Ya | Tidak | Tidak |
| Perluasan Tier 3 | Sulit | Mengikat wewenang ke peran, bertentangan dengan brief §4.3 | Daftar pemberi persetujuan yang lebih panjang |

**Apa yang terjadi bila tidak ada tanggapan.**

| | X. Dianggap setuju setelah tenggat | Y. Dianggap ditolak setelah tenggat | Z. Tetap menunggu; hanya dapat kedaluwarsa |
|---|---|---|---|
| Diam sebagai alat | Mengajukan saat pasangan tidak dapat menjawab | Menghukum pemohon | Tidak |
| Keputusan pemilik nomor 6 | Bertentangan | Bertentangan | Sesuai |

## Keputusan

**Opsi C dan Z**, dengan aturan berikut:

1. Pemberi persetujuan ditentukan satu fungsi di server dari `approval_authority`, dengan pemohon **selalu** dikecualikan, dan dibekukan pada permintaan saat diajukan. Pada preset "dua dewasa setara": pasangan pemohon.
2. Kelima konsep brief §4.3 disimpan terpisah sejak rumah tangga dibuat, walau pilot hanya punya satu preset.
3. Yang disetujui adalah **satu versi syarat**. Versi syarat tidak dapat diubah dan bernomor naik; usul balik adalah versi baru yang menggantikan versi sebelumnya. Suatu versi disetujui bila pemohon dan setiap pemberi persetujuan yang disyaratkan menyatakan setuju atas versi yang sama, masing-masing dengan perintahnya sendiri; setiap tanggapan menyebut nomor versi yang ditanggapinya.
4. **Diam tidak pernah menyelesaikan apa pun.** Tidak ada transisi berbasis waktu menuju setuju, tolak, berlaku, atau dikonfirmasi. Waktu hanya dapat membuat permintaan kedaluwarsa, atau menutup otorisasi yang masa berlakunya habis tanpa membuat entri.
5. Persetujuan tidak dapat dicabut sepihak; perubahan setelah disetujui adalah perubahan syarat yang diterima pihak lain.
6. Tidak ada peran yang dapat menyetujui, mengonfirmasi, atau menutup atas nama orang dewasa lain. Pelaku hanya anggota atau `system`; AI tidak pernah menjadi pelaku (brief §5O).
7. Pengingat berbatas. **Keputusan pemilik produk, 11 November 2026:** pengingat ulang paling banyak satu kali per 6 jam per permintaan, dan pengingat yang jatuh pada jam tenang penerima disampaikan setelah jam tenang berakhir. Batas lainnya tetap usulan `architect` dan diajukan di checkpoint S2 (dokumen 06 §10, PG-06, INV-A30 sampai INV-A33; dokumen 17 S2-5, D-20, D-29). Setelah pemeriksaan awal `security-privacy` (SP-02, SP-03, SP-27) usulannya berbunyi: jendela 6 jam berlaku atas semua pemberitahuan yang meminta tanggapan per objek per pengirim, bukan hanya pengingat, dan pengingat otomatis dihitung sebagai berasal dari pihak yang menunggu; objeknya satu rantai permintaan, atau untuk usulan kesepakatan satu sasaran kesepakatan; paling banyak tiga pengingat ulang per pengirim seumur objek, tidak dimulai lagi oleh usul balik, pengajuan ulang, atau usulan baru atas sasaran yang sama; pengingat otomatis paling banyak dua per rantai permintaan; plafon harian per pengirim kepada penerima (pilot: enam, **[Asumsi]**), selebihnya masuk ringkasan; kalimat baku; **tunda 24 jam yang senyap**: pihak yang menunggu tidak diberi tahu dan tidak dapat mengetahuinya, pengingat yang ditekan selama tunda diterima dan ditahan seperti pada jam tenang, dan kabar "nanti aku jawab" hanya terkirim bila penerima sendiri memilihnya. Tidak ada tanda terima: "sampai" berarti dilepas server.
8. Tidak ada aturan yang membaca hubungan, jenis kelamin, siapa yang berpenghasilan atau berpenghasilan lebih besar, atau siapa pemegang uang. Aturan yang menyangkut penghasilan berlaku per penerima penghasilan dan sama bagi anggota dewasa mana pun (dokumen 06 §6.4).
9. **Asas pengawasan** (diusulkan `security-privacy` untuk dokumen 10; **[Usulan]**; dokumen 06 PG-14): seorang anggota dewasa boleh sendirian mengurangi pengawasan aplikasi (tanda, pemberitahuan) atas siapa pun atau menambah pengawasan atas dirinya sendiri; menambah pengawasan atas orang dewasa lain memerlukan persetujuan orang itu. Penerapan pertamanya: *jatah pribadi* hanya dapat diakhiri sendirian oleh pemiliknya, sedangkan kesepakatan atas tagihan berulang dapat diakhiri salah satu pihak (SP-24; dokumen 06 §6.3 K06). Asas ini batas bawah: keputusan pemilik produk yang menuntut dua orang tetap berlaku. Ia belum selaras dengan pilihan A dokumen 17 S2-4 (tidak seorang pun keluar sendiri dari ambang); itu dibiarkan terbuka sampai pemilik produk menjawab.

## Konsekuensi

- Simetri menjadi property test: menukar kedua anggota dewasa pada rangkaian perintah acak menghasilkan state yang tertukar persis (INV-A13).
- Permintaan dapat menunggu tanpa akhir bila tidak bertanggal. Diterima: itulah keputusan pemilik nomor 6; pemohon dapat membatalkan kapan saja.
- Konfirmasi custody mengikuti aturan 4 dan 6. **Keputusan pemilik produk, 11 November 2026:** serah terima tercatat seketika dan bertanda "belum dikonfirmasi" sampai penerima mengetuk terima; penerima dapat mengajukan keberatan kapan saja; sama ke dua arah (dokumen 06 §7). Tidak ada konfirmasi otomatis. Menutup keberatan bukan konfirmasi: status kembali ke yang berlaku sebelum keberatan, dan "dikonfirmasi" hanya lahir dari pernyataan pemegangnya sendiri (SP-04; 06 INV-A22). Tanda yang sama berlaku atas entri dua buku yang mengurangi tempat uang dana bersama yang dipegang pasangan (SP-05; **[Usulan]**).
- Jarak 6 jam lebih rapat daripada usulan awal 24 jam. Karena permintaan tidak bertanggal menunggu tanpa akhir, batas jumlah pengingat menjadi lebih penting, bukan kurang; selama batas itu belum diputuskan, pagar terhadap desakan berulang hanya jarak 6 jam dan jam tenang.
- Pemberi persetujuan yang berubah pikiran tidak punya tombol "cabut". Diterima: pencabutan sepihak tidak membatalkan belanja yang sudah terjadi dan hanya berguna sebagai tekanan.
- Rumah tangga dengan satu anggota dewasa tidak punya pemberi persetujuan: pengajuan ditolak dengan `no_approver`, dan tidak ada jalan lain menuju disetujui. Ini berlaku pada pilot sampai pasangan kedua bergabung, dan pada orang tua tunggal di tahap publik. Diterima untuk pilot: selama itu tidak ada permintaan anggaran, ambang, tanda, atau konfirmasi custody, dan pencatatan berjalan penuh; saat anggota dewasa kedua bergabung tidak ada yang berlaku surut (dokumen 06 §3.5). Preset "satu dewasa" mencatat rencana tanpa persetujuan; bentuknya Tier 3.
- Serah terima tidak punya pencatat bawaan: pemberi mencatat dan penerima mengonfirmasi, atau penerima mencatat dan catatannya langsung terkonfirmasi (dokumen 06 §7.1). Usulan sebelumnya, penerima sebagai pencatat bawaan, ditarik karena alur yang diuraikan pemilik produk adalah pemberi yang mencatat.
- Versi syarat pada setiap tanggapan menutup kasus "persetujuan tiba untuk syarat yang sudah digantikan", di jaringan buruk maupun saat kedua pihak bertindak bersamaan (dokumen 06 §12). Teknologi sinkronisasinya diputuskan di S4.
- Perluasan Tier 3 (beberapa pemberi persetujuan, urutan, batas, delegasi, wali) mengisi struktur yang sudah ada: `required_approvals`, `approval_mode`, `approval_authority`.
- Batas akhir visibilitas dan apa yang tidak pernah boleh dilakukan peran admin ditetapkan `security-privacy` di dokumen 10; bila bertentangan dengan ADR ini, penyelesaiannya ADR baru dengan masukan agen itu.
