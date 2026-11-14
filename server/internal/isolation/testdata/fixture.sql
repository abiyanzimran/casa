-- The isolation test's fixture (09.2 §3 layer 6; SEC-02): households X and
-- Y, each with a row in every table of public but currency, run as the
-- superuser. Every account, token, event, feed row, job and idempotency
-- record belongs to one of them. Identifiers of X begin with 1, of Y with 2;
-- token hashes are one byte repeated (0x1? for X, 0x2? for Y), as in
-- isolation_test.go. The clock is 2026-11-14 12:00 UTC: every token and
-- invitation is live, and both jobs are due, X's first. X's rows mirror Y's
-- (both invitations open at version 0 with a token), so a query that skipped
-- a policy would find X's row; their times (minutes and seconds 11:11) and
-- bytes differ from Y's, so a result holding one is caught. Household X2, on
-- X's side, holds the invitation X's invitee (account ...0012) accepted,
-- which invitation_invitee keys across households. Both household rules are
-- version 2, so InsertFirstHouseholdRule (version 1) can succeed.

INSERT INTO account (id, email, created_at, signup_via) VALUES
	('10000000-0000-4000-8000-000000000011', 'owner-x@example.test', '2026-11-14 11:11:11+00', 'setup_token'),
	('10000000-0000-4000-8000-000000000012', 'invitee-x@example.test', '2026-11-14 11:11:11+00', 'invitation'),
	('20000000-0000-4000-8000-000000000011', 'owner-y@example.test', '2026-11-14 12:00:00+00', 'setup_token');

INSERT INTO credential (account_id, salt, hash, time_cost, memory_kib, parallelism, updated_at) VALUES
	('10000000-0000-4000-8000-000000000011', decode(repeat('42', 32), 'hex'), decode(repeat('43', 32), 'hex'), 3, 65536, 1, '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000011', decode(repeat('40', 32), 'hex'), decode(repeat('41', 32), 'hex'), 3, 65536, 1, '2026-11-14 12:00:00+00');

INSERT INTO session (id, account_id, client_kind, device_mode, ua_summary, token_hash, client_version, created_at,
                     last_used_at, idle_timeout, absolute_expiry, token_issued_at) VALUES
	('10000000-0000-4000-8000-000000000031', '10000000-0000-4000-8000-000000000011', 'web', 'personal', 'Firefox',
	 decode(repeat('11', 32), 'hex'), '1.0.0', '2026-11-14 11:11:11+00', '2026-11-14 11:11:11+00', '30 days',
	 '2027-02-12 12:11:11+00', '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000031', '20000000-0000-4000-8000-000000000011', 'web', 'personal', 'Firefox',
	 decode(repeat('21', 32), 'hex'), '1.0.0', '2026-11-14 12:00:00+00', '2026-11-14 12:00:00+00', '30 days',
	 '2027-02-12 12:00:00+00', '2026-11-14 12:00:00+00');

INSERT INTO refresh_token (token_hash, session_id, account_id, created_at) VALUES
	(decode(repeat('12', 32), 'hex'), '10000000-0000-4000-8000-000000000031', '10000000-0000-4000-8000-000000000011', '2026-11-14 11:11:11+00'),
	(decode(repeat('22', 32), 'hex'), '20000000-0000-4000-8000-000000000031', '20000000-0000-4000-8000-000000000011', '2026-11-14 12:00:00+00');

INSERT INTO reauth_grant (id, session_id, account_id, purpose, created_at, expires_at, cookie_hash) VALUES
	('10000000-0000-4000-8000-000000000041', '10000000-0000-4000-8000-000000000031', '10000000-0000-4000-8000-000000000011',
	 'change_password', '2026-11-14 11:11:11+00', '2026-11-14 12:11:11+00', decode(repeat('17', 32), 'hex')),
	('20000000-0000-4000-8000-000000000041', '20000000-0000-4000-8000-000000000031', '20000000-0000-4000-8000-000000000011',
	 'change_password', '2026-11-14 12:00:00+00', '2026-11-14 12:05:00+00', decode(repeat('27', 32), 'hex'));

INSERT INTO recovery_token (token_hash, account_id, created_at, expires_at) VALUES
	(decode(repeat('13', 32), 'hex'), '10000000-0000-4000-8000-000000000011', '2026-11-14 11:11:11+00', '2026-11-14 13:11:11+00'),
	(decode(repeat('23', 32), 'hex'), '20000000-0000-4000-8000-000000000011', '2026-11-14 12:00:00+00', '2026-11-14 13:00:00+00');

-- Unused setup tokens, owned through their hash (isolation_test.go ownership).
INSERT INTO setup_token (token_hash, email, created_at, expires_at) VALUES
	(decode(repeat('14', 32), 'hex'), 'owner-x@example.test', '2026-11-14 11:11:11+00', '2026-11-21 12:11:11+00'),
	(decode(repeat('24', 32), 'hex'), 'owner-y@example.test', '2026-11-14 12:00:00+00', '2026-11-21 12:00:00+00');

-- The setup-token sign-ups are unused; the invitation sign-ups are used and
-- name their owner's account, which reads them in its own scope
-- (GetSignupInvitation, GetSignupVerification).
INSERT INTO signup_verification (id, token_hash, source, setup_token_hash, invitation_token_hash, email, created_at,
                                 expires_at, used_at, account_id) VALUES
	('10000000-0000-4000-8000-000000000051', decode(repeat('15', 32), 'hex'), 'setup_token', decode(repeat('14', 32), 'hex'), NULL,
	 'owner-x@example.test', '2026-11-14 11:11:11+00', '2026-11-15 12:11:11+00', NULL, NULL),
	('10000000-0000-4000-8000-000000000052', decode(repeat('18', 32), 'hex'), 'invitation', NULL, decode(repeat('19', 32), 'hex'),
	 'owner-x@example.test', '2026-11-14 11:11:11+00', '2026-11-15 12:11:11+00', '2026-11-14 11:11:11+00',
	 '10000000-0000-4000-8000-000000000011'),
	('20000000-0000-4000-8000-000000000051', decode(repeat('25', 32), 'hex'), 'setup_token', decode(repeat('24', 32), 'hex'), NULL,
	 'owner-y@example.test', '2026-11-14 12:00:00+00', '2026-11-15 12:00:00+00', NULL, NULL),
	('20000000-0000-4000-8000-000000000052', decode(repeat('28', 32), 'hex'), 'invitation', NULL, decode(repeat('29', 32), 'hex'),
	 'owner-y@example.test', '2026-11-14 12:00:00+00', '2026-11-15 12:00:00+00', '2026-11-14 12:00:00+00',
	 '20000000-0000-4000-8000-000000000011');

INSERT INTO security_event (id, account_id, session_id, kind, occurred_at) VALUES
	('10000000-0000-4000-8000-000000000061', '10000000-0000-4000-8000-000000000011', '10000000-0000-4000-8000-000000000031',
	 'sign_in', '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000061', '20000000-0000-4000-8000-000000000011', '20000000-0000-4000-8000-000000000031',
	 'sign_in', '2026-11-14 12:00:00+00');

INSERT INTO idempotency_record (scope_id, idempotency_key, command, request_hash, response) VALUES
	('10000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000081', 'CreateHousehold',
	 decode(repeat('34', 32), 'hex'), decode(repeat('35', 32), 'hex')),
	('20000000-0000-4000-8000-000000000001', '20000000-0000-4000-8000-000000000081', 'CreateHousehold',
	 decode(repeat('31', 32), 'hex'), decode(repeat('32', 32), 'hex'));

INSERT INTO household (id, name, time_zone, created_at) VALUES
	('10000000-0000-4000-8000-000000000001', 'Rumah X', 'Asia/Jakarta', '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000001', 'Rumah Y', 'Asia/Jakarta', '2026-11-14 12:00:00+00'),
	('10000000-0000-4000-8000-000000000002', 'Rumah X2', 'Asia/Jakarta', '2026-11-14 11:11:11+00');

INSERT INTO membership (household_id, member_id, account_id, display_name, role, permission_set, joined_at) VALUES
	('10000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000021', '10000000-0000-4000-8000-000000000011',
	 'Pemilik X', 'household_owner', 'two_equal_adults', '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000001', '20000000-0000-4000-8000-000000000021', '20000000-0000-4000-8000-000000000011',
	 'Pemilik Y', 'household_owner', 'two_equal_adults', '2026-11-14 12:00:00+00'),
	('10000000-0000-4000-8000-000000000002', '10000000-0000-4000-8000-000000000022', NULL, NULL,
	 'household_owner', 'two_equal_adults', '2026-11-14 11:11:11+00');

INSERT INTO approval_authority (household_id, member_id, request_type, valid_from) VALUES
	('10000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000021', 'all', '2026-11-13 12:11:11+00'),
	('20000000-0000-4000-8000-000000000001', '20000000-0000-4000-8000-000000000021', 'all', '2026-11-13 12:00:00+00');

INSERT INTO household_rule (household_id, version, preset_id, created_at) VALUES
	('10000000-0000-4000-8000-000000000001', 2, 'two_equal_adults', '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000001', 2, 'two_equal_adults', '2026-11-14 12:00:00+00');

-- The pilot writes no relationship (FR-FAM-006/AC-3); a member to itself
-- gives the table its rows.
INSERT INTO relationship (household_id, member_id, other_member_id, kind) VALUES
	('10000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000021', '10000000-0000-4000-8000-000000000021', 'fixture'),
	('20000000-0000-4000-8000-000000000001', '20000000-0000-4000-8000-000000000021', '20000000-0000-4000-8000-000000000021', 'fixture');

INSERT INTO invitation (household_id, id, inviter, token_hash, state, invitee_account_id, invitee_display_name,
                        created_at, expires_at, version) VALUES
	('10000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000071', '10000000-0000-4000-8000-000000000021',
	 decode(repeat('16', 32), 'hex'), 'open', NULL, NULL, '2026-11-14 11:11:11+00', '2026-11-21 12:11:11+00', 0),
	('20000000-0000-4000-8000-000000000001', '20000000-0000-4000-8000-000000000071', '20000000-0000-4000-8000-000000000021',
	 decode(repeat('26', 32), 'hex'), 'open', NULL, NULL, '2026-11-14 12:00:00+00', '2026-11-21 12:00:00+00', 0),
	('10000000-0000-4000-8000-000000000002', '10000000-0000-4000-8000-000000000072', '10000000-0000-4000-8000-000000000022',
	 NULL, 'accepted', '10000000-0000-4000-8000-000000000012', 'Tamu X', '2026-11-14 11:11:11+00', '2026-11-21 12:11:11+00', 1);

INSERT INTO object_event (household_id, object_kind, object_id, seq, actor_kind, actor, command, to_state, server_time) VALUES
	('10000000-0000-4000-8000-000000000001', 'household', '10000000-0000-4000-8000-000000000001', 0, 'member',
	 '10000000-0000-4000-8000-000000000021', 'CreateHousehold', 'created', '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000001', 'household', '20000000-0000-4000-8000-000000000001', 0, 'member',
	 '20000000-0000-4000-8000-000000000021', 'CreateHousehold', 'created', '2026-11-14 12:00:00+00');

INSERT INTO consent_record (account_id, kind, text_version, recorded_at) VALUES
	('10000000-0000-4000-8000-000000000011', 'adult_age', 1, '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000011', 'adult_age', 1, '2026-11-14 12:00:00+00');

INSERT INTO operator_notice_ack (account_id, issue_seq, variant, version, text_sha256, issued_at) VALUES
	('10000000-0000-4000-8000-000000000011', 1, 'couple', 1, decode(repeat('30', 32), 'hex'), '2026-11-14 11:11:11+00'),
	('20000000-0000-4000-8000-000000000011', 1, 'couple', 1, decode(repeat('30', 32), 'hex'), '2026-11-14 12:00:00+00');

INSERT INTO job (id, household_id, kind, run_at, unique_key, payload) VALUES
	('10000000-0000-4000-8000-000000000091', '10000000-0000-4000-8000-000000000001', 'household.expire_invitation',
	 '2026-11-14 10:11:11+00', '10000000-0000-4000-8000-000000000071', decode('10000000000040008000000000000071', 'hex')),
	('20000000-0000-4000-8000-000000000091', '20000000-0000-4000-8000-000000000001', 'household.expire_invitation',
	 '2026-11-14 11:00:00+00', '20000000-0000-4000-8000-000000000071', decode('20000000000040008000000000000071', 'hex'));

INSERT INTO reader_feed (household_id, member_id, pos, object_kind, object_id, change) VALUES
	('10000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000021', 1, 'household',
	 '10000000-0000-4000-8000-000000000001', 'upsert'),
	('20000000-0000-4000-8000-000000000001', '20000000-0000-4000-8000-000000000021', 1, 'household',
	 '20000000-0000-4000-8000-000000000001', 'upsert');

INSERT INTO reader_feed_counter (household_id, member_id, last_pos) VALUES
	('10000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000021', 1),
	('20000000-0000-4000-8000-000000000001', '20000000-0000-4000-8000-000000000021', 1);
