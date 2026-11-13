package database_test

import (
	"crypto/rand"
	"errors"
	"slices"
	"testing"

	"github.com/jackc/pgx/v5"
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
		scope := database.Scope{HouseholdID: h1, AccountID: account}

		requireNoScope := func(after string) {
			t.Helper()
			for _, setting := range []string{"app.household_id", "app.account_id"} {
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
