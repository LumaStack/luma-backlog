package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// populated builds a small, fixed backlog so output is byte-stable.
func populated(t *testing.T) *App {
	t.Helper()
	app, _ := initialized(t)
	seed(t, app)
	return app
}

func seed(t *testing.T, app *App) {
	t.Helper()
	for _, args := range [][]string{
		{"new", "work-item", "Payments v2"},
		{"new", "outcome", "The retry queue drains", "-w", "payments-v2"},
		{"new", "task", "Add the retry queue", "-w", "payments-v2"},
		{"new", "work-item", "Search relevance"},
	} {
		if code, _, e := run(t, app, args...); code != ExitOK {
			t.Fatalf("%v failed: %s", args, e)
		}
	}
}

// withProject adds a record whose type declares no workflow_status at all.
//
// A luma/project has no lifecycle — it is not todo, not done, not anything —
// and neither does it carry a stage. `new` cannot create one and `init` does
// not write one, so the fixture puts it there directly. Without it the whole
// non-worked, stage-less case is uncovered by construction.
func withProject(t *testing.T) *App {
	t.Helper()
	app, project := initialized(t)
	seed(t, app)

	const projectRecord = `---
type: luma/project
title: Example
disclosure_level: public
description: A project record, which has no workflow status.
---
`
	path := filepath.Join(project, ".luma", "PROJECT.md")
	if err := os.WriteFile(path, []byte(projectRecord), 0o644); err != nil {
		t.Fatal(err)
	}
	return app
}

func TestListTableReportsNoStatusForATypeThatDeclaresNone(t *testing.T) {
	app := withProject(t)
	code, out, errOut := run(t, app, "list")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	checkGolden(t, "list-with-project-table", out)
}

func TestShowProjectJSONOmitsStatus(t *testing.T) {
	app := withProject(t)
	code, out, errOut := run(t, app, "show", "PROJECT", "--json")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	checkGolden(t, "show-project-json", out)
}

func TestListJSONShape(t *testing.T) {
	app := populated(t)
	code, out, errOut := run(t, app, "list", "--json")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	checkGolden(t, "list-json", out)
}

func TestListFilteredJSONShape(t *testing.T) {
	app := populated(t)
	code, out, _ := run(t, app, "list", "outcome", "--json")
	if code != ExitOK {
		t.Fatalf("exit = %d", code)
	}
	checkGolden(t, "list-outcome-json", out)
}

func TestShowJSONShape(t *testing.T) {
	app := populated(t)
	code, out, errOut := run(t, app, "show", "the-retry-queue-drains", "--json")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	checkGolden(t, "show-json", out)
}

func TestEmptyListIsAnEmptyArrayNotNull(t *testing.T) {
	// A caller iterating the response should not have to special-case
	// "nothing yet" — null and [] are different shapes.
	app, _ := initialized(t)
	code, out, _ := run(t, app, "list", "--json")
	if code != ExitOK {
		t.Fatalf("exit = %d", code)
	}
	if strings.TrimSpace(out) != "[]" {
		t.Errorf("empty listing = %q, want []", strings.TrimSpace(out))
	}
}

func TestEmptyListIsNotAnError(t *testing.T) {
	// An empty backlog and an over-narrow filter are both ordinary. Exiting
	// non-zero would make a caller treat "none yet" as a failure.
	app, _ := initialized(t)
	if code, _, _ := run(t, app, "list"); code != ExitOK {
		t.Errorf("empty listing exited %d", code)
	}
}

func TestListFilters(t *testing.T) {
	app := populated(t)

	_, out, _ := run(t, app, "list", "-w", "payments-v2")
	if strings.Contains(out, "search-relevance") {
		t.Errorf("work item filter leaked another work item:\n%s", out)
	}

	_, out, _ = run(t, app, "list", "-s", "idea")
	if strings.Contains(out, "add-the-retry-queue") {
		t.Errorf("status filter matched a task with a different status:\n%s", out)
	}
}

func TestShowResolvesByPrefix(t *testing.T) {
	app := populated(t)
	code, out, errOut := run(t, app, "show", "the-retry")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	if !strings.Contains(out, "The retry queue drains") {
		t.Errorf("prefix did not resolve:\n%s", out)
	}
}

func TestShowRefusesAnAmbiguousReference(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "new", "work-item", "Payments alpha")
	run(t, app, "new", "work-item", "Payments beta")

	code, _, errOut := run(t, app, "show", "payments")
	if code == ExitOK {
		t.Fatal("an ambiguous reference resolved instead of erroring")
	}
	// Both candidates must be named: guessing is how the wrong record gets
	// edited and nobody finds out until later.
	for _, want := range []string{"payments-alpha", "payments-beta"} {
		if !strings.Contains(errOut, want) {
			t.Errorf("error did not list %s:\n%s", want, errOut)
		}
	}
}

func TestShowReportsNotFound(t *testing.T) {
	app := populated(t)
	code, _, _ := run(t, app, "show", "nothing-like-this")
	if code != ExitNotFound {
		t.Errorf("exit = %d, want %d (not found)", code, ExitNotFound)
	}
}

func TestListIgnoresNonRecords(t *testing.T) {
	// Journals, Type Definitions, and the bundle root are not units, and a
	// listing that included them would be wrong in a way that looks right.
	app := populated(t)
	_, out, _ := run(t, app, "list", "--json")
	for _, unwanted := range []string{"journal.md", "_types/", `"path": "index.md"`} {
		if strings.Contains(out, unwanted) {
			t.Errorf("listing included %s:\n%s", unwanted, out)
		}
	}
}
