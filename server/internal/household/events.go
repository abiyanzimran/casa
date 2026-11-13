package household

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/household/householddb"
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

// RecordEvent is the one writer of object_event (09.2 §1; INV-A01).
func RecordEvent(ctx context.Context, tx pgx.Tx, e Event) error {
	return householddb.New(tx).InsertObjectEvent(ctx, householddb.InsertObjectEventParams{
		HouseholdID: e.HouseholdID, ObjectKind: e.ObjectKind, ObjectID: e.ObjectID, Seq: e.Seq,
		ActorKind: e.ActorKind, Actor: e.Actor, Rule: optionalText(e.Rule),
		Command: e.Command, FromState: optionalText(e.FromState), ToState: e.ToState,
		IdempotencyKey: e.IdempotencyKey, ServerTime: e.At,
	})
}

// optionalText is NULL for "".
func optionalText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: s != ""}
}
