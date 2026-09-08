package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newCloseCommand(a *App) *cobra.Command {
	var reason string

	cmd := &cobra.Command{
		Use:   "close <work-item> <completed|rejected|canceled|superseded>",
		Short: "End a work item, recording why",
		Long: "Work ends for more reasons than success, so the terminal state is\n" +
			"\"closed\" and every closing records which ending it was.\n\n" +
			"Only completed is checked against the outcomes. The others close\n" +
			"freely: gating cancellation on completion would make it impossible\n" +
			"to stop work precisely because it was unfinished.\n\n" +
			"The disposition is positional because it is not optional — a close\n" +
			"without one is not a close. --reason carries prose, if there is any.",
		// Two, but the second is allowed to be missing so the failure is a
		// list of what was expected rather than cobra counting arguments.
		Args:         cobra.RangeArgs(1, 2),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			as := ""
			if len(args) == 2 {
				as = args[1]
			}
			res, err := s.CloseWorkItem(app.CloseRequest{Ref: args[0], As: as, Reason: reason})
			if err != nil {
				return err
			}

			observe(cmd.ErrOrStderr(), res.Observations)

			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "closed  %s (%s)\n", res.Path, res.Reason)
			if res.Retired > 0 {
				fmt.Fprintf(out, "%d retired outcome(s) were excluded from the count.\n", res.Retired)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&reason, "reason", "r", "",
		"why, in your words — free prose, not one of the dispositions")
	return cmd
}
