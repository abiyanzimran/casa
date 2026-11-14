package api_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"io/fs"
	"slices"
	"strings"
	"sync"
	"testing"
	"testing/fstest"
	"time"

	"connectrpc.com/connect/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"

	"github.com/abiyanzimran/casa/server/internal/api"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/household"
)

var (
	noticeUnread = []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_NOTICE_UNREAD}
	stateChanged = []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_STATE_CHANGED}
)

// publishedV1 is couple-1.json as published.
func publishedV1(t *testing.T) []byte {
	t.Helper()
	b, err := fs.ReadFile(household.PublishedNotices, "couple-1.json")
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// noticeFile is a test version with title and paragraphs.
func noticeFile(t *testing.T, title string, paragraphs ...household.NoticeParagraph) *fstest.MapFile {
	t.Helper()
	type paragraph struct {
		Text    string `json:"text"`
		Feature string `json:"feature"`
	}
	file := struct {
		Title      string      `json:"title"`
		Paragraphs []paragraph `json:"paragraphs"`
	}{Title: title}
	for _, p := range paragraphs {
		file.Paragraphs = append(file.Paragraphs, paragraph{p.Text, p.Feature})
	}
	b, err := json.Marshal(file)
	if err != nil {
		t.Fatal(err)
	}
	return &fstest.MapFile{Data: b}
}

// held reports whether a command of m is held by NOTICE_UNREAD, through the
// boundary: CreateHousehold, whose other refusals (already a member) also
// show it passed the hold.
func (h *householdHarness) held(t *testing.T, m member) bool {
	t.Helper()
	_, err := h.create(m, newHousehold())
	if err == nil {
		return false
	}
	codes := errorCodes(t, err)
	switch {
	case slices.Equal(codes, []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_HOUSEHOLD_CREATION_NOT_ALLOWED}):
		return false
	case connect.CodeOf(err) == connect.CodeFailedPrecondition && slices.Equal(codes, noticeUnread):
		return true
	}
	t.Fatalf("CreateHousehold: %v (codes %v), want success or NOTICE_UNREAD", err, codes)
	return false
}

// trail is the account's security trail lines of the notice kinds, as
// "kind detail", oldest first.
func (h *householdHarness) trail(t *testing.T, m member) []string {
	t.Helper()
	rows, err := h.admin.Query(t.Context(), `
		SELECT kind || ' ' || coalesce(detail, '-') FROM security_event
		 WHERE account_id = $1 AND kind IN ('notice_version', 'notice_acknowledged')
		 ORDER BY occurred_at, kind DESC`, m.account)
	if err != nil {
		t.Fatal(err)
	}
	var lines []string
	for rows.Next() {
		var line string
		if err := rows.Scan(&line); err != nil {
			t.Fatal(err)
		}
		lines = append(lines, line)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	return lines
}

func TestOperatorNotice(t *testing.T) {
	// FR-FAM-019/AC-1, the server part: the full text of the current
	// version; reading it in LY-05 before the membership is active is 1.10.
	t.Run("a new account reads the full current version, unacknowledged", func(t *testing.T) {
		h := startHousehold(t)
		n := h.notice(t, h.signUpUnread(t, "a@example.test"))
		var want struct {
			Title      string `json:"title"`
			Paragraphs []struct {
				Text    string `json:"text"`
				Feature string `json:"feature"`
			} `json:"paragraphs"`
		}
		if err := json.Unmarshal(publishedV1(t), &want); err != nil {
			t.Fatal(err)
		}
		if n.GetVariant() != "couple" || n.GetVersion() != 1 || n.GetTitle() != want.Title || n.GetAcknowledged() ||
			len(n.GetTextSha256()) != 32 || len(n.GetParagraphs()) != 5 || len(want.Paragraphs) != 5 {
			t.Fatalf("notice = %v, want couple 1 %q with 5 paragraphs, unacknowledged, a 32-byte hash", n, want.Title)
		}
		for i, p := range n.GetParagraphs() {
			text := strings.ReplaceAll(want.Paragraphs[i].Text, "{pengelola}", testOperatorName)
			if p.GetText() != text || p.GetFeature() != "" || strings.ContainsAny(p.GetText(), "{}") {
				t.Errorf("paragraph %d = %q %q, want %q without a tag", i+1, p.GetText(), p.GetFeature(), text)
			}
		}
		// FR-FAM-019/AC-11, the text part: the [Klien web] paragraph says a
		// changed web app leaves no line in the security history (10.1 §11.3).
		if !strings.HasSuffix(n.GetParagraphs()[3].GetText(), " Perubahan seperti itu tidak tercatat di riwayat keamanan akunmu.") {
			t.Fatalf("paragraph 4 = %q, want the [Klien web] paragraph with its third sentence", n.GetParagraphs()[3].GetText())
		}
		if !strings.Contains(n.GetParagraphs()[0].GetText(), "dikelola oleh "+testOperatorName+",") {
			t.Fatalf("paragraph 1 = %q, want the operator name filled in", n.GetParagraphs()[0].GetText())
		}
	})

	// FR-FAM-019/AC-2, the server part: the date and version are stored as
	// the account's own record. The member reading them ("terbaca olehnya")
	// lands in step 6 (the security trail reader, FR-FAM-022) and the 1.10
	// screen (lead ruling, 1.6 Task 4 review); the contract carries no date.
	t.Run("an acknowledgement stores its date and version as the account's own record", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		n := h.notice(t, a)
		h.clock.advance(time.Minute)
		if err := h.acknowledge(a, n); err != nil {
			t.Fatal(err)
		}
		var variant string
		var version int32
		var hash []byte
		var issued, acknowledged time.Time
		if err := h.admin.QueryRow(t.Context(), `
			SELECT variant, version, text_sha256, issued_at, acknowledged_at FROM operator_notice_ack WHERE account_id = $1`,
			a.account).Scan(&variant, &version, &hash, &issued, &acknowledged); err != nil {
			t.Fatal(err)
		}
		if variant != "couple" || version != 1 || !bytes.Equal(hash, n.GetTextSha256()) ||
			!issued.Equal(h.clock.Now().Add(-time.Minute)) || !acknowledged.Equal(h.clock.Now()) {
			t.Fatalf("record = %s %d %x issued %v acknowledged %v, want couple 1 %x, %v, %v",
				variant, version, hash, issued, acknowledged, n.GetTextSha256(), h.clock.Now().Add(-time.Minute), h.clock.Now())
		}
		if !h.notice(t, a).GetAcknowledged() {
			t.Fatal("the account does not read its acknowledgement")
		}
		// The record is the account's own: another account's scope sees
		// none of it (SEC-08; 09.2 §3 note of 14 November 2026).
		b := h.signUpUnread(t, "b@example.test")
		if h.notice(t, b).GetAcknowledged() {
			t.Fatal("another account reads the acknowledgement as its own")
		}
		if h.held(t, a) || !h.held(t, b) {
			t.Fatal("want a released and b held")
		}
	})

	// FR-FAM-019/AC-3, the server part: a new version holds until it is
	// acknowledged and the old record stays. The device learning it on the
	// next sync is proved in 1.7 (sync), the reading in 1.10 (screen).
	t.Run("a new version holds until it is read, and the old acknowledgement stays", func(t *testing.T) {
		h := startHousehold(t)
		v1 := &fstest.MapFile{Data: publishedV1(t)}
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": v1})
		a := h.signUpMember(t, "a@example.test")
		var before string
		if err := h.admin.QueryRow(t.Context(), "SELECT row(o.*)::text FROM operator_notice_ack o WHERE account_id = $1", a.account).Scan(&before); err != nil {
			t.Fatal(err)
		}
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": v1, "couple-2.json": noticeFile(t, "Judul", household.NoticeParagraph{Text: "Teks baru."})})
		if !h.held(t, a) {
			t.Fatal("version 2 does not hold")
		}
		n := h.notice(t, a)
		if n.GetVersion() != 2 || n.GetAcknowledged() {
			t.Fatalf("notice = version %d acknowledged %v, want 2, false", n.GetVersion(), n.GetAcknowledged())
		}
		if err := h.acknowledge(a, n); err != nil {
			t.Fatal(err)
		}
		if h.held(t, a) {
			t.Fatal("still held after acknowledging version 2")
		}
		var after string
		if err := h.admin.QueryRow(t.Context(), "SELECT row(o.*)::text FROM operator_notice_ack o WHERE account_id = $1 AND version = 1", a.account).Scan(&after); err != nil {
			t.Fatal(err)
		}
		if after != before {
			t.Fatalf("version 1 record = %s, want it unchanged: %s", after, before)
		}
	})

	t.Run("SEC-34 only the latest version releases the hold", func(t *testing.T) {
		h := startHousehold(t)
		v1 := &fstest.MapFile{Data: publishedV1(t)}
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": v1})
		a := h.signUpUnread(t, "a@example.test")
		old := h.notice(t, a)
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": v1, "couple-2.json": noticeFile(t, "Judul", household.NoticeParagraph{Text: "Teks baru."})})
		before := h.trail(t, a)
		err := h.acknowledge(a, old)
		wantCode(t, err, connect.CodeFailedPrecondition)
		if got := errorCodes(t, err); !slices.Equal(got, stateChanged) {
			t.Fatalf("error codes = %v, want %v", got, stateChanged)
		}
		// Nothing written: not the acknowledgement, not version 2's issue.
		if n := h.count(t, "SELECT count(*) FROM operator_notice_ack WHERE acknowledged_at IS NOT NULL OR version = 2"); n != 0 {
			t.Fatalf("%d rows acknowledged or of version 2, want 0", n)
		}
		if got := h.trail(t, a); !slices.Equal(got, before) {
			t.Fatalf("trail = %q, want unchanged %q", got, before)
		}
		if !h.held(t, a) {
			t.Fatal("not held after acknowledging an old version")
		}
		h.ack(t, a)
		if h.held(t, a) {
			t.Fatal("still held after acknowledging the latest version")
		}
	})

	t.Run("SEC-34 a changed operator name holds again and the old hash is refused", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		old := h.notice(t, a)
		h.reopen("Nama Lain", household.PublishedNotices)
		if !h.held(t, a) {
			t.Fatal("a changed operator name does not hold")
		}
		renamed := h.notice(t, a)
		if renamed.GetVersion() != old.GetVersion() || bytes.Equal(renamed.GetTextSha256(), old.GetTextSha256()) ||
			!strings.Contains(renamed.GetParagraphs()[0].GetText(), "Nama Lain") {
			t.Fatalf("notice after the rename = %v, want the same number with a new hash and the new name", renamed)
		}
		err := h.acknowledge(a, old)
		if got := errorCodes(t, err); !slices.Equal(got, stateChanged) {
			t.Fatalf("acknowledging the old hash: %v (codes %v), want %v", err, got, stateChanged)
		}
		if !h.held(t, a) {
			t.Fatal("released by the old hash")
		}
		if err := h.acknowledge(a, renamed); err != nil {
			t.Fatal(err)
		}
		if h.held(t, a) {
			t.Fatal("still held after acknowledging the new hash")
		}
	})

	// 10.1 note before §11.3 (A→B→A); 09.3 §1 "Mekanisme": the identity in
	// force is compared with the one last issued to the account, not with
	// any it acknowledged before.
	t.Run("SEC-34 renaming the operator back is a new issue that holds and writes its trail line", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		first := h.notice(t, a)
		h.clock.advance(time.Minute)
		h.reopen("Nama Lain", household.PublishedNotices)
		h.ack(t, a)
		h.clock.advance(time.Minute)
		h.reopen(testOperatorName, household.PublishedNotices)
		if !h.held(t, a) {
			t.Fatal("renaming the operator back does not hold")
		}
		back := h.notice(t, a)
		if back.GetAcknowledged() || !bytes.Equal(back.GetTextSha256(), first.GetTextSha256()) {
			t.Fatalf("notice after renaming back = acknowledged %v, hash %x; want unacknowledged, the first hash %x",
				back.GetAcknowledged(), back.GetTextSha256(), first.GetTextSha256())
		}
		if err := h.acknowledge(a, back); err != nil {
			t.Fatal(err)
		}
		if h.held(t, a) {
			t.Fatal("still held after acknowledging")
		}
		if n := h.count(t, "SELECT count(*) FROM operator_notice_ack WHERE account_id = $1", a.account); n != 3 {
			t.Fatalf("%d issue rows, want 3", n)
		}
		hexA := hex.EncodeToString(first.GetTextSha256())[:8]
		var hexB string
		for _, line := range h.trail(t, a) {
			if detail, ok := strings.CutPrefix(line, "notice_version couple:1:"); ok && detail != hexA {
				hexB = detail
			}
		}
		want := []string{
			"notice_version couple:1:" + hexA, "notice_acknowledged couple:1:" + hexA,
			"notice_version couple:1:" + hexB, "notice_acknowledged couple:1:" + hexB,
			"notice_version couple:1:" + hexA, "notice_acknowledged couple:1:" + hexA,
		}
		if got := h.trail(t, a); hexB == "" || !slices.Equal(got, want) {
			t.Fatalf("trail = %q, want %q", got, want)
		}
	})

	// 09.3 §1 note of 14 November 2026: earlier_version_acknowledged follows
	// the account's earlier issues, a rename back (A→B→A) included.
	t.Run("earlier_version_acknowledged follows the earlier issues, A, B, A included", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		c := h.signUpUnread(t, "c@example.test")
		earlier := func(m member) bool { return h.notice(t, m).GetEarlierVersionAcknowledged() }
		if earlier(a) || earlier(c) {
			t.Fatal("first sight: earlier_version_acknowledged is true, want false")
		}
		h.ack(t, a)
		if n := h.notice(t, a); !n.GetAcknowledged() || n.GetEarlierVersionAcknowledged() {
			t.Fatalf("acknowledged now = %v earlier = %v, want true and false", n.GetAcknowledged(), n.GetEarlierVersionAcknowledged())
		}
		h.clock.advance(time.Minute)
		h.reopen("Nama Lain", household.PublishedNotices) // B, a new version
		if n := h.notice(t, a); n.GetAcknowledged() || !n.GetEarlierVersionAcknowledged() {
			t.Fatalf("new version: acknowledged %v earlier %v, want false and true", n.GetAcknowledged(), n.GetEarlierVersionAcknowledged())
		}
		if earlier(c) { // its only earlier issue, A, was never acknowledged
			t.Fatal("an account whose only earlier issue was never acknowledged reads earlier true")
		}
		h.clock.advance(time.Minute)
		h.reopen(testOperatorName, household.PublishedNotices) // A again
		if n := h.notice(t, a); n.GetAcknowledged() || !n.GetEarlierVersionAcknowledged() {
			t.Fatalf("A after B: acknowledged %v earlier %v, want false and true", n.GetAcknowledged(), n.GetEarlierVersionAcknowledged())
		}
		if !earlier(a) { // a repeated read of the same issue keeps it
			t.Fatal("a repeated read lost earlier_version_acknowledged")
		}
		if earlier(c) { // A (unacknowledged), B (unacknowledged), A: still none acknowledged
			t.Fatal("c reads earlier true with no acknowledged issue")
		}
	})

	t.Run("SEC-34 every version writes one line in the account's security trail", func(t *testing.T) {
		h := startHousehold(t)
		v1 := &fstest.MapFile{Data: publishedV1(t)}
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": v1})
		a := h.signUpUnread(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		bTrail := h.trail(t, b)
		n1 := h.notice(t, a)
		h.notice(t, a)
		h.held(t, a)
		if got, want := h.trail(t, a), []string{"notice_version couple:1:" + hex.EncodeToString(n1.GetTextSha256())[:8]}; !slices.Equal(got, want) {
			t.Fatalf("trail after reads = %q, want %q", got, want)
		}
		for range 2 { // a second acknowledgement changes nothing and writes no line
			if err := h.acknowledge(a, n1); err != nil {
				t.Fatal(err)
			}
		}
		h.clock.advance(time.Minute)
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": v1, "couple-2.json": noticeFile(t, "Judul", household.NoticeParagraph{Text: "Teks baru."})})
		h.held(t, a)
		n2 := h.notice(t, a)
		hex1, hex2 := hex.EncodeToString(n1.GetTextSha256())[:8], hex.EncodeToString(n2.GetTextSha256())[:8]
		want := []string{"notice_version couple:1:" + hex1, "notice_acknowledged couple:1:" + hex1, "notice_version couple:2:" + hex2}
		if got := h.trail(t, a); !slices.Equal(got, want) {
			t.Fatalf("trail = %q, want %q", got, want)
		}
		// The trail of b, which read version 1 before, gets version 2 only
		// when b meets it.
		if got := h.trail(t, b); !slices.Equal(got, bTrail) {
			t.Fatalf("trail of b = %q, want unchanged %q", got, bTrail)
		}
	})

	t.Run("SEC-34 two first sights at once write one row and one trail line", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		waitBlocked, release := holdTables(t, h.db, "operator_notice_ack")
		var wg sync.WaitGroup
		errs := make([]error, 2)
		for i := range errs {
			wg.Go(func() {
				_, errs[i] = h.households.GetOperatorNotice(bearer(a.token), &casav1.GetOperatorNoticeRequest{RequestContext: identityRC(ios)})
			})
		}
		waitBlocked(2)
		release()
		wg.Wait()
		for _, err := range errs {
			wantOK(t, err)
		}
		if n := h.count(t, "SELECT count(*) FROM operator_notice_ack WHERE account_id = $1", a.account); n != 1 {
			t.Fatalf("%d rows, want 1", n)
		}
		if got := h.trail(t, a); len(got) != 1 {
			t.Fatalf("trail = %q, want one line", got)
		}
	})

	// FR-FAM-019/AC-8, the server part: the tag is sent per paragraph; the
	// client's "Berlaku saat fitur ini tersedia" is 1.10.
	t.Run("a paragraph with a feature is sent with its tag", func(t *testing.T) {
		h := startHousehold(t)
		h.reopen(testOperatorName, fstest.MapFS{"couple-1.json": noticeFile(t, "Judul", household.NoticeParagraph{Text: "Selalu."}, household.NoticeParagraph{Text: "Suasana hati.", Feature: "care"})})
		got := h.notice(t, h.signUpUnread(t, "a@example.test")).GetParagraphs()
		if len(got) != 2 || got[0].GetFeature() != "" || got[1].GetFeature() != "care" || got[1].GetText() != "Suasana hati." {
			t.Fatalf("paragraphs = %v, want the second tagged care", got)
		}
	})

	t.Run("a version number beyond int32 is a field violation", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		n := h.notice(t, a)
		_, err := h.households.AcknowledgeOperatorNotice(bearer(a.token), &casav1.AcknowledgeOperatorNoticeRequest{
			RequestContext: commandRC(), Variant: "couple", Version: 2147483648, TextSha256: n.GetTextSha256(),
		})
		wantCode(t, err, connect.CodeInvalidArgument)
		if got := violationPaths(t, err); !slices.Equal(got, []string{"version"}) {
			t.Fatalf("violation paths = %q, want [version]", got)
		}
	})
}

// heldMethods are the procedures NOTICE_UNREAD holds, from the contract:
// every command that is not notice_exempt, plus noticeHeldProcedures
// (09.3 §1; plan 1.6 decision 7).
func heldMethods(t *testing.T) (held, exempt []protoreflect.MethodDescriptor) {
	t.Helper()
	extra := api.NoticeHeldProcedures()
	for _, file := range casaFiles(t) {
		for i := range file.Services().Len() {
			methods := file.Services().Get(i).Methods()
			for j := range methods.Len() {
				m := methods.Get(j)
				command, isCommand := proto.GetExtension(m.Options(), casav1.E_Command).(*casav1.CommandOptions)
				isCommand = isCommand && proto.HasExtension(m.Options(), casav1.E_Command)
				switch {
				case isCommand && command.GetNoticeExempt():
					exempt = append(exempt, m)
				case isCommand || extra[procedureOf(m)]:
					delete(extra, procedureOf(m))
					held = append(held, m)
				}
			}
		}
	}
	if len(extra) != 0 {
		t.Fatalf("noticeHeldProcedures names procedures outside the contract: %v", extra)
	}
	return held, exempt
}

// noticeFreeReads are the online reads outside IdentityService, which
// NOTICE_UNREAD never holds (09.3 §1; plan 1.6 decision 16).
var noticeFreeReads = []string{
	casav1connect.HouseholdServiceListMembershipsProcedure,
	casav1connect.HouseholdServiceGetOperatorNoticeProcedure,
	casav1connect.HouseholdServicePreviewInvitationProcedure,
	casav1connect.HouseholdServiceGetOpenInvitationProcedure,
	casav1connect.HouseholdServiceGetOwnAcceptanceProcedure,
	casav1connect.SyncServiceSyncProcedure, // plan 1.7 decision 12
}

// methodOf is the contract's method of procedure.
func methodOf(t *testing.T, procedure string) protoreflect.MethodDescriptor {
	t.Helper()
	service, method, _ := strings.Cut(strings.TrimPrefix(procedure, "/"), "/")
	d, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(service + "." + method))
	m, ok := d.(protoreflect.MethodDescriptor)
	if err != nil || !ok {
		t.Fatalf("%s is not a method of the contract (%v)", procedure, err)
	}
	return m
}

func procedureOf(m protoreflect.MethodDescriptor) string {
	return "/" + string(m.Parent().FullName()) + "/" + string(m.Name())
}

// callWithContext calls m with a request holding only rc, as m's token.
func (h *householdHarness) callWithContext(m protoreflect.MethodDescriptor, token string, rc *casav1.RequestContext) error {
	req := dynamicpb.NewMessage(m.Input())
	req.Set(m.Input().Fields().ByName("request_context"), protoreflect.ValueOfMessage(rc.ProtoReflect()))
	return h.probe.CallUnary(bearer(token), connect.Spec{Procedure: procedureOf(m)}, req, dynamicpb.NewMessage(m.Output()))
}

// fingerprint is the content of every table, to show a request changed
// nothing.
func (h *householdHarness) fingerprint(t *testing.T) string {
	t.Helper()
	var tables []string
	rows, err := h.admin.Query(t.Context(), "SELECT tablename FROM pg_tables WHERE schemaname = 'public' ORDER BY tablename")
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			t.Fatal(err)
		}
		tables = append(tables, table)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	var all strings.Builder
	for _, table := range tables {
		var content string
		if err := h.admin.QueryRow(t.Context(),
			"SELECT coalesce(string_agg(row(t.*)::text, ',' ORDER BY row(t.*)::text), '') FROM "+table+" t").Scan(&content); err != nil {
			t.Fatal(err)
		}
		all.WriteString(table + ":" + content + "\n")
	}
	return all.String()
}

func TestNoticeHold(t *testing.T) {
	t.Run("SEC-34 every write that is not notice_exempt is held", func(t *testing.T) {
		held, _ := heldMethods(t)
		if !slices.ContainsFunc(held, func(m protoreflect.MethodDescriptor) bool { return m.Name() == "CreateHousehold" }) {
			t.Fatalf("held methods %v do not include CreateHousehold", held)
		}
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		h.notice(t, a) // the first sight is written here, not by the held requests
		before := h.fingerprint(t)
		for _, m := range held {
			err := h.callWithContext(m, a.token, commandRC())
			if connect.CodeOf(err) != connect.CodeFailedPrecondition || !slices.Equal(errorCodes(t, err), noticeUnread) {
				t.Errorf("%s: %v, want FailedPrecondition NOTICE_UNREAD", m.FullName(), err)
			}
		}
		if h.fingerprint(t) != before {
			t.Fatal("a held request changed a table")
		}
		if n := h.count(t, "SELECT count(*) FROM idempotency_record"); n != 0 {
			t.Fatalf("%d idempotency records, want 0", n)
		}
	})

	t.Run("notice_exempt commands and identity RPCs are never held", func(t *testing.T) {
		_, exempt := heldMethods(t)
		for _, name := range []string{"SignOut", "AcknowledgeOperatorNotice", "WithdrawAcceptance"} {
			if !slices.ContainsFunc(exempt, func(m protoreflect.MethodDescriptor) bool { return string(m.Name()) == name }) {
				t.Fatalf("exempt methods %v do not include %s", exempt, name)
			}
		}
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		for _, m := range exempt {
			tokens, err := h.signIn("a@example.test", "a long enough passphrase") // SignOut ends the session it is sent in
			if err != nil {
				t.Fatal(err)
			}
			err = h.callWithContext(m, tokens.GetAccessToken(), commandRC())
			if err != nil && (slices.Contains(errorCodes(t, err), casav1.ErrorCode_ERROR_CODE_NOTICE_UNREAD) || connect.CodeOf(err) == connect.CodeUnauthenticated) {
				t.Errorf("%s: %v, want it past the hold", m.FullName(), err)
			}
		}
		tokens, err := h.signIn("a@example.test", "a long enough passphrase")
		wantOK(t, err)
		_, err = h.refresh(tokens.GetRefreshToken())
		wantOK(t, err)
		_, err = h.reauthenticate(a.token, "a long enough passphrase")
		wantOK(t, err)
		_, err = h.client.StartRecovery(context.Background(), &casav1.StartRecoveryRequest{RequestContext: identityRC(ios), Email: "a@example.test"})
		wantOK(t, err)
		wantOK(t, h.signOut(a.token))
	})

	t.Run("reading is never held", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		h.notice(t, a)
		if got := h.memberships(t, a); len(got) != 0 {
			t.Fatalf("memberships = %v, want none", got)
		}
		for _, procedure := range noticeFreeReads {
			m := methodOf(t, procedure)
			if err := h.callWithContext(m, a.token, identityRC(ios)); err != nil && slices.Contains(errorCodes(t, err), casav1.ErrorCode_ERROR_CODE_NOTICE_UNREAD) {
				t.Errorf("%s: %v, want it past the hold", procedure, err)
			}
		}
	})

	// Carried from the Task 4 review (M2): a new read RPC cannot skip the
	// hold decision; it is named here or in noticeHeldProcedures.
	t.Run("every RPC that is not a command is a named read or held", func(t *testing.T) {
		held := api.NoticeHeldProcedures()
		for _, file := range casaFiles(t) {
			for i := range file.Services().Len() {
				service := file.Services().Get(i)
				if service.FullName() == "casa.v1.IdentityService" { // never held (09.3 §1)
					continue
				}
				for j := range service.Methods().Len() {
					m := service.Methods().Get(j)
					procedure := procedureOf(m)
					if proto.HasExtension(m.Options(), casav1.E_Command) {
						continue
					}
					if !slices.Contains(noticeFreeReads, procedure) && !held[procedure] {
						t.Errorf("%s is neither a command, a named read, nor in noticeHeldProcedures", procedure)
					}
					if slices.Contains(noticeFreeReads, procedure) && held[procedure] {
						t.Errorf("%s is both a named read and held", procedure)
					}
				}
			}
		}
		for _, procedure := range noticeFreeReads {
			methodOf(t, procedure)
		}
	})

	t.Run("another account is not affected", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		b := h.signUpMember(t, "b@example.test")
		if !h.held(t, a) || h.held(t, b) {
			t.Fatal("want a held and b not")
		}
	})

	// 10.1 §11 note of 14 November 2026 (d): CreateHousehold is not
	// notice_exempt (SEC-32).
	t.Run("creating a household before acknowledging the notice is held", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		_, err := h.create(a, newHousehold())
		wantCode(t, err, connect.CodeFailedPrecondition)
		if got := errorCodes(t, err); !slices.Equal(got, noticeUnread) {
			t.Fatalf("error codes = %v, want %v", got, noticeUnread)
		}
		h.wantNothingWritten(t)
	})

	// 09.3 §1 "Penolakan ini menunda, tidak menolak": the key is not stored,
	// so the same command is applied once after the acknowledgement.
	t.Run("a held command keeps its idempotency key free", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpUnread(t, "a@example.test")
		req := newHousehold()
		if _, err := h.create(a, req); !slices.Equal(errorCodes(t, err), noticeUnread) {
			t.Fatalf("first send: %v, want NOTICE_UNREAD", err)
		}
		h.ack(t, a)
		res, err := h.create(a, req)
		if err != nil {
			t.Fatal(err)
		}
		again, err := h.create(a, req)
		if err != nil || again.GetHouseholdId() != res.GetHouseholdId() {
			t.Fatalf("resend = %v (%v), want the stored %v", again, err, res)
		}
		if n := h.count(t, "SELECT count(*) FROM household"); n != 1 {
			t.Fatalf("%d households, want 1", n)
		}
	})
}
