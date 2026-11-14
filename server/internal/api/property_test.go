package api_test

// The stateful property test frame (14 §4.3 "Pembangkit perintah",
// "Pemeriksa", "Simetri"; 09.4 §4 rows "Invarian alur dan persetujuan" and
// "Isolasi"; plan 1.6 Q17; plan 1.7 decision 16): one random command
// generator from the contract, a machine of 2 to 5 accounts over the real
// server and database, and one checker per property (invariant_test.go). Step
// 2 adds the ledger's builders and the INV checkers to the same machine.

import (
	"cmp"
	"errors"
	"flag"
	"fmt"
	"maps"
	"net"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"pgregory.net/rapid"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/identity"
	"github.com/abiyanzimran/casa/server/internal/jobs"
)

// TestMain gives rapid the per-change defaults when the environment sets
// none: a fixed seed, since a per-change test is deterministic (ES §6), and
// 20 checks [Asumsi] (plan 1.7 decision 16; ruling P6). The weekly run sets
// RAPID_CHECKS and RAPID_SEED, which rapid reads itself. Fail files are
// never written: a shrunk failure becomes a named example test (14 §4.3).
func TestMain(m *testing.M) {
	defaults := map[string]string{"rapid.checks": "20", "rapid.seed": "1"}
	for name, env := range map[string]string{"rapid.checks": "RAPID_CHECKS", "rapid.seed": "RAPID_SEED"} {
		if os.Getenv(env) == "" {
			mustSetFlag(name, defaults[name])
		}
	}
	mustSetFlag("rapid.nofailfile", "true")
	os.Exit(m.Run())
}

func mustSetFlag(name, value string) {
	if err := flag.Set(name, value); err != nil {
		panic(err)
	}
}

const (
	passphrase = "a long enough passphrase"
	// unknownID is an id no object has, for a request that names nothing.
	unknownID = "00000000-0000-4000-8000-000000000000"
	// maxPages bounds a catch-up, so a feed that never ends fails instead
	// of looping.
	maxPages = 50
)

// propertyHarness is one household harness for the checks of one property:
// each check resets its database instead of making a new one.
type propertyHarness struct {
	*householdHarness
	runner *jobs.Runner
	start  time.Time
	// tables are the tables a check resets and counts: all but the
	// migrations' own and currency, whose rows 00001 seeds.
	tables []string
}

func startPropertyHarness(t *testing.T) *propertyHarness {
	t.Helper()
	h := &propertyHarness{householdHarness: startHousehold(t)}
	h.runner = h.householdHarness.runner(t)
	h.start = h.clock.Now()
	rows, err := h.admin.Query(t.Context(), `
		SELECT tablename FROM pg_tables
		 WHERE schemaname = 'public' AND tablename NOT IN ('currency', 'goose_db_version') ORDER BY tablename`)
	if err != nil {
		t.Fatal(err)
	}
	if h.tables, err = pgx.CollectRows(rows, pgx.RowTo[string]); err != nil {
		t.Fatal(err)
	}
	return h
}

// rowsSnapshot is what a resend may not change: the row count of every
// table but object_event and reader_feed, and every row of those two by key
// (events: "kind id seq" to "actor_kind rule"; feed: "member pos" to
// "object_kind object_id"), so that the rows a resend added can be told
// apart (checkResend); and the invitations whose expires_at is not after
// the clock.
type rowsSnapshot struct {
	counts       string
	events, feed map[string]string
	due          map[string]bool
}

func (h *propertyHarness) rows(t *rapid.T) rowsSnapshot {
	var counts []string
	for _, table := range h.tables {
		if table != "object_event" && table != "reader_feed" {
			counts = append(counts, fmt.Sprintf("'%s=' || (SELECT count(*) FROM %s)", table, pgx.Identifier{table}.Sanitize()))
		}
	}
	r := rowsSnapshot{events: map[string]string{}, feed: map[string]string{}, due: map[string]bool{}}
	if err := h.admin.QueryRow(t.Context(), "SELECT "+strings.Join(counts, " || ' ' || ")).Scan(&r.counts); err != nil {
		t.Fatal(err)
	}
	for into, query := range map[*map[string]string]string{
		&r.events: "SELECT concat_ws(' ', object_kind, object_id, seq), concat_ws(' ', actor_kind, coalesce(rule, '-')) FROM object_event",
		&r.feed:   "SELECT concat_ws(' ', member_id, pos), concat_ws(' ', object_kind, object_id) FROM reader_feed",
	} {
		rows, err := h.admin.Query(t.Context(), query)
		if err != nil {
			t.Fatal(err)
		}
		var key, value string
		if _, err := pgx.ForEachRow(rows, []any{&key, &value}, func() error { (*into)[key] = value; return nil }); err != nil {
			t.Fatal(err)
		}
	}
	due, err := h.admin.Query(t.Context(), "SELECT id::text FROM invitation WHERE expires_at <= $1", h.clock.Now())
	if err != nil {
		t.Fatal(err)
	}
	var id string
	if _, err := pgx.ForEachRow(due, []any{&id}, func() error { r.due[id] = true; return nil }); err != nil {
		t.Fatal(err)
	}
	return r
}

// account is one signed-up account of the machine, calling from a loopback
// address of its own, so the token-use rate limit (10.1 §8.1) counts each
// account's calls apart and alike in a replay.
type account struct {
	email, id, token string
	conn             *connect.Client
	identity         casav1connect.IdentityServiceClient
	households       casav1connect.HouseholdServiceClient
	syncs            casav1connect.SyncServiceClient
	// feeds is the device's sync model per household: what the client
	// holds after its last sync (09.3 §3), moved by the sync action.
	feeds map[string]*syncModel
	// synced is a second device per household that syncs after every
	// step, in checkSyncEqualsRead: a feed row missed by one command shows
	// at the next check.
	synced map[string]*syncModel
}

// syncModel is a device's cache of one household: its cursor and every
// object it holds, by object id.
type syncModel struct {
	cursor  string
	objects map[string]*casav1.Change
}

func (s *syncModel) clone() *syncModel {
	return &syncModel{cursor: s.cursor, objects: maps.Clone(s.objects)}
}

// apply applies one page as a client does under 09.3 §3 "Nilai enum yang
// tidak dikenal klien": an UPSERT of a known kind with its object is held;
// an unknown change kind, UNSPECIFIED included, and every PURGE drop the
// object id whatever its kind; an UPSERT of an unknown kind is skipped.
func (s *syncModel) apply(t *rapid.T, changes []*casav1.Change) {
	for _, c := range changes {
		// 09.3 §3 item 5: the server never sends UNSPECIFIED.
		if c.GetKind() == casav1.ChangeKind_CHANGE_KIND_UNSPECIFIED || c.GetObjectKind() == casav1.ObjectKind_OBJECT_KIND_UNSPECIFIED {
			t.Fatalf("the server sent an UNSPECIFIED change: %v", c)
		}
		switch {
		case c.GetKind() == casav1.ChangeKind_CHANGE_KIND_UPSERT && upsertCarries(c):
			s.objects[c.GetObjectId()] = c
		case c.GetKind() == casav1.ChangeKind_CHANGE_KIND_UPSERT:
			// skipped, not held (09.3 §3 item 3)
		default:
			delete(s.objects, c.GetObjectId())
		}
	}
}

// upsertCarries: c's object is there and of c's kind.
func upsertCarries(c *casav1.Change) bool {
	switch c.GetObjectKind() {
	case casav1.ObjectKind_OBJECT_KIND_HOUSEHOLD:
		return c.GetHousehold() != nil
	case casav1.ObjectKind_OBJECT_KIND_MEMBER:
		return c.GetMember() != nil
	case casav1.ObjectKind_OBJECT_KIND_INVITATION:
		return c.GetInvitation() != nil
	}
	return false
}

// step is one entry of the command log: a command or a resend, with its
// actor, request and normalized answer; or a clock or job step.
type step struct {
	actor     int
	procedure string
	request   proto.Message
	resend    bool
	answer    string
	ok        bool
	advance   time.Duration
	runJobs   bool
}

// resendResult is what the last resend saw (ADR-005).
type resendResult struct {
	first, again          string
	rowsBefore, rowsAfter rowsSnapshot
}

// householdMachine: 2 to 5 accounts, the command log (actor index,
// procedure, request, answer), each member's sync model (cursor and
// objects), the clock.
type householdMachine struct {
	h        *propertyHarness
	accounts []*account
	steps    []step
	names    *names
	// households and invitations are the ids the answers made, for the
	// builders to name; invitationHousehold is each invitation's household.
	households, invitations []string
	invitationHousehold     map[string]string
	tokens                  []string
	resent                  *resendResult
	// Memory of the checkers: every event and membership seen so far.
	seenEvents      map[string]string
	seenMemberships map[string]string
}

// newMachine resets the database and the clock, then signs up n accounts
// in order. order[i] is the account symbol of account i: a replay with
// actors swapped passes the swap, so that the same person carries the same
// symbol in both runs.
func (h *propertyHarness) newMachine(t *rapid.T, n int, order []int) *householdMachine {
	if _, err := h.admin.Exec(t.Context(), "TRUNCATE "+strings.Join(h.tables, ", ")); err != nil {
		t.Fatal(err)
	}
	h.clock.mu.Lock()
	h.clock.now = h.start
	h.clock.mu.Unlock()
	m := &householdMachine{
		h: h, names: newNames(), invitationHousehold: map[string]string{},
		seenEvents: map[string]string{}, seenMemberships: map[string]string{},
	}
	for i := range n {
		// A minute apart, so sign-up order is a difference INV-A13's swap
		// tests, not only the random ids.
		h.clock.advance(time.Minute)
		a := h.signUpAccount(t, fmt.Sprintf("person%d@example.test", i))
		m.names.add(a.email, fmt.Sprintf("account#%d", order[i]))
		m.names.add(a.id, fmt.Sprintf("account#%d", order[i]))
		m.accounts = append(m.accounts, a)
	}
	return m
}

// signUpAccount runs the closed sign-up for email from a new loopback
// address and acknowledges the operator notice, as the client does before
// anything else (FR-FAM-019/AC-1; SEC-34).
func (h *propertyHarness) signUpAccount(t *rapid.T, email string) *account {
	transport := &http.Transport{DialContext: (&net.Dialer{LocalAddr: &net.TCPAddr{IP: h.fixtureAddress()}}).DialContext}
	t.Cleanup(transport.CloseIdleConnections)
	conn := connect.NewClient(connecthttp.NewTransport(&http.Client{Transport: transport}, h.url))
	a := &account{
		email: email, conn: conn, identity: casav1connect.NewIdentityServiceClient(conn),
		households: casav1connect.NewHouseholdServiceClient(conn), syncs: casav1connect.NewSyncServiceClient(conn),
		feeds: map[string]*syncModel{}, synced: map[string]*syncModel{},
	}
	setup, _, err := identity.IssueSetupToken(t.Context(), h.pool, email, h.clock.Now())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := a.identity.StartSignUp(t.Context(), &casav1.StartSignUpRequest{
		RequestContext: identityRC(ios), SetupToken: setup, Email: email,
	}); err != nil {
		t.Fatal(err)
	}
	mails := h.mails()
	_, link, _ := strings.Cut(mails[len(mails)-1].Link, "#")
	res, err := a.identity.CompleteSignUp(t.Context(), &casav1.CompleteSignUpRequest{
		RequestContext: identityRC(ios), VerificationToken: link, Password: passphrase,
	})
	if err != nil {
		t.Fatal(err)
	}
	a.token = res.GetTokens().GetAccessToken()
	if err := h.admin.QueryRow(t.Context(), "SELECT id::text FROM account WHERE email = $1", email).Scan(&a.id); err != nil {
		t.Fatal(err)
	}
	n, err := a.households.GetOperatorNotice(bearer(a.token), &casav1.GetOperatorNoticeRequest{RequestContext: identityRC(ios)})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := a.households.AcknowledgeOperatorNotice(bearer(a.token), &casav1.AcknowledgeOperatorNoticeRequest{
		RequestContext: commandRC(), Variant: n.GetVariant(), Version: n.GetVersion(), TextSha256: n.GetTextSha256(),
	}); err != nil {
		t.Fatal(err)
	}
	return a
}

// signInAll gives every account a new access token once the clock moved.
func (m *householdMachine) signInAll(t *rapid.T) {
	for _, a := range m.accounts {
		res, err := a.identity.SignIn(t.Context(), &casav1.SignInRequest{RequestContext: identityRC(ios), Email: a.email, Password: passphrase})
		if err != nil {
			t.Fatal(err)
		}
		a.token = res.GetTokens().GetAccessToken()
	}
}

// householdMethod is the HouseholdService method of procedure.
func householdMethod(procedure string) protoreflect.MethodDescriptor {
	name := procedure[strings.LastIndex(procedure, "/")+1:]
	return casav1.File_casa_v1_household_proto.Services().ByName("HouseholdService").Methods().ByName(protoreflect.Name(name))
}

// send sends request to procedure as actor: one generic call from the
// contract's method descriptor, so a new command needs only its builder.
func (m *householdMachine) send(t *rapid.T, actor int, procedure string, request proto.Message) (proto.Message, error) {
	method := householdMethod(procedure)
	output, err := protoregistry.GlobalTypes.FindMessageByName(method.Output().FullName())
	if err != nil {
		t.Fatal(err)
	}
	res := output.New().Interface()
	err = m.accounts[actor].conn.CallUnary(bearer(m.accounts[actor].token),
		connect.Spec{StreamType: connect.StreamTypeUnary, Schema: method, Procedure: procedure}, request, res)
	if err != nil {
		var connectErr *connect.Error
		if !errors.As(err, &connectErr) {
			t.Fatalf("%s: %v", procedure, err)
		}
		return nil, err
	}
	return res, nil
}

// run sends s's request, records its normalized answer and appends s to the
// log. A first successful CreateInvitation is followed by its link, as the
// client shares one, so that F02 has a token to use.
func (m *householdMachine) run(t *rapid.T, s step) step {
	res, err := m.send(t, s.actor, s.procedure, s.request)
	s.ok = err == nil
	s.answer = m.names.answer(res, err)
	if s.ok && !s.resend {
		m.learn(t, s, res)
	}
	m.steps = append(m.steps, s)
	return s
}

// learn keeps the ids a successful answer made for the builders, and issues
// the link of a new invitation.
func (m *householdMachine) learn(t *rapid.T, s step, res proto.Message) {
	switch r := res.(type) {
	case *casav1.CreateHouseholdResponse:
		m.households = append(m.households, r.GetHouseholdId())
	case *casav1.CreateInvitationResponse:
		m.invitations = append(m.invitations, r.GetInvitationId())
		m.invitationHousehold[r.GetInvitationId()] = s.request.(*casav1.CreateInvitationRequest).GetHouseholdId()
		link, err := m.accounts[s.actor].households.IssueInvitationLink(bearer(m.accounts[s.actor].token), &casav1.IssueInvitationLinkRequest{
			RequestContext: identityRC(ios), HouseholdId: s.request.(*casav1.CreateInvitationRequest).GetHouseholdId(),
			InvitationId: r.GetInvitationId(),
		})
		if err != nil {
			t.Fatalf("link of a new invitation: %v", err)
		}
		m.names.add(link.GetInvitationToken(), fmt.Sprintf("token#%d", len(m.tokens)))
		m.tokens = append(m.tokens, link.GetInvitationToken())
	}
}

// warmStart draws how far the history starts: from the accounts alone, from
// a household of account 0, from that household with account 1's accepted
// invitation, or with account 1 as its second full adult (06 §3.6 F01, F02,
// F04). The steps are commands of the log like any other, so they are
// resent, replayed and checked; they only make the states the household
// rules govern, two full adults above all (INV-A13), common in a short
// history.
//
// check runs after every step, as after a step of Repeat, so the second
// adult's F04 at depth 3 reaches the inviter's device through its feed.
func (m *householdMachine) warmStart(t *rapid.T, check invariantCheck) {
	m.start(t, rapid.IntRange(0, 3).Draw(t, "warm start"), check)
}

// start runs the first depth steps of warmStart.
func (m *householdMachine) start(t *rapid.T, depth int, check invariantCheck) {
	if depth == 0 {
		return
	}
	m.mustRun(t, check, 0, casav1connect.HouseholdServiceCreateHouseholdProcedure, &casav1.CreateHouseholdRequest{
		RequestContext: commandRC(), Name: "Rumah Kita", TimeZone: "Asia/Jakarta", DisplayName: "Ayu", AgeDeclarationVersion: 1,
	})
	if depth == 1 {
		return
	}
	household := m.households[len(m.households)-1]
	m.mustRun(t, check, 0, casav1connect.HouseholdServiceCreateInvitationProcedure, &casav1.CreateInvitationRequest{
		RequestContext: commandRC(), HouseholdId: household,
	})
	invitation := m.invitations[len(m.invitations)-1]
	m.mustRun(t, check, 1, casav1connect.HouseholdServiceAcceptInvitationProcedure, &casav1.AcceptInvitationRequest{
		RequestContext: commandRC(), DisplayName: "Budi", AgeDeclarationVersion: 1,
		Source: &casav1.InvitationSource{Source: &casav1.InvitationSource_InvitationToken{InvitationToken: m.tokens[len(m.tokens)-1]}},
	})
	if depth == 2 {
		return
	}
	m.mustRun(t, check, 0, casav1connect.HouseholdServiceConfirmInviteeProcedure, &casav1.ConfirmInviteeRequest{
		RequestContext: commandRC(), HouseholdId: household, InvitationId: invitation,
	})
}

// mustRun runs a scripted command step that must succeed, then check
// (nil: none).
func (m *householdMachine) mustRun(t *rapid.T, check invariantCheck, actor int, procedure string, request proto.Message) {
	if s := m.run(t, step{actor: actor, procedure: procedure, request: request}); !s.ok {
		t.Fatalf("scripted %s: %s", procedure, s.answer)
	}
	if check != nil {
		check(t, m)
	}
}

// actions are the machine's actions (plan 1.7 Task 8 "Aksi mesin"): one per
// procedure of commandBuilders, resend, advanceClock, runDueJobs and sync;
// check runs after each.
func (m *householdMachine) actions(check invariantCheck) map[string]func(*rapid.T) {
	actions := map[string]func(*rapid.T){
		"resend":       m.resendAction,
		"advanceClock": m.advanceClock,
		"runDueJobs":   m.runDueJobs,
		"sync":         m.syncAction,
		"": func(t *rapid.T) {
			if check != nil {
				check(t, m)
			}
		},
	}
	for procedure, build := range commandBuilders {
		actions[procedure] = func(t *rapid.T) {
			actor, request := build(m, t)
			m.run(t, step{actor: actor, procedure: procedure, request: request})
		}
	}
	return actions
}

// resendAction sends a command of the log again with the same key (ADR-005).
// Only an accepted command has a first result to return: a refused one
// stored nothing and is a new attempt (ADR-005 "Keputusan").
func (m *householdMachine) resendAction(t *rapid.T) {
	var accepted []int
	for i, s := range m.steps {
		if s.ok && s.procedure != "" {
			accepted = append(accepted, i)
		}
	}
	if len(accepted) == 0 {
		t.Skip("no accepted command yet")
	}
	m.resend(t, m.steps[rapid.SampledFrom(accepted).Draw(t, "resent")])
}

// resend sends first again with its key and keeps what checkResend checks.
func (m *householdMachine) resend(t *rapid.T, first step) {
	before := m.h.rows(t)
	again := m.run(t, step{actor: first.actor, procedure: first.procedure, request: first.request, resend: true})
	m.resent = &resendResult{first: first.answer, again: again.answer, rowsBefore: before, rowsAfter: m.h.rows(t)}
}

// advanceClock moves the clock 1 minute to 8 days: two times in three less
// than a day, so that an invitation (7 days, SEC-28) lives through several
// steps, else past it or close. Then it signs every account in again, as
// the access token lives 15 minutes (SEC-15).
func (m *householdMachine) advanceClock(t *rapid.T) {
	most := 24 * 60
	if rapid.IntRange(0, 2).Draw(t, "days") == 0 {
		most = 8 * 24 * 60
	}
	d := time.Duration(rapid.IntRange(1, most).Draw(t, "minutes")) * time.Minute
	m.h.clock.advance(d)
	m.signInAll(t)
	m.steps = append(m.steps, step{advance: d})
}

// runDueJobs runs the due jobs once, as the runner's loop does each second
// (09.3 §9; plan 1.7 decision 2).
func (m *householdMachine) runDueJobs(t *rapid.T) {
	if _, err := m.h.runner.RunDue(t.Context()); err != nil {
		t.Fatal(err)
	}
	m.steps = append(m.steps, step{runJobs: true})
}

// memberships are the (account id, household id) pairs of every
// membership, read as admin, ordered by account index, then by the
// household's index in m.households.
func (m *householdMachine) memberships(t *rapid.T) [][2]string {
	rows, err := m.h.admin.Query(t.Context(), "SELECT account_id::text, household_id::text FROM membership")
	if err != nil {
		t.Fatal(err)
	}
	pairs, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) ([2]string, error) {
		var p [2]string
		return p, r.Scan(&p[0], &p[1])
	})
	if err != nil {
		t.Fatal(err)
	}
	// In machine terms, never by a database order: member ids are random,
	// so an order that falls back on them would make one drawn index name
	// another account from run to run, and a fixed seed would not fix the
	// history (ES §6).
	slices.SortFunc(pairs, func(a, b [2]string) int {
		return cmp.Or(cmp.Compare(m.accountIndex(a[0]), m.accountIndex(b[0])),
			cmp.Compare(slices.Index(m.households, a[1]), slices.Index(m.households, b[1])))
	})
	return pairs
}

// accountIndex is the machine's index of account id; -1 for none.
func (m *householdMachine) accountIndex(id string) int {
	return slices.IndexFunc(m.accounts, func(a *account) bool { return a.id == id })
}

// syncAction: one random member pulls its feed until has_more is false,
// with a random page size so pages split the feed and the snapshot.
func (m *householdMachine) syncAction(t *rapid.T) {
	pairs := m.memberships(t)
	if len(pairs) == 0 {
		t.Skip("no member yet")
	}
	p := pairs[rapid.IntRange(0, len(pairs)-1).Draw(t, "member")]
	m.syncMember(t, m.accounts[m.accountIndex(p[0])], p[1], rapid.SampledFrom([]uint32{0, 1, 2}).Draw(t, "limit"))
}

// syncMember pulls a's feed of household with page size limit.
func (m *householdMachine) syncMember(t *rapid.T, a *account, household string, limit uint32) {
	feed := a.feeds[household]
	if feed == nil {
		feed = &syncModel{objects: map[string]*casav1.Change{}}
		a.feeds[household] = feed
	}
	m.catchUp(t, a, household, feed, limit)
}

// catchUp syncs feed from its cursor until has_more is false.
func (m *householdMachine) catchUp(t *rapid.T, a *account, household string, feed *syncModel, limit uint32) {
	for range maxPages {
		res, err := a.syncs.Sync(bearer(a.token), syncRequest(household, feed.cursor, limit))
		if err != nil {
			t.Fatalf("sync of %s: %v", m.names.of(household), err)
		}
		if res.GetResyncRequired() {
			t.Fatal("resync_required without a pruned feed")
		}
		feed.apply(t, res.GetChanges())
		feed.cursor = res.GetNextCursor()
		if !res.GetHasMore() {
			return
		}
	}
	t.Fatalf("still has_more after %d pages", maxPages)
}

// pick draws mostly the newest of known, which is the live one, sometimes
// an older one, and rarely the unknown id.
func pick(t *rapid.T, label string, known []string) string {
	switch {
	case len(known) == 0 || rarely(t, label+" unknown"):
		return unknownID
	case rarely(t, label+" older"):
		return rapid.SampledFrom(known).Draw(t, label)
	}
	return known[len(known)-1]
}

// rarely is true one time in eight: the share of drawn requests that break
// one rule, so that most commands can succeed and a history goes deep.
func rarely(t *rapid.T, label string) bool {
	return rapid.IntRange(0, 7).Draw(t, label) == 0
}

// mostly is valid, or rarely one of invalid.
func mostly[V any](t *rapid.T, label string, valid V, invalid ...V) V {
	if rarely(t, label+" invalid") {
		return rapid.SampledFrom(invalid).Draw(t, label)
	}
	return valid
}

// actorIn draws, mostly, a member of household, and sometimes any account,
// so that most household commands come from someone who may send them.
func (m *householdMachine) actorIn(t *rapid.T, household string) int {
	var members []int
	for _, p := range m.memberships(t) {
		if p[1] == household {
			members = append(members, m.accountIndex(p[0]))
		}
	}
	if len(members) == 0 || rarely(t, "outsider") {
		return rapid.IntRange(0, len(m.accounts)-1).Draw(t, "actor")
	}
	return rapid.SampledFrom(members).Draw(t, "actor")
}

// outsider draws, mostly, an account that is a member nowhere.
func (m *householdMachine) outsider(t *rapid.T) int {
	members := map[string]bool{}
	for _, p := range m.memberships(t) {
		members[p[0]] = true
	}
	var outside []int
	for i, a := range m.accounts {
		if !members[a.id] {
			outside = append(outside, i)
		}
	}
	if len(outside) == 0 || rarely(t, "member") {
		return m.anyActor(t)
	}
	return rapid.SampledFrom(outside).Draw(t, "actor")
}

func (m *householdMachine) anyActor(t *rapid.T) int {
	return rapid.IntRange(0, len(m.accounts)-1).Draw(t, "actor")
}

// invitationCommand draws an invitation id, mostly with its own household,
// and an actor for F03, F04 and F05 by the inviter.
func (m *householdMachine) invitationCommand(t *rapid.T) (actor int, household, invitation string) {
	invitation = pick(t, "invitation", m.invitations)
	household, ok := m.invitationHousehold[invitation]
	if !ok || rarely(t, "other household") {
		household = pick(t, "household", m.households)
	}
	return m.actorIn(t, household), household, invitation
}

// commandBuilders has one builder per write command of HouseholdService:
// every method of it with the casa.v1.command option (14 §4.3 "Pembangkit
// perintah"; ruling P11: SignOut and ChangePassword do not touch household
// state). Step 2 widens the enumeration to every service of the contract,
// with an explicit list of the commands left out and why. Values are drawn
// mostly valid, with the refusals of each command in reach.
var commandBuilders = map[string]func(m *householdMachine, t *rapid.T) (actor int, request proto.Message){
	casav1connect.HouseholdServiceCreateHouseholdProcedure: func(m *householdMachine, t *rapid.T) (int, proto.Message) {
		return m.anyActor(t), &casav1.CreateHouseholdRequest{
			RequestContext:        commandRC(),
			Name:                  mostly(t, "name", "Rumah Kita", " "),
			TimeZone:              mostly(t, "time zone", "Asia/Jakarta", "Mars/Olympus"),
			DisplayName:           rapid.SampledFrom([]string{"Ayu", "Budi"}).Draw(t, "display name"),
			AgeDeclarationVersion: mostly[uint32](t, "age declaration", 1, 0, 2),
		}
	},
	casav1connect.HouseholdServiceAcknowledgeOperatorNoticeProcedure: func(m *householdMachine, t *rapid.T) (int, proto.Message) {
		actor := m.anyActor(t)
		a := m.accounts[actor]
		n, err := a.households.GetOperatorNotice(bearer(a.token), &casav1.GetOperatorNoticeRequest{RequestContext: identityRC(ios)})
		if err != nil {
			t.Fatal(err)
		}
		return actor, &casav1.AcknowledgeOperatorNoticeRequest{
			RequestContext: commandRC(), Variant: n.GetVariant(), TextSha256: n.GetTextSha256(),
			Version: n.GetVersion() + mostly[uint32](t, "version ahead", 0, 1),
		}
	},
	casav1connect.HouseholdServiceCreateInvitationProcedure: func(m *householdMachine, t *rapid.T) (int, proto.Message) {
		household := pick(t, "household", m.households)
		return m.actorIn(t, household), &casav1.CreateInvitationRequest{RequestContext: commandRC(), HouseholdId: household}
	},
	casav1connect.HouseholdServiceCancelInvitationProcedure: func(m *householdMachine, t *rapid.T) (int, proto.Message) {
		actor, household, invitation := m.invitationCommand(t)
		return actor, &casav1.CancelInvitationRequest{RequestContext: commandRC(), HouseholdId: household, InvitationId: invitation}
	},
	casav1connect.HouseholdServiceAcceptInvitationProcedure: func(m *householdMachine, t *rapid.T) (int, proto.Message) {
		source := &casav1.InvitationSource{Source: &casav1.InvitationSource_InvitationToken{
			InvitationToken: pick(t, "token", m.tokens),
		}}
		// from_signup: every machine account signed up with a setup token,
		// so this is the refusal path only (TOKEN_REJECTED, SEC-30).
		// ponytail: the success path needs sign-up through an invitation in
		// the machine; a known gap, proved by TestAcceptInvitation.
		if rarely(t, "from sign-up") {
			source = &casav1.InvitationSource{Source: &casav1.InvitationSource_FromSignup{FromSignup: true}}
		}
		return m.outsider(t), &casav1.AcceptInvitationRequest{
			RequestContext: commandRC(), Source: source,
			DisplayName:           rapid.SampledFrom([]string{"Budi", "Citra"}).Draw(t, "display name"),
			AgeDeclarationVersion: mostly[uint32](t, "age declaration", 1, 0),
		}
	},
	casav1connect.HouseholdServiceWithdrawAcceptanceProcedure: func(m *householdMachine, t *rapid.T) (int, proto.Message) {
		return m.anyActor(t), &casav1.WithdrawAcceptanceRequest{RequestContext: commandRC()}
	},
	casav1connect.HouseholdServiceConfirmInviteeProcedure: func(m *householdMachine, t *rapid.T) (int, proto.Message) {
		actor, household, invitation := m.invitationCommand(t)
		return actor, &casav1.ConfirmInviteeRequest{RequestContext: commandRC(), HouseholdId: household, InvitationId: invitation}
	},
	casav1connect.HouseholdServiceRejectInviteeProcedure: func(m *householdMachine, t *rapid.T) (int, proto.Message) {
		actor, household, invitation := m.invitationCommand(t)
		return actor, &casav1.RejectInviteeRequest{RequestContext: commandRC(), HouseholdId: household, InvitationId: invitation}
	},
}

// names maps the values that differ between two runs of one history (ids,
// link tokens, account emails) to symbols in order of birth, so that two
// runs compare equal when they differ only in those values (06 INV-A13
// "bandingkan sesudah memetakan balik").
type names struct {
	symbol, value map[string]string
	ids           int
}

func newNames() *names { return &names{symbol: map[string]string{}, value: map[string]string{}} }

func (n *names) add(value, symbol string) {
	n.symbol[value] = symbol
	n.value[symbol] = value
}

// of is value's symbol, or value itself when it has none.
func (n *names) of(value string) string {
	if s, ok := n.symbol[value]; ok {
		return s
	}
	return value
}

var uuidPattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

// learn gives every id of msg without a symbol the next one.
func (n *names) learn(msg proto.Message) {
	eachString(msg.ProtoReflect(), func(s string) string {
		if _, ok := n.symbol[s]; !ok && uuidPattern.MatchString(s) {
			n.ids++
			n.add(s, fmt.Sprintf("#%d", n.ids))
		}
		return s
	})
}

// normalize is msg in text form with every value replaced by its symbol.
func (n *names) normalize(msg proto.Message) string {
	c := proto.Clone(msg)
	eachString(c.ProtoReflect(), n.of)
	return prototext.MarshalOptions{}.Format(c)
}

// answer is a command's answer, normalized: the response after its new ids
// got symbols, or the error as a client sees it (code, message, details).
func (n *names) answer(res proto.Message, err error) string {
	if err != nil {
		var connectErr *connect.Error
		errors.As(err, &connectErr) // send fails the test on any other error
		answer := fmt.Sprintf("error %v %q", connectErr.Code(), connectErr.Message())
		for _, d := range connectErr.Details() {
			answer += fmt.Sprintf(" %s:%x", d.Type, d.Value)
		}
		return answer
	}
	n.learn(res)
	return n.normalize(res)
}

// mapped is request with every value of from's symbols replaced by n's
// value of the same symbol: the request of one run, sent in the other.
func (n *names) mapped(from *names, request proto.Message) proto.Message {
	c := proto.Clone(request)
	eachString(c.ProtoReflect(), func(s string) string {
		if v, ok := n.value[from.of(s)]; ok {
			return v
		}
		return s
	})
	return c
}

// eachString replaces every string field of msg, nested ones included, by
// f of its value.
func eachString(msg protoreflect.Message, f func(string) string) {
	msg.Range(func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		switch {
		case field.IsList() && field.Kind() == protoreflect.StringKind:
			list := value.List()
			for i := range list.Len() {
				list.Set(i, protoreflect.ValueOfString(f(list.Get(i).String())))
			}
		case field.IsList() && field.Message() != nil:
			list := value.List()
			for i := range list.Len() {
				eachString(list.Get(i).Message(), f)
			}
		case field.Kind() == protoreflect.StringKind:
			msg.Set(field, protoreflect.ValueOfString(f(value.String())))
		case field.Message() != nil && !field.IsMap():
			eachString(value.Message(), f)
		}
		return true
	})
}

// replay sends s of another run's log (names from) in m, with actor
// order[s.actor], and fails when its answer differs after mapping.
func (m *householdMachine) replay(t *rapid.T, i int, s step, from *names, order []int) {
	switch {
	case s.advance > 0:
		m.h.clock.advance(s.advance)
		m.signInAll(t)
	case s.runJobs:
		if _, err := m.h.runner.RunDue(t.Context()); err != nil {
			t.Fatal(err)
		}
	default:
		got := m.run(t, step{actor: order[s.actor], procedure: s.procedure, request: m.names.mapped(from, s.request), resend: s.resend})
		if got.answer != s.answer {
			t.Fatalf("step %d %s by account %d: answer\n%s\nwant, with the actors swapped,\n%s", i, s.procedure, s.actor, got.answer, s.answer)
		}
	}
}

// views are what each account reads, normalized and sorted: its
// memberships, and per household its sync snapshot and live invitation.
func (m *householdMachine) views(t *rapid.T) [][]string {
	views := make([][]string, len(m.accounts))
	for i, a := range m.accounts {
		res, err := a.households.ListMemberships(bearer(a.token), &casav1.ListMembershipsRequest{RequestContext: identityRC(ios)})
		if err != nil {
			t.Fatal(err)
		}
		var lines []string
		for _, ms := range res.GetMemberships() {
			hh := ms.GetHouseholdId()
			for _, p := range ms.GetMembers() {
				lines = append(lines, "member of "+m.names.of(hh)+" "+m.names.normalize(p))
			}
			ms = proto.CloneOf(ms)
			ms.Members = nil
			lines = append(lines, "membership "+m.names.normalize(ms))
			snapshot := &syncModel{objects: map[string]*casav1.Change{}}
			m.catchUp(t, a, hh, snapshot, 0)
			for _, c := range snapshot.objects {
				lines = append(lines, "sync "+m.names.of(hh)+" "+m.names.normalize(c))
			}
			open, err := a.households.GetOpenInvitation(bearer(a.token), &casav1.GetOpenInvitationRequest{RequestContext: identityRC(ios), HouseholdId: hh})
			if err != nil {
				t.Fatal(err)
			}
			lines = append(lines, "open invitation "+m.names.of(hh)+" "+m.names.normalize(open))
		}
		slices.Sort(lines)
		views[i] = lines
	}
	return views
}

func TestHouseholdProperties(t *testing.T) {
	// Ruling P11: HouseholdService only; SignOut and ChangePassword do not
	// touch household state.
	t.Run("every write command of HouseholdService has a builder", func(t *testing.T) {
		methods := casav1.File_casa_v1_household_proto.Services().ByName("HouseholdService").Methods()
		var want []string
		for i := range methods.Len() {
			if proto.HasExtension(methods.Get(i).Options(), casav1.E_Command) {
				want = append(want, "/casa.v1.HouseholdService/"+string(methods.Get(i).Name()))
			}
		}
		slices.Sort(want)
		if got := slices.Sorted(maps.Keys(commandBuilders)); !slices.Equal(got, want) {
			t.Fatalf("builders = %v, want one per command %v", got, want)
		}
	})

	// Each property is rapid.Check over the same machine with one checker
	// after every step (14 §4.3 "Pemeriksa").
	for _, p := range []struct {
		name  string
		check invariantCheck
	}{
		{"INV-A01 every invitation and household transition writes one event and old events never change", checkEvents},
		{"INV-A48 (a) to (c) hold after every random household command", checkMemberships},
		{"a resent command returns its first answer and writes nothing", checkResend},             // ADR-005
		{"sync after every command equals the direct read for every member", checkSyncEqualsRead}, // SEC-06; 14 §4.7
	} {
		t.Run(p.name, func(t *testing.T) {
			h := startPropertyHarness(t)
			rapid.Check(t, func(rt *rapid.T) {
				m := h.newMachine(rt, rapid.IntRange(2, 5).Draw(rt, "accounts"), identity5)
				m.warmStart(rt, p.check)
				rt.Repeat(m.actions(p.check))
			})
		})
	}

	// The named example of the resend exclusion (14 §4.3 "Penyusutan dan
	// benih"), so it runs every time and not only on a lucky seed: F01, the
	// clock past its expiry with the job not run, then F01 again with its
	// key. The resend writes the due expiry first (plan 1.6 decision 12)
	// and nothing else; a write of any other kind is still a failure.
	t.Run("a resend after the invitation's expiry is due writes only the expiry and returns its first answer", func(t *testing.T) {
		h := startPropertyHarness(t)
		// Examples must not draw: a draw turns the example into a skip.
		rapid.MakeFuzz(func(rt *rapid.T) {
			m := h.newMachine(rt, 2, identity5)
			m.start(rt, 1, nil)
			m.mustRun(rt, nil, 0, casav1connect.HouseholdServiceCreateInvitationProcedure, &casav1.CreateInvitationRequest{
				RequestContext: commandRC(), HouseholdId: m.households[0],
			})
			invite := m.steps[len(m.steps)-1]
			m.h.clock.advance(invitationLifetime + time.Minute)
			m.signInAll(rt)
			m.resend(rt, invite)
			r := *m.resent
			if len(r.rowsAfter.events) != len(r.rowsBefore.events)+1 || len(r.rowsAfter.feed) == len(r.rowsBefore.feed) {
				rt.Fatalf("the resend wrote %d events and %d feed rows, want the expiry event and its feed rows",
					len(r.rowsAfter.events)-len(r.rowsBefore.events), len(r.rowsAfter.feed)-len(r.rowsBefore.feed))
			}
			checkResend(rt, m)

			other := r.rowsAfter
			other.events = maps.Clone(other.events)
			other.events["invitation "+m.invitations[0]+" 9"] = "member -"
			if writes := resendWrites(r.rowsBefore, other); len(writes) != 1 {
				rt.Fatalf("a member event written by a resend: writes = %v, want it reported", writes)
			}
			other = r.rowsAfter
			other.feed = maps.Clone(other.feed)
			other.feed[m.accounts[0].id+" 99"] = "household " + m.households[0]
			if writes := resendWrites(r.rowsBefore, other); len(writes) != 1 {
				rt.Fatalf("a household feed row written by a resend: writes = %v, want it reported", writes)
			}
		})(t, nil)
	})

	// The named example of the fan-out (plan 1.7 review focus 4; 09.2 §6
	// step 6): the inviter's device holds the household before F04, so the
	// new member reaches it only through its feed, every run.
	t.Run("F04 reaches the inviter's device through its feed", func(t *testing.T) {
		h := startPropertyHarness(t)
		// Examples must not draw: a draw turns the example into a skip.
		rapid.MakeFuzz(func(rt *rapid.T) {
			m := h.newMachine(rt, 2, identity5)
			m.start(rt, 2, nil)
			m.syncMember(rt, m.accounts[0], m.households[0], 0)
			m.mustRun(rt, checkSyncEqualsRead, 0, casav1connect.HouseholdServiceConfirmInviteeProcedure, &casav1.ConfirmInviteeRequest{
				RequestContext: commandRC(), HouseholdId: m.households[0], InvitationId: m.invitations[0],
			})
		})(t, nil)
	})

	// INV-A13 with N = 2 (14 §4.3 "Simetri"; FR-FAM-001/AC-4, the part of
	// 1.7): the log replayed on new accounts with actors 0 and 1 swapped.
	t.Run("INV-A13 swapping the two full adults over a random sequence swaps every answer and view", func(t *testing.T) {
		h := startPropertyHarness(t)
		rapid.Check(t, func(rt *rapid.T) {
			m := h.newMachine(rt, rapid.IntRange(2, 5).Draw(rt, "accounts"), identity5)
			m.warmStart(rt, nil)
			rt.Repeat(m.actions(nil))
			views := m.views(rt)
			swap := []int{1, 0, 2, 3, 4}
			r := h.newMachine(rt, len(m.accounts), swap)
			for i, s := range m.steps {
				r.replay(rt, i, s, m.names, swap)
			}
			swapped := r.views(rt)
			for i := range m.accounts {
				if !slices.Equal(swapped[swap[i]], views[i]) {
					rt.Fatalf("account %d reads\n%s\nwant, with the actors swapped,\n%s", i,
						strings.Join(swapped[swap[i]], "\n"), strings.Join(views[i], "\n"))
				}
			}
		})
	})
}

// identity5 is the order of a run as generated.
var identity5 = []int{0, 1, 2, 3, 4}
