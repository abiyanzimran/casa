-- The per-reader feed of sync (ADR-017 item 3; 09.2 §5.1; 09.3 §3; plan 1.7
-- Task 4). sync.Changed, the one writer (09.2 §1 item 10), writes one row
-- per member who can read a changed object, in the command's transaction
-- (09.2 §6 step 6); pos rises without a gap per (household, member) through
-- reader_feed_counter. A row says only "object X changed": its content is
-- read through read at Sync (SEC-06). Every key leads with household_id (09.2
-- §3 layer 2) and names a member of that household. Expand only: the old
-- binary never reads these tables (09.2 §10).

-- +goose Up
CREATE TABLE reader_feed (
	household_id uuid NOT NULL,
	member_id    uuid NOT NULL,
	pos          bigint NOT NULL,
	-- Widened by each kind sync carries, as object_event.object_kind.
	object_kind  text NOT NULL CHECK (object_kind IN ('household', 'member', 'invitation')),
	object_id    uuid NOT NULL,
	change       text NOT NULL CHECK (change IN ('upsert', 'purge')),
	created_at   timestamptz NOT NULL DEFAULT now(), -- 09.2 §5.1: the pruning age (step 7)
	PRIMARY KEY (household_id, member_id, pos),
	FOREIGN KEY (household_id, member_id) REFERENCES membership
);

CREATE TABLE reader_feed_counter (
	household_id uuid NOT NULL,
	member_id    uuid NOT NULL,
	last_pos     bigint NOT NULL,
	PRIMARY KEY (household_id, member_id),
	FOREIGN KEY (household_id, member_id) REFERENCES membership
);

-- Layer 4 of 09.2 §3, without missing_ok. Both policies name app_rw from
-- the start (09.2 §3 note of 14 November 2026, Q1 (b), (d)): without TO a
-- policy would also be evaluated for job_runner's pruning (step 7) and
-- fail, since job_runner never sets app.household_id. The feed is
-- append-only for app_rw; the counter only moves last_pos. job_runner and
-- lifecycle_rw get nothing.
ALTER TABLE reader_feed ENABLE ROW LEVEL SECURITY;
ALTER TABLE reader_feed FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON reader_feed FOR ALL TO app_rw
	USING (true)
	WITH CHECK (household_id = current_setting('app.household_id')::uuid);
GRANT SELECT, INSERT ON reader_feed TO app_rw;

ALTER TABLE reader_feed_counter ENABLE ROW LEVEL SECURITY;
ALTER TABLE reader_feed_counter FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON reader_feed_counter FOR ALL TO app_rw
	USING (household_id = current_setting('app.household_id')::uuid)
	WITH CHECK (household_id = current_setting('app.household_id')::uuid);
GRANT SELECT, INSERT, UPDATE (last_pos) ON reader_feed_counter TO app_rw;

-- +goose Down
DROP TABLE reader_feed_counter;
DROP TABLE reader_feed;
