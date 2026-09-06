package cli

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newNewCommand(a *App) *cobra.Command {
	var workItem, kind string
	var project bool

	cmd := &cobra.Command{
		Use:   "new <" + strings.Join(app.Units, "|") + "> <title>",
		Short: "Create a record",
		Long: "Creates a record from a title. Everything else is derived: the file name\n" +
			"from the title, the work item from where you are, the timestamp and\n" +
			"actor from the environment.\n\n" +
			"Running it twice with the same title leaves the first one alone.\n\n" +
			"A decision states its level: --work-item for one belonging to that work\n" +
			"item, --project for a standing rule. It is never inferred from where the\n" +
			"command was run.\n\n" +
			"--kind classifies a work item by what it produces:\n" +
			"  defect   a fix\n" +
			"  request  an answer you already have the standing to give\n" +
			"  idea     a classification — it becomes one of the others\n" +
			"  inquiry  understanding, and the work items that follow\n" +
			"  change   the work itself\n\n" +
			"bug and ask are stored as defect and request; review, audit,\n" +
			"investigation and spike are stored as inquiry.",
		Args:         cobra.ExactArgs(2),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNew(a, cmd, args[0], args[1], workItem, kind, project, cmd.Flags().Changed("work-item"))
		},
	}
	cmd.Flags().StringVarP(&workItem, "work-item", "w", "",
		"the work item this belongs to (default: derived from the working directory)")
	cmd.Flags().StringVarP(&kind, "kind", "k", "",
		"what sort of work item this is (default: none, meaning ordinary work)")
	cmd.Flags().BoolVar(&project, "project", false,
		"a decision at the project level, rather than one belonging to a work item")
	return cmd
}

func runNew(a *App, cmd *cobra.Command, unit, title, workItem, kind string, project, workItemGiven bool) error {
	s, err := open(a)
	if err != nil {
		return err
	}
	defer s.Close()

	res, err := s.Create(app.CreateRequest{
		Unit:          unit,
		Title:         title,
		WorkItem:      workItem,
		WorkItemGiven: workItemGiven,
		Kind:          kind,
		Project:       project,
	})
	if err != nil {
		return err
	}

	if res.Unclassified {
		fmt.Fprintf(cmd.ErrOrStderr(),
			"luma-backlog: no kind — recorded as unclassified.\n"+
				"  --kind defect   something broke\n"+
				"  --kind request  somebody asked\n"+
				"  --kind idea     a thought nobody can judge yet\n"+
				"  --kind inquiry  going to look, and it will produce work\n"+
				"  --kind change   none of those\n"+
				"Leave it blank only when nobody has looked at this yet.\n")
	}

	out := cmd.OutOrStdout()
	if res.Created {
		fmt.Fprintf(out, "created  %s\n", res.Path)
	} else {
		fmt.Fprintf(out, "exists   %s\n", res.Path)
	}
	return nil
}
