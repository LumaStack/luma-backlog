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

// keyPattern matches a work item key: WORK-00002.
//
// Five digits rather than four, and the same accepted cost as an ADR number:
// two branches can both claim the next one and somebody renumbers on merge.
var keyPattern = regexp.MustCompile(`^([A-Z]+)-(\d+)$`)

// FormatKey renders a key from its number.
func FormatKey(number int) string {
	return fmt.Sprintf("%s-%05d", KeyPrefix, number)
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

// Ref is how a record is written and said: WORK-00002-lint-the-corpus.
//
// Key and slug joined, the way a decision's filename joins its number and slug.
// One string rather than two columns, because a reader wants one thing to copy
// and a listing with a key column leaves an empty cell on every record that
// carries none.
//
// It is a REFERENCE, not a path. The directory is still the slug alone, so
// nothing about this moves a file or changes what a record is (docs/spec.md
// §7.1).
func (i Item) Ref() string {
	// The key leads the directory name, so a work item's slug already carries
	// it and joining again would double it. Kept as its own method because it
	// is the published name for "the identifier you write", and because a
	// record created before keys were in paths still refs as its slug.
	return i.Slug()
}

// refPattern matches the joined form: WORK-00002-lint-the-corpus.
var refPattern = regexp.MustCompile(`^([A-Za-z]+-\d+)(-.*)$`)

// NormalizeRef upper-cases the key half of a joined reference and leaves the
// slug alone, since a slug is lower-case by construction and upper-casing it
// would stop it matching.
func NormalizeRef(ref string) string {
	m := refPattern.FindStringSubmatch(ref)
	if m == nil {
		return ref
	}
	return strings.ToUpper(m[1]) + m[2]
}
