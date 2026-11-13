package api_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

// holdTables takes LOCK TABLE … IN EXCLUSIVE MODE on the given tables in an
// admin transaction: INSERT, UPDATE and DELETE wait, SELECT does not.
// waitBlocked(n) loops on pg_locks until n backends wait on one of those
// tables (deadline 10 s, no fixed sleep); release commits the admin
// transaction so the waiting writes run at once (plan 1.6 decision 21).
func holdTables(t *testing.T, db databasetest.Database, tables ...string) (waitBlocked func(n int), release func()) {
	t.Helper()
	admin := db.Pool(t, db.AdminRole)
	tx, err := admin.Begin(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = tx.Rollback(context.Background()) }) // a no-op after release
	quoted := make([]string, len(tables))
	for i, table := range tables {
		quoted[i] = pgx.Identifier{table}.Sanitize()
	}
	if _, err := tx.Exec(t.Context(), "LOCK TABLE "+strings.Join(quoted, ", ")+" IN EXCLUSIVE MODE"); err != nil {
		t.Fatal(err)
	}
	waitBlocked = func(n int) {
		t.Helper()
		deadline := time.Now().Add(10 * time.Second)
		for {
			var waiting int
			if err := admin.QueryRow(t.Context(), `
				SELECT count(DISTINCT pid) FROM pg_locks
				 WHERE NOT granted AND relation = ANY($1::text[]::regclass[])`, tables).Scan(&waiting); err != nil {
				t.Fatal(err)
			}
			if waiting >= n {
				return
			}
			if time.Now().After(deadline) {
				t.Fatalf("%d backends wait on %v after 10 s, want %d", waiting, tables, n)
			}
			time.Sleep(5 * time.Millisecond) // a poll interval, not a wait for the outcome
		}
	}
	release = func() {
		t.Helper()
		if err := tx.Commit(t.Context()); err != nil {
			t.Fatal(err)
		}
	}
	return waitBlocked, release
}
