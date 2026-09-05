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

	// However the caller named the work item — directory, slug half, or key —
	// everything below needs the directory, because that is where a child is
	// written and what its work_item link must say.
	if s.WorkItem != "" {
		resolved, resolveErr := ResolveWorkItemDir(b, s.WorkItem)
		if resolveErr != nil {
			return Result{}, resolveErr
		}
		s.WorkItem = resolved
	}

	slug := Slugify(s.Title)
	if slug == "" {
		return Result{}, fmt.Errorf("title %q has no characters that can form a filename", s.Title)
	}

	// A work item's directory carries its key, so its path is not derivable
	// from the title alone either. Look for one whose slug half matches before
	// allocating anything.
	if s.Unit == WorkItem {
		existing, findErr := findWorkItemBySlug(b, slug)
		if findErr != nil {
			return Result{}, findErr
		}
		if existing != "" {
			return Result{Path: path.Join(BundleDir, "work-items", existing, "index.md"), Created: false}, nil
		}
	}

	// Allocated before the path, because the path contains it. Safe to do
	// here only because the existence check above already returned for a work
	// item that exists — a retry never reaches this line, so a number cannot
	// be burned by one.
	key := ""
	if s.Unit == WorkItem {
		highest, keyErr := highestKey(b)
		if keyErr != nil {
			return Result{}, keyErr
		}
		key = FormatKey(highest + 1)
		// The key leads the directory name, so a listing sorts by it and the
		// identifier on disk is the one people say. The slug follows, so the
		// directory still reads as what the work is.
		slug = key + "-" + slug
	}

	// A decision is numbered, so its filename is not derived from the title
	// alone and the existence check cannot be a path lookup.
	adr := 0
	if s.Unit == Decision {
		highest, existing, scanErr := scanDecisions(b, slug)
		if scanErr != nil {
			return Result{}, scanErr
		}
		if existing != "" {
			return Result{Path: existing, Created: false}, nil
		}
		adr = highest + 1
		slug = strings.TrimSuffix(decisionFilename(adr, slug), ".md")
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

	body, err := render(s, cfg, e, adr, key)
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

func render(s Spec, cfg config.Config, e env.Env, adr int, key string) ([]byte, error) {
	r, err := record.Parse([]byte("---\n---\n"))
	if err != nil {
		return nil, err
	}

	r.Set("type", s.Unit)
	if key != "" {
		// Second in the frontmatter, because it is the handle somebody reads
		// first and quotes in a commit or a conversation.
		r.Set("key", key)
	}
	r.Set("title", s.Title)

	switch s.Unit {
	case Decision:
		// Present and empty rather than absent, for the reason an outcome's
		// desired_state is: an empty field asks to be filled in and a missing
		// one is not noticed.
		//
		// `decided` is deliberately not today's date. It records when the
		// position became BINDING, which is not when the file appeared — a
		// decision can be drafted for a week and settled in a meeting, and
		// only one of those is the fact people cite. Stamping creation here
		// would assert something nobody has decided yet.
		r.Set("decided", "")
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
	if s.Unit == Decision {
		// A decision with no re-open condition becomes permanent by inertia —
		// not because anybody reaffirmed it, but because nobody knew what would
		// justify revisiting.
		r.Set("reopen_trigger", "")
	}
	if err := r.SetRaw("created", "{by: "+e.Actor.String()+", at: "+e.Now()+"}"); err != nil {
		return nil, err
	}

	r.SetBody(bodyFor(s.Unit, s.Title, adr))
	return r.Bytes()
}

// bodyFor is the starting shape for a unit's body. Sections are a starting
// point, not a form: leave one out rather than writing nothing under it
// (open-questions.md §17).
func bodyFor(unit, title string, adr int) string {
	h := "# " + title + "\n\n"
	if unit == Decision {
		h = fmt.Sprintf("# ADR-%04d: %s\n\n", adr, title)
	}
	switch unit {
	case WorkItem:
		return h + "## The problem\n\n## What is being delivered\n\n## Out of scope\n\n## Constraints\n"
	case Outcome:
		return h + "Why this matters, and anything needed to read the check correctly.\n"
	case Task:
		return h + "What is to be done, and how it will be verified.\n"
	case Decision:
		// Summary, Problem, Decision and Why are required of every record;
		// everything else is written only when it carries real content, since
		// a heading with nothing under it teaches the next author that the
		// headings matter more than the reasoning.
		return h +
			"## Summary\n\nOne sentence: what was decided.\n\n" +
			"## Problem\n\nWhy is anything being decided? The trigger, the forces, the constraints.\n\n" +
			"## Decision\n\nWhat was decided — \"We will …\".\n\n" +
			"## Why\n\nObservable reasoning, not assertion. Something a reader can check,\n" +
			"argue with, and one day find no longer true.\n\n" +
			"<!-- Optional, and only when they carry real content: Alternatives ·\n" +
			"     Tradeoffs · Assumptions · Revisit When · Follow-up · References.\n" +
			"     Record a path not taken as deferred with what would reopen it,\n" +
			"     never as rejected. -->\n"
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
