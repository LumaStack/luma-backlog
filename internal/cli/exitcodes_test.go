package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestExitCodes covers every code the build can reach, in one place.
//
// Exit codes are the most machine-facing part of the contract: a caller
// branches on 4 to re-read and retry, and on 5 to satisfy a condition first.
// Collapsing either into "something went wrong" removes the only signal an
// unattended agent has.
//
// They live in a Go test rather than a script test because the script
// frameworks in this ecosystem assert success or failure and not a specific
// status (docs/testing.md).
func TestExitCodes(t *testing.T) {
	t.Run("0 success", func(t *testing.T) {
		app, _ := initialized(t)
		if code, _, e := run(t, app, "list"); code != ExitOK {
			t.Errorf("exit = %d, want 0: %s", code, e)
		}
	})

	t.Run("2 usage — unknown command", func(t *testing.T) {
		app, _ := initialized(t)
		if code, _, _ := run(t, app, "definitely-not-a-command"); code != ExitUsage {
			t.Errorf("exit = %d, want %d", code, ExitUsage)
		}
	})

	t.Run("2 usage — bad arguments", func(t *testing.T) {
		app, _ := initialized(t)
		if code, _, _ := run(t, app, "new", "sprint", "Thing"); code != ExitUsage {
			t.Errorf("exit = %d, want %d", code, ExitUsage)
		}
	})

	t.Run("3 not found", func(t *testing.T) {
		app, _ := initialized(t)
		if code, _, _ := run(t, app, "show", "no-such-record"); code != ExitNotFound {
			t.Errorf("exit = %d, want %d", code, ExitNotFound)
		}
	})

	t.Run("4 conflict — re-read and retry", func(t *testing.T) {
		app, project := withDeliverable(t)
		_, out, _ := run(t, app, "show", "payments-v2", "--json")
		var seen struct {
			Hash string `json:"hash"`
		}
		json.Unmarshal([]byte(out), &seen)

		path := filepath.Join(project, ".luma", "backlog/deliverables/payments-v2/index.md")
		data, _ := os.ReadFile(path)
		os.WriteFile(path, append(data, []byte("\nchanged\n")...), 0o644)

		if code, _, _ := run(t, app, "set", "payments-v2", "a=b", "--if-unchanged", seen.Hash); code != ExitConflict {
			t.Errorf("exit = %d, want %d", code, ExitConflict)
		}
	})

	t.Run("5 refused — satisfy the condition first", func(t *testing.T) {
		app, _ := withOutcomes(t)
		if code, _, _ := run(t, app, "close", "payments-v2", "-r", "delivered"); code != ExitRefused {
			t.Errorf("exit = %d, want %d", code, ExitRefused)
		}
	})

	// 1 (unexpected) and 6 (already claimed) are unreachable in this build:
	// nothing returns a bare failure on a path a test can drive, and claiming
	// is out of scope. Recorded rather than faked — a test that manufactures
	// a code proves the constant exists, not that anything produces it.
}
