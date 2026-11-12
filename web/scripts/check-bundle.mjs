// Bundle gate (ADR-030 item 8; SEC-168): only allowlisted packages may reach
// code that runs in the browser, and no source map ships in the production
// build. Node standard library only.
//
// 1. The production build (`vite build`, into build/) has no .map file and no
//    sourceMappingURL comment.
// 2. A second build with BUNDLE_CHECK=1 writes hidden source maps into a
//    temporary directory; every node_modules package named in them must be on
//    the allowlist, and the list must not be empty and must contain svelte, so
//    an unreadable or empty map cannot pass.
// 3. static/ holds no script or WebAssembly file (SEC-168 (d)).
import { spawnSync } from 'node:child_process';
import { existsSync, mkdtempSync, readFileSync, readdirSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';

// ADR-030 item 8. Additions need a security-privacy review and an amendment
// of that item first.
const allowed = new Set([
	'svelte',
	'@sveltejs/kit',
	'dexie',
	'@connectrpc/connect',
	'@connectrpc/connect-web',
	'@bufbuild/protobuf',
	// Imported by the @sveltejs/kit client runtime itself (serializes load
	// data); covered by the "@sveltejs/kit (runtime klien)" entry (ADR-030
	// item 8, lead's note of 13 November 2026). Not for Casa code to import.
	'devalue'
]);
const allowedPrefixes = ['workbox-']; // pulled in by @vite-pwa/sveltekit

/** @param {string} name */
const isAllowed = (name) => allowed.has(name) || allowedPrefixes.some((p) => name.startsWith(p));

/** @param {Record<string, string>} env */
function build(env) {
	const result = spawnSync('pnpm', ['exec', 'vite', 'build'], {
		env: { ...process.env, ...env },
		stdio: ['ignore', 'ignore', 'inherit']
	});
	if (result.status !== 0) {
		throw new Error(`vite build failed (exit ${String(result.status)})`);
	}
}

/** @param {string} dir */
function files(dir) {
	return readdirSync(dir, { recursive: true, withFileTypes: true })
		.filter((entry) => entry.isFile())
		.map((entry) => join(entry.parentPath, entry.name));
}

// The package is the last node_modules segment, so pnpm's
// node_modules/.pnpm/<id>/node_modules/<package> layout resolves correctly.
/** @param {string} source */
function packageOf(source) {
	const names = [...source.matchAll(/node_modules\/((?:@[^/]+\/)?[^/]+)/g)].map((m) => m[1]);
	return names.at(-1);
}

/** @type {string[]} */
const problems = [];

// Files in static/ are copied into the build as they are, so a script there,
// loaded through <script src> in app.html, would bypass the source map check.
if (existsSync('static')) {
	for (const file of files('static')) {
		if (/\.(js|mjs|cjs|wasm)$/.test(file)) {
			problems.push(`SEC-168 (d): code in static/ bypasses the bundle allowlist: ${file}`);
		}
	}
}

delete process.env.BUNDLE_CHECK;
delete process.env.BUNDLE_CHECK_OUT;
build({});
for (const file of files('build')) {
	if (file.endsWith('.map')) problems.push(`source map in the production build: ${file}`);
	else if (/\.(js|css)$/.test(file) && readFileSync(file, 'utf8').includes('sourceMappingURL')) {
		problems.push(`sourceMappingURL in the production build: ${file}`);
	}
}

const checkDir = mkdtempSync(join(tmpdir(), 'casa-bundle-check-'));
/** @type {Set<string>} */
const packages = new Set();
try {
	build({ BUNDLE_CHECK: '1', BUNDLE_CHECK_OUT: checkDir });
	const maps = files(checkDir).filter((file) => file.endsWith('.map'));
	if (maps.length === 0) problems.push('no source maps were written, so nothing was checked');
	for (const map of maps) {
		/** @type {unknown} */
		const parsed = JSON.parse(readFileSync(map, 'utf8'));
		const sources =
			typeof parsed === 'object' && parsed !== null && 'sources' in parsed ? parsed.sources : null;
		if (!Array.isArray(sources)) {
			problems.push(`source map without a sources list: ${map}`);
			continue;
		}
		for (const source of sources) {
			const name = typeof source === 'string' ? packageOf(source) : undefined;
			if (name) packages.add(name);
		}
	}
} finally {
	rmSync(checkDir, { recursive: true, force: true });
}

const found = [...packages].sort();
console.log(`Packages in the client bundle: ${found.join(', ') || '(none)'}`);
if (found.length === 0) problems.push('no package found in the source maps');
if (!packages.has('svelte')) problems.push('svelte not found in the source maps');
for (const name of found.filter((n) => !isAllowed(n))) {
	problems.push(`${name} is not on the ADR-030 item 8 bundle allowlist`);
}

if (problems.length > 0) {
	for (const problem of problems) console.error(`check-bundle: ${problem}`);
	process.exit(1);
}
