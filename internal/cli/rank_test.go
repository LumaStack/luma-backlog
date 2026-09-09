package cli

import (
	"strings"
	"testing"
)

// Seeding leaves room to insert, and each new record lands behind the last.
func TestRankSeedsAndAppends(t *testing.T) {
	app, _ := initialized(t)
	for _, title := range []string{"Alpha", "Bravo", "Charlie"} {
		run(t, app, "work-item", "new", title)
	}
	for _, ref := range []string{"WORK-0001", "WORK-0002", "WORK-0003"} {
		if code, _, e := run(t, app, "work-item", "rank", ref, "--last"); code != ExitOK {
			t.Fatalf("%s: exit %d: %s", ref, code, e)
		}
	}
	_, out, _ := run(t, app, "work-item", "list")
	if order(out) != "Alpha Bravo Charlie" {
		t.Errorf("order = %q", order(out))
	}
}

// Moving one record writes one file and leaves its neighbors untouched, which
// is what a decimal ordering key buys (spec.md §9.6).
func TestRankMovesWithoutDisturbingNeighbors(t *testing.T) {
	app, _ := initialized(t)
	for _, title := range []string{"Alpha", "Bravo", "Charlie"} {
		run(t, app, "work-item", "new", title)
	}
	for _, ref := range []string{"WORK-0001", "WORK-0002", "WORK-0003"} {
		run(t, app, "work-item", "rank", ref, "--last")
	}
	_, before, _ := run(t, app, "show", "WORK-0002", "--json")

	run(t, app, "work-item", "rank", "WORK-0003", "--first")
	_, out, _ := run(t, app, "work-item", "list")
	if order(out) != "Charlie Alpha Bravo" {
		t.Errorf("order = %q", order(out))
	}
	_, after, _ := run(t, app, "show", "WORK-0002", "--json")
	if before != after {
		t.Errorf("ranking one record rewrote another:\nbefore %s\nafter  %s", before, after)
	}
}

// --before and --after place a record against a named neighbor.
func TestRankBeforeAndAfter(t *testing.T) {
	app, _ := initialized(t)
	for _, title := range []string{"Alpha", "Bravo", "Charlie"} {
		run(t, app, "work-item", "new", title)
	}
	for _, ref := range []string{"WORK-0001", "WORK-0002", "WORK-0003"} {
		run(t, app, "work-item", "rank", ref, "--last")
	}
	run(t, app, "work-item", "rank", "WORK-0003", "--before", "WORK-0001")
	_, out, _ := run(t, app, "work-item", "list")
	if order(out) != "Charlie Alpha Bravo" {
		t.Errorf("--before: order = %q", order(out))
	}
	run(t, app, "work-item", "rank", "WORK-0003", "--after", "WORK-0002")
	_, out, _ = run(t, app, "work-item", "list")
	if order(out) != "Alpha Bravo Charlie" {
		t.Errorf("--after: order = %q", order(out))
	}
}

// A record nobody has placed sorts after every record somebody has. Putting it
// first would let an unconsidered record outrank a considered one.
func TestUnrankedRecordsSortLast(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "work-item", "new", "Zulu")
	run(t, app, "work-item", "rank", "WORK-0002", "--first")

	_, out, _ := run(t, app, "work-item", "list")
	if order(out) != "Zulu Alpha" {
		t.Errorf("order = %q", order(out))
	}
}

// §9.6: the caller never computes an ordering key. If set can write one, that
// rule is decoration (ADR-0005).
func TestSetRefusesTheRankField(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, errOut := run(t, app, "set", "WORK-0001", "rank=010.0010.000")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "work-item rank") {
		t.Errorf("the refusal did not name the command to use instead:\n%s", errOut)
	}
}

// Rank orders records within a status. Ranking against a record at another
// status is asking for a position that cannot exist, and is refused rather
// than silently doing something else.
func TestRankAgainstAnotherStatusIsRefused(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "work-item", "new", "Bravo")
	run(t, app, "work-item", "rank", "WORK-0002", "--last")
	run(t, app, "work-item", "transition", "WORK-0002", "todo")

	code, _, errOut := run(t, app, "work-item", "rank", "WORK-0001", "--before", "WORK-0002")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "same one") {
		t.Errorf("the refusal did not explain why:\n%s", errOut)
	}
}

// Saying nothing about where is a usage error, not a default.
func TestRankNeedsToBeToldWhere(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	if code, _, _ := run(t, app, "work-item", "rank", "WORK-0001"); code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
}

// order reduces a listing to its titles, so a test reads as the order it means.
func order(listing string) string {
	var titles []string
	for _, line := range strings.Split(strings.TrimSpace(listing), "\n")[1:] {
		f := strings.Fields(line)
		if len(f) >= 3 {
			titles = append(titles, strings.Join(f[2:], " "))
		}
	}
	return strings.Join(titles, " ")
}
