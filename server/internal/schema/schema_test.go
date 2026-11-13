package schema_test

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
	"github.com/abiyanzimran/casa/server/internal/schema"
)

// TestEveryMigrationGoesUpDownUp proves each migration reversible (ES §7;
// 09.4 §4 row "Migrasi"; 14 §6.4): its down restores the catalog exactly, and
// its up applies again after that.
func TestEveryMigrationGoesUpDownUp(t *testing.T) {
	d := databasetest.Empty(t)
	db := d.SQL(t, "migrator")
	p, err := schema.NewProvider(db)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.ListSources()) == 0 {
		t.Fatal("no migrations embedded")
	}
	ctx := t.Context()
	before := catalogSnapshot(t, db)
	for _, m := range p.ListSources() {
		if _, err := p.UpByOne(ctx); err != nil {
			t.Fatalf("migration %d up: %v", m.Version, err)
		}
		after := catalogSnapshot(t, db)
		if _, err := p.Down(ctx); err != nil {
			t.Fatalf("migration %d down: %v", m.Version, err)
		}
		if got := catalogSnapshot(t, db); got != before {
			t.Fatalf("migration %d: down does not restore the catalog\nbefore:\n%s\nafter down:\n%s", m.Version, before, got)
		}
		if _, err := p.UpByOne(ctx); err != nil {
			t.Fatalf("migration %d up again: %v", m.Version, err)
		}
		if got := catalogSnapshot(t, db); got != after {
			t.Fatalf("migration %d: up after down differs from the first up\nfirst:\n%s\nagain:\n%s", m.Version, after, got)
		}
		before = after
	}
}

func TestUpRefusesARoleOtherThanMigrator(t *testing.T) {
	d := databasetest.Empty(t)
	db := d.SQL(t, d.AdminRole)
	_, err := schema.Up(t.Context(), db)
	want := "must run as role migrator, connected as " + d.AdminRole
	if err == nil || !strings.Contains(err.Error(), want) {
		t.Fatalf("Up as %s: err = %v, want it to contain %q", d.AdminRole, err, want)
	}
	var untouched bool
	if err := db.QueryRowContext(t.Context(), "SELECT to_regclass('goose_db_version') IS NULL").Scan(&untouched); err != nil {
		t.Fatal(err)
	}
	if !untouched {
		t.Fatal("Up created goose_db_version before refusing the role")
	}
}

func TestCurrency(t *testing.T) {
	d := databasetest.Migrated(t)
	app := d.Pool(t, "app_rw")

	t.Run("holds IDR with the product exponent 0", func(t *testing.T) {
		rows, err := app.Query(t.Context(), "SELECT code, exponent FROM currency ORDER BY code")
		if err != nil {
			t.Fatal(err)
		}
		type currency struct {
			Code     string
			Exponent int16
		}
		got, err := pgx.CollectRows(rows, pgx.RowToStructByPos[currency])
		if err != nil {
			t.Fatal(err)
		}
		// ADR-008; 07 §2.1, INV-39: IDR has exponent 0 in this product, not
		// the ISO 4217 minor unit 2.
		if len(got) != 1 || got[0] != (currency{"IDR", 0}) {
			t.Fatalf("currency = %v, want [{IDR 0}]", got)
		}
	})

	t.Run("app_rw cannot change it", func(t *testing.T) {
		// 09.2 §3 "Tabel global": currency is SELECT only for app_rw.
		for _, statement := range []string{
			"INSERT INTO currency (code, exponent) VALUES ('USD', 2)",
			"UPDATE currency SET exponent = 2",
			"DELETE FROM currency",
		} {
			_, err := app.Exec(t.Context(), statement)
			if code := databasetest.PgErrorCode(err); code != "42501" {
				t.Errorf("%s: err = %v, want insufficient_privilege (42501)", statement, err)
			}
		}
	})
}

// grantee names an aclexplode grantee; 0 is PUBLIC.
const grantee = "CASE a.grantee WHEN 0 THEN 'PUBLIC' ELSE a.grantee::regrole::text END"

// catalogQueries describe schema public, the database's privileges, and the
// installed extensions. goose_db_version is left out: it records the version,
// so it differs by design. Privileges are compared as effective privileges
// (aclexplode over the ACL or its default), not as ACL text, so a default
// (NULL) ACL equals the explicit ACL a down migration grants back.
var catalogQueries = []struct{ name, query string }{
	{"relations", `
		SELECT format('%s kind=%s owner=%s rls=%s force=%s', relname, relkind,
			relowner::regrole, relrowsecurity, relforcerowsecurity)
		  FROM pg_class
		 WHERE relnamespace = 'public'::regnamespace AND relname NOT LIKE 'goose\_db\_version%'`},
	{"relation privileges", `
		SELECT format('%s %s %s grantor=%s grantable=%s', c.relname, ` + grantee + `,
			a.privilege_type, a.grantor::regrole, a.is_grantable)
		  FROM pg_class c,
		       aclexplode(coalesce(c.relacl, acldefault((CASE c.relkind WHEN 'S' THEN 's' ELSE 'r' END)::"char", c.relowner))) a
		 WHERE c.relnamespace = 'public'::regnamespace AND c.relname NOT LIKE 'goose\_db\_version%'`},
	{"columns", `
		SELECT format('%s.%s %s not_null=%s default=%s', c.relname, a.attname,
			format_type(a.atttypid, a.atttypmod), a.attnotnull, pg_get_expr(d.adbin, d.adrelid))
		  FROM pg_attribute a
		  JOIN pg_class c ON c.oid = a.attrelid
		  LEFT JOIN pg_attrdef d ON d.adrelid = a.attrelid AND d.adnum = a.attnum
		 WHERE c.relnamespace = 'public'::regnamespace AND c.relname NOT LIKE 'goose\_db\_version%'
		   AND a.attnum > 0 AND NOT a.attisdropped`},
	{"column privileges", `
		SELECT format('%s.%s %s %s', c.relname, at.attname, ` + grantee + `, a.privilege_type)
		  FROM pg_attribute at
		  JOIN pg_class c ON c.oid = at.attrelid,
		       aclexplode(at.attacl) a
		 WHERE c.relnamespace = 'public'::regnamespace AND NOT at.attisdropped`},
	{"constraints", `
		SELECT format('%s %s %s', conrelid::regclass, conname, pg_get_constraintdef(oid))
		  FROM pg_constraint
		 WHERE connamespace = 'public'::regnamespace AND conname NOT LIKE 'goose\_db\_version%'`},
	{"indexes", `
		SELECT indexdef FROM pg_indexes
		 WHERE schemaname = 'public' AND tablename <> 'goose_db_version'`},
	{"policies", `
		SELECT format('%s %s %s %s %s using=%s check=%s', tablename, policyname, permissive,
			roles, cmd, qual, with_check)
		  FROM pg_policies WHERE schemaname = 'public'`},
	{"triggers", `
		SELECT pg_get_triggerdef(t.oid)
		  FROM pg_trigger t JOIN pg_class c ON c.oid = t.tgrelid
		 WHERE c.relnamespace = 'public'::regnamespace AND NOT t.tgisinternal`},
	{"functions", `
		SELECT pg_get_functiondef(oid) FROM pg_proc
		 WHERE pronamespace = 'public'::regnamespace AND prokind IN ('f', 'p')`},
	{"function privileges", `
		SELECT format('%s %s %s', p.oid::regprocedure, ` + grantee + `, a.privilege_type)
		  FROM pg_proc p, aclexplode(coalesce(p.proacl, acldefault('f', p.proowner))) a
		 WHERE p.pronamespace = 'public'::regnamespace`},
	{"types", `
		SELECT format('%s kind=%s', typname, typtype) FROM pg_type
		 WHERE typnamespace = 'public'::regnamespace AND typname NOT LIKE '%goose\_db\_version%'`},
	{"extensions", `SELECT format('%s %s', extname, extversion) FROM pg_extension`},
	{"default privileges", `
		SELECT format('%s %s %s', defaclrole::regrole, defaclobjtype, defaclacl) FROM pg_default_acl`},
	{"schema public privileges", `
		SELECT format('%s %s grantable=%s', ` + grantee + `, a.privilege_type, a.is_grantable)
		  FROM pg_namespace n, aclexplode(coalesce(n.nspacl, acldefault('n', n.nspowner))) a
		 WHERE n.nspname = 'public'`},
	{"database privileges", `
		SELECT format('%s %s grantable=%s', ` + grantee + `, a.privilege_type, a.is_grantable)
		  FROM pg_database d, aclexplode(coalesce(d.datacl, acldefault('d', d.datdba))) a
		 WHERE d.datname = current_database()`},
}

// catalogSnapshot returns catalogQueries' results as sorted text, one section
// per query.
func catalogSnapshot(t *testing.T, db *sql.DB) string {
	t.Helper()
	var b strings.Builder
	for _, q := range catalogQueries {
		rows, err := db.QueryContext(t.Context(), q.query+" ORDER BY 1")
		if err != nil {
			t.Fatalf("catalog %s: %v", q.name, err)
		}
		b.WriteString("-- " + q.name + "\n")
		for rows.Next() {
			var line string
			if err := rows.Scan(&line); err != nil {
				t.Fatalf("catalog %s: %v", q.name, err)
			}
			b.WriteString(line + "\n")
		}
		if err := rows.Err(); err != nil {
			t.Fatalf("catalog %s: %v", q.name, err)
		}
		_ = rows.Close()
	}
	return b.String()
}
