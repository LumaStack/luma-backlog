package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// breakRecord damages a record the way a careless hand edit does: the opening
// --- is gone, so nothing in the file is frontmatter any more and the record
// stops parsing.
func breakRecord(t *testing.T, project, rel string) {
	t.Helper()
	path := filepath.Join(project, ".luma", rel)
	if err := os.WriteFile(path, []byte("type: work-item\ntitle: broken\n---\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestListNamesWhatItSkipped(t *testing.T) {
	// The behavior this whole work item exists for: the record is still
	// skipped and the command still succeeds, but the reader is told.
	app, project := initialized(t)
	seed(t, app)
	breakRecord(t, project, wiPath(t, project, "payments-v2", "index.md"))

	code, out, errOut := run(t, app, "list")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	if !strings.Contains(errOut, "payments-v2/index.md") {
		t.Errorf("the skipped record was not named:\nstderr: %q", errOut)
	}
	if strings.Contains(out, "Payments v2") {
		t.Errorf("an unparseable record was listed anyway:\n%s", out)
	}
}

func TestSkipReportStaysOutOfStdout(t *testing.T) {
	// A caller piping the listing must be unaffected by a broken record. If
	// the warning reached stdout the JSON would stop parsing, which turns a
	// helpful message into a second outage.
	app, project := initialized(t)
	seed(t, app)
	breakRecord(t, project, wiPath(t, project, "payments-v2", "index.md"))

	code, out, errOut := run(t, app, "list", "--json")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	var rows []map[string]any
	if err := json.Unmarshal([]byte(out), &rows); err != nil {
		t.Fatalf("stdout stopped being valid JSON: %v\n%s", err, out)
	}
	if !strings.Contains(errOut, "payments-v2/index.md") {
		t.Errorf("nothing was reported on stderr:\n%q", errOut)
	}
}

func TestCloseReportsAnOutcomeItCouldNotRead(t *testing.T) {
	// The shape that makes silence unsafe: one outcome passes, another is
	// unreadable. The unreadable one is missing from Live, so it can never be
	// counted in Unpassing — the arithmetic runs on a denominator it has
	// quietly lost, and "delivered" comes out looking exactly like a correct
	// answer. Refusing is a separate work item; being told is this one.
	app, project := initialized(t)
	seed(t, app)
	if code, _, e := run(t, app, "new", "outcome", "Latency holds", "-w", "payments-v2"); code != ExitOK {
		t.Fatalf("new outcome failed: %s", e)
	}
	if code, _, e := run(t, app, "verify", "latency-holds", "-e", "measured"); code != ExitOK {
		t.Fatalf("verify failed: %s", e)
	}
	breakRecord(t, project, wiPath(t, project, "payments-v2", "outcomes", "the-retry-queue-drains.md"))

	_, _, errOut := run(t, app, "close", "payments-v2", "--reason", "delivered")
	if !strings.Contains(errOut, "the-retry-queue-drains.md") {
		t.Errorf("close drew a conclusion without mentioning the outcome it could not read:\nstderr: %q", errOut)
	}
}

func TestNothingIsSaidWhenNothingWasSkipped(t *testing.T) {
	// A warning that appears on ordinary runs is a warning people learn to
	// scroll past.
	app := populated(t)
	code, _, errOut := run(t, app, "list")
	if code != ExitOK {
		t.Fatalf("exit = %d", code)
	}
	if strings.TrimSpace(errOut) != "" {
		t.Errorf("a clean backlog produced noise on stderr: %q", errOut)
	}
}
