-- name: InsertHousehold :exec
-- The caller makes the id before the write (gen_random_uuid, ADR-017) and
-- sets it as the transaction's household: the policy's WITH CHECK is
-- id = app.household_id (09.2 §3 note of 14 November 2026). currency,
-- income_source_market and version take their defaults: IDR, ID, 0 (09.2
-- §5.1; FR-FAM-001/AC-7).
INSERT INTO household (id, name, time_zone, created_at)
VALUES (@id, @name, @time_zone, @created_at);

-- name: NewID :one
-- ADR-017: ids are made by gen_random_uuid. A household's id is made before
-- its write, which sets it as the transaction's household.
SELECT gen_random_uuid()::uuid AS id;

-- name: InsertMembership :one
-- 06 §3.1 role and permission set; member_id is made here (ADR-017).
INSERT INTO membership (household_id, account_id, display_name, role, permission_set, joined_at)
VALUES (@household_id, @account_id, @display_name, 'household_owner', 'two_equal_adults', @joined_at)
RETURNING member_id;

-- name: InsertApprovalAuthority :exec
-- 06 §3.1: every request type, no amount limit (FR-FAM-001/AC-2).
INSERT INTO approval_authority (household_id, member_id, request_type, valid_from)
VALUES (@household_id, @member_id, 'all', @valid_from);

-- name: InsertFirstHouseholdRule :exec
-- 06 §3.1: version 1, preset two_equal_adults, no threshold
-- (FR-FAM-001/AC-2); agreement_mode takes its default 'all'.
INSERT INTO household_rule (household_id, version, preset_id, created_at)
VALUES (@household_id, 1, 'two_equal_adults', @created_at);

-- name: InsertObjectEvent :exec
INSERT INTO object_event (household_id, object_kind, object_id, seq, actor_kind, actor, rule,
                          command, from_state, to_state, idempotency_key, server_time)
VALUES (@household_id, @object_kind, @object_id, @seq, @actor_kind, @actor, @rule,
        @command, @from_state, @to_state, @idempotency_key, @server_time);

-- name: HasAgeDeclaration :one
-- SEC-80: a statement of any text version counts (it does not expire).
SELECT EXISTS (SELECT 1 FROM consent_record WHERE account_id = @account_id AND kind = 'adult_age');

-- name: InsertAgeDeclaration :exec
-- A concurrent command of the same account (CreateHousehold and F02) may
-- have written the same declaration first: that is the intended end state,
-- never an error (09.2 §5.1, note on decision 20).
INSERT INTO consent_record (account_id, kind, text_version, recorded_at)
VALUES (@account_id, 'adult_age', @text_version, @recorded_at)
ON CONFLICT (account_id, kind, text_version) DO NOTHING;

-- name: MembershipsOfAccount :many
-- The account's own memberships, in a lookup transaction (Scope{AccountID};
-- 09.2 §3 layer 1): only household_id and the row's id (plan 1.6 decision
-- 3). Served by membership_one_household; at most 20 [Usulan].
SELECT household_id, member_id FROM membership
 WHERE account_id = @account_id
 ORDER BY joined_at, household_id
 LIMIT 20;

-- name: HouseholdMembers :many
-- A household's name and members, in the household's own transaction; the
-- account filter checks the membership again (10.1 §5.1 note 1.6 item 1;
-- SEC-151). At most 20 members [Usulan].
SELECT h.name AS household_name, m.member_id, m.display_name
  FROM household h
  JOIN membership m ON m.household_id = h.id
 WHERE h.id = @household_id
   AND EXISTS (SELECT 1 FROM membership own WHERE own.household_id = h.id AND own.account_id = @account_id)
 ORDER BY m.joined_at, m.member_id
 LIMIT 20;

-- name: LastNoticeIssue :one
-- 09.3 §1 "Mekanisme": the version last issued to the account, by the
-- primary key (one row).
SELECT issue_seq, variant, version, text_sha256, (acknowledged_at IS NOT NULL)::boolean AS acknowledged
  FROM operator_notice_ack
 WHERE account_id = @account_id
 ORDER BY issue_seq DESC
 LIMIT 1;

-- name: IssueNotice :execrows
-- 09.3 §1 "Mekanisme": a new issue goes at the last issue_seq + 1; a
-- concurrent first sight conflicts on the key and writes nothing, so the
-- caller writes the security trail line once per issue.
INSERT INTO operator_notice_ack (account_id, issue_seq, variant, version, text_sha256, issued_at)
VALUES (@account_id, @issue_seq, @variant, @version, @text_sha256, @issued_at)
ON CONFLICT (account_id, issue_seq) DO NOTHING;

-- name: AcknowledgeNotice :execrows
-- SEC-34: only the last issue, only with its identity; the first
-- acknowledgement's time stays (FR-FAM-019/AC-2), so a repeat changes no row.
UPDATE operator_notice_ack o SET acknowledged_at = @acknowledged_at::timestamptz
 WHERE o.account_id = @account_id
   AND o.issue_seq = (SELECT max(last.issue_seq) FROM operator_notice_ack last WHERE last.account_id = o.account_id)
   AND o.variant = @variant AND o.version = @version AND o.text_sha256 = @text_sha256
   AND o.acknowledged_at IS NULL;

-- name: MemberOf :one
-- The account's membership of the selected household, in a lookup
-- transaction (Scope{AccountID}; 09.2 §3 layer 1; SEC-38 (a)): only
-- household_id and the row's id (plan 1.6 decision 3). Served by the
-- membership (household_id, account_id) unique index.
SELECT household_id, member_id FROM membership
 WHERE household_id = @household_id AND account_id = @account_id;

-- name: MemberIn :one
-- The account's membership of the transaction's household, checked again
-- inside a command (SEC-151; 09.2 §6 step 2). Served by the membership
-- (household_id, account_id) unique index.
SELECT household_id, member_id FROM membership
 WHERE household_id = current_setting('app.household_id')::uuid AND account_id = @account_id;

-- name: FullAdultMembers :many
-- SEC-10; 06 §3.2: the one definition of the full adult members, in the
-- transaction's household. In the pilot every member is a household owner
-- of two equal adults; F01 keeps a household at two adults (06 §3.6), so a
-- third row is an error the caller reports, never a longer list (ES §3).
-- Served by the membership primary key.
SELECT member_id FROM membership
 WHERE household_id = current_setting('app.household_id')::uuid AND role = 'household_owner'
 ORDER BY member_id
 LIMIT 3;

-- name: InsertInvitation :one
-- 06 §3.6 F01: open, version 0, no token until IssueInvitationLink (10.1
-- SEC-28 note Q7). invitation_live refuses a second live invitation.
INSERT INTO invitation (household_id, inviter, state, created_at, expires_at)
VALUES (@household_id, @inviter, 'open', @created_at, @expires_at)
RETURNING id;

-- name: LiveInvitation :one
-- The household's open or accepted invitation, if any: only household_id
-- and id (plan 1.6 decision 3). Served by invitation_live.
SELECT household_id, id FROM invitation
 WHERE household_id = @household_id AND state IN ('open', 'accepted');

-- name: Invitation :one
-- One invitation of the household, for a command's guard and its
-- compare-and-set (plan 1.6 decision 11). By the primary key.
SELECT inviter, state, expires_at, version FROM invitation
 WHERE household_id = @household_id AND id = @id;

-- name: TransitionInvitation :execrows
-- A state transition by compare-and-set on the version and state read
-- (plan 1.6 decision 11; 09.2 §9): zero rows means another command came
-- first. version moves on every transition and expires_at never changes,
-- so a guard on expires_at checked against the row read still holds here.
UPDATE invitation SET state = @to_state, closed_at = @closed_at::timestamptz, version = version + 1
 WHERE household_id = @household_id AND id = @id AND version = @version AND state = @from_state;

-- name: ReplaceInvitationToken :execrows
-- IssueInvitationLink: a new token hash replaces the old one, so the old
-- link stops working; not a transition, so no version and no event (plan
-- 1.6 decision 11). Only while the invitation is open as read.
UPDATE invitation SET token_hash = @token_hash
 WHERE household_id = @household_id AND id = @id AND version = @version AND state = 'open';
