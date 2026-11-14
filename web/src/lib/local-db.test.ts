import 'fake-indexeddb/auto';
import Dexie from 'dexie';
import { afterEach, beforeEach, expect, test, vi } from 'vitest';

// Fresh module per test: after deleteLocalDb() the module refuses to open the database again
// until the page reloads, which a new module stands for here.
let db: typeof import('./local-db');

beforeEach(async () => {
	vi.resetModules();
	db = await import('./local-db');
});

afterEach(async () => {
	await Dexie.delete('casa');
});

test('SEC-165: importing the local database module creates no IndexedDB database', async () => {
	vi.resetModules();
	await import('./local-db');
	expect(await indexedDB.databases()).toEqual([]);
});

test('SEC-170: the local database holds only the outbox and meta stores', async () => {
	await db.localDb().open();
	expect(
		db
			.localDb()
			.tables.map((table) => table.name)
			.sort()
	).toEqual(['meta', 'outbox']);
	db.localDb().close();
});

test('SEC-171 (a): the persistent-storage answer is returned and kept', async () => {
	for (const [persist, expected] of [
		[() => Promise.resolve(true), true],
		[() => Promise.resolve(false), false],
		[() => Promise.reject(new Error('denied')), false],
		[undefined, false]
	] as const) {
		vi.stubGlobal('navigator', { storage: { persist } });
		expect(await db.requestPersistence(db.localDb())).toBe(expected);
		expect((await db.localDb().meta.get('persistence_granted'))?.value).toBe(expected);
	}
	db.localDb().close();
});

test('SEC-172 (a): deleting the local database removes it, and this page cannot open it again', async () => {
	await db.localDb().meta.put({ key: 'persistence_granted', value: true });
	await db.deleteLocalDb();
	expect(await indexedDB.databases()).toEqual([]);
	expect(() => db.localDb()).toThrow();
	expect(await indexedDB.databases()).toEqual([]);
});

test('SEC-165 (b): a database another tab deletes after the existence check is not made again', async () => {
	// Another tab deletes the database between the existence check and the open (plan
	// decision 5 would then read this browser as personal mode).
	vi.spyOn(Dexie, 'exists').mockResolvedValue(true);
	expect(await db.withExistingDb((existing) => existing.table('meta').count())).toBeUndefined();
	expect(await db.wipePending()).toBeUndefined();
	await db.markWipePending('session_revoked');
	vi.restoreAllMocks();
	expect(await indexedDB.databases()).toEqual([]);
});

test('SEC-172 (a): an existing database is read and written on its own connection', async () => {
	await db.localDb().meta.put({ key: 'persistence_granted', value: true });
	db.localDb().close();
	await db.markWipePending('sign_out');
	expect(await db.wipePending()).toBe('sign_out');
	expect(await db.withExistingDb((existing) => existing.table('outbox').count())).toBe(0);
});
