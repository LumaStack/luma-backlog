package backlog

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/lumastack/luma-backlog/internal/record"
	"github.com/lumastack/luma-backlog/internal/root"
)

// Item is a record plus where it lives.
type Item struct {
	Path     string // relative to the backlog, slash-separated
	Record   *record.Record
	WorkItem string // the work item it sits inside, if any

	// Raw is the file exactly as read. Kept so a caller can be handed a hash
	// of what it actually saw, which is what makes a later write safe.
	Raw []byte
}

// Hash identifies the content this item was read from.
func (i Item) Hash() string { return record.Hash(i.Raw) }

// Slug is the filename without its extension — for a work item, the
// directory name, since its record is index.md.
func (i Item) Slug() string {
	if i.Type() == WorkItem {
		return path.Base(path.Dir(i.Path))
	}
	return strings.TrimSuffix(path.Base(i.Path), ".md")
}

// Type is the record's unit, short form.
func (i Item) Type() string {
	t, _ := i.Record.Get("type")
	return t
}

// Title falls back to the slug, as the format allows.
func (i Item) Title() string {
	if t, ok := i.Record.Get("title"); ok && t != "" {
		return t
	}
	return i.Slug()
}

// Status is what to show for a record's state.
//
// For a worked unit it is the workflow position, or the configured default
// when absent — absence is meaningful rather than missing (docs/spec.md §4.2).
//
// For an outcome it is DERIVED from evidence, never read from a field. There
// is nothing to store that the verification record does not already say, and a
// stored copy could disagree with it.
func (i Item) Status(defaultStatus string) string {
	if i.Type() == Outcome {
		if i.Record.Has("verified") {
			return "passing"
		}
		return "unverified"
	}
	if s, ok := i.Record.Get("workflow_status"); ok && s != "" {
		return s
	}
	if !IsWorked(i.Type()) {
		if s, ok := i.Record.Get("stage"); ok && s != "" {
			return s
		}
		// A type that declares no workflow_status has no lifecycle for a
		// default to fill in. Falling through here stamped "todo" onto a
		// luma/project — a fact nobody wrote, and one a reader cannot tell
		// apart from a status somebody set deliberately.
		return ""
	}
	return defaultStatus
}

// Skip is a file the listing could not use, and the reason it could not.
//
// Skipping is deliberate: one malformed record must not make the whole backlog
// unlistable (docs/spec.md §4.1). Skipping in SILENCE is the part that is
// wrong. A record that vanishes without comment is an invisible absence, which
// is worse than a wrong answer because nothing prompts anyone to look for it.
//
// Permissive means the tool keeps working, not that it says nothing.
type Skip struct {
	Path string
	Err  error
}

// Filter narrows a listing. A zero Filter matches everything.
type Filter struct {
	Unit     string
	WorkItem string
	Status   string
	Kind     string
}

// List reads every record in the backlog, filtered.
//
// Unreadable files are skipped rather than fatal: one malformed record must
// not make the whole backlog unlistable, which is the permissive posture the
// format requires (docs/spec.md §4.1).
func List(b *root.Backlog, f Filter) ([]Item, []Skip, error) {
	var items []Item
	var skipped []Skip

	err := b.Walk(func(rel string) error {
		if !strings.HasSuffix(rel, ".md") || !isRecordPath(rel) {
			return nil
		}
		data, err := b.ReadFile(rel)
		if err != nil {
			skipped = append(skipped, Skip{Path: rel, Err: err})
			return nil
		}
		r, err := record.Parse(data)
		if err != nil {
			skipped = append(skipped, Skip{Path: rel, Err: err})
			return nil
		}
		it := Item{Path: rel, Record: r, WorkItem: WorkItemFromPath(rel), Raw: data}
		if matches(it, f) {
			items = append(items, it)
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	// A skip is reported whatever the filter, because a record that cannot be
	// read cannot be filtered: narrowing the listing must not narrow the
	// warning and hide the very record somebody is looking for.
	sort.Slice(skipped, func(a, b int) bool { return skipped[a].Path < skipped[b].Path })

	// Sorted by path: stable across runs and across machines, which is what
	// makes the output safe to pin in a golden file.
	sort.Slice(items, func(a, b int) bool { return items[a].Path < items[b].Path })
	return items, skipped, nil
}

// isRecordPath excludes files that are not units — the bundle root, journals,
// and Type Definitions.
func isRecordPath(rel string) bool {
	switch {
	case path.Base(rel) == "journal.md":
		return false
	case strings.HasPrefix(rel, "bundles/"):
		// Bundles hold contracts and policy — what is in force, not what is
		// intended — so nothing under them is a backlog record.
		return false
	case strings.HasPrefix(rel, "_types/"):
		// .luma/_types/ holds contracts for documents outside any bundle,
		// which are not this tool's records either.
		return false
	}
	return true
}

func matches(i Item, f Filter) bool {
	if f.Unit != "" && i.Type() != f.Unit {
		return false
	}
	if f.WorkItem != "" && i.WorkItem != f.WorkItem {
		return false
	}
	if f.Kind != "" {
		k, _ := i.Record.Get("kind")
		if CanonicalKind(k) != CanonicalKind(f.Kind) {
			return false
		}
	}
	if f.Status != "" {
		s, _ := i.Record.Get("workflow_status")
		if s != f.Status {
			return false
		}
	}
	return true
}

// Resolve finds a single record from a reference: a path, a slug, or an
// unambiguous prefix of one.
//
// An ambiguous reference is an error listing the candidates, never a guess.
// Picking one quietly is how the wrong record gets edited and nobody finds out
// until later.
func Resolve(b *root.Backlog, ref string) (Item, error) {
	// Skips are not surfaced here yet. Resolve answers "which record did you
	// mean", and a broken record reports as not found — misleading, but a
	// smaller wrong than this slice takes on. Tracked on
	// work-items/report-what-a-listing-skipped.
	items, _, err := List(b, Filter{})
	if err != nil {
		return Item{}, err
	}

	// A key is matched case-insensitively, so `work-00002` finds `WORK-00002`.
	// Somebody typing a handle from memory should not have to hold the shift
	// key to be understood.
	normalized := NormalizeKey(ref)

	var exact, prefix []Item
	for _, it := range items {
		switch {
		case it.Key() != "" && it.Key() == normalized:
			exact = append(exact, it)
		case it.Path == ref || it.Slug() == ref:
			exact = append(exact, it)
		case strings.HasPrefix(it.Slug(), ref):
			prefix = append(prefix, it)
		}
	}
	candidates := exact
	if len(candidates) == 0 {
		candidates = prefix
	}

	switch len(candidates) {
	case 0:
		return Item{}, fmt.Errorf("nothing matches %q", ref)
	case 1:
		return candidates[0], nil
	default:
		var paths []string
		for _, c := range candidates {
			paths = append(paths, c.Path)
		}
		return Item{}, fmt.Errorf("%q matches more than one record:\n  %s",
			ref, strings.Join(paths, "\n  "))
	}
}
