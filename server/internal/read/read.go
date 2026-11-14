// Package read is the one read path of the server (09.2 §1 row read, rule
// 2; SEC-06): the screen projections and the one cross-account read,
// AccountEmail (plan 1.6 Q10). It imports only database and its own query
// package (09.2 §1 item 10), so it never imports household or identity.
package read

import (
	"bytes"
	"context"
	"errors"
	"slices"
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

// errNoMembers: a household always has a member (06 §3.1), so none means a
// household other than the transaction's, which row-level security hides;
// the fan-out fails instead of writing nothing.
var errNoMembers = errors.New("no members of the household in this transaction")

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
// change with a household event (plan 1.7 decision 8; 09.2 §5.1). None is
// errNoMembers, more than maxMembers errTooManyMembers.
func Members(ctx context.Context, tx pgx.Tx, household pgtype.UUID) ([]pgtype.UUID, error) {
	members, err := readdb.New(tx).HouseholdMemberIDs(ctx, household)
	switch {
	case err != nil:
		return nil, err
	case len(members) == 0:
		return nil, errNoMembers
	case len(members) > maxMembers:
		return nil, errTooManyMembers
	}
	return members, nil
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
		row, live, err := s.liveInvitation(ctx, q)
		if err != nil || !live {
			return err
		}
		inv, found = OpenInvitation{ID: row.ID, State: row.State, ExpiresAt: row.ExpiresAt}, true
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

// OwnAcceptance is the caller's own acceptance as GetOwnAcceptance shows it:
// no household data (SEC-32).
type OwnAcceptance struct {
	DisplayName string
	ExpiresAt   time.Time
}

// OwnAcceptance reads account's accepted invitation in a lookup transaction,
// found from the account, never from the request (plan 1.6 decision 14).
// One past its expiry is lapsed: found is false and nothing is written, the
// same rule as liveInvitation (09.3 §1 note of 14 November 2026).
func (s *Service) OwnAcceptance(ctx context.Context, account pgtype.UUID) (a OwnAcceptance, found bool, err error) {
	var row readdb.OwnAcceptanceViewRow
	err = database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) (err error) {
		row, err = readdb.New(tx).OwnAcceptanceView(ctx, account)
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) || err == nil && !s.now().Before(row.ExpiresAt) {
		return OwnAcceptance{}, false, nil
	}
	if err != nil {
		return OwnAcceptance{}, false, err
	}
	return OwnAcceptance{DisplayName: row.InviteeDisplayName.String, ExpiresAt: row.ExpiresAt}, true, nil
}

// liveInvitation is the household's live invitation in q's transaction: open
// or accepted, and now before expires_at, whether its expiry is written yet
// or not (plan 1.6 decision 12; plan 1.7 decision 7). It is the one rule of
// GetOpenInvitation and of sync.
func (s *Service) liveInvitation(ctx context.Context, q *readdb.Queries) (readdb.LiveInvitationViewRow, bool, error) {
	row, err := q.LiveInvitationView(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return row, false, nil
	}
	return row, err == nil && s.now().Before(row.ExpiresAt), err
}

// CheckReader is ErrNotFound unless r's account holds the membership
// r.MemberID of tx's household, checked again inside the read transaction
// (SEC-151; ruling P5: the read path's own layer).
func CheckReader(ctx context.Context, tx pgx.Tx, r Reader) error {
	member, err := readdb.New(tx).ReaderMember(ctx, r.AccountID)
	if errors.Is(err, pgx.ErrNoRows) || err == nil && member != r.MemberID {
		return ErrNotFound
	}
	return err
}

// Household is the household object of sync (plan 1.7 decision 7).
type Household struct {
	ID       pgtype.UUID
	Name     string
	TimeZone string
	Currency string
}

// Invitation is the invitation object of sync: its state only, never the
// invitee's name or email (plan 1.6 Q12; plan 1.7 decision 7).
type Invitation struct {
	ID        pgtype.UUID
	State     string // open or accepted
	ExpiresAt time.Time
	Inviter   pgtype.UUID // member_id
}

// Object is one projection; exactly one field is set.
type Object struct {
	Household  *Household
	Member     *Member
	Invitation *Invitation
}

// snapshotOrder is the order of kinds in a snapshot (plan 1.7 decision 10).
var snapshotOrder = []ObjectKind{KindHousehold, KindMember, KindInvitation}

// nilID is the uuid before every other, where a kind's page starts.
var nilID = pgtype.UUID{Valid: true}

// kindPage is up to n live objects of kind in tx's household with an id
// after after, by id; errTooManyMembers when more than maxMembers members
// follow after, so no member is ever dropped from a page. A household has one household object and at most one
// live invitation (invitation_live), so only members need a keyset page.
func (s *Service) kindPage(ctx context.Context, q *readdb.Queries, kind ObjectKind, after pgtype.UUID, n int) ([]Projected, error) {
	var out []Projected
	switch kind {
	case KindHousehold:
		h, err := q.HouseholdView(ctx)
		if err != nil {
			return nil, err
		}
		out = append(out, Projected{ObjectRef{kind, h.ID}, Object{Household: &Household{ID: h.ID, Name: h.Name, TimeZone: h.TimeZone, Currency: h.Currency}}})
	case KindMember:
		// Always one more than maxMembers, whatever the page: more members
		// fail the read instead of dropping one (ES §3).
		rows, err := q.MemberViews(ctx, readdb.MemberViewsParams{After: after, Lim: maxMembers + 1})
		switch {
		case err != nil:
			return nil, err
		case len(rows) > maxMembers:
			return nil, errTooManyMembers
		}
		for _, m := range rows {
			out = append(out, Projected{ObjectRef{kind, m.MemberID}, Object{Member: &Member{MemberID: m.MemberID, DisplayName: m.DisplayName.String}}})
		}
	case KindInvitation:
		row, live, err := s.liveInvitation(ctx, q)
		if err != nil {
			return nil, err
		}
		if live {
			out = append(out, Projected{ObjectRef{kind, row.ID}, Object{Invitation: &Invitation{ID: row.ID, State: row.State, ExpiresAt: row.ExpiresAt, Inviter: row.Inviter}}})
		}
	}
	out = slices.DeleteFunc(out, func(p Projected) bool { return bytes.Compare(p.Ref.ID.Bytes[:], after.Bytes[:]) <= 0 })
	return out[:min(n, len(out))], nil
}

// Objects projects refs for r in tx, a read transaction of r's household; a
// ref that r cannot read (canRead) or that is not live (plan 1.7 decision 7)
// is absent from the map. Today the rule reads by kind alone (canRead); r is
// what step 2's rule reads.
func (s *Service) Objects(ctx context.Context, tx pgx.Tx, r Reader, refs []ObjectRef) (map[ObjectRef]Object, error) {
	q := readdb.New(tx)
	out := map[ObjectRef]Object{}
	for _, kind := range snapshotOrder {
		if !canRead(kind) || !slices.ContainsFunc(refs, func(ref ObjectRef) bool { return ref.Kind == kind }) {
			continue
		}
		all, err := s.kindPage(ctx, q, kind, nilID, maxMembers+1)
		if err != nil {
			return nil, err
		}
		for _, p := range all {
			if slices.Contains(refs, p.Ref) {
				out[p.Ref] = p.Object
			}
		}
	}
	return out, nil
}

// SnapshotMark is a position in the snapshot order (kind, then id); its zero
// value is the start.
type SnapshotMark struct {
	Kind  ObjectKind
	After pgtype.UUID
}

// Projected is one object of a snapshot page.
type Projected struct {
	Ref    ObjectRef
	Object Object
}

// errUnknownMark: a mark of no kind of the snapshot; the cursor codec makes
// none, so it is a programming error.
var errUnknownMark = errors.New("snapshot mark of an unknown kind")

// SnapshotPage returns up to limit live objects r can read after mark, in
// order of kind, then id (ADR-017 item 7; plan 1.7 decision 10), from tx, a
// read transaction of r's household; done when none is left after them, and
// next is the mark of the last one returned. Its order is by id, never time,
// so it does not tell how many objects r cannot read (SEC-04).
func (s *Service) SnapshotPage(ctx context.Context, tx pgx.Tx, r Reader, mark SnapshotMark, limit int) (objects []Projected, next SnapshotMark, done bool, err error) {
	if mark.Kind == "" {
		mark = SnapshotMark{Kind: snapshotOrder[0], After: nilID}
	}
	start := slices.Index(snapshotOrder, mark.Kind)
	if start < 0 {
		return nil, SnapshotMark{}, false, errUnknownMark
	}
	q := readdb.New(tx)
	// One object more than the page tells whether any is left.
	for i, kind := range snapshotOrder[start:] {
		need := limit + 1 - len(objects)
		if need <= 0 {
			break
		}
		if !canRead(kind) {
			continue
		}
		after := nilID
		if i == 0 {
			after = pgtype.UUID{Bytes: mark.After.Bytes, Valid: true}
		}
		page, err := s.kindPage(ctx, q, kind, after, need)
		if err != nil {
			return nil, SnapshotMark{}, false, err
		}
		objects = append(objects, page...)
	}
	done = len(objects) <= limit
	objects = objects[:min(limit, len(objects))]
	next = mark
	if len(objects) > 0 {
		last := objects[len(objects)-1].Ref
		next = SnapshotMark{Kind: last.Kind, After: last.ID}
	}
	return objects, next, done, nil
}
