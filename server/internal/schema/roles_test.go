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

	t.Run("SEC-09 application roles are not members of migrator", func(t *testing.T) {
		for _, role := range roles[1:] {
			var member bool
			if err := db.QueryRowContext(t.Context(),
				`SELECT pg_has_role($1, 'migrator', 'MEMBER')`, role).Scan(&member); err != nil {
				t.Fatalf("role %s: %v", role, err)
			}
			if member {
				t.Errorf("role %s is a member of migrator, so it owns every table and bypasses row-level security", role)
			}
		}
	})
}
