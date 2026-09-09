package cli

import (
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newTransitionCommand(a *App) *cobra.Command {
	var ifUnchanged, reason string
	var force bool

	cmd := &cobra.Command{
		Use:   "transition <work-item> <status>",
		Short: "Change a work item's workflow status",
		Long: "Changes where a work item sits on the workflow ladder.\n\n" +
			"A status change is an operation rather than a field write: it writes\n" +
			"workflow_status and rank together, so no record can be left with the two\n" +
			"disagreeing. `set` refuses the field for that reason, the way it already\n" +
			"refuses rank.\n\n" +
			"A transition re-enqueues the record at the back of its destination. Rank\n" +
			"before transitioning, not after.\n\n" +
			"Closing is a different command. It has its own refusals and its own\n" +
			"vocabulary, and this will not take you there.",
		Args:         cobra.ExactArgs(2),
		SilenceUsage: true,
		Aliases:      []string{"move"},
		Example: "  luma-backlog work-item transition WORK-0031 preparing\n" +
			"  luma-backlog work-item transition WORK-0031 prepared",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			res, err := s.Transition(app.TransitionRequest{
				Ref:         args[0],
				To:          args[1],
				IfUnchanged: ifUnchanged,
				Reason:      reason,
				Force:       force,
			})
			if err != nil {
				return err
			}
			// Only work items are ranked, so a task has none to report.
			if res.Rank != "" {
				fmt.Fprintf(cmd.OutOrStdout(), "%s  %s → %s (%s)\n",
					res.Path, res.From, res.To, res.Rank)
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "%s  %s → %s\n",
					res.Path, res.From, res.To)
			}
			if res.Journaled {
				fmt.Fprintf(cmd.OutOrStdout(), "journaled the reason\n")
			}
			// Advice goes to stderr at exit 0: the crossing happened, and
			// stdout stays clean for a caller piping it.
			for _, a := range res.Advice {
				fmt.Fprintf(cmd.ErrOrStderr(), "luma-backlog: %s\n", a)
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&force, "force", false,
		"start work with no outcomes anyway")
	cmd.Flags().StringVar(&reason, "reason", "",
		"why, in your words --- appended to the work item's journal")
	cmd.Flags().StringVar(&ifUnchanged, "if-unchanged", "",
		"the hash from show --json; refuses if the record changed since")
	return cmd
}
