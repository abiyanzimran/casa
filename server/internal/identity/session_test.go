package identity

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

func TestEndedSession(t *testing.T) {
	t.Run("SEC-15 an ended session is never touched or rotated", func(t *testing.T) {
		// Lead ruling on #32 (fail closed): the queries that move
		// last_used_at refuse a revoked, idle-expired or absolute-expired
		// session, whatever a caller checked before.
		password := "a long enough passphrase"
		for _, c := range []struct {
			name    string
			revoked bool
			after   time.Duration
		}{
			{"idle", false, 30 * 24 * time.Hour},
			{"absolute", false, 180 * 24 * time.Hour}, // idle too: app_rw cannot move absolute_expiry
			{"revoked", true, 0},
		} {
			t.Run(c.name, func(t *testing.T) {
				now := time.Now()
				s, pool, verification := startedSignUp(t, &now)
				tokens, err := s.CompleteSignUp(t.Context(), testClient, verification, "owner@example.test", password, noInvitation)
				if err != nil {
					t.Fatal(err)
				}
				session, err := s.Authenticate(t.Context(), tokens.AccessToken)
				if err != nil {
					t.Fatal(err)
				}
				if c.revoked {
					err := database.Write(t.Context(), pool, database.Scope{AccountID: session.AccountID}, func(tx pgx.Tx) error {
						return s.SignOut(t.Context(), tx, session)
					})
					if err != nil {
						t.Fatal(err)
					}
				}
				now = now.Add(c.after)
				var touched, rotated int64
				err = database.Write(t.Context(), pool, database.Scope{AccountID: session.AccountID}, func(tx pgx.Tx) error {
					q := identitydb.New(tx)
					var err error
					if touched, err = q.TouchSession(t.Context(), identitydb.TouchSessionParams{Now: now, ID: session.ID}); err != nil {
						return err
					}
					_, hash := NewToken()
					rotated, err = q.RotateAccessToken(t.Context(), identitydb.RotateAccessTokenParams{TokenHash: hash, Now: now, ID: session.ID})
					return err
				})
				if err != nil || touched != 0 || rotated != 0 {
					t.Fatalf("touched %d, rotated %d (%v); want neither", touched, rotated, err)
				}
			})
		}
	})
}
