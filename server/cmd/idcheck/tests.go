package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// TestRef is one test whose name starts with an ID. Skipped tests are
// reported but are never evidence (brief review focus 4). For a malformed
// name, ID is its first word and Reason says what is wrong.
type TestRef struct {
	ID, File string
	Line     int
	Skipped  bool
	Reason   string
}

// lookAlikeReason is the finding for an ID where none is read: a template
// literal or describe title, or a look-alike in another letter case
// (code-reviewer, 1.1c fix round 1).
const lookAlikeReason = "IDs belong at the start of a plain-string test/it title"

// testRoots are the directories scanned, each with the file suffix of the
// test files in it (14 §4.1 item 3: one name shape per language).
var testRoots = []struct{ dir, suffix string }{
	{"server", "_test.go"},
	{"web/src", ".test.ts"},
	{"web/tests", ".spec.ts"},
}

// skippedDirs hold third-party code, generated code, or test fixtures at any
// depth: never this repository's tests (brief review focus 5). Build output
// (web/build, web/.svelte-kit) lies outside the scanned roots, so a source
// directory named build is still scanned.
var skippedDirs = map[string]bool{"node_modules": true, "gen": true, "testdata": true}

var (
	// testNameID matches a test name that starts with one ID followed by a
	// space, a colon, or the end of the name.
	testNameID = regexp.MustCompile(`^(` + strings.Join([]string{
		frPattern + `/AC-\d+`, approvalInvariantPattern, invariantPattern, guardrailPattern,
		workedExamplePattern, scenarioPattern, securityPattern,
	}, "|") + `)(?:[ :]|$)`)
	// idLikePrefix matches a name that starts like an ID; one that does not
	// also match testNameID is a typo, not a plain name (brief review focus 3).
	// The hyphen is part of every prefix, so "INVALID input" is a plain name
	// (lead ruling 13 Nov 2026, 1.1c).
	idLikePrefix = regexp.MustCompile(`^(?:FR-|INV-|PG-|CTH-|SCN-|SEC-)`)
	// idLikePrefixAnyCase also catches look-alikes such as "inv-01".
	idLikePrefixAnyCase = regexp.MustCompile(`(?i)^(?:FR-|INV-|PG-|CTH-|SCN-|SEC-)`)
)

// scanTests reads the names of the tests under repoDir and returns those
// that start with an ID, and those that start like an ID but are not one.
func scanTests(repoDir string) (refs, malformed []TestRef, err error) {
	for _, root := range testRoots {
		dir := filepath.Join(repoDir, filepath.FromSlash(root.dir))
		if _, err := os.Stat(dir); err != nil {
			return nil, nil, fmt.Errorf("test root %s is missing", dir)
		}
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				if skippedDirs[d.Name()] {
					return filepath.SkipDir
				}
				return nil
			}
			if !strings.HasSuffix(path, root.suffix) {
				return nil
			}
			rel, err := filepath.Rel(repoDir, path)
			if err != nil {
				return err
			}
			names, err := readTestNames(path)
			if err != nil {
				return err
			}
			for _, n := range names {
				ref := TestRef{File: filepath.ToSlash(rel), Line: n.line, Skipped: n.skipped}
				id, isID := idOfTestName(n.name)
				switch {
				case n.plain && isID:
					ref.ID = id
					refs = append(refs, ref)
				case n.plain && idLikePrefix.MatchString(n.name):
					ref.ID = firstWord(n.name)
					ref.Reason = fmt.Sprintf("test name starts like an ID but %s is not one", ref.ID)
					malformed = append(malformed, ref)
				case idLikePrefixAnyCase.MatchString(n.name):
					ref.ID, ref.Skipped = firstWord(n.name), false
					ref.Reason = lookAlikeReason
					malformed = append(malformed, ref)
				}
			}
			return nil
		})
		if err != nil {
			return nil, nil, err
		}
	}
	return refs, malformed, nil
}

func idOfTestName(name string) (string, bool) {
	m := testNameID.FindStringSubmatch(name)
	if m == nil {
		return "", false
	}
	return m[1], true
}

func firstWord(name string) string {
	word, _, _ := strings.Cut(name, " ")
	return strings.TrimSuffix(word, ":")
}

// readTestNames returns every literal test name in one file with its line.
func readTestNames(path string) ([]testName, error) {
	if strings.HasSuffix(path, ".go") {
		return readGoTestNames(path)
	}
	content, err := os.ReadFile(path) //nolint:gosec // G304: the path is a test file found under the -repo flag of a CI tool, not user input
	if err != nil {
		return nil, err
	}
	names, err := readTypeScriptTestNames(string(content))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return names, nil
}

// readGoTestNames returns the first argument of every x.Run call that is a
// string literal; a name held in a variable is not read.
func readGoTestNames(path string) ([]testName, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	var names []testName
	var unquoteErr error
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok || len(call.Args) == 0 {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || sel.Sel.Name != "Run" {
			return true
		}
		if _, ok := sel.X.(*ast.Ident); !ok {
			return true
		}
		lit, ok := call.Args[0].(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return true
		}
		name, err := strconv.Unquote(lit.Value)
		if err != nil {
			unquoteErr = fmt.Errorf("%s: %w", fset.Position(lit.Pos()), err)
			return false
		}
		names = append(names, testName{
			name:    name,
			line:    fset.Position(lit.Pos()).Line,
			skipped: len(call.Args) > 1 && skipsDirectly(call.Args[1]),
			plain:   true,
		})
		return true
	})
	return names, unquoteErr
}

// skipsDirectly reports whether fn is a func literal whose body calls Skip,
// Skipf or SkipNow on its first parameter as a statement of its own (lead
// ruling 13 Nov 2026, 1.1c). Not seen, so such a test counts as evidence: a
// skip inside a branch or a deferred call, a skip in a parent test, and a
// subtest given as a named function; code-reviewer checks those by hand
// (14 §4.13).
func skipsDirectly(fn ast.Expr) bool {
	lit, ok := fn.(*ast.FuncLit)
	if !ok || len(lit.Type.Params.List) == 0 || len(lit.Type.Params.List[0].Names) == 0 {
		return false
	}
	param := lit.Type.Params.List[0].Names[0].Name
	for _, stmt := range lit.Body.List {
		expr, ok := stmt.(*ast.ExprStmt)
		if !ok {
			continue
		}
		call, ok := expr.X.(*ast.CallExpr)
		if !ok {
			continue
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		if x, ok := sel.X.(*ast.Ident); ok && x.Name == param {
			switch sel.Sel.Name {
			case "Skip", "Skipf", "SkipNow":
				return true
			}
		}
	}
	return false
}
