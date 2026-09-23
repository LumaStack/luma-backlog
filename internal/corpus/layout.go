package corpus

import (
	"fmt"
	"path"
	"strings"
)

// Unit names, as written in a record's type field before the namespace is
// applied.
const (
	WorkItem    = "work-item"
	Outcome     = "outcome"
	Task        = "task"
	Decision    = "decision"
	Exploration = "exploration"
)

// Units lists what can be created, in the order a person would meet them.
var Units = []string{WorkItem, Outcome, Task, Decision, Exploration}

// Where the tiers sit inside .luma/, per the luma directory layout policy.
// BundleDir holds the knowledge-format bundle — the backlog tier churns, and
// the bundle's config and type definitions travel with it. Decisions made
// outside any work item are records of what happened and live in the
// records tier under their kind.
const (
	BundleDir           = "backlog"
	RecordsDecisionsDir = "records/decisions"

	// WorkItemsDir is where work items live. Named once because PathFor writes
	// into it and the corpus walk accepts out of it — the two spelling it
	// separately is how a record stops being readable by the tool that wrote
	// it.
	WorkItemsDir = BundleDir + "/work-items"

	// ProjectFile is the project record the layout policy reserves. Not a unit
	// this tool creates: `show` reaches it and nothing else does, which is why
	// the walk has to name it rather than derive it.
	ProjectFile = "PROJECT.md"
)

// childDirs maps a unit to the directory it occupies inside a work item.
var childDirs = map[string]string{
	Outcome:     "outcomes",
	Task:        "tasks",
	Decision:    "decisions",
	Exploration: "explorations",
}

// IsWorked reports whether a unit carries a workflow status.
//
// A work item and a task are worked, and move through a sequence. An outcome
// is judged by evidence, a decision is ratified through stage, and
// an exploration is archived — none of them has a position in a workflow, and
// giving them one would create a declared state that can disagree with the
// real one.
func IsWorked(unit string) bool {
	return unit == WorkItem || unit == Task
}

// IsUnit reports whether a name is one this tool creates.
func IsUnit(unit string) bool {
	if unit == WorkItem {
		return true
	}
	_, ok := childDirs[unit]
	return ok
}

// PathFor returns the file a new record occupies.
//
// A work item's a directory whose index.md is the record; everything else
// is one file inside it. WorkItem membership is the only relationship
// encoded in the path, because it is the only one stable enough to be
// (docs/spec.md §7.3).
func PathFor(unit, slug, workItemSlug string) (string, error) {
	if slug == "" {
		return "", fmt.Errorf("a title is required: it becomes the filename")
	}
	if unit == WorkItem {
		return path.Join(WorkItemsDir, slug, "index.md"), nil
	}
	dir, ok := childDirs[unit]
	if !ok {
		return "", fmt.Errorf("unknown unit %q: expected one of %s", unit, strings.Join(Units, ", "))
	}
	if workItemSlug == "" {
		// A decision made outside any work item's legal and lives in the
		// records tier, per the luma directory layout policy — a bundle names
		// its kind, the policy owns the root. The others belong to work and
		// cannot float.
		if unit == Decision {
			return path.Join(RecordsDecisionsDir, slug), nil
		}
		return "", fmt.Errorf("a %s belongs to a work item: pass --work-item, or run inside one", unit)
	}
	return path.Join(WorkItemsDir, workItemSlug, dir, slug+".md"), nil
}

// WorkItemFromPath reads the work item a path sits inside, if any. This
// is how "from context" works when a command is run within a work item.
func WorkItemFromPath(rel string) string {
	parts := strings.Split(path.Clean(strings.TrimPrefix(rel, "/")), "/")
	for i, p := range parts {
		if p == "work-items" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// isChildDir reports whether a directory inside a work item holds records.
// Anything else there — journal.md, evidence/ — is not one.
func isChildDir(dir string) bool {
	for _, d := range childDirs {
		if d == dir {
			return true
		}
	}
	return false
}
