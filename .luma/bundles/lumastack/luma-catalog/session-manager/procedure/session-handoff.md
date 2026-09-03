---
type: procedure
title: Hand off a session
description: Transfer work to another agent or a future session so it resumes without losing what this one learned. Use before a compaction, before exiting, or when someone else takes over — the work is continuing, somewhere else.
---

# Hand off a session

**The work continues; it continues somewhere else.** Somebody is picking this up
and reasonably soon, which is what separates this from [[session-close]] — there,
nobody is.

| | |
| --- | --- |
| **reader** | a specific successor, named in step 1 |
| **budget** | generous. This runs once, and the context is ending anyway |
| **ends with** | a note aimed at that successor, and a prompt for the user to paste |

**Build for the successor.** Knowing who they are is what makes this a handoff
rather than a close, and it changes what you write: what they will already know,
what tools they have, what idiom they read.

## First: is a forced compaction close?

**A glance, and usually nothing follows from it.** If there is room, run the
steps as written and say nothing — see [[context-budget]].

If one is near, this is the procedure most at risk from it. Its budget assumes
the context is ending anyway, which is true — but **a half-finished handoff looks
complete**, and nothing in it announces that the dead ends were never reached.

So **do not begin a full handoff you cannot finish.** Write the note in one pass,
smallest useful version first, then enrich. Order it by what cannot be
recovered: dead ends, then decisions made aloud and never written, then what is
believed rather than confirmed. Where things stand is re-derivable from the
repository; none of those are.

**Step 1 still runs** — knowing the successor is what makes this a handoff, and
it costs one question. Steps 3 and 6 are the ones to compress.

## 1. Ask who is picking this up

Ask first, because the answer shapes everything after it:

> *Who is taking this over — the same agent after a compaction, a different
> model, a different tool, or a person?*

| successor | tailor by |
| --- | --- |
| **the same agent, post-compaction** | terse. They have the repository and the same tooling. Lead with the dead ends, which are what compaction destroys |
| **a different agent or model** | no tool assumptions. Name commands in full; do not assume your bundles, skills, or memory store are present |
| **a person** | plain prose, no frontmatter ceremony, and say what you would do next rather than instructing them |

**If nobody can be named, this is probably a close.** Say so and offer to run
[[session-close]] instead — a handoff without a recipient produces an artifact
aimed at nobody, and it is the weaker of the two.

## 2. Land the tree, with the user

**Never assume what should be committed.** Show them and ask:

```sh
git status --short && git diff --stat
```

**One touchpoint is better than none.** Ask in a single message — *here is what
is uncommitted; commit all of it, some of it, or leave it?* — rather than a
question per file or no question at all. Guessing here is how somebody's work
ends up in a commit whose message does not describe it.

**Committed is not enough.** A commit on a local branch does not survive a
machine boundary, and step 1 may have named a successor on other hardware.

```sh
git push -u origin HEAD
```

If pushing is wrong — the branch is not ready, the work is private — **say so in
the note explicitly**, because a successor elsewhere will otherwise look for a
branch that is not there.

Record: branch name, whether it is pushed, open pull request numbers and their
state.

## 3. Route the durable knowledge

Everything confirmed goes to a real home now — [[where-knowledge-goes]] finds
it. This is the last chance; after this the context is gone.

Read the existing session note first, so you route only what earlier checkpoints
did not.

- **Records** — decisions, findings, incidents.
- **Ideas** — including work you started and are not finishing.
- **Journal or log** — narrative, or per-event.
- **Memory** — only what is true of this operator, not this repository.
- **Documentation** — if you learned how something works and the repository
  documents such things, that is documentation, not a note.

**Each is conditional on the practice existing here.** Do it if established,
skip it if not, and never start one on the way out — see [[where-knowledge-goes]].

**Anything with no home stays in the note, flagged.** An unhoused kind of
knowledge is usually a missing bundle and worth surfacing — in step 8, filtered
to what is actually recommended.

## 4. Write the dead ends

**This is the highest-value paragraph in the whole handoff**, and the one a
compaction reliably destroys — summaries keep conclusions and discard refuted
paths.

For each: what was tried, what happened, and why it was abandoned. *Tried X,
failed because Y* saves the successor the hours it cost you. Without it they will
re-run every one of your failures, confidently.

## 5. Separate what you confirmed from what you believe

Go through the note and mark it. A successor inherits all of it with no way to
tell tested from assumed, and **an assumption read as fact is the failure this
bundle exists to prevent** — nothing announces it, and it becomes the ground
somebody else builds on.

Also write **what you did not do**: scope skipped, checks not run, files never
read. Silence reads as coverage.

## 6. Account for what is still running

Background processes, dev servers, worktrees, subagents. For each: **stop it, or
name it and say why it is still running.**

A handoff may legitimately pass over a running thing — there is somebody to
explain it to. What it may not do is leave one unmentioned.

## 7. Write the note

At `~/.local/state/luma/sessions/<project>/<branch>.md`, or committed to
`.luma/records/sessions/` if the successor is on another machine — see
[[session-continuity]] for the three reasons to commit one.

```yaml
---
type: session_note
kind: handoff
created: { by: agent:<model>, at: <timestamp> }
pinned: { branch: <branch>, commit: <short-sha>, pull_requests: [<numbers>] }
---
```

Use [templates/session-note.md](../templates/session-note.md). Order it for the
successor named in step 1 — for a post-compaction agent that means dead ends
first; for a stranger it means orientation first.

**Then check the invariant.** Read the note and ask of each line: *does this
exist anywhere else?* Anything that does not, route it now. The note is going to
be deleted by whoever consumes it.

### Pin last, and then stop working

**`pinned` is written after everything else, not with the frontmatter.** Steps 8
and 9 still follow, and anything that lands a commit, merges a pull request or
changes a branch between here and the end leaves the note describing a world
that no longer exists.

**A note stale by one minute is worse than one stale by a month.** A reader
checking a month-old pin expects drift and verifies. A reader checking a pin
written minutes ago finds a mismatch and concludes the note is much older than
it is — so they discount the dead ends too, which is the part they most needed.

**And once the note is written, the session is over.** Not *nearly* over —
over. Merging one more pull request, taking one more instruction, landing one
more fix: each of those invalidates the note you just handed over.

If work does continue, **the handoff has not happened yet.** Go back, do the
work, and write the note again. Re-pinning a note whose body is now also wrong
is patching a record instead of making one.

## 8. Report the gaps worth reporting

Once, here — not at every checkpoint.

**Only what is recommended or mandatory.** The catalog's `requires` field carries
the obligation, and that is the filter. A project with no journaling practice
does not need telling; a project with no secrets rule does.

**Report; do not offer to fix.** Name the gap and the command that would close
it, and leave it. Adoption is a durable change to the repository and the end of
a session is the worst moment to decide on one. See [[where-knowledge-goes]].

## 9. Give the user a prompt to paste

The deliverable. It must work as the **first message of a session that has none
of this context**.

Include: which repository and branch, where the note is and to read it first,
what is in flight, what to do next, and any standing expectation about how work
is done here.

**Tell them to run [[session-resume]]** if the successor has this bundle — that
is what consumes the note and deletes it. If they do not have it, say plainly in
the prompt that the note should be deleted once read, or it will be found stale
months later and believed.

Keep it pasteable: no markdown that only renders in one tool, no reference to
anything the successor cannot see.
