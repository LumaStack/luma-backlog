package config

import (
	"reflect"
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

	if parsed.LKFVersion != def.LKFVersion {
		t.Errorf("lkf_version: file %q, default %q", parsed.LKFVersion, def.LKFVersion)
	}
	if parsed.TypeNamespace != def.TypeNamespace {
		t.Errorf("type_namespace: file %q, default %q", parsed.TypeNamespace, def.TypeNamespace)
	}
	if !reflect.DeepEqual(parsed.WorkflowStatus, def.WorkflowStatus) {
		t.Errorf("workflow_status: file %v, default %v", parsed.WorkflowStatus, def.WorkflowStatus)
	}

	var fromFile, fromDefault map[string][]string
	if err := parsed.Columns.Decode(&fromFile); err != nil {
		t.Fatal(err)
	}
	if err := def.Columns.Decode(&fromDefault); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fromFile, fromDefault) {
		t.Errorf("columns: file %v, default %v", fromFile, fromDefault)
	}
}

func TestParseKeepsFallbacksForAbsentKeys(t *testing.T) {
	// A configuration written today must keep working when keys are added.
	c, err := Parse([]byte("type_namespace: acme/work\n"))
	if err != nil {
		t.Fatal(err)
	}
	if c.TypeNamespace != "acme/work" {
		t.Errorf("TypeNamespace = %q", c.TypeNamespace)
	}
	if got := c.DefaultStatusFor("work-item"); got != "idea" {
		t.Errorf("fallback lost: DefaultStatusFor = %q, want idea", got)
	}
}

func TestParseRejectsMalformedYAML(t *testing.T) {
	if _, err := Parse([]byte("workflow_status: [unclosed\n")); err == nil {
		t.Error("malformed configuration parsed without error")
	}
}

func TestParseRejectsAnEmptyNamespace(t *testing.T) {
	// Silently accepting this would leave every short type name unresolvable.
	if _, err := Parse([]byte("type_namespace: \"\"\n")); err == nil {
		t.Error("empty type_namespace accepted")
	}
}

func TestDefaultStatusIsTheFirstConfiguredValue(t *testing.T) {
	c := Default()
	if got, want := c.DefaultStatusFor("work-item"), "idea"; got != want {
		t.Errorf("work item default = %q, want %q", got, want)
	}
	if got, want := c.DefaultStatusFor("task"), "todo"; got != want {
		t.Errorf("task default = %q, want %q", got, want)
	}
	// An unknown unit falls back to the TASK vocabulary. The extra rungs
	// describe how far planning has gone on a backlog item, and only a
	// work item's one — stamping a new outcome "idea" would file it in
	// the Backlog column.
	for _, unit := range []string{"outcome", "exploration", "wave"} {
		if got, want := c.DefaultStatusFor(unit), "todo"; got != want {
			t.Errorf("%s default = %q, want %q", unit, got, want)
		}
	}
}

func TestQualify(t *testing.T) {
	c := Default()
	if got, want := c.Qualify("task"), "luma/backlog/task"; got != want {
		t.Errorf("Qualify(task) = %q, want %q", got, want)
	}
	// Already qualified wins, so a record from a foreign vocabulary survives.
	if got, want := c.Qualify("acme/thing/task"), "acme/thing/task"; got != want {
		t.Errorf("Qualify of a qualified name = %q, want %q", got, want)
	}
	if got := c.Qualify(""); got != "" {
		t.Errorf("Qualify(empty) = %q, want empty", got)
	}
}
