# Isi contoh layar gerbang 1.8: Beranda, Catat, Chat

Pemilik: `ux-designer`. Dibuat 13 November 2026 untuk FS-1 langkah 1, tugas 1.8 (gerbang contoh layar, 08 §19.8 dan §25.11). **Diperbarui 13 November 2026 untuk arah pemilik produk dari gambar rujukan (08 §19.13)**, sesudah ketiga arah §19.12 ditolak. Satu arah, tanpa pengalih arah. Dipakai `web-dev` untuk membangun halaman statis di `design-samples/`. Rujukan "§" tanpa nama dokumen berarti `docs/08-ux-ui-strategy.md`.

Berkas ini hanya isi: kartu, baris, angka, nama, kalimat, token, dan glif. Warna, huruf, komponen, kisi, dan gerak mengikuti §19.13; nama token §19.12.2 dan §19.13.2.

---

## 0. Aturan untuk ketiga contoh

| Hal | Aturan | Rujukan |
|---|---|---|
| Data | **Rekaan seluruhnya.** Tidak ada nama, angka, tempat uang, atau pesan dari rumah tangga pilot | §19.8 "Dipakai ulang untuk wawancara" |
| Orang | Dua anggota dewasa: **Sari** dan **Andi**. Pelihat di ketiga contoh adalah **Andi**, jadi Andi disebut "kamu". Lead memastikan kedua nama bukan nama anggota rumah tangga pilot | §18.1 "Sapaan" |
| Rumah tangga | **"Rumah Kenanga"** (rekaan) | §6.1 bagian 0 |
| Tempat uang | Tanpa nama bank atau layanan e-wallet: "Rekening Andi", "Rekening bersama", "E-wallet Sari", "Dompet Andi", "Dompet Sari" | §18.2 |
| Hari | Hari ini Jumat, 13 November; kemarin Kamis. Bulan anggaran 1–30 November | §6.1 bagian 2 |
| Arah visual | **Satu**: §19.13. Hanya tema dan kontras yang dapat ditimpa lewat alamat (`?tema=terang|gelap`, `&kontras=tinggi`), supaya pemilik produk membandingkan di perangkat tanpa mengubah setelan sistem | §19.13.9 butir 5 |
| Huruf | **Figtree** untuk seluruh teks, angka maupun isi; `spasi-kata` 0 | §19.13.4 |
| Penulisan uang | "Rp" lalu angka dengan titik ribuan, tanpa spasi, tanpa desimal, tanpa "rb" atau "jt". "Rp" unsur tersendiri dari token | §17.3, §19.4 |
| Token angka | Setiap jumlah memakai `angka-utama`, `angka-besar`, atau `angka-baris` dengan `tabular-nums`. Baris berkata dan berjumlah: kata di kiri, jumlah di kanan sebagai unsur sendiri | §19.4; §19.8 pemeriksaan 13 |
| "Rp" diredam | Pada `angka-utama` dan `angka-besar`: 60%, `tinta-2`, garis dasar sama. Pada `angka-baris`: sama dengan angkanya | §19.13.5 "Jumlah besar" |
| Warna jumlah | Semua jumlah `tinta`; tidak pernah `aksen`, `sorot`, atau warna makna. Jumlah di kalimat keterangan memakai warna kalimatnya, `tinta-2` | §19.9; §19.13.2 aturan 4 |
| Jumlah besar di wadah sempit | `angka-utama` dan `angka-besar` berhenti membesar pada lebar wadahnya, tidak pernah di bawah 40 px dan 28 px; faktor container query statis dikalibrasi untuk jumlah terpanjang (§0.1 "Uji lebar") | §19.4; §25.4 |
| Glif | Dari `web/static/glyphs/`, warna `tinta-2` (di atas pilihan terpilih: `di-atas-aksen`), selalu bersama kata, `aria-hidden` | §19.5 |
| Ikon antarmuka | Material Symbols gaya garis sebagai SVG milik sendiri; tab terpilih simbol terisi. Nama simbol usulan, padanan bermakna sama boleh | §25.4 "Ikon" |
| Warna | Hanya lewat token di `tokens.css` | §19.13.2 |
| Tombol biru | Paling banyak satu tombol biru berlabel per layar: Beranda, tombol Catat; lembar, Simpan di kepala. Tombol utama kartu **hitam** (`tinta`, label `kartu`); tombol lain **abu** (`sisip`, label `tinta`) | §17.8 aturan 4; §19.13.2 aturan 1 dan 2 |
| Pilihan terpilih | Isi `aksen`, label dan glif `di-atas-aksen`, **ikon centang** di depan label | §19.13.5 "Pil pemilih dan chip" |
| Gerak | Menurut §19.13.7, termasuk bentuk `prefers-reduced-motion` | §19.13.7 |
| Kata | Tanpa kata templat §19.8 butir 3, tanpa kata §18.1 "Kata yang tidak dipakai", tanpa kosakata akuntansi, tanpa "Anda", tanda seru, atau emoji. Kepala kolom dan label huruf kalimat biasa, bukan kapital | §18, §19.8 |

### 0.1 Angka yang dipakai bersama

Lead menghitung ulang: 4.185.000 + 2.150.000 + 18.400.000 = 3.200.000 + 12.500.000 + 9.035.000 = 24.735.000.

| Hal | Angka | Hubungan |
|---|---|---|
| Uang kita | Rp24.735.000 | Jumlah ketiga baris pemegang (06 INV-A24; 07 §5.1) |
| Dipegang kamu (Dompet Andi, Rekening Andi) | Rp4.185.000 | Termasuk Rp200.000 yang menunggu kamu mengonfirmasi |
| Dipegang Sari (Dompet Sari, E-wallet Sari) | Rp2.150.000 | Termasuk Rp50.000 yang sedang dibahas |
| Dipegang bersama (Rekening bersama) | Rp18.400.000 | — |
| Disiapkan untuk tagihan | Rp3.200.000 | Alokasi; tidak dikurangkan dari Uang kita (PRN-17) |
| Disisihkan untuk tabungan | Rp12.500.000 | Alokasi |
| Belum disiapkan untuk apa pun | Rp9.035.000 | 24.735.000 − 3.200.000 − 12.500.000 (07 INV-17) |
| Uang pribadi saya (Andi) | Rp1.725.000 | Tidak masuk Uang kita (07 §5.2) |
| Serah terima Sari ke Andi | Rp200.000 | Dompet Sari ke Dompet Andi; menunggu Andi |
| Minta anggaran Sari | Rp7.850.000, "Kulkas baru" | Menunggu Andi sejak Kam 19.40 |
| Minta anggaran Andi | Rp650.000, "Servis motor" | Menunggu Sari sejak Rab 08.15 |
| Anggaran Dapur | Batas Rp2.500.000; terpakai Rp1.880.000, di antaranya Rp200.000 belum dikonfirmasi; sisa Rp620.000 | Segmen: terpakai dikonfirmasi 1.680.000 (67,2%), belum dikonfirmasi 200.000 (8%), sisa 620.000 (24,8%); jumlah 100% |
| Anggaran Transportasi | Batas Rp800.000, terpakai Rp460.000, sisa Rp340.000 | Segmen: terpakai 57,5%, sisa 42,5% |

**Uji lebar.** "Rp88.888.888" hanya disisipkan skrip uji `web-dev` ke tempat angka utama untuk menguji 320 CSS px dan teks 200% (§25.11 pemeriksaan 17); faktor container query statis di `samples.css` dikalibrasi untuk jumlah itu. Uji pasangan "Rp1.111.111" dan "Rp8.888.888" (§19.4) dijalankan dengan Figtree.

---

## 1. Beranda (LY-20)

Judul halaman: **"Casa · Beranda"** (§25.9).

### 1.1 Kerangka per ukuran (§19.13.6)

| | Ringkas (< 600) | Sedang (600–839) | Lebar (840–1199) | Sangat lebar (≥ 1200) |
|---|---|---|---|---|
| Navigasi | Dok mengambang; tombol Catat bulat biru di kanannya | Bilah atas 72: simpul kiri, grup tab tengah, tombol biru "Catat" kanan | Sama, ditambah kata "Casa" di samping simpul | Sama dengan Lebar |
| Baris kepala isi | Simpul 32 · `h1` "Rumah Kenanga"; kanan kosong | `h1` "Rumah Kenanga"; kanan kosong | `h1`; kanan tombol bulat "Chat di samping" | Sama dengan Lebar |
| Kolom | 1 | 1, paling lebar 640, di tengah | Utama (kartu 1–2) dan samping (kartu 3–5) | Kolom 1: kartu 1–2; kolom 2: kartu 3–4; kolom 3: kartu 5 |
| Margin; jarak antarkartu; padding kartu | 16; 12; 16 | 24; 16; 20 | 24; 16; 20 utama, 16 samping | 32; 20; 20 |

**Dok** (Ringkas): Beranda (`home`, terpilih: pil `pil-navigasi` dengan ikon terisi dan kata "Beranda" berwarna `navigasi`; lencana "2"), Uang (`account_balance_wallet`), Chat (`chat_bubble`, lencana "2"), Saya (`person`). Tab yang tidak terpilih hanya ikon, dengan nama aksesibel berupa teks tak tampak ("Uang", "Chat, 2 belum dibaca", "Saya"); lencana pil `sorot` berangka `di-atas-sorot`. Tombol Catat: bulat 56, `aksen`, ikon `add` `di-atas-aksen`, nama aksesibel "Catat".

**Grup tab** (Sedang ke atas): empat tab berikon dan berkata; terpilih pil `tinta` dengan label `kartu`; `aria-current="page"`; lencana sama dengan dok. Tombol "Catat": pil biru 48 dengan ikon `add` dan kata "Catat".

### 1.2 Kartu, urutan dari atas

Setiap kartu: `kartu`, sudut 24, kepala kartu tinggi 44 (lingkaran ikon 40 `sisip` dengan ikon `aksen-tinta`, jarak 12, judul `judul-bagian`, kendali kanan bila ada).

**Kartu 1, Uang kita** (§6.1 bagian 2 sampai 4)

| # | Unsur | Isi persis | Tampil sebagai |
|---|---|---|---|
| 1 | Kepala kartu | Ikon `account_balance_wallet` · "Uang kita"; kanan kosong ("Sembunyikan angka" Tier 3 di web) | Kepala kartu |
| 2 | Angka utama | Rp24.735.000, rata kiri | `angka-utama` (2,5 rem Ringkas, 3 rem mulai Sedang, −0,03 em); "Rp" diredam |
| 3 | Keterangan | "Bulan ini, 1–30 November" | Teks kecil `tinta-2` |
| 4 | Garis | Satu garis `garis` selebar isi kartu | `garis` |
| 5 | Baris pemegang 1 | Glif **tunai** dan **rekening** bertumpuk di kolom glif 18 · "Dipegang kamu" · Rp4.185.000 rata kanan | Glif `tinta-2`; kata `tinta`; `angka-baris` |
| 5a | Baris anak | Ikon `info` · "Menunggu kamu mengonfirmasi" · Rp200.000. Tanpa wadah | Ikon dan kata teks terkecil `tinta-2`; `angka-baris` |
| 6 | Baris pemegang 2 | Glif **tunai** dan **e-wallet** · "Dipegang Sari" · Rp2.150.000 | Sama dengan 5 |
| 6a | Baris anak | Ikon `info` · "Sedang dibahas" · Rp50.000 | Sama dengan 5a |
| 7 | Baris pemegang 3 | Glif **rekening** · "Dipegang bersama" · Rp18.400.000 | Sama dengan 5 |
| 8 | Panel penjelas | Panel `sisip` sudut 16, padding 12, di dasar kartu: "Disiapkan untuk tagihan" Rp3.200.000 · "Disisihkan untuk tabungan" Rp12.500.000 · "Belum disiapkan untuk apa pun" Rp9.035.000 | Kata teks terkecil `tinta-2`; `angka-baris` `tinta` rata kanan |

**Kartu 2, Uang pribadi saya** (§6.1 bagian 5): kepala kartu ikon `lock` · "Uang pribadi saya" · Rp1.725.000 rata kanan (`angka-baris`) di baris kepala yang sama; di bawahnya "Hanya kamu yang melihat ini." `tinta-2`.

**Kartu 3, Menunggu** (§6.1 bagian 6)

| # | Unsur | Isi persis | Tampil sebagai |
|---|---|---|---|
| 1 | Kepala kartu | Ikon `schedule` · "Menunggu"; kanan pil "Lihat semua" (tautan ke LY-21) | Pil `sisip`, label `tinta` 500, tinggi 40, sasaran 44 |
| 2 | Subjudul | "Menunggu kamu" | `judul-subbagian` |
| 3 | Butir 1 | Lingkaran glif **tunai** (`sisip`, glif `tinta-2`) · "Terima dari Sari" · Rp200.000; di bawahnya tombol hitam "Sudah terima" tinggi 44 | `angka-baris`; tombol `tinta` + `kartu` |
| 4 | Butir 2 | "Sari minta anggaran" · di bawahnya "Kulkas baru" `tinta-2` · Rp7.850.000 · `chevron_right`; seluruh baris satu tautan | `angka-baris` |
| 5 | Subjudul | "Menunggu Sari" | `judul-subbagian` |
| 6 | Butir 3 | "Kamu minta anggaran" · di bawahnya "Servis motor, sejak Rab 08.15" `tinta-2` · Rp650.000; tanpa tombol | `angka-baris` |

**Kartu 4, Minggu ini** (§6.1 bagian 7): kepala kartu ikon `calendar_month` · "Minggu ini". Baris: "Listrik" · "Sen, 16 Nov" · Rp850.000; "Cicilan motor" · "Rab, 18 Nov" · Rp1.275.000; "Makan malam" · "sisa" · Rp25.000 dengan tombol abu "Selesaikan" tinggi 44.

**Kartu 5, Anggaran** (§6.1 bagian 8)

| # | Unsur | Isi persis | Tampil sebagai |
|---|---|---|---|
| 1 | Kepala kartu | Ikon `pie_chart` · "Anggaran"; kanan pil "Semua anggaran" | Pil `sisip` |
| 2 | Baris Dapur | "Dapur" · "sisa" Rp620.000 rata kanan; batang bersegmen selebar baris: `tinta` 67,2% · `aksen` berarsir 8% · `sorot` 24,8%, celah 3 px; kalimat "dari Rp2.500.000. Rp200.000 yang terpakai belum dikonfirmasi." `tinta-2` | Batang §19.13.5, tinggi 10, `aria-hidden`; `angka-baris`; jumlah di kalimat `angka-kalimat` |
| 3 | Baris Transportasi | "Transportasi" · "sisa" Rp340.000 · batang `tinta` 57,5% · `sorot` 42,5% · "dari Rp800.000" | Sama |
| 4 | Keterangan | Satu baris di bawah batang terakhir: titik `tinta` "Terpakai" · titik `aksen` "Belum dikonfirmasi" · titik `sorot` "Sisa" | Titik 8 px; teks terkecil `tinta-2`; `aria-hidden` (kalimat di baris sudah menyebut angkanya) |

**Pembaca layar.** Angka utama dibacakan "Uang kita, dua puluh empat juta tujuh ratus tiga puluh lima ribu rupiah" (§15). Lencana Beranda "2 menunggu", Chat "2 belum dibaca".

### 1.3 Panel Chat (Lebar ke atas)

Satu varian dengan panel Chat terbuka sesudah pengguna menekan "Chat di samping": `aside` berisi utas §3 tanpa dok dan tanpa bilah atas sendiri, padding 16. Lebar: menggantikan kolom samping. Sangat lebar: menggantikan kolom 3; kartu Anggaran pindah ke bawah Minggu ini di kolom 2.

### 1.4 Informasi peka (Beranda)

| Hal | Di aplikasi | Di contoh statis |
|---|---|---|
| Samaran angka | Berlaku untuk setiap jumlah saat fiturnya dibangun (§12.1, §12.2) | Tidak tampil (Tier 3 di web) |
| Halaman tersembunyi | Penutup netral `kertas` dengan simpul (§25.9) | Tidak dibangun |
| Judul halaman | "Casa · Beranda", tanpa nama rumah tangga atau angka | Sama |

---

## 2. Catat (LY-32, jenis Belanja)

Judul halaman saat lembar terbuka: **"Casa · Catat"**. Lembar di atas Beranda §1, diredupkan `selubung`. Ringkas: menempel di bawah, selebar layar, paling tinggi 90% viewport dinamis, sudut atas 28. Sedang ke atas: di tengah, paling lebar 560, sudut 28. Permukaan `permukaan-mengambang` (= `kartu`).

| # | Urutan §9.1 | Unsur | Isi persis | Tampil sebagai |
|---|---|---|---|---|
| 1 | — | Kepala, kiri | Tombol bulat `close`, nama aksesibel "Tutup" (menyimpan draf) | Bulat 44 `sisip`, ikon `tinta` |
| 2 | — | Kepala, tengah | `select` bawaan "Belanja" (Belanja, Uang masuk, Kasih ke, Terima dari, Pindah, Minta anggaran, Jalan-jalan), bergaya pil | Pil `sisip`, `expand_more` |
| 3 | 7 | Kepala, kanan | Tombol biru "Simpan", tinggi 48; satu-satunya tombol biru selama lembar terbuka | `aksen`, `di-atas-aksen` |
| 4 | 1 | Label | "Jumlah" (tampak) | Teks isi 500 |
| 5 | 1 | Isian jumlah | Panel `sisip` sudut 20 berisi "Rp" diredam lalu **87.000**; langsung difokus, cincin fokus `aksen-tinta` | `angka-besar` |
| 6 | 1 | Tombol "000" | Tombol abu di samping kanan isian, sama tinggi; nama aksesibel "000, tambah tiga nol" (teks tak tampak ", tambah tiga nol", tanpa `aria-label`) | `sisip`, `tinta` |
| 7 | 2 | Label | "Tempat uang" | Teks isi 500 |
| 8 | 2 | Chip 1, **terpilih** | Centang · glif **tunai** · "Dompet Andi" | `aksen`; label dan glif `di-atas-aksen` |
| 9 | 2 | Chip 2 | Glif **rekening** · "Rekening Andi" | `kartu` dengan garis `garis`; glif `tinta-2` |
| 10 | 2 | Chip 3 | Glif **rekening** · "Rekening bersama" | Sama |
| 11 | 3 | Kategori | Satu baris tombol di panel `sisip` sudut 16: "Kategori" kecil `tinta-2` di atas "Dapur" · "sisa" Rp620.000 rata kanan · `chevron_right` | `angka-baris` |
| 12 | 4 | Tanggal | Label "Tanggal"; chip "Hari ini" (**terpilih**, centang), "Kemarin", "Pilih tanggal" | Chip §19.13.5 |
| 13 | 5 | Rincian terlipat | `details` tertutup "Tambah rincian" | Teks `aksen-tinta` |
| 14 | 6 | Saran kontekstual | **Tidak ada** di contoh ini | — |
| 15 | 7 | Simpan di akhir isi | Tombol hitam "Simpan" selebar lembar, perintah yang sama dengan unsur 3 | `tinta`, `kartu` |

**Isian jumlah.** `inputmode="numeric"`, `autocomplete="off"`, huruf isian tidak di bawah 1 rem. **Gerak:** lembar naik 280 ms (Ringkas) atau memudar dan membesar dari 0,96 (Sedang ke atas); §19.13.7.

**Informasi peka.** Isian jumlah tetap terlihat saat angka disamarkan, sisa anggaran tersamar (Tier 3 di web, tidak tampil). Judul halaman tanpa jenis atau jumlah selain "Catat".

---

## 3. Chat (LY-60)

Judul halaman: **"Casa · Chat"**.

### 3.1 Kepala dan kerangka

Baris kepala: `h1` "Sari" (`judul-layar`), di bawahnya "online" teks kecil `tinta-2`; kanan kosong. Tanpa tombol kembali. Ringkas: dok di tengah dengan Chat terpilih (tanpa lencana; Beranda tetap "2"), **tanpa tombol Catat bulat**; Catat lewat tombol tambah di kotak tulis. Sedang ke atas: bilah atas dengan Chat terpilih; utas satu kolom paling lebar 640 di tengah.

### 3.2 Utas, dari atas

Gelembung paling lebar 80% dan 480 px, sudut 20, sudut di sisi pengirim 6; waktu kecil `tinta-2` di dalam gelembung. Tanpa foto profil atau avatar (§10.1).

| # | Jenis | Sisi | Isi persis | Tampil sebagai |
|---|---|---|---|---|
| 1 | Pemisah hari | Tengah | "Kemarin" | Teks kecil `tinta-2` |
| 2 | Baris peristiwa | Tengah | Glif **simpul** · "Disepakati: Jatah pribadi Sari, 16.10" | Teks kecil `tinta-2`, tanpa gelembung |
| 3 | Pesan Sari | Kiri | Nama "Sari" kecil di atasnya; "Kulkasnya makin berisik, tadi pagi bocor lagi." · 19.32 | `kartu`, teks `tinta` |
| 4 | Pesan Sari | Kiri | "Aku sudah lihat yang dua pintu di toko dekat pasar." · 19.33 | Sama, tanpa nama |
| 5 | **Kartu permintaan** | Selebar utas | §3.3 kartu A | — |
| 6 | Pesan sendiri | Kanan | "Aku lihat dulu, besok pagi kita bahas ya." · 21.05 · `done_all` dengan kata "Dibaca" bagi pembaca layar | `aksen-muda`, teks `tinta`; ikon `aksen-tinta` |
| 7 | Pemisah hari | Tengah | "Hari ini" | Sama dengan 1 |
| 8 | Pesan Sari | Kiri | Nama "Sari"; "Uang belanja minggu ini aku taruh di laci dapur." · 08.02 | Sama dengan 3 |
| 9 | **Kartu serah terima** | Selebar utas | §3.3 kartu B | — |

### 3.3 Kartu objek

`kartu` dengan garis 1 px `garis`, sudut 24, selebar kolom utas dan paling lebar 480, tidak berpihak. Kepala kartu: lingkaran ikon 40, jenis objek sebagai judul kecil, dan **pil status** di kanan. Tombol sama tinggi 44, sebaris sama lebar; pada lebar sempit atau teks besar menjadi tumpukan dengan tombol hitam di atas. Nama aksesibel tombol menyebut objeknya ("Setuju, minta anggaran Kulkas baru").

**Kartu A: permintaan menunggu, pelihat sebagai pemberi tanggapan**

| Butir §10.4 | Isi persis | Tampil sebagai |
|---|---|---|
| Kepala | Ikon `request_quote` · "Minta anggaran" · pil status "Menunggu kamu" | Lingkaran `sisip`; pil netral `sisip` + `tinta` |
| (2) Judul | "Kulkas baru" | `judul-bagian` |
| (3) Jumlah | Rp7.850.000 | `angka-besar`, "Rp" diredam |
| (4) Keadaan | "Menunggu kamu sejak Kam 19.40" | Teks kecil `tinta-2` |
| (5) Tombol | "Setuju" (hitam) · "Usul lain" (abu) · "Belum sepakat" (abu, tidak berwarna peringatan) | §19.13.5 "Tombol" |
| (6) | "Lihat rincian" | Tautan `aksen-tinta`, bergaris bawah |

**Kartu B: serah terima, pelihat sebagai penerima** (motif angka bersama pemegangnya)

| Butir §10.4 | Isi persis | Tampil sebagai |
|---|---|---|
| Kepala | Glif **tunai** di lingkaran `sisip` (`tinta-2`) · "Serah terima" · pil status "Menunggu kamu" | Pil netral |
| (2) Judul | "Sari kasih ke kamu" | `judul-bagian` |
| (3) Jumlah | Rp200.000, rata kiri | `angka-besar` |
| Motif | Garis `garis`, lalu glif **tunai** · "Dari Dompet Sari"; glif **tunai** · "Ke Dompet Andi" | §19.8 butir 5 |
| (4) Keadaan | "Menunggu kamu sejak 08.03" | Teks kecil `tinta-2` |
| (5) Tombol | "Sudah terima" (hitam) · "Tidak sesuai" (abu); sama tinggi | §9.7 |
| (6) | "Lihat rincian" | Tautan `aksen-tinta` |

### 3.4 Kotak tulis

Pil `kartu` tinggi 52, menempel 8 di atas dok (Ringkas) atau di dasar kolom utas (Sedang ke atas).

| Unsur | Isi persis |
|---|---|
| Tombol tambah | Bulat 44 `sisip`, ikon `add`, nama aksesibel "Tambah" (Foto, Kamera, Catat belanja, Minta anggaran, Kirim catatan yang sudah ada; menu tidak dibuka di contoh) |
| Isian | `textarea` kosong, label pembaca layar "Pesan untuk Sari", teks contoh "Tulis pesan" |
| Tombol kirim | Bulat 44 `aksen`, ikon `send` `di-atas-aksen`, nama aksesibel "Kirim" (tidak dihitung sebagai tombol biru, §17.8 aturan 4) |

### 3.5 Informasi peka (Chat)

Jumlah di kartu tersamar bila "Sembunyikan angka" menyala; teks bebas tidak (Tier 3 di web, tidak tampil). Judul halaman "Casa · Chat" tanpa nama Sari.

---

## 4. Ikon dan nama di layar utama

Satu set, satu manifest (`design-samples/manifest.webmanifest`). Hanya varian terang; varian gelap **[Native]**.

| Hal | Isi |
|---|---|
| Ikon `any` | `web/static/icons/casa-icon.svg`: simpul `sorot` `#C6F432` di atas `aksen` `#1757F0` (4,4:1), satu bentuk besar; PNG `icon-192.png`, `icon-512.png`, ikon layar utama Apple `icon-180.png` |
| Ikon `maskable` | `casa-icon-maskable.svg` dan `icon-maskable-512.png`: gambar yang sama, simpul seluruhnya di dalam lingkaran berjari-jari 295 px dari pusat kanvas 1024 |
| Nama | "Casa" (`name`, `short_name`), tanpa data |
| `theme-color` | `kertas`: `#F4F4F4` terang, `#111111` gelap |
| Simpul di baris kepala Ringkas | Bentuk dan warna ikon `any`, 32 px, `aria-hidden` |

---

## 5. Pemeriksaan sendiri 11 sampai 15 (§19.8) atas isi ini

| # | Pemeriksaan | Hasil atas isi | Yang dijaga `web-dev` |
|---|---|---|---|
| 11 | Glif dan ikon sendiri | Lulus: tempat uang selalu glif tunai, rekening, atau e-wallet; ikon aplikasi simpul | Simbol Material `payments`, `account_balance`, `wallet`, `savings` tidak dipakai di tempat glif. `account_balance_wallet` hanya ikon tab Uang dan kepala kartu Uang kita |
| 12 | Tanpa hiasan | Lulus: tanpa ilustrasi, gradasi, karakter, foto, spanduk, mosaik, atau kartu promosi. Arsiran segmen adalah penanda "belum dikonfirmasi" | Bayangan hanya dok, tombol Catat bulat, dan lembar |
| 13 | Angka bertoken | Lulus | `tabular-nums`; "Rp" unsur tersendiri |
| 14 | Motif hadir | Lulus: kartu Uang kita (baris 2 sampai 7) dan kartu serah terima | Urutan label, angka rata kiri, garis, baris pemegang tidak diubah per lebar |
| 15 | Kata templat | Lulus | Teks tombol, nama aksesibel, dan judul halaman juga tunduk |

---

## 6. Yang sengaja tidak ada

Tombol "Sembunyikan angka" dan samaran (Tier 3 di web); kata penyemangat (Tier 2); saran kata di chat (Tier 2); penutup netral dan keadaan luring (layar 1.10); grafik kolom dan tabel transaksi (bukan Beranda, Catat, atau Chat; komponennya di §19.13.5 untuk LY-31 dan ringkasan Tier 2); warna makna selain pil status netral (tidak ada keadaan yang memerlukannya di data rekaan ini); tren dan perbandingan antar-orang (§6.1).
