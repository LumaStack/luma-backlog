package cli

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/lumastack/luma-backlog/internal/backlog"
	"github.com/spf13/cobra"
)

func newShowCommand(app *App) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "show <record>",
		Short: "Read one record",
		Long: "Takes a slug, a path, or an unambiguous prefix.\n\n" +
			"An ambiguous reference lists the candidates rather than guessing.",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			b, cfg, _, err := openBacklog(app)
			if err != nil {
				return err
			}
			defer b.Close()

			it, err := backlog.Resolve(b, args[0])
			if err != nil {
				return coded{ExitNotFound, err}
			}

			out := cmd.OutOrStdout()
			if asJSON {
				rj, err := toRecordJSON(it, cfg.DefaultStatusFor(it.Type()))
				if err != nil {
					return failure("%w", err)
				}
				return writeJSON(out, rj)
			}

			fmt.Fprintf(out, "%s\n%s\n\n", it.Title(), strings.Repeat("─", len([]rune(it.Title()))))
			w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
			fmt.Fprintf(w, "path\t%s\n", it.Path)
			fmt.Fprintf(w, "type\t%s\n", it.Type())
			if k := it.Key(); k != "" {
				fmt.Fprintf(w, "key\t%s\n", k)
			}
			if st := it.Status(cfg.DefaultStatusFor(it.Type())); st != "" {
				fmt.Fprintf(w, "status\t%s\n", st)
			}
			for _, k := range it.Record.Keys() {
				switch k {
				case "type", "title", "workflow_status":
					continue
				}
				if v, ok := it.Record.Get(k); ok {
					fmt.Fprintf(w, "%s\t%s\n", k, v)
				}
			}
			w.Flush()
			if body := strings.TrimSpace(it.Record.Body()); body != "" {
				fmt.Fprintf(out, "\n%s\n", body)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "emit the record as JSON")
	return cmd
}
