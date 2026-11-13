-- name: InsertSetupToken :exec
-- SEC-158 (b): one token for one named person, bound to an email.
INSERT INTO setup_token (token_hash, email, created_at, expires_at)
VALUES (@token_hash, @email, @created_at, @expires_at);
