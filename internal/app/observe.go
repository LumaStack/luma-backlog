package app

import (
	"github.com/lumastack/luma-backlog/internal/backlog"
)

// Observations are things noticed while doing something else. They are
// reported and never refused (docs/spec.md §5.2): a wrong observation gets
// ignored, a wrong refusal stops work.
//
// The layer computes them; an adapter decides how to show them. That split is
// why they are values on a result rather than something written to a stream
// from in here.
type Observations struct {
	// Skipped is every record that could not be read.
	Skipped []Skip
	// Duplicates is every key held by more than one record.
	Duplicates []Duplicate
}

func skips(in []backlog.Skip) []Skip {
	if len(in) == 0 {
		return nil
	}
	out := make([]Skip, 0, len(in))
	for _, s := range in {
		out = append(out, Skip{Path: s.Path, Err: s.Err})
	}
	return out
}

// duplicateKeys runs over the whole work item set rather than over whatever a
// caller asked for: a filtered listing would miss a duplicate outside its
// filter, and a check that only fires when you were already looking in the
// right place is not a check.
func (s *Session) duplicateKeys() []Duplicate {
	items, _, err := backlog.List(s.Backlog, backlog.Filter{Unit: backlog.WorkItem})
	if err != nil {
		return nil
	}
	found := backlog.Duplicates(items)
	if len(found) == 0 {
		return nil
	}
	out := make([]Duplicate, 0, len(found))
	for _, d := range found {
		out = append(out, Duplicate{Key: d.Key, Paths: d.Paths})
	}
	return out
}
