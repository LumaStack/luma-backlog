package cli

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestWorkItemsGetSequentialKeys(t *testing.T) {
	// One sequence for the whole backlog. The number is what somebody says out
	// loud or writes in a commit, so it has to mean one record.
	app, project := initialized(t)
	for _, title := range []string{"Payments v2", "Search relevance", "Exports"} {
		if code, _, e := run(t, app, "new", "work-item", title, "--kind", "change"); code != ExitOK {
			t.Fatalf("new %q failed: %s", title, e)
		}
	}
	for path, want := range map[string]string{
		wiPath(t, project, "payments-v2", "index.md"):      "key: WORK-00001",
		wiPath(t, project, "search-relevance", "index.md"): "key: WORK-00002",
		wiPath(t, project, "exports", "index.md"):          "key: WORK-00003",
	} {
		if got := readFile(t, project, path); !strings.Contains(got, want) {
			t.Errorf("%s is missing %q:\n%s", path, want, got)
		}
	}
}

func TestOnlyWorkItemsCarryAKey(t *testing.T) {
	// A key identifies a unit of work. An outcome is a condition and a task is
	// coordination; neither is a thing somebody cites by handle.
	app, project := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	for _, tc := range []struct{ unit, path string }{
		{"outcome", wiPath(t, project, "payments-v2", "outcomes", "a-thing.md")},
		{"task", wiPath(t, project, "payments-v2", "tasks", "a-thing.md")},
	} {
		if code, _, e := run(t, app, "new", tc.unit, "A thing", "-w", "payments-v2"); code != ExitOK {
			t.Fatalf("new %s failed: %s", tc.unit, e)
		}
		if got := readFile(t, project, tc.path); strings.Contains(got, "key:") {
			t.Errorf("a %s carries a key:\n%s", tc.unit, got)
		}
	}
}

func TestAKeyResolvesLikeASlug(t *testing.T) {
	// The point of a handle is that you can use it. Case-insensitively too,
	// because somebody typing one from memory should not have to hold shift.
	app, _ := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	for _, ref := range []string{"WORK-00001", "work-00001"} {
		code, out, errOut := run(t, app, "show", ref, "--json")
		if code != ExitOK {
			t.Fatalf("show %s failed: %s", ref, errOut)
		}
		var rec map[string]any
		if err := json.Unmarshal([]byte(out), &rec); err != nil {
			t.Fatalf("show %s did not emit JSON: %v", ref, err)
		}
		// The slug is the directory name, which now leads with the key.
		if rec["slug"] != "WORK-00001-payments-v2" {
			t.Errorf("show %s resolved to %v", ref, rec["slug"])
		}
		if rec["key"] != "WORK-00001" {
			t.Errorf("show %s emitted key %v", ref, rec["key"])
		}
	}
}

func TestAskingTwiceDoesNotBurnAKey(t *testing.T) {
	// A work item's path is still its slug, so a retry returns the existing
	// record before a number is allocated. This is the trap the decision
	// numbering had to be rescued from, and it does not exist here — the test
	// is what stops it being reintroduced.
	app, project := initialized(t)
	for i := 0; i < 3; i++ {
		if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
			t.Fatalf("attempt %d failed: %s", i, e)
		}
	}
	if code, _, e := run(t, app, "new", "work-item", "Search relevance", "--kind", "change"); code != ExitOK {
		t.Fatalf("second item failed: %s", e)
	}
	if got := readFile(t, project, wiPath(t, project, "search-relevance", "index.md")); !strings.Contains(got, "key: WORK-00002") {
		t.Errorf("repeated asks burned a number:\n%s", got)
	}
}

func TestTheJoinedFormResolves(t *testing.T) {
	// WORK-00002-lint-the-corpus is how a work item is written and said, the
	// way a decision's filename joins its number and slug. All three forms
	// have to reach the same record, or the one people actually type is the
	// one that fails.
	app, _ := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Lint the corpus", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	for _, ref := range []string{
		"WORK-00001-lint-the-corpus",
		"work-00001-lint-the-corpus",
		"WORK-00001",
		"lint-the-corpus",
	} {
		code, out, errOut := run(t, app, "show", ref, "--json")
		if code != ExitOK {
			t.Fatalf("show %s failed: %s", ref, errOut)
		}
		var rec map[string]any
		if err := json.Unmarshal([]byte(out), &rec); err != nil {
			t.Fatalf("show %s did not emit JSON: %v", ref, err)
		}
		if rec["ref"] != "WORK-00001-lint-the-corpus" {
			t.Errorf("show %s emitted ref %v", ref, rec["ref"])
		}
	}
}

func TestARecordWithoutAKeyRefsAsItsSlug(t *testing.T) {
	// Only a work item carries a key, so an outcome's reference is its slug
	// and the identifier column is never empty.
	app, _ := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	if code, _, e := run(t, app, "new", "outcome", "It drains", "-w", "payments-v2"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	_, out, _ := run(t, app, "show", "it-drains", "--json")
	var rec map[string]any
	if err := json.Unmarshal([]byte(out), &rec); err != nil {
		t.Fatal(err)
	}
	if rec["ref"] != "it-drains" {
		t.Errorf("an outcome's ref = %v, want its slug", rec["ref"])
	}
	if _, hasKey := rec["key"]; hasKey {
		t.Errorf("an outcome carries a key: %v", rec["key"])
	}
}
