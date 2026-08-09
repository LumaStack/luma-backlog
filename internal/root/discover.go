// Package root finds the project root and hands out a bounded handle to the
// backlog directory.
//
// It is the only package permitted to reach the filesystem directly, and
// everything it returns is scoped so a caller cannot climb out. The failure
// this exists to prevent is not a loud escape — it is the upward walk leaving
// the intended tree and finding another repository, where operations SUCCEED
// against the wrong target and nothing reports a problem (docs/SPEC.md §9a.4).
package root

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Dir is the name of the backlog directory within a project.
const Dir = ".backlog"

// ErrNotFound means no project root was located before the search stopped.
var ErrNotFound = errors.New("no project root found: no .git between here and the ceiling")

// Discover walks upward from start looking for a directory that contains
// .git, and returns that directory.
//
// This function walks up, and it is the ONLY one that does. Everything
// downstream is handed an already-bounded root, so the dangerous operation
// happens once, in a place small enough to test exhaustively.
//
// ceiling, when non-empty, is a directory the search will not pass. The search
// stops after examining it. Tests set it so an escape fails loudly instead of
// finding the developer's own checkout and quietly succeeding.
//
// .git may be a file rather than a directory — that is how git marks a linked
// worktree or a submodule — so both count.
func Discover(start, ceiling string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("resolving %q: %w", start, err)
	}
	// Resolve symlinks before comparing paths. On macOS a temporary directory
	// is /var/folders/... while its real path is /private/var/folders/...,
	// so an unresolved ceiling silently never matches.
	if resolved, err := filepath.EvalSymlinks(dir); err == nil {
		dir = resolved
	}
	if ceiling != "" {
		if abs, err := filepath.Abs(ceiling); err == nil {
			ceiling = abs
			if resolved, err := filepath.EvalSymlinks(ceiling); err == nil {
				ceiling = resolved
			}
		}
	}

	for {
		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}

		if ceiling != "" && dir == ceiling {
			return "", ErrNotFound
		}
		parent := filepath.Dir(dir)
		if parent == dir { // reached the filesystem root
			return "", ErrNotFound
		}
		dir = parent
	}
}
