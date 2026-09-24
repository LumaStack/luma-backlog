package app

import (
	"sort"
	"strings"

	"github.com/lumastack/luma-backlog/internal/corpus"
)

// View is one record as a surface needs to show it.
//
// A view rather than the record itself, so an adapter never holds an engine
// type. That is what lets the containment rule be a build-time fact: a surface
// that cannot name the engine cannot reach past this layer to it
// (.luma/records/decisions/ADR-0004).
type View struct {
	Path string
	Type string
	// Key is the handle somebody quotes — BACK-0002. Empty where there is
	// none: only a work item carries one, and a record written before keys
	// existed has none either.
	Key string
	// Name is the two joined — BACK-0002-lint-the-corpus.
	Name  string
	Slug  string
	Title string
	// Status is empty when the record's type declares no workflow status.
	Status   string
	WorkItem string
	// Rank is work order --- <status ordinal>.<position>, ADR-0005. Empty on a
	// record nobody has placed, which is most of them: ranking one record
	// writes one file, so a status is not seeded just because it was read.
	Rank string
	// Completion is how many of a work item's live outcomes are proven, out of
	// how many there are. Nil on anything that is not a work item — absent
	// means not counted, never counted as zero.
	Completion *corpus.Counts
	// Created and Modified are the record's stamps. Empty where a record has
	// none --- one written before stamps existed, or one nobody has edited.
	Created  Stamp
	Modified Stamp
}

// Stamp is who did something and when.
type Stamp struct {
	By string
	At string
}

// Record is a view with everything a single-record read needs.
type Record struct {
	View
	// Hash identifies the content that was read. Pass it back to a set with
	// IfUnchanged and a write that would clobber somebody else's change is
	// refused rather than applied (docs/spec.md §6.3).
	Hash string
	// Fields is the frontmatter as written, including keys this tool knows
	// nothing about — dropping them would quietly hide another system's state.
	Fields map[string]any
	// Order is the frontmatter's key order, so a rendering can follow the file
	// rather than sorting and telling a reader something the file does not.
	Order []string
	// Raw is each field as the file holds it. Kept beside Fields because the
	// two renderings need different things — a structured consumer wants the
	// decoded value, and a person reading a terminal wants what is actually
	// written on the line.
	Raw  map[string]string
	Body string
}

func (s *Session) view(it corpus.Item) View {
	return View{
		Path:     it.Path,
		Type:     it.Type(),
		Key:      it.Key(),
		Name:     it.Name(),
		Slug:     it.Slug(),
		Title:    it.Title(),
		Status:   it.Status(s.Config.DefaultStatusFor(it.Type())),
		WorkItem: it.WorkItem,
		Rank:     rankOf(it),
		Created:  stampOf(it, "created"),
		Modified: stampOf(it, "modified"),
	}
}

func (s *Session) record(it corpus.Item) (Record, error) {
	fields := map[string]any{}
	order := it.Record.Keys()
	for _, k := range order {
		var v any
		if node := it.Record.Node(k); node != nil {
			if err := node.Decode(&v); err != nil {
				return Record{}, FailureError("decoding %s: %w", k, err)
			}
		}
		fields[k] = v
	}
	raw := map[string]string{}
	for _, k := range order {
		if v, ok := it.Record.Get(k); ok {
			raw[k] = v
			continue
		}
		// A list field, rendered as a person would read it out. Without this
		// every sequence in the frontmatter is invisible to any caller working
		// from Raw — `former_keys` on a migrated record, `advances` on a task
		// — because Get answers for scalars only. Fields still carries the
		// decoded value, so a structured consumer is unaffected either way.
		if l := it.Record.List(k); len(l) > 0 {
			raw[k] = strings.Join(l, ", ")
		}
	}
	v := s.view(it)
	// A single record costs one extra walk, which is what CompletionOf already
	// does. A listing uses the batched form instead.
	if it.Type() == corpus.WorkItem {
		if comp, cErr := corpus.CompletionOf(s.Backlog, it.Slug()); cErr == nil {
			c := comp.Counts()
			v.Completion = &c
		}
	}
	return Record{
		View:   v,
		Hash:   it.Hash(),
		Fields: fields,
		Order:  order,
		Raw:    raw,
		Body:   it.Record.Body(),
	}, nil
}

// Units are the record types a caller may name.
var Units = corpus.Units

// The unit names, re-exported so an adapter can say which record type it means
// without importing the engine to do it (ADR-0004). Units above gives the set;
// these give the members.
const (
	WorkItem    = corpus.WorkItem
	Outcome     = corpus.Outcome
	Task        = corpus.Task
	Decision    = corpus.Decision
	Exploration = corpus.Exploration
)

// Skip is a record that could not be read.
type Skip struct {
	Path string
	Err  error
}

// Duplicate is one key held by more than one record.
type Duplicate struct {
	Key   string
	Paths []string
}

// stampOf reads a {by, at} field, or an empty stamp where there is none.
func stampOf(it corpus.Item, key string) Stamp {
	by, at := it.Record.Stamp(key)
	return Stamp{By: by, At: at}
}

// rankOf reads a record's rank, or empty where it has none.
func rankOf(it corpus.Item) string {
	r, _ := it.Record.Get("rank")
	return r
}

// byWorkOrder sorts a listing the way the work is meant to be done: by where a
// record stands on the ladder, then by its position among the records standing
// with it, then by name.
//
// **The group comes from `workflow_status`, not from the rank string.**
// ADR-0005 says rank orders records *within* a status, and reading the group
// out of the rank made the rank the only thing that said which status a record
// was ordered within --- so a record with no rank fell below every status
// instead of to the back of its own, and one listing showed `captured` work
// above `in_progress` work with nothing explaining why.
//
// The position half is still compared as text, which is the whole reason it is
// zero-padded (ADR-0005, spec.md §9.6) --- text order and numeric order are the
// same order, so there is no arithmetic here to disagree with anything else
// that sorts the field.
//
// A record nobody has placed sorts to the back of its own status. A rank is a
// position somebody chose, so an unplaced record must not outrank a placed one
// --- but it must not leave its status either, which is the whole correction.
func (s *Session) byWorkOrder(views []View) {
	sort.SliceStable(views, func(i, j int) bool {
		a, b := views[i], views[j]
		if ga, gb := s.groupOf(a), s.groupOf(b); ga != gb {
			return ga < gb
		}
		pa, pb := positionOf(a.Rank), positionOf(b.Rank)
		if pa != pb {
			if pa == "" {
				return false
			}
			if pb == "" {
				return true
			}
			return pa < pb
		}
		return a.Name < b.Name
	})
}

// groupOf is where a record stands on its own unit's ladder.
//
// A status the vocabulary no longer carries has no ordinal, and neither does a
// record whose type declares no status at all. Both sort after every known
// status rather than being refused: configuration is a file people may edit
// (principles.md), a renamed status is reported as drift elsewhere
// (spec.md §5.2), and a listing that dropped or rejected those records would
// hide the very thing the drift report is telling somebody to fix.
func (s *Session) groupOf(v View) int {
	if n, ok := s.Config.LadderFor(v.Type).Ordinal(v.Status); ok {
		return n
	}
	return unplacedGroup
}

// unplacedGroup sorts after every ordinal a ladder can hold --- ordinalDigits
// caps a ladder at 999 statuses (corpus/rank.go), so this is past the end of
// any vocabulary rather than a number somebody might configure.
const unplacedGroup = 1000

// positionOf is the ordering half of a rank, or empty where the record has no
// rank or one that does not parse.
//
// A malformed rank is treated as absent for ordering and reported elsewhere:
// guessing at half of one would put a record somewhere nobody chose, which is
// worse than the back of its own status.
func positionOf(rank string) string {
	if rank == "" {
		return ""
	}
	_, pos, err := corpus.SplitRank(corpus.Rank(rank))
	if err != nil {
		return ""
	}
	return string(pos)
}
