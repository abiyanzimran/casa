package household

import (
	"bytes"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/household/householddb"
)

// FindInvitation finds the invitation of an invitation link by the token's
// hash, in a lookup transaction (Scope{TokenHash}; 09.2 §3): only one that
// is open and before its expiry, else ErrInvitationRejected, whatever the
// reason (SEC-30). It never writes the expiry (plan 1.6 decision 12).
func (s *Service) FindInvitation(ctx context.Context, tokenHash []byte) (household, invitation pgtype.UUID, err error) {
	var row householddb.OpenInvitationByTokenRow
	err = database.Read(ctx, s.pool, database.Scope{TokenHash: tokenHash}, func(tx pgx.Tx) (err error) {
		row, err = householddb.New(tx).OpenInvitationByToken(ctx, householddb.OpenInvitationByTokenParams{TokenHash: tokenHash, Now: s.now()})
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return household, invitation, ErrInvitationRejected
	}
	return row.HouseholdID, row.ID, err
}

// Preview is what the invitee reads before accepting (FAM.md
// FR-FAM-004/AC-1): the names only, no money data.
type Preview struct {
	HouseholdName, InviterDisplayName string
	// FormerMemberNotice is FR-FAM-004/AC-10, AC-11.
	FormerMemberNotice bool
}

// Preview reads the preview of the invitation FindInvitation found, in the
// household's own transaction; one no longer open, or past its expiry, is
// ErrInvitationRejected.
func (s *Service) Preview(ctx context.Context, household, invitation pgtype.UUID) (Preview, error) {
	var row householddb.InvitationPreviewRow
	err := database.Read(ctx, s.pool, database.Scope{HouseholdID: household}, func(tx pgx.Tx) (err error) {
		row, err = householddb.New(tx).InvitationPreview(ctx, householddb.InvitationPreviewParams{
			HouseholdID: household, ID: invitation, Now: s.now(),
		})
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Preview{}, ErrInvitationRejected
	}
	// ponytail: no former members until F06 and F07 (step 6, FR-FAM-004
	// AC-10, AC-11), so FormerMemberNotice stays false.
	return Preview{HouseholdName: row.HouseholdName, InviterDisplayName: row.InviterDisplayName.String}, err
}

// Accept is F02 by account (06 §3.6) in tx, whose scope is the household
// FindInvitation found and the account (plan 1.6 decision 13): the
// invitation is open, before its expiry, and still has tokenHash, else
// ErrInvitationRejected; the account is not a member of the household
// (06 §3.6 F02 guard), else ErrAlreadyMember; the household has fewer than
// two full adult members, else ErrInvitationRejected (FR-FAM-004/AC-8). It
// is then accepted by compare-and-set, its token cleared (SEC-28), with
// the event of the invitee and the adult age statement of version
// recordAge unless it is 0 (SEC-80). Another open acceptance of the account
// violates invitation_invitee: ErrAlreadyMember, never Internal (09.2 §5.1,
// note on decision 20).
func (s *Service) Accept(ctx context.Context, tx pgx.Tx, account, invitation pgtype.UUID, tokenHash []byte,
	displayName string, recordAge uint32, key pgtype.UUID,
) error {
	q := householddb.New(tx)
	now := s.now()
	row, err := q.Invitation(ctx, invitation)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return ErrInvitationRejected
	case err != nil:
		return err
	case row.State != "open" || !now.Before(row.ExpiresAt) || !bytes.Equal(row.TokenHash, tokenHash):
		return ErrInvitationRejected
	}
	if _, err := s.MemberIn(ctx, tx, account); !errors.Is(err, ErrNotFound) {
		if err == nil {
			err = ErrAlreadyMember
		}
		return err
	}
	adults, err := s.FullAdultMembers(ctx, tx)
	if err != nil {
		return err
	}
	if len(adults) >= maxFullAdults {
		return ErrInvitationRejected
	}
	n, err := q.AcceptInvitation(ctx, householddb.AcceptInvitationParams{
		InviteeAccountID: account, InviteeDisplayName: optionalText(displayName),
		ID: invitation, Version: row.Version, TokenHash: tokenHash,
	})
	var pgErr *pgconn.PgError
	switch {
	case errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "invitation_invitee":
		return ErrAlreadyMember
	case err != nil:
		return err
	case n == 0:
		return ErrInvitationRejected
	}
	if err := RecordEvent(ctx, tx, Event{
		HouseholdID: row.HouseholdID, ObjectKind: "invitation", ObjectID: invitation, Seq: row.Version + 1,
		ActorKind: "invitee", Command: acceptCommand, FromState: row.State, ToState: "accepted",
		IdempotencyKey: key, At: now,
	}); err != nil {
		return err
	}
	return recordAgeDeclaration(ctx, q, account, recordAge, now)
}

// Acceptance is an account's own open acceptance: its household and
// invitation, and whether it is still live.
type Acceptance struct {
	HouseholdID, InvitationID pgtype.UUID
	Live                      bool
}

// FindAcceptance reads account's own accepted invitation in a lookup
// transaction (Scope{AccountID}, the invitee clause; plan 1.6 decision 14),
// then hands it to expire, the token-free path's expiry writer (decision
// 12): a due one is written expired and returned with Live false. One that
// another command closed in between is returned live, so its caller
// refuses, and the next call no longer finds it. ErrNotFound when there is
// none.
func (s *Service) FindAcceptance(ctx context.Context, account pgtype.UUID) (Acceptance, error) {
	var row householddb.OwnAcceptanceRow
	err := database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) (err error) {
		row, err = householddb.New(tx).OwnAcceptance(ctx, account)
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Acceptance{}, ErrNotFound
	}
	if err != nil {
		return Acceptance{}, err
	}
	expired, err := s.expire(ctx, row.HouseholdID, account, row.ID)
	return Acceptance{HouseholdID: row.HouseholdID, InvitationID: row.ID, Live: !expired}, err
}

// Withdraw is F05 withdraw_acceptance by account (06 §3.6) in tx, whose
// scope is the household FindAcceptance found: only the invitee, else
// ErrNotFound (plan 1.6 decision 14).
func (s *Service) Withdraw(ctx context.Context, tx pgx.Tx, account, invitation pgtype.UUID, key pgtype.UUID) error {
	row, err := householddb.New(tx).Invitation(ctx, invitation)
	switch {
	case errors.Is(err, pgx.ErrNoRows) || err == nil && row.InviteeAccountID != account:
		return ErrNotFound
	case err != nil:
		return err
	}
	return s.withdraw(ctx, tx, invitation, row, key)
}

// withdraw is the one writer of a withdrawal (plan 1.6 decision 14; 10.1
// §10 note): the invitation read in row, accepted and before its expiry,
// else ErrStateChanged, becomes withdrawn by compare-and-set on its version
// and state, closed now, with the event of the invitee. The version pins the
// invitee too: invitee_account_id changes only with F02's version bump.
func (s *Service) withdraw(ctx context.Context, tx pgx.Tx, invitation pgtype.UUID, row householddb.InvitationRow, key pgtype.UUID) error {
	if row.State != "accepted" || !s.now().Before(row.ExpiresAt) {
		return ErrStateChanged
	}
	return s.transition(ctx, tx, row.HouseholdID, invitation, row, "withdrawn", s.now(), Event{
		ActorKind: "invitee", Command: withdrawCommand, IdempotencyKey: key,
	})
}
