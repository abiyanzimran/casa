package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strings"
	"time"
)

// ErrNoBundle means the binary was built without the web client (plan
// decision 2): casa serve then serves the API only.
var ErrNoBundle = errors.New("no index.html in the web bundle")

// contentTypes is the fixed Content-Type of each file type the bundle may
// hold. Fixed in code, not mime.TypeByExtension, which reads the machine's
// /etc/mime.types and so differs on the VPS (plan decision 3; SEC-166).
var contentTypes = map[string]string{
	".html":        "text/html; charset=utf-8",
	".js":          "text/javascript; charset=utf-8",
	".css":         "text/css; charset=utf-8",
	".json":        "application/json",
	".webmanifest": "application/manifest+json",
	".png":         "image/png",
	".svg":         "image/svg+xml",
	".woff2":       "font/woff2",
	".ico":         "image/x-icon",
	".txt":         "text/plain; charset=utf-8",
}

// immutablePrefix holds the files whose names carry a content hash (ADR-030
// item 7; plan decision 3).
const immutablePrefix = "_app/immutable/"

type bundleFile struct {
	content     []byte
	contentType string
	etag        string
}

// Bundle serves the web client's static files (ADR-030 items 3, 7, 9;
// SEC-166, SEC-169 (b); plan decision 3). It is read once by NewBundle and
// never changes after.
type Bundle struct {
	files map[string]bundleFile
	// scriptHashes are the CSP sources of index.html's inline scripts
	// (plan decision 4).
	scriptHashes []string
}

// NewBundle reads the bundle rooted at fsys. It fails at start, not when a
// file is served: on a file type outside contentTypes, or on an inline script
// in index.html the CSP could not allow by hash. Without index.html it
// returns ErrNoBundle. Files and directories whose name starts with "." are
// left out and never served (plan decision 1).
func NewBundle(fsys fs.FS) (*Bundle, error) {
	files := map[string]bundleFile{}
	err := fs.WalkDir(fsys, ".", func(name string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if name != "." && strings.HasPrefix(entry.Name(), ".") {
			if entry.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		contentType, ok := contentTypes[path.Ext(name)]
		if !ok {
			return fmt.Errorf("%s: file type outside the fixed map", name)
		}
		content, err := fs.ReadFile(fsys, name)
		if err != nil {
			return err
		}
		sum := sha256.Sum256(content)
		// A strong ETag: the first 16 hex digits of the SHA-256 (plan decision 3).
		files[name] = bundleFile{content: content, contentType: contentType, etag: `"` + hex.EncodeToString(sum[:8]) + `"`}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("read the web bundle: %w", err)
	}
	index, ok := files["index.html"]
	if !ok {
		return nil, ErrNoBundle
	}
	scriptHashes, err := inlineScriptHashes(index.content)
	if err != nil {
		return nil, fmt.Errorf("read the web bundle: index.html: %w", err)
	}
	return &Bundle{files: files, scriptHashes: scriptHashes}, nil
}

// inlineScriptHashes returns the CSP source 'sha256-<base64>' of each
// <script> without attributes in page, over the exact bytes between <script>
// and </script> (plan decision 4; ADR-030 item 7: no 'unsafe-inline'). Any
// other <script must load a file with src=; one that does not would be
// blocked by the CSP, so it is an error here.
func inlineScriptHashes(page []byte) ([]string, error) {
	var hashes []string
	rest := string(page)
	for {
		start := strings.Index(rest, "<script")
		if start < 0 {
			return hashes, nil
		}
		rest = rest[start:]
		if body, ok := strings.CutPrefix(rest, "<script>"); ok {
			script, after, closed := strings.Cut(body, "</script>")
			if !closed {
				return nil, errors.New("an inline script has no </script>")
			}
			sum := sha256.Sum256([]byte(script))
			hashes = append(hashes, "'sha256-"+base64.StdEncoding.EncodeToString(sum[:])+"'")
			rest = after
			continue
		}
		if tag, _, _ := strings.Cut(rest, ">"); !strings.Contains(tag, " src=") {
			return nil, errors.New("a script with attributes and no src= has no CSP hash")
		}
		rest = rest[len("<script"):]
	}
}

// ServeHTTP answers GET and HEAD with a bundle file, or with index.html for
// an app route (the SPA fallback), else 404 in plain text: never a directory
// listing, never the app page for a missing asset (plan decision 3).
func (b *Bundle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Clean against "/" first, so ".." never climbs out of the bundle.
	name := strings.TrimPrefix(path.Clean("/"+r.URL.Path), "/")
	if name == "" {
		name = "index.html"
	}
	file, ok := b.files[name]
	if !ok {
		if !isAppRoute(name) {
			http.NotFound(w, r)
			return
		}
		name, file = "index.html", b.files["index.html"]
	}
	header := w.Header()
	header.Set("Content-Type", file.contentType)
	header.Set("ETag", file.etag)
	// Replaces the no-store of WithSecurityHeaders. Everything outside
	// _app/immutable/, index.html and service-worker.js included, is
	// revalidated on every load (SEC-169 (b); lead ruling Q8).
	if strings.HasPrefix(name, immutablePrefix) {
		header.Set("Cache-Control", "public, max-age=31536000, immutable")
	} else {
		header.Set("Cache-Control", "no-cache")
	}
	http.ServeContent(w, r, name, time.Time{}, bytes.NewReader(file.content))
}

// isAppRoute reports whether a missing name is a client route, answered with
// index.html: its last segment has no dot (not a file name) and it is not
// under _app/, which holds only build files (plan decision 3).
func isAppRoute(name string) bool {
	return !strings.Contains(path.Base(name), ".") && name != "_app" && !strings.HasPrefix(name, "_app/")
}
