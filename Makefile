# Single source of the CI gate commands: CI calls these targets and nothing
# else, so a local run and a CI run check the same things (ES §2, §7).
# Gates: QG-01 format and lint, QG-02 duplication, QG-03 vulnerabilities,
# QG-05 tests (docs/14-engineering-roadmap.md §5.1), plus the repository
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

.PHONY: check-server server-fmt server-lint server-vuln server-test cpd-go secrets

check-server: server-fmt server-lint server-vuln server-test cpd-go secrets

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

server-test:
	cd server && go test -race -shuffle=on -coverprofile=coverage.out ./...

cpd-go:
	scripts/cpd.sh go server

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
