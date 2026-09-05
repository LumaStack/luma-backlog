package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// withTwoOutcomes builds the shape that made this unsafe: one outcome passing,
// one that will be broken by the caller. One good outcome was enough to satisfy
// a count that had quietly lost its denominator.
func withTwoOutcomes(t *testing.T) (*App, string) {
	t.Helper()
	app, project := initialized(t)
	for _, args := range [][]string{
		{"new", "work-item", "Payments v2", "--kind", "change"},
		{"new", "outcome", "Latency holds", "-w", "payments-v2"},
		{"new", "outcome", "The retry queue drains", "-w", "payments-v2"},
		{"verify", "latency-holds", "-e", "measured"},
		{"verify", "the-retry-queue-drains", "-e", "measured"},
	} {
		if code, _, e := run(t, app, args...); code != ExitOK {
			t.Fatalf("%v failed: %s", args, e)
		}
	}
	return app, project
}

func breakOutcome(t *testing.T, project, rel string) {
	t.Helper()
	path := filepath.Join(project, ".luma", rel)
	if err := os.WriteFile(path, []byte("type: outcome\ntitle: broken\n---\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestDeliveredIsRefusedWhenAnOutcomeCannotBeRead(t *testing.T) {
	// The reproduction. Before this, both outcomes verified and one file
	// broken produced exit 0 and a work item closed as delivered, with an
	// outcome sitting on disk that nothing had checked.
	app, project := withTwoOutcomes(t)
	breakOutcome(t, project, wiPath(t, project, "payments-v2", "outcomes", "the-retry-queue-drains.md"))

	code, _, errOut := run(t, app, "close", "payments-v2", "--reason", "delivered")
	if code == ExitOK {
		t.Fatalf("delivered was allowed over an unreadable outcome:\n%s", errOut)
	}
	if code != ExitRefused {
		t.Errorf("exit = %d, want ExitRefused (%d)", code, ExitRefused)
	}
	if !strings.Contains(errOut, "the-retry-queue-drains.md") {
		t.Errorf("the refusal did not name the file:\n%s", errOut)
	}
}

func TestOtherReasonsCloseOverAnUnreadableOutcome(t *testing.T) {
	// The way out, and the reason no force flag is needed. Canceling claims
	// nothing about evidence, so it needs no count — gating it would make it
	// impossible to stop work precisely because the record was broken.
	app, project := withTwoOutcomes(t)
	breakOutcome(t, project, wiPath(t, project, "payments-v2", "outcomes", "the-retry-queue-drains.md"))

	code, out, errOut := run(t, app, "close", "payments-v2", "--reason", "canceled")
	if code != ExitOK {
		t.Fatalf("canceling was blocked by an unreadable outcome: exit %d, %s", code, errOut)
	}
	if !strings.Contains(out, "closed") {
		t.Errorf("the work item did not close:\n%s", out)
	}
}

func TestASkipElsewhereDoesNotBlockThisWorkItem(t *testing.T) {
	// A corrupt record in somebody else's work is a real problem and not this
	// caller's. Blocking every close in the repository on it would be a
	// refusal nobody could act on from where they are standing.
	app, project := withTwoOutcomes(t)
	if code, _, e := run(t, app, "new", "work-item", "Search relevance", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	if code, _, e := run(t, app, "new", "outcome", "Results rank well", "-w", "search-relevance"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	breakOutcome(t, project, wiPath(t, project, "search-relevance", "outcomes", "results-rank-well.md"))

	code, out, errOut := run(t, app, "close", "payments-v2", "--reason", "delivered")
	if code != ExitOK {
		t.Fatalf("a skip in another work item blocked this one: exit %d, %s", code, errOut)
	}
	if !strings.Contains(out, "closed") {
		t.Errorf("the work item did not close:\n%s", out)
	}
}

func TestAnUnreadableOutcomeIsStillReported(t *testing.T) {
	// Refusing is the new half; saying so was already true and must stay true,
	// because a refusal that does not name the file leaves somebody hunting.
	app, project := withTwoOutcomes(t)
	breakOutcome(t, project, wiPath(t, project, "payments-v2", "outcomes", "the-retry-queue-drains.md"))

	_, _, errOut := run(t, app, "close", "payments-v2", "--reason", "canceled")
	if !strings.Contains(errOut, "skipped") {
		t.Errorf("the skip went unreported on a reason that closes freely:\n%q", errOut)
	}
}
