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
	confirmCommand  = "confirm_invitee"
	rejectCommand   = "reject_invitee"
	expireCommand   = "expire_invitation"
	expireRule      = "invitation.expired"
	// joinedState is the to_state of the household's own event when F04
	// adds a member [Usulan]: the household has no state machine, so its
	// event names what happened, as 'created' does.
	joinedState = "member_joined"
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
// A lost compare-and-set is ErrStateChanged and writes nothing. Only a
// withdrawal clears the invitee's account and display name
// (FR-FAM-004/AC-5); a rejection or an expiry keeps them, and the event
// keeps its actor.
func (s *Service) transition(ctx context.Context, tx pgx.Tx, household, invitation pgtype.UUID,
	row householddb.InvitationRow, to string, closedAt time.Time, event Event,
) error {
	n, err := householddb.New(tx).TransitionInvitation(ctx, householddb.TransitionInvitationParams{
		ToState: to, ClosedAt: closedAt, HouseholdID: household, ID: invitation, Version: row.Version, FromState: row.State,
		ClearInvitee: to == "withdrawn",
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

// inviterInvitation reads invitation for a command of inviter that expects
// it in state and not past its expiry (06 S-4; plan 1.6 decision 11):
// another household's invitation, or one inviter did not make, is
// ErrNotFound (10.1 SEC-28 note Q7 (a); FR-FAM-003 note of 14 November
// 2026; decision 14); any other state, or a due one, ErrStateChanged.
func (s *Service) inviterInvitation(ctx context.Context, tx pgx.Tx, inviter Member, invitation pgtype.UUID, state string) (householddb.InvitationRow, error) {
	row, err := householddb.New(tx).Invitation(ctx, invitation)
	switch {
	case errors.Is(err, pgx.ErrNoRows) || err == nil && row.Inviter != inviter.MemberID:
		return row, ErrNotFound
	case err != nil:
		return row, err
	case row.State != state || !s.now().Before(row.ExpiresAt):
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
		row, err := s.inviterInvitation(ctx, tx, member, invitation, "open")
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
	row, err := s.inviterInvitation(ctx, tx, inviter, invitation, "open")
	if err != nil {
		return err
	}
	return s.transition(ctx, tx, inviter.HouseholdID, invitation, row, "cancelled", s.now(), Event{
		ActorKind: "member", Actor: inviter.MemberID, Command: cancelCommand, IdempotencyKey: key,
	})
}

// ErrInviteeUnavailable: the invitee became a member of another household
// after accepting, so F04 cannot add it (09.2 §5.1, note on decision 20).
// Only the API layer's F04 handler sees it, and closes the invitation with
// CloseAsWithdrawn; it is never mapped to a contract code (plan 1.6
// decision 14).
var ErrInviteeUnavailable = errors.New("invitee is a member elsewhere")

// Confirm is F04 confirm_invitee by inviter (06 §3.6) in tx, whose scope is
// inviter's household: only the inviter, only an accepted invitation before
// its expiry (inviterInvitation). F04 never makes a third full adult (06
// §3.6 F01 keeps a household at two), so a household that already has two
// is ErrStateChanged [Usulan]. The invitation becomes confirmed with the
// inviter's event; the invitee becomes a member with the inviter's role,
// permission set and approval authority (06 §3.1; FR-FAM-004/AC-4); the
// household records member_joined with the new member as its actor, and the
// inviter is the actor of the invitation's confirmed event in the same
// transaction, so the joining has both its actors (09.2 §5.1). Confirm and Create are the only writers of
// membership (INV-A48 (a), (b)). An invitee who is a member anywhere
// inserts nothing: ErrInviteeUnavailable, and the caller rolls tx back.
func (s *Service) Confirm(ctx context.Context, tx pgx.Tx, inviter Member, invitation pgtype.UUID, key pgtype.UUID) (Member, error) {
	row, err := s.inviterInvitation(ctx, tx, inviter, invitation, "accepted")
	if err != nil {
		return Member{}, err
	}
	adults, err := s.FullAdultMembers(ctx, tx)
	if err != nil {
		return Member{}, err
	}
	if len(adults) >= maxFullAdults {
		return Member{}, ErrStateChanged
	}
	now := s.now()
	if err := s.transition(ctx, tx, inviter.HouseholdID, invitation, row, "confirmed", now, Event{
		ActorKind: "member", Actor: inviter.MemberID, Command: confirmCommand, IdempotencyKey: key,
	}); err != nil {
		return Member{}, err
	}
	q := householddb.New(tx)
	memberID, err := q.JoinMembership(ctx, householddb.JoinMembershipParams{
		AccountID: row.InviteeAccountID, DisplayName: row.InviteeDisplayName, JoinedAt: now,
		HouseholdID: inviter.HouseholdID, Inviter: inviter.MemberID,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Member{}, ErrInviteeUnavailable
	}
	if err != nil {
		return Member{}, err
	}
	if err := q.CopyApprovalAuthority(ctx, householddb.CopyApprovalAuthorityParams{
		MemberID: memberID, ValidFrom: now, HouseholdID: inviter.HouseholdID, Inviter: inviter.MemberID,
	}); err != nil {
		return Member{}, err
	}
	version, err := q.NextHouseholdVersion(ctx, inviter.HouseholdID)
	if err != nil {
		return Member{}, err
	}
	return Member{HouseholdID: inviter.HouseholdID, MemberID: memberID, AccountID: row.InviteeAccountID}, RecordEvent(ctx, tx, Event{
		HouseholdID: inviter.HouseholdID, ObjectKind: "household", ObjectID: inviter.HouseholdID, Seq: version,
		ActorKind: "member", Actor: memberID, Command: confirmCommand, ToState: joinedState,
		IdempotencyKey: key, At: now,
	})
}

// Reject is F05 reject_invitee by inviter (06 §3.6; FR-FAM-004/AC-3) in tx:
// only the inviter, only an accepted invitation before its expiry; rejected,
// closed now, with its event. The invitee's account and display name stay:
// only a withdrawal clears them (FR-FAM-004/AC-5).
func (s *Service) Reject(ctx context.Context, tx pgx.Tx, inviter Member, invitation pgtype.UUID, key pgtype.UUID) error {
	row, err := s.inviterInvitation(ctx, tx, inviter, invitation, "accepted")
	if err != nil {
		return err
	}
	return s.transition(ctx, tx, inviter.HouseholdID, invitation, row, "rejected", s.now(), Event{
		ActorKind: "member", Actor: inviter.MemberID, Command: rejectCommand, IdempotencyKey: key,
	})
}

// OpenInvitation is the household's live invitation as a member reads it
// (plan 1.6 decision 16).
type OpenInvitation struct {
	ID        pgtype.UUID
	State     string
	ExpiresAt time.Time
	// InviteeDisplayName and InviteeAccount are set only for the inviter of
	// an accepted invitation (decision 16; 10.1 §5.1 note 1.6 item 3 (a),
	// (b)).
	InviteeDisplayName string
	InviteeAccount     pgtype.UUID
}

// OpenInvitation reads viewer's household's live invitation in the
// household's own read transaction, with the membership checked again
// (SEC-151). One past its expiry is lapsed: found is false and nothing is
// written (decision 12).
func (s *Service) OpenInvitation(ctx context.Context, viewer Member) (inv OpenInvitation, found bool, err error) {
	err = database.Read(ctx, s.pool, database.Scope{HouseholdID: viewer.HouseholdID, AccountID: viewer.AccountID}, func(tx pgx.Tx) error {
		member, err := s.MemberIn(ctx, tx, viewer.AccountID)
		if err != nil {
			return err
		}
		row, err := householddb.New(tx).LiveInvitationView(ctx)
		if errors.Is(err, pgx.ErrNoRows) || err == nil && !s.now().Before(row.ExpiresAt) {
			return nil
		}
		if err != nil {
			return err
		}
		inv, found = OpenInvitation{ID: row.ID, State: row.State, ExpiresAt: row.ExpiresAt}, true
		if row.State == "accepted" && row.Inviter == member.MemberID {
			inv.InviteeDisplayName, inv.InviteeAccount = row.InviteeDisplayName.String, row.InviteeAccountID
		}
		return nil
	})
	return inv, found, err
}
