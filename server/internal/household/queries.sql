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
