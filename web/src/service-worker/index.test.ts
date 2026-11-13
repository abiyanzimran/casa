import { expect, test, vi } from 'vitest';

vi.mock('$app/env', () => ({ version: 'v42' }));
vi.mock('$app/manifest', () => ({
	immutable: [{ path: '_app/immutable/start.js' }],
	assets: [{ path: 'icons/icon-180.png' }, { path: 'LICENSE' }]
}));
const listeners = new Map<string, (event: unknown) => void>();
vi.mock('$app/service-worker', () => ({
	self: {
		addEventListener: (type: string, listener: (event: unknown) => void) =>
			listeners.set(type, listener)
	}
}));

/** Node's Request needs an absolute URL; in the worker a path resolves against its origin. */
class OriginRequest extends Request {
	constructor(input: string, init?: RequestInit) {
		super(new URL(input, 'https://casa.test'), init);
	}
}

test('SEC-169 (b): install fills one cache named with the version, every file fetched past the HTTP cache', async () => {
	const addAll = vi.fn<(requests: Request[]) => Promise<void>>(() => Promise.resolve());
	const open = vi.fn(() => Promise.resolve({ addAll }));
	vi.stubGlobal('caches', { open });
	vi.stubGlobal('Request', OriginRequest);
	await import('./index');

	const pending: Promise<unknown>[] = [];
	listeners.get('install')?.({ waitUntil: (promise: Promise<unknown>) => pending.push(promise) });
	await Promise.all(pending);

	expect(open.mock.calls).toEqual([['casa-shell-v42']]);
	const requests = addAll.mock.calls.flatMap(([batch]) => batch);
	expect(requests.map((request) => new URL(request.url).pathname).sort()).toEqual([
		'/',
		'/_app/immutable/start.js',
		'/icons/icon-180.png'
	]);
	expect(requests.map((request) => request.cache)).toEqual(['reload', 'reload', 'reload']);
});
