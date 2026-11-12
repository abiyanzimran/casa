package main

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

// testdata/docs holds lines copied from the real documents: one block of
// every definition shape, plus the shapes that must not count (a fenced
// block, a changelog table that repeats an ID).
func TestReadRequirementsReadsOnlyDefinitions(t *testing.T) {
	reqs, err := readRequirements("testdata/docs", false)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, r := range reqs {
		got[r.ID] = r.Required
	}
	want := map[string]bool{
		// Tier 1 FR in a module file.
		"FR-FAM-001/AC-1": true, "FR-FAM-001/AC-2": true, "FR-FAM-001/AC-3": true,
		"FR-FAM-001/AC-4": true, "FR-FAM-001/AC-5": true, "FR-FAM-001/AC-6": true,
		"FR-FAM-001/AC-7": true, "FR-FAM-001/AC-8": true, "FR-FAM-001/AC-9": true,
		"FR-FAM-001/AC-10": true,
		// Tier 2 FR in a module file: known, not required.
		"FR-FAM-025/AC-1": false, "FR-FAM-025/AC-2": false, "FR-FAM-025/AC-3": false,
		"FR-FAM-025/AC-4": false, "FR-FAM-025/AC-5": false, "FR-FAM-025/AC-6": false,
		"FR-FAM-025/AC-7": false, "FR-FAM-025/AC-8": false, "FR-FAM-025/AC-9": false,
		// Tier 1 FR in tier-2.md.
		"FR-CAL-002/AC-1": true, "FR-CAL-002/AC-2": true, "FR-CAL-002/AC-3": true,
		"INV-01": true, "INV-02": true, "INV-03": true,
		"INV-A35": true,
		"PG-14":   true, "PG-15": false,
		"CTH-01": true, "CTH-03b": true,
		// SCN-08 is "2 (+1)": not required until the lead decides (14 §4.4).
		"SCN-01": true, "SCN-02": true, "SCN-03": false, "SCN-07": false, "SCN-08": false,
		"SEC-01": true, "SEC-02": true, "SEC-158": true,
		// Store-only rows (Apple, Google): known, not required.
		"SEC-143": false, "SEC-145": false, "SEC-146": false,
	}
	if !maps.Equal(got, want) {
		t.Fatalf("requirements =\n%v\nwant\n%v", got, want)
	}
}

func TestReadRequirementsFailsWhenAKindIsEmpty(t *testing.T) {
	dir := copyDocs(t, "testdata/docs")
	writeFile(t, filepath.Join(dir, "12-scenarios.md"), "## 2. Ringkasan skenario\n")
	_, err := readRequirements(dir, false)
	if got, want := errString(err), "no SCN IDs found in the documents"; got != want {
		t.Fatalf("error = %q, want %q", got, want)
	}
}

func TestReadRequirementsFailsOnAnIDCellThatIsNotAnID(t *testing.T) {
	dir := copyDocs(t, "testdata/docs")
	writeFile(t, filepath.Join(dir, "07-financial-domain.md"),
		"| ID | Invarian | Cara menguji |\n|---|---|---|\n| INV-1 | x | y |\n")
	_, err := readRequirements(dir, false)
	want := filepath.Join(dir, "07-financial-domain.md") + `:3: "INV-1" is not an INV ID`
	if got := errString(err); got != want {
		t.Fatalf("error = %q, want %q", got, want)
	}
}

func TestReadRequirementsFailsOnAnUnmappedUntukValue(t *testing.T) {
	dir := copyDocs(t, "testdata/docs")
	writeFile(t, filepath.Join(dir, "10-security-privacy", "01-ancaman-dan-akses.md"),
		"| ID | Persyaratan | Untuk | Tahap | Label | Cara menguji |\n|---|---|---|---|---|---|\n| SEC-01 | x | Watch | Pilot | y | z |\n")
	_, err := readRequirements(dir, false)
	want := filepath.Join(dir, "10-security-privacy", "01-ancaman-dan-akses.md") +
		`:3: "Untuk" value "Watch" has no mapping; add it to untukAppliesToServerOrWeb`
	if got := errString(err); got != want {
		t.Fatalf("error = %q, want %q", got, want)
	}
}

func TestReadRequirementsChecksCountsOnlyWhenAsked(t *testing.T) {
	_, err := readRequirements("testdata/docs", true)
	want := "documents disagree with the checker numbers: Tier 1 FR 2, want 190; INV 3, want 41; " +
		"INV-A 1, want 53; PG known 2, want 15; PG required 1, want 14"
	if got := errString(err); got != want {
		t.Fatalf("error = %q, want %q", got, want)
	}
}

// The real documents hold exactly the checker numbers (14 §3.3, §4.2).
func TestReadRequirementsMatchesCheckerNumbersInRealDocs(t *testing.T) {
	if _, err := readRequirements("../../../docs", true); err != nil {
		t.Fatal(err)
	}
}

func copyDocs(t *testing.T, src string) string {
	t.Helper()
	dst := t.TempDir()
	if err := os.CopyFS(dst, os.DirFS(src)); err != nil {
		t.Fatal(err)
	}
	return dst
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func errString(err error) string {
	if err == nil {
		return "<nil>"
	}
	return err.Error()
}
