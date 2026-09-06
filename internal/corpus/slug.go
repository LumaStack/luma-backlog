// Package corpus is the body of records and what they mean: which file a unit
// lives in, what it contains when first written, how one is found, and what can
// be counted across them.
//
// It is the plural of internal/record, which knows one record's bytes and
// nothing about the others. Named for the corpus rather than for the backlog
// because everything here is the backlog — the word said nothing.
package corpus

import (
	"strings"
	"unicode"
)

// Slugify turns a title into a filename.
//
// Filenames are slugs derived from titles, in kebab-case, so a path is
// meaningful without a lookup. Numeric identifiers are deliberately not used:
// they need an allocator, they collide across branches, and they tell a reader
// nothing (docs/spec.md §7.4).
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
