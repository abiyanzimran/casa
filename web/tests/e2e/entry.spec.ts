import { Code } from '@connectrpc/connect';
import { expect, test, type Page } from '@playwright/test';
import { HouseholdService } from '../../src/lib/gen/casa/v1/household_pb';
import { failUnmocked, mockRpc, rpcError } from './api-mock';
import { choiceButton, safariTab, visitedPaths } from './support';

// page.route does not see requests a service worker handles (plan decision 19; W2).
test.use({ serviceWorkers: 'block' });

// 08 §25.6 LY-90, word for word.
const ly90 = {
	firstLine: 'Kalau kamu membuka halaman ini dari aplikasi lain, buka halaman ini di Safari dulu.',
	title: 'Pasang Casa di layar utama dulu',
	reason:
		'Di iPhone dan iPad, Casa dibuka dari ikonnya di layar utama. Dengan begitu catatanmu tidak terhapus oleh aturan pembersihan data Safari, dan kamu bisa menerima pemberitahuan.',
	steps: [
		'1. Ketuk ⋯ di samping alamat situs, lalu ketuk Bagikan.',
		'2. Pilih Tambah ke Layar Utama.',
		'3. Biarkan Buka sebagai Web App menyala, lalu ketuk Tambah.',
		'4. Tutup Safari, lalu buka Casa dari layar utama.'
	],
	shared: 'Pakai sebagai perangkat bersama atau pinjaman',
	sharedNote:
		'Untuk iPhone atau iPad milik orang lain. Catatanmu tidak disimpan di perangkat ini, dan kamu keluar otomatis sesudah 30 menit tidak dipakai.'
};

// 08 §25.6 LY-91, word for word (the two buttons: choiceButton).
const ly91 = {
	question: 'Ini perangkat pribadimu, atau perangkat bersama atau pinjaman?',
	personalNote:
		'Salinan catatan tersimpan di browser perangkat ini, jadi kamu bisa mencatat tanpa internet. Siapa pun yang bisa membuka browser ini dapat membaca salinan itu. Bisa memakai pemberitahuan dan kunci aplikasi.',
	sharedNote:
		'Catatanmu tidak disimpan di perangkat ini. Kamu keluar otomatis sesudah 30 menit tidak dipakai, dan paling lama 12 jam. Mencatat perlu internet.',
	limit:
		'Casa tidak bisa mencegah perekam layar atau ketikan, atau pengelola kata sandi, yang terpasang di perangkat ini.'
};

// 08 §7.6 ly01.* and §5.1.
const ly01Line = 'Uang dan urusan rumah tangga, diketahui bersama tanpa perlu saling bertanya.';

const listMemberships = HouseholdService.method.listMemberships;

/** No session: the session check answers Unauthenticated without a detail (SEC-161 (c) (5)). */
async function signedOut(page: Page) {
	await failUnmocked(page);
	return mockRpc(page, listMemberships, () => rpcError(Code.Unauthenticated));
}

async function expectLy01(page: Page): Promise<void> {
	await expect(page).toHaveURL(/\/masuk$/);
	await expect(page).toHaveTitle('Casa · Selamat datang');
	await expect(page.locator('h1')).toHaveText('Casa');
	await expect(page.locator('h1')).toBeFocused();
	await expect(page.getByText(ly01Line, { exact: true })).toBeVisible();
	await expect(page.getByRole('button', { name: 'Buat akun', exact: true })).toBeVisible();
	await expect(page.getByRole('button', { name: 'Sudah punya akun', exact: true })).toBeVisible();
}

test('SEC-171 (b): a Safari tab on iPhone or iPad shows LY-90 with both ways and never offers personal mode', async ({
	page
}) => {
	const checks = await signedOut(page);
	await safariTab(page);
	await page.goto('/');
	await expect(page).toHaveURL(/\/pasang$/);
	await expect(page).toHaveTitle('Casa · Pasang');
	const heading = page.locator('h1');
	await expect(heading).toHaveText(ly90.title);
	await expect(heading).toBeFocused();
	await expect(heading).toHaveAccessibleDescription(`${ly90.firstLine} ${ly90.reason}`);
	for (const text of [ly90.firstLine, ly90.reason, ...ly90.steps]) {
		await expect(page.getByText(text, { exact: true })).toBeVisible();
	}
	// The first line stands above both ways, the install way's title included.
	const firstLineBeforeTitle = await page
		.getByText(ly90.firstLine, { exact: true })
		.evaluate(
			(line, title) => !!(line.compareDocumentPosition(title) & Node.DOCUMENT_POSITION_FOLLOWING),
			await heading.elementHandle()
		);
	expect(firstLineBeforeTitle).toBe(true);
	await expect(page.locator('ol > li')).toHaveText(ly90.steps);
	const shared = page.getByRole('button', { name: ly90.shared, exact: true });
	await expect(shared).toBeVisible();
	await expect(shared).toHaveAccessibleDescription(ly90.sharedNote);
	// "Yang tidak ada": no sign-in field, no data, no personal mode.
	await expect(page.getByRole('button')).toHaveCount(1);
	await expect(page.locator('input, select, textarea, form')).toHaveCount(0);
	await expect(page.getByText(choiceButton.personal)).toHaveCount(0);
	// The session check is the only RPC (SEC-172 (b)); LY-90 sends nothing.
	expect(checks).toHaveLength(1);
	expect(checks[0]?.requestContext?.userInitiated).toBe(false);
});

test('SEC-171 (b): the shared way in a Safari tab opens sign-in in shared mode without LY-91', async ({
	page
}) => {
	const checks = await signedOut(page);
	await safariTab(page);
	const paths = visitedPaths(page);
	await page.goto('/');
	await page.getByRole('button', { name: ly90.shared, exact: true }).click();
	await expectLy01(page);
	expect(paths).not.toContain('/masuk/mode');
	await expect(page.getByText(ly91.question)).toHaveCount(0);
	expect(checks).toHaveLength(1);
});

test('SEC-171 (b): an uncertain display context shows LY-91, not LY-90', async ({ page }) => {
	await signedOut(page);
	// Neither navigator.standalone nor the standalone display mode: an ordinary browser.
	await page.goto('/');
	await expect(page).toHaveURL(/\/masuk\/mode$/);
	await expect(page).toHaveTitle('Casa · Mode perangkat');
	await expect(page.locator('h1')).toHaveText(ly91.question);
	await expect(page.getByText(ly90.title)).toHaveCount(0);
	await expect(page.getByText(ly90.firstLine)).toHaveCount(0);
});

test('SEC-171 (b): entry screens cannot be skipped by typing their address', async ({ page }) => {
	await signedOut(page);
	await safariTab(page);
	// LY-91 offers personal mode, which a Safari tab never offers.
	await page.goto('/masuk/mode');
	await expect(page).toHaveURL(/\/pasang$/);
	await expect(page.getByText(choiceButton.personal)).toHaveCount(0);
	// LY-01 before a way was chosen.
	await page.goto('/masuk');
	await expect(page).toHaveURL(/\/pasang$/);
});

test('SEC-165: LY-01 before the device mode is chosen shows LY-91', async ({ page }) => {
	await signedOut(page);
	await page.goto('/masuk');
	await expect(page).toHaveURL(/\/masuk\/mode$/);
	await expect(page.locator('h1')).toHaveText(ly91.question);
});

for (const choice of [choiceButton.personal, choiceButton.shared]) {
	test.describe(`choosing ${choice}`, () => {
		test('SEC-165: LY-91 has no preselected choice and states what the app cannot prevent', async ({
			page
		}) => {
			const checks = await signedOut(page);
			await page.goto('/');
			await expect(page.locator('h1')).toHaveText(ly91.question);
			await expect(page.getByText(ly91.limit, { exact: true })).toBeVisible();
			const personal = page.getByRole('button', { name: choiceButton.personal, exact: true });
			const shared = page.getByRole('button', { name: choiceButton.shared, exact: true });
			await expect(personal).toHaveAccessibleDescription(ly91.personalNote);
			await expect(shared).toHaveAccessibleDescription(ly91.sharedNote);
			await expect(page.getByRole('button')).toHaveCount(2);
			for (const button of [personal, shared]) {
				await expect(button).not.toHaveAttribute('aria-pressed');
				await expect(button).not.toHaveAttribute('aria-checked');
				await expect(button).not.toHaveAttribute('aria-selected');
				await expect(button).not.toBeFocused();
			}
			// Both equally wide, the column's width (08 §25.6 "Bentuk").
			const [a, b] = [await personal.boundingBox(), await shared.boundingBox()];
			expect(a?.width).toBe(b?.width);
			await page.getByRole('button', { name: choice, exact: true }).click();
			await expectLy01(page);
			expect(checks).toHaveLength(1);
		});
	});
}

const back = (page: Page) => page.getByRole('link', { name: 'Kembali', exact: true });

/** The Kembali link is the first element in main, right before the h1 (08 §25.4). */
async function expectBackFirst(page: Page): Promise<void> {
	const first = await page.evaluate(() => {
		const main = document.querySelector('main');
		return [main?.firstElementChild?.textContent.trim(), main?.children[1].tagName];
	});
	expect(first).toEqual(['Kembali', 'H1']);
}

test('08 §25.4: LY-01 to LY-03 start with Kembali to the step before; LY-90 and LY-91 have none', async ({
	page
}) => {
	await signedOut(page);
	await page.goto('/');
	await expect(page.locator('h1')).toHaveText(ly91.question);
	await expect(back(page)).toHaveCount(0);
	await page.getByRole('button', { name: choiceButton.personal, exact: true }).click();
	await expectLy01(page);
	await expectBackFirst(page);

	await page.getByRole('button', { name: 'Sudah punya akun', exact: true }).click();
	await expect(page).toHaveURL(/\/masuk\/akun$/);
	await expectBackFirst(page);
	await back(page).click();
	await expectLy01(page);

	await page.getByRole('button', { name: 'Buat akun', exact: true }).click();
	await expect(page).toHaveURL(/\/masuk\/daftar$/);
	await expectBackFirst(page);
	// A pasted invitation is held across Kembali; from LY-91 it goes on to LY-01, not Buat akun
	// (PR B review M7).
	await page
		.getByLabel('Tempel tautan', { exact: true })
		.fill(`${new URL(page.url()).origin}/invite#invitation-token`);
	await page.getByRole('button', { name: 'Lanjut', exact: true }).click();
	await back(page).click();
	await expectLy01(page);
	await back(page).click();
	await expect(page).toHaveURL(/\/masuk\/mode$/);
	await page.getByRole('button', { name: choiceButton.shared, exact: true }).click();
	await expectLy01(page);
	await page.getByRole('button', { name: 'Buat akun', exact: true }).click();
	await expect(
		page.getByText('Kamu membuka undangan Casa. Buat akun dulu, lalu undangannya tampil.', {
			exact: true
		})
	).toBeVisible();
});

test('08 §25.4: in a Safari tab, Kembali on LY-01 goes back to LY-90, which has none', async ({
	page
}) => {
	await signedOut(page);
	await safariTab(page);
	await page.goto('/');
	await expect(back(page)).toHaveCount(0);
	await page.getByRole('button', { name: ly90.shared, exact: true }).click();
	await expectLy01(page);
	await back(page).click();
	await expect(page).toHaveURL(/\/pasang$/);
	await expect(page.locator('h1')).toHaveText(ly90.title);
});
