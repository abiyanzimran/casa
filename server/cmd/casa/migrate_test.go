package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
	"github.com/abiyanzimran/casa/server/internal/schema"
)

func TestMigrate(t *testing.T) {
	t.Run("applies every migration and reports the version", func(t *testing.T) {
		d := databasetest.Empty(t)
		withCredential(t, d.URL("migrator"))
		p, err := schema.NewProvider(d.SQL(t, "migrator"))
		if err != nil {
			t.Fatal(err)
		}
		sources := p.ListSources()
		want := fmt.Sprintf("casa migrate: database at version %d\n", sources[len(sources)-1].Version)
		for _, attempt := range []string{"first", "second"} {
			code, stdout, stderr := runCasa(t, "migrate")
			if code != 0 || stdout != want {
				t.Fatalf("%s run: exit %d, stdout %q, stderr %q; want 0, %q", attempt, code, stdout, stderr, want)
			}
		}
	})

	t.Run("refuses a role other than migrator", func(t *testing.T) {
		d := databasetest.Empty(t)
		withCredential(t, d.URL(d.AdminRole))
		code, _, stderr := runCasa(t, "migrate")
		want := "casa migrate: must run as role migrator, connected as " + d.AdminRole
		if code != 1 || !strings.Contains(stderr, want) {
			t.Fatalf("exit %d, stderr %q; want 1 and %q", code, stderr, want)
		}
	})

	t.Run("never prints the connection string", func(t *testing.T) {
		// Built by concatenation so the secrets check sees no credential URI.
		const marker = "leak-marker-7f3"
		for _, credential := range []string{
			"postgres://migrator:" + marker + "@127.0.0.1:1/casa?connect_timeout=2",
			"postgres://migrator:" + marker + "@[bad",
		} {
			withCredential(t, credential)
			code, stdout, stderr := runCasa(t, "migrate")
			if code != 1 || strings.Contains(stdout+stderr, marker) {
				t.Errorf("exit %d, output %q; want 1 without the password (ES §4)", code, stdout+stderr)
			}
		}
	})

	t.Run("fails on an empty credential", func(t *testing.T) {
		// An empty connection string would fall back to the PG* environment.
		withCredential(t, "\n")
		code, _, stderr := runCasa(t, "migrate")
		if want := "casa migrate: credential migrator-database-url is empty\n"; code != 1 || stderr != want {
			t.Fatalf("exit %d, stderr %q; want 1, %q", code, stderr, want)
		}
	})

	t.Run("fails without CREDENTIALS_DIRECTORY", func(t *testing.T) {
		t.Setenv("CREDENTIALS_DIRECTORY", "")
		code, _, stderr := runCasa(t, "migrate")
		want := "casa migrate: CREDENTIALS_DIRECTORY is not set; casa reads secrets from systemd LoadCredential= (09.4 §1)\n"
		if code != 1 || stderr != want {
			t.Fatalf("exit %d, stderr %q; want 1, %q", code, stderr, want)
		}
	})

	t.Run("rejects extra arguments", func(t *testing.T) {
		code, _, stderr := runCasa(t, "migrate", "down")
		if code != 2 || stderr != "usage: casa migrate\n" {
			t.Fatalf("exit %d, stderr %q; want 2, %q", code, stderr, "usage: casa migrate\n")
		}
	})
}

// withCredential writes credential as systemd LoadCredential= would.
func withCredential(t *testing.T, credential string) {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "migrator-database-url"), []byte(credential+"\n"), 0o400); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CREDENTIALS_DIRECTORY", dir)
}

func runCasa(t *testing.T, args ...string) (code int, stdout, stderr string) {
	t.Helper()
	var out, errOut bytes.Buffer
	code = run(t.Context(), args, &out, &errOut)
	return code, out.String(), errOut.String()
}
