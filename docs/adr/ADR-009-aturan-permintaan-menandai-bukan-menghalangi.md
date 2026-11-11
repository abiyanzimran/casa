# ADR-009 — Aturan permintaan menandai, tidak menghalangi pencatatan

| | |
|---|---|
| Status | Diusulkan; menunggu persetujuan pemilik produk di checkpoint S2. Direvisi 11 November 2026 setelah keputusan pemilik produk 14, 16, dan 17 tentang jatah pribadi (konsekuensi "Jatah pribadi"), lalu atas tinjauan `critic` S2: label keputusan dibatasi pada isi ketiga keputusan itu |
| Tanggal | 11 November 2026 |
| Pengusul | `architect`; mengukuhkan asumsi A-20 dan menjawab Q-54 |
| Rincian | [06-activity-approval-architecture.md](../06-activity-approval-architecture.md) §3.5, §4.5, §5.4, §6.1, §6.2, §6.3 "Jatah pribadi"; INV-A11, INV-A17, INV-A20, INV-A37, INV-A38; [07-financial-domain.md](../07-financial-domain.md) §3.2, §3.8, INV-26 |

## Konteks

Keputusan pemilik produk nomor 5 (checkpoint S1): pemakaian dana bersama yang besar, tidak rutin, dan tidak direncanakan, di atas satu *ambang permintaan*, wajib diajukan lebih dulu. Tetapi Family OS hanya mencatat uang; ia tidak memegang atau memindahkannya (brief §1), sehingga tidak dapat mencegah uang dibelanjakan. Kriteria sukses nomor 1 menuntut catatan mengikuti kenyataan. Pencatat utama pelupa (brief §1.1). Brief §5D melarang alur persetujuan menjadi alat kendali.

Pertanyaannya: apa yang dilakukan sistem ketika pemakaian yang "wajib diajukan" dicatat tanpa permintaan, atau melampaui batas yang disetujui.

## Opsi

| | A. Tolak pencatatan | B. Tahan entri sampai disetujui | C. Simpan, lalu wajibkan persetujuan susulan | D. Simpan dan tandai |
|---|---|---|---|---|
| Catatan mengikuti kenyataan | Tidak: uangnya sudah keluar | Tidak: saldo tertinggal | Ya | Ya |
| Diam pasangan | — | Menahan catatan | Menggantungkan entri tanpa akhir | Tidak berakibat |
| Dapat dipakai menekan | Ya | Ya | Ya | Tidak |
| Pemakaian tetap terlihat | Tidak tercatat sama sekali | Tertunda | Ya | Ya |

## Keputusan

**Opsi D.** Tidak ada state permintaan, ambang, kesepakatan, atau keberatan yang dapat menyebabkan perintah entri yang sah menurut dokumen 07 ditolak atau ditahan.

- Pemakaian di atas ambang tanpa otorisasi dan tanpa kesepakatan rutin mendapat **tanda netral yang diturunkan server**, terlihat sama oleh kedua pasangan, dan pasangan diberi tahu satu kali.
- Realisasi yang melampaui batas otorisasi tampil sebagai kelebihan; pemohon boleh, tidak wajib, mengusulkan perubahan syarat.
- Tidak ada persetujuan susulan yang wajib. Yang tidak sepakat mengajukan *keberatan*, yang tidak mengubah angka.
- Tidak ada hitungan, peringkat, atau rekap tanda per orang.

## Konsekuensi

- Ambang permintaan adalah kesepakatan yang pelanggarannya terlihat, bukan kunci. Kriteria sukses nomor 4 (jalan-jalan dan pembelian tidak rutin datang sebagai permintaan, bukan kejutan) bergantung pada kebiasaan pasangan, bukan pada paksaan sistem (07 §10.3 butir 4). Ini dinyatakan terbuka kepada pemilik produk.
- Satu pemakaian besar dapat dipecah menjadi beberapa entri di bawah ambang; tidak dideteksi, tetap terlihat.
- Uang masuk yang diarahkan ke dana pribadi saat dicatat tidak pernah dihalangi dan tidak memerlukan persetujuan; pasangan melihat catatannya (keputusan pemilik produk, 11 November 2026; dokumen 06 §6.4). Bagian itu diperlakukan seperti penyisihan lain, sama dengan usulan dokumen 07 (T-13), supaya tanda tidak dapat dihindari dengan mengarahkan uang saat mencatat; semuanya **[Usulan]**: di dalam jatah pribadi tidak ditandai (bawaan D-48 dokumen 17), tanpa jatah yang berlaku dibandingkan dengan ambang (D-18, D-48), dan di atas jatah menurut pertanyaan dokumen 17 S2-1 butir (iii). Jawaban mana pun tidak mengubah keputusan opsi D.
- Tanda tidak boleh menjadi alat pantau (pemeriksaan awal `security-privacy`, SP-09 dan SP-10; **[Usulan]**): ambang selalu di atas nol; pemberitahuan langsung atas tanda paling banyak satu per penerima per hari, selebihnya masuk ringkasan; tautan susulan ke otorisasi hanya oleh pencatat entri atau pemohon otorisasi, dan pencabutannya hanya oleh pembuat tautan atau pencatat entri, sehingga pasangan tidak dapat memasang kembali tanda pada entri orang lain (dokumen 06 §4.5, §6.2, INV-A34). Apakah salah satu pihak boleh menarik diri dari kesepakatan ambang adalah pertanyaan pemilik produk; aturan yang berlaku tetap "hanya berganti bila keduanya setuju" (dokumen 06 §6.3).
- Baris `expense` yang ditanggung dana bersama dibandingkan dengan ambang di entri jenis apa pun (termasuk denda di dalam pelunasan, biaya admin di dalam pemindahan, dan perincian selisih kas), kecuali pada `contribution_expense` (dibayar dana pribadi) dan `advance` (sudah terhitung sebagai hak ganti); kewajiban yang didaftarkan sebagai saldo awal sesudah ambang berlaku ditandai satu kali; utang uang yang diterima (`loan_received`) tidak ditandai, karena bukan pemakaian dan uangnya terlihat masuk (dokumen 06 §6.2; 07 §3.8; **[Usulan]**).
- Kesepakatan rutin atas **tagihan berulang** dibandingkan per entri; kekerapan tidak diperiksa (dokumen 06 §6.2, §6.3; **[Usulan]**).
- **Jatah pribadi** (menggantikan konsekuensi sebelumnya "tidak ada batas kumulatif per periode"). **Yang diputuskan pemilik produk, 11 November 2026 (dokumen 17 keputusan 14, 16, 17), hanya tiga hal:** jatah pribadi adalah satu jumlah per orang per bulan, dijumlahkan sepanjang bulan itu, jumlahnya dapat diatur; jumlahnya disepakati berdua, satu mengusulkan dan yang lain mengonfirmasi; kenaikan dapat diajukan dengan alasan. `architect` semula merekomendasikan tidak menambah aturan kumulatif; keberatannya disampaikan kepada pemilik produk, yang memilih bentuk kumulatif. Keputusan opsi D tidak berubah: jatah tidak pernah menolak atau menahan entri. **Selebihnya [Usulan]** (dokumen 06 §6.3 "Jatah pribadi"), termasuk dua aturan yang paling menentukan:
  - **Di dalam jatah tidak ada tanda dan tidak ada pemberitahuan.** Ini dasar pertanyaan lead yang disampaikan kepada pemilik produk dan tidak ia bantah; ia bukan kata-kata pemilik produk dan dikukuhkan lewat bawaan D-48 di checkpoint S2.
  - **Bagian yang melampaui jatah** dibandingkan dengan ambang per entri, seperti penyisihan lain. Ini **terbuka**: dokumen 17 S2-1 butir (iii), dengan rancangan ini sebagai pilihan (a) dan "ditandai berapa pun begitu jatah terlewati" sebagai pilihan (b).
  - Semua penyisihan kepada anggota itu terhitung dengan sendirinya, tanpa ditautkan; yang tertaut ke otorisasi dan pelepasan klaim tidak terhitung.
  - Periodenya adalah periode anggaran yang memuat **tanggal kejadian** entri, bukan waktu simpan, supaya angka jatah sama dengan ringkasan bulanan dan entri terlambat atau offline tidak berpindah bulan. Jalan menghindar yang diterima: tanggal kejadian dimundurkan ke periode yang jatahnya masih longgar; tanggal pencatatan tampil pada entri bila berbeda dari tanggal kejadian, tanpa jamnya (SP-28).
  - Tanda dinilai satu kali, saat entri disimpan, terhadap jumlah periode yang dikenal server saat itu, dan jatah serta jumlah itu dicatat bersama tanda (catatan penilaian). Entri terlambat, pembalik, tautan susulan, dan perubahan jatah tidak pernah memasang atau menghapus tanda pada entri lain. Akibat yang diterima: entri mana yang ditandai bergantung pada urutan tiba di server, sehingga dua entri yang isinya sama dapat bertanda berbeda.
  - Akibat pilihan (a): jatah ditambah potongan yang tidak melebihi ambang tidak pernah ditandai. Pilihan (b) mengubah satu syarat tanda (dokumen 06 §6.2) dan tidak mengubah mekanisme catatan penilaian.
  - Jatah pribadi hanya dapat diakhiri sendirian oleh pemiliknya; pasangan mengurangi lewat usulan yang disetujui pemilik jatah (temuan `security-privacy` SP-24; dokumen 06 §6.3, PG-14). Catatan penilaian hanya dibaca per entri (SP-31).
  - Penilaian harus berurutan per anggota per periode di dalam transaksi tulis entri, dan urutan penerimaan server per rumah tangga harus tersimpan; ini syarat bagi dokumen 09 di S4.
- Tanda adalah fungsi dari *jumlah pemakaian* entri (07 §3.8; dibaca dari baris, bukan dari jenis entri), tautan yang berlaku atasnya termasuk tautan susulan, ambang, kesepakatan yang berlaku, **dan, untuk penyisihan kepada anggota berjatah, catatan penilaian entri itu**. Ia tetap dapat diuji, kini dengan dua pemeriksaan: tanda dihitung ulang dari entri dan catatan penilaiannya (INV-A17), dan catatan penilaian dihitung ulang dengan memutar ulang entri menurut urutan penerimaan server (INV-A37). Tanda **bukan lagi** fungsi murni dari entrinya sendiri. Pencatat tidak dapat memasang atau melepasnya.
- Pelunasan kewajiban yang sudah tercatat (pembayaran cicilan dan utang) tidak ditandai; pembelian secara cicilan dibandingkan dengan ambang satu kali saat dibeli, sebesar uang muka ditambah jumlah seluruh cicilan (07 §3.8, CTH-34). Cicilan karenanya tidak memerlukan kesepakatan rutin (dokumen 06 §6.2, §6.3).
- Tautan yang tidak sah (ke permintaan yang dibatalkan, atau bertanggal di luar masa hidup otorisasinya; 07 INV-26) tidak membuat entrinya ditolak: entri tersimpan tanpa tautan.
- Rumah tangga dengan satu anggota dewasa tidak punya ambang yang berlaku, jadi tidak ada tanda; saat anggota dewasa kedua bergabung dan menyetujui ambang, tidak ada entri lama yang ditandai surut (dokumen 06 §3.5).
- Pencatatan offline selalu diterima; tanda dihitung server saat entri tiba.
- Default Q-13 "kelebihan memerlukan persetujuan lanjutan" dilunakkan menjadi "boleh diajukan".
- `security-privacy` diminta menilai apakah tanda itu sendiri dapat menjadi alat pantau (dokumen 06 §15.5).
- Ditinjau ulang hanya bila produk kelak memegang atau memindahkan uang (brief §1, di luar MVP); pada saat itu penolakan menjadi mungkin secara teknis dan perlu keputusan baru.
