package api_test

import (
	"context"
	"net/http"
	"slices"
	"testing"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
)

// browser is one person's web client: its own HTTP client and its own
// session cookie, which it sends with every call and replaces when a
// response sets a new one, as a browser does (SEC-161).
type browser struct {
	t          *testing.T
	cookie     string
	identity   casav1connect.IdentityServiceClient
	households casav1connect.HouseholdServiceClient
}

func (h *householdHarness) newBrowser(t *testing.T) *browser {
	t.Helper()
	transport := &http.Transport{}
	t.Cleanup(transport.CloseIdleConnections)
	client := connect.NewClient(connecthttp.NewTransport(&http.Client{Transport: transport}, h.url))
	return &browser{
		t: t, identity: casav1connect.NewIdentityServiceClient(client),
		households: casav1connect.NewHouseholdServiceClient(client),
	}
}

// call runs fn as the app's fetch sends it: the cookie, the app's Origin,
// Sec-Fetch-Site and Safari's User-Agent; it keeps a cookie the response
// sets.
func (b *browser) call(fn func(ctx context.Context) error) error {
	b.t.Helper()
	ctx := sameOrigin(b.cookie)
	info, _ := connect.CallInfoForClientContext(ctx)
	info.RequestHeader().Set("User-Agent", safariIPhone)
	err := fn(ctx)
	if cookie := newCookie(b.t, info); cookie != nil {
		b.cookie = cookie.Value
	}
	return err
}

// do is call for a step that must succeed.
func (b *browser) do(what string, fn func(ctx context.Context) error) {
	b.t.Helper()
	if err := b.call(fn); err != nil {
		b.t.Fatalf("%s: %v", what, err)
	}
}

// webRC is a web client's request context; a command carries a fresh
// idempotency key.
func webRC(command bool) *casav1.RequestContext {
	rc := identityRC(web)
	if command {
		rc.IdempotencyKey = uuid4()
	}
	return rc
}

// signUp opens the emailed link in personal mode, then reads and
// acknowledges the operator notice (FR-FAM-019/AC-1; SEC-34).
func (b *browser) signUp(h *householdHarness, password string) {
	b.t.Helper()
	b.do("CompleteSignUp", func(ctx context.Context) error {
		_, err := b.identity.CompleteSignUp(ctx, &casav1.CompleteSignUpRequest{
			RequestContext: webRC(false), VerificationToken: h.lastLink(b.t), Password: password, DeviceMode: personal,
		})
		return err
	})
	if b.cookie == "" {
		b.t.Fatal("CompleteSignUp set no session cookie")
	}
	var n *casav1.GetOperatorNoticeResponse
	b.do("GetOperatorNotice", func(ctx context.Context) (err error) {
		n, err = b.households.GetOperatorNotice(ctx, &casav1.GetOperatorNoticeRequest{RequestContext: webRC(false)})
		return err
	})
	b.do("AcknowledgeOperatorNotice", func(ctx context.Context) error {
		_, err := b.households.AcknowledgeOperatorNotice(ctx, &casav1.AcknowledgeOperatorNoticeRequest{
			RequestContext: webRC(true), Variant: n.GetVariant(), Version: n.GetVersion(), TextSha256: n.GetTextSha256(),
		})
		return err
	})
}

// TestTwoAccountsJoinOneHousehold is 14 §3.3 step 1 "Selesai bila" over the
// API, as the web client calls it: two people, each with their own browser
// and session cookie, end as the two members of one household.
func TestTwoAccountsJoinOneHousehold(t *testing.T) {
	h := startHousehold(t)
	a, b := h.newBrowser(t), h.newBrowser(t)

	// A: the setup token's account creates the household and invites.
	a.do("A StartSignUp(setup_token)", func(ctx context.Context) error {
		_, err := a.identity.StartSignUp(ctx, &casav1.StartSignUpRequest{
			RequestContext: webRC(false), SetupToken: h.issue(t, "a@example.test"), Email: "a@example.test",
		})
		return err
	})
	a.signUp(h, "a long enough passphrase")
	var created *casav1.CreateHouseholdResponse
	a.do("A CreateHousehold", func(ctx context.Context) (err error) {
		created, err = a.households.CreateHousehold(ctx, &casav1.CreateHouseholdRequest{
			RequestContext: webRC(true), Name: "Rumah Kita", TimeZone: "Asia/Jakarta", DisplayName: "Ayu", AgeDeclarationVersion: 1,
		})
		return err
	})
	household := created.GetHouseholdId()
	var invitation *casav1.CreateInvitationResponse
	a.do("A CreateInvitation", func(ctx context.Context) (err error) {
		invitation, err = a.households.CreateInvitation(ctx, &casav1.CreateInvitationRequest{RequestContext: webRC(true), HouseholdId: household})
		return err
	})
	var link *casav1.IssueInvitationLinkResponse
	a.do("A IssueInvitationLink", func(ctx context.Context) (err error) {
		link, err = a.households.IssueInvitationLink(ctx, &casav1.IssueInvitationLinkRequest{
			RequestContext: webRC(false), HouseholdId: household, InvitationId: invitation.GetInvitationId(),
		})
		return err
	})

	// B: signs up through the link and accepts without pasting it again.
	b.do("B StartSignUp(invitation_token)", func(ctx context.Context) error {
		_, err := b.identity.StartSignUp(ctx, &casav1.StartSignUpRequest{
			RequestContext: webRC(false), InvitationToken: link.GetInvitationToken(), Email: "b@example.test",
		})
		return err
	})
	b.signUp(h, "a long enough passphrase")
	var preview *casav1.PreviewInvitationResponse
	b.do("B PreviewInvitation{from_signup}", func(ctx context.Context) (err error) {
		preview, err = b.households.PreviewInvitation(ctx, &casav1.PreviewInvitationRequest{RequestContext: webRC(false), Source: fromSignup()})
		return err
	})
	if preview.GetHouseholdName() != "Rumah Kita" || preview.GetInviterDisplayName() != "Ayu" {
		t.Fatalf("B's preview %v, want Rumah Kita by Ayu", preview)
	}
	b.do("B AcceptInvitation{from_signup}", func(ctx context.Context) error {
		_, err := b.households.AcceptInvitation(ctx, &casav1.AcceptInvitationRequest{
			RequestContext: webRC(true), Source: fromSignup(), DisplayName: "Budi", AgeDeclarationVersion: 1,
		})
		return err
	})

	// A: sees B's verified email and confirms (F04; SEC-31).
	var open *casav1.GetOpenInvitationResponse
	a.do("A GetOpenInvitation", func(ctx context.Context) (err error) {
		open, err = a.households.GetOpenInvitation(ctx, &casav1.GetOpenInvitationRequest{RequestContext: webRC(false), HouseholdId: household})
		return err
	})
	if open.GetState() != casav1.InvitationState_INVITATION_STATE_ACCEPTED || open.GetInviteeEmail() != "b@example.test" ||
		open.GetInviteeDisplayName() != "Budi" {
		t.Fatalf("A's open invitation %v, want accepted by Budi, b@example.test", open)
	}
	var confirmed *casav1.ConfirmInviteeResponse
	a.do("A ConfirmInvitee", func(ctx context.Context) (err error) {
		confirmed, err = a.households.ConfirmInvitee(ctx, &casav1.ConfirmInviteeRequest{
			RequestContext: webRC(true), HouseholdId: household, InvitationId: invitation.GetInvitationId(),
		})
		return err
	})

	// Both: one household with members A and B, and it is full.
	want := []string{created.GetMemberId() + " Ayu", confirmed.GetMemberId() + " Budi"}
	slices.Sort(want)
	for name, p := range map[string]*browser{"A": a, "B": b} {
		var res *casav1.ListMembershipsResponse
		p.do(name+" ListMemberships", func(ctx context.Context) (err error) {
			res, err = p.households.ListMemberships(ctx, &casav1.ListMembershipsRequest{RequestContext: webRC(false)})
			return err
		})
		ms := res.GetMemberships()
		if len(ms) != 1 || ms[0].GetHouseholdId() != household {
			t.Fatalf("%s's memberships %v, want only %s", name, ms, household)
		}
		var got []string
		for _, m := range ms[0].GetMembers() {
			got = append(got, m.GetMemberId()+" "+m.GetDisplayName())
		}
		slices.Sort(got)
		if !slices.Equal(got, want) {
			t.Fatalf("%s sees members %v, want %v", name, got, want)
		}
		err := p.call(func(ctx context.Context) error {
			_, err := p.households.CreateInvitation(ctx, &casav1.CreateInvitationRequest{RequestContext: webRC(true), HouseholdId: household})
			return err
		})
		wantCodes(t, err, connect.CodeFailedPrecondition, householdFull)
	}
}
