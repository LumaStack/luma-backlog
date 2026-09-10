---
type: work-item
key: WORK-0097
title: An authorization has nowhere to be recorded
description: backlog-move 0.37.0 says leaving preparing is refused until somebody with standing accepts the outcomes. Nothing records that acceptance, so the refusal cannot be enforced and the rule is prose again --- which is the shape that already failed twice.
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T19:06:17Z'}
---

# An authorization has nowhere to be recorded

## The problem

**`backlog-move` 0.37.0 adds a refusal the tool cannot run.** Leaving
`preparing` now requires that somebody with standing accepted the outcomes ---
and nothing anywhere records an acceptance, so the check has nothing to read.

**That is the shape that already failed twice.** A rule with no mechanism is
prose, and `CLAUDE.md` names the driver: measured compliance with prose-only
rules runs far below what a guarantee requires. This exact gate was crossed
without an answer on 2026-09-09 and again on 2026-09-10
([[records/violations/2026-09-10-173823-ladder-run-without-pausing]]), the
second time by an agent that had read the rule in the same turn.

**It is the one authorization a request cannot carry.** Every other one is given
by the crossing being asked for --- naming `in_progress` authorizes preparing,
scheduling and starting. Nobody can approve a definition of done that did not
exist when they spoke, so this is the only gate that must stop somebody, and the
only one with nowhere to put the answer.

## What is being delivered

**A field that records the acceptance, and the gate that reads it.**

Most likely an `accepted: {by, at}` stamp on an outcome, matching `created` and
`modified` exactly --- but the shape is the thing to decide, not assume:

- **On each outcome, or once on the work item?** Per outcome is finer and
  survives an outcome being added after acceptance; per work item is cheaper and
  is what the gate actually asks about.
- **What invalidates it.** An outcome edited after acceptance is no longer the
  thing that was accepted, and nothing would notice.
- **What the tool checks.** Presence, and that the accepting actor is not the
  proposing one --- which is [[work-items/WORK-0098-an-actor-cannot-say-it-has-standing]]'s
  problem, and the reason this can be built before that one is settled: presence
  is checkable now, standing is not.

## Out of scope

- **Whether standing can be proven** ---
  [[work-items/WORK-0098-an-actor-cannot-say-it-has-standing]]. This delivers
  the place to write it; that one makes the value trustworthy.
- **Whether `preparing` is renamed or split** ---
  [[work-items/WORK-0080-preparation-and-definition-may-be-two-different-things]].
  Whatever the rung is called, leaving it needs this.
- **Verification verdicts** ---
  [[work-items/WORK-0072-independent-verdicts-cancel-each-other-out]] is about
  many checkers disagreeing on whether an outcome *held*. This is one acceptance
  that its *definition* is good enough. Same second-party shape, different fact.

## Constraints

- **The doer is not the checker** (ADR-0007). Whatever is written must make a
  self-accepted outcome detectable.
- **Report, never refuse, on read** (`spec.md` §5.2). A record with no
  acceptance is ordinary; only the crossing is refused.
- **`--force` stays available and stays recorded.** There are legitimate reasons
  to go without --- trivial work, an outage --- and the journal entry is what
  lets anybody ask later whether the gate earned its place.
