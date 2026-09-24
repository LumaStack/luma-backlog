package root

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWalkTextSkipsBinaryAndGit(t *testing.T) {
	dir := t.TempDir()
	write := func(rel string, data []byte) {
		p := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, data, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("docs/notes.md", []byte("WORK-0031-a-thing"))
	write("internal/x.go", []byte("// WORK-0031-a-thing"))
	// A compiled binary is exactly what a naive walk corrupted.
	write("luma-backlog", []byte("ELF\x00\x01\x02 WORK-0031-a-thing"))
	write(".git/objects/ab/cdef", []byte("WORK-0031-a-thing"))

	p, err := OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { p.Close() })

	seen := map[string]bool{}
	if err := p.WalkText(func(rel string, _ []byte) error {
		seen[rel] = true
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if !seen["docs/notes.md"] || !seen["internal/x.go"] {
		t.Errorf("text files outside .luma were not visited: %v", seen)
	}
	if seen["luma-backlog"] {
		t.Error("a binary was visited — a rewrite would have corrupted it")
	}
	for k := range seen {
		if len(k) > 4 && k[:5] == ".git/" {
			t.Errorf("descended into .git: %s", k)
		}
	}
}

func TestProjectCannotEscapeTheRepository(t *testing.T) {
	dir := t.TempDir()
	outside := filepath.Join(t.TempDir(), "secret.md")
	if err := os.WriteFile(outside, []byte("not yours"), 0o644); err != nil {
		t.Fatal(err)
	}
	p, err := OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { p.Close() })

	// The fence moved from .luma/ to the repository; it did not open.
	if _, err := p.ReadFile("../" + filepath.Base(filepath.Dir(outside)) + "/secret.md"); err == nil {
		t.Error("read escaped the repository root")
	}
	if err := p.WriteFile("../escaped.md", []byte("x"), 0o644); err == nil {
		t.Error("write escaped the repository root")
	}
}
