# FS-1 1.7d — potongan kontrak untuk klien web: rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** satu PR server kecil yang memberi klien web potongan kontrak yang masih kurang, semuanya sudah diputus:
- `IdentityService.GetSession`, `HouseholdService.GetOwnAcceptance`, `GetOperatorNoticeResponse.earlier_version_acknowledged`, dan `casa setup-token -origin` (putusan `architect` Q3);
- `without_session` pada `CompleteSignUpRequest` dan `CompleteRecoveryRequest` (putusan `security-privacy` Q2, opsi B);
- cookie sesi yang dicabut atau kedaluwarsa dikirim kedaluwarsa, sehingga putaran LY-94 "Masuk" berhenti (empat syarat `security-privacy`);
- email yang diketik ulang di `CompleteSignUp` untuk setiap klien (10.4 catatan "halaman verifikasi").

**Arsitektur:**
- Tidak ada modul, tabel, migrasi, atau dependensi baru. Perubahan tinggal di `proto/casa/v1/identity.proto`, `household.proto`, `error.proto` (hanya bila Q1 memilih kode galat), modul `identity`, `read`, `household`, lapisan `api`, dan `cmd/casa`.
- `GetSession` ada di `identity` (ADR-030 butir 10, catatan 14 November 2026), bukan di `ListMembershipsResponse`, sehingga batas modul 09.2 §1 utuh.
- `GetOwnAcceptance` membaca lewat `read` (SEC-06; rencana 1.7 keputusan 6, putusan P5), tidak lewat `household.FindAcceptance`, karena fungsi itu menulis gugur, padahal bacaan tidak pernah mencatat kedaluwarsa (09.3 §1, catatan 14 November 2026).
- Cookie kedaluwarsa dipasang di satu tempat, `boundary.session` jalur cookie, oleh satu fungsi yang juga dipakai `SignOut`.

**Stack:** Go (toolchain di `server/go.mod`), connect-go v2, protovalidate (CEL), sqlc, buf. **Tanpa dependensi baru**, jadi catatan ES §8 tidak diperlukan.

**Spesifikasi:**
- `_progress.md` "Putusan atas rencana jalur web (14 November 2026)" dan "FS-1 1.7 PR C dan web PR B selesai (14 November 2026)".
- ADR-030 butir 5, butir 7 (catatan cache), butir 10 (catatan `GetSession`), butir 11 (catatan bentuk kontrak `without_session`).
- 09.3 §1, catatan 14 November 2026 (`earlier_version_acknowledged`, `GetOwnAcceptance`).
- 09.4 §2, baris "Pendaftaran tertutup", catatan 14 November 2026 (`-origin`).
- 10.5 SEC-161, SEC-171 (b) (2) catatan 14 November 2026 (Q2-1 sampai Q2-6), SEC-172 (d).
- 10.4 SEC-158, catatan 14 November 2026 "halaman verifikasi" butir 1 sampai 7, termasuk 2a (di `origin/main`).
- 10.1 §5.1 butir 5 (normalisasi email), §8.1 (pembatas pemakaian token per alamat, ambang 20), SEC-12, SEC-30.
- Empat syarat `security-privacy` atas pengosongan cookie, disampaikan lead di brief 1.7d. Belum tercatat di 10.5; lihat DR-1.
- `ENGINEERING-STANDARDS.md` §1 sampai §9.

**Pelaksana:** `backend-dev`. **Peninjau:** `code-reviewer` (dua lintasan). **Cabang:** `fs1/1.7d-kontrak-web`, satu PR.

**Langkah 0 (sebelum Task 1).** Cabang ini sekarang berdiri di atas `b6324ba` + `b802960` (catatan progres), **bukan** di atas `origin/main` `26d75ce` (merge PR C). Jadi uji isolasi per kueri dan uji properti PR C belum ada di cabang ini, padahal Task 2 dan Task 4 menyentuhnya. Jalankan `git fetch origin && git rebase origin/main`, lalu periksa bahwa `server/internal/isolation/isolation_test.go` ada.

Urutan task: 1 → 2 → 3 → 4 → 5 → 6. Task 2 menunggu putusan Q1; Task 3 sampai 5 boleh dikerjakan lebih dulu bila Q1 belum diputus.

## Keputusan yang diambil rencana ini

Bila satu butir menulis ulang bunyi dokumen, bunyi dokumen yang menang.

1. **`without_session`** (ADR-030 butir 11, catatan bentuk kontrak; 10.5 SEC-171 (b) (2) Q2-1 sampai Q2-5):
   - **Medan:** `bool without_session = 5` di kedua permintaan.
   - **Aturan CEL `device_mode.web`:** diganti persis menjadi `client_kind == 3 ? (without_session ? device_mode == 0 : device_mode != 0) : (device_mode == 0 && !without_session)`.
     - `id` tetap `device_mode.web`, karena klien memetakan galat menurut id aturan.
     - Kalimat `message` diperluas: "web clients choose a device mode or finish without a session; native clients send neither".
   - **Tanda dibaca seperti `device_mode`:** medan `identity.Client.WithoutSession bool` diisi `clientOf` lewat antarmuka `GetWithoutSession() bool`. Tanda ini sifat permintaan, sama dengan `DeviceMode`, dan dua medan itu diputuskan di tempat yang sama.
   - **`completeSignUp`:** akun, kredensial, dan pemakaian tautan tetap. Bila `client.WithoutSession`, `openSession` tidak dipanggil, jadi tidak ada baris `session`, `sign_in`, maupun token penyegar.
   - **`completeRecovery`:** `revokeSessions(…, keep = pgtype.UUID{}, "recovery", …)` sudah mencabut **semua** sesi. Bila `WithoutSession`, `openSession` dilewati. Baris `security_event` `recovery` tetap ditulis, dengan `SessionID` kosong dan `Detail = summarizeUserAgent(client.UserAgent)` (Q2-3; ADR-030 butir 11 tambahan pertama).
   - **`openedSession` di `api`:** `client.WithoutSession` → `(nil, nil)`, tanpa `Set-Cookie` apa pun (Q2-4). Pemeriksaan ini satu tempat untuk kedua handler.
   - **Batas laju dan galat tidak berubah:** jalur `tokenUse.attempt` dan `outcomeOf` sama, dengan atau tanpa tanda (Q2-5). `TOKEN_REJECTED` tetap byte-identik (SEC-30).
2. **Email yang diketik ulang** (10.4 catatan "halaman verifikasi" butir 2 dan 7):
   - **Medan:** usulan `string email = 6` di `CompleteSignUpRequest` **[Q1]**, dengan `(buf.validate.field).string = {min_len: 1, max_len: 254}`.
     - Sengaja **tanpa** `email: true`. Butir 7 meminta email yang sama "dengan huruf besar dan spasi di ujung" tetap melahirkan akun, sedangkan aturan `email` menolak spasi di batas.
     - Medan kosong ditolak `InvalidArgument` (butir 7).
     - `CompleteRecoveryRequest` tidak mendapat medan ini.
   - **Normalisasi satu implementasi** (ES §2): `normalizeEmail(email string) string` (`strings.ToLower(strings.TrimSpace(email))`) di `identity`.
     - Menggantikan dua salinan yang ada di `signin.go:40` dan `recovery.go:29` (10.1 §5.1 butir 5: "sama seperti pencarian `SignIn`").
     - Email baris pendaftaran sudah lolos aturan `email` kontrak saat `StartSignUp` atau `casa setup-token`, jadi isinya ASCII. Karena itu `strings.ToLower` di Go sama dengan perbandingan `citext`.
   - **Urutan di `completeSignUp`:** tautan dibaca dan diperiksa, lalu `invitationOpen` untuk sumber undangan, lalu **perbandingan email**, lalu `s.hasher.hash`.
     - Tidak cocok: `ErrEmailMismatch`, tanpa tulisan apa pun, dan tautan tetap hidup.
     - `outcomeOf` menghitung `ErrEmailMismatch` sebagai `Failed`, seperti `ErrTokenRejected`.
   - **Jawaban di batas:** menunggu Q1. Opsi A memakai pola `password.common`, opsi B memakai `ErrorCode` baru.
   - **Berlaku untuk semua `client_kind`, dengan atau tanpa `without_session`.**
   - **Kalimat email verifikasi** ("Tautan ini hanya untukmu. Jangan teruskan ke orang lain.") **tidak** di PR ini. Pengirim di server uji (`cmd/casa/mail.go`) hanya menulis `To`, `Kind`, `Link`, tanpa isi; templat isi datang bersama pengirim SES di langkah 7 (09.3 §9 catatan 14 November 2026).
3. **Cookie kedaluwarsa untuk sesi yang berakhir** (empat syarat `security-privacy`; ADR-030 butir 5; SEC-161; SEC-172 (d)):
   - **Fungsi baru `expireSessionCookie(ctx) error` di `api/cookie.go`:** isinya `setSessionCookie(ctx, identity.Cookie{MaxAge: -time.Second})`, dengan atribut identik `__Host-casa_session=`, `Path=/`, `Secure`, `HttpOnly`, `SameSite=Strict`, tanpa `Domain`, dan `Max-Age=0` (syarat 1).
     - `clearSessionCookie` memanggilnya, lalu menambah `Clear-Site-Data`.
     - Jadi atribut cookie kedaluwarsa punya satu implementasi.
   - **Hanya di `boundary.session`, cabang `withCookie`**, saat `sessions.Cookie` mengembalikan `identity.ErrSessionExpired` atau `identity.ErrSessionRevoked`. Termasuk jalur `reused` (pemakaian ulang terdeteksi, SEC-161 (c) aturan (2)), karena token itu cocok dengan baris yang dicabut.
     - `ErrUnauthenticated` (tak dikenal, rusak, token native) tetap jawaban seragam tanpa `Set-Cookie` (syarat 2; SEC-161 (c) aturan (5)).
     - Jalur bearer tidak pernah mendapat cookie.
     - Prosedur publik tidak mengautentikasi cookie, sehingga tidak pernah mengosongkannya; ini juga memenuhi Q2-4.
   - **Data lokal tidak disentuh:** tanpa `Clear-Site-Data` (syarat 2; SEC-172 (d)).
   - **`Cache-Control: no-store`** (syarat 3) sudah dipasang `WithSecurityHeaders` pada setiap jawaban API (ADR-030 butir 7, catatan cache). Uji menegaskannya pada jawaban ini, dan nilainya tidak dipasang dua kali.
   - `// ponytail:` sesi bersama yang berakhir tepat di antara autentikasi dan sentuhan `UserAction` (`checkClient`) dijawab `SESSION_EXPIRED` tanpa cookie kedaluwarsa. Permintaan berikutnya dari browser itu mengosongkannya. Ini balapan sempit yang tidak dapat diuji deterministik tanpa kait.
4. **`GetSession`** (ADR-030 butir 10, catatan 14 November 2026):
   - Hanya handler, tanpa kueri.
   - `account_id = uuidString(session.AccountID)`. `device_mode` dari `session.DeviceMode`: `"personal"` → `DEVICE_MODE_PERSONAL`, `"shared"` → `DEVICE_MODE_SHARED`. Sesi native selalu `"personal"` (`Authenticate`, `openSession`).
   - Bukan perintah (tanpa opsi `casa.v1.command`), jadi tidak ditahan dan tanpa kunci idempotensi. Interseptor tetap melakukan sentuhan dan rotasi biasa.
5. **`GetOwnAcceptance`** (09.3 §1 catatan 14 November 2026):
   - `read.OwnAcceptance(ctx, account) (OwnAcceptance, bool, error)` di `Scope{AccountID}`, dengan kueri baru `readdb.OwnAcceptanceView`: `invitee_display_name, expires_at`, `state = 'accepted'`, dilayani `invitation_invitee`.
   - Hidup bila `s.now().Before(expires_at)`, dengan `Now` yang sama seperti `liveInvitation`.
   - Tidak ada atau lewat masa: jawaban kosong, bukan `NotFound`, dan tidak menulis.
   - Tanpa `household_id`, nama rumah tangga, maupun `invitation_id` (SEC-32).
6. **`earlier_version_acknowledged`** (09.3 §1 catatan 14 November 2026; medan 7):
   - Kueri `LastNoticeIssue` mendapat kolom `earlier_acknowledged` (`EXISTS` atas kunci utama `(account_id, issue_seq)`).
   - `household.Notice` mendapat `EarlierAcknowledged bool`, diisi `CurrentNotice`:
     - identitas sama: `last.EarlierAcknowledged`;
     - terbitan baru: `last.Acknowledged || last.EarlierAcknowledged`;
     - belum ada baris: `false`.
   - Tanpa perjalanan tambahan ke basis data, dan perilaku tahan tidak berubah.
7. **`casa setup-token -origin`** (09.4 §2 catatan 14 November 2026):
   - Flag pilihan, diperiksa `checkOriginFlag` sebelum pool dibuka, sehingga origin yang salah keluar 2 tanpa token terbit.
   - Bila diisi, baris keempat `link: <origin>/signup#<token>` ditambahkan. Ketiga baris yang ada tidak berubah, dan catatan SEC-37 tetap.
   - Jalur `/signup#` adalah rute web `web/src/routes/signup/+page.svelte` dan pola SEC-175 (c).
   - Tidak ada log; perintah ini hanya menulis ke stdout operator.
8. **Butir parkir** (`AcknowledgeOperatorNotice` tanpa `idempotencyKey` dijawab "request_context required"):
   - Uji probe yang ada (`server_test.go`, "a write command without an idempotency key never reaches the handler") sudah menegaskan jalur `request_context.idempotency_key` dengan aturan `required`.
   - Task 6 langkah 1 mereproduksinya di handler yang sebenarnya.
     - Bila jawabannya sama dengan probe: tidak ada cacat server. Butir ditutup dengan laporan, dan lead diminta menunjukkan sumber pengamatannya.
     - Bila berbeda dan perbaikannya satu baris di `missingIdempotencyKey`: diperbaiki dengan uji itu sebagai uji merah.
     - Selain itu: ditunda dengan alasan di PR.
9. **Tanpa migrasi:**
   - `security_event.kind` sudah memuat `recovery` dan `session_id` boleh NULL (`00002_identity.sql`, `00007`).
   - Perubahan `LastNoticeIssue` hanya SQL kueri.
   - QG-05 (`server-test-old-binary`) tidak mendapat skema baru.
10. **Urutan masuk baru pemilik produk** (10.5 §3 catatan 14 November 2026, worktree web `1faee51`) **tidak memerlukan perubahan server.**
    - Pertanyaan mode tetap dijawab sebelum `SignIn` dan, saat daftar, sebelum kata sandi.
    - `SignInRequest.device_mode` dan `CompleteSignUpRequest.device_mode` tetap wajib bagi web.
    - Bawaan pribadi di aplikasi terpasang dan tautan "Masuk sebagai perangkat bersama" hanya urusan klien.

## Batasan global

- **Warisan.** Batasan 1.1a sampai 1.7 berlaku (ES §1 sampai §9; alat dipatok). Tidak ada dependensi baru.
- **SQL.**
  - Semua lewat sqlc dan berparameter.
  - Dua kueri berubah atau baru (`readdb.OwnAcceptanceView`, `householddb.LastNoticeIssue`). Keduanya punya indeks yang melayani dan `EXPLAIN` di deskripsi PR (ES §3).
  - Uji isolasi PR C mendapat entri `"readdb.OwnAcceptanceView.account_id": pair(inviteeX, inviteeY)` dan baris `ownCallFails` `{"readdb.OwnAcceptanceView", lookupScopeOnly}`, sama dengan `householddb.OwnAcceptance`.
- **Log dan galat.**
  - Tidak ada baris log baru.
  - Tidak ada email, token, nilai cookie, atau `ua_summary` di galat klien atau log (SEC-44; ES §4).
  - `Detail` jejak keamanan hanya `ua_summary`.
- **Uji.**
  - PostgreSQL 18 nyata; gagal keras tanpa `CASA_TEST_ADMIN_URL`; tanpa `t.Parallel` di paket basis data; jam lewat `testClock`.
  - Iterasi: `cd server && CASA_TEST_ADMIN_URL='postgres://postgres@127.0.0.1:5432/postgres?sslmode=disable' go test -race -p 1 -run '<Uji>' ./internal/<paket>/`. `make check` hanya di Task 6.
- **Harness.** Pemanggil `CompleteSignUp` di uji yang ada mendapat email **hanya pada pemasangan**: `fixtureSignUp` (`Email: start.Email`), `signedUp` dan `startedSignUp` di `identity`, `household_flow_test.go`, `property_test.go`, dan `websession_test.go`. Asersi dan nama uji lama tidak berubah.
- **Penamaan (QG-06).** ID SEC di nama uji hanya bila uji membuktikan aturan itu. Tidak ada AC Tier 1 yang dibuktikan utuh di PR ini, jadi tidak ada ID AC di nama uji. Komentar mengutip ID atau bagian dokumen.
- **Kontrak.**
  - Hanya tambahan: dua RPC, empat pesan, tiga medan, dan nol atau satu nilai enum menurut Q1.
  - `buf lint`, `buf breaking` (FILE), dan `make proto-generated` lulus. Kode Go dan TypeScript (`web/src/lib/gen`) dibangkitkan ulang dan di-commit (14 §4.6).
- **Migrasi.** Tidak ada.

## Fokus tinjauan

1. **Cookie basi di browser yang sama menghapus cookie baru.**
   - Permintaan lama yang masih dalam perjalanan dengan nilai sesi yang sudah dicabut dapat tiba sesudah `SignIn` baru di browser yang sama. `Max-Age=0`-nya lalu menghapus cookie baru, karena namanya sama.
   - Server tidak dapat mengetahui isi wadah cookie. Penutupnya ada di klien: SEC-172 (a) menghentikan semua permintaan dan menghapus sebelum layar lain, jadi `SignIn` datang sesudahnya.
   - Uji server membuktikan batas yang dapat dibuktikannya: `a stale cookie of one browser never ends the session of another` (Task 3). Sisa risikonya dinyatakan di Q3.
2. **Token tak dikenal mendapat cookie kedaluwarsa.** Itu akan membedakan "pernah ada" dari "tidak pernah ada". Diuji di Task 3 (`an unknown, malformed or native token keeps the uniform answer and sets no cookie`).
3. **Email tidak cocok sesudah hash atau sebelum pemeriksaan tautan.** Yang pertama membuka jalan DoS Argon2; yang kedua membuka galat tersendiri bagi tautan mati (SEC-30). Diuji di Task 2 (`an email mismatch runs no hash and keeps the link`; `a used or expired link gets the same TOKEN_REJECTED with the right or a wrong email`).
4. **`without_session` yang tetap membuat sesi atau memasang cookie,** atau pemulihan tanpa sesi yang melewatkan pencabutan atau jejak. Diuji di Task 1 (dua subuji SEC-171 (b) (2) dan `a without-session completion leaves a stale cookie untouched`).
5. **`GetOwnAcceptance` yang menulis gugur atau membuka rumah tangga.** Diuji di Task 4 (`GetOwnAcceptance never writes and an acceptance past its expiry reads empty`; `GetOwnAcceptance carries no household data`).

---

### Task 1: `without_session` pada penyelesaian pendaftaran dan pemulihan

**Berkas:**
- **Ubah:**
  - `proto/casa/v1/identity.proto`: medan 5 di kedua pesan, aturan CEL, komentar RPC `CompleteSignUp` dan `CompleteRecovery` ("…opens a session unless without_session")
  - hasil generate Go dan TypeScript
  - `server/internal/identity/identity.go` (`Client.WithoutSession`)
  - `server/internal/identity/signup.go`, `recovery.go`
  - `server/internal/api/identity.go` (`clientOf`, `openedSession`)
- **Uji:** `server/internal/api/websession_test.go` (`TestWithoutSession`)

**Antarmuka:**
- **Consumes:** `identity.Client`, `openSession`, `revokeSessions`, `summarizeUserAgent`, `identitydb.InsertSecurityEvent` (semuanya sudah ada).
- **Produces:**
  - `identity.Client.WithoutSession bool`, dengan komentar yang mengutip SEC-171 (b) (2) catatan 14 November 2026 Q2 dan ADR-030 butir 11;
  - `GetWithoutSession() bool` di kedua pesan hasil generate.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestWithoutSession(t *testing.T) {
	t.Run("SEC-171 (b) (2) a web sign-up completed without a session sets no cookie, returns no tokens and opens no session", …)
	// web, WithoutSession true, DeviceMode 0: OK; len(Set-Cookie) == 0; GetTokens() == nil;
	// count(session) == 0 and count(security_event kind 'sign_in') == 0 for the account;
	// the account and its credential exist; the same link again → TOKEN_REJECTED; SignIn with the password → OK
	t.Run("SEC-171 (b) (2) a web recovery without a session revokes every session and records recovery without a session", …)
	// account with one web cookie session and one native session; CompleteRecovery web, WithoutSession:
	// len(Set-Cookie) == 0, GetTokens() == nil; count(session WHERE revoked_at IS NULL) == 0;
	// exactly one security_event kind 'recovery' with session_id IS NULL and detail = "Safari · iOS" (webCall(safariIPhone));
	// the old cookie → SESSION_REVOKED; SignIn with the new password → OK
	t.Run("the flag from a native client, or with a device mode, is an invalid argument", …)
	// for both requests: {ios, true, 0}, {web, true, PERSONAL}, {web, false, 0} → InvalidArgument, rule "device_mode.web";
	// {ios, false, 0} and {web, false, PERSONAL} still OK
	t.Run("a rejected link gets the same TOKEN_REJECTED with and without the flag", …)
	// SEC-30: code, message and details equal for an unknown token, both requests
	t.Run("the token-use limiter counts the same with and without the flag", …)
	// one fixture peer: 10 unknown tokens without and 10 with the flag → the 21st (either form) is RATE_LIMITED (10.1 §8.1, threshold 20)
	t.Run("a without-session completion leaves a stale cookie untouched", …)
	// Q2-4: the request carries the cookie of a revoked session (sameOrigin(old)); the answer has no Set-Cookie at all
}
```

- [ ] **Langkah 2:** `make proto-generate`, lalu `go test -race -p 1 -run TestWithoutSession ./internal/api/` → GAGAL (medan ada, perilaku belum).
- [ ] **Langkah 3:** tulis keputusan 1: `clientOf`, `openedSession`, cabang `WithoutSession` di `completeSignUp` dan `completeRecovery`.
- [ ] **Langkah 4:** jalankan langkah 2, lalu `go test -race -p 1 ./internal/api/ ./internal/identity/` dan `make check-contract` → LULUS.
- [ ] **Langkah 5: commit** `Finish sign-up and recovery without a session for a Safari tab (SEC-171 (b) (2) note Q2; ADR-030 item 11)`.

### Task 2: Email yang diketik ulang di `CompleteSignUp`

Dimulai sesudah Q1 diputus. Di bawah ini ditulis dengan usulan: medan `email` = 6, dan opsi A (pelanggaran medan `email.mismatch`). Bila Q1 memilih lain, hanya nama dan isi `wantEmailMismatch` yang berubah.

**Berkas:**
- **Ubah:**
  - `proto/casa/v1/identity.proto` (medan 6)
  - `error.proto`, hanya bila Q1 = opsi B
  - hasil generate
  - `server/internal/identity/signup.go` (`ErrEmailMismatch`, `normalizeEmail`, urutan, `outcomeOf`)
  - `signin.go`, `recovery.go` (memakai `normalizeEmail`)
  - `server/internal/api/identity.go` (handler meneruskan `req.GetEmail()`; `identityError`)
  - pemasangan harness (Batasan global, "Harness")
- **Uji:** `server/internal/api/identity_test.go` (`TestSignUpEmailRetyped`), `server/internal/identity/signup_test.go` (baru)

**Antarmuka:**
- **Produces:**
  - `var ErrEmailMismatch = errors.New("identity: the typed email is not the link's")`;
  - `func normalizeEmail(email string) string`;
  - `func (s *Service) CompleteSignUp(ctx context.Context, client Client, verificationToken, email, password string, invitationOpen func(context.Context, []byte) error) (Tokens, error)`.
- **Jawaban** (`identityError`):
  - opsi A: `invalidFields(fieldViolation(<email field of CompleteSignUpRequest>, "email.mismatch", "email is not the one the link was sent to"))`;
  - opsi B: `errorWithDetail(connect.CodeFailedPrecondition, "email mismatch", &casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_EMAIL_MISMATCH})`.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
// identity_test.go
func TestSignUpEmailRetyped(t *testing.T) {
	// 10.4 SEC-158 note of 14 November 2026 "halaman verifikasi" items 2 and 7.
	t.Run("a valid link with another email is refused on its own answer, writes nothing and keeps the link", …)
	// both sources (setup token; invitation with h.invitationOpen returning nil): wantEmailMismatch;
	// count(account) and count(session) unchanged; signup_verification.used_at IS NULL;
	// then the same link with "  " + strings.ToUpper(email) + " " → OK, the account is born
	t.Run("a mismatch counts one failure on the address limiter", …)
	// one fixture peer, one valid link: 20 mismatches, then the right email → RATE_LIMITED; another peer → OK
	t.Run("a used or expired link gets the same TOKEN_REJECTED with the right or a wrong email", …)
	// used link, expired link (clock past expires_at) × {right, wrong} email: four errors byte-equal to tokenRejected()
	t.Run("an empty email is an invalid argument", …) // Email "" → InvalidArgument on field "email", rule "string.min_len"
	t.Run("the rule holds for a native client and without a session", …)
	// ios with a wrong email and web WithoutSession with a wrong email → wantEmailMismatch
}

// identity/signup_test.go
func TestCompleteSignUpEmail(t *testing.T) {
	t.Run("an email mismatch runs no hash and keeps the link", …)
	// before := s.hasher.count.Load(); CompleteSignUp with "other@example.test" → ErrEmailMismatch;
	// s.hasher.count.Load() == before; then the right email → OK
}
```

- [ ] **Langkah 2:** `make proto-generate`, lalu `go test -race -p 1 -run 'TestSignUpEmailRetyped|TestCompleteSignUpEmail' ./internal/api/ ./internal/identity/` → GAGAL.
- [ ] **Langkah 3:** tulis keputusan 2: `normalizeEmail` dan dua pemakainya yang lama, perbandingan sesudah `invitationOpen` dan sebelum `s.hasher.hash`, `outcomeOf`, `identityError`, handler, dan pemasangan harness.
- [ ] **Langkah 4:** jalankan langkah 2, lalu `go test -race -p 1 ./internal/api/ ./internal/identity/ ./internal/isolation/` dan `make check-contract` → LULUS. Uji properti PR C juga tetap hijau.
- [ ] **Langkah 5: commit** `Require the typed email when a verification link is opened, before the hash (10.4 SEC-158 note "halaman verifikasi" items 2, 7; Q1 ruling)`.

### Task 3: Cookie kedaluwarsa bagi sesi yang dicabut atau kedaluwarsa

**Berkas:**
- **Ubah:** `server/internal/api/cookie.go` (`expireSessionCookie`; `clearSessionCookie` memakainya), `server/internal/api/server.go` (`boundary.session`)
- **Uji:** `server/internal/api/websession_test.go`, di `TestCookieSession`. Bagian cookie dari `wantSiteCleared` dipisah menjadi `wantCookieExpired(t, info)`, dan `wantSiteCleared` memanggilnya, sehingga uji `SignOut` yang lama tidak berubah isinya.

**Antarmuka:**
- **Produces:** `func expireSessionCookie(ctx context.Context) error`, yang gagal tertutup seperti `setSessionCookie`.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
	t.Run("SEC-161 a revoked or expired session's cookie is expired with the same attributes, without touching site data", …)
	// three cases: signed out (revoked), 30 idle days (expired), a previous value reused after its window (revoked, reuse_detected):
	// error code as today; wantCookieExpired(t, info); no Clear-Site-Data header;
	// Cache-Control exactly "no-store"
	t.Run("an unknown, malformed or native token keeps the uniform answer and sets no cookie", …)
	// tokenOf(t), "not base64!", a value two rotations old, a native access token in the cookie:
	// Unauthenticated without detail; len(Set-Cookie) == 0; no Clear-Site-Data
	t.Run("a stale cookie of one browser never ends the session of another", …)
	// condition (4): one account, browser A (cookie a) and browser B (cookie b); A signs out;
	// a again → SESSION_REVOKED with the expired cookie; then b → OK, b's row revoked_at IS NULL and not touched by A's answer;
	// the same with A expired by 30 idle days while B is used daily
	t.Run("a native request of an ended session gets no cookie", …)
	// bearer token of a revoked native session → SESSION_REVOKED, len(Set-Cookie) == 0
```

- [ ] **Langkah 2:** `go test -race -p 1 -run 'TestCookieSession' ./internal/api/` → GAGAL (subuji pertama dan ketiga).
- [ ] **Langkah 3:** tulis keputusan 3. Komentar di `boundary.session` mengutip ADR-030 butir 5, SEC-161, SEC-172 (d), dan keempat syarat (sumbernya DR-1). Langkah 3 di komentar `NewServer` menyebut cookie kedaluwarsa.
- [ ] **Langkah 4:** `go test -race -p 1 ./internal/api/` → LULUS, termasuk `TestCookieRotation` dan `TestWebSessionOpening` tanpa suntingan.
- [ ] **Langkah 5: commit** `Expire the cookie of a revoked or expired session, never an unknown one (ADR-030 item 5; SEC-161; SEC-172 (d); LY-94)`.

### Task 4: `GetSession`, `GetOwnAcceptance`, dan `earlier_version_acknowledged`

**Kontrak:**

```proto
// identity.proto, IdentityService
  // Returns the caller's own account and the device mode of the session that
  // made the call (ADR-030 item 10 note of 14 November 2026; SEC-170 (a);
  // SEC-171 (e)); a native session is personal. Outside the queue: never held
  // by NOTICE_UNREAD, and it writes nothing beyond the boundary's touch and
  // rotation (09.3 §1).
  rpc GetSession(GetSessionRequest) returns (GetSessionResponse);
message GetSessionRequest { RequestContext request_context = 1 [(buf.validate.field).required = true]; }
message GetSessionResponse { string account_id = 1; DeviceMode device_mode = 2; }

// household.proto, HouseholdService
  // Returns the caller's own acceptance, found from the caller's account, never
  // from the request (plan 1.6 decision 14): the display name it gave and its
  // expiry, no household data (SEC-32). Empty when there is none or it is past
  // its expiry; it never writes, so an expiry is never recorded here. An online
  // read, never held by NOTICE_UNREAD (09.3 §1 note of 14 November 2026).
  rpc GetOwnAcceptance(GetOwnAcceptanceRequest) returns (GetOwnAcceptanceResponse);
message GetOwnAcceptanceRequest { RequestContext request_context = 1 [(buf.validate.field).required = true]; }
// Empty when the caller has no live acceptance.
message GetOwnAcceptanceResponse { string display_name = 1; google.protobuf.Timestamp expires_at = 2; }
// GetOperatorNoticeResponse
  // True when the account acknowledged an issue before the current one, A→B→A
  // included (09.3 §1 note of 14 November 2026); a read only.
  bool earlier_version_acknowledged = 7;
```

**Berkas:**
- **Ubah:**
  - kedua `.proto` dan hasil generate
  - `server/internal/api/identity.go` (handler `GetSession`; `deviceModeOf(mode string) casav1.DeviceMode`)
  - `server/internal/read/read.go` dan `queries.sql` (`OwnAcceptance`, `OwnAcceptanceView`), hasil sqlc
  - `server/internal/household/notice.go` dan `queries.sql` (`LastNoticeIssue`, `Notice.EarlierAcknowledged`), hasil sqlc
  - `server/internal/api/household.go` (handler `GetOwnAcceptance`; `GetOperatorNotice` mengisi medan 7)
  - `server/internal/isolation/isolation_test.go` (dua entri, Batasan global)
  - `server/internal/api/notice_test.go` (`noticeFreeReads` + `GetOwnAcceptance`)
- **Uji:** `websession_test.go` atau `identity_test.go` (`TestGetSession`), `invitation_test.go` (`TestGetOwnAcceptance`), `notice_test.go` (subuji di `TestOperatorNotice`), `read/read_test.go`

**Antarmuka:**
- **Produces:**
  - `type OwnAcceptance struct { DisplayName string; ExpiresAt time.Time }`;
  - `func (s *Service) OwnAcceptance(ctx context.Context, account pgtype.UUID) (OwnAcceptance, bool, error)` di `read`;
  - `household.Notice.EarlierAcknowledged bool`.

```sql
-- name: OwnAcceptanceView :one
-- The caller's own accepted invitation as GetOwnAcceptance shows it (09.3 §1
-- note of 14 November 2026), in a lookup transaction (Scope{AccountID}; the
-- invitee clause of 09.2 §3): only the display name and the expiry. Served
-- by invitation_invitee, which also makes it one row at most.
SELECT invitee_display_name, expires_at FROM invitation
 WHERE invitee_account_id = @account_id AND state = 'accepted';
```

`LastNoticeIssue` ditambah `EXISTS (SELECT 1 FROM operator_notice_ack e WHERE e.account_id = @account_id AND e.issue_seq < l.issue_seq AND e.acknowledged_at IS NOT NULL) AS earlier_acknowledged`, dengan alias `l` di tabel utama. Kunci utama melayani keduanya.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestGetSession(t *testing.T) {
	t.Run("GetSession answers the caller's account and the session's mode", …)
	// web personal → PERSONAL, web shared → SHARED, native → PERSONAL; account_id equals the account row's id
	t.Run("GetSession without a session answers like any authenticated RPC", …)
	// no cookie and no bearer → Unauthenticated without detail; revoked cookie → SESSION_REVOKED (with Task 3's expired cookie)
	t.Run("GetSession is never held and writes nothing but the boundary's touch", …)
	// unread notice: OK; counts of every identity table except session.last_used_at unchanged
}
func TestGetOwnAcceptance(t *testing.T) {
	t.Run("the invitee reads its own acceptance: display name and expiry", …)
	t.Run("GetOwnAcceptance carries no household data", …)
	// response descriptor has exactly fields display_name and expires_at (SEC-32)
	t.Run("GetOwnAcceptance never writes and an acceptance past its expiry reads empty", …)
	// clock past expires_at: empty response; invitation.state still 'accepted', count(object_event) unchanged
	t.Run("no acceptance reads empty, not NotFound", …) // a fresh account, and the inviter
	t.Run("GetOwnAcceptance is never held", …)
}
	// TestOperatorNotice
	t.Run("earlier_version_acknowledged follows the earlier issues, A, B, A included", …)
	// first sight: false; acknowledge; a new version: acknowledged false, earlier true;
	// operator name A (acknowledged) → B (not acknowledged) → A: earlier true on the second A;
	// an account whose only earlier issue was never acknowledged: false
```

Ditambah: `GetOwnAcceptance` di `noticeFreeReads`. Uji `every RPC that is not a command is a named read or held` merah tanpa baris ini, dan `reading is never held` memanggilnya.

- [ ] **Langkah 2:** `make proto-generate sqlc-generate`, lalu `go test -race -p 1 -run 'TestGetSession|TestGetOwnAcceptance|TestOperatorNotice|TestNoticeHold' ./internal/api/` → GAGAL.
- [ ] **Langkah 3:** tulis keputusan 4 sampai 6.
- [ ] **Langkah 4:** jalankan langkah 2, lalu `go test -race -p 1 ./internal/api/ ./internal/read/ ./internal/household/ ./internal/isolation/` dan `make check-contract sqlc-generated` → LULUS. Lampirkan `EXPLAIN` kedua kueri.
- [ ] **Langkah 5: commit** `Add GetSession, GetOwnAcceptance and earlier_version_acknowledged for the web client (ADR-030 item 10 note; 09.3 §1 note of 14 November 2026)`.

### Task 5: `casa setup-token -origin`

**Berkas:**
- **Ubah:**
  - `server/cmd/casa/setuptoken.go`: flag, `setupTokenUsage` = `"usage: casa setup-token -email address [-origin https://host]\n"`
  - `README.md` akar, baris 82: contoh dengan `-origin`
- **Uji:** `server/cmd/casa/setuptoken_test.go`

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
	t.Run("-origin adds the setup link after the three lines", …)
	// -origin https://casa.example.test: stdout = the three lines + "link: https://casa.example.test/signup#<token>\n",
	// the token in the link equals the one on line 1; stderr empty
	t.Run("an origin casa serve would refuse issues no token", …)
	// "http://casa.example.test", "https://casa.example.test/", "https://Casa.example.test": exit 2, stderr names -origin,
	// stdout empty, count(setup_token) == 0
```

Uji yang ada (`issues one token bound to one email and prints it once`) tetap tanpa suntingan dan membuktikan bahwa tanpa `-origin` tidak ada baris keempat.

- [ ] **Langkah 2:** `go test -race -p 1 -run TestSetupToken ./cmd/casa/` → GAGAL.
- [ ] **Langkah 3:** tulis keputusan 7.
- [ ] **Langkah 4:** jalankan langkah 2 → LULUS.
- [ ] **Langkah 5: commit** `Print the setup link with casa setup-token -origin (09.4 §2 note of 14 November 2026; SEC-158 (b))`.

### Task 6: Butir parkir, verifikasi, jalan lokal, bukti merah, PR

- [ ] **Langkah 1: butir parkir** (keputusan 8). Tulis subuji di `TestNoticeHold` atau `TestOperatorNotice`: `AcknowledgeOperatorNotice without an idempotency key names request_context.idempotency_key`, lewat handler yang sebenarnya, dengan `violationPaths` == `["request_context.idempotency_key"]` dan rule `required`.
  - Lulus seketika: commit sebagai uji regresi, lalu laporkan "tidak tereproduksi" beserta pertanyaan sumber ke lead.
  - Gagal: perbaiki bila satu baris; selain itu tunda, dengan alasan di PR.
- [ ] **Langkah 2:** `make check` → LULUS, termasuk `server-test-old-binary` dan `check-contract`.
- [ ] **Langkah 3: jalan lokal sungguhan** (ES §9 butir 5): `make casa db-up`, `.cache/casa migrate`, `casa serve` dengan kredensial uji, lalu `curl` dengan `Origin`. Keluaran dilampirkan dengan token, cookie, dan email disensor.
  - `casa setup-token -email … -origin https://127.0.0.1:…` mencetak baris `link:`.
  - `StartSignUp`, lalu `CompleteSignUp` web dengan email salah: jawaban Q1. Dengan `without_session` dan email benar: tanpa `Set-Cookie`.
  - `SignIn` web: `GetSession` → `PERSONAL`. `SignOut`, lalu `GetSession` dengan cookie lama: `SESSION_REVOKED` dan `Set-Cookie: __Host-casa_session=; Path=/; Max-Age=0; HttpOnly; Secure; SameSite=Strict`, tanpa `Clear-Site-Data`.
  - `GetOwnAcceptance` tanpa penerimaan: `{}`.
- [ ] **Langkah 4: bukti merah**, satu per satu, di PR draf sekali pakai:

| Pelanggaran buangan | Yang harus merah |
|---|---|
| `openedSession` tanpa cabang `WithoutSession` | `SEC-171 (b) (2) a web sign-up completed without a session …` |
| `completeRecovery` tetap memanggil `openSession` dengan tanda | `SEC-171 (b) (2) a web recovery without a session …` |
| Aturan CEL lama (tanda tidak disebut) | `the flag from a native client, or with a device mode, is an invalid argument` |
| `done(NotTried)` bila tanda diisi | `the token-use limiter counts the same with and without the flag` |
| Perbandingan email sesudah `s.hasher.hash` | `an email mismatch runs no hash and keeps the link` |
| Perbandingan email sebelum pemeriksaan tautan | `a used or expired link gets the same TOKEN_REJECTED …` |
| Perbandingan tanpa `normalizeEmail` | `a valid link with another email is refused …` (bagian huruf besar dan spasi) |
| `outcomeOf` tidak menghitung `ErrEmailMismatch` | `a mismatch counts one failure on the address limiter` |
| Cookie kedaluwarsa juga pada `ErrUnauthenticated` | `an unknown, malformed or native token keeps the uniform answer and sets no cookie` |
| `expireSessionCookie` memanggil `clearSessionCookie` (dengan `Clear-Site-Data`) | `SEC-161 a revoked or expired session's cookie is expired …` |
| `GetOwnAcceptance` lewat `household.FindAcceptance` | `GetOwnAcceptance never writes and an acceptance past its expiry reads empty` |
| `OwnAcceptance` tanpa pemeriksaan `expires_at` | yang sama |
| `EarlierAcknowledged` hanya dari baris terakhir | `earlier_version_acknowledged follows the earlier issues, A, B, A included` |
| `deviceModeOf` selalu `PERSONAL` | `GetSession answers the caller's account and the session's mode` |
| Token diterbitkan sebelum `checkOriginFlag` | `an origin casa serve would refuse issues no token` |

- [ ] **Langkah 5: deskripsi PR** memuat:
  - dua kueri dengan indeks dan `EXPLAIN`;
  - tautan run merah;
  - keluaran jalan lokal;
  - putusan Q1 sampai Q3 yang dipakai;
  - "tanpa dependensi baru, tanpa migrasi";
  - hasil butir parkir;
  - catatan bagi `web-dev` (Risiko 2).
- [ ] **Langkah 6:** `code-reviewer` meninjau dua lintasan (Opus untuk lintasan keamanan: cookie dan email). Temuan diperbaiki, lalu lead merge (merge commit).

## Yang sengaja tidak di PR ini

- **Klien (`web-dev`, web PR E):**
  - kolom email di halaman verifikasi, beserta kalimatnya (Q2);
  - aturan 2a (tidak memanggil `CompleteSignUp` saat sudah masuk), sebuah kewajiban klien yang tidak dapat diperiksa server (10.4 butir 2a);
  - pengiriman `without_session` di tab Safari iOS dan iPadOS;
  - pemakaian `GetSession`, `GetOwnAcceptance`, dan medan 7;
  - penghentian permintaan yang masih berjalan sebelum `SignIn` sesudah `session_revoked` atau `session_expired` (Fokus tinjauan 1).
- **Langkah 7:** kalimat email verifikasi di templat SES (keputusan 2).
- **Sentuhan `UserAction`** yang menemukan sesi berakhir tidak mengosongkan cookie (keputusan 3, `ponytail:`).

## Risiko

1. **Pendaftaran web di server uji dan pratinjau LAN gagal** sejak PR ini di-merge sampai web PR E mengirim `email`: `CompleteSignUp` tanpa `email` ditolak `InvalidArgument`. Pemilik produk yang mencoba mendaftar di perangkat akan melihat galat. Usulan: lead menjadwalkan web PR E tepat sesudahnya, atau `web-dev` menambah medan itu lebih dulu dalam perubahan kecil. Uji web di CI tidak terkena, karena Playwright memakai `cmd/webpreview` tanpa API.
2. **Berkas hasil generate bentrok** dengan cabang web `fs1/1.10b-tampilan`: `web/src/lib/gen/casa/v1/identity_pb.ts`, `household_pb.ts`, dan `error_pb.ts` bila Q1 = B. Penyelesaiannya: bangkitkan ulang sesudah rebase, jangan menggabung dengan tangan.
3. **Cabang belum di atas `origin/main`** (Langkah 0). Tanpa rebase, Task 2 dan Task 4 tidak dapat memperbarui uji isolasi dan uji properti.
4. **Sumber empat syarat cookie belum tercatat** di dokumen (DR-1). Komentar kode harus dapat mengutip bagian 10.5 sebelum merge.

## Berkas yang mungkin bentrok

- `web/src/lib/gen/` (dibangkitkan ulang).
- `server/internal/api/identity.go`, `household.go`, `server.go`, `cookie.go`.
- `server/internal/identity/signup.go`, `recovery.go`, `signin.go`.
- `server/internal/isolation/isolation_test.go`, `server/internal/api/property_test.go` (PR C; hanya pemasangan).

## Pertanyaan rencana

- **Q1 (`architect`): nama medan dan bentuk jawaban email tidak cocok.** 10.4 butir 2: "kode galat tersendiri (namanya ditetapkan `architect` di kontrak)". Usulan:
  - **Medan:** `string email = 6` di `CompleteSignUpRequest`, dengan `string = {min_len: 1, max_len: 254}`, tanpa `email: true` (alasan di keputusan 2).
  - **Jawaban, opsi A (disarankan):** pelanggaran medan `email` dengan `rule_id` `email.mismatch`, dalam bentuk Violations yang sama dengan `password.common` (putusan lead Q5 atas rencana 1.3). Alasannya:
    - kalimat 10.4 butir 6 tampil "di bawah isian", dan klien sudah memetakan (medan, aturan) ke kalimat lewat `fieldErrors`;
    - kontrak tidak mendapat nilai enum baru.
  - **Jawaban, opsi B:** `ERROR_CODE_EMAIL_MISMATCH = 18` dengan `connect.CodeFailedPrecondition` dan `ErrorDetail`, ditambah catatan di ADR-013 (pola catatan 1.6).
  - Keduanya memenuhi "bukan `TOKEN_REJECTED`".
- **Q2 (`ux-designer`): kata di 08 §7.6.** Kalimat dari 10.4 butir 6 belum ada di 08 §7.6 `origin/main`:
  - label "Email" dan petunjuk "Ketik email yang menerima tautan ini.";
  - galat "Email ini bukan email yang menerima tautan ini. Kalau tautan ini dikirim orang lain kepadamu, jangan dipakai. Minta undangan sendiri.";
  - baris email verifikasi "Tautan ini hanya untukmu. Jangan teruskan ke orang lain.".

  Server tidak membawa kata apa pun, jadi Q2 tidak menahan PR ini, tetapi menahan web PR E.
- **Q3 (`security-privacy`, informatif): sisa risiko Fokus tinjauan 1.** `Max-Age=0` dari permintaan lama yang tiba sesudah `SignIn` baru di browser yang sama menghapus cookie baru. Server tidak dapat mencegahnya. Usulan: penutupnya kewajiban klien di SEC-172 (a), yaitu menghentikan permintaan yang berjalan sebelum layar lain, dan dicatat sebagai sisa yang diterima. Tidak menahan PR.

## Permintaan perubahan dokumen

- **DR-1 (`security-privacy`, 10.5 SEC-161):** catatan bertanggal yang mencatat empat syarat pengosongan cookie. Bunyi usulan: "Jawaban `session_revoked` dan `session_expired` atas cookie yang cocok dengan baris sesi yang dicabut atau kedaluwarsa membawa cookie kedaluwarsa dengan atribut identik (`__Host-casa_session`, `Secure`, `HttpOnly`, `SameSite=Strict`, `Path=/`, tanpa `Domain`, nilai kosong, `Max-Age=0`; ADR-030 butir 5). Token tak dikenal atau rusak tetap jawaban seragam tanpa `Set-Cookie`. Data lokal tidak disentuh (tanpa `Clear-Site-Data`; SEC-172 (d)). Jawabannya `Cache-Control: no-store`. Uji: cookie basi satu browser tidak pernah mengakhiri sesi browser lain."
- **DR-2 (`architect`, ADR-013 catatan 1.6):** hanya bila Q1 = opsi B, satu baris untuk `ERROR_CODE_EMAIL_MISMATCH`.
- **DR-3 (`ux-designer`, 08 §7.6):** kalimat Q2.

## Selesai bila

- CI hijau (`contract`, `server`, `web`).
- Setiap baris bukti merah terbukti merah.
- `code-reviewer` tanpa temuan terbuka, dan lead sudah merge.
- Di `main`:
  - web dapat menyelesaikan pendaftaran dan pemulihan tanpa sesi;
  - `CompleteSignUp` menolak email lain sebelum hash;
  - sesi yang berakhir mengosongkan cookie-nya;
  - `GetSession`, `GetOwnAcceptance`, medan 7, dan `-origin` tersedia bagi web PR E.

## Putusan pra-eksekusi

- **Lead (pemindaian pra-eksekusi, 15 November 2026):**
  - R1: pemasangan harness Task 2 juga mencakup panggilan `CompleteSignUp` di `TestWithoutSession` (Task 1 lahir lebih dulu).
  - R2: baris bukti merah "`expireSessionCookie` memanggil `clearSessionCookie`" berputar tanpa ujung (`clearSessionCookie` memanggil `expireSessionCookie`); diganti "`expireSessionCookie` juga memasang `Clear-Site-Data`".
  - R3: urutan 1 → 2 → 3 → 4 → 5 → 6 (Q1 sudah diputus).

- **Q1 (architect): medan `string email = 6` di `CompleteSignUpRequest`, `min_len: 1`, `max_len: 254`, tanpa `email: true` (disetujui seperti usulan). Jawaban tidak cocok: opsi B, `ERROR_CODE_EMAIL_MISMATCH = 18`, `connect.CodeFailedPrecondition`, dengan `ErrorDetail`.** Alasan:
  - 10.4 butir 2 menulis "kode galat tersendiri (namanya ditetapkan architect di kontrak)", dan butir 7 membedakan "kode galat tersendiri" dari "ditolak sebagai argumen tidak sah" untuk medan kosong. Pelanggaran medan adalah `InvalidArgument`, yaitu bentuk masukan yang salah; email di sini bentuknya sah, yang tidak cocok adalah keadaan baris tautan. Itu `FailedPrecondition`, pola `ErrorCode` yang sama dengan 13 sampai 17 di ADR-013 catatan 1.6. `password.common` aturan kebijakan atas masukan itu sendiri, jadi bukan preseden.
  - Bukan `TOKEN_REJECTED`, dan SEC-30 utuh: kode itu tetap byte-identik untuk tautan terpakai, kedaluwarsa, atau tak dikenal dengan email benar maupun salah (butir 7). `EMAIL_MISMATCH` hanya dicapai pemegang tautan yang sah (pemeriksaan sesudah tautan lolos, sebelum hash), jadi ia tidak membuka sumur untuk menebak tautan. Yang bocor hanya "tebakan email ini salah" bagi pemegang tautan, dibatasi pembatas alamat (satu hitungan `Failed`) dan tautan tak hangus (sisa yang diterima, butir 4).
  - Klien memetakan kode itu ke isian email lewat satu entri peta kode ke medan, jadi alasan "di bawah isian" tidak memerlukan Violations.
  - Dilaksanakan: nilai enum 18 di `error.proto` (komentar gaya 13 sampai 17, rujuk 10.4 butir 2 dan ADR-013 catatan 1.6), pemetaan `ErrEmailMismatch` di fungsi pemetaan `identity`, satu pesan tetap. DR-2 sudah ditulis di ADR-013 catatan 1.6 (baris tabel baru; belum di-commit). Langkah 4 dan 5 serta `wantEmailMismatch` memakai kode ini; `error_pb.ts` ikut dibangkitkan ulang (risiko 2 berlaku).

- Q3 (security-privacy): sisa risiko diterima. `Max-Age=0` dari permintaan lama yang tiba sesudah `SignIn` baru di browser yang sama dapat menghapus cookie baru; akibatnya satu masuk ulang, tanpa kebocoran. Penutupnya kewajiban klien SEC-172 (a) (hentikan permintaan berjalan sebelum layar lain). Dicatat 15 November 2026 di 10.5 SEC-161, bersama empat syarat DR-1, sehingga komentar kode dapat mengutipnya. Tidak menahan PR.
