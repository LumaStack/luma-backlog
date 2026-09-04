package cli

import (
	"strings"
	"testing"
)

func TestKindIsWrittenOnlyWhenGiven(t *testing.T) {
	// Absence is the common case and means ordinary work — something we
	// decided to do, obliging nobody. Defaulting the field would make every
	// record claim a classification nobody chose, which is the mistake the
	// old `idea` default made one field over.
	app, project := initialized(t)

	if code, _, e := run(t, app, "new", "work-item", "Ordinary work"); code != ExitOK {
		t.Fatalf("new failed: %s", e)
	}
	if got := readFile(t, project, "backlog/work-items/ordinary-work/index.md"); strings.Contains(got, "kind:") {
		t.Errorf("a kind was written when none was given:\n%s", got)
	}

	if code, _, e := run(t, app, "new", "work-item", "A crash", "--kind", "bug"); code != ExitOK {
		t.Fatalf("new --kind failed: %s", e)
	}
	if got := readFile(t, project, "backlog/work-items/a-crash/index.md"); !strings.Contains(got, "kind: bug") {
		t.Errorf("--kind bug was not recorded:\n%s", got)
	}
}

func TestKindClassifiesWorkItemsOnly(t *testing.T) {
	// An outcome is a desired state and a task is coordination; neither is a
	// thing somebody requested or a thing that broke. Accepting the flag
	// there would store a field nothing can read.
	app, _ := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	code, _, errOut := run(t, app, "new", "outcome", "It drains", "-w", "payments-v2", "--kind", "bug")
	if code == ExitOK {
		t.Fatal("--kind was accepted on an outcome")
	}
	if !strings.Contains(errOut, "kind") {
		t.Errorf("the refusal did not name the flag: %q", errOut)
	}
}

func TestListFiltersByKind(t *testing.T) {
	// Filtering is what earns the field. A kind nothing can select on is a
	// label, and labels attract categories that fit language rather than use.
	app, _ := initialized(t)
	for _, args := range [][]string{
		{"new", "work-item", "A crash", "--kind", "bug"},
		{"new", "work-item", "Please add exports", "--kind", "request"},
		{"new", "work-item", "Ordinary work"},
	} {
		if code, _, e := run(t, app, args...); code != ExitOK {
			t.Fatalf("%v failed: %s", args, e)
		}
	}

	_, out, _ := run(t, app, "list", "--kind", "bug")
	if !strings.Contains(out, "a-crash") {
		t.Errorf("the bug was not listed:\n%s", out)
	}
	for _, unwanted := range []string{"please-add-exports", "ordinary-work"} {
		if strings.Contains(out, unwanted) {
			t.Errorf("--kind bug returned %s:\n%s", unwanted, out)
		}
	}
}
