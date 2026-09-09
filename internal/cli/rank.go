package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newRankCommand(a *App) *cobra.Command {
	var (
		before, after string
		first, last   bool
	)

	cmd := &cobra.Command{
		Use:   "rank <work-item>",
		Short: "Reorder a work item",
		Long: "Reorders a work item within the records at its own workflow status.\n\n" +
			"Rank is work order, and workflow status dominates it: a record at a later\n" +
			"status is ahead of every record at an earlier one, whatever its rank. So\n" +
			"ranking reorders a record among the others at its own status, and nowhere else.\n\n" +
			"You say where; the tool chooses the ordering key. `set` refuses the rank\n" +
			"field for the same reason.",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		Example: "  luma-backlog work-item rank WORK-0031 --first\n" +
			"  luma-backlog work-item rank WORK-0031 --before WORK-0022",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := app.RankRequest{Ref: args[0]}
			set := 0
			switch {
			case first:
				req.Where, set = app.RankFirst, set+1
			}
			if last {
				req.Where, set = app.RankLast, set+1
			}
			if before != "" {
				req.Where, req.Neighbor, set = app.RankBefore, before, set+1
			}
			if after != "" {
				req.Where, req.Neighbor, set = app.RankAfter, after, set+1
			}
			if set > 1 {
				return usageErr("pass one of --first, --last, --before or --after")
			}

			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			res, err := s.Rank(req)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "ranked  %s (%s)\n", res.Path, res.Rank)
			return nil
		},
	}
	cmd.Flags().BoolVar(&first, "first", false, "first at its status")
	cmd.Flags().BoolVar(&last, "last", false, "last at its status")
	cmd.Flags().StringVar(&before, "before", "", "immediately before this record")
	cmd.Flags().StringVar(&after, "after", "", "immediately after this record")
	return cmd
}
