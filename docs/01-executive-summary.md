# 01 — Ringkasan Eksekutif

| | |
|---|---|
| Pemilik | Lead |
| Tahap | S5 — Rencana dan sintesis, ditulis paling akhir (brief §0.3), 12 November 2026 |
| Status | Ditinjau `critic` (S5 B, lalu tinjauan ulang B2, 12 November 2026); direvisi untuk keputusan 42 dan ditinjau `critic` lagi (12 November 2026); temuan diperbaiki; direvisi untuk keputusan 43–50 (12 November 2026) |
| Merangkum | Brief §1.1; [02](02-product-principles.md), [04](04-feature-map.md), [09](09-technical-architecture.md), [13](13-product-roadmap.md), [16](16-final-recommendations.md), [18 §1.4](18-market-research.md#14-peluang-pembedaan-positioning-dan-value-proposition). Tidak menetapkan apa pun yang baru |

---

**Visi.** Aplikasi mobile tempat suami dan istri mengetahui keadaan uang dan urusan rumah tangga tanpa saling bertanya, dengan usaha mencatat serendah mungkin, dan yang membantu keduanya saling memperhatikan. Bertahap: lebih dulu dipakai setiap hari oleh rumah tangga pemilik produk, lalu menjadi produk publik untuk keluarga Indonesia (brief §1.1).

**Masalah inti.** Uang rumah tangga tidak terlihat dan tidak direncanakan, dan mengoordinasikannya berarti terus bertanya. Sesudah gaji dibagi ke cicilan, tagihan, dan kiriman untuk orang tua, lalu ditarik tunai dan berpindah tangan, tidak ada yang tahu siapa memegang berapa. Jalan-jalan diumumkan mendadak, anggaran dapur terlampaui, dan pencatat utama pelupa, jadi mencatat harus selesai dalam hitungan detik dan boleh terlambat (brief §1.1).

**Audiens.** Pertama satu rumah tangga: suami dan istri sebagai dua dewasa setara; pilot di iPhone istri dan iPad suami, Android menyusul (keputusan 42). Lalu rumah tangga Indonesia; segmen publik pertama yang diusulkan adalah rumah tangga dua dewasa dan satu dewasa **[Usulan]** (17 Q-214). Struktur keluarga lain sudah dirancang; kapan dibangun bergantung pada pemicunya (13 §3.8).

**Value proposition** (18 §1.4, **[Usulan]**).

> Family OS membuat suami dan istri sama-sama tahu berapa uang rumah tangga yang tersisa dan siapa yang memegangnya, serta menyepakati pengeluaran sebelum terjadi, tanpa harus saling bertanya dan dengan usaha mencatat serendah mungkin, sambil membantu keduanya mengurus rumah dan saling memperhatikan.

**Pembeda** (**[Usulan]**; belum diuji langsung terhadap aplikasi pesaing, 18 §6.4): custody sebagai konsep utama ("sisa berapa dan siapa memegang"); permintaan dua arah yang selesai sampai uangnya; arus uang keluarga Indonesia apa adanya (kiriman ke orang tua, pinjaman ke saudara, uang tunai, serah terima); uang dan perhatian di satu tempat, dengan pagar: dua dewasa setara, pribadi secara bawaan, tidak ada yang mengawasi atau bertindak atas nama orang lain. Pencatatan cepat bukan pembeda; Family OS harus setara cepatnya.

**MVP yang direkomendasikan** (16 §2).

- **Tier 1, sebelum data nyata (55 fitur):** custody dan angka utama, uang masuk dan pembagiannya, pengeluaran cepat, serah terima, jatah pribadi, anggaran bulanan, tagihan dan pos rutin, cicilan, utang dan piutang termasuk keluarga besar, tabungan sederhana, aset, ringkasan bulanan; permintaan dan persetujuan dua arah; notifikasi dengan jam tenang; chat rumah tangga ringan; kunci aplikasi dan halaman Bantuan.
- **Tier 2, bertahap di dalam pilot (18 fitur):** suasana hati dan tindak lanjut sakit, permintaan tindakan, belanja, masalah rumah, rutinitas, rencana, pendamping `AI` versi ringan.
- **Tidak di pilot:** web atau PC, sinkron bank, dan apa pun yang memegang atau memindahkan uang sungguhan.

**Cara membangunnya** (09, 14). Native penuh per platform, Swift dan Kotlin (keputusan pemilik produk), iOS dan iPadOS lebih dulu; satu server Go dengan PostgreSQL di Indonesia memegang semua aturan uang dan persetujuan. Biaya pilot per bulan sekitar Rp260.667 + PPN bila berlaku, ditambah USD 9,57 (paling tinggi USD 18,25) dan email; sebagian besar angka rupiah bergantung pada pilihan hosting yang ditunda (09.4 §7.1).

**Jalan ke depan** (13). Tier 1 di iOS dan iPadOS dengan data uji → gerbang data nyata → pilot tiga bulan dengan Tier 2 bergelombang dan aplikasi Android menyusul → evaluasi → bila lanjut: advokat, PSE, beta tertutup → rilis publik, freemium per rumah tangga dengan satu langganan rumah tangga **[Usulan]**. Tanpa tanggal tetap. Evaluasi paling cepat tiga bulan sesudah data nyata, bila kedua pasangan merasa nyaman dan aman dan 14 hari terakhir tanpa bug baru yang berat (keputusan 46). Keputusan berikutnya pemilik produk ada di 16 §5.
