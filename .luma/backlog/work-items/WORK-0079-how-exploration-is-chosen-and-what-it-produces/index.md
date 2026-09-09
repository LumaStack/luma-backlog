---
type: work-item
key: WORK-0079
title: How exploration is chosen and what it produces
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T17:25:51Z'}
description: 'select from a menu of exploration strategies — not yet defined — which then guide research, interviewing the user, prototyping, or others. each produces findings: an answer, a better way to measure outcomes, a mockup, well formed requirements. those findings define the outcomes and move the work forward. exploration generally goes two ways: used in preparation it defines THIS work item; used in progress on an inquiry it generates OTHER work items or deliverables. a very strong recommendation rather than a hard rule.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T17:25:51Z'}
---

# How exploration is chosen and what it produces

## The problem

**Exploration happens and nothing chooses how.** Somebody researches, or
prototypes, or asks the user, and which of those it should have been was never a
decision anybody made.

## How it should work

**Select from a menu of exploration strategies** — *not yet defined*. The
strategy then guides the work: research, interviewing the user, prototyping, and
many others.

**Each produces findings.** An answer, a better way to measure outcomes, a
mockup, well-formed requirements — and others. **Those findings are then used to
define the outcomes and move the work forward.**

## Exploration goes two ways

**Both are ordinary, and which one is in play changes what comes out.**

| where | what it produces |
| --- | --- |
| **in preparation** | it defines **this** work item |
| **in progress, on an inquiry** | it generates **other** work items, or deliverables |

**This is a very strong recommendation and not a hard rule.** The two directions
are how exploration is nearly always used, and somebody with a reason to use it
otherwise should not be stopped.

## What is being delivered

**The menu**, and enough of a definition of a finding that a strategy can say
what it produces.

## Constraints

- **A strong recommendation is not a refusal.** Whatever ships here advises and
  lets somebody past — the posture `workflow-status.md` already sets out, where
  *"shipping every gate as a refusal is the fastest way to teach people to route
  around the tool."*

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### The two directions are already in the specification, as the two placements

**`spec.md` §7.2.1 defines exploration and puts it in exactly two places:**
*"It lives in `explorations/` inside the work item it belongs to, **or as a work
item in its own right when the investigation is the work** (§2.1)."*

**Those are the two directions, arrived at independently.** Inside a work item,
exploration serves that work item — direction one. As a work item of its own, it
exists to produce other work — direction two.

**And direction two already has a kind.** `inquiry` is defined as producing
*"understanding, and more work items"*, and *"what comes out is work items, or a
report that generates them. It changes nothing itself, and finding nothing still
counts as done, because the looking was the work."*

**So the model is not new and does not need arguing for.** What is missing sits
on top of it.

### What actually does not exist

**The menu.** Nothing selects a strategy, and nothing records which was used. The
closest thing is `spec.md` §5.2's list of what `preparing` covers —
*"de-risking, estimating, spiking, splitting, checking feasibility, and
coordinating with whoever else is affected"* — which is prose describing
activities rather than a menu anybody picks from. **It is a reasonable seed for
one.**

**The finding.** §7.2.1 models only two endings — *derived into action*, where
an outcome or task is created from the exploration, and *kept as learning*,
where it is archived and stays findable. **Neither is a mockup, a set of
requirements, or a better way to measure an outcome.** Those are products with
shapes of their own, and the spec has no unit for them.

**A typed record for exploration.** `spec.md` §7.2.1 says exploration gets
*"its own directory, and its own type"*, and the directory is real —
`internal/corpus/layout.go:16` registers `exploration`, and `layout.go:37` maps
it to `explorations/`. **The type definition is not there:**
`.luma/bundles/local/backlog/_types/` defines `outcome` and `work-item` and
nothing else, so records are being written as `type: exploration` against a type
nobody declared.

### It is already in use, twice

**Two explorations exist in this corpus**, both written before any of this was
discussed:

- `WORK-0042/explorations/whether-a-task-runner-earns-its-place.md`
- `WORK-0059/explorations/what-the-ladder-walk-established.md`

**Both are direction one** — exploration serving the work item that holds it.
Their frontmatter is `type`, `title`, `work_item`, `stage`, `created`, and
neither records a strategy or names what it produced, **because there is nowhere
to put either.** That is the gap, observed rather than predicted.

### The seam with preparation

[[work-items/WORK-0077-how-preparation-work-is-tracked]] asks what **unit** holds
a preparation step and what it is called. **This asks what the step does and
what comes out of it.** Direction one is the same activity both records are
circling, so the boundary matters: WORK-0077 owns the container, this owns the
strategy and the product.

**Its example list is already exploration by another name** — *establish who
must coordinate and when*, *establish what deliverables are needed to begin*,
*define this work so outcomes are set*. **Those read as strategies, not as
tracking.** Worth checking whether one record should absorb the other rather
than both shipping half a mechanism.

### This is the third menu

**Three records now propose a configurable menu a project selects from:**

- [[work-items/WORK-0060-a-configurable-menu-of-working-modes]] — working modes.
- [[work-items/WORK-0078-how-duplicate-capture-is-handled]] — three capture
  modes.
- This one — exploration strategies.

**Whether that is one mechanism or three is worth deciding before any of them
ships.** Three menus with three configuration surfaces and three sets of rules
about who may add to them is how a tool becomes unlearnable —
[[records/decisions/ADR-0009-a-symbol-that-must-mean-one-thing-is-assigned-in-one-place]]
is the shape of the argument, applied to a mechanism rather than a name.

### On findings feeding outcomes

**Direction one ends by writing or sharpening outcomes, which is where the
`preparing` rung already stops.** `backlog-move` names outcome refinement as the
only thing in that rung with a natural stopping point: *"Scoping and breakdown
can run forever; outcome refinement stops when the outcome passes."*

**So a strategy in direction one has an exit condition already** — it is done
when the outcomes pass their tests. Direction two has no equivalent, because
*finding nothing still counts as done*, and that asymmetry is worth designing
against rather than discovering.

## References

- `docs/spec.md` §7.2.1 — exploration, its directory, its type, both placements,
  and both endings.
- `docs/spec.md` §5.2 — what `preparing` covers; the seed for a strategy menu.
- `internal/corpus/layout.go:16,37` — the exploration type and directory, built.
- `.luma/bundles/local/backlog/_types/` — where the exploration type is not.
- [[work-items/WORK-0077-how-preparation-work-is-tracked]] — the container for
  direction one.
- [[work-items/WORK-0060-a-configurable-menu-of-working-modes]] — the first menu.
- [[work-items/WORK-0078-how-duplicate-capture-is-handled]] — the second.
- [[work-items/WORK-0069-the-backlog-has-no-unit-for-a-delivery]] — deliverables,
  which direction two is said to produce.
