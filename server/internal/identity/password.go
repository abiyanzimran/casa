package identity

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	_ "embed" // the common password list
	"errors"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	"golang.org/x/crypto/argon2"
)

// PasswordParams are the Argon2id parameters stored with each hash, so they
// can rise later without breaking stored hashes (SEC-11; 09.4 §2).
type PasswordParams struct {
	Time, MemoryKiB     uint32
	Parallelism         uint8
	SaltBytes, KeyBytes uint32
}

// DefaultPasswordParams is the second recommended option of RFC 9106 §4:
// t = 3, p = 4, m = 64 MiB, a 128-bit salt and a 256-bit tag (SEC-11;
// 09.4 §2).
var DefaultPasswordParams = PasswordParams{Time: 3, MemoryKiB: 64 * 1024, Parallelism: 4, SaltBytes: 16, KeyBytes: 32}

// maxConcurrentHashes bounds the peak memory of hashing to 2 × 64 MiB
// (09.4 §2).
const maxConcurrentHashes = 2

type passwordHash struct {
	params     PasswordParams
	salt, hash []byte
}

// hasher runs Argon2id at most maxConcurrentHashes at a time and counts its
// runs, so a test can prove that a failed sign-in costs the same whatever
// the reason (SEC-12).
type hasher struct {
	params PasswordParams
	slots  chan struct{}
	count  atomic.Int64
}

func newHasher(params PasswordParams) *hasher {
	return &hasher{params: params, slots: make(chan struct{}, maxConcurrentHashes)}
}

func (h *hasher) hash(ctx context.Context, password string) (passwordHash, error) {
	salt := make([]byte, h.params.SaltBytes)
	_, _ = rand.Read(salt) // never returns an error (crypto/rand, Go 1.24 and later)
	key, err := h.derive(ctx, h.params, salt, password)
	if err != nil {
		return passwordHash{}, err
	}
	return passwordHash{params: h.params, salt: salt, hash: key}, nil
}

// verify compares in constant time, with the parameters stored with the hash.
func (h *hasher) verify(ctx context.Context, stored passwordHash, password string) (bool, error) {
	key, err := h.derive(ctx, stored.params, stored.salt, password)
	if err != nil {
		return false, err
	}
	return subtle.ConstantTimeCompare(key, stored.hash) == 1, nil
}

func (h *hasher) derive(ctx context.Context, params PasswordParams, salt []byte, password string) ([]byte, error) {
	select {
	case h.slots <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	defer func() { <-h.slots }()
	h.count.Add(1)
	return argon2.IDKey([]byte(password), salt, params.Time, params.MemoryKiB, params.Parallelism, params.KeyBytes), nil
}

// ErrPasswordTooCommon rejects a password on the common password list kept
// in the binary (SEC-11; 09.4 §2).
var ErrPasswordTooCommon = errors.New("password is on the common password list")

// commonPasswordFile is SecLists' list, filtered to entries of at least 12
// characters, lower-cased, sorted and unique; its header names the source,
// commit and licence (lead ruling Q10 on the 1.4 plan).
//
//go:embed common-passwords.txt
var commonPasswordFile string

var commonPasswords = sync.OnceValue(func() []string {
	var list []string
	for line := range strings.Lines(commonPasswordFile) {
		if line = strings.TrimSuffix(line, "\n"); !strings.HasPrefix(line, "#") {
			list = append(list, line)
		}
	}
	return list
})

// checkPasswordNotCommon compares in lower case, without Unicode
// normalization (lead ruling Q10 on the 1.4 plan).
func checkPasswordNotCommon(password string) error {
	if _, found := slices.BinarySearch(commonPasswords(), strings.ToLower(password)); found {
		return ErrPasswordTooCommon
	}
	return nil
}
