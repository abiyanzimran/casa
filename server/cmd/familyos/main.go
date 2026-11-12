// Command familyos is the Casa server binary (09.4 §6.2; ADR-030 item 9 keeps
// "Family OS" as the existing identifier). Subcommands are added by later
// FS-1 tasks: migrate (1.2) and serve (1.3).
package main

import (
	"fmt"
	"io"
	"os"
)

const usage = "usage: familyos <command>\n"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run executes the subcommand named by args[0] and returns the process exit
// code: 2 for a usage error, as is conventional for command-line tools.
// stdout stays in the signature for the subcommands that write results;
// none does yet.
//
// Write errors on stderr are discarded on purpose: stderr is the last place
// to report a problem, and the exit code still carries the failure.
func run(args []string, _, stderr io.Writer) int {
	if len(args) == 0 {
		_, _ = fmt.Fprint(stderr, usage)
		return 2
	}
	_, _ = fmt.Fprintf(stderr, "familyos: unknown command %q\n%s", args[0], usage)
	return 2
}
