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

`check-server` needs the test database (below) running. It runs every server gate that CI runs: format, lint, vulnerability scan, tests, the old-binary test (below), duplication check, and the secrets check over the whole history of the checked-out branch (every commit reachable from `HEAD`). `check-web` installs the web dependencies from the lockfile, then runs every web gate: format, lint, type check, vulnerability scan, duplication check, the client bundle check, the Vitest unit tests (`make web-unit` runs them alone), and the Playwright tests in Chromium and WebKit. `check-contract` runs the contract gates (below). `make check` runs all three.

## Database

The server tests run against a real PostgreSQL 18 in Docker, the image pinned in the `Makefile`, never a mock:

```sh
make db-up        # start it on 127.0.0.1:5432 (CASA_DB_PORT=... to use another port)
make server-test  # each test creates and drops its own database
make db-down
```

`make server-test-old-binary` (QG-05) builds this branch's `casa` binary and runs the tests of `refs/remotes/origin/main` (`BASE_REF=...` to use another ref; fetch it first) on databases that binary migrated, so a rollback to the old binary keeps working on the new schema (expand and contract).

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

All production SQL is in each module's `queries.sql` (today `server/internal/identity/queries.sql`, `server/internal/household/queries.sql` and `server/internal/command/queries.sql`) and goes through sqlc; the generated Go (`server/internal/*/*db/`) is committed and never edited by hand. `server/sqlc.yaml` names every migration. After changing a query or a migration:

```sh
make sqlc-generate    # regenerate, then commit the result
make sqlc-generated   # fails if the committed code differs (a check-server and CI gate)
```

## Server

```sh
cd server && go run ./cmd/casa serve -mail-dir /path/to/mail -listen 127.0.0.1:8080
```

`GET /healthz` answers `ok`. `serve` connects as `app_rw` with the credential `app-database-url` (below). `-mail-dir` (required, an existing directory) is the fake mail sender of the developer machine and the temporary test server: each verification or recovery mail becomes one file there, readable by its owner only; a real sender replaces it at FS-1 step 7. `serve` also needs the credential `job-database-url`, the connection string of role `job_runner`, with which its background jobs runner claims due jobs (the jobs themselves run as `app_rw`), and the credential `operator-name`: it fills `{pengelola}` in the operator notice, and changing it issues a new notice version and holds every account until they read it again. `-origin` (default `http://127.0.0.1:8080`) is the base of the emailed links. `-minimum-web-version MAJOR.MINOR.PATCH` sets the lowest web client version allowed to send a write command (default in `server/internal/api/version.go`).

A setup token opens the first account of a household (closed sign-up, SEC-158 (b)); everyone else joins through an invitation link (below). The operator issues one per named person, bound to that person's email and valid 7 days, and records the issuance in the operations log:

```sh
casa setup-token -email person@example.com
```

It connects as `app_rw` with the connection string in the systemd credential `app-database-url` (`$CREDENTIALS_DIRECTORY/app-database-url`) and prints the token once; the database keeps only its SHA-256.

### Test server runbook

- **A stopped job.** A job is retried with a growing wait up to 10 attempts, then stops: `serve` logs one JSON line at level `ERROR` with the message `job stopped`, the job's `kind`, `id` and failure `code`. A stopped `household.expire_invitation` job is fixed by the operator before the household makes its next invitation: until then the invitation's expiry is written only when someone acts on it, which can reveal when they acted (10.1 §10 note 1.7 Q3, condition 3). Find the stopped jobs as a superuser with `SELECT id, kind, unique_key, last_error FROM job WHERE locked_until = 'infinity';`. For `household.expire_invitation`, `unique_key` is the invitation's id: check that invitation with `SELECT state, expires_at FROM invitation WHERE id = '<unique_key>';`. An invitation still `open` or `accepted` after its `expires_at` is the one whose expiry was not written. `last_error` is the failure code (`handler`, `panic`, `timeout`, `unknown_kind`, or `lease` when the runner died during the last attempt); find the cause from it and the code, then let the job run again by clearing its lock as a superuser: `UPDATE job SET attempts = 0, locked_until = NULL, last_error = NULL WHERE id = '<id>';`.
- **The jobs runner lost the database.** `serve` prints `casa serve: the jobs runner stopped on a database error` and exits 1, so its supervisor restarts it (systemd `Restart=`, set up at step 7); the error's text is not printed (SEC-44). The PostgreSQL log holds the cause.

## Identity

Sign-up is closed: the person opens the emailed verification link and sets their own password. A native client then signs in with email and password and gets an access token (15 minutes, sent as `Authorization: Bearer`) and a refresh token that rotates on every use. `SignOut` revokes the session. `Reauthenticate` checks the password again and allows one `ChangePassword` within 5 minutes. `StartRecovery` mails a single-use recovery link (30 minutes); `CompleteRecovery` sets a new password with it and revokes every other session. Sign-in, re-authentication, recovery and token use are rate limited per account or email and per network address, in memory; every `StartSignUp` counts on its address, and a sign-up through an invitation also counts on that invitation (5 starts). Web sessions arrive with the cookie transport in FS-1 1.5.

## Household

`HouseholdService` serves the household: the operator notice (`GetOperatorNotice`, `AcknowledgeOperatorNotice`), `CreateHousehold`, `ListMemberships`, and the invitation commands. Until an account has acknowledged the latest operator notice, its household write commands (except acknowledging the notice and withdrawing its own acceptance) and `IssueInvitationLink` answer `NOTICE_UNREAD`. The notice names the operator from the `operator-name` credential of `casa serve` (above). A request's `household_id` only selects a household; the server checks the caller's membership on every call.

Onboarding, in order:

1. The first person signs up with a setup token: `StartSignUp{setup_token}`, the emailed verification link, then `CompleteSignUp` with their own password.
2. They read and acknowledge the operator notice, then `CreateHousehold` with the adult age statement.
3. They invite the second adult: `CreateInvitation`, then `IssueInvitationLink`, and share the link. A new link replaces the old one; the invitation lapses after 7 days.
4. The invitee signs up with the link's token: `StartSignUp{invitation_token}` with any email, the emailed verification link, then `CompleteSignUp`. An account made this way can join a household but never create one.
5. The invitee acknowledges the operator notice, then `PreviewInvitation{from_signup}` and `AcceptInvitation{from_signup}`, without pasting the link again. An account that already exists pastes the link instead (`invitation_token`).
6. The inviter sees the invitee's verified email in `GetOpenInvitation` and confirms with `ConfirmInvitee` (or refuses with `RejectInvitee`); only then is the invitee a member. A household holds two full adult members.

Every invalid, used, replaced, cancelled or lapsed token gets the same `TOKEN_REJECTED` answer.

