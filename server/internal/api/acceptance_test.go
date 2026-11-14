package api_test

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5"

	"github.com/abiyanzimran/casa/server/internal/command"
	"github.com/abiyanzimran/casa/server/internal/database"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

var (
	alreadyMember          = []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_ALREADY_HOUSEHOLD_MEMBER}
	ageDeclarationRequired = []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_AGE_DECLARATION_REQUIRED}
	rateLimited            = []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_RATE_LIMITED}
)

// invited is a household made through the API by its inviter, with an open
// invitation and the token of its link.
type invited struct {
	inviter                      member
	household, inviterMember, id string
	token                        string
	expiresAt                    time.Time
}

func (h *householdHarness) invited(t *testing.T, inviterEmail string) invited {
	t.Helper()
	a := h.signUpMember(t, inviterEmail)
	hh := h.mustCreate(t, a)
	inv := h.mustInvite(t, a, hh.GetHouseholdId())
	link := h.mustIssueLink(t, a, hh.GetHouseholdId(), inv.GetInvitationId())
	return invited{
		inviter: a, household: hh.GetHouseholdId(), inviterMember: hh.GetMemberId(), id: inv.GetInvitationId(),
		token: link.GetInvitationToken(), expiresAt: inv.GetExpiresAt().AsTime(),
	}
}

func tokenSource(token string) *casav1.InvitationSource {
	return &casav1.InvitationSource{Source: &casav1.InvitationSource_InvitationToken{InvitationToken: token}}
}

func (h *householdHarness) preview(m member, token string) (*casav1.PreviewInvitationResponse, error) {
	return h.households.PreviewInvitation(bearer(m.token), &casav1.PreviewInvitationRequest{
		RequestContext: identityRC(ios), Source: tokenSource(token),
	})
}

// acceptRequest accepts with token as Budi, with the current age declaration.
func acceptRequest(token string) *casav1.AcceptInvitationRequest {
	return &casav1.AcceptInvitationRequest{
		RequestContext: commandRC(), Source: tokenSource(token), DisplayName: "Budi", AgeDeclarationVersion: 1,
	}
}

func (h *householdHarness) accept(m member, req *casav1.AcceptInvitationRequest) (*casav1.AcceptInvitationResponse, error) {
	return h.households.AcceptInvitation(bearer(m.token), req)
}

func (h *householdHarness) acceptToken(m member, token string) error {
	_, err := h.accept(m, acceptRequest(token))
	return err
}

func (h *householdHarness) mustAccept(t *testing.T, m member, token string) {
	t.Helper()
	if err := h.acceptToken(m, token); err != nil {
		t.Fatal(err)
	}
}

func (h *householdHarness) withdraw(m member, req *casav1.WithdrawAcceptanceRequest) error {
	_, err := h.households.WithdrawAcceptance(bearer(m.token), req)
	return err
}

func withdrawRequest() *casav1.WithdrawAcceptanceRequest {
	return &casav1.WithdrawAcceptanceRequest{RequestContext: commandRC()}
}

// tokenRejection is the answer of StartSignUp to a setup token it never
// issued: the one rejection every invitation token failure must match byte
// for byte (SEC-30).
func (h *householdHarness) tokenRejection(t *testing.T) rejection {
	t.Helper()
	unknown, _ := identity.NewToken()
	err := h.startSignUp(unknown, "nobody@example.test")
	wantCodes(t, err, connect.CodePermissionDenied, []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED})
	return rejectionOf(t, err)
}

func wantRejection(t *testing.T, err error, want rejection, what string) {
	t.Helper()
	if err == nil {
		t.Fatalf("%s: accepted, want %v", what, want)
	}
	if got := rejectionOf(t, err); !got.equal(want) {
		t.Fatalf("%s: %v, want the same answer as %v", what, err, want)
	}
}

func tokenHash(t *testing.T, token string) []byte {
	t.Helper()
	sum := sha256.Sum256(decodeToken(t, token))
	return sum[:]
}

func TestAcceptInvitation(t *testing.T) {
	// FR-FAM-003/AC-3, the use side: a cancelled invitation cannot be used
	// (the closing and recording are proved in TestInvitations, "no
	// notification" in step 4).
	t.Run("SEC-30 every invalid invitation gets the one token rejection", func(t *testing.T) {
		h := startHousehold(t)
		b := h.signUpMember(t, "b@example.test")
		cancelled := h.invited(t, "c1@example.test")
		wantOK(t, h.cancelInvitation(cancelled.inviter, cancelled.household, cancelled.id))
		due := h.invited(t, "c2@example.test")
		if _, err := h.admin.Exec(t.Context(), "UPDATE invitation SET expires_at = $2 WHERE id = $1", due.id, h.clock.Now()); err != nil {
			t.Fatal(err)
		}
		replaced := h.invited(t, "c3@example.test")
		h.mustIssueLink(t, replaced.inviter, replaced.household, replaced.id)
		used := h.invited(t, "c4@example.test")
		h.mustAccept(t, h.signUpMember(t, "d@example.test"), used.token)
		never, _ := identity.NewToken()
		want := h.tokenRejection(t)
		before := h.fingerprint(t)
		for _, c := range []struct{ name, token string }{
			{"empty", ""}, {"malformed", "abc"}, {"never issued", never}, {"cancelled", cancelled.token},
			{"expired", due.token}, {"replaced", replaced.token}, {"accepted", used.token},
		} {
			_, err := h.preview(b, c.token)
			wantRejection(t, err, want, "PreviewInvitation "+c.name)
			wantRejection(t, h.acceptToken(b, c.token), want, "AcceptInvitation "+c.name)
		}
		if h.fingerprint(t) != before {
			t.Fatal("a rejected invitation token changed a table")
		}
	})

	t.Run("FR-FAM-004/AC-8 an old invitation used when the household has two adults is refused and makes no third member", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		h.addMember(t, inv.household, uuid4()) // test data: in 1.6 F01 never leaves an invitation beside two adults
		b := h.signUpMember(t, "b@example.test")
		want := h.tokenRejection(t)
		row := h.invitationRow(t, inv.id)
		wantRejection(t, h.acceptToken(b, inv.token), want, "AcceptInvitation into a full household")
		if n := h.count(t, "SELECT count(*) FROM membership WHERE household_id = $1", inv.household); n != 2 {
			t.Fatalf("%d memberships, want 2", n)
		}
		if got := h.invitationRow(t, inv.id); got != row {
			t.Fatalf("invitation = %q, want %q", got, row)
		}
	})

	// 10.1 §8.1 row "Pemakaian token"; 10.1 §10 note of 14 November 2026:
	// preview and F02 count against the limit of sign-up's token use.
	t.Run("SEC-30 failed invitation uses, by preview or accept, hit the rate limit", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		for i := range 20 {
			var err error
			if i%2 == 0 {
				_, err = h.preview(b, "abc")
			} else {
				err = h.acceptToken(b, "abc")
			}
			wantCodes(t, err, connect.CodePermissionDenied, []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED})
		}
		_, err := h.preview(b, inv.token)
		wantCodes(t, err, connect.CodeResourceExhausted, rateLimited)
		wantCodes(t, h.acceptToken(b, inv.token), connect.CodeResourceExhausted, rateLimited)
		unknown, _ := identity.NewToken()
		wantCodes(t, h.startSignUp(unknown, "x@example.test"), connect.CodeResourceExhausted, rateLimited)
		if got := h.invitationRow(t, inv.id); !strings.HasPrefix(got, "open 0 - ") {
			t.Fatalf("invitation = %q, want open", got)
		}
	})

	// FR-FAM-004/AC-1, the server part: what the preview holds. The order of
	// the screen and the operator notice in it are 1.10.
	t.Run("the invitee reads the household and inviter names before accepting, and no money data", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		before := h.fingerprint(t)
		res, err := h.preview(b, inv.token)
		if err != nil {
			t.Fatal(err)
		}
		if res.GetHouseholdName() != "Rumah Kita" || res.GetInviterDisplayName() != "Ayu" || res.GetFormerMemberNotice() {
			t.Fatalf("preview = %v, want Rumah Kita, Ayu, no former-member notice", res)
		}
		wantFields(t, res.ProtoReflect().Descriptor(), "household_name", "inviter_display_name", "former_member_notice")
		if h.fingerprint(t) != before {
			t.Fatal("a preview changed a table")
		}
	})

	// FR-FAM-004/AC-2, the part of this task: the acceptance waits and opens
	// nothing. The inviter seeing it is Task 7 (GetOpenInvitation); N-24 is
	// step 4.
	t.Run("an acceptance waits for the inviter and opens nothing", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		h.clock.advance(time.Minute)
		req := acceptRequest(inv.token)
		res, err := h.accept(b, req)
		if err != nil {
			t.Fatal(err)
		}
		if res.GetHouseholdId() != inv.household || res.GetInvitationId() != inv.id {
			t.Fatalf("response = %v, want %s %s", res, inv.household, inv.id)
		}
		if got, want := h.invitationRow(t, inv.id), "accepted 1 - "; got != want {
			t.Fatalf("invitation = %q, want %q (no token)", got, want)
		}
		var invitee, name string
		if err := h.admin.QueryRow(t.Context(), "SELECT invitee_account_id::text, invitee_display_name FROM invitation WHERE id = $1",
			inv.id).Scan(&invitee, &name); err != nil {
			t.Fatal(err)
		}
		if invitee != b.account || name != "Budi" {
			t.Fatalf("invitee = %s %q, want %s Budi", invitee, name, b.account)
		}
		events := h.invitationEvents(t, inv.id)
		want := fmt.Sprintf("1 invitee - - accept_invitation open accepted %s %s",
			strings.ToLower(req.GetRequestContext().GetIdempotencyKey()), utc(h.clock.Now()))
		if len(events) != 2 || events[1] != want {
			t.Fatalf("events = %q, want the creation and %q", events, want)
		}
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b.account); n != 0 {
			t.Fatalf("%d memberships of the invitee, want 0", n)
		}
		if got := h.memberships(t, b); len(got) != 0 {
			t.Fatalf("the invitee lists %v, want nothing", got)
		}
		// INV-A48 (c): every household RPC with the household's selector
		// answers the invitee as for a household that does not exist.
		for _, c := range selectorCases {
			wantHiddenNotFound(t, c.call(h, b, inv.household, inv.id))
		}
	})

	t.Run("SEC-38 (c) an account active in any household is refused before the token is read", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		c := h.signUpMember(t, "c@example.test")
		h.mustCreate(t, c)
		before := h.fingerprint(t)
		valid := h.acceptToken(c, inv.token)
		wantCodes(t, valid, connect.CodeFailedPrecondition, alreadyMember)
		for _, token := range []string{"abc", ""} {
			wantRejection(t, h.acceptToken(c, token), rejectionOf(t, valid), fmt.Sprintf("token %q", token))
		}
		// The limiter is not touched: 20 more refusals leave the token use
		// of the next caller free.
		for range 20 {
			wantCodes(t, h.acceptToken(c, "abc"), connect.CodeFailedPrecondition, alreadyMember)
		}
		if h.fingerprint(t) != before {
			t.Fatal("a refused acceptance changed a table")
		}
		if _, err := h.preview(h.signUpMember(t, "b@example.test"), inv.token); err != nil {
			t.Fatalf("preview after the refusals: %v", err)
		}
	})

	t.Run("SEC-32 accepting needs the acknowledged notice and the age declaration, checked before the token", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		u := h.signUpUnread(t, "u@example.test")
		h.notice(t, u)
		b := h.signUpMember(t, "b@example.test")
		before := h.fingerprint(t)
		for _, token := range []string{inv.token, "abc"} {
			wantCodes(t, h.acceptToken(u, token), connect.CodeFailedPrecondition, noticeUnread)
		}
		var first error
		for _, version := range []uint32{0, 2} {
			for _, token := range []string{inv.token, "abc"} {
				req := acceptRequest(token)
				req.AgeDeclarationVersion = version
				_, err := h.accept(b, req)
				wantCodes(t, err, connect.CodeFailedPrecondition, ageDeclarationRequired)
				if first == nil {
					first = err
				}
				wantRejection(t, err, rejectionOf(t, first), fmt.Sprintf("version %d, token %q", version, token))
			}
		}
		if h.fingerprint(t) != before {
			t.Fatal("a refused acceptance changed a table")
		}
		// The token was never used up (SEC-80 note of 14 November 2026).
		h.mustAccept(t, b, inv.token)
	})

	// FR-FAM-004/AC-12, the server part (with TestCreateHousehold's
	// subtest for creating); the statement asked on the screen is 1.10.
	t.Run("creating and accepting both need the age declaration, stored without a birth date", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		req := acceptRequest(inv.token)
		req.AgeDeclarationVersion = 0
		_, err := h.accept(b, req)
		wantCodes(t, err, connect.CodeFailedPrecondition, ageDeclarationRequired)
		if n := h.count(t, "SELECT count(*) FROM consent_record WHERE account_id = $1", b.account); n != 0 {
			t.Fatalf("%d consent records of the refused invitee, want 0", n)
		}
		if n := h.count(t, "SELECT count(*) FROM invitation WHERE invitee_account_id = $1", b.account); n != 0 {
			t.Fatalf("%d acceptances of the refused invitee, want 0", n)
		}
		h.mustAccept(t, b, inv.token)
		for _, m := range []member{inv.inviter, b} {
			var kind string
			var version int32
			if err := h.admin.QueryRow(t.Context(), "SELECT kind, text_version FROM consent_record WHERE account_id = $1",
				m.account).Scan(&kind, &version); err != nil {
				t.Fatal(err) // also fails on more than one row
			}
			if kind != "adult_age" || version != 1 {
				t.Fatalf("consent record = %s %d, want adult_age 1", kind, version)
			}
		}
		if n := h.count(t, "SELECT count(*) FROM information_schema.columns WHERE table_schema = 'public' AND column_name ILIKE '%birth%'"); n != 0 {
			t.Fatalf("%d birth columns, want 0", n)
		}
	})

	t.Run("SEC-28 an accepted token cannot be used again", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		h.mustAccept(t, h.signUpMember(t, "b@example.test"), inv.token)
		c := h.signUpMember(t, "c@example.test")
		want := h.tokenRejection(t)
		_, err := h.preview(c, inv.token)
		wantRejection(t, err, want, "PreviewInvitation")
		wantRejection(t, h.acceptToken(c, inv.token), want, "AcceptInvitation")
	})

	// Plan 1.6 decision 21: the interleaving is run in order, without
	// goroutines.
	t.Run("SEC-30 a link reissued between lookup and accept is refused uniformly", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		account, old := uuidOf(t, b.account), tokenHash(t, inv.token)
		householdID, invitationID, err := h.service.FindInvitation(t.Context(), old)
		if err != nil {
			t.Fatal(err)
		}
		fresh := h.mustIssueLink(t, inv.inviter, inv.household, inv.id)
		req := acceptRequest(inv.token)
		key := req.GetRequestContext().GetIdempotencyKey()
		_, err = command.RunForAccount(t.Context(), h.pool, database.Scope{HouseholdID: householdID, AccountID: account},
			casav1connect.HouseholdServiceAcceptInvitationProcedure, key, req,
			func(ctx context.Context, tx pgx.Tx) (*casav1.AcceptInvitationResponse, error) {
				return &casav1.AcceptInvitationResponse{}, h.service.Accept(ctx, tx, account, invitationID, old, "Budi", 1, uuidOf(t, key))
			})
		if !errors.Is(err, household.ErrInvitationRejected) {
			t.Fatalf("Accept with the replaced hash = %v, want ErrInvitationRejected", err)
		}
		if got, want := h.invitationRow(t, inv.id), fmt.Sprintf("open 0 - %x", tokenHash(t, fresh.GetInvitationToken())); got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		if n := h.count(t, "SELECT count(*) FROM idempotency_record WHERE command = $1", casav1connect.HouseholdServiceAcceptInvitationProcedure); n != 0 {
			t.Fatalf("%d idempotency records of AcceptInvitation, want 0", n)
		}
	})

	t.Run("FR-FAM-004/AC-5 a withdrawn acceptance leaves only the invitation's events", func(t *testing.T) {
		wantFields(t, (&casav1.WithdrawAcceptanceRequest{}).ProtoReflect().Descriptor(), "request_context")
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		h.mustAccept(t, b, inv.token)
		h.clock.advance(time.Minute)
		req := withdrawRequest()
		wantOK(t, h.withdraw(b, req))
		if got, want := h.invitationRow(t, inv.id), "withdrawn 2 "+utc(h.clock.Now())+" "; got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		events := h.invitationEvents(t, inv.id)
		want := fmt.Sprintf("2 invitee - - withdraw_acceptance accepted withdrawn %s %s",
			strings.ToLower(req.GetRequestContext().GetIdempotencyKey()), utc(h.clock.Now()))
		if len(events) != 3 || events[2] != want {
			t.Fatalf("events = %q, want creation, acceptance and %q", events, want)
		}
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b.account); n != 0 {
			t.Fatalf("%d memberships of the invitee, want 0", n)
		}
		// The invitee's identity leaves the invitation; its events keep their
		// actor kind (lead ruling on the Task 6 review).
		if n := h.count(t, "SELECT count(*) FROM invitation WHERE id = $1 AND invitee_account_id IS NULL AND invitee_display_name IS NULL",
			inv.id); n != 1 {
			t.Fatal("the withdrawn invitation still holds the invitee's account or display name")
		}
		for table, want := range map[string]int{"membership": 1, "approval_authority": 1, "household_rule": 1, "relationship": 0} {
			if n := h.count(t, "SELECT count(*) FROM "+table+" WHERE household_id = $1", inv.household); n != want {
				t.Fatalf("%d rows in %s, want %d", n, table, want)
			}
		}
		// A resend gets the stored answer (ADR-005).
		wantOK(t, h.withdraw(b, req))
		// The inviter can invite again (06 §3.6 F01: no live invitation).
		h.mustInvite(t, inv.inviter, inv.household)
	})

	// Lead ruling of 14 November 2026; SEC-151 (1); 10.1 §10 note.
	t.Run("withdrawing is protective and never held", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		h.mustAccept(t, b, inv.token)
		h.reopen("Nama Lain", household.PublishedNotices) // a new version for b
		if !h.held(t, b) {
			t.Fatal("b is not held")
		}
		wantOK(t, h.withdraw(b, withdrawRequest()))
		if got := h.invitationRow(t, inv.id); !strings.HasPrefix(got, "withdrawn 2 ") {
			t.Fatalf("invitation = %q, want withdrawn", got)
		}
	})

	t.Run("only the invitee can withdraw, and only while accepted", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		c := h.signUpMember(t, "c@example.test")
		wantHiddenNotFound(t, h.withdraw(b, withdrawRequest())) // nothing accepted yet
		h.mustAccept(t, b, inv.token)
		before := h.fingerprint(t)
		for _, m := range []member{inv.inviter, c} {
			wantHiddenNotFound(t, h.withdraw(m, withdrawRequest()))
		}
		if h.fingerprint(t) != before {
			t.Fatal("a refused withdrawal changed a table")
		}
		wantOK(t, h.withdraw(b, withdrawRequest()))
		wantHiddenNotFound(t, h.withdraw(b, withdrawRequest()))
	})

	// Plan 1.6 decision 21; "Fokus tinjauan" item 4.
	t.Run("two accounts accepting one token at once: exactly one", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		takers := []member{h.signUpMember(t, "b@example.test"), h.signUpMember(t, "c@example.test")}
		want := h.tokenRejection(t)
		// Barrier: both acceptances wait on the invitation table lock.
		waitBlocked, release := holdTables(t, h.db, "invitation")
		errs := make([]error, 2)
		var wg sync.WaitGroup
		for i := range errs {
			wg.Go(func() { errs[i] = h.acceptToken(takers[i], inv.token) })
		}
		waitBlocked(2)
		release()
		wg.Wait()
		var winner member
		var rejected int
		for i, err := range errs {
			if err == nil {
				winner = takers[i]
				continue
			}
			wantRejection(t, err, want, "the losing acceptance")
			rejected++
		}
		if rejected != 1 {
			t.Fatalf("%d rejected, want 1 (%v)", rejected, errs)
		}
		if n := h.count(t, "SELECT count(*) FROM invitation WHERE state = 'accepted' AND invitee_account_id = $1", winner.account); n != 1 {
			t.Fatal("the invitation is not accepted by the winner")
		}
		if n := h.count(t, "SELECT count(*) FROM object_event WHERE object_kind = 'invitation'"); n != 2 {
			t.Fatalf("%d invitation events, want 2", n)
		}
		// The losing transaction wrote nothing: one consent record and one
		// idempotency record, both the winner's.
		if n := h.count(t, "SELECT count(*) FROM consent_record WHERE account_id = ANY($1::uuid[])",
			[]string{takers[0].account, takers[1].account}); n != 1 {
			t.Fatalf("%d consent records of the two takers, want 1", n)
		}
		if n := h.count(t, "SELECT count(*) FROM idempotency_record WHERE command = $1", casav1connect.HouseholdServiceAcceptInvitationProcedure); n != 1 {
			t.Fatalf("%d idempotency records of AcceptInvitation, want 1", n)
		}
	})

	t.Run("one account accepting two households' invitations at once: exactly one", func(t *testing.T) {
		h := startHousehold(t)
		invs := []invited{h.invited(t, "a1@example.test"), h.invited(t, "a2@example.test")}
		b := h.signUpMember(t, "b@example.test")
		rows := []string{h.invitationRow(t, invs[0].id), h.invitationRow(t, invs[1].id)}
		// 18 failed uses first: the two racing uses make 20, the winner
		// takes its own back, and one more failure reaches the limit only
		// if the loser counted.
		for range 18 {
			_, err := h.preview(b, "abc")
			wantCodes(t, err, connect.CodePermissionDenied, []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED})
		}
		// Barrier: both acceptances wait on the invitation table lock.
		waitBlocked, release := holdTables(t, h.db, "invitation")
		errs := make([]error, 2)
		var wg sync.WaitGroup
		for i := range errs {
			wg.Go(func() { errs[i] = h.acceptToken(b, invs[i].token) })
		}
		waitBlocked(2)
		release()
		wg.Wait()
		var losers []invited
		for i, err := range errs {
			if err != nil {
				wantCodes(t, err, connect.CodeFailedPrecondition, alreadyMember)
				if got := h.invitationRow(t, invs[i].id); got != rows[i] {
					t.Fatalf("the loser's invitation = %q, want unchanged %q", got, rows[i])
				}
				losers = append(losers, invs[i])
			}
		}
		if len(losers) != 1 {
			t.Fatalf("errors = %v, want exactly one ALREADY_HOUSEHOLD_MEMBER", errs)
		}
		if n := h.count(t, "SELECT count(*) FROM invitation WHERE state = 'accepted'"); n != 1 {
			t.Fatalf("%d accepted invitations, want 1", n)
		}
		for query, want := range map[string]int{
			"SELECT count(*) FROM object_event WHERE command = 'accept_invitation'":           1,
			"SELECT count(*) FROM consent_record WHERE account_id = '" + b.account + "'":      1,
			"SELECT count(*) FROM idempotency_record WHERE command LIKE '%/AcceptInvitation'": 1,
		} {
			if n := h.count(t, query); n != want {
				t.Fatalf("%s = %d, want %d", query, n, want)
			}
		}
		_, err := h.preview(b, "abc")
		wantCodes(t, err, connect.CodePermissionDenied, []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED})
		_, err = h.preview(b, losers[0].token)
		wantCodes(t, err, connect.CodeResourceExhausted, rateLimited)
	})

	// 09.2 §5.1, note on decision 20: both lookups saw nothing; the F04
	// that follows is TestConfirmInvitee's "joined elsewhere" case (Task 7).
	t.Run("an account creating a household while accepting an invitation", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		// Barrier: the membership insert and the invitation update wait.
		waitBlocked, release := holdTables(t, h.db, "membership", "invitation")
		errs := make([]error, 2)
		var wg sync.WaitGroup
		wg.Go(func() { _, errs[0] = h.create(b, newHousehold()) })
		wg.Go(func() { errs[1] = h.acceptToken(b, inv.token) })
		waitBlocked(2)
		release()
		wg.Wait()
		for _, err := range errs {
			wantOK(t, err)
		}
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b.account); n != 1 {
			t.Fatalf("%d memberships of b, want 1", n)
		}
		if n := h.count(t, "SELECT count(*) FROM invitation WHERE state = 'accepted' AND invitee_account_id = $1", b.account); n != 1 {
			t.Fatalf("%d acceptances of b, want 1", n)
		}
	})
}

// expiredAcceptance is the fixture of TestExpiredAcceptance: b accepted
// h1's invitation, then the clock moved by after.
type expiredAcceptance struct {
	h  *householdHarness
	h1 invited
	b  member
}

func startExpiredAcceptance(t *testing.T, after time.Duration) expiredAcceptance {
	t.Helper()
	h := startHousehold(t)
	f := expiredAcceptance{h: h, h1: h.invited(t, "a@example.test"), b: h.signUpMember(t, "b@example.test")}
	h.mustAccept(t, f.b, f.h1.token)
	h.clock.advance(after)
	h.signInAgain(t, &f.b, "b@example.test")
	return f
}

// wantExpired checks that h1's acceptance was written expired, at its
// expiry, by one system event.
func (f expiredAcceptance) wantExpired(t *testing.T) {
	t.Helper()
	at := utc(f.h1.expiresAt)
	if got, want := f.h.invitationRow(t, f.h1.id), "expired 2 "+at+" "; got != want {
		t.Fatalf("invitation = %q, want %q", got, want)
	}
	events := f.h.invitationEvents(t, f.h1.id)
	want := "2 system - invitation.expired expire_invitation accepted expired - " + at
	if len(events) != 3 || events[2] != want || f.h.count(t, "SELECT count(*) FROM object_event WHERE rule = 'invitation.expired'") != 1 {
		t.Fatalf("events = %q, want one expiry %q", events, want)
	}
}

// TestExpiredAcceptance is the architect's X3 ruling (06 §3.6 "Gugur karena
// waktu"; 09.2 §5.1 note X3): token-free paths write the expiry first, token
// paths never write. Fixture: b accepted h1's invitation, clock + 7 d.
func TestExpiredAcceptance(t *testing.T) {
	t.Run("an expired own acceptance no longer blocks creating a household", func(t *testing.T) {
		f := startExpiredAcceptance(t, invitationLifetime)
		if _, err := f.h.create(f.b, newHousehold()); err != nil {
			t.Fatal(err)
		}
		f.wantExpired(t)
	})

	t.Run("an expired own acceptance no longer blocks accepting another invitation", func(t *testing.T) {
		f := startExpiredAcceptance(t, invitationLifetime)
		h2 := f.h.invited(t, "c@example.test")
		f.h.mustAccept(t, f.b, h2.token)
		f.wantExpired(t)
		if got := f.h.invitationRow(t, h2.id); got != "accepted 1 - " {
			t.Fatalf("h2's invitation = %q, want accepted", got)
		}
	})

	t.Run("a live own acceptance answers already a member", func(t *testing.T) {
		f := startExpiredAcceptance(t, 24*time.Hour)
		h2 := f.h.invited(t, "c@example.test")
		before := f.h.fingerprint(t)
		_, err := f.h.create(f.b, newHousehold())
		wantCodes(t, err, connect.CodeFailedPrecondition, alreadyMember)
		wantCodes(t, f.h.acceptToken(f.b, h2.token), connect.CodeFailedPrecondition, alreadyMember)
		if f.h.fingerprint(t) != before {
			t.Fatal("a refused command changed a table")
		}
	})

	// 10.1 §10 note X3 (2): the expiry is written by the caller's own
	// lookup, before the age declaration and the token.
	t.Run("SEC-30 the own expiry write and its answer do not depend on the token", func(t *testing.T) {
		var written []string
		for _, token := range []string{"valid", "abc", ""} {
			f := startExpiredAcceptance(t, invitationLifetime)
			h2 := f.h.invited(t, "c@example.test")
			if token == "valid" {
				token = h2.token
			}
			_ = f.h.acceptToken(f.b, token)
			f.wantExpired(t)
			events := f.h.invitationEvents(t, f.h1.id)
			written = append(written, f.h.invitationRow(t, f.h1.id)+" "+events[len(events)-1])
		}
		if written[0] != written[1] || written[1] != written[2] {
			t.Fatalf("expiry writes differ by token: %q", written)
		}
		f := startExpiredAcceptance(t, 24*time.Hour)
		h2 := f.h.invited(t, "c@example.test")
		valid := f.h.acceptToken(f.b, h2.token)
		wantCodes(t, valid, connect.CodeFailedPrecondition, alreadyMember)
		for _, token := range []string{"abc", ""} {
			wantRejection(t, f.h.acceptToken(f.b, token), rejectionOf(t, valid), fmt.Sprintf("token %q", token))
		}
	})

	t.Run("withdrawing an expired acceptance answers state changed and writes the expiry", func(t *testing.T) {
		f := startExpiredAcceptance(t, invitationLifetime)
		wantCodes(t, f.h.withdraw(f.b, withdrawRequest()), connect.CodeFailedPrecondition, stateChanged)
		f.wantExpired(t)
		wantHiddenNotFound(t, f.h.withdraw(f.b, withdrawRequest()))
	})

	t.Run("a token path never writes the expiry", func(t *testing.T) {
		f := startExpiredAcceptance(t, 0)
		h3 := f.h.invited(t, "c@example.test")
		f.h.clock.advance(invitationLifetime)
		f.h.signInAgain(t, &f.b, "b@example.test")
		d := f.h.signUpMember(t, "d@example.test")
		want := f.h.tokenRejection(t)
		rows := []string{f.h.invitationRow(t, h3.id), f.h.invitationRow(t, f.h1.id)}
		before := f.h.fingerprint(t)
		for _, m := range []member{d, f.b} {
			_, err := f.h.preview(m, h3.token)
			wantRejection(t, err, want, "PreviewInvitation")
		}
		wantRejection(t, f.h.acceptToken(d, h3.token), want, "AcceptInvitation")
		if got := []string{f.h.invitationRow(t, h3.id), f.h.invitationRow(t, f.h1.id)}; !slices.Equal(got, rows) {
			t.Fatalf("invitations = %q, want unchanged %q", got, rows)
		}
		if f.h.fingerprint(t) != before {
			t.Fatal("a token path changed a table")
		}
	})
}
