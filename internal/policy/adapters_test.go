package policy

import (
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

// adapterPackages are the surfaces. Each translates an input form into a
// request and a result into a rendering, and carries no judgment of its own.
//
// A package added here is one making that promise. A surface not listed is one
// nobody has decided about, which is worth noticing.
var adapterPackages = []string{
	"internal/cli",
	"internal/board",
	"internal/web",
}

// enginePackages are what an adapter must not reach.
//
// Not because they are secret, but because reaching them means going around
// internal/app — and every validation, every policy check and every mutation
// lives there. A surface that skipped it would skip the checks: the tool's only
// refusal sat in a Cobra command until this layer existed, protected by nothing
// except there being one caller.
var enginePackages = map[string]string{
	"internal/backlog": "the record engine",
	"internal/root":    "the filesystem handle",
	"internal/config":  "configuration loading",
}

// TestAdaptersDoNotReachPastTheApplicationLayer fails when a surface imports
// the engine directly.
//
// This is the mechanism ADR-0004 rests on. That record says adapters may not
// import the engine, "enforced by a containment test rather than by
// convention" — and a rule with nothing behind it is one the second person to
// touch the code breaks. It is the same move docs/spec.md §9a.4 already makes
// for filesystem access, applied to a second seam.
//
// Test files are exempt: a test may set up a corpus with the engine, because
// arranging a fixture is not the same as a surface performing an operation.
func TestAdaptersDoNotReachPastTheApplicationLayer(t *testing.T) {
	repo := repoRoot(t)

	type violation struct{ file, pkg, why string }
	var found []violation

	for _, adapter := range adapterPackages {
		dir := filepath.Join(repo, adapter)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue // a surface that does not exist yet
		}
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return nil
			}
			file, perr := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
			if perr != nil {
				return perr
			}
			for _, imp := range file.Imports {
				p, uerr := strconv.Unquote(imp.Path.Value)
				if uerr != nil {
					continue
				}
				for engine, why := range enginePackages {
					if strings.HasSuffix(p, "/"+engine) {
						rel, _ := filepath.Rel(repo, path)
						found = append(found, violation{rel, engine, why})
					}
				}
			}
			return nil
		})
		if err != nil {
			t.Fatalf("walking %s: %v", adapter, err)
		}
	}

	if len(found) == 0 {
		return
	}
	sort.Slice(found, func(i, j int) bool { return found[i].file < found[j].file })
	var b strings.Builder
	b.WriteString("a surface imports the engine directly:\n")
	for _, v := range found {
		b.WriteString("  " + v.file + " imports " + v.pkg + " (" + v.why + ")\n")
	}
	b.WriteString("\nAdapters translate; they do not decide. Go through internal/app, and if it\n")
	b.WriteString("cannot express what is needed, that is a gap in the layer rather than a\n")
	b.WriteString("reason to reach past it (.luma/records/decisions/ADR-0004).")
	t.Error(b.String())
}
