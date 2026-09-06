package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Ladder is one unit's workflow statuses in order, each carrying the ordinal
// that prefixes a rank (ADR-0005). The ordinal is what lets sorting the raw
// rank field give board order without reading configuration.
type Ladder struct {
	// Statuses in ladder order.
	Statuses []string
	// Ordinals maps a status to its sparse position. Sparse on purpose:
	// inserting a status between two others must not renumber anything.
	Ordinals map[string]int
}

// Len reports how many statuses the ladder holds.
func (l Ladder) Len() int { return len(l.Statuses) }

// Ordinal reports a status's ordinal, and whether the ladder knows it. A
// status the ladder does not carry has no place in the order, and guessing one
// would file the record somewhere nobody chose.
func (l Ladder) Ordinal(status string) (int, bool) {
	n, ok := l.Ordinals[status]
	return n, ok
}

// ordinalStep spaces derived ordinals so statuses can be inserted between two
// existing ones without renumbering. Only used for the list form.
const ordinalStep = 10

// ladderAt builds a ladder from alternating status and ordinal, for the
// built-in defaults where the numbers are chosen rather than derived.
func ladderAt(pairs ...any) Ladder {
	l := Ladder{Ordinals: map[string]int{}}
	for i := 0; i+1 < len(pairs); i += 2 {
		status, n := pairs[i].(string), pairs[i+1].(int)
		l.Statuses = append(l.Statuses, status)
		l.Ordinals[status] = n
	}
	return l
}

func ladderOf(statuses ...string) Ladder {
	l := Ladder{Statuses: statuses, Ordinals: make(map[string]int, len(statuses))}
	for i, s := range statuses {
		l.Ordinals[s] = (i + 1) * ordinalStep
	}
	return l
}

// UnmarshalYAML accepts both shapes a ladder may be written in.
//
// ADR-0005 asks for explicit sparse ordinals, so that inserting a status costs
// nothing and moving one touches only the records in it:
//
//	work-item: {captured: 10, todo: 50, closed: 90}
//
// A plain list keeps working and derives its ordinals by position:
//
//	work-item: [captured, todo, closed]
//
// The list form is not deprecated here --- every configuration written before
// this uses it, including this repository's own. What it gives up is the
// property the explicit form exists for: inserting a status shifts every
// ordinal after it, which rewrites the rank of every record at those statuses.
func (l *Ladder) UnmarshalYAML(n *yaml.Node) error {
	switch n.Kind {
	case yaml.SequenceNode:
		var statuses []string
		if err := n.Decode(&statuses); err != nil {
			return err
		}
		*l = ladderOf(statuses...)
		return nil

	case yaml.MappingNode:
		var pairs map[string]int
		if err := n.Decode(&pairs); err != nil {
			return err
		}
		// Mapping order in YAML is the document's order, which is the ladder
		// order. Reading the node directly keeps it; ranging the map loses it.
		out := Ladder{Ordinals: make(map[string]int, len(pairs))}
		for i := 0; i+1 < len(n.Content); i += 2 {
			status := n.Content[i].Value
			out.Statuses = append(out.Statuses, status)
			out.Ordinals[status] = pairs[status]
		}
		if err := out.checkOrdinals(); err != nil {
			return err
		}
		*l = out
		return nil
	}
	return fmt.Errorf("workflow_status must be a list of statuses or a mapping of status to ordinal, at line %d", n.Line)
}

// checkOrdinals refuses a ladder whose ordinals do not ascend with its order.
// Ascending order is the work order (ADR-0005); a ladder that reads one way and
// sorts another would put the board in an order nobody wrote down.
func (l Ladder) checkOrdinals() error {
	seen := map[int]string{}
	last := 0
	for i, s := range l.Statuses {
		n := l.Ordinals[s]
		if prev, dup := seen[n]; dup {
			return fmt.Errorf("workflow_status: %q and %q both have ordinal %d", prev, s, n)
		}
		seen[n] = s
		if i > 0 && n <= last {
			return fmt.Errorf("workflow_status: %q has ordinal %d, which does not come after %d --- ordinals must ascend with the ladder", s, n, last)
		}
		last = n
	}
	return nil
}
