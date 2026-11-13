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
		case refused(r):
			return "", ruleCharacters
		case !unicode.In(r, unicode.Mn, unicode.Me):
			// Variation selectors are Mn. ponytail: a ZWJ sequence counts
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

// refused is a SEC-155 (h) class: a control character (one line, so line
// breaks too), a line or paragraph separator, or a format character, which
// holds the bidirectional controls (U+202A to U+202E, U+2066 to U+2069) and
// the invisible characters such as U+200B. The one format character allowed,
// a ZWJ inside an emoji sequence, is decided by the caller.
func refused(r rune) bool {
	return unicode.In(r, unicode.Cc, unicode.Cf, unicode.Zl, unicode.Zp)
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
