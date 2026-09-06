package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newCloseCommand(a *App) *cobra.Command {
	var reason string

	cmd := &cobra.Command{
		Use:   "close <work-item>",
		Short: "End a work item, recording why",
		Long: "Work ends for more reasons than success, so the terminal state is\n" +
			"\"closed\" and every closing records why.\n\n" +
			"Only --reason delivered is checked against the outcomes. The others\n" +
			"close freely: gating cancellation on completion would make it\n" +
			"impossible to stop work precisely because it was unfinished.",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			res, err := s.CloseWorkItem(app.CloseRequest{Ref: args[0], Reason: reason})
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
		"delivered, canceled, superseded, or abandoned")
	return cmd
}
