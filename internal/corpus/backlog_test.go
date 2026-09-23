package corpus

import "testing"

func TestSlugify(t *testing.T) {
	for _, tc := range []struct{ in, want string }{
		{"Add retry queue", "add-retry-queue"},
		{"  Trim  me  ", "trim-me"},
		{"Retry queue: drain it!", "retry-queue-drain-it"},
		{"Retry queue - drain it", "retry-queue-drain-it"}, // agrees with the line above
		{"--leading and trailing--", "leading-and-trailing"},
		{"CAPS and 123", "caps-and-123"},
		{"café déjà vu", "café-déjà-vu"}, // letters are letters
		{"!!!", ""},                      // nothing usable
		{"", ""},
	} {
		if got := Slugify(tc.in); got != tc.want {
			t.Errorf("Slugify(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestPathFor(t *testing.T) {
	for _, tc := range []struct {
		unit, slug, workItem, want string
	}{
		{WorkItem, "payments", "", "backlog/work-items/payments/index.md"},
		{Outcome, "queue-drains", "payments", "backlog/work-items/payments/outcomes/queue-drains.md"},
		{Task, "add-queue", "payments", "backlog/work-items/payments/tasks/add-queue.md"},
		{Exploration, "spike", "payments", "backlog/work-items/payments/explorations/spike.md"},
		{Decision, "use-postgres", "payments", "backlog/work-items/payments/decisions/use-postgres.md"},
		// A decision made outside any work item's legal and sits at the top.
		{Decision, "use-postgres", "", "records/decisions/use-postgres"},
	} {
		got, err := PathFor(tc.unit, tc.slug, tc.workItem)
		if err != nil {
			t.Errorf("PathFor(%s) errored: %v", tc.unit, err)
			continue
		}
		if got != tc.want {
			t.Errorf("PathFor(%s, %s, %s) = %q, want %q", tc.unit, tc.slug, tc.workItem, got, tc.want)
		}
	}
}

func TestPathForRequiresAWorkItemWhereItMatters(t *testing.T) {
	// An outcome or task without a work item would float, and nothing
	// would ever count it toward completion.
	for _, unit := range []string{Outcome, Task, Exploration} {
		if _, err := PathFor(unit, "thing", ""); err == nil {
			t.Errorf("PathFor(%s) with no work item succeeded", unit)
		}
	}
}

func TestPathForRejectsUnknownUnits(t *testing.T) {
	if _, err := PathFor("sprint", "thing", "payments"); err == nil {
		t.Error("an unknown unit was accepted")
	}
}

func TestWorkItemFromPath(t *testing.T) {
	for _, tc := range []struct{ in, want string }{
		{".luma/work-items/payments/tasks", "payments"},
		{"backlog/work-items/payments", "payments"},
		{"backlog/work-items", ""},
		{"docs", ""},
		{"", ""},
	} {
		if got := WorkItemFromPath(tc.in); got != tc.want {
			t.Errorf("WorkItemFromPath(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// The corpus accepts only what this tool writes.
//
// It used to accept everything it had not been told to exclude, which is only
// correct if this tool owns .luma/. It shares the directory with the other luma
// tools, so each new one had to be predicted — and was not (BACK-0108).
func TestOnlyPathsThisToolWritesAreRecords(t *testing.T) {
	cases := []struct {
		path string
		want bool
		why  string
	}{
		{"PROJECT.md", true, "reserved by the layout policy and reached by show"},
		{"backlog/work-items/WORK-0001-a/index.md", true, "the work item itself"},
		{"backlog/work-items/WORK-0001-a/outcomes/x.md", true, "a child, in its unit's directory"},
		{"backlog/work-items/WORK-0001-a/tasks/x.md", true, "a child, in its unit's directory"},
		{"records/decisions/ADR-0001-x.md", true, "a decision outside any work item"},

		{"backlog/work-items/WORK-0001-a/journal.md", false, "a ledger, not a record"},
		{"backlog/work-items/WORK-0001-a/evidence/log.md", false, "an attachment"},
		{"backlog/work-items/WORK-0001-a/notes.md", false, "a stray file beside a record"},
		{"records/decisions/archived/ADR-0001-x.md", false, "spent; listing it reads as current"},
		{"records/violations/2026-01-01-x/violation.md", false, "written by a procedure, not by this tool"},
		{"backlog/ideas/routers.md", false, "another tool's ideas"},
		{"backlog/plans/bundle-publish.md", false, "another tool's plans"},
		{"bundles/local/backlog/policy/x.md", false, "what is in force, not what is intended"},
		{"_types/luma/backlog/task.md", false, "a contract, not a record"},
	}
	for _, c := range cases {
		if got := isRecordPath(c.path); got != c.want {
			t.Errorf("isRecordPath(%q) = %v, want %v — %s", c.path, got, c.want, c.why)
		}
	}
}
