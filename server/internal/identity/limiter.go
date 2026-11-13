package identity

import (
	"net/netip"
	"sync"
	"time"
)

// The rate limit of SEC-12 as 10.1 §8.1 sets it [Usulan] (lead ruling Q11
// on the 1.4 plan; 09.4 §2).
const (
	// maxWait caps the doubling wait.
	maxWait = 15 * time.Minute
	// forgetAfter drops a key this long after its last failure.
	forgetAfter = time.Hour
	// maxLimiterKeys bounds the memory of each limiter's key table.
	maxLimiterKeys = 100_000
)

// Limiter counts consecutive failures per key: after threshold failures, the
// next attempt waits until the last failure + min(2^(n−threshold) s, 15 min),
// n being the number of failures (10.1 §8.1). There is one Limiter per RPC
// and key kind, so account (email) keys and address keys fill separate tables
// (SEC-12).
//
// ponytail: process memory of the one binary (09.4 §2); moves to PostgreSQL
// when the public stage runs several processes.
type Limiter struct {
	threshold int
	// clearsOnSuccess: a success clears an account key, never an address
	// key (10.1 §8.1).
	clearsOnSuccess bool
	maxKeys         int
	now             func() time.Time
	mu              sync.Mutex
	keys            map[string]limiterKey
	// fullUntil: the table is full of keys past their threshold and none
	// can be forgotten before this time, so a new key is refused without
	// another scan (a flood of new keys costs no O(n) scan each).
	fullUntil time.Time
}

type limiterKey struct {
	failures    int
	lastFailure time.Time
}

// Outcome is how an allowed attempt ended.
type Outcome int

const (
	// Failed: the attempt was a failure, as counted when it started.
	Failed Outcome = iota
	// Succeeded: for an account key the key is cleared; for an address key
	// only this attempt's count is taken back.
	Succeeded
	// NotTried: the attempt did not run (for example another limiter
	// refused it), so its count is taken back.
	NotTried
)

// NewLimiter returns an empty Limiter with its failure threshold
// (10.1 §8.1), reading the time from now.
func NewLimiter(threshold int, clearsOnSuccess bool, now func() time.Time) *Limiter {
	return &Limiter{threshold: threshold, clearsOnSuccess: clearsOnSuccess, maxKeys: maxLimiterKeys, now: now, keys: map[string]limiterKey{}}
}

// Attempt reports whether key may try now and, if so, counts the attempt as
// a failure at once, before the caller's slow work (a password hash) runs,
// so attempts that start together cannot all pass. The caller always calls
// done with the outcome; for a refused attempt done does nothing. A new key
// with no room in a full table is refused (fail closed during a flood).
func (l *Limiter) Attempt(key string) (allowed bool, done func(Outcome)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := l.now()
	k, ok := l.live(key, now)
	if ok && l.waiting(k, now) || !ok && len(l.keys) >= l.maxKeys && (now.Before(l.fullUntil) || !l.evict(now)) {
		return false, func(Outcome) {}
	}
	l.keys[key] = limiterKey{failures: k.failures + 1, lastFailure: now}
	return true, func(outcome Outcome) { l.end(key, outcome, k.lastFailure, now) }
}

// end takes back a withdrawn attempt (NotTried, or Succeeded on a key a
// success does not clear): its count, and its time, previous, unless a
// later attempt has set the time since (attemptedAt no longer matches), so a
// withdrawn attempt moves neither the wait nor the forgetting.
func (l *Limiter) end(key string, outcome Outcome, previous, attemptedAt time.Time) {
	if outcome == Failed {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	k, ok := l.keys[key]
	switch {
	case !ok:
	case outcome == Succeeded && l.clearsOnSuccess, k.failures <= 1:
		delete(l.keys, key)
	default:
		k.failures--
		if k.lastFailure.Equal(attemptedAt) && !previous.IsZero() {
			k.lastFailure = previous
		}
		l.keys[key] = k
	}
}

// waiting reports whether k must wait before its next attempt.
func (l *Limiter) waiting(k limiterKey, now time.Time) bool {
	return k.failures >= l.threshold && now.Before(k.lastFailure.Add(wait(k.failures-l.threshold)))
}

// live returns key's entry unless it is forgotten.
func (l *Limiter) live(key string, now time.Time) (limiterKey, bool) {
	k, ok := l.keys[key]
	if ok && !now.Before(k.lastFailure.Add(forgetAfter)) {
		delete(l.keys, key)
		return limiterKey{}, false
	}
	return k, ok
}

// evict makes room for one key: it drops keys idle for forgetAfter, or else
// the key below the threshold whose last failure is oldest. A key that
// reached the threshold stays until it is forgotten, even when its current
// wait is over, so a flood of invented keys cannot reset a real account's
// count (10.1 §8.1). It reports false when nothing could go, and then sets
// fullUntil to the earliest time a key can be forgotten.
//
// ponytail: an O(n) scan of at most maxLimiterKeys, only when the table is
// full; a heap by last failure if a full table becomes common.
func (l *Limiter) evict(now time.Time) bool {
	var oldest string
	var forgettable time.Time
	for key, k := range l.keys {
		switch {
		case !now.Before(k.lastFailure.Add(forgetAfter)):
			delete(l.keys, key)
		case k.failures >= l.threshold:
			if at := k.lastFailure.Add(forgetAfter); forgettable.IsZero() || at.Before(forgettable) {
				forgettable = at
			}
		case oldest == "" || k.lastFailure.Before(l.keys[oldest].lastFailure):
			oldest = key
		}
	}
	if len(l.keys) < l.maxKeys {
		return true
	}
	if oldest == "" {
		l.fullUntil = forgettable
		return false
	}
	delete(l.keys, oldest)
	return true
}

// rateLimit is one RPC's rate limit: an account (or email) key and a
// network-address key, each in its own Limiter so neither kind can crowd the
// other out (SEC-12; 10.1 §8.1). account is nil where a token does not yet
// name an account.
type rateLimit struct {
	account, address *Limiter
}

// newRateLimit returns a rateLimit with the given thresholds; an account
// threshold of 0 means no account key.
func newRateLimit(accountThreshold, addressThreshold int, now func() time.Time) rateLimit {
	r := rateLimit{address: NewLimiter(addressThreshold, false, now)}
	if accountThreshold > 0 {
		r.account = NewLimiter(accountThreshold, true, now)
	}
	return r
}

// attempt counts one attempt for accountKey (ignored without an account
// key) and for the address of peer, or refuses it without counting either.
func (r rateLimit) attempt(accountKey, peer string) (allowed bool, done func(Outcome)) {
	accountDone := func(Outcome) {}
	if r.account != nil {
		if allowed, accountDone = r.account.Attempt(accountKey); !allowed {
			return false, accountDone
		}
	}
	allowed, addressDone := r.address.Attempt(addressKey(peer))
	if !allowed {
		accountDone(NotTried)
		return false, addressDone
	}
	return true, func(outcome Outcome) {
		accountDone(outcome)
		addressDone(outcome)
	}
}

// wait is min(2^exponent s, maxWait).
func wait(exponent int) time.Duration {
	if exponent >= 10 { // 2^10 s > 15 min; also keeps the shift from overflowing
		return maxWait
	}
	return min(time.Second<<exponent, maxWait)
}

// addressKey is the network-address key of a peer address "host:port"
// (10.1 §8.1): an IPv4 address on its own, an IPv6 address by its /64, and an
// IPv4-mapped IPv6 address as IPv4. One home or mobile connection holds a
// whole /64, so a full IPv6 address could change at every attempt.
func addressKey(peer string) string {
	addrPort, err := netip.ParseAddrPort(peer)
	if err != nil {
		return peer // unreachable for a TCP peer; still one key per value
	}
	addr := addrPort.Addr().Unmap().WithZone("")
	if addr.Is4() {
		return addr.String()
	}
	return netip.PrefixFrom(addr, 64).Masked().String()
}
