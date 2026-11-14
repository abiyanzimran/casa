package sync

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"math"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/read"
)

// The cursor of ADR-017 items 3 and 8 and 09.3 §3, in the layout of plan 1.7
// decision 9:
//
//	body   = version (1, = 1) | mode (1: 0 feed, 1 snapshot)
//	         | household (16) | member (16) | pos (8, uint64 big-endian)
//	         [ snapshot only: kind (1: 1 household, 2 member, 3 invitation) | after (16) ]
//	mac    = HMAC-SHA256(key, "casa.sync.cursor.v1" || body)[:16]
//	cursor = base64url without padding (body || mac)
const (
	cursorVersion   = 1
	modeFeed        = 0
	modeSnapshot    = 1
	macDomain       = "casa.sync.cursor.v1"
	macLen          = 16
	feedBodyLen     = 1 + 1 + 16 + 16 + 8
	snapshotBodyLen = feedBodyLen + 1 + 16
)

// kindBytes are the snapshot kinds of the layout; byKindByte inverts it.
var (
	kindBytes  = map[read.ObjectKind]byte{read.KindHousehold: 1, read.KindMember: 2, read.KindInvitation: 3}
	byKindByte = map[byte]read.ObjectKind{1: read.KindHousehold, 2: read.KindMember, 3: read.KindInvitation}
)

// cursorEncoding decodes strictly (ruling P9): unused trailing bits must be
// zero, so no second string decodes to the same bytes.
var cursorEncoding = base64.RawURLEncoding.Strict()

// ErrCursorRejected is the one answer to a cursor this server did not make
// for this household and member, whatever is wrong with it (plan 1.7
// decision 9). The API answers it like a household that does not exist
// (SEC-05), so it says nothing more.
var ErrCursorRejected = errors.New("cursor rejected")

// CursorKey is the server's 32-byte cursor MAC key (plan 1.7 decision 4).
type CursorKey struct{ key [32]byte }

// NewCursorKey makes a cursor key of exactly 32 bytes.
func NewCursorKey(key []byte) (CursorKey, error) {
	var k CursorKey
	if len(key) != len(k.key) {
		return CursorKey{}, errors.New("sync cursor key must be 32 bytes")
	}
	copy(k.key[:], key)
	return k, nil
}

// Cursor is a reader's place in Sync: Pos is the feed position, or P0 of a
// snapshot in progress (plan 1.7 decision 10).
type Cursor struct {
	Household, Member pgtype.UUID
	Pos               int64
	Snapshot          *read.SnapshotMark // nil in feed mode
}

func (k CursorKey) mac(body []byte) []byte {
	m := hmac.New(sha256.New, k.key[:])
	m.Write([]byte(macDomain))
	m.Write(body)
	return m.Sum(nil)[:macLen]
}

// Encode seals c. Pull makes a snapshot cursor only after a page with an
// object, so its mark always names a kind of the layout.
func (k CursorKey) Encode(c Cursor) string {
	mode := byte(modeFeed)
	if c.Snapshot != nil {
		mode = modeSnapshot
	}
	b := append([]byte{cursorVersion, mode}, c.Household.Bytes[:]...)
	b = append(b, c.Member.Bytes[:]...)
	b = binary.BigEndian.AppendUint64(b, uint64(c.Pos)) //nolint:gosec // Pos is a feed position, never negative
	if c.Snapshot != nil {
		b = append(b, kindBytes[c.Snapshot.Kind])
		b = append(b, c.Snapshot.After.Bytes[:]...)
	}
	return cursorEncoding.EncodeToString(append(b, k.mac(b)...))
}

// Decode refuses with ErrCursorRejected anything but a cursor k made for
// this household and member (plan 1.7 decision 9). The MAC is checked first,
// in constant time, so nothing of an unsealed body is read.
func (k CursorKey) Decode(s string, household, member pgtype.UUID) (Cursor, error) {
	raw, err := cursorEncoding.DecodeString(s)
	if err != nil || len(raw) != feedBodyLen+macLen && len(raw) != snapshotBodyLen+macLen {
		return Cursor{}, ErrCursorRejected
	}
	b, mac := raw[:len(raw)-macLen], raw[len(raw)-macLen:]
	if !hmac.Equal(mac, k.mac(b)) {
		return Cursor{}, ErrCursorRejected
	}
	wantLen := map[byte]int{modeFeed: feedBodyLen, modeSnapshot: snapshotBodyLen}
	if n, ok := wantLen[b[1]]; b[0] != cursorVersion || !ok || n != len(b) {
		return Cursor{}, ErrCursorRejected
	}
	c := Cursor{Household: uuidAt(b[2:]), Member: uuidAt(b[18:])}
	if c.Household != household || c.Member != member {
		return Cursor{}, ErrCursorRejected
	}
	pos := binary.BigEndian.Uint64(b[34:])
	if pos > math.MaxInt64 {
		return Cursor{}, ErrCursorRejected
	}
	c.Pos = int64(pos)
	if b[1] == modeSnapshot {
		kind, ok := byKindByte[b[feedBodyLen]]
		if !ok {
			return Cursor{}, ErrCursorRejected
		}
		c.Snapshot = &read.SnapshotMark{Kind: kind, After: uuidAt(b[feedBodyLen+1:])}
	}
	return c, nil
}

func uuidAt(b []byte) pgtype.UUID {
	id := pgtype.UUID{Valid: true}
	copy(id.Bytes[:], b)
	return id
}
