---
type: task
title: Build the noun-verb command tree
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:49:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T20:02:32Z'}
---

# Build the noun-verb command tree

**This one goes first.** Every other task in this work item lands inside the
tree it creates.

The noun is already data below the adapter: `app.CreateRequest.Unit` is a
string and `app.Filter` takes one. So the change is confined to
`internal/cli` --- the noun moves from a positional argument into a
constructor parameter:

```go
func newListCommand(a *App, unit string) *cobra.Command
```

registered once per noun. **One implementation per verb, parameterized** ---
not one per noun-verb pair. Copying the body five times is how `outcome list`
and `task list` drift apart.

Nouns: `work-item`, `outcome`, `task`, `decision`, `exploration`. Verb-only,
per ADR-0006: `init`, `board`, `contract`, `config`, `check`, `log`, `serve` ---
only those that exist today get built.

**Verified by:** `work-item list` resolves and `list work-item` does not;
`internal/app`, `internal/corpus` and `internal/record` are untouched by the
diff; the containment test in `internal/guards` still passes.
