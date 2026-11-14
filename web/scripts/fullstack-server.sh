#!/usr/bin/env bash
# A throwaway casa stack on this PC for the fullstack browser test and the pilot
# device preview (plan 1.9b/1.10 Task 12): PostgreSQL in the container
# casa-postgres-web, the migrated schema, two setup tokens, and casa serve with
# the web bundle of this worktree. Everything lives until the script ends
# (Ctrl-C or SIGTERM): the container, the credentials and the mail files are
# removed. Test data only.
#
#   web/scripts/fullstack-server.sh
#
# Environment (all optional):
#   CASA_ORIGIN        base of emailed links and the only accepted Origin
#                      (default http://127.0.0.1:8080)
#   CASA_LISTEN        casa's listen address (default 127.0.0.1:8080)
#   CASA_WEB_DB_PORT   loopback port of the container (default 55441; 55432 is
#                      taken by another project, 55440 is casa-postgres of make db-up)
#   CASA_SETUP_EMAILS  one setup token per address, space separated
#   CASA_RUN_DIR       an existing 0700 directory the caller owns; default: a
#                      new mktemp directory this script removes on exit
#
# It prints addresses and file paths, never a token, a password or a mail body.
# Files in the run directory: mail/ (one file per mail), setup-token-N (0600,
# the token only), ready (written when casa answers /healthz).
set -euo pipefail
umask 077

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)
origin=${CASA_ORIGIN:-http://127.0.0.1:8080}
listen=${CASA_LISTEN:-127.0.0.1:8080}
port=${CASA_WEB_DB_PORT:-55441}
emails=${CASA_SETUP_EMAILS:-"a@casa.test b@casa.test"}
container=casa-postgres-web
image=$(sed -n 's/^POSTGRES_IMAGE := //p' "$root/Makefile")
[ -n "$image" ] || { echo "fullstack-server: POSTGRES_IMAGE not found in Makefile" >&2; exit 1; }

own_run=
if [ -z "${CASA_RUN_DIR:-}" ]; then
	CASA_RUN_DIR=$(mktemp -d)
	own_run=1
fi
run=$CASA_RUN_DIR
cred=$run/credentials
mail=$run/mail
casa_pid=

cleanup() {
	trap '' INT TERM
	if [ -n "$casa_pid" ]; then
		kill "$casa_pid" 2>/dev/null || true
		wait "$casa_pid" 2>/dev/null || true
	fi
	docker rm -f "$container" >/dev/null 2>&1 || true
	rm -rf "$cred"
	if [ -n "$own_run" ]; then rm -rf "$run"; fi
}
trap cleanup EXIT
trap 'exit 130' INT
trap 'exit 143' TERM

if ss -ltn "sport = :$port" | grep -q LISTEN && ! docker ps -q -f "name=^$container$" | grep -q .; then
	echo "fullstack-server: port $port is in use by something else; set CASA_WEB_DB_PORT" >&2
	exit 1
fi

# The binary: this worktree's server and the web bundle embedded in it.
(cd "$root" && make casa >/dev/null)
casa=$root/.cache/casa

# The database. Only the container of this name is ever replaced.
docker rm -f "$container" >/dev/null 2>&1 || true
docker image inspect "$image" >/dev/null 2>&1 || docker pull -q "$image" >/dev/null
# trust: bound to 127.0.0.1, data on tmpfs, test data only (as make db-up).
docker run -d --rm --name "$container" -p "127.0.0.1:$port:5432" \
	-e POSTGRES_HOST_AUTH_METHOD=trust --tmpfs /var/lib/postgresql "$image" >/dev/null
ready=
for _ in $(seq 60); do
	if docker exec "$container" pg_isready -U postgres -h 127.0.0.1 >/dev/null 2>&1; then ready=1; break; fi
	sleep 1
done
[ -n "$ready" ] || { echo "fullstack-server: $container not ready after 60 s" >&2; exit 1; }
psql_admin() { docker exec -i "$container" psql -U postgres -h 127.0.0.1 -v ON_ERROR_STOP=1 -q "$@"; }
psql_admin <"$root/server/internal/schema/roles.sql" >/dev/null
psql_admin -c 'CREATE DATABASE casa OWNER migrator' >/dev/null

# Credentials (09.4 §1 LoadCredential= files): a 0700 directory, removed on exit.
mkdir -p -m 700 "$cred" "$mail"
url() { printf 'postgres://%s@127.0.0.1:%s/casa?sslmode=disable\n' "$1" "$port"; }
url migrator >"$cred/migrator-database-url"
url app_rw >"$cred/app-database-url"
url job_runner >"$cred/job-database-url"
echo 'Pengelola Uji' >"$cred/operator-name"
openssl rand -hex 32 >"$cred/sync-cursor-key"
export CREDENTIALS_DIRECTORY=$cred

"$casa" migrate >/dev/null
n=0
for email in $emails; do
	n=$((n + 1))
	# The token is the text after the prefix of the first line; it goes to a file, not the terminal.
	"$casa" setup-token -email "$email" | sed -n 's/^casa setup-token: //p' >"$run/setup-token-$n"
	[ -s "$run/setup-token-$n" ] || { echo "fullstack-server: no setup token for $email" >&2; exit 1; }
	echo "setup token $n for $email: $run/setup-token-$n"
done

"$casa" serve -listen "$listen" -origin "$origin" -mail-dir "$mail" >/dev/null &
casa_pid=$!
for _ in $(seq 50); do
	curl -sf "http://$listen/healthz" >/dev/null 2>&1 && break
	kill -0 "$casa_pid" 2>/dev/null || { echo "fullstack-server: casa serve stopped" >&2; exit 1; }
	sleep 0.2
done
curl -sf "http://$listen/healthz" >/dev/null || { echo "fullstack-server: casa serve not ready" >&2; exit 1; }
touch "$run/ready"
echo "casa serve: http://$listen (origin $origin)"
echo "mail files: $mail"
echo "run directory: $run"
wait "$casa_pid"
