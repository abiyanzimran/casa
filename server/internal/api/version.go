package api

import (
	"cmp"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
)

// MinimumWebVersion is the default lowest web client version allowed to send
// a write command; casa serve -minimum-web-version overrides it, and the
// operator raises it on a release that fixes a security hole (ADR-030 item 9;
// TM-45; SEC-169 (d)). Native clients have no minimum in the pilot (ADR-028
// web note item 5).
const MinimumWebVersion = "0.0.0"

// Version is a client build version, MAJOR.MINOR.PATCH, compared by number
// (ADR-030 item 9; lead ruling Q1 on the 1.3 plan).
type Version struct{ Major, Minor, Patch uint32 }

// clientVersionPattern is the contract's own rule for client_version, read
// from the descriptor so the format is defined once (ES §2).
var clientVersionPattern = regexp.MustCompile(proto.GetExtension(
	(&casav1.RequestContext{}).ProtoReflect().Descriptor().Fields().ByName("client_version").Options(),
	validate.E_Field,
).(*validate.FieldRules).GetString().GetPattern())

// ParseVersion parses s in the form the contract allows for client_version.
func ParseVersion(s string) (Version, error) {
	if !clientVersionPattern.MatchString(s) {
		return Version{}, fmt.Errorf("version %q is not MAJOR.MINOR.PATCH", s)
	}
	var parts [3]uint32
	for i, part := range strings.Split(s, ".") {
		n, err := strconv.ParseUint(part, 10, 32)
		if err != nil {
			return Version{}, fmt.Errorf("version %q: %w", s, err)
		}
		parts[i] = uint32(n)
	}
	return Version{Major: parts[0], Minor: parts[1], Patch: parts[2]}, nil
}

// Less reports whether v is older than w.
func (v Version) Less(w Version) bool {
	return cmp.Or(cmp.Compare(v.Major, w.Major), cmp.Compare(v.Minor, w.Minor), cmp.Compare(v.Patch, w.Patch)) < 0
}
