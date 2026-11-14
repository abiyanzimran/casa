package household

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// The rule ids of a refused name (plan 1.6 decision 8), reported as the
// field's violation.
const (
	ruleVisibleLength = "name.visible_length"
	ruleCharacters    = "name.characters"
)

const zwj = '\u200D'

// CleanName trims, normalises to NFC, refuses the SEC-155 (h) classes and
// counts visible characters (decision 8); rule is "" when s is accepted, else
// "name.visible_length" or "name.characters". The household name allows 60
// and the display name 40 visible characters (FAM.md FR-FAM-001, note of
// 14 November 2026 [Usulan]).
func CleanName(s string, max int) (clean string, rule string) {
	if !utf8.ValidString(s) {
		return "", ruleCharacters
	}
	clean = norm.NFC.String(strings.TrimSpace(s))
	runes := []rune(clean)
	visible := 0
	for i, r := range runes {
		switch {
		case r == zwj:
			if !inEmojiSequence(runes, i) {
				return "", ruleCharacters
			}
			visible++
		case unicode.Is(unicode.Variation_Selector, r):
			// SEC-155 (h): a selector is accepted only after a base it
			// visibly changes, else "Ayu" and "Ayu"+U+FE0F would be two
			// names that look the same. It is not counted (Mn).
			// ponytail: rare valid sequences outside selectorChangesBase's
			// list are refused (and emoji tag sequences, because the tag
			// characters are Cf); the upgrade path is the Unicode
			// emoji-variation-sequences data.
			if !selectorChangesBase(runes, i) {
				return "", ruleCharacters
			}
		case refused(r):
			return "", ruleCharacters
		case !unicode.In(r, unicode.Mn, unicode.Me):
			// ponytail: a ZWJ sequence counts
			// per base emoji, not per grapheme; a grapheme segmenter waits
			// for a dependency chosen when one is needed.
			visible++
		}
	}
	if visible == 0 || visible > max {
		return "", ruleVisibleLength
	}
	return clean, ""
}

// brailleBlank is U+2800, a symbol (So) that shows as nothing.
const brailleBlank = '\u2800'

// refused is a SEC-155 (h) class: a control character (one line, so line
// breaks too), a line or paragraph separator, a format character, which
// holds the bidirectional controls (U+202A to U+202E, U+2066 to U+2069) and
// invisible characters such as U+200B, or another invisible character: the
// default-ignorable code points that are not format characters (the Hangul
// fillers, the combining grapheme joiner) and the braille blank. The ZWJ
// (only inside an emoji sequence) and the variation selectors (only after a
// base they change visibly) are decided by the caller.
func refused(r rune) bool {
	return r == brailleBlank ||
		unicode.In(r, unicode.Cc, unicode.Cf, unicode.Zl, unicode.Zp, unicode.Other_Default_Ignorable_Code_Point)
}

// selectorChangesBase reports whether the variation selector at i follows a
// base it visibly changes (SEC-155 (h)): FE0E or FE0F after a symbol (So),
// after a keycap base only when U+20E3 follows, or after one of the text
// emoji bases U+203C, U+2049, U+2139, U+3030, U+303D; FE00 to FE0D after a Han
// rune or a math symbol (standardized variation sequences); E0100 to E01EF
// after a Han rune (ideographic); 180B to 180D and 180F after a Mongolian
// rune (free variation selectors).
func selectorChangesBase(runes []rune, i int) bool {
	if i == 0 {
		return false
	}
	sel, base := runes[i], runes[i-1]
	switch {
	case sel == 0xFE0E || sel == 0xFE0F:
		switch {
		case unicode.Is(unicode.So, base):
			return true
		case base == 0x203C || base == 0x2049 || base == 0x2139 || base == 0x3030 || base == 0x303D:
			return true
		case base >= '0' && base <= '9' || base == '#' || base == '*':
			return i+1 < len(runes) && runes[i+1] == 0x20E3
		}
		return false
	case sel <= 0xFE0D && sel >= 0xFE00:
		return unicode.Is(unicode.Han, base) || unicode.Is(unicode.Sm, base)
	case sel >= 0xE0100:
		return unicode.Is(unicode.Han, base)
	default: // 180B to 180D, 180F
		return unicode.Is(unicode.Mongolian, base)
	}
}

// inEmojiSequence reports whether the ZWJ at i joins two emoji: a symbol
// before it, with any variation selectors or skin tone modifiers between,
// and a symbol right after it. ponytail: "emoji" is approximated by the
// Unicode category So; the standard library has no Extended_Pictographic.
func inEmojiSequence(runes []rune, i int) bool {
	if i+1 >= len(runes) || !unicode.Is(unicode.So, runes[i+1]) {
		return false
	}
	for j := i - 1; j >= 0; j-- {
		switch r := runes[j]; {
		case unicode.Is(unicode.So, r):
			return true
		case unicode.In(r, unicode.Mn, unicode.Sk): // variation selector, skin tone modifier
		default:
			return false
		}
	}
	return false
}
