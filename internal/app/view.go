package app

import (
	"github.com/lumastack/luma-backlog/internal/corpus"
)

// View is one record as a surface needs to show it.
//
// A view rather than the record itself, so an adapter never holds an engine
// type. That is what lets the containment rule be a build-time fact: a surface
// that cannot name the engine cannot reach past this layer to it
// (.luma/records/decisions/ADR-0004).
type View struct {
	Path string
	Type string
	// Key is the handle somebody quotes — WORK-0002. Empty where there is
	// none: only a work item carries one, and a record written before keys
	// existed has none either.
	Key string
	// Name is the two joined — WORK-0002-lint-the-corpus.
	Name  string
	Slug  string
	Title string
	// Status is empty when the record's type declares no workflow status.
	Status   string
	WorkItem string
}

// Record is a view with everything a single-record read needs.
type Record struct {
	View
	// Hash identifies the content that was read. Pass it back to a set with
	// IfUnchanged and a write that would clobber somebody else's change is
	// refused rather than applied (docs/spec.md §6.3).
	Hash string
	// Fields is the frontmatter as written, including keys this tool knows
	// nothing about — dropping them would quietly hide another system's state.
	Fields map[string]any
	// Order is the frontmatter's key order, so a rendering can follow the file
	// rather than sorting and telling a reader something the file does not.
	Order []string
	// Raw is each field as the file holds it. Kept beside Fields because the
	// two renderings need different things — a structured consumer wants the
	// decoded value, and a person reading a terminal wants what is actually
	// written on the line.
	Raw  map[string]string
	Body string
}

func (s *Session) view(it corpus.Item) View {
	return View{
		Path:     it.Path,
		Type:     it.Type(),
		Key:      it.Key(),
		Name:     it.Name(),
		Slug:     it.Slug(),
		Title:    it.Title(),
		Status:   it.Status(s.Config.DefaultStatusFor(it.Type())),
		WorkItem: it.WorkItem,
	}
}

func (s *Session) record(it corpus.Item) (Record, error) {
	fields := map[string]any{}
	order := it.Record.Keys()
	for _, k := range order {
		var v any
		if node := it.Record.Node(k); node != nil {
			if err := node.Decode(&v); err != nil {
				return Record{}, FailureError("decoding %s: %w", k, err)
			}
		}
		fields[k] = v
	}
	raw := map[string]string{}
	for _, k := range order {
		if v, ok := it.Record.Get(k); ok {
			raw[k] = v
		}
	}
	return Record{
		View:   s.view(it),
		Hash:   it.Hash(),
		Fields: fields,
		Order:  order,
		Raw:    raw,
		Body:   it.Record.Body(),
	}, nil
}

// Units are the record types a caller may name.
var Units = corpus.Units

// Skip is a record that could not be read.
type Skip struct {
	Path string
	Err  error
}

// Duplicate is one key held by more than one record.
type Duplicate struct {
	Key   string
	Paths []string
}
