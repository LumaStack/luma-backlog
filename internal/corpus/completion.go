package corpus

import (
	"path"
	"strings"

	"github.com/lumastack/luma-backlog/internal/root"
)

// Completion is the arithmetic behind closing a work item.
type Completion struct {
	Live      []Item // outcomes counted toward completion
	Unpassing []Item // of those, the ones with no evidence
	Retired   []Item // archived, and excluded (docs/spec.md §4.4)

	// Skipped is every file under this work item's outcomes that could not be
	// read. Each one is a live outcome for all anybody knows, so the count has
	// lost its denominator and Complete() cannot be true.
	//
	// Scoped to this work item deliberately: a corrupt record belonging to
	// somebody else's work is a real problem and not this caller's, and
	// blocking every close in the repository on it would be a refusal nobody
	// could act on.
	Skipped []Skip
}

// Complete reports whether every live outcome has passed.
//
// False when anything was skipped. "Every live outcome" is a claim about a set,
// and a set read incompletely does not support it — an unreadable file is
// missing from Live, so it can never appear in Unpassing, and the arithmetic
// would come out clean on evidence nobody has seen. Completion is computed
// rather than asserted (docs/spec.md §2.4), and a count over an incomplete read
// is an assertion wearing a count's clothes.
func (c Completion) Complete() bool {
	return len(c.Live) > 0 && len(c.Unpassing) == 0 && len(c.Skipped) == 0
}

// Counts is a completion reduced to numbers, for a reader rather than a
// refusal. Proven is what a person means by "how far along is this".
type Counts struct {
	Proven  int
	Live    int
	Retired int
	Skipped int
}

// Counts reduces a completion to the four numbers worth showing.
func (c Completion) Counts() Counts {
	return Counts{
		Proven:  len(c.Live) - len(c.Unpassing),
		Live:    len(c.Live),
		Retired: len(c.Retired),
		Skipped: len(c.Skipped),
	}
}

// Completions counts every work item's outcomes in one walk.
//
// CompletionOf reads the corpus each time it is called, which is fine for the
// single record `show` reads and quadratic for a listing. A listing gets this
// instead: one walk, grouped by work item.
func Completions(b *root.Backlog) (map[string]Completion, error) {
	items, skipped, err := List(b, Filter{Unit: Outcome})
	if err != nil {
		return nil, err
	}
	out := map[string]Completion{}
	for _, it := range items {
		c := out[it.WorkItem]
		if s, _ := it.Record.Get("stage"); s == "archived" {
			c.Retired = append(c.Retired, it)
		} else {
			c.Live = append(c.Live, it)
			if !it.Record.Has("verified") {
				c.Unpassing = append(c.Unpassing, it)
			}
		}
		out[it.WorkItem] = c
	}
	// An unreadable file where an outcome would be counts against the work
	// item it sits under, the same as it does for a refusal — a work item
	// whose evidence cannot be read is not a work item with less evidence.
	for _, sk := range skipped {
		for wi, c := range out {
			if couldBeOutcomeOf(sk.Path, wi) {
				c.Skipped = append(c.Skipped, sk)
				out[wi] = c
			}
		}
	}
	return out, nil
}

// CompletionOf counts the outcomes of a work item.
//
// Derived by counting, never read from a field. There is nothing to store that
// the verification records do not already say, and a stored copy could
// disagree with them (docs/spec.md §2.4).
func CompletionOf(b *root.Backlog, workItem string) (Completion, error) {
	items, skipped, err := List(b, Filter{Unit: Outcome, WorkItem: workItem})
	if err != nil {
		return Completion{}, err
	}

	var c Completion
	for _, sk := range skipped {
		if couldBeOutcomeOf(sk.Path, workItem) {
			c.Skipped = append(c.Skipped, sk)
		}
	}
	for _, it := range items {
		// A retired outcome is archived rather than deleted, and excluded
		// from the arithmetic — otherwise retiring one could never let a
		// work item close, which is the point of retiring it.
		if s, _ := it.Record.Get("stage"); s == "archived" {
			c.Retired = append(c.Retired, it)
			continue
		}
		c.Live = append(c.Live, it)
		if !it.Record.Has("verified") {
			c.Unpassing = append(c.Unpassing, it)
		}
	}
	return c, nil
}

// CloseReason is why work ended. Only one of them is success, which is why
// the terminal state is "closed" rather than "done" (docs/spec.md §5.3.1).
type CloseReason string

const (
	Completed  CloseReason = "completed"
	Rejected   CloseReason = "rejected"
	Canceled   CloseReason = "canceled"
	Superseded CloseReason = "superseded"
)

// CloseReasons lists them in the order ADR-0007 sets them out.
//
// `completed` rather than `delivered`: the tool cannot observe a handover and
// can compute a count.
//
// `abandoned` is gone. The enum carries what the record cannot derive, and
// stopping without a decision is derivable — from a record that has one and a
// journal that stops. `rejected` earns its place by the same test: never
// crossed the first gate is a statement of intent nothing else holds.
var CloseReasons = []CloseReason{Completed, Rejected, Canceled, Superseded}

// IsCloseReason reports whether a value is one.
func IsCloseReason(s string) bool {
	for _, r := range CloseReasons {
		if CloseReason(s) == r {
			return true
		}
	}
	return false
}

// GatedOnCompletion reports whether a reason requires every outcome to pass.
//
// Only completion is gated. Gating cancellation would make it impossible to stop
// work precisely because it was unfinished — which is the only reason anyone
// ever cancels anything.
func (r CloseReason) GatedOnCompletion() bool { return r == Completed }

// couldBeOutcomeOf reports whether an unreadable file sits where an outcome of
// this work item would.
//
// It has to go by path, because a file that will not parse has no type to read.
// That is the whole difficulty: the record cannot say what it is, so its
// location is the only evidence available.
func couldBeOutcomeOf(rel, workItem string) bool {
	if workItem == "" {
		return false
	}
	dir := path.Join(BundleDir, "work-items", workItem, childDirs[Outcome])
	return strings.HasPrefix(path.Clean(rel), dir+"/")
}
