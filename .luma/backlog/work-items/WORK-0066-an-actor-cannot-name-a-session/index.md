---
type: work-item
key: WORK-0066
title: An actor cannot name a session
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T02:18:56Z'}
description: LUMA_BACKLOG_ACTOR is agent:<model>/<project>, which names a model and a project rather than a worker — two agent sessions running concurrently in two worktrees write identical actor strings, so the corpus cannot tell parallel agents apart
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T02:18:56Z'}
---

# An actor cannot name a session

## The problem

**`LUMA_BACKLOG_ACTOR="agent:<model>/<project>"` names a kind of worker, not a
worker.** Two agent sessions running at the same time in two worktrees write
**identical** actor strings, so nothing in the corpus can tell them apart.

**The concurrency is already assumed elsewhere.** The `git-worktrees` bundle
exists specifically to run concurrent agents in one repository, and this project
adopted it. So the tool supports a concurrency its identity format cannot
attribute.

**And it is the single-actor case baked into the identity itself** —
[[records/decisions/ADR-0010-a-team-of-people-and-a-team-of-agents-is-the-default-shape]]
makes many actors the default shape, which turns this from a future nicety into
a defect against a decision in force.

## What it blocks

**A work-in-progress limit.** ADR-0010 sets the unit at one human assignee or
one agent session. The session half is uncomputable, so the rule can be written
and never checked — `backlog-move` and `backlog-rundown` both now say so in
place of the check.

**Taking.** ADR-0008 makes a take exclusive and expiring, held by whoever is
working the task. Two sessions with the same actor string cannot hold two
distinct takes, and neither can release without releasing the other's.

**Staleness with no threshold.** The clean answer — a session ends, its take
expires, the work returns to `todo` — needs the corpus to know which session
ended.

## Constraints

- **Provenance is the point of the field** (`CLAUDE.md`). A change here must
  make attribution sharper, never vaguer.
- **It must not become an operating system username.**
  [[work-items/WORK-0035-the-operating-system-username-must-never-be-an-actor]]
  settled that, and a session identity is exactly the kind of thing a tool
  reaches for a machine name to produce.
- **Existing records carry the current form**, so adopting a new one needs a
  backfill or an explicit grandfather clause —
  [[adopting-a-rule-the-corpus-does-not-meet]].

## Open

**What a session identity is, and who mints it.** A random handle per session is
enough to distinguish, and carries nothing a reader can use. Something derived
from the worktree or the branch would be legible and might leak a path or a
machine name, which is the trap WORK-0035 already closed.

**Whether it belongs in the actor at all**, or beside it. The actor answers
*who*; a session answers *which run*. Those may be two fields —
[[work-items/WORK-0030-an-actor-can-act-on-behalf-of-another]] is the other
record already pulling on the shape of this one.

## References

- [[records/decisions/ADR-0010-a-team-of-people-and-a-team-of-agents-is-the-default-shape]]
- [[records/decisions/ADR-0008-taking-a-task-expires-and-owning-a-work-item-does-not]]
- [[work-items/WORK-0035-the-operating-system-username-must-never-be-an-actor]]
- [[work-items/WORK-0030-an-actor-can-act-on-behalf-of-another]]
