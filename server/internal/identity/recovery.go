package identity

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
)

// StartRecovery mails a single-use recovery link to email when it has an
// account, and does nothing otherwise; both answer nil (SEC-24, SEC-12).
// Every request counts on the email key and the address key, and nothing
// clears them (10.1 §8.1 row StartRecovery); over the limit it is
// ErrRateLimited.
//
// ponytail: the mail goes synchronously after commit, so a registered email
// answers one write and one send slower; the identity.send_mail job of 1.7
// takes it off the request path (10.1 §5.1 rule 6; 09.3 §9 note).
func (s *Service) StartRecovery(ctx context.Context, client Client, email string) error {
	email = strings.ToLower(strings.TrimSpace(email))
	emailKey := sha256.Sum256([]byte(email)) // no email in memory (10.1 §8.1)
	allowed, done := s.recovery.attempt(string(emailKey[:]), client.Address)
	if !allowed {
		return ErrRateLimited
	}
	done(Failed) // every request is a failure (10.1 §8.1 item 1)
	var account pgtype.UUID
	err := database.Read(ctx, s.pool, database.Scope{Email: email}, func(tx pgx.Tx) error {
		var err error
		account, err = identitydb.New(tx).GetAccountIDByEmail(ctx, email)
		return err
	})
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil
	case err != nil:
		return fmt.Errorf("look up the email: %w", err)
	}
	token, tokenHash := newToken()
	now := s.config.Now()
	err = database.Write(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) error {
		return identitydb.New(tx).InsertRecoveryToken(ctx, identitydb.InsertRecoveryTokenParams{
			TokenHash: tokenHash, AccountID: account, CreatedAt: now, ExpiresAt: now.Add(recoveryLinkLifetime),
		})
	})
	if err != nil {
		return fmt.Errorf("store the recovery link: %w", err)
	}
	// After commit only: a retried transaction must not mail twice.
	return s.config.SendMail(ctx, Message{To: email, Kind: MessageRecovery, Link: s.config.Origin + "/recover#" + token})
}

// CompleteRecovery sets newPassword with the emailed recovery token,
// revokes every live session of the account, and opens a session for this
// device, in one transaction (SEC-24). A missing, unknown, expired or used
// token is ErrTokenRejected, found before any hash; failures count on the
// token-use address key (10.1 §8.1). Without a hash slot the token is kept
// (item 3).
func (s *Service) CompleteRecovery(ctx context.Context, client Client, recoveryToken, newPassword string) (Tokens, error) {
	if client.Kind == "web" {
		return Tokens{}, ErrWebSessionNotYet
	}
	if err := checkPasswordNotCommon(newPassword); err != nil {
		return Tokens{}, err
	}
	allowed, done := s.tokenUse.attempt("", client.Address)
	if !allowed {
		return Tokens{}, ErrRateLimited
	}
	tokens, err := s.completeRecovery(ctx, client, recoveryToken, newPassword)
	if err != nil {
		done(outcomeOf(err))
		return Tokens{}, err
	}
	done(Succeeded)
	return tokens, nil
}

func (s *Service) completeRecovery(ctx context.Context, client Client, recoveryToken, newPassword string) (Tokens, error) {
	tokenHash, ok := hashToken(recoveryToken)
	if !ok {
		return Tokens{}, ErrTokenRejected
	}
	var row identitydb.GetRecoveryTokenRow
	err := database.Read(ctx, s.pool, database.Scope{TokenHash: tokenHash}, func(tx pgx.Tx) error {
		var err error
		row, err = identitydb.New(tx).GetRecoveryToken(ctx, tokenHash)
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) || err == nil && (row.UsedAt.Valid || !s.config.Now().Before(row.ExpiresAt)) {
		return Tokens{}, ErrTokenRejected
	}
	if err != nil {
		return Tokens{}, fmt.Errorf("read the recovery link: %w", err)
	}
	// Hashed outside the transaction, which may run more than once.
	hash, err := s.hasher.hash(ctx, newPassword)
	if err != nil {
		return Tokens{}, fmt.Errorf("hash the password: %w", err)
	}
	account := row.AccountID
	var tokens Tokens
	// 10.1 §5.1: the write needs only the account; the lookup found it.
	err = database.Write(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) error {
		now := s.config.Now()
		q := identitydb.New(tx)
		used, err := q.UseRecoveryToken(ctx, identitydb.UseRecoveryTokenParams{Now: now, TokenHash: tokenHash})
		if err != nil || used == 0 {
			return errors.Join(err, rejectedIfNone(used))
		}
		if err := replaceCredential(ctx, q, account, hash, now); err != nil {
			return err
		}
		revoked, err := q.RevokeLiveSessions(ctx, identitydb.RevokeLiveSessionsParams{Now: now, Reason: "recovery", AccountID: account})
		if err != nil {
			return err
		}
		for _, session := range revoked {
			err := q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
				AccountID: account, SessionID: session, Kind: "session_revoked",
				Detail: pgtype.Text{String: "recovery", Valid: true}, OccurredAt: now,
			})
			if err != nil {
				return err
			}
		}
		var session pgtype.UUID
		if tokens, session, err = openSession(ctx, tx, account, client, now); err != nil {
			return err
		}
		// SEC-24: the owner reads when, and from which device, through the
		// new session.
		return q.InsertSecurityEvent(ctx, identitydb.InsertSecurityEventParams{
			AccountID: account, SessionID: session, Kind: "recovery", OccurredAt: now,
		})
	})
	return tokens, err
}
