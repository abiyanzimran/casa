package identity

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	token, hash := newToken()
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil || len(token) != 43 || len(raw) != 32 {
		t.Fatalf("token %d characters decoding to %d bytes (%v), want 43 and 32", len(token), len(raw), err)
	}
	if sum := sha256.Sum256(raw); !bytes.Equal(hash, sum[:]) {
		t.Fatal("hash is not the SHA-256 of the raw token")
	}
	if got, ok := hashToken(token); !ok || !bytes.Equal(got, hash) {
		t.Fatalf("hashToken(token) = %x, %t; want %x, true", got, ok, hash)
	}
	if other, _ := newToken(); other == token {
		t.Fatal("two tokens are equal")
	}
	// The last character of a 43-character token carries 2 unused bits; a
	// value that sets them is not the canonical form of any token.
	nonCanonical := token[:42] + "B"
	if token[42] == 'A' || token[42] == 'B' {
		nonCanonical = token[:42] + "D"
	}
	for _, bad := range []string{"", "abc", token + "A", strings.Replace(token, token[:1], "+", 1), nonCanonical} {
		if _, ok := hashToken(bad); ok {
			t.Errorf("hashToken(%q) ok, want rejected", bad)
		}
	}
}
