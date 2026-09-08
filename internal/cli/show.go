package cli

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func newShowCommand(a *App) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "show <record>",
		Short: "Read one record",
		Long: "Takes a slug, a path, or an unambiguous prefix.\n\n" +
			"An ambiguous reference lists the candidates rather than guessing.",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := open(a)
			if err != nil {
				return err
			}
			defer s.Close()

			rec, err := s.Get(args[0])
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			if asJSON {
				return writeJSON(out, toRecordJSON(rec))
			}

			fmt.Fprintf(out, "%s\n%s\n\n", rec.Title, strings.Repeat("─", len([]rune(rec.Title))))
			w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
			fmt.Fprintf(w, "path\t%s\n", rec.Path)
			fmt.Fprintf(w, "type\t%s\n", rec.Type)
			if rec.Key != "" {
				fmt.Fprintf(w, "key\t%s\n", rec.Key)
				fmt.Fprintf(w, "name\t%s\n", rec.Name)
			}
			if rec.Status != "" {
				fmt.Fprintf(w, "status\t%s\n", rec.Status)
			}
			// Computed, and above the frontmatter because it is the one line
			// here that appears nowhere in the file. A work item with no
			// outcomes says so rather than showing 0 of 0, which reads as
			// progress measured against nothing.
			if c := rec.Completion; c != nil {
				switch {
				case c.Live == 0:
					fmt.Fprintf(w, "outcomes\tnone\n")
				default:
					line := fmt.Sprintf("%d of %d proven", c.Proven, c.Live)
					if c.Retired > 0 {
						line += fmt.Sprintf(", %d retired", c.Retired)
					}
					if c.Skipped > 0 {
						line += fmt.Sprintf(", %d unreadable", c.Skipped)
					}
					fmt.Fprintf(w, "outcomes\t%s\n", line)
				}
			}
			for _, k := range rec.Order {
				switch k {
				case "type", "title", "workflow_status":
					continue
				}
				if v, ok := rec.Raw[k]; ok {
					fmt.Fprintf(w, "%s\t%s\n", k, v)
				}
			}
			w.Flush()
			if body := strings.TrimSpace(rec.Body); body != "" {
				fmt.Fprintf(out, "\n%s\n", body)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "emit the record as JSON")
	return cmd
}
