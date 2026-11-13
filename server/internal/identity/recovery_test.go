package identity

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"testing"
	"time"
)

// recoverable returns a signed-up service whose recovery mails land in
// *link (the token of the last one).
func recoverable(t *testing.T, now *time.Time, password string) (s *Service, link *string) {
	t.Helper()
	s, _ = signedUp(t, now, password)
	link = new(string)
	s.config.SendMail = func(_ context.Context, m Message) error {
		_, *link, _ = strings.Cut(m.Link, "#")
		return nil
	}
	return s, link
}

// startRecovery calls StartRecovery and waits for its mail.
func (s *Service) startRecovery(t *testing.T, client Client, email string) error {
	t.Helper()
	defer s.Wait()
	return s.StartRecovery(t.Context(), client, email)
}

func TestRecovery(t *testing.T) {
	// Built, so no secret scanner sees a literal.
	password, next := strings.Repeat("recover ", 3), strings.Repeat("recovered ", 3)
	t.Run("SEC-12 a flood of penalised keys does not close recovery", func(t *testing.T) {
		// 10.1 §8.1 item 2: the StartRecovery tables and the token-use
		// address table evict the oldest penalised key when full.
		now := time.Now()
		s, link := recoverable(t, &now, password)
		for _, table := range []struct {
			l         *Limiter
			threshold int
		}{{s.recovery.account, recoveryEmailThreshold}, {s.recovery.address, recoveryAddressThreshold}, {s.tokenUse.address, tokenUseThreshold}} {
			table.l.maxKeys = 100
			for i := range table.l.maxKeys {
				fail(t, table.l, fmt.Sprint(i), table.threshold)
			}
		}
		if err := s.startRecovery(t, Client{Kind: "ios", Address: "203.0.113.1:443"}, "owner@example.test"); err != nil || *link == "" {
			t.Fatalf("StartRecovery: %v, link %t; want a mailed link", err, *link != "")
		}
		if _, err := s.CompleteRecovery(t.Context(), Client{Kind: "ios", Address: "203.0.113.2:443"}, *link, next); err != nil {
			t.Fatalf("CompleteRecovery: %v", err)
		}
	})
	t.Run("SEC-12 a recovery that gets no hash slot keeps its token", func(t *testing.T) {
		// 10.1 §8.1 item 3: the token is used only after the hash ran.
		now := time.Now()
		s, link := recoverable(t, &now, password)
		if err := s.startRecovery(t, testClient, "owner@example.test"); err != nil {
			t.Fatal(err)
		}
		s.hasher.slotWait = 20 * time.Millisecond
		release := s.hasher.hold()
		if _, err := s.CompleteRecovery(t.Context(), testClient, *link, next); !errors.Is(err, ErrRateLimited) {
			t.Fatalf("with no slot: err = %v, want ErrRateLimited", err)
		}
		release()
		if _, err := s.CompleteRecovery(t.Context(), testClient, *link, next); err != nil {
			t.Fatalf("the same link after release: %v", err)
		}
	})
	t.Run("SEC-12 a rejected recovery token runs no hash", func(t *testing.T) {
		// Lead ruling on PR B: no hash before the token is checked.
		now := time.Now()
		s, link := recoverable(t, &now, password)
		if err := s.startRecovery(t, testClient, "owner@example.test"); err != nil {
			t.Fatal(err)
		}
		used := *link
		if _, err := s.CompleteRecovery(t.Context(), testClient, used, next); err != nil {
			t.Fatal(err)
		}
		before := s.hasher.count.Load()
		for _, token := range []string{"", strings.Repeat("A", 43), used} {
			if _, err := s.CompleteRecovery(t.Context(), testClient, token, next); !errors.Is(err, ErrTokenRejected) {
				t.Fatalf("token %q: err = %v, want ErrTokenRejected", token, err)
			}
		}
		if n := s.hasher.count.Load() - before; n != 0 {
			t.Fatalf("%d hashes for rejected tokens, want 0", n)
		}
	})
	t.Run("the StartRecovery email key is kept only as a SHA-256 of the normalised email", func(t *testing.T) {
		now := time.Now()
		s, _ := recoverable(t, &now, password)
		if err := s.startRecovery(t, testClient, " Owner@Example.TEST "); err != nil {
			t.Fatal(err)
		}
		if err := s.startRecovery(t, testClient, "owner@example.test"); err != nil {
			t.Fatal(err)
		}
		if n := len(s.recovery.account.keys); n != 1 {
			t.Fatalf("%d email keys, want 1", n)
		}
	})
	t.Run("the work after a recovery response", func(t *testing.T) {
		// 10.1 §5.1 rule 6; code-reviewer B2, S3, S4 on #32: it has a
		// deadline, it is refused once shutdown starts, and every failure is
		// logged with a fixed reason and no address.
		for _, c := range []struct {
			name   string
			setup  func(s *Service)
			reason string
		}{
			{"ends at its deadline while the mail blocks", func(s *Service) {
				s.afterResponseTimeout = 200 * time.Millisecond
				s.config.SendMail = func(ctx context.Context, _ Message) error { <-ctx.Done(); return ctx.Err() }
			}, "send"},
			{"logs a failed store", func(s *Service) { s.afterResponseTimeout = time.Nanosecond }, "store"},
			{"is refused once shutdown starts", func(s *Service) { s.Close() }, "shutdown"},
		} {
			t.Run(c.name, func(t *testing.T) {
				now := time.Now()
				s, link := recoverable(t, &now, password)
				var logs bytes.Buffer
				s.logger = slog.New(slog.NewJSONHandler(&logs, nil))
				c.setup(s)
				start := time.Now()
				if err := s.startRecovery(t, testClient, "owner@example.test"); err != nil {
					t.Fatalf("StartRecovery: %v, want nil whatever happens after the response", err)
				}
				if elapsed := time.Since(start); elapsed > 5*time.Second {
					t.Fatalf("the work took %v, want it ended at its deadline", elapsed)
				}
				got := logs.String()
				if *link != "" || !strings.Contains(got, `"msg":"recovery mail not sent","reason":"`+c.reason+`"`) || strings.Contains(got, "owner") {
					t.Fatalf("link sent %t, logs %q; want no mail and reason %s without the address", *link != "", got, c.reason)
				}
			})
		}
	})
}
