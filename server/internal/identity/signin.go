package identity

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

// Session is the authenticated caller of a request.
type Session struct {
	ID, AccountID pgtype.UUID
	ClientKind    string
	// DeviceMode is "personal" or "shared"; native sessions are personal
	// (plan 1.5 decision 9).
	DeviceMode string
	// LastUsedAt is the base of the idle limit when the session was read;
	// TouchUserAction reads it (SEC-165 (g)).
	LastUsedAt time.Time
}

// SignIn opens a native or web session (SEC-12, SEC-15, SEC-161). An unknown
// email and a wrong password both get ErrCredentialsRejected after the same
// work: two read transactions and one password verification (10.1 §5.1 rule
// 5). The
// email is trimmed and lower-cased once, for both the lookup and the rate
// limit key.
func (s *Service) SignIn(ctx context.Context, client Client, email, password string) (Tokens, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	// The key is a hash so the table in memory holds no email (10.1 §8.1).
	emailKey := sha256.Sum256([]byte(email))
	allowed, done := s.signIn.attempt(string(emailKey[:]), client.Address)
	if !allowed {
		return Tokens{}, ErrRateLimited
	}
	account, known, stored, err := s.readCredential(ctx, email)
	if err != nil {
		done(NotTried)
		return Tokens{}, err
	}
	// Outside any transaction: a retried transaction must not hash twice.
	match, err := s.hasher.verify(ctx, stored, password)
	if err != nil {
		done(NotTried)
		return Tokens{}, err
	}
	if !known || !match {
		done(Failed)
		return Tokens{}, ErrCredentialsRejected
	}
	done(Succeeded)
	var tokens Tokens
	err = database.Write(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) error {
		var err error
		tokens, _, err = openSession(ctx, tx, account, client, s.config.Now())
		return err
	})
	if err != nil {
		return Tokens{}, fmt.Errorf("open the session: %w", err)
	}
	return tokens, nil
}

// readCredential finds the account of email and its credential. For an
// unknown email it still runs the credential read, with a random id, and
// returns the decoy, so both cases run the same transactions.
func (s *Service) readCredential(ctx context.Context, email string) (account pgtype.UUID, known bool, stored PasswordHash, err error) {
	// 10.1 §5.1 rules 3, 4: the email lookup is a read of its own and
	// selects only the id.
	err = database.Read(ctx, s.pool, database.Scope{Email: email}, func(tx pgx.Tx) error {
		var err error
		account, err = identitydb.New(tx).GetAccountIDByEmail(ctx, email)
		return err
	})
	known = err == nil
	if errors.Is(err, pgx.ErrNoRows) {
		account = pgtype.UUID{Bytes: [16]byte(randomBytes(16)), Valid: true}
	} else if err != nil {
		return account, false, PasswordHash{}, fmt.Errorf("look up the email: %w", err)
	}
	stored, found, err := s.credentialOf(ctx, account)
	if err != nil {
		return account, false, PasswordHash{}, err
	}
	return account, known && found, stored, nil
}

// credentialOf reads account's credential; without one it returns the decoy
// and false.
func (s *Service) credentialOf(ctx context.Context, account pgtype.UUID) (PasswordHash, bool, error) {
	var row identitydb.GetCredentialRow
	err := database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) error {
		var err error
		row, err = identitydb.New(tx).GetCredential(ctx, account)
		return err
	})
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return s.decoy, false, nil
	case err != nil:
		return PasswordHash{}, false, fmt.Errorf("read the credential: %w", err)
	}
	return PasswordHash{
		params: PasswordParams{
			Time: uint32(row.TimeCost), MemoryKiB: uint32(row.MemoryKib), Parallelism: uint8(row.Parallelism), //nolint:gosec // G115: stored from PasswordParams
			KeyBytes: uint32(len(row.Hash)), SaltBytes: uint32(len(row.Salt)), //nolint:gosec // G115: 16 and 32 bytes
		},
		salt: row.Salt, hash: row.Hash,
	}, true, nil
}

// Authenticate returns the session of a native access token (SEC-15). A
// missing, malformed or unknown token, a web session (its token travels only
// in a cookie, SEC-161 (b)), or a token older than 15 minutes is
// ErrUnauthenticated; a revoked or expired session says so.
func (s *Service) Authenticate(ctx context.Context, accessToken string) (Session, error) {
	hash, ok := hashToken(accessToken)
	if !ok {
		return Session{}, ErrUnauthenticated
	}
	var row identitydb.GetSessionByTokenHashRow
	err := database.Read(ctx, s.pool, database.Scope{TokenHash: hash}, func(tx pgx.Tx) error {
		var err error
		row, err = identitydb.New(tx).GetSessionByTokenHash(ctx, hash)
		return err
	})
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return Session{}, ErrUnauthenticated
	case err != nil:
		return Session{}, fmt.Errorf("read the session: %w", err)
	case row.ClientKind == "web":
		return Session{}, ErrUnauthenticated
	}
	now := s.config.Now()
	if err := sessionEnded(now, row.RevokedAt, row.LastUsedAt, row.IdleTimeout, row.AbsoluteExpiry); err != nil {
		return Session{}, err
	}
	if !now.Before(row.TokenIssuedAt.Add(accessTokenLifetime)) {
		return Session{}, ErrUnauthenticated
	}
	session := Session{ID: row.ID, AccountID: row.AccountID, ClientKind: row.ClientKind, DeviceMode: "personal", LastUsedAt: row.LastUsedAt}
	return s.touch(ctx, session, now, touchAfter)
}

// touch moves session's idle limit to now when last_used_at is more than
// after old; touchAfter spares a burst of requests a write on each one. A
// session that ended since it was read is ErrSessionExpired (fail closed).
func (s *Service) touch(ctx context.Context, session Session, now time.Time, after time.Duration) (Session, error) {
	if now.Sub(session.LastUsedAt) <= after {
		return session, nil // no transaction
	}
	var touched Session
	err := database.Write(ctx, s.pool, database.Scope{AccountID: session.AccountID}, func(tx pgx.Tx) error {
		var err error
		touched, err = touchSession(ctx, identitydb.New(tx), session, now, after)
		return err
	})
	if errors.Is(err, ErrSessionExpired) {
		return Session{}, err
	}
	if err != nil {
		return Session{}, fmt.Errorf("touch the session: %w", err)
	}
	return touched, nil
}

// touchSession is touch inside the caller's transaction.
func touchSession(ctx context.Context, q *identitydb.Queries, session Session, now time.Time, after time.Duration) (Session, error) {
	if now.Sub(session.LastUsedAt) <= after {
		return session, nil
	}
	touched, err := q.TouchSession(ctx, identitydb.TouchSessionParams{Now: now, ID: session.ID})
	switch {
	case err != nil:
		return Session{}, err
	case touched == 0:
		return Session{}, ErrSessionExpired // it ended since the read: fail closed
	}
	session.LastUsedAt = now
	return session, nil
}

// sessionEnded reports why a session no longer serves, or nil.
func sessionEnded(now time.Time, revokedAt pgtype.Timestamptz, lastUsedAt time.Time, idle time.Duration, absoluteExpiry time.Time) error {
	switch {
	case revokedAt.Valid:
		return ErrSessionRevoked
	case !now.Before(absoluteExpiry), !now.Before(lastUsedAt.Add(idle)):
		return ErrSessionExpired
	}
	return nil
}

// Refresh rotates a native session's access and refresh tokens (SEC-15). A
// refresh token presented a second time revokes its whole session, writes
// the security trail, and gets ErrSessionRevoked; an unknown one gets
// ErrUnauthenticated.
func (s *Service) Refresh(ctx context.Context, refreshToken string) (Tokens, error) {
	hash, ok := hashToken(refreshToken)
	if !ok {
		return Tokens{}, ErrUnauthenticated
	}
	var found identitydb.GetRefreshTokenRow
	err := database.Read(ctx, s.pool, database.Scope{TokenHash: hash}, func(tx pgx.Tx) error {
		var err error
		found, err = identitydb.New(tx).GetRefreshToken(ctx, hash)
		return err
	})
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return Tokens{}, ErrUnauthenticated
	case err != nil:
		return Tokens{}, fmt.Errorf("read the refresh token: %w", err)
	}
	var tokens Tokens
	reused := false
	// 10.1 §5.1: the write needs only the account; the lookup found it.
	err = database.Write(ctx, s.pool, database.Scope{AccountID: found.AccountID}, func(tx pgx.Tx) error {
		tokens, reused = Tokens{}, false
		now := s.config.Now()
		q := identitydb.New(tx)
		session, err := q.GetSession(ctx, found.SessionID)
		if err != nil {
			return err
		}
		if err := sessionEnded(now, session.RevokedAt, session.LastUsedAt, session.IdleTimeout, session.AbsoluteExpiry); err != nil {
			return err
		}
		used, err := q.UseRefreshToken(ctx, identitydb.UseRefreshTokenParams{Now: now, TokenHash: hash})
		if err != nil {
			return err
		}
		if used == 0 {
			reused = true
			return revokeSession(ctx, q, found.AccountID, found.SessionID, "reuse_detected", now)
		}
		tokens, err = rotate(ctx, q, found.AccountID, found.SessionID, now)
		return err
	})
	if err == nil && reused {
		return Tokens{}, ErrSessionRevoked
	}
	return tokens, err
}

// rotate gives a session a new access token and a new refresh token.
func rotate(ctx context.Context, q *identitydb.Queries, account, session pgtype.UUID, now time.Time) (Tokens, error) {
	access, accessHash := newToken()
	refresh, refreshHash := newToken()
	rotated, err := q.RotateAccessToken(ctx, identitydb.RotateAccessTokenParams{TokenHash: accessHash, Now: now, ID: session})
	switch {
	case err != nil:
		return Tokens{}, err
	case rotated == 0:
		return Tokens{}, ErrSessionExpired // it ended: fail closed
	}
	err = q.InsertRefreshToken(ctx, identitydb.InsertRefreshTokenParams{
		TokenHash: refreshHash, SessionID: session, AccountID: account, CreatedAt: now,
	})
	if err != nil {
		return Tokens{}, err
	}
	return Tokens{AccessToken: access, RefreshToken: refresh, AccessTokenExpiresAt: now.Add(accessTokenLifetime)}, nil
}

// revokeSession revokes a live session and records it in the security trail
// with its reason (SEC-16, SEC-39 (b)).
func revokeSession(ctx context.Context, q *identitydb.Queries, account, session pgtype.UUID, reason string, now time.Time) error {
	revoked, err := q.RevokeSession(ctx, identitydb.RevokeSessionParams{Now: now, Reason: reason, ID: session})
	if err != nil || revoked == 0 {
		return err
	}
	return q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
		AccountID: account, SessionID: session, Kind: "session_revoked",
		Detail: pgtype.Text{String: reason, Valid: true}, OccurredAt: now,
	})
}

// SignOut revokes session (reason 'user') and records it in the security
// trail (09.3 §1 "keluar akun"; SEC-16). It runs inside the command
// transaction of tx, whose scope is the session's account.
func (s *Service) SignOut(ctx context.Context, tx pgx.Tx, session Session) error {
	return revokeSession(ctx, identitydb.New(tx), session.AccountID, session.ID, "user", s.config.Now())
}
