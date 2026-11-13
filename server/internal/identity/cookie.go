package identity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

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

// Rotation of the cookie's value (SEC-161 (c) [Asumsi]; plan 1.5 decision
// 12).
const (
	// cookieRotateAfter: a successful request with a current value at
	// least this old gets a new value.
	cookieRotateAfter = 15 * time.Minute
	// previousCookieGrace: how long the value before a rotation still
	// serves requests already on their way.
	previousCookieGrace = 60 * time.Second
)

// AuthenticateCookie returns the web session of a cookie value (SEC-161 (a))
// and, when this request gave the session a new value, that cookie. The
// value is classified against its session under the row's lock, and every
// write of the outcome is in that one transaction (SEC-161 (c) note 1.5 rule
// (4); plan 1.5 decision 12):
//
//   - current: confirmed if it was not yet (rule (1)); rotated when at least
//     cookieRotateAfter old;
//   - previous, inside its window: served as it is;
//   - previous, after its window: if the current value is confirmed the
//     session is revoked for reuse and the trail gets one row; if not, the
//     current value is discarded and a new one issued, and the presented
//     value gets a new window (rule (2));
//   - discarded: revoked for reuse, one row in the trail (rule (3));
//   - unknown, malformed, or a native session's token: ErrUnauthenticated,
//     never ErrSessionRevoked, so the browser keeps its local data (rule
//     (5); SEC-161 (b)).
//
// An ended session says why. A personal session is touched as Authenticate
// touches a native one; a shared one only by TouchUserAction (SEC-165 (g);
// plan 1.5 decision 13).
func (s *Service) AuthenticateCookie(ctx context.Context, cookie string) (CookieSession, error) {
	hash, ok := HashToken(cookie)
	if !ok {
		return CookieSession{}, ErrUnauthenticated
	}
	// 10.1 §5.1: the lookup without a lock only finds the account.
	var account pgtype.UUID
	err := database.Read(ctx, s.pool, database.Scope{TokenHash: hash}, func(tx pgx.Tx) error {
		var err error
		account, err = identitydb.New(tx).GetSessionByCookieHash(ctx, hash)
		return err
	})
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return CookieSession{}, ErrUnauthenticated
	case err != nil:
		return CookieSession{}, fmt.Errorf("read the session: %w", err)
	}
	var session CookieSession
	reused := false
	err = database.Write(ctx, s.pool, database.Scope{AccountID: account, TokenHash: hash}, func(tx pgx.Tx) error {
		// fn may run again: only the committed attempt's new value is
		// returned.
		var err error
		session, reused, err = classifyCookie(ctx, identitydb.New(tx), hash, s.config.Now())
		return err
	})
	switch {
	case errors.Is(err, ErrUnauthenticated), errors.Is(err, ErrSessionExpired), errors.Is(err, ErrSessionRevoked):
		return CookieSession{}, err
	case err != nil:
		return CookieSession{}, fmt.Errorf("authenticate the cookie: %w", err)
	case reused:
		return CookieSession{}, ErrSessionRevoked
	}
	return session, nil
}

// classifyCookie is AuthenticateCookie's decision on the locked row of the
// value whose hash is hash; reused is true when it revoked the session.
func classifyCookie(ctx context.Context, q *identitydb.Queries, hash []byte, now time.Time) (CookieSession, bool, error) {
	row, err := q.GetSessionByCookieHashForUpdate(ctx, hash)
	switch {
	case errors.Is(err, pgx.ErrNoRows): // rotated away since the lookup: unknown now (rule (5))
		return CookieSession{}, false, ErrUnauthenticated
	case err != nil:
		return CookieSession{}, false, err
	case row.ClientKind != "web":
		return CookieSession{}, false, ErrUnauthenticated // SEC-161 (b): a native token is not a cookie
	}
	if err := sessionEnded(now, row.RevokedAt, row.LastUsedAt, row.IdleTimeout, row.AbsoluteExpiry); err != nil {
		return CookieSession{}, false, err
	}
	result := CookieSession{Session: Session{
		ID: row.ID, AccountID: row.AccountID, ClientKind: row.ClientKind, DeviceMode: row.DeviceMode, LastUsedAt: row.LastUsedAt,
		CookieHash: hash,
	}}
	switch {
	case row.Current && !now.Before(row.TokenIssuedAt.Add(cookieRotateAfter)):
		// The rotation leaves the new value unconfirmed, so no confirm.
		cookie, _, err := rotateCookie(ctx, q, result.Session, now)
		if err != nil {
			return CookieSession{}, false, err
		}
		result.NewCookie = &cookie
	case row.Current:
		if !row.TokenConfirmedAt.Valid { // rule (1)
			if err := q.ConfirmCookie(ctx, identitydb.ConfirmCookieParams{Now: now, ID: row.ID}); err != nil {
				return CookieSession{}, false, err
			}
		}
	case row.Discarded: // rule (3)
		return CookieSession{}, true, revokeSession(ctx, q, row.AccountID, row.ID, "reuse_detected", now)
	case row.Previous && row.PrevValidUntil.Valid && now.Before(row.PrevValidUntil.Time):
		// Inside the window: a request already on its way.
	case row.Previous && row.TokenConfirmedAt.Valid: // rule (2): two holders
		return CookieSession{}, true, revokeSession(ctx, q, row.AccountID, row.ID, "reuse_detected", now)
	case row.Previous: // rule (2): the new value never arrived
		cookie, err := reissueCookie(ctx, q, result.Session, hash, now)
		if err != nil {
			return CookieSession{}, false, err
		}
		result.NewCookie = &cookie
	default: // unreachable: the row matched one of the three
		return CookieSession{}, false, ErrUnauthenticated
	}
	if result.DeviceMode == "personal" {
		if result.Session, err = touchSession(ctx, q, result.Session, now, touchAfter); err != nil {
			return CookieSession{}, false, err
		}
	}
	return result, false, nil
}

// rotateCookie makes a new current value for session; the value the
// browser holds, the current one or an unconfirmed rotation's previous one
// (query RotateCookie), serves as the previous one until now +
// previousCookieGrace (SEC-161 (c)). The token is made here, inside the
// caller's transaction attempt. It also returns the hash of the new value.
func rotateCookie(ctx context.Context, q *identitydb.Queries, session Session, now time.Time) (Cookie, []byte, error) {
	value, hash := NewToken()
	expiry, err := q.RotateCookie(ctx, identitydb.RotateCookieParams{
		PrevValidUntil: now.Add(previousCookieGrace), TokenHash: hash, Now: now, ID: session.ID,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Cookie{}, nil, ErrSessionExpired // it ended: fail closed
	}
	if err != nil {
		return Cookie{}, nil, err
	}
	return sessionCookie(value, session.DeviceMode, expiry, now), hash, nil
}

// reissueCookie discards session's current value, which never reached the
// browser, remembers its hash, and issues a new one; presented, the previous
// value, gets a new window (SEC-161 (c) note 1.5 rules (2), (3)).
func reissueCookie(ctx context.Context, q *identitydb.Queries, session Session, presented []byte, now time.Time) (Cookie, error) {
	value, hash := NewToken()
	expiry, err := q.ReissueCookie(ctx, identitydb.ReissueCookieParams{
		TokenHash: hash, Presented: presented, PrevValidUntil: now.Add(previousCookieGrace), Now: now, ID: session.ID,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Cookie{}, ErrSessionExpired // it ended: fail closed
	}
	if err != nil {
		return Cookie{}, err
	}
	return sessionCookie(value, session.DeviceMode, expiry, now), nil
}

// sessionCookie is the cookie of value for a session of mode that ends at
// absoluteExpiry: in personal mode Max-Age is the remaining absolute
// lifetime, in shared mode there is none (plan 1.5 decision 7; SEC-161 (d),
// SEC-165 (a)).
func sessionCookie(value, mode string, absoluteExpiry, now time.Time) Cookie {
	cookie := Cookie{Value: value}
	if mode == "personal" {
		cookie.MaxAge = absoluteExpiry.Sub(now)
	}
	return cookie
}

// TouchUserAction moves the idle limit of a shared web session for a request
// the client marks as a user's action; nothing else does, so a tab left open
// is not kept alive by background requests, and the absolute limit never
// moves (SEC-165 (g); plan 1.5 decision 13). Any other session is left
// alone. A session that ended since it was read is ErrSessionExpired.
//
// Every action moves it, without touchAfter: the client warns 60 seconds
// before 30 minutes counted from the last input, so the server's limit must
// count from that same action (SEC-165 (g); lead ruling on the Task 4
// review, overriding decision 13's debounce for shared mode).
func (s *Service) TouchUserAction(ctx context.Context, session Session) error {
	if session.ClientKind != "web" || session.DeviceMode != "shared" {
		return nil
	}
	_, err := s.touch(ctx, session, s.config.Now(), 0)
	return err
}
