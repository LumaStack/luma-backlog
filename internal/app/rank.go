package app

import (
	"sort"

	"github.com/lumastack/luma-backlog/internal/corpus"
)

// Where a record goes in the order. Exactly one is set.
//
// Named by sequence rather than by the screen: --first and --last stay correct
// when a listing is drawn in reverse, where --top and --bottom would invert
// with it (spec.md §9.6).
type RankPosition int

const (
	RankFirst RankPosition = iota + 1
	RankLast
	RankBefore
	RankAfter
)

// RankRequest asks for a record to be reordered.
//
// The caller says where, never what the ordering key should be: §9.6 makes
// that the tool's job, and ADR-0005 makes `rank` the only way in so the rule
// is not decoration.
type RankRequest struct {
	Ref string
	// Where the record goes. Neighbor is required for Before and After.
	Where    RankPosition
	Neighbor string
}

type RankResult struct {
	Path string
	// Rank as written, for a caller that wants to report it.
	Rank string
}

// Rank reorders one record and writes one file. Its neighbors are untouched:
// a decimal ordering key is what buys that (§9.6).
func (s *Session) Rank(req RankRequest) (*RankResult, error) {
	if req.Where == 0 {
		return nil, UsageError("say where: --first, --last, --before <ref>, or --after <ref>")
	}
	if (req.Where == RankBefore || req.Where == RankAfter) && req.Neighbor == "" {
		return nil, UsageError("--before and --after need a record to sit against")
	}

	it, err := corpus.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, resolveError(err)
	}
	if it.Type() != corpus.WorkItem {
		return nil, UsageError("%s is a %s: only work items are ranked", it.Name(), it.Type())
	}

	status := it.Status(s.Config.DefaultStatusFor(it.Type()))
	ordinal, ok := s.Config.LadderFor(it.Type()).Ordinal(status)
	if !ok {
		return nil, UsageError(
			"%s is at status %q, which the configured ladder does not carry --- rank has no place to put it",
			it.Name(), status)
	}

	// Rank orders records within a status (ADR-0005), so only records sharing
	// this one's status are neighbors. Records at another status are ahead or
	// behind by the ordinal alone and nothing here can change that.
	peers, err := s.rankedPeers(it, status)
	if err != nil {
		return nil, err
	}

	pos, err := s.positionFor(req, peers)
	if err != nil {
		return nil, err
	}

	rank := string(corpus.MakeRank(ordinal, pos))
	it.Record.Set("rank", rank)
	if err := it.Record.SetRaw("modified",
		"{by: "+s.Env.Actor.String()+", at: "+s.Env.Now()+"}"); err != nil {
		return nil, FailureError("%w", err)
	}
	out, err := it.Record.Bytes()
	if err != nil {
		return nil, FailureError("serializing %s: %w", it.Path, err)
	}
	if err := s.Backlog.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return nil, FailureError("%w", err)
	}
	return &RankResult{Path: it.Path, Rank: rank}, nil
}

// peer is a neighbor and the position it occupies.
type peer struct {
	name string
	pos  corpus.Position
}

// rankedPeers returns the record's status-mates that already carry a rank, in
// order. The record being moved is excluded --- it is not its own neighbor.
//
// Records with no rank are skipped rather than assigned one. Ranking one
// record must write one file (§9.6); seeding the whole status here would make
// a reorder a multi-record write, which is the thing the scheme exists to
// avoid.
func (s *Session) rankedPeers(moving corpus.Item, status string) ([]peer, error) {
	items, _, err := corpus.List(s.Backlog, corpus.Filter{Unit: corpus.WorkItem, Status: status})
	if err != nil {
		return nil, FailureError("%w", err)
	}
	var out []peer
	for _, it := range items {
		if it.Path == moving.Path {
			continue
		}
		raw, ok := it.Record.Get("rank")
		if !ok || raw == "" {
			continue
		}
		_, pos, err := corpus.SplitRank(corpus.Rank(raw))
		if err != nil {
			// A malformed rank is reported, not repaired --- but it must not
			// stop a reorder of a different record.
			continue
		}
		out = append(out, peer{name: it.Name(), pos: pos})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].pos < out[j].pos })
	return out, nil
}

func (s *Session) positionFor(req RankRequest, peers []peer) (corpus.Position, error) {
	switch req.Where {
	case RankFirst:
		if len(peers) == 0 {
			return corpus.Between("", "")
		}
		return corpus.Between("", peers[0].pos)

	case RankLast:
		if len(peers) == 0 {
			return corpus.Between("", "")
		}
		return corpus.Between(peers[len(peers)-1].pos, "")
	}

	// --before and --after are relative to a named record, which has to be one
	// of the neighbors: ranking against something at another status would be
	// asking for a position that cannot exist.
	target, err := corpus.Resolve(s.Backlog, req.Neighbor)
	if err != nil {
		return "", resolveError(err)
	}
	at := -1
	for i, p := range peers {
		if p.name == target.Name() {
			at = i
			break
		}
	}
	if at < 0 {
		return "", UsageError(
			"%s is not ranked alongside it --- rank orders records within a status, and these are not at the same one",
			target.Name())
	}

	if req.Where == RankBefore {
		if at == 0 {
			return corpus.Between("", peers[0].pos)
		}
		return corpus.Between(peers[at-1].pos, peers[at].pos)
	}
	if at == len(peers)-1 {
		return corpus.Between(peers[at].pos, "")
	}
	return corpus.Between(peers[at].pos, peers[at+1].pos)
}
