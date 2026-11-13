import 'fake-indexeddb/auto';
import { afterEach, expect, test, vi } from 'vitest';
import { deleteLocalDb, localDb, requestPersistence } from './local-db';

afterEach(deleteLocalDb);

test('SEC-165: importing the local database module creates no IndexedDB database', async () => {
	await import('./local-db');
	expect(await indexedDB.databases()).toEqual([]);
});

test('SEC-170: the local database holds only the outbox and meta stores', async () => {
	await localDb().open();
	expect(
		localDb()
			.tables.map((table) => table.name)
			.sort()
	).toEqual(['meta', 'outbox']);
});

test('SEC-171 (a): the persistent-storage answer is returned and kept', async () => {
	for (const [persist, expected] of [
		[() => Promise.resolve(true), true],
		[() => Promise.resolve(false), false],
		[() => Promise.reject(new Error('denied')), false],
		[undefined, false]
	] as const) {
		vi.stubGlobal('navigator', { storage: { persist } });
		expect(await requestPersistence(localDb())).toBe(expected);
		expect((await localDb().meta.get('persistence_granted'))?.value).toBe(expected);
	}
});

test('SEC-172 (a): deleting the local database removes it and the next use starts empty', async () => {
	await localDb().meta.put({ key: 'persistence_granted', value: true });
	await deleteLocalDb();
	expect(await indexedDB.databases()).toEqual([]);
	expect(await localDb().meta.count()).toBe(0);
});
