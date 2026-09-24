package corpus

import (
	"errors"
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
	// Open narrows to records that have not ended. Terminal names the status
	// that counts as ended, passed in rather than read here, so matching stays
	// a pure function of the record and the filter.
	Open     bool
	Terminal string
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

// isRecordPath reports whether a path is one this tool writes.
//
// An ALLOWLIST, and that is the whole point. It was a denylist: every .md under
// .luma/ counted as a record unless it sat under one of four named exceptions.
// That is only correct if this tool owns the directory, and it does not —
// .luma/ is shared with the other luma tools, so every directory nobody had
// thought of yet was claimed by default.
//
// What that cost, all of it found in real repositories (BACK-0108): another
// tool's ideas and plans read as this tool's corpus; the files among them that
// were never records reported as corrupt ones on every command; and
// records/violations/, written by a procedure rather than by this binary,
// silently in the corpus here. Each fix extended the denylist by one directory,
// which meant predicting the next tool's layout — a race the denylist loses by
// construction.
//
// The accepted set is exactly what PathFor can produce, plus PROJECT.md, so
// what the tool reads and what it writes cannot drift apart. A caller has
// already checked the .md suffix.
func isRecordPath(rel string) bool {
	if rel == ProjectFile {
		return true
	}

	// A free-standing decision — the one unit that lives outside a work item.
	// Directly inside the directory, never below it: an archived decision is
	// spent, and listing it beside the live ones is how a superseded position
	// gets read as current.
	if dir, file := path.Split(rel); path.Clean(dir) == RecordsDecisionsDir {
		return file != ""
	}

	rest, ok := strings.CutPrefix(rel, WorkItemsDir+"/")
	if !ok {
		return false
	}
	parts := strings.Split(rest, "/")
	switch len(parts) {
	case 2:
		// The work item itself. Its journal sits beside index.md and is a
		// ledger rather than a record.
		return parts[1] == "index.md"
	case 3:
		// A child, under the directory its unit occupies. This is what keeps
		// evidence/ out without naming it: a work item may keep a transcript
		// or an export beside its records, and those are attachments.
		return isChildDir(parts[1])
	}
	return false
}

func matches(i Item, f Filter) bool {
	if f.Unit != "" && i.Type() != f.Unit {
		return false
	}
	// A caller may name a work item by its directory — WORK-00001-payments-v2 —
	// or by the slug half alone, or by its key. All three reach the same work,
	// and requiring the long form would make the key a tax rather than a handle.
	if f.WorkItem != "" && !matchesWorkItem(i.WorkItem, f.WorkItem) {
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
	if f.Open && f.Terminal != "" {
		// A record with no status reads as the first rung, which is never the
		// terminal one — so absence is open, and only an explicit terminal
		// value is closed.
		if s, _ := i.Record.Get("workflow_status"); s == f.Terminal {
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
// ErrAmbiguous reports a reference that names several records rather than none.
//
// It is a different failure from not finding one, and the difference decides
// what a caller should do: a missing record invites creating it, while an
// ambiguous reference means the record exists and the invocation has to be
// narrowed. Collapsing them is how a well-formed query produces a duplicate.
var ErrAmbiguous = errors.New("ambiguous reference")

// ErrNoMatch reports a reference that names no record. It carries the
// reference so an adapter can name it without parsing the message back out.
type ErrNoMatch struct{ Ref string }

func (e *ErrNoMatch) Error() string { return fmt.Sprintf("nothing matches %q", e.Ref) }

// ErrAmbiguousRef is a reference that names several records, with the
// candidates. It unwraps to ErrAmbiguous so callers matching the sentinel keep
// working.
type ErrAmbiguousRef struct {
	Ref   string
	Paths []string
}

func (e *ErrAmbiguousRef) Error() string {
	return fmt.Sprintf("%s: %q could be any of:\n  %s",
		ErrAmbiguous, e.Ref, strings.Join(e.Paths, "\n  "))
}

func (e *ErrAmbiguousRef) Unwrap() error { return ErrAmbiguous }

func Resolve(b *root.Backlog, ref string) (Item, error) {
	// Skips are not surfaced here yet. Resolve answers "which record did you
	// mean", and a broken record reports as not found — misleading, but a
	// smaller wrong than this slice takes on. Tracked on
	// work-items/report-what-a-listing-skipped.
	items, _, err := List(b, Filter{})
	if err != nil {
		return Item{}, err
	}

	// A key is matched by its parsed value — prefix and number — never as a
	// string, so `work-2`, `WORK---0002` and `WORK 2` all find `WORK-0002`.
	// Somebody typing a handle from memory should not have to hold the shift
	// key, count the padding, or land the dash to be understood. The joined
	// form normalizes its key half the same way, so `work-2-lint-the-corpus`
	// resolves as readily as it is written.
	normalizedName := NormalizeName(ref)

	var exact, former, prefix []Item
	for _, it := range items {
		switch {
		case it.Name() != "" && it.Name() == normalizedName:
			exact = append(exact, it)
		case it.Key() != "" && SameKey(it.Key(), ref):
			exact = append(exact, it)
		// A key the record used to answer to. Its own tier, below exact and
		// above everything else, so a live key always wins: during a
		// migration a key can briefly be one record's current key and
		// another's former one, and the record that holds it now is the
		// answer. Allocation never reissues a former key to a different
		// record, so outside that window the tier has one member at most.
		case it.HeldFormerKey(ref):
			former = append(former, it)
		case it.Path == ref || it.Slug() == ref:
			exact = append(exact, it)
		// The slug half alone, for a work item whose directory leads with a
		// key. It is what people typed before the key existed and what they
		// will keep typing, and the long form is not always to hand.
		case SlugOf(it.Slug()) == ref:
			exact = append(exact, it)
		case strings.HasPrefix(it.Slug(), ref) || strings.HasPrefix(SlugOf(it.Slug()), ref):
			prefix = append(prefix, it)
		}
	}
	candidates := exact
	if len(candidates) == 0 {
		candidates = former
	}
	if len(candidates) == 0 {
		candidates = scoped(items, ref)
	}
	if len(candidates) == 0 {
		candidates = prefix
	}

	switch len(candidates) {
	case 0:
		return Item{}, &ErrNoMatch{Ref: ref}
	case 1:
		return candidates[0], nil
	default:
		var paths []string
		for _, c := range candidates {
			paths = append(paths, c.Path)
		}
		return Item{}, &ErrAmbiguousRef{Ref: ref, Paths: paths}
	}
}

// scoped resolves a reference written the way a listing prints one ---
// WORK-0031/tasks/add-the-queue --- where the leading segment names the work
// item and the rest is the path beneath it. A person reading a path out of
// output should be able to type it back in, and before this they could not:
// the full on-disk path worked and nothing shorter did.
//
// The leading segment resolves by the same rules as any work item reference,
// so a key, a name, or the slug half all work. The remainder is matched
// against the path with and without the extension, since neither form is
// wrong to type.
func scoped(items []Item, ref string) []Item {
	head, tail, ok := strings.Cut(ref, "/")
	if !ok || head == "" || tail == "" {
		return nil
	}

	var dir string
	normalizedName := NormalizeName(head)
	for _, it := range items {
		if it.Type() != WorkItem {
			continue
		}
		if (it.Key() != "" && SameKey(it.Key(), head)) || it.Name() == normalizedName ||
			it.Slug() == head || SlugOf(it.Slug()) == head {
			dir = path.Dir(it.Path)
			break
		}
	}
	if dir == "" {
		return nil
	}

	want := path.Join(dir, tail)
	var found []Item
	for _, it := range items {
		if it.Path == want || it.Path == want+".md" {
			found = append(found, it)
		}
	}
	return found
}
