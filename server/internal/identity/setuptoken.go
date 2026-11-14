package identity

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

// setupTokenLifetime is how long a setup token opens sign-up (SEC-158 (b)
// [Asumsi]; 09.4 §2).
const setupTokenLifetime = 7 * 24 * time.Hour

// IssueSetupToken stores a new setup token bound to email and returns it; it
// is shown once and kept only as its hash (SEC-158 (b); 09.4 §2). The
// operator who runs it never sets or learns a password (SEC-26).
func IssueSetupToken(ctx context.Context, pool *pgxpool.Pool, email string, now time.Time) (token string, expiresAt time.Time, err error) {
	token, hash := NewToken()
	expiresAt = now.Add(setupTokenLifetime)
	err = database.Write(ctx, pool, database.Scope{TokenHash: hash}, func(tx pgx.Tx) error {
		return identitydb.New(tx).InsertSetupToken(ctx, identitydb.InsertSetupTokenParams{
			TokenHash: hash, Email: email, CreatedAt: now, ExpiresAt: expiresAt,
		})
	})
	if err != nil {
		return "", time.Time{}, fmt.Errorf("store the setup token: %w", err)
	}
	return token, expiresAt, nil
}
