package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
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
	if !offersCommand(errOut, "luma-backlog rank WORK-0001 <--first|--last|--before <ref>|--after <ref>>") {
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
	if !strings.Contains(errOut, "rank orders records within a status") {
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

// Repair gives every work item a rank, and running it twice changes nothing.
//
// Convergence is the property that makes a whole-corpus rewrite safe: the
// result is a pure function of creation order, so two people repairing the
// same state produce identical files and the merge resolves itself.
func TestRepairRanksEveryRecordAndConverges(t *testing.T) {
	app, project := initialized(t)
	for _, title := range []string{"Alpha", "Bravo", "Charlie"} {
		run(t, app, "work-item", "new", title)
	}
	// An unranked record, the case this exists for: hand-written records and
	// everything created before creation wrote a rank.
	if code, _, e := run(t, app, "set", "WORK-0002", "kind=defect"); code != ExitOK {
		t.Fatalf("set failed: %s", e)
	}
	stripRank(t, app, project, "WORK-0002")

	code, out, e := run(t, app, "rank", "repair")
	if code != ExitOK {
		t.Fatalf("repair failed: %s", e)
	}
	if !strings.Contains(out, "unranked → ") {
		t.Errorf("repair did not report filling an absent rank:\n%s", out)
	}
	for _, key := range []string{"WORK-0001", "WORK-0002", "WORK-0003"} {
		_, shown, _ := run(t, app, "show", key)
		if !strings.Contains(shown, "rank") {
			t.Errorf("%s still has no rank after repair:\n%s", key, shown)
		}
	}

	// Twice is the test. A repair that moves records every time it runs is a
	// repair nobody can put in a pipeline.
	_, again, _ := run(t, app, "rank", "repair")
	if !strings.Contains(again, "already carries the rank it should") {
		t.Errorf("repair is not idempotent:\n%s", again)
	}
}

// Repair writes nothing under --dry-run.
func TestRepairDryRunWritesNothing(t *testing.T) {
	app, project := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	stripRank(t, app, project, "WORK-0001")

	_, out, _ := run(t, app, "rank", "repair", "--dry-run")
	if !strings.Contains(out, "would rank") || !strings.Contains(out, "nothing was written") {
		t.Errorf("--dry-run did not say it wrote nothing:\n%s", out)
	}
	if _, shown, _ := run(t, app, "show", "WORK-0001", "--json"); strings.Contains(shown, `"rank": "0`) {
		t.Errorf("--dry-run wrote a rank:\n%s", shown)
	}
}

// stripRank removes a record's rank on disk, which no command will do --- the
// whole design writes rank with status and never alone.
func stripRank(t *testing.T, app *App, project, key string) {
	t.Helper()
	_, out, _ := run(t, app, "show", key, "--json")
	var rec struct{ Path string }
	if err := json.Unmarshal([]byte(out), &rec); err != nil {
		t.Fatalf("show --json: %v\n%s", err, out)
	}
	full := filepath.Join(project, ".luma", rec.Path)
	body, err := os.ReadFile(full)
	if err != nil {
		t.Fatalf("read %s: %v", full, err)
	}
	var kept []string
	for _, line := range strings.Split(string(body), "\n") {
		if !strings.HasPrefix(line, "rank: ") {
			kept = append(kept, line)
		}
	}
	if err := os.WriteFile(full, []byte(strings.Join(kept, "\n")), 0o644); err != nil {
		t.Fatalf("write %s: %v", full, err)
	}
}
