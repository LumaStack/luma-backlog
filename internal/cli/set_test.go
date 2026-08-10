package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func withDeliverable(t *testing.T) (*App, string) {
	t.Helper()
	app, project := initialized(t)
	if code, _, e := run(t, app, "new", "deliverable", "Payments v2"); code != ExitOK {
		t.Fatalf("new failed: %s", e)
	}
	return app, project
}

func TestSetChangesOnlyWhatWasNamed(t *testing.T) {
	app, project := withDeliverable(t)
	path := filepath.Join(project, ".backlog", "deliverables/payments-v2/index.md")

	// A key from another system, which this tool knows nothing about.
	data, _ := os.ReadFile(path)
	edited := strings.Replace(string(data), "lifecycle_status:", "some_other_tool: {cursor: abc123}\nlifecycle_status:", 1)
	if err := os.WriteFile(path, []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}

	if code, _, e := run(t, app, "set", "payments-v2", "workflow_status=in_progress"); code != ExitOK {
		t.Fatalf("set failed: %s", e)
	}

	r := readRecord(t, project, "deliverables/payments-v2/index.md")
	if got, _ := r.Get("workflow_status"); got != "in_progress" {
		t.Errorf("workflow_status = %q", got)
	}
	// Losing another system's state is silent data loss, not a reported bug.
	if !r.Has("some_other_tool") {
		t.Error("set dropped a key it did not recognize")
	}
	if got, _ := r.Get("title"); got != "Payments v2" {
		t.Errorf("set disturbed the title: %q", got)
	}
}

func TestSetStampsModified(t *testing.T) {
	app, project := withDeliverable(t)
	run(t, app, "set", "payments-v2", "workflow_status=todo")

	r := readRecord(t, project, "deliverables/payments-v2/index.md")
	if !r.Has("modified") {
		t.Error("no modified stamp after an edit")
	}
}

func TestSetDoesNotOverruleAnExplicitModified(t *testing.T) {
	// The tool should not overrule something it was just told.
	app, project := withDeliverable(t)
	run(t, app, "set", "payments-v2", "modified=yesterday")

	r := readRecord(t, project, "deliverables/payments-v2/index.md")
	if got, _ := r.Get("modified"); got != "yesterday" {
		t.Errorf("modified = %q, want the value that was passed", got)
	}
}

func TestSetRawWritesStructure(t *testing.T) {
	app, project := withDeliverable(t)
	if code, _, e := run(t, app, "set", "payments-v2", `blocked:=[{on: 2026-08-09, why: vendor}]`); code != ExitOK {
		t.Fatalf("set failed: %s", e)
	}
	r := readRecord(t, project, "deliverables/payments-v2/index.md")
	node := r.Node("blocked")
	if node == nil {
		t.Fatal("blocked absent")
	}
	var got []map[string]string
	if err := node.Decode(&got); err != nil {
		t.Fatalf("blocked is not a list of maps: %v", err)
	}
	if len(got) != 1 || got[0]["why"] != "vendor" {
		t.Errorf("blocked = %v", got)
	}
}

func TestSetKeepsWikilinksIntact(t *testing.T) {
	// A wikilink looks exactly like a YAML nested sequence. Guessing at the
	// value's shape would turn [[deliverables/x]] into [["deliverables/x"]],
	// which is why the raw form is opt-in rather than detected.
	app, project := withDeliverable(t)
	run(t, app, "set", "payments-v2", "supersedes=[[deliverables/payments-v1]]")

	r := readRecord(t, project, "deliverables/payments-v2/index.md")
	got, ok := r.Get("supersedes")
	if !ok {
		t.Fatal("supersedes absent, or not a scalar — it was parsed as a list")
	}
	if got != "[[deliverables/payments-v1]]" {
		t.Errorf("supersedes = %q, want the wikilink unchanged", got)
	}
}

func TestSetUnsetRemovesAField(t *testing.T) {
	app, project := withDeliverable(t)
	if code, _, e := run(t, app, "set", "payments-v2", "--unset", "lifecycle_status"); code != ExitOK {
		t.Fatalf("set failed: %s", e)
	}
	if readRecord(t, project, "deliverables/payments-v2/index.md").Has("lifecycle_status") {
		t.Error("--unset left the field in place")
	}
}

func TestSetRefusesAStaleWrite(t *testing.T) {
	app, project := withDeliverable(t)

	// Read, as a caller would.
	_, out, _ := run(t, app, "show", "payments-v2", "--json")
	var seen struct {
		Hash string `json:"hash"`
	}
	if err := json.Unmarshal([]byte(out), &seen); err != nil {
		t.Fatal(err)
	}
	if seen.Hash == "" {
		t.Fatal("show --json emitted no hash")
	}

	// Somebody else writes in between.
	path := filepath.Join(project, ".backlog", "deliverables/payments-v2/index.md")
	data, _ := os.ReadFile(path)
	os.WriteFile(path, append(data, []byte("\nsomeone else was here\n")...), 0o644)

	code, _, errOut := run(t, app, "set", "payments-v2", "workflow_status=todo", "--if-unchanged", seen.Hash)
	if code != ExitConflict {
		t.Fatalf("exit = %d, want %d (conflict)", code, ExitConflict)
	}
	if !strings.Contains(errOut, "re-read and retry") {
		t.Errorf("error did not say what to do:\n%s", errOut)
	}

	// And the other writer's change must survive the refusal.
	after, _ := os.ReadFile(path)
	if !strings.Contains(string(after), "someone else was here") {
		t.Error("a refused write still modified the record")
	}
}

func TestSetAcceptsAMatchingHash(t *testing.T) {
	app, _ := withDeliverable(t)
	_, out, _ := run(t, app, "show", "payments-v2", "--json")
	var seen struct {
		Hash string `json:"hash"`
	}
	json.Unmarshal([]byte(out), &seen)

	if code, _, e := run(t, app, "set", "payments-v2", "workflow_status=todo", "--if-unchanged", seen.Hash); code != ExitOK {
		t.Fatalf("a matching hash was refused: exit %d, %s", code, e)
	}
}

func TestSetRejectsMalformedArguments(t *testing.T) {
	app, _ := withDeliverable(t)
	for _, args := range [][]string{
		{"set", "payments-v2"},
		{"set", "payments-v2", "novalue"},
		{"set", "payments-v2", "=orphan"},
	} {
		if code, _, _ := run(t, app, args...); code != ExitUsage {
			t.Errorf("%v exited %d, want %d", args, code, ExitUsage)
		}
	}
}

func TestSetReportsNotFound(t *testing.T) {
	app, _ := withDeliverable(t)
	if code, _, _ := run(t, app, "set", "nothing-like-this", "a=b"); code != ExitNotFound {
		t.Errorf("exit code was not %d", ExitNotFound)
	}
}
