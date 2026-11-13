package identity

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

// openSession writes a new native session for account with its first
// refresh token, and the sign_in event of the security trail (SEC-15,
// SEC-16, SEC-39 (b)). Every sign-in counts as a new device (lead ruling Q9
// (d) on the 1.4 plan).
func openSession(ctx context.Context, tx pgx.Tx, account pgtype.UUID, client Client, now time.Time) (Tokens, error) {
	q := identitydb.New(tx)
	access, accessHash := newToken()
	refresh, refreshHash := newToken()
	session, err := q.InsertSession(ctx, identitydb.InsertSessionParams{
		AccountID: account, ClientKind: client.Kind, TokenHash: accessHash, ClientVersion: client.Version,
		Now: now, IdleTimeout: nativeIdleTimeout, AbsoluteExpiry: now.Add(nativeAbsoluteLifetime),
	})
	if err != nil {
		return Tokens{}, err
	}
	err = q.InsertRefreshToken(ctx, identitydb.InsertRefreshTokenParams{
		TokenHash: refreshHash, SessionID: session, AccountID: account, CreatedAt: now,
	})
	if err != nil {
		return Tokens{}, err
	}
	err = q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
		AccountID: account, SessionID: session, Kind: "sign_in", OccurredAt: now,
	})
	if err != nil {
		return Tokens{}, err
	}
	return Tokens{AccessToken: access, RefreshToken: refresh, AccessTokenExpiresAt: now.Add(accessTokenLifetime)}, nil
}
