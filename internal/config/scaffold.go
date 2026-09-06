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

# Where the work is. A selection gate sits between captured and unprepared, and
# another between prepared and todo — everything between them is preparation.
# An absent status means the first value here.
# The number is the status's ordinal, and it prefixes every rank at that
# status, so sorting the rank field alone gives board order (ADR-0005). They
# are spaced so a status can be inserted between two others without renumbering
# --- renumbering rewrites the rank of every record at every status after it.
#
# A status shared by two units carries the same ordinal in both, so a rank
# means the same thing whichever unit it is on.
workflow_status:
  work-item:
    captured:    10
    unprepared:  20
    preparing:   30
    prepared:    40
    todo:        50
    in_progress: 60
    closed:      70
  task:
    todo:        50
    in_progress: 60
    closed:      70

# Statuses grouped into board columns, so a precise vocabulary still renders as
# a legible board.
columns:
  Captured:    [captured]
  Preparing:   [unprepared, preparing, prepared]
  To Do:       [todo]
  In Progress: [in_progress]
  Closed:      [closed]
`
