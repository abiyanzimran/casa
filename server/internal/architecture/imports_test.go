// Package architecture holds only tests: the module's package dependency rules
// (ADR-014 item 4).
package architecture

import (
	"os/exec"
	"slices"
	"strings"
	"testing"
)

const modulePath = "github.com/abiyanzimran/casa/server/"

// allowedImports lists, per package of this module, the other packages of this
// module its production code may import (ADR-014 item 4; 09.2 §1 "Aturan
// dependensi"). Every package must be listed. Test files are not checked, so a
// test may import databasetest or another module to set up data. A row not
// backed by 09.2 §1 (its diagram or "Aturan dependensi") needs 09.2 §1 updated
// by the architect first.
var allowedImports = map[string][]string{
	"cmd/casa":                           {"internal/api", "internal/database", "internal/gen/casa/v1", "internal/household", "internal/identity", "internal/schema", "internal/webbundle"},
	"cmd/idcheck":                        {},
	"cmd/webpreview":                     {"internal/api", "internal/webbundle"},
	"internal/api":                       {"internal/command", "internal/database", "internal/gen/casa/v1", "internal/gen/casa/v1/casav1connect", "internal/household", "internal/identity"},
	"internal/command":                   {"internal/command/commanddb", "internal/database"},
	"internal/command/commanddb":         {},
	"internal/architecture":              {},
	"internal/database":                  {},
	"internal/database/databasetest":     {"internal/schema"},
	"internal/gen/casa/v1":               {},
	"internal/gen/casa/v1/casav1connect": {"internal/gen/casa/v1"},
	"internal/household":                 {"internal/database", "internal/household/householddb"},
	"internal/household/householddb":     {},
	"internal/identity":                  {"internal/database", "internal/identity/identitydb"},
	"internal/identity/identitydb":       {},
	"internal/jobs":                      {"internal/database", "internal/jobs/jobsdb"},
	"internal/jobs/jobsdb":               {},
	"internal/schema":                    {},
	// Only cmd/casa and cmd/webpreview embed the bundle; internal/api takes it
	// as an fs.FS (FS-1 1.5 plan, Task 1).
	"internal/webbundle": {},
}

func TestModuleImportsFollowTheAllowList(t *testing.T) {
	cmd := exec.Command("go", "list", "-f", "{{.ImportPath}}{{range .Imports}} {{.}}{{end}}", "./...")
	cmd.Dir = "../.." // the module root
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("go list: %v", err)
	}
	listed := map[string]bool{}
	for line := range strings.Lines(strings.TrimSpace(string(out))) {
		fields := strings.Fields(line)
		pkg := strings.TrimPrefix(fields[0], modulePath)
		listed[pkg] = true
		allowed, ok := allowedImports[pkg]
		if !ok {
			t.Errorf("%s is not declared in allowedImports (ADR-014 item 4)", pkg)
			continue
		}
		for _, imported := range fields[1:] {
			dependency, ours := strings.CutPrefix(imported, modulePath)
			if ours && !slices.Contains(allowed, dependency) {
				t.Errorf("%s imports %s, which allowedImports does not allow (ADR-014 item 4; 09.2 §1)", pkg, dependency)
			}
		}
	}
	for pkg := range allowedImports {
		if !listed[pkg] {
			t.Errorf("%s is declared but does not exist", pkg)
		}
	}
}
