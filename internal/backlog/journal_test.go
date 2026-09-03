package backlog

import (
	"strings"
	"testing"
)

const existing = `# Journal — payments

> The work item's memory. Newest entry first.

---

## ▶ 2026-08-08 — yesterday

Something learned yesterday.

---

## 2026-08-07 — before that

Older still.
`

func TestAppendOpensTodaysEntryAtTheTop(t *testing.T) {
	got := AppendLine(existing, "2026-08-09", "A thing learned today.")

	first := firstEntryLine(got)
	if !strings.Contains(first, "2026-08-09") {
		t.Errorf("today's entry is not first; first entry is %q", first)
	}
	// Nothing below may be rewritten.
	if !strings.Contains(got, "Something learned yesterday.") ||
		!strings.Contains(got, "Older still.") {
		t.Errorf("earlier entries were lost:\n%s", got)
	}
	if !strings.Contains(got, "# Journal — payments") {
		t.Error("the introduction was lost")
	}
}

func TestAppendAddsToTodaysExistingEntry(t *testing.T) {
	once := AppendLine(existing, "2026-08-09", "First thing.")
	twice := AppendLine(once, "2026-08-09", "Second thing.")

	if strings.Count(twice, "2026-08-09") != 1 {
		t.Errorf("a second entry was opened for the same day:\n%s", twice)
	}
	// Within a day, lines run in the order they were written: a day reads as
	// a narrative even though the days do not.
	if strings.Index(twice, "First thing.") > strings.Index(twice, "Second thing.") {
		t.Errorf("lines within a day are out of order:\n%s", twice)
	}
	if !strings.Contains(twice, "Something learned yesterday.") {
		t.Error("earlier entries were lost")
	}
}

func TestAppendToAJournalWithNoEntriesYet(t *testing.T) {
	head := "# Journal — payments\n\n> A blurb.\n\n---\n"
	got := AppendLine(head, "2026-08-09", "The first thing.")

	if !strings.HasPrefix(got, "# Journal — payments") {
		t.Errorf("the introduction moved:\n%s", got)
	}
	if !strings.Contains(got, "## ▶ 2026-08-09") || !strings.Contains(got, "The first thing.") {
		t.Errorf("the entry was not written:\n%s", got)
	}
}

func TestAppendToAnEmptyFile(t *testing.T) {
	got := AppendLine("", "2026-08-09", "The very first thing.")
	if !strings.Contains(got, "The very first thing.") {
		t.Errorf("nothing written:\n%s", got)
	}
}

func TestAppendLeavesAHandWrittenIntroductionAlone(t *testing.T) {
	// Most journals are hand-written, and an introduction can contain
	// anything — including text that looks structural.
	head := "# Journal\n\nSome prose.\n\n### A sub-heading\n\nMore prose.\n\n---\n"
	got := AppendLine(head, "2026-08-09", "A line.")
	if !strings.HasPrefix(got, head) {
		t.Errorf("the introduction was altered:\n--- got ---\n%s", got)
	}
}

func TestAppendDoesNotDisturbLaterEntries(t *testing.T) {
	got := AppendLine(existing, "2026-08-09", "A line.")
	tail := existing[strings.Index(existing, "## ▶ 2026-08-08"):]
	if !strings.Contains(got, tail) {
		t.Errorf("the existing entries were reflowed rather than left alone:\n%s", got)
	}
}

func firstEntryLine(s string) string {
	i := strings.Index(s, "\n## ")
	if i < 0 {
		return ""
	}
	rest := s[i+1:]
	if j := strings.IndexByte(rest, '\n'); j >= 0 {
		return rest[:j]
	}
	return rest
}
