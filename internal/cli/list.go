package cli

import (
	"fmt"
	"text/tabwriter"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

// newListCommand builds a listing for one record type. It is registered under
// each noun, and once more at the top level for work items --- see root.go.
//
// There is deliberately no listing across types. A run of work items, outcomes,
// tasks, decisions and explorations interleaved is not something anyone asked
// for; it was what the code happened to do when the type was an argument.
func newListCommand(a *App, unit string) *cobra.Command {
	var (
		asJSON   bool
		asTree   bool
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

			if asTree {
				return runTree(cmd, s, f, asJSON)
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

	// Only work items have anything hanging off them, so only they can be
	// shown as a tree.
	if unit == app.WorkItem {
		cmd.Flags().BoolVar(&asTree, "tree", false, "show each work item's outcomes and tasks beneath it")
	}
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
	return "Read many " + unit + " records"
}

func listLong(unit string) string {
	return "Lists " + unit + " records, optionally narrowed by work item, status, or kind."
}

// runTree renders a listing with each work item's outcomes and tasks beneath
// it. The children are not filtered --- see app.Session.Tree.
func runTree(cmd *cobra.Command, s *app.Session, f app.Filter, asJSON bool) error {
	res, err := s.Tree(f)
	if err != nil {
		return err
	}
	observe(cmd.ErrOrStderr(), res.Observations)

	out := cmd.OutOrStdout()
	if asJSON {
		rows := make([]nodeJSON, 0, len(res.Nodes))
		for _, n := range res.Nodes {
			row := nodeJSON{itemJSON: toItemJSON(n.View)}
			for _, c := range n.Children {
				row.Children = append(row.Children, toItemJSON(c))
			}
			rows = append(rows, row)
		}
		return writeJSON(out, rows)
	}

	if len(res.Nodes) == 0 {
		return nil
	}
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "TYPE\tSTATUS\tID\tTITLE")
	for _, n := range res.Nodes {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", n.Type, n.Status, n.Name, n.Title)
		for i, c := range n.Children {
			// The last child closes the branch, so the eye can tell where one
			// work item ends and the next begins without counting indents.
			branch := "├─"
			if i == len(n.Children)-1 {
				branch = "└─"
			}
			fmt.Fprintf(w, "%s %s\t%s\t%s\t%s\n", branch, c.Type, c.Status, c.Name, c.Title)
		}
	}
	return w.Flush()
}
