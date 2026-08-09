package root

import (
	"fmt"
	"os"
	"path/filepath"
)

// Backlog is a bounded handle to a project's backlog directory. Every path is
// interpreted relative to it, and none can escape it — including through ..
// or a symlink.
//
// This is depth rather than a guarantee. The standard library's implementation
// checks a path component at a time rather than in a single kernel operation,
// has had escapes patched more than once, and disclaims bind mounts outright.
// It raises the cost of the bug considerably; the fenced environment in
// docs/TESTING.md is the other half rather than a redundancy.
type Backlog struct {
	root *os.Root
	path string
}

// Open returns a handle to an existing backlog under projectRoot.
func Open(projectRoot string) (*Backlog, error) {
	p := filepath.Join(projectRoot, Dir)
	r, err := os.OpenRoot(p)
	if err != nil {
		return nil, fmt.Errorf("opening backlog at %s: %w", p, err)
	}
	return &Backlog{root: r, path: p}, nil
}

// Create makes the backlog directory if absent and returns a handle to it.
func Create(projectRoot string) (*Backlog, error) {
	p := filepath.Join(projectRoot, Dir)
	if err := os.MkdirAll(p, 0o755); err != nil {
		return nil, fmt.Errorf("creating backlog at %s: %w", p, err)
	}
	return Open(projectRoot)
}

// Path is the absolute location of the backlog, for messages and for git.
// It is not a way to reach the filesystem: callers use the methods below.
func (b *Backlog) Path() string { return b.path }

// Close releases the handle.
func (b *Backlog) Close() error { return b.root.Close() }

// ReadFile reads a file within the backlog.
func (b *Backlog) ReadFile(name string) ([]byte, error) { return b.root.ReadFile(name) }

// WriteFile writes a file within the backlog, creating parent directories.
func (b *Backlog) WriteFile(name string, data []byte, perm os.FileMode) error {
	if dir := filepath.Dir(name); dir != "." {
		if err := b.root.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}
	return b.root.WriteFile(name, data, perm)
}

// MkdirAll creates a directory within the backlog.
func (b *Backlog) MkdirAll(name string) error { return b.root.MkdirAll(name, 0o755) }

// Stat reports on a path within the backlog.
func (b *Backlog) Stat(name string) (os.FileInfo, error) { return b.root.Stat(name) }

// Exists reports whether a path is present.
func (b *Backlog) Exists(name string) bool {
	_, err := b.root.Stat(name)
	return err == nil
}

// Rename moves a path within the backlog.
func (b *Backlog) Rename(from, to string) error { return b.root.Rename(from, to) }
