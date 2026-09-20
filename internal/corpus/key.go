package corpus

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lumastack/luma-backlog/internal/config"
	"github.com/lumastack/luma-backlog/internal/root"
)

// keyPattern matches anything written as a key: WORK-0002, work-2,
// WORK---2, "WORK  2".
//
// The prefix rule is config.KeyPrefixRule — Jira Cloud project-key rules
// (WORK-0082): an uppercase letter first, then uppercase letters or digits,
// two to ten characters. That admits R2D2-7 and turns away a one-letter
// prefix — which is deliberate, so a stray `x-1` in prose stays a slug rather
// than becoming a key. One fragment shared with the config validator, so what
// a repository may configure and what this file can read cannot drift apart.
//
// The separator tolerates runs of dashes and spaces because people type keys
// from memory and quote them out of prose. What a sloppy spelling resolves to
// is decided by ParseKey; nothing here changes what gets written to disk.
var keyPattern = regexp.MustCompile(`^(` + config.KeyPrefixRule + `)[ -]+(\d+)$`)

// ParseKey reads a reference as a key, however it was spelled. This is the
// only reader — every comparison goes through it, so two spellings of one key
// cannot disagree anywhere (WORK-0082: keys are compared as parsed values,
// never as strings).
func ParseKey(ref string) (prefix string, number int, ok bool) {
	m := keyPattern.FindStringSubmatch(strings.ToUpper(ref))
	if m == nil {
		return "", 0, false
	}
	n, err := strconv.Atoi(m[2])
	if err != nil {
		// More digits than an int holds is not a key anybody was given.
		return "", 0, false
	}
	return m[1], n, true
}

// FormatKeyAs renders a key under a prefix. It is what NormalizeKey and
// allocation both funnel into, so there is exactly one form a key is ever
// printed in.
//
// Four digits, matching the ADR numbers, and the same accepted cost: two
// branches can both claim the next one and somebody repairs it on merge
// (records/decisions/ADR-0003).
//
// Padding exists so a lexical sort matches a numeric one, which is what `ls`,
// git and an editor give you. It stops working past 9999 — and by then nobody
// is reading a directory of ten thousand work items by eye, so the property
// fails exactly where it had stopped being worth anything. If that order ever
// matters at that scale, the fix is sorting numerically in the tool, which
// touches no record and does not disturb a key that is meant never to change.
// The width is a minimum, so a five-digit key renders as itself rather than
// being squeezed back to four.
func FormatKeyAs(prefix string, number int) string {
	return fmt.Sprintf("%s-%04d", prefix, number)
}

// IsKey reports whether a reference looks like a key rather than a slug.
func IsKey(ref string) bool {
	_, _, ok := ParseKey(ref)
	return ok
}

// NormalizeKey renders any spelling of a key in its canonical form, so
// `work---2` becomes `WORK-0002` — already correct to print, which is why the
// canonical form is the stored one rather than a lowercase comparison form:
// one representation means no bug about which one is in hand. Anything that is
// not a key is returned unchanged, since it is somebody's slug.
func NormalizeKey(ref string) string {
	if prefix, number, ok := ParseKey(ref); ok {
		return FormatKeyAs(prefix, number)
	}
	return ref
}

// SameKey reports whether two references name the same key, whatever their
// spelling. Where neither side parses as a key, plain equality answers —
// two slugs are the same by being the same string.
func SameKey(a, b string) bool {
	ap, an, aok := ParseKey(a)
	bp, bn, bok := ParseKey(b)
	if aok != bok {
		return false
	}
	if aok {
		return ap == bp && an == bn
	}
	return a == b
}

// highestKey reports the largest key number in use across the project.
//
// One sequence for the whole corpus. The number is what somebody says out loud
// or writes in a commit, so it has to mean one record — which is the same
// reason decision numbers are allocated project-wide rather than per directory.
func highestKey(b *root.Backlog) (int, error) {
	highest := 0
	items, _, err := List(b, Filter{Unit: WorkItem})
	if err != nil {
		return 0, err
	}
	for _, it := range items {
		k, ok := it.Record.Get("key")
		if !ok {
			continue
		}
		// Parsed, not pattern-matched: a key stored in an unusual spelling
		// must still count, or the next allocation reuses its number
		// (WORK-0082 — the same failure WORK-0040 hit from a different cause).
		if _, n, isKey := ParseKey(k); isKey && n > highest {
			highest = n
		}
	}
	return highest, nil
}

// Key returns the work item's key, or empty when it has none. A record written
// before keys existed has none, and that is not an error.
func (i Item) Key() string {
	k, _ := i.Record.Get("key")
	return k
}

// Name is what a record is called: WORK-0002-lint-the-corpus.
//
// Key and slug joined, the way a decision's filename joins its number and slug,
// and literally the name of the directory — so the model and the filesystem
// agree rather than each holding their own idea of what this record is called.
//
// A TITLE is prose for a person: "Lint the corpus". A NAME is what the thing is
// called, unique like the key and legible like the slug. Nothing else in the
// format uses the word, which is why it was free to take.
//
// Not "ref": in a tool that lives inside git, a ref is a branch or a tag, and
// borrowing a word the surrounding system has already claimed is the mistake
// that set aside `change` for a kind and `committed` for a rung.
func (i Item) Name() string {
	// The key leads the directory name, so a work item's slug already carries
	// it and joining again would double it. Kept as its own method because it
	// is the published name for "the identifier you write", and because a
	// record created before keys were in paths still refs as its slug.
	return i.Slug()
}

// namePattern matches the joined form: WORK-0002-lint-the-corpus. The key
// half takes only a single dash — a name comes off a directory listing, not
// out of prose, so the sloppy separators ParseKey tolerates have no business
// here, and a run of dashes inside a slug must stay a slug.
var namePattern = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9]{1,9})-(\d+)(-.*)$`)

// NormalizeName renders the key half of a name canonically and leaves the
// slug alone — a slug is lower-case by construction and upper-casing it would
// stop it matching. `work-2-lint-the-corpus` finds `WORK-0002-lint-the-corpus`.
func NormalizeName(ref string) string {
	m := namePattern.FindStringSubmatch(ref)
	if m == nil {
		return ref
	}
	n, err := strconv.Atoi(m[2])
	if err != nil {
		return ref
	}
	return FormatKeyAs(strings.ToUpper(m[1]), n) + m[3]
}
