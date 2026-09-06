package cli

import (
	"fmt"
	"text/tabwriter"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

// newListCommand builds a listing. An empty unit builds the top-level `list`,
// which reads every record type --- including ones no noun can reach, such as
// the project record. A named unit builds `<noun> list`, which is the same
// command narrowed, the way `git log <path>` narrows `git log`.
func newListCommand(a *App, unit string) *cobra.Command {
	var (
		asJSON   bool
		workItem string
		status   string
		kind     string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: listShort(unit),
		Long:  listLong(unit),
		Args:  cobra.NoArgs,
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

			f := app.Filter{Unit: unit, WorkItem: workItem, Status: status, Kind: kind}

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
	cmd.Flags().StringVarP(&status, "status", "s", "", "only records with this workflow status")

	// A work item does not belong to a work item, so the filter has nothing to
	// narrow there. --kind classifies work items and means nothing elsewhere.
	if unit != app.WorkItem {
		cmd.Flags().StringVarP(&workItem, "work-item", "w", "", "only records in this work item")
	}
	if unit == "" || unit == app.WorkItem {
		cmd.Flags().StringVarP(&kind, "kind", "k", "", "only work items of this kind")
	}
	return cmd
}

func listShort(unit string) string {
	if unit == "" {
		return "Read many records, of every type"
	}
	return "Read many " + unit + " records"
}

func listLong(unit string) string {
	if unit == "" {
		return "Lists every record, whatever its type --- including records that belong to\n" +
			"no unit and so cannot be reached through a noun.\n\n" +
			"To list one type, use its noun: `luma-backlog work-item list`."
	}
	return "Lists " + unit + " records, optionally narrowed by work item, status, or kind."
}
