package api_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/abiyanzimran/casa/server/internal/database"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

// invitationLifetime is SEC-28's 7 days, counted from F01.
const invitationLifetime = 7 * 24 * time.Hour

var (
	invitationPending = []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_INVITATION_PENDING}
	householdFull     = []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_HOUSEHOLD_FULL}
)

func (h *householdHarness) invite(m member, householdID string) (*casav1.CreateInvitationResponse, error) {
	return h.households.CreateInvitation(bearer(m.token), &casav1.CreateInvitationRequest{
		RequestContext: commandRC(), HouseholdId: householdID,
	})
}

func (h *householdHarness) mustInvite(t *testing.T, m member, householdID string) *casav1.CreateInvitationResponse {
	t.Helper()
	res, err := h.invite(m, householdID)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func (h *householdHarness) issueLink(m member, householdID, invitationID string) (*casav1.IssueInvitationLinkResponse, error) {
	return h.households.IssueInvitationLink(bearer(m.token), &casav1.IssueInvitationLinkRequest{
		RequestContext: identityRC(ios), HouseholdId: householdID, InvitationId: invitationID,
	})
}

func (h *householdHarness) mustIssueLink(t *testing.T, m member, householdID, invitationID string) *casav1.IssueInvitationLinkResponse {
	t.Helper()
	res, err := h.issueLink(m, householdID, invitationID)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func (h *householdHarness) cancelInvitation(m member, householdID, invitationID string) error {
	_, err := h.households.CancelInvitation(bearer(m.token), &casav1.CancelInvitationRequest{
		RequestContext: commandRC(), HouseholdId: householdID, InvitationId: invitationID,
	})
	return err
}

// signInAgain gives m a new access token once the clock has moved past the
// old one's life.
func (h *householdHarness) signInAgain(t *testing.T, m *member, email string) {
	t.Helper()
	tokens, err := h.signIn(email, "a long enough passphrase")
	if err != nil {
		t.Fatal(err)
	}
	m.token = tokens.GetAccessToken()
}

// addMember adds account to household as a second household owner, as test
// data for states the API does not reach in one step (a member of two
// households, two adults beside an invitation); it returns the member id.
func (h *householdHarness) addMember(t *testing.T, householdID, account string) string {
	t.Helper()
	var memberID string
	if err := h.admin.QueryRow(t.Context(), `
		INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
		VALUES ($1, $2, 'Budi', 'household_owner', 'two_equal_adults', $3) RETURNING member_id::text`,
		householdID, account, h.clock.Now().Add(time.Minute)).Scan(&memberID); err != nil {
		t.Fatal(err)
	}
	return memberID
}

// acceptAsFixture moves an open invitation to accepted by invitee, as test
// data: the invitee need not be a signed-up account.
func (h *householdHarness) acceptAsFixture(t *testing.T, invitationID, invitee string) {
	t.Helper()
	if _, err := h.admin.Exec(t.Context(), `
		UPDATE invitation SET state = 'accepted', invitee_account_id = $2, invitee_display_name = 'Budi',
		                      token_hash = NULL, version = version + 1
		 WHERE id = $1 AND state = 'open'`, invitationID, invitee); err != nil {
		t.Fatal(err)
	}
}

// invitationRow is an invitation's state, version, closing time and token
// hash, as "state version closed_at token_hash".
func (h *householdHarness) invitationRow(t *testing.T, invitationID string) string {
	t.Helper()
	var state string
	var version int64
	var closedAt pgtype.Timestamptz
	var tokenHash []byte
	if err := h.admin.QueryRow(t.Context(), "SELECT state, version, closed_at, token_hash FROM invitation WHERE id = $1",
		invitationID).Scan(&state, &version, &closedAt, &tokenHash); err != nil {
		t.Fatal(err)
	}
	closed := "-"
	if closedAt.Valid {
		closed = closedAt.Time.UTC().Format(time.RFC3339)
	}
	return fmt.Sprintf("%s %d %s %x", state, version, closed, tokenHash)
}

// invitationEvents are an invitation's events, oldest first, as "seq
// actor_kind actor rule command from_state to_state idempotency_key
// server_time" with "-" for NULL.
func (h *householdHarness) invitationEvents(t *testing.T, invitationID string) []string {
	t.Helper()
	return h.objectEvents(t, "invitation", invitationID)
}

// objectEvents are the events of one object, as invitationEvents prints
// them.
func (h *householdHarness) objectEvents(t *testing.T, kind, id string) []string {
	t.Helper()
	rows, err := h.admin.Query(t.Context(), `
		SELECT concat_ws(' ', seq, actor_kind, coalesce(actor::text, '-'), coalesce(rule, '-'), command,
		                 coalesce(from_state, '-'), to_state, coalesce(idempotency_key::text, '-'),
		                 to_char(server_time AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"'))
		  FROM object_event WHERE object_kind = $1 AND object_id = $2 ORDER BY seq`, kind, id)
	if err != nil {
		t.Fatal(err)
	}
	events, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		t.Fatal(err)
	}
	return events
}

// utc is t as invitationRow and invitationEvents print it.
func utc(t time.Time) string { return t.UTC().Format(time.RFC3339) }

// wantHiddenNotFound checks the one answer of an object the caller may not see:
// NotFound, "not found", no detail (SEC-05; FR-FAM-020/AC-1).
func wantHiddenNotFound(t *testing.T, err error) {
	t.Helper()
	r := rejectionOf(t, err)
	if r.code != connect.CodeNotFound || r.message != "not found" || len(r.details) != 0 {
		t.Fatalf("error = %v with %d details, want NotFound \"not found\" without detail", err, len(r.details))
	}
}

func wantCodes(t *testing.T, err error, code connect.Code, codes []casav1.ErrorCode) {
	t.Helper()
	wantCode(t, err, code)
	if got := errorCodes(t, err); !slices.Equal(got, codes) {
		t.Fatalf("error codes = %v, want %v", got, codes)
	}
}

func TestFullAdultMembers(t *testing.T) {
	t.Run("SEC-10 a new household's full adult members are its creator", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		res := h.mustCreate(t, a)
		scope := database.Scope{HouseholdID: uuidOf(t, res.GetHouseholdId()), AccountID: uuidOf(t, a.account)}
		fullAdults := func() ([]string, error) {
			var adults []string
			err := database.Read(t.Context(), h.pool, scope, func(tx pgx.Tx) error {
				ids, err := h.service.FullAdultMembers(t.Context(), tx)
				for _, id := range ids {
					adults = append(adults, uuidText(id))
				}
				return err
			})
			return adults, err
		}
		adults, err := fullAdults()
		if err != nil || !slices.Equal(adults, []string{res.GetMemberId()}) {
			t.Fatalf("full adult members = %v, %v; want [%s]", adults, err, res.GetMemberId())
		}
		// F01 keeps a household at two adults (06 §3.6): a third is an
		// error, never a longer list (ES §3).
		h.addMember(t, res.GetHouseholdId(), uuid4())
		h.addMember(t, res.GetHouseholdId(), uuid4())
		if adults, err := fullAdults(); err == nil {
			t.Fatalf("full adult members of three owners = %v, want an error", adults)
		}
	})
}

func TestInvitations(t *testing.T) {
	// FR-FAM-003/AC-1, the server part: the invitation and its link. The
	// notice before creating is the client's (1.10); single use on the use
	// side is TestAcceptInvitation's "SEC-28 an accepted token cannot be used
	// again".
	t.Run("a household of one adult makes one invitation with a link", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		req := &casav1.CreateInvitationRequest{RequestContext: commandRC(), HouseholdId: hh.GetHouseholdId()}
		inv, err := h.households.CreateInvitation(bearer(a.token), req)
		if err != nil {
			t.Fatal(err)
		}
		expiresAt := h.clock.Now().Add(invitationLifetime)
		if !inv.GetExpiresAt().AsTime().Equal(expiresAt) {
			t.Fatalf("expires_at = %v, want %v", inv.GetExpiresAt().AsTime(), expiresAt)
		}
		var household, inviter, state string
		var createdAt, rowExpiresAt time.Time
		if err := h.admin.QueryRow(t.Context(), `
			SELECT household_id::text, inviter::text, state, created_at, expires_at FROM invitation WHERE id = $1`,
			inv.GetInvitationId()).Scan(&household, &inviter, &state, &createdAt, &rowExpiresAt); err != nil {
			t.Fatal(err)
		}
		if household != hh.GetHouseholdId() || inviter != hh.GetMemberId() || state != "open" ||
			!createdAt.Equal(h.clock.Now()) || !rowExpiresAt.Equal(expiresAt) {
			t.Fatalf("invitation = %s %s %s %v %v", household, inviter, state, createdAt, rowExpiresAt)
		}
		if got, want := h.invitationRow(t, inv.GetInvitationId()), "open 0 - "; got != want {
			t.Fatalf("invitation = %q, want %q (no token before IssueInvitationLink)", got, want)
		}
		link := h.mustIssueLink(t, a, hh.GetHouseholdId(), inv.GetInvitationId())
		if len(link.GetInvitationToken()) != 43 || !link.GetExpiresAt().AsTime().Equal(expiresAt) {
			t.Fatalf("link = %d characters, expires %v; want 43 and %v", len(link.GetInvitationToken()), link.GetExpiresAt().AsTime(), expiresAt)
		}
		sum := sha256.Sum256(decodeToken(t, link.GetInvitationToken()))
		if got, want := h.invitationRow(t, inv.GetInvitationId()), fmt.Sprintf("open 0 - %x", sum); got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		want := []string{fmt.Sprintf("0 member %s - invite - open %s %s",
			hh.GetMemberId(), strings.ToLower(req.GetRequestContext().GetIdempotencyKey()), utc(h.clock.Now()))}
		if got := h.invitationEvents(t, inv.GetInvitationId()); !slices.Equal(got, want) {
			t.Fatalf("events = %q, want %q", got, want)
		}
	})

	// FR-FAM-003/AC-2, the server part: no second invitation. The app
	// showing the existing one is 1.10.
	t.Run("a second invitation is not made while one is pending", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		first := h.mustIssueLink(t, a, hh.GetHouseholdId(), inv.GetInvitationId())
		_, err := h.invite(a, hh.GetHouseholdId())
		wantCodes(t, err, connect.CodeFailedPrecondition, invitationPending)
		if n := h.count(t, "SELECT count(*) FROM invitation"); n != 1 {
			t.Fatalf("%d invitations, want 1", n)
		}
		// A new link for the same invitation: the old token finds nothing,
		// the expiry and version stay, no event (plan 1.6 decision 11).
		h.clock.advance(time.Minute)
		second := h.mustIssueLink(t, a, hh.GetHouseholdId(), inv.GetInvitationId())
		if second.GetInvitationToken() == first.GetInvitationToken() {
			t.Fatal("the reissued token equals the first")
		}
		if !second.GetExpiresAt().AsTime().Equal(first.GetExpiresAt().AsTime()) {
			t.Fatalf("reissued expiry %v, want the first %v", second.GetExpiresAt().AsTime(), first.GetExpiresAt().AsTime())
		}
		oldHash := sha256.Sum256(decodeToken(t, first.GetInvitationToken()))
		if n := h.count(t, "SELECT count(*) FROM invitation WHERE token_hash = $1", oldHash[:]); n != 0 {
			t.Fatalf("the old token finds %d invitations, want 0", n)
		}
		newHash := sha256.Sum256(decodeToken(t, second.GetInvitationToken()))
		if got, want := h.invitationRow(t, inv.GetInvitationId()), fmt.Sprintf("open 0 - %x", newHash); got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		if got := h.invitationEvents(t, inv.GetInvitationId()); len(got) != 1 {
			t.Fatalf("events = %q, want the creation only", got)
		}
	})

	// FR-FAM-003/AC-3, the part of this task: the invitation is closed and
	// the cancellation recorded. That it can no longer be used is proved by
	// TestAcceptInvitation's "SEC-30 every invalid invitation gets the one
	// token rejection"; "no notification" in step 4.
	t.Run("a cancelled invitation is recorded and closed", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		h.clock.advance(time.Minute)
		req := &casav1.CancelInvitationRequest{
			RequestContext: commandRC(), HouseholdId: hh.GetHouseholdId(), InvitationId: inv.GetInvitationId(),
		}
		if _, err := h.households.CancelInvitation(bearer(a.token), req); err != nil {
			t.Fatal(err)
		}
		if got, want := h.invitationRow(t, inv.GetInvitationId()), "cancelled 1 "+utc(h.clock.Now())+" "; got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		events := h.invitationEvents(t, inv.GetInvitationId())
		want := fmt.Sprintf("1 member %s - cancel_invitation open cancelled %s %s",
			hh.GetMemberId(), strings.ToLower(req.GetRequestContext().GetIdempotencyKey()), utc(h.clock.Now()))
		if len(events) != 2 || events[1] != want {
			t.Fatalf("events = %q, want the creation and %q", events, want)
		}
		// A resend is answered from the idempotency record (ADR-005).
		if _, err := h.households.CancelInvitation(bearer(a.token), req); err != nil {
			t.Fatalf("resend: %v", err)
		}
		if got := h.invitationEvents(t, inv.GetInvitationId()); len(got) != 2 {
			t.Fatalf("events after a resend = %q, want 2", got)
		}
	})

	// FR-FAM-003/AC-4, the inviter's part: the lapse is recorded and a new
	// invitation can be made. "Someone tries to use it" is TestAcceptInvitation's
	// "SEC-30 every invalid invitation gets the one token rejection".
	t.Run("an expired invitation lapses and a new one can be made", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		old := h.mustInvite(t, a, hh.GetHouseholdId())
		expiresAt := old.GetExpiresAt().AsTime()
		h.clock.advance(invitationLifetime + 3*time.Hour)
		h.signInAgain(t, &a, "a@example.test")
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		if inv.GetInvitationId() == old.GetInvitationId() {
			t.Fatal("the new invitation has the old id")
		}
		// 06 §3.6 "Gugur karena waktu": closed_at and the event's time are
		// expires_at, not the time it was written (decision 12).
		if got, want := h.invitationRow(t, old.GetInvitationId()), "expired 1 "+utc(expiresAt)+" "; got != want {
			t.Fatalf("old invitation = %q, want %q", got, want)
		}
		events := h.invitationEvents(t, old.GetInvitationId())
		want := "1 system - invitation.expired expire_invitation open expired - " + utc(expiresAt)
		if len(events) != 2 || events[1] != want {
			t.Fatalf("old invitation's events = %q, want the creation and %q", events, want)
		}
		if got := h.invitationRow(t, inv.GetInvitationId()); got != "open 0 - " {
			t.Fatalf("new invitation = %q, want open", got)
		}
	})

	// FR-FAM-003/AC-5, the server part: the server refuses. The membership
	// menu without the choice (FR-FAM-012) is 1.10.
	t.Run("a household of two adults cannot invite", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		h.addMember(t, hh.GetHouseholdId(), uuid4())
		_, err := h.invite(a, hh.GetHouseholdId())
		wantCodes(t, err, connect.CodeFailedPrecondition, householdFull)
		if n := h.count(t, "SELECT count(*) FROM invitation"); n != 0 {
			t.Fatalf("%d invitations, want 0", n)
		}
		if n := h.count(t, "SELECT count(*) FROM idempotency_record WHERE command LIKE '%/CreateInvitation'"); n != 0 {
			t.Fatalf("%d idempotency records of CreateInvitation, want 0", n)
		}
	})

	// FR-FAM-003/AC-6, the server part: the link answer holds nothing about
	// the household (SEC-29). The text shared with it is 1.10.
	t.Run("the link response carries only the token and its expiry", func(t *testing.T) {
		response := (&casav1.IssueInvitationLinkResponse{}).ProtoReflect().Descriptor()
		wantFields(t, response, "invitation_token", "expires_at")
		token := response.Fields().ByName("invitation_token").Options().(*descriptorpb.FieldOptions)
		if !token.GetDebugRedact() {
			t.Fatal("invitation_token is not debug_redact")
		}
	})

	// 10.1 SEC-28 note Q7 (d).
	t.Run("SEC-28 the link response is not cached", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		ctx, info := connect.NewClientContext(t.Context())
		info.RequestHeader().Set("Authorization", "Bearer "+a.token)
		if _, err := h.households.IssueInvitationLink(ctx, &casav1.IssueInvitationLinkRequest{
			RequestContext: identityRC(ios), HouseholdId: hh.GetHouseholdId(), InvitationId: inv.GetInvitationId(),
		}); err != nil {
			t.Fatal(err)
		}
		if got := info.ResponseHeader().Values("Cache-Control"); !slices.Equal(got, []string{"no-store"}) {
			t.Fatalf("Cache-Control = %q, want no-store", got)
		}
	})

	// 10.1 SEC-28 note Q7 (a); FR-FAM-003 note of 14 November 2026 (Q19).
	t.Run("SEC-28 only the inviter issues a link or cancels, and only while open", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		other := h.signUpMember(t, "b@example.test")
		h.addMember(t, hh.GetHouseholdId(), other.account) // test data
		outside := h.signUpMember(t, "c@example.test")
		before := h.fingerprint(t)
		for _, m := range []member{other, outside} {
			_, err := h.issueLink(m, hh.GetHouseholdId(), inv.GetInvitationId())
			wantHiddenNotFound(t, err)
			wantHiddenNotFound(t, h.cancelInvitation(m, hh.GetHouseholdId(), inv.GetInvitationId()))
		}
		if h.fingerprint(t) != before {
			t.Fatal("a refused request changed a table")
		}
		h.acceptAsFixture(t, inv.GetInvitationId(), other.account)
		before = h.fingerprint(t)
		_, err := h.issueLink(a, hh.GetHouseholdId(), inv.GetInvitationId())
		wantCodes(t, err, connect.CodeFailedPrecondition, stateChanged)
		wantCodes(t, h.cancelInvitation(a, hh.GetHouseholdId(), inv.GetInvitationId()), connect.CodeFailedPrecondition, stateChanged)
		if h.fingerprint(t) != before {
			t.Fatal("a refused request changed a table")
		}
	})

	// Plan 1.6 decision 7; 10.1 SEC-28 note Q7 (c).
	t.Run("issuing a link is held until the notice is acknowledged", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		h.reopen("Nama Lain", household.PublishedNotices) // a new version for a
		_, err := h.issueLink(a, hh.GetHouseholdId(), inv.GetInvitationId())
		wantCodes(t, err, connect.CodeFailedPrecondition, noticeUnread)
		if got := h.invitationRow(t, inv.GetInvitationId()); got != "open 0 - " {
			t.Fatalf("invitation = %q, want no token", got)
		}
		h.ack(t, a)
		h.mustIssueLink(t, a, hh.GetHouseholdId(), inv.GetInvitationId())
	})

	// Plan 1.6 decisions 11, 21: invitation_live decides; never Internal.
	t.Run("two concurrent invitations: exactly one is made", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		// Barrier: both inserts wait on the invitation table lock.
		waitBlocked, release := holdTables(t, h.db, "invitation")
		errs := make([]error, 2)
		var wg sync.WaitGroup
		for i := range errs {
			wg.Go(func() { _, errs[i] = h.invite(a, hh.GetHouseholdId()) })
		}
		waitBlocked(2)
		release()
		wg.Wait()
		var made, pending int
		for _, err := range errs {
			switch {
			case err == nil:
				made++
			case connect.CodeOf(err) == connect.CodeFailedPrecondition && slices.Equal(errorCodes(t, err), invitationPending):
				pending++
			default:
				t.Fatalf("error = %v, want INVITATION_PENDING", err)
			}
		}
		if made != 1 || pending != 1 {
			t.Fatalf("%d made and %d pending, want 1 and 1", made, pending)
		}
		if n := h.count(t, "SELECT count(*) FROM invitation"); n != 1 {
			t.Fatalf("%d invitations, want 1", n)
		}
		if n := h.count(t, "SELECT count(*) FROM object_event WHERE object_kind = 'invitation'"); n != 1 {
			t.Fatalf("%d invitation events, want 1", n)
		}
	})

	t.Run("SEC-28 the token is stored only as its hash", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		link := h.mustIssueLink(t, a, hh.GetHouseholdId(), inv.GetInvitationId())
		raw := decodeToken(t, link.GetInvitationToken())
		everything := h.fingerprint(t)
		for _, secret := range []string{link.GetInvitationToken(), hex.EncodeToString(raw)} {
			if strings.Contains(everything, secret) {
				t.Fatalf("the database holds the token as %q", secret)
			}
		}
		sum := sha256.Sum256(raw)
		if n := h.count(t, "SELECT count(*) FROM invitation WHERE token_hash = $1", sum[:]); n != 1 {
			t.Fatalf("%d invitations hold the token's hash, want 1", n)
		}
	})

	// 06 §3.6 "Gugur karena waktu" with change X3; 09.2 §5.1 note X3; plan
	// 1.6 decision 12: a token-free path writes the expiry before its
	// command, which then finds the invitation no longer open.
	t.Run("a token-free command on an expired invitation writes the expiry first", func(t *testing.T) {
		commands := []struct {
			name string
			call func(h *householdHarness, a member, hh, inv string) error
		}{
			{"CancelInvitation", func(h *householdHarness, a member, hh, inv string) error {
				return h.cancelInvitation(a, hh, inv)
			}},
			{"IssueInvitationLink", func(h *householdHarness, a member, hh, inv string) error {
				_, err := h.issueLink(a, hh, inv)
				return err
			}},
		}
		for _, from := range []string{"open", "accepted"} {
			for _, first := range commands {
				t.Run(from+" then "+first.name, func(t *testing.T) {
					h := startHousehold(t)
					a := h.signUpMember(t, "a@example.test")
					hh := h.mustCreate(t, a)
					inv := h.mustInvite(t, a, hh.GetHouseholdId())
					version := 1
					if from == "accepted" {
						h.acceptAsFixture(t, inv.GetInvitationId(), uuid4())
						version = 2
					}
					expiresAt := inv.GetExpiresAt().AsTime()
					h.clock.advance(invitationLifetime + 3*time.Hour)
					h.signInAgain(t, &a, "a@example.test")
					wantExpired := func() {
						t.Helper()
						if got, want := h.invitationRow(t, inv.GetInvitationId()), fmt.Sprintf("expired %d %s ", version, utc(expiresAt)); got != want {
							t.Fatalf("invitation = %q, want %q", got, want)
						}
						events := h.invitationEvents(t, inv.GetInvitationId())
						want := fmt.Sprintf("%d system - invitation.expired expire_invitation %s expired - %s", version, from, utc(expiresAt))
						if events[len(events)-1] != want || h.count(t, "SELECT count(*) FROM object_event WHERE rule IS NOT NULL") != 1 {
							t.Fatalf("events = %q, want one expiry %q", events, want)
						}
					}
					wantCodes(t, first.call(h, a, hh.GetHouseholdId(), inv.GetInvitationId()), connect.CodeFailedPrecondition, stateChanged)
					wantExpired()
					// The second command finds no live invitation: no
					// second expiry.
					for _, again := range commands {
						wantCodes(t, again.call(h, a, hh.GetHouseholdId(), inv.GetInvitationId()), connect.CodeFailedPrecondition, stateChanged)
					}
					wantExpired()
				})
			}
		}

		// The compare-and-set of expire: two token-free commands that both
		// found the invitation due write one expiry.
		t.Run("two at once write one expiry", func(t *testing.T) {
			h := startHousehold(t)
			a := h.signUpMember(t, "a@example.test")
			hh := h.mustCreate(t, a)
			inv := h.mustInvite(t, a, hh.GetHouseholdId())
			h.clock.advance(invitationLifetime)
			h.signInAgain(t, &a, "a@example.test")
			// Barrier: both expiry writes wait on the invitation table lock.
			waitBlocked, release := holdTables(t, h.db, "invitation")
			errs := make([]error, 2)
			var wg sync.WaitGroup
			for i := range errs {
				wg.Go(func() { errs[i] = h.cancelInvitation(a, hh.GetHouseholdId(), inv.GetInvitationId()) })
			}
			waitBlocked(2)
			release()
			wg.Wait()
			for _, err := range errs {
				wantCodes(t, err, connect.CodeFailedPrecondition, stateChanged)
			}
			if n := h.count(t, "SELECT count(*) FROM object_event WHERE rule = 'invitation.expired'"); n != 1 {
				t.Fatalf("%d expiry events, want 1", n)
			}
		})
	})

	// SEC-151; 09.2 §6 step 2: the check each command runs again in its own
	// transaction. A handler test with a membership ended between the
	// selector check and the command comes with F06 (step 6).
	t.Run("MemberIn refuses an account outside the household", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		outside := h.signUpMember(t, "c@example.test")
		scope := database.Scope{HouseholdID: uuidOf(t, hh.GetHouseholdId()), AccountID: uuidOf(t, outside.account)}
		err := database.Write(t.Context(), h.pool, scope, func(tx pgx.Tx) error {
			_, err := h.service.MemberIn(t.Context(), tx, uuidOf(t, outside.account))
			return err
		})
		if !errors.Is(err, household.ErrNotFound) {
			t.Fatalf("MemberIn of an outside account = %v, want ErrNotFound", err)
		}
	})
}

// uuidText is id in its text form.
func uuidText(id pgtype.UUID) string {
	value, _ := id.Value()
	s, _ := value.(string)
	return s
}

func confirmRequest(householdID, invitationID string) *casav1.ConfirmInviteeRequest {
	return &casav1.ConfirmInviteeRequest{RequestContext: commandRC(), HouseholdId: householdID, InvitationId: invitationID}
}

func (h *householdHarness) confirm(m member, req *casav1.ConfirmInviteeRequest) (*casav1.ConfirmInviteeResponse, error) {
	return h.households.ConfirmInvitee(bearer(m.token), req)
}

func (h *householdHarness) mustConfirm(t *testing.T, m member, householdID, invitationID string) string {
	t.Helper()
	res, err := h.confirm(m, confirmRequest(householdID, invitationID))
	if err != nil {
		t.Fatal(err)
	}
	return res.GetMemberId()
}

func rejectRequest(householdID, invitationID string) *casav1.RejectInviteeRequest {
	return &casav1.RejectInviteeRequest{RequestContext: commandRC(), HouseholdId: householdID, InvitationId: invitationID}
}

func (h *householdHarness) reject(m member, req *casav1.RejectInviteeRequest) error {
	_, err := h.households.RejectInvitee(bearer(m.token), req)
	return err
}

func (h *householdHarness) openInvitation(m member, householdID string) (*casav1.GetOpenInvitationResponse, error) {
	return h.households.GetOpenInvitation(bearer(m.token), &casav1.GetOpenInvitationRequest{
		RequestContext: identityRC(ios), HouseholdId: householdID,
	})
}

func (h *householdHarness) mustOpenInvitation(t *testing.T, m member, householdID string) *casav1.GetOpenInvitationResponse {
	t.Helper()
	res, err := h.openInvitation(m, householdID)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

// acceptedBy is inv after invitee accepted it as Budi through the API.
func (h *householdHarness) acceptedBy(t *testing.T, inviterEmail, inviteeEmail string) (invited, member) {
	t.Helper()
	inv := h.invited(t, inviterEmail)
	b := h.signUpMember(t, inviteeEmail)
	h.mustAccept(t, b, inv.token)
	return inv, b
}

// maskedEvent is an event as invitationEvents prints it, with its
// idempotency key shown only as present or absent and without its time.
func maskedEvent(event string) string {
	fields := strings.Fields(event)
	if fields[7] != "-" {
		fields[7] = "set"
	}
	return strings.Join(fields[:8], " ")
}

// emailRows counts the rows of the household tables and idempotency_record
// that hold email, as text or as the hex of its bytes.
func (h *householdHarness) emailRows(t *testing.T, email string) int {
	t.Helper()
	total := 0
	for _, table := range []string{
		"household", "membership", "approval_authority", "household_rule", "relationship", "invitation",
		"object_event", "idempotency_record",
	} {
		total += h.count(t, "SELECT count(*) FROM "+table+" t WHERE strpos(lower(row(t.*)::text), lower($1)) > 0 OR strpos(row(t.*)::text, $2) > 0",
			email, hex.EncodeToString([]byte(email)))
	}
	return total
}

func TestConfirmInvitee(t *testing.T) {
	accepted := casav1.InvitationState_INVITATION_STATE_ACCEPTED

	// FR-FAM-004/AC-2, the inviter's part: the inviter sees the invitee's
	// display name and verified email, whole. That the acceptance waits and
	// opens nothing is TestAcceptInvitation's "an acceptance waits for the
	// inviter and opens nothing"; N-24 is step 4; the screen's sentence is
	// 1.10.
	t.Run("the inviter sees the invitee's name and verified email, whole", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		got := h.mustOpenInvitation(t, inv.inviter, inv.household)
		if got.GetInvitationId() != inv.id || got.GetState() != casav1.InvitationState_INVITATION_STATE_OPEN ||
			!got.GetExpiresAt().AsTime().Equal(inv.expiresAt) || got.GetInviteeDisplayName() != "" || got.GetInviteeEmail() != "" {
			t.Fatalf("open invitation = %v, want %s open until %v without an invitee", got, inv.id, inv.expiresAt)
		}
		email := "budi.santoso+rumah@example.test"
		b := h.signUpMember(t, email)
		h.mustAccept(t, b, inv.token)
		before := h.fingerprint(t)
		got = h.mustOpenInvitation(t, inv.inviter, inv.household)
		if got.GetInvitationId() != inv.id || got.GetState() != accepted || !got.GetExpiresAt().AsTime().Equal(inv.expiresAt) ||
			got.GetInviteeDisplayName() != "Budi" || got.GetInviteeEmail() != email {
			t.Fatalf("open invitation = %v, want %s accepted by Budi <%s>", got, inv.id, email)
		}
		wantFields(t, got.ProtoReflect().Descriptor(), "invitation_id", "state", "expires_at", "invitee_display_name", "invitee_email")
		if h.fingerprint(t) != before {
			t.Fatal("reading the open invitation changed a table")
		}
	})

	t.Run("SEC-31 two invitees with one display name are told apart by email", func(t *testing.T) {
		h := startHousehold(t)
		inv, first := h.acceptedBy(t, "a@example.test", "budi@example.test")
		seenFirst := h.mustOpenInvitation(t, inv.inviter, inv.household)
		wantOK(t, h.reject(inv.inviter, rejectRequest(inv.household, inv.id)))
		next := h.mustInvite(t, inv.inviter, inv.household).GetInvitationId()
		link := h.mustIssueLink(t, inv.inviter, inv.household, next)
		second := h.signUpMember(t, "budi.lain@example.test")
		h.mustAccept(t, second, link.GetInvitationToken())
		seenSecond := h.mustOpenInvitation(t, inv.inviter, inv.household)
		if seenFirst.GetInviteeDisplayName() != "Budi" || seenSecond.GetInviteeDisplayName() != "Budi" {
			t.Fatalf("display names %q and %q, want Budi twice", seenFirst.GetInviteeDisplayName(), seenSecond.GetInviteeDisplayName())
		}
		if seenFirst.GetInviteeEmail() != "budi@example.test" || seenSecond.GetInviteeEmail() != "budi.lain@example.test" {
			t.Fatalf("emails %q and %q, want each invitee's own", seenFirst.GetInviteeEmail(), seenSecond.GetInviteeEmail())
		}
		if first.account == second.account || seenSecond.GetInvitationId() != next {
			t.Fatal("the second acceptance is not a second account's")
		}
	})

	// FR-FAM-004/AC-3, all but "no notification", which is step 4.
	t.Run("'not this person' ends the invitation and nothing was ever readable", func(t *testing.T) {
		h := startHousehold(t)
		inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
		h.clock.advance(time.Minute)
		req := rejectRequest(inv.household, inv.id)
		wantOK(t, h.reject(inv.inviter, req))
		if got, want := h.invitationRow(t, inv.id), "rejected 2 "+utc(h.clock.Now())+" "; got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		events := h.invitationEvents(t, inv.id)
		want := fmt.Sprintf("2 member %s - reject_invitee accepted rejected %s %s",
			inv.inviterMember, strings.ToLower(req.GetRequestContext().GetIdempotencyKey()), utc(h.clock.Now()))
		if len(events) != 3 || events[2] != want {
			t.Fatalf("events = %q, want creation, acceptance and %q", events, want)
		}
		// Only a withdrawal clears the invitee (FR-FAM-004/AC-5).
		if n := h.count(t, "SELECT count(*) FROM invitation WHERE id = $1 AND invitee_account_id = $2 AND invitee_display_name = 'Budi'",
			inv.id, b.account); n != 1 {
			t.Fatal("the rejected invitation lost its invitee")
		}
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b.account); n != 0 {
			t.Fatalf("%d memberships of the invitee, want 0", n)
		}
		if got := h.memberships(t, b); len(got) != 0 {
			t.Fatalf("the invitee lists %v, want nothing", got)
		}
		for _, c := range selectorCases {
			wantHiddenNotFound(t, c.call(h, b, inv.household, inv.id))
		}
		wantOK(t, h.reject(inv.inviter, req)) // a resend gets the stored answer (ADR-005)
		if got := h.invitationEvents(t, inv.id); len(got) != 3 {
			t.Fatalf("events after a resend = %q, want 3", got)
		}
		if got := h.mustOpenInvitation(t, inv.inviter, inv.household); got.GetInvitationId() != "" {
			t.Fatalf("open invitation = %v, want none", got)
		}
		h.mustInvite(t, inv.inviter, inv.household)
	})

	// FR-FAM-004/AC-4, all but N-25 (step 4) and the empty chat (step 5).
	t.Run("confirming makes an equal household owner", func(t *testing.T) {
		h := startHousehold(t)
		inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
		h.clock.advance(time.Minute)
		req := confirmRequest(inv.household, inv.id)
		res, err := h.confirm(inv.inviter, req)
		if err != nil {
			t.Fatal(err)
		}
		bMember, key := res.GetMemberId(), strings.ToLower(req.GetRequestContext().GetIdempotencyKey())
		if got, want := h.invitationRow(t, inv.id), "confirmed 2 "+utc(h.clock.Now())+" "; got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		events := h.invitationEvents(t, inv.id)
		want := fmt.Sprintf("2 member %s - confirm_invitee accepted confirmed %s %s", inv.inviterMember, key, utc(h.clock.Now()))
		if len(events) != 3 || events[2] != want {
			t.Fatalf("invitation events = %q, want creation, acceptance and %q", events, want)
		}
		// The joining, with the new member as its actor: with the inviter's
		// confirmation, both actors are recorded.
		events = h.objectEvents(t, "household", inv.household)
		want = fmt.Sprintf("1 member %s - confirm_invitee - member_joined %s %s", bMember, key, utc(h.clock.Now()))
		if len(events) != 2 || events[1] != want {
			t.Fatalf("household events = %q, want the creation and %q", events, want)
		}
		if n := h.count(t, "SELECT count(*) FROM household WHERE id = $1 AND version = 1", inv.household); n != 1 {
			t.Fatal("the household's version is not 1")
		}
		if n := h.count(t, `SELECT count(*) FROM membership WHERE household_id = $1 AND member_id = $2 AND account_id = $3
			AND display_name = 'Budi' AND joined_at = $4`, inv.household, bMember, b.account, h.clock.Now()); n != 1 {
			t.Fatal("the new membership is not the invitee's, named Budi, joined now")
		}
		// Role, permission set and approval authority identical to the
		// inviter's (06 §3.1).
		same := func(query string) {
			t.Helper()
			var a, other string
			if err := h.admin.QueryRow(t.Context(), query, inv.household, inv.inviterMember).Scan(&a); err != nil {
				t.Fatal(err)
			}
			if err := h.admin.QueryRow(t.Context(), query, inv.household, bMember).Scan(&other); err != nil {
				t.Fatal(err)
			}
			if a != other || a == "" {
				t.Fatalf("inviter %q, new member %q, want the same", a, other)
			}
		}
		same("SELECT string_agg(concat_ws(' ', role, permission_set), ',') FROM membership WHERE household_id = $1 AND member_id = $2")
		same(`SELECT string_agg(concat_ws(' ', request_type, coalesce(limit_amount::text, '-'), coalesce(currency, '-'),
			coalesce(valid_until::text, '-')), ',' ORDER BY request_type) FROM approval_authority WHERE household_id = $1 AND member_id = $2`)
		if n := h.count(t, "SELECT count(*) FROM approval_authority WHERE member_id = $1 AND valid_from = $2", bMember, h.clock.Now()); n != 1 {
			t.Fatal("the new member's approval authority is not valid from now")
		}
		for _, m := range []member{inv.inviter, b} {
			got := h.memberships(t, m)
			if len(got) != 1 || got[0].GetHouseholdId() != inv.household {
				t.Fatalf("memberships = %v, want %s", got, inv.household)
			}
			var members []string
			for _, p := range got[0].GetMembers() {
				members = append(members, p.GetMemberId()+" "+p.GetDisplayName())
			}
			if !slices.Equal(members, []string{inv.inviterMember + " Ayu", bMember + " Budi"}) {
				t.Fatalf("members = %q, want Ayu and Budi", members)
			}
		}
		if got := h.mustOpenInvitation(t, inv.inviter, inv.household); got.GetInvitationId() != "" {
			t.Fatalf("open invitation = %v, want none", got)
		}
		again, err := h.confirm(inv.inviter, req) // a resend gets the stored answer (ADR-005)
		if err != nil || again.GetMemberId() != bMember {
			t.Fatalf("resend = %v, %v; want member %s", again, err, bMember)
		}
		if n := h.count(t, "SELECT count(*) FROM membership WHERE household_id = $1", inv.household); n != 2 {
			t.Fatalf("%d memberships after a resend, want 2", n)
		}
	})

	// FR-FAM-004/AC-7, all but the "no notification" it takes from AC-3,
	// which is step 4: no automatic confirmation, and the acceptance lapses
	// as AC-3's (never active, nothing readable, a new invitation can be
	// made).
	t.Run("an acceptance unconfirmed past the invitation's life lapses, never activates", func(t *testing.T) {
		f := startExpiredAcceptance(t, invitationLifetime)
		f.h.signInAgain(t, &f.h1.inviter, "a@example.test")
		// A read finds it lapsed and writes nothing (plan 1.6 decision 12).
		row := f.h.invitationRow(t, f.h1.id)
		if got := f.h.mustOpenInvitation(t, f.h1.inviter, f.h1.household); got.GetInvitationId() != "" {
			t.Fatalf("open invitation = %v, want none", got)
		}
		if got := f.h.invitationRow(t, f.h1.id); got != row {
			t.Fatalf("a read changed the invitation to %q", got)
		}
		_, err := f.h.confirm(f.h1.inviter, confirmRequest(f.h1.household, f.h1.id))
		wantCodes(t, err, connect.CodeFailedPrecondition, stateChanged)
		f.wantExpired(t)
		if n := f.h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", f.b.account); n != 0 {
			t.Fatalf("%d memberships of the invitee, want 0", n)
		}
		if got := f.h.memberships(t, f.b); len(got) != 0 {
			t.Fatalf("the invitee lists %v, want nothing", got)
		}
		f.h.mustInvite(t, f.h1.inviter, f.h1.household) // lapsed as AC-3: a new invitation can be made
	})

	t.Run("FR-FAM-004/AC-9 only the inviter can activate the membership", func(t *testing.T) {
		h := startHousehold(t)
		inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
		c := h.signUpMember(t, "c@example.test")
		cMember := h.addMember(t, inv.household, c.account) // test data: another member of the household
		d := h.signUpMember(t, "d@example.test")            // an outside account
		e := h.signUpMember(t, "e@example.test")
		h.mustCreate(t, e) // a member of another household
		before := h.fingerprint(t)
		for _, m := range []member{b, c, d, e} {
			_, err := h.confirm(m, confirmRequest(inv.household, inv.id))
			wantHiddenNotFound(t, err)
		}
		// The invitee's own command does not activate it either.
		wantCodes(t, h.acceptToken(b, inv.token), connect.CodeFailedPrecondition, alreadyMember)
		if h.fingerprint(t) != before {
			t.Fatal("a refused activation changed a table")
		}
		if _, err := h.admin.Exec(t.Context(), "DELETE FROM membership WHERE member_id = $1", cMember); err != nil {
			t.Fatal(err)
		}
		h.mustConfirm(t, inv.inviter, inv.household, inv.id)
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b.account); n != 1 {
			t.Fatalf("%d memberships of the invitee after F04, want 1", n)
		}
		// system: the only system transition out of an acceptance is its
		// expiry (06 §3.6 F05 "(waktu)"; PG-04), never activation.
		inv2, b2 := h.acceptedBy(t, "a2@example.test", "b2@example.test")
		h.clock.advance(invitationLifetime)
		h.signInAgain(t, &b2, "b2@example.test")
		wantCodes(t, h.withdraw(b2, withdrawRequest()), connect.CodeFailedPrecondition, stateChanged)
		if got := h.invitationRow(t, inv2.id); !strings.HasPrefix(got, "expired 2 ") {
			t.Fatalf("invitation = %q, want expired", got)
		}
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b2.account); n != 0 {
			t.Fatalf("%d memberships of the second invitee, want 0", n)
		}
		if n := h.count(t, "SELECT count(*) FROM object_event WHERE actor_kind = 'system' AND to_state <> 'expired'"); n != 0 {
			t.Fatalf("%d system events to another state than expired, want 0", n)
		}
	})

	// Plan 1.6 decision 21; "Fokus tinjauan" item 4.
	t.Run("a confirm racing a withdrawal: exactly one wins", func(t *testing.T) {
		h := startHousehold(t)
		inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
		// Barrier: both transitions wait on the invitation table lock.
		waitBlocked, release := holdTables(t, h.db, "invitation")
		errs := make([]error, 2)
		var wg sync.WaitGroup
		wg.Go(func() { _, errs[0] = h.confirm(inv.inviter, confirmRequest(inv.household, inv.id)) })
		wg.Go(func() { errs[1] = h.withdraw(b, withdrawRequest()) })
		waitBlocked(2)
		release()
		wg.Wait()
		winner := -1
		for i, err := range errs {
			if err == nil {
				winner = i
				continue
			}
			wantCodes(t, err, connect.CodeFailedPrecondition, stateChanged)
		}
		if winner < 0 || errs[1-winner] == nil {
			t.Fatalf("errors = %v, want exactly one STATE_CHANGED", errs)
		}
		state, memberships := "confirmed", 1
		if winner == 1 {
			state, memberships = "withdrawn", 0
		}
		if got := h.invitationRow(t, inv.id); !strings.HasPrefix(got, state+" 2 ") {
			t.Fatalf("invitation = %q, want %s", got, state)
		}
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b.account); n != memberships {
			t.Fatalf("%d memberships of the invitee, want %d", n, memberships)
		}
		if got := h.invitationEvents(t, inv.id); len(got) != 3 {
			t.Fatalf("events = %q, want 3", got)
		}
	})

	// 09.2 §5.1, note on decision 20; 10.1 §10 (security-privacy's
	// amendment): the answer and the stored event do not tell the cause.
	t.Run("F04 after the invitee joined elsewhere looks exactly like F04 after a real withdrawal", func(t *testing.T) {
		run := func(t *testing.T, joinedElsewhere bool) (rejection, string) {
			t.Helper()
			h := startHousehold(t)
			inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
			if joinedElsewhere {
				// The state the race of CreateHousehold and F02 leaves
				// (TestAcceptInvitation); through the API a live acceptance
				// blocks CreateHousehold.
				e := h.signUpMember(t, "e@example.test")
				h.addMember(t, h.mustCreate(t, e).GetHouseholdId(), b.account)
			} else {
				wantOK(t, h.withdraw(b, withdrawRequest()))
			}
			h.clock.advance(time.Minute)
			req := confirmRequest(inv.household, inv.id)
			_, err := h.confirm(inv.inviter, req)
			wantCodes(t, err, connect.CodeFailedPrecondition, stateChanged)
			_, again := h.confirm(inv.inviter, req) // not recorded: a resend reads the withdrawn state
			wantCodes(t, again, connect.CodeFailedPrecondition, stateChanged)
			if n := h.count(t, "SELECT count(*) FROM idempotency_record WHERE command LIKE '%/ConfirmInvitee'"); n != 0 {
				t.Fatalf("%d idempotency records of ConfirmInvitee, want 0", n)
			}
			if got := h.invitationRow(t, inv.id); !strings.HasPrefix(got, "withdrawn 2 ") {
				t.Fatalf("invitation = %q, want withdrawn", got)
			}
			if n := h.count(t, "SELECT count(*) FROM invitation WHERE id = $1 AND invitee_account_id IS NULL AND invitee_display_name IS NULL",
				inv.id); n != 1 {
				t.Fatal("the withdrawn invitation still holds the invitee")
			}
			for _, table := range []string{"membership", "approval_authority"} {
				if n := h.count(t, "SELECT count(*) FROM "+table+" WHERE household_id = $1", inv.household); n != 1 {
					t.Fatalf("%d rows in %s of the inviting household, want 1", n, table)
				}
			}
			events := h.invitationEvents(t, inv.id)
			if len(events) != 3 {
				t.Fatalf("events = %q, want 3", events)
			}
			return rejectionOf(t, err), maskedEvent(events[2])
		}
		joined, joinedEvent := run(t, true)
		withdrawn, withdrawnEvent := run(t, false)
		if !joined.equal(withdrawn) {
			t.Fatalf("answers differ: joined elsewhere %v, withdrawn %v", joined, withdrawn)
		}
		if joinedEvent != withdrawnEvent || joinedEvent != "2 invitee - - withdraw_acceptance accepted withdrawn set" {
			t.Fatalf("events differ: joined elsewhere %q, withdrawn %q", joinedEvent, withdrawnEvent)
		}
	})

	t.Run("the invitee cannot create a household while its acceptance waits", func(t *testing.T) {
		h := startHousehold(t)
		inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
		h.clock.advance(24 * time.Hour)
		h.signInAgain(t, &inv.inviter, "a@example.test")
		h.signInAgain(t, &b, "b@example.test")
		_, err := h.create(b, newHousehold())
		wantCodes(t, err, connect.CodeFailedPrecondition, alreadyMember)
		h.mustConfirm(t, inv.inviter, inv.household, inv.id)
		if got := h.memberships(t, b); len(got) != 1 || got[0].GetHouseholdId() != inv.household {
			t.Fatalf("memberships = %v, want %s", got, inv.household)
		}
	})

	// 10.1 §5.1 note 1.6 item 3; SEC-31; FR-FAM-004/AC-2 ("tidak tersimpan
	// sebagai data rumah tangga").
	t.Run("the invitee's email is shown to nobody else and stored nowhere", func(t *testing.T) {
		h := startHousehold(t)
		email := "budi@example.test"
		inv, b := h.acceptedBy(t, "a@example.test", email)
		c := h.signUpMember(t, "c@example.test")
		cMember := h.addMember(t, inv.household, c.account) // test data: a member who is not the inviter
		got := h.mustOpenInvitation(t, c, inv.household)
		if got.GetInvitationId() != inv.id || got.GetState() != accepted || got.GetInviteeDisplayName() != "" || got.GetInviteeEmail() != "" {
			t.Fatalf("another member reads %v, want the accepted invitation without name or email", got)
		}
		for _, m := range []member{b, h.signUpMember(t, "d@example.test")} {
			_, err := h.openInvitation(m, inv.household)
			wantHiddenNotFound(t, err)
		}
		if _, err := h.admin.Exec(t.Context(), "DELETE FROM membership WHERE member_id = $1", cMember); err != nil {
			t.Fatal(err)
		}
		if got := h.mustOpenInvitation(t, inv.inviter, inv.household); got.GetInviteeEmail() != email {
			t.Fatalf("the inviter reads %q, want %q", got.GetInviteeEmail(), email)
		}
		h.mustConfirm(t, inv.inviter, inv.household, inv.id)
		if n := h.emailRows(t, email); n != 0 {
			t.Fatalf("%d household or idempotency rows hold the invitee's email, want 0", n)
		}
	})

	t.Run("RejectInvitee on an expired acceptance writes the expiry and answers state changed", func(t *testing.T) {
		f := startExpiredAcceptance(t, invitationLifetime)
		f.h.signInAgain(t, &f.h1.inviter, "a@example.test")
		for range 2 {
			wantCodes(t, f.h.reject(f.h1.inviter, rejectRequest(f.h1.household, f.h1.id)), connect.CodeFailedPrecondition, stateChanged)
			f.wantExpired(t)
		}
	})

	// 10.1 §10 note X3 (1)(a): the household never reads when an expiry
	// was written.
	t.Run("every household-readable time of an expiry is expires_at", func(t *testing.T) {
		paths := []struct {
			name     string
			accepted bool
			call     func(t *testing.T, h *householdHarness, inv invited, b member) error
		}{
			{"CreateInvitation", false, func(_ *testing.T, h *householdHarness, inv invited, _ member) error {
				_, err := h.invite(inv.inviter, inv.household)
				return err
			}},
			{"CancelInvitation", false, func(_ *testing.T, h *householdHarness, inv invited, _ member) error {
				return h.cancelInvitation(inv.inviter, inv.household, inv.id)
			}},
			{"IssueInvitationLink", false, func(_ *testing.T, h *householdHarness, inv invited, _ member) error {
				_, err := h.issueLink(inv.inviter, inv.household, inv.id)
				return err
			}},
			{"ConfirmInvitee", true, func(_ *testing.T, h *householdHarness, inv invited, _ member) error {
				_, err := h.confirm(inv.inviter, confirmRequest(inv.household, inv.id))
				return err
			}},
			{"RejectInvitee", true, func(_ *testing.T, h *householdHarness, inv invited, _ member) error {
				return h.reject(inv.inviter, rejectRequest(inv.household, inv.id))
			}},
			{"WithdrawAcceptance", true, func(_ *testing.T, h *householdHarness, _ invited, b member) error {
				return h.withdraw(b, withdrawRequest())
			}},
			{"CreateHousehold", true, func(_ *testing.T, h *householdHarness, _ invited, b member) error {
				_, err := h.create(b, newHousehold())
				return err
			}},
			{"AcceptInvitation", true, func(t *testing.T, h *householdHarness, _ invited, b member) error {
				return h.acceptToken(b, h.invited(t, "c@example.test").token)
			}},
		}
		for _, p := range paths {
			t.Run(p.name, func(t *testing.T) {
				h := startHousehold(t)
				inv := h.invited(t, "a@example.test")
				b := h.signUpMember(t, "b@example.test")
				if p.accepted {
					h.mustAccept(t, b, inv.token)
				}
				h.clock.advance(invitationLifetime + 72*time.Hour)
				h.signInAgain(t, &inv.inviter, "a@example.test")
				h.signInAgain(t, &b, "b@example.test")
				_ = p.call(t, h, inv, b) // each path's answer is proven with the path
				var closedAt, eventAt time.Time
				if err := h.admin.QueryRow(t.Context(), `
					SELECT i.closed_at, e.server_time FROM invitation i
					  JOIN object_event e ON e.object_id = i.id AND e.to_state = 'expired'
					 WHERE i.id = $1 AND i.state = 'expired'`, inv.id).Scan(&closedAt, &eventAt); err != nil {
					t.Fatal(err) // also fails on no expiry, or more than one
				}
				if !closedAt.Equal(inv.expiresAt) || !eventAt.Equal(inv.expiresAt) {
					t.Fatalf("closed_at %v, event time %v; want both %v", closedAt, eventAt, inv.expiresAt)
				}
			})
		}

		t.Run("no household response carries a version or updated_at", func(t *testing.T) {
			methods := casav1.File_casa_v1_household_proto.Services().ByName("HouseholdService").Methods()
			seen := map[protoreflect.FullName]bool{}
			var scan func(m protoreflect.MessageDescriptor)
			scan = func(m protoreflect.MessageDescriptor) {
				if seen[m.FullName()] {
					return
				}
				seen[m.FullName()] = true
				for i := range m.Fields().Len() {
					f := m.Fields().Get(i)
					// GetOperatorNoticeResponse.version is the notice text's
					// number (SEC-34), not a row's version.
					if (f.Name() == "version" && f.FullName() != "casa.v1.GetOperatorNoticeResponse.version") || f.Name() == "updated_at" {
						t.Errorf("%s is in a household response", f.FullName())
					}
					if f.Message() != nil {
						scan(f.Message())
					}
				}
			}
			for i := range methods.Len() {
				scan(methods.Get(i).Output())
			}
		})
	})

	// FR-FAM-001/AC-4, the part of 1.6: two owners (N = 2) get the same
	// result shape from every household RPC; over every permutation of N
	// members is 1.7.
	t.Run("after joining the two owners are symmetric", func(t *testing.T) {
		h := startHousehold(t)
		inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
		h.mustConfirm(t, inv.inviter, inv.household, inv.id)
		// shape is what one owner gets, without its own ids. An invitation
		// id that names no invitation: the confirmed one has an inviter,
		// which is the invitation's property, not the owner's.
		unknown := uuid4()
		outcome := func(err error) string {
			if err != nil {
				return fmt.Sprint(connect.CodeOf(err), errorCodes(t, err))
			}
			return "ok"
		}
		calls := map[string]func(m member) string{
			"CreateHousehold": func(m member) string { _, err := h.create(m, newHousehold()); return outcome(err) },
			"ListMemberships": func(m member) string {
				got := h.memberships(t, m)
				var shape []string
				for _, ms := range got {
					var members []string
					for _, p := range ms.GetMembers() {
						members = append(members, p.GetMemberId()+" "+p.GetDisplayName())
					}
					shape = append(shape, ms.GetHouseholdId()+" "+ms.GetHouseholdName()+" "+strings.Join(members, ","))
				}
				return strings.Join(shape, ";")
			},
			"GetOperatorNotice": func(m member) string {
				n := h.notice(t, m)
				return fmt.Sprint(n.GetVariant(), n.GetVersion(), n.GetAcknowledged(), hex.EncodeToString(n.GetTextSha256()))
			},
			"AcknowledgeOperatorNotice": func(m member) string { return outcome(h.acknowledge(m, h.notice(t, m))) },
			"CreateInvitation":          func(m member) string { _, err := h.invite(m, inv.household); return outcome(err) },
			"IssueInvitationLink": func(m member) string {
				_, err := h.issueLink(m, inv.household, unknown)
				return outcome(err)
			},
			"CancelInvitation": func(m member) string { return outcome(h.cancelInvitation(m, inv.household, unknown)) },
			"PreviewInvitation": func(m member) string {
				_, err := h.preview(m, "abc")
				return outcome(err)
			},
			"AcceptInvitation":   func(m member) string { return outcome(h.acceptToken(m, "abc")) },
			"WithdrawAcceptance": func(m member) string { return outcome(h.withdraw(m, withdrawRequest())) },
			"ConfirmInvitee": func(m member) string {
				_, err := h.confirm(m, confirmRequest(inv.household, unknown))
				return outcome(err)
			},
			"RejectInvitee": func(m member) string { return outcome(h.reject(m, rejectRequest(inv.household, unknown))) },
			"GetOpenInvitation": func(m member) string {
				res, err := h.openInvitation(m, inv.household)
				return outcome(err) + " " + res.GetInvitationId()
			},
		}
		methods := casav1.File_casa_v1_household_proto.Services().ByName("HouseholdService").Methods()
		for i := range methods.Len() {
			name := string(methods.Get(i).Name())
			call, ok := calls[name]
			if !ok {
				t.Errorf("%s has no symmetry call", name)
				continue
			}
			if a, other := call(inv.inviter), call(b); a != other {
				t.Errorf("%s: the inviter gets %q, the new member %q", name, a, other)
			}
		}
		if n := h.count(t, `SELECT count(DISTINCT (role, permission_set)) FROM membership WHERE household_id = $1`, inv.household); n != 1 {
			t.Fatalf("%d role and permission sets in the household, want 1", n)
		}
	})
}

// startSignUpWith calls StartSignUp with both token fields as given.
func (h *householdHarness) startSignUpWith(setupToken, invitationToken, email string) error {
	_, err := h.client.StartSignUp(context.Background(), &casav1.StartSignUpRequest{
		RequestContext: identityRC(ios), SetupToken: setupToken, InvitationToken: invitationToken, Email: email,
	})
	return err
}

// startInvited starts a sign-up through an invitation link's token.
func (h *householdHarness) startInvited(token, email string) error {
	return h.startSignUpWith("", token, email)
}

// startInvitedFrom starts a sign-up through the invitation token from the
// loopback address 127.0.0.n (clientFrom).
func (h *householdHarness) startInvitedFrom(t *testing.T, n byte, token, email string) error {
	t.Helper()
	_, err := h.clientFrom(t, net.IPv4(127, 0, 0, n)).StartSignUp(context.Background(), &casav1.StartSignUpRequest{
		RequestContext: identityRC(ios), InvitationToken: token, Email: email,
	})
	return err
}

// invitationSignupRow is the invitation token's sign-up row as "token_hash
// email expires_at account_id used_at".
func (h *householdHarness) invitationSignupRow(t *testing.T, token string) string {
	t.Helper()
	var row string
	if err := h.admin.QueryRow(t.Context(), `
		SELECT concat_ws(' ', encode(token_hash, 'hex'), email, expires_at, coalesce(account_id::text, '-'), coalesce(used_at::text, '-'))
		  FROM signup_verification WHERE invitation_token_hash = $1`, tokenHash(t, token)).Scan(&row); err != nil {
		t.Fatal(err)
	}
	return row
}

// accountOf is the id, as text, of the account with email; "" for none.
func (h *householdHarness) accountOf(t *testing.T, email string) string {
	t.Helper()
	var account string
	err := h.admin.QueryRow(t.Context(), "SELECT id::text FROM account WHERE email = $1", email).Scan(&account)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		t.Fatal(err)
	}
	return account
}

// openLink opens an emailed verification link from iOS with the test
// passphrase.
func (h *householdHarness) openLink(link string) (*casav1.SessionTokens, error) {
	return h.completeSignUp(ios, link, "a long enough passphrase")
}

// signUpInvited signs up email through the invitation token, opens the
// emailed link and acknowledges the operator notice, as a fixture from a
// peer of its own (fixtureClient).
func (h *householdHarness) signUpInvited(t *testing.T, token, email string) member {
	t.Helper()
	tokens := h.fixtureSignUp(t, &casav1.StartSignUpRequest{InvitationToken: token, Email: email}, "a long enough passphrase")
	m := member{token: tokens.GetAccessToken(), account: h.accountOf(t, email)}
	h.ack(t, m)
	return m
}

func fromSignup() *casav1.InvitationSource {
	return &casav1.InvitationSource{Source: &casav1.InvitationSource_FromSignup{FromSignup: true}}
}

func (h *householdHarness) previewFromSignup(m member) (*casav1.PreviewInvitationResponse, error) {
	return h.households.PreviewInvitation(bearer(m.token), &casav1.PreviewInvitationRequest{
		RequestContext: identityRC(ios), Source: fromSignup(),
	})
}

func (h *householdHarness) acceptFromSignup(m member) (*casav1.AcceptInvitationResponse, error) {
	req := acceptRequest("")
	req.Source = fromSignup()
	return h.accept(m, req)
}

// signupEmail is the email of the sign-up row of the invitation token.
func (h *householdHarness) signupEmail(t *testing.T, token string) string {
	t.Helper()
	var email string
	if err := h.admin.QueryRow(t.Context(), "SELECT email::text FROM signup_verification WHERE invitation_token_hash = $1",
		tokenHash(t, token)).Scan(&email); err != nil {
		t.Fatal(err)
	}
	return email
}

// TestSignUpByInvitation is closed sign-up's invitation path (SEC-158 (a)
// and its note of 14 November 2026; plan 1.6 decision 15; FAM.md
// FR-FAM-004 note "Melanjutkan sesudah mendaftar lewat undangan").
func TestSignUpByInvitation(t *testing.T) {
	// FR-FAM-001/AC-9, the invitation-token half: no token, an invalid
	// invitation, or both tokens, get the one rejection, and repeated
	// failures hit the rate limit. The setup-token half is 1.4's TestSignUp
	// ("FR-FAM-001/AC-9 …" subtests).
	t.Run("sign-up with an invalid invitation gets the one rejection", func(t *testing.T) {
		h := startHousehold(t)
		dead := h.deadInvitations(t)
		valid := h.invited(t, "c5@example.test")
		setup := h.issue(t, "x@example.test")
		never, _ := identity.NewToken()
		want := h.tokenRejection(t)
		mails, accounts := len(h.mails()), h.count(t, "SELECT count(*) FROM account")
		rows := h.count(t, "SELECT count(*) FROM signup_verification")
		for _, c := range []struct{ name, setup, invitation string }{
			{"both empty", "", ""},
			{"both set", setup, valid.token},
			{"malformed", "", "abc"},
			{"never issued", "", never},
			{"cancelled", "", dead.cancelled},
			{"expired", "", dead.expired},
			{"replaced", "", dead.replaced},
			{"accepted", "", dead.accepted},
		} {
			err := h.startSignUpWith(c.setup, c.invitation, "x@example.test")
			wantRejection(t, err, want, "StartSignUp "+c.name)
			if got := violations(t, err); got != nil {
				t.Fatalf("%s: violations %v, want none", c.name, got)
			}
		}
		if got := len(h.mails()); got != mails {
			t.Fatalf("%d mails sent for rejected sign-ups", got-mails)
		}
		if got := h.count(t, "SELECT count(*) FROM account"); got != accounts {
			t.Fatalf("%d accounts made by rejected sign-ups", got-accounts)
		}
		if got := h.count(t, "SELECT count(*) FROM signup_verification"); got != rows {
			t.Fatalf("%d sign-up rows written by rejected sign-ups", got-rows)
		}
		// Repeated failures hit the rate limit (SEC-12), and then a valid
		// invitation waits too.
		for i := 0; ; i++ {
			err := h.startInvited(never, "x@example.test")
			if connect.CodeOf(err) == connect.CodeResourceExhausted {
				break
			}
			wantRejection(t, err, want, "StartSignUp never issued, again")
			if i == 20 {
				t.Fatal("no rate limit after 20 more failed sign-ups")
			}
		}
		wantCodes(t, h.startInvited(valid.token, "x@example.test"), connect.CodeResourceExhausted, rateLimited)
	})

	t.Run("SEC-158 an invitation token opens one account for any email after the emailed link", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		mails := len(h.mails())
		wantOK(t, h.startInvited(inv.token, "Budi@Example.Test"))
		got := h.mails()[mails:]
		if len(got) != 1 || got[0].To != "Budi@Example.Test" || got[0].Kind != identity.MessageVerification ||
			!strings.HasPrefix(got[0].Link, testOrigin+"/signup/verify#") {
			t.Fatalf("mails %+v, want one verification link to Budi@Example.Test", got)
		}
		if h.accountOf(t, "budi@example.test") != "" {
			t.Fatal("an account before the link is opened")
		}
		tokens, err := h.openLink(h.lastLink(t))
		wantOK(t, err)
		if tokens.GetAccessToken() == "" {
			t.Fatal("no session for the new account")
		}
		var email, via, source string
		var born bool
		if err := h.admin.QueryRow(t.Context(), `
			SELECT a.email::text, a.signup_via, v.source, v.account_id = a.id AND v.used_at IS NOT NULL
			  FROM account a JOIN signup_verification v ON v.invitation_token_hash = $1
			 WHERE a.email = 'budi@example.test'`, tokenHash(t, inv.token)).Scan(&email, &via, &source, &born); err != nil {
			t.Fatal(err)
		}
		if email != "Budi@Example.Test" || via != "invitation" || source != "invitation" || !born {
			t.Fatalf("account %q via %q, sign-up %q used for it %t; want Budi@Example.Test, invitation, invitation, true", email, via, source, born)
		}
	})

	t.Run("SEC-158 an invitation cancelled before the link is opened opens no account", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		want := h.tokenRejection(t)
		wantOK(t, h.startInvited(inv.token, "b@example.test"))
		link := h.lastLink(t)
		wantOK(t, h.cancelInvitation(inv.inviter, inv.household, inv.id))
		_, err := h.openLink(link)
		wantRejection(t, err, want, "CompleteSignUp after the invitation was cancelled")
		if h.accountOf(t, "b@example.test") != "" {
			t.Fatal("an account opened through a cancelled invitation")
		}
	})

	t.Run("an unverified sign-up never uses up the link", func(t *testing.T) {
		// SEC-158 note (1) as changed on 14 November 2026: a restart before
		// an account is born replaces the link and its email, so F04 shows
		// the email actually verified (SEC-31).
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		want := h.tokenRejection(t)
		wantOK(t, h.startInvited(inv.token, "typo@example.test"))
		first := h.lastLink(t)
		wantOK(t, h.startInvited(inv.token, "b@example.test"))
		second := h.lastLink(t)
		if mails := h.mails(); mails[len(mails)-2].To != "typo@example.test" || mails[len(mails)-1].To != "b@example.test" {
			t.Fatalf("the two links went to %q and %q", mails[len(mails)-2].To, mails[len(mails)-1].To)
		}
		_, err := h.openLink(first)
		wantRejection(t, err, want, "CompleteSignUp with the replaced link")
		tokens, err := h.openLink(second)
		wantOK(t, err)
		if h.accountOf(t, "typo@example.test") != "" {
			t.Fatal("an account for the mistyped email")
		}
		b := member{token: tokens.GetAccessToken(), account: h.accountOf(t, "b@example.test")}
		h.ack(t, b)
		_, err = h.acceptFromSignup(b)
		wantOK(t, err)
		if got := h.mustOpenInvitation(t, inv.inviter, inv.household).GetInviteeEmail(); got != "b@example.test" {
			t.Fatalf("F04 shows %q, want b@example.test", got)
		}
	})

	t.Run("a link replaced between reading and writing it opens no account", func(t *testing.T) {
		// security-privacy ruling on Task 8 of plan 1.6: CompleteSignUp
		// writes only while the row still holds the link it read.
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		want := h.tokenRejection(t)
		wantOK(t, h.startInvited(inv.token, "typo@example.test"))
		first := h.lastLink(t)
		waitBlocked, release := holdTables(t, h.db, "account")
		done := make(chan error, 1)
		go func() {
			_, err := h.openLink(first)
			done <- err
		}()
		waitBlocked(1) // the link is read; the account insert waits
		wantOK(t, h.startInvited(inv.token, "b@example.test"))
		second := h.lastLink(t)
		release()
		wantRejection(t, <-done, want, "CompleteSignUp with a link replaced after it was read")
		if n := h.count(t, "SELECT count(*) FROM account WHERE email IN ('typo@example.test', 'b@example.test')"); n != 0 {
			t.Fatalf("%d accounts, want 0", n)
		}
		_, err := h.openLink(second)
		wantOK(t, err)
		if h.accountOf(t, "b@example.test") == "" {
			t.Fatal("the new link opened no account")
		}
	})

	t.Run("SEC-12 invitation sign-up starts are limited per invitation, registered and fresh alike", func(t *testing.T) {
		// 10.1 §8.1 row "Mulai pendaftaran per undangan": every start that
		// passes the invitation check counts on the invitation's key, 5,
		// whatever the address, and a registered email counts and is
		// answered as a fresh one (no oracle).
		h := startHousehold(t)
		registered := []string{"r1@example.test", "r2@example.test", "r3@example.test", "r4@example.test", "r5@example.test", "r6@example.test"}
		mixed := h.invited(t, registered[0])
		allRegistered := h.invited(t, registered[1])
		allFresh := h.invited(t, registered[2])
		other := h.invited(t, registered[3])
		h.signUpMember(t, registered[4])
		h.signUpMember(t, registered[5])
		address := byte(10)
		for name, c := range map[string]struct {
			token  string
			emails []string
		}{
			"alternating": {mixed.token, []string{registered[4], "f1@example.test", registered[5], "f2@example.test", registered[0], "f3@example.test"}},
			"registered":  {allRegistered.token, registered},
			"fresh":       {allFresh.token, []string{"g1@example.test", "g2@example.test", "g3@example.test", "g4@example.test", "g5@example.test", "g6@example.test"}},
		} {
			for i, email := range c.emails[:5] {
				address++
				mails := len(h.mails())
				if err := h.startInvitedFrom(t, address, c.token, email); err != nil {
					t.Fatalf("%s: start %d: %v, want OK", name, i+1, err)
				}
				if got := h.mails()[mails:]; len(got) != 1 || got[0].To != email {
					t.Fatalf("%s: start %d: mails %+v, want one to %s", name, i+1, got, email)
				}
			}
			row := h.invitationSignupRow(t, c.token)
			mails := len(h.mails())
			address++
			wantCodes(t, h.startInvitedFrom(t, address, c.token, c.emails[5]), connect.CodeResourceExhausted, rateLimited)
			if len(h.mails()) != mails || h.invitationSignupRow(t, c.token) != row {
				t.Fatalf("%s: the refused sixth start mailed or wrote the sign-up row", name)
			}
		}
		address++
		wantOK(t, h.startInvitedFrom(t, address, other.token, "o@example.test"))
	})

	t.Run("SEC-12 successful sign-up starts count on the address", func(t *testing.T) {
		// 10.1 §8.1 row "Pemakaian token": on StartSignUp every request the
		// address limiter allows counts, success too.
		h := startHousehold(t)
		var tokens []string
		for i := range 5 {
			tokens = append(tokens, h.invited(t, fmt.Sprintf("a%d@example.test", i)).token)
		}
		const address byte = 20
		for i := range 20 {
			wantOK(t, h.startInvitedFrom(t, address, tokens[i/5], fmt.Sprintf("b%d@example.test", i)))
		}
		wantCodes(t, h.startInvitedFrom(t, address, tokens[4], "c@example.test"), connect.CodeResourceExhausted, rateLimited)
	})

	t.Run("SEC-30 an email that already has an account, in any case, opens no second account", func(t *testing.T) {
		// StartSignUp answers it as a fresh email (the SEC-12 subtest
		// above); CompleteSignUp refuses it with the one rejection, and
		// the refusal writes nothing (Task 8 review, lead ruling (a)).
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		h.signUpMember(t, "taken@example.test")
		want := h.tokenRejection(t)
		accounts := h.count(t, "SELECT count(*) FROM account")
		wantOK(t, h.startInvited(inv.token, "Taken@Example.Test"))
		_, err := h.openLink(h.lastLink(t))
		wantRejection(t, err, want, "CompleteSignUp for an email that already has an account")
		if got := h.count(t, "SELECT count(*) FROM account"); got != accounts {
			t.Fatalf("%d accounts made, want 0", got-accounts)
		}
		if n := h.count(t, "SELECT count(*) FROM signup_verification WHERE invitation_token_hash = $1 AND account_id IS NULL AND used_at IS NULL",
			tokenHash(t, inv.token)); n != 1 {
			t.Fatal("the refused completion wrote the sign-up row")
		}
		// The link was not used up: another email still opens the account.
		h.signUpInvited(t, inv.token, "b@example.test")
	})

	t.Run("only an account actually born uses up the link", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		want := h.tokenRejection(t)
		h.signUpInvited(t, inv.token, "b@example.test")
		wantRejection(t, h.startInvited(inv.token, "c@example.test"), want, "StartSignUp with the link an account was born from")
		if got := h.signupEmail(t, inv.token); got != "b@example.test" {
			t.Fatalf("the born sign-up's email is now %q, want b@example.test", got)
		}
		fresh := h.mustIssueLink(t, inv.inviter, inv.household, inv.id)
		h.signUpInvited(t, fresh.GetInvitationToken(), "c@example.test")
	})

	t.Run("a setup token restarted with another email is rejected and changes nothing", func(t *testing.T) {
		// security-privacy ruling on Task 8: a setup-token row never
		// changes its email (SEC-158 (b)).
		h := startHousehold(t)
		want := h.tokenRejection(t)
		setup := h.issue(t, "owner@example.test")
		wantOK(t, h.startSignUp(setup, "owner@example.test"))
		link := h.lastLink(t)
		wantRejection(t, h.startSignUp(setup, "other@example.test"), want, "StartSignUp with the setup token and another email")
		var email string
		var linkHash []byte
		if err := h.admin.QueryRow(t.Context(), "SELECT email::text, token_hash FROM signup_verification WHERE setup_token_hash = $1",
			tokenHash(t, setup)).Scan(&email, &linkHash); err != nil {
			t.Fatal(err)
		}
		if email != "owner@example.test" || !slices.Equal(linkHash, tokenHash(t, link)) {
			t.Fatalf("the sign-up row holds %q and another link; want it unchanged", email)
		}
		_, err := h.openLink(link)
		wantOK(t, err)
	})

	// FR-FAM-001/AC-10, the invitation clause: an account made through an
	// invitation cannot create a household and only joins the inviter's.
	// The setup-token clauses are 1.4's TestSignUp ("FR-FAM-001/AC-10 …")
	// and TestCreateHousehold; the runbook clause (no command sets or shows
	// a password) is cmd/casa's setup-token tests and the runbook review.
	t.Run("an account made through an invitation joins and cannot create", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpInvited(t, inv.token, "b@example.test")
		notAllowed := []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_HOUSEHOLD_CREATION_NOT_ALLOWED}
		_, err := h.create(b, newHousehold())
		wantCodes(t, err, connect.CodePermissionDenied, notAllowed)
		_, err = h.acceptFromSignup(b)
		wantOK(t, err)
		h.mustConfirm(t, inv.inviter, inv.household, inv.id)
		if got := h.memberships(t, b); len(got) != 1 || got[0].GetHouseholdId() != inv.household {
			t.Fatalf("memberships %v, want only the inviter's household", got)
		}
		_, err = h.create(b, newHousehold())
		wantCodes(t, err, connect.CodePermissionDenied, notAllowed)
		if n := h.count(t, "SELECT count(*) FROM household"); n != 1 {
			t.Fatalf("%d households, want 1", n)
		}
	})

	t.Run("the invitee continues without pasting the link again", func(t *testing.T) {
		// FR-FAM-004/AC-14, the server part: from_signup (plan 1.6 decision
		// 15). The web client's part is proved in the web client's tests.
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		want := h.tokenRejection(t)
		c := h.signUpMember(t, "c@example.test") // made with a setup token
		_, err := h.previewFromSignup(c)
		wantRejection(t, err, want, "PreviewInvitation from_signup of a setup-token account")
		_, err = h.acceptFromSignup(c)
		wantRejection(t, err, want, "AcceptInvitation from_signup of a setup-token account")
		b := h.signUpInvited(t, inv.token, "b@example.test")
		p, err := h.previewFromSignup(b)
		wantOK(t, err)
		if p.GetHouseholdName() != "Rumah Kita" || p.GetInviterDisplayName() != "Ayu" {
			t.Fatalf("preview %v, want Rumah Kita by Ayu", p)
		}
		res, err := h.acceptFromSignup(b)
		wantOK(t, err)
		if res.GetHouseholdId() != inv.household || res.GetInvitationId() != inv.id {
			t.Fatalf("accepted %v, want the invitation %s of %s", res, inv.id, inv.household)
		}
	})

	t.Run("from_signup ends when the link is reissued, and the pasted new link works", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		want := h.tokenRejection(t)
		b := h.signUpInvited(t, inv.token, "b@example.test")
		fresh := h.mustIssueLink(t, inv.inviter, inv.household, inv.id).GetInvitationToken()
		_, err := h.previewFromSignup(b)
		wantRejection(t, err, want, "PreviewInvitation from_signup after a new link")
		_, err = h.acceptFromSignup(b)
		wantRejection(t, err, want, "AcceptInvitation from_signup after a new link")
		_, err = h.preview(b, fresh)
		wantOK(t, err)
		wantOK(t, h.acceptToken(b, fresh))
	})

	t.Run("from_signup ends when the invitation is cancelled or expired", func(t *testing.T) {
		h := startHousehold(t)
		want := h.tokenRejection(t)
		cancelled := h.invited(t, "a1@example.test")
		b1 := h.signUpInvited(t, cancelled.token, "b1@example.test")
		wantOK(t, h.cancelInvitation(cancelled.inviter, cancelled.household, cancelled.id))
		due := h.invited(t, "a2@example.test")
		b2 := h.signUpInvited(t, due.token, "b2@example.test")
		if _, err := h.admin.Exec(t.Context(), "UPDATE invitation SET expires_at = $2 WHERE id = $1", due.id, h.clock.Now()); err != nil {
			t.Fatal(err)
		}
		for name, b := range map[string]member{"cancelled": b1, "expired": b2} {
			_, err := h.previewFromSignup(b)
			wantRejection(t, err, want, "PreviewInvitation from_signup, "+name)
			_, err = h.acceptFromSignup(b)
			wantRejection(t, err, want, "AcceptInvitation from_signup, "+name)
		}
	})
}
