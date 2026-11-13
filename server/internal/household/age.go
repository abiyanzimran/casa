package household

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/database"
	"github.com/abiyanzimran/casa/server/internal/household/householddb"
)

// currentAgeDeclarationVersion is the version of the adult age statement
// the clients show; its text, "Saya berusia 18 tahun atau lebih.", is
// owned by 08 §7.1 (SEC-80).
const currentAgeDeclarationVersion = 1

// CheckAgeDeclaration is decision 9: record is the version to write, 0 for
// none. An account with an adult_age row of any version is not asked again
// (record 0); without one, only the current version is accepted, else
// ErrAgeDeclarationRequired, before any write (SEC-80 note of 14 November
// 2026).
func (s *Service) CheckAgeDeclaration(ctx context.Context, account pgtype.UUID, version uint32) (record uint32, err error) {
	var declared bool
	err = database.Read(ctx, s.pool, database.Scope{AccountID: account}, func(tx pgx.Tx) error {
		declared, err = householddb.New(tx).HasAgeDeclaration(ctx, account)
		return err
	})
	switch {
	case err != nil:
		return 0, err
	case declared:
		return 0, nil
	case version != currentAgeDeclarationVersion:
		return 0, ErrAgeDeclarationRequired
	}
	return version, nil
}
