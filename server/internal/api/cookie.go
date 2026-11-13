package api

import (
	"context"
	"net/http"
	"time"

	"connectrpc.com/connect/v2"

	"github.com/abiyanzimran/casa/server/internal/identity"
)

// sessionCookieName is the web session's cookie: the __Host- prefix makes
// the browser require Secure and Path=/ and refuse Domain (SEC-161 (b); plan
// 1.5 decision 7; ADR-030 item 5).
const sessionCookieName = "__Host-casa_session"

// setSessionCookie adds c to the response as the session cookie: Path=/,
// Secure, HttpOnly, SameSite=Strict, no Domain, and Max-Age in whole seconds
// rounded down, or none for a session cookie (plan 1.5 decision 7). The
// value goes only into this header, never into a body or a log (SEC-161
// (b); SEC-44).
func setSessionCookie(ctx context.Context, c identity.Cookie) {
	cookie := http.Cookie{
		Name: sessionCookieName, Value: c.Value, Path: "/",
		Secure: true, HttpOnly: true, SameSite: http.SameSiteStrictMode,
		MaxAge: int(c.MaxAge / time.Second),
	}
	if info, ok := connect.CallInfoForServerContext(ctx); ok {
		info.ResponseHeader().Add("Set-Cookie", cookie.String())
	}
}
