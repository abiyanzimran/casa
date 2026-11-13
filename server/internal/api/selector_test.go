package api_test

import (
	"slices"
	"testing"
	"time"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
)

// selectorCase is one session RPC that takes a household selector: the
// household_id of its request is checked against the caller's memberships
// and is never a proof of permission (SEC-38 (a); 09.2 §3 layer 1; plan 1.6
// decision 4).
type selectorCase struct {
	name string
	// onInvitation: the RPC also names an invitation of the household.
	onInvitation bool
	call         func(h *householdHarness, m member, household, invitation string) error
}

// selectorCases has one row per session RPC that uses the selector.
var selectorCases = []selectorCase{
	{"CreateInvitation", false, func(h *householdHarness, m member, household, _ string) error {
		_, err := h.invite(m, household)
		return err
	}},
	{"IssueInvitationLink", true, func(h *householdHarness, m member, household, invitation string) error {
		_, err := h.issueLink(m, household, invitation)
		return err
	}},
	{"CancelInvitation", true, func(h *householdHarness, m member, household, invitation string) error {
		return h.cancelInvitation(m, household, invitation)
	}},
	{"ConfirmInvitee", true, func(h *householdHarness, m member, household, invitation string) error {
		_, err := h.confirm(m, confirmRequest(household, invitation))
		return err
	}},
	{"RejectInvitee", true, func(h *householdHarness, m member, household, invitation string) error {
		return h.reject(m, rejectRequest(household, invitation))
	}},
	{"GetOpenInvitation", false, func(h *householdHarness, m member, household, _ string) error {
		_, err := h.openInvitation(m, household)
		return err
	}},
}

// twoHouseholds is account a, member of household x (made through the API)
// and of household y (test data: the pilot refuses a second household,
// FR-FAM-001 note of 14 November 2026), each with an open invitation a made
// through the API with that household's selector.
type twoHouseholds struct {
	h                        *householdHarness
	a                        member
	x, y                     string // household ids
	xMember, yMember         string
	xInvitation, yInvitation string
}

func startTwoHouseholds(t *testing.T) twoHouseholds {
	t.Helper()
	h := startHousehold(t)
	a := h.signUpMember(t, "a@example.test")
	x := h.mustCreate(t, a)
	dropOneHouseholdIndex(t, h.admin)
	var y, yMember string
	if err := h.admin.QueryRow(t.Context(), `
		WITH y AS (INSERT INTO household (id, name, time_zone, created_at)
		           VALUES (gen_random_uuid(), 'Rumah Y', 'Asia/Jakarta', $1) RETURNING id)
		INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
		SELECT id, $2, 'Ayu di Y', 'household_owner', 'two_equal_adults', $1 FROM y
		RETURNING household_id::text, member_id::text`,
		h.clock.Now().Add(time.Minute), a.account).Scan(&y, &yMember); err != nil {
		t.Fatal(err)
	}
	return twoHouseholds{
		h: h, a: a, x: x.GetHouseholdId(), y: y, xMember: x.GetMemberId(), yMember: yMember,
		xInvitation: h.mustInvite(t, a, x.GetHouseholdId()).GetInvitationId(),
		yInvitation: h.mustInvite(t, a, y).GetInvitationId(),
	}
}

// householdRows is every row of household's invitations, memberships and
// events, to show a request left the household unchanged.
func (h *householdHarness) householdRows(t *testing.T, household string) string {
	t.Helper()
	var rows string
	if err := h.admin.QueryRow(t.Context(), `
		SELECT concat_ws(E'\n',
			(SELECT string_agg(row(i.*)::text, ',' ORDER BY i.id) FROM invitation i WHERE household_id = $1),
			(SELECT string_agg(row(m.*)::text, ',' ORDER BY m.member_id) FROM membership m WHERE household_id = $1),
			(SELECT string_agg(row(e.*)::text, ',' ORDER BY e.object_id, e.seq) FROM object_event e WHERE household_id = $1))`,
		household).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	return rows
}

func TestHouseholdSelector(t *testing.T) {
	if len(selectorCases) == 0 {
		t.Fatal("selectorCases is empty")
	}

	t.Run("every RPC with a household selector has a selectorCases row", func(t *testing.T) {
		methods := casav1.File_casa_v1_household_proto.Services().ByName("HouseholdService").Methods()
		for i := range methods.Len() {
			m := methods.Get(i)
			if m.Input().Fields().ByName("household_id") == nil {
				continue
			}
			if !slices.ContainsFunc(selectorCases, func(c selectorCase) bool { return c.name == string(m.Name()) }) {
				t.Errorf("%s takes household_id but has no selectorCases row", m.Name())
			}
		}
	})

	t.Run("FR-FAM-020/AC-1 another household answers exactly as one that does not exist", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		h.mustCreate(t, a)
		b := h.signUpMember(t, "b@example.test")
		y := h.mustCreate(t, b)
		yInvitation := h.mustInvite(t, b, y.GetHouseholdId()).GetInvitationId()
		before := h.fingerprint(t)
		for _, c := range selectorCases {
			other := c.call(h, a, y.GetHouseholdId(), yInvitation)
			missing := c.call(h, a, uuid4(), yInvitation)
			wantHiddenNotFound(t, other)
			if !rejectionOf(t, other).equal(rejectionOf(t, missing)) {
				t.Errorf("%s: another household %v, a missing one %v", c.name, other, missing)
			}
		}
		if h.fingerprint(t) != before {
			t.Fatal("a request with another household's selector changed a table")
		}
	})

	t.Run("FR-FAM-020/AC-2 an account in two households sees each only through its own selector", func(t *testing.T) {
		f := startTwoHouseholds(t)
		h := f.h
		got := h.memberships(t, f.a)
		if len(got) != 2 {
			t.Fatalf("%d memberships, want 2", len(got))
		}
		for i, want := range []struct{ household, name, member string }{
			{f.x, "Rumah Kita", f.xMember + " Ayu"}, {f.y, "Rumah Y", f.yMember + " Ayu di Y"},
		} {
			var members []string
			for _, p := range got[i].GetMembers() {
				members = append(members, p.GetMemberId()+" "+p.GetDisplayName())
			}
			if got[i].GetHouseholdId() != want.household || got[i].GetHouseholdName() != want.name ||
				!slices.Equal(members, []string{want.member}) {
				t.Fatalf("membership %d = %v, want %s %s with [%s]", i, got[i], want.household, want.name, want.member)
			}
		}
		// Each invitation was made in the household of its selector, by
		// a's member there.
		for _, inv := range []struct{ id, household, inviter string }{
			{f.xInvitation, f.x, f.xMember}, {f.yInvitation, f.y, f.yMember},
		} {
			if n := h.count(t, "SELECT count(*) FROM invitation WHERE id = $1 AND household_id = $2 AND inviter = $3",
				inv.id, inv.household, inv.inviter); n != 1 {
				t.Fatalf("invitation %s is not in %s by %s", inv.id, inv.household, inv.inviter)
			}
		}
		// The own household's answer to each case, run in selectorCases
		// order on its open invitation: F01 finds it pending, the link is
		// issued, F03 cancels it, so F04 and the rejection find it closed.
		own := map[string]string{
			"CreateInvitation": "invitation_pending", "IssueInvitationLink": "ok", "CancelInvitation": "ok",
			"ConfirmInvitee": "state_changed", "RejectInvitee": "state_changed", "GetOpenInvitation": "ok",
		}
		for _, sel := range []struct{ household, invitation, other string }{
			{f.x, f.xInvitation, f.y}, {f.y, f.yInvitation, f.x},
		} {
			if got := h.mustOpenInvitation(t, f.a, sel.household); got.GetInvitationId() != sel.invitation {
				t.Fatalf("GetOpenInvitation with selector %s = %v, want its invitation %s", sel.household, got, sel.invitation)
			}
			for _, c := range selectorCases {
				want, ok := own[c.name]
				if !ok {
					t.Fatalf("%s has no own-household answer", c.name)
				}
				before := h.householdRows(t, sel.other)
				if got := answerOf(t, c.call(h, f.a, sel.household, sel.invitation)); got != want {
					t.Fatalf("%s with its own selector: %s, want %s", c.name, got, want)
				}
				if h.householdRows(t, sel.other) != before {
					t.Fatalf("%s with selector %s changed household %s", c.name, sel.household, sel.other)
				}
			}
			if n := h.count(t, "SELECT count(*) FROM invitation WHERE id = $1 AND household_id = $2 AND state = 'cancelled'",
				sel.invitation, sel.household); n != 1 {
				t.Fatalf("the invitation of %s was not cancelled through its own selector", sel.household)
			}
		}
	})

	// FR-FAM-020/AC-4, its first clause: the object of another household.
	// Its second clause, the selector of a membership that is no longer
	// active, comes with F06 (step 6).
	t.Run("an object of one household sent with the other's selector is not found", func(t *testing.T) {
		f := startTwoHouseholds(t)
		h := f.h
		before := h.fingerprint(t)
		for _, c := range selectorCases {
			if !c.onInvitation {
				continue
			}
			crossed := c.call(h, f.a, f.x, f.yInvitation)
			missing := c.call(h, f.a, f.x, uuid4())
			wantHiddenNotFound(t, crossed)
			if !rejectionOf(t, crossed).equal(rejectionOf(t, missing)) {
				t.Errorf("%s: another household's invitation %v, a missing one %v", c.name, crossed, missing)
			}
		}
		if h.fingerprint(t) != before {
			t.Fatal("a request with another household's object changed a table")
		}
	})
}
