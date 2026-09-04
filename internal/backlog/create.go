package backlog

import (
	"fmt"
	"path"
	"strings"

	"github.com/lumastack/luma-backlog/internal/config"
	"github.com/lumastack/luma-backlog/internal/env"
	"github.com/lumastack/luma-backlog/internal/record"
	"github.com/lumastack/luma-backlog/internal/root"
)

// Spec is what a caller asked for. Everything absent is derived or omitted —
// no field here is something the author should have to know about the format
// (docs/spec.md §9.0).
type Spec struct {
	Unit     string
	Title    string
	WorkItem string // slug; empty means derive from context or omit
	Kind     string // classifies a work item; empty means ordinary work
}

// Result reports what happened, so a caller can tell a creation from a
// no-op without parsing prose.
type Result struct {
	Path    string
	Created bool // false means it already existed and was left alone
}

// Create writes a new record.
//
// Idempotent by name (docs/spec.md §9.5): asking twice for the same record
// leaves the first one alone and reports it. An agent that retries after a
// dropped connection must not destroy the work of its first attempt.
func Create(b *root.Backlog, cfg config.Config, e env.Env, s Spec) (Result, error) {
	if !IsUnit(s.Unit) {
		return Result{}, fmt.Errorf("unknown unit %q: expected one of %s", s.Unit, strings.Join(Units, ", "))
	}
	if strings.TrimSpace(s.Title) == "" {
		return Result{}, fmt.Errorf("a title is required")
	}

	slug := Slugify(s.Title)
	if slug == "" {
		return Result{}, fmt.Errorf("title %q has no characters that can form a filename", s.Title)
	}

	rel, err := PathFor(s.Unit, slug, s.WorkItem)
	if err != nil {
		return Result{}, err
	}
	if s.Unit == Decision && s.WorkItem == "" {
		rel += ".md"
	}

	if b.Exists(rel) {
		return Result{Path: rel, Created: false}, nil
	}

	body, err := render(s, cfg, e)
	if err != nil {
		return Result{}, err
	}
	if err := b.WriteFileAtomic(rel, body, 0o644); err != nil {
		return Result{}, err
	}

	// A work item gets its journal at once. Somewhere to write has to exist
	// before anyone needs it, or the writing does not happen (§5.5).
	if s.Unit == WorkItem {
		j := path.Join(path.Dir(rel), "journal.md")
		if !b.Exists(j) {
			if err := b.WriteFileAtomic(j, []byte(journalFor(s.Title)), 0o644); err != nil {
				return Result{}, err
			}
		}
	}
	return Result{Path: rel, Created: true}, nil
}

func render(s Spec, cfg config.Config, e env.Env) ([]byte, error) {
	r, err := record.Parse([]byte("---\n---\n"))
	if err != nil {
		return nil, err
	}

	r.Set("type", s.Unit)
	r.Set("title", s.Title)

	switch s.Unit {
	case Outcome:
		// desired_state is the whole point of an outcome, so it is present
		// and empty rather than absent: an empty field asks to be filled in,
		// a missing one is not noticed.
		r.Set("desired_state", "")
		r.Set("verify_by", "")
	case Task, WorkItem:
	}
	if s.WorkItem != "" && s.Unit != WorkItem {
		r.Set("work_item", "[[work-items/"+s.WorkItem+"]]")
	}
	// Only units that are WORKED carry a workflow status. An outcome is
	// judged by its evidence, a decision is ratified through stage,
	// an exploration is archived. Giving an outcome a declared status would
	// put it beside the computed one, free to contradict it — which is the
	// unbacked assertion this design exists to distrust (docs/spec.md §4.4).
	if IsWorked(s.Unit) {
		r.Set("workflow_status", cfg.DefaultStatusFor(s.Unit))
	}
	// A kind is written only when somebody said one. Absence means ordinary
	// work — something we decided to do, obliging nobody — which is the
	// common case and is why the field is not defaulted (docs/spec.md §4.2).
	if s.Kind != "" {
		r.Set("kind", CanonicalKind(s.Kind))
	}
	r.Set("stage", "draft")
	if err := r.SetRaw("created", "{by: "+e.Actor.String()+", at: "+e.Now()+"}"); err != nil {
		return nil, err
	}

	r.SetBody(bodyFor(s.Unit, s.Title))
	return r.Bytes()
}

// bodyFor is the starting shape for a unit's body. Sections are a starting
// point, not a form: leave one out rather than writing nothing under it
// (open-questions.md §17).
func bodyFor(unit, title string) string {
	h := "# " + title + "\n\n"
	switch unit {
	case WorkItem:
		return h + "## The problem\n\n## What is being delivered\n\n## Out of scope\n\n## Constraints\n"
	case Outcome:
		return h + "Why this matters, and anything needed to read the check correctly.\n"
	case Task:
		return h + "What is to be done, and how it will be verified.\n"
	case Decision:
		return h + "## Context\n\n## What was chosen\n\n## What was not taken, and why\n\nRecord what would reopen an option, rather than calling it rejected.\n"
	case Exploration:
		return h + "## The question\n\n## What was found\n\n## What it means\n"
	}
	return h
}

func journalFor(title string) string {
	return "# Journal — " + title + "\n\n" +
		"> The work item's memory. Newest entry first; everything below the top block is\n" +
		"> historical. Append, never curate. Shape: `spec.md` §5.5.\n\n---\n"
}
