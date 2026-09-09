package corpus

import (
	"sort"
	"strings"
	"testing"
)

// The table in docs/spec.md §9.6, held as a test so the document and the code
// cannot drift apart.
func TestBetweenMatchesTheSpecifiedTable(t *testing.T) {
	for _, tc := range []struct{ name, before, after, want string }{
		{"seeding a new backlog", "", "", "0010.000"},
		{"between 0010 and 0020", "0010.000", "0020.000", "0015.000"},
		{"between 0010.000 and 0011.000", "0010.000", "0011.000", "0010.500"},
		{"squeezed", "0010.000", "0010.002", "0010.001"},
		{"appending", "0010.000", "", "0020.000"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Between(Position(tc.before), Position(tc.after))
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tc.want {
				t.Errorf("Between(%q, %q) = %q, want %q", tc.before, tc.after, got, tc.want)
			}
		})
	}
}

// Zero-padding is what makes lexicographic order and numeric order the same
// order, so a consumer that only compares text still sorts correctly. This is
// the property the whole scheme rests on.
func TestTextOrderIsNumericOrder(t *testing.T) {
	var got []string
	// Insert repeatedly at the front, which produces the values most likely to
	// lose their padding.
	positions := []Position{"0010.000", "0020.000", "0030.000"}
	for i := 0; i < 12; i++ {
		p, err := Between("", positions[0])
		if err != nil {
			t.Fatal(err)
		}
		positions = append([]Position{p}, positions...)
	}
	for _, p := range positions {
		got = append(got, string(p))
	}

	sorted := append([]string(nil), got...)
	sort.Strings(sorted)
	if strings.Join(got, " ") != strings.Join(sorted, " ") {
		t.Errorf("text order differs from insertion order:\n got %v\nsort %v", got, sorted)
	}
}

// Precision extends rather than the scheme failing (§9.6). Roughly ten
// bisections at the same position exhaust three decimals; a rebalance must
// never be mandatory, because it would arrive as a multi-record write in the
// middle of a drag.
func TestPrecisionExtendsWhenSqueezed(t *testing.T) {
	lo, hi := Position("0010.000"), Position("0010.001")
	for i := 0; i < 30; i++ {
		p, err := Between(lo, hi)
		if err != nil {
			t.Fatalf("bisection failed after %d squeezes: %v", i, err)
		}
		if p <= lo || p >= hi {
			t.Fatalf("bisection %d produced %q, not between %q and %q", i, p, lo, hi)
		}
		hi = p
	}
	if len(hi) <= len("0010.000") {
		t.Errorf("precision never extended: %q", hi)
	}
}

// Appending past the ceiling bisects toward it rather than failing (§9.6).
func TestAppendingPastTheCeilingBisects(t *testing.T) {
	p, err := Between("9990.000", "")
	if err != nil {
		t.Fatal(err)
	}
	if p <= "9990.000" {
		t.Errorf("append produced %q, which does not follow 9990.000", p)
	}
}

// --first allocates below the current minimum, because ascending order is the
// work order and a higher number is a lower rank (ADR-0005).
func TestTopAllocatesBelowTheMinimum(t *testing.T) {
	p, err := Between("", "0010.000")
	if err != nil {
		t.Fatal(err)
	}
	if p >= "0010.000" {
		t.Errorf("top produced %q, which is not above 0010.000 in work order", p)
	}
}

// A malformed rank is a hand edit or a migration that went wrong. Reporting it
// beats repairing it, which would hide both.
func TestARankThatDoesNotParseIsReported(t *testing.T) {
	for _, bad := range []Rank{"", "50", "50.", "notanordinal.0010.000", "50.notaposition"} {
		if _, _, err := SplitRank(bad); err == nil {
			t.Errorf("SplitRank(%q) accepted a malformed rank", bad)
		}
	}
}

// The ordinal is padded as well as the position. Without it 100 sorts before
// 20 as text, and the prefix that exists to make sorting work breaks it.
func TestTheOrdinalIsPaddedToo(t *testing.T) {
	low := MakeRank(20, "0010.000")
	high := MakeRank(100, "0010.000")
	if !(low < high) {
		t.Errorf("ordinal 20 (%q) did not sort before ordinal 100 (%q)", low, high)
	}
	ordinal, pos, err := SplitRank(high)
	if err != nil {
		t.Fatal(err)
	}
	if ordinal != 100 || pos != "0010.000" {
		t.Errorf("round trip gave %d, %q", ordinal, pos)
	}
}
