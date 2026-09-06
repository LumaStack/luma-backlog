package cli

import (
	"strings"
	"testing"
)

// ADR-0005: workflow_status and rank are always written together. A status
// change re-enqueues the record at the back of the destination --- a rank is a
// position in a queue, and leaving the queue does not carry it with you.
func TestAStatusChangeRewritesTheRank(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "work-item", "rank", "WORK-0001", "--bottom")

	_, before, _ := run(t, app, "show", "WORK-0001", "--json")
	if !strings.Contains(before, "010.") {
		t.Fatalf("expected a captured-ordinal rank:\n%s", before)
	}

	run(t, app, "set", "WORK-0001", "workflow_status=todo")
	_, after, _ := run(t, app, "show", "WORK-0001", "--json")
	if !strings.Contains(after, "050.") {
		t.Errorf("the rank prefix did not follow the status:\n%s", after)
	}
}

// Closing is a status change like any other and goes through the same
// operation, rather than setting the field itself.
func TestClosingRewritesTheRank(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")
	run(t, app, "work-item", "rank", "WORK-0001", "--bottom")
	run(t, app, "work-item", "close", "WORK-0001", "--reason", "canceled")

	_, out, _ := run(t, app, "show", "WORK-0001", "--json")
	if !strings.Contains(out, "070.") {
		t.Errorf("closing left the rank at the old status:\n%s", out)
	}
}

// Advancing records in rank order lands them in the same relative order,
// because each arrives behind the last.
func TestAdvancingInOrderPreservesOrder(t *testing.T) {
	app, _ := initialized(t)
	for _, title := range []string{"Alpha", "Bravo", "Charlie"} {
		run(t, app, "work-item", "new", title)
	}
	for _, ref := range []string{"WORK-0001", "WORK-0002", "WORK-0003"} {
		run(t, app, "work-item", "rank", ref, "--bottom")
	}
	for _, ref := range []string{"WORK-0001", "WORK-0002", "WORK-0003"} {
		run(t, app, "set", ref, "workflow_status=todo")
	}
	_, out, _ := run(t, app, "work-item", "list")
	if order(out) != "Alpha Bravo Charlie" {
		t.Errorf("advancing in rank order did not preserve it: %q", order(out))
	}
}

// A record that has never been ranked still gets one when its status changes:
// the two fields are written together, so there is no state where a record has
// a status the tool set and no rank to go with it.
func TestAnUnrankedRecordGetsARankOnAStatusChange(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	run(t, app, "set", "WORK-0001", "workflow_status=todo")
	_, out, _ := run(t, app, "show", "WORK-0001", "--json")
	if !strings.Contains(out, "050.0010.000") {
		t.Errorf("a status change left the record unranked:\n%s", out)
	}
}

// A status the ladder does not carry has no ordinal, so there is nowhere to
// file the record. Refused rather than written with a guessed prefix.
func TestAStatusTheLadderDoesNotCarryIsRefused(t *testing.T) {
	app, _ := initialized(t)
	run(t, app, "work-item", "new", "Alpha")

	code, _, errOut := run(t, app, "set", "WORK-0001", "workflow_status=marinating")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "marinating") {
		t.Errorf("the refusal did not name the status:\n%s", errOut)
	}
}
