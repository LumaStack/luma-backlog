package cli

import (
	"strings"
	"testing"
)

// withOutcomes builds a work item with two outcomes, neither verified.
func withOutcomes(t *testing.T) (*App, string) {
	t.Helper()
	app, project := withWorkItem(t)
	for _, title := range []string{"The queue drains", "Retries are durable"} {
		if code, _, e := run(t, app, "outcome", "new", title, "-w", "payments-v2"); code != ExitOK {
			t.Fatalf("new outcome failed: %s", e)
		}
	}
	return app, project
}

func TestCompletedIsRefusedWhileAnyOutcomeLacksEvidence(t *testing.T) {
	app, _ := withOutcomes(t)
	run(t, app, "outcome", "verify", "the-queue-drains", "-e", "ran the drain test")

	code, _, errOut := run(t, app, "work-item", "close", "payments-v2", "completed")
	if code != ExitRefused {
		t.Fatalf("exit = %d, want %d (refused)", code, ExitRefused)
	}
	// The refusal must name what is missing, or it is not actionable.
	if !strings.Contains(errOut, "retries-are-durable") {
		t.Errorf("the refusal did not name the unpassing outcome:\n%s", errOut)
	}
	if !strings.Contains(errOut, "1 of 2") {
		t.Errorf("the refusal did not say how far off it was:\n%s", errOut)
	}
}

func TestCompletedSucceedsOnceEveryOutcomeHasEvidence(t *testing.T) {
	app, project := withOutcomes(t)
	run(t, app, "outcome", "verify", "the-queue-drains", "-e", "ran the drain test")
	run(t, app, "outcome", "verify", "retries-are-durable", "-e", "killed the worker mid-flight")

	code, out, errOut := run(t, app, "work-item", "close", "payments-v2", "completed")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	if !strings.Contains(out, "completed") {
		t.Errorf("output did not record the disposition:\n%s", out)
	}

	r := readRecord(t, project, wiPath(t, project, "payments-v2", "index.md"))
	if got, _ := r.Get("workflow_status"); got != "closed" {
		t.Errorf("workflow_status = %q", got)
	}
	if !r.Has("closed") {
		t.Error("no closing record")
	}
}

func TestCancellingIsNeverGated(t *testing.T) {
	// The obvious wrong implementation: gating cancellation on completion
	// would make it impossible to stop work precisely because it was
	// unfinished, which is the only reason anyone ever cancels anything.
	app, project := withOutcomes(t)

	code, _, errOut := run(t, app, "work-item", "close", "payments-v2", "canceled")
	if code != ExitOK {
		t.Fatalf("cancelling unfinished work was refused: exit %d, %s", code, errOut)
	}
	r := readRecord(t, project, wiPath(t, project, "payments-v2", "index.md"))
	if got, _ := r.Get("workflow_status"); got != "closed" {
		t.Errorf("workflow_status = %q", got)
	}
}

func TestSupersededAndRejectedAreAlsoUngated(t *testing.T) {
	for _, reason := range []string{"superseded", "rejected"} {
		app, _ := withOutcomes(t)
		if code, _, e := run(t, app, "work-item", "close", "payments-v2", reason); code != ExitOK {
			t.Errorf("%s was refused: exit %d, %s", reason, code, e)
		}
	}
}

func TestRetiredOutcomesAreExcludedFromTheCount(t *testing.T) {
	// Otherwise retiring an outcome could never let a work item close,
	// which is the point of retiring it.
	app, _ := withOutcomes(t)
	run(t, app, "outcome", "verify", "the-queue-drains", "-e", "ran it")
	run(t, app, "set", "retries-are-durable", "stage=archived")

	code, out, errOut := run(t, app, "work-item", "close", "payments-v2", "completed")
	if code != ExitOK {
		t.Fatalf("a retired outcome still blocked completion: exit %d, %s", code, errOut)
	}
	if !strings.Contains(out, "retired") {
		t.Errorf("the exclusion was not reported:\n%s", out)
	}
}

func TestCompletedIsRefusedWithNoOutcomesAtAll(t *testing.T) {
	// Nothing says it was completed, so the claim has no basis. Vacuous
	// truth is the wrong answer here.
	app, _ := withWorkItem(t)
	code, _, errOut := run(t, app, "work-item", "close", "payments-v2", "completed")
	if code != ExitRefused {
		t.Fatalf("exit = %d, want %d", code, ExitRefused)
	}
	if !strings.Contains(errOut, "no outcomes") {
		t.Errorf("the refusal did not explain itself:\n%s", errOut)
	}
}

func TestCloseRequiresADisposition(t *testing.T) {
	app, _ := withOutcomes(t)
	code, _, errOut := run(t, app, "work-item", "close", "payments-v2")
	if code != ExitUsage {
		t.Fatalf("exit = %d, want %d", code, ExitUsage)
	}
	// Every one of them, because a caller who omitted it does not know which
	// they wanted, and naming one would read as the default.
	for _, want := range []string{"completed", "rejected", "canceled", "superseded"} {
		if !strings.Contains(errOut, want) {
			t.Errorf("the error did not offer %q:\n%s", want, errOut)
		}
	}
}

// `abandoned` was a disposition until ADR-0007 dropped it. A caller reaching
// for it should be told it is unknown, not have it quietly accepted.
func TestAbandonedIsNoLongerADisposition(t *testing.T) {
	app, _ := withOutcomes(t)
	if code, _, _ := run(t, app, "work-item", "close", "payments-v2", "abandoned"); code != ExitUsage {
		t.Error("abandoned was accepted after being dropped")
	}
}

func TestCloseRejectsAnUnknownDisposition(t *testing.T) {
	app, _ := withOutcomes(t)
	if code, _, _ := run(t, app, "work-item", "close", "payments-v2", "done"); code != ExitUsage {
		t.Error("an unknown reason was accepted")
	}
}

func TestVerifyAccumulates(t *testing.T) {
	// Several actors confirming the same outcome is the normal case.
	app, project := withOutcomes(t)
	run(t, app, "outcome", "verify", "the-queue-drains", "-e", "first check")
	run(t, app, "outcome", "verify", "the-queue-drains", "-e", "second check")

	r := readRecord(t, project, wiPath(t, project, "payments-v2", "outcomes", "the-queue-drains.md"))
	var entries []map[string]any
	if err := r.Node("verified").Decode(&entries); err != nil {
		t.Fatalf("verified is not a list: %v", err)
	}
	if len(entries) != 2 {
		t.Errorf("verified has %d entries, want 2", len(entries))
	}
	var evidence []map[string]any
	if err := r.Node("evidence").Decode(&evidence); err != nil {
		t.Fatalf("evidence is not a list: %v", err)
	}
	if len(evidence) != 2 {
		t.Errorf("evidence has %d entries, want 2", len(evidence))
	}
}

func TestVerifyWithoutEvidenceSaysSo(t *testing.T) {
	app, _ := withOutcomes(t)
	code, out, _ := run(t, app, "outcome", "verify", "the-queue-drains")
	if code != ExitOK {
		t.Fatalf("exit = %d", code)
	}
	// Permitted, because refusing would be an opinion the record does not
	// contradict — but not silent, because it is the claim this design
	// distrusts most.
	if !strings.Contains(out, "No evidence recorded") {
		t.Errorf("an unbacked verification passed without comment:\n%s", out)
	}
}

func TestVerifyRefusesANonOutcome(t *testing.T) {
	app, _ := withOutcomes(t)
	if code, _, _ := run(t, app, "outcome", "verify", "payments-v2"); code != ExitUsage {
		t.Error("a work item was accepted for verification")
	}
}

// A doer's claim must never gate a close. Gating on it would gate on the thing
// the design distrusts, and would let a doer clear their own work (ADR-0007).
func TestAssertingSuccessDoesNotLetAWorkItemComplete(t *testing.T) {
	app, _ := withOutcomes(t)
	if code, _, e := run(t, app, "outcome", "assert", "the-queue-drains", "succeeded"); code != ExitOK {
		t.Fatalf("assert failed: %s", e)
	}
	if code, _, _ := run(t, app, "work-item", "close", "payments-v2", "completed"); code != ExitRefused {
		t.Errorf("a doer's own claim cleared the close: exit %d, want %d", code, ExitRefused)
	}
}

// Claims accumulate, so a second attempt is visible. A single overwritten
// field would make "failed once, then succeeded" indistinguishable from
// "succeeded first time".
func TestAssertionsAccumulate(t *testing.T) {
	app, project := withOutcomes(t)
	run(t, app, "outcome", "assert", "the-queue-drains", "failed")
	run(t, app, "outcome", "assert", "the-queue-drains", "succeeded")

	r := readRecord(t, project, wiPath(t, project, "payments-v2", "outcomes", "the-queue-drains.md"))
	var entries []map[string]any
	if err := r.Node("asserted").Decode(&entries); err != nil {
		t.Fatalf("asserted is not a list: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("asserted has %d entries, want 2", len(entries))
	}
	if entries[0]["as"] != "failed" || entries[1]["as"] != "succeeded" {
		t.Errorf("the sequence was not preserved: %v", entries)
	}
}

func TestAssertRefusesAnUnknownClaim(t *testing.T) {
	app, _ := withOutcomes(t)
	if code, _, _ := run(t, app, "outcome", "assert", "the-queue-drains", "done"); code != ExitUsage {
		t.Error("an unknown claim was accepted")
	}
}

func TestAssertRefusesANonOutcome(t *testing.T) {
	app, _ := withOutcomes(t)
	if code, _, _ := run(t, app, "outcome", "assert", "payments-v2", "succeeded"); code != ExitUsage {
		t.Error("a work item was asserted")
	}
}
