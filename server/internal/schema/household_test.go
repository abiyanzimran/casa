package schema_test

import (
	"errors"
	"regexp"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

// householdTables are the tables of 00005_household; householdColumn names
// each one's household column.
var householdTables = []string{
	"household", "membership", "approval_authority", "household_rule", "relationship", "invitation", "object_event",
}

// lookupTables are the household tables whose policy has lookup clauses
// (09.2 §3, note of 14 November 2026): without a household they read
// nothing instead of failing.
var lookupTables = []string{"membership", "invitation"}

// householdFixture is two households written as admin: X with members A and
// B, Y with members A and D. Each has an approval authority per member, a
// rule, a relationship and object events. X has an open invitation with
// token hash tx. Y has a cancelled invitation that kept its token hash ty
// (F03 does not clear it) and an accepted one naming account C: invitation_live
// allows one live invitation per household, so Y's token invitation is not
// open.
type householdFixture struct {
	x, y                   pgtype.UUID // households
	a, b, c, d             pgtype.UUID // accounts
	ax, bx, ay, dy         pgtype.UUID // member ids: account, household
	tx, ty                 []byte      // invitation token hashes
	openX, cancelledY, byC pgtype.UUID // invitations
}

// dropOneHouseholdIndex removes membership_one_household from d only, so an
// account can be a member of two households (FR-FAM-020/AC-2, AC-4). The
// index is the pilot's product rule "one household per account", not a
// limit of the model (09.2 §5.1, note on decision 20 of the 1.6 plan).
func dropOneHouseholdIndex(t *testing.T, admin *pgxpool.Pool) {
	t.Helper()
	if _, err := admin.Exec(t.Context(), "DROP INDEX membership_one_household"); err != nil {
		t.Fatal(err)
	}
}

func newHouseholdFixture(t *testing.T, admin *pgxpool.Pool) householdFixture {
	t.Helper()
	f := householdFixture{
		x: randomUUID(), y: randomUUID(),
		a: randomUUID(), b: randomUUID(), c: randomUUID(), d: randomUUID(),
		ax: randomUUID(), bx: randomUUID(), ay: randomUUID(), dy: randomUUID(),
		tx: randomBytes(), ty: randomBytes(),
		openX: randomUUID(), cancelledY: randomUUID(), byC: randomUUID(),
	}
	now := time.Now()
	week := now.Add(7 * 24 * time.Hour)
	for _, s := range []struct {
		sql  string
		args []any
	}{
		{"INSERT INTO household (id, name, time_zone, created_at) VALUES ($1, 'X', 'Asia/Jakarta', $3), ($2, 'Y', 'Asia/Jakarta', $3)",
			[]any{f.x, f.y, now}},
		{`INSERT INTO membership (household_id, member_id, account_id, display_name, role, permission_set, joined_at)
		  VALUES ($1, $3, $5, 'A', 'household_owner', 'two_equal_adults', $9),
		         ($1, $4, $6, 'B', 'household_owner', 'two_equal_adults', $9),
		         ($2, $7, $5, 'A', 'household_owner', 'two_equal_adults', $9),
		         ($2, $8, $10, 'D', 'household_owner', 'two_equal_adults', $9)`,
			[]any{f.x, f.y, f.ax, f.bx, f.a, f.b, f.ay, f.dy, now, f.d}},
		{`INSERT INTO approval_authority (household_id, member_id, request_type, valid_from)
		  VALUES ($1, $3, 'all', $7), ($1, $4, 'all', $7), ($2, $5, 'all', $7), ($2, $6, 'all', $7)`,
			[]any{f.x, f.y, f.ax, f.bx, f.ay, f.dy, now}},
		{`INSERT INTO household_rule (household_id, version, preset_id, created_at)
		  VALUES ($1, 1, 'two_equal_adults', $3), ($2, 1, 'two_equal_adults', $3)`, []any{f.x, f.y, now}},
		{`INSERT INTO relationship (household_id, member_id, other_member_id, kind)
		  VALUES ($1, $3, $4, 'partner'), ($2, $5, $6, 'partner')`, []any{f.x, f.y, f.ax, f.bx, f.ay, f.dy}},
		{`INSERT INTO invitation (household_id, id, inviter, token_hash, state, created_at, expires_at, closed_at,
		                         invitee_account_id, invitee_display_name, version)
		  VALUES ($1, $3, $6, $7, 'open', $10, $11, NULL, NULL, NULL, 0),
		         ($2, $4, $8, $9, 'cancelled', $10, $11, $10, NULL, NULL, 1),
		         ($2, $5, $8, NULL, 'accepted', $10, $11, NULL, $12, 'C', 1)`,
			[]any{f.x, f.y, f.openX, f.cancelledY, f.byC, f.ax, f.tx, f.ay, f.ty, now, week, f.c}},
		{`INSERT INTO object_event (household_id, object_kind, object_id, seq, actor_kind, actor, command, to_state, server_time)
		  VALUES ($1, 'household', $1, 0, 'member', $3, 'CreateHousehold', 'created', $5),
		         ($2, 'household', $2, 0, 'member', $4, 'CreateHousehold', 'created', $5)`,
			[]any{f.x, f.y, f.ax, f.ay, now}},
		{`INSERT INTO object_event (household_id, object_kind, object_id, seq, actor_kind, command, from_state, to_state,
		                           idempotency_key, server_time)
		  VALUES ($1, 'invitation', $2, 1, 'invitee', 'AcceptInvitation', 'open', 'accepted', gen_random_uuid(), $3)`,
			[]any{f.y, f.byC, now}},
	} {
		if _, err := admin.Exec(t.Context(), s.sql, s.args...); err != nil {
			t.Fatalf("%s: %v", s.sql, err)
		}
	}
	return f
}

// errRolledBack ends a tryWrite transaction without committing.
var errRolledBack = errors.New("rolled back")

// tryWrite runs statement as app_rw in a database.Write with scope and rolls
// it back, so a statement that succeeds leaves the fixture unchanged. It
// returns the statement's error.
func tryWrite(t *testing.T, app *pgxpool.Pool, scope database.Scope, statement string, args ...any) error {
	t.Helper()
	err := database.Write(t.Context(), app, scope, func(tx pgx.Tx) error {
		if _, err := tx.Exec(t.Context(), statement, args...); err != nil {
			return err
		}
		return errRolledBack
	})
	if errors.Is(err, errRolledBack) {
		return nil
	}
	return err
}

// TestHouseholdTablesAreHouseholdScoped proves layer 4 of 09.2 §3 for the
// household tables and the lookup clauses of 10.1 §5.1 (note of 14 November
// 2026). Fixtures as admin, after dropOneHouseholdIndex (decision 2):
// households X and Y, account A member of both, account B member of X only;
// each with an approval authority, rule, invitation with a token hash (tx,
// ty), an accepted invitation naming account C in Y, and object events.
func TestHouseholdTablesAreHouseholdScoped(t *testing.T) {
	d := databasetest.Migrated(t)
	admin := d.Pool(t, d.AdminRole)
	app := d.Pool(t, "app_rw")
	dropOneHouseholdIndex(t, admin)
	f := newHouseholdFixture(t, admin)
	now := time.Now()

	t.Run("a household scope reads only its own rows", func(t *testing.T) {
		for _, table := range householdTables {
			got := readColumn[pgtype.UUID](t, app, database.Scope{HouseholdID: f.y}, table, householdColumn(table))
			if len(got) == 0 || slices.ContainsFunc(got, func(u pgtype.UUID) bool { return u != f.y }) {
				t.Errorf("%s: household Y reads households %v, want only Y %v", table, got, f.y)
			}
		}
	})

	t.Run("a household scope never reads another household through a lookup clause", func(t *testing.T) {
		// SEC-38 (b): a lookup clause lives only where no household is set.
		for _, c := range []struct {
			scope database.Scope
			table string
		}{
			{database.Scope{HouseholdID: f.x, AccountID: f.a}, "membership"},
			{database.Scope{HouseholdID: f.x, TokenHash: f.ty}, "invitation"},
			{database.Scope{HouseholdID: f.x, AccountID: f.c}, "invitation"},
		} {
			got := readColumn[pgtype.UUID](t, app, c.scope, c.table, "household_id")
			if len(got) == 0 || slices.Contains(got, f.y) {
				t.Errorf("%s with scope %+v reads households %v, want only X %v", c.table, c.scope, got, f.x)
			}
		}
	})

	t.Run("an account lookup finds its own memberships in every household and nothing else", func(t *testing.T) {
		got := readColumn[pgtype.UUID](t, app, database.Scope{AccountID: f.a}, "membership", "member_id")
		if len(got) != 2 || !slices.Contains(got, f.ax) || !slices.Contains(got, f.ay) {
			t.Errorf("account A finds members %v, want exactly A's in X and Y %v, %v", got, f.ax, f.ay)
		}
	})

	t.Run("a token lookup finds only its invitation", func(t *testing.T) {
		got := readColumn[pgtype.UUID](t, app, database.Scope{TokenHash: f.tx}, "invitation", "id")
		if !slices.Equal(got, []pgtype.UUID{f.openX}) {
			t.Errorf("token tx finds invitations %v, want only %v", got, f.openX)
		}
	})

	t.Run("an invitee lookup sees only invitations naming that account", func(t *testing.T) {
		// Plan 1.6 decision 3: whatever the state; the F05 query filters
		// state = 'accepted'.
		if got := readColumn[pgtype.UUID](t, app, database.Scope{AccountID: f.c}, "invitation", "id"); !slices.Equal(got, []pgtype.UUID{f.byC}) {
			t.Errorf("account C finds invitations %v, want only %v", got, f.byC)
		}
		if got := readColumn[pgtype.UUID](t, app, database.Scope{AccountID: f.b}, "invitation", "id"); len(got) != 0 {
			t.Errorf("account B, named by no invitation, finds %v", got)
		}
	})

	t.Run("a lookup scope cannot write", func(t *testing.T) {
		// 10.1 §5.1 note 1.6 item 1: a lookup clause is only in USING.
		// membership has no UPDATE grant at all (catalog privileges subtest).
		insertMembership := `INSERT INTO membership (household_id, account_id, role, permission_set, joined_at)
		                     VALUES ($1, $2, 'household_owner', 'two_equal_adults', now())`
		insertInvitation := `INSERT INTO invitation (household_id, inviter, state, created_at, expires_at)
		                     VALUES ($1, $2, 'cancelled', now(), now())`
		for _, scope := range []database.Scope{{AccountID: f.a}, {TokenHash: f.tx}} {
			for _, w := range []struct {
				statement string
				args      []any
			}{
				{insertMembership, []any{f.x, randomUUID()}},
				{insertInvitation, []any{f.x, f.ax}},
			} {
				if err := tryWrite(t, app, scope, w.statement, w.args...); databasetest.PgErrorCode(err) != "42501" {
					t.Errorf("scope %+v: %s: err = %v, want 42501", scope, w.statement, err)
				}
			}
		}
		for _, c := range []struct {
			scope     database.Scope
			statement string
			arg       any
		}{
			{database.Scope{TokenHash: f.tx}, "UPDATE invitation SET state = 'cancelled' WHERE token_hash = $1", f.tx},
			{database.Scope{AccountID: f.c}, "UPDATE invitation SET state = 'withdrawn' WHERE invitee_account_id = $1", f.c},
		} {
			if err := tryWrite(t, app, c.scope, c.statement, c.arg); databasetest.PgErrorCode(err) != "42501" {
				t.Errorf("%s: err = %v, want 42501 (WITH CHECK)", c.statement, err)
			}
		}
	})

	t.Run("a household scope cannot write rows of another household", func(t *testing.T) {
		for _, w := range []struct {
			statement string
			args      []any
		}{
			{"INSERT INTO household (id, name, time_zone, created_at) VALUES ($1, 'Z', 'Asia/Jakarta', $2)",
				[]any{randomUUID(), now}},
			{`INSERT INTO membership (household_id, account_id, role, permission_set, joined_at)
			  VALUES ($1, $2, 'household_owner', 'two_equal_adults', $3)`, []any{f.x, randomUUID(), now}},
			{"INSERT INTO approval_authority (household_id, member_id, request_type, valid_from) VALUES ($1, $2, 'all', $3)",
				[]any{f.x, f.ax, now.Add(time.Hour)}},
			{"INSERT INTO household_rule (household_id, version, preset_id, created_at) VALUES ($1, 2, 'two_equal_adults', $2)",
				[]any{f.x, now}},
			{"INSERT INTO invitation (household_id, inviter, state, created_at, expires_at) VALUES ($1, $2, 'cancelled', $3, $3)",
				[]any{f.x, f.ax, now}},
			{`INSERT INTO object_event (household_id, object_kind, object_id, seq, actor_kind, actor, command, to_state, server_time)
			  VALUES ($1, 'household', $1, 1, 'member', $2, 'x', 'x', $3)`, []any{f.x, f.ax, now}},
		} {
			if err := tryWrite(t, app, database.Scope{HouseholdID: f.y}, w.statement, w.args...); databasetest.PgErrorCode(err) != "42501" {
				t.Errorf("as Y: %s: err = %v, want 42501", w.statement, err)
			}
		}
	})

	t.Run("without a scope a lookup table is empty and a strict table fails", func(t *testing.T) {
		for _, table := range householdTables {
			var rows int
			err := database.Read(t.Context(), app, database.Scope{}, func(tx pgx.Tx) error {
				return tx.QueryRow(t.Context(), "SELECT count(*) FROM "+table).Scan(&rows)
			})
			if slices.Contains(lookupTables, table) {
				if err != nil || rows != 0 {
					t.Errorf("%s without a scope: %d rows, err %v; want 0, nil", table, rows, err)
				}
				continue
			}
			// 09.2 §3 layer 4 has no missing_ok: '' is not a uuid.
			if code := databasetest.PgErrorCode(err); code != "22P02" {
				t.Errorf("%s without a scope: err = %v, want invalid_text_representation (22P02)", table, err)
			}
		}
	})

	t.Run("FR-FAM-020/AC-6 object_event is append-only for app_rw", func(t *testing.T) {
		for _, statement := range []string{"UPDATE object_event SET command = 'x'", "DELETE FROM object_event"} {
			if err := tryWrite(t, app, database.Scope{HouseholdID: f.x}, statement); databasetest.PgErrorCode(err) != "42501" {
				t.Errorf("%s: err = %v, want insufficient_privilege (42501)", statement, err)
			}
		}
	})

	t.Run("relationship is read-only for app_rw", func(t *testing.T) {
		// 06 §3.1 row "Hubungan": no writer until Tier 3.
		err := tryWrite(t, app, database.Scope{HouseholdID: f.x},
			"INSERT INTO relationship (household_id, member_id, other_member_id, kind) VALUES ($1, $3, $2, 'partner')", f.x, f.ax, f.bx)
		if databasetest.PgErrorCode(err) != "42501" {
			t.Errorf("insert into relationship: err = %v, want 42501", err)
		}
	})

	t.Run("object_event names its actor exactly as its kind says", func(t *testing.T) {
		// 09.2 §5.1: member events name a member of the same household,
		// system events their rule, invitee events only an invitation.
		insert := `INSERT INTO object_event (household_id, object_kind, object_id, seq, actor_kind, actor, rule, command, to_state, server_time)
		           VALUES ($1, $2, $3, 7, $4, $5, $6, 'x', 'x', now())`
		var none pgtype.UUID
		for _, c := range []struct {
			name, objectKind, actorKind string
			actor                       pgtype.UUID
			rule                        any
			want                        string
		}{
			{"member without actor", "household", "member", none, nil, "23514"},
			{"member with a rule", "household", "member", f.ax, "invitation.expired", "23514"},
			{"system without rule", "invitation", "system", none, nil, "23514"},
			{"system with an actor", "invitation", "system", f.ax, "invitation.expired", "23514"},
			{"invitee with an actor", "invitation", "invitee", f.ax, nil, "23514"},
			{"invitee on a household event", "household", "invitee", none, nil, "23514"},
			{"actor of another household", "household", "member", f.ay, nil, "23503"},
			{"member of this household", "household", "member", f.ax, nil, ""},
			{"system with its rule", "invitation", "system", none, "invitation.expired", ""},
			{"invitee on an invitation", "invitation", "invitee", none, nil, ""},
		} {
			err := tryWrite(t, app, database.Scope{HouseholdID: f.x}, insert, f.x, c.objectKind, f.openX, c.actorKind, c.actor, c.rule)
			if got := databasetest.PgErrorCode(err); got != c.want || (c.want == "" && err != nil) {
				t.Errorf("%s: err = %v, want SQLSTATE %q", c.name, err, c.want)
			}
		}
	})

	t.Run("an amount and its currency are both set or both empty", func(t *testing.T) {
		// 09.2 §5.1, §7: an amount always carries its currency.
		for _, w := range []struct {
			statement string
			args      []any
		}{
			{`INSERT INTO approval_authority (household_id, member_id, request_type, limit_amount, currency, valid_from)
			  VALUES ($1, $2, 'all', $3, $4, now() + interval '1 day')`, []any{f.x, f.ax}},
			{`INSERT INTO household_rule (household_id, version, preset_id, request_threshold, currency, created_at)
			  VALUES ($1, 2, 'two_equal_adults', $2, $3, now())`, []any{f.x}},
		} {
			for _, c := range []struct {
				amount   any
				currency any
				want     string
			}{
				{int64(100000), nil, "23514"},
				{nil, "IDR", "23514"},
				{int64(100000), "IDR", ""},
				{nil, nil, ""},
			} {
				err := tryWrite(t, app, database.Scope{HouseholdID: f.x}, w.statement, append(w.args, c.amount, c.currency)...)
				if got := databasetest.PgErrorCode(err); got != c.want || (c.want == "" && err != nil) {
					t.Errorf("%s with amount %v, currency %v: err = %v, want SQLSTATE %q", w.statement, c.amount, c.currency, err, c.want)
				}
			}
		}
	})

	t.Run("FR-FAM-020/AC-3 no household table has a column meaning husband or wife", func(t *testing.T) {
		// PRN-03; 06 §14: neither a name nor an allowed value.
		spouse := regexp.MustCompile(`(?i)husband|wife|suami|istri`)
		names := queryStrings(t, admin, `
			SELECT table_name || '.' || column_name FROM information_schema.columns
			 WHERE table_schema = 'public' AND table_name = ANY ($1)
			UNION ALL
			SELECT pg_get_constraintdef(oid) FROM pg_constraint
			 WHERE conrelid = ANY (SELECT format('public.%I', t)::regclass FROM unnest($1::text[]) t)`, householdTables)
		if len(names) == 0 {
			t.Fatal("no columns found")
		}
		for _, name := range names {
			if spouse.MatchString(name) {
				t.Errorf("%s means husband or wife", name)
			}
		}
	})

	t.Run("no member data column for relationship, birth date, gender or income", func(t *testing.T) {
		// FR-FAM-006/AC-3, database part: the relationship table itself is
		// the only place for a relationship. household.income_source_market
		// is the household's market for the default income sources (07 §3.1
		// "Daftar bawaan per pasar"), not a member's income.
		forbidden := regexp.MustCompile(`(?i)birth|gender|sex|income|relationship`)
		columns := queryStrings(t, admin, `
			SELECT table_name || '.' || column_name FROM information_schema.columns
			 WHERE table_schema = 'public' AND table_name = ANY ($1)`, householdTables)
		for _, column := range columns {
			_, name, _ := strings.Cut(column, ".")
			if forbidden.MatchString(name) && column != "household.income_source_market" {
				t.Errorf("column %s holds member data FR-FAM-006/AC-3 forbids", column)
			}
		}
	})

	t.Run("only one invitation of a household is live", func(t *testing.T) {
		// 06 §3.6 F01: one open or accepted invitation per household.
		insert := `INSERT INTO invitation (household_id, inviter, state, created_at, expires_at, invitee_account_id)
		           VALUES ($1, $2, $3, now(), now() + interval '7 days', $4)`
		for state, want := range map[string]string{"open": "23505", "accepted": "23505", "cancelled": ""} {
			var invitee any
			if state == "accepted" {
				invitee = randomUUID()
			}
			err := tryWrite(t, app, database.Scope{HouseholdID: f.x}, insert, f.x, f.bx, state, invitee)
			if got := databasetest.PgErrorCode(err); got != want || (want == "" && err != nil) {
				t.Errorf("a second invitation in state %s: err = %v, want SQLSTATE %q", state, err, want)
			}
		}
	})
}

// TestOneHouseholdPerAccount runs on a fully migrated database (indexes
// kept): the two unique indexes of 09.2 §5.1 (note on decision 20 of the 1.6
// plan) hold across households, whatever row-level security shows.
func TestOneHouseholdPerAccount(t *testing.T) {
	d := databasetest.Migrated(t)
	admin := d.Pool(t, d.AdminRole)
	app := d.Pool(t, "app_rw")
	p, q := randomUUID(), randomUUID()
	a, b, e := randomUUID(), randomUUID(), randomUUID()
	ap, bq := randomUUID(), randomUUID()
	for _, s := range []struct {
		sql  string
		args []any
	}{
		{"INSERT INTO household (id, name, time_zone, created_at) VALUES ($1, 'P', 'Asia/Jakarta', now()), ($2, 'Q', 'Asia/Jakarta', now())",
			[]any{p, q}},
		{`INSERT INTO membership (household_id, member_id, account_id, role, permission_set, joined_at)
		  VALUES ($1, $2, $3, 'household_owner', 'two_equal_adults', now()), ($4, $5, $6, 'household_owner', 'two_equal_adults', now())`,
			[]any{p, ap, a, q, bq, b}},
		{`INSERT INTO invitation (household_id, inviter, state, created_at, expires_at, invitee_account_id)
		  VALUES ($1, $2, 'accepted', now(), now() + interval '7 days', $3)`, []any{p, ap, e}},
	} {
		if _, err := admin.Exec(t.Context(), s.sql, s.args...); err != nil {
			t.Fatalf("%s: %v", s.sql, err)
		}
	}

	t.Run("an account has one membership across households", func(t *testing.T) {
		err := tryWrite(t, app, database.Scope{HouseholdID: q}, `
			INSERT INTO membership (household_id, account_id, role, permission_set, joined_at)
			VALUES ($1, $2, 'household_owner', 'two_equal_adults', now())`, q, a)
		if code := databasetest.PgErrorCode(err); code != "23505" {
			t.Errorf("A's membership in a second household: err = %v, want unique_violation (23505)", err)
		}
	})

	t.Run("an account has one open acceptance across households", func(t *testing.T) {
		err := tryWrite(t, app, database.Scope{HouseholdID: q}, `
			INSERT INTO invitation (household_id, inviter, state, created_at, expires_at, invitee_account_id)
			VALUES ($1, $2, 'accepted', now(), now() + interval '7 days', $3)`, q, bq, e)
		if code := databasetest.PgErrorCode(err); code != "23505" {
			t.Errorf("a second accepted invitation naming one account: err = %v, want unique_violation (23505)", err)
		}
	})
}
