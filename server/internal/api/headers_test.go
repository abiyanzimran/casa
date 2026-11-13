package api_test

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"testing"

	"connectrpc.com/connect/v2"

	"github.com/abiyanzimran/casa/server/internal/api"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
)

// fixtureCSP is the policy of plan decision 5 for the fixture bundle, with
// style-src-attr from the SEC-166 note of FS-1 task 1.5 (the SvelteKit route
// announcer's style attribute, by its hash).
func fixtureCSP() string {
	sum := sha256.Sum256([]byte(fixtureScript))
	return "default-src 'none'; script-src 'self' 'sha256-" + base64.StdEncoding.EncodeToString(sum[:]) + "'; " +
		"style-src 'self'; style-src-attr 'unsafe-hashes' 'sha256-S8qMpvofolR8Mpjy4kQvEm7m1q8clzU4dfDH0AmvZjo='; " +
		"img-src 'self' blob:; font-src 'self'; connect-src 'self'; worker-src 'self'; " +
		"manifest-src 'self'; base-uri 'none'; form-action 'self'; frame-ancestors 'none'; object-src 'none'; " +
		"require-trusted-types-for 'script'; trusted-types svelte-trusted-html casa-sw"
}

// securityHeaders are the headers of plan decision 5 other than Cache-Control.
func securityHeaders(csp string) map[string]string {
	return map[string]string{
		"Content-Security-Policy":      csp,
		"Strict-Transport-Security":    "max-age=31536000; includeSubDomains",
		"X-Content-Type-Options":       "nosniff",
		"Referrer-Policy":              "no-referrer",
		"Cross-Origin-Opener-Policy":   "same-origin",
		"Cross-Origin-Resource-Policy": "same-origin",
		"Permissions-Policy":           "camera=(), microphone=(), geolocation=(), payment=(), usb=()",
	}
}

// wantSecurityHeaders checks the header set of plan decision 5 and that no
// CORS header is present (SEC-160 (b)).
func wantSecurityHeaders(t *testing.T, header http.Header, csp, cacheControl, request string) {
	t.Helper()
	for name, want := range securityHeaders(csp) {
		if got := header.Values(name); len(got) != 1 || got[0] != want {
			t.Errorf("%s: %s = %q, want exactly %q", request, name, got, want)
		}
	}
	if got := header.Get("Cache-Control"); got != cacheControl {
		t.Errorf("%s: Cache-Control = %q, want %q", request, got, cacheControl)
	}
	for name := range header {
		if strings.HasPrefix(name, "Access-Control-") {
			t.Errorf("%s: has CORS header %s", request, name)
		}
	}
}

// newFixtureMux serves the fixture bundle and one Connect procedure that
// always fails, through NewMux.
func newFixtureMux(t *testing.T) http.Handler {
	t.Helper()
	server := api.NewServer(api.Version{}, testOrigin, api.Sessions{}, nil)
	server.Register(connect.Method{
		Spec: connect.Spec{Procedure: casav1connect.IdentityServiceSignInProcedure},
		Handler: func(context.Context, connect.Spec, connect.ServerStream) error {
			return errors.New("not reached in these tests")
		},
	})
	return api.NewMux(server, newFixtureBundle(t))
}

func TestSecurityHeaders(t *testing.T) {
	mux := newFixtureMux(t)
	connectJSON := http.Header{"Content-Type": {"application/json"}}

	t.Run("SEC-166 every kind of answer carries the security headers", func(t *testing.T) {
		// SEC-166, the CI part: the header set of every kind of answer of the
		// origin is compared with plan decision 5 (SEC-160 (c); ADR-030 item
		// 7; Cache-Control per plan decisions 3 and 5). The rest of SEC-166 is
		// proved elsewhere: the injected-script check by the browser test of
		// task 2, the scan of the production headers at release step 7.
		for _, tc := range []struct {
			method, target string
			header         http.Header
			status         int
			cacheControl   string
		}{
			{http.MethodGet, "/", nil, http.StatusOK, "no-cache"},
			{http.MethodGet, "/money", nil, http.StatusOK, "no-cache"},
			{http.MethodGet, "/_app/immutable/entry/start.abc.js", nil, http.StatusOK, "public, max-age=31536000, immutable"},
			{http.MethodGet, "/service-worker.js", nil, http.StatusOK, "no-cache"},
			{http.MethodGet, "/manifest.webmanifest", nil, http.StatusOK, "no-cache"},
			{http.MethodGet, "/healthz", nil, http.StatusOK, "no-store"},
			{http.MethodPost, "/casa.v1.IdentityService/Unknown", connectJSON, http.StatusNotImplemented, "no-store"},
			{http.MethodGet, "/_app/immutable/missing.js", nil, http.StatusNotFound, "no-store"},
			{http.MethodPut, "/", nil, http.StatusMethodNotAllowed, "no-store"},
			// Range and If-Match are not served (plan decision 3), so
			// ServeContent never answers 416 or 412, whose error path drops
			// Cache-Control or keeps the immutable one.
			{http.MethodGet, "/", http.Header{"Range": {"bytes=999999-"}}, http.StatusOK, "no-cache"},
			{http.MethodGet, "/_app/immutable/entry/start.abc.js", http.Header{"If-Match": {`"nope"`}}, http.StatusOK, "public, max-age=31536000, immutable"},
		} {
			request := tc.method + " " + tc.target
			rec := serve(mux, tc.method, tc.target, tc.header)
			if rec.Code != tc.status {
				t.Errorf("%s = %d %q, want %d", request, rec.Code, rec.Body.String(), tc.status)
			}
			wantSecurityHeaders(t, rec.Header(), fixtureCSP(), tc.cacheControl, request)
		}
	})
	t.Run("the CSP allows the bundle's inline script by its hash, no 'unsafe-inline' or 'unsafe-eval'", func(t *testing.T) {
		// Plan decision 4; ADR-030 item 7: no 'unsafe-inline', no 'unsafe-eval'.
		// The one 'unsafe-hashes' (style-src-attr) is pinned by fixtureCSP.
		csp := serve(mux, http.MethodGet, "/", nil).Header().Get("Content-Security-Policy")
		sum := sha256.Sum256([]byte(fixtureScript))
		want := "script-src 'self' 'sha256-" + base64.StdEncoding.EncodeToString(sum[:]) + "';"
		if !strings.Contains(csp, want) || strings.Contains(csp, "'unsafe-inline'") || strings.Contains(csp, "'unsafe-eval'") {
			t.Fatalf("CSP = %q, want %q and no 'unsafe-inline' or 'unsafe-eval'", csp, want)
		}
	})
	t.Run("SEC-160 (b) a cross-origin preflight gets no CORS header", func(t *testing.T) {
		// SEC-160 (b): the API is same-origin only; no answer opens it to
		// another origin.
		rec := serve(mux, http.MethodOptions, casav1connect.IdentityServiceSignInProcedure, http.Header{
			"Origin":                        {"https://other.example"},
			"Access-Control-Request-Method": {http.MethodPost},
		})
		wantSecurityHeaders(t, rec.Header(), fixtureCSP(), "no-store", "OPTIONS SignIn")
	})
}
