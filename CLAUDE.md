# Family OS

## The one goal

A husband and wife know the state of their household money and household matters without having to ask each other, at the lowest possible recording effort.

Every document, decision, and (later) line of code must trace back to this goal and to the core problem in `family-os-brief.md` §1.1. Work that does not serve it is cut or deferred, and the reason is recorded.

## Current phase: planning only

No application code, project scaffolding, dependency installs, or boilerplate. Output is Markdown under `docs/`. Pseudocode, schema sketches, state tables, and Mermaid diagrams are allowed inside documents. The phase ends only when the owner says so.

## Sources of truth

Higher entries win. When two sources conflict, follow the higher one and log the conflict in `docs/17-open-questions-assumptions.md`.

1. `family-os-brief.md` — direction, scope, MVP tiers, stages, checkpoints.
2. `docs/00-glossary.md` — the single definition of every shared term.
3. `docs/adr/` — decisions already made.
4. `ENGINEERING-STANDARDS.md` — hard rules for the build phase; planning documents must be compatible with them.
5. Everything else in `docs/`.

## Instruction priority

1. This file and the brief.
2. Superpowers skills. During planning, the brief's stages S1–S6 are the process; do not run Superpowers brainstorming or plan-writing on top of them. In the build phase, Superpowers is the default workflow (plans, worktrees, test-driven development, subagent-driven development, code review).
3. Ponytail. It applies to scope: cut what the pilot household does not need. It never reduces the depth of Tier 1 documents, tests, input validation, or security.
4. Caveman. It applies to chat only, never to file contents.

## Language

| Where | Language |
|---|---|
| Everything under `docs/` | Professional, natural Indonesian; established technical terms stay in English |
| `CLAUDE.md`, the brief, `ENGINEERING-STANDARDS.md`, agent definitions | English |
| Code, comments, identifiers, commit messages | English |
| Conversation with the owner | Indonesian |

## Team

The main session is the lead: it assigns work per stage, merges results, keeps `docs/_progress.md` current, and stops at checkpoints. Agents are defined in `.claude/agents/`.

| Agent | Owns |
|---|---|
| `researcher` | `18-market-research.md`; sourced findings for other documents |
| `finance-domain` | `07-financial-domain.md`; financial terms in `00-glossary.md` |
| `product-analyst` | `02`, `03`, `04`, `05`, `11`, `12`, `13`, `17` |
| `ux-designer` | `08-ux-ui-strategy.md` |
| `architect` | `06`, `09`, `14`, `15`, `docs/adr/` |
| `security-privacy` | `10-security-privacy.md` |
| `critic` | Review reports; never writes deliverables |

The lead writes `README.md`, `_progress.md`, `00-glossary.md`, `16-final-recommendations.md`, and `01-executive-summary.md`. The glossary is the one shared file: `finance-domain` writes its financial-terms section, and the lead writes the rest.

## Build team (created when the build phase starts)

Agent files for the build phase are written by the lead only after the owner approves the stack at S4, because their content depends on it. Until then this table is the plan.

| Role | Does |
|---|---|
| `backend-dev` | Server, database, migrations, API contract |
| `ios-dev`, `android-dev` | One per platform if the UI is fully native; a single `mobile-dev` if S4 chooses a shared framework |
| `ux-designer` (continues) | Reviews every built screen against `docs/08-ux-ui-strategy.md` and the platform's guidelines |
| `code-reviewer` | Two-pass review of every change: meets the requirement, then code quality against `ENGINEERING-STANDARDS.md` |
| `qa` | Proves Tier 1 acceptance criteria and financial invariants end to end, on the pilot devices |

Every build agent follows `ENGINEERING-STANDARDS.md` without exception, including live version checks (§8), and works through the Superpowers workflow. Plugins and skills are installed only for the stack that was chosen, and only if a named role will use them.

## Working rules

- **One owner per document.** Another agent that needs a change asks the owner or the lead; it does not edit the file.
- **Authors do not review their own work.** `critic` reviews every stage before its checkpoint. Blocking findings are fixed before the lead presents the stage.
- **Checkpoints belong to the owner.** At every **STOP** in the brief §0.3, present the summary and wait. Do not start the next stage without an explicit "lanjut".
- **Evidence.** Use the labels in the brief §16. Never present an assumption as fact and never claim research that was not done.
- **Identifiers and traceability.** Use the scheme in the brief §16. Every Tier 1 requirement maps to a feature, a scenario, and a roadmap phase.
- **Depth follows the MVP tiers** in the brief §1.1: full for Tier 1, concise for Tier 2, one page for Tier 3.
- **Computed, not read.** The numbers in every financial worked example are checked by computation (a throwaway script kept outside the repository) before the checkpoint. Reading the table is not a check.
- **Deferred on purpose.** The mobile stack is recommended at stage S4, not earlier. Which platform is built first is the owner's decision when the build phase starts.
- **Git.** One commit per completed stage, plus one for each owner-approved revision. Commit only finished, reviewed work. Planning commits go to the main branch; branch-per-task starts with the build phase.
