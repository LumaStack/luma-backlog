package backlog

import (
	"github.com/lumastack/luma-backlog/internal/root"
)

// Completion is the arithmetic behind closing a deliverable.
type Completion struct {
	Live      []Item // outcomes counted toward completion
	Unpassing []Item // of those, the ones with no evidence
	Retired   []Item // archived, and excluded (docs/spec.md §4.4)
}

// Complete reports whether every live outcome has passed.
func (c Completion) Complete() bool { return len(c.Live) > 0 && len(c.Unpassing) == 0 }

// CompletionOf counts the outcomes of a deliverable.
//
// Derived by counting, never read from a field. There is nothing to store that
// the verification records do not already say, and a stored copy could
// disagree with them (docs/spec.md §2.4).
func CompletionOf(b *root.Backlog, deliverable string) (Completion, error) {
	items, err := List(b, Filter{Unit: Outcome, Deliverable: deliverable})
	if err != nil {
		return Completion{}, err
	}

	var c Completion
	for _, it := range items {
		// A retired outcome is archived rather than deleted, and excluded
		// from the arithmetic — otherwise retiring one could never let a
		// deliverable close, which is the point of retiring it.
		if s, _ := it.Record.Get("lifecycle"); s == "archived" {
			c.Retired = append(c.Retired, it)
			continue
		}
		c.Live = append(c.Live, it)
		if !it.Record.Has("verified") {
			c.Unpassing = append(c.Unpassing, it)
		}
	}
	return c, nil
}

// CloseReason is why work ended. Only one of them is success, which is why
// the terminal state is "closed" rather than "done" (docs/spec.md §5.3.1).
type CloseReason string

const (
	Delivered  CloseReason = "delivered"
	Canceled   CloseReason = "canceled"
	Superseded CloseReason = "superseded"
	Abandoned  CloseReason = "abandoned"
)

// CloseReasons lists them in the order they are offered.
var CloseReasons = []CloseReason{Delivered, Canceled, Superseded, Abandoned}

// IsCloseReason reports whether a value is one.
func IsCloseReason(s string) bool {
	for _, r := range CloseReasons {
		if CloseReason(s) == r {
			return true
		}
	}
	return false
}

// GatedOnCompletion reports whether a reason requires every outcome to pass.
//
// Only delivery is gated. Gating cancellation would make it impossible to stop
// work precisely because it was unfinished — which is the only reason anyone
// ever cancels anything.
func (r CloseReason) GatedOnCompletion() bool { return r == Delivered }
