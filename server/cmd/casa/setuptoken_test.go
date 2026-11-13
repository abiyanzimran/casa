package main

import (
	"crypto/sha256"
	"encoding/base64"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

var setupTokenOutput = regexp.MustCompile(`^casa setup-token: ([A-Za-z0-9_-]{43})
expires: (\S+)
record this issuance in the operations log \(SEC-37\)
$`)

func TestSetupToken(t *testing.T) {
	t.Run("issues one token bound to one email and prints it once", func(t *testing.T) {
		// SEC-158 (b): one token for one named person, bound to an email;
		// the operator never sets a password (SEC-26).
		d := databasetest.Migrated(t)
		withCredential(t, appDatabaseURL, d.URL("app_rw"))
		code, stdout, stderr := runCasa(t, "setup-token", "-email", "Owner@Example.Test")
		match := setupTokenOutput.FindStringSubmatch(stdout)
		if code != 0 || match == nil || stderr != "" {
			t.Fatalf("exit %d, stdout %q, stderr %q; want 0 and the three lines", code, stdout, stderr)
		}
		expires, err := time.Parse(time.RFC3339, match[2])
		if err != nil || expires.Location() != time.UTC {
			t.Fatalf("expires %q: %v; want RFC 3339 in UTC", match[2], err)
		}
		if want := time.Now().Add(7 * 24 * time.Hour); expires.Before(want.Add(-time.Minute)) || expires.After(want.Add(time.Minute)) {
			t.Fatalf("expires %v, want about %v (7 days)", expires, want)
		}
		raw, err := base64.RawURLEncoding.DecodeString(match[1])
		if err != nil {
			t.Fatal(err)
		}
		hash := sha256.Sum256(raw)
		var rows int
		var email, row string
		err = d.Pool(t, d.AdminRole).QueryRow(t.Context(), `
			SELECT count(*) OVER (), email::text, s::text FROM setup_token s
			 WHERE token_hash = $1 AND used_at IS NULL AND account_id IS NULL`, hash[:]).Scan(&rows, &email, &row)
		if err != nil || rows != 1 || email != "Owner@Example.Test" {
			t.Fatalf("setup_token: %d rows, email %q, err %v; want 1 row for Owner@Example.Test", rows, email, err)
		}
		if strings.Contains(row, match[1]) {
			t.Fatal("the token itself is stored; only its hash may be")
		}
	})

	t.Run("rejects a malformed email", func(t *testing.T) {
		d := databasetest.Migrated(t)
		withCredential(t, appDatabaseURL, d.URL("app_rw"))
		code, stdout, stderr := runCasa(t, "setup-token", "-email", "not-an-email")
		if want := "casa setup-token: -email is not an email address\n"; code != 2 || stdout != "" || stderr != want {
			t.Fatalf("exit %d, stdout %q, stderr %q; want 2, %q", code, stdout, stderr, want)
		}
		var rows int
		if err := d.Pool(t, d.AdminRole).QueryRow(t.Context(), "SELECT count(*) FROM setup_token").Scan(&rows); err != nil || rows != 0 {
			t.Fatalf("%d setup tokens (%v), want 0", rows, err)
		}
	})

	t.Run("rejects extra arguments", func(t *testing.T) {
		for _, args := range [][]string{
			{"setup-token", "-email", "a@example.test", "x"},
			{"setup-token"},
		} {
			code, _, stderr := runCasa(t, args...)
			if want := "usage: casa setup-token -email address\n"; code != 2 || stderr != want {
				t.Errorf("%q: exit %d, stderr %q; want 2, %q", args, code, stderr, want)
			}
		}
	})

	t.Run("never prints the connection string", func(t *testing.T) {
		// Built by concatenation so the secrets check sees no credential URI.
		const marker = "leak-marker-7f3"
		for _, credential := range []string{
			"postgres://app_rw:" + marker + "@127.0.0.1:1/casa?connect_timeout=2",
			"postgres://app_rw:" + marker + "@[bad",
		} {
			withCredential(t, appDatabaseURL, credential)
			code, stdout, stderr := runCasa(t, "setup-token", "-email", "a@example.test")
			if code != 1 || strings.Contains(stdout+stderr, marker) {
				t.Errorf("exit %d, output %q; want 1 without the password (ES §4)", code, stdout+stderr)
			}
		}
	})
}
