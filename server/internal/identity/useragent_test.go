package identity

import (
	"testing"
	"unicode/utf8"
)

func TestSummarizeUserAgent(t *testing.T) {
	// Plan 1.5 decision 10; SEC-161 (e) note of 13 November 2026: a fixed
	// list, no version, at most 64 characters.
	for _, c := range []struct{ name, ua, want string }{
		{"Safari iPhone",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 18_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Mobile/15E148 Safari/604.1",
			"Safari · iOS"},
		{"Safari iPad",
			"Mozilla/5.0 (iPad; CPU OS 18_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Mobile/15E148 Safari/604.1",
			"Safari · iPadOS"},
		{"Safari iPad asking for the desktop site",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Safari/605.1.15",
			"Safari · macOS/iPadOS"},
		{"Firefox iOS",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 18_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/132.0 Mobile/15E148 Safari/605.1.15",
			"Firefox · iOS"},
		{"Chrome Android on a Galaxy A55",
			"Mozilla/5.0 (Linux; Android 14; SM-A556E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Mobile Safari/537.36",
			"Chrome · Android"},
		{"Samsung Internet",
			"Mozilla/5.0 (Linux; Android 14; SAMSUNG SM-A556E) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/26.0 Chrome/122.0.0.0 Mobile Safari/537.36",
			"Samsung Internet · Android"},
		{"Edge Windows",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
			"Edge · Windows"},
		{"Chrome iOS",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 18_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/131.0.6778.73 Mobile/15E148 Safari/604.1",
			"Chrome · iOS"},
		{"Firefox Linux",
			"Mozilla/5.0 (X11; Linux x86_64; rv:132.0) Gecko/20100101 Firefox/132.0",
			"Firefox · Linux"},
		{"empty", "", "Browser · Other"},
		{"markup and a control character", "<script>\x00", "Browser · Other"},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := summarizeUserAgent(c.ua)
			if got != c.want {
				t.Fatalf("summarizeUserAgent(%q) = %q, want %q", c.ua, got, c.want)
			}
			if utf8.RuneCountInString(got) > 64 {
				t.Fatalf("%q is longer than 64 characters", got)
			}
		})
	}
}
