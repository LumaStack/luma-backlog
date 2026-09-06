---
type: work-item
key: WORK-0040
title: A decision number is reused when its file is absent
workflow_status: unprepared
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T17:40:00Z'}
---

# A decision number is reused when its file is absent

## The problem

**`new decision` allocates by scanning the directory**, so a record missing from
disk is a number the allocator believes is free.

**Observed, not theorised.** On 2026-09-06 five files vanished from the working
tree — including `ADR-0004` — and the next `new decision` issued **ADR-0004
again**. It was caught because the title looked wrong beside a familiar number.
Had it not been, two records would have shared a key, and every citation of
ADR-0004 would have been ambiguous.

## Why it is not the collision already recorded

[[backlog/work-items/WORK-0013-how-two-workstations-avoid-colliding]] and
[[records/decisions/ADR-0003-a-colliding-key-is-repaired-by-appending]] cover
**two actors** allocating at once, on separate machines, resolved at merge.

**This is one actor, one machine, no concurrency.** The sequence is derived from
what is on disk rather than from a high-water mark, so anything that removes a
file — a bad checkout, a failed sync, an interrupted operation, a `.gitignore`
mistake — silently frees a number that was never free.

## What is being delivered

Allocation that does not depend on every prior record being present. Options,
none chosen:

- **A high-water mark**, stored. Cheap, and it is state that can itself drift.
- **Scan git history as well as the working tree** — a number ever issued is
  never reissued. No new state, and it costs a history walk.
- **Detect rather than prevent**: allocate as now, and report a number holding
  two records wherever the corpus is read — which is what
  [[backlog/work-items/WORK-0014-detect-two-records-holding-one-key]] already
  built for work item keys.

**The third is probably right on its own**, because it catches a duplicate
however it arose rather than closing one route to it.

## Constraints

- **Never renumber automatically.** ADR-0003 settled that a colliding key is
  repaired by appending, and a citation of the old number has to keep meaning
  something.
- **Whatever is chosen must survive a partial checkout**, since that is the case
  that produced this.

## References

- `.luma/bundles/local/backlog/procedure/backlog-new.md` — *"do not name the file
  yourself — the number comes from one sequence across the whole project."*
- `docs/spec.md` §6.4 — identifier allocation.
