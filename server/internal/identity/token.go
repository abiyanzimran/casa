package identity

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// tokenBytes is the size of every token: access, refresh, setup,
// verification, recovery and invitation (SEC-15, SEC-24, SEC-28, SEC-158;
// 09.4 §2: 256 bits).
const tokenBytes = 32

var tokenEncoding = base64.RawURLEncoding.Strict()

// NewToken returns a random token as 43 characters of unpadded base64url and
// the SHA-256 of its raw bytes, the only form the database keeps (SEC-15).
// identity is the one maker of tokens: the API layer uses it for invitation
// links too (SEC-28; 09.2 §1 item 9).
func NewToken() (token string, hash []byte) {
	raw := randomBytes(tokenBytes)
	sum := sha256.Sum256(raw)
	return tokenEncoding.EncodeToString(raw), sum[:]
}

// HashToken returns the stored form of token; ok is false unless token is
// the canonical encoding of 32 bytes. A malformed token is treated like an
// unknown one (SEC-30).
func HashToken(token string) (hash []byte, ok bool) {
	if len(token) != tokenEncoding.EncodedLen(tokenBytes) {
		return nil, false
	}
	raw, err := tokenEncoding.DecodeString(token)
	if err != nil || len(raw) != tokenBytes {
		return nil, false
	}
	sum := sha256.Sum256(raw)
	return sum[:], true
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b) // never returns an error (crypto/rand, Go 1.24 and later)
	return b
}
