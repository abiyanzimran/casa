# FS-1 1.3 — Kontrak dasar dan rangka Connect: rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** kontrak Protobuf pertama di `proto/` (pesan dasar yang dibawa setiap RPC, kode galat, tabel jenis perintah), gerbang QG-04 (`buf format`, `buf lint`, `buf breaking`, kode hasil generate identik) di CI, kode hasil generate Go dan TypeScript yang di-commit, subperintah `casa serve` dengan rangka Connect dan `/healthz`, pemeriksaan batas di setiap permintaan (konteks permintaan, versi klien minimum `web`, kunci idempotensi perintah, protovalidate), dan uji kontrak 14 §4.6.

**Arsitektur:** kontrak satu sumber di `proto/casa/v1/` (workspace `buf.yaml` v2 di akar repositori, modul `proto`). `buf generate` menulis Go ke `server/internal/gen/casa/v1` dan TypeScript ke `web/src/lib/gen/`; keduanya di-commit, dan CI membangkitkan ulang lalu menolak beda (14 §4.6 butir 2). Paket baru `server/internal/api` memegang server Connect (`connectrpc.com/connect/v2`) dengan satu interseptor batas, mux HTTP dengan `/healthz` (09.4 §3), dan batas ukuran badan permintaan. `cmd/casa` menambah `serve`. Belum ada RPC produk: kabel Connect dibuktikan dengan metode uji sintetis yang dibangun dari deskriptor dinamis di dalam uji, lewat tumpukan HTTP yang sungguhan.

**Stack:** buf CLI (biner rilis resmi, SHA-256 dipatok di `Makefile`), `google.golang.org/protobuf` (runtime dan `protoc-gen-go` lewat direktif `tool` di `server/go.mod`), `connectrpc.com/connect/v2`, `buf.build/go/protovalidate` v1.x, modul BSR `buf.build/bufbuild/protovalidate` (dikunci `buf.lock`), `@bufbuild/protobuf` 2.x dan `@bufbuild/protoc-gen-es` 2.x (versi sama persis).

**Spesifikasi:** 14 §2.2 (kolom Protobuf/Connect), §3.3 langkah 1, §4.6, §4.7 (baris `NOTICE_UNREAD`: uji kontrak `protective` ⊆ `notice_exempt`), §5.1 QG-04; ADR-013 butir 4, 5 dan catatan web; ADR-005; ADR-028 butir 3, 5 dan catatan web butir 5; ADR-030 butir 3, 5, 8, 9; 09.1 §4; 09.2 §1, §4, §8; 09.3 §1 (tabel jenis perintah, `WRITE_FROZEN`, `NOTICE_UNREAD`, jeda urungkan); 09.4 §2 (`user_initiated`), §3 (`/healthz`), §4 (baris "Kontrak dan bantuan isian"), §5; SEC-169 (d) (10.5); `ENGINEERING-STANDARDS.md` §1, §2, §4, §6, §7, §8; rencana 1.1a ("`buf lint`/`buf breaking` dan kontrak: tugas 1.3").

**Pelaksana:** `backend-dev`. **Peninjau:** `code-reviewer` (dua lintasan). **Cabang:** `fs1/1.3-kontrak-dasar-connect` dari `main`.

## Keputusan yang diambil rencana ini

1. **Kode hasil generate di-commit, dan CI membangkitkan ulang serta menolak beda.** Dasarnya 14 §4.6 butir 2 ("dibangkitkan ulang di CI dan harus identik dengan yang di repositori") dan 14 §2.2 ("`buf generate` di CI; hasil generate tidak disunting tangan"). Akibat yang dikehendaki: `go build`, `go list` (uji arsitektur), govulncheck, golangci-lint, `svelte-check`, dan ESLint berjalan tanpa buf; `web-dev` memakai tipe tanpa memasang buf; perubahan kontrak terlihat utuh di diff PR.
2. **Tata letak.** `buf.yaml` dan `buf.gen.yaml` di akar (workspace v2, `modules: [{path: proto}]`, `deps: [buf.build/bufbuild/protovalidate]`), `buf.lock` di akar. Paket proto `casa.v1`. Go: managed mode `go_package_prefix: github.com/abiyanzimran/casa/server/internal/gen`, dimatikan untuk modul `buf.build/bufbuild/protovalidate` (paket Go-nya tetap paket BSR yang diimpor protovalidate-go), `paths=source_relative`, keluaran `server/internal/gen` → paket `casav1`. TypeScript: `protoc-gen-es` `target=ts`, `include_imports: true` (berkas `buf/validate/validate_pb.ts` ikut dibangkitkan; tipe WKT dari `@bufbuild/protobuf/wkt`), keluaran `web/src/lib/gen`.
3. **Plugin lokal, versinya dari lockfile yang sudah ada.** `protoc-gen-go` dijalankan `go -C server tool protoc-gen-go` (direktif `tool`, sehingga versinya sama dengan runtime `google.golang.org/protobuf`; ES §8 "Compatible as a set"). `protoc-gen-es` dari `web/node_modules/.bin` (dev dependency, versi sama persis dengan `@bufbuild/protobuf`). Tanpa plugin jarak jauh BSR. `protoc-gen-connect-go` dan `@connectrpc/connect`/`connect-web` **belum** dipasang: belum ada `service` di kontrak, jadi keduanya tidak menghasilkan atau dipakai apa pun (ES §1 "no unused dependencies"); keduanya datang bersama layanan pertama (1.4) dan 1.9b.
4. **Pesan dasar** (nomor medan final; `buf breaking` menjaganya sejak merge):

   ```proto
   // casa/v1/request.proto
   enum ClientKind { CLIENT_KIND_UNSPECIFIED = 0; CLIENT_KIND_IOS = 1; CLIENT_KIND_ANDROID = 2; CLIENT_KIND_WEB = 3; }
   // Carried as field 1 of every RPC request (contract test, Task 3).
   message RequestContext {
     ClientKind client_kind = 1;   // enum: defined_only, not_in [0]          (ADR-013 web note; 09.2 §5.1 session.client_kind)
     string client_version = 2;    // pattern per Q1                          (ADR-030 item 9; SEC-169 (d))
     bool user_initiated = 3;      // shared-mode idle timer                  (09.4 §2; ADR-030 item 5; SEC-165 (g))
     string idempotency_key = 4;   // UUIDv4, case-insensitive; IGNORE_IF_ZERO_VALUE; required for commands by the server (ADR-005; 09.2 §8)
   }
   // casa/v1/error.proto — the one domain error enum (ADR-013 item 4); clients act on ErrorDetail.code, never on the Connect code.
   enum ErrorCode {
     ERROR_CODE_UNSPECIFIED = 0;
     ERROR_CODE_CLIENT_UPGRADE_REQUIRED = 1; // ADR-028 item 5 + web note; ADR-030 item 9; SEC-169 (d)
     ERROR_CODE_SESSION_EXPIRED = 2;         // "session_expired": ADR-030 item 5; SEC-161 (d)
     ERROR_CODE_SESSION_REVOKED = 3;         // "session_revoked": ADR-030 item 5; SEC-172
     ERROR_CODE_WRITE_FROZEN = 4;            // ADR-028 item 3; 09.3 §1 (retryable)
     ERROR_CODE_NOTICE_UNREAD = 5;           // SEC-34; 09.3 §1 (retryable)
   }
   message ErrorDetail { ErrorCode code = 1; }
   // casa/v1/command.proto — the command kind table (09.3 §1; ADR-028 item 3), one entry per command RPC.
   message CommandOptions {
     optional uint32 hold_seconds = 1; // undo window (09.3 §1; 08 §9.3); presence required by the contract test
     bool protective = 2;              // SEC-151 (1)-(7); exempt from WRITE_FROZEN and NOTICE_UNREAD
     bool notice_exempt = 3;           // SEC-34; 09.3 §1
   }
   extend google.protobuf.MethodOptions { CommandOptions command = 50000; } // in-house extension range
   ```

   Pola `idempotency_key`: `^(?i)[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$` (UUIDv4 menurut 09.2 §8; huruf besar diterima karena `UUID().uuidString` Swift menulis huruf besar). Nama kode di dokumen (`session_expired`, `WRITE_FROZEN`) menjadi nilai enum berawalan `ERROR_CODE_` karena aturan `buf lint` STANDARD; kode lain (`state_changed`, `terms_superseded`, `idempotency_conflict`, `currency_locked`) ditambahkan oleh tugas yang pertama kali mengembalikannya (penambahan nilai enum lolos `buf breaking`).
5. **"Perintah tulis" = metode yang membawa opsi `(casa.v1.command)`** (tabel jenis perintah 09.3 §1, yang mencakup perintah antrean dan perintah daring biasa). RPC `identity` di luar tabel (masuk, autentikasi ulang, pembaruan token, pemulihan) bukan perintah tulis di sini. Dikonfirmasi lead (Q4).
6. **Satu interseptor batas, urutan tetap:** (1) `Spec.Schema` harus `protoreflect.MethodDescriptor`, bila tidak → `CodeInternal`, handler tidak dipanggil; (2) `request_context` ada dan lolos protovalidate, bila tidak → `CodeInvalidArgument`; (3) bila metode perintah dan `client_kind = WEB` dan `client_version` < minimum `web` → `CodeFailedPrecondition` dengan `ErrorDetail{code: CLIENT_UPGRADE_REQUIRED}`; (4) bila metode perintah dan `idempotency_key` kosong → `CodeInvalidArgument`; (5) protovalidate atas seluruh pesan → `CodeInvalidArgument`. Galat validasi membawa detail `buf.validate.Violations` (pesan kontrak, ter-generate untuk semua bahasa). Versi diperiksa **sebelum** validasi badan, supaya klien lama yang bentuk pesannya sudah tidak lolos aturan baru tetap menerima kode yang dapat dicoba ulang dan perintahnya tetap di antrean (SEC-169 (d); pola SEC-34 (4)), bukan penolakan yang membuang perintah.
7. **Tanpa RPC produk.** Dokumen tidak menamai RPC versi atau kesehatan; yang ada hanya `/healthz` HTTP biasa untuk pemantau uptime (09.4 §3; 14 §6.4, §7.3), yang disajikan di sini (200, badan `ok\n`, tanpa data, tanpa basis data). Kabel Connect dibuktikan di uji dengan layanan `casa.apitest.v1.ProbeService` yang dibangun dari `descriptorpb` di berkas uji dan didaftarkan lewat `connect.Server.Register`, dipanggil lewat `httptest.Server` + `connecthttp.Mount` + `connecthttp.NewTransport`. Klien Go hasil generate teruji sejak layanan pertama (1.4).
8. **Versi minimum `web`** (Q1, Q2): rencana ini memakai rekomendasi `MAJOR.MINOR.PATCH` numerik (pola `^(0|[1-9][0-9]{0,5})\.(0|[1-9][0-9]{0,5})\.(0|[1-9][0-9]{0,5})$`, dibandingkan per bagian) dan konstanta di biner (`api.MinimumWebVersion`, awal `0.0.0`, dinaikkan pada rilis yang memperbaiki celah; ADR-028 catatan web "konstanta versi minimum per platform ... dibangun sekarang"). `ios` dan `android` tanpa minimum di pilot (ADR-028 catatan web butir 5). Bila lead memutus lain, yang berubah hanya pola, `ParseVersion`, dan sumber nilai; urutan dan uji tetap.
9. **Batas badan permintaan 1 MiB** `[Usulan]` lewat `http.MaxBytesHandler` (stdlib) di `NewMux`; dinaikkan per jalur oleh tugas `files` bila perlu (Q9).

## Batasan global

- Batasan global 1.1a dan 1.2 berlaku: setiap alat dan dependensi dicek live saat dipasang (registri atau halaman rilis resmi, catatan rilis, advisori, isu kritis terbuka), stabil terbaru, dipatok tepat, dicatat di deskripsi PR (versi, tanggal cek, tautan, alasan) (ES §8). Pengecekan perencana 13 November 2026 (buf 1.75.0, connect-go v2.0.2, protobuf-go v1.36.12, protovalidate-go v1.4.0, `@bufbuild/protobuf` dan `protoc-gen-es` 2.17.0 terbit 5 November 2026) **bukan** versi untuk dipasang.
- Dependensi Go langsung baru hanya: `google.golang.org/protobuf`, `connectrpc.com/connect/v2`, `buf.build/go/protovalidate`, dan modul BSR `buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go` (diimpor kode hasil generate, jadi langsung menurut ADR-023 butir 6.6; Q8). **Tidak** `connectrpc.com/validate` (v0.x, tidak di ADR-023, dan untuk connect-go v1): interseptor validasi ditulis sendiri, beberapa baris di atas `protovalidate.Validate`. Modul tak langsung baru (mis. `github.com/google/cel-go`, `cel.dev/expr`) dicatat menurut ADR-023 butir 6.
- Dependensi web baru hanya `@bufbuild/protobuf` (dependensi bundel, tercantum di ADR-030 butir 8) dan `@bufbuild/protoc-gen-es` (pengembangan), lewat `pnpm add -E`; masa tunggu 7 hari ditegakkan `web/pnpm-workspace.yaml`.
- Biner buf diunduh lewat `scripts/fetch-verified.sh` dengan SHA-256 dari `sha256.txt` rilisnya, pola yang sama dengan golangci-lint (`linux-amd64` saja, komentar `ponytail:` yang sama).
- Hasil generate tidak pernah disunting tangan (ES §2). `cpd.sh` dan `idcheck` sudah melewati direktori `gen` dan `*.pb.go`/`*_pb.ts`; golangci-lint v2 melewati berkas bertanda generated.
- Tanpa log baru; keluaran `casa serve` hanya alamat dengar dan galat tanpa isi permintaan (ES §4; 09.4 §3).
- Nama uji yang menyebut ID diawali ID-nya (`SEC-169 (d) …`). Uji kontrak yang masih hampa (belum ada perintah) **tidak** diawali `SEC-151` atau `SEC-34`, supaya QG-06 tidak menganggap ID itu sudah terbukti; ID-nya dikutip di komentar.
- Komentar, pengenal, dan pesan commit berbahasa Inggris; setiap aturan di kode mengutip ID atau bagian dokumennya; pesan commit menyebut ID (ES §1, §7).

## Fokus tinjauan

1. **Klien lama yang pesannya tidak lagi lolos validasi terbaru** mendapat `InvalidArgument` dan perintahnya dibuang, padahal seharusnya `CLIENT_UPGRADE_REQUIRED` dan tetap di antrean: urutan keputusan 6; uji di Tugas 4.
2. **Versi dibandingkan sebagai teks** (`"1.10.0" < "1.9.0"`): uji numerik di Tugas 4.
3. **Kode hasil generate yang menyimpang dari kontrak** (proto diubah tanpa generate ulang, berkas hasil disunting tangan, atau berkas baru hasil generate yang belum di-commit): `proto-generated` memeriksa `git diff` **dan** berkas tak terlacak; dibuktikan merah di Tugas 6.
4. **Metode tanpa deskriptor** (didaftarkan tanpa `Spec.Schema`) diam-diam dianggap bacaan sehingga lolos pemeriksaan versi dan kunci idempotensi: gagal tertutup, uji di Tugas 4.
5. **Gerbang `buf breaking` yang hijau karena tidak membandingkan apa pun**: penjaga kontrak-pertama hanya lolos bila ref pembanding ada dan belum memuat `buf.yaml`; ref yang tidak ada gagal. Penjaga dihapus dan gerbang dibuktikan merah sesudah merge (Tugas 7).

## Yang sengaja tidak di tugas ini

- RPC `identity`, sesi, cookie, `session_expired`/`session_revoked` sebagai perilaku (1.4, 1.5); pemeriksaan CSRF dan `Origin`, penolakan Connect `GET` bagi sesi cookie, CSP dan header, bundel tersemat (1.5); `household` dan `NOTICE_UNREAD` sebagai perilaku (1.6); `read`, `sync`, `jobs`, `GET /v1/stream` dan bingkai SSE (1.7); kabel klien web (connect-web, `client_kind`/`client_version` di permintaan, layar muat ulang): 1.9b.
- `WRITE_FROZEN` sebagai perilaku (`casa ops write-freeze`, ADR-028 butir 3): tugas tahan tulis. Di sini hanya kodenya.
- Tabel `idempotency_record` dan pencatatan hasil perintah (09.2 §8): bersama perintah tulis pertama (1.4 atau 1.6). Di sini hanya kunci di kontrak dan kewajibannya bagi perintah.
- Pemakaian `user_initiated` oleh server (pengatur waktu diam sesi mode bersama): 1.5. Di sini hanya medannya.
- `protoc-gen-connect-go`, pembantu server uji untuk uji AC lewat klien Go hasil generate (14 §4.6 butir 4): bersama layanan pertama (1.4).
- TLS, autocert, unit systemd `casa serve`, server uji sementara (14 §7.8): langkah 7 dan WS-5.
- Pembangkit vektor emas bantuan isian (14 §3.3 langkah 1; §4.5): tidak termasuk daftar 1.3 dari lead; tugasnya ditetapkan lead (laporan).
- Penambahan ID ke `traceability/scope.txt`: lead, di akhir langkah 1 (14 §4.13).

## Berkas yang mungkin bentrok dengan 1.9a (`web-dev`, cabang `fs1/1.9a-inti-klien-web`)

`web/package.json` dan `web/pnpm-lock.yaml` (1.9a menambah `dexie`, `fake-indexeddb`, `vitest`, medan `imports`), `Makefile` (1.9a mengubah baris `.PHONY` web, `check-web`, `web-check`, menambah `web-unit`), `.github/workflows/ci.yml` (1.9a menambah langkah `web-unit` di job `web`). Tugas ini tidak menyentuh baris-baris itu: `.PHONY` kontrak ditulis di baris sendiri, target kontrak ditambahkan sesudah target server, job `contract` ditambahkan di akhir berkas, dan di `web/` hanya `src/lib/gen/` (baru), `.prettierignore`, `package.json`, dan `pnpm-lock.yaml`. Bila 1.9a lebih dulu masuk `main`, cabang ini di-rebase dan lockfile dibangkitkan ulang dengan `pnpm install`, bukan disunting tangan.

---

### Tugas 1: Modul `proto/`, kontrak dasar, gerbang `buf format`, `buf lint`, `buf breaking`

**Berkas:**
- Buat: `buf.yaml`, `buf.lock`, `proto/casa/v1/request.proto`, `proto/casa/v1/error.proto`, `proto/casa/v1/command.proto`
- Ubah: `Makefile` (variabel buf, target `proto-fmt`, `proto-lint`, `proto-breaking`, `check-contract`; `check` ikut memanggil `check-contract`), `.github/workflows/ci.yml` (job baru `contract`), `.gitignore` bila perlu

**Antarmuka:**
- Menghasilkan: paket proto `casa.v1` dengan `RequestContext`, `ClientKind`, `ErrorCode`, `ErrorDetail`, `CommandOptions`, ekstensi `command` (keputusan 4); target `make` `proto-fmt`, `proto-lint`, `proto-breaking`, `check-contract`; variabel `BUF`, `BUF_AGAINST_REF ?= refs/remotes/origin/main`.

- [ ] **Langkah 1: cek live** buf CLI (https://github.com/bufbuild/buf/releases; catatan rilis, isu terbuka berlabel bug pada `breaking`/`generate`), SHA-256 `buf-Linux-x86_64` dari `sha256.txt` rilis, dan versi modul BSR `buf.build/bufbuild/protovalidate` yang selaras dengan protovalidate-go terbaru.
- [ ] **Langkah 2: buktikan gerbang tanpa kontrak gagal.** Tulis dulu target `make` dan `buf.yaml` tanpa berkas `.proto`. `make proto-lint` → GAGAL dengan `Module "path: "proto"" had no .proto files` (dilihat perencana dengan buf 1.75.0). Gerbang yang tidak menemukan berkas tidak pernah hijau.
- [ ] **Langkah 3: target `make`.**
  - `BUF_VERSION`, `BUF_SHA256`, `BUF := $(CURDIR)/.cache/buf-$(BUF_VERSION)/buf`, aturan unduh lewat `fetch-verified.sh`.
  - `proto-fmt`: `$(BUF) format --diff --exit-code` (14 §2.2 baris Formatter, QG-01).
  - `proto-lint`: `$(BUF) lint` dengan aturan bawaan STANDARD (tanpa pengecualian di `buf.yaml`).
  - `proto-breaking`: `$(BUF) breaking --against '.git#ref=$(BUF_AGAINST_REF)'` dengan aturan bawaan `FILE` (paling ketat; JSON connect-web peka nama medan). **Penjaga kontrak pertama** sebelum perintah itu: `git rev-parse --verify -q '$(BUF_AGAINST_REF)^{commit}'` gagal → keluar 1 dengan `proto-breaking: $(BUF_AGAINST_REF) not found; fetch it first`; ref ada tetapi `git cat-file -e '$(BUF_AGAINST_REF):buf.yaml'` gagal → cetak `proto-breaking: no contract on $(BUF_AGAINST_REF) yet (first contract, FS-1 1.3)` dan keluar 0. Komentar: penjaga dihapus sesudah merge (Tugas 7).
  - `.PHONY` kontrak di baris sendiri; `check-contract: proto-fmt proto-lint proto-breaking` (Tugas 2 menambah `proto-generated`); `check: check-server check-web check-contract`.
- [ ] **Langkah 4: job CI `contract`** di akhir `ci.yml`: `actions/checkout` (SHA yang sama dengan job lain) dengan `fetch-depth: 0` dan `persist-credentials: false`; langkah `proto-fmt (QG-01)`, `proto-lint (QG-04)`, `proto-breaking (QG-04)`, masing-masing `if: ${{ !cancelled() }}` dan memanggil satu target `make`. (Tugas 2 menambah Go, pnpm, Node, dan `proto-generated`.)
- [ ] **Langkah 5: tulis tiga berkas `.proto`** persis keputusan 4, dengan komentar dokumen yang mengutip ID pada setiap medan dan nilai, dan aturan protovalidate sebagai anotasi `(buf.validate.field)`: `client_kind` `enum = {defined_only: true, not_in: [0]}`; `client_version` `string.pattern` keputusan 8; `idempotency_key` `string.pattern` keputusan 4 dengan `ignore = IGNORE_IF_ZERO_VALUE`. `buf dep update` menulis `buf.lock`.
- [ ] **Langkah 6:** `make proto-fmt proto-lint proto-breaking` → LULUS; `proto-breaking` mencetak pesan kontrak pertama (`main` belum punya `buf.yaml`).
- [ ] **Langkah 7: commit** `Add the base contract and the buf format, lint and breaking gates (QG-04; ADR-013 item 4; ADR-005; ADR-028 items 3, 5; ADR-030 items 5, 9)`.

### Tugas 2: `buf generate` untuk Go dan TypeScript, gerbang kode hasil generate identik

**Berkas:**
- Buat: `buf.gen.yaml`; hasil generate `server/internal/gen/casa/v1/*.pb.go`, `web/src/lib/gen/casa/v1/*_pb.ts`, `web/src/lib/gen/buf/validate/validate_pb.ts`
- Ubah: `server/go.mod`, `server/go.sum` (runtime protobuf, modul BSR protovalidate, direktif `tool google.golang.org/protobuf/cmd/protoc-gen-go`), `web/package.json`, `web/pnpm-lock.yaml`, `web/.prettierignore` (`src/lib/gen/`), `Makefile` (`proto-generate`, `proto-generated`; `check-contract` memanggil `proto-generated`), `.github/workflows/ci.yml` (job `contract`: `setup-go`, `pnpm/action-setup`, `setup-node`, `make web-install`, langkah `proto-generated (QG-04)`), `server/internal/architecture/imports_test.go`

**Antarmuka:**
- Mengonsumsi: berkas `.proto` Tugas 1.
- Menghasilkan: paket Go `github.com/abiyanzimran/casa/server/internal/gen/casa/v1` (`casav1`) dengan `RequestContext`, `ClientKind_CLIENT_KIND_WEB`, `ErrorCode_ERROR_CODE_CLIENT_UPGRADE_REQUIRED`, `ErrorDetail`, `CommandOptions`, `E_Command`; modul TS `$lib/gen/casa/v1/*_pb` (dipakai 1.9b).

- [ ] **Langkah 1: cek live** `google.golang.org/protobuf`, `@bufbuild/protobuf` dan `@bufbuild/protoc-gen-es` (umur ≥ 7 hari, versi keduanya sama), modul Go BSR `buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go` yang komitnya sama dengan `buf.lock`.
- [ ] **Langkah 2: tulis uji yang gagal.** Tambahkan ke `allowedImports` hanya baris yang belum diizinkan, sehingga uji arsitektur gagal begitu paket hasil generate ada: jalankan `buf generate` (langkah 3) **sebelum** menambah baris → `make server-test` GAGAL dengan `internal/gen/casa/v1 is not declared in allowedImports (ADR-014 item 4)`.
- [ ] **Langkah 3: `buf.gen.yaml`** (keputusan 2, 3):

  ```yaml
  version: v2
  managed:
    enabled: true
    disable:
      - module: buf.build/bufbuild/protovalidate
    override:
      - file_option: go_package_prefix
        value: github.com/abiyanzimran/casa/server/internal/gen
  inputs:
    - directory: proto
  plugins:
    - local: ["go", "-C", "server", "tool", "protoc-gen-go"]
      out: server/internal/gen
      opt: paths=source_relative
    - local: web/node_modules/.bin/protoc-gen-es
      out: web/src/lib/gen
      opt: target=ts
      include_imports: true
  ```

  Target `proto-generate`: `$(BUF) generate --clean` (menghapus kedua direktori `out` lebih dulu, sehingga berkas usang dari proto yang dihapus ikut hilang). Target `proto-generated`: `proto-generate`, lalu `git diff --exit-code -- server/internal/gen web/src/lib/gen` dan `git status --porcelain -- server/internal/gen web/src/lib/gen` harus kosong; bila tidak, cetak `proto-generated: generated code differs from the contract; run make proto-generate and commit` dan keluar 1.
- [ ] **Langkah 4:** tambahkan `"internal/gen/casa/v1": {}` ke `allowedImports` (menunggu baris 09.2 §1, Q3). `.prettierignore` menambah `src/lib/gen/`. ESLint tidak diubah: `protoc-gen-es` menulis `/* eslint-disable */` di kepala berkas (diperiksa di langkah 5; bila tidak, berhenti dan laporkan, karena `eslint.config.js` sedang diubah 1.9a).
- [ ] **Langkah 5:** `make proto-generated` → LULUS; `make check-server` → LULUS; `make web-install web-fmt web-lint web-check` → LULUS (TS hasil generate terkompilasi di bawah `svelte-check`).
- [ ] **Langkah 6: commit** `Generate Go and TypeScript from the contract and gate on drift (QG-04; 14 §4.6; ES §2)`.

### Tugas 3: Uji kontrak (14 §4.6, 09.3 §1)

**Berkas:**
- Buat: `server/internal/api/contract_test.go`, `server/internal/api/probe_test.go` (pembantu deskriptor dinamis, dipakai juga Tugas 4)
- Ubah: `server/go.mod`, `server/go.sum` bila perlu

**Antarmuka:**
- Mengonsumsi: `casav1.E_Command`, `casav1.CommandOptions`, `casav1.File_casa_v1_request_proto` (Tugas 2).
- Menghasilkan (berkas uji, `package api_test`):
  - `type probeMethod struct { name string; command *casav1.CommandOptions; clientStreaming bool; bareRequest bool }`
  - `func probeFile(t testing.TB, methods ...probeMethod) protoreflect.FileDescriptor`: `casa/apitest/v1/probe.proto`, paket `casa.apitest.v1`, mengimpor `casa/v1/request.proto`, `casa/v1/command.proto`, `google/protobuf/empty.proto`, `buf/validate/validate.proto`; pesan `ProbeRequest { casa.v1.RequestContext request_context = 1 [(buf.validate.field).required = true]; string note = 2 [(buf.validate.field).string.max_len = 3]; }` dan `BareRequest { string note = 1; }`; layanan `ProbeService`, metode menurut `methods` (respons `google.protobuf.Empty`; `bareRequest` memakai `BareRequest`). Dibangun dengan `protodesc.NewFile` atas `protoregistry.GlobalFiles`.
  - `func contractViolations(files []protoreflect.FileDescriptor) []string`: satu baris per pelanggaran, berbentuk `<service>.<method>: <rule>`.
  - `func casaFiles(t testing.TB) []protoreflect.FileDescriptor`: setiap berkas di `protoregistry.GlobalFiles` yang paketnya diawali `casa.` (berkas uji mengimpor paket `casav1` secara kosong agar terdaftar; komentar: paket proto baru menambah impornya di sini).

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
// TestContract holds the contract rules every RPC follows (14 §4.6; 09.3 §1;
// 09.1 §4). Each rule runs on the real contract (no violation) and on a probe
// file that breaks only that rule (exactly that violation), so a rule that
// checks nothing fails here, not later.
func TestContract(t *testing.T) {
	t.Run("the real contract has no violation", …)
	// contractViolations(casaFiles(t)) == nil
	cases := []struct{ name string; method probeMethod; want string }{
		{"every RPC request starts with a required request_context",
			probeMethod{name: "Bare", bareRequest: true},
			"casa.apitest.v1.ProbeService.Bare: request field 1 is not a required casa.v1.RequestContext named request_context"},
		{"no RPC streams from the client (09.1 §4; ADR-013 web note item 5)",
			probeMethod{name: "Upload", clientStreaming: true},
			"casa.apitest.v1.ProbeService.Upload: client streaming is not allowed"},
		{"every command declares hold_seconds",
			probeMethod{name: "NoHold", command: &casav1.CommandOptions{}},
			"casa.apitest.v1.ProbeService.NoHold: command without hold_seconds"},
		// 09.3 §1 "Uji kontrak"; SEC-151 is cited, not named: no protective command exists yet.
		{"a protective command has hold_seconds 0",
			probeMethod{name: "Leave", command: &casav1.CommandOptions{HoldSeconds: proto.Uint32(5), Protective: true, NoticeExempt: true}},
			"casa.apitest.v1.ProbeService.Leave: protective with hold_seconds 5"},
		// 09.3 §1 NOTICE_UNREAD "Uji kontrak: protective ⊆ notice_exempt"; SEC-34 cited, not named.
		{"a protective command is notice_exempt",
			probeMethod{name: "Revoke", command: &casav1.CommandOptions{HoldSeconds: proto.Uint32(0), Protective: true}},
			"casa.apitest.v1.ProbeService.Revoke: protective but not notice_exempt"},
	}
	// for each: contractViolations([]{probeFile(t, c.method)}) == []string{c.want}
}
```

  Uji itu juga memastikan probe yang patuh (`probeMethod{name: "Ok", command: &casav1.CommandOptions{HoldSeconds: proto.Uint32(5)}}`) tanpa pelanggaran.
- [ ] **Langkah 2:** `make server-test` → GAGAL (`contractViolations` belum ada).
- [ ] **Langkah 3:** tulis `probeFile`, `casaFiles`, `contractViolations` di berkas uji. `contractViolations` menelusuri setiap `Services()`/`Methods()`; `request_context` diperiksa lewat nama, nomor 1, `Message().FullName() == "casa.v1.RequestContext"`, dan `buf.validate.field` `required = true` pada opsinya.
- [ ] **Langkah 4:** `make server-test` → LULUS. Paket `internal/api` berisi uji saja sampai Tugas 4; tambahkan `"internal/api": {}` ke `allowedImports` (Q3).
- [ ] **Langkah 5: commit** `Add contract tests: request_context first, no client streaming, command kind table rules (14 §4.6; 09.3 §1; 09.1 §4)`.

### Tugas 4: Server Connect dengan interseptor batas dan versi klien minimum (`internal/api`)

**Berkas:**
- Buat: `server/internal/api/server.go`, `server/internal/api/version.go`, `server/internal/api/server_test.go`, `server/internal/api/version_test.go`
- Ubah: `server/go.mod`, `server/go.sum` (`connectrpc.com/connect/v2`, `buf.build/go/protovalidate`), `server/internal/architecture/imports_test.go`

**Antarmuka:**
- Mengonsumsi: `casav1` (Tugas 2); `probeFile`, `probeMethod` (Tugas 3).
- Menghasilkan, di `package api` (dipakai `cmd/casa` dan setiap layanan sejak 1.4):

```go
// Version is a client build version, MAJOR.MINOR.PATCH compared numerically (Q1).
type Version struct{ Major, Minor, Patch uint32 }
func ParseVersion(s string) (Version, error)
func (v Version) Less(w Version) bool

// MinimumWebVersion is the lowest web client version allowed to send a write
// command (ADR-030 item 9; ADR-028 web note item 5; SEC-169 (d)). Raised on a
// release that fixes a security hole (TM-45). Native clients have none in the pilot.
var MinimumWebVersion = Version{0, 0, 0}

// NewServer returns the Connect server every service registers on, with the
// boundary checks of every RPC in the order of plan decision 6.
func NewServer(minimumWebVersion Version) *connect.Server

// NewMux serves the procedures of server and GET /healthz (09.4 §3) behind a
// 1 MiB request body limit.
func NewMux(server *connect.Server) http.Handler
```

`NewServer` membuat satu `protovalidate.Validator` dan satu `connect.ServerInterceptor` yang membungkus `Receive` aliran: setelah pesan diterima, keputusan 6 dijalankan sebelum handler melihat pesan. `request_context` dibaca lewat `protoreflect` (medan bernama `request_context`), sehingga berlaku untuk pesan hasil generate maupun `dynamicpb`; pesan dinamis diubah ke `*casav1.RequestContext` lewat `proto.Marshal`/`proto.Unmarshal` (bukan `proto.Merge`, yang panik bila deskriptornya bukan objek yang sama). Pemetaan galat di satu fungsi: `CLIENT_UPGRADE_REQUIRED` → `connect.CodeFailedPrecondition` + `ErrorDetail`; validasi → `connect.CodeInvalidArgument` + `buf.validate.Violations`. Kode lain dipetakan oleh tugas yang memakainya.

- [ ] **Langkah 1: cek live** `connectrpc.com/connect/v2` (catatan rilis, `docs/v2-migration.md`, advisori) dan `buf.build/go/protovalidate` v1.x.
- [ ] **Langkah 2: tulis uji yang gagal.** `version_test.go`:

```go
func TestParseVersion(t *testing.T) {
	// valid: "0.0.0", "1.9.0", "1.10.0", "999999.0.1" → fields as written
	// invalid, each an error: "", "1.2", "1.2.3.4", "1.2.3-beta", "01.2.3", "1.2.x", " 1.2.3", "1000000.0.0"
}
func TestVersionLessComparesNumbers(t *testing.T) {
	// Less: {1,9,0}<{1,10,0} true; {1,10,0}<{1,9,0} false; {2,0,0}<{2,0,0} false; {1,9,9}<{2,0,0} true
}
```

  `server_test.go`: satu server per uji, `api.NewServer(api.Version{1, 9, 0})`, metode probe `Read` (tanpa opsi) dan `Write` (`command{hold_seconds: 5}`) didaftarkan lewat `server.Register(connect.Method{Spec: …, Handler: …})` dengan `Spec.Schema` = deskriptor metode dari `probeFile`; handler menghitung panggilan dan membalas `emptypb.Empty`. Dipanggil lewat `httptest.NewServer(api.NewMux(server))` dan `connect.NewClient(connecthttp.NewTransport(srv.Client(), srv.URL))` dengan codec proto. `rc(kind, version, key)` membangun `RequestContext`; `validKey` adalah UUIDv4 tetap.

```go
func TestBoundaryChecks(t *testing.T) {
	t.Run("SEC-169 (d) a write command from a web client below the minimum version is rejected with CLIENT_UPGRADE_REQUIRED and not applied", …)
	// Write, rc(WEB, "1.8.10", validKey) → connect.CodeOf(err) == CodeFailedPrecondition;
	// the one ErrorDetail decodes to code ERROR_CODE_CLIENT_UPGRADE_REQUIRED; Write handler calls == 0
	t.Run("SEC-169 (d) the minimum version is compared by number, not by text", …)
	// Write with "1.10.0" → OK; with "1.9.0" → OK (equal is allowed); Write handler calls == 2
	t.Run("SEC-169 (d) a read from a web client below the minimum version is served", …)
	// Read, rc(WEB, "0.1.0", "") → OK, Read handler calls == 1
	t.Run("SEC-169 (d) native clients have no minimum version in the pilot", …)
	// Write, rc(IOS, "0.0.1", validKey) and rc(ANDROID, "0.0.1", validKey) → OK (ADR-028 web note item 5)
	t.Run("SEC-169 (d) a below-minimum command the current contract finds invalid still gets CLIENT_UPGRADE_REQUIRED", …)
	// Write, rc(WEB, "1.0.0", validKey), note "toolong" → CodeFailedPrecondition + CLIENT_UPGRADE_REQUIRED, not InvalidArgument
	t.Run("a request without request_context never reaches the handler", …)
	// Read with no request_context → CodeInvalidArgument; Violations names field path "request_context"; calls == 0
	t.Run("a write command without an idempotency key never reaches the handler (ADR-005)", …)
	// Write, rc(WEB, "2.0.0", "") → CodeInvalidArgument; calls == 0
	t.Run("an idempotency key must be a UUIDv4 (ADR-005; 09.2 §8)", …)
	// "not-a-uuid" and a version-1 UUID → CodeInvalidArgument; validKey in upper case → OK
	t.Run("an unspecified client kind or a malformed client version is rejected", …)
	// Read with CLIENT_KIND_UNSPECIFIED; Read with "1.2" → CodeInvalidArgument each
	t.Run("the request body is validated after the version check", …)
	// Write, rc(WEB, "2.0.0", validKey), note "toolong" → CodeInvalidArgument; Violations names "note"; calls == 0
	t.Run("a method registered without its descriptor is rejected, never treated as a read", …)
	// Method with Spec.Schema nil → CodeInternal; handler calls == 0
	t.Run("GET /healthz answers ok without data (09.4 §3)", …)
	// GET → 200, body "ok\n"; POST /healthz → 405
	t.Run("a request body over 1 MiB is refused", …)
	// Read whose note makes the body 1 MiB + 1 byte → an error without a buf.validate.Violations detail
	// (the limit acts before validation); handler calls == 0
}
```

- [ ] **Langkah 3:** `make server-test` → GAGAL (`api.NewServer` belum ada).
- [ ] **Langkah 4:** tulis `version.go` dan `server.go`. Setiap aturan di kode mengutip ID (keputusan 4 sampai 9). `allowedImports`: `"internal/api": {"internal/gen/casa/v1"}` (Q3).
- [ ] **Langkah 5:** `make server-test` → LULUS; `make server-lint` → LULUS tanpa suppression baru.
- [ ] **Langkah 6: commit** `Add the Connect server with boundary checks and the web minimum client version (SEC-169 (d); ADR-030 item 9; ADR-028 item 5; ADR-005; ES §4)`.

### Tugas 5: Subperintah `casa serve`

**Berkas:**
- Buat: `server/cmd/casa/serve.go`, `server/cmd/casa/serve_test.go`
- Ubah: `server/cmd/casa/main.go` (sebaran `serve`; komentar paket), `server/internal/architecture/imports_test.go`

**Antarmuka:**
- Mengonsumsi: `api.NewServer`, `api.NewMux`, `api.MinimumWebVersion` (Tugas 4); `run` dari 1.2.
- Menghasilkan: `func serve(ctx context.Context, args []string, stdout, stderr io.Writer) int` di `package main`. Flag tunggal `-listen` (bawaan `127.0.0.1:8080`; TLS dan porta 443 di langkah 7). `http.Server{ReadHeaderTimeout: 10 * time.Second}` `[Usulan]`; saat `ctx` selesai, `Shutdown` dengan batas 10 detik, lalu kembali 0. Sesudah `net.Listen` berhasil, mencetak `casa serve: listening on <alamat sebenarnya>\n` ke `stdout`. Belum membuka basis data (belum ada RPC yang memakainya; `/healthz` tanpa data).

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestServe(t *testing.T) {
	t.Run("serves /healthz until the context ends", …)
	// run(ctx, {"serve", "-listen", "127.0.0.1:0"}) in a goroutine, stdout through io.Pipe;
	// read "casa serve: listening on 127.0.0.1:<port>\n"; GET /healthz → 200 "ok\n";
	// cancel ctx → run returns 0 within 15 s
	t.Run("fails when the address is in use", …)
	// net.Listen("tcp", "127.0.0.1:0") first, then serve on that address → 1;
	// stderr starts with "casa serve: listen tcp 127.0.0.1:"
	t.Run("rejects extra arguments", …)
	// {"serve", "extra"} → 2; stderr = "usage: casa serve [-listen address]\n"
}
```

- [ ] **Langkah 2:** `make server-test` → GAGAL.
- [ ] **Langkah 3:** tulis `serve.go`; `run` memanggil `serve` untuk `args[0] == "serve"`; komentar paket di `main.go` tidak lagi menyebut "serve comes with 1.3". `allowedImports`: `"cmd/casa": {"internal/api", "internal/schema"}`.
- [ ] **Langkah 4:** `make server-test` → LULUS. Jalankan biner sungguhan: `cd server && go run ./cmd/casa serve -listen 127.0.0.1:0`, lalu `curl -si http://127.0.0.1:<port>/healthz` → `HTTP/1.1 200` dan `ok`; `curl -si -X POST -H 'Content-Type: application/proto' --data-binary '' http://127.0.0.1:<port>/casa.v1.None/Nothing` → `404` (belum ada layanan); Ctrl-C → keluar 0. Catat keluarannya di PR (ES §9 butir 5).
- [ ] **Langkah 5: commit** `Add casa serve with the Connect skeleton and /healthz (09.4 §3; 14 §6.4; ADR-030 item 3)`.

### Tugas 6: README, bukti setiap gerbang merah, PR

- [ ] **Langkah 1: README akar** (bahasa Inggris): bagian "Contract": `proto/` adalah satu sumber; `make proto-generate` lalu commit hasilnya; `make check-contract`; `web-install` diperlukan untuk `protoc-gen-es`; hasil generate tidak disunting tangan. Bagian "Server": `casa serve -listen`.
- [ ] **Langkah 2:** push cabang, buka PR draf ke `main`. CI → HIJAU di job `server`, `web`, dan `contract`. Catat menit tertagih job `contract` (14 §4.3 "Anggaran CI").
- [ ] **Langkah 3: buktikan setiap gerbang dapat merah**, satu komit buangan per gerbang di cabang `fs1/1.3-redproof` (PR draf terpisah, ditutup dan cabangnya dihapus sesudahnya):

| Pelanggaran buangan | Yang harus merah |
|---|---|
| `.proto` dengan indentasi salah | `proto-fmt` |
| Nilai nol enum baru tidak berakhiran `_UNSPECIFIED` | `proto-lint` |
| Medan baru di `RequestContext` tanpa `make proto-generate` | `proto-generated` (diff) |
| Satu baris `*.pb.go` disunting tangan | `proto-generated` (diff) |
| Berkas `.proto` baru dibangkitkan tetapi `*_pb.ts`-nya tidak di-commit | `proto-generated` (berkas tak terlacak) |
| Layanan `casa.v1` sungguhan dengan metode `protective` ber-`hold_seconds` 5 | `TestContract/the_real_contract_has_no_violation` |
| Pemeriksaan versi dipindah sesudah validasi badan | `SEC-169 (d) a below-minimum command the current contract finds invalid still gets CLIENT_UPGRADE_REQUIRED` |
| Perbandingan versi lewat `strings.Compare` | `SEC-169 (d) the minimum version is compared by number, not by text` |
| `Spec.Schema` yang bukan deskriptor dianggap bacaan | `a method registered without its descriptor is rejected, never treated as a read` |
| `internal/api` mengimpor `internal/schema` | `TestModuleImportsFollowTheAllowList` |

  Setiap komit buangan patuh pada gerbang lain, sehingga yang merah hanya sasarannya. Catat tautan run CI merah di laporan.
- [ ] **Langkah 4: PR siap ditinjau.** Deskripsi PR memuat tabel dependensi dan alat (versi, tanggal cek, tautan, alasan; buf dengan SHA-256), modul tak langsung baru menurut ADR-023 butir 6 (yang 0.x atau pseudo-version disebut namanya), tautan run merah, menit CI job `contract`, keluaran langkah 4 Tugas 5 (ES §9 butir 5), dan daftar berkas yang bentrok dengan 1.9a beserta cara penyelesaiannya.

### Tugas 7: Sesudah merge — penjaga kontrak pertama dihapus, `buf breaking` dibuktikan merah

- [ ] **Langkah 1:** cabang `fs1/1.3-buf-breaking-guard` dari `main` yang sudah memuat kontrak: hapus penjaga kontrak pertama dari `proto-breaking` (bagian "ref tanpa `buf.yaml`"; pemeriksaan ref yang tidak ada tetap). PR → CI hijau: `proto-breaking` kini membandingkan sungguhan.
- [ ] **Langkah 2:** cabang buangan dari cabang itu: nomor medan `client_version` diubah dari 2 ke 5 → `proto-breaking` MERAH (`Field "2" with name "client_version" on message "RequestContext" was deleted` atau bunyi setara buf). Tautan run dicatat; cabang buangan dihapus.
- [ ] **Langkah 3:** `code-reviewer` meninjau PR penjaga; lead merge.

## Selesai bila

- CI hijau di PR; setiap gerbang di tabel Tugas 6 terbukti merah sekali (tautan tercatat); `buf breaking` terbukti merah sesudah merge (Tugas 7).
- Q1, Q2, Q3, Q4, dan Q8 dijawab lead sebelum merge; pertanyaan lain sudah sampai ke lead.
- `code-reviewer` dua lintasan tanpa temuan terbuka; lead merge ke `main`.

## Pertanyaan untuk lead (diulang di laporan)

- **Q1 (memblokir Tugas 4):** bentuk `client_version` tidak ditetapkan dokumen mana pun (09.2 §5.1 hanya `text`, "dibandingkan dengan versi minimum"). Rekomendasi: `MAJOR.MINOR.PATCH` numerik tanpa pra-rilis, sama untuk web (SvelteKit `version.name` diisi 1.9b), iOS, dan Android.
- **Q2:** sumber versi minimum: "konstanta versi minimum per platform" (ADR-028 catatan web butir 5) atau "disimpan di konfigurasi server" (ADR-030 butir 9). Rekomendasi: konstanta, karena bundel tersemat di biner sehingga rilis perbaikan celah selalu biner baru; latihan versi minimum di staging (14 §7.7) memakai biner staging dengan konstanta dinaikkan.
- **Q3 (memblokir merge):** baris 09.2 §1 untuk `internal/api` (lapisan API Connect: interseptor batas, mux, `/healthz`; mengimpor `internal/gen/casa/v1`, kelak modul domain) dan `internal/gen/casa/v1` (kode hasil generate; tidak mengimpor paket `internal/` lain), oleh `architect`.
- **Q4:** "perintah tulis" = metode bertanda `(casa.v1.command)`; RPC `identity` di luar tabel tidak ditolak versi minimum.
- **Q5:** ADR-013 butir 4 menyebut "kode galat ... per medan" sebagai bagian enum yang sama, tetapi tidak ada dokumen yang menamai nilainya. Rencana ini mengirim galat validasi sebagai `CodeInvalidArgument` dengan detail `buf.validate.Violations` (jalur medan dan `rule_id`, ter-generate untuk semua bahasa) tanpa nilai enum. `architect` menamai nilainya atau menerima bentuk ini.
- **Q6:** 09.2 §10 menulis `upgrade_required`; ADR-028 butir 5 dan ADR-030 butir 9 menulis `CLIENT_UPGRADE_REQUIRED`. Rencana mengikuti ADR; 09.2 §10 diselaraskan `architect`.
- **Q7:** nama kode di dokumen bercampur huruf kecil (`session_expired`) dan besar (`WRITE_FROZEN`); di kontrak semuanya menjadi `ERROR_CODE_<NAMA BESAR>` (aturan `buf lint` STANDARD; di TS menjadi `ErrorCode.SESSION_EXPIRED`). Usul: satu kalimat di 09.1 §4 tentang pemetaan ini.
- **Q8:** versi modul Go BSR `buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go` berbentuk `v1.36.12-<waktu>-<komit>.2`, yang secara sintaks semver adalah pra-rilis; ES §8 melarang pra-rilis. Perlu catatan di ADR-023 bahwa bentuk ini versi SDK hasil generate BSR, bukan pra-rilis.

## Putusan lead atas pertanyaan rencana (13 November 2026)

- **Q1:** `client_version` berbentuk angka `MAJOR.MINOR.PATCH` (tanpa awalan, tanpa pra-rilis); nilai lain ditolak sebagai galat validasi.
- **Q2:** versi minimum mengikuti ADR-030 butir 9 (dokumen khusus web, lebih baru dari catatan web ADR-028): **konfigurasi server** (flag `casa serve`), dengan nilai bawaan konstanta di kode. `architect` menyelaraskan catatan web ADR-028.
- **Q3:** baris `internal/api` dan `internal/gen/casa/v1` di 09.2 §1 ditulis `architect` sebelum merge.
- **Q4:** "perintah tulis" = method yang membawa opsi `CommandOptions`. RPC `identity` di luar tabel perintah (09.3 §1). `architect` menyelaraskan 09.2 §8.
- **Q5:** detail galat per medan memakai `buf.validate.Violations`; `architect` mencatatnya di ADR-013 butir 4.
- **Q7:** enum `ERROR_CODE_*` sesuai buf lint; `architect` menambah satu kalimat pemetaan nama dokumen ke enum.
- **Q8:** versi modul Go hasil BSR untuk protovalidate (bentuk `v…-<waktu>-<hash>.N`) adalah build beralamat commit, bukan pra-rilis yang dimaksud ES §8; diterima lewat catatan `architect` di ADR-023.
- **Q9:** batas ukuran permintaan 1 MiB diterima sebagai **[Usulan]**; `architect` mencatatnya di 09.4.
- **Q10:** `RequestContext`, `ClientKind`, `CommandOptions`, `ErrorCode`, `ErrorDetail`, `internal/api` adalah pengenal kode, bukan istilah bersama; tidak masuk glosarium.
- **Q11:** tanpa RPC versi atau kesehatan produk; layanan uji di berkas uji, seperti rencana.
- **Q12:** ukuran `validate_pb.ts` diukur 1.9b terhadap anggaran app shell.
- **Q13:** pembangkit vektor emas masuk 1.7 (rangka dan alat uji).
- **09.2 §10** `upgrade_required` diganti `CLIENT_UPGRADE_REQUIRED` oleh `architect`.
