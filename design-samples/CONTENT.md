# Isi contoh layar gerbang 1.8: Beranda, Catat, Chat

Pemilik: `ux-designer`. Dibuat 13 November 2026 untuk FS-1 langkah 1, tugas 1.8 (gerbang contoh layar, 08 §19.8 dan §25.11); diperbarui 13 November 2026 untuk gerbang ulang dengan tiga arah visual (08 §19.12) sesudah tinjauan QG-09. Dipakai `web-dev` untuk membangun halaman statis di `design-samples/`. Rujukan "§" tanpa nama dokumen berarti `docs/08-ux-ui-strategy.md`.

Berkas ini hanya isi: bagian, baris, angka, nama, kalimat, token yang dipakai, dan glif. Tata letak per lebar mengikuti §4.2, §6.3, §16.1, §17.8, §25.2, dan §25.3; bentuk komponen mengikuti §25.4; warna, huruf, dan bentuk per arah mengikuti §19.12. Nama token mengikuti §19.12.2; nama lama di tabel itu tidak dipakai lagi.

---

## 0. Aturan untuk ketiga contoh

| Hal | Aturan | Rujukan |
|---|---|---|
| Data | **Rekaan seluruhnya.** Tidak ada nama, angka, tempat uang, atau pesan dari rumah tangga pilot | §19.8 "Dipakai ulang untuk wawancara" |
| Orang | Dua anggota dewasa: **Sari** dan **Andi**, keduanya nama contoh yang sudah dipakai 08 (§6.1, §10.4, §9.5). Pelihat di ketiga contoh adalah **Andi**, jadi Andi disebut "kamu". Lead memastikan kedua nama bukan nama anggota rumah tangga pilot sebelum contoh ditunjukkan | Brief tugas; §18.1 "Sapaan" |
| Rumah tangga | **"Rumah Kenanga"** (rekaan; lead memeriksanya bersama kedua nama) | §6.1 bagian 0 |
| Tempat uang | Tanpa nama bank atau layanan e-wallet: "Rekening Andi", "Rekening bersama", "E-wallet Sari", "Dompet Andi", "Dompet Sari". Bagian `{bank}` dan `{layanan}` di §18.2 sengaja dikosongkan supaya contoh yang juga dipakai untuk wawancara 13 §6.3 tidak memuat merek | §18.2 baris "Tempat uang" |
| Hari | Hari ini Jumat, 13 November; kemarin Kamis. Bulan anggaran 1–30 November | §6.1 bagian 2 |
| Arah visual | Tiga arah dari satu set halaman, dialihkan lewat alamat (`?arah=limau|koral|biru`, `tema=terang|gelap`, `kontras=tinggi`); isi dan data sama di ketiga arah | §19.12.11 |
| Huruf | Satu keluarga huruf per arah untuk seluruh teks, angka maupun isi: Manrope (Arah 1), Plus Jakarta Sans (Arah 2, dengan `spasi-kata` 0,08 em), Figtree (Arah 3) | §19.12.6 |
| Penulisan uang | "Rp" lalu angka dengan titik ribuan, tanpa spasi, tanpa desimal, tanpa "rb" atau "jt". "Rp" ditulis sebagai unsur tersendiri dari token, bukan bagian teks angka | §17.3, §19.4 |
| Token angka | Setiap jumlah memakai satu token §19.4 (`angka-utama`, `angka-besar`, atau `angka-baris`) dengan `tabular-nums`. Tidak ada jumlah di dalam kalimat teks biasa: bila sebuah baris berisi kata dan jumlah, kata di kiri dan jumlah di kanan sebagai unsur sendiri (pemeriksaan 13) | §19.4, §19.8 butir 4 |
| "Rp" kecil | Pada `angka-utama` dan `angka-besar`: 60% ukuran angka, garis dasar sama, warna `tinta-2`. Pada `angka-baris`: ukuran dan warna sama dengan angkanya | §19.4 "Aturan angka"; §19.12.6 |
| Warna jumlah | Semua jumlah `tinta`, tidak pernah `aksen`, `aksen-tinta`, atau warna makna. Satu-satunya pengecualian: jumlah di dalam kalimat keterangan ("dari Rp2.500.000") memakai warna kalimatnya, `tinta-2` | §19.9; §19.12.4; §19.4 "Jumlah di baris dan di kalimat" |
| Jumlah besar di wadah sempit | `angka-utama` dan `angka-besar` berhenti membesar pada lebar wadahnya, tidak pernah di bawah 40 px dan 28 px; di contoh statis faktor container query tetap dikalibrasi untuk jumlah terpanjang yang mungkin muncul (lihat §0.1 "Uji lebar") | §19.4 "Teks besar"; §25.4 "Batas lebar jumlah" |
| Glif | Dari `web/static/glyphs/`, warna `tinta-2`, selalu bersama kata. Glif hiasan bagi pembaca layar (`aria-hidden`), karena katanya sudah ada | §19.5; §19.12.9; WCAG 1.4.1 |
| Ikon antarmuka | Material Symbols gaya garis sebagai SVG milik sendiri; tab terpilih memakai simbol terisi. Nama simbol di bawah adalah usulan, `web-dev` boleh memilih padanan yang maknanya sama | §17.5, §25.4 "Ikon" |
| Warna | Hanya lewat token §19.12 di `tokens.css`; tidak ada warna yang ditulis langsung di halaman | §19.12.11; §25.4 "Token" |
| Tombol berisi aksen | Paling banyak satu tombol berlabel berisi `aksen` per layar (Beranda: Catat; lembar: Simpan di kepala). Tombol kartu: tonal (`aksen-muda`, teks `aksen-tinta`) untuk tindakan utamanya, bergaris (`garis-tegas`, teks `aksen-tinta`) untuk yang lain | §17.8 aturan 4; §25.4 "Tombol utama dan tombol kartu" |
| Kata | Tanpa kata dari daftar kata templat §19.8 butir 3, tanpa kata §18.1 "Kata yang tidak dipakai", tanpa kosakata akuntansi §18.2, tanpa "Anda", tanpa tanda seru, tanpa emoji. Baris keterangan dan baris peristiwa dengan koma, bukan rangkaian titik tengah | §18, §18.1, §19.8 |

### 0.1 Angka yang dipakai bersama

Semua angka ketiga contoh berasal dari keadaan rekaan yang sama, supaya Beranda dan Chat saling cocok. Angka diganti 13 November 2026 dari deret "88.888.888" yang terbaca sebagai data uji (tinjauan QG-09, S4); lead menghitung ulang: 4.185.000 + 2.150.000 + 18.400.000 = 3.200.000 + 12.500.000 + 9.035.000 = 24.735.000.

| Hal | Angka | Hubungan |
|---|---|---|
| Uang kita | Rp24.735.000 | Jumlah ketiga baris pemegang (06 INV-A24; 07 §5.1) |
| Dipegang kamu (Dompet Andi, Rekening Andi) | Rp4.185.000 | Termasuk Rp200.000 yang menunggu kamu mengonfirmasi |
| Dipegang Sari (Dompet Sari, E-wallet Sari) | Rp2.150.000 | Termasuk Rp50.000 yang sedang dibahas |
| Dipegang bersama (Rekening bersama) | Rp18.400.000 | — |
| Pemeriksaan | 4.185.000 + 2.150.000 + 18.400.000 = 24.735.000 | Lead memeriksa dengan skrip (CLAUDE.md "Computed, not read") |
| Disiapkan untuk tagihan | Rp3.200.000 | Alokasi; tidak dikurangkan dari Uang kita (PRN-17) |
| Disisihkan untuk tabungan | Rp12.500.000 | Alokasi |
| Belum disiapkan untuk apa pun | Rp9.035.000 | 24.735.000 − 3.200.000 − 12.500.000 (07 INV-17) |
| Uang pribadi saya (Andi) | Rp1.725.000 | Tidak masuk Uang kita (07 §5.2) |
| Serah terima Sari ke Andi | Rp200.000 | Dompet Sari ke Dompet Andi; menunggu Andi |
| Minta anggaran Sari | Rp7.850.000, "Kulkas baru" | Menunggu Andi sejak Kam 19.40 |
| Minta anggaran Andi | Rp650.000, "Servis motor" | Menunggu Sari sejak Rab 08.15 |
| Anggaran Dapur | Batas Rp2.500.000, terpakai Rp1.880.000 (di antaranya Rp200.000 belum dikonfirmasi), sisa Rp620.000 | Batang 75,2% (1.880.000 / 2.500.000); bagian belum dikonfirmasi 8% (200.000 / 2.500.000) di ujung isi |
| Anggaran Transportasi | Batas Rp800.000, terpakai Rp460.000, sisa Rp340.000 | Batang 57,5% |

**Uji lebar.** "Rp88.888.888" **tidak lagi tampil** di contoh yang ditunjukkan. Ia hanya **disisipkan oleh skrip uji** `web-dev` (uji alir ulang dan uji tabular) ke tempat angka utama, untuk menguji lebar pada 320 CSS px dan teks 200% (§25.11 pemeriksaan 17); angka 8 adalah angka terlebar bila fitur tabular gagal. Faktor container query statis di `samples.css` tetap dikalibrasi untuk jumlah itu, supaya contoh yang ditunjukkan juga lolos bila angkanya kelak lebih panjang. Uji pasangan "Rp1.111.111" dan "Rp8.888.888" (§19.4) dijalankan `web-dev` di halaman uji sendiri.

---

## 1. Beranda (LY-20)

Judul halaman: **"Casa · Beranda"** (§25.9; SEC-175 (a)).

### 1.1 Ringkas (iPhone 13, Galaxy A55, jendela < 600 px)

Urutan dari atas. Kolom "Tampil sebagai" menyebut token; huruf adalah huruf arah (§19.12.6).

| # | Bagian §6.1 | Unsur | Isi persis | Tampil sebagai | Rujukan |
|---|---|---|---|---|---|
| 1 | 0 Judul | Bilah atas, `h1` | "Rumah Kenanga" | `judul-layar` (1,5 rem 700); latar `kertas` | §6.1 bagian 0; §19.4; §25.4 "Bilah atas" |
| 2 | 0 Judul | Tombol "Sembunyikan angka" | **Tidak tampil** di contoh: di web fiturnya Tier 3 | — | §12.1; §25.10 |
| 3 | 1 Kata penyemangat | — | **Tidak tampil**: Tier 2 | — | §6.1 bagian 1 |
| 4 | 2 Uang kita | Kartu | Satu-satunya kartu di Beranda: `kartu`, sudut `--sudut-kartu` arah (20, 24, atau 12), padding lebih besar, judul di dalam kartu; Arah 3 dengan garis 1 px `garis`. Memuat unsur 5 sampai 11 | `kartu` | §17.8 aturan 7; §19.12.7 |
| 5 | 2 Uang kita | Label | "Uang kita" | `judul-bagian`, `tinta` | §19.4; §18.2 |
| 6 | 2 Uang kita | Angka utama | Rp24.735.000, rata kiri | `angka-utama`; "Rp" kecil `tinta-2` | §19.2 butir 1; §19.4; §19.8 butir 5 |
| 7 | 2 Uang kita | Keterangan | "Bulan ini, 1–30 November" | Teks kecil `tinta-2` | §6.1 bagian 2 |
| 8 | 3 Pemegang | Garis tipis | Satu garis `garis` selebar isi kartu | `garis` | §19.8 butir 5 |
| 9 | 3 Pemegang | Baris pemegang 1 | Glif **tunai** dan **rekening** bertumpuk di kolom glif 18 px · "Dipegang kamu" · Rp4.185.000 rata kanan | Glif `tinta-2`; kata `tinta`; jumlah `angka-baris` | §6.1 bagian 3; §19.8 butir 5 |
| 9a | 3 Pemegang | Baris anak, di bawah baris 9 | Ikon info (`info`) di kolom glif · "Menunggu kamu mengonfirmasi" di kiri, sejajar dengan teks nama · Rp200.000 di kanan. **Tanpa wadah berwarna** | Ikon `tinta-2`; kata teks terkecil `tinta-2`; jumlah `angka-baris` | §6.1 bagian 3; §17.8 aturan 3; §19.8 butir 5 |
| 10 | 3 Pemegang | Baris pemegang 2 | Glif **tunai** dan **e-wallet** · "Dipegang Sari" · Rp2.150.000 rata kanan | Sama dengan baris 9 | §6.1 bagian 3 |
| 10a | 3 Pemegang | Baris anak, di bawah baris 10 | Ikon info · "Sedang dibahas" · Rp50.000. Tanpa wadah berwarna | Sama dengan 9a | §9.6; §18.1 "Sedang dibahas" |
| 11 | 3 Pemegang | Baris pemegang 3 | Glif **rekening** · "Dipegang bersama" · Rp18.400.000 rata kanan | Sama dengan baris 9 | §6.1 bagian 3 ("tempat uang bersama") |
| 12 | 4 Penjelas angka | Tiga baris di atas latar, di bawah kartu Uang kita, sejajar dengan teks kartu | "Disiapkan untuk tagihan" Rp3.200.000 · "Disisihkan untuk tabungan" Rp12.500.000 · "Belum disiapkan untuk apa pun" Rp9.035.000 | Kata teks terkecil `tinta-2`; jumlah `angka-baris` `tinta`, rata kanan | §6.1 bagian 4; §18.2 |
| 13 | 5 Uang pribadi saya | Garis pemisah, lalu satu baris | "Uang pribadi saya" · Rp1.725.000 rata kanan; di bawahnya "Hanya kamu yang melihat ini." | Garis `garis`; jumlah `angka-baris`; keterangan `tinta-2` | §6.1 bagian 5; §18.3 |
| 14 | 6 Menunggu | Judul di atas latar | "Menunggu" | `judul-bagian` | §6.1 bagian 6; §17.8 aturan 7 |
| 15 | 6 Menunggu | Subjudul | "Menunggu kamu" | `judul-subbagian` | §6.1 bagian 6; §19.4 |
| 16 | 6 Menunggu | Butir 1 (tindakan ringan) | Glif **tunai** · "Terima dari Sari" · Rp200.000; tombol bergaris "Sudah terima" di bawah, tinggi ≥ 48 px | Kelompok daftar `kartu` sudut `--sudut-kelompok` (Arah 3 tanpa wadah); jumlah `angka-baris`; tombol bergaris `garis-tegas`, teks `aksen-tinta` | §6.1 bagian 6; §9.2; §15 "Sasaran sentuh"; §19.12.5 |
| 17 | 6 Menunggu | Butir 2 (membuka rincian) | "Sari minta anggaran" · di bawahnya "Kulkas baru" `tinta-2` · Rp7.850.000 · ikon `chevron_right`. Seluruh baris satu tautan | Jumlah `angka-baris` | §6.1 bagian 6; §25.4 "Daftar dan bagian"; §17.8 "Panel samping" |
| 18 | 6 Menunggu | Subjudul | "Menunggu Sari" | `judul-subbagian` | §6.1 bagian 6 |
| 19 | 6 Menunggu | Butir 3 (tanpa tombol) | "Kamu minta anggaran" · di bawahnya "Servis motor, sejak Rab 08.15" `tinta-2` · Rp650.000 | Jumlah `angka-baris` | §6.1 bagian 6 ("tanpa tombol"); §18.1 |
| 20 | 6 Menunggu | Tautan | "Lihat semua" | Teks `aksen-tinta`, bergaris bawah | §6.1 bagian 6; §19.12.5 |
| 21 | 7 Minggu ini | Judul | "Minggu ini" | `judul-bagian` | §19.4 |
| 22 | 7 Minggu ini | Baris tagihan | "Listrik" · di bawahnya "Sen, 16 Nov" · Rp850.000 | Jumlah `angka-baris` | §6.1 bagian 7 |
| 23 | 7 Minggu ini | Baris cicilan | "Cicilan motor" · di bawahnya "Rab, 18 Nov" · Rp1.275.000 | Jumlah `angka-baris` | §6.1 bagian 7 |
| 24 | 7 Minggu ini | Baris persetujuan terbuka | "Makan malam" · di bawahnya "sisa" · Rp25.000 · tombol bergaris "Selesaikan" | Jumlah `angka-baris`; tombol bergaris, teks `aksen-tinta` | §6.1 bagian 7 |
| 25 | 8 Anggaran | Judul | "Anggaran" | `judul-bagian` | §19.4 |
| 26 | 8 Anggaran | Baris 1 | "Dapur" · "sisa" Rp620.000 rata kanan · di bawahnya batang selebar baris, isi 75,2%, dengan 8% terakhir dari isi itu sebagai bagian belum dikonfirmasi · kalimat "dari Rp2.500.000. Rp200.000 yang terpakai belum dikonfirmasi." `tinta-2` | Batang menurut arah: isi `--isi-batang` di atas `lintasan`, tebal `--tebal-batang`; bagian belum dikonfirmasi diarsir (Arah 1 dan 2) atau putus-putus (Arah 3); `aria-hidden`. Jumlah `angka-baris`; jumlah di kalimat `angka-kalimat` | §6.1 bagian 8; §19.9; §19.12.5; §25.4 "Batang" |
| 27 | 8 Anggaran | Baris 2 | "Transportasi" · "sisa" Rp340.000 · batang isi 57,5% · "dari Rp800.000" | Sama | Sama |
| 28 | 8 Anggaran | Tautan | "Semua anggaran" | Teks `aksen-tinta`, bergaris bawah | §6.1 bagian 8 |
| 29 | — | Tombol Catat mengambang | "Catat" dengan ikon `add`, kanan bawah di atas bilah navigasi; wadah `aksen`, isi `di-atas-aksen`, sudut `--sudut-tombol-utama`, bayangan `bayangan` di tema terang | `aksen`, `di-atas-aksen` | §25.2 baris Ringkas; §19.12.7 |
| 30 | — | Bilah navigasi bawah | Beranda (`home`, terpilih, lencana "2"), Uang (`account_balance_wallet`), Chat (`chat_bubble`, lencana "2"), Saya (`person`). Label kata selalu tampil; `aria-current="page"` pada yang terpilih | Latar `navigasi`, ikon dan kata `di-atas-navigasi`. Terpilih: Arah 1 pil `aksen` dengan ikon terisi dan kata `di-atas-aksen` 600; Arah 2 pil `aksen-muda` dengan kata `aksen-tinta` 600; Arah 3 pil `aksen-muda` dengan ikon dan kata `aksen` | §4.1; §25.2; §25.4 "Navigasi"; §19.12.5 |

**Batang anggaran.** Isinya adalah bagian yang **terpakai** (1.880.000 dari 2.500.000, 75,2%; §19.9), termasuk bagian yang belum dikonfirmasi; warnanya tidak berubah dekat atau lewat batas (§19.9; §20 D-7). Angka "sisa" selalu tertulis, jadi batang hanya penegas.

**Pembaca layar.** Angka utama dibacakan "Uang kita, dua puluh empat juta tujuh ratus tiga puluh lima ribu rupiah" (§15 "Pembaca layar"). Lencana Beranda dibacakan "2 menunggu", lencana Chat "2 belum dibaca" (§25.4 "Navigasi"). Batang dibacakan lewat kata di sampingnya saja.

### 1.2 Sedang (iPad mini 7 potret, 600–839 px)

- Rail 80 px di kiri, berlatar `navigasi`: tombol "Catat" 56 × 56 px di kepala rail (ikon tambah dan label di bawahnya), lalu Beranda, Uang, Chat, Saya dengan ikon dan label di bawahnya (§25.2 baris Sedang).
- **Satu kolom isi**, paling lebar 640 px, dipusatkan di ruang sisa rail, margin samping paling sedikit 24 px, padding kartu dan kelompok daftar 20 px; urutan sama dengan Ringkas, baris 1 sampai 28 (§6.3; §17.8). Bukan dua kolom.
- Judul "Rumah Kenanga" tetap di bilah atas isi.
- Tanpa tombol Catat mengambang (sudah di rail).

### 1.3 Lebar dan Sangat lebar (iPad mini 7 lanskap, laptop; ≥ 840 px)

- **Lebar (840–1199 px):** rail sama dengan Sedang. Dua panel: **panel utama** baris 1 sampai 13 (bagian 0 sampai 5), **panel pendamping** baris 14 sampai 28 (bagian 6 sampai 8), selebar sepertiga lebar isi, 320 sampai 400 px, jarak antarpanel 24 px (§6.3; §17.8).
- **Sangat lebar (≥ 1200 px):** sidebar 240 px berlabel dengan "Rumah Kenanga" di kepala dan tombol "Catat" di bawahnya (selebar sidebar dikurangi padding 16, tinggi 48); isi berhenti melebar pada 1200 px dan berada di tengah; panel pendamping paling sempit 336 px (§25.2; §17.8).
- **Padding di panel pendamping 16 px di setiap keadaan**, seperti kolom Ringkas; di panel utama 20 px (Lebar) atau 24 px (Sangat lebar) (§17.8 "Panel samping adalah kolom selebar ponsel"). Dengan ini baris 17 muat di 1133 dan 1280 tanpa bungkus (pemeriksaan 23).
- Tombol bergaris "Chat" (`chat_bubble`) di bilah atas.
- **Satu varian dengan panel Chat terbuka**: `aside` berisi utas Chat §3 tanpa bilah navigasi sendiri. Di Lebar ia menggantikan panel pendamping selama terbuka; di Sangat lebar ia menjadi panel ketiga bila lebar isi 1200 px (jendela ≥ 1504 px), dengan panel utama paling sempit 480 px dan dua panel samping 336 px (§6.3; §10.6; §25.3 "Beranda"). Di halaman ditulis sebagai keadaan sesudah pengguna menekan "Chat"; panel tidak pernah terbuka sendiri.

### 1.4 Informasi peka (Beranda)

| Hal | Di aplikasi | Di contoh statis |
|---|---|---|
| Samaran angka | Berlaku untuk setiap jumlah saat fiturnya dibangun (§12.1, §12.2 baris Beranda) | Tidak tampil (Tier 3 di web) |
| Halaman tersembunyi | Penutup netral `kertas` dengan simpul (§25.9) | Tidak dibangun: satu-satunya skrip contoh adalah pengalih arah, dan isinya hanya data rekaan |
| Judul halaman | "Casa · Beranda", tanpa nama rumah tangga atau angka | Sama |
| Tangkapan layar | Tidak dicegah (SEC-20) | Sama |

---

## 2. Catat (LY-32, jenis Belanja)

Judul halaman saat lembar terbuka: **"Casa · Catat"** (§16.2 "Judul jendela"; §25.9).

Lembar terbuka di atas Beranda contoh §1 (bagian di belakangnya diredupkan oleh selubung; lihat catatan "Selubung" di bawah). Ringkas: menempel di bawah, selebar layar, tinggi paling banyak 90% tinggi viewport dinamis, sudut atas `--sudut-lembar` arah (24, 28, atau 16). Sedang dan Lebar: di tengah, lebar paling banyak 560 px (§25.3 baris "Lembar"; §19.12.7).

| # | Urutan §9.1 | Unsur | Isi persis | Tampil sebagai | Rujukan |
|---|---|---|---|---|---|
| 1 | — | Kepala lembar, kiri | Tombol ikon `close`, label pembaca layar "Tutup" (menyimpan draf) | Ikon `tinta-2` | §9.3 "Draf"; §25.4 "Lembar" |
| 2 | — | Kepala lembar, tengah | `select` bawaan bertuliskan "Belanja"; pilihannya: Belanja, Uang masuk, Kasih ke, Terima dari, Pindah, Minta anggaran, Jalan-jalan | Pemilih bawaan; huruf arah | §9.1; §25.4 "Pilihan jenis catat" |
| 3 | 7 | Kepala lembar, kanan | Tombol "Simpan", wadah `aksen`, teks `di-atas-aksen`, tinggi ≥ 48 px. Satu-satunya tombol berisi aksen selama lembar terbuka | `aksen` | §25.3 "Papan ketik virtual"; §17.8 aturan 4 |
| 4 | 1 | Label isian | "Jumlah" (label tampak) | Teks isi | §25.5; §25.11 pemeriksaan 18 |
| 5 | 1 | Isian jumlah | "Rp" kecil di depan, lalu **87.000**; cincin fokus karena isian ini langsung difokus | `angka-besar`; "Rp" 60% `tinta-2`; batas `garis-tegas`, cincin fokus `aksen-tinta` | §9.1 urutan 1; §19.4; §19.12.5 |
| 6 | 1 | Tombol "000" | Tombol bergaris di **samping kanan** isian jumlah, tulisan "000", nama aksesibel **"000, tambah tiga nol"** (teks tampak lalu teks tak tampak ", tambah tiga nol", tanpa `aria-label`; WCAG 2.5.3; keputusan B 13 November 2026), sama tinggi dengan isian | Bergaris `garis-tegas`, teks `aksen-tinta` | §25.3 "Papan ketik virtual"; §16.2 "Papan angka di iPad" |
| 7 | 2 | Label kelompok | "Tempat uang" | Teks isi | §9.1 urutan 2 |
| 8 | 2 | Pilihan 1, **terpilih** (terakhir dipakai) | Glif **tunai** · "Dompet Andi" | Wadah `aksen-muda`, tanda pilih; kata `tinta` | §9.1 urutan 2; §19.5; §19.12.2 `aksen-muda` "butir terpilih" |
| 9 | 2 | Pilihan 2 | Glif **rekening** · "Rekening Andi" | Tanpa wadah, batas `garis` | §19.5 |
| 10 | 2 | Pilihan 3 | Glif **rekening** · "Rekening bersama" | Sama | §19.5 |
| 11 | 3 | Kategori | Satu baris tombol: "Kategori" kecil `tinta-2` di atas "Dapur" · "sisa" Rp620.000 rata kanan · ikon `chevron_right` | Jumlah `angka-baris`; angka sisa tidak berubah saat jumlah diketik | §9.1 urutan 3 |
| 12 | 4 | Tanggal | Label "Tanggal"; tiga tombol: "Hari ini" (**terpilih**, `aksen-muda`), "Kemarin", "Pilih tanggal" | Tombol tinggi ≥ 48 px | §9.1 urutan 4; §25.4 "Tanggal" |
| 13 | 5 | Rincian terlipat | `details` tertutup: "Tambah rincian" | Teks `aksen-tinta` | §9.1 urutan 5; §25.4 "Rincian terlipat" |
| 14 | 6 | Saran kontekstual | **Tidak ada** di contoh ini (jumlah di bawah ambang, tidak ada persetujuan terbuka di Dapur) | — | §9.4 |
| 15 | 7 | Simpan di akhir isi | Tombol "Simpan" selebar lembar, perintah yang sama dengan unsur 3 | Tonal: wadah `aksen-muda`, teks `aksen-tinta` (supaya hanya satu tombol berisi aksen) | §25.3 "Papan ketik virtual"; §17.8 aturan 4 |

**Isian jumlah.** `inputmode="numeric"`, tanpa isian otomatis (`autocomplete="off"`; SEC-175 (d)), ukuran huruf isian tidak di bawah 1 rem (§25.4 "Isian"). Di contoh statis angka 87.000 sudah tertulis sebagai nilai isian.

**Selubung.** Latar redup di belakang lembar memakai token `--selubung` (terang: `tinta` 40%; gelap: hitam 60%), ditetapkan di §19.12.3 dan §25.4.

**Informasi peka (Catat).** Isian jumlah tetap terlihat saat angka disamarkan, sisa anggaran tersamar (§12.1 "Mencatat saat tersamar"; fitur Tier 3 di web, tidak tampil di contoh). Judul halaman tanpa jenis atau jumlah selain "Catat".

---

## 3. Chat (LY-60)

Judul halaman: **"Casa · Chat"** (§25.9).

### 3.1 Kepala

| Unsur | Isi persis | Tampil sebagai | Rujukan |
|---|---|---|---|
| `h1` | "Sari" | `judul-layar` | §10.1 "Kepala"; §19.4 |
| Baris kehadiran | "online" | Teks kecil `tinta-2` | §10.3 |

Tujuan Chat tidak punya tombol kembali (tujuan tingkat atas). Bilah navigasi bawah sama dengan Beranda, dengan Chat terpilih dan tanpa lencana; Beranda tetap berlencana "2". Tidak ada tombol Catat mengambang di Chat; Catat lewat tombol tambah di kotak tulis (§25.2 baris Ringkas).

### 3.2 Utas, dari atas

Lebar gelembung paling banyak 80% lebar utas dan tidak lebih dari 480 px, juga di Sedang dan Lebar (§16.2 baris Chat). Gelembung sudut 12 px. Waktu kecil `tinta-2` di dalam gelembung. Tanpa foto profil atau avatar (§10.1).

| # | Jenis | Sisi | Isi persis | Tampil sebagai | Rujukan |
|---|---|---|---|---|---|
| 1 | Pemisah hari | Tengah | "Kemarin" | Teks kecil `tinta-2` | §10.1 "Utas" |
| 2 | Baris peristiwa | Tengah | Glif **simpul** · "Disepakati: Jatah pribadi Sari, 16.10" | Teks kecil `tinta-2`, tanpa gelembung; glif `tinta-2` | §10.4 "Baris Disepakati"; §18.1; §19.5 baris Simpul |
| 3 | Pesan Sari, pertama dalam kelompok | Awal (kiri) | Nama "Sari" kecil di atasnya; "Kulkasnya makin berisik, tadi pagi bocor lagi." · 19.32 | Wadah `kartu`, teks `tinta`; Arah 3: wadah `sisip` tanpa garis | §10.1 "Gelembung"; §19.12.2; §19.12.5 Arah 3 |
| 4 | Pesan Sari, dalam kelompok yang sama | Awal | "Aku sudah lihat yang dua pintu di toko dekat pasar." · 19.33 | Sama, tanpa nama | §10.1 (kelompok 5 menit) |
| 5 | **Kartu permintaan** | Selebar utas | Lihat §3.3 kartu A | — | §10.4 |
| 6 | Pesan sendiri | Akhir (kanan) | "Aku lihat dulu, besok pagi kita bahas ya." · 21.05 · tanda dua centang dengan kata "Dibaca" bagi pembaca layar | Wadah `aksen-muda`, teks `tinta`; ikon `done_all` `aksen-tinta` | §10.1; §10.2 baris "Dibaca"; §19.12.2 `aksen-muda` |
| 7 | Pemisah hari | Tengah | "Hari ini" | Sama dengan #1 | §10.1 |
| 8 | Pesan Sari | Awal | Nama "Sari"; "Uang belanja minggu ini aku taruh di laci dapur." · 08.02 | Sama dengan #3 | §10.1 |
| 9 | **Kartu serah terima** | Selebar utas | Lihat §3.3 kartu B | — | §10.4; §19.8 butir 5 |

Tidak ada jumlah uang di dalam teks pesan bebas, supaya setiap jumlah di contoh memakai token angka (pemeriksaan 13).

### 3.3 Kartu objek

Kedua kartu: lapisan isi, wadah `kartu` dengan batas 1 px `garis`, sudut `--sudut-kartu` arah, selebar kolom utas dan paling lebar 480 px, tidak berpihak ke sisi mana pun. Batas, lebar, dan letak itulah yang membedakan kartu dari gelembung (§10.4). Susunan menurut §10.4 butir (1) sampai (6). Tombol sama tinggi, ≥ 48 px; pada lebar sempit atau teks besar baris tombol menjadi tumpukan tegak selebar kartu dengan tombol utama di atas (§10.4; §15). Nama aksesibel setiap tombol menyebut objeknya, misalnya "Setuju, minta anggaran Kulkas baru" (§25.11 pemeriksaan 18).

**Kartu A: permintaan menunggu, pelihat sebagai pemberi tanggapan**

| Butir §10.4 | Isi persis | Tampil sebagai |
|---|---|---|
| (1) Jenis | "Minta anggaran" | Teks kecil `tinta-2` |
| (2) Judul | "Kulkas baru" | `judul-bagian` |
| (3) Jumlah | Rp7.850.000 | `angka-besar`, "Rp" kecil `tinta-2` |
| (4) Keadaan | "Menunggu kamu sejak Kam 19.40" | Teks kecil `tinta-2` |
| (5) Tombol | "Setuju" (tonal: wadah `aksen-muda`, teks `aksen-tinta`) · "Usul lain" (bergaris `garis-tegas`, teks `aksen-tinta`) · "Belum sepakat" (bergaris, teks `aksen-tinta`, tidak berwarna peringatan) | §10.4 tabel "Tombol per kartu" baris "Permintaan menunggu"; §10.4 "Bentuk tombol"; §25.4 "Tombol utama dan tombol kartu" |
| (6) | "Lihat rincian" | Tautan `aksen-tinta`, bergaris bawah |

**Kartu B: serah terima, pelihat sebagai penerima** (motif angka bersama pemegangnya)

| Butir §10.4 | Isi persis | Tampil sebagai |
|---|---|---|
| (1) Jenis | "Serah terima" | Teks kecil `tinta-2` |
| (2) Judul | "Sari kasih ke kamu" | `judul-bagian` |
| (3) Jumlah | Rp200.000, rata kiri | `angka-besar` |
| Motif | Garis tipis `garis`, lalu dua baris pemegang: glif **tunai** · "Dari Dompet Sari"; glif **tunai** · "Ke Dompet Andi". Jumlah tidak diulang di kedua baris karena sama dengan butir (3) | §19.8 butir 5 ("kartu serah terima: glif 'dari' dan 'ke'") |
| (4) Keadaan | "Menunggu kamu sejak 08.03" | Teks kecil `tinta-2` |
| (5) Tombol | "Sudah terima" (tonal: `aksen-muda`, teks `aksen-tinta`) · "Tidak sesuai" (bergaris). **Dua tombol sama tinggi** | §10.4 baris "Serah terima"; §9.7 |
| (6) | "Lihat rincian" | Tautan `aksen-tinta`, bergaris bawah |

Kartu A punya tiga tombol, bukan dua: §10.4 menetapkan tiga tombol untuk permintaan yang menunggu pemberi tanggapan. Kartu dengan dua tombol sama tinggi yang diminta rencana tugas dipenuhi kartu B.

### 3.4 Kotak tulis

| Unsur | Isi persis | Rujukan |
|---|---|---|
| Tombol tambah | Ikon `add`, label pembaca layar "Tambah" (Foto, Kamera, Catat belanja, Minta anggaran, Kirim catatan yang sudah ada; menu tidak dibuka di contoh) | §9.1 baris "Dari Chat" |
| Isian | `textarea` kosong dengan label tampak atau label pembaca layar "Pesan untuk Sari" dan teks contoh "Tulis pesan" | §10.1 "Kotak tulis"; §25.11 pemeriksaan 18 |
| Tombol kirim | Ikon `send` `aksen-tinta`, label pembaca layar "Kirim" | §10.1 |

### 3.5 Per lebar

- Ringkas: layar penuh, kotak tulis menempel di atas bilah navigasi (§16.1 baris Chat; §25.2).
- Sedang, Lebar, dan Sangat lebar: layar penuh di tujuan Chat dengan rail atau sidebar; utas satu kolom paling lebar 640 px di tengah; lebar gelembung dan kartu tidak ikut melebar (§16.2 baris Chat; §25.3 "Lebar baca").
- Panel Chat di Beranda Lebar dan Sangat lebar: isi yang sama, tanpa rail dan bilah navigasi sendiri, padding 16 px seperti panel samping lain (§1.3 di atas; §17.8).

### 3.6 Informasi peka (Chat)

Jumlah di kartu tersamar bila "Sembunyikan angka" menyala; teks bebas tidak (§12.2 baris LY-60; fitur Tier 3 di web, tidak tampil). Judul halaman "Casa · Chat" tanpa nama Sari (§25.9). Penutup netral seperti Beranda, tidak dibangun di contoh statis.

---

## 4. Ikon dan nama di layar utama

Satu set ikon per arah; manifest contoh per arah (`manifest-{arah}.webmanifest`) menunjuk set arahnya, sehingga ikon di layar utama iPhone 13 dan peluncur Galaxy A55 ikut dinilai (§19.12.11 butir 5). Hanya varian terang: varian gelap §19.12.9 adalah **[Native]**, karena manifest web dan `apple-touch-icon` tidak punya tempat untuk ikon gelap.

| Hal | Isi | Rujukan |
|---|---|---|
| Ikon `any` | `web/static/icons/casa-icon-limau.svg` (simpul limau `#A3EB6E` di atas `#141414`), `casa-icon-koral.svg` (simpul `#141414` di atas koral `#FF7F62`), `casa-icon-biru.svg` (simpul putih di atas biru `#2443C4`); satu bentuk besar. PNG per arah: `web/static/icons/{limau,koral,biru}-192.png` dan `-512.png`, ikon layar utama Apple `{limau,koral,biru}-180.png` | §19.5 baris "Klien web"; §19.12.9 |
| Ikon `maskable` | `web/static/icons/casa-icon-{limau,koral,biru}-maskable.svg` dan `{limau,koral,biru}-maskable-512.png`: gambar yang sama, simpul diperkecil sehingga seluruhnya di dalam lingkaran berjari-jari 295 px dari pusat kanvas 1024 (zona aman 80% berjari-jari 410 px) | §19.5 |
| Nama | "Casa" di manifest (`name` dan `short_name`) dan di bawah ikon, tanpa data | §19.5; SEC-175 (g); ADR-030 butir 9 |
| `theme-color` | `kertas` per tema dan per arah (Arah 3 terang `#FFFFFF`) | §25.4 "Token"; §19.12.3 |

---

## 5. Pemeriksaan sendiri 11 sampai 15 (§19.8) atas isi ini

| # | Pemeriksaan | Hasil atas isi | Yang harus dijaga `web-dev` |
|---|---|---|---|
| 11 | Glif dan ikon sendiri | Lulus: tempat uang selalu memakai glif tunai, rekening, atau e-wallet dari `web/static/glyphs/`; ikon aplikasi adalah simpul di ketiga arah | Jangan memakai simbol Material `payments`, `account_balance`, `wallet`, atau `savings` di tempat glif |
| 12 | Tanpa hiasan | Lulus: tidak ada ilustrasi, gradasi, karakter, foto, spanduk, atau pola. Arsiran dan garis putus-putus batang adalah penanda bagian belum dikonfirmasi, bukan hiasan | Tanpa bayangan selain lembar dan tombol Catat mengambang (§19.12.7) |
| 13 | Angka bertoken | Lulus: setiap jumlah punya token di tabel; tidak ada jumlah di dalam kalimat biasa | `tabular-nums` pada ketiga token; "Rp" sebagai unsur tersendiri |
| 14 | Motif hadir | Lulus: Beranda bagian 2 dan 3 (baris 6 sampai 11) dan kartu serah terima (kartu B) | Urutan label, angka rata kiri, garis tipis, baris pemegang tidak diubah per lebar |
| 15 | Kata templat | Lulus: tidak ada frasa dari daftar §19.8 butir 3, tanpa tanda seru dan emoji; baris keterangan dengan koma (§18.1) | Teks tombol, `aria-label`, dan judul halaman juga tunduk |

---

## 6. Yang sengaja tidak ada

Tombol "Sembunyikan angka" dan samaran (Tier 3 di web); kata penyemangat (Tier 2); saran kata di chat (Tier 2); penutup netral dan keadaan luring (keadaan browser diuji di layar 1.10, §25.11 pemeriksaan 20); warna makna `positif`, `peringatan`, `lebih` (tidak ada keadaan yang memerlukannya di data rekaan ini); grafik, tren, perbandingan antar-orang (§6.1).
