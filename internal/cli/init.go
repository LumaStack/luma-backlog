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
			out := cmd.OutOrStdout()
			if res.Created {
				fmt.Fprintf(out, "created  %s\n", res.ConfigFile)
			} else {
				fmt.Fprintf(out, "exists   %s\n", res.ConfigFile)
			}
			fmt.Fprintf(out, "\nBacklog ready at %s\n", res.Path)
			return nil
		},
	}
}
