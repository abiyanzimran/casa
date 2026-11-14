package household_test

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
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

// The random pool: ASCII, combining marks, emoji, ZWJ, Unicode spaces.
var goldenPool = []rune{
	'a', 'B', 'z', 'Q', '0', '9', '-', '.', ' ',
	'\u0301', '\u0323', // combining marks
	'\U0001F468', '\U0001F469', '\U0001F600', '\u2764', // emoji
	'\u200d',                               // ZWJ
	'\u3000', '\u00a0', '\u2003', '\ufe0f', // Unicode spaces, selector
}

func goldenCases() []goldenCase {
	var cases []goldenCase
	for _, max := range goldenMaxes {
		for _, f := range handPicked() {
			cases = append(cases, goldenCase{f(max), max})
		}
	}
	// Map order is random; sort for byte-determinism.
	sortCases(cases)
	rng := rand.New(rand.NewPCG(1, 7)) //nolint:gosec // G404: a fixed seed makes the vectors reproducible; not a secret
	for range 200 {
		n := rng.IntN(71)
		rs := make([]rune, n)
		for i := range rs {
			rs[i] = goldenPool[rng.IntN(len(goldenPool))]
		}
		cases = append(cases, goldenCase{string(rs), goldenMaxes[rng.IntN(len(goldenMaxes))]})
	}
	return cases
}

func sortCases(c []goldenCase) {
	for i := 1; i < len(c); i++ {
		for j := i; j > 0 && less(c[j], c[j-1]); j-- {
			c[j], c[j-1] = c[j-1], c[j]
		}
	}
}

func less(a, b goldenCase) bool {
	if a.max != b.max {
		return a.max > b.max
	}
	return a.input < b.input
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
}
