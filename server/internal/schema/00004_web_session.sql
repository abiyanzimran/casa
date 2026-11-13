-- The web session in a __Host- cookie (FS-1 1.5 plan decisions 12, 15;
-- SEC-161 (c); ADR-030 item 5; 09.2 §3, §5.1): when the browser proved it
-- holds the current value, the hash of a value discarded on a reissue, the
-- lookup of the previous and discarded values before the account is known,
-- and the cookie value a web re-authentication's grant is bound to. Expand only: the old binary runs on it unchanged (09.2 §10).

-- +goose Up
ALTER TABLE session
	ADD COLUMN token_confirmed_at timestamptz NULL,
	ADD COLUMN discarded_token_hash bytea NULL;

CREATE INDEX session_prev_token ON session (prev_token_hash) WHERE prev_token_hash IS NOT NULL;
CREATE INDEX session_discarded_token ON session (discarded_token_hash) WHERE discarded_token_hash IS NOT NULL;

-- Lookup clauses only in USING, never in WITH CHECK (10.1 §5.1 rules 1, 7).
ALTER POLICY account ON session
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	    OR prev_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	    OR discarded_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));

-- SEC-161 (c) note 1.5 rule (7): a web re-authentication's grant serves only
-- the cookie value issued with it, kept as its hash; NULL for a native
-- session's grant. Written once on INSERT (table-level grant), never updated.
ALTER TABLE reauth_grant ADD COLUMN cookie_hash bytea NULL;

-- No UPDATE on device_mode or ua_summary: the mode never changes within a
-- session (SEC-165).
GRANT UPDATE (prev_token_hash, prev_valid_until, token_confirmed_at, discarded_token_hash) ON session TO app_rw;

-- +goose Down
ALTER TABLE reauth_grant DROP COLUMN cookie_hash;

REVOKE UPDATE (prev_token_hash, prev_valid_until, token_confirmed_at, discarded_token_hash) ON session FROM app_rw;

ALTER POLICY account ON session
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));

DROP INDEX session_discarded_token;
DROP INDEX session_prev_token;

ALTER TABLE session
	DROP COLUMN discarded_token_hash,
	DROP COLUMN token_confirmed_at;
