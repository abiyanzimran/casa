### FR-CAL-002 — Pengingat tagihan

| | |
|---|---|
| Fitur | F-CAL-002 (Tier 1 di 04, S3-1 (d)) |
| Tier / MoSCoW | 1 / **Must**. Ditaruh di berkas ini hanya sebagai penunjuk; kriterianya diuji di FR-FIN-100/AC-1 dan FR-NTF-009/AC-5 |
| Masalah (brief §1.1) | `MSL-6` ("who pays which bill, and how much, must be clear"); `KB-06` |
| Aturan acuan | 06 §11.2 N-32, §13.4 (termasuk "belum dibayar" dibaca per kejadian); tagihan berulang dan penanggung jawabnya: 07 §3.15, PA-FIN-5 (F-FIN-011 di [FIN-2.md](FIN-2.md)); tanda atas pembayaran: 06 §6.2 |
| Kedudukan | **Dinaikkan ke Tier 1 sebagai N-32** (`ringkasan`; usulan `architect`, rekomendasi atas CA-08 dan PA-SCN-2; 06 P-36), **diterima lewat "lanjut"** di checkpoint S3 (12 November 2026), S3-1 (d). Angka tiga hari **[Asumsi]**. Bawaan semula berkas ini (tiga hari sebelum dan pada hari jatuh tempo, kepada penanggung jawab atau keduanya) diambil 06 N-32 |
| Usaha berulang | Nol di luar mencatat pembayarannya |

**Kebutuhan ini kini penunjuk.** FR-FIN-100/AC-1 (Tier 1, tagihan rutin) bergantung pada pengingat tagihan, dan pengingat itu tidak dapat menunggu Tier 2 (CA-08). 06 menjadikannya N-32 Tier 1: satu baris `ringkasan` tanpa jumlah tiga hari sebelum jatuh tempo **[Asumsi]** dan satu pada hari jatuh tempo, per kejadian, kepada penanggung jawab atau keduanya bila tidak ada; dibuang dari ringkasan bila kejadian itu sudah dibayar saat ringkasan dilepas; berhenti sesudah tanggal akhir atau saat tagihan diarsipkan; tidak bergantung pada state kesepakatan pos rutinnya. Kriteria penerimaannya kini di Tier 1: [FIN-2.md](FIN-2.md) FR-FIN-100/AC-1 dan [NTF.md](NTF.md) FR-NTF-001 (tabel prioritas) dan FR-NTF-009/AC-5. Ketiga AC lama di bawah dipertahankan sebagai uji N-32, bukan sebagai kebutuhan Tier 2 tersendiri. Dokumen 04 menempatkan F-CAL-002 di Tier 1 (S3-1 (d), diterima lewat "lanjut"); karena itu FR ini berlabel Must Tier 1 dan masuk tabel keterlacakan indeks §6.7. Cabang lama "bila ditolak, kembali Tier 2 Should dan FR-FIN-100/AC-1 dipersempit" tidak berlaku lagi.

**Kriteria penerimaan** (uji N-32)

- **AC-1** — *Given* tagihan berulang dengan penanggung jawab, *When* tanggal pengingat tiba, *Then* penanggung jawab menerima N-32 lewat ringkasan, tanpa jumlah dan tanpa layar kunci tersendiri; pengingat berjalan apa pun state kesepakatan pos rutinnya (06 N-32).
- **AC-2** — *Given* pengingat tagihan, *When* penanggung jawab mencatat pembayaran yang ditautkan ke kejadian tagihan itu, *Then* kejadian itu tampil sudah dibayar dan pengingatnya berhenti; pembayaran lebih awal untuk kejadian berikutnya tidak menghentikan pengingat kejadian yang masih terbuka; tidak ada tombol "sudah dibayar" tanpa entri (06 §13.4; 07 §3.15, PA-FIN-5).
- **AC-3** — *Given* tagihan yang lewat jatuh tempo tanpa pembayaran, *When* hari berganti, *Then* ia tetap tampil sebagai kejadian belum dibayar, tanpa tanda pada orang dan tanpa notifikasi tambahan selain dua baris N-32.

