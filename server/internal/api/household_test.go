package api_test

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/abiyanzimran/casa/server/internal/api"
	"github.com/abiyanzimran/casa/server/internal/database"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/household"
)

// householdHarness is the household service on the identity harness's
// server, so its accounts sign up through the real closed sign-up.
type householdHarness struct {
	*identityHarness
	households casav1connect.HouseholdServiceClient
	service    *household.Service
}

func startHousehold(t *testing.T) *householdHarness {
	t.Helper()
	h := &householdHarness{}
	h.identityHarness = startIdentity(t, func(server *connect.Server, id *identityHarness) {
		h.service = household.New(id.pool, id.clock.Now)
		api.RegisterHouseholdService(server, id.pool, h.service, id.service)
	})
	h.households = casav1connect.NewHouseholdServiceClient(h.probe)
	return h
}

// member is a signed-up account: its access token, and its id as text.
type member struct {
	token, account string
}

func (h *householdHarness) signUpMember(t *testing.T, email string) member {
	t.Helper()
	tokens := h.signUp(t, email, "a long enough passphrase")
	var account string
	if err := h.admin.QueryRow(t.Context(), "SELECT id::text FROM account WHERE email = $1", email).Scan(&account); err != nil {
		t.Fatal(err)
	}
	return member{token: tokens.GetAccessToken(), account: account}
}

// newHousehold is a valid request: name, time zone, display name and the
// current age declaration version (08 §7.1).
func newHousehold() *casav1.CreateHouseholdRequest {
	return &casav1.CreateHouseholdRequest{
		RequestContext: commandRC(), Name: "Rumah Kita", TimeZone: "Asia/Jakarta",
		DisplayName: "Ayu", AgeDeclarationVersion: 1,
	}
}

func (h *householdHarness) create(m member, req *casav1.CreateHouseholdRequest) (*casav1.CreateHouseholdResponse, error) {
	return h.households.CreateHousehold(bearer(m.token), req)
}

func (h *householdHarness) mustCreate(t *testing.T, m member) *casav1.CreateHouseholdResponse {
	t.Helper()
	res, err := h.create(m, newHousehold())
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func (h *householdHarness) memberships(t *testing.T, m member) []*casav1.Membership {
	t.Helper()
	res, err := h.households.ListMemberships(bearer(m.token), &casav1.ListMembershipsRequest{RequestContext: identityRC(ios)})
	if err != nil {
		t.Fatal(err)
	}
	return res.GetMemberships()
}

// wantNothingWritten checks that no household row, consent record or
// idempotency record exists.
func (h *householdHarness) wantNothingWritten(t *testing.T) {
	t.Helper()
	for _, table := range []string{"household", "membership", "object_event", "consent_record", "idempotency_record"} {
		if n := h.count(t, "SELECT count(*) FROM "+table); n != 0 {
			t.Fatalf("%d rows in %s, want 0", n, table)
		}
	}
}

// dropOneHouseholdIndex removes membership_one_household from this test's
// database only, so an account can be a member of more than one household
// (FR-FAM-020/AC-2, AC-4). The index is the pilot's product rule "one
// household per account", not a limit of the model (09.2 §5.1, note on
// decision 20 of the 1.6 plan).
func dropOneHouseholdIndex(t *testing.T, admin *pgxpool.Pool) {
	t.Helper()
	if _, err := admin.Exec(t.Context(), "DROP INDEX membership_one_household"); err != nil {
		t.Fatal(err)
	}
}

func TestCreateHousehold(t *testing.T) {
	notAllowed := casav1.ErrorCode_ERROR_CODE_HOUSEHOLD_CREATION_NOT_ALLOWED

	// FR-FAM-001/AC-1, the part of 1.6: a new household with a new id and
	// its creator as household owner. The start day of AC-1 is step 3
	// (FAM.md FR-FAM-001 note of 14 November 2026, Q5 (b)).
	t.Run("the creator becomes a household owner of a new household", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		res := h.mustCreate(t, a)
		var name, zone string
		var createdAt time.Time
		var version int64
		if err := h.admin.QueryRow(t.Context(), "SELECT name, time_zone, created_at, version FROM household WHERE id = $1",
			res.GetHouseholdId()).Scan(&name, &zone, &createdAt, &version); err != nil {
			t.Fatal(err)
		}
		if name != "Rumah Kita" || zone != "Asia/Jakarta" || !createdAt.Equal(h.clock.Now()) || version != 0 {
			t.Fatalf("household = %q %q %v %d, want Rumah Kita, Asia/Jakarta, %v, 0", name, zone, createdAt, version, h.clock.Now())
		}
		var role, permissions, account, display string
		var joinedAt time.Time
		if err := h.admin.QueryRow(t.Context(), `
			SELECT role, permission_set, account_id::text, display_name, joined_at FROM membership
			 WHERE household_id = $1 AND member_id = $2`,
			res.GetHouseholdId(), res.GetMemberId()).Scan(&role, &permissions, &account, &display, &joinedAt); err != nil {
			t.Fatal(err)
		}
		if role != "household_owner" || permissions != "two_equal_adults" || account != a.account || display != "Ayu" || !joinedAt.Equal(h.clock.Now()) {
			t.Fatalf("membership = %s %s %s %q %v", role, permissions, account, display, joinedAt)
		}
		if n := h.count(t, "SELECT count(*) FROM membership"); n != 1 {
			t.Fatalf("%d memberships, want 1", n)
		}
		// SEC-10: the one definition of the household's full adult members.
		var adults []string
		scope := database.Scope{HouseholdID: uuidOf(t, res.GetHouseholdId()), AccountID: uuidOf(t, a.account)}
		if err := database.Read(t.Context(), h.pool, scope, func(tx pgx.Tx) error {
			ids, err := h.service.FullAdultMembers(t.Context(), tx)
			for _, id := range ids {
				adults = append(adults, uuidText(id))
			}
			return err
		}); err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(adults, []string{res.GetMemberId()}) {
			t.Fatalf("full adult members = %v, want [%s]", adults, res.GetMemberId())
		}
	})

	t.Run("FR-FAM-001/AC-2 the concepts are written as separate rows and no relationship row", func(t *testing.T) {
		h := startHousehold(t)
		res := h.mustCreate(t, h.signUpMember(t, "a@example.test"))
		var authorityMember, requestType string
		var validFrom time.Time
		if err := h.admin.QueryRow(t.Context(), `
			SELECT member_id::text, request_type, valid_from FROM approval_authority
			 WHERE household_id = $1 AND limit_amount IS NULL AND currency IS NULL AND valid_until IS NULL`,
			res.GetHouseholdId()).Scan(&authorityMember, &requestType, &validFrom); err != nil {
			t.Fatal(err)
		}
		if authorityMember != res.GetMemberId() || requestType != "all" || !validFrom.Equal(h.clock.Now()) {
			t.Fatalf("approval authority = %s %s %v", authorityMember, requestType, validFrom)
		}
		var preset, mode string
		if err := h.admin.QueryRow(t.Context(), `
			SELECT preset_id, agreement_mode FROM household_rule
			 WHERE household_id = $1 AND version = 1 AND request_threshold IS NULL AND currency IS NULL`,
			res.GetHouseholdId()).Scan(&preset, &mode); err != nil {
			t.Fatal(err)
		}
		if preset != "two_equal_adults" || mode != "all" {
			t.Fatalf("household rule = %s %s", preset, mode)
		}
		for table, want := range map[string]int{"approval_authority": 1, "household_rule": 1, "relationship": 0} {
			if n := h.count(t, "SELECT count(*) FROM "+table); n != want {
				t.Fatalf("%d rows in %s, want %d", n, table, want)
			}
		}
	})

	// FR-FAM-001/AC-3, the contract part; the screen is 1.10.
	t.Run("no field but name, display name, time zone and age declaration is required", func(t *testing.T) {
		fields := (&casav1.CreateHouseholdRequest{}).ProtoReflect().Descriptor().Fields()
		var names []string
		for i := range fields.Len() {
			names = append(names, string(fields.Get(i).Name()))
		}
		want := []string{"request_context", "name", "time_zone", "display_name", "age_declaration_version"}
		if !slices.Equal(names, want) {
			t.Fatalf("fields = %v, want %v", names, want)
		}
	})

	t.Run("FR-FAM-001/AC-7 a household made with only a name is IDR and the Indonesian market", func(t *testing.T) {
		h := startHousehold(t)
		res := h.mustCreate(t, h.signUpMember(t, "a@example.test"))
		var currency, market string
		if err := h.admin.QueryRow(t.Context(), "SELECT currency, income_source_market FROM household WHERE id = $1",
			res.GetHouseholdId()).Scan(&currency, &market); err != nil {
			t.Fatal(err)
		}
		if currency != "IDR" || market != "ID" {
			t.Fatalf("currency, market = %s, %s; want IDR, ID", currency, market)
		}
	})

	t.Run("FR-FAM-001/AC-5 an account that is already a member cannot create a second household", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		h.mustCreate(t, a)
		_, err := h.create(a, newHousehold())
		wantErrorCode(t, err, connect.CodePermissionDenied, notAllowed)
		if n := h.count(t, "SELECT count(*) FROM household"); n != 1 {
			t.Fatalf("%d households, want 1", n)
		}
	})

	t.Run("FR-FAM-001/AC-10 an account made through an invitation cannot create a household", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		// Until Task 8 builds sign-up through an invitation (SEC-158 (a)).
		if _, err := h.admin.Exec(t.Context(), "UPDATE account SET signup_via = 'invitation' WHERE id = $1", a.account); err != nil {
			t.Fatal(err)
		}
		_, err := h.create(a, newHousehold())
		wantErrorCode(t, err, connect.CodePermissionDenied, notAllowed)
		h.wantNothingWritten(t)
	})

	// FR-FAM-004/AC-12, the part for creating (SEC-80); accepting is Task 6
	// and the screen 1.10.
	t.Run("creating without the age declaration is refused without effect, and recorded once when given", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		for _, version := range []uint32{0, 2} {
			req := newHousehold()
			req.AgeDeclarationVersion = version
			_, err := h.create(a, req)
			wantErrorCode(t, err, connect.CodeFailedPrecondition, casav1.ErrorCode_ERROR_CODE_AGE_DECLARATION_REQUIRED)
			h.wantNothingWritten(t)
		}
		h.mustCreate(t, a)
		var kind string
		var version int32
		var recordedAt time.Time
		if err := h.admin.QueryRow(t.Context(), "SELECT kind, text_version, recorded_at FROM consent_record WHERE account_id = $1",
			a.account).Scan(&kind, &version, &recordedAt); err != nil {
			t.Fatal(err)
		}
		if kind != "adult_age" || version != 1 || !recordedAt.Equal(h.clock.Now()) {
			t.Fatalf("consent record = %s %d %v", kind, version, recordedAt)
		}
		// SEC-80: the statement only, never a birth date.
		if n := h.count(t, "SELECT count(*) FROM information_schema.columns WHERE table_schema = 'public' AND column_name ILIKE '%birth%'"); n != 0 {
			t.Fatalf("%d birth columns, want 0", n)
		}
	})

	t.Run("an age declaration is asked once per account", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		if _, err := h.admin.Exec(t.Context(),
			"INSERT INTO consent_record (account_id, kind, text_version, recorded_at) VALUES ($1, 'adult_age', 1, now())", a.account); err != nil {
			t.Fatal(err)
		}
		req := newHousehold()
		req.AgeDeclarationVersion = 0
		if _, err := h.create(a, req); err != nil {
			t.Fatal(err)
		}
		if n := h.count(t, "SELECT count(*) FROM consent_record"); n != 1 {
			t.Fatalf("%d consent records, want 1", n)
		}
	})

	t.Run("the same command sent twice creates one household", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		req := newHousehold()
		first, err := h.create(a, req)
		if err != nil {
			t.Fatal(err)
		}
		again, err := h.create(a, req)
		if err != nil {
			t.Fatal(err)
		}
		if again.GetHouseholdId() != first.GetHouseholdId() || again.GetMemberId() != first.GetMemberId() {
			t.Fatalf("resend = %v, want %v", again, first)
		}
		if n := h.count(t, "SELECT count(*) FROM household"); n != 1 {
			t.Fatalf("%d households, want 1", n)
		}
	})

	t.Run("name and time zone violations name their rule", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		req := newHousehold()
		req.Name = "   "
		req.DisplayName = "Ayu\u202E"
		req.TimeZone = "Mars/Olympus"
		_, err := h.create(a, req)
		wantCode(t, err, connect.CodeInvalidArgument)
		got := violations(t, err)
		want := []string{"name name.visible_length", "time_zone time_zone.iana", "display_name name.characters"}
		if !slices.Equal(got, want) {
			t.Fatalf("violations = %v, want %v", got, want)
		}
		for _, zone := range []string{"Local", "../../etc/passwd"} {
			req := newHousehold()
			req.TimeZone = zone
			_, err := h.create(a, req)
			if got := violations(t, err); !slices.Equal(got, []string{"time_zone time_zone.iana"}) {
				t.Fatalf("time zone %q: violations = %v", zone, got)
			}
		}
		h.wantNothingWritten(t)
	})

	// INV-A01; 06 §2.5: the creation event, seq 0 (09.2 §5.1).
	t.Run("creating writes the household event with seq 0 by its member", func(t *testing.T) {
		h := startHousehold(t)
		req := newHousehold()
		res, err := h.create(h.signUpMember(t, "a@example.test"), req)
		if err != nil {
			t.Fatal(err)
		}
		var got string
		if err := h.admin.QueryRow(t.Context(), `
			SELECT concat_ws(' ', object_kind, object_id, seq, actor_kind, actor, coalesce(rule, '-'), command,
			                 coalesce(from_state, '-'), to_state, idempotency_key, (server_time = $2)::text)
			  FROM object_event WHERE household_id = $1`, res.GetHouseholdId(), h.clock.Now()).Scan(&got); err != nil {
			t.Fatal(err)
		}
		want := fmt.Sprintf("household %s 0 member %s - create_household - created %s true",
			res.GetHouseholdId(), res.GetMemberId(), strings.ToLower(req.GetRequestContext().GetIdempotencyKey()))
		if got != want {
			t.Fatalf("event = %q, want %q", got, want)
		}
	})

	// 09.2 §5.1, note on decision 20: the race MemberAnywhere cannot see is
	// caught by membership_one_household.
	t.Run("two creations by one account at once: exactly one household", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		waitBlocked, release := holdTables(t, h.db, "membership")
		errs := make([]error, 2)
		var wg sync.WaitGroup
		for i := range errs {
			wg.Go(func() { _, errs[i] = h.create(a, newHousehold()) })
		}
		waitBlocked(2)
		release()
		wg.Wait()
		var ok, already int
		for _, err := range errs {
			switch {
			case err == nil:
				ok++
			case slices.Equal(errorCodes(t, err), []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_ALREADY_HOUSEHOLD_MEMBER}) &&
				connect.CodeOf(err) == connect.CodeFailedPrecondition:
				already++
			default:
				t.Fatalf("error = %v, want ALREADY_HOUSEHOLD_MEMBER", err)
			}
		}
		if ok != 1 || already != 1 {
			t.Fatalf("%d created and %d refused, want 1 and 1", ok, already)
		}
		if n := h.count(t, "SELECT count(*) FROM membership"); n != 1 {
			t.Fatalf("%d memberships, want 1", n)
		}
		if n := h.count(t, "SELECT count(*) FROM household"); n != 1 {
			t.Fatalf("%d households, want 1", n)
		}
	})
}

func TestListMemberships(t *testing.T) {
	t.Run("lists the account's households with their members' display names only", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		res := h.mustCreate(t, a)
		b := h.signUpMember(t, "b@example.test")
		// Until F04 (Task 7) adds a second member through the API.
		var bMember string
		if err := h.admin.QueryRow(t.Context(), `
			INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
			VALUES ($1, $2, 'Budi', 'household_owner', 'two_equal_adults', $3) RETURNING member_id::text`,
			res.GetHouseholdId(), b.account, h.clock.Now().Add(time.Minute)).Scan(&bMember); err != nil {
			t.Fatal(err)
		}
		for _, m := range []member{a, b} {
			got := h.memberships(t, m)
			if len(got) != 1 {
				t.Fatalf("%d memberships, want 1", len(got))
			}
			own := map[member]string{a: res.GetMemberId(), b: bMember}[m]
			g := got[0]
			if g.GetHouseholdId() != res.GetHouseholdId() || g.GetHouseholdName() != "Rumah Kita" || g.GetMemberId() != own {
				t.Fatalf("membership = %v", g)
			}
			var people []string
			for _, p := range g.GetMembers() {
				people = append(people, p.GetMemberId()+" "+p.GetDisplayName())
			}
			if want := []string{res.GetMemberId() + " Ayu", bMember + " Budi"}; !slices.Equal(people, want) {
				t.Fatalf("members = %v, want %v", people, want)
			}
			// Display names only: no email, no account id.
			encoded := protojson.Format(&casav1.ListMembershipsResponse{Memberships: got})
			for _, secret := range []string{a.account, b.account, "a@example.test", "b@example.test"} {
				if strings.Contains(encoded, secret) {
					t.Fatalf("response carries %q", secret)
				}
			}
		}
		wantFields(t, (&casav1.Member{}).ProtoReflect().Descriptor(), "member_id", "display_name")
	})

	t.Run("an account without a household gets an empty list", func(t *testing.T) {
		h := startHousehold(t)
		h.mustCreate(t, h.signUpMember(t, "a@example.test"))
		if got := h.memberships(t, h.signUpMember(t, "b@example.test")); len(got) != 0 {
			t.Fatalf("memberships = %v, want none", got)
		}
	})

	// 20 households of 20 members each at most [Usulan] (plan 1.6 "Batasan
	// global", SQL).
	t.Run("the list is bounded", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		res := h.mustCreate(t, a)
		dropOneHouseholdIndex(t, h.admin)
		later := h.clock.Now().Add(time.Minute)
		// 20 more members in a's household and 20 more households of a,
		// all joined after a's own: the first listed is a's household.
		for _, s := range []struct {
			sql  string
			args []any
		}{
			{`INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
			  SELECT $1, gen_random_uuid(), 'M' || i, 'household_owner', 'two_equal_adults', $2
			    FROM generate_series(1, 20) i`, []any{res.GetHouseholdId(), later}},
			{`WITH h AS (
				INSERT INTO household (id, name, time_zone, created_at)
				SELECT gen_random_uuid(), 'H' || i, 'Asia/Jakarta', $1 FROM generate_series(1, 20) i RETURNING id)
			  INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
			  SELECT id, $2, 'Ayu', 'household_owner', 'two_equal_adults', $1 FROM h`, []any{later, a.account}},
		} {
			if _, err := h.admin.Exec(t.Context(), s.sql, s.args...); err != nil {
				t.Fatal(err)
			}
		}
		got := h.memberships(t, a)
		if len(got) != 20 {
			t.Fatalf("%d memberships, want 20", len(got))
		}
		if got[0].GetHouseholdId() != res.GetHouseholdId() || len(got[0].GetMembers()) != 20 {
			t.Fatalf("first membership %s has %d members, want %s with 20", got[0].GetHouseholdId(), len(got[0].GetMembers()), res.GetHouseholdId())
		}
	})
}

// wantFields checks that message has exactly the named fields.
func wantFields(t *testing.T, message protoreflect.MessageDescriptor, want ...string) {
	t.Helper()
	var got []string
	for i := range message.Fields().Len() {
		got = append(got, string(message.Fields().Get(i).Name()))
	}
	if !slices.Equal(got, want) {
		t.Fatalf("%s fields = %v, want %v", message.FullName(), got, want)
	}
}

// uuidOf parses a uuid in text form.
func uuidOf(t *testing.T, s string) pgtype.UUID {
	t.Helper()
	var id pgtype.UUID
	if err := id.Scan(s); err != nil {
		t.Fatal(err)
	}
	return id
}

func uuidText(id pgtype.UUID) string {
	v, _ := id.Value() // a valid uuid always has a text form
	s, _ := v.(string)
	return s
}
