---
type: work-item
key: WORK-0024
title: Two attempts at the same task cannot both exist
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T03:40:00Z'}
---

# Two attempts at the same task cannot both exist

## The problem

**Two rules in force meet on the same case and disagree.**

`spec.md` §4.6 — succession: when an attempt does not succeed and another
begins, *"a new record is created and links back to the one it follows."*

`spec.md` §9.5 — idempotency: *"creation is idempotent by name… running it twice
with the same name and the same content is a no-op that succeeds and returns the
existing record."*

So `task new "run the migration"` a second time returns the first record.
**Idempotency wins, and a second attempt at an identically-named task cannot be
created.** Retrying is exactly the case §4.6 exists for.

It applies to outcomes too, though less often — §4.4 expects those to be
tightened rather than repeated.

**Nothing is broken today**, because nothing retries yet. This is a fault in the
specification rather than in the binary.

## Why it is parked

The fix is a can of worms and the system does not run yet. **Whatever is chosen
later costs a rename pass**, because a single maintainer holds the whole corpus
— which is the cheapest this will ever be to migrate.

**And §9.5 may be what changes.** The idempotency rule was written before this
case was understood; the resolution may be to narrow it rather than to add
identity. That should be argued when the retry path is actually being built and
somebody can see which rule is load-bearing.

## What was already considered, so it is not re-derived

| Option | Note |
| --- | --- |
| **A short random suffix on every slug** — `run-the-migration-a3f` | Needs no allocation, so concurrent creators cannot collide. Costs three or four characters of noise on the overwhelming majority of records that are never repeated. |
| **A suffix only on collision** | Tidier day to day; a record's name then depends on what else happened to exist when it was created. |
| **A per-work-item counter** | Re-creates the failure `[[records/decisions/ADR-0003-a-colliding-key-is-repaired-by-appending]]` exists to repair — read the maximum and increment is what collides. Scoping it to one work item makes that rarer, not impossible. |
| **A key of its own for tasks and outcomes** | Set aside separately in `[[backlog/work-items/WORK-0023-refer-to-a-record-by-the-path-a-person-would-type]]` — references are always work-item-scoped, so global uniqueness buys nothing. |

**Uniqueness only has to hold within a work item**, whichever is chosen: a
reference always names the work item first.

## Re-open trigger

- **The retry path is built** — anything implementing §4.6's succession hits
  this immediately.
- Somebody needs two attempts at one task, or two outcomes with one title, and
  finds they cannot have them.

## References

- `docs/spec.md` §4.6, §9.5 — the two rules.
- `docs/spec.md` §7.1 — identity is path-based, which is why a frontmatter-only
  identifier would create two identities that can disagree.
