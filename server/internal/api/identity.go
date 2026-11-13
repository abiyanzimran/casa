package api

import (
	"context"
	"errors"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/abiyanzimran/casa/server/internal/command"
	"github.com/abiyanzimran/casa/server/internal/database"
	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/gen/casa/v1/casav1connect"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

// RegisterIdentityService serves casa.v1.IdentityService with s; its write
// commands run on pool through command.Run.
func RegisterIdentityService(server *connect.Server, pool *pgxpool.Pool, s *identity.Service) {
	casav1connect.RegisterIdentityServiceHandler(server, identityHandler{pool: pool, s: s})
}

type identityHandler struct {
	pool *pgxpool.Pool
	s    *identity.Service
}

func (h identityHandler) StartSignUp(ctx context.Context, req *casav1.StartSignUpRequest) (*casav1.StartSignUpResponse, error) {
	if err := h.s.StartSignUp(ctx, clientOf(ctx, req.GetRequestContext()), req.GetSetupToken(), req.GetEmail()); err != nil {
		return nil, identityError(err)
	}
	return &casav1.StartSignUpResponse{}, nil
}

func (h identityHandler) CompleteSignUp(ctx context.Context, req *casav1.CompleteSignUpRequest) (*casav1.CompleteSignUpResponse, error) {
	tokens, err := h.s.CompleteSignUp(ctx, clientOf(ctx, req.GetRequestContext()), req.GetVerificationToken(), req.GetPassword())
	if err != nil {
		return nil, identityError(err)
	}
	return &casav1.CompleteSignUpResponse{Tokens: sessionTokens(tokens)}, nil
}

func (h identityHandler) SignIn(ctx context.Context, req *casav1.SignInRequest) (*casav1.SignInResponse, error) {
	tokens, err := h.s.SignIn(ctx, clientOf(ctx, req.GetRequestContext()), req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, identityError(err)
	}
	return &casav1.SignInResponse{Tokens: sessionTokens(tokens)}, nil
}

func (h identityHandler) RefreshSession(ctx context.Context, req *casav1.RefreshSessionRequest) (*casav1.RefreshSessionResponse, error) {
	tokens, err := h.s.Refresh(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, identityError(err)
	}
	return &casav1.RefreshSessionResponse{Tokens: sessionTokens(tokens)}, nil
}

func (h identityHandler) SignOut(ctx context.Context, req *casav1.SignOutRequest) (*casav1.SignOutResponse, error) {
	session, _ := SessionFrom(ctx) // not public: the boundary authenticated it
	res, err := command.Run(ctx, h.pool, database.Scope{AccountID: session.AccountID},
		casav1connect.IdentityServiceSignOutProcedure, req.GetRequestContext().GetIdempotencyKey(), req,
		func(ctx context.Context, tx pgx.Tx) (*casav1.SignOutResponse, error) {
			return &casav1.SignOutResponse{}, h.s.SignOut(ctx, tx, session)
		})
	if err != nil {
		return nil, identityError(err)
	}
	return res, nil
}

// reauthPurposes maps the contract's purposes to identity's.
var reauthPurposes = map[casav1.ReauthPurpose]string{
	casav1.ReauthPurpose_REAUTH_PURPOSE_CHANGE_PASSWORD: identity.ReauthChangePassword,
}

func (h identityHandler) Reauthenticate(ctx context.Context, req *casav1.ReauthenticateRequest) (*casav1.ReauthenticateResponse, error) {
	purpose, ok := reauthPurposes[req.GetPurpose()]
	if !ok { // unreachable: the boundary allows only defined, non-zero purposes
		return nil, connect.NewError(connect.CodeInternal, "internal error")
	}
	session, _ := SessionFrom(ctx) // not public: the boundary authenticated it
	expiresAt, err := h.s.Reauthenticate(ctx, clientOf(ctx, req.GetRequestContext()), session, req.GetPassword(), purpose)
	if err != nil {
		return nil, identityError(err)
	}
	return &casav1.ReauthenticateResponse{ExpiresAt: timestamppb.New(expiresAt)}, nil
}

func (h identityHandler) ChangePassword(ctx context.Context, req *casav1.ChangePasswordRequest) (*casav1.ChangePasswordResponse, error) {
	session, _ := SessionFrom(ctx)
	scope, key := database.Scope{AccountID: session.AccountID}, req.GetRequestContext().GetIdempotencyKey()
	// A resend gets the first response before the grant check, which the
	// first send used up (ADR-005).
	res, found, err := command.Stored[*casav1.ChangePasswordResponse](ctx, h.pool, scope,
		casav1connect.IdentityServiceChangePasswordProcedure, key, req)
	if err != nil || found {
		return res, identityErrorOrNil(err)
	}
	// Hashed before the command transaction, which may run more than once;
	// only with an open grant (lead ruling on PR B).
	hash, err := h.s.HashNewPassword(ctx, session, req.GetNewPassword())
	if err != nil {
		return nil, newPasswordError(err, req)
	}
	res, err = command.Run(ctx, h.pool, scope, casav1connect.IdentityServiceChangePasswordProcedure, key, req,
		func(ctx context.Context, tx pgx.Tx) (*casav1.ChangePasswordResponse, error) {
			return &casav1.ChangePasswordResponse{}, h.s.ChangePassword(ctx, tx, session, hash)
		})
	if err != nil {
		return nil, identityError(err)
	}
	return res, nil
}

// identityErrorOrNil is identityError, and nil for nil.
func identityErrorOrNil(err error) error {
	if err == nil {
		return nil
	}
	return identityError(err)
}

// newPasswordError is identityError, with a common password reported on
// the request's new_password field.
func newPasswordError(err error, request proto.Message) error {
	if errors.Is(err, identity.ErrPasswordTooCommon) {
		return passwordTooCommon(request.ProtoReflect().Descriptor().Fields().ByName("new_password"))
	}
	return identityError(err)
}

func (h identityHandler) StartRecovery(ctx context.Context, req *casav1.StartRecoveryRequest) (*casav1.StartRecoveryResponse, error) {
	if err := h.s.StartRecovery(ctx, clientOf(ctx, req.GetRequestContext()), req.GetEmail()); err != nil {
		return nil, identityError(err)
	}
	return &casav1.StartRecoveryResponse{}, nil
}

func (h identityHandler) CompleteRecovery(ctx context.Context, req *casav1.CompleteRecoveryRequest) (*casav1.CompleteRecoveryResponse, error) {
	tokens, err := h.s.CompleteRecovery(ctx, clientOf(ctx, req.GetRequestContext()), req.GetRecoveryToken(), req.GetNewPassword())
	if err != nil {
		return nil, newPasswordError(err, req)
	}
	return &casav1.CompleteRecoveryResponse{Tokens: sessionTokens(tokens)}, nil
}

func sessionTokens(t identity.Tokens) *casav1.SessionTokens {
	return &casav1.SessionTokens{
		AccessToken: t.AccessToken, RefreshToken: t.RefreshToken,
		AccessTokenExpiresAt: timestamppb.New(t.AccessTokenExpiresAt),
	}
}

// clientOf builds the caller from request_context, already checked at the
// boundary, and the connection's peer address (10.1 §8.1: TLS ends in the
// binary, so the peer is the client).
func clientOf(ctx context.Context, rc *casav1.RequestContext) identity.Client {
	client := identity.Client{
		Kind:    strings.ToLower(strings.TrimPrefix(rc.GetClientKind().String(), "CLIENT_KIND_")),
		Version: rc.GetClientVersion(),
	}
	if info, ok := connect.CallInfoForServerContext(ctx); ok {
		client.Address = info.PeerAddr
	}
	return client
}

// identityError maps an identity error to the contract (lead ruling Q6 on
// the 1.4 plan; ADR-013 item 4). Each code has one fixed message, so every
// reason behind it gives a byte-identical response (SEC-30); any other error
// is internal and its text never reaches the client (09.2 §3 layer 3).
func identityError(err error) error {
	switch {
	case errors.Is(err, identity.ErrTokenRejected):
		return errorWithDetail(connect.CodePermissionDenied, "token rejected",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED})
	case errors.Is(err, identity.ErrRateLimited):
		return errorWithDetail(connect.CodeResourceExhausted, "rate limited",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_RATE_LIMITED})
	case errors.Is(err, identity.ErrWebSessionNotYet):
		return connect.NewError(connect.CodeUnimplemented, "web sessions come with the cookie transport")
	case errors.Is(err, identity.ErrPasswordTooCommon):
		return passwordTooCommon((&casav1.CompleteSignUpRequest{}).ProtoReflect().Descriptor().Fields().ByName("password"))
	case errors.Is(err, identity.ErrCredentialsRejected):
		return errorWithDetail(connect.CodeUnauthenticated, "credentials rejected",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_CREDENTIALS_REJECTED})
	case errors.Is(err, identity.ErrUnauthenticated):
		// No detail: the client refreshes, then signs in.
		return connect.NewError(connect.CodeUnauthenticated, "unauthenticated")
	case errors.Is(err, identity.ErrSessionExpired):
		return errorWithDetail(connect.CodeUnauthenticated, "session expired",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_SESSION_EXPIRED})
	case errors.Is(err, identity.ErrSessionRevoked):
		return errorWithDetail(connect.CodeUnauthenticated, "session revoked",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_SESSION_REVOKED})
	case errors.Is(err, identity.ErrReauthRequired):
		return errorWithDetail(connect.CodePermissionDenied, "re-authentication required",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_REAUTH_REQUIRED})
	case errors.Is(err, command.ErrIdempotencyConflict):
		return errorWithDetail(connect.CodeFailedPrecondition, "idempotency conflict",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_IDEMPOTENCY_CONFLICT})
	case errors.Is(err, database.ErrRetryExhausted):
		return errorWithDetail(connect.CodeAborted, "retry exhausted",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_RETRY_EXHAUSTED})
	}
	return connect.NewError(connect.CodeInternal, "internal error").WithCause(err)
}

// passwordTooCommon reports the password field in the same Violations form
// as protovalidate's own rules (SEC-11; lead ruling Q5 on the 1.3 plan).
func passwordTooCommon(field protoreflect.FieldDescriptor) error {
	violation := validate.Violation_builder{
		Field:   validate.FieldPath_builder{Elements: []*validate.FieldPathElement{fieldPathElement(field)}}.Build(),
		RuleId:  proto.String("password.common"),
		Message: proto.String("password is on the common password list"),
	}.Build()
	return errorWithDetail(connect.CodeInvalidArgument, "invalid request",
		validate.Violations_builder{Violations: []*validate.Violation{violation}}.Build())
}
