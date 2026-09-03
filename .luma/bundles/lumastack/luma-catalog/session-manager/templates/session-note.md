# Session note template

Copy the blocks to `~/.local/state/luma/sessions/<project>/<branch>.md`.
**Copy the blocks, not this file.**

Sections are dropped when empty rather than left with *none* under them — a note
is read in a hurry, and an empty heading is a line somebody has to read to learn
nothing.

**Everything here is written for someone who cannot see the conversation.** Paths
from the repository root, commands in full, no pronoun without a referent in the
note itself.

## It must also work for a reader who has none of this

**Nothing is guaranteed to consume a session note.** [[session-resume]] deletes
one — where it has been taken, applied into the running agent, and
remembered. Assume none of that. A note that only makes sense to somebody
holding this bundle is a note that outlives its usefulness and gets believed.

**So the note explains itself, in its own first lines**, before any content:

```markdown
> **Session note.** Working state left by the previous session so this one can
> pick up without re-deriving it. **Read it, act on it, then delete this file.**
> It is a pointer, not a record — everything in it that matters lives somewhere
> durable too. If it is more than a few days old, check its `pinned` values
> against the repository before trusting anything below.
```

**Four words are doing the work: *then delete this file*.** Without them the
note's ephemerality depends on a person remembering, and the failure is silent —
a stale note is trusted precisely because nothing announces it is stale.

## Frontmatter

```yaml
---
type: session_note
kind: checkpoint          # checkpoint | handoff | close
title: <what this session is doing, in one line>
created: { by: agent:<model>, at: 2026-08-20T09:00:00Z }
pinned: { branch: <branch>, commit: <short-sha>, pull_requests: [<numbers>] }
---
```

`pinned` is what lets a later reader tell **still true** from **was true** —
without it, every claim has to be trusted completely or discarded completely.

## Body

```markdown
## Where this is

<Branch, whether the tree is clean, what the current task is.
One paragraph.>

## Next

- [ ] <the next concrete action, and why this rather than something else>
- [ ] <...>

## Done this session

- [x] <what landed, with the path or pull request number>

## Dead ends

<What was tried, what happened, why it was abandoned. Do not skip this
section — it is the most expensive thing lost in a compaction and the
least likely to be recovered.>

- **<what was tried>** — <what happened>, so <why it was dropped>

## Believed, not confirmed

<Anything that explains what has been seen but has not been checked.
Unlabelled, a successor reads all of it as fact.>

## Not done

<Scope skipped, checks not run, files never read. Silence reads as
coverage.>

## Still running

<Background processes, dev servers, worktrees, subagents — each with why
it is still running, or stopped before writing this.>

## Routed

<What was written durably this session and where, so the next checkpoint
does not file it again.>

- <path or destination> — <what went there>

## Nowhere to put it

<Anything worth keeping that has no home in this repository. Usually a
missing bundle, and worth surfacing rather than inventing a directory for.>
```

## For a handoff, add

```markdown
## For the successor

<Who this was written for, from step 1 of the handoff procedure, and what
they can be assumed to already have: tooling, bundles, memory, prior
context.>
```

Order the note for that successor. A post-compaction agent wants the dead ends
first — everything else is recoverable from the repository. A stranger wants
orientation first.

## For a close, drop `## Next` and add

**A close writes no next steps.** Other people, agents and systems come through
before anybody reads it, and a recommended action is unfalsifiable once stale —
it reads as current no matter how wrong it has become. State can be checked
against the repository; a plan cannot.

```markdown
## Problems left behind

<What is broken, how to see it, and what was true before. State, not
instructions. "Some tests fail" is not state — the command, its output,
and the commit where it started failing is.>

## Deliberately abandoned

<What was dropped and why, so nobody restarts it.>

## Learnings

- **applied** — <what was changed, and where>
- **queued** — <what to do next time, written as work rather than as an
  observation>
```

**`## Next` belongs to checkpoint and handoff only.** Both are read within hours
by somebody who can still tell whether the plan holds.

**A close note should not survive.** [[session-close]] drains it into durable
homes and deletes it. If one is worth keeping for audit or study, it moves to
`.luma/records/sessions/<date>-<slug>.md` and becomes a record — dated,
append-only, never edited.
