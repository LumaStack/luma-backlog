---
type: work-item
key: WORK-0069
title: The backlog has no unit for a delivery
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T16:24:25Z'}
description: work items are the only grouping unit, so a set of them that constitutes a delivery point or business outcome has nowhere to live. it is a dimension — but a tracked one, because it always carries metadata the system holds, and because people will call it a project, epic, milestone, release or initiative and every one is right for somebody. research what tracked dimensions are, how they are sequenced when wanted, and how a delivery is finished.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T16:26:00Z'}
---

# The backlog has no unit for a delivery

## The problem

**Work items are all we have, so a set of them that adds up to something has to
be another work item.** That is the wrong shape and it shows immediately: *what
do we need before another project can use this* is not a piece of work, it is a
**delivery point** made of eight or ten pieces of work — and today it can only
be recorded as an eleventh.

**Milestones are the obvious word and the wrong one.** A milestone implies
sequence: they are passed in order. **Epics, projects and initiatives do not** —
they can be delivered in any order, or sequenced deliberately when that is
actually true. **The unit wanted is one that can be ordered but is not ordered
by construction.**

## The spec already refers to it and has no name for it

`spec.md` §2.3 defines a wave as one attempt at a set of outcomes, and explains
it as answering *"how many attempts a **delivery** needs."* **The delivery is
the thing this record is about, and nothing defines it.** A wave is an
iteration; the thing being iterated toward is unnamed.

## Why dimensions cannot answer it

§2.7 is explicit, and it names exactly the words in play:

> Dimensions — **projects, epics, milestones, initiatives, phases, releases,
> sprints** — are user-defined and not core to this project. **They carry no
> mechanics of their own.** … For the minimum viable product they classify and
> nothing more.

**A delivery needs mechanics.** It has to be able to say whether it is reached,
what it is waiting on, and whether it comes after another one. A classification
axis cannot hold any of that — `milestone: v1` on eight records says they share
a label, not that a delivery exists, and certainly not that it is finished.

## It is a dimension, but not a plain one

**Settled going in, and it is a third answer rather than either obvious one.**
Not a sixth record type, and not *all dimensions gain mechanics*. **Two classes
of dimension:**

- **Plain** — classifies and nothing more, exactly as §2.7 has it today.
- **Tracked** — always carries metadata the system has to hold, because the
  thing it names has state: is the delivery reached, what is it waiting on,
  does it come after another.

**And the reason it must stay a dimension is the naming.** People will call this
a project, an epic, a milestone, a release, an initiative, a phase — and every
one of them is right for somebody. Making it a type called `milestone` forces a
team that says `release` to rename a built-in, which is exactly the interop
mismatch §2.7 removed when `project` stopped being the unit name.

**This project already has the mechanism for that split**, one level down.
`workflow_status` is configurable vocabulary that *"carries no meaning to the
tool"* (§8), while the ladder's mechanics are the tool's. **A tracked dimension
is to `epic` and `milestone` what `workflow_status` is to `todo` and `closed`:**
the mechanics belong to the tool, the words belong to the project.

**Where each unit sits, now that the shape is clear:**

```
tracked dimension        ← work items live inside this
    work item
        outcome
        task
        wave             ← an attempt at a set of outcomes, inside the work item
```

**A wave is inside a work item; this is what a work item is inside of.** They
are opposite directions from the same record and neither answers for the other.

§2.7 already carries an undecided adjacent question — whether `project`, `epic`
and `milestone` ship as default-defined dimensions — and it is now the same
question as this one rather than a neighbour.

## What it has to answer

- **What the tracked-versus-plain distinction is called**, and how a project
  declares one. The names of individual dimensions are the project's; the name
  of the *class* is the tool's and has to be chosen once.
- **Whether it holds work items, outcomes, or both.** *Business outcome* is one
  of the phrasings, and an outcome is already a type — so whether a delivery is
  a set of work items or a set of outcomes is a real modelling question, not a
  wording one.
- **How it is sequenced when somebody wants sequence**, without sequence being
  the default.
- **How it is finished.** Every other unit here derives its state rather than
  storing it; a delivery presumably does too, but from what.
- **How it maps to what external trackers already do**, since §2.7 requires
  import and export of dimensions to be first-class and this sits beside them.

## Out of scope

**Whether it ships in the first release.** This is research; scheduling it is a
later decision.

## The instance that produced it

*What has to be true before another project can use luma-backlog* — WORK-0031's
remaining tasks, `--open`, publishing the bundle out of `local/`, a released
binary, and migration. Eight or so pieces, ordered only in part, adding up to one
thing that has no record.

## References

- `docs/spec.md` §2.3 — the wave, and the *delivery* it refers to without
  defining.
- `docs/spec.md` §2.7 — dimensions, and why they cannot carry this.
- [[work-items/WORK-0025-how-one-work-item-blocking-many-others-is-modeled]] —
  the other record about relationships between work items.
