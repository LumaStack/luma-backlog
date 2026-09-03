---
type: procedure
title: Resume a session
description: Pick up work another session left — find what it wrote, judge how much of it is still true, act on it, and destroy it. Use at the start of a session that is continuing somebody else's work.
---

# Resume a session

The arriving side of [[session-handoff]], and **the only thing that ever deletes
a session note.** Without this step notes accumulate, go stale, and a stale note
is worse than none: a successor trusts it, and it describes a tree that moved on.

| | |
| --- | --- |
| **run it** | at the start of a session continuing somebody else's work |
| **ends with** | the note consumed and deleted, and the first task started |

## 1. Find the note

```sh
ls ~/.local/state/luma/sessions/<project>/ 2>/dev/null
ls .luma/records/sessions/ 2>/dev/null
```

Keyed by branch, because git guarantees a branch is checked out in exactly one
worktree — so the note for the tree you are standing in is the one that names
your branch. On a detached `HEAD`, look for the short commit.

**Nothing there is a normal outcome.** Say so and orient from the repository:
`git log`, the open pull requests, the readme. Do not go looking through other
branches' notes for something that might be relevant.

## 2. Decide how much to trust it, before reading it as fact

Read the frontmatter first. `kind` and `pinned` are there to make this a
comparison rather than a judgement.

```sh
git log --oneline -1 && git status --short
```

| check | if it fails |
| --- | --- |
| is `pinned.commit` an ancestor of `HEAD`? | the tree moved. Every path and line number is suspect |
| is `pinned.branch` still checked out? | you may be in the wrong tree entirely |
| are `pinned.pull_requests` still open? | *next steps* about them are probably done |
| how old is `created`? | hours: trust it. Months: treat every claim as a lead to verify |

**A `kind: close` note is a red flag.** [[session-close]] deletes its note as its
last act, so one surviving means that close did not finish — and nothing in it
can be assumed to have reached a durable home. Treat every line as unrouted, and
route it in step 4.

## 3. Separate what it confirmed from what it believed

A well-written note marks this. **A note that does not mark it is entirely
belief** — you cannot tell which parts were tested, and assuming is how somebody
else's guess becomes your foundation.

**Read the dead ends before starting anything.** They are why the note exists,
and the whole return on it is not re-running work that already failed.

## 4. Route anything that has no other home

The invariant is that a note is a pointer and never the only copy — but the
session that wrote it may have run out of context before finishing. So check:

**Does anything in here exist nowhere else?** A decision, a finding, an idea, a
learning. If so, give it a durable home now via [[where-knowledge-goes]], before
you touch anything else. It is the last moment anybody will know it was there.

## 5. Confirm the plan before acting on it

**The note's *next* was written by somebody without today's information.** The
pull request may have merged, priorities may have moved, the user may want
something else entirely.

Say what you found and what you intend, in a few lines, and start. **Do not
silently execute a plan written weeks ago** — and do not re-derive it from
scratch either, which throws away the thing you were handed.

**A note with no *next* is not incomplete.** A `kind: close` note deliberately
records state and problems rather than plans, because a plan read after an
unknown gap is unfalsifiable. Take **problems left behind** as verified-at-the-
time observations, not as a queue: **check each against the repository before
believing it.** Something described as broken may have been fixed by whoever came
through in between, and the note has no way to know.

Then form the plan yourself and put it to the user, which is what the close was
counting on.

## 6. Delete the note

Once consumed, and only after step 4:

```sh
rm ~/.local/state/luma/sessions/<project>/<branch>.md
```

**This is the step that makes the note ephemeral.** Skip it and the next session
finds a note describing this one's starting state, believes it, and the error
compounds each time.

A note committed to `.luma/records/sessions/` is a record instead — **do not
delete that one.** It was kept deliberately, for audit or for study, and it is
append-only.

## 7. Start, and checkpoint early

The first [[session-checkpoint]] establishes your own state and confirms the
handoff actually took. **Run one as soon as something works** — the transition
between two sessions is exactly where things get dropped, and it is the moment
you will most regret having no record of.
