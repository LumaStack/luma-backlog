package config

import (
	"reflect"
	"strings"
	"testing"
)

// TestDefaultFileMatchesDefaults stops the written file and the built-in
// fallbacks drifting apart. They are two expressions of one thing, and nothing
// else would notice if they disagreed.
func TestDefaultFileMatchesDefaults(t *testing.T) {
	parsed, err := Parse([]byte(DefaultFile))
	if err != nil {
		t.Fatal(err)
	}
	def := Default()

	if parsed.KeyPrefix() != def.KeyPrefix() {
		t.Errorf("work_item_key: file %q, default %q", parsed.KeyPrefix(), def.KeyPrefix())
	}
	if !reflect.DeepEqual(parsed.WorkflowStatus, def.WorkflowStatus) {
		t.Errorf("workflow_status: file %v, default %v", parsed.WorkflowStatus, def.WorkflowStatus)
	}
}

func TestParseKeepsFallbacksForAbsentKeys(t *testing.T) {
	// A configuration written today must keep working when keys are added.
	c, err := Parse([]byte("work_item_key: BACK\n"))
	if err != nil {
		t.Fatal(err)
	}
	if c.KeyPrefix() != "BACK" {
		t.Errorf("KeyPrefix = %q", c.KeyPrefix())
	}
	if got := c.DefaultStatusFor("work-item"); got != "captured" {
		t.Errorf("fallback lost: DefaultStatusFor = %q, want captured", got)
	}
}

func TestParseRejectsMalformedYAML(t *testing.T) {
	if _, err := Parse([]byte("workflow_status: [unclosed\n")); err == nil {
		t.Error("malformed configuration parsed without error")
	}
}

func TestAnAbsentWorkItemKeyMeansWork(t *testing.T) {
	c, err := Parse([]byte("workflow_status:\n  task:\n    todo: 50\n"))
	if err != nil {
		t.Fatal(err)
	}
	if c.KeyPrefix() != "WORK" {
		t.Errorf("KeyPrefix = %q, want WORK", c.KeyPrefix())
	}
	// A zero-value Config never went through Default or Parse and must not
	// produce keys like "-0042".
	if got := (Config{}).KeyPrefix(); got != "WORK" {
		t.Errorf("zero-value KeyPrefix = %q, want WORK", got)
	}
}

func TestParseRefusesAPrefixOutsideTheJiraCloudRule(t *testing.T) {
	// An uppercase letter, then uppercase letters or digits, two to ten
	// characters. A prefix the validator accepted but the key reader could
	// not parse would create records the tool cannot find.
	for _, bad := range []string{"back", "X", "2AB", "ABCDEFGHIJK", "WORK_ITEMS", "BA-CK", `""`} {
		_, err := Parse([]byte("work_item_key: " + bad + "\n"))
		if err == nil {
			t.Errorf("work_item_key %q accepted; want a refusal", bad)
			continue
		}
		if !strings.Contains(err.Error(), KeyPrefixRule) {
			t.Errorf("refusal of %q does not name the rule: %v", bad, err)
		}
	}
}

func TestDefaultStatusIsTheFirstConfiguredValue(t *testing.T) {
	c := Default()
	if got, want := c.DefaultStatusFor("work-item"), "captured"; got != want {
		t.Errorf("work item default = %q, want %q", got, want)
	}
	if got, want := c.DefaultStatusFor("task"), "todo"; got != want {
		t.Errorf("task default = %q, want %q", got, want)
	}
	// An unknown unit falls back to the TASK vocabulary. The extra rungs
	// describe how far planning has gone on a backlog item, and only a
	// work item's one — stamping a new outcome "captured" would file it in
	// the Backlog column.
	for _, unit := range []string{"outcome", "exploration", "wave"} {
		if got, want := c.DefaultStatusFor(unit), "todo"; got != want {
			t.Errorf("%s default = %q, want %q", unit, got, want)
		}
	}
}
