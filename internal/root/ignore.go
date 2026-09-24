package root

import (
	"path"
	"strings"
)

// DefaultIgnores are skipped by every walk, without anybody asking.
//
// **Two kinds, and they are excluded for different reasons.** `.git` and
// `.luma/bundles` hold content that is not the project's to edit --- git's own
// storage, and vendored copies of somebody else's published bundles, where an
// edit is drift that the next adoption reverts or conflicts with.
//
// The rest are where each ecosystem puts code it did not write or output it
// regenerates. **Listed rather than detected**, because there is no reliable
// signal for "somebody else owns this file" --- and named in `--help` so a
// reader can see what was skipped rather than wonder.
//
// **This list will be wrong for somebody**, which is what `--ignore` is for.
var DefaultIgnores = []string{
	".git/**",
	".luma/bundles/**",
	"node_modules/**", // javascript
	"vendor/**",       // go, php
	"target/**",       // rust, java
	"dist/**", "build/**", "out/**",
	".venv/**", "venv/**", "__pycache__/**", // python
	".gradle/**", ".terraform/**",
}

// Ignore decides which paths a walk skips.
type Ignore struct{ patterns []string }

// Patterns is every pattern in effect, defaults first, so output can list what
// was skipped rather than leave somebody wondering.
func (ig Ignore) Patterns() []string { return append([]string{}, ig.patterns...) }

// NewIgnore combines the defaults with whatever the caller added.
func NewIgnore(extra ...string) Ignore {
	p := append([]string{}, DefaultIgnores...)
	return Ignore{patterns: append(p, extra...)}
}

// Match reports whether a slash-separated path is ignored.
//
// `**` matches any number of path segments, `*` matches within one, and a
// pattern with no wildcard matches that path or anything under it --- so
// `--ignore vendor` does what somebody typing it meant, without requiring them
// to know the syntax.
func (ig Ignore) Match(rel string) bool {
	rel = strings.TrimPrefix(path.Clean(rel), "./")
	for _, pat := range ig.patterns {
		if matchPattern(pat, rel) {
			return true
		}
	}
	return false
}

func matchPattern(pat, rel string) bool {
	if !strings.ContainsAny(pat, "*?[") {
		pat = strings.TrimSuffix(pat, "/")
		return rel == pat || strings.HasPrefix(rel, pat+"/")
	}
	return segMatch(strings.Split(pat, "/"), strings.Split(rel, "/"))
}

// segMatch walks pattern segments against path segments, with `**` able to
// consume any number of them.
func segMatch(pat, seg []string) bool {
	if len(pat) == 0 {
		return len(seg) == 0
	}
	if pat[0] == "**" {
		// `a/**` matches a itself as well as everything under it, which is
		// what somebody writing it expects.
		if len(pat) == 1 {
			return true
		}
		for i := 0; i <= len(seg); i++ {
			if segMatch(pat[1:], seg[i:]) {
				return true
			}
		}
		return false
	}
	if len(seg) == 0 {
		return false
	}
	if ok, err := path.Match(pat[0], seg[0]); err != nil || !ok {
		return false
	}
	return segMatch(pat[1:], seg[1:])
}
