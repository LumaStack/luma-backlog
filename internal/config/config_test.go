package config

import (
	"reflect"
	"testing"

	"github.com/lumastack/luma-backlog/internal/record"
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
	if got := c.DefaultStatusFor("deliverable"); got != "idea" {
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
	if got, want := c.DefaultStatusFor("deliverable"), "idea"; got != want {
		t.Errorf("deliverable default = %q, want %q", got, want)
	}
	if got, want := c.DefaultStatusFor("task"), "todo"; got != want {
		t.Errorf("task default = %q, want %q", got, want)
	}
	// An unknown unit falls back to the deliverable's vocabulary rather than
	// to nothing, so a new unit type is usable before it is configured.
	if got, want := c.DefaultStatusFor("wave"), "idea"; got != want {
		t.Errorf("unknown unit default = %q, want %q", got, want)
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

func TestBundleRootIsAValidRecord(t *testing.T) {
	// Parsed by the tool's own parser rather than a raw YAML unmarshal — the
	// file is frontmatter plus markdown, and checking it any other way would
	// be checking something the tool never does.
	r, err := record.Parse([]byte(BundleRootFile))
	if err != nil {
		t.Fatalf("bundle root does not parse: %v", err)
	}
	if typ, ok := r.Get("type"); !ok || typ == "" {
		t.Error("bundle root has no type — the format's one hard requirement")
	}
	// The regenerated keys must agree with the configuration they mirror.
	def := Default()
	if got, _ := r.Get("type_namespace"); got != def.TypeNamespace {
		t.Errorf("bundle root type_namespace = %q, config default %q", got, def.TypeNamespace)
	}
	if got, _ := r.Get("lkf_version"); got != def.LKFVersion {
		t.Errorf("bundle root lkf_version = %q, config default %q", got, def.LKFVersion)
	}
}
