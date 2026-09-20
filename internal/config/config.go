// Package config holds the settings a team owns, so the binary can hold none.
//
// Every choice the specification leaves open — what states exist, what things
// are called, how work is classified — resolves here rather than in code.
package config

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v3"
)

// FileName is the tool's configuration file, relative to .luma/. It lives in
// the config tier per the luma directory layout policy: one file per tool,
// named for the tool so nobody has to guess which binary a file belongs to.
// It was config/luma-corpus.yaml until WORK-0053: that name matched an
// internal package rename rather than the tool, so the one repository using
// the tool had a configuration that was never read.
const FileName = "config/luma-backlog.yaml"

// DefaultKeyPrefix is what an absent work_item_key means. The prefix is
// written INTO each record at creation rather than derived from configuration,
// so changing the setting changes what gets written and not what already
// exists — a derived key would silently rename every record in the corpus the
// moment the setting changed.
const DefaultKeyPrefix = "WORK"

// KeyPrefixRule is the shape a key prefix must have, as a regular expression
// fragment: an uppercase letter, then uppercase letters or digits, two to ten
// characters — Jira Cloud's project-key rules (WORK-0102). One fragment, so
// the validator here and the key reader in internal/corpus cannot drift.
const KeyPrefixRule = `[A-Z][A-Z0-9]{1,9}`

var keyPrefixPattern = regexp.MustCompile(`^` + KeyPrefixRule + `$`)

// Config is the settings a repository declares.
//
// Every field here is read by something. A key the binary parses and never
// consumes is a placeholder wearing a promise — it teaches an author the
// setting works, and nothing notices that it does not (WORK-0053 found
// exactly that: a whole file that was never read, invisible because its
// values matched the defaults).
type Config struct {
	WorkItemKey    string            `yaml:"work_item_key"`
	WorkflowStatus map[string]Ladder `yaml:"workflow_status"`
}

// Default returns the built-in fallbacks.
//
// These exist as well as being written to disk, not instead of: the file makes
// a default discoverable, and the fallback keeps a configuration written today
// working when new keys are added later.
func Default() Config {
	return Config{
		WorkItemKey: DefaultKeyPrefix,
		// A status shared by two units carries the same ordinal in both, so a
		// rank means the same thing whichever unit it is on. Deriving them per
		// unit would give a task's `todo` a different ordinal from a work
		// item's, and sorting raw ranks across units would interleave them
		// wrongly --- which is the one thing the ordinal prefix exists to
		// prevent (ADR-0005).
		WorkflowStatus: map[string]Ladder{
			"work-item": ladderAt(
				"captured", 10, "unprepared", 20, "preparing", 30, "prepared", 40,
				"todo", 50, "in_progress", 60, "closed", 70),
			"task": ladderAt("todo", 50, "in_progress", 60, "closed", 70),
		},
	}
}

// Parse reads a configuration file.
//
// Strict where records are permissive, and deliberately so: a record with an
// unrecognized field is tolerated because knowledge arrives incomplete, while
// a misspelt configuration key is a silent behavior change (docs/spec.md §8.7).
func Parse(data []byte) (Config, error) {
	c := Default()
	if err := yaml.Unmarshal(data, &c); err != nil {
		return Config{}, fmt.Errorf("parsing configuration: %w", err)
	}
	if !keyPrefixPattern.MatchString(c.WorkItemKey) {
		return Config{}, fmt.Errorf(
			"work_item_key %q is not a legal key prefix: an uppercase letter, then uppercase letters or digits, two to ten characters (^%s$)",
			c.WorkItemKey, KeyPrefixRule)
	}
	return c, nil
}

// KeyPrefix is the prefix new work item keys are written under. It falls back
// for a zero-value Config, which never went through Default or Parse.
func (c Config) KeyPrefix() string {
	if c.WorkItemKey == "" {
		return DefaultKeyPrefix
	}
	return c.WorkItemKey
}

// StatusesFor returns the workflow vocabulary for a unit.
//
// A unit with none of its own falls back to the TASK vocabulary, not the
// work item's. The extra rungs — captured, unprepared, preparing, prepared —
// describe how far the planning has gone on a backlog item, and a work item's
// the backlog item. An outcome or an exploration is never "something we might
// drop"; it is todo, in progress, or closed.
//
// Falling back to the work item's list instead would stamp a new outcome
// "captured", which is not merely odd — it would file it above the first
// selection gate, among the things nobody has decided to do.
func (c Config) StatusesFor(unit string) []string {
	return c.LadderFor(unit).Statuses
}

// LadderFor is StatusesFor with the ordinals kept, for callers that need to
// order across statuses rather than only name them.
func (c Config) LadderFor(unit string) Ladder {
	if l, ok := c.WorkflowStatus[unit]; ok && len(l.Statuses) > 0 {
		return l
	}
	if l, ok := c.WorkflowStatus["task"]; ok && len(l.Statuses) > 0 {
		return l
	}
	return c.WorkflowStatus["work-item"]
}

// DefaultStatusFor is what an absent workflow_status means: the first value in
// the configured vocabulary. A field is only safely optional when omitting it
// says something (docs/spec.md §4.2).
func (c Config) DefaultStatusFor(unit string) string {
	if s := c.StatusesFor(unit); len(s) > 0 {
		return s[0]
	}
	return ""
}

// TerminalStatusFor is the last value in the configured vocabulary — where
// work ends.
//
// Derived rather than configured, and it mirrors DefaultStatusFor taking the
// first: the ladder is ordered and its ordinals ascend, so the last rung is the
// one nothing follows. A project that renames `closed` keeps working, and one
// that adds a rung after it moves the terminal with no key to remember.
//
// Empty where a unit has no vocabulary, which reads as "nothing is terminal"
// and leaves every record open.
func (c Config) TerminalStatusFor(unit string) string {
	s := c.StatusesFor(unit)
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

