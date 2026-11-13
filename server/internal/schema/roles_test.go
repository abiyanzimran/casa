package schema_test

import (
	"testing"

	"github.com/abiyanzimran/casa/server/internal/database/databasetest"
)

func TestRoles(t *testing.T) {
	d := databasetest.Empty(t)
	db := d.SQL(t, d.AdminRole)
	roles := []string{"migrator", "app_rw", "job_runner", "lifecycle_rw"}

	t.Run("SEC-09 database roles are not superuser and cannot bypass row-level security", func(t *testing.T) {
		for _, role := range roles {
			var super, bypassRLS, createRole, createDB, replication, login bool
			err := db.QueryRowContext(t.Context(),
				`SELECT rolsuper, rolbypassrls, rolcreaterole, rolcreatedb, rolreplication, rolcanlogin
				   FROM pg_roles WHERE rolname = $1`, role).
				Scan(&super, &bypassRLS, &createRole, &createDB, &replication, &login)
			if err != nil {
				t.Fatalf("role %s: %v", role, err)
			}
			if super || bypassRLS || createRole || createDB || replication || !login {
				t.Errorf("role %s: superuser=%t bypassrls=%t createrole=%t createdb=%t replication=%t login=%t; "+
					"want only login (SEC-09; 09.2 §3 layer 5)",
					role, super, bypassRLS, createRole, createDB, replication, login)
			}
		}
	})

	// roles.sql resets attributes, not memberships. A membership would hand
	// over another role's rights: migrator's tables (owner, so no RLS),
	// lifecycle_rw's DELETE, or built-in roles such as pg_read_server_files.
	t.Run("SEC-09 database roles are members of no other role", func(t *testing.T) {
		for _, role := range roles {
			var granted string
			if err := db.QueryRowContext(t.Context(), `
				SELECT coalesce(string_agg(roleid::regrole::text, ', ' ORDER BY 1), '')
				  FROM pg_auth_members WHERE member = $1::regrole`, role).Scan(&granted); err != nil {
				t.Fatalf("role %s: %v", role, err)
			}
			if granted != "" {
				t.Errorf("role %s is a member of %v", role, granted)
			}
		}
	})
}
