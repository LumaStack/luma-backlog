package root

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// mkProject makes dir/.git so Discover treats dir as a project root.
func mkProject(t *testing.T, dir string) string {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dir, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestDiscoverFindsTheNearestRoot(t *testing.T) {
	base := t.TempDir()
	outer := mkProject(t, filepath.Join(base, "outer"))
	inner := mkProject(t, filepath.Join(outer, "sub", "inner"))
	deep := filepath.Join(inner, "a", "b")
	if err := os.MkdirAll(deep, 0o755); err != nil {
		t.Fatal(err)
	}

	got, err := Discover(deep, base)
	if err != nil {
		t.Fatal(err)
	}
	// The NEAREST root, not the outermost — otherwise nested checkouts
	// resolve to the wrong project.
	if want, _ := filepath.EvalSymlinks(inner); got != want {
		t.Errorf("Discover = %q, want %q", got, want)
	}
}

func TestDiscoverStopsAtTheCeiling(t *testing.T) {
	base := t.TempDir()
	// A project ABOVE the ceiling stands in for the developer's own checkout.
	mkProject(t, base)
	fence := filepath.Join(base, "fence")
	start := filepath.Join(fence, "work")
	if err := os.MkdirAll(start, 0o755); err != nil {
		t.Fatal(err)
	}

	_, err := Discover(start, fence)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Discover climbed past the ceiling: err = %v", err)
	}
}

func TestDiscoverTreatsAGitFileAsARoot(t *testing.T) {
	// Linked worktrees and submodules mark the root with a .git FILE.
	base := t.TempDir()
	dir := filepath.Join(base, "worktree")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".git"), []byte("gitdir: /elsewhere\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := Discover(dir, base)
	if err != nil {
		t.Fatalf("a .git file was not recognized: %v", err)
	}
	if want, _ := filepath.EvalSymlinks(dir); got != want {
		t.Errorf("Discover = %q, want %q", got, want)
	}
}

func TestDiscoverResolvesSymlinksBeforeComparing(t *testing.T) {
	// On macOS t.TempDir() hands back /var/... while the real path is
	// /private/var/..., so an unresolved ceiling never matches and the walk
	// escapes silently. This is the exact bug the resolution step prevents.
	base := t.TempDir()
	real := filepath.Join(base, "real")
	if err := os.MkdirAll(filepath.Join(real, "work"), 0o755); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(base, "link")
	if err := os.Symlink(real, link); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	mkProject(t, base)

	if _, err := Discover(filepath.Join(link, "work"), real); !errors.Is(err, ErrNotFound) {
		t.Errorf("walk escaped through a symlinked path: err = %v", err)
	}
}

func TestBacklogCannotEscapeItsRoot(t *testing.T) {
	project := mkProject(t, filepath.Join(t.TempDir(), "project"))
	outside := filepath.Join(project, "SHOULD-NOT-EXIST.txt")

	b, err := Create(project)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()

	for _, name := range []string{
		"../SHOULD-NOT-EXIST.txt",
		"../../SHOULD-NOT-EXIST.txt",
		"nested/../../SHOULD-NOT-EXIST.txt",
	} {
		if err := b.WriteFile(name, []byte("escaped"), 0o644); err == nil {
			t.Errorf("WriteFile(%q) succeeded; it must not", name)
		}
	}
	if _, err := os.Stat(outside); err == nil {
		t.Fatal("a file was written outside the backlog")
	}
}

func TestBacklogCannotEscapeThroughASymlink(t *testing.T) {
	project := mkProject(t, filepath.Join(t.TempDir(), "project"))
	b, err := Create(project)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()

	// A symlink inside the backlog pointing out of it.
	if err := os.Symlink(project, filepath.Join(b.Path(), "out")); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	if err := b.WriteFile("out/ESCAPED.txt", []byte("escaped"), 0o644); err == nil {
		t.Error("wrote through a symlink out of the backlog")
	}
	if _, err := os.Stat(filepath.Join(project, "ESCAPED.txt")); err == nil {
		t.Fatal("a file was written outside the backlog via a symlink")
	}
}

func TestBacklogRoundTrip(t *testing.T) {
	project := mkProject(t, filepath.Join(t.TempDir(), "project"))
	b, err := Create(project)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()

	const name = "backlog/deliverables/thing/index.md"
	if err := b.WriteFile(name, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	if !b.Exists(name) {
		t.Error("Exists = false after writing")
	}
	got, err := b.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "hello" {
		t.Errorf("ReadFile = %q, want %q", got, "hello")
	}
	if !strings.HasSuffix(b.Path(), Dir) {
		t.Errorf("Path = %q, want it to end in %q", b.Path(), Dir)
	}
}
