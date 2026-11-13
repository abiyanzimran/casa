// Package identity holds accounts, credentials, sessions, rate limits and
// closed sign-up (09.2 §1; SEC-11 to SEC-27, SEC-158). It does not import
// the contract: the API layer maps its errors to contract codes (09.2 §1).
package identity

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Errors a caller maps to a response; each stands for every reason behind it,
// so a response never tells the reasons apart (SEC-30).
var (
	// ErrTokenRejected: a setup or verification token is missing,
	// malformed, unknown, expired, used, or bound to another email.
	ErrTokenRejected = errors.New("token rejected")
	// ErrRateLimited: too many failed attempts from the account or the
	// network address (SEC-12).
	ErrRateLimited = errors.New("rate limited")
	// ErrCredentialsRejected: the email is unknown or the password is
	// wrong, never told apart (SEC-12).
	ErrCredentialsRejected = errors.New("credentials rejected")
	// ErrUnauthenticated: no access token, an unknown one, or one older
	// than 15 minutes; the client refreshes, then signs in (SEC-15).
	ErrUnauthenticated = errors.New("unauthenticated")
	// ErrSessionExpired: the session passed its idle or absolute limit.
	ErrSessionExpired = errors.New("session expired")
	// ErrSessionRevoked: the session was revoked (SEC-15, SEC-24).
	ErrSessionRevoked = errors.New("session revoked")
	// ErrReauthRequired: the action needs an open re-authentication of this
	// session for its purpose (SEC-13).
	ErrReauthRequired = errors.New("re-authentication required")
)

// Lifetimes of SEC-15 and 09.4 §2, and of the native session (lead ruling Q5
// on the 1.4 plan [Usulan]: the same as a personal web session).
const (
	accessTokenLifetime      = 15 * time.Minute
	nativeIdleTimeout        = 30 * 24 * time.Hour
	nativeAbsoluteLifetime   = 180 * 24 * time.Hour
	verificationLinkLifetime = 30 * time.Minute // lead ruling Q4 on the 1.4 plan [Usulan]
	reauthLifetime           = 5 * time.Minute  // SEC-13 [Asumsi]; 09.4 §2
	recoveryLinkLifetime     = 30 * time.Minute // SEC-24 [Asumsi]; 09.4 §2
)

// Lifetimes of a web session by its device mode (plan 1.5 decision 8; 09.4
// §2).
const (
	personalIdleTimeout      = 30 * 24 * time.Hour  // SEC-161 (d) [Asumsi]
	personalAbsoluteLifetime = 180 * 24 * time.Hour // SEC-161 (d) [Asumsi]
	sharedIdleTimeout        = 30 * time.Minute     // SEC-165 (a) [Asumsi]
	sharedAbsoluteLifetime   = 12 * time.Hour       // SEC-165 (a) [Asumsi]
)

// Thresholds of 10.1 §8.1: every token use (StartSignUp, CompleteSignUp) by
// address; SignIn by email and by address; Reauthenticate by account and by
// address; StartRecovery by email and by address.
const (
	tokenUseThreshold        = 20
	signInThreshold          = 20
	reauthThreshold          = 20
	recoveryEmailThreshold   = 5
	recoveryAddressThreshold = 20
)

// afterResponseTimeout bounds the work after a response, such as a recovery
// mail (10.1 §5.1 rule 6) [Usulan].
const afterResponseTimeout = 30 * time.Second

// touchAfter is how stale last_used_at may get before a request updates it,
// so a burst of requests does not write on each one [Usulan].
const touchAfter = time.Minute

// Config is what the identity service needs from its host.
type Config struct {
	// Origin is the base of emailed links, such as https://casa.example.
	Origin string
	// SendMail delivers a message; it is called only after the transaction
	// that made the message's token has committed.
	SendMail func(context.Context, Message) error
	Now      func() time.Time
	// Password is DefaultPasswordParams in production.
	Password PasswordParams
	// Logger receives the failures of work done after a response; nil is
	// slog.Default().
	Logger *slog.Logger
}

// MessageKind is the kind of an emailed message.
type MessageKind int

const (
	// MessageVerification carries the link that finishes a sign-up.
	MessageVerification MessageKind = iota + 1
	// MessageRecovery carries the link that sets a new password (SEC-24).
	MessageRecovery
)

// Message is one email: a recipient and a one-time link (09.4 §2 "Pengiriman
// email": no household data). The link's token is in its fragment, which
// browsers do not send to a server or put in logs.
type Message struct {
	To   string
	Kind MessageKind
	Link string
}

// Client is who calls: the kind and version from request_context, the
// network peer address "host:port", and for a web client its device mode and
// User-Agent.
type Client struct {
	Kind    string // "ios", "android" or "web", as session.client_kind
	Version string
	Address string
	// DeviceMode is "personal" or "shared" for web and "" for native, whose
	// session is stored as personal (plan 1.5 decision 9; SEC-165).
	DeviceMode string
	// UserAgent is read only by summarizeUserAgent and never stored or
	// logged (SEC-161 (e)).
	UserAgent string
}

// Cookie is a web session's cookie (plan 1.5 decision 7; SEC-161 (b)).
type Cookie struct {
	Value string // the 43-character token; only its hash is stored
	// MaxAge is the session's remaining absolute lifetime; 0 makes a
	// session cookie, which shared mode uses (SEC-165 (a)).
	MaxAge time.Duration
}

// Tokens are what opening a session gives the client: a native session's
// tokens (SEC-15), or a web session's cookie, with the other fields empty
// (SEC-161 (b)).
type Tokens struct {
	AccessToken, RefreshToken string
	AccessTokenExpiresAt      time.Time
	Cookie                    Cookie
}

// Service is the identity module.
type Service struct {
	pool     *pgxpool.Pool
	config   Config
	hasher   *hasher
	tokenUse rateLimit
	signIn   rateLimit
	reauth   rateLimit
	recovery rateLimit
	// background is the work started after a response: recovery mail
	// until the identity.send_mail job of 1.7.
	background sync.WaitGroup
	// afterResponseTimeout bounds each piece of background work.
	afterResponseTimeout time.Duration
	// closed refuses background work once shutdown starts; mu orders it
	// with background.Go, so no Add runs during Close's Wait.
	mu     sync.Mutex
	closed bool
	logger *slog.Logger
	// decoy is verified when the email has no credential, so a sign-in
	// costs one hash whatever the reason it fails (SEC-12).
	decoy PasswordHash
}

// Wait returns when the work started after earlier responses is done,
// without refusing new work; tests call it to see that work. On shutdown
// casa serve calls Close instead.
func (s *Service) Wait() { s.background.Wait() }

// Close refuses new background work and waits for the work already
// started; casa serve calls it on shutdown, before closing the pool.
func (s *Service) Close() {
	s.mu.Lock()
	s.closed = true
	s.mu.Unlock()
	s.background.Wait()
}

// afterResponse runs work after the response, bounded by
// afterResponseTimeout and detached from the request's cancellation (10.1
// §5.1 rule 6); false once shutdown has started.
func (s *Service) afterResponse(ctx context.Context, work func(context.Context)) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return false
	}
	s.background.Go(func() {
		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), s.afterResponseTimeout)
		defer cancel()
		work(ctx)
	})
	return true
}

// New returns the identity service over pool.
func New(pool *pgxpool.Pool, config Config) *Service {
	logger := config.Logger
	if logger == nil {
		logger = slog.Default()
	}
	return &Service{
		afterResponseTimeout: afterResponseTimeout,
		logger:               logger,
		pool:                 pool,
		config:               config,
		hasher:               newHasher(config.Password),
		tokenUse:             newRateLimit(0, tokenUseThreshold, false, config.Now),
		signIn:               newRateLimit(signInThreshold, signInThreshold, true, config.Now),
		reauth:               newRateLimit(reauthThreshold, reauthThreshold, true, config.Now),
		recovery:             newRateLimit(recoveryEmailThreshold, recoveryAddressThreshold, false, config.Now),
		decoy: PasswordHash{
			params: config.Password,
			salt:   randomBytes(int(config.Password.SaltBytes)),
			hash:   randomBytes(int(config.Password.KeyBytes)),
		},
	}
}
