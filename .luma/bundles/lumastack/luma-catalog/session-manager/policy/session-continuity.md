---
type: policy
title: How a session survives its own end
description: The three ways a session ends, who reads what each one leaves behind, and the invariant that makes a session note safe to destroy.
matches: eager
---

# How a session survives its own end

A session ends three ways, and **the only thing that separates them is who reads
what you leave behind, and when.** Everything else — which steps run, how much
they may cost, how the prose is written — falls out of that.

| | reader | what it means | the note |
| --- | --- | --- | --- |
| [[session-checkpoint]] | you, minutes later | *continue here* | stays |
| [[session-handoff]] | a named successor | *continue somewhere else* | handed over |
| [[session-close]] | **a stranger who may be you, having forgotten** | ***we are done*** | **drained, then gone** |

Read the row before writing anything. An agent that writes every note the same
way produces one that is too terse for a stranger and too laborious to run four
times an hour.

**Handoff builds for a successor; close builds for nobody.** Handoff knows who
is next and produces things aimed at them — a note in their idiom, a prompt to
paste, context tailored to what they will and will not already have. Close has
no successor to aim at, so everything it produces has to stand on its own in the
repository, and it spends its effort shutting things down rather than setting
them up.

That is the whole difference, and it is why *handoff with nobody to hand to* is
not close. A handoff that cannot name its successor has lost the thing that
makes it a handoff.

**Why any of this is worth doing** is [[why-this-exists]] — read it when
deciding whether to keep, extend or adopt the bundle, not while using it.

## When one runs: events, not a clock

**A timer would interrupt mid-thought** — the most expensive moment to stop, and
the one where you have least worth saying, because you do not yet know what you
have learned. So each procedure fires on something that has already happened: a
seam, a transfer, an end. **A seam is already a pause.**

**The user's invocation outranks any heuristic.** If they ask for a procedure, run
it; if they declare a mode, take it. `session-close hard-stop` is declared rather
than inferred for exactly this reason, and urgency is not visible from inside a
session anyway.

**A harness may eventually fire one** — a checkpoint before a forced compaction
is the obvious candidate. That hook does not exist, which is the largest known
hole here. See [[context-budget]].

## The invariant: a note is a pointer, never the only copy

**Anything in a session note that would hurt to lose means an earlier step was
skipped.**

That single sentence is what makes the whole design safe. Notes are deleted
routinely — by [[session-resume]] when consumed, by [[session-close]] when
drained — and deletion is only defensible if nothing of value is exclusively
there. Decisions go to the decision record. Ideas go to the backlog. Learnings
go wherever [[where-knowledge-goes]] sends them. What remains in the note is a
resume point: where I am, what is in flight, what I would do next.

It is also a checkable rule rather than an aspiration. Before deleting a note,
read it and ask of each line: *does this exist anywhere else?* If something does
not, route it first. That is a step in both procedures that delete one.

## Confirmed earns a durable home; unconfirmed stays in the note

**A mid-session learning is often a hypothesis that gets falsified an hour
later.** Write it into a decision record and you have committed the project to
something you were wrong about, in an append-only tier where the correction is a
second document rather than an edit.

So the test is confidence, not importance:

- **Confirmed** — you observed it, tested it, or read it in the source. Route it.
- **Believed** — it explains what you have seen and you have not checked. It
  stays in the note, labelled as belief.

**Say which is which.** A successor inherits everything you write with no way to
tell tested from assumed, and an assumption read as fact is the failure mode
this whole bundle exists to prevent: nothing announces it, and it silently
becomes the ground somebody else builds on.

**Retraction is a new entry, never an edit.** When you find that something
recorded earlier is wrong, the fix is to append a correction that names what it
supersedes. Editing a record somebody may already have read is how two people
end up holding different versions of the same fact.

## Where a note lives

**Machine-local, by default:**

```
~/.local/state/luma/sessions/<project>/<branch>.md
```

`XDG_STATE_HOME` is exactly this: state that persists between runs and is not
important enough to back up. Which is precisely true here — by the invariant
above, losing a note costs a resume point and nothing else.

**Not in `.luma/`.** Everything there is committed without exception, and an
uncommitted file in that tree breaks the property the layout depends on. A note
is uncommitted by nature, so it cannot live there.

**Keyed by branch, because git already guarantees the key is unique.** A branch
is checked out in exactly one worktree, so two agents working the same
repository in separate worktrees cannot collide — and neither can read the
other's note by accident. On a detached `HEAD`, use the short commit instead.

## When a note is committed instead

Three reasons, and no others:

**It has to cross a machine boundary.** A successor on different hardware cannot
read your state directory.

**Somebody will audit it.** The exchange needs to be inspectable later.

**You want to study it.** Notes read in aggregate are how this practice gets
better — where sessions actually lose things, which steps get skipped.

Then it goes to `.luma/records/sessions/<date>-<slug>.md`, in the *what
happened* tier, and it is a record from that moment on: dated, append-only,
never edited. **The default is not to commit.** A committed note is a permanent
artifact created to solve a temporary problem, and a directory of them is
something somebody has to prune later.

## Assume nothing will consume it

**[[session-resume]] is what deletes a note — where it exists.** It has to be
taken, applied into whatever agent is running, and actually invoked. Assume
none of those held.

**So every note explains itself in its own first lines**: what it is, that it
should be deleted after use, and that its age is worth checking. Those lines are
the only mechanism that works when the bundle is absent, `apply` failed, or
the successor simply never ran the procedure.

**This is the weakest part of the design and it is structural.** The whole model
rests on notes being consumed and destroyed, and nothing in the format can force
that. A self-describing note narrows the gap; it does not close it. What would
close it is a harness that runs the arriving procedure automatically — see
[[future-hooks]].

## Write for a reader who cannot see the conversation

Every agent gets this wrong by default, because the conversation is still in
front of it while it writes.

**Nothing may depend on shared context.** No *the file we were editing*, no *as
discussed*, no *that approach*, no pronoun without a referent in the note
itself. Name files by path from the repository root. Name commands in full, as
they would be run. Name people and agents as actors — `human:fsmith`,
`agent:opus-5`.

The test is mechanical: **could this be read cold, by someone with the
repository and nothing else?** For a `close` note that is the exit condition.
For a checkpoint it is merely good practice, since the reader is you and you
still remember — but you will remember less than you expect after a compaction.

## What every note carries

Regardless of which procedure wrote it:

- **Where I am** — branch, whether the tree is clean, what the current task is.
- **What is in flight** — half-finished edits, open pull requests, anything
  running.
- **What I would do next**, and why that rather than something else. **Checkpoint
  and handoff only** — see below.
- **The dead ends.** What was tried and did not work. This is the most expensive
  thing lost in a compaction and the least likely to be recovered, because a
  summary keeps conclusions and discards refuted paths. The next agent will
  cheerfully re-run your failures unless you name them.
- **What was not done** — scope skipped, checks not run, files not read. Silence
  reads as coverage.
- **What is believed rather than confirmed.**

## A plan is only good for as long as its reader arrives

**Checkpoint and handoff write next steps. Close does not.**

The difference is the gap. A checkpoint is read minutes later and a handoff
within hours or days, by somebody who can still tell whether the plan holds. A
close is read at an unknown time, after other people, other agents and other
systems have come through — and **a stale plan is unfalsifiable.** A reader
cannot detect that *next: update the gate path* was done in September by
somebody else; it reads as current no matter how wrong it has become.

**State survives the gap; plans do not.** *`tests/test_gate.py` fails at commit
`4a9c1f2`* stays checkable forever — compare it against what is there now and you
learn something either way. So close records **what is true, what is broken, and
what was deliberately abandoned**, and trusts whoever arrives to plan from
accurate state faster than they could audit a stale one.

This is also the honest reading of *stopping hard*. A session cut short mid-task
has left problems, not a plan — and writing the plan it would have followed
dresses an interruption up as a roadmap.

## Cost is a first-class constraint

**A step that costs more context than it saves is a net loss**, and this is not
a marginal concern: these procedures run inside the context they are protecting.

It binds hardest on [[session-checkpoint]], which runs repeatedly during the
work. It barely binds on [[session-close]], which runs once and has nothing left
to protect. Each procedure states its own budget, and when a step would blow it,
the correct move is to **write down that the step was skipped** rather than
skipping it silently.
