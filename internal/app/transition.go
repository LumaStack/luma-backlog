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
	// Force proceeds past the one refusal a transition carries. Never set by
	// default, and never set on a caller's behalf: overriding a refusal is a
	// decision that belongs to whoever is accountable (ADR-0008).
	Force bool
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
	// Tasks carry a workflow status too, and their own ladder (`todo`,
	// `in_progress`, `closed`). They transition the same way --- applyStatus
	// already writes rank only for work items, since only work items are
	// ranked.
	if it.Type() != corpus.WorkItem && it.Type() != corpus.Task {
		return nil, UsageError(
			"%s is a %s: only work items and tasks carry a workflow status",
			it.Name(), it.Type())
	}
	isWorkItem := it.Type() == corpus.WorkItem

	// Optimistic concurrency, the same contract `set` offers (§6.3). Conflict
	// means re-read and retry, which is different advice from "something
	// broke" — and it is the distinction a retrying agent depends on.
	if req.IfUnchanged != "" && it.Hash() != req.IfUnchanged {
		return nil, ConflictError(
			"%s changed since you read it — re-read and retry\n  you saw:  %s\n  it is now: %s",
			it.Path, shortHash(req.IfUnchanged), shortHash(it.Hash()))
	}

	var advice, forced []string
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
	if isWorkItem && req.To == terminal && terminal != "" {
		return nil, UsageError(
			"%q is terminal — use: work-item close %s <completed|rejected|canceled|superseded>",
			req.To, req.Ref)
	}

	// Two refusals, both narrow, and both take --force. A refusal that cannot
	// be overridden is the tool holding an opinion; one that can is the tool
	// making somebody say they meant it (spec.md §5.0).
	counts, kind := 0, ""
	if isWorkItem {
		var err error
		if counts, err = s.outcomeCounts(it); err != nil {
			return nil, err
		}
		kind, _ = it.Record.Get("kind")
	}

	// Leaving the pile with `kind: idea` still on it. The type defines an idea
	// as "a classification that becomes one of the others" --- transitional by
	// construction --- so carrying one past the first gate files unformed work
	// beside formed work with nothing marking the difference.
	if isWorkItem && s.leavesThePile(it, from, req.To) && kind == "idea" {
		if !req.Force {
			return nil, RefusedError(
				"%s is still an idea, and an idea is a classification on its way to "+
					"something else.\n\nResolve it --- set %s kind=<defect|request|inquiry|change> "+
					"--- or pass --force.", req.Ref, req.Ref)
		}
		forced = append(forced,
			"selected while still an idea --- unformed work is now queued beside formed work")
	}

	// Starting work nobody can tell is finished. Deliberately "does one exist"
	// rather than "are they good": any opinion beyond zero would be the tool
	// holding a view about how work gets defined.
	if isWorkItem && req.To == s.startedStatusFor(it) && counts == 0 {
		if !req.Force {
			return nil, RefusedError(
				"%s has no outcomes, so nothing says when it is finished.\n\n"+
					"Write one --- outcome new \"<what must be true>\" -w %s --- or pass --force.",
				req.Ref, req.Ref)
		}
		forced = append(forced,
			"started with no outcomes --- nothing says when it is finished")
	}

	// Leaving the shaping rung is where the work is supposed to have been
	// worked out. Warned rather than refused: not all work is equal, and some
	// is trivial or urgent.
	if isWorkItem && from == s.shapingStatusFor(it) {
		tasks, err := s.taskCount(it)
		if err != nil {
			return nil, err
		}
		if missing := missingWork(counts, tasks); missing != "" {
			advice = append(advice,
				"leaving "+s.shapingStatusFor(it)+" with "+missing+" --- "+req.Ref+
					" was shaped without them, and whoever picks it up pays for that")
		}
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
	if isWorkItem && reopening && strings.TrimSpace(req.Reason) == "" {
		advice = append(advice,
			"reopening "+req.Ref+" and nothing records why --- pass --reason; "+
				"the closed entry stays, and no field holds the un-ending")
	}

	// The reason goes to the journal rather than onto the record. A work item's
	// journal is already the place reasoning lives, it is append-only so a
	// second crossing does not overwrite the first, and a field would hold only
	// the most recent — which is the half of the history worth least.
	// A forced crossing is written down whether or not a reason was given.
	// Announcing it on stderr tells whoever is at the terminal; the journal is
	// what lets anybody ask later how often this happens and what it cost ---
	// which is the same reason a skipped gate gets a line.
	for _, f := range forced {
		if _, err := s.Journal(JournalRequest{
			WorkItem: s.journalTargetFor(it),
			Line:     "FORCED " + from + " → " + req.To + ": " + f,
		}); err != nil {
			return nil, err
		}
		advice = append(advice, "forced: "+f+" --- recorded in the journal")
	}

	journaled := false
	if reason := strings.TrimSpace(req.Reason); reason != "" {
		if _, err := s.Journal(JournalRequest{
			WorkItem: s.journalTargetFor(it),
			Line:     it.Name() + " " + from + " → " + req.To + ": " + reason,
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

// outcomeCounts is how many live outcomes a work item has. Zero is the only
// number this file cares about.
func (s *Session) outcomeCounts(it corpus.Item) (int, error) {
	c, err := corpus.CompletionOf(s.Backlog, it.Slug())
	if err != nil {
		return 0, FailureError("%w", err)
	}
	return c.Counts().Live, nil
}

// startedStatusFor is the rung that means somebody is working on it.
//
// Derived as the one before terminal, the same way TerminalStatusFor is
// derived as the last: the ladder is ordered, so the rung nothing follows but
// the ending is where work is being done. A project that renames `in_progress`
// keeps working.
//
// **This is the weakest inference in the file, and it is worth saying so.**
// `workflow-status.md` is explicit that the tool attaches no meaning to any
// value — so nothing actually tells it which rung means started, and a project
// that adds a rung between `in_progress` and `closed` moves this silently. The
// honest fix is for the ladder to say which rung is which, and that does not
// exist yet.
func (s *Session) startedStatusFor(it corpus.Item) string {
	l := s.Config.LadderFor(it.Type())
	if len(l.Statuses) < 2 {
		return ""
	}
	return l.Statuses[len(l.Statuses)-2]
}


// leavesThePile reports a crossing of the first gate: out of the rung where
// things wait to be judged, into the pipeline where they have been chosen.
func (s *Session) leavesThePile(it corpus.Item, from, to string) bool {
	pile := s.Config.DefaultStatusFor(it.Type())
	return from == pile && to != pile
}

// shapingStatusFor is the rung where the work gets worked out.
//
// Positional, and carrying the same weakness as startedStatusFor: nothing in
// the ladder says which rung shapes work, so this counts from the pile. The
// default ladder puts it third --- captured, unprepared, preparing --- and a
// project that inserts a rung before it moves this silently.
func (s *Session) shapingStatusFor(it corpus.Item) string {
	l := s.Config.LadderFor(it.Type())
	if len(l.Statuses) < 3 {
		return ""
	}
	return l.Statuses[2]
}

// taskCount is how many tasks a work item has. Zero is the only number that
// matters here.
func (s *Session) taskCount(it corpus.Item) (int, error) {
	items, _, err := corpus.List(s.Backlog, corpus.Filter{Unit: corpus.Task, WorkItem: it.Slug()})
	if err != nil {
		return 0, FailureError("%w", err)
	}
	return len(items), nil
}

// missingWork names what a shaped work item does not have, or nothing.
func missingWork(outcomes, tasks int) string {
	switch {
	case outcomes == 0 && tasks == 0:
		return "no outcomes and no tasks"
	case outcomes == 0:
		return "no outcomes"
	case tasks == 0:
		return "no tasks"
	}
	return ""
}

// journalTargetFor is whose journal a line belongs in. Only work items have
// one, so a task's reasoning goes to its parent --- which is where somebody
// reading the work would look for it. The path already carries the parent,
// since children nest under the work item they belong to (spec.md §7.2).
func (s *Session) journalTargetFor(it corpus.Item) string {
	if it.Type() == corpus.WorkItem {
		return it.Slug()
	}
	const under = "work-items/"
	rest := it.Path
	if i := strings.Index(rest, under); i >= 0 {
		rest = rest[i+len(under):]
	}
	if i := strings.Index(rest, "/"); i > 0 {
		return rest[:i]
	}
	return ""
}
