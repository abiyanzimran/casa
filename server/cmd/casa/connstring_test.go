package main

import (
	"strings"
	"testing"
)

func TestConnString(t *testing.T) {
	t.Run("returns a parsable connection string", func(t *testing.T) {
		const want = "postgres://app_rw@127.0.0.1:5432/casa?pool_max_conns=4"
		withCredential(t, "db-url", want)
		if got, err := connString("db-url"); err != nil || got != want {
			t.Fatalf("got %q, %v; want %q", got, err, want)
		}
	})
	t.Run("refuses an empty credential", func(t *testing.T) {
		// pgx would fall back to the PG* environment variables.
		withCredential(t, "db-url", "")
		if _, err := connString("db-url"); err == nil || err.Error() != "credential db-url is empty" {
			t.Fatalf("err = %v, want the empty credential refused", err)
		}
	})
	t.Run("never quotes an unparsable connection string", func(t *testing.T) {
		// ES §4: the parse error can quote the password.
		for _, bad := range []string{"postgres://u:s3cret-pw@host:notaport/db", "postgres://u:s3cret-pw@h/db?pool_max_conns=x"} {
			withCredential(t, "db-url", bad)
			_, err := connString("db-url")
			if err == nil || strings.Contains(err.Error(), "s3cret") {
				t.Fatalf("%q: err = %v, want a refusal without the password", bad, err)
			}
		}
	})
}
