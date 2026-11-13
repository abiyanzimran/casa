package api_test

import (
	"context"
	"errors"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"reflect"
	"slices"
	"strings"
	"sync/atomic"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/go/protovalidate"
	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
	"connectrpc.com/connect/v2/connectproto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/abiyanzimran/casa/server/internal/api"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

// probeSessions are the sessions the probe's authentication accepts, by
// credential: a native one per bearer token and a web one for the cookie, so
// each client kind calls with a session of its own kind (plan 1.5 decision
// 11 step 5).
var probeSessions = map[string]identity.Session{
	"probe-token":   {ClientKind: "ios", DeviceMode: "personal"},
	"probe-android": {ClientKind: "android", DeviceMode: "personal"},
	"probe-cookie":  {ClientKind: "web", DeviceMode: "personal"},
}

// probeAuthenticate accepts the probe's credentials of kind ("web" for the
// cookie, else a bearer token).
func probeAuthenticate(web bool) func(context.Context, string) (identity.Session, error) {
	return func(_ context.Context, credential string) (identity.Session, error) {
		session, ok := probeSessions[credential]
		if !ok || (session.ClientKind == "web") != web {
			return identity.Session{}, identity.ErrUnauthenticated
		}
		return session, nil
	}
}

const uuidV4 = "0f8b2a6e-3c1d-4e5f-9a7b-1c2d3e4f5a6b" // a UUIDv4

// probe is a real Connect server over HTTP with the synthetic ProbeService:
// Read (no command option), Write (a command, hold_seconds 5), and
// Schemaless (registered without its descriptor). The handlers count the
// requests they actually receive.
type probe struct {
	file    protoreflect.FileDescriptor
	url     string
	client  *connect.Client
	gzipped *connect.Client // sends gzip-compressed requests
	calls   map[string]*atomic.Int32
}

func startProbe(t *testing.T) *probe {
	t.Helper()
	p := &probe{
		file:  probeFile(t, probeMethod{name: "Read"}, probeMethod{name: "Write", command: &casav1.CommandOptions{HoldSeconds: proto.Uint32(5)}}),
		calls: map[string]*atomic.Int32{},
	}
	server := api.NewServer(api.Version{Major: 1, Minor: 9}, testOrigin, api.Sessions{
		Bearer: probeAuthenticate(false),
		Cookie: func(ctx context.Context, cookie string) (identity.CookieSession, error) {
			session, err := probeAuthenticate(true)(ctx, cookie)
			return identity.CookieSession{Session: session}, err
		},
	})
	methods := p.file.Services().Get(0).Methods()
	read := methods.ByName("Read")
	for i := range methods.Len() {
		m := methods.Get(i)
		server.Register(p.method(string(m.Name()), m, m.Input()))
	}
	server.Register(p.method("Schemaless", nil, read.Input()))
	srv := httptest.NewServer(api.NewMux(server, nil))
	t.Cleanup(srv.Close)
	p.url = srv.URL
	p.client = connect.NewClient(connecthttp.NewTransport(srv.Client(), srv.URL))
	p.gzipped = connect.NewClient(connecthttp.NewTransport(srv.Client(), srv.URL, connecthttp.WithSendGzip()))
	return p
}

// method registers name with schema as Spec.Schema (nil for Schemaless).
func (p *probe) method(name string, schema protoreflect.MethodDescriptor, input protoreflect.MessageDescriptor) connect.Method {
	calls := &atomic.Int32{}
	p.calls[name] = calls
	spec := connect.Spec{Procedure: "/casa.apitest.v1.ProbeService/" + name}
	if schema != nil {
		spec.Schema = schema
	}
	return connect.Method{Spec: spec, Handler: func(ctx context.Context, _ connect.Spec, stream connect.ServerStream) error {
		if err := stream.Receive(dynamicpb.NewMessage(input)); err != nil {
			return err
		}
		if session, ok := api.SessionFrom(ctx); !ok || !slices.ContainsFunc(slices.Collect(maps.Values(probeSessions)), func(s identity.Session) bool { return reflect.DeepEqual(s, session) }) {
			return connect.NewError(connect.CodeInternal, "the handler did not get the session")
		}
		calls.Add(1)
		return stream.Send(&emptypb.Empty{})
	}}
}

// call sends a ProbeRequest with rc (nil: no request_context) and note, in a
// valid session of rc's client kind: a web one in the cookie from the app's
// origin, as a browser sends it, or a native one as a bearer token.
func (p *probe) call(client *connect.Client, method string, rc *casav1.RequestContext, note string) error {
	switch rc.GetClientKind() {
	case casav1.ClientKind_CLIENT_KIND_WEB:
		return p.send(client, map[string]string{"Cookie": "__Host-casa_session=probe-cookie", "Origin": testOrigin}, method, rc, note)
	case casav1.ClientKind_CLIENT_KIND_ANDROID:
		return p.callWith(client, "Bearer probe-android", method, rc, note)
	}
	return p.callWith(client, "Bearer probe-token", method, rc, note)
}

// callWith is call with the Authorization header authorization ("": none).
func (p *probe) callWith(client *connect.Client, authorization, method string, rc *casav1.RequestContext, note string) error {
	header := map[string]string{}
	if authorization != "" {
		header["Authorization"] = authorization
	}
	return p.send(client, header, method, rc, note)
}

// send is call with the request headers header.
func (p *probe) send(client *connect.Client, header map[string]string, method string, rc *casav1.RequestContext, note string) error {
	input := p.file.Services().Get(0).Methods().ByName("Read").Input()
	req := dynamicpb.NewMessage(input)
	if rc != nil {
		req.Set(input.Fields().ByName("request_context"), protoreflect.ValueOfMessage(rc.ProtoReflect()))
	}
	req.Set(input.Fields().ByName("note"), protoreflect.ValueOfString(note))
	spec := connect.Spec{Procedure: "/casa.apitest.v1.ProbeService/" + method}
	ctx, info := connect.NewClientContext(context.Background())
	for name, value := range header {
		info.RequestHeader().Set(name, value)
	}
	return client.CallUnary(ctx, spec, req, &emptypb.Empty{})
}

func rc(kind casav1.ClientKind, version, key string) *casav1.RequestContext {
	return &casav1.RequestContext{ClientKind: kind, ClientVersion: version, IdempotencyKey: key}
}

// details decodes the error details of err.
func details(t *testing.T, err error) []proto.Message {
	t.Helper()
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		t.Fatalf("error %v is not a *connect.Error", err)
	}
	var messages []proto.Message
	for _, detail := range connectErr.Details() {
		message, err := connectproto.UnmarshalErrorDetail(detail)
		if err != nil {
			t.Fatalf("error detail %s: %v", detail.Type, err)
		}
		messages = append(messages, message)
	}
	return messages
}

// errorCodes returns the codes of the casa.v1.ErrorDetail details of err.
func errorCodes(t *testing.T, err error) []casav1.ErrorCode {
	t.Helper()
	var codes []casav1.ErrorCode
	for _, detail := range details(t, err) {
		if detail, ok := detail.(*casav1.ErrorDetail); ok {
			codes = append(codes, detail.GetCode())
		}
	}
	return codes
}

// violationPaths returns the field paths of the buf.validate.Violations
// details of err.
func violationPaths(t *testing.T, err error) []string {
	t.Helper()
	var paths []string
	for _, detail := range details(t, err) {
		if violations, ok := detail.(*validate.Violations); ok {
			for _, violation := range violations.GetViolations() {
				paths = append(paths, protovalidate.FieldPathString(violation.GetField()))
			}
		}
	}
	return paths
}

func wantCode(t *testing.T, err error, want connect.Code) {
	t.Helper()
	if got := connect.CodeOf(err); err == nil || got != want {
		t.Fatalf("error = %v (code %v), want code %v", err, got, want)
	}
}

func wantOK(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
}

func (p *probe) wantCalls(t *testing.T, method string, want int32) {
	t.Helper()
	if got := p.calls[method].Load(); got != want {
		t.Fatalf("%s handler received %d requests, want %d", method, got, want)
	}
}

func TestBoundaryChecks(t *testing.T) {
	web, ios, android := casav1.ClientKind_CLIENT_KIND_WEB, casav1.ClientKind_CLIENT_KIND_IOS, casav1.ClientKind_CLIENT_KIND_ANDROID
	upgradeRequired := []casav1.ErrorCode{casav1.ErrorCode_ERROR_CODE_CLIENT_UPGRADE_REQUIRED}

	t.Run("SEC-169 (d) a write command from a web client below the minimum version is rejected with CLIENT_UPGRADE_REQUIRED and not applied", func(t *testing.T) {
		p := startProbe(t)
		err := p.call(p.client, "Write", rc(web, "1.8.10", uuidV4), "")
		wantCode(t, err, connect.CodeFailedPrecondition)
		if got := errorCodes(t, err); !slices.Equal(got, upgradeRequired) {
			t.Fatalf("error codes = %v, want %v", got, upgradeRequired)
		}
		p.wantCalls(t, "Write", 0)
	})
	t.Run("SEC-169 (d) the minimum version is compared by number, not by text", func(t *testing.T) {
		p := startProbe(t)
		wantOK(t, p.call(p.client, "Write", rc(web, "1.10.0", uuidV4), ""))
		wantOK(t, p.call(p.client, "Write", rc(web, "1.9.0", uuidV4), "")) // equal to the minimum
		p.wantCalls(t, "Write", 2)
	})
	t.Run("SEC-169 (d) a read from a web client below the minimum version is served", func(t *testing.T) {
		p := startProbe(t)
		wantOK(t, p.call(p.client, "Read", rc(web, "0.1.0", ""), ""))
		p.wantCalls(t, "Read", 1)
	})
	t.Run("SEC-169 (d) native clients have no minimum version in the pilot", func(t *testing.T) {
		// ADR-028 web note item 5.
		p := startProbe(t)
		wantOK(t, p.call(p.client, "Write", rc(ios, "0.0.1", uuidV4), ""))
		wantOK(t, p.call(p.client, "Write", rc(android, "0.0.1", uuidV4), ""))
		p.wantCalls(t, "Write", 2)
	})
	t.Run("SEC-169 (d) a below-minimum command the current contract finds invalid still gets CLIENT_UPGRADE_REQUIRED", func(t *testing.T) {
		// An old client must get a retryable code that keeps the command in
		// its queue, not a rejection that drops it.
		p := startProbe(t)
		err := p.call(p.client, "Write", rc(web, "1.0.0", uuidV4), "toolong")
		wantCode(t, err, connect.CodeFailedPrecondition)
		if got := errorCodes(t, err); !slices.Equal(got, upgradeRequired) {
			t.Fatalf("error codes = %v, want %v", got, upgradeRequired)
		}
		p.wantCalls(t, "Write", 0)
	})
	t.Run("SEC-169 (d) a below-minimum command with a malformed idempotency key still gets CLIENT_UPGRADE_REQUIRED", func(t *testing.T) {
		// A later tightening of the key rules must not turn an old client's
		// queued command into invalid_argument (lead ruling M2 on PR #20).
		p := startProbe(t)
		err := p.call(p.client, "Write", rc(web, "1.0.0", "not-a-uuid"), "")
		wantCode(t, err, connect.CodeFailedPrecondition)
		if got := errorCodes(t, err); !slices.Equal(got, upgradeRequired) {
			t.Fatalf("error codes = %v, want %v", got, upgradeRequired)
		}
		p.wantCalls(t, "Write", 0)
	})
	t.Run("SEC-15 a request without a valid session never reaches the handler", func(t *testing.T) {
		p := startProbe(t)
		for _, authorization := range []string{"", "Bearer x", "Basic cHJvYmU6dG9rZW4=", "probe-token", "Bearer  probe-token"} {
			err := p.callWith(p.client, authorization, "Read", rc(web, "2.0.0", ""), "")
			wantCode(t, err, connect.CodeUnauthenticated)
			if got := details(t, err); got != nil {
				t.Fatalf("Authorization %q: details %v, want none (refresh, then sign in)", authorization, got)
			}
		}
		// Checked before the request is read: no validation detail either.
		wantCode(t, p.callWith(p.client, "", "Write", nil, "toolong"), connect.CodeUnauthenticated)
		p.wantCalls(t, "Read", 0)
		p.wantCalls(t, "Write", 0)
	})
	t.Run("a request without request_context never reaches the handler", func(t *testing.T) {
		p := startProbe(t)
		err := p.call(p.client, "Read", nil, "")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violationPaths(t, err), []string{"request_context"}; !slices.Equal(got, want) {
			t.Fatalf("violation paths = %q, want %q", got, want)
		}
		p.wantCalls(t, "Read", 0)
	})
	t.Run("a write command without an idempotency key never reaches the handler (ADR-005)", func(t *testing.T) {
		p := startProbe(t)
		err := p.call(p.client, "Write", rc(web, "2.0.0", ""), "")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violationPaths(t, err), []string{"request_context.idempotency_key"}; !slices.Equal(got, want) {
			t.Fatalf("violation paths = %q, want %q", got, want)
		}
		p.wantCalls(t, "Write", 0)
	})
	t.Run("an idempotency key must be a UUIDv4 (ADR-005; 09.2 §8)", func(t *testing.T) {
		p := startProbe(t)
		for _, key := range []string{"not-a-uuid", "6ba7b810-9dad-11d1-80b4-00c04fd430c8"} { // the second is a version 1 UUID
			err := p.call(p.client, "Write", rc(web, "2.0.0", key), "")
			wantCode(t, err, connect.CodeInvalidArgument)
			if got, want := violationPaths(t, err), []string{"request_context.idempotency_key"}; !slices.Equal(got, want) {
				t.Fatalf("key %q: violation paths = %q, want %q", key, got, want)
			}
		}
		p.wantCalls(t, "Write", 0)
		wantOK(t, p.call(p.client, "Write", rc(web, "2.0.0", strings.ToUpper(uuidV4)), ""))
		p.wantCalls(t, "Write", 1)
	})
	t.Run("an unspecified client kind or a malformed client version is rejected", func(t *testing.T) {
		p := startProbe(t)
		err := p.call(p.client, "Read", rc(casav1.ClientKind_CLIENT_KIND_UNSPECIFIED, "2.0.0", ""), "")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violationPaths(t, err), []string{"request_context.client_kind"}; !slices.Equal(got, want) {
			t.Fatalf("violation paths = %q, want %q", got, want)
		}
		err = p.call(p.client, "Read", rc(web, "1.2", ""), "")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violationPaths(t, err), []string{"request_context.client_version"}; !slices.Equal(got, want) {
			t.Fatalf("violation paths = %q, want %q", got, want)
		}
		p.wantCalls(t, "Read", 0)
	})
	t.Run("a write command from a web client with a malformed client version is rejected", func(t *testing.T) {
		// Step 2 checks client_version before the version check reads it.
		p := startProbe(t)
		err := p.call(p.client, "Write", rc(web, "1.2", uuidV4), "")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violationPaths(t, err), []string{"request_context.client_version"}; !slices.Equal(got, want) {
			t.Fatalf("violation paths = %q, want %q", got, want)
		}
		p.wantCalls(t, "Write", 0)
	})
	t.Run("the request body is validated after the version check", func(t *testing.T) {
		p := startProbe(t)
		err := p.call(p.client, "Write", rc(web, "2.0.0", uuidV4), "toolong")
		wantCode(t, err, connect.CodeInvalidArgument)
		if got, want := violationPaths(t, err), []string{"note"}; !slices.Equal(got, want) {
			t.Fatalf("violation paths = %q, want %q", got, want)
		}
		p.wantCalls(t, "Write", 0)
	})
	t.Run("a method registered without its descriptor is rejected, never treated as a read", func(t *testing.T) {
		p := startProbe(t)
		wantCode(t, p.call(p.client, "Schemaless", rc(web, "0.1.0", ""), ""), connect.CodeInternal)
		p.wantCalls(t, "Schemaless", 0)
	})
	t.Run("GET /healthz answers ok without data (09.4 §3)", func(t *testing.T) {
		p := startProbe(t)
		res, err := http.Get(p.url + "/healthz")
		if err != nil {
			t.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		_ = res.Body.Close()
		if err != nil || res.StatusCode != http.StatusOK || string(body) != "ok\n" {
			t.Fatalf("GET /healthz = %d %q (%v), want 200 \"ok\\n\"", res.StatusCode, body, err)
		}
		res, err = http.Post(p.url+"/healthz", "text/plain", nil)
		if err != nil {
			t.Fatal(err)
		}
		_ = res.Body.Close()
		if res.StatusCode != http.StatusMethodNotAllowed {
			t.Fatalf("POST /healthz = %d, want 405", res.StatusCode)
		}
	})
	t.Run("a request over 1 MiB is refused before validation", func(t *testing.T) {
		p := startProbe(t)
		// Just under the limit, the request reaches validation (note is too long).
		wantCode(t, p.call(p.client, "Read", rc(web, "2.0.0", ""), strings.Repeat("a", 1<<20-64)), connect.CodeInvalidArgument)
		err := p.call(p.client, "Read", rc(web, "2.0.0", ""), strings.Repeat("a", 1<<20+1))
		wantCode(t, err, connect.CodeResourceExhausted)
		if got := violationPaths(t, err); got != nil {
			t.Fatalf("violation paths = %q, want none: the limit acts before validation", got)
		}
		p.wantCalls(t, "Read", 0)
	})
	t.Run("a compressed request over 1 MiB once decompressed is refused", func(t *testing.T) {
		p := startProbe(t)
		wantCode(t, p.call(p.gzipped, "Read", rc(web, "2.0.0", ""), strings.Repeat("a", 2<<20)), connect.CodeResourceExhausted)
		p.wantCalls(t, "Read", 0)
	})
}
