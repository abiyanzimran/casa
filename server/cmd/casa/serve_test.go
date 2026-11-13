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

// served gives serve its database credential and returns a mail directory.
func served(t *testing.T) (mailDir string) {
	t.Helper()
	withCredential(t, appDatabaseURL, databasetest.Migrated(t).URL("app_rw"))
	return t.TempDir()
}

// TestMain gives every test a web bundle, so no test depends on whether
// make bundle-copy has filled the embedded copy (plan decision 2).
func TestMain(m *testing.M) {
	embeddedBundle = fstest.MapFS{"dist/index.html": {Data: []byte("<!doctype html><title>Casa</title>")}}
	os.Exit(m.Run())
}

// startServe runs casa serve on a free port until the returned cancel is
// called; wait then returns its exit code and stderr.
func startServe(t *testing.T) (address string, cancel func(), wait func() (int, string)) {
	t.Helper()
	mailDir := served(t)
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
	return address, cancel, wait
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
		address, cancel, wait := startServe(t)
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
		address, cancel, wait := startServe(t)
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
		code, _, stderr := runCasa(t, "serve", "-mail-dir", served(t), "-listen", taken.Addr().String())
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
