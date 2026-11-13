// Package webbundle embeds the web client's build in the casa binary
// (ADR-030 item 3; 09.4 §1 "Bundel web"). make bundle-copy fills dist/ from
// web/build; in a build without it, dist/ holds only .gitkeep and api.NewBundle
// answers ErrNoBundle.
package webbundle

import "embed"

// FS holds dist/ and everything in it, dot files included; the server never
// serves a dot file (api.NewBundle).
//
//go:embed all:dist
var FS embed.FS
