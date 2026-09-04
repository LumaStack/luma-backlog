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

	if code, _, e := run(t, app, "new", "work-item", "A crash", "--kind", "defect"); code != ExitOK {
		t.Fatalf("new --kind failed: %s", e)
	}
	if got := readFile(t, project, "backlog/work-items/a-crash/index.md"); !strings.Contains(got, "kind: defect") {
		t.Errorf("--kind defect was not recorded:\n%s", got)
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
		{"new", "work-item", "A crash", "--kind", "defect"},
		{"new", "work-item", "Please add exports", "--kind", "request"},
		{"new", "work-item", "Ordinary work", "--kind", "change"},
	} {
		if code, _, e := run(t, app, args...); code != ExitOK {
			t.Fatalf("%v failed: %s", args, e)
		}
	}

	_, out, _ := run(t, app, "list", "--kind", "defect")
	if !strings.Contains(out, "a-crash") {
		t.Errorf("the defect was not listed:\n%s", out)
	}
	for _, unwanted := range []string{"please-add-exports", "ordinary-work"} {
		if strings.Contains(out, unwanted) {
			t.Errorf("--kind bug returned %s:\n%s", unwanted, out)
		}
	}
}

func TestAliasesAreAcceptedAndCanonicalIsStored(t *testing.T) {
	// Canonical names are emitted, aliases are accepted (spec.md §9.1). The
	// point is that an external tracker emitting "bug" lands as a defect with
	// nobody mapping anything, and the file on disk says one thing regardless
	// of which word the caller reached for.
	app, project := initialized(t)
	for _, args := range [][]string{
		{"new", "work-item", "A crash", "--kind", "bug"},
		{"new", "work-item", "Please add exports", "--kind", "ask"},
	} {
		if code, _, e := run(t, app, args...); code != ExitOK {
			t.Fatalf("%v failed: %s", args, e)
		}
	}
	for path, want := range map[string]string{
		"backlog/work-items/a-crash/index.md":            "kind: defect",
		"backlog/work-items/please-add-exports/index.md": "kind: request",
	} {
		if got := readFile(t, project, path); !strings.Contains(got, want) {
			t.Errorf("%s did not store %q:\n%s", path, want, got)
		}
	}
}

func TestFilteringByAliasFindsCanonicalRecords(t *testing.T) {
	// Somebody who types the familiar word must find the record, or the alias
	// is only half implemented and the second half is the confusing half.
	app, _ := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "A crash", "--kind", "defect"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	_, out, _ := run(t, app, "list", "--kind", "bug")
	if !strings.Contains(out, "a-crash") {
		t.Errorf("--kind bug did not find a record stored as defect:\n%s", out)
	}
}

func TestAnUnknownKindIsKeptAsWritten(t *testing.T) {
	// Values are opaque and somebody else's vocabulary is not an error.
	app, project := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Something", "--kind", "chore"); code != ExitOK {
		t.Fatalf("new failed: %s", e)
	}
	if got := readFile(t, project, "backlog/work-items/something/index.md"); !strings.Contains(got, "kind: chore") {
		t.Errorf("an unrecognized kind was not kept:\n%s", got)
	}
}
