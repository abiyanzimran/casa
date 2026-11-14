package main

import (
	"bufio"
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
	"time"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

// served gives serve its database, job database and operator-name
// credentials and returns a mail directory and the database.
func served(t *testing.T) (mailDir string, d databasetest.Database) {
	t.Helper()
	d = databasetest.Migrated(t)
	withCredentials(t, map[string]string{ //nolint:gosec // G101: test database URLs and a test name, no secret
		appDatabaseURL: d.URL("app_rw"), jobDatabaseURL: d.URL("job_runner"), operatorNameCredential: "Pengelola Uji",
	})
	return t.TempDir(), d
}

// TestMain gives every test a web bundle, so no test depends on whether
// make bundle-copy has filled the embedded copy (plan decision 2).
func TestMain(m *testing.M) {
	embeddedBundle = fstest.MapFS{"dist/index.html": {Data: []byte("<!doctype html><title>Casa</title>")}}
	os.Exit(m.Run())
}

// startServe runs casa serve on a free port until the returned cancel is
// called; wait then returns its exit code and stderr. d is its database.
func startServe(t *testing.T) (address string, cancel func(), wait func() (int, string), d databasetest.Database) {
	t.Helper()
	mailDir, d := served(t)
	ctx, cancel := context.WithCancel(t.Context())
	t.Cleanup(cancel)
	stdoutReader, stdoutWriter := io.Pipe()
	var stderr strings.Builder
	done := make(chan int, 1)
	go func() {
		done <- run(ctx, []string{"serve", "-mail-dir", mailDir, "-listen", "127.0.0.1:0"}, stdoutWriter, &stderr)
		_ = stdoutWriter.Close()
	}()
	line, err := bufio.NewReader(stdoutReader).ReadString('\n')
	if err != nil {
		t.Fatalf("reading stdout: %v", err)
	}
	address, ok := strings.CutPrefix(strings.TrimSuffix(line, "\n"), "casa serve: listening on ")
	if !ok || !strings.HasPrefix(address, "127.0.0.1:") {
		t.Fatalf("stdout = %q, want \"casa serve: listening on 127.0.0.1:<port>\\n\"", line)
	}
	wait = func() (int, string) {
		t.Helper()
		select {
		case code := <-done:
			return code, stderr.String()
		case <-time.After(15 * time.Second):
			t.Fatal("serve did not return within 15 s of the context ending")
			return 0, ""
		}
	}
	return address, cancel, wait, d
}

// get fetches path from the server at address.
func get(t *testing.T, address, path string) (*http.Response, string) {
	t.Helper()
	res, err := http.Get("http://" + address + path)
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	return res, string(body)
}

func TestServe(t *testing.T) {
	t.Run("serves /healthz until the context ends", func(t *testing.T) {
		address, cancel, wait, _ := startServe(t)
		if res, body := get(t, address, "/healthz"); res.StatusCode != http.StatusOK || body != "ok\n" {
			t.Fatalf("GET /healthz = %d %q, want 200 \"ok\\n\"", res.StatusCode, body)
		}
		if res, body := get(t, address, "/"); res.StatusCode != http.StatusOK || !strings.Contains(body, "<title>Casa</title>") {
			t.Fatalf("GET / = %d %q, want 200 and the app page", res.StatusCode, body)
		}
		cancel()
		if code, stderr := wait(); code != 0 || stderr != "" {
			t.Fatalf("exit code = %d, stderr %q; want 0 and nothing", code, stderr)
		}
	})
	t.Run("serves the API without a web bundle and says so", func(t *testing.T) {
		// Plan decision 2 (lead ruling Q5): the API still serves; the static
		// routes answer 404, with the security headers (SEC-166).
		withBundle := embeddedBundle
		embeddedBundle = fstest.MapFS{"dist/.gitkeep": {Data: []byte("")}}
		t.Cleanup(func() { embeddedBundle = withBundle })
		address, cancel, wait, _ := startServe(t)
		if res, body := get(t, address, "/healthz"); res.StatusCode != http.StatusOK || body != "ok\n" {
			t.Fatalf("GET /healthz = %d %q, want 200 \"ok\\n\"", res.StatusCode, body)
		}
		res, body := get(t, address, "/")
		if res.StatusCode != http.StatusNotFound || res.Header.Get("Content-Security-Policy") == "" ||
			res.Header.Get("X-Content-Type-Options") != "nosniff" || res.Header.Get("Cache-Control") != "no-store" {
			t.Fatalf("GET / = %d %q, header %v; want 404 with the security headers", res.StatusCode, body, res.Header)
		}
		cancel()
		want := "casa serve: no web bundle in this binary; build it with make casa\n"
		if code, stderr := wait(); code != 0 || stderr != want {
			t.Fatalf("exit code = %d, stderr %q; want 0 and %q", code, stderr, want)
		}
	})
	t.Run("fails when the address is in use", func(t *testing.T) {
		taken, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = taken.Close() }()
		mailDir, _ := served(t)
		code, _, stderr := runCasa(t, "serve", "-mail-dir", mailDir, "-listen", taken.Addr().String())
		if code != 1 || !strings.HasPrefix(stderr, "casa serve: listen tcp 127.0.0.1:") {
			t.Fatalf("exit %d, stderr %q; want 1 and \"casa serve: listen tcp 127.0.0.1:...\"", code, stderr)
		}
	})
	t.Run("rejects extra arguments and unknown flags", func(t *testing.T) {
		for _, args := range [][]string{{"serve", "extra"}, {"serve", "-port", "1"}} {
			code, _, stderr := runCasa(t, args...)
			if code != 2 || stderr != serveUsage {
				t.Fatalf("%q: exit %d, stderr %q; want 2 and %q", args, code, stderr, serveUsage)
			}
		}
	})
	t.Run("rejects a minimum web version that is not MAJOR.MINOR.PATCH", func(t *testing.T) {
		// ADR-030 item 9; lead ruling Q2: the minimum is server configuration.
		code, _, stderr := runCasa(t, "serve", "-mail-dir", t.TempDir(), "-listen", "127.0.0.1:0", "-minimum-web-version", "1.2")
		want := "casa serve: -minimum-web-version: version \"1.2\" is not MAJOR.MINOR.PATCH\n" + serveUsage
		if code != 2 || stderr != want {
			t.Fatalf("exit %d, stderr %q; want 2 and %q", code, stderr, want)
		}
	})
	t.Run("rejects an -origin that is not a bare origin", func(t *testing.T) {
		// Plan 1.5 Task 4: the origin check compares Origin with this value
		// as a string (SEC-162 (a)), so it must be scheme://host[:port] as a
		// browser sends it; plain http only on the loopback host.
		t.Setenv("CREDENTIALS_DIRECTORY", t.TempDir()) // an accepted -origin fails next, on the credential
		for _, c := range []struct {
			origin   string
			accepted bool
		}{
			{"https://a.example/x", false},
			{"https://a.example/", false},
			{"https://a.example?q", false},
			{"https://a.example#f", false},
			{"ftp://a", false},
			{"http://a.example", false},
			{"https://u@a.example", false},
			{"https://A.example", false},
			{"https://", false},
			{"a.example", false},
			{"https://a.example:", false},
			{"https://a.example:443", false},
			{"http://localhost:80", false},
			{"https://a.example:0", false},
			{"https://a.example:65536", false},
			{"http://localhost:08080", false},
			{"https://a.example.", false},
			{"https://café.example", false},
			{"https://xn--caf-dma.example", true},
			{"http://localhost", true},
			{"https://a.example:65535", true},
			{"http://localhost:8080", true},
			{"http://127.0.0.1:8080", true},
			{"http://[::1]:8080", true},
			{"https://a.example", true},
			{"https://a.example:8443", true},
		} {
			code, _, stderr := runCasa(t, "serve", "-mail-dir", t.TempDir(), "-listen", "127.0.0.1:0", "-origin", c.origin)
			switch {
			case c.accepted && (code != 1 || !strings.HasPrefix(stderr, "casa serve: read credential")):
				t.Errorf("-origin %q: exit %d, stderr %q; want it accepted", c.origin, code, stderr)
			case !c.accepted && (code != 2 || !strings.HasPrefix(stderr, "casa serve: -origin: ") || !strings.HasSuffix(stderr, serveUsage)):
				t.Errorf("-origin %q: exit %d, stderr %q; want 2 and \"casa serve: -origin: …\"", c.origin, code, stderr)
			}
		}
	})
	t.Run("requires -mail-dir, an existing directory", func(t *testing.T) {
		code, _, stderr := runCasa(t, "serve", "-listen", "127.0.0.1:0")
		if code != 2 || stderr != serveUsage {
			t.Fatalf("exit %d, stderr %q; want 2 and %q", code, stderr, serveUsage)
		}
		missing := filepath.Join(t.TempDir(), "missing")
		code, _, stderr = runCasa(t, "serve", "-mail-dir", missing, "-listen", "127.0.0.1:0")
		if want := "casa serve: -mail-dir: " + missing + " is not a directory\n" + serveUsage; code != 2 || stderr != want {
			t.Fatalf("exit %d, stderr %q; want 2 and %q", code, stderr, want)
		}
	})
	t.Run("fails without the database credential", func(t *testing.T) {
		t.Setenv("CREDENTIALS_DIRECTORY", t.TempDir())
		code, _, stderr := runCasa(t, "serve", "-mail-dir", t.TempDir(), "-listen", "127.0.0.1:0")
		if code != 1 || !strings.HasPrefix(stderr, "casa serve: read credential app-database-url:") {
			t.Fatalf("exit %d, stderr %q; want 1 and the missing credential", code, stderr)
		}
	})
	t.Run("fails without the operator-name credential", func(t *testing.T) {
		// 09.3 §1 "Mekanisme": {pengelola} of the operator notice is filled
		// from this credential; without it no notice can be shown (SEC-34).
		database := databasetest.Migrated(t).URL("app_rw")
		withCredential(t, appDatabaseURL, database)
		code, _, stderr := runCasa(t, "serve", "-mail-dir", t.TempDir(), "-listen", "127.0.0.1:0")
		if code != 1 || !strings.HasPrefix(stderr, "casa serve: read credential operator-name:") {
			t.Fatalf("exit %d, stderr %q; want 1 and the missing credential", code, stderr)
		}
		withCredentials(t, map[string]string{appDatabaseURL: database, operatorNameCredential: " "})
		code, _, stderr = runCasa(t, "serve", "-mail-dir", t.TempDir(), "-listen", "127.0.0.1:0")
		if want := "casa serve: credential operator-name is empty\n"; code != 1 || stderr != want {
			t.Fatalf("exit %d, stderr %q; want 1 and %q", code, stderr, want)
		}
	})
	t.Run("serve refuses to start without the job database credential", func(t *testing.T) {
		// Plan 1.7 decision 4: the jobs runner claims as job_runner (09.2 §3).
		database := databasetest.Migrated(t)
		withCredentials(t, map[string]string{ //nolint:gosec // G101: a test database URL and a test name, no secret
			appDatabaseURL: database.URL("app_rw"), operatorNameCredential: "Pengelola Uji",
		})
		code, _, stderr := runCasa(t, "serve", "-mail-dir", t.TempDir(), "-listen", "127.0.0.1:0")
		if code != 1 || !strings.HasPrefix(stderr, "casa serve: read credential job-database-url:") {
			t.Fatalf("exit %d, stderr %q; want 1 and the missing credential", code, stderr)
		}
		fake := "s3cret-pw"
		withCredentials(t, map[string]string{ //nolint:gosec // G101: a fake password for an unreachable database
			appDatabaseURL: database.URL("app_rw"), operatorNameCredential: "Pengelola Uji",
			jobDatabaseURL: "postgres://job_runner:" + fake + "@127.0.0.1:1/casa?connect_timeout=2",
		})
		code, stdout, stderr := runCasa(t, "serve", "-mail-dir", t.TempDir(), "-listen", "127.0.0.1:0")
		if code != 1 || !strings.HasPrefix(stderr, "casa serve: open the job database:") || strings.Contains(stdout+stderr, fake) {
			t.Fatalf("exit %d, stdout %q, stderr %q; want 1, the error, and no password", code, stdout, stderr)
		}
	})
	t.Run("serve runs due jobs", func(t *testing.T) {
		// 09.3 §9; plan 1.7 decision 1: a test invitation past its
		// expires_at, with its job, expires without any request.
		_, cancel, wait, d := startServe(t)
		admin := d.Pool(t, d.AdminRole)
		var invitation string
		if err := admin.QueryRow(t.Context(), `
			WITH h AS (INSERT INTO household (name, time_zone, created_at) VALUES ('Uji', 'Asia/Jakarta', now() - interval '8 days') RETURNING id),
			     m AS (INSERT INTO membership (household_id, display_name, role, permission_set, joined_at)
			           SELECT id, 'Ayu', 'household_owner', 'two_equal_adults', now() - interval '8 days' FROM h RETURNING household_id, member_id),
			     i AS (INSERT INTO invitation (household_id, inviter, state, created_at, expires_at)
			           SELECT household_id, member_id, 'open', now() - interval '8 days', now() - interval '1 day' FROM m RETURNING household_id, id, expires_at),
			     j AS (INSERT INTO job (household_id, kind, run_at, unique_key, payload)
			           SELECT household_id, 'household.expire_invitation', expires_at, id::text, uuid_send(id) FROM i)
			SELECT id::text FROM i`).Scan(&invitation); err != nil {
			t.Fatal(err)
		}
		deadline := time.Now().Add(5 * time.Second)
		for state := ""; state != "expired"; {
			if time.Now().After(deadline) {
				t.Fatalf("invitation state %q 5 s later, want expired", state)
			}
			time.Sleep(20 * time.Millisecond)
			if err := admin.QueryRow(t.Context(), "SELECT state FROM invitation WHERE id = $1", invitation).Scan(&state); err != nil {
				t.Fatal(err)
			}
		}
		cancel()
		if code, stderr := wait(); code != 0 || stderr != "" {
			t.Fatalf("exit code = %d, stderr %q; want 0 and nothing", code, stderr)
		}
	})
	t.Run("serve exits when the jobs runner loses the database", func(t *testing.T) {
		// Plan 1.7 Task 2: Run returns a database failure to its caller;
		// serve stops and exits 1 so its supervisor restarts it, without the
		// error's text (SEC-44).
		_, _, wait, d := startServe(t)
		admin := d.Pool(t, d.AdminRole)
		deadline := time.Now().Add(5 * time.Second)
		for terminated := 0; terminated == 0; {
			if time.Now().After(deadline) {
				t.Fatal("no listening job_runner connection 5 s after the start")
			}
			time.Sleep(20 * time.Millisecond)
			if err := admin.QueryRow(t.Context(), `
				SELECT count(pg_terminate_backend(pid)) FROM pg_stat_activity
				 WHERE usename = 'job_runner' AND datname = current_database() AND query LIKE '%LISTEN jobs%'`).Scan(&terminated); err != nil {
				t.Fatal(err)
			}
		}
		want := "casa serve: the jobs runner stopped on a database error\n"
		if code, stderr := wait(); code != 1 || stderr != want {
			t.Fatalf("exit code = %d, stderr %q; want 1 and %q", code, stderr, want)
		}
	})
	t.Run("never prints the connection string", func(t *testing.T) {
		// ES §4: a database that cannot be reached is reported without the
		// password.
		fake := "s3cret-pw" // built up, so it does not read as a credential
		withCredential(t, appDatabaseURL, "postgres://app_rw:"+fake+"@127.0.0.1:1/casa?connect_timeout=2")
		code, stdout, stderr := runCasa(t, "serve", "-mail-dir", t.TempDir(), "-listen", "127.0.0.1:0")
		if code != 1 || !strings.HasPrefix(stderr, "casa serve: open the database:") || strings.Contains(stdout+stderr, fake) {
			t.Fatalf("exit %d, stdout %q, stderr %q; want 1, the error, and no password", code, stdout, stderr)
		}
	})
}
