// Package cli assembles the command tree and maps failures onto the exit
// codes in docs/spec.md §9.4.
package cli

import (
	"fmt"
	"io"
	"os"

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

	if err := root.Execute(); err != nil {
		fmt.Fprintln(stderr, "luma-backlog:", err)
		return codeFor(err)
	}
	return ExitOK
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
	addGroups(root)

	root.AddCommand(inGroup(newInitCommand(app), groupExtra))
	addNouns(root, app)

	// show and set are top-level while the cross-type question is unsettled.
	// See nouns.go.
	root.AddCommand(inGroup(newShowCommand(app), groupCore))
	root.AddCommand(inGroup(newSetCommand(app), groupCore))

	// `list` is `work-item list`. The tool is called backlog; listing the
	// backlog means listing work items, which is the reading of the command
	// name and the overwhelmingly common case.
	root.AddCommand(inGroup(newListCommand(app, backlogUnit), groupCore))

	applyHelp(root)
	return root
}
