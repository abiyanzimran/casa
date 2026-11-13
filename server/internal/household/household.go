// Package household holds households, memberships, approval authorities,
// household rules, relationships, invitations and their object events (09.2
// §1; 06 §3.1, §3.6). It does not import identity or the contract: the API
// layer composes commands that touch both and maps errors to contract codes
// (09.2 §1 item 9).
package household

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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
