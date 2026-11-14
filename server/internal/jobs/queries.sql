-- name: InsertJob :exec
-- Run by app_rw in the domain transaction (outbox; 09.3 §9, §10). The same
-- household, kind and unique key is written once (09.2 §5.1 note of 14
-- November 2026); attempts takes its default 0.
INSERT INTO job (household_id, kind, run_at, unique_key, payload)
VALUES (@household_id, @kind, @run_at, sqlc.narg(unique_key), @payload)
ON CONFLICT (household_id, kind, unique_key) WHERE unique_key IS NOT NULL DO NOTHING;

-- name: Notify :exec
-- PostgreSQL delivers it at commit, so a runner never wakes for a job that
-- does not exist (09.3 §9).
SELECT pg_notify('jobs', '');

-- name: Listen :exec
-- The runner's connection waits on it between checks (09.3 §9).
LISTEN jobs;

-- name: ClaimDue :one
-- Run by job_runner (09.2 §3 row "job (pengambilan)"; plan 1.7 decision 2):
-- the oldest due job no other runner holds, by job_due for a new job and
-- job_locked for a retry or a lease that ran out. The attempt counts when
-- claimed, so a runner that dies mid-job still uses one.
WITH due AS (
	SELECT d.id FROM job d
	 WHERE d.run_at <= @now AND d.attempts < @max_attempts::int
	   AND (d.locked_until IS NULL OR d.locked_until <= @now)
	 ORDER BY d.run_at
	 LIMIT 1
	   FOR UPDATE SKIP LOCKED
)
UPDATE job SET locked_until = @lease_until::timestamptz, attempts = job.attempts + 1
  FROM due
 WHERE job.id = due.id
RETURNING job.id, job.household_id, job.kind, job.payload, job.attempts;

-- name: DeleteJob :execrows
-- Run by app_rw in the job's household scope once its handler succeeded
-- (09.2 §3 note of 14 November 2026).
DELETE FROM job WHERE id = @id;

-- name: FailJob :exec
-- Run by job_runner: the retry waits through locked_until; a stopped job
-- waits forever and flags the operator (09.3 §9). last_error is a code only
-- (SEC-44).
UPDATE job
   SET last_error = @last_error,
       locked_until = CASE WHEN @stopped::boolean THEN 'infinity'::timestamptz ELSE @retry_at::timestamptz END
 WHERE id = @id;
