-- Sign-up through an invitation link (SEC-158 (a) and its note of 14
-- November 2026; 09.2 §3 note of 14 November 2026, §5.1; 10.1 §5.1 note 1.6
-- item 4; plan 1.6 decision 15; security-privacy ruling on Task 8).
-- signup_verification gets the invitation source: one row per invitation
-- token hash, which never changes; a restart before an account is born
-- replaces the row's link and email; account_id names the account born from
-- the row and is set once (WHERE account_id IS NULL in the query), for both
-- sources. No foreign key on account_id (09.2 §5.1); account deletion removes
-- the rows by account_id (SEC-105). Expand only: the old binary writes only
-- setup_token rows, through the same columns, and runs on it unchanged
-- (09.2 §10).

-- +goose Up
ALTER TABLE signup_verification
	ADD COLUMN invitation_token_hash bytea NULL UNIQUE,
	ADD COLUMN account_id uuid NULL,
	DROP CONSTRAINT signup_verification_source_check,
	ADD CONSTRAINT signup_verification_source_check CHECK (source IN ('setup_token', 'invitation')),
	DROP CONSTRAINT signup_verification_check,
	ADD CONSTRAINT signup_verification_check
		CHECK ((source = 'setup_token') = (setup_token_hash IS NOT NULL)
		   AND (source = 'invitation') = (invitation_token_hash IS NOT NULL));

-- SignupInvitation finds an account's sign-up by account_id = A; one account
-- is born from one row.
CREATE UNIQUE INDEX signup_verification_account ON signup_verification (account_id) WHERE account_id IS NOT NULL;

-- The link's own hash and the account only find the row; it is written only
-- through the hash of its source token (10.1 §5.1 rule 1).
ALTER POLICY token ON signup_verification
	USING (token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	    OR setup_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	    OR invitation_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	    OR account_id = nullif(current_setting('app.account_id', true), '')::uuid)
	WITH CHECK (setup_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	         OR invitation_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));
GRANT UPDATE (email, account_id, invitation_token_hash) ON signup_verification TO app_rw;

-- +goose Down
-- Fails by design once a row has the invitation source: the narrower checks
-- would reject it. A rollback is the old binary on this schema, not this
-- Down (09.2 §10).
REVOKE UPDATE (email, account_id) ON signup_verification FROM app_rw;
ALTER POLICY token ON signup_verification
	USING (token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	    OR setup_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'))
	WITH CHECK (setup_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));
DROP INDEX signup_verification_account;
ALTER TABLE signup_verification
	DROP CONSTRAINT signup_verification_check,
	ADD CONSTRAINT signup_verification_check CHECK ((source = 'setup_token') = (setup_token_hash IS NOT NULL)),
	DROP CONSTRAINT signup_verification_source_check,
	ADD CONSTRAINT signup_verification_source_check CHECK (source IN ('setup_token')),
	DROP COLUMN account_id,
	DROP COLUMN invitation_token_hash;
