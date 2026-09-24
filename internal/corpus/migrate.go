package corpus

import (
	"path"
	"strings"

	"github.com/lumastack/luma-backlog/internal/root"
)

// KeyRename is one record's key and directory name, before and after.
type KeyRename struct {
	OldKey  string
	NewKey  string
	OldName string
	NewName string
	// Dir is the work item's directory name before the rename --- the same as
	// OldName, kept separate because a caller repointing links reads the name
	// and a caller moving files reads the directory, and conflating them is
	// how one of the two gets missed.
	Dir string
}

// KeyCollision is a record that could not take its target key because another
// record answers to it.
type KeyCollision struct {
	// Key is the key that could not be taken.
	Key string
	// Record is the work item that wanted it.
	Record string
	// HeldBy is the work item that already answers to it.
	HeldBy string
	// Former is true where the holder answers to it as a key it used to have
	// rather than as its current one. Both block, and saying which is the
	// difference between "rename that one first" and "that key is retired".
	Former bool
}

// KeyPlan is what a migration would do, computed before anything is written.
type KeyPlan struct {
	// Renames are the records that will move, in corpus order.
	Renames []KeyRename
	// Collisions are the records that will not, each naming what blocked it.
	Collisions []KeyCollision
	// AlreadyCorrect counts the records whose key is at the target prefix
	// already. They are not touched, and saying how many were skipped for the
	// good reason keeps that distinct from the ones that failed.
	AlreadyCorrect int
}

// PlanKeyMigration works out which records move to the target prefix and which
// cannot, without writing anything.
//
// **Separated from applying it because a plan is worth seeing first.** A
// corpus-wide rewrite is the least recoverable thing this tool does, and the
// same reasoning that gave `rank repair` a dry run applies here with 111
// records rather than one.
//
// **Only the prefix moves.** WORK-0123 becomes BACK-0123 and keeps its number,
// so a diff stays reviewable and an external reference stays recognizable even
// before anybody updates it. A number is only ever changed by --renumber, which
// operates on what this reports as a collision.
//
// **A record already at the target prefix is left alone**, not rewritten to
// itself. A corpus can hold a mixture --- this one holds two prefixes --- and a
// migration that assumed a single source would touch records that need nothing
// doing to them.
//
// **A collision is against any OTHER record, current key or former.** A record
// reclaiming a key from its own former_keys is not a collision: it is the same
// record, so the promise that a key resolves to exactly one record still holds,
// and without the exception migrating back would be refused for everything that
// ever moved.
func PlanKeyMigration(b *root.Backlog, target string) (KeyPlan, error) {
	items, _, err := List(b, Filter{Unit: WorkItem})
	if err != nil {
		return KeyPlan{}, err
	}

	var plan KeyPlan
	for _, it := range items {
		key := it.Key()
		prefix, number, ok := ParseKey(key)
		if !ok {
			// A record written before keys existed. Nothing to migrate and
			// not an error --- allocation will give it one if anybody asks.
			continue
		}
		if prefix == target {
			plan.AlreadyCorrect++
			continue
		}

		want := FormatKeyAs(target, number)
		if holder, former, taken := heldByAnother(items, want, it); taken {
			plan.Collisions = append(plan.Collisions, KeyCollision{
				Key:    want,
				Record: it.Name(),
				HeldBy: holder,
				Former: former,
			})
			continue
		}

		plan.Renames = append(plan.Renames, KeyRename{
			OldKey:  NormalizeKey(key),
			NewKey:  want,
			OldName: it.Name(),
			NewName: want + "-" + SlugOf(it.Slug()),
			Dir:     it.WorkItem,
		})
	}
	return plan, nil
}

// heldByAnother reports whether a key is answered to by a record other than
// the one asking, and whether that record answers to it as a former key.
//
// The exclusion is by path rather than by key, because the record asking is
// mid-rename by definition and its key is the thing being changed.
func heldByAnother(items []Item, key string, asking Item) (holder string, former, taken bool) {
	for _, it := range items {
		if it.Path == asking.Path {
			continue
		}
		if it.Key() != "" && SameKey(it.Key(), key) {
			return it.Name(), false, true
		}
		if it.HeldFormerKey(key) {
			return it.Name(), true, true
		}
	}
	return "", false, false
}

// RewriteNamesIn returns text with every old work item name replaced by its
// new one.
//
// **Names move and keys do not, which is the whole rule.** A full name ---
// BACK-0031-reshape-the-command-surface --- does not survive a migration: it is
// a name rather than a key, resolution falls back to string equality, and a
// slug defeats that. A bare key does survive, because former_keys answers for
// it, and rewriting one can be wrong: it may name a work item in a different
// project using the same prefix. So one must move and the other must not.
//
// **The boundary is required rather than convenient.** A replacement anchored
// on a name alone would corrupt a longer name that starts with it. No name in
// this corpus is a prefix of another today, and that is not something to
// depend on.
func RewriteNamesIn(text string, renames []KeyRename) (string, int) {
	changed := 0
	for _, r := range renames {
		// The progress guarantee, and the only one this loop has: `from`
		// strictly increases each turn because `end > i >= from`, which holds
		// only while OldName is non-empty. An empty name makes strings.Index
		// return 0 forever. Stated rather than left looking like a nil check,
		// since it is what makes the loop terminate.
		if r.OldName == "" {
			continue
		}
		// Scanned with an explicit offset. Restarting the search from the top
		// after a skipped match is an infinite loop, because the match that
		// was skipped is still the first one found.
		var out strings.Builder
		from := 0
		for {
			i := strings.Index(text[from:], r.OldName)
			if i < 0 {
				break
			}
			i += from
			end := i + len(r.OldName)
			if !endsName(text, end) {
				// A longer name starting with this one. Carry it through
				// unchanged and keep looking past it.
				out.WriteString(text[from:end])
				from = end
				continue
			}
			out.WriteString(text[from:i])
			out.WriteString(r.NewName)
			from = end
			changed++
		}
		out.WriteString(text[from:])
		text = out.String()
	}
	return text, changed
}

// endsName reports whether a name ends at this offset rather than running on
// into a longer one. A slug character continuing means this was a prefix of
// something else.
func endsName(text string, at int) bool {
	if at >= len(text) {
		return true
	}
	c := text[at]
	return !(c == '-' || c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'))
}

// NameOfKey renders the directory name a key and slug make together.
func NameOfKey(key, slug string) string { return key + "-" + SlugOf(slug) }

// WorkItemPath is where a work item's own record lives, given its directory.
func WorkItemPath(dir string) string {
	return path.Join(WorkItemsDir, dir, "index.md")
}
