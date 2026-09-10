---
type: work-item
key: WORK-0064
title: A task cannot record why it ended
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T00:36:20Z'}
description: 'a task is defined as an attempt at part of the work, but there is no task close and no disposition vocabulary — a task ends by transitioning it to closed and the record cannot say whether it succeeded, was tried and failed, or was dropped; the history of attempts loses the half that is worth keeping. it now blocks a warning close needs: warn on a completed close if any task was not successful, which has to tell failed from cancelled or found-unnecessary and so cannot be a boolean.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:24:27Z'}
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

## What it has to enable

**Warn on a completed close if any task was not successful.** Added
2026-09-09, when `close … completed` began refusing while any task is open and
this was the half that could not be built.

**It is a warning and not a bar**, deliberately. Attempting a task several times
is ordinary and some attempts fail — a person needs to see that and decide
whether it points at a problem, rather than be stopped by it.

**Which sets the shape of the disposition.** A boolean will not do: the warning
has to tell a task that **failed** from one **cancelled** or **found
unnecessary**, because only the first is worth looking at. That is the same
argument ADR-0007 made for the work item's four dispositions, and the same test
— *the enum carries what the record cannot*.

**`backlog-move` 0.33.0 carries the row already**, marked unbuilt against this
record, so the promise is visible rather than quietly missing.

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
