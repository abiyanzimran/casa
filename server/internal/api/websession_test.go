package api_test

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"slices"
	"strings"
	"sync"
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
			RequestContext: identityRC(ios), VerificationToken: h.lastLink(t), Email: email, Password: password,
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
						RequestContext: identityRC(c.kind), VerificationToken: token, Email: email, Password: password, DeviceMode: c.mode,
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
			RequestContext: identityRC(web), VerificationToken: h.lastLink(t), Email: email, Password: password, DeviceMode: personal,
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
			RequestContext: identityRC(web), VerificationToken: h.lastLink(t), Email: email, Password: password, DeviceMode: personal,
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
			RequestContext: identityRC(ios), VerificationToken: h.lastLink(t), Email: email, Password: password,
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

	// wantSiteCleared checks a web sign-out's answer (SEC-164 (b); plan 1.5
	// decision 14): Clear-Site-Data exactly "cookies", "storage" (no
	// "cache"), and one Set-Cookie that expires the session cookie with the
	// attributes of decision 7, never a new value.
	wantSiteCleared := func(t *testing.T, info *connect.CallInfo) {
		t.Helper()
		if got := info.ResponseHeader().Values("Clear-Site-Data"); len(got) != 1 || got[0] != `"cookies", "storage"` {
			t.Fatalf("Clear-Site-Data %q, want exactly %q", got, `"cookies", "storage"`)
		}
		cookie, line := sessionCookie(t, info)
		if cookie.Name != sessionCookieName || cookie.Value != "" || cookie.MaxAge != -1 || !strings.Contains(line, "; Max-Age=0") ||
			!cookie.Secure || !cookie.HttpOnly || cookie.SameSite != http.SameSiteStrictMode || cookie.Path != "/" ||
			strings.Contains(strings.ToLower(line), "domain") {
			t.Fatalf("Set-Cookie %q: want %s= expired (Max-Age=0), Secure, HttpOnly, SameSite=Strict, Path=/ and no Domain",
				cookieShape(line), sessionCookieName)
		}
	}

	t.Run("a web sign-out clears the cookie and the site data", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		for _, mode := range []casav1.DeviceMode{personal, shared} {
			cookie := h.webSignIn(t, email, password, mode)
			ctx := sameOrigin(cookie)
			info, _ := connect.CallInfoForClientContext(ctx)
			wantOK(t, h.webSignOut(ctx))
			wantSiteCleared(t, info)
			wantErrorCode(t, h.read(sameOrigin(cookie), web, true), connect.CodeUnauthenticated, revoked)
		}

		// SignOut as the first request after the 15-minute mark: the
		// boundary rotates the value, and the answer must expire the cookie
		// instead of handing out a new value for a revoked session
		// (SEC-161 (c); plan 1.5 decision 14).
		cookie := h.webSignIn(t, email, password, personal)
		h.clock.advance(15 * time.Minute)
		ctx := sameOrigin(cookie)
		info, _ := connect.CallInfoForClientContext(ctx)
		wantOK(t, h.webSignOut(ctx))
		wantSiteCleared(t, info)
		wantErrorCode(t, h.read(sameOrigin(cookie), web, true), connect.CodeUnauthenticated, revoked)
		if n := h.count(t, "SELECT count(*) FROM session WHERE client_kind = 'web' AND revoked_at IS NULL"); n != 0 {
			t.Fatalf("%d web sessions still live, want none", n)
		}
	})

	t.Run("a native sign-out sends no Clear-Site-Data", func(t *testing.T) {
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		ctx := bearer(tokens.GetAccessToken())
		info, _ := connect.CallInfoForClientContext(ctx)
		_, err := h.client.SignOut(ctx, &casav1.SignOutRequest{RequestContext: commandRC()})
		wantOK(t, err)
		if got := info.ResponseHeader().Values("Clear-Site-Data"); len(got) != 0 {
			t.Fatalf("Clear-Site-Data %q on a native sign-out, want none", got)
		}
		if got := info.ResponseHeader().Values("Set-Cookie"); len(got) != 0 {
			t.Fatalf("%d Set-Cookie headers on a native sign-out, want none", len(got))
		}
	})

	t.Run("SEC-162 (a) a cookie request from another origin is rejected without effect", func(t *testing.T) {
		// Focus 2 of plan 1.5: staging on a subdomain is same-site, so only
		// the origin check holds it (SEC-41 (d)). This proves SEC-162 (a)
		// note 1.5 clause (1), a request with the cookie. Clause (2), a web
		// RPC without Origin, is the subtest "a public web RPC from another
		// origin is rejected"; clause (3), a claimed native client, is "a
		// request claiming a native client with a foreign Origin is rejected".
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
		signedIn := h.clock.Now()
		cookie := h.webSignIn(t, email, password, personal)
		// Past touchAfter, so an authenticated request would move last_used_at.
		h.clock.advance(2 * time.Minute)
		both := func(origin string) context.Context {
			ctx := cookieCall(cookie, origin, "")
			info, _ := connect.CallInfoForClientContext(ctx)
			info.RequestHeader().Set("Authorization", "Bearer "+tokens.GetAccessToken())
			return ctx
		}
		wantCrossOriginRejected(t, h.read(both(""), ios, true))
		wantCrossOriginRejected(t, h.read(both(other), ios, true))
		wantUntouched(t, h, cookie, signedIn)
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
		// This proves the server half of SEC-165 (g); the client half (idle
		// counted from input, the warning, signing itself out) is 1.10's.
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

		// Lead ruling on Task 4 review: every user action moves the idle
		// limit, so the client's 60-second warning, counted from the last
		// input, never comes after the server has ended the session.
		start = h.clock.Now()
		early := h.webSignIn(t, email, password, shared)
		at(t, start, 1)
		wantOK(t, h.read(sameOrigin(early), web, true))
		h.clock.advance(29*time.Minute + 30*time.Second) // minute 30:30
		wantOK(t, h.read(sameOrigin(early), web, false))
		at(t, start, 31) // the action at minute 1 + 30 minutes
		wantErrorCode(t, h.read(sameOrigin(early), web, false), connect.CodeUnauthenticated, expired)

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

// newCookie returns the session cookie a response set, checked for the
// attributes every session cookie has, or nil for none; more than one fails
// the test (a response sets at most one session cookie).
func newCookie(t *testing.T, info *connect.CallInfo) *http.Cookie {
	t.Helper()
	if len(info.ResponseHeader().Values("Set-Cookie")) == 0 {
		return nil
	}
	cookie, line := sessionCookie(t, info)
	wantHostCookie(t, cookie, line)
	return cookie
}

// webRead calls the probe Read from the app with cookie as a user's action
// and returns the session cookie the response set (nil for none), which a
// browser would store.
func (h *identityHarness) webRead(t *testing.T, cookie string) (*http.Cookie, error) {
	t.Helper()
	ctx := sameOrigin(cookie)
	info, _ := connect.CallInfoForClientContext(ctx)
	err := h.read(ctx, web, true)
	return newCookie(t, info), err
}

func TestCookieRotation(t *testing.T) {
	// SEC-161 (c) and its note "FS-1 tugas 1.5" rules (1) to (5); plan 1.5
	// decision 12. The clock is the test's: no sleep.
	password, next := strings.Repeat("cookie rotation ", 2), strings.Repeat("rotated again ", 2) // built, so no secret scanner sees a literal
	const email = "owner@example.test"
	const day = 24 * time.Hour
	revoked := casav1.ErrorCode_ERROR_CODE_SESSION_REVOKED

	// served reads with cookie and wants it served without a new cookie.
	served := func(t *testing.T, h *identityHarness, cookie string) {
		t.Helper()
		got, err := h.webRead(t, cookie)
		wantOK(t, err)
		if got != nil {
			t.Fatal("the response set a new cookie; want none")
		}
	}
	// rotated reads with cookie and wants it served with a new value, which
	// it returns.
	rotated := func(t *testing.T, h *identityHarness, cookie string) *http.Cookie {
		t.Helper()
		got, err := h.webRead(t, cookie)
		wantOK(t, err)
		if got == nil || got.Value == cookie {
			t.Fatal("no new cookie value; want one")
		}
		return got
	}
	// revocations counts the sessions revoked for reuse and their rows in
	// the security trail.
	revocations := func(t *testing.T, h *identityHarness) (sessions, events int) {
		t.Helper()
		return h.count(t, "SELECT count(*) FROM session WHERE revoke_reason = 'reuse_detected'"),
			h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'session_revoked' AND detail = 'reuse_detected'")
	}
	// wantRevokedNone checks that no session is revoked and that the
	// security trail holds only the sign_in row of every session opened and
	// the reauths re-authentications the test made: anything more is a row
	// the rotation wrote (SEC-161 (c) note 1.5 rules (2), (7): "tanpa baris
	// jejak keamanan", "tanpa pencabutan").
	wantRevokedNone := func(t *testing.T, h *identityHarness, reauths int) {
		t.Helper()
		if n := h.count(t, "SELECT count(*) FROM session WHERE revoked_at IS NOT NULL"); n != 0 {
			t.Fatalf("%d sessions revoked, want none", n)
		}
		if events, sessions := h.count(t, "SELECT count(*) FROM security_event"), h.count(t, "SELECT count(*) FROM session"); events != sessions+reauths {
			t.Fatalf("%d security_event rows for %d sessions opened and %d re-authentications; want only theirs", events, sessions, reauths)
		}
	}
	wantNoRevocation := func(t *testing.T, h *identityHarness) { t.Helper(); wantRevokedNone(t, h, 0) }
	// reauth re-authenticates for a password change in ctx and returns the
	// call's info, which holds the response headers.
	reauth := func(h *identityHarness, ctx context.Context, kind casav1.ClientKind, password string) (*connect.CallInfo, error) {
		info, _ := connect.CallInfoForClientContext(ctx)
		_, err := h.client.Reauthenticate(ctx, &casav1.ReauthenticateRequest{
			RequestContext: identityRC(kind), Password: password, Purpose: casav1.ReauthPurpose_REAUTH_PURPOSE_CHANGE_PASSWORD,
		})
		return info, err
	}
	// changePassword is a web ChangePassword in ctx.
	changePassword := func(h *identityHarness, ctx context.Context) (*connect.CallInfo, error) {
		info, _ := connect.CallInfoForClientContext(ctx)
		rc := commandRC()
		rc.ClientKind = web
		_, err := h.client.ChangePassword(ctx, &casav1.ChangePasswordRequest{RequestContext: rc, NewPassword: next})
		return info, err
	}
	reauthRequired := casav1.ErrorCode_ERROR_CODE_REAUTH_REQUIRED
	wantReuseRevoked := func(t *testing.T, h *identityHarness) {
		t.Helper()
		if sessions, events := revocations(t, h); sessions != 1 || events != 1 {
			t.Fatalf("%d sessions revoked for reuse with %d security_event rows; want 1 and 1", sessions, events)
		}
	}
	// lostRotation signs in, rotates the value at 15 minutes and drops the
	// new cookie, as a response the browser never stored; it returns the
	// value the browser still holds, the dropped one, and the rotation time.
	lostRotation := func(t *testing.T, h *identityHarness) (held, dropped string, at time.Time) {
		t.Helper()
		h.signUp(t, email, password)
		held = h.webSignIn(t, email, password, personal)
		h.clock.advance(15 * time.Minute)
		return held, rotated(t, h, held).Value, h.clock.Now()
	}

	t.Run("the cookie value changes after 15 minutes", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		signedIn := h.clock.Now()
		a := h.webSignIn(t, email, password, personal)
		h.clock.advance(15*time.Minute - time.Second)
		served(t, h, a) // 14:59
		h.clock.advance(time.Second)
		b := rotated(t, h, a) // 15:00
		// Plan 1.5 decision 7: Max-Age is the remaining absolute lifetime.
		if want := int((180*day - 15*time.Minute).Seconds()); b.MaxAge != want {
			t.Fatalf("Max-Age %d, want %d", b.MaxAge, want)
		}
		served(t, h, b.Value)
		if n := h.count(t, "SELECT count(*) FROM session WHERE token_issued_at = $1", signedIn.Add(15*time.Minute)); n != 1 {
			t.Fatal("the rotation did not record when the new value was issued")
		}

		// A shared session's new value is a session cookie too (SEC-165 (a)).
		shared := h.webSignIn(t, email, password, shared)
		h.clock.advance(15 * time.Minute)
		sharedNext, err := h.webRead(t, shared)
		wantOK(t, err)
		if sharedNext == nil || sharedNext.Value == shared || sharedNext.MaxAge != 0 || sharedNext.RawExpires != "" {
			t.Fatal("shared rotation: want a new value in a cookie without Max-Age or Expires")
		}
		wantNoRevocation(t, h)
	})

	t.Run("the previous value serves for 60 seconds", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		a := h.webSignIn(t, email, password, personal)
		h.clock.advance(15 * time.Minute)
		b := rotated(t, h, a)
		served(t, h, b.Value) // the browser stored it: confirmed
		h.clock.advance(59 * time.Second)
		served(t, h, a) // a request that left before the rotation
		served(t, h, b.Value)
		wantNoRevocation(t, h)
	})

	t.Run("a lost new cookie does not revoke the session", func(t *testing.T) {
		// Focus 1 of plan 1.5; SEC-161 (c) note 1.5 rule (2): the dropped
		// value was never confirmed, so the server issues again.
		h := startIdentity(t)
		a, b, at := lostRotation(t, h)
		h.clock.advance(at.Add(61 * time.Second).Sub(h.clock.Now()))
		c := rotated(t, h, a)
		if c.Value == b {
			t.Fatal("the reissue sent the dropped value again; want a new one")
		}
		if want := int((180*day - 15*time.Minute - 61*time.Second).Seconds()); c.MaxAge != want {
			t.Fatalf("Max-Age %d, want %d", c.MaxAge, want)
		}
		wantNoRevocation(t, h)
		served(t, h, a) // the presented value got a new 60-second window
		served(t, h, c.Value)
		wantNoRevocation(t, h)
	})

	t.Run("a previous value used after the new one revokes the session", func(t *testing.T) {
		// SEC-161 (c): "dipakai sesudah itu, seluruh sesi dicabut dan satu
		// baris ditulis di jejak keamanan pemiliknya"; note 1.5 rule (2)
		// when the current value is confirmed.
		h := startIdentity(t)
		h.signUp(t, email, password)
		a := h.webSignIn(t, email, password, personal)
		h.clock.advance(15 * time.Minute)
		at := h.clock.Now()
		b := rotated(t, h, a)
		h.clock.advance(time.Second)
		served(t, h, b.Value) // used once: confirmed (rule (1))
		h.clock.advance(at.Add(61 * time.Second).Sub(h.clock.Now()))
		_, err := h.webRead(t, a)
		wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
		wantReuseRevoked(t, h)
		_, err = h.webRead(t, b.Value)
		wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
		wantReuseRevoked(t, h)
	})

	t.Run("a discarded value revokes the session", func(t *testing.T) {
		// SEC-161 (c) note 1.5 rule (3): only a second holder can bring the
		// value the browser never received.
		h := startIdentity(t)
		a, b, at := lostRotation(t, h)
		h.clock.advance(at.Add(61 * time.Second).Sub(h.clock.Now()))
		c := rotated(t, h, a)
		wantNoRevocation(t, h)
		_, err := h.webRead(t, b)
		wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
		wantReuseRevoked(t, h)
		_, err = h.webRead(t, c.Value)
		wantErrorCode(t, err, connect.CodeUnauthenticated, revoked)
	})

	t.Run("an unknown value is no session, never a revocation", func(t *testing.T) {
		// SEC-161 (c) note 1.5 rule (5): the server cannot tell whose it is,
		// so the browser keeps its local data.
		h := startIdentity(t)
		h.signUp(t, email, password)
		a := h.webSignIn(t, email, password, personal)
		h.clock.advance(15 * time.Minute)
		b := rotated(t, h, a)
		served(t, h, b.Value)
		h.clock.advance(15 * time.Minute)
		c := rotated(t, h, b.Value)
		served(t, h, c.Value)
		h.clock.advance(2 * time.Minute)
		// a is two rotations old: no longer current, previous or discarded.
		for _, value := range []string{tokenOf(t), a} {
			got, err := h.webRead(t, value)
			wantCode(t, err, connect.CodeUnauthenticated)
			if d := details(t, err); d != nil || got != nil {
				t.Fatalf("details %v, new cookie %t; want neither (sign in again)", d, got != nil)
			}
		}
		wantNoRevocation(t, h)
		served(t, h, c.Value)
	})

	t.Run("two requests at the rotation point rotate once", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password)
		a := h.webSignIn(t, email, password, personal)
		h.clock.advance(15 * time.Minute)
		cookies := h.concurrentReads(t, a, a)
		if len(cookies) != 1 {
			t.Fatalf("%d new cookies, want exactly 1", len(cookies))
		}
		h.clock.advance(59 * time.Second)
		served(t, h, a) // the other request's value keeps its window
		served(t, h, cookies[0])
		wantNoRevocation(t, h)
	})

	t.Run("two requests with the old value after the window reissue once", func(t *testing.T) {
		// SEC-161 (c) note 1.5 rule (4): one decision per row, so the second
		// is the previous value inside the new window, and no response sets
		// a value the first already discarded.
		h := startIdentity(t)
		a, _, at := lostRotation(t, h)
		h.clock.advance(at.Add(61 * time.Second).Sub(h.clock.Now()))
		cookies := h.concurrentReads(t, a, a)
		if len(cookies) != 1 {
			t.Fatalf("%d new cookies, want exactly 1", len(cookies))
		}
		wantNoRevocation(t, h)
		served(t, h, cookies[0])
		wantNoRevocation(t, h)
	})

	t.Run("a rotation on a failing request still delivers the new cookie", func(t *testing.T) {
		// Plan 1.5 decision 11 step 3: the boundary sets it before the
		// handler runs.
		h := startIdentity(t)
		h.signUp(t, email, password)
		a := h.webSignIn(t, email, password, personal)
		h.clock.advance(15 * time.Minute)
		info, err := changePassword(h, sameOrigin(a))
		wantErrorCode(t, err, connect.CodePermissionDenied, reauthRequired)
		b := newCookie(t, info)
		if b == nil || b.Value == a {
			t.Fatal("the failing request set no new cookie; want one")
		}
		served(t, h, b.Value)
	})

	t.Run("a web re-authentication rotates the cookie", func(t *testing.T) {
		// SEC-161 (c): "selalu saat masuk, autentikasi ulang, dan
		// pemulihan" (closes session fixation); sign-in and recovery open a
		// new session, so this is the re-authentication. Native sessions
		// never get a cookie (SEC-161 (b)).
		h := startIdentity(t)
		tokens := h.signUp(t, email, password)
		a := h.webSignIn(t, email, password, personal)
		info, err := reauth(h, sameOrigin(a), web, next) // wrong password: no rotation
		wantErrorCode(t, err, connect.CodeUnauthenticated, casav1.ErrorCode_ERROR_CODE_CREDENTIALS_REJECTED)
		if newCookie(t, info) != nil {
			t.Fatal("a failed re-authentication set a cookie; want none")
		}
		info, err = reauth(h, sameOrigin(a), web, password)
		wantOK(t, err)
		b := newCookie(t, info)
		if b == nil || b.Value == a || b.MaxAge != int((180*day).Seconds()) {
			t.Fatal("a web re-authentication set no new cookie, or one without the remaining lifetime; want one")
		}
		served(t, h, b.Value)
		served(t, h, a) // inside its 60-second window
		info, err = reauth(h, bearer(tokens.GetAccessToken()), ios, password)
		wantOK(t, err)
		if n := len(info.ResponseHeader().Values("Set-Cookie")); n != 0 {
			t.Fatalf("a native re-authentication set %d cookies, want none", n)
		}
		// At the rotation point the boundary rotates and the handler rotates
		// again: the response carries only the last value.
		h.clock.advance(15 * time.Minute)
		info, err = reauth(h, sameOrigin(b.Value), web, password)
		wantOK(t, err)
		c := newCookie(t, info)
		if c == nil || c.Value == b.Value {
			t.Fatal("a web re-authentication at the rotation point set no new cookie; want one")
		}
		served(t, h, c.Value)
		// The value the browser held (b, confirmed) stays the previous one:
		// a request with it already on its way is still served.
		h.clock.advance(59 * time.Second)
		served(t, h, b.Value)

		_, err = changePassword(h, sameOrigin(c.Value))
		wantOK(t, err) // the grant of the value issued with it (SEC-161 (c) note 1.5 rule (7))
	})

	t.Run("a forced rotation keeps the value the browser holds and never discards the dropped one", func(t *testing.T) {
		// SEC-161 (c) note 1.5 rule (6): a re-authentication carried by the
		// previous value while the current one is unconfirmed keeps the
		// previous value, with a new window; the unconfirmed value it drops
		// may have reached the browser, so it is not discarded (3) and is
		// answered by (5), never revoked.
		h := startIdentity(t)
		a, b, at := lostRotation(t, h)
		h.clock.advance(10 * time.Second)
		info, err := reauth(h, sameOrigin(a), web, password)
		wantOK(t, err)
		c := newCookie(t, info)
		if c == nil || c.Value == a || c.Value == b {
			t.Fatal("the re-authentication set no new value; want one")
		}
		h.clock.advance(at.Add(69 * time.Second).Sub(h.clock.Now())) // a's first window ended at 60 s
		served(t, h, a)
		got, err := h.webRead(t, b)
		wantCode(t, err, connect.CodeUnauthenticated)
		if d := details(t, err); d != nil || got != nil {
			t.Fatalf("details %v, new cookie %t; want neither (rule (5))", d, got != nil)
		}
		wantRevokedNone(t, h, 1)
		served(t, h, c.Value)
	})

	t.Run("a web re-authentication grant serves only the value issued with it", func(t *testing.T) {
		// SEC-161 (c) note 1.5 rule (7): the value held before the
		// re-authentication, inside its window, gets no grant, without a
		// revocation; the value issued with the grant uses it.
		h := startIdentity(t)
		h.signUp(t, email, password)
		a := h.webSignIn(t, email, password, personal)
		info, err := reauth(h, sameOrigin(a), web, password)
		wantOK(t, err)
		b := newCookie(t, info)
		if b == nil {
			t.Fatal("the re-authentication set no new value; want one")
		}
		_, err = changePassword(h, sameOrigin(a))
		wantErrorCode(t, err, connect.CodePermissionDenied, reauthRequired)
		wantRevokedNone(t, h, 1)
		_, err = changePassword(h, sameOrigin(b.Value))
		wantOK(t, err)
	})

	t.Run("a value reissued after a lost re-authentication response has no grant", func(t *testing.T) {
		// SEC-161 (c) note 1.5 rule (7): a value issued later by any
		// rotation, here (2), is not the one the grant was issued with.
		h := startIdentity(t)
		h.signUp(t, email, password)
		a := h.webSignIn(t, email, password, personal)
		info, err := reauth(h, sameOrigin(a), web, password)
		wantOK(t, err)
		if newCookie(t, info) == nil { // dropped: never stored
			t.Fatal("the re-authentication set no new value; want one")
		}
		h.clock.advance(61 * time.Second)
		info, err = changePassword(h, sameOrigin(a))
		wantErrorCode(t, err, connect.CodePermissionDenied, reauthRequired)
		c := newCookie(t, info)
		if c == nil || c.Value == a {
			t.Fatal("no reissue; want a new value")
		}
		_, err = changePassword(h, sameOrigin(c.Value))
		wantErrorCode(t, err, connect.CodePermissionDenied, reauthRequired)
		wantRevokedNone(t, h, 1)
		if n := h.count(t, "SELECT count(*) FROM reauth_grant WHERE used_at IS NULL"); n != 1 {
			t.Fatalf("%d open grants, want the one, unused", n)
		}
	})

	t.Run("a personal cookie in its last second is sent expired, never as a session cookie", func(t *testing.T) {
		// Plan 1.5 decision 7: Max-Age is the remaining lifetime rounded
		// down; under a second it would round to none, which makes a
		// session cookie, so it is sent as already expired.
		h := startIdentity(t)
		h.signUp(t, email, password)
		signedIn := h.clock.Now()
		cookie := h.webSignIn(t, email, password, personal)
		for range 180 / 29 {
			h.clock.advance(29 * day) // inside the 30-day idle limit
			cookie = rotated(t, h, cookie).Value
		}
		h.clock.advance(signedIn.Add(180*day - 500*time.Millisecond).Sub(h.clock.Now()))
		ctx := sameOrigin(cookie)
		info, _ := connect.CallInfoForClientContext(ctx)
		wantOK(t, h.read(ctx, web, true))
		last := newCookie(t, info)
		line := info.ResponseHeader().Get("Set-Cookie")
		if last == nil || last.MaxAge >= 0 || !strings.Contains(line, "; Max-Age=0") {
			t.Fatalf("Set-Cookie %q; want Max-Age=0", cookieShape(line))
		}
	})
}

// concurrentReads sends one web read per cookie value and returns the new
// cookie values the responses set; every read must succeed. The reads are
// made to overlap: a transaction of the test locks the values' session row
// until every read waits on that lock, so all of them decide on the row at
// once (SEC-161 (c) note 1.5 rule (4)).
func (h *identityHarness) concurrentReads(t *testing.T, cookies ...string) []string {
	t.Helper()
	ctx := t.Context()
	lock, err := h.admin.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = lock.Rollback(ctx) }()
	hash := sha256.Sum256(decodeToken(t, cookies[0]))
	if _, err := lock.Exec(ctx, "SELECT 1 FROM session WHERE token_hash = $1 OR prev_token_hash = $1 FOR UPDATE", hash[:]); err != nil {
		t.Fatal(err)
	}
	infos := make([]*connect.CallInfo, len(cookies))
	errs := make([]error, len(cookies))
	var wg sync.WaitGroup
	for i, cookie := range cookies {
		call := sameOrigin(cookie)
		infos[i], _ = connect.CallInfoForClientContext(call)
		wg.Go(func() { errs[i] = h.read(call, web, true) })
	}
	for waiting := 0; waiting < len(cookies); { // no sleep: each query takes its own time
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if err := h.admin.QueryRow(ctx, `
			SELECT count(*) FROM pg_stat_activity
			 WHERE datname = current_database() AND wait_event_type = 'Lock'`).Scan(&waiting); err != nil {
			t.Fatal(err)
		}
	}
	if err := lock.Rollback(ctx); err != nil {
		t.Fatal(err)
	}
	wg.Wait()
	var values []string
	for i, info := range infos {
		wantOK(t, errs[i])
		if cookie := newCookie(t, info); cookie != nil {
			values = append(values, cookie.Value)
		}
	}
	return values
}

// tokenOf returns a well-formed random token that no session holds.
func tokenOf(t *testing.T) string {
	t.Helper()
	var raw [32]byte
	if _, err := rand.Read(raw[:]); err != nil {
		t.Fatal(err)
	}
	return base64.RawURLEncoding.EncodeToString(raw[:])
}

// TestWithoutSession covers without_session on CompleteSignUp and
// CompleteRecovery (ADR-030 item 11; SEC-171 (b) (2) note of 14 November
// 2026, Q2-1 to Q2-5).
func TestWithoutSession(t *testing.T) {
	password, next := strings.Repeat("no session ", 3), strings.Repeat("no session next ", 2) // built, so no secret scanner sees a literal
	const email = "owner@example.test"
	rejected := casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED

	completeSignUp := func(h *identityHarness, ctx context.Context, kind casav1.ClientKind, link string, mode casav1.DeviceMode, without bool) error {
		_, err := h.client.CompleteSignUp(ctx, &casav1.CompleteSignUpRequest{
			RequestContext: identityRC(kind), VerificationToken: link, Email: h.lastEmail(), Password: password, DeviceMode: mode, WithoutSession: without,
		})
		return err
	}
	completeRecovery := func(h *identityHarness, ctx context.Context, kind casav1.ClientKind, link string, mode casav1.DeviceMode, without bool) error {
		_, err := h.client.CompleteRecovery(ctx, &casav1.CompleteRecoveryRequest{
			RequestContext: identityRC(kind), RecoveryToken: link, NewPassword: next, DeviceMode: mode, WithoutSession: without,
		})
		return err
	}

	t.Run("SEC-171 (b) (2) a web sign-up completed without a session sets no cookie, returns no tokens and opens no session", func(t *testing.T) {
		h := startIdentity(t)
		wantOK(t, h.startSignUp(h.issue(t, email), email))
		link := h.lastLink(t)
		ctx, info := webCall(safariIPhone)
		res, err := h.client.CompleteSignUp(ctx, &casav1.CompleteSignUpRequest{
			RequestContext: identityRC(web), VerificationToken: link, Email: email, Password: password, WithoutSession: true,
		})
		wantOK(t, err)
		if n := len(info.ResponseHeader().Values("Set-Cookie")); n != 0 || res.GetTokens() != nil {
			t.Fatalf("%d Set-Cookie headers, tokens %t; want none", n, res.GetTokens() != nil)
		}
		if n := h.count(t, "SELECT count(*) FROM session"); n != 0 {
			t.Fatalf("%d sessions, want none", n)
		}
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'sign_in'"); n != 0 {
			t.Fatalf("%d sign_in events, want none", n)
		}
		if n := h.count(t, "SELECT count(*) FROM account WHERE email = $1", email); n != 1 {
			t.Fatalf("%d accounts, want 1", n)
		}
		if n := h.count(t, "SELECT count(*) FROM credential"); n != 1 {
			t.Fatalf("%d credentials, want 1", n)
		}
		ctx, _ = webCall(safariIPhone)
		wantErrorCode(t, completeSignUp(h, ctx, web, link, 0, true), connect.CodePermissionDenied, rejected)
		_, err = h.signIn(email, password)
		wantOK(t, err)
	})

	t.Run("SEC-171 (b) (2) a web recovery without a session revokes every session and records recovery without a session", func(t *testing.T) {
		h := startIdentity(t)
		h.signUp(t, email, password) // a native session
		cookie := h.webSignIn(t, email, password, personal)
		_, err := h.startRecovery(email)
		wantOK(t, err)
		ctx, info := webCall(safariIPhone)
		res, err := h.client.CompleteRecovery(ctx, &casav1.CompleteRecoveryRequest{
			RequestContext: identityRC(web), RecoveryToken: h.lastLink(t), NewPassword: next, WithoutSession: true,
		})
		wantOK(t, err)
		if n := len(info.ResponseHeader().Values("Set-Cookie")); n != 0 || res.GetTokens() != nil {
			t.Fatalf("%d Set-Cookie headers, tokens %t; want none", n, res.GetTokens() != nil)
		}
		if n := h.count(t, "SELECT count(*) FROM session WHERE revoked_at IS NULL"); n != 0 {
			t.Fatalf("%d live sessions, want none", n)
		}
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'recovery' AND session_id IS NULL AND detail = $1", "Safari · iOS"); n != 1 {
			t.Fatalf("%d recovery events without a session and with the browser summary, want 1", n)
		}
		if n := h.count(t, "SELECT count(*) FROM security_event WHERE kind = 'recovery'"); n != 1 {
			t.Fatalf("%d recovery events, want exactly 1", n)
		}
		wantErrorCode(t, h.read(sameOrigin(cookie), web, true), connect.CodeUnauthenticated, casav1.ErrorCode_ERROR_CODE_SESSION_REVOKED)
		_, err = h.signIn(email, next)
		wantOK(t, err)
	})

	t.Run("the flag from a native client, or with a device mode, is an invalid argument", func(t *testing.T) {
		h := startIdentity(t)
		token := strings.Repeat("A", 43)
		type call func(*identityHarness, context.Context, casav1.ClientKind, string, casav1.DeviceMode, bool) error
		for name, complete := range map[string]call{"CompleteSignUp": completeSignUp, "CompleteRecovery": completeRecovery} {
			for _, c := range []struct {
				kind    casav1.ClientKind
				without bool
				mode    casav1.DeviceMode
				invalid bool
			}{
				{ios, true, 0, true}, {web, true, personal, true}, {web, false, 0, true},
				{ios, false, 0, false}, {web, false, personal, false}, {web, true, 0, false},
			} {
				ctx, _ := webCall(safariIPhone)
				err := complete(h, ctx, c.kind, token, c.mode, c.without)
				if !c.invalid {
					wantErrorCode(t, err, connect.CodePermissionDenied, rejected) // valid shape: the unknown link is rejected
					continue
				}
				wantCode(t, err, connect.CodeInvalidArgument)
				if got := ruleIDs(t, err); !slices.Equal(got, []string{"device_mode.web"}) {
					t.Errorf("%s %+v: rules %v, want [device_mode.web]", name, c, got)
				}
			}
		}
	})

	t.Run("a rejected link gets the same TOKEN_REJECTED with and without the flag", func(t *testing.T) {
		// SEC-30: code, message and details are equal.
		h := startIdentity(t)
		token := strings.Repeat("B", 43)
		ctx, _ := webCall(safariIPhone)
		withSignUp := rejectionOf(t, completeSignUp(h, ctx, web, token, 0, true))
		withRecovery := rejectionOf(t, completeRecovery(h, ctx, web, token, 0, true))
		for name, without := range map[string]rejection{
			"CompleteSignUp":   rejectionOf(t, completeSignUp(h, ctx, web, token, personal, false)),
			"CompleteRecovery": rejectionOf(t, completeRecovery(h, ctx, web, token, personal, false)),
		} {
			with := withSignUp
			if name == "CompleteRecovery" {
				with = withRecovery
			}
			if !with.equal(without) || with.code != connect.CodePermissionDenied {
				t.Errorf("%s: rejection with the flag %+v differs from without %+v", name, with, without)
			}
		}
	})

	t.Run("the token-use limiter counts the same with and without the flag", func(t *testing.T) {
		// 10.1 §8.1: 20 failed token uses per address, the 21st is limited.
		h := startIdentity(t)
		token := strings.Repeat("C", 43)
		for i := range 20 {
			ctx, _ := webCall(safariIPhone)
			if i%2 == 0 {
				wantErrorCode(t, completeSignUp(h, ctx, web, token, 0, true), connect.CodePermissionDenied, rejected)
			} else {
				wantErrorCode(t, completeSignUp(h, ctx, web, token, personal, false), connect.CodePermissionDenied, rejected)
			}
		}
		for _, without := range []bool{true, false} {
			mode := personal
			if without {
				mode = 0
			}
			ctx, _ := webCall(safariIPhone)
			wantErrorCode(t, completeSignUp(h, ctx, web, token, mode, without), connect.CodeResourceExhausted, casav1.ErrorCode_ERROR_CODE_RATE_LIMITED)
		}
	})

	t.Run("a without-session completion leaves a stale cookie untouched", func(t *testing.T) {
		// Q2-4: a request carrying the cookie of a revoked session gets no
		// Set-Cookie at all.
		h := startIdentity(t)
		h.signUp(t, email, password)
		stale := h.webSignIn(t, email, password, personal)
		wantOK(t, h.webSignOut(sameOrigin(stale)))
		wantOK(t, h.startSignUp(h.issue(t, "second@example.test"), "second@example.test"))
		ctx := sameOrigin(stale)
		info, _ := connect.CallInfoForClientContext(ctx)
		wantOK(t, completeSignUp(h, ctx, web, h.lastLink(t), 0, true))
		if got := info.ResponseHeader().Values("Set-Cookie"); len(got) != 0 {
			t.Fatalf("%d Set-Cookie headers, want none", len(got))
		}
	})
}
