// Package database opens the transactions every module runs in, each scoped
// to one household and one account for row-level security (SEC-08; 09.2 §3
// layer 4, §6; ADR-016 items 3 and 6).
package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/schema"
)

var _ = schema.Roles

// Scope is the household and account a transaction acts for. It is built only
// from values already checked against the session's membership (09.2 §3 layer 1;
// SEC-01); an invalid (zero) field means "none".
type Scope struct {
	HouseholdID pgtype.UUID
	AccountID   pgtype.UUID
}

// setScope is the first statement of every transaction. is_local true is
// SET LOCAL: the settings end with the transaction, so a pooled connection
// never carries them into the next one (09.2 §6). Both are always set, to the
// empty string when absent, so every transaction starts the same whatever ran
// on its connection before; the empty string is not a uuid, so a household
// table read without a household fails instead of returning rows.
const setScope = `SELECT set_config('app.household_id', coalesce($1::uuid::text, ''), true),
       set_config('app.account_id', coalesce($2::uuid::text, ''), true)`

// Write runs fn in one SERIALIZABLE transaction (09.2 §6) with s set for that
// transaction only.
//
// ponytail: no retry on 40001 and 40P01 yet; it comes with the first
// concurrent writer (1.4), after the retry count of 09.2 §6 is settled.
func Write(ctx context.Context, pool *pgxpool.Pool, s Scope, fn func(pgx.Tx) error) error {
	return inTransaction(ctx, pool, pgx.TxOptions{IsoLevel: pgx.Serializable}, s, fn)
}

// Read runs fn in one REPEATABLE READ READ ONLY transaction (09.2 §6) with s set
// the same way.
func Read(ctx context.Context, pool *pgxpool.Pool, s Scope, fn func(pgx.Tx) error) error {
	return inTransaction(ctx, pool, pgx.TxOptions{IsoLevel: pgx.RepeatableRead, AccessMode: pgx.ReadOnly}, s, fn)
}

func inTransaction(ctx context.Context, pool *pgxpool.Pool, options pgx.TxOptions, s Scope, fn func(pgx.Tx) error) error {
	return pgx.BeginTxFunc(ctx, pool, options, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, setScope, s.HouseholdID, s.AccountID); err != nil {
			return fmt.Errorf("set transaction scope: %w", err)
		}
		return fn(tx)
	})
}
