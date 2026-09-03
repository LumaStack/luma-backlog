package cli

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"github.com/spf13/cobra"
)

func newCloseCommand(app *App) *cobra.Command {
	var reason string

	cmd := &cobra.Command{
		Use:   "close <work-item>",
		Short: "End a work item, recording why",
		Long: "Work ends for more reasons than success, so the terminal state is\n" +
			"\"closed\" and every closing records why.\n\n" +
			"Only --reason delivered is checked against the outcomes. The others\n" +
			"close freely: gating cancellation on completion would make it\n" +
			"impossible to stop work precisely because it was unfinished.",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClose(app, cmd, args[0], reason)
		},
	}
	cmd.Flags().StringVarP(&reason, "reason", "r", "",
		"delivered, canceled, superseded, or abandoned")
	return cmd
}

func runClose(app *App, cmd *cobra.Command, ref, reason string) error {
	b, _, _, err := openBacklog(app)
	if err != nil {
		return err
	}
	defer b.Close()

	if reason == "" {
		return usageErr("--reason is required: %s\n\n"+
			"Closing silently is how a backlog loses its own history — cancelled work\n"+
			"and delivered work look identical afterwards.", reasonList())
	}
	if !backlog.IsCloseReason(reason) {
		return usageErr("unknown reason %q: expected %s", reason, reasonList())
	}

	it, err := backlog.Resolve(b, ref)
	if err != nil {
		return coded{ExitNotFound, err}
	}
	if it.Type() != backlog.WorkItem {
		return usageErr("%s is a %s — close applies to a work item", it.Slug(), it.Type())
	}

	c, err := backlog.CompletionOf(b, it.Slug())
	if err != nil {
		return failure("%w", err)
	}

	// The tool's only refusal, and it holds a caller to their OWN declarations
	// rather than to an opinion of its own (docs/spec.md §5.0).
	if backlog.CloseReason(reason).GatedOnCompletion() {
		if !c.Complete() && len(c.Live) == 0 {
			return coded{ExitRefused, fmt.Errorf(
				"%s has no outcomes, so there is nothing that says it was delivered.\n"+
					"Declare what done means, or close with a different reason.", it.Slug())}
		}
		if len(c.Unpassing) > 0 {
			var names []string
			for _, o := range c.Unpassing {
				names = append(names, "  "+o.Slug())
			}
			return coded{ExitRefused, fmt.Errorf(
				"%s cannot be delivered: %d of %d outcomes have no evidence.\n%s\n\n"+
					"Verify them, retire the ones that no longer apply, or close with a\n"+
					"different reason.",
				it.Slug(), len(c.Unpassing), len(c.Live), strings.Join(names, "\n"))}
		}
	}

	it.Record.Set("workflow_status", "closed")
	if err := it.Record.SetRaw("closed", fmt.Sprintf("{on: %s, reason: %s, by: %s}",
		app.Env.Today(), reason, app.Env.Actor.String())); err != nil {
		return failure("%w", err)
	}
	if err := it.Record.SetRaw("modified",
		"{by: "+app.Env.Actor.String()+", at: "+app.Env.Now()+"}"); err != nil {
		return failure("%w", err)
	}

	out, err := it.Record.Bytes()
	if err != nil {
		return failure("%w", err)
	}
	if err := b.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return failure("%w", err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "closed  %s (%s)\n", it.Path, reason)
	if len(c.Retired) > 0 {
		fmt.Fprintf(cmd.OutOrStdout(), "%d retired outcome(s) were excluded from the count.\n", len(c.Retired))
	}
	return nil
}

func reasonList() string {
	var s []string
	for _, r := range backlog.CloseReasons {
		s = append(s, string(r))
	}
	return strings.Join(s, ", ")
}
