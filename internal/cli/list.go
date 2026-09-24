package cli

import (
	"fmt"
	"io"
	"strings"
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
		onlyOpen bool
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

			f := app.Filter{Unit: unit, WorkItem: workItem, Status: status, Kind: kind, Open: onlyOpen}

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
				emptyListing(cmd.ErrOrStderr(), unit, f)
				return nil
			}
			w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "KEY\tSTATUS\tTITLE")
			for _, it := range res.Items {
				fmt.Fprintf(w, "%s\t%s\t%s\n", refOf(it), it.Status, it.Title)
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
	// The question anybody asks first, and the one that could not be expressed:
	// --status matches a single value, so "everything except closed" needed a
	// second language to answer (docs/design/mvp.md).
	cmd.Flags().BoolVar(&onlyOpen, "open", false, "only work that has not ended")

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

// emptyListing says why a listing came back empty.
//
// Printing nothing was the defect: it left somebody unable to tell a backlog
// with nothing in it from a filter that matched none of one — and those need
// opposite things said. An empty corpus needs the command that fills it; a
// filter that matched nothing needs to show the filter, so an empty result is
// distinguishable from a wrong question (policy/output-patterns).
//
// To stderr, for the reason a skip goes there: stdout stays a listing and
// nothing else, so a caller piping it is unaffected while a person still reads
// this.
func emptyListing(w io.Writer, unit string, f app.Filter) {
	unit = unitOr(unit)
	if narrowing := narrowing(f); narrowing != "" {
		fmt.Fprintf(w, "\nNo %ss match\n  %s\n\n", unitWords(unit), narrowing)
		fmt.Fprintf(w, "See them all with:\n  %s\n", listing(unit))
		return
	}
	fmt.Fprintf(w, "\nNo %ss yet\n\n", unitWords(unit))
	fmt.Fprintf(w, "Add one with:\n  %s\n", howToCreate(unit))
}

// unitOr resolves the bare `list`, which carries no unit and means work items.
func unitOr(unit string) string {
	if unit == "" {
		return app.WorkItem
	}
	return unit
}

// unitWords is a unit as a sentence says it — `work-item` is a type name, and
// "No work-items yet" reads as a field rather than as English. Every unit
// pluralizes with a trailing s, so callers add it.
func unitWords(unit string) string { return strings.ReplaceAll(unit, "-", " ") }

// narrowing renders the filter that ran, as the flags somebody would type.
// Empty when nothing narrowed the listing, which is what separates the two
// empty states.
func narrowing(f app.Filter) string {
	var parts []string
	if f.WorkItem != "" {
		parts = append(parts, "--work-item "+f.WorkItem)
	}
	if f.Status != "" {
		parts = append(parts, "--status "+f.Status)
	}
	if f.Kind != "" {
		parts = append(parts, "--kind "+f.Kind)
	}
	if f.Open {
		parts = append(parts, "--open")
	}
	return strings.Join(parts, " ")
}

// listing is the unnarrowed form of this listing. Work items reach it at the
// top level, since `list` is `work-item list` (root.go).
func listing(unit string) string {
	if unit == app.WorkItem {
		return "luma-backlog list"
	}
	return "luma-backlog " + unit + " list"
}

// howToCreate is the command that adds one, ready to type.
//
// An outcome, a task and an exploration belong to a work item and `new`
// refuses without one, so the flag is part of the command rather than
// something to discover from the refusal.
func howToCreate(unit string) string {
	cmd := "luma-backlog " + unit + ` new "<title>"`
	switch unit {
	case app.Outcome, app.Task, app.Exploration:
		cmd += " --work-item <key>"
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
		emptyListing(cmd.ErrOrStderr(), f.Unit, f)
		return nil
	}
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "KEY\tSTATUS\tTITLE")
	for _, n := range res.Nodes {
		fmt.Fprintf(w, "%s\t%s\t%s\n", refOf(n.View), n.Status, n.Title)
		for i, c := range n.Children {
			// The last child closes the branch, so the eye finds where one
			// work item ends without counting rows.
			branch := "├─"
			if i == len(n.Children)-1 {
				branch = "└─"
			}
			fmt.Fprintf(w, "%s %s\t%s\t%s\n", branch, marker(c.Type), c.Status, c.Title)
		}
	}
	return w.Flush()
}

// refOf is what a row is called. A work item shows its key --- the handle
// somebody says out loud. Anything else shows its name, since that is the only
// handle it has.
func refOf(v app.View) string {
	if v.Key != "" {
		return v.Key
	}
	return v.Name
}

// marker labels a child by what it is rather than naming it.
//
// Outcomes and tasks have no key, and their slug is their title in kebab case
// --- printing it puts the same sentence on the row twice and pushes every
// other column right. The type is the useful thing at this width; to act on a
// child, `task list --work-item <ref>` gives the handles.
func marker(unit string) string {
	switch unit {
	case app.Outcome:
		return "OUT"
	case app.Task:
		return "TASK"
	case app.Decision:
		return "DEC"
	case app.Exploration:
		return "EXP"
	}
	return strings.ToUpper(unit)
}
