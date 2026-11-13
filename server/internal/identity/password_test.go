package identity

import (
	"bytes"
	"context"
	"errors"
	"slices"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

// cheapParams keep the tests fast; only the first subtest hashes with the
// production parameters.
var cheapParams = PasswordParams{Time: 1, MemoryKiB: 64, Parallelism: 1, SaltBytes: 16, KeyBytes: 32}

func TestPassword(t *testing.T) {
	ctx := t.Context()

	t.Run("SEC-11 the default parameters are Argon2id t=3 p=4 m=64 MiB with a 128-bit salt and a 256-bit tag", func(t *testing.T) {
		want := PasswordParams{Time: 3, MemoryKiB: 65536, Parallelism: 4, SaltBytes: 16, KeyBytes: 32}
		if DefaultPasswordParams != want {
			t.Fatalf("DefaultPasswordParams = %+v, want %+v (09.4 §2)", DefaultPasswordParams, want)
		}
		h := newHasher(DefaultPasswordParams)
		stored, err := h.hash(ctx, "correct horse battery staple")
		if err != nil {
			t.Fatal(err)
		}
		if len(stored.salt) != 16 || len(stored.hash) != 32 || stored.params != want {
			t.Fatalf("salt %d bytes, hash %d bytes, params %+v; want 16, 32, %+v", len(stored.salt), len(stored.hash), stored.params, want)
		}
		for password, want := range map[string]bool{"correct horse battery staple": true, "wrong": false} {
			if ok, err := h.verify(ctx, stored, password); err != nil || ok != want {
				t.Errorf("verify(%q) = %t, %v; want %t", password, ok, err, want)
			}
		}
	})
	t.Run("a hash verifies with its own stored parameters", func(t *testing.T) {
		stored, err := newHasher(cheapParams).hash(ctx, "a long passphrase")
		if err != nil {
			t.Fatal(err)
		}
		other := cheapParams
		other.Time = 2
		if ok, err := newHasher(other).verify(ctx, stored, "a long passphrase"); err != nil || !ok {
			t.Fatalf("verify with other default parameters = %t, %v; want true", ok, err)
		}
	})
	t.Run("two hashes of one password differ", func(t *testing.T) {
		h := newHasher(cheapParams)
		first, err1 := h.hash(ctx, "a long passphrase")
		second, err2 := h.hash(ctx, "a long passphrase")
		if err := errors.Join(err1, err2); err != nil {
			t.Fatal(err)
		}
		if bytes.Equal(first.salt, second.salt) || bytes.Equal(first.hash, second.hash) {
			t.Fatal("two hashes share a salt or a hash")
		}
	})
	t.Run("at most two hashes run at once", func(t *testing.T) {
		h := newHasher(cheapParams)
		h.slots <- struct{}{}
		h.slots <- struct{}{}
		short, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
		defer cancel()
		if _, err := h.hash(short, "a long passphrase"); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("third hash: err = %v, want context.DeadlineExceeded", err)
		}
		if got := h.count.Load(); got != 0 {
			t.Fatalf("%d hashes ran, want 0", got)
		}
	})
	t.Run("SEC-11 a password on the common list is refused whatever its case", func(t *testing.T) {
		first := commonPasswords()[0]
		for _, password := range []string{first, strings.ToUpper(first)} {
			if err := checkPasswordNotCommon(password); !errors.Is(err, ErrPasswordTooCommon) {
				t.Errorf("checkPasswordNotCommon(%q) = %v, want ErrPasswordTooCommon", password, err)
			}
		}
		if err := checkPasswordNotCommon("correct horse battery staple 7"); err != nil {
			t.Errorf("an uncommon password: %v", err)
		}
	})
	t.Run("the common list holds only lower-case entries of at least 12 characters, sorted and unique", func(t *testing.T) {
		list := commonPasswords()
		if len(list) < 1000 {
			t.Fatalf("%d entries, want the embedded list", len(list))
		}
		for _, entry := range list {
			if utf8.RuneCountInString(entry) < 12 || strings.ToLower(entry) != entry {
				t.Fatalf("entry %q is shorter than 12 characters or not lower case", entry)
			}
		}
		// checkPasswordNotCommon searches it by binary search.
		if !slices.IsSorted(list) || len(slices.Compact(slices.Clone(list))) != len(list) {
			t.Fatal("the list is not sorted and unique")
		}
	})
}
