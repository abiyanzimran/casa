package main

import (
	"fmt"
	"slices"
	"strings"
)

// A small TypeScript lexer, enough to find test titles outside comments,
// strings, template literals and regular expressions (code-reviewer, 1.1c
// fix round 1). It does not parse TypeScript; it only tokenizes.

type tokenKind int

const (
	identToken    tokenKind = iota // identifier or keyword
	stringToken                    // '…' or "…"; text is the value
	templateToken                  // `…`; text is the raw literal text
	valueToken                     // number or regular expression literal
	otherToken                     // one punctuation character
)

type tsToken struct {
	kind tokenKind
	text string
	line int
}

// regexAfterKeyword lists the keywords after which "/" starts a regular
// expression rather than a division.
var regexAfterKeyword = map[string]bool{
	"return": true, "typeof": true, "case": true, "in": true, "of": true, "new": true, "delete": true,
	"void": true, "throw": true, "instanceof": true, "yield": true, "await": true,
}

type tsLexer struct {
	src    string
	pos    int
	line   int
	tokens []tsToken
}

func lexTypeScript(src string) ([]tsToken, error) {
	l := &tsLexer{src: src, line: 1}
	if err := l.code(false); err != nil {
		return nil, err
	}
	return l.tokens, nil
}

// code lexes until the end of the source or, inside a template
// substitution, until its closing "}".
func (l *tsLexer) code(inSubstitution bool) error {
	depth := 0
	for l.pos < len(l.src) {
		c := l.src[l.pos]
		switch {
		case c == '\n':
			l.line++
			l.pos++
		case c == ' ' || c == '\t' || c == '\r':
			l.pos++
		case strings.HasPrefix(l.src[l.pos:], "//"):
			for l.pos < len(l.src) && l.src[l.pos] != '\n' {
				l.pos++
			}
		case strings.HasPrefix(l.src[l.pos:], "/*"):
			end := strings.Index(l.src[l.pos+2:], "*/")
			if end < 0 {
				return fmt.Errorf("line %d: unterminated block comment", l.line)
			}
			l.advance(end + 4)
		case c == '\'' || c == '"':
			if err := l.quoted(c); err != nil {
				return err
			}
		case c == '`':
			if err := l.template(); err != nil {
				return err
			}
		case c == '/' && l.regexAllowed():
			l.regex()
		case isIdentStart(c):
			start := l.pos
			for l.pos < len(l.src) && isIdentPart(l.src[l.pos]) {
				l.pos++
			}
			l.tokens = append(l.tokens, tsToken{identToken, l.src[start:l.pos], l.line})
		case c >= '0' && c <= '9':
			start := l.pos
			for l.pos < len(l.src) && (isIdentPart(l.src[l.pos]) || l.src[l.pos] == '.') {
				l.pos++
			}
			l.tokens = append(l.tokens, tsToken{valueToken, l.src[start:l.pos], l.line})
		default:
			if c == '{' {
				depth++
			}
			if c == '}' {
				if inSubstitution && depth == 0 {
					l.pos++
					return nil
				}
				depth--
			}
			l.tokens = append(l.tokens, tsToken{otherToken, string(c), l.line})
			l.pos++
		}
	}
	if inSubstitution {
		return fmt.Errorf("line %d: unterminated template substitution", l.line)
	}
	return nil
}

// advance moves n bytes ahead, counting the newlines it passes.
func (l *tsLexer) advance(n int) {
	l.line += strings.Count(l.src[l.pos:l.pos+n], "\n")
	l.pos += n
}

func (l *tsLexer) quoted(quote byte) error {
	line := l.line
	l.pos++
	var value strings.Builder
	for l.pos < len(l.src) {
		c := l.src[l.pos]
		switch {
		case c == quote:
			l.pos++
			l.tokens = append(l.tokens, tsToken{stringToken, value.String(), line})
			return nil
		case c == '\n':
			return fmt.Errorf("line %d: unterminated string", line)
		case c == '\\' && l.pos+1 < len(l.src):
			l.advance(1)
			value.WriteByte(l.src[l.pos])
			l.advance(1)
		default:
			value.WriteByte(c)
			l.pos++
		}
	}
	return fmt.Errorf("line %d: unterminated string", line)
}

// template lexes a template literal; the code of each ${…} substitution is
// lexed too, so a nested string or template does not end the literal.
func (l *tsLexer) template() error {
	// The token goes in before the tokens of its substitutions, so a call
	// sees the template as its first argument; its text is set at the end.
	index := len(l.tokens)
	l.tokens = append(l.tokens, tsToken{templateToken, "", l.line})
	line := l.line
	l.pos++
	var raw strings.Builder
	for l.pos < len(l.src) {
		c := l.src[l.pos]
		switch {
		case c == '`':
			l.pos++
			l.tokens[index].text = raw.String()
			return nil
		case c == '\\' && l.pos+1 < len(l.src):
			raw.WriteString(l.src[l.pos : l.pos+2])
			l.advance(2)
		case strings.HasPrefix(l.src[l.pos:], "${"):
			raw.WriteString("${…}")
			l.pos += 2
			if err := l.code(true); err != nil {
				return err
			}
		default:
			raw.WriteByte(c)
			l.advance(1)
		}
	}
	return fmt.Errorf("line %d: unterminated template literal", line)
}

// regexAllowed reports whether "/" here starts a regular expression: not
// after a value (identifier other than a keyword, number, string, template,
// regular expression, or a ")", "]" or "}" closing an expression).
func (l *tsLexer) regexAllowed() bool {
	if len(l.tokens) == 0 {
		return true
	}
	prev := l.tokens[len(l.tokens)-1]
	switch prev.kind {
	case identToken:
		return regexAfterKeyword[prev.text]
	case stringToken, templateToken, valueToken:
		return false
	}
	return prev.text != ")" && prev.text != "]" && prev.text != "}"
}

// regex skips a regular expression literal and its flags. One that reaches
// the end of the line is read as a division instead.
func (l *tsLexer) regex() {
	inClass := false
	for i := l.pos + 1; i < len(l.src); i++ {
		switch c := l.src[i]; {
		case c == '\n':
			l.tokens = append(l.tokens, tsToken{otherToken, "/", l.line})
			l.pos++
			return
		case c == '\\':
			i++
		case c == '[':
			inClass = true
		case c == ']':
			inClass = false
		case c == '/' && !inClass:
			i++
			for i < len(l.src) && isIdentPart(l.src[i]) {
				i++
			}
			l.tokens = append(l.tokens, tsToken{valueToken, l.src[l.pos:i], l.line})
			l.pos = i
			return
		}
	}
	l.tokens = append(l.tokens, tsToken{otherToken, "/", l.line})
	l.pos++
}

func isIdentStart(c byte) bool {
	return c == '_' || c == '$' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c >= 0x80
}

func isIdentPart(c byte) bool {
	return isIdentStart(c) || (c >= '0' && c <= '9')
}

// testName is one title read from a test file. plain is false where an ID
// does not belong: a template-literal title, or a describe title.
type testName struct {
	name    string
	line    int
	skipped bool
	plain   bool
}

// readTypeScriptTestNames finds every test(…), it(…) and describe(…) call
// chain outside comments and literals, taking the title from the first
// argument when it is a string or template literal; for a chain with .each,
// from the call that follows the table. Only test or it, plain or with
// .only, .skip, .fixme or .todo, and a plain-string title, can carry an ID.
// A file with a describe chain that is skipped or conditional has every test
// skipped.
func readTypeScriptTestNames(src string) ([]testName, error) {
	tokens, err := lexTypeScript(src)
	if err != nil {
		return nil, err
	}
	var names []testName
	fileSkipped := false
	for i, tok := range tokens {
		if tok.kind != identToken || (tok.text != "test" && tok.text != "it" && tok.text != "describe") {
			continue
		}
		if i > 0 && tokens[i-1].kind == otherToken && tokens[i-1].text == "." {
			continue // a method such as /re/.test(…)
		}
		chain := []string{tok.text}
		j := i + 1
		for j+1 < len(tokens) && isPunct(tokens[j], ".") && tokens[j+1].kind == identToken {
			chain = append(chain, tokens[j+1].text)
			j += 2
		}
		kind := classifyChain(chain)
		if kind.describe && kind.skipped {
			fileSkipped = true
		}
		title, ok := callTitle(tokens, j, slices.Contains(chain, "each"))
		if !ok {
			continue
		}
		names = append(names, testName{
			name:    title.text,
			line:    tok.line,
			skipped: kind.skipped,
			plain:   kind.plainTest && title.kind == stringToken,
		})
	}
	if fileSkipped {
		for i := range names {
			names[i].skipped = true
		}
	}
	return names, nil
}

type chainKind struct {
	describe  bool // a describe block, at any depth of modifiers
	plainTest bool // test or it, plain or with one of .only .skip .fixme .todo
	skipped   bool
}

// describeSkips are the modifiers that skip a describe block, or run it
// only on a condition (code-reviewer, 1.1c fix round 2).
var describeSkips = []string{"skip", "fixme", "todo", "skipIf", "runIf"}

func classifyChain(chain []string) chainKind {
	if slices.Contains(chain, "describe") {
		skipped := slices.ContainsFunc(chain, func(m string) bool { return slices.Contains(describeSkips, m) })
		return chainKind{describe: true, skipped: skipped}
	}
	rest := chain[1:]
	switch {
	case len(rest) == 0 || (len(rest) == 1 && rest[0] == "only"):
		return chainKind{plainTest: true}
	case len(rest) == 1 && (rest[0] == "skip" || rest[0] == "fixme" || rest[0] == "todo"):
		return chainKind{plainTest: true, skipped: true}
	}
	return chainKind{} // test.each, it.concurrent, test.step and the like
}

// callTitle returns the string or template literal that opens the call at
// tokens[open]. With each, a call whose first argument is a table is
// followed by a second call, which holds the title.
func callTitle(tokens []tsToken, open int, each bool) (tsToken, bool) {
	if open+1 >= len(tokens) || !isPunct(tokens[open], "(") {
		return tsToken{}, false
	}
	if first := tokens[open+1]; first.kind == stringToken || first.kind == templateToken {
		return first, true
	}
	if !each {
		return tsToken{}, false
	}
	depth := 0
	for k := open; k < len(tokens); k++ {
		switch {
		case isPunct(tokens[k], "("):
			depth++
		case isPunct(tokens[k], ")"):
			depth--
			if depth == 0 {
				return callTitle(tokens, k+1, false)
			}
		}
	}
	return tsToken{}, false
}

func isPunct(tok tsToken, text string) bool {
	return tok.kind == otherToken && tok.text == text
}
