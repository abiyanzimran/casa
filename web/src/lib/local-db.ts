// The local IndexedDB database (10.5 §5; SEC-170; ADR-024 item 7). Stores arrive with their
// features: the Sync mirror (1.9b, step 2), form drafts (step 2), the photo cache (step 5),
// the app lock verifier (step 6). Nothing here goes to localStorage or sessionStorage.
import Dexie, { type EntityTable } from 'dexie';

/**
 * One queued write command (09.3 §1). The idempotency key is made when the command is
 * saved and sent unchanged on every attempt (ADR-005; UUIDv4, ADR-017). The payload is the
 * caller's encoded request; the outbox never reads it.
 */
export interface OutboxEntry {
	idempotencyKey: string;
	kind: string;
	payload: Uint8Array;
	/** Protective commands skip the undo window and are sent first (09.3 §1; SEC-151). */
	protective: boolean;
	createdAt: number;
	/** createdAt plus the command kind's undo window (09.3 §1); not sent before this. */
	releaseAt: number;
	/** The server's rejection code; the entry stays queued and counted as unsent (08 §13.2). */
	rejectedCode?: string;
}

/** Device facts kept across sessions; the union grows with its users. */
export type MetaRow =
	{ key: 'persistence_granted'; value: boolean } | { key: 'wipe_pending'; value: true };

const name = 'casa';

/** Exported for tests that need a second tab's connection; the app uses localDb(). */
export class LocalDb extends Dexie {
	declare outbox: EntityTable<OutboxEntry, 'idempotencyKey'>;
	declare meta: EntityTable<MetaRow, 'key'>;

	constructor() {
		super(name);
		this.version(1).stores({ outbox: 'idempotencyKey', meta: 'key' });
	}
}

let instance: LocalDb | undefined;
let wiped = false;

/**
 * This tab's database. Creating it does not open IndexedDB; Dexie opens on the first
 * operation. Shared-device mode never calls this (SEC-165; ADR-024 item 7 (e)). After
 * deleteLocalDb() it throws until the page reloads, so nothing recreates the wiped database.
 */
export function localDb(): LocalDb {
	if (wiped) throw new Error('local data was wiped; the page reloads');
	return (instance ??= new LocalDb());
}

/** Whether the local database exists, without creating it (SEC-171 (e)). */
export function localDbExists(): Promise<boolean> {
	return Dexie.exists(name);
}

/**
 * Deletes the local database (SEC-164 (c), SEC-172 (a)). The old instance is closed without
 * auto-reopen, so a late write from a send still in flight fails instead of creating the
 * database again, and localDb() refuses for the rest of this page (the wipe then reloads).
 */
export async function deleteLocalDb(): Promise<void> {
	wiped = true;
	const closing = instance;
	instance = undefined;
	closing?.close({ disableAutoOpen: true });
	await Dexie.delete(name);
}

/**
 * Asks the browser to keep this site's storage and records the answer for LY-95
 * (SEC-171 (a)). A rejection or a browser without persist() counts as not granted, which is
 * the answer to show, so it is stored rather than thrown.
 */
export async function requestPersistence(db: LocalDb): Promise<boolean> {
	const storage: Partial<StorageManager> = navigator.storage;
	const granted = storage.persist ? await storage.persist().catch(() => false) : false;
	await db.meta.put({ key: 'persistence_granted', value: granted });
	return granted;
}

/**
 * Marks the local database as being wiped, before the wipe starts (SEC-172 (a)). The marker
 * lives in the database it guards, never in localStorage (10.5 §5): a successful deletion
 * removes it, a failed one leaves it for the next open. Without a database there is nothing to
 * mark and nothing is created (SEC-165 (b)). Its own connection, so a page that already wiped
 * once (localDb() refuses) can still mark a database another tab brought back.
 */
export async function markWipePending(): Promise<void> {
	if (!(await localDbExists())) return;
	const db = new LocalDb();
	try {
		await db.meta.put({ key: 'wipe_pending', value: true });
	} finally {
		db.close();
	}
}

/** Whether a wipe was started and did not finish (SEC-172 (a)); creates no database. */
export async function wipePending(): Promise<boolean> {
	if (!(await localDbExists())) return false;
	const db = new LocalDb();
	try {
		return (await db.meta.get('wipe_pending')) !== undefined;
	} finally {
		db.close();
	}
}
