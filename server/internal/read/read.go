// Package read is the one read path of the server (09.2 §1 row read, rule
// 2; SEC-06): the screen projections and the one cross-account read,
// AccountEmail (plan 1.6 Q10). It imports only database and its own query
// package (09.2 §1 item 10), so it never imports household or identity.
package read

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/read/readdb"
)

// ObjectKind is a kind of object the feed carries (09.3 §3; plan 1.7
// decision 7): reader_feed.object_kind.
type ObjectKind string

// The kinds of 1.7 (plan 1.6 Q12).
const (
	KindHousehold  ObjectKind = "household"
	KindMember     ObjectKind = "member"
	KindInvitation ObjectKind = "invitation"
)

// ObjectRef names one object of a household.
type ObjectRef struct {
	Kind ObjectKind
	ID   pgtype.UUID
}

// maxMembers bounds the member queries [Usulan] (plan 1.6 "Batasan global").
const maxMembers = 20

// errTooManyMembers: more members than the bounded query lists, so a reader
// would be left out (ES §3).
var errTooManyMembers = errors.New("more household members than the read path lists")

// canRead is the one read rule (plan 1.7 decision 7, ruling P5): whether a
// member of the household reads an object of kind. Today every member reads
// the household, its members and its live invitation (06 §3.1; plan 1.6
// decision 16); step 2 widens it here, per kind.
func canRead(kind ObjectKind) bool {
	switch kind {
	case KindHousehold, KindMember, KindInvitation:
		return true
	}
	return false
}

// Members are household's member ids in tx, by id: the member objects that
// change with a household event (plan 1.7 decision 8; 09.2 §5.1).
func Members(ctx context.Context, tx pgx.Tx, household pgtype.UUID) ([]pgtype.UUID, error) {
	members, err := readdb.New(tx).HouseholdMemberIDs(ctx, household)
	if err == nil && len(members) > maxMembers {
		return nil, errTooManyMembers
	}
	return members, err
}

// Readers are the members who can read an object of kind in household after
// the change, in tx, by id: today every member, for every kind (plan 1.7
// decision 7; 09.3 §3 "Fan-out saat menulis").
func Readers(ctx context.Context, tx pgx.Tx, household pgtype.UUID, kind ObjectKind) ([]pgtype.UUID, error) {
	if !canRead(kind) {
		return nil, nil
	}
	return Members(ctx, tx, household)
}

// ErrNotFound: not a membership of the account, whether the household
// exists or not; one answer for all (SEC-05; FR-FAM-020/AC-1, AC-4).
var ErrNotFound = errors.New("not found")

// Service is the one read path (09.2 §1 rule 2; SEC-06).
type Service struct {
	pool *pgxpool.Pool
	now  func() time.Time
}

// New returns the read service over pool, reading time from now.
func New(pool *pgxpool.Pool, now func() time.Time) *Service {
	return &Service{pool: pool, now: now}
}

// Reader is the member a read is for, its membership already checked by the
// caller (09.2 §3 layer 1).
type Reader struct{ HouseholdID, AccountID, MemberID pgtype.UUID }

// Member is a listed member: what other members see of it.
type Member struct {
	MemberID    pgtype.UUID
	DisplayName string
}

// Membership is one household of an account, with the account's member id
// and the household's members.
type Membership struct {
	HouseholdID   pgtype.UUID
	HouseholdName string
	MemberID      pgtype.UUID
	Members       []Member
}

// Memberships lists account's households, at most 20, each with at most 20
// members [Usulan] (plan 1.6 "Batasan global"). The households are found in
// a lookup transaction; each is then read in its own scope, with the
// membership checked again (09.2 §3 layer 1; 10.1 §5.1 note 1.6 item 1).
func (s *Service) Memberships(ctx context.Context, account pgtype.UUID) ([]Membership, error) {
	var found []readdb.MembershipsOfAccountRow
	err := database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) (err error) {
		found, err = readdb.New(tx).MembershipsOfAccount(ctx, account)
		return err
	})
	if err != nil {
		return nil, err
	}
	// ponytail: 1+N reads. membership_one_household limits it to one
	// iteration until F-FAM-009; RLS `id = H` forces one tx per household.
	memberships := make([]Membership, 0, len(found))
	for _, f := range found {
		var rows []readdb.HouseholdMembersRow
		err := database.Read(ctx, s.pool, database.Scope{HouseholdID: f.HouseholdID, AccountID: account}, func(tx pgx.Tx) (err error) {
			rows, err = readdb.New(tx).HouseholdMembers(ctx, readdb.HouseholdMembersParams{
				HouseholdID: f.HouseholdID, AccountID: account,
			})
			return err
		})
		if err != nil {
			return nil, err
		}
		if len(rows) == 0 { // the membership ended between the two reads
			continue
		}
		m := Membership{HouseholdID: f.HouseholdID, HouseholdName: rows[0].HouseholdName, MemberID: f.MemberID}
		for _, r := range rows {
			m.Members = append(m.Members, Member{MemberID: r.MemberID, DisplayName: r.DisplayName.String})
		}
		memberships = append(memberships, m)
	}
	return memberships, nil
}

// OpenInvitation is the household's live invitation as a member reads it
// (plan 1.6 decision 16).
type OpenInvitation struct {
	ID        pgtype.UUID
	State     string
	ExpiresAt time.Time
	Inviter   pgtype.UUID
	// InviteeDisplayName and InviteeAccount are set only for the inviter of
	// an accepted invitation (decision 16; 10.1 §5.1 note 1.6 item 3 (a),
	// (b)).
	InviteeDisplayName string
	InviteeAccount     pgtype.UUID
}

// OpenInvitation reads r's household's live invitation in the household's
// own read transaction, with the membership checked again (SEC-151);
// ErrNotFound without it. One past its expiry is lapsed: found is false and
// nothing is written (plan 1.6 decision 12).
func (s *Service) OpenInvitation(ctx context.Context, r Reader) (inv OpenInvitation, found bool, err error) {
	err = database.Read(ctx, s.pool, database.Scope{HouseholdID: r.HouseholdID, AccountID: r.AccountID}, func(tx pgx.Tx) error {
		q := readdb.New(tx)
		member, err := q.ReaderMember(ctx, r.AccountID)
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		if err != nil {
			return err
		}
		row, err := q.LiveInvitationView(ctx)
		if errors.Is(err, pgx.ErrNoRows) || err == nil && !s.now().Before(row.ExpiresAt) {
			return nil
		}
		if err != nil {
			return err
		}
		inv, found = OpenInvitation{ID: row.ID, State: row.State, ExpiresAt: row.ExpiresAt, Inviter: row.Inviter}, true
		if row.State == "accepted" && row.Inviter == member {
			inv.InviteeDisplayName, inv.InviteeAccount = row.InviteeDisplayName.String, row.InviteeAccountID
		}
		return nil
	})
	return inv, found, err
}

// AccountEmail is account's email, read in the account's own scope: the one
// cross-account read of the server (10.1 §5.1 note 1.6 item 3; plan 1.6
// Q10). Its only caller is the GetOpenInvitation handler, for the inviter of
// an accepted invitation, with the invitee's account taken from that
// invitation's row (SEC-31); a second caller needs a review of 10.1.
func (s *Service) AccountEmail(ctx context.Context, account pgtype.UUID) (string, error) {
	var email string
	err := database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) (err error) {
		email, err = readdb.New(tx).AccountEmail(ctx, account)
		return err
	})
	return email, err
}
