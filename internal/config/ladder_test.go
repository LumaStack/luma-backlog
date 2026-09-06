package config

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func ladderFrom(t *testing.T, doc string) (Config, error) {
	t.Helper()
	var c Config
	err := yaml.Unmarshal([]byte(doc), &c)
	return c, err
}

// Every configuration written before ordinals existed uses a list, including
// this repository's own. It must keep working, and derive ordinals that leave
// room to insert between them.
func TestAListDerivesSparseOrdinals(t *testing.T) {
	c, err := ladderFrom(t, "workflow_status:\n  work-item: [captured, todo, closed]\n")
	if err != nil {
		t.Fatal(err)
	}
	l := c.LadderFor("work-item")
	if got := l.Statuses; len(got) != 3 || got[0] != "captured" || got[2] != "closed" {
		t.Fatalf("statuses = %v", got)
	}
	for _, tc := range []struct {
		status string
		want   int
	}{{"captured", 10}, {"todo", 20}, {"closed", 30}} {
		if n, ok := l.Ordinal(tc.status); !ok || n != tc.want {
			t.Errorf("ordinal(%s) = %d, %v; want %d", tc.status, n, ok, tc.want)
		}
	}
}

// The explicit form is what ADR-0005 asks for, and the one that lets a status
// be inserted without renumbering.
func TestAMappingKeepsItsOrdinalsAndItsOrder(t *testing.T) {
	c, err := ladderFrom(t, "workflow_status:\n  work-item:\n    captured: 10\n    todo: 50\n    closed: 90\n")
	if err != nil {
		t.Fatal(err)
	}
	l := c.LadderFor("work-item")
	// Document order is ladder order. Ranging a Go map would lose it.
	if got := strings.Join(l.Statuses, ","); got != "captured,todo,closed" {
		t.Errorf("order = %s", got)
	}
	if n, _ := l.Ordinal("todo"); n != 50 {
		t.Errorf("todo ordinal = %d, want 50", n)
	}
}

// Ascending order is the work order. A ladder that reads one way and sorts
// another would put the board in an order nobody wrote down, so it is refused
// at load rather than discovered on screen.
func TestOrdinalsMustAscendWithTheLadder(t *testing.T) {
	for _, tc := range []struct{ name, doc, want string }{
		{"descending", "workflow_status:\n  work-item:\n    captured: 50\n    todo: 10\n", "does not come after"},
		{"duplicate", "workflow_status:\n  work-item:\n    captured: 10\n    todo: 10\n", "both have ordinal"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ladderFrom(t, tc.doc)
			if err == nil {
				t.Fatal("accepted a ladder whose ordinals do not ascend")
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("error did not explain the problem: %v", err)
			}
		})
	}
}

// A status the ladder does not carry has no place in the order. Reporting that
// beats inventing an ordinal and filing the record somewhere nobody chose.
func TestAnUnknownStatusHasNoOrdinal(t *testing.T) {
	c, _ := ladderFrom(t, "workflow_status:\n  work-item: [captured, todo]\n")
	if n, ok := c.LadderFor("work-item").Ordinal("marinating"); ok {
		t.Errorf("unknown status reported ordinal %d", n)
	}
}
