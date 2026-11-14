package household

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/household/householddb"
	"github.com/abiyanzimran/casa/server/internal/read"
	"github.com/abiyanzimran/casa/server/internal/sync"
)

// Event is one object event (06 §2.5; INV-A01): a transition of one
// object, by a member, an invitee or a rule of the system.
type Event struct {
	HouseholdID pgtype.UUID
	ObjectKind  string
	ObjectID    pgtype.UUID
	Seq         int64 // the object's version after the transition; 0 for its creation (09.2 §5.1)
	ActorKind   string
	Actor       pgtype.UUID
	Rule        string // Actor only for member, Rule only for system (Q14)
	// Command is the command of 06 that made the transition; FromState is
	// "" for a creation.
	Command, FromState, ToState string
	IdempotencyKey              pgtype.UUID
	At                          time.Time
}

// feedKinds maps an object_event kind to the kind the feed carries (plan 1.7
// Task 4).
var feedKinds = map[string]read.ObjectKind{
	"household":  read.KindHousehold,
	"invitation": read.KindInvitation,
}

// RecordEvent is the one writer of object_event (09.2 §1; INV-A01). Every
// change of a household object writes an event, so the event's transaction
// also writes the object's feed rows through sync.Changed, and no write path
// can forget them (09.2 §1 item 10, §6 step 6; ADR-017 item 3; plan 1.7
// "Fan-out lewat satu titik").
func RecordEvent(ctx context.Context, tx pgx.Tx, e Event) error {
	kind, ok := feedKinds[e.ObjectKind]
	if !ok {
		return fmt.Errorf("record event: object kind %q has no feed kind", e.ObjectKind)
	}
	if err := householddb.New(tx).InsertObjectEvent(ctx, householddb.InsertObjectEventParams{
		HouseholdID: e.HouseholdID, ObjectKind: e.ObjectKind, ObjectID: e.ObjectID, Seq: e.Seq,
		ActorKind: e.ActorKind, Actor: e.Actor, Rule: optionalText(e.Rule),
		Command: e.Command, FromState: optionalText(e.FromState), ToState: e.ToState,
		IdempotencyKey: e.IdempotencyKey, ServerTime: e.At,
	}); err != nil {
		return err
	}
	return sync.Changed(ctx, tx, e.HouseholdID, kind, e.ObjectID)
}

// optionalText is NULL for "".
func optionalText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: s != ""}
}
