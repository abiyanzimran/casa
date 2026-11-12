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
// reported but are never evidence (brief review focus 4).
type TestRef struct {
	ID, File string
	Line     int
	Skipped  bool
}

// testRoots are the directories scanned, each with the file suffix of the
// test files in it (14 §4.1 item 3: one name shape per language).
var testRoots = []struct{ dir, suffix string }{
	{"server", "_test.go"},
	{"web/src", ".test.ts"},
	{"web/tests", ".spec.ts"},
}

// skippedDirs hold third-party code, build output, generated code, or test
// fixtures: never this repository's tests (brief review focus 5).
var skippedDirs = map[string]bool{
	"node_modules": true, ".svelte-kit": true, "build": true, "gen": true, "testdata": true,
}

var (
	// testNameID matches a test name that starts with one ID followed by a
	// space, a colon, or the end of the name.
	testNameID = regexp.MustCompile(`^(` + strings.Join([]string{
		frPattern + `/AC-\d+`, approvalInvariantPattern, invariantPattern, guardrailPattern,
		workedExamplePattern, scenarioPattern, securityPattern,
	}, "|") + `)(?:[ :]|$)`)
	// idLikePrefix matches a name that starts like an ID; one that does not
	// also match testNameID is a typo, not a plain name (brief review focus 3).
	idLikePrefix = regexp.MustCompile(`^(?:FR-|INV|PG-|CTH-|SCN-|SEC-)`)
	// typeScriptTest matches test( and it( with a string literal title, with
	// .only or .skip; not a method such as /re/.test(…).
	typeScriptTest = regexp.MustCompile(`(?:^|[^.\w$])(test|it)(\.only|\.skip)?\(\s*(?:'([^'\n]*)'|"([^"\n]*)"|` + "`([^`$]*)`" + `)`)
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
			names, err := readTestNames(path, filepath.ToSlash(rel))
			if err != nil {
				return err
			}
			for _, n := range names {
				switch id, ok := idOfTestName(n.ID); {
				case ok:
					n.ID = id
					refs = append(refs, n)
				case idLikePrefix.MatchString(n.ID):
					n.ID, _, _ = strings.Cut(n.ID, " ")
					n.ID = strings.TrimSuffix(n.ID, ":")
					malformed = append(malformed, n)
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

// readTestNames returns every literal test name in one file, the whole name
// in ID, with its line.
func readTestNames(path, rel string) ([]TestRef, error) {
	if strings.HasSuffix(path, ".go") {
		return readGoTestNames(path, rel)
	}
	return readTypeScriptTestNames(path, rel)
}

// readGoTestNames returns the first argument of every x.Run call that is a
// string literal; a name held in a variable is not read.
func readGoTestNames(path, rel string) ([]TestRef, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	var names []TestRef
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
		names = append(names, TestRef{ID: name, File: rel, Line: fset.Position(lit.Pos()).Line})
		return true
	})
	return names, unquoteErr
}

func readTypeScriptTestNames(path, rel string) ([]TestRef, error) {
	content, err := os.ReadFile(path) //nolint:gosec // G304: the path is a test file found under the -repo flag of a CI tool, not user input
	if err != nil {
		return nil, err
	}
	text := string(content)
	var names []TestRef
	for _, m := range typeScriptTest.FindAllStringSubmatchIndex(text, -1) {
		// Groups: 1 test or it, 2 .only or .skip, 3 to 5 the title in ', " or `.
		var name string
		for g := 3; g <= 5; g++ {
			if m[2*g] >= 0 {
				name = text[m[2*g]:m[2*g+1]]
			}
		}
		names = append(names, TestRef{
			ID:      name,
			File:    rel,
			Line:    strings.Count(text[:m[2]], "\n") + 1,
			Skipped: m[4] >= 0 && text[m[4]:m[5]] == ".skip",
		})
	}
	return names, nil
}
