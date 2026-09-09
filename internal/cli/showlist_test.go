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
		{"work-item", "new", "Payments v2"},
		{"outcome", "new", "The retry queue drains", "-w", "payments-v2"},
		{"task", "new", "Add the retry queue", "-w", "payments-v2"},
		{"work-item", "new", "Search relevance"},
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

// A record type that declares no workflow status must render without one
// rather than inventing a blank-looking value. The project record is the case,
// and `show` is what reaches it --- no listing does, since a listing is scoped
// to one unit and the project record is not one.
func TestShowTableReportsNoStatusForATypeThatDeclaresNone(t *testing.T) {
	app := withProject(t)
	code, out, errOut := run(t, app, "show", "PROJECT")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	checkGolden(t, "show-project-table", out)
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
	code, out, errOut := run(t, app, "work-item", "list", "--json")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	checkGolden(t, "list-json", out)
}

func TestListFilteredJSONShape(t *testing.T) {
	app := populated(t)
	code, out, _ := run(t, app, "outcome", "list", "--json")
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
	run(t, app, "work-item", "new", "Payments alpha")
	run(t, app, "work-item", "new", "Payments beta")

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

// A path read out of a listing should be typeable back in. Before this, only
// the full on-disk path resolved --- WORK-0031/tasks/<slug>, the form a person
// actually sees and repeats, returned "nothing matches".
func TestShowResolvesAKeyScopedPath(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "task", "new", "Add the queue", "-w", "payments-v2")

	for _, ref := range []string{
		"WORK-0001/tasks/add-the-queue",
		"work-0001/tasks/add-the-queue",             // keys match case-insensitively
		"WORK-0001-payments-v2/tasks/add-the-queue", // the joined name
		"payments-v2/tasks/add-the-queue",           // the slug half
		"WORK-0001/tasks/add-the-queue.md",          // extension is not wrong to type
	} {
		code, out, errOut := run(t, app, "show", ref)
		if code != ExitOK {
			t.Errorf("%s: exit = %d, want %d\n%s", ref, code, ExitOK, errOut)
			continue
		}
		if !strings.Contains(out, "Add the queue") {
			t.Errorf("%s: resolved to the wrong record:\n%s", ref, out)
		}
	}
}

// Scoping must not become a second way to guess. An unknown work item is not
// found rather than falling back to a loose match on the tail.
func TestAKeyScopedPathWithAnUnknownScopeIsNotFound(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "task", "new", "Add the queue", "-w", "payments-v2")

	if code, _, _ := run(t, app, "show", "WORK-9999/tasks/add-the-queue"); code != ExitNotFound {
		t.Errorf("exit = %d, want %d (not found)", code, ExitNotFound)
	}
}

// The bare slug is what people typed before scoping existed and must keep
// working --- scoping is an addition, not a replacement.
func TestABareSlugStillResolves(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "task", "new", "Add the queue", "-w", "payments-v2")

	code, out, errOut := run(t, app, "show", "add-the-queue")
	if code != ExitOK {
		t.Fatalf("exit = %d: %s", code, errOut)
	}
	if !strings.Contains(out, "Add the queue") {
		t.Errorf("resolved to the wrong record:\n%s", out)
	}
}

// The most-used read command, and the one the noun-verb shape exists for.
func TestANounListsItsOwnRecords(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "outcome", "new", "The queue drains", "-w", "payments-v2")

	code, out, errOut := run(t, app, "work-item", "list")
	if code != ExitOK {
		t.Fatalf("exit = %d: %s", code, errOut)
	}
	if !strings.Contains(out, "Payments v2") {
		t.Errorf("work-item list did not list the work item:\n%s", out)
	}
	if strings.Contains(out, "The queue drains") {
		t.Errorf("work-item list included an outcome:\n%s", out)
	}
}

// `list` is `work-item list`. The tool is called backlog; listing the backlog
// means listing work items, not a run of every record type interleaved.
func TestBareListIsWorkItemList(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "outcome", "new", "The queue drains", "-w", "payments-v2")

	code, out, errOut := run(t, app, "list")
	if code != ExitOK {
		t.Fatalf("exit = %d: %s", code, errOut)
	}
	if !strings.Contains(out, "Payments v2") {
		t.Errorf("list did not list the work item:\n%s", out)
	}
	if strings.Contains(out, "The queue drains") {
		t.Errorf("list included an outcome:\n%s", out)
	}
}

// The redundant path is gone: a record type is never a positional argument.
func TestListDoesNotTakeAUnitPositionally(t *testing.T) {
	app, _ := initialized(t)
	if code, _, _ := run(t, app, "list", "work-item"); code != ExitUsage {
		t.Errorf("exit = %d, want %d --- `list work-item` said the same thing twice", code, ExitUsage)
	}
}

// The tree is what replaced listing every type interleaved: a work item with
// what actually hangs off it, rather than a run of unrelated records.
func TestListTreeShowsChildrenBeneathTheirWorkItem(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "outcome", "new", "The queue drains", "-w", "payments-v2")
	run(t, app, "task", "new", "Add the queue", "-w", "payments-v2")

	code, out, errOut := run(t, app, "list", "--tree")
	if code != ExitOK {
		t.Fatalf("exit = %d: %s", code, errOut)
	}
	for _, want := range []string{"Payments v2", "The queue drains", "Add the queue"} {
		if !strings.Contains(out, want) {
			t.Errorf("tree missed %q:\n%s", want, out)
		}
	}
	// Without --tree the same command shows work items only.
	_, flat, _ := run(t, app, "list")
	if strings.Contains(flat, "Add the queue") {
		t.Errorf("a flat listing included a task:\n%s", flat)
	}
}

// The filter narrows the work items, never their children. Asking for
// prepared work items and being shown only their prepared tasks would hide the
// ones nobody has started, which is usually the reason for looking.
func TestATreeFilterNarrowsWorkItemsNotChildren(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "task", "new", "Add the queue", "-w", "payments-v2")
	run(t, app, "work-item", "transition", "payments-v2", "todo")

	code, out, errOut := run(t, app, "list", "--tree", "--status", "todo")
	if code != ExitOK {
		t.Fatalf("exit = %d: %s", code, errOut)
	}
	if !strings.Contains(out, "Add the queue") {
		t.Errorf("filtering the work items also filtered their children:\n%s", out)
	}
}

// Only work items have anything hanging off them.
func TestOnlyWorkItemsOfferATree(t *testing.T) {
	app, _ := initialized(t)
	if code, _, _ := run(t, app, "task", "list", "--tree"); code != ExitUsage {
		t.Errorf("task list accepted --tree")
	}
}

// A listing shows the key, the status and the title --- and not the type,
// which repeats the command, nor the slug, which repeats the title.
func TestAListingShowsKeyStatusAndTitle(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")

	_, out, _ := run(t, app, "work-item", "list")
	if !strings.Contains(out, "KEY") || !strings.Contains(out, "WORK-0001") {
		t.Errorf("the listing did not show the key:\n%s", out)
	}
	for _, unwanted := range []string{"TYPE", "work-item", "payments-v2"} {
		if strings.Contains(out, unwanted) {
			t.Errorf("the listing still shows %q:\n%s", unwanted, out)
		}
	}
}

// A child is marked by what it is rather than named. Its slug is its title in
// kebab case, so printing it would put the same sentence on the row twice.
func TestATreeMarksChildrenByType(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "task", "new", "Add the queue", "-w", "payments-v2")
	run(t, app, "outcome", "new", "The queue drains", "-w", "payments-v2")

	_, out, _ := run(t, app, "work-item", "list", "--tree")
	for _, want := range []string{"TASK", "OUT", "Add the queue", "The queue drains"} {
		if !strings.Contains(out, want) {
			t.Errorf("tree missing %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "add-the-queue") {
		t.Errorf("tree printed a child's slug beside its own title:\n%s", out)
	}
}

// Tasks come before outcomes. Tasks are what somebody looks at almost every
// time; an outcome reads "unverified" for nearly the whole life of a work item,
// so leading with them puts constant text where the useful rows belong.
func TestTasksComeBeforeOutcomes(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "outcome", "new", "The queue drains", "-w", "payments-v2")
	run(t, app, "task", "new", "Add the queue", "-w", "payments-v2")

	_, out, _ := run(t, app, "work-item", "list", "--tree")
	task, outcome := strings.Index(out, "TASK"), strings.Index(out, "OUT ")
	if task < 0 || outcome < 0 {
		t.Fatalf("tree did not contain both:\n%s", out)
	}
	if task > outcome {
		t.Errorf("outcomes came before tasks:\n%s", out)
	}
}

// The last child closes the branch, so the eye finds where one work item ends
// without counting rows.
func TestTheLastChildClosesTheBranch(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "task", "new", "Add the queue", "-w", "payments-v2")
	run(t, app, "task", "new", "Drain it", "-w", "payments-v2")

	_, out, _ := run(t, app, "work-item", "list", "--tree")
	if strings.Count(out, "└─") != 1 {
		t.Errorf("expected exactly one closing branch:\n%s", out)
	}
}

// A listing carries each record's stamps, so a caller can answer "what changed
// recently" without opening every file. Additive to the --json shape
// (spec.md §9.9).
func TestListingCarriesStamps(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")

	_, out, _ := run(t, app, "work-item", "list", "--json")
	if !strings.Contains(out, `"created"`) {
		t.Errorf("listing omitted created:\n%s", out)
	}
	if !strings.Contains(out, `"at"`) || !strings.Contains(out, `"by"`) {
		t.Errorf("a stamp lost its parts:\n%s", out)
	}
}

// A record nobody has edited has no modified stamp, and the field is absent
// rather than empty --- absent says the record has none, where {"by":"","at":""}
// says the tool read one and found nothing in it.
func TestAnUneditedRecordHasNoModifiedStamp(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")

	_, out, _ := run(t, app, "work-item", "list", "--json")
	if strings.Contains(out, `"modified"`) {
		t.Errorf("an unedited record carried a modified stamp:\n%s", out)
	}

	run(t, app, "work-item", "transition", "WORK-0001", "todo")
	_, after, _ := run(t, app, "work-item", "list", "--json")
	if !strings.Contains(after, `"modified"`) {
		t.Errorf("an edited record has no modified stamp:\n%s", after)
	}
}

// The question anybody asks first, and the one that could not be expressed:
// --status matches a single value, so "everything except closed" needed a
// second language to answer.
func TestOpenExcludesOnlyTheTerminalStatus(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Still going")
	run(t, app, "work-item", "new", "Finished")
	run(t, app, "work-item", "close", "finished", "canceled")

	_, out, _ := run(t, app, "work-item", "list", "--open")
	if strings.Contains(out, "Finished") {
		t.Errorf("--open included a closed work item:\n%s", out)
	}
	if !strings.Contains(out, "Still going") {
		t.Errorf("--open dropped an open work item:\n%s", out)
	}
}

// A record with no status reads as the first rung, which is never terminal.
func TestOpenIncludesARecordWithNoStatus(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Unstamped")
	_, out, _ := run(t, app, "work-item", "list", "--open")
	if !strings.Contains(out, "Unstamped") {
		t.Errorf("--open dropped a record that has not ended:\n%s", out)
	}
}

// Passing both says two different things about one listing, and picking one
// for the caller leaves them believing something untrue about what they got.
func TestOpenAndStatusTogetherAreRefused(t *testing.T) {
	app, _ := initialized(t)
	if code, _, _ := run(t, app, "work-item", "list", "--open", "-s", "closed"); code != ExitUsage {
		t.Error("--open and --status were accepted together")
	}
}
