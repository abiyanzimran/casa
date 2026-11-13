package identity

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

// openSession writes a new session for account and the sign_in event of the
// security trail (SEC-15, SEC-16, SEC-39 (b)). Every sign-in counts as a new
// device (lead ruling Q9 (d) on the 1.4 plan).
//
// A native session gets an access token and its first refresh token. A web
// session gets one token, carried only in its cookie, no refresh token, the
// lifetimes of its device mode, and the summary of its User-Agent (SEC-161
// (b), (d), (e); SEC-165 (a); plan 1.5 decisions 7 to 10).
func openSession(ctx context.Context, tx pgx.Tx, account pgtype.UUID, client Client, now time.Time) (Tokens, pgtype.UUID, error) {
	idle, lifetime, mode, summary := nativeIdleTimeout, nativeAbsoluteLifetime, "personal", ""
	if client.Kind == "web" {
		switch client.DeviceMode {
		case "personal":
			idle, lifetime = personalIdleTimeout, personalAbsoluteLifetime
		case "shared":
			idle, lifetime = sharedIdleTimeout, sharedAbsoluteLifetime
		default: // unreachable: the boundary requires a mode from web (plan 1.5 decision 9)
			return Tokens{}, pgtype.UUID{}, errors.New("a web session needs a device mode")
		}
		mode, summary = client.DeviceMode, summarizeUserAgent(client.UserAgent)
	}
	q := identitydb.New(tx)
	token, tokenHash := newToken()
	session, err := q.InsertSession(ctx, identitydb.InsertSessionParams{
		AccountID: account, ClientKind: client.Kind, DeviceMode: mode, UaSummary: summary, TokenHash: tokenHash,
		ClientVersion: client.Version, Now: now, IdleTimeout: idle, AbsoluteExpiry: now.Add(lifetime),
	})
	if err != nil {
		return Tokens{}, pgtype.UUID{}, err
	}
	err = q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
		AccountID: account, SessionID: session, Kind: "sign_in", OccurredAt: now,
	})
	if err != nil {
		return Tokens{}, pgtype.UUID{}, err
	}
	if client.Kind == "web" {
		return Tokens{Cookie: sessionCookie(token, mode, now.Add(lifetime), now)}, session, nil
	}
	refresh, refreshHash := newToken()
	err = q.InsertRefreshToken(ctx, identitydb.InsertRefreshTokenParams{
		TokenHash: refreshHash, SessionID: session, AccountID: account, CreatedAt: now,
	})
	if err != nil {
		return Tokens{}, pgtype.UUID{}, err
	}
	return Tokens{AccessToken: token, RefreshToken: refresh, AccessTokenExpiresAt: now.Add(accessTokenLifetime)}, session, nil
}
