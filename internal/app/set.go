package app

import (
	"fmt"
	"github.com/lumastack/luma-backlog/internal/corpus"
)

// Assignment is one field being written.
//
// Raw distinguishes a string from a value to be parsed. The two are separate
// because a wikilink looks like a YAML list: guessing would turn
// [[work-items/x]] into a nested sequence. How a caller expresses that is its
// own business — this layer takes the distinction already made.
type Assignment struct {
	Field string
	Value string
	Raw   bool
}

// SetRequest changes fields on a record.
type SetRequest struct {
	Ref         string
	Assignments []Assignment
	Unset       []string
	// IfUnchanged is the hash the caller last saw. When set, a write that
	// would clobber a change it never saw is refused rather than applied.
	IfUnchanged string
}

// SetResult describes what was written.
type SetResult struct{ Subject }

// Set changes only the fields named. Everything else — including keys this
// tool knows nothing about — is left exactly as it was.
func (s *Session) Set(req SetRequest) (*SetResult, error) {
	if len(req.Assignments) == 0 && len(req.Unset) == 0 {
		return nil, Refuse(Usage, Refusal{
			Problem: "Nothing to change",
			LeadIn:  "Set a field with",
			Command: "luma-backlog set " + req.Ref + " <field>=<value>",
		})
	}

	it, err := corpus.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, resolveError(err)
	}

	// Optimistic concurrency: the caller states what it saw, and a write that
	// would clobber a change it never saw is refused rather than applied
	// (docs/spec.md §6.3). Conflict means re-read and retry, which is
	// different advice from "something broke" — and it is the distinction a
	// retrying agent depends on.
	if req.IfUnchanged != "" && it.Hash() != req.IfUnchanged {
		return nil, Refuse(Conflict, Refusal{
			Problem: it.Path + " changed since you read it",
			Detail: []string{
				"you saw:   " + shortHash(req.IfUnchanged),
				"it is now: " + shortHash(it.Hash()),
			},
			LeadIn: "Re-read it with",
			// Re-reading is the whole remedy, and the hash comes back with it.
			Command: "luma-backlog show " + req.Ref + " --json",
		})
	}

	assignedModified := false
	for _, a := range req.Assignments {
		if a.Field == "" {
			return nil, Refuse(Usage, Refusal{
				Problem: "A field name is required",
				LeadIn:  "Set a field with",
				Command: "luma-backlog set " + req.Ref + " <field>=<value>",
			})
		}
		if a.Field == "modified" {
			assignedModified = true
		}
		// §9.6 says the caller never computes an ordering key. If `set
		// rank=0010.500` works, that rule is decoration --- so rank is the only
		// way in through the tool (ADR-0005). Editing the file by hand stays
		// available, as it does for everything.
		if a.Field == "rank" {
			return nil, Refuse(Usage, Refusal{
				Problem: "rank is not set directly",
				Detail:  []string{"it is written with the status, so the two cannot disagree"},
				LeadIn:  "Move it with",
				Command: "luma-backlog rank " + req.Ref + " <--first|--last|--before <ref>|--after <ref>>",
			})
		}
		// A status change is one operation that writes rank too (ADR-0005), and
		// the destination rung may require things a field write cannot check.
		// So it is refused here for the same reason `rank` is: letting the
		// assignment through would make the operation optional.
		if a.Field == "workflow_status" {
			return nil, Refuse(Usage, Refusal{
				Problem: "workflow_status is not set directly",
				Detail:  []string{"a move writes the rank too, and checks what the rung requires"},
				LeadIn:  "Move it with",
				Command: fmt.Sprintf("luma-backlog transition %s %s", req.Ref, a.Value),
			})
		}
		if a.Raw {
			if err := it.Record.SetRaw(a.Field, a.Value); err != nil {
				return nil, Refuse(Usage, Refusal{
					Problem: capitalized(err.Error()),
					Detail:  []string{a.Field + "=" + a.Value},
					Note:    "A raw value is written into the frontmatter as given, so it has to be valid there.",
				})
			}
			continue
		}
		it.Record.Set(a.Field, a.Value)
	}
	for _, key := range req.Unset {
		// Removing workflow_status is a status change --- absence reads as the
		// first configured value (spec.md §4.2), so it silently moves the
		// record to the bottom of the ladder without writing rank.
		if key == "workflow_status" {
			return nil, Refuse(Usage, Refusal{
				Problem: "workflow_status cannot be unset",
				Detail:  []string{"absence reads as " + s.Config.DefaultStatusFor(it.Type())},
				LeadIn:  "Move it with",
				Command: "luma-backlog transition " + req.Ref + " <status>",
			})
		}
		if key == "rank" {
			return nil, Refuse(Usage, Refusal{
				Problem: "rank cannot be unset",
				Detail:  []string{"it is written with the status"},
				LeadIn:  "Move it with",
				Command: "luma-backlog transition " + req.Ref + " <status>",
			})
		}
		it.Record.Remove(key)
	}

	// modified advances on edit, as the format defines it. Written after the
	// caller's own changes so an explicit modified: wins — the tool should not
	// overrule something it was just told.
	if !assignedModified {
		if err := it.Record.SetRaw("modified",
			"{by: "+s.Env.Actor.String()+", at: "+s.Env.Now()+"}"); err != nil {
			return nil, FailureError("%w", err)
		}
	}

	out, err := it.Record.Bytes()
	if err != nil {
		return nil, FailureError("serializing %s: %w", it.Path, err)
	}
	if err := s.Backlog.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return nil, FailureError("%w", err)
	}
	return &SetResult{subjectOf(it)}, nil
}

func shortHash(h string) string {
	if len(h) > 12 {
		return h[:12]
	}
	return h
}
