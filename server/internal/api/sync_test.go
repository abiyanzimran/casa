package api_test

import (
	"bytes"
	"crypto/rand"
	"io"
	"maps"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"testing/fstest"
	"time"

	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/abiyanzimran/casa/server/internal/api"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/read"
	casasync "github.com/abiyanzimran/casa/server/internal/sync"
)

// newCursorKey is a new random cursor key, as the test server makes one
// (plan 1.7 decision 4; 14 §7.8).
func newCursorKey(t *testing.T) casasync.CursorKey {
	t.Helper()
	var b [32]byte
	_, _ = rand.Read(b[:]) // crypto/rand never fails (Go 1.24)
	key, err := casasync.NewCursorKey(b[:])
	if err != nil {
		t.Fatal(err)
	}
	return key
}

// syncRequest is a foreground Sync of household from cursor; limit 0 is the
// server's default.
func syncRequest(householdID, cursor string, limit uint32) *casav1.SyncRequest {
	return &casav1.SyncRequest{
		RequestContext: identityRC(ios), HouseholdId: householdID, Cursor: cursor,
		Purpose: casav1.SyncPurpose_SYNC_PURPOSE_FOREGROUND, Limit: limit,
	}
}

func (h *householdHarness) syncPage(m member, householdID, cursor string, limit uint32) (*casav1.SyncResponse, error) {
	return h.syncs.Sync(bearer(m.token), syncRequest(householdID, cursor, limit))
}

func (h *householdHarness) mustSync(t *testing.T, m member, householdID, cursor string) *casav1.SyncResponse {
	t.Helper()
	res, err := h.syncPage(m, householdID, cursor, 0)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

// rawSync sends req as connect JSON over plain HTTP and returns the status
// and body exactly as they arrive, to compare two answers byte for byte.
func (h *householdHarness) rawSync(t *testing.T, m member, req *casav1.SyncRequest) (int, []byte) {
	t.Helper()
	body, err := protojson.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	httpReq, err := http.NewRequestWithContext(t.Context(), http.MethodPost, h.url+casav1connect.SyncServiceSyncProcedure, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+m.token)
	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = res.Body.Close() }()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	return res.StatusCode, got
}

// syncCache is a device's cache: each live object synced, by kind and id,
// as the deterministic wire form of its change (09.3 §3).
type syncCache map[string]string

func (c syncCache) apply(t *testing.T, changes []*casav1.Change) {
	t.Helper()
	for _, ch := range changes {
		key := ch.GetObjectKind().String() + " " + ch.GetObjectId()
		switch ch.GetKind() {
		case casav1.ChangeKind_CHANGE_KIND_UPSERT:
			wire, err := proto.MarshalOptions{Deterministic: true}.Marshal(ch)
			if err != nil {
				t.Fatal(err)
			}
			c[key] = string(wire)
		case casav1.ChangeKind_CHANGE_KIND_PURGE:
			if ch.GetObject() != nil {
				t.Fatalf("purge of %s carries an object", key)
			}
			delete(c, key)
		default:
			t.Fatalf("change of kind %v", ch.GetKind())
		}
	}
}

// catchUp syncs m from cursor with limit until has_more is false, applying
// every page to c, and returns the last next_cursor; a cursor that never
// ends fails the test instead of looping.
func (h *householdHarness) catchUp(t *testing.T, m member, householdID, cursor string, limit uint32, c syncCache) string {
	t.Helper()
	for range 50 {
		res, err := h.syncPage(m, householdID, cursor, limit)
		if err != nil {
			t.Fatal(err)
		}
		if res.GetResyncRequired() {
			t.Fatal("unexpected resync_required")
		}
		c.apply(t, res.GetChanges())
		if !res.GetHasMore() {
			return res.GetNextCursor()
		}
		cursor = res.GetNextCursor()
	}
	t.Fatal("still has_more after 50 pages")
	return ""
}

// directCache is what m reads of the household online, in sync's shape: the
// household and its members from ListMemberships, with newHousehold's time
// zone and currency, and the live invitation from GetOpenInvitation, whose
// inviter, not in that answer, is inviter.
func (h *householdHarness) directCache(t *testing.T, m member, householdID, inviter string) syncCache {
	t.Helper()
	c := syncCache{}
	upsert := func(kind casav1.ObjectKind, id string, change *casav1.Change) {
		change.Kind, change.ObjectKind, change.ObjectId = casav1.ChangeKind_CHANGE_KIND_UPSERT, kind, id
		c.apply(t, []*casav1.Change{change})
	}
	for _, ms := range h.memberships(t, m) {
		if ms.GetHouseholdId() != householdID {
			continue
		}
		upsert(casav1.ObjectKind_OBJECT_KIND_HOUSEHOLD, householdID, &casav1.Change{Object: &casav1.Change_Household{Household: &casav1.Household{
			HouseholdId: householdID, Name: ms.GetHouseholdName(), TimeZone: "Asia/Jakarta", Currency: "IDR",
		}}})
		for _, p := range ms.GetMembers() {
			upsert(casav1.ObjectKind_OBJECT_KIND_MEMBER, p.GetMemberId(), &casav1.Change{Object: &casav1.Change_Member{Member: p}})
		}
	}
	if inv := h.mustOpenInvitation(t, m, householdID); inv.GetInvitationId() != "" {
		upsert(casav1.ObjectKind_OBJECT_KIND_INVITATION, inv.GetInvitationId(), &casav1.Change{Object: &casav1.Change_Invitation{Invitation: &casav1.Invitation{
			InvitationId: inv.GetInvitationId(), State: inv.GetState(), ExpiresAt: inv.GetExpiresAt(), InviterMemberId: inviter,
		}}})
	}
	return c
}

// changeOf is the one change of kind among changes.
func changeOf(t *testing.T, changes []*casav1.Change, kind casav1.ObjectKind) *casav1.Change {
	t.Helper()
	var found []*casav1.Change
	for _, ch := range changes {
		if ch.GetObjectKind() == kind {
			found = append(found, ch)
		}
	}
	if len(found) != 1 {
		t.Fatalf("%d changes of %v in %v, want 1", len(found), kind, changes)
	}
	return found[0]
}

// 09.3 §3; ADR-017 items 3, 6, 7, 8; ADR-018 item 2; plan 1.7 decisions 9
// to 14.
func TestSync(t *testing.T) {
	t.Run("SEC-05 a cursor of another member, of another household, or changed in one character is not found like a household that does not exist", func(t *testing.T) {
		f := startTwoHouseholds(t)
		h := f.h
		b := h.signUpMember(t, "b@example.test")
		h.addMember(t, f.x, b.account)
		c := h.signUpMember(t, "c@example.test")
		h.mustCreate(t, c)
		own := h.mustSync(t, f.a, f.x, "").GetNextCursor()
		otherHousehold := h.mustSync(t, f.a, f.y, "").GetNextCursor()
		otherMember := h.mustSync(t, b, f.x, "").GetNextCursor()

		missing := syncRequest(uuid4(), "", 0)
		_, wantErr := h.syncs.Sync(bearer(f.a.token), missing)
		wantHiddenNotFound(t, wantErr)
		want := rejectionOf(t, wantErr)
		wantStatus, wantBody := h.rawSync(t, f.a, missing)

		type attempt struct {
			name       string
			caller     member
			householdX bool // else a household that does not exist
			cursor     string
		}
		attempts := []attempt{
			{"another member's cursor", f.a, true, otherMember},
			{"another household's cursor", f.a, true, otherHousehold},
			{"a truncated cursor", f.a, true, own[:len(own)-1]},
			{"a cursor that is not base64url", f.a, true, "not a cursor!"},
			{"a non-member without a cursor", c, true, ""},
			{"a non-member with a member's cursor", c, true, own},
			{"a household that does not exist with a valid cursor", f.a, false, own},
		}
		for i := range len(own) {
			changed := []byte(own)
			changed[i] = map[bool]byte{true: 'B', false: 'A'}[changed[i] == 'A']
			attempts = append(attempts, attempt{"the cursor changed at character " + strconv.Itoa(i), f.a, true, string(changed)})
		}
		before := h.fingerprint(t)
		for _, a := range attempts {
			householdID := uuid4()
			if a.householdX {
				householdID = f.x
			}
			req := syncRequest(householdID, a.cursor, 0)
			_, err := h.syncs.Sync(bearer(a.caller.token), req)
			if err == nil || !rejectionOf(t, err).equal(want) {
				t.Errorf("%s (%q): %v, want exactly %v", a.name, a.cursor, err, wantErr)
				continue
			}
			if status, body := h.rawSync(t, a.caller, req); status != wantStatus || !bytes.Equal(body, wantBody) {
				t.Errorf("%s: HTTP %d %s, want %d %s", a.name, status, body, wantStatus, wantBody)
			}
		}
		if h.fingerprint(t) != before {
			t.Fatal("a refused sync changed a table")
		}
		// The member's own cursor still works.
		if _, err := h.syncPage(f.a, f.x, own, 0); err != nil {
			t.Fatalf("own cursor: %v", err)
		}
	})

	t.Run("INV-A48 (c) an invitee before F04 reads nothing through sync", func(t *testing.T) {
		h := startHousehold(t)
		inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
		_, err := h.syncPage(b, inv.household, "", 0)
		_, missing := h.syncPage(b, uuid4(), "", 0)
		wantHiddenNotFound(t, err)
		if !rejectionOf(t, err).equal(rejectionOf(t, missing)) {
			t.Fatalf("invitee before F04: %v, a household that does not exist: %v", err, missing)
		}
	})

	t.Run("both members reach the same household state through sync", func(t *testing.T) {
		h := startHousehold(t)
		inv := h.invited(t, "a@example.test")
		a := inv.inviter
		// Limit 1: every answer is one page, so paging runs through the API.
		byA := syncCache{}
		cursorA := h.catchUp(t, a, inv.household, "", 1, byA)
		b := h.signUpMember(t, "b@example.test")
		h.mustAccept(t, b, inv.token)
		cursorA = h.catchUp(t, a, inv.household, cursorA, 1, byA)
		if want := h.directCache(t, a, inv.household, inv.inviterMember); !maps.Equal(byA, want) {
			t.Fatalf("after the acceptance the inviter's sync holds %q, the direct read %q", byA, want)
		}
		h.mustConfirm(t, a, inv.household, inv.id)
		h.catchUp(t, a, inv.household, cursorA, 1, byA) // a through the feed
		byB := syncCache{}
		h.catchUp(t, b, inv.household, "", 1, byB) // b through the snapshot
		direct := h.directCache(t, a, inv.household, "")
		if len(direct) != 3 || !maps.Equal(byA, direct) || !maps.Equal(byB, direct) || !maps.Equal(h.directCache(t, b, inv.household, ""), direct) {
			t.Fatalf("after F04: a's sync %q, b's sync %q, direct read %q; want all equal, the household and two members", byA, byB, direct)
		}
	})

	// FR-FAM-019/AC-3 and AC-10, the sync part (plan 1.7 decision 12); the
	// screen is 1.10.
	t.Run("a new notice version shows in the next sync and acknowledging clears it", func(t *testing.T) {
		h := startHousehold(t)
		v1 := &fstest.MapFile{Data: publishedV1(t)}
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": v1})
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a).GetHouseholdId()
		res := h.mustSync(t, a, hh, "")
		if res.GetOperatorNoticeUnacknowledged() {
			t.Fatal("unacknowledged with version 1 acknowledged")
		}
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": v1, "couple-2.json": noticeFile(t, "Judul", household.NoticeParagraph{Text: "Teks baru."})})
		res = h.mustSync(t, a, hh, res.GetNextCursor())
		if !res.GetOperatorNoticeUnacknowledged() {
			t.Fatal("version 2 does not show in the next sync")
		}
		// Issued on first sight, through sync too (09.3 §1 "Terbit saat
		// pertama terlihat"): one trail line per version.
		issued := slices.DeleteFunc(h.trail(t, a), func(line string) bool { return !strings.HasPrefix(line, "notice_version ") })
		if len(issued) != 2 || !strings.HasPrefix(issued[1], "notice_version couple:2:") {
			t.Fatalf("notice_version lines %q, want versions 1 and 2", issued)
		}
		h.ack(t, a)
		if h.mustSync(t, a, hh, res.GetNextCursor()).GetOperatorNoticeUnacknowledged() {
			t.Fatal("still unacknowledged after the acknowledgement")
		}
	})

	t.Run("sync is a read: never held by NOTICE_UNREAD and needs no idempotency key", func(t *testing.T) {
		method := casav1.File_casa_v1_sync_proto.Services().ByName("SyncService").Methods().ByName("Sync")
		if proto.HasExtension(method.Options(), casav1.E_Command) {
			t.Fatal("Sync carries the casa.v1.command option")
		}
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a).GetHouseholdId()
		h.reopen(testOperatorName, fstest.MapFS{
			"couple-1.json": {Data: publishedV1(t)},
			"couple-2.json": noticeFile(t, "Judul", household.NoticeParagraph{Text: "Teks baru."}),
		})
		if !h.held(t, a) {
			t.Fatal("version 2 does not hold commands")
		}
		req := syncRequest(hh, "", 0)
		if req.GetRequestContext().GetIdempotencyKey() != "" {
			t.Fatal("the request carries an idempotency key")
		}
		res, err := h.syncs.Sync(bearer(a.token), req)
		if err != nil {
			t.Fatalf("Sync of a held account: %v, want it served", err)
		}
		if !res.GetOperatorNoticeUnacknowledged() || len(res.GetChanges()) != 2 {
			t.Fatalf("Sync = %v, want the household and its member, and the notice unacknowledged", res)
		}
		if n := h.count(t, "SELECT count(*) FROM idempotency_record WHERE command = $1", casav1connect.SyncServiceSyncProcedure); n != 0 {
			t.Fatalf("%d idempotency records of Sync, want 0", n)
		}
	})

	t.Run("the invitee's email and display name never travel in sync", func(t *testing.T) {
		forbidden := regexp.MustCompile(`(?i)email|invitee`)
		seen := map[protoreflect.FullName]bool{}
		var walk func(protoreflect.MessageDescriptor)
		walk = func(m protoreflect.MessageDescriptor) {
			if seen[m.FullName()] {
				return
			}
			seen[m.FullName()] = true
			for i := range m.Fields().Len() {
				f := m.Fields().Get(i)
				if forbidden.MatchString(string(f.Name())) {
					t.Errorf("%s has field %s", m.FullName(), f.Name())
				}
				if f.Message() != nil {
					walk(f.Message())
				}
			}
		}
		walk((&casav1.SyncResponse{}).ProtoReflect().Descriptor())

		h := startHousehold(t)
		email := "budi.rahasia@example.test"
		inv := h.invited(t, "a@example.test")
		cursor := h.mustSync(t, inv.inviter, inv.household, "").GetNextCursor()
		h.mustAccept(t, h.signUpMember(t, email), inv.token)
		want := &casav1.Invitation{
			InvitationId: inv.id, State: casav1.InvitationState_INVITATION_STATE_ACCEPTED,
			ExpiresAt: timestamppb.New(inv.expiresAt), InviterMemberId: inv.inviterMember,
		}
		for path, res := range map[string]*casav1.SyncResponse{
			"feed": h.mustSync(t, inv.inviter, inv.household, cursor), "snapshot": h.mustSync(t, inv.inviter, inv.household, ""),
		} {
			if got := changeOf(t, res.GetChanges(), casav1.ObjectKind_OBJECT_KIND_INVITATION).GetInvitation(); !proto.Equal(got, want) {
				t.Errorf("%s: invitation %v, want %v", path, got, want)
			}
			wire, err := proto.Marshal(res)
			if err != nil {
				t.Fatal(err)
			}
			if bytes.Contains(bytes.ToLower(wire), []byte(email)) || bytes.Contains(wire, []byte("Budi")) {
				t.Errorf("%s: the answer carries the invitee's email or display name", path)
			}
		}
	})

	t.Run("a pruned feed answers resync_required with no changes and no cursor", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a).GetHouseholdId()
		cursor := h.mustSync(t, a, hh, "").GetNextCursor()
		inv := h.mustInvite(t, a, hh)
		// The pruning job is step 7 (plan 1.7 decision 5); the superuser
		// removes the rows the cursor still needs.
		if _, err := h.admin.Exec(t.Context(), "DELETE FROM reader_feed WHERE household_id = $1", hh); err != nil {
			t.Fatal(err)
		}
		if res := h.mustSync(t, a, hh, cursor); !proto.Equal(res, &casav1.SyncResponse{ResyncRequired: true}) {
			t.Fatalf("Sync = %v, want only resync_required", res)
		}
		// Starting again with an empty cursor gets the snapshot.
		res := h.mustSync(t, a, hh, "")
		if got := changeOf(t, res.GetChanges(), casav1.ObjectKind_OBJECT_KIND_INVITATION); got.GetObjectId() != inv.GetInvitationId() {
			t.Fatalf("snapshot invitation %v, want %s", got, inv.GetInvitationId())
		}
	})

	t.Run("a sync without purpose is refused", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a).GetHouseholdId()
		for purpose, rule := range map[casav1.SyncPurpose]string{0: "purpose enum.not_in", 7: "purpose enum.defined_only"} {
			req := syncRequest(hh, "", 0)
			req.Purpose = purpose
			_, err := h.syncs.Sync(bearer(a.token), req)
			wantCode(t, err, connect.CodeInvalidArgument)
			if got := violations(t, err); !slices.Equal(got, []string{rule}) {
				t.Errorf("purpose %d: violations %q, want [%s]", purpose, got, rule)
			}
		}
	})

	// Plan 1.7 decision 14: 0 is 100, at most 500 [Usulan].
	t.Run("a page limit above 500 is refused", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a).GetHouseholdId()
		_, err := h.syncPage(a, hh, "", 501)
		wantCode(t, err, connect.CodeInvalidArgument)
		if got := violations(t, err); !slices.Equal(got, []string{"limit uint32.lte"}) {
			t.Fatalf("violations %q, want [limit uint32.lte]", got)
		}
		if _, err := h.syncPage(a, hh, "", 500); err != nil {
			t.Fatalf("limit 500: %v", err)
		}
	})

	// 10.1 §10 note 1.7 Q3 condition 1, the Sync half of TestInvitationExpiryJob's
	// "the job and the lazy path write the same expiry".
	t.Run("the job and the lazy path give the same sync answer", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		aHousehold := h.mustCreate(t, a).GetHouseholdId()
		byJob := h.mustInvite(t, a, aHousehold).GetInvitationId()
		c := h.signUpMember(t, "c@example.test")
		cHousehold := h.mustCreate(t, c).GetHouseholdId()
		byLazyPath := h.mustInvite(t, c, cHousehold).GetInvitationId()
		cursorA := h.mustSync(t, a, aHousehold, "").GetNextCursor()
		cursorC := h.mustSync(t, c, cHousehold, "").GetNextCursor()
		h.clock.advance(invitationLifetime + 3*time.Hour)
		h.signInAgain(t, &a, "a@example.test")
		h.signInAgain(t, &c, "c@example.test")
		next := h.mustInvite(t, c, cHousehold).GetInvitationId() // X3: F01 writes the due expiry first
		if ran := h.runDue(t); ran != 2 {
			t.Fatalf("ran %d jobs, want 2 (one writes, one finds the invitation closed)", ran)
		}
		fromJob := h.mustSync(t, a, aHousehold, cursorA)
		fromLazyPath := h.mustSync(t, c, cHousehold, cursorC)
		job, lazy := fromJob.GetChanges(), fromLazyPath.GetChanges()
		if len(job) != 1 || len(lazy) != 2 || job[0].GetObjectId() != byJob || lazy[0].GetObjectId() != byLazyPath ||
			lazy[1].GetObjectId() != next || lazy[1].GetKind() != casav1.ChangeKind_CHANGE_KIND_UPSERT {
			t.Fatalf("job path %v, lazy path %v; want the expired invitation, and on the lazy path then the new one", job, lazy)
		}
		masked := func(ch *casav1.Change) *casav1.Change {
			m, _ := proto.Clone(ch).(*casav1.Change)
			m.ObjectId = ""
			return m
		}
		wantPurge := &casav1.Change{Kind: casav1.ChangeKind_CHANGE_KIND_PURGE, ObjectKind: casav1.ObjectKind_OBJECT_KIND_INVITATION}
		if !proto.Equal(masked(job[0]), wantPurge) || !proto.Equal(masked(lazy[0]), wantPurge) ||
			fromJob.GetHasMore() || fromLazyPath.GetHasMore() || fromJob.GetResyncRequired() || fromLazyPath.GetResyncRequired() {
			t.Fatalf("job path %v, lazy path %v; want the same purge apart from the object id", fromJob, fromLazyPath)
		}
	})
}

// syncChange fails closed on a change it cannot put in the contract, so the
// handler answers internal instead of sending an object a client misreads
// (09.3 §3; final review of PR B).
func TestSyncChange(t *testing.T) {
	id := pgtype.UUID{Bytes: [16]byte{1}, Valid: true}
	for name, c := range map[string]casasync.Change{
		"an object kind the contract does not name": {Kind: casasync.Purge, Ref: read.ObjectRef{Kind: "account", ID: id}},
		"an invitation state that is not live": {Kind: casasync.Upsert, Ref: read.ObjectRef{Kind: read.KindInvitation, ID: id},
			Object: read.Object{Invitation: &read.Invitation{ID: id, State: "expired"}}},
		"an upsert without an object": {Kind: casasync.Upsert, Ref: read.ObjectRef{Kind: read.KindMember, ID: id}},
		"a change kind that is neither upsert nor purge": {Kind: 0, Ref: read.ObjectRef{Kind: read.KindMember, ID: id},
			Object: read.Object{Member: &read.Member{MemberID: id}}},
	} {
		if got, err := api.SyncChange(c); err == nil {
			t.Errorf("%s: %v, want an error", name, got)
		}
	}
	got, err := api.SyncChange(casasync.Change{Kind: casasync.Purge, Ref: read.ObjectRef{Kind: read.KindInvitation, ID: id}})
	want := &casav1.Change{Kind: casav1.ChangeKind_CHANGE_KIND_PURGE, ObjectKind: casav1.ObjectKind_OBJECT_KIND_INVITATION, ObjectId: uuidText(id)}
	if err != nil || !proto.Equal(got, want) {
		t.Fatalf("purge = %v, %v; want %v", got, err, want)
	}
}
