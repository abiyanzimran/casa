package api

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
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

// NoticeHeld is the NOTICE_UNREAD check of NewServer step 6a: in its own
// database.Write, the account meets the latest notice version (issued on
// first sight with its security trail line) and is held until it
// acknowledged it (SEC-34; 09.3 §1 "Mekanisme").
func NoticeHeld(pool *pgxpool.Pool, h *household.Service) func(context.Context, pgtype.UUID) (bool, error) {
	return func(ctx context.Context, account pgtype.UUID) (held bool, err error) {
		err = database.Write(ctx, pool, database.Scope{AccountID: account}, func(tx pgx.Tx) error {
			n, err := currentNotice(ctx, tx, h, account)
			held = !n.Acknowledged
			return err
		})
		return held, err
	}
}

// currentNotice is household.CurrentNotice with the security trail line of a
// version issued to the account, in the same tx: one line per version per
// account (SEC-34; 10.1 §11 note of 14 November 2026 (c)). identity and
// household do not import each other, so it is composed here (09.2 §1 item
// 9).
func currentNotice(ctx context.Context, tx pgx.Tx, h *household.Service, account pgtype.UUID) (household.Notice, error) {
	n, issued, err := h.CurrentNotice(ctx, tx, account)
	if err == nil && issued {
		err = identity.RecordSecurityEvent(ctx, tx, account, "notice_version", noticeDetail(n), h.Now())
	}
	return n, err
}

// noticeDetail is a notice trail line's detail: variant, number and the
// first 8 hex characters of the hash, so a version issued again for a new
// operator name is not two identical lines (10.1 §11 note of 14 November
// 2026 (c)).
func noticeDetail(n household.Notice) string {
	return fmt.Sprintf("%s:%d:%s", n.Variant, n.Version, hex.EncodeToString(n.TextSHA256)[:8])
}

// GetOperatorNotice is an online read that may issue the version on first
// sight, so it runs in a database.Write (09.3 §1 "Mekanisme").
func (s householdHandler) GetOperatorNotice(ctx context.Context, _ *casav1.GetOperatorNoticeRequest) (*casav1.GetOperatorNoticeResponse, error) {
	session, _ := SessionFrom(ctx) // not public: the boundary authenticated it
	var n household.Notice
	err := database.Write(ctx, s.pool, database.Scope{AccountID: session.AccountID}, func(tx pgx.Tx) (err error) {
		n, err = currentNotice(ctx, tx, s.h, session.AccountID)
		return err
	})
	if err != nil {
		return nil, householdError(err)
	}
	res := &casav1.GetOperatorNoticeResponse{
		Variant: n.Variant, Version: uint32(n.Version), Title: n.Title, //nolint:gosec // G115: a version is positive (latestNotice)
		Acknowledged: n.Acknowledged, TextSha256: n.TextSHA256,
	}
	for _, p := range n.Paragraphs {
		res.Paragraphs = append(res.Paragraphs, &casav1.NoticeParagraph{Text: p.Text, Feature: p.Feature})
	}
	return res, nil
}

// AcknowledgeOperatorNotice is an account-level command, notice_exempt
// (09.3 §1): in one transaction the account meets the latest version, the
// acknowledgement is checked against it and written, and the first
// acknowledgement writes one security trail line (SEC-34; FR-FAM-019/AC-2).
// Any other version is STATE_CHANGED and nothing is written.
func (s householdHandler) AcknowledgeOperatorNotice(ctx context.Context, req *casav1.AcknowledgeOperatorNoticeRequest) (*casav1.AcknowledgeOperatorNoticeResponse, error) {
	session, _ := SessionFrom(ctx) // not public: the boundary authenticated it
	account := session.AccountID
	res, err := command.Run(ctx, s.pool, database.Scope{AccountID: account},
		casav1connect.HouseholdServiceAcknowledgeOperatorNoticeProcedure, req.GetRequestContext().GetIdempotencyKey(), req,
		func(ctx context.Context, tx pgx.Tx) (*casav1.AcknowledgeOperatorNoticeResponse, error) {
			n, err := currentNotice(ctx, tx, s.h, account)
			if err != nil {
				return nil, err
			}
			version := int32(req.GetVersion()) //nolint:gosec // G115: the contract bounds it to 2147483647
			if err := s.h.AcknowledgeNotice(ctx, tx, account, req.GetVariant(), version, req.GetTextSha256()); err != nil {
				return nil, err
			}
			if !n.Acknowledged {
				err = identity.RecordSecurityEvent(ctx, tx, account, "notice_acknowledged", noticeDetail(n), s.h.Now())
			}
			return &casav1.AcknowledgeOperatorNoticeResponse{}, err
		})
	if err != nil {
		return nil, householdError(err)
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
	case errors.Is(err, household.ErrStateChanged):
		return errorWithDetail(connect.CodeFailedPrecondition, "state changed", // no state disclosed
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_STATE_CHANGED})
	case errors.Is(err, household.ErrAlreadyMember):
		return errorWithDetail(connect.CodeFailedPrecondition, "already a household member",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_ALREADY_HOUSEHOLD_MEMBER})
	}
	return commandError(err)
}
