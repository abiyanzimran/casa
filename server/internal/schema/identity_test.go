package schema_test

import (
	"crypto/rand"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

// identityFixture is one account with a row in every account-owned identity
// table, plus a setup token and its sign-up link, each with its own hash.
type identityFixture struct {
	account, session                  pgtype.UUID
	email                             string
	sessionHash, refreshHash          []byte
	recoveryHash, setupHash, linkHash []byte
}

// accountTables are the identity tables an account owns, each with the
// column that names the owner.
var accountTables = map[string]string{
	"account":        "id",
	"credential":     "account_id",
	"session":        "account_id",
	"refresh_token":  "account_id",
	"reauth_grant":   "account_id",
	"recovery_token": "account_id",
	"security_event": "account_id",
}

// tokenTables are the tables found by app.token_hash, each with its hash
// columns.
var tokenTables = map[string][]string{
	"session":             {"token_hash"},
	"refresh_token":       {"token_hash"},
	"recovery_token":      {"token_hash"},
	"setup_token":         {"token_hash"},
	"signup_verification": {"token_hash", "setup_token_hash"},
}

func randomBytes() []byte {
	b := make([]byte, 32)
	_, _ = rand.Read(b) // never returns an error (crypto/rand, Go 1.24 and later)
	return b
}

func randomUUID() pgtype.UUID {
	u := pgtype.UUID{Valid: true}
	_, _ = rand.Read(u.Bytes[:]) // never returns an error (crypto/rand, Go 1.24 and later)
	return u
}

// newIdentityFixture writes f through database.Write as app_rw, so the
// policies' WITH CHECK clauses pass for an account's own rows.
func newIdentityFixture(t *testing.T, app *pgxpool.Pool, email string) identityFixture {
	t.Helper()
	ctx := t.Context()
	f := identityFixture{
		account: randomUUID(), session: randomUUID(), email: email,
		sessionHash: randomBytes(), refreshHash: randomBytes(), recoveryHash: randomBytes(),
		setupHash: randomBytes(), linkHash: randomBytes(),
	}
	now := time.Now()
	err := database.Write(ctx, app, database.Scope{AccountID: f.account}, func(tx pgx.Tx) error {
		for _, s := range []struct {
			sql  string
			args []any
		}{
			{"INSERT INTO account (id, email, created_at, signup_via) VALUES ($1, $2, $3, 'setup_token')",
				[]any{f.account, email, now}},
			{`INSERT INTO credential (account_id, salt, hash, time_cost, memory_kib, parallelism, updated_at)
			  VALUES ($1, '\x00', '\x00', 1, 64, 1, $2)`, []any{f.account, now}},
			{`INSERT INTO session (id, account_id, client_kind, ua_summary, token_hash, client_version,
			                       created_at, last_used_at, idle_timeout, absolute_expiry, token_issued_at)
			  VALUES ($1, $2, 'ios', '', $3, '1.0.0', $4, $4, '30 days', $4, $4)`,
				[]any{f.session, f.account, f.sessionHash, now}},
			{"INSERT INTO refresh_token (token_hash, session_id, account_id, created_at) VALUES ($1, $2, $3, $4)",
				[]any{f.refreshHash, f.session, f.account, now}},
			{`INSERT INTO reauth_grant (session_id, account_id, purpose, created_at, expires_at)
			  VALUES ($1, $2, 'change_password', $3, $3)`, []any{f.session, f.account, now}},
			{"INSERT INTO recovery_token (token_hash, account_id, created_at, expires_at) VALUES ($1, $2, $3, $3)",
				[]any{f.recoveryHash, f.account, now}},
			{"INSERT INTO security_event (account_id, session_id, kind, occurred_at) VALUES ($1, $2, 'sign_in', $3)",
				[]any{f.account, f.session, now}},
		} {
			if _, err := tx.Exec(ctx, s.sql, s.args...); err != nil {
				return fmt.Errorf("%s: %w", s.sql, err)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	err = database.Write(ctx, app, database.Scope{TokenHash: f.setupHash}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, "INSERT INTO setup_token (token_hash, email, created_at, expires_at) VALUES ($1, $2, $3, $3)",
			f.setupHash, email, now); err != nil {
			return err
		}
		_, err := tx.Exec(ctx, `INSERT INTO signup_verification (token_hash, source, setup_token_hash, email, created_at, expires_at)
		                        VALUES ($1, 'setup_token', $2, $3, $4, $4)`, f.linkHash, f.setupHash, email, now)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return f
}

// readColumn returns column of every row of table that scope can read.
func readColumn[T any](t *testing.T, app *pgxpool.Pool, scope database.Scope, table, column string) []T {
	t.Helper()
	var got []T
	err := database.Read(t.Context(), app, scope, func(tx pgx.Tx) error {
		rows, err := tx.Query(t.Context(), fmt.Sprintf("SELECT %s FROM %s", column, table))
		if err != nil {
			return err
		}
		got, err = pgx.CollectRows(rows, pgx.RowTo[T])
		return err
	})
	if err != nil {
		t.Fatalf("read %s.%s with scope %+v: %v", table, column, scope, err)
	}
	return got
}

// TestIdentityTablesAreAccountScoped proves the global identity tables of
// 09.2 §3 isolate accounts (ES §4; SEC-08) and that the lookup settings of
// lead ruling Q2 on the 1.4 plan find only the row they name.
func TestIdentityTablesAreAccountScoped(t *testing.T) {
	d := databasetest.Migrated(t)
	ctx := t.Context()
	app := d.Pool(t, "app_rw")
	a := newIdentityFixture(t, app, "Owner@Example.Test")
	b := newIdentityFixture(t, app, "partner@example.test")

	t.Run("an account scope reads only its own rows", func(t *testing.T) {
		for table, owner := range accountTables {
			got := readColumn[pgtype.UUID](t, app, database.Scope{AccountID: b.account}, table, owner)
			if len(got) == 0 || slices.ContainsFunc(got, func(u pgtype.UUID) bool { return u != b.account }) {
				t.Errorf("%s: account B reads owners %v, want only B %v", table, got, b.account)
			}
		}
		for table := range tokenTables {
			if _, owned := accountTables[table]; owned {
				continue
			}
			if got := readColumn[[]byte](t, app, database.Scope{AccountID: b.account}, table, "token_hash"); len(got) != 0 {
				t.Errorf("%s: an account scope reads %d rows, want 0", table, len(got))
			}
		}
	})

	t.Run("an account scope cannot write rows of another account", func(t *testing.T) {
		for _, statement := range []string{
			`INSERT INTO session (account_id, client_kind, ua_summary, token_hash, client_version, created_at,
			                      last_used_at, idle_timeout, absolute_expiry, token_issued_at)
			 VALUES ($1, 'ios', '', '\x01', '1.0.0', now(), now(), '1 day', now(), now())`,
			"INSERT INTO security_event (account_id, kind, occurred_at) VALUES ($1, 'sign_in', now())",
		} {
			err := database.Write(ctx, app, database.Scope{AccountID: b.account}, func(tx pgx.Tx) error {
				_, err := tx.Exec(ctx, statement, a.account)
				return err
			})
			if code := databasetest.PgErrorCode(err); code != "42501" {
				t.Errorf("%s: err = %v, want a row-level security violation (42501)", statement, err)
			}
		}
		var updated int64
		err := database.Write(ctx, app, database.Scope{AccountID: b.account}, func(tx pgx.Tx) error {
			tag, err := tx.Exec(ctx, "UPDATE credential SET hash = '\\x01' WHERE account_id = $1", a.account)
			updated = tag.RowsAffected()
			return err
		})
		if err != nil || updated != 0 {
			t.Errorf("update of A's credential from B: %d rows, err %v; want 0, nil", updated, err)
		}
	})

	t.Run("a token hash finds only its own row", func(t *testing.T) {
		hashes := map[string][]byte{
			"session.token_hash":                   a.sessionHash,
			"refresh_token.token_hash":             a.refreshHash,
			"recovery_token.token_hash":            a.recoveryHash,
			"setup_token.token_hash":               a.setupHash,
			"signup_verification.token_hash":       a.linkHash,
			"signup_verification.setup_token_hash": a.setupHash,
		}
		for table, columns := range tokenTables {
			for _, column := range columns {
				hash := hashes[table+"."+column]
				got := readColumn[[]byte](t, app, database.Scope{TokenHash: hash}, table, column)
				if len(got) != 1 || !slices.Equal(got[0], hash) {
					t.Errorf("%s by %s: %d rows, want exactly the one with that hash", table, column, len(got))
				}
			}
		}
	})

	t.Run("an email finds only its own account, whatever its case", func(t *testing.T) {
		got := readColumn[pgtype.UUID](t, app, database.Scope{Email: "owner@EXAMPLE.test"}, "account", "id")
		if !slices.Equal(got, []pgtype.UUID{a.account}) {
			t.Errorf("account by email: %v, want only A %v", got, a.account)
		}
	})

	t.Run("without a scope nothing is readable and nothing fails", func(t *testing.T) {
		for table := range accountTables {
			if got := readColumn[pgtype.UUID](t, app, database.Scope{}, table, accountTables[table]); len(got) != 0 {
				t.Errorf("%s: %d rows without a scope, want 0", table, len(got))
			}
		}
		for _, table := range []string{"setup_token", "signup_verification", "idempotency_record"} {
			if got := readColumn[int](t, app, database.Scope{}, table, "1"); len(got) != 0 {
				t.Errorf("%s: %d rows without a scope, want 0", table, len(got))
			}
		}
	})

	t.Run("idempotency_record follows the scope id (09.2 §3)", func(t *testing.T) {
		config, err := pgxpool.ParseConfig(d.URL("app_rw"))
		if err != nil {
			t.Fatal(err)
		}
		config.MaxConns = 1
		one, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(one.Close)
		household := randomUUID()
		insert := "INSERT INTO idempotency_record (scope_id, idempotency_key, command, request_hash, response) VALUES ($1, gen_random_uuid(), 'x', '\\x00', '\\x00')"
		for _, w := range []struct {
			scope database.Scope
			id    pgtype.UUID
		}{
			{database.Scope{AccountID: a.account}, a.account},
			{database.Scope{HouseholdID: household}, household},
		} {
			err := database.Write(ctx, one, w.scope, func(tx pgx.Tx) error {
				_, err := tx.Exec(ctx, insert, w.id)
				return err
			})
			if err != nil {
				t.Fatal(err)
			}
		}
		// The household write ran last on the same connection.
		if got := readColumn[pgtype.UUID](t, one, database.Scope{AccountID: a.account}, "idempotency_record", "scope_id"); !slices.Equal(got, []pgtype.UUID{a.account}) {
			t.Errorf("account A reads scope ids %v, want only %v", got, a.account)
		}
		if got := readColumn[pgtype.UUID](t, one, database.Scope{AccountID: b.account}, "idempotency_record", "scope_id"); len(got) != 0 {
			t.Errorf("account B reads scope ids %v, want none", got)
		}
	})

	t.Run("the security trail is append-only for app_rw", func(t *testing.T) {
		for _, statement := range []string{
			"UPDATE security_event SET detail = 'x'",
			"DELETE FROM security_event",
		} {
			err := database.Write(ctx, app, database.Scope{AccountID: a.account}, func(tx pgx.Tx) error {
				_, err := tx.Exec(ctx, statement)
				return err
			})
			if code := databasetest.PgErrorCode(err); code != "42501" {
				t.Errorf("%s: err = %v, want insufficient_privilege (42501)", statement, err)
			}
		}
	})

	t.Run("an email is unique without regard to case", func(t *testing.T) {
		id := randomUUID()
		err := database.Write(ctx, app, database.Scope{AccountID: id}, func(tx pgx.Tx) error {
			_, err := tx.Exec(ctx, "INSERT INTO account (id, email, created_at, signup_via) VALUES ($1, 'OWNER@example.TEST', now(), 'setup_token')", id)
			return err
		})
		if code := databasetest.PgErrorCode(err); code != "23505" {
			t.Errorf("second account for %s in another case: err = %v, want unique_violation (23505)", a.email, err)
		}
	})
}
