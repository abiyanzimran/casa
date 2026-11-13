package household

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/household/householddb"
)

// invitationLifetime is SEC-28's 7 days [Asumsi], counted from F01 and never
// extended by a new link (10.1 SEC-28 note Q7 (b)).
const invitationLifetime = 7 * 24 * time.Hour

// The object_event values of the invitation transitions: the snake_case
// commands of 06 §3.6 and English state names (09.2 §5.1, note of 14
// November 2026). The time transition of F03 and F05 has no command in 06;
// expireCommand names it [Usulan], with the rule that fired.
const (
	inviteCommand   = "invite"
	acceptCommand   = "accept_invitation"
	cancelCommand   = "cancel_invitation"
	withdrawCommand = "withdraw_acceptance"
	expireCommand   = "expire_invitation"
	expireRule      = "invitation.expired"
)

var (
	// ErrInvitationPending: the household already has an invitation that
	// has not lapsed (06 §3.6 F01; FR-FAM-003/AC-2).
	ErrInvitationPending = errors.New("an invitation is pending")
	// ErrHouseholdFull: the household has two full adult members (06 §3.6
	// F01; FR-FAM-003/AC-5).
	ErrHouseholdFull = errors.New("household full")
)

// Invitation is a made invitation: its id and the end of its validity.
type Invitation struct {
	ID        pgtype.UUID
	ExpiresAt time.Time
}

// ExpireDueInvitation writes the expiry of the household's due live
// invitation, if any, in its own database.Write (plan 1.6 decision 12): the
// token-free paths call it before their command, so the command finds the
// invitation no longer open (06 §3.6 "Gugur karena waktu", change X3).
func (s *Service) ExpireDueInvitation(ctx context.Context, inviter Member) error {
	var found householddb.LiveInvitationRow
	err := database.Read(ctx, s.pool, database.Scope{HouseholdID: inviter.HouseholdID, AccountID: inviter.AccountID}, func(tx pgx.Tx) (err error) {
		found, err = householddb.New(tx).LiveInvitation(ctx, inviter.HouseholdID)
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}
	_, err = s.expire(ctx, found.HouseholdID, inviter.AccountID, found.ID)
	return err
}

// expire is the one expiry writer (plan 1.6 decision 12; 09.2 §5.1 note X3).
// In its own database.Write it reads the version, state and expires_at, then
// compare-and-sets an open or accepted invitation with now ≥ expires_at to
// expired: actor system, rule invitation.expired, closed_at and the event's
// time = expires_at, so the household never reads when it was written (10.1
// §10 note X3 (1)(a)). expired is true when the row is (now) expired, false
// when it is live or closed otherwise; a lost compare-and-set writes
// nothing and reports the row as read, not expired. Under SERIALIZABLE a
// concurrent write is a serialization failure instead, and the retry reads
// the row again.
func (s *Service) expire(ctx context.Context, household, account, invitation pgtype.UUID) (expired bool, err error) {
	err = database.Write(ctx, s.pool, database.Scope{HouseholdID: household, AccountID: account}, func(tx pgx.Tx) error {
		row, err := householddb.New(tx).Invitation(ctx, invitation)
		if err != nil {
			return err
		}
		if !live(row.State) || s.now().Before(row.ExpiresAt) {
			expired = row.State == "expired"
			return nil
		}
		err = s.transition(ctx, tx, household, invitation, row, "expired", row.ExpiresAt, Event{
			ActorKind: "system", Rule: expireRule, Command: expireCommand,
		})
		expired = err == nil
		if errors.Is(err, ErrStateChanged) {
			return nil
		}
		return err
	})
	return expired, err
}

// live is a state in which an invitation is not yet closed (06 §3.6:
// undangan_terbuka, menunggu_konfirmasi_pengundang).
func live(state string) bool { return state == "open" || state == "accepted" }

// transition moves invitation from the state read in row to to, by
// compare-and-set on its version and state (plan 1.6 decision 11), and
// records event with the transition's seq, states and object, at closedAt.
// A lost compare-and-set is ErrStateChanged and writes nothing.
func (s *Service) transition(ctx context.Context, tx pgx.Tx, household, invitation pgtype.UUID,
	row householddb.InvitationRow, to string, closedAt time.Time, event Event,
) error {
	n, err := householddb.New(tx).TransitionInvitation(ctx, householddb.TransitionInvitationParams{
		ToState: to, ClosedAt: closedAt, HouseholdID: household, ID: invitation, Version: row.Version, FromState: row.State,
	})
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrStateChanged
	}
	event.HouseholdID, event.ObjectKind, event.ObjectID = household, "invitation", invitation
	event.Seq, event.FromState, event.ToState, event.At = row.Version+1, row.State, to, closedAt
	return RecordEvent(ctx, tx, event)
}

// CreateInvitation is F01 (06 §3.6) in tx, whose scope is inviter's
// household: an open invitation without a token, valid invitationLifetime,
// and its creation event (seq 0). Guards: fewer than two full adult
// members, else ErrHouseholdFull; no other live invitation, held by
// invitation_live, whose violation is ErrInvitationPending, never Internal
// (ADR-013 note 1.6). The third guard, no member in keluar_tenggang, holds
// by itself until F06 adds membership.state (FR-FAM-003/AC-7, step 6; plan
// 1.6 decision 1).
func (s *Service) CreateInvitation(ctx context.Context, tx pgx.Tx, inviter Member, key pgtype.UUID) (Invitation, error) {
	adults, err := s.FullAdultMembers(ctx, tx)
	if err != nil {
		return Invitation{}, err
	}
	if len(adults) >= maxFullAdults {
		return Invitation{}, ErrHouseholdFull
	}
	now := s.now()
	inv := Invitation{ExpiresAt: now.Add(invitationLifetime)}
	inv.ID, err = householddb.New(tx).InsertInvitation(ctx, householddb.InsertInvitationParams{
		HouseholdID: inviter.HouseholdID, Inviter: inviter.MemberID, CreatedAt: now, ExpiresAt: inv.ExpiresAt,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "invitation_live" {
			return Invitation{}, ErrInvitationPending
		}
		return Invitation{}, err
	}
	return inv, RecordEvent(ctx, tx, Event{
		HouseholdID: inviter.HouseholdID, ObjectKind: "invitation", ObjectID: inv.ID, Seq: 0,
		ActorKind: "member", Actor: inviter.MemberID, Command: inviteCommand, ToState: "open",
		IdempotencyKey: key, At: now,
	})
}

// openInvitation reads invitation for a command of inviter that expects it
// open and not past its expiry (06 S-4; plan 1.6 decision 11): another
// household's invitation, or one inviter did not make, is ErrNotFound
// (10.1 SEC-28 note Q7 (a); FR-FAM-003 note of 14 November 2026); any other
// state, or a due one, ErrStateChanged.
func (s *Service) openInvitation(ctx context.Context, tx pgx.Tx, inviter Member, invitation pgtype.UUID) (householddb.InvitationRow, error) {
	row, err := householddb.New(tx).Invitation(ctx, invitation)
	switch {
	case errors.Is(err, pgx.ErrNoRows) || err == nil && row.Inviter != inviter.MemberID:
		return row, ErrNotFound
	case err != nil:
		return row, err
	case row.State != "open" || !s.now().Before(row.ExpiresAt):
		return row, ErrStateChanged
	}
	return row, nil
}

// IssueInvitationLink sets tokenHash as the link of inviter's open
// invitation, in its own database.Write, replacing the previous hash so the
// old link stops working (10.1 SEC-28 note Q7). It is not a transition: no
// version bump and no event (plan 1.6 decision 11). The membership is
// checked again inside it (SEC-151). expiresAt is the expiry set by F01.
func (s *Service) IssueInvitationLink(ctx context.Context, inviter Member, invitation pgtype.UUID, tokenHash []byte) (expiresAt time.Time, err error) {
	err = database.Write(ctx, s.pool, database.Scope{HouseholdID: inviter.HouseholdID, AccountID: inviter.AccountID}, func(tx pgx.Tx) error {
		member, err := s.MemberIn(ctx, tx, inviter.AccountID)
		if err != nil {
			return err
		}
		row, err := s.openInvitation(ctx, tx, member, invitation)
		if err != nil {
			return err
		}
		n, err := householddb.New(tx).ReplaceInvitationToken(ctx, householddb.ReplaceInvitationTokenParams{
			TokenHash: tokenHash, HouseholdID: inviter.HouseholdID, ID: invitation, Version: row.Version,
		})
		if err == nil && n == 0 {
			err = ErrStateChanged
		}
		expiresAt = row.ExpiresAt
		return err
	})
	return expiresAt, err
}

// CancelInvitation is F03 by inviter (06 §3.6) in tx: only the inviter, only
// an open invitation (FR-FAM-003 note of 14 November 2026); cancelled,
// closed now, with its event.
func (s *Service) CancelInvitation(ctx context.Context, tx pgx.Tx, inviter Member, invitation pgtype.UUID, key pgtype.UUID) error {
	row, err := s.openInvitation(ctx, tx, inviter, invitation)
	if err != nil {
		return err
	}
	return s.transition(ctx, tx, inviter.HouseholdID, invitation, row, "cancelled", s.now(), Event{
		ActorKind: "member", Actor: inviter.MemberID, Command: cancelCommand, IdempotencyKey: key,
	})
}
