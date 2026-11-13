package identity

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

// StartSignUp checks setupToken and the email it is bound to, and mails a
// verification link to that email (SEC-158 (b); lead ruling Q4 on the 1.4
// plan). A token that is missing, malformed, unknown, expired, used, bound to
// another email, or whose email already has an account gets ErrTokenRejected
// and no mail (SEC-30).
func (s *Service) StartSignUp(ctx context.Context, client Client, setupToken, email string) error {
	allowed, done := s.tokenUse.attempt("", client.Address)
	if !allowed {
		return ErrRateLimited
	}
	boundEmail, setupHash, err := s.checkSetupToken(ctx, setupToken, email)
	if err != nil {
		done(outcomeOf(err))
		return err
	}
	done(Succeeded)
	link, linkHash := newToken()
	now := s.config.Now()
	err = database.Write(ctx, s.pool, database.Scope{TokenHash: setupHash}, func(tx pgx.Tx) error {
		return identitydb.New(tx).UpsertSignupVerification(ctx, identitydb.UpsertSignupVerificationParams{
			TokenHash: linkHash, SetupTokenHash: setupHash, Email: boundEmail,
			CreatedAt: now, ExpiresAt: now.Add(verificationLinkLifetime),
		})
	})
	if err != nil {
		return fmt.Errorf("store the verification link: %w", err)
	}
	// After commit only: a retried transaction must not mail twice.
	return s.config.SendMail(ctx, Message{To: boundEmail, Kind: MessageVerification, Link: s.config.Origin + "/signup/verify#" + link})
}

// checkSetupToken returns the email setupToken is bound to, as issued, and
// the token's hash, or ErrTokenRejected.
func (s *Service) checkSetupToken(ctx context.Context, setupToken, email string) (string, []byte, error) {
	setupHash, ok := hashToken(setupToken)
	if !ok {
		return "", nil, ErrTokenRejected
	}
	var row identitydb.GetSetupTokenForEmailRow
	err := database.Read(ctx, s.pool, database.Scope{TokenHash: setupHash}, func(tx pgx.Tx) error {
		var err error
		row, err = identitydb.New(tx).GetSetupTokenForEmail(ctx, identitydb.GetSetupTokenForEmailParams{TokenHash: setupHash, Email: email})
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) || err == nil && (row.UsedAt.Valid || !s.config.Now().Before(row.ExpiresAt)) {
		return "", nil, ErrTokenRejected
	}
	if err != nil {
		return "", nil, fmt.Errorf("read the setup token: %w", err)
	}
	// 10.1 §5.1 rule 3: the email lookup is a read of its own.
	err = database.Read(ctx, s.pool, database.Scope{Email: row.Email}, func(tx pgx.Tx) error {
		_, err := identitydb.New(tx).GetAccountIDByEmail(ctx, row.Email)
		return err
	})
	switch {
	case err == nil:
		return "", nil, ErrTokenRejected // the email already has an account
	case !errors.Is(err, pgx.ErrNoRows):
		return "", nil, fmt.Errorf("look up the email: %w", err)
	}
	return row.Email, setupHash, nil
}

// CompleteSignUp opens the account, its credential and its first session in
// one transaction when the emailed link is opened, and uses up the setup
// token (SEC-158; FR-FAM-001 AC-10). The person sets their own password here
// (SEC-26). A missing, unknown, expired or used link gets ErrTokenRejected.
func (s *Service) CompleteSignUp(ctx context.Context, client Client, verificationToken, password string) (Tokens, error) {
	if client.Kind == "web" {
		return Tokens{}, ErrWebSessionNotYet
	}
	if err := checkPasswordNotCommon(password); err != nil {
		return Tokens{}, err
	}
	allowed, done := s.tokenUse.attempt("", client.Address)
	if !allowed {
		return Tokens{}, ErrRateLimited
	}
	tokens, err := s.completeSignUp(ctx, client, verificationToken, password)
	if err != nil {
		done(outcomeOf(err))
		return Tokens{}, err
	}
	done(Succeeded)
	return tokens, nil
}

func (s *Service) completeSignUp(ctx context.Context, client Client, verificationToken, password string) (Tokens, error) {
	linkHash, ok := hashToken(verificationToken)
	if !ok {
		return Tokens{}, ErrTokenRejected
	}
	var link identitydb.GetSignupVerificationRow
	err := database.Read(ctx, s.pool, database.Scope{TokenHash: linkHash}, func(tx pgx.Tx) error {
		var err error
		link, err = identitydb.New(tx).GetSignupVerification(ctx, linkHash)
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) || err == nil && (link.UsedAt.Valid || !s.config.Now().Before(link.ExpiresAt)) {
		return Tokens{}, ErrTokenRejected
	}
	if err != nil {
		return Tokens{}, fmt.Errorf("read the verification link: %w", err)
	}
	// Hashed outside the transaction, which may run more than once.
	credential, err := s.hasher.hash(ctx, password)
	if err != nil {
		return Tokens{}, fmt.Errorf("hash the password: %w", err)
	}
	account := link.NewAccountID
	var tokens Tokens
	scope := database.Scope{AccountID: account, TokenHash: link.SetupTokenHash}
	err = database.Write(ctx, s.pool, scope, func(tx pgx.Tx) error {
		now := s.config.Now()
		q := identitydb.New(tx)
		err := q.InsertAccount(ctx, identitydb.InsertAccountParams{ID: account, Email: link.Email, CreatedAt: now})
		if uniqueViolation(err) {
			return ErrTokenRejected // the email got an account meanwhile
		}
		if err != nil {
			return err
		}
		used, err := q.UseSetupToken(ctx, identitydb.UseSetupTokenParams{Now: now, AccountID: account, TokenHash: link.SetupTokenHash})
		if err != nil || used == 0 {
			return errors.Join(err, rejectedIfNone(used))
		}
		used, err = q.UseSignupVerification(ctx, identitydb.UseSignupVerificationParams{
			Now: now, TokenHash: linkHash, SetupTokenHash: link.SetupTokenHash,
		})
		if err != nil || used == 0 {
			return errors.Join(err, rejectedIfNone(used))
		}
		err = q.InsertCredential(ctx, identitydb.InsertCredentialParams{
			AccountID: account, Salt: credential.salt, Hash: credential.hash,
			TimeCost: int32(credential.params.Time), MemoryKib: int32(credential.params.MemoryKiB), //nolint:gosec // G115: Argon2 parameters fit in int32
			Parallelism: int16(credential.params.Parallelism), UpdatedAt: now,
		})
		if err != nil {
			return err
		}
		tokens, _, err = openSession(ctx, tx, account, client, now)
		return err
	})
	return tokens, err
}

func rejectedIfNone(rows int64) error {
	if rows == 0 {
		return ErrTokenRejected
	}
	return nil
}

func uniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

// outcomeOf counts a rejected token as a failed attempt; any other error
// (the server's own) is not the caller's attempt (10.1 §8.1).
func outcomeOf(err error) Outcome {
	if errors.Is(err, ErrTokenRejected) {
		return Failed
	}
	return NotTried
}
