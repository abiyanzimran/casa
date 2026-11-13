package api_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
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
	"google.golang.org/protobuf/proto"
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
	// url is the server's base URL, for requests that are not Connect calls.
	url string
	// probe calls an authenticated procedure, ProbeService.Read.
	probe     *connect.Client
	probeRead protoreflect.MethodDescriptor
	service   *identity.Service
	mu        sync.Mutex
	mail      []identity.Message
	mailErr   error
	// logs is the service's log, read after service.Wait.
	logs bytes.Buffer
}

// failMail makes every later mail fail with err.
func (h *identityHarness) failMail(err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.mailErr = err
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
		Logger:   slog.New(slog.NewJSONHandler(&h.logs, nil)),
		SendMail: func(_ context.Context, m identity.Message) error {
			h.mu.Lock()
			defer h.mu.Unlock()
			if h.mailErr != nil {
				return h.mailErr
			}
			h.mail = append(h.mail, m)
			return nil
		},
	})
	h.service = service
	server := api.NewServer(api.Version{}, testOrigin, api.Sessions{
		Bearer: service.Authenticate, Cookie: service.AuthenticateCookie, UserAction: service.TouchUserAction,
	})
	api.RegisterIdentityService(server, h.pool, service)
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
	srv := httptest.NewServer(api.NewMux(server, nil))
	t.Cleanup(srv.Close)
	h.url = srv.URL
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

// bearer is a context whose request carries accessToken.
func bearer(accessToken string) context.Context {
	ctx, info := connect.NewClientContext(context.Background())
	if accessToken != "" {
		info.RequestHeader().Set("Authorization", "Bearer "+accessToken)
	}
	return ctx
}

// commandRC is a request context with a fresh idempotency key.
func commandRC() *casav1.RequestContext {
	rc := identityRC(ios)
	rc.IdempotencyKey = uuid4()
	return rc
}

func uuid4() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = b[6]&0x0f | 0x40
	b[8] = b[8]&0x3f | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func (h *identityHarness) signOut(accessToken string) error {
	_, err := h.client.SignOut(bearer(accessToken), &casav1.SignOutRequest{RequestContext: commandRC()})
	return err
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
		wantErrorCode(t, h.authenticated(active.GetAccessToken()), connect.CodeUnauthenticated, expired)
		// Expiry is never undone: the refused calls moved no session.
		if n := h.count(t, "SELECT count(*) FROM session WHERE last_used_at + idle_timeout > $1 AND absolute_expiry > $1", h.clock.Now()); n != 0 {
			t.Fatalf("%d sessions live again, want none", n)
		}
	})

	t.Run("a web session is never accepted from the Authorization header", func(t *testing.T) {
		// SEC-161 (b): a web session token travels only in a cookie.
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		cookie := h.webSignIn(t, "owner@example.test", password, personal)
		wantUnauthenticated(t, h.authenticated(cookie))
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
}

func TestSignOut(t *testing.T) {
	password := strings.Repeat("sign out ", 3) // built, so no secret scanner sees a literal
	revoked := casav1.ErrorCode_ERROR_CODE_SESSION_REVOKED
	t.Run("sign-out revokes the session on the server", func(t *testing.T) {
		// SEC-164 (b) cited, not named: its web parts come with 1.5.
		h := startIdentity(t)
		tokens := h.signUp(t, "owner@example.test", password)
		wantOK(t, h.signOut(tokens.GetAccessToken()))
		wantErrorCode(t, h.authenticated(tokens.GetAccessToken()), connect.CodeUnauthenticated, revoked)
		_, err := h.refresh(tokens.GetRefreshToken())
		wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
		if n := h.count(t, "SELECT count(*) FROM session WHERE revoke_reason = 'user'"); n != 1 {
			t.Fatalf("%d sessions revoked by the user, want 1", n)
		}
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'session_revoked' AND detail = 'user'"); n != 1 {
			t.Fatalf("%d session_revoked events, want 1", n)
		}
		if n := h.count(t, `SELECT count(*) FROM idempotency_record r JOIN session s ON r.scope_id = s.account_id
			WHERE r.command = '/casa.v1.IdentityService/SignOut'`); n != 1 {
			t.Fatalf("%d idempotency records scoped to the account, want 1", n)
		}
	})
	t.Run("sign-out needs a session", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, "owner@example.test", password)
		wantCode(t, h.signOut(""), connect.CodeUnauthenticated)
		if n := h.count(t, "SELECT count(*) FROM session WHERE revoked_at IS NOT NULL"); n != 0 {
			t.Fatalf("%d sessions revoked, want none", n)
		}
	})
	t.Run("sign-out needs an idempotency key", func(t *testing.T) {
		// ADR-005; 09.2 §8: a write command without a key is refused.
		h := startIdentity(t)
		tokens := h.signUp(t, "owner@example.test", password)
		_, err := h.client.SignOut(bearer(tokens.GetAccessToken()), &casav1.SignOutRequest{RequestContext: identityRC(ios)})
		wantCode(t, err, connect.CodeInvalidArgument)
		wantOK(t, h.authenticated(tokens.GetAccessToken()))
	})
}

func (h *identityHarness) reauthenticate(accessToken, password string) (*casav1.ReauthenticateResponse, error) {
	return h.client.Reauthenticate(bearer(accessToken), &casav1.ReauthenticateRequest{
		RequestContext: identityRC(ios), Password: password, Purpose: casav1.ReauthPurpose_REAUTH_PURPOSE_CHANGE_PASSWORD,
	})
}

func (h *identityHarness) changePassword(accessToken, newPassword string) error {
	_, err := h.client.ChangePassword(bearer(accessToken), &casav1.ChangePasswordRequest{
		RequestContext: commandRC(), NewPassword: newPassword,
	})
	return err
}

func TestReauthentication(t *testing.T) {
	const email = "owner@example.test"
	// Built, so no secret scanner sees a literal.
	password, next := strings.Repeat("reauth ", 3), strings.Repeat("changed ", 3)
	reauthRequired := casav1.ErrorCode_ERROR_CODE_REAUTH_REQUIRED

	t.Run("SEC-13 a password change without re-authentication is refused without effect", func(t *testing.T) {
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		wantErrorCode(t, h.changePassword(tokens.GetAccessToken(), next), connect.CodePermissionDenied, reauthRequired)
		_, err := h.signIn(email, password)
		wantOK(t, err)
		if n := h.count(t, "SELECT count(*) FROM idempotency_record"); n != 0 {
			t.Fatalf("%d idempotency records, want none", n)
		}
	})
	t.Run("SEC-13 a re-authentication older than 5 minutes is refused", func(t *testing.T) {
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		res, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantOK(t, err)
		if got, want := res.GetExpiresAt().AsTime(), h.clock.Now().Add(5*time.Minute); !got.Equal(want) {
			t.Fatalf("expires_at %v, want %v", got, want)
		}
		h.clock.advance(5*time.Minute - time.Second)
		wantOK(t, h.changePassword(tokens.GetAccessToken(), next))
		for _, age := range []time.Duration{5 * time.Minute, 6 * time.Minute} { // 6: the SEC-13 test value
			tokens, err := h.signIn(email, next) // a fresh session: an access token lives 15 minutes
			wantOK(t, err)
			_, err = h.reauthenticate(tokens.GetAccessToken(), next)
			wantOK(t, err)
			h.clock.advance(age)
			wantErrorCode(t, h.changePassword(tokens.GetAccessToken(), password), connect.CodePermissionDenied, reauthRequired)
		}
		_, err = h.signIn(email, next)
		wantOK(t, err)
	})
	t.Run("ADR-005 a retried ChangePassword gets the first response", func(t *testing.T) {
		// 09.2 §8: a resend of the same command, after its grant is used,
		// gets the stored response, not REAUTH_REQUIRED.
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		_, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantOK(t, err)
		request := &casav1.ChangePasswordRequest{RequestContext: commandRC(), NewPassword: next}
		for range 2 {
			_, err := h.client.ChangePassword(bearer(tokens.GetAccessToken()), request)
			wantOK(t, err)
		}
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'password_changed'"); n != 1 {
			t.Fatalf("%d password changes, want 1", n)
		}
		_, err = h.signIn(email, next)
		wantOK(t, err)
	})
	t.Run("SEC-13 two simultaneous ChangePassword with one grant", func(t *testing.T) {
		// One grant serves one action under concurrency; the serializable
		// retry decides which.
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		_, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantOK(t, err)
		var wg sync.WaitGroup
		results := make([]error, 2)
		for i, newPassword := range []string{next, strings.Repeat("other ", 3)} {
			wg.Go(func() { results[i] = h.changePassword(tokens.GetAccessToken(), newPassword) })
		}
		wg.Wait()
		ok := slices.IndexFunc(results, func(err error) bool { return err == nil })
		if ok < 0 || results[1-ok] == nil {
			t.Fatalf("results %v, want exactly one success", results)
		}
		wantErrorCode(t, results[1-ok], connect.CodePermissionDenied, reauthRequired)
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'password_changed'"); n != 1 {
			t.Fatalf("%d password_changed events, want 1", n)
		}
	})
	t.Run("SEC-13 a re-authentication serves one action", func(t *testing.T) {
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		_, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantOK(t, err)
		wantOK(t, h.changePassword(tokens.GetAccessToken(), next))
		wantErrorCode(t, h.changePassword(tokens.GetAccessToken(), password), connect.CodePermissionDenied, reauthRequired)
	})
	t.Run("a re-authentication is bound to its session and purpose", func(t *testing.T) {
		h := startIdentity(t)
		first := h.signUp(t, email, password)
		second, err := h.signIn(email, password)
		wantOK(t, err)
		_, err = h.reauthenticate(first.GetAccessToken(), password)
		wantOK(t, err)
		wantErrorCode(t, h.changePassword(second.GetAccessToken(), next), connect.CodePermissionDenied, reauthRequired)
		for _, purpose := range []casav1.ReauthPurpose{casav1.ReauthPurpose_REAUTH_PURPOSE_UNSPECIFIED, 99} {
			_, err := h.client.Reauthenticate(bearer(second.GetAccessToken()), &casav1.ReauthenticateRequest{
				RequestContext: identityRC(ios), Password: password, Purpose: purpose,
			})
			wantCode(t, err, connect.CodeInvalidArgument)
		}
		wantOK(t, h.changePassword(first.GetAccessToken(), next))
	})
	t.Run("a re-authentication needs a session", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		_, err := h.reauthenticate("", password)
		wantCode(t, err, connect.CodeUnauthenticated)
	})
	t.Run("a wrong password at re-authentication is rejected and rate limited", func(t *testing.T) {
		// SEC-12 cited, not named; 10.1 §8.1 row Reauthenticate.
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		for range 20 {
			_, err := h.reauthenticate(tokens.GetAccessToken(), "a wrong passphrase")
			wantErrorCode(t, err, connect.CodeUnauthenticated, casav1.ErrorCode_ERROR_CODE_CREDENTIALS_REJECTED)
		}
		_, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantErrorCode(t, err, connect.CodeResourceExhausted, casav1.ErrorCode_ERROR_CODE_RATE_LIMITED)
		if n := h.count(t, "SELECT count(*) FROM reauth_grant"); n != 0 {
			t.Fatalf("%d grants, want none", n)
		}
	})
	t.Run("after a password change only the new password signs in", func(t *testing.T) {
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		_, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantOK(t, err)
		wantOK(t, h.changePassword(tokens.GetAccessToken(), next))
		_, err = h.signIn(email, password)
		wantErrorCode(t, err, connect.CodeUnauthenticated, casav1.ErrorCode_ERROR_CODE_CREDENTIALS_REJECTED)
		_, err = h.signIn(email, next)
		wantOK(t, err)
	})
	t.Run("a new password on the common list is refused", func(t *testing.T) {
		// SEC-11; the grant is not consumed.
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		_, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantOK(t, err)
		err = h.changePassword(tokens.GetAccessToken(), "password1234")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violations(t, err), []string{"new_password password.common"}; !slices.Equal(got, want) {
			t.Fatalf("violations %v, want %v", got, want)
		}
		err = h.changePassword(tokens.GetAccessToken(), "elevenchars")
		if got, want := violations(t, err), []string{"new_password string.min_len"}; !slices.Equal(got, want) {
			t.Fatalf("violations %v, want %v", got, want)
		}
		wantOK(t, h.changePassword(tokens.GetAccessToken(), next))
	})
	t.Run("re-authentication and a password change are written to the account's security trail", func(t *testing.T) {
		// FR-FAM-022/AC-5 cited, not named: its screen comes later.
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		_, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantOK(t, err)
		wantOK(t, h.changePassword(tokens.GetAccessToken(), next))
		var kinds []string
		rows, err := h.admin.Query(t.Context(), `SELECT kind || ':' || coalesce(detail, '') FROM security_event
			WHERE session_id IS NOT NULL AND kind <> 'sign_in' ORDER BY occurred_at, kind DESC`)
		if err == nil {
			kinds, err = pgx.CollectRows(rows, pgx.RowTo[string])
		}
		if want := []string{"reauthentication:change_password", "password_changed:"}; err != nil || !slices.Equal(kinds, want) {
			t.Fatalf("events %v (%v), want %v", kinds, err, want)
		}
	})
	t.Run("the change-password request leaves no password trace", func(t *testing.T) {
		// SEC-11; decision 8 of the 1.4 plan.
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		_, err := h.reauthenticate(tokens.GetAccessToken(), password)
		wantOK(t, err)
		wantOK(t, h.changePassword(tokens.GetAccessToken(), next))
		cleared := sha256.Sum256(nil) // the request without request_context and new_password
		if n := h.count(t, "SELECT count(*) FROM idempotency_record WHERE request_hash = $1", cleared[:]); n != 1 {
			t.Fatalf("%d records with the cleared request hash, want 1", n)
		}
		sum := sha256.Sum256([]byte(next))
		traces := []string{next, fmt.Sprintf("%x", sum)}
		rows, err := h.admin.Query(t.Context(), "SELECT tablename FROM pg_tables WHERE schemaname = 'public'")
		if err != nil {
			t.Fatal(err)
		}
		tables, err := pgx.CollectRows(rows, pgx.RowTo[string])
		if err != nil || len(tables) < 10 {
			t.Fatalf("tables %v, %v", tables, err)
		}
		found := func(trace string) (holding []string) {
			for _, table := range tables {
				query := fmt.Sprintf("SELECT count(*) FROM %s t WHERE strpos(t::text, $1) > 0", pgx.Identifier{table}.Sanitize())
				if h.count(t, query, trace) != 0 {
					holding = append(holding, table)
				}
			}
			return holding
		}
		if got := found(email); len(got) == 0 { // the scan finds what is there
			t.Fatal("the scan does not find the email")
		}
		for _, trace := range traces {
			if got := found(trace); len(got) != 0 {
				t.Errorf("tables %v hold a trace of the new password", got)
			}
		}
	})
}

// startRecovery calls StartRecovery and waits for its mail, which leaves
// the request path.
func (h *identityHarness) startRecovery(email string) (*casav1.StartRecoveryResponse, error) {
	res, err := h.client.StartRecovery(context.Background(), &casav1.StartRecoveryRequest{RequestContext: identityRC(ios), Email: email})
	h.service.Wait()
	return res, err
}

func (h *identityHarness) completeRecovery(kind casav1.ClientKind, token, newPassword string) (*casav1.SessionTokens, error) {
	res, err := h.client.CompleteRecovery(context.Background(), &casav1.CompleteRecoveryRequest{
		RequestContext: identityRC(kind), RecoveryToken: token, NewPassword: newPassword,
	})
	return res.GetTokens(), err
}

func TestRecovery(t *testing.T) {
	const email = "owner@example.test"
	// Built, so no secret scanner sees a literal.
	password, next := strings.Repeat("recover ", 3), strings.Repeat("recovered ", 3)
	rejected := casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED
	rateLimited := casav1.ErrorCode_ERROR_CODE_RATE_LIMITED

	t.Run("SEC-24 a recovery link works once", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		_, err := h.startRecovery(email)
		wantOK(t, err)
		mails := h.mails()
		last := mails[len(mails)-1]
		if len(mails) != 2 || last.To != email || last.Kind != identity.MessageRecovery ||
			!strings.HasPrefix(last.Link, testOrigin+"/recover#") || len(h.lastLink(t)) != 43 {
			t.Fatalf("mails %d, last to %q kind %d link of %d characters; want a recovery link to %s", len(mails), last.To, last.Kind, len(last.Link), email)
		}
		link := h.lastLink(t)
		tokens, err := h.completeRecovery(ios, link, next)
		wantOK(t, err)
		wantOK(t, h.authenticated(tokens.GetAccessToken()))
		_, again := h.completeRecovery(ios, link, password)
		_, unknown := h.completeRecovery(ios, strings.Repeat("A", 43), password)
		wantErrorCode(t, again, connect.CodePermissionDenied, rejected)
		if !rejectionOf(t, again).equal(rejectionOf(t, unknown)) {
			t.Fatalf("a used link %v and an unknown one %v are told apart", again, unknown)
		}
		_, err = h.signIn(email, next)
		wantOK(t, err)
	})
	t.Run("SEC-24 after recovery the old sessions cannot refresh", func(t *testing.T) {
		h := startIdentity(t)
		first := h.signUp(t, email, password)
		second, err := h.signIn(email, password)
		wantOK(t, err)
		_, err = h.startRecovery(email)
		wantOK(t, err)
		tokens, err := h.completeRecovery(ios, h.lastLink(t), next)
		wantOK(t, err)
		revoked := casav1.ErrorCode_ERROR_CODE_SESSION_REVOKED
		for _, old := range []*casav1.SessionTokens{first, second} {
			_, err := h.refresh(old.GetRefreshToken())
			wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
			wantErrorCode(t, h.authenticated(old.GetAccessToken()), connect.CodeUnauthenticated, revoked)
		}
		wantOK(t, h.authenticated(tokens.GetAccessToken()))
		if n := h.count(t, "SELECT count(*) FROM session WHERE revoke_reason = 'recovery'"); n != 2 {
			t.Fatalf("%d sessions revoked by recovery, want 2", n)
		}
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'session_revoked' AND detail = 'recovery'"); n != 2 {
			t.Fatalf("%d session_revoked events, want 2", n)
		}
		// SEC-24: the owner reads when and from which device.
		if n := h.count(t, `SELECT count(*) FROM security_event e JOIN session s ON s.id = e.session_id
			WHERE e.kind = 'recovery' AND s.revoked_at IS NULL AND s.client_kind = 'ios'`); n != 1 {
			t.Fatalf("%d recovery events naming the new session, want 1", n)
		}
	})
	t.Run("SEC-24 recovery revokes expired sessions too but trails only live ones", func(t *testing.T) {
		// Fail closed (lead ruling on #32): every unrevoked session is
		// revoked, so no path can revive an expired one; the trail names
		// only the sessions that were still live (code-reviewer M1).
		h := startIdentity(t)
		h.signUp(t, email, password)
		h.clock.advance(30 * 24 * time.Hour) // the first session ends idle
		_, err := h.signIn(email, password)
		wantOK(t, err)
		_, err = h.startRecovery(email)
		wantOK(t, err)
		_, err = h.completeRecovery(ios, h.lastLink(t), next)
		wantOK(t, err)
		if n := h.count(t, "SELECT count(*) FROM session WHERE revoke_reason = 'recovery'"); n != 2 {
			t.Fatalf("%d sessions revoked by recovery, want 2 (the expired one too)", n)
		}
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'session_revoked'"); n != 1 {
			t.Fatalf("%d session_revoked events, want 1", n)
		}
	})
	t.Run("SEC-24 a recovery link expires after 30 minutes", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		_, err := h.startRecovery(email)
		wantOK(t, err)
		link := h.lastLink(t)
		h.clock.advance(30 * time.Minute)
		_, err = h.completeRecovery(ios, link, next)
		wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
		_, err = h.signIn(email, password)
		wantOK(t, err)
	})
	t.Run("SEC-30 recovery answers alike for known and unknown email", func(t *testing.T) {
		// SEC-12; lead ruling on PR B: the answer does not depend on the
		// mail, which leaves the request path.
		h := startIdentity(t)
		h.signUp(t, email, password)
		unknown, err := h.startRecovery("nobody@example.test")
		wantOK(t, err)
		known, err := h.startRecovery(email)
		wantOK(t, err)
		if mails := h.mails(); len(mails) != 2 || mails[1].To != email {
			t.Fatalf("%d mails, want the sign-up link and one recovery link to %s", len(mails), email)
		}
		h.failMail(errors.New("mail down"))
		failed, err := h.startRecovery(email)
		wantOK(t, err)
		for _, res := range []*casav1.StartRecoveryResponse{known, failed} {
			if a, b := marshal(t, res), marshal(t, unknown); !bytes.Equal(a, b) {
				t.Fatalf("responses %x and %x differ", a, b)
			}
		}
		logs := h.logs.String()
		if !strings.Contains(logs, `"msg":"recovery mail not sent"`) || strings.Contains(logs, "owner") || strings.Contains(logs, "mail down") {
			t.Fatalf("logs %q, want the failure logged without the email or the error text", logs)
		}
	})
	t.Run("SEC-12 recovery mail for one email stops after 5 requests", func(t *testing.T) {
		// 10.1 §8.1 row StartRecovery: every request counts; nothing clears.
		h := startIdentity(t)
		h.signUp(t, email, password)
		for _, address := range []string{email, "nobody@example.test"} {
			for i := range 5 {
				_, err := h.startRecovery(address)
				wantOK(t, err)
				if i == 0 && address == email {
					_, err := h.completeRecovery(ios, h.lastLink(t), next) // a success clears nothing
					wantOK(t, err)
				}
			}
			_, err := h.startRecovery(strings.ToUpper(address))
			wantErrorCode(t, err, connect.CodeResourceExhausted, rateLimited)
		}
		if mails := h.mails(); len(mails) != 1+5 {
			t.Fatalf("%d mails, want the sign-up link and 5 recovery links", len(mails))
		}
	})
	t.Run("SEC-12 recovery requests are limited per address", func(t *testing.T) {
		h := startIdentity(t)
		for i := range 20 {
			_, err := h.startRecovery(fmt.Sprintf("nobody%d@example.test", i))
			wantOK(t, err)
		}
		_, err := h.startRecovery("someone@example.test")
		wantErrorCode(t, err, connect.CodeResourceExhausted, rateLimited)
	})
	t.Run("SEC-12 a sign-in lockout does not block recovery", func(t *testing.T) {
		// 10.1 §8.1 "Penguncian oleh orang lain".
		h := startIdentity(t)
		h.signUp(t, email, password)
		for range 20 {
			_, err := h.signIn(email, "a wrong passphrase")
			wantErrorCode(t, err, connect.CodeUnauthenticated, casav1.ErrorCode_ERROR_CODE_CREDENTIALS_REJECTED)
		}
		_, err := h.signIn(email, password)
		wantErrorCode(t, err, connect.CodeResourceExhausted, rateLimited)
		_, err = h.startRecovery(email)
		wantOK(t, err)
		tokens, err := h.completeRecovery(ios, h.lastLink(t), next)
		wantOK(t, err)
		wantOK(t, h.authenticated(tokens.GetAccessToken()))
	})
	t.Run("SEC-12 failed token uses are limited per address", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password) // one successful token use; it is withdrawn, not counted
		_, err := h.startRecovery(email)
		wantOK(t, err)
		link := h.lastLink(t)
		for i := range 20 {
			_, err := h.completeRecovery(ios, fmt.Sprintf("%043d", i), next)
			wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
		}
		_, err = h.completeRecovery(ios, link, next)
		wantErrorCode(t, err, connect.CodeResourceExhausted, rateLimited)
	})
	t.Run("SEC-24 a reset uses up every other recovery link", func(t *testing.T) {
		// Lead ruling on PR B: a recovery and a password change each leave
		// no older recovery link usable.
		h := startIdentity(t)
		h.signUp(t, email, password)
		links := make([]string, 3)
		for i := range links {
			_, err := h.startRecovery(email)
			wantOK(t, err)
			links[i] = h.lastLink(t)
		}
		_, err := h.completeRecovery(ios, links[1], next)
		wantOK(t, err)
		for _, i := range []int{0, 2} {
			_, err := h.completeRecovery(ios, links[i], password)
			wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
		}
		_, err = h.startRecovery(email)
		wantOK(t, err)
		link := h.lastLink(t)
		tokens, err := h.signIn(email, next)
		wantOK(t, err)
		_, err = h.reauthenticate(tokens.GetAccessToken(), next)
		wantOK(t, err)
		wantOK(t, h.changePassword(tokens.GetAccessToken(), password))
		_, err = h.completeRecovery(ios, link, next)
		wantErrorCode(t, err, connect.CodePermissionDenied, rejected)
		_, err = h.signIn(email, password)
		wantOK(t, err)
	})
	t.Run("SEC-11 a recovery with a common password is refused and the link stays usable", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		_, err := h.startRecovery(email)
		wantOK(t, err)
		link := h.lastLink(t)
		for weak, want := range map[string]string{"password1234": "new_password password.common", "elevenchars": "new_password string.min_len"} {
			_, err := h.completeRecovery(ios, link, weak)
			wantCode(t, err, connect.CodeInvalidArgument)
			if got := violations(t, err); !slices.Equal(got, []string{want}) {
				t.Fatalf("violations %v, want [%s]", got, want)
			}
		}
		_, err = h.completeRecovery(ios, link, next)
		wantOK(t, err)
	})
}

func marshal(t *testing.T, m proto.Message) []byte {
	t.Helper()
	b, err := proto.MarshalOptions{Deterministic: true}.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestChangePasswordRevokesOtherSessions(t *testing.T) {
	// SEC-24 (security-privacy decision on #32): a password change revokes
	// every other live session of the account in its transaction and keeps
	// the one it came from.
	const email = "owner@example.test"
	password, next := strings.Repeat("reauth ", 3), strings.Repeat("changed ", 3)
	revoked := casav1.ErrorCode_ERROR_CODE_SESSION_REVOKED
	h := startIdentity(t)
	x := h.signUp(t, email, password)
	y := h.webSignIn(t, email, password, personal) // a web session in its cookie
	z, err := h.signIn(email, password)
	wantOK(t, err)
	_, err = h.reauthenticate(x.GetAccessToken(), password)
	wantOK(t, err)
	wantOK(t, h.changePassword(x.GetAccessToken(), next))
	wantOK(t, h.authenticated(x.GetAccessToken()))
	_, err = h.refresh(x.GetRefreshToken())
	wantOK(t, err)
	wantErrorCode(t, h.authenticated(z.GetAccessToken()), connect.CodeUnauthenticated, revoked)
	_, err = h.refresh(z.GetRefreshToken())
	wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
	wantErrorCode(t, h.read(sameOrigin(y), web, true), connect.CodeUnauthenticated, revoked)
	if n := h.count(t, `SELECT count(*) FROM security_event WHERE kind = 'session_revoked' AND detail = 'password_changed'`); n != 2 {
		t.Fatalf("%d session_revoked rows, want 2", n)
	}
	if n := h.count(t, "SELECT count(*) FROM session WHERE revoke_reason = 'password_changed'"); n != 2 {
		t.Fatalf("%d sessions revoked by the change, want 2", n)
	}
}

// decodeToken returns the 32 random bytes of a token (base64url, no padding).
func decodeToken(t *testing.T, token string) []byte {
	t.Helper()
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}
