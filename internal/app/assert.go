package app

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/corpus"
)

// AssertRequest records what a doer claims about an outcome.
type AssertRequest struct {
	Ref string
	// As is the claim — one of corpus.Assertions.
	As string
}

// AssertResult describes the claim.
type AssertResult struct {
	Path string
	// Attempts is how many claims the outcome now carries, including this one.
	// One is a first attempt; more is a second look, which is the fact a single
	// overwritten field could not have represented.
	Attempts int
	// Contradicts is true when a checker has already recorded a verdict. Said
	// out loud, never refused: a doer claiming success over a disproven
	// outcome is exactly the disagreement ADR-0007 exists to keep visible.
	Contradicts bool
}

// Assert appends the doer's claim to an outcome.
//
// It never gates anything. Closing a work item as completed reads the
// checker's verdict and not this — gating on the doer's claim would gate on
// the thing the design distrusts, and would let a doer clear their own work
// (ADR-0007).
func (s *Session) Assert(req AssertRequest) (*AssertResult, error) {
	if req.As == "" {
		return nil, Refuse(Usage, Refusal{
			Problem: "Claim required",
			Detail: []string{
				"an outcome nobody attempted and one somebody tried and failed",
				"look identical otherwise",
			},
			LeadIn:  "Record one with",
			Command: fmt.Sprintf("luma-backlog outcome assert %s <%s>", req.Ref, strings.Join(assertionNames(), "|")),
		})
	}
	if !corpus.IsAssertion(req.As) {
		return nil, Refuse(Usage, Refusal{
			Problem: "Unknown claim " + req.As,
			Detail:  []string{"expected " + assertionList()},
			LeadIn:  "Record one with",
			Command: fmt.Sprintf("luma-backlog outcome assert %s <%s>", req.Ref, strings.Join(assertionNames(), "|")),
		})
	}

	it, err := corpus.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, resolveError(err)
	}
	if it.Type() != corpus.Outcome {
		return nil, Refuse(Usage, Refusal{
			Problem: fmt.Sprintf("%s is a %s", it.Slug(), it.Type()),
			Detail:  []string{"only an outcome is asserted"},
			LeadIn:  "See the outcomes with",
			Command: "luma-backlog outcome list",
		})
	}

	// Appended, never replaced. Overwriting would be the one place in this
	// design that destroys history, and it would make "failed once, then
	// succeeded" indistinguishable from "succeeded first time" (ADR-0007).
	if err := appendToList(it.Record, "asserted", map[string]string{
		"by": s.Env.Actor.String(),
		"at": s.Env.Now(),
		"as": req.As,
	}); err != nil {
		return nil, FailureError("%w", err)
	}

	out, err := it.Record.Bytes()
	if err != nil {
		return nil, FailureError("%w", err)
	}
	if err := s.Backlog.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return nil, FailureError("%w", err)
	}

	return &AssertResult{
		Path:        it.Path,
		Attempts:    countList(it.Record, "asserted"),
		Contradicts: it.Record.Has("verified"),
	}, nil
}

func assertionList() string { return strings.Join(assertionNames(), " or ") }

// assertionNames is the vocabulary as plain strings, for a caller that renders
// them rather than compares them.
func assertionNames() []string {
	names := make([]string, 0, len(corpus.Assertions))
	for _, a := range corpus.Assertions {
		names = append(names, string(a))
	}
	return names
}
