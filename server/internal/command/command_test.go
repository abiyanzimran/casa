package command

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
)

const procedure = "/casa.test.v1.Probe/Count"

type harness struct {
	pool, admin *pgxpool.Pool
	scope       database.Scope
	applied     int // apply calls, retries included
	mu          sync.Mutex
}

// start migrates a test database and makes probe_counter, a table apply
// increments, so a test can see an effect that committed.
func start(t *testing.T) *harness {
	t.Helper()
	d := databasetest.Migrated(t)
	h := &harness{pool: d.Pool(t, "app_rw"), admin: d.Pool(t, d.AdminRole), scope: database.Scope{AccountID: newUUID()}}
	_, err := h.admin.Exec(t.Context(), `CREATE TABLE probe_counter (n bigint NOT NULL);
		INSERT INTO probe_counter VALUES (0);
		GRANT SELECT, UPDATE ON probe_counter TO app_rw`)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

func newUUID() pgtype.UUID {
	var u pgtype.UUID
	_, _ = rand.Read(u.Bytes[:])
	u.Bytes[6] = u.Bytes[6]&0x0f | 0x40
	u.Bytes[8] = u.Bytes[8]&0x3f | 0x80
	u.Valid = true
	return u
}

func uuidString(u pgtype.UUID) string {
	b := u.Bytes
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// count is an apply that increments probe_counter and answers its value.
func (h *harness) count(ctx context.Context, tx pgx.Tx) (*wrapperspb.Int64Value, error) {
	h.mu.Lock()
	h.applied++
	h.mu.Unlock()
	var n int64
	err := tx.QueryRow(ctx, "UPDATE probe_counter SET n = n + 1 RETURNING n").Scan(&n)
	return wrapperspb.Int64(n), err
}

func (h *harness) counter(t *testing.T) int64 {
	t.Helper()
	var n int64
	if err := h.admin.QueryRow(t.Context(), "SELECT n FROM probe_counter").Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func (h *harness) records(t *testing.T) int {
	t.Helper()
	var n int
	if err := h.admin.QueryRow(t.Context(), "SELECT count(*) FROM idempotency_record").Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestRun(t *testing.T) {
	t.Run("a command sent twice is applied once and both get the first response", func(t *testing.T) {
		// ADR-005; 09.2 §8 (INV-10).
		h := start(t)
		key := uuidString(newUUID())
		for range 2 {
			got, err := Run(t.Context(), h.pool, h.scope, procedure, key, wrapperspb.String("a"), h.count)
			if err != nil || got.GetValue() != 1 {
				t.Fatalf("got %v, %v; want 1", got, err)
			}
		}
		var command string
		if err := h.admin.QueryRow(t.Context(), "SELECT command FROM idempotency_record").Scan(&command); err != nil || command != procedure {
			t.Fatalf("command %q, %v; want %q", command, err, procedure)
		}
		if h.applied != 1 || h.counter(t) != 1 || h.records(t) != 1 {
			t.Fatalf("applied %d, counter %d, records %d; want 1 each", h.applied, h.counter(t), h.records(t))
		}
	})
	t.Run("the same key with a different request is a conflict", func(t *testing.T) {
		h := start(t)
		key := uuidString(newUUID())
		if _, err := Run(t.Context(), h.pool, h.scope, procedure, key, wrapperspb.String("a"), h.count); err != nil {
			t.Fatal(err)
		}
		if _, err := Run(t.Context(), h.pool, h.scope, procedure, key, wrapperspb.String("b"), h.count); !errors.Is(err, ErrIdempotencyConflict) {
			t.Fatalf("err = %v, want ErrIdempotencyConflict", err)
		}
		if _, err := Run(t.Context(), h.pool, h.scope, "/casa.test.v1.Probe/Other", key, wrapperspb.String("a"), h.count); !errors.Is(err, ErrIdempotencyConflict) {
			t.Fatalf("another procedure: err = %v, want ErrIdempotencyConflict", err)
		}
		if h.applied != 1 {
			t.Fatalf("applied %d times, want 1", h.applied)
		}
	})
	t.Run("two concurrent sends commit once", func(t *testing.T) {
		h := start(t)
		key := uuidString(newUUID())
		var wg sync.WaitGroup
		results := make([]*wrapperspb.Int64Value, 2)
		errs := make([]error, 2)
		for i := range 2 {
			wg.Go(func() {
				results[i], errs[i] = Run(context.Background(), h.pool, h.scope, procedure, key, wrapperspb.String("a"), h.count)
			})
		}
		wg.Wait()
		for i := range 2 {
			if errs[i] != nil || results[i].GetValue() != 1 {
				t.Fatalf("send %d: %v, %v; want 1", i, results[i], errs[i])
			}
		}
		if h.counter(t) != 1 || h.records(t) != 1 {
			t.Fatalf("counter %d, records %d; want the effect once", h.counter(t), h.records(t))
		}
	})
	t.Run("the request context is not part of the request hash", func(t *testing.T) {
		h := start(t)
		key := uuidString(newUUID())
		request := func(version string, userInitiated bool) *casav1.SignInRequest {
			return &casav1.SignInRequest{
				RequestContext: &casav1.RequestContext{ClientKind: casav1.ClientKind_CLIENT_KIND_IOS, ClientVersion: version, UserInitiated: userInitiated, IdempotencyKey: key},
				Email:          "owner@example.test",
			}
		}
		apply := func(ctx context.Context, tx pgx.Tx) (*wrapperspb.Int64Value, error) { return h.count(ctx, tx) }
		if _, err := Run(t.Context(), h.pool, h.scope, procedure, key, request("1.0.0", true), apply); err != nil {
			t.Fatal(err)
		}
		if got, err := Run(t.Context(), h.pool, h.scope, procedure, key, request("1.0.1", false), apply); err != nil || got.GetValue() != 1 {
			t.Fatalf("got %v, %v; want the stored response", got, err)
		}
	})
	t.Run("debug_redact fields are not part of the request hash", func(t *testing.T) {
		// Decision 8 of the 1.4 plan: else the hash of a password command
		// is an unsalted hash of the password (SEC-11).
		h := start(t)
		key := uuidString(newUUID())
		request := func(password string) *casav1.SignInRequest {
			return &casav1.SignInRequest{
				RequestContext: &casav1.RequestContext{ClientKind: casav1.ClientKind_CLIENT_KIND_IOS, ClientVersion: "1.0.0", IdempotencyKey: key},
				Email:          "owner@example.test", Password: password,
			}
		}
		for _, password := range []string{"first passphrase", "second passphrase"} {
			if _, err := Run(t.Context(), h.pool, h.scope, procedure, key, request(password), h.count); err != nil {
				t.Fatalf("password %q: %v", password, err)
			}
		}
		cleared, err := proto.MarshalOptions{Deterministic: true}.Marshal(&casav1.SignInRequest{Email: "owner@example.test"})
		if err != nil {
			t.Fatal(err)
		}
		want := sha256.Sum256(cleared)
		var stored []byte
		if err := h.admin.QueryRow(t.Context(), "SELECT request_hash FROM idempotency_record").Scan(&stored); err != nil || string(stored) != string(want[:]) {
			t.Fatalf("request_hash %x, %v; want %x", stored, err, want)
		}
	})
	t.Run("a response with a debug_redact field is refused before it is stored", func(t *testing.T) {
		h := start(t)
		apply := func(ctx context.Context, tx pgx.Tx) (*casav1.SignInResponse, error) {
			_, err := h.count(ctx, tx)
			return &casav1.SignInResponse{Tokens: &casav1.SessionTokens{AccessToken: "secret"}}, err
		}
		if _, err := Run(t.Context(), h.pool, h.scope, procedure, uuidString(newUUID()), wrapperspb.String("a"), apply); !errors.Is(err, errSecretInResponse) {
			t.Fatalf("err = %v, want errSecretInResponse", err)
		}
		if h.records(t) != 0 || h.counter(t) != 0 {
			t.Fatalf("records %d, counter %d; want nothing committed", h.records(t), h.counter(t))
		}
	})
	t.Run("a failed apply leaves no record", func(t *testing.T) {
		h := start(t)
		key := uuidString(newUUID())
		failing := func(context.Context, pgx.Tx) (*wrapperspb.Int64Value, error) { return nil, errors.New("refused") }
		if _, err := Run(t.Context(), h.pool, h.scope, procedure, key, wrapperspb.String("a"), failing); err == nil {
			t.Fatal("no error")
		}
		if h.records(t) != 0 {
			t.Fatal("a record was stored for a failed apply")
		}
		if got, err := Run(t.Context(), h.pool, h.scope, procedure, key, wrapperspb.String("a"), h.count); err != nil || got.GetValue() != 1 {
			t.Fatalf("the next send: %v, %v; want it applied", got, err)
		}
	})
	t.Run("an account-level command uses the account as scope id", func(t *testing.T) {
		// 09.2 §8: account-level commands use scope_id = account_id.
		h := start(t)
		if _, err := Run(t.Context(), h.pool, h.scope, procedure, uuidString(newUUID()), wrapperspb.String("a"), h.count); err != nil {
			t.Fatal(err)
		}
		var scopeID pgtype.UUID
		if err := h.admin.QueryRow(t.Context(), "SELECT scope_id FROM idempotency_record").Scan(&scopeID); err != nil || scopeID != h.scope.AccountID {
			t.Fatalf("scope_id %v, %v; want the account %v", scopeID, err, h.scope.AccountID)
		}
	})
	t.Run("an account-level command keeps the account as scope id when a household is set", func(t *testing.T) {
		// 09.2 §8: CreateHousehold and AcceptInvitation carry a household in the
		// transaction but keep scope_id = account_id.
		h := start(t)
		household := newUUID()
		key := uuidString(newUUID())
		scope := database.Scope{HouseholdID: household, AccountID: h.scope.AccountID}
		if _, err := RunForAccount(t.Context(), h.pool, scope, procedure, key, wrapperspb.String("a"), h.count); err != nil {
			t.Fatal(err)
		}
		var scopeID pgtype.UUID
		if err := h.admin.QueryRow(t.Context(), "SELECT scope_id FROM idempotency_record").Scan(&scopeID); err != nil || scopeID != h.scope.AccountID {
			t.Fatalf("scope_id %v, %v; want the account %v", scopeID, err, h.scope.AccountID)
		}
		resend := database.Scope{HouseholdID: newUUID(), AccountID: h.scope.AccountID}
		if got, err := RunForAccount(t.Context(), h.pool, resend, procedure, key, wrapperspb.String("a"), h.count); err != nil || got.GetValue() != 1 {
			t.Fatalf("resend under another household: %v, %v; want the stored response", got, err)
		}
		if h.applied != 1 || h.records(t) != 1 {
			t.Fatalf("applied %d, records %d; want 1 each", h.applied, h.records(t))
		}
	})
}

func TestStored(t *testing.T) {
	t.Run("answers the stored response with the conflict rules of Run", func(t *testing.T) {
		// ADR-005; 09.2 §8.
		h := start(t)
		key := uuidString(newUUID())
		stored := func(request proto.Message) (*wrapperspb.Int64Value, bool, error) {
			return Stored[*wrapperspb.Int64Value](t.Context(), h.pool, h.scope, procedure, key, request)
		}
		if got, found, err := stored(wrapperspb.String("a")); err != nil || found || got != nil {
			t.Fatalf("before Run: %v, %t, %v; want not found", got, found, err)
		}
		if _, err := Run(t.Context(), h.pool, h.scope, procedure, key, wrapperspb.String("a"), h.count); err != nil {
			t.Fatal(err)
		}
		if got, found, err := stored(wrapperspb.String("a")); err != nil || !found || got.GetValue() != 1 {
			t.Fatalf("after Run: %v, %t, %v; want the first response", got, found, err)
		}
		if _, _, err := stored(wrapperspb.String("b")); !errors.Is(err, ErrIdempotencyConflict) {
			t.Fatalf("another request: err = %v, want ErrIdempotencyConflict", err)
		}
	})
}
