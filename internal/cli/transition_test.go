package cli

import (
	"strings"
	"testing"
)

// A status change is an operation, not a field write. `set` refuses the field
// for the same reason it refuses `rank` (ADR-0005), and the refusal has to say
// what to use instead --- removing a way in that fails blankly is worse than
// not removing it (spec.md §9.9).
func TestSetRefusesWorkflowStatusAndNamesTransition(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, errOut := run(t, app, "set", "WORK-0001", "workflow_status=todo")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "transition") {
		t.Errorf("the refusal did not name what to use instead:\n%s", errOut)
	}

	_, out, _ := run(t, app, "show", "WORK-0001", "--json")
	if !strings.Contains(out, `"workflow_status": "captured"`) {
		t.Errorf("a refused set changed the record anyway:\n%s", out)
	}
}

// Absence reads as the first configured value (spec.md §4.2), so unsetting the
// field silently sends a record to the bottom of the ladder --- and does it
// without writing rank, which is the pair ADR-0005 says never separates.
func TestSetRefusesToUnsetWorkflowStatus(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "work-item", "transition", "WORK-0001", "todo")

	code, _, errOut := run(t, app, "set", "WORK-0001", "--unset", "workflow_status")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "transition") {
		t.Errorf("the refusal did not name what to use instead:\n%s", errOut)
	}

	_, out, _ := run(t, app, "show", "WORK-0001", "--json")
	if !strings.Contains(out, `"workflow_status": "todo"`) {
		t.Errorf("a refused unset changed the record anyway:\n%s", out)
	}
}

// Rank is written by the same operation, so unsetting it separates a pair that
// no command will write apart.
func TestSetRefusesToUnsetRank(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, _ := run(t, app, "set", "WORK-0001", "--unset", "rank")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
}

// Every rung the ladder carries is reachable, and each writes both fields.
func TestTransitionReachesEveryRungBelowTerminal(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	for _, status := range []string{"unprepared", "preparing", "prepared", "todo", "in_progress"} {
		if code, _, e := run(t, app, "work-item", "transition", "WORK-0001", status); code != ExitOK {
			t.Fatalf("transition to %s failed: %s", status, e)
		}
		_, out, _ := run(t, app, "show", "WORK-0001", "--json")
		if !strings.Contains(out, `"workflow_status": "`+status+`"`) {
			t.Errorf("%s was not written:\n%s", status, out)
		}
		if !strings.Contains(out, `"rank"`) {
			t.Errorf("transition to %s left the record unranked:\n%s", status, out)
		}
	}
}

// `close` carries its own refusals, its own disposition vocabulary and its own
// --force. A second door into the terminal status would skip all three, which
// is the defect recorded as WORK-0073.
func TestTransitionToTerminalIsRefusedAndNamesClose(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, errOut := run(t, app, "work-item", "transition", "WORK-0001", "closed")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "close") {
		t.Errorf("the refusal did not name close:\n%s", errOut)
	}

	_, out, _ := run(t, app, "show", "WORK-0001", "--json")
	if strings.Contains(out, `"workflow_status": "closed"`) {
		t.Error("a refused transition closed the record anyway")
	}
}

// A status the ladder does not carry has no place in the order, and guessing
// one would file the record somewhere nobody chose.
func TestTransitionToAnUnknownStatusIsRefused(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, errOut := run(t, app, "work-item", "transition", "WORK-0001", "marinating")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "marinating") {
		t.Errorf("the refusal did not name the status:\n%s", errOut)
	}
	// Naming the status is not enough on its own --- a caller who guessed wrong
	// needs to see what the ladder actually holds.
	if !strings.Contains(errOut, "captured") || !strings.Contains(errOut, "in_progress") {
		t.Errorf("the refusal did not show the ladder:\n%s", errOut)
	}
}

// Exit codes are distinguishable because an agent's next move depends on why
// something failed (spec.md §9.4, ADR-0006). A missing record is not a usage
// error --- the invocation was fine and the record was not there.
func TestTransitionOnAMissingRecordExitsNotFound(t *testing.T) {
	app, _ := initialized(t)

	code, _, _ := run(t, app, "work-item", "transition", "WORK-9999", "todo")
	if code != ExitNotFound {
		t.Errorf("exit = %d, want %d", code, ExitNotFound)
	}
}

// Only work items carry a workflow status.
func TestTransitionRefusesANonWorkItem(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "outcome", "new", "The queue drains", "-w", "WORK-0001")

	code, _, errOut := run(t, app, "work-item", "transition", "the-queue-drains", "todo")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "outcome") {
		t.Errorf("the refusal did not say what the record is:\n%s", errOut)
	}
}

// The optimistic-concurrency contract `set` offers is offered here too: a write
// that would clobber a change it never saw is refused rather than applied
// (spec.md §6.3), and conflict is its own exit code because retrying is correct.
func TestTransitionRefusesAStaleHash(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, _ := run(t, app, "work-item", "transition", "WORK-0001", "todo",
		"--if-unchanged", "0000000000000000")
	if code != ExitConflict {
		t.Errorf("exit = %d, want %d", code, ExitConflict)
	}
}

// `move` is not a verb this interface has, with one exception: it may be an
// alias. Help and output still call it transition (spec.md §9.2).
func TestMoveIsAnAliasAndNotAName(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	if code, _, e := run(t, app, "work-item", "move", "WORK-0001", "todo"); code != ExitOK {
		t.Fatalf("the alias did not resolve: %s", e)
	}

	_, help, _ := run(t, app, "work-item", "--help")
	if strings.Contains(strings.ToLower(help), "move") {
		t.Errorf("move is named in help rather than only aliased:\n%s", help)
	}
	if !strings.Contains(help, "transition") {
		t.Errorf("transition is missing from help:\n%s", help)
	}
}

// Only work items carry a workflow status or a rank, so the noun adds a word
// and removes no ambiguity. Reaching rank only under the noun is why it read as
// missing to somebody who went looking for it.
func TestTransitionAndRankAreReachableWithoutTheNoun(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	if code, _, e := run(t, app, "transition", "WORK-0001", "todo"); code != ExitOK {
		t.Errorf("transition is not reachable at the top level: %s", e)
	}
	if code, _, e := run(t, app, "rank", "WORK-0001", "--first"); code != ExitOK {
		t.Errorf("rank is not reachable at the top level: %s", e)
	}
	// The same operation, not a second implementation.
	_, out, _ := run(t, app, "show", "WORK-0001", "--json")
	if !strings.Contains(out, `"workflow_status": "todo"`) {
		t.Errorf("the top-level form did not write the status:\n%s", out)
	}
}
