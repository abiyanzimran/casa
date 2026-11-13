-- The identity tables and idempotency_record (09.2 §3 "Tabel global",
-- §5.1, §8; lead rulings Q2, Q3 and Q8 on the 1.4 plan).
--
-- Every table here is global (no household_id). Its one policy reads the
-- transaction's scope (internal/database): app.account_id for the account's
-- own rows, and, where a row must be found before the account is known,
-- app.email or app.token_hash as a lookup key (lead ruling Q2). Each setting
-- is read with missing_ok and nullif, so an absent one matches nothing
-- instead of failing (09.2 §3, note of 13 November 2026). Every time is
-- written by the server from its clock; none has a default except
-- idempotency_record.created_at, as 09.2 §5.1 has it.

-- +goose Up
-- citext compares emails without regard to case (09.2 §5.1; SEC-158 "dalam
-- bentuk yang dinormalkan"). A trusted extension: the database owner,
-- migrator, may create it.
CREATE EXTENSION citext;

CREATE TABLE account (
	id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	email      citext UNIQUE NOT NULL,
	time_zone  text NULL, -- 06 §2.4: empty means the household's zone
	locale     text NOT NULL DEFAULT 'id',
	created_at timestamptz NOT NULL,
	-- SEC-158: only an account made with a setup token (b) may create a
	-- household; an invited account (a) joins the inviter's.
	signup_via text NOT NULL CHECK (signup_via IN ('setup_token', 'invitation'))
);

-- Argon2id hash with its own parameters, so they can rise later without
-- breaking stored hashes (SEC-11; 09.4 §2).
CREATE TABLE credential (
	account_id  uuid PRIMARY KEY REFERENCES account (id),
	salt        bytea NOT NULL,
	hash        bytea NOT NULL,
	time_cost   int NOT NULL,
	memory_kib  int NOT NULL,
	parallelism smallint NOT NULL,
	updated_at  timestamptz NOT NULL
);

-- 09.2 §5.1, plus token_issued_at (the 15-minute access token, SEC-15) and
-- the revoke reason 'recovery' (SEC-24) (lead ruling Q3).
CREATE TABLE session (
	id               uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	account_id       uuid NOT NULL REFERENCES account (id),
	client_kind      text NOT NULL CHECK (client_kind IN ('ios', 'android', 'web')),
	device_mode      text NOT NULL DEFAULT 'personal' CHECK (device_mode IN ('personal', 'shared')),
	ua_summary       text NOT NULL,
	token_hash       bytea NOT NULL UNIQUE,
	prev_token_hash  bytea NULL,
	prev_valid_until timestamptz NULL,
	client_version   text NOT NULL,
	created_at       timestamptz NOT NULL,
	last_used_at     timestamptz NOT NULL,
	idle_timeout     interval NOT NULL,
	absolute_expiry  timestamptz NOT NULL,
	token_issued_at  timestamptz NOT NULL,
	revoked_at       timestamptz NULL,
	revoke_reason    text NULL CHECK (revoke_reason IN ('user', 'reuse_detected', 'exit', 'purge', 'recovery')),
	CHECK (client_kind = 'web' OR device_mode = 'personal')
);
-- Revoking every live session of an account (SEC-24).
CREATE INDEX session_account_live ON session (account_id) WHERE revoked_at IS NULL;

-- Native refresh tokens, rotated on every use; a used one presented again
-- revokes its session (SEC-15).
CREATE TABLE refresh_token (
	token_hash bytea PRIMARY KEY,
	session_id uuid NOT NULL REFERENCES session (id),
	account_id uuid NOT NULL REFERENCES account (id),
	created_at timestamptz NOT NULL,
	used_at    timestamptz NULL
);

-- A re-authentication, bound to its session and purpose, for one action
-- (SEC-13).
CREATE TABLE reauth_grant (
	id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	session_id uuid NOT NULL REFERENCES session (id),
	account_id uuid NOT NULL REFERENCES account (id),
	purpose    text NOT NULL CHECK (purpose IN ('change_password')),
	created_at timestamptz NOT NULL,
	expires_at timestamptz NOT NULL,
	used_at    timestamptz NULL
);
CREATE INDEX reauth_grant_open ON reauth_grant (session_id, purpose) WHERE used_at IS NULL;

-- Recovery through an emailed link (SEC-24).
CREATE TABLE recovery_token (
	token_hash bytea PRIMARY KEY,
	account_id uuid NOT NULL REFERENCES account (id),
	created_at timestamptz NOT NULL,
	expires_at timestamptz NOT NULL,
	used_at    timestamptz NULL
);

-- A one-time token the operator issues for one named person, bound to that
-- person's email (SEC-158 (b); 09.4 §2).
CREATE TABLE setup_token (
	token_hash bytea PRIMARY KEY,
	email      citext NOT NULL,
	created_at timestamptz NOT NULL,
	expires_at timestamptz NOT NULL,
	used_at    timestamptz NULL,
	account_id uuid NULL REFERENCES account (id)
);

-- The emailed link that finishes a sign-up, for both paths of SEC-158 (09.2
-- §5.1; lead rulings Q4, Q8 on the 1.4 plan). The account is made when the
-- link is opened, so no unverified account exists. 1.4 has only the source
-- 'setup_token'; 1.6 widens the source CHECK to 'invitation' and adds its
-- reference column. One link per setup token: a new request replaces
-- token_hash and expires_at.
CREATE TABLE signup_verification (
	id               uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	token_hash       bytea NOT NULL UNIQUE,
	source           text NOT NULL CHECK (source IN ('setup_token')),
	setup_token_hash bytea NULL UNIQUE REFERENCES setup_token (token_hash),
	email            citext NOT NULL,
	created_at       timestamptz NOT NULL,
	expires_at       timestamptz NOT NULL,
	used_at          timestamptz NULL,
	CHECK ((source = 'setup_token') = (setup_token_hash IS NOT NULL))
);

-- The account's security trail, append-only for app_rw (SEC-16, SEC-39 (b);
-- glossary "Jejak keamanan").
CREATE TABLE security_event (
	id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	account_id  uuid NOT NULL REFERENCES account (id),
	session_id  uuid NULL REFERENCES session (id),
	kind        text NOT NULL CHECK (kind IN ('sign_in', 'session_revoked', 'reauthentication', 'password_changed', 'recovery')),
	detail      text NULL, -- never a token, a password or an email
	occurred_at timestamptz NOT NULL
);

-- 09.2 §5.1, §8.
CREATE TABLE idempotency_record (
	scope_id        uuid NOT NULL,
	idempotency_key uuid NOT NULL,
	command         text NOT NULL,
	request_hash    bytea NOT NULL,
	response        bytea NOT NULL,
	created_at      timestamptz NOT NULL DEFAULT now(),
	PRIMARY KEY (scope_id, idempotency_key)
);

-- Every table forces row-level security (SEC-08; ADR-016 item 3).
ALTER TABLE account ENABLE ROW LEVEL SECURITY;
ALTER TABLE account FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON account
	USING (id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR email = nullif(current_setting('app.email', true), '')::citext)
	WITH CHECK (id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR email = nullif(current_setting('app.email', true), '')::citext);
GRANT SELECT, INSERT ON account TO app_rw;

ALTER TABLE credential ENABLE ROW LEVEL SECURITY;
ALTER TABLE credential FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON credential
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid)
	WITH CHECK (account_id = nullif(current_setting('app.account_id', true), '')::uuid);
GRANT SELECT, INSERT, UPDATE ON credential TO app_rw;

ALTER TABLE session ENABLE ROW LEVEL SECURITY;
ALTER TABLE session FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON session
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'))
	WITH CHECK (account_id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));
GRANT SELECT, INSERT, UPDATE ON session TO app_rw;

ALTER TABLE refresh_token ENABLE ROW LEVEL SECURITY;
ALTER TABLE refresh_token FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON refresh_token
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'))
	WITH CHECK (account_id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));
GRANT SELECT, INSERT, UPDATE ON refresh_token TO app_rw;

ALTER TABLE reauth_grant ENABLE ROW LEVEL SECURITY;
ALTER TABLE reauth_grant FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON reauth_grant
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid)
	WITH CHECK (account_id = nullif(current_setting('app.account_id', true), '')::uuid);
GRANT SELECT, INSERT, UPDATE ON reauth_grant TO app_rw;

ALTER TABLE recovery_token ENABLE ROW LEVEL SECURITY;
ALTER TABLE recovery_token FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON recovery_token
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'))
	WITH CHECK (account_id = nullif(current_setting('app.account_id', true), '')::uuid
	    OR token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));
GRANT SELECT, INSERT, UPDATE ON recovery_token TO app_rw;

-- Found only by its own hash: no account owns it until the sign-up ends.
ALTER TABLE setup_token ENABLE ROW LEVEL SECURITY;
ALTER TABLE setup_token FORCE ROW LEVEL SECURITY;
CREATE POLICY token ON setup_token
	USING (token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'))
	WITH CHECK (token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));
GRANT SELECT, INSERT, UPDATE ON setup_token TO app_rw;

-- Found by its own hash (the link) or by its setup token's hash (the holder
-- of the setup token replaces its link).
ALTER TABLE signup_verification ENABLE ROW LEVEL SECURITY;
ALTER TABLE signup_verification FORCE ROW LEVEL SECURITY;
CREATE POLICY token ON signup_verification
	USING (token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	    OR setup_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'))
	WITH CHECK (token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex')
	    OR setup_token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'));
GRANT SELECT, INSERT, UPDATE ON signup_verification TO app_rw;

ALTER TABLE security_event ENABLE ROW LEVEL SECURITY;
ALTER TABLE security_event FORCE ROW LEVEL SECURITY;
CREATE POLICY account ON security_event
	USING (account_id = nullif(current_setting('app.account_id', true), '')::uuid)
	WITH CHECK (account_id = nullif(current_setting('app.account_id', true), '')::uuid);
GRANT SELECT, INSERT ON security_event TO app_rw;

-- 09.2 §3 "Tabel infrastruktur dengan pemilik campuran", word for word.
ALTER TABLE idempotency_record ENABLE ROW LEVEL SECURITY;
ALTER TABLE idempotency_record FORCE ROW LEVEL SECURITY;
CREATE POLICY scope ON idempotency_record
	USING (scope_id = nullif(current_setting('app.household_id', true), '')::uuid
	    OR scope_id = nullif(current_setting('app.account_id', true), '')::uuid);
GRANT SELECT, INSERT ON idempotency_record TO app_rw;

-- +goose Down
DROP TABLE idempotency_record;
DROP TABLE security_event;
DROP TABLE signup_verification;
DROP TABLE setup_token;
DROP TABLE recovery_token;
DROP TABLE reauth_grant;
DROP TABLE refresh_token;
DROP TABLE session;
DROP TABLE credential;
DROP TABLE account;
DROP EXTENSION citext;
