package api_test

import (
	"fmt"
	"strings"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/emptypb" // registers google/protobuf/empty.proto for probeFile

	// Every package of the contract is imported here so that its files are
	// registered for casaFiles; a new proto package adds its import.
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
)

// probeMethod describes one method of the synthetic ProbeService the tests
// build, since the contract has no product RPC yet.
type probeMethod struct {
	name            string
	command         *casav1.CommandOptions // nil: not a write command
	clientStreaming bool
	bareRequest     bool // request without request_context
}

// probeFile builds casa/apitest/v1/probe.proto from descriptors: a request
// that carries request_context like every real RPC request, plus a note field
// limited to 3 characters so a test can make the body invalid.
func probeFile(t testing.TB, methods ...probeMethod) protoreflect.FileDescriptor {
	t.Helper()
	required := &descriptorpb.FieldOptions{}
	proto.SetExtension(required, validate.E_Field, validate.FieldRules_builder{Required: proto.Bool(true)}.Build())
	short := &descriptorpb.FieldOptions{}
	proto.SetExtension(short, validate.E_Field, validate.FieldRules_builder{
		String: validate.StringRules_builder{MaxLen: proto.Uint64(3)}.Build(),
	}.Build())
	optional := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
	service := &descriptorpb.ServiceDescriptorProto{Name: proto.String("ProbeService")}
	for _, m := range methods {
		input := ".casa.apitest.v1.ProbeRequest"
		if m.bareRequest {
			input = ".casa.apitest.v1.BareRequest"
		}
		options := &descriptorpb.MethodOptions{}
		if m.command != nil {
			proto.SetExtension(options, casav1.E_Command, m.command)
		}
		service.Method = append(service.Method, &descriptorpb.MethodDescriptorProto{
			Name:            proto.String(m.name),
			InputType:       proto.String(input),
			OutputType:      proto.String(".google.protobuf.Empty"),
			ClientStreaming: proto.Bool(m.clientStreaming),
			Options:         options,
		})
	}
	file, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    proto.String("casa/apitest/v1/probe.proto"),
		Package: proto.String("casa.apitest.v1"),
		Syntax:  proto.String("proto3"),
		Dependency: []string{
			"buf/validate/validate.proto",
			"casa/v1/command.proto",
			"casa/v1/request.proto",
			"google/protobuf/empty.proto",
		},
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("ProbeRequest"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name: proto.String("request_context"), JsonName: proto.String("requestContext"), Number: proto.Int32(1), Label: optional,
						Type: descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(), TypeName: proto.String(".casa.v1.RequestContext"),
						Options: required,
					},
					{
						Name: proto.String("note"), JsonName: proto.String("note"), Number: proto.Int32(2), Label: optional,
						Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(), Options: short,
					},
				},
			},
			{
				Name: proto.String("BareRequest"),
				Field: []*descriptorpb.FieldDescriptorProto{{
					Name: proto.String("note"), JsonName: proto.String("note"), Number: proto.Int32(1), Label: optional,
					Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				}},
			},
		},
		Service: []*descriptorpb.ServiceDescriptorProto{service},
	}, protoregistry.GlobalFiles)
	if err != nil {
		t.Fatalf("probe file: %v", err)
	}
	return file
}

// casaFiles returns every registered file of a casa.* proto package.
func casaFiles(t testing.TB) []protoreflect.FileDescriptor {
	t.Helper()
	var files []protoreflect.FileDescriptor
	protoregistry.GlobalFiles.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		if strings.HasPrefix(string(file.Package()), "casa.") {
			files = append(files, file)
		}
		return true
	})
	return files
}

// contractViolations returns one line per broken contract rule, as
// "<service>.<method>: <rule>" (14 §4.6; 09.3 §1; 09.1 §4).
func contractViolations(files []protoreflect.FileDescriptor) []string {
	var violations []string
	for _, file := range files {
		services := file.Services()
		for i := range services.Len() {
			methods := services.Get(i).Methods()
			for j := range methods.Len() {
				violations = append(violations, methodViolations(methods.Get(j))...)
			}
		}
	}
	return violations
}

func methodViolations(method protoreflect.MethodDescriptor) []string {
	var rules []string
	// 09.1 §4; ADR-013 web note item 5: no client streaming, so every RPC
	// works over HTTP/1.1 from the browser.
	if method.IsStreamingClient() {
		rules = append(rules, "client streaming is not allowed")
	}
	// 14 §4.6; ADR-030 item 9: every request carries the request context the
	// boundary checks read.
	if !startsWithRequestContext(method.Input()) {
		rules = append(rules, "request field 1 is not a required casa.v1.RequestContext named request_context")
	}
	if proto.HasExtension(method.Options(), casav1.E_Command) {
		rules = append(rules, commandViolations(proto.GetExtension(method.Options(), casav1.E_Command).(*casav1.CommandOptions))...)
	}
	violations := make([]string, len(rules))
	for i, rule := range rules {
		violations[i] = fmt.Sprintf("%s: %s", method.FullName(), rule)
	}
	return violations
}

func startsWithRequestContext(request protoreflect.MessageDescriptor) bool {
	field := request.Fields().ByNumber(1)
	if field == nil || field.Name() != "request_context" || field.Message() == nil ||
		field.Message().FullName() != (&casav1.RequestContext{}).ProtoReflect().Descriptor().FullName() {
		return false
	}
	rules, _ := proto.GetExtension(field.Options(), validate.E_Field).(*validate.FieldRules)
	return rules.GetRequired()
}

// commandViolations checks one row of the command kind table (09.3 §1;
// 14 §4.6; ADR-028 item 3).
func commandViolations(command *casav1.CommandOptions) []string {
	var rules []string
	if command.HoldSeconds == nil {
		rules = append(rules, "command without hold_seconds")
	}
	// SEC-151: a protective command is never held.
	if command.GetProtective() && command.GetHoldSeconds() != 0 {
		rules = append(rules, fmt.Sprintf("protective with hold_seconds %d", command.GetHoldSeconds()))
	}
	// SEC-34; 09.3 §1: protective ⊆ notice_exempt.
	if command.GetProtective() && !command.GetNoticeExempt() {
		rules = append(rules, "protective but not notice_exempt")
	}
	return rules
}
