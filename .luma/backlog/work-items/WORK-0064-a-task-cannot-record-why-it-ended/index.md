---
type: work-item
key: WORK-0064
title: A task cannot record why it ended
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T00:36:20Z'}
description: a task is defined as an attempt at part of the work, but there is no task close and no disposition vocabulary — a task ends by transitioning it to closed and the record cannot say whether it succeeded, was tried and failed, or was dropped; the history of attempts loses the half that is worth keeping
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T00:36:20Z'}
---

# A task cannot record why it ended

## The problem

**A task is defined as an attempt.** `luma-backlog task --help` says so
outright — *"An attempt at part of the work."* An attempt is a thing that can
fail, and nothing on a task can say that it did.

**There is no `task close`.** `close` is registered on `work-item` only, so a
task ends with `transition <ref> closed` and the record carries no
disposition. Succeeded, tried and abandoned, overtaken by a better approach,
found unnecessary — all of them look identical afterwards.

**This is the exact loss ADR-0007 gave work items four dispositions to
prevent.** Its test — *the enum carries what the record cannot* — applies
unchanged here: nothing elsewhere on a closed task reconstructs whether it
worked.

## Why it went unnoticed

**Tasks have been read as plan items.** `backlog-journal`'s *what stays out*
lists *"remaining work (tasks)"*, which is forward-looking, while the type
definition is backward-looking. **Two documents in force disagree about what a
task is**, and under the planning reading a disposition is pointless — a plan
item is done or not done.

Under the attempt reading it is the most valuable field on the record.

## What it blocks

**`spec.md` §4.6 succession has nothing to fire on.** *"When an attempt does not
succeed and another begins, a new record is created and links back to the one it
follows."* That presupposes the corpus knows an attempt did not succeed.
Nothing does.
[[work-items/WORK-0024-two-attempts-at-the-same-task-cannot-both-exist]] is the
other half of the same gap — it cannot create the successor; this cannot mark
the predecessor.

**And the history of struggle is the thing being lost.** Where an approach was
tried and abandoned is what a later reader most needs and what nothing else
holds — the journal carries reasoning, not per-attempt outcomes.

## Constraints

- **The vocabulary is already designed.** ADR-0007 did the thinking for work
  items — `completed`, `canceled`, `rejected`, `superseded`. Whether a task
  wants the same four, fewer, or different ones is the open question; inventing
  a second unrelated vocabulary is the thing to avoid.
- **Cheap, and it does not wait on waves.**
  [[work-items/WORK-0019-a-ledger-of-attempts-against-an-outcome]] says the
  richer version needs waves, which are specified and unbuilt. A disposition on
  a task is available now and independent of that.

## References

- [[work-items/WORK-0019-a-ledger-of-attempts-against-an-outcome]] — the deeper
  version, blocked on waves.
- [[work-items/WORK-0024-two-attempts-at-the-same-task-cannot-both-exist]] — the
  successor half.
- [[work-items/WORK-0031-reshape-the-command-surface]] — its journal is where
  the absence of `task close` was first noticed and left undecided.
