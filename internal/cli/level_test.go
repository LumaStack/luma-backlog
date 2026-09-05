package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestADecisionNeedsItsLevelStated(t *testing.T) {
	// The cost of getting a level wrong is asymmetric: a decision filed at the
	// wrong level is not visibly broken, it is simply somewhere nobody looks.
	// So it is stated rather than defaulted.
	app, _ := initialized(t)
	code, _, errOut := run(t, app, "new", "decision", "Use the new queue")
	if code == ExitOK {
		t.Fatal("a decision was created with no level stated")
	}
	for _, want := range []string{"level", "--work-item", "--project"} {
		if !strings.Contains(errOut, want) {
			t.Errorf("the error did not mention %q:\n%s", want, errOut)
		}
	}
}

func TestEachLevelLandsWhereItBelongs(t *testing.T) {
	app, project := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	for _, args := range [][]string{
		{"new", "decision", "Use the new queue", "--project"},
		{"new", "decision", "Retry inside the worker", "-w", "payments-v2"},
	} {
		if code, _, e := run(t, app, args...); code != ExitOK {
			t.Fatalf("%v failed: %s", args, e)
		}
	}
	// Numbering is one sequence across both levels, so a number means one
	// record wherever it sits.
	for _, rel := range []string{
		"records/decisions/ADR-0001-use-the-new-queue.md",
		wiPath(t, project, "payments-v2", "decisions", "ADR-0002-retry-inside-the-worker.md"),
	} {
		if _, err := os.Stat(filepath.Join(project, ".luma", rel)); err != nil {
			t.Errorf("expected %s: %v", rel, err)
		}
	}
}

func TestBothLevelsAtOnceIsRefused(t *testing.T) {
	app, _ := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	code, _, errOut := run(t, app, "new", "decision", "Contradiction", "--project", "-w", "payments-v2")
	if code == ExitOK {
		t.Fatal("a decision was created with two levels")
	}
	if !strings.Contains(errOut, "different levels") {
		t.Errorf("the error did not say why:\n%s", errOut)
	}
}

func TestStandingInAWorkItemDoesNotDecideTheLevel(t *testing.T) {
	// The whole point. A person at a terminal is usually standing where they
	// are working; an agent runs from the repository root whatever it is
	// doing. Letting the working directory answer produced a repository where
	// every decision is project-level and none was placed on purpose.
	app, project := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	app.WorkingDir = filepath.Join(project, ".luma", "backlog", "work-items", "payments-v2")

	code, _, errOut := run(t, app, "new", "decision", "Retry inside the worker")
	if code == ExitOK {
		t.Fatal("standing inside a work item decided the level")
	}
	if !strings.Contains(errOut, "level") {
		t.Errorf("the error did not say what was missing:\n%s", errOut)
	}

	// And the working directory still answers WHICH work item, which is the
	// job it is good at — only the level is withheld from it.
	if code, _, e := run(t, app, "new", "task", "Route jobs"); code != ExitOK {
		t.Fatalf("a task lost its work item from the working directory: %s", e)
	}
	if _, err := os.Stat(filepath.Join(project, ".luma",
		wiPath(t, project, "payments-v2", "tasks", "route-jobs.md"))); err != nil {
		t.Errorf("the task did not land in the work item: %v", err)
	}
}
