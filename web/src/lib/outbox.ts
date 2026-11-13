// The command outbox: saved at once, held for the undo window, sent by one tab at a time
// (09.3 §1; ADR-030 item 10; ADR-005; SEC-171 (d), (f), SEC-173 (a)). It does not know the
// contract: the caller encodes the payload, gives the command kind's undo window and
// protective flag from the contract's command-kind table (ADR-013 item 4), and sends.
import type { LocalDb, OutboxEntry } from './local-db';

export type SendResult =
	{ kind: 'accepted' } | { kind: 'retry' } | { kind: 'rejected'; code: string };
export type Send = (entry: OutboxEntry) => Promise<SendResult>;

export interface NewCommand {
	kind: string;
	payload: Uint8Array;
	holdSeconds: number;
	protective: boolean;
}

/**
 * Due: not rejected, and its undo window has ended. A createdAt in the future means the
 * device clock went back after saving; elapsed time is then unknown, so the command is sent
 * rather than held for hours.
 */
function isDue(entry: OutboxEntry, now: number): boolean {
	return entry.rejectedCode === undefined && (entry.releaseAt <= now || entry.createdAt > now);
}

/** 09.3 §1 send order: protective first, then by release time, then by save time. */
function sendOrder(a: OutboxEntry, b: OutboxEntry): number {
	return (
		Number(b.protective) - Number(a.protective) ||
		a.releaseAt - b.releaseAt ||
		a.createdAt - b.createdAt
	);
}

// ES §4: the error is reported, never the command or its payload.
function report(error: unknown): void {
	console.error(error);
}

export function createOutbox(db: LocalDb, send: Send, clock: () => number = Date.now) {
	let started = false;
	let releaseTimer: ReturnType<typeof setTimeout> | undefined;

	function trigger(): void {
		flush().catch(report);
	}

	// ponytail: reads the whole outbox each pass; it holds unsent commands only, so it stays small.
	async function sendDue(): Promise<void> {
		for (;;) {
			const now = clock();
			const due = (await db.outbox.toArray()).filter((e) => isDue(e, now)).sort(sendOrder);
			if (due.length === 0) return;
			for (const entry of due) {
				const result = await send(entry);
				if (result.kind === 'retry') return;
				if (result.kind === 'accepted') await db.outbox.delete(entry.idempotencyKey);
				else await db.outbox.update(entry.idempotencyKey, { rejectedCode: result.code });
			}
		}
	}

	/** While started, wakes the outbox when the nearest held command's undo window ends. */
	async function armReleaseTimer(): Promise<void> {
		clearTimeout(releaseTimer);
		releaseTimer = undefined;
		if (!started) return;
		const now = clock();
		const held = (await db.outbox.toArray()).filter(
			(e) => e.rejectedCode === undefined && !isDue(e, now)
		);
		if (held.length === 0) return;
		const next = Math.min(...held.map((e) => e.releaseAt));
		releaseTimer = setTimeout(trigger, next - now);
	}

	/**
	 * Sends every due command, one tab at a time per browser profile (Web Locks; SEC-173 (a)).
	 * Accepted commands leave the outbox; rejected ones stay with their code and still count
	 * as unsent (08 §13.2); a retry answer or a thrown error ends the pass with the command
	 * untouched, and the error is passed on.
	 */
	async function flush(): Promise<void> {
		await navigator.locks.request('outbox', async () => {
			try {
				await sendDue();
			} finally {
				await armReleaseTimer();
			}
		});
	}

	return {
		/**
		 * Saves the command at once with a new idempotency key (ADR-005, ADR-017), then tries to
		 * send. A write error rejects with the original error and nothing is queued
		 * (SEC-171 (f)); a send error after saving is reported, never thrown at the caller.
		 */
		async enqueue(command: NewCommand): Promise<OutboxEntry> {
			const createdAt = clock();
			const entry: OutboxEntry = {
				idempotencyKey: crypto.randomUUID(),
				kind: command.kind,
				payload: command.payload,
				protective: command.protective,
				createdAt,
				releaseAt: createdAt + command.holdSeconds * 1000
			};
			await db.outbox.add(entry);
			trigger();
			return entry;
		},

		/**
		 * Removes a command still inside its undo window (09.3 §1). The clock is read inside
		 * the write transaction, so it cannot race a flush taking the same command.
		 */
		async undo(idempotencyKey: string): Promise<boolean> {
			return db.transaction('rw', db.outbox, async () => {
				const entry = await db.outbox.get(idempotencyKey);
				if (!entry || isDue(entry, clock())) return false;
				await db.outbox.delete(idempotencyKey);
				return true;
			});
		},

		/** Every queued command, rejected ones included (SEC-171 (c); 08 §13.2). */
		unsentCount(): Promise<number> {
			return db.outbox.count();
		},

		flush,

		/**
		 * Sends now, on `online`, and when the page becomes visible (ADR-030 item 10;
		 * SEC-171 (d)), and arms the release timer. Returns a function that stops all of it.
		 */
		start(
			win: EventTarget = window,
			doc: EventTarget & { visibilityState: DocumentVisibilityState } = document
		): () => void {
			const onVisibility = () => {
				if (doc.visibilityState === 'visible') trigger();
			};
			started = true;
			win.addEventListener('online', trigger);
			doc.addEventListener('visibilitychange', onVisibility);
			trigger();
			return () => {
				started = false;
				clearTimeout(releaseTimer);
				win.removeEventListener('online', trigger);
				doc.removeEventListener('visibilitychange', onVisibility);
			};
		}
	};
}
