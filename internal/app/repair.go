package app

import (
	"sort"

	"github.com/lumastack/luma-backlog/internal/corpus"
)

// RepairRequest asks for every work item's rank to be recomputed.
type RepairRequest struct {
	// DryRun reports what would change and writes nothing.
	DryRun bool
}

// RankChange is one record's rank before and after.
type RankChange struct {
	Path string
	Name string
	// From is empty where the record had no rank at all, which is the case
	// this exists for.
	From string
	To   string
}

// RepairResult is what repair found and what it did.
type RepairResult struct {
	Examined int
	Changed  []RankChange
	// Unplaceable names records whose status carries no ordinal, so no rank
	// can be computed for them. Reported, never refused.
	Unplaceable []string
	DryRun      bool
	Observations
}

// RepairRanks renumbers every work item so that all of them carry a rank and
// every prefix agrees with the status the record declares.
//
// **A record with no rank has no position among its peers**, and nothing
// downstream can tell "nobody has placed this" from "placed last". ADR-0005
// promised this command and it did not exist, which is why most of this
// project's own corpus was unranked
// (.luma/backlog/work-items/BACK-0095-there-is-no-unranked-work).
//
// **Records are numbered in the order they were created**, within each status.
// Creation order is already recorded, it is meaningful, and it is the same for
// everybody --- so this is a pure function of the corpus. Two actors repairing
// the same state independently produce byte-identical output, which turns a
// whole-corpus rewrite from a guaranteed merge conflict into one that resolves
// itself. Re-running it changes nothing.
//
// **It renumbers rather than filling gaps.** Preserving existing positions and
// slotting unranked records around them would make the result depend on the
// order repairs happened in, which is the property above. Ranks are cheap to
// restate and an ordering somebody chose is restated by ranking again.
//
// **Never automatic.** A corpus rewrite that happens because a binary was
// upgraded is the least recoverable thing available, so this is a command
// somebody runs, and `--dry-run` shows the whole change first.
func (s *Session) RepairRanks(req RepairRequest) (*RepairResult, error) {
	items, skipped, err := corpus.List(s.Backlog, corpus.Filter{Unit: corpus.WorkItem})
	if err != nil {
		return nil, FailureError("%w", err)
	}

	byStatus := map[string][]corpus.Item{}
	for _, it := range items {
		status, _ := it.Record.Get("workflow_status")
		byStatus[status] = append(byStatus[status], it)
	}

	out := &RepairResult{
		Examined:     len(items),
		DryRun:       req.DryRun,
		Observations: Observations{Skipped: skips(skipped), Duplicates: s.duplicateKeys()},
	}

	ladder := s.Config.LadderFor(corpus.WorkItem)
	// Statuses in a fixed order, so a dry run reads the same twice.
	statuses := make([]string, 0, len(byStatus))
	for status := range byStatus {
		statuses = append(statuses, status)
	}
	sort.Strings(statuses)

	for _, status := range statuses {
		group := byStatus[status]
		ordinal, ok := ladder.Ordinal(status)
		if !ok {
			for _, it := range group {
				out.Unplaceable = append(out.Unplaceable, it.Name())
			}
			continue
		}
		byCreation(group)
		positions := corpus.PositionsFor(len(group))
		for i, it := range group {
			want := string(corpus.MakeRank(ordinal, positions[i]))
			have, _ := it.Record.Get("rank")
			if have == want {
				continue
			}
			out.Changed = append(out.Changed, RankChange{
				Path: it.Path, Name: it.Name(), From: have, To: want,
			})
			if req.DryRun {
				continue
			}
			it.Record.Set("rank", want)
			if err := s.write(it); err != nil {
				return nil, err
			}
		}
	}
	return out, nil
}

// byCreation orders records the way they were created, oldest first.
//
// A record with no created stamp --- written before stamps existed, or by hand
// --- sorts after every record that has one, then by name. It cannot be placed
// in the sequence honestly, and putting it at the end is the only answer that
// does not invent a time for it.
func byCreation(items []corpus.Item) {
	sort.SliceStable(items, func(i, j int) bool {
		_, a := items[i].Record.Stamp("created")
		_, b := items[j].Record.Stamp("created")
		if a != b {
			if a == "" {
				return false
			}
			if b == "" {
				return true
			}
			return a < b
		}
		return items[i].Name() < items[j].Name()
	})
}

// write stamps a record as modified and puts it back on disk.
func (s *Session) write(it corpus.Item) error {
	if err := it.Record.SetRaw("modified",
		"{by: "+s.Env.Actor.String()+", at: "+s.Env.Now()+"}"); err != nil {
		return FailureError("%w", err)
	}
	out, err := it.Record.Bytes()
	if err != nil {
		return FailureError("serializing %s: %w", it.Path, err)
	}
	if err := s.Backlog.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return FailureError("%w", err)
	}
	return nil
}
