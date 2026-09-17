package app

import (
	"path/filepath"

	"github.com/lumastack/luma-backlog/internal/corpus"
)

// CreateRequest asks for a record.
type CreateRequest struct {
	Unit  string
	Title string
	// WorkItem names the work item this belongs to. Empty means derive it
	// from where the caller was standing.
	WorkItem string
	// WorkItemGiven distinguishes "not supplied" from "supplied as empty",
	// which is what lets a decision's level be stated rather than inferred.
	WorkItemGiven bool
	// Kind classifies a work item by what it produces.
	Kind string
	// Project states that a decision is a standing rule.
	Project bool
	// Description is prose about the record, written at creation.
	Description string
}

// CreateResult describes the record.
type CreateResult struct {
	Path string
	// Created is false when the record already existed. Creation is
	// idempotent by name (docs/spec.md §9.5).
	Created bool
	// Unclassified is true when a work item was created carrying no kind.
	// Reported, never refused — see Create.
	Unclassified bool
}

// Create makes a record.
func (s *Session) Create(req CreateRequest) (*CreateResult, error) {
	workItem := req.WorkItem
	if workItem == "" {
		workItem = s.workItemFromWorkingDir()
	}

	// A decision's level is stated, never inferred. The working directory can
	// say WHICH work item, and it must not be allowed to say WHICH LEVEL: a
	// person at a terminal is usually standing where they are working, and an
	// agent runs from the repository root whatever it is doing, so the context
	// it would infer from is a constant. Every decision in this repository is
	// project-level because every one of them was created from the root.
	//
	// The cost of being wrong is asymmetric. A decision filed at the wrong
	// level is not visibly broken; it is simply somewhere nobody looks.
	if req.Unit == corpus.Decision {
		if req.Project && req.WorkItemGiven {
			return nil, UsageError("--project and --work-item say different levels: pass one")
		}
		if !req.Project && !req.WorkItemGiven {
			return nil, UsageError("a decision needs its level stated:\n" +
				"  --work-item <slug>  it belongs to that work item, and is a point-in-time record\n" +
				"  --project           it is a standing rule for the project\n" +
				"Most decisions are work item decisions. Promotion is a separate act.")
		}
		if req.Project {
			// Stated as project-level, so the working directory does not get
			// to attach it to whatever the caller happened to be standing in.
			workItem = ""
		}
	}

	if req.Kind != "" && req.Unit != corpus.WorkItem {
		return nil, UsageError("--kind classifies a work item; %s does not take one", req.Unit)
	}

	rank, err := s.rankAtCreation(req.Unit)
	if err != nil {
		return nil, err
	}

	res, err := corpus.Create(s.Backlog, s.Config, s.Env, corpus.Spec{
		Unit:        req.Unit,
		Title:       req.Title,
		WorkItem:    workItem,
		Kind:        req.Kind,
		Description: req.Description,
		Rank:        rank,
	})
	if err != nil {
		return nil, UsageError("%w", err)
	}

	// Avoided, not denied. Blank is honest when nobody has looked yet — an
	// issue synced from elsewhere, a note taken in a hurry — and it is the
	// wrong answer the rest of the time, because whoever writes a record
	// usually knows what they are holding. Left free, blank becomes the path
	// of least resistance and the field stops meaning anything.
	//
	// So: say so and continue (docs/spec.md §5.0). Nothing is refused.
	unclassified := res.Created && req.Unit == corpus.WorkItem && req.Kind == ""

	return &CreateResult{Path: res.Path, Created: res.Created, Unclassified: unclassified}, nil
}

// workItemFromWorkingDir reads the work item from where the command was run,
// so somebody working inside one does not have to name it.
func (s *Session) workItemFromWorkingDir() string {
	rel, err := filepath.Rel(s.Root, s.WorkingDir)
	if err != nil {
		return ""
	}
	return corpus.WorkItemFromPath(filepath.ToSlash(rel))
}

// rankAtCreation is the rank a new record is written with.
//
// **Every work item carries a rank from the moment it exists.** There is no
// unranked state: a record with no rank has no position among its peers, and
// nothing downstream can tell "nobody has placed this" from "placed last" ---
// so the question is answered once, here, rather than by every reader
// (.luma/backlog/work-items/WORK-0095-there-is-no-unranked-work).
//
// It lands at the back of the default status. Arriving says nothing about a
// record relative to the ones already there, and an unconsidered record must
// not outrank a considered one (ADR-0005).
//
// **It asks the allocator for a position rather than computing one.** A value
// derived from the record alone --- its key, its timestamp --- needs no peer
// read, and cannot place the record behind one somebody ranked `--last`, since
// an explicit rank allocates from the observed maximum and can exceed any
// number a key would give. So the peers get read, which is also what makes
// this survive a change to how positions are allocated: the scheme lives in
// corpus.Between and this calls it.
//
// Two sessions creating at the same moment may compute the same position. That
// is a tie rather than a collision --- both records exist, neither is lost, and
// a listing breaks the tie by name (byWorkOrder).
func (s *Session) rankAtCreation(unit string) (string, error) {
	if unit != corpus.WorkItem {
		return "", nil
	}
	status := s.Config.DefaultStatusFor(unit)
	ordinal, ok := s.Config.LadderFor(unit).Ordinal(status)
	if !ok {
		// The default status carries no ordinal, which means configuration
		// somebody is entitled to edit disagrees with itself. Creating an
		// unranked record is better than refusing to create one at all, and
		// the drift is reported wherever records are read (spec.md §5.2).
		return "", nil
	}
	peers, err := s.rankedPeers("", status)
	if err != nil {
		return "", err
	}
	pos, err := corpus.Between(edges(peers, false))
	if err != nil {
		return "", FailureError("%w", err)
	}
	return string(corpus.MakeRank(ordinal, pos)), nil
}
