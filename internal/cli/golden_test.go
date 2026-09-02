package cli

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

// update rewrites the golden files. Deliberately awkward to reach for.
//
// A golden file is a contract test: a diff in one IS a breaking change
// (docs/spec.md §9a.5). Regenerating by reflex records the bug as expected
// behavior, which is the whole risk — so this is run after reading a failure,
// never as a routine step.
var update = flag.Bool("update", false, "rewrite golden files after reading the diff")

func checkGolden(t *testing.T, name, got string) {
	t.Helper()
	path := filepath.Join("testdata", name+".golden")

	if *update {
		if err := os.MkdirAll("testdata", 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("updated %s", path)
		return
	}

	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("no golden file at %s — run: go test ./internal/cli -update\n"+
			"then READ the diff before committing it.\ngot:\n%s", path, got)
	}
	if got != string(want) {
		t.Errorf("output shape changed — this is a breaking change unless intended.\n"+
			"--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}
