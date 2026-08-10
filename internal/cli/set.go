package cli

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"github.com/spf13/cobra"
)

func newSetCommand(app *App) *cobra.Command {
	var (
		ifUnchanged string
		unset       []string
	)

	cmd := &cobra.Command{
		Use:   "set <record> [field=value ...]",
		Short: "Change fields on a record",
		Long: "Changes only the fields named. Everything else — including keys this\n" +
			"tool knows nothing about — is left exactly as it was.\n\n" +
			"  field=value    a string\n" +
			"  field:=value   parsed as YAML, for a list or a map\n\n" +
			"The two are separate because a wikilink looks like a YAML list:\n" +
			"guessing would turn [[deliverables/x]] into a nested sequence.",
		Args:         cobra.MinimumNArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSet(app, cmd, args[0], args[1:], unset, ifUnchanged)
		},
	}
	cmd.Flags().StringVar(&ifUnchanged, "if-unchanged", "",
		"the hash from `show --json`; refuses if the record changed since")
	cmd.Flags().StringArrayVar(&unset, "unset", nil, "remove a field entirely")
	return cmd
}

func runSet(app *App, cmd *cobra.Command, ref string, assignments, unset []string, ifUnchanged string) error {
	b, cfg, _, err := openBacklog(app)
	if err != nil {
		return err
	}
	defer b.Close()

	if len(assignments) == 0 && len(unset) == 0 {
		return usageErr("nothing to change: pass field=value, or --unset field")
	}

	it, err := backlog.Resolve(b, ref)
	if err != nil {
		return coded{ExitNotFound, err}
	}

	// Optimistic concurrency: the caller states what it saw, and a write that
	// would clobber a change it never saw is refused rather than applied
	// (docs/SPEC.md §6.3). Exit 4 means re-read and retry, which is different
	// advice from "something broke" — and it is the distinction a retrying
	// agent depends on.
	if ifUnchanged != "" && it.Hash() != ifUnchanged {
		return coded{ExitConflict, fmt.Errorf(
			"%s changed since you read it — re-read and retry\n  you saw:  %s\n  it is now: %s",
			it.Path, short(ifUnchanged), short(it.Hash()))}
	}

	for _, a := range assignments {
		key, value, ok := strings.Cut(a, "=")
		if !ok || key == "" {
			return usageErr("%q is not field=value", a)
		}
		if raw := strings.HasSuffix(key, ":"); raw {
			key = strings.TrimSuffix(key, ":")
			if err := it.Record.SetRaw(key, value); err != nil {
				return usageErr("%w", err)
			}
			continue
		}
		it.Record.Set(key, value)
	}
	for _, key := range unset {
		it.Record.Remove(key)
	}

	// modified advances on edit, as the format defines it. Written after the
	// caller's own changes so an explicit modified: wins — the tool should not
	// overrule something it was just told.
	if !assigned(assignments, "modified") {
		if err := it.Record.SetRaw("modified",
			"{by: "+app.Env.Actor.String()+", at: "+app.Env.Now()+"}"); err != nil {
			return failure("%w", err)
		}
	}

	out, err := it.Record.Bytes()
	if err != nil {
		return failure("serializing %s: %w", it.Path, err)
	}
	if err := b.WriteFileAtomic(it.Path, out, 0o644); err != nil {
		return failure("%w", err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "updated  %s\n", it.Path)
	_ = cfg
	return nil
}

func assigned(assignments []string, key string) bool {
	for _, a := range assignments {
		if k, _, _ := strings.Cut(a, "="); strings.TrimSuffix(k, ":") == key {
			return true
		}
	}
	return false
}

func short(hash string) string {
	if len(hash) > 12 {
		return hash[:12]
	}
	return hash
}
