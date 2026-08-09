// Package cli assembles the command tree and maps failures onto the exit
// codes in docs/SPEC.md §9.4.
package cli

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// version is set at build time via -ldflags.
var version = "dev"

// Exit codes are part of the published contract (docs/SPEC.md §9.4), so
// they are named rather than written as bare integers at call sites.
const (
	ExitOK       = 0
	ExitError    = 1 // unexpected; stop and surface it
	ExitUsage    = 2 // fix the invocation; never retry unchanged
	ExitNotFound = 3
	ExitConflict = 4 // the record changed underneath; re-read and retry
	ExitRefused  = 5 // a validated act did not pass its check
	ExitClaimed  = 6
)

// Main runs the command tree and returns a process exit code. Streams are
// passed in rather than reached for, so tests drive it without touching the
// real process.
func Main(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	root := newRootCommand()
	root.SetArgs(args)
	root.SetIn(stdin)
	root.SetOut(stdout)
	root.SetErr(stderr)

	if err := root.Execute(); err != nil {
		return ExitUsage
	}
	return ExitOK
}

func newRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "luma-backlog",
		Short: "A git-native backlog, worked by people and agents at the same time",
		Long: "A backlog that lives inside your git repository as plain markdown.\n" +
			"Records conform to the Luma Knowledge Format.",
		Version: version,
		// Without this, Cobra treats an unknown command as a positional
		// argument and exits 0 — found by the test below on its first run.
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// With no arguments this will open the board (docs/SPEC.md §11).
			// Until that exists, say so rather than printing help and implying
			// there is nothing here.
			fmt.Fprintln(cmd.OutOrStdout(), "luma-backlog "+version+" — no commands are implemented yet.")
			fmt.Fprintln(cmd.OutOrStdout(), "See docs/SPEC.md for the design, and .backlog/ for what is being built.")
			return nil
		},
	}
	return root
}
