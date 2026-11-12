# 01 — Ringkasan Eksekutif

| | |
|---|---|
| Pemilik | Lead |
| Tahap | S5 — Rencana dan sintesis, ditulis paling akhir (brief §0.3), 12 November 2026; S6 (12 November 2026): pass konsistensi |
| Status | Ditinjau `critic` (S5 B, lalu tinjauan ulang B2, 12 November 2026); direvisi untuk keputusan 42 dan ditinjau `critic` lagi (12 November 2026); temuan diperbaiki; direvisi untuk keputusan 43–50 dan 51–55 (12 November 2026); S6: keputusan 56–68; sesudah S6 (13 November 2026): keputusan 69 (klien web lebih dulu) dan 70 (nama Casa) |
| Merangkum | Brief §1.1; [02](02-product-principles.md), [04](04-feature-map.md), [09](09-technical-architecture.md), [13](13-product-roadmap.md), [16](16-final-recommendations.md), [18 §1.4](18-market-research.md#14-peluang-pembedaan-positioning-dan-value-proposition). Tidak menetapkan apa pun yang baru |

---

**Nama.** *Casa*, nama sementara selama pemakaian keluarga dan belum go public (keputusan 70); *Family OS* adalah nama kerja proyek perencanaan.

**Visi.** Aplikasi tempat suami dan istri mengetahui keadaan uang dan urusan rumah tangga tanpa saling bertanya, dengan usaha mencatat serendah mungkin, dan yang membantu keduanya saling memperhatikan. Dipakai setiap hari oleh rumah tangga pemilik produk, dirancang dan dibangun dengan standar produk publik untuk keluarga Indonesia. Go public tidak direncanakan, tetapi tetap terbuka tanpa pembangunan ulang (brief §1.1; keputusan 51).

**Masalah inti.** Uang rumah tangga tidak terlihat dan tidak direncanakan, dan mengoordinasikannya berarti terus bertanya. Sesudah gaji dibagi ke cicilan, tagihan, dan kiriman untuk orang tua, lalu ditarik tunai dan berpindah tangan, tidak ada yang tahu siapa memegang berapa. Jalan-jalan diumumkan mendadak, anggaran dapur terlampaui, dan pencatat utama pelupa, jadi mencatat harus selesai dalam hitungan detik dan boleh terlambat (brief §1.1).

**Audiens.** Pertama satu rumah tangga: suami dan istri sebagai dua dewasa setara; pilot lewat klien web di iPhone istri, iPad dan Galaxy A55 suami, dan browser desktop; aplikasi native iOS dan Android menyusul (keputusan 69). Orang tua dan kerabat dewasa yang kelak diberi login dihitung pengguna luar, dengan syarat keamanannya lebih dulu (keputusan 52, 63). Bila kelak go public, segmen publik pertama yang diusulkan adalah rumah tangga dua dewasa dan satu dewasa **[Usulan]** (17 Q-214). Struktur keluarga lain sudah dirancang; kapan dibangun bergantung pada pemicunya (13 §3.8).

**Value proposition** (18 §1.4, **[Usulan]**).

> Family OS membuat suami dan istri sama-sama tahu berapa uang rumah tangga yang tersisa dan siapa yang memegangnya, serta menyepakati pengeluaran sebelum terjadi, tanpa harus saling bertanya dan dengan usaha mencatat serendah mungkin, sambil membantu keduanya mengurus rumah dan saling memperhatikan.

**Pembeda** (**[Usulan]**; belum diuji langsung terhadap aplikasi pesaing, 18 §6.4): custody sebagai konsep utama ("sisa berapa dan siapa memegang"); permintaan dua arah yang selesai sampai uangnya; arus uang keluarga Indonesia apa adanya (kiriman ke orang tua, pinjaman ke saudara, uang tunai, serah terima); uang dan perhatian di satu tempat, dengan pagar: dua dewasa setara, pribadi secara bawaan, tidak ada yang mengawasi atau bertindak atas nama orang lain. Pencatatan cepat bukan pembeda; Family OS harus setara cepatnya.

**MVP yang direkomendasikan** (16 §2).

- **Tier 1, sebelum data nyata (55 fitur):** custody dan angka utama, uang masuk dan pembagiannya, pengeluaran cepat, serah terima, jatah pribadi, anggaran bulanan, tagihan dan pos rutin, cicilan, utang dan piutang termasuk keluarga besar, tabungan sederhana, aset, ringkasan bulanan; permintaan dan persetujuan dua arah; notifikasi dengan jam tenang; chat rumah tangga ringan; kunci aplikasi dan halaman Bantuan.
- **Tier 2, bertahap di dalam pilot (19 fitur):** suasana hati dan tindak lanjut sakit, permintaan tindakan, belanja, masalah rumah, rutinitas, rencana, pendamping `AI` versi ringan, faktor kedua dan kode pemulihan opsional per akun.
- **Tidak di pilot:** sinkron bank, dan apa pun yang memegang atau memindahkan uang sungguhan.

**Cara membangunnya** (09, 14, ADR-030). Klien web (PWA) lebih dulu: responsif, dapat dipakai luring, satu UI untuk semua perangkat, tanpa Mac; tata letak, arah visual, dan kata-katanya sama dengan aplikasi native, tetapi belum terasa sepenuhnya iPhone atau Android (keputusan 69). Aplikasi native penuh per platform, Swift dan Kotlin (keputusan pemilik produk), menyusul sesudah data nyata: iOS dan iPadOS lebih dulu, lalu Android. Satu server Go dengan PostgreSQL di Indonesia memegang semua aturan uang dan persetujuan. Biaya pilot per bulan sekitar Rp260.667 + PPN bila berlaku, ditambah USD 1,32 selama jalur web (paling tinggi USD 10) dan email; USD 9,57 (paling tinggi USD 18,25) sesudah jalur iOS native dimulai; sebagian besar angka rupiah bergantung pada pilihan hosting yang ditunda (09.4 §7.1).

**Jalan ke depan** (13). Tier 1 di klien web dengan data uji → gerbang data nyata → pemakaian sehari-hari dengan Tier 2 bergelombang, lalu aplikasi iOS native dan aplikasi Android menyusul → tinjauan tiga bulan, bukan gerbang → pemakaian dan pemeliharaan. Tanpa tanggal tetap. Dua jalur opsional, tidak dijadwalkan: sebelum orang tua diberi login, gerbang pengguna luar pertama (butir keamanan dan pemberitahuan pengelola server versi keluarga; menurut advokat tanpa PSE selama tertutup untuk keluarga; keputusan 52, 59); bila kelak go public, GR-3 dengan syarat keputusan 46, urusan hukum dan dokumen privasi (keputusan 67, 68), beta tertutup, lalu rilis publik freemium **[Usulan]**. Keputusan berikutnya pemilik produk ada di 16 §5.
