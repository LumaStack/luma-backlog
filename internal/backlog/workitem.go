package backlog

import (
	"path"
	"regexp"
	"strings"

	"github.com/lumastack/luma-backlog/internal/root"
)

// dirPattern splits a work item directory into its key and its slug.
var dirPattern = regexp.MustCompile(`^([A-Z]+-\d+)-(.+)$`)

// SlugOf returns the slug half of a work item directory name, or the name
// itself when it carries no key — a directory written before keys were in
// paths still has to resolve.
func SlugOf(dirName string) string {
	if m := dirPattern.FindStringSubmatch(dirName); m != nil {
		return m[2]
	}
	return dirName
}

// findWorkItemBySlug returns the directory of a work item whose slug half
// matches, if one exists.
//
// The key is in the path now, so the path is no longer derivable from the title
// alone and the existence check cannot be a path lookup. Without this, asking
// twice for the same work item would allocate a second key and create a second
// directory for one piece of work — the trap the decision numbering had to be
// rescued from, arriving by a different route.
func findWorkItemBySlug(b *root.Backlog, slug string) (string, error) {
	var found string
	err := b.Walk(func(rel string) error {
		if path.Base(rel) != "index.md" {
			return nil
		}
		dir := path.Base(path.Dir(rel))
		if !strings.HasPrefix(path.Clean(rel), path.Join(BundleDir, "work-items")+"/") {
			return nil
		}
		if SlugOf(dir) == slug {
			found = dir
		}
		return nil
	})
	return found, err
}

// matchesWorkItem reports whether a filter names this work item.
//
// Three forms reach the same work: the directory as written
// (WORK-00001-payments-v2), the slug half alone (payments-v2), and the key
// (WORK-00001, case-insensitively). Requiring the long form everywhere would
// make the key a tax rather than a handle, and the short forms were what people
// typed before the key existed.
func matchesWorkItem(dirName, filter string) bool {
	if dirName == filter {
		return true
	}
	if SlugOf(dirName) == filter {
		return true
	}
	if m := dirPattern.FindStringSubmatch(dirName); m != nil {
		return strings.EqualFold(m[1], filter)
	}
	return false
}

// ResolveWorkItemDir turns a reference to a work item into its directory name.
//
// A caller names a work item however they have it to hand — the directory, the
// slug half, or the key — and everything downstream needs the one true
// directory, because that is where children are written and what their
// work_item link has to say. Resolving once here keeps every caller from having
// to know the key is in the path.
//
// An unrecognized reference is returned unchanged, so a work item that does not
// exist yet fails where it already failed rather than somewhere new.
func ResolveWorkItemDir(b *root.Backlog, ref string) (string, error) {
	if ref == "" {
		return "", nil
	}
	items, _, err := List(b, Filter{Unit: WorkItem})
	if err != nil {
		return "", err
	}
	for _, it := range items {
		if matchesWorkItem(it.WorkItem, ref) {
			return it.WorkItem, nil
		}
	}
	return ref, nil
}
