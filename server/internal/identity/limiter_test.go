package identity

import (
	"fmt"
	"testing"
	"time"
)

type testClock struct{ now time.Time }

func (c *testClock) Now() time.Time          { return c.now }
func (c *testClock) advance(d time.Duration) { c.now = c.now.Add(d) }

func failN(l *Limiter, key string, n int) {
	for range n {
		l.Fail(key)
	}
}

func TestLimiter(t *testing.T) {
	newLimiter := func() (*Limiter, *testClock) {
		clock := &testClock{now: time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)}
		return NewLimiter(clock.Now), clock
	}
	t.Run("SEC-12 after 20 consecutive failures the next attempt waits", func(t *testing.T) {
		l, clock := newLimiter()
		failN(l, "k", 19)
		if !l.Allow("k") {
			t.Fatal("blocked after 19 failures")
		}
		l.Fail("k")
		if l.Allow("k") {
			t.Fatal("allowed right after 20 failures")
		}
		clock.advance(time.Second)
		if !l.Allow("k") {
			t.Fatal("blocked 1 s after the 20th failure")
		}
		l.Fail("k") // the 21st waits 2 s
		clock.advance(1900 * time.Millisecond)
		if l.Allow("k") {
			t.Fatal("allowed 1.9 s after the 21st failure")
		}
		clock.advance(100 * time.Millisecond)
		if !l.Allow("k") {
			t.Fatal("blocked 2 s after the 21st failure")
		}
	})
	t.Run("the wait stops growing at 15 minutes", func(t *testing.T) {
		l, clock := newLimiter()
		failN(l, "k", 100)
		clock.advance(15*time.Minute - time.Nanosecond)
		if l.Allow("k") {
			t.Fatal("allowed before 15 minutes")
		}
		clock.advance(time.Nanosecond)
		if !l.Allow("k") {
			t.Fatal("blocked after 15 minutes")
		}
	})
	t.Run("a success clears the key", func(t *testing.T) {
		l, _ := newLimiter()
		failN(l, "k", 20)
		l.Succeed("k")
		if !l.Allow("k") {
			t.Fatal("blocked after a success")
		}
		l.Fail("k")
		if !l.Allow("k") {
			t.Fatal("one failure after a success blocks")
		}
	})
	t.Run("keys are independent", func(t *testing.T) {
		l, _ := newLimiter()
		failN(l, "a", 20)
		if !l.Allow("b") {
			t.Fatal("b blocked by a's failures")
		}
	})
	t.Run("an idle key is forgotten after an hour", func(t *testing.T) {
		l, clock := newLimiter()
		failN(l, "k", 20)
		clock.advance(time.Hour + time.Second)
		l.Fail("k")
		if !l.Allow("k") {
			t.Fatal("blocked: the old failures still count")
		}
	})
	t.Run("the key table stays bounded", func(t *testing.T) {
		l, clock := newLimiter()
		failN(l, "attacked", 20)
		for i := range maxLimiterKeys + 1 {
			clock.advance(time.Microsecond)
			l.Fail(fmt.Sprint(i))
			if i == maxLimiterKeys/2 {
				failN(l, "attacked", 1) // still under attack: recent
			}
		}
		if got := len(l.keys); got > maxLimiterKeys {
			t.Fatalf("%d keys, want at most %d", got, maxLimiterKeys)
		}
		if l.Allow("attacked") {
			t.Fatal("a key that failed recently was evicted before older ones")
		}
	})
}
