package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/lumastack/luma-backlog/internal/backlog"
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
	// Ref is the two joined — WORK-00002-lint-the-corpus — which is how a
	// record is written and said. Key and slug are kept beside it so a
	// consumer can use either half without parsing this one apart.
	Ref   string `json:"ref"`
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

func toItemJSON(i backlog.Item, defaultStatus string) itemJSON {
	return itemJSON{
		Path:     i.Path,
		Type:     i.Type(),
		Key:      i.Key(),
		Ref:      i.Ref(),
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
