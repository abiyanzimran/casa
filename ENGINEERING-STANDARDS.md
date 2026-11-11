# Engineering Standards

Hard rules for the build phase. They are stack-neutral; stage S4 adds the stack-specific tooling (formatter, linter, test runner) by ADR. A rule here can be changed only by an ADR approved by the owner.

Planning documents must not recommend anything that breaks these rules.

## 1. Clean code

- One responsibility per module, type, and function. If a name needs "and", split it.
- Names say what a thing is in the domain's words, using the terms in `docs/00-glossary.md`. No abbreviations that are not in the glossary.
- No dead code, commented-out code, unused parameters, or unused dependencies in a merged change.
- No abstraction without a second real use: no interface with one implementation, no configuration for a value that never changes.
- Formatter and linter run in CI with zero warnings. A suppression needs a comment giving the reason.
- Comments explain why, not what. Public APIs and every financial rule carry a comment that cites the requirement or business-rule ID.
- Errors are handled or returned with context. Nothing is swallowed, and nothing fails silently.

## 2. No duplicated code

- Before writing a helper, type, or query, search the codebase for an existing one and reuse it.
- Every business rule has exactly one authoritative implementation. The server is always the final authority on money and approvals. Where a client must apply a rule while offline, it uses that same implementation (shared or generated code), never a hand-written copy, and the server re-checks the result on sync. Where rules run is decided at S4.
- The API contract is defined once and client types are generated from it. Hand-written copies of request or response types are not allowed.
- Validation rules, error codes, and constants are defined once and shared.
- Screens are built from the design system's components. A second copy of a component is a defect.
- CI runs a duplication check; a new duplicate block fails the build.

## 3. Performance

Targets marked `[DEFAULT]` are starting values, confirmed or changed at S4 by measurement on the pilot devices (iPhone 13, Galaxy A55).

| Area | Target |
|---|---|
| Cold start to a usable home screen | `[DEFAULT]` ≤ 2 s |
| Recording an expense | `[DEFAULT]` ≤ 3 taps after opening the entry screen, saved locally at once |
| Scrolling and transitions | No dropped frames at the device's refresh rate on the pilot devices |
| API response, 95th percentile, on the pilot VPS | `[DEFAULT]` ≤ 200 ms for reads, ≤ 400 ms for writes |
| App download size | `[DEFAULT]` ≤ 30 MB |

- Every list is paginated or virtualized. No unbounded query reaches the database.
- Every query used by a screen has an index that serves it; a query plan is reviewed for each new one.
- No N+1 queries. No network or disk work on the UI thread.
- A performance claim is backed by a measurement in the change description. No optimization without one.

## 4. Security and privacy

- Authorization is enforced on the server for every request, by household and by member. Client-side checks are for display only.
- Every query is scoped to a household. A test proves that a member of one household cannot read or change another household's data.
- All input is validated at the server boundary. All SQL is parameterized.
- Secrets never enter the repository, logs, or client bundles. Tokens are stored in the platform's secure storage.
- Transport is TLS only. Sensitive data at rest is encrypted as decided in `docs/10-security-privacy.md`.
- Logs and analytics never contain amounts, balances, mood or health content, message text, or personal identifiers.
- Lock-screen notifications never show amounts or wellbeing details.
- Dependencies are pinned and scanned for known vulnerabilities in CI. A new dependency needs a stated reason and a check of its license and maintenance.
- Wellbeing data is visible only to the audience its owner chose, and that choice is revocable.

## 5. Financial correctness

- Money is stored and computed as integers in the currency's smallest unit. Floating point is never used for money.
- Every write that changes money runs in one database transaction and carries an idempotency key, so a retry or a double tap cannot record it twice.
- Financial history is append-only. A correction is a reversing or adjusting entry, never an edit or delete of the original.
- A transfer between family-controlled holdings is never income or expense.
- The invariants defined in `docs/07-financial-domain.md` are automated tests that run on every change; a failure blocks the merge.
- Each worked example in `docs/07-financial-domain.md` exists as an executable test with the same numbers.

## 6. Testing

- Test-driven: a failing test is written before the code that makes it pass.
- Every Tier 1 acceptance criterion in `docs/05-functional-requirements.md` has an automated test that names its requirement ID.
- Financial and approval logic have property-based tests for their invariants in addition to example tests.
- A bug fix starts with a test that reproduces the bug.
- Tests are deterministic. A flaky test is fixed or removed within the same week, never ignored.

## 7. Review and delivery

- Work happens on a branch per task, never directly on the main branch.
- Every change is reviewed by a person or review agent other than its author, in two passes: does it meet the requirement, then is the code sound.
- A change does one thing. Refactoring and behavior change go in separate changes.
- Commit messages state what changed and why, and cite the requirement or ADR ID.
- CI must pass before merge: format, lint, duplication check, tests, vulnerability scan.
- Database schema changes go through versioned, reversible migrations.

## 8. Dependencies and versions

Applies to everything installed: language toolchains, frameworks, SDKs, libraries, build tools, the database, and container images.

- **Checked live.** At the moment of installing or upgrading, the current version is looked up in the official registry or release page. A version number from memory, a tutorial, or an article is never used.
- **Newest stable only.** No alpha, beta, release candidate, canary, or nightly build, and no API its maintainers mark experimental or unstable. An exception needs an ADR.
- **Known problems checked.** Before adopting a version, read its release notes and check the project's security advisories and open critical issues. If the newest stable release has a known defect that affects this product, use the newest stable release without that defect, and record the reason and the condition for moving up.
- **Compatible as a set.** Versions must be supported together (framework, OS SDK, build tool, language). Follow the framework's own compatibility table, not the newest number of each part.
- **Pinned.** Exact versions and a committed lockfile. No floating ranges.
- **Recorded.** Every added or upgraded dependency states, in its change description: version, date checked, source link, and why it is needed.
- **Maintained.** No deprecated or unmaintained package, and no deprecated API.
- **Kept current.** Updates are checked `[DEFAULT]` monthly, and at once for a security advisory. An upgrade is its own change and passes the full test suite.
- Minimum supported OS versions are set by ADR at S4.

## 9. Definition of done

A task is done only when all of these are true:

1. It traces to a requirement ID and serves the goal in `CLAUDE.md`.
2. Its tests exist, were seen to fail first, and now pass.
3. CI is green.
4. It was reviewed and the findings were resolved.
5. It was run on a real pilot device or against the real API, not only in tests.
6. Documentation and ADRs affected by it are updated in the same change.
