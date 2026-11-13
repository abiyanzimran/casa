package household_test

import (
	"strings"
	"testing"

	"github.com/abiyanzimran/casa/server/internal/household"
)

// TestCleanName is the name rule of FAM.md FR-FAM-001 (note of 14 November
// 2026, input form): trimmed, NFC, one line, the SEC-155 (h) character
// classes refused, visible characters counted (plan 1.6 decision 8).
func TestCleanName(t *testing.T) {
	family := "\U0001F468\u200D\U0001F469\u200D\U0001F467" // man ZWJ woman ZWJ girl: 3 base emoji
	cases := []struct {
		name, in  string
		max       int
		want      string
		wantRule  string
		wantClean bool
	}{
		{"trimmed", "  Rumah  ", 60, "Rumah", "", true},
		{"NFD becomes NFC", "Cafe\u0301", 60, "Caf\u00E9", "", true},
		{"60 visible with combining marks", strings.Repeat("a\u0301\u0327", 60), 60, strings.Repeat("\u00E1\u0327", 60), "", true},
		{"61 visible", strings.Repeat("a", 61), 60, "", "name.visible_length", false},
		{"empty", "", 60, "", "name.visible_length", false},
		{"spaces only", " \t\u3000 ", 60, "", "name.visible_length", false},
		{"right-to-left override", "Rumah\u202E", 60, "", "name.characters", false},
		{"line separator", "Rumah\u2028Kita", 60, "", "name.characters", false},
		{"paragraph separator", "Rumah\u2029Kita", 60, "", "name.characters", false},
		{"bell", "Rumah\u0007", 60, "", "name.characters", false},
		{"new line", "Rumah\nKita", 60, "", "name.characters", false},
		{"zero width space", "Ru\u200Bmah", 60, "", "name.characters", false},
		{"isolate", "\u2066Rumah\u2069", 60, "", "name.characters", false},
		{"ZWJ outside an emoji sequence", "Ru\u200Dmah", 60, "", "name.characters", false},
		{"ZWJ at the end", "\U0001F468\u200D", 60, "", "name.characters", false},
		{"invalid UTF-8", "Rumah\xff", 60, "", "name.characters", false},
		// SEC-155 (h): invisible characters that are not format characters.
		{"Hangul filler alone", "\u3164", 60, "", "name.characters", false},
		{"Hangul fillers", "\u3164\u3164\u3164", 60, "", "name.characters", false},
		{"Hangul choseong and jungseong fillers", "\u115F\u1160", 60, "", "name.characters", false},
		{"halfwidth Hangul filler", "\uFFA0", 60, "", "name.characters", false},
		{"combining grapheme joiner", "Ayu\u034F", 60, "", "name.characters", false},
		{"variation selector outside an emoji", "Ayu\uFE0F", 60, "", "name.characters", false},
		{"variation selector at the start", "\uFE0FAyu", 60, "", "name.characters", false},
		{"braille blank", "\u2800", 60, "", "name.characters", false},
		{"braille blank inside", "Ayu\u2800Budi", 60, "", "name.characters", false},
		{"emoji with its variation selector", "\u2764\uFE0F", 60, "\u2764\uFE0F", "", true},
		{"emoji ZWJ sequence", "Rumah " + family, 60, "Rumah " + family, "", true},
		// ponytail ceiling: a ZWJ sequence counts per base emoji, not per
		// grapheme, so 20 families are 60 visible characters.
		{"emoji ZWJ sequence counted per base emoji", strings.Repeat(family, 20), 60, strings.Repeat(family, 20), "", true},
		{"emoji ZWJ sequence over the limit", strings.Repeat(family, 20) + "a", 60, "", "name.visible_length", false},
		{"emoji with variation selector", "\u2764\uFE0F\u200D\U0001F525", 60, "\u2764\uFE0F\u200D\U0001F525", "", true},
		{"display name 40", strings.Repeat("b", 40), 40, strings.Repeat("b", 40), "", true},
		{"display name 41", strings.Repeat("b", 41), 40, "", "name.visible_length", false},
		{"non-Latin script", "  \u0930\u0938\u094B\u0908 \u0918\u0930  ", 60, "\u0930\u0938\u094B\u0908 \u0918\u0930", "", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			clean, rule := household.CleanName(c.in, c.max)
			if rule != c.wantRule {
				t.Fatalf("rule = %q, want %q", rule, c.wantRule)
			}
			if c.wantClean && clean != c.want {
				t.Fatalf("clean = %+q, want %+q", clean, c.want)
			}
		})
	}
}
