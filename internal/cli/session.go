package cli

import (
	"fmt"
	"io"

	"github.com/lumastack/luma-backlog/internal/app"
)

// open starts a session from the ambient facts this command was given.
//
// Every command that touches records goes through here, and none of them
// reaches for the environment itself — the layer takes actor and root as
// values (.luma/records/decisions/ADR-0004).
func open(a *App) (*app.Session, error) {
	return app.Open(a.Env, a.WorkingDir, a.Ceiling)
}

// observe renders what the layer noticed while working.
//
// To stderr, always: stdout stays a clean listing and --json stays parseable,
// so a caller piping the output is unaffected while still being told.
//
// Writes nothing when there is nothing to say. A warning that appears on
// ordinary runs is one people learn to scroll past, and these have to survive
// being ignored for months before they matter once.
func observe(w io.Writer, o app.Observations) {
	for _, s := range o.Skipped {
		fmt.Fprintf(w, "luma-backlog: skipped %s: %v\n", s.Path, s.Err)
	}
	for _, d := range o.Duplicates {
		fmt.Fprintf(w, "luma-backlog: %s is held by %d records:\n", d.Key, len(d.Paths))
		for _, p := range d.Paths {
			fmt.Fprintf(w, "  %s\n", p)
		}
		fmt.Fprintf(w, "A key is meant to name one record, and every citation of this one is ambiguous.\n")
	}
}
