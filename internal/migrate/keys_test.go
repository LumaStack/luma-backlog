package migrate

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lumastack/luma-backlog/internal/root"
)

// repo builds a project with a backlog and whatever else is asked for, so a
// migration can be run end to end against real files.
func repo(t *testing.T, files map[string]string) (string, *root.Backlog) {
	t.Helper()
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	b, err := root.Create(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { b.Close() })
	for rel, body := range files {
		p := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir, b
}

func workItem(key, slug string) string {
	return "---\ntype: work-item\nkey: " + key + "\ntitle: " + slug +
		"\nworkflow_status: captured\nstage: draft\n---\n\n# " + slug + "\n"
}

func read(t *testing.T, dir, rel string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(dir, rel))
	if err != nil {
		t.Fatalf("reading %s: %v", rel, err)
	}
	return string(b)
}

func TestMigrationMovesKeysNamesAndDirectories(t *testing.T) {
	dir, b := repo(t, map[string]string{
		".luma/backlog/work-items/WORK-0031-reshape/index.md":             workItem("WORK-0031", "reshape"),
		".luma/backlog/work-items/WORK-0031-reshape/outcomes/it-works.md": "---\ntype: outcome\nwork_item: '[[work-items/WORK-0031-reshape]]'\n---\n\n# it\n",
		"docs/notes.md": "See [[work-items/WORK-0031-reshape]] and WORK-0031 in prose.\n",
		"internal/x.go": "// cited in work-items/WORK-0031-reshape for a reason\n",
	})

	res, err := Keys(dir, b, Options{Target: "BACK"})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Renames) != 1 {
		t.Fatalf("renames = %d, want 1", len(res.Renames))
	}

	// The directory moved.
	if _, err := os.Stat(filepath.Join(dir, ".luma/backlog/work-items/BACK-0031-reshape")); err != nil {
		t.Fatalf("directory did not move: %v", err)
	}

	// The record carries the new key and remembers the old one.
	idx := read(t, dir, ".luma/backlog/work-items/BACK-0031-reshape/index.md")
	if !strings.Contains(idx, "key: BACK-0031") {
		t.Errorf("key not rewritten:\n%s", idx)
	}
	if !strings.Contains(idx, `former_keys: ["WORK-0031"]`) {
		t.Errorf("former_keys not recorded:\n%s", idx)
	}

	// The child's link followed.
	child := read(t, dir, ".luma/backlog/work-items/BACK-0031-reshape/outcomes/it-works.md")
	if !strings.Contains(child, "[[work-items/BACK-0031-reshape]]") {
		t.Errorf("child link not repointed:\n%s", child)
	}

	// Outside .luma: the name moved, the bare key did not.
	notes := read(t, dir, "docs/notes.md")
	if !strings.Contains(notes, "[[work-items/BACK-0031-reshape]]") {
		t.Errorf("docs link not repointed:\n%s", notes)
	}
	if !strings.Contains(notes, "WORK-0031 in prose") {
		t.Errorf("a bare key was rewritten without the flag:\n%s", notes)
	}
	if !strings.Contains(read(t, dir, "internal/x.go"), "work-items/BACK-0031-reshape") {
		t.Error("a name in source was not repointed")
	}

	// And the bare key was reported rather than silently left.
	if len(res.BareKeys) == 0 {
		t.Error("bare keys were not reported")
	}
}

func TestDryRunWritesNothing(t *testing.T) {
	dir, b := repo(t, map[string]string{
		".luma/backlog/work-items/WORK-0031-reshape/index.md": workItem("WORK-0031", "reshape"),
		"docs/notes.md": "[[work-items/WORK-0031-reshape]]\n",
	})
	before := read(t, dir, "docs/notes.md")

	res, err := Keys(dir, b, Options{Target: "BACK", DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Renames) != 1 || len(res.Files) == 0 {
		t.Errorf("a dry run must still report what it would do: %+v", res)
	}
	if read(t, dir, "docs/notes.md") != before {
		t.Error("a dry run wrote to a file")
	}
	if _, err := os.Stat(filepath.Join(dir, ".luma/backlog/work-items/WORK-0031-reshape")); err != nil {
		t.Error("a dry run moved a directory")
	}
}

func TestCollisionLeavesOneRecordAndContinues(t *testing.T) {
	dir, b := repo(t, map[string]string{
		".luma/backlog/work-items/WORK-0015-wants/index.md": workItem("WORK-0015", "wants"),
		".luma/backlog/work-items/BACK-0015-has/index.md":   workItem("BACK-0015", "has"),
		".luma/backlog/work-items/WORK-0016-free/index.md":  workItem("WORK-0016", "free"),
	})
	res, err := Keys(dir, b, Options{Target: "BACK"})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Collisions) != 1 || res.Collisions[0].Key != "BACK-0015" {
		t.Fatalf("collisions = %+v, want one on BACK-0015", res.Collisions)
	}
	// The run continued.
	if _, err := os.Stat(filepath.Join(dir, ".luma/backlog/work-items/BACK-0016-free")); err != nil {
		t.Error("an unblocked record did not migrate")
	}
	// The blocked one is untouched.
	if _, err := os.Stat(filepath.Join(dir, ".luma/backlog/work-items/WORK-0015-wants")); err != nil {
		t.Error("a blocked record was moved anyway")
	}
}

func TestRenumberGivesTheBlockedRecordAFreeKey(t *testing.T) {
	dir, b := repo(t, map[string]string{
		".luma/backlog/work-items/WORK-0015-wants/index.md": workItem("WORK-0015", "wants"),
		".luma/backlog/work-items/BACK-0015-has/index.md":   workItem("BACK-0015", "has"),
	})
	res, err := Keys(dir, b, Options{Target: "BACK", Renumber: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Collisions) != 0 {
		t.Fatalf("collisions = %+v, want none under --renumber", res.Collisions)
	}
	if len(res.Renames) != 1 || res.Renames[0].NewKey != "BACK-0016" {
		t.Fatalf("renames = %+v, want BACK-0016", res.Renames)
	}
	idx := read(t, dir, ".luma/backlog/work-items/BACK-0016-wants/index.md")
	if !strings.Contains(idx, `former_keys: ["WORK-0015"]`) {
		t.Errorf("a renumbered record must still answer to its old key:\n%s", idx)
	}
}

func TestIncludeBareKeysRewritesProse(t *testing.T) {
	dir, b := repo(t, map[string]string{
		".luma/backlog/work-items/WORK-0031-reshape/index.md":   workItem("WORK-0031", "reshape"),
		".luma/backlog/work-items/WORK-0031-reshape/journal.md": "# Journal\n\nJournaled on WORK-0031.\n",
		"docs/notes.md": "WORK-0031 and WORK-0999 which is not ours.\n",
	})
	res, err := Keys(dir, b, Options{Target: "BACK", IncludeBareKeys: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.BareKeys) != 0 {
		t.Errorf("bare keys should have been rewritten, not reported: %+v", res.BareKeys)
	}
	notes := read(t, dir, "docs/notes.md")
	if !strings.Contains(notes, "BACK-0031 and") {
		t.Errorf("a bare key was not rewritten:\n%s", notes)
	}
	if !strings.Contains(notes, "WORK-0999") {
		t.Errorf("a key belonging to no record here was rewritten:\n%s", notes)
	}
	j := read(t, dir, ".luma/backlog/work-items/BACK-0031-reshape/journal.md")
	if !strings.Contains(j, "Journaled on BACK-0031") {
		t.Errorf("journals are in scope and were not rewritten:\n%s", j)
	}
}

func TestRunningItTwiceChangesNothing(t *testing.T) {
	dir, b := repo(t, map[string]string{
		".luma/backlog/work-items/WORK-0031-reshape/index.md": workItem("WORK-0031", "reshape"),
		"docs/notes.md": "[[work-items/WORK-0031-reshape]]\n",
	})
	if _, err := Keys(dir, b, Options{Target: "BACK"}); err != nil {
		t.Fatal(err)
	}
	after := read(t, dir, ".luma/backlog/work-items/BACK-0031-reshape/index.md")
	notes := read(t, dir, "docs/notes.md")

	res, err := Keys(dir, b, Options{Target: "BACK"})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Renames) != 0 {
		t.Errorf("a second run migrated something: %+v", res.Renames)
	}
	if read(t, dir, ".luma/backlog/work-items/BACK-0031-reshape/index.md") != after {
		t.Error("a second run changed a record")
	}
	if read(t, dir, "docs/notes.md") != notes {
		t.Error("a second run changed a file")
	}
}

func TestDryRunPredictsTheRealRun(t *testing.T) {
	// A dry run skips stamping, so without accounting for that it reads its own
	// un-stamped input and reports each record's own `key:` field as a bare key
	// "remaining" --- over-counting by exactly the number of records moved. A
	// dry run that does not predict the run is worse than none, because it is
	// believed.
	files := map[string]string{
		".luma/backlog/work-items/WORK-0031-reshape/index.md": workItem("WORK-0031", "reshape"),
		".luma/backlog/work-items/WORK-0040-other/index.md":   workItem("WORK-0040", "other"),
	}
	dirA, bA := repo(t, files)
	dry, err := Keys(dirA, bA, Options{Target: "BACK", DryRun: true})
	if err != nil {
		t.Fatal(err)
	}

	dirB, bB := repo(t, files)
	if _, err := Keys(dirB, bB, Options{Target: "BACK"}); err != nil {
		t.Fatal(err)
	}
	real2, err := Keys(dirB, bB, Options{Target: "BACK", DryRun: true})
	if err != nil {
		t.Fatal(err)
	}

	// After a real run nothing remains; the dry run beforehand must have said
	// the same rather than naming each record's own key.
	if len(real2.BareKeys) != 0 {
		t.Fatalf("after a real run, bare keys remain: %+v", real2.BareKeys)
	}
	if len(dry.BareKeys) != 0 {
		t.Errorf("the dry run reported bare keys the real run does not leave: %+v", dry.BareKeys)
	}
}
