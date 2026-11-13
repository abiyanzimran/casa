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

// ReauthChangePassword is the purpose of a re-authentication for
// ChangePassword, the first action that needs one (SEC-13; lead ruling Q7 on
// the 1.4 plan).
const ReauthChangePassword = "change_password"

// Reauthenticate checks the password of session's account again and grants
// one action of purpose in that session until the returned time (SEC-13). A
// wrong password is ErrCredentialsRejected. Attempts are limited per account
// and per address and count only when the hash ran (SEC-12; 10.1 §8.1).
//
// A web session also gets a new cookie value, returned for the response,
// in the same transaction as the grant, which serves only that value
// (SEC-161 (c): "selalu saat … autentikasi ulang", closing session
// fixation; note 1.5 rules (6), (7)). A native session gets no cookie and
// its grant is unchanged.
func (s *Service) Reauthenticate(ctx context.Context, client Client, session Session, password, purpose string) (time.Time, *Cookie, error) {
	allowed, done := s.reauth.attempt(string(session.AccountID.Bytes[:]), client.Address)
	if !allowed {
		return time.Time{}, nil, ErrRateLimited
	}
	stored, found, err := s.credentialOf(ctx, session.AccountID)
	if err != nil {
		done(NotTried)
		return time.Time{}, nil, err
	}
	match, err := s.hasher.verify(ctx, stored, password)
	if err != nil {
		done(NotTried)
		return time.Time{}, nil, err
	}
	if !found || !match {
		done(Failed)
		return time.Time{}, nil, ErrCredentialsRejected
	}
	done(Succeeded)
	now := s.config.Now()
	expiresAt := now.Add(reauthLifetime)
	var cookie *Cookie
	err = database.Write(ctx, s.pool, database.Scope{AccountID: session.AccountID}, func(tx pgx.Tx) error {
		cookie = nil // fn may run again: only the committed attempt's value
		q := identitydb.New(tx)
		var cookieHash []byte // NULL for native
		if session.ClientKind == "web" {
			rotated, err := rotateCookie(ctx, q, session, now)
			if err != nil {
				return err
			}
			cookie = &rotated
			cookieHash, _ = hashToken(rotated.Value) // always well-formed: newToken made it
		}
		err := q.InsertReauthGrant(ctx, identitydb.InsertReauthGrantParams{
			SessionID: session.ID, AccountID: session.AccountID, Purpose: purpose, CreatedAt: now, ExpiresAt: expiresAt,
			CookieHash: cookieHash,
		})
		if err != nil {
			return err
		}
		return q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
			AccountID: session.AccountID, SessionID: session.ID, Kind: "reauthentication",
			Detail: pgtype.Text{String: purpose, Valid: true}, OccurredAt: now,
		})
	})
	switch {
	case errors.Is(err, ErrSessionExpired):
		return time.Time{}, nil, err
	case err != nil:
		return time.Time{}, nil, fmt.Errorf("grant the re-authentication: %w", err)
	}
	return expiresAt, cookie, nil
}

// hashPassword refuses a password on the common list, then hashes it with
// Argon2id (SEC-11). It runs outside any transaction, which may run more than
// once; without a free hash slot it is ErrRateLimited.
func (s *Service) hashPassword(ctx context.Context, password string) (PasswordHash, error) {
	if err := checkPasswordNotCommon(password); err != nil {
		return PasswordHash{}, err
	}
	return s.hasher.hash(ctx, password)
}

// ChangePassword uses one open change_password grant of session, replaces
// the account's credential with hash, and revokes every other live session
// of the account, inside the command transaction of tx (SEC-13; SEC-151 (7);
// SEC-24 per the security-privacy decision on PR #32: session stays, no
// household event, no opt-out). Without a grant it is ErrReauthRequired.
func (s *Service) ChangePassword(ctx context.Context, tx pgx.Tx, session Session, hash PasswordHash) error {
	now := s.config.Now()
	q := identitydb.New(tx)
	used, err := q.UseReauthGrant(ctx, identitydb.UseReauthGrantParams{
		Now: now, SessionID: session.ID, Purpose: ReauthChangePassword, CookieHash: session.CookieHash,
	})
	switch {
	case err != nil:
		return err
	case used == 0:
		return ErrReauthRequired
	}
	if err := replaceCredential(ctx, q, session.AccountID, hash, now); err != nil {
		return err
	}
	if err := revokeSessions(ctx, q, session.AccountID, session.ID, "password_changed", now); err != nil {
		return err
	}
	return q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
		AccountID: session.AccountID, SessionID: session.ID, Kind: "password_changed", OccurredAt: now,
	})
}

// replaceCredential writes hash as account's credential and uses up every
// open recovery link of the account, so no older link resets it again
// (SEC-15, SEC-24; lead ruling on PR B).
func replaceCredential(ctx context.Context, q *identitydb.Queries, account pgtype.UUID, hash PasswordHash, now time.Time) error {
	updated, err := q.UpdateCredential(ctx, identitydb.UpdateCredentialParams{
		AccountID: account, Salt: hash.salt, Hash: hash.hash,
		TimeCost: int32(hash.params.Time), MemoryKib: int32(hash.params.MemoryKiB), //nolint:gosec // G115: Argon2 parameters fit in int32
		Parallelism: int16(hash.params.Parallelism), UpdatedAt: now,
	})
	switch {
	case err != nil:
		return err
	case updated != 1:
		return fmt.Errorf("%d credentials updated, want 1", updated)
	}
	return q.UseAccountRecoveryTokens(ctx, identitydb.UseAccountRecoveryTokensParams{Now: now, AccountID: account})
}

// HashNewPassword hashes the new password of ChangePassword, outside the
// command transaction. It hashes only when session holds an open
// change_password grant, so a session without one cannot spend the hash
// slots SignIn needs (lead ruling on PR B; 10.1 §8.1 item 3); else it is
// ErrReauthRequired. ChangePassword still uses the grant atomically: this
// read is not the authority.
func (s *Service) HashNewPassword(ctx context.Context, session Session, password string) (PasswordHash, error) {
	var open bool
	err := database.Read(ctx, s.pool, database.Scope{AccountID: session.AccountID}, func(tx pgx.Tx) error {
		var err error
		open, err = identitydb.New(tx).HasOpenReauthGrant(ctx, identitydb.HasOpenReauthGrantParams{
			SessionID: session.ID, Purpose: ReauthChangePassword, Now: s.config.Now(), CookieHash: session.CookieHash,
		})
		return err
	})
	switch {
	case err != nil:
		return PasswordHash{}, fmt.Errorf("read the re-authentication: %w", err)
	case !open:
		return PasswordHash{}, ErrReauthRequired
	}
	return s.hashPassword(ctx, password)
}
