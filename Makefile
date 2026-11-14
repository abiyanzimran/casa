# Single source of the CI gate commands: CI calls these targets and nothing
# else, so a local run and a CI run check the same things (ES §2, §7).
# Gates: QG-01 format and lint, QG-02 duplication, QG-03 vulnerabilities,
# QG-04 contract, QG-05 tests, QG-06 ID traceability
# (docs/14-engineering-roadmap.md §5.1, §4.6, §4.13), plus the repository
# secrets check (09.4 §5).

SHELL := /bin/bash
.SHELLFLAGS := -euo pipefail -c

GOLANGCI_LINT_VERSION := 2.15.0
# SHA-256 of golangci-lint-2.15.0-linux-amd64.tar.gz from
# https://github.com/golangci/golangci-lint/releases/tag/v2.15.0
GOLANGCI_LINT_SHA256 := 9b7fe6a6bf6334b98decf8d86b11415ea7889afdabdd1959fdce23af2a9ebcea
GOLANGCI_LINT := $(CURDIR)/.cache/golangci-lint-$(GOLANGCI_LINT_VERSION)/golangci-lint

GOVULNCHECK_VERSION := v1.9.0
BETTERLEAKS_VERSION := v1.10.1

OSV_SCANNER_VERSION := 2.7.1
# SHA-256 of osv-scanner_linux_amd64 from
# https://github.com/google/osv-scanner/releases/tag/v2.7.1
OSV_SCANNER_SHA256 := fb7dab19a8bd9bbe5c3189f09d4b7ee1820c9553a23bac202648a56efd2f6ee8
OSV_SCANNER := $(CURDIR)/.cache/osv-scanner-$(OSV_SCANNER_VERSION)/osv-scanner
# Overridable on the command line to prove the gate red on a throwaway lockfile.
WEB_LOCKFILE := web/pnpm-lock.yaml

BUF_VERSION := 1.75.0
# SHA-256 of buf-Linux-x86_64 from
# https://github.com/bufbuild/buf/releases/tag/v1.75.0 (sha256.txt)
BUF_SHA256 := fad1e38d6909bbb583a4da81d8c02a01046e2cd663bde7e337dd24277393764b
BUF := $(CURDIR)/.cache/buf-$(BUF_VERSION)/buf
SQLC_VERSION := 1.31.1
# SHA-256 of sqlc_1.31.1_linux_amd64.tar.gz from
# https://github.com/sqlc-dev/sqlc/releases/tag/v1.31.1
SQLC_SHA256 := 497ae4fcdfa64c5b0c311ffe4c2bd991e43991e82e5367792ed78bc2dca27354
SQLC := $(CURDIR)/.cache/sqlc-$(SQLC_VERSION)/sqlc

# The contract buf breaking compares against (QG-04: the main branch).
BUF_AGAINST_REF ?= refs/remotes/origin/main

.PHONY: check check-server server-fmt server-lint server-vuln server-test server-property-long server-test-old-binary cpd-go idcheck secrets db-up db-down golden-generate
.PHONY: sqlc-generate sqlc-generated
.PHONY: check-web web-install web-fmt web-lint web-check web-vuln cpd-ts web-bundle web-unit web-e2e
.PHONY: bundle-copy casa web-preview
.PHONY: check-contract proto-fmt proto-lint proto-breaking proto-generate proto-generated

check: check-server check-web check-contract

check-server: server-fmt server-lint server-vuln sqlc-generated server-test server-test-old-binary cpd-go idcheck secrets

# gofmt is taken from the toolchain go.mod selects, not whatever is on PATH.
server-fmt:
	gofmt="$$(cd server && go env GOROOT)/bin/gofmt"; \
	unformatted="$$("$$gofmt" -l server)"; \
	if [ -n "$$unformatted" ]; then echo "gofmt needed:"; echo "$$unformatted"; exit 1; fi

server-lint: $(GOLANGCI_LINT)
	cd server && go vet ./...
	cd server && $(GOLANGCI_LINT) run ./...

server-vuln:
	cd server && go run golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION) ./...

# PostgreSQL 18 for the server tests (14 §4.1 item 1: a real database, never a
# mock), the same image locally and in CI. Official image, built from the PGDG
# packages (ADR-016 item 1). Tag 18.7-trixie, digest from
# https://hub.docker.com/_/postgres (checked 2026-11-13).
POSTGRES_IMAGE := postgres:18.7-trixie@sha256:9f4fd4aa918f52ae3d6f3b687f1de90d8f867cd4d520021545a9a50c120eae3b
# Overridable when the host already uses 5432 for something else.
CASA_DB_PORT ?= 5432
# trust is for this throwaway test container only: bound to 127.0.0.1, data on
# tmpfs, test data only (14 §7.8); no password enters the repository.
CASA_TEST_ADMIN_URL ?= postgres://postgres@127.0.0.1:$(CASA_DB_PORT)/postgres?sslmode=disable

# The image is pulled first, up to 3 times 10 s apart: a registry error is
# often transient (it once failed a CI run).
db-up:
	docker image inspect $(POSTGRES_IMAGE) >/dev/null 2>&1 || for i in 1 2 3; do \
		if docker pull -q $(POSTGRES_IMAGE) >/dev/null; then break; fi; \
		if [ $$i = 3 ]; then echo "docker pull failed 3 times"; exit 1; fi; \
		echo "docker pull failed; retrying in 10 s"; sleep 10; \
	done
	docker run -d --rm --name casa-postgres -p 127.0.0.1:$(CASA_DB_PORT):5432 \
		-e POSTGRES_HOST_AUTH_METHOD=trust --tmpfs /var/lib/postgresql $(POSTGRES_IMAGE) >/dev/null
	for _ in $$(seq 60); do \
		if docker exec casa-postgres pg_isready -U postgres -h 127.0.0.1 >/dev/null 2>&1; then exit 0; fi; \
		sleep 1; \
	done; \
	echo "casa-postgres not ready after 60 s"; docker logs casa-postgres; exit 1

db-down:
	docker rm -f casa-postgres

server-test:
	cd server && CASA_TEST_ADMIN_URL='$(CASA_TEST_ADMIN_URL)' go test -race -shuffle=on -coverprofile=coverage.out ./...

# Rewrites the name length golden vectors from household.CleanName (ES §2 (c);
# 14 §4.5; plan 1.7 decision 17). The test then checks the file is current.
golden-generate:
	cd server && CASA_GOLDEN_UPDATE=1 go test -count=1 -run '^TestGoldenNameLength$$' ./internal/household/

# The weekly long run of the stateful property test (14 §4.3 "Jumlah kasus";
# plan 1.7 decision 16): 1000 checks per property from a new seed, printed
# first so a failure can be replayed with RAPID_SEED. server-test runs the
# same test with 20 checks and seed 1 (TestMain in internal/api).
server-property-long:
	seed="$$(od -An -N8 -tu8 /dev/urandom | tr -d ' ')"; echo "RAPID_SEED=$$seed"; \
	cd server && RAPID_CHECKS=1000 RAPID_SEED="$$seed" CASA_TEST_ADMIN_URL='$(CASA_TEST_ADMIN_URL)' \
		go test -race -count=1 -timeout 5h -run '^TestHouseholdProperties$$' ./internal/api/

# QG-05 "biner lama di skema baru" (09.4 §4 row "Migrasi"; 09.2 §10: a
# rollback runs the old binary on the expand-contract schema). The tests of
# BASE_REF run against databases migrated by this branch's casa binary. The
# base ref's catalog, architecture, isolation and ID-check tests are left out: they check
# its own old schema and imports, not the old code's behaviour. -count=1:
# the test cache does not see the migrate binary, which the tests only exec.
BASE_REF ?= refs/remotes/origin/main
server-test-old-binary:
	@set -e; \
	if ! git rev-parse --verify -q '$(BASE_REF)^{commit}' >/dev/null; then \
		echo "server-test-old-binary: $(BASE_REF) not found; fetch it first"; exit 1; \
	fi; \
	mkdir -p .cache; \
	(cd server && go build -o '$(CURDIR)/.cache/casa-head' ./cmd/casa); \
	git worktree remove --force .cache/base-tree 2>/dev/null || true; \
	git worktree add --detach -q .cache/base-tree '$(BASE_REF)'; \
	trap 'git worktree remove --force .cache/base-tree' EXIT; \
	cd .cache/base-tree/server && \
	CASA_TEST_ADMIN_URL='$(CASA_TEST_ADMIN_URL)' CASA_TEST_MIGRATE_BINARY='$(CURDIR)/.cache/casa-head' \
		go test -race -p 1 -count=1 $$(go list ./... | grep -v -e /internal/schema -e /internal/architecture -e /internal/isolation -e /cmd/idcheck)

# Writes the Go code of every module's queries.sql (server/sqlc.yaml): all
# production SQL goes through sqlc, parameterized (ES §3, §4).
sqlc-generate: $(SQLC)
	cd server && $(SQLC) generate

# The committed query code must be exactly what sqlc generates (ES §2): no
# hand edit, no missed regeneration, no uncommitted new file.
sqlc-generated: sqlc-generate
	if ! git diff --exit-code -- ':(glob)server/internal/*/*db/**' \
		|| [ -n "$$(git status --porcelain -- ':(glob)server/internal/*/*db/**')" ]; then \
		git status --short -- ':(glob)server/internal/*/*db/**'; \
		echo "sqlc-generated: generated code differs from the queries; run make sqlc-generate and commit"; \
		exit 1; \
	fi

cpd-go:
	scripts/cpd.sh go server

# QG-06 (14 §4.13): every test that names an ID names a real one, and every ID
# in the scope of a finished step has a test.
idcheck:
	cd server && go run ./cmd/idcheck -docs ../docs -repo .. -scope ../traceability/scope.txt

# Betterleaks reads every commit reachable from HEAD (--log-opts=HEAD, the
# whole history of the branch being checked), not only the changed files,
# with the tool's default rules. Not --all: the CI checkout (fetch-depth 0)
# also fetches other branches, whose test fixtures are ignored only on them.
# Inline allow comments are ignored so every exception sits in a reviewable
# ignore file with its reason; findings are redacted from the output. Live validation stays off: nothing leaves the machine.
secrets:
	go run github.com/betterleaks/betterleaks@$(BETTERLEAKS_VERSION) git \
		--no-banner --redact --ignore-gitleaks-allow --log-opts=HEAD .

# ponytail: linux-amd64 only (the CI runner and current dev machine); other
# platforms need their own pinned checksum.
$(GOLANGCI_LINT):
	scripts/fetch-verified.sh \
		https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64.tar.gz \
		$(GOLANGCI_LINT_SHA256) $@.tar.gz
	tar -xzf $@.tar.gz -C $(@D) --strip-components=1 golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64/golangci-lint
	rm $@.tar.gz

# ponytail: linux-amd64 only, like golangci-lint above.
$(SQLC):
	scripts/fetch-verified.sh \
		https://github.com/sqlc-dev/sqlc/releases/download/v$(SQLC_VERSION)/sqlc_$(SQLC_VERSION)_linux_amd64.tar.gz \
		$(SQLC_SHA256) $@.tar.gz
	tar -xzf $@.tar.gz -C $(@D) sqlc
	rm $@.tar.gz

# ADR-030 item 3: the web bundle is embedded in the casa binary. bundle-copy
# builds web/ and replaces the embedded copy; casa builds the release binary.
bundle-copy:
	cd web && pnpm build
	find server/internal/webbundle/dist -mindepth 1 ! -name .gitkeep -delete
	cp -R web/build/. server/internal/webbundle/dist/

casa: bundle-copy
	cd server && go build -trimpath -o '$(CURDIR)/.cache/casa' ./cmd/casa

# The Playwright server (plan 1.5 decision 6): the embedded bundle with the
# production headers, no API and no database. Never deployed.
web-preview: bundle-copy
	cd server && go run ./cmd/webpreview -listen 127.0.0.1:4173

# Contract gates (QG-04; 14 §4.6): proto/ is the one source of the API
# contract (ES §2; ADR-013 item 4).
check-contract: proto-fmt proto-lint proto-breaking proto-generated

proto-fmt: $(BUF)
	$(BUF) format --diff --exit-code

proto-lint: $(BUF)
	$(BUF) lint

# A missing ref fails: a gate that compares nothing is never green.
proto-breaking: $(BUF)
	if ! git rev-parse --verify -q '$(BUF_AGAINST_REF)^{commit}' >/dev/null; then \
		echo "proto-breaking: $(BUF_AGAINST_REF) not found; fetch it first"; exit 1; \
	fi; \
	$(BUF) breaking --against '.git#ref=$(BUF_AGAINST_REF)'

# Writes the Go and TypeScript code of the contract (buf.gen.yaml); --clean
# removes stale output first. protoc-gen-es comes from web-install.
proto-generate: $(BUF)
	$(BUF) generate --clean

# The committed generated code must be exactly what the contract generates
# (14 §4.6 item 2; ES §2): no hand edit, no missed regeneration, no
# uncommitted new file.
proto-generated: proto-generate
	if ! git diff --exit-code -- server/internal/gen web/src/lib/gen \
		|| [ -n "$$(git status --porcelain -- server/internal/gen web/src/lib/gen)" ]; then \
		git status --short -- server/internal/gen web/src/lib/gen; \
		echo "proto-generated: generated code differs from the contract; run make proto-generate and commit"; \
		exit 1; \
	fi

# ponytail: linux-amd64 only, like golangci-lint above.
$(BUF):
	scripts/fetch-verified.sh \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Linux-x86_64 \
		$(BUF_SHA256) $@
	chmod +x $@

# Web client gates (ADR-030 items 8 and 18). web-install comes first; the
# others assume web/node_modules is in place.
check-web: web-install web-fmt web-lint web-check web-vuln cpd-ts web-bundle web-unit web-e2e

web-install:
	cd web && pnpm install --frozen-lockfile

web-fmt:
	cd web && pnpm exec prettier --check .

# svelte-kit sync writes the generated tsconfig and route types that the
# type-aware lint rules and svelte-check read.
web-lint:
	cd web && pnpm exec svelte-kit sync && pnpm exec eslint --max-warnings 0 .

# The service workers have their own tsconfig (WebWorker types), which svelte-check
# does not read.
web-check:
	cd web && pnpm exec svelte-kit sync && pnpm exec svelte-check --fail-on-warnings
	cd web && pnpm exec tsc -p src/service-worker && pnpm exec tsc -p emergency

web-vuln: $(OSV_SCANNER)
	$(OSV_SCANNER) scan source --lockfile $(WEB_LOCKFILE)

cpd-ts:
	scripts/cpd.sh typescript web

web-bundle:
	cd web && node scripts/check-bundle.mjs

# Vitest unit tests (QG-05). svelte-kit sync writes the generated tsconfig
# and $app types the tests compile against.
web-unit:
	cd web && pnpm exec svelte-kit sync && pnpm exec vitest run

web-e2e:
	cd web && pnpm exec playwright test

# ponytail: linux-amd64 only, like golangci-lint above.
$(OSV_SCANNER):
	scripts/fetch-verified.sh \
		https://github.com/google/osv-scanner/releases/download/v$(OSV_SCANNER_VERSION)/osv-scanner_linux_amd64 \
		$(OSV_SCANNER_SHA256) $@
	chmod +x $@
