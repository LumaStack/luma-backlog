package corpus

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lumastack/luma-backlog/internal/root"
)

// migratedBacklog writes records that a migration would have produced, which
// no creation path writes: a current key plus the keys the record used to
// answer to.
//
// Everything below constructs its corpus deliberately, and that is the point.
// A collision cannot occur in an ordinary corpus — allocation draws from one
// monotonic sequence, so the number ranges never overlap — which makes these
// paths the ones most likely to ship broken and least likely to be noticed by
// anybody using the tool.
func migratedBacklog(t *testing.T, records map[string][]string) *root.Backlog {
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

	for name, keys := range records {
		body := "---\ntype: work-item\nkey: " + keys[0] + "\n"
		if len(keys) > 1 {
			body += "former_keys: ["
			for i, k := range keys[1:] {
				if i > 0 {
					body += ", "
				}
				body += `"` + k + `"`
			}
			body += "]\n"
		}
		body += "title: " + name + "\nworkflow_status: captured\nstage: draft\n---\n\n# " + name + "\n"
		rel := "backlog/work-items/" + name + "/index.md"
		if err := b.WriteFileAtomic(rel, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return b
}

func TestFormerKeyResolvesToItsRecord(t *testing.T) {
	b := migratedBacklog(t, map[string][]string{
		"BACK-0036-a-migrated-record": {"BACK-0036", "WORK-0036"},
		"BACK-0040-never-migrated":    {"BACK-0040"},
	})

	// The old key still answers, in every spelling the current one accepts.
	for _, ref := range []string{"WORK-0036", "work-36", "WORK---0036", "work  36"} {
		it, err := Resolve(b, ref)
		if err != nil {
			t.Errorf("Resolve(%q): %v", ref, err)
			continue
		}
		if it.Key() != "BACK-0036" {
			t.Errorf("Resolve(%q) found %s, want BACK-0036", ref, it.Key())
		}
	}

	// And the record reports the current key, so a reader learns the new one
	// rather than being handed back what they typed.
	it, err := Resolve(b, "WORK-0036")
	if err != nil {
		t.Fatal(err)
	}
	if got := it.FormerKeys(); len(got) != 1 || got[0] != "WORK-0036" {
		t.Errorf("FormerKeys() = %v, want [WORK-0036]", got)
	}
}

func TestChainedMigrationKeepsEveryFormerKey(t *testing.T) {
	// WORK -> BACK -> PROJ. The oldest key must still resolve; a single-valued
	// field would have dropped it on the second migration.
	b := migratedBacklog(t, map[string][]string{
		"PROJ-0036-migrated-twice": {"PROJ-0036", "WORK-0036", "BACK-0036"},
	})

	for _, ref := range []string{"WORK-0036", "BACK-0036", "PROJ-0036"} {
		it, err := Resolve(b, ref)
		if err != nil {
			t.Errorf("Resolve(%q): %v", ref, err)
			continue
		}
		if it.Key() != "PROJ-0036" {
			t.Errorf("Resolve(%q) found %s, want PROJ-0036", ref, it.Key())
		}
	}
}

func TestCurrentKeyBeatsAnotherRecordsFormerKey(t *testing.T) {
	// The transient state a migration passes through: one record already holds
	// WORK-0050 as its live key while another lists it as former. Allocation
	// never produces this outside a migration, and resolution must not be
	// ambiguous while it lasts.
	b := migratedBacklog(t, map[string][]string{
		"WORK-0050-holds-it-now":    {"WORK-0050"},
		"BACK-0050-used-to-hold-it": {"BACK-0050", "WORK-0050"},
	})

	it, err := Resolve(b, "WORK-0050")
	if err != nil {
		t.Fatalf("Resolve(WORK-0050): %v", err)
	}
	if it.Key() != "WORK-0050" {
		t.Errorf("Resolve(WORK-0050) found %s, want the record that holds it now", it.Key())
	}
}

func TestHeldKeyCoversCurrentAndFormer(t *testing.T) {
	// What allocation asks: is this key taken by anybody, ever. The three
	// cases it has to separate are current, former, and neither.
	b := migratedBacklog(t, map[string][]string{
		"BACK-0036-a-migrated-record": {"BACK-0036", "WORK-0036"},
	})
	items, _, err := List(b, Filter{})
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	it := items[0]

	for _, c := range []struct {
		ref    string
		held   bool
		former bool
	}{
		{"BACK-0036", true, false},  // current
		{"WORK-0036", true, true},   // former
		{"back-36", true, false},    // current, sloppily spelled
		{"work-36", true, true},     // former, sloppily spelled
		{"BACK-0099", false, false}, // neither
	} {
		if got := it.HeldKey(c.ref); got != c.held {
			t.Errorf("HeldKey(%q) = %v, want %v", c.ref, got, c.held)
		}
		if got := it.HeldFormerKey(c.ref); got != c.former {
			t.Errorf("HeldFormerKey(%q) = %v, want %v", c.ref, got, c.former)
		}
	}
}

func TestRecordWithoutFormerKeysHasNone(t *testing.T) {
	// Nearly every record. An absent field is the ordinary case and must not
	// read as a fault or as a one-entry list holding the empty string.
	b := migratedBacklog(t, map[string][]string{
		"WORK-0040-never-migrated": {"WORK-0040"},
	})
	items, _, err := List(b, Filter{})
	if err != nil {
		t.Fatal(err)
	}
	if got := items[0].FormerKeys(); len(got) != 0 {
		t.Errorf("FormerKeys() = %v, want empty", got)
	}
	if items[0].HeldFormerKey("WORK-0040") {
		t.Error("a record's current key must not read as a former one")
	}
}

func TestWorkItemFlagAcceptsAFormerKey(t *testing.T) {
	// `--work-item` matches on the directory name, and a migration renames it,
	// old key survives only in former_keys. Without this the flag would be the
	// one door a migrated key could not open, and it is the door every child
	// record is created through.
	b := migratedBacklog(t, map[string][]string{
		"BACK-0036-a-migrated-record": {"BACK-0036", "WORK-0036"},
		"BACK-0040-never-migrated":    {"BACK-0040"},
	})

	for _, ref := range []string{"WORK-0036", "work-36", "BACK-0036", "back-36"} {
		dir, err := ResolveWorkItemDir(b, ref)
		if err != nil {
			t.Errorf("ResolveWorkItemDir(%q): %v", ref, err)
			continue
		}
		if dir != "BACK-0036-a-migrated-record" {
			t.Errorf("ResolveWorkItemDir(%q) = %q, want BACK-0036-a-migrated-record", ref, dir)
		}
	}
}

func TestWorkItemFlagPrefersTheRecordHoldingTheKeyNow(t *testing.T) {
	b := migratedBacklog(t, map[string][]string{
		"WORK-0050-holds-it-now":    {"WORK-0050"},
		"BACK-0050-used-to-hold-it": {"BACK-0050", "WORK-0050"},
	})

	dir, err := ResolveWorkItemDir(b, "WORK-0050")
	if err != nil {
		t.Fatalf("ResolveWorkItemDir(WORK-0050): %v", err)
	}
	if dir != "WORK-0050-holds-it-now" {
		t.Errorf("ResolveWorkItemDir(WORK-0050) = %q, want the record holding it now", dir)
	}
}
