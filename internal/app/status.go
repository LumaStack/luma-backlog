package app

import (
	"github.com/lumastack/luma-backlog/internal/config"
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
// **Advancing re-enqueues the record at the back of the destination; going
// backwards puts it at the front** (ADR-0005, amended 2026-09-17,
// .luma/backlog/work-items/WORK-0095-there-is-no-unranked-work).
//
// Advancing to the back preserves order in the ordinary case: several records
// advanced in rank order land in the same relative order, because each arrives
// behind the last.
//
// Going backwards to the front is chosen on which error is recoverable rather
// than on what the move means. Often it means nothing about the record ---
// capacity vanishing sends work back from `todo` and says nothing about the
// work, and a reopened record may be an old defect nobody fixed or work that
// was never really complete. **A record placed too high is visible at the top
// of a listing and gets corrected; one placed too low is invisible and stays
// wrong.** So burying something is left as an act somebody performs --- `rank
// --last` --- rather than something a default does quietly.
//
// A batch sent backwards therefore comes out reversed, each arrival pushing
// the previous one down. That is recorded rather than solved: backwards
// movement is rare, and a bulk drain is arguably a different operation from
// sending one record back.
func (s *Session) applyStatus(it corpus.Item, status string) error {
	// Read before writing: the direction is the difference between where the
	// record is and where it is going, and a caller cannot pass it wrongly if
	// it is never passed.
	from, _ := it.Record.Get("workflow_status")
	it.Record.Set("workflow_status", status)

	// Only work items are ranked today. ADR-0005 names `rank` for the field it
	// changes rather than a gesture, so it survives tasks being ranked later;
	// until they are, a task carries a status and no rank.
	if it.Type() != corpus.WorkItem {
		return nil
	}

	ladder := s.Config.LadderFor(it.Type())
	ordinal, ok := ladder.Ordinal(status)
	if !ok {
		return UsageError(
			"%q is not a status this project carries --- rank has no ordinal for it", status)
	}

	peers, err := s.rankedPeers(it, status)
	if err != nil {
		return err
	}
	pos, err := corpus.Between(edges(peers, regressing(ladder, from, ordinal)))
	if err != nil {
		return FailureError("%w", err)
	}
	it.Record.Set("rank", string(corpus.MakeRank(ordinal, pos)))
	return nil
}

// regressing reports whether a crossing is going back down the ladder.
//
// A record whose current status carries no ordinal is not regressing: a status
// the vocabulary no longer holds gives nothing to compare against, and a first
// status --- a record being ranked for the first time --- is an arrival rather
// than a return.
func regressing(ladder config.Ladder, from string, to int) bool {
	was, ok := ladder.Ordinal(from)
	return ok && to < was
}

// edges is the pair `corpus.Between` allocates inside: the neighbor above the
// new position and the neighbor below it. One side is always empty, because a
// record arriving at a status goes to one end of it and never between two
// records somebody already ordered.
func edges(peers []peer, toTheFront bool) (before, after corpus.Position) {
	if len(peers) == 0 {
		return "", ""
	}
	if toTheFront {
		return "", peers[0].pos
	}
	return peers[len(peers)-1].pos, ""
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
