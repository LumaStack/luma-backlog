package cli

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"github.com/spf13/cobra"
)

func newListCommand(app *App) *cobra.Command {
	var (
		asJSON   bool
		workItem string
		status   string
	)

	cmd := &cobra.Command{
		Use:   "list [" + strings.Join(backlog.Units, "|") + "]",
		Short: "Read many records",
		Long:  "Lists records, optionally narrowed by unit, work item, or status.",
		Args:  cobra.MaximumNArgs(1),
		// Deliberately not an error when nothing matches: an empty backlog and
		// an over-narrow filter are both ordinary, and exiting non-zero would
		// make a caller treat "none yet" as a failure.
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			b, cfg, _, err := openBacklog(app)
			if err != nil {
				return err
			}
			defer b.Close()

			f := backlog.Filter{WorkItem: workItem, Status: status}
			if len(args) == 1 {
				if !backlog.IsUnit(args[0]) {
					return usageErr("unknown unit %q: expected one of %s",
						args[0], strings.Join(backlog.Units, ", "))
				}
				f.Unit = args[0]
			}

			items, err := backlog.List(b, f)
			if err != nil {
				return failure("%w", err)
			}

			out := cmd.OutOrStdout()
			if asJSON {
				rows := make([]itemJSON, 0, len(items))
				for _, it := range items {
					rows = append(rows, toItemJSON(it, cfg.DefaultStatusFor(it.Type())))
				}
				// An empty result is [] rather than null: a caller iterating
				// the response should not have to special-case nothing.
				return writeJSON(out, rows)
			}

			if len(items) == 0 {
				fmt.Fprintln(out, "No records match.")
				return nil
			}
			w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "TYPE\tSTATUS\tSLUG\tTITLE")
			for _, it := range items {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
					it.Type(), it.Status(cfg.DefaultStatusFor(it.Type())), it.Slug(), it.Title())
			}
			return w.Flush()
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "emit the listing as JSON")
	cmd.Flags().StringVarP(&workItem, "work-item", "w", "", "only records in this work item")
	cmd.Flags().StringVarP(&status, "status", "s", "", "only records with this workflow status")
	return cmd
}
