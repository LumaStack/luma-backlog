---
type: work-item
key: WORK-0028
title: Whether an event is a kind of work item
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T05:30:00Z'}
---

# Whether an event is a kind of work item

## The problem

If meetings and retrospectives can be work, and work becomes work items, then
**there may be a kind nobody has named** — an `event`.

It matters beyond taxonomy: **an event can block work**, and if the system is to
show that, the data model has to be able to say it.

The specific question is whether an event is **just another inquiry**. Inquiries
today are questions needing answers — spikes, experiments, reviews. They deliver
a report, an understanding, de-risking, an estimate, and usually more work items;
they never deliver the feature. A meeting looks a lot like that.

Or is an event different enough — in what it delivers, in the metadata it
carries, or in how the system should treat it — to stand apart?

## The test a sixth kind has to pass

`_types/work-item` states it:

> **What would earn a sixth.** Something that produces **none** of a fix, an
> answer, a classification, more work, or the work itself.

**On output alone, an event probably fails.** A meeting produces understanding
and more work items, which is exactly what an inquiry produces. A meeting that
decides something produces a **decision**, which is a record type rather than a
kind.

## The precedent that likely settles it

The same passage handles a candidate shaped identically:

> an **incident** is a defect plus urgency, and **urgency is a different axis**

By that reasoning, **an event is work plus a scheduled time, and time is a
different axis.**

And it is the fourth instance of a pattern this project keeps finding.
`workflow-status.md` records two more: *requester data is a second axis, not a
kind*, and *authority is a third axis*. Each time, something that looked like a
kind turned out to be a property that can co-occur with any kind — which is the
test `spec.md` §4.1.1 already applies to `blocked` and `paused`.

**A meeting can be an inquiry, a defect triage, or a request review.** If the
kind is still free to vary, `event` is not a kind.

## What the test does not measure, which is the real question

The documented test asks only **what a record produces**. The maintainer is
asking two further things it is silent on:

- **Different metadata?** An event has a **time**, and nothing else here does.
- **Different system behavior?** This is the sharper one. Work items are
  **ranked**, and rank is *work order*
  ([[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]).
  **A meeting is not in the work order** — its position is imposed by a
  calendar, not chosen. Something that cannot be ranked sits oddly under a model
  where rank is how everything is sequenced.

So this inquiry may be less about events and more about **whether the sixth-kind
test is complete**. If a candidate can differ in behavior while matching on
output, the test as written cannot see it.

## The noise argument

Events are noisier than inquiries, and that alone might justify separating them
— so a view of *what to get done* is not polluted by *what to show up to*.

**That is a filtering need, and filtering has a mechanism.** Dimensions (§2.7,
§3) are attributes that group and filter every view without adding a type or a
kind. If noise is the whole complaint, a dimension answers it. If it is not,
what remains is the behavior question above.

## What this produces

A decision, and whatever follows from it — possibly a sixth kind, possibly an
axis, possibly a change to the test itself. **Concluding that events are
inquiries with a dimension on them is a complete result.**

## Related

[[backlog/work-items/WORK-0026-what-deserves-to-be-a-work-item]] asks whether
meetings should be work items at all. **If that answers *no*, this question does
not arise** — so they are best taken in that order.

[[backlog/work-items/WORK-0025-how-one-work-item-blocking-many-others-is-modeled]]
carries the blocking half. An event that blocks work needs the same mechanism
anything else blocking work needs.

## References

- `.luma/bundles/local/backlog/_types/work-item` — the kinds, and the test.
- `docs/workflow-status.md` — requester and authority as axes rather than kinds.
- `docs/spec.md` §4.1.1 — anything that can be true alongside a value is a
  separate field.
- `docs/spec.md` §2.7, §3 — dimensions.
