package root

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Project is a bounded handle to the whole repository, not just its backlog.
//
// **Why a second fence exists, and why it is not a weakening.** `Backlog`
// stops at `.luma/`, which was the right boundary while every write this tool
// made was a record. A key migration is not: a work item's name appears in
// prose, in documentation, and in source comments, and a migration that
// stopped at `.luma/` would leave those pointing at a directory that no longer
// exists --- silently, because a name is not a key and nothing resolves it.
//
// **The property spec.md §9a.4 actually states is preserved exactly.** Its
// concern is the tool "writing outside their repository", not outside `.luma/`.
// This handle is an os.Root at the repository root, so traversal out --- via
// `..`, via a symlink --- is resisted the same way, by the same mechanism, with
// the same caveats. The fence moves; it does not open.
//
// It carries the narrowest API that a repository-wide rewrite needs. Anything
// wanting more should have a reason written down, because every method here is
// one more thing reaching past `.luma/`.
type Project struct {
	root *os.Root
	path string
}

// OpenProject returns a handle to the repository containing a backlog.
func OpenProject(projectRoot string) (*Project, error) {
	r, err := os.OpenRoot(projectRoot)
	if err != nil {
		return nil, fmt.Errorf("opening project at %s: %w", projectRoot, err)
	}
	return &Project{root: r, path: projectRoot}, nil
}

// Path is the absolute location of the repository, for messages.
func (p *Project) Path() string { return p.path }

// Close releases the handle.
func (p *Project) Close() error { return p.root.Close() }

// ReadFile reads a file relative to the repository root.
func (p *Project) ReadFile(name string) ([]byte, error) { return p.root.ReadFile(name) }

// WriteFile replaces a file's contents relative to the repository root.
func (p *Project) WriteFile(name string, data []byte, perm os.FileMode) error {
	f, err := p.root.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return fmt.Errorf("opening %s: %w", name, err)
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		return fmt.Errorf("writing %s: %w", name, err)
	}
	return f.Close()
}

// Rename moves a file or directory within the repository.
func (p *Project) Rename(from, to string) error { return p.root.Rename(from, to) }

// skipDirs are never descended into.
//
// `.git` because rewriting object storage would corrupt the repository, and
// because nothing in it is a name anybody wrote. The rest are build output and
// vendored code: not ours to edit, regenerated anyway, and the place a naive
// walk does the most damage.
var skipDirs = map[string]bool{
	".git": true, "node_modules": true, "vendor": true, "dist": true, "build": true,
}

// WalkText visits every text file in the repository, giving each path relative
// to it, slash-separated whatever the platform.
//
// **Binary files are skipped, and this is the check that matters.** Counting
// key mentions across the tree once matched inside the compiled binary sitting
// in the repository root, on a Go runtime error string. A rewrite doing the
// same would corrupt whatever it touched --- an executable, an image, a
// vendored archive --- and the damage would not look like a migration bug.
//
// The test is a NUL byte in the first few kilobytes, which is what git itself
// uses to decide the same question. It is a heuristic, and it is the same
// heuristic everything else in this ecosystem already trusts.
func (p *Project) WalkText(fn func(relPath string, data []byte) error) error {
	return fs.WalkDir(p.root.FS(), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if skipDirs[d.Name()] {
				return fs.SkipDir
			}
			return nil
		}
		if strings.HasPrefix(d.Name(), ".") && d.Name() != ".gitignore" {
			return nil
		}
		data, readErr := p.root.ReadFile(path)
		if readErr != nil {
			// A file that cannot be read is not one to rewrite. A broken
			// symlink or a permission the operator meant is not this
			// command's business to report on.
			return nil
		}
		if isBinary(data) {
			return nil
		}
		return fn(filepath.ToSlash(path), data)
	})
}

// isBinary reports whether content looks like something no one typed.
func isBinary(data []byte) bool {
	head := data
	if len(head) > 8000 {
		head = head[:8000]
	}
	return bytes.IndexByte(head, 0) >= 0
}
