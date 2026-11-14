import 'fake-indexeddb/auto';
import { Code, ConnectError } from '@connectrpc/connect';
import { beforeEach, expect, test, vi } from 'vitest';
import { ErrorCode, ErrorDetailSchema } from './gen/casa/v1/error_pb';

let updateModule: typeof import('./app-update.svelte');

beforeEach(async () => {
	vi.resetModules();
	updateModule = await import('./app-update.svelte');
});

/** A browser whose registration has `waiting` (or not); controllerchange fires on demand. */
function browser(waiting: boolean, registered = true, controlled = true) {
	const worker = { postMessage: vi.fn() };
	const registration = {
		update: vi.fn(() => Promise.resolve()),
		waiting: waiting ? worker : null
	};
	let changed = () => {};
	const reload = vi.fn();
	vi.stubGlobal('navigator', {
		serviceWorker: {
			controller: controlled ? {} : null,
			getRegistration: () => Promise.resolve(registered ? registration : undefined),
			addEventListener: (_type: string, listener: () => void) => {
				changed = listener;
			}
		}
	});
	vi.stubGlobal('location', { reload });
	return {
		worker,
		registration,
		reload,
		fireControllerChange: () => {
			changed();
		}
	};
}

const upgradeRequired = new ConnectError('old', Code.FailedPrecondition, undefined, [
	{ desc: ErrorDetailSchema, value: { code: ErrorCode.CLIENT_UPGRADE_REQUIRED } }
]);

test('SEC-169 (c): a waiting version is applied only without an open form (none waiting)', async () => {
	const b = browser(false);
	expect(await updateModule.applyUpdate(() => false)).toBe('none');
	expect(b.registration.update).toHaveBeenCalledTimes(1);
	expect(b.worker.postMessage).not.toHaveBeenCalled();
});

test('SEC-169 (c): a waiting version is applied only without an open form (no open form)', async () => {
	const b = browser(true);
	const applied = updateModule.applyUpdate(() => false);
	await vi.waitFor(() => {
		expect(b.worker.postMessage).toHaveBeenCalledWith({ type: 'SKIP_WAITING' });
	});
	expect(b.reload).not.toHaveBeenCalled();
	b.fireControllerChange();
	expect(await applied).toBe('reloaded');
	expect(b.reload).toHaveBeenCalledTimes(1);
});

test('SEC-169 (c): a waiting version is applied only without an open form (open form)', async () => {
	const b = browser(true);
	expect(await updateModule.applyUpdate(() => true)).toBe('deferred');
	expect(b.worker.postMessage).not.toHaveBeenCalled();
});

test('SEC-169 (c): a deferred update is applied when the last open form closes', async () => {
	const b = browser(true);
	const closeFirst = updateModule.openForm();
	const closeSecond = updateModule.openForm();
	expect(await updateModule.applyUpdate(updateModule.hasOpenForms)).toBe('deferred');
	closeFirst();
	await new Promise((resolve) => setTimeout(resolve, 0));
	// Closing one of two forms does not look for the version again.
	expect(b.registration.update).toHaveBeenCalledTimes(1);
	closeSecond();
	await vi.waitFor(() => {
		expect(b.worker.postMessage).toHaveBeenCalledWith({ type: 'SKIP_WAITING' });
	});
});

test('SEC-169 (d): CLIENT_UPGRADE_REQUIRED sets upgrade.required, other failures do not', () => {
	updateModule.noteUpgradeRequired(new ConnectError('x', Code.Unavailable));
	expect(updateModule.upgrade.required).toBe(false);
	updateModule.noteUpgradeRequired(upgradeRequired);
	expect(updateModule.upgrade.required).toBe(true);
});

test('SEC-169 (d): the session check answered CLIENT_UPGRADE_REQUIRED sets upgrade.required', async () => {
	const api = await import('./api');
	vi.spyOn(api.api.household, 'listMemberships').mockRejectedValue(upgradeRequired);
	vi.stubGlobal('window', { addEventListener: vi.fn() });
	vi.stubGlobal('document', { addEventListener: vi.fn() });
	vi.spyOn(console, 'error').mockImplementation(() => {});
	const session = await import('./session.svelte');
	await session.checkSession();
	expect(updateModule.upgrade.required).toBe(true);
});

test('08 §25.7: Perbarui with no registration reloads, unless a form is open', async () => {
	const b = browser(false, false);
	await updateModule.updateNow();
	expect(b.reload).toHaveBeenCalledTimes(1);
	// The reloaded page (fresh modules), now with a form open.
	vi.resetModules();
	updateModule = await import('./app-update.svelte');
	const close = updateModule.openForm();
	await updateModule.updateNow();
	expect(b.reload).toHaveBeenCalledTimes(1);
	close();
	await vi.waitFor(() => {
		expect(b.reload).toHaveBeenCalledTimes(2);
	});
});

test('SEC-169 (c): a page no worker controls reloads without waiting for controllerchange', async () => {
	const b = browser(true, true, false);
	expect(await updateModule.applyUpdate(() => false)).toBe('reloaded');
	expect(b.worker.postMessage).toHaveBeenCalledWith({ type: 'SKIP_WAITING' });
	expect(b.reload).toHaveBeenCalledTimes(1);
});

test('08 §25.7: a second tap while an update runs does nothing', async () => {
	const b = browser(true);
	const first = updateModule.updateNow();
	await updateModule.updateNow();
	await vi.waitFor(() => {
		expect(b.worker.postMessage).toHaveBeenCalledTimes(1);
	});
	b.fireControllerChange();
	await first;
	expect(b.registration.update).toHaveBeenCalledTimes(1);
});

test('08 §25.7: the update row says the second sentence only signed in on a personal device', () => {
	expect(updateModule.updateRow(true)).toBeUndefined(); // no row before CLIENT_UPGRADE_REQUIRED
	updateModule.noteUpgradeRequired(upgradeRequired);
	expect(updateModule.updateRow(true)).toEqual({
		text: 'Aplikasi perlu diperbarui untuk mengirim catatan. Catatanmu tersimpan di perangkat ini.',
		button: { label: 'Perbarui', busy: false }
	});
	expect(updateModule.updateRow(false)).toEqual({
		text: 'Aplikasi perlu diperbarui untuk mengirim catatan.',
		button: { label: 'Perbarui', busy: false }
	});
});

test('08 §25.7: while Perbarui looks for a version its button is busy and says Memeriksa…', async () => {
	const b = browser(false);
	let finish: () => void = () => undefined;
	b.registration.update.mockImplementationOnce(
		() =>
			new Promise<void>((resolve) => {
				finish = resolve;
			})
	);
	updateModule.noteUpgradeRequired(upgradeRequired);
	const updating = updateModule.updateNow();
	await vi.waitFor(() => {
		expect(b.registration.update).toHaveBeenCalled();
	});
	expect(updateModule.updateRow(true)?.button).toEqual({ label: 'Memeriksa…', busy: true });
	finish();
	await updating;
	expect(b.reload).toHaveBeenCalledTimes(1);
});

test('08 §25.7: an update deferred by an open form says so without a button, and a failed one returns to Perbarui', async () => {
	const b = browser(true);
	updateModule.noteUpgradeRequired(upgradeRequired);
	const close = updateModule.openForm();
	await updateModule.updateNow();
	expect(updateModule.updateRow(true)).toEqual({ text: 'Diperbarui setelah isianmu ditutup.' });
	vi.spyOn(console, 'error').mockImplementation(() => {});
	b.registration.update.mockRejectedValueOnce(new Error('network'));
	close();
	await vi.waitFor(() => {
		expect(updateModule.updateRow(true)?.button).toEqual({ label: 'Perbarui', busy: false });
	});
});
