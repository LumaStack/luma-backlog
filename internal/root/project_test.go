package root

import (
	"os"
	"path/filepath"
	"strings"
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
	if err := p.WalkText(NewIgnore(), func(rel string, _ []byte) error {
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

func TestProjectRefusesToWriteInsideGit(t *testing.T) {
	// The walk skips .git, which governs reading. This governs writing —
	// corrupting object storage or refs does not look like a migration bug, it
	// looks like a broken repository, and it damages the history that would
	// have let somebody undo the migration.
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, ".git", "objects"), 0o755); err != nil {
		t.Fatal(err)
	}
	p, err := OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { p.Close() })

	for _, name := range []string{".git/config", ".git", "./.git/objects/ab", ".git/refs/heads/main"} {
		if err := p.WriteFile(name, []byte("x"), 0o644); err == nil {
			t.Errorf("WriteFile(%q) was allowed", name)
		}
	}
	if err := p.Rename("docs", ".git/docs"); err == nil {
		t.Error("Rename into .git was allowed")
	}
	if err := p.Rename(".git/config", "config"); err == nil {
		t.Error("Rename out of .git was allowed")
	}
	// A path merely containing the letters is not the git directory.
	if err := p.WriteFile("notes.gitignore-sample", []byte("x"), 0o644); err != nil {
		t.Errorf("an ordinary file was refused: %v", err)
	}
}

func TestWalkTextSkipsAdoptedBundles(t *testing.T) {
	// An adopted bundle is a vendored copy of somebody else's published
	// content. Editing one is drift: the copy stops matching what it was taken
	// from, and the next adoption reverts the edit or conflicts with it.
	dir := t.TempDir()
	write := func(rel string) {
		p := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte("WORK-0031 mentioned here"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write(".luma/bundles/lumastack/luma-catalog/backlog/BUNDLE.md")
	write(".luma/backlog/work-items/WORK-0031-a/index.md")
	write("docs/notes.md")

	p, err := OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { p.Close() })

	seen := map[string]bool{}
	if err := p.WalkText(NewIgnore(), func(rel string, _ []byte) error {
		seen[rel] = true
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	for k := range seen {
		if strings.HasPrefix(k, ".luma/bundles/") {
			t.Errorf("descended into an adopted bundle: %s", k)
		}
	}
	// The rest of .luma is still ours and must be visited.
	if !seen[".luma/backlog/work-items/WORK-0031-a/index.md"] || !seen["docs/notes.md"] {
		t.Errorf("the exclusion was too broad: %v", seen)
	}
}

func TestIgnoreMatchesTheWaySomebodyWouldType(t *testing.T) {
	ig := NewIgnore("internal/**/*_test.go", "generated")
	cases := []struct {
		path string
		want bool
	}{
		{".luma/bundles/lumastack/luma-catalog/backlog/BUNDLE.md", true}, // default
		{"node_modules/x/y.js", true},                                    // default
		{"vendor/pkg/a.go", true},                                        // default
		{"internal/cli/key_test.go", true},                               // ** spanning a segment
		{"internal/a/b/c_test.go", true},                                 // ** spanning several
		{"internal/cli/key.go", false},                                   // not a fixture
		{"generated", true},                                              // a bare name matches itself
		{"generated/client/api.go", true},                                // and everything under it
		{"docs/notes.md", false},
		{".luma/backlog/work-items/BACK-0001-a/index.md", false}, // ours
		{"regenerated/thing.go", false},                          // not a prefix match on a name
	}
	for _, c := range cases {
		if got := ig.Match(c.path); got != c.want {
			t.Errorf("Match(%q) = %v, want %v", c.path, got, c.want)
		}
	}
}
