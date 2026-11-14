-- The feed writer's queries (ADR-017 item 3; 09.2 §5.1, §6 step 6; plan 1.7
-- decision 8). Each runs in the command's database.Write, whose household
-- scope row-level security checks (09.2 §3 layer 4).

-- name: BumpCounter :one
-- Moves the reader's last position on by n and returns it, so the reader's
-- n new rows take last_pos - n + 1 … last_pos, without a gap. Two commands
-- that reach the same reader collide on this row under SERIALIZABLE and one
-- is retried (ADR-017 "Konsekuensi"). Served by reader_feed_counter_pkey.
INSERT INTO reader_feed_counter AS c (household_id, member_id, last_pos)
VALUES (@household_id, @member_id, @n::bigint)
ON CONFLICT (household_id, member_id) DO UPDATE SET last_pos = c.last_pos + excluded.last_pos
RETURNING last_pos;

-- name: InsertFeedRows :exec
-- One upsert row per (reader, object) of a change, all in one statement:
-- the four arrays are one row each index (set-returning functions in a
-- select list run in step; a shorter array would leave a NULL, which NOT
-- NULL refuses). purge rows come with the first narrowing of access (F06,
-- step 6; plan 1.7 decision 8).
INSERT INTO reader_feed (household_id, member_id, pos, object_kind, object_id, change)
SELECT @household_id, unnest(@member_ids::uuid[]), unnest(@positions::bigint[]),
       unnest(@object_kinds::text[]), unnest(@object_ids::uuid[]), 'upsert';
