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
	Path        string // relative to the backlog, slash-separated
	Record      *record.Record
	Deliverable string // the deliverable it sits inside, if any

	// Raw is the file exactly as read. Kept so a caller can be handed a hash
	// of what it actually saw, which is what makes a later write safe.
	Raw []byte
}

// Hash identifies the content this item was read from.
func (i Item) Hash() string { return record.Hash(i.Raw) }

// Slug is the filename without its extension — for a deliverable, the
// directory name, since its record is index.md.
func (i Item) Slug() string {
	if i.Type() == Deliverable {
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
// when absent — absence is meaningful rather than missing (docs/SPEC.md §4.2).
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
		if s, ok := i.Record.Get("lifecycle_status"); ok && s != "" {
			return s
		}
	}
	return defaultStatus
}

// Filter narrows a listing. A zero Filter matches everything.
type Filter struct {
	Unit        string
	Deliverable string
	Status      string
}

// List reads every record in the backlog, filtered.
//
// Unreadable files are skipped rather than fatal: one malformed record must
// not make the whole backlog unlistable, which is the permissive posture the
// format requires (docs/SPEC.md §4.1).
func List(b *root.Backlog, f Filter) ([]Item, error) {
	var items []Item

	err := b.Walk(func(rel string) error {
		if !strings.HasSuffix(rel, ".md") || !isRecordPath(rel) {
			return nil
		}
		data, err := b.ReadFile(rel)
		if err != nil {
			return nil
		}
		r, err := record.Parse(data)
		if err != nil {
			return nil
		}
		it := Item{Path: rel, Record: r, Deliverable: DeliverableFromPath(rel), Raw: data}
		if matches(it, f) {
			items = append(items, it)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sorted by path: stable across runs and across machines, which is what
	// makes the output safe to pin in a golden file.
	sort.Slice(items, func(a, b int) bool { return items[a].Path < items[b].Path })
	return items, nil
}

// isRecordPath excludes files that are not units — the bundle root, journals,
// and Type Definitions.
func isRecordPath(rel string) bool {
	switch {
	case rel == "index.md":
		return false
	case path.Base(rel) == "journal.md":
		return false
	case strings.HasPrefix(rel, "_types/"):
		return false
	}
	return true
}

func matches(i Item, f Filter) bool {
	if f.Unit != "" && i.Type() != f.Unit {
		return false
	}
	if f.Deliverable != "" && i.Deliverable != f.Deliverable {
		return false
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
	items, err := List(b, Filter{})
	if err != nil {
		return Item{}, err
	}

	var exact, prefix []Item
	for _, it := range items {
		switch {
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
