# ADR-016 — PostgreSQL 18: tenancy, isolasi, transaksi, idempotensi, dan berkas

| | |
|---|---|
| Status | **Diterima** pemilik produk di checkpoint S4, 12 November 2026: PostgreSQL 18, satu skema dengan `household_id` di setiap kunci, RLS, `SERIALIZABLE`, foto di basis data pada pilot (D-99, diterima dengan "lanjut"; [_progress.md](../_progress.md), "Checkpoint S4: jawaban pemilik produk") |
| Tanggal | 12 November 2026 (S4 fase C; diperbarui fase D; status checkpoint S4); 13 November 2026 catatan FS-1 tugas 1.2 di butir 3 (bunyi kebijakan, pembuatan peran) |
| Pengusul | `architect` |
| Mengubah | [ADR-005](ADR-005-idempotency-key-per-perintah-tulis.md): menetapkan bentuk kunci, cara keunikan dijaga, dan jangka simpan yang diserahkannya ke dokumen 09. [ADR-008](ADR-008-jumlah-uang-bilangan-bulat.md): menetapkan representasi di basis data, kontrak, dan perangkat |
| Rincian | [09.2 §2 sampai §12](../09-technical-architecture/02-data-dan-skema.md) |

## Konteks

Brief §9.3 meminta keputusan atas pemodelan relasional, konsistensi transaksi, tingkat isolasi, idempotensi, kendali konkurensi, multi-tenancy (isolasi rumah tangga), dan evolusi skema.

Persyaratan yang berlaku:

- **07:** ledger append-only dengan baris berimbang per pemilik (INV-02), entri idempoten (INV-10), dan hasil bersamaan sama dengan berurutan (INV-23).
- **06:** compare-and-set (S-4), `seq` tanpa celah (S-5), dan penilaian berurutan per anggota per periode (S-8).
- **Dokumen 10:** `household_id` di setiap kunci (SEC-02), RLS (SEC-08), dan peran bukan superuser (SEC-09).

Riset 18 §5.5 mencatat PostgreSQL 18.6 (didukung sampai 14 November 2030). Ia juga mencatat bahwa RLS dilewati pemilik dan peran `BYPASSRLS`, dan bahwa unique serta foreign key diperiksa di luar RLS.

## Opsi

| Kriteria | **PostgreSQL 18** | SQLite server | Firestore |
|---|---|---|---|
| Serializable sejati | Ya (SSI) | Satu penulis | Tidak |
| RLS | Ya | Tidak | Aturan sendiri |
| PITR | Base backup + WAL | Alat replikasi berkas | Terkelola, per pemakaian |
| Jalan ke tahap publik | Server sendiri, replika | Terbatas | Per operasi (BT-5) |

| Kriteria tenancy | **Satu skema, `household_id` di setiap kunci** | Skema per rumah tangga | Basis data per rumah tangga |
|---|---|---|---|
| Migrasi | Satu kali | Sebanyak rumah tangga | Sebanyak rumah tangga |
| Satu akun di banyak rumah tangga | Wajar | Lintas skema | Lintas basis data |
| Isolasi | Kunci + RLS + uji | Skema | Paling kuat |

| Kriteria berkas pilot | **Di PostgreSQL (`bytea`)** | Disk VPS di luar direktori web | Penyimpanan objek self-hosted |
|---|---|---|---|
| Cadangan dan PITR | Satu jalur | Jalur kedua | Jalur kedua dan layanan tambahan |
| Penghapusan transaksional | Ya | Berkas yatim mungkin | Berkas yatim mungkin |
| Skala | Sampai batas pindah | Besar | Besar |

## Keputusan

1. **PostgreSQL 18** (minor stabil terbaru), dengan satu skema untuk semua rumah tangga.
2. **`household_id` memimpin setiap primary key, unique index, dan foreign key** milik rumah tangga. Rujukan lintas rumah tangga mustahil secara skema (SEC-02). Galat konstrain dipetakan ke galat domain tetap, dan "tidak ada" berbentuk sama dengan "tidak boleh" (SEC-05).
3. **RLS dengan `FORCE ROW LEVEL SECURITY` sejak pilot**, memakai `SET LOCAL app.household_id` per transaksi. `household_id` dari permintaan hanya pemilih: nilainya dicocokkan dengan keanggotaan aktif akun sesi sebelum dipasang (09.2 §3). `idempotency_record` dan `job` punya kebijakan untuk pemilik campuran (rumah tangga atau akun; job sistem tanpa rumah tangga).
   - Catatan 13 November 2026 (FS-1 tugas 1.2): pada koneksi pool, setelan `app.…` yang tidak dipasang di transaksi itu terbaca `''`, bukan NULL, sesudah koneksi pernah menjalankan `SET LOCAL`. Kebijakan yang membaca setelan dengan `missing_ok` karena itu memakai `nullif(current_setting('app.…', true), '')`; bunyi kebijakan, sumber, dan ujinya di 09.2 §3.
   - Catatan 13 November 2026 (FS-1 tugas 1.2): peran butir 4 berlaku untuk seluruh cluster dan dibuat sekali per cluster oleh superuser lewat `roles.sql`, bukan oleh migrasi (09.2 §3 lapis 5).
4. **Peran:**
   - `app_rw`: bukan pemilik, tanpa `BYPASSRLS` dan `SUPERUSER`.
   - `migrator`: pemilik skema.
   - `lifecycle_rw`: satu-satunya yang boleh `DELETE` atas data yang dapat dihapus menurut ADR-007.
   - `job_runner`: hanya mengambil baris `job` lintas rumah tangga; isi job dijalankan sebagai `app_rw` dengan rumah tangga job itu.
   - `backup`.
5. **Append-only lewat hak akses.** `entry`, `line`, `entry_link`, `object_event`, dan `security_event` hanya mendapat `INSERT` dan `SELECT` untuk semua peran aplikasi (INV-09); penghapusan karena jangka simpan lewat `lifecycle_rw`. `migrator` sebagai pemilik secara teknis dapat menulis ulang; pagarnya adalah aturan migrasi (butir 11), uji hash inti sebelum dan sesudah setiap migrasi (SEC-108), dan tinjauan dua pass.
6. **Isolasi:**
   - Setiap perintah tulis berjalan `SERIALIZABLE` dan dicoba ulang pada `40001` dan `40P01` sampai 5 kali.
     - Catatan 13 November 2026 (FS-1 1.4, putusan lead Q1): 5 coba ulang sesudah percobaan pertama (paling banyak 6 percobaan), jeda acak yang berlipat dari 10 ms **[Usulan]**; sesudahnya `RETRY_EXHAUSTED`. Rinciannya di 09.2 §6.
   - Bacaan berjalan `REPEATABLE READ READ ONLY`.
   - Penilaian jatah dan ambang mengambil kunci baris `assessment_lane(household_id, member_id, period_start)`, sehingga urutannya tetap (S-8).
7. **Konkurensi:**
   - Compare-and-set atas `version` objek (S-4); `terms_version` di tanggapan (S-3).
   - `object_event.seq` = versi objek, tanpa celah (S-5).
   - Tanda air `read_state` dengan `GREATEST`.
8. **Idempotensi** (menuntaskan ADR-005):
   - **Bentuk:** kunci UUIDv4 dibuat klien per niat dan disimpan di antrean sebelum dikirim.
   - **Keunikan:** `PRIMARY KEY (scope_id, idempotency_key)` di `idempotency_record`, beserta `request_hash`. Isi yang berbeda ditolak `idempotency_conflict`. Rekaman ditulis dalam transaksi yang sama dengan efeknya.
   - **Jangka simpan:** entri uang menyimpan kuncinya permanen di `entry`, sehingga unik per rumah tangga. Perintah lain disimpan 90 hari, dan klien tidak mengirim otomatis perintah non-uang yang lebih tua dari 30 hari.
9. **Uang** (melengkapi ADR-008):
   - **Basis data:** `bigint` dengan `currency char(3)` di setiap kolom jumlah, termasuk aturan, anggaran, otorisasi, reservasi, jadwal, ambang, dan jatah.
   - **INV-35:** foreign key `(household_id, currency) → household(id, currency) ON UPDATE RESTRICT` adalah satu-satunya penjaganya.
   - **Eksponen:** tabel `currency` milik produk, dengan IDR = 0, tidak pernah dari pustaka mata uang platform (ISO 4217 menyebut 2); setiap ekspor membawa `minor_unit_exponent` (07 INV-39).
   - **Kontrak dan perangkat:** protobuf `int64` di kontrak; `Int64` atau `Long` di perangkat; tanpa `float` di jalur uang.
10. **Keseimbangan INV-02 dan INV-03** diperiksa satu fungsi Go di `ledger` sebelum menulis. Tidak ada trigger yang mengulang persamaannya. Pemeriksaan malam memakai fungsi yang sama atas seluruh data.
11. **Evolusi skema:**
    - Migrasi goose berversi dengan arah naik dan turun, dalam pola expand–contract, sehingga biner sebelumnya tetap berjalan di atas skema baru.
    - Migrasi tidak pernah menulis ulang baris inti uang.
    - Basis data perangkat diperlakukan sebagai cache yang dapat dibangun ulang, kecuali tabel antrean.
12. **Foto di pilot disimpan di PostgreSQL** (`attachment_blob`), dengan batas pindah ke penyimpanan objek self-hosted pada 5 GB foto atau 20 GB basis data. Server hanya menerima JPEG dan re-encode ke JPEG saja dengan encoder standar Go (`image/jpeg` tidak menulis segmen metadata menurut kode sumbernya); PNG ditolak, karena `image/png` menulis chunk tambahan tRNS (09.2 §11, CR-A-07). `security-privacy` menerimanya dengan empat syarat SEC-66: uji korpus membuktikan keluaran tanpa APPn/COM (bagian PNG dipenuhi dengan tidak ada keluaran PNG); batas dimensi diperiksa sebelum dekode; format yang tidak dapat di-encode ulang ditolak; bila syarat gagal, libvips dengan `keep NONE` lewat cgo, dengan dependensi native dan permukaan dekoder C sebagai risiko yang diterima (09.2 §11).
13. **Pencarian Tier 1** memakai filter berindeks dan `ILIKE`, dibatasi satu rumah tangga, lewat jalur baca. `tsvector` dengan konfigurasi `indonesian` dan `pg_trgm` ditambahkan saat volume menuntut.

## Konsekuensi

- Isolasi rumah tangga dijaga tiga lapis: kode, kunci, dan RLS. Uji lintas rumah tangga dibangkitkan untuk setiap kueri.
- `SERIALIZABLE` menambah coba ulang di bawah persaingan. Pada beban pilot hal itu jarang. Jumlah coba ulang dipantau (09.4 §3).
- Hak akses append-only membuat koreksi apa pun harus berupa entri pembalik (ADR-002, INV-18). Kesalahan migrasi tidak dapat "membetulkan" baris lama.
- Foto memperbesar cadangan basis data. Batas pindah sudah dinyatakan.
- Satu pengecualian keunikan lintas rumah tangga tetap ada, yaitu email akun. Pendaftaran dan undangan karena itu bertanggapan seragam (SEC-30).
