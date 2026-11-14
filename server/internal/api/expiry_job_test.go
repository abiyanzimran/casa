package api_test

import (
	"errors"
	"log/slog"
	"testing"
	"time"

	"connectrpc.com/connect/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/household"
	"github.com/abiyanzimran/casa/server/internal/jobs"
)

// runner is the jobs runner of casa serve on the harness's database, with
// the harness's clock for both the claim and the household service (plan
// 1.7 decision 1: one source of time).
func (h *householdHarness) runner(t *testing.T) *jobs.Runner {
	t.Helper()
	return jobs.NewRunner(jobs.Config{
		Claim: h.db.Pool(t, "job_runner"), App: h.pool, Now: h.clock.Now, Logger: slog.New(slog.DiscardHandler),
		Handlers: map[string]jobs.Handler{household.ExpireInvitationKind: h.service.ExpireInvitationJob},
	})
}

// runDue runs the due jobs once, as the runner's loop does each second.
func (h *householdHarness) runDue(t *testing.T) int {
	t.Helper()
	ran, err := h.runner(t).RunDue(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	return ran
}

// expiryJob is the one job row of invitation: everything the runner hands
// its handler, plus its run_at, unique key, attempts, lock and last error.
type expiryJob struct {
	job         jobs.Job
	runAt       time.Time
	uniqueKey   string
	attempts    int32
	lockedUntil pgtype.Timestamptz
	lastError   pgtype.Text
}

// jobOf reads the job whose unique key is invitationID; found is false
// when there is none.
func (h *householdHarness) jobOf(t *testing.T, invitationID string) (j expiryJob, found bool) {
	t.Helper()
	err := h.admin.QueryRow(t.Context(), `
		SELECT id, household_id, kind, payload, run_at, unique_key, attempts, locked_until, last_error
		  FROM job WHERE unique_key = $1`, invitationID).
		Scan(&j.job.ID, &j.job.HouseholdID, &j.job.Kind, &j.job.Payload, &j.runAt, &j.uniqueKey,
			&j.attempts, &j.lockedUntil, &j.lastError)
	if errors.Is(err, pgx.ErrNoRows) {
		return j, false
	}
	if err != nil {
		t.Fatal(err)
	}
	j.job.Attempt = j.attempts
	return j, true
}

func (h *householdHarness) mustJobOf(t *testing.T, invitationID string) expiryJob {
	t.Helper()
	j, found := h.jobOf(t, invitationID)
	if !found {
		t.Fatalf("no job for invitation %s", invitationID)
	}
	return j
}

// expiryEvent is the expiry event of an invitation from from at expiresAt,
// as invitationEvents prints it (06 §3.6 "Gugur karena waktu").
func expiryEvent(seq, from string, expiresAt time.Time) string {
	return seq + " system - invitation.expired expire_invitation " + from + " expired - " + utc(expiresAt)
}

// 06 §3.6 "Gugur karena waktu" with X3; 09.2 §5.1 note X3; 09.3 §9 note of
// 14 November 2026 (architect Q1 (e)); 10.1 §10 note 1.7 Q3; plan 1.7
// decision 1.
func TestInvitationExpiryJob(t *testing.T) {
	t.Run("F01 schedules the expiry at expires_at", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		if n := h.count(t, "SELECT count(*) FROM job"); n != 1 {
			t.Fatalf("%d jobs, want 1", n)
		}
		j := h.mustJobOf(t, inv.GetInvitationId())
		id := uuidOf(t, inv.GetInvitationId())
		if j.job.Kind != "household.expire_invitation" || j.job.HouseholdID != uuidOf(t, hh.GetHouseholdId()) ||
			!j.runAt.Equal(inv.GetExpiresAt().AsTime()) || j.uniqueKey != inv.GetInvitationId() ||
			string(j.job.Payload) != string(id.Bytes[:]) || j.attempts != 0 || j.lockedUntil.Valid || j.lastError.Valid {
			t.Fatalf("job = %+v, want kind household.expire_invitation in the household at %v keyed and carrying the invitation id",
				j, inv.GetExpiresAt().AsTime())
		}
	})

	t.Run("a refused F01 schedules nothing", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		h.addMember(t, hh.GetHouseholdId(), uuid4())
		_, err := h.invite(a, hh.GetHouseholdId())
		wantCodes(t, err, connect.CodeFailedPrecondition, householdFull)
		if n := h.count(t, "SELECT count(*) FROM job"); n != 0 {
			t.Fatalf("%d jobs, want 0", n)
		}
	})

	t.Run("at expires_at the job writes the expiry at expires_at", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		expiresAt := inv.GetExpiresAt().AsTime()
		h.clock.advance(invitationLifetime - time.Second)
		if ran := h.runDue(t); ran != 0 {
			t.Fatalf("ran %d jobs 1 s before expires_at, want 0", ran)
		}
		h.clock.advance(time.Second)
		if ran := h.runDue(t); ran != 1 {
			t.Fatalf("ran %d jobs at expires_at, want 1", ran)
		}
		if got, want := h.invitationRow(t, inv.GetInvitationId()), "expired 1 "+utc(expiresAt)+" "; got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		events := h.invitationEvents(t, inv.GetInvitationId())
		if len(events) != 2 || events[1] != expiryEvent("1", "open", expiresAt) {
			t.Fatalf("events = %q, want the creation and %q", events, expiryEvent("1", "open", expiresAt))
		}
		if _, found := h.jobOf(t, inv.GetInvitationId()); found {
			t.Fatal("the job is still there after it ran")
		}
	})

	t.Run("INV-A48 (b) an accepted invitation past its expiry is expired by the job, never confirmed", func(t *testing.T) {
		h := startHousehold(t)
		inv, b := h.acceptedBy(t, "a@example.test", "b@example.test")
		h.clock.advance(invitationLifetime)
		if ran := h.runDue(t); ran != 1 {
			t.Fatalf("ran %d jobs, want 1", ran)
		}
		if got, want := h.invitationRow(t, inv.id), "expired 2 "+utc(inv.expiresAt)+" "; got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		events := h.invitationEvents(t, inv.id)
		if got := events[len(events)-1]; got != expiryEvent("2", "accepted", inv.expiresAt) {
			t.Fatalf("last event = %q, want %q", got, expiryEvent("2", "accepted", inv.expiresAt))
		}
		// No system transition makes a membership (INV-A48 (b)), and F04
		// after it finds the invitation closed.
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b.account); n != 0 {
			t.Fatalf("%d memberships of the invitee, want 0", n)
		}
		h.signInAgain(t, &inv.inviter, "a@example.test")
		_, err := h.confirm(inv.inviter, confirmRequest(inv.household, inv.id))
		wantCodes(t, err, connect.CodeFailedPrecondition, stateChanged)
		if n := h.count(t, "SELECT count(*) FROM membership WHERE account_id = $1", b.account); n != 0 {
			t.Fatalf("%d memberships of the invitee after F04, want 0", n)
		}
	})

	t.Run("the job on a cancelled invitation writes nothing and is done", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		if err := h.cancelInvitation(a, hh.GetHouseholdId(), inv.GetInvitationId()); err != nil {
			t.Fatal(err)
		}
		row, events := h.invitationRow(t, inv.GetInvitationId()), h.invitationEvents(t, inv.GetInvitationId())
		h.clock.advance(invitationLifetime)
		if ran := h.runDue(t); ran != 1 {
			t.Fatalf("ran %d jobs, want 1", ran)
		}
		if got := h.invitationRow(t, inv.GetInvitationId()); got != row {
			t.Fatalf("invitation = %q, want it unchanged %q", got, row)
		}
		if got := h.invitationEvents(t, inv.GetInvitationId()); len(got) != len(events) {
			t.Fatalf("events = %q, want unchanged %q", got, events)
		}
		if _, found := h.jobOf(t, inv.GetInvitationId()); found {
			t.Fatal("the job is still there after it ran")
		}
	})

	// 09.3 §9: every handler is idempotent; a re-run after a commit whose
	// job deletion failed writes nothing more (plan 1.7 review focus 3).
	t.Run("a re-run of the expiry writes nothing more", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		j := h.mustJobOf(t, inv.GetInvitationId())
		h.clock.advance(invitationLifetime)
		for range 2 {
			if err := h.service.ExpireInvitationJob(t.Context(), j.job); err != nil {
				t.Fatalf("ExpireInvitationJob = %v, want done", err)
			}
		}
		if n := h.count(t, "SELECT count(*) FROM object_event WHERE rule = 'invitation.expired'"); n != 1 {
			t.Fatalf("%d expiry events, want 1", n)
		}
	})

	// 09.3 §9 note of 14 November 2026 (2): a live invitation the
	// compare-and-set finds not yet due is never done without its expiry
	// (ruling P3: run_at is lowered by a fixture to reach that case).
	t.Run("a live invitation the clock finds not yet due is retried, not done", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		expiresAt := inv.GetExpiresAt().AsTime()
		j := h.mustJobOf(t, inv.GetInvitationId())
		h.clock.advance(invitationLifetime - time.Second)
		if err := h.service.ExpireInvitationJob(t.Context(), j.job); !errors.Is(err, household.ErrNotYetDue) {
			t.Fatalf("ExpireInvitationJob 1 s early = %v, want ErrNotYetDue", err)
		}
		if _, err := h.admin.Exec(t.Context(), "UPDATE job SET run_at = $2 WHERE unique_key = $1",
			inv.GetInvitationId(), h.clock.Now()); err != nil {
			t.Fatal(err)
		}
		if ran := h.runDue(t); ran != 1 {
			t.Fatalf("ran %d jobs, want 1", ran)
		}
		retry := h.clock.Now().Add(10 * time.Second) // the first backoff (plan 1.7 decision 2)
		j, found := h.jobOf(t, inv.GetInvitationId())
		if !found || j.attempts != 1 || j.lastError.String != "handler" || !j.lockedUntil.Time.Equal(retry) {
			t.Fatalf("job = %+v (found %v), want attempt 1, last_error handler, locked until %v", j, found, retry)
		}
		if got := h.invitationRow(t, inv.GetInvitationId()); got != "open 0 - " {
			t.Fatalf("invitation = %q, want still open", got)
		}
		// The retry, after its backoff, writes the expiry at expires_at.
		h.clock.advance(10 * time.Second)
		if ran := h.runDue(t); ran != 1 {
			t.Fatalf("ran %d jobs on the retry, want 1", ran)
		}
		if got, want := h.invitationRow(t, inv.GetInvitationId()), "expired 1 "+utc(expiresAt)+" "; got != want {
			t.Fatalf("invitation = %q, want %q", got, want)
		}
		if _, found := h.jobOf(t, inv.GetInvitationId()); found {
			t.Fatal("the job is still there after the expiry")
		}
	})

	// Plan 1.7 Task 2: a payload that is not an invitation id is an error,
	// so the job stops after its retries and flags the operator.
	t.Run("a payload that is not an invitation id is an error", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		j := h.mustJobOf(t, h.mustInvite(t, a, hh.GetHouseholdId()).GetInvitationId())
		h.clock.advance(invitationLifetime)
		for _, payload := range [][]byte{nil, j.job.Payload[:15], append(j.job.Payload, 0)} {
			j.job.Payload = payload
			if err := h.service.ExpireInvitationJob(t.Context(), j.job); err == nil || errors.Is(err, household.ErrNotYetDue) {
				t.Fatalf("payload of %d bytes: %v, want an error", len(payload), err)
			}
		}
		if n := h.count(t, "SELECT count(*) FROM object_event WHERE rule = 'invitation.expired'"); n != 0 {
			t.Fatalf("%d expiry events, want 0", n)
		}
	})

	// 09.3 §9 note of 14 November 2026 (1); SEC-28 note Q7 (b).
	t.Run("issuing a link keeps expires_at and the job's run_at", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		hh := h.mustCreate(t, a)
		inv := h.mustInvite(t, a, hh.GetHouseholdId())
		expiresAt := inv.GetExpiresAt().AsTime()
		h.clock.advance(time.Hour)
		h.signInAgain(t, &a, "a@example.test")
		link := h.mustIssueLink(t, a, hh.GetHouseholdId(), inv.GetInvitationId())
		h.mustIssueLink(t, a, hh.GetHouseholdId(), inv.GetInvitationId())
		if !link.GetExpiresAt().AsTime().Equal(expiresAt) {
			t.Fatalf("link expires %v, want %v", link.GetExpiresAt().AsTime(), expiresAt)
		}
		var stored time.Time
		if err := h.admin.QueryRow(t.Context(), "SELECT expires_at FROM invitation WHERE id = $1",
			inv.GetInvitationId()).Scan(&stored); err != nil {
			t.Fatal(err)
		}
		if j := h.mustJobOf(t, inv.GetInvitationId()); !stored.Equal(expiresAt) || !j.runAt.Equal(expiresAt) ||
			h.count(t, "SELECT count(*) FROM job") != 1 {
			t.Fatalf("expires_at %v, run_at %v; want both %v and one job", stored, j.runAt, expiresAt)
		}
	})

	// 10.1 §10 note 1.7 Q3 condition 1: the job and the X3 lazy path call
	// the same writer, so the events differ only in household, object id
	// and seq. The Sync half of the comparison comes with Sync (Task 6).
	t.Run("the job and the lazy path write the same expiry", func(t *testing.T) {
		h := startHousehold(t)
		a := h.signUpMember(t, "a@example.test")
		byJob := h.mustInvite(t, a, h.mustCreate(t, a).GetHouseholdId())
		c := h.signUpMember(t, "c@example.test")
		lazyHousehold := h.mustCreate(t, c).GetHouseholdId()
		byLazyPath := h.mustInvite(t, c, lazyHousehold)
		h.clock.advance(invitationLifetime + 3*time.Hour)
		h.signInAgain(t, &c, "c@example.test")
		h.mustInvite(t, c, lazyHousehold) // X3: F01 writes the due expiry first
		if ran := h.runDue(t); ran != 2 {
			t.Fatalf("ran %d jobs, want 2 (one writes, one finds the invitation closed)", ran)
		}
		event := func(invitationID string) string {
			t.Helper()
			var e string
			if err := h.admin.QueryRow(t.Context(), `
				SELECT (to_jsonb(e) - 'household_id' - 'object_id' - 'seq')::text FROM object_event e
				 WHERE object_id = $1 AND to_state = 'expired'`, invitationID).Scan(&e); err != nil {
				t.Fatal(err) // also fails on no expiry, or more than one
			}
			return e
		}
		if byJobEvent, lazyEvent := event(byJob.GetInvitationId()), event(byLazyPath.GetInvitationId()); byJobEvent != lazyEvent {
			t.Fatalf("expiry by the job %s, by the lazy path %s; want equal", byJobEvent, lazyEvent)
		}
		if got, want := h.invitationRow(t, byJob.GetInvitationId()), h.invitationRow(t, byLazyPath.GetInvitationId()); got != want {
			t.Fatalf("invitation expired by the job %q, by the lazy path %q; want equal", got, want)
		}
		if n := h.count(t, "SELECT count(*) FROM job"); n != 1 {
			t.Fatalf("%d jobs, want 1: the new invitation's", n)
		}
	})
}
