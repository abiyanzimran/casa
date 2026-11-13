// Package api is the Connect API layer (09.2 §1): the server every service
// registers on, with the boundary checks each RPC passes before its handler
// sees the request (ES §4), and the HTTP routes of the casa binary.
package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/go/protovalidate"
	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
	"connectrpc.com/connect/v2/connectproto"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

// maxRequestBytes bounds every request message, on the wire and once
// decompressed [Usulan] (lead ruling Q9 on the 1.3 plan; 09.4).
const maxRequestBytes = 1 << 20

// requestContextField is field 1 of every RPC request (14 §4.6; the contract
// test in this package enforces it).
const requestContextField = "request_context"

var requestContextName = (&casav1.RequestContext{}).ProtoReflect().Descriptor().FullName()

// NewServer returns the Connect server every service registers on. Before a
// handler sees a request, it checks, in this order (plan 1.5 decision 11):
//
//  1. the method has its descriptor, else CodeInternal (fail closed: without
//     it a write command cannot be told from a read);
//  2. a request carrying the session cookie, an Origin or a Sec-Fetch-Site
//     header, whatever its procedure and claimed client kind, is a POST from
//     origin, else CodePermissionDenied, before the request is received
//     (SEC-162 (a), (b); ADR-030 item 6; checkOrigin);
//  3. outside publicProcedures, "Authorization: Bearer <access token>" is
//     accepted by sessions.Bearer (SEC-15; 09.4 §2), or without it the
//     session cookie by sessions.Cookie (SEC-161 (a)), before the request is
//     received, else CodeUnauthenticated; a cookie value rotated by the
//     authentication is set on the response here, so it is sent even when
//     the handler fails; the handler reads the session with SessionFrom;
//  4. request_context is present and its client_kind and client_version are
//     valid (protovalidate), else CodeInvalidArgument;
//  5. a web client's request passes the check of step 2 too, so a web RPC
//     without Origin is rejected (SEC-162 (a) note 1.5 (2)); an
//     authenticated request's client_kind is the session's, else
//     CodeInvalidArgument; a user-initiated request of a shared web session
//     moves its idle limit through sessions.UserAction (SEC-165 (g));
//  6. a write command from a web client is at least minimumWebVersion, else
//     CodeFailedPrecondition with CLIENT_UPGRADE_REQUIRED (ADR-030 item 9;
//     SEC-169 (d));
//     6a. a write command that is not notice_exempt, or a procedure in
//     noticeHeldProcedures, from an account noticeHeld reports as held, is
//     CodeFailedPrecondition with NOTICE_UNREAD (SEC-34; 09.3 §1
//     "Mekanisme"; plan 1.6 decision 7);
//  7. a write command carries an idempotency key, else CodeInvalidArgument
//     (ADR-005; 09.2 §8);
//  8. the whole request, idempotency key form included, is valid
//     (protovalidate), else CodeInvalidArgument.
//
// The hold of step 6a defers, it does not refuse: it comes before the key and
// the body, so the command is neither validated in full nor recorded in
// idempotency_record, and the same command is applied once after the
// acknowledgement (09.3 §1). The version is checked before the body and
// before the key's form, so an old client whose request no longer passes the
// current rules still gets the retryable CLIENT_UPGRADE_REQUIRED and keeps
// the command in its queue, rather than a rejection that drops it (SEC-169
// (d); ADR-030 item 9). A write command is a method with the casa.v1.command
// option (09.3 §1; lead ruling Q4).
// Validation errors carry a buf.validate.Violations detail (lead ruling Q5).
// Connect serves GET only for NO_SIDE_EFFECTS methods, which the contract
// has none of; step 2 would refuse one with the cookie.
func NewServer(minimumWebVersion Version, origin string, sessions Sessions, noticeHeld func(context.Context, pgtype.UUID) (bool, error)) *connect.Server {
	return connect.NewServer(boundary{minimumWebVersion: minimumWebVersion, origin: origin, sessions: sessions, noticeHeld: noticeHeld}.intercept)
}

// noticeHeldProcedures are the online RPCs outside the command kind table
// that NOTICE_UNREAD still holds, by full procedure name (10.1 §11 note of
// 14 November 2026 (d); plan 1.6 decision 7). The identity RPCs and reads
// are never held (09.3 §1).
var noticeHeldProcedures = map[string]bool{}

// Sessions authenticates the requests outside publicProcedures.
type Sessions struct {
	Bearer     func(context.Context, string) (identity.Session, error)       // identity.Service.Authenticate
	Cookie     func(context.Context, string) (identity.CookieSession, error) // identity.Service.AuthenticateCookie
	UserAction func(context.Context, identity.Session) error                 // identity.Service.TouchUserAction
}

// publicProcedures are the only procedures served without a session: the
// identity RPCs used before one exists (09.2 §1; SEC-15).
var publicProcedures = map[string]bool{
	casav1connect.IdentityServiceStartSignUpProcedure:      true,
	casav1connect.IdentityServiceCompleteSignUpProcedure:   true,
	casav1connect.IdentityServiceSignInProcedure:           true,
	casav1connect.IdentityServiceRefreshSessionProcedure:   true,
	casav1connect.IdentityServiceStartRecoveryProcedure:    true,
	casav1connect.IdentityServiceCompleteRecoveryProcedure: true,
}

type sessionKey struct{}

// SessionFrom returns the session a request was authenticated with; false
// for a public procedure.
func SessionFrom(ctx context.Context) (identity.Session, bool) {
	session, ok := ctx.Value(sessionKey{}).(identity.Session)
	return session, ok
}

// NewMux serves the procedures of server, GET /healthz and, when bundle is
// not nil, the web client on every other path (ADR-030 item 3); without a
// bundle the static routes answer 404 (plan decision 2). Every answer carries
// the security headers (WithSecurityHeaders; SEC-166).
func NewMux(server *connect.Server, bundle *Bundle) http.Handler {
	mux := http.NewServeMux()
	// The uptime monitor's check: no data and no database (09.4 §3; 14 §6.4).
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "ok\n") // a failed write leaves the monitor to report the outage
	})
	connecthttp.Mount(mux, server, connecthttp.WithReadMaxBytes(maxRequestBytes))
	if bundle != nil {
		// Every method, so Bundle answers other than GET and HEAD with 405.
		mux.Handle("/", bundle)
	}
	return WithSecurityHeaders(mux, bundle)
}

type boundary struct {
	minimumWebVersion Version
	// origin is the app's origin, scheme://host[:port], as browsers send it.
	origin   string
	sessions Sessions
	// noticeHeld reports whether an account has an unacknowledged notice
	// version (NoticeHeld).
	noticeHeld func(context.Context, pgtype.UUID) (bool, error)
}

func (b boundary) intercept(next connect.ServerFunc) connect.ServerFunc {
	return func(ctx context.Context, spec connect.Spec, stream connect.ServerStream) error {
		method, ok := spec.Schema.(protoreflect.MethodDescriptor)
		if !ok {
			return connect.Errorf(connect.CodeInternal, "procedure %s has no method descriptor", spec.Procedure)
		}
		header := &connect.Header{}
		if info, ok := connect.CallInfoForServerContext(ctx); ok {
			header = info.RequestHeader()
		}
		cookie, withCookie := sessionCookieValue(header)
		if withCookie || header.Has("Origin") || header.Has("Sec-Fetch-Site") {
			if err := b.sameOrigin(ctx, header); err != nil {
				return err
			}
		}
		if !publicProcedures[spec.Procedure] {
			session, err := b.session(ctx, header, cookie, withCookie)
			if err != nil {
				return err
			}
			ctx = context.WithValue(ctx, sessionKey{}, session)
		}
		command, _ := proto.GetExtension(method.Options(), casav1.E_Command).(*casav1.CommandOptions)
		isCommand := proto.HasExtension(method.Options(), casav1.E_Command)
		return next(ctx, spec, &checkedStream{
			ServerStream: stream, ctx: ctx, header: header, boundary: b, method: method, command: isCommand,
			held: isCommand && !command.GetNoticeExempt() || noticeHeldProcedures[spec.Procedure],
		})
	}
}

// sameOrigin is step 2 of NewServer: a POST whose Origin is the app's
// (SEC-162 (a), (b)).
func (b boundary) sameOrigin(ctx context.Context, header *connect.Header) error {
	if info, ok := connecthttp.ServerInfoForContext(ctx); !ok || info.HTTPMethod() != http.MethodPost {
		return crossOriginRejected()
	}
	return checkOrigin(header, b.origin)
}

// checkOrigin accepts a request only if it has exactly one Origin, equal to
// origin as a string (no case folding, prefix or suffix match: a staging
// subdomain is same-site, so SameSite=Strict does not hold it, SEC-41 (d)),
// and a Sec-Fetch-Site, if any, of same-origin (SEC-162 (a); ADR-030 item
// 6). /v1/stream reuses it (SEC-163).
func checkOrigin(header *connect.Header, origin string) error {
	if values := header.Values("Origin"); len(values) != 1 || values[0] != origin {
		return crossOriginRejected()
	}
	if site := header.Values("Sec-Fetch-Site"); len(site) > 1 || len(site) == 1 && site[0] != "same-origin" {
		return crossOriginRejected()
	}
	return nil
}

// crossOriginRejected is the one answer of the origin check, without detail.
func crossOriginRejected() error {
	return connect.NewError(connect.CodePermissionDenied, "cross-origin request rejected")
}

// session is step 3 of NewServer: the bearer access token, or without one
// the session cookie.
func (b boundary) session(ctx context.Context, header *connect.Header, cookie string, withCookie bool) (identity.Session, error) {
	token, bearer := strings.CutPrefix(header.Get("Authorization"), "Bearer ")
	switch {
	case bearer:
		session, err := b.sessions.Bearer(ctx, token)
		if err != nil {
			return identity.Session{}, identityError(err)
		}
		return session, nil
	case withCookie:
		session, err := b.sessions.Cookie(ctx, cookie)
		if err != nil {
			return identity.Session{}, identityError(err)
		}
		if session.NewCookie != nil {
			if err := setSessionCookie(ctx, *session.NewCookie); err != nil {
				return identity.Session{}, identityError(err)
			}
		}
		return session.Session, nil
	}
	return identity.Session{}, identityError(identity.ErrUnauthenticated)
}

// checkedStream runs the boundary checks on each request message as it is
// received, so the handler gets either a checked request or the error.
type checkedStream struct {
	connect.ServerStream
	ctx      context.Context // the request's, with its session: step 5 of NewServer reads it
	header   *connect.Header
	boundary boundary
	method   protoreflect.MethodDescriptor
	command  bool // a write command (09.3 §1)
	held     bool // subject to NOTICE_UNREAD (step 6a)
}

func (s *checkedStream) Receive(msg any) error {
	if err := s.ServerStream.Receive(msg); err != nil {
		return err
	}
	request, ok := msg.(proto.Message)
	if !ok {
		return connect.Errorf(connect.CodeInternal, "procedure %s received a %T, not a proto message", s.method.FullName(), msg)
	}
	return s.boundary.check(s.ctx, s.header, request, s.command, s.held)
}

func (b boundary) check(ctx context.Context, header *connect.Header, request proto.Message, command, held bool) error {
	descriptor := request.ProtoReflect().Descriptor()
	if err := protovalidate.Validate(request, protovalidate.WithFilter(requestContextOnly(descriptor))); err != nil {
		return validationError(err)
	}
	requestContext, err := readRequestContext(request)
	if err != nil {
		return connect.NewError(connect.CodeInternal, "request context unreadable").WithCause(err)
	}
	if err := b.checkClient(ctx, header, requestContext); err != nil {
		return err
	}
	if command && requestContext.GetClientKind() == casav1.ClientKind_CLIENT_KIND_WEB {
		version, err := ParseVersion(requestContext.GetClientVersion())
		if err != nil { // unreachable: step 2 checked the same pattern
			return connect.NewError(connect.CodeInternal, "client version unreadable").WithCause(err)
		}
		if version.Less(b.minimumWebVersion) {
			return errorWithDetail(connect.CodeFailedPrecondition, "client upgrade required",
				&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_CLIENT_UPGRADE_REQUIRED})
		}
	}
	// WRITE_FROZEN comes before this step when ADR-028 item 3 is built
	// (09.3 §1 order: protective, WRITE_FROZEN, NOTICE_UNREAD).
	if held {
		if err := b.checkNotice(ctx); err != nil {
			return err
		}
	}
	if command && requestContext.GetIdempotencyKey() == "" {
		return missingIdempotencyKey(descriptor)
	}
	if err := protovalidate.Validate(request); err != nil {
		return validationError(err)
	}
	return nil
}

// checkNotice is step 6a of NewServer, in its own transaction (09.3 §1
// "Mekanisme").
func (b boundary) checkNotice(ctx context.Context) error {
	session, ok := SessionFrom(ctx)
	if !ok { // unreachable: no public procedure is a command or held; fail closed
		return connect.NewError(connect.CodeInternal, "internal error")
	}
	held, err := b.noticeHeld(ctx, session.AccountID)
	if err != nil {
		return householdError(err)
	}
	if held {
		return errorWithDetail(connect.CodeFailedPrecondition, "notice unread",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_NOTICE_UNREAD})
	}
	return nil
}

// checkClient is step 5 of NewServer.
func (b boundary) checkClient(ctx context.Context, header *connect.Header, requestContext *casav1.RequestContext) error {
	kind := clientKindOf(requestContext)
	if kind == "web" { // SEC-162 (a) note 1.5 (2): closes login CSRF on the public RPCs
		if err := b.sameOrigin(ctx, header); err != nil {
			return err
		}
	}
	session, ok := SessionFrom(ctx)
	if !ok {
		return nil
	}
	if session.ClientKind != kind {
		return connect.NewError(connect.CodeInvalidArgument, "client kind does not match the session")
	}
	if requestContext.GetUserInitiated() && session.ClientKind == "web" && session.DeviceMode == "shared" {
		// SEC-165 (g): only a user's action moves a shared session's idle limit.
		if err := b.sessions.UserAction(ctx, session); err != nil {
			return identityError(err)
		}
	}
	return nil
}

// requestContextOnly limits validation to what the version check needs: the
// presence of request_context and the rules of its client_kind and
// client_version. Everything else, the idempotency key's form included, waits
// for step 5, so a later tightening of those rules never turns an old
// client's queued command into invalid_argument instead of
// CLIENT_UPGRADE_REQUIRED (SEC-169 (d); lead ruling M2 on PR #20).
func requestContextOnly(request protoreflect.MessageDescriptor) protovalidate.Filter {
	return protovalidate.FilterFunc(func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
		switch parent := d.Parent(); {
		case parent != nil && parent.FullName() == request.FullName():
			return d.Name() == requestContextField
		case parent != nil && parent.FullName() == requestContextName:
			return d.Name() == "client_kind" || d.Name() == "client_version"
		default:
			// The request and request_context themselves: their message
			// rules wait for step 5.
			return false
		}
	})
}

// readRequestContext copies request_context through its wire form, so a
// generated request and a dynamic one take the same path.
func readRequestContext(request proto.Message) (*casav1.RequestContext, error) {
	message := request.ProtoReflect()
	field := message.Descriptor().Fields().ByName(requestContextField)
	if field == nil || field.Message() == nil || field.Message().FullName() != requestContextName {
		return nil, fmt.Errorf("%s has no %s field of type %s", message.Descriptor().FullName(), requestContextField, requestContextName)
	}
	wire, err := proto.Marshal(message.Get(field).Message().Interface())
	if err != nil {
		return nil, err
	}
	requestContext := &casav1.RequestContext{}
	return requestContext, proto.Unmarshal(wire, requestContext)
}

// missingIdempotencyKey reports request_context.idempotency_key as required,
// in the same Violations form as protovalidate's own required rule.
func missingIdempotencyKey(request protoreflect.MessageDescriptor) error {
	contextField := request.Fields().ByName(requestContextField)
	keyField := contextField.Message().Fields().ByName("idempotency_key")
	violation := validate.Violation_builder{
		Field: validate.FieldPath_builder{
			Elements: []*validate.FieldPathElement{fieldPathElement(contextField), fieldPathElement(keyField)},
		}.Build(),
		RuleId:  proto.String("required"),
		Message: proto.String("value is required for a write command"),
	}.Build()
	return errorWithDetail(connect.CodeInvalidArgument, "invalid request",
		validate.Violations_builder{Violations: []*validate.Violation{violation}}.Build())
}

func fieldPathElement(field protoreflect.FieldDescriptor) *validate.FieldPathElement {
	return validate.FieldPathElement_builder{
		FieldNumber: proto.Int32(int32(field.Number())),
		FieldName:   proto.String(string(field.Name())),
		FieldType:   descriptorpb.FieldDescriptorProto_Type(field.Kind()).Enum(),
	}.Build()
}

func validationError(err error) error {
	var invalid *protovalidate.ValidationError
	if errors.As(err, &invalid) {
		return errorWithDetail(connect.CodeInvalidArgument, "invalid request", invalid.ToProto())
	}
	// A compilation or runtime error is a broken rule in the contract, not
	// the caller's fault.
	return connect.NewError(connect.CodeInternal, "request validation failed").WithCause(err)
}

func errorWithDetail(code connect.Code, message string, detail proto.Message) error {
	errorDetail, err := connectproto.NewErrorDetail(detail)
	if err != nil {
		return connect.NewError(connect.CodeInternal, "error detail unencodable").WithCause(err)
	}
	return connect.NewError(code, message).WithDetail(errorDetail)
}
