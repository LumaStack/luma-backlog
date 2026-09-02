package cli

import (
	"errors"
	"fmt"

	"github.com/lumastack/luma-backlog/internal/config"
	"github.com/lumastack/luma-backlog/internal/root"
	"github.com/spf13/cobra"
)

func newInitCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create a backlog in this repository",
		Long: "Creates .luma/ with the backlog bundle and records tier, and a\n" +
			"configuration file written out in full.\n\n" +
			"Safe to run again: nothing existing is overwritten, and anything\n" +
			"missing is created.",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runInit(app, cmd)
		},
	}
}

func runInit(app *App, cmd *cobra.Command) error {
	projectRoot, err := root.Discover(app.WorkingDir, app.Ceiling)
	if err != nil {
		if errors.Is(err, root.ErrNotFound) {
			return usageErr("no git repository here or above %s.\n"+
				"A backlog belongs to a repository — run `git init` first, or move somewhere inside one.",
				app.WorkingDir)
		}
		return failure("finding the project root: %w", err)
	}

	b, err := root.Create(projectRoot)
	if err != nil {
		return failure("%w", err)
	}
	defer b.Close()

	out := cmd.OutOrStdout()

	// Nothing existing is overwritten. Running init twice is an ordinary
	// thing to do — often to pick up a file added by a later version — and a
	// command that clobbered a team's configuration for it would be a trap.
	created, err := writeIfAbsent(b, config.FileName, config.DefaultFile)
	if err != nil {
		return failure("%w", err)
	}
	report(out, created, config.FileName)

	created, err = writeIfAbsent(b, "backlog/index.md", config.BundleRootFile)
	if err != nil {
		return failure("%w", err)
	}
	report(out, created, "backlog/index.md")

	for _, dir := range []string{"backlog/deliverables", "backlog/_types", "records/decisions"} {
		if err := b.MkdirAll(dir); err != nil {
			return failure("creating %s: %w", dir, err)
		}
	}

	fmt.Fprintf(out, "\nBacklog ready at %s\n", b.Path())
	return nil
}

func writeIfAbsent(b *root.Backlog, name, contents string) (bool, error) {
	if b.Exists(name) {
		return false, nil
	}
	if err := b.WriteFileAtomic(name, []byte(contents), 0o644); err != nil {
		return false, err
	}
	return true, nil
}

func report(out interface{ Write([]byte) (int, error) }, created bool, name string) {
	if created {
		fmt.Fprintf(out, "created  %s\n", name)
	} else {
		fmt.Fprintf(out, "exists   %s\n", name)
	}
}
