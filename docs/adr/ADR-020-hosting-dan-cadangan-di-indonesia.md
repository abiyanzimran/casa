# ADR-020 — Hosting dan cadangan di Indonesia

| | |
|---|---|
| Status | **Diusulkan** (`architect`, 12 November 2026). **Ditunda pemilik produk sampai menjelang deploy** (S4-2 (i); kata-katanya: "#S4-2: nanti aja pas mau deploy"; [_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk", jawaban pertama). Syarat server dan cadangan di Indonesia (keputusan 10) tetap mengikat pilihan apa pun. Penyedia, harga, dan pihak kontrak di bawah diperiksa ulang saat S4-2 diajukan lagi. **Arti "menjelang deploy" (keputusan 45, 12 November 2026; jawabannya "Sebelum data nyata (Saran)"):** sebelum deploy pertama yang akan memegang data nyata (GR-2); sampai itu FS-1 memakai server uji sementara berisi data uji saja, yang dibuang (14 §7.8). Karena cadangan, uji pulih, dan staging disiapkan di VPS ini pada FS-1 langkah 7 (QG-30 butir 7 sampai 9), S4-2 diajukan ulang sebelum langkah 7 dimulai **[Usulan]** |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4; keputusan 45) |
| Pengusul | `architect` |
| Mengubah | — |
| Rincian | [09.1 §10](../09-technical-architecture/01-pilihan-teknologi.md); [09.4 §1, §6, §7](../09-technical-architecture/04-operasi-uji-biaya.md) |

## Konteks

Pemilik produk menginginkan server aplikasi di-host di Indonesia (keputusan 10, dokumen 17: "Ia menginginkan AI-nya Claude atau ChatGPT, dan server aplikasi di-host di Indonesia"). Brief §1 membatasi pilot pada satu VPS sederhana tanpa layanan terkelola berharga per pemakaian.

Dokumen 10 mengajukan dua syarat:

- SEC-43: cadangan dienkripsi sebelum meninggalkan VPS, disimpan di Indonesia pada pihak kontrak di Indonesia, dengan jendela pemulihan 30 hari, dan pemulihannya diuji sebelum pilot;
- SEC-96: data yang dihapus bertahan di salinan mana pun paling lama 38 hari (jendela 30 hari + 7 hari interval full + 1 hari jeda salinan);
- ADR hosting mencatat badan hukum pihak kontrak dan lokasi pusat data, untuk VPS dan tujuan cadangan (10 §5).

## Opsi

| Kriteria | **Biznet Gio NEO Lite MS 4.2** | Biznet Gio NEO Lite Pro | Akamai Jakarta | AWS Lightsail Jakarta |
|---|---|---|---|---|
| Spesifikasi, harga | 2 vCPU, 4 GB, 60 GB SSD, Rp139.000 | 2 vCPU, 4 GB, NVMe, Rp559.000 | 4 GB USD 28,80 | 4 GB USD 24 |
| Pihak kontrak | **PT Biznet Gio Nusantara** **[Terverifikasi]** | Sama | Tidak diteliti | PT Amazon Web Services Indonesia **[Terverifikasi]** |
| Pusat data | NEO Cloud di Cimanggis dan MidPlaza (halaman Biznet Gio, 2024; 18 §5 fase D) **[Terverifikasi]**; yang melayani NEO Lite **[Perlu validasi]** | Sama | Jakarta | Jakarta |
| Data pribadi di luar Indonesia | Kebijakan privasi: data pribadi "dapat diproses atau disimpan di luar wilayah Republik Indonesia"; cakupan terbaca: data akun pelanggan (18 §5 fase D) **[Terverifikasi]**; apakah mencakup isi VPS **[Perlu validasi]** advokat | Sama | Tidak diteliti | Tidak diteliti |
| Tagihan per jam | Belum dikukuhkan untuk NEO Lite (18 §5 fase D) | — | — | — |

| Kriteria tujuan cadangan | **VPS IDCloudHost (Indonesia)** | VPS Biznet Gio XS kedua | Object storage per GB |
|---|---|---|---|
| Pihak kontrak | **PT Cloud Hosting Indonesia** **[Perlu validasi]** | PT Biznet Gio Nusantara (sama dengan utama) | Bergantung penyedia |
| Lokasi | Pilih lokasi Indonesia saat memesan (misalnya Jakarta atau Bogor) **[Perlu validasi]** | Indonesia | — |
| Terpisah dari kegagalan akun penyedia utama | **Ya** | Tidak | Bergantung |
| Harga | Sekitar Rp105.000 untuk 4 GB **[Perlu validasi]**; paket lebih kecil mungkin cukup | Rp59.000 | Per pemakaian (BT-5) |

## Keputusan

1. **VPS utama: Biznet Gio NEO Lite MS 4.2.**

   | | |
   |---|---|
   | Pihak kontrak | **PT Biznet Gio Nusantara** |
   | Pusat data | Di Indonesia: NEO Cloud di Cimanggis dan MidPlaza menurut halaman Biznet Gio sendiri. Kota pusat data VPS ini diminta tertulis dari penyedia sebelum data nyata masuk |
   | Kebijakan privasi | Penyedia diminta menyatakan tertulis bahwa kalimat "di luar wilayah Republik Indonesia" tidak mencakup isi VPS pelanggan, sebelum data nyata masuk. Penilaian advokat atasnya (L-01, L-14) dilakukan sebelum tahap publik, karena pemilik produk memilih advokat sebelum tahap publik (S4-4 (iii) A) |
   | Spesifikasi | 2 vCPU, 4 GB RAM, 60 GB SSD |
   | Harga | Rp139.000 per bulan (PPN tidak disebut) |

2. **Tujuan cadangan: VPS di IDCloudHost, di lokasi Indonesia.**

   | | |
   |---|---|
   | Pihak kontrak | **PT Cloud Hosting Indonesia** (dikukuhkan dari syarat layanan saat memesan) |
   | Pusat data | Lokasi Indonesia yang dipilih saat memesan, dicatat di runbook. Lokasi Singapura tidak dipakai |
   | Harga | Paket terkecil yang memuat repositori dan salinan pelindungnya |

   Bila pemilik produk memilih 09-B pilihan B, tujuan cadangannya VPS Biznet Gio NEO Lite XS 1.1 (Rp59.000), dengan pihak kontrak dan risiko akun yang sama dengan VPS utama.
3. **Cadangan:**
   - pgBackRest 2.60.0, berupa base backup fisik dan WAL dengan `archive_timeout = 60 s`;
   - dikirim lewat SFTP ke VPS cadangan;
   - dienkripsi `aes-256-cbc` di VPS utama, dengan kata sandi yang tidak pernah ada di VPS cadangan;
   - jangka simpan berbasis waktu: jendela pemulihan 30 hari (`repo1-retention-full-type=time`, `repo1-retention-full=30`);
   - **salinan pelindung yang hanya menambah** di VPS cadangan: setiap hari berkas baru repositori disalin (`rsync --ignore-existing`, tanpa `--delete`) ke direktori yang tidak dapat ditulis VPS utama; salinan dipangkas menurut umur set cadangan di manifest pgBackRest, bukan umur salinan;
   - akibatnya data yang dihapus bertahan paling lama **38 hari** di salinan mana pun (SEC-96), diperiksa `ops.backup_verify`;
   - jurnal penghapusan dikirim bersama WAL dan dijalankan ulang saat pemulihan (SEC-97).
4. **Uji pemulihan** sebelum pilot memakai data nyata, lalu bulanan: pulihkan ke instans sementara di VPS cadangan (data tetap di Indonesia), jalankan pemeriksaan invarian dan jurnal penghapusan, lalu hapus instans itu.
5. **Sasaran:** RPO ≤ 1 menit, RTO ≤ 4 jam **[Usulan]**.
6. **Pengukuran sebelum menambah ukuran.** Ukuran VPS dikukuhkan setelah pengukuran p95 (ES §3). Naik ke MM 8.4 (Rp269.000) hanya bila terukur perlu.

## Konsekuensi

- Data aplikasi dan cadangannya berada di Indonesia pada dua badan hukum Indonesia. Pertanyaan transfer lintas batas tersisa pada push (Apple, Google) dan `AI` (ADR-021). Untuk VPS milik penyedia asing di pusat data Jakarta, pertanyaannya hanya muncul bila pemilik produk memilih 09-B C (AWS Lightsail, PT Amazon Web Services Indonesia); bila itu dipilih, advokat menilai apakah pemrosesan oleh induk usaha di luar negeri terjadi. Email (ADR-026) memakai AWS Jakarta dan hanya membawa alamat email dan teks templat.
- Kota pusat data Biznet Gio dan cakupan kalimat "di luar wilayah Republik Indonesia" di kebijakan privasinya belum terbukti tertulis (RISK-04). Keduanya harus selesai sebelum data nyata.
- Status `repo1-type=sftp` di pgBackRest 2.60.0 dan dukungan "penuh" PG18 belum dikukuhkan **[Perlu validasi]** (RISK-05). WAL-G v3.0.9 menjadi alternatif bila uji pemulihan gagal.
- Dua penyedia berarti dua tagihan dan dua akun untuk dijaga.
- L-14 (pendaftaran PSE) berlaku untuk kedua penyedia dan milik advokat, sebelum tahap publik (S4-4 (iii) A); risiko pemutusan akses bila pilot ternyata wajib terdaftar (W-32) diterima pemilik produk.
- Catatan ketentuan prosesor (18 §5.8 butir tinjauan 2) untuk kedua penyedia dibuat saat S4-2 diputuskan, sebelum data nyata, dan hanya tinjauan advokat atas teksnya yang menunggu tahap publik (dokumen 10 §3.1, §4.1 butir 4). Waktu itu, juga untuk DPIA, **diterima atas saran lead** (Q-208; pemilik produk menjawab "pasang semua", [_progress.md](../_progress.md), "Hal terakhir checkpoint S4"): DPIA dan catatan ketentuan prosesor adalah pekerjaan fase bangun, selesai sebelum data nyata dipakai.
- Bila sebelum tahap publik ada pengguna di luar rumah tangga pilot, advokat menjawab L-14 (bersama L-01 dan L-24) lebih dulu. Pemicu ini berlaku (Q-209, diterima dengan cara yang sama).
