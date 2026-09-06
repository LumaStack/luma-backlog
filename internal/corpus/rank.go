package corpus

import (
	"fmt"
	"math/big"
	"strings"
)

// A rank is <status ordinal>.<position> --- ADR-0005, docs/spec.md §9.6.
//
// Both halves are zero-padded to a fixed width, which is the whole point:
// lexicographic order and numeric order are then the same order, so anything
// that can compare text sorts a backlog correctly. There is no comparator to
// implement and therefore no way to get ordering silently wrong.
//
// The ordinal is padded as well as the position. Without it `100` sorts before
// `20` as text, and the prefix that exists to make sorting work would be the
// thing breaking it.
const (
	ordinalDigits  = 3 // up to 999 statuses in a ladder
	positionDigits = 4 // docs/spec.md §9.6: four digits, a point, three decimals
	positionScale  = 3

	// seedStep spaces the first ranks so there is room to insert between them
	// without immediately needing decimals.
	seedStep = 10
	// positionCeiling is the largest whole position. Appending past it bisects
	// toward the ceiling rather than failing (§9.6).
	positionCeiling = 9990
)

// Rank is a whole rank field as stored.
type Rank string

// Position is the ordering half of a rank, without the status ordinal.
type Position string

// FirstPosition is what a queue's first record gets.
const FirstPosition = Position("0010.000")

// MakeRank joins an ordinal and a position.
func MakeRank(ordinal int, p Position) Rank {
	return Rank(fmt.Sprintf("%0*d.%s", ordinalDigits, ordinal, p))
}

// SplitRank takes a rank apart. A rank that does not parse reports so rather
// than being repaired: a malformed rank is a record somebody edited by hand or
// a migration that went wrong, and guessing at it would hide both.
func SplitRank(r Rank) (ordinal int, p Position, err error) {
	s := string(r)
	i := strings.Index(s, ".")
	if i < 0 || i == len(s)-1 {
		return 0, "", fmt.Errorf("rank %q is not <ordinal>.<position>", r)
	}
	if _, err := fmt.Sscanf(s[:i], "%d", &ordinal); err != nil {
		return 0, "", fmt.Errorf("rank %q has no ordinal", r)
	}
	rest := Position(s[i+1:])
	if _, err := parsePosition(rest); err != nil {
		return 0, "", err
	}
	return ordinal, rest, nil
}

func parsePosition(p Position) (*big.Rat, error) {
	r, ok := new(big.Rat).SetString(string(p))
	if !ok {
		return nil, fmt.Errorf("position %q is not a decimal", p)
	}
	return r, nil
}

// formatPosition writes a position at its natural precision --- at least three
// decimals, and more only when the value genuinely needs them.
//
// The width is a normal form, not a hard limit (§9.6). Bisection halves, and
// halving a finite decimal stays finite, so a value is always representable
// exactly; extending is what makes a rebalance never mandatory rather than a
// multi-record write arriving mid-drag.
func formatPosition(r *big.Rat) Position {
	scale := positionScale
	for !isExactAt(r, scale) {
		scale++
	}
	s := r.FloatString(scale)
	// FloatString gives no leading zeros; the padding is what keeps text order
	// and numeric order agreeing.
	if i := strings.Index(s, "."); i < positionDigits {
		s = strings.Repeat("0", positionDigits-i) + s
	}
	return Position(s)
}

func isExactAt(r *big.Rat, scale int) bool {
	shifted := new(big.Rat).Mul(r, new(big.Rat).SetInt(pow10(scale)))
	return shifted.IsInt()
}

func pow10(n int) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
}

// Between allocates a position between two neighbors. Either may be empty,
// meaning there is nothing on that side.
//
// Allocation is by bisection so that moving one record writes one file and
// leaves its neighbors untouched (§9.6). Integer positions would rewrite every
// record after the moved one --- churn on the most visible operation the board
// has, and contention whenever two actors reorder at once.
func Between(before, after Position) (Position, error) {
	switch {
	case before == "" && after == "":
		return FirstPosition, nil

	case before == "":
		// Nothing above: step below the first, or bisect toward zero when
		// there is no room. --top allocates a key *below* the current minimum,
		// because ascending order is the work order (ADR-0005).
		hi, err := parsePosition(after)
		if err != nil {
			return "", err
		}
		step := new(big.Rat).SetInt64(seedStep)
		if lo := new(big.Rat).Sub(hi, step); lo.Sign() > 0 {
			return formatPosition(lo), nil
		}
		return formatPosition(mid(new(big.Rat), hi)), nil

	case after == "":
		// Nothing below: step past the last, or bisect toward the ceiling.
		lo, err := parsePosition(before)
		if err != nil {
			return "", err
		}
		step := new(big.Rat).SetInt64(seedStep)
		next := new(big.Rat).Add(lo, step)
		if next.Cmp(new(big.Rat).SetInt64(positionCeiling)) <= 0 {
			return formatPosition(next), nil
		}
		ceiling := new(big.Rat).SetFrac64(99999999, 10000)
		return formatPosition(mid(lo, ceiling)), nil

	default:
		lo, err := parsePosition(before)
		if err != nil {
			return "", err
		}
		hi, err := parsePosition(after)
		if err != nil {
			return "", err
		}
		if lo.Cmp(hi) >= 0 {
			return "", fmt.Errorf("cannot rank between %q and %q: they are not in order", before, after)
		}
		return formatPosition(mid(lo, hi)), nil
	}
}

func mid(a, b *big.Rat) *big.Rat {
	sum := new(big.Rat).Add(a, b)
	return sum.Quo(sum, new(big.Rat).SetInt64(2))
}
