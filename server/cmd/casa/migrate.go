package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/abiyanzimran/casa/server/internal/schema"
)

// migratorDatabaseURL names the LoadCredential= file holding the connection
// string of role migrator.
const migratorDatabaseURL = "migrator-database-url"

// migrate applies every pending migration as migrator. A release runs it
// before the new binary is installed (09.4 §1 "Rilis server"; 14 §6.4). There
// is no down direction here: a rollback installs the previous binary over the
// expand–contract schema (09.2 §10).
func migrate(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) != 0 {
		_, _ = fmt.Fprint(stderr, "usage: casa migrate\n")
		return 2
	}
	version, err := migrateUp(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "casa migrate: %v\n", err)
		return 1
	}
	_, _ = fmt.Fprintf(stdout, "casa migrate: database at version %d\n", version)
	return 0
}

func migrateUp(ctx context.Context) (int64, error) {
	connString, err := readCredential(migratorDatabaseURL)
	if err != nil {
		return 0, err
	}
	if connString == "" {
		// pgx would fall back to the PG* environment variables.
		return 0, fmt.Errorf("credential %s is empty", migratorDatabaseURL)
	}
	config, err := pgx.ParseConfig(connString)
	if err != nil {
		// The parse error can quote the connection string, password
		// included (ES §4), so it is replaced, not wrapped.
		return 0, fmt.Errorf("cannot parse the connection string in credential %s", migratorDatabaseURL)
	}
	db := stdlib.OpenDB(*config)
	defer func() { _ = db.Close() }()
	return schema.Up(ctx, db)
}

// readCredential reads the secret systemd passes with LoadCredential=, which
// keeps secrets out of the repository, the binary, and the environment
// (09.4 §1; SEC-42). Trailing whitespace and newlines are dropped.
func readCredential(name string) (string, error) {
	dir := os.Getenv("CREDENTIALS_DIRECTORY")
	if dir == "" {
		return "", errors.New("CREDENTIALS_DIRECTORY is not set; casa reads secrets from systemd LoadCredential= (09.4 §1)")
	}
	b, err := fs.ReadFile(os.DirFS(dir), name)
	if err != nil {
		return "", fmt.Errorf("read credential %s: %w", name, err)
	}
	return strings.TrimRight(string(b), " \t\r\n"), nil
}
