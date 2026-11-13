package api

import (
	"errors"
	"fmt"
	"testing"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connectproto"

	"github.com/abiyanzimran/casa/server/internal/command"
	"github.com/abiyanzimran/casa/server/internal/database"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
)

func TestIdentityErrorRetryExhausted(t *testing.T) {
	// ADR-013 (1.4 note): an exhausted SERIALIZABLE retry is Aborted with
	// RETRY_EXHAUSTED (09.2 §6), so the client may try again.
	err := identityError(fmt.Errorf("open the session: %w", database.ErrRetryExhausted))
	var ce *connect.Error
	if !errors.As(err, &ce) || ce.Code() != connect.CodeAborted || ce.Message() != "retry exhausted" {
		t.Fatalf("got %v, want Aborted %q", err, "retry exhausted")
	}
	ds := ce.Details()
	if len(ds) != 1 {
		t.Fatalf("%d details, want 1", len(ds))
	}
	v, err := connectproto.UnmarshalErrorDetail(ds[0])
	if d, ok := v.(*casav1.ErrorDetail); err != nil || !ok || d.GetCode() != casav1.ErrorCode_ERROR_CODE_RETRY_EXHAUSTED {
		t.Fatalf("detail %v (%v), want ERROR_CODE_RETRY_EXHAUSTED", v, err)
	}
}

func TestIdentityErrorIdempotencyConflict(t *testing.T) {
	// 09.2 §8; ADR-013 (1.4 note): FailedPrecondition with
	// IDEMPOTENCY_CONFLICT, and the command is not applied.
	err := identityError(fmt.Errorf("sign out: %w", command.ErrIdempotencyConflict))
	var ce *connect.Error
	if !errors.As(err, &ce) || ce.Code() != connect.CodeFailedPrecondition || ce.Message() != "idempotency conflict" {
		t.Fatalf("got %v, want FailedPrecondition %q", err, "idempotency conflict")
	}
	ds := ce.Details()
	if len(ds) != 1 {
		t.Fatalf("%d details, want 1", len(ds))
	}
	v, err := connectproto.UnmarshalErrorDetail(ds[0])
	if d, ok := v.(*casav1.ErrorDetail); err != nil || !ok || d.GetCode() != casav1.ErrorCode_ERROR_CODE_IDEMPOTENCY_CONFLICT {
		t.Fatalf("detail %v (%v), want ERROR_CODE_IDEMPOTENCY_CONFLICT", v, err)
	}
}
