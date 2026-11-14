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

/**
 * A pasted Casa link's kind and token, read without opening the link (SEC-175 (c) (2)): the
 * exact origin of this app, an exact link path, and a fragment. Same origin also means https
 * only, because the production origin is https and the app does not run on a plain-http origin
 * other than the loopback one it is tested on (SEC-160 (e): no secure context, no Casa).
 */
export function parsePastedLink(
	text: string,
	origin: string
): { kind: 'invitation' | 'setup'; token: string } | 'not-casa' {
	let url: URL;
	try {
		url = new URL(text.trim());
	} catch {
		return 'not-casa';
	}
	const kind = pastedKinds[url.pathname];
	const token = url.hash.slice(1);
	if (url.origin !== origin || !kind || !token) return 'not-casa';
	return { kind, token };
}
