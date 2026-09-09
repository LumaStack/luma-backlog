---
type: work-item
key: WORK-0074
title: An active work item, remembered outside the repository
workflow_status: preparing
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T19:36:09Z'}
description: what somebody is working on right now, so commands stop needing -w on every call. not in git — it is per actor rather than per corpus, and two agents in worktrees would clobber one pointer. per machine to begin with, per user eventually.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T19:37:22Z'}
rank: 030.0010.000
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

## Open

- **What clears it.** Closing the work item is the obvious candidate; so is
  never, leaving a closed record as the last thing you touched.
- **What happens when it names a record that is gone** — deleted, renamed, or in
  a branch you are no longer on. Reporting rather than failing is this project's
  habit.
- **Whether `-w` still wins**, which it should, and whether an explicit flag that
  disagrees with the active item is worth saying out loud.
- **Whether setting it is a command of its own** or a side effect of moving a
  work item to `in_progress`.

## References

- [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]] — the session whose
  journal counts the cost.
