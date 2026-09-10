---
type: work-item
key: WORK-0077
title: How preparation work is tracked
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T16:55:53Z'}
description: 'preparation steps do not get work done — they prepare the work item so work can start. are they tasks, or their own classification, and if their own thing, what are they called? a default list people add to or remove from: establish who must coordinate and when, establish what deliverables are needed to begin, define the work so outcomes are set, verify outcomes are measurable. a step can also name what it blocks, and not always the same work status. three concepts that may or may not be fleshed out here: whether complete and delivered become two different things; whether outcomes should be complete before a work item is prepared — requirement, recommendation or optional, leaning recommendation; and whether plan-building and work-doing are separated by name as well as by kind.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T16:57:24Z'}
---

# How preparation work is tracked

## The problem

**Preparation steps do not get any work done.** They prepare the work item so
that work can start. That is a different thing from the work, and today there is
only one unit for both.

**So: are they tasks, or are they their own classification?**

## What preparation steps look like

**A default list, which people then add to or remove from.** Something like:

- Establish who needs to coordinate, and when.
- Establish what deliverables are necessary in order to begin this work.
- Define this work — outcomes are set.
- Verify outcomes are measurable.

**Another might look nothing like it:**

- Scope this work and break it down into smaller pieces.
- Estimate how long this will take.
- Deliver estimates to XYZ.
- Establish how we will consult with legal on branding choices, and at what
  point in the workflow this must happen — **what does it block?** `in_progress`?
  `todo`? `closed`? Delivery?

**The second list is the one that shows the shape.** A preparation step can name
what it blocks, and what it blocks is not always the same work status.

## Three concepts that may or may not be fleshed out here

**None is settled and any of them may turn out to belong elsewhere.**

**Do we introduce a concept where complete and delivered are two different
things?**

**Do we introduce a concept where outcomes should be completed before work is
considered prepared** — and is that a requirement, a recommendation, or optional
by default? **Leaning recommendation.**

**Do we separate preparing tasks from normal tasks because one is building a
plan and the other is doing the work** — and ideally give the first a different,
distinct name?

## What is being delivered

**A decision on the classification**, and enough of the two concepts above to
know whether they belong here or in records of their own.

## Constraints

- **Configuration may change the words, the count, and the order — not the
  mechanics** (`workflow-status.md`). A default list somebody edits is exactly
  that shape, so a preparation step has to carry no meaning the tool relies on.
- **`spec.md` §5.0** is the test any configurable surface has to pass.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### The argument that these are not tasks already exists, and it is sharp

**`spec.md` §2.5 defines a task as coordination rather than specification** —
*"it exists because actors work concurrently, not because anybody needs a
plan."* On that definition a preparation step qualifies: somebody picks it up
and finishes it.

**But the corpus already names the condition that separates them.**
`when-a-work-item-splits` carries `task.advances-nothing` — a task attached to no
outcome — and treats it as a defect to be fixed by writing the missing outcome
or moving the task elsewhere.

**A preparation step advances no outcome by construction**, because one of its
jobs is to produce the outcomes. *Define this work — outcomes are set* cannot
advance an outcome that does not exist yet. So modeling preparation steps as
tasks makes every one of them permanently satisfy a condition the system treats
as broken, and the diagnosis in that policy would fire on correct records.

**That is the strongest argument for a separate classification**, and it is
stronger than the intuition that started this, because it is checkable.

### The prior analysis, and what it already establishes

**`workflow-status.md` §"Preparing may hold many gates, sequential or parallel"
is this question, worked through** — and the branding example above lands inside
it. It draws product, legal, compliance, security and engineering as gates
*inside* the preparing phase, and settles three things worth not re-deriving:

- **Which gates apply is a fact about the work item, not the repository.** One
  change needs legal signoff and the next does not, so the set is computed per
  record. It calls this *"a larger departure than adding work statuses, because two work
  items in the same repository no longer take the same path."*
- **Most gates should advise rather than stop.** *"Shipping every gate as a
  refusal is the fastest way to teach people to route around the tool."*
- **The gate itself is not modeled.** *"Who owns one, what it requires before
  work may cross, and whether it applies at all are three things that do not
  exist."*

**And it already answers the "what does this block" question structurally.**
Sequential gates fit the ladder as more work statuses — configuration, no code. Parallel
ones do not, because a record holds one `workflow_status` and something waiting
on legal *while* being shaped by engineering is in two states at once. It points
at `spec.md` §4.1.1 and `blocked` as the precedent: a thing that travels
alongside the position rather than being a position in it.

**So if a preparation step can name what it blocks, it is closer to `blocked`
than to `workflow_status`** — which makes
[[work-items/WORK-0065-blocked-is-a-flag-rather-than-a-rung]] the neighbouring
design rather than an unrelated record.

### Complete versus delivered collides with a decision in force

**ADR-0007 chose `completed` over `delivered` deliberately**, and `spec.md`
§5.3.1 states the reason: *"the tool cannot observe a handover and can compute a
count."* So introducing the distinction is a re-open, not a gap.

**The argument for re-opening is already in `spec.md` §2.2**, which is the
tension rather than an oversight: a work item *"is the only unit that draws a
**delivery boundary**. Outcomes state what done means and waves attempt it, but
neither carries the handover to someone outside the work, and neither can say
*this is delivered*. That is a distinct act, and it needs a distinct unit."*

**One document says handover is a distinct act that needs its own unit; the
other says the tool cannot observe it, so the word was dropped.** Both are in
force. That is worth deciding on its own terms and may not belong in this record
at all — see the note on scope below.

### Outcomes before prepared has a record about the failure mode

**The lean toward *recommendation* is consistent with everything around it** —
the advise-rather-than-stop posture above, and `spec.md` §5.4, where a team that
needs the hard version authors it.

**The risk it opens is already tracked.**
[[work-items/WORK-0032-how-goalpost-fitting-is-discouraged-without-being-prevented]]
exists because outcomes get edited to match what was found. Requiring outcomes
at the preparation gate moves the moment they are written earlier, which makes
them more likely to be revised later rather than less — so whichever strength is
chosen, WORK-0032 is where the cost lands.

### On the third concept, and the name

**Plan-building versus work-doing is the sharpest statement of the question**,
and it is the one that survives contact with `spec.md` §2.5. That section
defines a task as coordination — *"how the work of getting there is divided,
ordered, and owned"* — which is a description of doing, not of planning. A step
whose output is the plan itself does not fit the sentence.

**A distinct name is not decoration; it is what stops the two merging back
together.** `spec.md` §11 and the vocabulary rules mean anything named here is
named once and everywhere —
[[records/decisions/ADR-0009-a-symbol-that-must-mean-one-thing-is-assigned-in-one-place]]
is the decision in force, so the name has one assignment point rather than being
introduced separately per surface.

**The naming has a live precedent in this corpus.** WORK-0074's journal records
choosing `pick` and rejecting `select`, `active`, `current`, `take`, `claim`,
`slate`, `cue` and *on deck* — each on the grounds that the surrounding system
had already spent the word. The same test applies here, and the obvious
candidates are already spent: **`stage` is a field**, **`gate` is used for the
transitions in `workflow-status.md`**, **`wave` is an attempt at a set of
outcomes** (`spec.md` §2.3), and **`step` is generic enough to mean nothing**.
That is a constraint on the answer rather than an answer.

### Where this may split

**Three questions are inside this record and they have different answers.**

- **The classification, and the name** — are preparation steps tasks or their
  own thing, and if their own thing, what they are called. This is the question,
  and it is answerable now.
- **Complete versus delivered** — a re-open of ADR-0007, independent of how
  preparation is modeled. Possibly its own record, and it overlaps
  [[work-items/WORK-0069-the-backlog-has-no-unit-for-a-delivery]], which asks for
  a *unit* for a delivery where this asks for a *status*.
- **Outcomes before prepared** — a gate policy, which needs the classification
  decided first but nothing else.

**Kept together deliberately, because splitting before the classification is
decided would produce records that cannot say what done means.** Named here so
the seam is visible when it is time.

## References

- `docs/workflow-status.md` §"Preparing may hold many gates, sequential or
  parallel" — the prior analysis, and the branding example.
- `docs/workflow-status.md` §"What configuration may change" — the words, the
  count and the order; not the mechanics.
- `docs/spec.md` §2.5 — a task is coordination, not specification.
- `docs/spec.md` §2.2 — the work item draws the delivery boundary.
- `docs/spec.md` §5.3.1 — closed is not the same as completed, and why
  `delivered` was not the word.
- [[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]]
  — the decision that chose `completed`.
- [[work-items/WORK-0065-blocked-is-a-flag-rather-than-a-rung]] — the shape a
  step that names what it blocks would take.
- [[work-items/WORK-0032-how-goalpost-fitting-is-discouraged-without-being-prevented]]
  — where the cost of requiring outcomes early lands.
- [[work-items/WORK-0069-the-backlog-has-no-unit-for-a-delivery]] — a unit for a
  delivery, where this touches a status for one.
- [[records/decisions/ADR-0009-a-symbol-that-must-mean-one-thing-is-assigned-in-one-place]]
  — one assignment point for any name this produces.
- [[work-items/WORK-0074-an-active-work-item-remembered-outside-the-repository]]
  — its journal is the worked example of naming against a spent vocabulary.
- [[work-items/WORK-0060-a-configurable-menu-of-working-modes]] — the other
  place a per-project configurable menu is being designed.
