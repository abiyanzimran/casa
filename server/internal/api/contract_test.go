package api_test

import (
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
)

// protoFilesOnDisk counts the .proto files of the contract, so a package
// casaFiles does not see cannot pass the contract test unchecked.
func protoFilesOnDisk(t *testing.T) int {
	t.Helper()
	count := 0
	err := filepath.WalkDir("../../../proto/casa", func(path string, entry fs.DirEntry, err error) error {
		if err == nil && !entry.IsDir() && strings.HasSuffix(path, ".proto") {
			count++
		}
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return count
}

// TestContract holds the rules every RPC of the contract follows (14 §4.6;
// 09.3 §1; 09.1 §4). Each rule runs on the real contract (no violation) and on
// a probe file that breaks only that rule (exactly that violation), so a rule
// that checks nothing fails here, not later.
func TestContract(t *testing.T) {
	t.Run("the real contract has no violation", func(t *testing.T) {
		files := casaFiles(t)
		if want := protoFilesOnDisk(t); len(files) != want {
			t.Fatalf("%d casa.* files registered, %d .proto files under proto/casa: import every contract package in probe_test.go", len(files), want)
		}
		if got := contractViolations(files); got != nil {
			t.Fatalf("violations:\n%v", got)
		}
	})
	t.Run("a compliant command has no violation", func(t *testing.T) {
		file := probeFile(t, probeMethod{name: "Ok", command: &casav1.CommandOptions{HoldSeconds: proto.Uint32(5)}})
		if got := contractViolations([]protoreflect.FileDescriptor{file}); got != nil {
			t.Fatalf("violations: %v", got)
		}
	})
	cases := []struct {
		name   string
		method probeMethod
		want   string
	}{
		{
			"every RPC request starts with a required request_context",
			probeMethod{name: "Bare", bareRequest: true},
			"casa.apitest.v1.ProbeService.Bare: request field 1 is not a required casa.v1.RequestContext named request_context",
		},
		{
			"no RPC streams from the client (09.1 §4; ADR-013 web note item 5)",
			probeMethod{name: "Upload", clientStreaming: true},
			"casa.apitest.v1.ProbeService.Upload: client streaming is not allowed",
		},
		{
			"every command declares hold_seconds",
			probeMethod{name: "NoHold", command: &casav1.CommandOptions{}},
			"casa.apitest.v1.ProbeService.NoHold: command without hold_seconds",
		},
		// 09.3 §1 "Uji kontrak"; 14 §4.6. SEC-151 is cited, not named in the
		// test name: no protective command exists yet.
		{
			"a protective command has hold_seconds 0",
			probeMethod{name: "Leave", command: &casav1.CommandOptions{HoldSeconds: proto.Uint32(5), Protective: true, NoticeExempt: true}},
			"casa.apitest.v1.ProbeService.Leave: protective with hold_seconds 5",
		},
		// 09.3 §1 NOTICE_UNREAD "Uji kontrak: protective ⊆ notice_exempt".
		// SEC-34 is cited, not named: no protective command exists yet.
		{
			"a protective command is notice_exempt",
			probeMethod{name: "Revoke", command: &casav1.CommandOptions{HoldSeconds: proto.Uint32(0), Protective: true}},
			"casa.apitest.v1.ProbeService.Revoke: protective but not notice_exempt",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := contractViolations([]protoreflect.FileDescriptor{probeFile(t, c.method)})
			if want := []string{c.want}; !slices.Equal(got, want) {
				t.Fatalf("violations = %q, want %q", got, want)
			}
		})
	}
}
