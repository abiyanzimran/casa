package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Requirement is one ID that a test may name. Required means a release, and
// the scope of a finished step, needs a test for it (14 §4.13 item 1).
type Requirement struct {
	ID       string
	Required bool
}

// Checker numbers: what the documents are known to hold, so a reader that
// silently stops matching after a format change fails instead of passing.
const (
	tier1FRCount           = 190 // 05 §6 "Hitungan yang berlaku"; 14 §3.3
	invariantCount         = 41  // 07 §6 INV-01 to INV-41; 14 §4.2
	approvalInvariantCount = 53  // 06 INV-A01 to INV-A53; 14 §4.2
	guardrailCount         = 15  // 06 §10 PG-01 to PG-15
	requiredGuardrailCount = 14  // PG-15 only from FS-3 (14 §4.2, §8)
)

// ID shapes, one per kind (14 §4.1 item 3). CTH has letter variants such as
// CTH-03b (14 §4.4).
const (
	frPattern                = `FR-[A-Z]+-\d{3}`
	approvalInvariantPattern = `INV-A\d{2}`
	invariantPattern         = `INV-\d{2}`
	guardrailPattern         = `PG-\d{2}`
	workedExamplePattern     = `CTH-\d{2,}[a-z]?`
	scenarioPattern          = `SCN-\d{2}`
	securityPattern          = `SEC-\d{2,3}`
)

// frHeadingPrefix starts a line that must be a well-formed FR heading.
const frHeadingPrefix = "### FR-"

// kinds lists every ID kind; each must be non-empty in the documents.
var kinds = []string{"FR", "INV", "INV-A", "PG", "CTH", "SCN", "SEC"}

// deferredGuardrails are known PG IDs outside the FS-1 to FS-2 release scope.
var deferredGuardrails = map[string]bool{"PG-15": true} // FS-3 (14 §4.2, §8)

// untukAppliesToServerOrWeb maps each value of the SEC "Untuk" column (or
// "Toko" in the store tables of 10.4 §2) to whether the requirement applies
// to the server or the web client, and so is due before GR-2 when its stage
// is "Pilot" (10 §0 "Pilot per klien"). The "Untuk" values name the
// document or owner that implements the requirement (10 §0 "Kolom Untuk"),
// not a platform, and a requirement that names no platform applies to the web
// client as written (10 §0); store values (Apple, Google) are store policy for
// the native apps (mapping confirmed by lead ruling 13 Nov 2026, 1.1c). An
// unmapped value fails the check.
var untukAppliesToServerOrWeb = map[string]bool{
	"05": true, "06": true, "07": true, "08": true, "09": true, "10": true, "14": true,
	"Ops": true, "ADR-018": true, "ADR-019": true, "ADR-030": true, "Semua jalur": true,
	"Apple": false, "Apple 5.1.1(v)": false, "Google": false,
}

var (
	frHeading      = regexp.MustCompile(`^### (` + frPattern + `) — `)
	tierRow        = regexp.MustCompile(`^\| Tier / MoSCoW \| ([12]) / `)
	acBullet       = regexp.MustCompile(`^- \*\*(AC-\d+)\*\*`)
	sectionHeading = regexp.MustCompile(`^#{1,3} `)
	// 07 §7 writes most examples as "#### CTH-01 — …" and some as numbered
	// sections, "### 7.10 CTH-37 — …"; workedExampleLine catches both so a
	// third shape fails instead of being skipped.
	workedExampleLine = regexp.MustCompile(`^#{1,6} (?:[\d.]+ )?CTH-`)
	workedExampleHead = regexp.MustCompile(`^#{3,4} (?:\d+(?:\.\d+)* )?(` + workedExamplePattern + `) — `)
	invariantHeader   = regexp.MustCompile(`^\| ID \| Invarian \| Cara menguji \|`)
	guardrailHeader   = regexp.MustCompile(`^\| ID \| Pagar \| Cara menguji \|`)
	scenarioHeader    = regexp.MustCompile(`^\| SCN \| Judul \| Tier \|`)
	securityHeader    = regexp.MustCompile(`^\| ID \| Persyaratan \| Untuk \| Tahap \|`)
	storeHeader       = regexp.MustCompile(`^\| ID \| Persyaratan \| Toko \| Tahap \|`)
	scenarioTier      = regexp.MustCompile(`^([123])(?: \(\+[123](?:, [123])*\))?$`)
	untukSeparator    = regexp.MustCompile(`[,;]`)
	idCellInvariant   = regexp.MustCompile(`^` + invariantPattern + `$`)
	idCellApprovalInv = regexp.MustCompile(`^` + approvalInvariantPattern + `$`)
	idCellGuardrail   = regexp.MustCompile(`^` + guardrailPattern + `$`)
	idCellScenario    = regexp.MustCompile(`^` + scenarioPattern + `$`)
	idCellSecurity    = regexp.MustCompile(`^` + securityPattern + `$`)
)

type kindCount struct{ known, required int }

func kindOf(id string) string {
	if strings.HasPrefix(id, "INV-A") {
		return "INV-A"
	}
	kind, _, _ := strings.Cut(id, "-")
	return kind
}

func countByKind(reqs []Requirement) map[string]kindCount {
	counts := map[string]kindCount{}
	for _, r := range reqs {
		c := counts[kindOf(r.ID)]
		c.known++
		if r.Required {
			c.required++
		}
		counts[kindOf(r.ID)] = c
	}
	return counts
}

// readRequirements reads every ID defined in the planning documents under
// docsDir (14 §4.13 item 1), sorted by ID. Only definition lines count:
// FR headings and their AC bullets, the rows of the INV, INV-A, PG, SCN and
// SEC definition tables, and CTH headings, all outside fenced code blocks.
// verifyCounts compares the result with the checker numbers; only tests on
// excerpts pass false.
func readRequirements(docsDir string, verifyCounts bool) ([]Requirement, error) {
	reqs := requirementSet{}
	tier1FRs, err := readFunctionalRequirements(filepath.Join(docsDir, "05-functional-requirements"), reqs)
	if err != nil {
		return nil, err
	}
	readers := []func(string, requirementSet) error{
		readInvariants, readApprovalInvariantsAndGuardrails, readWorkedExamples, readScenarios, readSecurityRequirements,
	}
	for _, read := range readers {
		if err := read(docsDir, reqs); err != nil {
			return nil, err
		}
	}

	list := make([]Requirement, 0, len(reqs))
	for id, required := range reqs {
		list = append(list, Requirement{ID: id, Required: required})
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })

	counts := countByKind(list)
	for _, kind := range kinds {
		if counts[kind].known == 0 {
			return nil, fmt.Errorf("no %s IDs found in the documents", kind)
		}
		if counts[kind].required == 0 {
			return nil, fmt.Errorf("no required %s IDs found in the documents", kind)
		}
	}
	if verifyCounts {
		if err := checkCounts(tier1FRs, counts); err != nil {
			return nil, err
		}
	}
	return list, nil
}

func checkCounts(tier1FRs int, counts map[string]kindCount) error {
	checks := []struct {
		name      string
		got, want int
	}{
		{"Tier 1 FR", tier1FRs, tier1FRCount},
		{"INV", counts["INV"].known, invariantCount},
		{"INV-A", counts["INV-A"].known, approvalInvariantCount},
		{"PG known", counts["PG"].known, guardrailCount},
		{"PG required", counts["PG"].required, requiredGuardrailCount},
	}
	var diffs []string
	for _, c := range checks {
		if c.got != c.want {
			diffs = append(diffs, fmt.Sprintf("%s %d, want %d", c.name, c.got, c.want))
		}
	}
	if diffs != nil {
		return errors.New("documents disagree with the checker numbers: " + strings.Join(diffs, "; "))
	}
	return nil
}

// requirementSet maps an ID to whether it is required.
type requirementSet map[string]bool

func (s requirementSet) add(id string, required bool, path string, line int) error {
	if _, ok := s[id]; ok {
		return fmt.Errorf("%s:%d: %s is defined twice", path, line, id)
	}
	s[id] = required
	return nil
}

// readFunctionalRequirements reads FR-…/AC-n from every file in dir. An AC
// belongs to the FR heading above it until the next heading of level 1 to 3;
// the FR's tier is the first number of its "Tier / MoSCoW" row, the row the
// 05 §6 count is taken from. Only Tier 1 is required.
func readFunctionalRequirements(dir string, reqs requirementSet) (int, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return 0, err
	}
	type fr struct {
		tier string
		acs  []string
		path string
		line int
	}
	var order []string
	frs := map[string]*fr{}
	for _, path := range paths {
		lines, err := readDocumentLines(path)
		if err != nil {
			return 0, err
		}
		var current *fr
		for i, line := range lines {
			n := i + 1
			if strings.HasPrefix(line, frHeadingPrefix) {
				m := frHeading.FindStringSubmatch(line)
				if m == nil {
					return 0, fmt.Errorf("%s:%d: FR heading does not start with an FR ID and \" — \"", path, n)
				}
				if _, ok := frs[m[1]]; ok {
					return 0, fmt.Errorf("%s:%d: %s is defined twice", path, n, m[1])
				}
				current = &fr{path: path, line: n}
				frs[m[1]] = current
				order = append(order, m[1])
				continue
			}
			if sectionHeading.MatchString(line) {
				current = nil
				continue
			}
			if m := acBullet.FindStringSubmatch(line); m != nil {
				if current == nil {
					return 0, fmt.Errorf("%s:%d: %s is not under an FR heading", path, n, m[1])
				}
				current.acs = append(current.acs, m[1])
				continue
			}
			if current != nil && strings.HasPrefix(line, "| Tier / MoSCoW |") {
				m := tierRow.FindStringSubmatch(line)
				if m == nil {
					return 0, fmt.Errorf("%s:%d: \"Tier / MoSCoW\" does not start with 1 or 2", path, n)
				}
				current.tier = m[1]
			}
		}
	}

	tier1 := 0
	for _, id := range order {
		f := frs[id]
		if f.tier == "" {
			return 0, fmt.Errorf("%s:%d: %s has no \"Tier / MoSCoW\" row", f.path, f.line, id)
		}
		if f.tier == "1" {
			tier1++
		}
		for _, ac := range f.acs {
			if err := reqs.add(id+"/"+ac, f.tier == "1", f.path, f.line); err != nil {
				return 0, err
			}
		}
	}
	return tier1, nil
}

func readInvariants(docsDir string, reqs requirementSet) error {
	path := filepath.Join(docsDir, "07-financial-domain.md")
	rows, err := readTable(path, invariantHeader)
	if err != nil {
		return err
	}
	return addIDColumn(rows, path, idCellInvariant, "INV", reqs, func(string) bool { return true })
}

func readApprovalInvariantsAndGuardrails(docsDir string, reqs requirementSet) error {
	path := filepath.Join(docsDir, "06-activity-approval-architecture.md")
	rows, err := readTable(path, invariantHeader)
	if err != nil {
		return err
	}
	if err := addIDColumn(rows, path, idCellApprovalInv, "INV-A", reqs, func(string) bool { return true }); err != nil {
		return err
	}
	rows, err = readTable(path, guardrailHeader)
	if err != nil {
		return err
	}
	return addIDColumn(rows, path, idCellGuardrail, "PG", reqs, func(id string) bool { return !deferredGuardrails[id] })
}

func addIDColumn(rows []tableRow, path string, shape *regexp.Regexp, kind string, reqs requirementSet, required func(string) bool) error {
	for _, row := range rows {
		id := row.cells[0]
		if !shape.MatchString(id) {
			return fmt.Errorf("%s:%d: %q is not an %s ID", path, row.line, id, kind)
		}
		if err := reqs.add(id, required(id), path, row.line); err != nil {
			return err
		}
	}
	return nil
}

func readWorkedExamples(docsDir string, reqs requirementSet) error {
	path := filepath.Join(docsDir, "07-financial-domain.md")
	lines, err := readDocumentLines(path)
	if err != nil {
		return err
	}
	for i, line := range lines {
		if !workedExampleLine.MatchString(line) {
			continue
		}
		m := workedExampleHead.FindStringSubmatch(line)
		if m == nil {
			return fmt.Errorf("%s:%d: CTH heading does not start with a CTH ID and \" — \"", path, i+1)
		}
		if err := reqs.add(m[1], true, path, i+1); err != nil {
			return err
		}
	}
	return nil
}

// readScenarios reads the 12 §2 table. A scenario is required when the first
// tier in its Tier column is 1 ("1", "1 (+2, 3)"), or when the column holds
// "(+1)": its Tier 1 part is tested in FS-1 too (14 §4.4; lead ruling 13 Nov
// 2026, 1.1c, which includes SCN-39 as 12 §2 marks it).
func readScenarios(docsDir string, reqs requirementSet) error {
	path := filepath.Join(docsDir, "12-scenarios.md")
	rows, err := readTable(path, scenarioHeader)
	if err != nil {
		return err
	}
	for _, row := range rows {
		id := row.cells[0]
		if !idCellScenario.MatchString(id) {
			return fmt.Errorf("%s:%d: %q is not an SCN ID", path, row.line, id)
		}
		m := scenarioTier.FindStringSubmatch(cell(row, 2))
		if m == nil {
			return fmt.Errorf("%s:%d: %s has tier %q, not \"N\" or \"N (+N, …)\"", path, row.line, id, cell(row, 2))
		}
		required := m[1] == "1" || strings.Contains(cell(row, 2), "(+1)")
		if err := reqs.add(id, required, path, row.line); err != nil {
			return err
		}
	}
	return nil
}

// readSecurityRequirements reads the SEC tables of 10. Every clause of the
// required rule is a lead ruling of 13 Nov 2026 (1.1c).
func readSecurityRequirements(docsDir string, reqs requirementSet) error {
	rows, err := readSecurityRows(docsDir)
	if err != nil {
		return err
	}
	for _, row := range rows {
		id := row.cells[0]
		if !idCellSecurity.MatchString(id) {
			return fmt.Errorf("%s:%d: %q is not a SEC ID", row.path, row.line, id)
		}
		applies := false
		for _, value := range untukValues(cell(row.tableRow, 2)) {
			mapped, ok := untukAppliesToServerOrWeb[value]
			if !ok {
				return fmt.Errorf("%s:%d: \"Untuk\" value %q has no mapping; add it to untukAppliesToServerOrWeb", row.path, row.line, value)
			}
			applies = applies || mapped
		}
		// The stage, trimmed of spaces and bold markers, begins with "Pilot"
		// (lead ruling 13 Nov 2026, 1.1c; 10 §0 "Pilot per klien").
		stage := cell(row.tableRow, 3)
		pilot := strings.HasPrefix(strings.TrimLeft(strings.TrimSpace(stage), "*"), "Pilot")
		// "bila" makes the row conditional on a later build decision (lead
		// ruling 13 Nov 2026, 1.1c).
		conditional := strings.Contains(stage, "bila")
		// Store rows ("Toko" tables, 10.4 §2) gate the native store tracks
		// (14 §3.8, §3.9); the web client is in no store (lead ruling 13 Nov
		// 2026, 1.1c).
		required := applies && pilot && !conditional && !row.store
		if err := reqs.add(id, required, row.path, row.line); err != nil {
			return err
		}
	}
	return nil
}

type securityRow struct {
	tableRow
	path  string
	store bool // from a "Toko" table
}

func readSecurityRows(docsDir string) ([]securityRow, error) {
	paths, err := filepath.Glob(filepath.Join(docsDir, "10-security-privacy", "*.md"))
	if err != nil {
		return nil, err
	}
	var rows []securityRow
	for _, path := range paths {
		for _, kind := range []struct {
			header *regexp.Regexp
			store  bool
		}{{securityHeader, false}, {storeHeader, true}} {
			table, err := readTable(path, kind.header)
			if err != nil {
				return nil, err
			}
			for _, row := range table {
				rows = append(rows, securityRow{tableRow: row, path: path, store: kind.store})
			}
		}
	}
	return rows, nil
}

// untukValues splits one "Untuk" cell, such as "09, 08" or
// "Apple 5.1.1(v); Google", into its values.
func untukValues(cellText string) []string {
	var values []string
	for _, v := range untukSeparator.Split(cellText, -1) {
		if v = strings.TrimSpace(v); v != "" {
			values = append(values, v)
		}
	}
	return values
}

type tableRow struct {
	line  int
	cells []string
}

func cell(row tableRow, i int) string {
	if i < len(row.cells) {
		return row.cells[i]
	}
	return ""
}

// readTable returns the body rows of every table in path whose header line
// matches header. A table ends at the first line that does not start with "|".
func readTable(path string, header *regexp.Regexp) ([]tableRow, error) {
	lines, err := readDocumentLines(path)
	if err != nil {
		return nil, err
	}
	var rows []tableRow
	inTable := false
	for i, line := range lines {
		switch {
		case header.MatchString(line):
			inTable = true
		case !strings.HasPrefix(line, "|"):
			inTable = false
		case inTable && !strings.HasPrefix(line, "|---") && !strings.HasPrefix(line, "| ---"):
			rows = append(rows, tableRow{line: i + 1, cells: splitCells(line)})
		}
	}
	return rows, nil
}

func splitCells(line string) []string {
	const escapedPipe = "\x00"
	line = strings.ReplaceAll(strings.TrimSpace(line), `\|`, escapedPipe)
	parts := strings.Split(strings.TrimSuffix(strings.TrimPrefix(line, "|"), "|"), "|")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(strings.ReplaceAll(p, escapedPipe, "|"))
	}
	return parts
}

// readDocumentLines returns the lines of a markdown file with every line of a
// fenced code block, fences included, blanked, so line numbers stay true and
// nothing inside a fence counts as a definition.
func readDocumentLines(path string) ([]string, error) {
	content, err := os.ReadFile(path) //nolint:gosec // G304: the path is a planning document under the -docs flag of a CI tool, not user input
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	inFence := false
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inFence = !inFence
			lines[i] = ""
			continue
		}
		if inFence {
			lines[i] = ""
		}
	}
	return lines, nil
}
