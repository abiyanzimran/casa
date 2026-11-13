// Package command is the write-command path (09.2 §1, §6, §8): one
// serializable transaction per command that answers a repeated request from
// its idempotency record, refuses a reused key, or applies the command and
// stores its response (ADR-005).
package command

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/abiyanzimran/casa/server/internal/command/commanddb"
	"github.com/abiyanzimran/casa/server/internal/database"
)

// ErrIdempotencyConflict: the key was used before for another request
// (09.2 §8 idempotency_conflict).
var ErrIdempotencyConflict = errors.New("idempotency key reused with a different request")

// errSecretInResponse: a response would store a secret in
// idempotency_record (decision 8 of the 1.4 plan; SEC-11).
var errSecretInResponse = errors.New("a command response carries a debug_redact field")

// requestContextField is field 1 of every request (14 §4.6).
const requestContextField = "request_context"

// Run applies one write command at most once per (scope, idempotency key):
// inside one database.Write it returns the stored response of an earlier
// identical request, refuses another request under the same key, or runs
// apply and stores its response (ADR-005; 09.2 §6 steps 1 and 8, §8). The
// scope id is the household, or the account for an account-level command
// that carries no household; one that carries a household uses RunForAccount.
// apply may run more than once (serializable retry): no effect outside tx.
func Run[Response proto.Message](ctx context.Context, pool *pgxpool.Pool, scope database.Scope,
	procedure, idempotencyKey string, request proto.Message,
	apply func(context.Context, pgx.Tx) (Response, error),
) (Response, error) {
	return run(ctx, pool, scope, scopeID(scope), procedure, idempotencyKey, request, apply)
}

// RunForAccount is Run with the account as the idempotency scope while the
// transaction may also carry a household: an account-level command such as
// creating a household or accepting an invitation (09.2 §8).
func RunForAccount[Response proto.Message](ctx context.Context, pool *pgxpool.Pool, scope database.Scope,
	procedure, idempotencyKey string, request proto.Message,
	apply func(context.Context, pgx.Tx) (Response, error),
) (Response, error) {
	return run(ctx, pool, scope, scope.AccountID, procedure, idempotencyKey, request, apply)
}

func run[Response proto.Message](ctx context.Context, pool *pgxpool.Pool, scope database.Scope, scopeOf pgtype.UUID,
	procedure, idempotencyKey string, request proto.Message,
	apply func(context.Context, pgx.Tx) (Response, error),
) (Response, error) {
	var response Response
	r, err := newRecord(scopeOf, procedure, idempotencyKey, request)
	if err != nil {
		return response, err
	}
	err = database.Write(ctx, pool, scope, func(tx pgx.Tx) error {
		q := commanddb.New(tx)
		var found bool
		if response, found, err = stored[Response](ctx, q, r); err != nil || found {
			return err
		}
		if response, err = apply(ctx, tx); err != nil {
			return err
		}
		if hasSecret(response.ProtoReflect()) {
			return errSecretInResponse
		}
		stored, err := proto.MarshalOptions{Deterministic: true}.Marshal(response)
		if err != nil {
			return err
		}
		return q.InsertIdempotencyRecord(ctx, commanddb.InsertIdempotencyRecordParams{
			ScopeID: r.scopeID, IdempotencyKey: r.key, Command: procedure, RequestHash: r.hash, Response: stored,
		})
	})
	if err != nil {
		var zero Response
		return zero, err
	}
	return response, nil
}

// Stored returns the stored response of an earlier identical request, with
// the conflict rules of Run, in a read of its own: a handler calls it before
// work that must not repeat on a resend, such as an authorisation check the
// first send used up or a password hash (ADR-005; 09.2 §8). Run still
// decides; Stored only answers early. It keys on the household first, so an
// account-level caller passes an account-only scope.
func Stored[Response proto.Message](ctx context.Context, pool *pgxpool.Pool, scope database.Scope,
	procedure, idempotencyKey string, request proto.Message,
) (response Response, found bool, err error) {
	r, err := newRecord(scopeID(scope), procedure, idempotencyKey, request)
	if err != nil {
		return response, false, err
	}
	err = database.Read(ctx, pool, scope, func(tx pgx.Tx) error {
		response, found, err = stored[Response](ctx, commanddb.New(tx), r)
		return err
	})
	return response, found, err
}

// record is what identifies one command's idempotency record.
type record struct {
	scopeID, key pgtype.UUID
	procedure    string
	hash         []byte
}

// scopeID is the household, or the account when the scope has no household
// (09.2 §8).
func scopeID(scope database.Scope) pgtype.UUID {
	if scope.HouseholdID.Valid {
		return scope.HouseholdID
	}
	return scope.AccountID
}

func newRecord(scopeOf pgtype.UUID, procedure, idempotencyKey string, request proto.Message) (record, error) {
	r := record{scopeID: scopeOf, procedure: procedure}
	if err := r.key.Scan(idempotencyKey); err != nil {
		return r, fmt.Errorf("idempotency key: %w", err)
	}
	var err error
	r.hash, err = requestHash(request)
	return r, err
}

// stored reads r's record: the stored response, ErrIdempotencyConflict for
// another request under the same key, or not found.
func stored[Response proto.Message](ctx context.Context, q *commanddb.Queries, r record) (Response, bool, error) {
	var response Response
	row, err := q.GetIdempotencyRecord(ctx, commanddb.GetIdempotencyRecordParams{ScopeID: r.scopeID, IdempotencyKey: r.key})
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return response, false, nil
	case err != nil:
		return response, false, err
	case row.Command != r.procedure || !bytes.Equal(row.RequestHash, r.hash):
		return response, false, ErrIdempotencyConflict
	}
	response = newMessage[Response]()
	return response, true, proto.Unmarshal(row.Response, response)
}

func newMessage[M proto.Message]() M {
	var zero M
	return zero.ProtoReflect().New().Interface().(M)
}

// requestHash is SHA-256 of the deterministic encoding of request without
// request_context and without any debug_redact field (09.2 §8; decision 8 of
// the 1.4 plan): the hash of a password command must not be an unsalted hash
// of the password (SEC-11).
func requestHash(request proto.Message) ([]byte, error) {
	canonical := proto.Clone(request)
	message := canonical.ProtoReflect()
	if field := message.Descriptor().Fields().ByName(requestContextField); field != nil {
		message.Clear(field)
	}
	clearSecrets(message)
	encoded, err := proto.MarshalOptions{Deterministic: true}.Marshal(canonical)
	if err != nil {
		return nil, err
	}
	sum := sha256.Sum256(encoded)
	return sum[:], nil
}

func redacted(field protoreflect.FieldDescriptor) bool {
	options, ok := field.Options().(*descriptorpb.FieldOptions)
	return ok && options.GetDebugRedact()
}

// clearSecrets clears every debug_redact field of message, nested ones
// included.
func clearSecrets(message protoreflect.Message) {
	message.Range(func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		switch {
		case redacted(field):
			message.Clear(field)
		default:
			eachMessage(field, value, clearSecrets)
		}
		return true
	})
}

// hasSecret reports whether message sets a debug_redact field anywhere.
func hasSecret(message protoreflect.Message) bool {
	found := false
	message.Range(func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		found = redacted(field)
		eachMessage(field, value, func(m protoreflect.Message) { found = found || hasSecret(m) })
		return !found
	})
	return found
}

// eachMessage calls fn on every message held by a set field.
func eachMessage(field protoreflect.FieldDescriptor, value protoreflect.Value, fn func(protoreflect.Message)) {
	switch {
	case field.IsMap():
		if field.MapValue().Message() != nil {
			value.Map().Range(func(_ protoreflect.MapKey, v protoreflect.Value) bool { fn(v.Message()); return true })
		}
	case field.Message() == nil:
	case field.IsList():
		for i := range value.List().Len() {
			fn(value.List().Get(i).Message())
		}
	default:
		fn(value.Message())
	}
}
