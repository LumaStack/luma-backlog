package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/lumastack/luma-backlog/internal/backlog"
)

// The JSON shapes below are published contract (docs/SPEC.md §9.3), which is
// why they are explicit structs rather than whatever a record happens to hold.
// Marshalling an internal type would make every refactor a breaking change
// without anyone noticing.

// itemJSON is one record in a listing.
type itemJSON struct {
	Path        string `json:"path"`
	Type        string `json:"type"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Deliverable string `json:"deliverable,omitempty"`
}

// recordJSON is one record in full. Fields carries the frontmatter as written,
// including keys this tool knows nothing about — dropping them here would
// quietly hide another system's state from anything reading our output.
type recordJSON struct {
	itemJSON
	Fields map[string]any `json:"fields"`
	Body   string         `json:"body"`
}

func toItemJSON(i backlog.Item, defaultStatus string) itemJSON {
	return itemJSON{
		Path:        i.Path,
		Type:        i.Type(),
		Slug:        i.Slug(),
		Title:       i.Title(),
		Status:      i.Status(defaultStatus),
		Deliverable: i.Deliverable,
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
