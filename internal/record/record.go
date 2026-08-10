// Package record parses and serializes a record: YAML frontmatter followed by
// a markdown body.
//
// It touches no filesystem — parsing is pure, so it can be tested exhaustively
// without fixtures, and the one package allowed to do input and output stays
// small (docs/SPEC.md §9a.4).
//
// The governing rule is that **anything not understood survives untouched**.
// That is inherited from the format, and it is load-bearing architecture: it
// is how other systems store their own state in our records. Losing a key is
// silent data loss rather than a bug anyone reports.
package record

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

const fence = "---"

// ErrNoFrontmatter means the file did not open with a YAML block.
var ErrNoFrontmatter = errors.New("no frontmatter: a record starts with ---")

// Record is a parsed record.
//
// Frontmatter is held as a YAML node rather than a map, deliberately. A map
// loses key order, so every write would reshuffle the block and produce a diff
// touching lines nobody changed — which fights "changes land as clean diffs"
// and makes review useless.
type Record struct {
	frontmatter *yaml.Node // a mapping node, or nil when the block was empty
	body        string
}

// Parse reads a record. It is permissive by design: unknown keys, unknown
// types, and unresolved links are all fine, and none is a reason to reject.
func Parse(data []byte) (*Record, error) {
	text := strings.ReplaceAll(string(data), "\r\n", "\n")

	if !strings.HasPrefix(text, fence+"\n") {
		return nil, ErrNoFrontmatter
	}
	rest := text[len(fence)+1:]

	// The closing fence may be the very next line, giving empty frontmatter.
	// Searching only for "\n---" misses that case, because there is nothing
	// before it — a record with no fields yet then fails to parse at all.
	var raw string
	var end int
	if strings.HasPrefix(rest, fence) {
		raw, end = "", -1
	} else {
		end = strings.Index(rest, "\n"+fence)
		if end < 0 {
			return nil, fmt.Errorf("unterminated frontmatter: no closing %s", fence)
		}
		raw = rest[:end]
	}

	// After the closing fence comes the newline that ends the fence line,
	// then conventionally a blank line before the body. Strip both, and put
	// them back on the way out, so a round trip is byte-identical.
	body := rest[end+len(fence)+1:]
	body = strings.TrimPrefix(body, "\n")
	body = strings.TrimPrefix(body, "\n")

	r := &Record{body: body}
	if strings.TrimSpace(raw) != "" {
		var doc yaml.Node
		if err := yaml.Unmarshal([]byte(raw), &doc); err != nil {
			return nil, fmt.Errorf("parsing frontmatter: %w", err)
		}
		if len(doc.Content) > 0 {
			r.frontmatter = doc.Content[0]
		}
	}
	return r, nil
}

// Bytes serializes the record back to a file's contents.
func (r *Record) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fence + "\n")

	if r.frontmatter != nil && len(r.frontmatter.Content) > 0 {
		enc := yaml.NewEncoder(&buf)
		// Two spaces, matching the format's own examples and every record
		// written by hand so far. A change here rewrites every file.
		enc.SetIndent(2)
		if err := enc.Encode(r.frontmatter); err != nil {
			return nil, fmt.Errorf("encoding frontmatter: %w", err)
		}
		if err := enc.Close(); err != nil {
			return nil, fmt.Errorf("closing encoder: %w", err)
		}
	}

	buf.WriteString(fence + "\n")
	if r.body != "" {
		buf.WriteString("\n")
		buf.WriteString(r.body)
	}
	return buf.Bytes(), nil
}

// Body returns the markdown after the frontmatter.
func (r *Record) Body() string { return r.body }

// SetBody replaces the markdown after the frontmatter.
func (r *Record) SetBody(s string) { r.body = s }

// Get returns a scalar field's value.
func (r *Record) Get(key string) (string, bool) {
	if _, v := r.find(key); v != nil && v.Kind == yaml.ScalarNode {
		return v.Value, true
	}
	return "", false
}

// Has reports whether a key is present, whatever its shape.
func (r *Record) Has(key string) bool {
	_, v := r.find(key)
	return v != nil
}

// Set assigns a scalar field, replacing the value in place when the key
// already exists so its position — and everything around it — is unchanged.
// A new key is appended rather than inserted, for the same reason.
func (r *Record) Set(key, value string) {
	if r.frontmatter == nil {
		r.frontmatter = &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	}
	if _, v := r.find(key); v != nil {
		v.Kind = yaml.ScalarNode
		v.Tag = "!!str"
		v.Value = value
		v.Content = nil
		return
	}
	r.frontmatter.Content = append(r.frontmatter.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: key},
		&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: value},
	)
}

// SetRaw assigns a field from YAML source, so a value can be a map or a list
// rather than a string.
//
// Set writes a scalar, which is right for a title and wrong for a timestamped
// actor: quoting "{by: …, at: …}" produces a string that looks structured and
// is not, and every reader afterwards has to re-parse it by hand.
func (r *Record) SetRaw(key, yamlValue string) error {
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte(yamlValue), &doc); err != nil {
		return fmt.Errorf("parsing value for %s: %w", key, err)
	}
	if len(doc.Content) == 0 {
		return fmt.Errorf("value for %s is empty", key)
	}
	value := doc.Content[0]

	if r.frontmatter == nil {
		r.frontmatter = &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	}
	if _, existing := r.find(key); existing != nil {
		*existing = *value
		return nil
	}
	r.frontmatter.Content = append(r.frontmatter.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: key},
		value,
	)
	return nil
}

// Node returns the value node for a key, so a caller can decode a shape this
// package does not model. Nil when absent.
func (r *Record) Node(key string) *yaml.Node {
	_, v := r.find(key)
	return v
}

// Remove deletes a field. Reports whether it was there.
func (r *Record) Remove(key string) bool {
	if r.frontmatter == nil {
		return false
	}
	for i := 0; i+1 < len(r.frontmatter.Content); i += 2 {
		if r.frontmatter.Content[i].Value == key {
			r.frontmatter.Content = append(
				r.frontmatter.Content[:i], r.frontmatter.Content[i+2:]...)
			return true
		}
	}
	return false
}

// Keys lists the frontmatter keys in the order they appear.
func (r *Record) Keys() []string {
	if r.frontmatter == nil {
		return nil
	}
	keys := make([]string, 0, len(r.frontmatter.Content)/2)
	for i := 0; i+1 < len(r.frontmatter.Content); i += 2 {
		keys = append(keys, r.frontmatter.Content[i].Value)
	}
	return keys
}

// find locates a key and its value node.
func (r *Record) find(key string) (k, v *yaml.Node) {
	if r.frontmatter == nil {
		return nil, nil
	}
	for i := 0; i+1 < len(r.frontmatter.Content); i += 2 {
		if r.frontmatter.Content[i].Value == key {
			return r.frontmatter.Content[i], r.frontmatter.Content[i+1]
		}
	}
	return nil, nil
}

// Hash is a content hash of the serialized record, for detecting a write that
// would clobber a change nobody saw (docs/SPEC.md §6.3).
//
// Content rather than modification time: timestamps are too coarse and are
// vulnerable to clock skew, and the format's own `modified` advances only on
// meaningful change, so it is not a reliable witness either.
func Hash(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}
