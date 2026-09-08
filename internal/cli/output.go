package cli

import (
	"encoding/json"
	"io"

	"github.com/lumastack/luma-backlog/internal/app"
)

// The JSON shapes below are published contract (docs/spec.md §9.3), which is
// why they are explicit structs rather than whatever a record happens to hold.
// Marshalling an internal type would make every refactor a breaking change
// without anyone noticing.

// itemJSON is one record in a listing.
// completionJSON is the count a board puts on a card and a listing can show
// without provoking a refusal to find out (docs/spec.md §9.3).
type completionJSON struct {
	Proven  int `json:"proven"`
	Live    int `json:"live"`
	Retired int `json:"retired,omitempty"`
	Skipped int `json:"skipped,omitempty"`
}

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
	// Created and Modified are omitted where a record carries no stamp, rather
	// than emitted empty --- absent says "this record has none", where
	// {"by":"","at":""} says the tool read one and found nothing in it.
	// Completion is how many live outcomes are proven, out of how many there
	// are. Omitted on anything that is not a work item — absent says "not
	// counted", where a zeroed object would say "counted, and none".
	Completion *completionJSON `json:"completion,omitempty"`
	Created    *stampJSON      `json:"created,omitempty"`
	Modified   *stampJSON      `json:"modified,omitempty"`
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

func toItemJSON(v app.View) itemJSON {
	return itemJSON{
		Path:     v.Path,
		Type:     v.Type,
		Key:      v.Key,
		Name:     v.Name,
		Slug:     v.Slug,
		Title:    v.Title,
		Status:   v.Status,
		WorkItem: v.WorkItem,
		Created:  toStampJSON(v.Created),
		Modified: toStampJSON(v.Modified),
		Completion: func() *completionJSON {
			if v.Completion == nil {
				return nil
			}
			return &completionJSON{
				Proven:  v.Completion.Proven,
				Live:    v.Completion.Live,
				Retired: v.Completion.Retired,
				Skipped: v.Completion.Skipped,
			}
		}(),
	}
}

func toRecordJSON(r app.Record) recordJSON {
	return recordJSON{
		itemJSON: toItemJSON(r.View),
		Hash:     r.Hash,
		Fields:   r.Fields,
		Body:     r.Body,
	}
}

// writeJSON emits indented JSON with a trailing newline, so output is
// diffable, greppable, and pleasant in a terminal.
func writeJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// nodeJSON is a work item with what hangs off it. Children are added rather
// than the shape being changed, so a consumer reading a flat listing is
// unaffected (docs/spec.md §9.9).
type nodeJSON struct {
	itemJSON
	Children []itemJSON `json:"children,omitempty"`
}

// stampJSON is who did something and when.
type stampJSON struct {
	By string `json:"by"`
	At string `json:"at"`
}

func toStampJSON(s app.Stamp) *stampJSON {
	if s.By == "" && s.At == "" {
		return nil
	}
	return &stampJSON{By: s.By, At: s.At}
}
