// Package isolation holds only tests: the generated cross-household
// isolation test of every sqlc query (09.2 §3 layer 6; 09.4 §4 row
// "Isolasi"; SEC-02; ES §4; plan 1.7 decision 15).
//
// Every query is found in the generated code, so a new query is tested
// without a line here unless one of its parameters has no isolation value,
// which fails the test. Each query runs as household Y's member: once with
// Y's values, which must succeed; then, for every identifier parameter, with
// household X's value and with a fresh random value, every other parameter
// Y's; and, with two identifiers or more, with all of them X's and all of
// them fresh. Each X run must answer like its fresh run (SQLSTATE, row
// count, and every boolean and integer in the result), no run may return an
// X value, and no row X owns may change. Another household's value then
// behaves exactly like one that never existed.
package isolation

import (
	"bytes"
	"context"
	"crypto/rand"
	_ "embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"maps"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/abiyanzimran/casa/server/internal/command/commanddb"
	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
	"github.com/abiyanzimran/casa/server/internal/household/householddb"
	"github.com/abiyanzimran/casa/server/internal/identity/identitydb"
	"github.com/abiyanzimran/casa/server/internal/jobs/jobsdb"
	"github.com/abiyanzimran/casa/server/internal/read/readdb"
	"github.com/abiyanzimran/casa/server/internal/sync/syncdb"
)

// fixture writes households X and Y as the superuser: a row of each in every
// table of public but currency, with the identifiers below.
//
//go:embed testdata/fixture.sql
var fixture string

// Fixture identifiers (testdata/fixture.sql): X's begin with 1, Y's with 2.
var (
	householdX, householdY           = uuid("10000000-0000-4000-8000-000000000001"), uuid("20000000-0000-4000-8000-000000000001")
	accountX, accountY               = uuid("10000000-0000-4000-8000-000000000011"), uuid("20000000-0000-4000-8000-000000000011")
	inviteeX, inviteeY               = uuid("10000000-0000-4000-8000-000000000012"), uuid("20000000-0000-4000-8000-000000000012")
	memberX, memberY                 = uuid("10000000-0000-4000-8000-000000000021"), uuid("20000000-0000-4000-8000-000000000021")
	sessionX, sessionY               = uuid("10000000-0000-4000-8000-000000000031"), uuid("20000000-0000-4000-8000-000000000031")
	invitationX, invitationY         = uuid("10000000-0000-4000-8000-000000000071"), uuid("20000000-0000-4000-8000-000000000071")
	idempotencyKeyX, idempotencyKeyY = uuid("10000000-0000-4000-8000-000000000081"), uuid("20000000-0000-4000-8000-000000000081")
	jobX, jobY                       = uuid("10000000-0000-4000-8000-000000000091"), uuid("20000000-0000-4000-8000-000000000091")
	emailX, emailY                   = "owner-x@example.test", "owner-y@example.test"

	sessionTokenX, sessionTokenY       = hash(0x11), hash(0x21)
	refreshTokenX, refreshTokenY       = hash(0x12), hash(0x22)
	recoveryTokenX, recoveryTokenY     = hash(0x13), hash(0x23)
	setupTokenX, setupTokenY           = hash(0x14), hash(0x24)
	signupLinkX, signupLinkY           = hash(0x15), hash(0x25)
	invitationTokenX, invitationTokenY = hash(0x16), hash(0x26)
	cookieHashX, cookieHashY           = hash(0x17), hash(0x27)
)

// otherValuesX are X's fixture values no parameter carries: a result holding
// one of them has read X's data too.
var otherValuesX = []any{
	uuid("10000000-0000-4000-8000-000000000041"), // reauth_grant
	uuid("10000000-0000-4000-8000-000000000051"), // signup_verification (setup token)
	uuid("10000000-0000-4000-8000-000000000052"), // signup_verification (invitation)
	uuid("10000000-0000-4000-8000-000000000061"), // security_event
	uuid("10000000-0000-4000-8000-000000000002"), // household X2
	uuid("10000000-0000-4000-8000-000000000022"), // X2's member
	uuid("10000000-0000-4000-8000-000000000072"), // X2's accepted invitation
	hash(0x18), // the invitation sign-up's link
	hash(0x19), // the invitation sign-up's invitation token
	hash(0x42), // credential salt
	hash(0x43), // credential hash
	hash(0x34), // idempotency request hash
	hash(0x35), // idempotency response
	"invitee-x@example.test", "Rumah X", "Rumah X2", "Pemilik X", "Tamu X",
}

// now is the fixture's clock: every token and invitation of the fixture is
// live at it, and both jobs are due.
var now = time.Date(2026, 11, 14, 12, 0, 0, 0, time.UTC)

// queryPackages makes the Queries of every sqlc package. A package the
// generated code has and this map lacks fails the test.
var queryPackages = map[string]func(pgx.Tx) any{
	"commanddb":   func(tx pgx.Tx) any { return commanddb.New(tx) },
	"householddb": func(tx pgx.Tx) any { return householddb.New(tx) },
	"identitydb":  func(tx pgx.Tx) any { return identitydb.New(tx) },
	"jobsdb":      func(tx pgx.Tx) any { return jobsdb.New(tx) },
	"readdb":      func(tx pgx.Tx) any { return readdb.New(tx) },
	"syncdb":      func(tx pgx.Tx) any { return syncdb.New(tx) },
}

// queryRoles names the database role a query runs as in production when it
// is not app_rw (ruling P7; 09.2 §3 layer 5; internal/jobs).
var queryRoles = map[string]string{
	"jobsdb.ClaimDue": "job_runner",
	"jobsdb.FailJob":  "job_runner",
	"jobsdb.Listen":   "job_runner",
}

// isolationValue is an identifier pair (x, y), or a fixed value for a
// parameter that is not an identifier.
type isolationValue struct {
	x, y  any
	fixed any
}

func pair(x, y any) isolationValue { return isolationValue{x: x, y: y} }
func fixed(v any) isolationValue   { return isolationValue{fixed: v} }

func (v isolationValue) identifier() bool { return v.x != nil }

// isolationValues maps every parameter name (snake case) to its value. A
// name whose meaning differs by query (id, token_hash, …) is keyed
// "<package>.<Query>.<name>" instead; that key wins. An integer converts to
// the parameter's integer type.
var isolationValues = map[string]isolationValue{
	// Identifiers.
	"household_id":          pair(householdX, householdY),
	"scope_id":              pair(householdX, householdY),
	"account_id":            pair(accountX, accountY),
	"invitee_account_id":    pair(inviteeX, inviteeY),
	"member_id":             pair(memberX, memberY),
	"inviter":               pair(memberX, memberY),
	"actor":                 pair(memberX, memberY),
	"member_ids":            pair([]pgtype.UUID{memberX}, []pgtype.UUID{memberY}),
	"object_id":             pair(householdX, householdY), // the object of X's event and feed row
	"object_ids":            pair([]pgtype.UUID{householdX}, []pgtype.UUID{householdY}),
	"session_id":            pair(sessionX, sessionY),
	"idempotency_key":       pair(idempotencyKeyX, idempotencyKeyY),
	"email":                 pair(emailX, emailY),
	"setup_token_hash":      pair(setupTokenX, setupTokenY),
	"invitation_token_hash": pair(invitationTokenX, invitationTokenY),
	"cookie_hash":           pair(cookieHashX, cookieHashY),
	"presented":             pair(sessionTokenX, sessionTokenY),
	"unique_key":            pair(text(invitationX.String()), text(invitationY.String())),

	// Y's value is an id no membership holds; membership.account_id has no
	// foreign key (identity mapping, 07 §8.2), so Y's own call inserts.
	"householddb.InsertMembership.account_id": pair(accountX, inviteeY),
	"householddb.JoinMembership.account_id":   pair(accountX, inviteeY),
	// The invitee of X2's accepted invitation.
	"householddb.OwnAcceptance.account_id": pair(inviteeX, inviteeY),
	// A new record's key is one no record holds.
	"commanddb.InsertIdempotencyRecord.idempotency_key": pair(idempotencyKeyX, uuid("20000000-0000-4000-8000-000000000082")),

	"householddb.InsertHousehold.id":        pair(householdX, householdY),
	"householddb.NextHouseholdVersion.id":   pair(householdX, householdY),
	"householddb.Invitation.id":             pair(invitationX, invitationY),
	"householddb.TransitionInvitation.id":   pair(invitationX, invitationY),
	"householddb.ReplaceInvitationToken.id": pair(invitationX, invitationY),
	"householddb.InvitationPreview.id":      pair(invitationX, invitationY),
	"householddb.AcceptInvitation.id":       pair(invitationX, invitationY),
	"identitydb.InsertAccount.id":           pair(accountX, accountY),
	"identitydb.GetSignupVia.id":            pair(accountX, accountY),
	"identitydb.GetSession.id":              pair(sessionX, sessionY),
	"identitydb.ConfirmCookie.id":           pair(sessionX, sessionY),
	"identitydb.RotateCookie.id":            pair(sessionX, sessionY),
	"identitydb.ReissueCookie.id":           pair(sessionX, sessionY),
	"identitydb.TouchSession.id":            pair(sessionX, sessionY),
	"identitydb.RevokeSession.id":           pair(sessionX, sessionY),
	"identitydb.RotateAccessToken.id":       pair(sessionX, sessionY),
	// keep is an id no session has, so Y's own call revokes Y's one session.
	"identitydb.RevokeSessions.keep": pair(sessionX, uuid("20000000-0000-4000-8000-000000000032")),
	"jobsdb.DeleteJob.id":            pair(jobX, jobY),
	"jobsdb.FailJob.id":              pair(jobX, jobY),
	"readdb.AccountEmail.id":         pair(accountX, accountY),

	// The Y value of a new token (0x6?) is one no row holds, so Y's own call
	// writes it.
	"householddb.ReplaceInvitationToken.token_hash":            pair(invitationTokenX, invitationTokenY),
	"householddb.OpenInvitationByToken.token_hash":             pair(invitationTokenX, invitationTokenY),
	"householddb.AcceptInvitation.token_hash":                  pair(invitationTokenX, invitationTokenY),
	"identitydb.InsertSetupToken.token_hash":                   pair(setupTokenX, hash(0x61)),
	"identitydb.GetSetupTokenForEmail.token_hash":              pair(setupTokenX, setupTokenY),
	"identitydb.UseSetupToken.token_hash":                      pair(setupTokenX, setupTokenY),
	"identitydb.UpsertSignupVerification.token_hash":           pair(signupLinkX, hash(0x68)),
	"identitydb.UpsertInvitationSignupVerification.token_hash": pair(signupLinkX, hash(0x69)),
	"identitydb.GetSignupVerification.token_hash":              pair(hash(0x18), hash(0x28)), // the invitation sign-ups, read in their account's scope
	"identitydb.UseSignupVerification.token_hash":              pair(signupLinkX, signupLinkY),
	"identitydb.InsertSession.token_hash":                      pair(sessionTokenX, hash(0x62)),
	"identitydb.GetSessionByTokenHash.token_hash":              pair(sessionTokenX, sessionTokenY),
	"identitydb.GetSessionByCookieHash.token_hash":             pair(sessionTokenX, sessionTokenY),
	"identitydb.GetSessionByCookieHashForUpdate.token_hash":    pair(sessionTokenX, sessionTokenY),
	"identitydb.RotateCookie.token_hash":                       pair(sessionTokenX, hash(0x63)),
	"identitydb.ReissueCookie.token_hash":                      pair(sessionTokenX, hash(0x64)),
	"identitydb.RotateAccessToken.token_hash":                  pair(sessionTokenX, hash(0x65)),
	"identitydb.InsertRefreshToken.token_hash":                 pair(refreshTokenX, hash(0x66)),
	"identitydb.GetRefreshToken.token_hash":                    pair(refreshTokenX, refreshTokenY),
	"identitydb.UseRefreshToken.token_hash":                    pair(refreshTokenX, refreshTokenY),
	"identitydb.InsertRecoveryToken.token_hash":                pair(recoveryTokenX, hash(0x67)),
	"identitydb.GetRecoveryToken.token_hash":                   pair(recoveryTokenX, recoveryTokenY),
	"identitydb.UseRecoveryToken.token_hash":                   pair(recoveryTokenX, recoveryTokenY),

	// Fixed values: times, counts, states, and content.
	"now":                                   fixed(now),
	"created_at":                            fixed(now),
	"joined_at":                             fixed(now),
	"valid_from":                            fixed(now),
	"recorded_at":                           fixed(now),
	"issued_at":                             fixed(now),
	"acknowledged_at":                       fixed(now),
	"occurred_at":                           fixed(now),
	"server_time":                           fixed(now),
	"updated_at":                            fixed(now),
	"closed_at":                             fixed(now),
	"run_at":                                fixed(now),
	"expires_at":                            fixed(now.Add(7 * 24 * time.Hour)),
	"absolute_expiry":                       fixed(now.Add(90 * 24 * time.Hour)),
	"prev_valid_until":                      fixed(now.Add(time.Minute)),
	"lease_until":                           fixed(now.Add(5 * time.Minute)),
	"retry_at":                              fixed(now.Add(10 * time.Second)),
	"idle_timeout":                          fixed(30 * 24 * time.Hour),
	"name":                                  fixed("Rumah Baru"),
	"time_zone":                             fixed("Asia/Jakarta"),
	"display_name":                          fixed(text("Anggota Baru")),
	"invitee_display_name":                  fixed(text("Anggota Baru")),
	"version":                               fixed(0), // the compare-and-set version of Y's invitation
	"householddb.AcknowledgeNotice.version": fixed(1), // the operator notice's version
	"householddb.IssueNotice.version":       fixed(1),
	"variant":                               fixed("couple"),
	"text_sha256":                           fixed(hash(0x30)),
	"text_version":                          fixed(2), // a version neither account has stated, so a write would land
	"issue_seq":                             fixed(2),
	"from_state":                            fixed("open"),
	"householddb.InsertObjectEvent.from_state": fixed(pgtype.Text{}),
	"to_state":                 fixed("withdrawn"),
	"clear_invitee":            fixed(false),
	"object_kind":              fixed("household"),
	"actor_kind":               fixed("member"),
	"rule":                     fixed(pgtype.Text{}),
	"seq":                      fixed(1),
	"command":                  fixed("CreateHousehold"),
	"request_hash":             fixed(hash(0x31)),
	"response":                 fixed(hash(0x32)),
	"kind":                     fixed("sign_in"),
	"jobsdb.InsertJob.kind":    fixed("household.expire_invitation"),
	"payload":                  fixed(invitationY.Bytes[:]),
	"detail":                   fixed(pgtype.Text{}),
	"client_kind":              fixed("web"),
	"device_mode":              fixed("personal"),
	"ua_summary":               fixed("Firefox"),
	"client_version":           fixed("1.0.0"),
	"salt":                     fixed(hash(0x40)),
	"hash":                     fixed(hash(0x41)),
	"time_cost":                fixed(3),
	"memory_kib":               fixed(65536),
	"parallelism":              fixed(1),
	"purpose":                  fixed("change_password"),
	"reason":                   fixed("user"),
	"signup_via":               fixed("setup_token"),
	"attempt":                  fixed(0),
	"max_attempts":             fixed(10),
	"last_error":               fixed(text("handler")),
	"stopped":                  fixed(false),
	"n":                        fixed(1),
	"lim":                      fixed(100),
	"after":                    fixed(0),
	"readdb.MemberViews.after": fixed(pgtype.UUID{Valid: true}), // a page marker, not a lookup
	"positions":                fixed([]int64{100}),
	"object_kinds":             fixed([]string{"household"}),
}

// isolationException lets one (query, parameter) answer X's value
// differently from a fresh one; parameter "" is the run with Y's values
// only, "*" the run with every identifier X's. It waives the comparison
// only: the result must still hold no X value and no row X owns may change,
// unless crossHousehold, for a role that reads every household by design.
type isolationException struct {
	query, param, reason string
	crossHousehold       bool
}

// isolationExceptions are the only failures allowed, each with its source.
// An exception that no longer fails fails the test.
var isolationExceptions = []isolationException{
	{"householddb.InsertMembership", "account_id", membershipOneHousehold, false},
	{"householddb.JoinMembership", "account_id", membershipOneHousehold, false},
	{"householddb.AcceptInvitation", "invitee_account_id", "invitation_invitee: one open acceptance per account across " +
		"households, a recorded exception to 09.2 §3 layer 2 keyed by the caller's own account; the handler gives " +
		"only the session's account (09.2 §5.1 note on plan 1.6 decision 20)", false},
	{"identitydb.InsertSession", "token_hash", randomTokenHash, false},
	{"identitydb.RotateCookie", "token_hash", randomTokenHash, false},
	{"identitydb.ReissueCookie", "token_hash", randomTokenHash, false},
	{"identitydb.RotateAccessToken", "token_hash", randomTokenHash, false},
	{"identitydb.InsertRefreshToken", "token_hash", randomTokenHash, false},
	{"identitydb.InsertRecoveryToken", "token_hash", randomTokenHash, false},
	{"identitydb.InsertRefreshToken", "session_id", sessionForeignKey, false},
	{"identitydb.InsertReauthGrant", "session_id", sessionForeignKey, false},
	{"identitydb.InsertSecurityEvent", "session_id", sessionForeignKey, false},
	// ClaimDue returns and leases X's due job, and FailJob writes it: both
	// read and change every household's job rows by design, so the leak and
	// digest checks cannot hold for them.
	{"jobsdb.ClaimDue", "", jobRunnerAcrossHouseholds, true},
	{"jobsdb.FailJob", "id", jobRunnerAcrossHouseholds, true},
}

// ownCallFails are the queries whose call with Y's values cannot succeed in
// the test's scope, each with its source. Every other query must answer Y's
// own call without an error and, unless it is exec, with at least one row,
// so a difference X makes is not hidden behind an error both runs share. An
// entry that no longer fails fails the test.
var ownCallFails = []struct{ query, reason string }{
	{"identitydb.GetSetupTokenForEmail", lookupScopeOnly},
	{"identitydb.InsertSetupToken", lookupScopeOnly},
	{"identitydb.UseSetupToken", lookupScopeOnly},
	{"identitydb.UpsertSignupVerification", lookupScopeOnly},
	{"identitydb.UpsertInvitationSignupVerification", lookupScopeOnly},
	{"identitydb.UseSignupVerification", lookupScopeOnly},
	{"householddb.OwnAcceptance", lookupScopeOnly},
	{"identitydb.InsertAccount", "the account it writes must be the scope's (WITH CHECK), and the scope's account " +
		"exists; sign-up writes in the new account's own scope (09.2 §3 global tables), the production-scope pass " +
		"(10.1 §5.1 catatan 14 November 2026 (FS-1 1.7 PR C, uji isolasi per kueri) item 5)"},
	{"householddb.InsertHousehold", "the household it writes must be the scope's (WITH CHECK, 09.2 §3 note of 14 " +
		"November 2026), and the scope's household exists; CreateHousehold's scope is the new id"},
	{"identitydb.InsertCredential", "one credential per account (primary key), and Y's account holds the one " +
		"GetCredential and UpdateCredential read"},
	{"householddb.InsertInvitation", "one live invitation per household (06 §3.6 F01; invitation_live), and Y's " +
		"live invitation is the one the other invitation queries act on"},
}

const (
	membershipOneHousehold = "membership_one_household: one household per account across households, a recorded " +
		"exception to 09.2 §3 layer 2 keyed by the caller's own account; the handler gives only the session's " +
		"account (09.2 §5.1 note on plan 1.6 decision 20)"
	randomTokenHash = "a new token's hash is unique across accounts, but it is the SHA-256 of a 256-bit secret " +
		"the server draws, so no caller can name another's (10.1 §5.1 row \"Pemeriksaan unique di luar RLS\")"
	sessionForeignKey = "the foreign key to session(id) is checked outside row-level security; a recorded " +
		"exception, without a schema change (10.1 §5.1 catatan 14 November 2026 (FS-1 1.7 PR C, uji isolasi per kueri))"
	lookupScopeOnly = "its policy clause holds only in a lookup scope (Scope{TokenHash} or Scope{AccountID} " +
		"without a household), which the fixed scope never sets; the production-scope pass is due before GR-4 " +
		"(10.1 §5.1 catatan 14 November 2026 (FS-1 1.7 PR C, uji isolasi per kueri) item 5)"
	jobRunnerAcrossHouseholds = "job_runner claims and fails jobs across households by design; payload holds " +
		"identifiers only (09.2 §3 row \"job (pengambilan)\" and its note of 14 November 2026, Q1 (b); ruling P7)"
)

// ownership is, per table, the rows a household owns: $1 holds the
// household's id and its accounts' ids, $2 its token hashes. The X digest
// covers the rows these select for X.
var ownership = map[string]string{ //nolint:gosec // G101: SQL predicates naming token columns, not credentials
	"household":           "id = ANY($1::uuid[])",
	"membership":          "household_id = ANY($1::uuid[])",
	"approval_authority":  "household_id = ANY($1::uuid[])",
	"household_rule":      "household_id = ANY($1::uuid[])",
	"relationship":        "household_id = ANY($1::uuid[])",
	"invitation":          "household_id = ANY($1::uuid[])",
	"object_event":        "household_id = ANY($1::uuid[])",
	"job":                 "household_id = ANY($1::uuid[])",
	"reader_feed":         "household_id = ANY($1::uuid[])",
	"reader_feed_counter": "household_id = ANY($1::uuid[])",
	"idempotency_record":  "scope_id = ANY($1::uuid[])",
	"account":             "id = ANY($1::uuid[])",
	"credential":          "account_id = ANY($1::uuid[])",
	"session":             "account_id = ANY($1::uuid[])",
	"refresh_token":       "account_id = ANY($1::uuid[])",
	"reauth_grant":        "account_id = ANY($1::uuid[])",
	"recovery_token":      "account_id = ANY($1::uuid[])",
	"security_event":      "account_id = ANY($1::uuid[])",
	"consent_record":      "account_id = ANY($1::uuid[])",
	"operator_notice_ack": "account_id = ANY($1::uuid[])",
	"setup_token":         "account_id = ANY($1::uuid[]) OR token_hash = ANY($2::bytea[])",
	"signup_verification": "account_id = ANY($1::uuid[]) OR setup_token_hash = ANY($2::bytea[]) OR invitation_token_hash = ANY($2::bytea[])",
}

type owner struct {
	ids    []pgtype.UUID
	tokens [][]byte
}

var (
	ownerX = owner{
		ids:    []pgtype.UUID{householdX, uuid("10000000-0000-4000-8000-000000000002"), accountX, inviteeX},
		tokens: [][]byte{setupTokenX, invitationTokenX},
	}
	ownerY = owner{
		ids:    []pgtype.UUID{householdY, accountY, inviteeY},
		tokens: [][]byte{setupTokenY, invitationTokenY},
	}
)

// query is one method of a sqlc Queries, found in the generated code.
type query struct {
	pkg, name string
	kind      string   // exec, execrows, one or many
	params    []string // snake case, in order
}

func (q query) String() string { return q.pkg + "." + q.name }

// outcome is what one run of a query answers.
type outcome struct {
	code    string // SQLSTATE, empty without an error
	rows    int    // rows returned or changed; 0 for exec
	scalars []any  // every boolean and integer in the result, in order
	leaked  []any  // X values in the result
	digest  string // of every row X owns, after the query
}

// isValueX treats every :11:11 time as X's, so no Y or fixed time may fall there.
func TestNoYOrFixedTimeIsAnXMark(t *testing.T) {
	for name, v := range isolationValues {
		for _, c := range []any{v.fixed, v.y} {
			if at, ok := c.(time.Time); ok && at.Minute() == 11 && at.Second() == 11 {
				t.Errorf("%s: %v falls on :11:11, the X mark", name, at)
			}
		}
	}
}

func TestXTimesAreRecognised(t *testing.T) {
	for _, at := range []string{"2026-11-13T12:11:11Z", "2027-02-12T12:11:11Z", "2026-11-14T13:11:11Z"} {
		parsed, err := time.Parse(time.RFC3339, at)
		if err != nil {
			t.Fatal(err)
		}
		if !isValueX(parsed) {
			t.Errorf("%s is an X time and must be recognised", at)
		}
	}
	if isValueX(now) {
		t.Error("the fixture clock is not an X time")
	}
}

func TestEveryQueryIsIsolated(t *testing.T) {
	queries := discoverQueries(t)
	d := databasetest.Migrated(t)
	ctx := t.Context()
	admin := d.Pool(t, d.AdminRole)
	if _, err := admin.Exec(ctx, fixture); err != nil {
		t.Fatalf("fixture: %v", err)
	}
	digestSQL := buildDigestSQL()
	before := digest(t, admin, digestSQL)

	t.Run("every table holds rows of both households in the fixture", func(t *testing.T) {
		// 09.2 §3: currency is the product's table, the same for everyone;
		// goose_db_version is the migration tool's.
		rows, err := admin.Query(ctx, `
			SELECT table_name FROM information_schema.tables
			 WHERE table_schema = 'public' AND table_type = 'BASE TABLE'
			   AND table_name NOT IN ('goose_db_version', 'currency')
			 ORDER BY table_name`)
		if err != nil {
			t.Fatal(err)
		}
		tables, err := pgx.CollectRows(rows, pgx.RowTo[string])
		if err != nil {
			t.Fatal(err)
		}
		if got, want := tables, slices.Sorted(maps.Keys(ownership)); !slices.Equal(got, want) {
			t.Fatalf("tables %v, ownership covers %v", got, want)
		}
		for _, table := range tables {
			for name, o := range map[string]owner{"X": ownerX, "Y": ownerY} {
				var n int
				args := []any{o.ids}
				if strings.Contains(ownership[table], "$2") {
					args = append(args, o.tokens)
				}
				err := admin.QueryRow(ctx, "SELECT count(*) FROM "+table+" WHERE "+ownership[table], args...).Scan(&n)
				if err != nil {
					t.Fatal(err)
				}
				if n == 0 {
					t.Errorf("%s has no row of household %s", table, name)
				}
			}
		}
	})

	t.Run("every query parameter has an isolation value", func(t *testing.T) {
		used := map[string]bool{}
		for _, q := range queries {
			t.Run(q.String(), func(t *testing.T) {
				for _, p := range q.params {
					key, ok := valueKey(q, p)
					if !ok {
						t.Errorf("parameter %s has no entry in isolationValues", p)
					}
					used[key] = true
				}
			})
		}
		for key := range isolationValues {
			if !used[key] {
				t.Errorf("isolationValues[%q] is used by no query", key)
			}
		}
		for name := range queryRoles {
			if !slices.ContainsFunc(queries, func(q query) bool { return q.String() == name }) {
				t.Errorf("queryRoles names %s, which no query is", name)
			}
		}
	})

	pools := map[string]*pgxpool.Pool{"app_rw": rolePool(t, d, "app_rw"), "job_runner": rolePool(t, d, "job_runner")}
	triggered, ownCallTriggered := map[string]bool{}, map[string]bool{}

	t.Run("SEC-02 another household's values behave as values that never existed", func(t *testing.T) {
		for _, q := range queries {
			t.Run(q.String(), func(t *testing.T) {
				r := runner{pool: pools[roleOf(q)], digestSQL: digestSQL, query: q}
				ys := map[string]any{}
				for _, p := range q.params {
					key, ok := valueKey(q, p)
					if !ok {
						t.Fatalf("parameter %s has no isolation value", p)
					}
					v := isolationValues[key]
					ys[p] = v.fixed
					if v.identifier() {
						ys[p] = v.y
					}
				}
				var failures []string
				check := func(param string, compared, leakOrChanged []string) {
					e, excepted := exception(q, param)
					if excepted && len(compared) > 0 || e.crossHousehold && len(leakOrChanged) > 0 {
						triggered[q.String()+"/"+param] = true
						t.Logf("%s: excepted (%s): %s", describe(param), e.reason, strings.Join(slices.Concat(compared, leakOrChanged), "; "))
						compared = nil
						if e.crossHousehold {
							leakOrChanged = nil
						}
					}
					if problems := slices.Concat(compared, leakOrChanged); len(problems) > 0 {
						failures = append(failures, describe(param)+": "+strings.Join(problems, "; "))
					}
				}
				compare := func(param string, x, fresh map[string]any) {
					runX, runR := r.run(t, x), r.run(t, fresh)
					var compared []string
					if runX.code != runR.code || runX.rows != runR.rows || !reflect.DeepEqual(runX.scalars, runR.scalars) {
						compared = append(compared, fmt.Sprintf("X answers (code %q, rows %d, %v), a fresh value (code %q, rows %d, %v)",
							runX.code, runX.rows, runX.scalars, runR.code, runR.rows, runR.scalars))
					}
					check(param, compared, leakOrChange(runX, before))
				}

				base := r.run(t, ys)
				check("", nil, leakOrChange(base, before))
				ownFails := base.code != "" || q.kind != "exec" && base.rows == 0
				if reason, declared := ownCallFailure(q); declared {
					if ownFails {
						ownCallTriggered[q.String()] = true
						t.Logf("Y's own call fails (code %q, rows %d), declared: %s", base.code, base.rows, reason)
					}
				} else if ownFails {
					failures = append(failures, fmt.Sprintf("Y's own call fails (code %q, rows %d), which would hide what X changes", base.code, base.rows))
				}

				allX, allFresh := maps.Clone(ys), maps.Clone(ys)
				identifiers := 0
				for _, p := range q.params {
					key, _ := valueKey(q, p)
					v := isolationValues[key]
					if !v.identifier() {
						continue
					}
					identifiers++
					x, fresh := maps.Clone(ys), maps.Clone(ys)
					x[p], fresh[p] = v.x, freshLike(t, v.x)
					allX[p], allFresh[p] = v.x, fresh[p]
					compare(p, x, fresh)
				}
				// A query that pins two identifiers together (household and
				// id) answers nothing when only one is X's, whatever the
				// policies say; with all of them X's it would find X's row.
				if identifiers > 1 {
					compare("*", allX, allFresh)
				}
				for _, f := range failures {
					t.Error(f)
				}
			})
		}
	})

	t.Run("every declared exception still applies", func(t *testing.T) {
		for _, e := range isolationExceptions {
			if !triggered[e.query+"/"+e.param] {
				t.Errorf("%s %s no longer fails; remove its exception", e.query, describe(e.param))
			}
		}
		for _, o := range ownCallFails {
			if !ownCallTriggered[o.query] {
				t.Errorf("%s: Y's own call no longer fails; remove it from ownCallFails", o.query)
			}
		}
	})

	if after := digest(t, admin, digestSQL); after != before {
		t.Errorf("rows of household X changed after the forced rollbacks")
	}
}

func describe(param string) string {
	switch param {
	case "":
		return "with Y's values only"
	case "*":
		return "with every identifier X's"
	}
	return "parameter " + param
}

func leakOrChange(o outcome, before string) []string {
	var problems []string
	if len(o.leaked) > 0 {
		problems = append(problems, fmt.Sprintf("the result holds X values %v", o.leaked))
	}
	if o.digest != before {
		problems = append(problems, "rows X owns changed")
	}
	return problems
}

func exception(q query, param string) (isolationException, bool) {
	for _, e := range isolationExceptions {
		if e.query == q.String() && e.param == param {
			return e, true
		}
	}
	return isolationException{}, false
}

func ownCallFailure(q query) (string, bool) {
	for _, o := range ownCallFails {
		if o.query == q.String() {
			return o.reason, true
		}
	}
	return "", false
}

func valueKey(q query, param string) (string, bool) {
	for _, key := range []string{q.String() + "." + param, param} {
		if _, ok := isolationValues[key]; ok {
			return key, true
		}
	}
	return "", false
}

func roleOf(q query) string {
	if role, ok := queryRoles[q.String()]; ok {
		return role
	}
	return "app_rw"
}

var (
	// annotation is sqlc's header line in each query's SQL constant.
	annotation = regexp.MustCompile(`^-- name: (\w+) :(\w+)`)
	// outDirectory is the output directory of one package in sqlc.yaml.
	outDirectory = regexp.MustCompile(`(?m)^\s+out:\s*(\S+)\s*$`)
	// queryKinds are the kinds the runner knows how to count; a batch or
	// copyfrom query fails the test until it does.
	queryKinds = []string{"exec", "execrows", "one", "many"}
)

// discoverQueries reads every method of *Queries, the kind of its query, and
// its parameters, from the generated code of every package sqlc.yaml names.
func discoverQueries(t *testing.T) []query {
	t.Helper()
	config, err := os.ReadFile("../../sqlc.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var files []string
	for _, m := range outDirectory.FindAllStringSubmatch(string(config), -1) {
		matched, err := filepath.Glob(filepath.Join("../..", m[1], "*.go"))
		if err != nil {
			t.Fatal(err)
		}
		files = append(files, matched...)
	}
	var queries []query
	packages := map[string]bool{}
	for _, file := range files {
		f, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		pkg := f.Name.Name
		packages[pkg] = true
		structs, kinds := map[string][]string{}, map[string]string{}
		for _, decl := range f.Decls {
			g, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range g.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if st, ok := s.Type.(*ast.StructType); ok {
						for _, field := range st.Fields.List {
							for _, n := range field.Names {
								structs[s.Name.Name] = append(structs[s.Name.Name], snake(n.Name))
							}
						}
					}
				case *ast.ValueSpec:
					for _, v := range s.Values {
						if lit, ok := v.(*ast.BasicLit); ok && lit.Kind == token.STRING {
							sql, err := strconv.Unquote(lit.Value)
							if err != nil {
								t.Fatal(err)
							}
							if m := annotation.FindStringSubmatch(sql); m != nil {
								kinds[m[1]] = m[2]
							}
						}
					}
				}
			}
		}
		for _, decl := range f.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			// WithTx is sqlc's own helper in db.go, not a query.
			if !ok || fn.Recv == nil || !isQueriesReceiver(fn.Recv) || fn.Name.Name == "WithTx" {
				continue
			}
			q := query{pkg: pkg, name: fn.Name.Name, kind: kinds[fn.Name.Name]}
			if !slices.Contains(queryKinds, q.kind) {
				t.Fatalf("%s: query kind %q, the runner knows %v", q, q.kind, queryKinds)
			}
			for _, field := range fn.Type.Params.List[1:] { // [0] is ctx
				if ident, ok := field.Type.(*ast.Ident); ok && structs[ident.Name] != nil {
					q.params = append(q.params, structs[ident.Name]...)
					continue
				}
				for _, n := range field.Names {
					q.params = append(q.params, snake(n.Name))
				}
			}
			queries = append(queries, q)
		}
	}
	if got, want := slices.Sorted(maps.Keys(packages)), slices.Sorted(maps.Keys(queryPackages)); !slices.Equal(got, want) {
		t.Fatalf("sqlc packages %v, queryPackages %v", got, want)
	}
	if len(queries) == 0 {
		t.Fatal("no query found")
	}
	return queries
}

func isQueriesReceiver(recv *ast.FieldList) bool {
	star, ok := recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return false
	}
	ident, ok := star.X.(*ast.Ident)
	return ok && ident.Name == "Queries"
}

// snake turns a Go name into the snake case of its SQL parameter:
// InviteeAccountID → invitee_account_id, TextSha256 → text_sha256.
func snake(name string) string {
	var b strings.Builder
	runes := []rune(name)
	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) && !unicode.IsUpper(runes[i-1]) {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

// rolePool logs in as the superuser and sets role for the session: a
// query in database.Write runs as role, bound by its grants and policies
// (ruling P10), and the same transaction can RESET ROLE to read the X digest
// past row-level security before the forced rollback, which also restores
// the role.
func rolePool(t *testing.T, d databasetest.Database, role string) *pgxpool.Pool {
	t.Helper()
	config, err := pgxpool.ParseConfig(d.URL(d.AdminRole))
	if err != nil {
		t.Fatal(err)
	}
	config.MaxConns = 1
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, "SET ROLE "+pgx.Identifier{role}.Sanitize())
		return err
	}
	pool, err := pgxpool.NewWithConfig(t.Context(), config)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// buildDigestSQL hashes every row X owns, table by table (ownership).
func buildDigestSQL() string {
	var parts []string
	for _, table := range slices.Sorted(maps.Keys(ownership)) {
		parts = append(parts, fmt.Sprintf(
			"SELECT '%s' AS t, md5(coalesce(string_agg(r::text, E'\\n' ORDER BY r::text), '')) AS d FROM %s r WHERE %s",
			table, pgx.Identifier{table}.Sanitize(), ownership[table]))
	}
	return "SELECT md5(string_agg(t || ':' || d, ',' ORDER BY t)) FROM (" + strings.Join(parts, " UNION ALL ") + ") s"
}

type querier interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func digest(t *testing.T, db querier, digestSQL string) string {
	t.Helper()
	var d string
	if err := db.QueryRow(t.Context(), digestSQL, ownerX.ids, ownerX.tokens).Scan(&d); err != nil {
		t.Fatalf("X digest: %v", err)
	}
	return d
}

type runner struct {
	pool      *pgxpool.Pool
	digestSQL string
	query     query
}

var errForcedRollback = errors.New("forced rollback")

// run calls the query once in database.Write as household Y's member, then
// reads the X digest in the same transaction and rolls it back.
func (r runner) run(t *testing.T, args map[string]any) outcome {
	t.Helper()
	ctx := t.Context()
	var o outcome
	scope := database.Scope{HouseholdID: householdY, AccountID: accountY}
	err := database.Write(ctx, r.pool, scope, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, "SAVEPOINT isolation"); err != nil {
			return err
		}
		result, err := r.call(ctx, tx, args)
		switch {
		case err == nil:
			o.rows, o.scalars, o.leaked = inspect(r.query.kind, result)
		case errors.Is(err, pgx.ErrNoRows):
		default:
			if o.code = databasetest.PgErrorCode(err); o.code == "" {
				return fmt.Errorf("not a database answer: %w", err)
			}
			if _, err := tx.Exec(ctx, "ROLLBACK TO SAVEPOINT isolation"); err != nil {
				return err
			}
		}
		if _, err := tx.Exec(ctx, "RESET ROLE"); err != nil {
			return err
		}
		o.digest = digest(t, tx, r.digestSQL)
		return errForcedRollback
	})
	if !errors.Is(err, errForcedRollback) {
		t.Fatalf("run: %v", err)
	}
	return o
}

// call runs the query's method with args by name, and returns its result
// (invalid for exec) and its error.
func (r runner) call(ctx context.Context, tx pgx.Tx, args map[string]any) (reflect.Value, error) {
	method := reflect.ValueOf(queryPackages[r.query.pkg](tx)).MethodByName(r.query.name)
	in := []reflect.Value{reflect.ValueOf(ctx)}
	if method.Type().NumIn() == 2 {
		pt := method.Type().In(1)
		var arg reflect.Value
		if pt.Kind() == reflect.Struct && strings.HasSuffix(pt.Name(), "Params") {
			arg = reflect.New(pt).Elem()
			for i := range pt.NumField() {
				arg.Field(i).Set(convert(args[snake(pt.Field(i).Name)], pt.Field(i).Type))
			}
		} else {
			arg = convert(args[r.query.params[0]], pt)
		}
		in = append(in, arg)
	}
	out := method.Call(in)
	err, _ := out[len(out)-1].Interface().(error)
	if len(out) == 1 {
		return reflect.Value{}, err
	}
	return out[0], err
}

func convert(v any, to reflect.Type) reflect.Value {
	value := reflect.ValueOf(v)
	if !value.Type().AssignableTo(to) {
		value = value.Convert(to)
	}
	return value
}

// inspect counts the rows of a result and collects its booleans, integers
// and X values.
func inspect(kind string, result reflect.Value) (rows int, scalars, leaked []any) {
	switch kind {
	case "exec":
		return 0, nil, nil
	case "execrows":
		return int(result.Int()), nil, nil
	case "many":
		rows = result.Len()
	default:
		rows = 1
	}
	walk(result, func(v any) {
		switch v := v.(type) {
		case bool, int16, int32, int64:
			scalars = append(scalars, v)
		default:
			if isValueX(v) {
				leaked = append(leaked, v)
			}
		}
	})
	return rows, scalars, leaked
}

var (
	uuidType = reflect.TypeFor[pgtype.UUID]()
	textType = reflect.TypeFor[pgtype.Text]()
	timeType = reflect.TypeFor[time.Time]()
)

// walk visits every leaf of a query result: UUIDs, texts, times, byte
// strings, strings, booleans and integers.
func walk(v reflect.Value, visit func(any)) {
	switch {
	case v.Type() == uuidType, v.Type() == textType, v.Type() == timeType:
		visit(v.Interface())
	case v.Type() == reflect.TypeFor[time.Duration]():
	case v.Kind() == reflect.Slice && v.Type().Elem().Kind() == reflect.Uint8:
		visit(v.Bytes())
	case v.Kind() == reflect.Slice:
		for i := range v.Len() {
			walk(v.Index(i), visit)
		}
	case v.Kind() == reflect.Struct:
		for i := range v.NumField() {
			walk(v.Field(i), visit)
		}
	default:
		visit(v.Interface())
	}
}

func isValueX(v any) bool {
	// Every X time of the fixture is written at minute 11, second 11; no Y,
	// fixed or fresh time falls there (xMark below).
	if t, ok := v.(time.Time); ok && t.Minute() == 11 && t.Second() == 11 {
		return true
	}
	for _, x := range valuesX() {
		if equalValue(v, x) {
			return true
		}
	}
	return false
}

// valuesX is every X value of the fixture: the X side of each identifier
// pair, and otherValuesX.
func valuesX() []any {
	values := slices.Clone(otherValuesX)
	for _, v := range isolationValues {
		if !v.identifier() {
			continue
		}
		if ids, ok := v.x.([]pgtype.UUID); ok {
			for _, id := range ids {
				values = append(values, id)
			}
			continue
		}
		values = append(values, v.x)
	}
	return values
}

func equalValue(v, x any) bool {
	if t, ok := v.(pgtype.Text); ok {
		v = t.String
	}
	if t, ok := x.(pgtype.Text); ok {
		x = t.String
	}
	switch v := v.(type) {
	case pgtype.UUID:
		x, ok := x.(pgtype.UUID)
		return ok && v.Valid && v == x
	case []byte:
		x, ok := x.([]byte)
		return ok && bytes.Equal(v, x)
	case string:
		x, ok := x.(string)
		return ok && strings.EqualFold(v, x)
	case time.Time:
		x, ok := x.(time.Time)
		return ok && v.Equal(x)
	}
	return false
}

// freshLike is a random value of x's type, which no row holds.
func freshLike(t *testing.T, x any) any {
	t.Helper()
	switch x := x.(type) {
	case pgtype.UUID:
		var id pgtype.UUID
		_, _ = rand.Read(id.Bytes[:]) // never returns an error (crypto/rand, Go 1.24 and later)
		id.Valid = true
		return id
	case []pgtype.UUID:
		ids := make([]pgtype.UUID, len(x))
		for i := range ids {
			ids[i] = freshLike(t, x[i]).(pgtype.UUID)
		}
		return ids
	case []byte:
		b := make([]byte, len(x))
		_, _ = rand.Read(b)
		return b
	case string:
		return rand.Text()
	case pgtype.Text:
		return text(rand.Text())
	}
	t.Fatalf("no fresh value for %T", x)
	return nil
}

func uuid(s string) pgtype.UUID {
	var id pgtype.UUID
	if err := id.Scan(s); err != nil {
		panic(err)
	}
	return id
}

func hash(b byte) []byte { return bytes.Repeat([]byte{b}, 32) }

func text(s string) pgtype.Text { return pgtype.Text{String: s, Valid: true} }
