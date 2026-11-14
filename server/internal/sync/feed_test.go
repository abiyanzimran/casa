package sync_test

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"slices"
	gosync "sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/jobs"
	"github.com/abiyanzimran/casa/server/internal/read"
	"github.com/abiyanzimran/casa/server/internal/sync"
)

var start = time.Date(2026, 11, 14, 9, 0, 0, 0, time.UTC)

// harness runs household's commands as app_rw against a migrated database
// and reads the feed as the superuser.
type harness struct {
	admin, app *pgxpool.Pool
	h          *household.Service
	now        time.Time
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	d := databasetest.Migrated(t)
	x := &harness{admin: d.Pool(t, d.AdminRole), app: d.Pool(t, "app_rw"), now: start}
	x.h = household.New(x.app, func() time.Time { return x.now }, "Pengelola", household.PublishedNotices)
	return x
}

// newID is a version 4 UUID; crypto/rand.Read never fails.
func newID() pgtype.UUID {
	var id pgtype.UUID
	_, _ = rand.Read(id.Bytes[:])
	id.Bytes[6], id.Bytes[8] = id.Bytes[6]&0x0f|0x40, id.Bytes[8]&0x3f|0x80
	id.Valid = true
	return id
}

// newTokenHash is the SHA-256 of a fresh link token.
func newTokenHash() []byte {
	token := newID()
	hash := sha256.Sum256(token.Bytes[:])
	return hash[:]
}

// write runs fn as one command in household's scope with account.
func (x *harness) write(t *testing.T, household, account pgtype.UUID, fn func(pgx.Tx) error) error {
	t.Helper()
	return database.Write(t.Context(), x.app, database.Scope{HouseholdID: household, AccountID: account}, fn)
}

func (x *harness) must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// create makes a household through Create and returns its creator.
func (x *harness) create(t *testing.T) household.Member {
	t.Helper()
	id, err := x.h.NewHouseholdID(t.Context())
	x.must(t, err)
	account := newID()
	var m household.Member
	x.must(t, x.write(t, id, account, func(tx pgx.Tx) (err error) {
		m, err = x.h.Create(t.Context(), tx, account, id, household.NewHousehold{Name: "Rumah", TimeZone: "Asia/Jakarta", DisplayName: "Ayu"}, 0, newID())
		return err
	}))
	return m
}

// addMember writes a second member as the superuser, as test data; it has
// no feed row until a change it can read.
func (x *harness) addMember(t *testing.T, household pgtype.UUID) pgtype.UUID {
	t.Helper()
	var id pgtype.UUID
	x.must(t, x.admin.QueryRow(t.Context(), `
		INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
		VALUES ($1, gen_random_uuid(), 'Budi', 'household_owner', 'two_equal_adults', $2) RETURNING member_id`,
		household, x.now).Scan(&id))
	return id
}

// invite is F01 by m and returns the invitation's id.
func (x *harness) invite(t *testing.T, m household.Member) household.Invitation {
	t.Helper()
	var inv household.Invitation
	x.must(t, x.write(t, m.HouseholdID, m.AccountID, func(tx pgx.Tx) (err error) {
		inv, err = x.h.CreateInvitation(t.Context(), tx, m, newID())
		return err
	}))
	return inv
}

// accept issues a link for the invitation and accepts it as a new account,
// which it returns.
func (x *harness) accept(t *testing.T, m household.Member, inv pgtype.UUID) pgtype.UUID {
	t.Helper()
	token := newTokenHash()
	_, err := x.h.IssueInvitationLink(t.Context(), m, inv, token)
	x.must(t, err)
	invitee := newID()
	x.must(t, x.write(t, m.HouseholdID, invitee, func(tx pgx.Tx) error {
		return x.h.Accept(t.Context(), tx, invitee, inv, token, "Budi", 0, newID())
	}))
	return invitee
}

type feedRow struct {
	Pos    int64
	Kind   string
	ID     pgtype.UUID
	Change string
}

// feed is member's feed in household, by position.
func (x *harness) feed(t *testing.T, household, member pgtype.UUID) []feedRow {
	t.Helper()
	rows, err := x.admin.Query(t.Context(), `
		SELECT pos, object_kind, object_id, change FROM reader_feed
		 WHERE household_id = $1 AND member_id = $2 ORDER BY pos`, household, member)
	x.must(t, err)
	got, err := pgx.CollectRows(rows, pgx.RowToStructByPos[feedRow])
	x.must(t, err)
	return got
}

// feedCount is the number of feed and counter rows of every household.
func (x *harness) feedCount(t *testing.T) int {
	t.Helper()
	var n int
	x.must(t, x.admin.QueryRow(t.Context(),
		"SELECT (SELECT count(*) FROM reader_feed) + (SELECT count(*) FROM reader_feed_counter)").Scan(&n))
	return n
}

// lastPos is member's counter in household.
func (x *harness) lastPos(t *testing.T, household, member pgtype.UUID) int64 {
	t.Helper()
	var n int64
	x.must(t, x.admin.QueryRow(t.Context(),
		"SELECT last_pos FROM reader_feed_counter WHERE household_id = $1 AND member_id = $2", household, member).Scan(&n))
	return n
}

// since are the rows after position before.
func since(rows []feedRow, before int64) []feedRow {
	return slices.DeleteFunc(rows, func(r feedRow) bool { return r.Pos <= before })
}

func upsert(pos int64, kind read.ObjectKind, id pgtype.UUID) feedRow {
	return feedRow{Pos: pos, Kind: string(kind), ID: id, Change: "upsert"}
}

// byID orders member ids as PostgreSQL orders uuid: byte by byte.
func byID(ids ...pgtype.UUID) []pgtype.UUID {
	slices.SortFunc(ids, func(a, b pgtype.UUID) int { return bytes.Compare(a.Bytes[:], b.Bytes[:]) })
	return ids
}

func wantFeed(t *testing.T, name string, got, want []feedRow) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("%s:\n got  %+v\n want %+v", name, got, want)
	}
}

// TestReaderFeed proves the fan-out of ADR-017 item 3 and 09.2 §6 step 6:
// every change of a household object writes, in the command's own
// transaction, one row per member who can read it, at gapless positions per
// reader (09.3 §3; plan 1.7 decision 8).
func TestReaderFeed(t *testing.T) {
	t.Run("creating a household writes the household and the creator to the creator's feed", func(t *testing.T) {
		x := newHarness(t)
		a := x.create(t)
		wantFeed(t, "creator", x.feed(t, a.HouseholdID, a.MemberID), []feedRow{
			upsert(1, read.KindHousehold, a.HouseholdID),
			upsert(2, read.KindMember, a.MemberID),
		})
		if n := x.lastPos(t, a.HouseholdID, a.MemberID); n != 2 {
			t.Errorf("last_pos %d, want 2", n)
		}
	})

	// Each case brings a fresh household's invitation to the state its
	// transition needs, adds a second member where the transition allows
	// one (F01 and F02 keep a household below two full adults, 06 §3.6), and
	// runs the transition: each member's feed grows by exactly one
	// invitation row (focus item 4 of plan 1.7).
	t.Run("every invitation transition writes one row per member", func(t *testing.T) {
		// asInviter runs a command of the inviter a.
		asInviter := func(cmd func(t *testing.T, x *harness, a household.Member, inv pgtype.UUID, tx pgx.Tx) error) func(*testing.T, *harness, household.Member, pgtype.UUID, pgtype.UUID) error {
			return func(t *testing.T, x *harness, a household.Member, inv, _ pgtype.UUID) error {
				return x.write(t, a.HouseholdID, a.AccountID, func(tx pgx.Tx) error { return cmd(t, x, a, inv, tx) })
			}
		}
		for _, c := range []struct {
			name string
			// accepted: the transition starts from an accepted invitation;
			// secondMember: a second member joins before it.
			accepted, secondMember bool
			// run is the transition; nil for F01, the invitation's creation.
			run func(t *testing.T, x *harness, a household.Member, inv, invitee pgtype.UUID) error
		}{
			{name: "F01"},
			{name: "F02 accept", run: func(t *testing.T, x *harness, a household.Member, inv, _ pgtype.UUID) error {
				x.accept(t, a, inv)
				return nil
			}},
			{name: "F05 withdraw", accepted: true, secondMember: true,
				run: func(t *testing.T, x *harness, a household.Member, inv, invitee pgtype.UUID) error {
					return x.write(t, a.HouseholdID, invitee, func(tx pgx.Tx) error {
						return x.h.Withdraw(t.Context(), tx, invitee, inv, newID())
					})
				}},
			{name: "CloseAsWithdrawn", accepted: true, secondMember: true,
				run: asInviter(func(t *testing.T, x *harness, _ household.Member, inv pgtype.UUID, tx pgx.Tx) error {
					return x.h.CloseAsWithdrawn(t.Context(), tx, inv)
				})},
			{name: "F03 cancel", secondMember: true,
				run: asInviter(func(t *testing.T, x *harness, a household.Member, inv pgtype.UUID, tx pgx.Tx) error {
					return x.h.CancelInvitation(t.Context(), tx, a, inv, newID())
				})},
			{name: "F05 reject", accepted: true, secondMember: true,
				run: asInviter(func(t *testing.T, x *harness, a household.Member, inv pgtype.UUID, tx pgx.Tx) error {
					return x.h.Reject(t.Context(), tx, a, inv, newID())
				})},
			{name: "expiry job", secondMember: true,
				run: func(t *testing.T, x *harness, a household.Member, inv, _ pgtype.UUID) error {
					x.now = x.now.Add(8 * 24 * time.Hour)
					return x.h.ExpireInvitationJob(t.Context(), jobs.Job{HouseholdID: a.HouseholdID, Payload: inv.Bytes[:]})
				}},
		} {
			t.Run(c.name, func(t *testing.T) {
				x := newHarness(t)
				a := x.create(t)
				members := []pgtype.UUID{a.MemberID}
				before := map[pgtype.UUID]int64{a.MemberID: 2} // the household and a
				inv := x.invite(t, a)
				if c.run != nil {
					var invitee pgtype.UUID
					if c.accepted {
						invitee = x.accept(t, a, inv.ID)
					}
					if c.secondMember {
						members = append(members, x.addMember(t, a.HouseholdID))
					}
					for _, m := range members {
						before[m] = int64(len(x.feed(t, a.HouseholdID, m)))
					}
					x.must(t, c.run(t, x, a, inv.ID, invitee))
				}
				for _, m := range members {
					wantFeed(t, "new rows of member "+m.String(), since(x.feed(t, a.HouseholdID, m), before[m]), []feedRow{
						upsert(before[m]+1, read.KindInvitation, inv.ID),
					})
				}
			})
		}
	})

	t.Run("F04 writes the household, every member and the invitation to the members' feeds, the new member included", func(t *testing.T) {
		// Ruling P4: Confirm writes the invitation's event before the
		// membership, so the invitation reaches the members of that moment
		// (the inviter); the new member receives the household and every
		// member, all at once.
		x := newHarness(t)
		a := x.create(t)
		inv := x.invite(t, a)
		x.accept(t, a, inv.ID)
		before := int64(len(x.feed(t, a.HouseholdID, a.MemberID)))
		var b household.Member
		x.must(t, x.write(t, a.HouseholdID, a.AccountID, func(tx pgx.Tx) (err error) {
			b, err = x.h.Confirm(t.Context(), tx, a, inv.ID, newID())
			return err
		}))
		members := byID(a.MemberID, b.MemberID)
		wantFeed(t, "inviter", since(x.feed(t, a.HouseholdID, a.MemberID), before), []feedRow{
			upsert(before+1, read.KindInvitation, inv.ID),
			upsert(before+2, read.KindHousehold, a.HouseholdID),
			upsert(before+3, read.KindMember, members[0]),
			upsert(before+4, read.KindMember, members[1]),
		})
		wantFeed(t, "new member", x.feed(t, a.HouseholdID, b.MemberID), []feedRow{
			upsert(1, read.KindHousehold, a.HouseholdID),
			upsert(2, read.KindMember, members[0]),
			upsert(3, read.KindMember, members[1]),
		})
	})

	t.Run("issuing a link writes no feed row", func(t *testing.T) {
		// Plan 1.7 decision 8: no event, no projected change.
		x := newHarness(t)
		a := x.create(t)
		inv := x.invite(t, a)
		before := x.feedCount(t)
		_, err := x.h.IssueInvitationLink(t.Context(), a, inv.ID, newTokenHash())
		x.must(t, err)
		if n := x.feedCount(t); n != before {
			t.Errorf("feed and counter rows %d after issuing a link, want %d", n, before)
		}
	})

	t.Run("a refused command writes no feed row", func(t *testing.T) {
		x := newHarness(t)
		a := x.create(t)
		inv := x.invite(t, a)
		x.must(t, x.write(t, a.HouseholdID, a.AccountID, func(tx pgx.Tx) error {
			return x.h.CancelInvitation(t.Context(), tx, a, inv.ID, newID())
		}))
		full := x.create(t)
		x.addMember(t, full.HouseholdID)
		// F04 refused after its fan-out: Confirm writes the invitation's
		// event, and so its feed rows, before it finds the invitee a member
		// elsewhere (ErrInviteeUnavailable; 09.2 §5.1 note on decision 20),
		// so only the command's own transaction keeps those rows out.
		late := x.create(t)
		lateInv := x.invite(t, late)
		invitee := x.accept(t, late, lateInv.ID)
		elsewhere := x.create(t)
		_, err := x.admin.Exec(t.Context(), `
			INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
			VALUES ($1, $2, 'Budi', 'household_owner', 'two_equal_adults', $3)`, elsewhere.HouseholdID, invitee, x.now)
		x.must(t, err)
		before := x.feedCount(t)
		for _, c := range []struct {
			name string
			m    household.Member
			run  func(pgx.Tx) error
			want error
		}{
			{"HOUSEHOLD_FULL", full, func(tx pgx.Tx) error {
				_, err := x.h.CreateInvitation(t.Context(), tx, full, newID())
				return err
			}, household.ErrHouseholdFull},
			{"STATE_CHANGED", a, func(tx pgx.Tx) error {
				return x.h.CancelInvitation(t.Context(), tx, a, inv.ID, newID())
			}, household.ErrStateChanged},
			{"F04 invitee a member elsewhere", late, func(tx pgx.Tx) error {
				_, err := x.h.Confirm(t.Context(), tx, late, lateInv.ID, newID())
				return err
			}, household.ErrInviteeUnavailable},
		} {
			if err := x.write(t, c.m.HouseholdID, c.m.AccountID, c.run); !errors.Is(err, c.want) {
				t.Fatalf("%s: err = %v, want %v", c.name, err, c.want)
			}
			if n := x.feedCount(t); n != before {
				t.Errorf("%s: feed and counter rows %d, want %d", c.name, n, before)
			}
		}
	})

	t.Run("positions have no gap per reader under concurrent commands", func(t *testing.T) {
		// Plan 1.7 decision 8: two commands that reach the same reader
		// collide under SERIALIZABLE and the loser is retried (ADR-017
		// "Konsekuensi"), so each reader's positions stay 1, 2, 3, ….
		x := newHarness(t)
		var hh pgtype.UUID
		x.must(t, x.admin.QueryRow(t.Context(),
			"INSERT INTO household (name, time_zone, created_at) VALUES ('Rumah', 'Asia/Jakarta', $1) RETURNING id", start).Scan(&hh))
		members := []pgtype.UUID{x.addMember(t, hh), x.addMember(t, hh)}
		const workers, each = 2, 10
		errs := make(chan error, workers*each)
		var wg gosync.WaitGroup
		for range workers {
			wg.Go(func() {
				for range each {
					errs <- database.Write(t.Context(), x.app, database.Scope{HouseholdID: hh}, func(tx pgx.Tx) error {
						return sync.Changed(t.Context(), tx, hh, read.KindInvitation, newID())
					})
				}
			})
		}
		wg.Wait()
		close(errs)
		for err := range errs {
			x.must(t, err)
		}
		for _, m := range members {
			var got []int64
			for _, r := range x.feed(t, hh, m) {
				got = append(got, r.Pos)
			}
			want := make([]int64, workers*each)
			for i := range want {
				want[i] = int64(i + 1)
			}
			if !slices.Equal(got, want) {
				t.Errorf("member %s positions %v, want 1..%d", m, got, workers*each)
			}
			if n := x.lastPos(t, hh, m); n != workers*each {
				t.Errorf("member %s last_pos %d, want %d", m, n, workers*each)
			}
		}
	})
}
