package corpus

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lumastack/luma-backlog/internal/root"
)

// The spellings come from WORK-0082's capture, and every one of them names
// WORK-0074. The table is the record's own list, pinned.
var spellings = []string{
	"WORK-0074",
	"WORK-74",
	"work-74",
	"WORK-0000000074",
	"WORK      74",
	"WORK---74",
	"WoRk-74",
}

func TestParseKeyReadsEverySpelling(t *testing.T) {
	for _, s := range spellings {
		prefix, number, ok := ParseKey(s)
		if !ok {
			t.Errorf("ParseKey(%q) is not a key", s)
			continue
		}
		if prefix != "WORK" || number != 74 {
			t.Errorf("ParseKey(%q) = %q, %d; want WORK, 74", s, prefix, number)
		}
	}
}

func TestParseKeyFollowsJiraCloudPrefixRules(t *testing.T) {
	// An uppercase letter first, then uppercase letters or digits, two to
	// ten characters. R2D2 is Atlassian's own example of a legal key.
	if prefix, number, ok := ParseKey("R2D2-7"); !ok || prefix != "R2D2" || number != 7 {
		t.Errorf("ParseKey(R2D2-7) = %q, %d, %v; want R2D2, 7, true", prefix, number, ok)
	}

	for _, s := range []string{
		"X-1",           // one-letter prefix: below Jira Cloud's minimum
		"ABCDEFGHIJK-1", // eleven characters: past the maximum
		"2XY-4",         // first character is not a letter
		"WORK_74",       // underscore is not a separator here
		"74",            // a bare number is not a key (deferred, WORK-0082)
		"add-2-buttons", // a slug with a digit inside stays a slug
	} {
		if _, _, ok := ParseKey(s); ok {
			t.Errorf("ParseKey(%q) claims a key; want a slug", s)
		}
	}
}

func TestNormalizeKeyRendersTheCanonicalForm(t *testing.T) {
	for _, s := range spellings {
		if got := NormalizeKey(s); got != "WORK-0074" {
			t.Errorf("NormalizeKey(%q) = %q, want WORK-0074", s, got)
		}
	}
	// The width is a minimum: a five-digit key is never squeezed back.
	if got := NormalizeKey("WORK-12345"); got != "WORK-12345" {
		t.Errorf("NormalizeKey(WORK-12345) = %q, want WORK-12345", got)
	}
	// Not a key: returned untouched, since it is somebody's slug.
	if got := NormalizeKey("lint-the-corpus"); got != "lint-the-corpus" {
		t.Errorf("NormalizeKey(lint-the-corpus) = %q, want it unchanged", got)
	}
}

func TestSameKeyComparesParsedValues(t *testing.T) {
	for _, s := range spellings {
		if !SameKey(s, "WORK-0074") {
			t.Errorf("SameKey(%q, WORK-0074) = false", s)
		}
	}
	if SameKey("WORK-74", "WORK-75") {
		t.Error("SameKey(WORK-74, WORK-75) = true")
	}
	if SameKey("WORK-74", "TASK-74") {
		t.Error("SameKey(WORK-74, TASK-74) = true; prefixes differ")
	}
	if SameKey("WORK-74", "lint-the-corpus") {
		t.Error("SameKey(WORK-74, lint-the-corpus) = true; a key never equals a slug")
	}
	if !SameKey("lint-the-corpus", "lint-the-corpus") {
		t.Error("SameKey on two equal slugs = false; plain equality should answer")
	}
}

func TestNormalizeNameCanonicalizesTheKeyHalf(t *testing.T) {
	for ref, want := range map[string]string{
		"work-74-lint-the-corpus":   "WORK-0074-lint-the-corpus",
		"WORK-0074-lint-the-corpus": "WORK-0074-lint-the-corpus",
		"WoRk-0074-lint-the-corpus": "WORK-0074-lint-the-corpus",
		"lint-the-corpus":           "lint-the-corpus",
	} {
		if got := NormalizeName(ref); got != want {
			t.Errorf("NormalizeName(%q) = %q, want %q", ref, got, want)
		}
	}
}

// keyedBacklog writes a corpus by hand, because the fixtures need keys the
// creation path would never write — that is the point of the tests below.
func keyedBacklog(t *testing.T, keys map[string]string) *root.Backlog {
	t.Helper()
	project := filepath.Join(t.TempDir(), "project")
	if err := os.MkdirAll(filepath.Join(project, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	b, err := root.Create(project)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { b.Close() })

	for name, key := range keys {
		body := "---\n" +
			"type: work-item\n" +
			"key: " + key + "\n" +
			"title: " + name + "\n" +
			"workflow_status: captured\n" +
			"stage: draft\n" +
			"---\n\n# " + name + "\n"
		rel := "backlog/work-items/" + name + "/index.md"
		if err := b.WriteFileAtomic(rel, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return b
}

func TestResolveFindsEveryKeySpelling(t *testing.T) {
	b := keyedBacklog(t, map[string]string{
		"WORK-0074-large-uploads-fail": "WORK-0074",
		"WORK-0075-another-record":     "WORK-0075",
		"R2D2-0007-a-jira-cloud-key":   "R2D2-0007",
	})

	refs := append([]string{}, spellings...)
	// The joined form resolves under the same rule for its key half.
	refs = append(refs, "work-74-large-uploads-fail", "WoRk-0074-large-uploads-fail")

	before, err := b.ReadFile("backlog/work-items/WORK-0074-large-uploads-fail/index.md")
	if err != nil {
		t.Fatal(err)
	}
	for _, ref := range refs {
		it, err := Resolve(b, ref)
		if err != nil {
			t.Errorf("Resolve(%q): %v", ref, err)
			continue
		}
		if it.Key() != "WORK-0074" {
			t.Errorf("Resolve(%q) found %s, want WORK-0074", ref, it.Key())
		}
	}
	// A prefix under Jira Cloud rules resolves too — recognition is not only
	// for WORK.
	it, err := Resolve(b, "r2d2-7")
	if err != nil {
		t.Errorf("Resolve(r2d2-7): %v", err)
	} else if it.Key() != "R2D2-0007" {
		t.Errorf("Resolve(r2d2-7) found %s, want R2D2-0007", it.Key())
	}

	// Resolution reads; it must never rewrite what is stored.
	after, err := b.ReadFile("backlog/work-items/WORK-0074-large-uploads-fail/index.md")
	if err != nil {
		t.Fatal(err)
	}
	if string(before) != string(after) {
		t.Error("resolving rewrote a stored record")
	}
}

func TestDuplicatesSeeOneKeyAcrossSpellings(t *testing.T) {
	// Two records holding two spellings of one key is one duplicate, not two
	// keys — otherwise the detector WORK-0014 shipped is blind to exactly the
	// collision ADR-0003 exists to repair.
	b := keyedBacklog(t, map[string]string{
		"WORK-0074-large-uploads-fail": "WORK-0074",
		"WORK-74-the-same-key-again":   "WORK-74",
	})
	items, _, err := List(b, Filter{})
	if err != nil {
		t.Fatal(err)
	}
	dups := Duplicates(items)
	if len(dups) != 1 {
		t.Fatalf("Duplicates = %d, want 1: %v", len(dups), dups)
	}
	if dups[0].Key != "WORK-0074" || len(dups[0].Paths) != 2 {
		t.Errorf("Duplicates[0] = %+v; want key WORK-0074 with both paths", dups[0])
	}
}

func TestHighestKeyCountsUnusualSpellings(t *testing.T) {
	// A key stored unpadded must still hold its number, or the next
	// allocation reuses it — WORK-0040's failure from a different cause.
	b := keyedBacklog(t, map[string]string{
		"WORK-0074-large-uploads-fail": "WORK-0074",
		"WORK-88-stored-unpadded":      "WORK-88",
	})
	highest, err := highestKey(b)
	if err != nil {
		t.Fatal(err)
	}
	if highest != 88 {
		t.Errorf("highestKey = %d, want 88", highest)
	}
}
