-- The job queue of the jobs module (09.3 §9; 09.2 §3 "Tabel infrastruktur
-- dengan pemilik campuran", §5.1, both with their notes of 14 November 2026,
-- architect ruling Q1 (b); plan 1.7 decisions 2, 3). A domain transaction
-- enqueues a job of its own household (household_id NULL: a system job,
-- written only outside a household); job_runner claims due jobs across
-- households and records retries only through locked_until, attempts and
-- last_error; the job itself runs as app_rw, which deletes it once it ran.
-- payload holds identifiers only and last_error a code only (SEC-44). No
-- foreign key on household_id: a system job has none (09.2 §5.1). Expand
-- only: the old binary never reads this table (09.2 §10).

-- +goose Up
CREATE TABLE job (
	id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	household_id uuid NULL,
	-- Widened by each kind, as object_event.object_kind.
	kind         text NOT NULL CHECK (kind IN ('household.expire_invitation')),
	run_at       timestamptz NOT NULL,
	unique_key   text NULL,
	payload      bytea NOT NULL,
	attempts     int NOT NULL DEFAULT 0 CHECK (attempts >= 0),
	locked_until timestamptz NULL,
	last_error   text NULL
);

-- The same job is not enqueued twice. household_id leads (09.2 §3 layer 2),
-- so another household's job never collides with, or silently skips, this
-- household's ON CONFLICT DO NOTHING; NULLS NOT DISTINCT keeps a system job
-- unique; WHERE frees a job without a unique key.
CREATE UNIQUE INDEX job_unique ON job (household_id, kind, unique_key) NULLS NOT DISTINCT
	WHERE unique_key IS NOT NULL;
-- ClaimDue: a new job by run_at, and a job waiting for its retry or whose
-- lease ran out by locked_until.
CREATE INDEX job_due ON job (run_at) WHERE locked_until IS NULL;
CREATE INDEX job_locked ON job (locked_until) WHERE locked_until IS NOT NULL;

-- Both policies name their role: permissive policies are ORed, so the claim
-- policy's USING (true) without TO would show app_rw every household's jobs.
ALTER TABLE job ENABLE ROW LEVEL SECURITY;
ALTER TABLE job FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON job FOR ALL TO app_rw
	USING (household_id = nullif(current_setting('app.household_id', true), '')::uuid
	    OR (household_id IS NULL AND nullif(current_setting('app.household_id', true), '')::uuid IS NULL))
	WITH CHECK (household_id = nullif(current_setting('app.household_id', true), '')::uuid
	         OR (household_id IS NULL AND nullif(current_setting('app.household_id', true), '')::uuid IS NULL));
CREATE POLICY claim ON job FOR ALL TO job_runner USING (true);
GRANT SELECT, INSERT, DELETE ON job TO app_rw;
GRANT SELECT, UPDATE (locked_until, attempts, last_error) ON job TO job_runner;

-- +goose Down
DROP TABLE job;
