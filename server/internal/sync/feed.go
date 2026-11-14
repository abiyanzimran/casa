// Package sync is the per-reader feed, its cursor and the Sync read (09.2 §1
// row sync; ADR-017 item 3; ADR-018). It takes the members and the read rule
// from read, never from household (09.2 §1 item 10). A file that needs the
// standard library's sync imports it as gosync.
package sync

import (
	"bytes"
	"context"
	"maps"
	"slices"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/read"
	"github.com/abiyanzimran/casa/server/internal/sync/syncdb"
)

// Changed writes one upsert row per reader of the object in tx, the command's
// transaction (09.2 §6 step 6; ADR-017 item 3): it is the one writer of
// reader_feed (09.2 §1 item 10). A household object also writes every member
// object, since the membership changes only with a household event (plan 1.7
// decision 8; 09.2 §5.1). Each reader's rows take the next positions of its
// counter, without a gap (09.3 §3).
func Changed(ctx context.Context, tx pgx.Tx, household pgtype.UUID, kind read.ObjectKind, id pgtype.UUID) error {
	objects := []read.ObjectRef{{Kind: kind, ID: id}}
	if kind == read.KindHousehold {
		members, err := read.Members(ctx, tx, household)
		if err != nil {
			return err
		}
		for _, m := range members {
			objects = append(objects, read.ObjectRef{Kind: read.KindMember, ID: m})
		}
	}
	// Each reader gets, in order, the objects whose kind it can read.
	readersOf := map[read.ObjectKind][]pgtype.UUID{}
	perReader := map[pgtype.UUID][]read.ObjectRef{}
	for _, o := range objects {
		readers, ok := readersOf[o.Kind]
		if !ok {
			var err error
			if readers, err = read.Readers(ctx, tx, household, o.Kind); err != nil {
				return err
			}
			readersOf[o.Kind] = readers
		}
		for _, r := range readers {
			perReader[r] = append(perReader[r], o)
		}
	}
	// One order of counter rows for every command, so two commands wait on
	// each other instead of deadlocking. ponytail: untested; if the order
	// broke, the cost is a deadlock found after deadlock_timeout and retried
	// by database.Write (40P01), never a wrong position.
	readers := slices.SortedFunc(maps.Keys(perReader), func(a, b pgtype.UUID) int { return bytes.Compare(a.Bytes[:], b.Bytes[:]) })

	q := syncdb.New(tx)
	rows := syncdb.InsertFeedRowsParams{HouseholdID: household}
	for _, r := range readers {
		own := perReader[r]
		last, err := q.BumpCounter(ctx, syncdb.BumpCounterParams{HouseholdID: household, MemberID: r, N: int64(len(own))})
		if err != nil {
			return err
		}
		for i, o := range own {
			rows.MemberIds = append(rows.MemberIds, r)
			rows.Positions = append(rows.Positions, last-int64(len(own))+int64(i)+1)
			rows.ObjectKinds = append(rows.ObjectKinds, string(o.Kind))
			rows.ObjectIds = append(rows.ObjectIds, o.ID)
		}
	}
	if len(readers) == 0 {
		return nil
	}
	return q.InsertFeedRows(ctx, rows)
}
