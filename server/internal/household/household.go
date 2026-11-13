// Package household holds households, memberships, approval authorities,
// household rules, relationships, invitations and their object events (09.2
// §1; 06 §3.1, §3.6). It does not import identity or the contract: the API
// layer composes commands that touch both and maps errors to contract codes
// (09.2 §1 item 9).
package household

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Errors the API layer maps to the contract (ADR-013 note 1.6; plan 1.6
// decision 17).
var (
	// ErrNotFound: not a membership of the account, another household's
	// object, or an actor who is not a party to it (SEC-05).
	ErrNotFound = errors.New("not found")
	// ErrCreationNotAllowed: the account was made through an invitation, or
	// is already a member of a household (SEC-158 (a); FR-FAM-001 AC-5,
	// AC-10).
	ErrCreationNotAllowed = errors.New("household creation not allowed")
	// ErrAgeDeclarationRequired: the account has no adult age statement and
	// the command carries none of the current text version (SEC-80).
	ErrAgeDeclarationRequired = errors.New("age declaration required")
	// ErrAlreadyMember: the account is already a member of a household,
	// caught by membership_one_household when a concurrent command made it
	// one (09.2 §5.1, note on decision 20).
	ErrAlreadyMember = errors.New("already a household member")
)

// Service is the household module.
type Service struct {
	pool *pgxpool.Pool
	now  func() time.Time
}

// New returns the household service over pool, reading time from now.
func New(pool *pgxpool.Pool, now func() time.Time) *Service {
	return &Service{pool: pool, now: now}
}
