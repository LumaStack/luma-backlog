---
type: work-item
key: WORK-0059
title: How ad hoc work should be done
workflow_status: closed
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:54:30Z'}
description: act fast on a vague direction, plow toward a solution taking shortcuts, and work out what the work is while doing it — outcomes and tasks as an afterthought rather than fleshed out first, because specifying up front would produce a worse result when you do not know what you want until you see it. backlog-move is the vehicle.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T03:09:34Z'}
rank: 070.0040.000
closed: {on: 2026-09-09, as: completed, by: 'agent:claude-opus-5/luma-backlog', reason: 'The maintainer read backlog-move and said it was good enough, which is the whole of that outcome''s check; the record-of-itself outcome verified on all three checks once the maintainer ruled that the captured transcript is the full conversation they asked for. Closed as completed rather than delivered — ADR-0007 renamed that word.'}
---

# How ad hoc work should be done

## The problem

**Some work should not be routed through capture, refine and select**, and
nothing says which. The ladder describes a discipline this project mostly does
not follow for small work, and the exception has never been written down.

Two halves, and the second is the stronger claim.

**The gates.** Acting fast on a vague direction, plowing toward a solution and
taking shortcuts, rather than stopping to file a record and cross two gates
before touching anything.

**Outcomes and tasks as an afterthought.** Not skipped — written after, once
there is something to look at. Fleshing them out first would produce a *worse*
result here, not merely a slower one, because **what is wanted is not known
until it is seen.** A specification written before that is a guess, and it
anchors the work to the guess.

**What this produces is a better version of what already exists**, not a list of
work items. Walking through the thing is how it gets fixed. Recording
observations into another work item and then working that item is the proper
approach, and is what should happen with all things equal — this is the case
where they are not.

## Editor mode

**The mode has a name, and the analogy is the argument.** Editing a book: you
cannot come up with the edits before reading, because you do not know what needs
editing until you see it. Enumerating outcomes first is not expensive here, it
is **impossible**.

**And it is dangerous.** Overused, misused, or declared when the truth is
impatience or an attempt to get out from under the system, it is
indistinguishable from making a mess — and it is *more* dangerous than making a
mess plainly, because it arrives with a justification attached.

**So what this has to produce is a test somebody can fail**, not a permission
somebody can claim.

## What it has to answer

- **What the mode is called**, and whether *editor mode* survives contact.
- **When it is acceptable**, and when it should be discouraged.
- **How the guidelines pivot** — what `CLAUDE.md`, the ladder and the procedures
  say differently once this mode exists.
- **How it is told apart from impatience** wearing its clothes.
- **Which shortcuts are safe, and which are always a bad idea.**
- **Whether the system can recognize which mode is in use**, name it on request
  or on recognition, help judge whether it is being abused, and apply different
  guidelines accordingly.

## The brief may not exist either

**Editing a book, and not for length.** *All I know is I do not like it and I
want to make it better* — so there is no criterion to state up front, only
dissatisfaction.

**Dissatisfaction is not the absence of a standard.** It is a standard that is
not articulable yet. Knowing something is wrong before being able to say why is
the ordinary condition of editing, and **a system that demands the why first
forbids the reading that would produce it.**

**So the test is not whether the criterion precedes the work. It is whether it
arrives.** If the editing finishes and nobody can still say what it was for,
that is thrashing — diagnosed after, which is the only time it can be
diagnosed at all. Same structure as
[[work-items/WORK-0032-how-goalpost-fitting-is-discouraged-without-being-prevented]]'s
second rank: written late, and still proven.

## Flexible, and still discouraging

**Asking the user how they want to make it better gets in the way of them making
it better.** That is the constraint, and it rules out a gate: anything that
stops the work to extract a specification destroys the thing it was protecting.

**The system already has the idiom for this — *observed, never refused*.** It is
what `list` does with skips and duplicates, and what
[[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]] specifies for
conditions. Applied here: never block ad hoc work, but name it, and report when
it looks like abuse.

## Declared: this work went around the system

**Written for somebody who was not here.** The maintainer considers this mode
acceptable and ideal; an observer — compliance, leadership, security — may
reasonably consider it a problem. Both can be true, so what is owed is a record
that lets them decide, not a defense.

**What was skipped, as of the move to `in_progress`:**

- **Both selection gates, in one move.** `captured` → `in_progress`. Nobody
  decided this would become work, then separately decided to do it now.
- **No outcomes.** Nothing states what done looks like, so nothing can say this
  is finished or judge whether it drifted.
- **No review, and no second actor.** One maintainer and one agent. Every
  multi-actor check the model implies produced no friction here, which is not
  the same as passing.
- **The capture procedure's corpus check** was skipped for WORK-0060.
- **A journal-line convention was invented** with no decision record behind it.

**Nobody has reviewed any of it.** Same category as
[[work-items/WORK-0003-review-and-audit-the-implementation]], which records the
Go being merged unreviewed for the same reason and calls it out rather than
leaving it implicit.

**Recording is not yet surfacing.** These facts live here and in this work
item's journal, and nothing indexes them —
[[work-items/WORK-0061-surface-work-that-went-around-the-system-to-observers]]
is where that becomes findable. Until then the interim convention is literal
prefixes in the journal: `shortcut taken` for a skipped step,
`hypothesis, untested` for an unproven claim.

## What is being delivered

**Answers to the above**, and a better `backlog-move`.

`backlog-move` is the vehicle because it is the least examined procedure in the
bundle — two commits, written in one pass, never reopened — while carrying both
selection gates, which is where the whole ladder model lives
([[work-items/WORK-0050-review-the-backlog-procedures]]).

## Out of scope

**The other six procedures.** If the method works it is the template for them,
and that is a later decision rather than a widening of this one.

**Changing the ladder.** This finds out what ad hoc work needs. Whether the
model gains a rung, a flag or nothing is downstream.

## Constraints

- **Outcomes arrive after the work, and still get proven.**
  [[work-items/WORK-0032-how-goalpost-fitting-is-discouraged-without-being-prevented]]
  ranks that second of three and allows it explicitly. The condition is the
  whole discipline: written late is fine, written late and never proven is
  goalpost fitting.
- **The journal is heavy, and it is the instrument.** Every shortcut gets a line
  **as it is taken** — what it skipped, before anyone knows whether it worked.
  Written afterwards, every shortcut reads as justified and the experiment
  produces a story rather than evidence.
- **The journal is not a deferral queue.** A line saying *this should be fixed
  later* is the proper approach sneaking back in. If it can be fixed in flight,
  fix it.
- **Every interaction is journaled, not only the notable ones.** Filtering for
  what seems worth keeping would apply the judgment under test to the evidence
  for it. This journal is an **instrument log** and deliberately denser than
  `spec.md` §5.5 asks for anywhere else — a future reader should not copy it as
  the house style.

## Why this is an inquiry that produces no work items

**It is self-demonstrating.** Ad hoc work about ad hoc work, so the journal is
instrument and specimen at once — the record of how this was done *is* the
finding. Nothing else in the corpus has that property.

**The kind does not quite fit and is still the closest.** `inquiry` is defined
as *understanding, and the work items that follow*; no work items follow this
one. The gloss is wrong rather than the classification.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**The project already works this way and has never said when it may.**
Everything done on 2026-09-07 before this record existed — renaming
`backlog-next` to `backlog-rundown`, the version bump, the Risks reformat,
repointing every inbound link — was ad hoc. No work item, no outcomes, no gate,
and a real result. There is no record of the policy that permitted it.

**One risk reported this morning was probably misread**, and this experiment
should settle it. WORK-0031 closing nine tasks while sitting at `prepared` was
reported as a status that had stopped being true. It may instead be ad hoc work
behaving correctly against a ladder that has no rung for it.

**[[work-items/WORK-0029-separate-quick-capture-from-thoughtful-capture]] is the
same question one level down**, already answered for capture alone, and its
principle generalizes: *"a quick capture that produces a better record at the
cost of two extra turns has failed. It is a different procedure with a different
success condition — not the same one with the reading skipped."* It also appears
to be delivered and unclosed.

**One data point exists already.** `backlog-move` was run for the first time
against the real corpus an hour before this record, and immediately produced a
finding — WORK-0057's rank landing at the front of eleven records rather than
the back, which the procedure says should not happen.

## References

- [[work-items/WORK-0050-review-the-backlog-procedures]] — decides what a good
  procedure is; this gets the evidence from one of them.
- [[work-items/WORK-0032-how-goalpost-fitting-is-discouraged-without-being-prevented]]
  — the ordering that permits outcomes after the fact.
- [[work-items/WORK-0029-separate-quick-capture-from-thoughtful-capture]] — the
  same question, for capture.
- `CLAUDE.md` — the bootstrap order, and the three drivers that promote prose to
  a command.
