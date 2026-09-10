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
	// in_progress refuses without one, which is its own test below.
	run(t, app, "outcome", "new", "The queue drains", "-w", "WORK-0001")

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

// backlog-move records that nothing captures the reasoning for crossing a gate.
// A reason is somewhere for it to go when there is some --- optional, because
// prose demanded of somebody with nothing to say is prose nobody reads.
func TestTransitionJournalsAReasonWhenGivenOne(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	if code, _, e := run(t, app, "transition", "WORK-0001", "prepared",
		"--reason", "the outcomes pass their checks"); code != ExitOK {
		t.Fatalf("transition failed: %s", e)
	}

	_, journal, _ := run(t, app, "work-item", "journal", "-w", "WORK-0001")
	if !strings.Contains(journal, "the outcomes pass their checks") {
		t.Errorf("the reason did not reach the journal:\n%s", journal)
	}
	// Which crossing it explains is part of the entry, or a later reader has
	// prose with nothing to attach it to.
	if !strings.Contains(journal, "captured → prepared") {
		t.Errorf("the entry did not say which crossing:\n%s", journal)
	}
}

// Absent means nothing extra is written. A journal line on every crossing would
// be an event log, and the journal is deliberately not one (spec.md §5.5).
func TestTransitionWithoutAReasonWritesNoJournalLine(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "transition", "WORK-0001", "todo")

	_, journal, _ := run(t, app, "work-item", "journal", "-w", "WORK-0001")
	if strings.Contains(journal, "→ todo") {
		t.Errorf("a crossing with no reason wrote to the journal anyway:\n%s", journal)
	}
}

// Reopening is the one crossing that should always say why: the closed entry
// keeps the ending and no field holds the un-ending. Strongly encouraged rather
// than required --- it warns, the crossing happens, and stdout stays clean.
func TestReopeningWithoutAReasonIsAdvisedAgainstButAllowed(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "work-item", "close", "WORK-0001", "canceled")

	code, out, errOut := run(t, app, "transition", "WORK-0001", "todo")
	if code != ExitOK {
		t.Fatalf("reopening was refused: exit %d, %s", code, errOut)
	}
	if !strings.Contains(errOut, "why") {
		t.Errorf("reopening without a reason was not advised against:\n%s", errOut)
	}
	if strings.Contains(out, "why") {
		t.Errorf("advice reached stdout:\n%s", out)
	}

	_, shown, _ := run(t, app, "show", "WORK-0001", "--json")
	if !strings.Contains(shown, `"workflow_status": "todo"`) {
		t.Errorf("the reopen did not happen:\n%s", shown)
	}
}

// With a reason, there is nothing to advise about.
func TestReopeningWithAReasonIsSilent(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "work-item", "close", "WORK-0001", "canceled")

	_, _, errOut := run(t, app, "transition", "WORK-0001", "todo",
		"--reason", "the budget came back")
	if strings.Contains(errOut, "why") {
		t.Errorf("a reopen carrying a reason was advised against anyway:\n%s", errOut)
	}
}

// The ordinary rungs expect nothing. captured to unprepared always has the same
// answer, and a prompt whose answer is always the same teaches people to type
// past it.
func TestAnOrdinaryCrossingIsNotAdvisedAbout(t *testing.T) {
	app, _ := initialized(t)
	// A kind, or leaving the pile warns about its absence --- its own test.
	run(t, app, "work-item", "new", "Alpha", "--kind", "change")

	_, _, errOut := run(t, app, "transition", "WORK-0001", "unprepared")
	if strings.TrimSpace(errOut) != "" {
		t.Errorf("an ordinary crossing said something:\n%s", errOut)
	}
}

// The one refusal a transition carries: starting work nobody can tell is
// finished. Deliberately narrow --- not "are the outcomes good" but "does one
// exist", because any opinion beyond zero would be the tool holding a view
// about how work gets defined.
func TestStartingWithNoOutcomesIsRefused(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, errOut := run(t, app, "transition", "WORK-0001", "in_progress")
	if code != ExitRefused {
		t.Errorf("exit = %d, want %d", code, ExitRefused)
	}
	if !strings.Contains(errOut, "outcome new") {
		t.Errorf("the refusal did not say how to fix it:\n%s", errOut)
	}

	_, out, _ := run(t, app, "show", "WORK-0001", "--json")
	if strings.Contains(out, `"workflow_status": "in_progress"`) {
		t.Error("a refused transition started the work anyway")
	}
}

// --force proceeds, and says so rather than going quiet: a forced start is
// exactly the thing a later reader needs to know happened.
func TestStartingWithNoOutcomesCanBeForced(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, errOut := run(t, app, "transition", "WORK-0001", "in_progress", "--force")
	if code != ExitOK {
		t.Fatalf("--force did not proceed: exit %d, %s", code, errOut)
	}
	if !strings.Contains(errOut, "forced") {
		t.Errorf("a forced start was silent:\n%s", errOut)
	}
}

// An outcome is all it takes. The bar is zero, not quality.
func TestStartingWithOneOutcomeIsAllowed(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "outcome", "new", "The queue drains", "-w", "WORK-0001")

	if code, _, e := run(t, app, "transition", "WORK-0001", "in_progress"); code != ExitOK {
		t.Fatalf("a work item with an outcome was refused: %s", e)
	}
}

// Leaving the shaping rung is where the work was supposed to be worked out.
// Warned rather than refused: not all work is equal, and some is trivial.
func TestLeavingPreparingWithoutOutcomesOrTasksWarns(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "change")
	run(t, app, "transition", "WORK-0001", "preparing")

	code, out, errOut := run(t, app, "transition", "WORK-0001", "prepared")
	if code != ExitOK {
		t.Fatalf("leaving preparing was refused: exit %d, %s", code, errOut)
	}
	if !strings.Contains(errOut, "no outcomes and no tasks") {
		t.Errorf("leaving preparing unshaped was not warned about:\n%s", errOut)
	}
	if strings.Contains(out, "no outcomes") {
		t.Errorf("the warning reached stdout:\n%s", out)
	}
}

// It names what is actually missing rather than both every time.
func TestLeavingPreparingNamesOnlyWhatIsMissing(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "change")
	run(t, app, "outcome", "new", "The queue drains", "-w", "WORK-0001")
	run(t, app, "transition", "WORK-0001", "preparing")

	_, _, errOut := run(t, app, "transition", "WORK-0001", "prepared")
	if !strings.Contains(errOut, "no tasks") {
		t.Errorf("the missing half was not named:\n%s", errOut)
	}
	if strings.Contains(errOut, "no outcomes") {
		t.Errorf("it complained about outcomes that exist:\n%s", errOut)
	}
}

// An idea is a classification on its way to something else. Carrying one past
// the first gate files unformed work beside formed work.
func TestLeavingThePileAsAnIdeaIsRefused(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "idea")

	code, _, errOut := run(t, app, "transition", "WORK-0001", "unprepared")
	if code != ExitRefused {
		t.Errorf("exit = %d, want %d", code, ExitRefused)
	}
	if !strings.Contains(errOut, "kind=") {
		t.Errorf("the refusal did not say how to fix it:\n%s", errOut)
	}
}

// Every refusal takes --force, and says it was forced.
func TestLeavingThePileAsAnIdeaCanBeForced(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "idea")

	code, _, errOut := run(t, app, "transition", "WORK-0001", "unprepared", "--force")
	if code != ExitOK {
		t.Fatalf("--force did not proceed: exit %d, %s", code, errOut)
	}
	if !strings.Contains(errOut, "forced") {
		t.Errorf("a forced selection was silent:\n%s", errOut)
	}
}

// A kind that is not idea passes freely --- the refusal is about ideas, not
// about classification being present.
func TestLeavingThePileWithARealKindIsSilent(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "defect")

	code, _, errOut := run(t, app, "transition", "WORK-0001", "unprepared")
	if code != ExitOK {
		t.Fatalf("a defect was refused: %s", errOut)
	}
	if strings.TrimSpace(errOut) != "" {
		t.Errorf("an ordinary selection said something:\n%s", errOut)
	}
}

// Announcing a force tells whoever is at the terminal; the journal is what lets
// anybody ask later how often this happens and what it cost. Same reason a
// skipped gate gets a line.
func TestAForcedCrossingIsRecordedInTheJournal(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "change")
	run(t, app, "transition", "WORK-0001", "in_progress", "--force")

	_, journal, _ := run(t, app, "work-item", "journal", "-w", "WORK-0001")
	if !strings.Contains(journal, "FORCED") {
		t.Errorf("a forced crossing left no journal entry:\n%s", journal)
	}
	if !strings.Contains(journal, "no outcomes") {
		t.Errorf("the entry did not say what was overridden:\n%s", journal)
	}
	if !strings.Contains(journal, "→ in_progress") {
		t.Errorf("the entry did not say which crossing:\n%s", journal)
	}
}

// A force with no refusal to override records nothing. --force is not a mode.
func TestForceWithNothingToOverrideRecordsNothing(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "change")
	run(t, app, "outcome", "new", "The queue drains", "-w", "WORK-0001")
	run(t, app, "transition", "WORK-0001", "in_progress", "--force")

	_, journal, _ := run(t, app, "work-item", "journal", "-w", "WORK-0001")
	if strings.Contains(journal, "FORCED") {
		t.Errorf("--force wrote an entry with nothing to override:\n%s", journal)
	}
}

// Tasks carry a workflow status and their own ladder. Nothing else can change
// it --- there is no task close --- so transition has to accept them, including
// all the way to their terminal.
func TestATaskTransitionsThroughItsOwnLadder(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "task", "new", "Do the thing", "-w", "WORK-0001")

	for _, status := range []string{"in_progress", "closed"} {
		if code, _, e := run(t, app, "transition", "do-the-thing", status); code != ExitOK {
			t.Fatalf("task transition to %s failed: %s", status, e)
		}
	}
	_, out, _ := run(t, app, "show", "do-the-thing", "--json")
	if !strings.Contains(out, `"workflow_status": "closed"`) {
		t.Errorf("the task did not reach closed:\n%s", out)
	}
	// Only work items are ranked (ADR-0005), so a task gains no rank on the way.
	if strings.Contains(out, `"rank"`) {
		t.Errorf("a task was given a rank:\n%s", out)
	}
}

// The work-item checks are work-item concepts and must not fire on a task ---
// a task has no outcomes and no kind, and would fail every one of them.
func TestATaskIsNotHeldToWorkItemChecks(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "task", "new", "Do the thing", "-w", "WORK-0001")

	code, _, errOut := run(t, app, "transition", "do-the-thing", "in_progress")
	if code != ExitOK {
		t.Fatalf("a task was refused for having no outcomes: %s", errOut)
	}
	if strings.TrimSpace(errOut) != "" {
		t.Errorf("a task transition said something:\n%s", errOut)
	}
}

// An outcome carries no workflow status, and saying so is better than a
// confusing refusal about ladders.
func TestAnOutcomeCannotTransition(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "outcome", "new", "The queue drains", "-w", "WORK-0001")

	code, _, errOut := run(t, app, "transition", "the-queue-drains", "todo")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "outcome") {
		t.Errorf("the refusal did not say what the record is:\n%s", errOut)
	}
}

// A missing kind is a different thing from `idea`: nobody has classified it,
// rather than somebody having classified it as not-yet-classifiable. Warned,
// because the gate criterion is about the problem being understood, not filed.
func TestLeavingThePileWithNoKindWarns(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, errOut := run(t, app, "transition", "WORK-0001", "unprepared")
	if code != ExitOK {
		t.Fatalf("a kindless work item was refused: %s", errOut)
	}
	if !strings.Contains(errOut, "no kind") {
		t.Errorf("a missing kind was not warned about:\n%s", errOut)
	}
}

// spec.md §5.2: an outcome with no verify_by is outcome.unmeasured. Without a
// check no task can be the last one.
func TestLeavingPreparingWithAnUnmeasuredOutcomeWarns(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "change")
	run(t, app, "outcome", "new", "The queue drains", "-w", "WORK-0001")
	run(t, app, "task", "new", "Do the thing", "-w", "WORK-0001")
	run(t, app, "transition", "WORK-0001", "preparing")

	_, _, errOut := run(t, app, "transition", "WORK-0001", "prepared")
	if !strings.Contains(errOut, "cannot be checked") {
		t.Errorf("an unmeasured outcome was not warned about:\n%s", errOut)
	}
	// The shaping warning is separate and must not fire: both exist here.
	if strings.Contains(errOut, "no outcomes") {
		t.Errorf("it complained about outcomes that exist:\n%s", errOut)
	}
}

// Queuing says: you are promising to start something nobody can tell is
// finished. A different message from leaving preparing, at a different moment.
func TestQueuingWithNoOutcomesWarns(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "change")

	code, out, errOut := run(t, app, "transition", "WORK-0001", "todo")
	if code != ExitOK {
		t.Fatalf("queuing was refused: exit %d, %s", code, errOut)
	}
	if !strings.Contains(errOut, "whoever picks it up") {
		t.Errorf("queuing without outcomes was not warned about:\n%s", errOut)
	}
	if strings.Contains(out, "whoever picks it up") {
		t.Errorf("the warning reached stdout:\n%s", out)
	}
}

// A task left open under a closed work item advertises work nobody can pick up.
// Warned, never refused, and never auto-closed --- that would invent a
// disposition nobody chose.
func TestClosingWithOpenTasksWarnsAndDoesNotCloseThem(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha", "--kind", "change")
	run(t, app, "task", "new", "Do the thing", "-w", "WORK-0001")

	code, _, errOut := run(t, app, "work-item", "close", "WORK-0001", "canceled")
	if code != ExitOK {
		t.Fatalf("closing was refused: exit %d, %s", code, errOut)
	}
	if !strings.Contains(errOut, "advertise work nobody can pick up") {
		t.Errorf("an open task was not warned about:\n%s", errOut)
	}

	_, task, _ := run(t, app, "show", "do-the-thing", "--json")
	if strings.Contains(task, `"workflow_status": "closed"`) {
		t.Error("closing the work item closed its task, inventing a disposition")
	}
}
