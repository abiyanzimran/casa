package main

import (
	"bufio"
	"context"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	t.Run("serves /healthz until the context ends", func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()
		stdoutReader, stdoutWriter := io.Pipe()
		var stderr strings.Builder
		done := make(chan int, 1)
		go func() {
			done <- run(ctx, []string{"serve", "-listen", "127.0.0.1:0"}, stdoutWriter, &stderr)
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
		res, err := http.Get("http://" + address + "/healthz")
		if err != nil {
			t.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		_ = res.Body.Close()
		if err != nil || res.StatusCode != http.StatusOK || string(body) != "ok\n" {
			t.Fatalf("GET /healthz = %d %q (%v), want 200 \"ok\\n\"", res.StatusCode, body, err)
		}
		cancel()
		select {
		case code := <-done:
			if code != 0 {
				t.Fatalf("exit code = %d, stderr %q; want 0", code, stderr.String())
			}
		case <-time.After(15 * time.Second):
			t.Fatal("serve did not return within 15 s of the context ending")
		}
	})
	t.Run("fails when the address is in use", func(t *testing.T) {
		taken, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = taken.Close() }()
		code, _, stderr := runCasa(t, "serve", "-listen", taken.Addr().String())
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
		code, _, stderr := runCasa(t, "serve", "-listen", "127.0.0.1:0", "-minimum-web-version", "1.2")
		want := "casa serve: -minimum-web-version: version \"1.2\" is not MAJOR.MINOR.PATCH\n" + serveUsage
		if code != 2 || stderr != want {
			t.Fatalf("exit %d, stderr %q; want 2 and %q", code, stderr, want)
		}
	})
}
