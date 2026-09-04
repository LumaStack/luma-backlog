package cli

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"github.com/lumastack/luma-backlog/internal/root"
	"github.com/spf13/cobra"
)

func newJournalCommand(app *App) *cobra.Command {
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
			b, _, projectRoot, err := openBacklog(app)
			if err != nil {
				return err
			}
			defer b.Close()

			slug, err := resolveJournalWorkItem(b, workItem, projectRoot, app.WorkingDir)
			if err != nil {
				return err
			}
			rel := path.Join(backlog.BundleDir, "work-items", slug, "journal.md")

			current := ""
			if data, err := b.ReadFile(rel); err == nil {
				current = string(data)
			}

			if len(args) == 0 {
				if strings.TrimSpace(current) == "" {
					fmt.Fprintf(cmd.OutOrStdout(), "The journal for %s is empty.\n", slug)
					return nil
				}
				fmt.Fprint(cmd.OutOrStdout(), current)
				return nil
			}

			line := strings.TrimSpace(args[0])
			if line == "" {
				return usageErr("nothing to write")
			}
			if err := b.WriteFileAtomic(rel, []byte(
				backlog.AppendLine(current, app.Env.Today(), line)), 0o644); err != nil {
				return failure("%w", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "written  %s\n", rel)
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

// resolveJournalWorkItem decides whose journal to write to.
//
// In order: the flag, then the working directory, then — only when there is
// exactly one work item — that one. The last is a real convenience early on
// and errs safe: the moment there are two, it stops guessing and names them.
//
// The precedence matters more here than for other commands. Capture has to
// cost one invocation, and requiring a flag every time is the friction that
// stops it happening at all.
func resolveJournalWorkItem(b *root.Backlog, flag, projectRoot, workingDir string) (string, error) {
	if flag != "" {
		return flag, nil
	}
	if fromDir := workItemFromWorkingDir(projectRoot, workingDir); fromDir != "" {
		return fromDir, nil
	}

	// Deferred with Resolve, and for the same reason — see load.go.
	items, _, err := backlog.List(b, backlog.Filter{Unit: backlog.WorkItem})
	if err != nil {
		return "", failure("%w", err)
	}
	switch len(items) {
	case 0:
		return "", usageErr("no work items yet — create one first")
	case 1:
		return items[0].Slug(), nil
	}

	slugs := make([]string, 0, len(items))
	for _, it := range items {
		slugs = append(slugs, it.Slug())
	}
	sort.Strings(slugs)
	return "", usageErr("more than one work item — say which with --work-item:\n  %s",
		strings.Join(slugs, "\n  "))
}
