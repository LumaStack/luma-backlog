// Package policy holds checks that are rules rather than behavior.
//
// They live as tests so they run on every `go test`, not only in continuous
// integration. A guardrail you have to push to discover is one people learn
// to work around.
package policy

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// allowedPackage is the one place permitted to reach the filesystem directly.
// Everything else takes a bounded handle from it (docs/spec.md §9a.4).
const allowedPackage = "internal/root"

// deniedFuncs are the os functions that touch the filesystem by path. Calls
// that do not take a path — Getenv, Exit, Args, the standard streams — are
// fine and deliberately absent.
var deniedFuncs = map[string]bool{
	"Open": true, "OpenFile": true, "OpenRoot": true, "Create": true, "CreateTemp": true,
	"ReadFile": true, "WriteFile": true, "ReadDir": true,
	"Mkdir": true, "MkdirAll": true, "MkdirTemp": true,
	"Remove": true, "RemoveAll": true, "Rename": true, "Truncate": true,
	"Stat": true, "Lstat": true, "Symlink": true, "Link": true, "Readlink": true,
	"Chmod": true, "Chown": true, "Chtimes": true,
}

// TestFilesystemAccessIsConfinedToOnePackage fails when production code
// reaches the filesystem outside the package allowed to.
//
// This is enforced rather than conventional because agents write code here,
// and a convention nobody carries across sessions is not a guardrail.
//
// Test files are exempt: fixtures legitimately build real directories, and
// they are not what ships.
func TestFilesystemAccessIsConfinedToOnePackage(t *testing.T) {
	repo := repoRoot(t)
	var violations []string

	err := filepath.WalkDir(repo, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if name := d.Name(); name == ".git" || name == "vendor" || strings.HasPrefix(name, ".") && name != "." {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		rel, _ := filepath.Rel(repo, path)
		if strings.HasPrefix(filepath.ToSlash(rel), allowedPackage+"/") {
			return nil
		}
		violations = append(violations, scanFile(t, path, rel)...)
		return nil
	})
	if err != nil {
		t.Fatalf("walking %s: %v", repo, err)
	}

	sort.Strings(violations)
	if len(violations) > 0 {
		t.Errorf("filesystem access outside %s/:\n  %s\n\n"+
			"Take a *root.Backlog and use its methods. If this genuinely belongs\n"+
			"in %s, move it there rather than widening this rule.",
			allowedPackage, strings.Join(violations, "\n  "), allowedPackage)
	}
}

func scanFile(t *testing.T, path, rel string) []string {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		t.Fatalf("parsing %s: %v", rel, err)
	}

	var found []string
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		pkg, ok := sel.X.(*ast.Ident)
		if !ok {
			return true
		}
		if (pkg.Name == "os" || pkg.Name == "ioutil") && deniedFuncs[sel.Sel.Name] {
			pos := fset.Position(call.Pos())
			found = append(found, rel+":"+strconv.Itoa(pos.Line)+" calls "+pkg.Name+"."+sel.Sel.Name)
		}
		return true
	})
	return found
}

// repoRoot climbs to the directory holding go.mod.
func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	for {
		// os.Stat is fine here: this is a test file, and test files are
		// exempt from the very rule being enforced below.
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find go.mod above the test directory")
		}
		dir = parent
	}
}
