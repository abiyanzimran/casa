// Package household holds households, memberships, approval authorities,
// household rules, relationships, invitations and their object events (09.2
// §1; 06 §3.1, §3.6). It does not import identity or the contract: the API
// layer composes commands that touch both and maps errors to contract codes
// (09.2 §1 item 9).
package household

import (
	"errors"
	"io/fs"
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
	// ErrStateChanged: the object is no longer in the state the command was
	// made for, such as an acknowledgement of a notice version that is not
	// the latest (10.1 §11 note of 14 November 2026 (b)).
	ErrStateChanged = errors.New("state changed")
)

// Service is the household module.
type Service struct {
	pool *pgxpool.Pool
	now  func() time.Time
	// operatorName fills {pengelola} in the operator notice; notices are its
	// versions, PublishedNotices in production (09.3 §1 "Mekanisme").
	operatorName string
	notices      fs.FS
}

// New returns the household service over pool, reading time from now, with
// the server's operator name and the notice versions to show.
func New(pool *pgxpool.Pool, now func() time.Time, operatorName string, notices fs.FS) *Service {
	return &Service{pool: pool, now: now, operatorName: operatorName, notices: notices}
}

// Now is the service's clock, so a security trail line the API layer writes
// for a household command carries the same time as the command's rows.
func (s *Service) Now() time.Time { return s.now() }
