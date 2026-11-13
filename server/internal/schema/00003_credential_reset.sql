-- A credential reset (SEC-15, SEC-24; lead rulings and the security-privacy
-- decision on PR #32): the open recovery links of an account, so a reset can
-- use them all up, and the revoke reason of the sessions a password change
-- ends. Expand only: the old binary runs on it unchanged (09.2 §10).

-- +goose Up
CREATE INDEX recovery_token_account_open ON recovery_token (account_id) WHERE used_at IS NULL;

ALTER TABLE session
	DROP CONSTRAINT session_revoke_reason_check,
	ADD CONSTRAINT session_revoke_reason_check
		CHECK (revoke_reason IN ('user', 'reuse_detected', 'exit', 'purge', 'recovery', 'password_changed'));

-- +goose Down
-- Fails by design once a session has revoke_reason 'password_changed': the
-- narrower check would reject that row. A rollback is the old binary on this
-- schema, not this Down (09.2 §10).
ALTER TABLE session
	DROP CONSTRAINT session_revoke_reason_check,
	ADD CONSTRAINT session_revoke_reason_check
		CHECK (revoke_reason IN ('user', 'reuse_detected', 'exit', 'purge', 'recovery'));

DROP INDEX recovery_token_account_open;
