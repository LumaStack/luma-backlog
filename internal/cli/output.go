package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"github.com/lumastack/luma-backlog/internal/root"
)

// The JSON shapes below are published contract (docs/spec.md §9.3), which is
// why they are explicit structs rather than whatever a record happens to hold.
// Marshalling an internal type would make every refactor a breaking change
// without anyone noticing.

// itemJSON is one record in a listing.
type itemJSON struct {
	Path string `json:"path"`
	Type string `json:"type"`
	// Key is the handle somebody quotes — WORK-00002. Omitted where there is
	// none: only a work item carries one, and a record written before keys
	// existed has none either.
	Key string `json:"key,omitempty"`
	// Name is the two joined — WORK-0002-lint-the-corpus — which is what the
	// record is called and what its directory is named. Key and slug are kept
	// beside it so a consumer can use either half without parsing this apart.
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
	// Status is omitted rather than emptied when the record's type declares
	// no workflow status. A consumer can then tell "no lifecycle" from a
	// lifecycle whose value happens to be blank.
	Status   string `json:"status,omitempty"`
	WorkItem string `json:"work_item,omitempty"`
}

// recordJSON is one record in full. Fields carries the frontmatter as written,
// including keys this tool knows nothing about — dropping them here would
// quietly hide another system's state from anything reading our output.
type recordJSON struct {
	itemJSON
	// Hash identifies the content that was read. Pass it back to `set
	// --if-unchanged` and a write that would clobber someone else's change is
	// refused rather than applied (docs/spec.md §6.3).
	Hash   string         `json:"hash"`
	Fields map[string]any `json:"fields"`
	Body   string         `json:"body"`
}

// reportSkipped names every record that could not be read.
//
// It writes nothing when there is nothing to say, so ordinary runs are as quiet
// as they were. The path comes first because the path is what somebody has to
// go and open.
func reportSkipped(w io.Writer, skipped []backlog.Skip) {
	for _, s := range skipped {
		fmt.Fprintf(w, "luma-backlog: skipped %s: %v\n", s.Path, s.Err)
	}
}

// reportDuplicateKeys names every key held by more than one record.
//
// Run over the whole work item set rather than over whatever a caller happened
// to ask for: a filtered listing would miss a duplicate outside its filter, and
// a check that only fires when you were already looking in the right place is
// not a check.
//
// Silent when there is nothing to say. A warning that appears on ordinary runs
// is one people learn to scroll past, and this one has to survive being ignored
// for months before it matters once.
func reportDuplicateKeys(w io.Writer, b *root.Backlog) {
	items, _, err := backlog.List(b, backlog.Filter{Unit: backlog.WorkItem})
	if err != nil {
		return
	}
	for _, d := range backlog.Duplicates(items) {
		fmt.Fprintf(w, "luma-backlog: %s is held by %d records:\n", d.Key, len(d.Paths))
		for _, p := range d.Paths {
			fmt.Fprintf(w, "  %s\n", p)
		}
		fmt.Fprintf(w, "A key is meant to name one record, and every citation of this one is ambiguous.\n")
	}
}

func toItemJSON(i backlog.Item, defaultStatus string) itemJSON {
	return itemJSON{
		Path:     i.Path,
		Type:     i.Type(),
		Key:      i.Key(),
		Name:     i.Name(),
		Slug:     i.Slug(),
		Title:    i.Title(),
		Status:   i.Status(defaultStatus),
		WorkItem: i.WorkItem,
	}
}

func toRecordJSON(i backlog.Item, defaultStatus string) (recordJSON, error) {
	fields := map[string]any{}
	for _, k := range i.Record.Keys() {
		var v any
		if node := i.Record.Node(k); node != nil {
			if err := node.Decode(&v); err != nil {
				return recordJSON{}, fmt.Errorf("decoding %s: %w", k, err)
			}
		}
		fields[k] = v
	}
	return recordJSON{
		itemJSON: toItemJSON(i, defaultStatus),
		Hash:     i.Hash(),
		Fields:   fields,
		Body:     i.Record.Body(),
	}, nil
}

// writeJSON emits indented JSON with a trailing newline, so output is
// diffable, greppable, and pleasant in a terminal.
func writeJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
