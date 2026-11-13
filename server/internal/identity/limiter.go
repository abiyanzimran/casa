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
	// maxLimiterKeys bounds the memory of the key table.
	maxLimiterKeys = 100_000
)

// Limiter counts consecutive failures per key: after threshold failures, the
// next attempt waits until the last failure + min(2^(n−threshold) s, 15 min),
// n being the number of failures (10.1 §8.1). A caller keeps separate keys
// per account (or normalized email) and per network address (SEC-12).
//
// ponytail: process memory of the one binary (09.4 §2); moves to PostgreSQL
// when the public stage runs several processes.
type Limiter struct {
	threshold int
	now       func() time.Time
	mu        sync.Mutex
	keys      map[string]limiterKey
}

type limiterKey struct {
	failures    int
	lastFailure time.Time
}

// NewLimiter returns an empty Limiter with its failure threshold (10.1 §8.1:
// one limiter per RPC), reading the time from now.
func NewLimiter(threshold int, now func() time.Time) *Limiter {
	return &Limiter{threshold: threshold, now: now, keys: map[string]limiterKey{}}
}

// Allow reports whether key may try now.
func (l *Limiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	k, ok := l.live(key)
	return !ok || k.failures < l.threshold || !l.now().Before(k.lastFailure.Add(wait(k.failures-l.threshold)))
}

// Fail records a failed attempt of key.
func (l *Limiter) Fail(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	k, ok := l.live(key)
	if !ok && len(l.keys) >= maxLimiterKeys {
		l.evict()
	}
	l.keys[key] = limiterKey{failures: k.failures + 1, lastFailure: l.now()}
}

// Succeed clears key.
func (l *Limiter) Succeed(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.keys, key)
}

// live returns key's entry unless it is forgotten.
func (l *Limiter) live(key string) (limiterKey, bool) {
	k, ok := l.keys[key]
	if ok && !l.now().Before(k.lastFailure.Add(forgetAfter)) {
		delete(l.keys, key)
		return limiterKey{}, false
	}
	return k, ok
}

// evict drops forgotten keys, or else the key whose last failure is oldest,
// so a key under attack, which keeps failing, is the last to go.
//
// ponytail: an O(n) scan of at most maxLimiterKeys, only when the table is
// full; a heap by last failure if a full table becomes common.
func (l *Limiter) evict() {
	var oldest string
	for key, k := range l.keys {
		if !l.now().Before(k.lastFailure.Add(forgetAfter)) {
			delete(l.keys, key)
		} else if oldest == "" || k.lastFailure.Before(l.keys[oldest].lastFailure) {
			oldest = key
		}
	}
	if len(l.keys) >= maxLimiterKeys {
		delete(l.keys, oldest)
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
