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
-- The email is the one the token is bound to, so it never changes.
INSERT INTO signup_verification (token_hash, source, setup_token_hash, email, created_at, expires_at)
VALUES (@token_hash, 'setup_token', @setup_token_hash, @email, @created_at, @expires_at)
ON CONFLICT (setup_token_hash) DO UPDATE SET token_hash = excluded.token_hash, expires_at = excluded.expires_at;

-- name: UpsertInvitationSignupVerification :execrows
-- SEC-158 note (1) of 14 November 2026: one row per invitation token hash. A
-- restart before an account is born replaces the link and its email (the
-- email is free, SEC-31); once an account is born from the row, nothing
-- changes and zero rows come back. Served by the unique index on
-- invitation_token_hash.
INSERT INTO signup_verification (token_hash, source, invitation_token_hash, email, created_at, expires_at)
VALUES (@token_hash, 'invitation', @invitation_token_hash, @email, @created_at, @expires_at)
ON CONFLICT (invitation_token_hash) DO UPDATE
   SET token_hash = excluded.token_hash, email = excluded.email, expires_at = excluded.expires_at;

-- name: GetSignupVerification :one
-- The id of the account the link would open is drawn here (09.2 §4).
SELECT source, setup_token_hash, invitation_token_hash, email, expires_at, used_at,
       gen_random_uuid()::uuid AS new_account_id
  FROM signup_verification
 WHERE token_hash = @token_hash;

-- name: InsertAccount :exec
-- signup_via is the sign-up row's source (SEC-158: only setup_token may
-- create a household).
INSERT INTO account (id, email, created_at, signup_via)
VALUES (@id, @email, @created_at, @signup_via);

-- name: UseSetupToken :execrows
-- SEC-158: the token ends with the account it opened.
UPDATE setup_token SET used_at = @now::timestamptz, account_id = @account_id
 WHERE token_hash = @token_hash AND used_at IS NULL AND expires_at > @now::timestamptz;

-- name: UseSignupVerification :execrows
-- The link opens one account: compare-and-set on the link that was read and
-- its email, and account_id is set once (10.1 §5.1 note 1.6 item 4;
-- security-privacy ruling on Task 8). A link replaced since it was read
-- changes nothing. Served by the unique index on token_hash.
UPDATE signup_verification SET used_at = @now::timestamptz, account_id = @account_id
 WHERE token_hash = @token_hash AND email = @email AND account_id IS NULL
   AND used_at IS NULL AND expires_at > @now::timestamptz;

-- name: GetSignupInvitation :one
-- from_signup (FR-FAM-004/AC-14; plan 1.6 decision 15): the invitation token
-- hash of the sign-up the account was born from; none for a setup token.
-- Served by signup_verification_account.
SELECT invitation_token_hash::bytea FROM signup_verification
 WHERE account_id = @account_id AND source = 'invitation';

-- name: InsertCredential :exec
INSERT INTO credential (account_id, salt, hash, time_cost, memory_kib, parallelism, updated_at)
VALUES (@account_id, @salt, @hash, @time_cost, @memory_kib, @parallelism, @updated_at);

-- name: InsertSession :one
-- device_mode and ua_summary are set once, here (SEC-165; SEC-161 (e)).
INSERT INTO session (account_id, client_kind, device_mode, ua_summary, token_hash, client_version, created_at,
                     last_used_at, idle_timeout, absolute_expiry, token_issued_at)
VALUES (@account_id, @client_kind, @device_mode, @ua_summary, @token_hash, @client_version, @now,
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

-- name: GetSessionByCookieHash :one
-- SEC-161 (a), (c); 10.1 §5.1: a web request's cookie value is the current,
-- the previous, or the discarded value of at most one session. This lookup,
-- before the account is known and without a lock, finds only the account;
-- GetSessionByCookieHashForUpdate decides. Served by the unique index on
-- token_hash and the partial indexes session_prev_token and
-- session_discarded_token (09.2 §3, §5.1).
SELECT account_id
  FROM session
 WHERE token_hash = @token_hash OR prev_token_hash = @token_hash OR discarded_token_hash = @token_hash;

-- name: GetSessionByCookieHashForUpdate :one
-- SEC-161 (c) note 1.5 rule (4): one decision per row. The row is locked for
-- the transaction that classifies the value and writes the outcome (plan 1.5
-- decision 12); served by the same indexes as GetSessionByCookieHash.
SELECT id, account_id, client_kind, device_mode,
       token_hash = @token_hash AS current,
       coalesce(prev_token_hash = @token_hash, false)::boolean AS previous,
       coalesce(discarded_token_hash = @token_hash, false)::boolean AS discarded,
       token_issued_at, token_confirmed_at, prev_valid_until, last_used_at, idle_timeout, absolute_expiry, revoked_at
  FROM session
 WHERE token_hash = @token_hash OR prev_token_hash = @token_hash OR discarded_token_hash = @token_hash
   FOR UPDATE;

-- name: ConfirmCookie :exec
-- SEC-161 (c) note 1.5 rule (1): the browser brought the current value back,
-- so it holds it.
UPDATE session SET token_confirmed_at = @now::timestamptz
 WHERE id = @id AND token_confirmed_at IS NULL;

-- name: RotateCookie :one
-- SEC-161 (c): a new current value, not confirmed yet; the value the browser
-- holds serves as the previous one until prev_valid_until. That is the
-- current value, unless the current value is unconfirmed while the previous
-- one is still in its window: then the browser may hold only the previous
-- one (a forced rotation right after the boundary rotated or reissued, or
-- carried by the previous value), and it stays, with a new window (lead
-- ruling on the Task 5 review). A rotation by age never meets that case: the
-- window ends long before the current value is 15 minutes old. Fail closed
-- like TouchSession.
UPDATE session
   SET prev_token_hash = CASE WHEN token_confirmed_at IS NULL AND prev_valid_until > @now
                              THEN prev_token_hash ELSE token_hash END,
       prev_valid_until = @prev_valid_until::timestamptz,
       token_hash = @token_hash, token_issued_at = @now, token_confirmed_at = NULL
 WHERE id = @id AND revoked_at IS NULL
   AND last_used_at + idle_timeout > @now AND absolute_expiry > @now
RETURNING absolute_expiry;

-- name: ReissueCookie :one
-- SEC-161 (c) note 1.5 rules (2), (3): the current value never reached the
-- browser; it is discarded (and remembered), a new one is issued, and the
-- presented value is the previous one with a new window. Fail closed like
-- TouchSession.
UPDATE session
   SET discarded_token_hash = token_hash, token_hash = @token_hash,
       prev_token_hash = @presented, prev_valid_until = @prev_valid_until::timestamptz,
       token_issued_at = @now, token_confirmed_at = NULL
 WHERE id = @id AND revoked_at IS NULL
   AND last_used_at + idle_timeout > @now AND absolute_expiry > @now
RETURNING absolute_expiry;

-- name: TouchSession :execrows
-- The base of the idle limit. Fail closed: an ended session is never moved,
-- whatever the caller checked (lead ruling on #32).
UPDATE session SET last_used_at = @now
 WHERE id = @id AND revoked_at IS NULL
   AND last_used_at + idle_timeout > @now AND absolute_expiry > @now;

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

-- name: RotateAccessToken :execrows
-- Fail closed like TouchSession.
UPDATE session SET token_hash = @token_hash, token_issued_at = @now, last_used_at = @now
 WHERE id = @id AND revoked_at IS NULL
   AND last_used_at + idle_timeout > @now AND absolute_expiry > @now;

-- name: InsertReauthGrant :exec
-- cookie_hash: the hash of the cookie value issued with a web grant, NULL
-- for native (SEC-161 (c) note 1.5 rule (7)).
INSERT INTO reauth_grant (session_id, account_id, purpose, created_at, expires_at, cookie_hash)
VALUES (@session_id, @account_id, @purpose, @created_at, @expires_at, sqlc.narg(cookie_hash)::bytea);

-- name: HasOpenReauthGrant :one
-- The read-only check before hashing a new password; served by
-- reauth_grant_open. A web grant serves only the request that presents the
-- cookie value issued with it; a native grant (NULL) only a native request
-- (SEC-161 (c) note 1.5 rule (7)).
SELECT EXISTS (
	SELECT 1 FROM reauth_grant
	 WHERE session_id = @session_id AND purpose = @purpose AND used_at IS NULL AND expires_at > @now::timestamptz
	   AND cookie_hash IS NOT DISTINCT FROM sqlc.narg(cookie_hash)::bytea);

-- name: UseReauthGrant :execrows
-- Uses one open grant of the session for purpose (SEC-13), bound to the
-- presented cookie value as HasOpenReauthGrant is (SEC-161 (c) note 1.5 rule
-- (7)); served by reauth_grant_open.
UPDATE reauth_grant g SET used_at = @now::timestamptz
 WHERE g.used_at IS NULL AND g.id = (
	SELECT o.id FROM reauth_grant o
	 WHERE o.session_id = @session_id AND o.purpose = @purpose AND o.used_at IS NULL AND o.expires_at > @now
	   AND o.cookie_hash IS NOT DISTINCT FROM sqlc.narg(cookie_hash)::bytea
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

-- name: RevokeSessions :many
-- Every unrevoked session of the account except keep (none when NULL),
-- expired ones too, so no path can revive one (SEC-24; lead ruling on #32:
-- fail closed); live says which were still live, for the trail. Served by
-- session_account_live.
UPDATE session SET revoked_at = @now::timestamptz, revoke_reason = @reason::text
 WHERE account_id = @account_id AND revoked_at IS NULL
   AND id IS DISTINCT FROM sqlc.narg(keep)::uuid
RETURNING id, (last_used_at + idle_timeout > @now::timestamptz AND absolute_expiry > @now::timestamptz)::boolean AS live;

-- name: UseAccountRecoveryTokens :exec
-- A credential reset uses up every open recovery link of the account
-- (SEC-15, SEC-24); served by recovery_token_account_open.
UPDATE recovery_token SET used_at = @now::timestamptz
 WHERE account_id = @account_id AND used_at IS NULL;

-- name: GetSignupVia :one
-- SEC-158: how the account was made; only (b) setup_token may create a
-- household.
SELECT signup_via FROM account WHERE id = @id;

-- name: GetAccountEmail :one
-- 10.1 §5.1 note 1.6 item 3: only the email, by the primary key, for the
-- inviter's F04 screen (SEC-31).
SELECT email FROM account WHERE id = @id;
