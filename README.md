# Casa

Casa is a household app for a husband and wife: they know the state of their household money and household matters without having to ask each other, at the lowest recording effort. The planning documents, in Indonesian, start at [`docs/README.md`](docs/README.md); how the project is run, and by whom, is in [`CLAUDE.md`](CLAUDE.md). Code follows [`ENGINEERING-STANDARDS.md`](ENGINEERING-STANDARDS.md).

## Layout

| Folder | Contents |
|---|---|
| `server/` | Go server (`casa` binary) |
| `proto/` | API contract (Protobuf) |
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
make check-contract
```

`check-server` needs the test database (below) running. It runs every server gate that CI runs: format, lint, vulnerability scan, tests, duplication check, and the secrets check over the whole history of the checked-out branch (every commit reachable from `HEAD`). `check-web` installs the web dependencies from the lockfile, then runs every web gate: format, lint, type check, vulnerability scan, duplication check, the client bundle check, the Vitest unit tests (`make web-unit` runs them alone), and the Playwright tests in Chromium and WebKit. `check-contract` runs the contract gates (below). `make check` runs all three.

## Database

The server tests run against a real PostgreSQL 18 in Docker, the image pinned in the `Makefile`, never a mock:

```sh
make db-up        # start it on 127.0.0.1:5432 (CASA_DB_PORT=... to use another port)
make server-test  # each test creates and drops its own database
make db-down
```

`server-test` reads the cluster from `CASA_TEST_ADMIN_URL`, by default `postgres://postgres@127.0.0.1:5432/postgres?sslmode=disable`; without a reachable cluster the database tests fail, they are never skipped. On a real cluster, `server/internal/schema/roles.sql` is run once by a superuser to create the database roles, and `casa migrate` then applies the migrations as `migrator`.

## Contract

`proto/` is the one source of the API contract. The generated Go (`server/internal/gen/`) and TypeScript (`web/src/lib/gen/`) are committed and never edited by hand. After changing a `.proto` file:

```sh
make web-install      # once: protoc-gen-es comes from the web dev dependencies
make proto-generate   # regenerate both, then commit the result
make check-contract   # buf format, buf lint, buf breaking against origin/main, generated code up to date
```

`proto-breaking` compares with `refs/remotes/origin/main` (`BUF_AGAINST_REF=...` to use another ref), so fetch it first.

## Queries

All production SQL is in each module's `queries.sql` (today `server/internal/identity/queries.sql`) and goes through sqlc; the generated Go (`server/internal/*/*db/`) is committed and never edited by hand. `server/sqlc.yaml` names every migration. After changing a query or a migration:

```sh
make sqlc-generate    # regenerate, then commit the result
make sqlc-generated   # fails if the committed code differs (a check-server and CI gate)
```

## Server

```sh
cd server && go run ./cmd/casa serve -listen 127.0.0.1:8080
```

`GET /healthz` answers `ok`. `-minimum-web-version MAJOR.MINOR.PATCH` sets the lowest web client version allowed to send a write command (default in `server/internal/api/version.go`).

A setup token opens the first account of a household (closed sign-up, SEC-158 (b)). The operator issues one per named person, bound to that person's email and valid 7 days, and records the issuance in the operations log:

```sh
casa setup-token -email person@example.com
```

It connects as `app_rw` with the connection string in the systemd credential `app-database-url` (`$CREDENTIALS_DIRECTORY/app-database-url`) and prints the token once; the database keeps only its SHA-256.
