package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"testing"
)

// writeTree writes files (slash paths relative to the root) under a new
// temporary directory and returns it. The fixture repository is built here
// rather than in testdata because the repository ignores node_modules/.
func writeTree(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for path, content := range files {
		writeFile(t, filepath.Join(root, filepath.FromSlash(path)), content)
	}
	return root
}

const goTestFixture = `package x

import "testing"

func TestX(t *testing.T) {
	t.Run("FR-FAM-001/AC-1 creator gets a household", func(t *testing.T) {})
	t.Run("INV-A35: exit is never refused", func(t *testing.T) {})
	t.Run("CTH-03b", func(t *testing.T) {})
	name := "INV-02 not a literal"
	t.Run(name, func(t *testing.T) {})
	t.Run("FR-FAM-01/AC-1 typo", func(t *testing.T) {})
	t.Run("plain name", func(t *testing.T) {})
	t.Run("INVALID input is a plain name", func(t *testing.T) {})
	t.Run("PG-02 skipped", func(t *testing.T) { t.Skip("later") })
	t.Run("PG-03 skipped", func(t *testing.T) { t.Skipf("later %d", 1) })
	t.Run("PG-04 skipped", func(t *testing.T) { t.SkipNow() })
	t.Run("PG-05 skips only sometimes", func(t *testing.T) {
		if testing.Short() {
			t.Skip("short")
		}
	})
	t.Run("inv-01 lower case", func(t *testing.T) {})
}
`

// Each line is one case; the line numbers are asserted below.
const tsTestFixture = `import { describe, expect, test, it } from 'vitest';
// comment line 2
test('INV-01 integer amounts', () => {});
it("SEC-158: closed registration", () => {});
test.skip('PG-01 symmetric', () => {});
test.only('SCN-01 meal money', () => {});
test('INV-1 typo', () => {});
expect(/x/.test('INV-03 not a test')).toBe(false);
test('INVALID input', () => {});
// test('INV-04 commented out', () => {});
/* test('INV-05 in a block comment', () => {}); */
const s = "test('INV-06 inside a string')";
test(` + "`INV-07 template title`" + `, () => {});
describe('INV-08 describe title', () => {});
test.describe('SEC-01 describe title', () => {});
test.todo('INV-09 later');
test.fixme('INV-10 broken', () => {});
test('inv-01 lower case', () => {});
test('Plain title, test(\'INV-11\') inside it', () => {});
const re = /'/; test('SEC-02 after a regex with a quote', () => {});
const tpl = ` + "`${s ? 'a' : `b${'c'}`}`" + `; test('SEC-03 after a nested template', () => {});
it.todo('inv-02 lower case later');
`

func TestScanTestsReadsIDsFromTestNames(t *testing.T) {
	repo := writeTree(t, map[string]string{
		"server/x/x_test.go":                  goTestFixture,
		"server/x/testdata/y_test.go":         goTestFixture,
		"server/gen/z_test.go":                goTestFixture,
		"web/src/lib/money.test.ts":           tsTestFixture,
		"web/src/lib/money.ts":                tsTestFixture,
		"web/src/lib/build/d.test.ts":         "test('INV-12 in a source directory named build', () => {});\n",
		"web/src/gen/e.test.ts":               tsTestFixture,
		"web/tests/e2e/shell.spec.ts":         "test('SEC-175: title', async () => {});\n",
		"web/tests/e2e/skipped.spec.ts":       "test.describe.skip('group', () => {\n\ttest('PG-06 in a skipped group', async () => {});\n});\n",
		"web/tests/e2e/fixme.spec.ts":         "describe.fixme('group', () => {\n\tit('PG-07 in a fixme group', () => {});\n});\n",
		"web/node_modules/pkg/a.test.ts":      tsTestFixture,
		"web/.svelte-kit/output/b.test.ts":    tsTestFixture,
		"web/build/c.spec.ts":                 tsTestFixture,
		"web/tests/node_modules/d.spec.ts":    tsTestFixture,
		"docs/notes_test.go":                  goTestFixture,
		"web/src/lib/not-a-test.spec.ts.orig": tsTestFixture,
	})
	refs, malformed, err := scanTests(repo)
	if err != nil {
		t.Fatal(err)
	}
	const money = "web/src/lib/money.test.ts"
	wantRefs := []TestRef{
		{ID: "FR-FAM-001/AC-1", File: "server/x/x_test.go", Line: 6},
		{ID: "INV-A35", File: "server/x/x_test.go", Line: 7},
		{ID: "CTH-03b", File: "server/x/x_test.go", Line: 8},
		{ID: "PG-02", File: "server/x/x_test.go", Line: 14, Skipped: true},
		{ID: "PG-03", File: "server/x/x_test.go", Line: 15, Skipped: true},
		{ID: "PG-04", File: "server/x/x_test.go", Line: 16, Skipped: true},
		{ID: "PG-05", File: "server/x/x_test.go", Line: 17},
		{ID: "INV-12", File: "web/src/lib/build/d.test.ts", Line: 1},
		{ID: "INV-01", File: money, Line: 3},
		{ID: "SEC-158", File: money, Line: 4},
		{ID: "PG-01", File: money, Line: 5, Skipped: true},
		{ID: "SCN-01", File: money, Line: 6},
		{ID: "INV-09", File: money, Line: 16, Skipped: true},
		{ID: "INV-10", File: money, Line: 17, Skipped: true},
		{ID: "SEC-02", File: money, Line: 20},
		{ID: "SEC-03", File: money, Line: 21},
		{ID: "PG-07", File: "web/tests/e2e/fixme.spec.ts", Line: 2, Skipped: true},
		{ID: "SEC-175", File: "web/tests/e2e/shell.spec.ts", Line: 1},
		{ID: "PG-06", File: "web/tests/e2e/skipped.spec.ts", Line: 2, Skipped: true},
	}
	if !slices.Equal(refs, wantRefs) {
		t.Fatalf("refs =\n%v\nwant\n%v", refs, wantRefs)
	}
	const notAnID = "test name starts like an ID but %s is not one"
	wantMalformed := []TestRef{
		{ID: "FR-FAM-01/AC-1", File: "server/x/x_test.go", Line: 11, Reason: fmt.Sprintf(notAnID, "FR-FAM-01/AC-1")},
		{ID: "inv-01", File: "server/x/x_test.go", Line: 22, Reason: lookAlikeReason},
		{ID: "INV-1", File: money, Line: 7, Reason: fmt.Sprintf(notAnID, "INV-1")},
		{ID: "INV-07", File: money, Line: 13, Reason: lookAlikeReason},
		{ID: "INV-08", File: money, Line: 14, Reason: lookAlikeReason},
		{ID: "SEC-01", File: money, Line: 15, Reason: lookAlikeReason},
		{ID: "inv-01", File: money, Line: 18, Reason: lookAlikeReason},
		{ID: "inv-02", File: money, Line: 22, Reason: lookAlikeReason},
	}
	if !slices.Equal(malformed, wantMalformed) {
		t.Fatalf("malformed =\n%v\nwant\n%v", malformed, wantMalformed)
	}
}

func TestScanTestsFailsWhenATestRootIsMissing(t *testing.T) {
	repo := writeTree(t, map[string]string{"server/go.mod": "module x\n"})
	_, _, err := scanTests(repo)
	want := "test root " + filepath.Join(repo, "web", "src") + " is missing"
	if got := errString(err); got != want {
		t.Fatalf("error = %q, want %q", got, want)
	}
}
