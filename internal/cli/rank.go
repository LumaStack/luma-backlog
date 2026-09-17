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
	cmd.AddCommand(newRankRepairCommand(a))
	return cmd
}

// newRankRepairCommand renumbers every work item so all of them carry a rank.
//
// A subcommand of `rank` rather than a verb of its own: it changes the same
// field, for the same reason, and somebody looking for it will look there.
func newRankRepairCommand(a *App) *cobra.Command {
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "repair",
		Short: "Give every work item a rank, in creation order",
		Long: "Recomputes every work item's rank so that all of them carry one and every\n" +
			"prefix agrees with the status the record declares.\n\n" +
			"Records are numbered in the order they were created, within each status.\n" +
			"That makes the result a pure function of the corpus: two people repairing\n" +
			"the same state produce identical files, and re-running changes nothing.\n\n" +
			"It renumbers rather than filling gaps, so an ordering somebody chose by\n" +
			"hand is replaced. Rank again afterwards to restate it.\n\n" +
			"This rewrites many records at once. Run it with --dry-run first.",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		Example: "  luma-backlog rank repair --dry-run\n" +
			"  luma-backlog rank repair",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			res, err := s.RepairRanks(app.RepairRequest{DryRun: dryRun})
			if err != nil {
				return err
			}
			observe(cmd.ErrOrStderr(), res.Observations)

			out := cmd.OutOrStdout()
			for _, c := range res.Changed {
				from := c.From
				if from == "" {
					from = "unranked"
				}
				fmt.Fprintf(out, "%s  %s  %s → %s\n", verb(res.DryRun), c.Name, from, c.To)
			}
			for _, name := range res.Unplaceable {
				fmt.Fprintf(cmd.ErrOrStderr(),
					"luma-backlog: %s declares a status this project does not carry --- no rank can be computed for it\n", name)
			}
			if len(res.Changed) == 0 {
				fmt.Fprintf(out, "every one of the %d work items already carries the rank it should\n", res.Examined)
				return nil
			}
			fmt.Fprintf(out, "\n%d of %d changed%s\n", len(res.Changed), res.Examined, suffix(res.DryRun))
			return nil
		},
	}
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "show what would change and write nothing")
	return cmd
}

// verb labels a row by whether it happened.
func verb(dryRun bool) string {
	if dryRun {
		return "would rank"
	}
	return "ranked   "
}

func suffix(dryRun bool) string {
	if dryRun {
		return "; nothing was written --- re-run without --dry-run"
	}
	return ""
}
