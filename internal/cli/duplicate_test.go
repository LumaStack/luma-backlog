package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// collide rewrites one work item's key to match another's, which is what a
// merge produces: two branches each allocated the same next key, and git merged
// two different files with no conflict to raise.
func collide(t *testing.T, project, slug, key string) {
	t.Helper()
	p := filepath.Join(project, ".luma", wiPath(t, project, slug, "index.md"))
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	out := make([]string, 0)
	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, "key: ") {
			line = "key: " + key
		}
		out = append(out, line)
	}
	if err := os.WriteFile(p, []byte(strings.Join(out, "\n")), 0o644); err != nil {
		t.Fatal(err)
	}
}

func twoWorkItems(t *testing.T) (*App, string) {
	t.Helper()
	app, project := initialized(t)
	for _, title := range []string{"Payments v2", "Search relevance"} {
		if code, _, e := run(t, app, "new", "work-item", title, "--kind", "change"); code != ExitOK {
			t.Fatalf("new %q failed: %s", title, e)
		}
	}
	return app, project
}

func TestACleanCorpusSaysNothingAboutKeys(t *testing.T) {
	// A warning that appears on ordinary runs is one people learn to scroll
	// past, and this one has to survive being ignored for months before it
	// matters once.
	app, _ := twoWorkItems(t)
	code, _, errOut := run(t, app, "list")
	if code != ExitOK {
		t.Fatalf("exit = %d", code)
	}
	if strings.TrimSpace(errOut) != "" {
		t.Errorf("a clean corpus produced output about keys:\n%q", errOut)
	}
}

func TestADuplicateKeyIsReportedAndNamesBoth(t *testing.T) {
	// Knowing a collision exists without knowing where is worse than not
	// knowing, so the report names both records and the key they share.
	app, project := twoWorkItems(t)
	collide(t, project, "search-relevance", "WORK-0001")

	code, out, errOut := run(t, app, "list")
	if code != ExitOK {
		t.Errorf("a duplicate key stopped the listing: exit %d", code)
	}
	if !strings.Contains(out, "payments-v2") || !strings.Contains(out, "search-relevance") {
		t.Errorf("the listing dropped records:\n%s", out)
	}
	for _, want := range []string{"WORK-0001", "payments-v2", "search-relevance"} {
		if !strings.Contains(errOut, want) {
			t.Errorf("the report did not name %q:\n%s", want, errOut)
		}
	}
}

func TestAFilteredListingStillSeesADuplicate(t *testing.T) {
	// The check runs over the whole work item set rather than the rows being
	// shown. A check that only fires when you were already looking in the right
	// place is not a check.
	app, project := twoWorkItems(t)
	collide(t, project, "search-relevance", "WORK-0001")

	_, _, errOut := run(t, app, "list", "outcome")
	if !strings.Contains(errOut, "WORK-0001") {
		t.Errorf("a listing that excludes work items missed the duplicate:\n%q", errOut)
	}
}

func TestClosingReportsADuplicateKey(t *testing.T) {
	// Closing writes a terminal state, so it is where acting on the wrong
	// record costs most — and a citation of this close could land on either.
	app, project := twoWorkItems(t)
	if code, _, e := run(t, app, "new", "outcome", "It drains", "-w", "payments-v2"); code != ExitOK {
		t.Fatalf("setup failed: %s", e)
	}
	if code, _, e := run(t, app, "verify", "it-drains", "-e", "measured"); code != ExitOK {
		t.Fatalf("verify failed: %s", e)
	}
	collide(t, project, "search-relevance", "WORK-0001")

	code, out, errOut := run(t, app, "close", "payments-v2", "--reason", "delivered")
	if code != ExitOK {
		t.Fatalf("a duplicate key blocked a close: exit %d, %s", code, errOut)
	}
	if !strings.Contains(out, "closed") {
		t.Errorf("the work item did not close:\n%s", out)
	}
	if !strings.Contains(errOut, "WORK-0001") {
		t.Errorf("closing said nothing about the duplicate:\n%q", errOut)
	}
}
