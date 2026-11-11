# ADR-003 — Dana pribadi di luar garis kendali rumah tangga; satu buku per pemilik uang

| | |
|---|---|
| Status | Diusulkan; menunggu persetujuan pemilik produk di checkpoint S2. Satu konsekuensi sudah diputuskan pemilik produk pada 11 November 2026 (uang masuk yang diarahkan ke dana pribadi terlihat oleh pasangan) dan ditandai di tempatnya |
| Tanggal | 11 November 2026 |
| Pengusul | `finance-domain` (rekomendasi, 07 §1.5 ADR-B); ditulis `architect` |
| Rincian | [07-financial-domain.md](../07-financial-domain.md) §2.2, §3.1, §3.12, §5; INV-02, INV-03, INV-08, INV-20, INV-24, INV-25, INV-27, INV-28; CTH-03b |

Aturan uang kanonis di dokumen 07. ADR ini mencatat keputusannya dan tidak mengubah isinya.

## Konteks

Keputusan pemilik produk di checkpoint S1: dana pribadi ada dan saldonya tertutup bagi pasangan (nomor 3); pembayaran keperluan rumah tangga dari dana pribadi terlihat oleh keduanya sebagai kontribusi; hadiah untuk satu orang menjadi dana pribadinya (nomor 2). Kriteria sukses nomor 1 (brief §1.1) menuntut "berapa uang kita" terjawab sama bagi kedua pasangan. Angka rumah tangga yang memuat dana pribadi tidak dapat ditampilkan sama kepada keduanya tanpa membocorkan dana pribadi lewat pengurangan (Q-23).

## Opsi

| | A. Dana pribadi di dalam garis | B. Dana pribadi di luar garis |
|---|---|---|
| Angka utama | Berbeda per pelihat, atau membocorkan lewat pengurangan | Identik bagi keduanya: hanya dana bersama |
| Kontribusi dari dana pribadi | Pemindahan internal; tidak bermakna sebagai kontribusi | Uang atau nilai yang melintas masuk garis |
| Kekayaan bersih rumah tangga | Memuat harta pribadi | Tidak memuat harta pribadi |
| Perlindungan dana pribadi | Lewat penyaringan tampilan | Lewat struktur |

## Keputusan

**Opsi B.** Di dalam garis kendali rumah tangga hanya dana bersama. Setiap pemilik uang (`shared`, atau `member(M)`) punya satu buku; setiap entri seimbang di tiap buku yang disentuhnya; perpindahan nilai antar-pemilik berjumlah nol; klaim antar-pemilik selalu berupa sepasang posisi yang sama besar. Buku pribadi berada di ledger yang sama dan tunduk pada invarian yang sama, tetapi hanya terlihat oleh pemiliknya.

Tiga aturan yang menyertai keputusan ini (07 §2.2, §3.1; **[Usulan]** `finance-domain`, kecuali keterlihatan gaji, yang keputusan pemilik produk 11 November 2026):

1. **Aturan tulis (INV-25).** Entri yang memuat baris berpemilik `member(M)` hanya diterima dari M. Dua pengecualian, keduanya tanpa baris tempat uang atau aset milik M: (1) pelepasan klaim dana bersama atas M (`claim_waiver`), dicatat anggota dewasa selain M, yang hanya mengurangi kewajiban M; (2) penutupan klaim antar-pemilik saat M keluar (`exit_closure`), ditulis sistem (07 §8.2, INV-28; **[Usulan]**). Entri yang menyentuh buku pribadi dua anggota tidak punya perintah di MVP (07 T-12).
2. **Penghasilan yang dicatat selalu meninggalkan catatan bersama (INV-24).** Ini bukan jaminan bahwa penghasilan tidak dapat disembunyikan; lihat "Celah label" di bawah (SP-22). Baris pemasukan bersumber gaji atau hasil usaha selalu berpemilik `shared`. **[Usulan]** yang menunggu pemilik produk di checkpoint S2 (SP-01): bagian pribadi di dalam entri itu hanya sah bila pencatatnya adalah penerima penghasilan itu sendiri (07 INV-27; dokumen 06 §6.4). Aturannya per penerima penghasilan dan sama bagi anggota dewasa mana pun; kedua pasangan di rumah tangga pilot berpenghasilan (pernyataan pemilik produk, 11 November 2026). Bagian yang diarahkan pencatatnya ke dana pribadinya keluar garis dalam entri `income` yang sama, sebagai sepasang baris `owner_transfer`, sehingga jumlah penuh dan bagian pribadinya terbaca dari buku rumah tangga saja (CTH-03b).
3. **Buku pribadi melekat pada anggota, bukan pada akun** **[Usulan]** `architect`. `member(M)` berarti anggota dalam satu rumah tangga; orang yang ikut dua rumah tangga punya dua buku pribadi, dan tidak ada entri lintas rumah tangga (dokumen 06 §14). Menunggu konfirmasi `finance-domain`.

## Konsekuensi

- *Uang yang ada* dan kekayaan bersih rumah tangga dihitung dari buku rumah tangga saja dan identik bagi kedua pasangan (INV-20). "Berapa harta keluarga seluruhnya" tidak dijawab di layar mana pun.
- Yang terlihat pasangan dari dana pribadi hanyalah persinggungannya dengan dana bersama; kebocoran yang tersisa didaftar terbuka di 07 §5.3 (B1 sampai B5) untuk dinilai `security-privacy`.
- Tempat uang berpemilik `shared` tidak dapat bervisibilitas `private` (dokumen 06 §9.3, V-1).
- Satu tempat uang memuat uang tepat satu pemilik; dompet fisik campuran dicatat sebagai dua tempat uang (07 §2.3), dengan biaya usaha yang dinyatakan di 07 T-03.
- Uang masuk yang diarahkan ke dana pribadi saat dicatat (K-16, Q-56). **Keputusan pemilik produk, 11 November 2026:** pencatat uang masuk memilih tujuannya (dana bersama, langsung dipakai, atau dana pribadi; gaji termasuk), dan pasangan melihat catatannya, yaitu jumlah uang masuk dan jumlah yang diarahkan ke dana pribadi; saldo dana pribadi dan pemakaiannya sesudah itu tetap tertutup. Usulan `finance-domain` dan `architect` bahwa gaji selalu masuk dana bersama tidak diambil. Akibatnya bagi opsi B: buku rumah tangga memuat pemasukan penuh dan baris `owner_transfer` sebesar bagian pribadinya, dalam satu entri `income` dengan baris di dua buku (07 §3.1, INV-24, CTH-03b); sisi bacanya dijaga dokumen 06 §6.4 (V-8, INV-A29). Garis kendali dan INV-20 tidak berubah.
- **Celah label**, yang tersisa setelah keputusan itu dan tidak dapat ditutup model tanpa membuka seluruh pemasukan dana pribadi: INV-24 berlaku atas pemasukan yang **diberi sumber** gaji atau hasil usaha, dan sumber diisi pencatat. Penghasilan yang dicatat sebagai hadiah atau "lainnya" langsung di buku pribadi, atau tidak dicatat sama sekali, tidak meninggalkan catatan bersama; yang terlihat hanya bahwa gaji bulan itu tidak muncul (07 §5.3 B5). Selain itu tidak ada persetujuan dan tidak ada batas bagian yang dapat diarahkan (dokumen 06 §6.4; 07 §10.3 butir 1).
- Belum diputuskan pemilik produk dan diajukan di checkpoint S2 (dokumen 17 S2-1): apakah hanya penerima penghasilan yang boleh mengarahkannya (butir i; SP-01; dokumen 06 §6.4); apakah hadiah untuk satu orang, yang menurut usulan `finance-domain` tidak terlihat sama sekali, tetap demikian (butir ii); bagaimana bagian yang melampaui *jatah pribadi* ditandai (butir iii); dan apakah pekerjaan lepas dicatat sebagai hasil usaha (butir iv). Bagian yang diarahkan diperlakukan seperti penyisihan lain terhadap *ambang permintaan* dan jatah pribadi sebagai bawaan **[Usulan]**: di dalam jatah tidak ditandai (D-48), tanpa jatah mengikuti ambang (D-18; default Q-49).
- Atas entri dua buku, "kedua pasangan boleh mengoreksi" terbatas: hanya pemilik buku pribadinya yang dapat membaliknya; pasangannya dapat menyesuaikan sisi rumah tangga dan mengajukan keberatan (07 §3.12; dokumen 06 §8.1). Pengarahan gaji atau talangan karenanya tidak dapat dibatalkan pasangan sendirian.
- Karena buku pribadi melekat pada anggota, setiap entri dan setiap baris milik tepat satu rumah tangga, dan uji isolasi rumah tangga (`ENGINEERING-STANDARDS.md` §4) berlaku tanpa pengecualian bagi buku pribadi. Biayanya: orang di dua rumah tangga mencatat uang pribadinya terpisah di tiap rumah tangga. Apa yang boleh digabungkan di perangkat pemilik akun ditetapkan dokumen 10.
- Invarian diperiksa atas data penuh semua buku, bukan atas tampilan seorang anggota. Cara pembatasan akses baca dijalankan diputuskan di S4 (dokumen 09 dan 10).
