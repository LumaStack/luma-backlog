// Package config holds the settings a team owns, so the binary can hold none.
//
// Every choice the specification leaves open — what states exist, what things
// are called, how work is classified — resolves here rather than in code.
package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// FileName is the configuration file within the backlog directory.
const FileName = "config.yml"

// Config is the settings a repository declares.
type Config struct {
	LKFVersion     string              `yaml:"lkf_version"`
	TypeNamespace  string              `yaml:"type_namespace"`
	WorkflowStatus map[string][]string `yaml:"workflow_status"`
	Columns        yaml.Node           `yaml:"columns"`
}

// Default returns the built-in fallbacks.
//
// These exist as well as being written to disk, not instead of: the file makes
// a default discoverable, and the fallback keeps a configuration written today
// working when new keys are added later.
func Default() Config {
	var columns yaml.Node
	_ = yaml.Unmarshal([]byte(defaultColumns), &columns)
	if len(columns.Content) > 0 {
		columns = *columns.Content[0]
	}
	return Config{
		LKFVersion:    "0.0.2",
		TypeNamespace: "luma/backlog",
		WorkflowStatus: map[string][]string{
			"deliverable": {"idea", "preparing", "actionable", "todo", "in_progress", "closed"},
			"task":        {"todo", "in_progress", "closed"},
		},
		Columns: columns,
	}
}

const defaultColumns = `
Backlog:     [idea, preparing, actionable]
To Do:       [todo]
In Progress: [in_progress]
Closed:      [closed]
`

// Parse reads a configuration file.
//
// Strict where records are permissive, and deliberately so: a record with an
// unrecognised field is tolerated because knowledge arrives incomplete, while
// a misspelt configuration key is a silent behaviour change (docs/SPEC.md §8.7).
func Parse(data []byte) (Config, error) {
	c := Default()
	if err := yaml.Unmarshal(data, &c); err != nil {
		return Config{}, fmt.Errorf("parsing configuration: %w", err)
	}
	if c.TypeNamespace == "" {
		return Config{}, fmt.Errorf("type_namespace is empty: short type names could not be resolved")
	}
	return c, nil
}

// StatusesFor returns the workflow vocabulary for a unit, falling back to the
// deliverable's when a unit has none of its own.
func (c Config) StatusesFor(unit string) []string {
	if s, ok := c.WorkflowStatus[unit]; ok && len(s) > 0 {
		return s
	}
	return c.WorkflowStatus["deliverable"]
}

// DefaultStatusFor is what an absent workflow_status means: the first value in
// the configured vocabulary. A field is only safely optional when omitting it
// says something (docs/SPEC.md §4.2).
func (c Config) DefaultStatusFor(unit string) string {
	if s := c.StatusesFor(unit); len(s) > 0 {
		return s[0]
	}
	return ""
}

// Qualify expands a short type name using the declared namespace. An already
// qualified name is returned unchanged — it is always legal and always wins.
func (c Config) Qualify(typeName string) string {
	if typeName == "" || c.TypeNamespace == "" {
		return typeName
	}
	for _, ch := range typeName {
		if ch == '/' {
			return typeName
		}
	}
	return c.TypeNamespace + "/" + typeName
}
