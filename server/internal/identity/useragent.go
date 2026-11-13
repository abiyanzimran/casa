package identity

import "strings"

// summarizeUserAgent names a web session's browser and system for the
// session list as "<browser> · <system>", from a fixed list and without
// versions (plan 1.5 decision 10; SEC-161 (e) note of 13 November 2026). The
// result is only ever a label from the list, so no part of the User-Agent is
// stored or logged. The browser order matters: Edge and Samsung Internet also
// name Chrome, and Chrome also names Safari. "Browser" and "Other" are
// placeholder labels until ux-designer words them (08).
func summarizeUserAgent(ua string) string {
	has := func(s string) bool { return strings.Contains(ua, s) }
	browser := "Browser"
	switch {
	case has("Edg/"):
		browser = "Edge"
	case has("SamsungBrowser/"):
		browser = "Samsung Internet"
	case has("Firefox/"), has("FxiOS/"):
		browser = "Firefox"
	case has("Chrome/"), has("CriOS/"):
		browser = "Chrome"
	case has("Safari/") && has("Version/"):
		browser = "Safari"
	}
	system := "Other"
	switch {
	case has("iPhone"):
		system = "iOS"
	case has("iPad"):
		system = "iPadOS"
	case has("Android"):
		system = "Android"
	case has("Windows"):
		system = "Windows"
	case has("Macintosh"):
		// Safari on an iPad asks for the desktop site by default.
		system = "macOS/iPadOS"
	case has("Linux"):
		system = "Linux"
	}
	return browser + " · " + system
}
