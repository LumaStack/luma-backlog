package cli

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"github.com/lumastack/luma-backlog/internal/config"
	"github.com/lumastack/luma-backlog/internal/root"
	"github.com/spf13/cobra"
)

func newNewCommand(app *App) *cobra.Command {
	var workItem, kind string

	cmd := &cobra.Command{
		Use:   "new <" + strings.Join(backlog.Units, "|") + "> <title>",
		Short: "Create a record",
		Long: "Creates a record from a title. Everything else is derived: the file name\n" +
			"from the title, the work item from where you are, the timestamp and\n" +
			"actor from the environment.\n\n" +
			"Running it twice with the same title leaves the first one alone.\n\n" +
			"--kind classifies a work item — bug, request, idea. Leave it off for\n" +
			"ordinary work, which is most of it. A kind says what has to happen\n" +
			"before the record can be judged: fix it, answer them, or develop it.",
		Args:         cobra.ExactArgs(2),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNew(app, cmd, args[0], args[1], workItem, kind)
		},
	}
	cmd.Flags().StringVarP(&workItem, "work-item", "w", "",
		"the work item this belongs to (default: derived from the working directory)")
	cmd.Flags().StringVarP(&kind, "kind", "k", "",
		"what sort of work item this is (default: none, meaning ordinary work)")
	return cmd
}

func runNew(app *App, cmd *cobra.Command, unit, title, workItem, kind string) error {
	b, cfg, projectRoot, err := openBacklog(app)
	if err != nil {
		return err
	}
	defer b.Close()

	if workItem == "" {
		workItem = workItemFromWorkingDir(projectRoot, app.WorkingDir)
	}

	if kind != "" && unit != backlog.WorkItem {
		return usageErr("--kind classifies a work item; %s does not take one", unit)
	}

	res, err := backlog.Create(b, cfg, app.Env, backlog.Spec{
		Unit:     unit,
		Title:    title,
		WorkItem: workItem,
		Kind:     kind,
	})
	if err != nil {
		return usageErr("%w", err)
	}

	out := cmd.OutOrStdout()
	if res.Created {
		fmt.Fprintf(out, "created  %s\n", res.Path)
	} else {
		fmt.Fprintf(out, "exists   %s\n", res.Path)
	}
	return nil
}

// openBacklog resolves the project root and opens the backlog with its
// configuration. Shared by every command that touches records.
func openBacklog(app *App) (*root.Backlog, config.Config, string, error) {
	projectRoot, err := root.Discover(app.WorkingDir, app.Ceiling)
	if err != nil {
		if errors.Is(err, root.ErrNotFound) {
			return nil, config.Config{}, "", usageErr(
				"no git repository here or above %s", app.WorkingDir)
		}
		return nil, config.Config{}, "", failure("finding the project root: %w", err)
	}

	b, err := root.Open(projectRoot)
	if err != nil {
		return nil, config.Config{}, "", usageErr(
			"no backlog in %s — run `luma-backlog init` first", projectRoot)
	}

	cfg := config.Default()
	if data, err := b.ReadFile(config.FileName); err == nil {
		parsed, perr := config.Parse(data)
		if perr != nil {
			b.Close()
			// Strict where records are permissive: a misspelt configuration
			// key is a silent behavior change (docs/spec.md §8.7).
			return nil, config.Config{}, "", failure("%w", perr)
		}
		cfg = parsed
	}
	return b, cfg, projectRoot, nil
}

// workItemFromWorkingDir reads the work item from where the command was
// run, so someone working inside one does not have to name it.
func workItemFromWorkingDir(projectRoot, workingDir string) string {
	rel, err := filepath.Rel(projectRoot, workingDir)
	if err != nil {
		return ""
	}
	return backlog.WorkItemFromPath(filepath.ToSlash(rel))
}
