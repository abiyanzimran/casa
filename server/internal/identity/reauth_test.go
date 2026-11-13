package identity

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestChangePasswordCost(t *testing.T) {
	// Built, so no secret scanner sees a literal.
	password, next := strings.Repeat("reauth ", 3), strings.Repeat("changed ", 3)
	t.Run("SEC-12 change password without a reauth grant runs no hash", func(t *testing.T) {
		// Lead ruling on PR B: without an open grant a session cannot spend
		// the hash slots SignIn needs (10.1 §8.1 item 3; SEC-13).
		now := time.Now()
		s, _ := signedUp(t, &now, password)
		tokens, err := s.SignIn(t.Context(), testClient, "owner@example.test", password)
		if err != nil {
			t.Fatal(err)
		}
		session, err := s.Authenticate(t.Context(), tokens.AccessToken)
		if err != nil {
			t.Fatal(err)
		}
		before := s.hasher.count.Load()
		for range 3 {
			if _, err := s.HashNewPassword(t.Context(), session, next); !errors.Is(err, ErrReauthRequired) {
				t.Fatalf("err = %v, want ErrReauthRequired", err)
			}
		}
		if n := s.hasher.count.Load() - before; n != 0 {
			t.Fatalf("%d hashes without a grant, want 0", n)
		}
		if _, err := s.Reauthenticate(t.Context(), testClient, session, password, ReauthChangePassword); err != nil {
			t.Fatal(err)
		}
		now = now.Add(reauthLifetime) // expired: still no hash
		if _, err := s.HashNewPassword(t.Context(), session, next); !errors.Is(err, ErrReauthRequired) {
			t.Fatalf("expired grant: err = %v, want ErrReauthRequired", err)
		}
		if _, err := s.Reauthenticate(t.Context(), testClient, session, password, ReauthChangePassword); err != nil {
			t.Fatal(err)
		}
		before = s.hasher.count.Load()
		if _, err := s.HashNewPassword(t.Context(), session, next); err != nil {
			t.Fatal(err)
		}
		if n := s.hasher.count.Load() - before; n != 1 {
			t.Fatalf("%d hashes with a grant, want 1", n)
		}
	})
}
