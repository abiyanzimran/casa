// Package identity holds accounts, credentials, sessions, rate limits and
// closed sign-up (09.2 §1; SEC-11 to SEC-27, SEC-158). It does not import
// the contract: the API layer maps its errors to contract codes (09.2 §1).
package identity

import (
	"context"
	"errors"
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
	// ErrWebSessionNotYet: a web session token travels only in a cookie
	// (SEC-161 (b)), which comes with 1.5 (lead ruling Q13 on the 1.4 plan).
	ErrWebSessionNotYet = errors.New("web sessions come with the cookie transport")
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
)

// Lifetimes of SEC-15 and 09.4 §2, and of the native session (lead ruling Q5
// on the 1.4 plan [Usulan]: the same as a personal web session).
const (
	accessTokenLifetime      = 15 * time.Minute
	nativeIdleTimeout        = 30 * 24 * time.Hour
	nativeAbsoluteLifetime   = 180 * 24 * time.Hour
	verificationLinkLifetime = 30 * time.Minute // lead ruling Q4 on the 1.4 plan [Usulan]
)

// Thresholds of 10.1 §8.1: every token use (StartSignUp, CompleteSignUp) by
// address; SignIn by email and by address.
const (
	tokenUseThreshold = 20
	signInThreshold   = 20
)

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
}

// MessageKind is the kind of an emailed message.
type MessageKind int

const (
	// MessageVerification carries the link that finishes a sign-up.
	MessageVerification MessageKind = iota + 1
)

// Message is one email: a recipient and a one-time link (09.4 §2 "Pengiriman
// email": no household data). The link's token is in its fragment, which
// browsers do not send to a server or put in logs.
type Message struct {
	To   string
	Kind MessageKind
	Link string
}

// Client is who calls: the kind and version from request_context and the
// network peer address "host:port".
type Client struct {
	Kind    string // "ios", "android" or "web", as session.client_kind
	Version string
	Address string
}

// Tokens are a native session's tokens (SEC-15).
type Tokens struct {
	AccessToken, RefreshToken string
	AccessTokenExpiresAt      time.Time
}

// Service is the identity module.
type Service struct {
	pool     *pgxpool.Pool
	config   Config
	hasher   *hasher
	tokenUse rateLimit
	signIn   rateLimit
	// decoy is verified when the email has no credential, so a sign-in
	// costs one hash whatever the reason it fails (SEC-12).
	decoy passwordHash
}

// New returns the identity service over pool.
func New(pool *pgxpool.Pool, config Config) *Service {
	return &Service{
		pool:     pool,
		config:   config,
		hasher:   newHasher(config.Password),
		tokenUse: newRateLimit(0, tokenUseThreshold, false, config.Now),
		signIn:   newRateLimit(signInThreshold, signInThreshold, true, config.Now),
		decoy: passwordHash{
			params: config.Password,
			salt:   randomBytes(int(config.Password.SaltBytes)),
			hash:   randomBytes(int(config.Password.KeyBytes)),
		},
	}
}
