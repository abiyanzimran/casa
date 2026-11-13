// Command casa is the Casa server binary (09.4 §6.2; owner decision 71,
// docs/17). Subcommands: migrate, serve and setup-token.
package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	_ "time/tzdata" // the IANA database for household time zones, whatever the host has (06 §2.4)
)

const usage = "usage: casa <command>\n"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	code := run(ctx, os.Args[1:], os.Stdout, os.Stderr)
	stop()
	os.Exit(code)
}

// run executes the subcommand named by args[0] and returns the process exit
// code: 0 on success, 1 on failure, 2 for a usage error, as is conventional
// for command-line tools.
//
// Write errors on stdout and stderr are discarded on purpose: they are the
// last place to report a problem, and the exit code still carries the result.
func run(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		_, _ = fmt.Fprint(stderr, usage)
		return 2
	}
	switch args[0] {
	case "migrate":
		return migrate(ctx, args[1:], stdout, stderr)
	case "serve":
		return serve(ctx, args[1:], stdout, stderr)
	case "setup-token":
		return setupToken(ctx, args[1:], stdout, stderr)
	}
	_, _ = fmt.Fprintf(stderr, "casa: unknown command %q\n%s", args[0], usage)
	return 2
}
