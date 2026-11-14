import { Code } from '@connectrpc/connect';
import { expect, test, type Page } from '@playwright/test';
import { ErrorCode } from '../../src/lib/gen/casa/v1/error_pb';
import { HouseholdService } from '../../src/lib/gen/casa/v1/household_pb';
import { DeviceMode, IdentityService } from '../../src/lib/gen/casa/v1/identity_pb';
import { failUnmocked, mockRpc, rpcError, type RpcError } from './api-mock';
import { choiceButton, expectNavigationStateOnly, safariTab, type Choice } from './support';

// page.route does not see requests a service worker handles (plan decision 19; W2).
test.use({ serviceWorkers: 'block' });

// 08 §7.6 signUp.*, checkEmail.*, verify.*, ly90.fromSetup; §7.2 SEC-30; §25.6; word for word.
const closed =
	'Akun Casa dibuat lewat tautan: tautan undangan dari anggota keluarga, atau tautan pemasangan dari pengelola server.';
const pasteHint = 'Salin tautan dari pesan yang kamu terima, lalu tempel di sini.';
const notCasa = 'Tautan ini bukan tautan Casa. Salin lagi tautan dari pesan yang kamu terima.';
const withInvitation = 'Kamu membuka undangan Casa. Buat akun dulu, lalu undangannya tampil.';
const withSetup =
	'Tautan pemasangan siap dipakai. Pakai email yang kamu berikan ke pengelola server.';
const emailHint =
	'Pakai email yang hanya kamu yang membukanya. Kami kirim tautan untuk melanjutkan ke email ini.';
const passwordLater = 'Kata sandi dibuat sesudah kamu membuka tautan itu.';
const emailInvalid = 'Tulis email lengkap, misalnya nama@contoh.id.';
const setupRejected =
	'Tautan pemasangan ini tidak dapat dipakai. Pastikan emailnya sama dengan yang kamu berikan ke pengelola server, atau minta tautan baru kepadanya.';
const invitationRejected =
	'Undangan ini tidak dapat dipakai. Minta undangan baru kepada yang mengundangmu.';
const rateLimited = 'Terlalu banyak percobaan. Tunggu beberapa menit, lalu coba lagi.';
const sent =
	'Kami mengirim tautan ke nama@contoh.id. Buka tautan itu untuk membuat kata sandi dan melanjutkan.';
const checkHelp = 'Belum masuk dalam beberapa menit? Periksa folder spam, atau kirim ulang.';
const resent = 'Tautan baru sudah dikirim. Pakai tautan di email yang paling baru.';
const passwordHint =
	'Paling sedikit 12 karakter. Tidak perlu angka atau simbol; kalimat pendek yang mudah kamu ingat juga boleh.';
const doNotSave = 'Jangan simpan kata sandi di browser perangkat ini.';
const tooShort = 'Kata sandi paling sedikit 12 karakter.';
const tooLong = 'Kata sandi ini terlalu panjang. Pakai paling banyak 1.024 karakter.';
const common = 'Kata sandi ini terlalu umum dan mudah ditebak. Pilih yang lain.';
const verifyRejected =
	'Tautan ini sudah tidak dapat dipakai. Kalau akunmu sudah jadi, masuk saja. Kalau belum, buka lagi tautan undangan atau tautan pemasangan, lalu daftar lagi.';
const fromSetup =
	'Kamu membuka tautan pemasangan. Sesudah Casa terpasang, bukalah, pilih Buat akun, lalu tempel tautan pemasangan dari pesan yang kamu terima.';

const email = 'nama@contoh.id';
const password = 'kalimat pendek yang mudah diingat';
// The size of a server token: 32 random bytes in unpadded base64url (identity.NewToken).
const token = 'A'.repeat(43);

/**
 * The server for these tests: no session until CompleteSignUp answers with its next reply (a
 * session by default). StartSignUp and CompleteSignUp answer with their queued replies, then
 * success. urls holds the page's address at every call the app made.
 */
async function server(page: Page, replies: { start?: RpcError[]; complete?: RpcError[] } = {}) {
	const urls: string[] = [];
	let signedIn = false;
	await failUnmocked(page);
	await mockRpc(page, HouseholdService.method.listMemberships, () => {
		urls.push(page.url());
		return signedIn ? { memberships: [] } : rpcError(Code.Unauthenticated);
	});
	const starts = await mockRpc(page, IdentityService.method.startSignUp, () => {
		urls.push(page.url());
		return replies.start?.shift() ?? {};
	});
	const completes = await mockRpc(page, IdentityService.method.completeSignUp, () => {
		urls.push(page.url());
		const reply = replies.complete?.shift();
		if (reply) return reply;
		signedIn = true;
		return {};
	});
	return { urls, starts, completes };
}

const emailField = (page: Page) => page.getByLabel('Email', { exact: true });
const pasteField = (page: Page) => page.getByLabel('Tempel tautan', { exact: true });
const passwordField = (page: Page) => page.getByLabel('Kata sandi', { exact: true });
const sendLink = (page: Page) => page.getByRole('button', { name: 'Kirim tautan', exact: true });
const proceed = (page: Page) => page.getByRole('button', { name: 'Lanjut', exact: true });

/** Chooses the device mode on LY-91 and taps LY-01 "Buat akun" (08 §25.6 flow, §5.1). */
async function chooseAndSignUp(page: Page, choice: Choice = 'personal'): Promise<void> {
	await expect(page).toHaveURL(/\/masuk\/mode$/);
	await page.getByRole('button', { name: choiceButton[choice], exact: true }).click();
	await page.getByRole('button', { name: 'Buat akun', exact: true }).click();
	await expect(page).toHaveURL(/\/masuk\/daftar$/);
	await expect(page).toHaveTitle('Casa · Buat akun');
	await expect(page.locator('h1')).toHaveText('Buat akun');
	await expect(page.locator('h1')).toBeFocused();
}

/** Opens Casa from a setup link and reaches LY-02 Buat akun with the token held. */
async function openSetupLink(page: Page): Promise<void> {
	await page.goto(`/signup#${token}`);
	await chooseAndSignUp(page);
	await expect(page.getByText(withSetup, { exact: true })).toBeVisible();
}

/** Opens Casa without a link, then pastes text into the closed sign-up's field. */
async function paste(page: Page, text: string): Promise<void> {
	await pasteField(page).fill(text);
	await proceed(page).click();
}

async function openWithPastedInvitation(page: Page): Promise<void> {
	await page.goto('/');
	await chooseAndSignUp(page);
	await paste(page, `${new URL(page.url()).origin}/invite#invitation-token`);
	await expect(page.getByText(withInvitation, { exact: true })).toBeVisible();
}

/** What the browser keeps where a token could linger: address, history state, storage. */
function kept(page: Page) {
	return page.evaluate(async () => ({
		address: location.href + JSON.stringify(history.state),
		storage: JSON.stringify([Object.entries(localStorage), Object.entries(sessionStorage)]),
		databases: (await indexedDB.databases()).map((database) => database.name)
	}));
}

test('SEC-175 (c): a setup link is read from the fragment and removed from the address bar before any request', async ({
	page
}) => {
	const { urls, starts } = await server(page);
	await openSetupLink(page);
	await emailField(page).fill(email);
	await sendLink(page).click();
	await expect(page).toHaveURL(/\/masuk\/periksa-email$/);

	expect(starts).toHaveLength(1);
	expect(starts[0]?.setupToken).toBe(token);
	expect(urls.length).toBeGreaterThan(1);
	for (const url of urls) expect(url, 'the address at a request').not.toContain('#');
	const { address, storage, databases } = await kept(page);
	expect(address).not.toContain(token);
	expect(storage).not.toContain(token);
	expect(databases).toEqual([]);
	await expectNavigationStateOnly(page);
	// Every history entry of this tab: LY-02 Buat akun, LY-01, then where /signup was.
	for (const path of [/\/masuk\/daftar$/, /\/masuk$/, /\/masuk\/mode$/]) {
		await page.goBack();
		await expect(page).toHaveURL(path);
		expect(page.url()).not.toContain(token);
	}
});

test('SEC-158: StartSignUp carries the setup token or the invitation token, never both', async ({
	page
}) => {
	const { starts } = await server(page);
	await openSetupLink(page);
	await emailField(page).fill(email);
	await sendLink(page).click();
	await expect(page).toHaveURL(/\/masuk\/periksa-email$/);

	await openWithPastedInvitation(page);
	await emailField(page).fill(email);
	await sendLink(page).click();
	await expect(page).toHaveURL(/\/masuk\/periksa-email$/);

	expect(starts).toHaveLength(2);
	expect(starts[0]).toMatchObject({
		email,
		setupToken: token,
		invitationToken: '',
		requestContext: { userInitiated: true }
	});
	expect(starts[1]).toMatchObject({
		email,
		setupToken: '',
		invitationToken: 'invitation-token',
		requestContext: { userInitiated: true }
	});
});

test('SEC-158: without a held link sign-up is closed, and a pasted link is checked without a request', async ({
	page
}) => {
	const { starts } = await server(page);
	await page.goto('/');
	await chooseAndSignUp(page);
	await expect(page.getByText(closed, { exact: true })).toBeVisible();
	await expect(emailField(page)).toHaveCount(0);
	await expect(pasteField(page)).toHaveAccessibleDescription(pasteHint);
	// SEC-175 (c) (1): the pasted link is never offered for autofill.
	await expect(pasteField(page)).toHaveAttribute('autocomplete', 'off');

	const origin = new URL(page.url()).origin;
	for (const text of ['halo', 'https://evil.example/invite#t', `${origin}/undang#t`]) {
		await paste(page, text);
		await expect(pasteField(page)).toHaveAccessibleDescription(`${pasteHint} ${notCasa}`);
		await expect(pasteField(page)).toHaveAttribute('aria-invalid', 'true');
	}
	await paste(page, `  ${origin}/signup#${token} `);
	await expect(page.getByText(withSetup, { exact: true })).toBeVisible();
	await expect(pasteField(page)).toHaveCount(0);
	await expect(emailField(page)).toHaveAccessibleDescription(emailHint);
	await expect(page.getByText(passwordLater, { exact: true })).toBeVisible();
	expect(starts).toHaveLength(0);

	// "Sudah punya akun? Masuk" (08 §7.6 signUp.closed) opens LY-02 Masuk.
	await page.goto('/');
	await chooseAndSignUp(page);
	await page.getByRole('link', { name: 'Sudah punya akun? Masuk', exact: true }).click();
	await expect(page).toHaveURL(/\/masuk\/akun$/);
});

test('SEC-30: every rejected sign-up start shows one message per link type', async ({ page }) => {
	// The server gives every reason the one TOKEN_REJECTED (SEC-30); the sentence follows the kind
	// of link the client holds, and names no field (08 §7.6).
	const rejected = () => rpcError(Code.PermissionDenied, ErrorCode.TOKEN_REJECTED);
	await server(page, { start: [rejected(), rejected()] });
	await openSetupLink(page);
	await emailField(page).fill(email);
	await sendLink(page).click();
	await expect(page.getByRole('alert')).toHaveText(setupRejected);
	await expect(emailField(page)).toHaveValue(email);
	await expect(page.locator('[aria-invalid="true"]')).toHaveCount(0);

	await openWithPastedInvitation(page);
	await emailField(page).fill(email);
	await sendLink(page).click();
	await expect(page.getByRole('alert')).toHaveText(invitationRejected);
	await expect(page.locator('[aria-invalid="true"]')).toHaveCount(0);
	await expect(page).toHaveURL(/\/masuk\/daftar$/);
});

test('08 §7.6: an email the server rejects by its shape is marked on the email field', async ({
	page
}) => {
	await server(page, {
		start: [
			rpcError(Code.InvalidArgument, undefined, [{ field: 'email', ruleId: 'string.email' }]),
			rpcError(Code.ResourceExhausted, ErrorCode.RATE_LIMITED)
		]
	});
	await openSetupLink(page);
	await emailField(page).fill('nama');
	await sendLink(page).click();
	await expect(emailField(page)).toHaveAttribute('aria-invalid', 'true');
	await expect(emailField(page)).toHaveAccessibleDescription(`${emailHint} ${emailInvalid}`);
	await expect(page.getByRole('alert')).toHaveText('');

	await emailField(page).fill(email);
	await sendLink(page).click();
	await expect(page.getByRole('alert')).toHaveText(rateLimited);
	await expect(emailField(page)).not.toHaveAttribute('aria-invalid', 'true');
});

test('SEC-12: LY-03 says a link was sent without saying an account exists, resends the same request, and keeps the email to change', async ({
	page
}) => {
	const { starts } = await server(page);
	await openSetupLink(page);
	await emailField(page).fill(email);
	await sendLink(page).click();
	await expect(page).toHaveURL(/\/masuk\/periksa-email$/);
	await expect(page).toHaveTitle('Casa · Periksa email');
	await expect(page.locator('h1')).toHaveText('Periksa emailmu');
	await expect(page.locator('h1')).toBeFocused();
	await expect(page.getByText(sent, { exact: true })).toBeVisible();
	await expect(page.getByText(checkHelp, { exact: true })).toBeVisible();

	await page.getByRole('button', { name: 'Kirim ulang', exact: true }).click();
	await expect(page.getByRole('status').filter({ hasText: resent })).toHaveText(resent);
	expect(starts).toHaveLength(2);
	expect(starts[1]).toEqual(starts[0]);

	await page.getByRole('link', { name: 'Ganti email', exact: true }).click();
	await expect(page).toHaveURL(/\/masuk\/daftar$/);
	await expect(emailField(page)).toHaveValue(email);
	await expect(page.getByText(withSetup, { exact: true })).toBeVisible();
});

for (const choice of ['personal', 'shared'] as const) {
	test.describe(`choosing ${choice}`, () => {
		test('SEC-165 (a): completing sign-up asks for the device mode and a password, then signs in', async ({
			page
		}) => {
			const { urls, completes } = await server(page);
			await page.goto('/signup/verify#verification-token');
			await expect(page).toHaveURL(/\/masuk\/mode$/);
			await page.getByRole('button', { name: choiceButton[choice], exact: true }).click();
			await expect(page).toHaveURL(/\/signup\/verify$/);
			await expect(page).toHaveTitle('Casa · Buat kata sandi');
			await expect(page.locator('h1')).toHaveText('Buat kata sandi');
			await expect(page.locator('h1')).toBeFocused();
			// SEC-175 (d), SEC-165 (e): a new password is offered for saving in personal mode only.
			await expect(passwordField(page)).toHaveAttribute(
				'autocomplete',
				choice === 'personal' ? 'new-password' : 'off'
			);
			await expect(passwordField(page)).toHaveAccessibleDescription(
				choice === 'personal' ? passwordHint : `${passwordHint} ${doNotSave}`
			);
			await passwordField(page).fill(password);
			await proceed(page).click();
			await expect(page).toHaveURL(/127\.0\.0\.1:\d+\/$/);

			expect(completes).toHaveLength(1);
			expect(completes[0]).toMatchObject({
				verificationToken: 'verification-token',
				password,
				deviceMode: choice === 'personal' ? DeviceMode.PERSONAL : DeviceMode.SHARED,
				requestContext: { userInitiated: true }
			});
			for (const url of urls) expect(url, 'the address at a request').not.toContain('#');
			expect((await kept(page)).address).not.toContain('verification-token');
		});
	});
}

test('SEC-11: a too short or common password is reported on the password field from the server answer', async ({
	page
}) => {
	const violation = (ruleId: string) =>
		rpcError(Code.InvalidArgument, undefined, [{ field: 'password', ruleId }]);
	const { completes } = await server(page, {
		complete: [
			violation('string.min_len'),
			violation('password.common'),
			violation('string.max_len')
		]
	});
	await page.goto('/signup/verify#verification-token');
	await page.getByRole('button', { name: choiceButton.personal, exact: true }).click();
	for (const [typed, message] of [
		['pendek', tooShort],
		['password1234', common],
		['x'.repeat(1025), tooLong]
	] as const) {
		await passwordField(page).fill(typed);
		await proceed(page).click();
		await expect(passwordField(page)).toHaveAccessibleDescription(`${passwordHint} ${message}`);
		await expect(passwordField(page)).toHaveAttribute('aria-invalid', 'true');
		// A field error is never the alert above the button (08 §7.6), and the field keeps its text.
		await expect(page.getByRole('alert')).toHaveText('');
		await expect(passwordField(page)).toHaveValue(typed);
	}
	expect(completes).toHaveLength(3);
	// The client checks no length or list itself (plan decision 7): every try reached the server.
	expect(completes.map((request) => request.password)).toEqual([
		'pendek',
		'password1234',
		'x'.repeat(1025)
	]);
});

test('SEC-30: a rejected verification link shows one sentence and a way to sign in', async ({
	page
}) => {
	await server(page, { complete: [rpcError(Code.PermissionDenied, ErrorCode.TOKEN_REJECTED)] });
	await page.goto('/signup/verify#verification-token');
	await page.getByRole('button', { name: choiceButton.personal, exact: true }).click();
	await passwordField(page).fill(password);
	await proceed(page).click();
	await expect(page.getByRole('alert')).toHaveText(verifyRejected);
	await expect(passwordField(page)).toHaveCount(0);
	await page.getByRole('button', { name: 'Masuk', exact: true }).click();
	await expect(page).toHaveURL(/\/masuk\/akun$/);
});

test('SEC-171 (b) (2): until the contract has the no-session flag, a verification link in a Safari tab shows LY-90 and sends no token', async ({
	page
}) => {
	const { urls, completes } = await server(page);
	await safariTab(page);
	await page.goto('/signup/verify#verification-token');
	await expect(page).toHaveURL(/\/pasang$/);
	await expect(page.locator('h1')).toHaveText('Pasang Casa di layar utama dulu');
	expect(completes).toHaveLength(0);
	for (const url of urls) expect(url).not.toContain('#');
	expect((await kept(page)).address).not.toContain('verification-token');
});

test('SEC-171 (b): a setup link in a Safari tab shows LY-90 with its line and drops the token unsent', async ({
	page
}) => {
	const { urls, starts } = await server(page);
	await safariTab(page);
	await page.goto(`/signup#${token}`);
	await expect(page).toHaveURL(/\/pasang$/);
	await expect(page.getByText(fromSetup, { exact: true })).toBeVisible();
	for (const url of urls) expect(url).not.toContain('#');
	expect((await kept(page)).address).not.toContain(token);
	// The shared way in this tab finds no link held: sign-up is closed (08 §7.6 ly90.fromSetup).
	await page
		.getByRole('button', { name: 'Pakai sebagai perangkat bersama atau pinjaman', exact: true })
		.click();
	await page.getByRole('button', { name: 'Buat akun', exact: true }).click();
	await expect(page.getByText(closed, { exact: true })).toBeVisible();
	expect(starts).toHaveLength(0);
});

test('08 §7.6: LY-02 Masuk offers sign-up under its button', async ({ page }) => {
	await server(page);
	await page.goto('/');
	await page.getByRole('button', { name: choiceButton.personal, exact: true }).click();
	await page.getByRole('button', { name: 'Sudah punya akun', exact: true }).click();
	await page.getByRole('link', { name: 'Belum punya akun? Buat akun', exact: true }).click();
	await expect(page).toHaveURL(/\/masuk\/daftar$/);
	await expect(page.locator('h1')).toHaveText('Buat akun');
});
