package api_test

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/abiyanzimran/casa/server/internal/database"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/household"
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
// data until F04 (Task 7) does it through the API; it returns the member id.
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
// data until F02 (Task 6) does it through the API.
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
	rows, err := h.admin.Query(t.Context(), `
		SELECT concat_ws(' ', seq, actor_kind, coalesce(actor::text, '-'), coalesce(rule, '-'), command,
		                 coalesce(from_state, '-'), to_state, coalesce(idempotency_key::text, '-'),
		                 to_char(server_time AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"'))
		  FROM object_event WHERE object_kind = 'invitation' AND object_id = $1 ORDER BY seq`, invitationID)
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
	t.Run("FR-FAM-003/AC-1 a household of one adult makes one single-use invitation", func(t *testing.T) {
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

	t.Run("FR-FAM-003/AC-2 a second invitation is not made while one is pending", func(t *testing.T) {
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

	t.Run("FR-FAM-003/AC-4 an expired invitation lapses and a new one can be made", func(t *testing.T) {
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

	t.Run("FR-FAM-003/AC-5 a household of two adults cannot invite", func(t *testing.T) {
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

	// SEC-151; 09.2 §6 step 2: a command checks the membership again in its
	// own transaction, which a membership ended since the selector check
	// fails.
	t.Run("the membership is checked again inside the command", func(t *testing.T) {
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
