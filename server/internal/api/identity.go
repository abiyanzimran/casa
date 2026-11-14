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
// commands run on pool through command.Run. invitationOpen is sign-up's
// invitation check (InvitationOpen), passed in because identity and
// household do not import each other (09.2 §1 item 9).
func RegisterIdentityService(server *connect.Server, pool *pgxpool.Pool, s *identity.Service,
	invitationOpen func(context.Context, []byte) error,
) {
	casav1connect.RegisterIdentityServiceHandler(server, identityHandler{pool: pool, s: s, invitationOpen: invitationOpen})
}

type identityHandler struct {
	pool           *pgxpool.Pool
	s              *identity.Service
	invitationOpen func(context.Context, []byte) error
}

func (h identityHandler) StartSignUp(ctx context.Context, req *casav1.StartSignUpRequest) (*casav1.StartSignUpResponse, error) {
	source := identity.SignUpSource{SetupToken: req.GetSetupToken(), InvitationToken: req.GetInvitationToken()}
	if err := h.s.StartSignUp(ctx, clientOf(ctx, req), source, req.GetEmail(), h.invitationOpen); err != nil {
		return nil, identityError(err)
	}
	return &casav1.StartSignUpResponse{}, nil
}

func (h identityHandler) CompleteSignUp(ctx context.Context, req *casav1.CompleteSignUpRequest) (*casav1.CompleteSignUpResponse, error) {
	client := clientOf(ctx, req)
	tokens, err := h.s.CompleteSignUp(ctx, client, req.GetVerificationToken(), req.GetEmail(), req.GetPassword(), h.invitationOpen)
	if err != nil {
		return nil, identityError(err)
	}
	opened, err := openedSession(ctx, client, tokens)
	if err != nil {
		return nil, identityError(err)
	}
	return &casav1.CompleteSignUpResponse{Tokens: opened}, nil
}

func (h identityHandler) SignIn(ctx context.Context, req *casav1.SignInRequest) (*casav1.SignInResponse, error) {
	client := clientOf(ctx, req)
	tokens, err := h.s.SignIn(ctx, client, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, identityError(err)
	}
	opened, err := openedSession(ctx, client, tokens)
	if err != nil {
		return nil, identityError(err)
	}
	return &casav1.SignInResponse{Tokens: opened}, nil
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
	if session.ClientKind == "web" { // SEC-164 (b): only after the revocation committed
		if err := clearSessionCookie(ctx); err != nil {
			return nil, identityError(err)
		}
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
	expiresAt, cookie, err := h.s.Reauthenticate(ctx, clientOf(ctx, req), session, req.GetPassword(), purpose)
	if err != nil {
		return nil, identityError(err)
	}
	if cookie != nil {
		// SEC-161 (c): a web re-authentication always rotates the cookie
		// (session fixation), and its grant serves only this value (note
		// 1.5 rule (7)).
		if err := setSessionCookie(ctx, *cookie); err != nil {
			return nil, identityError(err)
		}
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
	if err := h.s.StartRecovery(ctx, clientOf(ctx, req), req.GetEmail()); err != nil {
		return nil, identityError(err)
	}
	return &casav1.StartRecoveryResponse{}, nil
}

func (h identityHandler) CompleteRecovery(ctx context.Context, req *casav1.CompleteRecoveryRequest) (*casav1.CompleteRecoveryResponse, error) {
	client := clientOf(ctx, req)
	tokens, err := h.s.CompleteRecovery(ctx, client, req.GetRecoveryToken(), req.GetNewPassword())
	if err != nil {
		return nil, newPasswordError(err, req)
	}
	opened, err := openedSession(ctx, client, tokens)
	if err != nil {
		return nil, identityError(err)
	}
	return &casav1.CompleteRecoveryResponse{Tokens: opened}, nil
}

// openedSession hands a new session to its client: a web session only as the
// cookie, with tokens left unset in the response (SEC-161 (b)); a native one
// as tokens. A completion without a session answers with neither and sets no
// Set-Cookie at all, a stale cookie included (SEC-171 (b) (2) note Q2-4;
// ADR-030 item 11). A cookie that cannot be set is an internal error.
func openedSession(ctx context.Context, client identity.Client, tokens identity.Tokens) (*casav1.SessionTokens, error) {
	if client.WithoutSession {
		return nil, nil
	}
	if client.Kind == "web" {
		return nil, setSessionCookie(ctx, tokens.Cookie)
	}
	return sessionTokens(tokens), nil
}

func sessionTokens(t identity.Tokens) *casav1.SessionTokens {
	return &casav1.SessionTokens{
		AccessToken: t.AccessToken, RefreshToken: t.RefreshToken,
		AccessTokenExpiresAt: timestamppb.New(t.AccessTokenExpiresAt),
	}
}

// deviceModes maps the contract's device modes to session.device_mode; the
// unspecified mode, which only native clients send, maps to "".
var deviceModes = map[casav1.DeviceMode]string{
	casav1.DeviceMode_DEVICE_MODE_PERSONAL: "personal",
	casav1.DeviceMode_DEVICE_MODE_SHARED:   "shared",
}

// clientOf builds the caller from request_context and, when the request has
// one, device_mode and without_session, all already checked at the boundary (plan 1.5 decision
// 9), the connection's peer address (10.1 §8.1: TLS ends in the binary, so
// the peer is the client), and the User-Agent, which identity only
// summarizes (SEC-161 (e)).
func clientOf(ctx context.Context, req interface {
	GetRequestContext() *casav1.RequestContext
}) identity.Client {
	rc := req.GetRequestContext()
	client := identity.Client{Kind: clientKindOf(rc), Version: rc.GetClientVersion()}
	if withMode, ok := req.(interface{ GetDeviceMode() casav1.DeviceMode }); ok {
		client.DeviceMode = deviceModes[withMode.GetDeviceMode()]
	}
	if without, ok := req.(interface{ GetWithoutSession() bool }); ok {
		client.WithoutSession = without.GetWithoutSession()
	}
	if info, ok := connect.CallInfoForServerContext(ctx); ok {
		client.Address = info.PeerAddr
		client.UserAgent = info.RequestHeader().Get("User-Agent")
	}
	return client
}

// clientKindOf is request_context.client_kind as session.client_kind
// writes it: "ios", "android" or "web".
func clientKindOf(rc *casav1.RequestContext) string {
	return strings.ToLower(strings.TrimPrefix(rc.GetClientKind().String(), "CLIENT_KIND_"))
}

// identityError maps an identity error to the contract (lead ruling Q6 on
// the 1.4 plan; ADR-013 item 4). Each code has one fixed message, so every
// reason behind it gives a byte-identical response (SEC-30); any other error
// is internal and its text never reaches the client (09.2 §3 layer 3).
func identityError(err error) error {
	switch {
	case errors.Is(err, identity.ErrTokenRejected):
		return tokenRejected()
	case errors.Is(err, identity.ErrRateLimited):
		return errorWithDetail(connect.CodeResourceExhausted, "rate limited",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_RATE_LIMITED})
	case errors.Is(err, identity.ErrEmailMismatch):
		// 10.4 SEC-158 note "halaman verifikasi" item 2; ADR-013 note 1.6.
		return errorWithDetail(connect.CodeFailedPrecondition, "email mismatch",
			&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_EMAIL_MISMATCH})
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
	}
	return commandError(err)
}

// tokenRejected is the one answer to every token that is not accepted,
// whatever the reason, for the identity tokens and the invitation token
// alike, so no two reasons differ by a byte (SEC-30).
func tokenRejected() error {
	return errorWithDetail(connect.CodePermissionDenied, "token rejected",
		&casav1.ErrorDetail{Code: casav1.ErrorCode_ERROR_CODE_TOKEN_REJECTED})
}

// commandError maps the errors of the write-command path every module
// shares; any other error is internal and its text never reaches the client
// (09.2 §3 layer 3).
func commandError(err error) error {
	switch {
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
	return invalidFields(fieldViolation(field, "password.common", "password is on the common password list"))
}

// fieldViolation is a rule the server checks beyond the contract, on one
// field, in the form of protovalidate's own violations.
func fieldViolation(field protoreflect.FieldDescriptor, ruleID, message string) *validate.Violation {
	return validate.Violation_builder{
		Field:   validate.FieldPath_builder{Elements: []*validate.FieldPathElement{fieldPathElement(field)}}.Build(),
		RuleId:  proto.String(ruleID),
		Message: proto.String(message),
	}.Build()
}

// invalidFields is the invalid-argument answer with violations, as the
// boundary's validation gives it (lead ruling Q5 on the 1.3 plan).
func invalidFields(violations ...*validate.Violation) error {
	return errorWithDetail(connect.CodeInvalidArgument, "invalid request",
		validate.Violations_builder{Violations: violations}.Build())
}
