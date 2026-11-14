# 09.2 — Modul, data, dan skema

Bagian dari [dokumen 09](../09-technical-architecture.md). Berkas ini menetapkan:

- batas modul di dalam monolit;
- tempat setiap jenis data disimpan;
- sketsa skema relasional yang diturunkan dari 06 dan 07;
- isolasi rumah tangga, transaksi, idempotensi, konkurensi, dan evolusi skema;
- penyimpanan berkas dan pencarian.

Entitas domain dan medannya **kanonis di 06 dan 07**. Berkas ini hanya menambah kolom teknis (kunci, versi, indeks, peran basis data), dan tidak mendefinisikan ulang aturan apa pun. Semua butir **[Usulan]** `architect`, kecuali yang diberi label lain.

---

## 1. Modul di dalam monolit

Satu biner Go, satu proses, dan satu basis data (ADR-014). Setiap modul adalah satu paket di bawah `internal/`. Modul memiliki tabelnya sendiri dan hanya menulis tabelnya sendiri. Modul lain memanggilnya lewat fungsi Go di dalam **transaksi yang sama**, tidak lewat jaringan.

| Modul | Isi | Tabel milik | Sumber aturan |
|---|---|---|---|
| `identity` | Akun, kata sandi, sesi per perangkat, token penyegar, autentikasi ulang, pemulihan email, batas laju, pendaftaran tertutup dengan token pemasangan. Catatan 13 November 2026 (FS-1 1.4): kueri sqlc-nya di `internal/identity/identitydb` | `account`, `credential`, `session`, `refresh_token`, `reauth_grant`, `recovery_token`, `setup_token`, `signup_verification`, `security_event` | Dokumen 10 SEC-11 sampai SEC-27, SEC-158 |
| `household` | Rumah tangga, keanggotaan, peran, undangan, keluar dan masa tenggang, pihak luar, pemberitahuan operator, `full_adult_members`, sebutan anggota per pembaca (Tier 2, §5.5). Catatan 14 November 2026 (FS-1 tugas 1.6, putusan `architect` Q1, Q3, Q14): juga keempat konsep 06 §3.1 selain hubungan yang ditulis saat rumah tangga dibuat, tabel hubungan, dan **satu-satunya penulis `object_event`** (06 §2.5, INV-A01): modul lain yang menulis Peristiwa (`request` sejak langkah 4, `budget` untuk kesepakatan) memanggil fungsi `household` di transaksi yang sama, karena semuanya sudah bergantung pada `household` dan `household` tidak bergantung pada mereka. Teks pemberitahuan pengelola **bukan tabel**: setiap versi ikut di biner sebagai data yang di-`embed` di paket `household` (SEC-35, terbit bersama build), dan uji menyematkan hash setiap versi yang sudah terbit sehingga versi lama tidak dapat diubah; `operator_notice` dihapus dari daftar tabel. `membership.account_id` dan `display_name` adalah pemetaan identitas anggota (07 §8.2); `identity_map` di `ledger` hanya untuk pihak luar, dan perlu tidaknya tabel itu di samping `external_party` diputus di langkah yang membangun pihak luar. Kueri sqlc di `internal/household/householddb` | `household`, `membership`, `approval_authority`, `household_rule`, `relationship`, `invitation`, `object_event`, `external_party`, `operator_notice_ack`, `consent_record`, `member_label`. Ditambah tugasnya: `responsibility` (langkah 3, bersama tagihan berulang) | 06 §2.5, §3; SEC-10, SEC-28 sampai SEC-34, SEC-109, SEC-155 |
| `ledger` | Entri, baris, tautan, rincian, posisi (tempat uang, piutang, kewajiban, aset), rekonsiliasi | `entry`, `line`, `entry_link`, `entry_detail`, `identity_map`, `holding`, `receivable`, `liability`, `asset`, `reconciliation` | 07 §1 sampai §3; ADR-002, ADR-004, ADR-007 |
| `budget` | Kategori, tujuan, periode, anggaran, otorisasi, reservasi, jatah, tagihan berulang, cicilan, jadwal penyisihan, sumber penghasilan, penilaian tanda; **kecocokan kategori** tautan ke pos rutin dan larangan mengganti kategori entri yang tertaut (satu fungsi; SEC-92; 07 §3.15 butir 7b, INV-37) | `category`, `purpose`, `budget_period_setting`, `budget`, `budget_rule`, `authorization`, `reservation`, `installment`, `recurring_bill`, `set_aside_schedule`, `income_source`, `assessment_lane`, `entry_flag` | 07 §2, §3, §6; 06 S-8 |
| `request` | Permintaan, versi syarat, tanggapan, kesepakatan rumah tangga. Catatan 14 November 2026 (FS-1 tugas 1.6): Peristiwa ditulis lewat fungsi `household` (baris `household`); `request` menambah `approval_mode` ke `household_rule` (expand) bersama `resolve_approvers` | `request`, `request_terms`, `request_response` | 06 §2, §4 sampai §8 |
| `notify` | Niat notifikasi, penahanan jam tenang, ringkasan, pengingat, token push | `notification`, `push_token`, `digest_state` | 06 §10, §11; SEC-50 sampai SEC-58 |
| `chat` | Percakapan, pesan, kartu, tanda air, kehadiran | `conversation`, `message`, `read_state`, `presence_last_active` | 06 §16; ADR-011 |
| `files` | Lampiran dan isinya, re-encode, thumbnail | `attachment`, `attachment_blob` | SEC-65, SEC-66 |
| `sync` | Umpan per pembaca, kursor, pembuangan | `reader_feed`, `reader_feed_counter` | ADR-017, ADR-018 |
| `jobs` | Antrean pekerjaan dan penjadwal | `job` | Berkas 03 §9 |
| `lifecycle` | Ekspor, penghapusan 90 hari, jurnal penghapusan, hapus akun | `export_link`, `deletion_journal` | SEC-94 sampai SEC-106 |
| `ai` (Tier 2) | Kumpulan kalimat, penghitung biaya, saklar mati | `sentence_pool`, `ai_usage` | Dokumen 10 §10.4; ADR-021 |
| `care`, `shopping`, `maintenance`, `plan` (Tier 2) | Sesuai 06 §13 dan tier-2.md | Tabel masing-masing | 06 §13 |
| `schema` (infrastruktur, `internal/schema`) | Migrasi goose (§10) dan tabel infrastruktur yang tidak dimiliki modul domain. Catatan 13 November 2026 (FS-1 tugas 1.2): `currency` dibuat di tugas 1.2. Baris `idempotency_record` ditulis jalur perintah di dalam transaksi perintah itu (§8); `schema` memiliki DDL-nya | `currency`, `idempotency_record` | ADR-016 butir 8, 9, 11 |
| `database` (infrastruktur, `internal/database`) | Catatan 13 November 2026 (FS-1 tugas 1.2): transaksi bercakupan. `Scope` membawa konteks (`app.household_id` atau `app.account_id`) yang dipasang `SET LOCAL` di awal transaksi; `Write` membuka transaksi `SERIALIZABLE`, `Read` membuka `REPEATABLE READ READ ONLY` (§3 lapis 4, §6). Tidak mengimpor paket `internal/` lain | Tanpa tabel | ADR-016 |
| `api` (lapisan API, `internal/api`) | Catatan 13 November 2026 (FS-1 1.3, kontrak dasar): server Connect tempat setiap layanan mendaftar, mux HTTP dengan `/healthz` (09.4 §3) dan batas ukuran permintaan (09.4 §1), serta satu interseptor batas dengan urutan tetap: deskriptor metode → `request_context` → versi klien minimum (perintah tulis dari `web`) → kunci idempotensi (perintah tulis, §8) → protovalidate atas seluruh pesan. Versi diperiksa sebelum validasi badan, sehingga klien lama mendapat `CLIENT_UPGRADE_REQUIRED` dan perintahnya tetap di antrean. Catatan 13 November 2026 (FS-1 1.4): langkah autentikasi sesi masuk **sesudah deskriptor dan sebelum pesan diterima**, sehingga urutannya deskriptor metode → autentikasi sesi → `request_context` → versi klien minimum → kunci idempotensi → protovalidate. Setiap prosedur wajib membawa `Authorization: Bearer <token akses>` yang diterima `identity`, kecuali daftar prosedur publik, yang hanya boleh memuat prosedur `casa.v1.IdentityService` yang dipakai sebelum ada sesi (pendaftaran, masuk, penyegaran, pemulihan); uji `internal/api` menolak prosedur publik di luar layanan itu. Di 1.4 jalur header hanya menerima sesi `ios` dan `android`; sesi `web` lewat cookie di 1.5 (SEC-161 (b)). Catatan 13 November 2026 (FS-1 tugas 1.5, rencana 1.5 keputusan 11): urutannya menjadi deskriptor metode → **pemeriksaan origin pra-isi** → autentikasi sesi (`Authorization: Bearer` untuk native, atau cookie `__Host-casa_session` untuk web; cookie hasil rotasi dipasang di sini) → `request_context` → **pemeriksaan origin bagi `client_kind = WEB`, kecocokan jenis klien dengan sesi, sentuhan diam mode bersama dari `user_initiated`** → versi klien minimum → kunci idempotensi → protovalidate (cakupan di ADR-030 butir 6 catatan tanggal itu). `api` juga memuat penyaji bundel (`Bundle`) dan pembungkus header keamanan di depan seluruh mux (ADR-030 butir 7), tetapi menerima bundel sebagai `fs.FS` dan tidak mengimpor `internal/webbundle` | Tanpa tabel | ADR-013 butir 4; ADR-005; ADR-030 butir 6, 7, 9; SEC-169 (d); SEC-15; SEC-162 |
| `webbundle` (infrastruktur, `internal/webbundle`) | Catatan 13 November 2026 (FS-1 tugas 1.5): hanya `//go:embed all:dist`, salinan `web/build` yang diisi `make bundle-copy` (ADR-030 butir 3). Tanpa logika dan tanpa impor paket `internal/` | Tanpa tabel | ADR-030 butir 3 |
| `cmd/webpreview` (alat pengembangan) | Catatan 13 November 2026 (FS-1 tugas 1.5, putusan lead Q6): biner yang tidak pernah di-deploy; menyajikan `internal/webbundle` lewat `api.Bundle` dan pembungkus header keamanan yang sama dengan `casa`, tanpa API dan basis data. Server Playwright (09.4 §4) | Tanpa tabel | ADR-030 butir 3, 7 |
| `gen/casa/v1` (kode hasil generate, `internal/gen/casa/v1`) | Catatan 13 November 2026 (FS-1 1.3, kontrak dasar): kode Go hasil `buf generate` dari `proto/casa/v1`. Tidak disunting tangan (ES §2) dan dibangkitkan ulang di CI yang menolak beda (14 §4.6). Tidak mengimpor paket `internal/` lain | Tanpa tabel | ADR-013 butir 4 |
| `gen/casa/v1/casav1connect` (kode hasil generate, `internal/gen/casa/v1/casav1connect`) | Catatan 13 November 2026 (FS-1 1.4): kode handler dan klien Connect hasil `protoc-gen-connect-go` dalam `buf generate` yang sama. Tidak disunting tangan dan diperiksa CI seperti `gen/casa/v1`. Hanya mengimpor `internal/gen/casa/v1` | Tanpa tabel | ADR-013 butir 4 |
| `command` (infrastruktur, `internal/command`) | Catatan 13 November 2026 (FS-1 1.4): jalur perintah. `command.Run` menjalankan satu perintah tulis di dalam satu `database.Write`: mengembalikan respons tersimpan untuk permintaan identik, menolak isi berbeda di kunci yang sama (`IDEMPOTENCY_CONFLICT`), atau menjalankan perintahnya lalu menulis `idempotency_record` (§6 langkah 1 dan 8, §8). Medan kontrak bertanda `debug_redact` dikosongkan sebelum `request_hash` dihitung, dan respons yang memuat medan itu tidak disimpan (SEC-11). Kueri sqlc-nya di `internal/command/commanddb` | Tanpa tabel (menulis `idempotency_record` milik `schema`) | ADR-005; ADR-016 butir 8 |
| `identity/identitydb`, `command/commanddb` (kode hasil generate sqlc) | Catatan 13 November 2026 (FS-1 1.4): kode Go hasil `sqlc generate` dari `queries.sql` modulnya. Tidak disunting tangan; CI menolak beda (`sqlc-generated`). Tidak mengimpor paket `internal/` lain | Tanpa tabel | ES §2, §3 |
| `read` | **Satu jalur baca**: `can_read(member, object)`, proyeksi layar, laporan, ekspor, dan masukan `AI` | Tanpa tabel; fungsi dipakai semua modul | SEC-06; LM-07 |

**Aturan dependensi** (diperiksa uji arsitektur di CI, ADR-014):

1. `ledger` hanya ditulis oleh `ledger`. `budget` dan `request` memanggil fungsinya.
2. `read` adalah satu-satunya jalan bagi modul mana pun untuk membaca data milik orang lain. Ekspor, laporan, `sync`, dan `ai` memakai `read`, tidak menulis kueri sendiri (SEC-06).
3. `chat` tidak membaca isi objek lain. Kartu dirender dari `read` saat ditarik (INV-A44).
4. Tidak ada modul yang membaca `read_state` atau `presence` selain `chat` (LM-11; INV-A45).
5. Paket infrastruktur (catatan 13 November 2026, FS-1 tugas 1.2): `internal/database` tidak mengimpor paket `internal/` lain. Yang boleh mengimpor `schema` hanya `cmd/casa` (perintah `casa migrate`) dan `internal/database/databasetest` (menyiapkan basis data uji lewat migrasi yang sama); modul domain tidak mengimpornya.
6. Lapisan API dan kode hasil generate (catatan 13 November 2026, FS-1 1.3, kontrak dasar): `internal/gen/casa/v1` tidak mengimpor paket `internal/` lain. `internal/api` hanya mengimpor `internal/gen/casa/v1`, dan tidak pernah `internal/schema`. `cmd/casa` mengimpor `internal/api` dan `internal/schema`. Tepi baru, misalnya `internal/api` ke modul domain sejak layanan pertama (1.4), ditambahkan ke aturan ini dan ke daftar izin uji arsitektur (`allowedImports`) di PR yang sama.
7. Tepi FS-1 1.4 (catatan 13 November 2026, putusan lead Q12; menggantikan daftar `internal/api` dan `cmd/casa` di butir 6):

   | Paket | Boleh mengimpor (paket `internal/` saja) |
   |---|---|
   | `cmd/casa` | `internal/api`, `internal/database`, `internal/gen/casa/v1`, `internal/identity`, `internal/schema` |
   | `internal/api` | `internal/command`, `internal/database`, `internal/gen/casa/v1`, `internal/gen/casa/v1/casav1connect`, `internal/identity` |
   | `internal/identity` | `internal/database`, `internal/identity/identitydb` |
   | `internal/command` | `internal/command/commanddb`, `internal/database` |
   | `internal/identity/identitydb`, `internal/command/commanddb` | tidak ada |
   | `internal/gen/casa/v1/casav1connect` | `internal/gen/casa/v1` |

   `internal/identity` dan `internal/command` tidak saling mengimpor dan tidak mengimpor kode kontrak: lapisan API yang memetakan pesan kontrak ke panggilan modul dan galat modul ke kode kontrak (ADR-013, catatan 13 November 2026 FS-1 1.4).
8. Tepi FS-1 tugas 1.5 (catatan 13 November 2026; menambah butir 7): `cmd/casa` juga mengimpor `internal/webbundle`; `cmd/webpreview` hanya mengimpor `internal/api` dan `internal/webbundle`; `internal/webbundle` tidak mengimpor paket `internal/` lain. **Hanya `cmd/casa` dan `cmd/webpreview` yang mengimpor `internal/webbundle`**; `internal/api` tidak, sehingga uji `internal/api` tidak bergantung pada build web. Uji arsitektur (`allowedImports`) memuat tepi ini.
9. Tepi FS-1 tugas 1.6 (catatan 14 November 2026, putusan `architect` Q10; menambah butir 7 dan 8):

   | Paket | Tepi baru (paket `internal/` saja) |
   |---|---|
   | `cmd/casa` | `internal/household` |
   | `internal/api` | `internal/household` |
   | `internal/household` | `internal/database`, `internal/household/householddb` (seluruh daftarnya) |
   | `internal/household/householddb` | tidak ada |

   **`household` dan `identity` tidak saling mengimpor.** Perintah yang menyentuh keduanya disusun lapisan API di dalam satu transaksi `command.Run` atau `command.RunForAccount`; `identity` menerima pemeriksaan undangan sebagai parameter fungsi, bukan sebagai impor. Ekspor baru `identity`: `RecordSecurityEvent`, `SignupVia`, `AccountEmail`, `TokenUseAttempt`, `NewToken`, `HashToken`, `SignupInvitation`; ekspor baru `command`: `RunForAccount`, `StoredForAccount` (§8: `scope_id = account_id`). Token undangan dibuat dan di-hash hanya oleh `identity.NewToken` dan `identity.HashToken`, dipanggil lapisan API: `household` hanya menerima hash, sehingga pembuat token tetap satu implementasi. Pembacaan email penerima lewat `identity.AccountEmail` oleh lapisan API (rencana 1.6 keputusan 16) adalah pengecualian sementara atas aturan 2: sejak `read` dibangun (1.7), pembacaan itu lewat `read`.
10. Tepi FS-1 1.7 (catatan 14 November 2026, putusan `architect` Q1 (a); menambah butir 7 sampai 9):

    | Paket | Tepi baru (paket `internal/` saja) |
    |---|---|
    | `internal/jobs` | `internal/database`, `internal/jobs/jobsdb` (seluruh daftarnya) |
    | `internal/read` | `internal/database`, `internal/read/readdb` (seluruh daftarnya) |
    | `internal/sync` | `internal/database`, `internal/read`, `internal/sync/syncdb` (seluruh daftarnya) |
    | `internal/household` | `internal/jobs`, `internal/sync` |
    | `internal/api` | `internal/read`, `internal/sync` |
    | `cmd/casa` | `internal/jobs`, `internal/read`, `internal/sync` |
    | `internal/jobs/jobsdb`, `internal/read/readdb`, `internal/sync/syncdb`, `internal/isolation` (uji saja) | tidak ada |

    Arahnya `household → sync → read` dan `household → jobs`. `read` dan `jobs` tidak mengimpor modul domain mana pun, jadi tidak ada siklus. `sync` mengambil daftar anggota dan `can_read` dari `read`, tidak dari `household`. `sync.Changed` adalah satu-satunya penulis `reader_feed`: `household.RecordEvent` memanggilnya di transaksi yang sama, dan modul yang kelak mengubah proyeksi tanpa Peristiwa memanggilnya langsung (tepinya ditambahkan di sini bersama modul itu). Penangan job didaftarkan `cmd/casa` ke `jobs.Runner` sebagai fungsi, sehingga `jobs` tidak mengimpor modul yang job-nya ia jalankan. **Pengecualian sementara butir 9 berakhir:** `identity.AccountEmail` pindah menjadi `read.AccountEmail` dan ekspornya dihapus dari `identity`; aturan 2 kembali berlaku tanpa pengecualian.

```mermaid
flowchart LR
  subgraph api["Lapisan API (Connect + SSE)"]
    H[Handler per layanan]
  end
  H --> identity
  H --> household
  H --> request
  H --> ledger
  H --> budget
  H --> chat
  H --> files
  H --> sync
  H --> read
  household --> sync
  household --> jobs
  request --> budget
  request --> ledger
  budget --> ledger
  chat -.kartu.-> read
  sync --> read
  lifecycle --> read
  ai --> read
  request --> notify
  ledger --> notify
  chat --> notify
  notify --> jobs
  household --> lifecycle
  subgraph db[(PostgreSQL 18)]
  end
  identity & household & request & ledger & budget & chat & files & sync & notify & jobs & lifecycle & ai & read --> db
```

---

## 2. Tempat setiap jenis data (brief §9.3)

| Jenis data | Penyimpanan | Alasan | Sisi perangkat |
|---|---|---|---|
| Catatan keuangan (entri, baris, tautan, posisi) | PostgreSQL, tabel append-only (§6) | Transaksi Serializable; INV-09 dijaga hak akses | Salinan baca sesuai SEC-47; antrean perintah |
| Akun, kredensial, sesi | PostgreSQL, tabel global (bukan milik rumah tangga) | Satu akun di banyak rumah tangga (06 §14) | Token di Keychain atau Keystore (SEC-15) |
| Hubungan dan keanggotaan | PostgreSQL (`membership`); hubungan kekerabatan **tidak disimpan** (06 §3.1, PG-01) | — | Daftar anggota yang terbaca |
| Chat | PostgreSQL (`message`), kelas peka (SEC-63) | Dua peserta; volume kecil | Salinan pesan yang terbaca |
| Notifikasi | PostgreSQL (`notification`, dengan `hold_until` dan `released_at`) | Penahanan jam tenang dan ringkasan adalah state di server | Tidak disimpan aplikasi selain pusat notifikasi OS |
| Lampiran (foto) | **PostgreSQL** (`attachment_blob`, `bytea`) di pilot (§10) | Satu jalur cadangan dan PITR; penghapusan transaksional; tidak ada direktori web | Cache foto di direktori cache, dikecualikan dari cadangan |
| Persediaan (`INV`, Tier 3) | PostgreSQL, modul sendiri saat dibangun | Sama dengan modul lain | — |
| Rencana (`PLN`, Tier 2) | PostgreSQL; di MVP menumpang syarat permintaan jalan-jalan (F-REQ-002) | Tidak ada objek rencana terpisah | Salinan baca |
| Riwayat audit | `object_event` per objek (06 S-5); entri itu sendiri (append-only); `security_event` untuk sesi dan akses operator, terbaca hanya oleh pemilik akun (SEC-16, SEC-17) | Tidak ada riwayat lintas objek menurut orang (06 AT-3) | — |
| Pencarian | PostgreSQL: filter berindeks + `ILIKE` per rumah tangga di Tier 1; `tsvector` `indonesian` + `pg_trgm` saat volume menuntut (§11) | Tanpa mesin pencari terpisah | — |
| Kehadiran | Online dan mengetik: **memori proses saja**. Terakhir aktif: satu baris per (anggota, rumah tangga) yang ditimpa, tanpa riwayat (D-95) | 06 §16.6; INV-A45 | — |
| Status `CARE` yang tidak dibagikan | **Perangkat pemiliknya saja** (10-D A, diterima di S4-6 (i); ADR-024 butir 5) | Tidak pernah meninggalkan perangkat sebelum dibagikan | Basis data lokal, Class A |
| Sebutan anggota per pembaca (F-FAM-021, Tier 2) | PostgreSQL (`member_label`), setelan akun milik pembuatnya (§5.5) | Berlaku di semua perangkat pembuat (08 §7.4); terbaca hanya oleh pembuat (SEC-155 (b)) | Salinan baca hanya di perangkat pembuat |
| Log | journald, 14 hari (SEC-44) | — | Tidak ada log isi di perangkat |
| Cadangan | VPS cadangan di Indonesia, terenkripsi sebelum keluar (ADR-020) | SEC-43 | Basis data lokal dikecualikan dari cadangan OS (SEC-46) |

---

## 3. Isolasi rumah tangga dan multi-tenancy

**Model:** satu skema untuk semua rumah tangga, dan kolom `household_id` di setiap tabel milik rumah tangga. Skema atau basis data per rumah tangga ditolak: migrasi dan cadangan berlipat sebanyak rumah tangga tanpa manfaat bagi tim satu sampai tiga orang.

| Lapisan | Mekanisme | Persyaratan |
|---|---|---|
| 1. Otorisasi di server | Setiap handler memanggil `authorize(command, actor)` dan setiap bacaan lewat `read.can_read`. Satu akun dapat menjadi anggota beberapa rumah tangga (06 §14), jadi permintaan boleh membawa `household_id` **sebagai pemilih saja**: sebelum apa pun dibaca, server memeriksa bahwa akun sesi itu punya keanggotaan aktif (atau dalam masa tenggang, menurut perintahnya) di rumah tangga itu; bila tidak, jawabannya sama dengan "tidak ada" (SEC-05). Sesudah lolos, `household_id` yang sudah diperiksa itulah yang dipasang sebagai `SET LOCAL app.household_id` dan diteruskan ke kueri sqlc; nilai mentah dari badan permintaan tidak pernah dipakai langsung | SEC-01, ES §4 |
| 2. Kunci | **`household_id` memimpin setiap primary key, unique index, dan foreign key.** Contoh: `PRIMARY KEY (household_id, id)`, `FOREIGN KEY (household_id, holding_id) REFERENCES holding (household_id, id)`. Rujukan lintas rumah tangga mustahil secara skema | SEC-02 |
| 3. Pesan galat | Pelanggaran unique atau foreign key dipetakan ke galat domain yang tetap. Teks galat PostgreSQL tidak pernah diteruskan ke klien. "Tidak ada" dan "tidak boleh" berbentuk identik, termasuk kode HTTP dan waktu tanggapnya, sejauh dapat diukur | SEC-02, SEC-05 |
| 4. RLS | `ALTER TABLE … ENABLE ROW LEVEL SECURITY; ALTER TABLE … FORCE ROW LEVEL SECURITY;` dengan kebijakan `household_id = current_setting('app.household_id')::uuid`. Setiap transaksi memulai dengan `SET LOCAL app.household_id`. Dinyalakan sejak pilot (dianjurkan dokumen 10 §4) | SEC-08 |
| 5. Peran basis data | `app_rw`: bukan pemilik tabel, tanpa `BYPASSRLS`, tanpa `SUPERUSER`. `migrator`: pemilik skema, dipakai hanya oleh perintah migrasi. `lifecycle_rw`: satu-satunya peran dengan `DELETE` atas `entry_detail`, `identity_map`, dan data yang dihapus menurut ADR-007. `job_runner`: mengambil job (tabel di bawah). `backup`: peran replikasi untuk pgBackRest. Catatan 13 November 2026 (FS-1 tugas 1.2): peran berlaku untuk seluruh cluster, bukan per basis data ("Database roles are global across a database cluster installation", [postgresql.org/docs/18/database-roles.html](https://www.postgresql.org/docs/18/database-roles.html)), dan hanya superuser atau peran ber-`CREATEROLE` yang dapat membuat peran ([postgresql.org/docs/18/role-attributes.html](https://www.postgresql.org/docs/18/role-attributes.html); keduanya diakses 13 November 2026) **[Terverifikasi]**. `migrator` tidak diberi `CREATEROLE`, jadi peran dibuat oleh berkas `roles.sql` yang dijalankan **sekali per cluster oleh superuser** (peran `postgres` dari paket PGDG), di VPS, staging, dan CI sebelum migrasi pertama, bukan oleh migrasi goose. Berkas itu tidak memuat kata sandi; migrasi hanya memberi hak (`GRANT`) atas objek milik `migrator` kepada peran yang sudah ada. Catatan 13 November 2026 (FS-1 1.2, temuan `code-reviewer` I4): karena nama peran tetap dan berlaku se-cluster, **staging dan produksi memakai cluster PostgreSQL sendiri-sendiri** di VPS yang sama (`pg_createcluster 18 staging`, port lain; ADR-028 butir 1), bukan dua basis data di satu cluster. Dengan begitu `app_rw` staging dan `app_rw` produksi adalah peran yang berbeda dengan nama sama, `roles.sql` dan migrasi tidak berubah, dan cadangan pgBackRest produksi (salinan fisik seluruh cluster) tidak memuat staging | SEC-09 |
| 6. Uji | Untuk setiap kueri yang disiapkan sqlc, uji lintas rumah tangga yang dibangkitkan: data rumah tangga X tidak pernah kembali bagi sesi rumah tangga Y (§9 berkas 04) | ES §4; LM-09 |

**Tabel global** (tanpa `household_id`): `account`, `credential`, `session`, `refresh_token`, `reauth_grant`, `recovery_token`, `security_event`, `push_token` (milik sesi), `currency` (tabel produk), dan `sentence_pool` (tanpa data rumah tangga). Kebijakan RLS-nya memakai `app.account_id`; `currency` dan `sentence_pool` hanya `SELECT` bagi `app_rw`.

Catatan 14 November 2026 (FS-1 tugas 1.6; putusan `architect` atas rencana 1.6 Q3 dan F05, putusan `security-privacy` atas Q2 dan Q6 di [10.1 §5.1](../10-security-privacy/01-ancaman-dan-akses.md) catatan tanggal itu dan [10.3 SEC-80](../10-security-privacy/03-anak-penyalahgunaan-siklus.md)). Dua tabel global baru milik `household`, keduanya dengan kebijakan `account_id = A` di `USING` dan `WITH CHECK`, dan keduanya di daftar penghapusan akun (SEC-105) walau tanpa foreign key ke `account`: `operator_notice_ack` (`SELECT, INSERT, UPDATE (acknowledged_at)`) dan `consent_record` (`SELECT, INSERT`). Tabel rumah tangga memakai kebijakan lapis 4 (`household_id = current_setting('app.household_id')::uuid`, tanpa `missing_ok`), kecuali:

| Tabel | `USING` | `WITH CHECK` |
|---|---|---|
| `household` | `id = H` (kuncinya `id`) | `id = H` |
| `membership` | `household_id = H′ OR (H′ IS NULL AND account_id = A)` | `household_id = H′` |
| `invitation` | `household_id = H′ OR (H′ IS NULL AND token_hash = T) OR (H′ IS NULL AND invitee_account_id = A)` | `household_id = H′` |
| `signup_verification` (milik `identity`, migrasi `00008`) | `token_hash = T OR setup_token_hash = T OR invitation_token_hash = T OR account_id = A` | `setup_token_hash = T OR invitation_token_hash = T` |

`H` = `current_setting('app.household_id')::uuid`; `H′` = `nullif(current_setting('app.household_id', true), '')::uuid`. Klausa pencarian hanya hidup di transaksi pencarian murni (`H′ IS NULL`) dan hanya di `USING`, sehingga transaksi yang membawa rumah tangga X tidak pernah membaca keanggotaan atau undangan rumah tangga Y (SEC-38 (b)), dan cakupan pencarian tidak dapat menulis. Transaksi pencarian hanya memilih `household_id` dan pengenal baris; sesudahnya perintah berjalan dengan `household_id` dari baris yang ditemukan, tidak pernah dari permintaan. Hak `UPDATE` `signup_verification`: `(token_hash, expires_at, used_at, account_id)`; `account_id` diisi sekali (`WHERE account_id IS NULL`), pengecualian kedua aturan 2 di atas. Diubah 14 November 2026 (FS-1 1.6 PR C tugas 8; putusan `security-privacy`, [10.1 §5.1](../10-security-privacy/01-ancaman-dan-akses.md) catatan 1.6 butir 4 dan [10.4 SEC-158](../10-security-privacy/04-ai-toko-regulasi-perilaku.md) catatan (1); diterima lead menurut FR-FAM-004 catatan (3) dan rencana 1.6 keputusan 15) **[Usulan]**: hak `UPDATE` `app_rw` atas `signup_verification` tepat `(token_hash, email, expires_at, used_at, account_id)`. `email` bukan kolom pemilik, jadi daftar pengecualian aturan 2 tidak berubah. Kueri mengubah `email` hanya pada baris `source = 'invitation'` dan hanya `WHERE account_id IS NULL`.

- **F02** menemukan undangan lewat `token_hash = T`; compare-and-set di transaksi tulis mensyaratkan `version` **dan** `token_hash` sama dengan hash yang dibawa (10.1 §5.1 catatan 1.6 butir 2).
- **F05 `WithdrawAcceptance`** (putusan `architect`, menerima usulan `security-privacy`): penerima bukan anggota, jadi rumah tangganya ditemukan lewat klausa `invitee_account_id = A` di transaksi pencarian `Read(Scope{AccountID})`, dilayani indeks `invitation_invitee` (§5.1). Compare-and-set mensyaratkan `version`, `state = 'accepted'`, dan `invitee_account_id = A`. Alternatif yang ditolak: pemilih `household_id` dari klien (tetap memerlukan pencarian yang sama untuk memeriksanya, jadi hanya menambah medan), dan menyimpan rumah tangga penerima di tabel `identity` (tepi `identity` ke data rumah tangga).

Catatan 13 November 2026 (FS-1 1.4, putusan lead Q2; dikukuhkan `security-privacy` dengan penyesuaian, [10.1 §5.1](../10-security-privacy/01-ancaman-dan-akses.md)). `setup_token` dan `signup_verification` (§5.1) juga tabel global. Masuk, autentikasi token, penyegaran, pemulihan, dan pendaftaran harus menemukan satu baris **sebelum akun diketahui**, jadi kebijakan `app.account_id` saja membuatnya mustahil sebagai `app_rw`. `database.Scope` karena itu membawa dua kunci pencarian lagi, dipasang dengan pola `SET LOCAL` yang sama per transaksi: `app.email` (dibandingkan sebagai `citext`) dan `app.token_hash` (SHA-256 token dalam hex), dengan aturan `nullif` di atas. Keduanya kunci pencarian, bukan hak, dan tunduk pada empat aturan 10.1 §5.1:

1. **Pencarian hanya menemukan, tidak pernah menulis.** Pada tabel yang punya akun, klausa pencarian hanya ada di `USING`; `WITH CHECK` hanya klausa akun. Tabel tanpa akun (`setup_token`, `signup_verification`) menulis hanya lewat token sumbernya.
2. **`UPDATE` per kolom.** `app_rw` tidak pernah mendapat `UPDATE` atas `id`, `account_id`, atau kolom pemilik lain, kecuali `setup_token.account_id` (diisi sekali saat dikonsumsi).
3. **Email hanya untuk pencarian murni.** `database.Write` menolak `Scope` yang membawa `Email`; `database.Read` menolak `Email` bersama `AccountID` atau `HouseholdID`. Penolakan terjadi sebelum `BEGIN`.
4. **Pencarian lewat email hanya memilih `id`**; medan akun lain dibaca sesudahnya dengan cakupan akun.

Satu kebijakan `FOR ALL` per tabel. `A` = `nullif(current_setting('app.account_id', true), '')::uuid`, `E` = `nullif(current_setting('app.email', true), '')::citext`, `T` = `decode(nullif(current_setting('app.token_hash', true), ''), 'hex')`:

| Tabel | `USING` | `WITH CHECK` | Hak `app_rw` |
|---|---|---|---|
| `account` | `id = A OR email = E` | `id = A` | `SELECT, INSERT` |
| `credential` | `account_id = A` | `account_id = A` | `SELECT, INSERT, UPDATE (salt, hash, time_cost, memory_kib, parallelism, updated_at)` |
| `session` | `account_id = A OR token_hash = T OR prev_token_hash = T OR discarded_token_hash = T` (dua klausa terakhir sejak migrasi `00004_web_session`, FS-1 tugas 1.5) | `account_id = A` | `SELECT, INSERT, UPDATE (token_hash, token_issued_at, last_used_at, revoked_at, revoke_reason, prev_token_hash, prev_valid_until, token_confirmed_at, discarded_token_hash)`; tanpa `UPDATE` atas `device_mode` dan `ua_summary` (mode tidak berubah di tengah sesi, SEC-165) |
| `refresh_token`, `recovery_token` | `account_id = A OR token_hash = T` | `account_id = A` | `SELECT, INSERT, UPDATE (used_at)` |
| `reauth_grant` | `account_id = A` | `account_id = A` | `SELECT, INSERT, UPDATE (used_at)`; `cookie_hash` (sejak migrasi `00004_web_session`, FS-1 tugas 1.5) hanya ditulis saat `INSERT`, tanpa `UPDATE` |
| `setup_token` | `token_hash = T` | `token_hash = T` | `SELECT, INSERT, UPDATE (used_at, account_id)` |
| `signup_verification` | `token_hash = T OR setup_token_hash = T OR invitation_token_hash = T OR account_id = A` (dua klausa terakhir sejak migrasi `00008_invitation_signup`, FS-1 tugas 1.6; `account_id = A` hanya menemukan, tidak pernah menulis) | `setup_token_hash = T OR invitation_token_hash = T` | `SELECT, INSERT, UPDATE (token_hash, email, expires_at, used_at, account_id)`; `email` hanya diubah pada baris `source = 'invitation'` yang `WHERE account_id IS NULL` (pendaftaran yang diulang sebelum akun lahir mengganti email-nya), `account_id` diisi sekali |
| `security_event` | `account_id = A` | `account_id = A` | `SELECT, INSERT` (tambah-saja; ADR-016 butir 5) |
| `idempotency_record` | teks di tabel berikut, tanpa perubahan | sama | `SELECT, INSERT` |

Catatan 13 November 2026 (FS-1 tugas 1.5; rencana 1.5 keputusan 12 dan 15): klausa `prev_token_hash = T` dan `discarded_token_hash = T` pada `session` adalah klausa pencarian menurut aturan 1 di atas dan 10.1 §5.1 aturan 7: hanya di `USING`, tidak pernah di `WITH CHECK`. Keduanya ada supaya nilai cookie sebelumnya dan nilai buangan dapat ditemukan sebelum akun diketahui (ADR-030 butir 5 catatan tanggal itu). Autentikasi cookie selalu dua transaksi: pembacaan pencarian dengan `Scope{TokenHash}` yang hanya mengambil `account_id`, lalu `database.Write` dengan `Scope{AccountID, TokenHash}` yang mengunci baris (`FOR UPDATE`), menggolongkan nilai, dan menulis. Kueri pencarian mencakup ketiga kolom hash dan dilayani indeks unik `token_hash` serta dua indeks parsial di §5.1. Uji RLS 10.1 §5.1 aturan 7 mencakup kedua klausa.

`signup_verification`: token tautan (`token_hash = T`) hanya menemukan baris; menulisnya hanya mungkin dengan hash token pemasangan sumbernya (`setup_token_hash = T`), dan `setup_token_hash` tidak pernah dapat diubah. Permintaan tautan ulang (`StartSignUp`, `Write` dengan `T` = hash token pemasangan) mengganti `token_hash` dan `expires_at` baris yang sama. `CompleteSignUp` membaca baris dengan `T` = hash tautan untuk mendapat `setup_token_hash`, lalu menulis dalam satu transaksi dengan `Scope{AccountID: id baru, TokenHash: setup_token_hash}`: `used_at` di `signup_verification`, `used_at` dan `account_id` di `setup_token`, lalu `account`, `credential`, dan sesi. 1.6 menambah klausa sumber undangan di kedua sisi menurut aturan yang sama.

Catatan 14 November 2026 (FS-1 1.6 PR C tugas 8; putusan `security-privacy` yang sama seperti di atas) **[Usulan]**: pada baris bersumber undangan, memulai lagi (`StartSignUp` dengan `T` = hash token undangan) mengganti `token_hash`, `expires_at`, **dan** `email` baris yang sama, hanya selama `account_id IS NULL`; baris token pemasangan tidak pernah berganti `email` (tokennya terikat pada email itu, SEC-158 (b)), dan memulai lagi dengan email lain ditolak seragam. `CompleteSignUp` menulis `signup_verification` dengan compare-and-set `WHERE token_hash = <hash tautan verifikasi yang dibawa> AND account_id IS NULL`, mengisi `used_at` dan `account_id` untuk kedua sumber, dan mengambil email akun dari baris yang lolos penjaga itu, bukan dari pembacaan sebelumnya. Bila nol baris berubah (tautan sudah diganti oleh permintaan ulang di antara baca dan tulis, atau akun sudah lahir), tautan ditolak seragam dan tidak ada akun dibuat. Tanpa penjaga ini, baris dapat berakhir dengan email dan tautan pengganti sementara akun lahir dari email lama. Penjaga terhadap pemegang tautan undangan tetap F04: akun yang lahir bergabung hanya setelah pengundang menyetujuinya dengan email terverifikasi akun ditampilkan (SEC-31).

Tanpa setelan apa pun, setiap tabel di atas kosong dan tidak ada galat. Alur sebelum autentikasi selalu dua transaksi: baca dengan cakupan pencarian, lalu tulis dengan `app.account_id`. `TokenHash` di `Write` hanya diperlukan untuk `setup_token` dan `signup_verification`. `job_runner` dan `lifecycle_rw` tanpa hak atas tabel `identity` sampai tugas yang memerlukannya. Uji wajib, termasuk "pencarian menemukan tetapi tidak menulis" dan hak kolom, ada di 10.1 §5.1; uji hak kolom juga memeriksa bahwa `signup_verification.setup_token_hash` tidak dapat di-`UPDATE`.

**Tabel infrastruktur dengan pemilik campuran** (CS-A-20):

| Tabel | Kebijakan RLS bagi `app_rw` | Catatan |
|---|---|---|
| `idempotency_record` | `scope_id = nullif(current_setting('app.household_id', true), '')::uuid OR scope_id = nullif(current_setting('app.account_id', true), '')::uuid` | `scope_id` adalah rumah tangga, atau akun untuk perintah tingkat akun (§8) |
| `job` | `household_id = nullif(current_setting('app.household_id', true), '')::uuid OR (household_id IS NULL AND nullif(current_setting('app.household_id', true), '') IS NULL)` | Transaksi domain menulis job rumah tangganya sendiri; job sistem (`household_id` NULL, misalnya `ai.pool_build`, `ops.*`) hanya ditulis di luar konteks rumah tangga |
| `job` (pengambilan) | Peran tersendiri `job_runner`: hanya `SELECT` dan `UPDATE (locked_until, attempts, last_error)` atas `job`, dengan kebijakan `USING (true)`; tanpa hak atas tabel lain | Pekerja mengambil job lintas rumah tangga dengan peran ini, lalu menjalankan isinya sebagai `app_rw` dengan `SET LOCAL app.household_id = job.household_id`. `payload` hanya berisi pengenal, jadi melihat semua baris `job` tidak membuka isi rumah tangga |

Catatan 14 November 2026 (FS-1 1.7, putusan `architect` Q1 (b)) **[Usulan]**. Hak `app_rw` atas `job`: `SELECT, INSERT, DELETE`, tanpa `UPDATE`. Transaksi domain mengantrekan; pekerja menghapus job yang berhasil sebagai `app_rw` dengan `Scope{HouseholdID: job.household_id}` (job sistem: `Scope{}`). Coba ulang dan penghentian hanya lewat `locked_until`, `attempts`, dan `last_error`, jadi hak `job_runner` tidak berubah. Kebijakan baris `app_rw` di atas dipasang `FOR ALL TO app_rw` (`USING` dan `WITH CHECK`), dan kebijakan pengambilan `FOR ALL TO job_runner USING (true)`. **Keduanya wajib menyebut perannya:** kebijakan permisif digabung dengan OR, jadi kebijakan `USING (true)` tanpa `TO` membuat `app_rw` melihat job setiap rumah tangga. Gerbang katalog memeriksa kedua kebijakan beserta perannya. Indeks dan keunikan di §5.1, catatan tanggal itu.

Catatan 14 November 2026 (FS-1 1.7, putusan `architect` Q1 (d); dibangun di langkah 7 bersama `sync.prune_feed` dan `idempotency.prune`) **[Usulan]**. Pemangkasan adalah job sistem yang penangannya memakai pool `job_runner`, bukan `app_rw`: ia lintas rumah tangga dan tidak membaca isi, sedangkan `app_rw` tidak punya bacaan lintas rumah tangga. Ini satu-satunya pengecualian atas "isi job dijalankan sebagai `app_rw`". Tambahan bagi `job_runner`, tanpa hak lain:

| Tabel | Hak `job_runner` | Kebijakan `TO job_runner` |
|---|---|---|
| `reader_feed` | `SELECT (household_id, member_id, pos, created_at)`, `DELETE` | `FOR SELECT` dan `FOR DELETE`, keduanya `USING (created_at < now() - interval '90 days')` |
| `idempotency_record` | `SELECT (scope_id, idempotency_key, created_at)`, `DELETE` | sama |

- Kebijakan rumah tangga `reader_feed` (lapis 4, tanpa `missing_ok`) dipasang `TO app_rw` **sejak tabelnya dibuat di 1.7**. Tanpa `TO`, kebijakan itu juga dinilai bagi `job_runner` dan gagal karena `app.household_id` tidak dipasang. Kebijakan `idempotency_record` sudah memakai `missing_ok` dan bernilai salah bagi `job_runner`, jadi tidak perlu diubah.
- Kebijakan adalah batas atas. Kueri membawa batasnya sendiri (`$cutoff` dari `Now`, paling baru 90 hari lalu), sehingga dapat diuji tanpa menunggu, dan baris yang lebih muda dari 90 hari tetap tidak dapat dihapus apa pun kuerinya. Kolom yang terbaca hanya kunci dan waktu; `object_kind`, `object_id`, `command`, `request_hash`, dan `response` tidak.
- Berhalaman (ES §3): `DELETE … WHERE (kunci) IN (SELECT kunci … WHERE created_at < $cutoff ORDER BY created_at LIMIT 1000)`, satu transaksi per putaran, diulang sampai nol baris. Indeks `reader_feed (created_at)` dan `idempotency_record (created_at)` ditambahkan di migrasi itu.
- Pemangkasan menurut umur seharusnya hanya membuang awalan umpan setiap pembaca, tetapi `created_at` = `now()` awal transaksi tidak dijamin searah dengan `pos`. Karena itu deteksi `resync_required` sejak 1.7 tidak bergantung pada urutan jam (berkas 03 §3, catatan tanggal itu): baris yang dikembalikan harus tepat `P + 1, P + 2, …`, dan halaman yang tidak penuh harus berakhir di `last_pos`.

Catatan 13 November 2026 (FS-1 tugas 1.2): pada koneksi pool yang pernah menjalankan `SET LOCAL app.…`, setelan yang tidak dipasang di transaksi berikutnya terbaca `''`, bukan NULL. Setelan khusus lahir sebagai placeholder teks saat pertama kali di-`SET`, dan sesudah transaksinya selesai placeholder itu tetap ada dengan nilai kosong sampai koneksi ditutup. Dokumentasi resmi hanya menyatakan bahwa `current_setting(…, true)` mengembalikan NULL bila setelan tidak ada ([postgresql.org/docs/18/functions-admin.html](https://www.postgresql.org/docs/18/functions-admin.html), diakses 13 November 2026) **[Terverifikasi]**. Perilaku `''` sesudah `SET LOCAL` berasal dari laporan pgsql-bugs 17 Februari 2018 dengan contoh yang dapat diulang ([postgresql.org/message-id/CA+Q86ij0KDCB0G45G509-8q0DNR611gcKG-sSM83GA1EBL7boA@mail.gmail.com](https://www.postgresql.org/message-id/CA%2BQ86ij0KDCB0G45G509-8q0DNR611gcKG-sSM83GA1EBL7boA%40mail.gmail.com), diakses 13 November 2026) **[Terverifikasi]** sebagai bunyi sumber; dokumentasi resmi tidak menyebutnya, jadi uji 1.2 membuktikannya ulang di PostgreSQL 18 yang dipasang. Akibatnya `current_setting(…, true)::uuid` gagal di-cast dan `… IS NULL` tidak pernah benar. Karena itu **setiap kebijakan yang membaca setelan dengan `missing_ok` memakai `nullif(current_setting('app.…', true), '')`**, termasuk kebijakan `app.account_id` tabel global bila membacanya dengan `missing_ok`. Kebijakan lapis 4 sengaja tetap tanpa `missing_ok`: transaksi yang lupa `SET LOCAL` gagal dengan galat, tidak membaca kosong diam-diam. Uji 1.2, pada satu koneksi yang sama: transaksi tingkat akun sesudah transaksi rumah tangga membaca `idempotency_record` akunnya, dan job sistem hanya terbaca di luar konteks rumah tangga.

**Kebocoran lewat unique check di luar RLS** (18 §5.5). PostgreSQL memeriksa unique dan foreign key tanpa memandang RLS. Karena `household_id` memimpin setiap kunci, tabrakan hanya mungkin di dalam rumah tangga yang sama. Satu-satunya keunikan lintas rumah tangga adalah email akun. Tanggapan pendaftaran dan undangan karena itu seragam (SEC-30).

**`full_adult_members`** (SEC-10) adalah satu fungsi Go di modul `household` yang membaca `membership`. Setiap penentuan pihak kesepakatan, tingkat `adults`, dan penerima bawaan memanggilnya. Fungsi SQL tiruan tidak dibuat.

---

## 4. Pengenal

Ringkasan [ADR-017](../adr/ADR-017-pengenal-acak-dan-kursor-per-pembaca.md):

- Setiap pengenal yang keluar ke klien adalah **UUIDv4**: 122 bit acak, dibuat `gen_random_uuid()` di PostgreSQL. UUIDv7, ULID, snowflake, dan bigint berurutan tidak dipakai untuk apa pun yang terbaca anggota (SEC-03; 06 AT-6).
- Untuk menyederhanakan, satu pengenal dipakai per baris, internal maupun eksternal. Urutan tidak pernah diambil dari pengenal: urutan memakai kolom `seq`, `event_date`, atau `server_received_at`.
- **Pengenal ditetapkan server.** Objek yang dibuat offline dirujuk di perangkat lewat `idempotency_key` perintah pembuatnya (`pending_ref`). Server memetakannya ke pengenal yang ia tetapkan, di dalam rumah tangga yang sama. Klien tidak pernah dapat memilih pengenal, sehingga tidak dapat menebak atau menabrakkan pengenal rumah tangga lain.

---

## 5. Sketsa skema

Notasi SQL-semu. Kolom domain ditulis "(07 §n)" dan tidak diulang. Setiap tabel milik rumah tangga punya `household_id uuid NOT NULL` sebagai kolom pertama kunci, dan setiap foreign key-nya memuat `household_id` (SEC-02). Bentuk itu tidak diulang di setiap tabel di bawah. Skema tidak menambah kolom data pribadi di luar yang ditetapkan 06 dan 07; khususnya tidak ada kolom tanggungan, anak, atau hubungan kekerabatan selain yang 06 izinkan (SEC-85; tanggungan Tier 3).

### 5.1 Infrastruktur

```sql
-- Tabel produk: eksponen dari produk, bukan dari ISO 4217 mentah (ADR-008).
currency (
  code        char(3) PRIMARY KEY,       -- ISO 4217
  exponent    smallint NOT NULL CHECK (exponent BETWEEN 0 AND 4)
)  -- baris awal: ('IDR', 0)

household (
  id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  currency    char(3) NOT NULL DEFAULT 'IDR' REFERENCES currency(code),
  time_zone   text NOT NULL,              -- IANA, waktu rumah tangga (06 §2.4)
  version     bigint NOT NULL DEFAULT 0,  -- compare-and-set
  ...,                                    -- 06 §2.4, 07 §2.1
  UNIQUE (id, currency)                   -- sasaran FK mata uang (§7)
)

account (
  id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  email       citext UNIQUE NOT NULL,
  time_zone   text NULL,                  -- 06 §2.4: kosong = zona rumah tangga
  locale      text NOT NULL DEFAULT 'id',
  created_at  timestamptz NOT NULL,       -- catatan 13 November 2026 (FS-1 1.4)
  signup_via  text NOT NULL CHECK (signup_via IN ('setup_token','invitation'))
                                          -- SEC-158: hanya akun 'setup_token' yang boleh membuat rumah tangga
)

-- Satu baris per perangkat atau browser yang masuk (SEC-15, SEC-161, SEC-165; ADR-030 butir 5).
-- Native: token akses berumur pendek + refresh_token. Web: satu token di cookie __Host-, diputar di sini.
session (
  id               uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id       uuid NOT NULL REFERENCES account(id),
  client_kind      text NOT NULL CHECK (client_kind IN ('ios','android','web')),
  device_mode      text NOT NULL DEFAULT 'personal' CHECK (device_mode IN ('personal','shared')),
  ua_summary       text NOT NULL,         -- untuk daftar sesi (FR-FAM-022). Catatan 13 November 2026 (FS-1 tugas 1.5):
                                          -- web "<browser> · <sistem>" dari daftar tetap, tanpa versi, ≤ 64 karakter,
                                          -- misalnya "Safari · iOS" (aturan di 10.5 SEC-161 (e)); User-Agent mentah
                                          -- tidak pernah disimpan; native ''
  token_hash       bytea NOT NULL UNIQUE, -- SHA-256 token akses (native) atau token cookie (web)
  prev_token_hash  bytea NULL,            -- web: nilai sebelum rotasi
  prev_valid_until timestamptz NULL,      -- web: +60 detik. Catatan 13 November 2026 (FS-1 tugas 1.5): pakai ulang
                                          -- sesudahnya mencabut sesi hanya bila token_confirmed_at terisi; bila
                                          -- kosong, terbit ulang tanpa pencabutan (ADR-030 butir 5; SEC-161 (c))
  token_confirmed_at   timestamptz NULL,  -- web, sejak 1.5: diisi permintaan pertama yang membawa nilai sekarang
                                          -- (terbukti sampai di browser); dikosongkan setiap rotasi
  discarded_token_hash bytea NULL,        -- web, sejak 1.5: hash nilai terakhir yang dibuang saat terbit ulang;
                                          -- dibawa kemudian → cabut 'reuse_detected' (SEC-161 (c) aturan (3))
  client_version   text NOT NULL,         -- dibandingkan dengan versi minimum per client_kind (ADR-028 butir 5)
  created_at       timestamptz NOT NULL,
  last_used_at     timestamptz NOT NULL,  -- dasar batas diam
  idle_timeout     interval NOT NULL,     -- web pribadi 30 hari; web bersama 30 menit; native 30 hari [Usulan] (1.4 Q5)
  absolute_expiry  timestamptz NOT NULL,  -- web pribadi +180 hari; web bersama +12 jam; native +180 hari [Usulan] (1.4 Q5)
  token_issued_at  timestamptz NOT NULL,  -- 1.4: dasar umur token akses 15 menit (SEC-15); 1.5: dasar rotasi cookie 15 menit
  revoked_at       timestamptz NULL,
  revoke_reason    text NULL CHECK (revoke_reason IN ('user','reuse_detected','exit','purge','recovery',
                                                          'password_changed')),
                                          -- 'recovery' sejak 1.4, 'password_changed' sejak migrasi
                                          -- 00003_credential_reset (1.4 PR B): keduanya mencabut semua sesi
                                          -- lain (SEC-24); berakhir karena umur tidak dicatat di sini
  CHECK (client_kind = 'web' OR device_mode = 'personal')   -- mode bersama hanya di web
)
CREATE INDEX session_account_live ON session (account_id) WHERE revoked_at IS NULL;  -- cabut semua (SEC-24)
-- Catatan 13 November 2026 (FS-1 tugas 1.5, migrasi 00004_web_session, expand saja): pencarian nilai cookie
-- sebelumnya dan buangan (§3).
CREATE INDEX session_prev_token ON session (prev_token_hash) WHERE prev_token_hash IS NOT NULL;
CREATE INDEX session_discarded_token ON session (discarded_token_hash) WHERE discarded_token_hash IS NOT NULL;
-- Pencabutan fail-closed (SEC-24, 1.4 PR B): pemulihan atau ganti kata sandi mencabut SEMUA sesi akun
-- dengan revoked_at IS NULL (ChangePassword mempertahankan sesinya sendiri), termasuk yang sudah
-- berakhir karena umur. Baris security_event 'session_revoked' ditulis hanya untuk sesi yang saat itu
-- masih hidup (dalam batas diam dan batas mutlak).
-- Sesi berakhir tidak dapat dihidupkan lagi (SEC-15): touch (last_used_at) dan rotasi token akses
-- menolak sesi yang dicabut, melewati batas diam, atau melewati absolute_expiry di dalam SQL-nya
-- sendiri (syarat WHERE), dan dijawab SESSION_REVOKED atau SESSION_EXPIRED.

-- Catatan 13 November 2026 (FS-1 1.4, putusan lead Q3 dan Q8): tabel identity lain. Semua token
-- 32 byte acak; yang disimpan hanya SHA-256-nya. Kebijakan RLS di §3 ("Tabel global").
credential (                              -- SEC-11; 09.4 §2
  account_id   uuid PRIMARY KEY REFERENCES account(id),
  salt         bytea NOT NULL,
  hash         bytea NOT NULL,
  time_cost    int NOT NULL,              -- parameter Argon2id disimpan per baris agar dapat dinaikkan
  memory_kib   int NOT NULL,
  parallelism  smallint NOT NULL,
  updated_at   timestamptz NOT NULL
)

refresh_token (                           -- SEC-15: satu keluarga per sesi, diputar setiap dipakai
  token_hash   bytea PRIMARY KEY,
  session_id   uuid NOT NULL REFERENCES session(id),
  account_id   uuid NOT NULL REFERENCES account(id),
  created_at   timestamptz NOT NULL,
  used_at      timestamptz NULL           -- dipakai lagi sesudah terisi: sesi dicabut 'reuse_detected'
)

reauth_grant (                            -- SEC-13: terikat sesi dan tujuan, 5 menit, satu tindakan
  id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  session_id   uuid NOT NULL REFERENCES session(id),
  account_id   uuid NOT NULL REFERENCES account(id),
  purpose      text NOT NULL CHECK (purpose IN ('change_password')),  -- tujuan lain ditambah tugasnya
  created_at   timestamptz NOT NULL,
  expires_at   timestamptz NOT NULL,
  used_at      timestamptz NULL,
  cookie_hash  bytea NULL                 -- Catatan 13 November 2026 (FS-1 tugas 1.5, migrasi 00004_web_session):
                                          -- hash nilai cookie yang diterbitkan di jawaban autentikasi ulang web;
                                          -- hibah hanya berlaku bagi permintaan yang membawa nilai itu. NULL untuk
                                          -- sesi native. Ditulis sekali saat INSERT (10.5 SEC-161 (c) aturan (7))
)
CREATE INDEX reauth_grant_open ON reauth_grant (session_id, purpose) WHERE used_at IS NULL;

recovery_token (                          -- SEC-24: tautan email sekali pakai, 30 menit
  token_hash   bytea PRIMARY KEY,
  account_id   uuid NOT NULL REFERENCES account(id),
  created_at   timestamptz NOT NULL,
  expires_at   timestamptz NOT NULL,
  used_at      timestamptz NULL           -- reset kredensial (ChangePassword, CompleteRecovery) mengisi
                                          -- used_at semua token terbuka akun itu: tautan lama mati
)
CREATE INDEX recovery_token_account_open ON recovery_token (account_id) WHERE used_at IS NULL;
                                          -- migrasi 00003 (FS-1 1.4 PR B)

setup_token (                             -- SEC-158 (b); 09.4 §2: terikat satu email, 7 hari
  token_hash   bytea PRIMARY KEY,
  email        citext NOT NULL,
  created_at   timestamptz NOT NULL,
  expires_at   timestamptz NOT NULL,
  used_at      timestamptz NULL,
  account_id   uuid NULL REFERENCES account(id)   -- akun yang lahir darinya
)

-- Tautan verifikasi email pendaftaran, umum untuk kedua jalur SEC-158 (Q8). Akun baru dibuat saat
-- tautan dibuka dan kata sandi dipasang, sehingga tidak ada akun "belum terverifikasi".
-- 1.4 hanya sumber 'setup_token'; 1.6 memperluas CHECK sumber ke 'invitation' dan menambah
-- kolom rujukan undangannya (perubahan expand, §10).
signup_verification (
  id                uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  token_hash        bytea NOT NULL UNIQUE,              -- SHA-256 token tautan
  source            text NOT NULL CHECK (source IN ('setup_token')),
  setup_token_hash  bytea NULL UNIQUE REFERENCES setup_token(token_hash),
                                                        -- satu tautan berlaku per token pemasangan;
                                                        -- permintaan ulang mengganti token_hash dan expires_at
  email             citext NOT NULL,                    -- alamat yang diverifikasi
  created_at        timestamptz NOT NULL,
  expires_at        timestamptz NOT NULL,               -- 30 menit (1.4 Q4) [Usulan]
  used_at           timestamptz NULL,
  CHECK ((source = 'setup_token') = (setup_token_hash IS NOT NULL))
)

security_event (                          -- jejak keamanan; tambah-saja (ADR-016 butir 5); SEC-16, SEC-39 (b)
  id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id   uuid NOT NULL REFERENCES account(id),
  session_id   uuid NULL REFERENCES session(id),
  kind         text NOT NULL CHECK (kind IN
                 ('sign_in','session_revoked','reauthentication','password_changed','recovery')),
  detail       text NULL,                 -- session_revoked: user | reuse_detected | recovery | password_changed;
                                          -- reauthentication: tujuannya. Tidak pernah token, kata sandi, atau email
  occurred_at  timestamptz NOT NULL
)

idempotency_record (
  scope_id         uuid NOT NULL,          -- household_id, atau account_id untuk perintah tingkat akun
  idempotency_key  uuid NOT NULL,          -- dibuat klien, satu per niat (ADR-005)
  command          text NOT NULL,
  request_hash     bytea NOT NULL,         -- SHA-256 isi perintah yang dikanonkan
  response         bytea NOT NULL,         -- protobuf hasil pertama
  created_at       timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (scope_id, idempotency_key)
)

object_event (                            -- 06 §2.5: Peristiwa per objek
  household_id  uuid, object_kind text, object_id uuid,
  seq           bigint NOT NULL,          -- naik tanpa celah per objek (06 S-5)
  server_time   timestamptz NOT NULL DEFAULT now(),
  ...,                                    -- 06 §2.5
  PRIMARY KEY (household_id, object_kind, object_id, seq)
)

-- Catatan 14 November 2026 (FS-1 tugas 1.6, migrasi 00005_household; putusan architect atas rencana 1.6
-- Q1, Q3, Q8, Q14). Bentuk tabel rumah tangga, semuanya milik household (§1) dan menggantikan sketsa
-- household dan object_event di atas. Kebijakan RLS lapis 4 (§3), kecuali household, membership, dan
-- invitation (tabel di §3, catatan 14 November 2026). Tanpa FK dari tabel rumah tangga ke account: FK itu tidak dapat diawali
-- household_id (§3 lapis 2) dan melintasi batas modul. currency pada approval_authority dan household_rule
-- NULL bila jumlahnya NULL: FK komposit dengan satu kolom NULL tidak diperiksa ("MATCH SIMPLE allows any of
-- the foreign key columns to be null; if any of them are null, the row is not required to have a match",
-- postgresql.org/docs/18/sql-createtable.html, diakses 14 November 2026) [Terverifikasi], sehingga baris
-- tanpa jumlah tidak menghalangi ganti mata uang sebelum ada baris uang (FR-FAM-001/AC-8, langkah 2; 07 T-45).
-- Catatan 14 November 2026 (putusan lead, tinjauan Tugas 1, sudah di 00005_household.sql): karena itu juga,
-- household_rule dan object_event mendapat FK biasa (household_id) → household (id), sebab FK kompositnya
-- tidak diperiksa saat currency atau actor NULL (baris pilot biasa, setiap Peristiwa invitee dan system);
-- FK (household_id, currency) memakai ON UPDATE RESTRICT (§5.2 butir 4).
household (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(), name text NOT NULL, time_zone text NOT NULL,
  currency char(3) NOT NULL DEFAULT 'IDR' REFERENCES currency (code),
  income_source_market text NOT NULL DEFAULT 'ID' CHECK (income_source_market IN ('ID')),  -- 07 §3.1
  version bigint NOT NULL DEFAULT 0, created_at timestamptz NOT NULL,
  UNIQUE (id, currency)
)
membership (                              -- 06 §3.1 peran dan izin. Sampai F06 dibangun: baris ada = aktif;
                                          -- F06 menambah state DEFAULT 'active' (expand)
  household_id uuid NOT NULL REFERENCES household (id),
  member_id    uuid NOT NULL DEFAULT gen_random_uuid(),  -- pengenal internal anggota; tidak pernah diputus
  account_id   uuid NULL,                 -- pemetaan identitas anggota (07 §8.2): diputus = NULL
  display_name text NULL,                 -- FR-FAM-006; diputus bersama account_id
  role         text NOT NULL CHECK (role IN ('household_owner')),
  permission_set text NOT NULL CHECK (permission_set IN ('two_equal_adults')),
  joined_at    timestamptz NOT NULL,
  PRIMARY KEY (household_id, member_id), UNIQUE (household_id, account_id)
)
CREATE UNIQUE INDEX membership_one_household ON membership (account_id) WHERE account_id IS NOT NULL;
                                          -- keanggotaan akun sendiri (§3 lapis 1) dan aturan produk pilot
                                          -- "satu rumah tangga per akun"; lihat catatan balapan di bawah.
                                          -- Catatan 14 November 2026 (putusan lead, tinjauan Tugas 1):
                                          -- tidak ada indeks membership_account; migrasi F-FAM-009 yang
                                          -- menghapus indeks unik ini menambah indeks biasa (account_id)
approval_authority (                      -- 06 §3.1 wewenang persetujuan
  household_id, member_id, request_type text NOT NULL CHECK (request_type IN ('all')),
  limit_amount bigint NULL, currency char(3) NULL,   -- §7: jumlah selalu berpasangan; keduanya NULL = tanpa batas
  valid_from timestamptz NOT NULL, valid_until timestamptz NULL,
  CHECK ((limit_amount IS NULL) = (currency IS NULL)),
  PRIMARY KEY (household_id, member_id, request_type, valid_from),
  FOREIGN KEY (household_id, member_id) REFERENCES membership,
  FOREIGN KEY (household_id, currency) REFERENCES household (id, currency) ON UPDATE RESTRICT   -- §5.2 butir 4
)
household_rule (                          -- 06 §3.1 aturan rumah tangga; tidak dapat diubah, versi baru = baris baru
  household_id, version int NOT NULL,
  preset_id text NOT NULL CHECK (preset_id IN ('two_equal_adults')),
  request_threshold bigint NULL, currency char(3) NULL,
  agreement_mode text NOT NULL DEFAULT 'all' CHECK (agreement_mode IN ('all')),
  created_at timestamptz NOT NULL,
  CHECK ((request_threshold IS NULL) = (currency IS NULL)),
  PRIMARY KEY (household_id, version),
  FOREIGN KEY (household_id) REFERENCES household (id),   -- komposit di bawah tidak diperiksa saat currency NULL
  FOREIGN KEY (household_id, currency) REFERENCES household (id, currency) ON UPDATE RESTRICT   -- §5.2 butir 4
)  -- approval_mode ditambah request (langkah 4, expand) bersama resolve_approvers: belum ada aturan yang membacanya
relationship (                            -- 06 §3.1 hubungan: tabel ada, tanpa baris di pilot
  household_id, member_id, other_member_id, kind text NOT NULL,
  PRIMARY KEY (household_id, member_id, other_member_id),
  FOREIGN KEY (household_id, member_id) REFERENCES membership,
  FOREIGN KEY (household_id, other_member_id) REFERENCES membership
)  -- app_rw hanya SELECT: tidak ada penulis sampai Tier 3
invitation (                              -- 06 §3.6 F01 sampai F05; aturan token 10.1 SEC-28 (catatan 1.6, Q7)
  household_id, id uuid DEFAULT gen_random_uuid(), inviter uuid NOT NULL,
  token_hash bytea NULL,                  -- NULL sejak F01 sampai IssueInvitationLink pertama; setiap terbitan
                                          -- ulang mengganti hash (tautan lama mati); F02 mengosongkannya (sekali
                                          -- pakai). Token 256 bit hanya di jawaban IssueInvitationLink
  state text NOT NULL CHECK (state IN ('open','accepted','confirmed','cancelled','rejected','withdrawn','expired')),
                                          -- 06: open = undangan_terbuka, accepted = menunggu_konfirmasi_pengundang,
                                          -- confirmed = F04 selesai; empat sisanya = gugur dengan sebabnya
  invitee_account_id uuid NULL,           -- pemetaan identitas penerima, seperti membership.account_id
  invitee_display_name text NULL,         -- catatan 14 November 2026 (FS-1 tugas 1.6, PR B, putusan lead):
                                          -- penarikan (F05 withdraw_acceptance, juga CloseAsWithdrawn saat
                                          -- penerima sudah bergabung di tempat lain) mengosongkan keduanya di
                                          -- compare-and-set yang sama, karena FR-FAM-004/AC-5 ("tanpa jejak data
                                          -- di rumah tangga itu selain Peristiwa undangan"); Peristiwanya tetap
                                          -- berpelaku invitee. Penolakan, pembatalan, dan gugur karena waktu
                                          -- mempertahankannya sampai penghapusan akun (langkah 6)
  created_at timestamptz NOT NULL, expires_at timestamptz NOT NULL,
  closed_at timestamptz NULL,             -- gugur karena waktu: = expires_at, bukan saat ditulis
  version bigint NOT NULL DEFAULT 0,
  PRIMARY KEY (household_id, id),
  FOREIGN KEY (household_id, inviter) REFERENCES membership
)
CREATE INDEX invitation_token ON invitation (token_hash) WHERE token_hash IS NOT NULL;
CREATE UNIQUE INDEX invitation_live ON invitation (household_id) WHERE state IN ('open','accepted');
CREATE UNIQUE INDEX invitation_invitee ON invitation (invitee_account_id) WHERE state = 'accepted';
                                          -- F05 penerima (§3) dan satu penerimaan terbuka per akun
-- Catatan 14 November 2026 (FS-1 tugas 1.6, putusan architect atas rencana 1.6 keputusan 20): balapan
-- "sudah anggota di mana pun" (Q5 (d), Q18). Pemeriksaan keanggotaan berjalan di transaksi pencarian, dan
-- transaksi tulis yang membawa rumah tangga tidak dapat mengulangnya (§3). Penegaknya karena itu dua indeks
-- unik di atas: PostgreSQL memeriksa unique tanpa memandang RLS (§3 "Kebocoran lewat unique check"), jadi
-- keduanya berlaku lintas rumah tangga. Keduanya pengecualian tercatat atas §3 lapis 2 (kunci tidak diawali
-- household_id); kuncinya akun pemanggil sendiri. Keduanya aturan produk pilot dan dihapus (contract)
-- bersama F-FAM-009 (Tier 3); saat F06 menambah membership.state, predikat membership_one_household menjadi
-- state yang masih menghitung keanggotaan (aktif dan keluar_tenggang).
--   * CreateHousehold ×2, CreateHousehold + F04, F02 ×2 dari satu akun: pelanggaran unik di transaksi
--     pemanggil → galat modul ALREADY_HOUSEHOLD_MEMBER (17), tidak pernah Internal.
--   * F04 ketika penerima sudah menjadi anggota rumah tangga lain sesudah F02 (CreateHousehold yang balapan
--     dengan F02; tidak dapat dicegah indeks undangan saja, karena konfliknya lintas dua tabel):
--     INSERT membership … ON CONFLICT (account_id) WHERE account_id IS NOT NULL DO NOTHING ("index_predicate
--     … Used to allow inference of partial unique indexes", postgresql.org/docs/18/sql-insert.html, diakses
--     14 November 2026) [Terverifikasi]. Bila tidak ada baris tersisip (diubah 14 November 2026 atas
--     SEC-38 (c), 10.1 §10, diterima lead): transaksi F04 dibatalkan, lalu lapisan API menjalankan satu
--     database.Write lagi yang menutup undangan **persis seperti penarikan F05 oleh penerima**: to_state
--     'withdrawn', actor_kind 'invitee', perintah sama dengan F05, tanpa rule, idempotency_key UUIDv4 acak
--     baru (penarikan sungguhan selalu membawa kunci, jadi NULL akan membedakannya);
--     compare-and-set atas version dan state 'accepted' (kalah → tidak ada tulisan). F04 menjawab
--     STATE_CHANGED, sama dengan kalah balapan dari penarikan sungguhan, dan tidak tercatat di
--     idempotency_record, sehingga kiriman ulang menjawab STATE_CHANGED dari state 'withdrawn'. Peristiwa itu
--     sengaja tidak membedakan sebabnya: pelaku tercatat jenis invitee walau penerima tidak memerintahkannya,
--     pengecualian tercatat atas INV-A01 demi SEC-38 (c), karena sebab yang sebenarnya membuka bahwa penerima
--     bergabung ke rumah tangga lain. INV-A48 (b) tetap (tidak ada transisi menuju aktif).
--   * Kode 17 di jalur F02 (dikukuhkan security-privacy, 10.1 §10): transaksi F02 tidak menulis apa pun saat
--     menjawab 17, dan percobaan itu dihitung pembatas tokenUse.
--   * Uji: setiap pasangan di atas dengan dua goroutine di balik penghalang → tepat satu keanggotaan atau satu
--     penerimaan terbuka per akun; F04 sesudah penerima membuat rumah tangga → STATE_CHANGED, undangan
--     'withdrawn', tanpa membership, dan baris object_event-nya sama medan demi medan (selain pengenal dan
--     waktu) dengan penarikan F05 sungguhan; F02 yang dijawab 17 → nol baris baru dan satu hitungan tokenUse.
-- Catatan 14 November 2026 (FS-1 tugas 1.6, putusan architect atas X3; mengubah putusan Q8, 06 §3.6
-- "Gugur karena waktu"): undangan 'accepted' yang lewat masa berlakunya. Predikat indeks tidak dapat
-- memuat waktu sekarang (baris indeks hanya dinilai ulang saat baris ditulis), jadi gugur harus ditulis.
-- Aturannya: **jalur tanpa token menulis gugur lebih dulu; jalur token tidak pernah menulis.**
--   * Penulisan gugur: database.Write tersendiri dengan Scope{HouseholdID: household_id baris yang ditemukan,
--     AccountID}, compare-and-set atas version, state IN ('open','accepted'), expires_at <= now →
--     'expired', actor_kind 'system', rule 'invitation.expired', closed_at = expires_at, dan server_time
--     Peristiwanya = expires_at (saat gugur berlaku; waktu penulisan tidak disimpan, supaya tidak membuka kapan
--     penerima atau pengundang memakai aplikasi). Kalah compare-and-set → tidak ada tulisan. Sesudahnya
--     perintah berjalan di transaksinya sendiri. Catatan 14 November 2026 (FS-1 tugas 1.6, PR B, putusan
--     lead): command Peristiwa ini 'expire_invitation' (kata kerja snake_case seperti 'create_household',
--     'accept_invitation'), dengan rule 'invitation.expired' dan actor_kind 'system'.
--   * CreateHousehold dan F02 (sebelum token dibaca, sesudah tahan pemberitahuan): pencarian
--     invitee_account_id = A menemukan penerimaan akun sendiri; yang lewat ditulis gugur lalu tidak
--     menghalangi; yang belum lewat → 17. Urutan F02: tahan, keanggotaan dan penerimaan sendiri (termasuk
--     penulisan gugur ini), pernyataan usia, token. Waktunya hanya bergantung pada keadaan akun pemanggil.
--   * WithdrawAcceptance atas penerimaan sendiri yang lewat: gugur ditulis, jawaban STATE_CHANGED.
--   * F01, F03, F04, IssueInvitationLink atas undangan rumah tangga sendiri yang lewat: gugur ditulis, lalu
--     F01 berlanjut dan yang lain menjawab STATE_CHANGED.
--   * F02 dan PreviewInvitation lewat token atas undangan yang lewat: tidak menulis, penolakan token seragam
--     (alasan Q8 tetap).
--   * Uji: penerimaan lewat 7 hari tanpa F04 → penerima dapat F02 undangan lain dan CreateHousehold;
--     WithdrawAcceptance → STATE_CHANGED dan undangan 'expired'; server_time Peristiwa = expires_at.
--   * security-privacy memeriksa: server_time = expires_at cukup menyembunyikan waktu aktivitas dari
--     rumah tangga pengundang; dan di 1.7 baris reader_feed untuk gugur yang ditulis belakangan tidak
--     membuka waktu itu lewat saat munculnya di umpan.
object_event (                            -- 06 §2.5, INV-A01; tambah-saja (ADR-016 butir 5)
  household_id,
  object_kind  text NOT NULL CHECK (object_kind IN ('household','invitation')),  -- diperluas tiap modul penulis
  object_id    uuid NOT NULL,
  seq          bigint NOT NULL,           -- = version objek sesudah transisi (ADR-016 butir 7 "seq = versi
                                          -- objek"); Peristiwa pembuatan seq 0
  actor_kind   text NOT NULL CHECK (actor_kind IN ('member','invitee','system')),
  actor        uuid NULL,                 -- member_id
  rule         text NULL,                 -- system: pengenal aturan pemicu, mis. 'invitation.expired'
  command      text NOT NULL, from_state text NULL, to_state text NOT NULL,
                                          -- catatan 14 November 2026 (putusan lead, tinjauan PR A 1.6): command =
                                          -- kata kerja snake_case 06 §3.6 (pembuatan rumah tangga
                                          -- 'create_household' dengan to_state 'created'; undangan 'invite',
                                          -- 'accept_invitation', 'cancel_invitation', …); state = nama state
                                          -- bahasa Inggris ('open', 'accepted', …)
                                          -- catatan 14 November 2026 (FS-1 tugas 1.6, PR B, putusan lead): F04
                                          -- menulis, selain Peristiwa 'confirmed' undangan, Peristiwa rumah tangga
                                          -- 'member_joined' (command 'confirm_invitee', from_state NULL, seq =
                                          -- version rumah tangga yang dinaikkan, idempotency_key sama dengan
                                          -- Peristiwa undangan). Pelakunya anggota yang bergabung; pengundang yang
                                          -- menjalankan F04 adalah pelaku Peristiwa 'confirmed' undangan di
                                          -- transaksi yang sama dengan kunci yang sama, jadi pasangan Peristiwa itu
                                          -- lengkap tanpa kolom tambahan
  idempotency_key uuid NULL, server_time timestamptz NOT NULL,
  CHECK ((actor_kind = 'member') = (actor IS NOT NULL)),
  CHECK ((actor_kind = 'system') = (rule IS NOT NULL)),
  CHECK (actor_kind <> 'invitee' OR object_kind = 'invitation'),
  PRIMARY KEY (household_id, object_kind, object_id, seq),
  FOREIGN KEY (household_id) REFERENCES household (id),   -- komposit di bawah tidak diperiksa saat actor NULL
  FOREIGN KEY (household_id, actor) REFERENCES membership (household_id, member_id)
)  -- terms_version (INV-A01) ditambah request bersama request_terms
operator_notice_ack (                     -- tabel global (§3); SEC-34; FR-FAM-019/AC-3 riwayat
  account_id uuid NOT NULL, variant text NOT NULL CHECK (variant IN ('couple','family','general')),
  version int NOT NULL,
  text_sha256 bytea NOT NULL,             -- SHA-256 teks yang sudah diisi nama pengelola: identitas versi bagi akun
  issued_at timestamptz NOT NULL, acknowledged_at timestamptz NULL,
  issue_seq int NOT NULL,                 -- catatan 14 November 2026 (putusan security-privacy, 10.1 sebelum
                                          -- §11.3): satu baris per terbitan, bukan per identitas. Terbit bila
                                          -- identitas sekarang ≠ identitas baris issue_seq terbesar akun itu
                                          -- (A → B → A menerbitkan A lagi); INSERT dengan issue_seq = terbesar + 1
                                          -- ON CONFLICT DO NOTHING, jadi dua pertama-terlihat serentak menulis satu
                                          -- baris dan satu jejak. Ditahan selama baris terbesar belum diakui;
                                          -- pengakuan hanya atas baris terbesar dengan hash yang sama
  PRIMARY KEY (account_id, issue_seq)
)  -- hash atas penyandian berawalan panjang dari varian, nomor, judul, dan setiap paragraf beserta tanda
   -- fiturnya; pengakuan membawa hash yang ditampilkan (10.1 catatan 14 November 2026 sesudah SEC-159, Q3 (a), (b))
consent_record (                          -- tabel global (§3); SEC-80, SEC-109 (a); tanpa tanggal lahir
  account_id uuid NOT NULL, kind text NOT NULL CHECK (kind IN ('adult_age')),
  text_version int NOT NULL, recorded_at timestamptz NOT NULL,
  PRIMARY KEY (account_id, kind, text_version)
)  -- akun yang sudah punya baris adult_age tidak ditanya lagi saat versi teks berganti

-- Catatan 14 November 2026 (FS-1 tugas 1.6, migrasi 00008, expand): signup_verification (identity) mendapat
-- sumber undangan. Kebijakan dan hak di §3.
--   source CHECK (source IN ('setup_token','invitation')),
--   invitation_token_hash bytea NULL UNIQUE,   -- tidak pernah dapat diubah
--   account_id uuid NULL,                      -- akun yang lahir darinya; diisi sekali, untuk kedua sumber
--   CHECK (tepat satu dari setup_token_hash dan invitation_token_hash terisi, sesuai source)
--   CREATE UNIQUE INDEX signup_verification_account ON signup_verification (account_id)
--     WHERE account_id IS NOT NULL;            -- satu baris per akun yang lahir
-- Diubah 14 November 2026 (FS-1 1.6 PR C tugas 8, putusan security-privacy) [Usulan]: account_id tanpa FK
-- ke account; syaratnya, penghapusan akun (FR-FAM-023, SEC-105) menghapus baris signup_verification menurut
-- account_id karena baris itu memuat email (tugas kemudian, bukan tugas 8). email diganti hanya pada baris
-- 'invitation' selama account_id IS NULL (§3).

reader_feed (                             -- ADR-017: umpan per pembaca
  household_id uuid, member_id uuid,
  pos          bigint NOT NULL,           -- naik tanpa celah per (rumah tangga, anggota)
  object_kind  text NOT NULL, object_id uuid NOT NULL,
  change       text NOT NULL CHECK (change IN ('upsert','purge')),
  created_at   timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (household_id, member_id, pos)
)
reader_feed_counter (household_id uuid, member_id uuid, last_pos bigint NOT NULL,
                     PRIMARY KEY (household_id, member_id))

job (
  id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  household_id uuid NULL,
  kind        text NOT NULL,
  run_at      timestamptz NOT NULL,
  unique_key  text NULL,                  -- pekerjaan yang sama tidak diantrekan dua kali
  payload     bytea NOT NULL,             -- hanya pengenal, tidak pernah isi (SEC-44)
  attempts    int NOT NULL DEFAULT 0,
  locked_until timestamptz NULL,
  last_error  text NULL,                  -- kode galat saja
  UNIQUE (kind, unique_key)
)
CREATE INDEX job_due ON job (run_at) WHERE locked_until IS NULL;
-- Catatan 14 November 2026 (FS-1 1.7, putusan architect Q1 (b)); migrasi 1.7 menggantikan UNIQUE (kind, unique_key):
--   CHECK (kind IN ('household.expire_invitation'))   -- diperluas per jenis, seperti object_event.object_kind
--   CHECK (attempts >= 0)
--   CREATE UNIQUE INDEX job_unique ON job (household_id, kind, unique_key) NULLS NOT DISTINCT
--     WHERE unique_key IS NOT NULL;
--     -- household_id memimpin (§3 lapis 2): job rumah tangga lain tidak pernah bertabrakan dengan, atau
--     -- dilewati diam-diam oleh, ON CONFLICT DO NOTHING rumah tangga ini. NULLS NOT DISTINCT menjaga job
--     -- sistem (household_id NULL) tetap unik; WHERE membebaskan job tanpa unique_key. Enqueue:
--     -- INSERT … ON CONFLICT (household_id, kind, unique_key) WHERE unique_key IS NOT NULL DO NOTHING.
--     -- NULLS [NOT] DISTINCT dan WHERE dalam satu CREATE UNIQUE INDEX: postgresql.org/docs/18/sql-createindex.html,
--     -- diakses 14 November 2026 [Terverifikasi]
--   CREATE INDEX job_locked ON job (locked_until) WHERE locked_until IS NOT NULL;  -- coba ulang dan sewa habis
```

### 5.2 Inti uang (dari 07 §1 sampai §2)

```sql
entry (
  household_id uuid, id uuid DEFAULT gen_random_uuid(),
  type         text NOT NULL,             -- 07 §1: jenis entri
  event_date   date NOT NULL,
  recorded_at  timestamptz NOT NULL DEFAULT now(),   -- cap waktu penuh tidak dikirim ke perangkat; jalur baca
                                                     -- mengirim hanya tanggalnya (`recorded_on`, tanpa jam)
                                                     -- kepada setiap pembaca entri (AT-6; dokumen 10; 03 §2)
  actor        uuid NULL,                 -- NULL = system (exit_closure; INV-25)
  corrects     uuid NULL,                 -- entri asal pembalik (INV-18)
  idempotency_key uuid NOT NULL,          -- disimpan permanen (§8)
  currency     char(3) NOT NULL,
  PRIMARY KEY (household_id, id),
  UNIQUE (household_id, idempotency_key),
  UNIQUE (household_id, corrects),        -- paling banyak satu pembalik (INV-18)
  FOREIGN KEY (household_id, currency) REFERENCES household (id, currency)
)

line (
  household_id uuid, entry_id uuid, line_no smallint,
  owner_kind   text NOT NULL CHECK (owner_kind IN ('shared','member')),
  owner_member uuid NULL,
  role         text NOT NULL,             -- position | cause | allocation (07 §1)
  position_kind text NULL, position_id uuid NULL,
  cause        text NULL,                 -- 07 §1: sebab
  delta        bigint NULL,               -- satuan terkecil (ADR-008)
  effect       bigint NULL,
  currency     char(3) NOT NULL,
  PRIMARY KEY (household_id, entry_id, line_no),
  FOREIGN KEY (household_id, entry_id) REFERENCES entry (household_id, id),
  FOREIGN KEY (household_id, currency) REFERENCES household (id, currency),
  CHECK ((owner_kind = 'member') = (owner_member IS NOT NULL))
)
CREATE INDEX line_by_position ON line (household_id, position_kind, position_id);

entry_detail (...)   -- 07 §2: dapat dihapus (ADR-007); FK ke entry, bukan sebaliknya
identity_map (...)   -- 07 §8: pemetaan identitas yang dapat dipotong
```

Bentuk kolom ini sketsa. Nama dan pembagian `role` mengikuti 07 §1 saat dibangun. Yang mengikat adalah lima keputusan teknisnya:

1. **Append-only lewat hak akses.** `app_rw` hanya punya `INSERT` dan `SELECT` atas `entry`, `line`, dan `entry_link`. `UPDATE` dan `DELETE` tidak diberikan kepada **peran aplikasi mana pun** (`app_rw`, `lifecycle_rw`, `job_runner`), sehingga INV-09 tidak bergantung pada disiplin kode. `migrator` sebagai pemilik skema secara teknis dapat menulis ulang baris; karena itu migrasi tidak pernah menyentuh baris inti uang (§10), uji migrasi memeriksa hash inti sebelum dan sesudah setiap migrasi (SEC-108), dan setiap migrasi ditinjau dua pass (ES §7). Penghapusan menurut ADR-007 hanya menyentuh `entry_detail` dan `identity_map` lewat peran `lifecycle_rw`. Uji: `UPDATE entry` sebagai `app_rw` dan `lifecycle_rw` gagal dengan galat izin. Aturan hak yang sama berlaku atas tabel riwayat **`object_event`** dan **`security_event`**: `app_rw` hanya `INSERT` dan `SELECT`; penghapusan karena jangka simpan hanya lewat `lifecycle_rw` (dokumen 10).
2. **Saldo diturunkan dari baris** (ADR-004). Di pilot saldo dihitung saat dibaca dengan `SUM(delta)` lewat `line_by_position`. Volume pilot (ribuan baris per tahun **[Asumsi]**) membuatnya jauh di bawah p95 200 ms. Cache saldo baru ditambah bila pengukuran menuntut, dan harus lulus INV-04.
3. **Keseimbangan diperiksa di satu tempat.** INV-02 dan INV-03 diperiksa satu fungsi Go di modul `ledger` sebelum `INSERT`, di dalam transaksi. Fungsi itu satu-satunya implementasinya. Trigger basis data yang menghitung ulang persamaan yang sama **tidak** dibuat, karena itu implementasi kedua (ES §2). Pengamannya adalah uji invarian yang memblokir merge (berkas 04 §4), ditambah pemeriksaan malam yang menghitung ulang INV-02 sampai INV-08 atas seluruh data. Pemeriksaan malam itu memakai fungsi Go yang sama dan memberi tanda ke operator bila ada selisih.
4. **Mata uang dijaga foreign key.** Setiap tabel yang menyimpan jumlah uang punya `(household_id, currency) → household(id, currency)` dengan `ON UPDATE RESTRICT`. Itu mencakup entri, baris, tempat uang, dan juga anggaran, aturan, otorisasi, reservasi, jadwal, ambang, dan jatah. INV-35 karena itu dijaga di satu tempat, yaitu skema: setiap baris uang harus bermata uang rumah tangga, dan begitu ada satu baris uang, `household.currency` tidak dapat diganti. Kode Go tidak memeriksanya lagi. Ia hanya memetakan pelanggaran foreign key itu ke galat domain `currency_locked` atau `currency_mismatch` (§3 lapisan 3).
5. **`CHECK` tanda dan bentuk** untuk yang tidak memerlukan baris lain, misalnya `amount > 0` pada perintah yang menuntutnya dan INV-32 bentuk tempat uang.

### 5.3 Permintaan, chat, notifikasi (dari 06)

```sql
request (household_id, id, kind, state, version bigint NOT NULL, terms_version int NOT NULL, ...)  -- 06 §2.2
request_terms (household_id, request_id, terms_version, ..., PRIMARY KEY (household_id, request_id, terms_version))
request_response (household_id, request_id, responder, terms_version, decision, idempotency_key, ...)

conversation (household_id, id, participants uuid[] NOT NULL, ...)          -- 06 §16.2; peserta dibekukan
message (
  household_id uuid, conversation_id uuid, id uuid,
  seq          bigint NOT NULL,          -- per percakapan; INTERNAL, tidak keluar ke klien (ADR-017)
  kind         text NOT NULL,            -- text | photo | card | event_line (06 §16.2)
  author       uuid NULL,
  server_received_at timestamptz NOT NULL DEFAULT now(),  -- waktu tampil (SEC-64)
  idempotency_key uuid NOT NULL,
  removed      boolean NOT NULL DEFAULT false,            -- kartu status yang hilang (06 §13.5)
  PRIMARY KEY (household_id, conversation_id, seq),
  UNIQUE (household_id, id),
  UNIQUE (household_id, idempotency_key)
)
read_state (household_id, conversation_id, member_id,
            delivered_upto_seq bigint NOT NULL DEFAULT 0,
            read_upto_seq      bigint NOT NULL DEFAULT 0,
            CHECK (read_upto_seq <= delivered_upto_seq))  -- hanya naik: UPDATE … SET x = GREATEST(x, $n)
presence_last_active (household_id, member_id, last_active_floor timestamptz,  -- dipotong ke menit (SEC-60; S4-5 (ii) A)
                      PRIMARY KEY (household_id, member_id))

notification (
  household_id uuid, id uuid, recipient uuid,
  priority     text NOT NULL,             -- perlu_tanggapan | kabar | ringkasan (06 §11)
  catalog_id   text NOT NULL,             -- N-01 sampai N-32
  subject_kind text, subject_id uuid,
  collapse_group text NULL,               -- N-28 (SEC-57)
  hold_until   timestamptz NULL,          -- jam tenang (06 §11.1)
  released_at  timestamptz NULL,
  push_ref     uuid NOT NULL DEFAULT gen_random_uuid(),  -- pengenal opak di payload (SEC-50)
  ...
)
-- Satu tabel untuk ketiga kanal, sehingga penghapusan SEC-58 punya satu jalur (ADR-019 catatan web).
push_token (
  session_id  uuid PRIMARY KEY REFERENCES session(id) ON DELETE CASCADE,
  platform    text NOT NULL CHECK (platform IN ('apns','fcm','webpush')),
  token       text NOT NULL,              -- token APNs/FCM, atau endpoint Web Push (host dari daftar izin, ADR-030 butir 14)
  p256dh      bytea NULL,                 -- kunci publik langganan, hanya webpush (RFC 8291)
  auth        bytea NULL,                 -- rahasia autentikasi 16 byte, hanya webpush
  updated_at  timestamptz NOT NULL,
  CHECK ((platform = 'webpush') = (p256dh IS NOT NULL AND auth IS NOT NULL))
)

-- Mekanisme Apple 1.2, Tier 1 (10-G A, S4-7 A; FR-COM-019; 06 §16.12; SEC-153, SEC-154)
member_mute (household_id, viewer uuid, muted_member uuid, created_at,
             PRIMARY KEY (household_id, viewer, muted_member))   -- milik viewer; tidak terbaca anggota lain
-- content_report TIDAK dibuat selama pengelola server anggota rumah tangga pelapor (pilot):
-- "Laporkan" hanya membuka halaman Bantuan (SEC-154 (c)). Dibuat sejak build pertama ke TestFlight
-- eksternal atau App Store, atau sejak pengguna luar pertama membuat rumah tangga tanpa pengelola,
-- mana yang lebih dulu (SEC-154; keputusan 52), dan hanya dipakai bila pengelola bukan anggota
-- rumah tangga pelapor; tidak pernah diteruskan kepada orang yang dilaporkan.
content_report (household_id, id, reporter uuid, subject_kind text, subject_id uuid,
                attached_excerpt text NULL,  -- hanya butir yang dipilih pelapor sendiri
                created_at, handled_at NULL) -- dibaca pengelola lewat `casa ops report`, bukan API
```

### 5.4 Berkas

```sql
attachment (household_id, id, owner_kind, owner_id, media_type text, width int, height int,
            byte_size int, created_by uuid, ...)
attachment_blob (household_id, attachment_id, variant text CHECK (variant IN ('full','thumb')),
                 data bytea NOT NULL,
                 PRIMARY KEY (household_id, attachment_id, variant))
```

### 5.5 Sebutan anggota per pembaca (F-FAM-021, Tier 2; ringkas)

Tambahan pemilik produk di checkpoint S4 ([_progress.md](../_progress.md), "Keputusan checkpoint S4"). Aturan keamanannya kanonis di **SEC-155 (a) sampai (j)** (10.4 §6.3); tempat tampil dan kata-katanya di 08 §7.4; persyaratannya FR-FAM-025. Sketsa ini hanya menunjukkan bahwa SEC-155 dapat ditegakkan dan diuji dengan komponen yang sudah ada. **[Usulan]** `architect`.

```sql
member_label (
  household_id uuid, viewer uuid, subject uuid,  -- viewer = pembuat (anggota berakun); subject = anggota lain, termasuk tanggungan
  label        text NOT NULL,                    -- sesudah dipangkas dan NFC; 1..30 karakter tampak (SEC-155 (h))
  label_key    text NOT NULL,                    -- label_key(label): NFC lalu case folding Unicode; dihitung satu fungsi Go
  updated_at   timestamptz NOT NULL DEFAULT now(),  -- ditimpa; tanpa riwayat (SEC-155 (b))
  PRIMARY KEY (household_id, viewer, subject),
  UNIQUE (household_id, viewer, label_key),      -- satu pembuat tidak memberi sebutan sama kepada dua anggota
  CHECK (viewer <> subject),
  FOREIGN KEY (household_id, viewer)  REFERENCES membership (household_id, member_id),
  FOREIGN KEY (household_id, subject) REFERENCES membership (household_id, member_id)
)
-- Sebutan yang dibuat saat mengundang (SEC-155 (a)) menempel pada undangan, tidak tampil di mana pun
-- sampai keanggotaan aktif, menjadi baris member_label pada F04, dan dihapus bersama undangan pada F05
-- atau pembatalan (SEC-155 (g)(3)).
invitation (..., inviter_label text NULL, inviter_label_key text NULL)
```

| Hal | Penerapan | SEC-155 |
|---|---|---|
| Perintah | `set_member_label` (buat atau ganti) dan `clear_member_label`, oleh pembuat saja, untuk `viewer` = anggota pemanggil. Online-first, seperti pengaturan lain (03 §1); tanpa jeda urungkan, Peristiwa, kartu, atau notifikasi. F04 `confirm_invitee` memindahkan `inviter_label` menjadi baris milik pengundang dalam transaksi yang sama | (a), (b) |
| Bentuk masukan | Anotasi protovalidate di kontrak, ditegakkan server (ADR-013 butir 1): 1 sampai 30 karakter tampak sesudah dipangkas dan NFC, satu baris, tanpa karakter terlarang SEC-155 (h). Aplikasi hanya memeriksa **wajib diisi dan panjang teks**, yaitu bantuan isian (c) dari pengecualian ES §2 (ADR-013 butir 3), dengan vektor uji emas yang sama. Karakter terlarang, normalisasi NFC, dan duplikat **hanya diperiksa server**, karena daftar pengecualian ES §2 (c) tertutup | (h) |
| Tanpa duplikat | Dalam satu rumah tangga, satu pembuat tidak boleh memberi `label_key` yang sama kepada dua anggota. `label_key` dihitung **satu fungsi Go** di modul `household` (NFC lalu case folding Unicode), tidak dihitung ulang di basis data atau aplikasi. Penegaknya dua: `UNIQUE (household_id, viewer, label_key)` atas `member_label`, dan di dalam transaksi `SERIALIZABLE` yang sama perintah `set_member_label` serta pembuatan undangan juga menolak kunci yang sudah dipakai `inviter_label_key` undangan terbuka milik pembuat yang sama (satu `UNIQUE` tidak dapat mencakup dua tabel). Penolakannya galat per medan `label_duplicate`, yang hanya pernah dilihat pembuat sendiri, sehingga tidak membuka apa pun kepada orang lain (SEC-155 (b)). Karena kedua pintu masuk diperiksa, pemindahan pada F04 tidak dapat bertabrakan; bila tetap bertabrakan (misalnya data lama), F04 tetap berjalan dan sebutan undangan itu tidak dipindah, karena sebutan tidak pernah boleh menghalangi keanggotaan **[Usulan]** | (h) |
| Baca | `read.can_read` mengembalikan baris `member_label` hanya bila pembaca = `viewer`. Fan-out `reader_feed` karena itu hanya menulis baris untuk `viewer`; anggota lain tidak menerima `upsert` atau `purge` apa pun, sehingga umpan, potret awal, dan galat mereka sama dengan keadaan tanpa sebutan. Tidak ada kueri di jalur baca orang lain yang menyentuh `member_label`, jadi waktu tanggapan mereka juga tidak bergantung pada ada atau tidaknya sebutan. `viewer` adalah keanggotaan akun pembuat di rumah tangga itu, sehingga sebutan berlaku di semua sesi akun itu dan tidak lintas rumah tangga. Perintah atas baris milik orang lain tidak mungkin, karena kuncinya selalu diisi dari sesi pemanggil, bukan dari isi perintah | (b) |
| Penerapan di layar | Perangkat pembuat mengganti nama dari pengenal anggota dengan sebutan dari salinan bacanya sendiri, saat menampilkan saja. Server tidak pernah menyisipkan sebutan ke badan pesan, kartu, Peristiwa, teks ringkasan, payload push, atau teks notifikasi (payload push memang tanpa nama, SEC-50) | (c), (d) |
| `AI` | Modul `ai` tidak membaca `member_label`; aturan dependensi 2 di §1 ditambah daftar tolak untuk tabel ini, diperiksa uji arsitektur | (e) |
| Ekspor | `lifecycle` memuatnya hanya di ekspor milik `viewer` yang dibuat selama ia aktif, sebagai bagian tersendiri; tidak di ekspor anggota lain dan tidak di jendela ekspor sesudah keluar | (f) |
| Penghapusan | Dalam transaksi F07 anggota mana pun: `DELETE` semua baris dengan `viewer` **atau** `subject` = anggota itu di rumah tangga itu (dua arah). Hapus akun (FR-FAM-023): semua baris oleh dan tentang akun itu di setiap rumah tangga. Selama F06 baris tetap tersimpan, sehingga F08 memulihkannya; yang keluar tidak membacanya karena sesinya sudah dicabut. Setiap penghapusan menghasilkan `purge` hanya untuk `viewer` (SEC-48); cadangan mengikuti jangka 38 hari (D-104) | (g) |
| Log dan lintas rumah tangga | Medan `label` tidak ada di daftar izin log (09.4 §3); `household_id` di kunci dan RLS seperti tabel lain | (i) |

**Uji** (masuk rangkaian privasi 09.4 §4): uji SEC-155 (b) dijalankan dengan pola uji INV-A53, yaitu setiap API baca dan `Sync` dipanggil sebagai orang yang disebut dan sebagai anggota ketiga, dengan dan tanpa sebutan, dan hasilnya dibandingkan byte demi byte selain pengenal dan waktu; ditambah uji (c), (f), (g), dan (h) persis seperti kolom "Cara menguji" SEC-155.

Uji tambahan untuk keunikan: "Bunda" dan "bunda", serta bentuk NFC dan NFD dari sebutan yang sama, untuk dua anggota oleh pembuat yang sama ditolak `label_duplicate`; pembuat lain boleh memakai sebutan yang sama; sebutan pada undangan terbuka ikut dihitung; F04 tidak pernah gagal karena sebutan.

**Selisih antardokumen, diselesaikan lead (12 November 2026):** (1) panjang 1 sampai 30 karakter tampak (SEC-155 (h)); (2) tanpa duplikat per pembuat per rumah tangga sesudah NFC dan case folding, dimasukkan `security-privacy` ke SEC-155 (h); (3) ekspor pembuat memuat sebutannya di bagian tersendiri (SEC-155 (f)); (4) sebutan dibuat saat mengundang (SEC-155 (a)) dan baru tampil sesudah keanggotaan aktif. Sketsa di atas sudah mengikuti keempatnya.

---

## 6. Transaksi dan tingkat isolasi

| Jenis kerja | Tingkat isolasi | Alasan |
|---|---|---|
| Setiap perintah tulis | **`SERIALIZABLE`**, dicoba ulang otomatis pada `40001` (serialization failure) dan `40P01` (deadlock), paling banyak 5 kali dengan jeda acak yang membesar | INV-23: hasil dua perintah bersamaan sama dengan menjalankannya berurutan (18 §5.5) |
| Bacaan layar, laporan, ekspor | `REPEATABLE READ READ ONLY` | Satu potret konsisten untuk beberapa kueri satu layar (misalnya uang yang ada dan rinciannya) |
| Pekerjaan latar | Sama dengan perintah tulis | Pekerjaan adalah perintah sistem |

Catatan 13 November 2026 (FS-1 1.4, putusan lead Q1): "paling banyak 5 kali" berarti **5 coba ulang sesudah percobaan pertama, jadi paling banyak 6 percobaan**. Jeda sebelum coba ulang ke-`n` acak penuh dalam `[0, 10 ms × 2^(n−1))`, yaitu batas atas 10, 20, 40, 80, dan 160 ms **[Usulan]**. Sesudah coba ulang kelima gagal, `database.Write` mengembalikan galat yang dikirim ke klien sebagai `RETRY_EXHAUSTED` (§9; ADR-013). Bacaan tidak pernah dicoba ulang. Karena isi transaksi dapat berjalan lebih dari sekali, tidak ada efek di luar transaksi di dalamnya (email, hash kata sandi, batas laju): semuanya sebelum atau sesudah `Write`.

**Penilaian berurutan per anggota per periode** (06 S-8; INV-A37). SSI saja sudah mencegah hasil yang tidak berurutan, tetapi urutan pemenangnya tidak dapat diduga. Untuk membuat urutannya tetap, setiap entri yang dinilai terhadap jatah atau ambang mengambil kunci baris `SELECT … FOR UPDATE` atas `assessment_lane(household_id, member_id, period_start)` sebelum menghitung `terpakai_sebelum`. Entri yang tiba bersamaan dinilai sesuai urutan perolehan kunci itu, yang sama dengan urutan commit. Contoh urutan 07 §10.3 butir 7 menjadi uji konkurensi.

**Apa yang terjadi dalam satu transaksi perintah** (06 S-4):

```text
BEGIN ISOLATION LEVEL SERIALIZABLE
  SET LOCAL app.household_id = $checked_household   -- pemilih yang sudah dicocokkan dengan keanggotaan akun sesi (§3 lapisan 1)
  SET LOCAL app.account_id   = $session.account
  -- 1. idempotensi
  SELECT response FROM idempotency_record WHERE (scope_id, key) = …
    → ada dan request_hash sama : COMMIT, kembalikan response
    → ada dan request_hash beda : galat idempotency_conflict
  -- 2. otorisasi ulang (SEC-151): aktor, peran, state, visibilitas saat ini
  -- 3. compare-and-set: UPDATE request SET … , version = version + 1
  --                     WHERE id = $id AND version = $expected  → 0 baris: state_changed / terms_superseded
  -- 4. tulis domain: entry + line (+ entry_flag dari penilaian server)
  -- 5. object_event dengan seq = versi baru objek
  -- 6. reader_feed: satu baris per anggota yang can_read(objek) sesudah perubahan (ADR-017)
  -- 7. notification + job (outbox; berkas 03 §10)
  -- 8. INSERT idempotency_record (…, response)
COMMIT  → NOTIFY 'feed' (membangunkan hub SSE dan pekerja)
```

---

## 7. Uang: bilangan bulat dan mata uang (sweep M5)

- **Kolom:** setiap jumlah uang berjenis `bigint`, dalam satuan terkecil menurut `currency.exponent`, dan berdampingan dengan kolom `currency char(3)`. Ini berlaku di entri dan baris, dan juga di setiap jumlah yang tersimpan di luar entri: `budget`, `budget_rule`, `authorization`, `reservation`, `installment`, `recurring_bill`, `set_aside_schedule`, ambang, dan jatah (INV-01 sejak M5).
- **Eksponen IDR milik produk.** ISO 4217 mencantumkan IDR dengan minor unit **2** (18 §5, fase D) **[Terverifikasi]**; produk menetapkan **0** (ADR-008; 07 §2.1, INV-39). Eksponen hanya dibaca dari tabel `currency` milik produk, **tidak pernah** dari pustaka mata uang platform (misalnya `Currency.getDefaultFractionDigits` atau `NumberFormatter` dengan aturan ISO), dan uji unit klien serta server menolak pemakaian semacam itu di jalur uang. Setiap ekspor membawa `minor_unit_exponent` bersama `currency` dan bilangan bulat yang tersimpan tanpa diubah (SEC-103; 07 INV-39).
- **Di kabel:** protobuf `int64` dan `string currency`. Bentuk JSON protobuf untuk `int64` adalah string, sehingga klien tanpa int64 native pun tidak kehilangan presisi (kriteria ADR-008). Swift `Int64` dan Kotlin `Long` adalah tipe native.
- **Di perangkat:** `Int64` atau `Long`, tanpa `Double` di mana pun dalam jalur uang. Uji unit klien menolak tipe pecahan di model uang hasil generate.
- **Pembulatan:** tidak ada. Pembagian (misalnya bagian per orang) adalah aturan 07 dan dihitung di server.

---

## 8. Idempotensi (menuntaskan ADR-005)

ADR-005 menyerahkan tiga hal ke dokumen 09. Ketiganya ditetapkan di sini.

| Hal | Keputusan |
|---|---|
| **Bentuk kunci** | UUIDv4 dibuat klien satu kali **per niat pengguna** (ketukan "simpan"), disimpan bersama perintah di antrean lokal sebelum dikirim, dan dipakai ulang pada setiap pengiriman ulang. Server menyimpan `request_hash` (SHA-256 isi perintah yang dikanonkan). Kunci yang sama dengan isi berbeda ditolak `idempotency_conflict`, tidak diproses |
| **Cara keunikan dijaga** | `PRIMARY KEY (scope_id, idempotency_key)` pada `idempotency_record`, ditulis dalam transaksi yang sama dengan efeknya. Dua kiriman bersamaan: satu commit; yang lain mendapat pelanggaran unique atau `40001`, dicoba ulang, menemukan rekamannya, lalu mengembalikan hasil pertama (INV-10). Untuk entri uang, `entry.idempotency_key` juga unik per rumah tangga |
| **Jangka simpan** | **Perintah uang (entri):** permanen, karena kuncinya kolom `entry`. Antrean offline berapa lama pun tidak dapat menghasilkan entri ganda. **Perintah lain:** `idempotency_record` disimpan **90 hari**. Klien tidak mengirim otomatis perintah non-uang yang lebih tua dari **30 hari**; perintah itu ditampilkan kepada pemiliknya untuk dikirim ulang atau dibuang **[Usulan]**. Selisih 60 hari itu memberi ruang bagi jam perangkat yang salah |

**Perintah tulis** (catatan 13 November 2026, FS-1 1.3, kontrak dasar) adalah RPC yang membawa opsi metode `CommandOptions`, yaitu satu baris di tabel jenis perintah (09.3 §1), baik perintah antrean maupun perintah daring biasa. Hanya perintah tulis yang wajib membawa kunci idempotensi; server menolak perintah tulis tanpa kunci. RPC `identity` (masuk, autentikasi ulang, pembaruan token, pemulihan lewat email) berada di luar tabel jenis perintah dan tidak membawa kunci idempotensi dari mekanisme ini (09.3 §1).

Perintah tingkat akun di dalam tabel jenis perintah (membuat rumah tangga, menerima undangan) memakai `scope_id = account_id`.

**Pengecualian atas "coba ulang mengembalikan hasil perintah pertama" (ADR-005):** `SignOut` yang dicoba ulang dengan kunci idempotensi yang sama mendapat `SESSION_REVOKED` dari session interceptor, bukan hasil kosong yang tersimpan, karena sesi yang dinamainya sudah dicabut oleh kiriman pertama. Hal ini diterima karena klien tetap berakhir dalam keadaan keluar pada kedua jalur (catatan 13 November 2026, FS-1 1.4, code-reviewer M2 pada PR #32).

---

## 9. Kendali konkurensi

| Mekanisme | Dipakai untuk | Galat ke klien |
|---|---|---|
| Compare-and-set atas `version` objek | Setiap transisi permintaan, kesepakatan, dan objek berstate (06 S-4) | `state_changed`, beserta state terbaru yang **boleh dibaca** pemanggil (SEC-152) |
| `terms_version` di tanggapan | Persetujuan atas syarat yang digantikan (06 S-3) | `terms_superseded`, beserta syarat terbaru |
| Penulisan tambah-saja | Entri uang: tidak ada tabrakan; tanda dihitung saat diterima (06 S-7) | Tidak ada |
| Kunci baris jalur penilaian | Jatah dan ambang per anggota per periode (§6) | Tidak ada; hanya urutan |
| Tanda air `GREATEST` | `read_state` (06 S-11 (c)) | Tidak ada |
| SSI + coba ulang | Semua perintah | `retry_exhausted` (jarang; dicatat sebagai metrik) |

---

## 10. Evolusi skema

| Lapisan | Aturan |
|---|---|
| Basis data | Migrasi berversi dengan goose, masing-masing dengan arah naik dan turun (ES §7). Pola **expand–contract**: tambah kolom atau tabel, isi, alihkan kode, baru kemudian hapus yang lama di rilis berikutnya. Biner sebelumnya tetap berjalan di atas skema baru, sehingga rollback cukup dengan memasang biner lama. Baris inti uang tidak pernah ditulis ulang oleh migrasi (INV-09). Migrasi data atas entri dilakukan dengan entri baru, bukan `UPDATE` |
| Kontrak API | Aturan protobuf: nomor medan tidak dipakai ulang, medan hanya ditambah, nilai enum tak dikenal ditangani klien sebagai "tidak dikenal" yang aman. `buf breaking` memblokir merge yang merusak kompatibilitas |
| Aplikasi lama di tangan pengguna | Server menerima versi kontrak sampai N−1 (tahap publik). Di bawah itu server membalas `CLIENT_UPGRADE_REQUIRED` (ADR-028 butir 5; nama diselaraskan 13 November 2026, FS-1 1.3), dan aplikasi menampilkan layar perbarui. Di pilot cukup N bagi native; bagi web versi minimum berlaku sejak pilot (ADR-030 butir 9) |
| Basis data perangkat | Basis data lokal diperlakukan sebagai **cache yang dapat dibangun ulang**: bila skemanya berubah tidak kompatibel, cache dibuang dan disinkron ulang dari server. **Kecuali tabel antrean perintah**, yang skemanya dijaga stabil dan dimigrasi, karena isinya belum ada di server |

---

## 11. Penyimpanan berkas dan pembuangan metadata (D-94)

**Pilot: isi foto di PostgreSQL** (`attachment_blob`, `bytea`, disimpan otomatis sebagai TOAST) **[Usulan]**.

Alasannya:

- Satu jalur cadangan dan PITR mencakup foto (SEC-43).
- Menghapus pesan atau rincian menghapus foto dan thumbnail-nya dalam transaksi yang sama (SEC-65), tanpa berkas yatim.
- Tidak ada direktori di sistem berkas yang dapat salah dilayani.
- Otorisasi per permintaan memakai jalur baca yang sama dengan objeknya.

Perkiraan volume pilot: 5 foto per hari × 300 KB ≈ 1,5 MB per hari ≈ 550 MB per tahun **[Asumsi]**, jauh di bawah disk 60 GB.

**Batas pindah:** bila isi `attachment_blob` melewati 5 GB atau basis data melewati 20 GB **[Usulan]**, isi foto dipindah ke penyimpanan objek self-hosted di VPS terpisah. Garage atau SeaweedFS adalah kandidatnya (18 §5.7); MinIO diarsipkan 25 April 2026. Tabel `attachment` dan API-nya tidak berubah.

**Alur unggah:**

1. Pemilihan memakai pemilih foto sistem (PHPicker, Android Photo Picker; SEC-67).
2. Perangkat **menerapkan orientasi EXIF ke piksel** (memutar gambarnya), mengubah ruang warna ke sRGB, mengecilkan sisi terpanjang ke 2048 piksel **[Usulan]**, dan meng-encode ulang ke JPEG. Ini juga mengubah HEIC menjadi JPEG. Orientasi dan profil warna harus diselesaikan di sini, karena langkah 4 membuang keduanya: tanpa itu foto tegak tampil miring dan foto Display P3 iPhone tampil pudar (bacaan `researcher`, 18 §5 fase D, **[Usulan]**).
3. Server **hanya menerima JPEG** (format yang selalu dihasilkan langkah 2); masukan lain, termasuk PNG, ditolak. Server juga menolak lebih dulu berkas di atas 4 MB atau dimensi di atas 4096 piksel, memakai `image.DecodeConfig` sebelum dekode penuh. Ini mencegah bom dekompresi.
4. Server mendekode dan **meng-encode ulang ke JPEG saja** dengan paket standar Go `image/jpeg`, lalu membuat thumbnail JPEG. `image/jpeg` hanya menulis SOI, DQT, SOF0, DHT, SOS, dan EOI, tanpa segmen APPn apa pun (jadi tanpa EXIF, XMP, IPTC, JFIF, ICC, atau thumbnail tertanam) ([go.dev/src/image/jpeg/writer.go](https://go.dev/src/image/jpeg/writer.go), diakses 12 November 2026) dari kode sumber. **Server tidak pernah mengeluarkan PNG** (revisi fase D, CR-A-07): `image/png` menulis tRNS, chunk tambahan untuk transparansi gambar berpalet ([go.dev/src/image/png/writer.go](https://go.dev/src/image/png/writer.go)), sehingga keluaran PNG tidak akan memenuhi "hanya chunk kritis". Karena perangkat sudah mengirim JPEG, menolak PNG tidak menghilangkan fungsi apa pun, dan pertanyaan tRNS hilang. Karena belum dibuktikan pada berkas, klaim "keluaran tanpa metadata" tetap **[Perlu validasi]** sampai uji korpus SEC-66 lulus (syarat (a) di bawah).
5. Hanya hasil re-encode yang disimpan; berkas asli dibuang dari memori. Format yang tidak dapat di-encode ulang oleh paket standar Go **ditolak**, tidak pernah disimpan sebagai berkas asli.

**Encoder standar Go untuk SEC-66.** `security-privacy` menerima encoder standar Go sebagai ganti libvips, dengan empat syarat (dokumen 10, SEC-66), yang semuanya masuk rancangan ini:

| Syarat | Di mana |
|---|---|
| (a) Uji korpus membuktikan keluaran JPEG tanpa penanda APPn dan COM. Bagian PNG dipenuhi dengan **tidak ada keluaran PNG sama sekali** (langkah 3 dan 4); rumusan SEC-66 (a) disesuaikan `security-privacy` | Uji SEC-66 di CI (berkas 04 §4), dengan korpus berisi EXIF, XMP, IPTC, ICC, GPS, thumbnail tertanam, dan komentar JPEG; ditambah uji bahwa PNG (termasuk yang berisi `tEXt`, `iTXt`, `zTXt`, `eXIf`) ditolak dan tidak tersimpan |
| (b) Batas dimensi piksel diperiksa sebelum dekode | Langkah 3 (`image.DecodeConfig`) |
| (c) Format yang tidak dapat di-encode ulang ditolak, tidak pernah disimpan sebagai asli | Langkah 5 |
| (d) Bila salah satu syarat gagal, libvips dengan `keep NONE` lewat binding cgo | Jalan mundur; risikonya dicatat di ADR-016 |

Alasannya: encoder Go memory-safe, tanpa cgo dan tanpa pustaka C, sehingga satu dependensi native dan satu permukaan serangan dekoder hilang.

**Penayangan:** foto diambil lewat RPC `GetAttachment(id, variant)` dengan token sesi biasa. Tidak ada URL publik dan tidak ada tautan unduh tersendiri, sehingga syarat "tautan berumur pendek" SEC-65 tidak muncul. Aplikasi menyimpan cache di direktori cache yang dikecualikan dari cadangan.

---

## 12. Pencarian dalam bahasa Indonesia

| Tahap | Kebutuhan | Mekanisme |
|---|---|---|
| Tier 1 | Mencari objek di layarnya sendiri (catatan, permintaan, tempat uang), menurut kategori, tanggal, dan teks pendek; pencarian pesan chat **tidak** (F-COM-007 Tier 3) | Filter berindeks `(household_id, …)` ditambah `ILIKE '%…%'` atas judul dan catatan, dibatasi satu rumah tangga dan dipaginasi. Volume pilot membuatnya cukup **[Asumsi]**, diukur menurut ES §3 |
| Volume besar atau tahap publik | Pencarian teks yang toleran imbuhan dan salah ketik | Kolom `tsvector` dengan konfigurasi `indonesian` (Snowball, sejak PG12; mutu stemmer terbatas, 18 §5.5), ditambah indeks `pg_trgm` untuk salah ketik dan nama. Tanpa mesin pencari terpisah |
| Pencarian chat (Tier 3) | — | Dinilai bersama dokumen 10 saat naik: isi chat adalah kelas peka (SEC-63) |

Semua pencarian lewat jalur baca (SEC-06). Hasil difilter `can_read` di kueri yang sama, tidak sesudahnya. Dengan begitu jumlah hasil pun tidak membocorkan apa pun (SEC-07).
