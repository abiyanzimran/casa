// Link tokens in the fragment of a Casa link (SEC-175 (c); plan decision 8): read once, then
// gone from the address bar and its history entry; held in memory only (session.pendingToken).

/**
 * The token in the address bar's fragment, which leaves the address and the current history
 * entry before this returns, so no request, history entry, or Referer ever carries it
 * (SEC-175 (c)). The router's own history state is kept.
 */
export function takeFragmentToken(): string | undefined {
	const token = location.hash.slice(1);
	history.replaceState(history.state, '', location.pathname);
	return token || undefined;
}

/** The kind of token each link path of the server carries (plan decisions 8, 13). */
const pastedKinds: Record<string, 'invitation' | 'setup' | undefined> = {
	'/invite': 'invitation',
	'/signup': 'setup'
};

type PastedLink = { kind: 'invitation' | 'setup'; token: string; link: string };

/** The token alphabet of SEC-28 and SEC-158: base64url (server identity/token.go). */
const tokenChars = '[A-Za-z0-9_-]+';

/** A pasted link read with the exact rules of SEC-175 (c) (2), or undefined. */
function readLink(text: string, origin: string): PastedLink | undefined {
	let url: URL;
	try {
		url = new URL(text);
	} catch {
		return undefined;
	}
	const kind = pastedKinds[url.pathname];
	const token = url.hash.slice(1);
	if (url.origin !== origin || !kind || !token) return undefined;
	return { kind, token, link: text };
}

/**
 * The Casa link in pasted text, read without opening it (SEC-175 (c) (2) and the note "tempel
 * tautan dalam teks", 14 November 2026): the exact origin of this app, an exact link path
 * (`/invite`, `/signup`), and a fragment of token characters, found after the start of the text,
 * a space, a bracket, or a quote, so a link inside another URL is not taken; each hit is
 * re-checked with `new URL`. The same link twice is one; two different links are 'many' and
 * none is picked; none is 'not-casa'. Same origin also means https only, because the production
 * origin is https and the app does not run on a plain-http origin other than the loopback one
 * it is tested on (SEC-160 (e): no secure context, no Casa). The rest of the text is ignored.
 */
export function parsePastedLink(text: string, origin: string): PastedLink | 'not-casa' | 'many' {
	const escaped = origin.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
	const pattern = new RegExp(
		`(?<=^|[\\s(\\[<"'‘“])${escaped}/(?:invite|signup)#${tokenChars}`,
		'g'
	);
	const found = new Map<string, PastedLink>();
	for (const [hit] of text.matchAll(pattern)) {
		const link = readLink(hit, origin);
		if (link) found.set(hit, link);
	}
	if (found.size > 1) return 'many';
	return [...found.values()][0] ?? 'not-casa';
}
