package app

import (
	"github.com/lumastack/luma-backlog/internal/backlog"
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
type SetResult struct{ Path string }

// Set changes only the fields named. Everything else — including keys this
// tool knows nothing about — is left exactly as it was.
func (s *Session) Set(req SetRequest) (*SetResult, error) {
	if len(req.Assignments) == 0 && len(req.Unset) == 0 {
		return nil, UsageError("nothing to change: pass field=value, or --unset field")
	}

	it, err := backlog.Resolve(s.Backlog, req.Ref)
	if err != nil {
		return nil, &Error{Kind: NotFound, Err: err}
	}

	// Optimistic concurrency: the caller states what it saw, and a write that
	// would clobber a change it never saw is refused rather than applied
	// (docs/spec.md §6.3). Conflict means re-read and retry, which is
	// different advice from "something broke" — and it is the distinction a
	// retrying agent depends on.
	if req.IfUnchanged != "" && it.Hash() != req.IfUnchanged {
		return nil, ConflictError(
			"%s changed since you read it — re-read and retry\n  you saw:  %s\n  it is now: %s",
			it.Path, shortHash(req.IfUnchanged), shortHash(it.Hash()))
	}

	assignedModified := false
	for _, a := range req.Assignments {
		if a.Field == "" {
			return nil, UsageError("a field name is required")
		}
		if a.Field == "modified" {
			assignedModified = true
		}
		if a.Raw {
			if err := it.Record.SetRaw(a.Field, a.Value); err != nil {
				return nil, UsageError("%w", err)
			}
			continue
		}
		it.Record.Set(a.Field, a.Value)
	}
	for _, key := range req.Unset {
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
	return &SetResult{Path: it.Path}, nil
}

func shortHash(h string) string {
	if len(h) > 12 {
		return h[:12]
	}
	return h
}
