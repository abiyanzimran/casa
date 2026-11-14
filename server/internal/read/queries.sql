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

-- name: ReaderMember :one
-- The reader's membership of the transaction's household, checked again
-- inside the read transaction (SEC-151; plan 1.7 P5: the read path's own
-- layer, household.MemberIn stays the command's). Served by
-- membership_one_household or the (household_id, account_id) unique index.
SELECT member_id FROM membership
 WHERE household_id = current_setting('app.household_id')::uuid AND account_id = @account_id;

-- name: LiveInvitationView :one
-- The household's open or accepted invitation as its members read it (plan
-- 1.6 decision 16), in the household's own transaction; the caller decides
-- what of the invitee it shows. Served by invitation_live.
SELECT id, inviter, state, expires_at, invitee_account_id, invitee_display_name FROM invitation
 WHERE household_id = current_setting('app.household_id')::uuid AND state IN ('open', 'accepted');

-- name: AccountEmail :one
-- 10.1 §5.1 note 1.6 item 3: only the email, by the primary key, for the
-- inviter's F04 screen (SEC-31).
SELECT email FROM account WHERE id = @id;

-- name: HouseholdMemberIDs :many
-- The household's members by id, in the command's transaction: the readers
-- and the member objects of the fan-out (plan 1.7 decisions 7, 8). By id, so
-- every command bumps the readers' counters in one order. Served by an
-- index leading with household_id (membership_pkey or the (household_id,
-- account_id) key). At most 20 [Usulan] (plan 1.6 "Batasan global"); a 21st
-- row is an error, never a reader left out (ES §3).
SELECT member_id FROM membership
 WHERE household_id = @household_id
 ORDER BY member_id
 LIMIT 21;

-- name: HouseholdView :one
-- The household object of sync (plan 1.7 decision 7), in the household's
-- own read transaction. Served by an index on id (household_pkey or the
-- (id, currency) key).
SELECT id, name, time_zone, currency FROM household
 WHERE id = current_setting('app.household_id')::uuid;

-- name: MemberViews :many
-- The member objects of sync (plan 1.7 decision 7) after a member id, by id:
-- a snapshot page (ADR-017 item 7) or, from the nil id, every member. The
-- caller bounds lim by its page or by maxMembers + 1 (ES §3). Served by an
-- index leading with household_id; the at most 20 rows are sorted in memory.
SELECT member_id, display_name FROM membership
 WHERE household_id = current_setting('app.household_id')::uuid AND member_id > @after
 ORDER BY member_id
 LIMIT @lim;

-- name: ProbeByLabel :one
-- Red proof only: a parameter the isolation test knows no value for.
SELECT display_name FROM membership
 WHERE household_id = current_setting('app.household_id')::uuid AND display_name = @probe_label
 LIMIT 1;
