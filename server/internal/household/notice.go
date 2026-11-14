package household

import (
	"bytes"
	"context"
	"crypto/sha256"
	"embed"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/abiyanzimran/casa/server/internal/household/householddb"
)

//go:embed notice/*.json
var publishedNotices embed.FS

// PublishedNotices are the published versions of the operator notice, one
// file <variant>-<version>.json each, embedded so a version ships with its
// build (SEC-35; 09.3 §1 "Mekanisme"; plan 1.6 decision 6). A published file
// never changes: TestPublishedNoticesAreFrozen pins every one, so new text is
// a new number.
var PublishedNotices fs.FS = mustSub(publishedNotices, "notice")

func mustSub(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err) // unreachable: dir is a constant valid path
	}
	return sub
}

// pilotNoticeVariant is the variant every account reads in 1.6: the pilot
// couple's (10.1 §11.3).
// ponytail: one variant for all; FR-FAM-019/AC-9 (Tier 3) chooses it per account.
const pilotNoticeVariant = "couple"

// operatorPlaceholder is filled with the server's operator name when a
// version is shown (09.3 §1 "Mekanisme").
const operatorPlaceholder = "{pengelola}"

// NoticeParagraph is one paragraph of a version: plain text, and the
// feature it describes, "" for one that always applies (SEC-34 (5)).
type NoticeParagraph struct{ Text, Feature string }

// Notice is a version of the operator notice as an account reads it, with
// {pengelola} filled in. Variant, Version and TextSHA256 are its identity
// for the account (09.3 §1 "Mekanisme").
type Notice struct {
	Variant      string
	Version      int32
	Title        string
	Paragraphs   []NoticeParagraph
	TextSHA256   []byte
	Acknowledged bool
	// EarlierAcknowledged is true when the account acknowledged an issue
	// before this one, A→B→A included (09.3 §1 note of 14 November 2026).
	EarlierAcknowledged bool
}

// noticeHash is SHA-256 over the length-prefixed variant, version (decimal),
// title and each paragraph's text and feature, each field after its length
// as a big-endian uint32, so moving a paragraph boundary or changing a tag
// is a new identity (10.1 §11 note of 14 November 2026 (a)).
func noticeHash(n Notice) []byte {
	fields := []string{n.Variant, strconv.Itoa(int(n.Version)), n.Title}
	for _, p := range n.Paragraphs {
		fields = append(fields, p.Text, p.Feature)
	}
	h := sha256.New()
	for _, field := range fields {
		_ = binary.Write(h, binary.BigEndian, uint32(len(field))) //nolint:gosec // G115: a notice field is far below 4 GiB
		_, _ = h.Write([]byte(field))                             // a hash.Hash never fails a write
	}
	return h.Sum(nil)
}

// latestNotice is the highest version of variant in the notice files, with
// the operator name filled in and its hash. A malformed file fails closed: no
// notice is shown and every held command waits (SEC-34).
// ponytail: reads and parses the files on each call; they are a few KiB.
func (s *Service) latestNotice(variant string) (Notice, error) {
	names, err := fs.Glob(s.notices, variant+"-*.json")
	if err != nil {
		return Notice{}, err
	}
	n := Notice{Variant: variant}
	var latest string
	for _, name := range names {
		number := strings.TrimSuffix(strings.TrimPrefix(name, variant+"-"), ".json")
		version, err := strconv.ParseInt(number, 10, 32)
		if err != nil || version <= 0 || strconv.FormatInt(version, 10) != number {
			return Notice{}, fmt.Errorf("notice file %s: the version is not a number from 1 to 2147483647", name)
		}
		if int32(version) > n.Version {
			n.Version, latest = int32(version), name
		}
	}
	if latest == "" {
		return Notice{}, fmt.Errorf("no notice of variant %s", variant)
	}
	b, err := fs.ReadFile(s.notices, latest)
	if err != nil {
		return Notice{}, err
	}
	var file struct {
		Title      string `json:"title"`
		Paragraphs []struct {
			Text    string `json:"text"`
			Feature string `json:"feature"`
		} `json:"paragraphs"`
	}
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&file); err != nil {
		return Notice{}, fmt.Errorf("notice file %s: %w", latest, err)
	}
	if file.Title == "" || len(file.Paragraphs) == 0 {
		return Notice{}, fmt.Errorf("notice file %s: no title or no paragraph", latest)
	}
	fill := strings.NewReplacer(operatorPlaceholder, s.operatorName).Replace
	n.Title = fill(file.Title)
	for _, p := range file.Paragraphs {
		n.Paragraphs = append(n.Paragraphs, NoticeParagraph{Text: fill(p.Text), Feature: p.Feature})
	}
	n.TextSHA256 = noticeHash(n)
	return n, nil
}

// CurrentNotice is the latest version for account, in tx whose scope is the
// account. The version is issued when its identity differs from the one last
// issued to the account, a rename back A→B→A included: a new row at the last
// issue_seq + 1 (INSERT … ON CONFLICT DO NOTHING). issued is true only when
// this call inserted, so the caller writes the security trail line in the
// same tx, once per issue; a concurrent first sight fails serialization and
// retries (09.3 §1 "Mekanisme"; 10.1 note before §11.3; SEC-34).
// Acknowledged is the last issue's; EarlierAcknowledged covers the issues
// before it.
func (s *Service) CurrentNotice(ctx context.Context, tx pgx.Tx, account pgtype.UUID) (n Notice, issued bool, err error) {
	if n, err = s.latestNotice(pilotNoticeVariant); err != nil {
		return Notice{}, false, err
	}
	q := householddb.New(tx)
	last, err := q.LastNoticeIssue(ctx, account)
	switch {
	case err == nil && last.Variant == n.Variant && last.Version == n.Version && bytes.Equal(last.TextSha256, n.TextSHA256):
		n.Acknowledged, n.EarlierAcknowledged = last.Acknowledged, last.EarlierAcknowledged
		return n, false, nil
	case err != nil && !errors.Is(err, pgx.ErrNoRows):
		return Notice{}, false, err
	}
	inserted, err := q.IssueNotice(ctx, householddb.IssueNoticeParams{
		AccountID: account, IssueSeq: last.IssueSeq + 1, Variant: n.Variant, Version: n.Version, TextSha256: n.TextSHA256,
		IssuedAt: s.now(),
	})
	if err != nil {
		return Notice{}, false, err
	}
	if inserted != 1 {
		// SERIALIZABLE makes this unreachable; never answer OK while the
		// account stays held without an issued notice.
		return Notice{}, false, errors.New("household: operator notice issue was not inserted")
	}
	n.EarlierAcknowledged = last.Acknowledged || last.EarlierAcknowledged // no row yet: both false
	return n, true, nil
}

// AcknowledgeNotice records account's acknowledgement of the version it was
// shown, in tx after CurrentNotice issued the latest one. Only the last
// issue, with all three of variant, version and hash, is accepted, else
// ErrStateChanged (SEC-34; 10.1 §11 note of 14 November 2026 (b)).
// acknowledged is false for a repeat: the first acknowledgement's time stays
// (FR-FAM-019/AC-2) and no second trail line is written.
func (s *Service) AcknowledgeNotice(ctx context.Context, tx pgx.Tx, account pgtype.UUID, variant string, version int32, textSHA256 []byte) (acknowledged bool, err error) {
	latest, err := s.latestNotice(pilotNoticeVariant)
	if err != nil {
		return false, err
	}
	if variant != latest.Variant || version != latest.Version || !bytes.Equal(textSHA256, latest.TextSHA256) {
		return false, ErrStateChanged
	}
	changed, err := householddb.New(tx).AcknowledgeNotice(ctx, householddb.AcknowledgeNoticeParams{
		AcknowledgedAt: s.now(), AccountID: account, Variant: variant, Version: version, TextSha256: textSHA256,
	})
	return changed == 1, err
}
