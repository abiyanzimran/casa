-- name: GetIdempotencyRecord :one
-- 09.2 §6 step 1, §8: served by the primary key.
SELECT command, request_hash, response FROM idempotency_record
 WHERE scope_id = @scope_id AND idempotency_key = @idempotency_key;

-- name: InsertIdempotencyRecord :exec
-- 09.2 §6 step 8; created_at is the database default (09.2 §5.1).
INSERT INTO idempotency_record (scope_id, idempotency_key, command, request_hash, response)
VALUES (@scope_id, @idempotency_key, @command, @request_hash, @response);
