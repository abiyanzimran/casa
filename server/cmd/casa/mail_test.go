package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/abiyanzimran/casa/server/internal/identity"
)

func TestWriteMail(t *testing.T) {
	// 14 §7.8 item 2: the test server's fake mail sender.
	dir := t.TempDir()
	send := writeMail(dir)
	message := identity.Message{To: "owner@example.test", Kind: identity.MessageRecovery, Link: "http://127.0.0.1:8080/recover#token"}
	if err := send(t.Context(), message); err != nil {
		t.Fatal(err)
	}
	files, err := os.ReadDir(dir)
	if err != nil || len(files) != 1 {
		t.Fatalf("%d files (%v), want 1", len(files), err)
	}
	name := files[0].Name()
	if !regexp.MustCompile(`^\d{8}T\d{6}\.\d{9}Z-recovery\.txt$`).MatchString(name) {
		t.Fatalf("file name %q, want <UTC time>-recovery.txt", name)
	}
	info, err := files[0].Info()
	if err != nil || info.Mode().Perm() != 0o600 {
		t.Fatalf("mode %v (%v), want 0600", info.Mode().Perm(), err)
	}
	body, err := fs.ReadFile(os.DirFS(dir), name)
	if want := "To: owner@example.test\nKind: recovery\nLink: http://127.0.0.1:8080/recover#token\n"; err != nil || string(body) != want {
		t.Fatalf("body %q (%v), want %q", body, err, want)
	}
	message.Kind = identity.MessageVerification
	if err := send(t.Context(), message); err != nil {
		t.Fatal(err)
	}
	if matches, _ := filepath.Glob(filepath.Join(dir, "*-verification.txt")); len(matches) != 1 {
		t.Fatalf("%d verification files, want 1", len(matches))
	}
}
