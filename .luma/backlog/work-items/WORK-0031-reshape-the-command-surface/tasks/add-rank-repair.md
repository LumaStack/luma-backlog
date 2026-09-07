---
type: task
title: Add rank repair
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T19:55:00Z'}
---

# Add rank repair

The third mechanism
[[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]
names for holding `workflow_status` and `rank` together. The first two are
built: one operation writes both, and no adapter can reach past it. Drift is
already **detected** --- a listing reports every record whose prefix disagrees
with the ordinal its status now carries. Nothing repairs it in bulk.

**It needs no history and no heuristic.** The prefix is recomputed from
`workflow_status`, which the record already carries. The position component is
left exactly as it is --- bisection works on the position alone, and rewriting
it would silently reorder a queue somebody arranged by hand.

**Today the message says to re-set the status on each record**, which does the
same thing one at a time. That is fine for the handful a rename produces and
wrong for a corpus.

## What is to be done

- Recompute the ordinal prefix for every drifted record, leaving the position.
- Report what it changed, and what it could not --- a status the ladder no
  longer carries at all has no ordinal to compute, and that record needs a
  person.
- **Multi-record write**, so it carries the guarantees `spec.md` §9.6 states
  for those.

## Verified by

- A hand-edited vocabulary produces drift; repair clears it and the listing
  stops reporting any.
- Positions are unchanged by the repair --- the order within each status is the
  same before and after.
- A record at a status the ladder does not carry is reported rather than
  guessed at.
