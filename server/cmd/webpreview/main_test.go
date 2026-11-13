package main

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"testing/fstest"
	"time"
)

func TestWebPreview(t *testing.T) {
	t.Run("serves the bundle with the production headers until the context ends", func(t *testing.T) {
		// Plan decision 6: the e2e tests run on the embedded bundle bytes
		// and the production headers (SEC-166).
		embeddedBundle = fstest.MapFS{"dist/index.html": {Data: []byte("<!doctype html><title>Casa</title>")}}
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()
		stdoutReader, stdoutWriter := io.Pipe()
		var stderr bytes.Buffer
		done := make(chan int, 1)
		go func() {
			done <- run(ctx, []string{"-listen", "127.0.0.1:0"}, stdoutWriter, &stderr)
			_ = stdoutWriter.Close()
		}()
		line, err := bufio.NewReader(stdoutReader).ReadString('\n')
		if err != nil {
			t.Fatalf("reading stdout: %v", err)
		}
		url, ok := strings.CutPrefix(strings.TrimSuffix(line, "\n"), "webpreview: listening on ")
		if !ok || !strings.HasPrefix(url, "http://127.0.0.1:") {
			t.Fatalf("stdout = %q, want \"webpreview: listening on http://127.0.0.1:<port>\\n\"", line)
		}
		res, err := http.Get(url + "/money")
		if err != nil {
			t.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		_ = res.Body.Close()
		if err != nil || res.StatusCode != http.StatusOK || !strings.Contains(string(body), "<title>Casa</title>") ||
			!strings.HasPrefix(res.Header.Get("Content-Security-Policy"), "default-src 'none';") {
			t.Fatalf("GET /money = %d %q (%v), header %v; want the app page with the CSP", res.StatusCode, body, err, res.Header)
		}
		cancel()
		select {
		case code := <-done:
			if code != 0 || stderr.Len() != 0 {
				t.Fatalf("exit code = %d, stderr %q; want 0 and nothing", code, stderr.String())
			}
		case <-time.After(15 * time.Second):
			t.Fatal("webpreview did not return within 15 s of the context ending")
		}
	})
	t.Run("fails without a web bundle", func(t *testing.T) {
		embeddedBundle = fstest.MapFS{"dist/.gitkeep": {Data: []byte("")}}
		var stdout, stderr bytes.Buffer
		code := run(t.Context(), []string{"-listen", "127.0.0.1:0"}, &stdout, &stderr)
		if want := "webpreview: no web bundle; run make bundle-copy\n"; code != 1 || stderr.String() != want || stdout.Len() != 0 {
			t.Fatalf("exit %d, stdout %q, stderr %q; want 1 and %q", code, stdout.String(), stderr.String(), want)
		}
	})
}
