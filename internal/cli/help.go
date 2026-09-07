package cli

import (
	"strings"

	"github.com/spf13/cobra"
)

// Help follows gh where CLIG is silent about layout (ADR-0006): uppercase
// section headers without colons, commands grouped rather than listed flat,
// examples, and a place to read more.
//
// The template is written out rather than patched into Cobra's default,
// because every section here differs from it.
//
// `IsAvailableCommand` is false for the help command --- Cobra excludes it
// deliberately --- so every listing has to say `or ... (eq .Name "help")` to
// show it. Cobra's own template does; dropping the clause silently lost `help`
// from the listing.
const (
	groupCore    = "core"
	groupRecords = "records"
	groupExtra   = "extra"
)

// addGroups declares the sections the root command's listing is divided into.
//
// Flat, the listing interleaved `completion` and `help` between commands that
// do the work, as though they were peers. The division is by what a reader is
// looking for: a verb to run, a kind of record to work on, or the machinery.
func addGroups(root *cobra.Command) {
	root.AddGroup(
		&cobra.Group{ID: groupCore, Title: "CORE COMMANDS"},
		&cobra.Group{ID: groupRecords, Title: "RECORD COMMANDS"},
		&cobra.Group{ID: groupExtra, Title: "ADDITIONAL COMMANDS"},
	)
	root.SetHelpCommandGroupID(groupExtra)
	root.SetCompletionCommandGroupID(groupExtra)
}

const rootExample = `  $ luma-backlog init
  $ luma-backlog work-item new "Payments v2" --kind change
  $ luma-backlog outcome new "The retry queue drains" -w payments-v2
  $ luma-backlog task new "Add the retry queue" -w payments-v2
  $ luma-backlog list --tree`

const learnMore = `  Use ` + "`luma-backlog <command> --help`" + ` for a command's own options.
  Read the design at https://github.com/LumaStack/luma-backlog`

// usageTemplateFor renders one command's usage. Sections a command does not
// have are omitted rather than printed empty.
const usageTemplate = `USAGE{{if .HasAvailableSubCommands}}
  {{.CommandPath}} <command> <subcommand> [flags]{{else}}
  {{.UseLine}}{{end}}
{{if .HasAvailableSubCommands}}{{range $g := .Groups}}
{{$g.Title}}{{range $c := $.Commands}}{{if (and (eq $c.GroupID $g.ID) (or $c.IsAvailableCommand (eq $c.Name "help")))}}
  {{rpad $c.Name $c.NamePadding }} {{$c.Short}}{{end}}{{end}}
{{end}}{{if not .AllChildCommandsHaveGroup}}
COMMANDS{{range $c := .Commands}}{{if (and (eq $c.GroupID "") (or $c.IsAvailableCommand (eq $c.Name "help")))}}
  {{rpad $c.Name $c.NamePadding }} {{$c.Short}}{{end}}{{end}}
{{end}}{{end}}{{if .HasAvailableLocalFlags}}
FLAGS
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}{{if .HasAvailableInheritedFlags}}
INHERITED FLAGS
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}{{if .HasExample}}
EXAMPLES
{{.Example}}
{{end}}{{if .HasAvailableSubCommands}}
LEARN MORE
` + learnMore + `
{{end}}`

// applyHelp installs the layout on a command and everything under it. Cobra
// resolves a template through the parent, so setting it on the root is enough
// --- but the root is the only command with groups, and the template has to
// stay correct for a leaf, which is what the else branches are for.
func applyHelp(root *cobra.Command) {
	root.SetUsageTemplate(strings.TrimLeft(usageTemplate, "\n"))
	root.Example = rootExample

	// Cobra adds `help` lazily, during Execute --- after this template has
	// rendered, so it was missing from the listing it belongs in. Asking for
	// it here puts it where a reader looks. It has to come after the
	// subcommands, because Cobra declines to add it to a command that has none
	// yet, which is why calling it earlier did nothing.
	root.InitDefaultHelpCmd()
}
