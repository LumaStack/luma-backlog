---
type: work-item
key: WORK-00014
title: Detect two records holding one key
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T00:17:27Z'}
---

# Detect two records holding one key

## The problem

Keys are allocated optimistically and collide across branches, which `spec.md` §6.4 says cannot be solved by local means. **Nothing notices.** Two records can hold `WORK-00013` indefinitely, and the handle somebody quotes in a commit stops meaning one record without anyone being told.

A merge does not catch it either. Two work items created on two branches touch different files, so git merges them cleanly and the duplicate arrives quietly.

## What is being delivered

A check that reports two records holding one key. It runs where the tool already walks the corpus, so it costs a comparison rather than a pass.

**Detection, not allocation.** It is independent of which scheme
[[work-items/WORK-00013-how-two-workstations-avoid-colliding]] settles on, works against the
counter that ships today, and is the thing that makes an escape valve usable —
a repair nobody knows is needed does not happen.

## Out of scope

**Repairing the collision.** What a duplicate becomes — a fraction, the next integer, something else — is that inquiry's to answer. Reporting one is useful before it is answered, and choosing the repair here would pre-empt it.

**Refusing to run.** A duplicate key is a real problem and not a reason to stop working: the records are still readable, and everything except the handle still resolves. Say it and continue, which is the posture for everything that is not the one refusal (`spec.md` §5.0).

## Constraints

- Reported on stderr, like the skip reports, so a caller piping a listing is unaffected.
- Silent when there is nothing to say. A warning that fires on ordinary runs is one people learn to scroll past.
- It must name both records, because knowing a collision exists without knowing where is a worse position than not knowing.
