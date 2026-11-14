-- name: BumpCounter :one
-- In the command's database.Write, whose household scope row-level security
-- checks (09.2 §3 layer 4; 09.2 §6 step 6).
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

-- name: LastPos :one
-- The reader's last feed position, in Sync's read transaction: P0 of a
-- snapshot, or L of the resync checks (plan 1.7 decisions 10, 11). No row
-- is 0. Served by reader_feed_counter_pkey.
SELECT coalesce((SELECT last_pos FROM reader_feed_counter
                  WHERE household_id = @household_id AND member_id = @member_id), 0)::bigint;

-- name: FeedAfter :many
-- The reader's feed rows after a position, by position, in Sync's read
-- transaction; the caller asks for its page plus one (ADR-017 item 3; plan
-- 1.7 decisions 11, 14). Served by reader_feed_pkey (household_id,
-- member_id, pos).
SELECT pos, object_kind, object_id, change FROM reader_feed
 WHERE household_id = @household_id AND member_id = @member_id AND pos > @after
 ORDER BY pos
 LIMIT @lim;
