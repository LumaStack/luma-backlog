package cli

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newJournalCommand(a *App) *cobra.Command {
	var workItem string

	cmd := &cobra.Command{
		Use:   "journal [text]",
		Short: "Write to a work item's memory, or read it",
		Long: "With text, appends one line to today's entry, opening today's entry if\n" +
			"there is not one. With nothing, shows the journal.\n\n" +
			"No file to open, no heading to write, no decision about where it goes:\n" +
			"friction at the moment of writing is what loses the learning.",
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		Example: "  luma-backlog journal \"the ceiling must be symlink-resolved\"\n" +
			"  luma-backlog journal -- \"--use-hold pins the source snapshot\"",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			line := ""
			if len(args) == 1 {
				line = args[0]
			}
			res, err := s.Journal(app.JournalRequest{WorkItem: workItem, Line: line})
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			if res.Written {
				fmt.Fprintf(out, "written  %s\n", res.Path)
				return nil
			}
			if strings.TrimSpace(res.Content) == "" {
				fmt.Fprintf(out, "The journal for %s is empty.\n", res.Slug)
				return nil
			}
			fmt.Fprint(out, res.Content)
			return nil
		},
	}
	cmd.Flags().StringVarP(&workItem, "work-item", "w", "", "whose journal (default: derived)")

	// Journal text very often begins with a flag name — "--use-hold pins the
	// source snapshot" is exactly the sort of thing worth capturing — and the
	// parser sees a flag. The escape is standard, but nobody thinks of it
	// while being told their note is an unknown flag, so the error says so.
	cmd.SetFlagErrorFunc(func(c *cobra.Command, err error) error {
		if strings.HasPrefix(err.Error(), "unknown flag") {
			return usageErr("%w\n\nIf that was the text and not a flag, put -- in front of it:\n"+
				"  luma-backlog journal -- \"--your text here\"", err)
		}
		return err
	})
	return cmd
}
