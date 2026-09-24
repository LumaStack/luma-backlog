// Package migrate rewrites a corpus in ways ordinary commands must not.
//
// **It exists as a package so a guard can name it.** A key migration needs a
// handle reaching past `.luma/` — a work item's name appears in prose, in
// documentation and in source comments, and stopping at the backlog directory
// would leave all of those pointing at a directory that no longer exists. That
// capability is fine here and wrong everywhere else, and
// `internal/guards` fails the build if anything outside this package asks for
// it.
//
// **Which is why migrations did not need their own binary.** The argument for
// splitting them out was never subject matter; it was that they need something
// nothing else should have. A guard buys that separation for the cost of one
// test and leaves the tool one thing to install.
package migrate

import (
	"github.com/lumastack/luma-backlog/internal/root"
)

// openRepository returns the wide handle. The one call in the codebase that
// the guard permits.
func openRepository(projectRoot string) (*root.Project, error) {
	return root.OpenProject(projectRoot)
}
