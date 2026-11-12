package main

import (
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
}
`

const tsTestFixture = `import { expect, test, it } from 'vitest';

test('INV-01 integer amounts', () => {});
it("SEC-158: closed registration", () => {});
test.skip('PG-01 symmetric', () => {});
test.only('SCN-01 meal money', () => {});
test('INV-1 typo', () => {});
expect(/x/.test('INV-03 not a test')).toBe(false);
`

func TestScanTestsReadsIDsFromTestNames(t *testing.T) {
	repo := writeTree(t, map[string]string{
		"server/x/x_test.go":                  goTestFixture,
		"server/x/testdata/y_test.go":         goTestFixture,
		"server/gen/z_test.go":                goTestFixture,
		"web/src/lib/money.test.ts":           tsTestFixture,
		"web/src/lib/money.ts":                tsTestFixture,
		"web/tests/e2e/shell.spec.ts":         "test('SEC-175: title', async () => {});\n",
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
	wantRefs := []TestRef{
		{ID: "FR-FAM-001/AC-1", File: "server/x/x_test.go", Line: 6},
		{ID: "INV-A35", File: "server/x/x_test.go", Line: 7},
		{ID: "CTH-03b", File: "server/x/x_test.go", Line: 8},
		{ID: "INV-01", File: "web/src/lib/money.test.ts", Line: 3},
		{ID: "SEC-158", File: "web/src/lib/money.test.ts", Line: 4},
		{ID: "PG-01", File: "web/src/lib/money.test.ts", Line: 5, Skipped: true},
		{ID: "SCN-01", File: "web/src/lib/money.test.ts", Line: 6},
		{ID: "SEC-175", File: "web/tests/e2e/shell.spec.ts", Line: 1},
	}
	if !slices.Equal(refs, wantRefs) {
		t.Fatalf("refs =\n%v\nwant\n%v", refs, wantRefs)
	}
	wantMalformed := []TestRef{
		{ID: "FR-FAM-01/AC-1", File: "server/x/x_test.go", Line: 11},
		{ID: "INV-1", File: "web/src/lib/money.test.ts", Line: 7},
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
