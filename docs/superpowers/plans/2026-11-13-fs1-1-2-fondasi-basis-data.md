# FS-1 1.2 — Fondasi basis data: rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** PostgreSQL 18 nyata di lokal dan CI; peran basis data dari 09.2 §3 lapisan 5; migrasi goose berversi yang terbukti naik-turun-naik; RLS dipaksakan dengan gerbang katalog yang gagal untuk setiap tabel baru yang tidak patuh; satu mekanisme transaksi yang memasang `app.household_id` dan `app.account_id` per transaksi; subperintah `casa migrate`; dan uji arsitektur impor antarpaket (ADR-014 butir 4).

**Arsitektur:** dua paket infrastruktur, bukan modul domain: `internal/schema` (berkas migrasi tersemat, `roles.sql`, penjalan goose) dan `internal/database` (pembuka transaksi bercakupan). Pembantu uji `internal/database/databasetest` membuat satu basis data baru per uji di klaster PostgreSQL uji, sehingga paket uji yang berjalan bersamaan tidak saling ganggu. Peran adalah objek tingkat klaster, jadi dibuat `roles.sql` (dijalankan superuser sekali per klaster), bukan migrasi; semua yang tingkat basis data (hak skema, tabel, kebijakan) ada di migrasi. PostgreSQL lokal dan CI dinyalakan oleh satu target `make db-up`, sehingga CI dan lokal tidak berbeda (ES §2, §7; fokus tinjauan 4 di 1.1a).

**Stack:** Go (toolchain di `server/go.mod`), pgx v5 (`pgx`, `pgxpool`, `pgconn`, `pgtype`, `stdlib`), goose v3 (API `Provider`), image Docker resmi `postgres` 18 (dibangun dari paket PGDG, ADR-016), GitHub Actions.

**Spesifikasi:** 14 §3.3 langkah 1 (kolom server); 09.2 §1 "Aturan dependensi", §3, §4, §5.1, §6, §10; 09.4 §1 (`LoadCredential`, rilis `casa migrate`), §4 (baris Migrasi, Arsitektur, Isolasi); 14 §4.1 butir 1, §5.1 QG-05, §6.4; ADR-014 butir 2 sampai 4; ADR-015 butir 3; ADR-016 butir 3, 4, 6, 11; SEC-08, SEC-09 (10.1); `ENGINEERING-STANDARDS.md` §1, §4, §7, §8.

**Pelaksana:** `backend-dev`. **Peninjau:** `code-reviewer` (dua lintasan). **Cabang:** `fs1/1.2-fondasi-basis-data` dari `main`.

## Batasan global

- Batasan global 1.1a berlaku: setiap alat dan dependensi dicek live saat dipasang (registri atau halaman rilis resmi, catatan rilis, advisori, isu kritis terbuka), stabil terbaru, dipatok tepat, dicatat di deskripsi PR (versi, tanggal cek, tautan, alasan) (ES §8). Versi di dokumen perencanaan (pgx v5.13.0, goose v3.30.0, PostgreSQL 18.6) **bukan** versi untuk dipasang.
- Dependensi baru hanya `github.com/jackc/pgx/v5` dan `github.com/pressly/goose/v3`. Tanpa testcontainers, pustaka UUID, testify, atau ORM. Pengenal memakai `pgtype.UUID` (bawaan pgx, sama dengan keluaran sqlc nanti).
- **Modul tak langsung 0.x** yang ikut masuk `go.mod` lewat pgx atau goose (mis. `golang.org/x/text`, `golang.org/x/sync`, modul ber-pseudo-versi): didaftar di laporan beserta versinya. ADR-023 tidak menyebut dependensi tak langsung; lead memutuskan sebelum merge (lihat "Celah dokumen" di laporan).
- Image PostgreSQL: `postgres:18.<minor terbaru>-<varian Debian terbaru>`, dipatok ke **digest `sha256` penuh** dengan komentar tag-nya, satu kali di `Makefile`.
- Nama peran persis (09.2 §3 lapisan 5): `migrator`, `app_rw`, `job_runner`, `lifecycle_rw`. Nama setelan persis: `app.household_id`, `app.account_id` (09.2 §3, §6).
- Uji basis data memakai PostgreSQL nyata (14 §4.1 butir 1). Bila `CASA_TEST_ADMIN_URL` tidak terpasang atau server tidak terjangkau, uji **gagal** (`t.Fatalf`), tidak pernah `t.Skip`.
- SQL selalu berparameter. Nama basis data uji (yang tidak dapat diparameterkan) disanitasi dengan `pgx.Identifier{name}.Sanitize()`. SQL di tugas ini (satu pernyataan `set_config` dan SQL uji) ditulis tangan; sqlc datang bersama kueri domain pertama (1.4).
- Tidak ada log; keluaran CLI tidak pernah memuat string koneksi atau isi berkas kredensial (ES §4).
- Biner `casa`, folder `server/cmd/casa` (keputusan pemilik produk 71). Jalur modul `github.com/abiyanzimran/casa/server`.
- Komentar, pengenal, dan pesan commit berbahasa Inggris; setiap aturan di kode mengutip ID atau bagian dokumennya; pesan commit menyebut ID (ES §1, §7).

## Fokus tinjauan

1. **Cakupan bocor antartransaksi pada koneksi pool yang dipakai ulang** (`SET` alih-alih `SET LOCAL`, atau transaksi yang gagal): transaksi berikutnya tanpa rumah tangga harus gagal membaca, bukan melihat rumah tangga sebelumnya. Uji di Tugas 4 dengan `MaxConns = 1`.
2. **Uji basis data yang lulus atau dilewati tanpa basis data**: pembantu gagal keras; dibuktikan merah di Tugas 7 dengan menjalankan `make server-test` tanpa `make db-up`.
3. **Migrasi dijalankan sebagai superuser** (pemilik tabel menjadi `postgres`, yang selalu melewati RLS): `schema.Up` menolak peran selain `migrator`, dan gerbang katalog memeriksa pemilik setiap tabel (Tugas 2, 3, 5).
4. **Jalan memutar RLS lewat hak akses**: kebijakan permisif kedua (di-OR-kan, mis. `USING (true)`), hak `TRUNCATE` (tidak tunduk RLS), atau keanggotaan peran aplikasi di `migrator`. Gerbang katalog dan uji SEC-09 menolak ketiganya (Tugas 1, 3).
5. **String koneksi berkata sandi tercetak di stderr** saat koneksi gagal atau URL rusak: uji penanda di Tugas 5.

## Yang sengaja tidak di tugas ini

- Tabel milik `identity` (`account`, `session`, …; 1.4) dan `household` (`household`, `membership`, …; 1.6). Tabel `household` sendiri memakai kolom `id`, bukan `household_id`; aturan kebijakannya di gerbang katalog ditambahkan 1.6.
- Tabel nyata bercakupan rumah tangga: belum ada yang dimiliki 1.2. RLS rumah tangga dibuktikan dengan tabel uji `rls_probe` yang dibuat di dalam uji (Tugas 4), dan gerbang katalog dibuktikan merah dengan migrasi buangan (Tugas 7). Satu-satunya tabel nyata 1.2 adalah `currency`: tabel infrastruktur 09.2 §5.1 tanpa modul pemilik, diperlukan FK `household.currency` (1.6) dan kolom uang (langkah 2).
- `idempotency_record` (tabel infrastruktur 09.2 §5.1): bersama perintah tulis pertama (1.3 atau 1.4), sesudah celah kebijakan `''` dijawab (laporan).
- Coba ulang `SERIALIZABLE` pada `40001`/`40P01` (09.2 §6): bersama penulis bersamaan pertama (1.4), sesudah lead menjawab "5 kali" = 5 percobaan atau 5 coba ulang (laporan).
- Peran `backup`: atributnya mengikuti kebutuhan pgBackRest yang dicek saat dipasang (langkah 7, 09.4 §6).
- Autentikasi peran di VPS (`pg_hba`, kata sandi lewat `LoadCredential` atau peer lewat soket Unix) dan unit systemd untuk `casa migrate`: langkah 7 (runbook).
- "Biner lama berjalan di atas skema baru" (09.4 §4; QG-05): belum dapat diuji karena belum ada migrasi di `main`; pertama bermakna di 1.4 (laporan).
- Kontrak dan `buf` (1.3), sesi web (1.5), `read`, `sync`, `jobs`, dan pembangkit uji isolasi per kueri sqlc (1.7).
- Larangan `float32`/`float64` di paket uang, larangan `UPDATE`/`DELETE` atas tabel inti uang di kode, hak append-only, dan uji hash inti SEC-108 sebelum dan sesudah migrasi: langkah 2, bersama `ledger`.
- Penambahan ID ke `traceability/scope.txt`: lead, di akhir langkah 1 (14 §4.13).

---

### Tugas 1: PostgreSQL lokal dan CI, `roles.sql`, pembantu `databasetest`, uji SEC-09

**Berkas:**
- Ubah: `Makefile` (target `db-up`, `db-down`; `server-test` mengekspor `CASA_TEST_ADMIN_URL`), `.github/workflows/ci.yml` (langkah `db-up` sebelum `server-test`)
- Ubah: `server/go.mod`, `server/go.sum`
- Buat: `server/internal/schema/roles.sql`, `server/internal/schema/schema.go`, `server/internal/schema/roles_test.go`
- Buat: `server/internal/database/databasetest/databasetest.go`

**Antarmuka:**
- Menghasilkan, di `package schema`: `//go:embed roles.sql` sebagai `var Roles string`.
- Menghasilkan, di `package databasetest`:
  - `type Database struct { Name, AdminRole string }`
  - `func Empty(t testing.TB) Database`: memastikan peran `roles.sql` ada, membuat basis data `casa_test_<16 hex acak>` dengan `OWNER migrator`, mendaftarkan `DROP DATABASE … WITH (FORCE)` di `t.Cleanup`.
  - `func (d Database) URL(role string) string`: `CASA_TEST_ADMIN_URL` dengan pengguna diganti `role` dan basis data diganti `d.Name`.
  - `func (d Database) Pool(t testing.TB, role string) *pgxpool.Pool` dan `func (d Database) SQL(t testing.TB, role string) *sql.DB`: ditutup di `t.Cleanup` (urutan LIFO menutupnya sebelum `DROP DATABASE`).
  - `Migrated` ditambahkan di Tugas 2.

- [ ] **Langkah 1: cek live** versi stabil terbaru pgx v5 (https://github.com/jackc/pgx/releases, pkg.go.dev) dan image `postgres` 18 (https://hub.docker.com/_/postgres, tag dan digest; catatan rilis https://www.postgresql.org/docs/release/). Cek juga Go stabil terbaru terhadap `go 1.27.4` di `go.mod`; menaikkannya adalah perubahan tersendiri, bukan bagian tugas ini, kecuali ada advisori keamanan.
- [ ] **Langkah 2: target `make`.**
  - `POSTGRES_IMAGE := postgres:<tag>@sha256:<digest>` dengan komentar sumber.
  - `db-up`: `docker run -d --rm --name casa-postgres -p 127.0.0.1:5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust --tmpfs /var/lib/postgresql $(POSTGRES_IMAGE)`, lalu tunggu `docker exec casa-postgres pg_isready -U postgres` paling lama 60 detik; gagal bila lewat. `trust` hanya untuk kontainer uji yang terikat `127.0.0.1` (14 §7.8: data uji saja); tidak ada kata sandi di repositori.
  - `db-down`: `docker rm -f casa-postgres`.
  - `server-test`: `CASA_TEST_ADMIN_URL ?= postgres://postgres@127.0.0.1:5432/postgres?sslmode=disable`, diekspor ke `go test`.
  - CI job `server`: langkah `db-up` (memanggil `make db-up`) sebelum `server-test`. Tidak memakai blok `services:`, agar image dipatok satu kali di `Makefile`.
- [ ] **Langkah 3: tulis uji yang gagal** di `roles_test.go` (`package schema_test`):

```go
func TestRoles(t *testing.T) {
	d := databasetest.Empty(t)
	db := d.SQL(t, d.AdminRole)
	roles := []string{"migrator", "app_rw", "job_runner", "lifecycle_rw"}
	t.Run("SEC-09 database roles are not superuser and cannot bypass row-level security", func(t *testing.T) {
		for _, role := range roles {
			// SELECT rolsuper, rolbypassrls, rolcreaterole, rolcreatedb, rolreplication, rolcanlogin
			//   FROM pg_roles WHERE rolname = $1
			// want: false, false, false, false, false, true; a missing role fails
		}
	})
	t.Run("SEC-09 application roles are not members of migrator", func(t *testing.T) {
		for _, role := range roles[1:] {
			// SELECT pg_has_role($1, 'migrator', 'MEMBER') → false
		}
	})
}
```

- [ ] **Langkah 4:** `make db-up && make server-test` → GAGAL (paket `databasetest` belum ada).
- [ ] **Langkah 5: tulis `roles.sql`.** Untuk setiap peran: blok `DO` yang `CREATE ROLE` bila belum ada, lalu `ALTER ROLE <r> LOGIN NOSUPERUSER NOCREATEDB NOCREATEROLE NOREPLICATION NOBYPASSRLS` (supaya menjalankan ulang juga membetulkan atribut). Tanpa kata sandi, tanpa pernyataan psql, tanpa `CREATE DATABASE`, sehingga dapat dijalankan dalam satu transaksi. Komentar kepala: dijalankan superuser sekali per klaster (`psql -v ON_ERROR_STOP=1 -f roles.sql`), lalu basis data dibuat `CREATE DATABASE casa OWNER migrator` (migrasi 00001 mengatur hak basis data sebagai pemiliknya); autentikasi peran diatur di `pg_hba` pada langkah 7; mengapa peran bukan migrasi (tingkat klaster; `migrator` tanpa `CREATEROLE`, SEC-09). Rujuk 09.2 §3 lapisan 5 dan ADR-016 butir 4 untuk setiap peran.
- [ ] **Langkah 6: tulis `databasetest`.** `Empty` menjalankan `schema.Roles` di satu transaksi admin yang diawali `SELECT pg_advisory_xact_lock(<konstanta>)`, karena paket uji berjalan bersamaan dan `CREATE ROLE` saling tabrak. Pesan bila env kosong: `CASA_TEST_ADMIN_URL is not set; run make db-up and make server-test`. `AdminRole` diambil dari pengguna di `CASA_TEST_ADMIN_URL`.
- [ ] **Langkah 7:** `make server-test` → LULUS. `make check-server` → LULUS.
- [ ] **Langkah 8: commit** `Add PostgreSQL 18 for tests and database roles (SEC-09; 09.2 §3 layer 5; ADR-016 item 4)`.

### Tugas 2: Migrasi goose, migrasi pertama, uji naik-turun-naik

**Berkas:**
- Buat: `server/internal/schema/00001_currency.sql`, `server/internal/schema/schema_test.go`
- Ubah: `server/internal/schema/schema.go`, `server/internal/database/databasetest/databasetest.go`, `server/go.mod`, `server/go.sum`

**Antarmuka:**
- Mengonsumsi: `databasetest.Empty`, `Database.SQL`, `Database.Pool` (Tugas 1).
- Menghasilkan, di `package schema`:
  - `func NewProvider(db *sql.DB) (*goose.Provider, error)`: `goose.NewProvider(goose.DialectPostgres, db, fsys)` atas `//go:embed *.sql` (tanpa `roles.sql`: sematkan pola `0*.sql`), registri global goose dimatikan.
  - `func Up(ctx context.Context, db *sql.DB) (version int64, err error)`: memeriksa `SELECT current_user` = `migrator`, bila tidak mengembalikan galat `must run as role migrator, connected as <role>` tanpa menjalankan apa pun; lalu menerapkan semua migrasi tertunda dan mengembalikan versi akhir.
- Menghasilkan, di `package databasetest`: `func Migrated(t testing.TB) Database` (= `Empty`, lalu `schema.Up` sebagai `migrator`).

- [ ] **Langkah 1: cek live** goose v3 stabil terbaru (https://github.com/pressly/goose/releases), catatan rilis dan isu terbuka tentang `Provider`.
- [ ] **Langkah 2: tulis uji yang gagal** di `schema_test.go`:

```go
// TestEveryMigrationGoesUpDownUp proves each migration reversible (ES §7;
// 09.4 §4 row "Migrasi"; 14 §6.4).
func TestEveryMigrationGoesUpDownUp(t *testing.T) {
	d := databasetest.Empty(t)
	db := d.SQL(t, "migrator")
	p, err := schema.NewProvider(db)
	// fatal on err; fatal if len(p.ListSources()) == 0 ("no migrations embedded")
	before := catalogSnapshot(t, db)
	for _, m := range p.ListSources() {
		// UpByOne → after := catalogSnapshot
		// Down → catalogSnapshot must equal before, else
		//   t.Fatalf("migration %d: down does not restore the catalog\nbefore:\n%s\nafter down:\n%s", …)
		// UpByOne → catalogSnapshot must equal after
		before = after
	}
}

func TestUpRefusesARoleOtherThanMigrator(t *testing.T) {
	d := databasetest.Empty(t)
	_, err := schema.Up(t.Context(), d.SQL(t, d.AdminRole))
	// err must contain "must run as role migrator, connected as " + d.AdminRole
	// goose_db_version must not exist afterwards (to_regclass('goose_db_version') IS NULL)
}

func TestCurrency(t *testing.T) {
	d := databasetest.Migrated(t)
	app := d.Pool(t, "app_rw")
	t.Run("holds IDR with the product exponent 0", …)
	// SELECT code, exponent FROM currency ORDER BY code → exactly [("IDR", 0)]
	t.Run("app_rw cannot change it", …)
	// INSERT ('USD', 2), UPDATE … SET exponent = 2, DELETE: each fails with PgError code 42501
}
```

`catalogSnapshot(t, db) string` (helper uji di berkas yang sama): teks berurutan dari skema `public`, tanpa `goose_db_version` dan tanpa nama basis data: `pg_class` (nama, `relkind`, pemilik, `relrowsecurity`, `relforcerowsecurity`, hak), kolom (nama, tipe, `NOT NULL`, default), `pg_get_constraintdef` setiap constraint, `indexdef`, seluruh kolom `pg_policies`, `pg_get_functiondef` fungsi di `public`, hak skema `public`, dan hak `current_database()`. Hak dibandingkan sebagai hak efektif, bukan teks ACL: `aclexplode(coalesce(acl, acldefault(<jenis>, <pemilik>)))`, diurutkan. Tanpa itu `datacl` NULL (bawaan) dan ACL eksplisit yang setara hasil `GRANT … TO PUBLIC` di bagian turun terlihat berbeda.

- [ ] **Langkah 3:** `make server-test` → GAGAL (`schema.NewProvider` belum ada).
- [ ] **Langkah 4: tulis `00001_currency.sql`**, satu transaksi goose. Bagian naik, berurutan:
  1. `REVOKE ALL ON SCHEMA public FROM PUBLIC; GRANT USAGE ON SCHEMA public TO app_rw, job_runner, lifecycle_rw;`
  2. Blok `DO` dengan `format('%I', current_database())`: `REVOKE ALL ON DATABASE … FROM PUBLIC; GRANT CONNECT ON DATABASE … TO app_rw, job_runner, lifecycle_rw;` (basis data staging dan produksi tidak saling memberi `CONNECT`, ADR-028 butir 5).
  3. `currency (code char(3) PRIMARY KEY, exponent smallint NOT NULL CHECK (exponent BETWEEN 0 AND 4))` persis 09.2 §5.1; `INSERT ('IDR', 0)` (ADR-008; 07 §2.1) **sebelum** RLS dinyalakan, karena dengan `FORCE` pemilik pun tunduk kebijakan dan kebijakan hanya `SELECT`.
  4. `ENABLE` dan `FORCE ROW LEVEL SECURITY`; `CREATE POLICY currency_read ON currency FOR SELECT USING (true)`; `GRANT SELECT ON currency TO app_rw` (09.2 §3 "Tabel global": hanya `SELECT` bagi `app_rw`).

  Bagian turun membalik semuanya, termasuk mengembalikan hak bawaan PostgreSQL 18: `GRANT USAGE ON SCHEMA public TO PUBLIC` dan `GRANT CONNECT, TEMPORARY ON DATABASE … TO PUBLIC`. Komentar di kepala: menambah mata uang kelak memerlukan `NO FORCE` sementara di migrasi yang sama, atau kebijakan `INSERT` bagi `migrator`.
- [ ] **Langkah 5: tulis `NewProvider`, `Up`, `Migrated`.**
- [ ] **Langkah 6:** `make server-test` → LULUS.
- [ ] **Langkah 7: commit** `Add goose migrations and the currency table (ADR-016 items 3 and 11; 09.2 §5.1, §10; ADR-008)`.

### Tugas 3: Gerbang katalog: setiap tabel dideklarasikan, dimiliki `migrator`, dan memaksakan RLS

**Berkas:** Buat `server/internal/schema/catalog_test.go`.

**Antarmuka:**
- Mengonsumsi: `databasetest.Migrated` (Tugas 2).
- Menghasilkan: peta `tables` yang ditambah satu baris oleh setiap tugas yang menambah tabel (1.4, 1.6, …), dengan sumber dokumennya.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
// tables declares every table in schema public. A migration that adds a table
// adds its row here, citing the document that defines the table and its grants.
var tables = map[string]tableRule{
	"currency": {global: true, grants: map[string]string{"app_rw": "SELECT"},
		source: "09.2 §3 \"Tabel global\", §5.1; ADR-008"},
}

type tableRule struct {
	global bool              // no household_id column (09.2 §3 "Tabel global")
	grants map[string]string // role → privileges, comma-separated in the order of allPrivileges; absent role: none
	source string
}

// householdPolicy is the one policy of every household table, as PostgreSQL
// deparses 09.2 §3 layer 4.
const householdPolicy = "(household_id = (current_setting('app.household_id'::text))::uuid)"

var allPrivileges = []string{"SELECT", "INSERT", "UPDATE", "DELETE", "TRUNCATE", "REFERENCES", "TRIGGER", "MAINTAIN"}

func TestCatalog(t *testing.T) { /* d := databasetest.Migrated(t); subtests below */ }
```

Subuji di `TestCatalog` (nama persis, satu klaim masing-masing):

| Subuji | Menegaskan |
|---|---|
| `every table is declared` | himpunan tabel `public` (relkind `r`, `p`; tanpa `goose_db_version`) sama dengan kunci `tables`; pesan menyebut tabel yang tak dideklarasikan atau yang hilang |
| `SEC-08 every table is owned by migrator and forces row-level security` | per tabel: pemilik `migrator`, `relrowsecurity` dan `relforcerowsecurity` benar, minimal satu kebijakan |
| `SEC-08 a household table has exactly one policy, the household policy` | tabel `!global`: punya kolom `household_id`; tepat satu baris `pg_policies`, `cmd = 'ALL'`, `roles = {public}`, `permissive = 'PERMISSIVE'`, `qual` dan `with_check` sama dengan `householdPolicy`. Tabel `global`: tidak punya kolom `household_id` |
| `SEC-08 role privileges match the declaration` | untuk `app_rw`, `job_runner`, `lifecycle_rw`, dan `PUBLIC`, hasil `has_table_privilege` atas `allPrivileges` sama dengan `grants`; selain itu tidak satu pun deklarasi memuat `TRUNCATE` (tidak tunduk RLS), `REFERENCES`, `TRIGGER`, atau `MAINTAIN` |
| `goose_db_version is not reachable by application roles` | ketiga peran aplikasi tanpa hak apa pun atas `goose_db_version` |
| `the database and schema public grant nothing to PUBLIC` | `aclexplode(datacl)` dan `aclexplode(nspacl)` tanpa penerima `0`; `app_rw` tanpa `CREATE` di `public` |
| `household_id leads every key of a household table (SEC-02)` | tabel `!global`: kolom pertama setiap primary key, unique constraint, unique index, dan foreign key (sisi asal dan sisi rujukan) adalah `household_id` |

Nama subuji terakhir tidak diawali `SEC-02`, karena cara uji SEC-02 di 10.1 adalah uji API (FR-FAM-020/AC-1 diperluas); ini hanya pagar skemanya.

- [ ] **Langkah 2:** `make server-test` → GAGAL. Cara melihat merah sebelum peta lengkap: jalankan dulu dengan `tables` kosong → `every table is declared` gagal menyebut `currency`.
- [ ] **Langkah 3:** isi baris `currency`; tulis kueri katalog dengan `has_table_privilege(role, table, priv)` (untuk `PUBLIC` pakai `aclexplode(relacl)` dengan penerima `0`).
- [ ] **Langkah 4:** `make server-test` → LULUS.
- [ ] **Langkah 5: commit** `Gate every table on owner, forced RLS, one household policy and declared grants (SEC-08, SEC-09, SEC-02; ADR-016 items 2-4)`.

### Tugas 4: Transaksi bercakupan rumah tangga dan akun (`internal/database`)

**Berkas:** Buat `server/internal/database/database.go`, `server/internal/database/database_test.go`.

**Antarmuka:**
- Mengonsumsi: `databasetest.Migrated`, `Database.Pool` (Tugas 1, 2).
- Menghasilkan, di `package database` (dipakai setiap modul mulai 1.4):

```go
// Scope is the household and account a transaction acts for. It is built only
// from values already checked against the session's membership (09.2 §3 layer 1;
// SEC-01); an invalid (zero) field means "none".
type Scope struct {
	HouseholdID pgtype.UUID
	AccountID   pgtype.UUID
}

// Write runs fn in one SERIALIZABLE transaction (09.2 §6) with s set for that
// transaction only.
func Write(ctx context.Context, pool *pgxpool.Pool, s Scope, fn func(pgx.Tx) error) error

// Read runs fn in one REPEATABLE READ READ ONLY transaction (09.2 §6) with s set
// the same way.
func Read(ctx context.Context, pool *pgxpool.Pool, s Scope, fn func(pgx.Tx) error) error
```

Keduanya memakai `pgx.BeginTxFunc`. Pernyataan pertama setiap transaksi, satu kali bolak-balik:
`SELECT set_config('app.household_id', coalesce($1::uuid::text, ''), true), set_config('app.account_id', coalesce($2::uuid::text, ''), true)`.
Argumen ketiga `true` = `SET LOCAL` (09.2 §6). Kedua setelan **selalu** dipasang, dengan `''` bila tidak ada, sehingga keadaannya sama di setiap transaksi apa pun riwayat koneksi pool-nya. Galat dari `fn` dikembalikan dengan rantai utuh (`errors.Is` berlaku). Tanpa coba ulang (lihat "Yang sengaja tidak").

- [ ] **Langkah 1: tulis uji yang gagal.** Fixture dibuat di dalam uji sebagai `migrator` sesudah `Migrated`: `rls_probe (household_id uuid NOT NULL, id uuid NOT NULL DEFAULT gen_random_uuid(), PRIMARY KEY (household_id, id))`, `ENABLE` dan `FORCE ROW LEVEL SECURITY`, satu kebijakan `USING` dan `WITH CHECK` dengan ekspresi 09.2 §3 lapisan 4, `GRANT SELECT, INSERT, UPDATE, DELETE` ke `app_rw`. Data: dua baris rumah tangga H1, satu baris H2, ditulis lewat `Write`.

```go
func TestScope(t *testing.T) {
	t.Run("SEC-08 a transaction reads only rows of its household, with no filter in the query", …)
	// Read(H1): SELECT household_id FROM rls_probe → 2 rows, all H1; Read(H2) → 1 row, H2
	t.Run("SEC-08 a row for another household cannot be written", …)
	// Write(H1): INSERT household_id = H2 → PgError 42501;
	//   UPDATE rls_probe SET household_id = H2 → PgError 42501;
	//   DELETE FROM rls_probe WHERE household_id = H2 → RowsAffected 0;
	// then Read(H2) still sees 1 row
	t.Run("SEC-08 without a household a household table cannot be read", …)
	// Read(Scope{AccountID: A}): SELECT FROM rls_probe → PgError 22P02, no rows
	t.Run("SEC-08 the table owner is filtered too", …)
	// as migrator: BEGIN; set_config('app.household_id', H1, true); count(*) → 2 (FORCE)
	t.Run("SEC-08 the scope ends with its transaction", …)
	// pool with MaxConns 1, so every call below reuses one connection.
	// Write(H1) commits; then, outside Read and Write (they overwrite both settings
	// and would hide a leak), plain pool queries on that connection:
	//   SELECT coalesce(current_setting('app.household_id', true), '') → ''
	//   SELECT count(*) FROM rls_probe → PgError 22P02
	// Write(H1) with fn returning errSentinel → errors.Is(err, errSentinel);
	// the same two plain queries give the same results
	t.Run("write is serializable and read is repeatable read and read only", …)
	// in Write: SHOW transaction_isolation = "serializable";
	// in Read: "repeatable read", SHOW transaction_read_only = "on"; INSERT → PgError 25006
	t.Run("the account is set when given", …)
	// Read(Scope{AccountID: A}): current_setting('app.account_id') = A as text;
	// Read(Scope{}): ''
}
```

- [ ] **Langkah 2:** `make server-test` → GAGAL (`database.Write` belum ada).
- [ ] **Langkah 3:** tulis `Scope`, `Write`, `Read`.
- [ ] **Langkah 4:** `make server-test` → LULUS.
- [ ] **Langkah 5: commit** `Add household- and account-scoped transactions (SEC-08; 09.2 §3 layer 4, §6; ADR-016 items 3 and 6)`.

### Tugas 5: Subperintah `casa migrate`

**Berkas:**
- Buat: `server/cmd/casa/migrate.go`, `server/cmd/casa/migrate_test.go`
- Ubah: `server/cmd/casa/main.go`, `server/cmd/casa/main_test.go`

**Antarmuka:**
- Mengonsumsi: `schema.Up`, `schema.NewProvider` (Tugas 2); `databasetest.Empty`, `Database.URL` (Tugas 1).
- Menghasilkan: `func run(ctx context.Context, args []string, stdout, stderr io.Writer) int` (menggantikan `run(args, stderr)`; `main` memberi `signal.NotifyContext` untuk `os.Interrupt` dan `syscall.SIGTERM`); `func readCredential(name string) (string, error)` di `package main`, dipakai lagi oleh `serve` (1.3). Kode keluar: 0 berhasil, 1 gagal, 2 salah pakai.

- [ ] **Langkah 1: tulis uji yang gagal.** Setiap subuji menulis berkas `migrator-database-url` di `t.TempDir()` dan `t.Setenv("CREDENTIALS_DIRECTORY", dir)`.

```go
func TestMigrate(t *testing.T) {
	t.Run("applies every migration and reports the version", …)
	// credential = d.URL("migrator") of databasetest.Empty; run(ctx, {"migrate"}) = 0;
	// stdout = fmt.Sprintf("casa migrate: database at version %d\n", <last version of ListSources>);
	// a second run: 0 and the same stdout
	t.Run("refuses a role other than migrator", …)
	// credential = d.URL(d.AdminRole) → 1; stderr contains
	// "casa migrate: must run as role migrator, connected as " + d.AdminRole
	t.Run("never prints the connection string", …)
	// for each credential in {"postgres://migrator:leak-marker-7f3@127.0.0.1:1/casa?connect_timeout=2",
	//                          "postgres://migrator:leak-marker-7f3@[bad"}: exit 1;
	// stderr does not contain "leak-marker-7f3"
	t.Run("fails without CREDENTIALS_DIRECTORY", …)
	// Setenv("CREDENTIALS_DIRECTORY", "") → 1; stderr =
	// "casa migrate: CREDENTIALS_DIRECTORY is not set; casa reads secrets from systemd LoadCredential= (09.4 §1)\n"
	t.Run("rejects extra arguments", …)
	// run(ctx, {"migrate", "down"}) → 2; stderr = "usage: casa migrate\n"
}
```

- [ ] **Langkah 2:** `make server-test` → GAGAL.
- [ ] **Langkah 3:** tulis `migrate.go`: baca kredensial `migrator-database-url` (spasi dan baris baru di ujung dibuang), buka `*sql.DB` lewat `pgx/v5/stdlib`, panggil `schema.Up`. Galat dicetak `casa migrate: <galat>`; galat yang memuat string koneksi (parse URL) diganti pesan tetap `cannot parse the connection string in credential migrator-database-url`. Komentar: rilis menjalankan `casa migrate` sebagai `migrator` sebelum biner baru dipasang (09.4 §1 "Rilis server"; 14 §6.4); tidak ada arah turun di CLI, karena rollback = biner lama di atas skema expand–contract (09.2 §10). Sesuaikan dua uji lama ke tanda tangan `run` yang baru.
- [ ] **Langkah 4:** `make server-test` → LULUS.
- [ ] **Langkah 5: commit** `Add casa migrate (09.4 §1 release; 14 §6.4; SEC-09)`.

### Tugas 6: Uji arsitektur impor (ADR-014 butir 4)

**Berkas:** Buat `server/internal/architecture/imports_test.go` (paket berisi uji saja).

**Antarmuka:**
- Menghasilkan: peta `allowedImports` yang ditambah satu baris oleh setiap tugas yang menambah paket, mengutip 09.2 §1. Baris yang tidak tercantum di 09.2 §1 (diagram atau "Aturan dependensi") memerlukan pembaruan 09.2 §1 oleh `architect` lebih dulu.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
// allowedImports lists, per package of this module, the other packages of this
// module its production code may import (ADR-014 item 4; 09.2 §1 "Aturan
// dependensi"). Every package must be listed. Test files are not checked, so a
// test may import databasetest or another module to set up data.
var allowedImports = map[string][]string{
	"cmd/casa":                       {"internal/schema"},
	"cmd/idcheck":                    {},
	"internal/architecture":          {},
	"internal/database":              {},
	"internal/database/databasetest": {"internal/schema"},
	"internal/schema":                {},
}

func TestModuleImportsFollowTheAllowList(t *testing.T) {
	// go list -f '{{.ImportPath}}{{range .Imports}} {{.}}{{end}}' ./...  (Dir: module root)
	// for each package P (path without the module prefix):
	//   not in allowedImports → t.Errorf("%s is not declared in allowedImports (ADR-014 item 4)", P)
	//   each import I with the module prefix and I not in allowedImports[P] →
	//     t.Errorf("%s imports %s, which allowedImports does not allow (ADR-014 item 4; 09.2 §1)", P, I)
	// each key of allowedImports missing from go list → t.Errorf("%s is declared but does not exist", key)
}
```

- [ ] **Langkah 2:** jalankan dengan peta kosong → GAGAL, menyebut keenam paket.
- [ ] **Langkah 3:** isi peta seperti di atas.
- [ ] **Langkah 4:** `make server-test` → LULUS.
- [ ] **Langkah 5: commit** `Add the module import architecture test (ADR-014 item 4; 09.2 §1)`.

### Tugas 7: README, bukti setiap gerbang merah, PR

- [ ] **Langkah 1: README akar**: prasyarat Docker kini dipakai; satu bagian pendek "Database": `make db-up`, `make server-test`, `make db-down`; `CASA_TEST_ADMIN_URL` bawaan; `roles.sql` dijalankan superuser sekali per klaster. Bahasa Inggris.
- [ ] **Langkah 2:** push cabang, buka PR draf ke `main`. CI → HIJAU, termasuk langkah `db-up`.
- [ ] **Langkah 3: buktikan setiap gerbang dapat merah**, satu komit buangan per gerbang di cabang `fs1/1.2-redproof` (PR draf terpisah, ditutup dan cabangnya dihapus sesudahnya). Harapan untuk masing-masing:

| Pelanggaran buangan | Yang harus merah |
|---|---|
| Langkah CI `db-up` dihapus | `server-test` gagal dengan pesan koneksi atau `CASA_TEST_ADMIN_URL`; tidak ada uji basis data yang `SKIP` |
| Migrasi `00002` yang `Down`-nya tidak menghapus satu kolom yang ditambah `Up` | `TestEveryMigrationGoesUpDownUp` |
| Migrasi `00002` dengan tabel `household_id` tanpa `FORCE ROW LEVEL SECURITY`, dideklarasikan di `tables` | `SEC-08 every table is owned by migrator and forces row-level security` |
| Tabel yang sama dengan kebijakan kedua `USING (true)` | `SEC-08 a household table has exactly one policy, the household policy` |
| `GRANT TRUNCATE ON currency TO app_rw` | `SEC-08 role privileges match the declaration` |
| Tabel `household_id` dengan `UNIQUE (id)` | `household_id leads every key of a household table (SEC-02)` |
| `NOBYPASSRLS` diganti `BYPASSRLS` untuk `app_rw` di `roles.sql` (kontainer CI baru, jadi tidak tersisa) | `SEC-09 database roles are not superuser and cannot bypass row-level security` |
| `set_config(…, false)` di `database.go` | `SEC-08 the scope ends with its transaction` |
| `NO FORCE` di fixture `rls_probe` | `SEC-08 the table owner is filtered too` |
| Pemeriksaan `current_user` dihapus dari `schema.Up` | `TestUpRefusesARoleOtherThanMigrator` dan `refuses a role other than migrator` |
| `internal/database` mengimpor `internal/schema` | `TestModuleImportsFollowTheAllowList` |

Setiap komit buangan patuh pada semua gerbang lain, sehingga yang merah hanya gerbang sasarannya. Catat tautan run CI merah di laporan.
- [ ] **Langkah 4: PR siap ditinjau.** Deskripsi PR memuat tabel dependensi dan image (versi, tanggal cek, tautan, alasan), daftar modul tak langsung 0.x, tautan run merah, dan catatan ES §9 butir 5: `casa migrate` dan seluruh uji dijalankan terhadap PostgreSQL 18 nyata; tidak ada perilaku pengguna atau perangkat di tugas ini.

## Selesai bila

- CI hijau di PR; setiap gerbang di tabel Tugas 7 terbukti merah sekali (tautan tercatat).
- Modul tak langsung 0.x diputuskan lead; pertanyaan di laporan sudah sampai ke lead.
- `code-reviewer` dua lintasan tanpa temuan terbuka; lead merge ke `main`.
