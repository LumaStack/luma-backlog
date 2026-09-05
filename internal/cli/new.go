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
	var project bool

	cmd := &cobra.Command{
		Use:   "new <" + strings.Join(backlog.Units, "|") + "> <title>",
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
			return runNew(app, cmd, args[0], args[1], workItem, kind, project, cmd.Flags().Changed("work-item"))
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

func runNew(app *App, cmd *cobra.Command, unit, title, workItem, kind string, project, workItemGiven bool) error {
	b, cfg, projectRoot, err := openBacklog(app)
	if err != nil {
		return err
	}
	defer b.Close()

	if workItem == "" {
		workItem = workItemFromWorkingDir(projectRoot, app.WorkingDir)
	}

	// A decision's level is stated, never inferred. The working directory can
	// say WHICH work item, and it must not be allowed to say WHICH LEVEL: a
	// person at a terminal is usually standing where they are working, and an
	// agent runs from the repository root whatever it is doing, so the context
	// it would infer from is a constant. Every decision in this repository is
	// project-level because every one of them was created from the root.
	//
	// The cost of being wrong is asymmetric. A decision filed at the wrong
	// level is not visibly broken; it is simply somewhere nobody looks.
	if unit == backlog.Decision {
		if project && workItemGiven {
			return usageErr("--project and --work-item say different levels: pass one")
		}
		if !project && !workItemGiven {
			return usageErr("a decision needs its level stated:\n" +
				"  --work-item <slug>  it belongs to that work item, and is a point-in-time record\n" +
				"  --project           it is a standing rule for the project\n" +
				"Most decisions are work item decisions. Promotion is a separate act.")
		}
		if project {
			// Stated as project-level, so the working directory does not get
			// to attach it to whatever the caller happened to be standing in.
			workItem = ""
		}
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

	// Avoided, not denied. Blank is honest when nobody has looked yet — an
	// issue synced from elsewhere, a note taken in a hurry — and it is the
	// wrong answer the rest of the time, because whoever writes a record
	// usually knows what they are holding. Left free, blank becomes the path
	// of least resistance and the field stops meaning anything.
	//
	// So: say so and continue (docs/spec.md §5.0). Nothing is refused.
	if res.Created && unit == backlog.WorkItem && kind == "" {
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
