package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newVerifyCommand(a *App) *cobra.Command {
	var evidence string

	cmd := &cobra.Command{
		Use:   "verify <outcome> <proven|disproven|inconclusive>",
		Short: "Record that an outcome has been confirmed",
		Long: "Records what a checker found, and the evidence it rests on.\n\n" +
			"The verdict is positional because recording proof has to be said out\n" +
			"loud — an outcome nobody checked and one somebody disproved are\n" +
			"different facts, and a command that could only say yes conflated them.\n\n" +
			"Verification accumulates: several actors confirming the same outcome is\n" +
			"the normal case, and a human entry raises the derived trust tier with no\n" +
			"special handling.",
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
			res, err := s.Verify(app.VerifyRequest{Ref: args[0], As: as, Evidence: evidence})
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "verified  %s (%s)\n", res.Path, as)
			if res.NoEvidence {
				fmt.Fprintln(out,
					"\nNo evidence recorded. An unbacked confirmation is the claim this\n"+
						"design distrusts most — pass --evidence next time.")
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&evidence, "evidence", "e", "", "what the confirmation rests on")
	return cmd
}
