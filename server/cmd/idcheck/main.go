// Command idcheck is the ID traceability checker (QG-06; 14 §4.13). It reads
// the IDs that tests must prove from the planning documents and the test
// names in the repository, and fails when a test names an ID that does not
// exist, when a name starts like an ID but is not one, when an ID in the
// scope of a finished step has no test, and, with -release, when any
// required ID has no test (QG-20, QG-30).
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"strings"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run checks the repository and returns the process exit code: 0 when
// nothing is found, 1 for findings or unreadable input, 2 for a usage error.
// Findings go to stdout, one per line, with the ID, file, and line.
//
// Write errors on stdout and stderr are discarded on purpose: the exit code
// still carries the result.
func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("idcheck", flag.ContinueOnError)
	flags.SetOutput(stderr)
	docsDir := flags.String("docs", "", "planning documents directory (docs/)")
	repoDir := flags.String("repo", "", "repository root whose tests are scanned")
	scopePath := flags.String("scope", "", "IDs of finished steps that must have a test (traceability/scope.txt)")
	release := flags.Bool("release", false, "require a test for every required ID (QG-20, QG-30)")
	listUntuk := flags.Bool("list-untuk", false, `print each value of the SEC "Untuk" column with its mapping, and exit`)
	counts := flags.Bool("counts", false, "print the known and required IDs per kind, and exit")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if *docsDir == "" {
		_, _ = fmt.Fprintln(stderr, "idcheck: -docs is required")
		return 2
	}

	var err error
	switch {
	case *listUntuk:
		err = printUntukValues(*docsDir, stdout)
	case *counts:
		err = printCounts(*docsDir, stdout)
	case *repoDir == "":
		_, _ = fmt.Fprintln(stderr, "idcheck: -repo is required")
		return 2
	default:
		var findings int
		findings, err = check(*docsDir, *repoDir, *scopePath, *release, stdout)
		if err == nil && findings > 0 {
			return 1
		}
	}
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "idcheck: %v\n", err)
		return 1
	}
	return 0
}

// check prints every finding and warning and returns the number of findings.
func check(docsDir, repoDir, scopePath string, release bool, out io.Writer) (int, error) {
	reqs, err := readRequirements(docsDir, true)
	if err != nil {
		return 0, err
	}
	refs, malformed, err := scanTests(repoDir)
	if err != nil {
		return 0, err
	}
	required := map[string]bool{}
	for _, r := range reqs {
		required[r.ID] = r.Required
	}

	findings := 0
	report := func(format string, a ...any) {
		findings++
		_, _ = fmt.Fprintf(out, format+"\n", a...)
	}
	tested := map[string]bool{}
	for _, ref := range refs {
		switch _, known := required[ref.ID]; {
		case !known:
			report("%s:%d: test names unknown ID %s", ref.File, ref.Line, ref.ID)
		case ref.Skipped:
			_, _ = fmt.Fprintf(out, "%s:%d: warning: skipped test is not evidence for %s\n", ref.File, ref.Line, ref.ID)
		default:
			tested[ref.ID] = true
		}
	}
	for _, ref := range malformed {
		report("%s:%d: test name starts like an ID but %s is not one", ref.File, ref.Line, ref.ID)
	}

	if scopePath != "" {
		scope, err := readScope(scopePath)
		if err != nil {
			return 0, err
		}
		for _, entry := range scope {
			switch {
			case !required[entry.id]:
				report("%s:%d: %s is not a required ID", scopePath, entry.line, entry.id)
			case !tested[entry.id]:
				report("%s:%d: %s has no test", scopePath, entry.line, entry.id)
			}
		}
	}
	if release {
		for _, r := range reqs {
			if r.Required && !tested[r.ID] {
				report("release: %s has no test", r.ID)
			}
		}
	}
	return findings, nil
}

type scopeEntry struct {
	id   string
	line int
}

// readScope reads one ID per line; "#" starts a comment (14 §4.13 item 2).
func readScope(path string) ([]scopeEntry, error) {
	file, err := os.Open(path) //nolint:gosec // G304: the path is the scope list under the -scope flag of a CI tool, not user input
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }() // read-only: a close error loses nothing
	var entries []scopeEntry
	lines := bufio.NewScanner(file)
	for n := 1; lines.Scan(); n++ {
		text, _, _ := strings.Cut(lines.Text(), "#")
		if id := strings.TrimSpace(text); id != "" {
			entries = append(entries, scopeEntry{id: id, line: n})
		}
	}
	return entries, lines.Err()
}

func printUntukValues(docsDir string, out io.Writer) error {
	rows, err := readSecurityRows(docsDir)
	if err != nil {
		return err
	}
	rowsPerValue := map[string]int{}
	for _, row := range rows {
		for _, value := range untukValues(cell(row.tableRow, 2)) {
			rowsPerValue[value]++
		}
	}
	for _, value := range slices.Sorted(maps.Keys(rowsPerValue)) {
		mapping := "unmapped"
		if applies, ok := untukAppliesToServerOrWeb[value]; ok {
			mapping = map[bool]string{true: "server or web", false: "not server or web"}[applies]
		}
		_, _ = fmt.Fprintf(out, "%s\t%s\t%d rows\n", value, mapping, rowsPerValue[value])
	}
	return nil
}

func printCounts(docsDir string, out io.Writer) error {
	reqs, err := readRequirements(docsDir, false)
	if err != nil {
		return err
	}
	counts := countByKind(reqs)
	_, _ = fmt.Fprintln(out, "kind\tknown\trequired")
	for _, kind := range kinds {
		_, _ = fmt.Fprintf(out, "%s\t%d\t%d\n", kind, counts[kind].known, counts[kind].required)
	}
	return nil
}
