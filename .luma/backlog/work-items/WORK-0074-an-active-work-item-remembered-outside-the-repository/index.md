---
type: work-item
key: WORK-0074
title: An active work item, remembered outside the repository
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T19:36:09Z'}
description: what somebody is working on right now, so commands stop needing -w on every call. not in git — it is per actor rather than per corpus, and two agents in worktrees would clobber one pointer. per machine to begin with, per user eventually.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:15:12Z'}
rank: 010.0140.000
---

# An active work item, remembered outside the repository

## The problem

**Every command that belongs to a work item needs telling which one.** `-w` on
every journal line, every outcome, every task. In one session on 2026-09-07/08
that flag was typed several hundred times against the same work item.

**The mechanism that should prevent it does not fire for agents.** Context is
derived from the working directory — *the work item from where you are* — and an
agent runs from the repository root whatever it is doing. So the derivation
returns nothing precisely when the caller types the most.

**And `in_progress` cannot answer it.** That is a fact about the record, shared
by everyone. *What am I working on* is a fact about one actor at one moment, and
this corpus has had two work items `in_progress` at once for most of a day.

## Why it is not in git

**It is per actor, not per corpus.** Two agents in two worktrees are working two
different things against one backlog; a committed pointer gives them one, and
whoever writes last wins.

**And it is not a fact about the project.** Nothing about the corpus changes
when somebody starts and stops. Committing it would put a personal cursor in
everybody's history and make every `git status` dirty.

**`.luma/` cannot hold it.** That directory is committed-only, which is the
invariant that makes it trustworthy — see `luma-layout`.

## Scope

**`pick` and nothing else.** One active work item, replaced by picking another.
No queue, no stack, no second verb — an interruption means picking the other
thing and picking back, which is two commands rather than a structure nobody can
see. *Queues are deferred rather than rejected; re-open when returning turns out
to cost more than re-picking does.*

**`pick` does not move the work item.** Picking a `captured` record to journal a
thought about it is ordinary; dragging it across two gates is not. The pick says
where somebody is, the status says what the work is.

**Per machine first, per user eventually.** On one machine with one person they
are the same thing, and true per-user across machines needs synchronisation
nobody has asked for. Starting per machine costs nothing later: a per-user store
would read the same key.

**Keyed by corpus**, so one machine can hold an active work item for each
repository rather than one globally.

## Constraints

- **Filesystem access is confined to one package** — `internal/guards` fails the
  build otherwise. Today that package is scoped to the backlog root, so reaching
  a location outside the repository widens what it is for, deliberately rather
  than by accident.
- **It must survive `git clean` and not appear in `git status`.** Both follow
  from being outside the working tree, which is the point.
- **Absence is normal.** No active work item is the ordinary state, and every
  command must work without one — this is a shortcut, never a requirement.

## Settled while planning

- **The verb is `pick`.** `select` is spent — `workflow-status.md` says *both
  gates are a selection* — and `take`, `release` and `claim` are ADR-0008's and
  §4.5's. `active` and `current` both name the concept well and are noun-shaped.
- **Closing the picked work item clears the pick**, and nothing else does.
- **`-w` wins over the pick, silently.** An explicit flag disagreeing with the
  active item is a legitimate act, not something to comment on.
- **A pick that no longer resolves is reported and ignored**, never fatal.

## Open

- **Whether a second verb is wanted at all** — most tools of this shape have
  none, because you switch rather than clear.

## References

- [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]] — the session whose
  journal counts the cost.
