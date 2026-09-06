---
type: work-item
key: WORK-0026
title: What deserves to be a work item
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T04:20:00Z'}
---

# What deserves to be a work item

## The problem

**Meetings are work. Handoffs are work. Manual quality assurance is work. Design
reviews are work.** Any honest account of where a week went includes them.

Counting only what gets coded **hides where the time actually goes** — you
cannot see what has to happen before coding starts, or where it is being wasted.
Counting everything **buries the backlog in churn**, and a system that costs
more to maintain than it returns stops being maintained.

**The tool must work at both extremes** — track none of it, track all of it,
neither one degraded. That is not in question.

**What is missing is a recommended balance**: one default pattern that serves
most situations, offered as a starting point people can deliberately depart
from. Assuming such a pattern exists, which is itself part of the question.

## Two tests the design already implies

Neither is written as guidance, and both may answer more of this than expected.

**Does it hand something over?** `spec.md` §2.2 says a work item is *"the only
unit that draws a delivery boundary… a bounded body of work that, once complete,
hands something over."* A design review that produces decisions hands something
over. A status meeting does not.

**Can you state an outcome for it?** Completion here is evidenced, not asserted
(§2.4, `principles.md`). A record with no statable outcome can never be
completed by the arithmetic — only closed. `spec.md` §5.2 already reports
exactly this as `work-item.unarticulated`, so **the tool already flags the
records this inquiry would discourage.** Whether that flag is the whole answer
is worth deciding before anything is built.

## The reframe worth testing first

**A work item is one of six record types, and most of what is listed above has a
cheaper home.** The question may be less *how much do you track* and more *what
does this belong in*:

| The work | Might belong in |
| --- | --- |
| A meeting that decided something | a **decision** — that is the type's whole purpose |
| A meeting that taught something | a **journal** line on the work it affected (§5.5) |
| A design review producing more work | an **inquiry** work item, or an **exploration** |
| Manual quality assurance | **`verify_by` on an outcome** — §4.4.2 explicitly permits an ordered list of steps somebody performs |
| A handoff | a **journal** entry; the session procedures already write these |

If that holds, the recommendation is a routing table rather than a threshold,
and the burden objection mostly dissolves — none of those cost a work item.

**Manual quality assurance is the most interesting case**, because it is real
work that already has a home. `verify_by` is deliberately unconstrained, and
steps a person follows are one of the shapes it was left open for.

## What this produces

A recommendation, and whatever records it turns out to need. **It may conclude
that no single pattern serves most situations**, and that is a complete result —
knowing there is no default worth shipping is worth as much as finding one.

## Constraints

- **A recommendation is not a mechanism.** `principles.md` keeps opinions out of
  the binary and in configuration; this must land as guidance somebody reads,
  never as a check that refuses a record for being the wrong sort of thing
  (§5.0).
- **Both extremes stay first-class.** A team tracking every meeting and a team
  tracking only code must both be served without either being the degraded case.
- **Kinds classify what a record produces** — a fix, an answer, a
  classification, more work, or the work itself. Anything proposed here has to
  survive that test rather than adding a sixth kind
  ([[records/decisions/ADR-0001-the-backlog-unit-is-a-work-item]]).

## References

- `docs/spec.md` §2.2 — the delivery boundary.
- `docs/spec.md` §2.4, §4.4.2 — outcomes, and why `verify_by` is unconstrained.
- `docs/spec.md` §5.2 — `work-item.unarticulated`.
- `docs/spec.md` §5.5 — the journal.
