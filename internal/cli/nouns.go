package cli

import (
	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

// The command tree is noun then verb (docs/spec.md §9.1, ADR-0006): a record
// type is never a positional argument. The noun a verb runs under is passed to
// its constructor, so each verb has one implementation rather than one per
// noun-verb pair --- copying a verb five times is how the nouns drift apart.
//
// `show` and `set` are deliberately absent. They reach records that are not
// creatable units --- `.luma/PROJECT.md` is type `luma/project`, and
// `corpus.Units` is what can be *created* --- so scoping them to a noun would
// strand every such record. They keep their top-level form until that is
// settled. See the work item's task, "Decide where cross-type listing lives".
//
// `list` is here as well as at the top level, and the two are not duplicates.
// The top-level one reads every type, including those same unreachable
// records; this one is that command narrowed, the way `git log <path>` narrows
// `git log`. What went is the redundant path --- `list <unit>` as a positional
// argument, which said the same thing twice.

// nounHelp is what each noun is for, shown in the root command's listing.
var nounHelp = map[string]string{
	app.WorkItem:    "A piece of work, and everything hanging off it",
	app.Outcome:     "What must be true for work to be done",
	app.Task:        "An attempt at part of the work",
	app.Decision:    "A position settled, with its reasoning",
	app.Exploration: "A question worked through",
}

// addNouns builds one command per record type and hangs the verbs that apply
// to it underneath.
func addNouns(root *cobra.Command, a *App) {
	for _, unit := range app.Units {
		noun := &cobra.Command{
			Use:   unit,
			Short: nounHelp[unit],
			// Without NoArgs, Cobra takes a verb this noun does not have as a
			// positional argument, prints help and exits 0 --- so `task close`
			// looks like it worked. The same trap as the root command
			// (root.go), one level down.
			Args:          cobra.NoArgs,
			SilenceUsage:  true,
			SilenceErrors: true,
			// A noun on its own is not an action; show what it can do.
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
		}
		noun.AddCommand(newNewCommand(a, unit))
		noun.AddCommand(newListCommand(a, unit))

		switch unit {
		case app.WorkItem:
			noun.AddCommand(newCloseCommand(a))
			noun.AddCommand(newJournalCommand(a))
		case app.Outcome:
			noun.AddCommand(newVerifyCommand(a))
		}

		root.AddCommand(noun)
	}
}
