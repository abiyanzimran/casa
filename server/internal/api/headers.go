package api

import (
	"net/http"
	"strings"
)

// announcerStyleHash allows exactly one style attribute: the one on the route
// announcer (<div id="svelte-announcer" style="position: absolute; ...">) in
// SvelteKit's runtime root.svelte, which Casa cannot change (SEC-166 note of
// FS-1 task 1.5). It is the SHA-256 of that attribute's value. Casa code never
// gets a style hash; its styles come from stylesheets ('self').
const announcerStyleHash = "'sha256-S8qMpvofolR8Mpjy4kQvEm7m1q8clzU4dfDH0AmvZjo='"

// WithSecurityHeaders sets the security headers of every answer of the origin
// before next writes (SEC-166, SEC-160 (c); ADR-030 item 7; plan decision 5):
// API answers, errors, 404s and static files alike. Cache-Control defaults to
// no-store; Bundle replaces it for the files it serves. No answer carries an
// Access-Control-* header: the API is same-origin only (SEC-160 (b)). bundle
// may be nil; the CSP then allows no inline script.
func WithSecurityHeaders(next http.Handler, bundle *Bundle) http.Handler {
	scriptSources := []string{"'self'"}
	if bundle != nil {
		scriptSources = append(scriptSources, bundle.scriptHashes...)
	}
	// svelte-trusted-html stays until the 1.11 prototype decides whether
	// Svelte needs it (ADR-030 "Yang menunggu").
	csp := "default-src 'none'; script-src " + strings.Join(scriptSources, " ") +
		"; style-src 'self'; style-src-attr 'unsafe-hashes' " + announcerStyleHash +
		"; img-src 'self' blob:; font-src 'self'; connect-src 'self'; worker-src 'self'" +
		"; manifest-src 'self'; base-uri 'none'; form-action 'self'; frame-ancestors 'none'; object-src 'none'" +
		"; require-trusted-types-for 'script'; trusted-types svelte-trusted-html casa-sw"
	headers := [][2]string{
		{"Content-Security-Policy", csp},
		{"Strict-Transport-Security", "max-age=31536000; includeSubDomains"},
		{"X-Content-Type-Options", "nosniff"},
		{"Referrer-Policy", "no-referrer"},
		{"Cross-Origin-Opener-Policy", "same-origin"},
		{"Cross-Origin-Resource-Policy", "same-origin"},
		{"Permissions-Policy", "camera=(), microphone=(), geolocation=(), payment=(), usb=()"},
		{"Cache-Control", "no-store"},
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		for _, h := range headers {
			header.Set(h[0], h[1])
		}
		next.ServeHTTP(w, r)
	})
}
