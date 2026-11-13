// Command webpreview serves the embedded web bundle with the production
// security headers, without the API or a database, for the Playwright tests
// (plan decision 6): every e2e test runs on the bytes and the CSP the casa
// binary serves (SEC-166; ADR-030 item 7). It is never deployed.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abiyanzimran/casa/server/internal/api"
	"github.com/abiyanzimran/casa/server/internal/webbundle"
)

const usage = "usage: webpreview [-listen address]\n"

// embeddedBundle holds the web bundle under dist/; tests replace it so they
// do not depend on whether make bundle-copy has run.
var embeddedBundle fs.FS = webbundle.FS

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	code := run(ctx, os.Args[1:], os.Stdout, os.Stderr)
	stop()
	os.Exit(code)
}

// run serves the bundle until ctx ends and returns the exit code: 0 after a
// clean shutdown, 1 on failure, 2 for a usage error.
func run(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("webpreview", flag.ContinueOnError)
	flags.SetOutput(io.Discard) // usage is printed instead
	listen := flags.String("listen", "127.0.0.1:4173", "")
	if err := flags.Parse(args); err != nil || flags.NArg() != 0 {
		_, _ = fmt.Fprint(stderr, usage)
		return 2
	}
	files, err := fs.Sub(embeddedBundle, "dist")
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "webpreview: %v\n", err)
		return 1
	}
	bundle, err := api.NewBundle(files)
	if errors.Is(err, api.ErrNoBundle) {
		_, _ = fmt.Fprint(stderr, "webpreview: no web bundle; run make bundle-copy\n")
		return 1
	}
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "webpreview: %v\n", err)
		return 1
	}
	listener, err := new(net.ListenConfig).Listen(ctx, "tcp", *listen)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "webpreview: %v\n", err)
		return 1
	}
	_, _ = fmt.Fprintf(stdout, "webpreview: listening on http://%s\n", listener.Addr())
	server := &http.Server{
		Handler: api.WithSecurityHeaders(bundle, bundle),
		// Bounds a client that never finishes its headers, as casa serve.
		ReadHeaderTimeout: 10 * time.Second,
	}
	served := make(chan error, 1)
	go func() { served <- server.Serve(listener) }()
	select {
	case err := <-served:
		_, _ = fmt.Fprintf(stderr, "webpreview: %v\n", err)
		return 1
	case <-ctx.Done():
	}
	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		_, _ = fmt.Fprintf(stderr, "webpreview: shutdown: %v\n", err)
		return 1
	}
	if err := <-served; !errors.Is(err, http.ErrServerClosed) {
		_, _ = fmt.Fprintf(stderr, "webpreview: %v\n", err)
		return 1
	}
	return 0
}
