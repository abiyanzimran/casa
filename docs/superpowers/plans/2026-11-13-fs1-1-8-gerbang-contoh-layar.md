# FS-1 1.8 — Gerbang contoh layar web (08 §19.8, §25.11): rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** contoh layar statis Beranda, Catat, dan Chat yang dirender dengan token web arah C, terang dan gelap, di empat konteks pilot, beserta ikon dan nama Casa di layar utama. Pemilik produk dan pasangannya menilainya, lalu token dibekukan.

**Arsitektur:** token ditulis **sekali** sebagai CSS custom property di `web/src/lib/design/tokens.css`. Berkas ini menjadi sumber token klien web (08 §25.4) dan dipakai lagi oleh layar 1.10, bukan disalin. Contoh layar adalah HTML dan CSS statis tanpa JavaScript dan tanpa dependensi di `design-samples/`, yang merujuk token dan huruf lewat symlink. Contoh disajikan dari PC pemilik produk ke jaringan rumah hanya selama sesi gerbang.

**Stack:** HTML, CSS (custom property, media query, container query), SVG. Huruf Plus Jakarta Sans WOFF2 di origin sendiri (SEC-160 (d)).

**Spesifikasi:** `docs/08-ux-ui-strategy.md` §19.3 sampai §19.11, §25.1 sampai §25.4, §25.9, §25.11, §6.3, §16.1, §17.3, §18; `docs/adr/ADR-030-klien-web-pwa-lebih-dulu.md` butir 9 (nama Casa); `docs/10-security-privacy/05-klien-web.md` SEC-175.

**Pelaksana:** `ux-designer` (glif, ikon, isi contoh, tinjauan) dan `web-dev` (token, halaman, ikon raster, penyajian). **Peninjau:** `code-reviewer` untuk berkas di `web/`; `ux-designer` untuk QG-09 atas karya `web-dev`. **Cabang:** `fs1/1.8-design-gate`.

## Batasan global

- **Data buatan saja** (08 §19.8): nama, sebutan, jumlah, tempat uang, dan pesan rekaan. Tidak pernah data atau tangkapan layar rumah tangga pilot. Nama orang diambil dari contoh di 08 (mis. "Sari"); lead memastikan nama itu bukan nama anggota rumah tangga pilot sebelum ditunjukkan.
- Tanpa JavaScript, analitik, CDN, atau sumber luar. Semua berkas dari origin yang sama.
- Teks layar mengikuti 08 §18 dan daftar kata templat §19.8 butir 3 (tanpa "Selamat datang kembali", "Dashboard", "Yuk", tanda seru, emoji, dan seterusnya).
- Penulisan uang menurut §17.3: "Rp" dari token, angka dengan pemisah ribuan `id-ID`, tanpa desimal dan tanpa singkatan.
- Aset pihak ketiga (Plus Jakarta Sans, ikon Material Symbols): versi, tanggal cek, tautan, lisensi, dan checksum dicatat di PR (ES §8). Lisensi Material Symbols dicek live (08 §23.2 nomor 27); bila tidak jelas, berhenti dan lapor ke lead.
- `document.title` dan nama manifest "Casa" (SEC-175 (a), (g)).

## Fokus tinjauan

1. Mode gelap yang hanya membalik latar: setiap token punya nilai gelap dari §19.3, dan tidak ada warna yang ditulis langsung di halaman.
2. Teks 200% dan lebar 320 CSS px: jumlah uang tidak pernah terpotong atau tergulir dua arah (08 §15, §25.5 pemeriksaan 17).
3. Angka tabular tidak sama lebar di Safari (WebKit) karena fitur `tnum` hilang dari subset huruf.
4. Ikon `maskable` yang simpulnya terpotong zona aman di peluncur Galaxy A55.
5. Server LAN yang tanpa sengaja membuka seluruh repositori (dokumen perencanaan) ke jaringan rumah.

---

### Tugas A: Glif, ikon, dan isi contoh (`ux-designer`)

**Berkas:**
- Buat: `web/static/glyphs/{tunai,rekening,e-wallet,simpul}.svg`, `web/static/icons/casa-icon.svg` (simpul di atas `teal`), `web/static/icons/casa-icon-maskable.svg`
- Buat: `design-samples/CONTENT.md` (isi tiap contoh layar: bagian, baris, angka, nama rekaan, kalimat; bahasa Indonesia; rujukan nomor bagian 08 untuk setiap unsur)

- [ ] **Langkah 1:** gambar keempat glif menurut §19.5 (kisi 24, ketebalan garis setara ikon Material Symbols di sampingnya, warna `currentColor` sehingga mengikuti `tinta`/`tinta-2`, tidak pernah `teal`).
- [ ] **Langkah 2:** gambar ikon aplikasi: simpul tebal terang di atas `teal`, satu bentuk besar, tanpa unsur dari daftar "Yang tidak dipakai" §19.5. Versi `maskable`: simpul di dalam lingkaran zona aman berdiameter 80% kanvas.
- [ ] **Langkah 3:** tulis `CONTENT.md`: Beranda (bagian 1 sampai 3 menurut §6 dengan motif "angka bersama pemegangnya" §19.8 butir 5, termasuk satu bagian belum dikonfirmasi berwadah `pasir`), Catat (LY-32 dengan isian jumlah, tombol "000" di samping isian, pemilih tempat uang berglif, "Simpan" di kepala lembar, §25.3), dan Chat (utas dengan pesan sendiri `teal-muda`, pesan pasangan `kertas-2`, satu kartu permintaan dengan dua tombol sama tinggi, §10.4). Jumlah terpanjang yang muncul paling tidak `Rp88.888.888`, untuk menguji lebar.
- [ ] **Langkah 4:** periksa sendiri terhadap pemeriksaan 11 sampai 15 (§19.8). Serahkan ke `web-dev`.

### Tugas B: Token web (`web-dev`)

**Berkas:** Buat `web/src/lib/design/tokens.css`.

- [ ] **Langkah 1:** setiap token §19.3 menjadi custom property bernama sama (`--kertas`, `--kertas-1` sampai `--kertas-3`, `--tinta`, `--tinta-2`, `--garis-tegas`, `--garis`, `--teal`, `--di-atas-teal`, `--teal-muda`, `--pasir`), nilai terang di `:root`, gelap di `@media (prefers-color-scheme: dark)`, varian kontras tinggi di `@media (prefers-contrast: more)` (dan gabungannya dengan gelap), dan aturan `forced-colors: active` menurut §25.4. `color-scheme: light dark`.
- [ ] **Langkah 2:** token huruf dan angka §25.4 tabel "Huruf": `angka-utama` 2,5 rem, `angka-besar` 1,75 rem, `angka-baris` 1,0625 rem, `judul-bagian` 1,25 rem, isi 1 rem, terkecil 0,8125 rem; `font-variant-numeric: tabular-nums` pada token angka; `@font-face` Plus Jakarta Sans bobot 500, 600, 700 dari `/fonts/` dengan `font-display: swap`.
- [ ] **Langkah 3:** token jarak, sudut, gerak §25.4: kelipatan 4 px; sudut kartu 12, wadah tanda 8, tombol 12, sudut atas lembar 16; transisi paling lama 200 ms, pudar atau seketika di bawah `prefers-reduced-motion: reduce`.
- [ ] **Langkah 4: nilai kontras tinggi dihitung, bukan ditebak** (§19.3 "Kontras tinggi": sasaran 7:1 untuk semua teks). Lead memeriksa semua pasangan §19.3 di kedua tema dan varian kontras tinggi dengan skrip di scratchpad (rumus luminans relatif WCAG 2.2, dipotong satu desimal) terhadap nilai **di berkas**, bukan di tabel 08. Pasangan yang gagal sasaran dilaporkan ke `ux-designer`, yang menggeser nilainya di 08 §19.3 dan di `tokens.css` bersamaan.

### Tugas C: Huruf dan ikon raster (`web-dev`)

**Berkas:** Buat `web/static/fonts/PlusJakartaSans-{500,600,700}.woff2` (atau satu berkas variabel bila rilis resminya menyediakannya), `web/static/fonts/OFL.txt`, `web/static/icons/{icon-192.png,icon-512.png,icon-maskable-512.png,apple-touch-icon-180.png}`.

- [ ] **Langkah 1:** unduh Plus Jakarta Sans dari rilis resmi proyeknya (dicek live), verifikasi checksum bila disediakan, sertakan lisensi OFL. Pastikan fitur OpenType `tnum` ada di berkas yang dipakai.
- [ ] **Langkah 2:** rasterkan ikon SVG Tugas A ke keempat PNG dengan alat yang ada di mesin (ImageMagick `convert` tersedia; bila hasil SVG-nya cacat, render lewat browser tanpa kepala). Periksa visual tiap PNG.

### Tugas D: Halaman contoh (`web-dev`)

**Berkas:**
- Buat: `design-samples/index.html` (daftar tautan ketiga layar), `design-samples/beranda.html`, `design-samples/catat.html`, `design-samples/chat.html`, `design-samples/samples.css`, `design-samples/manifest.webmanifest`
- Buat symlink: `design-samples/tokens.css` → `../web/src/lib/design/tokens.css`, `design-samples/fonts` → `../web/static/fonts`, `design-samples/glyphs` → `../web/static/glyphs`, `design-samples/icons` → `../web/static/icons`
- Buat: `design-samples/serve.sh`

- [ ] **Langkah 1:** setiap halaman memakai elemen HTML dari tabel "Komponen per pola" §25.4 (`nav`, `header` dengan `h1`, `section`, `dialog` terbuka statis untuk Catat, `details`, `button`), isi dari `CONTENT.md`, dan hanya `var(--…)` dari `tokens.css` untuk warna, huruf, jarak. `samples.css` hanya berisi tata letak contoh, tidak ada nilai warna.
- [ ] **Langkah 2:** tata letak per keadaan menurut §25.3 dan §16.1: Ringkas (< 600), Sedang (600 sampai 839), Lebar (≥ 840), dengan Beranda dua kolom di Sedang dan Lebar dan panel Chat `aside` sekitar 30% (≥ 320 px) di Lebar; lembar Catat menempel di bawah di Ringkas dan di tengah ≤ 560 px di Sedang dan Lebar. `viewport-fit=cover` dan `env(safe-area-inset-*)`; meta viewport tanpa `maximum-scale` dan tanpa `user-scalable=no`.
- [ ] **Langkah 3:** `<head>`: `<title>Casa</title>` (Beranda boleh "Casa", layar lain "Casa"; tanpa data), `<link rel="manifest">` dengan `name`/`short_name` "Casa", `display: standalone`, tanpa `orientation`, ikon 192/512/maskable; `<link rel="apple-touch-icon" href="icons/apple-touch-icon-180.png">`; `theme-color` terang dan gelap dari token `kertas`.
- [ ] **Langkah 4: `serve.sh`** menjalankan `python3 -m http.server <port> --bind <IP LAN> --directory design-samples` (stdlib; IP LAN diberikan sebagai argumen, tidak pernah `0.0.0.0` diam-diam). **Uji batas:** dari mesin lain atau `curl`, minta `/../docs/README.md`, `/%2e%2e/CLAUDE.md`, dan `/tokens.css/../../CLAUDE.md` → semuanya 404; `/tokens.css` → 200. Catat hasilnya di PR.

### Tugas E: Pemeriksaan sebelum ditunjukkan

- [ ] **Langkah 1: angka tabular** (§19.11 butir 2): di Chromium dan WebKit tanpa kepala (Playwright, sesudah 1.1b merge; atau `npx` di scratchpad bila 1.1b belum merge), lebar render "Rp1.111.111" dan "Rp8.888.888" dengan token `angka-utama` sama persis. Catat angkanya.
- [ ] **Langkah 2: lebar dan teks besar** (§19.11 butir 5): tangkapan layar tiap halaman pada 320, 390, 744, 1133, 1280 CSS px, terang dan gelap, zoom 200%. Tidak ada gulir mendatar dan tidak ada jumlah terpotong. Tangkapan disimpan di scratchpad lead, **tidak** di repositori.
- [ ] **Langkah 3: tinjauan `ux-designer`** atas halaman (pemeriksaan 1, 3 sampai 15, 16 sampai 21 yang berlaku untuk halaman statis). Temuan diperbaiki sebelum langkah berikut.
- [ ] **Langkah 4: `code-reviewer`** atas `web/src/lib/design/tokens.css`, `web/static/**`, dan `design-samples/serve.sh`.

### Tugas F: Gerbang bersama pemilik produk dan pasangannya (**STOP**)

- [ ] **Langkah 1:** lead menjalankan `serve.sh` di PC pemilik produk dan memberi alamatnya. Pemilik produk dan pasangannya membuka contoh di: iPhone 13 (dipasang ke layar utama lewat Bagikan → Tambahkan ke Layar Utama), iPad mini 7 potret dan lanskap, Galaxy A55 di Chrome (tambahkan ke layar utama), dan satu browser komputer lebar; terang dan gelap (ganti tema sistem). Mereka melihat ikon dan nama "Casa" di layar utama iPhone dan di peluncur Galaxy A55.
  - Batas yang disampaikan lead sebelumnya: alamat LAN berupa `http`, jadi Chrome di Galaxy A55 mungkin membuat pintasan yang terbuka di tab, bukan aplikasi terpasang. Untuk gerbang ini yang dinilai adalah rupa, bukan perilaku pemasangan. Perilaku pemasangan dibuktikan di 1.12 dan di perangkat sesudah server uji ada.
- [ ] **Langkah 2:** lead menanyakan, dengan kata-kata 08 §19.8: apakah tampilannya terasa seperti templat; apakah simpul terbaca sebagai lambang tak hingga atau pita (§19.5 "Risiko"); apakah angka dan pemegangnya mudah dibaca. Jawaban dicatat **kata demi kata** di `_progress.md`.
- [ ] **Langkah 3:** bila dinilai masih seperti templat, yang diubah hanya kelima unsur §19.8 (glif, ikon, nada kata, tipografi angka, motif), bukan arah C. Ulangi Tugas A sampai E untuk unsur yang berubah, lalu langkah 1.
- [ ] **Langkah 4: token dibekukan.** `ux-designer` mencatat di 08 §19.3 dan §25.11 bahwa token dibekukan pada tanggal itu, dengan nilai akhir bila ada yang bergeser. Server LAN dimatikan.
- [ ] **Langkah 5: commit dan merge** sesudah tinjauan selesai dan CI hijau. Pesan contoh: `Freeze direction-C web design tokens after sample-screen gate (08 §19.8, §25.11)`.

## Selesai bila

- Pemilik produk dan pasangannya menyetujui contoh layar; kata-katanya tercatat di `_progress.md`.
- `tokens.css`, glif, ikon, dan huruf ada di `web/` dan dibekukan di 08.
- Tinjauan `ux-designer` dan `code-reviewer` tanpa temuan terbuka.
- Contoh layar tetap di `design-samples/` untuk uji prototipe 13 §6.3 (§19.8 "Dipakai ulang").
