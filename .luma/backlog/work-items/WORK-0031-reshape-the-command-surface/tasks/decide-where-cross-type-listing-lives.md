---
type: task
title: Decide where cross-type listing lives
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T17:05:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T16:43:48Z'}
---

# Decide where cross-type listing lives

**Blocks finishing [[work-items/WORK-0031-reshape-the-command-surface/tasks/build-the-noun-verb-command-tree]].**
Until it is settled, `list` keeps its current top-level shape and the noun
commands deliberately do **not** carry a `list` verb --- two ways to list one
type would be worse than one inconsistency.

## The gap

`spec.md` §9.2 puts `list` under the verbs universal to every noun.
ADR-0006 enumerates the verb-only commands --- `init`, `board`, `contract`,
`config`, `check`, `log`, `serve` --- and **`list` is not among them.**

Between them, nothing lists across record types.

**That is the form actually in use.** Of 21 `list` call sites in the test
suite, **18 are bare `list`** with no record type. Only three name one. The
most-used shape of the most-used read command has no home in the reshaped
surface.

Nothing suggests that was intended. It reads as cross-type listing not having
been considered when the verb-only set was enumerated.

## The options

**Add `list` to the verb-only set.** `backlog list` for everything,
`backlog work-item list` for one type. They are not one command wearing two
hats --- the cross-type form has no noun to attach to, the way `git log` and
`git log <path>` differ. Needs ADR-0006 amended, since that set is enumerated.

**Drop cross-type listing.** Cleanest surface, and a real capability lost until
the board arrives. Rewrites the 18 call sites.

**Something else** --- a filter on a top-level command, or leaving it to the
board and saying so.

## Verified by

The decision is recorded rather than assumed, `list` appears in exactly one
place per meaning, and no record type is reachable as a positional argument
to a verb.
