---
type: decision
title: A team of people and a team of agents is the default shape
decided: 2026-09-08
stage: provisional
reopen_trigger: the reduced single-actor path acquires ceremony — a field nobody can fill, a step with one possible answer, a prompt that always resolves the same way — which is the failure this record exists to prevent and cannot be argued away
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T02:10:00Z'}
---

# ADR-0010: A team of people and a team of agents is the default shape

## Summary

**Everything is designed for many actors working at once** — a team of people
with a team of agents — and the single-actor case is the **exception that
reduces out of it**. Not the base case that gets extended later.

**Reduction makes the multi-actor machinery invisible, never absent.** One
record shape at every team size, so growing a team is not a migration.

## Problem

**The rules this project writes keep encoding one maintainer as though it were
universal.** Three shipped on 2026-09-07 alone, all in one procedure, all caught
by a reader rather than by anything mechanical:

| written as | true only because |
| --- | --- |
| *"a work item may be created, worked and closed inside one evening"* | one person, one schedule |
| *"sessions, not weeks"* | one worker whose memory ends with a session |
| *"several things `in_progress` at once is itself a finding"* | one worker, so *several* and *several by one person* are the same |

**None is wrong here and all three are wrong anywhere else.** A ten-person team
with ten items in flight is healthy; the finding is one *assignee* holding
several. The single-actor reading is invisible from inside the single-actor
case, which is why writing it down keeps happening.

**And the assumption reaches the identity format.**
`LUMA_BACKLOG_ACTOR="agent:<model>/<project>"` names a model and a project, not
a worker — so two agent sessions running concurrently in two worktrees write
**identical** actor strings. The `git-worktrees` bundle exists specifically to
run concurrent agents in one repository, so the corpus already assumes a
concurrency it cannot attribute.

## Decision

**Design for many, reduce to one.**

**A single-person team is still a team**, because the agents are members of it.
The interesting concurrency is present at every size; only the number of humans
changes.

**Reduction is invisibility, not absence.** This is the whole content of the
word *gracefully*:

- **Invisible** — the field exists, and is filled with the only candidate. One
  record shape. Growing a team means more actors appear.
- **Absent** — the field is missing, so the record shape differs by team size,
  and growing a team becomes a migration.

The second is what happens when nobody says which was meant.

**The unit of concurrency is a human assignee or an agent session**, not an
agent identity — one agent can hold many sessions at once. `ADR-0008` already
separates the two lifetimes this needs: a session **takes**, and the take
expires; a human **owns**, and ownership does not.

## Consequences

**A work-in-progress limit is per-taker, not per-owner.** A human may own five
work items and have one taken. That is not a contradiction; it is the two
relationships ADR-0008 defines.

**Staleness stops needing a threshold.** A session ends, its take expires,
nobody is on it, and the work returns to `todo`. No elapsed time, no judgment —
which is what four attempts to write a duration into `backlog-move` were
reaching for and never found.

**The actor format is the exception hardcoded.** It cannot identify a session,
so nothing above is computable until it can. That is a defect under this record
rather than a future nicety.

**Half of this is testable here and half is not.** The multi-actor paths will go
unexercised — one maintainer, and no second human. **The reduced path is
exercised constantly**, which is the half that usually fails: if the single-actor
case ever feels like paperwork, reduction was done wrong. That is the reopen
trigger.

## Alternatives

**Design for one actor and extend later.** *Deferred, and it is the default
outcome of doing nothing.* It matches this project's bootstrap order — lead with
the thing, backfill the mechanism — and it is cheaper today. It was set aside
because the extension is not additive: single-actor assumptions reach the
identity format, the record shape and the prose, and the three fossils above
show they are written without anybody noticing. *Re-open if the multi-actor
machinery proves to cost the solo case more than the fossils cost.*

**Two record shapes, one per team size.** *Deferred.* Cheaper for small teams
and it makes growth a migration — and this project already has
`WORK-0037` and `WORK-0022` open on how expensive corpus migration is.
*Re-open if invisibility turns out to be unachievable for some field.*
