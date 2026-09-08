package app

import (
	"github.com/lumastack/luma-backlog/internal/corpus"
)

// AbandonRequest records that an outcome is no longer required.
type AbandonRequest struct {
	Ref string
	// Reason is prose: why it was given up on. Optional, and the thing a
	// retrospective actually reads.
	Reason string
}

// AbandonResult describes the abandonment.
type AbandonResult struct {
	Path string
	// StillCounted is always true, and is reported so nobody expects
	// otherwise. See Abandon.
	StillCounted bool
	// Verdicts is how many findings the outcome already carries. Abandoning
	// one that was checked is a different act from abandoning one nobody
	// looked at, and the difference is worth seeing at the moment it happens.
	Verdicts int
}

// Abandon records that an outcome is no longer required.
//
// It is a decision, not a finding — which is why `verify` cannot write it. A
// checker determines whether a condition holds; abandoning says the condition
// is no longer asked. One verb for both would let a doer acting as their own
// checker make an unmet outcome disappear by verifying it away (ADR-0007).
//
// **It does not remove the outcome from the count.** Abandoning explains why
// an outcome is unmet; it does not make the work item completable. A close
// still needs `--force`, and the arithmetic still says two of five — which is
// the truth, and is the same reason --force never touches the outcomes.
// Letting abandonment clean the count would make it the cheapest way out of a
// commitment, which is the opposite of what the word is for.
func (s *Session) Abandon(req AbandonRequest) (*AbandonResult, error) {
	it, err := corpus.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, resolveError(err)
	}
	if it.Type() != corpus.Outcome {
		return nil, UsageError("%s is a %s — only an outcome is abandoned", it.Slug(), it.Type())
	}

	entry := map[string]string{"by": s.Env.Actor.String(), "at": s.Env.Now()}
	if req.Reason != "" {
		entry["reason"] = req.Reason
	}
	// A list, like every other event on a record here. It also leaves room for
	// an abandonment to be reversed later by appending rather than by deleting
	// the only evidence it happened.
	if err := appendToList(it.Record, "abandoned", entry); err != nil {
		return nil, FailureError("%w", err)
	}

	out, err := it.Record.Bytes()
	if err != nil {
		return nil, FailureError("%w", err)
	}
	if err := s.Backlog.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return nil, FailureError("%w", err)
	}

	return &AbandonResult{
		Path:         it.Path,
		StillCounted: true,
		Verdicts:     countList(it.Record, "verified"),
	}, nil
}
