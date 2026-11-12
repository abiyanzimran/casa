# FS-1 1.1a — Kerangka repositori dan CI server: rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** repositori satu-monorepo dengan modul Go server yang paling kecil, dan CI GitHub Actions yang menjalankan gerbang server QG-01, QG-02, QG-03, QG-05 (uji Go) serta pemeriksaan rahasia, semuanya hijau dan terbukti gagal saat dilanggar.

**Arsitektur:** satu repositori (`proto/`, `server/`, `web/`, `golden/`; 14 §2.2 "Satu repositori"). Perintah gerbang ditulis sekali di `Makefile` dan dipanggil oleh CI maupun pengembang, supaya CI dan lokal tidak berbeda (ES §2). CI berjalan pada `pull_request` dan `push` ke `main`.

**Stack:** Go (versi stabil terbaru, dicek live), golangci-lint v2, govulncheck, PMD CPD, alat pemeriksa rahasia (dipilih di tugas ini), GitHub Actions.

**Spesifikasi:** `docs/14-engineering-roadmap.md` §2.2, §3.3 langkah 1, §5.1 (QG-01 sampai QG-03, QG-05); `docs/09-technical-architecture/04-operasi-uji-biaya.md` §4, §5; `docs/adr/ADR-025-alat-kualitas-kode.md` butir 1, 4, 6, 7; `ENGINEERING-STANDARDS.md` §1, §7, §8.

**Pelaksana:** `backend-dev`. **Peninjau:** `code-reviewer` (dua lintasan). **Cabang:** `fs1/1.1a-repo-ci-server`.

## Batasan global

- Setiap alat dan dependensi: versi dicek live di registri atau halaman rilis resmi saat dipasang; stabil terbaru; dipatok tepat; dicatat di deskripsi PR: versi, tanggal cek, tautan sumber, alasan (ES §8). Versi di dokumen perencanaan (mis. "Go 1.27") **bukan** versi untuk dipasang.
- GitHub Action pihak ketiga dipatok ke **SHA komit penuh**, dengan komentar nomor versinya. `permissions: contents: read` di tingkat workflow.
- Tidak ada alat yang mengunggah kode ke layanan selain GitHub (ADR-025 "Konsekuensi").
- Jalur modul Go: `github.com/abiyanzimran/casa/server`. Nama biner: `familyos` (09.4 §6.2; ADR-030 butir 9: "Family OS" tetap pengenal yang sudah ada).
- Komentar, pengenal, dan pesan commit berbahasa Inggris; pesan commit menyebut ID (ES §7).

## Fokus tinjauan

1. Gerbang yang "hijau" karena tidak memeriksa apa pun (pola jalur salah, direktori kosong, `|| true`): setiap gerbang harus terbukti merah pada pelanggaran nyata (Tugas langkah 6).
2. Unduhan alat di CI tanpa verifikasi checksum: setiap biner yang diunduh langsung diverifikasi SHA-256-nya.
3. Pemeriksa rahasia yang hanya memindai berkas yang berubah, bukan seluruh riwayat cabang.
4. Perbedaan antara perintah lokal dan perintah CI: CI hanya memanggil target `make`.
5. Ambang CPD yang dipilih sendiri oleh pelaksana: ambang ditetapkan `code-reviewer` (ADR-025 butir 4).

## Yang sengaja tidak di tugas ini

- `buf lint`/`buf breaking` dan kontrak: tugas 1.3 (proto pertama lahir di sana; gerbang tanpa berkas tidak membuktikan apa pun).
- PostgreSQL di CI, migrasi, uji arsitektur impor: tugas 1.2.
- Larangan `float32`/`float64` di paket uang: langkah 2, bersama paket `ledger` (uji arsitektur tanpa paketnya hampa).
- Toolchain dan CI web: 1.1b. Pemeriksa ID (QG-06): 1.1c.

---

### Tugas 1: Modul Go dan biner `familyos`

**Berkas:**
- Buat: `server/go.mod`, `server/cmd/familyos/main.go`, `server/cmd/familyos/main_test.go`
- Buat: `.gitignore` (biner, `coverage.out`, `node_modules/`, `.env*`)

**Antarmuka:**
- Menghasilkan: `func run(args []string, stdout, stderr io.Writer) int` di `package main`. `main` hanya memanggil `os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))`. Tugas 1.2 menambah subperintah `migrate`, tugas 1.3 `serve`, lewat `switch` di `run`.

- [ ] **Langkah 1: cek live versi Go stabil terbaru** di https://go.dev/dl/ (dan catatan rilisnya serta advisori keamanannya, ES §8). Pasang di mesin bila lokal lebih tua (lokal saat ini go1.26.3). Tulis `go.mod` dengan `go <versi>` sesuai itu.
- [ ] **Langkah 2: tulis uji yang gagal**

```go
func TestRunWithoutCommandPrintsUsageAndExits2(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := run(nil, &out, &errOut); code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
	if !strings.HasPrefix(errOut.String(), "usage: familyos <command>") {
		t.Fatalf("stderr = %q", errOut.String())
	}
	if out.Len() != 0 {
		t.Fatalf("stdout must be empty, got %q", out.String())
	}
}

func TestRunUnknownCommandExits2(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := run([]string{"nope"}, &out, &errOut); code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
	if !strings.Contains(errOut.String(), `unknown command "nope"`) {
		t.Fatalf("stderr = %q", errOut.String())
	}
}
```

- [ ] **Langkah 3:** `cd server && go test ./...` → GAGAL (`run` belum ada).
- [ ] **Langkah 4:** tulis `run`. Baris pemakaian persis `usage: familyos <command>\n`; perintah tak dikenal: `familyos: unknown command "<x>"\n` lalu baris pemakaian.
- [ ] **Langkah 5:** `go test -race -shuffle=on ./...` → LULUS.

### Tugas 2: `Makefile` sebagai satu sumber perintah gerbang server

**Berkas:** Buat `Makefile` (akar), `server/.golangci.yml`, `scripts/cpd.sh`.

Target (nama tetap; dipakai 1.1b, 1.1c, 1.2, 1.3):

| Target | Isi |
|---|---|
| `server-fmt` | `test -z "$(gofmt -l server)"`, dan bila tidak kosong mencetak daftar berkasnya lalu gagal |
| `server-lint` | `go vet ./...` lalu `golangci-lint run ./...` di `server/` |
| `server-vuln` | `govulncheck ./...` di `server/`, versi govulncheck dipatok (`go run golang.org/x/vuln/cmd/govulncheck@vX.Y.Z`) |
| `server-test` | `go test -race -shuffle=on -coverprofile=coverage.out ./...` di `server/` (14 §2.2: `-race`, `-shuffle`; cakupan dilaporkan tanpa ambang, 14 §4.13) |
| `cpd-go` | `scripts/cpd.sh go server` |
| `secrets` | pemeriksa rahasia atas seluruh riwayat git (Tugas 3) |
| `check-server` | semua target di atas berurutan |

- [ ] **Langkah 1: `.golangci.yml` versi 2.** Linter: himpunan `standard` ditambah `gosec`, `errorlint`, `misspell`. Formatter `gofmt`. Tanpa `nolint` tanpa alasan: aktifkan `nolintlint` dengan `require-explanation: true` (ES §1 "A suppression needs a comment giving the reason"). Versi golangci-lint dicek live, v2.x stabil.
- [ ] **Langkah 2: `scripts/cpd.sh <language> <dir>`.** Mengunduh distribusi PMD versi stabil terbaru (dicek live) ke `.cache/pmd/` bila belum ada, **memverifikasi SHA-256** terhadap nilai yang ditulis di skrip (diambil dari halaman rilis resmi), lalu menjalankan `pmd cpd --language <language> --minimum-tokens <ambang> --dir <dir> --fail-on-violation true`. `.cache/` masuk `.gitignore`. **Ambang token**: lead meminta nilainya ke `code-reviewer` sebelum tugas ini dimulai (ADR-025 butir 4) dan menuliskannya di sini: `__________`. Pelaksana tidak memilih sendiri.
- [ ] **Langkah 3:** `make check-server` lokal → LULUS.

### Tugas 3: Pemeriksa rahasia (14 §2.2 baris "Rahasia di repositori")

- [ ] **Langkah 1: pilih alat.** Syarat: berjalan lokal dan di CI tanpa mengunggah kode ke layanan lain; dipelihara aktif; lisensi terbuka; versi stabil; biner diunduh dengan checksum diverifikasi atau dipasang lewat `go install` berversi tepat; dapat memindai seluruh riwayat git. Calon yang wajib dinilai: gitleaks; boleh menambah satu pembanding. Catat penilaiannya (versi, tanggal cek, tautan, lisensi, tanggal rilis terakhir, isu kritis terbuka) di laporan.
- [ ] **Langkah 2:** isi target `secrets` (pindai seluruh riwayat; gagal bila ada temuan). Konfigurasi bawaan alat; pengecualian hanya dengan alasan tertulis.
- [ ] **Langkah 3: buktikan merah.** Di cabang buangan lokal (tidak di-push), komit berkas berisi kunci palsu berpola nyata (mis. kunci AWS contoh dokumentasi yang dikenali alat) → `make secrets` GAGAL. Hapus cabang buangan.
- [ ] **Langkah 4: ADR.** Pilihan alat memerlukan revisi ADR-025 (14 §2.2). Pelaksana **tidak** menyunting ADR; ia menyerahkan usulannya di laporan. Lead meneruskan ke `architect`, yang menulis catatan ADR-025 di cabang yang sama sebelum merge (ES §9 butir 6).

### Tugas 4: Workflow CI

**Berkas:** Buat `.github/workflows/ci.yml`.

- [ ] **Langkah 1:** pemicu `pull_request` (ke `main`) dan `push` (ke `main`); `concurrency` per ref dengan `cancel-in-progress: true`; `permissions: contents: read`.
- [ ] **Langkah 2:** job `server` di `ubuntu-latest` dipatok ke label versi (mis. `ubuntu-24.04`, dicek live): `actions/checkout` dengan `fetch-depth: 0` (pemeriksa rahasia butuh riwayat), `actions/setup-go` dengan `go-version-file: server/go.mod`, `actions/setup-java` (distribusi Temurin, LTS terbaru dicek live) untuk PMD, lalu satu langkah per target `make` (supaya kegagalan terlihat per gerbang). golangci-lint dipasang dengan cara yang dipatok versinya dan diverifikasi.
- [ ] **Langkah 3:** unggah `server/coverage.out` sebagai artefak (laporan, bukan gerbang).

### Tugas 5: README akar

**Berkas:** Buat `README.md` di akar: satu paragraf apa ini (Casa; rujuk `docs/README.md` dan `CLAUDE.md`), tata letak folder, prasyarat lokal (Go, Java untuk PMD, Docker untuk PostgreSQL mulai 1.2), dan `make check-server`. Bahasa Inggris (berkas di luar `docs/`; `CLAUDE.md` "Language" hanya mengatur `docs/` dan berkas yang disebut). Tanpa isi lain.

### Tugas 6: Bukti setiap gerbang merah, lalu hijau

- [ ] **Langkah 1:** push cabang, buka PR draf ke `main`. CI → HIJAU.
- [ ] **Langkah 2: buktikan setiap gerbang dapat merah**, satu komit buangan per gerbang di cabang `fs1/1.1a-redproof` (PR draf terpisah, ditutup dan cabangnya dihapus sesudahnya): berkas Go tak terformat → `server-fmt` merah; `err` tak diperiksa → `server-lint` merah; uji gagal → `server-test` merah; dua blok identik melebihi ambang → `cpd-go` merah. Untuk `server-vuln`: jalankan lokal di modul sementara di luar repositori yang mengimpor versi rentan yang dikenal (mis. versi lama `golang.org/x/net` dengan advisori yang fungsinya dipanggil) → merah. Catat tautan run CI merah di laporan.
- [ ] **Langkah 3: perlindungan cabang.** Lead (bukan pelaksana) mencoba mewajibkan job `server` hijau sebelum merge ke `main` lewat `gh api` (aturan cabang atau ruleset). Bila GitHub menolak karena paket akun (repositori privat di akun gratis), catat di `_progress.md` bahwa gerbang ditegakkan oleh proses lead: merge hanya sesudah CI hijau dan temuan `code-reviewer` selesai (`CLAUDE.md` "Git").
- [ ] **Langkah 4: commit dan PR siap ditinjau.** Pesan commit contoh: `Add repository skeleton and server CI gates (QG-01, QG-02, QG-03, QG-05; ADR-025)`. Deskripsi PR memuat tabel dependensi/alat (versi, tanggal cek, tautan, alasan) dan tautan run merah dari langkah 2.

## Selesai bila

- CI hijau di PR; setiap gerbang terbukti merah sekali (tautan tercatat).
- Catatan ADR-025 untuk alat rahasia ditulis `architect` di cabang yang sama.
- `code-reviewer` dua lintasan tanpa temuan terbuka; lead merge ke `main`.
- ES §9 butir 5 (dijalankan terhadap perangkat atau API nyata) **tidak berlaku** untuk tugas tanpa perilaku pengguna; dicatat demikian di PR.
