package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newAssertCommand(a *App) *cobra.Command {
	return &cobra.Command{
		Use:   "assert <outcome> <succeeded|failed>",
		Short: "Record what the doer claims",
		Long: "Adds the doer's claim to an outcome. It never gates anything —\n" +
			"closing a work item as completed reads the checker's verdict and\n" +
			"not this, because gating on the doer's claim would let a doer\n" +
			"clear their own work.\n\n" +
			"Claims accumulate. A second one is a second attempt, and the\n" +
			"sequence is the record: failed then succeeded is a different\n" +
			"history from succeeded first time.",
		// Two, but the second may be missing so the failure names the claims
		// rather than counting arguments.
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
			res, err := s.Assert(app.AssertRequest{Ref: args[0], As: as})
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "asserted  %s (%s)\n", res.Path, as)
			if res.Attempts > 1 {
				fmt.Fprintf(out, "Attempt %d — the earlier claims are still on the record.\n", res.Attempts)
			}
			if res.Contradicts {
				fmt.Fprintln(out,
					"\nA checker has already recorded a verdict on this outcome. Both\n"+
						"stand; the disagreement is the thing worth seeing.")
			}
			return nil
		},
	}
}
