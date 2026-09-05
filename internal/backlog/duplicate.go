package backlog

import "sort"

// Duplicate is one key held by more than one record.
//
// A key is meant to name one record — it is what somebody quotes in a commit,
// says out loud, and cites from another record, and every one of those is wrong
// the moment two records answer to it.
//
// This cannot be prevented locally. Keys are allocated optimistically because
// the alternative is a coordinator, which docs/spec.md §6.1 forbids: independent
// work must never serialize. Two workstations read the same highest key and both
// allocate the next one, and §6.4 is explicit that across branches or machines
// this is not solvable by local means.
//
// **And a merge will not raise it.** Two work items created on two branches
// touch entirely different files, so git merges them cleanly and the duplicate
// arrives with no conflict and no warning. Nothing else in the system is
// looking, which is why this looks.
type Duplicate struct {
	Key   string
	Paths []string
}

// Duplicates reports every key held by more than one record.
//
// Detection, not repair. What a duplicate should become is
// backlog/work-items/WORK-0013-how-two-workstations-avoid-colliding, and
// records/decisions/ADR-0003 leans toward sending one of them to the end of the
// sequence. Reporting is useful before that is settled, and it is what makes any
// repair usable at all: a repair nobody knows is needed does not happen.
func Duplicates(items []Item) []Duplicate {
	byKey := map[string][]string{}
	for _, it := range items {
		if k := it.Key(); k != "" {
			byKey[k] = append(byKey[k], it.Path)
		}
	}

	var dups []Duplicate
	for key, paths := range byKey {
		if len(paths) < 2 {
			continue
		}
		sort.Strings(paths)
		dups = append(dups, Duplicate{Key: key, Paths: paths})
	}
	// Sorted so the report is byte-stable across runs and across machines,
	// which is what makes it safe to pin in a test.
	sort.Slice(dups, func(a, b int) bool { return dups[a].Key < dups[b].Key })
	return dups
}
