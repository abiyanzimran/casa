package api

import (
	"context"
	"errors"
	"time"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/command"
	"github.com/abiyanzimran/casa/server/internal/database"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

// The visible-character limits of FAM.md FR-FAM-001, note of 14 November
// 2026 (Q5 (a)) [Usulan].
const (
	maxHouseholdName = 60
	maxDisplayName   = 40
)

// RegisterHouseholdService serves casa.v1.HouseholdService with h. A command
// that touches identity too is composed here, since the two modules do not
// import each other (09.2 §1 item 9).
func RegisterHouseholdService(server *connect.Server, pool *pgxpool.Pool, h *household.Service, id *identity.Service) {
	casav1connect.RegisterHouseholdServiceHandler(server, householdHandler{pool: pool, h: h, id: id})
}

type householdHandler struct {
	pool *pgxpool.Pool
	h    *household.Service
	id   *identity.Service
}

// CreateHousehold is an account-level command (09.2 §8; plan 1.6 decisions
// 5, 8). Its checks, in order: the names and the time zone; a resend gets
// the stored response, before the membership check that the first send
// turned into a refusal (ADR-005); the account is no member of any
// household (FR-FAM-001/AC-5)
// and was made with a setup token (AC-10; SEC-158), both
// HOUSEHOLD_CREATION_NOT_ALLOWED; the adult age statement (SEC-80); then
// the write, where a concurrent creation is caught as ALREADY_HOUSEHOLD_MEMBER.
func (s householdHandler) CreateHousehold(ctx context.Context, req *casav1.CreateHouseholdRequest) (*casav1.CreateHouseholdResponse, error) {
	h, err := newHousehold(req)
	if err != nil {
		return nil, err
	}
	session, _ := SessionFrom(ctx) // not public: the boundary authenticated it
	account, key := session.AccountID, req.GetRequestContext().GetIdempotencyKey()
	procedure := casav1connect.HouseholdServiceCreateHouseholdProcedure
	// A scope without a household keys the record on the account, as
	// RunForAccount does.
	res, found, err := command.Stored[*casav1.CreateHouseholdResponse](ctx, s.pool, database.Scope{AccountID: account}, procedure, key, req)
	if err != nil || found {
		return res, householdErrorOrNil(err)
	}
	member, err := s.h.MemberAnywhere(ctx, account)
	if err != nil {
		return nil, householdError(err)
	}
	if member {
		return nil, householdError(household.ErrCreationNotAllowed)
	}
	via, err := s.id.SignupVia(ctx, account)
	if err != nil {
		return nil, householdError(err)
	}
	if via != "setup_token" {
		return nil, householdError(household.ErrCreationNotAllowed)
	}
	recordAge, err := s.h.CheckAgeDeclaration(ctx, account, req.GetAgeDeclarationVersion())
	if err != nil {
		return nil, householdError(err)
	}
	id, err := s.h.NewHouseholdID(ctx)
	if err != nil {
		return nil, householdError(err)
	}
	res, err = command.RunForAccount(ctx, s.pool, database.Scope{HouseholdID: id, AccountID: account}, procedure, key, req,
		func(ctx context.Context, tx pgx.Tx) (*casav1.CreateHouseholdResponse, error) {
			m, err := s.h.Create(ctx, tx, account, id, h, recordAge, idempotencyKey(key))
			if err != nil {
				return nil, err
			}
			return &casav1.CreateHouseholdResponse{HouseholdId: uuidString(m.HouseholdID), MemberId: uuidString(m.MemberID)}, nil
		})
	if err != nil {
		return nil, householdError(err)
	}
	return res, nil
}

// newHousehold checks the names with household.CleanName and the time zone
// against the IANA database (06 §2.4; cmd/casa embeds time/tzdata), each
// refusal reported on its field (plan 1.6 decision 8).
func newHousehold(req *casav1.CreateHouseholdRequest) (household.NewHousehold, error) {
	fields := req.ProtoReflect().Descriptor().Fields()
	var violations []*validate.Violation
	name, rule := household.CleanName(req.GetName(), maxHouseholdName)
	if rule != "" {
		violations = append(violations, fieldViolation(fields.ByName("name"), rule, "name is not accepted"))
	}
	// "Local" is the server's own zone, not an IANA name.
	if _, err := time.LoadLocation(req.GetTimeZone()); err != nil || req.GetTimeZone() == "Local" {
		violations = append(violations, fieldViolation(fields.ByName("time_zone"), "time_zone.iana", "time zone is not an IANA name"))
	}
	display, rule := household.CleanName(req.GetDisplayName(), maxDisplayName)
	if rule != "" {
		violations = append(violations, fieldViolation(fields.ByName("display_name"), rule, "display name is not accepted"))
	}
	if violations != nil {
		return household.NewHousehold{}, invalidFields(violations...)
	}
	return household.NewHousehold{Name: name, TimeZone: req.GetTimeZone(), DisplayName: display}, nil
}

func (s householdHandler) ListMemberships(ctx context.Context, _ *casav1.ListMembershipsRequest) (*casav1.ListMembershipsResponse, error) {
	session, _ := SessionFrom(ctx) // not public: the boundary authenticated it
	memberships, err := s.h.Memberships(ctx, session.AccountID)
	if err != nil {
		return nil, householdError(err)
	}
	res := &casav1.ListMembershipsResponse{}
	for _, m := range memberships {
		membership := &casav1.Membership{
			HouseholdId: uuidString(m.HouseholdID), HouseholdName: m.HouseholdName, MemberId: uuidString(m.MemberID),
		}
		for _, p := range m.Members {
			membership.Members = append(membership.Members, &casav1.Member{MemberId: uuidString(p.MemberID), DisplayName: p.DisplayName})
		}
		res.Memberships = append(res.Memberships, membership)
	}
	return res, nil
}

// uuidString is id in its text form.
func uuidString(id pgtype.UUID) string {
	value, _ := id.Value() // a valid uuid always has a text form
	s, _ := value.(string)
	return s
}

// idempotencyKey is the command's key, already checked by the boundary and
// by command (09.2 §8), as the event's idempotency_key.
func idempotencyKey(key string) pgtype.UUID {
	var id pgtype.UUID
	_ = id.Scan(key) // checked: an unreadable key fails command.Stored first
	return id
}

// householdErrorOrNil is householdError, and nil for nil.
func householdErrorOrNil(err error) error {
	if err == nil {
		return nil
	}
	return householdError(err)
}

// householdError maps a household error to the contract (ADR-013 note 1.6;
// plan 1.6 decision 17). Each code has one fixed message; any other error is
// internal and its text never reaches the client (09.2 §3 layer 3).
func householdError(err error) error {
	switch {
	case errors.Is(err, household.ErrNotFound):
		return connect.NewError(connect.CodeNotFound, "not found") // no detail (SEC-05)
	case errors.Is(err, household.ErrAgeDeclarationRequired):
		return errorWithDetail(connect.CodeFailedPrecondition, "age declaration required",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_AGE_DECLARATION_REQUIRED})
	case errors.Is(err, household.ErrCreationNotAllowed):
		return errorWithDetail(connect.CodePermissionDenied, "household creation not allowed",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_HOUSEHOLD_CREATION_NOT_ALLOWED})
	case errors.Is(err, household.ErrAlreadyMember):
		return errorWithDetail(connect.CodeFailedPrecondition, "already a household member",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_ALREADY_HOUSEHOLD_MEMBER})
	}
	return commandError(err)
}
