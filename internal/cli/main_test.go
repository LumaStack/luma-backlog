package cli

import (
	"os"
	"testing"
)

// TestMain clears the environment the tests run in.
//
// Nothing here shells out yet, so the git lockdown in docs/testing.md is not
// needed and is deliberately not written — it would be a fence around a gate
// nobody has built. What IS needed is this: the tool reads its actor from the
// environment, so a developer with LUMA_BACKLOG_ACTOR set would see different
// behavior from continuous integration, and only sometimes.
//
// When the tool starts running git, the full lockdown belongs here — pointers
// unset FIRST, since an inherited GIT_DIR is resolved before the ceiling is
// ever parsed and voids everything after it.
func TestMain(m *testing.M) {
	for _, key := range []string{
		"LUMA_BACKLOG_ACTOR",
		// Cleared now rather than later: these cost nothing while unused, and
		// each is a way for a developer's machine to disagree with everyone
		// else's exactly once, at the worst moment.
		"GIT_DIR", "GIT_WORK_TREE", "GIT_INDEX_FILE", "GIT_COMMON_DIR",
		"GIT_OBJECT_DIRECTORY", "GIT_ALTERNATE_OBJECT_DIRECTORIES", "GIT_NAMESPACE",
	} {
		os.Unsetenv(key)
	}
	os.Setenv("TZ", "UTC")
	os.Setenv("LC_ALL", "C")
	os.Exit(m.Run())
}
