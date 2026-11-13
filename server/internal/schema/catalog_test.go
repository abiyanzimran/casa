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
		policy: "cmd=SELECT roles={public} PERMISSIVE using=true check=",
		source: "09.2 §3 \"Tabel global\", §5.1; ADR-008"},
	"account": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policy: lookupPolicy("((id = "+settingAccount+") OR (email = "+settingEmail+"))", "(id = "+settingAccount+")"),
		source: "09.2 §3 \"Tabel global\", §5.1; lead rulings Q2, Q3 on the 1.4 plan"},
	"credential": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"salt", "hash", "time_cost", "memory_kib", "parallelism", "updated_at"},
		policy:  ownAccount,
		source:  "09.2 §3 \"Tabel global\"; SEC-11; lead ruling Q3 on the 1.4 plan"},
	"session": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		// 00004_web_session: the cookie columns, and no UPDATE on device_mode
		// or ua_summary (SEC-165).
		updates: []string{"token_hash", "token_issued_at", "last_used_at", "revoked_at", "revoke_reason",
			"prev_token_hash", "prev_valid_until", "token_confirmed_at", "discarded_token_hash"},
		policy: lookupPolicy("((account_id = "+settingAccount+") OR (token_hash = "+settingTokenHash+
			") OR (prev_token_hash = "+settingTokenHash+") OR (discarded_token_hash = "+settingTokenHash+"))",
			"(account_id = "+settingAccount+")"),
		source: "09.2 §3 \"Tabel global\", §5.1; SEC-15, SEC-161 (c); lead rulings Q2, Q3 on the 1.4 plan; plan 1.5 decision 15"},
	"refresh_token": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"used_at"},
		policy:  accountOrToken,
		source:  "09.2 §3 \"Tabel global\"; SEC-15; lead rulings Q2, Q3 on the 1.4 plan"},
	"reauth_grant": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"used_at"},
		policy:  ownAccount,
		source:  "09.2 §3 \"Tabel global\"; SEC-13; SEC-161 (c) note 1.5 rule (7) (cookie_hash, 00004); lead ruling Q3 on the 1.4 plan"},
	"recovery_token": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"used_at"},
		policy:  accountOrToken,
		source:  "09.2 §3 \"Tabel global\"; SEC-24; lead rulings Q2, Q3 on the 1.4 plan"},
	"setup_token": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"used_at", "account_id"},
		policy:  lookupPolicy("(token_hash = "+settingTokenHash+")", "(token_hash = "+settingTokenHash+")"),
		source:  "09.2 §3 \"Tabel global\", §5.1; SEC-158 (b); 09.4 §2; lead rulings Q2, Q3 on the 1.4 plan"},
	"signup_verification": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"token_hash", "expires_at", "used_at"},
		policy: lookupPolicy("((token_hash = "+settingTokenHash+") OR (setup_token_hash = "+settingTokenHash+"))",
			"(setup_token_hash = "+settingTokenHash+")"),
		source: "09.2 §3 \"Tabel global\", §5.1; SEC-158; lead rulings Q2, Q4, Q8 on the 1.4 plan"},
	"security_event": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policy: ownAccount,
		source: "09.2 §3 \"Tabel global\"; SEC-16, SEC-39 (b): append-only; lead ruling Q3 on the 1.4 plan"},
	"idempotency_record": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policy: "cmd=ALL roles={public} PERMISSIVE using=((scope_id = " + settingHousehold + ") OR (scope_id = " + settingAccount + ")) check=",
		source: "09.2 §3 \"Tabel infrastruktur dengan pemilik campuran\", §5.1, §8"},
	// 00005_household: the household tables (09.2 §5.1, note of 14 November
	// 2026; plan 1.6 decisions 2, 3). app_rw never deletes a household row.
	"household": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"version"},
		policy:  lookupPolicy("(id = "+strictHousehold+")", "(id = "+strictHousehold+")"),
		source:  "09.2 §3 note of 14 November 2026, §5.1; plan 1.6 decisions 2, 3"},
	"membership": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policy: lookupPolicy("((household_id = "+settingHousehold+") OR (("+settingHousehold+" IS NULL) AND (account_id = "+settingAccount+")))",
			"(household_id = "+settingHousehold+")"),
		source: "09.2 §3 note of 14 November 2026, §5.1; 10.1 §5.1 note 1.6 item 1; 06 §3.1"},
	"approval_authority": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		source: "09.2 §5.1 note of 14 November 2026; 06 §3.1"},
	"household_rule": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		source: "09.2 §5.1 note of 14 November 2026; 06 §3.1"},
	// No writer until Tier 3 (06 §3.1 row "Hubungan"; FR-FAM-006/AC-3).
	"relationship": {grants: map[string]string{"app_rw": "SELECT"},
		source: "09.2 §5.1 note of 14 November 2026; 06 §3.1"},
	"invitation": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"token_hash", "state", "invitee_account_id", "invitee_display_name", "closed_at", "version"},
		policy: lookupPolicy("((household_id = "+settingHousehold+") OR (("+settingHousehold+" IS NULL) AND (token_hash = "+settingTokenHash+
			")) OR (("+settingHousehold+" IS NULL) AND (invitee_account_id = "+settingAccount+")))",
			"(household_id = "+settingHousehold+")"),
		source: "09.2 §3 note of 14 November 2026, §5.1; 10.1 §5.1 note 1.6 item 1; 06 §3.6"},
	// Append-only for app_rw (FR-FAM-020/AC-6; ADR-016 item 5).
	"object_event": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		source: "09.2 §5.1 note of 14 November 2026; 06 §2.5, INV-A01"},
}

type tableRule struct {
	global bool              // no household_id column (09.2 §3 "Tabel global")
	grants map[string]string // role → privileges, comma-separated in the order of allPrivileges; absent role: none
	// updates are the only columns app_rw may UPDATE, as column grants
	// (10.1 §5.1 rule 2; 09.2 §3).
	updates []string
	// policy is the one policy of the table as PostgreSQL deparses it. A
	// global table always declares it; for a household table empty means
	// householdPolicy exactly, and only the tables with lookup clauses of
	// 09.2 §3 (note of 14 November 2026) declare their own.
	policy string
	source string
}

// The settings of 09.2 §3 as PostgreSQL deparses them in a policy: A, E and
// T of the "Tabel global" table (lead ruling Q2 on the 1.4 plan), and the
// household setting read with missing_ok.
const (
	settingAccount   = "(NULLIF(current_setting('app.account_id'::text, true), ''::text))::uuid"
	settingEmail     = "(NULLIF(current_setting('app.email'::text, true), ''::text))::citext"
	settingTokenHash = "decode(NULLIF(current_setting('app.token_hash'::text, true), ''::text), 'hex'::text)" //nolint:gosec // G101: a policy expression naming a setting, no credential
	settingHousehold = "(NULLIF(current_setting('app.household_id'::text, true), ''::text))::uuid"
	// strictHousehold is H of 09.2 §3 layer 4, without missing_ok: a
	// transaction that sets no household fails instead of reading nothing.
	strictHousehold = "(current_setting('app.household_id'::text))::uuid"
)

// The policies of the identity tables of 09.2 §3 "Tabel global": a lookup
// clause only in USING, never in WITH CHECK (10.1 §5.1 rule 1).
var (
	ownAccount     = lookupPolicy("(account_id = "+settingAccount+")", "(account_id = "+settingAccount+")")
	accountOrToken = lookupPolicy("((account_id = "+settingAccount+") OR (token_hash = "+settingTokenHash+"))",
		"(account_id = "+settingAccount+")")
)

func lookupPolicy(using, check string) string {
	return "cmd=ALL roles={public} PERMISSIVE using=" + using + " check=" + check
}

// householdPolicy is the one policy of every household table that declares
// none, as PostgreSQL deparses 09.2 §3 layer 4.
var householdPolicy = lookupPolicy("(household_id = "+strictHousehold+")", "(household_id = "+strictHousehold+")")

// householdColumn is the column that names a household table's household:
// the household table's own key is id (09.2 §3 note of 14 November 2026).
func householdColumn(table string) string {
	if table == "household" {
		return "id"
	}
	return "household_id"
}

// keyExceptions are the unique indexes whose key does not start with
// household_id, each a recorded exception to 09.2 §3 layer 2: both are keyed
// by the caller's own account and hold one membership and one open
// acceptance per account across households (09.2 §5.1, note on decision 20
// of the 1.6 plan). The household table itself is the third exception.
var keyExceptions = []string{"membership_one_household", "invitation_invitee"}

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

	t.Run("SEC-08 a household table has exactly one policy, its declared policy", func(t *testing.T) {
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
			if !hasHouseholdID && householdColumn(name) == "household_id" {
				t.Errorf("household table %s has no household_id column", name)
			}
			want := rule.policy
			if want == "" {
				want = householdPolicy
			}
			if policies := tablePolicies(t, db, name); len(policies) != 1 || policies[0] != want {
				t.Errorf("table %s policies:\n%s\nwant exactly:\n%s (%s)", name, strings.Join(policies, "\n"), want, rule.source)
			}
		}
	})

	t.Run("SEC-08 every global table has exactly its declared policy", func(t *testing.T) {
		for name, rule := range tables {
			if !rule.global {
				continue
			}
			if got := tablePolicies(t, db, name); len(got) != 1 || got[0] != rule.policy {
				t.Errorf("table %s policies:\n%s\nwant exactly:\n%s\n(%s)", name, strings.Join(got, "\n"), rule.policy, rule.source)
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
		// no barrier for a role with them (any role can set the scope
		// settings), so every column grant must be declared in updates.
		var want []string
		for name, rule := range tables {
			for _, column := range rule.updates {
				want = append(want, name+"."+column+" app_rw UPDATE")
			}
		}
		slices.Sort(want)
		got := queryStrings(t, db, `
			SELECT format('%s.%s %s %s', c.relname, a.attname, g.grantee::regrole, g.privilege_type)
			  FROM pg_attribute a JOIN pg_class c ON c.oid = a.attrelid, aclexplode(a.attacl) g
			 WHERE c.relnamespace = 'public'::regnamespace AND c.relname <> 'goose_db_version'
			   AND a.attacl IS NOT NULL
			 ORDER BY 1`)
		slices.Sort(got)
		if !slices.Equal(got, want) {
			t.Errorf("column grants:\n%s\nwant exactly:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
		}
	})

	t.Run("app_rw never updates the owner of an identity row", func(t *testing.T) {
		// 10.1 §5.1 rule 2: a lookup can find a row, but nothing can move a
		// row to another owner; setup_token.account_id is filled once, when
		// the token is used.
		got := queryStrings(t, db, `
			SELECT format('%s.%s', table_name, column_name)
			  FROM information_schema.column_privileges
			 WHERE grantee = 'app_rw' AND privilege_type = 'UPDATE' AND table_schema = 'public'
			   AND column_name IN ('id', 'account_id', 'session_id', 'setup_token_hash', 'scope_id')
			 ORDER BY 1`)
		if want := []string{"setup_token.account_id"}; !slices.Equal(got, want) {
			t.Errorf("app_rw may update the owner columns %v, want only %v", got, want)
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
			// The household table's key is its own id; its currency
			// foreign key names the product table (09.2 §5.1).
			if rule.global || name == "household" {
				continue
			}
			// A foreign key to the household table leads with its id.
			bad := queryStrings(t, db, `
				WITH t AS (SELECT format('public.%I', $1::text)::regclass AS oid)
				SELECT pg_get_constraintdef(con.oid)
				  FROM pg_constraint con, t
				 WHERE con.conrelid = t.oid AND con.contype IN ('p', 'u', 'f')
				   AND ((SELECT attname FROM pg_attribute
				          WHERE attrelid = con.conrelid AND attnum = con.conkey[1]) IS DISTINCT FROM 'household_id'
				        OR (con.contype = 'f' AND (SELECT attname FROM pg_attribute
				              WHERE attrelid = con.confrelid AND attnum = con.confkey[1])
				            IS DISTINCT FROM CASE WHEN con.confrelid = 'public.household'::regclass THEN 'id' ELSE 'household_id' END))
				UNION ALL
				SELECT pg_get_indexdef(i.indexrelid)
				  FROM pg_index i, t
				 WHERE i.indrelid = t.oid AND i.indisunique
				   AND i.indexrelid::regclass::text <> ALL ($2::text[])
				   AND (SELECT attname FROM pg_attribute
				         WHERE attrelid = i.indrelid AND attnum = i.indkey[0]) IS DISTINCT FROM 'household_id'
				 ORDER BY 1`, name, keyExceptions)
			for _, key := range bad {
				t.Errorf("table %s: %s does not start with household_id (SEC-02; 09.2 §3 layer 2)", name, key)
			}
		}
	})
}

// tablePolicies returns the policies of table as PostgreSQL deparses them.
func tablePolicies(t *testing.T, db *pgxpool.Pool, table string) []string {
	t.Helper()
	return queryStrings(t, db, `
		SELECT format('cmd=%s roles=%s %s using=%s check=%s', cmd, roles, permissive, qual, with_check)
		  FROM pg_policies WHERE schemaname = 'public' AND tablename = $1
		 ORDER BY 1`, table)
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
