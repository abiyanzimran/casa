-- The open recovery links of an account, so a credential reset can use them
-- all up (SEC-15, SEC-24; lead ruling on PR B). Expand only: the old binary
-- runs on it unchanged (09.2 §10).

-- +goose Up
CREATE INDEX recovery_token_account_open ON recovery_token (account_id) WHERE used_at IS NULL;

-- +goose Down
DROP INDEX recovery_token_account_open;
