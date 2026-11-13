package api_test

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"slices"
	"strings"
	"testing"
	"time"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"connectrpc.com/connect/v2"
	"google.golang.org/protobuf/proto"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
)

const sessionCookieName = "__Host-casa_session"

var (
	web      = casav1.ClientKind_CLIENT_KIND_WEB
	personal = casav1.DeviceMode_DEVICE_MODE_PERSONAL
	shared   = casav1.DeviceMode_DEVICE_MODE_SHARED
)

// safariIPhone is the User-Agent of Safari on an iPhone.
const safariIPhone = "Mozilla/5.0 (iPhone; CPU iPhone OS 18_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Mobile/15E148 Safari/604.1"

// webCall is one call's context with userAgent, and the call's info, which
// holds the response headers once the call returns.
func webCall(userAgent string) (context.Context, *connect.CallInfo) {
	ctx, info := connect.NewClientContext(context.Background())
	info.RequestHeader().Set("User-Agent", userAgent)
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
		h.signUp(t, email, password)
		ctx, info := webCall(safariIPhone)
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
		}{
			{"web without a mode", web, casav1.DeviceMode_DEVICE_MODE_UNSPECIFIED},
			{"ios in shared mode", ios, shared},
			{"ios in personal mode", ios, personal},
		} {
			for name, call := range map[string]func() error{
				"SignIn": func() error {
					_, err := h.client.SignIn(context.Background(), &casav1.SignInRequest{
						RequestContext: identityRC(c.kind), Email: email, Password: password, DeviceMode: c.mode,
					})
					return err
				},
				"CompleteSignUp": func() error {
					_, err := h.client.CompleteSignUp(context.Background(), &casav1.CompleteSignUpRequest{
						RequestContext: identityRC(c.kind), VerificationToken: token, Password: password, DeviceMode: c.mode,
					})
					return err
				},
				"CompleteRecovery": func() error {
					_, err := h.client.CompleteRecovery(context.Background(), &casav1.CompleteRecoveryRequest{
						RequestContext: identityRC(c.kind), RecoveryToken: token, NewPassword: next, DeviceMode: c.mode,
					})
					return err
				},
			} {
				err := call()
				wantCode(t, err, connect.CodeInvalidArgument)
				if got := ruleIDs(t, err); !slices.Equal(got, []string{"device_mode.web"}) {
					t.Errorf("%s, %s: rules %v, want [device_mode.web]", c.name, name, got)
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
