package api_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/abiyanzimran/casa/server/internal/api"
	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

const testOrigin = "https://casa.example"

// cheapPassword keeps the Argon2 cost of the API tests low; the production
// parameters are proven in package identity (SEC-11).
var cheapPassword = identity.PasswordParams{Time: 1, MemoryKiB: 64, Parallelism: 1, SaltBytes: 16, KeyBytes: 32}

// testClock is a clock the test moves; handlers read it concurrently.
type testClock struct {
	mu  sync.Mutex
	now time.Time
}

func (c *testClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *testClock) advance(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = c.now.Add(d)
}

// identityHarness is the identity service over real HTTP and a real
// PostgreSQL 18 database (09.4 §4 row "AC Tier 1"), with a test clock and a
// mail sender that keeps the messages.
type identityHarness struct {
	db     databasetest.Database
	pool   *pgxpool.Pool
	admin  *pgxpool.Pool
	clock  *testClock
	client casav1connect.IdentityServiceClient
	// probe calls an authenticated procedure, ProbeService.Read.
	probe     *connect.Client
	probeRead protoreflect.MethodDescriptor
	mu        sync.Mutex
	mail      []identity.Message
}

func startIdentity(t *testing.T) *identityHarness {
	t.Helper()
	d := databasetest.Migrated(t)
	h := &identityHarness{
		db: d, pool: d.Pool(t, "app_rw"), admin: d.Pool(t, d.AdminRole),
		clock: &testClock{now: time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)},
	}
	service := identity.New(h.pool, identity.Config{
		Origin:   testOrigin,
		Now:      h.clock.Now,
		Password: cheapPassword,
		SendMail: func(_ context.Context, m identity.Message) error {
			h.mu.Lock()
			defer h.mu.Unlock()
			h.mail = append(h.mail, m)
			return nil
		},
	})
	server := api.NewServer(api.Version{}, service.Authenticate)
	api.RegisterIdentityService(server, service)
	h.probeRead = probeFile(t, probeMethod{name: "Read"}).Services().Get(0).Methods().ByName("Read")
	server.Register(connect.Method{
		Spec: connect.Spec{Procedure: "/casa.apitest.v1.ProbeService/Read", Schema: h.probeRead},
		Handler: func(_ context.Context, _ connect.Spec, stream connect.ServerStream) error {
			if err := stream.Receive(dynamicpb.NewMessage(h.probeRead.Input())); err != nil {
				return err
			}
			return stream.Send(&emptypb.Empty{})
		},
	})
	srv := httptest.NewServer(api.NewMux(server))
	t.Cleanup(srv.Close)
	client := connect.NewClient(connecthttp.NewTransport(srv.Client(), srv.URL))
	h.client = casav1connect.NewIdentityServiceClient(client)
	h.probe = client
	return h
}

func identityRC(kind casav1.ClientKind) *casav1.RequestContext {
	return &casav1.RequestContext{ClientKind: kind, ClientVersion: "1.0.0"}
}

var ios = casav1.ClientKind_CLIENT_KIND_IOS

// issue issues a setup token as casa setup-token does.
func (h *identityHarness) issue(t *testing.T, email string) string {
	t.Helper()
	token, _, err := identity.IssueSetupToken(t.Context(), h.pool, email, h.clock.Now())
	if err != nil {
		t.Fatal(err)
	}
	return token
}

func (h *identityHarness) startSignUp(setupToken, email string) error {
	_, err := h.client.StartSignUp(context.Background(), &casav1.StartSignUpRequest{
		RequestContext: identityRC(ios), SetupToken: setupToken, Email: email,
	})
	return err
}

func (h *identityHarness) completeSignUp(kind casav1.ClientKind, link, password string) (*casav1.SessionTokens, error) {
	res, err := h.client.CompleteSignUp(context.Background(), &casav1.CompleteSignUpRequest{
		RequestContext: identityRC(kind), VerificationToken: link, Password: password,
	})
	return res.GetTokens(), err
}

// authenticated calls an authenticated procedure with accessToken.
func (h *identityHarness) authenticated(accessToken string) error {
	req := dynamicpb.NewMessage(h.probeRead.Input())
	req.Set(h.probeRead.Input().Fields().ByName("request_context"), protoreflect.ValueOfMessage(identityRC(ios).ProtoReflect()))
	ctx, info := connect.NewClientContext(context.Background())
	info.RequestHeader().Set("Authorization", "Bearer "+accessToken)
	return h.probe.CallUnary(ctx, connect.Spec{Procedure: "/casa.apitest.v1.ProbeService/Read"}, req, &emptypb.Empty{})
}

func (h *identityHarness) signIn(email, password string) (*casav1.SessionTokens, error) {
	res, err := h.client.SignIn(context.Background(), &casav1.SignInRequest{
		RequestContext: identityRC(ios), Email: email, Password: password,
	})
	return res.GetTokens(), err
}

func (h *identityHarness) refresh(refreshToken string) (*casav1.SessionTokens, error) {
	res, err := h.client.RefreshSession(context.Background(), &casav1.RefreshSessionRequest{
		RequestContext: identityRC(ios), RefreshToken: refreshToken,
	})
	return res.GetTokens(), err
}

func (h *identityHarness) mails() []identity.Message {
	h.mu.Lock()
	defer h.mu.Unlock()
	return slices.Clone(h.mail)
}

// lastLink returns the token in the fragment of the last mail's link.
func (h *identityHarness) lastLink(t *testing.T) string {
	t.Helper()
	mails := h.mails()
	if len(mails) == 0 {
		t.Fatal("no mail sent")
	}
	_, token, ok := strings.Cut(mails[len(mails)-1].Link, "#")
	if !ok {
		t.Fatalf("link %q has no fragment", mails[len(mails)-1].Link)
	}
	return token
}

// signUp runs the whole closed sign-up for email and returns the session.
func (h *identityHarness) signUp(t *testing.T, email, password string) *casav1.SessionTokens {
	t.Helper()
	if err := h.startSignUp(h.issue(t, email), email); err != nil {
		t.Fatal(err)
	}
	tokens, err := h.completeSignUp(ios, h.lastLink(t), password)
	if err != nil {
		t.Fatal(err)
	}
	return tokens
}

func (h *identityHarness) count(t *testing.T, query string, args ...any) int {
	t.Helper()
	var n int
	if err := h.admin.QueryRow(t.Context(), query, args...).Scan(&n); err != nil {
		t.Fatalf("%s: %v", query, err)
	}
	return n
}

// rejection is everything a client can see of an error.
type rejection struct {
	code    connect.Code
	message string
	details []connect.ErrorDetail
}

func rejectionOf(t *testing.T, err error) rejection {
	t.Helper()
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		t.Fatalf("error %v is not a *connect.Error", err)
	}
	r := rejection{code: connectErr.Code(), message: connectErr.Message()}
	for _, d := range connectErr.Details() {
		r.details = append(r.details, connect.ErrorDetail{Type: d.Type, Value: d.Value})
	}
	return r
}

func (r rejection) equal(o rejection) bool {
	return r.code == o.code && r.message == o.message &&
		slices.EqualFunc(r.details, o.details, func(a, b connect.ErrorDetail) bool {
			return a.Type == b.Type && bytes.Equal(a.Value, b.Value)
		})
}

func wantErrorCode(t *testing.T, err error, code connect.Code, detail casav1.ErrorCode) {
	t.Helper()
	wantCode(t, err, code)
	if got := errorCodes(t, err); !slices.Equal(got, []casav1.ErrorCode{detail}) {
		t.Fatalf("error codes = %v, want [%v]", got, detail)
	}
}

// violations returns the field path and rule id of each violation in err.
func violations(t *testing.T, err error) []string {
	t.Helper()
	var got []string
	for _, detail := range details(t, err) {
		if v, ok := detail.(*validate.Violations); ok {
			for _, violation := range v.GetViolations() {
				got = append(got, fmt.Sprintf("%s %s", violation.GetField().GetElements()[0].GetFieldName(), violation.GetRuleId()))
			}
		}
	}
	return got
}

func TestSignUp(t *testing.T) {
	const password = "a long enough passphrase"
	rejected := casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED

	t.Run("FR-FAM-001/AC-9 a sign-up without a valid setup token gets one identical rejection", func(t *testing.T) {
		h := startIdentity(t)
		used := h.issue(t, "used@example.test")
		if err := h.startSignUp(used, "used@example.test"); err != nil {
			t.Fatal(err)
		}
		if _, err := h.completeSignUp(ios, h.lastLink(t), password); err != nil {
			t.Fatal(err)
		}
		mailsBefore := len(h.mails())
		expired := h.issue(t, "late@example.test")
		bound := h.issue(t, "a@example.test")
		h.clock.advance(7*24*time.Hour - time.Second)
		stillValid := h.issue(t, "fresh@example.test")
		h.clock.advance(time.Second) // expired is now exactly 7 days old
		accounts := h.count(t, "SELECT count(*) FROM account")
		neverIssued := strings.Repeat("A", 43)
		var first rejection
		for i, c := range []struct{ name, token, email string }{
			{"no token", "", "x@example.test"},
			{"malformed", "abc", "x@example.test"},
			{"never issued", neverIssued, "x@example.test"},
			{"expired", expired, "late@example.test"},
			{"already used", used, "used@example.test"},
			{"bound to another email", bound, "b@example.test"},
		} {
			err := h.startSignUp(c.token, c.email)
			wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
			if got := violations(t, err); got != nil {
				t.Fatalf("%s: violations %v, want none", c.name, got)
			}
			r := rejectionOf(t, err)
			if i == 0 {
				first = r
			} else if !r.equal(first) {
				t.Fatalf("%s: rejection %+v differs from %+v", c.name, r, first)
			}
		}
		if got := len(h.mails()); got != mailsBefore {
			t.Fatalf("%d mails sent for rejected sign-ups", got-mailsBefore)
		}
		if got := h.count(t, "SELECT count(*) FROM account"); got != accounts {
			t.Fatalf("%d accounts made by rejected sign-ups", got-accounts)
		}
		if err := h.startSignUp(stillValid, "fresh@example.test"); err != nil {
			t.Fatalf("a token just under 7 days old: %v", err)
		}
		// CompleteSignUp answers a bad link with the same rejection.
		expiredLink := h.lastLink(t)
		h.clock.advance(30 * time.Minute)
		for _, c := range []struct{ name, link string }{
			{"no link", ""},
			{"malformed link", "abc"},
			{"unknown link", neverIssued},
			{"expired link", expiredLink},
		} {
			_, err := h.completeSignUp(ios, c.link, password)
			wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
			if r := rejectionOf(t, err); !r.equal(first) {
				t.Fatalf("CompleteSignUp, %s: rejection %+v differs from %+v", c.name, r, first)
			}
		}
		if got := h.count(t, "SELECT count(*) FROM account"); got != accounts {
			t.Fatalf("%d accounts made by rejected links", got-accounts)
		}
	})

	t.Run("FR-FAM-001/AC-9 repeated failed sign-ups hit the rate limit", func(t *testing.T) {
		h := startIdentity(t)
		valid := h.issue(t, "owner@example.test")
		for range 20 {
			wantErrorCode(t, h.startSignUp("", "owner@example.test"), connect.CodePermissionDenied, rejected)
		}
		wantErrorCode(t, h.startSignUp(valid, "owner@example.test"), connect.CodeResourceExhausted, casav1.ErrorCode_ERROR_CODE_RATE_LIMITED)
		if len(h.mails()) != 0 {
			t.Fatal("a rate-limited request sent mail")
		}
		h.clock.advance(time.Second)
		wantOK(t, h.startSignUp(valid, "owner@example.test"))
	})

	t.Run("FR-FAM-001/AC-10 a setup token opens one account after the emailed link is opened", func(t *testing.T) {
		h := startIdentity(t)
		setup := h.issue(t, "Owner@Example.Test")
		wantOK(t, h.startSignUp(setup, "owner@example.test"))
		mails := h.mails()
		if len(mails) != 1 || mails[0].To != "Owner@Example.Test" || mails[0].Kind != identity.MessageVerification ||
			!strings.HasPrefix(mails[0].Link, testOrigin+"/signup/verify#") || len(h.lastLink(t)) != 43 {
			t.Fatalf("mail %+v, want one verification link to Owner@Example.Test", mails)
		}
		if got := h.count(t, "SELECT count(*) FROM account"); got != 0 {
			t.Fatalf("%d accounts before the link is opened, want 0", got)
		}
		tokens, err := h.completeSignUp(ios, h.lastLink(t), password)
		wantOK(t, err)
		if len(tokens.GetAccessToken()) != 43 || len(tokens.GetRefreshToken()) != 43 ||
			!tokens.GetAccessTokenExpiresAt().AsTime().Equal(h.clock.Now().Add(15*time.Minute)) {
			t.Fatalf("tokens %v, want 43-character tokens and an access token for 15 minutes", tokens)
		}
		var email, via string
		var setupUsed bool
		err = h.admin.QueryRow(t.Context(), `
			SELECT a.email::text, a.signup_via, s.used_at IS NOT NULL AND s.account_id = a.id
			  FROM account a, setup_token s`).Scan(&email, &via, &setupUsed)
		if err != nil || email != "Owner@Example.Test" || via != "setup_token" || !setupUsed {
			t.Fatalf("account %q via %q, setup token used for it %t, err %v", email, via, setupUsed, err)
		}
		if got := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'sign_in' AND session_id IS NOT NULL"); got != 1 {
			t.Fatalf("%d sign_in events, want 1", got)
		}
	})

	t.Run("FR-FAM-001/AC-10 a used setup token or link is rejected like any invalid token", func(t *testing.T) {
		h := startIdentity(t)
		setup := h.issue(t, "owner@example.test")
		wantOK(t, h.startSignUp(setup, "owner@example.test"))
		link := h.lastLink(t)
		_, err := h.completeSignUp(ios, link, password)
		wantOK(t, err)
		none := rejectionOf(t, h.startSignUp("", "owner@example.test"))
		_, err = h.completeSignUp(ios, link, password)
		if r := rejectionOf(t, err); !r.equal(none) {
			t.Fatalf("the used link: %+v, want %+v", r, none)
		}
		if r := rejectionOf(t, h.startSignUp(setup, "owner@example.test")); !r.equal(none) {
			t.Fatalf("the used setup token: %+v, want %+v", r, none)
		}
	})

	t.Run("SEC-158 a verification link expires after 30 minutes and a new one replaces it", func(t *testing.T) {
		h := startIdentity(t)
		setup := h.issue(t, "owner@example.test")
		wantOK(t, h.startSignUp(setup, "owner@example.test"))
		old := h.lastLink(t)
		h.clock.advance(30 * time.Minute)
		_, err := h.completeSignUp(ios, old, password)
		wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
		wantOK(t, h.startSignUp(setup, "owner@example.test"))
		fresh := h.lastLink(t)
		_, err = h.completeSignUp(ios, old, password)
		wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
		_, err = h.completeSignUp(ios, fresh, password)
		wantOK(t, err)
	})

	t.Run("SEC-11 a password under 12 characters or on the common list is refused and the link stays usable", func(t *testing.T) {
		h := startIdentity(t)
		wantOK(t, h.startSignUp(h.issue(t, "owner@example.test"), "owner@example.test"))
		link := h.lastLink(t)
		_, err := h.completeSignUp(ios, link, "elevenchars")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violations(t, err), []string{"password string.min_len"}; !slices.Equal(got, want) {
			t.Fatalf("violations %v, want %v", got, want)
		}
		_, err = h.completeSignUp(ios, link, "password1234")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violations(t, err), []string{"password password.common"}; !slices.Equal(got, want) {
			t.Fatalf("violations %v, want %v", got, want)
		}
		_, err = h.completeSignUp(ios, link, password)
		wantOK(t, err)
	})

	t.Run("SEC-11 the database holds no password and no unsalted password hash", func(t *testing.T) {
		h := startIdentity(t)
		setup := h.issue(t, "owner@example.test")
		wantOK(t, h.startSignUp(setup, "owner@example.test"))
		link := h.lastLink(t)
		tokens, err := h.completeSignUp(ios, link, password)
		wantOK(t, err)
		sum := sha256.Sum256([]byte(password))
		secrets := [][]byte{[]byte(password), sum[:], []byte(setup), []byte(link),
			[]byte(tokens.GetAccessToken()), []byte(tokens.GetRefreshToken())}
		rows, err := h.admin.Query(t.Context(), `
			SELECT c.table_name, c.column_name FROM information_schema.columns c
			  JOIN information_schema.tables t USING (table_schema, table_name)
			 WHERE c.table_schema = 'public' AND t.table_type = 'BASE TABLE'
			   AND (c.data_type IN ('text', 'bytea', 'character') OR c.udt_name = 'citext')`)
		if err != nil {
			t.Fatal(err)
		}
		columns, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) ([2]string, error) {
			var c [2]string
			return c, row.Scan(&c[0], &c[1])
		})
		if err != nil || len(columns) < 10 {
			t.Fatalf("%d columns, err %v", len(columns), err)
		}
		for _, c := range columns {
			query := fmt.Sprintf("SELECT count(*) FROM %s WHERE position($1::bytea IN convert_to(%s::text, 'UTF8')) > 0 OR position($1::bytea IN %s::text::bytea) > 0",
				pgx.Identifier{c[0]}.Sanitize(), pgx.Identifier{c[1]}.Sanitize(), pgx.Identifier{c[1]}.Sanitize())
			if c[0] == "goose_db_version" {
				continue
			}
			for i, secret := range secrets {
				if n := h.count(t, query, secret); n != 0 {
					t.Errorf("%s.%s holds secret %d", c[0], c[1], i)
				}
			}
		}
	})

	t.Run("an email that already has an account cannot sign up again", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		err := h.startSignUp(h.issue(t, "OWNER@example.test"), "OWNER@example.test")
		wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
	})

	t.Run("a web client cannot open a session before the cookie transport exists (1.5)", func(t *testing.T) {
		// SEC-161 (b): a web session token travels only in a cookie.
		h := startIdentity(t)
		wantOK(t, h.startSignUp(h.issue(t, "owner@example.test"), "owner@example.test"))
		_, err := h.completeSignUp(casav1.ClientKind_CLIENT_KIND_WEB, h.lastLink(t), password)
		wantCode(t, err, connect.CodeUnimplemented)
		_, err = h.completeSignUp(ios, h.lastLink(t), password)
		wantOK(t, err)
	})

	t.Run("SEC-12 failed token uses are limited per address", func(t *testing.T) {
		// 10.1 §8.1: CompleteSignUp shares the token-use limit.
		h := startIdentity(t)
		wantOK(t, h.startSignUp(h.issue(t, "owner@example.test"), "owner@example.test"))
		link := h.lastLink(t)
		for range 20 {
			_, err := h.completeSignUp(ios, strings.Repeat("B", 43), password)
			wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
		}
		_, err := h.completeSignUp(ios, link, password)
		wantErrorCode(t, err, connect.CodeResourceExhausted, casav1.ErrorCode_ERROR_CODE_RATE_LIMITED)
	})
}

func TestSessions(t *testing.T) {
	const password = "a long enough passphrase"
	credentialsRejected := casav1.ErrorCode_ERROR_CODE_CREDENTIALS_REJECTED
	rateLimited := casav1.ErrorCode_ERROR_CODE_RATE_LIMITED
	revoked := casav1.ErrorCode_ERROR_CODE_SESSION_REVOKED
	expired := casav1.ErrorCode_ERROR_CODE_SESSION_EXPIRED

	// wantUnauthenticated: no detail means "refresh, then sign in".
	wantUnauthenticated := func(t *testing.T, err error) {
		t.Helper()
		wantCode(t, err, connect.CodeUnauthenticated)
		if got := details(t, err); got != nil {
			t.Fatalf("details %v, want none", got)
		}
	}

	t.Run("SEC-12 a failed sign-in does not tell an unknown email from a wrong password", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		_, unknown := h.signIn("nobody@example.test", password)
		_, wrong := h.signIn("owner@example.test", "a wrong passphrase")
		wantErrorCode(t, unknown, connect.CodeUnauthenticated, credentialsRejected)
		if a, b := rejectionOf(t, unknown), rejectionOf(t, wrong); !a.equal(b) {
			t.Fatalf("unknown email %+v, wrong password %+v; want identical", a, b)
		}
	})

	t.Run("SEC-12 one address trying many accounts is limited", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		for i := range 20 {
			_, err := h.signIn(fmt.Sprintf("guess%d@example.test", i), password)
			wantErrorCode(t, err, connect.CodeUnauthenticated, credentialsRejected)
		}
		_, err := h.signIn("owner@example.test", password)
		wantErrorCode(t, err, connect.CodeResourceExhausted, rateLimited)
	})

	t.Run("SEC-12 a successful sign-in does not reset the address key", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		for range 19 {
			_, err := h.signIn("partner@example.test", password)
			wantErrorCode(t, err, connect.CodeUnauthenticated, credentialsRejected)
		}
		_, err := h.signIn("owner@example.test", password)
		wantOK(t, err)
		_, err = h.signIn("partner@example.test", password)
		wantErrorCode(t, err, connect.CodeUnauthenticated, credentialsRejected)
		_, err = h.signIn("owner@example.test", password)
		wantErrorCode(t, err, connect.CodeResourceExhausted, rateLimited)
	})

	t.Run("SEC-15 a refresh token is rotated on every use", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		first, err := h.signIn("owner@example.test", password)
		wantOK(t, err)
		wantOK(t, h.authenticated(first.GetAccessToken()))
		second, err := h.refresh(first.GetRefreshToken())
		wantOK(t, err)
		if second.GetRefreshToken() == first.GetRefreshToken() || second.GetAccessToken() == first.GetAccessToken() {
			t.Fatal("refresh returned the same tokens")
		}
		wantUnauthenticated(t, h.authenticated(first.GetAccessToken()))
		wantOK(t, h.authenticated(second.GetAccessToken()))
	})

	t.Run("SEC-15 reusing a refresh token revokes the whole session", func(t *testing.T) {
		h := startIdentity(t)
		first := h.signUp(t, "owner@example.test", password)
		second, err := h.refresh(first.GetRefreshToken())
		wantOK(t, err)
		_, err = h.refresh(first.GetRefreshToken())
		wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
		_, err = h.refresh(second.GetRefreshToken())
		wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
		wantErrorCode(t, h.authenticated(second.GetAccessToken()), connect.CodeUnauthenticated, revoked)
		if n := h.count(t, "SELECT count(*) FROM session WHERE revoke_reason = 'reuse_detected'"); n != 1 {
			t.Fatalf("%d sessions revoked for reuse, want 1", n)
		}
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'session_revoked' AND detail = 'reuse_detected'"); n != 1 {
			t.Fatalf("%d session_revoked events, want 1", n)
		}
	})

	t.Run("two refreshes with one token at once: one succeeds, the session is revoked", func(t *testing.T) {
		// SEC-15 under concurrency; the serializable retry decides the order.
		h := startIdentity(t)
		first := h.signUp(t, "owner@example.test", password)
		var wg sync.WaitGroup
		results := make([]error, 2)
		tokens := make([]*casav1.SessionTokens, 2)
		for i := range 2 {
			wg.Go(func() { tokens[i], results[i] = h.refresh(first.GetRefreshToken()) })
		}
		wg.Wait()
		ok := slices.IndexFunc(results, func(err error) bool { return err == nil })
		if ok < 0 || results[1-ok] == nil {
			t.Fatalf("results %v, want exactly one success", results)
		}
		wantErrorCode(t, results[1-ok], connect.CodeUnauthenticated, revoked)
		wantErrorCode(t, h.authenticated(tokens[ok].GetAccessToken()), connect.CodeUnauthenticated, revoked)
	})

	t.Run("an access token expires after 15 minutes and the refresh token renews it", func(t *testing.T) {
		h := startIdentity(t)
		first := h.signUp(t, "owner@example.test", password)
		h.clock.advance(15*time.Minute - time.Second)
		wantOK(t, h.authenticated(first.GetAccessToken()))
		h.clock.advance(time.Second)
		wantUnauthenticated(t, h.authenticated(first.GetAccessToken()))
		second, err := h.refresh(first.GetRefreshToken())
		wantOK(t, err)
		wantOK(t, h.authenticated(second.GetAccessToken()))
	})

	t.Run("a native session ends after 30 idle days or 180 days", func(t *testing.T) {
		// Lead ruling Q5 on the 1.4 plan [Usulan].
		h := startIdentity(t)
		idle := h.signUp(t, "owner@example.test", password)
		h.clock.advance(30 * 24 * time.Hour)
		wantErrorCode(t, h.authenticated(idle.GetAccessToken()), connect.CodeUnauthenticated, expired)
		_, err := h.refresh(idle.GetRefreshToken())
		wantErrorCode(t, err, connect.CodeUnauthenticated, expired)

		active, err := h.signIn("owner@example.test", password)
		wantOK(t, err)
		for day := 29; day < 180; day += 29 {
			h.clock.advance(29 * 24 * time.Hour)
			active, err = h.refresh(active.GetRefreshToken())
			wantOK(t, err)
		}
		h.clock.advance(6 * 24 * time.Hour) // day 180
		_, err = h.refresh(active.GetRefreshToken())
		wantErrorCode(t, err, connect.CodeUnauthenticated, expired)
	})

	t.Run("a web session is never accepted from the Authorization header", func(t *testing.T) {
		// SEC-161 (b): a web session token travels only in a cookie (1.5).
		h := startIdentity(t)
		tokens := h.signUp(t, "owner@example.test", password)
		if _, err := h.admin.Exec(t.Context(), "UPDATE session SET client_kind = 'web'"); err != nil {
			t.Fatal(err)
		}
		wantUnauthenticated(t, h.authenticated(tokens.GetAccessToken()))
	})

	t.Run("a sign-in writes the account's security trail", func(t *testing.T) {
		// SEC-16, SEC-39 (b).
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		_, err := h.signIn("owner@example.test", password)
		wantOK(t, err)
		if n := h.count(t, "SELECT count(DISTINCT session_id) FROM security_event WHERE kind = 'sign_in'"); n != 2 {
			t.Fatalf("%d sign_in events with distinct sessions, want 2 (sign-up and sign-in)", n)
		}
	})

	t.Run("a web client cannot sign in before the cookie transport exists (1.5)", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		_, err := h.client.SignIn(context.Background(), &casav1.SignInRequest{
			RequestContext: identityRC(casav1.ClientKind_CLIENT_KIND_WEB), Email: "owner@example.test", Password: password,
		})
		wantCode(t, err, connect.CodeUnimplemented)
	})
}
