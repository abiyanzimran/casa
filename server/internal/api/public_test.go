package api

import (
	"strings"
	"testing"
)

func TestOnlyIdentityProceduresArePublic(t *testing.T) {
	// 09.2 §1 (api row): only casa.v1.IdentityService procedures used before
	// a session exists may skip authentication (SEC-15).
	if len(publicProcedures) == 0 {
		t.Fatal("no public procedure: sign-up and sign-in would be unreachable")
	}
	for procedure := range publicProcedures {
		if !strings.HasPrefix(procedure, "/casa.v1.IdentityService/") {
			t.Errorf("public procedure %s is outside casa.v1.IdentityService", procedure)
		}
	}
}
