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
	// Reason is why the work ended — one of corpus.CloseReasons.
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
	Observations
}

// CloseWorkItem ends a work item.
//
// This carries the tool's only refusal, and it holds a caller to their OWN
// declarations rather than to an opinion of its own (docs/spec.md §5.0).
func (s *Session) CloseWorkItem(req CloseRequest) (*CloseResult, error) {
	if req.Reason == "" {
		return nil, UsageError("--reason is required: %s\n\n"+
			"Closing silently is how a backlog loses its own history — cancelled work\n"+
			"and delivered work look identical afterwards.", reasonList())
	}
	if !corpus.IsCloseReason(req.Reason) {
		return nil, UsageError("unknown reason %q: expected %s", req.Reason, reasonList())
	}

	it, err := corpus.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, &Error{Kind: NotFound, Err: err}
	}
	if it.Type() != corpus.WorkItem {
		return nil, UsageError("%s is a %s — close applies to a work item", it.Slug(), it.Type())
	}

	c, err := corpus.CompletionOf(s.Backlog, it.Slug())
	if err != nil {
		return nil, FailureError("%w", err)
	}

	if corpus.CloseReason(req.Reason).GatedOnCompletion() {
		// Refused for want of an answer rather than on an opinion. An outcome
		// that cannot be read is missing from the count, so it can never be
		// counted as failing, and "delivered" would come out clean on evidence
		// nobody has seen. That is the one thing this design cannot allow: a
		// wrong answer that looks exactly like a right one.
		//
		// Every other reason closes freely. None of them claims the work
		// succeeded, so none of them needs a count — which is also the way out
		// when a file is beyond repair.
		if len(c.Skipped) > 0 {
			var b strings.Builder
			fmt.Fprintf(&b, "%s cannot be delivered: %d outcome(s) could not be read, so there is no count.\n",
				it.Slug(), len(c.Skipped))
			for _, sk := range c.Skipped {
				fmt.Fprintf(&b, "  %s: %v\n", sk.Path, sk.Err)
			}
			b.WriteString("\nAn unreadable outcome might be failing, and nothing here can tell.\n")
			b.WriteString("Repair the file, or close with a reason that claims nothing about evidence.")
			return nil, RefusedError("%s", b.String())
		}

		if !c.Complete() && len(c.Live) == 0 {
			return nil, RefusedError(
				"%s has no outcomes, so there is nothing that says it was delivered.\n"+
					"Declare what done means, or close with a different reason.", it.Slug())
		}
		if len(c.Unpassing) > 0 {
			var names []string
			for _, o := range c.Unpassing {
				names = append(names, "  "+o.Slug())
			}
			return nil, RefusedError(
				"%s cannot be delivered: %d of %d outcomes have no evidence.\n%s\n\n"+
					"Verify them, retire the ones that no longer apply, or close with a\n"+
					"different reason.",
				it.Slug(), len(c.Unpassing), len(c.Live), strings.Join(names, "\n"))
		}
	}

	// Status and rank move together, always (ADR-0005). Closing is a status
	// change like any other, so it goes through the one operation that writes
	// both rather than setting the field itself.
	if err := s.applyStatus(it, "closed"); err != nil {
		return nil, err
	}
	if err := it.Record.SetRaw("closed", fmt.Sprintf("{on: %s, reason: %s, by: %s}",
		s.Env.Today(), req.Reason, s.Env.Actor.String())); err != nil {
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

	return &CloseResult{
		Path:    it.Path,
		Reason:  req.Reason,
		Retired: len(c.Retired),
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
