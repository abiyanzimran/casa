# Casa

Casa is a household app for a husband and wife: they know the state of their household money and household matters without having to ask each other, at the lowest recording effort. The planning documents, in Indonesian, start at [`docs/README.md`](docs/README.md); how the project is run, and by whom, is in [`CLAUDE.md`](CLAUDE.md). Code follows [`ENGINEERING-STANDARDS.md`](ENGINEERING-STANDARDS.md).

## Layout

| Folder | Contents |
|---|---|
| `server/` | Go server (`casa` binary) |
| `proto/` | API contract (Protobuf); from task 1.3 |
| `web/` | Web client (PWA); from task 1.1b |
| `golden/` | Golden test vectors for the client input aids; later in FS-1 step 1 |
| `scripts/` | Helper scripts used by the `Makefile` |
| `docs/` | Planning documents and ADRs |

## Local prerequisites

- Go: the version in `server/go.mod` (an older `go` downloads it automatically).
- Java (any current LTS) to run PMD CPD for the duplication check.
- Node.js: the version in `web/.nvmrc` (`nvm install` in `web/`).
- pnpm: the version in `"packageManager"` of `web/package.json` (a newer `pnpm` switches to it automatically).
- Playwright's Chromium and WebKit, once: `cd web && pnpm exec playwright install --with-deps chromium webkit`.
- Docker, for PostgreSQL 18 in the server tests.
- `make`, `curl`, `unzip`, `sha256sum` (linux-amd64).

## Checks

```sh
make check-server
make check-web
```

`check-server` needs the test database (below) running. It runs every server gate that CI runs: format, lint, vulnerability scan, tests, duplication check, and the secrets check over the whole git history. `check-web` installs the web dependencies from the lockfile, then runs every web gate: format, lint, type check, vulnerability scan, duplication check, the client bundle check, and the Playwright tests in Chromium and WebKit. `make check` runs both.

## Database

The server tests run against a real PostgreSQL 18 in Docker, the image pinned in the `Makefile`, never a mock:

```sh
make db-up        # start it on 127.0.0.1:5432 (CASA_DB_PORT=... to use another port)
make server-test  # each test creates and drops its own database
make db-down
```

`server-test` reads the cluster from `CASA_TEST_ADMIN_URL`, by default `postgres://postgres@127.0.0.1:5432/postgres?sslmode=disable`; without a reachable cluster the database tests fail, they are never skipped. On a real cluster, `server/internal/schema/roles.sql` is run once by a superuser to create the database roles, and `casa migrate` then applies the migrations as `migrator`.
