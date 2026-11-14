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
		policies: []string{"cmd=SELECT roles={public} PERMISSIVE using=true check="},
		source:   "09.2 §3 \"Tabel global\", §5.1; ADR-008"},
	"account": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policies: []string{policyText("((id = "+settingAccount+") OR (email = "+settingEmail+"))", "(id = "+settingAccount+")")},
		source:   "09.2 §3 \"Tabel global\", §5.1; lead rulings Q2, Q3 on the 1.4 plan"},
	"credential": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates:  []string{"salt", "hash", "time_cost", "memory_kib", "parallelism", "updated_at"},
		policies: []string{ownAccount},
		source:   "09.2 §3 \"Tabel global\"; SEC-11; lead ruling Q3 on the 1.4 plan"},
	"session": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		// 00004_web_session: the cookie columns, and no UPDATE on device_mode
		// or ua_summary (SEC-165).
		updates: []string{"token_hash", "token_issued_at", "last_used_at", "revoked_at", "revoke_reason",
			"prev_token_hash", "prev_valid_until", "token_confirmed_at", "discarded_token_hash"},
		policies: []string{policyText("((account_id = "+settingAccount+") OR (token_hash = "+settingTokenHash+
			") OR (prev_token_hash = "+settingTokenHash+") OR (discarded_token_hash = "+settingTokenHash+"))",
			"(account_id = "+settingAccount+")")},
		source: "09.2 §3 \"Tabel global\", §5.1; SEC-15, SEC-161 (c); lead rulings Q2, Q3 on the 1.4 plan; plan 1.5 decision 15"},
	"refresh_token": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates:  []string{"used_at"},
		policies: []string{accountOrToken},
		source:   "09.2 §3 \"Tabel global\"; SEC-15; lead rulings Q2, Q3 on the 1.4 plan"},
	"reauth_grant": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates:  []string{"used_at"},
		policies: []string{ownAccount},
		source:   "09.2 §3 \"Tabel global\"; SEC-13; SEC-161 (c) note 1.5 rule (7) (cookie_hash, 00004); lead ruling Q3 on the 1.4 plan"},
	"recovery_token": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates:  []string{"used_at"},
		policies: []string{accountOrToken},
		source:   "09.2 §3 \"Tabel global\"; SEC-24; lead rulings Q2, Q3 on the 1.4 plan"},
	"setup_token": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates:  []string{"used_at", "account_id"},
		policies: []string{policyText("(token_hash = "+settingTokenHash+")", "(token_hash = "+settingTokenHash+")")},
		source:   "09.2 §3 \"Tabel global\", §5.1; SEC-158 (b); 09.4 §2; lead rulings Q2, Q3 on the 1.4 plan"},
	// 00008_invitation_signup: the invitation source; email is replaced on
	// a restart before an account is born, account_id is set once (10.1
	// §5.1 note 1.6 item 4; security-privacy ruling on Task 8 of plan 1.6).
	"signup_verification": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"token_hash", "email", "expires_at", "used_at", "account_id"},
		policies: []string{policyText("((token_hash = "+settingTokenHash+") OR (setup_token_hash = "+settingTokenHash+
			") OR (invitation_token_hash = "+settingTokenHash+") OR (account_id = "+settingAccount+"))",
			"((setup_token_hash = "+settingTokenHash+") OR (invitation_token_hash = "+settingTokenHash+"))")},
		source: "09.2 §3 \"Tabel global\", §5.1; SEC-158; lead rulings Q2, Q4, Q8 on the 1.4 plan; plan 1.6 decision 15; 10.1 §5.1 note 1.6 item 4"},
	"security_event": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policies: []string{ownAccount},
		source:   "09.2 §3 \"Tabel global\"; SEC-16, SEC-39 (b): append-only; lead ruling Q3 on the 1.4 plan"},
	"idempotency_record": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policies: []string{"cmd=ALL roles={public} PERMISSIVE using=((scope_id = " + settingHousehold + ") OR (scope_id = " + settingAccount + ")) check="},
		source:   "09.2 §3 \"Tabel infrastruktur dengan pemilik campuran\", §5.1, §8"},
	// 00005_household: the household tables (09.2 §5.1, note of 14 November
	// 2026; plan 1.6 decisions 2, 3). app_rw never deletes a household row.
	"household": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates:  []string{"version"},
		policies: []string{policyText("(id = "+strictHousehold+")", "(id = "+strictHousehold+")")},
		source:   "09.2 §3 note of 14 November 2026, §5.1; plan 1.6 decisions 2, 3"},
	"membership": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policies: []string{policyText("((household_id = "+settingHousehold+") OR (("+settingHousehold+" IS NULL) AND (account_id = "+settingAccount+")))",
			"(household_id = "+settingHousehold+")")},
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
		policies: []string{policyText("((household_id = "+settingHousehold+") OR (("+settingHousehold+" IS NULL) AND (token_hash = "+settingTokenHash+
			")) OR (("+settingHousehold+" IS NULL) AND (invitee_account_id = "+settingAccount+")))",
			"(household_id = "+settingHousehold+")")},
		source: "09.2 §3 note of 14 November 2026, §5.1; 10.1 §5.1 note 1.6 item 1; 06 §3.6"},
	// Append-only for app_rw (FR-FAM-020/AC-6; ADR-016 item 5).
	"object_event": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		source: "09.2 §5.1 note of 14 November 2026; 06 §2.5, INV-A01"},
	// 00006_consent_record: a global table of household (plan 1.6 decision 9).
	"consent_record": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policies: []string{ownAccount},
		source:   "09.2 §3 note of 14 November 2026, §5.1; SEC-80 note of 14 November 2026, SEC-109 (a)"},
	// 00007_operator_notice_ack: a global table of household (plan 1.6
	// decision 6); history stays, only the acknowledgement is filled in.
	"operator_notice_ack": {global: true, grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates:  []string{"acknowledged_at"},
		policies: []string{ownAccount},
		source:   "09.2 §3 note of 14 November 2026, §5.1; SEC-34; 10.1 §11 note of 14 November 2026"},
	// 00009_job: a table with a mixed owner (CS-A-20). app_rw enqueues in
	// the domain transaction and deletes a job that ran; job_runner claims
	// across households and changes only the retry columns. Both policies
	// name their role: permissive policies are ORed, so USING (true)
	// without TO would show app_rw every household's jobs (09.2 §3 note of
	// 14 November 2026, Q1 (b); plan 1.7 decision 3).
	"job": {grants: map[string]string{"app_rw": "SELECT,INSERT,DELETE", "job_runner": "SELECT"},
		runnerUpdates: []string{"locked_until", "attempts", "last_error"},
		policies: []string{
			rolePolicyText("app_rw", "((household_id = "+settingHousehold+") OR ((household_id IS NULL) AND ("+settingHousehold+" IS NULL)))",
				"((household_id = "+settingHousehold+") OR ((household_id IS NULL) AND ("+settingHousehold+" IS NULL)))"),
			rolePolicyText("job_runner", "true", ""),
		},
		source: "09.2 §3 \"Tabel infrastruktur dengan pemilik campuran\" and its note of 14 November 2026, §5.1; plan 1.7 decision 3"},
	// 00010_reader_feed: the per-reader feed of sync (ADR-017 item 3),
	// append-only for app_rw. Its layer 4 policy names app_rw from the
	// start, so the pruning policy of job_runner (step 7) is not ANDed with
	// a household setting it never sets (09.2 §3 note of 14 November 2026,
	// Q1 (d)).
	"reader_feed": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		policies: []string{rolePolicyText("app_rw", "(household_id = "+strictHousehold+")", "(household_id = "+strictHousehold+")")},
		source:   "09.2 §3 note of 14 November 2026 (Q1 (d)), §5.1; ADR-017 item 3; plan 1.7 Task 4"},
	// The counter moves only forward, by the fan-out (plan 1.7 decision 8).
	"reader_feed_counter": {grants: map[string]string{"app_rw": "SELECT,INSERT"},
		updates: []string{"last_pos"},
		source:  "09.2 §5.1; ADR-017 item 3; plan 1.7 Task 4"},
}

type tableRule struct {
	global bool              // no household_id column (09.2 §3 "Tabel global")
	grants map[string]string // role → privileges, comma-separated in the order of allPrivileges; absent role: none
	// updates are the only columns app_rw may UPDATE, as column grants
	// (10.1 §5.1 rule 2; 09.2 §3).
	updates []string
	// runnerUpdates are the only columns job_runner may UPDATE (09.2 §3 row
	// "job (pengambilan)").
	runnerUpdates []string
	// policies are the table's policies as PostgreSQL deparses them, each
	// with the roles it applies to, in the order tablePolicies returns
	// them. A global table always declares them; for a household table
	// empty means householdPolicy alone, and only the tables with lookup
	// clauses or a second role of 09.2 §3 (notes of 14 November 2026)
	// declare their own.
	policies []string
	source   string
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
	ownAccount     = policyText("(account_id = "+settingAccount+")", "(account_id = "+settingAccount+")")
	accountOrToken = policyText("((account_id = "+settingAccount+") OR (token_hash = "+settingTokenHash+"))",
		"(account_id = "+settingAccount+")")
)

// policyText is one FOR ALL policy for every role as tablePolicies prints it.
func policyText(using, check string) string {
	return rolePolicyText("public", using, check)
}

// rolePolicyText is one FOR ALL policy TO role as tablePolicies prints it.
func rolePolicyText(role, using, check string) string {
	return "cmd=ALL roles={" + role + "} PERMISSIVE using=" + using + " check=" + check
}

// householdPolicy is the one policy of every household table that declares
// none, as PostgreSQL deparses 09.2 §3 layer 4.
var householdPolicy = policyText("(household_id = "+strictHousehold+")", "(household_id = "+strictHousehold+")")

// householdColumn is the column that names a household table's household:
// the household table's own key is id (09.2 §3 note of 14 November 2026).
func householdColumn(table string) string {
	if table == "household" {
		return "id"
	}
	return "household_id"
}

// keyExceptions are the unique indexes whose key does not start with
// household_id, each a recorded exception to 09.2 §3 layer 2, with its exact
// definition. The first two are keyed by the caller's own account and hold
// one membership and one open acceptance per account across households (09.2
// §5.1, note on decision 20 of the 1.6 plan). job_pkey is job's primary key
// id of 09.2 §5.1: household_id is NULL for a system job, so it cannot lead
// a primary key; job_unique leads with it (09.2 §5.1 note of 14 November
// 2026). The household table itself is the fourth exception.
var keyExceptions = map[string]string{
	"membership_one_household": "CREATE UNIQUE INDEX membership_one_household ON public.membership USING btree (account_id) WHERE (account_id IS NOT NULL)",
	"invitation_invitee":       "CREATE UNIQUE INDEX invitation_invitee ON public.invitation USING btree (invitee_account_id) WHERE (state = 'accepted'::text)",
	"job_pkey":                 "CREATE UNIQUE INDEX job_pkey ON public.job USING btree (id)",
}

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

	t.Run("SEC-08 a household table has exactly its declared policies", func(t *testing.T) {
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
			want := rule.policies
			if want == nil {
				want = []string{householdPolicy}
			}
			if policies := tablePolicies(t, db, name); !slices.Equal(policies, want) {
				t.Errorf("table %s policies:\n%s\nwant exactly:\n%s (%s)", name, strings.Join(policies, "\n"), strings.Join(want, "\n"), rule.source)
			}
		}
	})

	t.Run("SEC-08 every global table has exactly its declared policy", func(t *testing.T) {
		for name, rule := range tables {
			if !rule.global {
				continue
			}
			if got := tablePolicies(t, db, name); !slices.Equal(got, rule.policies) {
				t.Errorf("table %s policies:\n%s\nwant exactly:\n%s\n(%s)", name, strings.Join(got, "\n"), strings.Join(rule.policies, "\n"), rule.source)
			}
		}
	})

	t.Run("each policy of job names its role", func(t *testing.T) {
		// 09.2 §3 note of 14 November 2026 (Q1 (b)): a policy without TO
		// applies to every role, and permissive policies are ORed, so the
		// claim policy's USING (true) would show app_rw every household's
		// jobs.
		got := queryStrings(t, db, `
			SELECT roles::text FROM pg_policies
			 WHERE schemaname = 'public' AND tablename = 'job'
			 ORDER BY 1`)
		if want := []string{"{app_rw}", "{job_runner}"}; !slices.Equal(got, want) {
			t.Errorf("roles of the job policies: %v, want exactly %v (a policy without TO is {public})", got, want)
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
			for _, column := range rule.runnerUpdates {
				want = append(want, name+"."+column+" job_runner UPDATE")
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
		// row to another owner; setup_token.account_id and
		// signup_verification.account_id are filled once, when the account
		// is born (note 1.6 item 4), and a sign-up's source token hash never
		// changes.
		got := queryStrings(t, db, `
			SELECT format('%s.%s', table_name, column_name)
			  FROM information_schema.column_privileges
			 WHERE grantee = 'app_rw' AND privilege_type = 'UPDATE' AND table_schema = 'public'
			   AND column_name IN ('id', 'account_id', 'session_id', 'setup_token_hash', 'invitation_token_hash', 'scope_id')
			 ORDER BY 1`)
		if want := []string{"setup_token.account_id", "signup_verification.account_id"}; !slices.Equal(got, want) {
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

	t.Run("a household's currency cannot change under an amount (INV-35)", func(t *testing.T) {
		// 09.2 §5.2 item 4; ADR-016: every (household_id, currency) foreign
		// key to household is ON UPDATE RESTRICT.
		got := queryStrings(t, db, `
			SELECT format('%s %s', con.conrelid::regclass, pg_get_constraintdef(con.oid))
			  FROM pg_constraint con
			 WHERE con.contype = 'f' AND con.confrelid = 'public.household'::regclass
			   AND (SELECT attname FROM pg_attribute WHERE attrelid = con.conrelid AND attnum = con.conkey[2]) = 'currency'
			   AND con.confupdtype <> 'r'
			 ORDER BY 1`)
		for _, key := range got {
			t.Errorf("%s is not ON UPDATE RESTRICT", key)
		}
		if n := queryStrings(t, db, `
			SELECT conname::text FROM pg_constraint
			 WHERE contype = 'f' AND confrelid = 'public.household'::regclass AND cardinality(conkey) = 2`); len(n) == 0 {
			t.Error("no currency foreign key to household found")
		}
	})

	t.Run("household_id leads every key of a household table (SEC-02)", func(t *testing.T) {
		// An exception exists exactly as declared, so it cannot widen
		// unnoticed.
		for name, want := range keyExceptions {
			got := queryStrings(t, db, "SELECT coalesce(pg_get_indexdef(to_regclass(format('public.%I', $1::text))), 'missing')", name)
			if len(got) != 1 || got[0] != want {
				t.Errorf("key exception %s: %v, want exactly %q", name, got, want)
			}
		}
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
				   AND con.conname <> ALL ($2::text[])
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
				 ORDER BY 1`, name, slices.Collect(maps.Keys(keyExceptions)))
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
