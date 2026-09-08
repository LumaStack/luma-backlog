package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newAbandonCommand(a *App) *cobra.Command {
	var reason string

	cmd := &cobra.Command{
		Use:   "abandon <outcome>",
		Short: "Record that an outcome is no longer required",
		Long: "A decision, not a finding. A checker determines whether a condition\n" +
			"holds; abandoning says the condition is no longer asked — which is\n" +
			"why verify cannot record it.\n\n" +
			"It does not make the work item completable. The outcome stays in the\n" +
			"count as unmet, a close still needs --force, and the arithmetic still\n" +
			"disagrees with the claim. That is the point: both roads here are\n" +
			"failures — a requirement we could not state, or one we committed to\n" +
			"and dropped — and both are worth a retrospective.",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			res, err := s.Abandon(app.AbandonRequest{Ref: args[0], Reason: reason})
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "abandoned  %s\n", res.Path)
			fmt.Fprintln(out, "Still counted as unmet. This explains the gap; it does not close it.")
			if res.Verdicts > 0 {
				fmt.Fprintf(out, "%d verdict(s) already recorded — they stand.\n", res.Verdicts)
			}
			if reason == "" {
				fmt.Fprintln(out,
					"\nNo reason recorded. A retrospective asks which of two things happened —\n"+
						"a requirement nobody could state, or one we dropped — and nothing here says.")
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&reason, "reason", "r", "", "why it was given up on")
	return cmd
}
