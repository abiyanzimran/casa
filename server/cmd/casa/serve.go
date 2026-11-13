package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/abiyanzimran/casa/server/internal/api"
	"github.com/abiyanzimran/casa/server/internal/identity"
	"github.com/abiyanzimran/casa/server/internal/webbundle"
)

// embeddedBundle holds the web bundle under dist/; tests replace it so they
// do not depend on whether make bundle-copy has run.
var embeddedBundle fs.FS = webbundle.FS

const serveUsage = "usage: casa serve -mail-dir directory [-listen address] [-origin url] [-minimum-web-version MAJOR.MINOR.PATCH]\n"

// serve runs the API server until ctx ends, then shuts it down gracefully
// (09.4 §3; 14 §6.4; ADR-030 item 3). TLS on port 443 comes with FS-1 step 7.
// It prints only the listen address and errors, never request content (ES §4).
func serve(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("serve", flag.ContinueOnError)
	flags.SetOutput(io.Discard) // serveUsage is printed instead
	listen := flags.String("listen", "127.0.0.1:8080", "")
	// The base of emailed links (09.4 §2).
	origin := flags.String("origin", "http://127.0.0.1:8080", "")
	mailDir := flags.String("mail-dir", "", "")
	// ADR-030 item 9; lead ruling Q2: the minimum web version is server
	// configuration, raised by the operator on a release that fixes a hole.
	minimum := flags.String("minimum-web-version", api.MinimumWebVersion, "")
	if err := flags.Parse(args); err != nil || flags.NArg() != 0 || *mailDir == "" {
		_, _ = fmt.Fprint(stderr, serveUsage)
		return 2
	}
	if info, err := os.Stat(*mailDir); err != nil || !info.IsDir() {
		_, _ = fmt.Fprintf(stderr, "casa serve: -mail-dir: %s is not a directory\n%s", *mailDir, serveUsage)
		return 2
	}
	minimumWebVersion, err := api.ParseVersion(*minimum)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "casa serve: -minimum-web-version: %v\n%s", err, serveUsage)
		return 2
	}
	bundle, err := openBundle(stderr)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	pool, err := openAppPool(ctx)
	if err == nil {
		// Fail at start, not at the first request. pgx errors name the
		// host and user at most, never the password (ES §4).
		if err = pool.Ping(ctx); err != nil {
			pool.Close()
			err = fmt.Errorf("open the database: %w", err)
		}
	}
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	identityService := identity.New(pool, identity.Config{
		Origin: *origin, SendMail: writeMail(*mailDir), Now: time.Now, Password: identity.DefaultPasswordParams,
		Logger: slog.New(slog.NewJSONHandler(stderr, nil)), // JSON to journald (09.4 §3)
	})
	// closeAll ends the background work, then the pool. It runs only when
	// no handler can still use the pool: before serving, or after a clean
	// Shutdown. After a Shutdown that timed out, handlers may still run, so
	// the pool is left to the process exit (code-reviewer S4 on #32).
	closeAll := func() {
		identityService.Close()
		pool.Close()
	}
	apiServer := api.NewServer(minimumWebVersion, identityService.Authenticate)
	api.RegisterIdentityService(apiServer, pool, identityService)
	listener, err := new(net.ListenConfig).Listen(ctx, "tcp", *listen)
	if err != nil {
		closeAll()
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	_, _ = fmt.Fprintf(stdout, "casa serve: listening on %s\n", listener.Addr())
	server := &http.Server{
		Handler: api.NewMux(apiServer, bundle),
		// [Usulan]: bounds a client that never finishes its headers.
		ReadHeaderTimeout: 10 * time.Second,
		// step 7 (09.4 TLS): ReadTimeout and IdleTimeout before the binary faces the internet; no WriteTimeout (SSE)
	}
	served := make(chan error, 1)
	go func() { served <- server.Serve(listener) }()
	select {
	case err := <-served:
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	case <-ctx.Done():
	}
	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		_, _ = fmt.Fprintf(stderr, "casa serve: shutdown: %v\n", err)
		return 1
	}
	closeAll()
	if err := <-served; !errors.Is(err, http.ErrServerClosed) {
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	return 0
}

// openBundle reads the embedded web bundle (ADR-030 item 3). A binary built
// without it serves the API only and says so (plan decision 2, lead ruling
// Q5); the release binary is always built with make casa.
func openBundle(stderr io.Writer) (*api.Bundle, error) {
	files, err := fs.Sub(embeddedBundle, "dist")
	if err != nil {
		return nil, fmt.Errorf("open the web bundle: %w", err)
	}
	bundle, err := api.NewBundle(files)
	if errors.Is(err, api.ErrNoBundle) {
		_, _ = fmt.Fprintln(stderr, "casa serve: no web bundle in this binary; build it with make casa")
		return nil, nil
	}
	return bundle, err
}
