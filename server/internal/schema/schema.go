// Package schema holds the database schema of Casa: the cluster roles and the
// versioned goose migrations (ADR-016 items 4 and 11; 09.2 §3 layer 5, §10).
package schema

import _ "embed"

// Roles creates the database roles of 09.2 §3 layer 5 (SEC-09). It is run by a
// superuser once per cluster, not by a migration; see roles.sql.
//
//go:embed roles.sql
var Roles string
