// Package schema holds the database schema of Casa: the cluster roles and the
// versioned goose migrations (ADR-016 items 4 and 11; 09.2 §3 layer 5, §10).
package schema

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

// Roles creates the database roles of 09.2 §3 layer 5 (SEC-09). It is run by a
// superuser once per cluster, not by a migration; see roles.sql.
//
//go:embed roles.sql
var Roles string

// migrations are the numbered goose files; roles.sql is not one of them.
//
//go:embed 0*.sql
var migrations embed.FS

// NewProvider returns a goose provider over the embedded migrations. Each
// migration has an up and a down part (ES §7; 09.2 §10).
func NewProvider(db *sql.DB) (*goose.Provider, error) {
	return goose.NewProvider(goose.DialectPostgres, db, migrations, goose.WithDisableGlobalRegistry(true))
}

// Up applies every pending migration and returns the resulting version. It
// runs only as migrator: tables created by any other role would be owned by
// it, and a superuser owner bypasses row-level security (SEC-08, SEC-09;
// 09.2 §3 layer 5).
func Up(ctx context.Context, db *sql.DB) (int64, error) {
	var role string
	if err := db.QueryRowContext(ctx, "SELECT current_user").Scan(&role); err != nil {
		return 0, fmt.Errorf("read current role: %w", err)
	}
	if role != "migrator" {
		return 0, fmt.Errorf("must run as role migrator, connected as %s", role)
	}
	p, err := NewProvider(db)
	if err != nil {
		return 0, fmt.Errorf("load migrations: %w", err)
	}
	if _, err := p.Up(ctx); err != nil {
		return 0, fmt.Errorf("apply migrations: %w", err)
	}
	version, err := p.GetDBVersion(ctx)
	if err != nil {
		return 0, fmt.Errorf("read schema version: %w", err)
	}
	return version, nil
}
