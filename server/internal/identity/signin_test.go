package identity

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

// testClient is the native client of the service-level tests.
var testClient = Client{Kind: "ios", Version: "1.0.0", Address: "192.0.2.1:443"}

// startedSignUp returns a service whose clock is *now, after StartSignUp for
// owner@example.test, with the verification token of the emailed link.
func startedSignUp(t *testing.T, now *time.Time) (s *Service, pool *pgxpool.Pool, verification string) {
	t.Helper()
	ctx := t.Context()
	pool = databasetest.Migrated(t).Pool(t, "app_rw")
	var link string
	s = New(pool, Config{
		Origin: "https://casa.example", Now: func() time.Time { return *now }, Password: cheapParams,
		SendMail: func(_ context.Context, m Message) error { link = m.Link; return nil },
	})
	token, _, err := IssueSetupToken(ctx, pool, "owner@example.test", *now)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.StartSignUp(ctx, testClient, token, "owner@example.test"); err != nil {
		t.Fatal(err)
	}
	_, verification, _ = strings.Cut(link, "#")
	return s, pool, verification
}

// signedUp returns a service whose clock is *now, with one account
// owner@example.test signed up with password.
func signedUp(t *testing.T, now *time.Time, password string) (*Service, *pgxpool.Pool) {
	t.Helper()
	s, pool, verification := startedSignUp(t, now)
	if _, err := s.CompleteSignUp(t.Context(), testClient, verification, password); err != nil {
		t.Fatal(err)
	}
	return s, pool
}

// TestSignInHashesAlike proves a failed sign-in costs the same for an
// unknown email as for a wrong password: one Argon2 verification and the
// same number of database transactions (SEC-12; 10.1 §5.1 rule 5).
func TestSignInHashesAlike(t *testing.T) {
	now := time.Now()
	s, pool := signedUp(t, &now, "a long enough passphrase")
	client := Client{Kind: "ios", Version: "1.0.0", Address: "192.0.2.1:443"}
	cost := func(email string) (hashes, transactions int64) {
		t.Helper()
		hashesBefore, acquiresBefore := s.hasher.count.Load(), pool.Stat().AcquireCount()
		if _, err := s.SignIn(t.Context(), client, email, "a wrong passphrase"); !errors.Is(err, ErrCredentialsRejected) {
			t.Fatalf("sign-in as %s: %v, want ErrCredentialsRejected", email, err)
		}
		return s.hasher.count.Load() - hashesBefore, pool.Stat().AcquireCount() - acquiresBefore
	}
	unknownHashes, unknownTransactions := cost("nobody@example.test")
	wrongHashes, wrongTransactions := cost("owner@example.test")
	if unknownHashes != 1 || wrongHashes != 1 || unknownTransactions != wrongTransactions {
		t.Fatalf("unknown email: %d hashes, %d transactions; wrong password: %d hashes, %d transactions; want 1 hash each and equal transactions",
			unknownHashes, unknownTransactions, wrongHashes, wrongTransactions)
	}
}

func TestSignInLimits(t *testing.T) {
	const password = "a long enough passphrase"
	t.Run("SEC-12 after 20 consecutive failed sign-ins the next waits, for any email and any letter case", func(t *testing.T) {
		// Each attempt comes from its own address, so only the email key
		// can stop the 21st (10.1 §8.1).
		for _, email := range []string{"owner@example.test", "nobody@example.test"} {
			now := time.Now()
			s, _ := signedUp(t, &now, password)
			signIn := func(i int, email, password string) error {
				_, err := s.SignIn(t.Context(), Client{Kind: "ios", Address: fmt.Sprintf("198.51.100.%d:443", i)}, email, password)
				return err
			}
			for i := range 20 {
				variant := strings.ToUpper(email[:1]) + email[1:]
				if i%2 == 0 {
					variant = " " + strings.ToUpper(email) + " "
				}
				if err := signIn(i, variant, "a wrong passphrase"); !errors.Is(err, ErrCredentialsRejected) {
					t.Fatalf("%s attempt %d: %v, want ErrCredentialsRejected", email, i, err)
				}
			}
			now = now.Add(time.Second - time.Millisecond)
			if err := signIn(20, email, password); !errors.Is(err, ErrRateLimited) {
				t.Fatalf("%s 21st attempt: %v, want ErrRateLimited", email, err)
			}
			if email == "owner@example.test" {
				now = now.Add(time.Millisecond)
				if err := signIn(21, email, password); err != nil {
					t.Fatalf("after the wait: %v", err)
				}
			}
		}
	})
	t.Run("the email key is kept only as a SHA-256 of the normalised email", func(t *testing.T) {
		// 10.1 §8.1: the limiter table in memory holds no email address.
		now := time.Now()
		s, _ := signedUp(t, &now, password)
		if _, err := s.SignIn(t.Context(), Client{Kind: "ios", Address: "198.51.100.1:443"}, " Owner@Example.TEST ", "a wrong passphrase"); !errors.Is(err, ErrCredentialsRejected) {
			t.Fatal(err)
		}
		sum := sha256.Sum256([]byte("owner@example.test"))
		if _, ok := s.signIn.account.keys[string(sum[:])]; !ok || len(s.signIn.account.keys) != 1 {
			t.Fatalf("email keys %q, want only the SHA-256 of owner@example.test", slices.Collect(maps.Keys(s.signIn.account.keys)))
		}
	})
}

func TestHashSlots(t *testing.T) {
	// 10.1 §8.1 item 3: a request that gets no hash slot ran no Argon2, so
	// it keeps no count and uses up no token.
	const password = "a long enough passphrase"
	t.Run("SEC-12 an attempt that gets no hash slot is not counted", func(t *testing.T) {
		now := time.Now()
		s, _ := signedUp(t, &now, password)
		s.hasher.slotWait = 20 * time.Millisecond
		signIn := func(password string) error {
			_, err := s.SignIn(t.Context(), testClient, "owner@example.test", password)
			return err
		}
		for range 19 {
			if err := signIn("a wrong passphrase"); !errors.Is(err, ErrCredentialsRejected) {
				t.Fatalf("err = %v, want ErrCredentialsRejected", err)
			}
		}
		release := s.hasher.hold()
		for range 5 {
			if err := signIn("a wrong passphrase"); !errors.Is(err, ErrRateLimited) {
				t.Fatalf("with no slot: err = %v, want ErrRateLimited", err)
			}
		}
		release()
		if err := signIn("a wrong passphrase"); !errors.Is(err, ErrCredentialsRejected) {
			t.Fatalf("the 20th failure: err = %v, want ErrCredentialsRejected", err)
		}
		if err := signIn(password); !errors.Is(err, ErrRateLimited) {
			t.Fatalf("after 20 failures: err = %v, want ErrRateLimited", err)
		}
	})
	t.Run("SEC-12 a sign-up that gets no hash slot keeps its link", func(t *testing.T) {
		now := time.Now()
		s, _, verification := startedSignUp(t, &now)
		s.hasher.slotWait = 20 * time.Millisecond
		release := s.hasher.hold()
		if _, err := s.CompleteSignUp(t.Context(), testClient, verification, password); !errors.Is(err, ErrRateLimited) {
			t.Fatalf("with no slot: err = %v, want ErrRateLimited", err)
		}
		release()
		// The same link still opens the account: neither it nor the setup
		// token was used up, and no account was made.
		if _, err := s.CompleteSignUp(t.Context(), testClient, verification, password); err != nil {
			t.Fatalf("the same link after release: %v", err)
		}
	})
}
