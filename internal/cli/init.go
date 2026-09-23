package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newInitCommand(a *App) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create a backlog in this repository",
		Long: "Creates .luma/ with the backlog bundle and records tier, and a\n" +
			"configuration file written out in full.\n\n" +
			"Safe to run again: nothing existing is overwritten, and anything\n" +
			"missing is created.",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			res, err := app.Init(a.Env, a.WorkingDir, a.Ceiling)
			if err != nil {
				return err
			}
			// The heading is the end state, which is true on a re-run as well;
			// the column beside each file says what that file needed. One shape
			// covers both, and it takes another row rather than a rewrite if
			// init ever writes a second file.
			state := "created"
			if !res.Created {
				state = "exists"
			}
			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "\nInitialized backlog\n  %-7s  %s\n\n", state, res.ConfigFile)
			fmt.Fprintf(out, "Add a work item with:\n  %s\n", howToCreate(app.WorkItem))
			return nil
		},
	}
}
