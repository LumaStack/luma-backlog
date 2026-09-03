---
type: procedure
title: Checkpoint a session
description: Snapshot where the work is so a crash or a compaction costs nothing, without stopping. Use mid-session after something starts working, before anything irreversible, or at a natural seam.
---

# Checkpoint a session

**Work continues immediately afterwards.** That is what separates this from
[[session-handoff]] and [[session-close]], and it sets everything below.

| | |
| --- | --- |
| **reader** | you, minutes from now, possibly with no memory of this conversation |
| **budget** | small. It runs repeatedly inside the context it is protecting |
| **ends with** | a short report and a checklist. Then back to work |

**If a step would cost more than it saves, skip it and write down that you did.**
A checkpoint that eats the context it exists to preserve has done harm.

## When to run one

**After something works that did not work before.** The highest-value moment
available: you have just learned something, and it is the thing most likely to
be lost.

**Before anything irreversible** — a force push, a migration, a mass rename, a
destructive command.

**At a seam** — the suite goes green, a pull request opens, a subtask finishes.

**Before a long autonomous run**, where nobody will be watching if it goes wrong.

### When not to

**Mid-edit or mid-debug.** You do not yet know what you have learned, and
whatever you write will need retracting.

**When nothing has changed since the last one.** Say so and stop — see step 0.

## First: is a forced compaction close?

**A glance, and usually nothing follows from it.** If there is room, run the
steps as written and do not mention the check — see [[context-budget]].

If one is near, two things change. **Order by what cannot be recovered** — dead
ends first, since a summary keeps conclusions and discards refuted paths. And
**write it as a handoff rather than a checkpoint**, because a compaction changes
who your reader is: they will be working from a summary, and a terse note that
assumes shared context will fail them.

## 0. Is there anything to checkpoint?

Read the existing note first:

```sh
cat ~/.local/state/luma/sessions/<project>/<branch>.md 2>/dev/null
```

**If nothing meaningful has changed since it was written, stop and say so.** A
no-op checkpoint should cost one file read and one sentence. Manufacturing work
to look useful is how a cheap habit becomes an expensive one nobody runs.

Reading it first also tells you **what has already been routed**, so this
checkpoint handles only the delta. Without that, a session that checkpoints four
times files the same learning four times.

## 1. Snapshot where the tree is

```sh
git status --short && git log --oneline -3
```

Capture: the branch, whether the tree is clean, the **shape** of any uncommitted
diff — which files, roughly what changed — and any open pull request.

**Do not force a commit.** Half-finished work committed to look tidy pollutes
history and is a cost somebody pays later. Record that the changes are
uncommitted and what they are. Commit only what is coherent on its own.

## 2. Note what is still running

Background processes, dev servers, worktrees you created, subagents in flight.
**A crash is the case this procedure is insurance against**, and orphaned
processes are what it leaves behind.

## 3. Route what is confirmed

Only what you have **observed, tested, or read in the source** — see
[[session-continuity]] for why belief stays in the note. Then
[[where-knowledge-goes]] decides the destination.

- **Records** — a decision made, a finding, an incident.
- **Ideas** — something worth doing that nobody is doing.
- **Journal or log** — narrative, or per-event.
- **Memory** — only what is true of this operator rather than this repository.

**Every one of these is conditional on the practice already existing here.** Do
it if established, skip it if it is not — and **never start one on the way past**,
which creates something nobody agreed to and makes the next session's detection
wrong. The detection rule is in [[where-knowledge-goes]].

**Nothing has a home? Leave it in the note and say so.** Do not invent a
directory. Skips are collected here and reported once at the end of the session,
not at every checkpoint.

## 4. Retract anything an earlier checkpoint got wrong

You have learned things since. If something routed earlier is now known to be
wrong, **append a correction naming what it supersedes.** Never edit the
original — somebody may already have read it.

This is what makes step 3 safe to run repeatedly.

## 5. Defer, do not ask

**Anything ambiguous is parked in the note with a marker**, not raised now.
Five clarifying questions mid-task costs more than the whole checkpoint saves,
and [[session-handoff]] or [[session-close]] will resolve it with time to spare.

The exception is anything that would be *destructive* to guess at. Ask about
that, always.

## 6. The backlog — only if you are looking for work

**Skip this entirely if you are mid-task.** Reading the backlog to refocus is
useful; reading it out of completeness while you already know what you are doing
is pure cost.

Run it when the current task just finished, when you are genuinely unsure what is
next, or when the session has drifted and needs re-anchoring.

## 7. Write the note

Update in place at `~/.local/state/luma/sessions/<project>/<branch>.md` — see
[templates/session-note.md](../templates/session-note.md).

```yaml
---
type: session_note
kind: checkpoint
created: { by: agent:<model>, at: <timestamp> }
pinned: { branch: <branch>, commit: <short-sha>, pull_requests: [<numbers>] }
---
```

**The checklist lives here, not only in the report.** A checklist that exists
only in a message is gone next turn — which is the failure this procedure is
supposed to prevent. Render it for the user from the note; do not keep two
copies.

## 8. Report, briefly

**Not a status meeting.** Do not recap what the user just watched you do. Weight
it toward what they cannot see:

- **Done** and **next**, as checkboxes.
- **What was written durably, and where.** Paths, so they can check.
- **What is uncertain**, named as uncertain.
- **What was skipped for budget** — that is this checkpoint's own choice and
  worth knowing now.

**Not the missing practices.** A project without journaling does not need
telling four times a day. Those go in the note and surface once, at
[[session-handoff]] or [[session-close]].

Then continue working.
