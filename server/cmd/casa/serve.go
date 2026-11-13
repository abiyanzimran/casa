package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/abiyanzimran/casa/server/internal/api"
)

const serveUsage = "usage: casa serve [-listen address] [-minimum-web-version MAJOR.MINOR.PATCH]\n"

// serve runs the API server until ctx ends, then shuts it down gracefully
// (09.4 §3; 14 §6.4; ADR-030 item 3). TLS on port 443 comes with FS-1 step 7.
// It prints only the listen address and errors, never request content (ES §4).
func serve(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("serve", flag.ContinueOnError)
	flags.SetOutput(io.Discard) // serveUsage is printed instead
	listen := flags.String("listen", "127.0.0.1:8080", "")
	// ADR-030 item 9; lead ruling Q2: the minimum web version is server
	// configuration, raised by the operator on a release that fixes a hole.
	minimum := flags.String("minimum-web-version", api.MinimumWebVersion, "")
	if err := flags.Parse(args); err != nil || flags.NArg() != 0 {
		_, _ = fmt.Fprint(stderr, serveUsage)
		return 2
	}
	minimumWebVersion, err := api.ParseVersion(*minimum)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "casa serve: -minimum-web-version: %v\n%s", err, serveUsage)
		return 2
	}
	listener, err := new(net.ListenConfig).Listen(ctx, "tcp", *listen)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	_, _ = fmt.Fprintf(stdout, "casa serve: listening on %s\n", listener.Addr())
	server := &http.Server{
		Handler: api.NewMux(api.NewServer(minimumWebVersion)),
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
	if err := <-served; !errors.Is(err, http.ErrServerClosed) {
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	return 0
}
