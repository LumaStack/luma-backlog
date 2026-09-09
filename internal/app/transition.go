package app

import (
	"strings"

	"github.com/lumastack/luma-backlog/internal/corpus"
)

// TransitionRequest moves a work item to another workflow status.
//
// A status change is an operation, not a field write. `set` refuses the field
// for the same reason it refuses `rank` (ADR-0005): the caller says where, and
// the tool writes both fields together so no record can be produced whose
// status and rank disagree.
//
// The verb is `transition` and never `move`. In a tool whose records are files
// in a git repository, `move` reads as `mv` — a relocation, which is the one
// operation this design forbids (`spec.md` §9.2).
type TransitionRequest struct {
	Ref string
	// To is the destination rung, named in the configured vocabulary.
	To string
	// IfUnchanged is the hash the caller last saw. When set, a write that
	// would clobber a change it never saw is refused rather than applied.
	IfUnchanged string
	// Reason is why, in the caller's words. Optional, and appended to the work
	// item's journal rather than stored on the record.
	//
	// Optional deliberately. Requiring it would make every crossing carry
	// prose, and prose demanded of somebody who has nothing to say is prose
	// nobody reads. `backlog-move` records that nothing captures the reasoning
	// for crossing a gate today; this is somewhere for it to go when there is
	// some, not an obligation to produce it.
	Reason string
}

// TransitionResult describes what was written.
type TransitionResult struct {
	Path string
	// From and To are the statuses either side of the change. From is the
	// status the record actually held, which is not always what the caller
	// assumed.
	From string
	To   string
	// Rank as written. A transition re-enqueues the record at the back of its
	// destination, so this always changes.
	Rank string
	// Journaled is true when a reason was given and written.
	Journaled bool
	// Advice is what the caller should hear but is not stopped by. Rendered to
	// stderr at exit 0: the crossing happened.
	Advice []string
}

// Transition changes a work item's workflow status, writing status and rank
// together.
//
// Closing is deliberately not reachable here. `close` carries its own
// refusals, its own disposition vocabulary and its own `--force`, and a second
// door into the terminal status would skip all three — which is the defect
// recorded as WORK-0073.
func (s *Session) Transition(req TransitionRequest) (*TransitionResult, error) {
	if req.To == "" {
		return nil, UsageError("say where: transition <work-item> <status>")
	}

	it, err := corpus.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, resolveError(err)
	}
	if it.Type() != corpus.WorkItem {
		return nil, UsageError(
			"%s is a %s: only work items carry a workflow status", it.Name(), it.Type())
	}

	// Optimistic concurrency, the same contract `set` offers (§6.3). Conflict
	// means re-read and retry, which is different advice from "something
	// broke" — and it is the distinction a retrying agent depends on.
	if req.IfUnchanged != "" && it.Hash() != req.IfUnchanged {
		return nil, ConflictError(
			"%s changed since you read it — re-read and retry\n  you saw:  %s\n  it is now: %s",
			it.Path, shortHash(req.IfUnchanged), shortHash(it.Hash()))
	}

	from := it.Status(s.Config.DefaultStatusFor(it.Type()))
	terminal := s.Config.TerminalStatusFor(it.Type())
	reopening := terminal != "" && from == terminal

	ladder := s.Config.LadderFor(it.Type())
	if _, ok := ladder.Ordinal(req.To); !ok {
		return nil, UsageError(
			"%q is not a status this project carries — the ladder runs: %s",
			req.To, strings.Join(ladder.Statuses, " → "))
	}

	// Reaching `closed` through here would bypass the outcome gate, the
	// disposition, the `closed` event and the `--force` requirement — every
	// check `close` exists to perform (spec.md §5.3.1).
	if req.To == terminal && terminal != "" {
		return nil, UsageError(
			"%q is terminal — use: work-item close %s <completed|rejected|canceled|superseded>",
			req.To, req.Ref)
	}

	if err := s.applyStatus(it, req.To); err != nil {
		return nil, err
	}
	rank, _ := it.Record.Get("rank")
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

	// Reopening is the one crossing that should always say why, because
	// nothing else records it: the `closed` list keeps the ending, and no
	// field holds the un-ending. Strongly encouraged rather than required ---
	// a project that wants the hard version can have it, and most do not.
	//
	// Nothing is expected of the ordinary rungs. `captured` to `unprepared`
	// always has the same answer, and a prompt whose answer is always the same
	// teaches people to type past it.
	var advice []string
	if reopening && strings.TrimSpace(req.Reason) == "" {
		advice = append(advice,
			"reopening "+it.Name()+" and nothing records why --- pass --reason; "+
				"the closed entry stays, and no field holds the un-ending")
	}

	// The reason goes to the journal rather than onto the record. A work item's
	// journal is already the place reasoning lives, it is append-only so a
	// second crossing does not overwrite the first, and a field would hold only
	// the most recent — which is the half of the history worth least.
	journaled := false
	if reason := strings.TrimSpace(req.Reason); reason != "" {
		if _, err := s.Journal(JournalRequest{
			WorkItem: it.Slug(),
			Line:     from + " → " + req.To + ": " + reason,
		}); err != nil {
			return nil, err
		}
		journaled = true
	}

	return &TransitionResult{
		Path:      it.Path,
		From:      from,
		To:        req.To,
		Rank:      rank,
		Journaled: journaled,
		Advice:    advice,
	}, nil
}
