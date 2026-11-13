-- The household tables (09.2 §5.1, note of 14 November 2026; plan 1.6
-- decisions 2, 3, 20): the five concepts of 06 §3.1 each in its own table,
-- invitations as the state machine of 06 §3.6, and the per-object events of
-- 06 §2.5. All are owned by the household module (09.2 §1).
--
-- There is no foreign key to account: it could not lead with household_id
-- (09.2 §3 layer 2) and would cross a module boundary. Every time is written
-- by the server from its clock; none has a default. Expand only: the old
-- binary runs on it unchanged (09.2 §10).

-- +goose Up
CREATE TABLE household (
	id                   uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name                 text NOT NULL,
	time_zone            text NOT NULL, -- IANA, the household's time (06 §2.4)
	currency             char(3) NOT NULL DEFAULT 'IDR' REFERENCES currency (code),
	income_source_market text NOT NULL DEFAULT 'ID' CHECK (income_source_market IN ('ID')), -- 07 §3.1
	version              bigint NOT NULL DEFAULT 0, -- compare-and-set (09.2 §9)
	created_at           timestamptz NOT NULL,
	UNIQUE (id, currency) -- target of the currency foreign keys (09.2 §7)
);

-- 06 §3.1 role and permission set. Until F06 is built a row means an active
-- membership; F06 adds state DEFAULT 'active' (expand).
CREATE TABLE membership (
	household_id   uuid NOT NULL REFERENCES household (id),
	member_id      uuid NOT NULL DEFAULT gen_random_uuid(), -- never cut from its rows
	account_id     uuid NULL, -- identity mapping (07 §8.2): cut = NULL
	display_name   text NULL, -- FR-FAM-006; cut with account_id
	role           text NOT NULL CHECK (role IN ('household_owner')),
	permission_set text NOT NULL CHECK (permission_set IN ('two_equal_adults')),
	joined_at      timestamptz NOT NULL,
	PRIMARY KEY (household_id, member_id),
	UNIQUE (household_id, account_id)
);
-- An account's own memberships across households (09.2 §3 layer 1).
CREATE INDEX membership_account ON membership (account_id) WHERE account_id IS NOT NULL;
-- The pilot's product rule "one household per account", held across
-- households because a unique check ignores row-level security; a recorded
-- exception to 09.2 §3 layer 2, dropped (contract) with F-FAM-009 (09.2
-- §5.1, note on decision 20).
CREATE UNIQUE INDEX membership_one_household ON membership (account_id) WHERE account_id IS NOT NULL;

-- 06 §3.1 approval authority. An amount always carries its currency; both
-- NULL means no limit. A composite foreign key with a NULL column is not
-- checked (MATCH SIMPLE), so a row without an amount never blocks a change
-- of the household's currency (09.2 §5.1, §7).
CREATE TABLE approval_authority (
	household_id uuid NOT NULL,
	member_id    uuid NOT NULL,
	request_type text NOT NULL CHECK (request_type IN ('all')),
	limit_amount bigint NULL,
	currency     char(3) NULL,
	valid_from   timestamptz NOT NULL,
	valid_until  timestamptz NULL,
	CHECK ((limit_amount IS NULL) = (currency IS NULL)),
	PRIMARY KEY (household_id, member_id, request_type, valid_from),
	FOREIGN KEY (household_id, member_id) REFERENCES membership,
	FOREIGN KEY (household_id, currency) REFERENCES household (id, currency)
);

-- 06 §3.1 household rule: never changed; a new version is a new row.
-- approval_mode is added with request (step 4, expand).
CREATE TABLE household_rule (
	household_id      uuid NOT NULL,
	version           int NOT NULL,
	preset_id         text NOT NULL CHECK (preset_id IN ('two_equal_adults')),
	request_threshold bigint NULL,
	currency          char(3) NULL,
	agreement_mode    text NOT NULL DEFAULT 'all' CHECK (agreement_mode IN ('all')),
	created_at        timestamptz NOT NULL,
	CHECK ((request_threshold IS NULL) = (currency IS NULL)),
	PRIMARY KEY (household_id, version),
	FOREIGN KEY (household_id, currency) REFERENCES household (id, currency)
);

-- 06 §3.1 relationship: the table exists, with no rows in the pilot
-- (FR-FAM-006/AC-3); no writer until Tier 3.
CREATE TABLE relationship (
	household_id    uuid NOT NULL,
	member_id       uuid NOT NULL,
	other_member_id uuid NOT NULL,
	kind            text NOT NULL,
	PRIMARY KEY (household_id, member_id, other_member_id),
	FOREIGN KEY (household_id, member_id) REFERENCES membership,
	FOREIGN KEY (household_id, other_member_id) REFERENCES membership
);

-- 06 §3.6 F01 to F05. open = undangan_terbuka, accepted =
-- menunggu_konfirmasi_pengundang, confirmed = F04 done; the other four end
-- the invitation with their cause.
CREATE TABLE invitation (
	household_id         uuid NOT NULL,
	id                   uuid NOT NULL DEFAULT gen_random_uuid(),
	inviter              uuid NOT NULL, -- member_id
	-- SHA-256 of the 256-bit link token: NULL from F01 until the first
	-- IssueInvitationLink; each reissue replaces it, F02 clears it
	-- (10.1 SEC-28, note Q7).
	token_hash           bytea NULL,
	state                text NOT NULL CHECK (state IN ('open', 'accepted', 'confirmed', 'cancelled', 'rejected', 'withdrawn', 'expired')),
	invitee_account_id   uuid NULL, -- identity mapping, as membership.account_id
	invitee_display_name text NULL,
	created_at           timestamptz NOT NULL,
	expires_at           timestamptz NOT NULL,
	closed_at            timestamptz NULL, -- expired: = expires_at, not the time written
	version              bigint NOT NULL DEFAULT 0,
	PRIMARY KEY (household_id, id),
	FOREIGN KEY (household_id, inviter) REFERENCES membership
);
-- F02 and PreviewInvitation find an invitation by its token.
CREATE INDEX invitation_token ON invitation (token_hash) WHERE token_hash IS NOT NULL;
-- One live invitation per household (06 §3.6 F01).
CREATE UNIQUE INDEX invitation_live ON invitation (household_id) WHERE state IN ('open', 'accepted');
-- F05 by the invitee (09.2 §3), and one open acceptance per account across
-- households: the second recorded exception to 09.2 §3 layer 2 (09.2 §5.1,
-- note on decision 20).
CREATE UNIQUE INDEX invitation_invitee ON invitation (invitee_account_id) WHERE state = 'accepted';

-- 06 §2.5 per-object events, INV-A01; append-only for app_rw (ADR-016 item
-- 5; FR-FAM-020/AC-6). seq is the object's version after the transition,
-- 0 for its creation event. terms_version (INV-A01) comes with request.
CREATE TABLE object_event (
	household_id    uuid NOT NULL,
	object_kind     text NOT NULL CHECK (object_kind IN ('household', 'invitation')), -- widened by each writing module
	object_id       uuid NOT NULL,
	seq             bigint NOT NULL,
	actor_kind      text NOT NULL CHECK (actor_kind IN ('member', 'invitee', 'system')),
	actor           uuid NULL, -- member_id
	rule            text NULL, -- system: the rule that fired, such as 'invitation.expired'
	command         text NOT NULL,
	from_state      text NULL,
	to_state        text NOT NULL,
	idempotency_key uuid NULL,
	server_time     timestamptz NOT NULL,
	CHECK ((actor_kind = 'member') = (actor IS NOT NULL)),
	CHECK ((actor_kind = 'system') = (rule IS NOT NULL)),
	CHECK (actor_kind <> 'invitee' OR object_kind = 'invitation'),
	PRIMARY KEY (household_id, object_kind, object_id, seq),
	FOREIGN KEY (household_id, actor) REFERENCES membership (household_id, member_id)
);

-- Every table forces row-level security (SEC-08; ADR-016 item 3). Layer 4
-- of 09.2 §3 reads app.household_id without missing_ok: a transaction that
-- sets no household fails instead of reading nothing. household,
-- membership and invitation have the policies of 09.2 §3 (note of 14
-- November 2026): a lookup clause lives only in a pure lookup transaction
-- (no household set) and only in USING, so a lookup never writes and a
-- transaction for household X never reads household Y (SEC-38 (b); 10.1
-- §5.1 note 1.6 item 1). app_rw never deletes; job_runner and lifecycle_rw
-- get nothing (plan 1.6 decision 2).
ALTER TABLE household ENABLE ROW LEVEL SECURITY;
ALTER TABLE household FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON household
	USING (id = current_setting('app.household_id')::uuid)
	WITH CHECK (id = current_setting('app.household_id')::uuid);
GRANT SELECT, INSERT, UPDATE (version) ON household TO app_rw;

ALTER TABLE membership ENABLE ROW LEVEL SECURITY;
ALTER TABLE membership FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON membership
	USING (household_id = nullif(current_setting('app.household_id', true), '')::uuid
	    OR (nullif(current_setting('app.household_id', true), '')::uuid IS NULL
	        AND account_id = nullif(current_setting('app.account_id', true), '')::uuid))
	WITH CHECK (household_id = nullif(current_setting('app.household_id', true), '')::uuid);
GRANT SELECT, INSERT ON membership TO app_rw;

ALTER TABLE approval_authority ENABLE ROW LEVEL SECURITY;
ALTER TABLE approval_authority FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON approval_authority
	USING (household_id = current_setting('app.household_id')::uuid)
	WITH CHECK (household_id = current_setting('app.household_id')::uuid);
GRANT SELECT, INSERT ON approval_authority TO app_rw;

ALTER TABLE household_rule ENABLE ROW LEVEL SECURITY;
ALTER TABLE household_rule FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON household_rule
	USING (household_id = current_setting('app.household_id')::uuid)
	WITH CHECK (household_id = current_setting('app.household_id')::uuid);
GRANT SELECT, INSERT ON household_rule TO app_rw;

ALTER TABLE relationship ENABLE ROW LEVEL SECURITY;
ALTER TABLE relationship FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON relationship
	USING (household_id = current_setting('app.household_id')::uuid)
	WITH CHECK (household_id = current_setting('app.household_id')::uuid);
GRANT SELECT ON relationship TO app_rw;

-- The invitee clause opens any invitation naming the account, whatever its
-- state (09.2 §3); the F05 lookup filters state = 'accepted'.
ALTER TABLE invitation ENABLE ROW LEVEL SECURITY;
ALTER TABLE invitation FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON invitation
	USING (household_id = nullif(current_setting('app.household_id', true), '')::uuid
	    OR (nullif(current_setting('app.household_id', true), '')::uuid IS NULL
	        AND token_hash = decode(nullif(current_setting('app.token_hash', true), ''), 'hex'))
	    OR (nullif(current_setting('app.household_id', true), '')::uuid IS NULL
	        AND invitee_account_id = nullif(current_setting('app.account_id', true), '')::uuid))
	WITH CHECK (household_id = nullif(current_setting('app.household_id', true), '')::uuid);
GRANT SELECT, INSERT, UPDATE (token_hash, state, invitee_account_id, invitee_display_name, closed_at, version)
	ON invitation TO app_rw;

ALTER TABLE object_event ENABLE ROW LEVEL SECURITY;
ALTER TABLE object_event FORCE ROW LEVEL SECURITY;
CREATE POLICY household ON object_event
	USING (household_id = current_setting('app.household_id')::uuid)
	WITH CHECK (household_id = current_setting('app.household_id')::uuid);
GRANT SELECT, INSERT ON object_event TO app_rw;

-- +goose Down
DROP TABLE object_event;
DROP TABLE invitation;
DROP TABLE relationship;
DROP TABLE household_rule;
DROP TABLE approval_authority;
DROP TABLE membership;
DROP TABLE household;
