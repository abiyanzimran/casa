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
		case unicode.Is(unicode.Variation_Selector, r):
			// The security goal is no invisible name (SEC-155 (h)): a
			// selector after a visible base is harmless (emoji, keycap,
			// CJK ideographic and Mongolian sequences); one that is first
			// or follows an invisible or space rune could stand alone, and
			// one after a plain letter makes two names that look the same
			// ("Ayu" and "Ayu"+U+FE0F). It is not counted (Mn).
			// ponytail: emoji tag sequences (subdivision flags) stay refused
			// because the tag characters are Cf; allow them by a rule for
			// a tag run after a flag base if a user needs one.
			if i == 0 || unicode.IsSpace(runes[i-1]) || refused(runes[i-1]) || invisibleAfterLetter(runes[i-1]) {
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

// invisibleAfterLetter reports whether a variation selector after base changes
// nothing visible (SEC-155 (h)): base is a letter, other than a Han or
// Mongolian one (ideographic and free variation selectors) or U+2139, a
// letter-category emoji base.
func invisibleAfterLetter(base rune) bool {
	return unicode.IsLetter(base) && base != '\u2139' &&
		!unicode.Is(unicode.Han, base) && !unicode.Is(unicode.Mongolian, base)
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
