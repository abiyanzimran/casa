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
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/api"
	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/identity"
	"github.com/abiyanzimran/casa/server/internal/jobs"
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
	// The app's origin: the base of emailed links (09.4 §2) and the only
	// Origin a browser's request may carry (SEC-162 (a)).
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
	if err := checkOriginFlag(*origin); err != nil {
		_, _ = fmt.Fprintf(stderr, "casa serve: -origin: %v\n%s", err, serveUsage)
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
	pool, err := openPinged(ctx, appDatabaseURL, "the database")
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	operatorName, err := readOperatorName()
	if err != nil {
		pool.Close()
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	// The jobs runner claims as job_runner (09.2 §3 row "job
	// (pengambilan)"; plan 1.7 decision 4).
	jobPool, err := openPinged(ctx, jobDatabaseURL, "the job database")
	if err != nil {
		pool.Close()
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	}
	logger := slog.New(slog.NewJSONHandler(stderr, nil)) // JSON to journald (09.4 §3)
	identityService := identity.New(pool, identity.Config{
		Origin: *origin, SendMail: writeMail(*mailDir), Now: time.Now, Password: identity.DefaultPasswordParams,
		Logger: logger,
	})
	// closeAll ends the background work, then the pools. It runs only when
	// no handler or job can still use them: before serving, or after the
	// runner stopped and a clean Shutdown. After a Shutdown that timed out,
	// handlers may still run, so the pools are left to the process exit
	// (code-reviewer S4 on #32).
	closeAll := func() {
		identityService.Close()
		jobPool.Close()
		pool.Close()
	}
	// One clock for the claim and the expiry's compare-and-set (plan 1.7
	// decision 1; 09.3 §9 note of 14 November 2026 (2)).
	now := time.Now
	householdService := household.New(pool, now, operatorName, household.PublishedNotices)
	runner := jobs.NewRunner(jobs.Config{
		Claim: jobPool, App: pool, Now: now, Logger: logger,
		Handlers: map[string]jobs.Handler{household.ExpireInvitationKind: householdService.ExpireInvitationJob},
	})
	apiServer := api.NewServer(minimumWebVersion, *origin, api.Sessions{
		Bearer: identityService.Authenticate, Cookie: identityService.AuthenticateCookie,
		UserAction: identityService.TouchUserAction,
	}, api.NoticeHeld(pool, householdService))
	api.RegisterIdentityService(apiServer, pool, identityService, api.InvitationOpen(householdService))
	api.RegisterHouseholdService(apiServer, pool, householdService, identityService)
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
	// The runner stops with ctx, after the job it runs (09.3 §9). Run
	// returns a database failure without retrying (plan 1.7 Task 2): serve
	// then stops and exits 1, and its supervisor restarts it, so a runner
	// that cannot reach the database is never left stopped while the API
	// still serves (10.1 §10 note 1.7 Q3 condition 3).
	runnerCtx, stopRunner := context.WithCancel(ctx)
	defer stopRunner()
	runnerDone := make(chan struct{})
	var runnerErr error
	go func() {
		defer close(runnerDone)
		runnerErr = runner.Run(runnerCtx)
	}()
	select {
	case err := <-served:
		stopRunner()
		<-runnerDone
		_, _ = fmt.Fprintf(stderr, "casa serve: %v\n", err)
		return 1
	case <-runnerDone:
	case <-ctx.Done():
	}
	stopRunner()
	<-runnerDone
	code := 0
	if runnerErr != nil {
		// The error's text is not printed: it may quote PostgreSQL's
		// (SEC-44; ES §4).
		_, _ = fmt.Fprintln(stderr, "casa serve: the jobs runner stopped on a database error")
		code = 1
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
	return code
}

// jobDatabaseURL names the LoadCredential= file holding the connection
// string of role job_runner (plan 1.7 decision 4).
const jobDatabaseURL = "job-database-url"

// openPinged opens a pool with the connection string in credential and
// pings it, so serve fails at start, not at the first request. pgx errors
// name the host and user at most, never the password (ES §4).
func openPinged(ctx context.Context, credential, what string) (*pgxpool.Pool, error) {
	pool, err := openPool(ctx, credential, what)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("open %s: %w", what, err)
	}
	return pool, nil
}

// operatorNameCredential holds the name that fills {pengelola} in the
// operator notice: configuration, so no person's name enters the repository
// (09.3 §1 "Mekanisme"; SEC-34).
const operatorNameCredential = "operator-name"

// readOperatorName reads operatorNameCredential; an empty name is refused,
// and the name itself is never printed (ES §4).
func readOperatorName() (string, error) {
	name, err := readCredential(operatorNameCredential)
	if err == nil && strings.TrimSpace(name) == "" {
		err = fmt.Errorf("credential %s is empty", operatorNameCredential)
	}
	return name, err
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

// checkOriginFlag accepts only an origin as a browser serializes it in the
// Origin header, since the origin check compares the two as strings
// (SEC-162 (a)): scheme://host[:port] in lower-case ASCII, without path,
// query, fragment, user info, a trailing dot, an empty or default port, or a
// port outside 1 to 65535 or with a leading zero; https, or http on a loopback host for
// development. Anything else would fail every web request closed without
// saying why. The error never repeats the value, which might hold user info.
func checkOriginFlag(origin string) error {
	u, err := url.Parse(origin)
	if err != nil || u.Host == "" || u.Scheme+"://"+u.Host != origin || strings.ToLower(origin) != origin ||
		strings.HasSuffix(u.Host, ":") || strings.HasSuffix(u.Hostname(), ".") ||
		strings.ContainsFunc(u.Host, func(r rune) bool { return r > unicode.MaxASCII }) {
		return errors.New("want scheme://host[:port] in lower-case ASCII, without path, query, fragment, user info or trailing dot")
	}
	if port := u.Port(); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil || strconv.Itoa(n) != port || n < 1 || n > 65535 || u.Scheme == "https" && n == 443 || u.Scheme == "http" && n == 80 {
			return errors.New("want a port from 1 to 65535 other than the scheme's default, which browsers leave out")
		}
	}
	switch u.Hostname() {
	case "localhost", "127.0.0.1", "::1":
		if u.Scheme == "http" {
			return nil
		}
	}
	if u.Scheme != "https" {
		return errors.New("want https, or http on localhost, 127.0.0.1 or ::1")
	}
	return nil
}
