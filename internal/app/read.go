package app

import (
	"strings"

	"github.com/lumastack/luma-backlog/internal/corpus"
)

// Filter narrows a listing.
type Filter struct {
	Unit     string
	WorkItem string
	Status   string
	Kind     string
}

// Get reads one record. A reference is a slug, a path, or an unambiguous
// prefix; an ambiguous one is an error rather than a guess (docs/spec.md §9.1).
func (s *Session) Get(ref string) (Record, error) {
	it, err := corpus.Resolve(s.Backlog, ref)
	if err != nil {
		return Record{}, resolveError(err)
	}
	return s.record(it)
}

// ListResult is a listing and what was noticed while producing it.
type ListResult struct {
	Items []View
	Observations
}

// List reads many records, narrowed by the filter.
//
// Matching nothing is not an error: an empty backlog and an over-narrow filter
// are both ordinary, and failing would make a caller treat "none yet" as a
// fault (docs/spec.md §9.3).
func (s *Session) List(f Filter) (*ListResult, error) {
	if f.Unit != "" && !corpus.IsUnit(f.Unit) {
		return nil, UsageError("unknown unit %q: expected one of %s",
			f.Unit, strings.Join(corpus.Units, ", "))
	}
	items, skipped, err := corpus.List(s.Backlog, corpus.Filter{
		Unit:     f.Unit,
		WorkItem: f.WorkItem,
		Status:   f.Status,
		Kind:     f.Kind,
	})
	if err != nil {
		return nil, FailureError("%w", err)
	}
	// One walk for every work item's counts, attached after the loop. Calling
	// CompletionOf per item would read the corpus once per work item.
	completions, err := corpus.Completions(s.Backlog)
	if err != nil {
		return nil, FailureError("%w", err)
	}
	views := make([]View, 0, len(items))
	for _, it := range items {
		v := s.view(it)
		if it.Type() == corpus.WorkItem {
			c := completions[it.Slug()].Counts()
			v.Completion = &c
		}
		views = append(views, v)
	}
	byWorkOrder(views)
	return &ListResult{
		Items: views,
		Observations: Observations{
			Skipped:    skips(skipped),
			Duplicates: s.duplicateKeys(),
			Drifted:    s.statusDrift(items),
		},
	}, nil
}
