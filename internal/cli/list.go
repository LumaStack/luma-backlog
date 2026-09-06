package cli

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newListCommand(a *App) *cobra.Command {
	var (
		asJSON   bool
		workItem string
		status   string
		kind     string
	)

	cmd := &cobra.Command{
		Use:   "list [" + strings.Join(app.Units, "|") + "]",
		Short: "Read many records",
		Long:  "Lists records, optionally narrowed by unit, work item, status, or kind.",
		Args:  cobra.MaximumNArgs(1),
		// Deliberately not an error when nothing matches: an empty backlog and
		// an over-narrow filter are both ordinary, and exiting non-zero would
		// make a caller treat "none yet" as a failure.
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			f := app.Filter{WorkItem: workItem, Status: status, Kind: kind}
			if len(args) == 1 {
				f.Unit = args[0]
			}

			res, err := s.List(f)
			if err != nil {
				return err
			}

			// To stderr, always: stdout stays a clean listing and --json stays
			// parseable, so a caller piping the output is unaffected while
			// still being told.
			observe(cmd.ErrOrStderr(), res.Observations)

			out := cmd.OutOrStdout()
			if asJSON {
				rows := make([]itemJSON, 0, len(res.Items))
				for _, it := range res.Items {
					rows = append(rows, toItemJSON(it))
				}
				return writeJSON(out, rows)
			}

			if len(res.Items) == 0 {
				return nil
			}
			w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "TYPE\tSTATUS\tID\tTITLE")
			for _, it := range res.Items {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", it.Type, it.Status, it.Name, it.Title)
			}
			return w.Flush()
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "emit the listing as JSON")
	cmd.Flags().StringVarP(&workItem, "work-item", "w", "", "only records in this work item")
	cmd.Flags().StringVarP(&status, "status", "s", "", "only records with this workflow status")
	cmd.Flags().StringVarP(&kind, "kind", "k", "", "only work items of this kind")
	return cmd
}
