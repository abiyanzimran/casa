package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"connectrpc.com/connect/v2"

	"github.com/abiyanzimran/casa/server/internal/identity"
)

// sessionCookieName is the web session's cookie: the __Host- prefix makes
// the browser require Secure and Path=/ and refuse Domain (SEC-161 (b); plan
// 1.5 decision 7; ADR-030 item 5).
const sessionCookieName = "__Host-casa_session"

// sessionCookieValue returns the value of the request's session cookie, read
// as net/http reads a request's cookies (the first of that name), and false
// when there is none.
func sessionCookieValue(header *connect.Header) (string, bool) {
	request := http.Request{Header: http.Header{"Cookie": header.Values("Cookie")}}
	cookie, err := request.Cookie(sessionCookieName)
	if err != nil {
		return "", false
	}
	return cookie.Value, true
}

// setSessionCookie adds c to the response as the session cookie: Path=/,
// Secure, HttpOnly, SameSite=Strict, no Domain, and Max-Age in whole seconds
// rounded down, or none for a session cookie (plan 1.5 decision 7). The
// value goes only into this header, never into a body or a log (SEC-161
// (b); SEC-44). Without the call's response headers it fails closed, so a
// committed session is never answered as opened without its cookie; the
// error never carries the value.
func setSessionCookie(ctx context.Context, c identity.Cookie) error {
	info, ok := connect.CallInfoForServerContext(ctx)
	if !ok {
		return errors.New("no call info to set the session cookie on")
	}
	cookie := http.Cookie{
		Name: sessionCookieName, Value: c.Value, Path: "/",
		Secure: true, HttpOnly: true, SameSite: http.SameSiteStrictMode,
		MaxAge: int(c.MaxAge / time.Second),
	}
	info.ResponseHeader().Add("Set-Cookie", cookie.String())
	return nil
}
