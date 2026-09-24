package corpus

import "testing"

// Allocation runs out, and when it does it refuses rather than returning a
// position a neighbor already holds.
//
// **The numbers are the point.** Repeated allocation at one spot costs about a
// decimal digit each time, so the scheme has a budget --- and that budget is
// small enough to be reached by ordinary use, not only by abuse. A record that
// everything else queues in front of consumes the front budget one move at a
// time. Measured 2026-09-17 for
// .luma/backlog/work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key.
//
// The bounds are asserted loosely, as floors rather than exact values: a change
// to how positions are allocated should move them, and a test pinned to the
// exact number would fail on an improvement. What must not change is that
// exhaustion is an error.
func TestAllocationExhaustsLoudlyRatherThanColliding(t *testing.T) {
	cases := []struct {
		name  string
		floor int
		next  func(Position) (Position, error)
	}{
		// Appending is cheap while whole steps fit inside the range --- 999 of
		// them --- and bisects against the ceiling afterwards.
		{"append", 999, func(p Position) (Position, error) { return Between(p, "") }},
		// Prepending bisects from the first move: the seed is one step above
		// zero, so there is no room to step down into.
		{"prepend", 100, func(p Position) (Position, error) { return Between("", p) }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cur := FirstPosition
			for n := 1; ; n++ {
				next, err := c.next(cur)
				if err != nil {
					if n <= c.floor {
						t.Errorf("%s ran out after %d allocations, below the %d this scheme is expected to manage: %v",
							c.name, n, c.floor, err)
					}
					return
				}
				// The failure this guards against: a rounded position equal to
				// the neighbor it was meant to separate, which loses the order
				// and reports nothing.
				if c.name == "append" && next <= cur {
					t.Fatalf("append returned %q, not after %q --- two records would share a position", next, cur)
				}
				if c.name == "prepend" && next >= cur {
					t.Fatalf("prepend returned %q, not before %q --- two records would share a position", next, cur)
				}
				cur = next
				if n > 100000 {
					t.Fatalf("%s never exhausted, which means the budget is unbounded and this test measures nothing", c.name)
				}
			}
		})
	}
}
