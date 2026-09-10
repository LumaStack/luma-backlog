package cli

import (
	"strings"

	"github.com/spf13/cobra"
	"testing"
)

func rootHelp(t *testing.T) string {
	t.Helper()
	app, _ := initialized(t)
	code, out, errOut := run(t, app)
	if code != ExitOK {
		t.Fatalf("exit = %d: %s", code, errOut)
	}
	return out
}

// Uppercase headings without colons, in the order gh uses them (ADR-0006).
func TestHelpSectionsAreInOrder(t *testing.T) {
	out := rootHelp(t)
	want := []string{"USAGE", "CORE COMMANDS", "RECORD COMMANDS", "ADDITIONAL COMMANDS", "FLAGS", "EXAMPLES", "LEARN MORE"}
	at := -1
	for _, section := range want {
		i := strings.Index(out, "\n"+section+"\n")
		if i < 0 {
			t.Fatalf("no %q section:\n%s", section, out)
		}
		if i < at {
			t.Errorf("%q is out of order:\n%s", section, out)
		}
		at = i
	}
	if strings.Contains(out, "Available Commands:") || strings.Contains(out, "Usage:") {
		t.Errorf("Cobra's own headings survived:\n%s", out)
	}
}

// Every command is filed under a heading. Flat, `completion` and `help` sat
// between commands that do the work as though they were peers.
func TestEveryCommandIsGrouped(t *testing.T) {
	out := rootHelp(t)
	if strings.Contains(out, "\nCOMMANDS\n") {
		t.Errorf("a command was left ungrouped:\n%s", out)
	}
	for section, commands := range map[string][]string{
		"CORE COMMANDS":       {"list", "show", "set"},
		"RECORD COMMANDS":     {"work-item", "outcome", "task", "decision", "exploration"},
		"ADDITIONAL COMMANDS": {"init", "completion", "help"},
	} {
		body := section3(out, section)
		for _, c := range commands {
			if !strings.Contains(body, c) {
				t.Errorf("%s missing from %s:\n%s", c, section, out)
			}
		}
	}
}

// IsAvailableCommand is false for the help command --- Cobra excludes it --- so
// a listing that forgets to say so loses it silently.
func TestHelpItselfIsListed(t *testing.T) {
	if !strings.Contains(section3(rootHelp(t), "ADDITIONAL COMMANDS"), "help") {
		t.Error("the help command is missing from its own listing")
	}
}

// Every example must name a command that exists --- a stale one teaches a
// surface that is gone, which is exactly what happened when `journal` moved
// under `work-item`. This walks the tree; it does not run anything.
func TestExamplesNameCommandsThatExist(t *testing.T) {
	app, _ := initialized(t)
	for _, line := range strings.Split(section3(rootHelp(t), "EXAMPLES"), "\n") {
		line = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "$"))
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 || fields[0] != "luma-backlog" {
			t.Errorf("example is not a luma-backlog invocation: %q", line)
			continue
		}

		// Walk while each word names a child. Cobra's Find returns the root
		// with the arguments unconsumed rather than an error, so asking it
		// whether a command exists answers yes for anything.
		cmd := newRootCommand(app)
		depth := 0
		for _, word := range fields[1:] {
			if strings.HasPrefix(word, "-") {
				break
			}
			var child *cobra.Command
			for _, c := range cmd.Commands() {
				if c.Name() == word {
					child = c
					break
				}
			}
			if child == nil {
				break
			}
			cmd, depth = child, depth+1
		}
		if depth == 0 {
			t.Errorf("example names no command that exists: %q", line)
		}
	}
}

// A leaf keeps its own arguments and gets the same headings.
func TestALeafUsesTheSameLayout(t *testing.T) {
	app, _ := initialized(t)
	_, out, _ := run(t, app, "work-item", "rank", "--help")
	if !strings.Contains(out, "\nUSAGE\n") || !strings.Contains(out, "\nFLAGS\n") {
		t.Errorf("a leaf did not get the layout:\n%s", out)
	}
	if !strings.Contains(out, "luma-backlog work-item rank <work-item>") {
		t.Errorf("a leaf lost its own arguments:\n%s", out)
	}
}

// section3 returns the body of one uppercase section.
func section3(out, name string) string {
	i := strings.Index(out, "\n"+name+"\n")
	if i < 0 {
		return ""
	}
	rest := out[i+len(name)+2:]
	if j := strings.Index(rest, "\n\n"); j >= 0 {
		return rest[:j]
	}
	return rest
}

// Cobra sorts alphabetically, which throws away two orderings the code already
// states: the core verbs by how often they are reached for, and corpus.Units by
// the model's hierarchy. Alphabetical is a third ordering that means nothing.
func TestCommandsListInTheOrderTheyAreDeclared(t *testing.T) {
	app, _ := initialized(t)
	_, out, _ := run(t, app, "--help")

	for _, group := range [][]string{
		{"list", "show", "set", "rank", "transition"},
		{"work-item", "outcome", "task", "decision", "exploration"},
	} {
		at := -1
		for _, name := range group {
			i := strings.Index(out, "\n  "+name+" ")
			if i < 0 {
				t.Fatalf("%q missing from the listing:\n%s", name, out)
			}
			if i < at {
				t.Errorf("%q listed out of order:\n%s", name, out)
			}
			at = i
		}
	}
}
