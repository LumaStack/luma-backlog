package cli

import (
	"fmt"
	"strings"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newVerifyCommand(app *App) *cobra.Command {
	var evidence string

	cmd := &cobra.Command{
		Use:   "verify <outcome>",
		Short: "Record that an outcome has been confirmed",
		Long: "Adds a confirmation event, and the evidence it rests on.\n\n" +
			"Verification accumulates: several actors confirming the same outcome is\n" +
			"the normal case, and a human entry raises the derived trust tier with no\n" +
			"special handling.",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			b, _, _, err := openBacklog(app)
			if err != nil {
				return err
			}
			defer b.Close()

			it, err := backlog.Resolve(b, args[0])
			if err != nil {
				return coded{ExitNotFound, err}
			}
			if it.Type() != backlog.Outcome {
				return usageErr("%s is a %s — only an outcome is verified", it.Slug(), it.Type())
			}

			at, by := app.Env.Now(), app.Env.Actor.String()

			if err := appendToList(it.Record, "verified",
				map[string]string{"by": by, "at": at}); err != nil {
				return failure("%w", err)
			}

			// Evidence sits BESIDE verified rather than inside it, because
			// verified is a core format field and inheritance is add-only, so
			// the key cannot be added there (format-requests.md §3). The two
			// correlate on by and at rather than by position, since parallel
			// lists matched by index break the first time one is hand-edited.
			if strings.TrimSpace(evidence) != "" {
				if err := appendToList(it.Record, "evidence",
					map[string]string{"by": by, "at": at, "what": evidence}); err != nil {
					return failure("%w", err)
				}
			}

			out, err := it.Record.Bytes()
			if err != nil {
				return failure("%w", err)
			}
			if err := b.WriteFileAtomic(it.Path, out, 0o644); err != nil {
				return failure("%w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "verified  %s\n", it.Path)
			if strings.TrimSpace(evidence) == "" {
				fmt.Fprintln(cmd.OutOrStdout(),
					"\nNo evidence recorded. An unbacked confirmation is the claim this\n"+
						"design distrusts most — pass --evidence next time.")
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&evidence, "evidence", "e", "", "what the confirmation rests on")
	return cmd
}

// appendToList adds an entry to a list-valued field, creating the list when
// absent and tolerating a single bare entry written by hand.
func appendToList(r interface {
	Node(string) *yaml.Node
	SetRaw(string, string) error
}, key string, entry map[string]string) error {
	var existing []map[string]any
	if node := r.Node(key); node != nil {
		if err := node.Decode(&existing); err != nil {
			// A single bare mapping is legal and means a one-element list,
			// following the format's own handling of verified.
			var one map[string]any
			if err2 := node.Decode(&one); err2 != nil {
				return fmt.Errorf("%s is neither a list nor an entry: %w", key, err)
			}
			existing = []map[string]any{one}
		}
	}
	next := map[string]any{}
	for k, v := range entry {
		next[k] = v
	}
	existing = append(existing, next)

	encoded, err := yaml.Marshal(existing)
	if err != nil {
		return err
	}
	return r.SetRaw(key, string(encoded))
}
