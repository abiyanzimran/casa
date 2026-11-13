package api_test

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/abiyanzimran/casa/server/internal/api"
)

// fixtureScript is the inline script of the fixture app page, the bytes
// between <script> and </script> exactly (plan decision 4).
const fixtureScript = "\n\t\t\twindow.started = true;\n\t\t"

const fixtureIndex = "<!doctype html><html><head>" +
	`<script src="/_app/immutable/entry/start.abc.js" type="module"></script>` +
	"</head><body><div><script>" + fixtureScript + "</script></div></body></html>"

// fixtureBundle is a small web bundle with each kind of file the server
// treats differently (plan decision 3).
func fixtureBundle() fstest.MapFS {
	return fstest.MapFS{
		"index.html":                         {Data: []byte(fixtureIndex)},
		"_app/immutable/entry/start.abc.js":  {Data: []byte("export const start = 1;\n")},
		"service-worker.js":                  {Data: []byte("self.addEventListener('install', () => {});\n")},
		"manifest.webmanifest":               {Data: []byte(`{"name":"Casa"}`)},
		"fonts/x.woff2":                      {Data: []byte("wOF2 font bytes")},
		".gitkeep":                           {Data: []byte("")},
		"a/.hidden":                          {Data: []byte("hidden bytes")},
		".well-known/security.txt.untracked": {Data: []byte("dot directory bytes")},
	}
}

func newFixtureBundle(t *testing.T) *api.Bundle {
	t.Helper()
	bundle, err := api.NewBundle(fixtureBundle())
	if err != nil {
		t.Fatalf("NewBundle: %v", err)
	}
	return bundle
}

func serve(handler http.Handler, method, target string, header http.Header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, target, nil)
	for name, values := range header {
		req.Header[name] = values
	}
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

// wantFile checks a 200 answer with the given body, type and caching.
func wantFile(t *testing.T, rec *httptest.ResponseRecorder, target, body, contentType, cacheControl string) {
	t.Helper()
	if rec.Code != http.StatusOK || rec.Body.String() != body {
		t.Fatalf("GET %s = %d %q, want 200 %q", target, rec.Code, rec.Body.String(), body)
	}
	if got := rec.Header().Get("Content-Type"); got != contentType {
		t.Errorf("GET %s Content-Type = %q, want %q", target, got, contentType)
	}
	if got := rec.Header().Get("Cache-Control"); got != cacheControl {
		t.Errorf("GET %s Cache-Control = %q, want %q", target, got, cacheControl)
	}
}

// wantNotFound checks a plain-text 404: never the app page, never a listing.
func wantNotFound(t *testing.T, rec *httptest.ResponseRecorder, target string) {
	t.Helper()
	if rec.Code != http.StatusNotFound || !strings.HasPrefix(rec.Header().Get("Content-Type"), "text/plain") {
		t.Fatalf("GET %s = %d %q (%s), want 404 text/plain", target, rec.Code, rec.Body.String(), rec.Header().Get("Content-Type"))
	}
}

func TestBundle(t *testing.T) {
	bundle := newFixtureBundle(t)
	files := fixtureBundle()

	t.Run("serves the app page at / and for a route path", func(t *testing.T) {
		for _, target := range []string{"/", "/money/budget"} {
			wantFile(t, serve(bundle, http.MethodGet, target, nil), target, fixtureIndex, "text/html; charset=utf-8", "no-cache")
		}
	})
	t.Run("serves hashed files as immutable", func(t *testing.T) {
		// ADR-030 item 7: hashed file names never change content.
		target := "/_app/immutable/entry/start.abc.js"
		wantFile(t, serve(bundle, http.MethodGet, target, nil), target,
			string(files["_app/immutable/entry/start.abc.js"].Data), "text/javascript; charset=utf-8", "public, max-age=31536000, immutable")
	})
	t.Run("serves the service worker and other static files with no-cache", func(t *testing.T) {
		// SEC-169 (b): the browser revalidates the service worker and the app
		// page on every load, so a fixed release reaches every device.
		wantFile(t, serve(bundle, http.MethodGet, "/service-worker.js", nil), "/service-worker.js",
			string(files["service-worker.js"].Data), "text/javascript; charset=utf-8", "no-cache")
		wantFile(t, serve(bundle, http.MethodGet, "/fonts/x.woff2", nil), "/fonts/x.woff2",
			string(files["fonts/x.woff2"].Data), "font/woff2", "no-cache")
	})
	t.Run("never answers a missing asset with the app page", func(t *testing.T) {
		for _, target := range []string{"/_app/immutable/missing.js", "/favicon.ico", "/_app/", "/_app"} {
			wantNotFound(t, serve(bundle, http.MethodGet, target, nil), target)
		}
	})
	t.Run("never serves outside the bundle or a dot file", func(t *testing.T) {
		for _, target := range []string{"/../../go.mod", "/%2e%2e/x", "/.gitkeep", "/a/.hidden", "/.well-known/security.txt.untracked", "/a/", "/fonts/"} {
			rec := serve(bundle, http.MethodGet, target, nil)
			if rec.Code == http.StatusOK && rec.Body.String() == fixtureIndex {
				continue // the app page for a route path
			}
			wantNotFound(t, rec, target)
			for _, file := range files {
				if len(file.Data) > 0 && strings.Contains(rec.Body.String(), string(file.Data)) {
					t.Fatalf("GET %s served bundle bytes %q", target, rec.Body.String())
				}
			}
		}
	})
	t.Run("answers If-None-Match with 304", func(t *testing.T) {
		// Plan decision 3: a strong ETag of the first 16 hex digits of the
		// content's SHA-256.
		sum := sha256.Sum256([]byte(fixtureIndex))
		etag := `"` + hex.EncodeToString(sum[:])[:16] + `"`
		rec := serve(bundle, http.MethodGet, "/", nil)
		if got := rec.Header().Get("ETag"); got != etag {
			t.Fatalf("ETag = %q, want %q", got, etag)
		}
		rec = serve(bundle, http.MethodGet, "/money", http.Header{"If-None-Match": {etag}})
		if rec.Code != http.StatusNotModified || rec.Body.Len() != 0 {
			t.Fatalf("GET with If-None-Match = %d %q, want 304 and no body", rec.Code, rec.Body.String())
		}
	})
	t.Run("serves the manifest as application/manifest+json", func(t *testing.T) {
		wantFile(t, serve(bundle, http.MethodGet, "/manifest.webmanifest", nil), "/manifest.webmanifest",
			string(files["manifest.webmanifest"].Data), "application/manifest+json", "no-cache")
	})
	t.Run("answers HEAD without a body and other methods with 405", func(t *testing.T) {
		rec := serve(bundle, http.MethodHead, "/", nil)
		if rec.Code != http.StatusOK || rec.Body.Len() != 0 || rec.Header().Get("Content-Type") != "text/html; charset=utf-8" {
			t.Fatalf("HEAD / = %d %q %q, want 200, text/html and no body", rec.Code, rec.Body.String(), rec.Header().Get("Content-Type"))
		}
		for _, method := range []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions} {
			rec := serve(bundle, method, "/", nil)
			if rec.Code != http.StatusMethodNotAllowed || rec.Header().Get("Allow") != "GET, HEAD" || strings.Contains(rec.Body.String(), fixtureScript) {
				t.Fatalf("%s / = %d %q, Allow %q; want 405 and Allow \"GET, HEAD\"", method, rec.Code, rec.Body.String(), rec.Header().Get("Allow"))
			}
		}
	})
	t.Run("rejects a bundle without index.html", func(t *testing.T) {
		// Plan decision 2: casa serve then serves the API only.
		_, err := api.NewBundle(fstest.MapFS{".gitkeep": {Data: []byte("")}})
		if !errors.Is(err, api.ErrNoBundle) {
			t.Fatalf("NewBundle(empty) = %v, want ErrNoBundle", err)
		}
	})
	t.Run("rejects a bundle with a file type outside the fixed map", func(t *testing.T) {
		withWasm := fixtureBundle()
		withWasm["x.wasm"] = &fstest.MapFile{Data: []byte("\x00asm")}
		if _, err := api.NewBundle(withWasm); err == nil || errors.Is(err, api.ErrNoBundle) {
			t.Fatalf("NewBundle(x.wasm) = %v, want a file type error", err)
		}
	})
	t.Run("rejects an index.html with an inline script that has attributes", func(t *testing.T) {
		// Plan decision 4: the CSP would have no hash for it.
		withModule := fixtureBundle()
		withModule["index.html"] = &fstest.MapFile{Data: []byte(`<html><script type="module">start()</script></html>`)}
		if _, err := api.NewBundle(withModule); err == nil || errors.Is(err, api.ErrNoBundle) {
			t.Fatalf("NewBundle(script with attributes) = %v, want an error", err)
		}
	})
}
