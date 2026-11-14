package household_test

import (
	"bytes"
	"cmp"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/abiyanzimran/casa/server/internal/household"
)

// The name length golden vectors (ES §2 (c); 14 §4.5; ADR-013 item 3; plan
// 1.7 decision 17): the web client's length hint is checked against what
// household.CleanName answers. Run with CASA_GOLDEN_UPDATE=1 to rewrite the
// file (make golden-generate).
const goldenNameLengthFile = "../../../golden/input-shape/name-length.json"

// goldenMaxes are the household name (60) and display name (40) limits
// (FAM.md FR-FAM-001, note of 14 November 2026).
var goldenMaxes = []int{60, 40}

type goldenCase struct {
	input string
	max   int
}

// handPicked names each case so the list can be reviewed; invisible
// characters are Go escapes, never literals (plan 1.7 P12).
func handPicked() map[string]func(max int) string {
	return map[string]func(int) string{
		"empty":                 func(int) string { return "" },
		"space":                 func(int) string { return " " },
		"no-break space":        func(int) string { return "\u00a0" },
		"ideographic space pad": func(int) string { return "\u3000a\u3000" },
		"one letter":            func(int) string { return "a" },
		"exactly max letters":   func(m int) string { return strings.Repeat("a", m) },
		"max plus one letters":  func(m int) string { return strings.Repeat("a", m+1) },
		"e acute":               func(int) string { return "\u00e9" },
		"max e acute":           func(m int) string { return strings.Repeat("\u00e9", m) },
		"family zwj sequence":   func(int) string { return "\U0001F468\u200d\U0001F469\u200d\U0001F467" },
		"name with selector":    func(int) string { return "Ayu\ufe0f" },
		"zero width space":      func(int) string { return "\u200b" },
		"byte order mark":       func(int) string { return "\ufeff" },
		"bidi override inside":  func(int) string { return "a\u202eb" },
	}
}

// The full random pool (mode 3): ASCII, combining marks, emoji, ZWJ, Unicode
// spaces and a variation selector. Most strings from it are refused with
// name.characters; they pin the server-only rule.
var goldenPool = []rune{
	'a', 'B', 'z', 'Q', '0', '9', '-', '.', ' ',
	'\u0301', '\u0323', // combining marks
	'\U0001F468', '\U0001F469', '\U0001F600', '\u2764', // emoji
	'\u200d',                     // ZWJ
	'\u3000', '\u00a0', '\u2003', // Unicode spaces
	'\ufe0f', // variation selector
}

// goldenUnit is a piece of a valid-only string and the visible characters
// CleanName counts for it: a combining mark is not counted, a ZWJ sequence
// counts per base emoji (names.go).
type goldenUnit struct {
	text    string
	visible int
}

var goldenValidUnits = []goldenUnit{
	{"a", 1}, {"B", 1}, {"z", 1}, {"\u00e9", 1}, {"7", 1}, {"-", 1}, {" ", 1},
	{"\U0001F600", 1}, {"\u2764", 1},
	{"a\u0301", 1}, {"x\u0323", 1}, // a base with a combining mark
	{"\U0001F468\u200d\U0001F469", 2}, {"\U0001F468\u200d\U0001F469\u200d\U0001F467", 3},
}

// validString builds a string of exactly target visible characters from
// valid units only; a space is never first or last (CleanName trims it).
func validString(rng *rand.Rand, target int) string {
	var b strings.Builder
	for vis := 0; vis < target; {
		u := goldenValidUnits[rng.IntN(len(goldenValidUnits))]
		if vis+u.visible > target || (u.text == " " && (vis == 0 || vis == target-1)) {
			u = goldenUnit{"a", 1}
		}
		b.WriteString(u.text)
		vis += u.visible
	}
	return b.String()
}

// goldenCases are the hand-picked cases, then 200 random strings from
// rand.New(rand.NewPCG(1, 7)). Per string a mode is drawn: modes 0 and 1 are
// valid-only strings within 3 visible characters of the limit, mode 2 valid-only
// 30 to 70, mode 3 the full pool (length 0 to 70). Without the valid modes
// refusals would hide the length rule (ES §2 (c)).
func goldenCases() []goldenCase {
	var cases []goldenCase
	for _, max := range goldenMaxes {
		for _, f := range handPicked() {
			cases = append(cases, goldenCase{f(max), max})
		}
	}
	// Map order is random; sort for byte-determinism.
	slices.SortFunc(cases, func(a, b goldenCase) int {
		return cmp.Or(cmp.Compare(b.max, a.max), cmp.Compare(a.input, b.input))
	})
	rng := rand.New(rand.NewPCG(1, 7)) //nolint:gosec // G404: a fixed seed makes the vectors reproducible; not a secret
	for range 200 {
		max := goldenMaxes[rng.IntN(len(goldenMaxes))]
		switch mode := rng.IntN(4); mode {
		case 0, 1:
			cases = append(cases, goldenCase{validString(rng, max-3+rng.IntN(7)), max})
		case 2:
			cases = append(cases, goldenCase{validString(rng, 30+rng.IntN(41)), max})
		default:
			rs := make([]rune, rng.IntN(71))
			for i := range rs {
				rs[i] = goldenPool[rng.IntN(len(goldenPool))]
			}
			cases = append(cases, goldenCase{string(rs), max})
		}
	}
	return cases
}

// jsonString escapes every rune outside printable ASCII as \uXXXX so the file
// can be reviewed and carries no invisible character.
func jsonString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for _, r := range s {
		switch {
		case r == '"' || r == '\\':
			b.WriteByte('\\')
			b.WriteRune(r)
		case r >= 0x20 && r < 0x7f:
			b.WriteRune(r)
		case r >= 0x10000:
			r -= 0x10000
			fmt.Fprintf(&b, `\u%04x\u%04x`, 0xd800+(r>>10), 0xdc00+(r&0x3ff))
		default:
			fmt.Fprintf(&b, `\u%04x`, r)
		}
	}
	b.WriteByte('"')
	return b.String()
}

func goldenBytes() []byte {
	var b bytes.Buffer
	b.WriteString(`{"source":"household.CleanName (FAM.md FR-FAM-001 note of 14 November 2026)","cases":[` + "\n")
	cases := goldenCases()
	for i, c := range cases {
		_, rule := household.CleanName(c.input, c.max)
		fmt.Fprintf(&b, `{"input":%s,"max":%d,"rule":%s}`, jsonString(c.input), c.max, jsonString(rule))
		if i < len(cases)-1 {
			b.WriteByte(',')
		}
		b.WriteByte('\n')
	}
	b.WriteString("]}\n")
	return b.Bytes()
}

func TestGoldenNameLength(t *testing.T) {
	path := filepath.FromSlash(goldenNameLengthFile)
	want := goldenBytes()

	t.Run("the vector file equals what CleanName produces", func(t *testing.T) {
		if os.Getenv("CASA_GOLDEN_UPDATE") == "1" {
			if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, want, 0o600); err != nil {
				t.Fatal(err)
			}
		}
		got, err := os.ReadFile(path) //nolint:gosec // G304: path is a package constant
		if err != nil {
			t.Fatalf("read vector file: %v (run make golden-generate)", err)
		}
		if !bytes.Equal(got, want) {
			t.Fatal("vector file differs from CleanName output (run make golden-generate)")
		}
		if !utf8.Valid(got) || bytes.IndexFunc(got, func(r rune) bool { return r > 0x7e }) >= 0 {
			t.Fatal("vector file must be printable ASCII (escapes only)")
		}
	})

	t.Run("the file holds every hand-picked case", func(t *testing.T) {
		got, err := os.ReadFile(path) //nolint:gosec // G304: path is a package constant
		if err != nil {
			t.Fatal(err)
		}
		for _, max := range goldenMaxes {
			for name, f := range handPicked() {
				_, rule := household.CleanName(f(max), max)
				line := fmt.Sprintf(`{"input":%s,"max":%d,"rule":%s}`, jsonString(f(max)), max, jsonString(rule))
				if !bytes.Contains(got, []byte(line)) {
					t.Errorf("missing case %q for max %d", name, max)
				}
			}
		}
	})

	t.Run("fixed rules hold for both limits", func(t *testing.T) {
		for _, max := range goldenMaxes {
			for _, c := range []struct{ name, in, rule string }{
				{"exactly max", strings.Repeat("a", max), ""},
				{"max plus one", strings.Repeat("a", max+1), "name.visible_length"},
				{"zero width space", "\u200b", "name.characters"},
				{"family sequence", "\U0001F468\u200d\U0001F469\u200d\U0001F467", ""},
			} {
				if _, rule := household.CleanName(c.in, max); rule != c.rule {
					t.Errorf("%s, max %d: rule %q, want %q", c.name, max, rule, c.rule)
				}
			}
		}
	})
}
