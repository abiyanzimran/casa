# ADR-021 — Jalur `AI`: kumpulan kalimat (R1) dan batas biaya

| | |
|---|---|
| Status | **Diusulkan** (`architect`, 12 November 2026); akan diputuskan pemilik produk di checkpoint S4: 10-B (jalur; bawaan dan saran A = R1) dan 09-D (penyedia) |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D) |
| Pengusul | `architect`, atas R1 dari `security-privacy` (dokumen 10 §10.4 §1.2; 10-B) |
| Mengubah | — |
| Rincian | [09.1 §11](../09-technical-architecture/01-pilihan-teknologi.md); [09.3 §11](../09-technical-architecture/03-sinkron-waktu-nyata-notifikasi.md) |

## Konteks

Brief §9.5 meminta perbandingan antara AI hosted dan self-hosted atas mutu bahasa Indonesia, privasi, biaya pada anggaran pilot, dan beban operasi, lalu satu rekomendasi dengan batas biaya dan fallback. Semua panggilan harus lewat server.

Pemilik produk menginginkan Claude atau ChatGPT, dan server di Indonesia (keputusan 10, dokumen 17). Riset menemukan:

- Claude dan ChatGPT tidak dapat diproses di Indonesia lewat jalur mana pun yang ditemukan (18 §5.9);
- sesudah PP 33/2026 berlaku, peneliti tidak menemukan jalur bagi pemanggilan harian (18 §2.8; **[Perlu validasi]** advokat).

R1 dari dokumen 10 memenuhi kedua keinginan itu: `AI` menulis kumpulan kalimat tanpa data rumah tangga, dan server di Indonesia memilih kalimat bagi tiap orang. `AI` adalah Tier 2.

## Opsi

| Kriteria | **Hosted Claude, R1** | Hosted ChatGPT, R1 | Self-hosted SEA-LION v3 9B |
|---|---|---|---|
| Mutu bahasa Indonesia | Klaim vendor 97,1% sampai 98,6% Global-MMLU (18 §5.9) | Tidak ada tolok ukur primer | Belum terbukti |
| Data rumah tangga ke luar negeri | Tidak ada | Tidak ada | Tidak ada |
| Biaya per pembangunan (0,06 juta token masuk, 0,12 juta keluar **[Asumsi]**) | Sonnet 5.5: 0,06 × 2 + 0,12 × 10 = **USD 1,32** | gpt-6.1-terra: 0,06 × 1 + 0,12 × 5 = **USD 0,66**, separuhnya | VPS 8 sampai 16 GB, atau sewa jam-jaman saat membangun kumpulan |
| Beban operasi | Kunci API, batas belanja | Sama | Model, runtime, mutu |
| Batas belanja di penyedia | Spend limit per workspace, tidak di Default Workspace | Tidak diteliti | — |

## Keputusan

1. **Jalur R1**, bila 10-B = A.
   - Pekerjaan `ai.pool_build` (bulanan, atau saat templat berubah) mengirim templat produk tanpa nama, judul, jumlah, tanggal, status, atau foto (SEC-120, SEC-121, SEC-123, SEC-69).
   - Hasilnya disaring otomatis dengan daftar kalimat terlarang, batas panjang, dan bahasa (SEC-130), lalu disimpan sebagai `sentence_pool` berversi.
2. **Saat dipakai tidak ada pemanggilan `AI`.**
   - Pemilih di server memakai kunci kategori dari jalur baca dan rotasi tanpa pengulangan, dan tidak pernah membandingkan orang (SEC-129).
   - "Minta ubah" berarti kalimat lain dari kumpulan.
3. **Penyedia: Claude lewat API Anthropic** (09-D pilihan A), dengan model Sonnet 5.5 untuk membangun kumpulan. Pihak kontraknya Anthropic, PBC (18 §5.9). Alasannya mutu bahasa Indonesia yang punya klaim terukur; ChatGPT separuh biayanya per pembangunan, tetapi selisihnya hanya USD 0,66 per bulan pada satu pembangunan bulanan. Bila 09-D = B, penyedianya OpenAI dengan gpt-6.1-terra, dan pihak kontraknya belum diteliti.
4. **Batas biaya di dua tempat** (SEC-127):
   - spend limit bulanan USD 10 pada workspace khusus (bukan Default Workspace);
   - penghitung `ai_usage` di server dengan batas **per bulan USD 10** dan **per pekerjaan `ai.pool_build` USD 3** **[Asumsi]**. Batas **per hari** (USD 0,33) hanya untuk jalur R2 `ai.suggest` (10-B B), tidak untuk pembangunan kumpulan;
   - sebelum setiap panggilan, biaya terburuknya dihitung (token masuk × harga masuk + `max_tokens` × harga keluar), dan panggilan tidak dilakukan bila sisa batas mana pun lebih kecil;
   - pembangunan yang berhenti karena batas **tidak menerbitkan kumpulan setengah jadi**: kumpulan baru diaktifkan hanya bila selesai dan lolos saringan, dan kumpulan lama tetap berlaku utuh (09.3 §11.1).
5. **Saklar mati** (SEC-128) adalah setelan server yang dibaca setiap panggilan dan setiap pemilihan.
6. **Fallback:** bila penyedia gagal atau batas tercapai, kumpulan lama tetap dipakai. Bila kumpulan kosong, tidak ada saran. Pencatatan, permintaan, dan persetujuan tidak pernah bergantung pada modul `ai` (LM-02; brief §5O).
7. **Pilihan 10-B lain:**
   - **B:** menambah jalur `ai.suggest` berdata minimum (Haiku 5.5) dengan persetujuan tercatat, kunci R2 (SEC-122), dan saklar tanggal PP yang mematikannya otomatis.
   - **C:** model self-hosted dijalankan hanya saat membangun kumpulan, di VPS sewa jam-jaman atau di mesin di Indonesia, tidak di VPS pilot.
   - **D:** modul `ai` mati.

## Konsekuensi

- Tidak ada data rumah tangga yang keluar ke penyedia `AI`. Pertanyaan transfer lintas batas untuk `AI` gugur di pilot selama R1 dipakai.
- Saran bersifat umum menurut kategori. Brief §5O "asks the AI to revise it" dipenuhi sebagai "saran lain dari kumpulan", dan selisihnya dicatat lead (dokumen 10 §8 butir 5).
- Biaya `AI` tidak tumbuh dengan jumlah rumah tangga, karena satu kumpulan dipakai semua.
- Mutu kalimat perlu diuji buta dengan pasangan pilot sebelum memilih model (18 §5.9 butir 4). Templat dan kumpulan dapat diganti tanpa rilis aplikasi.
