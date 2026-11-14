package identity

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

// SignUpSource is the token a sign-up starts from: a setup token (SEC-158
// (b)) or the token of an invitation link (SEC-158 (a)). Exactly one is set;
// anything else gets ErrTokenRejected (plan 1.6 decision 15).
type SignUpSource struct{ SetupToken, InvitationToken string }

// StartSignUp checks source and email and mails a verification link to the
// email (SEC-158; lead ruling Q4 on the 1.4 plan). A setup token must be
// unused, unexpired and bound to email, and the link goes to the email as
// issued. An invitation token must pass invitationOpen, the household's
// check shared with F02 (SEC-158 note (2)); the email is free (SEC-31),
// and one that already has an account is answered like a fresh one. A token
// that fails, a setup token whose email already has an account, or an
// invitation link an account was already born from gets ErrTokenRejected and
// no mail (SEC-30).
func (s *Service) StartSignUp(ctx context.Context, client Client, source SignUpSource, email string,
	invitationOpen func(context.Context, []byte) error,
) error {
	allowed, done := s.tokenUse.attempt("", client.Address)
	if !allowed {
		return ErrRateLimited
	}
	to, link, err := s.storeVerificationLink(ctx, source, email, invitationOpen)
	if err != nil {
		done(outcomeOf(err))
		return err
	}
	done(Succeeded)
	// After commit only: a retried transaction must not mail twice.
	return s.config.SendMail(ctx, Message{To: to, Kind: MessageVerification, Link: s.config.Origin + "/signup/verify#" + link})
}

// storeVerificationLink checks source and writes a new verification link in
// its sign-up row, replacing the old link, through the source token's hash
// (10.1 §5.1 rule 1). It returns the email to mail and the link.
func (s *Service) storeVerificationLink(ctx context.Context, source SignUpSource, email string,
	invitationOpen func(context.Context, []byte) error,
) (to, link string, err error) {
	link, linkHash := NewToken()
	now := s.config.Now()
	expiresAt := now.Add(verificationLinkLifetime)
	switch {
	case source.SetupToken != "" && source.InvitationToken == "":
		var setupHash []byte
		to, setupHash, err = s.checkSetupToken(ctx, source.SetupToken, email)
		if err != nil {
			return "", "", err
		}
		err = database.Write(ctx, s.pool, database.Scope{TokenHash: setupHash}, func(tx pgx.Tx) error {
			return identitydb.New(tx).UpsertSignupVerification(ctx, identitydb.UpsertSignupVerificationParams{
				TokenHash: linkHash, SetupTokenHash: setupHash, Email: to, CreatedAt: now, ExpiresAt: expiresAt,
			})
		})
		if err != nil {
			return "", "", fmt.Errorf("store the verification link: %w", err)
		}
		return to, link, nil
	case source.InvitationToken != "" && source.SetupToken == "":
		invitationHash, ok := HashToken(source.InvitationToken)
		if !ok {
			return "", "", ErrTokenRejected
		}
		if err := invitationOpen(ctx, invitationHash); err != nil {
			return "", "", err
		}
		// No email lookup here: a live invitation link must not tell
		// whether an email has an account, by answer or by rate-limit
		// count (SEC-12, SEC-30; 10.1 §8.1). CompleteSignUp refuses an
		// email that has one, and writes nothing.
		var n int64
		err = database.Write(ctx, s.pool, database.Scope{TokenHash: invitationHash}, func(tx pgx.Tx) (err error) {
			n, err = identitydb.New(tx).UpsertInvitationSignupVerification(ctx, identitydb.UpsertInvitationSignupVerificationParams{
				TokenHash: linkHash, InvitationTokenHash: invitationHash, Email: email, CreatedAt: now, ExpiresAt: expiresAt,
			})
			return err
		})
		if err != nil {
			return "", "", fmt.Errorf("store the verification link: %w", err)
		}
		if n == 0 {
			return "", "", ErrTokenRejected // an account was born from this link (SEC-158 note (1))
		}
		return email, link, nil
	}
	return "", "", ErrTokenRejected
}

// checkSetupToken returns the email setupToken is bound to, as issued, and
// the token's hash, or ErrTokenRejected.
func (s *Service) checkSetupToken(ctx context.Context, setupToken, email string) (string, []byte, error) {
	setupHash, ok := HashToken(setupToken)
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
	if err := s.noAccountFor(ctx, row.Email); err != nil {
		return "", nil, err
	}
	return row.Email, setupHash, nil
}

// noAccountFor is ErrTokenRejected when email already has an account. 10.1
// §5.1 rule 3: the email lookup is a read of its own.
func (s *Service) noAccountFor(ctx context.Context, email string) error {
	err := database.Read(ctx, s.pool, database.Scope{Email: email}, func(tx pgx.Tx) error {
		_, err := identitydb.New(tx).GetAccountIDByEmail(ctx, email)
		return err
	})
	switch {
	case err == nil:
		return ErrTokenRejected
	case !errors.Is(err, pgx.ErrNoRows):
		return fmt.Errorf("look up the email: %w", err)
	}
	return nil
}

// CompleteSignUp opens the account, its credential and its first session in
// one transaction when the emailed link is opened, and uses up the link and,
// for a setup token, the token (SEC-158; FR-FAM-001 AC-10). The person sets
// their own password here (SEC-26). A missing, unknown, expired, used or
// replaced link, or an invitation that no longer passes invitationOpen,
// gets ErrTokenRejected.
func (s *Service) CompleteSignUp(ctx context.Context, client Client, verificationToken, password string,
	invitationOpen func(context.Context, []byte) error,
) (Tokens, error) {
	if err := checkPasswordNotCommon(password); err != nil {
		return Tokens{}, err
	}
	allowed, done := s.tokenUse.attempt("", client.Address)
	if !allowed {
		return Tokens{}, ErrRateLimited
	}
	tokens, err := s.completeSignUp(ctx, client, verificationToken, password, invitationOpen)
	if err != nil {
		done(outcomeOf(err))
		return Tokens{}, err
	}
	done(Succeeded)
	return tokens, nil
}

func (s *Service) completeSignUp(ctx context.Context, client Client, verificationToken, password string,
	invitationOpen func(context.Context, []byte) error,
) (Tokens, error) {
	linkHash, ok := HashToken(verificationToken)
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
	// The row is written through its source token's hash (10.1 §5.1 rule 1).
	sourceHash := link.SetupTokenHash
	if link.Source == "invitation" {
		// SEC-158 note (2): the invitation is checked again when the link
		// is opened.
		if err := invitationOpen(ctx, link.InvitationTokenHash); err != nil {
			return Tokens{}, err
		}
		sourceHash = link.InvitationTokenHash
	}
	// Hashed outside the transaction, which may run more than once.
	credential, err := s.hasher.hash(ctx, password)
	if err != nil {
		return Tokens{}, fmt.Errorf("hash the password: %w", err)
	}
	account := link.NewAccountID
	var tokens Tokens
	err = database.Write(ctx, s.pool, database.Scope{AccountID: account, TokenHash: sourceHash}, func(tx pgx.Tx) error {
		now := s.config.Now()
		q := identitydb.New(tx)
		err := q.InsertAccount(ctx, identitydb.InsertAccountParams{ID: account, Email: link.Email, CreatedAt: now, SignupVia: link.Source})
		if uniqueViolation(err) {
			// The email has an account: one made meanwhile, or, through an
			// invitation, one StartSignUp did not look up. The transaction
			// rolls back, so the sign-up row stays as it was.
			return ErrTokenRejected
		}
		if err != nil {
			return err
		}
		if link.Source == "setup_token" {
			used, err := q.UseSetupToken(ctx, identitydb.UseSetupTokenParams{Now: now, AccountID: account, TokenHash: sourceHash})
			if err != nil || used == 0 {
				return errors.Join(err, rejectedIfNone(used))
			}
		}
		// The account's email is the row's while it still holds the link
		// that was read (security-privacy ruling on Task 8 of plan 1.6).
		used, err := q.UseSignupVerification(ctx, identitydb.UseSignupVerificationParams{
			Now: now, AccountID: account, TokenHash: linkHash, Email: link.Email,
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

// SignupInvitation is the invitation token hash of the sign-up account was
// born from, for from_signup (FR-FAM-004/AC-14; plan 1.6 decision 15), read
// through the account_id = A clause (09.2 §3); an account made with a setup
// token has none: ErrTokenRejected.
func (s *Service) SignupInvitation(ctx context.Context, account pgtype.UUID) ([]byte, error) {
	var hash []byte
	err := database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) (err error) {
		hash, err = identitydb.New(tx).GetSignupInvitation(ctx, account)
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrTokenRejected
	}
	return hash, err
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
