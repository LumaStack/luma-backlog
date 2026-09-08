package app

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/corpus"
	"gopkg.in/yaml.v3"
)

// VerifyRequest records that an outcome has been confirmed.
type VerifyRequest struct {
	Ref string
	// As is what the checker found — one of corpus.Verdicts.
	As       string
	Evidence string
}

// VerifyResult describes the confirmation.
type VerifyResult struct {
	Path string
	// NoEvidence is true when the confirmation rests on nothing recorded.
	// Reported, never refused.
	NoEvidence bool
}

// Verify adds a confirmation event, and the evidence it rests on.
//
// Verification accumulates: several actors confirming the same outcome is the
// normal case, and a human entry raises the derived trust tier with no special
// handling (docs/spec.md §4.7).
func (s *Session) Verify(req VerifyRequest) (*VerifyResult, error) {
	if req.As == "" {
		return nil, UsageError("a verdict is required: %s\n\n"+
			"Recording proof has to be said out loud. An outcome nobody checked and\n"+
			"one somebody checked and disproved are different facts.", verdictList())
	}
	// Refused BY NAME rather than by omission. `abandoned` is a state an
	// outcome can be in, so validating against the state list would let it
	// through here — and a checker who can abandon an outcome can abandon
	// their own, through the command that exists to be independent of them
	// (ADR-0007). Deciding is not finding out.
	if req.As == "abandoned" {
		return nil, UsageError("abandoned is a decision, not a finding — use `outcome abandon`.\n"+
			"A verdict is one of: %s", verdictList())
	}
	if !corpus.IsVerdict(req.As) {
		return nil, UsageError("unknown verdict %q: expected %s", req.As, verdictList())
	}

	it, err := corpus.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, resolveError(err)
	}
	if it.Type() != corpus.Outcome {
		return nil, UsageError("%s is a %s — only an outcome is verified", it.Slug(), it.Type())
	}

	at, by := s.Env.Now(), s.Env.Actor.String()

	if err := appendToList(it.Record, "verified",
		map[string]string{"by": by, "at": at, "as": req.As}); err != nil {
		return nil, FailureError("%w", err)
	}

	// Evidence sits BESIDE verified rather than inside it, because verified is
	// a core format field and inheritance is add-only, so the key cannot be
	// added there (format-requests.md §3). The two correlate on by and at
	// rather than by position, since parallel lists matched by index break the
	// first time one is hand-edited.
	if strings.TrimSpace(req.Evidence) != "" {
		if err := appendToList(it.Record, "evidence",
			map[string]string{"by": by, "at": at, "what": req.Evidence}); err != nil {
			return nil, FailureError("%w", err)
		}
	}

	out, err := it.Record.Bytes()
	if err != nil {
		return nil, FailureError("%w", err)
	}
	if err := s.Backlog.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return nil, FailureError("%w", err)
	}

	return &VerifyResult{
		Path:       it.Path,
		NoEvidence: strings.TrimSpace(req.Evidence) == "",
	}, nil
}

// appendToList adds an entry to a list-valued field, creating the list when
// absent and tolerating a single bare entry written by hand.
func appendToList(r interface {
	Node(string) *yaml.Node
	SetRaw(string, string) error
}, key string, entry map[string]string) error {
	var existing []map[string]any
	if node := r.Node(key); node != nil {
		if err := node.Decode(&existing); err != nil {
			// A single bare mapping is legal and means a one-element list,
			// following the format's own handling of verified.
			var one map[string]any
			if err2 := node.Decode(&one); err2 != nil {
				return fmt.Errorf("%s is neither a list nor an entry: %w", key, err)
			}
			existing = []map[string]any{one}
		}
	}
	next := map[string]any{}
	for k, v := range entry {
		next[k] = v
	}
	existing = append(existing, next)

	encoded, err := yaml.Marshal(existing)
	if err != nil {
		return err
	}
	return r.SetRaw(key, string(encoded))
}

// countList reports how many entries a list field holds. A bare mapping counts
// as one, the same way appendToList reads it.
func countList(r interface{ Node(string) *yaml.Node }, key string) int {
	node := r.Node(key)
	if node == nil {
		return 0
	}
	var many []map[string]any
	if err := node.Decode(&many); err == nil {
		return len(many)
	}
	var one map[string]any
	if err := node.Decode(&one); err == nil {
		return 1
	}
	return 0
}

func verdictList() string {
	names := make([]string, 0, len(corpus.Verdicts))
	for _, v := range corpus.Verdicts {
		names = append(names, string(v))
	}
	return strings.Join(names, ", ")
}
