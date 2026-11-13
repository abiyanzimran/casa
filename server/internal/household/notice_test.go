package household

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"maps"
	"slices"
	"testing"
	"testing/fstest"
)

// publishedNoticeHashes pins the SHA-256 of every published notice file
// (09.3 §1 "Mekanisme"; SEC-34, SEC-35): a published version never changes,
// so new text is always a new file with a new number, and a new file adds
// its line here.
//
// couple-1.json was replaced once before any issue to anyone, for the
// [Klien web] sentence of FR-FAM-019/AC-11 (10.1 note before §11.3: it stays
// v1); from the first deploy on, a change is a new number.
var publishedNoticeHashes = map[string]string{
	"couple-1.json": "02b46c0f6c1295a05c2b71b1fe6850517158ecd2e50cc4e4eb040e09f151b468",
}

func TestPublishedNoticesAreFrozen(t *testing.T) {
	files, err := fs.Glob(PublishedNotices, "*")
	if err != nil {
		t.Fatal(err)
	}
	if want := slices.Sorted(maps.Keys(publishedNoticeHashes)); !slices.Equal(files, want) {
		t.Fatalf("published notice files = %q, want %q: a new version adds its hash here", files, want)
	}
	for _, name := range files {
		b, err := fs.ReadFile(PublishedNotices, name)
		if err != nil {
			t.Fatal(err)
		}
		if got := sha256.Sum256(b); hex.EncodeToString(got[:]) != publishedNoticeHashes[name] {
			t.Errorf("%s changed (SHA-256 %x): a published version is never edited; publish a new number", name, got)
		}
	}
}

func TestNoticeHash(t *testing.T) {
	base := Notice{Variant: "couple", Version: 1, Title: "T", Paragraphs: []NoticeParagraph{{Text: "ab"}, {Text: "c"}}}
	// 10.1 §11 note of 14 November 2026 (a): every field is length-prefixed,
	// so no two different notices share an encoding.
	t.Run("moving a paragraph boundary changes the hash", func(t *testing.T) {
		moved := base
		moved.Paragraphs = []NoticeParagraph{{Text: "a"}, {Text: "bc"}}
		if bytes.Equal(noticeHash(base), noticeHash(moved)) {
			t.Fatal("same hash after moving a paragraph boundary")
		}
	})
	t.Run("changing a feature tag changes the hash", func(t *testing.T) {
		tagged := base
		tagged.Paragraphs = []NoticeParagraph{{Text: "ab", Feature: "care"}, {Text: "c"}}
		if bytes.Equal(noticeHash(base), noticeHash(tagged)) {
			t.Fatal("same hash after tagging a paragraph")
		}
	})
	t.Run("the operator name is part of the hash", func(t *testing.T) {
		notices := fstest.MapFS{"couple-1.json": {Data: []byte(`{"title": "T", "paragraphs": [{"text": "oleh {pengelola}.", "feature": ""}]}`)}}
		a, err := New(nil, nil, "Pengelola Uji", notices).latestNotice(pilotNoticeVariant)
		if err != nil {
			t.Fatal(err)
		}
		b, err := New(nil, nil, "Nama Lain", notices).latestNotice(pilotNoticeVariant)
		if err != nil {
			t.Fatal(err)
		}
		if a.Paragraphs[0].Text != "oleh Pengelola Uji." || bytes.Equal(a.TextSHA256, b.TextSHA256) {
			t.Fatalf("text %q, hashes %x and %x; want the name filled in and two hashes", a.Paragraphs[0].Text, a.TextSHA256, b.TextSHA256)
		}
	})
	t.Run("the hash is of the length-prefixed fields", func(t *testing.T) {
		// variant, version in decimal, title, then each paragraph's text and
		// feature, each after its length as a big-endian uint32.
		encoded := "\x00\x00\x00\x06couple" + "\x00\x00\x00\x011" + "\x00\x00\x00\x01T" +
			"\x00\x00\x00\x02ab" + "\x00\x00\x00\x00" + "\x00\x00\x00\x01c" + "\x00\x00\x00\x00"
		if want := sha256.Sum256([]byte(encoded)); !bytes.Equal(noticeHash(base), want[:]) {
			t.Fatalf("hash %x, want %x", noticeHash(base), want)
		}
	})
}

func TestLatestNotice(t *testing.T) {
	file := func(title string) *fstest.MapFile {
		return &fstest.MapFile{Data: []byte(`{"title": "` + title + `", "paragraphs": [{"text": "p", "feature": ""}]}`)}
	}
	t.Run("the highest number of the variant is the latest", func(t *testing.T) {
		notices := fstest.MapFS{"couple-2.json": file("two"), "couple-10.json": file("ten"), "family-11.json": file("other")}
		n, err := New(nil, nil, "P", notices).latestNotice("couple")
		if err != nil || n.Version != 10 || n.Title != "ten" {
			t.Fatalf("latest = %d %q (%v), want 10 \"ten\"", n.Version, n.Title, err)
		}
	})
	t.Run("a malformed file or name fails closed", func(t *testing.T) {
		for name, notices := range map[string]fstest.MapFS{
			"none":          {},
			"leading zero":  {"couple-01.json": file("x")},
			"zero":          {"couple-0.json": file("x")},
			"beyond int32":  {"couple-2147483648.json": file("x")},
			"unknown field": {"couple-1.json": {Data: []byte(`{"title": "x", "paragraphs": [], "extra": 1}`)}},
			"no paragraph":  {"couple-1.json": {Data: []byte(`{"title": "x", "paragraphs": []}`)}},
			"empty title":   {"couple-1.json": {Data: []byte(`{"title": "", "paragraphs": [{"text": "p", "feature": ""}]}`)}},
		} {
			if _, err := New(nil, nil, "P", notices).latestNotice("couple"); err == nil {
				t.Errorf("%s: no error", name)
			}
		}
	})
}
