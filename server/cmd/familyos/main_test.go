package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunWithoutCommandPrintsUsageAndExits2(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := run(nil, &out, &errOut); code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
	if !strings.HasPrefix(errOut.String(), "usage: familyos <command>") {
		t.Fatalf("stderr = %q", errOut.String())
	}
	if out.Len() != 0 {
		t.Fatalf("stdout must be empty, got %q", out.String())
	}
}

func TestRunUnknownCommandExits2(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := run([]string{"nope"}, &out, &errOut); code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
	if !strings.Contains(errOut.String(), `unknown command "nope"`) {
		t.Fatalf("stderr = %q", errOut.String())
	}
}
