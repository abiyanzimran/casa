# Isi contoh layar gerbang 1.8: Beranda, Catat, Chat

Pemilik: `ux-designer`. Dibuat 13 November 2026 untuk FS-1 langkah 1, tugas 1.8 (gerbang contoh layar, 08 §19.8 dan §25.11). Dipakai `web-dev` untuk membangun halaman statis di `design-samples/`. Rujukan "§" tanpa nama dokumen berarti `docs/08-ux-ui-strategy.md`.

Berkas ini hanya isi: bagian, baris, angka, nama, kalimat, token yang dipakai, dan glif. Tata letak per lebar mengikuti §6.3, §16.1, §25.2, dan §25.3; bentuk komponen mengikuti §25.4.

---

## 0. Aturan untuk ketiga contoh

| Hal | Aturan | Rujukan |
|---|---|---|
| Data | **Rekaan seluruhnya.** Tidak ada nama, angka, tempat uang, atau pesan dari rumah tangga pilot | §19.8 "Dipakai ulang untuk wawancara" |
| Orang | Dua anggota dewasa: **Sari** dan **Andi**, keduanya nama contoh yang sudah dipakai 08 (§6.1, §10.4, §9.5). Pelihat di ketiga contoh adalah **Andi**, jadi Andi disebut "kamu". Lead memastikan kedua nama bukan nama anggota rumah tangga pilot sebelum contoh ditunjukkan | Brief tugas; §18.1 "Sapaan" |
| Rumah tangga | **"Rumah Kenanga"** (rekaan; lead memeriksanya bersama kedua nama) | §6.1 bagian 0 |
| Tempat uang | Tanpa nama bank atau layanan e-wallet: "Rekening Andi", "Rekening bersama", "E-wallet Sari", "Dompet Andi", "Dompet Sari". Bagian `{bank}` dan `{layanan}` di §18.2 sengaja dikosongkan supaya contoh yang juga dipakai untuk wawancara 13 §6.3 tidak memuat merek | §18.2 baris "Tempat uang" |
| Hari | Hari ini Jumat, 13 November; kemarin Kamis. Bulan anggaran 1–30 November | §6.1 bagian 2 |
| Penulisan uang | "Rp" lalu angka dengan titik ribuan, tanpa spasi, tanpa desimal, tanpa "rb" atau "jt". "Rp" ditulis sebagai unsur tersendiri dari token, bukan bagian teks angka | §17.3, §19.4 |
| Token angka | Setiap jumlah memakai satu token §19.4 (`angka-utama`, `angka-besar`, atau `angka-baris`) dengan `tabular-nums`. Tidak ada jumlah di dalam kalimat teks biasa: bila sebuah baris berisi kata dan jumlah, kata di kiri dan jumlah di kanan sebagai unsur sendiri (pemeriksaan 13) | §19.4, §19.8 butir 4 |
| "Rp" kecil | Pada `angka-utama` dan `angka-besar`: 60% ukuran angka, garis dasar sama, warna `tinta-2`. Pada `angka-baris`: ukuran dan warna sama dengan angkanya | §19.4 "Aturan angka" |
| Warna jumlah | Semua jumlah `tinta`, tidak pernah `teal`, hijau, atau merah | §19.9 |
| Glif | Dari `web/static/glyphs/`, warna `tinta-2`, selalu bersama kata. Glif hiasan bagi pembaca layar (`aria-hidden`), karena katanya sudah ada | §19.5; WCAG 1.4.1 |
| Ikon antarmuka | Material Symbols gaya garis sebagai SVG milik sendiri; nama simbol di bawah adalah usulan, `web-dev` boleh memilih padanan yang maknanya sama | §17.5, §25.4 "Ikon" |
| Warna | Hanya lewat token §19.3 dan §25.4; tidak ada warna yang ditulis langsung di halaman | §25.4 "Token" |
| Kata | Tanpa kata dari daftar kata templat §19.8 butir 3, tanpa kata §18.1 "Kata yang tidak dipakai", tanpa kosakata akuntansi §18.2, tanpa "Anda", tanpa tanda seru, tanpa emoji | §18, §19.8 |

### 0.1 Angka yang dipakai bersama

Semua angka ketiga contoh berasal dari keadaan rekaan yang sama, supaya Beranda dan Chat saling cocok.

| Hal | Angka | Hubungan |
|---|---|---|
| Uang kita | Rp88.888.888 | Jumlah ketiga baris pemegang (06 INV-A24; 07 §5.1) |
| Dipegang kamu (Dompet Andi, Rekening Andi) | Rp10.488.888 | Termasuk Rp200.000 yang menunggu kamu mengonfirmasi |
| Dipegang Sari (Dompet Sari, E-wallet Sari) | Rp2.150.000 | Termasuk Rp50.000 yang sedang dibahas |
| Dipegang bersama (Rekening bersama) | Rp76.250.000 | — |
| Pemeriksaan | 10.488.888 + 2.150.000 + 76.250.000 = 88.888.888 | Lead memeriksa dengan skrip (CLAUDE.md "Computed, not read") |
| Disiapkan untuk tagihan | Rp3.200.000 | Alokasi; tidak dikurangkan dari Uang kita (PRN-17) |
| Disisihkan untuk tabungan | Rp60.000.000 | Alokasi |
| Belum disiapkan untuk apa pun | Rp25.688.888 | 88.888.888 − 3.200.000 − 60.000.000 (07 INV-17) |
| Uang pribadi saya (Andi) | Rp1.725.000 | Tidak masuk Uang kita (07 §5.2) |
| Serah terima Sari ke Andi | Rp200.000 | Dompet Sari ke Dompet Andi; menunggu Andi |
| Minta anggaran Sari | Rp7.850.000, "Kulkas baru" | Menunggu Andi sejak Kam 19.40 |
| Minta anggaran Andi | Rp650.000, "Servis motor" | Menunggu Sari sejak Rab 08.15 |
| Anggaran Dapur | Batas Rp2.500.000, terpakai Rp1.880.000, sisa Rp620.000 | Batang 75,2% (1.880.000 / 2.500.000) |
| Anggaran Transportasi | Batas Rp800.000, terpakai Rp460.000, sisa Rp340.000 | Batang 57,5% |

**Jumlah terpanjang:** Rp88.888.888 pada `angka-utama` (Beranda) untuk menguji lebar pada 320 CSS px dan teks 200% (§25.11 pemeriksaan 17). Angka 8 adalah angka terlebar bila fitur tabular gagal, sehingga angka ini juga memperlihatkan kegagalan `tnum` di Safari (fokus tinjauan 3). Uji pasangan "Rp1.111.111" dan "Rp8.888.888" (§19.4) dijalankan `web-dev` di halaman uji sendiri, bukan di contoh yang ditunjukkan.

---

## 1. Beranda (LY-20)

Judul halaman: **"Casa · Beranda"** (§25.9; SEC-175 (a)).

### 1.1 Ringkas (iPhone 13, Galaxy A55, jendela < 600 px)

Urutan dari atas. Kolom "Tampil sebagai" menyebut token; huruf isi lain adalah huruf sistem (§25.4 "Huruf").

| # | Bagian §6.1 | Unsur | Isi persis | Tampil sebagai | Rujukan |
|---|---|---|---|---|---|
| 1 | 0 Judul | Bilah atas, `h1` | "Rumah Kenanga" | Huruf sistem, tebal; latar `kertas` | §6.1 bagian 0; §25.4 "Bilah atas" |
| 2 | 0 Judul | Tombol "Sembunyikan angka" | **Tidak tampil** di contoh: fiturnya Tier 3 (F-FIN-027) | — | §12.1 baris "Tier" |
| 3 | 1 Kata penyemangat | — | **Tidak tampil**: Tier 2 | — | §6.1 bagian 1 |
| 4 | 2 Uang kita | Wadah bagian | Wadah `kertas-2`, sudut 12 px, memuat unsur 5 sampai 11 | `kertas-2` | §17.7; §19.3 |
| 5 | 2 Uang kita | Label | "Uang kita" | `judul-bagian`, `tinta` | §19.4; §18.2 |
| 6 | 2 Uang kita | Angka utama | Rp88.888.888, rata kiri | `angka-utama`; "Rp" kecil `tinta-2` | §19.2 butir 1; §19.4; §19.8 butir 5 |
| 7 | 2 Uang kita | Keterangan | "Bulan ini, 1–30 November" | Teks kecil `tinta-2` | §6.1 bagian 2 |
| 8 | 3 Pemegang | Garis tipis | Satu garis `garis` selebar wadah | `garis` | §19.8 butir 5 |
| 9 | 3 Pemegang | Baris pemegang 1 | Glif **tunai** dan **rekening** · "Dipegang kamu" · Rp10.488.888 rata kanan | Glif `tinta-2`; kata `tinta`; jumlah `angka-baris` | §6.1 bagian 3; §19.8 butir 5 |
| 9a | 3 Pemegang | Bagian belum dikonfirmasi, di bawah baris 9 | Wadah `pasir`, sudut 8 px: "Menunggu kamu mengonfirmasi" di kiri, Rp200.000 di kanan | `pasir`; kata `tinta`; jumlah `angka-baris` | §6.1 bagian 3; §9.6 baris "Belum dikonfirmasi"; §19.8 butir 5 |
| 10 | 3 Pemegang | Baris pemegang 2 | Glif **tunai** dan **e-wallet** · "Dipegang Sari" · Rp2.150.000 rata kanan | Sama dengan baris 9 | §6.1 bagian 3 |
| 10a | 3 Pemegang | Bagian dibahas, di bawah baris 10 | Wadah `pasir` dengan ikon info (`info`): "Sedang dibahas" di kiri, Rp50.000 di kanan | Sama dengan 9a | §9.6 baris "Keberatan"; §19.9 |
| 11 | 3 Pemegang | Baris pemegang 3 | Glif **rekening** · "Dipegang bersama" · Rp76.250.000 rata kanan | Sama dengan baris 9 | §6.1 bagian 3 ("tempat uang bersama") |
| 12 | 4 Penjelas angka | Tiga baris, di bawah wadah Uang kita | "Disiapkan untuk tagihan" Rp3.200.000 · "Disisihkan untuk tabungan" Rp60.000.000 · "Belum disiapkan untuk apa pun" Rp25.688.888 | Kata `tinta-2`; jumlah `angka-baris` `tinta`, rata kanan | §6.1 bagian 4; §18.2 |
| 13 | 5 Uang pribadi saya | Garis pemisah, lalu satu baris | "Uang pribadi saya" · Rp1.725.000 rata kanan; di bawahnya "Hanya kamu yang melihat ini." | Garis `garis`; jumlah `angka-baris`; keterangan `tinta-2` | §6.1 bagian 5; §18.3 |
| 14 | 6 Menunggu | Judul | "Menunggu" | `judul-bagian` | §6.1 bagian 6; §19.4 |
| 15 | 6 Menunggu | Subjudul | "Menunggu kamu" | Teks isi tebal | §6.1 bagian 6 |
| 16 | 6 Menunggu | Butir 1 (tindakan ringan) | Glif **tunai** · "Terima dari Sari" · Rp200.000; tombol bergaris "Sudah terima" di bawah atau di kanan, tinggi ≥ 48 px | Jumlah `angka-baris`; tombol bergaris `garis-tegas`, teks `teal` | §6.1 bagian 6; §9.2 baris Beranda; §15 "Sasaran sentuh" |
| 17 | 6 Menunggu | Butir 2 (membuka rincian) | "Sari minta anggaran" · di bawahnya "Kulkas baru" `tinta-2` · Rp7.850.000 · ikon `chevron_right`. Seluruh baris satu tautan | Jumlah `angka-baris` | §6.1 bagian 6; §25.4 "Daftar dan bagian" |
| 18 | 6 Menunggu | Subjudul | "Menunggu Sari" | Teks isi tebal | §6.1 bagian 6 |
| 19 | 6 Menunggu | Butir 3 (tanpa tombol) | "Kamu minta anggaran" · di bawahnya "Servis motor · sejak Rab 08.15" `tinta-2` · Rp650.000 | Jumlah `angka-baris` | §6.1 bagian 6 ("tanpa tombol") |
| 20 | 6 Menunggu | Tautan | "Lihat semua" | Teks `teal` | §6.1 bagian 6 |
| 21 | 7 Minggu ini | Judul | "Minggu ini" | `judul-bagian` | §19.4 |
| 22 | 7 Minggu ini | Baris tagihan | "Listrik" · di bawahnya "Sen, 16 Nov" · Rp850.000 | Jumlah `angka-baris` | §6.1 bagian 7 |
| 23 | 7 Minggu ini | Baris cicilan | "Cicilan motor" · di bawahnya "Rab, 18 Nov" · Rp1.275.000 | Jumlah `angka-baris` | §6.1 bagian 7 |
| 24 | 7 Minggu ini | Baris persetujuan terbuka | "Makan malam" · di bawahnya "sisa" · Rp25.000 · tombol bergaris "Selesaikan" | Jumlah `angka-baris` | §6.1 bagian 7 (kata contohnya dipecah menjadi kata dan jumlah, lihat §0 "Token angka") |
| 25 | 8 Anggaran | Judul | "Anggaran" | `judul-bagian` | §19.4 |
| 26 | 8 Anggaran | Baris 1 | "Dapur" · "sisa" Rp620.000 rata kanan · di bawahnya batang selebar baris, isi 75,2% · "dari Rp2.500.000" `tinta-2` | Batang isi `tinta-2` di atas lintasan `kertas-3`, `aria-hidden`; jumlah `angka-baris` | §6.1 bagian 8; §19.9; §25.4 "Batang" |
| 27 | 8 Anggaran | Baris 2 | "Transportasi" · "sisa" Rp340.000 · batang isi 57,5% · "dari Rp800.000" | Sama | Sama |
| 28 | 8 Anggaran | Tautan | "Semua anggaran" | Teks `teal` | §6.1 bagian 8 |
| 29 | — | Tombol Catat mengambang | "Catat" dengan ikon `add`, kanan bawah di atas bilah navigasi; wadah `teal`, isi `di-atas-teal` | `teal`, `di-atas-teal` | §25.2 baris Ringkas |
| 30 | — | Bilah navigasi bawah | Beranda (`home`, terpilih, lencana "2"), Uang (`account_balance_wallet`), Chat (`chat_bubble`, lencana "2"), Saya (`person`). Label kata selalu tampil. Tujuan terpilih: ikon dan kata `teal`, `aria-current="page"` | `teal` hanya pada yang terpilih | §4.1; §25.2; §25.4 "Navigasi" |

**Batang anggaran.** Isinya adalah bagian yang **terpakai** (1.880.000 dari 2.500.000, 75,2%; §19.9), dan warnanya tidak berubah dekat atau lewat batas (§19.9; §20 D-7). Angka "sisa" selalu tertulis, jadi batang hanya penegas.

**Pembaca layar.** Angka utama dibacakan "Uang kita, delapan puluh delapan juta delapan ratus delapan puluh delapan ribu delapan ratus delapan puluh delapan rupiah" (§15 "Pembaca layar"). Lencana Beranda dibacakan "2 menunggu", lencana Chat "2 belum dibaca" (§25.4 "Navigasi"). Batang dibacakan lewat kata di sampingnya saja.

### 1.2 Sedang (iPad mini 7 potret, 600–839 px)

- Rail di kiri: kepala rail "Rumah Kenanga", lalu tombol "Catat" di puncak, lalu Beranda, Uang, Chat, Saya dengan ikon dan label di bawahnya (§25.2 baris Sedang).
- Isi dua kolom grid: **kiri** baris 4 sampai 13 (bagian 2 sampai 5); **kanan** baris 14 sampai 28 (bagian 6 sampai 8) (§6.3).
- Tanpa tombol Catat mengambang (sudah di rail).

### 1.3 Lebar (iPad mini 7 lanskap, laptop; ≥ 840 px)

- Rail sama dengan Sedang; dua kolom sama dengan Sedang; kolom teks paling lebar sekitar 680 px, di layar sangat lebar isi berhenti melebar dan berada di tengah (§25.3 "Lebar baca").
- Tombol ikon "Chat" (`chat_bubble`, label pembaca layar "Buka chat") di bilah atas.
- **Satu varian dengan panel Chat terbuka**: `aside` di kanan, sekitar 30% lebar dan tidak lebih sempit dari 320 px, berisi utas Chat §3 di bawah tanpa bilah navigasi sendiri (§6.3; §10.6; §25.3 "Beranda"). Varian ini menguji kartu di wadah sempit lewat container query. Di halaman ditulis sebagai keadaan sesudah pengguna menekan "Chat"; panel tidak pernah terbuka sendiri.

### 1.4 Informasi peka (Beranda)

| Hal | Di aplikasi | Di contoh statis |
|---|---|---|
| Samaran angka | Berlaku untuk setiap jumlah saat fiturnya dibangun (§12.1, §12.2 baris Beranda) | Tidak tampil (Tier 3) |
| Halaman tersembunyi | Penutup netral `kertas` dengan simpul (§25.9) | Tidak dibangun: contoh tanpa JavaScript dan hanya berisi data rekaan |
| Judul halaman | "Casa · Beranda", tanpa nama rumah tangga atau angka | Sama |
| Tangkapan layar | Tidak dicegah (SEC-20) | Sama |

---

## 2. Catat (LY-32, jenis Belanja)

Judul halaman saat lembar terbuka: **"Casa · Catat"** (§16.2 "Judul jendela"; §25.9).

Lembar terbuka di atas Beranda contoh §1 (bagian di belakangnya diredupkan oleh selubung; lihat catatan "Selubung" di bawah). Ringkas: menempel di bawah, selebar layar, tinggi paling banyak 90% tinggi viewport dinamis, sudut atas 16 px. Sedang dan Lebar: di tengah, lebar paling banyak 560 px (§25.3 baris "Lembar"; §25.4 "Jarak").

| # | Urutan §9.1 | Unsur | Isi persis | Tampil sebagai | Rujukan |
|---|---|---|---|---|---|
| 1 | — | Kepala lembar, kiri | Tombol ikon `close`, label pembaca layar "Tutup" (menyimpan draf) | Ikon `tinta-2` | §9.3 "Draf"; §25.4 "Lembar" |
| 2 | — | Kepala lembar, tengah | `select` bawaan bertuliskan "Belanja"; pilihannya: Belanja, Uang masuk, Kasih ke, Terima dari, Pindah, Minta anggaran, Jalan-jalan | Huruf sistem | §9.1; §25.4 "Pilihan jenis catat" |
| 3 | 7 | Kepala lembar, kanan | Tombol "Simpan", wadah `teal`, teks `di-atas-teal`, tinggi ≥ 48 px | `teal` | §25.3 "Papan ketik virtual" ("Simpan" di kepala lembar) |
| 4 | 1 | Label isian | "Jumlah" (label tampak) | Teks isi | §25.5; §25.11 pemeriksaan 18 |
| 5 | 1 | Isian jumlah | "Rp" kecil di depan, lalu **87.000**; cincin fokus `teal` karena isian ini langsung difokus | `angka-besar`; "Rp" 60% `tinta-2`; batas `garis-tegas`, fokus `teal` | §9.1 urutan 1; §19.4; §19.2 butir 2 |
| 6 | 1 | Tombol "000" | Tombol bergaris di **samping kanan** isian jumlah, tulisan "000", label pembaca layar "Tambah tiga nol", sama tinggi dengan isian | Bergaris `garis-tegas`, teks `teal` | §25.3 "Papan ketik virtual"; §16.2 "Papan angka di iPad" |
| 7 | 2 | Label kelompok | "Tempat uang" | Teks isi | §9.1 urutan 2 |
| 8 | 2 | Pilihan 1, **terpilih** (terakhir dipakai) | Glif **tunai** · "Dompet Andi" | Wadah `teal-muda`, tanda pilih; kata `tinta` | §9.1 urutan 2; §19.5; §19.3 `teal-muda` "butir terpilih" |
| 9 | 2 | Pilihan 2 | Glif **rekening** · "Rekening Andi" | Tanpa wadah, batas `garis` | §19.5 |
| 10 | 2 | Pilihan 3 | Glif **rekening** · "Rekening bersama" | Sama | §19.5 |
| 11 | 3 | Kategori | Satu baris tombol: "Kategori" kecil `tinta-2` di atas "Dapur" · "sisa" Rp620.000 rata kanan · ikon `chevron_right` | Jumlah `angka-baris`; angka sisa tidak berubah saat jumlah diketik | §9.1 urutan 3 |
| 12 | 4 | Tanggal | Label "Tanggal"; tiga tombol: "Hari ini" (**terpilih**, `teal-muda`), "Kemarin", "Pilih tanggal" | Tombol tinggi ≥ 48 px | §9.1 urutan 4; §25.4 "Tanggal" |
| 13 | 5 | Rincian terlipat | `details` tertutup: "Tambah rincian" | Teks `teal` | §9.1 urutan 5; §25.4 "Rincian terlipat" |
| 14 | 6 | Saran kontekstual | **Tidak ada** di contoh ini (jumlah di bawah ambang, tidak ada persetujuan terbuka di Dapur) | — | §9.4 |
| 15 | 7 | Simpan di akhir isi | Tombol "Simpan" selebar lembar, sama dengan unsur 3 | `teal` | §25.3 "Papan ketik virtual" |

**Isian jumlah.** `inputmode="numeric"`, tanpa isian otomatis (`autocomplete="off"`; SEC-175 (d)), ukuran huruf isian tidak di bawah 1 rem (§25.4 "Isian"). Di contoh statis angka 87.000 sudah tertulis sebagai nilai isian.

**Selubung.** Latar redup di belakang lembar memakai token `--selubung` (terang: `tinta` 40%; gelap: hitam 60%), ditetapkan di §19.3 dan §25.4 pada 13 November 2026. `web-dev` menuliskannya di `tokens.css` bersama token lain.

**Informasi peka (Catat).** Isian jumlah tetap terlihat saat angka disamarkan, sisa anggaran tersamar (§12.1 "Mencatat saat tersamar"; fitur Tier 3, tidak tampil di contoh). Judul halaman tanpa jenis atau jumlah selain "Catat".

---

## 3. Chat (LY-60)

Judul halaman: **"Casa · Chat"** (§25.9).

### 3.1 Kepala

| Unsur | Isi persis | Tampil sebagai | Rujukan |
|---|---|---|---|
| `h1` | "Sari" | Huruf sistem tebal | §10.1 "Kepala" |
| Baris kehadiran | "online" | Teks kecil `tinta-2` | §10.3 |

Tujuan Chat tidak punya tombol kembali (tujuan tingkat atas). Bilah navigasi bawah sama dengan Beranda, dengan Chat terpilih dan tanpa lencana; Beranda tetap berlencana "2". Tidak ada tombol Catat mengambang di Chat; Catat lewat tombol tambah di kotak tulis (§25.2 baris Ringkas).

### 3.2 Utas, dari atas

Lebar gelembung paling banyak 80% lebar utas dan tidak lebih dari 480 px, juga di Sedang dan Lebar (§16.2 baris Chat). Gelembung sudut 12 px. Waktu kecil `tinta-2` di dalam gelembung. Tanpa foto profil atau avatar (§10.1).

| # | Jenis | Sisi | Isi persis | Tampil sebagai | Rujukan |
|---|---|---|---|---|---|
| 1 | Pemisah hari | Tengah | "Kemarin" | Teks kecil `tinta-2` | §10.1 "Utas" |
| 2 | Baris peristiwa | Tengah | Glif **simpul** · "Disepakati · Jatah pribadi Sari · 16.10" | Teks kecil `tinta-2`, tanpa gelembung; glif `tinta-2` | §10.4 "Baris peristiwa"; §19.5 baris Simpul |
| 3 | Pesan Sari, pertama dalam kelompok | Awal (kiri) | Nama "Sari" kecil di atasnya; "Kulkasnya makin berisik, tadi pagi bocor lagi." · 19.32 | Wadah `kertas-2`, teks `tinta` | §10.1 "Gelembung"; §19.3 `kertas-2` |
| 4 | Pesan Sari, dalam kelompok yang sama | Awal | "Aku sudah lihat yang dua pintu di toko dekat pasar." · 19.33 | Sama, tanpa nama | §10.1 (kelompok 5 menit) |
| 5 | **Kartu permintaan** | Selebar utas | Lihat §3.3 kartu A | — | §10.4 |
| 6 | Pesan sendiri | Akhir (kanan) | "Aku lihat dulu, besok pagi kita bahas ya." · 21.05 · tanda dua centang `teal` dengan kata "Dibaca" bagi pembaca layar | Wadah `teal-muda`, teks `tinta`; ikon `done_all` `teal` | §10.1; §10.2 baris "Dibaca"; §19.3 `teal-muda` |
| 7 | Pemisah hari | Tengah | "Hari ini" | Sama dengan #1 | §10.1 |
| 8 | Pesan Sari | Awal | Nama "Sari"; "Uang belanja minggu ini aku taruh di laci dapur." · 08.02 | Sama dengan #3 | §10.1 |
| 9 | **Kartu serah terima** | Selebar utas | Lihat §3.3 kartu B | — | §10.4; §19.8 butir 5 |

Tidak ada jumlah uang di dalam teks pesan bebas, supaya setiap jumlah di contoh memakai token angka (pemeriksaan 13).

### 3.3 Kartu objek

Kedua kartu: lapisan isi, wadah `kertas-2` dengan batas 1 px `garis` (karena gelembung Sari juga `kertas-2`), sudut 12 px, selebar kolom utas dan paling lebar 480 px, tidak berpihak ke sisi mana pun. Susunan menurut §10.4 butir (1) sampai (6). Tombol sama tinggi, ≥ 48 px; pada lebar sempit atau teks besar baris tombol menjadi tumpukan tegak selebar kartu dengan tombol utama di atas (§10.4; §15). Nama aksesibel setiap tombol menyebut objeknya, misalnya "Setuju, minta anggaran Kulkas baru" (§25.11 pemeriksaan 18).

**Kartu A: permintaan menunggu, pelihat sebagai pemberi tanggapan**

| Butir §10.4 | Isi persis | Tampil sebagai |
|---|---|---|
| (1) Jenis | "Minta anggaran" | Teks kecil `tinta-2` |
| (2) Judul | "Kulkas baru" | `judul-bagian` |
| (3) Jumlah | Rp7.850.000 | `angka-besar`, "Rp" kecil `tinta-2` |
| (4) Keadaan | "Menunggu kamu sejak Kam 19.40" | Teks kecil `tinta-2` |
| (5) Tombol | "Setuju" (wadah `teal`, teks `di-atas-teal`) · "Usul lain" (bergaris `garis-tegas`, teks `teal`) · "Belum sepakat" (bergaris, teks `teal`, tidak berwarna peringatan) | §10.4 tabel "Tombol per kartu" baris "Permintaan menunggu"; §25.4 "Tombol utama dan tombol kartu" |
| (6) | "Lihat rincian" | Tautan `teal` |

**Kartu B: serah terima, pelihat sebagai penerima** (motif angka bersama pemegangnya)

| Butir §10.4 | Isi persis | Tampil sebagai |
|---|---|---|
| (1) Jenis | "Serah terima" | Teks kecil `tinta-2` |
| (2) Judul | "Sari kasih ke kamu" | `judul-bagian` |
| (3) Jumlah | Rp200.000, rata kiri | `angka-besar` |
| Motif | Garis tipis `garis`, lalu dua baris pemegang: glif **tunai** · "Dari Dompet Sari"; glif **tunai** · "Ke Dompet Andi". Jumlah tidak diulang di kedua baris karena sama dengan butir (3) | §19.8 butir 5 ("kartu serah terima: glif 'dari' dan 'ke'") |
| (4) Keadaan | "Menunggu kamu sejak 08.03" | Teks kecil `tinta-2` |
| (5) Tombol | "Sudah terima" (wadah `teal`) · "Tidak sesuai" (bergaris). **Dua tombol sama tinggi** | §10.4 baris "Serah terima"; §9.7 |
| (6) | "Lihat rincian" | Tautan `teal` |

Kartu A punya tiga tombol, bukan dua: §10.4 menetapkan tiga tombol untuk permintaan yang menunggu pemberi tanggapan. Kartu dengan dua tombol sama tinggi yang diminta rencana tugas dipenuhi kartu B.

### 3.4 Kotak tulis

| Unsur | Isi persis | Rujukan |
|---|---|---|
| Tombol tambah | Ikon `add`, label pembaca layar "Tambah" (Foto, Kamera, Catat belanja, Minta anggaran, Kirim catatan yang sudah ada; menu tidak dibuka di contoh) | §9.1 baris "Dari Chat" |
| Isian | `textarea` kosong dengan label tampak atau label pembaca layar "Pesan untuk Sari" dan teks contoh "Tulis pesan" | §10.1 "Kotak tulis"; §25.11 pemeriksaan 18 |
| Tombol kirim | Ikon `send`, label pembaca layar "Kirim" | §10.1 |

### 3.5 Per lebar

- Ringkas: layar penuh, kotak tulis menempel di atas bilah navigasi (§16.1 baris Chat; §25.2).
- Sedang dan Lebar: layar penuh di tujuan Chat dengan rail; lebar gelembung dan kartu tidak ikut melebar (§16.2 baris Chat).
- Panel 30% di Beranda Lebar: isi yang sama, tanpa rail dan bilah navigasi sendiri (§1.3 di atas).

### 3.6 Informasi peka (Chat)

Jumlah di kartu tersamar bila "Sembunyikan angka" menyala; teks bebas tidak (§12.2 baris LY-60; fitur Tier 3, tidak tampil). Judul halaman "Casa · Chat" tanpa nama Sari (§25.9). Penutup netral seperti Beranda, tidak dibangun di contoh statis.

---

## 4. Ikon dan nama di layar utama

| Hal | Isi | Rujukan |
|---|---|---|
| Ikon `any` | `web/static/icons/casa-icon.svg`: simpul `kertas` terang di atas `teal` terang, satu bentuk besar. `web-dev` membuat PNG 192 dan 512 px dan ikon layar utama Apple 180 px dari berkas ini | §19.5 baris "Klien web" |
| Ikon `maskable` | `web/static/icons/casa-icon-maskable.svg`: gambar yang sama, simpul diperkecil sehingga seluruhnya di dalam lingkaran berjari-jari 295 px dari pusat kanvas 1024 (zona aman 80% berjari-jari 410 px) | §19.5; fokus tinjauan 4 |
| Nama | "Casa" di manifest (`name` dan `short_name`) dan di bawah ikon, tanpa data | §19.5; SEC-175 (g); ADR-030 butir 9 |
| `theme-color` | `kertas` per tema | §25.4 "Token" |

---

## 5. Pemeriksaan sendiri 11 sampai 15 (§19.8) atas isi ini

| # | Pemeriksaan | Hasil atas isi | Yang harus dijaga `web-dev` |
|---|---|---|---|
| 11 | Glif dan ikon sendiri | Lulus: tempat uang selalu memakai glif tunai, rekening, atau e-wallet dari `web/static/glyphs/`; ikon aplikasi adalah simpul | Jangan memakai simbol Material `payments`, `account_balance`, `wallet`, atau `savings` di tempat glif |
| 12 | Tanpa hiasan | Lulus: tidak ada ilustrasi, gradasi, karakter, foto, spanduk, atau pola | Tanpa bayangan selain lembar dan tombol Catat mengambang (§25.4) |
| 13 | Angka bertoken | Lulus: setiap jumlah punya token di tabel; tidak ada jumlah di dalam kalimat biasa | `tabular-nums` pada ketiga token; "Rp" sebagai unsur tersendiri |
| 14 | Motif hadir | Lulus: Beranda bagian 2 dan 3 (baris 6 sampai 11) dan kartu serah terima (kartu B) | Urutan label, angka rata kiri, garis tipis, baris pemegang tidak diubah per lebar |
| 15 | Kata templat | Lulus: tidak ada frasa dari daftar §19.8 butir 3, tanpa tanda seru dan emoji | Teks tombol, `aria-label`, dan judul halaman juga tunduk |

---

## 6. Yang sengaja tidak ada

Tombol "Sembunyikan angka" dan samaran (Tier 3); kata penyemangat (Tier 2); saran kata di chat (Tier 2); penutup netral dan keadaan luring (perlu JavaScript; keadaan browser diuji di layar 1.10, §25.11 pemeriksaan 20); grafik, tren, perbandingan antar-orang (§6.1).
