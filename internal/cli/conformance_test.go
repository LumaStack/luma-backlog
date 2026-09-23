package cli

import (
	"strings"
	"testing"

	"github.com/lumastack/luma-backlog/internal/env"
	"github.com/spf13/cobra"
)

// Every message a command prints is checked against the house shape.
//
// This exists because finding them by hand did not work. Messages were
// converted one at a time as somebody tripped over them, and the ones nobody
// tripped over stayed in whatever style they were written in — so the tool
// spoke four dialects and no single reader ever saw enough of them to notice.
//
// The rules are `lumastack/luma-catalog/command-line-interface`,
// policy/output-patterns. What is checked here is what can be checked
// mechanically: shape, not wording.
// messageCase is one message a command can produce.
type messageCase struct {
	name string
	args []string
	// setup runs first, against the same project.
	setup [][]string
	// bare runs against a repository with no backlog, and outside runs
	// with no repository at all. Without these the two refusals somebody
	// meets FIRST were the two nothing checked.
	bare    bool
	outside bool
}

// messages is the battery. Add a row when adding a message: one nobody
// exercises here is one nobody is checking.
//
// Every check below runs over ALL of it. Keeping a second list was how the
// offered-command check ended up exercising no offered commands.
var messages = []messageCase{
	{name: "not initialized", args: []string{"list"}, bare: true},
	{name: "not initialized, another verb", args: []string{"work-item", "list"}, bare: true},
	{name: "no repository", args: []string{"init"}, outside: true},
	{name: "no repository, another verb", args: []string{"list"}, outside: true},

	{name: "unknown command", args: []string{"work-tiem"}},
	{name: "unknown command, no near miss", args: []string{"frobnicate"}},
	{name: "unknown subcommand", args: []string{"work-item", "lst"}},
	{name: "unknown flag", args: []string{"list", "--stauts", "closed"}},
	{name: "wrong argument count", args: []string{"show"}},

	{name: "empty, nothing yet", args: []string{"list"}},
	{name: "empty, filter matched nothing", args: []string{"list", "--status", "closed"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},

	{name: "reference names nothing", args: []string{"show", "nope-not-here"}},
	{name: "reference names several", args: []string{"show", "a"},
		setup: [][]string{{"work-item", "new", "Alpha one"}, {"work-item", "new", "Alpha two"}}},

	{name: "missing input, disposition", args: []string{"work-item", "close", "alpha"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},
	{name: "unknown disposition", args: []string{"work-item", "close", "alpha", "finished"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},
	{name: "missing input, verdict", args: []string{"outcome", "verify", "the-queue-drains"},
		setup: [][]string{{"work-item", "new", "Alpha"}, {"outcome", "new", "The queue drains", "-w", "alpha"}}},
	{name: "missing input, claim", args: []string{"outcome", "assert", "the-queue-drains"},
		setup: [][]string{{"work-item", "new", "Alpha"}, {"outcome", "new", "The queue drains", "-w", "alpha"}}},
	{name: "missing input, rank", args: []string{"work-item", "rank", "alpha"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},
	{name: "missing input, transition", args: []string{"work-item", "transition", "alpha"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},
	{name: "missing input, set", args: []string{"set", "alpha"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},
	{name: "missing input, title", args: []string{"work-item", "new"}},

	{name: "wrong type for the verb", args: []string{"work-item", "close", "the-queue-drains", "completed"},
		setup: [][]string{{"work-item", "new", "Alpha"}, {"outcome", "new", "The queue drains", "-w", "alpha"}}},
	{name: "field that is not set directly", args: []string{"set", "alpha", "workflow_status=todo"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},
	{name: "status the ladder does not carry", args: []string{"work-item", "transition", "alpha", "nowhere"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},
	{name: "terminal status through transition", args: []string{"work-item", "transition", "alpha", "closed"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},

	{name: "precondition, no outcomes", args: []string{"work-item", "close", "alpha", "completed"},
		setup: [][]string{{"work-item", "new", "Alpha"}}},
	{name: "precondition, outcomes unproven", args: []string{"work-item", "close", "alpha", "completed"},
		setup: [][]string{{"work-item", "new", "Alpha"}, {"outcome", "new", "The queue drains", "-w", "alpha"}}},

	// Success paths. Left out at first, on the assumption that only refusals
	// have shape — and two nonconforming reports sat behind that assumption
	// until somebody ran the commands by hand.
	{name: "report, created", args: []string{"work-item", "new", "Alpha"}},
	{name: "report, created with a kind", args: []string{"work-item", "new", "Alpha", "--kind", "change"}},
	{name: "report, created child", args: []string{"outcome", "new", "The queue drains", "-w", "alpha"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"}}},
	{name: "report, moved", args: []string{"work-item", "transition", "alpha", "todo"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"}}},
	{name: "report, ranked", args: []string{"work-item", "rank", "alpha", "--first"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"}}},
	{name: "report, set", args: []string{"set", "alpha", "description=hello"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"}}},
	{name: "report, journalled", args: []string{"work-item", "journal", "-w", "alpha", "learned a thing"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"}}},
	{name: "report, verified", args: []string{"outcome", "verify", "the-queue-drains", "proven", "-e", "ran it"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"},
			{"outcome", "new", "The queue drains", "-w", "alpha"}}},
	{name: "report, verified with no evidence", args: []string{"outcome", "verify", "the-queue-drains", "proven"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"},
			{"outcome", "new", "The queue drains", "-w", "alpha"}}},
	{name: "report, abandoned", args: []string{"outcome", "abandon", "the-queue-drains"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"},
			{"outcome", "new", "The queue drains", "-w", "alpha"}}},
	{name: "report, abandoned after a verdict", args: []string{"outcome", "abandon", "the-queue-drains", "-r", "dropped"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"},
			{"outcome", "new", "The queue drains", "-w", "alpha"},
			{"outcome", "verify", "the-queue-drains", "proven", "-e", "ran it"}}},
	{name: "report, closed", args: []string{"work-item", "close", "alpha", "canceled", "-r", "not doing it"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"}}},
	{name: "report, initialized", args: []string{"init"}, bare: true},
	{name: "report, initialized again", args: []string{"init"}},
	{name: "listing", args: []string{"list"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"}}},
	{name: "listing as a tree", args: []string{"list", "--tree"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"},
			{"outcome", "new", "The queue drains", "-w", "alpha"}}},
	{name: "show", args: []string{"show", "alpha"},
		setup: [][]string{{"work-item", "new", "Alpha", "--kind", "change"}}},
}

// project builds the repository this case expects.
func (c messageCase) project(t *testing.T) *App {
	t.Helper()
	var app *App
	switch {
	case c.outside:
		base := t.TempDir()
		app = &App{Env: env.New(), WorkingDir: base, Ceiling: base}
	case c.bare:
		app, _ = newApp(t)
	default:
		app, _ = initialized(t)
	}
	for _, s := range c.setup {
		if code, _, e := run(t, app, s...); code != ExitOK {
			t.Fatalf("setup %v failed: %s", s, e)
		}
	}
	return app
}

func TestEveryMessageFollowsTheOutputPatterns(t *testing.T) {
	for _, c := range messages {
		t.Run(c.name, func(t *testing.T) {
			app := c.project(t)
			_, out, errOut := run(t, app, c.args...)
			checkShape(t, "stderr", errOut)
			checkShape(t, "stdout", out)
		})
	}
}

// checkShape holds a message to the parts of the grammar a machine can see.
func checkShape(t *testing.T, stream, msg string) {
	t.Helper()
	if strings.TrimSpace(msg) == "" {
		return
	}

	// Backticks are shell syntax: a reader copying a line that contains them
	// gets command substitution rather than the command.
	if strings.Contains(msg, "`") {
		t.Errorf("%s wraps something in backticks:\n%s", stream, msg)
	}

	lines := strings.Split(strings.TrimRight(msg, "\n"), "\n")

	// A laid-out block and the tool-name prefix are alternatives, never both:
	// the prefix lands on the heading and fights the layout. A prefixed
	// message is the one-line diagnostic form and stays one line.
	for i, l := range lines {
		if !strings.HasPrefix(l, "luma-backlog:") {
			continue
		}
		rest := lines[i+1:]
		for _, r := range rest {
			if strings.TrimSpace(r) != "" && !strings.HasPrefix(r, "luma-backlog:") &&
				!strings.HasPrefix(r, "  ") {
				t.Errorf("%s: a prefixed diagnostic ran on into a block:\n%s", stream, msg)
				break
			}
		}
	}

	// A command is introduced by a colon and takes the indented line beneath
	// it. The lead-in is what says the command is for, so a bare colon line is
	// a slot half-filled.
	for i, l := range lines {
		if !strings.HasSuffix(l, ":") || strings.HasPrefix(l, " ") {
			continue
		}
		if strings.TrimSuffix(l, ":") == "" {
			t.Errorf("%s: a command lead-in says nothing:\n%s", stream, msg)
		}
		if i+1 >= len(lines) || !strings.HasPrefix(lines[i+1], "  ") ||
			strings.TrimSpace(lines[i+1]) == "" {
			t.Errorf("%s: a lead-in %q has no command under it:\n%s", stream, l, msg)
		}
	}

	// Detail is indented two spaces and belongs to the line above it, so the
	// first line of a block can never be indented.
	if strings.HasPrefix(lines[0], " ") {
		t.Errorf("%s opens with an indented line, so nothing owns it:\n%s", stream, msg)
	}
}

// A message that offers a command offers one somebody can actually run.
//
// The shape check cannot catch a command that is well-formed and wrong — a
// renamed verb, a flag that no longer exists — and that is exactly the decay a
// catalog of worked examples suffers. So every command any message prints is
// fed back through the tool, and must reach a real command.
func TestOfferedCommandsAreRealCommands(t *testing.T) {
	for _, c := range messages {
		t.Run(c.name, func(t *testing.T) {
			app := c.project(t)
			_, out, errOut := run(t, app, c.args...)

			for _, cmd := range append(offeredCommands(errOut), offeredCommands(out)...) {
				// Placeholders are the caller's to fill in, so the line is
				// checked as far as the first one.
				var fields []string
				for _, f := range strings.Fields(cmd) {
					if strings.HasPrefix(f, "<") || strings.HasPrefix(f, `"<`) {
						break
					}
					fields = append(fields, f)
				}
				if len(fields) == 0 || fields[0] != "luma-backlog" {
					continue // git init, or a path — not ours to run
				}
				if !commandExists(app, fields[1:]) {
					t.Errorf("offered a command that does not exist: %q", cmd)
				}
			}
		})
	}
}

// A report about a record names it the way somebody will type it next.
//
// The shape checks pass on `created  backlog/work-items/WORK-0001-x/index.md`
// — verb column, path, nothing half-filled — and it still breaks the rule that
// matters: showing-records asks that a key arrive with its title at least once
// in ephemeral output, and a path buries the key in a filename and the title in
// a slug. Shape was green while the content was wrong, twice, so the content is
// checked too.
func TestReportsNameTheRecordNotJustItsPath(t *testing.T) {
	// Each is a command that writes, and the key and title it must show.
	cases := []struct {
		name  string
		args  []string
		setup [][]string
		key   string
		title string
	}{
		{name: "created", args: []string{"work-item", "new", "Payments v2", "--kind", "change"},
			key: "WORK-0001", title: "Payments v2"},
		{name: "exists", args: []string{"work-item", "new", "Payments v2", "--kind", "change"},
			setup: [][]string{{"work-item", "new", "Payments v2", "--kind", "change"}},
			key:   "WORK-0001", title: "Payments v2"},
		{name: "updated", args: []string{"set", "payments-v2", "description=hello"},
			key: "WORK-0001", title: "Payments v2"},
		{name: "ranked", args: []string{"work-item", "rank", "payments-v2", "--first"},
			key: "WORK-0001", title: "Payments v2"},
		{name: "moved", args: []string{"work-item", "transition", "payments-v2", "todo"},
			key: "WORK-0001", title: "Payments v2"},
		{name: "closed", args: []string{"work-item", "close", "payments-v2", "canceled", "-r", "no"},
			key: "WORK-0001", title: "Payments v2"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			app, _ := initialized(t)
			for _, s := range c.setup {
				if code, _, e := run(t, app, s...); code != ExitOK {
					t.Fatalf("setup %v failed: %s", s, e)
				}
			}
			if len(c.setup) == 0 && c.args[1] != "new" {
				if code, _, e := run(t, app, "work-item", "new", "Payments v2", "--kind", "change"); code != ExitOK {
					t.Fatalf("setup failed: %s", e)
				}
			}
			code, out, errOut := run(t, app, c.args...)
			if code != ExitOK {
				t.Fatalf("exit = %d: %s", code, errOut)
			}
			if c.key != "" && !strings.Contains(out, c.key+" · ") {
				t.Errorf("the report did not name the key as a handle:\n%s", out)
			}
			if !strings.Contains(out, c.title) {
				t.Errorf("the report did not carry the title %q:\n%s", c.title, out)
			}
		})
	}
}

// A keyless record is titled, because that is the only handle it has.
func TestReportsNameKeylessRecordsByTitle(t *testing.T) {
	app, _ := initialized(t)
	for _, s := range [][]string{
		{"work-item", "new", "Payments v2", "--kind", "change"},
		{"outcome", "new", "The queue drains", "-w", "payments-v2"},
	} {
		if code, _, e := run(t, app, s...); code != ExitOK {
			t.Fatalf("setup %v failed: %s", s, e)
		}
	}
	_, out, _ := run(t, app, "outcome", "verify", "the-queue-drains", "proven", "-e", "ran it")
	if !strings.Contains(out, "The queue drains") {
		t.Errorf("the report did not carry the title:\n%s", out)
	}
	if strings.Contains(out, " · ") {
		t.Errorf("an outcome carries no key, so nothing should precede its title:\n%s", out)
	}
}

// commandExists walks the command tree for a verb path.
//
// Running the command with --help does NOT answer this: Cobra prints the
// parent's help and exits zero for an unknown subcommand, so the obvious probe
// passes on exactly the input it was written to catch.
func commandExists(a *App, path []string) bool {
	cur := newRootCommand(a)
	for i, name := range path {
		if strings.HasPrefix(name, "-") {
			return true // a flag ends the verb path
		}
		var next *cobra.Command
		for _, c := range cur.Commands() {
			if c.Name() == name {
				next = c
				break
			}
		}
		if next == nil {
			// Not a verb, so it is an argument — and whether that is legal is
			// the reached command's own rule. Treating every token as a verb
			// reported `outcome verify <ref>` as nonexistent, because the
			// reference is not a subcommand and was never meant to be.
			if cur.Args == nil {
				return true
			}
			err := cur.Args(cur, path[i:])
			if err == nil {
				return true
			}
			// NoArgs is how a noun says "that token was meant to be a verb",
			// and its message is Cobra's own. Any other complaint is about HOW
			// MANY arguments follow — which is wrong by construction here,
			// since the placeholders were stripped before the walk.
			return !strings.Contains(err.Error(), "unknown command")
		}
		cur = next
	}
	return true
}

// offeredCommands pulls the indented line under every lead-in.
func offeredCommands(msg string) []string {
	var out []string
	lines := strings.Split(msg, "\n")
	for i, l := range lines {
		if strings.HasSuffix(l, ":") && !strings.HasPrefix(l, " ") && i+1 < len(lines) {
			if c := strings.TrimSpace(lines[i+1]); c != "" {
				out = append(out, c)
			}
		}
	}
	return out
}
