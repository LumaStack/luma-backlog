package backlog

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
		unit, slug, deliverable, want string
	}{
		{Deliverable, "payments", "", "backlog/deliverables/payments/index.md"},
		{Outcome, "queue-drains", "payments", "backlog/deliverables/payments/outcomes/queue-drains.md"},
		{Task, "add-queue", "payments", "backlog/deliverables/payments/tasks/add-queue.md"},
		{Exploration, "spike", "payments", "backlog/deliverables/payments/explorations/spike.md"},
		{Decision, "use-postgres", "payments", "backlog/deliverables/payments/decisions/use-postgres.md"},
		// A decision made outside any deliverable is legal and sits at the top.
		{Decision, "use-postgres", "", "records/decisions/use-postgres"},
	} {
		got, err := PathFor(tc.unit, tc.slug, tc.deliverable)
		if err != nil {
			t.Errorf("PathFor(%s) errored: %v", tc.unit, err)
			continue
		}
		if got != tc.want {
			t.Errorf("PathFor(%s, %s, %s) = %q, want %q", tc.unit, tc.slug, tc.deliverable, got, tc.want)
		}
	}
}

func TestPathForRequiresADeliverableWhereItMatters(t *testing.T) {
	// An outcome or task without a deliverable would float, and nothing
	// would ever count it toward completion.
	for _, unit := range []string{Outcome, Task, Exploration} {
		if _, err := PathFor(unit, "thing", ""); err == nil {
			t.Errorf("PathFor(%s) with no deliverable succeeded", unit)
		}
	}
}

func TestPathForRejectsUnknownUnits(t *testing.T) {
	if _, err := PathFor("sprint", "thing", "payments"); err == nil {
		t.Error("an unknown unit was accepted")
	}
}

func TestDeliverableFromPath(t *testing.T) {
	for _, tc := range []struct{ in, want string }{
		{".luma/deliverables/payments/tasks", "payments"},
		{"backlog/deliverables/payments", "payments"},
		{"backlog/deliverables", ""},
		{"docs", ""},
		{"", ""},
	} {
		if got := DeliverableFromPath(tc.in); got != tc.want {
			t.Errorf("DeliverableFromPath(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
