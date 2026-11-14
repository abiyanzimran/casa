package read_test

import (
	"bytes"
	"errors"
	"slices"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
	"github.com/abiyanzimran/casa/server/internal/read"
)

var now = time.Date(2026, 11, 14, 9, 0, 0, 0, time.UTC)

type harness struct {
	admin, app *pgxpool.Pool
	read       *read.Service
}

func start(t *testing.T) *harness {
	t.Helper()
	d := databasetest.Migrated(t)
	app := d.Pool(t, "app_rw")
	return &harness{admin: d.Pool(t, d.AdminRole), app: app, read: read.New(app, func() time.Time { return now })}
}

func (h *harness) exec(t *testing.T, sql string, args ...any) {
	t.Helper()
	if _, err := h.admin.Exec(t.Context(), sql, args...); err != nil {
		t.Fatal(err)
	}
}

func (h *harness) uuid(t *testing.T, sql string, args ...any) pgtype.UUID {
	t.Helper()
	var id pgtype.UUID
	if err := h.admin.QueryRow(t.Context(), sql, args...).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// household writes a household as the superuser, bypassing row-level
// security, and returns its id.
func (h *harness) household(t *testing.T, name string) pgtype.UUID {
	t.Helper()
	return h.uuid(t, `INSERT INTO household (name, time_zone, created_at) VALUES ($1, 'Asia/Jakarta', $2) RETURNING id`, name, now)
}

// member writes account's membership of household, joined at joined, and
// returns its member id.
func (h *harness) member(t *testing.T, household, account pgtype.UUID, name string, joined time.Time) pgtype.UUID {
	t.Helper()
	return h.uuid(t, `
		INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
		VALUES ($1, $2, $3, 'household_owner', 'two_equal_adults', $4) RETURNING member_id`,
		household, account, name, joined)
}

func (h *harness) newID(t *testing.T) pgtype.UUID {
	t.Helper()
	return h.uuid(t, "SELECT gen_random_uuid()")
}

func TestRead(t *testing.T) {
	// 20 households of 20 members each at most [Usulan] (plan 1.6 "Batasan
	// global", SQL), unchanged by the move (plan 1.7 decision 6).
	t.Run("memberships moved to read keep their bounds", func(t *testing.T) {
		h := start(t)
		h.exec(t, "DROP INDEX membership_one_household")
		account := h.newID(t)
		first := h.household(t, "Rumah Kita")
		own := h.member(t, first, account, "Ayu", now)
		later := now.Add(time.Minute)
		h.exec(t, `INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
			SELECT $1, gen_random_uuid(), 'M' || i, 'household_owner', 'two_equal_adults', $2 FROM generate_series(1, 20) i`,
			first, later)
		h.exec(t, `WITH h AS (
				INSERT INTO household (name, time_zone, created_at)
				SELECT 'H' || i, 'Asia/Jakarta', $1 FROM generate_series(1, 20) i RETURNING id)
			INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
			SELECT id, $2, 'Ayu', 'household_owner', 'two_equal_adults', $1 FROM h`, later, account)

		got, err := h.read.Memberships(t.Context(), account)
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != 20 {
			t.Fatalf("%d memberships, want 20", len(got))
		}
		g := got[0]
		if g.HouseholdID != first || g.HouseholdName != "Rumah Kita" || g.MemberID != own || len(g.Members) != 20 {
			t.Fatalf("first membership = %v %q %v with %d members, want %v Rumah Kita %v with 20",
				g.HouseholdID, g.HouseholdName, g.MemberID, len(g.Members), first, own)
		}
		if g.Members[0] != (read.Member{MemberID: own, DisplayName: "Ayu"}) {
			t.Fatalf("first member = %v, want %v Ayu", g.Members[0], own)
		}
	})

	// Plan 1.6 decision 16; 10.1 §5.1 note 1.6 item 3 (a), (b).
	t.Run("the open invitation shows the invitee only to the inviter of an accepted invitation", func(t *testing.T) {
		h := start(t)
		household := h.household(t, "Rumah Kita")
		inviterAccount, otherAccount, invitee := h.newID(t), h.newID(t), h.newID(t)
		inviter := h.member(t, household, inviterAccount, "Ayu", now)
		other := h.member(t, household, otherAccount, "Budi", now.Add(time.Minute))
		expires := now.Add(time.Hour)
		id := h.uuid(t, `
			INSERT INTO invitation (household_id, inviter, state, invitee_account_id, invitee_display_name, created_at, expires_at)
			VALUES ($1, $2, 'accepted', $3, 'Citra', $4, $5) RETURNING id`,
			household, inviter, invitee, now, expires)

		inv, found, err := h.read.OpenInvitation(t.Context(), read.Reader{HouseholdID: household, AccountID: inviterAccount, MemberID: inviter})
		want := read.OpenInvitation{ID: id, State: "accepted", ExpiresAt: expires, Inviter: inviter, InviteeDisplayName: "Citra", InviteeAccount: invitee}
		inv.ExpiresAt = inv.ExpiresAt.UTC() // the instant, not the session's zone
		if err != nil || !found || inv != want {
			t.Fatalf("inviter reads %+v %v %v, want %+v", inv, found, err, want)
		}
		inv, found, err = h.read.OpenInvitation(t.Context(), read.Reader{HouseholdID: household, AccountID: otherAccount, MemberID: other})
		want.InviteeDisplayName, want.InviteeAccount = "", pgtype.UUID{}
		inv.ExpiresAt = inv.ExpiresAt.UTC()
		if err != nil || !found || inv != want {
			t.Fatalf("other member reads %+v %v %v, want %+v", inv, found, err, want)
		}
	})

	// Plan 1.6 decision 12: a read never writes the expiry.
	t.Run("an invitation past its expiry is not open and nothing is written", func(t *testing.T) {
		h := start(t)
		household := h.household(t, "Rumah Kita")
		account := h.newID(t)
		inviter := h.member(t, household, account, "Ayu", now)
		id := h.uuid(t, `
			INSERT INTO invitation (household_id, inviter, state, created_at, expires_at)
			VALUES ($1, $2, 'open', $3, $3) RETURNING id`, household, inviter, now)

		inv, found, err := h.read.OpenInvitation(t.Context(), read.Reader{HouseholdID: household, AccountID: account, MemberID: inviter})
		if err != nil || found || inv != (read.OpenInvitation{}) {
			t.Fatalf("OpenInvitation = %+v %v %v, want none", inv, found, err)
		}
		var state string
		var events int
		if err := h.admin.QueryRow(t.Context(), `SELECT state, (SELECT count(*) FROM object_event WHERE object_id = $1)
			FROM invitation WHERE id = $1`, id).Scan(&state, &events); err != nil {
			t.Fatal(err)
		}
		if state != "open" || events != 0 {
			t.Fatalf("invitation %s with %d events, want open with none", state, events)
		}
	})

	// SEC-151, SEC-05: the membership is checked again in the read
	// transaction; none is the answer for a household that does not exist.
	t.Run("a reader without membership in the household is not found", func(t *testing.T) {
		h := start(t)
		household := h.household(t, "Rumah Kita")
		inviter := h.member(t, household, h.newID(t), "Ayu", now)
		h.exec(t, `INSERT INTO invitation (household_id, inviter, state, created_at, expires_at)
			VALUES ($1, $2, 'open', $3, $4)`, household, inviter, now, now.Add(time.Hour))
		elsewhere := h.household(t, "Rumah Lain")
		outsider := h.newID(t)
		outsiderMember := h.member(t, elsewhere, outsider, "Dewi", now)

		for name, r := range map[string]read.Reader{
			"a member of another household":   {HouseholdID: household, AccountID: outsider, MemberID: outsiderMember},
			"a household that does not exist": {HouseholdID: h.newID(t), AccountID: outsider, MemberID: outsiderMember},
		} {
			inv, found, err := h.read.OpenInvitation(t.Context(), r)
			if !errors.Is(err, read.ErrNotFound) || found || inv != (read.OpenInvitation{}) {
				t.Fatalf("%s: OpenInvitation = %+v %v %v, want ErrNotFound", name, inv, found, err)
			}
		}
	})

	// ES §3: the member query is bounded, and a household past the bound
	// fails the fan-out instead of leaving a reader out of the feed.
	t.Run("readers are every member, by id, and never a list cut at the bound", func(t *testing.T) {
		h := start(t)
		household := h.household(t, "Rumah Kita")
		readers := func() ([]pgtype.UUID, error) {
			var got []pgtype.UUID
			err := database.Write(t.Context(), h.app, database.Scope{HouseholdID: household}, func(tx pgx.Tx) (err error) {
				got, err = read.Readers(t.Context(), tx, household, read.KindInvitation)
				return err
			})
			return got, err
		}
		var want []pgtype.UUID
		for range 20 {
			want = append(want, h.member(t, household, h.newID(t), "M", now))
		}
		slices.SortFunc(want, func(a, b pgtype.UUID) int { return bytes.Compare(a.Bytes[:], b.Bytes[:]) })
		if got, err := readers(); err != nil || !slices.Equal(got, want) {
			t.Fatalf("Readers of 20 members = %v, %v; want every member by id", got, err)
		}
		h.member(t, household, h.newID(t), "M", now)
		if got, err := readers(); err == nil {
			t.Fatalf("Readers of 21 members = %v, want an error", got)
		}
	})
}
