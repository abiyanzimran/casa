import { Code } from '@connectrpc/connect';
import { expect, test, type Page, type Route } from '@playwright/test';
import { ErrorCode } from '../../src/lib/gen/casa/v1/error_pb';
import { HouseholdService } from '../../src/lib/gen/casa/v1/household_pb';
import { DeviceMode, IdentityService } from '../../src/lib/gen/casa/v1/identity_pb';
import { failUnmocked, mockRpc, rpcError, type RpcError } from './api-mock';
import { choiceButton, expectNavigationStateOnly, safariTab, type Choice } from './support';

// page.route does not see requests a service worker handles (plan decision 19; W2). The tests
// that count service worker registrations allow them again below.
test.use({ serviceWorkers: 'block' });

// 08 §7.6 signIn.*, common.rateLimited, §25.6 LY-02 (2), §13.2; word for word.
const rejected = 'Email atau kata sandi belum cocok. Periksa keduanya, lalu coba lagi.';
const rateLimited = 'Terlalu banyak percobaan. Tunggu beberapa menit, lalu coba lagi.';
const doNotSave = 'Jangan simpan kata sandi di browser perangkat ini.';
const needsConnection = 'Ini perlu sambungan internet. Coba lagi saat tersambung.';
const serverTrouble = 'Ada gangguan di server. Coba lagi beberapa saat lagi.';

const email = 'nama@contoh.id';
const password = 'kalimat pendek yang mudah diingat';

// 08 §25.6: the Safari-tab way into shared mode (LY-90).
const sharedTabWay = 'Pakai sebagai perangkat bersama atau pinjaman';

/**
 * The server for these tests: no session until SignIn answers with the next reply (a session
 * by default), then a session. Returns the SignIn requests.
 */
async function server(page: Page, ...replies: RpcError[]) {
	let signedIn = false;
	await failUnmocked(page);
	await mockRpc(page, HouseholdService.method.listMemberships, () =>
		signedIn ? { memberships: [] } : rpcError(Code.Unauthenticated)
	);
	return mockRpc(page, IdentityService.method.signIn, () => {
		const reply = replies.shift();
		if (reply) return reply;
		signedIn = true;
		return {};
	});
}

/** Opens Casa and goes to LY-02 Masuk with the device mode chosen (in a Safari tab: LY-90). */
async function openSignIn(page: Page, choice: Choice, tab = false): Promise<void> {
	await page.goto('/');
	await page
		.getByRole('button', { name: tab ? sharedTabWay : choiceButton[choice], exact: true })
		.click();
	await page.getByRole('button', { name: 'Sudah punya akun', exact: true }).click();
	await expect(page).toHaveURL(/\/masuk\/akun$/);
	await expect(page).toHaveTitle('Casa · Masuk');
	await expect(page.locator('h1')).toHaveText('Masuk');
	await expect(page.locator('h1')).toBeFocused();
}

const emailField = (page: Page) => page.getByLabel('Email', { exact: true });
const passwordField = (page: Page) => page.getByLabel('Kata sandi', { exact: true });
const signInButton = (page: Page) => page.getByRole('button', { name: 'Masuk', exact: true });

async function signIn(page: Page): Promise<void> {
	await emailField(page).fill(email);
	await passwordField(page).fill(password);
	await signInButton(page).click();
}

/** What a browser keeps of Casa: databases, service worker registrations, caches. */
function stored(page: Page) {
	return page.evaluate(async () => ({
		databases: (await indexedDB.databases()).map((database) => database.name),
		registrations: (await navigator.serviceWorker.getRegistrations()).length,
		caches: await caches.keys()
	}));
}

for (const choice of ['personal', 'shared'] as const) {
	test.describe(`choosing ${choice}`, () => {
		test('SEC-165 (a): sign-in sends the chosen device mode', async ({ page }) => {
			const calls = await server(page);
			await openSignIn(page, choice);
			await signIn(page);
			await expect(page).toHaveURL(/127\.0\.0\.1:\d+\/$/);
			expect(calls).toHaveLength(1);
			expect(calls[0]).toMatchObject({
				email,
				password,
				deviceMode: choice === 'personal' ? DeviceMode.PERSONAL : DeviceMode.SHARED,
				requestContext: { userInitiated: true }
			});
		});
	});
}

test('SEC-165: LY-02 before the device mode is chosen shows LY-91', async ({ page }) => {
	await server(page);
	await page.goto('/masuk/akun');
	await expect(page).toHaveURL(/\/masuk\/mode$/);
});

test.describe('with service workers allowed', () => {
	test.use({ serviceWorkers: 'allow' });

	for (const tab of [false, true]) {
		test.describe(tab ? 'in a Safari tab through LY-90' : 'through LY-91', () => {
			test('SEC-165 (b): after a shared-mode sign-in nothing is stored and no service worker is registered', async ({
				page
			}) => {
				await server(page);
				if (tab) await safariTab(page);
				await openSignIn(page, 'shared', tab);
				await signIn(page);
				await expect(page).toHaveURL(/127\.0\.0\.1:\d+\/$/);
				// Reopened, a shared session finds no database and stays shared (plan decision 5).
				await page.reload();
				await expect(page.locator('h1')).toHaveText('Casa');
				expect(await stored(page)).toEqual({ databases: [], registrations: 0, caches: [] });
				await expectNavigationStateOnly(page);
			});
		});
	}

	test('SEC-171 (a): after a personal-mode sign-in the database exists, persistence was asked and the service worker registers', async ({
		page
	}) => {
		await server(page);
		await openSignIn(page, 'personal');
		// Nothing before the sign-in (SEC-171 (e)).
		expect(await stored(page)).toEqual({ databases: [], registrations: 0, caches: [] });
		await signIn(page);
		await expect(page).toHaveURL(/127\.0\.0\.1:\d+\/$/);
		await expect.poll(async () => (await stored(page)).registrations).toBe(1);
		// The answer is stored once persist() settles, after the screen moved on: polled.
		const persistence = () =>
			page.evaluate(
				() =>
					new Promise<unknown>((resolve, reject) => {
						const open = indexedDB.open('casa');
						open.onerror = () => {
							reject(open.error ?? new Error('open failed'));
						};
						open.onsuccess = () => {
							const get = open.result
								.transaction('meta')
								.objectStore('meta')
								.get('persistence_granted');
							get.onsuccess = () => {
								open.result.close();
								resolve(get.result);
							};
						};
					})
			);
		await expect
			.poll(persistence)
			.toEqual({ key: 'persistence_granted', value: expect.any(Boolean) });
	});
});

test('SEC-175 (d): personal mode offers email and password autocomplete; shared mode turns both off and says not to save', async ({
	page
}) => {
	await server(page);
	await openSignIn(page, 'personal');
	await expect(emailField(page)).toHaveAttribute('autocomplete', 'username');
	await expect(passwordField(page)).toHaveAttribute('autocomplete', 'current-password');
	await expect(passwordField(page)).toHaveAttribute('type', 'password');
	await expect(page.getByText(doNotSave)).toHaveCount(0);

	await openSignIn(page, 'shared');
	await expect(emailField(page)).toHaveAttribute('autocomplete', 'off');
	await expect(passwordField(page)).toHaveAttribute('autocomplete', 'off');
	await expect(page.getByText(doNotSave, { exact: true })).toBeVisible();
	await expect(passwordField(page)).toHaveAccessibleDescription(doNotSave);
});

test('SEC-12: a rejected sign-in shows one message whichever of email and password is wrong', async ({
	page
}) => {
	await server(
		page,
		rpcError(Code.Unauthenticated, ErrorCode.CREDENTIALS_REJECTED),
		rpcError(Code.ResourceExhausted, ErrorCode.RATE_LIMITED)
	);
	await openSignIn(page, 'personal');
	await signIn(page);
	await expect(page.getByRole('alert')).toHaveText(rejected);
	// 08 §7.6: only the password is emptied, and only after CREDENTIALS_REJECTED.
	await expect(emailField(page)).toHaveValue(email);
	await expect(passwordField(page)).toHaveValue('');
	await expect(page).toHaveURL(/\/masuk\/akun$/);

	await passwordField(page).fill(password);
	await signInButton(page).click();
	await expect(page.getByRole('alert')).toHaveText(rateLimited);
	await expect(emailField(page)).toHaveValue(email);
	await expect(passwordField(page)).toHaveValue(password);
});

test('08 §13.2: a sign-in without a connection or with a server fault says so and keeps both fields', async ({
	page
}) => {
	await server(page, rpcError(Code.Internal));
	await openSignIn(page, 'personal');
	await signIn(page);
	await expect(page.getByRole('alert')).toHaveText(serverTrouble);
	await expect(emailField(page)).toHaveValue(email);
	await expect(passwordField(page)).toHaveValue(password);

	const offline = async (route: Route) => {
		await route.abort();
	};
	await page.route('**/casa.v1.IdentityService/SignIn', offline);
	await signInButton(page).click();
	await expect(page.getByRole('alert')).toHaveText(needsConnection);
	await expect(passwordField(page)).toHaveValue(password);
	await page.unroute('**/casa.v1.IdentityService/SignIn', offline);
	await signInButton(page).click();
	await expect(page).toHaveURL(/127\.0\.0\.1:\d+\/$/);
});

test('08 §25.5: the password field shows its text on request and never blocks paste', async ({
	page
}) => {
	await server(page);
	await openSignIn(page, 'personal');
	const show = page.getByRole('button', { name: 'Tampilkan kata sandi', exact: true });
	await expect(show).toHaveText('Tampilkan');
	await show.click();
	await expect(passwordField(page)).toHaveAttribute('type', 'text');
	const hide = page.getByRole('button', { name: 'Sembunyikan kata sandi', exact: true });
	await expect(hide).toHaveText('Sembunyikan');
	await hide.click();
	await expect(passwordField(page)).toHaveAttribute('type', 'password');
	for (const field of [emailField(page), passwordField(page)]) {
		const blocked = await field.evaluate((input) => {
			const paste = new Event('paste', { bubbles: true, cancelable: true });
			input.dispatchEvent(paste);
			return paste.defaultPrevented;
		});
		expect(blocked).toBe(false);
	}
});

test('08 §7.6: a sign-in field the server rejects by its shape is marked on that field', async ({
	page
}) => {
	// buf.validate on SignInRequest: email and password max_len (Task 5 review, minor 4).
	await server(
		page,
		rpcError(Code.InvalidArgument, undefined, [{ field: 'email', ruleId: 'string.max_len' }]),
		rpcError(Code.InvalidArgument, undefined, [{ field: 'password', ruleId: 'string.max_len' }])
	);
	await openSignIn(page, 'personal');
	await signIn(page);
	await expect(emailField(page)).toHaveAttribute('aria-invalid', 'true');
	await expect(emailField(page)).toHaveAccessibleDescription(
		'Tulis email lengkap, misalnya nama@contoh.id.'
	);
	await expect(page.getByRole('alert')).toHaveText('');

	await signInButton(page).click();
	await expect(emailField(page)).not.toHaveAttribute('aria-invalid', 'true');
	await expect(passwordField(page)).toHaveAttribute('aria-invalid', 'true');
	await expect(passwordField(page)).toHaveAccessibleDescription(
		'Kata sandi ini terlalu panjang. Pakai paling banyak 1.024 karakter.'
	);
	await expect(passwordField(page)).toHaveValue(password);
	await expect(page.getByRole('alert')).toHaveText('');
});
