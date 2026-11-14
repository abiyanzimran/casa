package jobs_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log/slog"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
	"github.com/abiyanzimran/casa/server/internal/jobs"
)

// kind is the one job kind the table accepts in 1.7 (plan 1.7 decision 2).
const kind = "household.expire_invitation"

func newUUID() pgtype.UUID {
	var u pgtype.UUID
	_, _ = rand.Read(u.Bytes[:])
	u.Bytes[6] = u.Bytes[6]&0x0f | 0x40
	u.Bytes[8] = u.Bytes[8]&0x3f | 0x80
	u.Valid = true
	return u
}

// clock is a test clock: Now reads it, advance moves it. Microseconds only,
// as timestamptz stores them.
type clock struct {
	mu  sync.Mutex
	now time.Time
}

func newClock() *clock {
	return &clock{now: time.Date(2026, 11, 14, 9, 0, 0, 0, time.UTC)}
}

func (c *clock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *clock) advance(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = c.now.Add(d)
}

type harness struct {
	admin, app, runner *pgxpool.Pool
	clock              *clock
}

func start(t *testing.T) *harness {
	t.Helper()
	d := databasetest.Migrated(t)
	return &harness{admin: d.Pool(t, d.AdminRole), app: d.Pool(t, "app_rw"), runner: d.Pool(t, "job_runner"), clock: newClock()}
}

func (h *harness) newRunner(handlers map[string]jobs.Handler, logger *slog.Logger) *jobs.Runner {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	return jobs.NewRunner(jobs.Config{Claim: h.runner, App: h.app, Now: h.clock.Now, Logger: logger, Handlers: handlers})
}

// enqueue writes j in its own committed household transaction, as a domain
// command would (outbox; 09.3 §9).
func (h *harness) enqueue(t *testing.T, j jobs.NewJob) {
	t.Helper()
	err := database.Write(t.Context(), h.app, database.Scope{HouseholdID: j.HouseholdID}, func(tx pgx.Tx) error {
		return jobs.Enqueue(t.Context(), tx, j)
	})
	if err != nil {
		t.Fatalf("enqueue: %v", err)
	}
}

// dueJob is a household job due now with a fresh household.
func (h *harness) dueJob() jobs.NewJob {
	return jobs.NewJob{HouseholdID: newUUID(), Kind: kind, RunAt: h.clock.Now(), Payload: []byte("id")}
}

// insertAsAdmin writes a job row as the superuser, bypassing row-level
// security, and returns its id.
func (h *harness) insertAsAdmin(t *testing.T, household pgtype.UUID, uniqueKey string) pgtype.UUID {
	t.Helper()
	var id pgtype.UUID
	err := h.admin.QueryRow(t.Context(), `
		INSERT INTO job (household_id, kind, run_at, unique_key, payload)
		VALUES ($1, $2, $3, nullif($4, ''), '\x00') RETURNING id`,
		household, kind, h.clock.Now(), uniqueKey).Scan(&id)
	if err != nil {
		t.Fatalf("insert job as admin: %v", err)
	}
	return id
}

type jobRow struct {
	Attempts    int32
	LockedUntil pgtype.Timestamptz
	LastError   pgtype.Text
}

func (h *harness) row(t *testing.T, id pgtype.UUID) (jobRow, bool) {
	t.Helper()
	var r jobRow
	err := h.admin.QueryRow(t.Context(), "SELECT attempts, locked_until, last_error FROM job WHERE id = $1", id).
		Scan(&r.Attempts, &r.LockedUntil, &r.LastError)
	if errors.Is(err, pgx.ErrNoRows) {
		return r, false
	}
	if err != nil {
		t.Fatal(err)
	}
	return r, true
}

func (h *harness) count(t *testing.T) int {
	t.Helper()
	var n int
	if err := h.admin.QueryRow(t.Context(), "SELECT count(*) FROM job").Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func (h *harness) onlyJobID(t *testing.T) pgtype.UUID {
	t.Helper()
	var id pgtype.UUID
	if err := h.admin.QueryRow(t.Context(), "SELECT id FROM job").Scan(&id); err != nil {
		t.Fatalf("the one job: %v", err)
	}
	return id
}

func (h *harness) runDue(t *testing.T, r *jobs.Runner) int {
	t.Helper()
	ran, err := r.RunDue(t.Context())
	if err != nil {
		t.Fatalf("RunDue: %v", err)
	}
	return ran
}

func ids(t *testing.T, pool *pgxpool.Pool, read func(func(pgx.Tx) error) error) []pgtype.UUID {
	t.Helper()
	var got []pgtype.UUID
	err := read(func(tx pgx.Tx) error {
		rows, err := tx.Query(t.Context(), "SELECT id FROM job ORDER BY id")
		if err != nil {
			return err
		}
		got, err = pgx.CollectRows(rows, pgx.RowTo[pgtype.UUID])
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func sortedIDs(u ...pgtype.UUID) []pgtype.UUID {
	slices.SortFunc(u, func(a, b pgtype.UUID) int { return bytes.Compare(a.Bytes[:], b.Bytes[:]) })
	return u
}

// TestJobTable proves the job table of 09.2 §3, §5.1 and plan 1.7 decision 3.
func TestJobTable(t *testing.T) {
	h := start(t)
	ctx := t.Context()
	x, y := newUUID(), newUUID()
	jobX := h.insertAsAdmin(t, x, "")
	jobY := h.insertAsAdmin(t, y, "")
	system := h.insertAsAdmin(t, pgtype.UUID{}, "")

	t.Run("a household transaction sees only its own jobs, and a system job only outside a household", func(t *testing.T) {
		got := ids(t, h.app, func(fn func(pgx.Tx) error) error {
			return database.Read(ctx, h.app, database.Scope{HouseholdID: x}, fn)
		})
		if want := []pgtype.UUID{jobX}; !slices.Equal(got, want) {
			t.Errorf("household X reads %v, want only %v", got, want)
		}
		got = ids(t, h.app, func(fn func(pgx.Tx) error) error {
			return database.Read(ctx, h.app, database.Scope{}, fn)
		})
		if want := []pgtype.UUID{system}; !slices.Equal(got, want) {
			t.Errorf("no household reads %v, want only the system job %v", got, want)
		}
	})

	t.Run("a household transaction cannot write a job of another household or a system job", func(t *testing.T) {
		for name, household := range map[string]pgtype.UUID{"another household": y, "a system job": {}} {
			err := database.Write(ctx, h.app, database.Scope{HouseholdID: x}, func(tx pgx.Tx) error {
				return jobs.Enqueue(ctx, tx, jobs.NewJob{HouseholdID: household, Kind: kind, RunAt: h.clock.Now(), Payload: []byte("id")})
			})
			if code := databasetest.PgErrorCode(err); code != "42501" {
				t.Errorf("%s: err = %v, want 42501", name, err)
			}
		}
	})

	t.Run("job_runner reads every job and updates only locked_until, attempts and last_error", func(t *testing.T) {
		got := ids(t, h.runner, func(fn func(pgx.Tx) error) error {
			return database.Write(ctx, h.runner, database.Scope{}, fn)
		})
		if want := sortedIDs(jobX, jobY, system); !slices.Equal(got, want) {
			t.Errorf("job_runner reads %v, want every job %v", got, want)
		}
		if _, err := h.runner.Exec(ctx, "UPDATE job SET locked_until = now(), attempts = 1, last_error = 'handler' WHERE id = $1", jobY); err != nil {
			t.Errorf("update the retry columns: %v", err)
		}
		for _, statement := range []string{
			"UPDATE job SET run_at = now() WHERE id = $1",
			"UPDATE job SET payload = '\\x01' WHERE id = $1",
			"DELETE FROM job WHERE id = $1",
		} {
			_, err := h.runner.Exec(ctx, statement, jobY)
			if code := databasetest.PgErrorCode(err); code != "42501" {
				t.Errorf("%s: err = %v, want 42501", statement, err)
			}
		}
	})

	t.Run("app_rw cannot update a job", func(t *testing.T) {
		err := database.Write(ctx, h.app, database.Scope{HouseholdID: x}, func(tx pgx.Tx) error {
			_, err := tx.Exec(ctx, "UPDATE job SET attempts = 1 WHERE id = $1", jobX)
			return err
		})
		if code := databasetest.PgErrorCode(err); code != "42501" {
			t.Errorf("err = %v, want 42501", err)
		}
	})

	t.Run("a job of an unknown kind is refused", func(t *testing.T) {
		err := database.Write(ctx, h.app, database.Scope{HouseholdID: x}, func(tx pgx.Tx) error {
			return jobs.Enqueue(ctx, tx, jobs.NewJob{HouseholdID: x, Kind: "household.unknown", RunAt: h.clock.Now(), Payload: []byte("id")})
		})
		if code := databasetest.PgErrorCode(err); code != "23514" {
			t.Errorf("err = %v, want check_violation (23514)", err)
		}
	})

	t.Run("the same kind and unique key in two households are two jobs, and a system job is written once", func(t *testing.T) {
		if _, err := h.admin.Exec(ctx, "DELETE FROM job"); err != nil {
			t.Fatal(err)
		}
		for _, household := range []pgtype.UUID{x, y, {}, {}} {
			h.enqueue(t, jobs.NewJob{HouseholdID: household, Kind: kind, RunAt: h.clock.Now(), UniqueKey: "k", Payload: []byte("id")})
		}
		var households, systemJobs int
		err := h.admin.QueryRow(ctx, "SELECT count(household_id), count(*) FILTER (WHERE household_id IS NULL) FROM job").
			Scan(&households, &systemJobs)
		if err != nil {
			t.Fatal(err)
		}
		if households != 2 || systemJobs != 1 {
			t.Errorf("household jobs %d, system jobs %d; want 2 and 1", households, systemJobs)
		}
	})
}

// TestRunner proves the jobs runner of 09.3 §9 and plan 1.7 decision 2.
func TestRunner(t *testing.T) {
	errHandler := errors.New("handler failed: secret detail")

	t.Run("a due job runs once with attempt 1 and is gone", func(t *testing.T) {
		h := start(t)
		j := h.dueJob()
		h.enqueue(t, j)
		var got []jobs.Job
		r := h.newRunner(map[string]jobs.Handler{kind: func(_ context.Context, job jobs.Job) error {
			got = append(got, job)
			return nil
		}}, nil)
		if ran := h.runDue(t, r); ran != 1 {
			t.Fatalf("ran %d, want 1", ran)
		}
		if len(got) != 1 || got[0].Attempt != 1 || got[0].HouseholdID != j.HouseholdID || got[0].Kind != kind ||
			!bytes.Equal(got[0].Payload, j.Payload) || !got[0].ID.Valid {
			t.Fatalf("handler saw %+v, want one run of the job with attempt 1", got)
		}
		if n := h.count(t); n != 0 {
			t.Errorf("%d jobs left, want 0", n)
		}
		if ran := h.runDue(t, r); ran != 0 {
			t.Errorf("second RunDue ran %d, want 0", ran)
		}
	})

	t.Run("a job before its run_at does not run", func(t *testing.T) {
		h := start(t)
		j := h.dueJob()
		j.RunAt = h.clock.Now().Add(time.Second)
		h.enqueue(t, j)
		r := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error { return nil }}, nil)
		if ran := h.runDue(t, r); ran != 0 {
			t.Fatalf("ran %d before run_at, want 0", ran)
		}
		h.clock.advance(time.Second)
		if ran := h.runDue(t, r); ran != 1 {
			t.Fatalf("ran %d at run_at, want 1", ran)
		}
	})

	t.Run("a failed job waits 10 s, then 20 s, and stops after 10 attempts with its code", func(t *testing.T) {
		h := start(t)
		h.enqueue(t, h.dueJob())
		id := h.onlyJobID(t)
		attempts := 0
		r := h.newRunner(map[string]jobs.Handler{kind: func(_ context.Context, job jobs.Job) error {
			attempts++
			if int(job.Attempt) != attempts {
				t.Errorf("run %d has Attempt %d", attempts, job.Attempt)
			}
			return errHandler
		}}, nil)
		if ran := h.runDue(t, r); ran != 1 {
			t.Fatalf("first run: ran %d", ran)
		}
		// 09.3 §9: min(10 s × 2^(n−1), 1 h) after the nth attempt [Usulan].
		for n, wait := range []time.Duration{10, 20, 40, 80, 160, 320, 640, 1280, 2560} {
			wait *= time.Second
			h.clock.advance(wait - time.Microsecond)
			if ran := h.runDue(t, r); ran != 0 {
				t.Fatalf("after attempt %d: ran %d before the %v backoff ended", n+1, ran, wait)
			}
			h.clock.advance(time.Microsecond)
			if ran := h.runDue(t, r); ran != 1 {
				t.Fatalf("after attempt %d: ran %d when the %v backoff ended, want 1", n+1, ran, wait)
			}
		}
		got, ok := h.row(t, id)
		if !ok || got.Attempts != jobs.MaxAttempts || got.LastError.String != "handler" ||
			got.LockedUntil.InfinityModifier != pgtype.Infinity {
			t.Fatalf("stopped job = %+v (exists %t), want attempts 10, last_error handler, locked_until infinity", got, ok)
		}
		h.clock.advance(1000 * time.Hour)
		if ran := h.runDue(t, r); ran != 0 || attempts != jobs.MaxAttempts {
			t.Errorf("stopped job ran again: RunDue %d, handler runs %d", ran, attempts)
		}
	})

	t.Run("a job of 10 attempts without a lease is not claimed", func(t *testing.T) {
		// Plan 1.7 ruling P8: attempts < 10 holds even with no locked_until.
		h := start(t)
		id := h.insertAsAdmin(t, newUUID(), "")
		if _, err := h.admin.Exec(t.Context(), "UPDATE job SET attempts = 10 WHERE id = $1", id); err != nil {
			t.Fatal(err)
		}
		r := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error { return nil }}, nil)
		if ran := h.runDue(t, r); ran != 0 {
			t.Fatalf("ran %d, want 0", ran)
		}
	})

	t.Run("a stopped job is logged with its kind and id and no payload", func(t *testing.T) {
		h := start(t)
		j := h.dueJob()
		j.Payload = []byte("payload-identifier")
		h.enqueue(t, j)
		id := h.onlyJobID(t)
		if _, err := h.admin.Exec(t.Context(), "UPDATE job SET attempts = 9 WHERE id = $1", id); err != nil {
			t.Fatal(err)
		}
		var logs bytes.Buffer
		r := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error { return errHandler }},
			slog.New(slog.NewJSONHandler(&logs, nil)))
		if ran := h.runDue(t, r); ran != 1 {
			t.Fatalf("ran %d, want 1", ran)
		}
		lines := strings.Split(strings.TrimSpace(logs.String()), "\n")
		if len(lines) != 1 {
			t.Fatalf("log lines %q, want one", lines)
		}
		var entry map[string]any
		if err := json.Unmarshal([]byte(lines[0]), &entry); err != nil {
			t.Fatal(err)
		}
		uuid, _ := id.Value()
		if entry["level"] != "ERROR" || entry["msg"] != "job stopped" || entry["kind"] != kind ||
			entry["id"] != uuid || entry["code"] != "handler" {
			t.Errorf("log entry %v, want ERROR job stopped with kind, id %v and code handler", entry, uuid)
		}
		// SEC-44: never the payload or the handler's error text.
		for _, secret := range []string{string(j.Payload), base64.StdEncoding.EncodeToString(j.Payload), errHandler.Error()} {
			if strings.Contains(logs.String(), secret) {
				t.Errorf("log %q contains %q", logs.String(), secret)
			}
		}
	})

	t.Run("a job whose handler panics fails with code panic", func(t *testing.T) {
		h := start(t)
		h.enqueue(t, h.dueJob())
		id := h.onlyJobID(t)
		r := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error { panic("boom") }}, nil)
		if ran := h.runDue(t, r); ran != 1 {
			t.Fatalf("ran %d, want 1", ran)
		}
		if got, ok := h.row(t, id); !ok || got.LastError.String != "panic" || got.Attempts != 1 {
			t.Errorf("job = %+v (exists %t), want attempts 1, last_error panic", got, ok)
		}
	})

	t.Run("a job of a kind without handler fails with code unknown_kind", func(t *testing.T) {
		h := start(t)
		h.enqueue(t, h.dueJob())
		id := h.onlyJobID(t)
		r := h.newRunner(map[string]jobs.Handler{}, nil)
		if ran := h.runDue(t, r); ran != 1 {
			t.Fatalf("ran %d, want 1", ran)
		}
		if got, ok := h.row(t, id); !ok || got.LastError.String != "unknown_kind" {
			t.Errorf("job = %+v (exists %t), want last_error unknown_kind", got, ok)
		}
	})

	t.Run("a job whose 10th attempt died is stopped and logged", func(t *testing.T) {
		// A runner that died during the last attempt left attempts 10 and a
		// finite lease; once it ends, the job stops without running again
		// and flags the operator (09.3 §9; plan 1.7 fix round 1).
		h := start(t)
		h.enqueue(t, h.dueJob())
		id := h.onlyJobID(t)
		_, err := h.admin.Exec(t.Context(), "UPDATE job SET locked_until = $2, attempts = 10 WHERE id = $1",
			id, h.clock.Now().Add(-time.Second))
		if err != nil {
			t.Fatal(err)
		}
		var logs bytes.Buffer
		handled := false
		r := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error {
			handled = true
			return nil
		}}, slog.New(slog.NewJSONHandler(&logs, nil)))
		if ran := h.runDue(t, r); ran != 1 || handled {
			t.Fatalf("RunDue %d, handler ran %t; want the job claimed once and its handler not run", ran, handled)
		}
		got, ok := h.row(t, id)
		if !ok || got.Attempts != jobs.MaxAttempts || got.LastError.String != "lease" ||
			got.LockedUntil.InfinityModifier != pgtype.Infinity {
			t.Fatalf("job = %+v (exists %t), want attempts 10, last_error lease, locked_until infinity", got, ok)
		}
		var entry map[string]any
		if err := json.Unmarshal(logs.Bytes(), &entry); err != nil || entry["msg"] != "job stopped" || entry["code"] != "lease" {
			t.Errorf("log %q, want one job stopped line with code lease (%v)", logs.String(), err)
		}
		h.clock.advance(1000 * time.Hour)
		if ran := h.runDue(t, r); ran != 0 {
			t.Errorf("stopped job claimed again: RunDue %d", ran)
		}
	})

	t.Run("a runner that lost its lease neither deletes nor fails the job", func(t *testing.T) {
		// The handler overran its lease and another runner holds the job
		// (attempt 2): the late outcome of attempt 1 must not touch it.
		h := start(t)
		h.enqueue(t, h.dueJob())
		id := h.onlyJobID(t)
		lease := h.clock.Now().Add(5 * time.Minute)
		takeOver := func() {
			_, err := h.admin.Exec(t.Context(), "UPDATE job SET attempts = attempts + 1, locked_until = $2 WHERE id = $1", id, lease)
			if err != nil {
				t.Error(err)
			}
		}
		for _, outcome := range []error{nil, errHandler} {
			r := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error {
				takeOver()
				return outcome
			}}, nil)
			if _, err := h.admin.Exec(t.Context(), "UPDATE job SET locked_until = NULL WHERE id = $1", id); err != nil {
				t.Fatal(err)
			}
			h.runDue(t, r)
			got, ok := h.row(t, id)
			if !ok || got.LastError.Valid || !got.LockedUntil.Time.Equal(lease) {
				t.Errorf("handler outcome %v: job = %+v (exists %t), want it untouched under the new holder's lease", outcome, got, ok)
			}
		}
	})

	t.Run("an enqueued job notifies channel jobs at commit only", func(t *testing.T) {
		// 09.3 §9: the runner is woken by LISTEN jobs after commit; a job
		// that rolled back wakes nobody. Run's own LISTEN is proven only
		// indirectly (Run also checks every second).
		h := start(t)
		listener, err := h.admin.Acquire(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		defer listener.Release()
		if _, err := listener.Exec(t.Context(), "LISTEN jobs"); err != nil {
			t.Fatal(err)
		}
		defer func() { _, _ = listener.Exec(context.Background(), "UNLISTEN jobs") }()
		noNotification := func(when string) {
			ctx, cancel := context.WithTimeout(t.Context(), 200*time.Millisecond)
			defer cancel()
			if n, err := listener.Conn().WaitForNotification(ctx); !errors.Is(err, context.DeadlineExceeded) {
				t.Errorf("%s: notification %+v, err %v; want none", when, n, err)
			}
		}
		j := h.dueJob()
		err = database.Write(t.Context(), h.app, database.Scope{HouseholdID: j.HouseholdID}, func(tx pgx.Tx) error {
			if err := jobs.Enqueue(t.Context(), tx, j); err != nil {
				return err
			}
			return errors.New("rollback")
		})
		if err == nil {
			t.Fatal("the transaction committed")
		}
		noNotification("after a rollback")
		h.enqueue(t, j)
		ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
		defer cancel()
		if n, err := listener.Conn().WaitForNotification(ctx); err != nil || n.Channel != "jobs" {
			t.Fatalf("after a commit: notification %+v, err %v; want one on jobs", n, err)
		}
		noNotification("after the one notification")
	})

	t.Run("a lease that ran out is claimed again", func(t *testing.T) {
		// A runner that died mid-job left its lease and its counted attempt
		// (plan 1.7 decision 2).
		h := start(t)
		h.enqueue(t, h.dueJob())
		id := h.onlyJobID(t)
		_, err := h.admin.Exec(t.Context(), "UPDATE job SET locked_until = $2, attempts = 1 WHERE id = $1",
			id, h.clock.Now().Add(-time.Second))
		if err != nil {
			t.Fatal(err)
		}
		var attempt int32
		r := h.newRunner(map[string]jobs.Handler{kind: func(_ context.Context, job jobs.Job) error {
			attempt = job.Attempt
			return nil
		}}, nil)
		if ran := h.runDue(t, r); ran != 1 || attempt != 2 {
			t.Fatalf("ran %d with Attempt %d, want 1 with Attempt 2", ran, attempt)
		}
	})

	t.Run("a job held by a running handler is not claimed by a second RunDue", func(t *testing.T) {
		// The lease of plan 1.7 decision 2: while the first runner's handler
		// runs (same clock, before the lease ends), the claim has committed
		// and holds no row lock, so only locked_until keeps the job from a
		// second runner. Deterministic, unlike the two-runner race below.
		h := start(t)
		h.enqueue(t, h.dueJob())
		started, release := make(chan struct{}), make(chan struct{})
		runs := 0
		first := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error {
			runs++
			close(started)
			<-release
			return nil
		}}, nil)
		done := make(chan error, 1)
		go func() {
			_, err := first.RunDue(t.Context())
			done <- err
		}()
		select {
		case <-started:
		case <-time.After(5 * time.Second):
			t.Fatal("the first runner did not start the job within 5 s")
		}
		second := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error {
			t.Error("the second runner ran the held job")
			return nil
		}}, nil)
		ran, err := second.RunDue(t.Context())
		close(release)
		if err != nil || ran != 0 {
			t.Errorf("second RunDue: ran %d, err %v; want 0 while the job is held", ran, err)
		}
		if err := <-done; err != nil {
			t.Fatalf("first RunDue: %v", err)
		}
		if runs != 1 || h.count(t) != 0 {
			t.Errorf("handler runs %d, jobs left %d; want 1 and 0", runs, h.count(t))
		}
	})

	t.Run("two runners at once claim each job once", func(t *testing.T) {
		h := start(t)
		const n = 20
		for range n {
			h.enqueue(t, h.dueJob())
		}
		var mu sync.Mutex
		seen := map[pgtype.UUID]int{}
		handlers := map[string]jobs.Handler{kind: func(_ context.Context, job jobs.Job) error {
			mu.Lock()
			defer mu.Unlock()
			seen[job.ID]++
			return nil
		}}
		runners := []*jobs.Runner{h.newRunner(handlers, nil), h.newRunner(handlers, nil)}
		barrier := make(chan struct{})
		errs := make([]error, len(runners))
		var wg sync.WaitGroup
		for i, r := range runners {
			wg.Go(func() {
				<-barrier
				_, errs[i] = r.RunDue(t.Context())
			})
		}
		close(barrier)
		wg.Wait()
		if err := errors.Join(errs...); err != nil {
			t.Fatalf("RunDue: %v", err)
		}
		if len(seen) != n {
			t.Errorf("%d jobs ran, want %d", len(seen), n)
		}
		for id, runs := range seen {
			if runs != 1 {
				t.Errorf("job %v ran %d times", id, runs)
			}
		}
		if left := h.count(t); left != 0 {
			t.Errorf("%d jobs left", left)
		}
	})

	t.Run("an enqueued job with the same household, kind and unique key is written once", func(t *testing.T) {
		h := start(t)
		j := h.dueJob()
		j.UniqueKey = "invitation-id"
		h.enqueue(t, j)
		h.enqueue(t, j)
		if n := h.count(t); n != 1 {
			t.Errorf("%d jobs, want 1", n)
		}
	})

	t.Run("a job enqueued in a rolled-back transaction does not exist", func(t *testing.T) {
		h := start(t)
		j := h.dueJob()
		errRollback := errors.New("rollback")
		err := database.Write(t.Context(), h.app, database.Scope{HouseholdID: j.HouseholdID}, func(tx pgx.Tx) error {
			if err := jobs.Enqueue(t.Context(), tx, j); err != nil {
				return err
			}
			return errRollback
		})
		if !errors.Is(err, errRollback) {
			t.Fatalf("Write: %v", err)
		}
		if n := h.count(t); n != 0 {
			t.Errorf("%d jobs, want 0", n)
		}
	})

	t.Run("Run wakes on the notification of a committed job", func(t *testing.T) {
		h := start(t)
		ran := make(chan struct{}, 1)
		r := h.newRunner(map[string]jobs.Handler{kind: func(context.Context, jobs.Job) error {
			ran <- struct{}{}
			return nil
		}}, nil)
		ctx, cancel := context.WithCancel(t.Context())
		done := make(chan error, 1)
		go func() { done <- r.Run(ctx) }()
		h.enqueue(t, h.dueJob())
		select {
		case <-ran:
		case <-time.After(5 * time.Second):
			t.Error("the job did not run within 5 s")
		}
		cancel()
		if err := <-done; err != nil {
			t.Errorf("Run: %v", err)
		}
	})

	t.Run("Run returns nil when its context ended before it started", func(t *testing.T) {
		// A clean shutdown during casa serve's startup is not a database
		// error (Run's contract; PR A review).
		h := start(t)
		ctx, cancel := context.WithCancel(t.Context())
		cancel()
		if err := h.newRunner(map[string]jobs.Handler{}, nil).Run(ctx); err != nil {
			t.Errorf("Run: %v, want nil", err)
		}
	})

	t.Run("Run returns when its context ends, after the running job finishes", func(t *testing.T) {
		h := start(t)
		started, release := make(chan struct{}), make(chan struct{})
		r := h.newRunner(map[string]jobs.Handler{kind: func(ctx context.Context, _ jobs.Job) error {
			close(started)
			<-release
			return ctx.Err() // the job's context outlives Run's
		}}, nil)
		h.enqueue(t, h.dueJob())
		ctx, cancel := context.WithCancel(t.Context())
		done := make(chan error, 1)
		go func() { done <- r.Run(ctx) }()
		select {
		case <-started:
		case <-time.After(5 * time.Second):
			t.Fatal("the job did not start within 5 s")
		}
		cancel()
		select {
		case err := <-done:
			t.Fatalf("Run returned %v while its job was running", err)
		default:
		}
		close(release)
		select {
		case err := <-done:
			if err != nil {
				t.Errorf("Run: %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("Run did not return within 5 s")
		}
		if n := h.count(t); n != 0 {
			t.Errorf("%d jobs left, want the finished job deleted", n)
		}
	})
}
