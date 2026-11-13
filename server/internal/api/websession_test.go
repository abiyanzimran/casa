package api_test

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"slices"
	"strings"
	"testing"
	"time"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"connectrpc.com/connect/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/emptypb"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
)

const sessionCookieName = "__Host-casa_session"

var (
	web      = casav1.ClientKind_CLIENT_KIND_WEB
	personal = casav1.DeviceMode_DEVICE_MODE_PERSONAL
	shared   = casav1.DeviceMode_DEVICE_MODE_SHARED
)

// safariIPhone is the User-Agent of Safari on an iPhone.
const safariIPhone = "Mozilla/5.0 (iPhone; CPU iPhone OS 18_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Mobile/15E148 Safari/604.1"

// webCall is one call's context with userAgent and the app's Origin, as a
// browser sends them, and the call's info, which holds the response headers
// once the call returns.
func webCall(userAgent string) (context.Context, *connect.CallInfo) {
	ctx, info := connect.NewClientContext(context.Background())
	info.RequestHeader().Set("User-Agent", userAgent)
	info.RequestHeader().Set("Origin", testOrigin)
	return ctx, info
}

// sessionCookie returns the one Set-Cookie of a response, parsed.
func sessionCookie(t *testing.T, info *connect.CallInfo) (*http.Cookie, string) {
	t.Helper()
	lines := info.ResponseHeader().Values("Set-Cookie")
	if len(lines) != 1 {
		t.Fatalf("%d Set-Cookie headers, want 1", len(lines))
	}
	cookie, err := http.ParseSetCookie(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	return cookie, lines[0]
}

// wantHostCookie checks the attributes every session cookie has (plan 1.5
// decision 7; SEC-161 (b)).
func wantHostCookie(t *testing.T, cookie *http.Cookie, line string) {
	t.Helper()
	if cookie.Name != sessionCookieName || len(cookie.Value) != 43 || !cookie.Secure || !cookie.HttpOnly ||
		cookie.SameSite != http.SameSiteStrictMode || cookie.Path != "/" || cookie.Domain != "" ||
		strings.Contains(strings.ToLower(line), "domain") {
		t.Fatalf("Set-Cookie %q: want %s, a 43-character value, Secure, HttpOnly, SameSite=Strict, Path=/ and no Domain",
			cookieShape(line), sessionCookieName)
	}
}

// cookieShape is a Set-Cookie line without its value, for a failure message.
func cookieShape(line string) string {
	name, rest, _ := strings.Cut(line, "=")
	_, attributes, _ := strings.Cut(rest, ";")
	return name + "=<value>;" + attributes
}

// sessionRow is what the session table holds about the cookie's session.
type sessionRow struct {
	clientKind, deviceMode, uaSummary string
	idle                              time.Duration
	absolute                          time.Time
	refreshTokens                     int
}

func (h *identityHarness) sessionOfCookie(t *testing.T, value string) sessionRow {
	t.Helper()
	hash := sha256.Sum256(decodeToken(t, value))
	var r sessionRow
	err := h.admin.QueryRow(t.Context(), `
		SELECT client_kind, device_mode, ua_summary, idle_timeout, absolute_expiry,
		       (SELECT count(*) FROM refresh_token WHERE session_id = s.id)
		  FROM session s WHERE token_hash = $1`, hash[:]).
		Scan(&r.clientKind, &r.deviceMode, &r.uaSummary, &r.idle, &r.absolute, &r.refreshTokens)
	if err != nil {
		t.Fatalf("session of the cookie: %v", err)
	}
	return r
}

// wantNativeTokens checks a native response: tokens, and no Set-Cookie
// (SEC-161 (b): the cookie is only for web).
func wantNativeTokens(t *testing.T, info *connect.CallInfo, tokens *casav1.SessionTokens) {
	t.Helper()
	if n := len(info.ResponseHeader().Values("Set-Cookie")); n != 0 || tokens.GetAccessToken() == "" || tokens.GetRefreshToken() == "" {
		t.Fatalf("native response: %d Set-Cookie headers, tokens %t; want none and tokens", n, tokens.GetAccessToken() != "")
	}
}

// ruleIDs returns the rule id of every violation in err.
func ruleIDs(t *testing.T, err error) []string {
	t.Helper()
	var got []string
	for _, detail := range details(t, err) {
		if v, ok := detail.(*validate.Violations); ok {
			for _, violation := range v.GetViolations() {
				got = append(got, violation.GetRuleId())
			}
		}
	}
	return got
}

func TestWebSessionOpening(t *testing.T) {
	password, next := strings.Repeat("web session ", 2), strings.Repeat("web recovered ", 2) // built, so no secret scanner sees a literal
	const email = "owner@example.test"
	const day = 24 * time.Hour

	t.Run("SEC-161 (b) a web sign-in sets only the __Host- cookie", func(t *testing.T) {
		h := startIdentity(t)
		// The native variants of all three RPCs answer with tokens and set
		// no cookie, even with a browser's User-Agent.
		wantOK(t, h.startSignUp(h.issue(t, email), email))
		ctx, info := webCall(safariIPhone)
		signedUp, err := h.client.CompleteSignUp(ctx, &casav1.CompleteSignUpRequest{
			RequestContext: identityRC(ios), VerificationToken: h.lastLink(t), Password: password,
		})
		wantOK(t, err)
		wantNativeTokens(t, info, signedUp.GetTokens())
		ctx, info = webCall(safariIPhone)
		signedIn, err := h.client.SignIn(ctx, &casav1.SignInRequest{
			RequestContext: identityRC(ios), Email: email, Password: password,
		})
		wantOK(t, err)
		wantNativeTokens(t, info, signedIn.GetTokens())
		_, err = h.startRecovery(email)
		wantOK(t, err)
		ctx, info = webCall(safariIPhone)
		recovered, err := h.client.CompleteRecovery(ctx, &casav1.CompleteRecoveryRequest{
			RequestContext: identityRC(ios), RecoveryToken: h.lastLink(t), NewPassword: next,
		})
		wantOK(t, err)
		wantNativeTokens(t, info, recovered.GetTokens())
		password := next

		ctx, info = webCall(safariIPhone)
		res, err := h.client.SignIn(ctx, &casav1.SignInRequest{
			RequestContext: identityRC(web), Email: email, Password: password, DeviceMode: personal,
		})
		wantOK(t, err)
		if res.GetTokens() != nil {
			t.Fatal("a web sign-in answered with tokens; a web session travels only in the cookie")
		}
		cookie, line := sessionCookie(t, info)
		wantHostCookie(t, cookie, line)
		// Plan 1.5 decision 7: Max-Age is the remaining absolute lifetime.
		if cookie.MaxAge != int((180 * day).Seconds()) {
			t.Fatalf("Max-Age %d, want %d (180 days)", cookie.MaxAge, int((180 * day).Seconds()))
		}
		row := h.sessionOfCookie(t, cookie.Value)
		// SEC-161 (d): personal mode idles out after 30 days, ends after 180.
		if row.clientKind != "web" || row.deviceMode != "personal" || row.idle != 30*day ||
			!row.absolute.Equal(h.clock.Now().Add(180*day)) || row.refreshTokens != 0 {
			t.Fatalf("session %+v; want web, personal, 30 days idle, 180 days absolute, no refresh token", row)
		}
	})

	t.Run("a shared-mode sign-in sets a session cookie without Max-Age or Expires", func(t *testing.T) {
		// SEC-165 (a): shared mode ends with the browser, after 30 idle
		// minutes, or after 12 hours.
		h := startIdentity(t)
		h.signUp(t, email, password)
		ctx, info := webCall(safariIPhone)
		res, err := h.client.SignIn(ctx, &casav1.SignInRequest{
			RequestContext: identityRC(web), Email: email, Password: password, DeviceMode: shared,
		})
		wantOK(t, err)
		cookie, line := sessionCookie(t, info)
		wantHostCookie(t, cookie, line)
		if res.GetTokens() != nil || cookie.MaxAge != 0 || cookie.RawExpires != "" ||
			strings.Contains(strings.ToLower(line), "max-age") || strings.Contains(strings.ToLower(line), "expires") {
			t.Fatalf("Set-Cookie %q or tokens in the response; want a session cookie only", cookieShape(line))
		}
		row := h.sessionOfCookie(t, cookie.Value)
		if row.clientKind != "web" || row.deviceMode != "shared" || row.idle != 30*time.Minute ||
			!row.absolute.Equal(h.clock.Now().Add(12*time.Hour)) || row.refreshTokens != 0 {
			t.Fatalf("session %+v; want web, shared, 30 minutes idle, 12 hours absolute, no refresh token", row)
		}
	})

	t.Run("a web client must choose a device mode and a native client must not", func(t *testing.T) {
		// Plan 1.5 decision 9; SEC-165: no preselected mode.
		h := startIdentity(t)
		h.signUp(t, email, password)
		token := strings.Repeat("A", 43)
		for _, c := range []struct {
			name string
			kind casav1.ClientKind
			mode casav1.DeviceMode
			rule string
		}{
			{"web without a mode", web, casav1.DeviceMode_DEVICE_MODE_UNSPECIFIED, "device_mode.web"},
			{"web with an undefined mode", web, casav1.DeviceMode(7), "enum.defined_only"},
			{"ios in shared mode", ios, shared, "device_mode.web"},
			{"ios in personal mode", ios, personal, "device_mode.web"},
		} {
			ctx, _ := webCall(safariIPhone) // the Origin a web call needs (SEC-162 (a))
			for name, call := range map[string]func() error{
				"SignIn": func() error {
					_, err := h.client.SignIn(ctx, &casav1.SignInRequest{
						RequestContext: identityRC(c.kind), Email: email, Password: password, DeviceMode: c.mode,
					})
					return err
				},
				"CompleteSignUp": func() error {
					_, err := h.client.CompleteSignUp(ctx, &casav1.CompleteSignUpRequest{
						RequestContext: identityRC(c.kind), VerificationToken: token, Password: password, DeviceMode: c.mode,
					})
					return err
				},
				"CompleteRecovery": func() error {
					_, err := h.client.CompleteRecovery(ctx, &casav1.CompleteRecoveryRequest{
						RequestContext: identityRC(c.kind), RecoveryToken: token, NewPassword: next, DeviceMode: c.mode,
					})
					return err
				},
			} {
				err := call()
				wantCode(t, err, connect.CodeInvalidArgument)
				if got := ruleIDs(t, err); !slices.Equal(got, []string{c.rule}) {
					t.Errorf("%s, %s: rules %v, want [%s]", c.name, name, got, c.rule)
				}
			}
		}
		if n := h.count(t, "SELECT count(*) FROM session"); n != 1 {
			t.Fatalf("%d sessions, want only the sign-up's", n)
		}
	})

	t.Run("web sign-up completion and web recovery open a cookie session", func(t *testing.T) {
		h := startIdentity(t)
		wantOK(t, h.startSignUp(h.issue(t, email), email))
		ctx, info := webCall(safariIPhone)
		signedUp, err := h.client.CompleteSignUp(ctx, &casav1.CompleteSignUpRequest{
			RequestContext: identityRC(web), VerificationToken: h.lastLink(t), Password: password, DeviceMode: personal,
		})
		wantOK(t, err)
		cookie, line := sessionCookie(t, info)
		wantHostCookie(t, cookie, line)
		if signedUp.GetTokens() != nil || cookie.MaxAge != int((180*day).Seconds()) {
			t.Fatalf("sign-up: tokens %t, Max-Age %d; want no tokens and 180 days", signedUp.GetTokens() != nil, cookie.MaxAge)
		}
		if row := h.sessionOfCookie(t, cookie.Value); row.clientKind != "web" || row.deviceMode != "personal" || row.refreshTokens != 0 {
			t.Fatalf("sign-up session %+v; want web, personal, no refresh token", row)
		}

		_, err = h.startRecovery(email)
		wantOK(t, err)
		ctx, info = webCall(safariIPhone)
		recovered, err := h.client.CompleteRecovery(ctx, &casav1.CompleteRecoveryRequest{
			RequestContext: identityRC(web), RecoveryToken: h.lastLink(t), NewPassword: next, DeviceMode: shared,
		})
		wantOK(t, err)
		recoveredCookie, line := sessionCookie(t, info)
		wantHostCookie(t, recoveredCookie, line)
		if recovered.GetTokens() != nil || recoveredCookie.MaxAge != 0 || recoveredCookie.Value == cookie.Value {
			t.Fatalf("recovery: tokens %t, Max-Age %d; want no tokens, a session cookie and a new value",
				recovered.GetTokens() != nil, recoveredCookie.MaxAge)
		}
		if row := h.sessionOfCookie(t, recoveredCookie.Value); row.clientKind != "web" || row.deviceMode != "shared" || row.refreshTokens != 0 {
			t.Fatalf("recovery session %+v; want web, shared, no refresh token", row)
		}
		if n := h.count(t, "SELECT count(*) FROM session WHERE revoke_reason = 'recovery'"); n != 1 {
			t.Fatalf("%d sessions revoked by the recovery, want the sign-up's", n)
		}
	})

	t.Run("the cookie value is in no response body, table or summary", func(t *testing.T) {
		// SEC-161 (b), (e); SEC-44. net/http refuses a header value with a
		// control character other than tab before any handler, so the tab is
		// the control character that can reach the server.
		h := startIdentity(t)
		marker := "casa-marker-" + uuid4()
		userAgent := safariIPhone + " " + marker + "\t<script>"
		var values []string
		var bodies [][]byte
		keep := func(info *connect.CallInfo, res proto.Message) {
			cookie, _ := sessionCookie(t, info)
			values = append(values, cookie.Value)
			bodies = append(bodies, marshal(t, res))
		}

		wantOK(t, h.startSignUp(h.issue(t, email), email))
		ctx, info := webCall(userAgent)
		signedUp, err := h.client.CompleteSignUp(ctx, &casav1.CompleteSignUpRequest{
			RequestContext: identityRC(web), VerificationToken: h.lastLink(t), Password: password, DeviceMode: personal,
		})
		wantOK(t, err)
		keep(info, signedUp)
		ctx, info = webCall(userAgent)
		signedIn, err := h.client.SignIn(ctx, &casav1.SignInRequest{
			RequestContext: identityRC(web), Email: email, Password: password, DeviceMode: shared,
		})
		wantOK(t, err)
		keep(info, signedIn)
		_, err = h.startRecovery(email)
		wantOK(t, err)
		ctx, info = webCall(userAgent)
		recovered, err := h.client.CompleteRecovery(ctx, &casav1.CompleteRecoveryRequest{
			RequestContext: identityRC(web), RecoveryToken: h.lastLink(t), NewPassword: next, DeviceMode: personal,
		})
		wantOK(t, err)
		keep(info, recovered)

		var forbidden []string
		for _, value := range values {
			raw := decodeToken(t, value)
			forbidden = append(forbidden, value, hex.EncodeToString(raw), base64.StdEncoding.EncodeToString(raw))
		}
		for i, body := range bodies {
			for _, f := range forbidden {
				if strings.Contains(string(body), f) {
					t.Errorf("response body %d holds a cookie value", i)
				}
			}
		}
		forbidden = append(forbidden, marker)
		for _, table := range []string{
			"account", "credential", "session", "refresh_token", "reauth_grant", "recovery_token",
			"setup_token", "signup_verification", "security_event", "idempotency_record",
		} {
			var dump string
			if err := h.admin.QueryRow(t.Context(), "SELECT coalesce(string_agg(t::text, E'\\n'), '') FROM "+table+" t").Scan(&dump); err != nil {
				t.Fatal(err)
			}
			for _, f := range forbidden {
				if strings.Contains(dump, f) {
					t.Errorf("table %s holds a cookie value or the raw User-Agent", table)
				}
			}
		}
		if logs := h.logs.String(); slices.ContainsFunc(forbidden, func(f string) bool { return strings.Contains(logs, f) }) {
			t.Error("the service log holds a cookie value or the raw User-Agent")
		}
		if n := h.count(t, "SELECT count(*) FROM session WHERE client_kind = 'web' AND ua_summary = 'Safari · iOS'"); n != 3 {
			t.Fatalf("%d web sessions summarized as Safari · iOS, want 3", n)
		}
	})

	t.Run("ua_summary is filled for web and empty for native", func(t *testing.T) {
		// SEC-161 (e); plan 1.5 decision 10.
		h := startIdentity(t)
		wantOK(t, h.startSignUp(h.issue(t, email), email))
		ctx, _ := webCall(safariIPhone)
		_, err := h.client.CompleteSignUp(ctx, &casav1.CompleteSignUpRequest{
			RequestContext: identityRC(ios), VerificationToken: h.lastLink(t), Password: password,
		})
		wantOK(t, err)
		ctx, info := webCall(safariIPhone)
		_, err = h.client.SignIn(ctx, &casav1.SignInRequest{
			RequestContext: identityRC(web), Email: email, Password: password, DeviceMode: personal,
		})
		wantOK(t, err)
		cookie, _ := sessionCookie(t, info)
		if row := h.sessionOfCookie(t, cookie.Value); row.uaSummary != "Safari · iOS" {
			t.Fatalf("web ua_summary %q, want %q", row.uaSummary, "Safari · iOS")
		}
		if n := h.count(t, "SELECT count(*) FROM session WHERE client_kind = 'ios' AND ua_summary = '' AND device_mode = 'personal'"); n != 1 {
			t.Fatalf("%d native sessions with an empty ua_summary in personal mode, want 1", n)
		}
	})
}

// cookieCall is a call's context carrying the session cookie value (none
// for ""), and the Origin and Sec-Fetch-Site headers (none for "").
func cookieCall(cookie, origin, fetchSite string) context.Context {
	ctx, info := connect.NewClientContext(context.Background())
	if cookie != "" {
		info.RequestHeader().Set("Cookie", sessionCookieName+"="+cookie)
	}
	if origin != "" {
		info.RequestHeader().Set("Origin", origin)
	}
	if fetchSite != "" {
		info.RequestHeader().Set("Sec-Fetch-Site", fetchSite)
	}
	return ctx
}

// sameOrigin is a browser's call from the app itself with the cookie.
func sameOrigin(cookie string) context.Context { return cookieCall(cookie, testOrigin, "same-origin") }

// webSignIn signs in from the web in mode and returns the cookie's value.
func (h *identityHarness) webSignIn(t *testing.T, email, password string, mode casav1.DeviceMode) string {
	t.Helper()
	ctx, info := webCall(safariIPhone)
	_, err := h.client.SignIn(ctx, &casav1.SignInRequest{
		RequestContext: identityRC(web), Email: email, Password: password, DeviceMode: mode,
	})
	wantOK(t, err)
	cookie, _ := sessionCookie(t, info)
	return cookie.Value
}

// read calls the authenticated ProbeService.Read as kind; userInitiated is
// request_context.user_initiated (SEC-165 (g)).
func (h *identityHarness) read(ctx context.Context, kind casav1.ClientKind, userInitiated bool) error {
	rc := identityRC(kind)
	rc.UserInitiated = userInitiated
	req := dynamicpb.NewMessage(h.probeRead.Input())
	req.Set(h.probeRead.Input().Fields().ByName("request_context"), protoreflect.ValueOfMessage(rc.ProtoReflect()))
	return h.probe.CallUnary(ctx, connect.Spec{Procedure: "/casa.apitest.v1.ProbeService/Read"}, req, &emptypb.Empty{})
}

// webSignOut is SignOut from a web client in ctx.
func (h *identityHarness) webSignOut(ctx context.Context) error {
	rc := commandRC()
	rc.ClientKind = web
	_, err := h.client.SignOut(ctx, &casav1.SignOutRequest{RequestContext: rc})
	return err
}

// wantCrossOriginRejected checks the one answer of the origin check: no
// detail and a fixed message (plan 1.5 decision 11 step 2).
func wantCrossOriginRejected(t *testing.T, err error) {
	t.Helper()
	wantCode(t, err, connect.CodePermissionDenied)
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) || connectErr.Message() != "cross-origin request rejected" || len(connectErr.Details()) != 0 {
		t.Fatalf("error %v; want \"cross-origin request rejected\" without detail", err)
	}
}

func TestCookieSession(t *testing.T) {
	password := strings.Repeat("cookie session ", 2) // built, so no secret scanner sees a literal
	const email = "owner@example.test"
	const day = 24 * time.Hour
	expired := casav1.ErrorCode_ERROR_CODE_SESSION_EXPIRED
	revoked := casav1.ErrorCode_ERROR_CODE_SESSION_REVOKED
	const other = "https://evil.example"

	// wantUntouched checks that the cookie's session was neither revoked nor
	// touched since lastUsed, and that no command was recorded.
	wantUntouched := func(t *testing.T, h *identityHarness, cookie string, lastUsed time.Time) {
		t.Helper()
		hash := sha256.Sum256(decodeToken(t, cookie))
		if n := h.count(t, "SELECT count(*) FROM session WHERE token_hash = $1 AND revoked_at IS NULL AND last_used_at = $2", hash[:], lastUsed); n != 1 {
			t.Fatal("the session was revoked or touched; want no effect")
		}
		if n := h.count(t, "SELECT count(*) FROM idempotency_record"); n != 0 {
			t.Fatalf("%d idempotency records, want none", n)
		}
	}

	t.Run("a same-origin request with the session cookie is authenticated", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		cookie := h.webSignIn(t, email, password, personal)
		wantOK(t, h.read(sameOrigin(cookie), web, true))
		// Sec-Fetch-Site is optional: a browser without it is judged by Origin.
		wantOK(t, h.read(cookieCall(cookie, testOrigin, ""), web, true))
		wantOK(t, h.webSignOut(sameOrigin(cookie)))
		if n := h.count(t, "SELECT count(*) FROM session WHERE client_kind = 'web' AND revoke_reason = 'user'"); n != 1 {
			t.Fatalf("%d web sessions signed out, want 1", n)
		}
	})

	t.Run("SEC-162 (a) a cookie request from another origin is rejected without effect", func(t *testing.T) {
		// Focus 2 of plan 1.5: staging on a subdomain is same-site, so only
		// the origin check holds it (SEC-41 (d)).
		h := startIdentity(t)
		h.signUp(t, email, password)
		signedIn := h.clock.Now()
		cookie := h.webSignIn(t, email, password, personal)
		// Past touchAfter, so an authenticated request would move last_used_at.
		h.clock.advance(2 * time.Minute)
		for _, c := range []struct{ origin, fetchSite string }{
			{"https://staging.casa.example", "same-origin"},
			{"null", "same-origin"},
			{"https://casa.example:8443", "same-origin"},
			{"http://casa.example", "same-origin"},
			{"https://casa.example/", "same-origin"},
			{"https://CASA.example", "same-origin"},
			{"https://casa.example.evil.example", "same-origin"},
			{"", "same-origin"},
			{"", ""},
			{testOrigin, "same-site"},
			{testOrigin, "cross-site"},
			{testOrigin, "none"},
			{testOrigin, "Same-Origin"},
		} {
			err := h.webSignOut(cookieCall(cookie, c.origin, c.fetchSite))
			if connect.CodeOf(err) != connect.CodePermissionDenied {
				t.Errorf("Origin %q, Sec-Fetch-Site %q: %v; want permission denied", c.origin, c.fetchSite, err)
				continue
			}
			wantCrossOriginRejected(t, err)
		}
		wantUntouched(t, h, cookie, signedIn)
	})

	t.Run("a public web RPC from another origin is rejected", func(t *testing.T) {
		// SEC-162 (a) note 1.5 (2): closes login CSRF; also a same-origin
		// web RPC without Origin.
		h := startIdentity(t)
		h.signUp(t, email, password)
		for _, c := range []struct{ origin, fetchSite string }{
			{other, ""}, {"", ""}, {"", "same-origin"}, {testOrigin, "cross-site"},
		} {
			ctx := cookieCall("", c.origin, c.fetchSite)
			info, _ := connect.CallInfoForClientContext(ctx)
			_, err := h.client.SignIn(ctx, &casav1.SignInRequest{
				RequestContext: identityRC(web), Email: email, Password: password, DeviceMode: personal,
			})
			wantCrossOriginRejected(t, err)
			if n := len(info.ResponseHeader().Values("Set-Cookie")); n != 0 {
				t.Fatalf("Origin %q: %d Set-Cookie headers, want none", c.origin, n)
			}
		}
		if n := h.count(t, "SELECT count(*) FROM session"); n != 1 {
			t.Fatalf("%d sessions, want only the sign-up's", n)
		}
	})

	t.Run("a request claiming a native client with a foreign Origin is rejected", func(t *testing.T) {
		// SEC-162 (a) note 1.5 (3): client_kind is the sender's claim.
		h := startIdentity(t)
		h.signUp(t, email, password)
		for _, c := range []struct{ origin, fetchSite string }{{other, ""}, {other, "cross-site"}, {"", "cross-site"}} {
			_, err := h.client.SignIn(cookieCall("", c.origin, c.fetchSite), &casav1.SignInRequest{
				RequestContext: identityRC(ios), Email: email, Password: password,
			})
			wantCrossOriginRejected(t, err)
		}
		if n := h.count(t, "SELECT count(*) FROM session"); n != 1 {
			t.Fatalf("%d sessions, want only the sign-up's", n)
		}
	})

	t.Run("a native request without Origin is accepted", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		tokens, err := h.signIn(email, password)
		wantOK(t, err)
		wantOK(t, h.authenticated(tokens.GetAccessToken()))
	})

	t.Run("SEC-162 (b) a form POST with the cookie has no effect", func(t *testing.T) {
		// A cross-site form can send only these types, none a Connect type.
		h := startIdentity(t)
		h.signUp(t, email, password)
		signedIn := h.clock.Now()
		cookie := h.webSignIn(t, email, password, personal)
		h.clock.advance(2 * time.Minute)
		for _, contentType := range []string{"text/plain", "application/x-www-form-urlencoded", "multipart/form-data; boundary=x"} {
			for _, origin := range []string{other, testOrigin} {
				req, err := http.NewRequestWithContext(t.Context(), http.MethodPost,
					h.url+casav1connect.IdentityServiceSignOutProcedure, strings.NewReader("x"))
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Set("Content-Type", contentType)
				req.Header.Set("Origin", origin)
				req.Header.Set("Cookie", sessionCookieName+"="+cookie)
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatal(err)
				}
				_ = res.Body.Close()
				if res.StatusCode != http.StatusUnsupportedMediaType {
					t.Errorf("%s from %s: status %d, want 415 (not a Connect call)", contentType, origin, res.StatusCode)
				}
			}
		}
		// No Connect GET for a cookie session: every procedure has side
		// effects, so Connect refuses the method before the boundary.
		req, err := http.NewRequestWithContext(t.Context(), http.MethodGet,
			h.url+casav1connect.IdentityServiceSignOutProcedure+"?encoding=json&message={}", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Origin", testOrigin)
		req.Header.Set("Cookie", sessionCookieName+"="+cookie)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		_ = res.Body.Close()
		if res.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("GET: status %d, want 405", res.StatusCode)
		}
		wantUntouched(t, h, cookie, signedIn)
	})

	t.Run("a cookie and a bearer token together still need the origin", func(t *testing.T) {
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		cookie := h.webSignIn(t, email, password, personal)
		both := func(origin string) context.Context {
			ctx := cookieCall(cookie, origin, "")
			info, _ := connect.CallInfoForClientContext(ctx)
			info.RequestHeader().Set("Authorization", "Bearer "+tokens.GetAccessToken())
			return ctx
		}
		wantCrossOriginRejected(t, h.read(both(""), ios, true))
		wantCrossOriginRejected(t, h.read(both(other), ios, true))
		// From the app's origin the bearer token decides (decision 11 step 3).
		wantOK(t, h.read(both(testOrigin), ios, true))
	})

	t.Run("the client kind must match the session", func(t *testing.T) {
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		cookie := h.webSignIn(t, email, password, personal)
		mismatch := func(t *testing.T, err error) {
			t.Helper()
			wantCode(t, err, connect.CodeInvalidArgument)
			var connectErr *connect.Error
			if !errors.As(err, &connectErr) || connectErr.Message() != "client kind does not match the session" {
				t.Fatalf("error %v; want \"client kind does not match the session\"", err)
			}
		}
		mismatch(t, h.read(sameOrigin(cookie), ios, true))
		withBearer := cookieCall("", testOrigin, "same-origin")
		info, _ := connect.CallInfoForClientContext(withBearer)
		info.RequestHeader().Set("Authorization", "Bearer "+tokens.GetAccessToken())
		mismatch(t, h.read(withBearer, web, true))
	})

	t.Run("a native access token in the cookie is not a session", func(t *testing.T) {
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		for _, value := range []string{tokens.GetAccessToken(), tokens.GetRefreshToken(), strings.Repeat("A", 43), "x"} {
			err := h.read(sameOrigin(value), web, true)
			wantCode(t, err, connect.CodeUnauthenticated)
			if got := details(t, err); got != nil {
				t.Fatalf("details %v, want none (sign in again)", got)
			}
		}
	})

	t.Run("an expired web session says session expired and a revoked one session revoked", func(t *testing.T) {
		// SEC-161 (d): personal mode ends after 30 idle days or 180 days.
		h := startIdentity(t)
		h.signUp(t, email, password)
		idle := h.webSignIn(t, email, password, personal)
		h.clock.advance(30 * day)
		wantErrorCode(t, h.read(sameOrigin(idle), web, true), connect.CodeUnauthenticated, expired)

		active := h.webSignIn(t, email, password, personal)
		for range 180 / 29 {
			h.clock.advance(29 * day)
			wantOK(t, h.read(sameOrigin(active), web, false))
		}
		h.clock.advance(6 * day) // day 180
		wantErrorCode(t, h.read(sameOrigin(active), web, true), connect.CodeUnauthenticated, expired)

		signedOut := h.webSignIn(t, email, password, personal)
		wantOK(t, h.webSignOut(sameOrigin(signedOut)))
		wantErrorCode(t, h.read(sameOrigin(signedOut), web, true), connect.CodeUnauthenticated, revoked)
	})

	t.Run("SEC-165 (g) only user-initiated requests keep a shared session alive", func(t *testing.T) {
		// Focus 5 of plan 1.5: background requests never move the idle
		// limit of a shared session, and nothing moves the absolute one.
		h := startIdentity(t)
		h.signUp(t, email, password)
		at := func(t *testing.T, start time.Time, minute int) {
			t.Helper()
			h.clock.advance(start.Add(time.Duration(minute) * time.Minute).Sub(h.clock.Now()))
		}

		start := h.clock.Now()
		background := h.webSignIn(t, email, password, shared)
		for minute := 5; minute < 30; minute += 5 {
			at(t, start, minute)
			wantOK(t, h.read(sameOrigin(background), web, false))
		}
		at(t, start, 30)
		wantErrorCode(t, h.read(sameOrigin(background), web, false), connect.CodeUnauthenticated, expired)

		start = h.clock.Now()
		kept := h.webSignIn(t, email, password, shared)
		for minute := 5; minute < 29; minute += 5 {
			at(t, start, minute)
			wantOK(t, h.read(sameOrigin(kept), web, false))
		}
		at(t, start, 29)
		wantOK(t, h.read(sameOrigin(kept), web, true)) // "Tetap masuk"
		for minute := 30; minute <= 58; minute += 4 {
			at(t, start, minute)
			wantOK(t, h.read(sameOrigin(kept), web, false))
		}
		at(t, start, 59)
		wantErrorCode(t, h.read(sameOrigin(kept), web, true), connect.CodeUnauthenticated, expired)

		start = h.clock.Now()
		used := h.webSignIn(t, email, password, shared)
		for minute := 10; minute < 12*60; minute += 10 {
			at(t, start, minute)
			wantOK(t, h.read(sameOrigin(used), web, true))
		}
		at(t, start, 12*60)
		wantErrorCode(t, h.read(sameOrigin(used), web, true), connect.CodeUnauthenticated, expired)
	})
}
