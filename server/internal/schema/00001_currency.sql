-- Database and schema privileges, and the product's currency table.
--
-- Adding a currency later needs, in the same migration, a temporary
-- ALTER TABLE currency NO FORCE ROW LEVEL SECURITY (with FORCE even the owner
-- is bound by the policies, and the only policy is SELECT), or an INSERT policy
-- for migrator.

-- +goose Up
-- Nothing in this database for roles not named here (SEC-09; 09.2 §3 layer 5).
REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT USAGE ON SCHEMA public TO app_rw, job_runner, lifecycle_rw;

-- Staging and production databases do not grant each other's roles CONNECT
-- (ADR-028 item 5).
-- +goose StatementBegin
DO $$
BEGIN
	EXECUTE format('REVOKE ALL ON DATABASE %I FROM PUBLIC', current_database());
	EXECUTE format('GRANT CONNECT ON DATABASE %I TO app_rw, job_runner, lifecycle_rw', current_database());
END
$$;
-- +goose StatementEnd

-- 09.2 §5.1. The exponent is the product's, not ISO 4217's (ADR-008).
CREATE TABLE currency (
	code     char(3) PRIMARY KEY, -- ISO 4217
	exponent smallint NOT NULL CHECK (exponent BETWEEN 0 AND 4)
);
-- IDR has exponent 0 in this product (ADR-008; 07 §2.1, INV-39). Inserted
-- before FORCE ROW LEVEL SECURITY, which binds the owner too.
INSERT INTO currency (code, exponent) VALUES ('IDR', 0);

-- Every table forces row-level security (SEC-08; ADR-016 item 3). currency is
-- a global table, read-only for app_rw (09.2 §3 "Tabel global").
ALTER TABLE currency ENABLE ROW LEVEL SECURITY;
ALTER TABLE currency FORCE ROW LEVEL SECURITY;
CREATE POLICY currency_read ON currency FOR SELECT USING (true);
GRANT SELECT ON currency TO app_rw;

-- +goose Down
DROP TABLE currency;

-- Back to the PostgreSQL 18 defaults.
-- +goose StatementBegin
DO $$
BEGIN
	EXECUTE format('REVOKE CONNECT ON DATABASE %I FROM app_rw, job_runner, lifecycle_rw', current_database());
	EXECUTE format('GRANT CONNECT, TEMPORARY ON DATABASE %I TO PUBLIC', current_database());
END
$$;
-- +goose StatementEnd

REVOKE USAGE ON SCHEMA public FROM app_rw, job_runner, lifecycle_rw;
GRANT USAGE ON SCHEMA public TO PUBLIC;
