package main

import (
	"testing"
)

func TestRunWithoutCommandPrintsUsageAndExits2(t *testing.T) {
	code, _, stderr := runCasa(t)
	if code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
	if got, want := stderr, "usage: casa <command>\n"; got != want {
		t.Fatalf("stderr = %q, want %q", got, want)
	}
}

func TestRunUnknownCommandExits2(t *testing.T) {
	code, _, stderr := runCasa(t, "nope")
	if code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
	want := "casa: unknown command \"nope\"\n" + "usage: casa <command>\n"
	if got := stderr; got != want {
		t.Fatalf("stderr = %q, want %q", got, want)
	}
}
