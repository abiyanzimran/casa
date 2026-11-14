# FS-1 1.7 — rangka `read`, `sync`, `jobs` dan alat uji: rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** membangun tiga modul server dan tiga alat uji yang tersisa di kolom server 14 §3.3 langkah 1:
- `jobs`: mesin pekerjaan latar di atas tabel `job` (outbox di transaksi domain, `FOR UPDATE SKIP LOCKED`, `LISTEN jobs`, periksa tiap detik, coba ulang sampai 10 kali), dengan pekerjaan pertama: gugur undangan pada `expires_at`;
- `read`: satu jalur baca (SEC-06); proyeksi layar yang sudah ada dan `AccountEmail` pindah ke sana (putusan 1.6 Q10);
- `sync`: umpan per pembaca (ADR-017), fan-out di dalam transaksi perintah, kursor ber-MAC, RPC `Sync` dengan potret awal dan `resync_required`, untuk objek yang ada hari ini: rumah tangga, anggota, undangan hidup (putusan 1.6 Q12);
- pembangkit uji isolasi per kueri sqlc, rangka uji properti berstate (INV-A01, INV-A13, INV-A48, sinkron = bacaan langsung), dan pembangkit vektor emas bantuan isian (c) panjang nama.

Hasilnya: klien web 1.9b punya kontrak `Sync` yang lengkap untuk langkah 1, dan setiap perubahan rumah tangga sampai ke perangkat setiap anggota lewat jalur baca yang sama dengan layar.

**Arsitektur:**
- **Paket.** Tiga paket domain baru: `internal/jobs` (+ `jobsdb`), `internal/read` (+ `readdb`), `internal/sync` (+ `syncdb`). Nama paket `sync` mengikuti 09.2 §1; berkas di paket itu yang memerlukan pustaka standar mengimpornya sebagai `gosync`.
- **Arah impor** (09.2 §1 butir 10, catatan 14 November 2026): `household → sync → read`, `household → jobs`; `read` dan `jobs` hanya ke `database` dan paket sqlc-nya. `read` tidak mengimpor `household`, jadi tidak ada siklus. Penangan job didaftarkan `cmd/casa` ke `jobs.Runner` sebagai fungsi, jadi `jobs` tidak pernah mengimpor modul yang job-nya ia jalankan.
- **Fan-out lewat satu titik.** `household.RecordEvent` adalah satu-satunya penulis Peristiwa (09.2 §1) dan setiap perubahan rumah tangga menulis Peristiwa (INV-A01, PG-08). Karena itu `RecordEvent` memanggil `sync.Changed`, satu-satunya penulis `reader_feed` (09.2 §1 butir 10), di transaksi yang sama, dan tidak ada jalur tulis yang dapat lupa.
- **Pekerja.** `casa serve` menjalankan satu `jobs.Runner` dengan dua pool: `job_runner` untuk mengambil, `app_rw` untuk menjalankan dan menghapus (09.2 §3).
- **Bentuk tugas.** Potongan vertikal dengan uji terhadap PostgreSQL 18 nyata, seperti 1.6.

**Stack:**
- Go (toolchain di `server/go.mod`), pgx v5 (`LISTEN`/`WaitForNotification`), goose v3, sqlc (biner yang dipatok), connect-go v2, protovalidate, stdlib `crypto/hmac`, `crypto/sha256`, `encoding/base64`, `math/rand/v2`, `go/parser`, `reflect`.
- **Satu dependensi baru, hanya untuk uji:** `pgregory.net/rapid` **v1.3.0** (keputusan 16). Tidak ada dependensi npm baru.

**Spesifikasi:**
- 14 §3.3 langkah 1 (kolom server: "kerangka `read`, `sync` (umpan per pembaca, kursor, potret awal), `jobs`; pembangkit uji isolasi; rangka uji properti INV dan INV-A; pembangkit vektor emas"), §4.2, §4.3, §4.5, §4.6, §4.7 baris "Umpan per pembaca dan potret awal", §4.8.
- 09.2 §1 (baris `read`, `sync`, `jobs`, aturan dependensi 2 dan butir 9), §3 (lapis 1, 4, 5, 6; `job` dan `job` pengambilan), §5.1 (`reader_feed`, `reader_feed_counter`, `job`, catatan X3), §6, §8.
- 09.3 §1 ("Membaca tetap boleh", "Terbit saat pertama terlihat"), §3, §4 (hanya untuk "Yang sengaja tidak"), §9, §10.
- 09.4 §4 (baris Isolasi, Invarian alur dan persetujuan, Kontrak dan bantuan isian).
- ADR-005, ADR-013 butir 3, ADR-014, ADR-015 butir 4, ADR-016, ADR-017, ADR-018 butir 2, 4, 6, ADR-023.
- 06 INV-A01, INV-A13, INV-A48 (a) sampai (c) beserta kolom "Cara menguji"; FR-FAM-001/AC-4; FR-FAM-019/AC-3, AC-10 (bagian sinkron).
- 10.1 SEC-01, SEC-02, SEC-04, SEC-05, SEC-06; 10.2 SEC-47, SEC-48; 10.3 §3.2.
- Rencana 1.6: keputusan 12, 16; putusan Q8, Q10, Q12, Q17, X3; `_progress.md` "FS-1 1.6 selesai (14 November 2026)".
- `ENGINEERING-STANDARDS.md` §1 sampai §9.

**Pelaksana:** `backend-dev`. **Peninjau:** `code-reviewer` (dua lintasan) per PR. **Cabang:**
- `fs1/1.7a-jobs`: PR A, Tugas 1 dan 2;
- `fs1/1.7b-read-sync`: PR B, Tugas 3 sampai 6, dari `main` sesudah PR A di-merge;
- `fs1/1.7c-test-tools`: PR C, Tugas 7 sampai 9, dari `main` sesudah PR B di-merge.

Tugas 10 (verifikasi) dijalankan per PR. PR A lebih dulu karena X3 mensyaratkan job gugur ada sebelum `sync` mengirim Peristiwa undangan. Semua pertanyaan rencana sudah diputus ("Putusan atas pertanyaan rencana" di akhir).

## Keputusan yang diambil rencana ini

Semua keputusan di bawah sudah diputus (bagian "Putusan atas pertanyaan rencana" di akhir). Bila satu butir menulis ulang bunyi dokumen, bunyi dokumen yang menang.

1. **Gugur undangan lewat job terjadwal per undangan, bukan sapuan berkala** (X3; 09.3 §9 catatan 14 November 2026, putusan Q1 (e); 10.1 §10 catatan 1.7 Q3).
   - F01 mengantrekan, di transaksinya sendiri, satu job `household.expire_invitation`: `household_id` = rumah tangga undangan, `run_at = expires_at`, `unique_key` = pengenal undangan (teks), `payload` = 16 byte pengenal undangan.
   - Penangan memanggil penulis gugur yang sudah ada (`expire`, keputusan 12 rencana 1.6) dengan `Scope{HouseholdID}` tanpa akun. Tiga hasil:
     - undangan sudah tertutup: tidak ada tulisan, job selesai;
     - gugur tertulis: job selesai;
     - undangan masih hidup tetapi compare-and-set menyatakan belum jatuh tempo: galat `ErrNotYetDue`, dicoba ulang lewat jeda, tidak pernah selesai tanpa gugur tertulis (Q1 (e) syarat 2).
   - Pengambilan dan compare-and-set membaca satu sumber jam, `Now` yang sama (`jobs.Config.Now` dan `household.New(…, now, …)` diberi fungsi yang sama di `serve` dan di uji).
   - `expires_at` tidak berubah sesudah F01; `IssueInvitationLink` tidak memperpanjangnya (SEC-28 catatan Q7 (b)). Perubahan apa pun kelak memindahkan `run_at` job-nya di transaksi yang sama (Q1 (e) syarat 1).
   - Alasannya: sapuan berkala lintas rumah tangga memerlukan bacaan lintas rumah tangga yang tidak dimiliki `app_rw` (09.2 §3), sedangkan job per undangan berjalan di rumah tangganya sendiri. Gugur juga tertulis pada `expires_at` ditambah latensi pekerja (≤ 1 detik bila antrean kosong), tidak bergantung pada aktivitas siapa pun, sehingga baris umpannya tidak membuka waktu aktivitas (10.1 §10 catatan 1.7 Q3; sisa jendelanya ditinjau ulang sebelum GR-4).
   - Undangan uji yang lahir sebelum migrasi tidak diisi ulang: server uji hanya berisi data uji (14 §7.8), dan jalur malas X3 tetap menulis gugurnya. Basis data produksi (langkah 7) tidak boleh memuat undangan hidup tanpa job (10.1 §10 catatan 1.7 Q3 syarat 2).
   - Job dan jalur malas memanggil penulis `expire` yang sama, sehingga semua yang terbaca klien identik selain `pos` dan `seq` (10.1 §10 catatan 1.7 Q3 syarat 1; uji di Tugas 2 dan Tugas 6).
2. **Mesin `jobs`** (09.3 §9; 09.2 §3, §5.1):
   - **Tabel** menurut 09.2 §5.1 dengan catatan 14 November 2026 (putusan Q1 (b)):
     - `CHECK (kind IN ('household.expire_invitation'))`, diperluas per jenis seperti `object_event.object_kind`; `CHECK (attempts >= 0)`;
     - `UNIQUE (kind, unique_key)` diganti `CREATE UNIQUE INDEX job_unique ON job (household_id, kind, unique_key) NULLS NOT DISTINCT WHERE unique_key IS NOT NULL` (`household_id` memimpin, §3 lapis 2; job sistem tetap unik);
     - `job_due` untuk job baru; `job_locked ON job (locked_until) WHERE locked_until IS NOT NULL` untuk job yang menunggu coba ulang atau sewanya habis.
   - **Antre:** `INSERT … ON CONFLICT (household_id, kind, unique_key) WHERE unique_key IS NOT NULL DO NOTHING`.
   - **Ambil satu per satu.** `ClaimDue` dijalankan `job_runner` di satu `database.Write`: satu baris `run_at <= now AND attempts < 10 AND (locked_until IS NULL OR locked_until <= now)` urut `run_at`, `FOR UPDATE SKIP LOCKED`, lalu `locked_until = now + 5 menit` (sewa) dan `attempts = attempts + 1`. Percobaan dihitung saat diambil, jadi proses yang mati di tengah job tetap menghitung percobaannya. `// ponytail:` satu pekerja berurutan; batasnya satu job lambat menunda yang lain; jalan naiknya beberapa goroutine pengambil, karena `SKIP LOCKED` sudah aman untuk itu.
   - **Berhasil:** job dihapus `app_rw` dengan `Scope{HouseholdID: job.household_id}` (job sistem: `Scope{}`), di transaksi sesudah penangan. Penangan wajib idempoten (09.3 §9), jadi penghapusan yang gagal hanya menjalankan ulang penangan.
   - **Gagal:** `job_runner` menulis `last_error` (kode saja: `handler`, `panic`, `timeout`, `unknown_kind`) dan `locked_until = now + jeda`. Jeda sesudah percobaan ke-*n* = min(10 detik × 2^(n−1), 1 jam) **[Usulan]**: 10, 20, 40, …, 2560 detik. Coba ulang lewat `locked_until`, karena `job_runner` hanya boleh mengubah `locked_until`, `attempts`, `last_error` (09.2 §3).
   - **Berhenti dan menandai operator:** sesudah percobaan ke-10 gagal, `locked_until = 'infinity'` dan satu baris log `ERROR` "job stopped" dengan `kind`, `id`, dan kode. Tidak pernah `payload`, teks galat Go, atau isi lain (SEC-44).
   - **Batas waktu penangan** 4 menit **[Usulan]**, di bawah sewa 5 menit.
   - **Bangun:** `Run` memegang satu koneksi `job_runner` dengan `LISTEN jobs` dan menunggu `WaitForNotification` paling lama 1 detik, lalu `RunDue`. `Enqueue` menjalankan `pg_notify('jobs', '')` di transaksi domain; PostgreSQL mengantar `NOTIFY` saat commit, jadi tidak ada bangun tanpa job.
   - **Tanpa job berkala di 1.7.** Satu-satunya jenis terjadwal per objek (keputusan 1), jadi penjadwal berkala tidak dibuat (YAGNI); ia datang bersama job berkala pertama.
3. **Hak atas `job`** (09.2 §3 catatan 14 November 2026, putusan Q1 (b)):
   - `app_rw`: `SELECT, INSERT, DELETE`, tanpa `UPDATE`; kebijakan `FOR ALL TO app_rw` dengan teks 09.2 §3 (`household_id = H′ OR (household_id IS NULL AND H′ IS NULL)`) di `USING` dan `WITH CHECK`;
   - `job_runner`: `SELECT, UPDATE (locked_until, attempts, last_error)`, tidak berubah; kebijakan `FOR ALL TO job_runner USING (true)`.
   - **Keduanya wajib menyebut perannya:** kebijakan permisif digabung OR, jadi `USING (true)` tanpa `TO` membuat `app_rw` melihat job setiap rumah tangga. Gerbang katalog memeriksa setiap kebijakan `job` beserta perannya (`tableRule.policies` berisi teks deparse dengan `roles={…}`).
4. **Kredensial baru `casa serve`** (pola `LoadCredential=` 09.4 §1): `job-database-url` (string koneksi peran `job_runner`) dan `sync-cursor-key` (64 karakter heksa, 32 byte acak, misalnya `openssl rand -hex 32`). Keduanya wajib; tanpa itu `serve` gagal saat mulai dan tidak mencetak nilainya. Server uji membuat kunci baru; produksi membuat kunci baru di langkah 7 (14 §7.8). Rotasi kunci tidak dibuat di pilot; kunci baru membuat setiap kursor ditolak, dan klien memulai lagi dengan potret (keputusan 9).
5. **Yang ditunda** (alasan di "Yang sengaja tidak"):
   - `identity.send_mail` ke langkah 7, bersama pengirim SES (09.3 §9 catatan 14 November 2026; 10.1 §9 catatan 1.7 Q2);
   - `sync.prune_feed`, `idempotency.prune` ke langkah 7 (09.2 §3 catatan 14 November 2026, putusan Q1 (d));
   - `identity.prune_tokens` dan pembersihan baris pendaftaran 24 jam (10.3 §3.2) ke langkah 6, bersama hak `DELETE` atas tabel `identity` yang memang dibawa `lifecycle` (catatan `_progress.md` 1.6), tetap sebelum GR-4.
6. **`read` memegang proyeksi layar** (09.2 §1 baris `read`, aturan 2; SEC-06; putusan 1.6 Q10).
   - `ListMemberships` dan `GetOpenInvitation` membaca lewat `read.Memberships` dan `read.OpenInvitation`. `household.Memberships`, `membershipsOf`, `OpenInvitation` dan kuerinya pindah, tanpa perubahan perilaku: uji 1.6 tetap lulus tanpa disunting.
   - `identity.AccountEmail` pindah menjadi `read.AccountEmail` dan ekspornya dihapus dari `identity`, sehingga pengecualian sementara 09.2 §1 butir 9 berakhir (butir 10). Pemanggilnya tetap satu, `GetOpenInvitation`; aturan isinya tidak berubah (hanya email, hanya bagi pengundang undangan `accepted`).
   - Bacaan penjaga di jalur tulis (`FullAdultMembers`, `MemberIn`, `LiveInvitation`, token F02, `PreviewInvitation`) tetap di `household`: itu bagian perintah, bukan proyeksi.
   - `read` mendefinisikan `Reader` sendiri, karena tidak mengimpor `household`. Lapisan API mengubah `household.Member` menjadi `read.Reader`.
7. **Objek sinkron dan proyeksinya** (09.3 §3 "Isi diambil lewat jalur baca"; putusan 1.6 Q12):
   - `Household {household_id, name, time_zone, currency}`;
   - `Member {member_id, display_name}` (pesan kontrak yang sudah ada, dipakai ulang);
   - `Invitation {invitation_id, state, expires_at, inviter_member_id}`, hanya selama hidup (`open` atau `accepted`) **dan** `now < expires_at`, aturan yang sama dengan `GetOpenInvitation`. Nama tampilan dan email penerima tidak pernah masuk sinkron (Q12: hanya keadaan; email hanya daring).
   - Undangan yang tidak lagi hidup dikirim sebagai `PURGE`: cache perangkat hanya memuat objek hidup (09.3 §3 butir 2 potret), dan padanan "kosong" `GetOpenInvitation` di cache adalah membuang objeknya.
   - `can_read` hari ini: setiap anggota rumah tangga membaca ketiga jenis objek itu (06 §3.1; rencana 1.6 keputusan 16, `GetOpenInvitation` terbaca setiap anggota). Aturannya satu fungsi `canRead` di `read` dengan `switch` per jenis, sehingga langkah 2 menambah visibilitas di satu tempat.
8. **Fan-out** (09.2 §6 langkah 6; ADR-017 butir 3):
   - `sync.Changed(ctx, tx, household, kind, id)` adalah satu-satunya penulis `reader_feed` (09.2 §1 butir 10) dan menulis satu baris `upsert` per anggota rumah tangga sesudah perubahan. Untuk Peristiwa rumah tangga (`created`, `member_joined`) ia menulis objek rumah tangga dan setiap objek anggota, karena keanggotaan berubah hanya bersama Peristiwa rumah tangga (09.2 §5.1).
   - `pos` naik tanpa celah lewat `reader_feed_counter`: satu `INSERT … ON CONFLICT DO UPDATE SET last_pos = last_pos + n RETURNING last_pos` per pembaca. Dua perintah yang mengenai pembaca yang sama saling bertabrakan di bawah `SERIALIZABLE` dan dicoba ulang (ADR-017 "Konsekuensi").
   - `IssueInvitationLink` tidak menulis Peristiwa dan tidak mengubah proyeksi (token tidak diproyeksikan), jadi tidak menulis umpan.
   - Baris `purge` saat akses menyempit dibuat bersama penyempitan pertama (F06, langkah 6). `PURGE` saat ditarik sudah ada di 1.7 (keputusan 7).
   - `NOTIFY feed` ditunda ke langkah 5, bersama hub SSE: belum ada pendengarnya.
9. **Kursor** (ADR-017 butir 3, 8; 09.3 §3). Tata letak (tidak ditentukan tanda tangan, jadi ditulis di sini):

   ```text
   body   = version (1 byte, = 1) | mode (1 byte: 0 feed, 1 snapshot)
            | household (16) | member (16) | pos (8, uint64 big-endian)
            [ mode 1 saja: kind (1 byte: 1 household, 2 member, 3 invitation) | after (16) ]
   mac    = HMAC-SHA256(key, "casa.sync.cursor.v1" || body)[:16]
   cursor = base64url tanpa padding (body || mac)
   ```

   - Ditolak dengan satu galat `ErrCursorRejected`: base64 rusak, panjang tidak cocok dengan mode, versi atau mode atau jenis tak dikenal, MAC tidak sama (`hmac.Equal`), dan `household` atau `member` yang bukan milik pembaca.
   - Di API, `ErrCursorRejected` dijawab persis seperti rumah tangga yang tidak ada (SEC-05).
   - Komentar kontrak menyatakan bagi klien: `NotFound` dengan kursor tidak kosong → buang kursor, coba sekali dengan kursor kosong; `NotFound` lagi → rumah tangga tidak terbaca.
10. **Potret berhalaman lintas panggilan** (ADR-017 butir 7 dan 09.3 §3, masing-masing dengan catatan 14 November 2026, putusan Q1 (c)).
    - Kursor kosong: `P0` = `last_pos` pembaca (0 bila belum ada baris penghitung), dibaca di transaksi halaman pertama. Setiap halaman satu transaksi `REPEATABLE READ READ ONLY` sendiri. Urutannya jenis (rumah tangga, anggota, undangan), lalu pengenal; `P0` dan tanda halaman dibawa di kursor mode 1.
    - Sesudah halaman terakhir, `next_cursor` = kursor umpan di `P0`. Setiap perubahan sesudah `P0` punya baris umpan, jadi objek yang berubah, lahir, atau tertutup di antara halaman sampai lewat umpan. `UPSERT` ganda tidak merusak karena idempoten.
    - Alasannya: satu transaksi tidak dapat hidup di antara dua RPC tanpa state di server. Ukuran potret hari ini satu halaman, tetapi halaman tetap dibangun (ES §3).
11. **`resync_required`** (ADR-017 butir 6; 09.3 §3 "Pemangkasan" catatan 14 November 2026, putusan Q1 (d)): kursor umpan `P` dengan `last_pos = L` mendapat resync bila:
    - `P > L` (kursor dari basis data yang dipulihkan);
    - baris yang dikembalikan bukan tepat `P + 1, P + 2, …` tanpa celah;
    - halaman yang tidak penuh tidak berakhir di `L`.

    Pemeriksaan tidak bergantung pada urutan `created_at` (09.2 §3 catatan Q1 (d)).

    Resync dijawab tanpa perubahan, dengan `next_cursor` kosong. Deteksi ini dibangun dan diuji sekarang dengan menghapus baris sebagai admin, walau job pemangkasnya ditunda (keputusan 5).
12. **Versi pemberitahuan lewat sinkron** (09.3 §1 "Terbit saat pertama terlihat"; FR-FAM-019/AC-3, AC-10): handler `Sync` memanggil `api.NoticeHeld` yang sudah ada, sebelum membaca, dan mengisi `SyncResponse.operator_notice_unacknowledged`. Teksnya tetap lewat `GetOperatorNotice`, yang daring selamanya (Q12). `Sync` adalah bacaan: tanpa `CommandOptions`, tanpa kunci idempotensi, tidak pernah ditahan `NOTICE_UNREAD` (09.3 §1 "Membaca tetap boleh").
13. **`purpose` ada sejak sekarang** (ADR-018 butir 2, 4): `SyncPurpose` wajib (`defined_only`, bukan 0). Server baru membacanya di langkah 5 (tanda terkirim, INV-A33). Menambahkannya sebagai medan wajib kelak akan merusak klien yang tidak mengirimnya.
14. **Batas halaman:** `limit` 0 berarti 100; paling banyak 500 (`lte: 500`) **[Usulan]**.
15. **Pembangkit uji isolasi** (09.2 §3 lapis 6; 09.4 §4 baris Isolasi; SEC-02 kolom "Cara menguji"). Algoritmanya ditulis di Tugas 7 karena tanda tangan tidak menentukannya. Intinya: setiap kueri sqlc ditemukan dari kode hasil generate, lalu setiap parameter pengenal diganti satu per satu dengan nilai rumah tangga lain (X), dan hasilnya harus sama persis dengan nilai yang tidak pernah ada.
16. **Uji properti** (14 §4.3; 09.4 §4; putusan 1.6 Q17):
    - **Pustaka.** `pgregory.net/rapid` v1.3.0 dipilih karena menyusutkan otomatis dan punya mesin state (`T.Repeat`), dua hal yang diminta 14 §4.3. Alternatifnya tidak memadai: `testing/quick` dibekukan, tanpa penyusutan dan tanpa mesin state; fuzzing `testing.F` tanpa rangkaian berstate; `gopter` 0.x memerlukan baris ADR-023. Rapid ≥ 1.0, jadi ADR-023 tidak berlaku; catatan ES §8 ada di Tugas 8.
    - **Pemeriksa di kode uji.** 14 §4.3 **[Usulan]** ingin pemeriksa yang sama dipakai `ops.invariant_check` (langkah 7). Sampai pemakai produksi itu ada, pemeriksa tinggal di berkas `_test.go` (ES §1: tanpa kode mati) dan dipindah ke paket produksi di langkah 7 tanpa ditulis ulang.
    - **Benih.** Per perubahan memakai benih tetap (`RAPID_SEED=1`), karena uji harus deterministik (ES §6). Putaran mingguan memakai benih baru, sesuai 14 §4.3. Kegagalan mencetak benihnya; kegagalan yang sudah disusutkan disimpan sebagai uji contoh bernama (14 §4.3). Berkas gagal rapid tidak di-commit (`-rapid.nofailfile` dan `.gitignore`).
    - **Jumlah kasus:** `RAPID_CHECKS=20` per perubahan **[Asumsi]**, waktunya diukur dan dicantumkan di PR C; mingguan 1000.
    - **INV (07)** belum punya objek (uang datang di langkah 2). Rangkanya adalah mesin dan pembangkit perintah yang sama; langkah 2 menambah pembangun perintah dan pemeriksa INV.
17. **Vektor emas** (ES §2 (c); 14 §4.5; 01 §3 butir 3). Hari ini hanya bantuan (c) panjang teks yang punya aturan server: nama rumah tangga 60 dan nama tampilan 40 karakter tampak sesudah dipangkas dan NFC (`household.CleanName`; FAM.md FR-FAM-001 catatan 14 November 2026). Berkas `golden/input-shape/name-length.json` dibangkitkan dari `CleanName`. Bantuan (a), (b), dan (c) bilangan bulat positif datang bersama `ledger` di langkah 2. Pengguna TypeScript dan gerbang "berkas berubah tanpa uji klien yang lulus" ditambahkan bersama pengguna pertamanya, layar buat rumah tangga di 1.10 (`web-dev`).

## Batasan global

- **Warisan.** Batasan 1.1a sampai 1.6 berlaku: ES §1 sampai §9; alat dengan versi yang dipatok. Satu-satunya dependensi baru adalah rapid, untuk uji, dengan catatan ES §8.
- **SQL.**
  - Semua SQL produksi lewat sqlc dan berparameter.
  - Setiap kueri baru punya indeks; `EXPLAIN` dilampirkan di PR (ES §3).
  - Tidak ada kueri tanpa batas: umpan dan potret dengan `LIMIT` (keputusan 14); anggota tetap paling banyak 20 (1.6).
- **Log dan galat.** Satu baris log baru saja: "job stopped" (keputusan 2), dengan `kind`, `id`, dan kode. Tidak ada token, email, nama, isi `payload`, atau teks galat PostgreSQL di log, di galat klien, atau di keluaran `casa` (SEC-44; ES §4).
- **Uji.**
  - PostgreSQL 18 nyata, gagal keras tanpa `CASA_TEST_ADMIN_URL`.
  - Tanpa `t.Parallel` di paket yang menyentuh basis data.
  - Jam lewat `Now`/`Config.Now`, tanpa tidur. `Runner.RunDue` dipanggil langsung oleh uji; hanya uji `Run` yang menunggu (dengan batas 5 detik).
  - Iterasi: `cd server && CASA_TEST_ADMIN_URL='postgres://postgres@127.0.0.1:5432/postgres?sslmode=disable' go test -race -p 1 -run '<Uji>' ./internal/<paket>/`. `make check` hanya di langkah akhir PR.
- **Kode khusus uji (X7).** Kait dan fixture hanya di `_test.go`. Setiap aturan punya satu implementasi:
  - penulis gugur (`expire`);
  - penulis Peristiwa (`RecordEvent`);
  - penulis umpan (`sync.Changed`);
  - aturan baca (`read.canRead`);
  - codec kursor (`CursorKey`).
- **Penamaan (QG-06, X6).** Uji membawa ID AC hanya bila membuktikan seluruh AC. ID SEC, INV, dan INV-A boleh di nama uji yang membuktikan aturan itu. Pengenal mengikuti glosarium; komentar, pengenal, dan pesan commit berbahasa Inggris; setiap aturan di kode mengutip ID atau bagian dokumennya.
- **Migrasi.** Hanya menambah (09.2 §10), dengan arah turun; `server-test-old-binary` tetap hijau.
- **Kontrak.** `buf lint` dan `buf breaking` lulus; kode Go dan TypeScript dibangkitkan ulang dan di-commit (14 §4.6).

## Fokus tinjauan

1. **Kursor palsu, asing, atau diubah.** Kursor anggota lain, rumah tangga lain, satu karakter diubah, terpotong, atau bukan base64 harus dijawab persis seperti rumah tangga yang tidak ada (kode, pesan, detail). Galat kursor tersendiri akan membuka bahwa rumah tangga itu ada. Diuji di Tugas 5 (`TestCursor`) dan Tugas 6 (`SEC-05 a cursor of another member, of another household, or changed in one character is not found like a household that does not exist`).
2. **Perubahan di antara halaman potret.** Objek yang lahir, berubah, atau tertutup di antara dua halaman tidak boleh hilang dan tidak boleh tertinggal basi sesudah umpan dikejar. Diuji di Tugas 5 (`a change between snapshot pages arrives through the feed`).
3. **Job berjalan dua kali atau tidak pernah.** Dua pekerja serentak, sewa yang habis sesudah proses mati, dan penangan yang dijalankan ulang sesudah commit: tepat satu Peristiwa gugur, tidak ada job yang macet diam-diam. Diuji di Tugas 1 (`two runners at once claim each job once`, `a lease that ran out is claimed again`) dan Tugas 2 (`a re-run of the expiry writes nothing more`).
4. **Jalur tulis tanpa baris umpan.** Gugur oleh job, F04, penarikan, dan penutupan seperti penarikan (`CloseAsWithdrawn`) yang tidak sampai ke perangkat, sehingga sinkron berbeda dari bacaan langsung. Diuji di Tugas 4 (`every invitation transition writes one row per member`) dan Tugas 8 (`sync after every command equals the direct read for every member`).
5. **Kursor di luar umpan yang ada.** Kursor yang lebih tua dari baris yang tersisa, atau lebih maju dari penghitung (basis data dipulihkan), harus menghasilkan `resync_required`, tidak pernah lompatan diam-diam. Diuji di Tugas 5 (`a cursor behind the pruned feed asks for a resync`, `a cursor ahead of the feed asks for a resync`).

## Yang sengaja tidak di tugas ini

- **Langkah 2:**
  - vektor emas (a) selisih baris, (b) bagi rata, (c) bilangan bulat positif;
  - pembangun perintah dan pemeriksa INV-01 sampai INV-41, termasuk INV-10 (kiriman ulang di 1.7 diuji sebagai ADR-005, tanpa ID INV);
  - uji SEC-04 byte demi byte (memerlukan entri pribadi).
- **Langkah 5:** hub SSE `GET /v1/stream`, `NOTIFY feed`, pembacaan `purpose` (tanda terkirim, INV-A33), `retract_push_refs` (bersama push, langkah 4).
- **Langkah 6 (`lifecycle`):**
  - baris `purge` saat akses menyempit (F06) dan uji SEC-48 dengan pembaca yang kehilangan akses;
  - `identity.prune_tokens`; pembersihan baris pendaftaran yang tidak selesai paling lambat 24 jam sesudah tautannya kedaluwarsa (10.3 §3.2), dengan jalur hapus dan hibah yang ditulis `architect` di 09.2. Keduanya memerlukan hak `DELETE` atas tabel `identity`, yang memang dibawa `lifecycle`. Batas akhirnya tetap sebelum GR-4.
- **Langkah 7:**
  - `identity.send_mail`, bersama pengirim SES (09.3 §9 catatan 14 November 2026). Job itu memenuhi enam syarat `security-privacy` di 10.1 §9 catatan 1.7 Q2: (1) `payload` hanya pengenal; (2) token dibuat penangan saat mengirim, hanya hash disimpan, coba ulang membuat token baru; (3) keadaan diperiksa ulang sebelum membuat token; (4) batas laju SEC-12 dihitung saat permintaan; (5) pemulihan mengantre sesudah jawaban, pendaftaran di transaksi barisnya; (6) pencarian `signup_verification` menurut pengenal lewat jalur yang ditulis `architect` di 09.2 §3;
  - `sync.prune_feed` dan `idempotency.prune` menurut 09.2 §3 catatan 14 November 2026 (putusan Q1 (d)): penangan memakai pool `job_runner`, hak `SELECT` atas kolom kunci dan waktu serta `DELETE` dengan kebijakan `TO job_runner USING (created_at < now() - interval '90 days')`, `$cutoff` dari `Now`, berhalaman 1000, indeks `created_at` di migrasi itu;
  - isi ulang job gugur bagi undangan hidup tanpa job sebelum `sync` melayani klien produksi (10.1 §10 catatan 1.7 Q3 syarat 2);
  - pemeriksa invarian pindah ke kode produksi untuk `ops.invariant_check`.
- **1.9b (`web-dev`):** putaran sinkron klien, penyimpanan kursor, dan reaksi atas `resync_required` dan `NotFound` (keputusan 9). Bagian SSE 1.9b menunggu hub di langkah 5 (konflik dilaporkan ke lead).
- **1.10 (`web-dev`):** pengguna TypeScript vektor panjang nama beserta gerbangnya; layar LY-05 dari `operator_notice_unacknowledged`.

**AC sebagian.** Uji AC berikut tidak membawa ID di nama (X6):

| AC | Bagian di 1.7 (uji) | Sisanya |
|---|---|---|
| FR-FAM-001/AC-4 | permutasi N = 2 atas rangkaian acak (Tugas 8, subuji `INV-A13 …`) | N > 2 mustahil di pilot (paling banyak dua dewasa penuh); "setiap tampilan": layar 1.10 |
| FR-FAM-019/AC-3 | perangkat mengetahui versi baru pada sinkron berikutnya (Tugas 6) | membaca di layar: 1.10 |
| FR-FAM-019/AC-10 | sinkron menyatakan versi yang belum diakui (Tugas 6) | LY-05 seketika dan antrean klien: 1.10 |

## Celah langkah 1 (kolom server) sesudah 1.7

Butir kolom server 14 §3.3 langkah 1 yang belum dibangun 1.2 sampai 1.7:
- vektor emas (a), (b), dan (c) bilangan bulat positif: menunggu `ledger` (langkah 2);
- pemeriksa INV (07): menunggu objek uang (langkah 2); rangkanya ada.

Tidak di kolom itu dan dipindah ke langkah 7: `identity.send_mail` (09.3 §9 catatan 14 November 2026). `traceability/scope.txt` diperbarui lead di akhir langkah 1.

## Putusan pra-eksekusi (14 November 2026, lead)

Hasil pemindaian pra-eksekusi. Putusan ini mengungguli teks tugas yang disebut.

- **P1 (Tugas 2, 3, 6; "uji 1.6 tidak disunting").** Harness `server/internal/api/household_test.go` (dan harness API lain yang memanggil fungsi `Register…Service`) boleh disunting **hanya pada pemasangan** (argumen baru, runner, `read`, `sync`, `Sync`). Asersi, nama, dan isi uji 1.6 tidak berubah. Klaim "diff kosong" diganti: `git diff main -- server/internal/api/*_test.go` hanya menyentuh fungsi harness.
- **P2 (Tugas 2 langkah 3).** `expire` tetap satu-satunya penulis gugur, tetapi mengembalikan keadaan yang dibacanya, sehingga job membedakan "sudah tertutup dalam keadaan lain" (selesai) dari "masih hidup tetapi belum jatuh tempo" (`ErrNotYetDue`).
- **P3 (Tugas 2, uji "a live invitation the clock finds not yet due is retried").** Uji menurunkan `run_at` lewat fixture admin di `_test.go` sebelum `RunDue`, lalu menegaskan job dicoba ulang dan undangan tetap `open`.
- **P4 (Tugas 4, uji F04).** `Confirm` tidak diubah urutannya. Peristiwa undangan dari F04 sampai ke umpan anggota yang ada saat Peristiwa ditulis (pengundang); anggota baru menerima baris rumah tangga dan anggota. Uji menegaskan persis itu.
- **P5 (keputusan 7, X7; Tugas 4, 5).** Aturan baca satu implementasi: `read.canRead` lahir di Tugas 4, dan `Readers` (fan-out) serta proyeksi Tugas 5 memanggilnya. `read` boleh punya kueri keanggotaan sendiri di `readdb` untuk jalur baca; `household.MemberIn` tetap pemeriksaan sisi perintah (SEC-151). Keduanya dicatat sebagai dua lapis, bukan dua implementasi satu aturan.
- **P6 (Tugas 8, Makefile).** Tidak ada flag `-rapid.*` di Makefile. Uji properti membaca `RAPID_CHECKS` dan `RAPID_SEED` sendiri (di `TestMain` paket ujinya, lewat `flag.Set` sebelum `m.Run`), dengan bawaan seed tetap dan 20 putaran bila variabel kosong. Job CI mingguan mengisi variabel itu.
- **P7 (Tugas 7).** Peran dicatat per paket kueri dan uji menjalankan kueri dengan peran itu. Kueri `job_runner` (`ClaimDue`, `FailJob`, dan sejenisnya) dijalankan sebagai `job_runner` dan masuk daftar pengecualian berdasar sumber (job sistem lintas rumah tangga menurut desain, 09.2 §3 catatan Q1 (b)).
- **P8 (Tugas 10, bukti merah "`attempts < 10` dibuang").** Diganti uji dengan fixture admin `attempts = 10` dan `locked_until` kosong: job tidak diklaim. Baris bukti merah memakai uji itu.
- **P9 (keputusan 9; Tugas 5, 6).** Kursor didekode dengan `base64.RawURLEncoding.Strict()`, sehingga perubahan satu karakter tidak pernah menghasilkan byte yang sama.
- **P10 (Tugas 7 langkah 4).** Tidak menyalin `setScope`. Kueri dijalankan lewat `database.Write` (atau fungsi ekspor `database` yang sudah ada) pada pool peran yang dicatat; digest dibaca pool admin sebelum dan sesudah rollback paksa.
- **P11 (Tugas 8).** Uji "every write command has a builder" dibatasi ke `HouseholdService` dan dinamai begitu; `SignOut` dan `ChangePassword` tidak menyentuh keadaan rumah tangga.
- **P12 (Tugas 9).** Karakter tak terlihat dan spasi ditulis sebagai escape Go (` `, `​`, `﻿`, `‮`, `️`), tidak sebagai literal.

## Berkas yang mungkin bentrok

- **`web/src/lib/gen/` dibangkitkan ulang:** `casa/v1/sync_pb.ts` baru; `household_pb.ts` berubah.
- **`server/cmd/casa/serve.go`:** dua kredensial baru dan runner.
- **`server/internal/api/household.go`:** `ListMemberships` dan `GetOpenInvitation` membaca lewat `read`.
- **`server/go.mod`, `server/go.sum`:** rapid.
- **`Makefile`, `.github/workflows/`:** `golden-generate`, `RAPID_*`, alur mingguan.

---

### Tugas 1: Tabel `job` dan mesin `jobs`

**Berkas:**
- **Buat:**
  - `server/internal/schema/00009_job.sql`
  - `server/internal/jobs/jobs.go`, `server/internal/jobs/queries.sql`, hasil sqlc `server/internal/jobs/jobsdb/*.go`
  - `server/internal/jobs/jobs_test.go`
- **Ubah:**
  - `server/internal/schema/catalog_test.go`: baris `job`; `tableRule.policy` menjadi `policies []string` (satu butir untuk tabel biasa, dua untuk `job`), dan teks yang dibandingkan memuat peran setiap kebijakan
  - `server/sqlc.yaml` (blok `jobsdb`)
  - `server/internal/architecture/imports_test.go`: `"internal/jobs": {"internal/database", "internal/jobs/jobsdb"}`, `"internal/jobs/jobsdb": {}` (09.2 §1 butir 10)

**Antarmuka:**
- **Consumes:** `database.Write`, `database.Read`, `database.Scope`.
- **Produces:**

```go
// MaxAttempts is the retry bound of 09.3 §9: the job stops after the 10th failed attempt.
const MaxAttempts = 10

// Job is a claimed job: its payload holds identifiers only (09.2 §5.1; SEC-44).
type Job struct {
	ID, HouseholdID pgtype.UUID // HouseholdID invalid for a system job
	Kind            string
	Payload         []byte
	Attempt         int32 // 1 for the first run
}

// Handler runs one job; it must be idempotent (09.3 §9).
type Handler func(context.Context, Job) error

// NewJob is a job to enqueue in the domain transaction (outbox; 09.3 §9, §10).
type NewJob struct {
	HouseholdID pgtype.UUID
	Kind        string
	RunAt       time.Time
	UniqueKey   string
	Payload     []byte
}

// Enqueue writes j in tx with ON CONFLICT (household_id, kind, unique_key)
// WHERE unique_key IS NOT NULL DO NOTHING and pg_notify('jobs', ''), delivered
// at commit (09.2 §5.1 note of 14 November 2026).
func Enqueue(ctx context.Context, tx pgx.Tx, j NewJob) error

type Config struct {
	Claim    *pgxpool.Pool // role job_runner
	App      *pgxpool.Pool // role app_rw
	Now      func() time.Time
	Logger   *slog.Logger
	Handlers map[string]Handler
}

func NewRunner(c Config) *Runner
// RunDue claims and runs due jobs one at a time until none is due; ran counts them.
func (r *Runner) RunDue(ctx context.Context) (ran int, err error)
// Run waits on LISTEN jobs for at most 1 s, then RunDue, until ctx ends; it
// returns after the running job finishes.
func (r *Runner) Run(ctx context.Context) error
```

Kueri `jobsdb`: `InsertJob` (`:exec`), `Notify` (`SELECT pg_notify('jobs', '')`), `ClaimDue` (`:one`, keputusan 2), `DeleteJob` (`:execrows`), `FailJob` (`:exec`: `last_error`, `locked_until`).

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
// TestJobTable proves the job table of 09.2 §3, §5.1 and plan 1.7 decision 3.
func TestJobTable(t *testing.T) {
	t.Run("a household transaction sees only its own jobs, and a system job only outside a household", …)
	// fixtures as admin: job of X, job of Y, system job; Scope{HouseholdID: X} → X's only;
	// Scope{} → the system job only
	t.Run("a household transaction cannot write a job of another household or a system job", …) // → 42501
	t.Run("job_runner reads every job and updates only locked_until, attempts and last_error", …)
	// UPDATE run_at or payload as job_runner → 42501; DELETE → 42501
	t.Run("app_rw cannot update a job", …) // → 42501
	t.Run("a job of an unknown kind is refused", …) // → 23514
	t.Run("the same kind and unique key in two households are two jobs, and a system job is written once", …)
}

// in catalog_test: the job policies name their roles (09.2 §3 note Q1 (b))
t.Run("each policy of job names its role", …) // TO app_rw and TO job_runner; a policy without TO fails

func TestRunner(t *testing.T) {
	t.Run("a due job runs once with attempt 1 and is gone", …)
	t.Run("a job before its run_at does not run", …)
	t.Run("a failed job waits 10 s, then 20 s, and stops after 10 attempts with its code", …)
	// clock advanced by the backoff each time: runs exactly at 10 s, 20 s, 40 s, …;
	// after the 10th failure: attempts 10, last_error "handler", locked_until infinity; RunDue runs it no more
	t.Run("a stopped job is logged with its kind and id and no payload", …)
	// slog to a buffer: one ERROR "job stopped" with kind and id; the payload bytes and the handler's error text absent
	t.Run("a job whose handler panics fails with code panic", …)
	t.Run("a job of a kind without handler fails with code unknown_kind", …)
	t.Run("a lease that ran out is claimed again", …)
	// admin sets locked_until = now − 1 s and attempts 1 on a due job → it runs with Attempt 2
	t.Run("two runners at once claim each job once", …)
	// 20 due jobs, two Runners calling RunDue behind a barrier; each job id seen exactly once
	t.Run("an enqueued job with the same household, kind and unique key is written once", …)
	t.Run("a job enqueued in a rolled-back transaction does not exist", …)
	t.Run("Run wakes on the notification of a committed job", …)
	// Run in a goroutine; Enqueue a job due now in a committed Write; handler signals within 5 s
	t.Run("Run returns when its context ends, after the running job finishes", …)
}
```

- [ ] **Langkah 2:** `go test -race -p 1 ./internal/jobs/ ./internal/schema/` → GAGAL (paket dan tabel tidak ada).
- [ ] **Langkah 3:** tulis migrasi (naik dan turun), kueri, `sqlc.yaml`; `make sqlc-generate`; tulis `jobs.go` menurut keputusan 2. `backoff(attempt int32) time.Duration` = min(10 s × 2^(attempt−1), 1 h). Penangan dijalankan dengan `context.WithTimeout(context.WithoutCancel(ctx), 4*time.Minute)` dan `recover`.
- [ ] **Langkah 4:** `go test -race -p 1 ./internal/jobs/ ./internal/schema/ ./internal/architecture/` dan `make sqlc-generated` → LULUS. `EXPLAIN` `ClaimDue` (job baru dan job dengan `locked_until` lewat) disimpan untuk PR.
- [ ] **Langkah 5: commit** `Add the job table and the jobs runner: outbox enqueue, SKIP LOCKED claim, backoff to 10 attempts, operator flag (09.3 §9; 09.2 §3, §5.1; ADR-015 item 4)`.

### Tugas 2: Gugur undangan oleh job dan pekerja di `casa serve`

**Berkas:**
- **Ubah:**
  - `server/internal/household/invitation.go`: `CreateInvitation` mengantrekan job; penangan `ExpireInvitationJob` (jenisnya sudah di CHECK `job` Tugas 1)
  - `server/cmd/casa/serve.go`, `server/cmd/casa/serve_test.go`: kredensial `job-database-url`, `jobs.NewRunner`, `Run` di goroutine, dihentikan sebelum pool ditutup
  - `server/internal/architecture/imports_test.go`: `internal/household` + `internal/jobs`; `cmd/casa` + `internal/jobs` (09.2 §1 butir 10)
  - `README.md` akar (kredensial baru, runbook server uji). Runbook menambah: job `household.expire_invitation` yang berhenti (log `ERROR` "job stopped") diperbaiki operator sebelum undangan berikutnya dibuat (10.1 §10 catatan 1.7 Q3 syarat 3)
- **Uji:** `server/internal/api/expiry_job_test.go` (harness 1.6)

**Antarmuka:**
- **Consumes:** `jobs.Enqueue`, `jobs.Job`, `jobs.NewRunner` (Tugas 1); `(*household.Service).expire` (1.6).
- **Produces:**

```go
// ExpireInvitationKind is the job of plan 1.7 decision 1.
const ExpireInvitationKind = "household.expire_invitation"

// ErrNotYetDue: the invitation is live but the compare-and-set finds it not
// yet due; the job is retried, never done (09.3 §9 note of 14 November 2026 (2)).
var ErrNotYetDue = errors.New("invitation not yet due")

// ExpireInvitationJob writes the expiry of the invitation in j.Payload in
// j's household when it is live and due (06 §3.6 "Gugur karena waktu"; X3);
// a closed invitation is done without a write; a live one not yet due is
// ErrNotYetDue.
func (s *Service) ExpireInvitationJob(ctx context.Context, j jobs.Job) error
```

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestInvitationExpiryJob(t *testing.T) {
	t.Run("F01 schedules the expiry at expires_at", …)
	// one job: kind household.expire_invitation, household_id = the household, run_at = expires_at,
	// unique_key = the invitation id, payload = its 16 bytes
	t.Run("a refused F01 schedules nothing", …) // HOUSEHOLD_FULL → no job row
	t.Run("at expires_at the job writes the expiry at expires_at", …)
	// clock to expires_at, RunDue: state expired, closed_at = expires_at, one event actor_kind system,
	// rule invitation.expired, command expire_invitation, server_time = expires_at; the job row is gone
	t.Run("INV-A48 (b) an accepted invitation past its expiry is expired by the job, never confirmed", …)
	t.Run("the job on a cancelled invitation writes nothing and is done", …)
	t.Run("a re-run of the expiry writes nothing more", …) // ExpireInvitationJob twice → one event
	t.Run("a live invitation the clock finds not yet due is retried, not done", …)
	// handler called with the clock 1 s before expires_at → ErrNotYetDue; through RunDue: job kept with
	// last_error "handler" and locked_until = now + 10 s; at expires_at the next run writes the expiry
	t.Run("issuing a link keeps expires_at and the job's run_at", …)
	t.Run("the job and the lazy path write the same expiry", …)
	// two equivalent invitations in two households; one expired by the job, one by the X3 lazy path (F01
	// after expires_at): object_event rows equal apart from household, object id and seq (10.1 §10 note
	// 1.7 Q3 condition 1). The Sync half of the comparison is in Task 6: Sync does not exist in PR A
}
```

Ditambah di `serve_test.go`: `serve refuses to start without the job database credential` dan `serve runs due jobs` (undangan uji yang `expires_at`-nya sudah lewat menjadi `expired` tanpa permintaan lain, ditunggu paling lama 5 detik).

- [ ] **Langkah 2:** `go test -race -p 1 -run 'TestInvitationExpiryJob' ./internal/api/ && go test -race -p 1 ./cmd/casa/` → GAGAL.
- [ ] **Langkah 3:** tulis kodenya. `Enqueue` di transaksi F01 sesudah undangan ditulis. `ExpireInvitationJob` memanggil `s.expire(ctx, j.HouseholdID, pgtype.UUID{}, id)`: `expired` → selesai; tertutup dengan state lain → selesai; masih hidup → `ErrNotYetDue`. Payload yang bukan 16 byte adalah galat (job berhenti dan menandai operator). `serve` memberi `time.Now` yang sama kepada `household.New` dan `jobs.Config.Now`.
- [ ] **Langkah 4:** jalankan perintah langkah 2 → LULUS; seluruh `./internal/api/` tetap hijau.
- [ ] **Langkah 5: commit** `Expire each invitation by a job at its expires_at and run the jobs runner in casa serve (06 §3.6; 09.2 §5.1 note X3; 09.3 §9)`.

### Tugas 3: Modul `read` memegang proyeksi layar dan `AccountEmail`

Refactor tanpa perubahan perilaku (ES §7: terpisah dari perubahan perilaku). Bukti utamanya: uji 1.6 lulus tanpa disunting.

**Berkas:**
- **Buat:** `server/internal/read/read.go`, `server/internal/read/queries.sql`, hasil sqlc `server/internal/read/readdb/*.go`
- **Ubah:**
  - `server/internal/household/members.go`, `invitation.go`, `queries.sql`: `Memberships`, `membershipsOf`, `OpenInvitation` dan kueri yang hanya dipakai mereka dihapus
  - `server/internal/identity/identity.go`, `queries.sql`: `AccountEmail` dan `GetAccountEmail` dihapus
  - `server/internal/api/household.go`: `ListMemberships`, `GetOpenInvitation`; `RegisterHouseholdService` menerima `*read.Service`
  - `server/cmd/casa/serve.go`: `read.New`
  - `server/sqlc.yaml`; `server/internal/architecture/imports_test.go`: `"internal/read": {"internal/database", "internal/read/readdb"}`, `"internal/read/readdb": {}`, `internal/api` dan `cmd/casa` + `internal/read` (09.2 §1 butir 10)

**Antarmuka:**
- **Consumes:** `database.Read`.
- **Produces:**

```go
// Service is the one read path (09.2 §1 rule 2; SEC-06).
type Service struct{ /* pool, now */ }
func New(pool *pgxpool.Pool, now func() time.Time) *Service

// Reader is the member a read is for, its membership already checked by the caller (09.2 §3 layer 1).
type Reader struct{ HouseholdID, AccountID, MemberID pgtype.UUID }

type Member struct {
	MemberID    pgtype.UUID
	DisplayName string
}
type Membership struct {
	HouseholdID   pgtype.UUID
	HouseholdName string
	MemberID      pgtype.UUID
	Members       []Member
}
// Memberships is household.Memberships moved unchanged (bounds of plan 1.6: 20 and 20).
func (s *Service) Memberships(ctx context.Context, account pgtype.UUID) ([]Membership, error)

type OpenInvitation struct {
	ID                 pgtype.UUID
	State              string
	ExpiresAt          time.Time
	Inviter            pgtype.UUID
	InviteeDisplayName string      // only for the inviter of an accepted invitation
	InviteeAccount     pgtype.UUID // idem
}
// OpenInvitation is household.OpenInvitation moved unchanged; the membership
// is checked again in the read transaction (SEC-151); ErrNotFound without it.
func (s *Service) OpenInvitation(ctx context.Context, r Reader) (inv OpenInvitation, found bool, err error)

// AccountEmail is identity.AccountEmail moved unchanged (plan 1.6 Q10): read
// in the account's own scope, email only; its one caller is GetOpenInvitation.
func (s *Service) AccountEmail(ctx context.Context, account pgtype.UUID) (string, error)

var ErrNotFound = errors.New("not found")
```

- [ ] **Langkah 1: tulis uji yang gagal:** `server/internal/read/read_test.go`, `TestRead`:
  - `memberships moved to read keep their bounds` (21 rumah tangga dan 21 anggota sebagai admin → 20 dan 20);
  - `the open invitation shows the invitee only to the inviter of an accepted invitation`;
  - `an invitation past its expiry is not open and nothing is written`;
  - `a reader without membership in the household is not found`.
- [ ] **Langkah 2:** `go test -race -p 1 ./internal/read/` → GAGAL (paket tidak ada).
- [ ] **Langkah 3:** pindahkan kode dan kueri; ganti pemanggil di handler.
- [ ] **Langkah 4:** `go test -race -p 1 ./internal/read/ ./internal/household/ ./internal/identity/ ./internal/api/ ./internal/architecture/` dan `make sqlc-generated` → LULUS, dengan `git diff --stat main -- server/internal/api/*_test.go` kosong (uji 1.6 tidak disunting).
- [ ] **Langkah 5: commit** `Move the screen projections and AccountEmail behind read, the one read path (09.2 §1 rule 2; SEC-06; plan 1.6 Q10)`.

### Tugas 4: Umpan per pembaca dan fan-out di transaksi perintah

**Berkas:**
- **Buat:**
  - `server/internal/schema/00010_reader_feed.sql`
  - `server/internal/sync/feed.go`, `server/internal/sync/queries.sql`, hasil sqlc `server/internal/sync/syncdb/*.go`
  - `server/internal/sync/feed_test.go`
- **Ubah:**
  - `server/internal/read/read.go`, `queries.sql`: `ObjectKind`, `ObjectRef`, `Readers`
  - `server/internal/household/events.go`: `RecordEvent` memanggil `sync.Changed`
  - `server/internal/schema/catalog_test.go`: baris `reader_feed`, `reader_feed_counter`
  - `server/sqlc.yaml`; `imports_test.go`: `"internal/sync": {"internal/database", "internal/read", "internal/sync/syncdb"}`, `"internal/sync/syncdb": {}`, `internal/household` + `internal/sync` (09.2 §1 butir 10)

**Migrasi:** kedua tabel seperti 09.2 §5.1, ditambah:
- `household_id` dan `member_id` `NOT NULL`;
- `FOREIGN KEY (household_id, member_id) REFERENCES membership`;
- `object_kind CHECK IN ('household','member','invitation')`.

RLS lapis 4 ketat pada keduanya; kebijakan rumah tangga `reader_feed` dipasang `TO app_rw` sejak dibuat (09.2 §3 catatan 14 November 2026, putusan Q1 (d)), dan gerbang katalog memeriksa perannya. Hak `app_rw`: `reader_feed` `SELECT, INSERT` (tambah-saja); `reader_feed_counter` `SELECT, INSERT, UPDATE (last_pos)`. `job_runner` dan `lifecycle_rw` tanpa hak.

**Antarmuka:**
- **Consumes:** `household.RecordEvent` (1.6).
- **Produces:**

```go
// package read
type ObjectKind string
const (
	KindHousehold  ObjectKind = "household"
	KindMember     ObjectKind = "member"
	KindInvitation ObjectKind = "invitation"
)
type ObjectRef struct {
	Kind ObjectKind
	ID   pgtype.UUID
}
// Readers are the members who can read an object of household after the
// change: today every member, for every kind (plan 1.7 decision 7).
func Readers(ctx context.Context, tx pgx.Tx, household pgtype.UUID, kind ObjectKind) ([]pgtype.UUID, error)

// package sync
// Changed writes one upsert row per reader of the object in tx, the command's
// transaction (09.2 §6 step 6; ADR-017 item 3). A household object also
// writes every member object (plan 1.7 decision 8).
func Changed(ctx context.Context, tx pgx.Tx, household pgtype.UUID, kind read.ObjectKind, id pgtype.UUID) error
```

`RecordEvent` memetakan `object_kind` Peristiwa `household` → `read.KindHousehold`, `invitation` → `read.KindInvitation`.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestReaderFeed(t *testing.T) {
	t.Run("creating a household writes the household and the creator to the creator's feed", …) // pos 1, 2
	t.Run("every invitation transition writes one row per member", …)
	// table: F01, F02, F05 withdraw, CloseAsWithdrawn, F03, reject, expiry job; each → one invitation row per member
	t.Run("F04 writes the household, every member and the invitation to the members' feeds, the new member included", …)
	t.Run("issuing a link writes no feed row", …)
	t.Run("a refused command writes no feed row", …) // HOUSEHOLD_FULL, STATE_CHANGED
	t.Run("positions have no gap per reader under concurrent commands", …)
	// two goroutines × 10 Changed calls in separate Writes for one household → each member's pos = 1..20
}

// in schema: reader feed tables are household-scoped and append-only
t.Run("reader_feed is household-scoped and app_rw cannot update or delete it", …) // 42501
t.Run("a feed row of another household cannot be written", …)                   // 42501
```

- [ ] **Langkah 2:** `go test -race -p 1 ./internal/sync/ ./internal/schema/` → GAGAL.
- [ ] **Langkah 3:** tulis migrasi, kueri (`BumpCounter`: upsert penghitung, `RETURNING last_pos`; `InsertFeedRows` dengan `unnest`), `Readers`, `Changed`, dan panggilannya di `RecordEvent`.
- [ ] **Langkah 4:** jalankan perintah langkah 2, lalu seluruh `./internal/household/ ./internal/api/ ./internal/architecture/` dan `make sqlc-generated` → LULUS. `EXPLAIN` `Readers`, `BumpCounter`, `InsertFeedRows` untuk PR.
- [ ] **Langkah 5: commit** `Write the per-reader feed in every command transaction through the one event writer (ADR-017 item 3; 09.2 §6 step 6, §5.1; INV-A01)`.

### Tugas 5: Kursor, proyeksi sinkron, umpan, dan potret

**Berkas:**
- **Buat:** `server/internal/sync/cursor.go`, `cursor_test.go`, `pull.go`, `pull_test.go`
- **Ubah:** `server/internal/read/read.go`, `queries.sql` (`Objects`, `SnapshotPage`, proyeksi), `server/internal/sync/queries.sql` (`LastPos`, `FeedAfter`)

**Antarmuka:**
- **Consumes:** `read.ObjectRef`, `read.Reader` (Tugas 3, 4); `syncdb` (Tugas 4).
- **Produces:**

```go
// package read
type Household struct {
	ID       pgtype.UUID
	Name     string
	TimeZone string
	Currency string
}
type Invitation struct {
	ID        pgtype.UUID
	State     string // open or accepted
	ExpiresAt time.Time
	Inviter   pgtype.UUID
}
// Object is one projection; exactly one field is set.
type Object struct {
	Household  *Household
	Member     *Member
	Invitation *Invitation
}
// Objects projects refs for r in tx; a ref that r cannot read or that is not
// live (decision 7) is absent from the map.
func (s *Service) Objects(ctx context.Context, tx pgx.Tx, r Reader, refs []ObjectRef) (map[ObjectRef]Object, error)

// SnapshotMark is a position in the snapshot order (kind, then id); its zero value is the start.
type SnapshotMark struct {
	Kind  ObjectKind
	After pgtype.UUID
}
type Projected struct {
	Ref    ObjectRef
	Object Object
}
// SnapshotPage returns up to limit live objects r can read after mark, in order (ADR-017 item 7).
func (s *Service) SnapshotPage(ctx context.Context, tx pgx.Tx, r Reader, mark SnapshotMark, limit int) (objects []Projected, next SnapshotMark, done bool, err error)

// package sync
type CursorKey struct{ /* 32-byte key */ }
func NewCursorKey(key []byte) (CursorKey, error) // len(key) != 32 → error
type Cursor struct {
	Household, Member pgtype.UUID
	Pos               int64
	Snapshot          *read.SnapshotMark // nil in feed mode
}
func (k CursorKey) Encode(c Cursor) string
// Decode refuses with ErrCursorRejected anything but a cursor k made for this household and member (decision 9).
func (k CursorKey) Decode(s string, household, member pgtype.UUID) (Cursor, error)
var ErrCursorRejected = errors.New("cursor rejected")

type ChangeKind int
const (
	Upsert ChangeKind = iota + 1
	Purge
)
type Change struct {
	Kind   ChangeKind
	Ref    read.ObjectRef
	Object read.Object // zero for Purge
}
type Page struct {
	Changes        []Change
	Next           string
	HasMore        bool
	ResyncRequired bool
}
type Service struct{ /* pool, read, key */ }
func New(pool *pgxpool.Pool, r *read.Service, key CursorKey) *Service
// Pull is one Sync page for r (09.3 §3; decisions 9 to 11), in one REPEATABLE
// READ READ ONLY transaction; read.ErrNotFound when r is no longer a member.
func (s *Service) Pull(ctx context.Context, r read.Reader, cursor string, limit int) (Page, error)
```

**Algoritma `Pull`** (selain keputusan 10 dan 11): mode umpan membaca `FeedAfter(P, limit+1)` dan menerapkan ketiga syarat resync keputusan 11 sebelum memproyeksikan. Setiap baris menjadi `Upsert` bila `Objects` memuatnya, selain itu `Purge`; baris `change = 'purge'` selalu `Purge`. `Next` = kursor umpan pada `pos` baris terakhir yang dikembalikan (atau `P` bila tidak ada), dan `HasMore = len(rows) > limit`. Mode potret: `HasMore = !done || last_pos > P0`.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestCursor(t *testing.T) { // no database
	t.Run("a feed and a snapshot cursor round-trip", …)
	t.Run("a cursor of another member or another household is rejected", …)
	t.Run("a cursor with any one bit flipped is rejected", …) // every byte position of the decoded value
	t.Run("an empty body, a truncated cursor, bad base64 and an unknown version, mode or kind are rejected", …)
	t.Run("a cursor made with another key is rejected", …)
	t.Run("a key that is not 32 bytes is refused", …)
}

func TestPull(t *testing.T) {
	t.Run("an empty cursor returns the household, its members and its live invitation, then the feed position", …)
	t.Run("the snapshot pages by kind and id with the limit", …) // limit 1 → 4 pages, has_more until the last
	t.Run("a change between snapshot pages arrives through the feed", …)
	// page 1 (limit 1) taken; then F01 through household.Service; then the rest of the snapshot, then the
	// feed from P0 → the reader holds the invitation exactly as read.Objects projects it, once F03 a purge
	t.Run("a feed cursor returns each change once in order and the next cursor", …)
	t.Run("an invitation no longer live is a purge", …)       // F03 → PURGE invitation
	t.Run("an invitation past its expiry but not yet written is a purge", …)
	t.Run("a cursor behind the pruned feed asks for a resync", …) // admin deletes the reader's first rows
	t.Run("a cursor ahead of the feed asks for a resync", …)      // Encode with Pos = last_pos + 1
	t.Run("a gap inside the returned rows asks for a resync", …)  // admin deletes the reader's row P+2 only
	t.Run("a page that is not full and does not end at last_pos asks for a resync", …) // admin deletes the reader's last rows
	t.Run("projections carry no invitee name or email", …)        // read.Invitation has neither; accepted invitation
}
```

- [ ] **Langkah 2:** `go test -race -p 1 ./internal/sync/ ./internal/read/` → GAGAL.
- [ ] **Langkah 3:** tulis codec menurut tata letak keputusan 9; tulis `Objects`, `SnapshotPage`, `canRead`, `Pull`.
- [ ] **Langkah 4:** jalankan perintah langkah 2 dan `make sqlc-generated` → LULUS. `EXPLAIN` `FeedAfter`, `LastPos`, dan kueri halaman potret per jenis untuk PR.
- [ ] **Langkah 5: commit** `Pull a reader's feed and initial snapshot through the read path, with a MAC-bound cursor and resync on a pruned feed (ADR-017 items 3, 6, 7, 8; 09.3 §3; SEC-05, SEC-06)`.

### Tugas 6: Kontrak dan RPC `Sync`

**Berkas:**
- **Buat:** `proto/casa/v1/sync.proto`, `server/internal/api/sync.go`, `server/internal/api/sync_test.go`
- **Ubah:**
  - `proto/casa/v1/household.proto`: pesan `Household`, `Invitation`
  - hasil generate Go dan TypeScript
  - `server/internal/api/selector_test.go`: baris `Sync`
  - `server/cmd/casa/serve.go`, `serve_test.go`: kredensial `sync-cursor-key`, `sync.New`, `RegisterSyncService`
  - `imports_test.go`: `internal/api` dan `cmd/casa` + `internal/sync` (09.2 §1 butir 10)
  - `README.md` akar: `SyncService`, kredensial

**Kontrak:**

```proto
service SyncService {
  // The reader's feed or initial snapshot (09.3 §3; ADR-017, ADR-018). A read:
  // no CommandOptions, never held by NOTICE_UNREAD. NotFound for a household
  // the caller is not a member of and for any cursor the server did not give
  // this member; a client then drops the cursor and tries once with an empty
  // one (plan 1.7 decision 9).
  rpc Sync(SyncRequest) returns (SyncResponse);
}
enum SyncPurpose { SYNC_PURPOSE_UNSPECIFIED = 0; SYNC_PURPOSE_FOREGROUND = 1; SYNC_PURPOSE_BACKGROUND = 2; }
message SyncRequest {
  RequestContext request_context = 1 [(buf.validate.field).required = true];
  string household_id = 2 [(buf.validate.field).string.uuid = true];
  string cursor = 3 [(buf.validate.field).string.max_len = 256];
  SyncPurpose purpose = 4 [(buf.validate.field).enum = {defined_only: true, not_in: [0]}];
  uint32 limit = 5 [(buf.validate.field).uint32.lte = 500]; // 0 = 100
}
enum ChangeKind { CHANGE_KIND_UNSPECIFIED = 0; CHANGE_KIND_UPSERT = 1; CHANGE_KIND_PURGE = 2; }
enum ObjectKind { OBJECT_KIND_UNSPECIFIED = 0; OBJECT_KIND_HOUSEHOLD = 1; OBJECT_KIND_MEMBER = 2; OBJECT_KIND_INVITATION = 3; }
message Change {
  ChangeKind kind = 1;
  ObjectKind object_kind = 2;
  string object_id = 3;
  oneof object { Household household = 4; Member member = 5; Invitation invitation = 6; } // empty on PURGE
}
message SyncResponse {
  repeated Change changes = 1;
  string next_cursor = 2;
  bool has_more = 3;
  bool resync_required = 4;          // drop the cache, never the outbox; start again with an empty cursor
  bool operator_notice_unacknowledged = 5; // SEC-34: show LY-05 from GetOperatorNotice
}
// household.proto
message Household { string household_id = 1; string name = 2; string time_zone = 3; string currency = 4; }
message Invitation { string invitation_id = 1; InvitationState state = 2; google.protobuf.Timestamp expires_at = 3; string inviter_member_id = 4; }
```

**Antarmuka:**
- **Consumes:** `household.Member` (selektor 1.6), `api.NoticeHeld` (1.6), `sync.Pull`, `sync.NewCursorKey` (Tugas 5).
- **Produces:** `func RegisterSyncService(server *connect.Server, h *household.Service, s *sync.Service, noticeHeld func(context.Context, pgtype.UUID) (bool, error))`. Handler: `h.Member` (tanpa penulisan gugur malas: `Sync` adalah bacaan seperti `GetOpenInvitation`) → `noticeHeld` → `s.Pull`. `ErrCursorRejected` dan `read.ErrNotFound` → galat yang sama dengan `household.ErrNotFound`.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestSync(t *testing.T) {
	t.Run("SEC-05 a cursor of another member, of another household, or changed in one character is not found like a household that does not exist", …)
	// compare code, message and details with a Sync on a random household id
	t.Run("INV-A48 (c) an invitee before F04 reads nothing through sync", …)
	// accepted invitee with the inviter's household id → identical to a household that does not exist
	t.Run("both members reach the same household state through sync", …)
	// after F04: creator (feed) and new member (snapshot) hold equal objects; equal to ListMemberships and GetOpenInvitation
	t.Run("a new notice version shows in the next sync and acknowledging clears it", …)
	// fstest notices as in 1.6: operator_notice_unacknowledged true, then false after AcknowledgeOperatorNotice
	t.Run("sync is a read: never held by NOTICE_UNREAD and needs no idempotency key", …)
	t.Run("the invitee's email and display name never travel in sync", …)
	// SyncResponse descriptor (recursively) has no field matching (?i)email|invitee; accepted invitation → id, state, expires_at, inviter only
	t.Run("a pruned feed answers resync_required with no changes and no cursor", …)
	t.Run("a sync without purpose is refused", …) // InvalidArgument
	t.Run("the job and the lazy path give the same sync answer", …)
	// Sync half of Task 2's comparison (10.1 §10 note 1.7 Q3 condition 1): the inviter's Sync after each
	// expiry gives the same PURGE change apart from pos, object id and cursor
```

Ditambah: baris `{"Sync", false, …}` di `selectorCases`, sehingga `every RPC with a household selector has a selectorCases row` dan ketiga subuji pemilih 1.6 berlaku untuk `Sync`; dan di `serve_test.go`: `serve refuses a cursor key that is not 64 hex characters`.

- [ ] **Langkah 2:** `make proto-generate` lalu `go test -race -p 1 -run 'TestSync|TestHouseholdSelector' ./internal/api/ && go test -race -p 1 ./cmd/casa/` → GAGAL.
- [ ] **Langkah 3:** tulis handler dan pemasangan di `serve`.
- [ ] **Langkah 4:** jalankan perintah langkah 2, `make check-contract`, dan `go test -race -p 1 ./internal/api/ ./internal/architecture/` → LULUS.
- [ ] **Langkah 5: commit** `Add the Sync RPC: selector check, notice state, feed and snapshot pages through read (09.3 §3; ADR-017, ADR-018 item 2; SEC-05; INV-A48 (c); FR-FAM-019 AC-3, AC-10 sync part)`.

### Tugas 7: Pembangkit uji isolasi per kueri

**Berkas:**
- **Buat:** `server/internal/isolation/isolation_test.go`, `server/internal/isolation/testdata/fixture.sql`
- **Ubah:** `imports_test.go`: `"internal/isolation": {}` (paket berisi uji saja, seperti `internal/architecture`)

**Antarmuka:**
- **Consumes:** semua `*Queries` sqlc: `commanddb`, `householddb`, `identitydb`, `jobsdb`, `readdb`, `syncdb`; `databasetest.Migrated`.
- **Produces:** tidak ada untuk tugas lain.

**Algoritma** (keputusan 15):
1. **Temukan kueri.** `go/parser` membaca `server/internal/*/*db/*.sql.go`. Setiap metode `*Queries` adalah satu kueri; parameter skalar dan medan struct parameter diratakan menjadi daftar (nama, tipe).
2. **Fixture.** `fixture.sql` (admin, UUID tetap) membuat rumah tangga X dan Y lengkap: setiap tabel di skema `public` punya baris X dan baris Y, dan setiap akun, token, Peristiwa, umpan, job, dan rekaman idempotensi terikat ke salah satunya.
3. **Nilai.** `isolationValues` memetakan setiap nama parameter ke salah satu dari dua bentuk:
   - pasangan pengenal `(x, y)`, untuk nama seperti `household_id`, `id`, `member_id`, `account_id`, `token_hash`, `email`;
   - nilai tetap, untuk nama yang bukan pengenal, seperti `limit`, `now`, `state`, `version`.

   Nama tanpa entri menggagalkan uji.
4. **Jalankan.** Untuk setiap kueri dan setiap parameter pengenal *p*: parameter lain diisi nilai Y (panggilan sah bagi Y), lalu dua run:
   - run X: *p* = nilai X;
   - run R: *p* = nilai acak segar bertipe sama.

   Setiap run berjalan di koneksi admin: `BEGIN`; `SET LOCAL ROLE app_rw`; cakupan `Scope{HouseholdID: Y, AccountID: A_y}` (pernyataan `set_config` yang sama dengan `database`); `SAVEPOINT`; panggil kueri. Bila galat: `ROLLBACK TO SAVEPOINT`. Lalu `RESET ROLE`; hitung digest semua baris X; `ROLLBACK`.
5. **Hasil yang dicatat per run:** kode SQLSTATE atau kosong, jumlah baris hasil atau terpengaruh, apakah hasil memuat nilai X mana pun (pindai reflektif atas UUID, `[]byte`, dan string fixture X), dan digest X.
6. **Lulus** bila run X dan run R punya kode dan jumlah yang sama, run X tidak memuat nilai X, dan digest X sama dengan digest awal.

**Pengecualian** hanya lewat `isolationExceptions []struct{ query, param, reason string }`. Setiap alasan mengutip sumber. Calon yang diketahui:
- `membership_one_household` dan `invitation_invitee` pada parameter akun (09.2 §5.1 keputusan 20: kuncinya akun pemanggil sendiri, dan handler hanya memberi akun sesi);
- email akun (09.2 §3: satu-satunya keunikan lintas rumah tangga; jawaban API seragam, SEC-30).

Pengecualian yang tidak lagi terpicu menggagalkan uji.

- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestEveryQueryIsIsolated(t *testing.T) {
	t.Run("every table holds rows of both households in the fixture", …)
	// information_schema tables of public, minus goose_db_version and currency (product table, 09.2 §3)
	t.Run("every query parameter has an isolation value", …)
	// one subtest per query, named "<package>.<Query>"
	t.Run("SEC-02 another household's values behave as values that never existed", …)
	t.Run("every declared exception still applies", …)
}
```

- [ ] **Langkah 2:** `go test -race -p 1 ./internal/isolation/` → GAGAL (paket dan fixture tidak ada).
- [ ] **Langkah 3:** tulis fixture, penemu kueri, peta nilai, dan pengecualian. Setiap temuan nyata (bukan salah fixture) dihentikan dan dilaporkan ke lead sebelum diperbaiki di tugas sendiri.
- [ ] **Langkah 4:** jalankan perintah langkah 2 → LULUS; catat jumlah kueri dan waktu uji untuk PR.
- [ ] **Langkah 5: commit** `Generate a cross-household isolation test for every sqlc query (09.2 §3 layer 6; SEC-02; ES §4)`.

### Tugas 8: Rangka uji properti berstate

**Berkas:**
- **Buat:** `server/internal/api/property_test.go` (mesin, pembangkit perintah), `server/internal/api/invariant_test.go` (pemeriksa), `.github/workflows/property-weekly.yml`
- **Ubah:**
  - `server/go.mod`, `server/go.sum`: `pgregory.net/rapid v1.3.0`
  - `Makefile`: `server-test` mengekspor `RAPID_CHECKS ?= 20`, `RAPID_SEED ?= 1` dan memberi `-rapid.nofailfile`; target `server-property-long` (`RAPID_CHECKS=1000`, benih acak yang dicetak)
  - `.gitignore`: `server/**/testdata/rapid/`

**Antarmuka:**
- **Consumes:** harness `api_test` 1.6 (`startHousehold`, `signUpMember`, `h.clock`); `jobs.Runner.RunDue` (Tugas 1); RPC `Sync` (Tugas 6).
- **Produces** (dipakai langkah 2 untuk INV):

```go
// commandBuilders has one builder per write command of the contract: every
// method with the casa.v1.command option (14 §4.3 "Pembangkit perintah").
var commandBuilders map[string]func(m *householdMachine, t *rapid.T) (actor int, request proto.Message)

// householdMachine: 2 to 5 accounts, the command log (actor index, procedure,
// request, answer), each member's sync model (cursor and objects), the clock.
type householdMachine struct{ /* … */ }

// invariantCheck runs after every step over the whole database (14 §4.3 "Pemeriksa").
type invariantCheck func(t *rapid.T, m *householdMachine)
```

**Aksi mesin:**
- satu aksi per prosedur di `commandBuilders`;
- `resend`: kirim ulang perintah acak dari log dengan kunci yang sama;
- `advanceClock`: 1 menit sampai 8 hari;
- `runDueJobs`;
- `sync`: satu anggota acak menarik umpan sampai `has_more` false.

**Pemeriksa:**
- **INV-A01:** per objek, Peristiwa `seq` 0..`version` tanpa celah; Peristiwa yang sudah terlihat tidak berubah (hash prefiks).
- **INV-A48:**
  - (a) baris keanggotaan tidak pernah hilang atau berganti peran dan izin;
  - (b) setiap keanggotaan selain pembuat punya undangan `confirmed` dengan Peristiwa F04 berpelaku pengundang, dan tidak ada Peristiwa `system` menuju `confirmed`;
  - (c) setiap akun yang bukan anggota mendapat jawaban "tidak ada" dari `Sync` dan `GetOpenInvitation`, dan tidak melihat rumah tangga itu di `ListMemberships`.
- **Kiriman ulang:** jawaban sama dengan yang pertama, dan jumlah baris semua tabel tidak berubah.
- **Sinkron = bacaan:** untuk setiap anggota, (model inkremental) = (potret dari kursor kosong) = (proyeksi `ListMemberships` + `GetOpenInvitation`), sesudah dinormalkan.
- **INV-A13 / FR-FAM-001/AC-4:** sesudah rangkaian selesai, log diputar ulang pada akun baru dengan pelaku 0 dan 1 ditukar. Jawaban setiap perintah, dan pandangan setiap orang (`ListMemberships`, potret `Sync`, `GetOpenInvitation`), harus sama sesudah pengenal dipetakan menurut urutan lahirnya.

- [ ] **Langkah 0: catatan ES §8 rapid** (dicek 14 November 2026, disalin ke deskripsi PR C):
  - **Versi:** v1.3.0, rilis stabil terbaru. Waktu tag di proxy 2026-03-30T15:56:16Z ([proxy.golang.org/pgregory.net/rapid/@latest](https://proxy.golang.org/pgregory.net/rapid/@latest)); daftar versi v1.0.0 sampai v1.3.0, tanpa pra-rilis.
  - **Catatan rilis** ([github.com/flyingmutant/rapid/releases](https://github.com/flyingmutant/rapid/releases)): v1.3.0 menambah override `Make[T]` dan `t.Output`, tanpa perubahan yang merusak.
  - **Modul:** `go 1.23`, tanpa `require` (nol dependensi).
  - **Lisensi:** MPL-2.0. Hanya diimpor berkas `_test.go`, jadi tidak masuk biner `casa`.
  - **Kerentanan:** OSV kosong untuk `pgregory.net/rapid`.
  - **Isu terbuka:** tidak ada yang kritis. #63 (panik di `checkFailFile` bila kode berubah) tidak berlaku karena berkas gagal tidak dipakai; #97 (penyusutan nilai besar) tidak berlaku karena rangkaian perintah kecil.
  - **Alasan:** keputusan 16.
  - **ADR-023:** tidak berlaku (≥ 1.0).

  Langkah ini mengecek ulang versi terbaru saat dipasang. Bila sudah lebih baru, versi itu yang dipakai dan catatannya diperbarui.
- [ ] **Langkah 1: tulis uji yang gagal.**

```go
func TestHouseholdProperties(t *testing.T) {
	t.Run("every write command of the contract has a builder", …) // from the service descriptors
	t.Run("INV-A01 every invitation and household transition writes one event and old events never change", …)
	t.Run("INV-A48 (a) to (c) hold after every random household command", …)
	t.Run("INV-A13 swapping the two full adults over a random sequence swaps every answer and view", …)
	t.Run("a resent command returns its first answer and writes nothing", …)          // ADR-005
	t.Run("sync after every command equals the direct read for every member", …)     // SEC-06; 14 §4.7
}
```

Setiap subuji properti adalah `rapid.Check(t, …)` atas mesin yang sama dengan satu pemeriksa.
- [ ] **Langkah 2:** `go get pgregory.net/rapid@v1.3.0`; `go test -race -p 1 -run TestHouseholdProperties ./internal/api/` → GAGAL (`commandBuilders` kosong).
- [ ] **Langkah 3:** tulis mesin, pembangun untuk setiap perintah `HouseholdService` yang ada, pemeriksa, target Makefile, dan alur mingguan (cron Senin 03.00 UTC, `make db-up server-property-long`; kegagalan diperlakukan sama dengan kegagalan per perubahan, 14 §4.3).
- [ ] **Langkah 4:** `make server-test` → LULUS. Waktu `TestHouseholdProperties` dengan `RAPID_CHECKS=20` dicatat untuk PR (anggaran CI 14 §4.3). `make server-vuln` → LULUS dengan rapid.
- [ ] **Langkah 5: commit** `Add the stateful property test frame with INV-A01, INV-A13, INV-A48 and sync-equals-read checks (14 §4.3, §4.7; 09.4 §4; plan 1.6 Q17)`.

### Tugas 9: Pembangkit vektor emas panjang nama

**Berkas:**
- **Buat:** `golden/input-shape/name-length.json` (dibangkitkan), `server/internal/household/golden_test.go`
- **Ubah:** `Makefile`: target `golden-generate` (`CASA_GOLDEN_UPDATE=1 go test -run TestGoldenNameLength ./internal/household/`)

**Bentuk berkas** (JSON, satu baris per kasus, deterministik):

```json
{"source": "household.CleanName (FAM.md FR-FAM-001 note of 14 November 2026)",
 "cases": [{"input": "…", "max": 60, "rule": ""}, …]}
```

`rule` adalah keluaran `CleanName`: `""`, `"name.visible_length"`, atau `"name.characters"`. Pengguna klien kelak menegaskan bantuan panjangnya pada baris `""` (sah) dan `name.visible_length` (tidak sah) saja; `name.characters` aturan server (ES §2: server memeriksa ulang (c)).

**Kasus pilihan tangan,** masing-masing untuk `max` 60 dan 40:
- `""`, `" "`, `" "`, `"　a　"`;
- `"a"`; tepat `max` dan `max + 1` huruf;
- `"é"` dan `max` kali `"é"` (NFC);
- `"👨‍👩‍👧"`;
- `"Ayu️"`, `"​"`, `"﻿"`, `"a‮b"`.

**Kasus acak:** 200 string, panjang 0 sampai 70, dari kumpulan (ASCII, tanda gabung, emoji, ZWJ, spasi Unicode), dengan `rand.New(rand.NewPCG(1, 7))`.

- [ ] **Langkah 1: tulis uji yang gagal:** `TestGoldenNameLength`:
  - `the vector file equals what CleanName produces`: dibangkitkan di memori lalu dibandingkan byte demi byte dengan berkas; dengan `CASA_GOLDEN_UPDATE=1` berkas ditulis ulang;
  - `the file holds every hand-picked case`.
- [ ] **Langkah 2:** `go test -race -p 1 -run TestGoldenNameLength ./internal/household/` → GAGAL (berkas tidak ada).
- [ ] **Langkah 3:** tulis pembangkit; `make golden-generate`.
- [ ] **Langkah 4:** jalankan perintah langkah 2 tanpa `CASA_GOLDEN_UPDATE` → LULUS.
- [ ] **Langkah 5: commit** `Generate the name length golden vectors from the server's reference (ES §2 (c); 14 §4.5; ADR-013 item 3)`.

### Tugas 10: Verifikasi, jalan lokal, QG-05, bukti merah, PR

Langkah 2 dan 3 dijalankan per PR: A di akhir Tugas 2, B di akhir Tugas 6, C di sini.

- [ ] **Langkah 1:** `make check` → LULUS, termasuk `server-test-old-binary` (biner `main` di atas migrasi baru PR itu).
- [ ] **Langkah 2: jalan lokal sungguhan** (ES §9 butir 5), dengan `make casa db-up`, `.cache/casa migrate`, dan kredensial baru di `CREDENTIALS_DIRECTORY` uji. Keluaran dilampirkan dengan token, cookie, kursor, dan email disensor.

  | Bagian | Versi basis data | Alur |
  |---|---|---|
  | A | `database at version 9` | daftar → rumah tangga → `CreateInvitation`; sebagai admin di basis data pengembang (data uji) `UPDATE invitation SET expires_at = now()` dan `UPDATE job SET run_at = now()`; dalam ±1 detik undangan `expired`, job hilang; log tanpa isi |
  | B | `database at version 10` | `Sync` kursor kosong (`limit` 1, empat halaman) → `CreateInvitation` → `Sync` dengan kursor → `UPSERT` undangan → `CancelInvitation` → `PURGE`; kursor yang diubah → `NotFound` |
  | C | `database at version 10` | `make server-test` dengan waktu `TestHouseholdProperties` dan `TestEveryQueryIsIsolated`; `make golden-generate` tanpa beda |

- [ ] **Langkah 3: bukti merah**, satu per satu, di PR draf sekali pakai per PR:

| PR | Pelanggaran buangan | Yang harus merah |
|---|---|---|
| A | `ClaimDue` tanpa `locked_until` | `two runners at once claim each job once` |
| A | Gagal tanpa jeda (`locked_until` tidak diisi) | `a failed job waits 10 s, then 20 s, and stops after 10 attempts with its code` |
| A | `attempts < 10` dibuang dari `ClaimDue` | yang sama |
| A | Log "job stopped" memuat `payload` | `a stopped job is logged with its kind and id and no payload` |
| A | `Enqueue` lewat pool, bukan `tx` | `a job enqueued in a rolled-back transaction does not exist` |
| A | `GRANT DELETE ON job TO job_runner` | `job_runner reads every job and updates only locked_until, attempts and last_error` |
| A | Kebijakan `job_runner` tanpa `TO job_runner` | `each policy of job names its role` dan `a household transaction sees only its own jobs, …` |
| A | `job_unique` tanpa `household_id` | `the same kind and unique key in two households are two jobs, …` |
| A | Penangan menganggap "belum jatuh tempo" selesai | `a live invitation the clock finds not yet due is retried, not done` |
| A | F01 tanpa `Enqueue` | `F01 schedules the expiry at expires_at` |
| A | Penangan gugur menulis `closed_at = now` | `at expires_at the job writes the expiry at expires_at` |
| B | `read.OpenInvitation` mengisi penerima bagi setiap anggota | uji 1.6 `the invitee's email is shown to nobody else and stored nowhere` |
| B | `RecordEvent` tanpa `sync.Changed` | `every invitation transition writes one row per member` |
| B | Peristiwa rumah tangga hanya menulis objek rumah tangga | `F04 writes the household, every member and the invitation …` |
| B | `pos` = `max(pos) + 1` tanpa penghitung | `positions have no gap per reader under concurrent commands` |
| B | `Decode` tanpa perbandingan anggota | `a cursor of another member or another household is rejected` dan `SEC-05 a cursor of another member …` |
| B | Kursor akhir potret = posisi halaman terakhir, bukan `P0` | `a change between snapshot pages arrives through the feed` |
| B | Pemeriksaan resync dibuang | `a cursor behind the pruned feed asks for a resync` |
| B | Resync hanya memeriksa baris pertama | `a gap inside the returned rows asks for a resync` |
| B | Kebijakan rumah tangga `reader_feed` tanpa `TO app_rw` | gerbang katalog `reader_feed` |
| B | Undangan tertutup diproyeksikan sebagai `UPSERT` | `an invitation no longer live is a purge` |
| B | Handler `Sync` tanpa `h.Member` | `INV-A48 (c) an invitee before F04 reads nothing through sync` |
| B | Handler `Sync` tanpa `noticeHeld` | `a new notice version shows in the next sync and acknowledging clears it` |
| C | Kebijakan `reader_feed` `USING (true)` | `SEC-02 another household's values behave as values that never existed` (subuji `syncdb.*`) |
| C | Kueri baru dengan parameter tanpa nilai isolasi | `every query parameter has an isolation value` |
| C | Fixture tanpa baris `job` | `every table holds rows of both households in the fixture` |
| C | `Confirm` tanpa Peristiwa `member_joined` | `INV-A01 …` |
| C | `CancelInvitation` dengan `RecordEvent` yang melewati `sync.Changed` | `sync after every command equals the direct read for every member` |
| C | F01 ditolak bila pelaku bukan pembuat rumah tangga | `INV-A13 …` |
| C | `CleanName` menghitung tanda gabung | `the vector file equals what CleanName produces` |

- [ ] **Langkah 4: deskripsi PR** memuat:
  - kueri sqlc baru dengan indeks dan `EXPLAIN`;
  - tautan run merah;
  - keluaran jalan lokal;
  - catatan ES §8 rapid (PR C) dan waktu uji properti dan isolasi;
  - putusan yang dipakai.
- [ ] **Langkah 5:** `code-reviewer` meninjau dua lintasan; temuan diperbaiki; vonis siap merge; lead merge (merge commit).

## Selesai bila

- PR A, B, dan C:
  - CI hijau (`contract`, `server`, `web`);
  - setiap baris bukti merah terbukti merah;
  - `code-reviewer` tanpa temuan terbuka;
  - lead sudah merge.
- Di `main`:
  - undangan gugur oleh job pada `expires_at`;
  - `Sync` memberi setiap anggota keadaan rumah tangga yang sama dengan bacaan langsung;
  - setiap kueri sqlc lulus uji isolasi yang dibangkitkan;
  - `TestHouseholdProperties` hijau dengan benih tetap.

## Putusan atas pertanyaan rencana (14 November 2026)

- **Q1 (a)** (`architect`, diterima): tepi impor 09.2 §1 butir 10. Tambahannya:
  - penangan job didaftarkan `cmd/casa` ke `jobs.Runner` sebagai fungsi, dan `jobs` tidak pernah mengimpor `household`;
  - `sync.Changed` adalah satu-satunya penulis `reader_feed`;
  - ekspor `identity.AccountEmail` dihapus.

  Diterapkan: Arsitektur; keputusan 6, 8; Tugas 1 sampai 4, 6, 7.
- **Q1 (b)** (`architect`, diterima dengan perubahan; 09.2 §3, §5.1 catatan 14 November 2026):
  - `app_rw` `SELECT, INSERT, DELETE` atas `job`, tanpa `UPDATE`; hak `job_runner` tetap; indeks `job_locked`;
  - kedua kebijakan menyebut perannya (`TO app_rw`, `TO job_runner`), dan gerbang katalog memeriksa perannya;
  - `UNIQUE (kind, unique_key)` diganti `job_unique` (`household_id, kind, unique_key`, `NULLS NOT DISTINCT`, `WHERE unique_key IS NOT NULL`), dan `Enqueue` memakai `ON CONFLICT` atas indeks itu.

  Diterapkan: keputusan 2, 3; Tugas 1, 10.
- **Q1 (c)** (`architect`, diterima; 09.3 §3 dan ADR-017 butir 7, catatan 14 November 2026): satu transaksi `REPEATABLE READ READ ONLY` per halaman, `P0` di kursor ber-MAC. Diterapkan: keputusan 10; Tugas 5.
- **Q1 (d)** (`architect`, diterima dengan perubahan; 09.2 §3 catatan 14 November 2026, 09.3 §3 "Pemangkasan"):
  - rancangan pemangkasan dibangun di langkah 7;
  - di 1.7, kebijakan rumah tangga `reader_feed` sudah `TO app_rw`;
  - aturan resync diperketat menjadi tiga syarat, masing-masing dengan uji.

  Diterapkan: keputusan 5, 11; Tugas 4, 5, 10; "Yang sengaja tidak".
- **Q1 (e)** (`architect`, diterima dengan perubahan; 09.3 §9 catatan 14 November 2026): job per undangan memenuhi "job gugur berkala". Syaratnya:
  - `expires_at` tidak berubah sesudah F01, dan perubahan apa pun kelak memindahkan `run_at` di transaksi yang sama;
  - undangan hidup yang belum jatuh tempo menurut compare-and-set adalah galat yang dicoba ulang, bukan selesai;
  - pengambilan dan compare-and-set membaca `Now` yang sama.

  Diterapkan: keputusan 1; Tugas 2, 10.
- **Q2** (`architect`, diterima; `security-privacy`, diterima dengan enam syarat di 10.1 §9 catatan 1.7 Q2): `identity.send_mail` pindah ke langkah 7 bersama pengirim SES; perilaku 1.4 berlaku sampai itu. Diterapkan: keputusan 5; "Yang sengaja tidak"; "Celah langkah 1".
- **Q3** (`security-privacy`, diterima dengan syarat; 10.1 §10 catatan 1.7 Q3):
  - sisa jendela (≤ 1 detik biasanya; sampai sekitar 85 menit saat coba ulang; tanpa batas sesudah job berhenti) diterima untuk pilot dan ditinjau ulang sebelum GR-4;
  - syarat 1: job dan jalur malas menulis gugur yang identik selain `pos` dan `seq`; uji di Tugas 2 (Peristiwa) dan Tugas 6 (`Sync`, karena `Sync` belum ada di PR A);
  - syarat 2: tidak ada undangan hidup di produksi tanpa job (langkah 7);
  - syarat 3: job gugur yang berhenti adalah perbaikan operator di runbook.

  Diterapkan: keputusan 1; Tugas 2, 6; "Yang sengaja tidak".
- **Rancangan kursor** (`security-privacy`): tanpa cacat. Keputusan 9 tetap.
