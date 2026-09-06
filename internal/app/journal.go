package app

import (
	"path"
	"sort"
	"strings"

	"github.com/lumastack/luma-backlog/internal/backlog"
)

// JournalRequest reads or appends to a work item's journal.
type JournalRequest struct {
	// WorkItem names whose journal. Empty means derive it.
	WorkItem string
	// Line is the text to append. Empty means read.
	Line string
}

// JournalResult is the journal, or where a line went.
type JournalResult struct {
	Slug string
	Path string
	// Content is the journal as it stands, on a read.
	Content string
	// Written is true when a line was appended.
	Written bool
}

// Journal appends one line to today's entry, opening today's entry if there is
// not one — or shows the journal when given nothing to write.
//
// No file to open, no heading to write, no decision about where it goes:
// friction at the moment of writing is what loses the learning.
func (s *Session) Journal(req JournalRequest) (*JournalResult, error) {
	slug, err := s.journalWorkItem(req.WorkItem)
	if err != nil {
		return nil, err
	}
	rel := path.Join(backlog.BundleDir, "work-items", slug, "journal.md")

	current := ""
	if data, err := s.Backlog.ReadFile(rel); err == nil {
		current = string(data)
	}

	if req.Line == "" {
		return &JournalResult{Slug: slug, Path: rel, Content: current}, nil
	}

	line := strings.TrimSpace(req.Line)
	if line == "" {
		return nil, UsageError("nothing to write")
	}
	if err := s.Backlog.WriteFileAtomic(rel, []byte(
		backlog.AppendLine(current, s.Env.Today(), line)), 0o644); err != nil {
		return nil, FailureError("%w", err)
	}
	return &JournalResult{Slug: slug, Path: rel, Written: true}, nil
}

// journalWorkItem decides whose journal to write to.
//
// In order: the name given, then the working directory, then — only when there
// is exactly one work item — that one. The last is a real convenience early on
// and errs safe: the moment there are two, it stops guessing and names them.
//
// The precedence matters more here than for other operations. Capture has to
// cost one invocation, and requiring the work item every time is the friction
// that stops it happening at all.
func (s *Session) journalWorkItem(given string) (string, error) {
	if given != "" {
		// Resolved, not taken as written. A caller names a work item by its
		// directory, its slug half, or its key, and the journal has to reach
		// the same record everything else does — otherwise it writes to
		// work-items/<whatever-was-typed>/journal.md and quietly creates a
		// directory that is not a work item at all.
		dir, err := backlog.ResolveWorkItemDir(s.Backlog, given)
		if err != nil {
			return "", FailureError("%w", err)
		}
		if !s.Backlog.Exists(path.Join(backlog.BundleDir, "work-items", dir, "index.md")) {
			return "", UsageError("no work item %q", given)
		}
		return dir, nil
	}
	if fromDir := s.workItemFromWorkingDir(); fromDir != "" {
		return fromDir, nil
	}

	// Deferred with Resolve, and for the same reason — see load.go.
	items, _, err := backlog.List(s.Backlog, backlog.Filter{Unit: backlog.WorkItem})
	if err != nil {
		return "", FailureError("%w", err)
	}
	switch len(items) {
	case 0:
		return "", UsageError("no work items yet — create one first")
	case 1:
		return items[0].Slug(), nil
	}

	slugs := make([]string, 0, len(items))
	for _, it := range items {
		slugs = append(slugs, it.Slug())
	}
	sort.Strings(slugs)
	return "", UsageError("more than one work item — say which with --work-item:\n  %s",
		strings.Join(slugs, "\n  "))
}
