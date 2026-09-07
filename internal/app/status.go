package app

import (
	"github.com/lumastack/luma-backlog/internal/corpus"
)

// applyStatus moves a record to a workflow status and rewrites its rank to
// match, in one operation.
//
// **`workflow_status` and `rank` are always written together** (ADR-0005).
// A stale prefix is wrong data, and the whole ordering design rests on the two
// agreeing --- so nothing here offers a way to write one without the other. Any
// caller changing a status goes through this, which is what makes the
// invariant unable to be expressed otherwise rather than merely required.
//
// The record is re-enqueued at the back of the destination. A rank is a
// position in a queue, and leaving the queue does not carry the position with
// you. This also preserves order in the ordinary case --- advancing records in
// rank order lands them in the same relative order, because each arrives
// behind the last.
func (s *Session) applyStatus(it corpus.Item, status string) error {
	it.Record.Set("workflow_status", status)

	// Only work items are ranked today. ADR-0005 names `rank` for the field it
	// changes rather than a gesture, so it survives tasks being ranked later;
	// until they are, a task carries a status and no rank.
	if it.Type() != corpus.WorkItem {
		return nil
	}

	ordinal, ok := s.Config.LadderFor(it.Type()).Ordinal(status)
	if !ok {
		return UsageError(
			"%q is not a status this project carries --- rank has no ordinal for it", status)
	}

	peers, err := s.rankedPeers(it, status)
	if err != nil {
		return err
	}
	var back corpus.Position
	if len(peers) > 0 {
		back = peers[len(peers)-1].pos
	}
	pos, err := corpus.Between(back, "")
	if err != nil {
		return FailureError("%w", err)
	}
	it.Record.Set("rank", string(corpus.MakeRank(ordinal, pos)))
	return nil
}

// StatusDrift is a record whose rank prefix disagrees with its status.
//
// Configuration is a file people are entitled to edit (principles.md), so this
// is observed and never refused (spec.md §5.2). It happens when a status is
// renamed, inserted or reordered by hand, which changes what ordinal a status
// carries without touching any record.
type StatusDrift struct {
	Path string
	// Status the record declares, and the ordinal that status carries now.
	Status  string
	Ordinal int
	// Rank as stored, whose prefix says otherwise.
	Rank string
}

// statusDrift reports records whose rank disagrees with their status.
func (s *Session) statusDrift(items []corpus.Item) []StatusDrift {
	var out []StatusDrift
	for _, it := range items {
		raw, ok := it.Record.Get("rank")
		if !ok || raw == "" {
			continue
		}
		have, _, err := corpus.SplitRank(corpus.Rank(raw))
		if err != nil {
			continue
		}
		status := it.Status(s.Config.DefaultStatusFor(it.Type()))
		want, known := s.Config.LadderFor(it.Type()).Ordinal(status)
		if !known || have == want {
			continue
		}
		out = append(out, StatusDrift{Path: it.Path, Status: status, Ordinal: want, Rank: raw})
	}
	return out
}
