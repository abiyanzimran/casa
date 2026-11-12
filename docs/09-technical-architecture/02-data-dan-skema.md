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
| `identity` | Akun, kata sandi, sesi per perangkat, token penyegar, autentikasi ulang, pemulihan email, batas laju, pendaftaran tertutup dengan token pemasangan | `account`, `credential`, `session`, `refresh_token`, `reauth_grant`, `recovery_token`, `setup_token`, `security_event` | Dokumen 10 SEC-11 sampai SEC-27, SEC-158 |
| `household` | Rumah tangga, keanggotaan, peran, undangan, keluar dan masa tenggang, pihak luar, pemberitahuan operator, `full_adult_members`, sebutan anggota per pembaca (Tier 2, §5.5) | `household`, `membership`, `invitation`, `external_party`, `operator_notice`, `operator_notice_ack`, `consent_record`, `member_label` | 06 §3; SEC-10, SEC-28 sampai SEC-34, SEC-109, SEC-155 |
| `ledger` | Entri, baris, tautan, rincian, posisi (tempat uang, piutang, kewajiban, aset), rekonsiliasi | `entry`, `line`, `entry_link`, `entry_detail`, `identity_map`, `holding`, `receivable`, `liability`, `asset`, `reconciliation` | 07 §1 sampai §3; ADR-002, ADR-004, ADR-007 |
| `budget` | Kategori, tujuan, periode, anggaran, otorisasi, reservasi, jatah, tagihan berulang, cicilan, jadwal penyisihan, sumber penghasilan, penilaian tanda; **kecocokan kategori** tautan ke pos rutin dan larangan mengganti kategori entri yang tertaut (satu fungsi; SEC-92; 07 §3.15 butir 7b, INV-37) | `category`, `purpose`, `budget_period_setting`, `budget`, `budget_rule`, `authorization`, `reservation`, `installment`, `recurring_bill`, `set_aside_schedule`, `income_source`, `assessment_lane`, `entry_flag` | 07 §2, §3, §6; 06 S-8 |
| `request` | Permintaan, versi syarat, tanggapan, kesepakatan rumah tangga, Peristiwa | `request`, `request_terms`, `request_response`, `object_event` | 06 §2, §4 sampai §8 |
| `notify` | Niat notifikasi, penahanan jam tenang, ringkasan, pengingat, token push | `notification`, `push_token`, `digest_state` | 06 §10, §11; SEC-50 sampai SEC-58 |
| `chat` | Percakapan, pesan, kartu, tanda air, kehadiran | `conversation`, `message`, `read_state`, `presence_last_active` | 06 §16; ADR-011 |
| `files` | Lampiran dan isinya, re-encode, thumbnail | `attachment`, `attachment_blob` | SEC-65, SEC-66 |
| `sync` | Umpan per pembaca, kursor, pembuangan | `reader_feed`, `reader_feed_counter` | ADR-017, ADR-018 |
| `jobs` | Antrean pekerjaan dan penjadwal | `job` | Berkas 03 §9 |
| `lifecycle` | Ekspor, penghapusan 90 hari, jurnal penghapusan, hapus akun | `export_link`, `deletion_journal` | SEC-94 sampai SEC-106 |
| `ai` (Tier 2) | Kumpulan kalimat, penghitung biaya, saklar mati | `sentence_pool`, `ai_usage` | Dokumen 10 §10.4; ADR-021 |
| `care`, `shopping`, `maintenance`, `plan` (Tier 2) | Sesuai 06 §13 dan tier-2.md | Tabel masing-masing | 06 §13 |
| `read` | **Satu jalur baca**: `can_read(member, object)`, proyeksi layar, laporan, ekspor, dan masukan `AI` | Tanpa tabel; fungsi dipakai semua modul | SEC-06; LM-07 |

**Aturan dependensi** (diperiksa uji arsitektur di CI, ADR-014):

1. `ledger` hanya ditulis oleh `ledger`. `budget` dan `request` memanggil fungsinya.
2. `read` adalah satu-satunya jalan bagi modul mana pun untuk membaca data milik orang lain. Ekspor, laporan, `sync`, dan `ai` memakai `read`, tidak menulis kueri sendiri (SEC-06).
3. `chat` tidak membaca isi objek lain. Kartu dirender dari `read` saat ditarik (INV-A44).
4. Tidak ada modul yang membaca `read_state` atau `presence` selain `chat` (LM-11; INV-A45).

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
  identity & household & request & ledger & budget & chat & files & sync & notify & jobs & lifecycle & ai --> db
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
| 5. Peran basis data | `app_rw`: bukan pemilik tabel, tanpa `BYPASSRLS`, tanpa `SUPERUSER`. `migrator`: pemilik skema, dipakai hanya oleh perintah migrasi. `lifecycle_rw`: satu-satunya peran dengan `DELETE` atas `entry_detail`, `identity_map`, dan data yang dihapus menurut ADR-007. `job_runner`: mengambil job (tabel di bawah). `backup`: peran replikasi untuk pgBackRest | SEC-09 |
| 6. Uji | Untuk setiap kueri yang disiapkan sqlc, uji lintas rumah tangga yang dibangkitkan: data rumah tangga X tidak pernah kembali bagi sesi rumah tangga Y (§9 berkas 04) | ES §4; LM-09 |

**Tabel global** (tanpa `household_id`): `account`, `credential`, `session`, `refresh_token`, `reauth_grant`, `recovery_token`, `security_event`, `push_token` (milik sesi), `currency` (tabel produk), dan `sentence_pool` (tanpa data rumah tangga). Kebijakan RLS-nya memakai `app.account_id`; `currency` dan `sentence_pool` hanya `SELECT` bagi `app_rw`.

**Tabel infrastruktur dengan pemilik campuran** (CS-A-20):

| Tabel | Kebijakan RLS bagi `app_rw` | Catatan |
|---|---|---|
| `idempotency_record` | `scope_id = current_setting('app.household_id', true)::uuid OR scope_id = current_setting('app.account_id', true)::uuid` | `scope_id` adalah rumah tangga, atau akun untuk perintah tingkat akun (§8) |
| `job` | `household_id = current_setting('app.household_id', true)::uuid OR (household_id IS NULL AND current_setting('app.household_id', true) IS NULL)` | Transaksi domain menulis job rumah tangganya sendiri; job sistem (`household_id` NULL, misalnya `ai.pool_build`, `ops.*`) hanya ditulis di luar konteks rumah tangga |
| `job` (pengambilan) | Peran tersendiri `job_runner`: hanya `SELECT` dan `UPDATE (locked_until, attempts, last_error)` atas `job`, dengan kebijakan `USING (true)`; tanpa hak atas tabel lain | Pekerja mengambil job lintas rumah tangga dengan peran ini, lalu menjalankan isinya sebagai `app_rw` dengan `SET LOCAL app.household_id = job.household_id`. `payload` hanya berisi pengenal, jadi melihat semua baris `job` tidak membuka isi rumah tangga |

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
  locale      text NOT NULL DEFAULT 'id'
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
push_token (session_id uuid PRIMARY KEY, platform text, token text, updated_at timestamptz)

-- Mekanisme Apple 1.2, Tier 1 (10-G A, S4-7 A; FR-COM-019; 06 §16.12; SEC-153, SEC-154)
member_mute (household_id, viewer uuid, muted_member uuid, created_at,
             PRIMARY KEY (household_id, viewer, muted_member))   -- milik viewer; tidak terbaca anggota lain
-- content_report TIDAK dibuat di pilot: "Laporkan" hanya membuka halaman Bantuan (SEC-154 (c)).
-- Sesudah pilot, hanya bila pengelola bukan anggota rumah tangga pelapor; tidak pernah diteruskan
-- kepada orang yang dilaporkan.
content_report (household_id, id, reporter uuid, subject_kind text, subject_id uuid,
                attached_excerpt text NULL,  -- hanya butir yang dipilih pelapor sendiri
                created_at, handled_at NULL) -- dibaca pengelola lewat `familyos ops report`, bukan API
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

Perintah tingkat akun (masuk, membuat rumah tangga, menerima undangan) memakai `scope_id = account_id`.

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
| Aplikasi lama di tangan pengguna | Server menerima versi kontrak sampai N−1 (tahap publik). Di bawah itu server membalas `upgrade_required`, dan aplikasi menampilkan layar perbarui. Di pilot cukup N |
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
