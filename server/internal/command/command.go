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
// scope id is the household, or the account for an account-level command.
// apply may run more than once (serializable retry): no effect outside tx.
func Run[Response proto.Message](ctx context.Context, pool *pgxpool.Pool, scope database.Scope,
	procedure, idempotencyKey string, request proto.Message,
	apply func(context.Context, pgx.Tx) (Response, error),
) (Response, error) {
	var response Response
	var key pgtype.UUID
	if err := key.Scan(idempotencyKey); err != nil {
		return response, fmt.Errorf("idempotency key: %w", err)
	}
	scopeID := scope.HouseholdID
	if !scopeID.Valid {
		scopeID = scope.AccountID
	}
	hash, err := requestHash(request)
	if err != nil {
		return response, err
	}
	err = database.Write(ctx, pool, scope, func(tx pgx.Tx) error {
		q := commanddb.New(tx)
		record, err := q.GetIdempotencyRecord(ctx, commanddb.GetIdempotencyRecordParams{ScopeID: scopeID, IdempotencyKey: key})
		switch {
		case err == nil && (record.Command != procedure || !bytes.Equal(record.RequestHash, hash)):
			return ErrIdempotencyConflict
		case err == nil:
			response = newMessage[Response]()
			return proto.Unmarshal(record.Response, response)
		case !errors.Is(err, pgx.ErrNoRows):
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
			ScopeID: scopeID, IdempotencyKey: key, Command: procedure, RequestHash: hash, Response: stored,
		})
	})
	if err != nil {
		var zero Response
		return zero, err
	}
	return response, nil
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
