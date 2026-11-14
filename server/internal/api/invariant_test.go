package api_test

// The checkers of the property test (14 §4.3 "Pemeriksa"): each reads the
// whole database, or every account's reads, after every step. They stay in
// test code until ops.invariant_check uses them at step 7 (plan 1.7
// decision 16; ES §1).

import (
	"errors"
	"slices"
	"strings"
	"time"

	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/encoding/prototext"
	"pgregory.net/rapid"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/household"
)

// invariantCheck runs after every step over the whole database (14 §4.3
// "Pemeriksa").
type invariantCheck func(t *rapid.T, m *householdMachine)

// adminStrings runs query as admin and returns its one text column.
func (m *householdMachine) adminStrings(t *rapid.T, query string, args ...any) []string {
	rows, err := m.h.admin.Query(t.Context(), query, args...)
	if err != nil {
		t.Fatal(err)
	}
	got, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		t.Fatal(err)
	}
	return got
}

// checkEvents is INV-A01 with its "Cara menguji": each household and
// invitation has exactly the events seq 0 to its version, so every
// transition wrote one; no event belongs to no object; and every event seen
// before is still there, unchanged (06 §2.5).
func checkEvents(t *rapid.T, m *householdMachine) {
	if bad := m.adminStrings(t, `
		SELECT format('%s %s version %s has events %s', o.kind, o.id, o.version,
		              coalesce(array_agg(e.seq ORDER BY e.seq) FILTER (WHERE e.seq IS NOT NULL), '{}'))
		  FROM (SELECT id AS household_id, 'household' AS kind, id, version FROM household
		        UNION ALL SELECT household_id, 'invitation', id, version FROM invitation) o
		  LEFT JOIN object_event e ON (e.household_id, e.object_kind, e.object_id) = (o.household_id, o.kind, o.id)
		 GROUP BY o.kind, o.id, o.version
		HAVING coalesce(array_agg(e.seq ORDER BY e.seq) FILTER (WHERE e.seq IS NOT NULL), '{}')
		       IS DISTINCT FROM ARRAY(SELECT generate_series(0, o.version))
		UNION ALL
		SELECT format('event %s of %s %s has no object', e.seq, e.object_kind, e.object_id)
		  FROM object_event e
		 WHERE NOT EXISTS (SELECT FROM household h WHERE e.object_kind = 'household' AND h.id = e.object_id)
		   AND NOT EXISTS (SELECT FROM invitation i WHERE e.object_kind = 'invitation' AND (i.household_id, i.id) = (e.household_id, e.object_id))`); len(bad) > 0 {
		t.Fatalf("INV-A01: %s", strings.Join(bad, "; "))
	}
	rows, err := m.h.admin.Query(t.Context(), `
		SELECT concat_ws(' ', object_kind, object_id, seq),
		       concat_ws(' ', household_id, actor_kind, coalesce(actor::text, '-'), coalesce(rule, '-'), command,
		                 coalesce(from_state, '-'), to_state, coalesce(idempotency_key::text, '-'), server_time)
		  FROM object_event`)
	if err != nil {
		t.Fatal(err)
	}
	events := map[string]string{}
	var key, event string
	if _, err := pgx.ForEachRow(rows, []any{&key, &event}, func() error { events[key] = event; return nil }); err != nil {
		t.Fatal(err)
	}
	for key, was := range m.seenEvents {
		if events[key] != was {
			t.Fatalf("INV-A01: event %s was %q, is %q", key, was, events[key])
		}
	}
	m.seenEvents = events
}

// checkMemberships is INV-A48 (a) to (c) as far as 1.7 builds them (06
// INV-A48 "Cara menguji"; F06 to F08 come at step 6).
func checkMemberships(t *rapid.T, m *householdMachine) {
	// (a) No membership ends or changes its role or permission set.
	rows, err := m.h.admin.Query(t.Context(), "SELECT household_id::text || ' ' || member_id, role || ' ' || permission_set FROM membership")
	if err != nil {
		t.Fatal(err)
	}
	now := map[string]string{}
	var key, value string
	if _, err := pgx.ForEachRow(rows, []any{&key, &value}, func() error { now[key] = value; return nil }); err != nil {
		t.Fatal(err)
	}
	for key, was := range m.seenMemberships {
		if now[key] != was {
			t.Fatalf("INV-A48 (a): membership %s was %q, is %q", key, was, now[key])
		}
	}
	m.seenMemberships = now

	// (b) A membership other than the creator's comes only from F04 by the
	// inviter, and no system transition leads to confirmed.
	if bad := m.adminStrings(t, `
		SELECT format('membership %s %s', m.household_id, m.member_id) FROM membership m
		 WHERE m.member_id <> (SELECT actor FROM object_event c
		                        WHERE c.household_id = m.household_id AND c.object_kind = 'household' AND c.seq = 0)
		   AND NOT EXISTS (
		       SELECT FROM invitation i JOIN object_event e
		           ON (e.household_id, e.object_kind, e.object_id) = (i.household_id, 'invitation', i.id)
		        WHERE i.household_id = m.household_id AND i.state = 'confirmed' AND i.invitee_account_id = m.account_id
		          AND e.to_state = 'confirmed' AND e.actor_kind = 'member' AND e.actor = i.inviter)
		UNION ALL
		SELECT format('system event %s of %s %s to confirmed', seq, object_kind, object_id)
		  FROM object_event WHERE actor_kind = 'system' AND to_state = 'confirmed'`); len(bad) > 0 {
		t.Fatalf("INV-A48 (b): %s", strings.Join(bad, "; "))
	}

	// (c) An account that is not a member, an invitee before F04 included,
	// reads nothing of the household: Sync and GetOpenInvitation answer as
	// for a household that does not exist (SEC-05), and ListMemberships
	// does not list it.
	households := m.adminStrings(t, "SELECT id::text FROM household")
	members := map[[2]string]bool{}
	for _, p := range m.memberships(t) {
		members[p] = true
	}
	for i, a := range m.accounts {
		listed, err := a.households.ListMemberships(bearer(a.token), &casav1.ListMembershipsRequest{RequestContext: identityRC(ios)})
		if err != nil {
			t.Fatal(err)
		}
		for _, hh := range households {
			if members[[2]string{a.id, hh}] {
				continue
			}
			if slices.ContainsFunc(listed.GetMemberships(), func(ms *casav1.Membership) bool { return ms.GetHouseholdId() == hh }) {
				t.Fatalf("INV-A48 (c): account %d lists household %s", i, m.names.of(hh))
			}
			_, err := a.syncs.Sync(bearer(a.token), syncRequest(hh, "", 0))
			wantHidden(t, err, "Sync", i, m.names.of(hh))
			_, err = a.households.GetOpenInvitation(bearer(a.token), &casav1.GetOpenInvitationRequest{RequestContext: identityRC(ios), HouseholdId: hh})
			wantHidden(t, err, "GetOpenInvitation", i, m.names.of(hh))
		}
	}
}

// wantHidden checks the one answer of a household that does not exist:
// NotFound, "not found", no detail (SEC-05).
func wantHidden(t *rapid.T, err error, read string, actor int, household string) {
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) || connectErr.Code() != connect.CodeNotFound || connectErr.Message() != "not found" ||
		len(connectErr.Details()) != 0 {
		t.Fatalf("INV-A48 (c): %s by non-member account %d of %s = %v, want the NotFound of a household that does not exist",
			read, actor, household, err)
	}
}

// checkResend: a resent accepted command gets its first answer and writes
// nothing (ADR-005).
func checkResend(t *rapid.T, m *householdMachine) {
	r := m.resent
	m.resent = nil
	if r == nil {
		return
	}
	if r.again != r.first {
		t.Fatalf("ADR-005: resend answered\n%s\nwant the first answer\n%s", r.again, r.first)
	}
	if writes := resendWrites(r.rowsBefore, r.rowsAfter); len(writes) > 0 {
		t.Fatalf("ADR-005: the resend wrote %s", strings.Join(writes, "; "))
	}
}

// resendWrites lists what changed between before and after that a resend
// may not write. The one exception is the lazy expiry: a token-free
// household command writes the due expiry of the household's live
// invitation before it reads its idempotency record (09.3 §9; plan 1.6
// decision 12). That write is a system transition any access may trigger,
// idempotent with the expiry job, not the command written again, which is
// what ADR-005 forbids (lead ruling, Task 8 fix round). So a new
// invitation.expired event by system, and the new feed rows of that
// invitation, are allowed; any other added, changed or lost row is not.
func resendWrites(before, after rowsSnapshot) []string {
	var writes []string
	if before.counts != after.counts {
		writes = append(writes, "rows from "+before.counts+" to "+after.counts)
	}
	expired := map[string]bool{}
	for key, event := range after.events {
		was, existed := before.events[key]
		switch {
		case existed && was == event:
		case !existed && event == "system invitation.expired" && strings.HasPrefix(key, "invitation "):
			expired[strings.Fields(key)[1]] = true
		default:
			writes = append(writes, "event "+key+" "+event)
		}
	}
	for key, row := range after.feed {
		was, existed := before.feed[key]
		if existed && was == row {
			continue
		}
		if f := strings.Fields(row); existed || f[0] != "invitation" || !expired[f[1]] {
			writes = append(writes, "feed row "+key+" "+row)
		}
	}
	for _, rows := range [][2]map[string]string{{before.events, after.events}, {before.feed, after.feed}} {
		for key := range rows[0] {
			if _, ok := rows[1][key]; !ok {
				writes = append(writes, "lost row "+key)
			}
		}
	}
	slices.Sort(writes)
	return writes
}

// checkSyncEqualsRead: for every member, both device models caught up from
// their cursors (the one the sync action moves, on a copy, and the one
// synced after every step, in place) equal a snapshot from an empty cursor,
// and the snapshot equals the direct reads ListMemberships and
// GetOpenInvitation in the fields those carry (SEC-06; 14 §4.7; plan 1.7
// review focus 4).
//
// The one allowed difference is X3's: an invitation whose expiry job is due
// but has not run yet stays in the model until the job writes the expiry
// and its feed rows (plan 1.7 decision 1; 10.1 §10 note 1.7 Q3). A job
// that ran and wrote no feed row is caught, because the job is gone.
func checkSyncEqualsRead(t *rapid.T, m *householdMachine) {
	due := m.adminStrings(t, "SELECT unique_key FROM job WHERE kind = $1 AND run_at <= $2",
		household.ExpireInvitationKind, m.h.clock.Now())
	for _, p := range m.memberships(t) {
		i := m.accountIndex(p[0])
		a, hh := m.accounts[i], p[1]
		lagging := &syncModel{objects: map[string]*casav1.Change{}}
		if feed := a.feeds[hh]; feed != nil {
			lagging = feed.clone()
		}
		if a.synced[hh] == nil {
			a.synced[hh] = &syncModel{objects: map[string]*casav1.Change{}}
		}
		snapshot := &syncModel{objects: map[string]*casav1.Change{}}
		m.catchUp(t, a, hh, snapshot, 0)
		for device, model := range map[string]*syncModel{"lagging": lagging, "synced": a.synced[hh]} {
			m.catchUp(t, a, hh, model, 0)
			got := changeLines(model, due)
			if want := changeLines(snapshot, nil); !slices.Equal(got, want) {
				t.Fatalf("account %d, household %s: the %s device holds\n%s\nwant the snapshot\n%s",
					i, m.names.of(hh), device, strings.Join(got, "\n"), strings.Join(want, "\n"))
			}
		}
		if got, want := readLines(snapshot), m.directRead(t, a, hh); !slices.Equal(got, want) {
			t.Fatalf("account %d, household %s: the snapshot\n%s\nwant the direct read\n%s",
				i, m.names.of(hh), strings.Join(got, "\n"), strings.Join(want, "\n"))
		}
	}
}

// changeLines are a model's objects in text form, sorted, without the
// objects of ids skip.
func changeLines(s *syncModel, skip []string) []string {
	var lines []string
	for id, c := range s.objects {
		if !slices.Contains(skip, id) {
			lines = append(lines, prototext.MarshalOptions{}.Format(c))
		}
	}
	slices.Sort(lines)
	return lines
}

// readLines are a model's objects in the fields the direct reads carry,
// sorted.
func readLines(s *syncModel) []string {
	var lines []string
	for _, c := range s.objects {
		switch {
		case c.GetHousehold() != nil:
			lines = append(lines, householdLine(c.GetHousehold().GetHouseholdId(), c.GetHousehold().GetName()))
		case c.GetMember() != nil:
			lines = append(lines, memberLine(c.GetMember()))
		case c.GetInvitation() != nil:
			i := c.GetInvitation()
			lines = append(lines, invitationLine(i.GetInvitationId(), i.GetState(), i.GetExpiresAt().AsTime()))
		}
	}
	slices.Sort(lines)
	return lines
}

// directRead is what ListMemberships and GetOpenInvitation show a of
// household, as readLines prints it.
func (m *householdMachine) directRead(t *rapid.T, a *account, household string) []string {
	listed, err := a.households.ListMemberships(bearer(a.token), &casav1.ListMembershipsRequest{RequestContext: identityRC(ios)})
	if err != nil {
		t.Fatal(err)
	}
	var lines []string
	for _, ms := range listed.GetMemberships() {
		if ms.GetHouseholdId() != household {
			continue
		}
		lines = append(lines, householdLine(ms.GetHouseholdId(), ms.GetHouseholdName()))
		for _, p := range ms.GetMembers() {
			lines = append(lines, memberLine(p))
		}
	}
	open, err := a.households.GetOpenInvitation(bearer(a.token), &casav1.GetOpenInvitationRequest{RequestContext: identityRC(ios), HouseholdId: household})
	if err != nil {
		t.Fatal(err)
	}
	if open.GetInvitationId() != "" {
		lines = append(lines, invitationLine(open.GetInvitationId(), open.GetState(), open.GetExpiresAt().AsTime()))
	}
	slices.Sort(lines)
	return lines
}

func householdLine(id, name string) string { return "household " + id + " " + name }

func memberLine(p *casav1.Member) string {
	return "member " + p.GetMemberId() + " " + p.GetDisplayName()
}

func invitationLine(id string, state casav1.InvitationState, expiresAt time.Time) string {
	return "invitation " + id + " " + state.String() + " " + expiresAt.UTC().Format(time.RFC3339Nano)
}
