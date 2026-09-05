---
type: work-item
key: WORK-0021
title: Rank by position rather than by neighbor
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T20:44:00Z'}
---

# Rank by position rather than by neighbor

## The problem

`rank` takes `--before`, `--after`, `--top` and `--bottom` — every one of which names *another record*. People think in ordinals: `docs/design-mvp.md` writes rank as *#1 → #2 → #3 → #4*, and "make this third" is the sentence somebody actually says.

Expressing that today means first finding out what is currently third.

## What is being delivered

Nothing yet. **Deferred, not rejected** — captured so the reasoning survives and the option can be raised again without re-deriving it.

Two forms of the same want:

- **`rank <work-item> --at <n>`** — place at an ordinal position.
- **Applying a whole ordering at once**, which is what an agent wants after re-prioritizing a batch, and which `--at` serves badly anyway.

## Why it was not taken now

**`--at 3` names a slot in a list nobody specified.** Third among all work items, third in the `todo` column, or third in whatever the board is filtered to — all defensible, all different. `--before WORK-0012` names a record and cannot be misread.

**It reintroduces index semantics that `spec.md` §9.6 removed on purpose.** Positions are decimal ordering keys precisely so a move writes one record; a caller thinking in indices is a caller whose intent breaks when somebody else inserts between the read and the write.

**The ordinals people say out loud already have flags.** Speech is "put it first" and "stick it at the bottom" — `--top` and `--bottom`. `--at` buys positions two through n−1, which is the range nobody names.

**The cost is asymmetric.** `spec.md` §9.9 makes additions non-breaking and removals breaking, so shipping without it is free to reverse and shipping with it is not.

## Re-open trigger

- **Rank's scope is settled as something narrower than global** — per column, per work item, per view — which would remove the ambiguity that is the main objection.
- **Real use shows people repeatedly resolving neighbors by hand** to express an ordinal, which is the friction this would remove.
- **A batch re-prioritization workflow appears**, at which point the bulk form is worth designing and `--at` may fall out of it.

## References

- `[[backlog/work-items/WORK-0017-specify-the-minimum-viable-product]]`
- `docs/spec.md` §9.6 — the ordering key and why positions were rejected as storage.
