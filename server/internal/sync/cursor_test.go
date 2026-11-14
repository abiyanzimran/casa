package sync_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/read"
	"github.com/abiyanzimran/casa/server/internal/sync"
)

var testKeyBytes = bytes.Repeat([]byte{7}, 32)

func testKey(t *testing.T) sync.CursorKey {
	t.Helper()
	k, err := sync.NewCursorKey(testKeyBytes)
	if err != nil {
		t.Fatal(err)
	}
	return k
}

// sealed is a cursor of the layout of plan 1.7 decision 9 with a valid MAC
// over any body, so a test can reach the checks behind the MAC.
func sealed(body []byte) string {
	mac := hmac.New(sha256.New, testKeyBytes)
	mac.Write([]byte("casa.sync.cursor.v1"))
	mac.Write(body)
	return base64.RawURLEncoding.EncodeToString(append(body, mac.Sum(nil)[:16]...))
}

// body is the layout of plan 1.7 decision 9; snapshot appends kind and after.
func body(version, mode byte, household, member pgtype.UUID, pos uint64, snapshot ...byte) []byte {
	b := append([]byte{version, mode}, household.Bytes[:]...)
	b = append(b, member.Bytes[:]...)
	b = binary.BigEndian.AppendUint64(b, pos)
	return append(b, snapshot...)
}

func rejected(t *testing.T, name string, k sync.CursorKey, s string, household, member pgtype.UUID) {
	t.Helper()
	if c, err := k.Decode(s, household, member); !errors.Is(err, sync.ErrCursorRejected) {
		t.Errorf("%s: Decode = %+v, %v; want ErrCursorRejected", name, c, err)
	}
}

// TestCursor proves the cursor of ADR-017 items 3 and 8 and plan 1.7
// decision 9: a server MAC binds the position to the household and the
// member, so a cursor of anyone else, or changed in any way, is refused with
// the one error the API answers like a household that does not exist
// (SEC-05).
func TestCursor(t *testing.T) {
	h, m := newID(), newID()
	feed := sync.Cursor{Household: h, Member: m, Pos: 42}
	after := newID()
	snapshot := sync.Cursor{Household: h, Member: m, Pos: 7, Snapshot: &read.SnapshotMark{Kind: read.KindMember, After: after}}

	t.Run("a feed and a snapshot cursor round-trip", func(t *testing.T) {
		k := testKey(t)
		for _, c := range []sync.Cursor{feed, snapshot,
			{Household: h, Member: m, Pos: 0, Snapshot: &read.SnapshotMark{Kind: read.KindHousehold, After: h}},
			{Household: h, Member: m, Pos: 1<<63 - 1, Snapshot: &read.SnapshotMark{Kind: read.KindInvitation, After: after}},
		} {
			got, err := k.Decode(k.Encode(c), h, m)
			if err != nil {
				t.Fatalf("Decode(Encode(%+v)): %v", c, err)
			}
			if got.Household != c.Household || got.Member != c.Member || got.Pos != c.Pos ||
				(got.Snapshot == nil) != (c.Snapshot == nil) || c.Snapshot != nil && *got.Snapshot != *c.Snapshot {
				t.Errorf("round trip of %+v gave %+v", c, got)
			}
		}
	})

	t.Run("a cursor of another member or another household is rejected", func(t *testing.T) {
		k := testKey(t)
		for _, c := range []sync.Cursor{feed, snapshot} {
			s := k.Encode(c)
			rejected(t, "another member", k, s, h, newID())
			rejected(t, "another household", k, s, newID(), m)
			rejected(t, "household and member swapped", k, s, m, h)
		}
	})

	t.Run("a cursor with any one bit flipped is rejected", func(t *testing.T) {
		k := testKey(t)
		for _, c := range []sync.Cursor{feed, snapshot} {
			raw, err := base64.RawURLEncoding.DecodeString(k.Encode(c))
			if err != nil {
				t.Fatal(err)
			}
			for i := range raw {
				for bit := range 8 {
					changed := bytes.Clone(raw)
					changed[i] ^= 1 << bit
					rejected(t, "flipped", k, base64.RawURLEncoding.EncodeToString(changed), h, m)
				}
			}
		}
	})

	t.Run("an empty body, a truncated cursor, bad base64 and an unknown version, mode or kind are rejected", func(t *testing.T) {
		k := testKey(t)
		s := k.Encode(snapshot)
		raw, _ := base64.RawURLEncoding.DecodeString(s)
		// A feed cursor of 58 bytes ends in 4 unused bits; ruling P9's strict
		// decoding refuses them set, so no second string decodes to the
		// same bytes.
		fs := k.Encode(feed)
		const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
		trailing := fs[:len(fs)-1] + string(alphabet[strings.IndexByte(alphabet, fs[len(fs)-1])|1])
		for name, bad := range map[string]string{
			"empty":                           "",
			"empty body with a MAC":           sealed(nil),
			"truncated by one character":      s[:len(s)-1],
			"truncated by one byte":           base64.RawURLEncoding.EncodeToString(raw[:len(raw)-1]),
			"the MAC only":                    base64.RawURLEncoding.EncodeToString(raw[len(raw)-16:]),
			"not base64":                      "!" + s[1:],
			"padded":                          base64.URLEncoding.EncodeToString(raw) + "=",
			"a + of the standard alphabet":    s[:5] + "+" + s[6:],
			"a / of the standard alphabet":    s[:5] + "/" + s[6:],
			"non-zero trailing bits":          trailing,
			"version 0":                       sealed(body(0, 0, h, m, 1)),
			"version 2":                       sealed(body(2, 0, h, m, 1)),
			"mode 2":                          sealed(body(1, 2, h, m, 1)),
			"feed mode at snapshot length":    sealed(body(1, 0, h, m, 1, 2, 1)),
			"snapshot mode at feed length":    sealed(body(1, 1, h, m, 1)),
			"kind 0":                          sealed(body(1, 1, h, m, 1, append([]byte{0}, after.Bytes[:]...)...)),
			"kind 4":                          sealed(body(1, 1, h, m, 1, append([]byte{4}, after.Bytes[:]...)...)),
			"a position beyond int64":         sealed(body(1, 0, h, m, 1<<63)),
			"a valid body with one more byte": sealed(append(body(1, 0, h, m, 1), 0)),
		} {
			rejected(t, name, k, bad, h, m)
		}
		// The sealed helper writes what Encode writes, so the cases above
		// fail on the field they name, not on the MAC.
		if got, err := k.Decode(sealed(body(1, 1, h, m, 7, append([]byte{2}, after.Bytes[:]...)...)), h, m); err != nil || got.Pos != 7 {
			t.Fatalf("a sealed valid body: %+v, %v", got, err)
		}
	})

	t.Run("a cursor made with another key is rejected", func(t *testing.T) {
		other, err := sync.NewCursorKey(bytes.Repeat([]byte{8}, 32))
		if err != nil {
			t.Fatal(err)
		}
		rejected(t, "another key", testKey(t), other.Encode(feed), h, m)
	})

	t.Run("a key that is not 32 bytes is refused", func(t *testing.T) {
		for _, n := range []int{0, 16, 31, 33, 64} {
			if _, err := sync.NewCursorKey(make([]byte, n)); err == nil {
				t.Errorf("a key of %d bytes was accepted", n)
			}
		}
	})
}
