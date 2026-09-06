package app

// Node is a work item with what hangs off it.
//
// The listing and the tree are one call apart on purpose: a board wants the
// same assembly a terminal does, and an adapter that built it itself would be
// carrying judgment (ADR-0004).
type Node struct {
	View
	// Children are the records belonging to this work item, outcomes before
	// tasks --- what must be true before how it is being attempted.
	Children []View
}

// TreeResult is a tree and what was noticed while producing it.
type TreeResult struct {
	Nodes []Node
	Observations
}

// Tree lists work items and gathers each one's outcomes and tasks.
//
// The filter narrows the work items, never their children: asking for
// in-progress work items and being shown only their in-progress tasks would
// hide the ones nobody has started, which is usually the reason for looking.
func (s *Session) Tree(f Filter) (*TreeResult, error) {
	f.Unit = WorkItem
	listed, err := s.List(f)
	if err != nil {
		return nil, err
	}

	out := &TreeResult{Observations: listed.Observations}
	for _, wi := range listed.Items {
		node := Node{View: wi}
		for _, unit := range []string{Outcome, Task} {
			kids, err := s.List(Filter{Unit: unit, WorkItem: wi.Name})
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, kids.Items...)
			out.Observations.Skipped = append(out.Observations.Skipped, kids.Skipped...)
		}
		out.Nodes = append(out.Nodes, node)
	}
	return out, nil
}
