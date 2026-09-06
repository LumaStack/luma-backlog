package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newVerifyCommand(a *App) *cobra.Command {
	var evidence string

	cmd := &cobra.Command{
		Use:   "verify <outcome>",
		Short: "Record that an outcome has been confirmed",
		Long: "Adds a confirmation event, and the evidence it rests on.\n\n" +
			"Verification accumulates: several actors confirming the same outcome is\n" +
			"the normal case, and a human entry raises the derived trust tier with no\n" +
			"special handling.",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			res, err := s.Verify(app.VerifyRequest{Ref: args[0], Evidence: evidence})
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "verified  %s\n", res.Path)
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
