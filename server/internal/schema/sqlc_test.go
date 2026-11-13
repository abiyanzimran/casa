package schema_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestSQLCReadsEveryMigration keeps sqlc.yaml's schema list complete: sqlc
// reads the migrations named there, not the directory (roles.sql is not a
// migration), so a missed migration would generate code for an old schema
// (ES §2).
func TestSQLCReadsEveryMigration(t *testing.T) {
	config, err := os.ReadFile("../../sqlc.yaml")
	if err != nil {
		t.Fatal(err)
	}
	migrations, err := filepath.Glob("0*.sql")
	if err != nil || len(migrations) == 0 {
		t.Fatalf("no migrations found (%v)", err)
	}
	// Every module's block reads every migration.
	blocks := strings.Count(string(config), "- engine: postgresql")
	for _, m := range migrations {
		if n := strings.Count(string(config), "internal/schema/"+m); n != blocks {
			t.Errorf("migration %s is in %d of the %d sqlc.yaml blocks (sqlc would generate code for an old schema)", m, n, blocks)
		}
	}
}
