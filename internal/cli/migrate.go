package cli

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newMigrateCommand(a *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate <what>",
		Short: "Rewrite a corpus in ways ordinary commands do not",
		Long: "Migrations change the shape of records already written, across the whole\n" +
			"repository. They are grouped here because there will be more than one, and\n" +
			"because a corpus-wide rewrite is the least recoverable thing this tool does:\n" +
			"every one of them takes --dry-run, and none of them happens on its own.",
		SilenceUsage: true,
	}
	cmd.AddCommand(newMigrateKeysCommand(a))
	return cmd
}

func newMigrateKeysCommand(a *App) *cobra.Command {
	var dryRun, renumber, includeBareKeys bool

	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Move every work item to the configured key prefix",
		Long: "Moves every work item to the key prefix named in configuration, and\n" +
			"repoints everything that referred to it.\n\n" +
			"Only the prefix moves: WORK-0123 becomes BACK-0123 and keeps its number, so\n" +
			"the change stays reviewable and a reference held elsewhere stays recognizable.\n" +
			"A record already at the prefix is left alone.\n\n" +
			"Names are rewritten everywhere in the repository, because a work item's name\n" +
			"appears in prose, in documentation and in source comments, and a name does not\n" +
			"survive the move. Bare keys are left alone, because the old key keeps resolving\n" +
			"and because one may name a work item in a different project.",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		Example: "  luma-backlog migrate keys --dry-run\n" +
			"  luma-backlog migrate keys\n" +
			"  luma-backlog migrate keys --renumber",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			res, err := s.MigrateKeys(app.MigrateKeysRequest{
				DryRun:          dryRun,
				Renumber:        renumber,
				IncludeBareKeys: includeBareKeys,
			})
			if err != nil {
				return err
			}
			renderMigration(cmd.OutOrStdout(), cmd.ErrOrStderr(), res)
			return nil
		},
	}
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "report the whole change and write nothing")
	cmd.Flags().BoolVar(&renumber, "renumber", false,
		"give a record blocked by a collision the next free number")
	cmd.Flags().BoolVar(&includeBareKeys, "include-bare-keys", false,
		"also rewrite keys written without their slug")
	return cmd
}

// renderMigration writes the three things a reader needs, in the order they
// need them.
//
// **The mapping goes to stdout and everything else to stderr.** Whoever is
// updating an external system wants old-to-new and nothing else, pipeable. The
// blocked list and the bare keys are for the person running it.
func renderMigration(out, errOut io.Writer, res *app.MigrateKeysResult) {
	for _, m := range res.Moved {
		fmt.Fprintf(out, "%s\t%s\n", m.OldKey, m.NewKey)
	}

	if res.DryRun {
		fmt.Fprintf(errOut, "\nNothing was written --- this was a dry run.\n")
	}

	w := tabwriter.NewWriter(errOut, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "\nmoved\t%d\n", len(res.Moved))
	if res.AlreadyCorrect > 0 {
		fmt.Fprintf(w, "already correct\t%d\n", res.AlreadyCorrect)
	}
	fmt.Fprintf(w, "files rewritten\t%d\n", len(res.Files))
	w.Flush()

	if len(res.Blocked) > 0 {
		fmt.Fprintf(errOut, "\n%d record(s) could not migrate --- their keys are held:\n", len(res.Blocked))
		for _, b := range res.Blocked {
			held := "by"
			if b.Former {
				held = "formerly by"
			}
			fmt.Fprintf(errOut, "  %s is held %s %s\n", b.Key, held, b.HeldBy)
		}
		fmt.Fprintf(errOut, "\nGive them new numbers with:\n  luma-backlog migrate keys --renumber\n")
	}

	if len(res.BareKeys) > 0 {
		total := 0
		for _, f := range res.BareKeys {
			total += len(f.Keys)
		}
		fmt.Fprintf(errOut, "\n%d old key(s) remain written without a slug, in %d file(s):\n",
			total, len(res.BareKeys))
		for _, f := range res.BareKeys {
			fmt.Fprintf(errOut, "  %s\t%s\n", f.Path, strings.Join(f.Keys, " "))
		}
		fmt.Fprintf(errOut, "\nThey still resolve, and one may name another project's work.\n")
		fmt.Fprintf(errOut, "Rewrite them too with:\n  luma-backlog migrate keys --include-bare-keys\n")
	}
}
