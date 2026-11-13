# Single source of the CI gate commands: CI calls these targets and nothing
# else, so a local run and a CI run check the same things (ES §2, §7).
# Gates: QG-01 format and lint, QG-02 duplication, QG-03 vulnerabilities,
# QG-05 tests, QG-06 ID traceability (docs/14-engineering-roadmap.md §5.1,
# §4.13), plus the repository secrets check (09.4 §5).

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

.PHONY: check check-server server-fmt server-lint server-vuln server-test cpd-go idcheck secrets db-up db-down
.PHONY: check-web web-install web-fmt web-lint web-check web-vuln cpd-ts web-bundle web-unit web-e2e

check: check-server check-web

check-server: server-fmt server-lint server-vuln server-test cpd-go idcheck secrets

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

db-up:
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

cpd-go:
	scripts/cpd.sh go server

# QG-06 (14 §4.13): every test that names an ID names a real one, and every ID
# in the scope of a finished step has a test.
idcheck:
	cd server && go run ./cmd/idcheck -docs ../docs -repo .. -scope ../traceability/scope.txt

# Betterleaks' default git scan reads every commit reachable from any ref
# (git log --full-history --all), not only the changed files, with the
# tool's default rules. Inline allow comments are ignored so every exception
# sits in a reviewable ignore file with its reason; findings are redacted
# from the output. Live validation stays off: nothing leaves the machine.
secrets:
	go run github.com/betterleaks/betterleaks@$(BETTERLEAKS_VERSION) git \
		--no-banner --redact --ignore-gitleaks-allow .

# ponytail: linux-amd64 only (the CI runner and current dev machine); other
# platforms need their own pinned checksum.
$(GOLANGCI_LINT):
	scripts/fetch-verified.sh \
		https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64.tar.gz \
		$(GOLANGCI_LINT_SHA256) $@.tar.gz
	tar -xzf $@.tar.gz -C $(@D) --strip-components=1 golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64/golangci-lint
	rm $@.tar.gz

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

web-check:
	cd web && pnpm exec svelte-kit sync && pnpm exec svelte-check --fail-on-warnings

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
