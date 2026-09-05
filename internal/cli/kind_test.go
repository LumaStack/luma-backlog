package cli

import (
	"os"
	"path/filepath"
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
	if got := readFile(t, project, wiPath(t, project, "ordinary-work", "index.md")); strings.Contains(got, "kind:") {
		t.Errorf("a kind was written when none was given:\n%s", got)
	}

	if code, _, e := run(t, app, "new", "work-item", "A crash", "--kind", "defect"); code != ExitOK {
		t.Fatalf("new --kind failed: %s", e)
	}
	if got := readFile(t, project, wiPath(t, project, "a-crash", "index.md")); !strings.Contains(got, "kind: defect") {
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
		wiPath(t, project, "a-crash", "index.md"):            "kind: defect",
		wiPath(t, project, "please-add-exports", "index.md"): "kind: request",
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
	if got := readFile(t, project, wiPath(t, project, "something", "index.md")); !strings.Contains(got, "kind: chore") {
		t.Errorf("an unrecognized kind was not kept:\n%s", got)
	}
}

func TestBlankKindIsNudgedNotRefused(t *testing.T) {
	// Avoided, not denied. Capture has to stay free — an issue arriving from
	// elsewhere genuinely has no kind until somebody looks — but blank left
	// frictionless becomes the path of least resistance, and then every idea
	// arrives blank and the field means nothing.
	app, _ := initialized(t)
	code, out, errOut := run(t, app, "new", "work-item", "Something arrived")
	if code != ExitOK {
		t.Fatalf("capture was refused: exit %d, %s", code, errOut)
	}
	if !strings.Contains(out, "created") {
		t.Errorf("the record was not created:\n%s", out)
	}
	if !strings.Contains(errOut, "no kind") {
		t.Errorf("nothing was said about the missing kind:\n%q", errOut)
	}
	for _, k := range []string{"defect", "request", "idea", "change"} {
		if !strings.Contains(errOut, k) {
			t.Errorf("the nudge did not name %q:\n%s", k, errOut)
		}
	}
}

func TestNoNudgeWhenAKindIsGiven(t *testing.T) {
	// A warning that fires on correct use is one people learn to ignore.
	app, _ := initialized(t)
	_, _, errOut := run(t, app, "new", "work-item", "A crash", "--kind", "defect")
	if strings.Contains(errOut, "no kind") {
		t.Errorf("a classified record was nudged anyway:\n%q", errOut)
	}
}

func TestOtherUnitsAreNotNudged(t *testing.T) {
	// Only a work item takes a kind, so only a work item can be missing one.
	app, _ := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	_, _, errOut := run(t, app, "new", "outcome", "It drains", "-w", "payments-v2")
	if strings.Contains(errOut, "no kind") {
		t.Errorf("an outcome was nudged about kind:\n%q", errOut)
	}
}

func TestDecisionsAreNumberedFromOneSequence(t *testing.T) {
	// One sequence for the whole project, not one per directory. A decision
	// inside a work item and one in the records tier must not both be
	// ADR-0001, because the number is what somebody cites in a commit message
	// or says out loud, and it has to mean one record.
	app, project := initialized(t)
	if code, _, e := run(t, app, "new", "work-item", "Payments v2", "--kind", "change"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	for _, args := range [][]string{
		{"new", "decision", "Catalogs do not inherit", "--project"},
		{"new", "decision", "Retry inside the worker", "-w", "payments-v2"},
		{"new", "decision", "Store evidence as events", "--project"},
	} {
		if code, _, e := run(t, app, args...); code != ExitOK {
			t.Fatalf("%v failed: %s", args, e)
		}
	}
	for _, want := range []string{
		"records/decisions/ADR-0001-catalogs-do-not-inherit.md",
		wiPath(t, project, "payments-v2", "decisions", "ADR-0002-retry-inside-the-worker.md"),
		"records/decisions/ADR-0003-store-evidence-as-events.md",
	} {
		if _, err := os.Stat(filepath.Join(project, ".luma", want)); err != nil {
			t.Errorf("expected %s: %v", want, err)
		}
	}
}

func TestAskingTwiceForADecisionDoesNotBurnANumber(t *testing.T) {
	// Idempotent by name, and the number must not defeat that. The filename
	// is no longer derived from the title alone, so the existence check
	// cannot be a path lookup — asking twice has to find the first record
	// rather than allocate a second number for the same title.
	app, _ := initialized(t)
	if code, _, e := run(t, app, "new", "decision", "Catalogs do not inherit", "--project"); code != ExitOK {
		t.Fatalf("first create failed: %s", e)
	}
	code, out, _ := run(t, app, "new", "decision", "Catalogs do not inherit", "--project")
	if code != ExitOK {
		t.Fatalf("second create exited %d", code)
	}
	if !strings.Contains(out, "exists") || !strings.Contains(out, "ADR-0001") {
		t.Errorf("the second ask did not find the first record:\n%s", out)
	}
	if code, _, _ := run(t, app, "new", "decision", "Something else", "--project"); code != ExitOK {
		t.Fatal("third create failed")
	}
	_, out, _ = run(t, app, "list", "decision")
	if strings.Contains(out, "ADR-0003") {
		t.Errorf("a number was burned by the repeated ask:\n%s", out)
	}
}

func TestANewDecisionCarriesTheContractsFields(t *testing.T) {
	// decided and reopen_trigger are present and empty rather than absent: an
	// empty field asks to be filled in, a missing one is not noticed. decided
	// is NOT stamped with today, because it records when the position became
	// binding, which is not when the file appeared.
	app, project := initialized(t)
	if code, _, e := run(t, app, "new", "decision", "Catalogs do not inherit", "--project"); code != ExitOK {
		t.Fatalf("new failed: %s", e)
	}
	got := readFile(t, project, "records/decisions/ADR-0001-catalogs-do-not-inherit.md")
	for _, want := range []string{
		`decided: ""`,
		`reopen_trigger: ""`,
		"# ADR-0001: Catalogs do not inherit",
		"## Summary", "## Problem", "## Decision", "## Why",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("a new decision is missing %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, "## Context") || strings.Contains(got, "## What was chosen") {
		t.Errorf("the pre-bundle section shape survived:\n%s", got)
	}
}

func TestInquiryInstancesAreStoredAsInquiry(t *testing.T) {
	// review, audit, investigation and spike are INSTANCES of an inquiry
	// rather than synonyms for it, so the alias loses a shade of meaning that
	// bug against defect does not. Accepted, because without it some records
	// say spike and some say inquiry, and the filter that earns the field
	// stops finding half of them.
	app, project := initialized(t)
	for _, tc := range []struct{ typed, slug string }{
		{"review", "read-the-code"},
		{"audit", "check-the-books"},
		{"investigation", "what-broke-on-tuesday"},
		{"spike", "can-we-use-parquet"},
	} {
		if code, _, e := run(t, app, "new", "work-item", strings.ReplaceAll(tc.slug, "-", " "), "--kind", tc.typed); code != ExitOK {
			t.Fatalf("--kind %s failed: %s", tc.typed, e)
		}
		got := readFile(t, project, wiPath(t, project, tc.slug, "index.md"))
		if !strings.Contains(got, "kind: inquiry") {
			t.Errorf("--kind %s was not stored as inquiry:\n%s", tc.typed, got)
		}
	}

	// And all four are findable as one, which is the point of aliasing them.
	_, out, _ := run(t, app, "list", "--kind", "inquiry")
	for _, slug := range []string{"read-the-code", "check-the-books", "what-broke-on-tuesday", "can-we-use-parquet"} {
		if !strings.Contains(out, slug) {
			t.Errorf("%s was not listed as an inquiry:\n%s", slug, out)
		}
	}
}
