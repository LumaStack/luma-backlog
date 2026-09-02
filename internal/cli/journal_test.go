package cli

import (
	"strings"
	"testing"
)

func TestJournalCapturesInOneInvocation(t *testing.T) {
	app, project := withDeliverable(t)

	code, out, errOut := run(t, app, "journal", "--", "--use-hold pins the source snapshot")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	if !strings.Contains(out, "journal.md") {
		t.Errorf("output did not name the file:\n%s", out)
	}

	// The outcome this command exists for: one command, and the line is there.
	body := readFile(t, project, "backlog/deliverables/payments-v2/journal.md")
	if !strings.Contains(body, "--use-hold pins the source snapshot") {
		t.Errorf("the line was not written:\n%s", body)
	}
	if !strings.Contains(body, "2026-08-09") {
		t.Errorf("today's entry was not opened:\n%s", body)
	}
}

func TestJournalKeepsAddingToTheSameDay(t *testing.T) {
	app, project := withDeliverable(t)
	run(t, app, "journal", "First.")
	run(t, app, "journal", "Second.")

	body := readFile(t, project, "backlog/deliverables/payments-v2/journal.md")
	if strings.Count(body, "2026-08-09") != 1 {
		t.Errorf("a second entry was opened for the same day:\n%s", body)
	}
	if strings.Index(body, "First.") > strings.Index(body, "Second.") {
		t.Errorf("lines are out of order within the day:\n%s", body)
	}
}

func TestJournalWithNoArgumentReadsIt(t *testing.T) {
	app, _ := withDeliverable(t)
	run(t, app, "journal", "Something worth keeping.")

	code, out, _ := run(t, app, "journal")
	if code != ExitOK {
		t.Fatalf("exit = %d", code)
	}
	if !strings.Contains(out, "Something worth keeping.") {
		t.Errorf("reading did not show the entry:\n%s", out)
	}
}

func TestJournalDerivesTheDeliverableFromOneBeingTheOnlyOne(t *testing.T) {
	// Capture has to cost one invocation. Requiring -d every time is the
	// friction that stops it happening at all.
	app, _ := withDeliverable(t)
	if code, _, e := run(t, app, "journal", "No flag needed."); code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, e)
	}
}

func TestJournalNamesTheCandidatesWhenAmbiguous(t *testing.T) {
	app, _ := withDeliverable(t)
	run(t, app, "new", "deliverable", "Search relevance")

	code, _, errOut := run(t, app, "journal", "Which one?")
	if code != ExitUsage {
		t.Fatalf("exit = %d, want %d", code, ExitUsage)
	}
	for _, want := range []string{"payments-v2", "search-relevance"} {
		if !strings.Contains(errOut, want) {
			t.Errorf("error did not list %s:\n%s", want, errOut)
		}
	}
}

func TestJournalPrefersTheWorkingDirectory(t *testing.T) {
	app, project := withDeliverable(t)
	run(t, app, "new", "deliverable", "Search relevance")

	// Two deliverables, so it would be ambiguous — except we are inside one.
	app.WorkingDir = project + "/.luma/deliverables/search-relevance"
	if code, _, e := run(t, app, "journal", "Context wins."); code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, e)
	}
	body := readFile(t, project, "backlog/deliverables/search-relevance/journal.md")
	if !strings.Contains(body, "Context wins.") {
		t.Errorf("written to the wrong journal:\n%s", body)
	}
}

func TestJournalNeverRewritesWhatIsBelow(t *testing.T) {
	app, project := withDeliverable(t)
	run(t, app, "journal", "Day one.")

	before := readFile(t, project, "backlog/deliverables/payments-v2/journal.md")
	app.Env.Clock = fixedAt(t, "2026-08-10T09:00:00Z")
	run(t, app, "journal", "Day two.")

	after := readFile(t, project, "backlog/deliverables/payments-v2/journal.md")
	// Newest first, and the earlier day survives byte for byte.
	if strings.Index(after, "2026-08-10") > strings.Index(after, "2026-08-09") {
		t.Errorf("the newer entry is not first:\n%s", after)
	}
	tail := before[strings.Index(before, "## "):]
	if !strings.Contains(after, strings.TrimRight(tail, "\n")) {
		t.Errorf("the earlier entry was rewritten:\n--- before ---\n%s\n--- after ---\n%s", before, after)
	}
}

func TestJournalSuggestsTheEscapeForTextThatLooksLikeAFlag(t *testing.T) {
	// Text beginning with -- is common in a journal, and being told it is an
	// unknown flag helps nobody unless the fix comes with it.
	app, _ := withDeliverable(t)
	code, _, errOut := run(t, app, "journal", "--use-hold pins the source snapshot")
	if code != ExitUsage {
		t.Fatalf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "put -- in front") {
		t.Errorf("the error did not offer the fix:\n%s", errOut)
	}
}

func TestJournalRejectsEmptyText(t *testing.T) {
	app, _ := withDeliverable(t)
	if code, _, _ := run(t, app, "journal", "   "); code != ExitUsage {
		t.Error("empty text was accepted")
	}
}
