# ADR-026 — Pengiriman email: Amazon SES Jakarta

| | |
|---|---|
| Status | **Diusulkan** (`architect`, 12 November 2026); akan diputuskan pemilik produk di checkpoint S4 |
| Tanggal | 12 November 2026 (S4 fase D) |
| Pengusul | `architect`, atas temuan `critic` CS-A-11 |
| Mengubah | Menuntaskan butir email yang ditunda di dokumen 09 fase C |
| Rincian | [09.4 §2, §3, §7](../09-technical-architecture/04-operasi-uji-biaya.md) |

## Konteks

Pilot memerlukan email untuk verifikasi alamat, pemulihan akun dengan tautan sekali pakai 30 menit (SEC-24), undangan bila F04 memakai email (SEC-31), dan tanda ke operator bila cadangan atau pemeriksaan invarian gagal. Volumenya puluhan pesan per bulan **[Asumsi]**.

Batasan: server di Indonesia (BT-6); "no usage-priced managed services unless clearly justified" (BT-5); pihak penerima data sesedikit mungkin (dokumen 10, L-16).

Fakta yang diperiksa 12 November 2026:

- Amazon SES tersedia di `ap-southeast-3` (Jakarta) dengan endpoint API `email.ap-southeast-3.amazonaws.com`; **endpoint SMTP SES tidak tersedia di Jakarta** ([docs.aws.amazon.com/general/latest/gr/ses.html](https://docs.aws.amazon.com/general/latest/gr/ses.html)) **[Terverifikasi]**. Kuota kirim bawaan per wilayah 200 pesan per 24 jam (halaman yang sama) **[Terverifikasi]**.
- Sandbox SES berlaku per wilayah; di sandbox email hanya dapat dikirim ke alamat yang terverifikasi; keluar dari sandbox lewat permohonan *production access* (halaman SES "Request production access"; 18 §5.10) **[Terverifikasi]**.
- Pihak kontrak AWS bagi pelanggan Indonesia: PT Amazon Web Services Indonesia (09.1 §10) **[Terverifikasi]**.
- IDCloudHost menawarkan email hosting, tetapi harga dan izin SMTP keluar untuk pengiriman aplikasi tidak terbaca di halamannya. Biznet Gio tidak punya email hosting tersendiri, hanya add-on Plesk mulai Rp439.000 per bulan (18 §5.11) **[Terverifikasi]**, jauh di atas biaya SES pada volume pilot. Banyak penyedia VPS memblokir porta 25 keluar **[Asumsi]**.
- Tarif SES per pesan **[Perlu validasi]**.

## Opsi

| Kriteria | **A. Amazon SES Jakarta, API HTTPS** | B. SMTP email hosting bertarif tetap (IDCloudHost; Biznet Gio hanya add-on Plesk Rp439.000 per bulan) | C. Server SMTP sendiri di VPS | D. Penyedia transaksional lain (luar negeri) |
|---|---|---|---|---|
| Lokasi pemrosesan | Jakarta | Indonesia **[Perlu validasi]** | VPS sendiri | Luar negeri |
| Harga | Per pesan, mendekati 0 pada volume pilot | Tetap; harga tidak ditemukan | 0 | Paket gratis atau per pesan |
| Keterkiriman (tidak masuk spam) | Tinggi dengan SPF, DKIM, DMARC **[Asumsi]** | Bergantung reputasi IP bersama | Rendah: IP VPS baru tanpa reputasi; porta 25 mungkin diblokir | Tinggi |
| Dukungan aplikasi | API terdokumentasi; AWS SDK for Go v2 `sesv2` (1.x) | SMTP stdlib `net/smtp` | Komponen tambahan untuk dijaga | API |
| Penerima data tambahan | AWS (sudah opsi 09-B C) | Penyedia hosting (mungkin sama dengan cadangan) | Tidak ada | Satu pihak asing lagi |
| Terbukti dapat dipakai | Ya | Belum: izin pengiriman aplikasi tidak diketahui | Ya, tetapi mutu kirim buruk | Ya |

## Keputusan

1. **Amazon SES di `ap-southeast-3` lewat API HTTPS** (opsi A), dipanggil dari biner Go dengan AWS SDK for Go v2 (`sesv2`) **[Usulan]**.
2. **Alasan pengecualian BT-5** ("clearly justified"): tidak ada jalur bertarif tetap yang terverifikasi dapat dipakai aplikasi di Indonesia; biayanya mendekati nol pada volume pilot; dan pemrosesannya di Jakarta. Server SMTP sendiri ditolak karena email pemulihan yang masuk spam membuat pemulihan akun gagal (RISK-06).
3. **Isi yang dikirim:** hanya alamat penerima dan teks templat produk dengan tautan sekali pakai. Tidak ada nama rumah tangga, jumlah, judul, atau isi lain. Templat ada di kode, bukan di SES.
4. **Domain:** SPF, DKIM (Easy DKIM SES), dan DMARC di domain produk sebelum pilot.
5. **Pilot dalam sandbox:** penerimanya alamat pemilik produk dan pasangannya, yang diverifikasi di SES. Keluar dari sandbox (permohonan ke AWS) dilakukan sebelum pengguna di luar rumah tangga pilot.
6. **Kredensial:** kunci akses IAM dengan izin `ses:SendEmail` saja, dimuat lewat `LoadCredential=` systemd (09.4 §1; SEC-42).
7. **Tanda ke operator** memakai jalur yang sama ke alamat operator.
8. **Bila opsi B terbukti** (izin SMTP keluar untuk aplikasi, harga tetap, di Indonesia), ia dapat menggantikan SES tanpa mengubah modul lain, karena pengiriman dibungkus satu antarmuka di modul `notify`.

## Konsekuensi

- Satu penerima data tambahan (AWS) untuk alamat email saja. Dicatat di daftar prosesor dokumen 10 dan butir advokat L-16.
- Satu akun dan tagihan AWS tambahan, kecuali pemilik produk memilih 09-B C (Lightsail), yang memakai akun yang sama.
- Pengiriman email adalah pekerjaan latar (`job`), sehingga SES yang sedang gagal tidak menghalangi perintah apa pun; job dicoba ulang dengan jeda.
