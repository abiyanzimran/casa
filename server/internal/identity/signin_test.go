package identity

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

// TestSignInHashesAlike proves a failed sign-in costs the same for an
// unknown email as for a wrong password: one Argon2 verification and the
// same number of database transactions (SEC-12; 10.1 §5.1 rule 5).
func TestSignInHashesAlike(t *testing.T) {
	ctx := t.Context()
	d := databasetest.Migrated(t)
	pool := d.Pool(t, "app_rw")
	var link string
	s := New(pool, Config{
		Origin: "https://casa.example", Now: time.Now, Password: cheapParams,
		SendMail: func(_ context.Context, m Message) error { link = m.Link; return nil },
	})
	client := Client{Kind: "ios", Version: "1.0.0", Address: "192.0.2.1:443"}
	token, _, err := IssueSetupToken(ctx, pool, "owner@example.test", time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if err := s.StartSignUp(ctx, client, token, "owner@example.test"); err != nil {
		t.Fatal(err)
	}
	_, verification, _ := strings.Cut(link, "#")
	if _, err := s.CompleteSignUp(ctx, client, verification, "a long enough passphrase"); err != nil {
		t.Fatal(err)
	}

	cost := func(email string) (hashes, transactions int64) {
		t.Helper()
		hashesBefore, acquiresBefore := s.hasher.count.Load(), pool.Stat().AcquireCount()
		if _, err := s.SignIn(ctx, client, email, "a wrong passphrase"); !errors.Is(err, ErrCredentialsRejected) {
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
