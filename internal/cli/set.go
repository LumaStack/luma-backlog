package cli

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/app"
	"github.com/spf13/cobra"
)

func newSetCommand(a *App) *cobra.Command {
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
			"guessing would turn [[work-items/x]] into a nested sequence.",
		Args:         cobra.MinimumNArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSet(a, cmd, args[0], args[1:], unset, ifUnchanged)
		},
	}
	cmd.Flags().StringVar(&ifUnchanged, "if-unchanged", "",
		"the hash from `show --json`; refuses if the record changed since")
	cmd.Flags().StringArrayVar(&unset, "unset", nil, "remove a field entirely")
	return cmd
}

func runSet(a *App, cmd *cobra.Command, ref string, assignments, unset []string, ifUnchanged string) error {
	s, err := open(a)
	if err != nil {
		return err
	}
	defer s.Close()

	parsed := make([]app.Assignment, 0, len(assignments))
	for _, raw := range assignments {
		key, value, ok := strings.Cut(raw, "=")
		if !ok || key == "" {
			return usageErr("%q is not field=value", raw)
		}
		// field:=value is parsed as YAML; field=value is a string. The two are
		// separate because a wikilink looks like a YAML list.
		isRaw := strings.HasSuffix(key, ":")
		parsed = append(parsed, app.Assignment{
			Field: strings.TrimSuffix(key, ":"),
			Value: value,
			Raw:   isRaw,
		})
	}

	res, err := s.Set(app.SetRequest{
		Ref:         ref,
		Assignments: parsed,
		Unset:       unset,
		IfUnchanged: ifUnchanged,
	})
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "updated  %s\n", res.Path)
	return nil
}
