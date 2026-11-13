-- The adult age statement (SEC-80, note of 14 November 2026; SEC-109 (a);
-- 09.2 §5.1, §3 note of 14 November 2026; plan 1.6 decision 9). A global
-- table owned by household: one row per account and text version, written
-- in the transaction of CreateHousehold or AcceptInvitation. The statement,
-- its text version and the server's time only, never a birth date. No
-- foreign key to account (it would cross a module boundary); account
-- deletion removes the rows (SEC-105, step 6). Expand only: the old binary
-- runs on it unchanged (09.2 §10).

-- +goose Up
CREATE TABLE consent_record (
	account_id   uuid NOT NULL,
	kind         text NOT NULL CHECK (kind IN ('adult_age')),
	text_version int NOT NULL, -- the version of the text shown (08 §7.1)
	recorded_at  timestamptz NOT NULL,
	PRIMARY KEY (account_id, kind, text_version)
);

-- Read and written only in the account's own scope (SEC-80 note).
ALTER TABLE consent_record ENABLE ROW LEVEL SECURITY;
ALTER TABLE consent_record FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON consent_record
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid)
	WITH CHECK (account_id = nullif(current_setting('app.account_id', true), '')::uuid);
GRANT SELECT, INSERT ON consent_record TO app_rw;

-- +goose Down
DROP TABLE consent_record;
