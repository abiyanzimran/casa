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

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/go/protovalidate"
	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
	"connectrpc.com/connect/v2/connectproto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/schema"
)

// probeSchema exists only to prove the allow-list test.
var probeSchema = schema.Up

// maxRequestBytes bounds every request message, on the wire and once
// decompressed [Usulan] (lead ruling Q9 on the 1.3 plan; 09.4).
const maxRequestBytes = 1 << 20

// requestContextField is field 1 of every RPC request (14 §4.6; the contract
// test in this package enforces it).
const requestContextField = "request_context"

// NewServer returns the Connect server every service registers on. Before a
// handler sees a request, it checks, in this order:
//
//  1. the method has its descriptor, else CodeInternal (fail closed: without
//     it a write command cannot be told from a read);
//  2. request_context is present and valid (protovalidate), else
//     CodeInvalidArgument;
//  3. a write command from a web client is at least minimumWebVersion, else
//     CodeFailedPrecondition with CLIENT_UPGRADE_REQUIRED (ADR-030 item 9;
//     SEC-169 (d));
//  4. a write command carries an idempotency key, else CodeInvalidArgument
//     (ADR-005; 09.2 §8);
//  5. the whole request is valid (protovalidate), else CodeInvalidArgument.
//
// The version is checked before the body, so an old client whose request no
// longer passes the current rules still gets the retryable
// CLIENT_UPGRADE_REQUIRED and keeps the command in its queue, rather than a
// rejection that drops it (SEC-169 (d); ADR-030 item 9). A write command is a
// method with the casa.v1.command option (09.3 §1; lead ruling Q4).
// Validation errors carry a buf.validate.Violations detail (lead ruling Q5).
func NewServer(minimumWebVersion Version) *connect.Server {
	_ = probeSchema
	return connect.NewServer(boundary{minimumWebVersion: minimumWebVersion}.intercept)
}

// NewMux serves the procedures of server and GET /healthz.
func NewMux(server *connect.Server) http.Handler {
	mux := http.NewServeMux()
	// The uptime monitor's check: no data and no database (09.4 §3; 14 §6.4).
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "ok\n") // a failed write leaves the monitor to report the outage
	})
	connecthttp.Mount(mux, server, connecthttp.WithReadMaxBytes(maxRequestBytes))
	return mux
}

type boundary struct {
	minimumWebVersion Version
}

func (b boundary) intercept(next connect.ServerFunc) connect.ServerFunc {
	return func(ctx context.Context, spec connect.Spec, stream connect.ServerStream) error {
		method, ok := spec.Schema.(protoreflect.MethodDescriptor)
		if !ok {
			return connect.Errorf(connect.CodeInternal, "procedure %s has no method descriptor", spec.Procedure)
		}
		return next(ctx, spec, &checkedStream{ServerStream: stream, boundary: b, method: method})
	}
}

// checkedStream runs the boundary checks on each request message as it is
// received, so the handler gets either a checked request or the error.
type checkedStream struct {
	connect.ServerStream
	boundary boundary
	method   protoreflect.MethodDescriptor
}

func (s *checkedStream) Receive(msg any) error {
	if err := s.ServerStream.Receive(msg); err != nil {
		return err
	}
	request, ok := msg.(proto.Message)
	if !ok {
		return connect.Errorf(connect.CodeInternal, "procedure %s received a %T, not a proto message", s.method.FullName(), msg)
	}
	return s.boundary.check(request, proto.HasExtension(s.method.Options(), casav1.E_Command))
}

func (b boundary) check(request proto.Message, command bool) error {
	descriptor := request.ProtoReflect().Descriptor()
	if err := protovalidate.Validate(request, protovalidate.WithFilter(requestContextOnly(descriptor))); err != nil {
		return validationError(err)
	}
	requestContext, err := readRequestContext(request)
	if err != nil {
		return connect.NewError(connect.CodeInternal, "request context unreadable").WithCause(err)
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
	if command && requestContext.GetIdempotencyKey() == "" {
		return missingIdempotencyKey(descriptor)
	}
	if err := protovalidate.Validate(request); err != nil {
		return validationError(err)
	}
	return nil
}

// requestContextOnly limits validation to request_context and what it
// contains: the request's own other fields, oneofs, and message rules wait for
// step 5.
func requestContextOnly(request protoreflect.MessageDescriptor) protovalidate.Filter {
	return protovalidate.FilterFunc(func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
		if d.FullName() == request.FullName() {
			return false
		}
		if parent := d.Parent(); parent != nil && parent.FullName() == request.FullName() {
			return d.Name() == requestContextField
		}
		return true
	})
}

// readRequestContext copies request_context through its wire form, so a
// generated request and a dynamic one take the same path.
func readRequestContext(request proto.Message) (*casav1.RequestContext, error) {
	message := request.ProtoReflect()
	field := message.Descriptor().Fields().ByName(requestContextField)
	want := (&casav1.RequestContext{}).ProtoReflect().Descriptor().FullName()
	if field == nil || field.Message() == nil || field.Message().FullName() != want {
		return nil, fmt.Errorf("%s has no %s field of type %s", message.Descriptor().FullName(), requestContextField, want)
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
