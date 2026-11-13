// Package databasetest gives each test its own database in the PostgreSQL 18
// test cluster named by CASA_TEST_ADMIN_URL (14 §4.1 item 1: a real
// database, never a mock). A separate database per test keeps test packages
// that run at the same time from seeing each other's data.
//
// Without the cluster a test fails; it is never skipped, so a green run always
// means the database tests ran.
package databasetest

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"net/url"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // registers the "pgx" database/sql driver

	"github.com/abiyanzimran/casa/server/internal/schema"
)

// rolesLock serializes running roles.sql across test packages: CREATE ROLE and
// ALTER ROLE of the same role in two transactions at once fail.
const rolesLock = 7_202_611_130_001

// Database is one test database, owned by migrator.
type Database struct {
	Name string
	// AdminRole is the superuser of CASA_TEST_ADMIN_URL.
	AdminRole string
	admin     *url.URL
}

// Empty makes sure the roles of roles.sql exist, then creates an empty
// database owned by migrator, as the runbook does (see roles.sql). The
// database is dropped when the test ends.
func Empty(t testing.TB) Database {
	t.Helper()
	admin := adminURL(t)
	ctx := t.Context()
	conn, err := pgx.Connect(ctx, admin.String())
	if err != nil {
		t.Fatalf("connect to CASA_TEST_ADMIN_URL: %v; run make db-up", err)
	}
	t.Cleanup(func() { _ = conn.Close(context.Background()) })

	err = pgx.BeginFunc(ctx, conn, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, "SELECT pg_advisory_xact_lock($1)", rolesLock); err != nil {
			return err
		}
		_, err := tx.Exec(ctx, schema.Roles)
		return err
	})
	if err != nil {
		t.Fatalf("run roles.sql: %v", err)
	}

	d := Database{Name: "casa_test_" + randomHex(), AdminRole: admin.User.Username(), admin: admin}
	name := pgx.Identifier{d.Name}.Sanitize()
	if _, err := conn.Exec(ctx, "CREATE DATABASE "+name+" OWNER migrator"); err != nil {
		t.Fatalf("create database: %v", err)
	}
	// t.Context is already cancelled when cleanups run. Registered after the
	// connection's Close, so it runs before it; pools a test opens later are
	// closed before this.
	t.Cleanup(func() {
		if _, err := conn.Exec(context.Background(), "DROP DATABASE "+name+" WITH (FORCE)"); err != nil {
			t.Errorf("drop database %s: %v", d.Name, err)
		}
	})
	return d
}

// Migrated is Empty with every migration applied as migrator.
func Migrated(t testing.TB) Database {
	t.Helper()
	d := Empty(t)
	if _, err := schema.Up(t.Context(), d.SQL(t, "migrator")); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return d
}

// URL is CASA_TEST_ADMIN_URL with the user replaced by role and the database
// by d. The test roles have no password; the test cluster trusts local
// connections (14 §7.8).
func (d Database) URL(role string) string {
	u := *d.admin
	u.User = url.User(role)
	u.Path = "/" + d.Name
	return u.String()
}

// Pool opens a pgx pool on d as role, closed when the test ends.
func (d Database) Pool(t testing.TB, role string) *pgxpool.Pool {
	t.Helper()
	pool, err := pgxpool.New(t.Context(), d.URL(role))
	if err != nil {
		t.Fatalf("open pool as %s: %v", role, err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// SQL opens a database/sql handle on d as role, closed when the test ends.
func (d Database) SQL(t testing.TB, role string) *sql.DB {
	t.Helper()
	db, err := sql.Open("pgx", d.URL(role))
	if err != nil {
		t.Fatalf("open database/sql as %s: %v", role, err)
	}
	t.Cleanup(func() { _ = db.Close() })
	return db
}

func adminURL(t testing.TB) *url.URL {
	t.Helper()
	raw := os.Getenv("CASA_TEST_ADMIN_URL")
	if raw == "" {
		t.Fatal("CASA_TEST_ADMIN_URL is not set; run make db-up and make server-test")
	}
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatal("CASA_TEST_ADMIN_URL is not a valid URL")
	}
	if u.User.Username() == "" {
		t.Fatal("CASA_TEST_ADMIN_URL names no user")
	}
	return u
}

func randomHex() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b) // never returns an error (crypto/rand, Go 1.24 and later)
	return hex.EncodeToString(b)
}

// PgErrorCode returns the SQLSTATE of a PostgreSQL error in err's chain, or
// the empty string.
func PgErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
