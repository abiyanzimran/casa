package identity

import (
	"context"
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
func (s *Service) Reauthenticate(ctx context.Context, client Client, session Session, password, purpose string) (time.Time, error) {
	allowed, done := s.reauth.attempt(string(session.AccountID.Bytes[:]), client.Address)
	if !allowed {
		return time.Time{}, ErrRateLimited
	}
	stored, found, err := s.credentialOf(ctx, session.AccountID)
	if err != nil {
		done(NotTried)
		return time.Time{}, err
	}
	match, err := s.hasher.verify(ctx, stored, password)
	if err != nil {
		done(NotTried)
		return time.Time{}, err
	}
	if !found || !match {
		done(Failed)
		return time.Time{}, ErrCredentialsRejected
	}
	done(Succeeded)
	now := s.config.Now()
	expiresAt := now.Add(reauthLifetime)
	err = database.Write(ctx, s.pool, database.Scope{AccountID: session.AccountID}, func(tx pgx.Tx) error {
		q := identitydb.New(tx)
		err := q.InsertReauthGrant(ctx, identitydb.InsertReauthGrantParams{
			SessionID: session.ID, AccountID: session.AccountID, Purpose: purpose, CreatedAt: now, ExpiresAt: expiresAt,
		})
		if err != nil {
			return err
		}
		return q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
			AccountID: session.AccountID, SessionID: session.ID, Kind: "reauthentication",
			Detail: pgtype.Text{String: purpose, Valid: true}, OccurredAt: now,
		})
	})
	if err != nil {
		return time.Time{}, fmt.Errorf("grant the re-authentication: %w", err)
	}
	return expiresAt, nil
}

// HashPassword refuses a password on the common list, then hashes it with
// Argon2id (SEC-11). It runs outside any transaction, which may run more than
// once; without a free hash slot it is ErrRateLimited.
func (s *Service) HashPassword(ctx context.Context, password string) (PasswordHash, error) {
	if err := checkPasswordNotCommon(password); err != nil {
		return PasswordHash{}, err
	}
	return s.hasher.hash(ctx, password)
}

// ChangePassword uses one open change_password grant of session and
// replaces the account's credential with hash, inside the command
// transaction of tx (SEC-13; SEC-151 (7)). Without a grant it is
// ErrReauthRequired.
func (s *Service) ChangePassword(ctx context.Context, tx pgx.Tx, session Session, hash PasswordHash) error {
	now := s.config.Now()
	q := identitydb.New(tx)
	used, err := q.UseReauthGrant(ctx, identitydb.UseReauthGrantParams{Now: now, SessionID: session.ID, Purpose: ReauthChangePassword})
	switch {
	case err != nil:
		return err
	case used == 0:
		return ErrReauthRequired
	}
	if err := replaceCredential(ctx, q, session.AccountID, hash, now); err != nil {
		return err
	}
	return q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
		AccountID: session.AccountID, SessionID: session.ID, Kind: "password_changed", OccurredAt: now,
	})
}

// replaceCredential writes hash as account's credential.
func replaceCredential(ctx context.Context, q *identitydb.Queries, account pgtype.UUID, hash PasswordHash, now time.Time) error {
	updated, err := q.UpdateCredential(ctx, identitydb.UpdateCredentialParams{
		AccountID: account, Salt: hash.salt, Hash: hash.hash,
		TimeCost: int32(hash.params.Time), MemoryKib: int32(hash.params.MemoryKiB), //nolint:gosec // G115: Argon2 parameters fit in int32
		Parallelism: int16(hash.params.Parallelism), UpdatedAt: now,
	})
	if err == nil && updated != 1 {
		err = fmt.Errorf("%d credentials updated, want 1", updated)
	}
	return err
}
