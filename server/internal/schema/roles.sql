-- Database roles of 09.2 §3 layer 5 and ADR-016 item 4 (SEC-09).
--
-- Run once per PostgreSQL cluster by a superuser, before the first migration:
--
--   psql -v ON_ERROR_STOP=1 -f roles.sql
--   CREATE DATABASE casa OWNER migrator;
--
-- Migration 00001 then sets the database and schema privileges as the
-- database owner. Roles are cluster objects, shared by every database of the
-- cluster, so they are not a migration: a migration runs as migrator, and
-- migrator has no CREATEROLE (SEC-09). How each role authenticates (pg_hba,
-- password through systemd LoadCredential=, or peer over the Unix socket) is
-- set up by the operations runbook (14 §3.3 step 7), so no password is here.
--
-- Running this file again is safe and also resets the attributes. It holds no
-- psql meta-command and no CREATE DATABASE, so it runs in one transaction.
--
-- migrator: owns the schema and every table; used only by casa migrate.
-- app_rw: the server's role; owns nothing, so row-level security applies.
-- job_runner: claims job rows across households; the job itself runs as app_rw.
-- lifecycle_rw: the only role with DELETE on data that ADR-007 deletes.
-- The backup role comes with pgBackRest at step 7 (09.4 §6).
DO $$
DECLARE
	role_name text;
BEGIN
	FOREACH role_name IN ARRAY ARRAY['migrator', 'app_rw', 'job_runner', 'lifecycle_rw'] LOOP
		IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = role_name) THEN
			EXECUTE format('CREATE ROLE %I', role_name);
		END IF;
		-- Not superuser and no BYPASSRLS: either would skip row-level
		-- security (SEC-08, SEC-09).
		EXECUTE format('ALTER ROLE %I LOGIN NOSUPERUSER NOCREATEDB NOCREATEROLE NOREPLICATION NOBYPASSRLS',
			role_name);
	END LOOP;
END
$$;
