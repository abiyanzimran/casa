package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

const realDocs = "../../../docs"

// fixtureRepo returns a repository whose only Go test is goTest and whose
// only web test is webTest.
func fixtureRepo(t *testing.T, goTest, webTest string) string {
	t.Helper()
	return writeTree(t, map[string]string{
		"server/x/x_test.go":  "package x\n\nimport \"testing\"\n\nfunc TestX(t *testing.T) {\n" + goTest + "}\n",
		"web/src/a.test.ts":   webTest,
		"web/tests/e2e/.keep": "",
	})
}

func runIDCheck(t *testing.T, args ...string) (code int, stdout, stderr string) {
	t.Helper()
	var out, errOut bytes.Buffer
	code = run(args, &out, &errOut)
	return code, out.String(), errOut.String()
}

func TestRunPassesWhenEveryTestNamesAKnownID(t *testing.T) {
	repo := fixtureRepo(t, "\tt.Run(\"INV-01 integer amounts\", func(t *testing.T) {})\n", "test('SEC-175: title', () => {});\n")
	code, stdout, stderr := runIDCheck(t, "-docs", realDocs, "-repo", repo)
	if code != 0 || stdout != "" || stderr != "" {
		t.Fatalf("code %d, stdout %q, stderr %q; want 0 and no output", code, stdout, stderr)
	}
}

func TestRunFailsOnATestWithAnUnknownID(t *testing.T) {
	repo := fixtureRepo(t, "\tt.Run(\"INV-99 x\", func(t *testing.T) {})\n", "")
	code, stdout, _ := runIDCheck(t, "-docs", realDocs, "-repo", repo)
	want := "server/x/x_test.go:6: test names unknown ID INV-99\n"
	if code != 1 || stdout != want {
		t.Fatalf("code %d, stdout %q; want 1, %q", code, stdout, want)
	}
}

func TestRunFailsOnATestNameThatOnlyLooksLikeAnID(t *testing.T) {
	repo := fixtureRepo(t, "", "test('INV-1 typo', () => {});\n")
	code, stdout, _ := runIDCheck(t, "-docs", realDocs, "-repo", repo)
	want := "web/src/a.test.ts:1: test name starts like an ID but INV-1 is not one\n"
	if code != 1 || stdout != want {
		t.Fatalf("code %d, stdout %q; want 1, %q", code, stdout, want)
	}
}

func TestRunFailsOnAScopeIDWithoutATest(t *testing.T) {
	repo := fixtureRepo(t, "\tt.Run(\"INV-01 integer amounts\", func(t *testing.T) {})\n", "test.skip('PG-01 symmetric', () => {});\n")
	scope := filepath.Join(repo, "traceability", "scope.txt")
	writeFile(t, scope, "# step 1\nINV-01\nPG-01 # skipped only\nINV-02\n")
	code, stdout, _ := runIDCheck(t, "-docs", realDocs, "-repo", repo, "-scope", scope)
	want := "web/src/a.test.ts:1: warning: skipped test is not evidence for PG-01\n" +
		scope + ":3: PG-01 has no test\n" +
		scope + ":4: INV-02 has no test\n"
	if code != 1 || stdout != want {
		t.Fatalf("code %d, stdout %q; want 1, %q", code, stdout, want)
	}
}

func TestRunFailsOnAScopeIDThatIsNotRequired(t *testing.T) {
	repo := fixtureRepo(t, "\tt.Run(\"SCN-03 tier 3\", func(t *testing.T) {})\n", "")
	scope := filepath.Join(repo, "traceability", "scope.txt")
	writeFile(t, scope, "SCN-03\nINV-99\n")
	code, stdout, _ := runIDCheck(t, "-docs", realDocs, "-repo", repo, "-scope", scope)
	want := scope + ":1: SCN-03 is not a required ID\n" + scope + ":2: INV-99 is not a required ID\n"
	if code != 1 || stdout != want {
		t.Fatalf("code %d, stdout %q; want 1, %q", code, stdout, want)
	}
}

func TestRunPassesWhenEveryScopeIDHasATest(t *testing.T) {
	repo := fixtureRepo(t, "\tt.Run(\"INV-01 integer amounts\", func(t *testing.T) {})\n", "")
	scope := filepath.Join(repo, "traceability", "scope.txt")
	writeFile(t, scope, "INV-01\n")
	code, stdout, _ := runIDCheck(t, "-docs", realDocs, "-repo", repo, "-scope", scope)
	if code != 0 || stdout != "" {
		t.Fatalf("code %d, stdout %q; want 0 and no output", code, stdout)
	}
}

func TestRunReleaseFailsOnEveryRequiredIDWithoutATest(t *testing.T) {
	repo := fixtureRepo(t, "\tt.Run(\"INV-01 integer amounts\", func(t *testing.T) {})\n", "")
	code, stdout, _ := runIDCheck(t, "-docs", realDocs, "-repo", repo, "-release")
	reqs, err := readRequirements(realDocs, true)
	if err != nil {
		t.Fatal(err)
	}
	var want strings.Builder
	for _, r := range reqs {
		if r.Required && r.ID != "INV-01" {
			want.WriteString("release: " + r.ID + " has no test\n")
		}
	}
	if code != 1 || stdout != want.String() {
		t.Fatalf("code %d, stdout %q; want 1, %q", code, stdout, want.String())
	}
}

func TestRunListsUntukValues(t *testing.T) {
	code, stdout, _ := runIDCheck(t, "-docs", "testdata/docs", "-list-untuk")
	want := "09\tserver or web\t2 rows\n" +
		"Apple\tnot server or web\t2 rows\n" +
		"Google\tnot server or web\t3 rows\n" +
		"Semua jalur\tserver or web\t1 rows\n"
	if code != 0 || stdout != want {
		t.Fatalf("code %d, stdout %q; want 0, %q", code, stdout, want)
	}
}

func TestRunPrintsCountsPerKind(t *testing.T) {
	code, stdout, _ := runIDCheck(t, "-docs", "testdata/docs", "-counts")
	want := "kind\tknown\trequired\n" +
		"FR\t22\t13\n" +
		"INV\t3\t3\n" +
		"INV-A\t1\t1\n" +
		"PG\t2\t1\n" +
		"CTH\t3\t3\n" +
		"SCN\t5\t2\n" +
		"SEC\t6\t3\n"
	if code != 0 || stdout != want {
		t.Fatalf("code %d, stdout %q; want 0, %q", code, stdout, want)
	}
}

func TestRunWithoutDocsIsAUsageError(t *testing.T) {
	code, _, stderr := runIDCheck(t, "-repo", ".")
	want := "idcheck: -docs is required\n"
	if code != 2 || stderr != want {
		t.Fatalf("code %d, stderr %q; want 2, %q", code, stderr, want)
	}
}
