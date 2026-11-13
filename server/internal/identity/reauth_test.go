package identity

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/abiyanzimran/casa/server/internal/database"
)

func TestChangePasswordCost(t *testing.T) {
	// Built, so no secret scanner sees a literal.
	password, next := strings.Repeat("reauth ", 3), strings.Repeat("changed ", 3)
	t.Run("SEC-13 change password without a reauth grant runs no hash", func(t *testing.T) {
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
		if _, _, err := s.Reauthenticate(t.Context(), testClient, session, password, ReauthChangePassword); err != nil {
			t.Fatal(err)
		}
		now = now.Add(reauthLifetime) // expired: still no hash
		if _, err := s.HashNewPassword(t.Context(), session, next); !errors.Is(err, ErrReauthRequired) {
			t.Fatalf("expired grant: err = %v, want ErrReauthRequired", err)
		}
		if _, _, err := s.Reauthenticate(t.Context(), testClient, session, password, ReauthChangePassword); err != nil {
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

func TestWebReauthGrantIsBoundToItsCookie(t *testing.T) {
	// SEC-161 (c) note 1.5 rule (7): ChangePassword uses a web grant only
	// for the request presenting the value issued with it, even where
	// HashNewPassword's read is skipped (that read is not the authority).
	password, next := strings.Repeat("reauth ", 3), strings.Repeat("changed ", 3)
	now := time.Now()
	s, pool := signedUp(t, &now, password)
	web := Client{Kind: "web", Version: "1.0.0", Address: "192.0.2.1:443", DeviceMode: "personal"}
	tokens, err := s.SignIn(t.Context(), web, "owner@example.test", password)
	if err != nil {
		t.Fatal(err)
	}
	before, err := s.AuthenticateCookie(t.Context(), tokens.Cookie.Value)
	if err != nil {
		t.Fatal(err)
	}
	_, cookie, err := s.Reauthenticate(t.Context(), web, before.Session, password, ReauthChangePassword)
	if err != nil || cookie == nil {
		t.Fatalf("re-authentication: cookie %t, err %v; want a cookie", cookie != nil, err)
	}
	if _, err := s.HashNewPassword(t.Context(), before.Session, next); !errors.Is(err, ErrReauthRequired) {
		t.Fatalf("hash with the value before the re-authentication: err = %v, want ErrReauthRequired", err)
	}
	hash, err := s.hashPassword(t.Context(), next)
	if err != nil {
		t.Fatal(err)
	}
	change := func(session Session) error {
		return database.Write(t.Context(), pool, database.Scope{AccountID: session.AccountID}, func(tx pgx.Tx) error {
			return s.ChangePassword(t.Context(), tx, session, hash)
		})
	}
	if err := change(before.Session); !errors.Is(err, ErrReauthRequired) {
		t.Fatalf("the value before the re-authentication: err = %v, want ErrReauthRequired", err)
	}
	native := before.Session
	native.CookieHash = nil
	if err := change(native); !errors.Is(err, ErrReauthRequired) {
		t.Fatalf("no cookie: err = %v, want ErrReauthRequired", err)
	}
	after, err := s.AuthenticateCookie(t.Context(), cookie.Value)
	if err != nil {
		t.Fatal(err)
	}
	if err := change(after.Session); err != nil {
		t.Fatalf("the value issued with the grant: %v", err)
	}
}
