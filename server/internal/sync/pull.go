package sync

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/read"
	"github.com/abiyanzimran/casa/server/internal/sync/syncdb"
)

// ChangeKind is what a device does with an object (09.3 §3).
type ChangeKind int

// The kinds of change of 09.3 §3.
const (
	Upsert ChangeKind = iota + 1
	Purge
)

// Change is one object for a device's cache: Upsert holds it, Purge drops
// it (09.3 §3; plan 1.7 decision 7).
type Change struct {
	Kind   ChangeKind
	Ref    read.ObjectRef
	Object read.Object // zero for Purge
}

// Page is one Sync answer (09.3 §3).
type Page struct {
	Changes        []Change
	Next           string
	HasMore        bool
	ResyncRequired bool
}

// Service reads Sync pages (09.2 §1 row sync).
type Service struct {
	pool *pgxpool.Pool
	read *read.Service
	key  CursorKey
}

// New returns the sync service over pool, projecting through r and sealing
// cursors with key.
func New(pool *pgxpool.Pool, r *read.Service, key CursorKey) *Service {
	return &Service{pool: pool, read: r, key: key}
}

// errBadLimit: the API sets the page size (plan 1.7 decision 14), so a
// limit below one is a programming error.
var errBadLimit = errors.New("sync page limit below one")

// Pull is one Sync page for r (09.3 §3; plan 1.7 decisions 9 to 11), in one
// REPEATABLE READ READ ONLY transaction (database.Read; ADR-017 item 7, note
// of 14 November 2026); read.ErrNotFound when r is no longer a member,
// ErrCursorRejected for a cursor not made for r. An empty cursor starts the
// snapshot.
func (s *Service) Pull(ctx context.Context, r read.Reader, cursor string, limit int) (Page, error) {
	if limit < 1 {
		return Page{}, errBadLimit
	}
	var c *Cursor
	if cursor != "" {
		decoded, err := s.key.Decode(cursor, r.HouseholdID, r.MemberID)
		if err != nil {
			return Page{}, err
		}
		c = &decoded
	}
	var page Page
	err := database.Read(ctx, s.pool, database.Scope{HouseholdID: r.HouseholdID, AccountID: r.AccountID}, func(tx pgx.Tx) error {
		if err := read.CheckReader(ctx, tx, r); err != nil {
			return err
		}
		last, err := syncdb.New(tx).LastPos(ctx, syncdb.LastPosParams{HouseholdID: r.HouseholdID, MemberID: r.MemberID})
		if err != nil {
			return err
		}
		switch {
		case c == nil:
			page, err = s.snapshot(ctx, tx, r, last, read.SnapshotMark{}, last, limit)
		case c.Snapshot != nil:
			page, err = s.snapshot(ctx, tx, r, c.Pos, *c.Snapshot, last, limit)
		default:
			page, err = s.feed(ctx, tx, r, c.Pos, last, limit)
		}
		return err
	})
	return page, err
}

// snapshot is a page of the initial snapshot from mark, with p0 the feed
// position its first page read (plan 1.7 decision 10). After the last page
// the cursor is the feed at p0, so every change after p0, between pages
// included, arrives through the feed.
func (s *Service) snapshot(ctx context.Context, tx pgx.Tx, r read.Reader, p0 int64, mark read.SnapshotMark, last int64, limit int) (Page, error) {
	objects, next, done, err := s.read.SnapshotPage(ctx, tx, r, mark, limit)
	if err != nil {
		return Page{}, err
	}
	page := Page{HasMore: !done || last > p0}
	for _, o := range objects {
		page.Changes = append(page.Changes, Change{Kind: Upsert, Ref: o.Ref, Object: o.Object})
	}
	c := Cursor{Household: r.HouseholdID, Member: r.MemberID, Pos: p0}
	if !done {
		c.Snapshot = &next
	}
	page.Next = s.key.Encode(c)
	return page, nil
}

// feed is the page of r's feed after p. It asks for resync_required, with
// no change and no cursor, when the feed that exists cannot continue p
// (09.3 §3 "Pemangkasan", note of 14 November 2026; plan 1.7 decision 11):
// p is beyond last, the rows are not exactly p+1, p+2, …, or a page that is
// not full does not end at last. None of it reads created_at.
func (s *Service) feed(ctx context.Context, tx pgx.Tx, r read.Reader, p, last int64, limit int) (Page, error) {
	if p > last {
		return Page{ResyncRequired: true}, nil
	}
	rows, err := syncdb.New(tx).FeedAfter(ctx, syncdb.FeedAfterParams{
		HouseholdID: r.HouseholdID, MemberID: r.MemberID, After: p, Lim: int32(limit + 1), //nolint:gosec // the API caps limit at 500 (plan 1.7 decision 14)
	})
	if err != nil {
		return Page{}, err
	}
	for i, row := range rows {
		if row.Pos != p+int64(i)+1 {
			return Page{ResyncRequired: true}, nil
		}
	}
	more := len(rows) > limit
	rows = rows[:min(limit, len(rows))]
	end := p
	if len(rows) > 0 {
		end = rows[len(rows)-1].Pos
	}
	if !more && end != last {
		return Page{ResyncRequired: true}, nil
	}

	refs := make([]read.ObjectRef, len(rows))
	for i, row := range rows {
		refs[i] = read.ObjectRef{Kind: read.ObjectKind(row.ObjectKind), ID: row.ObjectID}
	}
	objects, err := s.read.Objects(ctx, tx, r, refs)
	if err != nil {
		return Page{}, err
	}
	page := Page{HasMore: more, Next: s.key.Encode(Cursor{Household: r.HouseholdID, Member: r.MemberID, Pos: end})}
	for i, row := range rows {
		// A row that is purge, or an object r can no longer read or that is
		// no longer live, is a purge (09.3 §3; plan 1.7 decision 7).
		o, ok := objects[refs[i]]
		if row.Change == "purge" || !ok {
			page.Changes = append(page.Changes, Change{Kind: Purge, Ref: refs[i]})
			continue
		}
		page.Changes = append(page.Changes, Change{Kind: Upsert, Ref: refs[i], Object: o})
	}
	return page, nil
}
