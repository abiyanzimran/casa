package identity

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

// CookieSession is a web session authenticated by its cookie.
type CookieSession struct {
	Session
	// NewCookie is set when this request rotated the cookie's value (SEC-161
	// (c); plan 1.5 decision 12); the API layer sends it with the response.
	NewCookie *Cookie
}

// AuthenticateCookie returns the web session of a cookie value (SEC-161 (a)).
// A missing, malformed or unknown value, or the token of a native session,
// is ErrUnauthenticated, never ErrSessionRevoked, so the browser keeps its
// local data (SEC-161 (c) rule (5)); an ended session says why. A personal
// session is touched as Authenticate touches a native one; a shared one only
// by TouchUserAction (SEC-165 (g); plan 1.5 decision 13).
func (s *Service) AuthenticateCookie(ctx context.Context, cookie string) (CookieSession, error) {
	hash, ok := hashToken(cookie)
	if !ok {
		return CookieSession{}, ErrUnauthenticated
	}
	var row identitydb.GetSessionByCookieHashRow
	err := database.Read(ctx, s.pool, database.Scope{TokenHash: hash}, func(tx pgx.Tx) error {
		var err error
		row, err = identitydb.New(tx).GetSessionByCookieHash(ctx, hash)
		return err
	})
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return CookieSession{}, ErrUnauthenticated
	case err != nil:
		return CookieSession{}, fmt.Errorf("read the session: %w", err)
	case row.ClientKind != "web":
		return CookieSession{}, ErrUnauthenticated // SEC-161 (b): a native token is not a cookie
	case !row.Current:
		// A previous or discarded value: fail closed until rotation
		// classifies them (SEC-161 (c); plan 1.5 decision 12).
		return CookieSession{}, ErrUnauthenticated
	}
	now := s.config.Now()
	if err := sessionEnded(now, row.RevokedAt, row.LastUsedAt, row.IdleTimeout, row.AbsoluteExpiry); err != nil {
		return CookieSession{}, err
	}
	session := Session{ID: row.ID, AccountID: row.AccountID, ClientKind: row.ClientKind, DeviceMode: row.DeviceMode, LastUsedAt: row.LastUsedAt}
	if session.DeviceMode == "personal" {
		if session, err = s.touch(ctx, session, now); err != nil {
			return CookieSession{}, err
		}
	}
	return CookieSession{Session: session}, nil
}

// TouchUserAction moves the idle limit of a shared web session for a request
// the client marks as a user's action; nothing else does, so a tab left open
// is not kept alive by background requests, and the absolute limit never
// moves (SEC-165 (g); plan 1.5 decision 13). Any other session is left
// alone. A session that ended since it was read is ErrSessionExpired.
func (s *Service) TouchUserAction(ctx context.Context, session Session) error {
	if session.ClientKind != "web" || session.DeviceMode != "shared" {
		return nil
	}
	_, err := s.touch(ctx, session, s.config.Now())
	return err
}
