package sync_test

import (
	"errors"
	"reflect"
	"slices"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/read"
	"github.com/abiyanzimran/casa/server/internal/sync"
)

func reader(m household.Member) read.Reader {
	return read.Reader{HouseholdID: m.HouseholdID, AccountID: m.AccountID, MemberID: m.MemberID}
}

func (x *harness) pull(t *testing.T, r read.Reader, cursor string, limit int) sync.Page {
	t.Helper()
	p, err := x.sync.Pull(t.Context(), r, cursor, limit)
	x.must(t, err)
	return p
}

// decode is the cursor s of r, which must be valid.
func decode(t *testing.T, r read.Reader, s string) sync.Cursor {
	t.Helper()
	c, err := testKey(t).Decode(s, r.HouseholdID, r.MemberID)
	if err != nil {
		t.Fatalf("decode %q: %v", s, err)
	}
	return c
}

// feedCursor is a feed cursor of r at pos, as the server makes one.
func feedCursor(t *testing.T, r read.Reader, pos int64) string {
	t.Helper()
	return testKey(t).Encode(sync.Cursor{Household: r.HouseholdID, Member: r.MemberID, Pos: pos})
}

// cache is a device's cache of r: the live objects it holds (09.3 §3).
type cache map[read.ObjectRef]read.Object

func (c cache) apply(changes []sync.Change) {
	for _, ch := range changes {
		switch ch.Kind {
		case sync.Upsert:
			c[ch.Ref] = ch.Object
		case sync.Purge:
			delete(c, ch.Ref)
		}
	}
}

// catchUp pulls from cursor until has_more is false and applies every page;
// a cursor that stops moving fails the test instead of looping.
func (x *harness) catchUp(t *testing.T, r read.Reader, c cache, cursor string, limit int) string {
	t.Helper()
	for range 50 {
		p := x.pull(t, r, cursor, limit)
		if p.ResyncRequired {
			t.Fatal("unexpected resync_required")
		}
		c.apply(p.Changes)
		if !p.HasMore {
			return p.Next
		}
		cursor = p.Next
	}
	t.Fatal("still has_more after 50 pages")
	return ""
}

// objects projects refs for r through read.Objects in a read transaction.
func (x *harness) objects(t *testing.T, r read.Reader, refs ...read.ObjectRef) map[read.ObjectRef]read.Object {
	t.Helper()
	var got map[read.ObjectRef]read.Object
	x.must(t, database.Read(t.Context(), x.app, database.Scope{HouseholdID: r.HouseholdID, AccountID: r.AccountID}, func(tx pgx.Tx) (err error) {
		got, err = x.read.Objects(t.Context(), tx, r, refs)
		return err
	}))
	return got
}

func refs(changes []sync.Change) []read.ObjectRef {
	var out []read.ObjectRef
	for _, c := range changes {
		out = append(out, c.Ref)
	}
	return out
}

func kinds(changes []sync.Change) []sync.ChangeKind {
	var out []sync.ChangeKind
	for _, c := range changes {
		out = append(out, c.Kind)
	}
	return out
}

func wantResync(t *testing.T, name string, p sync.Page) {
	t.Helper()
	if !reflect.DeepEqual(p, sync.Page{ResyncRequired: true}) {
		t.Errorf("%s: %+v, want only resync_required (no change, no cursor)", name, p)
	}
}

// deleteFeed removes r's feed rows at the positions as the superuser: the
// pruning job is step 7 (plan 1.7 decision 5), the detection is now
// (decision 11).
func (x *harness) deleteFeed(t *testing.T, r read.Reader, positions ...int64) {
	t.Helper()
	_, err := x.admin.Exec(t.Context(), "DELETE FROM reader_feed WHERE household_id = $1 AND member_id = $2 AND pos = ANY($3)",
		r.HouseholdID, r.MemberID, positions)
	x.must(t, err)
}

// cancel is F03 by m.
func (x *harness) cancel(t *testing.T, m household.Member, inv pgtype.UUID) {
	t.Helper()
	x.must(t, x.write(t, m.HouseholdID, m.AccountID, func(tx pgx.Tx) error {
		return x.h.CancelInvitation(t.Context(), tx, m, inv, newID())
	}))
}

// fiveRows makes a's feed 1 household, 2 a, 3 and 4 a first invitation
// (F01, F03), 5 a second invitation (F01).
func (x *harness) fiveRows(t *testing.T) (household.Member, household.Invitation, household.Invitation) {
	t.Helper()
	a := x.create(t)
	first := x.invite(t, a)
	x.cancel(t, a, first.ID)
	second := x.invite(t, a)
	if n := x.lastPos(t, a.HouseholdID, a.MemberID); n != 5 {
		t.Fatalf("last_pos %d, want 5", n)
	}
	return a, first, second
}

// TestPull proves Sync's read of 09.3 §3 and ADR-017 items 3, 6 and 7: an
// initial snapshot through the read path, paged across calls with the feed
// position of its first page, then the feed, each change projected by the
// same projections as the screens (SEC-06), and resync_required for a cursor
// outside the feed that exists (plan 1.7 decisions 9 to 11).
func TestPull(t *testing.T) {
	t.Run("an empty cursor returns the household, its members and its live invitation, then the feed position", func(t *testing.T) {
		x := newHarness(t)
		a := x.create(t)
		inv := x.invite(t, a)
		b := x.addMember(t, a.HouseholdID)
		r := reader(a)
		p := x.pull(t, r, "", 100)
		members := byID(a.MemberID, b)
		wantRefs := []read.ObjectRef{
			{Kind: read.KindHousehold, ID: a.HouseholdID},
			{Kind: read.KindMember, ID: members[0]},
			{Kind: read.KindMember, ID: members[1]},
			{Kind: read.KindInvitation, ID: inv.ID},
		}
		if !slices.Equal(refs(p.Changes), wantRefs) || !slices.Equal(kinds(p.Changes), []sync.ChangeKind{sync.Upsert, sync.Upsert, sync.Upsert, sync.Upsert}) {
			t.Fatalf("changes %+v, want upserts of %+v", p.Changes, wantRefs)
		}
		if got := *p.Changes[0].Object.Household; got != (read.Household{ID: a.HouseholdID, Name: "Rumah", TimeZone: "Asia/Jakarta", Currency: "IDR"}) {
			t.Errorf("household %+v", got)
		}
		names := map[pgtype.UUID]string{a.MemberID: "Ayu", b: "Budi"}
		for _, c := range p.Changes[1:3] {
			if got := *c.Object.Member; got != (read.Member{MemberID: c.Ref.ID, DisplayName: names[c.Ref.ID]}) {
				t.Errorf("member %+v", got)
			}
		}
		got := *p.Changes[3].Object.Invitation
		if got.ID != inv.ID || got.State != "open" || !got.ExpiresAt.Equal(inv.ExpiresAt) || got.Inviter != a.MemberID {
			t.Errorf("invitation %+v, want %v open until %v by %v", got, inv.ID, inv.ExpiresAt, a.MemberID)
		}
		if p.HasMore || p.ResyncRequired {
			t.Errorf("has_more %v resync %v, want both false", p.HasMore, p.ResyncRequired)
		}
		if c := decode(t, r, p.Next); c.Snapshot != nil || c.Pos != 3 {
			t.Errorf("next cursor %+v, want the feed at last_pos 3", c)
		}
		again := x.pull(t, r, p.Next, 100)
		if len(again.Changes) != 0 || again.HasMore || decode(t, r, again.Next).Pos != 3 {
			t.Errorf("the feed after the snapshot: %+v, want nothing more at 3", again)
		}
	})

	t.Run("the snapshot pages by kind and id with the limit", func(t *testing.T) {
		x := newHarness(t)
		a := x.create(t)
		x.invite(t, a)
		x.addMember(t, a.HouseholdID)
		r := reader(a)
		whole := x.pull(t, r, "", 100)
		var paged []sync.Change
		cursor := ""
		for page := 1; ; page++ {
			p := x.pull(t, r, cursor, 1)
			if len(p.Changes) != 1 {
				t.Fatalf("page %d has %d changes, want 1", page, len(p.Changes))
			}
			paged = append(paged, p.Changes...)
			c := decode(t, r, p.Next)
			if page < 4 && (!p.HasMore || c.Snapshot == nil || c.Pos != 3) {
				t.Fatalf("page %d: has_more %v, cursor %+v; want more of the snapshot at P0 3", page, p.HasMore, c)
			}
			if page == 4 {
				if p.HasMore || c.Snapshot != nil || c.Pos != 3 {
					t.Fatalf("last page: has_more %v, cursor %+v; want the feed at 3", p.HasMore, c)
				}
				break
			}
			cursor = p.Next
		}
		if !reflect.DeepEqual(paged, whole.Changes) {
			t.Errorf("paged %+v\nwhole %+v", paged, whole.Changes)
		}
	})

	t.Run("a change between snapshot pages arrives through the feed", func(t *testing.T) {
		// Focus item 2 of plan 1.7; 09.3 §3 note of 14 November 2026: every
		// change after P0 has a feed row, so an object born or closed
		// between two pages reaches the device.
		x := newHarness(t)
		a := x.create(t)
		r := reader(a)
		dev := cache{}
		first := x.pull(t, r, "", 1)
		dev.apply(first.Changes)
		if !first.HasMore || decode(t, r, first.Next).Pos != 2 {
			t.Fatalf("first page %+v, want more at P0 2", first)
		}
		inv := x.invite(t, a)
		// And a change to the object the first page already returned. No
		// command renames a household yet, so the test renames it as the
		// superuser and writes the feed through sync.Changed, the one writer.
		_, err := x.admin.Exec(t.Context(), "UPDATE household SET name = 'Rumah Baru' WHERE id = $1", a.HouseholdID)
		x.must(t, err)
		x.must(t, x.write(t, a.HouseholdID, a.AccountID, func(tx pgx.Tx) error {
			return sync.Changed(t.Context(), tx, a.HouseholdID, read.KindHousehold, a.HouseholdID)
		}))
		next := x.catchUp(t, r, dev, first.Next, 1)
		home := read.ObjectRef{Kind: read.KindHousehold, ID: a.HouseholdID}
		if h := dev[home].Household; h == nil || h.Name != "Rumah Baru" {
			t.Errorf("device holds household %+v, want the renamed one", dev[home])
		}
		ref := read.ObjectRef{Kind: read.KindInvitation, ID: inv.ID}
		direct := x.objects(t, r, ref)
		if !reflect.DeepEqual(dev[ref], direct[ref]) || dev[ref].Invitation == nil {
			t.Errorf("device holds %+v, read.Objects projects %+v", dev[ref], direct[ref])
		}
		if c := decode(t, r, next); c.Snapshot != nil || c.Pos != 5 {
			t.Errorf("cursor after catching up %+v, want the feed at 5", c)
		}
		x.cancel(t, a, inv.ID)
		x.catchUp(t, r, dev, next, 1)
		if _, held := dev[ref]; held {
			t.Error("the device still holds the cancelled invitation")
		}
		fresh := cache{}
		x.catchUp(t, r, fresh, "", 100)
		if !reflect.DeepEqual(dev, fresh) {
			t.Errorf("device %+v, a fresh snapshot %+v", dev, fresh)
		}
	})

	t.Run("a feed cursor returns each change once in order and the next cursor", func(t *testing.T) {
		x := newHarness(t)
		a, first, second := x.fiveRows(t)
		r := reader(a)
		p := x.pull(t, r, feedCursor(t, r, 2), 2)
		firstRef := read.ObjectRef{Kind: read.KindInvitation, ID: first.ID}
		if !slices.Equal(refs(p.Changes), []read.ObjectRef{firstRef, firstRef}) || !p.HasMore || decode(t, r, p.Next).Pos != 4 {
			t.Fatalf("page 1 %+v, want positions 3 and 4 and more", p)
		}
		p = x.pull(t, r, p.Next, 2)
		secondRef := read.ObjectRef{Kind: read.KindInvitation, ID: second.ID}
		if !slices.Equal(refs(p.Changes), []read.ObjectRef{secondRef}) || p.HasMore || decode(t, r, p.Next).Pos != 5 {
			t.Fatalf("page 2 %+v, want position 5 and no more", p)
		}
		if p.Changes[0].Kind != sync.Upsert || !reflect.DeepEqual(p.Changes[0].Object, x.objects(t, r, secondRef)[secondRef]) {
			t.Errorf("the live invitation %+v, want its projection", p.Changes[0])
		}
		p = x.pull(t, r, p.Next, 2)
		if len(p.Changes) != 0 || p.HasMore || decode(t, r, p.Next).Pos != 5 {
			t.Errorf("caught up %+v, want nothing at 5", p)
		}
	})

	t.Run("an invitation no longer live is a purge", func(t *testing.T) {
		// Plan 1.7 decision 7: the device cache holds live objects only.
		x := newHarness(t)
		a := x.create(t)
		inv := x.invite(t, a)
		x.cancel(t, a, inv.ID)
		r := reader(a)
		p := x.pull(t, r, feedCursor(t, r, 2), 100)
		if !slices.Equal(kinds(p.Changes), []sync.ChangeKind{sync.Purge, sync.Purge}) || p.Changes[0].Object != (read.Object{}) {
			t.Errorf("changes %+v, want two bare purges", p.Changes)
		}
		if snap := x.pull(t, r, "", 100); slices.Contains(refs(snap.Changes), read.ObjectRef{Kind: read.KindInvitation, ID: inv.ID}) {
			t.Error("the snapshot holds the cancelled invitation")
		}
	})

	t.Run("a purge row is a purge even for an object still readable", func(t *testing.T) {
		// Plan 1.7 Task 5: a row with change = 'purge' is always a purge.
		// The writer makes none before F06 (step 6), so the test writes one
		// as the superuser.
		x := newHarness(t)
		a := x.create(t)
		r := reader(a)
		_, err := x.admin.Exec(t.Context(), `
			WITH c AS (UPDATE reader_feed_counter SET last_pos = 3 WHERE household_id = $1 AND member_id = $2)
			INSERT INTO reader_feed (household_id, member_id, pos, object_kind, object_id, change)
			VALUES ($1, $2, 3, 'household', $1, 'purge')`, a.HouseholdID, a.MemberID)
		x.must(t, err)
		p := x.pull(t, r, feedCursor(t, r, 2), 100)
		if !reflect.DeepEqual(p.Changes, []sync.Change{{Kind: sync.Purge, Ref: read.ObjectRef{Kind: read.KindHousehold, ID: a.HouseholdID}}}) {
			t.Errorf("changes %+v, want a purge of the household", p.Changes)
		}
	})

	t.Run("an invitation past its expiry but not yet written is a purge", func(t *testing.T) {
		// Plan 1.7 decision 7: live is open or accepted and now < expires_at,
		// the rule of GetOpenInvitation, whether the expiry is written or not.
		x := newHarness(t)
		a := x.create(t)
		inv := x.invite(t, a)
		x.now = inv.ExpiresAt
		r := reader(a)
		p := x.pull(t, r, feedCursor(t, r, 2), 100)
		if !slices.Equal(kinds(p.Changes), []sync.ChangeKind{sync.Purge}) {
			t.Errorf("changes %+v, want one purge", p.Changes)
		}
		if snap := x.pull(t, r, "", 100); slices.Contains(refs(snap.Changes), read.ObjectRef{Kind: read.KindInvitation, ID: inv.ID}) {
			t.Error("the snapshot holds the lapsed invitation")
		}
		var state string
		x.must(t, x.admin.QueryRow(t.Context(), "SELECT state FROM invitation WHERE id = $1", inv.ID).Scan(&state))
		if state != "open" {
			t.Fatalf("state %q, want the expiry not yet written", state)
		}
	})

	t.Run("a cursor behind the pruned feed asks for a resync", func(t *testing.T) {
		// Focus item 5; 09.3 §3 "Pemangkasan".
		x := newHarness(t)
		a, _, _ := x.fiveRows(t)
		r := reader(a)
		x.deleteFeed(t, r, 1, 2)
		wantResync(t, "cursor 0", x.pull(t, r, feedCursor(t, r, 0), 100))
		wantResync(t, "cursor 1", x.pull(t, r, feedCursor(t, r, 1), 100))
		if p := x.pull(t, r, feedCursor(t, r, 2), 100); p.ResyncRequired || len(p.Changes) != 3 {
			t.Errorf("a cursor at the pruned edge: %+v, want the three rows left", p)
		}
	})

	t.Run("a cursor ahead of the feed asks for a resync", func(t *testing.T) {
		// A cursor from a restored database: P > last_pos.
		x := newHarness(t)
		a, _, _ := x.fiveRows(t)
		r := reader(a)
		wantResync(t, "last_pos + 1", x.pull(t, r, feedCursor(t, r, 6), 100))
		other := x.create(t)
		ro := reader(other)
		wantResync(t, "ahead of a short feed", x.pull(t, ro, feedCursor(t, ro, 3), 100))
		// A counter restored behind rows that remain: only P > last_pos
		// tells it, since rows 4 and 5 follow cursor 3 on a full page.
		_, err := x.admin.Exec(t.Context(), "UPDATE reader_feed_counter SET last_pos = 2 WHERE household_id = $1 AND member_id = $2",
			r.HouseholdID, r.MemberID)
		x.must(t, err)
		wantResync(t, "a counter behind its rows", x.pull(t, r, feedCursor(t, r, 3), 1))
	})

	t.Run("a gap inside the returned rows asks for a resync", func(t *testing.T) {
		x := newHarness(t)
		a, _, _ := x.fiveRows(t)
		r := reader(a)
		x.deleteFeed(t, r, 4)
		wantResync(t, "rows 3 and 5", x.pull(t, r, feedCursor(t, r, 2), 100))
		wantResync(t, "rows 3 and 5 on a page of one", x.pull(t, r, feedCursor(t, r, 2), 1))
	})

	t.Run("a page that is not full and does not end at last_pos asks for a resync", func(t *testing.T) {
		x := newHarness(t)
		a, _, _ := x.fiveRows(t)
		r := reader(a)
		x.deleteFeed(t, r, 4, 5)
		wantResync(t, "row 3 of 5", x.pull(t, r, feedCursor(t, r, 2), 100))
		wantResync(t, "nothing after 3 of 5", x.pull(t, r, feedCursor(t, r, 3), 100))
	})

	t.Run("projections carry no invitee name or email", func(t *testing.T) {
		// Plan 1.6 Q12: sync carries the invitation's state only; the
		// invitee's name and email stay online (GetOpenInvitation).
		x := newHarness(t)
		a := x.create(t)
		inv := x.invite(t, a)
		x.accept(t, a, inv.ID)
		p := x.pull(t, reader(a), "", 100)
		last := p.Changes[len(p.Changes)-1]
		if last.Ref.ID != inv.ID || last.Object.Invitation.State != "accepted" || last.Object.Invitation.Inviter != a.MemberID {
			t.Fatalf("last change %+v, want the accepted invitation", last)
		}
		var fields []string
		for f := range reflect.TypeFor[read.Invitation]().Fields() {
			fields = append(fields, f.Name)
		}
		if !slices.Equal(fields, []string{"ID", "State", "ExpiresAt", "Inviter"}) {
			t.Errorf("read.Invitation fields %v, want only ID, State, ExpiresAt, Inviter", fields)
		}
	})

	t.Run("a reader no longer a member is not found and a cursor of another member is rejected", func(t *testing.T) {
		x := newHarness(t)
		a := x.create(t)
		r := reader(a)
		for name, stranger := range map[string]read.Reader{
			"another account": {HouseholdID: a.HouseholdID, AccountID: newID(), MemberID: a.MemberID},
			"another member":  {HouseholdID: a.HouseholdID, AccountID: a.AccountID, MemberID: newID()},
		} {
			if _, err := x.sync.Pull(t.Context(), stranger, "", 100); !errors.Is(err, read.ErrNotFound) {
				t.Errorf("%s: err %v, want read.ErrNotFound", name, err)
			}
		}
		b := x.addMember(t, a.HouseholdID)
		theirs := feedCursor(t, read.Reader{HouseholdID: a.HouseholdID, MemberID: b}, 0)
		if _, err := x.sync.Pull(t.Context(), r, theirs, 100); !errors.Is(err, sync.ErrCursorRejected) {
			t.Errorf("another member's cursor: err %v, want ErrCursorRejected", err)
		}
	})
}
