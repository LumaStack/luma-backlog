package backlog

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lumastack/luma-backlog/internal/root"
)

// KeyPrefix is the only prefix supported today.
//
// It is written INTO the record rather than derived from configuration, so a
// repository that later chooses its own prefix changes what gets written and
// not what already exists. A derived key would silently rename every record in
// the corpus the moment the setting changed.
const KeyPrefix = "WORK"

// keyPattern matches a work item key: WORK-0002.
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
var keyPattern = regexp.MustCompile(`^([A-Z]+)-(\d+)$`)

// FormatKey renders a key from its number.
func FormatKey(number int) string {
	return fmt.Sprintf("%s-%04d", KeyPrefix, number)
}

// IsKey reports whether a reference looks like a key rather than a slug.
func IsKey(ref string) bool {
	return keyPattern.MatchString(strings.ToUpper(ref))
}

// NormalizeKey upper-cases a key so `work-00002` finds `WORK-00002`. Anything
// that is not a key is returned unchanged, since it is somebody's slug.
func NormalizeKey(ref string) string {
	if IsKey(ref) {
		return strings.ToUpper(ref)
	}
	return ref
}

// highestKey reports the largest key number in use across the project.
//
// One sequence for the whole backlog. The number is what somebody says out loud
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
		m := keyPattern.FindStringSubmatch(strings.ToUpper(k))
		if m == nil {
			continue
		}
		if n, convErr := strconv.Atoi(m[2]); convErr == nil && n > highest {
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

// namePattern matches the joined form: WORK-0002-lint-the-corpus.
var namePattern = regexp.MustCompile(`^([A-Za-z]+-\d+)(-.*)$`)

// NormalizeName upper-cases the key half of a name and leaves the slug alone,
// since a slug is lower-case by construction and upper-casing it would stop it
// matching.
func NormalizeName(ref string) string {
	m := namePattern.FindStringSubmatch(ref)
	if m == nil {
		return ref
	}
	return strings.ToUpper(m[1]) + m[2]
}
