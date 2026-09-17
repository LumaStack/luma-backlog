package corpus

import (
	"math/big"
	"testing"
	"time"
)

// A position whose value cannot terminate in base ten is rounded, not spun on.
//
// formatPosition extends precision until a value is exactly representable, and
// nothing in the type system says a position has to be a finite decimal. One
// that is not --- 9990/20001, which is what dividing a range between twenty
// thousand records gives --- made the search run forever: no error, no wrong
// answer, just a process that never returned. Found 2026-09-17.
//
// The bound is a tripwire for a caller doing arithmetic this scheme does not
// do, so the assertion is only that it returns.
func TestFormatPositionTerminatesOnARepeatingValue(t *testing.T) {
	done := make(chan Position, 1)
	go func() {
		done <- formatPosition(new(big.Rat).SetFrac64(9990, 20001))
	}()
	select {
	case got := <-done:
		if got == "" {
			t.Errorf("no position came back")
		}
	case <-time.After(10 * time.Second):
		t.Fatal("formatPosition did not return on a repeating value")
	}
}
