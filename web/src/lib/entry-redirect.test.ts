import { beforeEach, expect, test, vi } from 'vitest';
import {
	entryRedirect,
	nextScreen,
	session,
	type DeviceMode,
	type EntryScreen,
	type SignInScreen
} from './session.svelte';

type Path = EntryScreen | SignInScreen;
type Context = 'safari-tab' | 'other';
type Expected = Path | '/' | undefined;

const paths: Path[] = [
	'/pasang',
	'/masuk/mode',
	'/masuk',
	'/masuk/akun',
	'/masuk/daftar',
	'/masuk/periksa-email'
];

/**
 * Signed out, what each path does (undefined: it shows), per display context and device mode
 * (08 §25.6 flow): LY-90 only in a Safari tab, LY-91 never there, LY-01 and LY-02 only with a
 * mode, an earlier step still reachable.
 */
const signedOut: Record<Context, Record<DeviceMode | 'none', Record<Path, Expected>>> = {
	'safari-tab': {
		none: {
			'/pasang': undefined,
			'/masuk/mode': '/pasang',
			'/masuk': '/pasang',
			'/masuk/akun': '/pasang',
			'/masuk/daftar': '/pasang',
			'/masuk/periksa-email': '/pasang'
		},
		personal: {
			'/pasang': undefined,
			'/masuk/mode': '/pasang',
			'/masuk': '/pasang',
			'/masuk/akun': '/pasang',
			'/masuk/daftar': '/pasang',
			'/masuk/periksa-email': '/pasang'
		},
		shared: {
			'/pasang': undefined,
			'/masuk/mode': '/masuk',
			'/masuk': undefined,
			'/masuk/akun': undefined,
			'/masuk/daftar': undefined,
			'/masuk/periksa-email': undefined
		}
	},
	other: {
		none: {
			'/pasang': '/masuk/mode',
			'/masuk/mode': undefined,
			'/masuk': '/masuk/mode',
			'/masuk/akun': '/masuk/mode',
			'/masuk/daftar': '/masuk/mode',
			'/masuk/periksa-email': '/masuk/mode'
		},
		personal: {
			'/pasang': '/masuk',
			'/masuk/mode': undefined,
			'/masuk': undefined,
			'/masuk/akun': undefined,
			'/masuk/daftar': undefined,
			'/masuk/periksa-email': undefined
		},
		shared: {
			'/pasang': '/masuk',
			'/masuk/mode': undefined,
			'/masuk': undefined,
			'/masuk/akun': undefined,
			'/masuk/daftar': undefined,
			'/masuk/periksa-email': undefined
		}
	}
};

const rows = (['safari-tab', 'other'] as const).flatMap((context) =>
	(['none', 'personal', 'shared'] as const).flatMap((mode) =>
		paths.flatMap((path) => [
			{
				state: 'signed-out' as const,
				context,
				mode,
				path,
				expected: signedOut[context][mode][path]
			},
			// Once signed in, every entry and sign-in screen sends the person on to "/".
			{ state: 'signed-in' as const, context, mode, path, expected: '/' as const }
		])
	)
);

beforeEach(() => {
	session.state = { kind: 'checking' };
	session.mode = undefined;
});

test.each(rows)(
	'SEC-171 (b), SEC-165: $state in $context with mode $mode, $path goes to $expected',
	({ state, context, mode, path, expected }) => {
		// navigator.standalone exists only in Safari on iPhone and iPad; false there is a tab.
		vi.stubGlobal('navigator', context === 'safari-tab' ? { standalone: false } : {});
		session.state =
			state === 'signed-in' ? { kind: 'signed-in', memberships: [] } : { kind: state };
		session.mode = mode === 'none' ? undefined : mode;
		expect(entryRedirect(path)).toBe(expected);
		// Where a redirect goes is a screen that shows, so no redirect loops.
		const next = nextScreen();
		if (next === '/') expect(state).toBe('signed-in');
		else expect(entryRedirect(next)).toBeUndefined();
	}
);
