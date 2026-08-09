package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lumastack/luma-backlog/internal/record"
)

func initialised(t *testing.T) (*App, string) {
	t.Helper()
	app, project := newApp(t)
	if code, _, e := run(t, app, "init"); code != ExitOK {
		t.Fatalf("init failed: %s", e)
	}
	return app, project
}

func readRecord(t *testing.T, project, rel string) *record.Record {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(project, ".backlog", rel))
	if err != nil {
		t.Fatal(err)
	}
	r, err := record.Parse(data)
	if err != nil {
		t.Fatalf("%s does not parse: %v", rel, err)
	}
	return r
}

func TestNewDeliverableDerivesEverythingFromTheTitle(t *testing.T) {
	app, project := initialised(t)

	code, out, errOut := run(t, app, "new", "deliverable", "Payments v2")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	const want = "deliverables/payments-v2/index.md"
	if !strings.Contains(out, want) {
		t.Errorf("output did not name the path:\n%s", out)
	}

	r := readRecord(t, project, want)
	if got, _ := r.Get("type"); got != "deliverable" {
		t.Errorf("type = %q, want the short form", got)
	}
	if got, _ := r.Get("title"); got != "Payments v2" {
		t.Errorf("title = %q", got)
	}
	// Absent means the first configured value, so it is written explicitly
	// rather than left to be inferred by every reader.
	if got, _ := r.Get("workflow_status"); got != "idea" {
		t.Errorf("workflow_status = %q, want idea", got)
	}
	if !r.Has("created") {
		t.Error("no created stamp")
	}

	// A journal exists at once: somewhere to write has to be there before
	// anyone needs it, or the writing does not happen.
	if _, err := os.Stat(filepath.Join(project, ".backlog", "deliverables/payments-v2/journal.md")); err != nil {
		t.Errorf("no journal alongside the deliverable: %v", err)
	}
}

func TestNewIsIdempotentByName(t *testing.T) {
	app, project := initialised(t)
	run(t, app, "new", "deliverable", "Payments v2")

	path := filepath.Join(project, ".backlog", "deliverables/payments-v2/index.md")
	if err := os.WriteFile(path, []byte("---\ntype: deliverable\ntitle: Edited\n---\n\nmine\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out, _ := run(t, app, "new", "deliverable", "Payments v2")
	if code != ExitOK {
		t.Fatalf("second create failed with %d", code)
	}
	if !strings.Contains(out, "exists") {
		t.Errorf("second create did not report the record as existing:\n%s", out)
	}
	// An agent retrying after a dropped connection must not destroy the work
	// of its first attempt.
	after, _ := os.ReadFile(path)
	if !strings.Contains(string(after), "mine") {
		t.Error("a second create overwrote an edited record")
	}
}

func TestNewOutcomeTakesItsDeliverableFromAFlag(t *testing.T) {
	app, project := initialised(t)
	run(t, app, "new", "deliverable", "Payments v2")

	code, _, errOut := run(t, app, "new", "outcome", "The queue drains", "-d", "payments-v2")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}

	r := readRecord(t, project, "deliverables/payments-v2/outcomes/the-queue-drains.md")
	if got, _ := r.Get("deliverable"); got != "[[deliverables/payments-v2]]" {
		t.Errorf("deliverable = %q", got)
	}
	// desired_state is present and empty rather than absent: an empty field
	// asks to be filled in, a missing one is not noticed.
	if !r.Has("desired_state") {
		t.Error("no desired_state on a new outcome")
	}
}

func TestNewTakesItsDeliverableFromTheWorkingDirectory(t *testing.T) {
	app, project := initialised(t)
	run(t, app, "new", "deliverable", "Payments v2")

	// Working inside a deliverable should not require naming it.
	app.WorkingDir = filepath.Join(project, ".backlog", "deliverables", "payments-v2")

	if code, _, e := run(t, app, "new", "task", "Add the queue"); code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, e)
	}
	r := readRecord(t, project, "deliverables/payments-v2/tasks/add-the-queue.md")
	if got, _ := r.Get("deliverable"); got != "[[deliverables/payments-v2]]" {
		t.Errorf("deliverable was not derived from the working directory: %q", got)
	}
}

func TestNewRefusesAFloatingOutcome(t *testing.T) {
	app, _ := initialised(t)
	code, _, errOut := run(t, app, "new", "outcome", "Floating")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "deliverable") {
		t.Errorf("error did not say what was missing:\n%s", errOut)
	}
}

func TestNewRejectsAnUnknownUnit(t *testing.T) {
	app, _ := initialised(t)
	code, _, errOut := run(t, app, "new", "sprint", "Something")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "sprint") {
		t.Errorf("error did not name the unit given:\n%s", errOut)
	}
}

func TestNewNeedsABacklog(t *testing.T) {
	app, _ := newApp(t) // no init
	code, _, errOut := run(t, app, "new", "deliverable", "Thing")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "init") {
		t.Errorf("error did not say how to fix it:\n%s", errOut)
	}
}

func TestNewWritesNothingOutsideTheBacklog(t *testing.T) {
	app, project := initialised(t)
	before := snapshot(t, project)

	run(t, app, "new", "deliverable", "Payments v2")
	run(t, app, "new", "outcome", "The queue drains", "-d", "payments-v2")

	for path := range snapshot(t, project) {
		if _, existed := before[path]; existed {
			continue
		}
		if !strings.HasPrefix(path, ".backlog/") {
			t.Errorf("wrote outside the backlog: %s", path)
		}
	}
}
