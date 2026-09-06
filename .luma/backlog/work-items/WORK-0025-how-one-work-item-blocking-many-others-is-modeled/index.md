---
type: work-item
key: WORK-0025
title: How one work item blocking many others is modeled
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T04:00:00Z'}
---

# How one work item blocking many others is modeled

## The problem

**Some work items block most of the others, and there is no honest way to say
so.** Every option has a serious cost, which is why this is an inquiry rather
than a change.

The case is not hypothetical.
[[backlog/work-items/WORK-0018-extract-the-application-layer]] blocks nearly
everything in this backlog right now, and nothing in the corpus records it.

**What exists today says *that*, not *by what*.** `spec.md` §4.2.1 makes
`blocked` a field whose presence means blocked — a marker, deliberately not a
status (§4.1.1) and deliberately not a column (§11.2). `depends_on` is a
relationship, but it lives on the **task** (§4.5), not on the work item. So a
work item can be marked blocked and cannot say what is blocking it.

## The options, and what is wrong with each

**Link it to every other work item.** Impractical at any size. Worse, it decays
silently: every new work item is unblocked until somebody remembers to link it,
and the one that gets missed is invisible precisely because the volume of
correct data hides it.

**Declare that it blocks everything before it, and record exceptions.** Cheap to
state and puts the burden on people to notice where it is wrong. Being wrong is
the default state until somebody catches it.

**Conditional blocking** — tags, kinds, or some predicate saying what is
blocked. Expressive, and it is a rules engine arriving through a side door
(`spec.md` §8.6, `open-questions.md` §6).

**Introduce a container** — epic, project, milestone, initiative — so containers
block containers and the work-item graph stays quiet. Avoids the noise, and adds
a level this project deliberately does not have: ADR-0001 chose a single unit,
and level-legibility was one of the three tests it had to pass.

**A chain** — 1 blocks 2, 2 blocks 3 and 4, 4 blocks 5, 6 and 7. Least data on
disk. Highest churn, hardest to keep true, and if somebody does keep it true the
upkeep may cost more than the answer is worth.

## Two questions that may collapse the rest

**What is blocking actually for?** The options cost very different amounts
depending on the answer, and it has not been asked.

- **To warn somebody off** — *do not start this yet*. Then it can be lossy. A
  missed link costs one person one wasted start, and the cheapest option is
  probably good enough.
- **To compute what is startable** — a scheduling input, consumed by something
  choosing work. Then it must be complete and correct, and every option above
  gets expensive because completeness is the expensive part.

`principles.md` puts *deciding what to work on next* upstream of this tool,
which argues for the first — but the board showing *three of eight in progress
are blocked* (§11.2) is closer to the second.

**Could dimensions already carry this?** §2.7 and §3 make a dimension an
attribute that groups and filters every view, with optional records of its own.
Blocking between dimensions is the container option **without a new level** —
the grouping exists, and nothing new is invented. Worth ruling in or out before
the heavier options are weighed.

## What this produces

A decision, and the work items that follow from it. **Nothing is built from this
work item** — if it concludes that the cheapest lossy option is enough, that is
a complete and successful result.

## Constraints

- **Whatever is chosen is observed, not enforced.** `spec.md` §5.0 permits
  refusing only what the caller's own record contradicts, and blocking is
  advice — `workflow-status.md` is explicit that a wrong refusal teaches people
  to reach for a force flag.
- **It must not become a rules engine** (§8.6). Configuration declares
  vocabulary and bindings, never behavior.
- **Correctness that depends on remembering is not correctness** (`CLAUDE.md`,
  §9a.4). An option whose data is right only while somebody maintains it by hand
  should be judged as an option that will be wrong.

## References

- `docs/spec.md` §4.2.1 — `blocked` as a marker rather than a status.
- `docs/spec.md` §4.5 — `depends_on`, which exists only on tasks.
- `docs/spec.md` §2.7, §3 — dimensions.
- `[[records/decisions/ADR-0001-the-backlog-unit-is-a-work-item]]` — why there is
  one unit and no container above it.
