Notasi: `saldo(p)` = jumlah `delta` semua baris posisi `p`; `U(o)`, `P(o)`, `A(o)`, `K(o)` = jumlah saldo tempat uang, piutang, aset, kewajiban berpemilik `o`; `KB(o) = U(o) + P(o) + A(o) − K(o)`.

| ID | Invarian | Cara menguji |
|---|---|---|
| INV-01 | Setiap `amount`, `delta`, `effect` adalah bilangan bulat dalam satuan terkecil dan membawa mata uang; semua baris satu entri bermata uang sama. Sejak penyisiran S4 (**[Usulan]** `finance-domain`, M5): hal yang sama berlaku atas setiap jumlah uang yang tersimpan di luar entri, yaitu di aturan, kesepakatan, jadwal, anggaran, otorisasi, dan reservasi (§2.1) | Pemeriksaan tipe dan skema; perintah berisi pecahan ditolak; perintah yang menyimpan jumlah aturan tanpa mata uang ditolak |
| INV-02 | **Keseimbangan per pemilik.** Untuk setiap entri `e` dan setiap pemilik `o`: `Σ delta(tempat uang) + Σ delta(piutang) + Σ delta(aset) − Σ delta(kewajiban) = Σ effect(sebab)`, dihitung atas baris `e` yang berpemilik `o` | Untuk setiap entri yang tersimpan |
| INV-03 | **Perpindahan antar-pemilik berjumlah nol.** Dalam satu entri, `Σ effect` semua baris `owner_transfer` di semua buku = 0 | Untuk setiap entri |

### 7.2 Saldo awal dan custody

#### CTH-01 — Saldo awal

#### CTH-03b — Gaji masuk, sebagian langsung ke dana pribadi (berdiri sendiri)

### 7.10 CTH-37 — Ringkasan bulanan rumah tangga contoh


| Bagian | Perubahan |
|---|---|
| INV-26 | Syarat tautan berlaku saat tautan dibuat; otorisasi `open` tanpa syarat tanggal; tautan susulan ke tagihan berulang hanya oleh pencatat; uji ditambah |

Not a definition: a fenced block in a document is never read as one.

```markdown
#### CTH-99 — Fenced example

| ID | Invarian | Cara menguji |
|---|---|---|
| INV-99 | Fenced invariant | none |
```
