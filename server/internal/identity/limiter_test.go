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
		return NewLimiter(20, clock.Now), clock
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
	t.Run("the threshold is set per limiter", func(t *testing.T) {
		clock := &testClock{now: time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)}
		l := NewLimiter(5, clock.Now)
		failN(l, "k", 4)
		if !l.Allow("k") {
			t.Fatal("blocked after 4 failures")
		}
		l.Fail("k")
		if l.Allow("k") {
			t.Fatal("allowed right after 5 failures")
		}
		clock.advance(time.Second)
		if !l.Allow("k") {
			t.Fatal("blocked 1 s after the 5th failure")
		}
	})
	t.Run("a full table drops the key whose last failure is oldest", func(t *testing.T) {
		l, clock := newLimiter()
		l.Fail("oldest")
		clock.advance(time.Microsecond)
		failN(l, "attacked", 20)
		for i := range maxLimiterKeys - 1 { // one more than fits
			clock.advance(time.Microsecond)
			l.Fail(fmt.Sprint(i))
		}
		if got := len(l.keys); got != maxLimiterKeys {
			t.Fatalf("%d keys, want %d", got, maxLimiterKeys)
		}
		if _, ok := l.keys["oldest"]; ok {
			t.Fatal("the key with the oldest failure is still there")
		}
		if l.Allow("attacked") {
			t.Fatal("a newer key that is waiting was dropped")
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
