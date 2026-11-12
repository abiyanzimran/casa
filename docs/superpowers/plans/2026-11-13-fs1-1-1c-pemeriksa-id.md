# FS-1 1.1c — Pemeriksa ID (QG-06, 14 §4.13): rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** program Go kecil yang membaca pengenal yang harus terbukti dari dokumen perencanaan dan nama uji di repositori, lalu gagal bila ada uji dengan ID yang tidak ada, atau ID dalam cakupan langkah yang sudah selesai tanpa uji.

**Arsitektur:** `server/cmd/idcheck` (package `main`, tanpa dependensi baru). Tiga bagian dalam satu paket: pembaca dokumen, pemindai nama uji, dan pembanding. Cakupan per langkah ada di `traceability/scope.txt`. Dijalankan lewat target `make idcheck` dan langkah CI di job `server`.

**Stack:** Go stdlib (`regexp`, `bufio`, `go/parser` untuk berkas Go, `io/fs`).

**Spesifikasi:** 14 §4.13 dan §5.1 (QG-06); ES §6 ("Every Tier 1 acceptance criterion ... has an automated test that names its requirement ID").

**Pelaksana:** `backend-dev`. **Peninjau:** `code-reviewer`. **Cabang:** `fs1/1.1c-idcheck` dari `main` sesudah 1.1a merge.

## Batasan global

- Batasan global 1.1a berlaku. Tanpa dependensi baru.
- **Bentuk nama uji yang dihitung** (satu bentuk per bahasa; bentuk lain tidak dihitung):
  - Go: argumen pertama `t.Run` yang berupa string literal dan **diawali** ID, mis. `t.Run("FR-FAM-001/AC-1 creator gets a household", …)`. Dibaca dengan `go/parser`, bukan regex atas teks.
  - TypeScript (Vitest, Playwright): judul string literal `test(` atau `it(` yang diawali ID, termasuk `test.only`/`test.skip` (yang terakhir dilaporkan sebagai peringatan dan **tidak** dihitung sebagai bukti).
  - Swift dan Kotlin: ditambahkan saat jalurnya mulai (§3.8, §3.9); tidak di tugas ini.
- **Pola ID di awal nama uji:** `FR-[A-Z]+-\d{3}/AC-\d+`, `INV-A\d{2}`, `INV-\d{2}`, `PG-\d{2}`, `CTH-\d{2,}`, `SCN-\d{2}`, `SEC-\d{2,3}`, diikuti spasi, titik dua, atau akhir string. Satu uji boleh menyebut satu ID saja.

## Sumber ID (14 §4.13 butir 1)

| Jenis | Dokumen | Dikenal (boleh dirujuk uji) | Wajib (cakupan dan rilis) | Angka pemeriksa |
|---|---|---|---|---|
| FR/AC | `docs/05-functional-requirements/{FAM,FIN-1,FIN-2,REQ,COM,NTF}.md`, `tier-2.md` | Semua `FR-…/AC-n`: bullet `- **AC-n**` di bawah heading `### FR-XXX-NNN — …` | Hanya FR Tier 1 (FR yang baris indeksnya tidak bertanda "(Tier 2" atau "(Tier 3") | Jumlah FR Tier 1 = **190** (14 §3.3: "semua 190 FR Must Tier 1 di FS-1") |
| INV | `docs/07-financial-domain.md`, baris tabel `\| INV-NN \|` | INV-01 sampai INV-41 | Semua | **41** |
| INV-A | `docs/06-activity-approval-architecture.md`, baris `\| INV-ANN \|` | INV-A01 sampai INV-A53 | Semua | **53** |
| PG | `docs/06-…`, baris `\| PG-NN \|` | PG-01 sampai PG-15 | PG-01 sampai PG-14 (PG-15 sejak FS-3, 14 §4.2) | Dikenal **15**, wajib **14** |
| CTH | `docs/07-…`, heading `#### CTH-NN — …` | Semua | Semua | Dilaporkan; lead mencocokkan dengan 07 §7 |
| SCN | `docs/12-scenarios.md` §2, tabel berkolom `\| SCN \| Judul \| Tier \|` | Semua | Kolom Tier = 1 | Dilaporkan; lead mencocokkan dengan 12 |
| SEC | `docs/10-security-privacy/*.md`, tabel berkolom `\| ID \| Persyaratan \| Untuk \| Tahap \|` | Semua `SEC-NN` | Tahap memuat "Pilot" **dan** kolom Untuk berlaku bagi server atau klien web (bukan khusus iOS atau Android native; 10 indeks "Pilot per klien") | Dilaporkan |

Nilai kolom "Untuk" yang berbeda-beda dikumpulkan dan dicetak oleh `idcheck -list-untuk`. Pelaksana menulis pemetaan nilai → berlaku bagi web/server atau tidak di satu peta di kode. **Nilai yang maknanya tidak jelas: berhenti dengan status NEEDS_CONTEXT dan sebut nilainya**; lead memutuskan.

## Fokus tinjauan

1. Parser dokumen yang diam-diam menemukan 0 ID karena format berubah: setiap jenis punya angka pemeriksa atau batas bawah, dan program gagal (bukan lulus) bila sebuah jenis kosong.
2. ID di dalam blok kode markdown atau di teks biasa (rujukan silang) yang ikut terhitung sebagai definisi: hanya baris definisi menurut tabel di atas.
3. Uji dengan ID salah ketik (`INV-1`, `FR-FAM-01/AC-1`) yang lolos karena tidak cocok pola: nama uji yang **mirip** ID (diawali `FR-`, `INV`, `PG-`, `CTH-`, `SCN-`, `SEC-`) tetapi tidak cocok pola lengkap juga gagal.
4. `test.skip` yang dihitung sebagai bukti.
5. Uji di `node_modules/`, `.svelte-kit/`, atau kode hasil generate yang ikut dipindai.

---

### Tugas 1: Pembaca dokumen

**Berkas:** Buat `server/cmd/idcheck/docs.go`, `server/cmd/idcheck/docs_test.go`, `server/cmd/idcheck/testdata/docs/…` (cuplikan **asli** yang disalin dari dokumen, cukup untuk setiap bentuk baris, termasuk satu FR Tier 2 dan satu SEC khusus Android).

**Antarmuka (dipakai Tugas 3):**
- `type Requirement struct { ID string; Required bool }`
- `func readRequirements(docsDir string) ([]Requirement, error)` — galat bila satu jenis kosong, atau bila jumlah FR Tier 1, INV, INV-A, PG tidak sama dengan angka pemeriksa.

- [ ] **Langkah 1: uji gagal** atas `testdata/docs`: ID yang diharapkan persis (daftar ditulis di uji), FR Tier 2 dikenal tapi tidak wajib, SEC Android dikenal tapi tidak wajib, ID di blok kode tidak terhitung.
- [ ] **Langkah 2:** `go test ./cmd/idcheck/` → GAGAL.
- [ ] **Langkah 3:** tulis `readRequirements`. Angka pemeriksa 190, 41, 53, 15/14 ditulis sebagai konstanta bernama dengan komentar sumbernya; fungsi pemeriksa angka dilewati hanya untuk `testdata` lewat parameter, bukan variabel global.
- [ ] **Langkah 4:** uji kedua atas `../docs` sungguhan: angka pemeriksa cocok. → LULUS. **Bila tidak cocok, jangan ubah angka pemeriksa**; laporkan selisihnya (lead menentukan apakah dokumen atau parser yang salah).

### Tugas 2: Pemindai nama uji

**Berkas:** Buat `server/cmd/idcheck/tests.go`, `server/cmd/idcheck/tests_test.go`, `server/cmd/idcheck/testdata/repo/…` (berkas `_test.go` dan `.spec.ts`/`.test.ts` contoh).

**Antarmuka:**
- `type TestRef struct { ID, File string; Line int; Skipped bool }`
- `func scanTests(repoDir string) (refs []TestRef, malformed []TestRef, err error)` — memindai `server/**/*_test.go` dan `web/src/**/*.test.ts`, `web/tests/**/*.spec.ts`; melewati `node_modules`, `.svelte-kit`, `build`, dan direktori `gen`.

- [ ] **Langkah 1: uji gagal:** `t.Run` literal berawalan ID terbaca; `t.Run` dengan variabel tidak terbaca; `test('INV-01 …')`, `it("SEC-158: …")`, `test.skip('PG-01 …')` (Skipped), `test('INV-1 …')` masuk `malformed`.
- [ ] **Langkah 2:** GAGAL → tulis → LULUS.

### Tugas 3: Pembanding, cakupan, dan CLI

**Berkas:** Buat `server/cmd/idcheck/main.go`, `server/cmd/idcheck/main_test.go`, `traceability/scope.txt`. Ubah `Makefile`, `.github/workflows/ci.yml`.

**Antarmuka:** `func run(args []string, stdout, stderr io.Writer) int`, bendera `-docs`, `-repo`, `-scope`, `-release`, `-list-untuk`.

- [ ] **Langkah 1: uji gagal** untuk keempat aturan: (a) uji dengan ID tak dikenal → kode 1; (b) nama mirip ID (malformed) → kode 1; (c) ID di `scope.txt` tanpa uji (atau hanya uji `skip`) → kode 1; (d) dengan `-release`, ID wajib mana pun tanpa uji → kode 1; selain itu → kode 0. Keluaran menyebut ID, berkas, dan baris.
- [ ] **Langkah 2:** GAGAL → tulis → LULUS.
- [ ] **Langkah 3:** `traceability/scope.txt`: satu ID per baris, `#` komentar; isi awal hanya komentar kepala yang menjelaskan bahwa daftar diperbarui di akhir setiap langkah 14 §3.3 (14 §4.13 butir 2). Baris ID di scope yang bukan ID wajib dikenal → kode 1.
- [ ] **Langkah 4:** target `idcheck` (`go run ./cmd/idcheck -docs ../docs -repo .. -scope ../traceability/scope.txt` dari `server/`) masuk `check-server`; langkah CI sendiri di job `server`.
- [ ] **Langkah 5:** PR draf; CI HIJAU. Bukti merah di cabang buangan: satu uji bernama `t.Run("INV-99 x", …)` → merah. Catat tautan run.
- [ ] **Langkah 6:** laporan memuat angka per jenis (dikenal, wajib) dari dokumen sungguhan dan daftar nilai kolom "Untuk" beserta pemetaannya. Pesan commit contoh: `Add ID checker for test traceability (QG-06; 14 §4.13)`.

## Selesai bila

- CI hijau; satu bukti merah tercatat; angka per jenis dilaporkan dan dicocokkan lead dengan dokumen (CTH, SCN, SEC dihitung ulang lead dengan skrip terpisah).
- `code-reviewer` tanpa temuan terbuka; lead merge.
