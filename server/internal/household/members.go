package household

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/household/householddb"
)

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
