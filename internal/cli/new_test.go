package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lumastack/luma-backlog/internal/env"
	"github.com/lumastack/luma-backlog/internal/record"
)

func initialized(t *testing.T) (*App, string) {
	t.Helper()
	app, project := newApp(t)
	if code, _, e := run(t, app, "init"); code != ExitOK {
		t.Fatalf("init failed: %s", e)
	}
	return app, project
}

func readFile(t *testing.T, project, rel string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(project, ".luma", rel))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func fixedAt(t *testing.T, rfc3339 string) env.Clock {
	t.Helper()
	at, err := time.Parse(time.RFC3339, rfc3339)
	if err != nil {
		t.Fatal(err)
	}
	return env.FixedClock{At: at}
}

func readRecord(t *testing.T, project, rel string) *record.Record {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(project, ".luma", rel))
	if err != nil {
		t.Fatal(err)
	}
	r, err := record.Parse(data)
	if err != nil {
		t.Fatalf("%s does not parse: %v", rel, err)
	}
	return r
}

func TestNewWorkItemDerivesEverythingFromTheTitle(t *testing.T) {
	app, project := initialized(t)

	code, out, errOut := run(t, app, "work-item", "new", "Payments v2")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	want := wiPath(t, project, "payments-v2", "index.md")
	if !strings.Contains(out, want) {
		t.Errorf("output did not name the path:\n%s", out)
	}

	r := readRecord(t, project, want)
	if got, _ := r.Get("type"); got != "work-item" {
		t.Errorf("type = %q, want the short form", got)
	}
	if got, _ := r.Get("title"); got != "Payments v2" {
		t.Errorf("title = %q", got)
	}
	// Absent means the first configured value, so it is written explicitly
	// rather than left to be inferred by every reader.
	if got, _ := r.Get("workflow_status"); got != "captured" {
		t.Errorf("workflow_status = %q, want idea", got)
	}
	if !r.Has("created") {
		t.Error("no created stamp")
	}

	// A journal exists at once: somewhere to write has to be there before
	// anyone needs it, or the writing does not happen.
	if _, err := os.Stat(filepath.Join(project, ".luma", wiPath(t, project, "payments-v2", "journal.md"))); err != nil {
		t.Errorf("no journal alongside the work item: %v", err)
	}
}

func TestNewIsIdempotentByName(t *testing.T) {
	app, project := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")

	path := filepath.Join(project, ".luma", wiPath(t, project, "payments-v2", "index.md"))
	if err := os.WriteFile(path, []byte("---\ntype: work item\ntitle: Edited\n---\n\nmine\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out, _ := run(t, app, "work-item", "new", "Payments v2")
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

func TestNewOutcomeTakesItsWorkItemFromAFlag(t *testing.T) {
	app, project := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")

	code, _, errOut := run(t, app, "outcome", "new", "The queue drains", "-w", "payments-v2")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}

	r := readRecord(t, project, wiPath(t, project, "payments-v2", "outcomes", "the-queue-drains.md"))
	if got, _ := r.Get("work_item"); got != "[[work-items/"+wiDir(t, project, "payments-v2")+"]]" {
		t.Errorf("work item = %q", got)
	}
	// desired_state is present and empty rather than absent: an empty field
	// asks to be filled in, a missing one is not noticed.
	if !r.Has("desired_state") {
		t.Error("no desired_state on a new outcome")
	}
}

func TestJudgedUnitsGetNoWorkflowStatus(t *testing.T) {
	// An outcome is judged by evidence; a decision is ratified through
	// stage; an exploration is archived. A declared status on any
	// of them would sit beside the real state and could disagree with it
	// (docs/spec.md §4.4).
	app, project := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")

	for _, tc := range []struct{ unit, path string }{
		{"outcome", wiPath(t, project, "payments-v2", "outcomes", "a-thing.md")},
		{"decision", wiPath(t, project, "payments-v2", "decisions", "ADR-0001-a-thing.md")},
		{"exploration", wiPath(t, project, "payments-v2", "explorations", "a-thing.md")},
	} {
		if code, _, e := run(t, app, tc.unit, "new", "A thing", "-w", "payments-v2"); code != ExitOK {
			t.Fatalf("new %s failed: %s", tc.unit, e)
		}
		if r := readRecord(t, project, tc.path); r.Has("workflow_status") {
			t.Errorf("a new %s carries a workflow_status", tc.unit)
		}
	}

	// A task is worked, so it does carry one.
	run(t, app, "task", "new", "Do it", "-w", "payments-v2")
	if r := readRecord(t, project, wiPath(t, project, "payments-v2", "tasks", "do-it.md")); !r.Has("workflow_status") {
		t.Error("a new task has no workflow_status")
	}
}

func TestNewTakesItsWorkItemFromTheWorkingDirectory(t *testing.T) {
	app, project := initialized(t)
	run(t, app, "work-item", "new", "Payments v2")

	// Working inside a work item should not require naming it.
	app.WorkingDir = filepath.Join(project, ".luma", "backlog/work-items", "payments-v2")

	if code, _, e := run(t, app, "task", "new", "Add the queue"); code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, e)
	}
	r := readRecord(t, project, wiPath(t, project, "payments-v2", "tasks", "add-the-queue.md"))
	if got, _ := r.Get("work_item"); got != "[[work-items/"+wiDir(t, project, "payments-v2")+"]]" {
		t.Errorf("work item was not derived from the working directory: %q", got)
	}
}

func TestNewRefusesAFloatingOutcome(t *testing.T) {
	app, _ := initialized(t)
	code, _, errOut := run(t, app, "outcome", "new", "Floating")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "work-item") {
		t.Errorf("error did not say what was missing:\n%s", errOut)
	}
}

// The record type used to be an argument, so an unknown one was a validation
// error. It is now the command, so an unknown one is an unknown command ---
// still exit 2, still naming what was typed, but refused before any argument
// is parsed.
func TestAnUnknownNounIsAnUnknownCommand(t *testing.T) {
	app, _ := initialized(t)
	code, _, errOut := run(t, app, "sprint", "new", "Something")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "sprint") {
		t.Errorf("error did not name what was typed:\n%s", errOut)
	}
}

// Every noun carries the verbs the tree gives it, and nothing carries a verb
// it has no business with. §9.2 divides these into universal and domain verbs;
// this holds the division the tree actually builds.
func TestEachNounCarriesItsOwnVerbs(t *testing.T) {
	app, _ := initialized(t)
	for _, tc := range []struct {
		noun, verb string
		want       int
	}{
		{"work-item", "new", ExitOK},
		{"outcome", "new", ExitOK},
		{"task", "new", ExitOK},
		{"decision", "new", ExitOK},
		{"exploration", "new", ExitOK},
		{"work-item", "close", ExitOK},
		{"outcome", "verify", ExitOK},
		{"work-item", "journal", ExitOK},
		// Domain verbs do not leak onto nouns they mean nothing for.
		{"task", "close", ExitUsage},
		{"task", "verify", ExitUsage},
		{"decision", "close", ExitUsage},
		{"outcome", "close", ExitUsage},
		{"work-item", "verify", ExitUsage},
	} {
		// Asymmetric on purpose. A verb that exists needs --help, because
		// invoking it bare fails on its missing arguments. A verb that does
		// not exist must be probed *without* --help, since Cobra honors the
		// help flag before it validates arguments --- `task close --help`
		// prints task's help and exits 0 whether or not close is real.
		args := []string{tc.noun, tc.verb}
		if tc.want == ExitOK {
			args = append(args, "--help")
		}
		code, _, errOut := run(t, app, args...)
		if code != tc.want {
			t.Errorf("%v: exit = %d, want %d\n%s", args, code, tc.want, errOut)
		}
	}
}

// --kind classifies a work item and --project states a decision's level.
// Registering them only where they apply makes a wrong one a usage error at
// the door rather than a validation error after parsing.
func TestFlagsAreRegisteredOnlyOnTheNounsThatTakeThem(t *testing.T) {
	app, _ := initialized(t)
	for _, tc := range []struct {
		noun, flag string
		want       int
	}{
		{"work-item", "--kind", ExitOK},
		{"decision", "--project", ExitOK},
		{"task", "--kind", ExitUsage},
		{"outcome", "--kind", ExitUsage},
		{"work-item", "--project", ExitUsage},
		{"task", "--project", ExitUsage},
	} {
		code, _, _ := run(t, app, tc.noun, "new", "--help")
		if code != ExitOK {
			t.Fatalf("%s new --help failed", tc.noun)
		}
		_, out, _ := run(t, app, tc.noun, "new", "--help")
		has := strings.Contains(out, tc.flag)
		if has != (tc.want == ExitOK) {
			t.Errorf("%s new: %s present = %v, want %v", tc.noun, tc.flag, has, tc.want == ExitOK)
		}
	}
}

func TestNewNeedsABacklog(t *testing.T) {
	app, _ := newApp(t) // no init
	code, _, errOut := run(t, app, "work-item", "new", "Thing")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d", code, ExitUsage)
	}
	if !strings.Contains(errOut, "init") {
		t.Errorf("error did not say how to fix it:\n%s", errOut)
	}
}

func TestNewWritesNothingOutsideTheBacklog(t *testing.T) {
	app, project := initialized(t)
	before := snapshot(t, project)

	run(t, app, "work-item", "new", "Payments v2")
	run(t, app, "outcome", "new", "The queue drains", "-w", "payments-v2")

	for path := range snapshot(t, project) {
		if _, existed := before[path]; existed {
			continue
		}
		if !strings.HasPrefix(path, ".luma/") {
			t.Errorf("wrote outside the backlog: %s", path)
		}
	}
}

// wiDir finds a work item's directory by its slug half.
//
// The key leads the directory name and is allocated in creation order, so
// hard-coding WORK-0001 into an assertion would couple every test to the order
// of its fixture. Tests say what they mean — this work item — and let the
// helper find it.
func wiDir(t *testing.T, project, slug string) string {
	t.Helper()
	matches, err := filepath.Glob(filepath.Join(project, ".luma", "backlog", "work-items", "*-"+slug))
	if err != nil || len(matches) != 1 {
		t.Fatalf("expected exactly one work item directory for %q, got %v (%v)", slug, matches, err)
	}
	return filepath.Base(matches[0])
}

// wiPath builds a path inside a work item, found by slug.
func wiPath(t *testing.T, project, slug string, rest ...string) string {
	t.Helper()
	parts := append([]string{"backlog", "work-items", wiDir(t, project, slug)}, rest...)
	return filepath.Join(parts...)
}
