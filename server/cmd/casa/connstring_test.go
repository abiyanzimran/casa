package main

import (
	"strings"
	"testing"
)

func TestConnString(t *testing.T) {
	t.Run("returns a parsable connection string", func(t *testing.T) {
		const want = "postgres://app_rw@127.0.0.1:5432/casa?pool_max_conns=4"
		withCredential(t, "db-url", want)
		if got, err := connString("db-url", true); err != nil || got != want {
			t.Fatalf("got %q, %v; want %q", got, err, want)
		}
	})
	t.Run("refuses an empty credential", func(t *testing.T) {
		// pgx would fall back to the PG* environment variables.
		withCredential(t, "db-url", "")
		if _, err := connString("db-url", true); err == nil || err.Error() != "credential db-url is empty" {
			t.Fatalf("err = %v, want the empty credential refused", err)
		}
	})
	t.Run("never quotes an unparsable connection string", func(t *testing.T) {
		// ES §4: the parse error can quote the password.
		for _, bad := range []string{"postgres://u:s3cret-pw@host:notaport/db", "postgres://u:s3cret-pw@h/db?pool_max_conns=x"} {
			withCredential(t, "db-url", bad)
			_, err := connString("db-url", true)
			if err == nil || strings.Contains(err.Error(), "s3cret") {
				t.Fatalf("%q: err = %v, want a refusal without the password", bad, err)
			}
		}
	})
	t.Run("a credential for one connection refuses pool parameters", func(t *testing.T) {
		// casa migrate connects through database/sql, where pgx would send a
		// pool_ parameter to the server as a setting; refuse it at the
		// check, without the password (ES §4).
		fake := "s3cret-pw" // built up, so it does not read as a credential
		s := "postgres://u:" + fake + "@127.0.0.1:1/db?pool_max_conns=4"
		withCredential(t, "db-url", s)
		if got, err := connString("db-url", true); err != nil || got != s {
			t.Fatalf("for a pool: %q, %v; want it accepted", got, err)
		}
		_, err := connString("db-url", false)
		if err == nil || strings.Contains(err.Error(), "s3cret") {
			t.Fatalf("for one connection: err = %v, want a refusal without the password", err)
		}
		withCredential(t, migratorDatabaseURL, s)
		code, _, stderr := runCasa(t, "migrate")
		if want := "casa migrate: credential migrator-database-url has pool_ parameters, which do not apply to one connection\n"; code != 1 || stderr != want {
			t.Fatalf("migrate: code %d, stderr %q; want 1, %q", code, stderr, want)
		}
	})
}
