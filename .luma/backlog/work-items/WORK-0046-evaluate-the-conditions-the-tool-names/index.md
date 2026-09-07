---
type: work-item
key: WORK-0046
title: Evaluate the conditions the tool names
workflow_status: captured
kind: change
stage: draft
description: Ten conditions are specified by name and none is evaluated.
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T01:00:00Z'}
---

# Evaluate the conditions the tool names

## The problem

**`spec.md` §5.2 names a fixed set of conditions and nothing computes them.**
The set is deliberately closed --- *"not a general query language ... a query
language rich enough to express arbitrary conditions is a rules engine"* --- so
this is a finite piece of work rather than an open one.

Four of them are the mechanical half of
`.luma/bundles/local/backlog/policy/when-a-work-item-splits.md`, which without
them is a prose-only rule:

| condition | fires when |
| --- | --- |
| `task.advances-nothing` | a task is attached to no outcome |
| `work-item.unarticulated` | a work item has no outcomes at all |
| `outcome.unmeasured` | an outcome has no `verify_by` |
| `work-item.drifted` | work happened and no outcome was verified or revised |

**`work-item.drifted` has already fired unobserved.** WORK-0031 had nine tasks
closed, no outcome touched, and progress reported from task counts. Nothing
noticed because nothing evaluates it.

## Roughly what it covers

The `check` command, and the conditions themselves. Reported the way a listing
reports skips and duplicates --- **observed, never refused**.

## Open before this can be planned

- **`work-item.drifted` needs a definition of "work happened"** that does not
  require a wave, since waves are not built. Tasks closing since the last
  outcome was touched is the obvious reading and belongs in a decision rather
  than an implementation.
- **Whether conditions are computed on demand or reported by every read.**
  Duplicates and skips are already reported by every listing; these may want the
  same treatment rather than a command nobody runs.
