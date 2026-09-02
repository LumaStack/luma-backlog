package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestNoCommandWritesOutsideTheBacklog runs every command that mutates
// anything and diffs the entire working tree.
//
// Snapshot-and-diff rather than an assertion about intent, because the failure
// worth catching is the silent one: the upward walk leaving the fixture and
// operating on the developer's own checkout, where everything SUCCEEDS. A test
// that asserted "the tool meant to stay inside" would pass while it did not.
func TestNoCommandWritesOutsideTheBacklog(t *testing.T) {
	app, project := initialized(t)
	before := snapshotAll(t, project)

	for _, args := range [][]string{
		{"init"},
		{"new", "deliverable", "Payments v2"},
		{"new", "outcome", "The queue drains", "-d", "payments-v2"},
		{"new", "task", "Add the queue", "-d", "payments-v2"},
		{"new", "decision", "Use a queue", "-d", "payments-v2"},
		{"new", "exploration", "Queue options", "-d", "payments-v2"},
		{"set", "payments-v2", "workflow_status=in_progress"},
		{"journal", "a line worth keeping"},
		{"verify", "the-queue-drains", "-e", "ran it"},
		{"close", "payments-v2", "-r", "delivered"},
		{"list"},
		{"show", "payments-v2"},
	} {
		if code, _, e := run(t, app, args...); code != ExitOK {
			t.Fatalf("%v exited %d: %s", args, code, e)
		}
	}

	for path, digest := range snapshotAll(t, project) {
		if strings.HasPrefix(path, ".luma/") {
			continue
		}
		if was, existed := before[path]; !existed {
			t.Errorf("created a file outside the backlog: %s", path)
		} else if was != digest {
			t.Errorf("modified a file outside the backlog: %s", path)
		}
	}
	for path := range before {
		if _, still := snapshotAll(t, project)[path]; !still {
			t.Errorf("removed a file: %s", path)
		}
	}
}

// TestCommandsRefuseToEscapeUpward runs from a nested directory and through a
// symlink, which is where root discovery goes wrong if it goes wrong at all.
func TestCommandsRefuseToEscapeUpward(t *testing.T) {
	app, project := withDeliverable(t)

	nested := filepath.Join(project, "src", "deep", "nested")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	app.WorkingDir = nested
	if code, _, e := run(t, app, "list"); code != ExitOK {
		t.Errorf("running from a nested directory failed: %d %s", code, e)
	}

	link := filepath.Join(filepath.Dir(project), "link-to-project")
	if err := os.Symlink(project, link); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	app.WorkingDir = link
	if code, _, e := run(t, app, "list"); code != ExitOK {
		t.Errorf("running through a symlink failed: %d %s", code, e)
	}
}

// snapshotAll records every file and its content, so a modification is caught
// as well as a creation.
func snapshotAll(t *testing.T, dir string) map[string]string {
	t.Helper()
	found := map[string]string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(dir, path)
		found[filepath.ToSlash(rel)] = hashOf(data)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return found
}

func hashOf(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
