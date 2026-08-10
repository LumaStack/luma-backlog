package backlog

import (
	"strings"
)

// entryHeading marks the start of a journal entry.
const entryHeading = "## "

// AppendLine puts one line into today's entry, opening today's entry at the
// top when there is not one.
//
// Entries are newest first and are PREPENDED — nothing below is ever
// rewritten. Lines within a day run in the order they were written, because a
// day reads as a narrative even though the days do not.
//
// Tolerant of a hand-written journal, because most of them are: anything
// before the first entry heading is left exactly as it is, whatever it
// contains.
func AppendLine(journal, today, line string) string {
	if journal == "" {
		journal = "# Journal\n\n"
	}
	if !strings.HasSuffix(journal, "\n") {
		journal += "\n"
	}

	head, rest := splitAtFirstEntry(journal)

	if startsWithTodaysEntry(rest, today) {
		nextEntry := indexOfEntry(rest[len(entryHeading):])
		if nextEntry < 0 {
			return head + strings.TrimRight(rest, "\n") + "\n" + line + "\n"
		}
		cut := len(entryHeading) + nextEntry
		return head + strings.TrimRight(rest[:cut], "\n") + "\n" + line + "\n\n" + rest[cut:]
	}

	entry := entryHeading + "▶ " + today + "\n\n" + line + "\n"
	if rest == "" {
		return head + entry
	}
	return head + entry + "\n" + rest
}

// splitAtFirstEntry divides the introduction from the entries.
func splitAtFirstEntry(journal string) (head, rest string) {
	if i := indexOfEntry(journal); i >= 0 {
		return journal[:i], journal[i:]
	}
	// No entries yet: everything is introduction, and the first entry goes
	// after it with a blank line between.
	if !strings.HasSuffix(journal, "\n\n") {
		journal += "\n"
	}
	return journal, ""
}

// indexOfEntry finds the first entry heading at the start of a line.
func indexOfEntry(s string) int {
	if strings.HasPrefix(s, entryHeading) {
		return 0
	}
	if i := strings.Index(s, "\n"+entryHeading); i >= 0 {
		return i + 1
	}
	return -1
}

func startsWithTodaysEntry(rest, today string) bool {
	if !strings.HasPrefix(rest, entryHeading) {
		return false
	}
	end := strings.IndexByte(rest, '\n')
	if end < 0 {
		end = len(rest)
	}
	return strings.Contains(rest[:end], today)
}
