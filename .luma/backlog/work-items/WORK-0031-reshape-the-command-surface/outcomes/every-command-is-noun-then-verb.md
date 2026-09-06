---
type: outcome
title: Every command is noun then verb
desired_state: "The command tree matches `spec.md` §9.1 --- noun then verb, with verb-only where no noun applies. No command takes a record type as a positional argument."
verify_by:
  - "Confirm `work-item list`, `outcome verify`, `task new` resolve, and that `list work-item` no longer does."
  - "Confirm the verb-only commands are the ones ADR-0006 names: `init`, `board`, `contract`, `config`, `check`, `log`, `serve`."
  - "Confirm each universal verb has one implementation parameterized by noun, not one per noun-verb pair."
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:45:00Z'}
---

# Every command is noun then verb

ADR-0006 confirmed §9.1 and made **the code wrong, not the specification.**
This is where the code catches up.

The noun is already data below the adapter --- `app.CreateRequest.Unit` is a
string, and the filter takes one too. So the noun moves from a positional
argument into a constructor parameter, and `internal/app`, `internal/corpus`
and `internal/record` are untouched. **If this change reaches below
`internal/cli`, something has gone wrong.**
