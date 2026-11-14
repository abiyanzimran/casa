package identity

import (
	"errors"
	"testing"
	"time"
)

func TestCompleteSignUpEmail(t *testing.T) {
	t.Run("an email mismatch runs no hash and keeps the link", func(t *testing.T) {
		// 10.4 SEC-158 note "halaman verifikasi" item 2: refused before the hash.
		now := time.Now()
		s, _, verification := startedSignUp(t, &now)
		before := s.hasher.count.Load()
		_, err := s.CompleteSignUp(t.Context(), testClient, verification, "other@example.test", "a long enough passphrase", noInvitation)
		if !errors.Is(err, ErrEmailMismatch) {
			t.Fatalf("err = %v, want ErrEmailMismatch", err)
		}
		if n := s.hasher.count.Load() - before; n != 0 {
			t.Fatalf("%d hashes, want none", n)
		}
		if _, err := s.CompleteSignUp(t.Context(), testClient, verification, "owner@example.test", "a long enough passphrase", noInvitation); err != nil {
			t.Fatalf("the right email: %v", err)
		}
	})
}
