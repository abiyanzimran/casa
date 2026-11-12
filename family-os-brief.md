# PROJECT BRIEF — FAMILY OPERATING SYSTEM (Family OS)
## Product Discovery, Domain Design, UX Strategy & Technical Architecture — PLANNING PHASE ONLY

---

## 0. HOW TO RUN THIS BRIEF (READ FIRST)

This brief is large. Do not try to produce every deliverable in one pass — that produces long but shallow documents and risks losing context midway. Work in the stages below, keep a progress log, and stop at each checkpoint.

### 0.1 Hard constraints

- **Planning only.** Do not write application code, initialize projects, install dependencies, generate boilerplate, or create implementation files. Pseudocode, schema sketches, state tables, and Mermaid diagrams inside documentation are allowed when they clarify a design.
- **Language.** All documents under `docs/` are written in professional, natural Indonesian. Instruction files (this brief, `CLAUDE.md`, `ENGINEERING-STANDARDS.md`, agent definitions) are in English. Established technical terms (ledger, offline-first, state machine, RBAC, etc.) stay in English.
- **Honesty about evidence.** Never present an assumption as fact, never invent competitor capabilities, statistics, or API availability, and never claim research was done unless you actually performed it.

### 0.2 Progress log

Create and maintain `docs/_progress.md`. After every stage, record:
- Documents completed and their status (draft / reviewed / final).
- Key decisions made (with links to ADRs).
- Assumptions adopted.
- Open questions still waiting on the product owner.
- The next stage to execute.

If your context is cleared or the session restarts, read `docs/_progress.md` and `docs/00-glossary.md` first, then resume from the recorded next stage.

### 0.3 Execution stages and checkpoints

| Stage | Output | Checkpoint |
|---|---|---|
| **S1 — Framing** | `README.md` (skeleton + links), `00-glossary.md` (first version; financial terms are working definitions, finalized in S2), `17-open-questions-assumptions.md` (first version), `18-market-research.md` (market and regulation at S1; extended whenever a later stage commissions research) | **STOP.** Present your top 10 most decision-critical questions with a proposed default for each. Wait for my answers, or for "lanjut pakai default". |
| **S2 — Domain foundation** | `02-product-principles.md`, `03-personas-household-models.md`, `07-financial-domain.md`, `06-activity-approval-architecture.md` | **STOP.** Summarize the financial model and approval model in ≤ 1 page for my review. These two documents constrain everything else. |
| **S3 — Product definition** | `04-feature-map.md`, `05-functional-requirements.md`, `11-innovation.md`, `12-scenarios.md` | **STOP.** Short summary of the MVP feature cut. Wait for "lanjut". |
| **S4 — Experience & architecture** | `08-ux-ui-strategy.md`, `09-technical-architecture.md` (+ ADRs), `10-security-privacy.md` | **STOP.** Present the recommended stack and architecture with the main tradeoffs. |
| **S5 — Plan & synthesis** | `13-product-roadmap.md`, `14-engineering-roadmap.md`, `15-risk-assessment.md`, `16-final-recommendations.md`, then `01-executive-summary.md` **last** | — |
| **S6 — Consistency pass** | Fix contradictions, broken links, undefined terms, untraced requirements | Report what was changed. |

An owner decision recorded at a checkpoint (in `docs/17-open-questions-assumptions.md`) ranks with this brief on the matter it decides. Where this brief's wording has not been updated yet, the decision wins and the difference is logged in `docs/17-open-questions-assumptions.md`.

Write the executive summary last, because it must summarize decisions that do not exist yet at the start.

---

## 1. CONTEXT & CONSTRAINTS

These constraints shape every recommendation, especially technology choice, scope, and roadmap. Values marked `[DEFAULT]` are my current assumptions; treat them as given unless I change them.

| Item | Value |
|---|---|
| Product owner | Solo founder / small studio, technical background |
| Team size for v1 | `[DEFAULT]` 1–3 developers |
| Team strengths | Strong in Go and Rust backend; extensive mobile experience. The owner is open to any mobile stack and wants the best recommendation, not the most familiar one |
| First users (pilot) | The owner's own household: husband and wife first, for several months. Parents and children get logins later. Public launch only after the pilot shows sustained daily use. Architecture and privacy design must not need a rewrite to go public |
| Pilot devices | iPhone 13 (wife, main recorder), iPad mini 7 and Samsung Galaxy A55 (owner). Both spouses must be able to use the app. Real data starts on iOS and iPadOS; Android is built after real data starts and joins the pilot before the beta (decision 42) |
| Build environment | The owner develops on Linux and has no Mac. Any iOS recommendation must state how the app gets built, tested on a real iPhone, and distributed, and what that costs |
| Engineering standards | `ENGINEERING-STANDARDS.md` holds the hard rules for the build phase (clean code, no duplication, performance, security, financial correctness, testing, review). Every recommendation must be compatible with it; `14-engineering-roadmap.md` adopts it and adds the stack-specific tooling |
| Team and working rules | `CLAUDE.md` defines the single goal, the order of sources of truth, document ownership per agent, and the review rule |
| Experience bar | WhatsApp-class: instant cold start, light, reliable on poor networks, professional finish. Recording something must take seconds |
| Primary launch market | `[DEFAULT]` Indonesia (IDR, Bahasa Indonesia first, English second); design for later internationalization |
| Platforms | iPhone, Android phones, iPad, Android tablets, foldables. Phones come first, except the iPad mini 7: it is a daily pilot device, so every Tier 1 flow it uses is laid out for it before real data. See the tiers in §1.1 |
| Target time to MVP | No fixed deadline. Prefer the smallest scope the pilot household will use daily over a date |
| Monthly infrastructure budget at launch | One modest VPS during the pilot. No usage-priced managed services unless clearly justified. The AI provider for §5O is the one accepted exception, under its hard cost cap |
| Business model | `[DEFAULT]` Undecided. Evaluate freemium per household, family subscription, and one-time purchase; recommend one with reasoning |
| **Money movement scope** | `[DEFAULT]` **The app records and organizes money; it does not hold, transfer, or process real funds.** No payment gateway, no stored value, no e-wallet. All "transfers" are records of real-world movements. Evaluate the regulatory and product implications of ever crossing this line, but do not design for it in MVP |
| Bank / e-wallet sync | `[DEFAULT]` Not in MVP. Investigate what is realistically available in Indonesia and label findings as verified or unverified |

If any of these values makes a requirement in this brief unrealistic, say so explicitly and propose the adjustment.

### 1.1 Goal, core problem & success criteria

**Goal.** Staged: first a product the owner's own household uses every day, then a public product for Indonesian families. The pilot household is the primary source of truth for MVP scope.

**Core problem (real, from the pilot household).** Household money is invisible and unplanned, and coordinating it costs constant asking.

- **Money is not visible.** The husband keeps asking "is the monthly money still there, how much do we have left, what assets do we own?" Only grocery spending is recorded today; nothing else is.
- **Custody is unclear.** After ATM withdrawals and handovers, nobody knows who holds how much or where the money went.
- **Actual money flow.** Salary of about IDR 10,000,000 arrives; installments and bills (electricity, wifi, work subscriptions) are sometimes paid first; IDR 3,000,000 goes to the husband's mother; the wife holds the rest; the husband later asks for, or borrows back, money from what the wife holds. Occasional support or loans to the husband's sibling and father often go unrecorded.
- **Spending and outings are unplanned.** Outings are announced at the last minute and grocery budgets are regularly exceeded. Wanted: either spouse proposes an outing or purchase in advance (date, places, planned items, budget); the other checks the budget and approves, discusses, or counters.
- **Recording is a burden.** The main recorder is forgetful, so entry must take seconds and tolerate late entry.
- **Secondary pains.** A spouse's mood is unknown unless asked; house issues (e.g., leaking roof) should become a task closed with photos and notes; routine chores are forgotten (e.g., cooking rice); "what should we cook?"; illness needs follow-up (fever reported, family checks again after N days, the sick member confirms recovered, still sick, or going to a clinic); who pays which bill, and how much, must be clear.
- **Warmth takes effort.** The couple wants the app to help them be attentive and affectionate: a kind word in the morning, help finding the right words when the other is unwell or low, and things one asks of the other kept as tasks until done. See §5D and §5O.

**MVP priority tiers.** Set by the owner. Challenge them, with reasons, if research or the financial model says otherwise; a change takes effect only after the owner approves it at a checkpoint.

| Tier | Scope | Documentation depth |
|---|---|---|
| **1 — Core** | `FAM` as far as one household with two adult accounts needs: creation, invitation, and the roles and permissions that `FIN` and `REQ` use. `FIN`: custody, income allocation, expenses, monthly budgets per category, with a period start day chosen by the household (default: the 1st), recurring bills, installments, debts and loans (between spouses and to extended family), simple savings tracking, assets and net worth, summary reports. `NTF` for requests and approvals. `REQ`: budget and outing requests, discussion and counter-proposal on the request itself, approval, settlement of the remainder. `COM` light household chat: free messages, photos, object cards (requests, agreements, tasks, entries, handovers, shared care status) whose buttons run the object's own commands, delivered and read marks, and presence; no voice notes, groups, reactions, or search (owner decision, S3). `NTF` quiet hours and digests, and bill reminders (owner decision, S3 checkpoint). `FAM` account protection: optional app lock, the member's own sessions and security history, account deletion, and a Help page. `COM` hide-for-me and report on the household chat (owner decision, S4 checkpoint) | Full requirements with acceptance criteria |
| **2 — Light versions** | `SHP` shopping lists with estimated and actual cost against a budget; `MNT` issue to task; `CARE` voluntary mood status (what the member feels and, optionally, why; once the member has chosen to share with their partner, the partner is notified on every change) and illness follow-up; `CAL`/`NTF` routine reminders; `PLN` only as far as outing requests need it; `REQ` action requests: a task one member asks of another, which the assignee can accept or decline and the requester can cancel, and which stays pending until the assignee marks it done; `AI` light version: morning greeting, suggested words of support, warm reminders, monthly narrative; `FAM` member labels: each member may name the other members for their own screens only (owner decision, S4 checkpoint) | Concise requirements |
| **3 — Deferred** | `AI` open-ended advice conversations, `COM` beyond the light household chat (voice notes, groups, reactions, search), `INV`, `MEAL`, `GOL` beyond simple savings tracking, `TML`, online shopping integrations, multiple households, dedicated tablet and foldable screen designs (layouts must still adapt to any width from day one; see §10.3) | One-page outline per module |

Any part of a module not named in Tier 1 or Tier 2 is Tier 3.

Tiers set what is built first, not what is designed: the data model, rules, flows, terms, and screens are designed for every household of the public stage, and the pilot household is their first user (owner direction, S3).

This depth rule overrides §5, §10.3, §12, and §16 wherever they ask for uniform detail across all modules. A scenario that falls in Tier 3 scope is written at outline level and needs no requirement IDs.

Deferring a module limits its documentation and its build, not the foundations: the Tier 1 data model and access model must not block Tier 3, in particular household isolation and one person belonging to more than one household.

**Success after three months of pilot use.**

1. Either spouse can answer "how much money is left and who holds it" from the app, without asking the other.
2. Both spouses record without being reminded.
3. The monthly summary shows spending by category and per person, who paid what, savings, and assets.
4. Outings and non-routine purchases arrive as requests in the app, not as surprises.
5. Each spouse says the app helped them notice and respond to how the other was doing. This one is judged by the couple, not measured.

**Guardrail.** Requests run in both directions (each spouse asks the other), and mood sharing is always the sharer's choice. See §5D and §8.

---

## 2. YOUR MANDATE

Act as a senior cross-functional product team: product strategy, domain/financial modeling, UX research and design, mobile and backend architecture, security and privacy, and QA/accessibility.

Your job is not to convert my requirements into a feature checklist. Your job is to:

1. Understand the underlying family problems behind each requirement.
2. Challenge requirements that add complexity, privacy risk, financial inconsistency, or poor UX — and propose a better alternative.
3. Find important problems I have not mentioned.
4. Make decisions. When there are several options, compare them briefly and then **recommend one**. Do not end sections with open-ended option lists.
5. Keep the product simple for users even where the internal model is sophisticated.

Treat this as a product that could eventually serve thousands to millions of families, but do not over-engineer v1 for traffic it will not have.

When information is missing, make a reasonable assumption, record it in `17-open-questions-assumptions.md`, and continue. Only stop at the defined checkpoints.

---

## 3. PRODUCT VISION

Build a unified **Family Operating System**: one integrated mobile environment where a family organizes, communicates, plans, manages money, shares responsibilities, makes decisions, and cares for one another.

It must go beyond existing family organizers, expense trackers, chat apps, and task managers by connecting these capabilities into one coherent experience — not by bundling them side by side.

### Product philosophy

The app should feel like a **digital home for the family**: warm, calm, trustworthy, modern, and easy for both young and older members. It must not feel like an enterprise dashboard disguised as a family app.

**Core principle: sophisticated systems, effortless everyday experiences.**

**Ultimate purpose:** help families live more organized, financially aware, coordinated, and connected lives *with less effort* — not make them track more information. Every feature must justify the recurring effort it asks of users.

---

## 4. FAMILY STRUCTURES & IDENTITY MODEL

### 4.1 Supported structures

Do not hardcode "husband + wife + children". Support at least:

1. Married couple without children.
2. Couple with children.
3. Single-parent household.
4. Parents with adult children.
5. Extended families.
6. Grandparents raising grandchildren.
7. One family across multiple residences.
8. Families financially supporting parents or in-laws.
9. Guardians and dependents.
10. Household helpers or caregivers with restricted access.

### 4.2 Accounts and membership

Every eligible member can have an individual account. Design: family creation, invitations, onboarding, membership, multiple households, profiles, age-appropriate access, guardianship, administration and ownership transfer, member exit (an adult leaves on their own; no adult can remove another adult; owner decision, S3), account recovery, household separation, and the boundary between shared and personal information.

Explore how one person can belong to multiple households (e.g., a married son who still supports his parents' household) without leaking information between them.

### 4.3 Five independent concepts

These must be modeled separately and must not automatically derive from each other:

| Concept | Meaning | Example values |
|---|---|---|
| **Relationship** | How members are related | husband, wife, child, parent, in-law, sibling, guardian |
| **Role** | System-level standing | owner, admin, member, dependent, guest |
| **Permission** | What a member can view or modify | view shared budget, edit shopping list |
| **Responsibility** | What a member is expected to manage | groceries, electricity bill, school pickup |
| **Approval authority** | Which requests a member can approve, and up to what limit | approve purchases ≤ IDR 500,000 |

A husband must not automatically have more permissions than a wife, and a wife must not automatically control all finances. Each family configures its own household rules, with sensible presets so families are not forced to configure everything on day one.

---

## 5. CORE MODULES

For each module, analyze: purpose, target users, problems solved, features, primary and secondary workflows, relationships with other modules, permissions, data and privacy requirements, edge cases, notifications, UX considerations, and implementation priority.

The feature lists below are starting points, not a specification. Add, merge, cut, or defer as justified.

**Module codes** (use them in all identifiers): `FAM`, `COM`, `FIN`, `REQ`, `SHP`, `INV`, `PLN`, `GOL`, `MNT`, `MEAL`, `CARE`, `CAL`, `TML`, `NTF`, `AI`.

### A — Family & Household Management (`FAM`)
Family and household creation, invitations, member profiles, relationship visualization, configurable roles, fine-grained permissions, guardian/dependent links, household-specific roles, contact directory, announcements, member activity history, responsibility delegation, emergency contacts.

### B — Family Communication (`COM`)
Group chat, private conversations, topic threads, voice messages, attachments, reactions, pins, announcements, polls, shared decisions, search, and **contextual conversations attached to activities**.

Contextual communication is the key requirement: a shopping request, a roof repair, or a vacation plan each carries its own discussion, decisions, expenses, and updates.

**Decide explicitly:** is a full messaging platform needed for MVP, or do contextual discussions on activities deliver more value while families keep using WhatsApp for general chat? Justify the decision, including the cost of building and moderating real-time chat.

### C — Family Financial Management (`FIN`)
One of the most important modules. Must support detailed financial organization without becoming hard to use.

- **Income:** salary, business income, allowances, gifts, refunds, other; recurring and irregular; income ownership and allocation.
- **Expenses:** daily/weekly/monthly/annual, recurring and one-time, shared vs personal vs for a specific member, categories and subcategories, receipts, notes, history.
- **Money & balances:** cash, bank accounts, e-wallets, shared funds, personal funds, funds held by another member, transfers, allocated vs unallocated money, savings, emergency funds, debts and receivables.
- **Budgets:** daily/weekly/monthly/custom periods, per category, per member, per household, per event; allocation, remaining balance, overspending warnings, adjustments, rollover rules, approval, utilization analysis. **Pilot scope (owner decision, S1 and S2 checkpoints):** Tier 1 builds monthly budgets per category with a household-chosen start day. Other periods, per-member and per-event budgets, and rollover rules are Tier 3.
- **Reports:** income vs expense, cash flow, budget vs actual, trends, by category / member / beneficiary / household, recurring costs, savings progress, history, export.

**Custody example.** The family has IDR 10,000,000 in total:
- Husband holds IDR 4,000,000.
- Wife holds IDR 3,000,000.
- Mother-in-law holds IDR 1,000,000 for household meals.
- Shared bank account holds IDR 2,000,000.

The model must distinguish, for any amount of money:
1. Who **owns** it (legally or by family agreement).
2. Who **currently holds** it.
3. **Where** it is stored.
4. What **purpose** it is allocated to.
5. Who is **authorized to spend** it.
6. Who **benefits** from the spending.

Money moving between members, accounts, or wallets must never be double-counted. Compare a full double-entry ledger, a simplified ledger, and a pure transaction list; recommend one that guarantees accuracy and auditability while showing ordinary families a simple interface with no accounting vocabulary.

### D — Requests, Permissions & Approvals (`REQ`)
Scenarios: husband requests a spending budget from wife; wife requests a purchase budget from husband; child asks parents for permission; member proposes a large purchase; parent asks for help; reimbursement request; request to use shared funds; outing proposal.

**Reference workflow:**
1. Husband proposes buying food, requested budget IDR 100,000.
2. Wife is notified.
3. Wife approves, rejects, discusses, or counter-proposes an amount.
4. The request becomes authorized.
5. Husband buys the food; actual spend is IDR 75,000.
6. He attaches the receipt.
7. The remaining IDR 25,000 is reconciled correctly (define what "correctly" means in each custody situation).
8. Financial records and activity history update.

Explore: configurable rules, single and multiple approvers, sequential approval, thresholds, deadlines, delegation, modification requests, rejection reasons, cancellation, partial completion, reimbursement, settlement, audit trails, and dispute handling.

These are **different request types** and must not be modeled as one financial transaction: asking permission, requesting a budget, requesting reimbursement, requesting an action, proposing a decision.

An **action request** is a task one member asks of another (for example, one spouse asks the other to fix something). The assignee accepts or declines, the requester can cancel, and an accepted task stays pending, with reminders, until the assignee marks it done.

Approval flows must support cooperation, not coercive control or surveillance. Define concrete design safeguards for this (see §8).

### E — Shopping & Purchasing (`SHP`)
- **Lists:** shared and personal lists, groceries, household supplies, quantity and units, priority, preferred brand, estimated price, assigned shopper, status, purchase confirmation, recurring lists, history.
- **Online:** evaluate Shopee, Alfagift, KlikIndomaret, and other marketplaces or grocery delivery services. **Do not assume official APIs exist.** For each, research and classify what is realistically possible — official partner API, deep link, product link, receipt import, manual order tracking — and label each finding as *verified (with source)* or *unverified*.
- **Offline:** Alfamart, Indomaret, supermarkets, traditional markets (pasar), wholesale stores, local merchants (warung). Include receipts, actual cost, remaining budget, and reconciliation. Note that many traditional-market purchases have no receipt.

### F — Household Inventory (`INV`)
Consumables such as rice, cooking oil, eggs, milk, toiletries, and cleaning products: current and minimum stock, units, estimated consumption, expiry, low-stock reminders, shopping-list suggestions, purchase-price history, supply estimates.

Families will not manually update dozens of items daily. Propose low-effort approaches (e.g., stock updated as a side effect of purchases, "running low" one-tap reports, consumption estimates from purchase intervals) and decide whether full inventory belongs in an early release at all.

### G — Family Plans & Activities (`PLN`)
Eating out, vacations, visiting relatives, weekend outings, birthdays, school activities, medical appointments, religious or cultural gatherings (e.g., mudik, arisan, hajatan), renovations, major purchases.

A plan may include: title, description, organizer, participants, schedule, location, estimated budget, approval status, discussion, checklist, linked expenses, progress, attachments, completion summary. Plans must generate financial records without duplicate data entry.

### H — Goals & Savings (`GOL`)
Emergency fund, school fees, house, vehicle, vacation, renovation, wedding, retirement, household purchases. Goal creation, target amount and date, individual and shared contributions, progress, milestones, prioritization, funding projections, completion, cancellation and fund reallocation.

Strictly distinguish **money earmarked for a goal** from **money that actually exists**.

### I — Household Issues & Maintenance (`MNT`)
Leaking roof, broken lights, plumbing, electrical, appliances, vehicles, internet outages, urgent repairs. Reporting, severity, photos, assignee, discussion, cost estimate, budget request, scheduling, progress, receipts, resolution notes, maintenance history, recurring maintenance reminders. Include emergency handling and prioritization.

### J — Cooking & Meal Planning (`MEAL`)
Daily and weekly meal plans, family recipes, ingredients, portions, cooking responsibilities, dietary preferences, ingredient availability, shopping-list generation, meal budgets, schedules, leftovers. Connect with inventory and shopping. Assess honestly whether this belongs in an early phase.

### K — Wellbeing & Care (`CARE`)
Tired, sick, low energy, stressed, mood updates, needing rest, needing help, menstrual-cycle and PMS information, recovery status, medication reminders, care responsibilities.

Requirements:
- Private by default; nothing is shared automatically with all members.
- Explicit, granular, revocable consent before any sharing.
- No unsupported medical conclusions and no diagnostic language.
- No intrusive monitoring; support and empathy are the purpose.
- Design a way to **ask for help without revealing medical details** (e.g., "Saya butuh istirahat hari ini" without stating why).

Evaluate whether cycle tracking should be built at all, given the sensitivity of the data, versus integrating with dedicated apps or deferring it.

### L — Calendar & Routines (`CAL`)
Shared and personal calendars, family events, school schedules, medication and bill reminders, chores, recurring responsibilities, upcoming payments, birthdays and anniversaries, custom reminders, time zones (members may live in different cities or countries), and calendar sync options (verify what Google Calendar and Apple Calendar integration realistically allow).

### M — Timeline & History (`TML`)
Milestones, completed goals, repairs, past events, major purchases, financial milestones, photos, important decisions. Define the boundary between the **operational audit history** (immutable, for correctness) and a **family memory album** (curated, emotional, editable).

### N — Notifications & Activity Feed (`NTF`)
Approval requests and outcomes, upcoming bills, budget warnings, shopping reminders, announcements, urgent issues, plans and events, task assignments, help requests. Design priority levels, quiet hours, digests, granular preferences, and **privacy-safe notification content** on lock screens. Minimize notification fatigue.

### O — Family Companion (`AI`)
An AI companion that helps members care for each other and keeps the household warm. It assists; the people act.

- **Morning greeting.** On opening the app, a short personal line of encouragement or affection.
- **Suggested words of support.** When a member shares a low mood or an illness, their partner is prompted with a suggested message of encouragement or praise. The partner sends it as is, edits it by hand, or asks for another suggestion. During the pilot no member's text, status, or object titles are sent to the AI provider: the AI writes pools of phrases without household data, and the server in Indonesia picks among them (owner decision, S4 checkpoint). It is sent as a household-chat reply to that status's card, so it is also attached to the status (owner decision, S3).
- **Warm reminders.** Pending tasks, requests, and outings are worded kindly instead of as bare alerts.
- **Monthly narrative.** The month's money, what the couple achieved together, and impressions, written as a short narrative.
- **Advice and encouragement on request.**

Requirements:
- The AI never sends a message, approves a request, or changes a record on anyone's behalf. A person takes every action.
- It uses only what the member it is helping is already allowed to see. It never reveals, hints at, or paraphrases to one partner what the other kept private.
- It is a supportive companion, not a psychologist: no diagnosis, no therapy or medical claims. `security-privacy` defines, in document 10, what it does on signs of serious distress, self-harm, or violence, including pointing to professional help.
- Each member opts in separately and can turn it off. State plainly which data leaves the server for an AI provider, and send the minimum.
- The app is fully usable without it. An AI failure never blocks recording, requests, or approvals.
- Its wording is natural Indonesian, varied, and never generic or repetitive; it respects quiet hours.
- Usage cost has a hard cap.

Owner decisions:
1. **Timing.** The light version (morning greeting, suggested words of support, warm reminders, monthly narrative) is Tier 2 and part of the pilot. Open-ended advice conversations are Tier 3.
2. **Disclosure.** A drafted message carries no AI marker. Both partners know the feature exists, and the sender must read each message and send it personally.
3. **Data sent to the provider.** Only with the consent of the member the data is about, and only the minimum, never names:
   - Suggested words of support: the stated feeling or condition, without the member's free-text reason unless that member allows it.
   - Warm reminders: the title and due date of the task, request, or outing, without amounts.
   - The non-money part of the monthly narrative: counts and titles of completed tasks, requests, and outings.
4. **Money in the monthly narrative.** The AI receives aggregated figures only, never individual transactions.

`security-privacy` assesses decisions 2 to 4 at S4 against the research on UU PDP and recommends a change if the research calls for one. `product-analyst` recommends, in document 05 at S3, how to keep suggested affection from feeling fake; the owner decides at that checkpoint.

---

## 6. UNIVERSAL FAMILY ACTIVITY ENGINE

Evaluate whether modules should share reusable primitives so workflows connect instead of living in silos. Example chains:

- Issue reported → discussion → repair proposal → budget approval → payment → completion → household history.
- Meal plan → ingredient needs → inventory check → shopping list → budget approval → purchase → expense record → inventory update.
- Vacation idea → discussion → vote → plan approval → savings goal → budget allocation → booking records → expenses → completion summary.

Candidate primitives: Activity, Request, Proposal, Approval, Task, Assignment, Budget, Financial Transaction, Discussion, Attachment, Notification, Event, Timeline Entry, Audit Record.

Compare a unified generic workflow engine against explicit domain-specific workflows that share a small set of primitives. Do not force everything into one generic model if it harms clarity or financial correctness. Recommend one architecture and document it as an ADR.

---

## 7. FINANCIAL DATA CORRECTNESS

Financial accuracy is non-negotiable. Provide worked examples (tables showing balances before and after each step) for at least:

Spending from personal cash; spending from shared household cash; husband transfers money to wife; allocating money to mother-in-law; paying on behalf of another member; reimbursement; moving money between bank, e-wallet, and cash; full and partial refunds; recurring bills; split expenses; debt repayment; savings contributions; budget reservations; settling unused budget; cash reconciliation (counted cash ≠ recorded cash); incorrect transactions and corrections; concurrent updates; duplicate submissions; salary arrival and allocation as in §1.1; an installment payment; a loan to extended family and its repayment; one spouse borrowing back from money the other holds; an asset purchase and its effect on net worth.

Define each of these terms canonically in `00-glossary.md` and use them consistently: income, expense, transfer, allocation, authorization, reservation, liability, settlement, reimbursement, available balance, recorded balance.

Rules that must hold:
- A transfer between family-controlled accounts is **never** income or expense.
- Budget approval **does not** reduce actual cash unless money is genuinely transferred, reserved under an explicit rule, or spent.
- Corrections are made through reversing or adjusting entries, not by silently editing history.
- Reconcile the need for an immutable financial history with legitimate deletion rights (e.g., a member leaving or requesting data deletion).

---

## 8. PRIVACY, SECURITY & TRUST

The app may hold family relationships, finances, cash holdings, health status, private conversations, children's data, addresses, conflicts, and personal plans. Design privacy-first.

Investigate: authentication, biometrics, sessions, device management, account recovery, multi-household access control, fine-grained authorization, sensitive-data visibility, personal vs shared wallets, transaction integrity, audit trails, secure attachments, retention, export and deletion, encryption (in transit, at rest, and whether any data warrants end-to-end encryption), notification privacy, child-account protection, ownership transfer, household separation, abuse prevention, and emergency recovery.

Define behavior for:
- A couple separates.
- A member leaves the family.
- A device is lost or stolen.
- A child turns 18.
- The household admin loses access.
- Members disagree about a transaction.
- A member revokes previously granted access.
- The device is offline.
- Conflicting edits arrive from multiple devices.

**Abuse prevention is a first-class requirement.** Family software can be used for financial control or surveillance by one member over another. Define concrete safeguards (e.g., what an admin can never see, transparency about who can see what, ability for an adult member to keep personal funds private, safe exit from a household).

**Regulation.** Research the current Indonesian rules that apply — at minimum the Personal Data Protection Law (UU PDP) and any regulation on children's data and online child protection — plus Apple App Store and Google Play requirements (account deletion, data safety disclosure, children's apps). Cite sources with access dates and flag what needs legal review. You are not giving legal advice.

Family membership never grants unlimited access to another adult's private data.

**AI data flows.** For the companion in §5O, define what is sent to an AI provider, under whose consent, how long the provider keeps it, and how one member's private data is kept out of anything generated for another. Research how UU PDP treats sending mood and health data to a provider outside Indonesia, and flag it for legal review.

---

## 9. TECHNICAL ARCHITECTURE & TECHNOLOGY SELECTION

I am deliberately not prescribing the stack. Recommend it based on the product requirements **and** the constraints in §1 (team size, team skills, budget, timeline).

### 9.1 Mobile
Compare at least Flutter, React Native (incl. Expo), Kotlin Multiplatform (incl. Compose Multiplatform), and native Swift + Kotlin. Criteria: performance, iOS and Android quality, tablet and adaptive UI, native capabilities, accessibility, offline storage, maintainability, productivity for a small team, testing, ecosystem, long-term support, platform limits, total effort.

Weigh **native look and feel** (§10.1) above code sharing. A framework that draws its own widgets must prove it can match current system components on both platforms, and keep up with yearly OS design changes, or be rejected. Include the option of fully native UI on each platform with logic kept on the server or in a shared core.

### 9.2 Backend
Evaluate realistic options given the team's Go/Rust strength, as well as managed backends (e.g., Supabase, Firebase) where they reduce effort. Criteria: API performance, realtime, transactional integrity, background jobs, security, concurrency, productivity, operational burden, scalability, maintainability, cost.

### 9.3 Data
Recommend storage for financial records, accounts, relationships, chat, notifications, attachments, inventory, plans, audit history, and search. Address relational modeling, transactional consistency, isolation levels, idempotency, concurrency control, multi-tenancy (household isolation), and schema evolution.

### 9.4 Architectural decisions
Cover: modular monolith vs microservices (do not assume microservices are better), realtime transport, push notifications, offline-first vs online-first (decide per module — finance and shopping lists may need different answers), sync and conflict resolution, background and scheduled jobs, file storage, search indexing, event-driven workflows, ledger modeling, observability, deployment, CI/CD, automated testing, backup and disaster recovery, cost optimization.

Provide comparison matrices, state risks, **recommend one primary architecture**, and record each major decision as an ADR (`docs/adr/ADR-NNN-title.md`: context, options, decision, consequences). Explicitly list which decisions must be made now and which can be deferred.

Use web research for anything version-, pricing-, or capability-dependent, and cite sources with dates.

### 9.5 AI
For the companion in §5O, compare a hosted AI provider against a self-hosted model on Indonesian-language quality, privacy, cost on the pilot budget, and operational burden. Recommend one, with a cost cap and a fallback when the AI is unavailable. All AI calls go through the server, never directly from the app.

---

## 10. UX/UI STRATEGY

### 10.1 Visual identity
Warm, friendly, trustworthy, premium, calm, modern, clean, universal, minimal but not empty. Avoid corporate coldness, heavy gradients, gratuitous animation, and complex dashboards.

**Platform-native is a hard requirement.** On iPhone the app must feel fully like an iPhone app; on Android, fully like an Android app. Each platform follows its own current design language (Apple Human Interface Guidelines and Liquid Glass; Material 3 Expressive), its own navigation patterns, system components, gestures, typography, haptics, and motion. The two apps share brand color, iconography, wording, and information architecture, but not pixel-identical screens. Research the current state of both design languages before proposing the design system, and cite sources with dates. Minimal, clean, light, and easy to use take priority over decoration.

### 10.2 Themes
Light and dark are both designed on every platform. Android offers light, dark, and follow system inside the app; iPhone and iPad follow the system setting, as Apple's guidelines advise (owner decision, S4 checkpoint). Dark mode is designed, not inverted. Ensure contrast, readable typography, and consistent components in every theme.

### 10.3 Devices
Small and large Android phones, compact and large iPhones, Android tablets, iPad, large tablets, foldables (e.g., Samsung Galaxy Z Fold, and any foldable iPhone) including folded, unfolded, and mid-transition states; varied aspect ratios; portrait and landscape where appropriate; dynamic type; cutouts and safe areas. A tablet layout must not be a stretched phone layout. Explore adaptive navigation (bottom bar, navigation rail), split view, master–detail, responsive grids, adaptive modals, multitasking, and touch targets.

### 10.4 Questions to answer
- What does a user see immediately on opening the app, and does it differ by persona?
- How do complex finances stay understandable?
- How do requests and approvals appear and get resolved in one or two taps?
- How do the most common actions take very few steps?
- How do older members navigate comfortably?
- How does a user switch households?
- How is sensitive information hidden (e.g., balance masking, private mode)?
- Which features appear contextually instead of in menus?
- How does a large feature set avoid an overwhelming menu?

### 10.5 Deliverables
Information architecture, navigation structure, screen inventory and hierarchy, primary journeys per persona, onboarding, dashboard strategy, interaction patterns, empty / error / offline / loading states, permission-request UX, accessibility, phone vs tablet behavior, design-system recommendations (tokens, components, typography, color), and motion principles. Describe screens in words and diagrams; do not produce UI code.

---

## 11. ORIGINAL FEATURE INNOVATION (MANDATORY)

Propose **at least 15** feature concepts beyond this brief, covering areas such as financial transparency, mental load, responsibility sharing, parenting, elderly support, automation, predictive reminders, emotional wellbeing, collaborative decisions, safety, shared resources, and long-term planning. Consider Indonesian family realities (e.g., arisan, uang belanja, supporting parents, mudik, zakat/sedekah, school fee cycles, gotong royong) while keeping the product culturally adaptable.

For each: (1) user problem, (2) feature, (3) realistic scenario, (4) expected value, (5) disadvantages, (6) privacy implications, (7) complexity, (8) priority.

Separate genuinely valuable ideas from ones that only impress in demos. Reject AI features that do not solve a real problem.

---

## 12. END-TO-END SCENARIOS

Write **at least 20** scenarios, including:

1. Husband requests money to buy food.
2. Wife buys groceries with allocated household cash.
3. Child asks permission to go out.
4. Mother-in-law receives money for groceries.
5. Family plans a vacation.
6. Parents save for school fees.
7. Household runs out of cooking oil.
8. Leaking roof needs urgent repair.
9. A member is sick and needs help.
10. A member shares their mood privately.
11. Recurring electricity bill becomes due.
12. Member pays with personal money and requests reimbursement.
13. Purchase costs less than the approved budget.
14. Purchase exceeds the approved budget.
15. A refund is issued.
16. A member belongs to multiple households.
17. Guardian supervises a child's financial requests.
18. Two members edit the same shopping list simultaneously.
19. Member loses internet while shopping.
20. Household admin changes or loses account access.

For each: trigger, actors, preconditions, user interaction, system behavior, approval requirements, financial effects (with balances), notifications, outcome, and failure/edge cases. Reference requirement IDs so scenarios trace back to `05-functional-requirements.md`.

---

## 13. PRIORITIZATION & PHASING

Do not put everything in the first release. Use this phase structure, or justify a better one:

- **Phase 0 — Discovery & Validation:** validate problems, users, and assumptions. Include a concrete validation plan (who to interview, how many families, which questions, which prototype tests, and what result would change the plan).
- **Phase 1 — MVP:** the smallest coherent product that delivers real daily value and that one family would actually keep using.
- **Phase 2 — Core Family Operations.**
- **Phase 3 — Advanced Household Management.**
- **Phase 4 — Intelligence & Automation** (optional, problem-driven AI beyond the §5O light version, which is part of the MVP).
- **Phase 5 — Ecosystem & Scale.**

Use one prioritization framework consistently (e.g., RICE scores plus MoSCoW labels). Classify every feature as Must / Should / Could / Future / **Will not build** (with reasons). Show dependencies and sequencing, preferably as a Mermaid diagram.

---

## 14. CRITICAL RISKS & SIMPLICITY

Analyze failure points such as: too many features, complex navigation, excessive manual entry, intrusive approvals, surveillance, financial inaccuracy, notification overload, low adoption by older members, **only one family member actively using the app**, difficult onboarding, unclear money ownership, over-engineering, integration dependencies, operating cost, unreliable offline sync.

For each: why it matters, likelihood, impact, prevention, mitigation. Propose concrete ways to reduce cognitive load and recurring effort. The app must save time, not become another chore.

---

## 15. MARKET & DIFFERENTIATION

Research comparable categories: family organizers, shared calendars, personal finance managers, shared-expense trackers, household task managers, family communication apps, shopping-list apps — including apps popular in Indonesia.

Identify existing approaches, common limitations, differentiation opportunities, what users actually need, what is unnecessarily complex, why families abandon such apps, and what drives long-term retention.

Do not invent competitor features or statistics. Cite sources with dates; label anything unverified. Conclude with the strongest positioning and a one-sentence value proposition.

---

## 16. DOCUMENTATION STRUCTURE & CONVENTIONS

```
docs/
├── README.md                          # entry point, reading order, links
├── _progress.md                       # stage log (see §0.2)
├── 00-glossary.md                     # single canonical definition of every shared term
├── 01-executive-summary.md            # written LAST
├── 02-product-principles.md
├── 03-personas-household-models.md
├── 04-feature-map.md
├── 05-functional-requirements.md
├── 06-activity-approval-architecture.md
├── 07-financial-domain.md
├── 08-ux-ui-strategy.md
├── 09-technical-architecture.md
├── 10-security-privacy.md
├── 11-innovation.md
├── 12-scenarios.md
├── 13-product-roadmap.md
├── 14-engineering-roadmap.md
├── 15-risk-assessment.md
├── 16-final-recommendations.md
├── 17-open-questions-assumptions.md
├── 18-market-research.md
└── adr/
    └── ADR-001-....md
```

Split a document into a subfolder (e.g., `05-functional-requirements/FIN.md`) if it grows beyond roughly 1,500 lines.

### Required content per deliverable

| # | Document | Must contain |
|---|---|---|
| 01 | Executive summary | Vision, core problem, audience, value proposition, differentiation, recommended MVP in one page |
| 02 | Principles | Design, UX, privacy, and financial-integrity principles — each testable, not slogans |
| 03 | Personas & household models | Personas incl. an older member, a teenager, a helper/caregiver, a member in two households; household structures; core needs |
| 04 | Feature map | All modules, features, sub-features, relationships, dependencies |
| 05 | Functional requirements | Specs, business rules, permissions, workflows, acceptance criteria (Given/When/Then) |
| 06 | Activity & approval architecture | Workflow models, state machines, approval logic, settlement |
| 07 | Financial domain | Ledger concepts, balances, custody, budgeting, transactions, reporting, worked examples |
| 08 | UX/UI strategy | Everything in §10.5 |
| 09 | Technical architecture | Comparisons, recommended stack, rationale, components, data, realtime, offline |
| 10 | Security & privacy | Threat model, access control, sensitive data, child safety, abuse prevention, data lifecycle, regulation |
| 11 | Innovation | ≥ 15 concepts with assessment and priority |
| 12 | Scenarios | ≥ 20 scenarios with edge cases |
| 13 | Product roadmap | Phases, validation plan, milestones, dependencies, and the business-model recommendation asked for in §1 |
| 14 | Engineering roadmap | Workstreams, testing strategy (incl. financial invariant tests), quality gates, release strategy, operations |
| 15 | Risk assessment | Product, UX, security, financial, technical, business risks with mitigations |
| 16 | Final recommendations | MVP, architecture, critical decisions, deferred decisions, features to reconsider, top opportunities |
| 17 | Open questions & assumptions | Owner decisions needed, assumptions, research to validate, proposed defaults with justification |
| 18 | Market research | Everything in §15, plus every other research finding (regulation, integrations, platform design, technology), one section per topic, with sources and access dates. Other documents cite it. A finding another agent verifies itself is cited in that agent's own document and reported to the lead, who has the researcher add it here |

### Identifier scheme

| Type | Format | Example |
|---|---|---|
| Feature | `F-<MOD>-NNN` | `F-FIN-004` |
| Functional requirement | `FR-<MOD>-NNN` | `FR-REQ-012` |
| Business rule | `BR-<MOD>-NNN` | `BR-FIN-003` |
| Scenario | `SCN-NN` | `SCN-13` |
| Innovation concept | `IDEA-NN` | `IDEA-07` |
| Risk | `RISK-NN` | `RISK-05` |
| Decision record | `ADR-NNN` | `ADR-002` |

Maintain traceability: every Must-have requirement (all of Tier 1, plus any Tier 2 requirement labeled Must) maps to a feature and at least one scenario from S3, and to a roadmap phase from S5. Include a traceability table in `05-functional-requirements.md` or `13-product-roadmap.md`.

### Evidence labels

Tag every non-trivial claim. Where this brief says "unverified", use **[Perlu validasi]** and state what was searched and why it could not be confirmed:

- **[Terverifikasi]** — confirmed through research; include source and access date.
- **[Asumsi]** — reasonable belief, not verified.
- **[Usulan]** — your proposal or recommendation.
- **[Perlu validasi]** — must be checked with users, legal, or a partner before relying on it.

### Formatting

Use tables for comparisons, Mermaid for state machines, flows, and dependencies, and short paragraphs. Keep one canonical definition per concept (in the glossary) and link to it instead of redefining it.

---

## 17. QUALITY BAR

Before finalizing each major recommendation: understand the problem, consider alternatives, state tradeoffs, check dependencies and edge cases, challenge the assumption, assess complexity and long-term consequences, then choose the most defensible option.

- Comprehensive is not the same as complex. The best product is not the one with the most features.
- Avoid generic advice that could apply to any mobile app; tie every important recommendation to this Family OS and its users.
- Do not agree with my ideas by default. If something I asked for is a bad idea, say so plainly and propose a better alternative.

**Plan deeply. Question assumptions. Design for real families. Recommend decisively. Build nothing yet.**
