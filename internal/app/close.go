package app

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/corpus"
)

// CloseRequest ends a work item, recording why.
type CloseRequest struct {
	// Ref names the work item.
	Ref string
	// As is how the work ended — one of corpus.CloseReasons.
	As string
	// Force closes past every refusal. It never touches the outcomes: the
	// count still computes what it computed, so a reader sees a completed
	// record whose own arithmetic disagrees with it --- which is the truth.
	// Announced, and written to the journal.
	Force bool
	// Reason is prose: why, in the closer's words. Optional, and free text —
	// the enum says which ending, this says anything the enum cannot.
	Reason string
}

// CloseResult describes what was closed.
type CloseResult struct {
	Path   string
	Reason string
	// Retired counts outcomes excluded from the arithmetic because they were
	// archived. Worth saying: it explains a count that would otherwise look
	// wrong.
	Retired int
	// Advice is what the caller should hear but is not stopped by.
	Advice []string
	Observations
}

// CloseWorkItem ends a work item.
//
// This carries the tool's only refusal, and it holds a caller to their OWN
// declarations rather than to an opinion of its own (docs/spec.md §5.0).
func (s *Session) CloseWorkItem(req CloseRequest) (*CloseResult, error) {
	if req.As == "" {
		return nil, UsageError("a disposition is required: %s\n\n"+
			"Closing silently is how a backlog loses its own history — cancelled work\n"+
			"and completed work look identical afterwards.", reasonList())
	}
	if !corpus.IsCloseReason(req.As) {
		return nil, UsageError("unknown disposition %q: expected %s", req.As, reasonList())
	}

	it, err := corpus.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, resolveError(err)
	}
	if it.Type() != corpus.WorkItem {
		return nil, UsageError("%s is a %s — close applies to a work item", it.Slug(), it.Type())
	}

	c, err := corpus.CompletionOf(s.Backlog, it.Slug())
	if err != nil {
		return nil, FailureError("%w", err)
	}

	// Only a disposition that claims success is gated. Cancelling a work item
	// with open tasks is the ordinary case --- it is what being cancelled
	// means --- and gating it would make it impossible to stop work for being
	// unfinished (spec.md §5.3.1).
	var forced []string
	if corpus.CloseReason(req.As).GatedOnCompletion() {
		// Refused for want of an answer rather than on an opinion. An outcome
		// that cannot be read is missing from the count, so it can never be
		// counted as failing, and "completed" would come out clean on evidence
		// nobody has seen. That is the one thing this design cannot allow: a
		// wrong answer that looks exactly like a right one.
		//
		// Every other reason closes freely. None of them claims the work
		// succeeded, so none of them needs a count — which is also the way out
		// when a file is beyond repair.
		if len(c.Skipped) > 0 && !req.Force {
			var b strings.Builder
			fmt.Fprintf(&b, "%s cannot be completed: %d outcome(s) could not be read, so there is no count.\n",
				it.Slug(), len(c.Skipped))
			for _, sk := range c.Skipped {
				fmt.Fprintf(&b, "  %s: %v\n", sk.Path, sk.Err)
			}
			b.WriteString("\nAn unreadable outcome might be failing, and nothing here can tell.\n")
			b.WriteString("Repair the file, or close with a reason that claims nothing about evidence.")
			return nil, RefusedError("%s", b.String())
		}

		if len(c.Skipped) > 0 {
			forced = append(forced, plural(len(c.Skipped), "outcome")+" could not be read")
		}
		if !c.Complete() && len(c.Live) == 0 && !req.Force {
			return nil, RefusedError(
				"%s has no outcomes, so there is nothing that says it was completed.\n"+
					"Declare what done means, or close with a different disposition.", it.Slug())
		}
		if len(c.Live) == 0 {
			forced = append(forced, "no outcomes at all")
		}
		if len(c.Unpassing) > 0 && !req.Force {
			var names []string
			for _, o := range c.Unpassing {
				names = append(names, "  "+o.Slug())
			}
			return nil, RefusedError(
				"%s cannot be completed: %d of %d outcomes are not proven.\n%s\n\n"+
					"Verify them, or close with a different disposition. Abandoning one\n"+
					"records why it is unmet and does not clear this — a completed close\n"+
					"over an unmet outcome needs --force, and the count will say so.",
				it.Slug(), len(c.Unpassing), len(c.Live), strings.Join(names, "\n"))
		}
		if len(c.Unpassing) > 0 {
			forced = append(forced,
				fmt.Sprintf("%d of %d outcomes not proven", len(c.Unpassing), len(c.Live)))
		}

		// Last, because it is bookkeeping where the checks above are about
		// whether the work is done at all. A caller with both problems should
		// hear about the outcomes first.
		openTasks, terr := s.openTaskCount(it)
		if terr != nil {
			return nil, terr
		}
		if openTasks > 0 {
			if !req.Force {
				return nil, RefusedError(
					"%s cannot be completed: %s never reached a terminal status.\n\n"+
						"Close them --- any reason is fine, they may have failed or been\n"+
						"cancelled --- or pass --force.",
					req.Ref, plural(openTasks, "task"))
			}
			forced = append(forced, plural(openTasks, "task")+" still open")
		}
	}

	// Status and rank move together, always (ADR-0005). Closing is a status
	// change like any other, so it goes through the one operation that writes
	// both rather than setting the field itself.
	if err := s.applyStatus(it, "closed"); err != nil {
		return nil, err
	}
	// `as` rather than `reason`, because the field holds the disposition and
	// the prose is a separate thing somebody may or may not have said
	// (ADR-0007). Every record closed before this carries the old spelling;
	// migrating them is work-items/WORK-0037.
	closed := fmt.Sprintf("{on: %s, as: %s, by: %s}", s.Env.Today(), req.As, s.Env.Actor.String())
	if req.Reason != "" {
		closed = fmt.Sprintf("{on: %s, as: %s, by: %s, reason: %s}",
			s.Env.Today(), req.As, s.Env.Actor.String(), yamlQuote(req.Reason))
	}
	if err := it.Record.SetRaw("closed", closed); err != nil {
		return nil, FailureError("%w", err)
	}
	if err := it.Record.SetRaw("modified",
		"{by: "+s.Env.Actor.String()+", at: "+s.Env.Now()+"}"); err != nil {
		return nil, FailureError("%w", err)
	}

	out, err := it.Record.Bytes()
	if err != nil {
		return nil, FailureError("%w", err)
	}
	if err := s.Backlog.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return nil, FailureError("%w", err)
	}

	// A task left open and ready to start under a closed work item advertises
	// work nobody can pick up. Warned, never refused, and never auto-closed:
	// closing the stragglers would invent a disposition nobody chose, which is
	// exactly what --force refuses to do to outcomes. A work item is judged on
	// its outcomes and on nothing else (spec.md §2.4).
	open, err := s.openTaskCount(it)
	if err != nil {
		return nil, err
	}
	var advice []string
	for _, f := range forced {
		if _, jerr := s.Journal(JournalRequest{
			WorkItem: it.Slug(),
			Line:     "FORCED close as " + req.As + ": " + f,
		}); jerr != nil {
			return nil, jerr
		}
		advice = append(advice, "forced: "+f+" --- recorded in the journal")
	}
	if open > 0 && len(forced) == 0 {
		advice = append(advice,
			plural(open, "task")+" on "+req.Ref+" never reached a terminal status --- "+
				"they now advertise work nobody can pick up, and nothing here closed them")
	}

	return &CloseResult{
		Path:    it.Path,
		Reason:  req.As,
		Retired: len(c.Retired),
		Advice:  advice,
		Observations: Observations{
			// Only on a path that did not refuse. A refusal already names the
			// files it refused over, and saying the same path twice teaches a
			// reader to skim the part that matters.
			Skipped: skips(c.Skipped),
			// Closing is where acting on the wrong record costs most: it is
			// the one command that writes a terminal state, and a key naming
			// two records means a citation of this close could land on either.
			Duplicates: s.duplicateKeys(),
		},
	}, nil
}

func reasonList() string {
	var s []string
	for _, r := range corpus.CloseReasons {
		s = append(s, string(r))
	}
	return strings.Join(s, ", ")
}

// yamlQuote makes free prose safe inside an inline mapping. A colon or a hash
// in somebody's sentence would otherwise end the value early or start a
// comment, and the sentence most likely to contain one is the sentence
// explaining a refusal.
func yamlQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}

// openTaskCount is how many of a work item's tasks have not reached the
// terminal rung of their own ladder.
func (s *Session) openTaskCount(it corpus.Item) (int, error) {
	items, _, err := corpus.List(s.Backlog, corpus.Filter{Unit: corpus.Task, WorkItem: it.Slug()})
	if err != nil {
		return 0, FailureError("%w", err)
	}
	terminal := s.Config.TerminalStatusFor(corpus.Task)
	open := 0
	for _, t := range items {
		if t.Status(s.Config.DefaultStatusFor(corpus.Task)) != terminal {
			open++
		}
	}
	return open, nil
}
