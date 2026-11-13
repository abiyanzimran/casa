package database_test

import (
	"context"
	"crypto/rand"
	"errors"
	"slices"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

// rlsProbe is a household table shaped as 09.2 §3 layers 2 and 4 require. No
// real household table exists before 1.4, so the test creates one.
const rlsProbe = `
	CREATE TABLE rls_probe (
		household_id uuid NOT NULL,
		id           uuid NOT NULL DEFAULT gen_random_uuid(),
		PRIMARY KEY (household_id, id)
	);
	ALTER TABLE rls_probe ENABLE ROW LEVEL SECURITY;
	ALTER TABLE rls_probe FORCE ROW LEVEL SECURITY;
	CREATE POLICY household ON rls_probe
		USING (household_id = current_setting('app.household_id')::uuid)
		WITH CHECK (household_id = current_setting('app.household_id')::uuid);
	GRANT SELECT, INSERT, UPDATE, DELETE ON rls_probe TO app_rw;`

// mixedOwnerProbes carry the mixed-owner policies of 09.2 §3 word for word:
// idempotency_probe as idempotency_record, job_probe as job.
const mixedOwnerProbes = `
	CREATE TABLE idempotency_probe (scope_id uuid NOT NULL);
	ALTER TABLE idempotency_probe ENABLE ROW LEVEL SECURITY;
	ALTER TABLE idempotency_probe FORCE ROW LEVEL SECURITY;
	CREATE POLICY scope ON idempotency_probe
		USING (scope_id = nullif(current_setting('app.household_id', true), '')::uuid
		    OR scope_id = nullif(current_setting('app.account_id', true), '')::uuid);
	CREATE TABLE job_probe (household_id uuid);
	ALTER TABLE job_probe ENABLE ROW LEVEL SECURITY;
	ALTER TABLE job_probe FORCE ROW LEVEL SECURITY;
	CREATE POLICY household ON job_probe
		USING (household_id = nullif(current_setting('app.household_id', true), '')::uuid
		    OR (household_id IS NULL AND nullif(current_setting('app.household_id', true), '') IS NULL));
	GRANT SELECT, INSERT ON idempotency_probe, job_probe TO app_rw;`

var errSentinel = errors.New("sentinel")

func TestScope(t *testing.T) {
	d := databasetest.Migrated(t)
	ctx := t.Context()
	if _, err := d.Pool(t, "migrator").Exec(ctx, rlsProbe); err != nil {
		t.Fatal(err)
	}
	app := d.Pool(t, "app_rw")
	h1, h2, account := newUUID(), newUUID(), newUUID()
	for _, h := range []pgtype.UUID{h1, h1, h2} {
		err := database.Write(ctx, app, database.Scope{HouseholdID: h, AccountID: account}, func(tx pgx.Tx) error {
			_, err := tx.Exec(ctx, "INSERT INTO rls_probe (household_id) VALUES ($1)", h)
			return err
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Run("SEC-08 a transaction reads only rows of its household, with no filter in the query", func(t *testing.T) {
		for _, c := range []struct {
			household pgtype.UUID
			want      int
		}{{h1, 2}, {h2, 1}} {
			var got []pgtype.UUID
			err := database.Read(ctx, app, database.Scope{HouseholdID: c.household}, func(tx pgx.Tx) error {
				rows, err := tx.Query(ctx, "SELECT household_id FROM rls_probe")
				if err != nil {
					return err
				}
				got, err = pgx.CollectRows(rows, pgx.RowTo[pgtype.UUID])
				return err
			})
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != c.want {
				t.Errorf("household %v: %d rows, want %d", c.household, len(got), c.want)
			}
			for _, h := range got {
				if h != c.household {
					t.Errorf("household %v read a row of household %v", c.household, h)
				}
			}
		}
	})

	t.Run("SEC-08 a row for another household cannot be written", func(t *testing.T) {
		for _, statement := range []string{
			"INSERT INTO rls_probe (household_id) VALUES ($1)",
			"UPDATE rls_probe SET household_id = $1",
		} {
			err := database.Write(ctx, app, database.Scope{HouseholdID: h1}, func(tx pgx.Tx) error {
				_, err := tx.Exec(ctx, statement, h2)
				return err
			})
			if code := databasetest.PgErrorCode(err); code != "42501" {
				t.Errorf("%s: err = %v, want a row-level security violation (42501)", statement, err)
			}
		}
		var deleted int64
		err := database.Write(ctx, app, database.Scope{HouseholdID: h1}, func(tx pgx.Tx) error {
			tag, err := tx.Exec(ctx, "DELETE FROM rls_probe WHERE household_id = $1", h2)
			deleted = tag.RowsAffected()
			return err
		})
		if err != nil || deleted != 0 {
			t.Errorf("delete of household 2 rows from household 1: deleted %d, err %v; want 0, nil", deleted, err)
		}
		if got := countIn(t, app, h2); got != 1 {
			t.Errorf("household 2 has %d rows, want 1", got)
		}
	})

	t.Run("SEC-08 without a household a household table cannot be read", func(t *testing.T) {
		err := database.Read(ctx, app, database.Scope{AccountID: account}, func(tx pgx.Tx) error {
			_, err := tx.Exec(ctx, "SELECT FROM rls_probe")
			return err
		})
		if code := databasetest.PgErrorCode(err); code != "22P02" {
			t.Errorf("err = %v, want invalid uuid '' (22P02)", err)
		}
	})

	t.Run("SEC-08 the table owner is filtered too", func(t *testing.T) {
		var got int
		err := pgx.BeginFunc(ctx, d.Pool(t, "migrator"), func(tx pgx.Tx) error {
			if _, err := tx.Exec(ctx, "SELECT set_config('app.household_id', $1::uuid::text, true)", h1); err != nil {
				return err
			}
			return tx.QueryRow(ctx, "SELECT count(*) FROM rls_probe").Scan(&got)
		})
		if err != nil {
			t.Fatal(err)
		}
		if got != 2 {
			t.Errorf("migrator sees %d rows in household 1, want 2 (FORCE ROW LEVEL SECURITY)", got)
		}
	})

	t.Run("SEC-08 the scope ends with its transaction", func(t *testing.T) {
		// One connection, so every call below reuses it. Read and Write set
		// both settings and would hide a leak, so the checks are plain pool
		// queries outside them.
		one := oneConnection(t, d)
		scope := database.Scope{HouseholdID: h1, AccountID: account, TokenHash: []byte{1}}

		requireNoScope := func(after string) {
			t.Helper()
			for _, setting := range []string{"app.household_id", "app.account_id", "app.email", "app.token_hash"} {
				var value string
				if err := one.QueryRow(ctx, "SELECT coalesce(current_setting($1, true), '')", setting).Scan(&value); err != nil {
					t.Fatal(err)
				}
				if value != "" {
					t.Errorf("after %s, %s = %q on the reused connection, want ''", after, setting, value)
				}
			}
			_, err := one.Exec(ctx, "SELECT FROM rls_probe")
			if code := databasetest.PgErrorCode(err); code != "22P02" {
				t.Errorf("after %s, reading rls_probe outside a scope: err = %v, want 22P02", after, err)
			}
		}

		if err := database.Write(ctx, one, scope, func(pgx.Tx) error { return nil }); err != nil {
			t.Fatal(err)
		}
		requireNoScope("a committed write")
		err := database.Write(ctx, one, scope, func(pgx.Tx) error { return errSentinel })
		if !errors.Is(err, errSentinel) {
			t.Fatalf("err = %v, want it to wrap the error of fn", err)
		}
		requireNoScope("a failed write")
		if err := database.Read(ctx, one, database.Scope{Email: "a@example.test", TokenHash: []byte{1}}, func(pgx.Tx) error { return nil }); err != nil {
			t.Fatal(err)
		}
		requireNoScope("a lookup")
	})

	t.Run("SEC-08 the nullif policies of 09.2 §3 hold on a reused connection", func(t *testing.T) {
		if _, err := d.Pool(t, "migrator").Exec(ctx, mixedOwnerProbes); err != nil {
			t.Fatal(err)
		}
		one := oneConnection(t, d)

		// The hazard 09.2 §3 documents: a connection that never set the
		// setting reads NULL, but after one SET LOCAL it reads '' for good,
		// so the bare cast fails instead of giving NULL.
		const bare = "SELECT current_setting('app.household_id', true)::uuid IS NULL"
		var isNull bool
		if err := one.QueryRow(ctx, bare).Scan(&isNull); err != nil || !isNull {
			t.Fatalf("fresh connection: %s = %t, err %v; want true", bare, isNull, err)
		}
		for _, w := range []struct {
			scope     database.Scope
			statement string
			arg       any
		}{
			{database.Scope{HouseholdID: h1, AccountID: account}, "INSERT INTO idempotency_probe VALUES ($1)", h1},
			{database.Scope{HouseholdID: h1, AccountID: account}, "INSERT INTO job_probe VALUES ($1)", h1},
			{database.Scope{AccountID: account}, "INSERT INTO idempotency_probe VALUES ($1)", account},
			{database.Scope{}, "INSERT INTO job_probe VALUES ($1)", pgtype.UUID{}},
		} {
			err := database.Write(ctx, one, w.scope, func(tx pgx.Tx) error {
				_, err := tx.Exec(ctx, w.statement, w.arg)
				return err
			})
			if err != nil {
				t.Fatalf("%s with scope %+v: %v", w.statement, w.scope, err)
			}
		}
		if _, err := one.Exec(ctx, bare); databasetest.PgErrorCode(err) != "22P02" {
			t.Errorf("after SET LOCAL: %s err = %v, want invalid uuid '' (22P02)", bare, err)
		}

		for _, c := range []struct {
			name  string
			scope database.Scope
			query string
			want  []pgtype.UUID
		}{
			// An account-level command after a household one reads its own
			// record and not the household's.
			{"account-level idempotency", database.Scope{AccountID: account},
				"SELECT scope_id FROM idempotency_probe", []pgtype.UUID{account}},
			// A system job is visible only outside a household.
			{"jobs outside a household", database.Scope{},
				"SELECT household_id FROM job_probe", []pgtype.UUID{{}}},
			{"jobs of household 1", database.Scope{HouseholdID: h1},
				"SELECT household_id FROM job_probe", []pgtype.UUID{h1}},
		} {
			var got []pgtype.UUID
			err := database.Read(ctx, one, c.scope, func(tx pgx.Tx) error {
				rows, err := tx.Query(ctx, c.query)
				if err != nil {
					return err
				}
				got, err = pgx.CollectRows(rows, pgx.RowTo[pgtype.UUID])
				return err
			})
			if err != nil || !slices.Equal(got, c.want) {
				t.Errorf("%s: rows %v, err %v; want %v", c.name, got, err, c.want)
			}
		}
	})

	t.Run("write is serializable and read is repeatable read and read only", func(t *testing.T) {
		var isolation string
		err := database.Write(ctx, app, database.Scope{HouseholdID: h1}, func(tx pgx.Tx) error {
			return tx.QueryRow(ctx, "SHOW transaction_isolation").Scan(&isolation)
		})
		if err != nil || isolation != "serializable" {
			t.Errorf("write: isolation %q, err %v; want serializable (09.2 §6)", isolation, err)
		}
		var readOnly string
		err = database.Read(ctx, app, database.Scope{HouseholdID: h1}, func(tx pgx.Tx) error {
			if err := tx.QueryRow(ctx, "SHOW transaction_isolation").Scan(&isolation); err != nil {
				return err
			}
			return tx.QueryRow(ctx, "SHOW transaction_read_only").Scan(&readOnly)
		})
		if err != nil || isolation != "repeatable read" || readOnly != "on" {
			t.Errorf("read: isolation %q, read only %q, err %v; want repeatable read, on (09.2 §6)", isolation, readOnly, err)
		}
		err = database.Read(ctx, app, database.Scope{HouseholdID: h1}, func(tx pgx.Tx) error {
			_, err := tx.Exec(ctx, "INSERT INTO rls_probe (household_id) VALUES ($1)", h1)
			return err
		})
		if code := databasetest.PgErrorCode(err); code != "25006" {
			t.Errorf("insert in a read: err = %v, want read_only_sql_transaction (25006)", err)
		}
	})

	t.Run("the email and the token hash are set when given", func(t *testing.T) {
		for _, c := range []struct {
			scope            database.Scope
			email, tokenHash string
		}{
			{database.Scope{Email: "a@example.test", TokenHash: []byte{1, 2}}, "a@example.test", "0102"},
			{database.Scope{}, "", ""},
		} {
			var email, tokenHash string
			err := database.Read(ctx, app, c.scope, func(tx pgx.Tx) error {
				return tx.QueryRow(ctx, "SELECT current_setting('app.email'), current_setting('app.token_hash')").Scan(&email, &tokenHash)
			})
			if err != nil || email != c.email || tokenHash != c.tokenHash {
				t.Errorf("scope %+v: app.email = %q, app.token_hash = %q, err %v; want %q, %q", c.scope, email, tokenHash, err, c.email, c.tokenHash)
			}
		}
	})

	t.Run("an email lookup stands alone", func(t *testing.T) {
		// 10.1 §5.1 rule 3: an authenticated handler cannot read another
		// account through its email, even by a bug.
		email := "a@example.test"
		for _, c := range []struct {
			name string
			run  func(fn func(pgx.Tx) error) error
		}{
			{"write with an email", func(fn func(pgx.Tx) error) error {
				return database.Write(ctx, app, database.Scope{Email: email}, fn)
			}},
			{"read with an email and an account", func(fn func(pgx.Tx) error) error {
				return database.Read(ctx, app, database.Scope{Email: email, AccountID: account}, fn)
			}},
			{"read with an email and a household", func(fn func(pgx.Tx) error) error {
				return database.Read(ctx, app, database.Scope{Email: email, HouseholdID: h1}, fn)
			}},
		} {
			calls := 0
			err := c.run(func(pgx.Tx) error { calls++; return nil })
			if !errors.Is(err, database.ErrEmailLookupNotAlone) || calls != 0 {
				t.Errorf("%s: err %v after %d calls, want ErrEmailLookupNotAlone after 0", c.name, err, calls)
			}
		}
		calls := 0
		err := database.Read(ctx, app, database.Scope{Email: email}, func(pgx.Tx) error { calls++; return nil })
		if err != nil || calls != 1 {
			t.Errorf("read with an email alone: err %v after %d calls, want nil after 1", err, calls)
		}
	})

	t.Run("the account is set when given", func(t *testing.T) {
		for _, c := range []struct {
			scope database.Scope
			want  string
		}{
			{database.Scope{AccountID: account}, account.String()},
			{database.Scope{}, ""},
		} {
			var got string
			err := database.Read(ctx, app, c.scope, func(tx pgx.Tx) error {
				return tx.QueryRow(ctx, "SELECT current_setting('app.account_id')").Scan(&got)
			})
			if err != nil || got != c.want {
				t.Errorf("scope %+v: app.account_id = %q, err %v; want %q", c.scope, got, err, c.want)
			}
		}
	})
}

// counterProbe is a household table two writers update at once.
const counterProbe = `
	CREATE TABLE counter_probe (household_id uuid PRIMARY KEY, n int NOT NULL);
	ALTER TABLE counter_probe ENABLE ROW LEVEL SECURITY;
	ALTER TABLE counter_probe FORCE ROW LEVEL SECURITY;
	CREATE POLICY household ON counter_probe
		USING (household_id = current_setting('app.household_id')::uuid)
		WITH CHECK (household_id = current_setting('app.household_id')::uuid);
	GRANT SELECT, INSERT, UPDATE ON counter_probe TO app_rw;`

// TestWriteRetries proves the retry of 09.2 §6 (lead ruling Q1 on the 1.4
// plan: 5 retries after the first attempt).
func TestWriteRetries(t *testing.T) {
	d := databasetest.Migrated(t)
	ctx := t.Context()
	app := d.Pool(t, "app_rw")
	serializationFailure := &pgconn.PgError{Code: "40001"}

	// failing returns an fn that fails with err on its first n calls, and
	// the counter of its calls.
	failing := func(n int32, err error) (func(pgx.Tx) error, *atomic.Int32) {
		calls := &atomic.Int32{}
		return func(pgx.Tx) error {
			if calls.Add(1) <= n {
				return err
			}
			return nil
		}, calls
	}

	t.Run("a serialization failure is retried up to 5 times", func(t *testing.T) {
		fn, calls := failing(5, serializationFailure)
		if err := database.Write(ctx, app, database.Scope{}, fn); err != nil || calls.Load() != 6 {
			t.Fatalf("err %v after %d calls, want nil after 6", err, calls.Load())
		}
	})
	t.Run("a deadlock is retried", func(t *testing.T) {
		fn, calls := failing(1, &pgconn.PgError{Code: "40P01"})
		if err := database.Write(ctx, app, database.Scope{}, fn); err != nil || calls.Load() != 2 {
			t.Fatalf("err %v after %d calls, want nil after 2", err, calls.Load())
		}
	})
	t.Run("after 5 retries the error is ErrRetryExhausted", func(t *testing.T) {
		fn, calls := failing(100, serializationFailure)
		err := database.Write(ctx, app, database.Scope{}, fn)
		if !errors.Is(err, database.ErrRetryExhausted) || databasetest.PgErrorCode(err) != "40001" || calls.Load() != 6 {
			t.Fatalf("err %v after %d calls, want ErrRetryExhausted wrapping 40001 after 6", err, calls.Load())
		}
	})
	t.Run("other errors are not retried", func(t *testing.T) {
		fn, calls := failing(100, errSentinel)
		if err := database.Write(ctx, app, database.Scope{}, fn); !errors.Is(err, errSentinel) || calls.Load() != 1 {
			t.Fatalf("err %v after %d calls, want the sentinel after 1", err, calls.Load())
		}
	})
	t.Run("a read is never retried", func(t *testing.T) {
		fn, calls := failing(100, serializationFailure)
		if err := database.Read(ctx, app, database.Scope{}, fn); databasetest.PgErrorCode(err) != "40001" || calls.Load() != 1 {
			t.Fatalf("err %v after %d calls, want 40001 after 1", err, calls.Load())
		}
	})
	t.Run("a cancelled context stops the retries", func(t *testing.T) {
		cancelled, cancel := context.WithCancel(ctx)
		calls := 0
		err := database.Write(cancelled, app, database.Scope{}, func(pgx.Tx) error {
			calls++
			cancel()
			return serializationFailure
		})
		if !errors.Is(err, context.Canceled) || calls != 1 {
			t.Fatalf("err %v after %d calls, want context.Canceled after 1", err, calls)
		}
	})
	t.Run("two conflicting writers both commit", func(t *testing.T) {
		if _, err := d.Pool(t, "migrator").Exec(ctx, counterProbe); err != nil {
			t.Fatal(err)
		}
		h := newUUID()
		scope := database.Scope{HouseholdID: h}
		err := database.Write(ctx, app, scope, func(tx pgx.Tx) error {
			_, err := tx.Exec(ctx, "INSERT INTO counter_probe VALUES ($1, 0)", h)
			return err
		})
		if err != nil {
			t.Fatal(err)
		}
		// Both first attempts read n before either writes, so one of them
		// cannot serialize and must run again.
		var arrived sync.WaitGroup
		arrived.Add(2)
		var calls atomic.Int32
		errs := make(chan error, 2)
		for range 2 {
			go func() {
				first := true
				errs <- database.Write(ctx, app, scope, func(tx pgx.Tx) error {
					calls.Add(1)
					var n int
					if err := tx.QueryRow(ctx, "SELECT n FROM counter_probe").Scan(&n); err != nil {
						return err
					}
					if first {
						first = false
						arrived.Done()
						arrived.Wait()
					}
					_, err := tx.Exec(ctx, "UPDATE counter_probe SET n = $1", n+1)
					return err
				})
			}()
		}
		for range 2 {
			if err := <-errs; err != nil {
				t.Fatalf("writer: %v", err)
			}
		}
		var n int
		err = database.Read(ctx, app, scope, func(tx pgx.Tx) error {
			return tx.QueryRow(ctx, "SELECT n FROM counter_probe").Scan(&n)
		})
		if err != nil || n != 2 || calls.Load() < 3 {
			t.Fatalf("n = %d after %d calls, err %v; want 2 after at least 3", n, calls.Load(), err)
		}
	})
}

// oneConnection is an app_rw pool of one connection, so every call on it
// reuses the same session and a setting left behind would show.
func oneConnection(t *testing.T, d databasetest.Database) *pgxpool.Pool {
	t.Helper()
	config, err := pgxpool.ParseConfig(d.URL("app_rw"))
	if err != nil {
		t.Fatal(err)
	}
	config.MaxConns = 1
	pool, err := pgxpool.NewWithConfig(t.Context(), config)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func countIn(t *testing.T, app *pgxpool.Pool, household pgtype.UUID) int {
	t.Helper()
	var n int
	err := database.Read(t.Context(), app, database.Scope{HouseholdID: household}, func(tx pgx.Tx) error {
		return tx.QueryRow(t.Context(), "SELECT count(*) FROM rls_probe").Scan(&n)
	})
	if err != nil {
		t.Fatal(err)
	}
	return n
}

func newUUID() pgtype.UUID {
	u := pgtype.UUID{Valid: true}
	_, _ = rand.Read(u.Bytes[:]) // never returns an error (crypto/rand, Go 1.24 and later)
	return u
}
