package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // the database/sql driver "pgx"

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
	s, err := connString(migratorDatabaseURL, false)
	if err != nil {
		return 0, err
	}
	db, err := sql.Open("pgx", s)
	if err != nil {
		return 0, fmt.Errorf("open the database: %w", err)
	}
	defer func() { _ = db.Close() }()
	return schema.Up(ctx, db)
}

// connString reads the connection string in credential name and checks it
// parses, so no later parse error can quote it, password included (ES §4).
// An empty one is refused: pgx would fall back to the PG* environment
// variables. pooled: the string goes to pgxpool, which takes pool_
// parameters; otherwise it goes to one connection through database/sql,
// where pgx would send a pool_ parameter to the server as a setting, so it
// is refused here.
func connString(name string, pooled bool) (string, error) {
	s, err := readCredential(name)
	if err != nil {
		return "", err
	}
	if s == "" {
		return "", fmt.Errorf("credential %s is empty", name)
	}
	parse := func(s string) (map[string]string, error) {
		config, err := pgx.ParseConfig(s)
		if err != nil {
			return nil, err
		}
		return config.RuntimeParams, nil
	}
	if pooled {
		parse = func(s string) (map[string]string, error) {
			_, err := pgxpool.ParseConfig(s) // pool_ parameters first, then pgx's parse
			return nil, err
		}
	}
	settings, err := parse(s)
	if err != nil {
		// Replaced, not wrapped: the error can quote the string.
		return "", fmt.Errorf("cannot parse the connection string in credential %s", name)
	}
	for setting := range settings {
		if strings.HasPrefix(setting, "pool_") {
			return "", fmt.Errorf("credential %s has pool_ parameters, which do not apply to one connection", name)
		}
	}
	return s, nil
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
