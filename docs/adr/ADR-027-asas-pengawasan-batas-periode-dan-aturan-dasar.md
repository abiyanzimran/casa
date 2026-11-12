# ADR-027 — Asas pengawasan: batas periode dan aturan dasar

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: **opsi A** (D-113, diterima dengan "lanjut"; [_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk"). Butir (b) dan (c) dirumuskan sesuai S4-9 (i) A: ambang dibuat, diubah, dan diakhiri hanya dengan persetujuan semua pihak, untuk semua rumah tangga. Kata-katanya: "Ambang itu batas yang di tentukan/disepakati 2 orang" dan "ambang di akhiri berdua juga" (jawaban kedua dan keempat). S4-9 (ii) gugur, karena hanya relevan bila (i) B atau C |
| Tanggal | 12 November 2026 (S4 fase D; status checkpoint S4) |
| Pengusul | `security-privacy` (dokumen 10.3 §2.4.1; SP-33), ditulis `architect` |
| Mengubah | Bunyi 06 PG-14, kalimat kedua dan ketiga serta "Dua titik yang tidak selaras". Penerapannya di 06 tidak berubah |
| Rincian | [06 §10 PG-14](../06-activity-approval-architecture.md); [10.3 §2.4.1](../10-security-privacy/03-anak-penyalahgunaan-siklus.md) |

## Konteks

06 PG-14 merumuskan asas pengawasan: seorang anggota dewasa boleh sendirian mengurangi pengawasan atas siapa pun atau menambah pengawasan atas dirinya sendiri; menambah pengawasan atas orang dewasa lain memerlukan persetujuannya. Dua aturan yang sudah berlaku tidak sesuai dengan bunyi itu (K-34 titik (2) dan (3)):

- **Titik (2).** Jatah di atas ambang yang diakhiri pemiliknya tetap menjadi dasar penilaian sampai akhir periode anggaran itu. Ini rumusan lead yang dikukuhkan pemilik produk dengan "bener" (`_progress.md`, 11 November 2026; 07 T-29 pilihan 1), dipersempit SP-33 pada jatah yang lebih besar daripada ambang saat K06. Menurut bunyi PG-14, melepas pengawasan atas diri sendiri seharusnya berlaku seketika.
- **Titik (3).** Satu orang boleh mengakhiri kesepakatan pos rutin sendirian, seketika, dan pasangannya diberi tahu (S3-5 (i), diterima lewat "lanjut" di checkpoint S3; 06 §6.3 "Pos rutin" (d), (f); P-26; W-25). Akibatnya pembayaran pos itu oleh orang dewasa lain dapat menjadi lebih mungkin ditandai tanpa persetujuannya.

`security-privacy` menilai keduanya bertentangan dengan bunyi PG-14 tetapi terbatas, dan menyarankan perilakunya dipertahankan sementara asasnya dirumuskan ulang (10.3 §2.4.1).

## Opsi

| Opsi | Akibat |
|---|---|
| **A. Rumuskan ulang PG-14 dengan butir (a) sampai (e)** | Satu asas menjelaskan semua aturan yang berlaku; tidak ada pengecualian yang harus diingat |
| B. Pertahankan bunyi PG-14; catat titik (2) dan (3) sebagai dua pengecualian terbatas | Perilaku sama; asasnya tetap dengan dua catatan kaki |
| C. Ubah perilaku agar sesuai bunyi PG-14 sekarang | Titik (2): K06 menjadi jalan keluar sesudah kelonggaran jatah terpakai. Titik (3): orang yang tidak lagi setuju dengan jumlah per kali tidak dapat menarik persetujuannya. Keduanya mengubah keputusan yang sudah diterima pemilik produk |

## Keputusan

**Opsi A** (diterima). Bunyi baru PG-14:

- **(a)** Menambah pengawasan atas orang dewasa lain hanya berlaku dengan persetujuan orang itu (tidak berubah).
- **(b)** Melepas sesuatu yang hanya pernah menambah pengawasan atas diri sendiri boleh sendirian dan berlaku seketika: jatah yang tidak di atas ambang, termasuk 0; jatah selama tidak ada ambang yang berlaku (06 K06; 07 §3.4); berbagi data sendiri (06 V-7).
- **(c) Aturan batas periode.** Mengakhiri kesepakatan timbal balik yang dalam periode berjalan memberi kelonggaran sekaligus batas atas diri sendiri boleh sendirian. Pengakhirannya seketika, terlihat, dan diberitahukan, tetapi akibatnya atas penilaian berlaku pada batas periode, tidak di tengah periode dan tidak surut. **Sejak S4-9 (i) A butir ini hanya berlaku atas satu hal: jatah di atas ambang yang diakhiri pemiliknya (06 K06, `assessment_until`; titik (2)).** Kesepakatan ambang tidak termasuk; ia diatur butir (e): ambang tidak dapat diakhiri sendirian, dan diakhiri lewat usulan `rule` tanpa ambang yang disetujui semua pihak yang disyaratkan (06 §6.3 K01, K02, K05). Pengakhiran bersama itu bukan perintah sepihak, jadi aturan batas periode tidak berlaku atasnya; ia berlaku seketika atas entri yang disimpan sesudah K02 yang melengkapi, ke depan saja (07 §3.8, **[Usulan]** `finance-domain`). K07 dan `threshold_assessment_until` tidak dibangun dan tidak dirancang untuk tahap publik.
- **(d) Aturan kembali ke dasar yang disepakati.** Perintah sepihak boleh membuat pembayaran bersama yang dilakukan orang dewasa lain lebih mungkin ditandai hanya bila hasilnya kembali ke aturan dasar yang orang itu sudah setujui (ambang yang berlaku; tanpa ambang berarti tanpa tanda), ke depan saja, dan tidak pernah atas uang pribadi seseorang (titik (3)).
- **(e)** Asas ini batas bawah, bukan izin: keputusan pemilik produk yang menuntut semua pihak tetap berlaku (keputusan 16; ambang, S4-9 (i) A).

## Konsekuensi

- Perilaku produk tidak berubah oleh rumusan ini; tidak ada contoh terhitung yang berubah.
- K-34 titik (2) dan (3) ditutup. Titik (1) ditutup oleh S4-9 (i) A lewat butir (e): menambah atau mengurangi pengawasan lewat ambang selalu memerlukan semua pihak.
- Daftar uji PG-14 di 06 mendapat dua kasus baru, satu per aturan (06 §10 PG-14), ditambah satu kasus: `end` atau perintah sepihak apa pun atas `rule` ditolak, dan usulan `rule` tanpa ambang hanya berlaku sesudah semua pihak menyetujuinya.
- Opsi B (dua pengecualian terbatas) tidak dipilih.
