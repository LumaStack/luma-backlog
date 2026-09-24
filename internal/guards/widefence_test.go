package guards

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// wideFencePackage is the one place permitted to open a handle to the whole
// repository. Everything else takes a *root.Backlog, which stops at `.luma/`.
const wideFencePackage = "internal/migrate"

// wideFenceFuncs are the constructors that hand back more than the backlog.
var wideFenceFuncs = map[string]bool{"OpenProject": true}

// TestTheWideFenceIsConfinedToMigration fails when production code outside the
// migration package opens a handle reaching past `.luma/`.
//
// **Two fences exist and only one of them is ordinary.** `root.Backlog` stops
// at the backlog directory, and every command in this tool works through it.
// `root.Project` reaches the whole repository, which a key migration needs
// because a work item's name appears in prose, in documentation and in source
// comments --- and a migration that stopped at `.luma/` would leave all of
// those pointing at a directory that no longer exists.
//
// **The wider fence is a capability, not a default.** spec.md §9a.4 is content
// with either — its concern is writing outside the *repository* — so nothing
// about `root.Project` is unsafe. What would be unsafe is it becoming the
// handle everything happens to use, at which point an ordinary bug in an
// ordinary command can write to `internal/` or `docs/` rather than to a
// record.
//
// **This is the rule that keeps migrations separable without a second binary.**
// The argument for splitting them out was never subject matter; it was that
// they need a capability nothing else should have. A guard gives that
// separation for the cost of one test, and leaves the tool one thing to
// install.
//
// Test files are exempt, as in the filesystem rule: fixtures legitimately
// build real trees, and they are not what ships.
func TestTheWideFenceIsConfinedToMigration(t *testing.T) {
	repo := repoRoot(t)
	var violations []string

	err := filepath.WalkDir(repo, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if name := d.Name(); name == ".git" || name == "vendor" || (strings.HasPrefix(name, ".") && name != ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		rel := filepath.ToSlash(mustRel(t, repo, path))
		// root defines it; migrate is the one package allowed to ask for it.
		if strings.HasPrefix(rel, allowedPackage+"/") || strings.HasPrefix(rel, wideFencePackage+"/") {
			return nil
		}
		violations = append(violations, scanWideFence(t, path, rel)...)
		return nil
	})
	if err != nil {
		t.Fatalf("walking %s: %v", repo, err)
	}

	sort.Strings(violations)
	if len(violations) > 0 {
		t.Errorf("a handle reaching past .luma/ was opened outside %s/:\n  %s\n\n"+
			"Take a *root.Backlog instead. If this genuinely is migration work,\n"+
			"move it into %s rather than widening this rule.",
			wideFencePackage, strings.Join(violations, "\n  "), wideFencePackage)
	}
}

func scanWideFence(t *testing.T, path, rel string) []string {
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
		if pkg.Name == "root" && wideFenceFuncs[sel.Sel.Name] {
			pos := fset.Position(call.Pos())
			found = append(found, rel+":"+strconv.Itoa(pos.Line)+" calls root."+sel.Sel.Name)
		}
		return true
	})
	return found
}

func mustRel(t *testing.T, base, path string) string {
	t.Helper()
	rel, err := filepath.Rel(base, path)
	if err != nil {
		t.Fatalf("relativizing %s: %v", path, err)
	}
	return rel
}
