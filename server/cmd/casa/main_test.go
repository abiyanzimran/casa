package main

import (
	"bytes"
	"testing"
)

func TestRunWithoutCommandPrintsUsageAndExits2(t *testing.T) {
	var errOut bytes.Buffer
	if code := run(nil, &errOut); code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
	if got, want := errOut.String(), "usage: casa <command>\n"; got != want {
		t.Fatalf("stderr = %q, want %q", got, want)
	}
}

func TestRunUnknownCommandExits2(t *testing.T) {
	var errOut bytes.Buffer
	if code := run([]string{"nope"}, &errOut); code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
	want := "casa: unknown command \"nope\"\n" + "usage: casa <command>\n"
	if got := errOut.String(); got != want {
		t.Fatalf("stderr = %q, want %q", got, want)
	}
}
