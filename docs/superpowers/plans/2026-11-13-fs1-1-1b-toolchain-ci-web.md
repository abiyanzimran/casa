# FS-1 1.1b — Toolchain dan CI klien web: rencana implementasi

> **Untuk agen pelaksana:** SUB-SKILL WAJIB: superpowers:subagent-driven-development (disarankan) atau superpowers:executing-plans. Langkah memakai kotak centang (`- [ ]`).

**Tujuan:** proyek SvelteKit mode SPA yang paling kecil di `web/`, dengan gerbang web QG-01 sampai QG-03 dan QG-05 (Playwright) di CI, semuanya hijau dan terbukti merah saat dilanggar.

**Arsitektur:** satu paket pnpm di `web/`. Perintah gerbang ditambahkan ke `Makefile` akar dari 1.1a (target baru, pola yang sama), dan CI mendapat job `web` yang hanya memanggil target `make`.

**Stack:** Node.js LTS, pnpm, SvelteKit (mode SPA, `@sveltejs/adapter-static`), Svelte 5, Vite, TypeScript, Prettier + `prettier-plugin-svelte`, ESLint + `typescript-eslint` + `eslint-plugin-svelte`, `svelte-check`, Playwright (Chromium, WebKit), OSV-Scanner, PMD CPD.

**Spesifikasi:** ADR-030 butir 2, 4, 8, 9; ADR-025 "Catatan keputusan web lebih dulu"; 14 §2.2 tabel klien web, §5.1; 09.4 §4, §5; 10.5 SEC-167, SEC-168, SEC-175; ES §8 (masa tunggu 7 hari).

**Pelaksana:** `web-dev`. **Peninjau:** `code-reviewer`. **Cabang:** `fs1/1.1b-web-toolchain` dari `main` sesudah 1.1a merge.

## Batasan global

- Semua batasan global 1.1a berlaku (versi dicek live, dipatok tepat, dicatat di PR; Action dipatok SHA; tanpa unggah kode ke layanan lain; `.superpowers/` tidak pernah di-commit).
- **Masa tunggu 7 hari:** `minimumReleaseAge: 10080` di `web/pnpm-workspace.yaml`, berlaku untuk semua dependensi web (lebih ketat dari ES §8, yang hanya mewajibkan untuk dependensi bundel; satu aturan lebih sederhana). Pengecualian hanya untuk perbaikan keamanan lewat `minimumReleaseAgeExclude`, dicatat di PR.
- `strictDepBuilds: true` ditulis eksplisit. Setiap paket yang punya skrip build masuk `allowBuilds` dengan nilai eksplisit; `true` hanya bila hasil skripnya benar-benar diperlukan, dan setiap entri disebut di PR untuk disetujui `code-reviewer` (ADR-030 butir 8).
- Dependensi bundel hanya dari daftar ADR-030 butir 8: `svelte`, `@sveltejs/kit` (runtime klien), `dexie`, `@connectrpc/connect`, `@connectrpc/connect-web`, `@bufbuild/protobuf`, `workbox-*` yang ditarik `@vite-pwa/sveltekit`. Tugas ini hanya memakai `svelte` dan `@sveltejs/kit`.
- `document.title` = "Casa" (ADR-030 butir 9; SEC-175 (a)).
- Ambang CPD TypeScript **70**; pengecualian: tipe kontrak TS hasil generate, fixture vektor emas hasil generate, `node_modules/`, `.svelte-kit/`, keluaran build (`code-reviewer`, 13 November 2026).

## Fokus tinjauan

1. Larangan sink HTML (SEC-167) yang tidak menangkap bentuk lain: `el["innerHTML"] = …`, `outerHTML`, `insertAdjacentHTML`, `document.write`, `{@html}`, `setTimeout("…")`, `new Function`, `eval`. Setiap bentuk punya satu contoh pelanggaran di uji merah.
2. Pemeriksaan bundel yang lolos karena peta sumber kosong atau tidak terbaca: skrip gagal bila tidak menemukan satu pun paket (mis. `svelte` harus selalu tampak).
3. Peta sumber (`.map`) yang ikut ke keluaran build produksi.
4. Uji Playwright yang hanya berjalan di Chromium: kedua mesin wajib, dan CI gagal bila salah satu tidak terpasang.
5. `pnpm install` di CI tanpa `--frozen-lockfile`.

## Yang sengaja tidak di tugas ini

- Vitest: tugas 1.9 (logika TypeScript pertama; konfigurasi tanpa uji tidak membuktikan apa pun).
- `@vite-pwa/sveltekit`, Dexie, connect-web: 1.9. Kode hasil generate TS: 1.3.
- Pemeriksaan "bundel produksi tanpa penanda staging" (SEC-179): bersama tugas yang membuat build staging (halaman pemeriksaan diri); di sini belum ada build staging.
- Layar sungguhan: 1.10. Halaman di tugas ini hanya cangkang yang diganti 1.10.

---

### Tugas 1: Proyek SvelteKit SPA

**Berkas:** Buat `web/package.json`, `web/pnpm-workspace.yaml`, `web/pnpm-lock.yaml`, `web/.nvmrc`, `web/svelte.config.js`, `web/vite.config.ts`, `web/tsconfig.json`, `web/src/app.html`, `web/src/routes/+layout.ts`, `web/src/routes/+page.svelte`.

- [ ] **Langkah 1:** cek live Node.js LTS aktif terbaru (nodejs.org) dan pnpm stabil terbaru; tulis `.nvmrc` dan `"packageManager": "pnpm@<versi>"`. Cek live versi stabil SvelteKit, Svelte, adapter-static, Vite, TypeScript yang **berumur ≥ 7 hari** dan kompatibel sebagai satu set (ES §8 "Compatible as a set"; ikuti tabel kompatibilitas SvelteKit).
- [ ] **Langkah 2:** `+layout.ts`: `export const ssr = false;` dan `export const prerender = false;`. `svelte.config.js`: `adapter-static` dengan `fallback: 'index.html'`. `app.html`: `<title>Casa</title>`, `lang="id"`, meta viewport tanpa `maximum-scale` dan tanpa `user-scalable=no`.
- [ ] **Langkah 3:** `+page.svelte` berisi `<main><h1>Casa</h1></main>` saja.
- [ ] **Langkah 4:** `vite.config.ts`: `build.sourcemap` bernilai `'hidden'` hanya bila `process.env.BUNDLE_CHECK === '1'`, selain itu `false`.
- [ ] **Langkah 5:** `pnpm install` lalu `pnpm build` → berhasil; `find build -name '*.map'` kosong.

### Tugas 2: Uji Playwright yang gagal lebih dulu

**Berkas:** Buat `web/playwright.config.ts`, `web/tests/e2e/shell.spec.ts`.

- [ ] **Langkah 1:** konfigurasi dua proyek, `chromium` dan `webkit`; `webServer` menjalankan `pnpm build && pnpm preview` pada porta tetap.
- [ ] **Langkah 2: tulis uji** (sebelum Tugas 1 langkah 2–3 diterapkan, atau dengan judul sengaja salah, sehingga terlihat gagal):

```ts
test('app shell has the Casa title and one h1', async ({ page }) => {
	await page.goto('/');
	await expect(page).toHaveTitle('Casa');
	await expect(page.locator('h1')).toHaveCount(1);
});

test('unknown path falls back to the SPA shell', async ({ page }) => {
	await page.goto('/does-not-exist');
	await expect(page).toHaveTitle('Casa');
});
```

- [ ] **Langkah 3:** `pnpm exec playwright test` → GAGAL dulu, lalu LULUS di kedua proyek sesudah cangkang benar.

### Tugas 3: Format, lint, dan pemeriksaan tipe

**Berkas:** Buat `web/.prettierrc`, `web/.prettierignore`, `web/eslint.config.js`.

- [ ] **Langkah 1:** Prettier dengan `prettier-plugin-svelte`; abaikan `build/`, `.svelte-kit/`, `pnpm-lock.yaml`.
- [ ] **Langkah 2:** ESLint flat config: `typescript-eslint` (strict, type-checked) dan `eslint-plugin-svelte` (recommended). Larangan SEC-167: `no-eval`, `no-implied-eval`, `no-new-func`, `svelte/no-at-html-tags`, dan `no-restricted-properties`/`no-restricted-syntax` untuk `innerHTML`, `outerHTML` (termasuk bentuk akses `["innerHTML"]`), `insertAdjacentHTML`, `document.write`, `document.writeln`. Pesan aturan menyebut `SEC-167`.
- [ ] **Langkah 3:** `svelte-check --fail-on-warnings`.

### Tugas 4: Kerentanan, duplikasi, dan bundel

**Berkas:** Buat `web/scripts/check-bundle.mjs`. Ubah `scripts/cpd.sh` hanya bila perlu untuk TypeScript.

- [ ] **Langkah 1: OSV-Scanner** atas `web/pnpm-lock.yaml`: biner dari rilis resmi, versi dicek live, SHA-256 diverifikasi (pola sama dengan PMD di `scripts/cpd.sh`).
- [ ] **Langkah 2: CPD** TypeScript ambang 70 atas `web/src`. Cek live apakah PMD mendukung berkas `.svelte`. Bila tidak, CPD atas `.ts` saja dan catat di PR (ADR-025 butir 8: `code-reviewer` memeriksa komponen).
- [ ] **Langkah 3: `check-bundle.mjs`** (Node stdlib saja): menjalankan build dengan `BUNDLE_CHECK=1` ke direktori sementara, membaca setiap `.map`, mengambil nama paket dari jalur `node_modules/<paket>` (termasuk paket berskop `@x/y`), lalu gagal bila (a) ada paket di luar daftar izin ADR-030 butir 8, atau (b) tidak satu paket pun ditemukan, atau (c) `svelte` tidak ditemukan. Mencetak daftar paket yang ditemukan.

### Tugas 5: `Makefile` dan job CI `web`

- [ ] **Langkah 1:** target baru di `Makefile`: `web-install` (`pnpm install --frozen-lockfile`), `web-fmt` (`prettier --check`), `web-lint` (`eslint`), `web-check` (`svelte-check`), `web-vuln`, `cpd-ts`, `web-bundle`, `web-e2e`, dan `check-web` (semuanya); ditambah `check` = `check-server` + `check-web`. Target 1.1a tidak diubah.
- [ ] **Langkah 2:** job `web` di `.github/workflows/ci.yml`: `actions/setup-node` dengan `node-version-file: web/.nvmrc`, `pnpm/action-setup` (versi dari `packageManager`), cache store pnpm, `playwright install --with-deps chromium webkit`, lalu satu langkah per target. Unggah laporan Playwright sebagai artefak saat gagal.
- [ ] **Langkah 3:** perbarui `README.md` akar: prasyarat Node dan pnpm, `make check-web`.

### Tugas 6: Bukti merah, lalu hijau

- [ ] **Langkah 1:** PR draf ke `main`; CI → HIJAU di job `server` dan `web`.
- [ ] **Langkah 2:** di cabang buangan `fs1/1.1b-redproof` (PR draf, ditutup dan dihapus sesudahnya), satu komit per gerbang: berkas tak terformat → `web-fmt` merah; setiap bentuk sink di Fokus tinjauan butir 1 → `web-lint` merah (satu run boleh memuat semuanya; laporan ESLint harus menyebut semuanya); galat tipe → `web-check` merah; blok duplikat > 70 token → `cpd-ts` merah; impor paket di luar daftar izin ke kode klien (mis. paket kecil tanpa skrip build yang lolos masa tunggu, dipasang hanya di cabang buangan) → `web-bundle` merah; judul diubah → `web-e2e` merah di kedua proyek. `web-vuln`: lokal, lockfile sementara di luar repositori dengan versi paket yang punya advisori dikenal → merah. Catat tautan run.
- [ ] **Langkah 3:** commit; deskripsi PR memuat tabel dependensi (versi, tanggal cek, umur rilis ≥ 7 hari, tautan, alasan, bundel atau pengembangan), entri `allowBuilds`, dan tautan run merah. Pesan contoh: `Add SvelteKit SPA shell and web CI gates (QG-01..03, QG-05; ADR-030 items 2, 8; SEC-167, SEC-168)`.

## Selesai bila

- CI hijau di kedua job; setiap gerbang web terbukti merah sekali.
- `code-reviewer` menyetujui entri `allowBuilds` dan tidak ada temuan terbuka; lead merge.
- ES §9 butir 5: cangkang dibuka di browser PC (Chromium dan Firefox atau Safari bila ada) dan judul "Casa" terlihat; dicatat di PR.
