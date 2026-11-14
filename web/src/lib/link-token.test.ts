import { expect, test, vi } from 'vitest';
import { parsePastedLink, takeFragmentToken } from './link-token';

const origin = 'https://casa.example';

test('SEC-175 (c): a pasted invitation or setup link of the production origin gives its kind and token', () => {
	expect(parsePastedLink(`${origin}/invite#tok`, origin)).toEqual({
		kind: 'invitation',
		token: 'tok',
		link: `${origin}/invite#tok`
	});
	expect(parsePastedLink(`${origin}/signup#tok`, origin)).toEqual({
		kind: 'setup',
		token: 'tok',
		link: `${origin}/signup#tok`
	});
});

test('SEC-175 (c): spaces around a pasted link are trimmed', () => {
	expect(parsePastedLink(`  ${origin}/invite#tok \n`, origin)).toEqual({
		kind: 'invitation',
		token: 'tok',
		link: `${origin}/invite#tok`
	});
});

test('SEC-175 (c) note (i): a Casa link in the middle of other text is found, and only it is returned', () => {
	for (const text of [
		`Halo, ini undangannya: ${origin}/invite#tok_1-A. Terima kasih!`,
		`(${origin}/invite#tok_1-A)`,
		`"${origin}/invite#tok_1-A"`,
		`<${origin}/invite#tok_1-A>`,
		`baris satu\n${origin}/invite#tok_1-A\nbaris tiga`
	]) {
		expect(parsePastedLink(text, origin), text).toEqual({
			kind: 'invitation',
			token: 'tok_1-A',
			link: `${origin}/invite#tok_1-A`
		});
	}
});

test('SEC-175 (c) note (iii): the same link twice is one; two different links (token or kind) are many', () => {
	const one = `${origin}/invite#a`;
	expect(parsePastedLink(`${one} lalu lagi ${one}`, origin)).toMatchObject({ token: 'a' });
	expect(parsePastedLink(`${one} dan ${origin}/invite#b`, origin)).toBe('many');
	expect(parsePastedLink(`${one} dan ${origin}/signup#a`, origin)).toBe('many');
});

test('SEC-175 (c) note (ii): a Casa link nested in another URL, a look-alike origin or path, and /signup/verify in text are not taken', () => {
	for (const text of [
		`https://lain.example/?u=${origin}/invite#tok`,
		`x${origin}/invite#tok`,
		`lihat https://casa.example.evil.example/invite#tok`,
		`lihat https://casa.example:8443/invite#tok`,
		`lihat ${origin}/invite/#tok dan ${origin}/undangan#tok`,
		`lihat ${origin}/signup/verify#tok`,
		`lihat http://casa.example/invite#tok`
	]) {
		expect(parsePastedLink(text, origin), text).toBe('not-casa');
	}
});

test('SEC-175 (c): anything but https, the exact origin, an exact link path, and a fragment is not a Casa link', () => {
	for (const text of [
		'http://casa.example/invite#tok',
		'https://www.casa.example/invite#tok',
		'https://casa.example.evil.example/invite#tok',
		'https://evil.example/invite#tok',
		'https://casa.example:8443/invite#tok',
		'https://casa.example/signup/verify#tok',
		'https://casa.example/invite/#tok',
		'https://casa.example/undangan#tok',
		'https://casa.example/invite',
		'https://casa.example/invite#',
		'halo',
		''
	]) {
		expect(parsePastedLink(text, origin), text).toBe('not-casa');
	}
});

test('SEC-175 (c): the fragment token leaves the address bar and its history entry before it is returned', () => {
	const location = { hash: '#tok', pathname: '/signup' };
	const state = { 'sveltekit:history': 1 };
	const replaceState = vi.fn((_state: unknown, _unused: string, url: string) => {
		location.hash = '';
		location.pathname = url;
	});
	vi.stubGlobal('location', location);
	vi.stubGlobal('history', { state, replaceState });

	expect(takeFragmentToken()).toBe('tok');
	// The router's own state is kept; only the address changes, in place (no new entry).
	expect(replaceState).toHaveBeenCalledExactlyOnceWith(state, '', '/signup');
	expect(location.hash).toBe('');
});

test('SEC-175 (c): an address without a fragment gives no token', () => {
	vi.stubGlobal('location', { hash: '', pathname: '/signup' });
	vi.stubGlobal('history', { state: null, replaceState: vi.fn() });
	expect(takeFragmentToken()).toBeUndefined();
});
