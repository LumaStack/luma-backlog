package root

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func newBacklog(t *testing.T) *Backlog {
	t.Helper()
	project := filepath.Join(t.TempDir(), "project")
	if err := os.MkdirAll(filepath.Join(project, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	b, err := Create(project)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { b.Close() })
	return b
}

func TestWriteFileAtomicRoundTrips(t *testing.T) {
	b := newBacklog(t)
	const name = "deliverables/thing/index.md"

	if err := b.WriteFileAtomic(name, []byte("first"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := b.WriteFileAtomic(name, []byte("second"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := b.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "second" {
		t.Errorf("ReadFile = %q, want %q", got, "second")
	}
}

func TestWriteFileAtomicLeavesNoTemporaryFiles(t *testing.T) {
	// A staging file left behind would be picked up by anything that lists
	// records, and would look like a corrupt one.
	b := newBacklog(t)
	if err := b.WriteFileAtomic("a.md", []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	entries, err := os.ReadDir(b.Path())
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if strings.Contains(e.Name(), ".tmp-") {
			t.Errorf("temporary file left behind: %s", e.Name())
		}
	}
}

func TestWriteFileAtomicCannotEscape(t *testing.T) {
	b := newBacklog(t)
	if err := b.WriteFileAtomic("../ESCAPED.md", []byte("x"), 0o644); err == nil {
		t.Error("atomic write escaped the backlog")
	}
	if _, err := os.Stat(filepath.Join(filepath.Dir(b.Path()), "ESCAPED.md")); err == nil {
		t.Fatal("a file was written outside the backlog")
	}
}
