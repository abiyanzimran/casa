package identity

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type testClock struct{ now time.Time }

func (c *testClock) Now() time.Time          { return c.now }
func (c *testClock) advance(d time.Duration) { c.now = c.now.Add(d) }

// newTestLimiter returns a limiter whose full table evicts penalised keys;
// a test of the refusing policy sets refusesWhenFull.
func newTestLimiter(threshold int, clearsOnSuccess bool) (*Limiter, *testClock) {
	clock := &testClock{now: time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)}
	return NewLimiter(threshold, clearsOnSuccess, false, clock.Now), clock
}

// fail makes n attempts of key that each fail.
func fail(t *testing.T, l *Limiter, key string, n int) {
	t.Helper()
	for i := range n {
		allowed, done := l.Attempt(key)
		if !allowed {
			t.Fatalf("attempt %d of %s refused", i+1, key)
		}
		done(Failed)
	}
}

func allowed(l *Limiter, key string) bool {
	ok, done := l.Attempt(key)
	done(NotTried)
	return ok
}

func TestLimiter(t *testing.T) {
	t.Run("SEC-12 after 20 consecutive failures the next attempt waits", func(t *testing.T) {
		l, clock := newTestLimiter(20, true)
		fail(t, l, "k", 19)
		if !allowed(l, "k") {
			t.Fatal("refused after 19 failures")
		}
		fail(t, l, "k", 1)
		if allowed(l, "k") {
			t.Fatal("allowed right after 20 failures")
		}
		clock.advance(time.Second)
		fail(t, l, "k", 1) // the 21st waits 2 s
		clock.advance(1900 * time.Millisecond)
		if allowed(l, "k") {
			t.Fatal("allowed 1.9 s after the 21st failure")
		}
		clock.advance(100 * time.Millisecond)
		if !allowed(l, "k") {
			t.Fatal("refused 2 s after the 21st failure")
		}
	})
	t.Run("a refused attempt is not counted", func(t *testing.T) {
		l, clock := newTestLimiter(20, true)
		fail(t, l, "k", 20)
		for range 5 {
			if ok, done := l.Attempt("k"); ok {
				t.Fatal("allowed while waiting")
			} else {
				done(Failed) // a no-op for a refused attempt
			}
		}
		clock.advance(time.Second)
		if !allowed(l, "k") {
			t.Fatal("refused attempts lengthened the wait")
		}
	})
	t.Run("the wait stops growing at 15 minutes", func(t *testing.T) {
		l, clock := newTestLimiter(20, true)
		for range 100 {
			fail(t, l, "k", 1)
			clock.advance(15 * time.Minute)
		}
		fail(t, l, "k", 1)
		clock.advance(15*time.Minute - time.Nanosecond)
		if allowed(l, "k") {
			t.Fatal("allowed before 15 minutes")
		}
		clock.advance(time.Nanosecond)
		if !allowed(l, "k") {
			t.Fatal("refused after 15 minutes")
		}
	})
	t.Run("a success clears an account key", func(t *testing.T) {
		l, _ := newTestLimiter(20, true)
		fail(t, l, "k", 19)
		ok, done := l.Attempt("k")
		done(Succeeded)
		fail(t, l, "k", 19)
		if !ok || !allowed(l, "k") {
			t.Fatal("the failures before a success still count")
		}
	})
	t.Run("a success withdraws only its own count from the address key", func(t *testing.T) {
		// 10.1 §8.1: else one valid account resets the address every 19
		// tries at others.
		l, _ := newTestLimiter(20, false)
		fail(t, l, "a", 19)
		ok, done := l.Attempt("a")
		done(Succeeded)
		fail(t, l, "a", 1)
		if !ok || allowed(l, "a") {
			t.Fatal("want 19 failures + a success + 1 failure to make the address wait")
		}
	})
	t.Run("keys are independent", func(t *testing.T) {
		l, _ := newTestLimiter(20, true)
		fail(t, l, "a", 20)
		if !allowed(l, "b") {
			t.Fatal("b refused for a's failures")
		}
	})
	t.Run("an idle key is forgotten after an hour", func(t *testing.T) {
		l, clock := newTestLimiter(20, true)
		fail(t, l, "k", 20)
		clock.advance(time.Hour + time.Second)
		fail(t, l, "k", 1)
		if !allowed(l, "k") {
			t.Fatal("refused: the old failures still count")
		}
	})
	t.Run("the threshold is set per limiter", func(t *testing.T) {
		l, clock := newTestLimiter(5, true)
		fail(t, l, "k", 4)
		if !allowed(l, "k") {
			t.Fatal("refused after 4 failures")
		}
		fail(t, l, "k", 1)
		if allowed(l, "k") {
			t.Fatal("allowed right after 5 failures")
		}
		clock.advance(time.Second)
		if !allowed(l, "k") {
			t.Fatal("refused 1 s after the 5th failure")
		}
	})
	t.Run("concurrent attempts cannot pass the threshold", func(t *testing.T) {
		// Each attempt counts before its slow password hash runs, so attempts
		// that start together cannot all pass the check.
		l, _ := newTestLimiter(20, true)
		const attempts = 30
		var started, finished sync.WaitGroup
		started.Add(attempts)
		finished.Add(attempts)
		var mu sync.Mutex
		allowedCount := 0
		for range attempts {
			go func() {
				defer finished.Done()
				ok, done := l.Attempt("k")
				started.Done()
				started.Wait() // the hash stub: runs once every attempt has started
				if ok {
					mu.Lock()
					allowedCount++
					mu.Unlock()
				}
				done(Failed)
			}()
		}
		finished.Wait()
		if allowedCount != 20 {
			t.Fatalf("%d attempts allowed, want exactly 20", allowedCount)
		}
	})
	t.Run("a full table drops a forgotten key, then the oldest key below the threshold", func(t *testing.T) {
		l, clock := newTestLimiter(20, true)
		l.maxKeys = 3
		fail(t, l, "forgotten", 1)
		clock.advance(time.Hour)
		fail(t, l, "older", 1)
		clock.advance(time.Microsecond)
		fail(t, l, "newer", 1)
		clock.advance(time.Microsecond)
		fail(t, l, "x", 1) // the table is full: drops "forgotten"
		fail(t, l, "y", 1) // drops "older"
		_, forgotten := l.keys["forgotten"]
		_, older := l.keys["older"]
		_, newer := l.keys["newer"]
		if forgotten || older || !newer || len(l.keys) != 3 {
			t.Fatalf("keys %v; want forgotten and older dropped, newer kept", l.keys)
		}
	})
	t.Run("a full table never drops a key that reached the threshold", func(t *testing.T) {
		// Invented emails cost nothing; they must not reset a real account's
		// count, even after its current wait is over (10.1 §8.1 item 2).
		l, clock := newTestLimiter(20, true)
		l.refusesWhenFull = true
		l.maxKeys = 10
		fail(t, l, "attacked", 20)
		clock.advance(time.Second) // its wait is over; it is still kept
		for i := range 10 * l.maxKeys {
			fail(t, l, fmt.Sprint(i), 1)
		}
		if k, ok := l.keys["attacked"]; !ok || k.failures != 20 || len(l.keys) > l.maxKeys {
			t.Fatalf("%d keys, attacked %+v kept %t; want at most %d and attacked kept with 20 failures", len(l.keys), k, ok, l.maxKeys)
		}
		fail(t, l, "attacked", 1) // the 21st waits 2 s, as without the flood
		clock.advance(2*time.Second - time.Millisecond)
		if allowed(l, "attacked") {
			t.Fatal("allowed before the 2 s wait of the 21st failure")
		}
		clock.advance(time.Millisecond)
		if !allowed(l, "attacked") {
			t.Fatal("refused after the 2 s wait of the 21st failure")
		}
	})
	t.Run("a full table of penalised keys refuses a new key", func(t *testing.T) {
		l, clock := newTestLimiter(20, true)
		l.refusesWhenFull = true
		l.maxKeys = 10
		for i := range l.maxKeys {
			fail(t, l, fmt.Sprint(i), 20)
		}
		clock.advance(time.Second) // every wait is over, but the keys stay
		if allowed(l, "new") {
			t.Fatal("a new key was allowed with no room and nothing evictable")
		}
		clock.advance(time.Hour)
		if !allowed(l, "new") {
			t.Fatal("a new key refused although every key is forgotten")
		}
	})
	t.Run("a full table of penalised keys evicts the oldest when its limiter allows it", func(t *testing.T) {
		// 10.1 §8.1 item 2: every table but the two guarding a password.
		l, clock := newTestLimiter(20, false)
		l.maxKeys = 10
		for i := range l.maxKeys {
			fail(t, l, fmt.Sprint(i), 20)
			clock.advance(time.Minute)
		}
		if !allowed(l, "new") {
			t.Fatal("a new key was refused by an evicting table")
		}
		fail(t, l, "new", 1)
		if _, kept := l.keys["0"]; kept || len(l.keys) != l.maxKeys {
			t.Fatalf("keys %v; want key 0 (oldest last failure) evicted and the table full", l.keys)
		}
		for i := 1; i < l.maxKeys; i++ {
			if k := l.keys[fmt.Sprint(i)]; k.failures != 20 {
				t.Fatalf("penalised key %d: %+v, want kept with 20 failures", i, k)
			}
		}
	})
	t.Run("a full table makes room for 1% of its keys in one scan, in the order of 10.1 §8.1", func(t *testing.T) {
		// One scan per new key would cost O(n) under the mutex for every
		// request of a flood; one scan per 1% of the table amortises it.
		l, clock := newTestLimiter(20, false)
		l.maxKeys = 1000
		for i := range 5 { // the oldest keys, but penalised: kept while others can go
			fail(t, l, fmt.Sprintf("p%d", i), 20)
			clock.advance(time.Millisecond)
		}
		for i := range l.maxKeys - 5 {
			fail(t, l, fmt.Sprint(i), 1)
			clock.advance(time.Millisecond)
		}
		fail(t, l, "new0", 1)
		if l.scans != 1 || len(l.keys) != l.maxKeys-9 {
			t.Fatalf("%d scans, %d keys; want 1 scan that left room for 10 (%d keys)", l.scans, len(l.keys), l.maxKeys-9)
		}
		for i := range 10 {
			if _, kept := l.keys[fmt.Sprint(i)]; kept {
				t.Fatalf("key %d, among the 10 oldest below the threshold, was kept", i)
			}
		}
		for i := range 5 {
			if _, kept := l.keys[fmt.Sprintf("p%d", i)]; !kept {
				t.Fatalf("penalised key p%d was evicted before keys below the threshold", i)
			}
		}
		for i := 1; i < 10; i++ {
			fail(t, l, fmt.Sprintf("new%d", i), 1)
		}
		if l.scans != 1 {
			t.Fatalf("%d scans for the 9 new keys that fit, want none more", l.scans)
		}
		fail(t, l, "new10", 1)
		if l.scans != 2 {
			t.Fatalf("%d scans, want a second one once the room is used", l.scans)
		}
	})
	t.Run("a full table of penalised keys refuses new keys without a scan until one can be forgotten", func(t *testing.T) {
		l, clock := newTestLimiter(20, true)
		l.refusesWhenFull = true
		l.maxKeys = 10
		start := clock.now
		for i := range l.maxKeys {
			fail(t, l, fmt.Sprint(i), 20)
			clock.advance(time.Minute)
		}
		if allowed(l, "new") {
			t.Fatal("a new key was allowed with no room")
		}
		if want := start.Add(forgetAfter); !l.fullUntil.Equal(want) {
			t.Fatalf("fullUntil %v, want %v: the first key's last failure + 1 h", l.fullUntil, want)
		}
		// Until then the table is not scanned: an idle key planted in it
		// stays, and the new key is still refused.
		l.keys["idle"] = limiterKey{failures: 1, lastFailure: start.Add(-2 * forgetAfter)}
		delete(l.keys, "9")
		if allowed(l, "new") {
			t.Fatal("a new key was allowed before fullUntil")
		}
		if _, kept := l.keys["idle"]; !kept {
			t.Fatal("the table was scanned before fullUntil")
		}
		clock.now = start.Add(forgetAfter)
		if !allowed(l, "new") {
			t.Fatal("a new key refused once a key can be forgotten")
		}
	})
	t.Run("a withdrawn attempt does not move the wait", func(t *testing.T) {
		// NotTried and an address key's Succeeded take back their count;
		// they must take back their time too (10.1 §8.1).
		for _, outcome := range []Outcome{NotTried, Succeeded} {
			l, clock := newTestLimiter(20, false)
			fail(t, l, "k", 20)
			start := clock.now
			clock.advance(time.Second) // the wait after 20 failures is over
			ok, done := l.Attempt("k")
			if !ok {
				t.Fatal("refused after the wait")
			}
			done(outcome)
			if k := l.keys["k"]; k.failures != 20 || !k.lastFailure.Equal(start) {
				t.Fatalf("outcome %d: key %+v, want 20 failures, last at %v", outcome, k, start)
			}
			if !allowed(l, "k") {
				t.Fatalf("outcome %d: the withdrawn attempt restarted the wait", outcome)
			}
			clock.now = start.Add(forgetAfter)
			if _, ok := l.live("k", clock.now); ok {
				t.Fatalf("outcome %d: the withdrawn attempt delayed forgetting", outcome)
			}
		}
	})
	t.Run("each limiter has the table policy of 10.1 §8.1", func(t *testing.T) {
		// Item 2, last column: only the tables guarding a password refuse
		// new keys when full of penalised keys. Reauthenticate (refuse) and
		// StartRecovery (evict) come with PR B.
		s := New(nil, Config{Now: time.Now, Password: cheapParams})
		for _, c := range []struct {
			name   string
			l      *Limiter
			refuse bool
		}{
			{"SignIn email", s.signIn.account, true},
			{"SignIn address", s.signIn.address, false},
			{"token use address", s.tokenUse.address, false},
		} {
			if c.l.refusesWhenFull != c.refuse {
				t.Errorf("%s: refusesWhenFull %t, want %t", c.name, c.l.refusesWhenFull, c.refuse)
			}
		}
		if s.tokenUse.account != nil {
			t.Error("token use has an account table; a token names no account yet")
		}
	})
	t.Run("email and address keys live in separate tables", func(t *testing.T) {
		clock := &testClock{now: time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)}
		r := newRateLimit(20, 20, true, clock.Now)
		for i := range 50 { // 25 per address: the last 5 of each are refused
			ok, done := r.attempt(fmt.Sprintf("u%d@example.test", i), fmt.Sprintf("192.0.2.%d:443", i%2))
			if ok {
				done(Failed)
			}
		}
		if r.account == r.address || len(r.account.keys) != 40 || len(r.address.keys) != 2 {
			t.Fatalf("account keys %d, address keys %d; want 40 (refused emails not kept) and 2 in separate tables",
				len(r.account.keys), len(r.address.keys))
		}
	})
	t.Run("an attempt refused by the address key is not counted for the email", func(t *testing.T) {
		clock := &testClock{now: time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)}
		r := newRateLimit(20, 20, true, clock.Now)
		for i := range 20 {
			ok, done := r.attempt(fmt.Sprintf("u%d@example.test", i), "192.0.2.1:443")
			if !ok {
				t.Fatal("refused before the threshold")
			}
			done(Failed)
		}
		if ok, _ := r.attempt("victim@example.test", "192.0.2.1:443"); ok {
			t.Fatal("the address is past its threshold but was allowed")
		}
		if _, counted := r.account.keys["victim@example.test"]; counted {
			t.Fatal("the refused attempt was counted for the email")
		}
	})
}

func TestAddressKey(t *testing.T) {
	t.Run("IPv6 addresses in one /64 share a key", func(t *testing.T) {
		a, b := addressKey("[2001:db8:1:2::1]:443"), addressKey("[2001:db8:1:2:ffff::9]:5000")
		if a != b || a == addressKey("[2001:db8:1:3::1]:443") {
			t.Fatalf("keys %q, %q; want one key per /64", a, b)
		}
	})
	t.Run("IPv4 addresses do not share a key", func(t *testing.T) {
		if addressKey("192.0.2.1:443") == addressKey("192.0.2.2:443") {
			t.Fatal("two IPv4 addresses share a key")
		}
		if addressKey("192.0.2.1:443") != addressKey("192.0.2.1:5000") {
			t.Fatal("the port is part of the key")
		}
	})
	t.Run("an IPv4-mapped IPv6 address is keyed as IPv4", func(t *testing.T) {
		if a, b := addressKey("[::ffff:192.0.2.1]:443"), addressKey("192.0.2.1:443"); a != b {
			t.Fatalf("keys %q and %q, want equal", a, b)
		}
	})
}

// BenchmarkLimiterFullTable is the cost of one new key in a full table of
// maxLimiterKeys keys during a flood, scans included (10.1 §8.1 item 2).
func BenchmarkLimiterFullTable(b *testing.B) {
	l, clock := newTestLimiter(20, false)
	for i := range maxLimiterKeys {
		l.keys[fmt.Sprint(i)] = limiterKey{failures: 1, lastFailure: clock.now.Add(time.Duration(i))}
	}
	clock.advance(time.Second)
	b.ResetTimer()
	for i := range b.N {
		ok, done := l.Attempt(fmt.Sprintf("new%d", i))
		if !ok {
			b.Fatal("a new key was refused")
		}
		done(Failed)
	}
	b.ReportMetric(float64(l.scans)/float64(b.N), "scans/op")
}
