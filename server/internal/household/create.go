package household

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/household/householddb"
)

// Member is an acting member: a membership of the caller's account.
type Member struct{ HouseholdID, MemberID pgtype.UUID }

// NewHousehold is what creating a household needs, its names already
// cleaned by CleanName and its time zone checked.
type NewHousehold struct{ Name, TimeZone, DisplayName string }

// NewHouseholdID makes a household's id before its write, which sets it as
// the transaction's household (ADR-017; 09.2 §3 note of 14 November 2026).
func (s *Service) NewHouseholdID(ctx context.Context) (pgtype.UUID, error) {
	var id pgtype.UUID
	err := database.Read(ctx, s.pool, database.Scope{}, func(tx pgx.Tx) (err error) {
		id, err = householddb.New(tx).NewID(ctx)
		return err
	})
	return id, err
}

// MemberAnywhere reports whether account is a member of any household, in
// a lookup transaction (09.2 §3 layer 1; FR-FAM-001 note of 14 November
// 2026, Q5 (d)).
func (s *Service) MemberAnywhere(ctx context.Context, account pgtype.UUID) (bool, error) {
	memberships, err := s.membershipsOf(ctx, account)
	return len(memberships) > 0, err
}

// Create writes a new household in tx, whose scope is the household and
// account (FR-FAM-001; 06 §3.1): the household (IDR, market ID, version 0),
// account's membership as household owner of two equal adults, its approval
// authority for every request type without limit, household rule version 1
// without threshold, no relationship, the household's creation event, and
// the adult age statement of version recordAge unless it is 0 (SEC-80). A
// violation of membership_one_household is ErrAlreadyMember (09.2 §5.1, note
// on decision 20).
func (s *Service) Create(ctx context.Context, tx pgx.Tx, account, household pgtype.UUID, h NewHousehold, recordAge uint32, key pgtype.UUID) (Member, error) {
	now := s.now()
	q := householddb.New(tx)
	if err := q.InsertHousehold(ctx, householddb.InsertHouseholdParams{
		ID: household, Name: h.Name, TimeZone: h.TimeZone, CreatedAt: now,
	}); err != nil {
		return Member{}, err
	}
	memberID, err := q.InsertMembership(ctx, householddb.InsertMembershipParams{
		HouseholdID: household, AccountID: account, DisplayName: optionalText(h.DisplayName), JoinedAt: now,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "membership_one_household" {
			return Member{}, ErrAlreadyMember
		}
		return Member{}, err
	}
	if err := q.InsertApprovalAuthority(ctx, householddb.InsertApprovalAuthorityParams{
		HouseholdID: household, MemberID: memberID, ValidFrom: now,
	}); err != nil {
		return Member{}, err
	}
	if err := q.InsertFirstHouseholdRule(ctx, householddb.InsertFirstHouseholdRuleParams{
		HouseholdID: household, CreatedAt: now,
	}); err != nil {
		return Member{}, err
	}
	if err := RecordEvent(ctx, tx, Event{
		HouseholdID: household, ObjectKind: "household", ObjectID: household, Seq: 0,
		ActorKind: "member", Actor: memberID, Command: "create_household", ToState: "created",
		IdempotencyKey: key, At: now,
	}); err != nil {
		return Member{}, err
	}
	if recordAge != 0 {
		if err := q.InsertAgeDeclaration(ctx, householddb.InsertAgeDeclarationParams{
			AccountID: account, TextVersion: int32(recordAge), RecordedAt: now, //nolint:gosec // G115: recordAge is currentAgeDeclarationVersion (CheckAgeDeclaration)
		}); err != nil {
			return Member{}, err
		}
	}
	return Member{HouseholdID: household, MemberID: memberID}, nil
}

// FullAdultMembers are the full adult members of tx's household (SEC-10;
// 06 §3.2 full_adult_members; 09.2 §1): every caller that needs them uses
// this one definition.
func (s *Service) FullAdultMembers(ctx context.Context, tx pgx.Tx) ([]pgtype.UUID, error) {
	return householddb.New(tx).FullAdultMembers(ctx)
}
