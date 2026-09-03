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
# needs them, not in advance.

lkf_version:    0.0.2          # format grammar this bundle is written against
type_namespace: luma/backlog   # records write short type names; this resolves them

# Where the work is. The first three describe how far the planning has gone
# rather than where the work is, so a glance tells you what is an idea and what
# is ready to pick up. An absent status means the first value here.
workflow_status:
  work-item:   [idea, preparing, ready, todo, in_progress, closed]
  task:        [todo, in_progress, closed]

# Statuses grouped into board columns, so a precise vocabulary still renders as
# a legible board.
columns:
  Backlog:     [idea, preparing, ready]
  To Do:       [todo]
  In Progress: [in_progress]
  Closed:      [closed]
`
