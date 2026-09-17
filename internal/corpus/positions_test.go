package corpus

import "testing"

func TestPositionsForSpacesAndStaysOrdered(t *testing.T) {
	for _, n := range []int{1, 3, 100, 999, 1500, 20000} {
		got := PositionsFor(n)
		if len(got) != n {
			t.Fatalf("PositionsFor(%d) gave %d", n, len(got))
		}
		for i := 1; i < len(got); i++ {
			if !(got[i-1] < got[i]) {
				t.Fatalf("n=%d not ascending as text at %d: %q then %q", n, i, got[i-1], got[i])
			}
		}
	}
	if got := PositionsFor(3); string(got[0]) != "0010.000" || string(got[2]) != "0030.000" {
		t.Errorf("clean spacing lost: %v", got)
	}
}
