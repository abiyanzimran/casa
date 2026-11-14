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
function browser(waiting: boolean) {
	const worker = { postMessage: vi.fn() };
	const registration = {
		update: vi.fn(() => Promise.resolve()),
		waiting: waiting ? worker : null
	};
	let changed = () => {};
	const reload = vi.fn();
	vi.stubGlobal('navigator', {
		serviceWorker: {
			getRegistration: () => Promise.resolve(registration),
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
	await Promise.resolve();
	expect(b.worker.postMessage).not.toHaveBeenCalled();
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
