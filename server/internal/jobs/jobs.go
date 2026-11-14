// Package jobs is the background job runner of 09.3 §9 (09.2 §1 row jobs,
// §3, §5.1; ADR-015 item 4). A domain transaction enqueues a job in its own
// transaction (outbox); the runner claims due jobs one at a time as
// job_runner with FOR UPDATE SKIP LOCKED, runs the handler registered for the
// job's kind, then deletes the job as app_rw in the job's household, or
// records the failure for a retry with a growing backoff, and stops the job
// after MaxAttempts and flags the operator. Handlers are registered by
// cmd/casa as functions, so this package never imports the modules whose jobs
// it runs (plan 1.7, "Arah impor").
package jobs

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/jobs/jobsdb"
)

// MaxAttempts is the retry bound of 09.3 §9: the job stops after the 10th failed attempt.
const MaxAttempts = 10

const (
	// lease is how long a claimed job stays with its runner; a runner that
	// dies mid-job leaves it to be claimed again after it (plan 1.7
	// decision 2).
	lease = 5 * time.Minute
	// handlerTimeout is under the lease, so a job runs twice at once only
	// if its handler ignores its context past the timeout; even then the
	// late outcome cannot touch the new claim (DeleteJob, FailJob match the
	// attempt) (plan 1.7 decision 2) [Usulan].
	handlerTimeout = 4 * time.Minute
	// firstBackoff and maxBackoff bound the wait after a failed attempt:
	// min(10 s × 2^(n−1), 1 h) after the nth (plan 1.7 decision 2) [Usulan].
	firstBackoff = 10 * time.Second
	maxBackoff   = time.Hour
	// checkInterval is "memeriksa setiap detik" of 09.3 §9.
	checkInterval = time.Second
)

// The failure codes, the only text last_error and the log carry: never the
// handler's error text or the payload (SEC-44; plan 1.7 decision 2).
const (
	codeHandler     = "handler"
	codePanic       = "panic"
	codeTimeout     = "timeout"
	codeUnknownKind = "unknown_kind"
	// codeLease: the runner died during the last attempt, so its outcome is
	// unknown and the job stops without running again (plan 1.7 fix round 1).
	codeLease = "lease"
)

// Job is a claimed job: its payload holds identifiers only (09.2 §5.1; SEC-44).
type Job struct {
	ID, HouseholdID pgtype.UUID // HouseholdID invalid for a system job
	Kind            string
	Payload         []byte
	Attempt         int32 // 1 for the first run
}

// Handler runs one job; it must be idempotent (09.3 §9).
type Handler func(context.Context, Job) error

// NewJob is a job to enqueue in the domain transaction (outbox; 09.3 §9, §10).
type NewJob struct {
	HouseholdID pgtype.UUID
	Kind        string
	RunAt       time.Time
	UniqueKey   string
	Payload     []byte
}

// Enqueue writes j in tx with ON CONFLICT (household_id, kind, unique_key)
// WHERE unique_key IS NOT NULL DO NOTHING and a notification on channel jobs
// with an empty payload, delivered at commit (09.2 §5.1 note of 14 November
// 2026). An empty UniqueKey is none.
func Enqueue(ctx context.Context, tx pgx.Tx, j NewJob) error {
	q := jobsdb.New(tx)
	err := q.InsertJob(ctx, jobsdb.InsertJobParams{
		HouseholdID: j.HouseholdID,
		Kind:        j.Kind,
		RunAt:       j.RunAt,
		UniqueKey:   pgtype.Text{String: j.UniqueKey, Valid: j.UniqueKey != ""},
		Payload:     j.Payload,
	})
	if err != nil {
		return fmt.Errorf("insert job: %w", err)
	}
	if err := q.Notify(ctx); err != nil {
		return fmt.Errorf("notify jobs: %w", err)
	}
	return nil
}

// Config is what a Runner runs on. Claim and Now are the claim's pool and
// clock; the job's own writes use App (09.2 §3 row "job (pengambilan)").
type Config struct {
	Claim    *pgxpool.Pool // role job_runner
	App      *pgxpool.Pool // role app_rw
	Now      func() time.Time
	Logger   *slog.Logger
	Handlers map[string]Handler // by job kind
}

// Runner claims and runs due jobs.
type Runner struct {
	c Config
}

// NewRunner returns a Runner on c.
func NewRunner(c Config) *Runner {
	return &Runner{c: c}
}

// RunDue claims and runs due jobs one at a time until none is due; ran counts them,
// including a job stopped without running because its last attempt died.
// A job's failure is recorded on the job, not returned; err is a failure to
// claim or to record. ponytail: one sequential runner, so a slow job delays
// the others. Several claiming goroutines are safe (SKIP LOCKED) but would
// not scale as is: database.Write is SERIALIZABLE, so concurrent claims
// mostly end in 40001 retries; the way up is a READ COMMITTED claim
// transaction, where SKIP LOCKED alone keeps claims apart.
func (r *Runner) RunDue(ctx context.Context) (ran int, err error) {
	for ctx.Err() == nil {
		job, died, found, err := r.claim(ctx)
		if err != nil || !found {
			return ran, err
		}
		ran++
		if died {
			err = r.fail(context.WithoutCancel(ctx), job, codeLease)
		} else {
			err = r.run(ctx, job)
		}
		if err != nil {
			return ran, err
		}
	}
	return ran, ctx.Err()
}

// Run waits on LISTEN jobs for at most 1 s, then RunDue, until ctx ends; it
// returns after the running job finishes, and returns nil whenever ctx has
// ended, also before it began to listen. It returns the error of a
// database failure (claim, record, listen or wait) without retrying; the
// caller decides what follows (plan 1.7 Task 2, casa serve).
func (r *Runner) Run(ctx context.Context) error {
	pooled, err := r.c.Claim.Acquire(ctx)
	if err != nil {
		if ctx.Err() != nil {
			return nil
		}
		return fmt.Errorf("acquire the listening connection: %w", err)
	}
	// A listening connection never goes back to the pool: it would buffer
	// notifications for whoever used it next.
	conn := pooled.Hijack()
	defer func() { _ = conn.Close(context.WithoutCancel(ctx)) }()
	if err := jobsdb.New(conn).Listen(ctx); err != nil {
		if ctx.Err() != nil {
			return nil
		}
		return fmt.Errorf("listen on jobs: %w", err)
	}
	for {
		if _, err := r.RunDue(ctx); err != nil && ctx.Err() == nil {
			return err
		}
		wait, cancel := context.WithTimeout(ctx, checkInterval)
		_, err := conn.WaitForNotification(wait)
		cancel()
		if ctx.Err() != nil {
			return nil
		}
		if err != nil && !errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("wait for a job: %w", err)
		}
	}
}

// claim takes the oldest due job as job_runner and holds it for the lease,
// counting the attempt (plan 1.7 decision 2). died reports a job whose last
// attempt died with its runner.
func (r *Runner) claim(ctx context.Context) (job Job, died, found bool, err error) {
	now := r.c.Now()
	var row jobsdb.ClaimDueRow
	err = database.Write(ctx, r.c.Claim, database.Scope{}, func(tx pgx.Tx) (err error) {
		row, err = jobsdb.New(tx).ClaimDue(ctx, jobsdb.ClaimDueParams{
			Now:         now,
			LeaseUntil:  now.Add(lease),
			MaxAttempts: MaxAttempts,
		})
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Job{}, false, false, nil
	}
	if err != nil {
		return Job{}, false, false, fmt.Errorf("claim a due job: %w", err)
	}
	job = Job{ID: row.ID, HouseholdID: row.HouseholdID, Kind: row.Kind, Payload: row.Payload, Attempt: row.Attempts}
	return job, row.LastAttemptDied, true, nil
}

// run runs job and records its outcome. The job finishes and is recorded
// even when ctx ends, so Run returns only after it (plan 1.7 decision 2).
func (r *Runner) run(ctx context.Context, job Job) error {
	ctx = context.WithoutCancel(ctx)
	if code := r.handle(ctx, job); code != "" {
		return r.fail(ctx, job, code)
	}
	return r.delete(ctx, job)
}

// handle runs the job's handler and returns its failure code, or "" on
// success. The handler's error is reduced to a code on purpose: its text may
// carry household content, which never reaches the table or the log (SEC-44).
func (r *Runner) handle(ctx context.Context, job Job) (code string) {
	handler, ok := r.c.Handlers[job.Kind]
	if !ok {
		return codeUnknownKind
	}
	ctx, cancel := context.WithTimeout(ctx, handlerTimeout)
	defer cancel()
	defer func() {
		if recover() != nil {
			code = codePanic
		}
	}()
	switch err := handler(ctx, job); {
	case err == nil:
		return ""
	case errors.Is(ctx.Err(), context.DeadlineExceeded):
		return codeTimeout
	default:
		return codeHandler
	}
}

// delete removes a job that ran, as app_rw in its household (a system job
// outside any; 09.2 §3 note of 14 November 2026). A lease lost to another
// runner deletes nothing and is no error: that runner owns the job now.
func (r *Runner) delete(ctx context.Context, job Job) error {
	err := database.Write(ctx, r.c.App, database.Scope{HouseholdID: job.HouseholdID}, func(tx pgx.Tx) error {
		_, err := jobsdb.New(tx).DeleteJob(ctx, jobsdb.DeleteJobParams{ID: job.ID, Attempt: job.Attempt})
		return err
	})
	if err != nil {
		return fmt.Errorf("delete a job that ran: %w", err)
	}
	return nil
}

// fail records code and the retry's time as job_runner; after the last
// attempt the job waits forever and one ERROR line flags the operator with
// the kind, id and code only (09.3 §9; SEC-44). A lease lost to another
// runner records and logs nothing, as in delete.
func (r *Runner) fail(ctx context.Context, job Job, code string) error {
	stopped := job.Attempt >= MaxAttempts
	var recorded int64
	err := database.Write(ctx, r.c.Claim, database.Scope{}, func(tx pgx.Tx) error {
		var err error
		recorded, err = jobsdb.New(tx).FailJob(ctx, jobsdb.FailJobParams{
			ID:        job.ID,
			Attempt:   job.Attempt,
			LastError: pgtype.Text{String: code, Valid: true},
			Stopped:   stopped,
			RetryAt:   r.c.Now().Add(backoff(job.Attempt)),
		})
		return err
	})
	if err != nil {
		return fmt.Errorf("record a failed job: %w", err)
	}
	if stopped && recorded == 1 {
		r.c.Logger.LogAttrs(ctx, slog.LevelError, "job stopped",
			slog.String("kind", job.Kind), slog.String("id", job.ID.String()), slog.String("code", code))
	}
	return nil
}

// backoff is the wait after the attempt-th failed attempt, attempt ≥ 1:
// min(10 s × 2^(attempt−1), 1 h). The shift stops at 9, already past the
// hour, so it never overflows.
func backoff(attempt int32) time.Duration {
	return min(firstBackoff<<min(attempt-1, 9), maxBackoff)
}
