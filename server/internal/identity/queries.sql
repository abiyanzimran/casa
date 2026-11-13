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

-- name: GetCredential :one
SELECT salt, hash, time_cost, memory_kib, parallelism FROM credential
 WHERE account_id = @account_id;

-- name: GetSessionByTokenHash :one
-- SEC-15: every request checks its access token here, by the unique index.
SELECT id, account_id, client_kind, token_issued_at, last_used_at, idle_timeout, absolute_expiry, revoked_at
  FROM session
 WHERE token_hash = @token_hash;

-- name: TouchSession :exec
-- The base of the idle limit.
UPDATE session SET last_used_at = @now WHERE id = @id;

-- name: GetRefreshToken :one
SELECT session_id, account_id FROM refresh_token WHERE token_hash = @token_hash;

-- name: GetSession :one
SELECT last_used_at, idle_timeout, absolute_expiry, revoked_at FROM session WHERE id = @id;

-- name: UseRefreshToken :execrows
-- SEC-15: a refresh token is used once; zero rows means it was used before.
UPDATE refresh_token SET used_at = @now::timestamptz
 WHERE token_hash = @token_hash AND used_at IS NULL;

-- name: RevokeSession :execrows
UPDATE session SET revoked_at = @now::timestamptz, revoke_reason = @reason::text
 WHERE id = @id AND revoked_at IS NULL;

-- name: RotateAccessToken :exec
UPDATE session SET token_hash = @token_hash, token_issued_at = @now, last_used_at = @now
 WHERE id = @id;

-- name: InsertReauthGrant :exec
INSERT INTO reauth_grant (session_id, account_id, purpose, created_at, expires_at)
VALUES (@session_id, @account_id, @purpose, @created_at, @expires_at);

-- name: HasOpenReauthGrant :one
-- The read-only check before hashing a new password; served by
-- reauth_grant_open.
SELECT EXISTS (
	SELECT 1 FROM reauth_grant
	 WHERE session_id = @session_id AND purpose = @purpose AND used_at IS NULL AND expires_at > @now::timestamptz);

-- name: UseReauthGrant :execrows
-- Uses one open grant of the session for purpose (SEC-13); served by
-- reauth_grant_open.
UPDATE reauth_grant g SET used_at = @now::timestamptz
 WHERE g.used_at IS NULL AND g.id = (
	SELECT o.id FROM reauth_grant o
	 WHERE o.session_id = @session_id AND o.purpose = @purpose AND o.used_at IS NULL AND o.expires_at > @now
	 LIMIT 1);

-- name: UpdateCredential :execrows
UPDATE credential
   SET salt = @salt, hash = @hash, time_cost = @time_cost, memory_kib = @memory_kib,
       parallelism = @parallelism, updated_at = @updated_at
 WHERE account_id = @account_id;

-- name: InsertRecoveryToken :exec
INSERT INTO recovery_token (token_hash, account_id, created_at, expires_at)
VALUES (@token_hash, @account_id, @created_at, @expires_at);

-- name: GetRecoveryToken :one
SELECT account_id, expires_at, used_at FROM recovery_token WHERE token_hash = @token_hash;

-- name: UseRecoveryToken :execrows
-- SEC-24: a recovery link is used once, within its lifetime.
UPDATE recovery_token SET used_at = @now::timestamptz
 WHERE token_hash = @token_hash AND used_at IS NULL AND expires_at > @now::timestamptz;

-- name: RevokeLiveSessions :many
-- Every session of the account still live at now, except keep (none when
-- NULL), so an ended session gets no second end (SEC-24); served by
-- session_account_live.
UPDATE session SET revoked_at = @now::timestamptz, revoke_reason = @reason::text
 WHERE account_id = @account_id AND revoked_at IS NULL
   AND last_used_at + idle_timeout > @now::timestamptz AND absolute_expiry > @now::timestamptz
   AND id IS DISTINCT FROM sqlc.narg(keep)::uuid
RETURNING id;

-- name: UseAccountRecoveryTokens :exec
-- A credential reset uses up every open recovery link of the account
-- (SEC-15, SEC-24); served by recovery_token_account_open.
UPDATE recovery_token SET used_at = @now::timestamptz
 WHERE account_id = @account_id AND used_at IS NULL;
