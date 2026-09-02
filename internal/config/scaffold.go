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
  deliverable: [idea, preparing, actionable, todo, in_progress, closed]
  task:        [todo, in_progress, closed]

# Statuses grouped into board columns, so a precise vocabulary still renders as
# a legible board.
columns:
  Backlog:     [idea, preparing, actionable]
  To Do:       [todo]
  In Progress: [in_progress]
  Closed:      [closed]
`

// BundleRootFile is the generated bundle root.
//
// It carries lkf_version and type_namespace as regenerated keys so a tool that
// understands the knowledge format, but nothing about this one, can resolve
// short type names without parsing a private configuration file.
const BundleRootFile = `---
type: bundle
title: Backlog
lkf_version: 0.0.2
type_namespace: luma/backlog
---

# Backlog

The bundle root. This file is authored; ` + "`lkf_version`" + ` and ` + "`type_namespace`" + ` are
**regenerated keys** sourced from ` + "`config.yml`" + `, which is where they are edited.

They are mirrored here so a tool that understands the knowledge format, but
knows nothing about this one, can resolve ` + "`type: task`" + ` to its full name without
parsing a private configuration file. Anything reading this programmatically
should prefer ` + "`backlog config`" + `.
`
