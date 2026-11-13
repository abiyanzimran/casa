-- The operator notice per account (SEC-34; 09.2 §5.1, §3 note of 14 November
-- 2026; 09.3 §1 "Mekanisme"; plan 1.6 decision 6). The text of each version
-- is embedded in the binary (package household), so there is no table of
-- versions: a row is one issue of a version's identity to one account,
-- written when the identity in force differs from the one last issued to the
-- account (a rename A→B→A issues A again; 10.1 note before §11.3), and
-- acknowledged once. The history of every issue stays (FR-FAM-019/AC-3). A
-- global table owned by household, with no foreign key to account (it would
-- cross a module boundary); account deletion removes the rows (SEC-105,
-- step 6). The two new security trail kinds record each issue and each
-- acknowledgement (10.1 §11 note of 14 November 2026 (c)). Expand only: the
-- old binary runs on it unchanged (09.2 §10).

-- +goose Up
CREATE TABLE operator_notice_ack (
	account_id      uuid NOT NULL,
	-- The account's issues in order, from 1; the next issue is the highest
	-- plus 1, so two first sights at once write one row (09.2 §5.1).
	issue_seq       int NOT NULL CHECK (issue_seq > 0),
	variant         text NOT NULL CHECK (variant IN ('couple', 'family', 'general')),
	version         int NOT NULL CHECK (version > 0),
	-- SHA-256 of the length-prefixed variant, number, title and paragraphs
	-- with their feature tags, after {pengelola} is filled: the version's
	-- identity for the account (10.1 §11 note of 14 November 2026 (a)).
	text_sha256     bytea NOT NULL CHECK (length(text_sha256) = 32),
	issued_at       timestamptz NOT NULL,
	acknowledged_at timestamptz NULL,
	PRIMARY KEY (account_id, issue_seq)
);

-- Read and written only in the account's own scope.
ALTER TABLE operator_notice_ack ENABLE ROW LEVEL SECURITY;
ALTER TABLE operator_notice_ack FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON operator_notice_ack
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid)
	WITH CHECK (account_id = nullif(current_setting('app.account_id', true), '')::uuid);
GRANT SELECT, INSERT, UPDATE (acknowledged_at) ON operator_notice_ack TO app_rw;

ALTER TABLE security_event
	DROP CONSTRAINT security_event_kind_check,
	ADD CONSTRAINT security_event_kind_check
		CHECK (kind IN ('sign_in', 'session_revoked', 'reauthentication', 'password_changed', 'recovery',
		                'notice_version', 'notice_acknowledged'));

-- +goose Down
-- Fails by design once a security event has a notice kind: the narrower
-- check would reject that row. A rollback is the old binary on this schema,
-- not this Down (09.2 §10).
ALTER TABLE security_event
	DROP CONSTRAINT security_event_kind_check,
	ADD CONSTRAINT security_event_kind_check
		CHECK (kind IN ('sign_in', 'session_revoked', 'reauthentication', 'password_changed', 'recovery'));

DROP TABLE operator_notice_ack;
