package migrate

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/lumastack/luma-backlog/internal/corpus"
	"github.com/lumastack/luma-backlog/internal/record"
	"github.com/lumastack/luma-backlog/internal/root"
)

// Options is how a key migration was asked for.
type Options struct {
	// Target is the prefix every work item should end up under.
	Target string
	// DryRun computes everything and writes nothing.
	DryRun bool
	// Renumber gives a record blocked by a collision the next available
	// number instead of leaving it alone.
	Renumber bool
	// IncludeBareKeys also rewrites keys written without their slug. Off by
	// default because a bare key may name another project's work item.
	IncludeBareKeys bool
}

// FileChange is one file and how many replacements it took.
type FileChange struct {
	Path    string
	Changes int
}

// BareKeyFile is one file and the old keys still written in it.
type BareKeyFile struct {
	Path string
	Keys []string
}

// Result is what a migration did, or would have done.
type Result struct {
	Renames        []corpus.KeyRename
	Collisions     []corpus.KeyCollision
	AlreadyCorrect int
	Files          []FileChange
	// BareKeys are the keys left in place. Empty when IncludeBareKeys
	// rewrote them.
	BareKeys []BareKeyFile
	DryRun   bool
}

// Keys moves every work item to the target prefix and repoints what named it.
//
// **The order is forced and each step depends on the one before.** Records are
// rewritten first, so their own frontmatter carries the new key before anything
// reads it. Then names are rewritten across the repository, which catches the
// `work_item` link on every child and every mention in prose, documentation and
// source. Directories move last, because after they move the old paths are gone
// and a half-finished run would have nothing to find.
func Keys(projectRoot string, b *root.Backlog, opts Options) (*Result, error) {
	plan, err := corpus.PlanKeyMigration(b, opts.Target)
	if err != nil {
		return nil, err
	}

	res := &Result{
		Renames:        plan.Renames,
		Collisions:     plan.Collisions,
		AlreadyCorrect: plan.AlreadyCorrect,
		DryRun:         opts.DryRun,
	}

	if opts.Renumber && len(plan.Collisions) > 0 {
		extra, rerr := renumberBlocked(b, plan.Collisions, opts.Target)
		if rerr != nil {
			return nil, rerr
		}
		res.Renames = append(res.Renames, extra...)
		res.Collisions = nil
	}

	p, err := openRepository(projectRoot)
	if err != nil {
		return nil, err
	}
	defer p.Close()

	if !opts.DryRun {
		if err := stampRecords(p, res.Renames); err != nil {
			return nil, err
		}
	}

	files, bare, err := rewriteTree(p, res.Renames, opts)
	if err != nil {
		return nil, err
	}
	res.Files, res.BareKeys = files, bare

	if !opts.DryRun {
		if err := moveDirectories(p, res.Renames); err != nil {
			return nil, err
		}
	}
	return res, nil
}

// renumberBlocked gives each collided record a key nobody has ever held.
//
// Allocated one at a time against the corpus rather than from a single
// starting number, because two blocked records in one run must not be handed
// the same key --- and the allocator is the only thing that knows what is
// taken.
func renumberBlocked(b *root.Backlog, blocked []corpus.KeyCollision, target string) ([]corpus.KeyRename, error) {
	var out []corpus.KeyRename
	taken := map[string]bool{}
	for _, c := range blocked {
		next, err := corpus.NextAvailableKey(b, target)
		for taken[next] && err == nil {
			// Already promised to an earlier record in this same run. The
			// corpus does not know that yet, because nothing is written until
			// the plan is complete.
			_, n, _ := corpus.ParseKey(next)
			next = corpus.FormatKeyAs(target, n+1)
		}
		if err != nil {
			return nil, err
		}
		taken[next] = true
		out = append(out, corpus.KeyRename{
			OldKey:  keyOfName(c.Record),
			NewKey:  next,
			OldName: c.Record,
			NewName: next + "-" + corpus.SlugOf(c.Record),
			Dir:     c.Record,
		})
	}
	return out, nil
}

// keyOfName reads the key half of a directory name.
func keyOfName(name string) string {
	if i := strings.Index(name, "-"); i > 0 {
		if j := strings.Index(name[i+1:], "-"); j > 0 {
			return name[:i+1+j]
		}
	}
	return name
}

// stampRecords writes each migrating work item's new key and records the old
// one, so the key it used to answer to keeps resolving.
func stampRecords(p *root.Project, renames []corpus.KeyRename) error {
	for _, r := range renames {
		rel := path.Join(root.Dir, corpus.WorkItemPath(r.Dir))
		data, err := p.ReadFile(rel)
		if err != nil {
			return fmt.Errorf("reading %s: %w", rel, err)
		}
		rec, err := record.Parse(data)
		if err != nil {
			return fmt.Errorf("parsing %s: %w", rel, err)
		}
		former := rec.List("former_keys")
		former = append(former, r.OldKey)
		rec.Set("key", r.NewKey)
		if err := rec.SetRaw("former_keys", yamlList(former)); err != nil {
			return fmt.Errorf("writing former_keys on %s: %w", rel, err)
		}
		out, err := rec.Bytes()
		if err != nil {
			return fmt.Errorf("rendering %s: %w", rel, err)
		}
		if err := p.WriteFile(rel, out, 0o644); err != nil {
			return err
		}
	}
	return nil
}

// yamlList renders a flow sequence, the shape every other list field uses.
func yamlList(items []string) string {
	quoted := make([]string, 0, len(items))
	for _, s := range items {
		quoted = append(quoted, `"`+s+`"`)
	}
	return "[" + strings.Join(quoted, ", ") + "]"
}

// rewriteTree repoints every name across the repository, and collects or
// rewrites the bare keys depending on what was asked for.
func rewriteTree(p *root.Project, renames []corpus.KeyRename, opts Options) ([]FileChange, []BareKeyFile, error) {
	var files []FileChange
	var bare []BareKeyFile
	old := map[string]string{}
	for _, r := range renames {
		old[r.OldKey] = r.NewKey
	}

	// A migrating record's own `key:` field is rewritten by stampRecords, which
	// a dry run does not do. Without accounting for that here, a dry run reads
	// its own un-stamped input and reports each record's current key as one
	// "remaining" --- over-counting by exactly the number of records moved, and
	// naming files that need nothing. **A dry run that does not predict the run
	// is worse than no dry run**, because it is believed.
	stamped := map[string]string{}
	for _, r := range renames {
		stamped[path.Join(root.Dir, corpus.WorkItemPath(r.Dir))] = r.OldKey + "\x00" + r.NewKey
	}

	err := p.WalkText(func(rel string, data []byte) error {
		text := string(data)
		if pair, ok := stamped[rel]; ok {
			k := strings.SplitN(pair, "\x00", 2)
			text = strings.Replace(text, "key: "+k[0], "key: "+k[1], 1)
		}
		next, n := corpus.RewriteNamesIn(text, renames)

		found := bareKeysIn(next, old)
		if opts.IncludeBareKeys && len(found) > 0 {
			next, n = rewriteBareKeys(next, old, n)
		} else if len(found) > 0 {
			bare = append(bare, BareKeyFile{Path: rel, Keys: found})
		}

		if n == 0 {
			return nil
		}
		files = append(files, FileChange{Path: rel, Changes: n})
		if opts.DryRun {
			return nil
		}
		// On a real run the file already carries the stamped key, so writing
		// `next` is writing what stampRecords produced plus the name rewrite.
		return p.WriteFile(rel, []byte(next), 0o644)
	})
	sort.Slice(bare, func(i, j int) bool { return bare[i].Path < bare[j].Path })
	return files, bare, err
}

// bareKeysIn reports which migrated keys appear without their slug.
func bareKeysIn(text string, old map[string]string) []string {
	seen := map[string]bool{}
	var out []string
	for k := range old {
		for _, at := range occurrences(text, k) {
			end := at + len(k)
			if end < len(text) && text[end] == '-' {
				continue // part of a full name, already rewritten
			}
			if !seen[k] {
				seen[k] = true
				out = append(out, k)
			}
			break
		}
	}
	sort.Strings(out)
	return out
}

// rewriteBareKeys replaces migrated keys written without their slug.
func rewriteBareKeys(text string, old map[string]string, n int) (string, int) {
	for k, v := range old {
		var b strings.Builder
		from := 0
		for _, at := range occurrences(text, k) {
			if at < from {
				continue
			}
			end := at + len(k)
			if end < len(text) && text[end] == '-' {
				continue
			}
			b.WriteString(text[from:at])
			b.WriteString(v)
			from = end
			n++
		}
		b.WriteString(text[from:])
		text = b.String()
	}
	return text, n
}

// occurrences lists where a key appears, at a word boundary on the left so a
// longer key ending in these characters is not matched.
func occurrences(text, key string) []int {
	var out []int
	for i := 0; ; {
		j := strings.Index(text[i:], key)
		if j < 0 {
			return out
		}
		at := i + j
		if at == 0 || !isKeyChar(text[at-1]) {
			out = append(out, at)
		}
		i = at + len(key)
	}
}

func isKeyChar(c byte) bool {
	return c == '-' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// moveDirectories renames each work item's directory, last, because after this
// the old paths are gone.
func moveDirectories(p *root.Project, renames []corpus.KeyRename) error {
	for _, r := range renames {
		from := path.Join(root.Dir, corpus.WorkItemsDir, r.Dir)
		to := path.Join(root.Dir, corpus.WorkItemsDir, r.NewName)
		if err := p.Rename(from, to); err != nil {
			return fmt.Errorf("moving %s to %s: %w", from, to, err)
		}
	}
	return nil
}
