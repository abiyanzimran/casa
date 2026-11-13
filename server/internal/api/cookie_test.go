package api

import (
	"context"
	"strings"
	"testing"

	"github.com/abiyanzimran/casa/server/internal/identity"
)

func TestSetSessionCookieFailsClosed(t *testing.T) {
	// SEC-161 (b): a web session whose cookie cannot be set must not be
	// answered as opened; the error never carries the value (SEC-44).
	value := strings.Repeat("v", 43)
	err := setSessionCookie(context.Background(), identity.Cookie{Value: value})
	if err == nil || strings.Contains(err.Error(), value) {
		t.Fatalf("setSessionCookie without call info: %v; want an error without the value", err)
	}
}
