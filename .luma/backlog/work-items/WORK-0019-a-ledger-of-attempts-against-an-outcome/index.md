---
type: work-item
key: WORK-0019
title: A ledger of attempts against an outcome
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T20:14:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T00:40:00Z'}
---

# A ledger of attempts against an outcome

## The problem

**Half of this shipped early.** `asserted` was originally a single field
overwritten on each attempt; it is now an append-only list
([[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]]),
because a replaced field could not represent a second attempt and would have
been the only place in the design that destroys history. That was free — the
append already existed for `verified`.

**What remains is the part that is not free: an attempt is not a thing.** The
list records *who claimed what, when*, and nothing ties those entries to the
effort that produced them. Two assertions three weeks apart, one before a
rewrite and one after, are indistinguishable from two on the same afternoon.

`spec.md` §2.4 calls an outcome's loop convergence — *"an agent may attempt an
outcome many times — attempt, probe, adjust — and those passes are transient"* —
and §2.3 gives the durable unit a name: the **wave**. Until waves exist there is
nowhere to hang the context an attempt had.

**And an attempt in flight is still invisible.** During a second attempt the
last entry reads `failed`, which is stale rather than false. Nothing says
somebody is working on this right now. A wave or a claim (§6.5) would; neither
is in the first release.

## What is being delivered

Nothing yet. Captured so the shipped half is understood as a deliberate part
rather than an accident, and so the remaining half is not re-derived.

**What it would become:** each assertion carries the wave it belonged to, and a
wave carries what the attempt was for. That makes *how many attempts this took*
and *what changed between them* answerable, which is what a bare list cannot do.

## Constraints

- **Additive.** Assertions are already events in a list, so binding one to a
  wave adds a field rather than changing a shape — no contract break
  (`spec.md` §9.9).
- **Waves are the likely trigger.** An attempt ledger and a wave are close to
  the same idea seen from two ends, and building either without the other would
  probably be wrong.

## Re-open trigger

- **Waves ship**, at which point this is mostly already done.
- Somebody needs to know *what changed between attempt two and attempt three*
  and the assertion list cannot say.
- An attempt in flight needs to be visible, and neither claims nor waves have
  arrived to make it so.

## References

- `[[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]]` — the append-only list, and why it shipped early.
- `[[backlog/work-items/WORK-0017-specify-the-minimum-viable-product]]`
- `docs/spec.md` §2.3, §2.4 — waves, and convergence.
