// Package cli assembles the command tree and maps failures onto the exit
// codes in docs/spec.md §9.4.
package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/lumastack/luma-backlog/internal/env"
	"github.com/spf13/cobra"
)

// workingDir reports where the process was started. This is the one place the
// package reaches outside itself for a path; everything after takes it as a
// value.
func workingDir() (string, error) { return os.Getwd() }

// version is set at build time via -ldflags.
var version = "dev"

// Exit codes are part of the published contract (docs/spec.md §9.4), so
// they are named rather than written as bare integers at call sites.
const (
	ExitOK       = 0
	ExitError    = 1 // unexpected; stop and surface it
	ExitUsage    = 2 // fix the invocation; never retry unchanged
	ExitNotFound = 3
	ExitConflict = 4 // the record changed underneath; re-read and retry
	ExitRefused  = 5 // a validated act did not pass its check
	// There is deliberately no sixth. `6` was reserved for "already taken",
	// and taking does not ship (ADR-0008) — a code held for a feature nobody
	// has designed is a promise about a shape nobody chose. Adding one is
	// additive and removing one is breaking (spec.md §9.9), so six is the
	// reversible direction.
)

// App is what a command runs against: the ambient facts, gathered once, so no
// command reaches for them itself.
type App struct {
	Env env.Env
	// WorkingDir is where discovery starts.
	WorkingDir string
	// Ceiling bounds the upward walk. Empty means the filesystem root; tests
	// set it so an escape fails loudly rather than finding the developer's
	// own checkout and quietly succeeding.
	Ceiling string
}

// Main runs the command tree and returns a process exit code. Streams are
// passed in rather than reached for, so tests drive it without touching the
// real process.
func Main(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	wd, err := workingDir()
	if err != nil {
		fmt.Fprintln(stderr, "luma-backlog:", err)
		return ExitError
	}
	return Run(&App{Env: env.New(), WorkingDir: wd}, args, stdin, stdout, stderr)
}

// Run is Main with the ambient state supplied, which is how tests drive it.
func Run(app *App, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	root := newRootCommand(app)
	// Cobra falls back to os.Args when given nil, which would make the whole
	// point of passing arguments in — not touching the process — quietly
	// false. Found when `go test -update` leaked its own flag into the tool.
	if args == nil {
		args = []string{}
	}
	root.SetArgs(args)
	root.SetIn(stdin)
	root.SetOut(stdout)
	root.SetErr(stderr)

	// ExecuteC rather than Execute: it returns the command that was running,
	// which is what lets a parsing failure name the right help page and draw
	// suggestions from the right set of verbs.
	cmd, err := root.ExecuteC()
	if err != nil {
		if r, ok := parseRefusal(cmd, err); ok {
			renderRefusal(stderr, r)
		} else {
			report(stderr, err)
		}
		return codeFor(err)
	}
	return ExitOK
}

var unknownCommand = regexp.MustCompile(`^unknown command "([^"]+)"`)

// parseRefusal turns an argument- or flag-parsing failure into a refusal.
//
// These never reach the application layer — Cobra rejects them before a
// command runs — so its prose is what a reader gets unless something
// translates it. `unknown command "work-tiem" for "luma-backlog"` was the
// reported case: correct, and in nobody's house style.
//
// Matching on Cobra's message text is the part to dislike. There is no typed
// error to match instead, and the alternative is reimplementing the parsing to
// find out what it already knows.
func parseRefusal(cmd *cobra.Command, err error) (app.Refusal, bool) {
	if cmd == nil {
		return app.Refusal{}, false
	}
	msg := err.Error()
	help := cmd.CommandPath() + " --help"

	if m := unknownCommand.FindStringSubmatch(msg); m != nil {
		r := app.Refusal{
			Problem: "Unknown command " + m[1],
			LeadIn:  "See every command with",
			Command: help,
		}
		// Cobra already knows the near misses; asking it beats a second
		// distance function that can disagree with the one in the help.
		//
		// It suggests nothing while the distance is zero, and sets its own
		// default only on the way into Execute — after the error we are
		// holding. Two is that default, restated here rather than waited for.
		if cmd.SuggestionsMinimumDistance <= 0 {
			cmd.SuggestionsMinimumDistance = 2
		}
		if s := cmd.SuggestionsFor(m[1]); len(s) > 0 {
			r.Detail = []string{"did you mean " + strings.Join(s, ", ")}
		}
		return r, true
	}

	switch {
	case strings.HasPrefix(msg, "unknown flag:"),
		strings.HasPrefix(msg, "unknown shorthand flag:"),
		strings.HasPrefix(msg, "flag needs an argument:"),
		strings.HasPrefix(msg, "invalid argument"):
		return app.Refusal{
			Problem: capitalize(msg),
			LeadIn:  "See what " + cmd.Name() + " accepts with",
			Command: help,
		}, true
	case strings.Contains(msg, "arg(s), received"), strings.Contains(msg, "arg(s), only received"):
		return app.Refusal{
			Problem: "Wrong number of arguments",
			Detail:  []string{msg},
			LeadIn:  "See how " + cmd.Name() + " is called with",
			Command: help,
		}, true
	}
	return app.Refusal{}, false
}

// capitalize raises the first letter, so a borrowed message reads as a heading
// rather than as the middle of somebody else's sentence.
func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.TrimSuffix(s[1:], ":")
}

// report writes a failure to stderr.
//
// Most of them are one sentence and take the tool's name as a prefix, which is
// the convention and what keeps them readable when several tools write to one
// stream.
//
// A refusal somebody has to act on is laid out instead — what is wrong, the
// detail beneath it, and the command that resolves it. A prefix would sit on
// the heading and fight the layout, so these render whole. The layer hands over
// the facts and none of the shape (ADR-0004).
//
// A command is introduced by a colon and takes the line under it. No backticks:
// they are shell syntax, so a copied one becomes command substitution rather
// than the command — and the indentation delimits it without spending a
// character that can misfire.
func report(w io.Writer, err error) {
	var refusal *app.Refusal
	if errors.As(err, &refusal) {
		renderRefusal(w, *refusal)
		return
	}

	var notInit *app.NotInitialized
	if errors.As(err, &notInit) {
		fmt.Fprintf(w, "\nNo backlog found\n  missing %s\n\n"+
			"Get started with:\n  luma-backlog init\n", notInit.ConfigFile)
		return
	}

	var noRepo *app.NoRepository
	if errors.As(err, &noRepo) {
		fmt.Fprint(w, "\nNo git repository found\n  a backlog belongs to a repository\n\n"+
			"Start one with:\n  git init\n")
		return
	}

	var noDisposition *app.DispositionRequired
	if errors.As(err, &noDisposition) {
		// The choices sit in the action line rather than the detail block:
		// they are what somebody types, and the detail slot is worth more to
		// the reason the refusal exists at all.
		fmt.Fprintf(w, "\nDisposition required\n"+
			"  cancelled work and completed work look identical without one\n\n"+
			"Close it with:\n  luma-backlog work-item close %s <%s>\n",
			noDisposition.Ref, strings.Join(noDisposition.Choices, "|"))
		return
	}

	var cannot *app.CannotComplete
	if errors.As(err, &cannot) {
		reportCannotComplete(w, cannot)
		return
	}

	var ambiguous *app.Ambiguous
	if errors.As(err, &ambiguous) {
		fmt.Fprintf(w, "\nAmbiguous reference %s\n", ambiguous.Ref)
		for _, p := range ambiguous.Paths {
			fmt.Fprintf(w, "  %s\n", p)
		}
		fmt.Fprint(w, "\nName one of them, or pass the key.\n")
		return
	}

	var noMatch *app.NoMatch
	if errors.As(err, &noMatch) {
		// No detail line — the heading already carries the specifics, and a
		// section with nothing in it is left out rather than left empty.
		//
		// "record" where the operation took any: `show` reaches decisions and
		// PROJECT.md too, and naming the wrong unit sends somebody looking in
		// the wrong listing.
		what := "record"
		if noMatch.Unit != "" {
			what = unitWords(noMatch.Unit)
		}
		fmt.Fprintf(w, "\nNo %s matches %s\n\n"+
			"See what exists with:\n  %s\n", what, noMatch.Ref, listing(unitOr(noMatch.Unit)))
		return
	}

	fmt.Fprintln(w, "luma-backlog:", err)
}

// renderRefusal lays out a refusal: heading, indented detail, then either a
// command under a colon or closing prose.
//
// This is the one place the shape lives. Every message that used to build its
// own multi-line string drifted from every other one, because nothing compared
// them — and the reader met a different layout per command.
func renderRefusal(w io.Writer, r app.Refusal) {
	fmt.Fprintf(w, "\n%s\n", r.Problem)
	for _, d := range r.Detail {
		fmt.Fprintf(w, "  %s\n", d)
	}
	switch {
	case r.Command != "":
		// A colon opens the command and the line break closes it. No
		// backticks: they are shell syntax, so a copied one becomes command
		// substitution rather than the command.
		fmt.Fprintf(w, "\n%s:\n  %s\n", r.LeadIn, r.Command)
	case r.Note != "":
		fmt.Fprintf(w, "\n%s\n", r.Note)
	}
}

// reportCannotComplete renders a completed close the work item's own
// declarations refused.
//
// The closing line names a CONDITION rather than a command, and that is the
// whole difference from every other refusal here: the invocation was correct,
// so offering a corrected one would send somebody to run something that fails
// identically. Where there is an escape, it is named with what it costs.
func reportCannotComplete(w io.Writer, e *app.CannotComplete) {
	switch e.Gap {
	case app.OutcomesUnreadable:
		fmt.Fprintf(w, "\n%s cannot be completed\n", e.Ref)
		fmt.Fprintf(w, "  %s could not be read, so there is no count\n", plural(e.Count, "outcome"))
	case app.NoOutcomes:
		fmt.Fprintf(w, "\n%s cannot be completed\n", e.Ref)
		fmt.Fprint(w, "  no outcomes, so nothing says it was completed\n")
	case app.OutcomesUnproven:
		fmt.Fprintf(w, "\n%s cannot be completed\n", e.Ref)
		fmt.Fprintf(w, "  %d of %d outcomes are not proven\n", e.Count, e.Total)
	case app.TasksOpen:
		fmt.Fprintf(w, "\n%s cannot be completed\n", e.Ref)
		fmt.Fprintf(w, "  %s never reached a terminal status\n", plural(e.Count, "task"))
	}
	for _, n := range e.Names {
		fmt.Fprintf(w, "    %s\n", n)
	}

	fmt.Fprintln(w)
	switch e.Gap {
	case app.OutcomesUnreadable:
		fmt.Fprint(w, "An unreadable outcome might be failing, and nothing here can tell.\n"+
			"Repair the file, or close with a disposition that claims nothing about evidence.\n")
	case app.NoOutcomes:
		fmt.Fprint(w, "Declare what done meant, or close with a different disposition.\n")
	case app.OutcomesUnproven:
		fmt.Fprint(w, "Verify them, or close with a different disposition. Abandoning one records\n"+
			"why it is unmet and does not clear this — a completed close over an unmet\n"+
			"outcome needs --force, and the count will say so.\n")
	case app.TasksOpen:
		fmt.Fprint(w, "Close them — any disposition is fine, they may have failed or been\n"+
			"cancelled — or pass --force.\n")
	}
}

// report renders "<verb>  <subject>" and the path beneath it.
//
// The path used to BE the subject, which meant reading the key out of a
// filename and the title out of a slug. showing-records asks that a key arrive
// with its title at least once in ephemeral output, and the key is the handle
// for whatever command comes next — so it leads, and the path becomes detail.
//
// An outcome and a task carry no key, so they are titled and nothing else;
// that is the same rule a listing follows.
func reportSubject(w io.Writer, verb string, s app.Subject, extra ...string) {
	name := s.Title
	if s.Key != "" {
		name = s.Key + " · " + s.Title
	}
	if name == "" {
		name = s.Path // nothing to call it by; the path is better than blank
	}
	fmt.Fprintf(w, "%s  %s\n", verb, name)
	if s.Path != "" && name != s.Path {
		fmt.Fprintf(w, "  %s\n", s.Path)
	}
	for _, e := range extra {
		fmt.Fprintf(w, "  %s\n", e)
	}
}

// plural renders a count with its noun, so a message reads as English.
func plural(n int, noun string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, noun)
	}
	return fmt.Sprintf("%d %ss", n, noun)
}

func newRootCommand(app *App) *cobra.Command {
	root := &cobra.Command{
		Use:     "luma-backlog",
		Short:   "A git-native backlog, worked by people and agents at the same time",
		Long:    "A backlog that lives inside your git repository as plain markdown.",
		Version: version,
		// Without this, Cobra treats an unknown command as a positional
		// argument and exits 0 — found by the test below on its first run.
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// With no arguments this will open the board (docs/spec.md §11).
			// Until the board exists, print help: the no-argument case is a
			// person, and help is the honest thing to show one. The earlier
			// placeholder said no commands were implemented, which stopped
			// being true and nothing caught it.
			return cmd.Help()
		},
	}
	// Cobra sorts commands alphabetically, which throws away two orderings the
	// code already states: the core verbs are listed by how often they are
	// reached for, and `corpus.Units` is the model's own hierarchy with the
	// primary unit first. Alphabetical is a third ordering that means nothing
	// and silently overrules both.
	cobra.EnableCommandSorting = false

	addGroups(root)

	root.AddCommand(inGroup(newInitCommand(app), groupExtra))
	root.AddCommand(inGroup(newMigrateCommand(app), groupExtra))
	addNouns(root, app)

	// Ordered by how often each is reached for, not alphabetically. `list` and
	// `show` are the two reads and belong together; `set` writes fields; `rank`
	// and `transition` are the two operations, which is also the order somebody
	// meets them in.
	//
	// show and set are top-level while the cross-type question is unsettled.
	// See nouns.go.
	//
	// `list` is `work-item list`. The tool is called backlog; listing the
	// backlog means listing work items, which is the reading of the command
	// name and the overwhelmingly common case.
	root.AddCommand(inGroup(newListCommand(app, backlogUnit), groupCore))
	root.AddCommand(inGroup(newShowCommand(app), groupCore))
	root.AddCommand(inGroup(newSetCommand(app), groupCore))

	// `transition` and `rank` are `work-item transition` and `work-item rank`,
	// for the same reason: only work items carry a workflow status or a rank,
	// so the noun adds a word and removes no ambiguity. It is also the pair
	// typed most often, and `rank` reading as absent because it was only
	// reachable under the noun is the observed cost of not doing this.
	root.AddCommand(inGroup(newRankCommand(app), groupCore))
	root.AddCommand(inGroup(newTransitionCommand(app), groupCore))

	applyHelp(root)
	return root
}
