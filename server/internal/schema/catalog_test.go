package schema_test

import (
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

// tables declares every table in schema public. A migration that adds a table
// adds its row here, citing the document that defines the table and its grants.
var tables = map[string]tableRule{
	"currency": {global: true, grants: map[string]string{"app_rw": "SELECT"},
		source: "09.2 §3 \"Tabel global\", §5.1; ADR-008"},
}

type tableRule struct {
	global bool              // no household_id column (09.2 §3 "Tabel global")
	grants map[string]string // role → privileges, comma-separated in the order of allPrivileges; absent role: none
	source string
}

// householdPolicy is the one policy of every household table, as PostgreSQL
// deparses 09.2 §3 layer 4.
const householdPolicy = "(household_id = (current_setting('app.household_id'::text))::uuid)"

var allPrivileges = []string{"SELECT", "INSERT", "UPDATE", "DELETE", "TRUNCATE", "REFERENCES", "TRIGGER", "MAINTAIN"}

// applicationRoles are the roles of 09.2 §3 layer 5 that row-level security
// binds; migrator owns the tables.
var applicationRoles = []string{"app_rw", "job_runner", "lifecycle_rw"}

// TestCatalog is the schema gate of ADR-016 items 2 to 4: a table that skips
// forced row-level security, the household policy, or its declared grants
// fails here, before any query uses it.
func TestCatalog(t *testing.T) {
	d := databasetest.Migrated(t)
	db := d.Pool(t, d.AdminRole)

	// Views, materialized views and foreign tables count too: a materialized
	// view has no row-level security, and a view runs as its owner.
	t.Run("every table is declared", func(t *testing.T) {
		got := queryStrings(t, db, `
			SELECT relname FROM pg_class
			 WHERE relnamespace = 'public'::regnamespace AND relkind IN ('r', 'p', 'v', 'm', 'f')
			   AND relname <> 'goose_db_version'
			 ORDER BY 1`)
		for _, name := range got {
			if _, ok := tables[name]; !ok {
				t.Errorf("table or view %s is not declared in tables", name)
			}
		}
		for _, name := range slices.Sorted(maps.Keys(tables)) {
			if !slices.Contains(got, name) {
				t.Errorf("table %s is declared in tables but does not exist", name)
			}
		}
	})

	t.Run("SEC-08 every table is owned by migrator and forces row-level security", func(t *testing.T) {
		rows, err := db.Query(t.Context(), `
			SELECT c.relname, c.relowner::regrole::text, c.relrowsecurity, c.relforcerowsecurity,
			       (SELECT count(*) FROM pg_policy p WHERE p.polrelid = c.oid)
			  FROM pg_class c
			 WHERE c.relnamespace = 'public'::regnamespace AND c.relkind IN ('r', 'p')
			   AND c.relname <> 'goose_db_version'`)
		if err != nil {
			t.Fatal(err)
		}
		type table struct {
			Name, Owner     string
			Enabled, Forced bool
			Policies        int64
		}
		got, err := pgx.CollectRows(rows, pgx.RowToStructByPos[table])
		if err != nil {
			t.Fatal(err)
		}
		for _, c := range got {
			if c.Owner != "migrator" || !c.Enabled || !c.Forced || c.Policies == 0 {
				t.Errorf("table %s: owner=%s rls=%t force=%t policies=%d; want owner migrator, rls and force on, a policy",
					c.Name, c.Owner, c.Enabled, c.Forced, c.Policies)
			}
		}
	})

	t.Run("SEC-08 a household table has exactly one policy, the household policy", func(t *testing.T) {
		for name, rule := range tables {
			hasHouseholdID := queryBool(t, db, `
				SELECT EXISTS (SELECT FROM pg_attribute
				                WHERE attrelid = format('public.%I', $1::text)::regclass
				                  AND attname = 'household_id' AND NOT attisdropped)`, name)
			if rule.global {
				if hasHouseholdID {
					t.Errorf("table %s is declared global but has household_id", name)
				}
				continue
			}
			if !hasHouseholdID {
				t.Errorf("household table %s has no household_id column", name)
			}
			policies := queryStrings(t, db, `
				SELECT format('cmd=%s roles=%s %s using=%s check=%s', cmd, roles, permissive, qual, with_check)
				  FROM pg_policies WHERE schemaname = 'public' AND tablename = $1
				 ORDER BY 1`, name)
			want := "cmd=ALL roles={public} PERMISSIVE using=" + householdPolicy + " check=" + householdPolicy
			if len(policies) != 1 || policies[0] != want {
				t.Errorf("table %s policies:\n%s\nwant exactly:\n%s", name, strings.Join(policies, "\n"), want)
			}
		}
	})

	t.Run("SEC-08 role privileges match the declaration", func(t *testing.T) {
		for name, rule := range tables {
			for _, role := range applicationRoles {
				if got := tablePrivileges(t, db, role, name); got != rule.grants[role] {
					t.Errorf("table %s, role %s: privileges %q, declared %q (%s)", name, role, got, rule.grants[role], rule.source)
				}
			}
			publicGrants := queryStrings(t, db, `
				SELECT a.privilege_type
				  FROM pg_class c, aclexplode(c.relacl) a
				 WHERE c.oid = format('public.%I', $1::text)::regclass AND a.grantee = 0
				 ORDER BY 1`, name)
			if len(publicGrants) != 0 {
				t.Errorf("table %s grants %v to PUBLIC", name, publicGrants)
			}
			// TRUNCATE is not subject to row-level security; REFERENCES,
			// TRIGGER and MAINTAIN are owner business.
			for _, privileges := range rule.grants {
				for _, p := range []string{"TRUNCATE", "REFERENCES", "TRIGGER", "MAINTAIN"} {
					if slices.Contains(strings.Split(privileges, ","), p) {
						t.Errorf("table %s declares %s for an application role", name, p)
					}
				}
			}
		}
		// has_table_privilege and relacl do not see column grants, and RLS is
		// no barrier for a role with them (any role can set
		// app.household_id). No table may carry one until the declaration can
		// express it (09.2 §3: job_runner's column UPDATE on job).
		columnGrants := queryStrings(t, db, `
			SELECT format('%s.%s %s', c.relname, a.attname, a.attacl)
			  FROM pg_attribute a JOIN pg_class c ON c.oid = a.attrelid
			 WHERE c.relnamespace = 'public'::regnamespace AND c.relname <> 'goose_db_version'
			   AND a.attacl IS NOT NULL
			 ORDER BY 1`)
		for _, g := range columnGrants {
			t.Errorf("column grant %s is not declared", g)
		}
	})

	t.Run("goose_db_version is not reachable by application roles", func(t *testing.T) {
		for _, role := range applicationRoles {
			if got := tablePrivileges(t, db, role, "goose_db_version"); got != "" {
				t.Errorf("role %s has %s on goose_db_version", role, got)
			}
		}
	})

	t.Run("the database and schema public grant nothing to PUBLIC", func(t *testing.T) {
		// A NULL ACL means the built-in default, which grants PUBLIC CONNECT,
		// TEMPORARY and, on schema public, USAGE.
		got := queryStrings(t, db, `
			SELECT 'database ' || a.privilege_type
			  FROM pg_database d, aclexplode(coalesce(d.datacl, acldefault('d', d.datdba))) a
			 WHERE d.datname = current_database() AND a.grantee = 0
			UNION ALL
			SELECT 'schema public ' || a.privilege_type
			  FROM pg_namespace n, aclexplode(coalesce(n.nspacl, acldefault('n', n.nspowner))) a
			 WHERE n.nspname = 'public' AND a.grantee = 0
			 ORDER BY 1`)
		if len(got) != 0 {
			t.Errorf("PUBLIC holds %v", got)
		}
		for _, role := range applicationRoles {
			if queryBool(t, db, "SELECT has_schema_privilege($1, 'public', 'CREATE')", role) {
				t.Errorf("role %s can create objects in schema public", role)
			}
		}
	})

	t.Run("household_id leads every key of a household table (SEC-02)", func(t *testing.T) {
		for name, rule := range tables {
			if rule.global {
				continue
			}
			bad := queryStrings(t, db, `
				WITH t AS (SELECT format('public.%I', $1::text)::regclass AS oid)
				SELECT pg_get_constraintdef(con.oid)
				  FROM pg_constraint con, t
				 WHERE con.conrelid = t.oid AND con.contype IN ('p', 'u', 'f')
				   AND ((SELECT attname FROM pg_attribute
				          WHERE attrelid = con.conrelid AND attnum = con.conkey[1]) IS DISTINCT FROM 'household_id'
				        OR (con.contype = 'f' AND (SELECT attname FROM pg_attribute
				              WHERE attrelid = con.confrelid AND attnum = con.confkey[1]) IS DISTINCT FROM 'household_id'))
				UNION ALL
				SELECT pg_get_indexdef(i.indexrelid)
				  FROM pg_index i, t
				 WHERE i.indrelid = t.oid AND i.indisunique
				   AND (SELECT attname FROM pg_attribute
				         WHERE attrelid = i.indrelid AND attnum = i.indkey[0]) IS DISTINCT FROM 'household_id'
				 ORDER BY 1`, name)
			for _, key := range bad {
				t.Errorf("table %s: %s does not start with household_id (SEC-02; 09.2 §3 layer 2)", name, key)
			}
		}
	})
}

// tablePrivileges returns the privileges role holds on table, comma-separated
// in the order of allPrivileges.
func tablePrivileges(t *testing.T, db *pgxpool.Pool, role, table string) string {
	t.Helper()
	var got string
	err := db.QueryRow(t.Context(), `
		SELECT coalesce(string_agg(p, ',' ORDER BY n), '')
		  FROM unnest($3::text[]) WITH ORDINALITY AS u(p, n)
		 WHERE has_table_privilege($1, format('public.%I', $2::text), p)`,
		role, table, allPrivileges).Scan(&got)
	if err != nil {
		t.Fatalf("privileges of %s on %s: %v", role, table, err)
	}
	return got
}

func queryStrings(t *testing.T, db *pgxpool.Pool, query string, args ...any) []string {
	t.Helper()
	rows, err := db.Query(t.Context(), query, args...)
	if err != nil {
		t.Fatal(err)
	}
	got, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func queryBool(t *testing.T, db *pgxpool.Pool, query string, args ...any) bool {
	t.Helper()
	var got bool
	if err := db.QueryRow(t.Context(), query, args...).Scan(&got); err != nil {
		t.Fatal(err)
	}
	return got
}
