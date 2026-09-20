package config

// DefaultFile is what `init` writes.
//
// Written out in full rather than left implicit in the binary, so a team's
// first encounter with a default is a line they can read and change, not
// behavior they have to discover and then find a way to override
// (docs/spec.md §8.3). The comments are the point, which is why this is a
// document rather than a marshalled struct.
//
// TestDefaultFileMatchesDefaults keeps it honest against Default().
const DefaultFile = `# luma-backlog configuration — see docs/spec.md §8
# The source of truth. Deliberately minimal; keys are added when something
# needs them, not in advance — and only keys the binary actually reads.

# The prefix of every NEW work item key (WORK-0042). Optional; WORK when
# absent. Changing it never renames existing records — the prefix is written
# into each record at creation. Jira Cloud rules: an uppercase letter, then
# uppercase letters or digits, two to ten characters.
# work_item_key: WORK
`
