// Package database opens the transactions every module runs in, each scoped
// to one household and one account for row-level security (SEC-08; 09.2 §3
// layer 4, §6; ADR-016 items 3 and 6).
package database

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Scope is the household and account a transaction acts for. It is built only
// from values already checked against the session's membership (09.2 §3 layer 1;
// SEC-01); an invalid (zero) field means "none".
type Scope struct {
	HouseholdID pgtype.UUID
	AccountID   pgtype.UUID
	// Email and TokenHash let identity find one row before the account is
	// known (sign-in, token use); each is a lookup key, never a grant
	// (09.2 §3 global tables; 10.1 §5.1). An email is allowed only alone in
	// a Read (ErrEmailLookupNotAlone).
	Email     string
	TokenHash []byte
}

// ErrEmailLookupNotAlone refuses a scope that carries an email anywhere but
// in a read with no account and no household: an email only finds an
// account's id before sign-in (10.1 §5.1 rule 3; 09.2 §3).
var ErrEmailLookupNotAlone = errors.New("an email lookup runs alone in a read transaction")

// ErrRetryExhausted wraps the last 40001 or 40P01 after maxRetries retries
// (09.2 §6, §9).
var ErrRetryExhausted = errors.New("serializable transaction retries exhausted")

// maxRetries is "paling banyak 5 kali" of 09.2 §6 and ADR-016 item 6: five
// retries after the first attempt (lead ruling Q1 on the 1.4 plan).
const maxRetries = 5

// firstBackoff is the upper bound of the first retry's random wait; it
// doubles with each retry (09.2 §6 "jeda acak yang membesar") [Usulan].
const firstBackoff = 10 * time.Millisecond

// setScope is the first statement of every transaction. is_local true is
// SET LOCAL: the settings end with the transaction, so a pooled connection
// never carries them into the next one (09.2 §6). All are always set, to the
// empty string when absent, so every transaction starts the same whatever ran
// on its connection before; the empty string is not a uuid, so a household
// table read without a household fails instead of returning rows.
const setScope = `SELECT set_config('app.household_id', coalesce($1::uuid::text, ''), true),
       set_config('app.account_id', coalesce($2::uuid::text, ''), true),
       set_config('app.email', $3::text, true),
       set_config('app.token_hash', coalesce(encode($4::bytea, 'hex'), ''), true)`

// Write runs fn in one SERIALIZABLE transaction (09.2 §6) with s set for that
// transaction only. On a serialization failure (40001) or a deadlock (40P01)
// it runs fn again in a new transaction, up to maxRetries times, after a
// random wait. fn may therefore run up to six times and must have no effect
// outside its transaction: mail, password hashing and rate-limit updates
// happen before or after Write, never inside fn.
func Write(ctx context.Context, pool *pgxpool.Pool, s Scope, fn func(pgx.Tx) error) error {
	if s.Email != "" {
		return ErrEmailLookupNotAlone
	}
	options := pgx.TxOptions{IsoLevel: pgx.Serializable}
	for retry := 0; ; retry++ {
		err := inTransaction(ctx, pool, options, s, fn)
		if !retryable(err) {
			return err
		}
		if retry == maxRetries {
			return fmt.Errorf("%w: %w", ErrRetryExhausted, err)
		}
		// Full jitter in [0, firstBackoff × 2^retry).
		wait := time.NewTimer(rand.N(firstBackoff << retry)) //nolint:gosec // G404: a retry delay needs no unpredictability
		select {
		case <-ctx.Done():
			wait.Stop()
			return errors.Join(ctx.Err(), err)
		case <-wait.C:
		}
	}
}

func retryable(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && (pgErr.Code == "40001" || pgErr.Code == "40P01")
}

// Read runs fn in one REPEATABLE READ READ ONLY transaction (09.2 §6) with s set
// the same way. It is never retried.
func Read(ctx context.Context, pool *pgxpool.Pool, s Scope, fn func(pgx.Tx) error) error {
	if s.Email != "" && (s.AccountID.Valid || s.HouseholdID.Valid) {
		return ErrEmailLookupNotAlone
	}
	return inTransaction(ctx, pool, pgx.TxOptions{IsoLevel: pgx.RepeatableRead, AccessMode: pgx.ReadOnly}, s, fn)
}

func inTransaction(ctx context.Context, pool *pgxpool.Pool, options pgx.TxOptions, s Scope, fn func(pgx.Tx) error) error {
	return pgx.BeginTxFunc(ctx, pool, options, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, setScope, s.HouseholdID, s.AccountID, s.Email, s.TokenHash); err != nil {
			return fmt.Errorf("set transaction scope: %w", err)
		}
		return fn(tx)
	})
}
