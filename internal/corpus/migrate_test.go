package corpus

import (
	"strings"
	"testing"
	"time"
)

func TestPlanMovesOnlyThePrefix(t *testing.T) {
	b := migratedBacklog(t, map[string][]string{
		"WORK-0031-reshape-the-command-surface": {"WORK-0031"},
		"WORK-0102-a-second-one":                {"WORK-0102"},
	})
	plan, err := PlanKeyMigration(b, "BACK")
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Renames) != 2 || len(plan.Collisions) != 0 {
		t.Fatalf("renames=%d collisions=%d, want 2 and 0", len(plan.Renames), len(plan.Collisions))
	}
	for _, r := range plan.Renames {
		if r.OldKey == "WORK-0031" {
			if r.NewKey != "BACK-0031" {
				t.Errorf("NewKey = %s, want BACK-0031 — only the prefix moves", r.NewKey)
			}
			if r.NewName != "BACK-0031-reshape-the-command-surface" {
				t.Errorf("NewName = %s, want the slug kept", r.NewName)
			}
		}
	}
}

func TestPlanLeavesRecordsAlreadyAtTheTarget(t *testing.T) {
	// A mixed corpus, which is the state this project is actually in.
	b := migratedBacklog(t, map[string][]string{
		"WORK-0102-still-old": {"WORK-0102"},
		"BACK-0103-already":   {"BACK-0103"},
	})
	plan, err := PlanKeyMigration(b, "BACK")
	if err != nil {
		t.Fatal(err)
	}
	if plan.AlreadyCorrect != 1 {
		t.Errorf("AlreadyCorrect = %d, want 1", plan.AlreadyCorrect)
	}
	if len(plan.Renames) != 1 || plan.Renames[0].OldKey != "WORK-0102" {
		t.Errorf("renames = %+v, want only WORK-0102", plan.Renames)
	}
}

func TestPlanReportsACollisionAgainstALiveKey(t *testing.T) {
	b := migratedBacklog(t, map[string][]string{
		"WORK-0015-wants-it":  {"WORK-0015"},
		"BACK-0015-has-it":    {"BACK-0015"},
		"WORK-0016-unblocked": {"WORK-0016"},
	})
	plan, err := PlanKeyMigration(b, "BACK")
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Collisions) != 1 {
		t.Fatalf("collisions = %d, want 1", len(plan.Collisions))
	}
	c := plan.Collisions[0]
	if c.Key != "BACK-0015" || c.HeldBy != "BACK-0015-has-it" || c.Former {
		t.Errorf("collision = %+v, want BACK-0015 held live by BACK-0015-has-it", c)
	}
	// The run continues: the unblocked record still migrates.
	if len(plan.Renames) != 1 || plan.Renames[0].OldKey != "WORK-0016" {
		t.Errorf("renames = %+v, want WORK-0016 to migrate anyway", plan.Renames)
	}
}

func TestPlanReportsACollisionAgainstAFormerKey(t *testing.T) {
	b := migratedBacklog(t, map[string][]string{
		"WORK-0020-wants-it":   {"WORK-0020"},
		"PROJ-0030-retired-it": {"PROJ-0030", "BACK-0020"},
	})
	plan, err := PlanKeyMigration(b, "BACK")
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Collisions) != 1 || !plan.Collisions[0].Former {
		t.Fatalf("collisions = %+v, want one blocked by a former key", plan.Collisions)
	}
}

func TestPlanLetsARecordReclaimItsOwnFormerKey(t *testing.T) {
	// Migrating back. BACK-0040 used to be WORK-0040, and WORK is the target
	// again. Without the own-key exception this is refused for every record
	// that ever moved.
	b := migratedBacklog(t, map[string][]string{
		"BACK-0040-was-work": {"BACK-0040", "WORK-0040"},
	})
	plan, err := PlanKeyMigration(b, "WORK")
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Collisions) != 0 {
		t.Fatalf("collisions = %+v, want none — a record may reclaim its own", plan.Collisions)
	}
	if len(plan.Renames) != 1 || plan.Renames[0].NewKey != "WORK-0040" {
		t.Errorf("renames = %+v, want BACK-0040 back to WORK-0040", plan.Renames)
	}
}

func TestRewriteNamesLeavesBareKeysAlone(t *testing.T) {
	renames := []KeyRename{{
		OldKey: "WORK-0031", NewKey: "BACK-0031",
		OldName: "WORK-0031-reshape-the-command-surface",
		NewName: "BACK-0031-reshape-the-command-surface",
	}}
	in := "See [[work-items/WORK-0031-reshape-the-command-surface]] and also " +
		"[[work-items/WORK-0031-reshape-the-command-surface/outcomes/x]].\n" +
		"Journaled on WORK-0031, and `WORK-0031-reshape-the-command-surface` in prose."
	got, n := RewriteNamesIn(in, renames)
	if n != 3 {
		t.Errorf("changed = %d, want 3", n)
	}
	if want := "Journaled on WORK-0031,"; !strings.Contains(got, want) {
		t.Errorf("a bare key was rewritten; output:\n%s", got)
	}
	if strings.Contains(got, "WORK-0031-reshape") {
		t.Errorf("a full name survived; output:\n%s", got)
	}
}

func TestRewriteNamesWillNotCorruptALongerName(t *testing.T) {
	// The boundary. Without it, rewriting the shorter name eats the prefix of
	// the longer one and leaves a name that points at neither record.
	renames := []KeyRename{{
		OldName: "WORK-0031-reshape", NewName: "BACK-0031-reshape",
	}}
	in := "[[work-items/WORK-0031-reshape-the-command-surface]] and [[work-items/WORK-0031-reshape]]"
	got, n := RewriteNamesIn(in, renames)
	if n != 1 {
		t.Errorf("changed = %d, want 1 — only the exact name", n)
	}
	if !strings.Contains(got, "WORK-0031-reshape-the-command-surface") {
		t.Errorf("the longer name was corrupted; output:\n%s", got)
	}
	if !strings.Contains(got, "BACK-0031-reshape]]") {
		t.Errorf("the exact name was not rewritten; output:\n%s", got)
	}
}

// RewriteNamesIn ran forever on a name that is a prefix of a longer one: it
// advanced a local index and then let the loop restart the search from the
// top, so the match it had just skipped was the first one found again. No
// error, no wrong answer, just a call that never returned --- the same shape
// as the repeating-position hang found on 2026-09-17, and caught the same way.
//
// The bound is a tripwire rather than a performance assertion, so this checks
// only that it comes back.
func TestRewriteNamesTerminatesOnAPrefixName(t *testing.T) {
	renames := []KeyRename{{OldName: "WORK-0031-reshape", NewName: "BACK-0031-reshape"}}
	in := "[[work-items/WORK-0031-reshape-the-command-surface]] twice: " +
		"[[work-items/WORK-0031-reshape-the-command-surface]]"

	done := make(chan string, 1)
	go func() {
		out, _ := RewriteNamesIn(in, renames)
		done <- out
	}()
	select {
	case got := <-done:
		if got != in {
			t.Errorf("a longer name was rewritten:\n%s", got)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("RewriteNamesIn did not return on a name that prefixes a longer one")
	}
}

func TestAllocationErrorsRatherThanSpinning(t *testing.T) {
	// The bound exists so a wrong holding check surfaces as an error a caller
	// can act on, instead of a process that never returns. Nothing reachable
	// produces it, so this asserts the ordinary path still comes back rather
	// than contriving the failure.
	b := migratedBacklog(t, map[string][]string{"BACK-0001-a": {"BACK-0001"}})
	got, err := NextAvailableKey(b, "BACK")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "BACK-0002" {
		t.Errorf("NextAvailableKey = %s, want BACK-0002", got)
	}
}
