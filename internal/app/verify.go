package app

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"gopkg.in/yaml.v3"
)

// VerifyRequest records that an outcome has been confirmed.
type VerifyRequest struct {
	Ref      string
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
	it, err := backlog.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, &Error{Kind: NotFound, Err: err}
	}
	if it.Type() != backlog.Outcome {
		return nil, UsageError("%s is a %s — only an outcome is verified", it.Slug(), it.Type())
	}

	at, by := s.Env.Now(), s.Env.Actor.String()

	if err := appendToList(it.Record, "verified",
		map[string]string{"by": by, "at": at}); err != nil {
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
