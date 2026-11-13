-- name: InsertHousehold :exec
-- The caller makes the id before the write (gen_random_uuid, ADR-017) and
-- sets it as the transaction's household: the policy's WITH CHECK is
-- id = app.household_id (09.2 §3 note of 14 November 2026). currency,
-- income_source_market and version take their defaults: IDR, ID, 0 (09.2
-- §5.1; FR-FAM-001/AC-7).
INSERT INTO household (id, name, time_zone, created_at)
VALUES (@id, @name, @time_zone, @created_at);
