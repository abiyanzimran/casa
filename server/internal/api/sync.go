package api

import (
	"context"
	"errors"

	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/read"
	"github.com/abiyanzimran/casa/server/internal/sync"
)

// defaultSyncLimit is the page of a request with limit 0 [Usulan] (plan 1.7
// decision 14); the contract caps it at 500.
const defaultSyncLimit = 100

// RegisterSyncService serves casa.v1.SyncService: each member's feed and
// initial snapshot from s, with the selector checked by h and the notice
// state from noticeHeld (NoticeHeld) (09.3 §3; ADR-017; plan 1.7 decision
// 12).
func RegisterSyncService(server *connect.Server, h *household.Service, s *sync.Service, noticeHeld func(context.Context, pgtype.UUID) (bool, error)) {
	casav1connect.RegisterSyncServiceHandler(server, syncHandler{h: h, s: s, noticeHeld: noticeHeld})
}

type syncHandler struct {
	h          *household.Service
	s          *sync.Service
	noticeHeld func(context.Context, pgtype.UUID) (bool, error)
}

// Sync is a read (09.3 §1 "Membaca tetap boleh"): no idempotency key, never
// held by NOTICE_UNREAD. The selector check (SEC-38 (a)), without the lazy
// expiry write of member, since a read never writes an expiry (as
// GetOpenInvitation; plan 1.6 decision 12); then the notice state, which may
// issue a new version on first sight (09.3 §1 "Terbit saat pertama
// terlihat"; FR-FAM-019/AC-3, AC-10); then the page through read (SEC-06).
// A household the caller is not a member of and a cursor not made for this
// member get the one NotFound of a household that does not exist (SEC-05;
// plan 1.7 decision 9).
func (s syncHandler) Sync(ctx context.Context, req *casav1.SyncRequest) (*casav1.SyncResponse, error) {
	session, _ := SessionFrom(ctx) // not public: the boundary authenticated it
	m, err := s.h.Member(ctx, session.AccountID, checkedUUID(req.GetHouseholdId()))
	if err != nil {
		return nil, householdError(err)
	}
	unacknowledged := false
	limit := int(req.GetLimit())
	if limit == 0 {
		limit = defaultSyncLimit
	}
	page, err := s.s.Pull(ctx, read.Reader{HouseholdID: m.HouseholdID, AccountID: m.AccountID, MemberID: m.MemberID}, req.GetCursor(), limit)
	if err != nil {
		return nil, householdError(err)
	}
	res := &casav1.SyncResponse{
		NextCursor: page.Next, HasMore: page.HasMore, ResyncRequired: page.ResyncRequired,
		OperatorNoticeUnacknowledged: unacknowledged,
	}
	for _, c := range page.Changes {
		change, err := syncChange(c)
		if err != nil {
			return nil, householdError(err) // internal: its text never reaches the client
		}
		res.Changes = append(res.Changes, change)
	}
	return res, nil
}

// objectKinds are the contract's names of the synced kinds (plan 1.7
// decision 7).
var objectKinds = map[read.ObjectKind]casav1.ObjectKind{
	read.KindHousehold:  casav1.ObjectKind_OBJECT_KIND_HOUSEHOLD,
	read.KindMember:     casav1.ObjectKind_OBJECT_KIND_MEMBER,
	read.KindInvitation: casav1.ObjectKind_OBJECT_KIND_INVITATION,
}

// errUnsyncable: a change the contract cannot carry. It fails the Sync as
// internal rather than send an UNSPECIFIED kind or state, or an UPSERT
// without its object, that a client would misread (09.3 §3).
var errUnsyncable = errors.New("sync change the contract cannot carry")

// syncChange is c in the contract: a purge carries no object (09.3 §3).
func syncChange(c sync.Change) (*casav1.Change, error) {
	kind, ok := objectKinds[c.Ref.Kind]
	if !ok {
		return nil, errUnsyncable
	}
	out := &casav1.Change{ObjectKind: kind, ObjectId: uuidString(c.Ref.ID)}
	switch c.Kind {
	case sync.Purge:
		out.Kind = casav1.ChangeKind_CHANGE_KIND_PURGE
		return out, nil
	case sync.Upsert:
		out.Kind = casav1.ChangeKind_CHANGE_KIND_UPSERT
	default:
		return nil, errUnsyncable
	}
	switch o := c.Object; {
	case o.Household != nil:
		out.Object = &casav1.Change_Household{Household: &casav1.Household{
			HouseholdId: uuidString(o.Household.ID), Name: o.Household.Name, TimeZone: o.Household.TimeZone, Currency: o.Household.Currency,
		}}
	case o.Member != nil:
		out.Object = &casav1.Change_Member{Member: &casav1.Member{MemberId: uuidString(o.Member.MemberID), DisplayName: o.Member.DisplayName}}
	case o.Invitation != nil:
		state, ok := liveInvitationStates[o.Invitation.State]
		if !ok {
			return nil, errUnsyncable
		}
		out.Object = &casav1.Change_Invitation{Invitation: &casav1.Invitation{
			InvitationId: uuidString(o.Invitation.ID), State: state,
			ExpiresAt: timestamppb.New(o.Invitation.ExpiresAt), InviterMemberId: uuidString(o.Invitation.Inviter),
		}}
	default:
		return nil, errUnsyncable
	}
	return out, nil
}
