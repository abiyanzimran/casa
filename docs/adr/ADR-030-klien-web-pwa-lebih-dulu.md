# ADR-030 — Klien web (PWA) lebih dulu

| | |
|---|---|
| Status | **Diterima.** Web lebih dulu diputuskan pemilik produk sendiri (keputusan 69, dicatat `product-analyst` di 17): "web dibangun duluan aja yang penting web jalan dulu aja. ios dan android nanti. web akan jadi yang pertama untuk user experience". Stack dan rancangan di bawah **diterima atas saran lead**: "ya saya terima kan ini versi webnya, selebihnya saya ikut saran kamu pastikan semua valid dan matang wajib profesional" (keduanya [_progress.md](../_progress.md), "Sesudah S6: audit kesiapan bangun dan usulan web app", 12 November 2026). Butir yang **mengubah bunyi `ENGINEERING-STANDARDS.md`** (butir 15, 16, 17 dan bagian butir 8) menunggu pratinjau kata demi kata kepada pemilik produk, karena ES hanya berubah lewat ADR yang disetujuinya (ES baris 3); sampai itu ES berlaku |
| Tanggal | 12 November 2026 (sesudah S6, putaran web lebih dulu); 13 November 2026 amandemen sesudah tinjauan `security-privacy` (butir 5, 9, 11, 12, 14 menurut SEC-165 (g), (h), SEC-171 (b), SEC-174 (b), SEC-176 (g); LY-95; nama Casa); 13 November 2026 amandemen atas temuan `critic` dan 18 §5.13.10 (keputusan 69 dan 70; bawaan 69.1 sampai 69.6 di butir 8, 15, 16, 19 dan tabel menunggu; pratinjau ES tidak menahan FS-1; Safari sebelum 26 [Usulan]; `strictDepBuilds` versi [Perlu validasi]; butir 5 dan 11 mode bersama di tab Safari menurut SEC-171 (b) (5); butir 9 LY-95 wajib untuk pilot) |
| Pengusul | `architect`, atas 18 §5.13 (`researcher`) dan 10.5 (`security-privacy`, SEC-160 sampai SEC-179, TM-37 sampai TM-45) |
| Mengubah | Urutan platform ADR-012 (catatan keputusan 42); ADR-013, ADR-018, ADR-019, ADR-022, ADR-024, ADR-025, ADR-028, ADR-029 lewat catatan bertanggal di masing-masing; usulan bunyi ES §2, §3, §8 (butir 15 sampai 17) |
| Rincian | [09 indeks](../09-technical-architecture.md); [09.1 §14](../09-technical-architecture/01-pilihan-teknologi.md); [09.2 §5.1](../09-technical-architecture/02-data-dan-skema.md); [09.3 §4, §8](../09-technical-architecture/03-sinkron-waktu-nyata-notifikasi.md); [09.4 §2, §4, §5, §7.1](../09-technical-architecture/04-operasi-uji-biaya.md); [14 §3](../14-engineering-roadmap.md); 15 RISK-94 sampai RISK-97 |

## Konteks

- **Keputusan pemilik produk** (kutipan di baris Status): klien web dibangun pertama dan menjadi klien pertama untuk pengalaman pengguna; iOS dan Android native tetap direncanakan (ADR-012 butir 1 tidak dicabut) dan menyusul. Keputusan ini menggantikan keputusan 41 ("web tidak dibangun di pilot") dan urutan platform keputusan 42. Pemilik produk menerima bahwa selama web menjadi klien pertama, syarat tampilan native brief §10.1 dipenuhi oleh aplikasi native kemudian, bukan oleh web.
- **Bacaan lead [Usulan], dipakai sebagai rencana sampai dikoreksi 17:** GR-2 (data nyata) berjalan di klien web saja, di iPhone 13 (web app layar utama, iOS 26), iPad mini 7 (iPadOS 26), Galaxy A55 (Chrome), dan browser desktop.
- **Fakta yang menentukan** (18 §5.13, diakses 12 November 2026, **[Terverifikasi]** kecuali disebut):
  - Web Push di iOS dan iPadOS hanya untuk web app layar utama; tanpa keanggotaan Apple Developer (18 §5.13.3).
  - Background Sync tidak ada di Safari (18 §5.13.4).
  - Data skrip tab Safari terhapus sesudah 7 hari pemakaian Safari tanpa interaksi; web app layar utama dikecualikan dan **datanya terpisah dari tab Safari** (18 §5.13.2; [webkit.org/tracking-prevention](https://webkit.org/tracking-prevention/)).
  - `EventSource` tidak dapat mengirim header `Authorization` (18 §5.13.1).
  - connect-go menerima protokol Connect dari browser tanpa proxy; dari browser hanya server streaming (18 §5.13.1, §5.12).
  - PWA dibangun dan dideploy tanpa Mac; Web Inspector untuk iPhone dan iPad memerlukan Mac; Playwright WebKit di Linux bukan Safari (18 §5.13.8).
- **Batasan tim** (brief §1): satu sampai tiga pengembang, pemimpin kuat di Go dan Rust, pengembangan di Linux, satu VPS sederhana. Semua aturan uang dan persetujuan tetap satu implementasi di server (ADR-013); klien web, seperti klien native, tidak membawa aturan.
- **Model keamanan browser** dirancang `security-privacy` di 10.5; dokumen 10 §7.9 meminta ADR ini memuat 13 butir, dan 10 §8 butir 16, 18, 19, 20 meminta penyelesaian.

## Opsi

### O-1. Framework UI web

Kriteria: ringan di perangkat pilot (brief §1 "instant cold start"); dukungan PWA dan service worker yang terawat; versi mayor stabil (ES §8, ADR-023); klien Connect berjalan di atasnya; satu pengembang dapat menguasainya.

| Kriteria | **A. SvelteKit 3 (mode SPA) + Svelte 5** | B. Preact + Vite + vite-plugin-pwa | C. Vue 3 | D. React 19 |
|---|---|---|---|---|
| Ukuran runtime | Compiler; tanpa runtime besar (angka vendor tidak ada; diukur di prototipe) | Sekitar 5,0 kB min+gzip (Bundlephobia, **[Perlu validasi]**) | Sekitar 48,1 kB (**[Perlu validasi]**) | Tidak terbandingkan dengan alat yang sama |
| PWA | `@vite-pwa/sveltekit` 2.0.2 (peer SvelteKit ^3) | vite-plugin-pwa 2.0.0 | vite-plugin-pwa | vite-plugin-pwa |
| Router dan struktur bawaan | Ada (mode SPA, adapter-static) | Harus dirakit | Vue Router | Harus dipilih |
| Versi stabil | Svelte 5.57.3; SvelteKit 3.2.0 | Preact 11.2.0 | Vue 3.5.45 | React 19.3.1 |
| Trusted Types | Kebijakan `svelte-trusted-html` dari PR #16271 (**[Perlu validasi]** status merge) | Tergantung kode sendiri | **[Perlu validasi]** | **[Perlu validasi]** |

### O-2. Penyimpanan luring di browser

| Kriteria | **A. IndexedDB lewat Dexie 4** | B. SQLite WASM di OPFS (`@sqlite.org/sqlite-wasm` 3.54.1-build1) |
|---|---|---|
| Kebutuhan | Antrean perintah, potret `Sync`, kursor (ADR-017, ADR-018) | Sama, ditambah SQL lokal |
| Biaya | Satu pustaka JS | Worker, berkas WASM, satu koneksi (`opfs-sahpool`); CSP perlu `'wasm-unsafe-eval'` |
| Manfaat SQL sama dengan native | Tidak ada: native juga tidak berbagi kode klien (ADR-013) | Tidak ada |

### O-3. Autentikasi aliran SSE dari browser

| Kriteria | **A. `EventSource` dengan cookie sesi** | B. Pembaca SSE berbasis `fetch` dengan header | C. Server streaming Connect |
|---|---|---|---|
| Kode klien | Bawaan browser | Pustaka atau parser sendiri | connect-web |
| Titik akhir server | Sama, `GET /v1/stream` (ADR-018) | Sama | Titik akhir kedua |
| Token terbaca skrip | Tidak (`HttpOnly`) | Ya | Ya, bila token di header |

## Keputusan

### Platform dan stack

1. **Urutan platform:** klien web (PWA) → GR-2 di web saja → **iOS dan iPadOS native** (Swift dan SwiftUI) → **Android native** (Kotlin dan Compose). ADR-012 butir 1 sampai 6 tetap berlaku untuk aplikasi native. Urutan iOS lalu Android sesudah GR-2 dipertahankan dari keputusan 42, 50, dan 58. **[Usulan]** `architect`, untuk diputuskan pemilik produk saat jalur native dimulai: karena Galaxy A55 sudah dilayani web sejak GR-2, alasan awal keputusan 42 (istri di iPhone mendapat kehalusan paling banyak) masih menunjuk ke iOS. Urutan tidak diubah di sini.
2. **Stack web: O-1 A, O-2 A, O-3 A.**
   - SvelteKit 3 mode SPA (`export const ssr = false`, `@sveltejs/adapter-static` dengan `fallback`), Svelte 5, Vite, `@vite-pwa/sveltekit` (Workbox 7) dengan strategi `injectManifest`, karena service worker memuat handler `push` dan `notificationclick` sendiri.
   - Dexie 4 di atas IndexedDB.
   - `@connectrpc/connect-web` 2.x dan `@bufbuild/protobuf` 2.x; tipe dibangkitkan dari kontrak Protobuf yang sama (`buf generate`, ES §2).
   - `EventSource` untuk SSE ADR-018.
   - Versi di atas adalah versi 18 §5.13.1 pada 12 November 2026, **bukan** versi untuk dipasang; versi dicek live saat dipasang (ES §8).
   - **Alasan:** A paling ringan dengan router dan PWA bawaan; Dexie cukup karena aturan ada di server; cookie membuat token tidak terbaca skrip sekaligus memungkinkan `EventSource` bawaan.
3. **Satu origin, tanpa CORS** (SEC-160).
   - Biner Go yang sama melayani API Connect, SSE, dan bundel statis di satu origin produksi.
   - **[Usulan]** Bundel disematkan ke biner (`embed`, stdlib), sehingga satu artefak CI, satu hash, dan rollback sama dengan rollback server (09.4 §1). Harga: perubahan web saja tetap memulai ulang proses (sekitar satu detik).
   - Staging di origin lain (subdomain terpisah), dengan bundel staging yang tidak pernah masuk biner produksi (SEC-179).
   - Server tidak mengirim header CORS; permintaan lintas origin gagal di browser.
4. **Batas bawah browser** (SEC-160 (e)) **[Usulan]**:
   - Safari, iOS, dan iPadOS **17** ke atas (alasan: `Clear-Site-Data` dan `navigator.storage.estimate()` di Safari 17; 10.5).
   - Chromium yang masih menerima pembaruan keamanan (Chrome, Edge, Samsung Internet).
   - Firefox tidak diuji di pilot; boleh berjalan bila lolos deteksi fitur.
   - **Deteksi fitur** saat dibuka: IndexedDB, service worker, `navigator.storage`, Web Crypto `subtle`, dan Web Locks. Bila salah satu tidak ada, aplikasi tidak memuat layar masuk dan menampilkan kalimat browser tidak didukung.
   - **Yang diuji:**
     - Safari 26 di web app layar utama iPhone 13 dan iPad mini 7;
     - Chrome di Galaxy A55;
     - browser desktop yang dipakai rumah tangga;
     - Playwright Chromium dan WebKit di CI.
   - Web Locks Baseline "Widely available" sejak Maret 2022 ([MDN Web Locks API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Locks_API), diakses 12 November 2026) **[Terverifikasi]**.

### Sesi dan permintaan

5. **Sesi web** (SEC-161, SEC-165; parameter di 09.4 §2):
   - **Cookie:** satu cookie `__Host-` dengan `Secure`, `HttpOnly`, `SameSite=Strict`, `Path=/`, tanpa `Domain`.
   - **Token:** isinya token opak 256 bit; server menyimpan hash SHA-256 di `session`, dengan `client_kind = 'web'` dan `device_mode` (09.2 §5.1).
   - **Rotasi:** paling sering setiap 15 menit. Nilai sebelumnya diterima 60 detik untuk permintaan yang sedang berjalan. Pemakaian ulang sesudahnya mencabut sesi dan menulis `security_event`.
   - **Umur mode pribadi:** diam 30 hari, mutlak 180 hari; `Max-Age` sama dengan sisa umur mutlak.
   - **Umur mode bersama:** diam 30 menit, mutlak 12 jam. Cookie sesi tanpa `Max-Age` dan `Expires`. Data hanya di memori tab, tanpa IndexedDB, push, dan kunci aplikasi. Di iPhone dan iPad, mode bersama dimasuki langsung di tab Safari lewat LY-90 "Pakai sebagai perangkat bersama atau pinjaman", tanpa memasang ikon (butir 11; SEC-165 (d)).
   - **Masa diam mode bersama hanya dari tindakan orang** (SEC-165 (g)): klien menandai RPC yang lahir dari tindakan pengguna (navigasi, perintah, "Tetap masuk"); hanya RPC bertanda itu yang mengulang pengatur waktu diam di server. Denyut kehadiran, aliran SSE, dan `Sync` otomatis yang dipicu tanda SSE tidak mengulangnya. Tanda dipercaya dari klien, karena ancamannya sesi yang terlupa, bukan klien yang bermusuhan **[Usulan]**. Klien juga menghitung diam dari masukan pengguna dan keluar sendiri pada 30 menit, dengan peringatan 60 detik dan "Tetap masuk"; batas mutlak 12 jam tidak dapat diperpanjang. Medan tanda ada di kontrak (09.4 §2).
   - **Mode bersama tanpa LY-88** (SEC-165 (h)): pemasangan faktor kedua dan kode pemulihan tidak tersedia; layar meminta perangkat pribadi.
   - **Kode galat:**
     - `session_expired`: masuk lagi; data lokal tidak dihapus.
     - `session_revoked`: hapus di semua tab (butir 10).
   - **Daftar sesi** menampilkan browser, sistem operasi, dan mode, dari ringkasan `User-Agent` yang dipotong.
   - **Bacaan ES §4 "Tokens are stored in the platform's secure storage"** (10 §8 butir 18): bagi browser, cookie `HttpOnly` `Secure` `__Host-` yang tidak terbaca skrip **adalah** penyimpanan aman platform. Bunyi ES tidak diubah.
6. **CSRF** (SEC-162):
   - Setiap RPC dengan sesi cookie hanya lewat `POST`; Connect `GET` ditolak bagi sesi cookie.
   - Server menolak bila `Origin` tidak sama dengan origin produksi, atau `Sec-Fetch-Site` ada dan bukan `same-origin`. RPC tanpa `Origin` juga ditolak.
   - `GET /v1/stream` tanpa efek samping; memeriksa `Sec-Fetch-Site` yang sama.
   - Tanpa token CSRF di pilot: `SameSite=Strict` ditambah pemeriksaan header cukup.
   - Klien native (header `Authorization`) tidak terkena pemeriksaan ini, karena tidak membawa cookie.
7. **CSP dan header** (SEC-166, SEC-167):
   - **CSP dari server:** `default-src 'none'; script-src 'self'; style-src 'self'; img-src 'self' blob:; font-src 'self'; connect-src 'self'; worker-src 'self'; manifest-src 'self'; base-uri 'none'; form-action 'self'; frame-ancestors 'none'; object-src 'none'; require-trusted-types-for 'script'; trusted-types svelte-trusted-html`.
   - **Tanpa skrip inline.** Bila SvelteKit mode SPA memerlukan skrip pembuka, ia dihapus dari `index.html` atau diizinkan lewat hash build di `script-src`, tidak pernah lewat `'unsafe-inline'`.
   - **Gaya inline:** bila Svelte 5 memerlukan atribut `style` yang tidak dapat dihindari, `'unsafe-inline'` diizinkan **hanya** di `style-src`, dicatat di sini. Saat ini belum diperlukan **[Perlu validasi]** di prototipe.
   - Tanpa `'wasm-unsafe-eval'` (O-2 A) dan tanpa `'unsafe-eval'`.
   - **Trusted Types:** satu kebijakan bernama `svelte-trusted-html`, bila framework memerlukannya (PR Svelte #16271, status merge **[Perlu validasi]**). Lint melarang `innerHTML`, `{@html}`, `eval`, `new Function`, dan `setTimeout` berisi string, jadi tidak ada kebijakan lain.
   - Trusted Types Baseline 2026 "Newly available" sejak Februari 2026; direktif CSP-nya didukung sejak Safari 26 menurut BCD ([MDN Trusted Types API](https://developer.mozilla.org/en-US/docs/Web/API/Trusted_Types_API); 18 §5.13.10, diakses 13 November 2026) **[Terverifikasi]**. Bahwa Safari 17 sampai 25 mengabaikan direktif itu adalah simpulan dari ketiadaan dukungan, bukan bunyi sumber **[Usulan]**; karena itu lint tetap pengaman utama.
   - **Header lain:**
     - `X-Content-Type-Options: nosniff`;
     - `Referrer-Policy: no-referrer`;
     - `Cross-Origin-Opener-Policy: same-origin` dan `Cross-Origin-Resource-Policy: same-origin`;
     - `Permissions-Policy` mematikan `camera`, `microphone`, `geolocation`, `payment`, `usb`;
     - HSTS `max-age` sekurangnya satu tahun dengan `includeSubDomains`;
     - API dan SSE `Cache-Control: no-store`; `index.html` dan service worker `no-cache`; berkas berhash `immutable`.
   - Uji header di CI (09.4 §4) memblokir merge.

### Rantai pasok dan rilis

8. **Dependensi npm** (SEC-168; 10 §8 butir 16):
   - **Pengelola paket pnpm.** `strictDepBuilds` bawaannya `true`: instalasi gagal bila ada dependensi dengan skrip build yang belum ditinjau, dan paket di luar `allowBuilds` tidak menjalankan skrip ([pnpm.io/settings/build](https://pnpm.io/settings/build); 18 §5.13.10, diakses 13 November 2026) **[Terverifikasi]**. Versi pnpm sejak kapan bawaan ini berlaku tidak tercantum di sumber **[Perlu validasi]**; karena itu `strictDepBuilds: true` ditulis eksplisit di `pnpm-workspace.yaml`, tidak bergantung pada bawaan. Daftar `allowBuilds` kosong kecuali disetujui `code-reviewer`.
   - **Lockfile** `pnpm-lock.yaml` di repositori; instalasi CI `--frozen-lockfile`.
   - **Pemindai:** OSV-Scanner atas `pnpm-lock.yaml`, yang tercantum di daftar lockfile resminya ([google.github.io/osv-scanner](https://google.github.io/osv-scanner/supported-languages-and-lockfiles/); 18 §5.13.10, diakses 13 November 2026) **[Terverifikasi]**. Satu alat untuk Kotlin dan web (ADR-025 catatan web).
   - **Dependensi bundel** (yang masuk kode yang berjalan di browser) dibatasi pada daftar ini: `svelte`, `@sveltejs/kit` (runtime klien), `dexie`, `@connectrpc/connect`, `@connectrpc/connect-web`, `@bufbuild/protobuf`, `workbox-*` yang ditarik `@vite-pwa/sveltekit`. Tambahan memerlukan tinjauan `security-privacy` dan amandemen butir ini. Dependensi pengembangan (Vite, alat lint, Playwright) tidak masuk bundel; CI memeriksa daftar paket di peta bundel.
   - **Masa tunggu (isi diterima, bawaan 69.1; bunyi ES §8 menunggu pratinjau pemilik produk):** versi baru dependensi bundel baru dipasang sesudah berumur 7 hari, lewat `minimumReleaseAge: 10080` (menit) pnpm. Bawaan pnpm 1440 menit ([pnpm.io/settings/dependency-resolution](https://pnpm.io/settings/dependency-resolution); 18 §5.13.10, diakses 13 November 2026) **[Terverifikasi]**. Perbaikan keamanan dikecualikan lewat `minimumReleaseAgeExclude` per paket, dicatat di pull request-nya.
   - Sampai bunyi ES §8 disetujui, ES §8 berlaku apa adanya (tanpa masa tunggu) dan sisa butir ini tetap; tidak ada langkah yang tertahan.
9. **Build, service worker, dan versi** (SEC-169):
   - **Build:** bundel dibangun hanya di CI dari komit bertanda. Daftar hash setiap berkas disimpan sebagai artefak CI. Deploy memasang biner yang memuat bundel itu tanpa build ulang di VPS. Runbook rilis membandingkan hash biner di VPS dengan artefak CI.
   - **Cache service worker:** hanya app shell berhash. Tidak pernah jawaban API, SSE, foto, atau halaman berisi data.
   - **Pergantian versi:** versi baru menunggu (`skipWaiting` tidak otomatis). Ia aktif saat aplikasi dibuka berikutnya, atau atas ketukan "Muat ulang" ketika tidak ada isian terbuka. Antrean tidak pernah dibuang.
   - **Versi klien di kontrak** (ADR-013 catatan web): setiap permintaan membawa `client_kind` dan `client_version`. Server menolak perintah tulis dari versi di bawah versi minimum `web` dengan kode `CLIENT_UPGRADE_REQUIRED` (nama dari ADR-028 butir 5); perintah tetap di antrean.
   - Batas minimum disimpan di konfigurasi server dan dinaikkan pengelola pada rilis yang memperbaiki celah (TM-45). Ini berlaku sejak pilot bagi web, berbeda dari ADR-028 butir 5 bagi native.
   - **Service worker pengganti:** berkas darurat yang hanya mengosongkan cache app shell lalu mendaftar ulang, tanpa menyentuh IndexedDB. Disimpan di repositori dan dilatih di staging sebelum GR-2.
   - **LY-95 "Perangkat ini"** (08 §25.6) menampilkan versi dan hash bundel, **wajib untuk pilot** (SEC-169 (f)), sebagai alat periksa, bukan bukti (RISK-95).
   - **Nama yang terlihat pengguna:** manifest web app `name` dan `short_name` "Casa", juga awal `document.title` (SEC-175 (a), (g)), memakai nama kerja sementara Casa (keputusan 70, `_progress.md` "Keputusan pemilik produk: nama aplikasi"). "Family OS" tetap nama proyek perencanaan dan pengenal yang sudah ada.
   - Halaman pemeriksaan diri (SEC-179) hanya ada di build staging. CI menolak bundel produksi yang memuat penandanya atau alamat staging.

### Data di perangkat

10. **Penyimpanan, antrean, dan tab** (SEC-170 sampai SEC-173; ADR-024 catatan web):
    - Isi IndexedDB menurut tabel 10.5 §5. Status `CARE` yang tidak dibagikan (Tier 2) di store tersendiri; tanpa padanan Class A, dinyatakan.
    - `navigator.storage.persist()` diminta sesudah masuk di mode pribadi; hasilnya tampil di LY-95 "Perangkat ini".
    - **Pengiriman antrean:** saat aplikasi dibuka, saat peristiwa `online`, saat `visibilitychange` menjadi terlihat, dan sesudah setiap perintah baru.
    - **Kunci antartab:** `navigator.locks.request('outbox', …)` (Web Locks) memastikan satu pengirim per profil browser. Idempotency key (ADR-005) tetap penjaga terakhir.
    - **Siaran antartab:** `BroadcastChannel` membawa keluar akun, `session_revoked`, F06, dan kunci aplikasi ke semua tab.
    - **Hapus jarak jauh:** pada `session_revoked`, semua tab menghapus IndexedDB, Cache Storage selain app shell, langganan push, dan memori sebelum menampilkan layar lain.
    - **Pemeriksaan sesi saat dibuka:** paling lama 3 detik **[Asumsi]** sebelum data lokal tampil.
    - **Kehilangan dinyatakan, tidak disembunyikan** (SEC-171 (e), (f)).
11. **Pasang ke layar utama di iPhone dan iPad** (10 §8 butir 19; SEC-171 (b), dikukuhkan `security-privacy`):
    - Di iOS dan iPadOS, tab Safari, termasuk browser di dalam aplikasi lain, **tidak pernah menawarkan mode pribadi**. LY-90 menawarkan dua jalan (SEC-171 (b) (5); SEC-165 (d); SEC-161 (e)):
      1. **Perangkat sendiri:** petunjuk memasang ke layar utama; masuk mode pribadi hanya di web app yang terpasang.
      2. **"Pakai sebagai perangkat bersama atau pinjaman":** membuka masuk **mode bersama** langsung di tab itu, tanpa memasang ikon dan tanpa LY-91. Alasannya: mode bersama tidak menyimpan apa pun (butir 5), sedangkan ikon terpasang meninggalkan jejak di perangkat orang lain.
    - Deteksi lewat `matchMedia('(display-mode: standalone)')` dan `navigator.standalone`.
    - **Tautan di tab:** verifikasi email dan pemulihan kata sandi diselesaikan tanpa membuat sesi di tab (LY-92); token undangan dibuang dari bilah alamat dan tidak dikirim ke server dari tab, lalu diterima lewat tautan yang ditempel di web app (SEC-171 (b) (2), (3); SEC-175 (c)).
    - **Bila deteksi tidak pasti** (iPad dapat melapor seperti komputer): halaman diperlakukan sebagai browser biasa dan masuk diizinkan, dengan peringatan SEC-171 (c) bila berjalan di tab atau `persist()` ditolak, supaya tidak ada pemilik akun yang terkunci di luar (08 §25.6) **[Usulan]**. Risikonya kehilangan antrean (RISK-94), bukan kebocoran.
    - **Alasan:** data web app layar utama terpisah dari tab Safari (Konteks), sehingga masuk mode pribadi di tab menghasilkan sesi kedua yang datanya terhapus sesudah 7 hari dan tidak ikut ke web app. Masuk sekali di tempat yang benar mengurangi beban dan kehilangan antrean (TM-44).
    - Bagi pengguna, ini satu kali pemasangan per perangkat, bukan beban berulang.
    - Android dan desktop: pemasangan dianjurkan (ajakan sekali), tidak diwajibkan.
    - **Risiko sisa:** browser masih dapat menghapus data situs karena tekanan ruang atau tindakan pengguna; jumlah catatan belum terkirim selalu tampil. Penutup sebenarnya adalah aplikasi native (RISK-94).
    - Dikukuhkan `security-privacy` di SEC-171 (b), yang menggantikan "ajakan memasang" draf pertama.
12. **Kunci aplikasi web** (SEC-174):
    - Kode 6 digit saja.
    - **Pemverifikasi:** `PBKDF2-HMAC-SHA-256` dari Web Crypto (`crypto.subtle.deriveBits`), **600.000 iterasi**, garam acak 128 bit, keluaran 256 bit. Disimpan di IndexedDB bersama parameter, sehingga iterasi dapat dinaikkan kemudian. Angka iterasi dari OWASP Password Storage Cheat Sheet: "PBKDF2-HMAC-SHA256: 600,000 iterations" ([cheatsheetseries.owasp.org](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html); 18 §5.13.10, diakses 13 November 2026) **[Terverifikasi]**.
    - **Batas yang jujur:** ruang kode 6 digit hanya 10^6. Pemverifikasi memperlambat tebakan luring, tetapi tidak mencegahnya, sehingga kunci ini hanya menutup tampilan (SEC-174 (c)).
    - **Aturan penurunan iterasi** (SEC-174 (b)) **[Usulan]**: waktu derivasi diukur di prototipe pada perangkat pilot paling lambat yang memakai kunci (iPhone 13, iPad mini 7, Galaxy A55). Bila melebihi 1 detik, iterasi diturunkan ke kelipatan 50.000 terbesar yang selesai dalam 1 detik di perangkat itu, satu angka untuk semua perangkat, **tidak pernah di bawah 300.000**. Bila 300.000 pun melebihi 1 detik, angka itu tetap dan waktu bukanya diterima. Amandemen butir ini di dalam aturan itu tidak memerlukan tinjauan baru; di luar aturan itu memerlukan tinjauan `security-privacy`. Pemverifikasi diturunkan ulang dengan parameter baru saat kode benar dimasukkan berikutnya.
    - Kalimat tablet tampil di semua perangkat web, karena iPad tidak dapat dikenali dengan pasti dari browser (SEC-174 (a)); kunci tidak ditawarkan di mode bersama.
    - **Autentikasi ulang** memakai kata sandi akun, ditambah faktor kedua bila menyala.
    - **Passkey** (10 §8 butir 20) tidak di pilot. Bila kelak dirancang, passkey boleh dipakai untuk masuk, tetapi autentikasi ulang SEC-13 tetap meminta kata sandi atau faktor yang tidak dapat dipenuhi orang lain yang terdaftar di perangkat itu, kecuali ADR baru menyatakan lain dengan alasan.

### Waktu nyata dan notifikasi

13. **SSE di web** (ADR-018 catatan web; SEC-163):
    - `new EventSource('/v1/stream')`, same-origin, sehingga cookie ikut tanpa `withCredentials`.
    - Server mengautentikasi lewat cookie (web) atau header `Authorization` (native) di titik akhir yang sama.
    - Aliran ditutup seketika saat sesi dicabut atau berakhir, dengan peristiwa terakhir berisi kode galat.
    - Status "di depan" klien web diturunkan dari `document.visibilityState`.
14. **Web Push** (ADR-019 catatan web; SEC-176, SEC-177):
    - **Kanal ketiga** di samping APNs dan FCM, dikirim langsung dari biner Go dengan VAPID (RFC 8292) dan enkripsi payload RFC 8291.
    - **Pustaka:** `github.com/SherClockHolmes/webpush-go` v1.4.0 (MIT; 18 §5.13.3); status perawatannya **[Perlu validasi]** saat dipasang. Jalan mundur: implementasi RFC 8291 di atas stdlib (`crypto/ecdh`, `crypto/hkdf`, `crypto/aes`), dengan vektor uji RFC 8291 Lampiran A.
    - **Payload** sama dengan SEC-50, ditambah angka lencana, di bawah 4 KB, dan **selalu ditampilkan** (notifikasi tak terlihat mencabut izin di Safari). `tag` dan `Topic` opak, paling banyak 32 karakter.
    - **JWT VAPID:** ES256 P-256; `exp` paling lama 24 jam; diperbarui paling sering sekali per jam; `sub` alamat operasi umum. Kunci privat lewat `LoadCredential=` (09.4 §1).
    - **Langganan** per sesi di `push_token` dengan `platform = 'webpush'`, berisi endpoint, `p256dh`, dan `auth` (09.2 §5.1). Dihapus pada jalur yang sama dengan token SEC-58.
    - **Daftar izin host endpoint** (SEC-176 (g); TM-43). Diperiksa **saat berlangganan dan sebelum setiap pengiriman**:
      1. URL diurai dengan pengurai standar (`net/url`), bukan pencocokan string; skema `https`, port kosong atau 443, tanpa userinfo, host berupa nama, bukan literal IP.
      2. Host sama persis dengan baris tanpa `*`; untuk baris `*.`, host berakhiran titik ditambah sufiks itu (`x.push.apple.com` cocok; `push.apple.com.contoh.id` dan `evilpush.apple.com` tidak).
      3. Alamat hasil resolusi DNS bukan loopback, jaringan pribadi, link-local, atau alamat metadata (lapis kedua).
      4. Tanpa mengikuti redirect; batas waktu permintaan; isi jawaban tidak dibaca selain kode status.
      5. Panjang endpoint paling banyak 2.048 karakter **[Asumsi]**.

      Endpoint yang tidak cocok ditolak saat berlangganan tanpa permintaan keluar. Daftar disimpan di konfigurasi server dan diubah hanya lewat amandemen butir ini:

      | Host | Browser | Label |
      |---|---|---|
      | `*.push.apple.com` | Safari, iOS, iPadOS, macOS | **[Terverifikasi]** 18 §5.13.3 |
      | `fcm.googleapis.com` | Chrome, Edge Android, Samsung Internet | **[Perlu validasi]** |
      | `*.notify.windows.com` | Edge desktop | **[Perlu validasi]** |
      | `updates.push.services.mozilla.com` | Firefox | **[Perlu validasi]** |

      Tiga baris terakhir dari daftar pihak ketiga (`github.com/pushpad/known-push-services`), bukan daftar resmi vendor. Diperiksa di perangkat pilot sebelum GR-2: endpoint yang dihasilkan Chrome di Galaxy A55 dan browser desktop yang dipakai harus cocok; bila tidak, daftar diperbarui di sini.

### Standar rekayasa

15. **Bantuan isian di TypeScript** (ADR-013 butir 3; isi diterima, bawaan 69.2; bunyi ES §2 menunggu pratinjau pemilik produk):
    - Ketiga bantuan isian (selisih baris, bagi rata, bentuk masukan) ditulis sekali lagi di TypeScript untuk klien web.
    - Vektor uji emas dari acuan Go yang sama dijalankan di CI web (Vitest); beda hasil memblokir merge.
    - **protovalidate-es tidak dipilih:** menambah dependensi bundel dan runtime CEL yang kecocokannya dengan CSP tanpa `'unsafe-eval'` belum diperiksa.
    - **Usulan bunyi ES §2:** "…may be hand-written once in Swift, once in Kotlin, and once in TypeScript for the web client…"; sisa kalimat tetap.
   - **Sampai bunyi itu disetujui**, layar web tidak memuat bantuan isian, dan server menolak dengan kode galat per medan (ADR-013 opsi B), sehingga tidak ada kode yang melanggar ES yang berlaku. **Tidak ada yang tertahan:** FS-1 langkah 1 dan seterusnya berjalan; bantuan isian ditambahkan sesudah persetujuan sebagai perubahan biasa.
16. **Kinerja dan perangkat ukur** (ADR-029 catatan web; isi diterima, bawaan 69.3; bunyi ES §3 menunggu pratinjau pemilik produk):
    - Angka `[DEFAULT]` ES §3 dikukuhkan di klien web pada iPhone 13, iPad mini 7, dan Galaxy A55 sebelum GR-2.
    - Klien native diukur sebelum masing-masing memegang data nyata, dengan ADR kinerja yang sama direvisi.
    - "App download size ≤ 30 MB" berpadanan web: **ukuran transfer app shell terkompresi ≤ 500 KB [Usulan] [Asumsi]**, dikukuhkan sesudah prototipe.
    - **Usulan bunyi ES §3:** "…on the pilot devices: the web client on iPhone 13, iPad mini 7 and Galaxy A55 before the first real data; each native app on its pilot devices before it holds real data (ADR-029, ADR-030)."
17. **Masa tunggu dependensi bundel** (butir 8): **usulan tambahan ES §8:** "Exception: web bundle dependencies are installed only after a version has been public for 7 days; security fixes are exempt (ADR-030)."
18. **Alat kualitas web** (ADR-025 catatan web):
    - Formatter: Prettier dengan `prettier-plugin-svelte`.
    - Linter: ESLint dengan `typescript-eslint` dan `eslint-plugin-svelte`, ditambah `svelte-check` untuk tipe.
    - Pemindai: OSV-Scanner atas `pnpm-lock.yaml`.
    - Duplikasi: PMD CPD (dukungan TypeScript dan berkas `.svelte` **[Perlu validasi]**).
    - Uji: Vitest dan Playwright (Chromium dan WebKit).
    - Semua ≥ 1.0 **[Perlu validasi]** saat dipasang (ES §8, ADR-023), dan semuanya gerbang merge.

### Mac dan Apple

19. **Mac dan Apple** (ADR-022 catatan web): Apple Developer Program, Xcode Cloud, dan sewa Mac sekali **tidak diperlukan** sampai jalur iOS native dimulai. Sumber Mac sekali: **AWS EC2 Mac atau Mac pinjaman** (bawaan 69.6, diterima atas saran lead; temuan audit 2); pilihan di antara keduanya diambil pemilik produk saat jalur iOS native dimulai, sesuai pernyataannya "itu nanti aja" (`_progress.md`). Karena keduanya tidak memerlukan pemeriksaan lisensi penyedia lain, L-31 butir 3 tidak perlu dijawab. Scaleway tetap hanya sebagai catatan riwayat (ADR-022 butir 2).

## Yang diputuskan sekarang dan yang menunggu

| Sekarang (ADR ini) | Menunggu, dengan titiknya |
|---|---|
| Urutan web lebih dulu; stack O-1 A, O-2 A, O-3 A; satu origin; sesi, CSRF, CSP; npm dengan pnpm dan OSV-Scanner; service worker dan versi minimum; Web Push | Pratinjau bunyi ES §2, §3, §8 (butir 15 sampai 17; isinya diterima sebagai bawaan 69.1 sampai 69.3): lead kepada pemilik produk, sebaiknya sebelum FS-1 langkah 1. **Tidak menahan FS-1:** sampai bunyi disetujui, klien web tanpa bantuan isian (butir 15), tanpa masa tunggu dependensi (butir 8, 17), dan ES §3 dibaca apa adanya; bantuan isian TypeScript ditambahkan sesudahnya sebagai perubahan biasa |
| PBKDF2 600.000 | Waktu derivasi di perangkat: prototipe FS-1 langkah 1 |
| Batas bawah Safari 17 | Kebijakan Trusted Types Svelte, gaya inline, host push non-Apple: prototipe FS-1 langkah 1 sampai 2 |
| Urutan sesudah GR-2: iOS native lalu Android native (bawaan 69.5, diterima atas saran lead) | Kapan jalur iOS native dimulai di antara gelombang Tier 2: pemilik produk di GR-2 |
| Sumber Mac: AWS EC2 Mac atau Mac pinjaman (bawaan 69.6) | Pilihan di antara keduanya: saat jalur iOS native dimulai |
| — | Ukuran app shell ≤ 500 KB: dikukuhkan sesudah prototipe |

## Konsekuensi

- **Tiga UI kelak:** web, SwiftUI, dan Compose. Semua aturan tetap sekali di server; yang ditulis tiga kali hanya layar dan tiga bantuan isian, dijaga vektor emas. Kapasitas tim dinilai di RISK-53.
- **Web tidak terasa native** di ponsel: komponen sistem tidak dapat disamai, dan yang sama hanya IA, token arah C, dan kata-kata (18 §5.13.7). **Diterima pemilik produk** (kutipan di Status); syarat brief §10.1 dipenuhi aplikasi native (RISK-97).
- **Jalur kritis tanpa Apple:**
  - Siklus UI cepat di Linux (Vite, browser, Playwright).
  - RISK-01, RISK-52, dan RISK-92 keluar dari jalur GR-2; biaya Apple USD 99 per tahun hilang dari biaya pilot sampai iOS native (09.4 §7.1).
  - Harga: uji di iPhone dan iPad tanpa Web Inspector, jadi berupa pengamatan tangan dan halaman pemeriksaan diri staging (SEC-179).
- **Galaxy A55 memegang data nyata sejak GR-2** lewat Chrome. Pengeluaran pemilik produk di luar rumah tidak lagi menunggu jalur Android (RISK-18).
- **Antrean luring lebih rapuh daripada native** (TM-44; RISK-94).
- **Pengelola server dapat menyajikan kode web berbeda per akun tanpa jejak pemasangan** (TM-42; RISK-95). Di server pasangan saja, ini termasuk keadaan keputusan 25; ADR E2EE sebelum tahap publik harus menyatakan bahwa klien web tidak memberi perlindungan berarti terhadap pengelola (10 §8 butir 17).
- **Rantai pasok npm** menjadi permukaan baru (TM-38; RISK-96).
- **Versi klien minimum** sudah berlaku sejak pilot bagi web; aplikasi native mengikuti ADR-028 butir 5.
- **Tidak ada perubahan** pada kontrak aturan, skema uang, invarian 06 dan 07, atau ADR-013 butir 1 dan 2.
