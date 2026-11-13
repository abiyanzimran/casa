-- name: InsertSetupToken :exec
-- SEC-158 (b): one token for one named person, bound to an email.
INSERT INTO setup_token (token_hash, email, created_at, expires_at)
VALUES (@token_hash, @email, @created_at, @expires_at);

-- name: GetSetupTokenForEmail :one
-- SEC-158 (b): the token is found by its hash and matches only the email it
-- is bound to, compared as citext (the same rule as account.email).
SELECT email, expires_at, used_at FROM setup_token
 WHERE token_hash = @token_hash AND email = @email;

-- name: GetAccountIDByEmail :one
-- 10.1 §5.1 rule 4: a lookup by email selects only the id.
SELECT id FROM account WHERE email = @email;

-- name: UpsertSignupVerification :exec
-- One link per setup token: a new request replaces the old link (09.2 §5.1).
INSERT INTO signup_verification (token_hash, source, setup_token_hash, email, created_at, expires_at)
VALUES (@token_hash, 'setup_token', @setup_token_hash, @email, @created_at, @expires_at)
ON CONFLICT (setup_token_hash) DO UPDATE SET token_hash = excluded.token_hash, expires_at = excluded.expires_at;

-- name: GetSignupVerification :one
-- The id of the account the link would open is drawn here (09.2 §4).
SELECT setup_token_hash, email, expires_at, used_at, gen_random_uuid()::uuid AS new_account_id
  FROM signup_verification
 WHERE token_hash = @token_hash;

-- name: InsertAccount :exec
INSERT INTO account (id, email, created_at, signup_via)
VALUES (@id, @email, @created_at, 'setup_token');

-- name: UseSetupToken :execrows
-- SEC-158: the token ends with the account it opened.
UPDATE setup_token SET used_at = @now::timestamptz, account_id = @account_id
 WHERE token_hash = @token_hash AND used_at IS NULL AND expires_at > @now::timestamptz;

-- name: UseSignupVerification :execrows
UPDATE signup_verification SET used_at = @now::timestamptz
 WHERE token_hash = @token_hash AND setup_token_hash = @setup_token_hash
   AND used_at IS NULL AND expires_at > @now::timestamptz;

-- name: InsertCredential :exec
INSERT INTO credential (account_id, salt, hash, time_cost, memory_kib, parallelism, updated_at)
VALUES (@account_id, @salt, @hash, @time_cost, @memory_kib, @parallelism, @updated_at);

-- name: InsertSession :one
INSERT INTO session (account_id, client_kind, ua_summary, token_hash, client_version, created_at,
                     last_used_at, idle_timeout, absolute_expiry, token_issued_at)
VALUES (@account_id, @client_kind, '', @token_hash, @client_version, @now,
        @now, @idle_timeout, @absolute_expiry, @now)
RETURNING id;

-- name: InsertRefreshToken :exec
INSERT INTO refresh_token (token_hash, session_id, account_id, created_at)
VALUES (@token_hash, @session_id, @account_id, @created_at);

-- name: InsertSecurityEvent :exec
-- The account's security trail (SEC-16, SEC-39 (b)); detail never holds a
-- token, a password or an email.
INSERT INTO security_event (account_id, session_id, kind, detail, occurred_at)
VALUES (@account_id, @session_id, @kind, @detail, @occurred_at);
