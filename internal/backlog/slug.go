// Package backlog owns the layout and the creation of records: which file a
// unit lives in, and what it contains when it is first written.
package backlog

import (
	"strings"
	"unicode"
)

// Slugify turns a title into a filename.
//
// Filenames are slugs derived from titles, in kebab-case, so a path is
// meaningful without a lookup. Numeric identifiers are deliberately not used:
// they need an allocator, they collide across branches, and they tell a reader
// nothing (docs/SPEC.md §7.4).
func Slugify(title string) string {
	var b strings.Builder
	lastDash := true // leading dashes are suppressed

	for _, r := range strings.ToLower(strings.TrimSpace(title)) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			lastDash = false
		case !lastDash:
			// Any run of punctuation or space collapses to one dash, so
			// "retry queue: drain it!" and "retry queue - drain it" agree.
			b.WriteRune('-')
			lastDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}
