package household

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/household/householddb"
)

// maxFullAdults is the pilot preset's limit: F01 keeps a household at two
// full adult members (06 §3.6 F01).
const maxFullAdults = 2

// errTooManyFullAdults: more full adult members than F01 ever lets in, so
// the bounded query cannot list them all (ES §3).
var errTooManyFullAdults = errors.New("more full adult members than the preset allows")

// Member is account's membership of household, in a lookup transaction
// (Scope{AccountID}): the selector check of SEC-38 (a) (09.2 §3 layer 1;
// plan 1.6 decision 4). No membership, whether the household exists or
// not, is ErrNotFound (SEC-05; FR-FAM-020/AC-1, AC-4).
func (s *Service) Member(ctx context.Context, account, household pgtype.UUID) (Member, error) {
	var row householddb.MemberOfRow
	err := database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) (err error) {
		row, err = householddb.New(tx).MemberOf(ctx, householddb.MemberOfParams{HouseholdID: household, AccountID: account})
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return Member{}, ErrNotFound
	}
	return Member{HouseholdID: row.HouseholdID, MemberID: row.MemberID, AccountID: account}, err
}

// MemberIn is account's membership of tx's household, checked again inside
// a command (SEC-151; 09.2 §6 step 2); none is ErrNotFound.
func (s *Service) MemberIn(ctx context.Context, tx pgx.Tx, account pgtype.UUID) (Member, error) {
	row, err := householddb.New(tx).MemberIn(ctx, account)
	if errors.Is(err, pgx.ErrNoRows) {
		return Member{}, ErrNotFound
	}
	return Member{HouseholdID: row.HouseholdID, MemberID: row.MemberID, AccountID: account}, err
}

// FullAdultMembers are the full adult members of tx's household (SEC-10;
// 06 §3.2 full_adult_members; 09.2 §1): every caller that needs them uses
// this one definition. Its query is bounded (ES §3): more than two rows (F01
// keeps a household at two adults, 06 §3.6) is an error.
func (s *Service) FullAdultMembers(ctx context.Context, tx pgx.Tx) ([]pgtype.UUID, error) {
	members, err := householddb.New(tx).FullAdultMembers(ctx)
	if err == nil && len(members) > maxFullAdults {
		return nil, errTooManyFullAdults
	}
	return members, err
}

// Person is a listed member: what other members see of it.
type Person struct {
	MemberID    pgtype.UUID
	DisplayName string
}

// Membership is one household of an account, with the account's member id
// and the household's members.
type Membership struct {
	HouseholdID   pgtype.UUID
	HouseholdName string
	MemberID      pgtype.UUID
	Members       []Person
}

// Memberships lists account's households, at most 20, each with at most 20
// members [Usulan] (plan 1.6 "Batasan global"). The households are found in
// a lookup transaction; each is then read in its own scope, with the
// membership checked again (09.2 §3 layer 1; 10.1 §5.1 note 1.6 item 1).
func (s *Service) Memberships(ctx context.Context, account pgtype.UUID) ([]Membership, error) {
	found, err := s.membershipsOf(ctx, account)
	if err != nil {
		return nil, err
	}
	// ponytail: 1+N reads. membership_one_household limits it to one
	// iteration until F-FAM-009; RLS `id = H` forces one tx per household.
	memberships := make([]Membership, 0, len(found))
	for _, f := range found {
		var rows []householddb.HouseholdMembersRow
		err := database.Read(ctx, s.pool, database.Scope{HouseholdID: f.HouseholdID, AccountID: account}, func(tx pgx.Tx) (err error) {
			rows, err = householddb.New(tx).HouseholdMembers(ctx, householddb.HouseholdMembersParams{
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
			m.Members = append(m.Members, Person{MemberID: r.MemberID, DisplayName: r.DisplayName.String})
		}
		memberships = append(memberships, m)
	}
	return memberships, nil
}

// membershipsOf is account's own memberships, at most 20, in a lookup
// transaction.
func (s *Service) membershipsOf(ctx context.Context, account pgtype.UUID) ([]householddb.MembershipsOfAccountRow, error) {
	var rows []householddb.MembershipsOfAccountRow
	err := database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) (err error) {
		rows, err = householddb.New(tx).MembershipsOfAccount(ctx, account)
		return err
	})
	return rows, err
}
