# FS-1 1.9a — Inti klien web, bagian tanpa server: rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** bagian klien web yang tidak memerlukan server: Vitest sebagai gerbang, basis data lokal Dexie, antrean perintah dengan jeda urungkan dan kunci antartab, siaran antartab dan hapus lokal, service worker app shell beserta service worker darurat, deteksi fitur dengan kalimat LY-97, deteksi web app layar utama, dan permintaan penyimpanan tetap. Semuanya teruji, dan 1.9b tinggal memasang connect-web.

**Arsitektur:** modul TypeScript kecil di `web/src/lib/`, satu tanggung jawab per berkas, tanpa hierarki antarmuka. Antrean tidak mengenal kontrak: ia menyimpan payload biner dari pemanggil dan mengirim lewat satu parameter fungsi `send`, yang di 1.9b diisi klien connect-web. Nilai per jenis perintah (`hold_seconds`, `protective`) datang dari pemanggil, karena kelak dibaca dari tabel jenis perintah di kontrak (ADR-013 butir 4; 09.3 §1). Klien tidak menulis angka 5 sendiri. Basis data tidak pernah dibuka saat modul diimpor, karena mode bersama tidak memakai IndexedDB (SEC-165; ADR-024 butir 7 (e)) dan basis data dibuat saat masuk (SEC-171 (e)), yaitu di 1.10.

**Stack:** yang sudah ada dari 1.1b, ditambah `dexie` (bundel), `vitest` dan `fake-indexeddb` (pengembangan), serta alat service worker menurut keputusan di Tugas 5 langkah 1.

**Spesifikasi:** ADR-030 butir 2, 4, 8, 9, 10, 11; ADR-005; ADR-017 (UUIDv4); ADR-024 "Catatan keputusan web lebih dulu" butir 7; 10.5 §5, SEC-160 (e), SEC-164 (c), SEC-165, SEC-169 sampai SEC-173, SEC-175 (g); 09.3 §1 (urutan kirim, jeda urungkan 5 detik, perintah pelindung); 08 §9.3, §13.2, §25.6 (LY-94, LY-97, deteksi tab Safari), §25.7; FR-FIN-040/AC-17, AC-18; 14 §3.3 langkah 1, §4.7; ES §1, §2, §3, §6, §8.

**Pelaksana:** `web-dev`. **Peninjau:** `code-reviewer` (dua lintasan) untuk setiap tugas; `ux-designer` QG-09 untuk Tugas 2 (LY-97). **Cabang:** `fs1/1.9a-inti-klien-web` dari `main`. Tugas 1.2 berjalan bersamaan dan juga mengubah `Makefile` dan `.github/workflows/ci.yml`; konflik diselesaikan dengan mempertahankan target dan langkah keduanya.

## Batasan global

- Semua batasan global 1.1a dan 1.1b berlaku.
- **Versi dicek live saat dipasang** (ES §8), dipatok tepat (`pnpm add -E`). Masa tunggu 7 hari ditegakkan `minimumReleaseAge: 10080` yang sudah ada di `web/pnpm-workspace.yaml`. Rencana ini tidak menetapkan nomor versi. Paket 0.x yang bukan pra-rilis hanya boleh bila tercantum di ADR-023 (ES §8), juga untuk alat build.
- **Daftar dependensi bundel tertutup** (ADR-030 butir 8): `svelte`, `@sveltejs/kit` (runtime klien), `dexie`, `@connectrpc/*`, `@bufbuild/protobuf`, `workbox-*` yang ditarik `@vite-pwa/sveltekit`. `fake-indexeddb` dan `vitest` hanya `devDependencies` dan tidak pernah diimpor kode di luar `*.test.ts`. Kode service worker juga termasuk bundel, karena berjalan di browser.
- **Penyimpanan menurut 10.5 §5.** Basis data `casa` berisi dua store saja: `outbox` (perintah beserta idempotency key) dan `meta` (hasil `persist()`; kelak pengenal akun pemilik data dan kursor sinkron, 1.9b dan 1.10). Store data lain datang bersama fiturnya: cermin `Sync` (1.9b dan langkah 2), draf isian (langkah 2), cache foto (langkah 5), pemverifikasi kunci aplikasi (langkah 6), status `CARE` (Tier 2). Tidak ada apa pun di `localStorage` atau `sessionStorage`, dan tidak ada token di tempat mana pun.
- **Idempotency key:** `crypto.randomUUID()` (UUIDv4, ADR-017), dibuat saat perintah disimpan, bukan saat dikirim, dan sama di setiap percobaan (ADR-005).
- **Antrean tidak pernah dibuang diam-diam.** Satu-satunya penghapusan baris: diterima server, urungkan selama jeda, dan hapus lokal (yang jumlahnya dinyatakan di LY-94).
- **Log** (ES §4): galat yang dilaporkan ke `console.error` tidak pernah memuat payload perintah.
- **Nama uji:** diawali ID bila ada (`SEC-173 (a): …`, `FR-FIN-040/AC-17: …`), berupa string literal biasa pada `test(`/`it(`. Pemeriksa ID (QG-06) tidak membaca ID dari `test.each`, template literal, atau judul `describe`. Uji Vitest di `web/src/**/*.test.ts` dan uji Playwright di `web/tests/e2e/*.spec.ts`, yaitu akar yang dipindai pemeriksa ID.
- **Kedua mesin** (ADR-030 butir 4): setiap uji Playwright berjalan di Chromium dan WebKit. Bila sebuah uji tidak dapat berjalan di salah satunya, pelaksana berhenti dan melapor; tidak ada `test.skip` per mesin tanpa persetujuan lead.
- Teks layar dari 08, kata demi kata. Pengenal memakai istilah Inggris glosarium. Untuk istilah yang belum ada di glosarium dipakai sementara `outbox`, `wipe`, `sign_out`, `household_exit` (F06), dan konfirmasinya diminta ke lead (lihat laporan).

## Fokus tinjauan

1. **Jam perangkat dimundurkan sesudah menyimpan** (zona waktu, setelan manual): perintah dengan `releaseAt` jauh di depan jam tidak boleh tertahan berjam-jam. Perintah dianggap jatuh tempo juga bila `createdAt > now`. Uji di Tugas 4.
2. **Tab ditutup atau dimatikan di tengah jeda 5 detik:** perintah terkirim pada pembukaan berikutnya, hanya kesempatan mengurungkan yang hilang (09.3 §1). Uji di Tugas 4.
3. **Hapus lokal saat pengiriman sedang berjalan:** penulisan status yang terlambat tidak boleh membuat ulang basis data yang baru dihapus. Uji di Tugas 6.
4. **Galat tulis saat menyimpan** (`QuotaExceededError`): `enqueue` menolak dengan galat aslinya, tidak ada baris setengah jadi, dan tidak ada pengiriman (SEC-171 (f)). Uji di Tugas 4.
5. **Pesan siaran dari versi aplikasi lain** (tab lama sesudah pembaruan) atau pesan yang rusak: diabaikan, tanpa hapus lokal dan tanpa galat. Uji di Tugas 6.

## Yang sengaja tidak di tugas ini

- connect-web, `client_kind`/`client_version`, putaran sinkron dan kursor, pemeriksaan sesi 3 detik saat dibuka (SEC-172 (b)), SSE, pemicu hapus dari jawaban server, LY-96 (SEC-171 (e)), uji Playwright antrean ujung ke ujung dengan kirim sungguhan (tutup konteks lalu buka lagi; dua tab luring lalu daring): **1.9b**. Di sini perilaku yang sama dibuktikan di tingkat unit.
- Layar masuk, daftar, LY-90, LY-91, LY-94, LY-95, buat rumah tangga, undang; pemanggil `requestPersistence` dan `displayContext`; tampilan jumlah belum terkirim: **1.10** dan layar berikutnya. Jumlahnya sudah tersedia lewat `unsentCount()`.
- Kunci aplikasi dan penanganan sinyal `app_lock`: langkah 6 (SEC-174). Di sini hanya jenis pesannya.
- Pemicu "Muat ulang"/"Perbarui" dari klien (SEC-169 (c), 08 §25.7 versi minimum): 1.9b bersama `CLIENT_UPGRADE_REQUIRED`. Di sini service worker hanya menanggapi pesan `SKIP_WAITING`.
- Latihan service worker darurat di staging: langkah 7 (14 §3.3; SEC-169 (e)). LY-95 versi dan hash (SEC-169 (f)): bersama LY-95.
- Laporan cakupan Vitest (14 §2.2, "tanpa ambang"): bukan gerbang dan menambah satu dependensi pengembangan; menunggu putusan lead.

---

### Tugas 1: Vitest sebagai gerbang, dan deteksi fitur serta mode tampilan

**Berkas:**
- Buat: `web/src/lib/device.ts`, `web/src/lib/device.test.ts`
- Ubah: `web/package.json`, `web/pnpm-lock.yaml`, `web/vite.config.ts` (blok `test`), `Makefile` (target `web-unit`, masuk `check-web`), `.github/workflows/ci.yml` (langkah `web-unit (QG-05)` dengan `if: ${{ !cancelled() }}`, sebelum `web-e2e`), `README.md` akar (satu baris `make web-unit`)

**Antarmuka:**
- Menghasilkan:
  - `type Feature = 'indexedDB' | 'serviceWorker' | 'storage' | 'crypto.subtle' | 'locks'`
  - `missingFeatures(scope: BrowserScope = globalThis): Feature[]`. Pemeriksaannya berdasarkan nilai truthy, bukan `in`, karena di konteks tidak aman `navigator.serviceWorker` dan `crypto.subtle` bernilai `undefined`. `BrowserScope` hanya memuat sifat yang dibaca.
  - `type DisplayContext = 'standalone' | 'ios-safari-tab' | 'browser'`
  - `displayContext(scope: BrowserScope = globalThis): DisplayContext`

- [ ] **Langkah 1:** cek live versi stabil `vitest` yang berumur ≥ 7 hari dan kompatibel dengan Vite terpasang (peer `vite`). Pasang `pnpm add -DE vitest`. Tambahkan `test: { include: ['src/**/*.test.ts'], environment: 'node', restoreMocks: true, unstubGlobals: true }`. Target `web-unit`: `cd web && pnpm exec svelte-kit sync && pnpm exec vitest run`.
- [ ] **Langkah 2: tulis uji yang gagal**

```ts
const full = { indexedDB: {}, crypto: { subtle: {} }, navigator: { serviceWorker: {}, storage: {}, locks: {} } };

test('SEC-160 (e): a browser with all five required features misses none', () => {
	expect(missingFeatures(full)).toEqual([]);
});
test('SEC-160 (e): each required feature is reported when it is absent', () => {
	expect(missingFeatures({ ...full, indexedDB: undefined })).toEqual(['indexedDB']);
	expect(missingFeatures({ ...full, crypto: {} })).toEqual(['crypto.subtle']);
	for (const f of ['serviceWorker', 'storage', 'locks'] as const) {
		expect(missingFeatures({ ...full, navigator: { ...full.navigator, [f]: undefined } })).toEqual([f]);
	}
});
test('SEC-171 (b): standalone display mode or navigator.standalone true is the home-screen web app', () => {
	expect(displayContext(scope({ standaloneMedia: true }))).toBe('standalone');
	expect(displayContext(scope({ navigatorStandalone: true }))).toBe('standalone');
});
test('SEC-171 (b): navigator.standalone false is a Safari tab on iPhone or iPad', () => {
	expect(displayContext(scope({ navigatorStandalone: false }))).toBe('ios-safari-tab');
});
test('SEC-171 (b): without navigator.standalone the page is an ordinary browser', () => {
	// 08 §25.6: when detection is uncertain (iPad reporting as a computer), LY-91, not LY-90.
	expect(displayContext(scope({}))).toBe('browser');
});
```

  `scope({...})` adalah pembantu di berkas uji yang membuat `matchMedia` palsu untuk `'(display-mode: standalone)'` dan `navigator.standalone` hanya bila diberikan.

- [ ] **Langkah 3:** `make web-unit` → GAGAL (modul tidak ada).
- [ ] **Langkah 4:** tulis `device.ts`. Urutan di `displayContext`: `matchMedia('(display-mode: standalone)').matches` atau `navigator.standalone === true` menghasilkan `'standalone'`; lalu `navigator.standalone === false` menghasilkan `'ios-safari-tab'`; selain itu `'browser'`.
- [ ] **Langkah 5:** `make web-unit` → LULUS, 5 uji. `make web-lint web-check` → bersih.
- [ ] **Langkah 6: commit** `Add Vitest gate and browser feature and display-mode detection (SEC-160 (e), SEC-171 (b); ADR-030 items 4, 11)`

### Tugas 2: LY-97 "Browser tidak didukung" saat dibuka

**Berkas:**
- Buat: `web/src/routes/+layout.svelte`, `web/tests/e2e/unsupported.spec.ts`

**Antarmuka:**
- Memakai: `missingFeatures()` dari Tugas 1.
- Menghasilkan: layout akar yang menampilkan LY-97 sebagai pengganti seluruh isi bila `missingFeatures().length > 0`. Bila tidak, ia merender `children`. Tugas 6 menambah pendengar siaran ke berkas ini.

- [ ] **Langkah 1: tulis uji yang gagal.** Kalimat dari 08 §25.6 LY-97, kata demi kata:

```ts
const ly97 = 'Browser ini belum bisa menjalankan Casa dengan aman. Perbarui browser ini, atau buka Casa di browser lain yang lebih baru.';
const removals = [
	() => Object.defineProperty(window, 'indexedDB', { value: undefined }),
	() => Object.defineProperty(Navigator.prototype, 'serviceWorker', { get: () => undefined }),
	() => Object.defineProperty(Navigator.prototype, 'storage', { get: () => undefined }),
	() => Object.defineProperty(Crypto.prototype, 'subtle', { get: () => undefined }),
	() => Object.defineProperty(Navigator.prototype, 'locks', { get: () => undefined })
];

test('SEC-160 (e): without any one required feature only the unsupported-browser sentence shows', async ({ browser }) => {
	for (const removal of removals) {
		const page = await browser.newPage();
		await page.addInitScript(removal);
		await page.goto('/');
		await expect(page.getByText(ly97, { exact: true })).toBeVisible();
		await expect(page.locator('input, form, button')).toHaveCount(0); // "Tanpa isian masuk"
		await expect(page).toHaveTitle('Casa');
		await page.close();
	}
});
test('SEC-160 (e): a supported browser does not show the unsupported-browser sentence', async ({ page }) => {
	await page.goto('/');
	await expect(page.getByText(ly97)).toHaveCount(0);
});
```

- [ ] **Langkah 2:** `pnpm exec playwright test unsupported` → uji pertama GAGAL di kedua mesin.
- [ ] **Langkah 3:** `+layout.svelte`: `const missing = missingFeatures()` dihitung sekali saat dimuat. Bila ada yang hilang: `<main><h1>Casa</h1><p>{kalimat}</p></main>` dengan token dari `$lib/design/tokens.css`, tanpa komponen baru. Satu `h1` tetap, sehingga uji SEC-175 1.1b tetap berlaku.
- [ ] **Langkah 4:** `make web-e2e` → LULUS di `chromium` dan `webkit`, termasuk `shell.spec.ts`.
- [ ] **Langkah 5:** `ux-designer` QG-09: LY-97 di Ringkas, Sedang, Lebar (lebar 390, 744, 1280), terang dan gelap, WCAG 2.2 AA (kontras, zoom 200 %, `lang="id"`).
- [ ] **Langkah 6: commit** `Show LY-97 when a required browser feature is missing (SEC-160 (e); ADR-030 item 4)`

### Tugas 3: Basis data lokal dan penyimpanan tetap

**Berkas:**
- Buat: `web/src/lib/local-db.ts`, `web/src/lib/local-db.test.ts`
- Ubah: `web/package.json`, `web/pnpm-lock.yaml`

**Antarmuka:**
- Menghasilkan:
  - `interface OutboxEntry { idempotencyKey: string; kind: string; payload: Uint8Array; protective: boolean; createdAt: number; releaseAt: number; rejectedCode?: string }`. Komentarnya mengutip ADR-005, 09.3 §1.
  - `type MetaRow = { key: 'persistence_granted'; value: boolean }`. Union ini bertambah bersama pemakainya.
  - `class LocalDb extends Dexie` dengan nama basis data `'casa'`, versi 1, `outbox: '&idempotencyKey'` dan `meta: '&key'`. Kelasnya diekspor untuk uji dua tab.
  - `localDb(): LocalDb`, yaitu instans tunggal per tab yang dibuat malas. Membuat instans tidak membuka basis data, karena Dexie baru membuka pada operasi pertama.
  - `deleteLocalDb(): Promise<void>`: menutup instans tunggal tanpa buka ulang otomatis, menghapus `'casa'`, lalu mengosongkan instans tunggal. Dipakai Tugas 6.
  - `requestPersistence(db: LocalDb): Promise<boolean>`: memanggil `navigator.storage.persist()`, lalu menyimpan hasilnya di `meta` (SEC-171 (a)). Penolakan atau fungsi yang tidak ada dihitung `false` dan juga disimpan.

- [ ] **Langkah 1:** cek live `dexie` 4.x stabil yang berumur ≥ 7 hari dan `fake-indexeddb` stabil (lisensi, perawatan, advisori). `pnpm add -E dexie` dan `pnpm add -DE fake-indexeddb`.
- [ ] **Langkah 2: tulis uji yang gagal.** Berkas diawali `import 'fake-indexeddb/auto';`, dan `afterEach(deleteLocalDb)`.

```ts
test('SEC-165: importing the local database module creates no IndexedDB database', async () => {
	await import('./local-db');
	expect(await indexedDB.databases()).toEqual([]);
});
test('SEC-170: the local database holds only the outbox and meta stores', async () => {
	await localDb().open();
	expect(localDb().tables.map((t) => t.name).sort()).toEqual(['meta', 'outbox']);
});
test('SEC-171 (a): the persistent-storage answer is returned and kept', async () => {
	for (const [persist, expected] of [
		[() => Promise.resolve(true), true],
		[() => Promise.resolve(false), false],
		[() => Promise.reject(new Error('denied')), false],
		[undefined, false]
	] as const) {
		vi.stubGlobal('navigator', { ...navigator, storage: { persist } });
		expect(await requestPersistence(localDb())).toBe(expected);
		expect((await localDb().meta.get('persistence_granted'))?.value).toBe(expected);
	}
});
```

- [ ] **Langkah 3:** `make web-unit` → GAGAL.
- [ ] **Langkah 4:** tulis `local-db.ts` menurut Antarmuka.
- [ ] **Langkah 5:** `make web-unit` → LULUS. `make web-bundle` → LULUS. `dexie` belum masuk bundel karena belum diimpor kode aplikasi (masuk di Tugas 6).
- [ ] **Langkah 6: commit** `Add the local IndexedDB database with outbox and meta stores (10.5 §5; SEC-165, SEC-170, SEC-171 (a); ADR-024 item 7)`

### Tugas 4: Antrean dengan jeda urungkan, kunci antartab, dan pemicu kirim

**Berkas:**
- Buat: `web/src/lib/outbox.ts`, `web/src/lib/outbox.test.ts`

**Antarmuka:**
- Memakai: `LocalDb`, `OutboxEntry`, `localDb()` dari Tugas 3.
- Menghasilkan:

```ts
export type SendResult = { kind: 'accepted' } | { kind: 'retry' } | { kind: 'rejected'; code: string };
export type Send = (entry: OutboxEntry) => Promise<SendResult>;
export interface NewCommand { kind: string; payload: Uint8Array; holdSeconds: number; protective: boolean }
export function createOutbox(db: LocalDb, send: Send, clock: () => number = Date.now): {
	enqueue(command: NewCommand): Promise<OutboxEntry>;
	undo(idempotencyKey: string): Promise<boolean>;
	unsentCount(): Promise<number>;
	flush(): Promise<void>;
	start(win?: EventTarget, doc?: EventTarget & { visibilityState: DocumentVisibilityState }): () => void;
};
```

  Keputusan yang mengikat:
  - `enqueue` menulis baris dengan `releaseAt = createdAt + holdSeconds * 1000` (09.3 §1), lalu menjalankan `flush()` sebagai pemicu "sesudah setiap perintah baru". Galat tulis diteruskan apa adanya (SEC-171 (f)).
  - `undo` menghapus baris hanya bila `clock() < releaseAt`, dengan jam dibaca **di dalam** transaksi `rw`, supaya tidak bertabrakan dengan `flush` yang sedang mengambil baris yang sama.
  - `flush` berjalan di dalam `navigator.locks.request('outbox', …)` (ADR-030 butir 10). Ia mengulang sampai tidak ada baris jatuh tempo. Jatuh tempo berarti `rejectedCode` kosong dan (`releaseAt <= now` atau `createdAt > now`). Urutannya: `protective` lebih dulu, lalu `releaseAt`, lalu `createdAt` (09.3 §1 "urutan kirim"). Untuk setiap hasil: `accepted` menghapus baris; `rejected` mengisi `rejectedCode` dan baris tetap ada serta tetap dihitung belum terkirim (08 §13.2); `retry` menghentikan putaran tanpa mengubah baris. Galat dari `send` menghentikan putaran, baris tetap utuh, dan galat itu diteruskan.
  - Di akhir `flush`, satu pengatur waktu dipasang ke `releaseAt` terdekat yang belum jatuh tempo. Pengatur waktu sebelumnya dibatalkan.
  - `start` menjalankan `flush` seketika (saat dibuka), pada `online`, dan pada `visibilitychange` bila `visibilityState === 'visible'` (ADR-030 butir 10; SEC-171 (d)). Ia mengembalikan fungsi berhenti yang melepas pendengar dan pengatur waktu. Galat dari putaran yang dipicu peristiwa dilaporkan lewat `console.error(error)` tanpa isi baris.
  - `unsentCount()` = jumlah semua baris `outbox`, termasuk yang ditolak.

- [ ] **Langkah 1: tulis uji yang gagal.** Pakai `import 'fake-indexeddb/auto'`, `afterEach(deleteLocalDb)`, dan jam palsu `let t = 0; const clock = () => t;`. Perintah contoh: `const ordinary = { kind: 'test_command', payload: new Uint8Array([1]), holdSeconds: 5, protective: false }` (5 detik: 08 §9.3, 09.3 §1), `const protective = { ...ordinary, holdSeconds: 0, protective: true }`. `accepting` adalah `vi.fn<Send>()` yang menjawab `{ kind: 'accepted' }`.

```ts
test('FR-FIN-040/AC-17: a saved command is stored at once and not sent before its 5-second undo window ends', async () => {
	const outbox = createOutbox(localDb(), accepting, clock);
	await outbox.enqueue(ordinary);
	expect(await outbox.unsentCount()).toBe(1);
	t = 4999; await outbox.flush();
	expect(accepting).not.toHaveBeenCalled();
	t = 5000; await outbox.flush();
	expect(accepting).toHaveBeenCalledTimes(1);
	expect(await outbox.unsentCount()).toBe(0);
});
test('FR-FIN-040/AC-17: undo inside the undo window removes the command and nothing is sent', async () => {
	// enqueue at 0; undo at 4999 → true; flush at 5000 → send not called; unsentCount 0
});
test('FR-FIN-040/AC-17: undo after the undo window keeps the command', async () => {
	// enqueue at 0; undo at 5000 → false; unsentCount 1
});
test('FR-FIN-040/AC-17: a command saved before its tab closed is sent at the next open', async () => {
	// enqueue with db A at 0; A.close(); new LocalDb() as B; createOutbox(B, accepting, clock).flush() at 5000 → sent once
});
test('FR-FIN-040/AC-17: a device clock moved back after saving does not hold the command', async () => {
	// enqueue at 10_000; flush at 1_000 → sent once
});
test('SEC-151: a protective command is not held and is sent before older ordinary commands', async () => {
	// ordinary at 0, ordinary at 1, protective at 6000; flush at 6000 → kinds/keys in order [protective, first, second]
});
test('SEC-173 (a): two tabs flushing at once send every command exactly once', async () => {
	// two LocalDb instances, two outboxes, one shared slow send (await a 5 ms timer, records keys);
	// 3 commands at 0; t = 5000; await Promise.all([a.flush(), b.flush()]);
	// expect(sentKeys.sort()).toEqual(savedKeys.sort()); and navigator.locks.request was called with 'outbox'
});
test('FR-FIN-040/AC-18: a retry answer stops the pass and keeps every command queued', async () => {
	// 3 commands, send → retry: send called once, unsentCount 3
});
test('FR-FIN-040/AC-18: a rejected command stays queued with its code and is not sent again', async () => {
	// send → { kind: 'rejected', code: 'state_changed' }: rejectedCode 'state_changed', unsentCount 1; second flush: send not called again
});
test('FR-FIN-040/AC-18: a send that throws leaves the command queued', async () => {
	// await expect(outbox.flush()).rejects.toThrow(); unsentCount 1
});
test('SEC-171 (f): a failed write rejects the save with the original error and queues nothing', async () => {
	vi.spyOn(db.outbox, 'add').mockRejectedValue(new DOMException('full', 'QuotaExceededError'));
	await expect(outbox.enqueue(ordinary)).rejects.toMatchObject({ name: 'QuotaExceededError' });
	// unsentCount 0; send not called
});
test('the idempotency key is a UUIDv4 made at save time and kept on every attempt', async () => {
	// ADR-005, ADR-017. send → retry, then accepted: both calls get the same key, matching
	// /^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/
});
test('SEC-171 (d): the outbox sends on open, on online, and when the page becomes visible', async () => {
	// start(win, doc) with EventTargets: sends on start; on 'online'; on 'visibilitychange' with
	// visibilityState 'visible' but not 'hidden'; nothing after stop()
});
test('SEC-171 (d): a new command is sent when its undo window ends without another trigger', async () => {
	// vi.useFakeTimers({ toFake: ['setTimeout', 'clearTimeout', 'Date'] }) (fake-indexeddb needs the real setImmediate);
	// clock = Date.now; enqueue; advance 4999 ms → not sent; advance 1 ms → sent once
});
```

- [ ] **Langkah 2:** `make web-unit` → GAGAL (modul tidak ada).
- [ ] **Langkah 3:** tulis `outbox.ts` menurut keputusan di Antarmuka.
- [ ] **Langkah 4:** `make web-unit` → LULUS. Bukti kunci: bungkus `navigator.locks.request` sementara dihapus dari `flush` → uji `SEC-173 (a)` GAGAL (kunci terkirim dua kali), lalu dikembalikan. Catat di PR.
- [ ] **Langkah 5:** `grep -rnE --include='*.ts' --include='*.svelte' --exclude='*.test.ts' 'holdSeconds: *5|5000|5_000' web/src` → kosong (nilai jeda datang dari kontrak, ADR-013 butir 4).
- [ ] **Langkah 6: commit** `Add the outbox with undo window, Web Locks sender, and send triggers (ADR-030 item 10; ADR-005; SEC-171 (d), (f), SEC-173 (a); FR-FIN-040/AC-17, AC-18)`

### Tugas 5: Service worker app shell, manifest, dan service worker darurat

**Berkas:**
- Buat: `web/src/lib/app-shell.ts`, service worker (letaknya menurut langkah 1), `web/static/manifest.webmanifest`, `web/emergency/service-worker.js`, `web/src/lib/emergency-service-worker.test.ts`, `web/tests/e2e/service-worker.spec.ts`
- Ubah: `web/src/app.html` (`<link rel="manifest">`, `apple-touch-icon` `icons/icon-180.png`, dua `theme-color` dari `design-samples/index.html`), `web/vite.config.ts`, `web/tsconfig.json` (memuat `emergency`), `web/scripts/check-bundle.mjs`, `web/package.json` dan `web/pnpm-lock.yaml` (bila A)

**Antarmuka:**
- Menghasilkan: `export const APP_SHELL_CACHE_PREFIX = 'casa-shell'` di `app-shell.ts`, dipakai service worker dan Tugas 6. Setiap cache app shell bernama dengan awalan ini. Service worker menanggapi `postMessage({ type: 'SKIP_WAITING' })` dengan `skipWaiting()`. Itulah satu-satunya jalan aktif sebelum semua tab ditutup (ADR-030 butir 9; SEC-169 (c)).

- [ ] **Langkah 1: keputusan alat, sebelum kode apa pun.** Cek live `@vite-pwa/sveltekit`: versi yang mendukung `@sveltejs/kit` 3, umurnya, dan paket yang dibawanya ke kode service worker serta ke build. Pengecekan perencana pada 13 November 2026 **[Terverifikasi]** (registry.npmjs.org): hanya 2.0.0 (1 November), 2.0.1 (7 November), dan 2.0.2 (9 November 2026) yang ber-peer `@sveltejs/kit ^3.0.0`. Ketiganya memakai `@vite-pwa/workbox-swkit` 8.x (bergantung pada `idb`), bukan `workbox-*` 7, dan membawa `@unplugin-pwa/*` 0.0.x serta `verkit` 0.5. Bila hasil cek live sama, **berhenti dengan NEEDS_CONTEXT**: opsi A melanggar ADR-030 butir 8 (`@vite-pwa/workbox-swkit` dan `idb` tidak ada di daftar bundel) dan ES §8 (0.x tanpa ADR-023). Lead memilih:
  - **A.** `@vite-pwa/sveltekit` `injectManifest` sesudah amandemen ADR-030 butir 2 dan 8 (`architect`, tinjauan `security-privacy`) dan entri ADR-023. Service worker di `web/src/service-worker.ts`: `setCacheNameDetails({ prefix: APP_SHELL_CACHE_PREFIX })`, `precacheAndRoute(self.__WB_MANIFEST)`, `cleanupOutdatedCaches()`, `NavigationRoute(createHandlerBoundToURL('/index.html'))`, tanpa rute runtime. Registrasi otomatis Kit dimatikan dan registrasi dilakukan sesuai dokumentasi plugin.
  - **B.** Service worker bawaan SvelteKit 3 tanpa plugin: `web/src/service-worker/index.ts` dengan `tsconfig.json` sendiri yang meng-extend `$app/tsconfig/service-worker`; daftar dari `$app/manifest` (`immutable`, `assets`) dan `version` dari `$app/env`. `install` membuka `${APP_SHELL_CACHE_PREFIX}-${version}` dan `addAll` app shell; `activate` menghapus cache berawalan sama yang bukan versi ini; `fetch` hanya untuk GET sama-origin: navigasi dijawab `/index.html` dari cache, URL app shell dijawab dari cache, selain itu tidak memanggil `respondWith`. Tidak ada dependensi baru. Perlu amandemen ADR-030 butir 2 (alat), yang sifatnya mengurangi.

  Dalam kedua opsi: tanpa `skipWaiting()` otomatis dan tanpa `clientsClaim()`; tidak pernah meng-cache jawaban API, `/v1/stream`, foto, atau `_app/version.json`. `README.md`, `LICENSE`, dan `*.txt` di `static/` tidak masuk app shell.
- [ ] **Langkah 2: tulis uji yang gagal** (`service-worker.spec.ts`). Pembantu `cachedUrls(page)` mengembalikan semua pathname di semua cache. `expectedShell()` membaca `build/` dari sisi Node: semua berkas kecuali `service-worker.js`, `_app/version.json`, `*.map`, `README.md`, `LICENSE`, dan `*.txt`, dengan `index.html` sebagai `/index.html`.

```ts
test('SEC-169 (b): the service worker caches exactly the app shell', async ({ page }) => {
	await page.goto('/');
	await page.evaluate(() => navigator.serviceWorker.ready);
	expect((await cachedUrls(page)).sort()).toEqual(expectedShell().sort());
});
test('SEC-169 (b): API, stream, and other requests are never cached', async ({ page }) => {
	await page.goto('/');
	await page.evaluate(() => navigator.serviceWorker.ready);
	await page.evaluate(() => Promise.all([fetch('/v1/stream'), fetch(`/probe-${crypto.randomUUID()}`)]));
	expect((await cachedUrls(page)).filter((u) => u.startsWith('/v1/') || u.startsWith('/probe-'))).toEqual([]);
});
test('SEC-169 (b): after the first visit the app opens without a network', async ({ page, context }) => {
	await page.goto('/');
	await page.evaluate(() => navigator.serviceWorker.ready);
	await context.setOffline(true);
	await page.reload();
	await expect(page).toHaveTitle('Casa');
	await expect(page.locator('h1')).toHaveCount(1);
});
test('SEC-169 (c): a new service worker waits until it is told to take over', async ({ page }) => {
	// after ready: register the same script at '/service-worker.js?probe=1' (same options as the app's own
	// registration), wait for state 'installed'; expect reg.waiting not null and reg.active.scriptURL without
	// 'probe'; then reg.waiting.postMessage({ type: 'SKIP_WAITING' }); after 'controllerchange',
	// navigator.serviceWorker.controller.scriptURL contains 'probe=1'
});
test('SEC-175 (g): the web app manifest names the app Casa', async ({ request }) => {
	const manifest = await (await request.get('/manifest.webmanifest')).json();
	expect(manifest).toMatchObject({ name: 'Casa', short_name: 'Casa', display: 'standalone', start_url: '/', scope: '/' });
});
```

  `emergency-service-worker.test.ts` (Vitest, `fake-indexeddb/auto`): `self` palsu yang mengumpulkan pendengar, `caches` palsu dengan `['casa-shell-v1', 'other']`, dan `registration.unregister` sebagai `vi.fn`.

```ts
test('SEC-169 (e): the emergency service worker empties Cache Storage and unregisters without touching IndexedDB', async () => {
	// create IndexedDB 'casa' with one outbox row first; import '../../emergency/service-worker.js';
	// run the 'install' handler (skipWaiting called) and await the 'activate' waitUntil promise;
	// expect both caches deleted, unregister called once, and the 'casa' row still readable
});
```

- [ ] **Langkah 3:** `make web-unit web-e2e` → uji baru GAGAL di kedua mesin.
- [ ] **Langkah 4:** tulis service worker menurut opsi yang diputuskan, juga `manifest.webmanifest` (nilai dari `design-samples/manifest.webmanifest`, dengan `id`, `start_url`, dan `scope` `/`; ikon `/icons/…`), dan `emergency/service-worker.js`. Yang terakhir adalah JS biasa dengan `// @ts-check` dan tanpa impor, supaya tetap jalan saat build rusak; disajikan server sebagai pengganti `/service-worker.js` saat insiden (runbook, langkah 7). Isinya: `install` → `skipWaiting()`; `activate` → hapus **semua** cache (menurut SEC-169 (b), Cache Storage hanya berisi app shell), lalu `registration.unregister()`. Ia tidak menavigasi tab (menghindari putaran muat ulang selama server masih menyajikannya) dan tidak menyentuh IndexedDB. Halaman mendaftarkan service worker biasa lagi saat dibuka berikutnya. `skipWaiting` di sini disengaja dan dicatat di PR untuk dikukuhkan `security-privacy` terhadap SEC-169 (c).
- [ ] **Langkah 5: `check-bundle.mjs`** juga gagal bila build `BUNDLE_CHECK=1` tidak menghasilkan peta sumber untuk `service-worker.js`, dengan pesan `service worker source map not found, so its packages were not checked`. Dengan begitu kode service worker ikut diperiksa daftar izin.
- [ ] **Langkah 6:** `make web-unit web-e2e web-bundle` → LULUS di kedua mesin. `find web/build -name '*.map'` kosong. Catat daftar paket yang dicetak `check-bundle` di PR.
- [ ] **Langkah 7: commit** `Add app shell service worker, manifest, and emergency service worker (ADR-030 items 2, 9; SEC-169 (b), (c), (e), SEC-175 (g))`

### Tugas 6: Siaran antartab dan hapus lokal

**Berkas:**
- Buat: `web/src/lib/tab-signal.ts`, `web/src/lib/tab-signal.test.ts`, `web/src/lib/wipe.ts`, `web/src/lib/wipe.test.ts`, `web/tests/e2e/wipe.spec.ts`
- Ubah: `web/src/routes/+layout.svelte`

**Antarmuka:**
- Memakai: `localDb()`, `deleteLocalDb()` (Tugas 3); `createOutbox` (Tugas 4, di uji); `APP_SHELL_CACHE_PREFIX` (Tugas 5).
- Menghasilkan:

```ts
// tab-signal.ts — BroadcastChannel name 'casa' (ADR-030 item 10; SEC-173 (b))
export type WipeKind = 'sign_out' | 'session_revoked' | 'household_exit';
export type TabSignal = { kind: WipeKind; unsentDropped: number } | { kind: 'app_lock' };
export function postTabSignal(signal: TabSignal): void;
export function onTabSignal(handler: (signal: TabSignal) => void): () => void; // ignores anything not a TabSignal
// wipe.ts
export function wipeLocalData(): Promise<void>;
export function broadcastWipe(kind: WipeKind): Promise<number>; // returns unsentDropped
```

  Keputusan yang mengikat:
  - `wipeLocalData` menjalankan ketiga langkah dengan `Promise.allSettled`: `deleteLocalDb()`; menghapus setiap cache yang namanya tidak diawali `APP_SHELL_CACHE_PREFIX`; dan `getRegistration()` → `pushManager.getSubscription()` → `unsubscribe()` bila ada. Pencabutan langganan di server adalah tugas server (SEC-164 (b), SEC-176 (f)). Bila satu langkah gagal, langkah lain tetap berjalan, lalu fungsi menolak dengan `AggregateError` (ES §1: tidak ada yang gagal diam-diam).
  - `broadcastWipe` menghitung baris `outbox` (hanya bila `Dexie.exists('casa')`, supaya basis data tidak terbuat), lalu `postTabSignal({ kind, unsentDropped })`. `BroadcastChannel` mengantar pesan ke objek kanal lain di tab yang sama, sehingga **setiap** tab, termasuk pengirimnya, menghapus lewat satu jalan: pendengar di layout. Jumlah ini untuk kalimat LY-94 "{n} catatan yang belum terkirim dari perangkat ini ikut terhapus." (08 §25.6). Pemanggilnya (jawaban `session_revoked`, keluar akun, F06) ada di 1.9b dan langkah 6.
  - Layout: untuk `WipeKind`, `wiping = true` lebih dulu, sehingga isi dilepas dari DOM sebelum menghapus (SEC-172 (a) "sebelum menampilkan layar lain"). Lalu `wipeLocalData()`; galat dilaporkan `console.error`; akhirnya `location.replace('/')`, yang mengosongkan memori dan mengganti entri riwayat (SEC-164 (c)). Bila penghapusan gagal, pemeriksaan sesi saat dibuka (1.9b) mengulanginya karena server tetap menjawab `session_revoked`. `app_lock` belum ditangani (langkah 6).

- [ ] **Langkah 1: tulis uji yang gagal** (Vitest; `wipe.test.ts` dengan `fake-indexeddb/auto`, `caches` palsu, dan `navigator.serviceWorker` palsu lewat `vi.stubGlobal`):

```ts
test('SEC-173 (b): every signal kind reaches another channel in the same browser', async () => {
	// for each of sign_out, session_revoked, household_exit (unsentDropped 2) and app_lock:
	// onTabSignal handler receives exactly the posted object
});
test('SEC-173 (b): a malformed or unknown message is ignored', async () => {
	// raw BroadcastChannel('casa') posts 'x', { kind: 'nope' }, { kind: 'session_revoked' } (no count),
	// { kind: 'sign_out', unsentDropped: -1 }, then a valid sentinel; handler called once, with the sentinel
});
test('SEC-172 (a): the wipe deletes the local database, every cache but the app shell, and the push subscription', async () => {
	// one outbox row; caches ['casa-shell-v1', 'other']; after wipeLocalData():
	// Dexie.exists('casa') false; caches.delete called with 'other' only; unsubscribe called once
});
test('SEC-172 (a): a send that finishes after the wipe does not bring the database back', async () => {
	// flush with a send that waits on a deferred; wipeLocalData(); resolve accepted; await flush settled;
	// Dexie.exists('casa') false
});
test('SEC-172 (a): every wipe step runs when one fails, and the failure is reported', async () => {
	// caches.keys rejects: wipeLocalData rejects with AggregateError; Dexie.exists('casa') false; unsubscribe called
});
test('SEC-164 (c): broadcastWipe counts unsent commands and tells every tab', async () => {
	// two outbox rows; handler gets { kind: 'session_revoked', unsentDropped: 2 }; returns 2
});
test('SEC-172 (a): broadcastWipe without a local database creates none and reports 0', async () => {
	// expect(await broadcastWipe('sign_out')).toBe(0); indexedDB.databases() empty
});
```

  `wipe.spec.ts` (Playwright, kedua mesin):

```ts
test('SEC-172 (a): a session_revoked signal empties IndexedDB and non-shell caches in every open tab', async ({ context }) => {
	// pages A and B at '/', both after navigator.serviceWorker.ready;
	// in A: open raw IndexedDB 'casa' and close the connection (an open raw connection would block deletion);
	// caches.open('casa-probe').put('/probe', new Response('x'));
	// in A: new BroadcastChannel('casa').postMessage({ kind: 'session_revoked', unsentDropped: 0 });
	// wait for both pages to load again; in each: indexedDB.databases() has no 'casa',
	// caches.keys() has no 'casa-probe' and still has one key starting with 'casa-shell'
});
```

- [ ] **Langkah 2:** `make web-unit web-e2e` → uji baru GAGAL.
- [ ] **Langkah 3:** tulis `tab-signal.ts`, `wipe.ts`, dan pendengar di `+layout.svelte` menurut keputusan di Antarmuka. Pendengar dipasang di `$effect` dan dilepas saat dibongkar.
- [ ] **Langkah 4:** `make web-unit web-e2e web-bundle` → LULUS. `check-bundle` kini mencetak `dexie`.
- [ ] **Langkah 5: commit** `Add cross-tab signals and local wipe on sign-out, session_revoked, and F06 (ADR-030 item 10; ADR-024 item 7 (b), (d); SEC-164 (c), SEC-172 (a), SEC-173 (b))`

### Tugas 7: Bukti merah, ukuran app shell, PR

- [ ] **Langkah 1:** PR draf ke `main`; CI → HIJAU di job `server` dan `web`, termasuk langkah `web-unit (QG-05)`.
- [ ] **Langkah 2:** di cabang buangan `fs1/1.9a-redproof` (PR draf, ditutup dan dihapus sesudahnya), satu komit per gerbang baru: satu asersi Vitest dibalik → `web-unit` merah; impor paket kecil di luar daftar izin (lolos masa tunggu, tanpa skrip build, hanya di cabang buangan) ke kode service worker → `web-bundle` merah dengan nama paket itu. Catat tautan run.
- [ ] **Langkah 3: ukuran app shell** (ES §3, `[DEFAULT]` ≤ 500 KB terkompresi). `cd web && pnpm build && find build -type f ! -name '*.map' ! -name 'version.json' -print0 | xargs -0 -n1 sh -c 'gzip -9c "$0" | wc -c' | awk '{s+=$1} END {print s}'` → angka byte dicatat di PR. Bila di atas 500 000, berhenti dan laporkan.
- [ ] **Langkah 4:** deskripsi PR memuat:
  - tabel dependensi: versi, tanggal cek, umur rilis ≥ 7 hari, tautan, alasan, bundel atau pengembangan, lisensi;
  - keputusan Tugas 5 langkah 1 beserta tautan amandemen ADR;
  - bukti kunci Tugas 4 langkah 4;
  - tautan run merah;
  - ukuran app shell;
  - `skipWaiting` di service worker darurat untuk `security-privacy`.
- [ ] **Langkah 5: ES §9 butir 5.** Hanya dapat dilakukan dari origin **HTTPS**: service worker, `crypto.subtle`, Web Locks, dan `persist()` memerlukan konteks aman, sehingga `http://` di jaringan rumah akan menampilkan LY-97. Bila server uji sementara dengan TLS (14 §3.3 langkah 1, WS-5) belum ada, butir ini dinyatakan **belum diperiksa**. Yang diperiksa pemilik produk:
  - Galaxy A55, Chrome: Casa terbuka tanpa LY-97; buka sekali, mode pesawat, muat ulang → Casa tetap terbuka.
  - iPhone 13: di tab Safari, sama; lalu "Tambah ke Layar Utama", buka dari ikon, mode pesawat, tutup dan buka lagi → terbuka. Ini menguji service worker modul di Safari bila dipakai opsi B.
  - iPad mini 7: sama dengan iPhone 13.

  Hasil perangkat dilaporkan **belum terverifikasi** sampai lead mengukuhkannya. Hasil `displayContext` dan `persist()` belum terlihat tanpa LY-90 dan LY-95 (1.10).
- [ ] **Langkah 6:** `code-reviewer` dua lintasan; `ux-designer` QG-09 untuk LY-97; temuan diselesaikan; lead merge.

## Selesai bila

- CI hijau di kedua job; `web-unit` dan aturan service worker di `web-bundle` masing-masing terbukti merah sekali.
- Keputusan alat service worker (Tugas 5 langkah 1) tercatat di ADR-030 sebelum merge.
- `code-reviewer` dan `ux-designer` tanpa temuan terbuka; lead merge.
- ES §9 butir 5: hasil langkah 5 dicatat, atau dinyatakan belum diperiksa beserta sebabnya.

## Putusan lead atas temuan rencana (13 November 2026)

1. **Service worker: opsi B.** Service worker bawaan SvelteKit 3 (`src/service-worker.ts`, modul `$service-worker`: `build`, `files`, `version`), tanpa `@vite-pwa/sveltekit` dan tanpa dependensi baru. Alasan: `@vite-pwa/sveltekit` 2.x (satu-satunya yang mendukung SvelteKit 3) menarik `@vite-pwa/workbox-swkit` 8.x, `idb`, `@unplugin-pwa/*` 0.0.x, dan `verkit` 0.5, yang di luar daftar bundel ADR-030 butir 8 dan di luar ADR-023; service worker kita sudah ditulis tangan (`injectManifest`), jadi plugin tidak memberi apa pun yang diperlukan. Manifest web app ditulis sebagai berkas statis. Gerbang keputusan di Tugas 5 langkah 1 diganti putusan ini. `architect` mengamandemen ADR-030 butir 2 dan 8 (Workbox keluar dari daftar bundel).
2. **SEC-169 (b), berkas `static/` dan `index.html` tanpa hash:** diversikan lewat nama cache yang memuat `version`. Diterima lead; dikonfirmasi `security-privacy` dalam tinjauan sebelum GR-2.
3. **Deteksi fitur:** tambahkan `BroadcastChannel` dan `crypto.randomUUID` ke `missingFeatures()` (tujuh fitur). `architect` mengamandemen ADR-030 butir 4.
4. **Angka {n} LY-94 sesudah hapus:** diputuskan di 1.10.
5. **Coba ulang otomatis (08 §13.2):** 1.9a hanya memakai empat pemicu ADR-030 butir 10 ditambah pengatur waktu pelepasan. Apakah perlu pengatur waktu coba ulang dan nilainya diputuskan `architect`; bila perlu, dibangun di 1.9b bersama putaran sinkron.
6. **Perintah ditolak:** tetap di antrean dan tetap dihitung belum terkirim, seperti rencana. Cara membuangnya diputuskan bersama layar di langkah 2.
7. **Service worker darurat memanggil `skipWaiting`:** diterima lead (tidak menyentuh halaman atau IndexedDB; tujuannya pemulihan); dikonfirmasi `security-privacy` sebelum GR-2.
8. **Trusted Types dan `navigator.serviceWorker.register`:** masuk prototipe pengukur 1.11.
9. **Cek perangkat** tetap belum terbukti sampai server uji sementara dengan TLS ada (WS-5); dicatat sebagai belum terbukti, bukan lulus.
10. **Nama pengenal** `outbox`, `wipe`, `sign_out`, `household_exit` diterima; lead menambah padanan Inggrisnya di glosarium.
11. **ID AC di uji klien** (`FR-FIN-040/AC-17`, `AC-18`) dipertahankan: uji itu membuktikan bagian klien kriteria itu. Bukti penuh AC (server dan klien) ditentukan matriks bukti `qa` (14 §4.11), bukan oleh adanya uji.
12. **Laporan cakupan Vitest:** ditunda; bukan gerbang (14 §4.13).
13. **Konflik merge `Makefile` dan `ci.yml` dengan 1.2:** cabang yang di-merge kedua di-rebase oleh pelaksananya.
