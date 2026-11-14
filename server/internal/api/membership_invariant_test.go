package api_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect/v2"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

// The actors of the INV-A48 space (plan 1.6 decision 19).
const (
	inviter     = "inviter"
	invitee     = "invitee"
	otherMember = "other member"
	outside     = "outside account"
)

// invitationStates are the seven states of invitation.state (09.2 §5.1).
var invitationStates = []string{"open", "accepted", "confirmed", "cancelled", "rejected", "withdrawn", "expired"}

// membershipCommands are the seven commands of decision 19: F01,
// IssueInvitationLink, F03, F02, F05 by the invitee, F04, F05 by the
// inviter.
var membershipCommands = []string{
	"CreateInvitation", "IssueInvitationLink", "CancelInvitation", "AcceptInvitation",
	"WithdrawAcceptance", "ConfirmInvitee", "RejectInvitee",
}

// membershipCase is one point of the space: the invitation's state, whether
// the clock is past its expiry, the household's full adult members, the
// command and its actor.
type membershipCase struct {
	state   string
	after   bool
	adults  int
	command string
	actor   string
}

func (c membershipCase) String() string {
	when := "before"
	if c.after {
		when = "after"
	}
	return fmt.Sprintf("%s %s expiry, %d adults: %s by %s", c.state, when, c.adults, c.command, c.actor)
}

// writesExpiry: the command writes the expiry of a live invitation past its
// expiry before it runs (plan 1.6 decision 12): a token-free path of a
// member of the household, or the invitee's own acceptance check (F02,
// WithdrawAcceptance). A token path never writes.
func (c membershipCase) writesExpiry() bool {
	if !c.after || (c.state != "open" && c.state != "accepted") {
		return false
	}
	switch c.actor {
	case inviter, otherMember:
		return c.command != "AcceptInvitation" && c.command != "WithdrawAcceptance"
	case invitee:
		return c.state == "accepted" && (c.command == "AcceptInvitation" || c.command == "WithdrawAcceptance")
	}
	return false
}

// seen is the state the command meets: "lapsing" when it writes the expiry
// first, "due" for a live invitation past its expiry it does not write.
func (c membershipCase) seen() string {
	switch {
	case c.writesExpiry():
		return "lapsing"
	case c.after && (c.state == "open" || c.state == "accepted"):
		return "due"
	}
	return c.state
}

// membershipOutcome is one row of the outcome table: the first row whose
// command, actors (nil: any), seen states ("": any) and adults (0: any)
// match gives the answer and, for "ok", the invitation's new state ("":
// unchanged; "+open": a new open invitation beside it).
type membershipOutcome struct {
	command string
	actors  []string
	seen    string
	adults  int
	answer  string
	to      string
}

// membershipOutcomes is 06 §3.6 F01 to F05 with plan 1.6 decisions 10 to 14
// and 17.
var membershipOutcomes = []membershipOutcome{
	// F01: members only; full, then pending (FR-FAM-003/AC-2, AC-5).
	{"CreateInvitation", []string{invitee, outside}, "", 0, "not_found", ""},
	{"CreateInvitation", nil, "", 2, "household_full", ""},
	{"CreateInvitation", nil, "open accepted", 0, "invitation_pending", ""},
	{"CreateInvitation", nil, "", 0, "ok", "+open"},
	// IssueInvitationLink, F03: only the inviter, only open (10.1 SEC-28
	// note Q7 (a); FR-FAM-003 note of 14 November 2026).
	{"IssueInvitationLink", []string{inviter}, "open", 0, "ok", ""},
	{"IssueInvitationLink", []string{inviter}, "", 0, "state_changed", ""},
	{"IssueInvitationLink", nil, "", 0, "not_found", ""},
	{"CancelInvitation", []string{inviter}, "open", 0, "ok", "cancelled"},
	{"CancelInvitation", []string{inviter}, "", 0, "state_changed", ""},
	{"CancelInvitation", nil, "", 0, "not_found", ""},
	// F04: only the inviter, only accepted, never a third adult.
	{"ConfirmInvitee", []string{inviter}, "accepted", 1, "ok", "confirmed"},
	{"ConfirmInvitee", []string{inviter}, "", 0, "state_changed", ""},
	{"ConfirmInvitee", nil, "", 0, "not_found", ""},
	// F05 reject_invitee: only the inviter, only accepted.
	{"RejectInvitee", []string{inviter}, "accepted", 0, "ok", "rejected"},
	{"RejectInvitee", []string{inviter}, "", 0, "state_changed", ""},
	{"RejectInvitee", nil, "", 0, "not_found", ""},
	// F05 withdraw_acceptance: only the invitee's own acceptance.
	{"WithdrawAcceptance", []string{invitee}, "accepted", 0, "ok", "withdrawn"},
	{"WithdrawAcceptance", []string{invitee}, "lapsing", 0, "state_changed", ""},
	{"WithdrawAcceptance", nil, "", 0, "not_found", ""},
	// F02: a member anywhere, then the own live acceptance, then the token
	// (decision 13); a full household is the token rejection (AC-8).
	{"AcceptInvitation", []string{inviter, otherMember}, "", 0, "already_member", ""},
	{"AcceptInvitation", []string{invitee}, "accepted", 0, "already_member", ""},
	{"AcceptInvitation", nil, "open", 1, "ok", "accepted"},
	{"AcceptInvitation", nil, "", 0, "token_rejected", ""},
}

func (c membershipCase) want(t *testing.T) membershipOutcome {
	t.Helper()
	for _, o := range membershipOutcomes {
		if o.command == c.command && (o.actors == nil || slices.Contains(o.actors, c.actor)) &&
			(o.seen == "" || slices.Contains(strings.Fields(o.seen), c.seen())) && (o.adults == 0 || o.adults == c.adults) {
			return o
		}
	}
	t.Fatalf("no outcome for %v", c)
	return membershipOutcome{}
}

// answerOf names a command's answer as the outcome table does.
func answerOf(t *testing.T, err error) string {
	t.Helper()
	if err == nil {
		return "ok"
	}
	if connect.CodeOf(err) == connect.CodeNotFound {
		return "not_found"
	}
	names := map[casav1.ErrorCode]string{
		casav1.ErrorCode_ERROR_CODE_HOUSEHOLD_FULL:           "household_full",
		casav1.ErrorCode_ERROR_CODE_INVITATION_PENDING:       "invitation_pending",
		casav1.ErrorCode_ERROR_CODE_STATE_CHANGED:            "state_changed",
		casav1.ErrorCode_ERROR_CODE_ALREADY_HOUSEHOLD_MEMBER: "already_member",
		casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED:           "token_rejected",
	}
	if codes := errorCodes(t, err); len(codes) == 1 && names[codes[0]] != "" {
		return names[codes[0]]
	}
	return err.Error()
}

// membershipFixture is one household of the space: the inviter a made it
// through the API; c is a second full adult as test data when adults is 2;
// b and d are accounts outside it, b being the invitee the invitation names.
type membershipFixture struct {
	h          *householdHarness
	household  string
	adults     int
	actors     map[string]member
	kept       []string // the accounts whose memberships the reset keeps
	aMember    string
	inviteeAcc string
}

func startMembershipFixture(t *testing.T, adults int) *membershipFixture {
	t.Helper()
	h := startHousehold(t)
	a := h.signUpMember(t, "a@example.test")
	hh := h.mustCreate(t, a)
	f := &membershipFixture{
		h: h, household: hh.GetHouseholdId(), adults: adults, aMember: hh.GetMemberId(),
		actors: map[string]member{
			inviter: a, invitee: h.signUpMember(t, "b@example.test"), outside: h.signUpMember(t, "d@example.test"),
		},
		kept: []string{a.account},
	}
	f.inviteeAcc = f.actors[invitee].account
	if adults == 2 {
		c := h.signUpMember(t, "c@example.test")
		h.addMember(t, f.household, c.account)
		f.actors[otherMember] = c
		f.kept = append(f.kept, c.account)
	}
	return f
}

// reset brings the household back to its members and one invitation of a
// in state, expiring an hour from now or an hour ago, with its creation
// event, and returns the invitation's id and a link token: the one its
// token_hash holds where the state keeps one (open, cancelled), else a
// token that finds nothing.
func (f *membershipFixture) reset(t *testing.T, state string, after bool) (invitationID, token string) {
	t.Helper()
	h := f.h
	for _, s := range []struct {
		sql  string
		args []any
	}{
		{"DELETE FROM object_event WHERE household_id = $1 AND NOT (object_kind = 'household' AND seq = 0)", []any{f.household}},
		{`DELETE FROM approval_authority WHERE household_id = $1 AND member_id IN
			(SELECT member_id FROM membership WHERE household_id = $1 AND account_id <> ALL($2::uuid[]))`, []any{f.household, f.kept}},
		{"DELETE FROM membership WHERE household_id = $1 AND account_id <> ALL($2::uuid[])", []any{f.household, f.kept}},
		{"DELETE FROM invitation WHERE household_id = $1", []any{f.household}},
		{"UPDATE household SET version = 0 WHERE id = $1", []any{f.household}},
	} {
		if _, err := h.admin.Exec(t.Context(), s.sql, s.args...); err != nil {
			t.Fatal(err)
		}
	}
	now := h.clock.Now()
	expiresAt := now.Add(time.Hour)
	if after {
		expiresAt = now.Add(-time.Hour)
	}
	createdAt := expiresAt.Add(-invitationLifetime)
	token, hash := identity.NewToken()
	shape := map[string]struct {
		hash    bool
		invitee bool
		version int
		closed  bool
	}{
		"open": {true, false, 0, false}, "accepted": {false, true, 1, false}, "confirmed": {false, true, 2, true},
		"cancelled": {true, false, 1, true}, "rejected": {false, true, 2, true}, "withdrawn": {false, false, 2, true},
		"expired": {false, true, 2, true},
	}[state]
	var tokenHash []byte
	if shape.hash {
		tokenHash = hash
	}
	var inviteeAccount, inviteeName, closedAt any
	if shape.invitee {
		inviteeAccount, inviteeName = f.inviteeAcc, "Budi"
	}
	if shape.closed {
		closedAt = createdAt.Add(time.Minute)
		if state == "expired" {
			closedAt = expiresAt
		}
	}
	if err := h.admin.QueryRow(t.Context(), `
		INSERT INTO invitation (household_id, inviter, token_hash, state, invitee_account_id, invitee_display_name,
		                        created_at, expires_at, closed_at, version)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id::text`,
		f.household, f.aMember, tokenHash, state, inviteeAccount, inviteeName, createdAt, expiresAt, closedAt, shape.version,
	).Scan(&invitationID); err != nil {
		t.Fatal(err)
	}
	if _, err := h.admin.Exec(t.Context(), `
		INSERT INTO object_event (household_id, object_kind, object_id, seq, actor_kind, actor, command, to_state, server_time)
		VALUES ($1, 'invitation', $2, 0, 'member', $3, 'invite', 'open', $4)`,
		f.household, invitationID, f.aMember, createdAt); err != nil {
		t.Fatal(err)
	}
	return invitationID, token
}

// call sends c.command as m.
func (f *membershipFixture) call(c membershipCase, m member, invitationID, token string) error {
	h := f.h
	switch c.command {
	case "CreateInvitation":
		_, err := h.invite(m, f.household)
		return err
	case "IssueInvitationLink":
		_, err := h.issueLink(m, f.household, invitationID)
		return err
	case "CancelInvitation":
		return h.cancelInvitation(m, f.household, invitationID)
	case "AcceptInvitation":
		return h.acceptToken(m, token)
	case "WithdrawAcceptance":
		return h.withdraw(m, withdrawRequest())
	case "ConfirmInvitee":
		_, err := h.confirm(m, confirmRequest(f.household, invitationID))
		return err
	case "RejectInvitee":
		return h.reject(m, rejectRequest(f.household, invitationID))
	}
	panic("unknown command " + c.command)
}

// rows is every membership and approval authority outside the invitee's,
// to show a command changed no adult's role, permission set, account or
// authority (INV-A48 (a)).
func (f *membershipFixture) rows(t *testing.T) string {
	t.Helper()
	var rows string
	if err := f.h.admin.QueryRow(t.Context(), `
		SELECT concat_ws(E'\n',
			(SELECT string_agg(row(m.*)::text, ',' ORDER BY m.household_id, m.member_id) FROM membership m
			  WHERE m.account_id IS DISTINCT FROM $1),
			(SELECT string_agg(row(x.*)::text, ',' ORDER BY x.household_id, x.member_id) FROM approval_authority x
			  WHERE x.member_id NOT IN (SELECT member_id FROM membership WHERE account_id = $1)))`,
		f.inviteeAcc).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	return rows
}

// check runs c on a fresh invitation and checks its answer, the
// invitation's state and events, and the memberships against the outcome
// table.
func (f *membershipFixture) check(t *testing.T, c membershipCase) {
	t.Helper()
	h := f.h
	invitationID, token := f.reset(t, c.state, c.after)
	before := f.rows(t)
	want := c.want(t)
	if got := answerOf(t, f.call(c, f.actors[c.actor], invitationID, token)); got != want.answer {
		t.Errorf("%v: answer %s, want %s", c, got, want.answer)
		return
	}
	state, events := c.state, 1
	if c.writesExpiry() {
		state, events = "expired", 2
	}
	if want.answer == "ok" && want.to != "" && want.to != "+open" {
		state, events = want.to, events+1
	}
	var gotState string
	if err := h.admin.QueryRow(t.Context(), "SELECT state FROM invitation WHERE id = $1", invitationID).Scan(&gotState); err != nil {
		t.Fatal(err)
	}
	got := h.invitationEvents(t, invitationID)
	if gotState != state || len(got) != events {
		t.Errorf("%v: invitation %s with %d events %q, want %s with %d", c, gotState, len(got), got, state, events)
	}
	if c.writesExpiry() && !strings.Contains(got[1], " system - invitation.expired expire_invitation ") {
		t.Errorf("%v: events %q, want the system expiry second", c, got)
	}
	invitations := 1
	if want.answer == "ok" && want.to == "+open" {
		invitations = 2
	}
	if n := h.count(t, "SELECT count(*) FROM invitation WHERE household_id = $1", f.household); n != invitations {
		t.Errorf("%v: %d invitations, want %d", c, n, invitations)
	}
	// INV-A48 (a), (b): only F04 by the inviter makes a membership, and
	// no command changes another adult's.
	joined := 0
	if want.to == "confirmed" {
		joined = 1
	}
	if n := h.count(t, `SELECT count(*) FROM membership m JOIN membership a ON a.household_id = m.household_id
		AND a.member_id = $2 AND a.role = m.role AND a.permission_set = m.permission_set
		WHERE m.account_id = $1 AND m.household_id = $3`, f.inviteeAcc, f.aMember, f.household); n != joined {
		t.Errorf("%v: %d memberships of the invitee, want %d", c, n, joined)
	}
	if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", f.actors[outside].account); n != 0 {
		t.Errorf("%v: the outside account has %d memberships", c, n)
	}
	if after := f.rows(t); after != before {
		t.Errorf("%v: another adult's membership or authority changed:\n%s\nwant\n%s", c, after, before)
	}
}

// TestMembershipTransitions is INV-A48 (a) to (c) over the whole F01..F05
// space (decision 19). (d) and (e) come with F06 to F08 (step 6); the
// stateful property version is 1.7.
func TestMembershipTransitions(t *testing.T) {
	// The other member exists only beside a second full adult: with one
	// adult the household has no other member.
	// INV-A48 (d), (e) and the F06, F08 parts of (a), (b) come with step 6.
	t.Run("INV-A48 (a), (b) only F04 by the inviter makes a membership, and no command changes another adult's", func(t *testing.T) {
		cases := 0
		for _, state := range invitationStates {
			for _, after := range []bool{false, true} {
				for _, adults := range []int{1, 2} {
					// One database per group, closed when the group ends.
					t.Run(fmt.Sprintf("%s after=%t adults=%d", state, after, adults), func(t *testing.T) {
						f := startMembershipFixture(t, adults)
						for _, command := range membershipCommands {
							for _, actor := range []string{inviter, invitee, otherMember, outside} {
								if _, ok := f.actors[actor]; !ok {
									continue
								}
								f.check(t, membershipCase{state: state, after: after, adults: adults, command: command, actor: actor})
								cases++
							}
						}
					})
				}
			}
		}
		if want := 7 * 2 * 7 * (3 + 4); cases != want {
			t.Fatalf("%d cases, want %d", cases, want)
		}
	})

	t.Run("INV-A48 (c) before activation the invitee reads nothing of the household", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		readsNothing := func(when string) {
			t.Helper()
			for _, c := range selectorCases {
				if err := c.call(h, b, inv.household, inv.id); connect.CodeOf(err) != connect.CodeNotFound {
					t.Fatalf("%s: %s answers the invitee %v, want NotFound", when, c.name, err)
				}
			}
			if got := h.memberships(t, b); len(got) != 0 {
				t.Fatalf("%s: the invitee lists %v, want nothing", when, got)
			}
			_, err := h.preview(b, inv.token)
			wantRejection(t, err, h.tokenRejection(t), when+": PreviewInvitation with the used token")
		}
		h.mustAccept(t, b, inv.token)
		readsNothing("accepted")
		wantOK(t, h.reject(inv.inviter, rejectRequest(inv.household, inv.id)))
		readsNothing("rejected")
		inv.id = h.mustInvite(t, inv.inviter, inv.household).GetInvitationId()
		inv.token = h.mustIssueLink(t, inv.inviter, inv.household, inv.id).GetInvitationToken()
		h.mustAccept(t, b, inv.token)
		wantOK(t, h.withdraw(b, withdrawRequest()))
		readsNothing("withdrawn")
		inv.id = h.mustInvite(t, inv.inviter, inv.household).GetInvitationId()
		inv.token = h.mustIssueLink(t, inv.inviter, inv.household, inv.id).GetInvitationToken()
		h.mustAccept(t, b, inv.token)
		h.clock.advance(invitationLifetime)
		h.signInAgain(t, &b, "b@example.test")
		readsNothing("past its expiry")
	})

	// 06 §2.5; 09.2 §5.1: seq is the object's version after the transition.
	t.Run("INV-A01 every transition writes exactly one event and old events stay identical", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		household := h.mustCreate(t, a).GetHouseholdId()
		b := h.signUpMember(t, "b@example.test")
		var invitation, token string
		events := func() []string {
			t.Helper()
			rows, err := h.admin.Query(t.Context(), "SELECT row(e.*)::text FROM object_event e ORDER BY object_kind, object_id, seq")
			if err != nil {
				t.Fatal(err)
			}
			defer rows.Close()
			var all []string
			for rows.Next() {
				var r string
				if err := rows.Scan(&r); err != nil {
					t.Fatal(err)
				}
				all = append(all, r)
			}
			return all
		}
		invite := func() error {
			res, err := h.invite(a, household)
			if err == nil {
				invitation = res.GetInvitationId()
				token = h.mustIssueLink(t, a, household, invitation).GetInvitationToken()
			}
			return err
		}
		accept := func() error { return h.acceptToken(b, token) }
		steps := []struct {
			name   string
			run    func() error
			events int
		}{
			{"invite", invite, 1},
			{"cancel", func() error { return h.cancelInvitation(a, household, invitation) }, 1},
			{"invite", invite, 1},
			{"accept", accept, 1},
			{"reject", func() error { return h.reject(a, rejectRequest(household, invitation)) }, 1},
			{"invite", invite, 1},
			{"accept", accept, 1},
			{"withdraw", func() error { return h.withdraw(b, withdrawRequest()) }, 1},
			{"invite", invite, 1},
			{"invite after the expiry", func() error {
				h.clock.advance(invitationLifetime)
				h.signInAgain(t, &a, "a@example.test")
				h.signInAgain(t, &b, "b@example.test")
				return invite()
			}, 2}, // the old invitation's expiry and the new one's creation
			{"accept", accept, 1},
			{"confirm", func() error {
				_, err := h.confirm(a, confirmRequest(household, invitation))
				return err
			}, 2}, // the invitation's confirmation and the household's member_joined
		}
		for _, s := range steps {
			before := events()
			if err := s.run(); err != nil {
				t.Fatalf("%s: %v", s.name, err)
			}
			after := events()
			if len(after) != len(before)+s.events {
				t.Fatalf("%s: %d new events, want %d", s.name, len(after)-len(before), s.events)
			}
			for _, e := range before {
				if !slices.Contains(after, e) {
					t.Fatalf("%s changed or removed the event %s", s.name, e)
				}
			}
		}
		// Each object's last seq is its version.
		if n := h.count(t, `SELECT count(*) FROM (
			SELECT i.version = max(e.seq) AS ok FROM invitation i JOIN object_event e ON e.object_id = i.id GROUP BY i.id, i.version
			UNION ALL
			SELECT h.version = max(e.seq) FROM household h JOIN object_event e ON e.object_id = h.id GROUP BY h.id, h.version) v
			WHERE NOT ok`); n != 0 {
			t.Fatalf("%d objects whose last seq is not their version", n)
		}
	})
}
